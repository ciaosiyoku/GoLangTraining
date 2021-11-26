package main

import "fmt"

func main() {
	// var keyword
	var a = 12
	fmt.Println("a =", a)
	var b = 13
	fmt.Println("b =", b)
	var c = 14
	fmt.Println("c =", c)
	var d int
	fmt.Println("d =", d)
	var e bool
	fmt.Println("e =", e)
	var f string
	fmt.Println("f =", f)
	fmt.Println("-")

	// Short declaration
	x := 42
	fmt.Println("x =", x)
	y := 50 + 24
	fmt.Println("y =", y)
	z := "z = Osiyoku, Adeolu"
	fmt.Println(z)

	fmt.Println("-")

	// Reassign a value to an already declared variable
	x = 77
	fmt.Println("x new value is", x)
}
