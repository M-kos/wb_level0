package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type localeRepository struct {
	db *db.PostgresDB
}

func NewLocaleRepository(db *db.PostgresDB) (*localeRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &localeRepository{db: db}, nil
}

func (l *localeRepository) Create(ctx context.Context, locale *domains.Locale) (int, error) {
	row := l.db.Pool.QueryRow(ctx, createLocale, locale.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (l *localeRepository) FindById(ctx context.Context, id int) (*domains.Locale, error) {
	row := l.db.Pool.QueryRow(ctx, findLocaleById, id)

	var locale localeModel
	if err := row.Scan(&locale.ID, &locale.Name); err != nil {
		return nil, err
	}

	return &domains.Locale{
		ID:   locale.ID,
		Name: locale.Name,
	}, nil
}

func (l *localeRepository) FindByName(ctx context.Context, name string) (*domains.Locale, error) {
	row := l.db.Pool.QueryRow(ctx, findLocaleByName, name)

	var locale localeModel
	if err := row.Scan(&locale.ID, &locale.Name); err != nil {
		return nil, err
	}

	return &domains.Locale{
		ID:   locale.ID,
		Name: locale.Name,
	}, nil
}
