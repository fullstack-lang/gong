// generated code - do not edit
package barline

// insertion point of enum utility functions
// Utility function for EnumBackwardForward
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbackwardforward EnumBackwardForward) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbackwardforward {
	// insertion code per enum code
	case EnumBackwardForwardBackward:
		res = "backward"
	case EnumBackwardForwardForward:
		res = "forward"
	}
	return
}

func (enumbackwardforward *EnumBackwardForward) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "backward":
		*enumbackwardforward = EnumBackwardForwardBackward
		return
	case "forward":
		*enumbackwardforward = EnumBackwardForwardForward
		return
	default:
		return errUnkownEnum
	}
}

func (enumbackwardforward *EnumBackwardForward) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBackwardForwardBackward":
		*enumbackwardforward = EnumBackwardForwardBackward
	case "EnumBackwardForwardForward":
		*enumbackwardforward = EnumBackwardForwardForward
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbackwardforward *EnumBackwardForward) ToCodeString() (res string) {

	switch *enumbackwardforward {
	// insertion code per enum code
	case EnumBackwardForwardBackward:
		res = "EnumBackwardForwardBackward"
	case EnumBackwardForwardForward:
		res = "EnumBackwardForwardForward"
	}
	return
}

func (enumbackwardforward EnumBackwardForward) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBackwardForwardBackward")
	res = append(res, "EnumBackwardForwardForward")

	return
}

func (enumbackwardforward EnumBackwardForward) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "backward")
	res = append(res, "forward")

	return
}

// Utility function for EnumBarStyle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbarstyle EnumBarStyle) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbarstyle {
	// insertion code per enum code
	case EnumBarStyleRegular:
		res = "regular"
	case EnumBarStyleDotted:
		res = "dotted"
	case EnumBarStyleDashed:
		res = "dashed"
	case EnumBarStyleHeavy:
		res = "heavy"
	case EnumBarStyleLightLight:
		res = "light-light"
	case EnumBarStyleLightHeavy:
		res = "light-heavy"
	case EnumBarStyleHeavyLight:
		res = "heavy-light"
	case EnumBarStyleHeavyHeavy:
		res = "heavy-heavy"
	case EnumBarStyleTick:
		res = "tick"
	case EnumBarStyleShort:
		res = "short"
	case EnumBarStyleNone:
		res = "none"
	}
	return
}

func (enumbarstyle *EnumBarStyle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "regular":
		*enumbarstyle = EnumBarStyleRegular
		return
	case "dotted":
		*enumbarstyle = EnumBarStyleDotted
		return
	case "dashed":
		*enumbarstyle = EnumBarStyleDashed
		return
	case "heavy":
		*enumbarstyle = EnumBarStyleHeavy
		return
	case "light-light":
		*enumbarstyle = EnumBarStyleLightLight
		return
	case "light-heavy":
		*enumbarstyle = EnumBarStyleLightHeavy
		return
	case "heavy-light":
		*enumbarstyle = EnumBarStyleHeavyLight
		return
	case "heavy-heavy":
		*enumbarstyle = EnumBarStyleHeavyHeavy
		return
	case "tick":
		*enumbarstyle = EnumBarStyleTick
		return
	case "short":
		*enumbarstyle = EnumBarStyleShort
		return
	case "none":
		*enumbarstyle = EnumBarStyleNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumbarstyle *EnumBarStyle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBarStyleRegular":
		*enumbarstyle = EnumBarStyleRegular
	case "EnumBarStyleDotted":
		*enumbarstyle = EnumBarStyleDotted
	case "EnumBarStyleDashed":
		*enumbarstyle = EnumBarStyleDashed
	case "EnumBarStyleHeavy":
		*enumbarstyle = EnumBarStyleHeavy
	case "EnumBarStyleLightLight":
		*enumbarstyle = EnumBarStyleLightLight
	case "EnumBarStyleLightHeavy":
		*enumbarstyle = EnumBarStyleLightHeavy
	case "EnumBarStyleHeavyLight":
		*enumbarstyle = EnumBarStyleHeavyLight
	case "EnumBarStyleHeavyHeavy":
		*enumbarstyle = EnumBarStyleHeavyHeavy
	case "EnumBarStyleTick":
		*enumbarstyle = EnumBarStyleTick
	case "EnumBarStyleShort":
		*enumbarstyle = EnumBarStyleShort
	case "EnumBarStyleNone":
		*enumbarstyle = EnumBarStyleNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbarstyle *EnumBarStyle) ToCodeString() (res string) {

	switch *enumbarstyle {
	// insertion code per enum code
	case EnumBarStyleRegular:
		res = "EnumBarStyleRegular"
	case EnumBarStyleDotted:
		res = "EnumBarStyleDotted"
	case EnumBarStyleDashed:
		res = "EnumBarStyleDashed"
	case EnumBarStyleHeavy:
		res = "EnumBarStyleHeavy"
	case EnumBarStyleLightLight:
		res = "EnumBarStyleLightLight"
	case EnumBarStyleLightHeavy:
		res = "EnumBarStyleLightHeavy"
	case EnumBarStyleHeavyLight:
		res = "EnumBarStyleHeavyLight"
	case EnumBarStyleHeavyHeavy:
		res = "EnumBarStyleHeavyHeavy"
	case EnumBarStyleTick:
		res = "EnumBarStyleTick"
	case EnumBarStyleShort:
		res = "EnumBarStyleShort"
	case EnumBarStyleNone:
		res = "EnumBarStyleNone"
	}
	return
}

