// generated code - do not edit
package attributes

// insertion point of enum utility functions
// Utility function for EnumCancelLocation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumcancellocation EnumCancelLocation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumcancellocation {
	// insertion code per enum code
	case EnumCancelLocationLeft:
		res = "left"
	case EnumCancelLocationRight:
		res = "right"
	case EnumCancelLocationBeforeBarline:
		res = "before-barline"
	}
	return
}

func (enumcancellocation *EnumCancelLocation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enumcancellocation = EnumCancelLocationLeft
		return
	case "right":
		*enumcancellocation = EnumCancelLocationRight
		return
	case "before-barline":
		*enumcancellocation = EnumCancelLocationBeforeBarline
		return
	default:
		return errUnkownEnum
	}
}

func (enumcancellocation *EnumCancelLocation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumCancelLocationLeft":
		*enumcancellocation = EnumCancelLocationLeft
	case "EnumCancelLocationRight":
		*enumcancellocation = EnumCancelLocationRight
	case "EnumCancelLocationBeforeBarline":
		*enumcancellocation = EnumCancelLocationBeforeBarline
	default:
		err = errUnkownEnum
	}
	return
}

func (enumcancellocation *EnumCancelLocation) ToCodeString() (res string) {

	switch *enumcancellocation {
	// insertion code per enum code
	case EnumCancelLocationLeft:
		res = "EnumCancelLocationLeft"
	case EnumCancelLocationRight:
		res = "EnumCancelLocationRight"
	case EnumCancelLocationBeforeBarline:
		res = "EnumCancelLocationBeforeBarline"
	}
	return
}

func (enumcancellocation EnumCancelLocation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumCancelLocationLeft")
	res = append(res, "EnumCancelLocationRight")
	res = append(res, "EnumCancelLocationBeforeBarline")

	return
}

func (enumcancellocation EnumCancelLocation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "before-barline")

	return
}

// Utility function for EnumClefSign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumclefsign EnumClefSign) ToString() (res string) {

	// migration of former implementation of enum
	switch enumclefsign {
	// insertion code per enum code
	case EnumClefSignG:
		res = "G"
	case EnumClefSignF:
		res = "F"
	case EnumClefSignC:
		res = "C"
	case EnumClefSignPercussion:
		res = "percussion"
	case EnumClefSignTAB:
		res = "TAB"
	case EnumClefSignJianpu:
		res = "jianpu"
	case EnumClefSignNone:
		res = "none"
	}
	return
}

func (enumclefsign *EnumClefSign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "G":
		*enumclefsign = EnumClefSignG
		return
	case "F":
		*enumclefsign = EnumClefSignF
		return
	case "C":
		*enumclefsign = EnumClefSignC
		return
	case "percussion":
		*enumclefsign = EnumClefSignPercussion
		return
	case "TAB":
		*enumclefsign = EnumClefSignTAB
		return
	case "jianpu":
		*enumclefsign = EnumClefSignJianpu
		return
	case "none":
		*enumclefsign = EnumClefSignNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumclefsign *EnumClefSign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumClefSignG":
		*enumclefsign = EnumClefSignG
	case "EnumClefSignF":
		*enumclefsign = EnumClefSignF
	case "EnumClefSignC":
		*enumclefsign = EnumClefSignC
	case "EnumClefSignPercussion":
		*enumclefsign = EnumClefSignPercussion
	case "EnumClefSignTAB":
		*enumclefsign = EnumClefSignTAB
	case "EnumClefSignJianpu":
		*enumclefsign = EnumClefSignJianpu
	case "EnumClefSignNone":
		*enumclefsign = EnumClefSignNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumclefsign *EnumClefSign) ToCodeString() (res string) {

	switch *enumclefsign {
	// insertion code per enum code
	case EnumClefSignG:
		res = "EnumClefSignG"
	case EnumClefSignF:
		res = "EnumClefSignF"
	case EnumClefSignC:
		res = "EnumClefSignC"
	case EnumClefSignPercussion:
		res = "EnumClefSignPercussion"
	case EnumClefSignTAB:
		res = "EnumClefSignTAB"
	case EnumClefSignJianpu:
		res = "EnumClefSignJianpu"
	case EnumClefSignNone:
		res = "EnumClefSignNone"
	}
	return
}

