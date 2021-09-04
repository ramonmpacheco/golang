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

	stm, _ := db.Prepare("UPDATE users set name = ? where id = ?")
	stm.Exec("Anne", 1)

	stmt2, _ := db.Prepare("DELETE FROM users where id = ?")
	stmt2.Exec(3)
}
