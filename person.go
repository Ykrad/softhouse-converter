package main

import "fmt"

type Person struct {
	firstName          string
	lastName           string
	phone              Phone
	phoneInitialized   bool
	address            Address
	addressInitialized bool
	family             []Family
}

func (person Person) toXML() string {
	phone := ""
	if person.phoneInitialized {
		phone = person.phone.toXML()
	}

	address := ""
	if person.addressInitialized {
		address = person.address.toXML()
	}

	family := ""
	for _, familyMember := range person.family {
		family += familyMember.toXML()
	}

	return fmt.Sprintf("<person><firstname>%s</firstname><lastname>%s</lastname>%s%s%s</person>",
		person.firstName, person.lastName, phone, address, family)
}
