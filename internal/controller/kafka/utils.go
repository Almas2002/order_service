package kafka

import (
	"7kzu-order-service/pkg/logger"
	"context"
	"github.com/segmentio/kafka-go"
)

func (g *consumerGroup) successCommitMessage(ctx context.Context, msg kafka.Message, r *kafka.Reader) {
	g.m.SuccessKafkaEvent.Inc()
	if err := r.CommitMessages(ctx, msg); err != nil {
		logger.Errorf("consumerGroup.successCommitMessage topic: %s", msg.Topic)
	}
}

func (g *consumerGroup) errorCommitMessage(ctx context.Context, msg kafka.Message, r *kafka.Reader) {
	g.m.ErrorKafkaEvent.Inc()
	if err := r.CommitMessages(ctx, msg); err != nil {
		logger.Errorf("consumerGroup.errorCommitMessage topic: %s", msg.Topic)
	}
}
