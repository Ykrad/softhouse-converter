package main

import "fmt"

type Phone struct {
	mobile   string
	landline string
}

func (phone Phone) toXML() string {
	return fmt.Sprintf("<phone><mobile>%s</mobile><landline>%s</landline></phone>",
		phone.mobile, phone.landline)
}

type Address struct {
	street   string
	city     string
	areaCode string
}

func (address Address) toXML() string {
	areaCode := ""
	if address.areaCode != "" {
		areaCode = fmt.Sprintf("<areaCode>%s</areaCode>", address.areaCode)
	}

	return fmt.Sprintf("<address><street>%s</street><city>%s</city>%s</address>",
		address.street, address.city, areaCode)
}

type Family struct {
	name               string
	birthyear          string
	phone              Phone
	phoneInitialized   bool
	address            Address
	addressInitialized bool
}

func (family Family) toXML() string {
	address := ""
	if family.addressInitialized {
		address = family.address.toXML()
	}

	phone := ""
	if family.phoneInitialized {
		phone = family.phone.toXML()
	}

	return fmt.Sprintf("<family><name>%s</name><born>%s</born>%s%s</family>",
		family.name, family.birthyear, address, phone)
}

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
