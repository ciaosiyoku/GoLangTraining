package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total = ", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//To add up the variadic values
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For index position", i, "we are now adding", v, "which is now", sum)
	}

	return sum
}
