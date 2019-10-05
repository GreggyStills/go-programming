// ex02, pointers
/*
create a person struct
create a func called “changeMe” with a *person as a parameter
change the value stored at the *person address

Tips:
 to dereference a struct, use (*value).field
 p1.first
 (*p1).first
 “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
 https://golang.org/ref/spec#Selectors

create a value of type person
print out the value
in func main
  call “changeMe”
  print out the value
*/

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Changed"
}

func main() {
	p1 := person{
		first: "Bob",
		last:  "Smith",
		age:   90,
	}
	fmt.Println("Before:", p1)
	changeMe(&p1)
	fmt.Println("After:", p1)
}
