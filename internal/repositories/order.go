package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
	"time"
)

type OrderRepository struct {
	db *db.PostgresDB
}

func NewOrderRepository(db *db.PostgresDB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) Create(ctx context.Context, order *domains.Order) (int, error) {
	tx, err := or.db.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}

		err = tx.Commit(ctx)
	}()

	var regionId int
	row := tx.QueryRow(ctx, createRegion, order.Delivery.Address.Region.Name)
	if err = row.Scan(&regionId); err != nil {
		fmt.Println("[createRegion]:", err.Error())
		return 0, err
	}

	var cityId int
	row = tx.QueryRow(ctx, createCity, order.Delivery.Address.City.Name, regionId)
	if err = row.Scan(&cityId); err != nil {
		fmt.Println("[createCity]:", err.Error())
		return 0, err
	}

	var addressId int
	row = tx.QueryRow(ctx, createAddress, order.Delivery.Address.Zip, order.Delivery.Address.Address, cityId)
	if err = row.Scan(&addressId); err != nil {
		fmt.Println("[createAddress]:", err.Error())
		return 0, err
	}

	var customerId int
	row = tx.QueryRow(
		ctx,
		createCustomer,
		order.Delivery.Customer.FirstName,
		order.Delivery.Customer.LastName,
		order.Delivery.Customer.Phone,
		order.Delivery.Customer.Email,
	)
	if err = row.Scan(&customerId); err != nil {
		fmt.Println("[createCustomer]:", err.Error())
		return 0, err
	}

	var deliveryId int
	row = tx.QueryRow(ctx, createDelivery, customerId, addressId)
	if err = row.Scan(&deliveryId); err != nil {
		fmt.Println("[createDelivery]:", err.Error())
		return 0, err
	}

	var currencyId int
	row = tx.QueryRow(ctx, createCurrency, order.Payment.Currency.Name)
	if err = row.Scan(&currencyId); err != nil {
		fmt.Println("[createCurrency]:", err.Error())
		return 0, err
	}

	var providerId int
	row = tx.QueryRow(ctx, createProvider, order.Payment.Provider.Name)
	if err = row.Scan(&providerId); err != nil {
		fmt.Println("[createProvider]:", err.Error())
		return 0, err
	}

	var bankId int
	row = tx.QueryRow(ctx, createBank, order.Payment.Bank.Name)
	if err = row.Scan(&bankId); err != nil {
		fmt.Println("[createBank]:", err.Error())
		return 0, err
	}

	var paymentId int
	paymentTime := time.Unix(order.Payment.PaymentDt, 0).UTC()
	row = tx.QueryRow(
		ctx,
		createPayment,
		order.Payment.Transaction,
		order.Payment.RequestID,
		currencyId,
		providerId,
		order.Payment.Amount,
		paymentTime,
		bankId,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	)
	if err = row.Scan(&paymentId); err != nil {
		fmt.Println("[createPayment]:", err.Error())
		return 0, err
	}

	var localeId int
	row = tx.QueryRow(ctx, createLocale, order.Locale.Name)
	if err = row.Scan(&localeId); err != nil {
		fmt.Println("[createLocale]:", err.Error())
		return 0, err
	}

	var deliveryServiceId int
	row = tx.QueryRow(ctx, createDeliveryService, order.DeliveryService.Name)
	if err = row.Scan(&deliveryServiceId); err != nil {
		fmt.Println("[deliveryServiceId]:", err.Error())
		return 0, err
	}

	var orderId int
	row = tx.QueryRow(ctx, createOrder,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		deliveryId,
		paymentId,
		localeId,
		order.InternalSignature,
		order.CustomerID,
		deliveryServiceId,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	if err := row.Scan(&orderId); err != nil {
		fmt.Println("[createOrder]:", err.Error())
		return 0, err
	}

	fmt.Println("prder.Items", order.Items)

	for i, item := range order.Items {
		fmt.Println("Item", item)
		var brandId int
		row = tx.QueryRow(ctx, createBrand, item.Brand.Name)
		if err = row.Scan(&brandId); err != nil {
			fmt.Println("[createBrand]:", i, err.Error())
			return 0, err
		}

		var statusId int
		row = tx.QueryRow(ctx, createItemStatus, item.Status.Value)
		if err = row.Scan(&statusId); err != nil {
			fmt.Println("[createItemStatus]:", i, err.Error())
			return 0, err
		}

		var itemId int
		row = tx.QueryRow(ctx, createItem,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			brandId,
			statusId,
		)
		if err := row.Scan(&itemId); err != nil {
			fmt.Println("[createItem]:", i, err.Error())
			return 0, err
		}

		_, err = tx.Exec(ctx, createOrderItem, orderId, itemId)
		if err != nil {
			fmt.Println("[createOrderItem]:", i, err.Error())
			return 0, err
		}
	}

	return orderId, nil
}

func (or *OrderRepository) GetById(ctx context.Context, orderId string) (*domains.Order, error) {
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

	var paymentDt time.Time

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
		&paymentDt,
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

	payment.PaymentDt = paymentDt.Unix()

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

func (or *OrderRepository) List(ctx context.Context, limit int) ([]*domains.Order, error) {
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
	var paymentDt time.Time

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
			&paymentDt,
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

		payment.PaymentDt = paymentDt.Unix()

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
