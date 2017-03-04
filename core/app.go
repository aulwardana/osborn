package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(r *mux.Router, address string, port int) {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", address, port),
	}

	http.Handle("/", r)
	log.Println("Server started:" + server.Addr)
	log.Fatal(server.ListenAndServe())
}
