/*
Starting with this code ( https://play.golang.org/p/BVRZTdlUZ_ )
sort the []user by
- age
- last

Also sort each []string “Sayings” for each user
- print everything in a way that is pleasant
*/
package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type UsersByAge []user

func (u UsersByAge) Len() int {
	return len(u)
}
func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type UsersByLast []user

func (u UsersByLast) Len() int {
	return len(u)
}
func (u UsersByLast) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}
func (u UsersByLast) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user{u1, u2, u3}
	fmt.Println("==== Unsorted (original) ====")
	PrintJson(users)

	// your code goes here
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
	fmt.Println("\n\n==== Sorted (by Age) ====\n\n")
	sort.Sort(UsersByAge(users))
	PrintJson(users)

	fmt.Println("\n\n==== Sorted (by Last) ====\n\n")
	sort.Sort(UsersByLast(users))
	PrintJson(users)
}

func PrintJson(i interface{}) {
	bs, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("### JSON data follows:")
	fmt.Println(string(bs))
}
