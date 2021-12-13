package main

import "fmt"

type Adeh int

var x Adeh

func main() {
	fmt.Println("x =", x)
	fmt.Printf("%T\n", x)

	fmt.Println("")

	x = 42
	fmt.Println("x again is", x)
}
