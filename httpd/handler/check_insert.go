package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	check "github.com/pcrandall/travelDist/httpd/platform/check_shoes"
	"github.com/pcrandall/travelDist/utils"
)

type InsertCheckSuccess struct {
	Nice string `json:"Success"`
}

type InsertCheckFailure struct {
	Err error `json:"Error"`
}

func InsertCheck(c check.Adder) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var check check.Check
		utils.HttpPrettyPrintRequest(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(&check) // encode to json and send to client

		fmt.Printf("check value: %#v", check)

		db, err := sql.Open("sqlite3", "db/traveldistances.db")
		utils.CheckErr(err, "Error connecting to database")

		db.Ping()
		utils.CheckErr(err, "Error pinging database")

		defer db.Close()

		// insert shoecheck into table
		stmt, err := db.Prepare("INSERT INTO SHOE_CHECK(shuttle, distance, timestamp, notes, uuid, wear) VALUES(?,?,?,?,?,?);")
		utils.CheckErr(err, "Error preparing DB")

		check.UUID, err = utils.GenerateUUID()
		utils.CheckErr(err, "Error generating uuid")
		// fmt.Printf("check: %#v", check)

		res, err := stmt.Exec(&check.Shuttle, &check.Distance, &check.Timestamp, &check.Notes, &check.UUID, &check.Wear)

		utils.CheckErr(err, "Error inserting into check table")

		if err != nil {
			greatJob := &Success{"Great Success!"}
			fmt.Println(err)
			json.NewEncoder(w).Encode(&greatJob) // encode to json and send to client
		} else {
			badJob := &Failure{err}
			json.NewEncoder(w).Encode(&badJob) // encode to json and send to client
		}

		id, err := res.LastInsertId()
		fmt.Println("Last InsertID: ", id)
		utils.CheckErr(err, "Error getting last id")
	}
}
