package repositories

import (
	"github.com/M-kos/wb_level0/internal/domains"
	"time"
)

type addressModel struct {
	ID      int    `db:"id"`
	Zip     string `db:"zip"`
	Address string `db:"address"`
	CityID  int    `db:"city_id"`
}

type bankModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type brandModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type cityModel struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	RegionID int    `db:"region_id"`
}

type currencyModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type deliveryModel struct {
	ID         int `db:"id"`
	CustomerID int `db:"name"`
	AddressID  int `db:"address_id"`
}

type deliveryServiceModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type itemStatusModel struct {
	ID    int `db:"id"`
	Value int `db:"value"`
}

type localeModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type providerModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type regionModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type customerModel struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Phone     string `db:"phone"`
	Email     string `db:"email"`
}

type itemModel struct {
	ID          int    `db:"id"`
	ChrtID      int    `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       int    `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int    `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int    `db:"total_price"`
	NmID        int    `db:"nm_id"`
	BrandID     int    `db:"brand_id"`
	StatusID    int    `db:"status_id"`
}

func (i *itemModel) toDomain(brand brandModel, status itemStatusModel) *domains.Item {
	return &domains.Item{
		ID:          i.ID,
		ChrtID:      i.ChrtID,
		TrackNumber: i.TrackNumber,
		Price:       i.Price,
		Rid:         i.Rid,
		Name:        i.Name,
		Sale:        i.Sale,
		Size:        i.Size,
		TotalPrice:  i.TotalPrice,
		NmID:        i.NmID,
		Brand: domains.Brand{
			Name: brand.Name,
		},
		Status: domains.ItemStatus{
			Value: status.Value,
		},
	}
}

type paymentModel struct {
	ID           int    `db:"id"`
	Transaction  string `db:"transaction"`
	RequestID    string `db:"request_id"`
	CurrencyID   int    `db:"currency_id"`
	ProviderID   int    `db:"provider_id"`
	Amount       int    `db:"amount"`
	PaymentDt    int64  `db:"payment_dt"`
	BankID       int    `db:"bank_id"`
	DeliveryCost int    `db:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total"`
	CustomFee    int    `db:"custom_fee"`
}

func (p *paymentModel) toDomain(cur currencyModel, bank bankModel, provider providerModel) *domains.Payment {
	return &domains.Payment{
		ID:          p.ID,
		Transaction: p.Transaction,
		RequestID:   p.RequestID,
		Currency: domains.Currency{
			Name: cur.Name,
		},
		Provider: domains.Provider{
			Name: provider.Name,
		},
		Amount:    p.Amount,
		PaymentDt: p.PaymentDt,
		Bank: domains.Bank{
			Name: bank.Name,
		},
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}

type orderModel struct {
	ID                int       `db:"id"`
	OrderUID          string    `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	DeliveryID        int       `db:"delivery_id"`
	PaymentID         int       `db:"payment_id"`
	LocaleID          int       `db:"locale_id"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryServiceID int       `db:"delivery_service_id"`
	Shardkey          string    `db:"shardkey"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}

func (order *orderModel) toDomain(
	delivery deliveryModel,
	customer customerModel,
	address addressModel,
	city cityModel,
	region regionModel,
	payment paymentModel,
	currency currencyModel,
	provider providerModel,
	bank bankModel,
	locale localeModel,
	deliveryService deliveryServiceModel,
) *domains.Order {
	return &domains.Order{
		ID:          order.ID,
		OrderUID:    order.OrderUID,
		TrackNumber: order.TrackNumber,
		Entry:       order.Entry,
		Delivery: domains.Delivery{
			ID: delivery.ID,
			Customer: domains.Customer{
				ID:        customer.ID,
				FirstName: customer.FirstName,
				LastName:  customer.LastName,
				Phone:     customer.Phone,
				Email:     customer.Email,
			},
			Address: domains.Address{
				ID:      address.ID,
				Zip:     address.Zip,
				Address: address.Address,
				City: domains.City{
					ID:   city.ID,
					Name: city.Name,
				},
				Region: domains.Region{
					ID:   region.ID,
					Name: region.Name,
				},
			},
		},
		Payment: payment.toDomain(currency, bank, provider),
		Items:   nil,
		Locale: domains.Locale{
			ID:   locale.ID,
			Name: locale.Name,
		},
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService: domains.DeliveryService{
			ID:   deliveryService.ID,
			Name: deliveryService.Name,
		},
		Shardkey:    order.Shardkey,
		SmID:        order.SmID,
		DateCreated: order.DateCreated,
		OofShard:    order.OofShard,
	}
}
