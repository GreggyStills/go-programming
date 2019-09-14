package main

import "fmt"

func main() {
	// for [<init; condition; post>] {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	// this will crash, because var i is scoped to the for loop
	// fmt.Println(i)

	// this is an infinite loop
	// for {
	//    fmt.Println("hi")
	// }

	// switch statement
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}

	n := 4
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 4:
		fmt.Println("four")
	}

}
