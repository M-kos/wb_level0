package repositories

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/models"
)

var (
	//go:embed order-queries/create-address.sql
	createAddress string
	//go:embed order-queries/create-delivery.sql
	createDelivery string
	//go:embed order-queries/create-item.sql
	createItem string
	//go:embed order-queries/create-order.sql
	createOrder string
	//go:embed order-queries/create-order-item.sql
	createOrderItem string
	//go:embed order-queries/create-payment.sql
	createPayment string
	//go:embed order-queries/get-order-by-id.sql
	getOrderById string
	//go:embed order-queries/get-orders-by-limit.sql
	getOrdersByLimit string
)

type OrderRepository struct {
	db *db.PostgresDB
}

func NewOrderRepository(db *db.PostgresDB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) Create(ctx context.Context, order *models.Order) (int, error) {
	tx, err := or.db.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	row := tx.QueryRow(ctx, createAddress, order.Delivery.Zip, order.Delivery.Address, order.Delivery.City)
	var addressId int
	if err := row.Scan(&addressId); err != nil {
		return 0, fmt.Errorf("address create error: %w", err)
	}

	row = tx.QueryRow(ctx, createDelivery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Email, addressId)
	var deliveryId int
	if err := row.Scan(&deliveryId); err != nil {
		return 0, fmt.Errorf("delivery create error: %w", err)
	}

	//(transaction, request_id, currency_id, provider_id, amount, payment_dt, bank_id, delivery_cost,
	//                     goods_total, custom_fee)
	row = tx.QueryRow(ctx, createPayment, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency)
	var addressId int
	if err := row.Scan(&addressId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createAddressQuery, user.Street, user.Zip, cityId)

			if err := row.Scan(&countryId); err != nil {
				return domain.User{}, err
			}
		} else {
			return domain.User{}, err
		}
	}

	row = tx.QueryRow(ctx, createUserQuery, user.Email, user.FullName, addressId)

	err = row.Scan(&user.ID, &user.CreateTime)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
	return 0, nil
}

func (or *OrderRepository) GetById(ctx context.Context, orderId int) (*models.Order, error) {
	return &models.Order{}, nil
}

func (or *OrderRepository) List(ctx context.Context, limit int) ([]*models.Order, error) {
	return nil, nil
}
