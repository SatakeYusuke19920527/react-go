package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Book Struct (Model)

type Book struct {
	ID   string
	NAME string
}

// Get All Books
// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(books)
// }

func main() {

	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	// 全User取得
	cmd := "select * from test"
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

	// Initiate Router
	r := mux.NewRouter()

	// Route Hnadlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", r))
}

func getDB() sql.DB {
	db, err := sql.Open("postgres", "postgres://SatakeYusuke19920527:s29y27a25@host:5432/mydb")
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()
	return *db
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	w.Header().Set("Content-Type", "application/json")
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
	json.NewEncoder(w).Encode(m)
}
