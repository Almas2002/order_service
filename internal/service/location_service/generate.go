package location_service

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/proto"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (s *service) GenerateLocation(ctx context.Context, receiver *model.Address, sender *model.Address) (*model.Coordinate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationService.GenerateLocation")
	defer span.Finish()

	coordinates, err := s.locationServiceClie.GetLocationsByAddress(ctx, &proto.GetLocationByAddressRequest{
		From: &proto.LocationAddress{
			CityName: "Алматы",
			Address:  sender.Address,
		},
		To: &proto.LocationAddress{
			CityName: "Алматы",
			Address:  receiver.Address,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "GenerateLocation.GetLocationsByAddress")
	}

	return &model.Coordinate{
		From: model.Point{
			Lat: coordinates.From.GetLat(),
			Lon: coordinates.From.GetLon(),
		},
		To: model.Point{
			Lat: coordinates.To.GetLat(),
			Lon: coordinates.To.GetLon(),
		},
	}, err
}
