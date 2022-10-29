package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	path := createPath()
	db, err := sql.Open("mysql", path);if err != nil {
		log.Println(err)
		return
	}
	c, ioErr := ioutil.ReadFile("./init.sql")
	if ioErr != nil {
		log.Println(ioErr)
	}
	requests := strings.Split(string(c), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			log.Println(err)
		}
	}

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