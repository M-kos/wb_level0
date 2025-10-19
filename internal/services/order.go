package services

import (
	"context"
	"github.com/M-kos/wb_level0/internal/domains"
)

type OrderRepository interface {
	GetById(ctx context.Context, orderId string) (*domains.Order, error)
	Create(ctx context.Context, order *domains.Order) (int, error)
}

type Cache interface {
	Get(orderId string) (*domains.Order, bool)
	Set(order *domains.Order)
}

type OrderService struct {
	repo  OrderRepository
	cache Cache
}

func NewOrderService(repo OrderRepository, cache Cache) *OrderService {
	return &OrderService{
		repo:  repo,
		cache: cache,
	}
}

func (o *OrderService) GetById(ctx context.Context, orderId string) (*domains.Order, error) {
	order, ok := o.cache.Get(orderId)

	if ok {
		return order, nil
	}

	order, err := o.repo.GetById(ctx, orderId)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderService) Add(ctx context.Context, order *domains.Order) error {
	o.cache.Set(order)

	_, err := o.repo.Create(ctx, order)

	return err
}
