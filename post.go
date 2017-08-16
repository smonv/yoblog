package yoblog

// Post represent blog post
type Post struct {
	ID        string
	OwnerID   string
	Title     string
	Content   string
	CreatedAt int
	UpdatedAt int
}
