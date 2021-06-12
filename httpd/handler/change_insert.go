package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	change "github.com/pcrandall/travelDist/httpd/platform/change_shoes"
	"github.com/pcrandall/travelDist/utils"
)

type Success struct {
	Nice string `json:"Success"`
}

type Failure struct {
	Err error `json:"Error"`
}

func InsertChange(c change.Adder) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var change change.Change
		// var change models.Change
		utils.HttpPrettyPrintRequest(r)
		w.Header().Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(&change) // encode to json and send to client

		db, err := sql.Open("sqlite3", "db/traveldistances.db")
		utils.CheckErr(err, "Error connecting to database")

		db.Ping()
		utils.CheckErr(err, "Error pinging database")

		defer db.Close()

		// insert change into table
		stmt, err := db.Prepare("INSERT INTO CHANGE(shuttle, distance, timestamp, notes, uuid) VALUES(?,?,?,?,?);")
		utils.CheckErr(err, "Error preparing DB")
		// dist, err := strconv.Atoi(utils.TrimString(change.Distance))
		// utils.CheckErr(err, "Error converting distance to int")

		change.UUID, err = utils.GenerateUUID()
		utils.CheckErr(err, "Error generating uuid")
		fmt.Printf("change: %#v", change)
		// make the timestamp valid
		// 2020-09-02 15:16:15:415 --> 2020-09-02 15:16:15
		res, err := stmt.Exec(&change.Shuttle, &change.Distance, &change.Timestamp, &change.Notes, &change.UUID)

		utils.CheckErr(err, "Error inserting into change table")

		if err != nil {
			greatJob := &Success{"Great Success!"}
			fmt.Println(err)
			err = json.NewEncoder(w).Encode(&greatJob) // encode to json and send to client
		} else {
			badJob := &Failure{err}
			err = json.NewEncoder(w).Encode(&badJob) // encode to json and send to client
		}

		id, err := res.LastInsertId()
		fmt.Println("Last InsertID: ", id)
		utils.CheckErr(err, "Error getting last id")

		// utils.CheckErr(err, "JSON encoding error")
		// err = json.NewEncoder(w).Encode(&change) // encode to json and send to client
		// utils.CheckErr(err, "JSON encoding error")
		// if err != nil {
		// 	e := &Failure{err}
		// 	json.NewEncoder(w).Encode(e) // encode to json and send to client
		// }

	}
}
