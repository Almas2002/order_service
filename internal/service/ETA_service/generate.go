package ETA_service

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/proto"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (s *service) GenerateDistanceAndTime(ctx context.Context, coordinates *model.Coordinate, transportId uint32) (distance float32, duration uint32, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ETAService.GenerateDistanceAndTime")
	defer span.Finish()

	res, err := s.ETAServiceClient.GenerateETA(ctx, &proto.GenerateETARequest{
		From: &proto.Coordinate{
			Lat: coordinates.From.Lat,
			Lon: coordinates.From.Lon,
		},
		To: &proto.Coordinate{
			Lat: coordinates.To.Lat,
			Lon: coordinates.To.Lon,
		},
		TransportId: transportId,
	})
	if err != nil {
		return 0, 0, errors.Wrap(err, "ETA_service.GenerateDistanceAndTime")
	}

	return res.GetDistance(), res.GetDuration(), nil
}
