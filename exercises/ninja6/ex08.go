// ex08
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
package main

import "fmt"

func main() {
	c := func() func() int {
		return func() int {
			return 99
		}
	}
	d := c()
	fmt.Println(d())
}
