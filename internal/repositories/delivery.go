package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type deliveryRepository struct {
	db *db.PostgresDB
}

func NewDeliveryRepository(db *db.PostgresDB) (*deliveryRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &deliveryRepository{db: db}, nil
}

func (dr *deliveryRepository) Create(ctx context.Context, customerId, addressId int) (int, error) {
	row := dr.db.Pool.QueryRow(ctx, createDelivery, customerId, addressId)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (dr *deliveryRepository) FindById(ctx context.Context, id int) (*domains.Delivery, error) {
	row := dr.db.Pool.QueryRow(ctx, findDeliveryById, id)
	var delivery deliveryModel

	if err := row.Scan(&delivery.ID, &delivery.CustomerID, &delivery.AddressID); err != nil {
		return nil, err
	}

	row = dr.db.Pool.QueryRow(ctx, findCustomerById, delivery.CustomerID)
	var customer customerModel

	if err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Email); err != nil {
		return nil, err
	}

	row = dr.db.Pool.QueryRow(ctx, findAddress, delivery.AddressID)
	var address addressModel

	if err := row.Scan(&address.ID, &address.Zip, &address.Address, &address.CityID); err != nil {
		return nil, err
	}

	row = dr.db.Pool.QueryRow(ctx, findCityById, address.CityID)
	var city cityModel

	if err := row.Scan(&city.ID, &city.Name, &city.RegionID); err != nil {
		return nil, err
	}

	row = dr.db.Pool.QueryRow(ctx, findRegionById, city.RegionID)
	var region regionModel

	if err := row.Scan(&region.ID, &region.Name); err != nil {
		return nil, err
	}

	return &domains.Delivery{
		ID: delivery.ID,
		Customer: domains.Customer{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Phone:     customer.Phone,
			Email:     customer.Email,
		},
		Address: domains.Address{
			ID:      address.ID,
			Zip:     address.Zip,
			Address: address.Address,
			Region: domains.Region{
				Name: region.Name,
			},
			City: domains.City{
				Name: city.Name,
			},
		},
	}, nil
}
