/*
This exercise will reinforce our understanding of method sets:
- create a type person struct
- attach a method speak to type person using a pointer receiver (*person)
- create a type human interface
	to implicitly implement the interface, a human must have the speak method
- create func “saySomething”
	- have it take in a human as a parameter
	- have it call the speak method
- show the following in your code
	- you CAN pass a value of type *person into saySomething
	- you CANNOT pass a value of type person into saySomething

here is a hint if you need some help
https://play.golang.org/p/FAwcQbNtMG

Receivers       Values
------------------------------
(t T)           T and *T
(t *T)          *T
*/

package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p *Person) Speak(s string) error {
	fmt.Printf("%s %s, age %d, says \"%s\"\n", p.First, p.Last, p.Age, s)
	return nil
}

type Human interface {
	Speak(string) error
}

func SaySomething(h Human, s string) {
	h.Speak(s)
}

func main() {
	p1 := Person{
		First: "Bob",
		Last:  "Smith",
		Age:   46,
	}
	SaySomething(&p1, "Egad, I'm ancient!")
	//SaySomething(p1, "Egad, I'm ancient!")
}
