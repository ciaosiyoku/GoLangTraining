package main

import "fmt"

func main() {
	fmt.Println("Hello world!!!")

	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started was:", x)
	}
	g(1984)
}
