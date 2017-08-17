package yoblog

// Account represent blog user
type Account struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

// AccountStore inteface
type AccountStore interface {
	Create(account *Account) (accountID int, err error)
	GetByID(id int) (account Account, err error)
	Delete(id int) (err error)
}
