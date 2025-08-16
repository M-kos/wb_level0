package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type itemRepository struct {
	db *db.PostgresDB
}

func NewItemRepository(db *db.PostgresDB) (*itemRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &itemRepository{db: db}, nil
}

func (i *itemRepository) Create(ctx context.Context, item *domains.Item) (int, error) {
	row := i.db.Pool.QueryRow(ctx, findBrandByName, item.Brand.Name)
	var brand brandModel
	if err := row.Scan(&brand.ID, &brand.Name); err != nil {
		return 0, err
	}

	row = i.db.Pool.QueryRow(ctx, findItemStatusByValue, item.Status.Value)
	var status itemStatusModel
	if err := row.Scan(&status.ID, &status.Value); err != nil {
		return 0, err
	}

	row = i.db.Pool.QueryRow(ctx, createItem, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, brand.ID, status.ID)
	var itemId int
	if err := row.Scan(&itemId); err != nil {
		return 0, err
	}

	return itemId, nil
}

func (i *itemRepository) FindById(ctx context.Context, id int) (*domains.Item, error) {
	row := i.db.Pool.QueryRow(ctx, findItemById, id)
	var item itemModel

	if err := row.Scan(&item.ID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.BrandID, &item.StatusID); err != nil {
		return nil, err
	}

	row = i.db.Pool.QueryRow(ctx, findBrandById, item.BrandID)
	var brand brandModel
	if err := row.Scan(&brand.ID, &brand.Name); err != nil {
		return nil, err
	}

	row = i.db.Pool.QueryRow(ctx, findItemStatusById, item.StatusID)
	var status itemStatusModel
	if err := row.Scan(&status.ID, &status.Value); err != nil {
		return nil, err
	}

	return item.toDomain(brand, status), nil
}
func (i *itemRepository) ListByIds(ctx context.Context, ids []int) ([]*domains.Item, error) {
	rows, err := i.db.Pool.Query(ctx, listItemByIds, ids)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []itemModel

	for rows.Next() {
		var item itemModel
		if err := rows.Scan(&item.ID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.BrandID, &item.StatusID); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	var result []*domains.Item

	for _, item := range items {
		row := i.db.Pool.QueryRow(ctx, findBrandById, item.BrandID)
		var brand brandModel
		if err := row.Scan(&brand.ID, &brand.Name); err != nil {
			return nil, err
		}

		row = i.db.Pool.QueryRow(ctx, findItemStatusById, item.StatusID)
		var status itemStatusModel
		if err := row.Scan(&status.ID, &status.Value); err != nil {
			return nil, err
		}

		result = append(result, item.toDomain(brand, status))
	}

	return result, nil
}
