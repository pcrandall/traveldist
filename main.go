package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./matrix.txt")
	// Can't wait forever if a Navette is powered off
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output, err := os.Create("Collector Shoe Distances.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		ip := txt[0]
		nav := cleanString(txt[1:])

		fmt.Println(nav, "\nURL:", "http://"+ip+"/srm1TravelDistanceList.html")

		// Were writing to the text file here
		// _, err = io.WriteString(output, nav+"\n")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// output.Sync()

		res, err := client.Get("http://" + ip + "/srm1TravelDistanceList.html")
		if err != nil {
			io.WriteString(output, "¯\\_(ツ)_/¯ Guess I'm real dead\n")
			fmt.Println("ERROR:", err, "\n")
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("ERROR:", err, "\n")
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
				f := element[0]
				if s, err := strconv.ParseFloat(f[7:], 64); err == nil {
					km := s / 10000
					//skm := strconv.FormatFloat(km, 'f', 0, 64)
					fmt.Printf("Collector Shoe distance:\n\t %s\n\t %.f km\n\n", f[7:], km)
					io.WriteString(output, nav+"  "+f[7:]+"\n")
					output.Sync()
					// fmt.Println("\t", f[7:])
				}
			}
		}

		defer res.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func cleanString(s []string) string {
	// strings.Join converts the string array to string
	// strings.Replace removes the [] in the string
	// strings.Trim removes the whitespace
	n := strings.Join(s[1:], " ")
	n = strings.Replace(n, "[", "", -1)
	n = strings.Trim(n, " ")
	return n
}
