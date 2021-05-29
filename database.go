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
	// q := `SELECT
	// 	  DISTANCES.shuttle as t1_shuttle,
	// 	  DISTANCES.distance as t1_distance,
	// 	  max(CHANGE.distance) as t2_distance,
	// 	  (DISTANCES.distance-CHANGE.distance)/1000 as shoe_travel_difference,
	// 	  max(DISTANCES.timestamp) as t1_timestamp,
	// 	  CHANGE.timestamp as t2_timestamp,
	// 	  (SELECT
	// 	  printf("%03d", julianday(DISTANCES.timestamp)-julianday(CHANGE.timestamp))) as days_installed,
	// 	  CHANGE.notes as notes
	// 	FROM
	// 	  DISTANCES
	// 	INNER JOIN CHANGE on CHANGE.shuttle = DISTANCES.shuttle
	// 	GROUP BY t1_shuttle
	// 	ORDER BY shoe_travel_difference DESC`

	q := `SELECT * FROM shoe_travel;`

	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	checkErr(err, "Error connecting to database")
	db.Ping()
	checkErr(err, "Error pinging database")
	defer db.Close()

	rows, err := db.Query(q)
	checkErr(err, "Database Query error:  ")
	// var t1_distance, t2_distance, shoe_travel_difference int
	// var notes, days_installed, t1_shuttle, t1_timestamp, t2_timestamp string

	for rows.Next() {
		var dist travelDistances
		// rows.Scan(&t1_shuttle, &t1_distance, &t2_distance, &shoe_travel_difference, &t1_timestamp, &t2_timestamp, &days_installed, &notes)
		rows.Scan(&dist.t1_shuttle, &dist.t1_distance, &dist.t2_distance, &dist.shoe_travel_distance, &dist.t1_timestamp, &dist.t2_timestamp, &dist.days_installed, &dist.notes)
		checkErr(err, "")
		log.Println("t1_shuttle: ", dist.t1_shuttle, "t1_distance: ", dist.t1_distance, "t2_distance: ", dist.t2_distance, "shoe_travel_difference: ", dist.shoe_travel_distance, "t1_timestamp: ", dist.t1_timestamp, "t2_timestamp: ", dist.t2_timestamp, "days_installed: ", dist.days_installed, "notes: ", dist.notes)
		fmt.Println("t1_shuttle: ", dist.t1_shuttle, "t1_distance: ", dist.t1_distance, "t2_distance: ", dist.t2_distance, "shoe_travel_difference: ", dist.shoe_travel_distance, "t1_timestamp: ", dist.t1_timestamp, "t2_timestamp: ", dist.t2_timestamp, "days_installed: ", dist.days_installed, "notes: ", dist.notes)
		getDistances = append(getDistances, dist)
		// fmt.Println("ID: ", id, "Shuttle: ", shuttle, "Distance: ", distance, "Timestamp: ", timestamp)

		// 			row := new(travelDistances)
		// 			row.shuttle = cleanString(n.Name)
		// 			row.timestamp = cleanString(lastDate)
		// 			row.distance = total[7:]
		// 			allDistances = append(allDistances, *row)

	}

	json.NewEncoder(w).Encode(getDistances)
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
