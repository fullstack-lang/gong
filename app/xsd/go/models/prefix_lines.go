package models

import "strings"

// PrefixLines prefixes each line of the input string with "// "
func PrefixLines(input string) string {
	// Split the input string into lines
	lines := strings.Split(input, "\n")

	// Prefix each line with "// "
	for i, line := range lines {
		line = strings.ReplaceAll(line, "                ", "")
		lines[i] = "// " + line
	}

	// Join the lines back into a single string
	return "\n" + strings.Join(lines, "\n")
}
