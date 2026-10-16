// generated code - do not edit
package direction

// insertion point of enum utility functions
// Utility function for EnumBeamValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbeamvalue EnumBeamValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbeamvalue {
	// insertion code per enum code
	case EnumBeamValueBegin:
		res = "begin"
	case EnumBeamValueContinue:
		res = "continue"
	case EnumBeamValueEnd:
		res = "end"
	case EnumBeamValueForwardHook:
		res = "forward hook"
	case EnumBeamValueBackwardHook:
		res = "backward hook"
	}
	return
}

func (enumbeamvalue *EnumBeamValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "begin":
		*enumbeamvalue = EnumBeamValueBegin
		return
	case "continue":
		*enumbeamvalue = EnumBeamValueContinue
		return
	case "end":
		*enumbeamvalue = EnumBeamValueEnd
		return
	case "forward hook":
		*enumbeamvalue = EnumBeamValueForwardHook
		return
	case "backward hook":
		*enumbeamvalue = EnumBeamValueBackwardHook
		return
	default:
		return errUnkownEnum
	}
}

func (enumbeamvalue *EnumBeamValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBeamValueBegin":
		*enumbeamvalue = EnumBeamValueBegin
	case "EnumBeamValueContinue":
		*enumbeamvalue = EnumBeamValueContinue
	case "EnumBeamValueEnd":
		*enumbeamvalue = EnumBeamValueEnd
	case "EnumBeamValueForwardHook":
		*enumbeamvalue = EnumBeamValueForwardHook
	case "EnumBeamValueBackwardHook":
		*enumbeamvalue = EnumBeamValueBackwardHook
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbeamvalue *EnumBeamValue) ToCodeString() (res string) {

	switch *enumbeamvalue {
	// insertion code per enum code
	case EnumBeamValueBegin:
		res = "EnumBeamValueBegin"
	case EnumBeamValueContinue:
		res = "EnumBeamValueContinue"
	case EnumBeamValueEnd:
		res = "EnumBeamValueEnd"
	case EnumBeamValueForwardHook:
		res = "EnumBeamValueForwardHook"
	case EnumBeamValueBackwardHook:
		res = "EnumBeamValueBackwardHook"
	}
	return
}

func (enumbeamvalue EnumBeamValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBeamValueBegin")
	res = append(res, "EnumBeamValueContinue")
	res = append(res, "EnumBeamValueEnd")
	res = append(res, "EnumBeamValueForwardHook")
	res = append(res, "EnumBeamValueBackwardHook")

	return
}

func (enumbeamvalue EnumBeamValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "begin")
	res = append(res, "continue")
	res = append(res, "end")
	res = append(res, "forward hook")
	res = append(res, "backward hook")

	return
}

// Utility function for EnumBeaterValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumbeatervalue EnumBeaterValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumbeatervalue {
	// insertion code per enum code
	case EnumBeaterValueBow:
		res = "bow"
	case EnumBeaterValueChimeHammer:
		res = "chime hammer"
	case EnumBeaterValueCoin:
		res = "coin"
	case EnumBeaterValueDrumStick:
		res = "drum stick"
	case EnumBeaterValueFinger:
		res = "finger"
	case EnumBeaterValueFingernail:
		res = "fingernail"
	case EnumBeaterValueFist:
		res = "fist"
	case EnumBeaterValueGuiroScraper:
		res = "guiro scraper"
	case EnumBeaterValueHammer:
		res = "hammer"
	case EnumBeaterValueHand:
		res = "hand"
	case EnumBeaterValueJazzStick:
		res = "jazz stick"
	case EnumBeaterValueKnittingNeedle:
		res = "knitting needle"
	case EnumBeaterValueMetalHammer:
		res = "metal hammer"
	case EnumBeaterValueSlideBrushOnGong:
		res = "slide brush on gong"
	case EnumBeaterValueSnareStick:
		res = "snare stick"
	case EnumBeaterValueSpoonMallet:
		res = "spoon mallet"
	case EnumBeaterValueSuperball:
		res = "superball"
	case EnumBeaterValueTriangleBeater:
		res = "triangle beater"
	case EnumBeaterValueTriangleBeaterPlain:
		res = "triangle beater plain"
	case EnumBeaterValueWireBrush:
		res = "wire brush"
	}
	return
}

func (enumbeatervalue *EnumBeaterValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bow":
		*enumbeatervalue = EnumBeaterValueBow
		return
	case "chime hammer":
		*enumbeatervalue = EnumBeaterValueChimeHammer
		return
	case "coin":
		*enumbeatervalue = EnumBeaterValueCoin
		return
	case "drum stick":
		*enumbeatervalue = EnumBeaterValueDrumStick
		return
	case "finger":
		*enumbeatervalue = EnumBeaterValueFinger
		return
	case "fingernail":
		*enumbeatervalue = EnumBeaterValueFingernail
		return
	case "fist":
		*enumbeatervalue = EnumBeaterValueFist
		return
	case "guiro scraper":
		*enumbeatervalue = EnumBeaterValueGuiroScraper
		return
	case "hammer":
		*enumbeatervalue = EnumBeaterValueHammer
		return
	case "hand":
		*enumbeatervalue = EnumBeaterValueHand
		return
	case "jazz stick":
		*enumbeatervalue = EnumBeaterValueJazzStick
		return
	case "knitting needle":
		*enumbeatervalue = EnumBeaterValueKnittingNeedle
		return
	case "metal hammer":
		*enumbeatervalue = EnumBeaterValueMetalHammer
		return
	case "slide brush on gong":
		*enumbeatervalue = EnumBeaterValueSlideBrushOnGong
		return
	case "snare stick":
		*enumbeatervalue = EnumBeaterValueSnareStick
		return
	case "spoon mallet":
		*enumbeatervalue = EnumBeaterValueSpoonMallet
		return
	case "superball":
		*enumbeatervalue = EnumBeaterValueSuperball
		return
	case "triangle beater":
		*enumbeatervalue = EnumBeaterValueTriangleBeater
		return
	case "triangle beater plain":
		*enumbeatervalue = EnumBeaterValueTriangleBeaterPlain
		return
	case "wire brush":
		*enumbeatervalue = EnumBeaterValueWireBrush
		return
	default:
		return errUnkownEnum
	}
}

func (enumbeatervalue *EnumBeaterValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumBeaterValueBow":
		*enumbeatervalue = EnumBeaterValueBow
	case "EnumBeaterValueChimeHammer":
		*enumbeatervalue = EnumBeaterValueChimeHammer
	case "EnumBeaterValueCoin":
		*enumbeatervalue = EnumBeaterValueCoin
	case "EnumBeaterValueDrumStick":
		*enumbeatervalue = EnumBeaterValueDrumStick
	case "EnumBeaterValueFinger":
		*enumbeatervalue = EnumBeaterValueFinger
	case "EnumBeaterValueFingernail":
		*enumbeatervalue = EnumBeaterValueFingernail
	case "EnumBeaterValueFist":
		*enumbeatervalue = EnumBeaterValueFist
	case "EnumBeaterValueGuiroScraper":
		*enumbeatervalue = EnumBeaterValueGuiroScraper
	case "EnumBeaterValueHammer":
		*enumbeatervalue = EnumBeaterValueHammer
	case "EnumBeaterValueHand":
		*enumbeatervalue = EnumBeaterValueHand
	case "EnumBeaterValueJazzStick":
		*enumbeatervalue = EnumBeaterValueJazzStick
	case "EnumBeaterValueKnittingNeedle":
		*enumbeatervalue = EnumBeaterValueKnittingNeedle
	case "EnumBeaterValueMetalHammer":
		*enumbeatervalue = EnumBeaterValueMetalHammer
	case "EnumBeaterValueSlideBrushOnGong":
		*enumbeatervalue = EnumBeaterValueSlideBrushOnGong
	case "EnumBeaterValueSnareStick":
		*enumbeatervalue = EnumBeaterValueSnareStick
	case "EnumBeaterValueSpoonMallet":
		*enumbeatervalue = EnumBeaterValueSpoonMallet
	case "EnumBeaterValueSuperball":
		*enumbeatervalue = EnumBeaterValueSuperball
	case "EnumBeaterValueTriangleBeater":
		*enumbeatervalue = EnumBeaterValueTriangleBeater
	case "EnumBeaterValueTriangleBeaterPlain":
		*enumbeatervalue = EnumBeaterValueTriangleBeaterPlain
	case "EnumBeaterValueWireBrush":
		*enumbeatervalue = EnumBeaterValueWireBrush
	default:
		err = errUnkownEnum
	}
	return
}

func (enumbeatervalue *EnumBeaterValue) ToCodeString() (res string) {

	switch *enumbeatervalue {
	// insertion code per enum code
	case EnumBeaterValueBow:
		res = "EnumBeaterValueBow"
	case EnumBeaterValueChimeHammer:
		res = "EnumBeaterValueChimeHammer"
	case EnumBeaterValueCoin:
		res = "EnumBeaterValueCoin"
	case EnumBeaterValueDrumStick:
		res = "EnumBeaterValueDrumStick"
	case EnumBeaterValueFinger:
		res = "EnumBeaterValueFinger"
	case EnumBeaterValueFingernail:
		res = "EnumBeaterValueFingernail"
	case EnumBeaterValueFist:
		res = "EnumBeaterValueFist"
	case EnumBeaterValueGuiroScraper:
		res = "EnumBeaterValueGuiroScraper"
	case EnumBeaterValueHammer:
		res = "EnumBeaterValueHammer"
	case EnumBeaterValueHand:
		res = "EnumBeaterValueHand"
	case EnumBeaterValueJazzStick:
		res = "EnumBeaterValueJazzStick"
	case EnumBeaterValueKnittingNeedle:
		res = "EnumBeaterValueKnittingNeedle"
	case EnumBeaterValueMetalHammer:
		res = "EnumBeaterValueMetalHammer"
	case EnumBeaterValueSlideBrushOnGong:
		res = "EnumBeaterValueSlideBrushOnGong"
	case EnumBeaterValueSnareStick:
		res = "EnumBeaterValueSnareStick"
	case EnumBeaterValueSpoonMallet:
		res = "EnumBeaterValueSpoonMallet"
	case EnumBeaterValueSuperball:
		res = "EnumBeaterValueSuperball"
	case EnumBeaterValueTriangleBeater:
		res = "EnumBeaterValueTriangleBeater"
	case EnumBeaterValueTriangleBeaterPlain:
		res = "EnumBeaterValueTriangleBeaterPlain"
	case EnumBeaterValueWireBrush:
		res = "EnumBeaterValueWireBrush"
	}
	return
}

func (enumbeatervalue EnumBeaterValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumBeaterValueBow")
	res = append(res, "EnumBeaterValueChimeHammer")
	res = append(res, "EnumBeaterValueCoin")
	res = append(res, "EnumBeaterValueDrumStick")
	res = append(res, "EnumBeaterValueFinger")
	res = append(res, "EnumBeaterValueFingernail")
	res = append(res, "EnumBeaterValueFist")
	res = append(res, "EnumBeaterValueGuiroScraper")
	res = append(res, "EnumBeaterValueHammer")
	res = append(res, "EnumBeaterValueHand")
	res = append(res, "EnumBeaterValueJazzStick")
	res = append(res, "EnumBeaterValueKnittingNeedle")
	res = append(res, "EnumBeaterValueMetalHammer")
	res = append(res, "EnumBeaterValueSlideBrushOnGong")
	res = append(res, "EnumBeaterValueSnareStick")
	res = append(res, "EnumBeaterValueSpoonMallet")
	res = append(res, "EnumBeaterValueSuperball")
	res = append(res, "EnumBeaterValueTriangleBeater")
	res = append(res, "EnumBeaterValueTriangleBeaterPlain")
	res = append(res, "EnumBeaterValueWireBrush")

	return
}

func (enumbeatervalue EnumBeaterValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bow")
	res = append(res, "chime hammer")
	res = append(res, "coin")
	res = append(res, "drum stick")
	res = append(res, "finger")
	res = append(res, "fingernail")
	res = append(res, "fist")
	res = append(res, "guiro scraper")
	res = append(res, "hammer")
	res = append(res, "hand")
	res = append(res, "jazz stick")
	res = append(res, "knitting needle")
	res = append(res, "metal hammer")
	res = append(res, "slide brush on gong")
	res = append(res, "snare stick")
	res = append(res, "spoon mallet")
	res = append(res, "superball")
	res = append(res, "triangle beater")
	res = append(res, "triangle beater plain")
	res = append(res, "wire brush")

	return
}

// Utility function for EnumDegreeSymbolValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumdegreesymbolvalue EnumDegreeSymbolValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumdegreesymbolvalue {
	// insertion code per enum code
	case EnumDegreeSymbolValueMajor:
		res = "major"
	case EnumDegreeSymbolValueMinor:
		res = "minor"
	case EnumDegreeSymbolValueAugmented:
		res = "augmented"
	case EnumDegreeSymbolValueDiminished:
		res = "diminished"
	case EnumDegreeSymbolValueHalfDiminished:
		res = "half-diminished"
	}
	return
}

func (enumdegreesymbolvalue *EnumDegreeSymbolValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueMajor
		return
	case "minor":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueMinor
		return
	case "augmented":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueAugmented
		return
	case "diminished":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueDiminished
		return
	case "half-diminished":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueHalfDiminished
		return
	default:
		return errUnkownEnum
	}
}

func (enumdegreesymbolvalue *EnumDegreeSymbolValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumDegreeSymbolValueMajor":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueMajor
	case "EnumDegreeSymbolValueMinor":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueMinor
	case "EnumDegreeSymbolValueAugmented":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueAugmented
	case "EnumDegreeSymbolValueDiminished":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueDiminished
	case "EnumDegreeSymbolValueHalfDiminished":
		*enumdegreesymbolvalue = EnumDegreeSymbolValueHalfDiminished
	default:
		err = errUnkownEnum
	}
	return
}

func (enumdegreesymbolvalue *EnumDegreeSymbolValue) ToCodeString() (res string) {

	switch *enumdegreesymbolvalue {
	// insertion code per enum code
	case EnumDegreeSymbolValueMajor:
		res = "EnumDegreeSymbolValueMajor"
	case EnumDegreeSymbolValueMinor:
		res = "EnumDegreeSymbolValueMinor"
	case EnumDegreeSymbolValueAugmented:
		res = "EnumDegreeSymbolValueAugmented"
	case EnumDegreeSymbolValueDiminished:
		res = "EnumDegreeSymbolValueDiminished"
	case EnumDegreeSymbolValueHalfDiminished:
		res = "EnumDegreeSymbolValueHalfDiminished"
	}
	return
}

func (enumdegreesymbolvalue EnumDegreeSymbolValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumDegreeSymbolValueMajor")
	res = append(res, "EnumDegreeSymbolValueMinor")
	res = append(res, "EnumDegreeSymbolValueAugmented")
	res = append(res, "EnumDegreeSymbolValueDiminished")
	res = append(res, "EnumDegreeSymbolValueHalfDiminished")

	return
}

func (enumdegreesymbolvalue EnumDegreeSymbolValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "augmented")
	res = append(res, "diminished")
	res = append(res, "half-diminished")

	return
}

// Utility function for EnumDegreeTypeValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumdegreetypevalue EnumDegreeTypeValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumdegreetypevalue {
	// insertion code per enum code
	case EnumDegreeTypeValueAdd:
		res = "add"
	case EnumDegreeTypeValueAlter:
		res = "alter"
	case EnumDegreeTypeValueSubtract:
		res = "subtract"
	}
	return
}

