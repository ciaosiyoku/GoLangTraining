package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8, 5, 1, 3, 5, 7, 9}

	fmt.Println(x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
