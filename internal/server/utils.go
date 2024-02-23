package server

import (
	kafkaClient "7kzu-order-service/pkg/kafka"
	"7kzu-order-service/pkg/logger"
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafkaConn.kafkaClient")
	}
	s.kafkaConn = kafkaConn
	brokers, err := kafkaConn.Brokers()

	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}
	logger.Infof("kafka connected to brokers: %v", brokers)
	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		logger.Errorf("kafkaConn.Controller %v", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	logger.Infof("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		logger.Errorf("initKafkaTopics.DialContext %v", err)
		return
	}
	defer func() {
		if err = conn.Close(); err != nil {
			logger.Errorf("kafkaConn.Cancel %v", err)
		}

	}()

	logger.Infof("established new kafka controller connection: %s", controllerURI)

	createOrderTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.CreateOrder.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.CreateOrder.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.CreateOrder.ReplicationFactor,
	}

	updateOrderTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UpdateOrder.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UpdateOrder.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UpdateOrder.ReplicationFactor,
	}

	if err = conn.CreateTopics(
		createOrderTopic,
	); err != nil {
		logger.Errorf("kafkaConn.CreateTopic %v", err)
	}

	logger.Infof("kafka topics created or already exists: %+v", []kafka.TopicConfig{createOrderTopic, updateOrderTopic})

}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.CreateOrder.TopicName,
		s.cfg.KafkaTopics.UpdateOrder.TopicName,
	}
}

func (s *server) connectToGrpc(ctx context.Context, address string) (*grpc.ClientConn, error) {
	readerServiceConn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}

	return readerServiceConn, nil
}

func (s *server) connectGrpcServices(ctx context.Context) {
	s.connectToEtaGrpcService(ctx)
	s.connectToPriceGrpcService(ctx)
	s.connectToWeightGrpcService(ctx)
	s.connectToLocationGrpcService(ctx)

}