func (enumclefsign EnumClefSign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumClefSignG")
	res = append(res, "EnumClefSignF")
	res = append(res, "EnumClefSignC")
	res = append(res, "EnumClefSignPercussion")
	res = append(res, "EnumClefSignTAB")
	res = append(res, "EnumClefSignJianpu")
	res = append(res, "EnumClefSignNone")

	return
}

func (enumclefsign EnumClefSign) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "G")
	res = append(res, "F")
	res = append(res, "C")
	res = append(res, "percussion")
	res = append(res, "TAB")
	res = append(res, "jianpu")
	res = append(res, "none")

	return
}

// Utility function for EnumGroupSymbolValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumgroupsymbolvalue EnumGroupSymbolValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumgroupsymbolvalue {
	// insertion code per enum code
	case EnumGroupSymbolValueNone:
		res = "none"
	case EnumGroupSymbolValueBrace:
		res = "brace"
	case EnumGroupSymbolValueLine:
		res = "line"
	case EnumGroupSymbolValueBracket:
		res = "bracket"
	case EnumGroupSymbolValueSquare:
		res = "square"
	}
	return
}

func (enumgroupsymbolvalue *EnumGroupSymbolValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enumgroupsymbolvalue = EnumGroupSymbolValueNone
		return
	case "brace":
		*enumgroupsymbolvalue = EnumGroupSymbolValueBrace
		return
	case "line":
		*enumgroupsymbolvalue = EnumGroupSymbolValueLine
		return
	case "bracket":
		*enumgroupsymbolvalue = EnumGroupSymbolValueBracket
		return
	case "square":
		*enumgroupsymbolvalue = EnumGroupSymbolValueSquare
		return
	default:
		return errUnkownEnum
	}
}

func (enumgroupsymbolvalue *EnumGroupSymbolValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumGroupSymbolValueNone":
		*enumgroupsymbolvalue = EnumGroupSymbolValueNone
	case "EnumGroupSymbolValueBrace":
		*enumgroupsymbolvalue = EnumGroupSymbolValueBrace
	case "EnumGroupSymbolValueLine":
		*enumgroupsymbolvalue = EnumGroupSymbolValueLine
	case "EnumGroupSymbolValueBracket":
		*enumgroupsymbolvalue = EnumGroupSymbolValueBracket
	case "EnumGroupSymbolValueSquare":
		*enumgroupsymbolvalue = EnumGroupSymbolValueSquare
	default:
		err = errUnkownEnum
	}
	return
}

func (enumgroupsymbolvalue *EnumGroupSymbolValue) ToCodeString() (res string) {

	switch *enumgroupsymbolvalue {
	// insertion code per enum code
	case EnumGroupSymbolValueNone:
		res = "EnumGroupSymbolValueNone"
	case EnumGroupSymbolValueBrace:
		res = "EnumGroupSymbolValueBrace"
	case EnumGroupSymbolValueLine:
		res = "EnumGroupSymbolValueLine"
	case EnumGroupSymbolValueBracket:
		res = "EnumGroupSymbolValueBracket"
	case EnumGroupSymbolValueSquare:
		res = "EnumGroupSymbolValueSquare"
	}
	return
}

func (enumgroupsymbolvalue EnumGroupSymbolValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumGroupSymbolValueNone")
	res = append(res, "EnumGroupSymbolValueBrace")
	res = append(res, "EnumGroupSymbolValueLine")
	res = append(res, "EnumGroupSymbolValueBracket")
	res = append(res, "EnumGroupSymbolValueSquare")

	return
}

