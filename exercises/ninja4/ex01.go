package main

import "fmt"

func main() {
	var numbers [5]int
	for i := 0; i <= 4; i++ {
		numbers[i] = i * 3
	}
	for _, i := range numbers {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("%#v", numbers)
}
