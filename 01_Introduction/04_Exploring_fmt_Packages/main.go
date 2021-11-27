package main

import "fmt"

func main() {
	//Exploring fmt package
	var l = 42

	//Format printing
	fmt.Println(l)
	fmt.Printf("%T\n", l)
	fmt.Printf("%b\n", l)
	fmt.Printf("%x\n", l)
	fmt.Printf("%#x\n", l)
	fmt.Printf("%#x\t%b\t%T\t%x\n", l, l, l, l)

	fmt.Println("")

	//Sprint Printing
	m := fmt.Sprintf("%#x\t%b\t%T\t%x\n", l, l, l, l)
	fmt.Println("m =", m)
}
