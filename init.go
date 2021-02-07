package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {

	newpath := filepath.Join(".", "old")
	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		os.MkdirAll(newpath, os.ModePerm)
		// os.Mkdir("scannedIPs", 777)
	}

	travelDistFile, err = os.Create("tmp.xlsx")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(travelDistFile)
}
