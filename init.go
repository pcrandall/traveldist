package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pcrandall/travelDist/utils"
	"gopkg.in/yaml.v2"
)

// Config
type Config struct {
	Levels         []Levels       `yaml:"levels"`
	View           View           `yaml:"view"`
	Controller     Controller     `yaml:"controller"`
	ShoeParameters ShoeParameters `yaml:"shoe_change"`
}

// Levels
type Levels struct {
	Floor   int       `yaml:"floor"`
	Navette []Navette `yaml:"navette"`
}

// Navette
type Navette struct {
	Ip   string `yaml:"ip"`
	Name string `yaml:"name"`
}

// View
type View struct {
	Port string `yaml:"port"`
}

// Controller
type Controller struct {
	Port string `yaml:"port"`
}

// ShoeCheck
type ShoeParameters struct {
	Check    int `yaml:"check"`
	Interval int `yaml:"interval"`
}

func init() {
	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}

	dbpath := filepath.Join(".", "db")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		os.MkdirAll(dbpath, os.ModePerm)
	}
	GetConfig()
}

func GetConfig() {
	if _, err := os.Stat("./config/config.yml"); err == nil { // check if config file exists
		yamlFile, err := ioutil.ReadFile("./config/config.yml")
		utils.CheckErr(err, "init: error reading ./config/config.yml")
		err = yaml.Unmarshal(yamlFile, &config)
		utils.CheckErr(err, "init: error unmarshalling yaml into &config")
	} else if os.IsNotExist(err) { // config file not included, use embedded config
		yamlFile, err := Asset("./config/config.yml")
		utils.CheckErr(err, "init: error reading Asset(./config/config.yml)")
		err = yaml.Unmarshal(yamlFile, &config)
		utils.CheckErr(err, "init: error unmarshalling yaml into &config")
	} else {
		log.Panic("Schrodinger: file may or may not exist. See err for details.")
	}
}
