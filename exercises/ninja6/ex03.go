package main

import "fmt"

func main() {
	fmt.Println("Should print bar, then foo")
	defer fmt.Println(foo())
	fmt.Println(bar())
}

func foo() string {
	return "foo string"
}

func bar() (int, string) {
	return 9, "bar string"
}
