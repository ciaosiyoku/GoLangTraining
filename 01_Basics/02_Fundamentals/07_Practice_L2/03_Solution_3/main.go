package main

import "fmt"

const (
	// untyped constant
	a = 7

	//typed constant
	b string = "Adeh"
)

func main() {
	fmt.Println("a =", a)
	fmt.Printf("%T\n", a)

	fmt.Println("")

	fmt.Println("b =", b)
	fmt.Printf("%T\n", b)
}
