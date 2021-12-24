package main

import "fmt"

type person struct {
	first string
	last  string
	Age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

//func (r receiver) identifier(parameters) (return(s)) {...}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			Age:   37,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Adeh",
			last:  "Osiy",
			Age:   25,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println("")

	fmt.Println(sa2)
	sa2.speak()
}
