package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type bankRepository struct {
	db *db.PostgresDB
}

func NewBankRepository(db *db.PostgresDB) (*bankRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &bankRepository{db: db}, nil
}

func (b *bankRepository) Create(ctx context.Context, bank *domains.Bank) (int, error) {
	row := b.db.Pool.QueryRow(ctx, createBank, bank.Name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b *bankRepository) FindById(ctx context.Context, id int) (*domains.Bank, error) {
	row := b.db.Pool.QueryRow(ctx, findBankById, id)

	var bank bankModel
	if err := row.Scan(&bank.ID, &bank.Name); err != nil {
		return nil, err
	}

	return &domains.Bank{
		ID:   bank.ID,
		Name: bank.Name,
	}, nil
}

func (b *bankRepository) FindByName(ctx context.Context, name string) (*domains.Bank, error) {
	row := b.db.Pool.QueryRow(ctx, findBankByName, name)
	var bank bankModel
	if err := row.Scan(&bank.ID, &bank.Name); err != nil {
		return nil, err
	}

	return &domains.Bank{
		ID:   bank.ID,
		Name: bank.Name,
	}, nil
}
