package blogservice

import (
	"database/sql"

	"local.packages/userserivce"
)

type Article struct {
	user userserivce.User
	text string
}

type BlogRepositoryInterface interface {
	Read(string) ([]Article, error)
	Write(string, string) error
}

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) BlogRepositoryInterface {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) Read(userID string) ([]Article, error) {
	var ar []Article
	return ar, nil
}

func (r *BlogRepository) Write(userID string, text string) error {
	return nil
}
