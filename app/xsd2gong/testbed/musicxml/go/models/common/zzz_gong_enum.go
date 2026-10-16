// generated code - do not edit
package common

// insertion point of enum utility functions
// Utility function for EnumAboveBelow
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumabovebelow EnumAboveBelow) ToString() (res string) {

	// migration of former implementation of enum
	switch enumabovebelow {
	// insertion code per enum code
	case EnumAboveBelowAbove:
		res = "above"
	case EnumAboveBelowBelow:
		res = "below"
	}
	return
}

func (enumabovebelow *EnumAboveBelow) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "above":
		*enumabovebelow = EnumAboveBelowAbove
		return
	case "below":
		*enumabovebelow = EnumAboveBelowBelow
		return
	default:
		return errUnkownEnum
	}
}

func (enumabovebelow *EnumAboveBelow) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumAboveBelowAbove":
		*enumabovebelow = EnumAboveBelowAbove
	case "EnumAboveBelowBelow":
		*enumabovebelow = EnumAboveBelowBelow
	default:
		err = errUnkownEnum
	}
	return
}

func (enumabovebelow *EnumAboveBelow) ToCodeString() (res string) {

	switch *enumabovebelow {
	// insertion code per enum code
	case EnumAboveBelowAbove:
		res = "EnumAboveBelowAbove"
	case EnumAboveBelowBelow:
		res = "EnumAboveBelowBelow"
	}
	return
}

func (enumabovebelow EnumAboveBelow) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumAboveBelowAbove")
	res = append(res, "EnumAboveBelowBelow")

	return
}

func (enumabovebelow EnumAboveBelow) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "above")
	res = append(res, "below")

	return
}

// Utility function for EnumAccidentalValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumaccidentalvalue EnumAccidentalValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumaccidentalvalue {
	// insertion code per enum code
	case EnumAccidentalValueSharp:
		res = "sharp"
	case EnumAccidentalValueNatural:
		res = "natural"
	case EnumAccidentalValueFlat:
		res = "flat"
	case EnumAccidentalValueDoubleSharp:
		res = "double-sharp"
	case EnumAccidentalValueSharpSharp:
		res = "sharp-sharp"
	case EnumAccidentalValueFlatFlat:
		res = "flat-flat"
	case EnumAccidentalValueNaturalSharp:
		res = "natural-sharp"
	case EnumAccidentalValueNaturalFlat:
		res = "natural-flat"
	case EnumAccidentalValueQuarterFlat:
		res = "quarter-flat"
	case EnumAccidentalValueQuarterSharp:
		res = "quarter-sharp"
	case EnumAccidentalValueThreeQuartersFlat:
		res = "three-quarters-flat"
	case EnumAccidentalValueThreeQuartersSharp:
		res = "three-quarters-sharp"
	case EnumAccidentalValueSharpDown:
		res = "sharp-down"
	case EnumAccidentalValueSharpUp:
		res = "sharp-up"
	case EnumAccidentalValueNaturalDown:
		res = "natural-down"
	case EnumAccidentalValueNaturalUp:
		res = "natural-up"
	case EnumAccidentalValueFlatDown:
		res = "flat-down"
	case EnumAccidentalValueFlatUp:
		res = "flat-up"
	case EnumAccidentalValueDoubleSharpDown:
		res = "double-sharp-down"
	case EnumAccidentalValueDoubleSharpUp:
		res = "double-sharp-up"
	case EnumAccidentalValueFlatFlatDown:
		res = "flat-flat-down"
	case EnumAccidentalValueFlatFlatUp:
		res = "flat-flat-up"
	case EnumAccidentalValueArrowDown:
		res = "arrow-down"
	case EnumAccidentalValueArrowUp:
		res = "arrow-up"
	case EnumAccidentalValueTripleSharp:
		res = "triple-sharp"
	case EnumAccidentalValueTripleFlat:
		res = "triple-flat"
	case EnumAccidentalValueSlashQuarterSharp:
		res = "slash-quarter-sharp"
	case EnumAccidentalValueSlashSharp:
		res = "slash-sharp"
	case EnumAccidentalValueSlashFlat:
		res = "slash-flat"
	case EnumAccidentalValueDoubleSlashFlat:
		res = "double-slash-flat"
	case EnumAccidentalValueSharp1:
		res = "sharp-1"
	case EnumAccidentalValueSharp2:
		res = "sharp-2"
	case EnumAccidentalValueSharp3:
		res = "sharp-3"
	case EnumAccidentalValueSharp5:
		res = "sharp-5"
	case EnumAccidentalValueFlat1:
		res = "flat-1"
	case EnumAccidentalValueFlat2:
		res = "flat-2"
	case EnumAccidentalValueFlat3:
		res = "flat-3"
	case EnumAccidentalValueFlat4:
		res = "flat-4"
	case EnumAccidentalValueSori:
		res = "sori"
	case EnumAccidentalValueKoron:
		res = "koron"
	case EnumAccidentalValueOther:
		res = "other"
	}
	return
}

func (enumaccidentalvalue *EnumAccidentalValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sharp":
		*enumaccidentalvalue = EnumAccidentalValueSharp
		return
	case "natural":
		*enumaccidentalvalue = EnumAccidentalValueNatural
		return
	case "flat":
		*enumaccidentalvalue = EnumAccidentalValueFlat
		return
	case "double-sharp":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharp
		return
	case "sharp-sharp":
		*enumaccidentalvalue = EnumAccidentalValueSharpSharp
		return
	case "flat-flat":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlat
		return
	case "natural-sharp":
		*enumaccidentalvalue = EnumAccidentalValueNaturalSharp
		return
	case "natural-flat":
		*enumaccidentalvalue = EnumAccidentalValueNaturalFlat
		return
	case "quarter-flat":
		*enumaccidentalvalue = EnumAccidentalValueQuarterFlat
		return
	case "quarter-sharp":
		*enumaccidentalvalue = EnumAccidentalValueQuarterSharp
		return
	case "three-quarters-flat":
		*enumaccidentalvalue = EnumAccidentalValueThreeQuartersFlat
		return
	case "three-quarters-sharp":
		*enumaccidentalvalue = EnumAccidentalValueThreeQuartersSharp
		return
	case "sharp-down":
		*enumaccidentalvalue = EnumAccidentalValueSharpDown
		return
	case "sharp-up":
		*enumaccidentalvalue = EnumAccidentalValueSharpUp
		return
	case "natural-down":
		*enumaccidentalvalue = EnumAccidentalValueNaturalDown
		return
	case "natural-up":
		*enumaccidentalvalue = EnumAccidentalValueNaturalUp
		return
	case "flat-down":
		*enumaccidentalvalue = EnumAccidentalValueFlatDown
		return
	case "flat-up":
		*enumaccidentalvalue = EnumAccidentalValueFlatUp
		return
	case "double-sharp-down":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharpDown
		return
	case "double-sharp-up":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharpUp
		return
	case "flat-flat-down":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlatDown
		return
	case "flat-flat-up":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlatUp
		return
	case "arrow-down":
		*enumaccidentalvalue = EnumAccidentalValueArrowDown
		return
	case "arrow-up":
		*enumaccidentalvalue = EnumAccidentalValueArrowUp
		return
	case "triple-sharp":
		*enumaccidentalvalue = EnumAccidentalValueTripleSharp
		return
	case "triple-flat":
		*enumaccidentalvalue = EnumAccidentalValueTripleFlat
		return
	case "slash-quarter-sharp":
		*enumaccidentalvalue = EnumAccidentalValueSlashQuarterSharp
		return
	case "slash-sharp":
		*enumaccidentalvalue = EnumAccidentalValueSlashSharp
		return
	case "slash-flat":
		*enumaccidentalvalue = EnumAccidentalValueSlashFlat
		return
	case "double-slash-flat":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSlashFlat
		return
	case "sharp-1":
		*enumaccidentalvalue = EnumAccidentalValueSharp1
		return
	case "sharp-2":
		*enumaccidentalvalue = EnumAccidentalValueSharp2
		return
	case "sharp-3":
		*enumaccidentalvalue = EnumAccidentalValueSharp3
		return
	case "sharp-5":
		*enumaccidentalvalue = EnumAccidentalValueSharp5
		return
	case "flat-1":
		*enumaccidentalvalue = EnumAccidentalValueFlat1
		return
	case "flat-2":
		*enumaccidentalvalue = EnumAccidentalValueFlat2
		return
	case "flat-3":
		*enumaccidentalvalue = EnumAccidentalValueFlat3
		return
	case "flat-4":
		*enumaccidentalvalue = EnumAccidentalValueFlat4
		return
	case "sori":
		*enumaccidentalvalue = EnumAccidentalValueSori
		return
	case "koron":
		*enumaccidentalvalue = EnumAccidentalValueKoron
		return
	case "other":
		*enumaccidentalvalue = EnumAccidentalValueOther
		return
	default:
		return errUnkownEnum
	}
}

func (enumaccidentalvalue *EnumAccidentalValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumAccidentalValueSharp":
		*enumaccidentalvalue = EnumAccidentalValueSharp
	case "EnumAccidentalValueNatural":
		*enumaccidentalvalue = EnumAccidentalValueNatural
	case "EnumAccidentalValueFlat":
		*enumaccidentalvalue = EnumAccidentalValueFlat
	case "EnumAccidentalValueDoubleSharp":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharp
	case "EnumAccidentalValueSharpSharp":
		*enumaccidentalvalue = EnumAccidentalValueSharpSharp
	case "EnumAccidentalValueFlatFlat":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlat
	case "EnumAccidentalValueNaturalSharp":
		*enumaccidentalvalue = EnumAccidentalValueNaturalSharp
	case "EnumAccidentalValueNaturalFlat":
		*enumaccidentalvalue = EnumAccidentalValueNaturalFlat
	case "EnumAccidentalValueQuarterFlat":
		*enumaccidentalvalue = EnumAccidentalValueQuarterFlat
	case "EnumAccidentalValueQuarterSharp":
		*enumaccidentalvalue = EnumAccidentalValueQuarterSharp
	case "EnumAccidentalValueThreeQuartersFlat":
		*enumaccidentalvalue = EnumAccidentalValueThreeQuartersFlat
	case "EnumAccidentalValueThreeQuartersSharp":
		*enumaccidentalvalue = EnumAccidentalValueThreeQuartersSharp
	case "EnumAccidentalValueSharpDown":
		*enumaccidentalvalue = EnumAccidentalValueSharpDown
	case "EnumAccidentalValueSharpUp":
		*enumaccidentalvalue = EnumAccidentalValueSharpUp
	case "EnumAccidentalValueNaturalDown":
		*enumaccidentalvalue = EnumAccidentalValueNaturalDown
	case "EnumAccidentalValueNaturalUp":
		*enumaccidentalvalue = EnumAccidentalValueNaturalUp
	case "EnumAccidentalValueFlatDown":
		*enumaccidentalvalue = EnumAccidentalValueFlatDown
	case "EnumAccidentalValueFlatUp":
		*enumaccidentalvalue = EnumAccidentalValueFlatUp
	case "EnumAccidentalValueDoubleSharpDown":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharpDown
	case "EnumAccidentalValueDoubleSharpUp":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSharpUp
	case "EnumAccidentalValueFlatFlatDown":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlatDown
	case "EnumAccidentalValueFlatFlatUp":
		*enumaccidentalvalue = EnumAccidentalValueFlatFlatUp
	case "EnumAccidentalValueArrowDown":
		*enumaccidentalvalue = EnumAccidentalValueArrowDown
	case "EnumAccidentalValueArrowUp":
		*enumaccidentalvalue = EnumAccidentalValueArrowUp
	case "EnumAccidentalValueTripleSharp":
		*enumaccidentalvalue = EnumAccidentalValueTripleSharp
	case "EnumAccidentalValueTripleFlat":
		*enumaccidentalvalue = EnumAccidentalValueTripleFlat
	case "EnumAccidentalValueSlashQuarterSharp":
		*enumaccidentalvalue = EnumAccidentalValueSlashQuarterSharp
	case "EnumAccidentalValueSlashSharp":
		*enumaccidentalvalue = EnumAccidentalValueSlashSharp
	case "EnumAccidentalValueSlashFlat":
		*enumaccidentalvalue = EnumAccidentalValueSlashFlat
	case "EnumAccidentalValueDoubleSlashFlat":
		*enumaccidentalvalue = EnumAccidentalValueDoubleSlashFlat
	case "EnumAccidentalValueSharp1":
		*enumaccidentalvalue = EnumAccidentalValueSharp1
	case "EnumAccidentalValueSharp2":
		*enumaccidentalvalue = EnumAccidentalValueSharp2
	case "EnumAccidentalValueSharp3":
		*enumaccidentalvalue = EnumAccidentalValueSharp3
	case "EnumAccidentalValueSharp5":
		*enumaccidentalvalue = EnumAccidentalValueSharp5
	case "EnumAccidentalValueFlat1":
		*enumaccidentalvalue = EnumAccidentalValueFlat1
	case "EnumAccidentalValueFlat2":
		*enumaccidentalvalue = EnumAccidentalValueFlat2
	case "EnumAccidentalValueFlat3":
		*enumaccidentalvalue = EnumAccidentalValueFlat3
	case "EnumAccidentalValueFlat4":
		*enumaccidentalvalue = EnumAccidentalValueFlat4
	case "EnumAccidentalValueSori":
		*enumaccidentalvalue = EnumAccidentalValueSori
	case "EnumAccidentalValueKoron":
		*enumaccidentalvalue = EnumAccidentalValueKoron
	case "EnumAccidentalValueOther":
		*enumaccidentalvalue = EnumAccidentalValueOther
	default:
		err = errUnkownEnum
	}
	return
}