func (enumdegreetypevalue *EnumDegreeTypeValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "add":
		*enumdegreetypevalue = EnumDegreeTypeValueAdd
		return
	case "alter":
		*enumdegreetypevalue = EnumDegreeTypeValueAlter
		return
	case "subtract":
		*enumdegreetypevalue = EnumDegreeTypeValueSubtract
		return
	default:
		return errUnkownEnum
	}
}

func (enumdegreetypevalue *EnumDegreeTypeValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumDegreeTypeValueAdd":
		*enumdegreetypevalue = EnumDegreeTypeValueAdd
	case "EnumDegreeTypeValueAlter":
		*enumdegreetypevalue = EnumDegreeTypeValueAlter
	case "EnumDegreeTypeValueSubtract":
		*enumdegreetypevalue = EnumDegreeTypeValueSubtract
	default:
		err = errUnkownEnum
	}
	return
}

func (enumdegreetypevalue *EnumDegreeTypeValue) ToCodeString() (res string) {

	switch *enumdegreetypevalue {
	// insertion code per enum code
	case EnumDegreeTypeValueAdd:
		res = "EnumDegreeTypeValueAdd"
	case EnumDegreeTypeValueAlter:
		res = "EnumDegreeTypeValueAlter"
	case EnumDegreeTypeValueSubtract:
		res = "EnumDegreeTypeValueSubtract"
	}
	return
}

func (enumdegreetypevalue EnumDegreeTypeValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumDegreeTypeValueAdd")
	res = append(res, "EnumDegreeTypeValueAlter")
	res = append(res, "EnumDegreeTypeValueSubtract")

	return
}

func (enumdegreetypevalue EnumDegreeTypeValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "add")
	res = append(res, "alter")
	res = append(res, "subtract")

	return
}

// Utility function for EnumEffectValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumeffectvalue EnumEffectValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumeffectvalue {
	// insertion code per enum code
	case EnumEffectValueAnvil:
		res = "anvil"
	case EnumEffectValueAutoHorn:
		res = "auto horn"
	case EnumEffectValueBirdWhistle:
		res = "bird whistle"
	case EnumEffectValueCannon:
		res = "cannon"
	case EnumEffectValueDuckCall:
		res = "duck call"
	case EnumEffectValueGunShot:
		res = "gun shot"
	case EnumEffectValueKlaxonHorn:
		res = "klaxon horn"
	case EnumEffectValueLionsRoar:
		res = "lions roar"
	case EnumEffectValueLotusFlute:
		res = "lotus flute"
	case EnumEffectValueMegaphone:
		res = "megaphone"
	case EnumEffectValuePoliceWhistle:
		res = "police whistle"
	case EnumEffectValueSiren:
		res = "siren"
	case EnumEffectValueSlideWhistle:
		res = "slide whistle"
	case EnumEffectValueThunderSheet:
		res = "thunder sheet"
	case EnumEffectValueWindMachine:
		res = "wind machine"
	case EnumEffectValueWindWhistle:
		res = "wind whistle"
	}
	return
}

func (enumeffectvalue *EnumEffectValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "anvil":
		*enumeffectvalue = EnumEffectValueAnvil
		return
	case "auto horn":
		*enumeffectvalue = EnumEffectValueAutoHorn
		return
	case "bird whistle":
		*enumeffectvalue = EnumEffectValueBirdWhistle
		return
	case "cannon":
		*enumeffectvalue = EnumEffectValueCannon
		return
	case "duck call":
		*enumeffectvalue = EnumEffectValueDuckCall
		return
	case "gun shot":
		*enumeffectvalue = EnumEffectValueGunShot
		return
	case "klaxon horn":
		*enumeffectvalue = EnumEffectValueKlaxonHorn
		return
	case "lions roar":
		*enumeffectvalue = EnumEffectValueLionsRoar
		return
	case "lotus flute":
		*enumeffectvalue = EnumEffectValueLotusFlute
		return
	case "megaphone":
		*enumeffectvalue = EnumEffectValueMegaphone
		return
	case "police whistle":
		*enumeffectvalue = EnumEffectValuePoliceWhistle
		return
	case "siren":
		*enumeffectvalue = EnumEffectValueSiren
		return
	case "slide whistle":
		*enumeffectvalue = EnumEffectValueSlideWhistle
		return
	case "thunder sheet":
		*enumeffectvalue = EnumEffectValueThunderSheet
		return
	case "wind machine":
		*enumeffectvalue = EnumEffectValueWindMachine
		return
	case "wind whistle":
		*enumeffectvalue = EnumEffectValueWindWhistle
		return
	default:
		return errUnkownEnum
	}
}

func (enumeffectvalue *EnumEffectValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumEffectValueAnvil":
		*enumeffectvalue = EnumEffectValueAnvil
	case "EnumEffectValueAutoHorn":
		*enumeffectvalue = EnumEffectValueAutoHorn
	case "EnumEffectValueBirdWhistle":
		*enumeffectvalue = EnumEffectValueBirdWhistle
	case "EnumEffectValueCannon":
		*enumeffectvalue = EnumEffectValueCannon
	case "EnumEffectValueDuckCall":
		*enumeffectvalue = EnumEffectValueDuckCall
	case "EnumEffectValueGunShot":
		*enumeffectvalue = EnumEffectValueGunShot
	case "EnumEffectValueKlaxonHorn":
		*enumeffectvalue = EnumEffectValueKlaxonHorn
	case "EnumEffectValueLionsRoar":
		*enumeffectvalue = EnumEffectValueLionsRoar
	case "EnumEffectValueLotusFlute":
		*enumeffectvalue = EnumEffectValueLotusFlute
	case "EnumEffectValueMegaphone":
		*enumeffectvalue = EnumEffectValueMegaphone
	case "EnumEffectValuePoliceWhistle":
		*enumeffectvalue = EnumEffectValuePoliceWhistle
	case "EnumEffectValueSiren":
		*enumeffectvalue = EnumEffectValueSiren
	case "EnumEffectValueSlideWhistle":
		*enumeffectvalue = EnumEffectValueSlideWhistle
	case "EnumEffectValueThunderSheet":
		*enumeffectvalue = EnumEffectValueThunderSheet
	case "EnumEffectValueWindMachine":
		*enumeffectvalue = EnumEffectValueWindMachine
	case "EnumEffectValueWindWhistle":
		*enumeffectvalue = EnumEffectValueWindWhistle
	default:
		err = errUnkownEnum
	}
	return
}

func (enumeffectvalue *EnumEffectValue) ToCodeString() (res string) {

	switch *enumeffectvalue {
	// insertion code per enum code
	case EnumEffectValueAnvil:
		res = "EnumEffectValueAnvil"
	case EnumEffectValueAutoHorn:
		res = "EnumEffectValueAutoHorn"
	case EnumEffectValueBirdWhistle:
		res = "EnumEffectValueBirdWhistle"
	case EnumEffectValueCannon:
		res = "EnumEffectValueCannon"
	case EnumEffectValueDuckCall:
		res = "EnumEffectValueDuckCall"
	case EnumEffectValueGunShot:
		res = "EnumEffectValueGunShot"
	case EnumEffectValueKlaxonHorn:
		res = "EnumEffectValueKlaxonHorn"
	case EnumEffectValueLionsRoar:
		res = "EnumEffectValueLionsRoar"
	case EnumEffectValueLotusFlute:
		res = "EnumEffectValueLotusFlute"
	case EnumEffectValueMegaphone:
		res = "EnumEffectValueMegaphone"
	case EnumEffectValuePoliceWhistle:
		res = "EnumEffectValuePoliceWhistle"
	case EnumEffectValueSiren:
		res = "EnumEffectValueSiren"
	case EnumEffectValueSlideWhistle:
		res = "EnumEffectValueSlideWhistle"
	case EnumEffectValueThunderSheet:
		res = "EnumEffectValueThunderSheet"
	case EnumEffectValueWindMachine:
		res = "EnumEffectValueWindMachine"
	case EnumEffectValueWindWhistle:
		res = "EnumEffectValueWindWhistle"
	}
	return
}

func (enumeffectvalue EnumEffectValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumEffectValueAnvil")
	res = append(res, "EnumEffectValueAutoHorn")
	res = append(res, "EnumEffectValueBirdWhistle")
	res = append(res, "EnumEffectValueCannon")
	res = append(res, "EnumEffectValueDuckCall")
	res = append(res, "EnumEffectValueGunShot")
	res = append(res, "EnumEffectValueKlaxonHorn")
	res = append(res, "EnumEffectValueLionsRoar")
	res = append(res, "EnumEffectValueLotusFlute")
	res = append(res, "EnumEffectValueMegaphone")
	res = append(res, "EnumEffectValuePoliceWhistle")
	res = append(res, "EnumEffectValueSiren")
	res = append(res, "EnumEffectValueSlideWhistle")
	res = append(res, "EnumEffectValueThunderSheet")
	res = append(res, "EnumEffectValueWindMachine")
	res = append(res, "EnumEffectValueWindWhistle")

	return
}

func (enumeffectvalue EnumEffectValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "anvil")
	res = append(res, "auto horn")
	res = append(res, "bird whistle")
	res = append(res, "cannon")
	res = append(res, "duck call")
	res = append(res, "gun shot")
	res = append(res, "klaxon horn")
	res = append(res, "lions roar")
	res = append(res, "lotus flute")
	res = append(res, "megaphone")
	res = append(res, "police whistle")
	res = append(res, "siren")
	res = append(res, "slide whistle")
	res = append(res, "thunder sheet")
	res = append(res, "wind machine")
	res = append(res, "wind whistle")

	return
}

// Utility function for EnumGlassValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumglassvalue EnumGlassValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumglassvalue {
	// insertion code per enum code
	case EnumGlassValueGlassHarmonica:
		res = "glass harmonica"
	case EnumGlassValueGlassHarp:
		res = "glass harp"
	case EnumGlassValueWindChimes:
		res = "wind chimes"
	}
	return
}

func (enumglassvalue *EnumGlassValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "glass harmonica":
		*enumglassvalue = EnumGlassValueGlassHarmonica
		return
	case "glass harp":
		*enumglassvalue = EnumGlassValueGlassHarp
		return
	case "wind chimes":
		*enumglassvalue = EnumGlassValueWindChimes
		return
	default:
		return errUnkownEnum
	}
}

func (enumglassvalue *EnumGlassValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumGlassValueGlassHarmonica":
		*enumglassvalue = EnumGlassValueGlassHarmonica
	case "EnumGlassValueGlassHarp":
		*enumglassvalue = EnumGlassValueGlassHarp
	case "EnumGlassValueWindChimes":
		*enumglassvalue = EnumGlassValueWindChimes
	default:
		err = errUnkownEnum
	}
	return
}

func (enumglassvalue *EnumGlassValue) ToCodeString() (res string) {

	switch *enumglassvalue {
	// insertion code per enum code
	case EnumGlassValueGlassHarmonica:
		res = "EnumGlassValueGlassHarmonica"
	case EnumGlassValueGlassHarp:
		res = "EnumGlassValueGlassHarp"
	case EnumGlassValueWindChimes:
		res = "EnumGlassValueWindChimes"
	}
	return
}

func (enumglassvalue EnumGlassValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumGlassValueGlassHarmonica")
	res = append(res, "EnumGlassValueGlassHarp")
	res = append(res, "EnumGlassValueWindChimes")

	return
}

func (enumglassvalue EnumGlassValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "glass harmonica")
	res = append(res, "glass harp")
	res = append(res, "wind chimes")

	return
}

// Utility function for EnumHarmonyArrangement
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumharmonyarrangement EnumHarmonyArrangement) ToString() (res string) {

	// migration of former implementation of enum
	switch enumharmonyarrangement {
	// insertion code per enum code
	case EnumHarmonyArrangementVertical:
		res = "vertical"
	case EnumHarmonyArrangementHorizontal:
		res = "horizontal"
	case EnumHarmonyArrangementDiagonal:
		res = "diagonal"
	}
	return
}

func (enumharmonyarrangement *EnumHarmonyArrangement) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "vertical":
		*enumharmonyarrangement = EnumHarmonyArrangementVertical
		return
	case "horizontal":
		*enumharmonyarrangement = EnumHarmonyArrangementHorizontal
		return
	case "diagonal":
		*enumharmonyarrangement = EnumHarmonyArrangementDiagonal
		return
	default:
		return errUnkownEnum
	}
}

func (enumharmonyarrangement *EnumHarmonyArrangement) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHarmonyArrangementVertical":
		*enumharmonyarrangement = EnumHarmonyArrangementVertical
	case "EnumHarmonyArrangementHorizontal":
		*enumharmonyarrangement = EnumHarmonyArrangementHorizontal
	case "EnumHarmonyArrangementDiagonal":
		*enumharmonyarrangement = EnumHarmonyArrangementDiagonal
	default:
		err = errUnkownEnum
	}
	return
}

func (enumharmonyarrangement *EnumHarmonyArrangement) ToCodeString() (res string) {

	switch *enumharmonyarrangement {
	// insertion code per enum code
	case EnumHarmonyArrangementVertical:
		res = "EnumHarmonyArrangementVertical"
	case EnumHarmonyArrangementHorizontal:
		res = "EnumHarmonyArrangementHorizontal"
	case EnumHarmonyArrangementDiagonal:
		res = "EnumHarmonyArrangementDiagonal"
	}
	return
}

func (enumharmonyarrangement EnumHarmonyArrangement) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHarmonyArrangementVertical")
	res = append(res, "EnumHarmonyArrangementHorizontal")
	res = append(res, "EnumHarmonyArrangementDiagonal")

	return
}

func (enumharmonyarrangement EnumHarmonyArrangement) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "vertical")
	res = append(res, "horizontal")
	res = append(res, "diagonal")

	return
}

// Utility function for EnumHarmonyType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumharmonytype EnumHarmonyType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumharmonytype {
	// insertion code per enum code
	case EnumHarmonyTypeExplicit:
		res = "explicit"
	case EnumHarmonyTypeImplied:
		res = "implied"
	case EnumHarmonyTypeAlternate:
		res = "alternate"
	}
	return
}

func (enumharmonytype *EnumHarmonyType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "explicit":
		*enumharmonytype = EnumHarmonyTypeExplicit
		return
	case "implied":
		*enumharmonytype = EnumHarmonyTypeImplied
		return
	case "alternate":
		*enumharmonytype = EnumHarmonyTypeAlternate
		return
	default:
		return errUnkownEnum
	}
}

func (enumharmonytype *EnumHarmonyType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumHarmonyTypeExplicit":
		*enumharmonytype = EnumHarmonyTypeExplicit
	case "EnumHarmonyTypeImplied":
		*enumharmonytype = EnumHarmonyTypeImplied
	case "EnumHarmonyTypeAlternate":
		*enumharmonytype = EnumHarmonyTypeAlternate
	default:
		err = errUnkownEnum
	}
	return
}

func (enumharmonytype *EnumHarmonyType) ToCodeString() (res string) {

	switch *enumharmonytype {
	// insertion code per enum code
	case EnumHarmonyTypeExplicit:
		res = "EnumHarmonyTypeExplicit"
	case EnumHarmonyTypeImplied:
		res = "EnumHarmonyTypeImplied"
	case EnumHarmonyTypeAlternate:
		res = "EnumHarmonyTypeAlternate"
	}
	return
}

func (enumharmonytype EnumHarmonyType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumHarmonyTypeExplicit")
	res = append(res, "EnumHarmonyTypeImplied")
	res = append(res, "EnumHarmonyTypeAlternate")

	return
}

