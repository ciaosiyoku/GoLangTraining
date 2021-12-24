package main

import "fmt"

func main() {
	fi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(fi...)
	fmt.Println("Func foo prints", n)

	fmt.Println("")

	bi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := bar(bi)
	fmt.Println("Func bar prints", m)

}

func foo(xi ...int) int {
	tf := 0
	for _, v := range xi {
		tf += v
	}
	return tf
}

func bar(vi []int) int {
	tb := 0
	for _, v := range vi {
		tb += v
	}
	return tb
}
