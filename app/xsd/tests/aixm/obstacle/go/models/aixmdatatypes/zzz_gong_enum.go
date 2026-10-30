// generated code - do not edit
package aixmdatatypes

// insertion point of enum utility functions
// Utility function for Enum_CodeColourBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecolourbasetype_member_1 Enum_CodeColourBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecolourbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeColourBaseType_Member_1_YELLOW:
		res = "YELLOW"
	case Enum_CodeColourBaseType_Member_1_RED:
		res = "RED"
	case Enum_CodeColourBaseType_Member_1_WHITE:
		res = "WHITE"
	case Enum_CodeColourBaseType_Member_1_BLUE:
		res = "BLUE"
	case Enum_CodeColourBaseType_Member_1_GREEN:
		res = "GREEN"
	case Enum_CodeColourBaseType_Member_1_PURPLE:
		res = "PURPLE"
	case Enum_CodeColourBaseType_Member_1_ORANGE:
		res = "ORANGE"
	case Enum_CodeColourBaseType_Member_1_AMBER:
		res = "AMBER"
	case Enum_CodeColourBaseType_Member_1_BLACK:
		res = "BLACK"
	case Enum_CodeColourBaseType_Member_1_BROWN:
		res = "BROWN"
	case Enum_CodeColourBaseType_Member_1_GREY:
		res = "GREY"
	case Enum_CodeColourBaseType_Member_1_LIGHT_GREY:
		res = "LIGHT_GREY"
	case Enum_CodeColourBaseType_Member_1_MAGENTA:
		res = "MAGENTA"
	case Enum_CodeColourBaseType_Member_1_PINK:
		res = "PINK"
	case Enum_CodeColourBaseType_Member_1_VIOLET:
		res = "VIOLET"
	}
	return
}

func (enum_codecolourbasetype_member_1 *Enum_CodeColourBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "YELLOW":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_YELLOW
		return
	case "RED":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_RED
		return
	case "WHITE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_WHITE
		return
	case "BLUE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BLUE
		return
	case "GREEN":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_GREEN
		return
	case "PURPLE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_PURPLE
		return
	case "ORANGE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_ORANGE
		return
	case "AMBER":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_AMBER
		return
	case "BLACK":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BLACK
		return
	case "BROWN":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BROWN
		return
	case "GREY":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_GREY
		return
	case "LIGHT_GREY":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_LIGHT_GREY
		return
	case "MAGENTA":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_MAGENTA
		return
	case "PINK":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_PINK
		return
	case "VIOLET":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_VIOLET
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecolourbasetype_member_1 *Enum_CodeColourBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeColourBaseType_Member_1_YELLOW":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_YELLOW
	case "Enum_CodeColourBaseType_Member_1_RED":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_RED
	case "Enum_CodeColourBaseType_Member_1_WHITE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_WHITE
	case "Enum_CodeColourBaseType_Member_1_BLUE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BLUE
	case "Enum_CodeColourBaseType_Member_1_GREEN":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_GREEN
	case "Enum_CodeColourBaseType_Member_1_PURPLE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_PURPLE
	case "Enum_CodeColourBaseType_Member_1_ORANGE":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_ORANGE
	case "Enum_CodeColourBaseType_Member_1_AMBER":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_AMBER
	case "Enum_CodeColourBaseType_Member_1_BLACK":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BLACK
	case "Enum_CodeColourBaseType_Member_1_BROWN":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_BROWN
	case "Enum_CodeColourBaseType_Member_1_GREY":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_GREY
	case "Enum_CodeColourBaseType_Member_1_LIGHT_GREY":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_LIGHT_GREY
	case "Enum_CodeColourBaseType_Member_1_MAGENTA":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_MAGENTA
	case "Enum_CodeColourBaseType_Member_1_PINK":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_PINK
	case "Enum_CodeColourBaseType_Member_1_VIOLET":
		*enum_codecolourbasetype_member_1 = Enum_CodeColourBaseType_Member_1_VIOLET
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecolourbasetype_member_1 *Enum_CodeColourBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecolourbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeColourBaseType_Member_1_YELLOW:
		res = "Enum_CodeColourBaseType_Member_1_YELLOW"
	case Enum_CodeColourBaseType_Member_1_RED:
		res = "Enum_CodeColourBaseType_Member_1_RED"
	case Enum_CodeColourBaseType_Member_1_WHITE:
		res = "Enum_CodeColourBaseType_Member_1_WHITE"
	case Enum_CodeColourBaseType_Member_1_BLUE:
		res = "Enum_CodeColourBaseType_Member_1_BLUE"
	case Enum_CodeColourBaseType_Member_1_GREEN:
		res = "Enum_CodeColourBaseType_Member_1_GREEN"
	case Enum_CodeColourBaseType_Member_1_PURPLE:
		res = "Enum_CodeColourBaseType_Member_1_PURPLE"
	case Enum_CodeColourBaseType_Member_1_ORANGE:
		res = "Enum_CodeColourBaseType_Member_1_ORANGE"
	case Enum_CodeColourBaseType_Member_1_AMBER:
		res = "Enum_CodeColourBaseType_Member_1_AMBER"
	case Enum_CodeColourBaseType_Member_1_BLACK:
		res = "Enum_CodeColourBaseType_Member_1_BLACK"
	case Enum_CodeColourBaseType_Member_1_BROWN:
		res = "Enum_CodeColourBaseType_Member_1_BROWN"
	case Enum_CodeColourBaseType_Member_1_GREY:
		res = "Enum_CodeColourBaseType_Member_1_GREY"
	case Enum_CodeColourBaseType_Member_1_LIGHT_GREY:
		res = "Enum_CodeColourBaseType_Member_1_LIGHT_GREY"
	case Enum_CodeColourBaseType_Member_1_MAGENTA:
		res = "Enum_CodeColourBaseType_Member_1_MAGENTA"
	case Enum_CodeColourBaseType_Member_1_PINK:
		res = "Enum_CodeColourBaseType_Member_1_PINK"
	case Enum_CodeColourBaseType_Member_1_VIOLET:
		res = "Enum_CodeColourBaseType_Member_1_VIOLET"
	}
	return
}

func (enum_codecolourbasetype_member_1 Enum_CodeColourBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeColourBaseType_Member_1_YELLOW")
	res = append(res, "Enum_CodeColourBaseType_Member_1_RED")
	res = append(res, "Enum_CodeColourBaseType_Member_1_WHITE")
	res = append(res, "Enum_CodeColourBaseType_Member_1_BLUE")
	res = append(res, "Enum_CodeColourBaseType_Member_1_GREEN")
	res = append(res, "Enum_CodeColourBaseType_Member_1_PURPLE")
	res = append(res, "Enum_CodeColourBaseType_Member_1_ORANGE")
	res = append(res, "Enum_CodeColourBaseType_Member_1_AMBER")
	res = append(res, "Enum_CodeColourBaseType_Member_1_BLACK")
	res = append(res, "Enum_CodeColourBaseType_Member_1_BROWN")
	res = append(res, "Enum_CodeColourBaseType_Member_1_GREY")
	res = append(res, "Enum_CodeColourBaseType_Member_1_LIGHT_GREY")
	res = append(res, "Enum_CodeColourBaseType_Member_1_MAGENTA")
	res = append(res, "Enum_CodeColourBaseType_Member_1_PINK")
	res = append(res, "Enum_CodeColourBaseType_Member_1_VIOLET")

	return
}

func (enum_codecolourbasetype_member_1 Enum_CodeColourBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "YELLOW")
	res = append(res, "RED")
	res = append(res, "WHITE")
	res = append(res, "BLUE")
	res = append(res, "GREEN")
	res = append(res, "PURPLE")
	res = append(res, "ORANGE")
	res = append(res, "AMBER")
	res = append(res, "BLACK")
	res = append(res, "BROWN")
	res = append(res, "GREY")
	res = append(res, "LIGHT_GREY")
	res = append(res, "MAGENTA")
	res = append(res, "PINK")
	res = append(res, "VIOLET")

	return
}

// Utility function for Enum_CodeDayBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedaybasetype_member_1 Enum_CodeDayBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDayBaseType_Member_1_MON:
		res = "MON"
	case Enum_CodeDayBaseType_Member_1_TUE:
		res = "TUE"
	case Enum_CodeDayBaseType_Member_1_WED:
		res = "WED"
	case Enum_CodeDayBaseType_Member_1_THU:
		res = "THU"
	case Enum_CodeDayBaseType_Member_1_FRI:
		res = "FRI"
	case Enum_CodeDayBaseType_Member_1_SAT:
		res = "SAT"
	case Enum_CodeDayBaseType_Member_1_SUN:
		res = "SUN"
	case Enum_CodeDayBaseType_Member_1_WORK_DAY:
		res = "WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY:
		res = "BEF_WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY:
		res = "AFT_WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_HOL:
		res = "HOL"
	case Enum_CodeDayBaseType_Member_1_BEF_HOL:
		res = "BEF_HOL"
	case Enum_CodeDayBaseType_Member_1_AFT_HOL:
		res = "AFT_HOL"
	case Enum_CodeDayBaseType_Member_1_ANY_XSD:
		res = "ANY"
	case Enum_CodeDayBaseType_Member_1_BUSY_FRI:
		res = "BUSY_FRI"
	}
	return
}

