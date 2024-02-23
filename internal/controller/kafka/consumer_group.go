package kafka

import (
	"7kzu-order-service/internal/config"
	"7kzu-order-service/internal/metric"
	"7kzu-order-service/internal/service"
	"7kzu-order-service/pkg/logger"
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type consumerGroup struct {
	s   service.OrderService
	cfg *config.Config
	m   *metric.OrderService
}

func New(s service.OrderService, cfg *config.Config, m *metric.OrderService) *consumerGroup {
	return &consumerGroup{s, cfg, m}
}

func (g *consumerGroup) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msg, err := r.FetchMessage(ctx)
		if err != nil {
			logger.Error(err)
			return
		}

		logger.Infof("workerID: %d take message with partition %d topic %s", workerID, msg.Partition, msg.Topic)

		switch msg.Topic {
		case g.cfg.KafkaTopics.CreateOrder.TopicName:
			g.createOrderConsumerProcess(ctx, msg, r)
		case g.cfg.KafkaTopics.UpdateOrder.TopicName:
			g.updateOrderConsumerProcess(ctx, msg, r)
		}
	}
}
