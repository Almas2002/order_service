package kafka

import (
	"7kzu-order-service/internal/data/dto"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

func (g *consumerGroup) updateOrderConsumerProcess(ctx context.Context, msg kafka.Message, r *kafka.Reader) {
	g.m.UpdateOrderKafkaEvent.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, msg.Headers, "consumerGroup.updateOrderConsumerProcess")
	defer span.Finish()

	order := dto.UpdateOrderDto{}

	if err := json.Unmarshal(msg.Value, &order); err != nil {
		logger.Errorf("updateOrderConsumerProcess.Unmarshal: %v", err)
		g.errorCommitMessage(ctx, msg, r)
		tracing.TraceError(span, err)
		return
	}

	if err := g.s.UpdateOrder(ctx, order.ToModel()); err != nil {
		logger.Errorf("updateOrderConsumerProcess.Unmarshal: %v", err)
		g.errorCommitMessage(ctx, msg, r)
		tracing.TraceError(span, err)
		return
	}

	g.successCommitMessage(ctx, msg, r)
}
