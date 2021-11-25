package main

import "fmt"

func main() {
	fmt.Println("Hello world!!!")

	Adeh1()

}

//Declare a variable  and assign a value
func Adeh1() {
	// var keyword
	var a = 12
	fmt.Println("a =", a)
	var b = 13
	fmt.Println("b =", b)
	var c = 14
	fmt.Println("c =", c)

	// Short declaration
	x := 42
	fmt.Println("x =", x)
	y := 100 + 24
	fmt.Println("y =", y)
	z := "z = Osiyoku, Adeolu"
	fmt.Println(z)

	// Reassign a value to an already declared variable
	x = 99
	fmt.Println("x new value is", x)

}
