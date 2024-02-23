package order_service

import (
	"7kzu-order-service/internal/data/dto"
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/tracing"
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func (s *orderService) CreateOrder(ctx context.Context, order *model.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderService.CreateOrder")
	defer span.Finish()

	coordinates, err := s.locationService.GenerateLocation(ctx, order.Receiver.Address, order.Sender.Address)
	if err != nil {
		return err
	}

	transport, err := s.weightService.GenerateVolume(ctx, order.Dimensions)
	if err != nil {
		return err
	}

	distance, _, err := s.ETAService.GenerateDistanceAndTime(ctx, coordinates, transport.Id)
	if err != nil {
		return err
	}

	price, err := s.priceService.GeneratePrice(ctx, distance, transport.Id, order.Sender.Address.CityId)
	if err != nil {
		return err
	}

	order.Price = price
	order.TransportId = transport.Id

	if err = s.repo.SaveOrder(ctx, order); err != nil {
		return err
	}

	orderDto := dto.Order{
		ID:     order.ID.String(),
		TypeId: transport.Id,
		Volume: transport.Volume,
		Lat:    coordinates.From.Lat,
		Lot:    coordinates.From.Lon,
		From:   3000,
		To:     0,
	}

	marshal, err := json.Marshal(orderDto)
	if err != nil {
		return err
	}
	if !order.IsDraft {
		header := tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context())
		if err = s.producer.WriteMessages(ctx, kafka.Message{
			Topic:   "7kzu_search_courier",
			Value:   marshal,
			Headers: header,
		}); err != nil {
			return errors.Wrap(err, "CreateOrder.WriteMessages")
		}

	}

	return nil
}
