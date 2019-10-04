// ex07, Assign a func to a variable, then call that func
package main

import "fmt"

func main() {
	c := func() int {
		return 99
	}
	d := c()
	fmt.Println(d)
}
