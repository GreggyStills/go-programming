package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	fmt.Println("==== Test marshal ====")
	testMarshal()

	fmt.Println("==== Test unmarshal ====")
	testUnMarshal()
}

func testMarshal() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	PrintJson(p1)
	fmt.Println()
	PrintJson(people)
}

func testUnMarshal() {
	var op string = `
[
  {
    "First": "Bob",
    "Last": "Smith",
    "Age": 32
  },
  {
    "First": "Jane",
    "Last": "Smith",
    "Age": 27
  }
]
`
	var otherPeople []person

	bs := []byte(op)
	err := json.Unmarshal(bs, &otherPeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(otherPeople)
	PrintJson(otherPeople)

}

func PrintJson(i interface{}) {
	bs, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("### JSON data follows:")
	fmt.Println(string(bs))
}
