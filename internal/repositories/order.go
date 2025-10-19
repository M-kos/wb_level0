package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type OrderRepository struct {
	db *db.PostgresDB
}

func NewOrderRepository(db *db.PostgresDB) (*OrderRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &OrderRepository{db: db}, nil
}

func (or *OrderRepository) Create(ctx context.Context, orderRepository *domains.Order) (int, error) {
	row := or.db.Pool.QueryRow(ctx,
		orderRepository.OrderUID,
		orderRepository.TrackNumber,
		orderRepository.Entry,
		orderRepository.Delivery.ID,
		orderRepository.Payment.ID,
		orderRepository.Locale.ID,
		orderRepository.InternalSignature,
		orderRepository.Delivery.Customer.ID,
		orderRepository.DeliveryService.ID,
		orderRepository.Shardkey,
		orderRepository.SmID,
		orderRepository.DateCreated,
		orderRepository.OofShard,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (or *OrderRepository) GetById(ctx context.Context, orderId string) (*domains.Order, error) {
	var orderRepository orderModel
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
		&orderRepository.ID,
		&orderRepository.OrderUID,
		&orderRepository.TrackNumber,
		&orderRepository.Entry,
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
		&orderRepository.InternalSignature,
		&orderRepository.CustomerID,
		&deliveryService.ID,
		&deliveryService.Name,
		&orderRepository.Shardkey,
		&orderRepository.SmID,
		&orderRepository.DateCreated,
		&orderRepository.OofShard,
	); err != nil {
		return nil, err
	}

	rows, err := or.db.Pool.Query(ctx, getOrderItemsByOrderId, orderRepository.ID)
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

	dOrder := orderRepository.toDomain(
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

func (or *OrderRepository) List(ctx context.Context, limit int) ([]*domains.Order, error) {
	var orderRepository orderModel
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
			&orderRepository.ID,
			&orderRepository.OrderUID,
			&orderRepository.TrackNumber,
			&orderRepository.Entry,
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
			&orderRepository.InternalSignature,
			&orderRepository.CustomerID,
			&deliveryService.ID,
			&deliveryService.Name,
			&orderRepository.Shardkey,
			&orderRepository.SmID,
			&orderRepository.DateCreated,
			&orderRepository.OofShard,
		); err != nil {
			return nil, err
		}

		itemIdsRows, err := or.db.Pool.Query(ctx, getOrderItemsByOrderId, orderRepository.ID)
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

		dOrder := orderRepository.toDomain(
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
