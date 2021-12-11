package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 17}
	fmt.Println(x)
	x = append(x, 27, 32, 67)
	fmt.Println(x)

	y := []int{78, 88, 89, 99}
	x = append(x, y...)
	fmt.Println(x)

	//to delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
