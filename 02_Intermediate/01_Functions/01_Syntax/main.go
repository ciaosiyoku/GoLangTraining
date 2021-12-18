package main

import "fmt"

func main() {
	foo()
	bar("Adeh")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Adeh", "Osiy")
	fmt.Println(x)
	fmt.Println(y)
}

//func (r receiver) identifier(parameters) (return(s)) {...}

//We define our func with parameters

//We call and pass our func in arguments

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Printf("Hello, %v from bar\n", s)
}

//Return
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

//Multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
