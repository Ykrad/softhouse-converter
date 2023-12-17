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

func parsePhone(splitLine []string) Phone {
	phone := Phone{
		mobile:   splitLine[1],
		landline: splitLine[2],
	}
	return phone
}
