package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type OrderService struct {
	SuccessKafkaEvent prometheus.Counter
	ErrorKafkaEvent   prometheus.Counter

	CreateOrderKafkaEvent prometheus.Counter
	UpdateOrderKafkaEvent prometheus.Counter
}

func NewClient360ServiceMetrics(serviceName string) *OrderService {
	return &OrderService{
		SuccessKafkaEvent: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_events_total", serviceName),
			Help: "The total number of success kafka events",
		}),
		ErrorKafkaEvent: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_events_total", serviceName),
			Help: "The total number of error kafka events",
		}),
		CreateOrderKafkaEvent: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_order_kafka_event_total", serviceName),
			Help: "The total number of kafka events create order",
		}),
		UpdateOrderKafkaEvent: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_order_kafka_event_total", serviceName),
			Help: "The total number of kafka events update order",
		}),
	}
}
