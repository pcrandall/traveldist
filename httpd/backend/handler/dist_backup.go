package handler

import (
	"database/sql"
	"time"

	"github.com/pcrandall/travelDist/utils"
)

func Backup(location string) {
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
