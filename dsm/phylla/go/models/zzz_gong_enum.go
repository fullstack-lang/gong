// generated code - do not edit
package models

// insertion point of enum utility functions
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

// Utility function for ViewType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (viewtype ViewType) ToString() (res string) {

	// migration of former implementation of enum
	switch viewtype {
	// insertion code per enum code
	case VIEW_TREE_SVG_FORM:
		res = "Tree - SVG - Form"
	case VIEW_TREE_SVG_SLIDER:
		res = "Tree - SVG - Slider"
	case VIEW_TREE_3D_SLIDER:
		res = "Tree - 3D - Slider"
	}
	return
}

func (viewtype *ViewType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tree - SVG - Form":
		*viewtype = VIEW_TREE_SVG_FORM
		return
	case "Tree - SVG - Slider":
		*viewtype = VIEW_TREE_SVG_SLIDER
		return
	case "Tree - 3D - Slider":
		*viewtype = VIEW_TREE_3D_SLIDER
		return
	default:
		return errUnkownEnum
	}
}

func (viewtype *ViewType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "VIEW_TREE_SVG_FORM":
		*viewtype = VIEW_TREE_SVG_FORM
	case "VIEW_TREE_SVG_SLIDER":
		*viewtype = VIEW_TREE_SVG_SLIDER
	case "VIEW_TREE_3D_SLIDER":
		*viewtype = VIEW_TREE_3D_SLIDER
	default:
		err = errUnkownEnum
	}
	return
}

func (viewtype *ViewType) ToCodeString() (res string) {

	switch *viewtype {
	// insertion code per enum code
	case VIEW_TREE_SVG_FORM:
		res = "VIEW_TREE_SVG_FORM"
	case VIEW_TREE_SVG_SLIDER:
		res = "VIEW_TREE_SVG_SLIDER"
	case VIEW_TREE_3D_SLIDER:
		res = "VIEW_TREE_3D_SLIDER"
	}
	return
}

func (viewtype ViewType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "VIEW_TREE_SVG_FORM")
	res = append(res, "VIEW_TREE_SVG_SLIDER")
	res = append(res, "VIEW_TREE_3D_SLIDER")

	return
}

func (viewtype ViewType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tree - SVG - Form")
	res = append(res, "Tree - SVG - Slider")
	res = append(res, "Tree - 3D - Slider")

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
