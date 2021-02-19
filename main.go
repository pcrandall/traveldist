package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	output, err := os.Create("collectorShoesDistances.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	for _, val := range config.Levels {
		// fmt.Printf("Floor: [%d]\n", val.Floor)
		for _, v := range val.Navette {
			// fmt.Printf("Name: %s, IP: %s, Row: %s\n", v.Name, v.IP, v.Row)
			res, err := client.Get("http://" + v.IP + "/srm1TravelDistanceList.html")
			if err != nil {
				io.WriteString(output, "¯\\_(ツ)_/¯\n")
				fmt.Println("¯\\_(ツ)_/¯:", err)
				continue
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("¯\\_(ツ)_/¯:", err)
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
					if _, err := strconv.ParseFloat(f[7:], 64); err == nil {
						// if s, err := strconv.ParseFloat(f[7:], 64); err == nil {
						// km := s / 10000
						//skm := strconv.FormatFloat(km, 'f', 0, 64)
						// fmt.Printf("Collector Shoe distance:\n\t %s\n\t %.f km\n\n", f[7:], km)
						io.WriteString(output, v.Name+"  "+f[7:]+"\n")
						output.Sync()
						// fmt.Println("\t", f[7:])
					}
				}
			}

			defer res.Body.Close()
		}
	}

	//for scanner.Scan() {
	//	txt := strings.Split(scanner.Text(), " ")
	//	ip := txt[0]
	//	nav := cleanString(txt[1:])

	//	fmt.Println(nav, "\nURL:", "http://"+ip+"/srm1TravelDistanceList.html")

	//	Were writing to the text file here
	//	_, err = io.WriteString(output, nav+"\n")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	output.Sync()

	//	res, err := client.Get("http://" + ip + "/srm1TravelDistanceList.html")
	//	if err != nil {
	//		io.WriteString(output, "¯\\_(ツ)_/¯ Guess I'm real dead\n")
	//		fmt.Println("ERROR:", err, "\n")
	//		continue
	//	}

	//	body, err := ioutil.ReadAll(res.Body)
	//	if err != nil {
	//		fmt.Println("ERROR:", err, "\n")
	//		continue
	//	}

	//	pageContent := string(body)

	//	//<td>I</td><td>2020-09-02 15:16:15:415</td><td id="desc"></td><td>TD Total: 4598010 4598010 </td><td>Td: 0 4598010 </td></td> err <nil>
	//	// this returns all elements in the page that look like the above
	//	r := regexp.MustCompile(`Total:(.\d*|\s*|\d*)`)

	//	// give me all the matches
	//	submatchall := r.FindAllStringSubmatch(pageContent, -1)

	//	for index, element := range submatchall {
	//		// the most recent update is the last element, grab the distance from there.
	//		if index == len(submatchall)-1 {
	//			f := element[0]
	//			if s, err := strconv.ParseFloat(f[7:], 64); err == nil {
	//				km := s / 10000
	//				//skm := strconv.FormatFloat(km, 'f', 0, 64)
	//				fmt.Printf("Collector Shoe distance:\n\t %s\n\t %.f km\n\n", f[7:], km)
	//				io.WriteString(output, nav+"  "+f[7:]+"\n")
	//				output.Sync()
	//				// fmt.Println("\t", f[7:])
	//			}
	//		}
	//	}

	//	defer res.Body.Close()
	//}

	// // TODO implement excel things
	// 	file, err := excelize.OpenFile(os.Args[1])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// Get all the rows in the Sheet1.
	// 	rows := file.GetRows("Sheet1")
	// 	for _, row := range rows {
	// 		level(row[1], flr, row)
	// 	}

	// // Create new file and write sorted rows
	// sortedFile := excelize.NewFile()

	// sortedFile.SetColWidth("Sheet1", "A", "C", 25)

	// // style, err := sortedFile.NewStyle(`"alignment": {"horizontal": "right"}`)
	// // sortedFile.SetColStyle("Sheet1", "B:C", style)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var rowIndex int

	// sortedFile.SetSheetRow("Sheet1", "A1", &[]interface{}{"STOLOC", "LUID", "Verified LUID"})

	// for idx, sortedFlr := range flr.f4 {
	// 	// Sheet to write, row index to write, convert slice to interface
	// 	sortedFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(idx+2), &[]interface{}{sortedFlr[1], sortedFlr[2], sortedFlr[3]})
	// 	rowIndex = idx + 2
	// 	// fmt.Println("f4", sortedFlr)
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	//}

	// TODO you don't need this
	// func cleanString(s []string) string {
	// 	// strings.Join converts the string array to string
	// 	// strings.Replace removes the [] in the string
	// 	// strings.Trim removes the whitespace
	// 	n := strings.Join(s[1:], " ")
	// 	n = strings.Replace(n, "[", "", -1)
	// 	n = strings.Trim(n, " ")
	// 	return n
	// }
}