func (enumharmonytype EnumHarmonyType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "explicit")
	res = append(res, "implied")
	res = append(res, "alternate")

	return
}

// Utility function for EnumKindValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumkindvalue EnumKindValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumkindvalue {
	// insertion code per enum code
	case EnumKindValueMajor:
		res = "major"
	case EnumKindValueMinor:
		res = "minor"
	case EnumKindValueAugmented:
		res = "augmented"
	case EnumKindValueDiminished:
		res = "diminished"
	case EnumKindValueDominant:
		res = "dominant"
	case EnumKindValueMajorSeventh:
		res = "major-seventh"
	case EnumKindValueMinorSeventh:
		res = "minor-seventh"
	case EnumKindValueDiminishedSeventh:
		res = "diminished-seventh"
	case EnumKindValueAugmentedSeventh:
		res = "augmented-seventh"
	case EnumKindValueHalfDiminished:
		res = "half-diminished"
	case EnumKindValueMajorMinor:
		res = "major-minor"
	case EnumKindValueMajorSixth:
		res = "major-sixth"
	case EnumKindValueMinorSixth:
		res = "minor-sixth"
	case EnumKindValueDominantNinth:
		res = "dominant-ninth"
	case EnumKindValueMajorNinth:
		res = "major-ninth"
	case EnumKindValueMinorNinth:
		res = "minor-ninth"
	case EnumKindValueDominant11th:
		res = "dominant-11th"
	case EnumKindValueMajor11th:
		res = "major-11th"
	case EnumKindValueMinor11th:
		res = "minor-11th"
	case EnumKindValueDominant13th:
		res = "dominant-13th"
	case EnumKindValueMajor13th:
		res = "major-13th"
	case EnumKindValueMinor13th:
		res = "minor-13th"
	case EnumKindValueSuspendedSecond:
		res = "suspended-second"
	case EnumKindValueSuspendedFourth:
		res = "suspended-fourth"
	case EnumKindValueNeapolitan:
		res = "Neapolitan"
	case EnumKindValueItalian:
		res = "Italian"
	case EnumKindValueFrench:
		res = "French"
	case EnumKindValueGerman:
		res = "German"
	case EnumKindValuePedal:
		res = "pedal"
	case EnumKindValuePower:
		res = "power"
	case EnumKindValueTristan:
		res = "Tristan"
	case EnumKindValueOther:
		res = "other"
	case EnumKindValueNone:
		res = "none"
	}
	return
}

func (enumkindvalue *EnumKindValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enumkindvalue = EnumKindValueMajor
		return
	case "minor":
		*enumkindvalue = EnumKindValueMinor
		return
	case "augmented":
		*enumkindvalue = EnumKindValueAugmented
		return
	case "diminished":
		*enumkindvalue = EnumKindValueDiminished
		return
	case "dominant":
		*enumkindvalue = EnumKindValueDominant
		return
	case "major-seventh":
		*enumkindvalue = EnumKindValueMajorSeventh
		return
	case "minor-seventh":
		*enumkindvalue = EnumKindValueMinorSeventh
		return
	case "diminished-seventh":
		*enumkindvalue = EnumKindValueDiminishedSeventh
		return
	case "augmented-seventh":
		*enumkindvalue = EnumKindValueAugmentedSeventh
		return
	case "half-diminished":
		*enumkindvalue = EnumKindValueHalfDiminished
		return
	case "major-minor":
		*enumkindvalue = EnumKindValueMajorMinor
		return
	case "major-sixth":
		*enumkindvalue = EnumKindValueMajorSixth
		return
	case "minor-sixth":
		*enumkindvalue = EnumKindValueMinorSixth
		return
	case "dominant-ninth":
		*enumkindvalue = EnumKindValueDominantNinth
		return
	case "major-ninth":
		*enumkindvalue = EnumKindValueMajorNinth
		return
	case "minor-ninth":
		*enumkindvalue = EnumKindValueMinorNinth
		return
	case "dominant-11th":
		*enumkindvalue = EnumKindValueDominant11th
		return
	case "major-11th":
		*enumkindvalue = EnumKindValueMajor11th
		return
	case "minor-11th":
		*enumkindvalue = EnumKindValueMinor11th
		return
	case "dominant-13th":
		*enumkindvalue = EnumKindValueDominant13th
		return
	case "major-13th":
		*enumkindvalue = EnumKindValueMajor13th
		return
	case "minor-13th":
		*enumkindvalue = EnumKindValueMinor13th
		return
	case "suspended-second":
		*enumkindvalue = EnumKindValueSuspendedSecond
		return
	case "suspended-fourth":
		*enumkindvalue = EnumKindValueSuspendedFourth
		return
	case "Neapolitan":
		*enumkindvalue = EnumKindValueNeapolitan
		return
	case "Italian":
		*enumkindvalue = EnumKindValueItalian
		return
	case "French":
		*enumkindvalue = EnumKindValueFrench
		return
	case "German":
		*enumkindvalue = EnumKindValueGerman
		return
	case "pedal":
		*enumkindvalue = EnumKindValuePedal
		return
	case "power":
		*enumkindvalue = EnumKindValuePower
		return
	case "Tristan":
		*enumkindvalue = EnumKindValueTristan
		return
	case "other":
		*enumkindvalue = EnumKindValueOther
		return
	case "none":
		*enumkindvalue = EnumKindValueNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumkindvalue *EnumKindValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumKindValueMajor":
		*enumkindvalue = EnumKindValueMajor
	case "EnumKindValueMinor":
		*enumkindvalue = EnumKindValueMinor
	case "EnumKindValueAugmented":
		*enumkindvalue = EnumKindValueAugmented
	case "EnumKindValueDiminished":
		*enumkindvalue = EnumKindValueDiminished
	case "EnumKindValueDominant":
		*enumkindvalue = EnumKindValueDominant
	case "EnumKindValueMajorSeventh":
		*enumkindvalue = EnumKindValueMajorSeventh
	case "EnumKindValueMinorSeventh":
		*enumkindvalue = EnumKindValueMinorSeventh
	case "EnumKindValueDiminishedSeventh":
		*enumkindvalue = EnumKindValueDiminishedSeventh
	case "EnumKindValueAugmentedSeventh":
		*enumkindvalue = EnumKindValueAugmentedSeventh
	case "EnumKindValueHalfDiminished":
		*enumkindvalue = EnumKindValueHalfDiminished
	case "EnumKindValueMajorMinor":
		*enumkindvalue = EnumKindValueMajorMinor
	case "EnumKindValueMajorSixth":
		*enumkindvalue = EnumKindValueMajorSixth
	case "EnumKindValueMinorSixth":
		*enumkindvalue = EnumKindValueMinorSixth
	case "EnumKindValueDominantNinth":
		*enumkindvalue = EnumKindValueDominantNinth
	case "EnumKindValueMajorNinth":
		*enumkindvalue = EnumKindValueMajorNinth
	case "EnumKindValueMinorNinth":
		*enumkindvalue = EnumKindValueMinorNinth
	case "EnumKindValueDominant11th":
		*enumkindvalue = EnumKindValueDominant11th
	case "EnumKindValueMajor11th":
		*enumkindvalue = EnumKindValueMajor11th
	case "EnumKindValueMinor11th":
		*enumkindvalue = EnumKindValueMinor11th
	case "EnumKindValueDominant13th":
		*enumkindvalue = EnumKindValueDominant13th
	case "EnumKindValueMajor13th":
		*enumkindvalue = EnumKindValueMajor13th
	case "EnumKindValueMinor13th":
		*enumkindvalue = EnumKindValueMinor13th
	case "EnumKindValueSuspendedSecond":
		*enumkindvalue = EnumKindValueSuspendedSecond
	case "EnumKindValueSuspendedFourth":
		*enumkindvalue = EnumKindValueSuspendedFourth
	case "EnumKindValueNeapolitan":
		*enumkindvalue = EnumKindValueNeapolitan
	case "EnumKindValueItalian":
		*enumkindvalue = EnumKindValueItalian
	case "EnumKindValueFrench":
		*enumkindvalue = EnumKindValueFrench
	case "EnumKindValueGerman":
		*enumkindvalue = EnumKindValueGerman
	case "EnumKindValuePedal":
		*enumkindvalue = EnumKindValuePedal
	case "EnumKindValuePower":
		*enumkindvalue = EnumKindValuePower
	case "EnumKindValueTristan":
		*enumkindvalue = EnumKindValueTristan
	case "EnumKindValueOther":
		*enumkindvalue = EnumKindValueOther
	case "EnumKindValueNone":
		*enumkindvalue = EnumKindValueNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumkindvalue *EnumKindValue) ToCodeString() (res string) {

	switch *enumkindvalue {
	// insertion code per enum code
	case EnumKindValueMajor:
		res = "EnumKindValueMajor"
	case EnumKindValueMinor:
		res = "EnumKindValueMinor"
	case EnumKindValueAugmented:
		res = "EnumKindValueAugmented"
	case EnumKindValueDiminished:
		res = "EnumKindValueDiminished"
	case EnumKindValueDominant:
		res = "EnumKindValueDominant"
	case EnumKindValueMajorSeventh:
		res = "EnumKindValueMajorSeventh"
	case EnumKindValueMinorSeventh:
		res = "EnumKindValueMinorSeventh"
	case EnumKindValueDiminishedSeventh:
		res = "EnumKindValueDiminishedSeventh"
	case EnumKindValueAugmentedSeventh:
		res = "EnumKindValueAugmentedSeventh"
	case EnumKindValueHalfDiminished:
		res = "EnumKindValueHalfDiminished"
	case EnumKindValueMajorMinor:
		res = "EnumKindValueMajorMinor"
	case EnumKindValueMajorSixth:
		res = "EnumKindValueMajorSixth"
	case EnumKindValueMinorSixth:
		res = "EnumKindValueMinorSixth"
	case EnumKindValueDominantNinth:
		res = "EnumKindValueDominantNinth"
	case EnumKindValueMajorNinth:
		res = "EnumKindValueMajorNinth"
	case EnumKindValueMinorNinth:
		res = "EnumKindValueMinorNinth"
	case EnumKindValueDominant11th:
		res = "EnumKindValueDominant11th"
	case EnumKindValueMajor11th:
		res = "EnumKindValueMajor11th"
	case EnumKindValueMinor11th:
		res = "EnumKindValueMinor11th"
	case EnumKindValueDominant13th:
		res = "EnumKindValueDominant13th"
	case EnumKindValueMajor13th:
		res = "EnumKindValueMajor13th"
	case EnumKindValueMinor13th:
		res = "EnumKindValueMinor13th"
	case EnumKindValueSuspendedSecond:
		res = "EnumKindValueSuspendedSecond"
	case EnumKindValueSuspendedFourth:
		res = "EnumKindValueSuspendedFourth"
	case EnumKindValueNeapolitan:
		res = "EnumKindValueNeapolitan"
	case EnumKindValueItalian:
		res = "EnumKindValueItalian"
	case EnumKindValueFrench:
		res = "EnumKindValueFrench"
	case EnumKindValueGerman:
		res = "EnumKindValueGerman"
	case EnumKindValuePedal:
		res = "EnumKindValuePedal"
	case EnumKindValuePower:
		res = "EnumKindValuePower"
	case EnumKindValueTristan:
		res = "EnumKindValueTristan"
	case EnumKindValueOther:
		res = "EnumKindValueOther"
	case EnumKindValueNone:
		res = "EnumKindValueNone"
	}
	return
}

func (enumkindvalue EnumKindValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumKindValueMajor")
	res = append(res, "EnumKindValueMinor")
	res = append(res, "EnumKindValueAugmented")
	res = append(res, "EnumKindValueDiminished")
	res = append(res, "EnumKindValueDominant")
	res = append(res, "EnumKindValueMajorSeventh")
	res = append(res, "EnumKindValueMinorSeventh")
	res = append(res, "EnumKindValueDiminishedSeventh")
	res = append(res, "EnumKindValueAugmentedSeventh")
	res = append(res, "EnumKindValueHalfDiminished")
	res = append(res, "EnumKindValueMajorMinor")
	res = append(res, "EnumKindValueMajorSixth")
	res = append(res, "EnumKindValueMinorSixth")
	res = append(res, "EnumKindValueDominantNinth")
	res = append(res, "EnumKindValueMajorNinth")
	res = append(res, "EnumKindValueMinorNinth")
	res = append(res, "EnumKindValueDominant11th")
	res = append(res, "EnumKindValueMajor11th")
	res = append(res, "EnumKindValueMinor11th")
	res = append(res, "EnumKindValueDominant13th")
	res = append(res, "EnumKindValueMajor13th")
	res = append(res, "EnumKindValueMinor13th")
	res = append(res, "EnumKindValueSuspendedSecond")
	res = append(res, "EnumKindValueSuspendedFourth")
	res = append(res, "EnumKindValueNeapolitan")
	res = append(res, "EnumKindValueItalian")
	res = append(res, "EnumKindValueFrench")
	res = append(res, "EnumKindValueGerman")
	res = append(res, "EnumKindValuePedal")
	res = append(res, "EnumKindValuePower")
	res = append(res, "EnumKindValueTristan")
	res = append(res, "EnumKindValueOther")
	res = append(res, "EnumKindValueNone")

	return
}

func (enumkindvalue EnumKindValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "augmented")
	res = append(res, "diminished")
	res = append(res, "dominant")
	res = append(res, "major-seventh")
	res = append(res, "minor-seventh")
	res = append(res, "diminished-seventh")
	res = append(res, "augmented-seventh")
	res = append(res, "half-diminished")
	res = append(res, "major-minor")
	res = append(res, "major-sixth")
	res = append(res, "minor-sixth")
	res = append(res, "dominant-ninth")
	res = append(res, "major-ninth")
	res = append(res, "minor-ninth")
	res = append(res, "dominant-11th")
	res = append(res, "major-11th")
	res = append(res, "minor-11th")
	res = append(res, "dominant-13th")
	res = append(res, "major-13th")
	res = append(res, "minor-13th")
	res = append(res, "suspended-second")
	res = append(res, "suspended-fourth")
	res = append(res, "Neapolitan")
	res = append(res, "Italian")
	res = append(res, "French")
	res = append(res, "German")
	res = append(res, "pedal")
	res = append(res, "power")
	res = append(res, "Tristan")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for EnumLineEnd
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumlineend EnumLineEnd) ToString() (res string) {

	// migration of former implementation of enum
	switch enumlineend {
	// insertion code per enum code
	case EnumLineEndUp:
		res = "up"
	case EnumLineEndDown:
		res = "down"
	case EnumLineEndBoth:
		res = "both"
	case EnumLineEndArrow:
		res = "arrow"
	case EnumLineEndNone:
		res = "none"
	}
	return
}

func (enumlineend *EnumLineEnd) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enumlineend = EnumLineEndUp
		return
	case "down":
		*enumlineend = EnumLineEndDown
		return
	case "both":
		*enumlineend = EnumLineEndBoth
		return
	case "arrow":
		*enumlineend = EnumLineEndArrow
		return
	case "none":
		*enumlineend = EnumLineEndNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumlineend *EnumLineEnd) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLineEndUp":
		*enumlineend = EnumLineEndUp
	case "EnumLineEndDown":
		*enumlineend = EnumLineEndDown
	case "EnumLineEndBoth":
		*enumlineend = EnumLineEndBoth
	case "EnumLineEndArrow":
		*enumlineend = EnumLineEndArrow
	case "EnumLineEndNone":
		*enumlineend = EnumLineEndNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumlineend *EnumLineEnd) ToCodeString() (res string) {

	switch *enumlineend {
	// insertion code per enum code
	case EnumLineEndUp:
		res = "EnumLineEndUp"
	case EnumLineEndDown:
		res = "EnumLineEndDown"
	case EnumLineEndBoth:
		res = "EnumLineEndBoth"
	case EnumLineEndArrow:
		res = "EnumLineEndArrow"
	case EnumLineEndNone:
		res = "EnumLineEndNone"
	}
	return
}

