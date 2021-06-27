package check

type Getter interface {
	GetAll() []Check
}

type Adder interface {
	Add(item Check)
}

type Check struct {
	Shuttle   string `json:"Shuttle"`
	Distance  int    `json:"New_Check_Distance"`
	Timestamp string `json:"New_Check_Date"`
	Notes     string `json:"New_Check_Notes"`
	UUID      string `json:"Previous_Check_UUID"`
	Wear      string `json:"Wear"`
}

type ShoeCheck struct {
	Checks []Check
}

func New() *ShoeCheck {
	return &ShoeCheck{
		Checks: []Check{},
	}
}

func (r *ShoeCheck) Add(item Check) {
	r.Checks = append(r.Checks, item)
}

func (r *ShoeCheck) GetAll() []Check {
	return r.Checks
}
