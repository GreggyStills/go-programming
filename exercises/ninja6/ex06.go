// ex06, Build and use an anonymous func
package main

import "fmt"

func main() {
	c := func() int {
		return 99
	}()
	fmt.Println(c)
}
