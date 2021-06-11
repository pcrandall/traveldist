package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pcrandall/travelDist/httpd/platform/distances"
)

func DistPost(feed distances.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request := map[string]string{}
		// json.NewDecoder(r.Body).Decode(&request)

		// feed.Add(distances.Item{
		// 	Title: request["title"],
		// 	Post:  request["post"],
		// })

		// w.Write([]byte("Good job!"))

		var change ChangeShoe
		w.Header().Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(&change) // encode to json and send to client
		fmt.Printf("change: %#v\n\n", *r)
		log.Printf("change: %#v\n\n", *r)

		params := mux.Vars(r)
		log.Printf("params: %v\n\n", params)
		log.Printf("requestURI: %v\n\n", r.RequestURI)
		fmt.Printf("requestURI: %v\n\n", r.RequestURI)

		for v, p := range params {
			fmt.Printf("iter\n\n")
			fmt.Printf("v: %s p: %#v\n\n", v, p)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		log.Println("database.go.98: ", body)
		fmt.Println("database.go.98: ", body)
		// fmt.Println(json.NewEncoder(w).Encode(&body)) // encode to json and send to client)

	}
}
