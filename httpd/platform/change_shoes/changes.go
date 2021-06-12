package change

type Getter interface {
	GetAll() []Change
}

type Adder interface {
	Add(item Change)
}

type Change struct {
	Shuttle   string `json:"Shuttle"`
	Distance  int    `json:"New_Change_Distance"`
	Timestamp string `json:"New_Change_Date"`
	Notes     string `json:"New_Change_Notes"`
	UUID      string `json:"Previous_Change_UUID"`
}

type ShoeChange struct {
	Changes []Change
}

func New() *ShoeChange {
	return &ShoeChange{
		Changes: []Change{},
	}
}

func (r *ShoeChange) Add(item Change) {
	r.Changes = append(r.Changes, item)
}

func (r *ShoeChange) GetAll() []Change {
	return r.Changes
}
