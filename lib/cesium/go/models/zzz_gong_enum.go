// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for MarkerAnchorEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (markeranchorenum MarkerAnchorEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch markeranchorenum {
	// insertion code per enum code
	case MarkerAnchorCenter:
		res = "center"
	case MarkerAnchorTop:
		res = "top"
	case MarkerAnchorBottom:
		res = "bottom"
	case MarkerAnchorLeft:
		res = "left"
	case MarkerAnchorRight:
		res = "right"
	case MarkerAnchorTopLeft:
		res = "top-left"
	case MarkerAnchorTopRight:
		res = "top-right"
	case MarkerAnchorBottomLeft:
		res = "bottom-left"
	case MarkerAnchorBottomRight:
		res = "bottom-right"
	}
	return
}

func (markeranchorenum *MarkerAnchorEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "center":
		*markeranchorenum = MarkerAnchorCenter
		return
	case "top":
		*markeranchorenum = MarkerAnchorTop
		return
	case "bottom":
		*markeranchorenum = MarkerAnchorBottom
		return
	case "left":
		*markeranchorenum = MarkerAnchorLeft
		return
	case "right":
		*markeranchorenum = MarkerAnchorRight
		return
	case "top-left":
		*markeranchorenum = MarkerAnchorTopLeft
		return
	case "top-right":
		*markeranchorenum = MarkerAnchorTopRight
		return
	case "bottom-left":
		*markeranchorenum = MarkerAnchorBottomLeft
		return
	case "bottom-right":
		*markeranchorenum = MarkerAnchorBottomRight
		return
	default:
		return errUnkownEnum
	}
}

func (markeranchorenum *MarkerAnchorEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MarkerAnchorCenter":
		*markeranchorenum = MarkerAnchorCenter
	case "MarkerAnchorTop":
		*markeranchorenum = MarkerAnchorTop
	case "MarkerAnchorBottom":
		*markeranchorenum = MarkerAnchorBottom
	case "MarkerAnchorLeft":
		*markeranchorenum = MarkerAnchorLeft
	case "MarkerAnchorRight":
		*markeranchorenum = MarkerAnchorRight
	case "MarkerAnchorTopLeft":
		*markeranchorenum = MarkerAnchorTopLeft
	case "MarkerAnchorTopRight":
		*markeranchorenum = MarkerAnchorTopRight
	case "MarkerAnchorBottomLeft":
		*markeranchorenum = MarkerAnchorBottomLeft
	case "MarkerAnchorBottomRight":
		*markeranchorenum = MarkerAnchorBottomRight
	default:
		err = errUnkownEnum
	}
	return
}

func (markeranchorenum *MarkerAnchorEnum) ToCodeString() (res string) {

	switch *markeranchorenum {
	// insertion code per enum code
	case MarkerAnchorCenter:
		res = "MarkerAnchorCenter"
	case MarkerAnchorTop:
		res = "MarkerAnchorTop"
	case MarkerAnchorBottom:
		res = "MarkerAnchorBottom"
	case MarkerAnchorLeft:
		res = "MarkerAnchorLeft"
	case MarkerAnchorRight:
		res = "MarkerAnchorRight"
	case MarkerAnchorTopLeft:
		res = "MarkerAnchorTopLeft"
	case MarkerAnchorTopRight:
		res = "MarkerAnchorTopRight"
	case MarkerAnchorBottomLeft:
		res = "MarkerAnchorBottomLeft"
	case MarkerAnchorBottomRight:
		res = "MarkerAnchorBottomRight"
	}
	return
}

func (markeranchorenum MarkerAnchorEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MarkerAnchorCenter")
	res = append(res, "MarkerAnchorTop")
	res = append(res, "MarkerAnchorBottom")
	res = append(res, "MarkerAnchorLeft")
	res = append(res, "MarkerAnchorRight")
	res = append(res, "MarkerAnchorTopLeft")
	res = append(res, "MarkerAnchorTopRight")
	res = append(res, "MarkerAnchorBottomLeft")
	res = append(res, "MarkerAnchorBottomRight")

	return
}

func (markeranchorenum MarkerAnchorEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "center")
	res = append(res, "top")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "top-left")
	res = append(res, "top-right")
	res = append(res, "bottom-left")
	res = append(res, "bottom-right")

	return
}

// Utility function for MarkerPitchAlignmentEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (markerpitchalignmentenum MarkerPitchAlignmentEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch markerpitchalignmentenum {
	// insertion code per enum code
	case MarkerAlignmentAuto:
		res = "auto"
	case MarkerAlignmentMap:
		res = "map"
	case MarkerAlignmentViewport:
		res = "viewport"
	}
	return
}

func (markerpitchalignmentenum *MarkerPitchAlignmentEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "auto":
		*markerpitchalignmentenum = MarkerAlignmentAuto
		return
	case "map":
		*markerpitchalignmentenum = MarkerAlignmentMap
		return
	case "viewport":
		*markerpitchalignmentenum = MarkerAlignmentViewport
		return
	default:
		return errUnkownEnum
	}
}

func (markerpitchalignmentenum *MarkerPitchAlignmentEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MarkerAlignmentAuto":
		*markerpitchalignmentenum = MarkerAlignmentAuto
	case "MarkerAlignmentMap":
		*markerpitchalignmentenum = MarkerAlignmentMap
	case "MarkerAlignmentViewport":
		*markerpitchalignmentenum = MarkerAlignmentViewport
	default:
		err = errUnkownEnum
	}
	return
}

func (markerpitchalignmentenum *MarkerPitchAlignmentEnum) ToCodeString() (res string) {

	switch *markerpitchalignmentenum {
	// insertion code per enum code
	case MarkerAlignmentAuto:
		res = "MarkerAlignmentAuto"
	case MarkerAlignmentMap:
		res = "MarkerAlignmentMap"
	case MarkerAlignmentViewport:
		res = "MarkerAlignmentViewport"
	}
	return
}

func (markerpitchalignmentenum MarkerPitchAlignmentEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MarkerAlignmentAuto")
	res = append(res, "MarkerAlignmentMap")
	res = append(res, "MarkerAlignmentViewport")

	return
}

func (markerpitchalignmentenum MarkerPitchAlignmentEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "auto")
	res = append(res, "map")
	res = append(res, "viewport")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type GongEnumStringPointer interface {
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumStringField = GongEnumStringPointer

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type GongEnumIntPointer interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumIntField = GongEnumIntPointer

// Last line of the template
