package main

import "fmt"

func main() {
	numberSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	numberSlice = append(numberSlice, 52)
	fmt.Println(numberSlice)
	numberSlice = append(numberSlice, 53, 54, 55)
	fmt.Println(numberSlice)
	y := []int{56, 57, 58, 59, 60}
	numberSlice = append(numberSlice, y...)
	fmt.Println(numberSlice)
}