func (enumlineend EnumLineEnd) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLineEndUp")
	res = append(res, "EnumLineEndDown")
	res = append(res, "EnumLineEndBoth")
	res = append(res, "EnumLineEndArrow")
	res = append(res, "EnumLineEndNone")

	return
}

func (enumlineend EnumLineEnd) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "both")
	res = append(res, "arrow")
	res = append(res, "none")

	return
}

// Utility function for EnumMeasureNumberingValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enummeasurenumberingvalue EnumMeasureNumberingValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enummeasurenumberingvalue {
	// insertion code per enum code
	case EnumMeasureNumberingValueNone:
		res = "none"
	case EnumMeasureNumberingValueMeasure:
		res = "measure"
	case EnumMeasureNumberingValueSystem:
		res = "system"
	}
	return
}

func (enummeasurenumberingvalue *EnumMeasureNumberingValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueNone
		return
	case "measure":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueMeasure
		return
	case "system":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueSystem
		return
	default:
		return errUnkownEnum
	}
}

func (enummeasurenumberingvalue *EnumMeasureNumberingValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumMeasureNumberingValueNone":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueNone
	case "EnumMeasureNumberingValueMeasure":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueMeasure
	case "EnumMeasureNumberingValueSystem":
		*enummeasurenumberingvalue = EnumMeasureNumberingValueSystem
	default:
		err = errUnkownEnum
	}
	return
}

func (enummeasurenumberingvalue *EnumMeasureNumberingValue) ToCodeString() (res string) {

	switch *enummeasurenumberingvalue {
	// insertion code per enum code
	case EnumMeasureNumberingValueNone:
		res = "EnumMeasureNumberingValueNone"
	case EnumMeasureNumberingValueMeasure:
		res = "EnumMeasureNumberingValueMeasure"
	case EnumMeasureNumberingValueSystem:
		res = "EnumMeasureNumberingValueSystem"
	}
	return
}

func (enummeasurenumberingvalue EnumMeasureNumberingValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumMeasureNumberingValueNone")
	res = append(res, "EnumMeasureNumberingValueMeasure")
	res = append(res, "EnumMeasureNumberingValueSystem")

	return
}

func (enummeasurenumberingvalue EnumMeasureNumberingValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "measure")
	res = append(res, "system")

	return
}

// Utility function for EnumMembraneValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enummembranevalue EnumMembraneValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enummembranevalue {
	// insertion code per enum code
	case EnumMembraneValueBassDrum:
		res = "bass drum"
	case EnumMembraneValueBassDrumOnSide:
		res = "bass drum on side"
	case EnumMembraneValueBongos:
		res = "bongos"
	case EnumMembraneValueChineseTomtom:
		res = "Chinese tomtom"
	case EnumMembraneValueCongaDrum:
		res = "conga drum"
	case EnumMembraneValueCuica:
		res = "cuica"
	case EnumMembraneValueGobletDrum:
		res = "goblet drum"
	case EnumMembraneValueIndoAmericanTomtom:
		res = "Indo-American tomtom"
	case EnumMembraneValueJapaneseTomtom:
		res = "Japanese tomtom"
	case EnumMembraneValueMilitaryDrum:
		res = "military drum"
	case EnumMembraneValueSnareDrum:
		res = "snare drum"
	case EnumMembraneValueSnareDrumSnaresOff:
		res = "snare drum snares off"
	case EnumMembraneValueTabla:
		res = "tabla"
	case EnumMembraneValueTambourine:
		res = "tambourine"
	case EnumMembraneValueTenorDrum:
		res = "tenor drum"
	case EnumMembraneValueTimbales:
		res = "timbales"
	case EnumMembraneValueTomtom:
		res = "tomtom"
	}
	return
}

func (enummembranevalue *EnumMembraneValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*enummembranevalue = EnumMembraneValueBassDrum
		return
	case "bass drum on side":
		*enummembranevalue = EnumMembraneValueBassDrumOnSide
		return
	case "bongos":
		*enummembranevalue = EnumMembraneValueBongos
		return
	case "Chinese tomtom":
		*enummembranevalue = EnumMembraneValueChineseTomtom
		return
	case "conga drum":
		*enummembranevalue = EnumMembraneValueCongaDrum
		return
	case "cuica":
		*enummembranevalue = EnumMembraneValueCuica
		return
	case "goblet drum":
		*enummembranevalue = EnumMembraneValueGobletDrum
		return
	case "Indo-American tomtom":
		*enummembranevalue = EnumMembraneValueIndoAmericanTomtom
		return
	case "Japanese tomtom":
		*enummembranevalue = EnumMembraneValueJapaneseTomtom
		return
	case "military drum":
		*enummembranevalue = EnumMembraneValueMilitaryDrum
		return
	case "snare drum":
		*enummembranevalue = EnumMembraneValueSnareDrum
		return
	case "snare drum snares off":
		*enummembranevalue = EnumMembraneValueSnareDrumSnaresOff
		return
	case "tabla":
		*enummembranevalue = EnumMembraneValueTabla
		return
	case "tambourine":
		*enummembranevalue = EnumMembraneValueTambourine
		return
	case "tenor drum":
		*enummembranevalue = EnumMembraneValueTenorDrum
		return
	case "timbales":
		*enummembranevalue = EnumMembraneValueTimbales
		return
	case "tomtom":
		*enummembranevalue = EnumMembraneValueTomtom
		return
	default:
		return errUnkownEnum
	}
}

func (enummembranevalue *EnumMembraneValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumMembraneValueBassDrum":
		*enummembranevalue = EnumMembraneValueBassDrum
	case "EnumMembraneValueBassDrumOnSide":
		*enummembranevalue = EnumMembraneValueBassDrumOnSide
	case "EnumMembraneValueBongos":
		*enummembranevalue = EnumMembraneValueBongos
	case "EnumMembraneValueChineseTomtom":
		*enummembranevalue = EnumMembraneValueChineseTomtom
	case "EnumMembraneValueCongaDrum":
		*enummembranevalue = EnumMembraneValueCongaDrum
	case "EnumMembraneValueCuica":
		*enummembranevalue = EnumMembraneValueCuica
	case "EnumMembraneValueGobletDrum":
		*enummembranevalue = EnumMembraneValueGobletDrum
	case "EnumMembraneValueIndoAmericanTomtom":
		*enummembranevalue = EnumMembraneValueIndoAmericanTomtom
	case "EnumMembraneValueJapaneseTomtom":
		*enummembranevalue = EnumMembraneValueJapaneseTomtom
	case "EnumMembraneValueMilitaryDrum":
		*enummembranevalue = EnumMembraneValueMilitaryDrum
	case "EnumMembraneValueSnareDrum":
		*enummembranevalue = EnumMembraneValueSnareDrum
	case "EnumMembraneValueSnareDrumSnaresOff":
		*enummembranevalue = EnumMembraneValueSnareDrumSnaresOff
	case "EnumMembraneValueTabla":
		*enummembranevalue = EnumMembraneValueTabla
	case "EnumMembraneValueTambourine":
		*enummembranevalue = EnumMembraneValueTambourine
	case "EnumMembraneValueTenorDrum":
		*enummembranevalue = EnumMembraneValueTenorDrum
	case "EnumMembraneValueTimbales":
		*enummembranevalue = EnumMembraneValueTimbales
	case "EnumMembraneValueTomtom":
		*enummembranevalue = EnumMembraneValueTomtom
	default:
		err = errUnkownEnum
	}
	return
}

func (enummembranevalue *EnumMembraneValue) ToCodeString() (res string) {

	switch *enummembranevalue {
	// insertion code per enum code
	case EnumMembraneValueBassDrum:
		res = "EnumMembraneValueBassDrum"
	case EnumMembraneValueBassDrumOnSide:
		res = "EnumMembraneValueBassDrumOnSide"
	case EnumMembraneValueBongos:
		res = "EnumMembraneValueBongos"
	case EnumMembraneValueChineseTomtom:
		res = "EnumMembraneValueChineseTomtom"
	case EnumMembraneValueCongaDrum:
		res = "EnumMembraneValueCongaDrum"
	case EnumMembraneValueCuica:
		res = "EnumMembraneValueCuica"
	case EnumMembraneValueGobletDrum:
		res = "EnumMembraneValueGobletDrum"
	case EnumMembraneValueIndoAmericanTomtom:
		res = "EnumMembraneValueIndoAmericanTomtom"
	case EnumMembraneValueJapaneseTomtom:
		res = "EnumMembraneValueJapaneseTomtom"
	case EnumMembraneValueMilitaryDrum:
		res = "EnumMembraneValueMilitaryDrum"
	case EnumMembraneValueSnareDrum:
		res = "EnumMembraneValueSnareDrum"
	case EnumMembraneValueSnareDrumSnaresOff:
		res = "EnumMembraneValueSnareDrumSnaresOff"
	case EnumMembraneValueTabla:
		res = "EnumMembraneValueTabla"
	case EnumMembraneValueTambourine:
		res = "EnumMembraneValueTambourine"
	case EnumMembraneValueTenorDrum:
		res = "EnumMembraneValueTenorDrum"
	case EnumMembraneValueTimbales:
		res = "EnumMembraneValueTimbales"
	case EnumMembraneValueTomtom:
		res = "EnumMembraneValueTomtom"
	}
	return
}

func (enummembranevalue EnumMembraneValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumMembraneValueBassDrum")
	res = append(res, "EnumMembraneValueBassDrumOnSide")
	res = append(res, "EnumMembraneValueBongos")
	res = append(res, "EnumMembraneValueChineseTomtom")
	res = append(res, "EnumMembraneValueCongaDrum")
	res = append(res, "EnumMembraneValueCuica")
	res = append(res, "EnumMembraneValueGobletDrum")
	res = append(res, "EnumMembraneValueIndoAmericanTomtom")
	res = append(res, "EnumMembraneValueJapaneseTomtom")
	res = append(res, "EnumMembraneValueMilitaryDrum")
	res = append(res, "EnumMembraneValueSnareDrum")
	res = append(res, "EnumMembraneValueSnareDrumSnaresOff")
	res = append(res, "EnumMembraneValueTabla")
	res = append(res, "EnumMembraneValueTambourine")
	res = append(res, "EnumMembraneValueTenorDrum")
	res = append(res, "EnumMembraneValueTimbales")
	res = append(res, "EnumMembraneValueTomtom")

	return
}

func (enummembranevalue EnumMembraneValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bass drum")
	res = append(res, "bass drum on side")
	res = append(res, "bongos")
	res = append(res, "Chinese tomtom")
	res = append(res, "conga drum")
	res = append(res, "cuica")
	res = append(res, "goblet drum")
	res = append(res, "Indo-American tomtom")
	res = append(res, "Japanese tomtom")
	res = append(res, "military drum")
	res = append(res, "snare drum")
	res = append(res, "snare drum snares off")
	res = append(res, "tabla")
	res = append(res, "tambourine")
	res = append(res, "tenor drum")
	res = append(res, "timbales")
	res = append(res, "tomtom")

	return
}

// Utility function for EnumMetalValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enummetalvalue EnumMetalValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enummetalvalue {
	// insertion code per enum code
	case EnumMetalValueAgogo:
		res = "agogo"
	case EnumMetalValueAlmglocken:
		res = "almglocken"
	case EnumMetalValueBell:
		res = "bell"
	case EnumMetalValueBellPlate:
		res = "bell plate"
	case EnumMetalValueBellTree:
		res = "bell tree"
	case EnumMetalValueBrakeDrum:
		res = "brake drum"
	case EnumMetalValueCencerro:
		res = "cencerro"
	case EnumMetalValueChainRattle:
		res = "chain rattle"
	case EnumMetalValueChineseCymbal:
		res = "Chinese cymbal"
	case EnumMetalValueCowbell:
		res = "cowbell"
	case EnumMetalValueCrashCymbals:
		res = "crash cymbals"
	case EnumMetalValueCrotale:
		res = "crotale"
	case EnumMetalValueCymbalTongs:
		res = "cymbal tongs"
	case EnumMetalValueDomedGong:
		res = "domed gong"
	case EnumMetalValueFingerCymbals:
		res = "finger cymbals"
	case EnumMetalValueFlexatone:
		res = "flexatone"
	case EnumMetalValueGong:
		res = "gong"
	case EnumMetalValueHiHat:
		res = "hi-hat"
	case EnumMetalValueHighHatCymbals:
		res = "high-hat cymbals"
	case EnumMetalValueHandbell:
		res = "handbell"
	case EnumMetalValueJawHarp:
		res = "jaw harp"
	case EnumMetalValueJingleBells:
		res = "jingle bells"
	case EnumMetalValueMusicalSaw:
		res = "musical saw"
	case EnumMetalValueShellBells:
		res = "shell bells"
	case EnumMetalValueSistrum:
		res = "sistrum"
	case EnumMetalValueSizzleCymbal:
		res = "sizzle cymbal"
	case EnumMetalValueSleighBells:
		res = "sleigh bells"
	case EnumMetalValueSuspendedCymbal:
		res = "suspended cymbal"
	case EnumMetalValueTamTam:
		res = "tam tam"
	case EnumMetalValueTamTamWithBeater:
		res = "tam tam with beater"
	case EnumMetalValueTriangle:
		res = "triangle"
	case EnumMetalValueVietnameseHat:
		res = "Vietnamese hat"
	}
	return
}

func (enummetalvalue *EnumMetalValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "agogo":
		*enummetalvalue = EnumMetalValueAgogo
		return
	case "almglocken":
		*enummetalvalue = EnumMetalValueAlmglocken
		return
	case "bell":
		*enummetalvalue = EnumMetalValueBell
		return
	case "bell plate":
		*enummetalvalue = EnumMetalValueBellPlate
		return
	case "bell tree":
		*enummetalvalue = EnumMetalValueBellTree
		return
	case "brake drum":
		*enummetalvalue = EnumMetalValueBrakeDrum
		return
	case "cencerro":
		*enummetalvalue = EnumMetalValueCencerro
		return
	case "chain rattle":
		*enummetalvalue = EnumMetalValueChainRattle
		return
	case "Chinese cymbal":
		*enummetalvalue = EnumMetalValueChineseCymbal
		return
	case "cowbell":
		*enummetalvalue = EnumMetalValueCowbell
		return
	case "crash cymbals":
		*enummetalvalue = EnumMetalValueCrashCymbals
		return
	case "crotale":
		*enummetalvalue = EnumMetalValueCrotale
		return
	case "cymbal tongs":
		*enummetalvalue = EnumMetalValueCymbalTongs
		return
	case "domed gong":
		*enummetalvalue = EnumMetalValueDomedGong
		return
	case "finger cymbals":
		*enummetalvalue = EnumMetalValueFingerCymbals
		return
	case "flexatone":
		*enummetalvalue = EnumMetalValueFlexatone
		return
	case "gong":
		*enummetalvalue = EnumMetalValueGong
		return
	case "hi-hat":
		*enummetalvalue = EnumMetalValueHiHat
		return
	case "high-hat cymbals":
		*enummetalvalue = EnumMetalValueHighHatCymbals
		return
	case "handbell":
		*enummetalvalue = EnumMetalValueHandbell
		return
	case "jaw harp":
		*enummetalvalue = EnumMetalValueJawHarp
		return
	case "jingle bells":
		*enummetalvalue = EnumMetalValueJingleBells
		return
	case "musical saw":
		*enummetalvalue = EnumMetalValueMusicalSaw
		return
	case "shell bells":
		*enummetalvalue = EnumMetalValueShellBells
		return
	case "sistrum":
		*enummetalvalue = EnumMetalValueSistrum
		return
	case "sizzle cymbal":
		*enummetalvalue = EnumMetalValueSizzleCymbal
		return
	case "sleigh bells":
		*enummetalvalue = EnumMetalValueSleighBells
		return
	case "suspended cymbal":
		*enummetalvalue = EnumMetalValueSuspendedCymbal
		return
	case "tam tam":
		*enummetalvalue = EnumMetalValueTamTam
		return
	case "tam tam with beater":
		*enummetalvalue = EnumMetalValueTamTamWithBeater
		return
	case "triangle":
		*enummetalvalue = EnumMetalValueTriangle
		return
	case "Vietnamese hat":
		*enummetalvalue = EnumMetalValueVietnameseHat
		return
	default:
		return errUnkownEnum
	}
}

