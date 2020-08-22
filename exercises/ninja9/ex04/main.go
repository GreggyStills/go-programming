/*
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	for i := 0; i < 101; i++ {
		go increment()

	}
}

func increment() {
	mu.Lock()
	defer mu.Unlock()

	if counter%5 == 0 {
		fmt.Println()
	} else {
		fmt.Printf("\t")
	}
	counter++
	fmt.Printf("%d (gr: %d)", counter, runtime.NumGoroutine())
}
