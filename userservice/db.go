package userservice

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserDB struct {
	DB *sql.DB
}

func NewDb(dsn string) *UserDB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return &UserDB{
		DB: db,
	}
}
