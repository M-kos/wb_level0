package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type deliveryServiceRepository struct {
	db *db.PostgresDB
}

func NewDeliveryServiceRepository(db *db.PostgresDB) (*deliveryServiceRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &deliveryServiceRepository{db: db}, nil
}

func (ds *deliveryServiceRepository) Create(ctx context.Context, dService *domains.DeliveryService) (int, error) {
	row := ds.db.Pool.QueryRow(ctx, createDeliveryService, dService.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ds *deliveryServiceRepository) FindById(ctx context.Context, id int) (*domains.DeliveryService, error) {
	row := ds.db.Pool.QueryRow(ctx, findDeliveryServiceById, id)

	var dService deliveryServiceModel
	if err := row.Scan(&dService.ID, &dService.Name); err != nil {
		return nil, err
	}

	return &domains.DeliveryService{
		ID:   dService.ID,
		Name: dService.Name,
	}, nil
}

func (ds *deliveryServiceRepository) FindByName(ctx context.Context, name string) (*domains.DeliveryService, error) {
	row := ds.db.Pool.QueryRow(ctx, findDeliveryServiceByName, name)

	var dService deliveryServiceModel
	if err := row.Scan(&dService.ID, &dService.Name); err != nil {
		return nil, err
	}

	return &domains.DeliveryService{
		ID:   dService.ID,
		Name: dService.Name,
	}, nil
}