func (enum_codedaybasetype_member_1 *Enum_CodeDayBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MON":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_MON
		return
	case "TUE":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_TUE
		return
	case "WED":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_WED
		return
	case "THU":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_THU
		return
	case "FRI":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_FRI
		return
	case "SAT":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_SAT
		return
	case "SUN":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_SUN
		return
	case "WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_WORK_DAY
		return
	case "BEF_WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY
		return
	case "AFT_WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY
		return
	case "HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_HOL
		return
	case "BEF_HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BEF_HOL
		return
	case "AFT_HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_AFT_HOL
		return
	case "ANY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_ANY_XSD
		return
	case "BUSY_FRI":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BUSY_FRI
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedaybasetype_member_1 *Enum_CodeDayBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDayBaseType_Member_1_MON":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_MON
	case "Enum_CodeDayBaseType_Member_1_TUE":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_TUE
	case "Enum_CodeDayBaseType_Member_1_WED":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_WED
	case "Enum_CodeDayBaseType_Member_1_THU":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_THU
	case "Enum_CodeDayBaseType_Member_1_FRI":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_FRI
	case "Enum_CodeDayBaseType_Member_1_SAT":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_SAT
	case "Enum_CodeDayBaseType_Member_1_SUN":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_SUN
	case "Enum_CodeDayBaseType_Member_1_WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_WORK_DAY
	case "Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY
	case "Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY
	case "Enum_CodeDayBaseType_Member_1_HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_HOL
	case "Enum_CodeDayBaseType_Member_1_BEF_HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BEF_HOL
	case "Enum_CodeDayBaseType_Member_1_AFT_HOL":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_AFT_HOL
	case "Enum_CodeDayBaseType_Member_1_ANY_XSD":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_ANY_XSD
	case "Enum_CodeDayBaseType_Member_1_BUSY_FRI":
		*enum_codedaybasetype_member_1 = Enum_CodeDayBaseType_Member_1_BUSY_FRI
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedaybasetype_member_1 *Enum_CodeDayBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDayBaseType_Member_1_MON:
		res = "Enum_CodeDayBaseType_Member_1_MON"
	case Enum_CodeDayBaseType_Member_1_TUE:
		res = "Enum_CodeDayBaseType_Member_1_TUE"
	case Enum_CodeDayBaseType_Member_1_WED:
		res = "Enum_CodeDayBaseType_Member_1_WED"
	case Enum_CodeDayBaseType_Member_1_THU:
		res = "Enum_CodeDayBaseType_Member_1_THU"
	case Enum_CodeDayBaseType_Member_1_FRI:
		res = "Enum_CodeDayBaseType_Member_1_FRI"
	case Enum_CodeDayBaseType_Member_1_SAT:
		res = "Enum_CodeDayBaseType_Member_1_SAT"
	case Enum_CodeDayBaseType_Member_1_SUN:
		res = "Enum_CodeDayBaseType_Member_1_SUN"
	case Enum_CodeDayBaseType_Member_1_WORK_DAY:
		res = "Enum_CodeDayBaseType_Member_1_WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY:
		res = "Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY:
		res = "Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY"
	case Enum_CodeDayBaseType_Member_1_HOL:
		res = "Enum_CodeDayBaseType_Member_1_HOL"
	case Enum_CodeDayBaseType_Member_1_BEF_HOL:
		res = "Enum_CodeDayBaseType_Member_1_BEF_HOL"
	case Enum_CodeDayBaseType_Member_1_AFT_HOL:
		res = "Enum_CodeDayBaseType_Member_1_AFT_HOL"
	case Enum_CodeDayBaseType_Member_1_ANY_XSD:
		res = "Enum_CodeDayBaseType_Member_1_ANY_XSD"
	case Enum_CodeDayBaseType_Member_1_BUSY_FRI:
		res = "Enum_CodeDayBaseType_Member_1_BUSY_FRI"
	}
	return
}

func (enum_codedaybasetype_member_1 Enum_CodeDayBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDayBaseType_Member_1_MON")
	res = append(res, "Enum_CodeDayBaseType_Member_1_TUE")
	res = append(res, "Enum_CodeDayBaseType_Member_1_WED")
	res = append(res, "Enum_CodeDayBaseType_Member_1_THU")
	res = append(res, "Enum_CodeDayBaseType_Member_1_FRI")
	res = append(res, "Enum_CodeDayBaseType_Member_1_SAT")
	res = append(res, "Enum_CodeDayBaseType_Member_1_SUN")
	res = append(res, "Enum_CodeDayBaseType_Member_1_WORK_DAY")
	res = append(res, "Enum_CodeDayBaseType_Member_1_BEF_WORK_DAY")
	res = append(res, "Enum_CodeDayBaseType_Member_1_AFT_WORK_DAY")
	res = append(res, "Enum_CodeDayBaseType_Member_1_HOL")
	res = append(res, "Enum_CodeDayBaseType_Member_1_BEF_HOL")
	res = append(res, "Enum_CodeDayBaseType_Member_1_AFT_HOL")
	res = append(res, "Enum_CodeDayBaseType_Member_1_ANY_XSD")
	res = append(res, "Enum_CodeDayBaseType_Member_1_BUSY_FRI")

	return
}

func (enum_codedaybasetype_member_1 Enum_CodeDayBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MON")
	res = append(res, "TUE")
	res = append(res, "WED")
	res = append(res, "THU")
	res = append(res, "FRI")
	res = append(res, "SAT")
	res = append(res, "SUN")
	res = append(res, "WORK_DAY")
	res = append(res, "BEF_WORK_DAY")
	res = append(res, "AFT_WORK_DAY")
	res = append(res, "HOL")
	res = append(res, "BEF_HOL")
	res = append(res, "AFT_HOL")
	res = append(res, "ANY")
	res = append(res, "BUSY_FRI")

	return
}

// Utility function for Enum_CodeGroundLightingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codegroundlightingbasetype_member_1 Enum_CodeGroundLightingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codegroundlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGroundLightingBaseType_Member_1_BCN:
		res = "BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_IBN:
		res = "IBN"
	case Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN:
		res = "HEL_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_ABN:
		res = "ABN"
	case Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN:
		res = "MAR_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN:
		res = "RSP_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN:
		res = "TWR_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN:
		res = "HAZ_BCN"
	}
	return
}

func (enum_codegroundlightingbasetype_member_1 *Enum_CodeGroundLightingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_BCN
		return
	case "IBN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_IBN
		return
	case "HEL_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN
		return
	case "ABN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_ABN
		return
	case "MAR_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN
		return
	case "RSP_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN
		return
	case "TWR_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN
		return
	case "HAZ_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codegroundlightingbasetype_member_1 *Enum_CodeGroundLightingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeGroundLightingBaseType_Member_1_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_BCN
	case "Enum_CodeGroundLightingBaseType_Member_1_IBN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_IBN
	case "Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN
	case "Enum_CodeGroundLightingBaseType_Member_1_ABN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_ABN
	case "Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN
	case "Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN
	case "Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN
	case "Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN":
		*enum_codegroundlightingbasetype_member_1 = Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codegroundlightingbasetype_member_1 *Enum_CodeGroundLightingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codegroundlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGroundLightingBaseType_Member_1_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_IBN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_IBN"
	case Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_ABN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_ABN"
	case Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN"
	case Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN:
		res = "Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN"
	}
	return
}

func (enum_codegroundlightingbasetype_member_1 Enum_CodeGroundLightingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_BCN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_IBN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_HEL_BCN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_ABN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_MAR_BCN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_RSP_BCN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_TWR_BCN")
	res = append(res, "Enum_CodeGroundLightingBaseType_Member_1_HAZ_BCN")

	return
}

func (enum_codegroundlightingbasetype_member_1 Enum_CodeGroundLightingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BCN")
	res = append(res, "IBN")
	res = append(res, "HEL_BCN")
	res = append(res, "ABN")
	res = append(res, "MAR_BCN")
	res = append(res, "RSP_BCN")
	res = append(res, "TWR_BCN")
	res = append(res, "HAZ_BCN")

	return
}

// Utility function for Enum_CodeLightIntensityBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelightintensitybasetype_member_1 Enum_CodeLightIntensityBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelightintensitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightIntensityBaseType_Member_1_LIL:
		res = "LIL"
	case Enum_CodeLightIntensityBaseType_Member_1_LIM:
		res = "LIM"
	case Enum_CodeLightIntensityBaseType_Member_1_LIH:
		res = "LIH"
	case Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH:
		res = "LIL_LIH"
	case Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED:
		res = "PREDETERMINED"
	}
	return
}

