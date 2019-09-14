package main

import "fmt"

func main() {
	var numberSlice []int
	for i := 42; i < 52; i++ {
		numberSlice = append(numberSlice, i)
	}
	//fmt.Println(numberSlice)
	fmt.Println(numberSlice[:5])
	fmt.Println(numberSlice[5:])
	fmt.Println(numberSlice[2:7])
	fmt.Println(numberSlice[1:6])
}
