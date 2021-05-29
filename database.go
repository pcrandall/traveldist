package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func insertDatabase(val []shuttleDistance) {
	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	defer db.Close()
	checkErr(err, "Error connecting to database")

	db.Ping()
	checkErr(err, "Error pinging database")

	// insert
	for _, row := range val {

		fmt.Printf("ROW VALUES: %#v\tROW TYPES: %T", row, row)

		stmt, err := db.Prepare("INSERT INTO DISTANCES(shuttle, distance, timestamp) VALUES(?,?,?);")
		checkErr(err, "Error preparing DB")
		dist, err := strconv.Atoi(cleanString(row.distance))
		checkErr(err, "Error converting distance to int")

		res, err := stmt.Exec(&row.shuttle, &dist, &row.timestamp)
		checkErr(err, "Error inserting into db")

		id, err := res.LastInsertId()
		checkErr(err, "Error getting last id")

		fmt.Println("Last InsertID: ", id)
	}

	// statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS locations (id INTEGER PRIMARY KEY AUTOINCREMENT, location TEXT NOT NULL, description TEXT NOT NULL, area TEXT)")
	// statement.Exec()
	// query = cleanString(fmt.Sprintf("%%%s%%", query))
	rows, err := db.Query("SELECT * FROM DISTANCES")
	checkErr(err, "Database Query error:  ")
	var id, distance int
	var shuttle, timestamp string

	for rows.Next() {
		rows.Scan(&id, &shuttle, &distance, &timestamp)
		checkErr(err, "")
		fmt.Println("ID: ", id, "Shuttle: ", shuttle, "Distance: ", distance, "Timestamp: ", timestamp)

	}
}

func router() {

	router := mux.NewRouter()

	router.HandleFunc("/dist", getDists).Methods("GET")
	// router.HandleFunc("/dist", createDist).Methods("POST")
	// router.HandleFunc("/dist/{id}", getDist).Methods("GET")
	// router.HandleFunc("/dist/{id}", updateDist).Methods("PUT")
	// router.HandleFunc("/dist/{id}", deleteDist).Methods("DELETE")

	http.ListenAndServe(":8001", router)
}

func getDists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	defer db.Close()
	checkErr(err, "Error connecting to database")

	rows, err := db.Query(`SELECT * FROM shoe_travel;`)
	checkErr(err, "Database Query error:  ")

	for rows.Next() {
		var dist travelDistances
		rows.Scan(dist.t1_shuttle, dist.t1_distance, dist.t2_distance, dist.shoe_travel_distance, dist.t1_timestamp, dist.t2_timestamp, dist.days_installed, dist.notes)
		checkErr(err, "")
		log.Println("t1_shuttle: ", dist.t1_shuttle, "t1_distance: ", dist.t1_distance, "t2_distance: ", dist.t2_distance, "shoe_travel_difference: ", dist.shoe_travel_distance, "t1_timestamp: ", dist.t1_timestamp, "t2_timestamp: ", dist.t2_timestamp, "days_installed: ", dist.days_installed, "notes: ", dist.notes)
		fmt.Println("t1_shuttle: ", dist.t1_shuttle, "t1_distance: ", dist.t1_distance, "t2_distance: ", dist.t2_distance, "shoe_travel_difference: ", dist.shoe_travel_distance, "t1_timestamp: ", dist.t1_timestamp, "t2_timestamp: ", dist.t2_timestamp, "days_installed: ", dist.days_installed, "notes: ", dist.notes)
		getTravelDistances = append(getTravelDistances, dist)
	}

	json.NewEncoder(w).Encode(getTravelDistances)
	checkErr(err, "JSON encoding error")

}

// func createPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	title := keyVal["title"]
// 	_, err = stmt.Exec(title)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "New post was created")
// }
