package store

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	connStr := ""
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
