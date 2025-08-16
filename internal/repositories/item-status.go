package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type itemStatusRepository struct {
	db *db.PostgresDB
}

func NewItemStatusRepository(db *db.PostgresDB) (*itemStatusRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &itemStatusRepository{db: db}, nil
}

func (is *itemStatusRepository) Create(ctx context.Context, itemStatus *domains.ItemStatus) (int, error) {
	row := is.db.Pool.QueryRow(ctx, createItemStatus, itemStatus.Value)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (is *itemStatusRepository) FindById(ctx context.Context, id int) (*domains.ItemStatus, error) {
	row := is.db.Pool.QueryRow(ctx, findItemStatusById, id)

	var itemStatus itemStatusModel
	if err := row.Scan(&itemStatus.ID, &itemStatus.Value); err != nil {
		return nil, err
	}

	return &domains.ItemStatus{
		ID:    itemStatus.ID,
		Value: itemStatus.Value,
	}, nil
}

func (is *itemStatusRepository) FindByName(ctx context.Context, name string) (*domains.ItemStatus, error) {
	row := is.db.Pool.QueryRow(ctx, findItemStatusByValue, name)

	var itemStatus itemStatusModel
	if err := row.Scan(&itemStatus.ID, &itemStatus.Value); err != nil {
		return nil, err
	}

	return &domains.ItemStatus{
		ID:    itemStatus.ID,
		Value: itemStatus.Value,
	}, nil
}
