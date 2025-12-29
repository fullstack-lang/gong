// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for MatButtonAppearance
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (matbuttonappearance MatButtonAppearance) ToString() (res string) {

	// migration of former implementation of enum
	switch matbuttonappearance {
	// insertion code per enum code
	case MatButtonAppearanceText:
		res = "text"
	case MatButtonAppearanceFilled:
		res = "filled"
	case MatButtonAppearanceElevated:
		res = "elevated"
	case MatButtonAppearanceOutlined:
		res = "outlined"
	case MatButtonAppearanceTonal:
		res = "tonal"
	}
	return
}

func (matbuttonappearance *MatButtonAppearance) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "text":
		*matbuttonappearance = MatButtonAppearanceText
		return
	case "filled":
		*matbuttonappearance = MatButtonAppearanceFilled
		return
	case "elevated":
		*matbuttonappearance = MatButtonAppearanceElevated
		return
	case "outlined":
		*matbuttonappearance = MatButtonAppearanceOutlined
		return
	case "tonal":
		*matbuttonappearance = MatButtonAppearanceTonal
		return
	default:
		return errUnkownEnum
	}
}

func (matbuttonappearance *MatButtonAppearance) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MatButtonAppearanceText":
		*matbuttonappearance = MatButtonAppearanceText
	case "MatButtonAppearanceFilled":
		*matbuttonappearance = MatButtonAppearanceFilled
	case "MatButtonAppearanceElevated":
		*matbuttonappearance = MatButtonAppearanceElevated
	case "MatButtonAppearanceOutlined":
		*matbuttonappearance = MatButtonAppearanceOutlined
	case "MatButtonAppearanceTonal":
		*matbuttonappearance = MatButtonAppearanceTonal
	default:
		err = errUnkownEnum
	}
	return
}

func (matbuttonappearance *MatButtonAppearance) ToCodeString() (res string) {

	switch *matbuttonappearance {
	// insertion code per enum code
	case MatButtonAppearanceText:
		res = "MatButtonAppearanceText"
	case MatButtonAppearanceFilled:
		res = "MatButtonAppearanceFilled"
	case MatButtonAppearanceElevated:
		res = "MatButtonAppearanceElevated"
	case MatButtonAppearanceOutlined:
		res = "MatButtonAppearanceOutlined"
	case MatButtonAppearanceTonal:
		res = "MatButtonAppearanceTonal"
	}
	return
}

func (matbuttonappearance MatButtonAppearance) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MatButtonAppearanceText")
	res = append(res, "MatButtonAppearanceFilled")
	res = append(res, "MatButtonAppearanceElevated")
	res = append(res, "MatButtonAppearanceOutlined")
	res = append(res, "MatButtonAppearanceTonal")

	return
}

func (matbuttonappearance MatButtonAppearance) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "text")
	res = append(res, "filled")
	res = append(res, "elevated")
	res = append(res, "outlined")
	res = append(res, "tonal")

	return
}

// Utility function for MatButtonPaletteType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (matbuttonpalettetype MatButtonPaletteType) ToString() (res string) {

	// migration of former implementation of enum
	switch matbuttonpalettetype {
	// insertion code per enum code
	case MatButtonPaletteTypePrimary:
		res = "primary"
	case MatButtonPaletteTypeWarn:
		res = "warn"
	case MatButtonPaletteTypeAccent:
		res = "accent"
	}
	return
}

func (matbuttonpalettetype *MatButtonPaletteType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "primary":
		*matbuttonpalettetype = MatButtonPaletteTypePrimary
		return
	case "warn":
		*matbuttonpalettetype = MatButtonPaletteTypeWarn
		return
	case "accent":
		*matbuttonpalettetype = MatButtonPaletteTypeAccent
		return
	default:
		return errUnkownEnum
	}
}

func (matbuttonpalettetype *MatButtonPaletteType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MatButtonPaletteTypePrimary":
		*matbuttonpalettetype = MatButtonPaletteTypePrimary
	case "MatButtonPaletteTypeWarn":
		*matbuttonpalettetype = MatButtonPaletteTypeWarn
	case "MatButtonPaletteTypeAccent":
		*matbuttonpalettetype = MatButtonPaletteTypeAccent
	default:
		err = errUnkownEnum
	}
	return
}

