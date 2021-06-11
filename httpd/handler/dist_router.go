package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pcrandall/travelDist/httpd/platform/distances"
	"github.com/rs/cors"
)

func ChiRouter(port string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
	})
	r := chi.NewRouter()
	dist := distances.New()

	r.Get("/dist", DistGet(dist))
	r.Post("/dist", DistPost(dist))

	fmt.Println("Backend Server is ready and is listening at port :8001...")
	log.Fatal(http.ListenAndServe(port, c.Handler(r)))
}
