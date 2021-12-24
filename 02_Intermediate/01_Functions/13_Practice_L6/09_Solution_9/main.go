package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(a...)
	fmt.Println("The sum of the specified numbers =", s1)

	e1 := even(sum, a...)
	fmt.Println("The sum of the specified even numbers =", e1)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, xe ...int) int {
	var xi []int
	for _, v := range xe {
		if v%2 == 0 {
			xi = append(xi, v)
		}

	}
	return f(xi...)
}
