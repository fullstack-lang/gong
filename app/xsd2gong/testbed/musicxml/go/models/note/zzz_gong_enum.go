// generated code - do not edit
package note

// insertion point of enum utility functions
// Utility function for EnumArrowDirection
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumarrowdirection EnumArrowDirection) ToString() (res string) {

	// migration of former implementation of enum
	switch enumarrowdirection {
	// insertion code per enum code
	case EnumArrowDirectionLeft:
		res = "left"
	case EnumArrowDirectionUp:
		res = "up"
	case EnumArrowDirectionRight:
		res = "right"
	case EnumArrowDirectionDown:
		res = "down"
	case EnumArrowDirectionNorthwest:
		res = "northwest"
	case EnumArrowDirectionNortheast:
		res = "northeast"
	case EnumArrowDirectionSoutheast:
		res = "southeast"
	case EnumArrowDirectionSouthwest:
		res = "southwest"
	case EnumArrowDirectionLeftRight:
		res = "left right"
	case EnumArrowDirectionUpDown:
		res = "up down"
	case EnumArrowDirectionNorthwestSoutheast:
		res = "northwest southeast"
	case EnumArrowDirectionNortheastSouthwest:
		res = "northeast southwest"
	case EnumArrowDirectionOther:
		res = "other"
	}
	return
}

func (enumarrowdirection *EnumArrowDirection) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enumarrowdirection = EnumArrowDirectionLeft
		return
	case "up":
		*enumarrowdirection = EnumArrowDirectionUp
		return
	case "right":
		*enumarrowdirection = EnumArrowDirectionRight
		return
	case "down":
		*enumarrowdirection = EnumArrowDirectionDown
		return
	case "northwest":
		*enumarrowdirection = EnumArrowDirectionNorthwest
		return
	case "northeast":
		*enumarrowdirection = EnumArrowDirectionNortheast
		return
	case "southeast":
		*enumarrowdirection = EnumArrowDirectionSoutheast
		return
	case "southwest":
		*enumarrowdirection = EnumArrowDirectionSouthwest
		return
	case "left right":
		*enumarrowdirection = EnumArrowDirectionLeftRight
		return
	case "up down":
		*enumarrowdirection = EnumArrowDirectionUpDown
		return
	case "northwest southeast":
		*enumarrowdirection = EnumArrowDirectionNorthwestSoutheast
		return
	case "northeast southwest":
		*enumarrowdirection = EnumArrowDirectionNortheastSouthwest
		return
	case "other":
		*enumarrowdirection = EnumArrowDirectionOther
		return
	default:
		return errUnkownEnum
	}
}

func (enumarrowdirection *EnumArrowDirection) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumArrowDirectionLeft":
		*enumarrowdirection = EnumArrowDirectionLeft
	case "EnumArrowDirectionUp":
		*enumarrowdirection = EnumArrowDirectionUp
	case "EnumArrowDirectionRight":
		*enumarrowdirection = EnumArrowDirectionRight
	case "EnumArrowDirectionDown":
		*enumarrowdirection = EnumArrowDirectionDown
	case "EnumArrowDirectionNorthwest":
		*enumarrowdirection = EnumArrowDirectionNorthwest
	case "EnumArrowDirectionNortheast":
		*enumarrowdirection = EnumArrowDirectionNortheast
	case "EnumArrowDirectionSoutheast":
		*enumarrowdirection = EnumArrowDirectionSoutheast
	case "EnumArrowDirectionSouthwest":
		*enumarrowdirection = EnumArrowDirectionSouthwest
	case "EnumArrowDirectionLeftRight":
		*enumarrowdirection = EnumArrowDirectionLeftRight
	case "EnumArrowDirectionUpDown":
		*enumarrowdirection = EnumArrowDirectionUpDown
	case "EnumArrowDirectionNorthwestSoutheast":
		*enumarrowdirection = EnumArrowDirectionNorthwestSoutheast
	case "EnumArrowDirectionNortheastSouthwest":
		*enumarrowdirection = EnumArrowDirectionNortheastSouthwest
	case "EnumArrowDirectionOther":
		*enumarrowdirection = EnumArrowDirectionOther
	default:
		err = errUnkownEnum
	}
	return
}

func (enumarrowdirection *EnumArrowDirection) ToCodeString() (res string) {

	switch *enumarrowdirection {
	// insertion code per enum code
	case EnumArrowDirectionLeft:
		res = "EnumArrowDirectionLeft"
	case EnumArrowDirectionUp:
		res = "EnumArrowDirectionUp"
	case EnumArrowDirectionRight:
		res = "EnumArrowDirectionRight"
	case EnumArrowDirectionDown:
		res = "EnumArrowDirectionDown"
	case EnumArrowDirectionNorthwest:
		res = "EnumArrowDirectionNorthwest"
	case EnumArrowDirectionNortheast:
		res = "EnumArrowDirectionNortheast"
	case EnumArrowDirectionSoutheast:
		res = "EnumArrowDirectionSoutheast"
	case EnumArrowDirectionSouthwest:
		res = "EnumArrowDirectionSouthwest"
	case EnumArrowDirectionLeftRight:
		res = "EnumArrowDirectionLeftRight"
	case EnumArrowDirectionUpDown:
		res = "EnumArrowDirectionUpDown"
	case EnumArrowDirectionNorthwestSoutheast:
		res = "EnumArrowDirectionNorthwestSoutheast"
	case EnumArrowDirectionNortheastSouthwest:
		res = "EnumArrowDirectionNortheastSouthwest"
	case EnumArrowDirectionOther:
		res = "EnumArrowDirectionOther"
	}
	return
}

func (enumarrowdirection EnumArrowDirection) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumArrowDirectionLeft")
	res = append(res, "EnumArrowDirectionUp")
	res = append(res, "EnumArrowDirectionRight")
	res = append(res, "EnumArrowDirectionDown")
	res = append(res, "EnumArrowDirectionNorthwest")
	res = append(res, "EnumArrowDirectionNortheast")
	res = append(res, "EnumArrowDirectionSoutheast")
	res = append(res, "EnumArrowDirectionSouthwest")
	res = append(res, "EnumArrowDirectionLeftRight")
	res = append(res, "EnumArrowDirectionUpDown")
	res = append(res, "EnumArrowDirectionNorthwestSoutheast")
	res = append(res, "EnumArrowDirectionNortheastSouthwest")
	res = append(res, "EnumArrowDirectionOther")

	return
}

func (enumarrowdirection EnumArrowDirection) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "up")
	res = append(res, "right")
	res = append(res, "down")
	res = append(res, "northwest")
	res = append(res, "northeast")
	res = append(res, "southeast")
	res = append(res, "southwest")
	res = append(res, "left right")
	res = append(res, "up down")
	res = append(res, "northwest southeast")
	res = append(res, "northeast southwest")
	res = append(res, "other")

	return
}

