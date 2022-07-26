package models

import "time"

// GongfieldCoder return an instance of Type where each field
// encodes the index of the field.
// This allows for refactorable field name
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Astruct | []*Astruct | *AstructBstruct2Use | []*AstructBstruct2Use | *AstructBstructUse | []*AstructBstructUse | *Bstruct | []*Bstruct | *Dstruct | []*Dstruct
}

func GongfieldName[Type Gongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	}
	return ""
}
