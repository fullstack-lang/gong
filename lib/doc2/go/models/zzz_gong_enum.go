// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for ButtonType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (buttontype ButtonType) ToInt() (res int) {

	// migration of former implementation of enum
	switch buttontype {
	// insertion code per enum code
	case DUPLICATE:
		res = 0
	case EDIT_CANCEL:
		res = 1
	case EDIT:
		res = 2
	case REMOVE:
		res = 3
	case RENAME_CANCEL:
		res = 4
	case RENAME:
		res = 5
	case SAVE:
		res = 6
	}
	return
}

func (buttontype *ButtonType) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*buttontype = DUPLICATE
		return
	case 1:
		*buttontype = EDIT_CANCEL
		return
	case 2:
		*buttontype = EDIT
		return
	case 3:
		*buttontype = REMOVE
		return
	case 4:
		*buttontype = RENAME_CANCEL
		return
	case 5:
		*buttontype = RENAME
		return
	case 6:
		*buttontype = SAVE
		return
	default:
		return errUnkownEnum
	}
}

func (buttontype *ButtonType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DUPLICATE":
		*buttontype = DUPLICATE
	case "EDIT_CANCEL":
		*buttontype = EDIT_CANCEL
	case "EDIT":
		*buttontype = EDIT
	case "REMOVE":
		*buttontype = REMOVE
	case "RENAME_CANCEL":
		*buttontype = RENAME_CANCEL
	case "RENAME":
		*buttontype = RENAME
	case "SAVE":
		*buttontype = SAVE
	default:
		return errUnkownEnum
	}
	return
}

func (buttontype *ButtonType) ToCodeString() (res string) {

	switch *buttontype {
	// insertion code per enum code
	case DUPLICATE:
		res = "DUPLICATE"
	case EDIT_CANCEL:
		res = "EDIT_CANCEL"
	case EDIT:
		res = "EDIT"
	case REMOVE:
		res = "REMOVE"
	case RENAME_CANCEL:
		res = "RENAME_CANCEL"
	case RENAME:
		res = "RENAME"
	case SAVE:
		res = "SAVE"
	}
	return
}

func (buttontype ButtonType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DUPLICATE")
	res = append(res, "EDIT_CANCEL")
	res = append(res, "EDIT")
	res = append(res, "REMOVE")
	res = append(res, "RENAME_CANCEL")
	res = append(res, "RENAME")
	res = append(res, "SAVE")

	return
}

func (buttontype ButtonType) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)
	res = append(res, 4)
	res = append(res, 5)
	res = append(res, 6)

	return
}

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

// Utility function for TreeNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treenames TreeNames) ToString() (res string) {

	// migration of former implementation of enum
	switch treenames {
	// insertion code per enum code
	case Portfolio:
		res = "portfolio"
	case Model:
		res = "model"
	}
	return
}

func (treenames *TreeNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "portfolio":
		*treenames = Portfolio
		return
	case "model":
		*treenames = Model
		return
	default:
		return errUnkownEnum
	}
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Portfolio":
		*treenames = Portfolio
	case "Model":
		*treenames = Model
	default:
		return errUnkownEnum
	}
	return
}

func (treenames *TreeNames) ToCodeString() (res string) {

	switch *treenames {
	// insertion code per enum code
	case Portfolio:
		res = "Portfolio"
	case Model:
		res = "Model"
	}
	return
}

func (treenames TreeNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Portfolio")
	res = append(res, "Model")

	return
}

func (treenames TreeNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "portfolio")
	res = append(res, "model")

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
	int | ButtonType | GongEnumShapeType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*ButtonType | *GongEnumShapeType
	FromCodeString(input string) (err error)
}

// Last line of the template
