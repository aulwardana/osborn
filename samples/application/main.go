package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s %s\n", r.Host, r.URL.Path)
	})

	http.HandleFunc("/req", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s\n", r.Host, r.URL.Path)
	})

	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
