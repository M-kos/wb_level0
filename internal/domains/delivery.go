package domains

type Delivery struct {
	ID       int
	Customer Customer
	Address  Address
}
