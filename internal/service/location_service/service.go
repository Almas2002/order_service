package location_service

import (
	"7kzu-order-service/pkg/proto"
)

type service struct {
	locationServiceClie proto.LocationServiceClient
}

func New(locationService proto.LocationServiceClient) *service {
	return &service{locationService}
}
