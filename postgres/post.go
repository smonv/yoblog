package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/tthanh/yoblog"
)

// PostStore implement yoblog.PostStore
type PostStore struct {
	db *sqlx.DB
}

// NewPostStore ...
func NewPostStore(db *sqlx.DB) PostStore {
	return PostStore{
		db: db,
	}
}

func (s PostStore) Create(post *yoblog.Post) (postID string, err error) {
	post.ID = uuid.NewV4().String()

	now := time.Now()

	post.CreatedAt = now.Unix()
	post.UpdatedAt = now.Unix()

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(
		"INSERT INTO post VALUES($1, $2, $3, $4, $5, $6)",
		post.ID,
		post.OwnerID,
		post.Title,
		post.Content,
		post.CreatedAt,
		post.UpdatedAt,
	)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return "", sErr
		}

		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return post.ID, nil
}

func (s PostStore) GetByID(id string) (post yoblog.Post, err error) {
	err = s.db.Get(
		&post,
		"SELECT account.name, post.* FROM post INNER JOIN account ON post.owner_id = account.id WHERE post.id = $1",
		id,
	)

	return
}

func (s PostStore) GetByOwnerID(ownerID string) (posts []yoblog.Post, err error) {
	err = s.db.Select(
		&posts,
		"SELECT account.name, post.* FROM post INNER JOIN account ON post.owner_id = account.id WHERE post.owner_id = $1",
		ownerID,
	)

	return
}

func (s PostStore) GetAll() (posts []yoblog.Post, err error) {
	err = s.db.Select(&posts, "SELECT account.name, post.* FROM post INNER JOIN account ON post.owner_id = account.id")

	return
}

func (s PostStore) Delete(id string) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM comment WHERE comment.post_id = $1", id)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return sErr
		}

		return err
	}

	_, err = tx.Exec("DELETE FROM post WHERE id = $1", id)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return sErr
		}

		return err
	}

	return nil
}

func (s PostStore) CreateComment(comment *yoblog.Comment) (commentID string, err error) {
	comment.ID = uuid.NewV4().String()

	now := time.Now()

	comment.CreatedAt = now.Unix()
	comment.UpdatedAt = now.Unix()

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(
		"INSERT INTO comment VALUES($1, $2, $3, $4, $5, $6)",
		comment.ID,
		comment.OwnerID,
		comment.PostID,
		comment.Content,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	if err != nil {
		sErr := tx.Rollback()
		if sErr != nil {
			return "", sErr
		}

		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return comment.ID, nil
}

func (s PostStore) GetPostComments(postID string) (comments []yoblog.Comment, err error) {
	err = s.db.Select(
		&comments,
		"SELECT account.name, comment.* FROM comment INNER JOIN account ON comment.owner_id = account.id WHERE comment.post_id = $1",
		postID,
	)

	return
}
