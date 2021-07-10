package check

type Getter interface {
	GetAll() []Check
}

type Adder interface {
	Add(item Check)
}

type Check struct {
	Shuttle         string  `json:"Shuttle"`
	Distance        int     `json:"Distance"`
	Distance_1500km int     `json:"Distance_1500km"`
	Timestamp       string  `json:"Timestamp"`
	Notes           string  `json:"Notes"`
	UUID            string  `json:"UUID"`
	Wear            float64 `json:"Wear"`
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
