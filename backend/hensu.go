package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init")
}

func main() {
	go sub()
	go sub()
	for {
		fmt.Println("main loop")
		time.Sleep(2000 * time.Millisecond)
	}
}

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(1000 * time.Millisecond)
	}
}
