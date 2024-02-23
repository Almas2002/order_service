package ETA_service

import "7kzu-order-service/pkg/proto"

type service struct {
	ETAServiceClient proto.ETAServiceClient
}

func New(ETAServiceClient proto.ETAServiceClient) *service {
	return &service{ETAServiceClient}
}