func (enummetalvalue *EnumMetalValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumMetalValueAgogo":
		*enummetalvalue = EnumMetalValueAgogo
	case "EnumMetalValueAlmglocken":
		*enummetalvalue = EnumMetalValueAlmglocken
	case "EnumMetalValueBell":
		*enummetalvalue = EnumMetalValueBell
	case "EnumMetalValueBellPlate":
		*enummetalvalue = EnumMetalValueBellPlate
	case "EnumMetalValueBellTree":
		*enummetalvalue = EnumMetalValueBellTree
	case "EnumMetalValueBrakeDrum":
		*enummetalvalue = EnumMetalValueBrakeDrum
	case "EnumMetalValueCencerro":
		*enummetalvalue = EnumMetalValueCencerro
	case "EnumMetalValueChainRattle":
		*enummetalvalue = EnumMetalValueChainRattle
	case "EnumMetalValueChineseCymbal":
		*enummetalvalue = EnumMetalValueChineseCymbal
	case "EnumMetalValueCowbell":
		*enummetalvalue = EnumMetalValueCowbell
	case "EnumMetalValueCrashCymbals":
		*enummetalvalue = EnumMetalValueCrashCymbals
	case "EnumMetalValueCrotale":
		*enummetalvalue = EnumMetalValueCrotale
	case "EnumMetalValueCymbalTongs":
		*enummetalvalue = EnumMetalValueCymbalTongs
	case "EnumMetalValueDomedGong":
		*enummetalvalue = EnumMetalValueDomedGong
	case "EnumMetalValueFingerCymbals":
		*enummetalvalue = EnumMetalValueFingerCymbals
	case "EnumMetalValueFlexatone":
		*enummetalvalue = EnumMetalValueFlexatone
	case "EnumMetalValueGong":
		*enummetalvalue = EnumMetalValueGong
	case "EnumMetalValueHiHat":
		*enummetalvalue = EnumMetalValueHiHat
	case "EnumMetalValueHighHatCymbals":
		*enummetalvalue = EnumMetalValueHighHatCymbals
	case "EnumMetalValueHandbell":
		*enummetalvalue = EnumMetalValueHandbell
	case "EnumMetalValueJawHarp":
		*enummetalvalue = EnumMetalValueJawHarp
	case "EnumMetalValueJingleBells":
		*enummetalvalue = EnumMetalValueJingleBells
	case "EnumMetalValueMusicalSaw":
		*enummetalvalue = EnumMetalValueMusicalSaw
	case "EnumMetalValueShellBells":
		*enummetalvalue = EnumMetalValueShellBells
	case "EnumMetalValueSistrum":
		*enummetalvalue = EnumMetalValueSistrum
	case "EnumMetalValueSizzleCymbal":
		*enummetalvalue = EnumMetalValueSizzleCymbal
	case "EnumMetalValueSleighBells":
		*enummetalvalue = EnumMetalValueSleighBells
	case "EnumMetalValueSuspendedCymbal":
		*enummetalvalue = EnumMetalValueSuspendedCymbal
	case "EnumMetalValueTamTam":
		*enummetalvalue = EnumMetalValueTamTam
	case "EnumMetalValueTamTamWithBeater":
		*enummetalvalue = EnumMetalValueTamTamWithBeater
	case "EnumMetalValueTriangle":
		*enummetalvalue = EnumMetalValueTriangle
	case "EnumMetalValueVietnameseHat":
		*enummetalvalue = EnumMetalValueVietnameseHat
	default:
		err = errUnkownEnum
	}
	return
}

func (enummetalvalue *EnumMetalValue) ToCodeString() (res string) {

	switch *enummetalvalue {
	// insertion code per enum code
	case EnumMetalValueAgogo:
		res = "EnumMetalValueAgogo"
	case EnumMetalValueAlmglocken:
		res = "EnumMetalValueAlmglocken"
	case EnumMetalValueBell:
		res = "EnumMetalValueBell"
	case EnumMetalValueBellPlate:
		res = "EnumMetalValueBellPlate"
	case EnumMetalValueBellTree:
		res = "EnumMetalValueBellTree"
	case EnumMetalValueBrakeDrum:
		res = "EnumMetalValueBrakeDrum"
	case EnumMetalValueCencerro:
		res = "EnumMetalValueCencerro"
	case EnumMetalValueChainRattle:
		res = "EnumMetalValueChainRattle"
	case EnumMetalValueChineseCymbal:
		res = "EnumMetalValueChineseCymbal"
	case EnumMetalValueCowbell:
		res = "EnumMetalValueCowbell"
	case EnumMetalValueCrashCymbals:
		res = "EnumMetalValueCrashCymbals"
	case EnumMetalValueCrotale:
		res = "EnumMetalValueCrotale"
	case EnumMetalValueCymbalTongs:
		res = "EnumMetalValueCymbalTongs"
	case EnumMetalValueDomedGong:
		res = "EnumMetalValueDomedGong"
	case EnumMetalValueFingerCymbals:
		res = "EnumMetalValueFingerCymbals"
	case EnumMetalValueFlexatone:
		res = "EnumMetalValueFlexatone"
	case EnumMetalValueGong:
		res = "EnumMetalValueGong"
	case EnumMetalValueHiHat:
		res = "EnumMetalValueHiHat"
	case EnumMetalValueHighHatCymbals:
		res = "EnumMetalValueHighHatCymbals"
	case EnumMetalValueHandbell:
		res = "EnumMetalValueHandbell"
	case EnumMetalValueJawHarp:
		res = "EnumMetalValueJawHarp"
	case EnumMetalValueJingleBells:
		res = "EnumMetalValueJingleBells"
	case EnumMetalValueMusicalSaw:
		res = "EnumMetalValueMusicalSaw"
	case EnumMetalValueShellBells:
		res = "EnumMetalValueShellBells"
	case EnumMetalValueSistrum:
		res = "EnumMetalValueSistrum"
	case EnumMetalValueSizzleCymbal:
		res = "EnumMetalValueSizzleCymbal"
	case EnumMetalValueSleighBells:
		res = "EnumMetalValueSleighBells"
	case EnumMetalValueSuspendedCymbal:
		res = "EnumMetalValueSuspendedCymbal"
	case EnumMetalValueTamTam:
		res = "EnumMetalValueTamTam"
	case EnumMetalValueTamTamWithBeater:
		res = "EnumMetalValueTamTamWithBeater"
	case EnumMetalValueTriangle:
		res = "EnumMetalValueTriangle"
	case EnumMetalValueVietnameseHat:
		res = "EnumMetalValueVietnameseHat"
	}
	return
}

func (enummetalvalue EnumMetalValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumMetalValueAgogo")
	res = append(res, "EnumMetalValueAlmglocken")
	res = append(res, "EnumMetalValueBell")
	res = append(res, "EnumMetalValueBellPlate")
	res = append(res, "EnumMetalValueBellTree")
	res = append(res, "EnumMetalValueBrakeDrum")
	res = append(res, "EnumMetalValueCencerro")
	res = append(res, "EnumMetalValueChainRattle")
	res = append(res, "EnumMetalValueChineseCymbal")
	res = append(res, "EnumMetalValueCowbell")
	res = append(res, "EnumMetalValueCrashCymbals")
	res = append(res, "EnumMetalValueCrotale")
	res = append(res, "EnumMetalValueCymbalTongs")
	res = append(res, "EnumMetalValueDomedGong")
	res = append(res, "EnumMetalValueFingerCymbals")
	res = append(res, "EnumMetalValueFlexatone")
	res = append(res, "EnumMetalValueGong")
	res = append(res, "EnumMetalValueHiHat")
	res = append(res, "EnumMetalValueHighHatCymbals")
	res = append(res, "EnumMetalValueHandbell")
	res = append(res, "EnumMetalValueJawHarp")
	res = append(res, "EnumMetalValueJingleBells")
	res = append(res, "EnumMetalValueMusicalSaw")
	res = append(res, "EnumMetalValueShellBells")
	res = append(res, "EnumMetalValueSistrum")
	res = append(res, "EnumMetalValueSizzleCymbal")
	res = append(res, "EnumMetalValueSleighBells")
	res = append(res, "EnumMetalValueSuspendedCymbal")
	res = append(res, "EnumMetalValueTamTam")
	res = append(res, "EnumMetalValueTamTamWithBeater")
	res = append(res, "EnumMetalValueTriangle")
	res = append(res, "EnumMetalValueVietnameseHat")

	return
}

func (enummetalvalue EnumMetalValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "agogo")
	res = append(res, "almglocken")
	res = append(res, "bell")
	res = append(res, "bell plate")
	res = append(res, "bell tree")
	res = append(res, "brake drum")
	res = append(res, "cencerro")
	res = append(res, "chain rattle")
	res = append(res, "Chinese cymbal")
	res = append(res, "cowbell")
	res = append(res, "crash cymbals")
	res = append(res, "crotale")
	res = append(res, "cymbal tongs")
	res = append(res, "domed gong")
	res = append(res, "finger cymbals")
	res = append(res, "flexatone")
	res = append(res, "gong")
	res = append(res, "hi-hat")
	res = append(res, "high-hat cymbals")
	res = append(res, "handbell")
	res = append(res, "jaw harp")
	res = append(res, "jingle bells")
	res = append(res, "musical saw")
	res = append(res, "shell bells")
	res = append(res, "sistrum")
	res = append(res, "sizzle cymbal")
	res = append(res, "sleigh bells")
	res = append(res, "suspended cymbal")
	res = append(res, "tam tam")
	res = append(res, "tam tam with beater")
	res = append(res, "triangle")
	res = append(res, "Vietnamese hat")

	return
}

// Utility function for EnumNumeralMode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumnumeralmode EnumNumeralMode) ToString() (res string) {

	// migration of former implementation of enum
	switch enumnumeralmode {
	// insertion code per enum code
	case EnumNumeralModeMajor:
		res = "major"
	case EnumNumeralModeMinor:
		res = "minor"
	case EnumNumeralModeNaturalMinor:
		res = "natural minor"
	case EnumNumeralModeMelodicMinor:
		res = "melodic minor"
	case EnumNumeralModeHarmonicMinor:
		res = "harmonic minor"
	}
	return
}

func (enumnumeralmode *EnumNumeralMode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enumnumeralmode = EnumNumeralModeMajor
		return
	case "minor":
		*enumnumeralmode = EnumNumeralModeMinor
		return
	case "natural minor":
		*enumnumeralmode = EnumNumeralModeNaturalMinor
		return
	case "melodic minor":
		*enumnumeralmode = EnumNumeralModeMelodicMinor
		return
	case "harmonic minor":
		*enumnumeralmode = EnumNumeralModeHarmonicMinor
		return
	default:
		return errUnkownEnum
	}
}

func (enumnumeralmode *EnumNumeralMode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumNumeralModeMajor":
		*enumnumeralmode = EnumNumeralModeMajor
	case "EnumNumeralModeMinor":
		*enumnumeralmode = EnumNumeralModeMinor
	case "EnumNumeralModeNaturalMinor":
		*enumnumeralmode = EnumNumeralModeNaturalMinor
	case "EnumNumeralModeMelodicMinor":
		*enumnumeralmode = EnumNumeralModeMelodicMinor
	case "EnumNumeralModeHarmonicMinor":
		*enumnumeralmode = EnumNumeralModeHarmonicMinor
	default:
		err = errUnkownEnum
	}
	return
}

func (enumnumeralmode *EnumNumeralMode) ToCodeString() (res string) {

	switch *enumnumeralmode {
	// insertion code per enum code
	case EnumNumeralModeMajor:
		res = "EnumNumeralModeMajor"
	case EnumNumeralModeMinor:
		res = "EnumNumeralModeMinor"
	case EnumNumeralModeNaturalMinor:
		res = "EnumNumeralModeNaturalMinor"
	case EnumNumeralModeMelodicMinor:
		res = "EnumNumeralModeMelodicMinor"
	case EnumNumeralModeHarmonicMinor:
		res = "EnumNumeralModeHarmonicMinor"
	}
	return
}

func (enumnumeralmode EnumNumeralMode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumNumeralModeMajor")
	res = append(res, "EnumNumeralModeMinor")
	res = append(res, "EnumNumeralModeNaturalMinor")
	res = append(res, "EnumNumeralModeMelodicMinor")
	res = append(res, "EnumNumeralModeHarmonicMinor")

	return
}

func (enumnumeralmode EnumNumeralMode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "natural minor")
	res = append(res, "melodic minor")
	res = append(res, "harmonic minor")

	return
}

// Utility function for EnumOnOff
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumonoff EnumOnOff) ToString() (res string) {

	// migration of former implementation of enum
	switch enumonoff {
	// insertion code per enum code
	case EnumOnOffOn:
		res = "on"
	case EnumOnOffOff:
		res = "off"
	}
	return
}

func (enumonoff *EnumOnOff) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*enumonoff = EnumOnOffOn
		return
	case "off":
		*enumonoff = EnumOnOffOff
		return
	default:
		return errUnkownEnum
	}
}

func (enumonoff *EnumOnOff) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumOnOffOn":
		*enumonoff = EnumOnOffOn
	case "EnumOnOffOff":
		*enumonoff = EnumOnOffOff
	default:
		err = errUnkownEnum
	}
	return
}

func (enumonoff *EnumOnOff) ToCodeString() (res string) {

	switch *enumonoff {
	// insertion code per enum code
	case EnumOnOffOn:
		res = "EnumOnOffOn"
	case EnumOnOffOff:
		res = "EnumOnOffOff"
	}
	return
}

func (enumonoff EnumOnOff) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumOnOffOn")
	res = append(res, "EnumOnOffOff")

	return
}

func (enumonoff EnumOnOff) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "on")
	res = append(res, "off")

	return
}

// Utility function for EnumPedalType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumpedaltype EnumPedalType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumpedaltype {
	// insertion code per enum code
	case EnumPedalTypeStart:
		res = "start"
	case EnumPedalTypeStop:
		res = "stop"
	case EnumPedalTypeSostenuto:
		res = "sostenuto"
	case EnumPedalTypeChange:
		res = "change"
	case EnumPedalTypeContinue:
		res = "continue"
	case EnumPedalTypeDiscontinue:
		res = "discontinue"
	case EnumPedalTypeResume:
		res = "resume"
	}
	return
}

