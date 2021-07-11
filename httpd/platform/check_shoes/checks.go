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

// ViewCheck represents a row from 'view_check'.
type ViewCheck struct {
	Shuttle            string  `json:"Shuttle"`              // Shuttle
	ZeroDistance       int64   `json:"Zero_Distance"`        // Zero_Distance
	LastCheckDistance  int64   `json:"Last_Check_Distance"`  // Last_Check_Distance
	CheckTrigger       int64   `json:"Check_Trigger"`        // Check_Trigger
	CurrentDistance    int64   `json:"Current_Distance"`     // Current_Distance
	CheckShoes         bool    `json:"Check_Shoes"`          // Check_Shoes
	LastCheckTimestamp string  `json:"Last_Check_Timestamp"` // Last_Check_Timestamp
	LastCheckNotes     string  `json:"Last_Check_Notes"`     // Last_Check_Notes
	LastCheckUUID      string  `json:"Last_Check_UUID"`      // Last_Check_UUID
	LastCheckWear      float64 `json:"Last_Check_Wear"`      // Last_Check_Wear
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
