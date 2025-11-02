package consumer

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/logger"
	"strings"
	"time"
)

type KafkaConsumer struct {
	client sarama.ConsumerGroup
	config *config.Config
	logger logger.Logger
}

func NewKafkaConsumer(config *config.Config, logger logger.Logger) (*KafkaConsumer, error) {
	brokers := make([]string, 0, 1)
	var broker strings.Builder

	broker.WriteString(config.Kafka.Host)
	broker.WriteString(":")
	broker.WriteString(config.Kafka.Port)
	brokers = append(brokers, broker.String())

	kafkaConfig := sarama.NewConfig()

	client, err := sarama.NewConsumerGroup(brokers, config.Kafka.GroupID, kafkaConfig)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		client: client,
		config: config,
		logger: logger,
	}, nil
}

func (k *KafkaConsumer) RunConsume(ctx context.Context, handler sarama.ConsumerGroupHandler) error {
	var retryCount int

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := k.client.Consume(ctx, []string{k.config.Kafka.Topic}, handler)
			if err != nil {
				if ctx.Err() != nil {
					return ctx.Err()
				}

				retryCount++

				if retryCount > k.config.Kafka.MaxRetries {
					return fmt.Errorf("kafka consumer failed after %d retries: %w", k.config.Kafka.MaxRetries, err)
				}

				k.logger.Error("error from consume", "msg", err)
				time.Sleep(time.Duration(retryCount) * time.Second)
				continue
			}

			retryCount = 0
		}
	}
}

func (k *KafkaConsumer) Close() error {
	err := k.client.Close()
	if err != nil {
		return err
	}

	return nil
}
