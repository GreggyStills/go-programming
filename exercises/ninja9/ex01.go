/*
In addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exits
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go doit("foo")
	go doit("bar")
	//doit("foo")
	//doit("bar")
	for x := 0; x < 5; x++ {
		fmt.Printf("main_%02d ", x)
	}
	fmt.Println()

	wg.Wait()
}

func doit(s string) {
	wg.Add(1)
	defer wg.Done()

	fmt.Printf("=== %s ===\n", s)
	for x := 0; x < 10; x++ {
		fmt.Printf("%s_%02d ", s, x)
	}
	fmt.Println()
}
