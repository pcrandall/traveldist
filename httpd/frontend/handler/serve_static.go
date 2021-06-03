package frontend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pcrandall/travelDist/utils"
)

func ServeFrontEnd() {
	port := ":8000"
	// fileServer := http.FileServer(http.Dir("../../../site"))
	fileServer := http.FileServer(http.Dir("./site"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Frontend Server is ready and is listening at port :8000...\n")
	utils.OpenBrowser("http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}
