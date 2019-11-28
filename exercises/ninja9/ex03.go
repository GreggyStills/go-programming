/*
Using goroutines, create an incrementer program
- have a variable to hold the incrementer value
- launch a bunch of goroutines

each goroutine should:
- read the incrementer value
- store it in a new variable
- yield the processor with runtime.Gosched()
- increment the new variable
- write the value in the new variable back to the incrementer variable
- use waitgroups to wait for all of your goroutines to finish

The above will create a race condition.
Prove that it is a race condition by using the -race flag
if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	my_way()
	//todds_way()
}

func my_way() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	iters := 101
	wg.Add(iters)

	myInt := 0
	for i := 0; i < iters; i++ {
		if i%5 == 0 {
			fmt.Println()
		}
		go increment(&myInt)
		fmt.Printf("%d(gr: %d)\t", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println()
}

func increment(i *int) {
	runtime.Gosched() // yield the processor
	*i += 1
	wg.Done()
}

func todds_way() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg2 sync.WaitGroup
	wg2.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg2.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg2.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
