package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println("")

	fmt.Println("James is", m["James"])
	fmt.Println("Moneypenny is", m["Moneypenny"])

	fmt.Println("")

	//To check if values is stored in the map
	x, ok := m["James"]
	fmt.Println("James is", x)
	fmt.Println("Data stored?", ok)
	fmt.Println("")

	y, ok := m["Moneypenny"]
	fmt.Println("Moneypenny is", y)
	fmt.Println("Data stored?", ok)
	fmt.Println("")

	z, ok := m["Adeh"]
	fmt.Println("Adeh is", z)
	fmt.Println("Data stored?", ok)
	fmt.Println("")

	//Using If condition
	if a, ok := m["James"]; ok {
		fmt.Println("This is James Bond and he is", a)
	}

	if b, ok := m["Moneypenny"]; ok {
		fmt.Println("This is Miss Moneypenny and she is", b)
	}

	if c, ok := m["Adeh"]; ok {
		fmt.Println("This is Adeh and he is", c)
	}

	fmt.Println("")

	// To add additional element
	m["Adeh"] = 33

	//To range over the elements
	for a, v := range m {
		fmt.Println(a, v)
	}

	fmt.Println("")

	// To delete an entry from the map
	delete(m, "Moneypenny")
	fmt.Println("After deleting Moneypenny, we have")

	for b, v := range m {
		fmt.Println(b, v)
	}
}
