package main

import "fmt"

func main() {
	ns := makeNumbers(1, 100)
	s := sum(ns...)
	fmt.Println("Total: ", s)
}

// sum takes any number of integers as input, adds them all together, and returns the total
func sum(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

// makeNumbers returns a slice of consecutive integers, from a-z
func makeNumbers(a, z int) (numbers []int) {
	for n := a; n <= z; n++ {
		numbers = append(numbers, n)
	}
	return
}
