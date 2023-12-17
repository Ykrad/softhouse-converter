package main

import "fmt"

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

func parseAddress(splitLine []string) Address {
	address := Address{
		street: splitLine[1],
		city:   splitLine[2],
	}

	if len(splitLine) == 4 {
		address.areaCode = splitLine[3]
	}

	return address
}