func (enumaccidentalvalue *EnumAccidentalValue) ToCodeString() (res string) {

	switch *enumaccidentalvalue {
	// insertion code per enum code
	case EnumAccidentalValueSharp:
		res = "EnumAccidentalValueSharp"
	case EnumAccidentalValueNatural:
		res = "EnumAccidentalValueNatural"
	case EnumAccidentalValueFlat:
		res = "EnumAccidentalValueFlat"
	case EnumAccidentalValueDoubleSharp:
		res = "EnumAccidentalValueDoubleSharp"
	case EnumAccidentalValueSharpSharp:
		res = "EnumAccidentalValueSharpSharp"
	case EnumAccidentalValueFlatFlat:
		res = "EnumAccidentalValueFlatFlat"
	case EnumAccidentalValueNaturalSharp:
		res = "EnumAccidentalValueNaturalSharp"
	case EnumAccidentalValueNaturalFlat:
		res = "EnumAccidentalValueNaturalFlat"
	case EnumAccidentalValueQuarterFlat:
		res = "EnumAccidentalValueQuarterFlat"
	case EnumAccidentalValueQuarterSharp:
		res = "EnumAccidentalValueQuarterSharp"
	case EnumAccidentalValueThreeQuartersFlat:
		res = "EnumAccidentalValueThreeQuartersFlat"
	case EnumAccidentalValueThreeQuartersSharp:
		res = "EnumAccidentalValueThreeQuartersSharp"
	case EnumAccidentalValueSharpDown:
		res = "EnumAccidentalValueSharpDown"
	case EnumAccidentalValueSharpUp:
		res = "EnumAccidentalValueSharpUp"
	case EnumAccidentalValueNaturalDown:
		res = "EnumAccidentalValueNaturalDown"
	case EnumAccidentalValueNaturalUp:
		res = "EnumAccidentalValueNaturalUp"
	case EnumAccidentalValueFlatDown:
		res = "EnumAccidentalValueFlatDown"
	case EnumAccidentalValueFlatUp:
		res = "EnumAccidentalValueFlatUp"
	case EnumAccidentalValueDoubleSharpDown:
		res = "EnumAccidentalValueDoubleSharpDown"
	case EnumAccidentalValueDoubleSharpUp:
		res = "EnumAccidentalValueDoubleSharpUp"
	case EnumAccidentalValueFlatFlatDown:
		res = "EnumAccidentalValueFlatFlatDown"
	case EnumAccidentalValueFlatFlatUp:
		res = "EnumAccidentalValueFlatFlatUp"
	case EnumAccidentalValueArrowDown:
		res = "EnumAccidentalValueArrowDown"
	case EnumAccidentalValueArrowUp:
		res = "EnumAccidentalValueArrowUp"
	case EnumAccidentalValueTripleSharp:
		res = "EnumAccidentalValueTripleSharp"
	case EnumAccidentalValueTripleFlat:
		res = "EnumAccidentalValueTripleFlat"
	case EnumAccidentalValueSlashQuarterSharp:
		res = "EnumAccidentalValueSlashQuarterSharp"
	case EnumAccidentalValueSlashSharp:
		res = "EnumAccidentalValueSlashSharp"
	case EnumAccidentalValueSlashFlat:
		res = "EnumAccidentalValueSlashFlat"
	case EnumAccidentalValueDoubleSlashFlat:
		res = "EnumAccidentalValueDoubleSlashFlat"
	case EnumAccidentalValueSharp1:
		res = "EnumAccidentalValueSharp1"
	case EnumAccidentalValueSharp2:
		res = "EnumAccidentalValueSharp2"
	case EnumAccidentalValueSharp3:
		res = "EnumAccidentalValueSharp3"
	case EnumAccidentalValueSharp5:
		res = "EnumAccidentalValueSharp5"
	case EnumAccidentalValueFlat1:
		res = "EnumAccidentalValueFlat1"
	case EnumAccidentalValueFlat2:
		res = "EnumAccidentalValueFlat2"
	case EnumAccidentalValueFlat3:
		res = "EnumAccidentalValueFlat3"
	case EnumAccidentalValueFlat4:
		res = "EnumAccidentalValueFlat4"
	case EnumAccidentalValueSori:
		res = "EnumAccidentalValueSori"
	case EnumAccidentalValueKoron:
		res = "EnumAccidentalValueKoron"
	case EnumAccidentalValueOther:
		res = "EnumAccidentalValueOther"
	}
	return
}

func (enumaccidentalvalue EnumAccidentalValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumAccidentalValueSharp")
	res = append(res, "EnumAccidentalValueNatural")
	res = append(res, "EnumAccidentalValueFlat")
	res = append(res, "EnumAccidentalValueDoubleSharp")
	res = append(res, "EnumAccidentalValueSharpSharp")
	res = append(res, "EnumAccidentalValueFlatFlat")
	res = append(res, "EnumAccidentalValueNaturalSharp")
	res = append(res, "EnumAccidentalValueNaturalFlat")
	res = append(res, "EnumAccidentalValueQuarterFlat")
	res = append(res, "EnumAccidentalValueQuarterSharp")
	res = append(res, "EnumAccidentalValueThreeQuartersFlat")
	res = append(res, "EnumAccidentalValueThreeQuartersSharp")
	res = append(res, "EnumAccidentalValueSharpDown")
	res = append(res, "EnumAccidentalValueSharpUp")
	res = append(res, "EnumAccidentalValueNaturalDown")
	res = append(res, "EnumAccidentalValueNaturalUp")
	res = append(res, "EnumAccidentalValueFlatDown")
	res = append(res, "EnumAccidentalValueFlatUp")
	res = append(res, "EnumAccidentalValueDoubleSharpDown")
	res = append(res, "EnumAccidentalValueDoubleSharpUp")
	res = append(res, "EnumAccidentalValueFlatFlatDown")
	res = append(res, "EnumAccidentalValueFlatFlatUp")
	res = append(res, "EnumAccidentalValueArrowDown")
	res = append(res, "EnumAccidentalValueArrowUp")
	res = append(res, "EnumAccidentalValueTripleSharp")
	res = append(res, "EnumAccidentalValueTripleFlat")
	res = append(res, "EnumAccidentalValueSlashQuarterSharp")
	res = append(res, "EnumAccidentalValueSlashSharp")
	res = append(res, "EnumAccidentalValueSlashFlat")
	res = append(res, "EnumAccidentalValueDoubleSlashFlat")
	res = append(res, "EnumAccidentalValueSharp1")
	res = append(res, "EnumAccidentalValueSharp2")
	res = append(res, "EnumAccidentalValueSharp3")
	res = append(res, "EnumAccidentalValueSharp5")
	res = append(res, "EnumAccidentalValueFlat1")
	res = append(res, "EnumAccidentalValueFlat2")
	res = append(res, "EnumAccidentalValueFlat3")
	res = append(res, "EnumAccidentalValueFlat4")
	res = append(res, "EnumAccidentalValueSori")
	res = append(res, "EnumAccidentalValueKoron")
	res = append(res, "EnumAccidentalValueOther")

	return
}

func (enumaccidentalvalue EnumAccidentalValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sharp")
	res = append(res, "natural")
	res = append(res, "flat")
	res = append(res, "double-sharp")
	res = append(res, "sharp-sharp")
	res = append(res, "flat-flat")
	res = append(res, "natural-sharp")
	res = append(res, "natural-flat")
	res = append(res, "quarter-flat")
	res = append(res, "quarter-sharp")
	res = append(res, "three-quarters-flat")
	res = append(res, "three-quarters-sharp")
	res = append(res, "sharp-down")
	res = append(res, "sharp-up")
	res = append(res, "natural-down")
	res = append(res, "natural-up")
	res = append(res, "flat-down")
	res = append(res, "flat-up")
	res = append(res, "double-sharp-down")
	res = append(res, "double-sharp-up")
	res = append(res, "flat-flat-down")
	res = append(res, "flat-flat-up")
	res = append(res, "arrow-down")
	res = append(res, "arrow-up")
	res = append(res, "triple-sharp")
	res = append(res, "triple-flat")
	res = append(res, "slash-quarter-sharp")
	res = append(res, "slash-sharp")
	res = append(res, "slash-flat")
	res = append(res, "double-slash-flat")
	res = append(res, "sharp-1")
	res = append(res, "sharp-2")
	res = append(res, "sharp-3")
	res = append(res, "sharp-5")
	res = append(res, "flat-1")
	res = append(res, "flat-2")
	res = append(res, "flat-3")
	res = append(res, "flat-4")
	res = append(res, "sori")
	res = append(res, "koron")
	res = append(res, "other")

	return
}

// Utility function for EnumCssFontSize
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumcssfontsize EnumCssFontSize) ToString() (res string) {

	// migration of former implementation of enum
	switch enumcssfontsize {
	// insertion code per enum code
	case EnumCssFontSizeXxSmall:
		res = "xx-small"
	case EnumCssFontSizeXSmall:
		res = "x-small"
	case EnumCssFontSizeSmall:
		res = "small"
	case EnumCssFontSizeMedium:
		res = "medium"
	case EnumCssFontSizeLarge:
		res = "large"
	case EnumCssFontSizeXLarge:
		res = "x-large"
	case EnumCssFontSizeXxLarge:
		res = "xx-large"
	}
	return
}

func (enumcssfontsize *EnumCssFontSize) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "xx-small":
		*enumcssfontsize = EnumCssFontSizeXxSmall
		return
	case "x-small":
		*enumcssfontsize = EnumCssFontSizeXSmall
		return
	case "small":
		*enumcssfontsize = EnumCssFontSizeSmall
		return
	case "medium":
		*enumcssfontsize = EnumCssFontSizeMedium
		return
	case "large":
		*enumcssfontsize = EnumCssFontSizeLarge
		return
	case "x-large":
		*enumcssfontsize = EnumCssFontSizeXLarge
		return
	case "xx-large":
		*enumcssfontsize = EnumCssFontSizeXxLarge
		return
	default:
		return errUnkownEnum
	}
}

func (enumcssfontsize *EnumCssFontSize) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumCssFontSizeXxSmall":
		*enumcssfontsize = EnumCssFontSizeXxSmall
	case "EnumCssFontSizeXSmall":
		*enumcssfontsize = EnumCssFontSizeXSmall
	case "EnumCssFontSizeSmall":
		*enumcssfontsize = EnumCssFontSizeSmall
	case "EnumCssFontSizeMedium":
		*enumcssfontsize = EnumCssFontSizeMedium
	case "EnumCssFontSizeLarge":
		*enumcssfontsize = EnumCssFontSizeLarge
	case "EnumCssFontSizeXLarge":
		*enumcssfontsize = EnumCssFontSizeXLarge
	case "EnumCssFontSizeXxLarge":
		*enumcssfontsize = EnumCssFontSizeXxLarge
	default:
		err = errUnkownEnum
	}
	return
}

