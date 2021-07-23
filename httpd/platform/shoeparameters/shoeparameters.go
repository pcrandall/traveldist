package shoeparameters

var (
	shoeParameter *ShoeParameter
)

// ShoeParameter
type ShoeParameter struct {
	Check    int
	Interval int
	Min_Shoe float64
	Max_Shoe float64
}

func SetShoeParameters(check int, interval int, min float64, max float64) {
	shoeParameter = &ShoeParameter{check, interval, min, max}
}

func GetShoeParameters() *ShoeParameter {
	return shoeParameter
}
