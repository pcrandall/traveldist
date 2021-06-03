package handler

type ShuttleDistance struct {
	Shuttle   string
	Distance  string
	Timestamp string
}

type TravelDistances struct {
	T1_shuttle           string
	T1_distance          int
	T2_distance          int
	Shoe_travel_distance int
	T1_timestamp         string
	T2_timestamp         string
	Days_installed       string
	Notes                string
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

type ChangeShoe struct {
	Shuttle            string `json:"Shuttle"`
	NewChangeDistance  string `json:"New_Change_Distance"`
	NewChangeDate      string `json:"New_Change_Date"`
	NewChangeNotes     string `json:"New_Change_Notes"`
	PreviousChangeUUID string `json:"Previous_Change_UUID"`
}
