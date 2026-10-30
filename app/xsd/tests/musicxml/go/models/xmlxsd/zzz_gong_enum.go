// generated code - do not edit
package xmlxsd

// insertion point of enum utility functions
// Utility function for Enum_Lang_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_lang_member_1 Enum_Lang_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_lang_member_1 {
	// insertion code per enum code
	case Enum_Lang_Member_1_Empty:
		res = ""
	}
	return
}

func (enum_lang_member_1 *Enum_Lang_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "":
		*enum_lang_member_1 = Enum_Lang_Member_1_Empty
		return
	default:
		return errUnkownEnum
	}
}

func (enum_lang_member_1 *Enum_Lang_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Lang_Member_1_Empty":
		*enum_lang_member_1 = Enum_Lang_Member_1_Empty
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_lang_member_1 *Enum_Lang_Member_1) ToCodeString() (res string) {

	switch *enum_lang_member_1 {
	// insertion code per enum code
	case Enum_Lang_Member_1_Empty:
		res = "Enum_Lang_Member_1_Empty"
	}
	return
}

func (enum_lang_member_1 Enum_Lang_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Lang_Member_1_Empty")

	return
}

func (enum_lang_member_1 Enum_Lang_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "")

	return
}

// Utility function for Enum_Space
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_space Enum_Space) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_space {
	// insertion code per enum code
	case Enum_Space_Default_XSD:
		res = "default"
	case Enum_Space_Preserve:
		res = "preserve"
	}
	return
}

func (enum_space *Enum_Space) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "default":
		*enum_space = Enum_Space_Default_XSD
		return
	case "preserve":
		*enum_space = Enum_Space_Preserve
		return
	default:
		return errUnkownEnum
	}
}

func (enum_space *Enum_Space) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Space_Default_XSD":
		*enum_space = Enum_Space_Default_XSD
	case "Enum_Space_Preserve":
		*enum_space = Enum_Space_Preserve
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_space *Enum_Space) ToCodeString() (res string) {

	switch *enum_space {
	// insertion code per enum code
	case Enum_Space_Default_XSD:
		res = "Enum_Space_Default_XSD"
	case Enum_Space_Preserve:
		res = "Enum_Space_Preserve"
	}
	return
}

func (enum_space Enum_Space) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Space_Default_XSD")
	res = append(res, "Enum_Space_Preserve")

	return
}

func (enum_space Enum_Space) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "default")
	res = append(res, "preserve")

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
