package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"osborn/core/libs/mongodb"
	"osborn/samples/application/sensor"
)

func initDBNoSql() *mgo.Session {
	return mongodb.Connect(cnf.MongoDB().Hosts, cnf.MongoDB().User, cnf.MongoDB().Password, cnf.MongoDB().DBName)
}

func InsertData(w http.ResponseWriter, r *http.Request) {
	sensing := &sensor.SensorInsert{
		Code:        "1a",
		Temperature: 10,
		Humidity:    20,
	}

	sns := sensor.New(initDBNoSql(), cnf.MongoDB().DBName)
	err := sns.InsertSensorData(sensing)
	if err != nil {
		return
	}
}