// Utility function for EnumArrowStyle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumarrowstyle EnumArrowStyle) ToString() (res string) {

	// migration of former implementation of enum
	switch enumarrowstyle {
	// insertion code per enum code
	case EnumArrowStyleSingle:
		res = "single"
	case EnumArrowStyleDouble:
		res = "double"
	case EnumArrowStyleFilled:
		res = "filled"
	case EnumArrowStyleHollow:
		res = "hollow"
	case EnumArrowStylePaired:
		res = "paired"
	case EnumArrowStyleCombined:
		res = "combined"
	case EnumArrowStyleOther:
		res = "other"
	}
	return
}

func (enumarrowstyle *EnumArrowStyle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*enumarrowstyle = EnumArrowStyleSingle
		return
	case "double":
		*enumarrowstyle = EnumArrowStyleDouble
		return
	case "filled":
		*enumarrowstyle = EnumArrowStyleFilled
		return
	case "hollow":
		*enumarrowstyle = EnumArrowStyleHollow
		return
	case "paired":
		*enumarrowstyle = EnumArrowStylePaired
		return
	case "combined":
		*enumarrowstyle = EnumArrowStyleCombined
		return
	case "other":
		*enumarrowstyle = EnumArrowStyleOther
		return
	default:
		return errUnkownEnum
	}
}

func (enumarrowstyle *EnumArrowStyle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumArrowStyleSingle":
		*enumarrowstyle = EnumArrowStyleSingle
	case "EnumArrowStyleDouble":
		*enumarrowstyle = EnumArrowStyleDouble
	case "EnumArrowStyleFilled":
		*enumarrowstyle = EnumArrowStyleFilled
	case "EnumArrowStyleHollow":
		*enumarrowstyle = EnumArrowStyleHollow
	case "EnumArrowStylePaired":
		*enumarrowstyle = EnumArrowStylePaired
	case "EnumArrowStyleCombined":
		*enumarrowstyle = EnumArrowStyleCombined
	case "EnumArrowStyleOther":
		*enumarrowstyle = EnumArrowStyleOther
	default:
		err = errUnkownEnum
	}
	return
}

func (enumarrowstyle *EnumArrowStyle) ToCodeString() (res string) {

	switch *enumarrowstyle {
	// insertion code per enum code
	case EnumArrowStyleSingle:
		res = "EnumArrowStyleSingle"
	case EnumArrowStyleDouble:
		res = "EnumArrowStyleDouble"
	case EnumArrowStyleFilled:
		res = "EnumArrowStyleFilled"
	case EnumArrowStyleHollow:
		res = "EnumArrowStyleHollow"
	case EnumArrowStylePaired:
		res = "EnumArrowStylePaired"
	case EnumArrowStyleCombined:
		res = "EnumArrowStyleCombined"
	case EnumArrowStyleOther:
		res = "EnumArrowStyleOther"
	}
	return
}

func (enumarrowstyle EnumArrowStyle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumArrowStyleSingle")
	res = append(res, "EnumArrowStyleDouble")
	res = append(res, "EnumArrowStyleFilled")
	res = append(res, "EnumArrowStyleHollow")
	res = append(res, "EnumArrowStylePaired")
	res = append(res, "EnumArrowStyleCombined")
	res = append(res, "EnumArrowStyleOther")

	return
}

func (enumarrowstyle EnumArrowStyle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "single")
	res = append(res, "double")
	res = append(res, "filled")
	res = append(res, "hollow")
	res = append(res, "paired")
	res = append(res, "combined")
	res = append(res, "other")

	return
}

// Utility function for EnumBendShape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbendshape EnumBendShape) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbendshape {
	// insertion code per enum code
	case EnumBendShapeAngled:
		res = "angled"
	case EnumBendShapeCurved:
		res = "curved"
	}
	return
}

func (enumbendshape *EnumBendShape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "angled":
		*enumbendshape = EnumBendShapeAngled
		return
	case "curved":
		*enumbendshape = EnumBendShapeCurved
		return
	default:
		return errUnkownEnum
	}
}

func (enumbendshape *EnumBendShape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBendShapeAngled":
		*enumbendshape = EnumBendShapeAngled
	case "EnumBendShapeCurved":
		*enumbendshape = EnumBendShapeCurved
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbendshape *EnumBendShape) ToCodeString() (res string) {

	switch *enumbendshape {
	// insertion code per enum code
	case EnumBendShapeAngled:
		res = "EnumBendShapeAngled"
	case EnumBendShapeCurved:
		res = "EnumBendShapeCurved"
	}
	return
}

func (enumbendshape EnumBendShape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBendShapeAngled")
	res = append(res, "EnumBendShapeCurved")

	return
}

func (enumbendshape EnumBendShape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "angled")
	res = append(res, "curved")

	return
}

// Utility function for EnumBreathMarkValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbreathmarkvalue EnumBreathMarkValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbreathmarkvalue {
	// insertion code per enum code
	case EnumBreathMarkValueEmpty:
		res = ""
	case EnumBreathMarkValueComma:
		res = "comma"
	case EnumBreathMarkValueTick:
		res = "tick"
	case EnumBreathMarkValueUpbow:
		res = "upbow"
	case EnumBreathMarkValueSalzedo:
		res = "salzedo"
	}
	return
}

func (enumbreathmarkvalue *EnumBreathMarkValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "":
		*enumbreathmarkvalue = EnumBreathMarkValueEmpty
		return
	case "comma":
		*enumbreathmarkvalue = EnumBreathMarkValueComma
		return
	case "tick":
		*enumbreathmarkvalue = EnumBreathMarkValueTick
		return
	case "upbow":
		*enumbreathmarkvalue = EnumBreathMarkValueUpbow
		return
	case "salzedo":
		*enumbreathmarkvalue = EnumBreathMarkValueSalzedo
		return
	default:
		return errUnkownEnum
	}
}

func (enumbreathmarkvalue *EnumBreathMarkValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBreathMarkValueEmpty":
		*enumbreathmarkvalue = EnumBreathMarkValueEmpty
	case "EnumBreathMarkValueComma":
		*enumbreathmarkvalue = EnumBreathMarkValueComma
	case "EnumBreathMarkValueTick":
		*enumbreathmarkvalue = EnumBreathMarkValueTick
	case "EnumBreathMarkValueUpbow":
		*enumbreathmarkvalue = EnumBreathMarkValueUpbow
	case "EnumBreathMarkValueSalzedo":
		*enumbreathmarkvalue = EnumBreathMarkValueSalzedo
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbreathmarkvalue *EnumBreathMarkValue) ToCodeString() (res string) {

	switch *enumbreathmarkvalue {
	// insertion code per enum code
	case EnumBreathMarkValueEmpty:
		res = "EnumBreathMarkValueEmpty"
	case EnumBreathMarkValueComma:
		res = "EnumBreathMarkValueComma"
	case EnumBreathMarkValueTick:
		res = "EnumBreathMarkValueTick"
	case EnumBreathMarkValueUpbow:
		res = "EnumBreathMarkValueUpbow"
	case EnumBreathMarkValueSalzedo:
		res = "EnumBreathMarkValueSalzedo"
	}
	return
}

