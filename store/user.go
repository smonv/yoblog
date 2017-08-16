package store

import "github.com/jmoiron/sqlx"

// User implement yoblog.UserStore interface
type User struct {
	db *sqlx.DB
}
