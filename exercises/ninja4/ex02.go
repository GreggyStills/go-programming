package main

import "fmt"

func main() {
	var numberSlice []int
	for i := 0; i < 10; i++ {
		numberSlice = append(numberSlice, i*3)
	}
	for _, i := range numberSlice {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n%#v\n", numberSlice)
}
