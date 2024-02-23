package order_service

import (
	"7kzu-order-service/internal/repository"
	"7kzu-order-service/internal/service"
	"github.com/segmentio/kafka-go"
)

type orderService struct {
	repo            repository.OrderRepository
	locationService service.LocationService
	ETAService      service.ETAService
	priceService    service.PriceService
	weightService   service.WeightService
	producer        *kafka.Writer
}

func New(repo repository.OrderRepository, locationService service.LocationService, etaService service.ETAService, priceService service.PriceService, weightService service.WeightService, producer *kafka.Writer) *orderService {
	return &orderService{repo, locationService, etaService, priceService, weightService, producer}
}