func (enumgroupsymbolvalue EnumGroupSymbolValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "brace")
	res = append(res, "line")
	res = append(res, "bracket")
	res = append(res, "square")

	return
}

// Utility function for EnumShowFrets
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumshowfrets EnumShowFrets) ToString() (res string) {

	// migration of former implementation of enum
	switch enumshowfrets {
	// insertion code per enum code
	case EnumShowFretsNumbers:
		res = "numbers"
	case EnumShowFretsLetters:
		res = "letters"
	}
	return
}

func (enumshowfrets *EnumShowFrets) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "numbers":
		*enumshowfrets = EnumShowFretsNumbers
		return
	case "letters":
		*enumshowfrets = EnumShowFretsLetters
		return
	default:
		return errUnkownEnum
	}
}

func (enumshowfrets *EnumShowFrets) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumShowFretsNumbers":
		*enumshowfrets = EnumShowFretsNumbers
	case "EnumShowFretsLetters":
		*enumshowfrets = EnumShowFretsLetters
	default:
		err = errUnkownEnum
	}
	return
}

func (enumshowfrets *EnumShowFrets) ToCodeString() (res string) {

	switch *enumshowfrets {
	// insertion code per enum code
	case EnumShowFretsNumbers:
		res = "EnumShowFretsNumbers"
	case EnumShowFretsLetters:
		res = "EnumShowFretsLetters"
	}
	return
}

func (enumshowfrets EnumShowFrets) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumShowFretsNumbers")
	res = append(res, "EnumShowFretsLetters")

	return
}

func (enumshowfrets EnumShowFrets) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "numbers")
	res = append(res, "letters")

	return
}

// Utility function for EnumStaffType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstafftype EnumStaffType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstafftype {
	// insertion code per enum code
	case EnumStaffTypeOssia:
		res = "ossia"
	case EnumStaffTypeEditorial:
		res = "editorial"
	case EnumStaffTypeCue:
		res = "cue"
	case EnumStaffTypeAlternate:
		res = "alternate"
	case EnumStaffTypeRegular:
		res = "regular"
	}
	return
}

func (enumstafftype *EnumStaffType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ossia":
		*enumstafftype = EnumStaffTypeOssia
		return
	case "editorial":
		*enumstafftype = EnumStaffTypeEditorial
		return
	case "cue":
		*enumstafftype = EnumStaffTypeCue
		return
	case "alternate":
		*enumstafftype = EnumStaffTypeAlternate
		return
	case "regular":
		*enumstafftype = EnumStaffTypeRegular
		return
	default:
		return errUnkownEnum
	}
}

func (enumstafftype *EnumStaffType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStaffTypeOssia":
		*enumstafftype = EnumStaffTypeOssia
	case "EnumStaffTypeEditorial":
		*enumstafftype = EnumStaffTypeEditorial
	case "EnumStaffTypeCue":
		*enumstafftype = EnumStaffTypeCue
	case "EnumStaffTypeAlternate":
		*enumstafftype = EnumStaffTypeAlternate
	case "EnumStaffTypeRegular":
		*enumstafftype = EnumStaffTypeRegular
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstafftype *EnumStaffType) ToCodeString() (res string) {

	switch *enumstafftype {
	// insertion code per enum code
	case EnumStaffTypeOssia:
		res = "EnumStaffTypeOssia"
	case EnumStaffTypeEditorial:
		res = "EnumStaffTypeEditorial"
	case EnumStaffTypeCue:
		res = "EnumStaffTypeCue"
	case EnumStaffTypeAlternate:
		res = "EnumStaffTypeAlternate"
	case EnumStaffTypeRegular:
		res = "EnumStaffTypeRegular"
	}
	return
}

func (enumstafftype EnumStaffType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStaffTypeOssia")
	res = append(res, "EnumStaffTypeEditorial")
	res = append(res, "EnumStaffTypeCue")
	res = append(res, "EnumStaffTypeAlternate")
	res = append(res, "EnumStaffTypeRegular")

	return
}

