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
		err = errUnkownEnum
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

// Utility function for Criticality
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (criticality Criticality) ToString() (res string) {

	// migration of former implementation of enum
	switch criticality {
	// insertion code per enum code
	case DoActionCritical:
		res = "Critical"
	case DoActionDefault:
		res = "Default"
	}
	return
}

func (criticality *Criticality) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Critical":
		*criticality = DoActionCritical
		return
	case "Default":
		*criticality = DoActionDefault
		return
	default:
		return errUnkownEnum
	}
}

func (criticality *Criticality) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DoActionCritical":
		*criticality = DoActionCritical
	case "DoActionDefault":
		*criticality = DoActionDefault
	default:
		err = errUnkownEnum
	}
	return
}

func (criticality *Criticality) ToCodeString() (res string) {

	switch *criticality {
	// insertion code per enum code
	case DoActionCritical:
		res = "DoActionCritical"
	case DoActionDefault:
		res = "DoActionDefault"
	}
	return
}

func (criticality Criticality) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DoActionCritical")
	res = append(res, "DoActionDefault")

	return
}

func (criticality Criticality) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Critical")
	res = append(res, "Default")

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

// Utility function for TreeNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treenames TreeNames) ToString() (res string) {

	// migration of former implementation of enum
	switch treenames {
	// insertion code per enum code
	case ObjectTreeName:
		res = "Object Tree Name"
	case DiagramTreeName:
		res = "Diagram Tree Name"
	}
	return
}

func (treenames *TreeNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Object Tree Name":
		*treenames = ObjectTreeName
		return
	case "Diagram Tree Name":
		*treenames = DiagramTreeName
		return
	default:
		return errUnkownEnum
	}
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ObjectTreeName":
		*treenames = ObjectTreeName
	case "DiagramTreeName":
		*treenames = DiagramTreeName
	default:
		err = errUnkownEnum
	}
	return
}

func (treenames *TreeNames) ToCodeString() (res string) {

	switch *treenames {
	// insertion code per enum code
	case ObjectTreeName:
		res = "ObjectTreeName"
	case DiagramTreeName:
		res = "DiagramTreeName"
	}
	return
}

func (treenames TreeNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ObjectTreeName")
	res = append(res, "DiagramTreeName")

	return
}

func (treenames TreeNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Object Tree Name")
	res = append(res, "Diagram Tree Name")

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
	int | ButtonType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*ButtonType
	FromCodeString(input string) (err error)
}

// Last line of the template
