package main

import "fmt"

//Main line
func main() {
	//Print Hello world!!!
	fmt.Println("Hello world!!!")
	fmt.Println("-")

	//Declare a variable  and assign a value
	dvsv()

}

//Declare a variable  and assign a value
func dvsv() {
	// var keyword
	var a = 12
	fmt.Println("a =", a)
	var b = 13
	fmt.Println("b =", b)
	var c = 14
	fmt.Println("c =", c)

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
