package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	fmt.Println(mp)

	fmt.Println("")

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	fmt.Println("")

	for i, xs := range xp {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t Index position: %v\t Value is: %v\n", j, val)
		}
	}
}
