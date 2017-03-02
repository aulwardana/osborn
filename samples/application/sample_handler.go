package main

import (
	"fmt"
	"net/http"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo %s %s\n", r.Host, r.URL.Path)
	fmt.Println("halo")
}
