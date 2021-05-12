package userservice

import "database/sql"

type User struct {
	id   string
	name string
}

type UserRepositoryInterface interface {
	Register(string) (*User, error)
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

func (r *UserRepository) Register(name string) (*User, error) {
	return &User{
		id:   "",
		name: name,
	}, nil
}

func (r *UserRepository) Get(id string) (*User, error) {
	return &User{
		id:   "",
		name: "",
	}, nil
}
