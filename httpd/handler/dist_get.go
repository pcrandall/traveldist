package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pcrandall/travelDist/httpd/platform/distances"
	"github.com/pcrandall/travelDist/utils"
)

func GetDistances(_dist distances.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// items := feed.GetAll()
		// json.NewEncoder(w).Encode(items)

		w.Header().Set("Content-Type", "application/json")
		db, err := sql.Open("sqlite3", "db/traveldistances.db")
		utils.CheckErr(err, "database.go.60: Error connecting to database\t")
		defer db.Close()

		rows, err := db.Query(`SELECT * FROM clean_shoe_travel;`)
		utils.CheckErr(err, "database.go.64: Database Query error\t")

		keys := make(map[string]CleanTravelDistances)
		for rows.Next() {
			var dist CleanTravelDistances
			rows.Scan(&dist.Shuttle, &dist.Last_Updated, &dist.Shoe_Travel, &dist.Days_Installed, &dist.Shoes_Last_Distance, &dist.Shoes_Change_Distance, &dist.Shoes_Last_Changed, &dist.Notes, &dist.UUID)
			utils.CheckErr(err, "")
			keys[dist.Shuttle] = dist
			// log.Printf("database.go.72 getDists(): %#v\n", dist)
		}

		// log.Printf("KEYS: %#v\n\n", keys)
		json.NewEncoder(w).Encode(&keys) // encode to json and send to client
		utils.CheckErr(err, "JSON encoding error")
	}
}