func (enum_codelightintensitybasetype_member_1 *Enum_CodeLightIntensityBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LIL":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIL
		return
	case "LIM":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIM
		return
	case "LIH":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIH
		return
	case "LIL_LIH":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH
		return
	case "PREDETERMINED":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelightintensitybasetype_member_1 *Enum_CodeLightIntensityBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLightIntensityBaseType_Member_1_LIL":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIL
	case "Enum_CodeLightIntensityBaseType_Member_1_LIM":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIM
	case "Enum_CodeLightIntensityBaseType_Member_1_LIH":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIH
	case "Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH
	case "Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED":
		*enum_codelightintensitybasetype_member_1 = Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelightintensitybasetype_member_1 *Enum_CodeLightIntensityBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelightintensitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightIntensityBaseType_Member_1_LIL:
		res = "Enum_CodeLightIntensityBaseType_Member_1_LIL"
	case Enum_CodeLightIntensityBaseType_Member_1_LIM:
		res = "Enum_CodeLightIntensityBaseType_Member_1_LIM"
	case Enum_CodeLightIntensityBaseType_Member_1_LIH:
		res = "Enum_CodeLightIntensityBaseType_Member_1_LIH"
	case Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH:
		res = "Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH"
	case Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED:
		res = "Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED"
	}
	return
}

func (enum_codelightintensitybasetype_member_1 Enum_CodeLightIntensityBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLightIntensityBaseType_Member_1_LIL")
	res = append(res, "Enum_CodeLightIntensityBaseType_Member_1_LIM")
	res = append(res, "Enum_CodeLightIntensityBaseType_Member_1_LIH")
	res = append(res, "Enum_CodeLightIntensityBaseType_Member_1_LIL_LIH")
	res = append(res, "Enum_CodeLightIntensityBaseType_Member_1_PREDETERMINED")

	return
}

func (enum_codelightintensitybasetype_member_1 Enum_CodeLightIntensityBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LIL")
	res = append(res, "LIM")
	res = append(res, "LIH")
	res = append(res, "LIL_LIH")
	res = append(res, "PREDETERMINED")

	return
}

// Utility function for Enum_CodeLightSourceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelightsourcebasetype_member_1 Enum_CodeLightSourceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelightsourcebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightSourceBaseType_Member_1_FLOOD:
		res = "FLOOD"
	case Enum_CodeLightSourceBaseType_Member_1_STROBE:
		res = "STROBE"
	}
	return
}

func (enum_codelightsourcebasetype_member_1 *Enum_CodeLightSourceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FLOOD":
		*enum_codelightsourcebasetype_member_1 = Enum_CodeLightSourceBaseType_Member_1_FLOOD
		return
	case "STROBE":
		*enum_codelightsourcebasetype_member_1 = Enum_CodeLightSourceBaseType_Member_1_STROBE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelightsourcebasetype_member_1 *Enum_CodeLightSourceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLightSourceBaseType_Member_1_FLOOD":
		*enum_codelightsourcebasetype_member_1 = Enum_CodeLightSourceBaseType_Member_1_FLOOD
	case "Enum_CodeLightSourceBaseType_Member_1_STROBE":
		*enum_codelightsourcebasetype_member_1 = Enum_CodeLightSourceBaseType_Member_1_STROBE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelightsourcebasetype_member_1 *Enum_CodeLightSourceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelightsourcebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightSourceBaseType_Member_1_FLOOD:
		res = "Enum_CodeLightSourceBaseType_Member_1_FLOOD"
	case Enum_CodeLightSourceBaseType_Member_1_STROBE:
		res = "Enum_CodeLightSourceBaseType_Member_1_STROBE"
	}
	return
}

func (enum_codelightsourcebasetype_member_1 Enum_CodeLightSourceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLightSourceBaseType_Member_1_FLOOD")
	res = append(res, "Enum_CodeLightSourceBaseType_Member_1_STROBE")

	return
}

func (enum_codelightsourcebasetype_member_1 Enum_CodeLightSourceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FLOOD")
	res = append(res, "STROBE")

	return
}

// Utility function for Enum_CodeNotePurposeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenotepurposebasetype_member_1 Enum_CodeNotePurposeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenotepurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION:
		res = "DESCRIPTION"
	case Enum_CodeNotePurposeBaseType_Member_1_REMARK:
		res = "REMARK"
	case Enum_CodeNotePurposeBaseType_Member_1_WARNING:
		res = "WARNING"
	case Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER:
		res = "DISCLAIMER"
	}
	return
}

func (enum_codenotepurposebasetype_member_1 *Enum_CodeNotePurposeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DESCRIPTION":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION
		return
	case "REMARK":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_REMARK
		return
	case "WARNING":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_WARNING
		return
	case "DISCLAIMER":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenotepurposebasetype_member_1 *Enum_CodeNotePurposeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION
	case "Enum_CodeNotePurposeBaseType_Member_1_REMARK":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_REMARK
	case "Enum_CodeNotePurposeBaseType_Member_1_WARNING":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_WARNING
	case "Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER":
		*enum_codenotepurposebasetype_member_1 = Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenotepurposebasetype_member_1 *Enum_CodeNotePurposeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenotepurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION:
		res = "Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION"
	case Enum_CodeNotePurposeBaseType_Member_1_REMARK:
		res = "Enum_CodeNotePurposeBaseType_Member_1_REMARK"
	case Enum_CodeNotePurposeBaseType_Member_1_WARNING:
		res = "Enum_CodeNotePurposeBaseType_Member_1_WARNING"
	case Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER:
		res = "Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER"
	}
	return
}

func (enum_codenotepurposebasetype_member_1 Enum_CodeNotePurposeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNotePurposeBaseType_Member_1_DESCRIPTION")
	res = append(res, "Enum_CodeNotePurposeBaseType_Member_1_REMARK")
	res = append(res, "Enum_CodeNotePurposeBaseType_Member_1_WARNING")
	res = append(res, "Enum_CodeNotePurposeBaseType_Member_1_DISCLAIMER")

	return
}

func (enum_codenotepurposebasetype_member_1 Enum_CodeNotePurposeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DESCRIPTION")
	res = append(res, "REMARK")
	res = append(res, "WARNING")
	res = append(res, "DISCLAIMER")

	return
}

// Utility function for Enum_CodeObstacleAreaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeobstacleareabasetype_member_1 Enum_CodeObstacleAreaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeobstacleareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA1:
		res = "AREA1"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA2:
		res = "AREA2"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA3:
		res = "AREA3"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA4:
		res = "AREA4"
	case Enum_CodeObstacleAreaBaseType_Member_1_OLS:
		res = "OLS"
	case Enum_CodeObstacleAreaBaseType_Member_1_FAR77:
		res = "FAR77"
	case Enum_CodeObstacleAreaBaseType_Member_1_MANAGED:
		res = "MANAGED"
	}
	return
}

func (enum_codeobstacleareabasetype_member_1 *Enum_CodeObstacleAreaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AREA1":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA1
		return
	case "AREA2":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA2
		return
	case "AREA3":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA3
		return
	case "AREA4":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA4
		return
	case "OLS":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_OLS
		return
	case "FAR77":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_FAR77
		return
	case "MANAGED":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_MANAGED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeobstacleareabasetype_member_1 *Enum_CodeObstacleAreaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeObstacleAreaBaseType_Member_1_AREA1":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA1
	case "Enum_CodeObstacleAreaBaseType_Member_1_AREA2":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA2
	case "Enum_CodeObstacleAreaBaseType_Member_1_AREA3":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA3
	case "Enum_CodeObstacleAreaBaseType_Member_1_AREA4":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_AREA4
	case "Enum_CodeObstacleAreaBaseType_Member_1_OLS":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_OLS
	case "Enum_CodeObstacleAreaBaseType_Member_1_FAR77":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_FAR77
	case "Enum_CodeObstacleAreaBaseType_Member_1_MANAGED":
		*enum_codeobstacleareabasetype_member_1 = Enum_CodeObstacleAreaBaseType_Member_1_MANAGED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeobstacleareabasetype_member_1 *Enum_CodeObstacleAreaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeobstacleareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA1:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_AREA1"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA2:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_AREA2"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA3:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_AREA3"
	case Enum_CodeObstacleAreaBaseType_Member_1_AREA4:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_AREA4"
	case Enum_CodeObstacleAreaBaseType_Member_1_OLS:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_OLS"
	case Enum_CodeObstacleAreaBaseType_Member_1_FAR77:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_FAR77"
	case Enum_CodeObstacleAreaBaseType_Member_1_MANAGED:
		res = "Enum_CodeObstacleAreaBaseType_Member_1_MANAGED"
	}
	return
}

func (enum_codeobstacleareabasetype_member_1 Enum_CodeObstacleAreaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_AREA1")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_AREA2")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_AREA3")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_AREA4")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_OLS")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_FAR77")
	res = append(res, "Enum_CodeObstacleAreaBaseType_Member_1_MANAGED")

	return
}

func (enum_codeobstacleareabasetype_member_1 Enum_CodeObstacleAreaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AREA1")
	res = append(res, "AREA2")
	res = append(res, "AREA3")
	res = append(res, "AREA4")
	res = append(res, "OLS")
	res = append(res, "FAR77")
	res = append(res, "MANAGED")

	return
}

// Utility function for Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeobstacleassessmentsurfacebasetype_member_1 Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeobstacleassessmentsurfacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1:
		res = "40_TO_1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1:
		res = "72_TO_1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA:
		res = "MA"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL:
		res = "FINAL"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA:
		res = "PT_ENTRY_AREA"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY:
		res = "PRIMARY"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY:
		res = "SECONDARY"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1:
		res = "ZONE1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2:
		res = "ZONE2"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3:
		res = "ZONE3"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1:
		res = "AREA1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2:
		res = "AREA2"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3:
		res = "AREA3"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION:
		res = "TURN_INITIATION"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN:
		res = "TURN"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER:
		res = "DER"
	}
	return
}

