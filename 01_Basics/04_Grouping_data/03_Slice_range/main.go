package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 17}
	fmt.Println(len(x))

	fmt.Println("")

	fmt.Println(x)

	fmt.Println("")

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	fmt.Println("")

	for i, v := range x {
		fmt.Printf("For the position of %v\t we have value of %v\n", i, v)
	}
}
