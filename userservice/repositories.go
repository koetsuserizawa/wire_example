package userservice

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	Register(string, string) error
	Get(string) (*User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(id string, name string) error {
	q := "INSERT INTO user(user_id, user_name) VALUES(?, ?)"
	in, err := r.db.Prepare(q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	in.Exec(id, name)
	return nil
}

func (r *UserRepository) Get(id string) (*User, error) {
	var u User
	q := "SELECT user_id, user_name FROM user WHERE user_id = ?"
	err := r.db.QueryRow(q, id).Scan(&u.ID, &u.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &u, nil
}
