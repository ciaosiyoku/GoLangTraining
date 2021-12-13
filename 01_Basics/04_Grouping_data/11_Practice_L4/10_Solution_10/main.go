package main

import "fmt"

func main() {
	m := map[string][]string{
		"Bond_Jame":       []string{"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny_Miss": []string{"James Bond", "Literature", "Computer Science"},
		"No_Dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	// Adding a new record
	m["Adeh_Mr"] = []string{"Resilient", "Chocolate", "Engineering"}

	//Deleting a record
	delete(m, "No_Dr")

	fmt.Println("")

	for a, v := range m {
		fmt.Println("This is the record for", a)
		for b, v2 := range v {
			fmt.Println("\t", b, v2)
		}

	}
}
