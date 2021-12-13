package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)

	fmt.Println("")

	x[2] = 7
	fmt.Println(x)

	fmt.Println("")

	fmt.Println("The length of the array is", len(x))
}
