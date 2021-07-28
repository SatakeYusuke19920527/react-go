package main

import (
	"fmt"
	"strconv"
)

var i5 int = 2000

const Pi = 3.14

func main() {
	interfaceFunc()
	fmt.Println(Pi)
}

func interfaceFunc() {
	var str string = "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(i)
	fmt.Printf("s=%T\n", str)
}
