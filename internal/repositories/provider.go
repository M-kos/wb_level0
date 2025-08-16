package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type providerRepository struct {
	db *db.PostgresDB
}

func NewProviderRepository(db *db.PostgresDB) (*providerRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &providerRepository{db: db}, nil
}

func (p *providerRepository) Create(ctx context.Context, provider *domains.Provider) (int, error) {
	row := p.db.Pool.QueryRow(ctx, createProvider, provider.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *providerRepository) FindById(ctx context.Context, id int) (*domains.Provider, error) {
	row := p.db.Pool.QueryRow(ctx, findProviderById, id)

	var provider providerModel
	if err := row.Scan(&provider.ID, &provider.Name); err != nil {
		return nil, err
	}

	return &domains.Provider{
		ID:   provider.ID,
		Name: provider.Name,
	}, nil
}

func (p *providerRepository) FindByName(ctx context.Context, name string) (*domains.Provider, error) {
	row := p.db.Pool.QueryRow(ctx, findProviderByName, name)

	var provider providerModel
	if err := row.Scan(&provider.ID, &provider.Name); err != nil {
		return nil, err
	}

	return &domains.Provider{
		ID:   provider.ID,
		Name: provider.Name,
	}, nil
}
