package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// db取得
	db := getDB()
	m := getUsers(db)
	insertUser(m, db)
	deleteUser(m, db)
	fmt.Println(m)
}

func getDB() sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()
	return *db
}

func getUsers(db sql.DB) map[int]string {
	var (
		id   int
		name string
	)
	// 全User取得
	cmd := "select id, name from test"
	//取得するデータが1件の場合は、QueryRowも利用できる
	rows, _ := db.Query(cmd)
	defer rows.Close()

	m := make(map[int]string)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			break
		}
		m[id] = name
	}
	return m
}

func insertUser(m map[int]string, db sql.DB) {
	num := len(m) + 1
	uname := "Hello"
	rows, err := db.Query("insert into test values($1,$2)", num, uname)
	if err != nil {
		log.Fatalln("insert失敗", err)
	}
	defer rows.Close()
}

func deleteUser(m map[int]string, db sql.DB) {
	rows, err := db.Query("delete from test where id = $1", len(m))
	if err != nil {
		log.Fatalln("insert失敗", err)
	}
	defer rows.Close()
}
