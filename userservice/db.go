package userservice

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	host   string
	port   int
	user   string
	pass   string
	dbName string
}

func NewDb(cnf DbConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cnf.user, cnf.pass, cnf.host, cnf.port, cnf.dbName)
	db, err := sql.Open("mysql", dsn)
	return db, err
}