func (enumcssfontsize *EnumCssFontSize) ToCodeString() (res string) {

	switch *enumcssfontsize {
	// insertion code per enum code
	case EnumCssFontSizeXxSmall:
		res = "EnumCssFontSizeXxSmall"
	case EnumCssFontSizeXSmall:
		res = "EnumCssFontSizeXSmall"
	case EnumCssFontSizeSmall:
		res = "EnumCssFontSizeSmall"
	case EnumCssFontSizeMedium:
		res = "EnumCssFontSizeMedium"
	case EnumCssFontSizeLarge:
		res = "EnumCssFontSizeLarge"
	case EnumCssFontSizeXLarge:
		res = "EnumCssFontSizeXLarge"
	case EnumCssFontSizeXxLarge:
		res = "EnumCssFontSizeXxLarge"
	}
	return
}

func (enumcssfontsize EnumCssFontSize) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumCssFontSizeXxSmall")
	res = append(res, "EnumCssFontSizeXSmall")
	res = append(res, "EnumCssFontSizeSmall")
	res = append(res, "EnumCssFontSizeMedium")
	res = append(res, "EnumCssFontSizeLarge")
	res = append(res, "EnumCssFontSizeXLarge")
	res = append(res, "EnumCssFontSizeXxLarge")

	return
}

func (enumcssfontsize EnumCssFontSize) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "xx-small")
	res = append(res, "x-small")
	res = append(res, "small")
	res = append(res, "medium")
	res = append(res, "large")
	res = append(res, "x-large")
	res = append(res, "xx-large")

	return
}

// Utility function for EnumEnclosureShape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumenclosureshape EnumEnclosureShape) ToString() (res string) {

	// migration of former implementation of enum
	switch enumenclosureshape {
	// insertion code per enum code
	case EnumEnclosureShapeRectangle:
		res = "rectangle"
	case EnumEnclosureShapeSquare:
		res = "square"
	case EnumEnclosureShapeOval:
		res = "oval"
	case EnumEnclosureShapeCircle:
		res = "circle"
	case EnumEnclosureShapeBracket:
		res = "bracket"
	case EnumEnclosureShapeInvertedBracket:
		res = "inverted-bracket"
	case EnumEnclosureShapeTriangle:
		res = "triangle"
	case EnumEnclosureShapeDiamond:
		res = "diamond"
	case EnumEnclosureShapePentagon:
		res = "pentagon"
	case EnumEnclosureShapeHexagon:
		res = "hexagon"
	case EnumEnclosureShapeHeptagon:
		res = "heptagon"
	case EnumEnclosureShapeOctagon:
		res = "octagon"
	case EnumEnclosureShapeNonagon:
		res = "nonagon"
	case EnumEnclosureShapeDecagon:
		res = "decagon"
	case EnumEnclosureShapeNone:
		res = "none"
	}
	return
}

func (enumenclosureshape *EnumEnclosureShape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "rectangle":
		*enumenclosureshape = EnumEnclosureShapeRectangle
		return
	case "square":
		*enumenclosureshape = EnumEnclosureShapeSquare
		return
	case "oval":
		*enumenclosureshape = EnumEnclosureShapeOval
		return
	case "circle":
		*enumenclosureshape = EnumEnclosureShapeCircle
		return
	case "bracket":
		*enumenclosureshape = EnumEnclosureShapeBracket
		return
	case "inverted-bracket":
		*enumenclosureshape = EnumEnclosureShapeInvertedBracket
		return
	case "triangle":
		*enumenclosureshape = EnumEnclosureShapeTriangle
		return
	case "diamond":
		*enumenclosureshape = EnumEnclosureShapeDiamond
		return
	case "pentagon":
		*enumenclosureshape = EnumEnclosureShapePentagon
		return
	case "hexagon":
		*enumenclosureshape = EnumEnclosureShapeHexagon
		return
	case "heptagon":
		*enumenclosureshape = EnumEnclosureShapeHeptagon
		return
	case "octagon":
		*enumenclosureshape = EnumEnclosureShapeOctagon
		return
	case "nonagon":
		*enumenclosureshape = EnumEnclosureShapeNonagon
		return
	case "decagon":
		*enumenclosureshape = EnumEnclosureShapeDecagon
		return
	case "none":
		*enumenclosureshape = EnumEnclosureShapeNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumenclosureshape *EnumEnclosureShape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumEnclosureShapeRectangle":
		*enumenclosureshape = EnumEnclosureShapeRectangle
	case "EnumEnclosureShapeSquare":
		*enumenclosureshape = EnumEnclosureShapeSquare
	case "EnumEnclosureShapeOval":
		*enumenclosureshape = EnumEnclosureShapeOval
	case "EnumEnclosureShapeCircle":
		*enumenclosureshape = EnumEnclosureShapeCircle
	case "EnumEnclosureShapeBracket":
		*enumenclosureshape = EnumEnclosureShapeBracket
	case "EnumEnclosureShapeInvertedBracket":
		*enumenclosureshape = EnumEnclosureShapeInvertedBracket
	case "EnumEnclosureShapeTriangle":
		*enumenclosureshape = EnumEnclosureShapeTriangle
	case "EnumEnclosureShapeDiamond":
		*enumenclosureshape = EnumEnclosureShapeDiamond
	case "EnumEnclosureShapePentagon":
		*enumenclosureshape = EnumEnclosureShapePentagon
	case "EnumEnclosureShapeHexagon":
		*enumenclosureshape = EnumEnclosureShapeHexagon
	case "EnumEnclosureShapeHeptagon":
		*enumenclosureshape = EnumEnclosureShapeHeptagon
	case "EnumEnclosureShapeOctagon":
		*enumenclosureshape = EnumEnclosureShapeOctagon
	case "EnumEnclosureShapeNonagon":
		*enumenclosureshape = EnumEnclosureShapeNonagon
	case "EnumEnclosureShapeDecagon":
		*enumenclosureshape = EnumEnclosureShapeDecagon
	case "EnumEnclosureShapeNone":
		*enumenclosureshape = EnumEnclosureShapeNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumenclosureshape *EnumEnclosureShape) ToCodeString() (res string) {

	switch *enumenclosureshape {
	// insertion code per enum code
	case EnumEnclosureShapeRectangle:
		res = "EnumEnclosureShapeRectangle"
	case EnumEnclosureShapeSquare:
		res = "EnumEnclosureShapeSquare"
	case EnumEnclosureShapeOval:
		res = "EnumEnclosureShapeOval"
	case EnumEnclosureShapeCircle:
		res = "EnumEnclosureShapeCircle"
	case EnumEnclosureShapeBracket:
		res = "EnumEnclosureShapeBracket"
	case EnumEnclosureShapeInvertedBracket:
		res = "EnumEnclosureShapeInvertedBracket"
	case EnumEnclosureShapeTriangle:
		res = "EnumEnclosureShapeTriangle"
	case EnumEnclosureShapeDiamond:
		res = "EnumEnclosureShapeDiamond"
	case EnumEnclosureShapePentagon:
		res = "EnumEnclosureShapePentagon"
	case EnumEnclosureShapeHexagon:
		res = "EnumEnclosureShapeHexagon"
	case EnumEnclosureShapeHeptagon:
		res = "EnumEnclosureShapeHeptagon"
	case EnumEnclosureShapeOctagon:
		res = "EnumEnclosureShapeOctagon"
	case EnumEnclosureShapeNonagon:
		res = "EnumEnclosureShapeNonagon"
	case EnumEnclosureShapeDecagon:
		res = "EnumEnclosureShapeDecagon"
	case EnumEnclosureShapeNone:
		res = "EnumEnclosureShapeNone"
	}
	return
}

func (enumenclosureshape EnumEnclosureShape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumEnclosureShapeRectangle")
	res = append(res, "EnumEnclosureShapeSquare")
	res = append(res, "EnumEnclosureShapeOval")
	res = append(res, "EnumEnclosureShapeCircle")
	res = append(res, "EnumEnclosureShapeBracket")
	res = append(res, "EnumEnclosureShapeInvertedBracket")
	res = append(res, "EnumEnclosureShapeTriangle")
	res = append(res, "EnumEnclosureShapeDiamond")
	res = append(res, "EnumEnclosureShapePentagon")
	res = append(res, "EnumEnclosureShapeHexagon")
	res = append(res, "EnumEnclosureShapeHeptagon")
	res = append(res, "EnumEnclosureShapeOctagon")
	res = append(res, "EnumEnclosureShapeNonagon")
	res = append(res, "EnumEnclosureShapeDecagon")
	res = append(res, "EnumEnclosureShapeNone")

	return
}

func (enumenclosureshape EnumEnclosureShape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "rectangle")
	res = append(res, "square")
	res = append(res, "oval")
	res = append(res, "circle")
	res = append(res, "bracket")
	res = append(res, "inverted-bracket")
	res = append(res, "triangle")
	res = append(res, "diamond")
	res = append(res, "pentagon")
	res = append(res, "hexagon")
	res = append(res, "heptagon")
	res = append(res, "octagon")
	res = append(res, "nonagon")
	res = append(res, "decagon")
	res = append(res, "none")

	return
}

// Utility function for EnumFermataShape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumfermatashape EnumFermataShape) ToString() (res string) {

	// migration of former implementation of enum
	switch enumfermatashape {
	// insertion code per enum code
	case EnumFermataShapeNormal:
		res = "normal"
	case EnumFermataShapeAngled:
		res = "angled"
	case EnumFermataShapeSquare:
		res = "square"
	case EnumFermataShapeDoubleAngled:
		res = "double-angled"
	case EnumFermataShapeDoubleSquare:
		res = "double-square"
	case EnumFermataShapeDoubleDot:
		res = "double-dot"
	case EnumFermataShapeHalfCurve:
		res = "half-curve"
	case EnumFermataShapeCurlew:
		res = "curlew"
	case EnumFermataShapeEmpty:
		res = ""
	}
	return
}

func (enumfermatashape *EnumFermataShape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enumfermatashape = EnumFermataShapeNormal
		return
	case "angled":
		*enumfermatashape = EnumFermataShapeAngled
		return
	case "square":
		*enumfermatashape = EnumFermataShapeSquare
		return
	case "double-angled":
		*enumfermatashape = EnumFermataShapeDoubleAngled
		return
	case "double-square":
		*enumfermatashape = EnumFermataShapeDoubleSquare
		return
	case "double-dot":
		*enumfermatashape = EnumFermataShapeDoubleDot
		return
	case "half-curve":
		*enumfermatashape = EnumFermataShapeHalfCurve
		return
	case "curlew":
		*enumfermatashape = EnumFermataShapeCurlew
		return
	case "":
		*enumfermatashape = EnumFermataShapeEmpty
		return
	default:
		return errUnkownEnum
	}
}

func (enumfermatashape *EnumFermataShape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumFermataShapeNormal":
		*enumfermatashape = EnumFermataShapeNormal
	case "EnumFermataShapeAngled":
		*enumfermatashape = EnumFermataShapeAngled
	case "EnumFermataShapeSquare":
		*enumfermatashape = EnumFermataShapeSquare
	case "EnumFermataShapeDoubleAngled":
		*enumfermatashape = EnumFermataShapeDoubleAngled
	case "EnumFermataShapeDoubleSquare":
		*enumfermatashape = EnumFermataShapeDoubleSquare
	case "EnumFermataShapeDoubleDot":
		*enumfermatashape = EnumFermataShapeDoubleDot
	case "EnumFermataShapeHalfCurve":
		*enumfermatashape = EnumFermataShapeHalfCurve
	case "EnumFermataShapeCurlew":
		*enumfermatashape = EnumFermataShapeCurlew
	case "EnumFermataShapeEmpty":
		*enumfermatashape = EnumFermataShapeEmpty
	default:
		err = errUnkownEnum
	}
	return
}

func (enumfermatashape *EnumFermataShape) ToCodeString() (res string) {

	switch *enumfermatashape {
	// insertion code per enum code
	case EnumFermataShapeNormal:
		res = "EnumFermataShapeNormal"
	case EnumFermataShapeAngled:
		res = "EnumFermataShapeAngled"
	case EnumFermataShapeSquare:
		res = "EnumFermataShapeSquare"
	case EnumFermataShapeDoubleAngled:
		res = "EnumFermataShapeDoubleAngled"
	case EnumFermataShapeDoubleSquare:
		res = "EnumFermataShapeDoubleSquare"
	case EnumFermataShapeDoubleDot:
		res = "EnumFermataShapeDoubleDot"
	case EnumFermataShapeHalfCurve:
		res = "EnumFermataShapeHalfCurve"
	case EnumFermataShapeCurlew:
		res = "EnumFermataShapeCurlew"
	case EnumFermataShapeEmpty:
		res = "EnumFermataShapeEmpty"
	}
	return
}

