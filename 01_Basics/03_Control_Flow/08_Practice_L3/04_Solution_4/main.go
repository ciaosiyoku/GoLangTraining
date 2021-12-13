package main

import "fmt"

func main() {
	a := 1986
	for {
		if a > 2021 {
			break
		}
		fmt.Println(a)
		a++
	}
}
