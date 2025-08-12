package repositories

import (
	_ "embed"
	"github.com/M-kos/wb_level0/internal/db"
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

func (or *OrderRepository) Create() {

}

func (or *OrderRepository) GetById() {

}

func (or *OrderRepository) List() {

}
