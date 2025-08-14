package models

type Delivery struct {
	Name    string
	Phone   string
	Address Address
	Email   string
}
type DeliveryDbModel struct {
	Name      string
	Phone     string
	AddressID int
	Email     string
}
