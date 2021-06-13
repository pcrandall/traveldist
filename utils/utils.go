package utils

import (
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/pcrandall/travelDist/utils/stripansi"
	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	err         error
	Logfile     *os.File
	ErrLog      *log.Logger
	regexNumber = regexp.MustCompile(`^\d+$`) // regexNumber is a regex that matches a string that looks like an integer
)

func InitLog() {
	Logfile, err = os.OpenFile("logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	ErrLog = log.New(Logfile, "", log.Ldate|log.Ltime|log.LstdFlags)
	ErrLog.SetOutput(&lumberjack.Logger{
		Filename:   "logs/logfile.txt",
		MaxSize:    25, // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})

	log.SetOutput(Logfile)

	defer Logfile.Close()
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
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)
		ErrLog.Panicf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, errors.Wrap(err, str))
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
