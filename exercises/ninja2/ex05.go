package main

import "fmt"

func main() {
    var x string
    x = `
This is a raw string literal
it's got newlines and weird stuff "" %$^
`
    fmt.Println(x)
}
