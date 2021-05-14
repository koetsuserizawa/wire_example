package blogservice

import (
	"database/sql"
	"fmt"
)

type Article struct {
	UserID string
	Text   string
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
	q := "SELECT user_id, article_text FROM article WHERE user_id = ?"
	st, err := r.db.Prepare(q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := st.Query(userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var ar []Article
	for rows.Next() {
		var t Article
		err := rows.Scan(&t.UserID, &t.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ar = append(ar, t)
	}
	return ar, nil
}

func (r *BlogRepository) Write(userID string, text string) error {

	q := "INSERT INTO article(user_id, article_text) VALUES(?, ?)"
	in, err := r.db.Prepare(q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	in.Exec(userID, text)
	return nil
}
