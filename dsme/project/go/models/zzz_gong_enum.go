// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for CompletionEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (completionenum CompletionEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch completionenum {
	// insertion code per enum code
	case PERCENT_100:
		res = "100 %"
	case PERCENT_75:
		res = "75 %"
	case PERCENT_50:
		res = "50 %"
	case PERCENT_25:
		res = "25 %"
	case PERCENT_00:
		res = "0 %"
	}
	return
}

func (completionenum *CompletionEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "100 %":
		*completionenum = PERCENT_100
		return
	case "75 %":
		*completionenum = PERCENT_75
		return
	case "50 %":
		*completionenum = PERCENT_50
		return
	case "25 %":
		*completionenum = PERCENT_25
		return
	case "0 %":
		*completionenum = PERCENT_00
		return
	default:
		return errUnkownEnum
	}
}

func (completionenum *CompletionEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PERCENT_100":
		*completionenum = PERCENT_100
	case "PERCENT_75":
		*completionenum = PERCENT_75
	case "PERCENT_50":
		*completionenum = PERCENT_50
	case "PERCENT_25":
		*completionenum = PERCENT_25
	case "PERCENT_00":
		*completionenum = PERCENT_00
	default:
		err = errUnkownEnum
	}
	return
}

func (completionenum *CompletionEnum) ToCodeString() (res string) {

	switch *completionenum {
	// insertion code per enum code
	case PERCENT_100:
		res = "PERCENT_100"
	case PERCENT_75:
		res = "PERCENT_75"
	case PERCENT_50:
		res = "PERCENT_50"
	case PERCENT_25:
		res = "PERCENT_25"
	case PERCENT_00:
		res = "PERCENT_00"
	}
	return
}

func (completionenum CompletionEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PERCENT_100")
	res = append(res, "PERCENT_75")
	res = append(res, "PERCENT_50")
	res = append(res, "PERCENT_25")
	res = append(res, "PERCENT_00")

	return
}

func (completionenum CompletionEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "100 %")
	res = append(res, "75 %")
	res = append(res, "50 %")
	res = append(res, "25 %")
	res = append(res, "0 %")

	return
}

// Utility function for OrientationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (orientationtype OrientationType) ToString() (res string) {

	// migration of former implementation of enum
	switch orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype *OrientationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
		return
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
		return
	default:
		return errUnkownEnum
	}
}

func (orientationtype *OrientationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
	default:
		err = errUnkownEnum
	}
	return
}

func (orientationtype *OrientationType) ToCodeString() (res string) {

	switch *orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype OrientationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

	return
}

func (orientationtype OrientationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

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
