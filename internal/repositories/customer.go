package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type customerRepository struct {
	db *db.PostgresDB
}

func NewCustomerRepository(db *db.PostgresDB) (*customerRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &customerRepository{db: db}, nil
}

func (c *customerRepository) Create(ctx context.Context, customer *domains.Customer) (int, error) {
	row := c.db.Pool.QueryRow(ctx, createCustomer, customer.FirstName, customer.LastName, customer.Phone, customer.Email)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (c *customerRepository) FindById(ctx context.Context, id int) (*domains.Customer, error) {
	row := c.db.Pool.QueryRow(ctx, findCustomerById, id)

	var customer customerModel
	if err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Email); err != nil {
		return nil, err
	}

	return &domains.Customer{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Phone:     customer.Phone,
		Email:     customer.Email,
	}, nil
}

func (c *customerRepository) FindByPhone(ctx context.Context, phone string) (*domains.Customer, error) {
	row := c.db.Pool.QueryRow(ctx, findCustomerByPhone, phone)
	var customer customerModel
	if err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Email); err != nil {
		return nil, err
	}

	return &domains.Customer{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Phone:     customer.Phone,
		Email:     customer.Email,
	}, nil
}