func (enumbreathmarkvalue EnumBreathMarkValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBreathMarkValueEmpty")
	res = append(res, "EnumBreathMarkValueComma")
	res = append(res, "EnumBreathMarkValueTick")
	res = append(res, "EnumBreathMarkValueUpbow")
	res = append(res, "EnumBreathMarkValueSalzedo")

	return
}

func (enumbreathmarkvalue EnumBreathMarkValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "")
	res = append(res, "comma")
	res = append(res, "tick")
	res = append(res, "upbow")
	res = append(res, "salzedo")

	return
}

// Utility function for EnumCaesuraValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumcaesuravalue EnumCaesuraValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumcaesuravalue {
	// insertion code per enum code
	case EnumCaesuraValueNormal:
		res = "normal"
	case EnumCaesuraValueThick:
		res = "thick"
	case EnumCaesuraValueShort:
		res = "short"
	case EnumCaesuraValueCurved:
		res = "curved"
	case EnumCaesuraValueSingle:
		res = "single"
	case EnumCaesuraValueEmpty:
		res = ""
	}
	return
}

func (enumcaesuravalue *EnumCaesuraValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enumcaesuravalue = EnumCaesuraValueNormal
		return
	case "thick":
		*enumcaesuravalue = EnumCaesuraValueThick
		return
	case "short":
		*enumcaesuravalue = EnumCaesuraValueShort
		return
	case "curved":
		*enumcaesuravalue = EnumCaesuraValueCurved
		return
	case "single":
		*enumcaesuravalue = EnumCaesuraValueSingle
		return
	case "":
		*enumcaesuravalue = EnumCaesuraValueEmpty
		return
	default:
		return errUnkownEnum
	}
}

func (enumcaesuravalue *EnumCaesuraValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumCaesuraValueNormal":
		*enumcaesuravalue = EnumCaesuraValueNormal
	case "EnumCaesuraValueThick":
		*enumcaesuravalue = EnumCaesuraValueThick
	case "EnumCaesuraValueShort":
		*enumcaesuravalue = EnumCaesuraValueShort
	case "EnumCaesuraValueCurved":
		*enumcaesuravalue = EnumCaesuraValueCurved
	case "EnumCaesuraValueSingle":
		*enumcaesuravalue = EnumCaesuraValueSingle
	case "EnumCaesuraValueEmpty":
		*enumcaesuravalue = EnumCaesuraValueEmpty
	default:
		err = errUnkownEnum
	}
	return
}

func (enumcaesuravalue *EnumCaesuraValue) ToCodeString() (res string) {

	switch *enumcaesuravalue {
	// insertion code per enum code
	case EnumCaesuraValueNormal:
		res = "EnumCaesuraValueNormal"
	case EnumCaesuraValueThick:
		res = "EnumCaesuraValueThick"
	case EnumCaesuraValueShort:
		res = "EnumCaesuraValueShort"
	case EnumCaesuraValueCurved:
		res = "EnumCaesuraValueCurved"
	case EnumCaesuraValueSingle:
		res = "EnumCaesuraValueSingle"
	case EnumCaesuraValueEmpty:
		res = "EnumCaesuraValueEmpty"
	}
	return
}

func (enumcaesuravalue EnumCaesuraValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumCaesuraValueNormal")
	res = append(res, "EnumCaesuraValueThick")
	res = append(res, "EnumCaesuraValueShort")
	res = append(res, "EnumCaesuraValueCurved")
	res = append(res, "EnumCaesuraValueSingle")
	res = append(res, "EnumCaesuraValueEmpty")

	return
}

func (enumcaesuravalue EnumCaesuraValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "thick")
	res = append(res, "short")
	res = append(res, "curved")
	res = append(res, "single")
	res = append(res, "")

	return
}

// Utility function for EnumCircularArrow
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumcirculararrow EnumCircularArrow) ToString() (res string) {

	// migration of former implementation of enum
	switch enumcirculararrow {
	// insertion code per enum code
	case EnumCircularArrowClockwise:
		res = "clockwise"
	case EnumCircularArrowAnticlockwise:
		res = "anticlockwise"
	}
	return
}

func (enumcirculararrow *EnumCircularArrow) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "clockwise":
		*enumcirculararrow = EnumCircularArrowClockwise
		return
	case "anticlockwise":
		*enumcirculararrow = EnumCircularArrowAnticlockwise
		return
	default:
		return errUnkownEnum
	}
}

func (enumcirculararrow *EnumCircularArrow) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumCircularArrowClockwise":
		*enumcirculararrow = EnumCircularArrowClockwise
	case "EnumCircularArrowAnticlockwise":
		*enumcirculararrow = EnumCircularArrowAnticlockwise
	default:
		err = errUnkownEnum
	}
	return
}

func (enumcirculararrow *EnumCircularArrow) ToCodeString() (res string) {

	switch *enumcirculararrow {
	// insertion code per enum code
	case EnumCircularArrowClockwise:
		res = "EnumCircularArrowClockwise"
	case EnumCircularArrowAnticlockwise:
		res = "EnumCircularArrowAnticlockwise"
	}
	return
}

func (enumcirculararrow EnumCircularArrow) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumCircularArrowClockwise")
	res = append(res, "EnumCircularArrowAnticlockwise")

	return
}

func (enumcirculararrow EnumCircularArrow) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "clockwise")
	res = append(res, "anticlockwise")

	return
}

// Utility function for EnumFan
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumfan EnumFan) ToString() (res string) {

	// migration of former implementation of enum
	switch enumfan {
	// insertion code per enum code
	case EnumFanAccel:
		res = "accel"
	case EnumFanRit:
		res = "rit"
	case EnumFanNone:
		res = "none"
	}
	return
}

func (enumfan *EnumFan) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "accel":
		*enumfan = EnumFanAccel
		return
	case "rit":
		*enumfan = EnumFanRit
		return
	case "none":
		*enumfan = EnumFanNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumfan *EnumFan) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumFanAccel":
		*enumfan = EnumFanAccel
	case "EnumFanRit":
		*enumfan = EnumFanRit
	case "EnumFanNone":
		*enumfan = EnumFanNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumfan *EnumFan) ToCodeString() (res string) {

	switch *enumfan {
	// insertion code per enum code
	case EnumFanAccel:
		res = "EnumFanAccel"
	case EnumFanRit:
		res = "EnumFanRit"
	case EnumFanNone:
		res = "EnumFanNone"
	}
	return
}

