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
	id CHAR(36) PRIMARY KEY,
	email VARCHAR(50) UNIQUE,
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

	srv, err := service.New(
		service.SetAccountStore(accountStore),
		service.SetCookieStore([]byte("secret")),
	)

	r := mux.NewRouter()

	r.HandleFunc("/", srv.IndexHandler).Methods("GET")
	r.HandleFunc("/login", srv.LoginHandler).Methods("GET")
	r.HandleFunc("/callback", srv.CallbackHandler).Methods("GET")
	r.HandleFunc("/logout", srv.LogoutHandler).Methods("GET")

	httpSrv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(httpSrv.ListenAndServe())
}
