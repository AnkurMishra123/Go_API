package utilities

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	dataSourceName := "username:password@tcp(localhost:3306)/yourdb"

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database!")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
