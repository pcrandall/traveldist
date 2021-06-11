package distances

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
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

type Distance struct {
	Items []Item
}

func New() *Distance {
	return &Distance{
		Items: []Item{},
	}
}

func (r *Distance) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Distance) GetAll() []Item {
	return r.Items
}
