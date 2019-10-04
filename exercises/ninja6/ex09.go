// ex09
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument

package main

import "fmt"

func n99() int {
	return 99
}

func multiply(f func() int, m int) int {
	return f() * m
}

func main() {
	fmt.Println(multiply(n99, 1))
	fmt.Println(multiply(n99, 2))
}
