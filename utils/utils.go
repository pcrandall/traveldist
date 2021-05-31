package utils

import (
	"fmt"
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

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func WrapError(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func DebugErr(err error, str string) {
	if err != nil {
		e := WrapError(err, str)
		log.Println(e)
	}
}

func CheckErr(err error, str string) {
	if err != nil {
		e := WrapError(err, str)
		log.Println(e)
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
