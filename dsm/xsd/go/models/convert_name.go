package models

import (
	"strings"
	"unicode"
)

func xsdNameToGoIdentifier(s string) string {
	if len(s) == 0 {
		return s
	}

	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, " ", "_")

	// Convert the first rune to uppercase and append the rest of the string
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
