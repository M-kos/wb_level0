package models

type Payment struct {
	Transaction  string
	RequestID    string
	Currency     Currency
	Provider     Provider
	Amount       int
	PaymentDt    int64
	Bank         Bank
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

type PaymentDBModel struct {
	Transaction  string
	RequestID    string
	CurrencyID   int
	ProviderID   int
	Amount       int
	PaymentDt    int64
	BankID       int
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}
