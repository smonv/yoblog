package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/tthanh/yoblog/postgres"
	"github.com/tthanh/yoblog/service"
)

var schema = `
CREATE TABLE IF NOT EXISTS account (
	id SERIAL PRIMARY KEY,
	email VARCHAR(50) NOT NULL UNIQUE,
	name VARCHAR(50) NOT NULL,
	created_at INTEGER,
	updated_at INTEGER
);
`

func main() {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 user=postgres password=123456 dbname=postgres sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.MustExec(schema)

	accountStore := postgres.NewAccountStore(db)

	service := service.New(accountStore)

	r := mux.NewRouter()

	r.HandleFunc("/", service.IndexHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
