package main

import "fmt"

var a uint
var b uint8
var x int
var y int8
var z float64

func main() {
	a = 7
	b = 128
	x = -7
	y = -128
	z = 9.24

	fmt.Println("a =", a)
	fmt.Printf("%T\n", a)
	fmt.Println("b =", b)
	fmt.Printf("%T\n", b)

	fmt.Println("-")

	fmt.Println("x =", x)
	fmt.Printf("%T\n", x)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)

	fmt.Println("-")

	fmt.Println("z =", z)
	fmt.Printf("%T\n", z)
}
