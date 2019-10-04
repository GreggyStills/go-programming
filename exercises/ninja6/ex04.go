package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi my name is %s, and I am %d years old.\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Bob",
		last:  "Smith",
		age:   40,
	}
	p1.speak()

}
