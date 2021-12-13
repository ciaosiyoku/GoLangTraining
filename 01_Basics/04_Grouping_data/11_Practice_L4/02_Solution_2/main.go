package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8, 5, 1, 3, 5, 7, 9}

	for i, v := range x {
		fmt.Printf("For the position of %v\t we have value %v\n", i, v)
	}

	fmt.Println("")

	fmt.Printf("%T\n", x)
}
