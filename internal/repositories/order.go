package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type orderRepository struct {
	db *db.PostgresDB
}

func NewOrderRepository(db *db.PostgresDB) (*orderRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &orderRepository{db: db}, nil
}

func (or *orderRepository) Create(ctx context.Context, order *domains.Order) (int, error) {
	row := or.db.Pool.QueryRow(ctx,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Delivery.ID,
		order.Payment.ID,
		order.Locale.ID,
		order.InternalSignature,
		order.Delivery.Customer.ID,
		order.DeliveryService.ID,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (or *orderRepository) GetById(ctx context.Context, orderId int) (*domains.Order, error) {
	var order orderModel
	var delivery deliveryModel
	var customer customerModel
	var address addressModel
	var city cityModel
	var region regionModel
	var payment paymentModel
	var currency currencyModel
	var provider providerModel
	var bank bankModel
	var locale localeModel
	var deliveryService deliveryServiceModel

	row := or.db.Pool.QueryRow(ctx, getOrderById, orderId)

	if err := row.Scan(
		&order.ID,
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&delivery.ID,
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&address.ID,
		&address.Zip,
		&city.ID,
		&city.Name,
		&address.Address,
		&region.ID,
		&region.Name,
		&customer.Email,
		&payment.ID,
		&payment.Transaction,
		&payment.RequestID,
		&currency.ID,
		&currency.Name,
		&provider.ID,
		&provider.Name,
		&payment.Amount,
		&payment.PaymentDt,
		&bank.ID,
		&bank.Name,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
		&locale.ID,
		&locale.Name,
		&order.InternalSignature,
		&order.CustomerID,
		&deliveryService.ID,
		&deliveryService.Name,
		&order.Shardkey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
	); err != nil {
		return nil, err
	}

	rows, err := or.db.Pool.Query(ctx, getOrderItemsByOrderId, order.ID)
	if err != nil {
		return nil, err
	}

	var itemIds []int

	for rows.Next() {
		var itemId int
		if err := rows.Scan(&itemId); err != nil {
			return nil, err
		}

		itemIds = append(itemIds, itemId)
	}

	rows, err = or.db.Pool.Query(ctx, listItemByIds, itemIds)
	if err != nil {
		return nil, err
	}

	var items []itemModel

	for rows.Next() {
		var item itemModel
		if err := rows.Scan(
			&item.ID,
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.BrandID,
			&item.StatusID,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	defer rows.Close()

	var dItems []*domains.Item

	for _, item := range items {
		row := or.db.Pool.QueryRow(ctx, findBrandById, item.BrandID)
		var brand brandModel
		if err := row.Scan(&brand.ID, &brand.Name); err != nil {
			return nil, err
		}

		row = or.db.Pool.QueryRow(ctx, findItemStatusById, item.StatusID)
		var status itemStatusModel
		if err := row.Scan(&status.ID, &status.Value); err != nil {
			return nil, err
		}

		dItems = append(dItems, item.toDomain(brand, status))
	}

	dOrder := order.toDomain(
		delivery,
		customer,
		address,
		city,
		region,
		payment,
		currency,
		provider,
		bank,
		locale,
		deliveryService,
	)

	dOrder.Items = dItems

	return dOrder, nil
}

func (or *orderRepository) List(ctx context.Context, limit int) ([]*domains.Order, error) {
	var order orderModel
	var delivery deliveryModel
	var customer customerModel
	var address addressModel
	var city cityModel
	var region regionModel
	var payment paymentModel
	var currency currencyModel
	var provider providerModel
	var bank bankModel
	var locale localeModel
	var deliveryService deliveryServiceModel

	rows, err := or.db.Pool.Query(ctx, getOrdersByLimit, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*domains.Order

	for rows.Next() {
		if err := rows.Scan(
			&order.ID,
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&delivery.ID,
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Phone,
			&address.ID,
			&address.Zip,
			&city.ID,
			&city.Name,
			&address.Address,
			&region.ID,
			&region.Name,
			&customer.Email,
			&payment.ID,
			&payment.Transaction,
			&payment.RequestID,
			&currency.ID,
			&currency.Name,
			&provider.ID,
			&provider.Name,
			&payment.Amount,
			&payment.PaymentDt,
			&bank.ID,
			&bank.Name,
			&payment.DeliveryCost,
			&payment.GoodsTotal,
			&payment.CustomFee,
			&locale.ID,
			&locale.Name,
			&order.InternalSignature,
			&order.CustomerID,
			&deliveryService.ID,
			&deliveryService.Name,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
		); err != nil {
			return nil, err
		}

		itemIdsRows, err := or.db.Pool.Query(ctx, getOrderItemsByOrderId, order.ID)
		if err != nil {
			return nil, err
		}

		var itemIds []int

		for itemIdsRows.Next() {
			var itemId int
			if err := itemIdsRows.Scan(&itemId); err != nil {
				return nil, err
			}

			itemIds = append(itemIds, itemId)
		}

		itemIdsRows.Close()

		itemRows, err := or.db.Pool.Query(ctx, listItemByIds, itemIds)
		if err != nil {
			return nil, err
		}

		var items []itemModel

		for itemRows.Next() {
			var item itemModel
			if err := itemRows.Scan(
				&item.ID,
				&item.ChrtID,
				&item.TrackNumber,
				&item.Price,
				&item.Rid,
				&item.Name,
				&item.Sale,
				&item.Size,
				&item.TotalPrice,
				&item.NmID,
				&item.BrandID,
				&item.StatusID,
			); err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		itemRows.Close()

		var dItems []*domains.Item

		for _, item := range items {
			row := or.db.Pool.QueryRow(ctx, findBrandById, item.BrandID)
			var brand brandModel
			if err := row.Scan(&brand.ID, &brand.Name); err != nil {
				return nil, err
			}

			row = or.db.Pool.QueryRow(ctx, findItemStatusById, item.StatusID)
			var status itemStatusModel
			if err := row.Scan(&status.ID, &status.Value); err != nil {
				return nil, err
			}

			dItems = append(dItems, item.toDomain(brand, status))
		}

		dOrder := order.toDomain(
			delivery,
			customer,
			address,
			city,
			region,
			payment,
			currency,
			provider,
			bank,
			locale,
			deliveryService,
		)

		dOrder.Items = dItems

		orders = append(orders, dOrder)
	}

	return orders, nil
}
