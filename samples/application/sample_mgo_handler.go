package main

import (
	"gopkg.in/mgo.v2"

	"osborn/core/libs/mongodb"
)

type mongoSession struct {
	db *mgo.Session
}

func sessionDBNoSql(db *mgo.Session) *mongoSession {
	return &mongoSession{
		db: db,
	}
}

func initDBNoSql() {
	sessionDBNoSql(mongodb.Connect(cnf.MongoDB().Hosts, cnf.MongoDB().User, cnf.MongoDB().Password, cnf.MongoDB().DBName))
}
