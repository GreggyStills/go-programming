package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%d is divisible by 4!\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is divisible by 3!\n", i)
		} else {
			fmt.Printf("Darn! %d is not divisible by 3 or 4!\n", i)
		}

	}
	fmt.Println()
}
