package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)


const (
	path = "root:password@tcp(mysql:3306)/kartenspielen")
var (
	count = 0
)

func main() {
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