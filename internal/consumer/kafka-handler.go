package consumer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/logger"
	"strings"
	"time"
)

type HandlerService interface {
	HandleMessage(ctx context.Context, value []byte) error
}

type KafkaHandler struct {
	service  HandlerService
	logger   logger.Logger
	config   *config.Config
	producer sarama.SyncProducer
}

func NewConsumerHandler(service HandlerService, config *config.Config, logger logger.Logger) (*KafkaHandler, error) {
	kConfig := sarama.NewConfig()
	kConfig.Producer.RequiredAcks = sarama.WaitForAll
	kConfig.Producer.Return.Successes = true

	brokers := make([]string, 0, 1)
	var broker strings.Builder

	broker.WriteString(config.BrokerHost)
	broker.WriteString(":")
	broker.WriteString(config.Kafka.Port)
	brokers = append(brokers, broker.String())

	producer, err := sarama.NewSyncProducer(brokers, kConfig)
	if err != nil {
		return nil, err
	}

	return &KafkaHandler{
		service:  service,
		config:   config,
		logger:   logger,
		producer: producer,
	}, nil
}

func (kh *KafkaHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (kh *KafkaHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (kh *KafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				kh.logger.Error("message channel was closed")
				return nil
			}

			kh.logger.Info("message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

			err := kh.service.HandleMessage(session.Context(), message.Value)
			if err != nil {
				kh.logger.Error("error handling message", "Error", err)

				if kh.config.Kafka.DlqTopic != "" {
					dlqMsg := &sarama.ProducerMessage{
						Topic: kh.config.Kafka.DlqTopic,
						Key:   sarama.ByteEncoder(message.Key),
						Value: sarama.ByteEncoder(message.Value),
					}
					for attempt := 1; attempt <= kh.config.Kafka.MaxDlqRetries; attempt++ {
						_, _, pErr := kh.producer.SendMessage(dlqMsg)
						if pErr == nil {
							break
						}

						kh.logger.Error("failed to send to DLQ: %v", pErr)

						if strings.Contains(pErr.Error(), "connection") || strings.Contains(pErr.Error(), "timeout") {
							time.Sleep(time.Duration(attempt) * time.Second)
							continue
						}

						break
					}
				}
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}

func (kh *KafkaHandler) Close() error {
	return kh.producer.Close()
}