func (matbuttonpalettetype *MatButtonPaletteType) ToCodeString() (res string) {

	switch *matbuttonpalettetype {
	// insertion code per enum code
	case MatButtonPaletteTypePrimary:
		res = "MatButtonPaletteTypePrimary"
	case MatButtonPaletteTypeWarn:
		res = "MatButtonPaletteTypeWarn"
	case MatButtonPaletteTypeAccent:
		res = "MatButtonPaletteTypeAccent"
	}
	return
}

func (matbuttonpalettetype MatButtonPaletteType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MatButtonPaletteTypePrimary")
	res = append(res, "MatButtonPaletteTypeWarn")
	res = append(res, "MatButtonPaletteTypeAccent")

	return
}

func (matbuttonpalettetype MatButtonPaletteType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "primary")
	res = append(res, "warn")
	res = append(res, "accent")

	return
}

// Utility function for MatButtonType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (matbuttontype MatButtonType) ToString() (res string) {

	// migration of former implementation of enum
	switch matbuttontype {
	// insertion code per enum code
	case MatButtonTypeBasic:
		res = "matButton"
	case MatButtonTypeIcon:
		res = "matIconButton"
	case MatButtonTypeFab:
		res = "matFab"
	case MatButtonTypeMiniFab:
		res = "matMiniFab"
	case MatButtonTypeExtendedFab:
		res = "matFab extended"
	}
	return
}

func (matbuttontype *MatButtonType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "matButton":
		*matbuttontype = MatButtonTypeBasic
		return
	case "matIconButton":
		*matbuttontype = MatButtonTypeIcon
		return
	case "matFab":
		*matbuttontype = MatButtonTypeFab
		return
	case "matMiniFab":
		*matbuttontype = MatButtonTypeMiniFab
		return
	case "matFab extended":
		*matbuttontype = MatButtonTypeExtendedFab
		return
	default:
		return errUnkownEnum
	}
}

func (matbuttontype *MatButtonType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MatButtonTypeBasic":
		*matbuttontype = MatButtonTypeBasic
	case "MatButtonTypeIcon":
		*matbuttontype = MatButtonTypeIcon
	case "MatButtonTypeFab":
		*matbuttontype = MatButtonTypeFab
	case "MatButtonTypeMiniFab":
		*matbuttontype = MatButtonTypeMiniFab
	case "MatButtonTypeExtendedFab":
		*matbuttontype = MatButtonTypeExtendedFab
	default:
		err = errUnkownEnum
	}
	return
}

func (matbuttontype *MatButtonType) ToCodeString() (res string) {

	switch *matbuttontype {
	// insertion code per enum code
	case MatButtonTypeBasic:
		res = "MatButtonTypeBasic"
	case MatButtonTypeIcon:
		res = "MatButtonTypeIcon"
	case MatButtonTypeFab:
		res = "MatButtonTypeFab"
	case MatButtonTypeMiniFab:
		res = "MatButtonTypeMiniFab"
	case MatButtonTypeExtendedFab:
		res = "MatButtonTypeExtendedFab"
	}
	return
}

func (matbuttontype MatButtonType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MatButtonTypeBasic")
	res = append(res, "MatButtonTypeIcon")
	res = append(res, "MatButtonTypeFab")
	res = append(res, "MatButtonTypeMiniFab")
	res = append(res, "MatButtonTypeExtendedFab")

	return
}

func (matbuttontype MatButtonType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "matButton")
	res = append(res, "matIconButton")
	res = append(res, "matFab")
	res = append(res, "matMiniFab")
	res = append(res, "matFab extended")

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case ButtonStackName:
		res = "button"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "button":
		*stacksnames = ButtonStackName
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ButtonStackName":
		*stacksnames = ButtonStackName
	default:
		err = errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case ButtonStackName:
		res = "ButtonStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ButtonStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "button")

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
