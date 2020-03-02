package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

// turn "address" into "value" with "*address"
// turn "value" into "address" with "&value"

func (pointerToPerson *person) updateName(newFirstName string) { // give me the valeue  this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName // Pass by value, it won't update jim.firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
