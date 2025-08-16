package domains

type Address struct {
	ID      int
	Zip     string
	Address string
	City    City
	Region  Region
}
