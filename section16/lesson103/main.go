package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer DB.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
			name STRING
			age INT)`

	_, err := Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}
}
