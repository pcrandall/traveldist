package distances

type Getter interface {
	GetAll() []Distance
}

type Adder interface {
	Add(item Distance)
}

type Distance struct {
	Shuttle               string
	Last_Updated          string
	Shoes_Last_Distance   int
	Shoes_Change_Distance int
	Shoe_Travel           int
	Days_Installed        string
	Shoes_Last_Changed    string
	Notes                 string
	UUID                  string
}

type TravelDistances struct {
	Distances []Distance
}

func New() *TravelDistances {
	return &TravelDistances{
		Distances: []Distance{},
	}
}

func (r *TravelDistances) Add(item Distance) {
	r.Distances = append(r.Distances, item)
}

func (r *TravelDistances) GetAll() []Distance {
	return r.Distances
}
