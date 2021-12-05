package main

import "fmt"

func main() {
	for a := 10; a <= 100; a++ {

		x := a / 4
		y := a % 4

		fmt.Printf("%v divided 4 = %d remainder %d\n", a, x, y)

		//It can also be printed like this
		//fmt.Printf("When %v is divided by 4, the remainder is %v\n", a, a%4)

	}
}
