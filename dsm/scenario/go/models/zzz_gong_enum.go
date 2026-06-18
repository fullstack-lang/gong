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
	case DIRECTION_TOP:
		res = "TOP"
	case DIRECTION_BOTTOM:
		res = "BOTTOM"
	case DIRECTION_LEFT:
		res = "LEFT"
	case DIRECTION_RIGHT:
		res = "RIGHT"
	}
	return
}

func (direction *Direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TOP":
		*direction = DIRECTION_TOP
		return
	case "BOTTOM":
		*direction = DIRECTION_BOTTOM
		return
	case "LEFT":
		*direction = DIRECTION_LEFT
		return
	case "RIGHT":
		*direction = DIRECTION_RIGHT
		return
	default:
		return errUnkownEnum
	}
}

func (direction *Direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DIRECTION_TOP":
		*direction = DIRECTION_TOP
	case "DIRECTION_BOTTOM":
		*direction = DIRECTION_BOTTOM
	case "DIRECTION_LEFT":
		*direction = DIRECTION_LEFT
	case "DIRECTION_RIGHT":
		*direction = DIRECTION_RIGHT
	default:
		err = errUnkownEnum
	}
	return
}

func (direction *Direction) ToCodeString() (res string) {

	switch *direction {
	// insertion code per enum code
	case DIRECTION_TOP:
		res = "DIRECTION_TOP"
	case DIRECTION_BOTTOM:
		res = "DIRECTION_BOTTOM"
	case DIRECTION_LEFT:
		res = "DIRECTION_LEFT"
	case DIRECTION_RIGHT:
		res = "DIRECTION_RIGHT"
	}
	return
}

func (direction Direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DIRECTION_TOP")
	res = append(res, "DIRECTION_BOTTOM")
	res = append(res, "DIRECTION_LEFT")
	res = append(res, "DIRECTION_RIGHT")

	return
}

func (direction Direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TOP")
	res = append(res, "BOTTOM")
	res = append(res, "LEFT")
	res = append(res, "RIGHT")

	return
}

// Utility function for DirectionType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (directiontype DirectionType) ToString() (res string) {

	// migration of former implementation of enum
	switch directiontype {
	// insertion code per enum code
	case DIRECTION_UP:
		res = "DIRECTION_UP"
	case DIRECTION_DOWN:
		res = "DIRECTION_DOWN"
	}
	return
}

func (directiontype *DirectionType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DIRECTION_UP":
		*directiontype = DIRECTION_UP
		return
	case "DIRECTION_DOWN":
		*directiontype = DIRECTION_DOWN
		return
	default:
		return errUnkownEnum
	}
}

func (directiontype *DirectionType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DIRECTION_UP":
		*directiontype = DIRECTION_UP
	case "DIRECTION_DOWN":
		*directiontype = DIRECTION_DOWN
	default:
		err = errUnkownEnum
	}
	return
}

func (directiontype *DirectionType) ToCodeString() (res string) {

	switch *directiontype {
	// insertion code per enum code
	case DIRECTION_UP:
		res = "DIRECTION_UP"
	case DIRECTION_DOWN:
		res = "DIRECTION_DOWN"
	}
	return
}

func (directiontype DirectionType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DIRECTION_UP")
	res = append(res, "DIRECTION_DOWN")

	return
}

func (directiontype DirectionType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DIRECTION_UP")
	res = append(res, "DIRECTION_DOWN")

	return
}

// Utility function for FormNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (formnames FormNames) ToString() (res string) {

	// migration of former implementation of enum
	switch formnames {
	// insertion code per enum code
	case DefaultForm:
		res = "default form"
	case ModelForm:
		res = "Model Form"
	case ShapeForm:
		res = "Shape Form"
	}
	return
}

func (formnames *FormNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "default form":
		*formnames = DefaultForm
		return
	case "Model Form":
		*formnames = ModelForm
		return
	case "Shape Form":
		*formnames = ShapeForm
		return
	default:
		return errUnkownEnum
	}
}

func (formnames *FormNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DefaultForm":
		*formnames = DefaultForm
	case "ModelForm":
		*formnames = ModelForm
	case "ShapeForm":
		*formnames = ShapeForm
	default:
		err = errUnkownEnum
	}
	return
}

func (formnames *FormNames) ToCodeString() (res string) {

	switch *formnames {
	// insertion code per enum code
	case DefaultForm:
		res = "DefaultForm"
	case ModelForm:
		res = "ModelForm"
	case ShapeForm:
		res = "ShapeForm"
	}
	return
}

