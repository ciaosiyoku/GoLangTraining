package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println("")

	fmt.Printf("true && true\t%v\n", true && true)
	fmt.Printf("true && false\t%v\n", true && false)
	fmt.Printf("true || true\t%v\n", true || true)
	fmt.Printf("true || false\t%v\n", true || false)
	fmt.Printf("!true\t\t%v\n", !true)
	fmt.Printf("!false\t\t%v\n", !false)
}
