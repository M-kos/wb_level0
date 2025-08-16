package domains

type Payment struct {
	ID           int
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
