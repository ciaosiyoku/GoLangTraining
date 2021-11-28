package main

import "fmt"

func main() {
	var x int = 77
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	fmt.Println("")

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
