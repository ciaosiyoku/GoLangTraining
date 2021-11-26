package main

import "fmt"

//Main line
func main() {
	//Print Hello world!!!
	fmt.Println("Hello world!!!")
	fmt.Println("-")

	//Declare a variable  and assign a value
	dvsv()

	//Exploring types
	exty()

	//Exploring fmt packages
	exfp()

	//Creating and Converting User-Defined Types
	cudt()
}

//Declare a variable  and assign a value
func dvsv() {
	// var keyword
	var a = 12
	fmt.Println("a =", a)
	var b = 13
	fmt.Println("b =", b)
	var c = 14
	fmt.Println("c =", c)
	var d int
	fmt.Println("d =", d)
	var e bool
	fmt.Println("e =", e)
	var f string
	fmt.Println("f =", f)
	fmt.Println("-")

	// Short declaration
	x := 42
	fmt.Println("x =", x)
	y := 50 + 24
	fmt.Println("y =", y)
	z := "z = Osiyoku, Adeolu"
	fmt.Println(z)

	fmt.Println("-")

	// Reassign a value to an already declared variable
	x = 77
	fmt.Println("x new value is", x)
	fmt.Println("-")
}

// Exploring types
func exty() {
	var h = 33
	fmt.Println("h =", h)
	fmt.Printf("%T\n", h)
	fmt.Println("-")
	var i = "Adeh is learning golang"
	fmt.Println("i =", i)
	fmt.Printf("%T\n", i)
	fmt.Println("-")
	var j = true
	fmt.Println("j =", j)
	fmt.Printf("%T\n", j)
	fmt.Println("-")
}

//Exploring fmt package
func exfp() {
	var l = 42

	//Format printing
	fmt.Println(l)
	fmt.Printf("%T\n", l)
	fmt.Printf("%b\n", l)
	fmt.Printf("%x\n", l)
	fmt.Printf("%#x\n", l)
	fmt.Printf("%#x\t%b\t%T\t%x\n", l, l, l, l)

	fmt.Println("-")

	//Sprint Printing
	m := fmt.Sprintf("%#x\t%b\t%T\t%x\n", l, l, l, l)
	fmt.Println("m =", m)
}

//Creating and Converting User-Defined Types
var n int

type Adeh int

var o Adeh

func cudt() {
	n = 42
	fmt.Println("n =", n)
	fmt.Printf("%T\n", n)

	o = 43
	fmt.Println("o =", o)
	fmt.Printf("%T\n", o)

	//Converting Types
	n = int(o)
	fmt.Println("n =", n)
	fmt.Printf("%T\n", n)
}
