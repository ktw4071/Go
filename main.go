package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1
	// alex := person{"Taewoo", "Kim"}
	// 2
	// alex := person{firstName: "Taewoo", lastName: "Kim"} // Same as above, but explicitly assign values, not relying on order in struct
	// fmt.Println(alex)
	// 3
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)       // {} zerovalue is assigned as "" for string
	fmt.Printf("%+v", alex) // {firstName: lastName:} %+v fills field name and value
}
