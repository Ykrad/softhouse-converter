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

func parsePerson(splitLines [][]string) Person {
	person := Person{}

	for i := 0; i < len(splitLines); i++ {
		splitLine := splitLines[i]
		if splitLine[0] == "P" {
			person.firstName = splitLine[1]
			person.lastName = splitLine[2]
		} else if splitLine[0] == "T" {
			person.phone = parsePhone(splitLine)
			person.phoneInitialized = true
		} else if splitLine[0] == "A" {
			person.address = parseAddress(splitLine)
			person.addressInitialized = true
		} else if splitLine[0] == "F" {
			familyLines, stoppedAtIndex := getUnitLines(i, splitLines, familyValidation)
			person.family = append(person.family, parseFamily(familyLines))
			i = stoppedAtIndex
		}
	}

	return person
}

func personValidation(identifier string) bool {
	return identifier == "P"
}
