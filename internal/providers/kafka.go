package providers

import (
	"context"
	"encoding/json"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/domains"
	"github.com/M-kos/wb_level0/internal/dto"
	"github.com/M-kos/wb_level0/internal/logger"
	"github.com/segmentio/kafka-go"
	"strings"
)

type OrderService interface {
	Add(ctx context.Context, order *domains.Order) error
}

type KafkaConsumer struct {
	reader      *kafka.Reader
	config      *config.Config
	logger      logger.Logger
	oderService OrderService
}

func NewKafkaConsumer(config *config.Config, logger logger.Logger, oderService OrderService) *KafkaConsumer {
	return &KafkaConsumer{
		config:      config,
		logger:      logger,
		oderService: oderService,
	}
}

func (k *KafkaConsumer) Consume(ctx context.Context) {
	var url strings.Builder

	url.WriteString(k.config.Kafka.Host)
	url.WriteString(":")
	url.WriteString(k.config.Kafka.Port)

	k.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{url.String()},
		Topic:    k.config.Kafka.Topic,
		GroupID:  k.config.Kafka.GroupID,
		MinBytes: k.config.Kafka.MinBytes,
		MaxBytes: k.config.Kafka.MaxBytes,
	})

	defer func() {
		err := k.reader.Close()
		if err != nil {
			k.logger.Error("failed to close kafka: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := k.reader.ReadMessage(ctx)
			if err != nil {
				k.logger.Error("error reading message: %v", err)
				continue
			}

			var order dto.Order
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				k.logger.Error("invalid JSON: %v", err)
				continue
			}

			names := strings.Split(order.Delivery.Name, " ")

			var firstName string
			var lastName string

			if len(names) == 2 {
				firstName = names[0]
				lastName = names[1]
			}

			if len(names) == 1 {
				firstName = names[0]
			}

			items := make([]*domains.Item, len(order.Items))
			for _, item := range order.Items {
				items = append(items, &domains.Item{
					ChrtID:      item.ChrtID,
					TrackNumber: item.TrackNumber,
					Price:       item.Price,
					Rid:         item.Rid,
					Name:        item.Name,
					Sale:        item.Sale,
					Size:        item.Size,
					TotalPrice:  item.TotalPrice,
					NmID:        item.NmID,
					Brand: domains.Brand{
						Name: item.Brand,
					},
					Status: domains.ItemStatus{
						Value: item.Status,
					},
				})
			}

			err = k.oderService.Add(ctx, &domains.Order{
				OrderUID:    order.OrderUID,
				TrackNumber: order.TrackNumber,
				Entry:       order.Entry,
				Delivery: domains.Delivery{
					Customer: domains.Customer{
						FirstName: firstName,
						LastName:  lastName,
						Phone:     order.Delivery.Phone,
						Email:     order.Delivery.Email,
					},
					Address: domains.Address{
						Zip:     order.Delivery.Zip,
						Address: order.Delivery.Address,
						City: domains.City{
							Name: order.Delivery.City,
						},
						Region: domains.Region{
							Name: order.Delivery.Region,
						},
					},
				},
				Payment: &domains.Payment{
					Transaction: order.Payment.Transaction,
					RequestID:   order.Payment.RequestID,
					Currency: domains.Currency{
						Name: order.Payment.Currency,
					},
					Provider: domains.Provider{
						Name: order.Payment.Provider,
					},
					Amount:    order.Payment.Amount,
					PaymentDt: order.Payment.PaymentDt,
					Bank: domains.Bank{
						Name: order.Payment.Bank,
					},
					DeliveryCost: order.Payment.DeliveryCost,
					GoodsTotal:   order.Payment.GoodsTotal,
					CustomFee:    order.Payment.CustomFee,
				},
				Items: items,
				Locale: domains.Locale{
					Name: order.Locale,
				},
				InternalSignature: order.InternalSignature,
				CustomerID:        order.CustomerID,
				DeliveryService: domains.DeliveryService{
					Name: order.DeliveryService,
				},
				Shardkey:    order.Shardkey,
				SmID:        order.SmID,
				DateCreated: order.DateCreated,
				OofShard:    order.OofShard,
			})
			if err != nil {
				k.logger.Error("failed to add order: %v", err)
			}
		}
	}
}
