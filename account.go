package yoblog

// Account represent blog user
type Account struct {
	ID        string `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Name      string `db:"name" json:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

// AccountStore inteface
type AccountStore interface {
	Create(account *Account) (accountID string, err error)
	GetByID(id string) (account Account, err error)
	Delete(id string) (err error)
}
