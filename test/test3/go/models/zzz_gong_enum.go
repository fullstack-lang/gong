// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for EnumTypeInt
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtypeint EnumTypeInt) ToInt() (res int) {

	// migration of former implementation of enum
	switch enumtypeint {
	// insertion code per enum code
	case EnumTypeInt_Value1:
		res = 0
	case EnumTypeInt_Value2:
		res = 1
	}
	return
}

func (enumtypeint *EnumTypeInt) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*enumtypeint = EnumTypeInt_Value1
		return
	case 1:
		*enumtypeint = EnumTypeInt_Value2
		return
	default:
		return errUnkownEnum
	}
}

func (enumtypeint *EnumTypeInt) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTypeInt_Value1":
		*enumtypeint = EnumTypeInt_Value1
	case "EnumTypeInt_Value2":
		*enumtypeint = EnumTypeInt_Value2
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtypeint *EnumTypeInt) ToCodeString() (res string) {

	switch *enumtypeint {
	// insertion code per enum code
	case EnumTypeInt_Value1:
		res = "EnumTypeInt_Value1"
	case EnumTypeInt_Value2:
		res = "EnumTypeInt_Value2"
	}
	return
}

func (enumtypeint EnumTypeInt) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTypeInt_Value1")
	res = append(res, "EnumTypeInt_Value2")

	return
}

func (enumtypeint EnumTypeInt) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for EnumTypeString
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtypestring EnumTypeString) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtypestring {
	// insertion code per enum code
	case EnumTypeString_Value1:
		res = "Value1"
	case EnumTypeString_Value2:
		res = "Value2"
	}
	return
}

func (enumtypestring *EnumTypeString) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Value1":
		*enumtypestring = EnumTypeString_Value1
		return
	case "Value2":
		*enumtypestring = EnumTypeString_Value2
		return
	default:
		return errUnkownEnum
	}
}

func (enumtypestring *EnumTypeString) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTypeString_Value1":
		*enumtypestring = EnumTypeString_Value1
	case "EnumTypeString_Value2":
		*enumtypestring = EnumTypeString_Value2
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtypestring *EnumTypeString) ToCodeString() (res string) {

	switch *enumtypestring {
	// insertion code per enum code
	case EnumTypeString_Value1:
		res = "EnumTypeString_Value1"
	case EnumTypeString_Value2:
		res = "EnumTypeString_Value2"
	}
	return
}

func (enumtypestring EnumTypeString) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTypeString_Value1")
	res = append(res, "EnumTypeString_Value2")

	return
}

func (enumtypestring EnumTypeString) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Value1")
	res = append(res, "Value2")

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
	int | EnumTypeInt
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types | *EnumTypeInt
	FromCodeString(input string) (err error)
}

// Last line of the template
