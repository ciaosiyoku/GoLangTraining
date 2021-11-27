package main

import (
	"fmt"
)

const (
	// untyped constant
	a = 7
	b = 9.23
	//typed constant
	c string = "Adeh"
	d bool   = true
)

func main() {
	fmt.Println("")

	fmt.Println("a =", a)
	fmt.Printf("%T\n", a)

	fmt.Println("")

	fmt.Println("b =", b)
	fmt.Printf("%T\n", b)

	fmt.Println("")

	fmt.Println("c =", c)
	fmt.Printf("%T\n", c)

	fmt.Println("")

	fmt.Println("d =", d)
	fmt.Printf("%T\n", d)

	fmt.Println("")
}
