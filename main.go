package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/theckman/yacspin"
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	config        *Config
	logfile       *os.File
	err           error
	writeRows     map[string]string
	writeColumn   int
	writeFileName string
	oldFilename   string
	getNavette    string

	clear map[string]func()
)

func main() {
	//Initialize
	writeRows = make(map[string]string)
	cfg := yacspin.Config{
		Frequency:       200 * time.Millisecond,
		CharSet:         yacspin.CharSets[32],
		Suffix:          "Getting Travel Distances",
		SuffixAutoColon: false,
		Message:         getNavette,
		StopCharacter:   "√",
		StopMessage:     " Completed!",
		StopColors:      []string{"fgGreen"},
		Colors:          []string{"fgCyan"},
	}

	// printHeader("TRAVELDIST")

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

	file, err := excelize.OpenFile(writeFileName)
	if err != nil {
		log.Println("Latest travel distance file must be in the same directory as Travel Distance.exe")
		log.Panic(err)
		return
	}

	spinner, err := yacspin.New(cfg) // handle the error
	if err != nil {
		panic(err)
	}
	spinner.Start() // Start the spinner

	for _, val := range config.Levels {
		// fmt.Printf("Floor: [%d]\n", val.Floor)
		for _, v := range val.Navette {
			getNavette = v.Name
			writeRows[v.Name] = v.Row
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
					total := element[0] // Total: 15440653
					if _, err := strconv.ParseFloat(total[7:], 64); err == nil {
						row, _ := strconv.Atoi(v.Row)
						writeCoord := buildCoordinate(writeColumn, row)
						// file.SetCellValue(config.SheetName, writeCoord, total[7:])
						td, _ := strconv.Atoi(total[7:]) // convert to int here so workbook doesn't complain
						file.SetCellValue(config.SheetName, writeCoord, td)
					}
				}
			}
			defer res.Body.Close()
		}
	}

	spinner.Stop() // connected stop spinner

	oldFilename = writeFileName // store filename for comparison later

	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Panic(err)
	}

	// check that the new filename is not empty, if so use old filename
	if writeFileName == "" {
		if err := file.SaveAs(oldFilename); err != nil {
			log.Panic(err)
		}
	} else {
		// renamed file delete old file
		if err := file.SaveAs(writeFileName); err != nil {
			log.Panic(err)
		}
		if err = os.Remove(oldFilename); err != nil {
			log.Panic(err)
		}
	}
}

func buildCoordinate(n int, row int) string {
	// https://www.geeksforgeeks.org/find-excel-column-name-given-number
	str, result := "", ""
	for n > 0 {
		rem := n % 26
		if rem == 0 {
			str += string(90) // string(90) = "Z"
			n = (n / 26) - 1
		} else {
			str += string(rem + 64)
			n /= 26
		}
	}
	for _, v := range str {
		result = string(v) + result
	}
	// fmt.Println("coordinate here: ", result+strconv.Itoa(row))
	return result + strconv.Itoa(row)
}

func findLastColumn(inFile string, sheetName string) (write int, err error) {
	s, err := excelize.OpenFile(inFile)
	if err != nil {
		return -1, err
	}
	rows := s.GetRows(sheetName)
	// fmt.Println("row here: ", rows[5])
	// fmt.Println("row here: ", len(rows[5]))
	val := ""
	row := rows[5]

	for i := len(row) - 1; val == ""; i-- {
		val = string(row[i])
		if val != "" {
			writeColumn = i + 1
		}
	}
	writeColumn++
	// fmt.Println("writeColumn here: ", writeColumn)
	// fmt.Println("writeColumn here: ", buildCoordinate(writeColumn, 6))
	return writeColumn, err
}