func (formnames FormNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DefaultForm")
	res = append(res, "ModelForm")
	res = append(res, "ShapeForm")

	return
}

func (formnames FormNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "default form")
	res = append(res, "Model Form")
	res = append(res, "Shape Form")

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

// Utility function for ProbabilityEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (probabilityenum ProbabilityEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch probabilityenum {
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

func (probabilityenum *ProbabilityEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "100 %":
		*probabilityenum = PERCENT_100
		return
	case "75 %":
		*probabilityenum = PERCENT_75
		return
	case "50 %":
		*probabilityenum = PERCENT_50
		return
	case "25 %":
		*probabilityenum = PERCENT_25
		return
	case "0 %":
		*probabilityenum = PERCENT_00
		return
	default:
		return errUnkownEnum
	}
}

func (probabilityenum *ProbabilityEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PERCENT_100":
		*probabilityenum = PERCENT_100
	case "PERCENT_75":
		*probabilityenum = PERCENT_75
	case "PERCENT_50":
		*probabilityenum = PERCENT_50
	case "PERCENT_25":
		*probabilityenum = PERCENT_25
	case "PERCENT_00":
		*probabilityenum = PERCENT_00
	default:
		err = errUnkownEnum
	}
	return
}

func (probabilityenum *ProbabilityEnum) ToCodeString() (res string) {

	switch *probabilityenum {
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

func (probabilityenum ProbabilityEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PERCENT_100")
	res = append(res, "PERCENT_75")
	res = append(res, "PERCENT_50")
	res = append(res, "PERCENT_25")
	res = append(res, "PERCENT_00")

	return
}

func (probabilityenum ProbabilityEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "100 %")
	res = append(res, "75 %")
	res = append(res, "50 %")
	res = append(res, "25 %")
	res = append(res, "0 %")

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Weberv2specsStackName:
		res = "weberv2specs"
	case GongsvgStackName:
		res = "gongsvg"
	case GongtreeStackName:
		res = "gongtree"
	case GongtableStackName:
		res = "gongtable"
	case GongsimStackName:
		res = "gongsim"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "weberv2specs":
		*stacksnames = Weberv2specsStackName
		return
	case "gongsvg":
		*stacksnames = GongsvgStackName
		return
	case "gongtree":
		*stacksnames = GongtreeStackName
		return
	case "gongtable":
		*stacksnames = GongtableStackName
		return
	case "gongsim":
		*stacksnames = GongsimStackName
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Weberv2specsStackName":
		*stacksnames = Weberv2specsStackName
	case "GongsvgStackName":
		*stacksnames = GongsvgStackName
	case "GongtreeStackName":
		*stacksnames = GongtreeStackName
	case "GongtableStackName":
		*stacksnames = GongtableStackName
	case "GongsimStackName":
		*stacksnames = GongsimStackName
	default:
		err = errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Weberv2specsStackName:
		res = "Weberv2specsStackName"
	case GongsvgStackName:
		res = "GongsvgStackName"
	case GongtreeStackName:
		res = "GongtreeStackName"
	case GongtableStackName:
		res = "GongtableStackName"
	case GongsimStackName:
		res = "GongsimStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Weberv2specsStackName")
	res = append(res, "GongsvgStackName")
	res = append(res, "GongtreeStackName")
	res = append(res, "GongtableStackName")
	res = append(res, "GongsimStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "weberv2specs")
	res = append(res, "gongsvg")
	res = append(res, "gongtree")
	res = append(res, "gongtable")
	res = append(res, "gongsim")

	return
}

// Utility function for TreeNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treenames TreeNames) ToString() (res string) {

	// migration of former implementation of enum
	switch treenames {
	// insertion code per enum code
	case Sidebar:
		res = "sidebar"
	}
	return
}

func (treenames *TreeNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sidebar":
		*treenames = Sidebar
		return
	default:
		return errUnkownEnum
	}
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Sidebar":
		*treenames = Sidebar
	default:
		err = errUnkownEnum
	}
	return
}

func (treenames *TreeNames) ToCodeString() (res string) {

	switch *treenames {
	// insertion code per enum code
	case Sidebar:
		res = "Sidebar"
	}
	return
}

func (treenames TreeNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Sidebar")

	return
}

func (treenames TreeNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sidebar")

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
