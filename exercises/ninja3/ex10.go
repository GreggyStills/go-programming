package main

import "fmt"

func main() {
	t := true
	f := false
	fmt.Printf("%s && %s = %s\n", t, t, t && t)
	fmt.Printf("%s && %s = %s\n", t, f, t && f)
	fmt.Printf("%s || %s = %s\n", t, t, t || t)
	fmt.Printf("!%s = %s\n", t, !t)
}
