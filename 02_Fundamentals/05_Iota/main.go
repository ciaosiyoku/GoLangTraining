package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
	g
	h
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
	fmt.Println("-")
	fmt.Println("")

	fmt.Println("d =", d)
	fmt.Printf("%T\n", d)

	fmt.Println("")

	fmt.Println("e =", e)
	fmt.Printf("%T\n", e)

	fmt.Println("")

	fmt.Println("f =", f)
	fmt.Printf("%T\n", f)

	fmt.Println("")

	fmt.Println("g =", g)
	fmt.Printf("%T\n", g)

	fmt.Println("")

	fmt.Println("h =", h)
	fmt.Printf("%T\n", h)

	fmt.Println("")
}
