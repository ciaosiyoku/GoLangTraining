package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a) //The "&" operator shows the adderess in memory
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	fmt.Println("")

	b := &a
	fmt.Println(b)
	fmt.Println(*b) //The "*" operator gives you the value stored at an address

	fmt.Println("")

	*b = 43
	fmt.Println("The new value of a =", a)
}
