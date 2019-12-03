/*
Fix the race condition you created in exercise #3 by using package atomic
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	iters := 101
	wg.Add(iters)

	for i := 0; i < iters; i++ {
		go increment()
	}
	wg.Wait()
}

func increment() {
	atomic.AddInt64(&counter, 1)
	if atomic.LoadInt64(&counter)%5 == 0 {
		fmt.Println()
	}
	fmt.Printf("%d (gr: %d)\t", atomic.LoadInt64(&counter), runtime.NumGoroutine())
	wg.Done()
}
