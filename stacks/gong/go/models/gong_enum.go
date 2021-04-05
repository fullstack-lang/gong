package models

// GongEnum is a go const declared as enum that is selected by the gongc compiler
type GongEnum struct {
	Name           string
	GongEnumValues []*GongEnumValue
}
