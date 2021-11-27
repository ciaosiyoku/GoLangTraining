package main

import "fmt"

func main() {
	s := "Hello Adeh"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println("")

	// Show "Hello Adeh" in Decimal base 10
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println("")

	// Show "Hello Adeh" in UTF8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	fmt.Println("")

	// Show "Hello Adeh" in Hexa-decimal
	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
	}
}
