package handlers

import (
	"context"
	"encoding/json"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/domains"
	"github.com/M-kos/wb_level0/internal/dto"
	"github.com/M-kos/wb_level0/internal/logger"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type OrderService interface {
	GetById(ctx context.Context, id string) (*domains.Order, error)
}

type OrderHandler struct {
	config  *config.Config
	logger  logger.Logger
	service OrderService
}

func (oh *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")
	validate := validator.New()

	if err := validate.Var(orderId, "required"); err != nil {
		oh.logger.Error("[GetOrderById] order id validation err: ", err)
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	select {
	case <-r.Context().Done():
		oh.logger.Error("[GetOrderById] request cancelled by the client")
		http.Error(w, "request cancelled by the client", http.StatusRequestTimeout)
		return
	default:
		order, err := oh.service.GetById(r.Context(), orderId)
		if err != nil {
			oh.logger.Error("[GetOrderById] get order by id err: ", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var name strings.Builder
		name.WriteString(order.Delivery.Customer.FirstName)
		name.WriteString(order.Delivery.Customer.LastName)

		var items []dto.Item

		for _, item := range order.Items {
			items = append(items, dto.Item{
				ChrtID:      item.ChrtID,
				TrackNumber: item.TrackNumber,
				Price:       item.Price,
				Rid:         item.Rid,
				Name:        item.Name,
				Sale:        item.Sale,
				Size:        item.Size,
				TotalPrice:  item.TotalPrice,
				NmID:        item.NmID,
				Brand:       item.Brand.Name,
				Status:      item.Status.Value,
			})

		}

		err = json.NewEncoder(w).Encode(dto.Order{
			OrderUID:    order.OrderUID,
			TrackNumber: order.TrackNumber,
			Entry:       order.Entry,
			Delivery: dto.Delivery{
				Name:    name.String(),
				Phone:   order.Delivery.Customer.Phone,
				Zip:     order.Delivery.Address.Zip,
				City:    order.Delivery.Address.City.Name,
				Address: order.Delivery.Address.Address,
				Region:  order.Delivery.Address.Region.Name,
				Email:   order.Delivery.Customer.Email,
			},
			Payment: dto.Payment{
				Transaction:  order.Payment.Transaction,
				RequestID:    order.Payment.RequestID,
				Currency:     order.Payment.Currency.Name,
				Provider:     order.Payment.Provider.Name,
				Amount:       order.Payment.Amount,
				PaymentDt:    order.Payment.PaymentDt,
				Bank:         order.Payment.Bank.Name,
				DeliveryCost: order.Payment.DeliveryCost,
				GoodsTotal:   order.Payment.GoodsTotal,
				CustomFee:    order.Payment.CustomFee,
			},
			Items:             items,
			Locale:            order.Locale.Name,
			InternalSignature: order.InternalSignature,
			CustomerID:        order.CustomerID,
			DeliveryService:   order.DeliveryService.Name,
			Shardkey:          order.Shardkey,
			SmID:              order.SmID,
			DateCreated:       order.DateCreated.String(),
			OofShard:          order.OofShard,
		})

		if err != nil {
			oh.logger.Error("[GetOrderById] encode response err: ", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func NewOrderHandler(router *http.ServeMux, config *config.Config, logger logger.Logger, service OrderService) {
	handler := &OrderHandler{
		config:  config,
		logger:  logger,
		service: service,
	}

	router.HandleFunc("GET /orders/{orderId}", handler.GetOrderById)
}
