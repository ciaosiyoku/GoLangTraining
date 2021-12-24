package main

import "fmt"

func main() {
	x := func() {
		fmt.Println(`Assigned variable "x" to func`)
	}
	x()
	fmt.Printf("%T\n", x)

	y := func(yi int) {
		fmt.Println(`The year I started "go programming" was`, yi)
	}
	y(2021)
	fmt.Printf("%T\n", y)

	fmt.Println("")

	fmt.Println("Done")
}
