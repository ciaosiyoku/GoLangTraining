package main

import (
	"fmt"
)

//Scope of x
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)

	fmt.Println("")
	//code blocked "y"
	{
		y := 7
		fmt.Println("The value of y is ", y)
	}

	//you can't print "y" here
	//fmt.Println(y)

	fmt.Println("")

	foo()
	fmt.Println(x)

	fmt.Println("")

	bar()
	fmt.Println(x)

	fmt.Println("")

	a := incrementor()
	b := incrementor()

	fmt.Println("Incrementor value a1 is:", a())
	fmt.Println("Incrementor value a2 is:", a())
	fmt.Println("Incrementor value a3 is:", a())
	fmt.Println("Incrementor value a4 is:", a())
	fmt.Println("Incrementor value a5 is:", a())

	fmt.Println("")

	fmt.Println("Incrementor value b1 is:", b())
	fmt.Println("Incrementor value b2 is:", b())
	fmt.Println("Incrementor value b3 is:", b())
	fmt.Println("Incrementor value b4 is:", b())
}

func foo() {
	fmt.Println("Hello Adeh")
	x++
}

func bar() {
	fmt.Println("Hello Osiy")
	x *= 4
}

func incrementor() func() int {
	y := 0
	return func() int {
		y++
		y *= 2
		return y
	}
}
