package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type regionRepository struct {
	db *db.PostgresDB
}

func NewRegionRepository(db *db.PostgresDB) (*regionRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &regionRepository{db: db}, nil
}

func (r *regionRepository) Create(ctx context.Context, region *domains.Region) (int, error) {
	row := r.db.Pool.QueryRow(ctx, createRegion, region.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *regionRepository) FindById(ctx context.Context, id int) (*domains.Region, error) {
	row := r.db.Pool.QueryRow(ctx, findRegionById, id)

	var reg regionModel
	if err := row.Scan(&reg.ID, &reg.Name); err != nil {
		return nil, err
	}

	return &domains.Region{
		ID:   reg.ID,
		Name: reg.Name,
	}, nil
}

func (r *regionRepository) FindByName(ctx context.Context, name string) (*domains.Region, error) {
	row := r.db.Pool.QueryRow(ctx, findRegionByName, name)
	var reg regionModel
	if err := row.Scan(&reg.ID, &reg.Name); err != nil {
		return nil, err
	}

	return &domains.Region{
		ID:   reg.ID,
		Name: reg.Name,
	}, nil
}
