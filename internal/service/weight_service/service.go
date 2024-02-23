package weight_service

import "7kzu-order-service/pkg/proto"

type service struct {
	client proto.WeightServiceClient
}

func New(client proto.WeightServiceClient) *service {
	return &service{client: client}
}
