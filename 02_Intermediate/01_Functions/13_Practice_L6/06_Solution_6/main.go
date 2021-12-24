package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous func ran")
	}()

	fmt.Println("")

	func(x int) {
		fmt.Println("Anonymous func int ran", x)
	}(35)
}