func (enumfermatashape EnumFermataShape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumFermataShapeNormal")
	res = append(res, "EnumFermataShapeAngled")
	res = append(res, "EnumFermataShapeSquare")
	res = append(res, "EnumFermataShapeDoubleAngled")
	res = append(res, "EnumFermataShapeDoubleSquare")
	res = append(res, "EnumFermataShapeDoubleDot")
	res = append(res, "EnumFermataShapeHalfCurve")
	res = append(res, "EnumFermataShapeCurlew")
	res = append(res, "EnumFermataShapeEmpty")

	return
}

func (enumfermatashape EnumFermataShape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "angled")
	res = append(res, "square")
	res = append(res, "double-angled")
	res = append(res, "double-square")
	res = append(res, "double-dot")
	res = append(res, "half-curve")
	res = append(res, "curlew")
	res = append(res, "")

	return
}

// Utility function for EnumFontStyle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumfontstyle EnumFontStyle) ToString() (res string) {

	// migration of former implementation of enum
	switch enumfontstyle {
	// insertion code per enum code
	case EnumFontStyleNormal:
		res = "normal"
	case EnumFontStyleItalic:
		res = "italic"
	}
	return
}

func (enumfontstyle *EnumFontStyle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enumfontstyle = EnumFontStyleNormal
		return
	case "italic":
		*enumfontstyle = EnumFontStyleItalic
		return
	default:
		return errUnkownEnum
	}
}

func (enumfontstyle *EnumFontStyle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumFontStyleNormal":
		*enumfontstyle = EnumFontStyleNormal
	case "EnumFontStyleItalic":
		*enumfontstyle = EnumFontStyleItalic
	default:
		err = errUnkownEnum
	}
	return
}

func (enumfontstyle *EnumFontStyle) ToCodeString() (res string) {

	switch *enumfontstyle {
	// insertion code per enum code
	case EnumFontStyleNormal:
		res = "EnumFontStyleNormal"
	case EnumFontStyleItalic:
		res = "EnumFontStyleItalic"
	}
	return
}

func (enumfontstyle EnumFontStyle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumFontStyleNormal")
	res = append(res, "EnumFontStyleItalic")

	return
}

func (enumfontstyle EnumFontStyle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "italic")

	return
}

// Utility function for EnumFontWeight
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumfontweight EnumFontWeight) ToString() (res string) {

	// migration of former implementation of enum
	switch enumfontweight {
	// insertion code per enum code
	case EnumFontWeightNormal:
		res = "normal"
	case EnumFontWeightBold:
		res = "bold"
	}
	return
}

func (enumfontweight *EnumFontWeight) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enumfontweight = EnumFontWeightNormal
		return
	case "bold":
		*enumfontweight = EnumFontWeightBold
		return
	default:
		return errUnkownEnum
	}
}

func (enumfontweight *EnumFontWeight) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumFontWeightNormal":
		*enumfontweight = EnumFontWeightNormal
	case "EnumFontWeightBold":
		*enumfontweight = EnumFontWeightBold
	default:
		err = errUnkownEnum
	}
	return
}

func (enumfontweight *EnumFontWeight) ToCodeString() (res string) {

	switch *enumfontweight {
	// insertion code per enum code
	case EnumFontWeightNormal:
		res = "EnumFontWeightNormal"
	case EnumFontWeightBold:
		res = "EnumFontWeightBold"
	}
	return
}

func (enumfontweight EnumFontWeight) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumFontWeightNormal")
	res = append(res, "EnumFontWeightBold")

	return
}

func (enumfontweight EnumFontWeight) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "bold")

	return
}

// Utility function for EnumLeftCenterRight
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumleftcenterright EnumLeftCenterRight) ToString() (res string) {

	// migration of former implementation of enum
	switch enumleftcenterright {
	// insertion code per enum code
	case EnumLeftCenterRightLeft:
		res = "left"
	case EnumLeftCenterRightCenter:
		res = "center"
	case EnumLeftCenterRightRight:
		res = "right"
	}
	return
}

func (enumleftcenterright *EnumLeftCenterRight) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enumleftcenterright = EnumLeftCenterRightLeft
		return
	case "center":
		*enumleftcenterright = EnumLeftCenterRightCenter
		return
	case "right":
		*enumleftcenterright = EnumLeftCenterRightRight
		return
	default:
		return errUnkownEnum
	}
}

func (enumleftcenterright *EnumLeftCenterRight) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLeftCenterRightLeft":
		*enumleftcenterright = EnumLeftCenterRightLeft
	case "EnumLeftCenterRightCenter":
		*enumleftcenterright = EnumLeftCenterRightCenter
	case "EnumLeftCenterRightRight":
		*enumleftcenterright = EnumLeftCenterRightRight
	default:
		err = errUnkownEnum
	}
	return
}

func (enumleftcenterright *EnumLeftCenterRight) ToCodeString() (res string) {

	switch *enumleftcenterright {
	// insertion code per enum code
	case EnumLeftCenterRightLeft:
		res = "EnumLeftCenterRightLeft"
	case EnumLeftCenterRightCenter:
		res = "EnumLeftCenterRightCenter"
	case EnumLeftCenterRightRight:
		res = "EnumLeftCenterRightRight"
	}
	return
}

func (enumleftcenterright EnumLeftCenterRight) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLeftCenterRightLeft")
	res = append(res, "EnumLeftCenterRightCenter")
	res = append(res, "EnumLeftCenterRightRight")

	return
}

func (enumleftcenterright EnumLeftCenterRight) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "center")
	res = append(res, "right")

	return
}

// Utility function for EnumLeftRight
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumleftright EnumLeftRight) ToString() (res string) {

	// migration of former implementation of enum
	switch enumleftright {
	// insertion code per enum code
	case EnumLeftRightLeft:
		res = "left"
	case EnumLeftRightRight:
		res = "right"
	}
	return
}

func (enumleftright *EnumLeftRight) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enumleftright = EnumLeftRightLeft
		return
	case "right":
		*enumleftright = EnumLeftRightRight
		return
	default:
		return errUnkownEnum
	}
}

func (enumleftright *EnumLeftRight) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLeftRightLeft":
		*enumleftright = EnumLeftRightLeft
	case "EnumLeftRightRight":
		*enumleftright = EnumLeftRightRight
	default:
		err = errUnkownEnum
	}
	return
}

func (enumleftright *EnumLeftRight) ToCodeString() (res string) {

	switch *enumleftright {
	// insertion code per enum code
	case EnumLeftRightLeft:
		res = "EnumLeftRightLeft"
	case EnumLeftRightRight:
		res = "EnumLeftRightRight"
	}
	return
}

func (enumleftright EnumLeftRight) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLeftRightLeft")
	res = append(res, "EnumLeftRightRight")

	return
}

func (enumleftright EnumLeftRight) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for EnumLineLength
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumlinelength EnumLineLength) ToString() (res string) {

	// migration of former implementation of enum
	switch enumlinelength {
	// insertion code per enum code
	case EnumLineLengthShort:
		res = "short"
	case EnumLineLengthMedium:
		res = "medium"
	case EnumLineLengthLong:
		res = "long"
	}
	return
}

func (enumlinelength *EnumLineLength) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "short":
		*enumlinelength = EnumLineLengthShort
		return
	case "medium":
		*enumlinelength = EnumLineLengthMedium
		return
	case "long":
		*enumlinelength = EnumLineLengthLong
		return
	default:
		return errUnkownEnum
	}
}

func (enumlinelength *EnumLineLength) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLineLengthShort":
		*enumlinelength = EnumLineLengthShort
	case "EnumLineLengthMedium":
		*enumlinelength = EnumLineLengthMedium
	case "EnumLineLengthLong":
		*enumlinelength = EnumLineLengthLong
	default:
		err = errUnkownEnum
	}
	return
}

func (enumlinelength *EnumLineLength) ToCodeString() (res string) {

	switch *enumlinelength {
	// insertion code per enum code
	case EnumLineLengthShort:
		res = "EnumLineLengthShort"
	case EnumLineLengthMedium:
		res = "EnumLineLengthMedium"
	case EnumLineLengthLong:
		res = "EnumLineLengthLong"
	}
	return
}

func (enumlinelength EnumLineLength) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLineLengthShort")
	res = append(res, "EnumLineLengthMedium")
	res = append(res, "EnumLineLengthLong")

	return
}

func (enumlinelength EnumLineLength) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "short")
	res = append(res, "medium")
	res = append(res, "long")

	return
}

// Utility function for EnumLineShape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumlineshape EnumLineShape) ToString() (res string) {

	// migration of former implementation of enum
	switch enumlineshape {
	// insertion code per enum code
	case EnumLineShapeStraight:
		res = "straight"
	case EnumLineShapeCurved:
		res = "curved"
	}
	return
}

func (enumlineshape *EnumLineShape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "straight":
		*enumlineshape = EnumLineShapeStraight
		return
	case "curved":
		*enumlineshape = EnumLineShapeCurved
		return
	default:
		return errUnkownEnum
	}
}

func (enumlineshape *EnumLineShape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLineShapeStraight":
		*enumlineshape = EnumLineShapeStraight
	case "EnumLineShapeCurved":
		*enumlineshape = EnumLineShapeCurved
	default:
		err = errUnkownEnum
	}
	return
}

func (enumlineshape *EnumLineShape) ToCodeString() (res string) {

	switch *enumlineshape {
	// insertion code per enum code
	case EnumLineShapeStraight:
		res = "EnumLineShapeStraight"
	case EnumLineShapeCurved:
		res = "EnumLineShapeCurved"
	}
	return
}

func (enumlineshape EnumLineShape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLineShapeStraight")
	res = append(res, "EnumLineShapeCurved")

	return
}

func (enumlineshape EnumLineShape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "straight")
	res = append(res, "curved")

	return
}

// Utility function for EnumLineType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumlinetype EnumLineType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumlinetype {
	// insertion code per enum code
	case EnumLineTypeSolid:
		res = "solid"
	case EnumLineTypeDashed:
		res = "dashed"
	case EnumLineTypeDotted:
		res = "dotted"
	case EnumLineTypeWavy:
		res = "wavy"
	}
	return
}

func (enumlinetype *EnumLineType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "solid":
		*enumlinetype = EnumLineTypeSolid
		return
	case "dashed":
		*enumlinetype = EnumLineTypeDashed
		return
	case "dotted":
		*enumlinetype = EnumLineTypeDotted
		return
	case "wavy":
		*enumlinetype = EnumLineTypeWavy
		return
	default:
		return errUnkownEnum
	}
}

func (enumlinetype *EnumLineType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumLineTypeSolid":
		*enumlinetype = EnumLineTypeSolid
	case "EnumLineTypeDashed":
		*enumlinetype = EnumLineTypeDashed
	case "EnumLineTypeDotted":
		*enumlinetype = EnumLineTypeDotted
	case "EnumLineTypeWavy":
		*enumlinetype = EnumLineTypeWavy
	default:
		err = errUnkownEnum
	}
	return
}

func (enumlinetype *EnumLineType) ToCodeString() (res string) {

	switch *enumlinetype {
	// insertion code per enum code
	case EnumLineTypeSolid:
		res = "EnumLineTypeSolid"
	case EnumLineTypeDashed:
		res = "EnumLineTypeDashed"
	case EnumLineTypeDotted:
		res = "EnumLineTypeDotted"
	case EnumLineTypeWavy:
		res = "EnumLineTypeWavy"
	}
	return
}

func (enumlinetype EnumLineType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumLineTypeSolid")
	res = append(res, "EnumLineTypeDashed")
	res = append(res, "EnumLineTypeDotted")
	res = append(res, "EnumLineTypeWavy")

	return
}

func (enumlinetype EnumLineType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "solid")
	res = append(res, "dashed")
	res = append(res, "dotted")
	res = append(res, "wavy")

	return
}

// Utility function for EnumMarginType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enummargintype EnumMarginType) ToString() (res string) {

	// migration of former implementation of enum
	switch enummargintype {
	// insertion code per enum code
	case EnumMarginTypeOdd:
		res = "odd"
	case EnumMarginTypeEven:
		res = "even"
	case EnumMarginTypeBoth:
		res = "both"
	}
	return
}

