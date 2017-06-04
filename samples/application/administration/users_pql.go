package administration

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type UsersServicePql struct {
	db *sql.DB
}

func (d *UsersServicePql) Create(r *UsersInsert) error {
	var lastInsertId int
	var err error
	err = d.db.QueryRow("INSERT INTO public.users(name, email, age, city) VALUES($1,$2,$3,$4) returning uid;", "aul", "a@a.com", 12, "malang").Scan(&lastInsertId)
	log.Println("last inserted id =", lastInsertId)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func New(db *sql.DB) UsersService {
	svc := &UsersServicePql{
		db: db,
	}

	return svc
}
