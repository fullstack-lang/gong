package generator

import (
	"strings"
	"unicode"
)

// GenerateGoIdentifier takes a single-line string and converts it to a Go identifier
// with the first letter capitalized. It removes or replaces characters that are
// not valid in Go identifiers.
func GenerateGoIdentifier(input string) string {
	if len(input) == 0 {
		return ""
	}

	var sb strings.Builder
	capitalizeNext := true

	for _, r := range input {
		if unicode.IsSpace(r) || !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			sb.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			sb.WriteRune(r)
		}
	}

	identifier := sb.String()

	// Ensure the first character is uppercase if the input wasn't empty
	if len(identifier) > 0 && unicode.IsLower(rune(identifier[0])) {
		first := string(unicode.ToUpper(rune(identifier[0])))
		rest := ""
		if len(identifier) > 1 {
			rest = identifier[1:]
		}
		return first + rest
	}

	return identifier
}
