// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DataFlowType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dataflowtype DataFlowType) ToString() (res string) {

	// migration of former implementation of enum
	switch dataflowtype {
	// insertion code per enum code
	case DataFlow_Task2Task:
		res = "DataFlow_Task2Task"
	case DataFlow_ExternalParticipant2Task:
		res = "DataFlow_ExternalParticipant2Task"
	case DataFlow_Task2ExternalParticipant:
		res = "DataFlow_Task2ExternalParticipant"
	}
	return
}

func (dataflowtype *DataFlowType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DataFlow_Task2Task":
		*dataflowtype = DataFlow_Task2Task
		return
	case "DataFlow_ExternalParticipant2Task":
		*dataflowtype = DataFlow_ExternalParticipant2Task
		return
	case "DataFlow_Task2ExternalParticipant":
		*dataflowtype = DataFlow_Task2ExternalParticipant
		return
	default:
		return errUnkownEnum
	}
}

func (dataflowtype *DataFlowType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DataFlow_Task2Task":
		*dataflowtype = DataFlow_Task2Task
	case "DataFlow_ExternalParticipant2Task":
		*dataflowtype = DataFlow_ExternalParticipant2Task
	case "DataFlow_Task2ExternalParticipant":
		*dataflowtype = DataFlow_Task2ExternalParticipant
	default:
		err = errUnkownEnum
	}
	return
}

func (dataflowtype *DataFlowType) ToCodeString() (res string) {

	switch *dataflowtype {
	// insertion code per enum code
	case DataFlow_Task2Task:
		res = "DataFlow_Task2Task"
	case DataFlow_ExternalParticipant2Task:
		res = "DataFlow_ExternalParticipant2Task"
	case DataFlow_Task2ExternalParticipant:
		res = "DataFlow_Task2ExternalParticipant"
	}
	return
}

func (dataflowtype DataFlowType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DataFlow_Task2Task")
	res = append(res, "DataFlow_ExternalParticipant2Task")
	res = append(res, "DataFlow_Task2ExternalParticipant")

	return
}

func (dataflowtype DataFlowType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DataFlow_Task2Task")
	res = append(res, "DataFlow_ExternalParticipant2Task")
	res = append(res, "DataFlow_Task2ExternalParticipant")

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
