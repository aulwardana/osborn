package main

import (
	"log"
	"net/http"
	"osborn/core/modules/templates"
)

func main() {
	port := "8000"

	template.ServeAngular()

	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
