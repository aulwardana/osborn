package main

import (
	"log"
	"net/http"
	"osborn/core/modules/api"
	"osborn/core/modules/templates"

	"github.com/gorilla/mux"
)

func main() {
	port := "8000"
	r := mux.NewRouter()

	api.Route(r, "/halo", dummy, api.GET)
	template.ServeAngular(r, "/", "./assets/")

	http.Handle("/", r)
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
