package main

import "fmt"

func main() {
	ns := makeNumbers(1, 100)
	fmt.Println("The [foo] total is: ", foo(ns...))
	fmt.Println("The [bar] total is: ", bar(ns))
}

func foo(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func bar(numbers []int) (sum int) {
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
