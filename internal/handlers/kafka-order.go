package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/M-kos/wb_level0/internal/domains"
	"github.com/M-kos/wb_level0/internal/dto"
	"github.com/jackc/pgx/v5"
	"strings"
	"time"
)

type Service interface {
	Add(ctx context.Context, order *domains.Order) error
}

type KafkaOrderHandler struct {
	service Service
}

func NewKafkaOrderHandler(service Service) *KafkaOrderHandler {
	return &KafkaOrderHandler{
		service: service,
	}
}

func (koh *KafkaOrderHandler) HandleMessage(ctx context.Context, value []byte) error {
	var order dto.Order
	if err := json.Unmarshal(value, &order); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	err := order.Validate()
	if err != nil {
		return fmt.Errorf("order validation errors: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil
	default:
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

		items := make([]*domains.Item, 0, len(order.Items))
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

		dateCreated, err := time.Parse(time.RFC3339, order.DateCreated)
		if err != nil {
			return fmt.Errorf("invalid date created: %w", err)
		}

		err = koh.service.Add(ctx, &domains.Order{
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
			DateCreated: dateCreated,
			OofShard:    order.OofShard,
		})

		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil
			}
			return fmt.Errorf("failed to add order: %v", err)
		}

		return nil
	}
}
