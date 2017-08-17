package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	accountStore AccountStore
)

func init() {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 user=postgres password=123456 dbname=postgres sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}

	accountStore = NewAccountStore(db)
}
