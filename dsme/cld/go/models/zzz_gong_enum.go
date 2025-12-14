// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DominantBaselineType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dominantbaselinetype DominantBaselineType) ToString() (res string) {

	// migration of former implementation of enum
	switch dominantbaselinetype {
	// insertion code per enum code
	case DominantBaselineAuto:
		res = "auto"
	case DominantBaselineAlphabetic:
		res = "alphabetic"
	case DominantBaselineIdeographic:
		res = "ideographic"
	case DominantBaselineHanging:
		res = "hanging"
	case DominantBaselineMathematical:
		res = "mathematical"
	case DominantBaselineCentral:
		res = "central"
	case DominantBaselineMiddle:
		res = "middle"
	case DominantBaselineTextAfterEdge:
		res = "text-after-edge"
	case DominantBaselineTextBeforeEdge:
		res = "text-before-edge"
	case DominantBaselineTextTop:
		res = "text-top"
	case DominantBaselineUseScript:
		res = "use-script"
	case DominantBaselineNoChange:
		res = "no-change"
	case DominantBaselineResetSize:
		res = "reset-size"
	}
	return
}

func (dominantbaselinetype *DominantBaselineType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "auto":
		*dominantbaselinetype = DominantBaselineAuto
		return
	case "alphabetic":
		*dominantbaselinetype = DominantBaselineAlphabetic
		return
	case "ideographic":
		*dominantbaselinetype = DominantBaselineIdeographic
		return
	case "hanging":
		*dominantbaselinetype = DominantBaselineHanging
		return
	case "mathematical":
		*dominantbaselinetype = DominantBaselineMathematical
		return
	case "central":
		*dominantbaselinetype = DominantBaselineCentral
		return
	case "middle":
		*dominantbaselinetype = DominantBaselineMiddle
		return
	case "text-after-edge":
		*dominantbaselinetype = DominantBaselineTextAfterEdge
		return
	case "text-before-edge":
		*dominantbaselinetype = DominantBaselineTextBeforeEdge
		return
	case "text-top":
		*dominantbaselinetype = DominantBaselineTextTop
		return
	case "use-script":
		*dominantbaselinetype = DominantBaselineUseScript
		return
	case "no-change":
		*dominantbaselinetype = DominantBaselineNoChange
		return
	case "reset-size":
		*dominantbaselinetype = DominantBaselineResetSize
		return
	default:
		return errUnkownEnum
	}
}

func (dominantbaselinetype *DominantBaselineType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DominantBaselineAuto":
		*dominantbaselinetype = DominantBaselineAuto
	case "DominantBaselineAlphabetic":
		*dominantbaselinetype = DominantBaselineAlphabetic
	case "DominantBaselineIdeographic":
		*dominantbaselinetype = DominantBaselineIdeographic
	case "DominantBaselineHanging":
		*dominantbaselinetype = DominantBaselineHanging
	case "DominantBaselineMathematical":
		*dominantbaselinetype = DominantBaselineMathematical
	case "DominantBaselineCentral":
		*dominantbaselinetype = DominantBaselineCentral
	case "DominantBaselineMiddle":
		*dominantbaselinetype = DominantBaselineMiddle
	case "DominantBaselineTextAfterEdge":
		*dominantbaselinetype = DominantBaselineTextAfterEdge
	case "DominantBaselineTextBeforeEdge":
		*dominantbaselinetype = DominantBaselineTextBeforeEdge
	case "DominantBaselineTextTop":
		*dominantbaselinetype = DominantBaselineTextTop
	case "DominantBaselineUseScript":
		*dominantbaselinetype = DominantBaselineUseScript
	case "DominantBaselineNoChange":
		*dominantbaselinetype = DominantBaselineNoChange
	case "DominantBaselineResetSize":
		*dominantbaselinetype = DominantBaselineResetSize
	default:
		err = errUnkownEnum
	}
	return
}

