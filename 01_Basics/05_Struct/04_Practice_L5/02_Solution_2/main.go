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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first, v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("")
	}

}
