// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Enum_BookFormatEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_bookformatenum Enum_BookFormatEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_bookformatenum {
	// insertion code per enum code
	case Enum_BookFormatEnum_Paperback:
		res = "Paperback"
	case Enum_BookFormatEnum_Hardcover:
		res = "Hardcover"
	}
	return
}

func (enum_bookformatenum *Enum_BookFormatEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Paperback":
		*enum_bookformatenum = Enum_BookFormatEnum_Paperback
		return
	case "Hardcover":
		*enum_bookformatenum = Enum_BookFormatEnum_Hardcover
		return
	default:
		return errUnkownEnum
	}
}

func (enum_bookformatenum *Enum_BookFormatEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_BookFormatEnum_Paperback":
		*enum_bookformatenum = Enum_BookFormatEnum_Paperback
	case "Enum_BookFormatEnum_Hardcover":
		*enum_bookformatenum = Enum_BookFormatEnum_Hardcover
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_bookformatenum *Enum_BookFormatEnum) ToCodeString() (res string) {

	switch *enum_bookformatenum {
	// insertion code per enum code
	case Enum_BookFormatEnum_Paperback:
		res = "Enum_BookFormatEnum_Paperback"
	case Enum_BookFormatEnum_Hardcover:
		res = "Enum_BookFormatEnum_Hardcover"
	}
	return
}

func (enum_bookformatenum Enum_BookFormatEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_BookFormatEnum_Paperback")
	res = append(res, "Enum_BookFormatEnum_Hardcover")

	return
}

func (enum_bookformatenum Enum_BookFormatEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Paperback")
	res = append(res, "Hardcover")

	return
}

// Utility function for Enum_TitleType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_titletype Enum_TitleType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_titletype {
	// insertion code per enum code
	}
	return
}

func (enum_titletype *Enum_TitleType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_titletype *Enum_TitleType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_titletype *Enum_TitleType) ToCodeString() (res string) {

	switch *enum_titletype {
	// insertion code per enum code
	}
	return
}

func (enum_titletype Enum_TitleType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_titletype Enum_TitleType) CodeValues() (res []string) {

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
