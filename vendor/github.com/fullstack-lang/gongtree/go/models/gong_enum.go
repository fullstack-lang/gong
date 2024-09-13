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
		return errUnkownEnum
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
		return errUnkownEnum
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
