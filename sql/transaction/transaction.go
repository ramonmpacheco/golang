package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tran, _ := db.Begin()
	stmt, _ := tran.Prepare("INSERT INTO users(id, name) values(?,?)")

	stmt.Exec(2000, "John")
	stmt.Exec(2001, "Mary")
	_, err = stmt.Exec(1, "Petter") //duplicated

	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}

	tran.Commit()
}