func (enumbarstyle EnumBarStyle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBarStyleRegular")
	res = append(res, "EnumBarStyleDotted")
	res = append(res, "EnumBarStyleDashed")
	res = append(res, "EnumBarStyleHeavy")
	res = append(res, "EnumBarStyleLightLight")
	res = append(res, "EnumBarStyleLightHeavy")
	res = append(res, "EnumBarStyleHeavyLight")
	res = append(res, "EnumBarStyleHeavyHeavy")
	res = append(res, "EnumBarStyleTick")
	res = append(res, "EnumBarStyleShort")
	res = append(res, "EnumBarStyleNone")

	return
}

func (enumbarstyle EnumBarStyle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "regular")
	res = append(res, "dotted")
	res = append(res, "dashed")
	res = append(res, "heavy")
	res = append(res, "light-light")
	res = append(res, "light-heavy")
	res = append(res, "heavy-light")
	res = append(res, "heavy-heavy")
	res = append(res, "tick")
	res = append(res, "short")
	res = append(res, "none")

	return
}

// Utility function for EnumRightLeftMiddle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumrightleftmiddle EnumRightLeftMiddle) ToString() (res string) {

	// migration of former implementation of enum
	switch enumrightleftmiddle {
	// insertion code per enum code
	case EnumRightLeftMiddleRight:
		res = "right"
	case EnumRightLeftMiddleLeft:
		res = "left"
	case EnumRightLeftMiddleMiddle:
		res = "middle"
	}
	return
}

func (enumrightleftmiddle *EnumRightLeftMiddle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enumrightleftmiddle = EnumRightLeftMiddleRight
		return
	case "left":
		*enumrightleftmiddle = EnumRightLeftMiddleLeft
		return
	case "middle":
		*enumrightleftmiddle = EnumRightLeftMiddleMiddle
		return
	default:
		return errUnkownEnum
	}
}

func (enumrightleftmiddle *EnumRightLeftMiddle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumRightLeftMiddleRight":
		*enumrightleftmiddle = EnumRightLeftMiddleRight
	case "EnumRightLeftMiddleLeft":
		*enumrightleftmiddle = EnumRightLeftMiddleLeft
	case "EnumRightLeftMiddleMiddle":
		*enumrightleftmiddle = EnumRightLeftMiddleMiddle
	default:
		err = errUnkownEnum
	}
	return
}

func (enumrightleftmiddle *EnumRightLeftMiddle) ToCodeString() (res string) {

	switch *enumrightleftmiddle {
	// insertion code per enum code
	case EnumRightLeftMiddleRight:
		res = "EnumRightLeftMiddleRight"
	case EnumRightLeftMiddleLeft:
		res = "EnumRightLeftMiddleLeft"
	case EnumRightLeftMiddleMiddle:
		res = "EnumRightLeftMiddleMiddle"
	}
	return
}

func (enumrightleftmiddle EnumRightLeftMiddle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumRightLeftMiddleRight")
	res = append(res, "EnumRightLeftMiddleLeft")
	res = append(res, "EnumRightLeftMiddleMiddle")

	return
}

func (enumrightleftmiddle EnumRightLeftMiddle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "left")
	res = append(res, "middle")

	return
}

// Utility function for EnumStartStopDiscontinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartstopdiscontinue EnumStartStopDiscontinue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartstopdiscontinue {
	// insertion code per enum code
	case EnumStartStopDiscontinueStart:
		res = "start"
	case EnumStartStopDiscontinueStop:
		res = "stop"
	case EnumStartStopDiscontinueDiscontinue:
		res = "discontinue"
	}
	return
}

