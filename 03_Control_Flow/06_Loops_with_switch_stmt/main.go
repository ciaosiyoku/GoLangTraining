package main

import (
	"fmt"
)

func main() {
	//Switch on boolean value
	switch {
	case false:
		fmt.Println("Not True 1")
		fallthrough
	case (2 <= 3):
		fmt.Println("True 1")
		fallthrough
	case (2 == 3):
		fmt.Println("Not True 2")
		fallthrough
	case true:
		fmt.Println("True 2")
		fallthrough
	case (2 == 4):
		fmt.Println("Not True 3")
		fallthrough
	case (2 != 3):
		fmt.Println("True 3")
	}

	fmt.Println("")

	//Switch on Default
	switch {
	case (2 == 5):
		fmt.Println("Not True 4")
	default:
		fmt.Println("This is default")
	}

	fmt.Println("")

	// Switch on a value
	a := "Adeh"
	switch a {
	case "Adeh":
		fmt.Println("This is Adeh")
	case "Bond":
		fmt.Println("This is James bond")
	case "Q":
		fmt.Println("This is Q")
	}
}