func (enumpedaltype *EnumPedalType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumpedaltype = EnumPedalTypeStart
		return
	case "stop":
		*enumpedaltype = EnumPedalTypeStop
		return
	case "sostenuto":
		*enumpedaltype = EnumPedalTypeSostenuto
		return
	case "change":
		*enumpedaltype = EnumPedalTypeChange
		return
	case "continue":
		*enumpedaltype = EnumPedalTypeContinue
		return
	case "discontinue":
		*enumpedaltype = EnumPedalTypeDiscontinue
		return
	case "resume":
		*enumpedaltype = EnumPedalTypeResume
		return
	default:
		return errUnkownEnum
	}
}

func (enumpedaltype *EnumPedalType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumPedalTypeStart":
		*enumpedaltype = EnumPedalTypeStart
	case "EnumPedalTypeStop":
		*enumpedaltype = EnumPedalTypeStop
	case "EnumPedalTypeSostenuto":
		*enumpedaltype = EnumPedalTypeSostenuto
	case "EnumPedalTypeChange":
		*enumpedaltype = EnumPedalTypeChange
	case "EnumPedalTypeContinue":
		*enumpedaltype = EnumPedalTypeContinue
	case "EnumPedalTypeDiscontinue":
		*enumpedaltype = EnumPedalTypeDiscontinue
	case "EnumPedalTypeResume":
		*enumpedaltype = EnumPedalTypeResume
	default:
		err = errUnkownEnum
	}
	return
}

func (enumpedaltype *EnumPedalType) ToCodeString() (res string) {

	switch *enumpedaltype {
	// insertion code per enum code
	case EnumPedalTypeStart:
		res = "EnumPedalTypeStart"
	case EnumPedalTypeStop:
		res = "EnumPedalTypeStop"
	case EnumPedalTypeSostenuto:
		res = "EnumPedalTypeSostenuto"
	case EnumPedalTypeChange:
		res = "EnumPedalTypeChange"
	case EnumPedalTypeContinue:
		res = "EnumPedalTypeContinue"
	case EnumPedalTypeDiscontinue:
		res = "EnumPedalTypeDiscontinue"
	case EnumPedalTypeResume:
		res = "EnumPedalTypeResume"
	}
	return
}

func (enumpedaltype EnumPedalType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumPedalTypeStart")
	res = append(res, "EnumPedalTypeStop")
	res = append(res, "EnumPedalTypeSostenuto")
	res = append(res, "EnumPedalTypeChange")
	res = append(res, "EnumPedalTypeContinue")
	res = append(res, "EnumPedalTypeDiscontinue")
	res = append(res, "EnumPedalTypeResume")

	return
}

func (enumpedaltype EnumPedalType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "sostenuto")
	res = append(res, "change")
	res = append(res, "continue")
	res = append(res, "discontinue")
	res = append(res, "resume")

	return
}

// Utility function for EnumPitchedValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumpitchedvalue EnumPitchedValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumpitchedvalue {
	// insertion code per enum code
	case EnumPitchedValueCelesta:
		res = "celesta"
	case EnumPitchedValueChimes:
		res = "chimes"
	case EnumPitchedValueGlockenspiel:
		res = "glockenspiel"
	case EnumPitchedValueLithophone:
		res = "lithophone"
	case EnumPitchedValueMallet:
		res = "mallet"
	case EnumPitchedValueMarimba:
		res = "marimba"
	case EnumPitchedValueSteelDrums:
		res = "steel drums"
	case EnumPitchedValueTubaphone:
		res = "tubaphone"
	case EnumPitchedValueTubularChimes:
		res = "tubular chimes"
	case EnumPitchedValueVibraphone:
		res = "vibraphone"
	case EnumPitchedValueXylophone:
		res = "xylophone"
	}
	return
}

func (enumpitchedvalue *EnumPitchedValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "celesta":
		*enumpitchedvalue = EnumPitchedValueCelesta
		return
	case "chimes":
		*enumpitchedvalue = EnumPitchedValueChimes
		return
	case "glockenspiel":
		*enumpitchedvalue = EnumPitchedValueGlockenspiel
		return
	case "lithophone":
		*enumpitchedvalue = EnumPitchedValueLithophone
		return
	case "mallet":
		*enumpitchedvalue = EnumPitchedValueMallet
		return
	case "marimba":
		*enumpitchedvalue = EnumPitchedValueMarimba
		return
	case "steel drums":
		*enumpitchedvalue = EnumPitchedValueSteelDrums
		return
	case "tubaphone":
		*enumpitchedvalue = EnumPitchedValueTubaphone
		return
	case "tubular chimes":
		*enumpitchedvalue = EnumPitchedValueTubularChimes
		return
	case "vibraphone":
		*enumpitchedvalue = EnumPitchedValueVibraphone
		return
	case "xylophone":
		*enumpitchedvalue = EnumPitchedValueXylophone
		return
	default:
		return errUnkownEnum
	}
}

func (enumpitchedvalue *EnumPitchedValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumPitchedValueCelesta":
		*enumpitchedvalue = EnumPitchedValueCelesta
	case "EnumPitchedValueChimes":
		*enumpitchedvalue = EnumPitchedValueChimes
	case "EnumPitchedValueGlockenspiel":
		*enumpitchedvalue = EnumPitchedValueGlockenspiel
	case "EnumPitchedValueLithophone":
		*enumpitchedvalue = EnumPitchedValueLithophone
	case "EnumPitchedValueMallet":
		*enumpitchedvalue = EnumPitchedValueMallet
	case "EnumPitchedValueMarimba":
		*enumpitchedvalue = EnumPitchedValueMarimba
	case "EnumPitchedValueSteelDrums":
		*enumpitchedvalue = EnumPitchedValueSteelDrums
	case "EnumPitchedValueTubaphone":
		*enumpitchedvalue = EnumPitchedValueTubaphone
	case "EnumPitchedValueTubularChimes":
		*enumpitchedvalue = EnumPitchedValueTubularChimes
	case "EnumPitchedValueVibraphone":
		*enumpitchedvalue = EnumPitchedValueVibraphone
	case "EnumPitchedValueXylophone":
		*enumpitchedvalue = EnumPitchedValueXylophone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumpitchedvalue *EnumPitchedValue) ToCodeString() (res string) {

	switch *enumpitchedvalue {
	// insertion code per enum code
	case EnumPitchedValueCelesta:
		res = "EnumPitchedValueCelesta"
	case EnumPitchedValueChimes:
		res = "EnumPitchedValueChimes"
	case EnumPitchedValueGlockenspiel:
		res = "EnumPitchedValueGlockenspiel"
	case EnumPitchedValueLithophone:
		res = "EnumPitchedValueLithophone"
	case EnumPitchedValueMallet:
		res = "EnumPitchedValueMallet"
	case EnumPitchedValueMarimba:
		res = "EnumPitchedValueMarimba"
	case EnumPitchedValueSteelDrums:
		res = "EnumPitchedValueSteelDrums"
	case EnumPitchedValueTubaphone:
		res = "EnumPitchedValueTubaphone"
	case EnumPitchedValueTubularChimes:
		res = "EnumPitchedValueTubularChimes"
	case EnumPitchedValueVibraphone:
		res = "EnumPitchedValueVibraphone"
	case EnumPitchedValueXylophone:
		res = "EnumPitchedValueXylophone"
	}
	return
}

func (enumpitchedvalue EnumPitchedValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumPitchedValueCelesta")
	res = append(res, "EnumPitchedValueChimes")
	res = append(res, "EnumPitchedValueGlockenspiel")
	res = append(res, "EnumPitchedValueLithophone")
	res = append(res, "EnumPitchedValueMallet")
	res = append(res, "EnumPitchedValueMarimba")
	res = append(res, "EnumPitchedValueSteelDrums")
	res = append(res, "EnumPitchedValueTubaphone")
	res = append(res, "EnumPitchedValueTubularChimes")
	res = append(res, "EnumPitchedValueVibraphone")
	res = append(res, "EnumPitchedValueXylophone")

	return
}

func (enumpitchedvalue EnumPitchedValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "celesta")
	res = append(res, "chimes")
	res = append(res, "glockenspiel")
	res = append(res, "lithophone")
	res = append(res, "mallet")
	res = append(res, "marimba")
	res = append(res, "steel drums")
	res = append(res, "tubaphone")
	res = append(res, "tubular chimes")
	res = append(res, "vibraphone")
	res = append(res, "xylophone")

	return
}

// Utility function for EnumPrincipalVoiceSymbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumprincipalvoicesymbol EnumPrincipalVoiceSymbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enumprincipalvoicesymbol {
	// insertion code per enum code
	case EnumPrincipalVoiceSymbolHauptstimme:
		res = "Hauptstimme"
	case EnumPrincipalVoiceSymbolNebenstimme:
		res = "Nebenstimme"
	case EnumPrincipalVoiceSymbolPlain:
		res = "plain"
	case EnumPrincipalVoiceSymbolNone:
		res = "none"
	}
	return
}

func (enumprincipalvoicesymbol *EnumPrincipalVoiceSymbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Hauptstimme":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolHauptstimme
		return
	case "Nebenstimme":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolNebenstimme
		return
	case "plain":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolPlain
		return
	case "none":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumprincipalvoicesymbol *EnumPrincipalVoiceSymbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumPrincipalVoiceSymbolHauptstimme":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolHauptstimme
	case "EnumPrincipalVoiceSymbolNebenstimme":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolNebenstimme
	case "EnumPrincipalVoiceSymbolPlain":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolPlain
	case "EnumPrincipalVoiceSymbolNone":
		*enumprincipalvoicesymbol = EnumPrincipalVoiceSymbolNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumprincipalvoicesymbol *EnumPrincipalVoiceSymbol) ToCodeString() (res string) {

	switch *enumprincipalvoicesymbol {
	// insertion code per enum code
	case EnumPrincipalVoiceSymbolHauptstimme:
		res = "EnumPrincipalVoiceSymbolHauptstimme"
	case EnumPrincipalVoiceSymbolNebenstimme:
		res = "EnumPrincipalVoiceSymbolNebenstimme"
	case EnumPrincipalVoiceSymbolPlain:
		res = "EnumPrincipalVoiceSymbolPlain"
	case EnumPrincipalVoiceSymbolNone:
		res = "EnumPrincipalVoiceSymbolNone"
	}
	return
}

func (enumprincipalvoicesymbol EnumPrincipalVoiceSymbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumPrincipalVoiceSymbolHauptstimme")
	res = append(res, "EnumPrincipalVoiceSymbolNebenstimme")
	res = append(res, "EnumPrincipalVoiceSymbolPlain")
	res = append(res, "EnumPrincipalVoiceSymbolNone")

	return
}

func (enumprincipalvoicesymbol EnumPrincipalVoiceSymbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Hauptstimme")
	res = append(res, "Nebenstimme")
	res = append(res, "plain")
	res = append(res, "none")

	return
}

// Utility function for EnumShowTuplet
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumshowtuplet EnumShowTuplet) ToString() (res string) {

	// migration of former implementation of enum
	switch enumshowtuplet {
	// insertion code per enum code
	case EnumShowTupletActual:
		res = "actual"
	case EnumShowTupletBoth:
		res = "both"
	case EnumShowTupletNone:
		res = "none"
	}
	return
}

func (enumshowtuplet *EnumShowTuplet) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "actual":
		*enumshowtuplet = EnumShowTupletActual
		return
	case "both":
		*enumshowtuplet = EnumShowTupletBoth
		return
	case "none":
		*enumshowtuplet = EnumShowTupletNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumshowtuplet *EnumShowTuplet) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumShowTupletActual":
		*enumshowtuplet = EnumShowTupletActual
	case "EnumShowTupletBoth":
		*enumshowtuplet = EnumShowTupletBoth
	case "EnumShowTupletNone":
		*enumshowtuplet = EnumShowTupletNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumshowtuplet *EnumShowTuplet) ToCodeString() (res string) {

	switch *enumshowtuplet {
	// insertion code per enum code
	case EnumShowTupletActual:
		res = "EnumShowTupletActual"
	case EnumShowTupletBoth:
		res = "EnumShowTupletBoth"
	case EnumShowTupletNone:
		res = "EnumShowTupletNone"
	}
	return
}

func (enumshowtuplet EnumShowTuplet) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumShowTupletActual")
	res = append(res, "EnumShowTupletBoth")
	res = append(res, "EnumShowTupletNone")

	return
}

func (enumshowtuplet EnumShowTuplet) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "actual")
	res = append(res, "both")
	res = append(res, "none")

	return
}

// Utility function for EnumStaffDivideSymbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstaffdividesymbol EnumStaffDivideSymbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstaffdividesymbol {
	// insertion code per enum code
	case EnumStaffDivideSymbolDown:
		res = "down"
	case EnumStaffDivideSymbolUp:
		res = "up"
	case EnumStaffDivideSymbolUpDown:
		res = "up-down"
	}
	return
}

func (enumstaffdividesymbol *EnumStaffDivideSymbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*enumstaffdividesymbol = EnumStaffDivideSymbolDown
		return
	case "up":
		*enumstaffdividesymbol = EnumStaffDivideSymbolUp
		return
	case "up-down":
		*enumstaffdividesymbol = EnumStaffDivideSymbolUpDown
		return
	default:
		return errUnkownEnum
	}
}

func (enumstaffdividesymbol *EnumStaffDivideSymbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStaffDivideSymbolDown":
		*enumstaffdividesymbol = EnumStaffDivideSymbolDown
	case "EnumStaffDivideSymbolUp":
		*enumstaffdividesymbol = EnumStaffDivideSymbolUp
	case "EnumStaffDivideSymbolUpDown":
		*enumstaffdividesymbol = EnumStaffDivideSymbolUpDown
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstaffdividesymbol *EnumStaffDivideSymbol) ToCodeString() (res string) {

	switch *enumstaffdividesymbol {
	// insertion code per enum code
	case EnumStaffDivideSymbolDown:
		res = "EnumStaffDivideSymbolDown"
	case EnumStaffDivideSymbolUp:
		res = "EnumStaffDivideSymbolUp"
	case EnumStaffDivideSymbolUpDown:
		res = "EnumStaffDivideSymbolUpDown"
	}
	return
}

func (enumstaffdividesymbol EnumStaffDivideSymbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStaffDivideSymbolDown")
	res = append(res, "EnumStaffDivideSymbolUp")
	res = append(res, "EnumStaffDivideSymbolUpDown")

	return
}

func (enumstaffdividesymbol EnumStaffDivideSymbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "up-down")

	return
}

// Utility function for EnumStartStopChangeContinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartstopchangecontinue EnumStartStopChangeContinue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartstopchangecontinue {
	// insertion code per enum code
	case EnumStartStopChangeContinueStart:
		res = "start"
	case EnumStartStopChangeContinueStop:
		res = "stop"
	case EnumStartStopChangeContinueChange:
		res = "change"
	case EnumStartStopChangeContinueContinue:
		res = "continue"
	}
	return
}

func (enumstartstopchangecontinue *EnumStartStopChangeContinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueStart
		return
	case "stop":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueStop
		return
	case "change":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueChange
		return
	case "continue":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueContinue
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartstopchangecontinue *EnumStartStopChangeContinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartStopChangeContinueStart":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueStart
	case "EnumStartStopChangeContinueStop":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueStop
	case "EnumStartStopChangeContinueChange":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueChange
	case "EnumStartStopChangeContinueContinue":
		*enumstartstopchangecontinue = EnumStartStopChangeContinueContinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartstopchangecontinue *EnumStartStopChangeContinue) ToCodeString() (res string) {

	switch *enumstartstopchangecontinue {
	// insertion code per enum code
	case EnumStartStopChangeContinueStart:
		res = "EnumStartStopChangeContinueStart"
	case EnumStartStopChangeContinueStop:
		res = "EnumStartStopChangeContinueStop"
	case EnumStartStopChangeContinueChange:
		res = "EnumStartStopChangeContinueChange"
	case EnumStartStopChangeContinueContinue:
		res = "EnumStartStopChangeContinueContinue"
	}
	return
}

