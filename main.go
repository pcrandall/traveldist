package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"github.com/pcrandall/travelDist/workbook"
	// frames "github.com/pcrandall/mfsplaces/frames"
)

var (
	tableString        []shuttleDistance
	getTravelDistances []travelDistances

	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	config     *Config
	err        error
	getNavette string
	writeFile  bool
	clear      map[string]func()
	resize     map[string]func()
	restAPI    bool
)

func main() {
	flag.BoolVar(&writeFile, "w", false, "Write to excel workbook -w=true")
	flag.BoolVar(&restAPI, "r", false, "restAPI -r=true")
	flag.Parse()

	callResize()
	printHeader("TRAVELDIST")

	//Initialize
	logfile, err := os.OpenFile("./logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkErr(err, "Error Creating logfile.txt")
	log.SetOutput(logfile)
	defer logfile.Close()

	if writeFile {
		workbook.FindWorkbook(config.SheetName)
	}

	// front end server
	go ServeFrontEnd()

	if restAPI == false {
		//TODO add channels here
		navettes := config.Levels
		for _, nav := range navettes {
			// fmt.Printf("Floor: [%d]\n", val.Floor)
			for _, n := range nav.Navette {
				getNavette = n.Name
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

				r := regexp.MustCompile(`Total:[\d\s]{9}`)
				// give me all the matches
				submatchall := r.FindAllStringSubmatch(pageContent, -1)
				total := submatchall[len(submatchall)-1][0]
				//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
				//total looks like this now Total: 2912046
				/// total[7:] to trim the string
				row.shuttle = cleanString(n.Name)
				row.timestamp = cleanString(lastDate)
				row.distance = total[7:]
				tableString = append(tableString, *row)

				if writeFile {
					workbook.WriteFile(n.Row, config.SheetName, total[0])
				}
				defer res.Body.Close()
			}
		}

	}

	if writeFile {
		workbook.SaveFile()
	}

	//TODO
	if restAPI == false {
		RenderTable(tableString)
		insertDatabase(tableString)
	} else {
		DBRouter()
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
