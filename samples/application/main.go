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
}

func main() {
	settingsInit()

	r := mux.NewRouter()

	api.Route(r, "/halo", dummy, api.GET)
	template.ServeAngular(r, "/", cnf.Web().TemplateDir)

	core.Run(r, cnf.Web().Address, cnf.Web().Port)
}
