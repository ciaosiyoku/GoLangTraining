package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type SecretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := SecretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Adeolu",
		last:  "Osiyoku",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println("")

	fmt.Printf("First Name: %v\t Last Name: %v\t Age: %v\t ltk? %v\n", sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Printf("First Name: %v\t Last Name: %v\t Age: %v\n", p2.first, p2.last, p2.age)
}
