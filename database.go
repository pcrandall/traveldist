package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pcrandall/travelDist/utils"
	"github.com/rs/cors"
)

func insertDatabase(distances []shuttleDistance) {
	BackupDB("db/traveldistances.db")
	utils.CheckErr(err, "database.go.19 Error Backing up DB: ")

	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	utils.CheckErr(err, "Error connecting to database")

	db.Ping()
	utils.CheckErr(err, "Error pinging database")

	defer db.Close()

	// insert distances into db
	for _, row := range distances {
		stmt, err := db.Prepare("INSERT INTO DISTANCES(shuttle, distance, timestamp) VALUES(?,?,?);")
		utils.CheckErr(err, "Error preparing DB")
		dist, err := strconv.Atoi(utils.TrimString(row.distance))
		utils.CheckErr(err, "Error converting distance to int")

		// make the timestamp valid
		// 2020-09-02 15:16:15:415 --> 2020-09-02 15:16:15
		t := utils.CleanTimeStamp(row.timestamp)
		res, err := stmt.Exec(&row.shuttle, &dist, &t)
		utils.CheckErr(err, "Error inserting into db")

		_, err = res.LastInsertId()
		// id, err := res.LastInsertId()
		// log.Println("Last InsertID: ", id)
		utils.CheckErr(err, "database.go.37: Error getting last id")
	}
}

func DBRouter() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},          // All origins
		AllowedMethods: []string{"GET", "PUT"}, // Allowing only get, just an example
	})
	router := mux.NewRouter()
	router.HandleFunc("/dists", getDists).Methods("GET") // get all the distances from db
	// router.HandleFunc("/dist", createDist).Methods("POST")
	// router.HandleFunc("/dist/{id}", getDist).Methods("GET")
	// router.HandleFunc("/dist/{id}", updateDist).Methods("PUT")
	// router.HandleFunc("/dist/{id}", deleteDist).Methods("DELETE")
	fmt.Println("Backend Server is ready and is listening at port :8001...")
	log.Fatal(http.ListenAndServe(":8001", c.Handler(router)))
}

func getDists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	utils.CheckErr(err, "database.go.60: Error connecting to database\t")
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM clean_shoe_travel;`)
	utils.CheckErr(err, "database.go.64: Database Query error\t")

	keys := make(map[string]cleanTravelDistances)
	for rows.Next() {
		var dist cleanTravelDistances
		rows.Scan(&dist.Shuttle, &dist.Last_Updated, &dist.Shoe_Travel, &dist.Days_Installed, &dist.Shoes_Last_Distance, &dist.Shoes_Change_Distance, &dist.Shoes_Last_Changed, &dist.Notes, &dist.UUID)
		utils.CheckErr(err, "")
		keys[dist.Shuttle] = dist
		// log.Printf("database.go.72 getDists(): %#v\n", dist)
	}

	// log.Printf("KEYS: %#v\n\n", keys)
	json.NewEncoder(w).Encode(&keys) // encode to json and send to client
	utils.CheckErr(err, "JSON encoding error")
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

func BackupDB(location string) {
	db, err := sql.Open("sqlite3", location)
	utils.CheckErr(err, "Error connecting to database")
	db.Ping()
	utils.CheckErr(err, "Error pinging database"+location)
	defer db.Close()

	t := time.Now()
	backupLocation := "./db/backups/" + t.Format("2006-01-02_150405") + "__traveldistances.db"
	stmt, err := db.Prepare(`VACUUM main INTO ?;`)
	utils.CheckErr(err, "database.go.112 Error preparing DB: ")

	_, err = stmt.Exec(&backupLocation)
	utils.CheckErr(err, "database.go.115 Error backing up DB: ")
}
