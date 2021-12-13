package main

import "fmt"

func main() {
	x := [5]int{2, 4, 6, 8, 5}

	for i, v := range x {
		fmt.Printf("For the position of %v\t we have value %v\n", i, v)
	}

	fmt.Println("")

	fmt.Printf("%T\n", x)
}
