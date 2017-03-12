package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"osborn/core/libs/postgresql"
)

type psqlSession struct {
	db *sql.DB
}

func sessionDBSql(db *sql.DB) *psqlSession {
	return &psqlSession{
		db: db,
	}
}

func initDBSql() {
	sessionDBSql(postgresql.Connect(cnf.Postgres().Host, cnf.Postgres().Port, cnf.Postgres().User, cnf.Postgres().Password, cnf.Postgres().DBName, cnf.Postgres().SSLMode))
}
