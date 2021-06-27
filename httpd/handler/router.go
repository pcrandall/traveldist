package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	change "github.com/pcrandall/travelDist/httpd/platform/change_shoes"
	check "github.com/pcrandall/travelDist/httpd/platform/check_shoes"
	"github.com/pcrandall/travelDist/httpd/platform/distances"
	"github.com/rs/cors"
)

func ChiRouter(port string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
	})
	r := chi.NewRouter()
	distance := distances.New()
	change := change.New()
	check := check.New()

	r.Get("/dist", GetDistances(distance))
	r.Get("/distparam", GetShoeParameters())
	r.Post("/changeshoes", InsertChange(change))
	r.Post("/checkshoes", InsertCheck(check))

	fmt.Println("Backend Server is ready and is listening at port :8001...")
	log.Fatal(http.ListenAndServe(port, c.Handler(r)))
}
