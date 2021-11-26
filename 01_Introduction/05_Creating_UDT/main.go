package main

import "fmt"

var n int

type Adeh int

var o Adeh

func main() {

	//Creating User-Defined Types
	n = 42
	fmt.Println("n =", n)
	fmt.Printf("%T\n", n)

	o = 43
	fmt.Println("o =", o)
	fmt.Printf("%T\n", o)

	//Converting Types
	n = int(o)
	fmt.Println("n =", n)
	fmt.Printf("%T\n", n)
}