func (enumstafftype EnumStaffType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ossia")
	res = append(res, "editorial")
	res = append(res, "cue")
	res = append(res, "alternate")
	res = append(res, "regular")

	return
}

// Utility function for EnumStep
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstep EnumStep) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstep {
	// insertion code per enum code
	case EnumStepA:
		res = "A"
	case EnumStepB:
		res = "B"
	case EnumStepC:
		res = "C"
	case EnumStepD:
		res = "D"
	case EnumStepE:
		res = "E"
	case EnumStepF:
		res = "F"
	case EnumStepG:
		res = "G"
	}
	return
}

func (enumstep *EnumStep) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enumstep = EnumStepA
		return
	case "B":
		*enumstep = EnumStepB
		return
	case "C":
		*enumstep = EnumStepC
		return
	case "D":
		*enumstep = EnumStepD
		return
	case "E":
		*enumstep = EnumStepE
		return
	case "F":
		*enumstep = EnumStepF
		return
	case "G":
		*enumstep = EnumStepG
		return
	default:
		return errUnkownEnum
	}
}

func (enumstep *EnumStep) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStepA":
		*enumstep = EnumStepA
	case "EnumStepB":
		*enumstep = EnumStepB
	case "EnumStepC":
		*enumstep = EnumStepC
	case "EnumStepD":
		*enumstep = EnumStepD
	case "EnumStepE":
		*enumstep = EnumStepE
	case "EnumStepF":
		*enumstep = EnumStepF
	case "EnumStepG":
		*enumstep = EnumStepG
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstep *EnumStep) ToCodeString() (res string) {

	switch *enumstep {
	// insertion code per enum code
	case EnumStepA:
		res = "EnumStepA"
	case EnumStepB:
		res = "EnumStepB"
	case EnumStepC:
		res = "EnumStepC"
	case EnumStepD:
		res = "EnumStepD"
	case EnumStepE:
		res = "EnumStepE"
	case EnumStepF:
		res = "EnumStepF"
	case EnumStepG:
		res = "EnumStepG"
	}
	return
}

func (enumstep EnumStep) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStepA")
	res = append(res, "EnumStepB")
	res = append(res, "EnumStepC")
	res = append(res, "EnumStepD")
	res = append(res, "EnumStepE")
	res = append(res, "EnumStepF")
	res = append(res, "EnumStepG")

	return
}

func (enumstep EnumStep) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")
	res = append(res, "E")
	res = append(res, "F")
	res = append(res, "G")

	return
}

// Utility function for EnumTimeRelation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtimerelation EnumTimeRelation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtimerelation {
	// insertion code per enum code
	case EnumTimeRelationParentheses:
		res = "parentheses"
	case EnumTimeRelationBracket:
		res = "bracket"
	case EnumTimeRelationEquals:
		res = "equals"
	case EnumTimeRelationSlash:
		res = "slash"
	case EnumTimeRelationSpace:
		res = "space"
	case EnumTimeRelationHyphen:
		res = "hyphen"
	}
	return
}

func (enumtimerelation *EnumTimeRelation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "parentheses":
		*enumtimerelation = EnumTimeRelationParentheses
		return
	case "bracket":
		*enumtimerelation = EnumTimeRelationBracket
		return
	case "equals":
		*enumtimerelation = EnumTimeRelationEquals
		return
	case "slash":
		*enumtimerelation = EnumTimeRelationSlash
		return
	case "space":
		*enumtimerelation = EnumTimeRelationSpace
		return
	case "hyphen":
		*enumtimerelation = EnumTimeRelationHyphen
		return
	default:
		return errUnkownEnum
	}
}

