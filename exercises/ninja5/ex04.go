package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := struct {
		first      string
		last       string
		favFlavors []string
	}{
		first:      "Bob",
		last:       "Jones",
		favFlavors: []string{"chocolate", "raspberry"},
	}
	fmt.Printf("%T\n", p1)
	fmt.Printf("%v\n", p1)
}
