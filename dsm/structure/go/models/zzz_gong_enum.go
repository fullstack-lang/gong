// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DataFlowDirection
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dataflowdirection DataFlowDirection) ToString() (res string) {

	// migration of former implementation of enum
	switch dataflowdirection {
	// insertion code per enum code
	case Unspecified:
		res = "Unspecified"
	case Forward:
		res = "Forward"
	case Backward:
		res = "Backward"
	case BothWays:
		res = "BothWays"
	}
	return
}

func (dataflowdirection *DataFlowDirection) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Unspecified":
		*dataflowdirection = Unspecified
		return
	case "Forward":
		*dataflowdirection = Forward
		return
	case "Backward":
		*dataflowdirection = Backward
		return
	case "BothWays":
		*dataflowdirection = BothWays
		return
	default:
		return errUnkownEnum
	}
}

func (dataflowdirection *DataFlowDirection) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Unspecified":
		*dataflowdirection = Unspecified
	case "Forward":
		*dataflowdirection = Forward
	case "Backward":
		*dataflowdirection = Backward
	case "BothWays":
		*dataflowdirection = BothWays
	default:
		err = errUnkownEnum
	}
	return
}

func (dataflowdirection *DataFlowDirection) ToCodeString() (res string) {

	switch *dataflowdirection {
	// insertion code per enum code
	case Unspecified:
		res = "Unspecified"
	case Forward:
		res = "Forward"
	case Backward:
		res = "Backward"
	case BothWays:
		res = "BothWays"
	}
	return
}

func (dataflowdirection DataFlowDirection) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Unspecified")
	res = append(res, "Forward")
	res = append(res, "Backward")
	res = append(res, "BothWays")

	return
}

func (dataflowdirection DataFlowDirection) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Unspecified")
	res = append(res, "Forward")
	res = append(res, "Backward")
	res = append(res, "BothWays")

	return
}

// Utility function for DataFlowType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dataflowtype DataFlowType) ToString() (res string) {

	// migration of former implementation of enum
	switch dataflowtype {
	// insertion code per enum code
	case DataFlow_Port2Port:
		res = "Port2Port"
	case DataFlow_ExternalPart2Port:
		res = "ExternalPart2Port"
	case DataFlow_Port2ExternalPart:
		res = "Port2ExternalPart"
	}
	return
}

func (dataflowtype *DataFlowType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Port2Port":
		*dataflowtype = DataFlow_Port2Port
		return
	case "ExternalPart2Port":
		*dataflowtype = DataFlow_ExternalPart2Port
		return
	case "Port2ExternalPart":
		*dataflowtype = DataFlow_Port2ExternalPart
		return
	default:
		return errUnkownEnum
	}
}

func (dataflowtype *DataFlowType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DataFlow_Port2Port":
		*dataflowtype = DataFlow_Port2Port
	case "DataFlow_ExternalPart2Port":
		*dataflowtype = DataFlow_ExternalPart2Port
	case "DataFlow_Port2ExternalPart":
		*dataflowtype = DataFlow_Port2ExternalPart
	default:
		err = errUnkownEnum
	}
	return
}

func (dataflowtype *DataFlowType) ToCodeString() (res string) {

	switch *dataflowtype {
	// insertion code per enum code
	case DataFlow_Port2Port:
		res = "DataFlow_Port2Port"
	case DataFlow_ExternalPart2Port:
		res = "DataFlow_ExternalPart2Port"
	case DataFlow_Port2ExternalPart:
		res = "DataFlow_Port2ExternalPart"
	}
	return
}

func (dataflowtype DataFlowType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DataFlow_Port2Port")
	res = append(res, "DataFlow_ExternalPart2Port")
	res = append(res, "DataFlow_Port2ExternalPart")

	return
}

func (dataflowtype DataFlowType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Port2Port")
	res = append(res, "ExternalPart2Port")
	res = append(res, "Port2ExternalPart")

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

// Utility function for RectAnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (rectanchortype RectAnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch rectanchortype {
	// insertion code per enum code
	case RECT_TOP:
		res = "RECT_TOP"
	case RECT_TOP_LEFT:
		res = "RECT_TOP_LEFT"
	case RECT_TOP_RIGHT:
		res = "RECT_TOP_RIGHT"
	case RECT_BOTTOM:
		res = "RECT_BOTTOM"
	case RECT_BOTTOM_LEFT:
		res = "RECT_BOTTOM_LEFT"
	case RECT_BOTTOM_LEFT_LEFT:
		res = "RECT_BOTTOM_LEFT_LEFT"
	case RECT_BOTTOM_BOTTOM_LEFT:
		res = "RECT_BOTTOM_BOTTOM_LEFT"
	case RECT_BOTTOM_RIGHT:
		res = "RECT_BOTTOM_RIGHT"
	case RECT_BOTTOM_INSIDE_RIGHT:
		res = "RECT_BOTTOM_INSIDE_RIGHT"
	case RECT_LEFT:
		res = "RECT_LEFT"
	case RECT_RIGHT:
		res = "RECT_RIGHT"
	case RECT_CENTER:
		res = "RECT_CENTER"
	case RECT_CENTER_MIDDLE:
		res = "RECT_CENTER_MIDDLE"
	}
	return
}

