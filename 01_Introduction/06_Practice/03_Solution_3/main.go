package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)

	fmt.Println("-")

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println("s =", s)
}
