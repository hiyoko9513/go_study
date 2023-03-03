package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// todo ポスグレの場合のテーブル作成についてまとめる
//func createUsersTable(db *sql.DB) error {
//	cmd := `CREATE TABLE IF NOT EXISTS users (name STRING,age INT)`
//	_, err := db.Exec(cmd)
//	return err
//
//}

func main() {
	var Db *sql.DB
	Db, err := sql.Open("postgres", "user=postgres password=postgres host=postgres port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	//err = createUsersTable(Db)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
