package main

import "fmt"

func main() {
    x := 42
    y := x << 1
    fmt.Printf("decimal: %d\nbinary: %b\nhex: %#x\n", x,x,x)
    fmt.Printf("decimal: %d\nbinary: %b\nhex: %#x\n", y,y,y)
}
