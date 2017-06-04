package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"osborn/core/libs/postgresql"
	"osborn/samples/application/administration"
)

func initDBSql() *sql.DB {
	return postgresql.Connect(cnf.Postgres().Host, cnf.Postgres().Port, cnf.Postgres().User, cnf.Postgres().Password, cnf.Postgres().DBName, cnf.Postgres().SSLMode)
}

func InsertDataPql(w http.ResponseWriter, r *http.Request) {
	user := &administration.UsersInsert{
		Name:  "aul",
		Email: "a@a.com",
		Age:   25,
		City:  "Malang",
	}

	usr := administration.New(initDBSql())
	err := usr.Create(user)
	if err != nil {
		return
	}
}
