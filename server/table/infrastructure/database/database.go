package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {

	path := createPath()
	db, err := sql.Open("mysql", path);if err != nil {
		panic(err)
	}
	DB = db
}

func createPath() string {
	DB_HOST := os.Getenv("DB_HOST");if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	DB_PORT := os.Getenv("DB_PORT");if DB_PORT == "" {
		DB_PORT = "3306"
	}
	DB_NAME := os.Getenv("DB_NAME");if DB_NAME == "" {
		DB_NAME = "kartenspielen"
	}
	DB_USER := os.Getenv("DB_USER");if DB_USER == "" {
		DB_USER = "root"
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD");if DB_PASSWORD == "" {
		DB_PASSWORD = "password"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
}