func (enum_codeobstacleassessmentsurfacebasetype_member_1 *Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "40_TO_1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1
		return
	case "72_TO_1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1
		return
	case "MA":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA
		return
	case "FINAL":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL
		return
	case "PT_ENTRY_AREA":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA
		return
	case "PRIMARY":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY
		return
	case "SECONDARY":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY
		return
	case "ZONE1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1
		return
	case "ZONE2":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2
		return
	case "ZONE3":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3
		return
	case "AREA1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1
		return
	case "AREA2":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2
		return
	case "AREA3":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3
		return
	case "TURN_INITIATION":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION
		return
	case "TURN":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN
		return
	case "DER":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeobstacleassessmentsurfacebasetype_member_1 *Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN
	case "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER":
		*enum_codeobstacleassessmentsurfacebasetype_member_1 = Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeobstacleassessmentsurfacebasetype_member_1 *Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeobstacleassessmentsurfacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN"
	case Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER:
		res = "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER"
	}
	return
}

func (enum_codeobstacleassessmentsurfacebasetype_member_1 Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_40_TO_1")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_N_72_TO_1")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_MA")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_FINAL")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PT_ENTRY_AREA")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_PRIMARY")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_SECONDARY")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE1")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE2")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_ZONE3")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA1")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA2")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_AREA3")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN_INITIATION")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_TURN")
	res = append(res, "Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1_DER")

	return
}

func (enum_codeobstacleassessmentsurfacebasetype_member_1 Enum_CodeObstacleAssessmentSurfaceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "40_TO_1")
	res = append(res, "72_TO_1")
	res = append(res, "MA")
	res = append(res, "FINAL")
	res = append(res, "PT_ENTRY_AREA")
	res = append(res, "PRIMARY")
	res = append(res, "SECONDARY")
	res = append(res, "ZONE1")
	res = append(res, "ZONE2")
	res = append(res, "ZONE3")
	res = append(res, "AREA1")
	res = append(res, "AREA2")
	res = append(res, "AREA3")
	res = append(res, "TURN_INITIATION")
	res = append(res, "TURN")
	res = append(res, "DER")

	return
}

// Utility function for Enum_CodeStatusConstructionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusconstructionbasetype_member_1 Enum_CodeStatusConstructionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusconstructionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION:
		res = "IN_CONSTRUCTION"
	case Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED:
		res = "COMPLETED"
	case Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED:
		res = "DEMOLITION_PLANNED"
	case Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION:
		res = "IN_DEMOLITION"
	}
	return
}

func (enum_codestatusconstructionbasetype_member_1 *Enum_CodeStatusConstructionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IN_CONSTRUCTION":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION
		return
	case "COMPLETED":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED
		return
	case "DEMOLITION_PLANNED":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED
		return
	case "IN_DEMOLITION":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusconstructionbasetype_member_1 *Enum_CodeStatusConstructionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION
	case "Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED
	case "Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED
	case "Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION":
		*enum_codestatusconstructionbasetype_member_1 = Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusconstructionbasetype_member_1 *Enum_CodeStatusConstructionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusconstructionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION:
		res = "Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION"
	case Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED:
		res = "Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED"
	case Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED:
		res = "Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED"
	case Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION:
		res = "Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION"
	}
	return
}

func (enum_codestatusconstructionbasetype_member_1 Enum_CodeStatusConstructionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusConstructionBaseType_Member_1_IN_CONSTRUCTION")
	res = append(res, "Enum_CodeStatusConstructionBaseType_Member_1_COMPLETED")
	res = append(res, "Enum_CodeStatusConstructionBaseType_Member_1_DEMOLITION_PLANNED")
	res = append(res, "Enum_CodeStatusConstructionBaseType_Member_1_IN_DEMOLITION")

	return
}

func (enum_codestatusconstructionbasetype_member_1 Enum_CodeStatusConstructionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IN_CONSTRUCTION")
	res = append(res, "COMPLETED")
	res = append(res, "DEMOLITION_PLANNED")
	res = append(res, "IN_DEMOLITION")

	return
}

// Utility function for Enum_CodeStatusOperationsBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusoperationsbasetype_member_1 Enum_CodeStatusOperationsBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusoperationsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusOperationsBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED:
		res = "DOWNGRADED"
	case Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE:
		res = "UNSERVICEABLE"
	case Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS:
		res = "WORK_IN_PROGRESS"
	}
	return
}

func (enum_codestatusoperationsbasetype_member_1 *Enum_CodeStatusOperationsBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_NORMAL
		return
	case "DOWNGRADED":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED
		return
	case "UNSERVICEABLE":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE
		return
	case "WORK_IN_PROGRESS":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusoperationsbasetype_member_1 *Enum_CodeStatusOperationsBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusOperationsBaseType_Member_1_NORMAL":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_NORMAL
	case "Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED
	case "Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE
	case "Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS":
		*enum_codestatusoperationsbasetype_member_1 = Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusoperationsbasetype_member_1 *Enum_CodeStatusOperationsBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusoperationsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusOperationsBaseType_Member_1_NORMAL:
		res = "Enum_CodeStatusOperationsBaseType_Member_1_NORMAL"
	case Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED:
		res = "Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED"
	case Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE:
		res = "Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE"
	case Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS:
		res = "Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS"
	}
	return
}

func (enum_codestatusoperationsbasetype_member_1 Enum_CodeStatusOperationsBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusOperationsBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeStatusOperationsBaseType_Member_1_DOWNGRADED")
	res = append(res, "Enum_CodeStatusOperationsBaseType_Member_1_UNSERVICEABLE")
	res = append(res, "Enum_CodeStatusOperationsBaseType_Member_1_WORK_IN_PROGRESS")

	return
}

func (enum_codestatusoperationsbasetype_member_1 Enum_CodeStatusOperationsBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "DOWNGRADED")
	res = append(res, "UNSERVICEABLE")
	res = append(res, "WORK_IN_PROGRESS")

	return
}

// Utility function for Enum_CodeTimeEventBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetimeeventbasetype_member_1 Enum_CodeTimeEventBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetimeeventbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeEventBaseType_Member_1_SR:
		res = "SR"
	case Enum_CodeTimeEventBaseType_Member_1_SS:
		res = "SS"
	}
	return
}

func (enum_codetimeeventbasetype_member_1 *Enum_CodeTimeEventBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SR":
		*enum_codetimeeventbasetype_member_1 = Enum_CodeTimeEventBaseType_Member_1_SR
		return
	case "SS":
		*enum_codetimeeventbasetype_member_1 = Enum_CodeTimeEventBaseType_Member_1_SS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetimeeventbasetype_member_1 *Enum_CodeTimeEventBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTimeEventBaseType_Member_1_SR":
		*enum_codetimeeventbasetype_member_1 = Enum_CodeTimeEventBaseType_Member_1_SR
	case "Enum_CodeTimeEventBaseType_Member_1_SS":
		*enum_codetimeeventbasetype_member_1 = Enum_CodeTimeEventBaseType_Member_1_SS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetimeeventbasetype_member_1 *Enum_CodeTimeEventBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetimeeventbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeEventBaseType_Member_1_SR:
		res = "Enum_CodeTimeEventBaseType_Member_1_SR"
	case Enum_CodeTimeEventBaseType_Member_1_SS:
		res = "Enum_CodeTimeEventBaseType_Member_1_SS"
	}
	return
}

func (enum_codetimeeventbasetype_member_1 Enum_CodeTimeEventBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTimeEventBaseType_Member_1_SR")
	res = append(res, "Enum_CodeTimeEventBaseType_Member_1_SS")

	return
}

func (enum_codetimeeventbasetype_member_1 Enum_CodeTimeEventBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SR")
	res = append(res, "SS")

	return
}

// Utility function for Enum_CodeTimeEventCombinationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetimeeventcombinationbasetype_member_1 Enum_CodeTimeEventCombinationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetimeeventcombinationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST:
		res = "EARLIEST"
	case Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST:
		res = "LATEST"
	}
	return
}

func (enum_codetimeeventcombinationbasetype_member_1 *Enum_CodeTimeEventCombinationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EARLIEST":
		*enum_codetimeeventcombinationbasetype_member_1 = Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST
		return
	case "LATEST":
		*enum_codetimeeventcombinationbasetype_member_1 = Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetimeeventcombinationbasetype_member_1 *Enum_CodeTimeEventCombinationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST":
		*enum_codetimeeventcombinationbasetype_member_1 = Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST
	case "Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST":
		*enum_codetimeeventcombinationbasetype_member_1 = Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetimeeventcombinationbasetype_member_1 *Enum_CodeTimeEventCombinationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetimeeventcombinationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST:
		res = "Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST"
	case Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST:
		res = "Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST"
	}
	return
}

func (enum_codetimeeventcombinationbasetype_member_1 Enum_CodeTimeEventCombinationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTimeEventCombinationBaseType_Member_1_EARLIEST")
	res = append(res, "Enum_CodeTimeEventCombinationBaseType_Member_1_LATEST")

	return
}

func (enum_codetimeeventcombinationbasetype_member_1 Enum_CodeTimeEventCombinationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EARLIEST")
	res = append(res, "LATEST")

	return
}

