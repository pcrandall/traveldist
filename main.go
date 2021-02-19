package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
	config        *Config
	err           error
	writeColumn   int
	writeFileName string
	logfile       *os.File
)

func main() {

	logfile, err = os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logfile)

	defer logfile.Close()

	// find current workbook
	re := regexp.MustCompile(`^([a-zA-Z0-9\s_\\.\-\(\):])+\.(xlsx|xlsm)$`)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panic(err)
	}

	for _, file := range files {
		match := re.FindStringSubmatch(file.Name())

		if len(match) > 0 {
			dest := "./old/" + match[0]

			err = Copy(match[0], dest)

			if err != nil {
				log.Println("Latest travel distance file must be in the same directory as Travel Distance.exe")
				log.Panic(err)
			}
			writeColumn, err = findLastColumn(match[0], config.SheetName)
			writeFileName = match[0]
		}
	}

	// fmt.Println("writeColumn here: ", writeColumn)
	// fmt.Println("writeFileName here: ", writeFileName)

	file, err := excelize.OpenFile(writeFileName)
	if err != nil {
		log.Println("Latest travel distance file must be in the same directory as Travel Distance.exe")
		log.Panic(err)
		return
	}

	for _, val := range config.Levels {
		// fmt.Printf("Floor: [%d]\n", val.Floor)
		for _, v := range val.Navette {
			// fmt.Printf("Name: %s, IP: %s, Row: %s\n", v.Name, v.IP, v.Row)
			res, err := client.Get("http://" + v.IP + "/srm1TravelDistanceList.html")
			if err != nil {
				log.Println("Navette:", v.Name, "IP address:", v.IP, "Error:", err)
				fmt.Println("Navette:", v.Name, "IP address:", v.IP, "Error:", err)
				continue
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(v.Name, v.IP, err)
				fmt.Println(v.Name, v.IP, err)
				continue
			}

			pageContent := string(body)

			//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
			// this returns all elements in the page that look like the above
			r := regexp.MustCompile(`Total:(.\d*|\s*|\d*)`)

			// give me all the matches
			submatchall := r.FindAllStringSubmatch(pageContent, -1)

			for index, element := range submatchall {
				// the most recent update is the last element, grab the distance from there.
				if index == len(submatchall)-1 {
					total := element[0]
					// total is this:
					// Total: 15440653
					if _, err := strconv.ParseFloat(total[7:], 64); err == nil {
						row, _ := strconv.Atoi(v.Row)
						writeCoord := getColumn(writeColumn, row)
						// fmt.Println("getColumn here: ", writeCoord)
						fmt.Println(v.Name, total[7:])
						file.SetCellValue(config.SheetName, writeCoord, total[7:])
					}
				}
			}
			defer res.Body.Close()
		}
	}

	if err := file.Save(); err != nil {
		log.Panic(err)
	}

	userInput := ""
	fmt.Println("[Press enter to close]")
	fmt.Scanf("%s", &userInput)
}

func getColumn(idx int, row int) string {
	var coord string
	base := 26

	if idx > base {
		if idx < int(math.Pow(float64(base), 2)) {
			first := idx / base
			rem := idx % base
			// fmt.Println(fmt.Sprintf("%x", first))
			// fmt.Println(fmt.Sprintf("%x", rem))
			coord = string(first+64) + string(rem+64) + strconv.Itoa(row)
			// fmt.Println(string(first+64), string(rem+64))
		} else {
			log.Panic("13 years worth of columns is a long time. Great Job! It's time to start over now. Delete columns on page", config.SheetName)
		}
	} else {
		coord = string(idx + 64)
	}
	// fmt.Println(coord)
	return coord
}

func findLastColumn(inFile string, sheetName string) (write int, err error) {
	s, err := excelize.OpenFile(inFile)
	if err != nil {
		return -1, err
	}
	rows := s.GetRows(sheetName)
	for rowInt, row := range rows {
		for colInt, colCell := range row {
			if rowInt == 5 {
				value := colCell
				if value != "" {
					// fmt.Println(fmt.Sprintf("column: %v, value: %v", colInt+1, value))
					writeColumn = colInt + 1
				}
			}
		}
		if rowInt == 5 {
			break
		}
	}
	// rows start at 1, idx starts a 0, add one.
	writeColumn++
	return writeColumn, err
}
