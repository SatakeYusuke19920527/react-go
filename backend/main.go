package main

import (
	"fmt"
	"time"
)

func reciever(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)
	go reciever("1,goroutin", ch1)
	go reciever("2,goroutin", ch1)
	go reciever("3,goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(300 * time.Millisecond)
}
