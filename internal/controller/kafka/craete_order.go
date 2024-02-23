package kafka

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

func (g *consumerGroup) createOrderConsumerProcess(ctx context.Context, msg kafka.Message, r *kafka.Reader) {
	g.m.CreateOrderKafkaEvent.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, msg.Headers, "consumerGroup.createOrderConsumerProcess")
	defer span.Finish()

	order := model.Order{}

	if err := json.Unmarshal(msg.Value, &order); err != nil {
		logger.Errorf("createOrderConsumerProcess.Unmarshal: %v", err)
		g.errorCommitMessage(ctx, msg, r)
		tracing.TraceError(span, err)
		return
	}
	if err := g.s.CreateOrder(ctx, &order); err != nil {
		logger.Errorf("createOrderConsumerProcess.CreateOrder: %v", err)
		g.errorCommitMessage(ctx, msg, r)
		return
	}

	g.successCommitMessage(ctx, msg, r)
}
