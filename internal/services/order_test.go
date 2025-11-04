package services

import (
	"context"
	"errors"
	"github.com/M-kos/wb_level0/internal/domains"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var order = &domains.Order{
	ID:          1,
	OrderUID:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: domains.Delivery{
		ID: 1,
		Customer: domains.Customer{
			ID:        1,
			FirstName: "Test",
			LastName:  "Testov",
			Phone:     "+9720000000",
			Email:     "test@gmail.com",
		},
		Address: domains.Address{
			ID:      1,
			Zip:     "2639809",
			Address: "Ploshad Mira 15",
			City: domains.City{
				ID:   1,
				Name: "Kiryat Mozkin",
			},
			Region: domains.Region{
				ID:   1,
				Name: "Kraiot",
			},
		},
	},
	Payment: &domains.Payment{
		ID:          1,
		Transaction: "b563feb7b2b84b6test",
		RequestID:   "",
		Currency: domains.Currency{
			ID:   1,
			Name: "USD",
		},
		Provider: domains.Provider{
			ID:   1,
			Name: "USD",
		},
		Amount:    1817,
		PaymentDt: 1637907727,
		Bank: domains.Bank{
			ID:   1,
			Name: "alpha",
		},
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	},
	Items: []*domains.Item{
		{
			ID:          1,
			ChrtID:      9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			Rid:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmID:        2389212,
			Brand: domains.Brand{
				ID:   1,
				Name: "Vivienne Sabo",
			},
			Status: domains.ItemStatus{
				ID:    1,
				Value: 202,
			},
		},
	},
	Locale: domains.Locale{
		ID:   1,
		Name: "en",
	},
	InternalSignature: "",
	CustomerID:        "test",
	DeliveryService: domains.DeliveryService{
		ID:   1,
		Name: "meest",
	},
	Shardkey:    "9",
	SmID:        99,
	DateCreated: time.Time{},
	OofShard:    "1",
}

var ErrorAddOrder = errors.New("add error")

type StubRepository struct {
	store map[string]*domains.Order
}

func (sr *StubRepository) GetById(ctx context.Context, orderId string) (*domains.Order, error) {
	_ = ctx

	if v, ok := sr.store[orderId]; ok {
		return v, nil
	}

	return nil, pgx.ErrNoRows
}

func (sr *StubRepository) Create(ctx context.Context, order *domains.Order) (int, error) {
	if sr.store == nil {
		return 0, ErrorAddOrder
	}
	_ = ctx

	sr.store[order.OrderUID] = order

	return order.ID, nil
}

type StubCache struct {
	store map[string]*domains.Order
}

func (sc *StubCache) Get(orderId string) (*domains.Order, bool) {
	if v, ok := sc.store[orderId]; ok {
		return v, true
	}

	return nil, false
}

func (sc *StubCache) Set(order *domains.Order) {
	if sc.store == nil {
		return
	}

	sc.store[order.OrderUID] = order
}

func TestOrderService(t *testing.T) {
	tests := []struct {
		title         string
		testType      string
		order         *domains.Order
		repo          StubRepository
		cache         StubCache
		expectedErr   error
		expectedValue *domains.Order
	}{
		{
			title:    "success get without cache",
			testType: "get",
			order:    order,
			repo: StubRepository{
				store: map[string]*domains.Order{order.OrderUID: order},
			},
			cache:         StubCache{},
			expectedErr:   nil,
			expectedValue: order,
		},
		{
			title:    "success get with cache",
			testType: "get",
			order:    order,
			repo:     StubRepository{},
			cache: StubCache{
				store: map[string]*domains.Order{order.OrderUID: order},
			},
			expectedErr:   nil,
			expectedValue: order,
		},
		{
			title:         "failed get",
			testType:      "get",
			order:         order,
			repo:          StubRepository{},
			cache:         StubCache{},
			expectedErr:   pgx.ErrNoRows,
			expectedValue: nil,
		},
		{
			title:    "success add",
			testType: "add",
			order:    order,
			repo: StubRepository{
				store: make(map[string]*domains.Order, 1),
			},
			cache:         StubCache{},
			expectedErr:   nil,
			expectedValue: order,
		},
		{
			title:         "failed add",
			testType:      "add",
			order:         order,
			repo:          StubRepository{},
			cache:         StubCache{},
			expectedErr:   ErrorAddOrder,
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			service := NewOrderService(&tt.repo, &tt.cache)

			switch tt.testType {
			case "add":
				err := service.Add(context.Background(), tt.order)

				require.ErrorIs(t, err, tt.expectedErr)
			case "get":
				value, err := service.GetById(context.Background(), tt.order.OrderUID)

				if tt.expectedValue != nil {
					require.Equal(t, *tt.order, *value)

				} else {
					require.Nil(t, value)
				}

				require.ErrorIs(t, err, tt.expectedErr)
			}

		})
	}
}
