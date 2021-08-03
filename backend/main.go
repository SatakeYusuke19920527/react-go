package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m := getUsers(db, err)
	fmt.Println(m)

}

// 全userをget
func getUsers(db *sql.DB, err error) map[int]string {
	var (
		id   int
		name string
	)
	m := make(map[int]string)
	rows, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name)
		checkError(err)
		m[id] = name
	}
	return m
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
