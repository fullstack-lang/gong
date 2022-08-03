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
	case GongBasicField:
		fieldCoder := GongBasicField{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.BasicKindName = "1"
		fieldCoder.DeclaredType = "3"
		fieldCoder.CompositeStructName = "4"
		fieldCoder.Index = 5
		return (any)(fieldCoder).(Type)
	case GongEnum:
		fieldCoder := GongEnum{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Type = 1
		return (any)(fieldCoder).(Type)
	case GongEnumValue:
		fieldCoder := GongEnumValue{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Value = "1"
		return (any)(fieldCoder).(Type)
	case GongNote:
		fieldCoder := GongNote{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Body = "1"
		return (any)(fieldCoder).(Type)
	case GongStruct:
		fieldCoder := GongStruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case GongTimeField:
		fieldCoder := GongTimeField{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Index = 1
		fieldCoder.CompositeStructName = "2"
		return (any)(fieldCoder).(Type)
	case ModelPkg:
		fieldCoder := ModelPkg{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.PkgPath = "1"
		return (any)(fieldCoder).(Type)
	case PointerToGongStructField:
		fieldCoder := PointerToGongStructField{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Index = 2
		fieldCoder.CompositeStructName = "3"
		return (any)(fieldCoder).(Type)
	case SliceOfPointerToGongStructField:
		fieldCoder := SliceOfPointerToGongStructField{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Index = 2
		fieldCoder.CompositeStructName = "3"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *GongBasicField | []*GongBasicField | *GongEnum | []*GongEnum | *GongEnumValue | []*GongEnumValue | *GongNote | []*GongNote | *GongStruct | []*GongStruct | *GongTimeField | []*GongTimeField | *ModelPkg | []*ModelPkg | *PointerToGongStructField | []*PointerToGongStructField | *SliceOfPointerToGongStructField | []*SliceOfPointerToGongStructField
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
	case *GongBasicField:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "BasicKindName"
			}
			if field == "3" {
				return "DeclaredType"
			}
			if field == "4" {
				return "CompositeStructName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 5 {
				return "Index"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongEnum:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "Type"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongEnumValue:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Value"
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
	case *GongNote:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Body"
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
	case *GongStruct:
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
	case *GongTimeField:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "2" {
				return "CompositeStructName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "Index"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *ModelPkg:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "PkgPath"
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
	case *PointerToGongStructField:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "CompositeStructName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 2 {
				return "Index"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *SliceOfPointerToGongStructField:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "CompositeStructName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 2 {
				return "Index"
			}
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
