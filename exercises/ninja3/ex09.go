package main

import "fmt"

func main() {
	var s string = "favSport"
	switch s {
	case "notFav":
		fmt.Println("not your fav, oh well")
	case "favSport":
		fmt.Printf("woohoo, your fav!")
	}
	fmt.Println()
}