func (enumfan EnumFan) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumFanAccel")
	res = append(res, "EnumFanRit")
	res = append(res, "EnumFanNone")

	return
}

func (enumfan EnumFan) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "accel")
	res = append(res, "rit")
	res = append(res, "none")

	return
}

// Utility function for EnumHandbellValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumhandbellvalue EnumHandbellValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumhandbellvalue {
	// insertion code per enum code
	case EnumHandbellValueBelltree:
		res = "belltree"
	case EnumHandbellValueDamp:
		res = "damp"
	case EnumHandbellValueEcho:
		res = "echo"
	case EnumHandbellValueGyro:
		res = "gyro"
	case EnumHandbellValueHandMartellato:
		res = "hand martellato"
	case EnumHandbellValueMalletLift:
		res = "mallet lift"
	case EnumHandbellValueMalletTable:
		res = "mallet table"
	case EnumHandbellValueMartellato:
		res = "martellato"
	case EnumHandbellValueMartellatoLift:
		res = "martellato lift"
	case EnumHandbellValueMutedMartellato:
		res = "muted martellato"
	case EnumHandbellValuePluckLift:
		res = "pluck lift"
	case EnumHandbellValueSwing:
		res = "swing"
	}
	return
}

func (enumhandbellvalue *EnumHandbellValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "belltree":
		*enumhandbellvalue = EnumHandbellValueBelltree
		return
	case "damp":
		*enumhandbellvalue = EnumHandbellValueDamp
		return
	case "echo":
		*enumhandbellvalue = EnumHandbellValueEcho
		return
	case "gyro":
		*enumhandbellvalue = EnumHandbellValueGyro
		return
	case "hand martellato":
		*enumhandbellvalue = EnumHandbellValueHandMartellato
		return
	case "mallet lift":
		*enumhandbellvalue = EnumHandbellValueMalletLift
		return
	case "mallet table":
		*enumhandbellvalue = EnumHandbellValueMalletTable
		return
	case "martellato":
		*enumhandbellvalue = EnumHandbellValueMartellato
		return
	case "martellato lift":
		*enumhandbellvalue = EnumHandbellValueMartellatoLift
		return
	case "muted martellato":
		*enumhandbellvalue = EnumHandbellValueMutedMartellato
		return
	case "pluck lift":
		*enumhandbellvalue = EnumHandbellValuePluckLift
		return
	case "swing":
		*enumhandbellvalue = EnumHandbellValueSwing
		return
	default:
		return errUnkownEnum
	}
}

func (enumhandbellvalue *EnumHandbellValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHandbellValueBelltree":
		*enumhandbellvalue = EnumHandbellValueBelltree
	case "EnumHandbellValueDamp":
		*enumhandbellvalue = EnumHandbellValueDamp
	case "EnumHandbellValueEcho":
		*enumhandbellvalue = EnumHandbellValueEcho
	case "EnumHandbellValueGyro":
		*enumhandbellvalue = EnumHandbellValueGyro
	case "EnumHandbellValueHandMartellato":
		*enumhandbellvalue = EnumHandbellValueHandMartellato
	case "EnumHandbellValueMalletLift":
		*enumhandbellvalue = EnumHandbellValueMalletLift
	case "EnumHandbellValueMalletTable":
		*enumhandbellvalue = EnumHandbellValueMalletTable
	case "EnumHandbellValueMartellato":
		*enumhandbellvalue = EnumHandbellValueMartellato
	case "EnumHandbellValueMartellatoLift":
		*enumhandbellvalue = EnumHandbellValueMartellatoLift
	case "EnumHandbellValueMutedMartellato":
		*enumhandbellvalue = EnumHandbellValueMutedMartellato
	case "EnumHandbellValuePluckLift":
		*enumhandbellvalue = EnumHandbellValuePluckLift
	case "EnumHandbellValueSwing":
		*enumhandbellvalue = EnumHandbellValueSwing
	default:
		err = errUnkownEnum
	}
	return
}

func (enumhandbellvalue *EnumHandbellValue) ToCodeString() (res string) {

	switch *enumhandbellvalue {
	// insertion code per enum code
	case EnumHandbellValueBelltree:
		res = "EnumHandbellValueBelltree"
	case EnumHandbellValueDamp:
		res = "EnumHandbellValueDamp"
	case EnumHandbellValueEcho:
		res = "EnumHandbellValueEcho"
	case EnumHandbellValueGyro:
		res = "EnumHandbellValueGyro"
	case EnumHandbellValueHandMartellato:
		res = "EnumHandbellValueHandMartellato"
	case EnumHandbellValueMalletLift:
		res = "EnumHandbellValueMalletLift"
	case EnumHandbellValueMalletTable:
		res = "EnumHandbellValueMalletTable"
	case EnumHandbellValueMartellato:
		res = "EnumHandbellValueMartellato"
	case EnumHandbellValueMartellatoLift:
		res = "EnumHandbellValueMartellatoLift"
	case EnumHandbellValueMutedMartellato:
		res = "EnumHandbellValueMutedMartellato"
	case EnumHandbellValuePluckLift:
		res = "EnumHandbellValuePluckLift"
	case EnumHandbellValueSwing:
		res = "EnumHandbellValueSwing"
	}
	return
}

func (enumhandbellvalue EnumHandbellValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHandbellValueBelltree")
	res = append(res, "EnumHandbellValueDamp")
	res = append(res, "EnumHandbellValueEcho")
	res = append(res, "EnumHandbellValueGyro")
	res = append(res, "EnumHandbellValueHandMartellato")
	res = append(res, "EnumHandbellValueMalletLift")
	res = append(res, "EnumHandbellValueMalletTable")
	res = append(res, "EnumHandbellValueMartellato")
	res = append(res, "EnumHandbellValueMartellatoLift")
	res = append(res, "EnumHandbellValueMutedMartellato")
	res = append(res, "EnumHandbellValuePluckLift")
	res = append(res, "EnumHandbellValueSwing")

	return
}

func (enumhandbellvalue EnumHandbellValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "belltree")
	res = append(res, "damp")
	res = append(res, "echo")
	res = append(res, "gyro")
	res = append(res, "hand martellato")
	res = append(res, "mallet lift")
	res = append(res, "mallet table")
	res = append(res, "martellato")
	res = append(res, "martellato lift")
	res = append(res, "muted martellato")
	res = append(res, "pluck lift")
	res = append(res, "swing")

	return
}

