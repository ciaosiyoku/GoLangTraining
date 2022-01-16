package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Adeh",
		last:  "Osiy",
		age:   24,
	}

	p2 := person{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	saySomething(&p1)
	fmt.Println("-")
	p2.speak()
}
