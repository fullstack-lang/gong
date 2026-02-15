// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Criticality
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (criticality Criticality) ToString() (res string) {

	// migration of former implementation of enum
	switch criticality {
	// insertion code per enum code
	case CriticalityCritical:
		res = "Critical"
	case CriticalityDefault:
		res = "Default"
	}
	return
}

func (criticality *Criticality) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Critical":
		*criticality = CriticalityCritical
		return
	case "Default":
		*criticality = CriticalityDefault
		return
	default:
		return errUnkownEnum
	}
}

func (criticality *Criticality) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CriticalityCritical":
		*criticality = CriticalityCritical
	case "CriticalityDefault":
		*criticality = CriticalityDefault
	default:
		err = errUnkownEnum
	}
	return
}

func (criticality *Criticality) ToCodeString() (res string) {

	switch *criticality {
	// insertion code per enum code
	case CriticalityCritical:
		res = "CriticalityCritical"
	case CriticalityDefault:
		res = "CriticalityDefault"
	}
	return
}

func (criticality Criticality) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CriticalityCritical")
	res = append(res, "CriticalityDefault")

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
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

// Last line of the template
