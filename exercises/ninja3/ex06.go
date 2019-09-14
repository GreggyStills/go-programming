package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%d is divisible by 4!\n", i)
		}
	}
	fmt.Println()
}
