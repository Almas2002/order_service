package price_service

import "7kzu-order-service/pkg/proto"

type service struct {
	priceService proto.PriceServicesClient
}

func New(priceService proto.PriceServicesClient) *service {
	return &service{priceService: priceService}
}
