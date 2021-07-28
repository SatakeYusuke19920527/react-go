package main

import "fmt"

var i5 int = 2000

func main() {
	var i int = 100
	fmt.Println(i)
	var t, f bool = true, false
	fmt.Println(t, f)
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	s3 := "GOGO"
	fmt.Println(s3)

	s3 = "NONO"
	fmt.Println(s3)

	fmt.Println(i5)
	outer()
	floatFunc()
}

func outer() {
	var s4 string = "s4 in outer"
	fmt.Printf("%T\n", s4)
}

func floatFunc() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	fmt.Printf("%T\n", fl64)
}
