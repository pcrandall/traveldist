package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		nav := strings.Split(scanner.Text(), " ")
		ip := nav[0]
		nNav := nav[2:]
		fmt.Println("IP", ip, "Nav:", nNav)

		res, err := http.Get("http://" + ip + "/srm1TravelDistanceList.html")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		pageContent := string(body)
		fmt.Println("body:", pageContent, "err:", err)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
