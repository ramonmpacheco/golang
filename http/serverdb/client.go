package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userById(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:12345678@/cursogo")

	if err != nil {
		log.Fatal((err))
	}
	defer db.Close()

	var u user
	db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.Id, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:12345678@/cursogo")

	if err != nil {
		log.Fatal((err))
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()
	var users []user

	for rows.Next() {
		var user user
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