func (enummargintype *EnumMarginType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "odd":
		*enummargintype = EnumMarginTypeOdd
		return
	case "even":
		*enummargintype = EnumMarginTypeEven
		return
	case "both":
		*enummargintype = EnumMarginTypeBoth
		return
	default:
		return errUnkownEnum
	}
}

func (enummargintype *EnumMarginType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumMarginTypeOdd":
		*enummargintype = EnumMarginTypeOdd
	case "EnumMarginTypeEven":
		*enummargintype = EnumMarginTypeEven
	case "EnumMarginTypeBoth":
		*enummargintype = EnumMarginTypeBoth
	default:
		err = errUnkownEnum
	}
	return
}

func (enummargintype *EnumMarginType) ToCodeString() (res string) {

	switch *enummargintype {
	// insertion code per enum code
	case EnumMarginTypeOdd:
		res = "EnumMarginTypeOdd"
	case EnumMarginTypeEven:
		res = "EnumMarginTypeEven"
	case EnumMarginTypeBoth:
		res = "EnumMarginTypeBoth"
	}
	return
}

func (enummargintype EnumMarginType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumMarginTypeOdd")
	res = append(res, "EnumMarginTypeEven")
	res = append(res, "EnumMarginTypeBoth")

	return
}

func (enummargintype EnumMarginType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "odd")
	res = append(res, "even")
	res = append(res, "both")

	return
}

// Utility function for EnumMute
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enummute EnumMute) ToString() (res string) {

	// migration of former implementation of enum
	switch enummute {
	// insertion code per enum code
	case EnumMuteOn:
		res = "on"
	case EnumMuteOff:
		res = "off"
	case EnumMuteStraight:
		res = "straight"
	case EnumMuteCup:
		res = "cup"
	case EnumMuteHarmonNoStem:
		res = "harmon-no-stem"
	case EnumMuteHarmonStem:
		res = "harmon-stem"
	case EnumMuteBucket:
		res = "bucket"
	case EnumMutePlunger:
		res = "plunger"
	case EnumMuteHat:
		res = "hat"
	case EnumMuteSolotone:
		res = "solotone"
	case EnumMutePractice:
		res = "practice"
	case EnumMuteStopMute:
		res = "stop-mute"
	case EnumMuteStopHand:
		res = "stop-hand"
	case EnumMuteEcho:
		res = "echo"
	case EnumMutePalm:
		res = "palm"
	}
	return
}

func (enummute *EnumMute) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*enummute = EnumMuteOn
		return
	case "off":
		*enummute = EnumMuteOff
		return
	case "straight":
		*enummute = EnumMuteStraight
		return
	case "cup":
		*enummute = EnumMuteCup
		return
	case "harmon-no-stem":
		*enummute = EnumMuteHarmonNoStem
		return
	case "harmon-stem":
		*enummute = EnumMuteHarmonStem
		return
	case "bucket":
		*enummute = EnumMuteBucket
		return
	case "plunger":
		*enummute = EnumMutePlunger
		return
	case "hat":
		*enummute = EnumMuteHat
		return
	case "solotone":
		*enummute = EnumMuteSolotone
		return
	case "practice":
		*enummute = EnumMutePractice
		return
	case "stop-mute":
		*enummute = EnumMuteStopMute
		return
	case "stop-hand":
		*enummute = EnumMuteStopHand
		return
	case "echo":
		*enummute = EnumMuteEcho
		return
	case "palm":
		*enummute = EnumMutePalm
		return
	default:
		return errUnkownEnum
	}
}

func (enummute *EnumMute) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumMuteOn":
		*enummute = EnumMuteOn
	case "EnumMuteOff":
		*enummute = EnumMuteOff
	case "EnumMuteStraight":
		*enummute = EnumMuteStraight
	case "EnumMuteCup":
		*enummute = EnumMuteCup
	case "EnumMuteHarmonNoStem":
		*enummute = EnumMuteHarmonNoStem
	case "EnumMuteHarmonStem":
		*enummute = EnumMuteHarmonStem
	case "EnumMuteBucket":
		*enummute = EnumMuteBucket
	case "EnumMutePlunger":
		*enummute = EnumMutePlunger
	case "EnumMuteHat":
		*enummute = EnumMuteHat
	case "EnumMuteSolotone":
		*enummute = EnumMuteSolotone
	case "EnumMutePractice":
		*enummute = EnumMutePractice
	case "EnumMuteStopMute":
		*enummute = EnumMuteStopMute
	case "EnumMuteStopHand":
		*enummute = EnumMuteStopHand
	case "EnumMuteEcho":
		*enummute = EnumMuteEcho
	case "EnumMutePalm":
		*enummute = EnumMutePalm
	default:
		err = errUnkownEnum
	}
	return
}

func (enummute *EnumMute) ToCodeString() (res string) {

	switch *enummute {
	// insertion code per enum code
	case EnumMuteOn:
		res = "EnumMuteOn"
	case EnumMuteOff:
		res = "EnumMuteOff"
	case EnumMuteStraight:
		res = "EnumMuteStraight"
	case EnumMuteCup:
		res = "EnumMuteCup"
	case EnumMuteHarmonNoStem:
		res = "EnumMuteHarmonNoStem"
	case EnumMuteHarmonStem:
		res = "EnumMuteHarmonStem"
	case EnumMuteBucket:
		res = "EnumMuteBucket"
	case EnumMutePlunger:
		res = "EnumMutePlunger"
	case EnumMuteHat:
		res = "EnumMuteHat"
	case EnumMuteSolotone:
		res = "EnumMuteSolotone"
	case EnumMutePractice:
		res = "EnumMutePractice"
	case EnumMuteStopMute:
		res = "EnumMuteStopMute"
	case EnumMuteStopHand:
		res = "EnumMuteStopHand"
	case EnumMuteEcho:
		res = "EnumMuteEcho"
	case EnumMutePalm:
		res = "EnumMutePalm"
	}
	return
}

func (enummute EnumMute) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumMuteOn")
	res = append(res, "EnumMuteOff")
	res = append(res, "EnumMuteStraight")
	res = append(res, "EnumMuteCup")
	res = append(res, "EnumMuteHarmonNoStem")
	res = append(res, "EnumMuteHarmonStem")
	res = append(res, "EnumMuteBucket")
	res = append(res, "EnumMutePlunger")
	res = append(res, "EnumMuteHat")
	res = append(res, "EnumMuteSolotone")
	res = append(res, "EnumMutePractice")
	res = append(res, "EnumMuteStopMute")
	res = append(res, "EnumMuteStopHand")
	res = append(res, "EnumMuteEcho")
	res = append(res, "EnumMutePalm")

	return
}

func (enummute EnumMute) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "on")
	res = append(res, "off")
	res = append(res, "straight")
	res = append(res, "cup")
	res = append(res, "harmon-no-stem")
	res = append(res, "harmon-stem")
	res = append(res, "bucket")
	res = append(res, "plunger")
	res = append(res, "hat")
	res = append(res, "solotone")
	res = append(res, "practice")
	res = append(res, "stop-mute")
	res = append(res, "stop-hand")
	res = append(res, "echo")
	res = append(res, "palm")

	return
}

// Utility function for EnumNoteSizeType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumnotesizetype EnumNoteSizeType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumnotesizetype {
	// insertion code per enum code
	case EnumNoteSizeTypeCue:
		res = "cue"
	case EnumNoteSizeTypeGrace:
		res = "grace"
	case EnumNoteSizeTypeGraceCue:
		res = "grace-cue"
	case EnumNoteSizeTypeLarge:
		res = "large"
	}
	return
}

func (enumnotesizetype *EnumNoteSizeType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "cue":
		*enumnotesizetype = EnumNoteSizeTypeCue
		return
	case "grace":
		*enumnotesizetype = EnumNoteSizeTypeGrace
		return
	case "grace-cue":
		*enumnotesizetype = EnumNoteSizeTypeGraceCue
		return
	case "large":
		*enumnotesizetype = EnumNoteSizeTypeLarge
		return
	default:
		return errUnkownEnum
	}
}

func (enumnotesizetype *EnumNoteSizeType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumNoteSizeTypeCue":
		*enumnotesizetype = EnumNoteSizeTypeCue
	case "EnumNoteSizeTypeGrace":
		*enumnotesizetype = EnumNoteSizeTypeGrace
	case "EnumNoteSizeTypeGraceCue":
		*enumnotesizetype = EnumNoteSizeTypeGraceCue
	case "EnumNoteSizeTypeLarge":
		*enumnotesizetype = EnumNoteSizeTypeLarge
	default:
		err = errUnkownEnum
	}
	return
}

func (enumnotesizetype *EnumNoteSizeType) ToCodeString() (res string) {

	switch *enumnotesizetype {
	// insertion code per enum code
	case EnumNoteSizeTypeCue:
		res = "EnumNoteSizeTypeCue"
	case EnumNoteSizeTypeGrace:
		res = "EnumNoteSizeTypeGrace"
	case EnumNoteSizeTypeGraceCue:
		res = "EnumNoteSizeTypeGraceCue"
	case EnumNoteSizeTypeLarge:
		res = "EnumNoteSizeTypeLarge"
	}
	return
}

func (enumnotesizetype EnumNoteSizeType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumNoteSizeTypeCue")
	res = append(res, "EnumNoteSizeTypeGrace")
	res = append(res, "EnumNoteSizeTypeGraceCue")
	res = append(res, "EnumNoteSizeTypeLarge")

	return
}

func (enumnotesizetype EnumNoteSizeType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "cue")
	res = append(res, "grace")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for EnumNoteTypeValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumnotetypevalue EnumNoteTypeValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumnotetypevalue {
	// insertion code per enum code
	case EnumNoteTypeValueN1024th:
		res = "1024th"
	case EnumNoteTypeValueN512th:
		res = "512th"
	case EnumNoteTypeValueN256th:
		res = "256th"
	case EnumNoteTypeValueN128th:
		res = "128th"
	case EnumNoteTypeValueN64th:
		res = "64th"
	case EnumNoteTypeValueN32nd:
		res = "32nd"
	case EnumNoteTypeValueN16th:
		res = "16th"
	case EnumNoteTypeValueEighth:
		res = "eighth"
	case EnumNoteTypeValueQuarter:
		res = "quarter"
	case EnumNoteTypeValueHalf:
		res = "half"
	case EnumNoteTypeValueWhole:
		res = "whole"
	case EnumNoteTypeValueBreve:
		res = "breve"
	case EnumNoteTypeValueLong:
		res = "long"
	case EnumNoteTypeValueMaxima:
		res = "maxima"
	}
	return
}

func (enumnotetypevalue *EnumNoteTypeValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1024th":
		*enumnotetypevalue = EnumNoteTypeValueN1024th
		return
	case "512th":
		*enumnotetypevalue = EnumNoteTypeValueN512th
		return
	case "256th":
		*enumnotetypevalue = EnumNoteTypeValueN256th
		return
	case "128th":
		*enumnotetypevalue = EnumNoteTypeValueN128th
		return
	case "64th":
		*enumnotetypevalue = EnumNoteTypeValueN64th
		return
	case "32nd":
		*enumnotetypevalue = EnumNoteTypeValueN32nd
		return
	case "16th":
		*enumnotetypevalue = EnumNoteTypeValueN16th
		return
	case "eighth":
		*enumnotetypevalue = EnumNoteTypeValueEighth
		return
	case "quarter":
		*enumnotetypevalue = EnumNoteTypeValueQuarter
		return
	case "half":
		*enumnotetypevalue = EnumNoteTypeValueHalf
		return
	case "whole":
		*enumnotetypevalue = EnumNoteTypeValueWhole
		return
	case "breve":
		*enumnotetypevalue = EnumNoteTypeValueBreve
		return
	case "long":
		*enumnotetypevalue = EnumNoteTypeValueLong
		return
	case "maxima":
		*enumnotetypevalue = EnumNoteTypeValueMaxima
		return
	default:
		return errUnkownEnum
	}
}

