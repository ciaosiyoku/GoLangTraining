package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr1 := truck{
		vehicle: vehicle{
			doors: "4 doors",
			color: "dark grey",
		},
		fourWheel: true,
	}

	se1 := sedan{
		vehicle: vehicle{
			doors: "2 doors",
			color: "light pink",
		},
		luxury: false,
	}

	fmt.Println(tr1)
	fmt.Println(se1)

	fmt.Println("")

	fmt.Println("truck color:", tr1.color)
	fmt.Println("sedan color:", se1.color)

}
