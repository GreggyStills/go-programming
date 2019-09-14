package main

import "fmt"

func main() {
	var favSport string = "Anbo-Jyutsu"
	switch favSport {
	case "baseball":
		fmt.Println("baseball's not your fav? oh well")
	case "Anbo-Jyutsu":
		fmt.Printf("Anbo-Jyutsu?  Really?  Star Trek nerd alert!")
	}
	fmt.Println()
}
