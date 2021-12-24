package main

import "fmt"

func main() {
	a := foo()
	fmt.Println("func foo  =", a)

	fmt.Println("")

	xi, xs := bar()
	fmt.Printf(`func bar says "%v %v"`, xi, xs)
}

func foo() int {
	var x int = 6
	x *= 2
	return x
}

func bar() (string, int) {
	return "Big Brother started watching in", 1984
}
