// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (direction Direction) ToString() (res string) {

	// migration of former implementation of enum
	switch direction {
	// insertion code per enum code
	case Vertical:
		res = "vertical"
	case Horizontal:
		res = "horizontal"
	}
	return
}

func (direction *Direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "vertical":
		*direction = Vertical
		return
	case "horizontal":
		*direction = Horizontal
		return
	default:
		return errUnkownEnum
	}
}

func (direction *Direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Vertical":
		*direction = Vertical
	case "Horizontal":
		*direction = Horizontal
	default:
		return errUnkownEnum
	}
	return
}

func (direction *Direction) ToCodeString() (res string) {

	switch *direction {
	// insertion code per enum code
	case Vertical:
		res = "Vertical"
	case Horizontal:
		res = "Horizontal"
	}
	return
}

func (direction Direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Vertical")
	res = append(res, "Horizontal")

	return
}

func (direction Direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "vertical")
	res = append(res, "horizontal")

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
