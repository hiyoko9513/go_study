package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createUsersTable(db *sql.DB) error {
	cmd := `CREATE TABLE IF NOT EXISTS users (name STRING,age INT)`
	_, err := db.Exec(cmd)
	return err

}

func main() {
	var Db *sql.DB
	Db, _ = sql.Open("sqlite3", "./db/example.sql")
	defer Db.Close()

	err := createUsersTable(Db)
	if err != nil {
		log.Fatal(err)
	}

}