func (dominantbaselinetype *DominantBaselineType) ToCodeString() (res string) {

	switch *dominantbaselinetype {
	// insertion code per enum code
	case DominantBaselineAuto:
		res = "DominantBaselineAuto"
	case DominantBaselineAlphabetic:
		res = "DominantBaselineAlphabetic"
	case DominantBaselineIdeographic:
		res = "DominantBaselineIdeographic"
	case DominantBaselineHanging:
		res = "DominantBaselineHanging"
	case DominantBaselineMathematical:
		res = "DominantBaselineMathematical"
	case DominantBaselineCentral:
		res = "DominantBaselineCentral"
	case DominantBaselineMiddle:
		res = "DominantBaselineMiddle"
	case DominantBaselineTextAfterEdge:
		res = "DominantBaselineTextAfterEdge"
	case DominantBaselineTextBeforeEdge:
		res = "DominantBaselineTextBeforeEdge"
	case DominantBaselineTextTop:
		res = "DominantBaselineTextTop"
	case DominantBaselineUseScript:
		res = "DominantBaselineUseScript"
	case DominantBaselineNoChange:
		res = "DominantBaselineNoChange"
	case DominantBaselineResetSize:
		res = "DominantBaselineResetSize"
	}
	return
}

func (dominantbaselinetype DominantBaselineType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DominantBaselineAuto")
	res = append(res, "DominantBaselineAlphabetic")
	res = append(res, "DominantBaselineIdeographic")
	res = append(res, "DominantBaselineHanging")
	res = append(res, "DominantBaselineMathematical")
	res = append(res, "DominantBaselineCentral")
	res = append(res, "DominantBaselineMiddle")
	res = append(res, "DominantBaselineTextAfterEdge")
	res = append(res, "DominantBaselineTextBeforeEdge")
	res = append(res, "DominantBaselineTextTop")
	res = append(res, "DominantBaselineUseScript")
	res = append(res, "DominantBaselineNoChange")
	res = append(res, "DominantBaselineResetSize")

	return
}

func (dominantbaselinetype DominantBaselineType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "auto")
	res = append(res, "alphabetic")
	res = append(res, "ideographic")
	res = append(res, "hanging")
	res = append(res, "mathematical")
	res = append(res, "central")
	res = append(res, "middle")
	res = append(res, "text-after-edge")
	res = append(res, "text-before-edge")
	res = append(res, "text-top")
	res = append(res, "use-script")
	res = append(res, "no-change")
	res = append(res, "reset-size")

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

	return
}

// Utility function for TextAnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (textanchortype TextAnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch textanchortype {
	// insertion code per enum code
	case TEXT_ANCHOR_START:
		res = "start"
	case TEXT_ANCHOR_CENTER:
		res = "middle"
	case TEXT_ANCHOR_END:
		res = "end"
	}
	return
}

func (textanchortype *TextAnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*textanchortype = TEXT_ANCHOR_START
		return
	case "middle":
		*textanchortype = TEXT_ANCHOR_CENTER
		return
	case "end":
		*textanchortype = TEXT_ANCHOR_END
		return
	default:
		return errUnkownEnum
	}
}

func (textanchortype *TextAnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TEXT_ANCHOR_START":
		*textanchortype = TEXT_ANCHOR_START
	case "TEXT_ANCHOR_CENTER":
		*textanchortype = TEXT_ANCHOR_CENTER
	case "TEXT_ANCHOR_END":
		*textanchortype = TEXT_ANCHOR_END
	default:
		err = errUnkownEnum
	}
	return
}

func (textanchortype *TextAnchorType) ToCodeString() (res string) {

	switch *textanchortype {
	// insertion code per enum code
	case TEXT_ANCHOR_START:
		res = "TEXT_ANCHOR_START"
	case TEXT_ANCHOR_CENTER:
		res = "TEXT_ANCHOR_CENTER"
	case TEXT_ANCHOR_END:
		res = "TEXT_ANCHOR_END"
	}
	return
}

func (textanchortype TextAnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TEXT_ANCHOR_START")
	res = append(res, "TEXT_ANCHOR_CENTER")
	res = append(res, "TEXT_ANCHOR_END")

	return
}

func (textanchortype TextAnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "middle")
	res = append(res, "end")

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
	
	FromCodeString(input string) (err error)
}

// Last line of the template
