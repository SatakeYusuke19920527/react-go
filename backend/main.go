package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

type MyInt int

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u *User) SetName() {
	u.Name = "A"
}

type Users []*User

func main() {

	m := map[int]User{
		1: {"user1", 29},
		2: {"user2", 30},
		3: {"user3", 31},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{"user1", 29}: "tokyo",
		{"user2", 30}: "osaka",
		{"user3", 31}: "la",
	}

	fmt.Println(m2)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	var myInt MyInt = 100

	fmt.Println(myInt)

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
