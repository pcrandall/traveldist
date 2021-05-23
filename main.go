package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"github.com/theckman/yacspin"
	// frames "github.com/pcrandall/mfsplaces/frames"
)

type shuttleDistance struct {
	shuttle   string
	distance  string
	timestamp string
}

var (
	tableString []shuttleDistance
	client      = &http.Client{
		Timeout: 10 * time.Second,
	}

	cfg = yacspin.Config{
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

	config        *Config
	err           error
	excelFile     *excelize.File
	getNavette    string
	oldFilename   string
	writeFile     bool
	writeFileName string
	writeRows     map[string]string
	excelColumn   int
	clear         map[string]func()
	resize        map[string]func()
)

func main() {
	printHeader("TRAVELDIST")

	flag.BoolVar(&writeFile, "w", false, "Write to excel workbook -w=true")
	flag.Parse()

	//Initialize
	writeRows = make(map[string]string)
	logfile, err := os.OpenFile("./logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkErr(err, "Error Creating logfile.txt")
	log.SetOutput(logfile)
	defer logfile.Close()

	if writeFile {
		findWorkbook()
	}

	spinner, err := yacspin.New(cfg) // handle the error
	checkErr(err, "yacspin error: ")
	spinner.Start() // Start the spinner

	//TODO add channels here
	navettes := config.Levels
	for _, nav := range navettes {
		// fmt.Printf("Floor: [%d]\n", val.Floor)
		for _, n := range nav.Navette {
			getNavette = n.Name
			writeRows[n.Name] = n.Row
			// fmt.Printf("Name: %s, IP: %s, Row: %s\n", v.Name, v.IP, v.Row)
			res, err := client.Get("http://" + n.IP + "/srm1TravelDistanceList.html")
			checkErr(err, "Navette:"+n.Name+"IP address:"+n.IP)
			if err != nil {
				continue
			}

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "Navette:"+n.Name+"IP address:"+n.IP)
			if err != nil {
				continue
			}
			pageContent := string(body)

			// parse the page content and pull the relevant values
			row := new(shuttleDistance)
			date := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})[\s\d:]{13}`)
			//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
			// date returns 2020-09-02 15:16:15:415 from the above
			dateMatchAll := date.FindAllStringSubmatch(pageContent, -1)
			lastDate := dateMatchAll[len(dateMatchAll)-1][0]
			log.Println(n.Name, lastDate)

			r := regexp.MustCompile(`Total:[\d\s]{8}`)
			// give me all the matches
			submatchall := r.FindAllStringSubmatch(pageContent, -1)
			total := submatchall[len(submatchall)-1][0]
			//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
			//total looks like this now Total: 2912046
			/// total[7:] to trim the string
			row.shuttle = n.Name
			row.timestamp = lastDate
			row.distance = total[7:]
			tableString = append(tableString, *row)

			if writeFile {
				excelRow, _ := strconv.Atoi(n.Row)
				writeCoord := buildCoordinate(excelColumn, excelRow)
				excelFile.SetCellValue(config.SheetName, writeCoord, total[0])
			}
			defer res.Body.Close()
		}
	}

	spinner.Stop() // connected stop spinner

	RenderTable(tableString)

	insertDatabase(tableString)
	if writeFile {
		oldFilename = writeFileName // store filename for comparison later
		p := tea.NewProgram(initialModel())
		if err := p.Start(); err != nil {
			log.Panic(err)
		}

		if writeFileName == "" { // check that the new filename is not empty, if so use old filename
			if err := excelFile.SaveAs(oldFilename); err != nil {
				log.Panic(err)
			}
		} else {
			if err := excelFile.SaveAs(writeFileName); err != nil { // renamed file delete old file
				log.Panic(err)
			}
			if err = os.Remove(oldFilename); err != nil {
				log.Panic(err)
			}
		}
	}
}

func RenderTable(locations []shuttleDistance) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Shuttle", "Distance", "Timestamp"})
	table.SetTablePadding("\t") // pad with tabs
	table.SetAutoWrapText(true)
	table.SetAutoFormatHeaders(true)
	table.SetRowLine(true)

	for _, val := range locations {
		var row = []string{cleanString(val.shuttle), cleanString(val.distance), cleanString(val.timestamp)}
		table.Append(row)
		// fmt.Println(cleanString(val[0]), cleanString(val[1]), cleanString(val[2]))
	}
	table.Render() // Send output
}
