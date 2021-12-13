package main

import (
	"fmt"
)

func main() {
	// else conditional statement
	a := 42
	if a == 40 {
		fmt.Println("Our value was 40")
	} else {
		fmt.Println("Our value was not 40")
	}

	fmt.Println("")

	// else if conditional statement
	b := 44
	if b == 40 {
		fmt.Println("Our value was 40")
	} else if b == 41 {
		fmt.Println("Our value was 41")
	} else if b == 42 {
		fmt.Println("Our value was 42")
	} else if b == 43 {
		fmt.Println("Our value was 43")
	} else {
		fmt.Println("Our value was not 40, 41, 42 or 43")
	}
}
