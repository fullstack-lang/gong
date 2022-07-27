package models

import "time"

// GongfieldCoder return an instance of Type where each field
// encodes the index of the field.
// This allows for refactorable field name
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Astruct:
		fieldCoder := Astruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Date = time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.Aenum = "3"
		fieldCoder.Aenum_2 = "4"
		fieldCoder.Benum = "5"
		fieldCoder.CEnum = 6
		fieldCoder.CName = "7"
		fieldCoder.CFloatfield = 8.000000
		fieldCoder.Bstruct = &Bstruct{Name: "9"}
		fieldCoder.Floatfield = 10.000000
		fieldCoder.Intfield = 11
		fieldCoder.Duration1 = 13
		fieldCoder.Associationtob = &Bstruct{Name: "14"}
		fieldCoder.Anotherassociationtob_2 = &Bstruct{Name: "15"}
		fieldCoder.Anarrayofb = []*Bstruct{
			{
				Name: "16",
			},
		}
		fieldCoder.Anotherarrayofb = []*Bstruct{
			{
				Name: "17",
			},
		}
		fieldCoder.Anarrayofa = []*Astruct{
			{
				Name: "18",
			},
		}
		fieldCoder.AnarrayofbUse = []*AstructBstructUse{
			{
				Name: "19",
			},
		}
		fieldCoder.Anarrayofb2Use = []*AstructBstruct2Use{
			{
				Name: "20",
			},
		}
		fieldCoder.AnAstruct = &Astruct{Name: "21"}
		return (any)(fieldCoder).(Type)
	case AstructBstruct2Use:
		fieldCoder := AstructBstruct2Use{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Bstrcut2 = &Bstruct{Name: "1"}
		return (any)(fieldCoder).(Type)
	case AstructBstructUse:
		fieldCoder := AstructBstructUse{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Bstruct2 = &Bstruct{Name: "1"}
		return (any)(fieldCoder).(Type)
	case Bstruct:
		fieldCoder := Bstruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Floatfield = 1.000000
		fieldCoder.Floatfield2 = 2.000000
		fieldCoder.Intfield = 3
		return (any)(fieldCoder).(Type)
	case Dstruct:
		fieldCoder := Dstruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Astruct | []*Astruct | *AstructBstruct2Use | []*AstructBstruct2Use | *AstructBstructUse | []*AstructBstructUse | *Bstruct | []*Bstruct | *Dstruct | []*Dstruct
}

func GongfieldName[Type Gongstruct, FieldType Gongfield, AssociationFieldType Gongstruct](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Astruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "Aenum"
			}
			if field == "4" {
				return "Aenum_2"
			}
			if field == "5" {
				return "Benum"
			}
			if field == "7" {
				return "CName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == "6" {
				return "CEnum"
			}
			if field == "11" {
				return "Intfield"
			}
			if field == "13" {
				return "Duration1"
			}
		case *AssociationFieldType:
			if (*field).GetName() == "toto" {
				return "Associationtob"
			}
		case []*AssociationFieldType:
			if (*field[0]).GetName() == "toto" {
				return "Associationtob"
			}
		}
	case AstructBstruct2Use:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		}
	case AstructBstructUse:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		}
	case Bstruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == "3" {
				return "Intfield"
			}
		}
	case Dstruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
