package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type brandRepository struct {
	db *db.PostgresDB
}

func NewBrandRepository(db *db.PostgresDB) (*brandRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &brandRepository{db: db}, nil
}

func (b *brandRepository) Create(ctx context.Context, brand *domains.Brand) (int, error) {
	row := b.db.Pool.QueryRow(ctx, createBrand, brand.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b *brandRepository) FindById(ctx context.Context, id int) (*domains.Brand, error) {
	row := b.db.Pool.QueryRow(ctx, findBrandById, id)

	var brand brandModel
	if err := row.Scan(&brand.ID, &brand.Name); err != nil {
		return nil, err
	}

	return &domains.Brand{
		ID:   brand.ID,
		Name: brand.Name,
	}, nil
}

func (b *brandRepository) FindByName(ctx context.Context, name string) (*domains.Brand, error) {
	row := b.db.Pool.QueryRow(ctx, findBrandByName, name)
	var brand brandModel
	if err := row.Scan(&brand.ID, &brand.Name); err != nil {
		return nil, err
	}

	return &domains.Brand{
		ID:   brand.ID,
		Name: brand.Name,
	}, nil
}
