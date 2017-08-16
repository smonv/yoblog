package yoblog

// User represent blog user
type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt int
	UpdatedAt int
}

// Post represent blog post
type Post struct {
	ID        string
	OwnerID   string
	Title     string
	Content   string
	CreatedAt int
	UpdatedAt int
}
