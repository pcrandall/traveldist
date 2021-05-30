package main

type Config struct {
	SheetName string `yaml:"sheetname"`
	Levels    []struct {
		Floor   int `yaml:"floor"`
		Navette []struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
			Row  string `yaml:"row"`
		} `yaml:"navette"`
	} `yaml:"levels"`
}

type shuttleDistance struct {
	shuttle   string
	distance  string
	timestamp string
}

type travelDistances struct {
	T1_shuttle           string
	T1_distance          int
	T2_distance          int
	Shoe_travel_distance int
	T1_timestamp         string
	T2_timestamp         string
	Days_installed       string
	Notes                string
}

type cleanTravelDistances struct {
	Shuttle               string
	Last_Updated          string
	Shoes_Last_Distance   int
	Shoes_Change_Distance int
	Shoe_Travel           int
	Days_Installed        string
	Shoes_Last_Changed    string
	Notes                 string
}
