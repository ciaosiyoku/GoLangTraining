package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println("The first name is", p1)

	changeMe(&p1)

	fmt.Println("The re-assigned name is", p1)
}

func changeMe(p *person) {
	p.name = "Adeh Osiy"
	fmt.Println("")
}
