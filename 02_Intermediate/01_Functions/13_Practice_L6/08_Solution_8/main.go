package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1())
	fmt.Printf("%T\n", s1)

}

func foo() func() string {
	return func() string {
		return "Hello world"
	}
}
