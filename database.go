package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pcrandall/travelDist/utils"
	"github.com/rs/cors"
)

func insertDatabase(val []shuttleDistance) {
	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	utils.CheckErr(err, "Error connecting to database")
	db.Ping()
	utils.CheckErr(err, "Error pinging database")
	defer db.Close()

	// insert
	for _, row := range val {
		// fmt.Printf("ROW VALUES: %#v\tROW TYPES: %T", row, row)
		stmt, err := db.Prepare("INSERT INTO DISTANCES(shuttle, distance, timestamp) VALUES(?,?,?);")
		utils.CheckErr(err, "Error preparing DB")
		dist, err := strconv.Atoi(utils.CleanString(row.distance))
		utils.CheckErr(err, "Error converting distance to int")

		res, err := stmt.Exec(&row.shuttle, &dist, &row.timestamp)
		utils.CheckErr(err, "Error inserting into db")

		id, err := res.LastInsertId()
		utils.CheckErr(err, "Error getting last id")

		log.Println("Last InsertID: ", id)
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
	utils.CheckErr(err, "Error connecting to database")
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM clean_shoe_travel;`)
	utils.CheckErr(err, "Database Query error:  ")

	keys := make(map[string]cleanTravelDistances)
	for rows.Next() {
		var dist cleanTravelDistances
		rows.Scan(&dist.Shuttle, &dist.Last_Updated, &dist.Shoe_Travel, &dist.Days_Installed, &dist.Shoes_Last_Distance, &dist.Shoes_Change_Distance, &dist.Shoes_Last_Changed, &dist.Notes, &dist.UUID)
		utils.CheckErr(err, "")
		keys[dist.Shuttle] = dist
	}

	log.Printf("KEYS: %#v\n\n", keys)
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
