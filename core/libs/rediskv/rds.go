package rediskv

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

func Connect(host string, port int, password string) *redis.Pool {
	hosts := fmt.Sprintf("%s:%d", host, port)

	log.Println("Redis started:" + hosts)

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hosts)
			if err != nil {
				panic(err.Error())
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
