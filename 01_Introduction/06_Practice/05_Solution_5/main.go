package main

import "fmt"

type Adeh int

var x Adeh
var y int

func main() {
	fmt.Println("x =", x)
	fmt.Printf("%T\n", x)

	fmt.Println("-")

	x = 77
	fmt.Println("x again is", x)

	fmt.Println("-")

	y = int(x)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)
}
