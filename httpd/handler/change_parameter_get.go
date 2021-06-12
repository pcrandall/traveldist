package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pcrandall/travelDist/httpd/platform/shoeparameters"
	"github.com/pcrandall/travelDist/utils"
)

func GetShoeParameters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := shoeparameters.GetShoeParameters()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(params) // encode to json and send to client
		utils.CheckErr(err, "JSON encoding error")
	}
}
