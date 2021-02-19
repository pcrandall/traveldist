package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Levels []struct {
		Floor   int `yaml:"floor"`
		Navette []struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
			Row  string `yaml:"row"`
		} `yaml:"navette"`
	} `yaml:"levels"`
}

var (
	config *Config

	client = &http.Client{
		Timeout: 10 * time.Second,
	}
)

func init() {

	newpath := filepath.Join(".", "old")
	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		os.MkdirAll(newpath, os.ModePerm)
		// os.Mkdir("scannedIPs", 777)
	}

	travelDistFile, err := os.Create("tmp.xlsx")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(travelDistFile)

	GetConfig()
}

func GetConfig() {
	if _, err := os.Stat("./config/config.yml"); err == nil { // check if config file exists
		yamlFile, err := ioutil.ReadFile("./config/config.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) { // config file not included, use embedded config
		yamlFile, err := Asset("config/config.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Schrodinger: file may or may not exist. See err for details.")
		// panic(err)
	}
}
