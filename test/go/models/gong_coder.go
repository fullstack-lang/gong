package models

import "time"

// GongfieldCoder return an instance of Type where each field
// encodes the index of the field.
// This allows for refactorable field name
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	case Astruct:
		fieldCoder := Astruct{
			Cstruct: Cstruct{
				CName:       "6",
				CFloatfield: 7.0,
			},
			Name:                "1",
			Date:                time.Date(2, time.January, 0, 0, 0, 0, 0, time.UTC),
			Booleanfield:        false,
			Anotherbooleanfield: true,
			Aenum:               ENUM_VAL1,
			Floatfield:          3.0,
			Intfield:            4,
		}
		fieldCoder.CName = "6"
		fieldCoder.CFloatfield = 7.0
		fieldCoder.Associationtob = &Bstruct{Name: "10"}
		fieldCoder.Anotherassociationtob_2 = &Bstruct{Name: "12"}
		fieldCoder.Anarrayofb = []*Bstruct{
			{
				Name: "15",
			},
		}
		fieldCoder.Anotherarrayofb = []*Bstruct{
			{
				Name: "17",
			},
		}
		fieldCoder.Aenum = "-1"
		fieldCoder.CEnum = -10
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Astruct | *Bstruct | []*Astruct | []*Bstruct | AEnumType | CEnumTypeInt
}

func GongfieldName[Type Gongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	case Astruct:
		switch field := any(field).(type) {
		case string:
			if field == "1" {
				return "Name"
			}
			if field == "6" {
				return "CName"
			}
			return ""
		case int:
			if field == 4 {
				return "Intfield"
			}
			return ""
		case bool:
			// can't code more than 2 different values with a bool
			if !field {
				return "Booleanfield"
			}
			if field {
				return "Anotherbooleanfield"
			}
			return ""
		case float64:
			if field == 3.0 {
				return "Floatfield"
			}
			if field == 7.0 {
				return "CFloatfield"
			}
		case *Bstruct:
			if field.Name == "10" {
				return "Associationtob"
			}
			if field.Name == "12" {
				return "Anotherassociationtob_2"
			}
		case []*Bstruct:
			if field[0].Name == "15" {
				return "Anarrayofb"
			}
			if field[0].Name == "17" {
				return "Anarrayofb"
			}
		case AEnumType:
			if field == "-1" {
				return "Aenum"
			}
		case CEnumTypeInt:
			if field == -10 {
				return "CEnum"
			}
		}

	}
	return ""
}