// Utility function for Enum_CodeTimeReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetimereferencebasetype_member_1 Enum_CodeTimeReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetimereferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC:
		res = "UTC"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_12:
		res = "UTC-12"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_11:
		res = "UTC-11"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_10:
		res = "UTC-10"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_9:
		res = "UTC-9"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_8:
		res = "UTC-8"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_7:
		res = "UTC-7"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_6:
		res = "UTC-6"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_5:
		res = "UTC-5"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_4:
		res = "UTC-4"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_3:
		res = "UTC-3"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_2:
		res = "UTC-2"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_1:
		res = "UTC-1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1:
		res = "UTC+1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1:
		res = "UTC+2"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1:
		res = "UTC+3"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1:
		res = "UTC+4"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1:
		res = "UTC+5"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1:
		res = "UTC+6"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1:
		res = "UTC+7"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1:
		res = "UTC+8"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1:
		res = "UTC+9"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1:
		res = "UTC+10"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1:
		res = "UTC+11"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1:
		res = "UTC+12"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_13:
		res = "UTC+13"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_14:
		res = "UTC+14"
	}
	return
}

func (enum_codetimereferencebasetype_member_1 *Enum_CodeTimeReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "UTC":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC
		return
	case "UTC-12":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_12
		return
	case "UTC-11":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_11
		return
	case "UTC-10":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_10
		return
	case "UTC-9":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_9
		return
	case "UTC-8":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_8
		return
	case "UTC-7":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_7
		return
	case "UTC-6":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_6
		return
	case "UTC-5":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_5
		return
	case "UTC-4":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_4
		return
	case "UTC-3":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_3
		return
	case "UTC-2":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_2
		return
	case "UTC-1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_1
		return
	case "UTC+1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1
		return
	case "UTC+2":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1
		return
	case "UTC+3":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1
		return
	case "UTC+4":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1
		return
	case "UTC+5":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1
		return
	case "UTC+6":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1
		return
	case "UTC+7":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1
		return
	case "UTC+8":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1
		return
	case "UTC+9":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1
		return
	case "UTC+10":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1
		return
	case "UTC+11":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1
		return
	case "UTC+12":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1
		return
	case "UTC+13":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_13
		return
	case "UTC+14":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_14
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetimereferencebasetype_member_1 *Enum_CodeTimeReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_12
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_11
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_10
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_9
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_8
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_7
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_6
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_5
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_4
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_3
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_2
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_13":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_13
	case "Enum_CodeTimeReferenceBaseType_Member_1_UTC_14":
		*enum_codetimereferencebasetype_member_1 = Enum_CodeTimeReferenceBaseType_Member_1_UTC_14
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetimereferencebasetype_member_1 *Enum_CodeTimeReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetimereferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_12:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_11:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_10:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_9:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_8:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_7:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_6:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_5:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_4:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_3:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_2:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_13:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_13"
	case Enum_CodeTimeReferenceBaseType_Member_1_UTC_14:
		res = "Enum_CodeTimeReferenceBaseType_Member_1_UTC_14"
	}
	return
}

func (enum_codetimereferencebasetype_member_1 Enum_CodeTimeReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_1_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_2_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_3_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_4_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_5_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_6_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_7_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_8_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_9_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_10_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_11_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_12_1")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_13")
	res = append(res, "Enum_CodeTimeReferenceBaseType_Member_1_UTC_14")

	return
}

func (enum_codetimereferencebasetype_member_1 Enum_CodeTimeReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "UTC")
	res = append(res, "UTC-12")
	res = append(res, "UTC-11")
	res = append(res, "UTC-10")
	res = append(res, "UTC-9")
	res = append(res, "UTC-8")
	res = append(res, "UTC-7")
	res = append(res, "UTC-6")
	res = append(res, "UTC-5")
	res = append(res, "UTC-4")
	res = append(res, "UTC-3")
	res = append(res, "UTC-2")
	res = append(res, "UTC-1")
	res = append(res, "UTC+1")
	res = append(res, "UTC+2")
	res = append(res, "UTC+3")
	res = append(res, "UTC+4")
	res = append(res, "UTC+5")
	res = append(res, "UTC+6")
	res = append(res, "UTC+7")
	res = append(res, "UTC+8")
	res = append(res, "UTC+9")
	res = append(res, "UTC+10")
	res = append(res, "UTC+11")
	res = append(res, "UTC+12")
	res = append(res, "UTC+13")
	res = append(res, "UTC+14")

	return
}

// Utility function for Enum_CodeVerticalDatumBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeverticaldatumbasetype_member_1 Enum_CodeVerticalDatumBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeverticaldatumbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalDatumBaseType_Member_1_EGM_96:
		res = "EGM_96"
	case Enum_CodeVerticalDatumBaseType_Member_1_AHD:
		res = "AHD"
	case Enum_CodeVerticalDatumBaseType_Member_1_NAVD88:
		res = "NAVD88"
	}
	return
}

func (enum_codeverticaldatumbasetype_member_1 *Enum_CodeVerticalDatumBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EGM_96":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_EGM_96
		return
	case "AHD":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_AHD
		return
	case "NAVD88":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_NAVD88
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeverticaldatumbasetype_member_1 *Enum_CodeVerticalDatumBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVerticalDatumBaseType_Member_1_EGM_96":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_EGM_96
	case "Enum_CodeVerticalDatumBaseType_Member_1_AHD":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_AHD
	case "Enum_CodeVerticalDatumBaseType_Member_1_NAVD88":
		*enum_codeverticaldatumbasetype_member_1 = Enum_CodeVerticalDatumBaseType_Member_1_NAVD88
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeverticaldatumbasetype_member_1 *Enum_CodeVerticalDatumBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeverticaldatumbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalDatumBaseType_Member_1_EGM_96:
		res = "Enum_CodeVerticalDatumBaseType_Member_1_EGM_96"
	case Enum_CodeVerticalDatumBaseType_Member_1_AHD:
		res = "Enum_CodeVerticalDatumBaseType_Member_1_AHD"
	case Enum_CodeVerticalDatumBaseType_Member_1_NAVD88:
		res = "Enum_CodeVerticalDatumBaseType_Member_1_NAVD88"
	}
	return
}

func (enum_codeverticaldatumbasetype_member_1 Enum_CodeVerticalDatumBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVerticalDatumBaseType_Member_1_EGM_96")
	res = append(res, "Enum_CodeVerticalDatumBaseType_Member_1_AHD")
	res = append(res, "Enum_CodeVerticalDatumBaseType_Member_1_NAVD88")

	return
}

func (enum_codeverticaldatumbasetype_member_1 Enum_CodeVerticalDatumBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EGM_96")
	res = append(res, "AHD")
	res = append(res, "NAVD88")

	return
}

// Utility function for Enum_CodeVerticalStructureBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeverticalstructurebasetype_member_1 Enum_CodeVerticalStructureBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeverticalstructurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP:
		res = "AG_EQUIP"
	case Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA:
		res = "ANTENNA"
	case Enum_CodeVerticalStructureBaseType_Member_1_ARCH:
		res = "ARCH"
	case Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE:
		res = "BRIDGE"
	case Enum_CodeVerticalStructureBaseType_Member_1_BUILDING:
		res = "BUILDING"
	case Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR:
		res = "CABLE_CAR"
	case Enum_CodeVerticalStructureBaseType_Member_1_CATENARY:
		res = "CATENARY"
	case Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM:
		res = "COMPRESSED_AIR_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM:
		res = "CONTROL_MONITORING_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER:
		res = "CONTROL_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER:
		res = "COOLING_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_CRANE:
		res = "CRANE"
	case Enum_CodeVerticalStructureBaseType_Member_1_DAM:
		res = "DAM"
	case Enum_CodeVerticalStructureBaseType_Member_1_DOME:
		res = "DOME"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT:
		res = "ELECTRICAL_EXIT_LIGHT"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM:
		res = "ELECTRICAL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR:
		res = "ELEVATOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_FENCE:
		res = "FENCE"
	case Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM:
		res = "FUEL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_GATE:
		res = "GATE"
	case Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY:
		res = "GENERAL_UTILITY"
	case Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR:
		res = "GRAIN_ELEVATOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM:
		res = "HEAT_COOL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM:
		res = "INDUSTRIAL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE:
		res = "LIGHTHOUSE"
	case Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT:
		res = "MONUMENT"
	case Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM:
		res = "NATURAL_GAS_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT:
		res = "NATURAL_HIGHPOINT"
	case Enum_CodeVerticalStructureBaseType_Member_1_NAVAID:
		res = "NAVAID"
	case Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR:
		res = "NUCLEAR_REACTOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_POLE:
		res = "POLE"
	case Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT:
		res = "POWER_PLANT"
	case Enum_CodeVerticalStructureBaseType_Member_1_REFINERY:
		res = "REFINERY"
	case Enum_CodeVerticalStructureBaseType_Member_1_RIG:
		res = "RIG"
	case Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM:
		res = "SALTWATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_SIGN:
		res = "SIGN"
	case Enum_CodeVerticalStructureBaseType_Member_1_SPIRE:
		res = "SPIRE"
	case Enum_CodeVerticalStructureBaseType_Member_1_STACK:
		res = "STACK"
	case Enum_CodeVerticalStructureBaseType_Member_1_STADIUM:
		res = "STADIUM"
	case Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM:
		res = "STORM_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_TANK:
		res = "TANK"
	case Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON:
		res = "TETHERED_BALLOON"
	case Enum_CodeVerticalStructureBaseType_Member_1_TOWER:
		res = "TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY:
		res = "TRAMWAY"
	case Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE:
		res = "TRANSMISSION_LINE"
	case Enum_CodeVerticalStructureBaseType_Member_1_TREE:
		res = "TREE"
	case Enum_CodeVerticalStructureBaseType_Member_1_URBAN:
		res = "URBAN"
	case Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION:
		res = "VEGETATION"
	case Enum_CodeVerticalStructureBaseType_Member_1_WALL:
		res = "WALL"
	case Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM:
		res = "WASTEWATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM:
		res = "WATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER:
		res = "WATER_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL:
		res = "WINDMILL"
	case Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS:
		res = "WINDMILL_FARMS"
	}
	return
}

