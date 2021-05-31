package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"github.com/pcrandall/travelDist/frames"
	"github.com/pcrandall/travelDist/utils"
	"github.com/pcrandall/travelDist/workbook"
)

var (
	tableString []shuttleDistance

	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	config    *Config
	err       error
	writeFile bool
	restAPI   bool
)

func main() {
	flag.BoolVar(&writeFile, "w", false, "Write to excel workbook -w=true")
	flag.BoolVar(&restAPI, "r", false, "restAPI -r=true")
	flag.Parse()

	utils.ResizeWindow()
	utils.PrintHeader("TRAVELDIST")

	//Initialize
	logfile, err := os.OpenFile("./logs/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.CheckErr(err, "Error Creating logfile.txt")
	log.SetOutput(logfile)
	defer logfile.Close()

	//TODO get rid of this
	if writeFile {
		workbook.FindWorkbook(config.SheetName)
	}

	// initiate loading screen
	f := make(chan bool) // loading frames channel
	go frames.Start(f)

	var wg sync.WaitGroup
	if restAPI == false {
		navettes := config.Levels
		for _, nav := range navettes {
			for _, n := range nav.Navette {
				wg.Add(1)
				go ScrapPages(n.Name, n.IP, n.Row, &wg)
			}
		}
	}
	wg.Wait()

	f <- true // send the stop signal to the go func and close channel
	close(f)

	//TODO get rid of this
	if writeFile {
		workbook.SaveFile()
	}

	//TODO
	if restAPI == false {
		// RenderTable(tableString)
		insertDatabase(tableString)
		utils.PrintHeader("TRAVELDIST")
		go ServeFrontEnd() // front end server
		DBRouter()         // backend server
	} else {
		go ServeFrontEnd() // front end server
		DBRouter()         // backend server
	}
}

func ScrapPages(name string, ip string, excelRow string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := client.Get("http://" + ip + "/srm1TravelDistanceList.html")
	utils.DebugErr(err, "Navette:"+name+"IP address:"+ip)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	utils.DebugErr(err, "Navette:"+name+"IP address:"+ip)
	if err != nil {
		return
	}

	pageContent := string(body)

	// parse the page content and pull the relevant values
	dbRow := new(shuttleDistance)
	date := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})[\s\d:]{13}`)
	//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
	// date returns 2020-09-02 15:16:15:415 from the above
	dateMatchAll := date.FindAllStringSubmatch(pageContent, -1)
	lastDate := dateMatchAll[len(dateMatchAll)-1][0]
	log.Println(name, lastDate)

	r := regexp.MustCompile(`Total:[\d\s]{9}`)
	// give me all the matches
	submatchall := r.FindAllStringSubmatch(pageContent, -1)
	total := submatchall[len(submatchall)-1][0]
	//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
	//total looks like this now Total: 2912046
	/// total[7:] to trim the string
	dbRow.shuttle = utils.TrimString(name)
	dbRow.timestamp = utils.TrimString(lastDate)
	dbRow.distance = total[7:]
	tableString = append(tableString, *dbRow)

	if writeFile {
		workbook.WriteFile(excelRow, config.SheetName, total[0])
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
		var row = []string{utils.TrimString(val.shuttle), utils.TrimString(val.distance), utils.TrimString(val.timestamp)}
		table.Append(row)
		// fmt.Println(utils.CleanString(val[0]), utils.CleanString(val[1]), utils.CleanString(val[2]))
	}
	table.Render() // Send output
}
