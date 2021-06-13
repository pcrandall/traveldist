package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pcrandall/travelDist/frames"
	controllerHandler "github.com/pcrandall/travelDist/httpd/handler"
	"github.com/pcrandall/travelDist/httpd/platform/shoeparameters"
	"github.com/pcrandall/travelDist/utils"
	viewHandler "github.com/pcrandall/travelDist/view/handler"
)

var (
	Version     = "v0.1"
	ProgramName = "TRAVELDIST"
	client      = &http.Client{
		Timeout: 10 * time.Second,
	}
	config               *Config
	distancesSlice       []controllerHandler.ShuttleDistance
	loadingFramesChannel = make(chan bool)
	wg                   sync.WaitGroup

	err       error
	restAPI   bool
	writeFile bool
)

func main() {
	flag.BoolVar(&writeFile, "w", false, "Write to excel workbook -w=true")
	flag.BoolVar(&restAPI, "r", false, "restAPI -r=true")
	flag.Parse()

	// initialize logging
	utils.InitLog()
	log.SetOutput(utils.Logfile)

	// resize console window to default size
	utils.ResizeWindow()

	// print out program header
	utils.PrintHeader(ProgramName, Version)

	utils.ErrLog.Println("Travel Distances started")
	// // For testing only
	// utils.CheckErr(fmt.Errorf("BIG BAD ERROR\n"), "PLZ NO MORE!!!")

	// set the config parameters
	shoeparameters.SetShoeParameters(config.ShoeParameters.Check, config.ShoeParameters.Interval)

	if restAPI == true {
		go viewHandler.ServeView(config.View.Port) // front end server
		controllerHandler.ChiRouter(config.Controller.Port)
	} else {
		// initiate loading screen
		go frames.Start(loadingFramesChannel)
		navettes := config.Levels
		for _, nav := range navettes {
			for _, n := range nav.Navette {
				wg.Add(1)
				// read the level control web interface
				go ScrapPages(n.Name, n.Ip, &wg)
			}
		}
		// wait for all concurrent jobs to finish
		wg.Wait()

		loadingFramesChannel <- true // send the stop signal to the go func and close channel
		close(loadingFramesChannel)

		// insert new values into the DISTANCE table
		controllerHandler.InsertDistance(distancesSlice)
		utils.PrintHeader(ProgramName, Version)

		// Start servers
		go viewHandler.ServeView(config.View.Port)
		controllerHandler.ChiRouter(config.Controller.Port)
	}
}

func ScrapPages(name string, ip string, wg *sync.WaitGroup) {
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
	dbRow := new(controllerHandler.ShuttleDistance)
	// dbRow := &handler.ShuttleDistance{}

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

	dbRow.Shuttle = utils.TrimString(name)
	dbRow.Timestamp = utils.TrimString(lastDate)
	dbRow.Distance = total[7:]
	distancesSlice = append(distancesSlice, *dbRow)
}
