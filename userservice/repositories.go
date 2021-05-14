package userservice

import "database/sql"

type User struct {
	ID   string
	Name string
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
		ID:   "",
		Name: name,
	}, nil
}

func (r *UserRepository) Get(id string) (*User, error) {
	return &User{
		ID:   "",
		Name: "",
	}, nil
}
