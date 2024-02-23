package server

import (
	"7kzu-order-service/internal/config"
	kafkaConsumer "7kzu-order-service/internal/controller/kafka"
	"7kzu-order-service/internal/metric"
	"7kzu-order-service/internal/repository/address_repository"
	"7kzu-order-service/internal/repository/item_repository"
	"7kzu-order-service/internal/repository/order_repository"
	"7kzu-order-service/internal/repository/receiver_repository"
	"7kzu-order-service/internal/repository/sender_repository"
	"7kzu-order-service/internal/service/ETA_service"
	"7kzu-order-service/internal/service/location_service"
	"7kzu-order-service/internal/service/order_service"
	"7kzu-order-service/internal/service/price_service"
	"7kzu-order-service/internal/service/weight_service"
	kafkaClient "7kzu-order-service/pkg/kafka"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/postgres"
	"7kzu-order-service/pkg/proto"
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	cfg             *config.Config
	metric          *metric.OrderService
	kafkaConn       *kafka.Conn
	locationService proto.LocationServiceClient
	ETAService      proto.ETAServiceClient
	priceService    proto.PriceServicesClient
	weightService   proto.WeightServiceClient
}

func New(config *config.Config) *server {
	return &server{cfg: config}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	metrics := metric.NewClient360ServiceMetrics("order_service")

	pg := postgres.NewDB(s.cfg.Postgres)

	adrRepo := address_repository.New(pg)

	orderRepository := order_repository.New(pg, receiver_repository.New(adrRepo, pg), sender_repository.New(adrRepo, pg), item_repository.New(pg))

	s.connectGrpcServices(ctx)

	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:29092"),
		RequiredAcks: kafka.RequireAll,
	}

	orderService := order_service.New(orderRepository,
		location_service.New(s.locationService),
		ETA_service.New(s.ETAService),
		price_service.New(s.priceService),
		weight_service.New(s.weightService),
		writer,
	)

	consumeProcesses := kafkaConsumer.New(orderService, s.cfg, metrics)
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID)

	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), 6, consumeProcesses.ProcessMessages)

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer func() {
		if err := s.kafkaConn.Close(); err != nil {
			logger.Errorf("s.kafkaClose %v", err)
		}

	}()

	//defer func() {
	//	err := closeGrpcServer()
	//	if err != nil {
	//		s.log.WarnMsg("server.closeGrpcServer", err)
	//		return
	//	}
	//}()
	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}

	<-ctx.Done()
	//grpcServer.GracefulStop()
	return nil
}
