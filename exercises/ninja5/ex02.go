package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "Bob",
		last:       "Jones",
		favFlavors: []string{"chocolate", "raspberry"},
	}
	p2 := person{
		first:      "Amy",
		last:       "Smith",
		favFlavors: []string{"vanilla", "rocky road"},
	}
	people := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	printPerson(people["Jones"])
	printPerson(people["Smith"])
}

func printPerson(p person) {
	fmt.Printf("%s %s favorite ice cream flavors:\n", p.first, p.last)
	for i, f := range p.favFlavors {
		fmt.Printf("\t%d. %s\n", i+1, f)
	}
}
