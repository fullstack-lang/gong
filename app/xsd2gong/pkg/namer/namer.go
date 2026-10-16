package namer

import (
	"fmt"
	"strings"
	"unicode"
)

// ToExportedIdentifier converts an XSD name (e.g. "score-partwise", "above-below")
// into a valid exported Go identifier (e.g. "ScorePartwise", "AboveBelow").
func ToExportedIdentifier(xsdName string) string {
	if xsdName == "" {
		return "Empty"
	}
	// Strip namespace prefix if any
	if idx := strings.IndexByte(xsdName, ':'); idx >= 0 {
		xsdName = xsdName[idx+1:]
	}

	var sb strings.Builder
	capitalizeNext := true
	for _, r := range xsdName {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if sb.Len() == 0 && unicode.IsDigit(r) {
				sb.WriteString("N")
			}
			if capitalizeNext {
				sb.WriteRune(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				sb.WriteRune(r)
			}
		} else {
			capitalizeNext = true
		}
	}

	res := sb.String()
	if res == "" {
		return "Empty"
	}

	// Avoid collision with Gong's mandatory Name field or Go keywords
	switch res {
	case "Name":
		return "NameAttr"
	case "Stage", "Commit", "Unstage", "Checkout", "Reset":
		return res + "XSD"
	}

	return res
}

// ToEnumConstantName generates a valid Go constant identifier for an enum value.
func ToEnumConstantName(enumTypeName, val string) string {
	if val == "" {
		return enumTypeName + "_Empty"
	}
	// Special symbols in MusicXML or XSD enumerations
	cleaned := val
	cleaned = strings.ReplaceAll(cleaned, "+", "Plus")
	cleaned = strings.ReplaceAll(cleaned, "/", "Slash")
	cleaned = strings.ReplaceAll(cleaned, ".", "Dot")

	suffix := ToExportedIdentifier(cleaned)
	return enumTypeName + "_" + suffix
}

// Scope tracks used identifiers within a struct or package to resolve collisions deterministically.
type Scope struct {
	used map[string]int
}

func NewScope() *Scope {
	return &Scope{
		used: make(map[string]int),
	}
}

// Reserve marks an identifier as used.
func (s *Scope) Reserve(ident string) {
	if s.used[ident] == 0 {
		s.used[ident] = 1
	}
}

// Unique returns a unique exported Go identifier in this scope.
func (s *Scope) Unique(rawName string) string {
	base := ToExportedIdentifier(rawName)
	count := s.used[base]
	s.used[base] = count + 1
	if count == 0 {
		return base
	}
	res := fmt.Sprintf("%s_%d", base, count)
	s.used[res] = 1
	return res
}

