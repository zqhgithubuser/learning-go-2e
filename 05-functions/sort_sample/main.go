package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people) // [{Pat Patterson 37} {Tracy Bobdaughter 23} {Fred Fredson 18}]
	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people) // [{Tracy Bobdaughter 23} {Fred Fredson 18} {Pat Patterson 37}]
	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // [{Tracy Bobdaughter 23} {Fred Fredson 18} {Pat Patterson 37}]
}
