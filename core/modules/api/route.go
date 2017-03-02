package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

func Route(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), method string) {
	r.HandleFunc(path, handler).Methods(method)
}
