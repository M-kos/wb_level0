package services

import (
	"context"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/domains"
	"sync"
	"time"
)

type CacheOrderRepository interface {
	List(ctx context.Context, limit int) ([]*domains.Order, error)
}

type OrderCache struct {
	cache  map[string]*domains.Order
	mu     sync.RWMutex
	config *config.Config
	repo   CacheOrderRepository
}

func NewOrderCache(config *config.Config, repo CacheOrderRepository) *OrderCache {
	return &OrderCache{
		cache:  make(map[string]*domains.Order, config.CacheSize),
		config: config,
		repo:   repo,
	}
}

func (o *OrderCache) LoadCache(ctx context.Context) error {
	orders, err := o.repo.List(ctx, o.config.CacheSize)
	if err != nil {
		return err
	}

	o.mu.Lock()
	defer o.mu.Unlock()
	for _, orderRepository := range orders {
		o.cache[orderRepository.OrderUID] = orderRepository
	}

	return nil
}

func (o *OrderCache) Get(orderId string) (*domains.Order, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	orderRepository, ok := o.cache[orderId]

	return orderRepository, ok
}

func (o *OrderCache) Set(order *domains.Order) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if len(o.cache) < o.config.CacheSize {
		o.cache[order.OrderUID] = order
		return
	}

	if _, ok := o.cache[order.OrderUID]; ok {
		o.cache[order.OrderUID] = order
		return
	}

	var minDate time.Time
	var idMinDate string

	for _, orderRepository := range o.cache {
		if orderRepository.DateCreated.Before(minDate) {
			minDate = orderRepository.DateCreated
			idMinDate = orderRepository.OrderUID
		}
	}

	if idMinDate != "" {
		delete(o.cache, idMinDate)
	}

	o.cache[order.OrderUID] = order
}

func (o *OrderCache) Delete(orderId string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	delete(o.cache, orderId)
}
