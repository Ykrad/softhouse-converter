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

func parseFamily(splitLines [][]string) Family {
	family := Family{}

	for _, splitLine := range splitLines {
		if splitLine[0] == "F" {
			family.name = splitLine[1]
			family.birthyear = splitLine[2]
		} else if splitLine[0] == "T" {
			family.phone = parsePhone(splitLine)
			family.phoneInitialized = true
		} else if splitLine[0] == "A" {
			family.address = parseAddress(splitLine)
			family.addressInitialized = true
		}
	}

	return family
}

func familyValidation(identifier string) bool {
	return identifier == "F" || identifier == "P"
}
