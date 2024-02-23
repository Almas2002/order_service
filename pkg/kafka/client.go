package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(context context.Context, config *Config) (*kafka.Conn, error) {
	return kafka.DialContext(context, "tcp", config.Brokers[0])
}
