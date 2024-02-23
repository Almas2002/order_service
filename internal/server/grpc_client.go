package server

import (
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/proto"
	"context"
)

func (s *server) connectToEtaGrpcService(ctx context.Context) {
	client, err := s.connectToGrpc(ctx, ":9101")
	if err != nil {
		logger.Fatal("dont connect ETA service")
	}

	s.ETAService = proto.NewETAServiceClient(client)
}

func (s *server) connectToPriceGrpcService(ctx context.Context) {
	client, err := s.connectToGrpc(ctx, ":9100")
	if err != nil {
		logger.Fatal("dont connect price service")
	}

	s.priceService = proto.NewPriceServicesClient(client)
}

func (s *server) connectToWeightGrpcService(ctx context.Context) {
	client, err := s.connectToGrpc(ctx, ":9102")
	if err != nil {
		logger.Fatal("dont connect weight service")
	}

	s.weightService = proto.NewWeightServiceClient(client)
}

func (s *server) connectToLocationGrpcService(ctx context.Context) {
	client, err := s.connectToGrpc(ctx, ":9000")
	if err != nil {
		logger.Fatal("dont connect ETA service")
	}

	s.locationService = proto.NewLocationServiceClient(client)
}