func (enum_codeverticalstructurebasetype_member_1 *Enum_CodeVerticalStructureBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AG_EQUIP":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP
		return
	case "ANTENNA":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA
		return
	case "ARCH":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ARCH
		return
	case "BRIDGE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE
		return
	case "BUILDING":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_BUILDING
		return
	case "CABLE_CAR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR
		return
	case "CATENARY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CATENARY
		return
	case "COMPRESSED_AIR_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM
		return
	case "CONTROL_MONITORING_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM
		return
	case "CONTROL_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER
		return
	case "COOLING_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER
		return
	case "CRANE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CRANE
		return
	case "DAM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_DAM
		return
	case "DOME":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_DOME
		return
	case "ELECTRICAL_EXIT_LIGHT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT
		return
	case "ELECTRICAL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM
		return
	case "ELEVATOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR
		return
	case "FENCE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_FENCE
		return
	case "FUEL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM
		return
	case "GATE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GATE
		return
	case "GENERAL_UTILITY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY
		return
	case "GRAIN_ELEVATOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR
		return
	case "HEAT_COOL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM
		return
	case "INDUSTRIAL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM
		return
	case "LIGHTHOUSE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE
		return
	case "MONUMENT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT
		return
	case "NATURAL_GAS_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM
		return
	case "NATURAL_HIGHPOINT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT
		return
	case "NAVAID":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NAVAID
		return
	case "NUCLEAR_REACTOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR
		return
	case "POLE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_POLE
		return
	case "POWER_PLANT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT
		return
	case "REFINERY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_REFINERY
		return
	case "RIG":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_RIG
		return
	case "SALTWATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM
		return
	case "SIGN":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SIGN
		return
	case "SPIRE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SPIRE
		return
	case "STACK":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STACK
		return
	case "STADIUM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STADIUM
		return
	case "STORM_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM
		return
	case "TANK":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TANK
		return
	case "TETHERED_BALLOON":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON
		return
	case "TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TOWER
		return
	case "TRAMWAY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY
		return
	case "TRANSMISSION_LINE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE
		return
	case "TREE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TREE
		return
	case "URBAN":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_URBAN
		return
	case "VEGETATION":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION
		return
	case "WALL":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WALL
		return
	case "WASTEWATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM
		return
	case "WATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM
		return
	case "WATER_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER
		return
	case "WINDMILL":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL
		return
	case "WINDMILL_FARMS":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeverticalstructurebasetype_member_1 *Enum_CodeVerticalStructureBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP
	case "Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA
	case "Enum_CodeVerticalStructureBaseType_Member_1_ARCH":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ARCH
	case "Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE
	case "Enum_CodeVerticalStructureBaseType_Member_1_BUILDING":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_BUILDING
	case "Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR
	case "Enum_CodeVerticalStructureBaseType_Member_1_CATENARY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CATENARY
	case "Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER
	case "Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER
	case "Enum_CodeVerticalStructureBaseType_Member_1_CRANE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_CRANE
	case "Enum_CodeVerticalStructureBaseType_Member_1_DAM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_DAM
	case "Enum_CodeVerticalStructureBaseType_Member_1_DOME":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_DOME
	case "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT
	case "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR
	case "Enum_CodeVerticalStructureBaseType_Member_1_FENCE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_FENCE
	case "Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_GATE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GATE
	case "Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY
	case "Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR
	case "Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE
	case "Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT
	case "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT
	case "Enum_CodeVerticalStructureBaseType_Member_1_NAVAID":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NAVAID
	case "Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR
	case "Enum_CodeVerticalStructureBaseType_Member_1_POLE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_POLE
	case "Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT
	case "Enum_CodeVerticalStructureBaseType_Member_1_REFINERY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_REFINERY
	case "Enum_CodeVerticalStructureBaseType_Member_1_RIG":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_RIG
	case "Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_SIGN":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SIGN
	case "Enum_CodeVerticalStructureBaseType_Member_1_SPIRE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_SPIRE
	case "Enum_CodeVerticalStructureBaseType_Member_1_STACK":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STACK
	case "Enum_CodeVerticalStructureBaseType_Member_1_STADIUM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STADIUM
	case "Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_TANK":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TANK
	case "Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON
	case "Enum_CodeVerticalStructureBaseType_Member_1_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TOWER
	case "Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY
	case "Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE
	case "Enum_CodeVerticalStructureBaseType_Member_1_TREE":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_TREE
	case "Enum_CodeVerticalStructureBaseType_Member_1_URBAN":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_URBAN
	case "Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION
	case "Enum_CodeVerticalStructureBaseType_Member_1_WALL":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WALL
	case "Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM
	case "Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER
	case "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL
	case "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS":
		*enum_codeverticalstructurebasetype_member_1 = Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeverticalstructurebasetype_member_1 *Enum_CodeVerticalStructureBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeverticalstructurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP"
	case Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA"
	case Enum_CodeVerticalStructureBaseType_Member_1_ARCH:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_ARCH"
	case Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE"
	case Enum_CodeVerticalStructureBaseType_Member_1_BUILDING:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_BUILDING"
	case Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR"
	case Enum_CodeVerticalStructureBaseType_Member_1_CATENARY:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_CATENARY"
	case Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_CRANE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_CRANE"
	case Enum_CodeVerticalStructureBaseType_Member_1_DAM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_DAM"
	case Enum_CodeVerticalStructureBaseType_Member_1_DOME:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_DOME"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_FENCE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_FENCE"
	case Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_GATE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_GATE"
	case Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY"
	case Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE"
	case Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT"
	case Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT"
	case Enum_CodeVerticalStructureBaseType_Member_1_NAVAID:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_NAVAID"
	case Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR"
	case Enum_CodeVerticalStructureBaseType_Member_1_POLE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_POLE"
	case Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT"
	case Enum_CodeVerticalStructureBaseType_Member_1_REFINERY:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_REFINERY"
	case Enum_CodeVerticalStructureBaseType_Member_1_RIG:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_RIG"
	case Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_SIGN:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_SIGN"
	case Enum_CodeVerticalStructureBaseType_Member_1_SPIRE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_SPIRE"
	case Enum_CodeVerticalStructureBaseType_Member_1_STACK:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_STACK"
	case Enum_CodeVerticalStructureBaseType_Member_1_STADIUM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_STADIUM"
	case Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_TANK:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TANK"
	case Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON"
	case Enum_CodeVerticalStructureBaseType_Member_1_TOWER:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY"
	case Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE"
	case Enum_CodeVerticalStructureBaseType_Member_1_TREE:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_TREE"
	case Enum_CodeVerticalStructureBaseType_Member_1_URBAN:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_URBAN"
	case Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION"
	case Enum_CodeVerticalStructureBaseType_Member_1_WALL:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WALL"
	case Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM"
	case Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER"
	case Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL"
	case Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS:
		res = "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS"
	}
	return
}

func (enum_codeverticalstructurebasetype_member_1 Enum_CodeVerticalStructureBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_AG_EQUIP")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_ANTENNA")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_ARCH")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_BRIDGE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_BUILDING")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_CABLE_CAR")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_CATENARY")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_COMPRESSED_AIR_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_MONITORING_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_CONTROL_TOWER")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_COOLING_TOWER")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_CRANE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_DAM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_DOME")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_EXIT_LIGHT")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_ELECTRICAL_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_ELEVATOR")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_FENCE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_FUEL_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_GATE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_GENERAL_UTILITY")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_GRAIN_ELEVATOR")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_HEAT_COOL_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_INDUSTRIAL_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_LIGHTHOUSE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_MONUMENT")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_GAS_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_NATURAL_HIGHPOINT")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_NAVAID")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_NUCLEAR_REACTOR")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_POLE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_POWER_PLANT")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_REFINERY")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_RIG")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_SALTWATER_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_SIGN")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_SPIRE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_STACK")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_STADIUM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_STORM_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TANK")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TETHERED_BALLOON")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TOWER")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TRAMWAY")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TRANSMISSION_LINE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_TREE")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_URBAN")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_VEGETATION")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WALL")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WASTEWATER_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WATER_SYSTEM")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WATER_TOWER")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL")
	res = append(res, "Enum_CodeVerticalStructureBaseType_Member_1_WINDMILL_FARMS")

	return
}

