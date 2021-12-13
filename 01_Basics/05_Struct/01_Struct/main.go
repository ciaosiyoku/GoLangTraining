package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Adeolu",
		last:  "Osiyoku",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("")

	fmt.Printf("First Name: %v\t Last Name: %v\t Age: %v\n", p1.first, p1.last, p1.age)
	fmt.Printf("First Name: %v\t Last Name: %v\t Age: %v\n", p2.first, p2.last, p2.age)
}
