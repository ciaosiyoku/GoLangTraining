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
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

type hotdog int

func bar(h human) {
	fmt.Println("I was passed into bar", h)

	//Assertion
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar by", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar by secretAgent", h.(secretAgent).first)
	}
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

	p1 := person{
		first: "Dr.",
		last:  "Yes",
		Age:   42,
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println("")

	fmt.Println(sa2)
	sa2.speak()

	fmt.Println("")

	fmt.Println(p1)
	p1.speak()

	fmt.Println("")

	bar(sa1)
	bar(sa2)
	bar(p1)

	fmt.Println("")

	//Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
