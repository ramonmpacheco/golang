package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stm, _ := db.Prepare("INSERT INTO users(name) values(?)")
	stm.Exec("Mary")
	stm.Exec("John")

	res, _ := stm.Exec("Petter")

	id, _ := res.LastInsertId()
	fmt.Print("Petter id ", id)

	rows, _ := res.RowsAffected()
	fmt.Print("Rows affected ", rows)
}
