package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("./matrix.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		ip := txt[0]
		nav := cleanString(txt[1:])
		fmt.Println("IP", ip, "Nav:", nav)
		res, err := http.Get("http://" + ip + "/srm1TravelDistanceList.html")
		body, err := ioutil.ReadAll(res.Body)

		pageContent := string(body)

		r := regexp.MustCompile(`<tr.*?>(.*)</tr>`)

		submatchall := r.FindAllStringSubmatch(pageContent, -1)

		// fmt.Println("body:", pageContent, "err:", err)
		// if r.MatchString(pageContent) {
		// 	fmt.Println("body:", pageContent, "err:", err)
		// }

		for _, element := range submatchall {
			fmt.Println(element[1], "err", err)
		}

		defer res.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func cleanString(s []string) string {
	//nav := strings.Trim(strings.Replace(strings.Join(txt[1:], " "), "[", "", -1), " ")
	// strings.Join converts the string array to string
	// strings.Replace removes the [] in the string
	// strings.Trim removes the whitespace
	n := strings.Join(s[1:], " ")
	n = strings.Replace(n, "[", "", -1)
	n = strings.Trim(n, " ")
	return n
}
