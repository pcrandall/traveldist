package shoeparameters

var (
	shoeParameter *ShoeParameter
)

// ShoeParameter
type ShoeParameter struct {
	Check    int
	Interval int
}

func SetShoeParameters(check int, interval int) {
	shoeParameter = &ShoeParameter{check, interval}
}

func GetShoeParameters() *ShoeParameter {
	return shoeParameter
}
