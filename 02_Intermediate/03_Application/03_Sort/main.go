package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("Unsorted ints:", xi)
	sort.Ints(xi)
	fmt.Println("Sorted ints:", xi)

	fmt.Println("")

	fmt.Println("Unsorted strings:", xs)
	sort.Strings(xs)
	fmt.Println("Sorted strings:", xs)
}
