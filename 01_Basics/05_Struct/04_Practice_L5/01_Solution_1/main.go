package main

import (
	"fmt"
)

func main() {
	type person struct {
		first      string
		last       string
		favFlavors []string
	}

	p1 := person{
		first: "Adeh",
		last:  "Osiy",
		favFlavors: []string{
			"Strawberry",
			"Martini",
			"Pepsi",
			"Sprite",
		},
	}

	p2 := person{
		first: "Bunmi",
		last:  "Oguns",
		favFlavors: []string{
			"Vanilla",
			"CreamBerry",
			"Coke",
			"Chocolate",
		},
	}

	fmt.Printf("Name: %v\t %v\n", p1.first, p1.last)
	for a, v1 := range p1.favFlavors {
		fmt.Println("Favorite flavors:", a, v1)
	}

	fmt.Println("")

	fmt.Printf("Name: %v\t %v\n", p2.first, p2.last)
	for b, v2 := range p2.favFlavors {
		fmt.Println("Favorite flavors:", b, v2)
	}

}
