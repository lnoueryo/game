package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:password@tcp(172.18.0.2:3306)/kartenspielen")
	if err != nil {
		panic(err)
	}
	DB = db
}