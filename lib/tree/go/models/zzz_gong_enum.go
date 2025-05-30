// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for FontStyleEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fontstyleenum FontStyleEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch fontstyleenum {
	// insertion code per enum code
	case NORMAL:
		res = "NORMAL"
	case ITALIC:
		res = "ITALIC"
	}
	return
}

func (fontstyleenum *FontStyleEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*fontstyleenum = NORMAL
		return
	case "ITALIC":
		*fontstyleenum = ITALIC
		return
	default:
		return errUnkownEnum
	}
}

func (fontstyleenum *FontStyleEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*fontstyleenum = NORMAL
	case "ITALIC":
		*fontstyleenum = ITALIC
	default:
		err = errUnkownEnum
	}
	return
}

func (fontstyleenum *FontStyleEnum) ToCodeString() (res string) {

	switch *fontstyleenum {
	// insertion code per enum code
	case NORMAL:
		res = "NORMAL"
	case ITALIC:
		res = "ITALIC"
	}
	return
}

func (fontstyleenum FontStyleEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "ITALIC")

	return
}

func (fontstyleenum FontStyleEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "ITALIC")

	return
}

// Utility function for ToolTipPositionEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tooltippositionenum ToolTipPositionEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tooltippositionenum {
	// insertion code per enum code
	case Below:
		res = "below"
	case Above:
		res = "above"
	case Left:
		res = "left"
	case Right:
		res = "right"
	}
	return
}

func (tooltippositionenum *ToolTipPositionEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "below":
		*tooltippositionenum = Below
		return
	case "above":
		*tooltippositionenum = Above
		return
	case "left":
		*tooltippositionenum = Left
		return
	case "right":
		*tooltippositionenum = Right
		return
	default:
		return errUnkownEnum
	}
}

func (tooltippositionenum *ToolTipPositionEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Below":
		*tooltippositionenum = Below
	case "Above":
		*tooltippositionenum = Above
	case "Left":
		*tooltippositionenum = Left
	case "Right":
		*tooltippositionenum = Right
	default:
		err = errUnkownEnum
	}
	return
}

func (tooltippositionenum *ToolTipPositionEnum) ToCodeString() (res string) {

	switch *tooltippositionenum {
	// insertion code per enum code
	case Below:
		res = "Below"
	case Above:
		res = "Above"
	case Left:
		res = "Left"
	case Right:
		res = "Right"
	}
	return
}

func (tooltippositionenum ToolTipPositionEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Below")
	res = append(res, "Above")
	res = append(res, "Left")
	res = append(res, "Right")

	return
}

func (tooltippositionenum ToolTipPositionEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "below")
	res = append(res, "above")
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for TreeStackName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treestackname TreeStackName) ToString() (res string) {

	// migration of former implementation of enum
	switch treestackname {
	// insertion code per enum code
	case TreeStackDefaultName:
		res = "Tree"
	}
	return
}

func (treestackname *TreeStackName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tree":
		*treestackname = TreeStackDefaultName
		return
	default:
		return errUnkownEnum
	}
}

func (treestackname *TreeStackName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TreeStackDefaultName":
		*treestackname = TreeStackDefaultName
	default:
		err = errUnkownEnum
	}
	return
}

func (treestackname *TreeStackName) ToCodeString() (res string) {

	switch *treestackname {
	// insertion code per enum code
	case TreeStackDefaultName:
		res = "TreeStackDefaultName"
	}
	return
}

func (treestackname TreeStackName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TreeStackDefaultName")

	return
}

func (treestackname TreeStackName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tree")

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
	
	FromCodeString(input string) (err error)
}

// Last line of the template
