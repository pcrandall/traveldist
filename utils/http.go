package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func HttpPrettyPrintRequest(r *http.Request) ([]byte, error) {
	requestDump, err := httputil.DumpRequest(r, true)
	// Save a copy of this request for debugging.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	return requestDump, err
}
