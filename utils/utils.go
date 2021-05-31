package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/pcrandall/travelDist/utils/stripansi"
	"github.com/pkg/errors"
)

var (
	err         error
	logfile     *os.File
	regexNumber = regexp.MustCompile(`^\d+$`) // regexNumber is a regex that matches a string that looks like an integer
)

func init() {
	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}
	//Initialize
	logfile, err = os.OpenFile("./logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	CheckErr(err, "Error Creating logfile.txt")
	log.SetOutput(logfile)
	defer logfile.Close()
}

// only prints errors to log
func DebugErr(err error, str string) {
	if err != nil {
		log.Println(errors.Wrap(err, str))
	}
}

// prints to log then panic if err
func CheckErr(err error, str string) {
	if err != nil {
		log.Panicln(errors.Wrap(err, str))
	}
}

// GenerateUUID returns a uuid v4 in string
func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "generating uuid")
	}

	return u.String(), nil
}

// IsNumber checks if the given string is in the form of a number
func IsNumber(s string) bool {
	if s == "" {
		return false
	}
	return regexNumber.MatchString(s)
}

// 2020-09-02 15:16:15:415 --> 2020-09-02 15:16:15
func CleanTimeStamp(s string) string {
	regexTimeStamp := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
	if s == "" {
		return ""
	}
	clean := regexTimeStamp.FindString(s)
	return clean
}

// clean ansi codes and space from string
func StripString(s string) string {
	if s == "" {
		return ""
	}
	// return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
	return stripansi.Strip(strings.TrimSpace(s))
}

// trim space from string
func TrimString(str string) string {
	return strings.TrimSpace(str)
}

// strips all non alphanumeric except ./-_
func StripNONALPHANUMERIC(str string) string {
	const ansi = "[^a-zA-Z0-9_-]*"
	re := regexp.MustCompile(ansi)
	return re.ReplaceAllString(str, "")
}

// func GetConfig(str string, config interface{}) {
// 	fmt.Println("Bfore: ", config)
// 	if _, err := os.Stat(str); err == nil { // check if config file exists
// 		yamlFile, err := ioutil.ReadFile(str)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = yaml.Unmarshal(yamlFile, config)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		// } else if os.IsNotExist(err) { // config file not included, use embedded config
// 		// 	yamlFile, err := Asset(str)
// 		// 	if err != nil {
// 		// 		panic(err)
// 		// 	}
// 		// 	err = yaml.Unmarshal(yamlFile, &config)
// 		// 	if err != nil {
// 		// 		panic(err)
// 		// 	}
// 		// } else {
// 	} else {
// 		CheckErr(err, "Schrodinger: file may or may not exist. See err for details.")
// 	}

// 	fmt.Println("After: ", config)
// }
