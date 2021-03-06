package main

import (
	"fmt"
)

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
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	jim.updateName("Yufen")
	jim.print()
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
