package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type currencyRepository struct {
	db *db.PostgresDB
}

func NewCurrencyRepository(db *db.PostgresDB) (*currencyRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &currencyRepository{db: db}, nil
}

func (c *currencyRepository) Create(ctx context.Context, cur *domains.Currency) (int, error) {
	row := c.db.Pool.QueryRow(ctx, createCurrency, cur.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (c *currencyRepository) FindById(ctx context.Context, id int) (*domains.Currency, error) {
	row := c.db.Pool.QueryRow(ctx, findCurrencyById, id)

	var cur currencyModel
	if err := row.Scan(&cur.ID, &cur.Name); err != nil {
		return nil, err
	}

	return &domains.Currency{
		ID:   id,
		Name: cur.Name,
	}, nil
}

func (c *currencyRepository) FindByName(ctx context.Context, name string) (*domains.Currency, error) {
	row := c.db.Pool.QueryRow(ctx, findCurrencyByName, name)
	var cur currencyModel
	if err := row.Scan(&cur.ID, &cur.Name); err != nil {
		return nil, err
	}

	return &domains.Currency{
		ID:   cur.ID,
		Name: cur.Name,
	}, nil
}
