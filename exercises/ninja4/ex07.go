package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	ep := []string{"Eve", "Moneypenny", "Ohhh, James..."}

	people := [][]string{jb, ep}

	for _, p := range people {
		fmt.Println(p)
	}

	fmt.Println()

	for _, p := range people {
		for _, t := range p {
			fmt.Println(t)
		}
	}

}