func (enumtimerelation *EnumTimeRelation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTimeRelationParentheses":
		*enumtimerelation = EnumTimeRelationParentheses
	case "EnumTimeRelationBracket":
		*enumtimerelation = EnumTimeRelationBracket
	case "EnumTimeRelationEquals":
		*enumtimerelation = EnumTimeRelationEquals
	case "EnumTimeRelationSlash":
		*enumtimerelation = EnumTimeRelationSlash
	case "EnumTimeRelationSpace":
		*enumtimerelation = EnumTimeRelationSpace
	case "EnumTimeRelationHyphen":
		*enumtimerelation = EnumTimeRelationHyphen
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtimerelation *EnumTimeRelation) ToCodeString() (res string) {

	switch *enumtimerelation {
	// insertion code per enum code
	case EnumTimeRelationParentheses:
		res = "EnumTimeRelationParentheses"
	case EnumTimeRelationBracket:
		res = "EnumTimeRelationBracket"
	case EnumTimeRelationEquals:
		res = "EnumTimeRelationEquals"
	case EnumTimeRelationSlash:
		res = "EnumTimeRelationSlash"
	case EnumTimeRelationSpace:
		res = "EnumTimeRelationSpace"
	case EnumTimeRelationHyphen:
		res = "EnumTimeRelationHyphen"
	}
	return
}

func (enumtimerelation EnumTimeRelation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTimeRelationParentheses")
	res = append(res, "EnumTimeRelationBracket")
	res = append(res, "EnumTimeRelationEquals")
	res = append(res, "EnumTimeRelationSlash")
	res = append(res, "EnumTimeRelationSpace")
	res = append(res, "EnumTimeRelationHyphen")

	return
}

func (enumtimerelation EnumTimeRelation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "parentheses")
	res = append(res, "bracket")
	res = append(res, "equals")
	res = append(res, "slash")
	res = append(res, "space")
	res = append(res, "hyphen")

	return
}

// Utility function for EnumTimeSeparator
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtimeseparator EnumTimeSeparator) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtimeseparator {
	// insertion code per enum code
	case EnumTimeSeparatorNone:
		res = "none"
	case EnumTimeSeparatorHorizontal:
		res = "horizontal"
	case EnumTimeSeparatorDiagonal:
		res = "diagonal"
	case EnumTimeSeparatorVertical:
		res = "vertical"
	case EnumTimeSeparatorAdjacent:
		res = "adjacent"
	}
	return
}

func (enumtimeseparator *EnumTimeSeparator) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enumtimeseparator = EnumTimeSeparatorNone
		return
	case "horizontal":
		*enumtimeseparator = EnumTimeSeparatorHorizontal
		return
	case "diagonal":
		*enumtimeseparator = EnumTimeSeparatorDiagonal
		return
	case "vertical":
		*enumtimeseparator = EnumTimeSeparatorVertical
		return
	case "adjacent":
		*enumtimeseparator = EnumTimeSeparatorAdjacent
		return
	default:
		return errUnkownEnum
	}
}

func (enumtimeseparator *EnumTimeSeparator) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTimeSeparatorNone":
		*enumtimeseparator = EnumTimeSeparatorNone
	case "EnumTimeSeparatorHorizontal":
		*enumtimeseparator = EnumTimeSeparatorHorizontal
	case "EnumTimeSeparatorDiagonal":
		*enumtimeseparator = EnumTimeSeparatorDiagonal
	case "EnumTimeSeparatorVertical":
		*enumtimeseparator = EnumTimeSeparatorVertical
	case "EnumTimeSeparatorAdjacent":
		*enumtimeseparator = EnumTimeSeparatorAdjacent
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtimeseparator *EnumTimeSeparator) ToCodeString() (res string) {

	switch *enumtimeseparator {
	// insertion code per enum code
	case EnumTimeSeparatorNone:
		res = "EnumTimeSeparatorNone"
	case EnumTimeSeparatorHorizontal:
		res = "EnumTimeSeparatorHorizontal"
	case EnumTimeSeparatorDiagonal:
		res = "EnumTimeSeparatorDiagonal"
	case EnumTimeSeparatorVertical:
		res = "EnumTimeSeparatorVertical"
	case EnumTimeSeparatorAdjacent:
		res = "EnumTimeSeparatorAdjacent"
	}
	return
}

