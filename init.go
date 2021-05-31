package main

import (
	"os"
	"path/filepath"

	"github.com/pcrandall/travelDist/utils"
)

func init() {
	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}

	dbpath := filepath.Join(".", "db")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		os.MkdirAll(dbpath, os.ModePerm)
	}
	utils.GetConfig("./config/config.yml", &config)
}
