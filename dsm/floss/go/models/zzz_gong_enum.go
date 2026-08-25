// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for ComplexityType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (complexitytype ComplexityType) ToString() (res string) {

	// migration of former implementation of enum
	switch complexitytype {
	// insertion code per enum code
	case ComponentComplexity:
		res = "Component Complexity"
	case InterfaceComplexity:
		res = "Interface Complexity"
	case TopologicalComplexity:
		res = "Topological Complexity"
	}
	return
}

func (complexitytype *ComplexityType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Component Complexity":
		*complexitytype = ComponentComplexity
		return
	case "Interface Complexity":
		*complexitytype = InterfaceComplexity
		return
	case "Topological Complexity":
		*complexitytype = TopologicalComplexity
		return
	default:
		return errUnkownEnum
	}
}

func (complexitytype *ComplexityType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ComponentComplexity":
		*complexitytype = ComponentComplexity
	case "InterfaceComplexity":
		*complexitytype = InterfaceComplexity
	case "TopologicalComplexity":
		*complexitytype = TopologicalComplexity
	default:
		err = errUnkownEnum
	}
	return
}

func (complexitytype *ComplexityType) ToCodeString() (res string) {

	switch *complexitytype {
	// insertion code per enum code
	case ComponentComplexity:
		res = "ComponentComplexity"
	case InterfaceComplexity:
		res = "InterfaceComplexity"
	case TopologicalComplexity:
		res = "TopologicalComplexity"
	}
	return
}

func (complexitytype ComplexityType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ComponentComplexity")
	res = append(res, "InterfaceComplexity")
	res = append(res, "TopologicalComplexity")

	return
}

func (complexitytype ComplexityType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Component Complexity")
	res = append(res, "Interface Complexity")
	res = append(res, "Topological Complexity")

	return
}

// Utility function for FontSize
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fontsize FontSize) ToString() (res string) {

	// migration of former implementation of enum
	switch fontsize {
	// insertion code per enum code
	case FONT_SIZE_SMALL:
		res = "small"
	case FONT_SIZE_NORMAL:
		res = "normal"
	case FONT_SIZE_BIG:
		res = "big"
	case FONT_SIZE_VERY_BIG:
		res = "very big"
	}
	return
}

func (fontsize *FontSize) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "small":
		*fontsize = FONT_SIZE_SMALL
		return
	case "normal":
		*fontsize = FONT_SIZE_NORMAL
		return
	case "big":
		*fontsize = FONT_SIZE_BIG
		return
	case "very big":
		*fontsize = FONT_SIZE_VERY_BIG
		return
	default:
		return errUnkownEnum
	}
}

func (fontsize *FontSize) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FONT_SIZE_SMALL":
		*fontsize = FONT_SIZE_SMALL
	case "FONT_SIZE_NORMAL":
		*fontsize = FONT_SIZE_NORMAL
	case "FONT_SIZE_BIG":
		*fontsize = FONT_SIZE_BIG
	case "FONT_SIZE_VERY_BIG":
		*fontsize = FONT_SIZE_VERY_BIG
	default:
		err = errUnkownEnum
	}
	return
}

func (fontsize *FontSize) ToCodeString() (res string) {

	switch *fontsize {
	// insertion code per enum code
	case FONT_SIZE_SMALL:
		res = "FONT_SIZE_SMALL"
	case FONT_SIZE_NORMAL:
		res = "FONT_SIZE_NORMAL"
	case FONT_SIZE_BIG:
		res = "FONT_SIZE_BIG"
	case FONT_SIZE_VERY_BIG:
		res = "FONT_SIZE_VERY_BIG"
	}
	return
}

func (fontsize FontSize) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FONT_SIZE_SMALL")
	res = append(res, "FONT_SIZE_NORMAL")
	res = append(res, "FONT_SIZE_BIG")
	res = append(res, "FONT_SIZE_VERY_BIG")

	return
}

func (fontsize FontSize) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "small")
	res = append(res, "normal")
	res = append(res, "big")
	res = append(res, "very big")

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
