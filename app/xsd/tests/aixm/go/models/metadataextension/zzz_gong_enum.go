// generated code - do not edit
package metadataextension

// insertion point of enum utility functions
// Utility function for Enum_MD_ObligationCode_Type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_md_obligationcode_type Enum_MD_ObligationCode_Type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_md_obligationcode_type {
	// insertion code per enum code
	case Enum_MD_ObligationCode_Type_Mandatory:
		res = "mandatory"
	case Enum_MD_ObligationCode_Type_Optional:
		res = "optional"
	case Enum_MD_ObligationCode_Type_Conditional:
		res = "conditional"
	}
	return
}

func (enum_md_obligationcode_type *Enum_MD_ObligationCode_Type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "mandatory":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Mandatory
		return
	case "optional":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Optional
		return
	case "conditional":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Conditional
		return
	default:
		return errUnkownEnum
	}
}

func (enum_md_obligationcode_type *Enum_MD_ObligationCode_Type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_MD_ObligationCode_Type_Mandatory":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Mandatory
	case "Enum_MD_ObligationCode_Type_Optional":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Optional
	case "Enum_MD_ObligationCode_Type_Conditional":
		*enum_md_obligationcode_type = Enum_MD_ObligationCode_Type_Conditional
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_md_obligationcode_type *Enum_MD_ObligationCode_Type) ToCodeString() (res string) {

	switch *enum_md_obligationcode_type {
	// insertion code per enum code
	case Enum_MD_ObligationCode_Type_Mandatory:
		res = "Enum_MD_ObligationCode_Type_Mandatory"
	case Enum_MD_ObligationCode_Type_Optional:
		res = "Enum_MD_ObligationCode_Type_Optional"
	case Enum_MD_ObligationCode_Type_Conditional:
		res = "Enum_MD_ObligationCode_Type_Conditional"
	}
	return
}

func (enum_md_obligationcode_type Enum_MD_ObligationCode_Type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_MD_ObligationCode_Type_Mandatory")
	res = append(res, "Enum_MD_ObligationCode_Type_Optional")
	res = append(res, "Enum_MD_ObligationCode_Type_Conditional")

	return
}

func (enum_md_obligationcode_type Enum_MD_ObligationCode_Type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "mandatory")
	res = append(res, "optional")
	res = append(res, "conditional")

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
