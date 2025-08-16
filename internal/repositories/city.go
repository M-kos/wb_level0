package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

var ()

type cityRepository struct {
	db *db.PostgresDB
}

func NewCityRepository(db *db.PostgresDB) (*cityRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &cityRepository{db: db}, nil
}

func (c *cityRepository) Create(ctx context.Context, city *domains.City) (int, error) {
	row := c.db.Pool.QueryRow(ctx, createCity, city.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (c *cityRepository) FindById(ctx context.Context, id int) (*domains.City, error) {
	row := c.db.Pool.QueryRow(ctx, findCityById, id)

	var city cityModel
	if err := row.Scan(&city.ID, &city.Name, &city.RegionID); err != nil {
		return nil, err
	}

	return &domains.City{
		ID:   city.ID,
		Name: city.Name,
	}, nil
}

func (c *cityRepository) FindByName(ctx context.Context, name string) (*domains.City, error) {
	row := c.db.Pool.QueryRow(ctx, findCityByName, name)

	var city cityModel
	if err := row.Scan(&city.ID, &city.Name, &city.RegionID); err != nil {
		return nil, err
	}

	return &domains.City{
		ID:   city.ID,
		Name: city.Name,
	}, nil
}