func (enumstartstopchangecontinue EnumStartStopChangeContinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartStopChangeContinueStart")
	res = append(res, "EnumStartStopChangeContinueStop")
	res = append(res, "EnumStartStopChangeContinueChange")
	res = append(res, "EnumStartStopChangeContinueContinue")

	return
}

func (enumstartstopchangecontinue EnumStartStopChangeContinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "change")
	res = append(res, "continue")

	return
}

// Utility function for EnumStickLocation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsticklocation EnumStickLocation) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsticklocation {
	// insertion code per enum code
	case EnumStickLocationCenter:
		res = "center"
	case EnumStickLocationRim:
		res = "rim"
	case EnumStickLocationCymbalBell:
		res = "cymbal bell"
	case EnumStickLocationCymbalEdge:
		res = "cymbal edge"
	}
	return
}

func (enumsticklocation *EnumStickLocation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "center":
		*enumsticklocation = EnumStickLocationCenter
		return
	case "rim":
		*enumsticklocation = EnumStickLocationRim
		return
	case "cymbal bell":
		*enumsticklocation = EnumStickLocationCymbalBell
		return
	case "cymbal edge":
		*enumsticklocation = EnumStickLocationCymbalEdge
		return
	default:
		return errUnkownEnum
	}
}

func (enumsticklocation *EnumStickLocation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStickLocationCenter":
		*enumsticklocation = EnumStickLocationCenter
	case "EnumStickLocationRim":
		*enumsticklocation = EnumStickLocationRim
	case "EnumStickLocationCymbalBell":
		*enumsticklocation = EnumStickLocationCymbalBell
	case "EnumStickLocationCymbalEdge":
		*enumsticklocation = EnumStickLocationCymbalEdge
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsticklocation *EnumStickLocation) ToCodeString() (res string) {

	switch *enumsticklocation {
	// insertion code per enum code
	case EnumStickLocationCenter:
		res = "EnumStickLocationCenter"
	case EnumStickLocationRim:
		res = "EnumStickLocationRim"
	case EnumStickLocationCymbalBell:
		res = "EnumStickLocationCymbalBell"
	case EnumStickLocationCymbalEdge:
		res = "EnumStickLocationCymbalEdge"
	}
	return
}

func (enumsticklocation EnumStickLocation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStickLocationCenter")
	res = append(res, "EnumStickLocationRim")
	res = append(res, "EnumStickLocationCymbalBell")
	res = append(res, "EnumStickLocationCymbalEdge")

	return
}

func (enumsticklocation EnumStickLocation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "center")
	res = append(res, "rim")
	res = append(res, "cymbal bell")
	res = append(res, "cymbal edge")

	return
}

// Utility function for EnumStickMaterial
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstickmaterial EnumStickMaterial) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstickmaterial {
	// insertion code per enum code
	case EnumStickMaterialSoft:
		res = "soft"
	case EnumStickMaterialMedium:
		res = "medium"
	case EnumStickMaterialHard:
		res = "hard"
	case EnumStickMaterialShaded:
		res = "shaded"
	case EnumStickMaterialX:
		res = "x"
	}
	return
}

func (enumstickmaterial *EnumStickMaterial) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "soft":
		*enumstickmaterial = EnumStickMaterialSoft
		return
	case "medium":
		*enumstickmaterial = EnumStickMaterialMedium
		return
	case "hard":
		*enumstickmaterial = EnumStickMaterialHard
		return
	case "shaded":
		*enumstickmaterial = EnumStickMaterialShaded
		return
	case "x":
		*enumstickmaterial = EnumStickMaterialX
		return
	default:
		return errUnkownEnum
	}
}

func (enumstickmaterial *EnumStickMaterial) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStickMaterialSoft":
		*enumstickmaterial = EnumStickMaterialSoft
	case "EnumStickMaterialMedium":
		*enumstickmaterial = EnumStickMaterialMedium
	case "EnumStickMaterialHard":
		*enumstickmaterial = EnumStickMaterialHard
	case "EnumStickMaterialShaded":
		*enumstickmaterial = EnumStickMaterialShaded
	case "EnumStickMaterialX":
		*enumstickmaterial = EnumStickMaterialX
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstickmaterial *EnumStickMaterial) ToCodeString() (res string) {

	switch *enumstickmaterial {
	// insertion code per enum code
	case EnumStickMaterialSoft:
		res = "EnumStickMaterialSoft"
	case EnumStickMaterialMedium:
		res = "EnumStickMaterialMedium"
	case EnumStickMaterialHard:
		res = "EnumStickMaterialHard"
	case EnumStickMaterialShaded:
		res = "EnumStickMaterialShaded"
	case EnumStickMaterialX:
		res = "EnumStickMaterialX"
	}
	return
}

func (enumstickmaterial EnumStickMaterial) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStickMaterialSoft")
	res = append(res, "EnumStickMaterialMedium")
	res = append(res, "EnumStickMaterialHard")
	res = append(res, "EnumStickMaterialShaded")
	res = append(res, "EnumStickMaterialX")

	return
}

func (enumstickmaterial EnumStickMaterial) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "soft")
	res = append(res, "medium")
	res = append(res, "hard")
	res = append(res, "shaded")
	res = append(res, "x")

	return
}

// Utility function for EnumStickType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsticktype EnumStickType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsticktype {
	// insertion code per enum code
	case EnumStickTypeBassDrum:
		res = "bass drum"
	case EnumStickTypeDoubleBassDrum:
		res = "double bass drum"
	case EnumStickTypeGlockenspiel:
		res = "glockenspiel"
	case EnumStickTypeGum:
		res = "gum"
	case EnumStickTypeHammer:
		res = "hammer"
	case EnumStickTypeSuperball:
		res = "superball"
	case EnumStickTypeTimpani:
		res = "timpani"
	case EnumStickTypeWound:
		res = "wound"
	case EnumStickTypeXylophone:
		res = "xylophone"
	case EnumStickTypeYarn:
		res = "yarn"
	}
	return
}

func (enumsticktype *EnumStickType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*enumsticktype = EnumStickTypeBassDrum
		return
	case "double bass drum":
		*enumsticktype = EnumStickTypeDoubleBassDrum
		return
	case "glockenspiel":
		*enumsticktype = EnumStickTypeGlockenspiel
		return
	case "gum":
		*enumsticktype = EnumStickTypeGum
		return
	case "hammer":
		*enumsticktype = EnumStickTypeHammer
		return
	case "superball":
		*enumsticktype = EnumStickTypeSuperball
		return
	case "timpani":
		*enumsticktype = EnumStickTypeTimpani
		return
	case "wound":
		*enumsticktype = EnumStickTypeWound
		return
	case "xylophone":
		*enumsticktype = EnumStickTypeXylophone
		return
	case "yarn":
		*enumsticktype = EnumStickTypeYarn
		return
	default:
		return errUnkownEnum
	}
}

func (enumsticktype *EnumStickType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStickTypeBassDrum":
		*enumsticktype = EnumStickTypeBassDrum
	case "EnumStickTypeDoubleBassDrum":
		*enumsticktype = EnumStickTypeDoubleBassDrum
	case "EnumStickTypeGlockenspiel":
		*enumsticktype = EnumStickTypeGlockenspiel
	case "EnumStickTypeGum":
		*enumsticktype = EnumStickTypeGum
	case "EnumStickTypeHammer":
		*enumsticktype = EnumStickTypeHammer
	case "EnumStickTypeSuperball":
		*enumsticktype = EnumStickTypeSuperball
	case "EnumStickTypeTimpani":
		*enumsticktype = EnumStickTypeTimpani
	case "EnumStickTypeWound":
		*enumsticktype = EnumStickTypeWound
	case "EnumStickTypeXylophone":
		*enumsticktype = EnumStickTypeXylophone
	case "EnumStickTypeYarn":
		*enumsticktype = EnumStickTypeYarn
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsticktype *EnumStickType) ToCodeString() (res string) {

	switch *enumsticktype {
	// insertion code per enum code
	case EnumStickTypeBassDrum:
		res = "EnumStickTypeBassDrum"
	case EnumStickTypeDoubleBassDrum:
		res = "EnumStickTypeDoubleBassDrum"
	case EnumStickTypeGlockenspiel:
		res = "EnumStickTypeGlockenspiel"
	case EnumStickTypeGum:
		res = "EnumStickTypeGum"
	case EnumStickTypeHammer:
		res = "EnumStickTypeHammer"
	case EnumStickTypeSuperball:
		res = "EnumStickTypeSuperball"
	case EnumStickTypeTimpani:
		res = "EnumStickTypeTimpani"
	case EnumStickTypeWound:
		res = "EnumStickTypeWound"
	case EnumStickTypeXylophone:
		res = "EnumStickTypeXylophone"
	case EnumStickTypeYarn:
		res = "EnumStickTypeYarn"
	}
	return
}

func (enumsticktype EnumStickType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStickTypeBassDrum")
	res = append(res, "EnumStickTypeDoubleBassDrum")
	res = append(res, "EnumStickTypeGlockenspiel")
	res = append(res, "EnumStickTypeGum")
	res = append(res, "EnumStickTypeHammer")
	res = append(res, "EnumStickTypeSuperball")
	res = append(res, "EnumStickTypeTimpani")
	res = append(res, "EnumStickTypeWound")
	res = append(res, "EnumStickTypeXylophone")
	res = append(res, "EnumStickTypeYarn")

	return
}

func (enumsticktype EnumStickType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bass drum")
	res = append(res, "double bass drum")
	res = append(res, "glockenspiel")
	res = append(res, "gum")
	res = append(res, "hammer")
	res = append(res, "superball")
	res = append(res, "timpani")
	res = append(res, "wound")
	res = append(res, "xylophone")
	res = append(res, "yarn")

	return
}

// Utility function for EnumSyncType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsynctype EnumSyncType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsynctype {
	// insertion code per enum code
	case EnumSyncTypeNone:
		res = "none"
	case EnumSyncTypeTempo:
		res = "tempo"
	case EnumSyncTypeMostlyTempo:
		res = "mostly-tempo"
	case EnumSyncTypeMostlyEvent:
		res = "mostly-event"
	case EnumSyncTypeEvent:
		res = "event"
	case EnumSyncTypeAlwaysEvent:
		res = "always-event"
	}
	return
}

func (enumsynctype *EnumSyncType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enumsynctype = EnumSyncTypeNone
		return
	case "tempo":
		*enumsynctype = EnumSyncTypeTempo
		return
	case "mostly-tempo":
		*enumsynctype = EnumSyncTypeMostlyTempo
		return
	case "mostly-event":
		*enumsynctype = EnumSyncTypeMostlyEvent
		return
	case "event":
		*enumsynctype = EnumSyncTypeEvent
		return
	case "always-event":
		*enumsynctype = EnumSyncTypeAlwaysEvent
		return
	default:
		return errUnkownEnum
	}
}

func (enumsynctype *EnumSyncType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSyncTypeNone":
		*enumsynctype = EnumSyncTypeNone
	case "EnumSyncTypeTempo":
		*enumsynctype = EnumSyncTypeTempo
	case "EnumSyncTypeMostlyTempo":
		*enumsynctype = EnumSyncTypeMostlyTempo
	case "EnumSyncTypeMostlyEvent":
		*enumsynctype = EnumSyncTypeMostlyEvent
	case "EnumSyncTypeEvent":
		*enumsynctype = EnumSyncTypeEvent
	case "EnumSyncTypeAlwaysEvent":
		*enumsynctype = EnumSyncTypeAlwaysEvent
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsynctype *EnumSyncType) ToCodeString() (res string) {

	switch *enumsynctype {
	// insertion code per enum code
	case EnumSyncTypeNone:
		res = "EnumSyncTypeNone"
	case EnumSyncTypeTempo:
		res = "EnumSyncTypeTempo"
	case EnumSyncTypeMostlyTempo:
		res = "EnumSyncTypeMostlyTempo"
	case EnumSyncTypeMostlyEvent:
		res = "EnumSyncTypeMostlyEvent"
	case EnumSyncTypeEvent:
		res = "EnumSyncTypeEvent"
	case EnumSyncTypeAlwaysEvent:
		res = "EnumSyncTypeAlwaysEvent"
	}
	return
}

func (enumsynctype EnumSyncType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSyncTypeNone")
	res = append(res, "EnumSyncTypeTempo")
	res = append(res, "EnumSyncTypeMostlyTempo")
	res = append(res, "EnumSyncTypeMostlyEvent")
	res = append(res, "EnumSyncTypeEvent")
	res = append(res, "EnumSyncTypeAlwaysEvent")

	return
}

func (enumsynctype EnumSyncType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "tempo")
	res = append(res, "mostly-tempo")
	res = append(res, "mostly-event")
	res = append(res, "event")
	res = append(res, "always-event")

	return
}

// Utility function for EnumTipDirection
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtipdirection EnumTipDirection) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtipdirection {
	// insertion code per enum code
	case EnumTipDirectionUp:
		res = "up"
	case EnumTipDirectionDown:
		res = "down"
	case EnumTipDirectionLeft:
		res = "left"
	case EnumTipDirectionRight:
		res = "right"
	case EnumTipDirectionNorthwest:
		res = "northwest"
	case EnumTipDirectionNortheast:
		res = "northeast"
	case EnumTipDirectionSoutheast:
		res = "southeast"
	case EnumTipDirectionSouthwest:
		res = "southwest"
	}
	return
}

func (enumtipdirection *EnumTipDirection) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enumtipdirection = EnumTipDirectionUp
		return
	case "down":
		*enumtipdirection = EnumTipDirectionDown
		return
	case "left":
		*enumtipdirection = EnumTipDirectionLeft
		return
	case "right":
		*enumtipdirection = EnumTipDirectionRight
		return
	case "northwest":
		*enumtipdirection = EnumTipDirectionNorthwest
		return
	case "northeast":
		*enumtipdirection = EnumTipDirectionNortheast
		return
	case "southeast":
		*enumtipdirection = EnumTipDirectionSoutheast
		return
	case "southwest":
		*enumtipdirection = EnumTipDirectionSouthwest
		return
	default:
		return errUnkownEnum
	}
}

func (enumtipdirection *EnumTipDirection) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTipDirectionUp":
		*enumtipdirection = EnumTipDirectionUp
	case "EnumTipDirectionDown":
		*enumtipdirection = EnumTipDirectionDown
	case "EnumTipDirectionLeft":
		*enumtipdirection = EnumTipDirectionLeft
	case "EnumTipDirectionRight":
		*enumtipdirection = EnumTipDirectionRight
	case "EnumTipDirectionNorthwest":
		*enumtipdirection = EnumTipDirectionNorthwest
	case "EnumTipDirectionNortheast":
		*enumtipdirection = EnumTipDirectionNortheast
	case "EnumTipDirectionSoutheast":
		*enumtipdirection = EnumTipDirectionSoutheast
	case "EnumTipDirectionSouthwest":
		*enumtipdirection = EnumTipDirectionSouthwest
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtipdirection *EnumTipDirection) ToCodeString() (res string) {

	switch *enumtipdirection {
	// insertion code per enum code
	case EnumTipDirectionUp:
		res = "EnumTipDirectionUp"
	case EnumTipDirectionDown:
		res = "EnumTipDirectionDown"
	case EnumTipDirectionLeft:
		res = "EnumTipDirectionLeft"
	case EnumTipDirectionRight:
		res = "EnumTipDirectionRight"
	case EnumTipDirectionNorthwest:
		res = "EnumTipDirectionNorthwest"
	case EnumTipDirectionNortheast:
		res = "EnumTipDirectionNortheast"
	case EnumTipDirectionSoutheast:
		res = "EnumTipDirectionSoutheast"
	case EnumTipDirectionSouthwest:
		res = "EnumTipDirectionSouthwest"
	}
	return
}