func (enumnotetypevalue *EnumNoteTypeValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumNoteTypeValueN1024th":
		*enumnotetypevalue = EnumNoteTypeValueN1024th
	case "EnumNoteTypeValueN512th":
		*enumnotetypevalue = EnumNoteTypeValueN512th
	case "EnumNoteTypeValueN256th":
		*enumnotetypevalue = EnumNoteTypeValueN256th
	case "EnumNoteTypeValueN128th":
		*enumnotetypevalue = EnumNoteTypeValueN128th
	case "EnumNoteTypeValueN64th":
		*enumnotetypevalue = EnumNoteTypeValueN64th
	case "EnumNoteTypeValueN32nd":
		*enumnotetypevalue = EnumNoteTypeValueN32nd
	case "EnumNoteTypeValueN16th":
		*enumnotetypevalue = EnumNoteTypeValueN16th
	case "EnumNoteTypeValueEighth":
		*enumnotetypevalue = EnumNoteTypeValueEighth
	case "EnumNoteTypeValueQuarter":
		*enumnotetypevalue = EnumNoteTypeValueQuarter
	case "EnumNoteTypeValueHalf":
		*enumnotetypevalue = EnumNoteTypeValueHalf
	case "EnumNoteTypeValueWhole":
		*enumnotetypevalue = EnumNoteTypeValueWhole
	case "EnumNoteTypeValueBreve":
		*enumnotetypevalue = EnumNoteTypeValueBreve
	case "EnumNoteTypeValueLong":
		*enumnotetypevalue = EnumNoteTypeValueLong
	case "EnumNoteTypeValueMaxima":
		*enumnotetypevalue = EnumNoteTypeValueMaxima
	default:
		err = errUnkownEnum
	}
	return
}

func (enumnotetypevalue *EnumNoteTypeValue) ToCodeString() (res string) {

	switch *enumnotetypevalue {
	// insertion code per enum code
	case EnumNoteTypeValueN1024th:
		res = "EnumNoteTypeValueN1024th"
	case EnumNoteTypeValueN512th:
		res = "EnumNoteTypeValueN512th"
	case EnumNoteTypeValueN256th:
		res = "EnumNoteTypeValueN256th"
	case EnumNoteTypeValueN128th:
		res = "EnumNoteTypeValueN128th"
	case EnumNoteTypeValueN64th:
		res = "EnumNoteTypeValueN64th"
	case EnumNoteTypeValueN32nd:
		res = "EnumNoteTypeValueN32nd"
	case EnumNoteTypeValueN16th:
		res = "EnumNoteTypeValueN16th"
	case EnumNoteTypeValueEighth:
		res = "EnumNoteTypeValueEighth"
	case EnumNoteTypeValueQuarter:
		res = "EnumNoteTypeValueQuarter"
	case EnumNoteTypeValueHalf:
		res = "EnumNoteTypeValueHalf"
	case EnumNoteTypeValueWhole:
		res = "EnumNoteTypeValueWhole"
	case EnumNoteTypeValueBreve:
		res = "EnumNoteTypeValueBreve"
	case EnumNoteTypeValueLong:
		res = "EnumNoteTypeValueLong"
	case EnumNoteTypeValueMaxima:
		res = "EnumNoteTypeValueMaxima"
	}
	return
}

func (enumnotetypevalue EnumNoteTypeValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumNoteTypeValueN1024th")
	res = append(res, "EnumNoteTypeValueN512th")
	res = append(res, "EnumNoteTypeValueN256th")
	res = append(res, "EnumNoteTypeValueN128th")
	res = append(res, "EnumNoteTypeValueN64th")
	res = append(res, "EnumNoteTypeValueN32nd")
	res = append(res, "EnumNoteTypeValueN16th")
	res = append(res, "EnumNoteTypeValueEighth")
	res = append(res, "EnumNoteTypeValueQuarter")
	res = append(res, "EnumNoteTypeValueHalf")
	res = append(res, "EnumNoteTypeValueWhole")
	res = append(res, "EnumNoteTypeValueBreve")
	res = append(res, "EnumNoteTypeValueLong")
	res = append(res, "EnumNoteTypeValueMaxima")

	return
}

func (enumnotetypevalue EnumNoteTypeValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1024th")
	res = append(res, "512th")
	res = append(res, "256th")
	res = append(res, "128th")
	res = append(res, "64th")
	res = append(res, "32nd")
	res = append(res, "16th")
	res = append(res, "eighth")
	res = append(res, "quarter")
	res = append(res, "half")
	res = append(res, "whole")
	res = append(res, "breve")
	res = append(res, "long")
	res = append(res, "maxima")

	return
}

// Utility function for EnumOverUnder
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumoverunder EnumOverUnder) ToString() (res string) {

	// migration of former implementation of enum
	switch enumoverunder {
	// insertion code per enum code
	case EnumOverUnderOver:
		res = "over"
	case EnumOverUnderUnder:
		res = "under"
	}
	return
}

func (enumoverunder *EnumOverUnder) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "over":
		*enumoverunder = EnumOverUnderOver
		return
	case "under":
		*enumoverunder = EnumOverUnderUnder
		return
	default:
		return errUnkownEnum
	}
}

func (enumoverunder *EnumOverUnder) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumOverUnderOver":
		*enumoverunder = EnumOverUnderOver
	case "EnumOverUnderUnder":
		*enumoverunder = EnumOverUnderUnder
	default:
		err = errUnkownEnum
	}
	return
}

func (enumoverunder *EnumOverUnder) ToCodeString() (res string) {

	switch *enumoverunder {
	// insertion code per enum code
	case EnumOverUnderOver:
		res = "EnumOverUnderOver"
	case EnumOverUnderUnder:
		res = "EnumOverUnderUnder"
	}
	return
}

func (enumoverunder EnumOverUnder) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumOverUnderOver")
	res = append(res, "EnumOverUnderUnder")

	return
}

func (enumoverunder EnumOverUnder) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "over")
	res = append(res, "under")

	return
}

// Utility function for EnumSemiPitched
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsemipitched EnumSemiPitched) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsemipitched {
	// insertion code per enum code
	case EnumSemiPitchedHigh:
		res = "high"
	case EnumSemiPitchedMediumHigh:
		res = "medium-high"
	case EnumSemiPitchedMedium:
		res = "medium"
	case EnumSemiPitchedMediumLow:
		res = "medium-low"
	case EnumSemiPitchedLow:
		res = "low"
	case EnumSemiPitchedVeryLow:
		res = "very-low"
	}
	return
}

func (enumsemipitched *EnumSemiPitched) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "high":
		*enumsemipitched = EnumSemiPitchedHigh
		return
	case "medium-high":
		*enumsemipitched = EnumSemiPitchedMediumHigh
		return
	case "medium":
		*enumsemipitched = EnumSemiPitchedMedium
		return
	case "medium-low":
		*enumsemipitched = EnumSemiPitchedMediumLow
		return
	case "low":
		*enumsemipitched = EnumSemiPitchedLow
		return
	case "very-low":
		*enumsemipitched = EnumSemiPitchedVeryLow
		return
	default:
		return errUnkownEnum
	}
}

func (enumsemipitched *EnumSemiPitched) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSemiPitchedHigh":
		*enumsemipitched = EnumSemiPitchedHigh
	case "EnumSemiPitchedMediumHigh":
		*enumsemipitched = EnumSemiPitchedMediumHigh
	case "EnumSemiPitchedMedium":
		*enumsemipitched = EnumSemiPitchedMedium
	case "EnumSemiPitchedMediumLow":
		*enumsemipitched = EnumSemiPitchedMediumLow
	case "EnumSemiPitchedLow":
		*enumsemipitched = EnumSemiPitchedLow
	case "EnumSemiPitchedVeryLow":
		*enumsemipitched = EnumSemiPitchedVeryLow
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsemipitched *EnumSemiPitched) ToCodeString() (res string) {

	switch *enumsemipitched {
	// insertion code per enum code
	case EnumSemiPitchedHigh:
		res = "EnumSemiPitchedHigh"
	case EnumSemiPitchedMediumHigh:
		res = "EnumSemiPitchedMediumHigh"
	case EnumSemiPitchedMedium:
		res = "EnumSemiPitchedMedium"
	case EnumSemiPitchedMediumLow:
		res = "EnumSemiPitchedMediumLow"
	case EnumSemiPitchedLow:
		res = "EnumSemiPitchedLow"
	case EnumSemiPitchedVeryLow:
		res = "EnumSemiPitchedVeryLow"
	}
	return
}

func (enumsemipitched EnumSemiPitched) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSemiPitchedHigh")
	res = append(res, "EnumSemiPitchedMediumHigh")
	res = append(res, "EnumSemiPitchedMedium")
	res = append(res, "EnumSemiPitchedMediumLow")
	res = append(res, "EnumSemiPitchedLow")
	res = append(res, "EnumSemiPitchedVeryLow")

	return
}

func (enumsemipitched EnumSemiPitched) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "high")
	res = append(res, "medium-high")
	res = append(res, "medium")
	res = append(res, "medium-low")
	res = append(res, "low")
	res = append(res, "very-low")

	return
}

// Utility function for EnumStartNote
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartnote EnumStartNote) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartnote {
	// insertion code per enum code
	case EnumStartNoteUpper:
		res = "upper"
	case EnumStartNoteMain:
		res = "main"
	case EnumStartNoteBelow:
		res = "below"
	}
	return
}

func (enumstartnote *EnumStartNote) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upper":
		*enumstartnote = EnumStartNoteUpper
		return
	case "main":
		*enumstartnote = EnumStartNoteMain
		return
	case "below":
		*enumstartnote = EnumStartNoteBelow
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartnote *EnumStartNote) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartNoteUpper":
		*enumstartnote = EnumStartNoteUpper
	case "EnumStartNoteMain":
		*enumstartnote = EnumStartNoteMain
	case "EnumStartNoteBelow":
		*enumstartnote = EnumStartNoteBelow
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartnote *EnumStartNote) ToCodeString() (res string) {

	switch *enumstartnote {
	// insertion code per enum code
	case EnumStartNoteUpper:
		res = "EnumStartNoteUpper"
	case EnumStartNoteMain:
		res = "EnumStartNoteMain"
	case EnumStartNoteBelow:
		res = "EnumStartNoteBelow"
	}
	return
}

func (enumstartnote EnumStartNote) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartNoteUpper")
	res = append(res, "EnumStartNoteMain")
	res = append(res, "EnumStartNoteBelow")

	return
}

func (enumstartnote EnumStartNote) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upper")
	res = append(res, "main")
	res = append(res, "below")

	return
}

// Utility function for EnumStartStop
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartstop EnumStartStop) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartstop {
	// insertion code per enum code
	case EnumStartStopStart:
		res = "start"
	case EnumStartStopStop:
		res = "stop"
	}
	return
}

func (enumstartstop *EnumStartStop) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumstartstop = EnumStartStopStart
		return
	case "stop":
		*enumstartstop = EnumStartStopStop
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartstop *EnumStartStop) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartStopStart":
		*enumstartstop = EnumStartStopStart
	case "EnumStartStopStop":
		*enumstartstop = EnumStartStopStop
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartstop *EnumStartStop) ToCodeString() (res string) {

	switch *enumstartstop {
	// insertion code per enum code
	case EnumStartStopStart:
		res = "EnumStartStopStart"
	case EnumStartStopStop:
		res = "EnumStartStopStop"
	}
	return
}

func (enumstartstop EnumStartStop) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartStopStart")
	res = append(res, "EnumStartStopStop")

	return
}

func (enumstartstop EnumStartStop) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")

	return
}

// Utility function for EnumStartStopContinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartstopcontinue EnumStartStopContinue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartstopcontinue {
	// insertion code per enum code
	case EnumStartStopContinueStart:
		res = "start"
	case EnumStartStopContinueStop:
		res = "stop"
	case EnumStartStopContinueContinue:
		res = "continue"
	}
	return
}

func (enumstartstopcontinue *EnumStartStopContinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumstartstopcontinue = EnumStartStopContinueStart
		return
	case "stop":
		*enumstartstopcontinue = EnumStartStopContinueStop
		return
	case "continue":
		*enumstartstopcontinue = EnumStartStopContinueContinue
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartstopcontinue *EnumStartStopContinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartStopContinueStart":
		*enumstartstopcontinue = EnumStartStopContinueStart
	case "EnumStartStopContinueStop":
		*enumstartstopcontinue = EnumStartStopContinueStop
	case "EnumStartStopContinueContinue":
		*enumstartstopcontinue = EnumStartStopContinueContinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartstopcontinue *EnumStartStopContinue) ToCodeString() (res string) {

	switch *enumstartstopcontinue {
	// insertion code per enum code
	case EnumStartStopContinueStart:
		res = "EnumStartStopContinueStart"
	case EnumStartStopContinueStop:
		res = "EnumStartStopContinueStop"
	case EnumStartStopContinueContinue:
		res = "EnumStartStopContinueContinue"
	}
	return
}

