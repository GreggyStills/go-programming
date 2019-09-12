package main

import "fmt"

const (
    y = 2015
    y1 = y + iota
    y2 = y + iota
    y3 = y + iota
    y4 = y + iota
)

func main() {
    fmt.Println(y1,y2,y3,y4)
}