func (enumtipdirection EnumTipDirection) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTipDirectionUp")
	res = append(res, "EnumTipDirectionDown")
	res = append(res, "EnumTipDirectionLeft")
	res = append(res, "EnumTipDirectionRight")
	res = append(res, "EnumTipDirectionNorthwest")
	res = append(res, "EnumTipDirectionNortheast")
	res = append(res, "EnumTipDirectionSoutheast")
	res = append(res, "EnumTipDirectionSouthwest")

	return
}

func (enumtipdirection EnumTipDirection) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "northwest")
	res = append(res, "northeast")
	res = append(res, "southeast")
	res = append(res, "southwest")

	return
}

// Utility function for EnumUpDownStopContinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumupdownstopcontinue EnumUpDownStopContinue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumupdownstopcontinue {
	// insertion code per enum code
	case EnumUpDownStopContinueUp:
		res = "up"
	case EnumUpDownStopContinueDown:
		res = "down"
	case EnumUpDownStopContinueStop:
		res = "stop"
	case EnumUpDownStopContinueContinue:
		res = "continue"
	}
	return
}

func (enumupdownstopcontinue *EnumUpDownStopContinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enumupdownstopcontinue = EnumUpDownStopContinueUp
		return
	case "down":
		*enumupdownstopcontinue = EnumUpDownStopContinueDown
		return
	case "stop":
		*enumupdownstopcontinue = EnumUpDownStopContinueStop
		return
	case "continue":
		*enumupdownstopcontinue = EnumUpDownStopContinueContinue
		return
	default:
		return errUnkownEnum
	}
}

func (enumupdownstopcontinue *EnumUpDownStopContinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumUpDownStopContinueUp":
		*enumupdownstopcontinue = EnumUpDownStopContinueUp
	case "EnumUpDownStopContinueDown":
		*enumupdownstopcontinue = EnumUpDownStopContinueDown
	case "EnumUpDownStopContinueStop":
		*enumupdownstopcontinue = EnumUpDownStopContinueStop
	case "EnumUpDownStopContinueContinue":
		*enumupdownstopcontinue = EnumUpDownStopContinueContinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enumupdownstopcontinue *EnumUpDownStopContinue) ToCodeString() (res string) {

	switch *enumupdownstopcontinue {
	// insertion code per enum code
	case EnumUpDownStopContinueUp:
		res = "EnumUpDownStopContinueUp"
	case EnumUpDownStopContinueDown:
		res = "EnumUpDownStopContinueDown"
	case EnumUpDownStopContinueStop:
		res = "EnumUpDownStopContinueStop"
	case EnumUpDownStopContinueContinue:
		res = "EnumUpDownStopContinueContinue"
	}
	return
}

func (enumupdownstopcontinue EnumUpDownStopContinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumUpDownStopContinueUp")
	res = append(res, "EnumUpDownStopContinueDown")
	res = append(res, "EnumUpDownStopContinueStop")
	res = append(res, "EnumUpDownStopContinueContinue")

	return
}

func (enumupdownstopcontinue EnumUpDownStopContinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for EnumWedgeType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumwedgetype EnumWedgeType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumwedgetype {
	// insertion code per enum code
	case EnumWedgeTypeCrescendo:
		res = "crescendo"
	case EnumWedgeTypeDiminuendo:
		res = "diminuendo"
	case EnumWedgeTypeStop:
		res = "stop"
	case EnumWedgeTypeContinue:
		res = "continue"
	}
	return
}

func (enumwedgetype *EnumWedgeType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "crescendo":
		*enumwedgetype = EnumWedgeTypeCrescendo
		return
	case "diminuendo":
		*enumwedgetype = EnumWedgeTypeDiminuendo
		return
	case "stop":
		*enumwedgetype = EnumWedgeTypeStop
		return
	case "continue":
		*enumwedgetype = EnumWedgeTypeContinue
		return
	default:
		return errUnkownEnum
	}
}

func (enumwedgetype *EnumWedgeType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumWedgeTypeCrescendo":
		*enumwedgetype = EnumWedgeTypeCrescendo
	case "EnumWedgeTypeDiminuendo":
		*enumwedgetype = EnumWedgeTypeDiminuendo
	case "EnumWedgeTypeStop":
		*enumwedgetype = EnumWedgeTypeStop
	case "EnumWedgeTypeContinue":
		*enumwedgetype = EnumWedgeTypeContinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enumwedgetype *EnumWedgeType) ToCodeString() (res string) {

	switch *enumwedgetype {
	// insertion code per enum code
	case EnumWedgeTypeCrescendo:
		res = "EnumWedgeTypeCrescendo"
	case EnumWedgeTypeDiminuendo:
		res = "EnumWedgeTypeDiminuendo"
	case EnumWedgeTypeStop:
		res = "EnumWedgeTypeStop"
	case EnumWedgeTypeContinue:
		res = "EnumWedgeTypeContinue"
	}
	return
}

func (enumwedgetype EnumWedgeType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumWedgeTypeCrescendo")
	res = append(res, "EnumWedgeTypeDiminuendo")
	res = append(res, "EnumWedgeTypeStop")
	res = append(res, "EnumWedgeTypeContinue")

	return
}

func (enumwedgetype EnumWedgeType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "crescendo")
	res = append(res, "diminuendo")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for EnumWoodValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumwoodvalue EnumWoodValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumwoodvalue {
	// insertion code per enum code
	case EnumWoodValueBambooScraper:
		res = "bamboo scraper"
	case EnumWoodValueBoardClapper:
		res = "board clapper"
	case EnumWoodValueCabasa:
		res = "cabasa"
	case EnumWoodValueCastanets:
		res = "castanets"
	case EnumWoodValueCastanetsWithHandle:
		res = "castanets with handle"
	case EnumWoodValueClaves:
		res = "claves"
	case EnumWoodValueFootballRattle:
		res = "football rattle"
	case EnumWoodValueGuiro:
		res = "guiro"
	case EnumWoodValueLogDrum:
		res = "log drum"
	case EnumWoodValueMaraca:
		res = "maraca"
	case EnumWoodValueMaracas:
		res = "maracas"
	case EnumWoodValueQuijada:
		res = "quijada"
	case EnumWoodValueRainstick:
		res = "rainstick"
	case EnumWoodValueRatchet:
		res = "ratchet"
	case EnumWoodValueRecoReco:
		res = "reco-reco"
	case EnumWoodValueSandpaperBlocks:
		res = "sandpaper blocks"
	case EnumWoodValueSlitDrum:
		res = "slit drum"
	case EnumWoodValueTempleBlock:
		res = "temple block"
	case EnumWoodValueVibraslap:
		res = "vibraslap"
	case EnumWoodValueWhip:
		res = "whip"
	case EnumWoodValueWoodBlock:
		res = "wood block"
	}
	return
}

func (enumwoodvalue *EnumWoodValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bamboo scraper":
		*enumwoodvalue = EnumWoodValueBambooScraper
		return
	case "board clapper":
		*enumwoodvalue = EnumWoodValueBoardClapper
		return
	case "cabasa":
		*enumwoodvalue = EnumWoodValueCabasa
		return
	case "castanets":
		*enumwoodvalue = EnumWoodValueCastanets
		return
	case "castanets with handle":
		*enumwoodvalue = EnumWoodValueCastanetsWithHandle
		return
	case "claves":
		*enumwoodvalue = EnumWoodValueClaves
		return
	case "football rattle":
		*enumwoodvalue = EnumWoodValueFootballRattle
		return
	case "guiro":
		*enumwoodvalue = EnumWoodValueGuiro
		return
	case "log drum":
		*enumwoodvalue = EnumWoodValueLogDrum
		return
	case "maraca":
		*enumwoodvalue = EnumWoodValueMaraca
		return
	case "maracas":
		*enumwoodvalue = EnumWoodValueMaracas
		return
	case "quijada":
		*enumwoodvalue = EnumWoodValueQuijada
		return
	case "rainstick":
		*enumwoodvalue = EnumWoodValueRainstick
		return
	case "ratchet":
		*enumwoodvalue = EnumWoodValueRatchet
		return
	case "reco-reco":
		*enumwoodvalue = EnumWoodValueRecoReco
		return
	case "sandpaper blocks":
		*enumwoodvalue = EnumWoodValueSandpaperBlocks
		return
	case "slit drum":
		*enumwoodvalue = EnumWoodValueSlitDrum
		return
	case "temple block":
		*enumwoodvalue = EnumWoodValueTempleBlock
		return
	case "vibraslap":
		*enumwoodvalue = EnumWoodValueVibraslap
		return
	case "whip":
		*enumwoodvalue = EnumWoodValueWhip
		return
	case "wood block":
		*enumwoodvalue = EnumWoodValueWoodBlock
		return
	default:
		return errUnkownEnum
	}
}

func (enumwoodvalue *EnumWoodValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumWoodValueBambooScraper":
		*enumwoodvalue = EnumWoodValueBambooScraper
	case "EnumWoodValueBoardClapper":
		*enumwoodvalue = EnumWoodValueBoardClapper
	case "EnumWoodValueCabasa":
		*enumwoodvalue = EnumWoodValueCabasa
	case "EnumWoodValueCastanets":
		*enumwoodvalue = EnumWoodValueCastanets
	case "EnumWoodValueCastanetsWithHandle":
		*enumwoodvalue = EnumWoodValueCastanetsWithHandle
	case "EnumWoodValueClaves":
		*enumwoodvalue = EnumWoodValueClaves
	case "EnumWoodValueFootballRattle":
		*enumwoodvalue = EnumWoodValueFootballRattle
	case "EnumWoodValueGuiro":
		*enumwoodvalue = EnumWoodValueGuiro
	case "EnumWoodValueLogDrum":
		*enumwoodvalue = EnumWoodValueLogDrum
	case "EnumWoodValueMaraca":
		*enumwoodvalue = EnumWoodValueMaraca
	case "EnumWoodValueMaracas":
		*enumwoodvalue = EnumWoodValueMaracas
	case "EnumWoodValueQuijada":
		*enumwoodvalue = EnumWoodValueQuijada
	case "EnumWoodValueRainstick":
		*enumwoodvalue = EnumWoodValueRainstick
	case "EnumWoodValueRatchet":
		*enumwoodvalue = EnumWoodValueRatchet
	case "EnumWoodValueRecoReco":
		*enumwoodvalue = EnumWoodValueRecoReco
	case "EnumWoodValueSandpaperBlocks":
		*enumwoodvalue = EnumWoodValueSandpaperBlocks
	case "EnumWoodValueSlitDrum":
		*enumwoodvalue = EnumWoodValueSlitDrum
	case "EnumWoodValueTempleBlock":
		*enumwoodvalue = EnumWoodValueTempleBlock
	case "EnumWoodValueVibraslap":
		*enumwoodvalue = EnumWoodValueVibraslap
	case "EnumWoodValueWhip":
		*enumwoodvalue = EnumWoodValueWhip
	case "EnumWoodValueWoodBlock":
		*enumwoodvalue = EnumWoodValueWoodBlock
	default:
		err = errUnkownEnum
	}
	return
}

func (enumwoodvalue *EnumWoodValue) ToCodeString() (res string) {

	switch *enumwoodvalue {
	// insertion code per enum code
	case EnumWoodValueBambooScraper:
		res = "EnumWoodValueBambooScraper"
	case EnumWoodValueBoardClapper:
		res = "EnumWoodValueBoardClapper"
	case EnumWoodValueCabasa:
		res = "EnumWoodValueCabasa"
	case EnumWoodValueCastanets:
		res = "EnumWoodValueCastanets"
	case EnumWoodValueCastanetsWithHandle:
		res = "EnumWoodValueCastanetsWithHandle"
	case EnumWoodValueClaves:
		res = "EnumWoodValueClaves"
	case EnumWoodValueFootballRattle:
		res = "EnumWoodValueFootballRattle"
	case EnumWoodValueGuiro:
		res = "EnumWoodValueGuiro"
	case EnumWoodValueLogDrum:
		res = "EnumWoodValueLogDrum"
	case EnumWoodValueMaraca:
		res = "EnumWoodValueMaraca"
	case EnumWoodValueMaracas:
		res = "EnumWoodValueMaracas"
	case EnumWoodValueQuijada:
		res = "EnumWoodValueQuijada"
	case EnumWoodValueRainstick:
		res = "EnumWoodValueRainstick"
	case EnumWoodValueRatchet:
		res = "EnumWoodValueRatchet"
	case EnumWoodValueRecoReco:
		res = "EnumWoodValueRecoReco"
	case EnumWoodValueSandpaperBlocks:
		res = "EnumWoodValueSandpaperBlocks"
	case EnumWoodValueSlitDrum:
		res = "EnumWoodValueSlitDrum"
	case EnumWoodValueTempleBlock:
		res = "EnumWoodValueTempleBlock"
	case EnumWoodValueVibraslap:
		res = "EnumWoodValueVibraslap"
	case EnumWoodValueWhip:
		res = "EnumWoodValueWhip"
	case EnumWoodValueWoodBlock:
		res = "EnumWoodValueWoodBlock"
	}
	return
}

func (enumwoodvalue EnumWoodValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumWoodValueBambooScraper")
	res = append(res, "EnumWoodValueBoardClapper")
	res = append(res, "EnumWoodValueCabasa")
	res = append(res, "EnumWoodValueCastanets")
	res = append(res, "EnumWoodValueCastanetsWithHandle")
	res = append(res, "EnumWoodValueClaves")
	res = append(res, "EnumWoodValueFootballRattle")
	res = append(res, "EnumWoodValueGuiro")
	res = append(res, "EnumWoodValueLogDrum")
	res = append(res, "EnumWoodValueMaraca")
	res = append(res, "EnumWoodValueMaracas")
	res = append(res, "EnumWoodValueQuijada")
	res = append(res, "EnumWoodValueRainstick")
	res = append(res, "EnumWoodValueRatchet")
	res = append(res, "EnumWoodValueRecoReco")
	res = append(res, "EnumWoodValueSandpaperBlocks")
	res = append(res, "EnumWoodValueSlitDrum")
	res = append(res, "EnumWoodValueTempleBlock")
	res = append(res, "EnumWoodValueVibraslap")
	res = append(res, "EnumWoodValueWhip")
	res = append(res, "EnumWoodValueWoodBlock")

	return
}

func (enumwoodvalue EnumWoodValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bamboo scraper")
	res = append(res, "board clapper")
	res = append(res, "cabasa")
	res = append(res, "castanets")
	res = append(res, "castanets with handle")
	res = append(res, "claves")
	res = append(res, "football rattle")
	res = append(res, "guiro")
	res = append(res, "log drum")
	res = append(res, "maraca")
	res = append(res, "maracas")
	res = append(res, "quijada")
	res = append(res, "rainstick")
	res = append(res, "ratchet")
	res = append(res, "reco-reco")
	res = append(res, "sandpaper blocks")
	res = append(res, "slit drum")
	res = append(res, "temple block")
	res = append(res, "vibraslap")
	res = append(res, "whip")
	res = append(res, "wood block")

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
