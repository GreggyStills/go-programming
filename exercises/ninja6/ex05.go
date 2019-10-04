package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

func Info(s Shape) float64 {
	return s.Area()
}

func main() {
	s := Square{
		side: 20,
	}
	c := Circle{
		radius: 10,
	}
	fmt.Println("Square ", s.side, ", area:", Info(s))
	fmt.Println("Circle ", c.radius, " area:", Info(c))
}