func (enumstartstopcontinue EnumStartStopContinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartStopContinueStart")
	res = append(res, "EnumStartStopContinueStop")
	res = append(res, "EnumStartStopContinueContinue")

	return
}

func (enumstartstopcontinue EnumStartStopContinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for EnumStartStopSingle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumstartstopsingle EnumStartStopSingle) ToString() (res string) {

	// migration of former implementation of enum
	switch enumstartstopsingle {
	// insertion code per enum code
	case EnumStartStopSingleStart:
		res = "start"
	case EnumStartStopSingleStop:
		res = "stop"
	case EnumStartStopSingleSingle:
		res = "single"
	}
	return
}

func (enumstartstopsingle *EnumStartStopSingle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumstartstopsingle = EnumStartStopSingleStart
		return
	case "stop":
		*enumstartstopsingle = EnumStartStopSingleStop
		return
	case "single":
		*enumstartstopsingle = EnumStartStopSingleSingle
		return
	default:
		return errUnkownEnum
	}
}

func (enumstartstopsingle *EnumStartStopSingle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumStartStopSingleStart":
		*enumstartstopsingle = EnumStartStopSingleStart
	case "EnumStartStopSingleStop":
		*enumstartstopsingle = EnumStartStopSingleStop
	case "EnumStartStopSingleSingle":
		*enumstartstopsingle = EnumStartStopSingleSingle
	default:
		err = errUnkownEnum
	}
	return
}

func (enumstartstopsingle *EnumStartStopSingle) ToCodeString() (res string) {

	switch *enumstartstopsingle {
	// insertion code per enum code
	case EnumStartStopSingleStart:
		res = "EnumStartStopSingleStart"
	case EnumStartStopSingleStop:
		res = "EnumStartStopSingleStop"
	case EnumStartStopSingleSingle:
		res = "EnumStartStopSingleSingle"
	}
	return
}

func (enumstartstopsingle EnumStartStopSingle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumStartStopSingleStart")
	res = append(res, "EnumStartStopSingleStop")
	res = append(res, "EnumStartStopSingleSingle")

	return
}

func (enumstartstopsingle EnumStartStopSingle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")

	return
}

// Utility function for EnumSwingTypeValue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumswingtypevalue EnumSwingTypeValue) ToString() (res string) {

	// migration of former implementation of enum
	switch enumswingtypevalue {
	// insertion code per enum code
	case EnumSwingTypeValueN16th:
		res = "16th"
	case EnumSwingTypeValueEighth:
		res = "eighth"
	}
	return
}

func (enumswingtypevalue *EnumSwingTypeValue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "16th":
		*enumswingtypevalue = EnumSwingTypeValueN16th
		return
	case "eighth":
		*enumswingtypevalue = EnumSwingTypeValueEighth
		return
	default:
		return errUnkownEnum
	}
}

func (enumswingtypevalue *EnumSwingTypeValue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSwingTypeValueN16th":
		*enumswingtypevalue = EnumSwingTypeValueN16th
	case "EnumSwingTypeValueEighth":
		*enumswingtypevalue = EnumSwingTypeValueEighth
	default:
		err = errUnkownEnum
	}
	return
}

func (enumswingtypevalue *EnumSwingTypeValue) ToCodeString() (res string) {

	switch *enumswingtypevalue {
	// insertion code per enum code
	case EnumSwingTypeValueN16th:
		res = "EnumSwingTypeValueN16th"
	case EnumSwingTypeValueEighth:
		res = "EnumSwingTypeValueEighth"
	}
	return
}

func (enumswingtypevalue EnumSwingTypeValue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSwingTypeValueN16th")
	res = append(res, "EnumSwingTypeValueEighth")

	return
}

func (enumswingtypevalue EnumSwingTypeValue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "16th")
	res = append(res, "eighth")

	return
}

// Utility function for EnumSymbolSize
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumsymbolsize EnumSymbolSize) ToString() (res string) {

	// migration of former implementation of enum
	switch enumsymbolsize {
	// insertion code per enum code
	case EnumSymbolSizeFull:
		res = "full"
	case EnumSymbolSizeCue:
		res = "cue"
	case EnumSymbolSizeGraceCue:
		res = "grace-cue"
	case EnumSymbolSizeLarge:
		res = "large"
	}
	return
}

func (enumsymbolsize *EnumSymbolSize) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "full":
		*enumsymbolsize = EnumSymbolSizeFull
		return
	case "cue":
		*enumsymbolsize = EnumSymbolSizeCue
		return
	case "grace-cue":
		*enumsymbolsize = EnumSymbolSizeGraceCue
		return
	case "large":
		*enumsymbolsize = EnumSymbolSizeLarge
		return
	default:
		return errUnkownEnum
	}
}

func (enumsymbolsize *EnumSymbolSize) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumSymbolSizeFull":
		*enumsymbolsize = EnumSymbolSizeFull
	case "EnumSymbolSizeCue":
		*enumsymbolsize = EnumSymbolSizeCue
	case "EnumSymbolSizeGraceCue":
		*enumsymbolsize = EnumSymbolSizeGraceCue
	case "EnumSymbolSizeLarge":
		*enumsymbolsize = EnumSymbolSizeLarge
	default:
		err = errUnkownEnum
	}
	return
}

func (enumsymbolsize *EnumSymbolSize) ToCodeString() (res string) {

	switch *enumsymbolsize {
	// insertion code per enum code
	case EnumSymbolSizeFull:
		res = "EnumSymbolSizeFull"
	case EnumSymbolSizeCue:
		res = "EnumSymbolSizeCue"
	case EnumSymbolSizeGraceCue:
		res = "EnumSymbolSizeGraceCue"
	case EnumSymbolSizeLarge:
		res = "EnumSymbolSizeLarge"
	}
	return
}

func (enumsymbolsize EnumSymbolSize) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumSymbolSizeFull")
	res = append(res, "EnumSymbolSizeCue")
	res = append(res, "EnumSymbolSizeGraceCue")
	res = append(res, "EnumSymbolSizeLarge")

	return
}

func (enumsymbolsize EnumSymbolSize) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "full")
	res = append(res, "cue")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for EnumTextDirection
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtextdirection EnumTextDirection) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtextdirection {
	// insertion code per enum code
	case EnumTextDirectionLtr:
		res = "ltr"
	case EnumTextDirectionRtl:
		res = "rtl"
	case EnumTextDirectionLro:
		res = "lro"
	case EnumTextDirectionRlo:
		res = "rlo"
	}
	return
}

func (enumtextdirection *EnumTextDirection) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ltr":
		*enumtextdirection = EnumTextDirectionLtr
		return
	case "rtl":
		*enumtextdirection = EnumTextDirectionRtl
		return
	case "lro":
		*enumtextdirection = EnumTextDirectionLro
		return
	case "rlo":
		*enumtextdirection = EnumTextDirectionRlo
		return
	default:
		return errUnkownEnum
	}
}

func (enumtextdirection *EnumTextDirection) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTextDirectionLtr":
		*enumtextdirection = EnumTextDirectionLtr
	case "EnumTextDirectionRtl":
		*enumtextdirection = EnumTextDirectionRtl
	case "EnumTextDirectionLro":
		*enumtextdirection = EnumTextDirectionLro
	case "EnumTextDirectionRlo":
		*enumtextdirection = EnumTextDirectionRlo
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtextdirection *EnumTextDirection) ToCodeString() (res string) {

	switch *enumtextdirection {
	// insertion code per enum code
	case EnumTextDirectionLtr:
		res = "EnumTextDirectionLtr"
	case EnumTextDirectionRtl:
		res = "EnumTextDirectionRtl"
	case EnumTextDirectionLro:
		res = "EnumTextDirectionLro"
	case EnumTextDirectionRlo:
		res = "EnumTextDirectionRlo"
	}
	return
}

func (enumtextdirection EnumTextDirection) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTextDirectionLtr")
	res = append(res, "EnumTextDirectionRtl")
	res = append(res, "EnumTextDirectionLro")
	res = append(res, "EnumTextDirectionRlo")

	return
}

func (enumtextdirection EnumTextDirection) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ltr")
	res = append(res, "rtl")
	res = append(res, "lro")
	res = append(res, "rlo")

	return
}

// Utility function for EnumTiedType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtiedtype EnumTiedType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtiedtype {
	// insertion code per enum code
	case EnumTiedTypeStart:
		res = "start"
	case EnumTiedTypeStop:
		res = "stop"
	case EnumTiedTypeContinue:
		res = "continue"
	case EnumTiedTypeLetRing:
		res = "let-ring"
	}
	return
}

func (enumtiedtype *EnumTiedType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumtiedtype = EnumTiedTypeStart
		return
	case "stop":
		*enumtiedtype = EnumTiedTypeStop
		return
	case "continue":
		*enumtiedtype = EnumTiedTypeContinue
		return
	case "let-ring":
		*enumtiedtype = EnumTiedTypeLetRing
		return
	default:
		return errUnkownEnum
	}
}

func (enumtiedtype *EnumTiedType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTiedTypeStart":
		*enumtiedtype = EnumTiedTypeStart
	case "EnumTiedTypeStop":
		*enumtiedtype = EnumTiedTypeStop
	case "EnumTiedTypeContinue":
		*enumtiedtype = EnumTiedTypeContinue
	case "EnumTiedTypeLetRing":
		*enumtiedtype = EnumTiedTypeLetRing
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtiedtype *EnumTiedType) ToCodeString() (res string) {

	switch *enumtiedtype {
	// insertion code per enum code
	case EnumTiedTypeStart:
		res = "EnumTiedTypeStart"
	case EnumTiedTypeStop:
		res = "EnumTiedTypeStop"
	case EnumTiedTypeContinue:
		res = "EnumTiedTypeContinue"
	case EnumTiedTypeLetRing:
		res = "EnumTiedTypeLetRing"
	}
	return
}

func (enumtiedtype EnumTiedType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTiedTypeStart")
	res = append(res, "EnumTiedTypeStop")
	res = append(res, "EnumTiedTypeContinue")
	res = append(res, "EnumTiedTypeLetRing")

	return
}

func (enumtiedtype EnumTiedType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")
	res = append(res, "let-ring")

	return
}

// Utility function for EnumTopBottom
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtopbottom EnumTopBottom) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtopbottom {
	// insertion code per enum code
	case EnumTopBottomTop:
		res = "top"
	case EnumTopBottomBottom:
		res = "bottom"
	}
	return
}

func (enumtopbottom *EnumTopBottom) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enumtopbottom = EnumTopBottomTop
		return
	case "bottom":
		*enumtopbottom = EnumTopBottomBottom
		return
	default:
		return errUnkownEnum
	}
}

func (enumtopbottom *EnumTopBottom) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTopBottomTop":
		*enumtopbottom = EnumTopBottomTop
	case "EnumTopBottomBottom":
		*enumtopbottom = EnumTopBottomBottom
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtopbottom *EnumTopBottom) ToCodeString() (res string) {

	switch *enumtopbottom {
	// insertion code per enum code
	case EnumTopBottomTop:
		res = "EnumTopBottomTop"
	case EnumTopBottomBottom:
		res = "EnumTopBottomBottom"
	}
	return
}

func (enumtopbottom EnumTopBottom) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTopBottomTop")
	res = append(res, "EnumTopBottomBottom")

	return
}

func (enumtopbottom EnumTopBottom) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "bottom")

	return
}

// Utility function for EnumTremoloType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtremolotype EnumTremoloType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtremolotype {
	// insertion code per enum code
	case EnumTremoloTypeStart:
		res = "start"
	case EnumTremoloTypeStop:
		res = "stop"
	case EnumTremoloTypeSingle:
		res = "single"
	case EnumTremoloTypeUnmeasured:
		res = "unmeasured"
	}
	return
}

