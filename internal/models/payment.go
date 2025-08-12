package models

type Payment struct {
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDt    int64
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}
