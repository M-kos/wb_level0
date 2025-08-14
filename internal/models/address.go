package models

type Address struct {
	ID      int
	Zip     string
	Address string
	City    string
	Region  string
}
type AddressDbModel struct {
	ID      int
	Zip     string
	Address string
	CityID  int
}
