package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "webinar/pkg/test"

	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	log.Default().Println("Pinged!")

	const q = `
		INSERT INTO users(id, username) VALUES (2, "user1"), (3, "user2")
	`

	// "SELECT username FROM users WHERE id = ?"

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err.Error())
	}

	var usernames []string

	for rows.Next() {
		var username string

		if err := rows.Scan(&username); err != nil {
			log.Fatal(err.Error())
		}

		usernames = append(usernames, username)
	}

	fmt.Println(len(usernames), usernames)
}
