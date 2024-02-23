package kafka

import (
	"7kzu-order-service/pkg/logger"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"sync"
)

type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type ConsumerGroup interface {
	ConsumeTopic(ctx context.Context, groupId, topic string, poolSize int, worker Worker)
	GetNewKafkaReader(kafkaURL []string, topic, groupID string) *kafka.Reader
	GetNewKafkaWriter(topic string) *kafka.Writer
}

type consumerGroup struct {
	Brokers []string
	GroupID string
}

func NewConsumerGroup(brokers []string, groupID string) *consumerGroup {
	return &consumerGroup{Brokers: brokers, GroupID: groupID}
}

func (c *consumerGroup) GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

// GetNewKafkaWriter create new kafka producer
func (c *consumerGroup) GetNewKafkaWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(c.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}

	return w
}

func (c *consumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	r := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			logger.Error(err)
		}
	}()

	logger.Info("Starting consumer groupID: %s, topic: %+v, pool size: %v")

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
