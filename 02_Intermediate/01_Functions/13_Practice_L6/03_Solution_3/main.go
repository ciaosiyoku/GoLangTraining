package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	var x int = 24
	fmt.Println("The value of func foo is", x)
}

func bar() {
	var y int = 27
	fmt.Println("The value of func bar is", y)
}
