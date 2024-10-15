package main

import (
	"fmt"

	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

func main() {
	usPresidents := []struct {
		firstName  string
		middleName string
		lastName   string
		termStart  int
		termEnd    int
	}{
		{"George", "", "Washington", 1789, 1797},
		{"John", "", "Adams", 1797, 1801},
		{"Thomas", "", "Jefferson", 1801, 1809},
		{"James", "", "Madison", 1809, 1817},
		{"James", "", "Monroe", 1817, 1825},
		{"John", "Quincy", "Adams", 1825, 1829},
		{"John", "", "Tyler", 1841, 1845},
		{"James", "", "Polk", 1845, 1849},
		{"Grover", "", "Cleveland", 1885, 1889},
		{"Benjamin", "", "Harrison", 1889, 1893},
		{"Grover", "", "Cleveland", 1893, 1897},
		{"George", "Herbert Walker", "Bush", 1989, 1993},
		{"George", "Walker", "Bush", 2001, 2009},
		{"Barack", "Hussein", "Obama", 2009, 2017},
	}

	m := slicemultimap.New()

	for _, president := range usPresidents {
		m.Put(president.firstName, president.lastName)
	}

	for _, firstName := range m.KeySet() {
		lastNames, _ := m.Get(firstName)
		fmt.Printf("%v: %v\n", firstName, lastNames)
	}
}
