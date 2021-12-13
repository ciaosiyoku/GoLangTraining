package main

import "fmt"

func main() {
	a := 42
	if a == 40 {
		fmt.Println("Our value was 40")
	} else if a == 41 {
		fmt.Println("Our value was 41")
	} else {
		fmt.Println("Our value was not 40 or 41")
	}

}
