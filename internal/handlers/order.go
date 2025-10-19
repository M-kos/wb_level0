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

	if err := validate.Var(orderId, "required,uuid"); err != nil {
		oh.logger.Error("[GetOrderById] orderRepository id validation err: ", err)
		http.Error(w, "invalid orderRepository id", http.StatusBadRequest)
		return
	}

	orderRepository, err := oh.service.GetById(r.Context(), orderId)
	if err != nil {
		oh.logger.Error("[GetOrderById] get orderRepository by id err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var name strings.Builder
	name.WriteString(orderRepository.Delivery.Customer.FirstName)
	name.WriteString(orderRepository.Delivery.Customer.LastName)

	var items []dto.Item

	for _, item := range orderRepository.Items {
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
		OrderUID:    orderRepository.OrderUID,
		TrackNumber: orderRepository.TrackNumber,
		Entry:       orderRepository.Entry,
		Delivery: dto.Delivery{
			Name:    name.String(),
			Phone:   orderRepository.Delivery.Customer.Phone,
			Zip:     orderRepository.Delivery.Address.Zip,
			City:    orderRepository.Delivery.Address.City.Name,
			Address: orderRepository.Delivery.Address.Address,
			Region:  orderRepository.Delivery.Address.Region.Name,
			Email:   orderRepository.Delivery.Customer.Email,
		},
		Payment: dto.Payment{
			Transaction:  orderRepository.Payment.Transaction,
			RequestID:    orderRepository.Payment.RequestID,
			Currency:     orderRepository.Payment.Currency.Name,
			Provider:     orderRepository.Payment.Provider.Name,
			Amount:       orderRepository.Payment.Amount,
			PaymentDt:    orderRepository.Payment.PaymentDt,
			Bank:         orderRepository.Payment.Bank.Name,
			DeliveryCost: orderRepository.Payment.DeliveryCost,
			GoodsTotal:   orderRepository.Payment.GoodsTotal,
			CustomFee:    orderRepository.Payment.CustomFee,
		},
		Items:             items,
		Locale:            orderRepository.Locale.Name,
		InternalSignature: orderRepository.InternalSignature,
		CustomerID:        orderRepository.CustomerID,
		DeliveryService:   orderRepository.DeliveryService.Name,
		Shardkey:          orderRepository.Shardkey,
		SmID:              orderRepository.SmID,
		DateCreated:       orderRepository.DateCreated,
		OofShard:          orderRepository.OofShard,
	})
	if err != nil {
		oh.logger.Error("[GetOrderById] encode response err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func NewOrderHndler(router *http.ServeMux, config *config.Config, logger logger.Logger, service OrderService) {
	handler := &OrderHandler{
		config:  config,
		logger:  logger,
		service: service,
	}

	router.HandleFunc("GET /orders/{orderId}", handler.GetOrderById)
}
