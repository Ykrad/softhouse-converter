package main

import (
	"encoding/xml"
)

type Person struct {
	FirstName          string
	LastName           string
	Phone              Phone
	phoneInitialized   bool
	Address            Address
	addressInitialized bool
	Family             []Family
}

func (person Person) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	personStart := xml.StartElement{Name: xml.Name{Local: "person"}}
	if error := encoder.EncodeToken(personStart); error != nil {
		return error
	}

	if person.FirstName != "" {
		if error := encoder.EncodeElement(person.FirstName, xml.StartElement{Name: xml.Name{Local: "firstname"}}); error != nil {
			return error
		}
	}

	if person.LastName != "" {
		if error := encoder.EncodeElement(person.LastName, xml.StartElement{Name: xml.Name{Local: "lastname"}}); error != nil {
			return error
		}
	}

	if person.phoneInitialized {
		if error := encoder.EncodeElement(person.Phone, xml.StartElement{Name: xml.Name{Local: "phone"}}); error != nil {
			return error
		}
	}

	if person.addressInitialized {
		if error := encoder.EncodeElement(person.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); error != nil {
			return error
		}
	}

	if len(person.Family) > 0 {
		if error := encoder.EncodeElement(person.Family, xml.StartElement{Name: xml.Name{Local: "family"}}); error != nil {
			return error
		}
	}

	if error := encoder.EncodeToken(personStart.End()); error != nil {
		return error
	}

	return nil
}

func parsePerson(splitLines [][]string) Person {
	person := Person{}

	for i := 0; i < len(splitLines); i++ {
		splitLine := splitLines[i]
		if splitLine[0] == "P" {
			person.FirstName = splitLine[1]
			person.LastName = splitLine[2]
		} else if splitLine[0] == "T" {
			person.Phone = parsePhone(splitLine)
			person.phoneInitialized = true
		} else if splitLine[0] == "A" {
			person.Address = parseAddress(splitLine)
			person.addressInitialized = true
		} else if splitLine[0] == "F" {
			familyLines, stoppedAtIndex := getUnitLines(i, splitLines, familyValidation)
			person.Family = append(person.Family, parseFamily(familyLines))
			i = stoppedAtIndex
		}
	}

	return person
}

func personValidation(identifier string) bool {
	return identifier == "P"
}
