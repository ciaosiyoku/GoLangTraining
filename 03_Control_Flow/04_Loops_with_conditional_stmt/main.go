package main

import (
	"fmt"
)

func main() {
	//bool conditional statement
	a := 2
	b := 2

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if a == b {
		fmt.Println("003")
	}

	if a != b {
		fmt.Println("004")
	}

	fmt.Println("")

	//the "not" operator function
	if !true {
		fmt.Println("005")
	}

	if !false {
		fmt.Println("006")
	}

	if !(a == b) {
		fmt.Println("007")
	}

	if !(a != b) {
		fmt.Println("008")
	}

	fmt.Println("")

	//the initialization statement
	if c := 42; c == 42 {
		fmt.Println("009")
	}
}