// Utility function for EnumHarmonClosedLocation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumharmonclosedlocation EnumHarmonClosedLocation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumharmonclosedlocation {
	// insertion code per enum code
	case EnumHarmonClosedLocationRight:
		res = "right"
	case EnumHarmonClosedLocationBottom:
		res = "bottom"
	case EnumHarmonClosedLocationLeft:
		res = "left"
	case EnumHarmonClosedLocationTop:
		res = "top"
	}
	return
}

func (enumharmonclosedlocation *EnumHarmonClosedLocation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enumharmonclosedlocation = EnumHarmonClosedLocationRight
		return
	case "bottom":
		*enumharmonclosedlocation = EnumHarmonClosedLocationBottom
		return
	case "left":
		*enumharmonclosedlocation = EnumHarmonClosedLocationLeft
		return
	case "top":
		*enumharmonclosedlocation = EnumHarmonClosedLocationTop
		return
	default:
		return errUnkownEnum
	}
}

func (enumharmonclosedlocation *EnumHarmonClosedLocation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHarmonClosedLocationRight":
		*enumharmonclosedlocation = EnumHarmonClosedLocationRight
	case "EnumHarmonClosedLocationBottom":
		*enumharmonclosedlocation = EnumHarmonClosedLocationBottom
	case "EnumHarmonClosedLocationLeft":
		*enumharmonclosedlocation = EnumHarmonClosedLocationLeft
	case "EnumHarmonClosedLocationTop":
		*enumharmonclosedlocation = EnumHarmonClosedLocationTop
	default:
		err = errUnkownEnum
	}
	return
}

func (enumharmonclosedlocation *EnumHarmonClosedLocation) ToCodeString() (res string) {

	switch *enumharmonclosedlocation {
	// insertion code per enum code
	case EnumHarmonClosedLocationRight:
		res = "EnumHarmonClosedLocationRight"
	case EnumHarmonClosedLocationBottom:
		res = "EnumHarmonClosedLocationBottom"
	case EnumHarmonClosedLocationLeft:
		res = "EnumHarmonClosedLocationLeft"
	case EnumHarmonClosedLocationTop:
		res = "EnumHarmonClosedLocationTop"
	}
	return
}

func (enumharmonclosedlocation EnumHarmonClosedLocation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHarmonClosedLocationRight")
	res = append(res, "EnumHarmonClosedLocationBottom")
	res = append(res, "EnumHarmonClosedLocationLeft")
	res = append(res, "EnumHarmonClosedLocationTop")

	return
}

func (enumharmonclosedlocation EnumHarmonClosedLocation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for EnumHarmonClosedValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumharmonclosedvalue EnumHarmonClosedValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumharmonclosedvalue {
	// insertion code per enum code
	case EnumHarmonClosedValueYes:
		res = "yes"
	case EnumHarmonClosedValueNo:
		res = "no"
	case EnumHarmonClosedValueHalf:
		res = "half"
	}
	return
}

func (enumharmonclosedvalue *EnumHarmonClosedValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enumharmonclosedvalue = EnumHarmonClosedValueYes
		return
	case "no":
		*enumharmonclosedvalue = EnumHarmonClosedValueNo
		return
	case "half":
		*enumharmonclosedvalue = EnumHarmonClosedValueHalf
		return
	default:
		return errUnkownEnum
	}
}

func (enumharmonclosedvalue *EnumHarmonClosedValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHarmonClosedValueYes":
		*enumharmonclosedvalue = EnumHarmonClosedValueYes
	case "EnumHarmonClosedValueNo":
		*enumharmonclosedvalue = EnumHarmonClosedValueNo
	case "EnumHarmonClosedValueHalf":
		*enumharmonclosedvalue = EnumHarmonClosedValueHalf
	default:
		err = errUnkownEnum
	}
	return
}

func (enumharmonclosedvalue *EnumHarmonClosedValue) ToCodeString() (res string) {

	switch *enumharmonclosedvalue {
	// insertion code per enum code
	case EnumHarmonClosedValueYes:
		res = "EnumHarmonClosedValueYes"
	case EnumHarmonClosedValueNo:
		res = "EnumHarmonClosedValueNo"
	case EnumHarmonClosedValueHalf:
		res = "EnumHarmonClosedValueHalf"
	}
	return
}

func (enumharmonclosedvalue EnumHarmonClosedValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHarmonClosedValueYes")
	res = append(res, "EnumHarmonClosedValueNo")
	res = append(res, "EnumHarmonClosedValueHalf")

	return
}

func (enumharmonclosedvalue EnumHarmonClosedValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for EnumHoleClosedLocation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumholeclosedlocation EnumHoleClosedLocation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumholeclosedlocation {
	// insertion code per enum code
	case EnumHoleClosedLocationRight:
		res = "right"
	case EnumHoleClosedLocationBottom:
		res = "bottom"
	case EnumHoleClosedLocationLeft:
		res = "left"
	case EnumHoleClosedLocationTop:
		res = "top"
	}
	return
}

func (enumholeclosedlocation *EnumHoleClosedLocation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enumholeclosedlocation = EnumHoleClosedLocationRight
		return
	case "bottom":
		*enumholeclosedlocation = EnumHoleClosedLocationBottom
		return
	case "left":
		*enumholeclosedlocation = EnumHoleClosedLocationLeft
		return
	case "top":
		*enumholeclosedlocation = EnumHoleClosedLocationTop
		return
	default:
		return errUnkownEnum
	}
}

func (enumholeclosedlocation *EnumHoleClosedLocation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHoleClosedLocationRight":
		*enumholeclosedlocation = EnumHoleClosedLocationRight
	case "EnumHoleClosedLocationBottom":
		*enumholeclosedlocation = EnumHoleClosedLocationBottom
	case "EnumHoleClosedLocationLeft":
		*enumholeclosedlocation = EnumHoleClosedLocationLeft
	case "EnumHoleClosedLocationTop":
		*enumholeclosedlocation = EnumHoleClosedLocationTop
	default:
		err = errUnkownEnum
	}
	return
}

func (enumholeclosedlocation *EnumHoleClosedLocation) ToCodeString() (res string) {

	switch *enumholeclosedlocation {
	// insertion code per enum code
	case EnumHoleClosedLocationRight:
		res = "EnumHoleClosedLocationRight"
	case EnumHoleClosedLocationBottom:
		res = "EnumHoleClosedLocationBottom"
	case EnumHoleClosedLocationLeft:
		res = "EnumHoleClosedLocationLeft"
	case EnumHoleClosedLocationTop:
		res = "EnumHoleClosedLocationTop"
	}
	return
}

func (enumholeclosedlocation EnumHoleClosedLocation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHoleClosedLocationRight")
	res = append(res, "EnumHoleClosedLocationBottom")
	res = append(res, "EnumHoleClosedLocationLeft")
	res = append(res, "EnumHoleClosedLocationTop")

	return
}

