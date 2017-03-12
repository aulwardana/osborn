package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url : %s Path : %s\n", r.Host, r.URL.Path)
	fmt.Fprintf(w, "Mongodb : %s\n", cnf.MongoDB().Hosts)
	fmt.Fprintf(w, "Postgres : %s:%d\n", cnf.Postgres().Host, cnf.Postgres().Port)
	fmt.Fprintf(w, "Redis : %s:%d\n", cnf.Redis().Host, cnf.Redis().Port)
	fmt.Fprintf(w, "Mqtt : %s:%d\n", cnf.Mqtt().Url, cnf.Mqtt().Port)
}
