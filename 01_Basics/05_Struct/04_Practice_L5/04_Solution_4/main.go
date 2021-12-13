package main

import "fmt"

func main() {
	x := struct {
		first string
		last  string
		age   int
		car   bool
	}{
		first: "Adeh",
		last:  "Osiy",
		age:   27,
		car:   true,
	}

	fmt.Printf(" Name: %v %v\n Age: %v\n Own a car?: %v\n", x.first, x.last, x.age, x.car)
}
