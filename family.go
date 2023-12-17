package main

import "fmt"

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
