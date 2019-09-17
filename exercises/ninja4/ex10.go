package main

import "fmt"

func main() {
	faves := map[string][]string{
		"bond_james":     []string{"martinis", "women", "killing"},
		"moneypenny_eve": []string{"flowers", "flirting", "James Bond"},
	}
	faves["goldfinger_aurus"] = []string{"gold", "Jaws", "expecting people to die"}

	delete(faves, "moneypenny_eve")

	for person, things := range faves {
		fmt.Printf("%s:\n", person)
		for _, t := range things {
			fmt.Printf("\t%s\n", t)
		}
	}
}
