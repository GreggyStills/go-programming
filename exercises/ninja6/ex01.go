package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() string {
	return "foo string"
}

func bar() (int, string) {
	return 9, "bar string"
}
