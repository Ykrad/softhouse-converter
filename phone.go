package main

type Phone struct {
	Mobile   string `xml:"mobile,omitempty"`
	Landline string `xml:"landline,omitempty"`
}

func parsePhone(splitLine []string) Phone {
	phone := Phone{
		Mobile:   splitLine[1],
		Landline: splitLine[2],
	}
	return phone
}
