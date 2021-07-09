package handler

type ShuttleDistance struct {
	Shuttle   string
	Distance  string
	Timestamp string
}

type CleanTravelDistances struct {
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

// type CleanShoeCheck struct {
// 	Shuttle              string
// 	Shoes_Check_Distance int
// 	Last_Updated         string
// 	Notes                string
// 	UUID                 string
// 	Wear                 float64
// }
