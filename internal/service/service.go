package service

import (
	"7kzu-order-service/internal/data/model"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
}

type LocationService interface {
	GenerateLocation(ctx context.Context, receiver *model.Address, sender *model.Address) (*model.Coordinate, error)
}

type PriceService interface {
	GeneratePrice(ctx context.Context, distance float32, typeId, cityId uint32) (uint64, error)
}

type WeightService interface {
	GenerateVolume(ctx context.Context, dimensions *model.Dimensions) (*model.Transport, error)
}

type ETAService interface {
	GenerateDistanceAndTime(ctx context.Context, coordinates *model.Coordinate, transportId uint32) (distance float32, duration uint32, err error)
}
