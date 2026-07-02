package models

type GongEnumType int

const (
	Int GongEnumType = iota
	String
)

// GongEnum is a go const declared as enum that is selected by the gong generate compiler
type GongEnum struct {
	Name           string
	Type           GongEnumType
	GongEnumValues []*GongEnumValue
}
