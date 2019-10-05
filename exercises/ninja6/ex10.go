// ex10
// Closure is when we have “enclosed” the scope of a variable in some code block.
// create a func which “encloses” the scope of a variable

package main

import "fmt"

func main() {
	a := increment(0)
	b := increment(45)
	for n := 1; n <= 10; n++ {
		fmt.Println("a:", a(), "\tb:", b())
	}
}

func increment(i int) func() int {
	var x int = i
	return func() int {
		x++
		return x
	}
}
