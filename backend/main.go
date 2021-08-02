package main

import "fmt"

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	m2 := map[string]int{
		"A": 100,
		"B": 200,
		"C": 300,
	}
	fmt.Println(m2)

}
