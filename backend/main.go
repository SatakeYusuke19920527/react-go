package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func UpdateUser(user User) {
	user.Name = "update"
	user.Age = 30
}

func UpdateUser2(user *User) {
	user.Name = "update"
	user.Age = 30
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 29
	fmt.Println(user1)

	user2 := User{"user2", 29}
	fmt.Println(user2)

	user3 := User{"user3", 29}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	user5 := new(User)
	fmt.Println(user5)
	user6 := &User{}
	fmt.Println(user6)

	UpdateUser(user3)
	UpdateUser2(&user4)

	fmt.Println(user3)
	fmt.Println(user4)
}
