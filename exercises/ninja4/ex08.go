package main

import "fmt"

func main() {
	faves := map[string][]string{
		"bond_james":     []string{"martinis", "women", "killing"},
		"moneypenny_eve": []string{"flowers", "flirting", "James Bond"},
	}

	for k, v := range faves {
		fmt.Println(k, v)
	}
}
