package main

import "fmt"

func main() {
	age := 40

	for age > 30 {
		fmt.Printf("My age is: %d\n", age)
		break
	}
	fmt.Println()
}
