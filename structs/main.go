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
	// Can instantiate in three ways
	// 1.
	// ben := person{"Ben", "Cayton"} <- But this is a shitty way of doing things
	// 2.
	// Better to call out the properties explicitly upon instantiation
	// ben := person{
	// 	firstName: "Ben",
	// 	lastName:  "Cayton",
	// }
	// 3.

	var ben person
	//fmt.Println(ben)
	// %+v gives property names and values in the formatted print stmt

	ben.firstName = "Benjammin'"
	ben.lastName = "Cayton"
	ben.contactInfo.email = "benjamin.p.cayton@usace.army.mil"
	ben.contactInfo.zipCode = 80808
	ben.print()
	ben.updateName("Brad")
	ben.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Discussion on pointers
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