func (rectanchortype *RectAnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RECT_TOP":
		*rectanchortype = RECT_TOP
		return
	case "RECT_TOP_LEFT":
		*rectanchortype = RECT_TOP_LEFT
		return
	case "RECT_TOP_RIGHT":
		*rectanchortype = RECT_TOP_RIGHT
		return
	case "RECT_BOTTOM":
		*rectanchortype = RECT_BOTTOM
		return
	case "RECT_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT
		return
	case "RECT_BOTTOM_LEFT_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT_LEFT
		return
	case "RECT_BOTTOM_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_BOTTOM_LEFT
		return
	case "RECT_BOTTOM_RIGHT":
		*rectanchortype = RECT_BOTTOM_RIGHT
		return
	case "RECT_BOTTOM_INSIDE_RIGHT":
		*rectanchortype = RECT_BOTTOM_INSIDE_RIGHT
		return
	case "RECT_LEFT":
		*rectanchortype = RECT_LEFT
		return
	case "RECT_RIGHT":
		*rectanchortype = RECT_RIGHT
		return
	case "RECT_CENTER":
		*rectanchortype = RECT_CENTER
		return
	case "RECT_CENTER_MIDDLE":
		*rectanchortype = RECT_CENTER_MIDDLE
		return
	default:
		return errUnkownEnum
	}
}

func (rectanchortype *RectAnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RECT_TOP":
		*rectanchortype = RECT_TOP
	case "RECT_TOP_LEFT":
		*rectanchortype = RECT_TOP_LEFT
	case "RECT_TOP_RIGHT":
		*rectanchortype = RECT_TOP_RIGHT
	case "RECT_BOTTOM":
		*rectanchortype = RECT_BOTTOM
	case "RECT_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT
	case "RECT_BOTTOM_LEFT_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT_LEFT
	case "RECT_BOTTOM_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_BOTTOM_LEFT
	case "RECT_BOTTOM_RIGHT":
		*rectanchortype = RECT_BOTTOM_RIGHT
	case "RECT_BOTTOM_INSIDE_RIGHT":
		*rectanchortype = RECT_BOTTOM_INSIDE_RIGHT
	case "RECT_LEFT":
		*rectanchortype = RECT_LEFT
	case "RECT_RIGHT":
		*rectanchortype = RECT_RIGHT
	case "RECT_CENTER":
		*rectanchortype = RECT_CENTER
	case "RECT_CENTER_MIDDLE":
		*rectanchortype = RECT_CENTER_MIDDLE
	default:
		err = errUnkownEnum
	}
	return
}

func (rectanchortype *RectAnchorType) ToCodeString() (res string) {

	switch *rectanchortype {
	// insertion code per enum code
	case RECT_TOP:
		res = "RECT_TOP"
	case RECT_TOP_LEFT:
		res = "RECT_TOP_LEFT"
	case RECT_TOP_RIGHT:
		res = "RECT_TOP_RIGHT"
	case RECT_BOTTOM:
		res = "RECT_BOTTOM"
	case RECT_BOTTOM_LEFT:
		res = "RECT_BOTTOM_LEFT"
	case RECT_BOTTOM_LEFT_LEFT:
		res = "RECT_BOTTOM_LEFT_LEFT"
	case RECT_BOTTOM_BOTTOM_LEFT:
		res = "RECT_BOTTOM_BOTTOM_LEFT"
	case RECT_BOTTOM_RIGHT:
		res = "RECT_BOTTOM_RIGHT"
	case RECT_BOTTOM_INSIDE_RIGHT:
		res = "RECT_BOTTOM_INSIDE_RIGHT"
	case RECT_LEFT:
		res = "RECT_LEFT"
	case RECT_RIGHT:
		res = "RECT_RIGHT"
	case RECT_CENTER:
		res = "RECT_CENTER"
	case RECT_CENTER_MIDDLE:
		res = "RECT_CENTER_MIDDLE"
	}
	return
}

func (rectanchortype RectAnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RECT_TOP")
	res = append(res, "RECT_TOP_LEFT")
	res = append(res, "RECT_TOP_RIGHT")
	res = append(res, "RECT_BOTTOM")
	res = append(res, "RECT_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_LEFT_LEFT")
	res = append(res, "RECT_BOTTOM_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_RIGHT")
	res = append(res, "RECT_BOTTOM_INSIDE_RIGHT")
	res = append(res, "RECT_LEFT")
	res = append(res, "RECT_RIGHT")
	res = append(res, "RECT_CENTER")
	res = append(res, "RECT_CENTER_MIDDLE")

	return
}

func (rectanchortype RectAnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RECT_TOP")
	res = append(res, "RECT_TOP_LEFT")
	res = append(res, "RECT_TOP_RIGHT")
	res = append(res, "RECT_BOTTOM")
	res = append(res, "RECT_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_LEFT_LEFT")
	res = append(res, "RECT_BOTTOM_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_RIGHT")
	res = append(res, "RECT_BOTTOM_INSIDE_RIGHT")
	res = append(res, "RECT_LEFT")
	res = append(res, "RECT_RIGHT")
	res = append(res, "RECT_CENTER")
	res = append(res, "RECT_CENTER_MIDDLE")

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
