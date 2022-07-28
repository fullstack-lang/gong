package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Astruct:
		fieldCoder := Astruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Date = time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.Booleanfield = false
		fieldCoder.Aenum = "3"
		fieldCoder.Aenum_2 = "4"
		fieldCoder.Benum = "5"
		fieldCoder.CEnum = 6
		fieldCoder.CName = "7"
		fieldCoder.CFloatfield = 8.000000
		fieldCoder.Floatfield = 10.000000
		fieldCoder.Intfield = 11
		fieldCoder.Anotherbooleanfield = true
		fieldCoder.Duration1 = 13
		return (any)(fieldCoder).(Type)
	case AstructBstruct2Use:
		fieldCoder := AstructBstruct2Use{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case AstructBstructUse:
		fieldCoder := AstructBstructUse{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
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

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *Astruct:
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
			if field == 6 {
				return "CEnum"
			}
			if field == 11 {
				return "Intfield"
			}
			if field == 13 {
				return "Duration1"
			}
		case float64:
			// insertion point for field dependant name
			if field == 8.000000 {
				return "CFloatfield"
			}
			if field == 10.000000 {
				return "Floatfield"
			}
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "Date"
			}
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "Booleanfield"
			}
			if field == true {
				return "Anotherbooleanfield"
			}
		}
	case *AstructBstruct2Use:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *AstructBstructUse:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Bstruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 3 {
				return "Intfield"
			}
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "Floatfield"
			}
			if field == 2.000000 {
				return "Floatfield2"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Dstruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
