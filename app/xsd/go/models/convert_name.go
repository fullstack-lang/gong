package models

import (
	"strings"
	"unicode"
)

var goKeywords = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"defer":       true,
	"else":        true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

func isGoKeywordOrPredeclared(lower string) bool {
	if goKeywords[lower] {
		return true
	}
	switch lower {
	case "string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"float32", "float64", "complex64", "complex128",
		"bool", "byte", "rune", "error", "any", "comparable",
		"true", "false", "iota", "nil",
		"append", "cap", "clear", "close", "complex", "copy", "delete",
		"imag", "len", "make", "max", "min", "new", "panic", "print", "println", "real", "recover",
		"stage", "reference", "instance", "stager":
		return true
	}
	return false
}

func xsdNameToGoIdentifier(s string) string {
	if len(s) == 0 {
		return "Empty"
	}

	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	cleaned := b.String()
	if len(cleaned) == 0 {
		return "Empty"
	}

	runes := []rune(cleaned)
	if unicode.IsDigit(runes[0]) || runes[0] == '_' {
		cleaned = "N_" + cleaned
		runes = []rune(cleaned)
	}

	res := string(unicode.ToUpper(runes[0])) + string(runes[1:])
	if isGoKeywordOrPredeclared(strings.ToLower(res)) {
		res = res + "_XSD"
	}
	return res
}
