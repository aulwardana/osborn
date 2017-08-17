package main

import (
	"osborn/core"
	"osborn/core/config"
	"osborn/core/modules/api"
	"osborn/core/modules/templates"

	"github.com/gorilla/mux"
)

var (
	cnf *config.Config
)

func settingsInit() {
	if err := initConfig(); err != nil {
		panic(err)
	}

	initDBSql()
	initDBNoSql()
	initDBKV()
	initMqtt()
}

func main() {
	settingsInit()
	FlatProc()

	r := mux.NewRouter()

	/*For get server monitoring*/
	api.Route(r, "/server", server, api.GET)

	/*For testing nosql connection*/
	api.Route(r, "/insertDataMgo", InsertDataMgo, api.GET)

	/*For testing sql connection*/
	//api.Route(r, "/getDataPql", GetDataPql, api.GET)
	api.Route(r, "/insertDataPql", InsertDataPql, api.GET)
	//api.Route(r, "/updateDataPql", UpdateDataPql, api.PUT)
	//api.Route(r, "/deleteDataPql", DeleteDataPql, api.DELETE)

	/*For testing api connection*/
	api.Route(r, "/testGetAPI", APITestGet, api.GET)
	api.Route(r, "/testPostAPI", APITestPost, api.POST)

	/*Serve template / handling template*/
	template.ServeAngular(r, "/", cnf.Web().TemplateDir)

	/*Start middleware using osborn framework*/
	core.Run(r, cnf.Web().Address, cnf.Web().Port)
}
