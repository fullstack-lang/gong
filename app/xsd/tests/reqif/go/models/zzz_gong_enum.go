// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Enum_GLOBAL_REF
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_global_ref Enum_GLOBAL_REF) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_global_ref {
	// insertion code per enum code
	}
	return
}

func (enum_global_ref *Enum_GLOBAL_REF) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_global_ref *Enum_GLOBAL_REF) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_global_ref *Enum_GLOBAL_REF) ToCodeString() (res string) {

	switch *enum_global_ref {
	// insertion code per enum code
	}
	return
}

func (enum_global_ref Enum_GLOBAL_REF) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_global_ref Enum_GLOBAL_REF) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

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
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

// Last line of the template
