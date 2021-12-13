package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("")

	x[1] = 7
	x[4] = 9
	x[7] = 5

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("")

	x = append(x, 17, 24, 37, 12)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("")

	x = append(x, 11, 16, 22, 41)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("")

	x = append(x, 19, 22, 51, 76)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("")

	x = append(x, 17, 24, 37, 12, 19, 22, 51, 19, 22, 51, 76, 78)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
