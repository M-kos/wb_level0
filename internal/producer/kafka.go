package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/M-kos/wb_level0/internal/config"
	"log"
	"strings"
	"time"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	config   *config.Config
}

func NewKafkaProducer(config *config.Config) (*KafkaProducer, error) {
	brokers := make([]string, 0, 1)
	var broker strings.Builder

	broker.WriteString(config.Kafka.Host)
	broker.WriteString(":")
	broker.WriteString(config.Kafka.ExternalPort)
	brokers = append(brokers, broker.String())

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Retry.Max = config.Kafka.MaxRetries

	producer, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &KafkaProducer{
		producer: producer,
		config:   config,
	}, nil
}

func (kp *KafkaProducer) SendMessage(key, value []byte) error {
	var lastErr error

	msg := &sarama.ProducerMessage{
		Topic: kp.config.Kafka.Topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	for attempt := 1; attempt <= kp.config.Kafka.MaxRetries; attempt++ {
		_, _, err := kp.producer.SendMessage(msg)
		if err == nil {
			return nil
		}

		lastErr = err
		log.Printf("producer failed to send message (attempt %d): %v", attempt, err)

		if strings.Contains(err.Error(), "connection") || strings.Contains(err.Error(), "timeout") {
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}

		break
	}

	return fmt.Errorf("failed to send message %w", lastErr)
}

func (kp *KafkaProducer) Close() error {
	if err := kp.producer.Close(); err != nil {
		return fmt.Errorf("failed to close producer: %w", err)
	}

	return nil
}
