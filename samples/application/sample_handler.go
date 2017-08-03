package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type UsersData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
	City  string `json:"city"`
}

type Users []UsersData

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url : %s Path : %s\n", r.Host, r.URL.Path)
	fmt.Fprintf(w, "Mongodb : %s\n", cnf.MongoDB().Hosts)
	fmt.Fprintf(w, "Postgres : %s:%d\n", cnf.Postgres().Host, cnf.Postgres().Port)
	fmt.Fprintf(w, "Redis : %s:%d\n", cnf.Redis().Host, cnf.Redis().Port)
	fmt.Fprintf(w, "Mqtt : %s:%d\n", cnf.Mqtt().Url, cnf.Mqtt().Port)

	v, _ := mem.VirtualMemory()
	h, _ := host.Info()

	fmt.Fprintf(w, "\nTotal Memory: %v, Free Memory:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Fprintf(w, "OS: %v, Kernel Version:%v, Uptime:%d\n", h.OS, h.KernelVersion, h.Uptime)
}

func APITestPost(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	var usrDec UsersData
	err := body.Decode(&usrDec)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(usrDec)
}

func APITestGet(w http.ResponseWriter, r *http.Request) {
	users := Users{
		UsersData{Name: "Aul", Email: "auldesain@gmail.com", Age: 25, City: "Malang"},
		UsersData{Name: "Embuh", Email: "embuh@gmail.com", Age: 20, City: "Surabaya"},
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}
