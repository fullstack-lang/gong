// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Target
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (target Target) ToString() (res string) {

	// migration of former implementation of enum
	switch target {
	// insertion code per enum code
	case FILE:
		res = "FILE"
	case WEB:
		res = "WEB"
	}
	return
}

func (target *Target) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FILE":
		*target = FILE
		return
	case "WEB":
		*target = WEB
		return
	default:
		return errUnkownEnum
	}
}

func (target *Target) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FILE":
		*target = FILE
	case "WEB":
		*target = WEB
	default:
		return errUnkownEnum
	}
	return
}

func (target *Target) ToCodeString() (res string) {

	switch *target {
	// insertion code per enum code
	case FILE:
		res = "FILE"
	case WEB:
		res = "WEB"
	}
	return
}

func (target Target) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FILE")
	res = append(res, "WEB")

	return
}

func (target Target) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FILE")
	res = append(res, "WEB")

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
