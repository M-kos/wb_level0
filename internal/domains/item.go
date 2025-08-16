package domains

type Item struct {
	ID          int
	ChrtID      int
	TrackNumber string
	Price       int
	Rid         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  int
	NmID        int
	Brand       Brand
	Status      ItemStatus
}
