package main

import (
	"github.com/garyburd/redigo/redis"

	"osborn/core/libs/rediskv"
)

type redisSession struct {
	db *redis.Pool
}

func sessionDBKV(db *redis.Pool) *redisSession {
	return &redisSession{
		db: db,
	}
}

func initDBKV() {
	sessionDBKV(rediskv.Connect(cnf.Redis().Host, cnf.Redis().Port, cnf.Redis().Password))
}
