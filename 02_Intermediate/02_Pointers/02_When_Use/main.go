package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("The first value of x =", x)
	fmt.Println("The position of value x =", &x)

	fmt.Println("")

	foo(&x)

	fmt.Println("")

	fmt.Println("The second value of x =", x)
	fmt.Println("The position of value x =", &x)
}

func foo(y *int) {
	fmt.Println("The first value of y =", y)
	fmt.Println("The first value of y =", *y)

	fmt.Println("")

	*y = 43
	fmt.Println("The second value of y =", y)
	fmt.Println("The second value of y =", *y)
}
