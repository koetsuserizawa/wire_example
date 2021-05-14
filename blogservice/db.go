package blogservice

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN string
}

func NewDb(cnf Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cnf.DSN)
	return db, err
}