func (enumholeclosedlocation EnumHoleClosedLocation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for EnumHoleClosedValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumholeclosedvalue EnumHoleClosedValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumholeclosedvalue {
	// insertion code per enum code
	case EnumHoleClosedValueYes:
		res = "yes"
	case EnumHoleClosedValueNo:
		res = "no"
	case EnumHoleClosedValueHalf:
		res = "half"
	}
	return
}

func (enumholeclosedvalue *EnumHoleClosedValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enumholeclosedvalue = EnumHoleClosedValueYes
		return
	case "no":
		*enumholeclosedvalue = EnumHoleClosedValueNo
		return
	case "half":
		*enumholeclosedvalue = EnumHoleClosedValueHalf
		return
	default:
		return errUnkownEnum
	}
}

func (enumholeclosedvalue *EnumHoleClosedValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHoleClosedValueYes":
		*enumholeclosedvalue = EnumHoleClosedValueYes
	case "EnumHoleClosedValueNo":
		*enumholeclosedvalue = EnumHoleClosedValueNo
	case "EnumHoleClosedValueHalf":
		*enumholeclosedvalue = EnumHoleClosedValueHalf
	default:
		err = errUnkownEnum
	}
	return
}

func (enumholeclosedvalue *EnumHoleClosedValue) ToCodeString() (res string) {

	switch *enumholeclosedvalue {
	// insertion code per enum code
	case EnumHoleClosedValueYes:
		res = "EnumHoleClosedValueYes"
	case EnumHoleClosedValueNo:
		res = "EnumHoleClosedValueNo"
	case EnumHoleClosedValueHalf:
		res = "EnumHoleClosedValueHalf"
	}
	return
}

func (enumholeclosedvalue EnumHoleClosedValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHoleClosedValueYes")
	res = append(res, "EnumHoleClosedValueNo")
	res = append(res, "EnumHoleClosedValueHalf")

	return
}

func (enumholeclosedvalue EnumHoleClosedValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for EnumNoteheadValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumnoteheadvalue EnumNoteheadValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumnoteheadvalue {
	// insertion code per enum code
	case EnumNoteheadValueSlash:
		res = "slash"
	case EnumNoteheadValueTriangle:
		res = "triangle"
	case EnumNoteheadValueDiamond:
		res = "diamond"
	case EnumNoteheadValueSquare:
		res = "square"
	case EnumNoteheadValueCross:
		res = "cross"
	case EnumNoteheadValueX:
		res = "x"
	case EnumNoteheadValueCircleX:
		res = "circle-x"
	case EnumNoteheadValueInvertedTriangle:
		res = "inverted triangle"
	case EnumNoteheadValueArrowDown:
		res = "arrow down"
	case EnumNoteheadValueArrowUp:
		res = "arrow up"
	case EnumNoteheadValueCircled:
		res = "circled"
	case EnumNoteheadValueSlashed:
		res = "slashed"
	case EnumNoteheadValueBackSlashed:
		res = "back slashed"
	case EnumNoteheadValueNormal:
		res = "normal"
	case EnumNoteheadValueCluster:
		res = "cluster"
	case EnumNoteheadValueCircleDot:
		res = "circle dot"
	case EnumNoteheadValueLeftTriangle:
		res = "left triangle"
	case EnumNoteheadValueRectangle:
		res = "rectangle"
	case EnumNoteheadValueNone:
		res = "none"
	case EnumNoteheadValueDo:
		res = "do"
	case EnumNoteheadValueRe:
		res = "re"
	case EnumNoteheadValueMi:
		res = "mi"
	case EnumNoteheadValueFa:
		res = "fa"
	case EnumNoteheadValueFaUp:
		res = "fa up"
	case EnumNoteheadValueSo:
		res = "so"
	case EnumNoteheadValueLa:
		res = "la"
	case EnumNoteheadValueTi:
		res = "ti"
	case EnumNoteheadValueOther:
		res = "other"
	}
	return
}

func (enumnoteheadvalue *EnumNoteheadValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "slash":
		*enumnoteheadvalue = EnumNoteheadValueSlash
		return
	case "triangle":
		*enumnoteheadvalue = EnumNoteheadValueTriangle
		return
	case "diamond":
		*enumnoteheadvalue = EnumNoteheadValueDiamond
		return
	case "square":
		*enumnoteheadvalue = EnumNoteheadValueSquare
		return
	case "cross":
		*enumnoteheadvalue = EnumNoteheadValueCross
		return
	case "x":
		*enumnoteheadvalue = EnumNoteheadValueX
		return
	case "circle-x":
		*enumnoteheadvalue = EnumNoteheadValueCircleX
		return
	case "inverted triangle":
		*enumnoteheadvalue = EnumNoteheadValueInvertedTriangle
		return
	case "arrow down":
		*enumnoteheadvalue = EnumNoteheadValueArrowDown
		return
	case "arrow up":
		*enumnoteheadvalue = EnumNoteheadValueArrowUp
		return
	case "circled":
		*enumnoteheadvalue = EnumNoteheadValueCircled
		return
	case "slashed":
		*enumnoteheadvalue = EnumNoteheadValueSlashed
		return
	case "back slashed":
		*enumnoteheadvalue = EnumNoteheadValueBackSlashed
		return
	case "normal":
		*enumnoteheadvalue = EnumNoteheadValueNormal
		return
	case "cluster":
		*enumnoteheadvalue = EnumNoteheadValueCluster
		return
	case "circle dot":
		*enumnoteheadvalue = EnumNoteheadValueCircleDot
		return
	case "left triangle":
		*enumnoteheadvalue = EnumNoteheadValueLeftTriangle
		return
	case "rectangle":
		*enumnoteheadvalue = EnumNoteheadValueRectangle
		return
	case "none":
		*enumnoteheadvalue = EnumNoteheadValueNone
		return
	case "do":
		*enumnoteheadvalue = EnumNoteheadValueDo
		return
	case "re":
		*enumnoteheadvalue = EnumNoteheadValueRe
		return
	case "mi":
		*enumnoteheadvalue = EnumNoteheadValueMi
		return
	case "fa":
		*enumnoteheadvalue = EnumNoteheadValueFa
		return
	case "fa up":
		*enumnoteheadvalue = EnumNoteheadValueFaUp
		return
	case "so":
		*enumnoteheadvalue = EnumNoteheadValueSo
		return
	case "la":
		*enumnoteheadvalue = EnumNoteheadValueLa
		return
	case "ti":
		*enumnoteheadvalue = EnumNoteheadValueTi
		return
	case "other":
		*enumnoteheadvalue = EnumNoteheadValueOther
		return
	default:
		return errUnkownEnum
	}
}

