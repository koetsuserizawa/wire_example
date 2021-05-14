package blogservice

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ArticleDB struct {
	DB *sql.DB
}

func NewDb(dsn string) *ArticleDB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return &ArticleDB{
		DB: db,
	}
}
