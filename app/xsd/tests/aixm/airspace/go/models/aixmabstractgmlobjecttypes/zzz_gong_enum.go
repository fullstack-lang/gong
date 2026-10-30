// generated code - do not edit
package aixmabstractgmlobjecttypes

// insertion point of enum utility functions
// Utility function for Enum_Interpretation_SimpleType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_interpretation_simpletype Enum_Interpretation_SimpleType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_interpretation_simpletype {
	// insertion code per enum code
	case Enum_Interpretation_SimpleType_BASELINE:
		res = "BASELINE"
	case Enum_Interpretation_SimpleType_SNAPSHOT:
		res = "SNAPSHOT"
	case Enum_Interpretation_SimpleType_TEMPDELTA:
		res = "TEMPDELTA"
	case Enum_Interpretation_SimpleType_PERMDELTA:
		res = "PERMDELTA"
	}
	return
}

func (enum_interpretation_simpletype *Enum_Interpretation_SimpleType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BASELINE":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_BASELINE
		return
	case "SNAPSHOT":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_SNAPSHOT
		return
	case "TEMPDELTA":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_TEMPDELTA
		return
	case "PERMDELTA":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_PERMDELTA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_interpretation_simpletype *Enum_Interpretation_SimpleType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Interpretation_SimpleType_BASELINE":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_BASELINE
	case "Enum_Interpretation_SimpleType_SNAPSHOT":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_SNAPSHOT
	case "Enum_Interpretation_SimpleType_TEMPDELTA":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_TEMPDELTA
	case "Enum_Interpretation_SimpleType_PERMDELTA":
		*enum_interpretation_simpletype = Enum_Interpretation_SimpleType_PERMDELTA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_interpretation_simpletype *Enum_Interpretation_SimpleType) ToCodeString() (res string) {

	switch *enum_interpretation_simpletype {
	// insertion code per enum code
	case Enum_Interpretation_SimpleType_BASELINE:
		res = "Enum_Interpretation_SimpleType_BASELINE"
	case Enum_Interpretation_SimpleType_SNAPSHOT:
		res = "Enum_Interpretation_SimpleType_SNAPSHOT"
	case Enum_Interpretation_SimpleType_TEMPDELTA:
		res = "Enum_Interpretation_SimpleType_TEMPDELTA"
	case Enum_Interpretation_SimpleType_PERMDELTA:
		res = "Enum_Interpretation_SimpleType_PERMDELTA"
	}
	return
}

func (enum_interpretation_simpletype Enum_Interpretation_SimpleType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Interpretation_SimpleType_BASELINE")
	res = append(res, "Enum_Interpretation_SimpleType_SNAPSHOT")
	res = append(res, "Enum_Interpretation_SimpleType_TEMPDELTA")
	res = append(res, "Enum_Interpretation_SimpleType_PERMDELTA")

	return
}

func (enum_interpretation_simpletype Enum_Interpretation_SimpleType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BASELINE")
	res = append(res, "SNAPSHOT")
	res = append(res, "TEMPDELTA")
	res = append(res, "PERMDELTA")

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
