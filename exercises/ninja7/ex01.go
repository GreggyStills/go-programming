// ex01, pointers
// Create a value and assign it to a variable.
// Print the address of that value.

package main

import "fmt"

func main() {
	a := 99
	fmt.Printf("'a'\tmemory addr: %v\tvalue: %v\ttype:%T\n", &a, a, a)
}
