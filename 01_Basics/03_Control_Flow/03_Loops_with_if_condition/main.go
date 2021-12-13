package main

import (
	"fmt"
)

func main() {
	// Finding a remainder
	x := 83 / 40
	y := 83 % 40
	fmt.Printf("83/40 = %d\t remainder %d\n", x, y)

	fmt.Println("")

	//for "init statement"; "condition"; "post statement" {}

	//for statement with break condition
	a := 1
	for {
		if a > 5 {
			break
		}
		fmt.Println(a)
		a++
	}
	fmt.Println("Done")

	fmt.Println("")

	//for statement with break and continue condition
	b := 1
	for {
		b++
		if b > 10 {
			break
		}

		if b%2 != 0 {
			continue
		}
		fmt.Println(b)

	}
	fmt.Println("Done")
}
