// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GongEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (gongenumtype GongEnumType) ToInt() (res int) {

	// migration of former implementation of enum
	switch gongenumtype {
	// insertion code per enum code
	case Int:
		res = 0
	case String:
		res = 1
	}
	return
}

func (gongenumtype *GongEnumType) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*gongenumtype = Int
		return
	case 1:
		*gongenumtype = String
		return
	default:
		return errUnkownEnum
	}
}

func (gongenumtype *GongEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Int":
		*gongenumtype = Int
	case "String":
		*gongenumtype = String
	default:
		return errUnkownEnum
	}
	return
}

func (gongenumtype *GongEnumType) ToCodeString() (res string) {

	switch *gongenumtype {
	// insertion code per enum code
	case Int:
		res = "Int"
	case String:
		res = "String"
	}
	return
}

func (gongenumtype GongEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Int")
	res = append(res, "String")

	return
}

func (gongenumtype GongEnumType) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

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
	int | GongEnumType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*GongEnumType
	FromCodeString(input string) (err error)
}

// Last line of the template
