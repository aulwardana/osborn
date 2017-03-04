package main

import (
	"fmt"
	"log"
	"net/http"
	"osborn/core/config"
	"osborn/core/modules/api"
	"osborn/core/modules/templates"

	"github.com/gorilla/mux"
)

var (
	cnf *config.Config
)

func main() {
	if err := initConfig(); err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	api.Route(r, "/halo", dummy, api.GET)
	template.ServeAngular(r, "/", cnf.Web().TemplateDir)

	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cnf.Web().Address, cnf.Web().Port),
	}

	http.Handle("/", r)
	log.Println("Server started:" + server.Addr)
	log.Fatal(server.ListenAndServe())
}
