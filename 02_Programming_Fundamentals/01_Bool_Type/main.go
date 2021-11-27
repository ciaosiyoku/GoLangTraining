package main

import "fmt"

var x bool

func main() {
	fmt.Println("x is", x)
	x = true
	fmt.Println("x is now", x)

	fmt.Println("")

	//Double equal sign is for Equality Comparison
	//Single equal sign makes it an Assignment/Operator
	a := 7
	b := 14
	fmt.Println(a == b)
}
