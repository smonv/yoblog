package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/tthanh/yoblog"
)

// AccountStore implement yoblog.AccountStore interface
type AccountStore struct {
	db *sqlx.DB
}

// NewAccountStore create new store.User
func NewAccountStore(db *sqlx.DB) AccountStore {
	return AccountStore{db: db}
}

// Create implement AccountStore.Create
func (a AccountStore) Create(account *yoblog.Account) (accountID int, err error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	err = tx.QueryRow(
		"INSERT INTO account (email, name, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id",
		account.Email,
		account.Name,
		account.CreatedAt,
		account.UpdatedAt,
	).Scan(&accountID)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return -1, sErr
		}

		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return
}

// GetByID implement AccountStore.GetByID
func (a AccountStore) GetByID(id int) (account yoblog.Account, err error) {
	err = a.db.Get(&account, "SELECT * FROM account WHERE id = $1", id)

	return
}

// Delete implement AccountStore.Delete
func (a AccountStore) Delete(id int) (err error) {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM account WHERE id = $1", id)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return sErr
		}

		return err
	}

	return nil
}