func (enumstartstopdiscontinue *EnumStartStopDiscontinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueStart
		return
	case "stop":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueStop
		return
	case "discontinue":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueDiscontinue
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartstopdiscontinue *EnumStartStopDiscontinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartStopDiscontinueStart":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueStart
	case "EnumStartStopDiscontinueStop":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueStop
	case "EnumStartStopDiscontinueDiscontinue":
		*enumstartstopdiscontinue = EnumStartStopDiscontinueDiscontinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartstopdiscontinue *EnumStartStopDiscontinue) ToCodeString() (res string) {

	switch *enumstartstopdiscontinue {
	// insertion code per enum code
	case EnumStartStopDiscontinueStart:
		res = "EnumStartStopDiscontinueStart"
	case EnumStartStopDiscontinueStop:
		res = "EnumStartStopDiscontinueStop"
	case EnumStartStopDiscontinueDiscontinue:
		res = "EnumStartStopDiscontinueDiscontinue"
	}
	return
}

func (enumstartstopdiscontinue EnumStartStopDiscontinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartStopDiscontinueStart")
	res = append(res, "EnumStartStopDiscontinueStop")
	res = append(res, "EnumStartStopDiscontinueDiscontinue")

	return
}

func (enumstartstopdiscontinue EnumStartStopDiscontinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "discontinue")

	return
}

// Utility function for EnumSystemRelation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsystemrelation EnumSystemRelation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsystemrelation {
	// insertion code per enum code
	case EnumSystemRelationOnlyTop:
		res = "only-top"
	case EnumSystemRelationAlsoTop:
		res = "also-top"
	case EnumSystemRelationNone:
		res = "none"
	}
	return
}

func (enumsystemrelation *EnumSystemRelation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "only-top":
		*enumsystemrelation = EnumSystemRelationOnlyTop
		return
	case "also-top":
		*enumsystemrelation = EnumSystemRelationAlsoTop
		return
	case "none":
		*enumsystemrelation = EnumSystemRelationNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumsystemrelation *EnumSystemRelation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSystemRelationOnlyTop":
		*enumsystemrelation = EnumSystemRelationOnlyTop
	case "EnumSystemRelationAlsoTop":
		*enumsystemrelation = EnumSystemRelationAlsoTop
	case "EnumSystemRelationNone":
		*enumsystemrelation = EnumSystemRelationNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsystemrelation *EnumSystemRelation) ToCodeString() (res string) {

	switch *enumsystemrelation {
	// insertion code per enum code
	case EnumSystemRelationOnlyTop:
		res = "EnumSystemRelationOnlyTop"
	case EnumSystemRelationAlsoTop:
		res = "EnumSystemRelationAlsoTop"
	case EnumSystemRelationNone:
		res = "EnumSystemRelationNone"
	}
	return
}

func (enumsystemrelation EnumSystemRelation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSystemRelationOnlyTop")
	res = append(res, "EnumSystemRelationAlsoTop")
	res = append(res, "EnumSystemRelationNone")

	return
}

func (enumsystemrelation EnumSystemRelation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "only-top")
	res = append(res, "also-top")
	res = append(res, "none")

	return
}

// Utility function for EnumSystemRelationNumber
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsystemrelationnumber EnumSystemRelationNumber) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsystemrelationnumber {
	// insertion code per enum code
	case EnumSystemRelationNumberOnlyTop:
		res = "only-top"
	case EnumSystemRelationNumberOnlyBottom:
		res = "only-bottom"
	case EnumSystemRelationNumberAlsoTop:
		res = "also-top"
	case EnumSystemRelationNumberAlsoBottom:
		res = "also-bottom"
	case EnumSystemRelationNumberNone:
		res = "none"
	}
	return
}

