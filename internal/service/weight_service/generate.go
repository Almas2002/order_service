package weight_service

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/proto"
	"context"
	"github.com/pkg/errors"
)

func (s *service) GenerateVolume(ctx context.Context, dimensions *model.Dimensions) (*model.Transport, error) {
	res, err := s.client.GenerateWeight(ctx, &proto.DimensionsRequest{
		Weight: dimensions.Weight,
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
	})
	if err != nil {
		return nil, errors.Wrap(err, "weightService.GenerateWeight")
	}

	return &model.Transport{
		Id:     res.GetTransportId(),
		Volume: res.GetVolume(),
		To:     res.GetTo(),
		From:   res.GetFrom(),
	}, err
}
