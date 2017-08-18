package yoblog

// Post represent blog post
type Post struct {
	ID        string `db:"id"`
	OwnerID   string `db:"owner_id"`
	OwnerName string `db:"name"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type PostStore interface {
	Create(post *Post) (postID string, err error)
	GetByID(id string) (post Post, err error)
	GetByOwnerID(ownerID string) (posts []Post, err error)
	GetAll() (posts []Post, err error)
	Delete(id string) (err error)

	CreateComment(comment *Comment) (commentID string, err error)
	GetPostComments(postID string) (comments []Comment, err error)
}
