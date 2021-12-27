package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{
		First: "James",
		Age:   32,
	}
	p2 := Person{
		First: "Moneypenny",
		Age:   27,
	}
	p3 := Person{
		First: "Q",
		Age:   64,
	}
	p4 := Person{
		First: "M",
		Age:   56,
	}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Unsorted strings")
	fmt.Println(people)

	fmt.Println("\nSorted by Age")
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println("\nSorted by Name")
	sort.Sort(ByName(people))
	fmt.Println(people)
}
