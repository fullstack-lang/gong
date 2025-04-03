// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for FieldType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fieldtype FieldType) ToInt() (res int) {

	// migration of former implementation of enum
	switch fieldtype {
	// insertion code per enum code
	case FieldTypeInt:
		res = 0
	case FieldTypeFloat64:
		res = 1
	case FieldTypeBool:
		res = 2
	case FieldTypeTime:
		res = 3
	case FieldTypeDuration:
		res = 4
	case FieldTypePointer:
		res = 5
	case FieldTypeSliceOfPointer:
		res = 6
	}
	return
}

func (fieldtype *FieldType) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*fieldtype = FieldTypeInt
		return
	case 1:
		*fieldtype = FieldTypeFloat64
		return
	case 2:
		*fieldtype = FieldTypeBool
		return
	case 3:
		*fieldtype = FieldTypeTime
		return
	case 4:
		*fieldtype = FieldTypeDuration
		return
	case 5:
		*fieldtype = FieldTypePointer
		return
	case 6:
		*fieldtype = FieldTypeSliceOfPointer
		return
	default:
		return errUnkownEnum
	}
}

func (fieldtype *FieldType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FieldTypeInt":
		*fieldtype = FieldTypeInt
	case "FieldTypeFloat64":
		*fieldtype = FieldTypeFloat64
	case "FieldTypeBool":
		*fieldtype = FieldTypeBool
	case "FieldTypeTime":
		*fieldtype = FieldTypeTime
	case "FieldTypeDuration":
		*fieldtype = FieldTypeDuration
	case "FieldTypePointer":
		*fieldtype = FieldTypePointer
	case "FieldTypeSliceOfPointer":
		*fieldtype = FieldTypeSliceOfPointer
	default:
		return errUnkownEnum
	}
	return
}

func (fieldtype *FieldType) ToCodeString() (res string) {

	switch *fieldtype {
	// insertion code per enum code
	case FieldTypeInt:
		res = "FieldTypeInt"
	case FieldTypeFloat64:
		res = "FieldTypeFloat64"
	case FieldTypeBool:
		res = "FieldTypeBool"
	case FieldTypeTime:
		res = "FieldTypeTime"
	case FieldTypeDuration:
		res = "FieldTypeDuration"
	case FieldTypePointer:
		res = "FieldTypePointer"
	case FieldTypeSliceOfPointer:
		res = "FieldTypeSliceOfPointer"
	}
	return
}

func (fieldtype FieldType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FieldTypeInt")
	res = append(res, "FieldTypeFloat64")
	res = append(res, "FieldTypeBool")
	res = append(res, "FieldTypeTime")
	res = append(res, "FieldTypeDuration")
	res = append(res, "FieldTypePointer")
	res = append(res, "FieldTypeSliceOfPointer")

	return
}

func (fieldtype FieldType) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)
	res = append(res, 4)
	res = append(res, 5)
	res = append(res, 6)

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | FieldType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*FieldType
	FromCodeString(input string) (err error)
}

// Last line of the template
