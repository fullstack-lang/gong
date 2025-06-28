package models

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SanitizeFileName generates a safe filename from a given string.
//
// It performs the following actions:
// - Replaces illegal characters with a specified separator.
// - Replaces whitespace with the separator.
// - Transliterates non-ASCII characters to their ASCII equivalents.
// - Truncates the filename to a maximum length.
// - Avoids reserved filenames on Windows.
func SanitizeFileName(name string, separator string) string {
	// 1. Transliterate non-ASCII characters to ASCII
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	sanitizedName, _, _ := transform.String(t, name)

	// 2. Define a regular expression for illegal characters.
	// These include control characters, and characters not allowed in filenames
	// on Windows, macOS, and Linux.
	illegalChars := regexp.MustCompile(`[\p{Cc}\/\\:*?"<>|]`)
	sanitizedName = illegalChars.ReplaceAllString(sanitizedName, separator)

	// 3. Replace whitespace with the separator.
	whitespace := regexp.MustCompile(`\s+`)
	sanitizedName = whitespace.ReplaceAllString(sanitizedName, separator)

	// 4. Replace multiple separators with a single one.
	multiSeparator := regexp.MustCompile(regexp.QuoteMeta(separator) + "{2,}")
	sanitizedName = multiSeparator.ReplaceAllString(sanitizedName, separator)

	// 5. Trim leading and trailing separators.
	sanitizedName = strings.Trim(sanitizedName, separator)

	// 6. Truncate to a maximum length (e.g., 255 bytes, common limit).
	if len(sanitizedName) > 255 {
		sanitizedName = sanitizedName[:255]
		// Ensure we don't end with a partial multi-byte character
		for !unicode.IsLetter(rune(sanitizedName[len(sanitizedName)-1])) && !unicode.IsNumber(rune(sanitizedName[len(sanitizedName)-1])) {
			sanitizedName = sanitizedName[:len(sanitizedName)-1]
		}
	}

	// 7. Check for reserved filenames on Windows.
	reservedNames := []string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}
	for _, reserved := range reservedNames {
		if strings.ToUpper(sanitizedName) == reserved {
			sanitizedName = "_" + sanitizedName
			break
		}
	}

	return sanitizedName
}
