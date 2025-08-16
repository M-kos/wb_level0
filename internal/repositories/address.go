package repositories

import (
	"context"
	"fmt"

	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type addressRepository struct {
	db *db.PostgresDB
}

func NewAddressRepository(db *db.PostgresDB) (*addressRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &addressRepository{db: db}, nil
}

func (ar *addressRepository) Create(ctx context.Context, address domains.Address) (int, error) {
	var city cityModel
	row := ar.db.Pool.QueryRow(ctx, findCityByName, address.City)
	if err := row.Scan(&city.ID, &city.Name, &city.RegionID); err != nil {
		return 0, err
	}

	row = ar.db.Pool.QueryRow(ctx, createAddress, address.Zip, address.Address, city.ID)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *addressRepository) Find(ctx context.Context, id int) (*domains.Address, error) {
	row := ar.db.Pool.QueryRow(ctx, findAddress, id)
	var address addressModel

	if err := row.Scan(&address.ID, &address.Zip, &address.Address, &address.CityID); err != nil {
		return nil, err
	}

	row = ar.db.Pool.QueryRow(ctx, findCityById, address.CityID)
	var city cityModel
	if err := row.Scan(&city.ID, &city.Name, &city.RegionID); err != nil {
		return nil, err
	}

	row = ar.db.Pool.QueryRow(ctx, findRegionById, city.RegionID)
	var region regionModel
	if err := row.Scan(&region.ID, &region.Name); err != nil {
		return nil, err
	}

	return &domains.Address{
		ID:      address.ID,
		Zip:     address.Zip,
		Address: address.Address,
		City: domains.City{
			Name: city.Name,
		},
		Region: domains.Region{
			Name: region.Name,
		},
	}, nil
}
