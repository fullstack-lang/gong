// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GongEnumShapeType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (gongenumshapetype GongEnumShapeType) ToInt() (res int) {

	// migration of former implementation of enum
	switch gongenumshapetype {
	// insertion code per enum code
	case Int:
		res = 0
	case String:
		res = 1
	}
	return
}

func (gongenumshapetype *GongEnumShapeType) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*gongenumshapetype = Int
		return
	case 1:
		*gongenumshapetype = String
		return
	default:
		return errUnkownEnum
	}
}

func (gongenumshapetype *GongEnumShapeType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Int":
		*gongenumshapetype = Int
	case "String":
		*gongenumshapetype = String
	default:
		return errUnkownEnum
	}
	return
}

func (gongenumshapetype *GongEnumShapeType) ToCodeString() (res string) {

	switch *gongenumshapetype {
	// insertion code per enum code
	case Int:
		res = "Int"
	case String:
		res = "String"
	}
	return
}

func (gongenumshapetype GongEnumShapeType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Int")
	res = append(res, "String")

	return
}

func (gongenumshapetype GongEnumShapeType) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for MultiplicityType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (multiplicitytype MultiplicityType) ToString() (res string) {

	// migration of former implementation of enum
	switch multiplicitytype {
	// insertion code per enum code
	case ZERO_ONE:
		res = "0..1"
	case ONE:
		res = "1"
	case MANY:
		res = "*"
	}
	return
}

func (multiplicitytype *MultiplicityType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "0..1":
		*multiplicitytype = ZERO_ONE
		return
	case "1":
		*multiplicitytype = ONE
		return
	case "*":
		*multiplicitytype = MANY
		return
	default:
		return errUnkownEnum
	}
}

func (multiplicitytype *MultiplicityType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ZERO_ONE":
		*multiplicitytype = ZERO_ONE
	case "ONE":
		*multiplicitytype = ONE
	case "MANY":
		*multiplicitytype = MANY
	default:
		return errUnkownEnum
	}
	return
}

func (multiplicitytype *MultiplicityType) ToCodeString() (res string) {

	switch *multiplicitytype {
	// insertion code per enum code
	case ZERO_ONE:
		res = "ZERO_ONE"
	case ONE:
		res = "ONE"
	case MANY:
		res = "MANY"
	}
	return
}

func (multiplicitytype MultiplicityType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ZERO_ONE")
	res = append(res, "ONE")
	res = append(res, "MANY")

	return
}

func (multiplicitytype MultiplicityType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "0..1")
	res = append(res, "1")
	res = append(res, "*")

	return
}

// Utility function for NoteShapeLinkType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (noteshapelinktype NoteShapeLinkType) ToString() (res string) {

	// migration of former implementation of enum
	switch noteshapelinktype {
	// insertion code per enum code
	case NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
		res = "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE"
	case NOTE_SHAPE_LINK_TO_GONG_FIELD:
		res = "NOTE_SHAPE_LINK_TO_GONG_FIELD"
	}
	return
}

func (noteshapelinktype *NoteShapeLinkType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE":
		*noteshapelinktype = NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
		return
	case "NOTE_SHAPE_LINK_TO_GONG_FIELD":
		*noteshapelinktype = NOTE_SHAPE_LINK_TO_GONG_FIELD
		return
	default:
		return errUnkownEnum
	}
}

func (noteshapelinktype *NoteShapeLinkType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE":
		*noteshapelinktype = NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
	case "NOTE_SHAPE_LINK_TO_GONG_FIELD":
		*noteshapelinktype = NOTE_SHAPE_LINK_TO_GONG_FIELD
	default:
		return errUnkownEnum
	}
	return
}

func (noteshapelinktype *NoteShapeLinkType) ToCodeString() (res string) {

	switch *noteshapelinktype {
	// insertion code per enum code
	case NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
		res = "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE"
	case NOTE_SHAPE_LINK_TO_GONG_FIELD:
		res = "NOTE_SHAPE_LINK_TO_GONG_FIELD"
	}
	return
}

func (noteshapelinktype NoteShapeLinkType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE")
	res = append(res, "NOTE_SHAPE_LINK_TO_GONG_FIELD")

	return
}

func (noteshapelinktype NoteShapeLinkType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE")
	res = append(res, "NOTE_SHAPE_LINK_TO_GONG_FIELD")

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
		return errUnkownEnum
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

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case SvgStackName:
		res = "svg"
	case GongdocStackName:
		res = "gongdoc"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "svg":
		*stacksnames = SvgStackName
		return
	case "gongdoc":
		*stacksnames = GongdocStackName
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SvgStackName":
		*stacksnames = SvgStackName
	case "GongdocStackName":
		*stacksnames = GongdocStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case SvgStackName:
		res = "SvgStackName"
	case GongdocStackName:
		res = "GongdocStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SvgStackName")
	res = append(res, "GongdocStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "svg")
	res = append(res, "gongdoc")

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
	int | GongEnumShapeType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*GongEnumShapeType
	FromCodeString(input string) (err error)
}

// Last line of the template