func (enumsystemrelationnumber *EnumSystemRelationNumber) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "only-top":
		*enumsystemrelationnumber = EnumSystemRelationNumberOnlyTop
		return
	case "only-bottom":
		*enumsystemrelationnumber = EnumSystemRelationNumberOnlyBottom
		return
	case "also-top":
		*enumsystemrelationnumber = EnumSystemRelationNumberAlsoTop
		return
	case "also-bottom":
		*enumsystemrelationnumber = EnumSystemRelationNumberAlsoBottom
		return
	case "none":
		*enumsystemrelationnumber = EnumSystemRelationNumberNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumsystemrelationnumber *EnumSystemRelationNumber) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSystemRelationNumberOnlyTop":
		*enumsystemrelationnumber = EnumSystemRelationNumberOnlyTop
	case "EnumSystemRelationNumberOnlyBottom":
		*enumsystemrelationnumber = EnumSystemRelationNumberOnlyBottom
	case "EnumSystemRelationNumberAlsoTop":
		*enumsystemrelationnumber = EnumSystemRelationNumberAlsoTop
	case "EnumSystemRelationNumberAlsoBottom":
		*enumsystemrelationnumber = EnumSystemRelationNumberAlsoBottom
	case "EnumSystemRelationNumberNone":
		*enumsystemrelationnumber = EnumSystemRelationNumberNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsystemrelationnumber *EnumSystemRelationNumber) ToCodeString() (res string) {

	switch *enumsystemrelationnumber {
	// insertion code per enum code
	case EnumSystemRelationNumberOnlyTop:
		res = "EnumSystemRelationNumberOnlyTop"
	case EnumSystemRelationNumberOnlyBottom:
		res = "EnumSystemRelationNumberOnlyBottom"
	case EnumSystemRelationNumberAlsoTop:
		res = "EnumSystemRelationNumberAlsoTop"
	case EnumSystemRelationNumberAlsoBottom:
		res = "EnumSystemRelationNumberAlsoBottom"
	case EnumSystemRelationNumberNone:
		res = "EnumSystemRelationNumberNone"
	}
	return
}

func (enumsystemrelationnumber EnumSystemRelationNumber) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSystemRelationNumberOnlyTop")
	res = append(res, "EnumSystemRelationNumberOnlyBottom")
	res = append(res, "EnumSystemRelationNumberAlsoTop")
	res = append(res, "EnumSystemRelationNumberAlsoBottom")
	res = append(res, "EnumSystemRelationNumberNone")

	return
}

func (enumsystemrelationnumber EnumSystemRelationNumber) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "only-top")
	res = append(res, "only-bottom")
	res = append(res, "also-top")
	res = append(res, "also-bottom")
	res = append(res, "none")

	return
}

// Utility function for EnumWinged
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumwinged EnumWinged) ToString() (res string) {

	// migration of former implementation of enum
	switch enumwinged {
	// insertion code per enum code
	case EnumWingedNone:
		res = "none"
	case EnumWingedStraight:
		res = "straight"
	case EnumWingedCurved:
		res = "curved"
	case EnumWingedDoubleStraight:
		res = "double-straight"
	case EnumWingedDoubleCurved:
		res = "double-curved"
	}
	return
}

func (enumwinged *EnumWinged) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enumwinged = EnumWingedNone
		return
	case "straight":
		*enumwinged = EnumWingedStraight
		return
	case "curved":
		*enumwinged = EnumWingedCurved
		return
	case "double-straight":
		*enumwinged = EnumWingedDoubleStraight
		return
	case "double-curved":
		*enumwinged = EnumWingedDoubleCurved
		return
	default:
		return errUnkownEnum
	}
}

func (enumwinged *EnumWinged) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumWingedNone":
		*enumwinged = EnumWingedNone
	case "EnumWingedStraight":
		*enumwinged = EnumWingedStraight
	case "EnumWingedCurved":
		*enumwinged = EnumWingedCurved
	case "EnumWingedDoubleStraight":
		*enumwinged = EnumWingedDoubleStraight
	case "EnumWingedDoubleCurved":
		*enumwinged = EnumWingedDoubleCurved
	default:
		err = errUnkownEnum
	}
	return
}

func (enumwinged *EnumWinged) ToCodeString() (res string) {

	switch *enumwinged {
	// insertion code per enum code
	case EnumWingedNone:
		res = "EnumWingedNone"
	case EnumWingedStraight:
		res = "EnumWingedStraight"
	case EnumWingedCurved:
		res = "EnumWingedCurved"
	case EnumWingedDoubleStraight:
		res = "EnumWingedDoubleStraight"
	case EnumWingedDoubleCurved:
		res = "EnumWingedDoubleCurved"
	}
	return
}

func (enumwinged EnumWinged) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumWingedNone")
	res = append(res, "EnumWingedStraight")
	res = append(res, "EnumWingedCurved")
	res = append(res, "EnumWingedDoubleStraight")
	res = append(res, "EnumWingedDoubleCurved")

	return
}

func (enumwinged EnumWinged) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "straight")
	res = append(res, "curved")
	res = append(res, "double-straight")
	res = append(res, "double-curved")

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
