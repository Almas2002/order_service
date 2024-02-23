package price_service

import (
	"7kzu-order-service/pkg/proto"
	"context"
	"github.com/pkg/errors"
)

func (s *service) GeneratePrice(ctx context.Context, distance float32, typeId, cityId uint32) (uint64, error) {
	price, err := s.priceService.GeneratePrice(ctx, &proto.GeneratePriceRequest{
		Distance: distance,
		CityId:   cityId,
		TypeId:   typeId,
	})
	if err != nil {
		return 0, errors.Wrap(err, "priceService.GeneratePrice")
	}

	return uint64(price.GetPrice()), nil
}
