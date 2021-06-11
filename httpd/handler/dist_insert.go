package handler

import (
	"database/sql"
	"strconv"

	"github.com/pcrandall/travelDist/utils"
)

func InsertDistance(distances []ShuttleDistance) {
	// Backup("db/traveldistances.db")
	// utils.CheckErr(err, "database.go.19 Error Backing up DB: ")

	db, err := sql.Open("sqlite3", "db/traveldistances.db")
	utils.CheckErr(err, "Error connecting to database")

	db.Ping()
	utils.CheckErr(err, "Error pinging database")

	defer db.Close()

	// insert distances into db
	for _, row := range distances {
		stmt, err := db.Prepare("INSERT INTO DISTANCES(shuttle, distance, timestamp) VALUES(?,?,?);")
		utils.CheckErr(err, "Error preparing DB")
		dist, err := strconv.Atoi(utils.TrimString(row.Distance))
		utils.CheckErr(err, "Error converting distance to int")

		// make the timestamp valid
		// 2020-09-02 15:16:15:415 --> 2020-09-02 15:16:15
		t := utils.CleanTimeStamp(row.Timestamp)
		res, err := stmt.Exec(&row.Shuttle, &dist, &t)
		utils.CheckErr(err, "Error inserting into db")

		_, err = res.LastInsertId()
		// id, err := res.LastInsertId()
		// log.Println("Last InsertID: ", id)
		utils.CheckErr(err, "database.go.37: Error getting last id")
	}
}
