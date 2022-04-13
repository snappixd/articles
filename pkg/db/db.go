package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(user, password, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, dbName))
	if err != nil {
		return nil, err
	}

	return db, nil
}
