// generated code - do not edit
package basictypes

// insertion point of enum utility functions
// Utility function for Enum_NilReasonEnumeration_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_nilreasonenumeration_member_1 {
	// insertion code per enum code
	case Enum_NilReasonEnumeration_Member_1_Inapplicable:
		res = "inapplicable"
	case Enum_NilReasonEnumeration_Member_1_Missing:
		res = "missing"
	case Enum_NilReasonEnumeration_Member_1_Template:
		res = "template"
	case Enum_NilReasonEnumeration_Member_1_Unknown:
		res = "unknown"
	case Enum_NilReasonEnumeration_Member_1_Withheld:
		res = "withheld"
	}
	return
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "inapplicable":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Inapplicable
		return
	case "missing":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Missing
		return
	case "template":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Template
		return
	case "unknown":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Unknown
		return
	case "withheld":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Withheld
		return
	default:
		return errUnkownEnum
	}
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_NilReasonEnumeration_Member_1_Inapplicable":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Inapplicable
	case "Enum_NilReasonEnumeration_Member_1_Missing":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Missing
	case "Enum_NilReasonEnumeration_Member_1_Template":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Template
	case "Enum_NilReasonEnumeration_Member_1_Unknown":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Unknown
	case "Enum_NilReasonEnumeration_Member_1_Withheld":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Withheld
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) ToCodeString() (res string) {

	switch *enum_nilreasonenumeration_member_1 {
	// insertion code per enum code
	case Enum_NilReasonEnumeration_Member_1_Inapplicable:
		res = "Enum_NilReasonEnumeration_Member_1_Inapplicable"
	case Enum_NilReasonEnumeration_Member_1_Missing:
		res = "Enum_NilReasonEnumeration_Member_1_Missing"
	case Enum_NilReasonEnumeration_Member_1_Template:
		res = "Enum_NilReasonEnumeration_Member_1_Template"
	case Enum_NilReasonEnumeration_Member_1_Unknown:
		res = "Enum_NilReasonEnumeration_Member_1_Unknown"
	case Enum_NilReasonEnumeration_Member_1_Withheld:
		res = "Enum_NilReasonEnumeration_Member_1_Withheld"
	}
	return
}

func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Inapplicable")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Missing")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Template")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Unknown")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Withheld")

	return
}

func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "inapplicable")
	res = append(res, "missing")
	res = append(res, "template")
	res = append(res, "unknown")
	res = append(res, "withheld")

	return
}

// Utility function for Enum_SignType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_signtype Enum_SignType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_signtype {
	// insertion code per enum code
	case Enum_SignType_N__:
		res = "-"
	case Enum_SignType_N___1:
		res = "+"
	}
	return
}

func (enum_signtype *Enum_SignType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "-":
		*enum_signtype = Enum_SignType_N__
		return
	case "+":
		*enum_signtype = Enum_SignType_N___1
		return
	default:
		return errUnkownEnum
	}
}

func (enum_signtype *Enum_SignType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SignType_N__":
		*enum_signtype = Enum_SignType_N__
	case "Enum_SignType_N___1":
		*enum_signtype = Enum_SignType_N___1
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_signtype *Enum_SignType) ToCodeString() (res string) {

	switch *enum_signtype {
	// insertion code per enum code
	case Enum_SignType_N__:
		res = "Enum_SignType_N__"
	case Enum_SignType_N___1:
		res = "Enum_SignType_N___1"
	}
	return
}

func (enum_signtype Enum_SignType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SignType_N__")
	res = append(res, "Enum_SignType_N___1")

	return
}

func (enum_signtype Enum_SignType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "-")
	res = append(res, "+")

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
