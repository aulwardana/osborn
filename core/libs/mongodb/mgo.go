package mongodb

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func Connect(hosts string, username string, password string, database string) *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	log.Println("MongoDB started:" + hosts + " DBName:" + database)

	return session
}
