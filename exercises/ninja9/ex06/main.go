/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
- go run
- go build
- go install
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS:\t%s\nCPU:\t%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
}
