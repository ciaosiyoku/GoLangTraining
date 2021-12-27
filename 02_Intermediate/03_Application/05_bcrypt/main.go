package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := `password12345`
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pwd)
	fmt.Println(bs)

	fmt.Println("")

	loginpwd1 := `password12345`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpwd1))
	if err != nil {
		fmt.Println("Wrong login details, check and try again!!!")
		return
	}

	fmt.Println("You are logged in")
}
