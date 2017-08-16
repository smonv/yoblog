package yoblog

// User represent blog user
type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt int
	UpdatedAt int
}

// UserStore inteface
type UserStore interface {
	Create(user *User) (userID string, err error)
}
