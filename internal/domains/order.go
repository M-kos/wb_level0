package domains

type Order struct {
	ID                int
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          Delivery
	Payment           *Payment
	Items             []*Item
	Locale            Locale
	InternalSignature string
	CustomerID        string
	DeliveryService   DeliveryService
	Shardkey          string
	SmID              int
	DateCreated       string
	OofShard          string
}
