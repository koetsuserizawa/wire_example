package blogservice

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DbName string
}

func NewDb(cnf DbConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cnf.User, cnf.Pass, cnf.Host, cnf.Port, cnf.DbName)
	db, err := sql.Open("mysql", dsn)
	return db, err
}
