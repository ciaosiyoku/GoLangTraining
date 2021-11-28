package main

import "fmt"

func main() {
	g := (42 == 43)
	h := (42 <= 43)
	i := (42 >= 43)
	j := (42 != 43)
	k := (42 < 43)
	l := (42 > 43)

	fmt.Println("g =", g)
	fmt.Println("h =", h)
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)
	fmt.Println("l =", l)
}
