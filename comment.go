package yoblog

// Comment represent post comment
type Comment struct {
	ID        string `db:"id"`
	OwnerID   string `db:"owner_id"`
	OwnerName string `db:"name"`
	PostID    string `db:"post_id"`
	Content   string `db:"content"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}
