package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pcrandall/travelDist/httpd/platform/distances"
	"github.com/pcrandall/travelDist/utils"
)

type Success struct {
	Nice string `json:"Success"`
}

type Failure struct {
	Err error `json:"Error"`
}

func PostDistance(dist distances.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var change ChangeShoe
		utils.HttpPrettyPrintRequest(r)
		w.Header().Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(&change) // encode to json and send to client
		// fmt.Printf("change: %#v", change)

		greatJob := &Success{"Great Success!"}
		err := json.NewEncoder(w).Encode(&greatJob) // encode to json and send to client
		utils.CheckErr(err, "JSON encoding error")
		if err != nil {
			e := &Failure{err}
			json.NewEncoder(w).Encode(e) // encode to json and send to client
		}

	}
}