func (enum_codeverticalstructurebasetype_member_1 Enum_CodeVerticalStructureBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AG_EQUIP")
	res = append(res, "ANTENNA")
	res = append(res, "ARCH")
	res = append(res, "BRIDGE")
	res = append(res, "BUILDING")
	res = append(res, "CABLE_CAR")
	res = append(res, "CATENARY")
	res = append(res, "COMPRESSED_AIR_SYSTEM")
	res = append(res, "CONTROL_MONITORING_SYSTEM")
	res = append(res, "CONTROL_TOWER")
	res = append(res, "COOLING_TOWER")
	res = append(res, "CRANE")
	res = append(res, "DAM")
	res = append(res, "DOME")
	res = append(res, "ELECTRICAL_EXIT_LIGHT")
	res = append(res, "ELECTRICAL_SYSTEM")
	res = append(res, "ELEVATOR")
	res = append(res, "FENCE")
	res = append(res, "FUEL_SYSTEM")
	res = append(res, "GATE")
	res = append(res, "GENERAL_UTILITY")
	res = append(res, "GRAIN_ELEVATOR")
	res = append(res, "HEAT_COOL_SYSTEM")
	res = append(res, "INDUSTRIAL_SYSTEM")
	res = append(res, "LIGHTHOUSE")
	res = append(res, "MONUMENT")
	res = append(res, "NATURAL_GAS_SYSTEM")
	res = append(res, "NATURAL_HIGHPOINT")
	res = append(res, "NAVAID")
	res = append(res, "NUCLEAR_REACTOR")
	res = append(res, "POLE")
	res = append(res, "POWER_PLANT")
	res = append(res, "REFINERY")
	res = append(res, "RIG")
	res = append(res, "SALTWATER_SYSTEM")
	res = append(res, "SIGN")
	res = append(res, "SPIRE")
	res = append(res, "STACK")
	res = append(res, "STADIUM")
	res = append(res, "STORM_SYSTEM")
	res = append(res, "TANK")
	res = append(res, "TETHERED_BALLOON")
	res = append(res, "TOWER")
	res = append(res, "TRAMWAY")
	res = append(res, "TRANSMISSION_LINE")
	res = append(res, "TREE")
	res = append(res, "URBAN")
	res = append(res, "VEGETATION")
	res = append(res, "WALL")
	res = append(res, "WASTEWATER_SYSTEM")
	res = append(res, "WATER_SYSTEM")
	res = append(res, "WATER_TOWER")
	res = append(res, "WINDMILL")
	res = append(res, "WINDMILL_FARMS")

	return
}

// Utility function for Enum_CodeVerticalStructureMarkingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeverticalstructuremarkingbasetype_member_1 Enum_CodeVerticalStructureMarkingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeverticalstructuremarkingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR:
		res = "MONOCOLOUR"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED:
		res = "CHEQUERED"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS:
		res = "HBANDS"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS:
		res = "VBANDS"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG:
		res = "FLAG"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS:
		res = "MARKERS"
	}
	return
}

func (enum_codeverticalstructuremarkingbasetype_member_1 *Enum_CodeVerticalStructureMarkingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MONOCOLOUR":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR
		return
	case "CHEQUERED":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED
		return
	case "HBANDS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS
		return
	case "VBANDS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS
		return
	case "FLAG":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG
		return
	case "MARKERS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeverticalstructuremarkingbasetype_member_1 *Enum_CodeVerticalStructureMarkingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG
	case "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS":
		*enum_codeverticalstructuremarkingbasetype_member_1 = Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeverticalstructuremarkingbasetype_member_1 *Enum_CodeVerticalStructureMarkingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeverticalstructuremarkingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG"
	case Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS:
		res = "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS"
	}
	return
}

func (enum_codeverticalstructuremarkingbasetype_member_1 Enum_CodeVerticalStructureMarkingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MONOCOLOUR")
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_CHEQUERED")
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_HBANDS")
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_VBANDS")
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_FLAG")
	res = append(res, "Enum_CodeVerticalStructureMarkingBaseType_Member_1_MARKERS")

	return
}

func (enum_codeverticalstructuremarkingbasetype_member_1 Enum_CodeVerticalStructureMarkingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MONOCOLOUR")
	res = append(res, "CHEQUERED")
	res = append(res, "HBANDS")
	res = append(res, "VBANDS")
	res = append(res, "FLAG")
	res = append(res, "MARKERS")

	return
}

// Utility function for Enum_CodeVerticalStructureMaterialBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeverticalstructurematerialbasetype_member_1 Enum_CodeVerticalStructureMaterialBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeverticalstructurematerialbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK:
		res = "ADOBE_BRICK"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM:
		res = "ALUMINIUM"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK:
		res = "BRICK"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE:
		res = "CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS:
		res = "FIBREGLASS"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS:
		res = "GLASS"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON:
		res = "IRON"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY:
		res = "MASONRY"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL:
		res = "METAL"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD:
		res = "MUD"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT:
		res = "PLANT"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE:
		res = "PRESTRESSED_CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE:
		res = "REINFORCED_CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD:
		res = "SOD"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL:
		res = "STEEL"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE:
		res = "STONE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER:
		res = "TREATED_TIMBER"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD:
		res = "WOOD"
	}
	return
}

func (enum_codeverticalstructurematerialbasetype_member_1 *Enum_CodeVerticalStructureMaterialBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ADOBE_BRICK":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK
		return
	case "ALUMINIUM":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM
		return
	case "BRICK":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK
		return
	case "CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE
		return
	case "FIBREGLASS":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS
		return
	case "GLASS":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS
		return
	case "IRON":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON
		return
	case "MASONRY":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY
		return
	case "METAL":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL
		return
	case "MUD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD
		return
	case "PLANT":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT
		return
	case "PRESTRESSED_CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE
		return
	case "REINFORCED_CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE
		return
	case "SOD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD
		return
	case "STEEL":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL
		return
	case "STONE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE
		return
	case "TREATED_TIMBER":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER
		return
	case "WOOD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeverticalstructurematerialbasetype_member_1 *Enum_CodeVerticalStructureMaterialBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER
	case "Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD":
		*enum_codeverticalstructurematerialbasetype_member_1 = Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeverticalstructurematerialbasetype_member_1 *Enum_CodeVerticalStructureMaterialBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeverticalstructurematerialbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER"
	case Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD:
		res = "Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD"
	}
	return
}

func (enum_codeverticalstructurematerialbasetype_member_1 Enum_CodeVerticalStructureMaterialBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ADOBE_BRICK")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_ALUMINIUM")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_BRICK")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_CONCRETE")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_FIBREGLASS")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_GLASS")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_IRON")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MASONRY")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_METAL")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_MUD")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PLANT")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_PRESTRESSED_CONCRETE")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_REINFORCED_CONCRETE")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_SOD")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STEEL")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_STONE")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_TREATED_TIMBER")
	res = append(res, "Enum_CodeVerticalStructureMaterialBaseType_Member_1_WOOD")

	return
}

func (enum_codeverticalstructurematerialbasetype_member_1 Enum_CodeVerticalStructureMaterialBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ADOBE_BRICK")
	res = append(res, "ALUMINIUM")
	res = append(res, "BRICK")
	res = append(res, "CONCRETE")
	res = append(res, "FIBREGLASS")
	res = append(res, "GLASS")
	res = append(res, "IRON")
	res = append(res, "MASONRY")
	res = append(res, "METAL")
	res = append(res, "MUD")
	res = append(res, "PLANT")
	res = append(res, "PRESTRESSED_CONCRETE")
	res = append(res, "REINFORCED_CONCRETE")
	res = append(res, "SOD")
	res = append(res, "STEEL")
	res = append(res, "STONE")
	res = append(res, "TREATED_TIMBER")
	res = append(res, "WOOD")

	return
}

// Utility function for Enum_CodeYesNoBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeyesnobasetype_member_1 Enum_CodeYesNoBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeyesnobasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeYesNoBaseType_Member_1_YES:
		res = "YES"
	case Enum_CodeYesNoBaseType_Member_1_NO:
		res = "NO"
	}
	return
}

func (enum_codeyesnobasetype_member_1 *Enum_CodeYesNoBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "YES":
		*enum_codeyesnobasetype_member_1 = Enum_CodeYesNoBaseType_Member_1_YES
		return
	case "NO":
		*enum_codeyesnobasetype_member_1 = Enum_CodeYesNoBaseType_Member_1_NO
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeyesnobasetype_member_1 *Enum_CodeYesNoBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeYesNoBaseType_Member_1_YES":
		*enum_codeyesnobasetype_member_1 = Enum_CodeYesNoBaseType_Member_1_YES
	case "Enum_CodeYesNoBaseType_Member_1_NO":
		*enum_codeyesnobasetype_member_1 = Enum_CodeYesNoBaseType_Member_1_NO
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeyesnobasetype_member_1 *Enum_CodeYesNoBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeyesnobasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeYesNoBaseType_Member_1_YES:
		res = "Enum_CodeYesNoBaseType_Member_1_YES"
	case Enum_CodeYesNoBaseType_Member_1_NO:
		res = "Enum_CodeYesNoBaseType_Member_1_NO"
	}
	return
}

