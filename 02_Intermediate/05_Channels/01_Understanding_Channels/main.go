package main

import "fmt"

func main() {
	//first route (pass)
	c := make(chan int)

	go func() {
		c <- 18
	}()

	fmt.Println("pass =", <-c)

	fmt.Println("")

	//second route (buffer)
	d := make(chan int, 1)

	d <- 27

	fmt.Println("buffer =", <-d)

}
