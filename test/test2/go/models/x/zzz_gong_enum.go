// generated code - do not edit
package x

// insertion point of enum utility functions
// Utility function for XEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (xenumtype XEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch xenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (xenumtype *XEnumType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1":
		*xenumtype = ENUM_VAL1
		return
	case "ENUM_VAL2":
		*xenumtype = ENUM_VAL2
		return
	default:
		return errUnkownEnum
	}
}

func (xenumtype *XEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1":
		*xenumtype = ENUM_VAL1
	case "ENUM_VAL2":
		*xenumtype = ENUM_VAL2
	default:
		err = errUnkownEnum
	}
	return
}

func (xenumtype *XEnumType) ToCodeString() (res string) {

	switch *xenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (xenumtype XEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENUM_VAL1")
	res = append(res, "ENUM_VAL2")

	return
}

func (xenumtype XEnumType) CodeValues() (res []string) {

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

type GongEnumStringPointer interface {
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumStringField = GongEnumStringPointer

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type GongEnumIntPointer interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumIntField = GongEnumIntPointer

// Last line of the template
