package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Printf("%d mod 4 = %d\n", i, i%4)
	}
	fmt.Println()
}