func (enumtremolotype *EnumTremoloType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enumtremolotype = EnumTremoloTypeStart
		return
	case "stop":
		*enumtremolotype = EnumTremoloTypeStop
		return
	case "single":
		*enumtremolotype = EnumTremoloTypeSingle
		return
	case "unmeasured":
		*enumtremolotype = EnumTremoloTypeUnmeasured
		return
	default:
		return errUnkownEnum
	}
}

func (enumtremolotype *EnumTremoloType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTremoloTypeStart":
		*enumtremolotype = EnumTremoloTypeStart
	case "EnumTremoloTypeStop":
		*enumtremolotype = EnumTremoloTypeStop
	case "EnumTremoloTypeSingle":
		*enumtremolotype = EnumTremoloTypeSingle
	case "EnumTremoloTypeUnmeasured":
		*enumtremolotype = EnumTremoloTypeUnmeasured
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtremolotype *EnumTremoloType) ToCodeString() (res string) {

	switch *enumtremolotype {
	// insertion code per enum code
	case EnumTremoloTypeStart:
		res = "EnumTremoloTypeStart"
	case EnumTremoloTypeStop:
		res = "EnumTremoloTypeStop"
	case EnumTremoloTypeSingle:
		res = "EnumTremoloTypeSingle"
	case EnumTremoloTypeUnmeasured:
		res = "EnumTremoloTypeUnmeasured"
	}
	return
}

func (enumtremolotype EnumTremoloType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTremoloTypeStart")
	res = append(res, "EnumTremoloTypeStop")
	res = append(res, "EnumTremoloTypeSingle")
	res = append(res, "EnumTremoloTypeUnmeasured")

	return
}

func (enumtremolotype EnumTremoloType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")
	res = append(res, "unmeasured")

	return
}

// Utility function for EnumTrillStep
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtrillstep EnumTrillStep) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtrillstep {
	// insertion code per enum code
	case EnumTrillStepWhole:
		res = "whole"
	case EnumTrillStepHalf:
		res = "half"
	case EnumTrillStepUnison:
		res = "unison"
	}
	return
}

func (enumtrillstep *EnumTrillStep) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*enumtrillstep = EnumTrillStepWhole
		return
	case "half":
		*enumtrillstep = EnumTrillStepHalf
		return
	case "unison":
		*enumtrillstep = EnumTrillStepUnison
		return
	default:
		return errUnkownEnum
	}
}

func (enumtrillstep *EnumTrillStep) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTrillStepWhole":
		*enumtrillstep = EnumTrillStepWhole
	case "EnumTrillStepHalf":
		*enumtrillstep = EnumTrillStepHalf
	case "EnumTrillStepUnison":
		*enumtrillstep = EnumTrillStepUnison
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtrillstep *EnumTrillStep) ToCodeString() (res string) {

	switch *enumtrillstep {
	// insertion code per enum code
	case EnumTrillStepWhole:
		res = "EnumTrillStepWhole"
	case EnumTrillStepHalf:
		res = "EnumTrillStepHalf"
	case EnumTrillStepUnison:
		res = "EnumTrillStepUnison"
	}
	return
}

func (enumtrillstep EnumTrillStep) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTrillStepWhole")
	res = append(res, "EnumTrillStepHalf")
	res = append(res, "EnumTrillStepUnison")

	return
}

func (enumtrillstep EnumTrillStep) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "unison")

	return
}

// Utility function for EnumTwoNoteTurn
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtwonoteturn EnumTwoNoteTurn) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtwonoteturn {
	// insertion code per enum code
	case EnumTwoNoteTurnWhole:
		res = "whole"
	case EnumTwoNoteTurnHalf:
		res = "half"
	case EnumTwoNoteTurnNone:
		res = "none"
	}
	return
}

func (enumtwonoteturn *EnumTwoNoteTurn) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*enumtwonoteturn = EnumTwoNoteTurnWhole
		return
	case "half":
		*enumtwonoteturn = EnumTwoNoteTurnHalf
		return
	case "none":
		*enumtwonoteturn = EnumTwoNoteTurnNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumtwonoteturn *EnumTwoNoteTurn) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTwoNoteTurnWhole":
		*enumtwonoteturn = EnumTwoNoteTurnWhole
	case "EnumTwoNoteTurnHalf":
		*enumtwonoteturn = EnumTwoNoteTurnHalf
	case "EnumTwoNoteTurnNone":
		*enumtwonoteturn = EnumTwoNoteTurnNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtwonoteturn *EnumTwoNoteTurn) ToCodeString() (res string) {

	switch *enumtwonoteturn {
	// insertion code per enum code
	case EnumTwoNoteTurnWhole:
		res = "EnumTwoNoteTurnWhole"
	case EnumTwoNoteTurnHalf:
		res = "EnumTwoNoteTurnHalf"
	case EnumTwoNoteTurnNone:
		res = "EnumTwoNoteTurnNone"
	}
	return
}

func (enumtwonoteturn EnumTwoNoteTurn) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTwoNoteTurnWhole")
	res = append(res, "EnumTwoNoteTurnHalf")
	res = append(res, "EnumTwoNoteTurnNone")

	return
}

func (enumtwonoteturn EnumTwoNoteTurn) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "none")

	return
}

// Utility function for EnumUpDown
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumupdown EnumUpDown) ToString() (res string) {

	// migration of former implementation of enum
	switch enumupdown {
	// insertion code per enum code
	case EnumUpDownUp:
		res = "up"
	case EnumUpDownDown:
		res = "down"
	}
	return
}

func (enumupdown *EnumUpDown) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enumupdown = EnumUpDownUp
		return
	case "down":
		*enumupdown = EnumUpDownDown
		return
	default:
		return errUnkownEnum
	}
}

func (enumupdown *EnumUpDown) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumUpDownUp":
		*enumupdown = EnumUpDownUp
	case "EnumUpDownDown":
		*enumupdown = EnumUpDownDown
	default:
		err = errUnkownEnum
	}
	return
}

func (enumupdown *EnumUpDown) ToCodeString() (res string) {

	switch *enumupdown {
	// insertion code per enum code
	case EnumUpDownUp:
		res = "EnumUpDownUp"
	case EnumUpDownDown:
		res = "EnumUpDownDown"
	}
	return
}

func (enumupdown EnumUpDown) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumUpDownUp")
	res = append(res, "EnumUpDownDown")

	return
}

func (enumupdown EnumUpDown) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")

	return
}

// Utility function for EnumUprightInverted
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumuprightinverted EnumUprightInverted) ToString() (res string) {

	// migration of former implementation of enum
	switch enumuprightinverted {
	// insertion code per enum code
	case EnumUprightInvertedUpright:
		res = "upright"
	case EnumUprightInvertedInverted:
		res = "inverted"
	}
	return
}

func (enumuprightinverted *EnumUprightInverted) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upright":
		*enumuprightinverted = EnumUprightInvertedUpright
		return
	case "inverted":
		*enumuprightinverted = EnumUprightInvertedInverted
		return
	default:
		return errUnkownEnum
	}
}

func (enumuprightinverted *EnumUprightInverted) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumUprightInvertedUpright":
		*enumuprightinverted = EnumUprightInvertedUpright
	case "EnumUprightInvertedInverted":
		*enumuprightinverted = EnumUprightInvertedInverted
	default:
		err = errUnkownEnum
	}
	return
}

func (enumuprightinverted *EnumUprightInverted) ToCodeString() (res string) {

	switch *enumuprightinverted {
	// insertion code per enum code
	case EnumUprightInvertedUpright:
		res = "EnumUprightInvertedUpright"
	case EnumUprightInvertedInverted:
		res = "EnumUprightInvertedInverted"
	}
	return
}

func (enumuprightinverted EnumUprightInverted) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumUprightInvertedUpright")
	res = append(res, "EnumUprightInvertedInverted")

	return
}

func (enumuprightinverted EnumUprightInverted) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upright")
	res = append(res, "inverted")

	return
}

// Utility function for EnumValign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumvalign EnumValign) ToString() (res string) {

	// migration of former implementation of enum
	switch enumvalign {
	// insertion code per enum code
	case EnumValignTop:
		res = "top"
	case EnumValignMiddle:
		res = "middle"
	case EnumValignBottom:
		res = "bottom"
	case EnumValignBaseline:
		res = "baseline"
	}
	return
}

func (enumvalign *EnumValign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enumvalign = EnumValignTop
		return
	case "middle":
		*enumvalign = EnumValignMiddle
		return
	case "bottom":
		*enumvalign = EnumValignBottom
		return
	case "baseline":
		*enumvalign = EnumValignBaseline
		return
	default:
		return errUnkownEnum
	}
}

func (enumvalign *EnumValign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumValignTop":
		*enumvalign = EnumValignTop
	case "EnumValignMiddle":
		*enumvalign = EnumValignMiddle
	case "EnumValignBottom":
		*enumvalign = EnumValignBottom
	case "EnumValignBaseline":
		*enumvalign = EnumValignBaseline
	default:
		err = errUnkownEnum
	}
	return
}

func (enumvalign *EnumValign) ToCodeString() (res string) {

	switch *enumvalign {
	// insertion code per enum code
	case EnumValignTop:
		res = "EnumValignTop"
	case EnumValignMiddle:
		res = "EnumValignMiddle"
	case EnumValignBottom:
		res = "EnumValignBottom"
	case EnumValignBaseline:
		res = "EnumValignBaseline"
	}
	return
}

func (enumvalign EnumValign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumValignTop")
	res = append(res, "EnumValignMiddle")
	res = append(res, "EnumValignBottom")
	res = append(res, "EnumValignBaseline")

	return
}

func (enumvalign EnumValign) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")
	res = append(res, "baseline")

	return
}

// Utility function for EnumValignImage
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumvalignimage EnumValignImage) ToString() (res string) {

	// migration of former implementation of enum
	switch enumvalignimage {
	// insertion code per enum code
	case EnumValignImageTop:
		res = "top"
	case EnumValignImageMiddle:
		res = "middle"
	case EnumValignImageBottom:
		res = "bottom"
	}
	return
}

func (enumvalignimage *EnumValignImage) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enumvalignimage = EnumValignImageTop
		return
	case "middle":
		*enumvalignimage = EnumValignImageMiddle
		return
	case "bottom":
		*enumvalignimage = EnumValignImageBottom
		return
	default:
		return errUnkownEnum
	}
}

func (enumvalignimage *EnumValignImage) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumValignImageTop":
		*enumvalignimage = EnumValignImageTop
	case "EnumValignImageMiddle":
		*enumvalignimage = EnumValignImageMiddle
	case "EnumValignImageBottom":
		*enumvalignimage = EnumValignImageBottom
	default:
		err = errUnkownEnum
	}
	return
}

func (enumvalignimage *EnumValignImage) ToCodeString() (res string) {

	switch *enumvalignimage {
	// insertion code per enum code
	case EnumValignImageTop:
		res = "EnumValignImageTop"
	case EnumValignImageMiddle:
		res = "EnumValignImageMiddle"
	case EnumValignImageBottom:
		res = "EnumValignImageBottom"
	}
	return
}

func (enumvalignimage EnumValignImage) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumValignImageTop")
	res = append(res, "EnumValignImageMiddle")
	res = append(res, "EnumValignImageBottom")

	return
}

func (enumvalignimage EnumValignImage) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")

	return
}

// Utility function for EnumYesNo
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumyesno EnumYesNo) ToString() (res string) {

	// migration of former implementation of enum
	switch enumyesno {
	// insertion code per enum code
	case EnumYesNoYes:
		res = "yes"
	case EnumYesNoNo:
		res = "no"
	}
	return
}

func (enumyesno *EnumYesNo) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enumyesno = EnumYesNoYes
		return
	case "no":
		*enumyesno = EnumYesNoNo
		return
	default:
		return errUnkownEnum
	}
}

func (enumyesno *EnumYesNo) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumYesNoYes":
		*enumyesno = EnumYesNoYes
	case "EnumYesNoNo":
		*enumyesno = EnumYesNoNo
	default:
		err = errUnkownEnum
	}
	return
}

func (enumyesno *EnumYesNo) ToCodeString() (res string) {

	switch *enumyesno {
	// insertion code per enum code
	case EnumYesNoYes:
		res = "EnumYesNoYes"
	case EnumYesNoNo:
		res = "EnumYesNoNo"
	}
	return
}

func (enumyesno EnumYesNo) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumYesNoYes")
	res = append(res, "EnumYesNoNo")

	return
}

func (enumyesno EnumYesNo) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")

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
