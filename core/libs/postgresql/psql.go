package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(host string, port int, user string, password string, dbname string, sslmode string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conInfo := fmt.Sprintf("%s:%d", host, port)
	log.Println("Postgres started:" + conInfo + " DBName:" + dbname)

	return db
}
