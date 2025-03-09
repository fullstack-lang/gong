// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for AEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (aenumtype AEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch aenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (aenumtype *AEnumType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1":
		*aenumtype = ENUM_VAL1
		return
	case "ENUM_VAL2":
		*aenumtype = ENUM_VAL2
		return
	default:
		return errUnkownEnum
	}
}

func (aenumtype *AEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1":
		*aenumtype = ENUM_VAL1
	case "ENUM_VAL2":
		*aenumtype = ENUM_VAL2
	default:
		return errUnkownEnum
	}
	return
}

func (aenumtype *AEnumType) ToCodeString() (res string) {

	switch *aenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (aenumtype AEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENUM_VAL1")
	res = append(res, "ENUM_VAL2")

	return
}

func (aenumtype AEnumType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENUM_VAL1")
	res = append(res, "ENUM_VAL2")

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
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