func (enum_codeyesnobasetype_member_1 Enum_CodeYesNoBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeYesNoBaseType_Member_1_YES")
	res = append(res, "Enum_CodeYesNoBaseType_Member_1_NO")

	return
}

func (enum_codeyesnobasetype_member_1 Enum_CodeYesNoBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "YES")
	res = append(res, "NO")

	return
}

// Utility function for Enum_UomDistanceType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomdistancetype_member_1 Enum_UomDistanceType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomdistancetype_member_1 {
	// insertion code per enum code
	case Enum_UomDistanceType_Member_1_NM:
		res = "NM"
	case Enum_UomDistanceType_Member_1_KM:
		res = "KM"
	case Enum_UomDistanceType_Member_1_M:
		res = "M"
	case Enum_UomDistanceType_Member_1_FT:
		res = "FT"
	case Enum_UomDistanceType_Member_1_MI:
		res = "MI"
	case Enum_UomDistanceType_Member_1_CM:
		res = "CM"
	}
	return
}

func (enum_uomdistancetype_member_1 *Enum_UomDistanceType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_NM
		return
	case "KM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_KM
		return
	case "M":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_M
		return
	case "FT":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_FT
		return
	case "MI":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_MI
		return
	case "CM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_CM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomdistancetype_member_1 *Enum_UomDistanceType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomDistanceType_Member_1_NM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_NM
	case "Enum_UomDistanceType_Member_1_KM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_KM
	case "Enum_UomDistanceType_Member_1_M":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_M
	case "Enum_UomDistanceType_Member_1_FT":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_FT
	case "Enum_UomDistanceType_Member_1_MI":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_MI
	case "Enum_UomDistanceType_Member_1_CM":
		*enum_uomdistancetype_member_1 = Enum_UomDistanceType_Member_1_CM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomdistancetype_member_1 *Enum_UomDistanceType_Member_1) ToCodeString() (res string) {

	switch *enum_uomdistancetype_member_1 {
	// insertion code per enum code
	case Enum_UomDistanceType_Member_1_NM:
		res = "Enum_UomDistanceType_Member_1_NM"
	case Enum_UomDistanceType_Member_1_KM:
		res = "Enum_UomDistanceType_Member_1_KM"
	case Enum_UomDistanceType_Member_1_M:
		res = "Enum_UomDistanceType_Member_1_M"
	case Enum_UomDistanceType_Member_1_FT:
		res = "Enum_UomDistanceType_Member_1_FT"
	case Enum_UomDistanceType_Member_1_MI:
		res = "Enum_UomDistanceType_Member_1_MI"
	case Enum_UomDistanceType_Member_1_CM:
		res = "Enum_UomDistanceType_Member_1_CM"
	}
	return
}

func (enum_uomdistancetype_member_1 Enum_UomDistanceType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomDistanceType_Member_1_NM")
	res = append(res, "Enum_UomDistanceType_Member_1_KM")
	res = append(res, "Enum_UomDistanceType_Member_1_M")
	res = append(res, "Enum_UomDistanceType_Member_1_FT")
	res = append(res, "Enum_UomDistanceType_Member_1_MI")
	res = append(res, "Enum_UomDistanceType_Member_1_CM")

	return
}

func (enum_uomdistancetype_member_1 Enum_UomDistanceType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NM")
	res = append(res, "KM")
	res = append(res, "M")
	res = append(res, "FT")
	res = append(res, "MI")
	res = append(res, "CM")

	return
}

// Utility function for Enum_UomDistanceVerticalType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomdistanceverticaltype_member_1 Enum_UomDistanceVerticalType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomdistanceverticaltype_member_1 {
	// insertion code per enum code
	case Enum_UomDistanceVerticalType_Member_1_FT:
		res = "FT"
	case Enum_UomDistanceVerticalType_Member_1_M:
		res = "M"
	case Enum_UomDistanceVerticalType_Member_1_FL:
		res = "FL"
	case Enum_UomDistanceVerticalType_Member_1_SM:
		res = "SM"
	}
	return
}

func (enum_uomdistanceverticaltype_member_1 *Enum_UomDistanceVerticalType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FT":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_FT
		return
	case "M":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_M
		return
	case "FL":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_FL
		return
	case "SM":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_SM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomdistanceverticaltype_member_1 *Enum_UomDistanceVerticalType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomDistanceVerticalType_Member_1_FT":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_FT
	case "Enum_UomDistanceVerticalType_Member_1_M":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_M
	case "Enum_UomDistanceVerticalType_Member_1_FL":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_FL
	case "Enum_UomDistanceVerticalType_Member_1_SM":
		*enum_uomdistanceverticaltype_member_1 = Enum_UomDistanceVerticalType_Member_1_SM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomdistanceverticaltype_member_1 *Enum_UomDistanceVerticalType_Member_1) ToCodeString() (res string) {

	switch *enum_uomdistanceverticaltype_member_1 {
	// insertion code per enum code
	case Enum_UomDistanceVerticalType_Member_1_FT:
		res = "Enum_UomDistanceVerticalType_Member_1_FT"
	case Enum_UomDistanceVerticalType_Member_1_M:
		res = "Enum_UomDistanceVerticalType_Member_1_M"
	case Enum_UomDistanceVerticalType_Member_1_FL:
		res = "Enum_UomDistanceVerticalType_Member_1_FL"
	case Enum_UomDistanceVerticalType_Member_1_SM:
		res = "Enum_UomDistanceVerticalType_Member_1_SM"
	}
	return
}

func (enum_uomdistanceverticaltype_member_1 Enum_UomDistanceVerticalType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomDistanceVerticalType_Member_1_FT")
	res = append(res, "Enum_UomDistanceVerticalType_Member_1_M")
	res = append(res, "Enum_UomDistanceVerticalType_Member_1_FL")
	res = append(res, "Enum_UomDistanceVerticalType_Member_1_SM")

	return
}

func (enum_uomdistanceverticaltype_member_1 Enum_UomDistanceVerticalType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FT")
	res = append(res, "M")
	res = append(res, "FL")
	res = append(res, "SM")

	return
}

// Utility function for Enum_UomDurationType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomdurationtype_member_1 Enum_UomDurationType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomdurationtype_member_1 {
	// insertion code per enum code
	case Enum_UomDurationType_Member_1_HR:
		res = "HR"
	case Enum_UomDurationType_Member_1_MIN_XSD:
		res = "MIN"
	case Enum_UomDurationType_Member_1_SEC:
		res = "SEC"
	}
	return
}

func (enum_uomdurationtype_member_1 *Enum_UomDurationType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HR":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_HR
		return
	case "MIN":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_MIN_XSD
		return
	case "SEC":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_SEC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomdurationtype_member_1 *Enum_UomDurationType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomDurationType_Member_1_HR":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_HR
	case "Enum_UomDurationType_Member_1_MIN_XSD":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_MIN_XSD
	case "Enum_UomDurationType_Member_1_SEC":
		*enum_uomdurationtype_member_1 = Enum_UomDurationType_Member_1_SEC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomdurationtype_member_1 *Enum_UomDurationType_Member_1) ToCodeString() (res string) {

	switch *enum_uomdurationtype_member_1 {
	// insertion code per enum code
	case Enum_UomDurationType_Member_1_HR:
		res = "Enum_UomDurationType_Member_1_HR"
	case Enum_UomDurationType_Member_1_MIN_XSD:
		res = "Enum_UomDurationType_Member_1_MIN_XSD"
	case Enum_UomDurationType_Member_1_SEC:
		res = "Enum_UomDurationType_Member_1_SEC"
	}
	return
}

func (enum_uomdurationtype_member_1 Enum_UomDurationType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomDurationType_Member_1_HR")
	res = append(res, "Enum_UomDurationType_Member_1_MIN_XSD")
	res = append(res, "Enum_UomDurationType_Member_1_SEC")

	return
}

func (enum_uomdurationtype_member_1 Enum_UomDurationType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HR")
	res = append(res, "MIN")
	res = append(res, "SEC")

	return
}

// Utility function for Enum_UomLightIntensityType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomlightintensitytype_member_1 Enum_UomLightIntensityType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomlightintensitytype_member_1 {
	// insertion code per enum code
	case Enum_UomLightIntensityType_Member_1_CD:
		res = "CD"
	}
	return
}

func (enum_uomlightintensitytype_member_1 *Enum_UomLightIntensityType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CD":
		*enum_uomlightintensitytype_member_1 = Enum_UomLightIntensityType_Member_1_CD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomlightintensitytype_member_1 *Enum_UomLightIntensityType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomLightIntensityType_Member_1_CD":
		*enum_uomlightintensitytype_member_1 = Enum_UomLightIntensityType_Member_1_CD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomlightintensitytype_member_1 *Enum_UomLightIntensityType_Member_1) ToCodeString() (res string) {

	switch *enum_uomlightintensitytype_member_1 {
	// insertion code per enum code
	case Enum_UomLightIntensityType_Member_1_CD:
		res = "Enum_UomLightIntensityType_Member_1_CD"
	}
	return
}

func (enum_uomlightintensitytype_member_1 Enum_UomLightIntensityType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomLightIntensityType_Member_1_CD")

	return
}

func (enum_uomlightintensitytype_member_1 Enum_UomLightIntensityType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CD")

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
