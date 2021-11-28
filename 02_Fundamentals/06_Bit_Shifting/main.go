package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("")

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	fmt.Println("")

	x := 8
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Println("")

	kb2 := 1024
	mb2 := 1024 * kb2
	gb2 := 1024 * mb2

	fmt.Printf("%d\t\t%b\n", kb2, kb2)
	fmt.Printf("%d\t\t%b\n", mb2, mb2)
	fmt.Printf("%d\t%b\n", gb2, gb2)

	fmt.Println("")
}