func (enumnoteheadvalue *EnumNoteheadValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumNoteheadValueSlash":
		*enumnoteheadvalue = EnumNoteheadValueSlash
	case "EnumNoteheadValueTriangle":
		*enumnoteheadvalue = EnumNoteheadValueTriangle
	case "EnumNoteheadValueDiamond":
		*enumnoteheadvalue = EnumNoteheadValueDiamond
	case "EnumNoteheadValueSquare":
		*enumnoteheadvalue = EnumNoteheadValueSquare
	case "EnumNoteheadValueCross":
		*enumnoteheadvalue = EnumNoteheadValueCross
	case "EnumNoteheadValueX":
		*enumnoteheadvalue = EnumNoteheadValueX
	case "EnumNoteheadValueCircleX":
		*enumnoteheadvalue = EnumNoteheadValueCircleX
	case "EnumNoteheadValueInvertedTriangle":
		*enumnoteheadvalue = EnumNoteheadValueInvertedTriangle
	case "EnumNoteheadValueArrowDown":
		*enumnoteheadvalue = EnumNoteheadValueArrowDown
	case "EnumNoteheadValueArrowUp":
		*enumnoteheadvalue = EnumNoteheadValueArrowUp
	case "EnumNoteheadValueCircled":
		*enumnoteheadvalue = EnumNoteheadValueCircled
	case "EnumNoteheadValueSlashed":
		*enumnoteheadvalue = EnumNoteheadValueSlashed
	case "EnumNoteheadValueBackSlashed":
		*enumnoteheadvalue = EnumNoteheadValueBackSlashed
	case "EnumNoteheadValueNormal":
		*enumnoteheadvalue = EnumNoteheadValueNormal
	case "EnumNoteheadValueCluster":
		*enumnoteheadvalue = EnumNoteheadValueCluster
	case "EnumNoteheadValueCircleDot":
		*enumnoteheadvalue = EnumNoteheadValueCircleDot
	case "EnumNoteheadValueLeftTriangle":
		*enumnoteheadvalue = EnumNoteheadValueLeftTriangle
	case "EnumNoteheadValueRectangle":
		*enumnoteheadvalue = EnumNoteheadValueRectangle
	case "EnumNoteheadValueNone":
		*enumnoteheadvalue = EnumNoteheadValueNone
	case "EnumNoteheadValueDo":
		*enumnoteheadvalue = EnumNoteheadValueDo
	case "EnumNoteheadValueRe":
		*enumnoteheadvalue = EnumNoteheadValueRe
	case "EnumNoteheadValueMi":
		*enumnoteheadvalue = EnumNoteheadValueMi
	case "EnumNoteheadValueFa":
		*enumnoteheadvalue = EnumNoteheadValueFa
	case "EnumNoteheadValueFaUp":
		*enumnoteheadvalue = EnumNoteheadValueFaUp
	case "EnumNoteheadValueSo":
		*enumnoteheadvalue = EnumNoteheadValueSo
	case "EnumNoteheadValueLa":
		*enumnoteheadvalue = EnumNoteheadValueLa
	case "EnumNoteheadValueTi":
		*enumnoteheadvalue = EnumNoteheadValueTi
	case "EnumNoteheadValueOther":
		*enumnoteheadvalue = EnumNoteheadValueOther
	default:
		err = errUnkownEnum
	}
	return
}

func (enumnoteheadvalue *EnumNoteheadValue) ToCodeString() (res string) {

	switch *enumnoteheadvalue {
	// insertion code per enum code
	case EnumNoteheadValueSlash:
		res = "EnumNoteheadValueSlash"
	case EnumNoteheadValueTriangle:
		res = "EnumNoteheadValueTriangle"
	case EnumNoteheadValueDiamond:
		res = "EnumNoteheadValueDiamond"
	case EnumNoteheadValueSquare:
		res = "EnumNoteheadValueSquare"
	case EnumNoteheadValueCross:
		res = "EnumNoteheadValueCross"
	case EnumNoteheadValueX:
		res = "EnumNoteheadValueX"
	case EnumNoteheadValueCircleX:
		res = "EnumNoteheadValueCircleX"
	case EnumNoteheadValueInvertedTriangle:
		res = "EnumNoteheadValueInvertedTriangle"
	case EnumNoteheadValueArrowDown:
		res = "EnumNoteheadValueArrowDown"
	case EnumNoteheadValueArrowUp:
		res = "EnumNoteheadValueArrowUp"
	case EnumNoteheadValueCircled:
		res = "EnumNoteheadValueCircled"
	case EnumNoteheadValueSlashed:
		res = "EnumNoteheadValueSlashed"
	case EnumNoteheadValueBackSlashed:
		res = "EnumNoteheadValueBackSlashed"
	case EnumNoteheadValueNormal:
		res = "EnumNoteheadValueNormal"
	case EnumNoteheadValueCluster:
		res = "EnumNoteheadValueCluster"
	case EnumNoteheadValueCircleDot:
		res = "EnumNoteheadValueCircleDot"
	case EnumNoteheadValueLeftTriangle:
		res = "EnumNoteheadValueLeftTriangle"
	case EnumNoteheadValueRectangle:
		res = "EnumNoteheadValueRectangle"
	case EnumNoteheadValueNone:
		res = "EnumNoteheadValueNone"
	case EnumNoteheadValueDo:
		res = "EnumNoteheadValueDo"
	case EnumNoteheadValueRe:
		res = "EnumNoteheadValueRe"
	case EnumNoteheadValueMi:
		res = "EnumNoteheadValueMi"
	case EnumNoteheadValueFa:
		res = "EnumNoteheadValueFa"
	case EnumNoteheadValueFaUp:
		res = "EnumNoteheadValueFaUp"
	case EnumNoteheadValueSo:
		res = "EnumNoteheadValueSo"
	case EnumNoteheadValueLa:
		res = "EnumNoteheadValueLa"
	case EnumNoteheadValueTi:
		res = "EnumNoteheadValueTi"
	case EnumNoteheadValueOther:
		res = "EnumNoteheadValueOther"
	}
	return
}

