package main

import (
	"encoding/xml"
)

type Family struct {
	Name               string
	Birthyear          string
	Phone              Phone
	phoneInitialized   bool
	Address            Address
	addressInitialized bool
}

func (family Family) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	familyStart := xml.StartElement{Name: xml.Name{Local: "family"}}
	if error := encoder.EncodeToken(familyStart); error != nil {
		return error
	}

	if family.Name != "" {
		if error := encoder.EncodeElement(family.Name, xml.StartElement{Name: xml.Name{Local: "name"}}); error != nil {
			return error
		}
	}

	if family.Birthyear != "" {
		if error := encoder.EncodeElement(family.Birthyear, xml.StartElement{Name: xml.Name{Local: "born"}}); error != nil {
			return error
		}
	}

	if family.phoneInitialized {
		if error := encoder.EncodeElement(family.Phone, xml.StartElement{Name: xml.Name{Local: "phone"}}); error != nil {
			return error
		}
	}

	if family.addressInitialized {
		if error := encoder.EncodeElement(family.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); error != nil {
			return error
		}
	}

	if error := encoder.EncodeToken(familyStart.End()); error != nil {
		return error
	}

	return nil
}

func parseFamily(splitLines [][]string) Family {
	family := Family{}

	for _, splitLine := range splitLines {
		if splitLine[0] == "F" {
			family.Name = splitLine[1]
			family.Birthyear = splitLine[2]
		} else if splitLine[0] == "T" {
			family.Phone = parsePhone(splitLine)
			family.phoneInitialized = true
		} else if splitLine[0] == "A" {
			family.Address = parseAddress(splitLine)
			family.addressInitialized = true
		}
	}

	return family
}

func familyValidation(identifier string) bool {
	return identifier == "F" || identifier == "P"
}
