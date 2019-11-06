package main

import (
"fmt"
"runtime"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Bob",
		last:  "Smith",
	}
	p2 := person{
		first: "Jane",
		last:  "Doe",
	}
	p1.Greet()
	p2.Greet()

	fmt.Println("# CPUs:", runtime.NumCPU())
}

func (p person) Greet() {
	fmt.Println("Hi there, my name is", p.first, p.last)
}
