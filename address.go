package main

type Address struct {
	Street   string `xml:"street"`
	City     string `xml:"city"`
	AreaCode string `xml:"areaCode,omitempty"`
}

func parseAddress(splitLine []string) Address {
	address := Address{
		Street: splitLine[1],
		City:   splitLine[2],
	}

	if len(splitLine) == 4 {
		address.AreaCode = splitLine[3]
	}

	return address
}
