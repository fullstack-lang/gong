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
	case PERCENT_075:
		res = "75 %"
	case PERCENT_050:
		res = "50 %"
	case PERCENT_025:
		res = "25 %"
	case PERCENT_000:
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
		*completionenum = PERCENT_075
		return
	case "50 %":
		*completionenum = PERCENT_050
		return
	case "25 %":
		*completionenum = PERCENT_025
		return
	case "0 %":
		*completionenum = PERCENT_000
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
	case "PERCENT_075":
		*completionenum = PERCENT_075
	case "PERCENT_050":
		*completionenum = PERCENT_050
	case "PERCENT_025":
		*completionenum = PERCENT_025
	case "PERCENT_000":
		*completionenum = PERCENT_000
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
	case PERCENT_075:
		res = "PERCENT_075"
	case PERCENT_050:
		res = "PERCENT_050"
	case PERCENT_025:
		res = "PERCENT_025"
	case PERCENT_000:
		res = "PERCENT_000"
	}
	return
}

func (completionenum CompletionEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PERCENT_100")
	res = append(res, "PERCENT_075")
	res = append(res, "PERCENT_050")
	res = append(res, "PERCENT_025")
	res = append(res, "PERCENT_000")

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

// Utility function for LayoutDirection
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (layoutdirection LayoutDirection) ToInt() (res int) {

	// migration of former implementation of enum
	switch layoutdirection {
	// insertion code per enum code
	case Vertical:
		res = 0
	case Horizontal:
		res = 1
	}
	return
}

func (layoutdirection *LayoutDirection) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*layoutdirection = Vertical
		return
	case 1:
		*layoutdirection = Horizontal
		return
	default:
		return errUnkownEnum
	}
}

func (layoutdirection *LayoutDirection) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Vertical":
		*layoutdirection = Vertical
	case "Horizontal":
		*layoutdirection = Horizontal
	default:
		err = errUnkownEnum
	}
	return
}

func (layoutdirection *LayoutDirection) ToCodeString() (res string) {

	switch *layoutdirection {
	// insertion code per enum code
	case Vertical:
		res = "Vertical"
	case Horizontal:
		res = "Horizontal"
	}
	return
}

func (layoutdirection LayoutDirection) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Vertical")
	res = append(res, "Horizontal")

	return
}

func (layoutdirection LayoutDirection) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

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

// Utility function for TimeStepScaleEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (timestepscaleenum TimeStepScaleEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch timestepscaleenum {
	// insertion code per enum code
	case YEARS:
		res = "YEARS"
	case MONTHS:
		res = "MONTHS"
	case WEEKS:
		res = "WEEKS"
	case DAYS:
		res = "DAYS"
	}
	return
}

func (timestepscaleenum *TimeStepScaleEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "YEARS":
		*timestepscaleenum = YEARS
		return
	case "MONTHS":
		*timestepscaleenum = MONTHS
		return
	case "WEEKS":
		*timestepscaleenum = WEEKS
		return
	case "DAYS":
		*timestepscaleenum = DAYS
		return
	default:
		return errUnkownEnum
	}
}

func (timestepscaleenum *TimeStepScaleEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "YEARS":
		*timestepscaleenum = YEARS
	case "MONTHS":
		*timestepscaleenum = MONTHS
	case "WEEKS":
		*timestepscaleenum = WEEKS
	case "DAYS":
		*timestepscaleenum = DAYS
	default:
		err = errUnkownEnum
	}
	return
}

func (timestepscaleenum *TimeStepScaleEnum) ToCodeString() (res string) {

	switch *timestepscaleenum {
	// insertion code per enum code
	case YEARS:
		res = "YEARS"
	case MONTHS:
		res = "MONTHS"
	case WEEKS:
		res = "WEEKS"
	case DAYS:
		res = "DAYS"
	}
	return
}

func (timestepscaleenum TimeStepScaleEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "YEARS")
	res = append(res, "MONTHS")
	res = append(res, "WEEKS")
	res = append(res, "DAYS")

	return
}

func (timestepscaleenum TimeStepScaleEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "YEARS")
	res = append(res, "MONTHS")
	res = append(res, "WEEKS")
	res = append(res, "DAYS")

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
	int | LayoutDirection
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types | *LayoutDirection
	FromCodeString(input string) (err error)
}

// Last line of the template
