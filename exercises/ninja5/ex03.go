package main

import "fmt"

type vehicle struct {
	doors int
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
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Printf("%T: %v\n", t1, t1)
	fmt.Printf("%T: %v\n", s1, s1)
	fmt.Printf("%s\n%s\n", t1.color, s1.color)

}
