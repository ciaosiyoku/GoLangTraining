package main

import (
	"fmt"
)

func main() {
	//for "init statement"; "condition"; "post statement" {}

	//for statement with a single condition
	a := 1
	for a <= 9 {
		fmt.Println(a)
		a++
	}
	fmt.Println("Done")

	fmt.Println("")
}
