package foo

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db, err := sql.Open("mysql", "root:s29y27a25@/mydb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// m := getUsers(db, err)
	// fmt.Println(m)
	// insertData(db, err, m)

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

func insertData(db *sql.DB, err error, m map[int]string) {

	fmt.Println("----- Exec -----")
	ins, err := db.Prepare("INSERT INTO test (id,name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	id := len(m) + 1
	name := strconv.Itoa(id) + "satake"
	ins.Exec(id, name)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
