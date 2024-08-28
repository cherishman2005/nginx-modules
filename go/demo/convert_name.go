package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// convert converts name from CamelCase to UnderScoreCase
func convert(name string) string {
	var b bytes.Buffer
	for i, c := range name {
		if unicode.IsUpper(c) {
			if i > 0 {
				b.WriteString("_")
			}
			b.WriteRune(c)
		} else {
			b.WriteRune(unicode.ToUpper(c))
		}
	}
	return b.String()
}
func main() {
	s := convert("ClientConnActive")

	fmt.Println("s=", s)
}
