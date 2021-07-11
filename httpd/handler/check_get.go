package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	check "github.com/pcrandall/travelDist/httpd/platform/check_shoes"
	"github.com/pcrandall/travelDist/utils"
)

func GetCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		db, err := sql.Open("sqlite3", "db/traveldistances.db")
		utils.CheckErr(err, "check_get.go.17: Error connecting to database\t")
		defer db.Close()

		rows, err := db.Query(`SELECT * FROM view_check;`)
		utils.CheckErr(err, "check_get.go.21: Database Query error\t")

		keys := make(map[string]check.Check)
		for rows.Next() {
			var check check.Check
			rows.Scan(&check.Shuttle, &check.Distance, &check.Distance_1500km, &check.Timestamp, &check.Notes, &check.UUID, &check.Wear)
			utils.CheckErr(err, "")
			keys[check.Shuttle] = check
			// log.Printf("database.go.72 getDists(): %#v\n", dist)
		}
		// fmt.Printf("KEYS: %#v\n\n", keys)
		json.NewEncoder(w).Encode(&keys) // encode to json and send to client
		utils.CheckErr(err, "JSON encoding error")
	}
}
