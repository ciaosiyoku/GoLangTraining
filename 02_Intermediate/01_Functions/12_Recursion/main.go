package main

import (
	"fmt"
)

func main() {
	var a int = 7
	r := factorial(a)
	fmt.Printf("%v! (recursed) = %v\n", a, r)

	fmt.Println("")

	l := loopFact(a)
	fmt.Printf("%v! (looped) = %v\n", a, l)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Calculate a factorial using a loop
func loopFact(y int) int {
	total := 1
	for ; y > 0; y-- {
		total *= y
	}
	return total
}
