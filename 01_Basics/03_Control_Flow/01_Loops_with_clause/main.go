package main

import (
	"fmt"
)

func main() {
	//for "init statement"; "condition"; "post statement" {}

	//for statement with for clause
	//Outer loop
	for a := 0; a <= 9; a++ {

		//Inner loop
		for b := 0; b < 3; b++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", a, b)
		}
		fmt.Println("Done")
	}

	fmt.Println("")
}