func (enumtimeseparator EnumTimeSeparator) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTimeSeparatorNone")
	res = append(res, "EnumTimeSeparatorHorizontal")
	res = append(res, "EnumTimeSeparatorDiagonal")
	res = append(res, "EnumTimeSeparatorVertical")
	res = append(res, "EnumTimeSeparatorAdjacent")

	return
}

func (enumtimeseparator EnumTimeSeparator) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "horizontal")
	res = append(res, "diagonal")
	res = append(res, "vertical")
	res = append(res, "adjacent")

	return
}

// Utility function for EnumTimeSymbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtimesymbol EnumTimeSymbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtimesymbol {
	// insertion code per enum code
	case EnumTimeSymbolCommon:
		res = "common"
	case EnumTimeSymbolCut:
		res = "cut"
	case EnumTimeSymbolSingleNumber:
		res = "single-number"
	case EnumTimeSymbolNote:
		res = "note"
	case EnumTimeSymbolDottedNote:
		res = "dotted-note"
	case EnumTimeSymbolNormal:
		res = "normal"
	}
	return
}

func (enumtimesymbol *EnumTimeSymbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "common":
		*enumtimesymbol = EnumTimeSymbolCommon
		return
	case "cut":
		*enumtimesymbol = EnumTimeSymbolCut
		return
	case "single-number":
		*enumtimesymbol = EnumTimeSymbolSingleNumber
		return
	case "note":
		*enumtimesymbol = EnumTimeSymbolNote
		return
	case "dotted-note":
		*enumtimesymbol = EnumTimeSymbolDottedNote
		return
	case "normal":
		*enumtimesymbol = EnumTimeSymbolNormal
		return
	default:
		return errUnkownEnum
	}
}

func (enumtimesymbol *EnumTimeSymbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTimeSymbolCommon":
		*enumtimesymbol = EnumTimeSymbolCommon
	case "EnumTimeSymbolCut":
		*enumtimesymbol = EnumTimeSymbolCut
	case "EnumTimeSymbolSingleNumber":
		*enumtimesymbol = EnumTimeSymbolSingleNumber
	case "EnumTimeSymbolNote":
		*enumtimesymbol = EnumTimeSymbolNote
	case "EnumTimeSymbolDottedNote":
		*enumtimesymbol = EnumTimeSymbolDottedNote
	case "EnumTimeSymbolNormal":
		*enumtimesymbol = EnumTimeSymbolNormal
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtimesymbol *EnumTimeSymbol) ToCodeString() (res string) {

	switch *enumtimesymbol {
	// insertion code per enum code
	case EnumTimeSymbolCommon:
		res = "EnumTimeSymbolCommon"
	case EnumTimeSymbolCut:
		res = "EnumTimeSymbolCut"
	case EnumTimeSymbolSingleNumber:
		res = "EnumTimeSymbolSingleNumber"
	case EnumTimeSymbolNote:
		res = "EnumTimeSymbolNote"
	case EnumTimeSymbolDottedNote:
		res = "EnumTimeSymbolDottedNote"
	case EnumTimeSymbolNormal:
		res = "EnumTimeSymbolNormal"
	}
	return
}

func (enumtimesymbol EnumTimeSymbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTimeSymbolCommon")
	res = append(res, "EnumTimeSymbolCut")
	res = append(res, "EnumTimeSymbolSingleNumber")
	res = append(res, "EnumTimeSymbolNote")
	res = append(res, "EnumTimeSymbolDottedNote")
	res = append(res, "EnumTimeSymbolNormal")

	return
}

func (enumtimesymbol EnumTimeSymbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "common")
	res = append(res, "cut")
	res = append(res, "single-number")
	res = append(res, "note")
	res = append(res, "dotted-note")
	res = append(res, "normal")

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