func (enumnoteheadvalue EnumNoteheadValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumNoteheadValueSlash")
	res = append(res, "EnumNoteheadValueTriangle")
	res = append(res, "EnumNoteheadValueDiamond")
	res = append(res, "EnumNoteheadValueSquare")
	res = append(res, "EnumNoteheadValueCross")
	res = append(res, "EnumNoteheadValueX")
	res = append(res, "EnumNoteheadValueCircleX")
	res = append(res, "EnumNoteheadValueInvertedTriangle")
	res = append(res, "EnumNoteheadValueArrowDown")
	res = append(res, "EnumNoteheadValueArrowUp")
	res = append(res, "EnumNoteheadValueCircled")
	res = append(res, "EnumNoteheadValueSlashed")
	res = append(res, "EnumNoteheadValueBackSlashed")
	res = append(res, "EnumNoteheadValueNormal")
	res = append(res, "EnumNoteheadValueCluster")
	res = append(res, "EnumNoteheadValueCircleDot")
	res = append(res, "EnumNoteheadValueLeftTriangle")
	res = append(res, "EnumNoteheadValueRectangle")
	res = append(res, "EnumNoteheadValueNone")
	res = append(res, "EnumNoteheadValueDo")
	res = append(res, "EnumNoteheadValueRe")
	res = append(res, "EnumNoteheadValueMi")
	res = append(res, "EnumNoteheadValueFa")
	res = append(res, "EnumNoteheadValueFaUp")
	res = append(res, "EnumNoteheadValueSo")
	res = append(res, "EnumNoteheadValueLa")
	res = append(res, "EnumNoteheadValueTi")
	res = append(res, "EnumNoteheadValueOther")

	return
}

func (enumnoteheadvalue EnumNoteheadValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "slash")
	res = append(res, "triangle")
	res = append(res, "diamond")
	res = append(res, "square")
	res = append(res, "cross")
	res = append(res, "x")
	res = append(res, "circle-x")
	res = append(res, "inverted triangle")
	res = append(res, "arrow down")
	res = append(res, "arrow up")
	res = append(res, "circled")
	res = append(res, "slashed")
	res = append(res, "back slashed")
	res = append(res, "normal")
	res = append(res, "cluster")
	res = append(res, "circle dot")
	res = append(res, "left triangle")
	res = append(res, "rectangle")
	res = append(res, "none")
	res = append(res, "do")
	res = append(res, "re")
	res = append(res, "mi")
	res = append(res, "fa")
	res = append(res, "fa up")
	res = append(res, "so")
	res = append(res, "la")
	res = append(res, "ti")
	res = append(res, "other")

	return
}

// Utility function for EnumStemValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstemvalue EnumStemValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstemvalue {
	// insertion code per enum code
	case EnumStemValueDown:
		res = "down"
	case EnumStemValueUp:
		res = "up"
	case EnumStemValueDouble:
		res = "double"
	case EnumStemValueNone:
		res = "none"
	}
	return
}

func (enumstemvalue *EnumStemValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*enumstemvalue = EnumStemValueDown
		return
	case "up":
		*enumstemvalue = EnumStemValueUp
		return
	case "double":
		*enumstemvalue = EnumStemValueDouble
		return
	case "none":
		*enumstemvalue = EnumStemValueNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumstemvalue *EnumStemValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStemValueDown":
		*enumstemvalue = EnumStemValueDown
	case "EnumStemValueUp":
		*enumstemvalue = EnumStemValueUp
	case "EnumStemValueDouble":
		*enumstemvalue = EnumStemValueDouble
	case "EnumStemValueNone":
		*enumstemvalue = EnumStemValueNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstemvalue *EnumStemValue) ToCodeString() (res string) {

	switch *enumstemvalue {
	// insertion code per enum code
	case EnumStemValueDown:
		res = "EnumStemValueDown"
	case EnumStemValueUp:
		res = "EnumStemValueUp"
	case EnumStemValueDouble:
		res = "EnumStemValueDouble"
	case EnumStemValueNone:
		res = "EnumStemValueNone"
	}
	return
}

func (enumstemvalue EnumStemValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStemValueDown")
	res = append(res, "EnumStemValueUp")
	res = append(res, "EnumStemValueDouble")
	res = append(res, "EnumStemValueNone")

	return
}

func (enumstemvalue EnumStemValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "double")
	res = append(res, "none")

	return
}

// Utility function for EnumSyllabic
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsyllabic EnumSyllabic) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsyllabic {
	// insertion code per enum code
	case EnumSyllabicSingle:
		res = "single"
	case EnumSyllabicBegin:
		res = "begin"
	case EnumSyllabicEnd:
		res = "end"
	case EnumSyllabicMiddle:
		res = "middle"
	}
	return
}

func (enumsyllabic *EnumSyllabic) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*enumsyllabic = EnumSyllabicSingle
		return
	case "begin":
		*enumsyllabic = EnumSyllabicBegin
		return
	case "end":
		*enumsyllabic = EnumSyllabicEnd
		return
	case "middle":
		*enumsyllabic = EnumSyllabicMiddle
		return
	default:
		return errUnkownEnum
	}
}

func (enumsyllabic *EnumSyllabic) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSyllabicSingle":
		*enumsyllabic = EnumSyllabicSingle
	case "EnumSyllabicBegin":
		*enumsyllabic = EnumSyllabicBegin
	case "EnumSyllabicEnd":
		*enumsyllabic = EnumSyllabicEnd
	case "EnumSyllabicMiddle":
		*enumsyllabic = EnumSyllabicMiddle
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsyllabic *EnumSyllabic) ToCodeString() (res string) {

	switch *enumsyllabic {
	// insertion code per enum code
	case EnumSyllabicSingle:
		res = "EnumSyllabicSingle"
	case EnumSyllabicBegin:
		res = "EnumSyllabicBegin"
	case EnumSyllabicEnd:
		res = "EnumSyllabicEnd"
	case EnumSyllabicMiddle:
		res = "EnumSyllabicMiddle"
	}
	return
}

func (enumsyllabic EnumSyllabic) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSyllabicSingle")
	res = append(res, "EnumSyllabicBegin")
	res = append(res, "EnumSyllabicEnd")
	res = append(res, "EnumSyllabicMiddle")

	return
}

func (enumsyllabic EnumSyllabic) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "single")
	res = append(res, "begin")
	res = append(res, "end")
	res = append(res, "middle")

	return
}

// Utility function for EnumTapHand
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtaphand EnumTapHand) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtaphand {
	// insertion code per enum code
	case EnumTapHandLeft:
		res = "left"
	case EnumTapHandRight:
		res = "right"
	}
	return
}

func (enumtaphand *EnumTapHand) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enumtaphand = EnumTapHandLeft
		return
	case "right":
		*enumtaphand = EnumTapHandRight
		return
	default:
		return errUnkownEnum
	}
}

func (enumtaphand *EnumTapHand) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTapHandLeft":
		*enumtaphand = EnumTapHandLeft
	case "EnumTapHandRight":
		*enumtaphand = EnumTapHandRight
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtaphand *EnumTapHand) ToCodeString() (res string) {

	switch *enumtaphand {
	// insertion code per enum code
	case EnumTapHandLeft:
		res = "EnumTapHandLeft"
	case EnumTapHandRight:
		res = "EnumTapHandRight"
	}
	return
}

func (enumtaphand EnumTapHand) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTapHandLeft")
	res = append(res, "EnumTapHandRight")

	return
}

func (enumtaphand EnumTapHand) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

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
