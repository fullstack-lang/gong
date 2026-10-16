// generated code - do not edit
package score

// insertion point of enum utility functions
// Utility function for EnumGroupBarlineValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumgroupbarlinevalue EnumGroupBarlineValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumgroupbarlinevalue {
	// insertion code per enum code
	case EnumGroupBarlineValueYes:
		res = "yes"
	case EnumGroupBarlineValueNo:
		res = "no"
	case EnumGroupBarlineValueMensurstrich:
		res = "Mensurstrich"
	}
	return
}

func (enumgroupbarlinevalue *EnumGroupBarlineValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enumgroupbarlinevalue = EnumGroupBarlineValueYes
		return
	case "no":
		*enumgroupbarlinevalue = EnumGroupBarlineValueNo
		return
	case "Mensurstrich":
		*enumgroupbarlinevalue = EnumGroupBarlineValueMensurstrich
		return
	default:
		return errUnkownEnum
	}
}

func (enumgroupbarlinevalue *EnumGroupBarlineValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumGroupBarlineValueYes":
		*enumgroupbarlinevalue = EnumGroupBarlineValueYes
	case "EnumGroupBarlineValueNo":
		*enumgroupbarlinevalue = EnumGroupBarlineValueNo
	case "EnumGroupBarlineValueMensurstrich":
		*enumgroupbarlinevalue = EnumGroupBarlineValueMensurstrich
	default:
		err = errUnkownEnum
	}
	return
}

func (enumgroupbarlinevalue *EnumGroupBarlineValue) ToCodeString() (res string) {

	switch *enumgroupbarlinevalue {
	// insertion code per enum code
	case EnumGroupBarlineValueYes:
		res = "EnumGroupBarlineValueYes"
	case EnumGroupBarlineValueNo:
		res = "EnumGroupBarlineValueNo"
	case EnumGroupBarlineValueMensurstrich:
		res = "EnumGroupBarlineValueMensurstrich"
	}
	return
}

func (enumgroupbarlinevalue EnumGroupBarlineValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumGroupBarlineValueYes")
	res = append(res, "EnumGroupBarlineValueNo")
	res = append(res, "EnumGroupBarlineValueMensurstrich")

	return
}

func (enumgroupbarlinevalue EnumGroupBarlineValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "Mensurstrich")

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
