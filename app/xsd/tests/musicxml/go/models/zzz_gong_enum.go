// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Enum_Above_below
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_above_below Enum_Above_below) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_above_below {
	// insertion code per enum code
	case Enum_Above_below_Above:
		res = "above"
	case Enum_Above_below_Below:
		res = "below"
	}
	return
}

func (enum_above_below *Enum_Above_below) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "above":
		*enum_above_below = Enum_Above_below_Above
		return
	case "below":
		*enum_above_below = Enum_Above_below_Below
		return
	default:
		return errUnkownEnum
	}
}

func (enum_above_below *Enum_Above_below) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Above_below_Above":
		*enum_above_below = Enum_Above_below_Above
	case "Enum_Above_below_Below":
		*enum_above_below = Enum_Above_below_Below
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_above_below *Enum_Above_below) ToCodeString() (res string) {

	switch *enum_above_below {
	// insertion code per enum code
	case Enum_Above_below_Above:
		res = "Enum_Above_below_Above"
	case Enum_Above_below_Below:
		res = "Enum_Above_below_Below"
	}
	return
}

func (enum_above_below Enum_Above_below) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Above_below_Above")
	res = append(res, "Enum_Above_below_Below")

	return
}

func (enum_above_below Enum_Above_below) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "above")
	res = append(res, "below")

	return
}

// Utility function for Enum_Accidental_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_accidental_value Enum_Accidental_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_accidental_value {
	// insertion code per enum code
	case Enum_Accidental_value_Sharp:
		res = "sharp"
	case Enum_Accidental_value_Natural:
		res = "natural"
	case Enum_Accidental_value_Flat:
		res = "flat"
	case Enum_Accidental_value_Double_sharp:
		res = "double-sharp"
	case Enum_Accidental_value_Sharp_sharp:
		res = "sharp-sharp"
	case Enum_Accidental_value_Flat_flat:
		res = "flat-flat"
	case Enum_Accidental_value_Natural_sharp:
		res = "natural-sharp"
	case Enum_Accidental_value_Natural_flat:
		res = "natural-flat"
	case Enum_Accidental_value_Quarter_flat:
		res = "quarter-flat"
	case Enum_Accidental_value_Quarter_sharp:
		res = "quarter-sharp"
	case Enum_Accidental_value_Three_quarters_flat:
		res = "three-quarters-flat"
	case Enum_Accidental_value_Three_quarters_sharp:
		res = "three-quarters-sharp"
	case Enum_Accidental_value_Sharp_down:
		res = "sharp-down"
	case Enum_Accidental_value_Sharp_up:
		res = "sharp-up"
	case Enum_Accidental_value_Natural_down:
		res = "natural-down"
	case Enum_Accidental_value_Natural_up:
		res = "natural-up"
	case Enum_Accidental_value_Flat_down:
		res = "flat-down"
	case Enum_Accidental_value_Flat_up:
		res = "flat-up"
	case Enum_Accidental_value_Double_sharp_down:
		res = "double-sharp-down"
	case Enum_Accidental_value_Double_sharp_up:
		res = "double-sharp-up"
	case Enum_Accidental_value_Flat_flat_down:
		res = "flat-flat-down"
	case Enum_Accidental_value_Flat_flat_up:
		res = "flat-flat-up"
	case Enum_Accidental_value_Arrow_down:
		res = "arrow-down"
	case Enum_Accidental_value_Arrow_up:
		res = "arrow-up"
	case Enum_Accidental_value_Triple_sharp:
		res = "triple-sharp"
	case Enum_Accidental_value_Triple_flat:
		res = "triple-flat"
	case Enum_Accidental_value_Slash_quarter_sharp:
		res = "slash-quarter-sharp"
	case Enum_Accidental_value_Slash_sharp:
		res = "slash-sharp"
	case Enum_Accidental_value_Slash_flat:
		res = "slash-flat"
	case Enum_Accidental_value_Double_slash_flat:
		res = "double-slash-flat"
	case Enum_Accidental_value_Sharp_1:
		res = "sharp-1"
	case Enum_Accidental_value_Sharp_2:
		res = "sharp-2"
	case Enum_Accidental_value_Sharp_3:
		res = "sharp-3"
	case Enum_Accidental_value_Sharp_5:
		res = "sharp-5"
	case Enum_Accidental_value_Flat_1:
		res = "flat-1"
	case Enum_Accidental_value_Flat_2:
		res = "flat-2"
	case Enum_Accidental_value_Flat_3:
		res = "flat-3"
	case Enum_Accidental_value_Flat_4:
		res = "flat-4"
	case Enum_Accidental_value_Sori:
		res = "sori"
	case Enum_Accidental_value_Koron:
		res = "koron"
	case Enum_Accidental_value_Other:
		res = "other"
	}
	return
}

func (enum_accidental_value *Enum_Accidental_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sharp":
		*enum_accidental_value = Enum_Accidental_value_Sharp
		return
	case "natural":
		*enum_accidental_value = Enum_Accidental_value_Natural
		return
	case "flat":
		*enum_accidental_value = Enum_Accidental_value_Flat
		return
	case "double-sharp":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp
		return
	case "sharp-sharp":
		*enum_accidental_value = Enum_Accidental_value_Sharp_sharp
		return
	case "flat-flat":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat
		return
	case "natural-sharp":
		*enum_accidental_value = Enum_Accidental_value_Natural_sharp
		return
	case "natural-flat":
		*enum_accidental_value = Enum_Accidental_value_Natural_flat
		return
	case "quarter-flat":
		*enum_accidental_value = Enum_Accidental_value_Quarter_flat
		return
	case "quarter-sharp":
		*enum_accidental_value = Enum_Accidental_value_Quarter_sharp
		return
	case "three-quarters-flat":
		*enum_accidental_value = Enum_Accidental_value_Three_quarters_flat
		return
	case "three-quarters-sharp":
		*enum_accidental_value = Enum_Accidental_value_Three_quarters_sharp
		return
	case "sharp-down":
		*enum_accidental_value = Enum_Accidental_value_Sharp_down
		return
	case "sharp-up":
		*enum_accidental_value = Enum_Accidental_value_Sharp_up
		return
	case "natural-down":
		*enum_accidental_value = Enum_Accidental_value_Natural_down
		return
	case "natural-up":
		*enum_accidental_value = Enum_Accidental_value_Natural_up
		return
	case "flat-down":
		*enum_accidental_value = Enum_Accidental_value_Flat_down
		return
	case "flat-up":
		*enum_accidental_value = Enum_Accidental_value_Flat_up
		return
	case "double-sharp-down":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp_down
		return
	case "double-sharp-up":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp_up
		return
	case "flat-flat-down":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat_down
		return
	case "flat-flat-up":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat_up
		return
	case "arrow-down":
		*enum_accidental_value = Enum_Accidental_value_Arrow_down
		return
	case "arrow-up":
		*enum_accidental_value = Enum_Accidental_value_Arrow_up
		return
	case "triple-sharp":
		*enum_accidental_value = Enum_Accidental_value_Triple_sharp
		return
	case "triple-flat":
		*enum_accidental_value = Enum_Accidental_value_Triple_flat
		return
	case "slash-quarter-sharp":
		*enum_accidental_value = Enum_Accidental_value_Slash_quarter_sharp
		return
	case "slash-sharp":
		*enum_accidental_value = Enum_Accidental_value_Slash_sharp
		return
	case "slash-flat":
		*enum_accidental_value = Enum_Accidental_value_Slash_flat
		return
	case "double-slash-flat":
		*enum_accidental_value = Enum_Accidental_value_Double_slash_flat
		return
	case "sharp-1":
		*enum_accidental_value = Enum_Accidental_value_Sharp_1
		return
	case "sharp-2":
		*enum_accidental_value = Enum_Accidental_value_Sharp_2
		return
	case "sharp-3":
		*enum_accidental_value = Enum_Accidental_value_Sharp_3
		return
	case "sharp-5":
		*enum_accidental_value = Enum_Accidental_value_Sharp_5
		return
	case "flat-1":
		*enum_accidental_value = Enum_Accidental_value_Flat_1
		return
	case "flat-2":
		*enum_accidental_value = Enum_Accidental_value_Flat_2
		return
	case "flat-3":
		*enum_accidental_value = Enum_Accidental_value_Flat_3
		return
	case "flat-4":
		*enum_accidental_value = Enum_Accidental_value_Flat_4
		return
	case "sori":
		*enum_accidental_value = Enum_Accidental_value_Sori
		return
	case "koron":
		*enum_accidental_value = Enum_Accidental_value_Koron
		return
	case "other":
		*enum_accidental_value = Enum_Accidental_value_Other
		return
	default:
		return errUnkownEnum
	}
}

func (enum_accidental_value *Enum_Accidental_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Accidental_value_Sharp":
		*enum_accidental_value = Enum_Accidental_value_Sharp
	case "Enum_Accidental_value_Natural":
		*enum_accidental_value = Enum_Accidental_value_Natural
	case "Enum_Accidental_value_Flat":
		*enum_accidental_value = Enum_Accidental_value_Flat
	case "Enum_Accidental_value_Double_sharp":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp
	case "Enum_Accidental_value_Sharp_sharp":
		*enum_accidental_value = Enum_Accidental_value_Sharp_sharp
	case "Enum_Accidental_value_Flat_flat":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat
	case "Enum_Accidental_value_Natural_sharp":
		*enum_accidental_value = Enum_Accidental_value_Natural_sharp
	case "Enum_Accidental_value_Natural_flat":
		*enum_accidental_value = Enum_Accidental_value_Natural_flat
	case "Enum_Accidental_value_Quarter_flat":
		*enum_accidental_value = Enum_Accidental_value_Quarter_flat
	case "Enum_Accidental_value_Quarter_sharp":
		*enum_accidental_value = Enum_Accidental_value_Quarter_sharp
	case "Enum_Accidental_value_Three_quarters_flat":
		*enum_accidental_value = Enum_Accidental_value_Three_quarters_flat
	case "Enum_Accidental_value_Three_quarters_sharp":
		*enum_accidental_value = Enum_Accidental_value_Three_quarters_sharp
	case "Enum_Accidental_value_Sharp_down":
		*enum_accidental_value = Enum_Accidental_value_Sharp_down
	case "Enum_Accidental_value_Sharp_up":
		*enum_accidental_value = Enum_Accidental_value_Sharp_up
	case "Enum_Accidental_value_Natural_down":
		*enum_accidental_value = Enum_Accidental_value_Natural_down
	case "Enum_Accidental_value_Natural_up":
		*enum_accidental_value = Enum_Accidental_value_Natural_up
	case "Enum_Accidental_value_Flat_down":
		*enum_accidental_value = Enum_Accidental_value_Flat_down
	case "Enum_Accidental_value_Flat_up":
		*enum_accidental_value = Enum_Accidental_value_Flat_up
	case "Enum_Accidental_value_Double_sharp_down":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp_down
	case "Enum_Accidental_value_Double_sharp_up":
		*enum_accidental_value = Enum_Accidental_value_Double_sharp_up
	case "Enum_Accidental_value_Flat_flat_down":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat_down
	case "Enum_Accidental_value_Flat_flat_up":
		*enum_accidental_value = Enum_Accidental_value_Flat_flat_up
	case "Enum_Accidental_value_Arrow_down":
		*enum_accidental_value = Enum_Accidental_value_Arrow_down
	case "Enum_Accidental_value_Arrow_up":
		*enum_accidental_value = Enum_Accidental_value_Arrow_up
	case "Enum_Accidental_value_Triple_sharp":
		*enum_accidental_value = Enum_Accidental_value_Triple_sharp
	case "Enum_Accidental_value_Triple_flat":
		*enum_accidental_value = Enum_Accidental_value_Triple_flat
	case "Enum_Accidental_value_Slash_quarter_sharp":
		*enum_accidental_value = Enum_Accidental_value_Slash_quarter_sharp
	case "Enum_Accidental_value_Slash_sharp":
		*enum_accidental_value = Enum_Accidental_value_Slash_sharp
	case "Enum_Accidental_value_Slash_flat":
		*enum_accidental_value = Enum_Accidental_value_Slash_flat
	case "Enum_Accidental_value_Double_slash_flat":
		*enum_accidental_value = Enum_Accidental_value_Double_slash_flat
	case "Enum_Accidental_value_Sharp_1":
		*enum_accidental_value = Enum_Accidental_value_Sharp_1
	case "Enum_Accidental_value_Sharp_2":
		*enum_accidental_value = Enum_Accidental_value_Sharp_2
	case "Enum_Accidental_value_Sharp_3":
		*enum_accidental_value = Enum_Accidental_value_Sharp_3
	case "Enum_Accidental_value_Sharp_5":
		*enum_accidental_value = Enum_Accidental_value_Sharp_5
	case "Enum_Accidental_value_Flat_1":
		*enum_accidental_value = Enum_Accidental_value_Flat_1
	case "Enum_Accidental_value_Flat_2":
		*enum_accidental_value = Enum_Accidental_value_Flat_2
	case "Enum_Accidental_value_Flat_3":
		*enum_accidental_value = Enum_Accidental_value_Flat_3
	case "Enum_Accidental_value_Flat_4":
		*enum_accidental_value = Enum_Accidental_value_Flat_4
	case "Enum_Accidental_value_Sori":
		*enum_accidental_value = Enum_Accidental_value_Sori
	case "Enum_Accidental_value_Koron":
		*enum_accidental_value = Enum_Accidental_value_Koron
	case "Enum_Accidental_value_Other":
		*enum_accidental_value = Enum_Accidental_value_Other
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_accidental_value *Enum_Accidental_value) ToCodeString() (res string) {

	switch *enum_accidental_value {
	// insertion code per enum code
	case Enum_Accidental_value_Sharp:
		res = "Enum_Accidental_value_Sharp"
	case Enum_Accidental_value_Natural:
		res = "Enum_Accidental_value_Natural"
	case Enum_Accidental_value_Flat:
		res = "Enum_Accidental_value_Flat"
	case Enum_Accidental_value_Double_sharp:
		res = "Enum_Accidental_value_Double_sharp"
	case Enum_Accidental_value_Sharp_sharp:
		res = "Enum_Accidental_value_Sharp_sharp"
	case Enum_Accidental_value_Flat_flat:
		res = "Enum_Accidental_value_Flat_flat"
	case Enum_Accidental_value_Natural_sharp:
		res = "Enum_Accidental_value_Natural_sharp"
	case Enum_Accidental_value_Natural_flat:
		res = "Enum_Accidental_value_Natural_flat"
	case Enum_Accidental_value_Quarter_flat:
		res = "Enum_Accidental_value_Quarter_flat"
	case Enum_Accidental_value_Quarter_sharp:
		res = "Enum_Accidental_value_Quarter_sharp"
	case Enum_Accidental_value_Three_quarters_flat:
		res = "Enum_Accidental_value_Three_quarters_flat"
	case Enum_Accidental_value_Three_quarters_sharp:
		res = "Enum_Accidental_value_Three_quarters_sharp"
	case Enum_Accidental_value_Sharp_down:
		res = "Enum_Accidental_value_Sharp_down"
	case Enum_Accidental_value_Sharp_up:
		res = "Enum_Accidental_value_Sharp_up"
	case Enum_Accidental_value_Natural_down:
		res = "Enum_Accidental_value_Natural_down"
	case Enum_Accidental_value_Natural_up:
		res = "Enum_Accidental_value_Natural_up"
	case Enum_Accidental_value_Flat_down:
		res = "Enum_Accidental_value_Flat_down"
	case Enum_Accidental_value_Flat_up:
		res = "Enum_Accidental_value_Flat_up"
	case Enum_Accidental_value_Double_sharp_down:
		res = "Enum_Accidental_value_Double_sharp_down"
	case Enum_Accidental_value_Double_sharp_up:
		res = "Enum_Accidental_value_Double_sharp_up"
	case Enum_Accidental_value_Flat_flat_down:
		res = "Enum_Accidental_value_Flat_flat_down"
	case Enum_Accidental_value_Flat_flat_up:
		res = "Enum_Accidental_value_Flat_flat_up"
	case Enum_Accidental_value_Arrow_down:
		res = "Enum_Accidental_value_Arrow_down"
	case Enum_Accidental_value_Arrow_up:
		res = "Enum_Accidental_value_Arrow_up"
	case Enum_Accidental_value_Triple_sharp:
		res = "Enum_Accidental_value_Triple_sharp"
	case Enum_Accidental_value_Triple_flat:
		res = "Enum_Accidental_value_Triple_flat"
	case Enum_Accidental_value_Slash_quarter_sharp:
		res = "Enum_Accidental_value_Slash_quarter_sharp"
	case Enum_Accidental_value_Slash_sharp:
		res = "Enum_Accidental_value_Slash_sharp"
	case Enum_Accidental_value_Slash_flat:
		res = "Enum_Accidental_value_Slash_flat"
	case Enum_Accidental_value_Double_slash_flat:
		res = "Enum_Accidental_value_Double_slash_flat"
	case Enum_Accidental_value_Sharp_1:
		res = "Enum_Accidental_value_Sharp_1"
	case Enum_Accidental_value_Sharp_2:
		res = "Enum_Accidental_value_Sharp_2"
	case Enum_Accidental_value_Sharp_3:
		res = "Enum_Accidental_value_Sharp_3"
	case Enum_Accidental_value_Sharp_5:
		res = "Enum_Accidental_value_Sharp_5"
	case Enum_Accidental_value_Flat_1:
		res = "Enum_Accidental_value_Flat_1"
	case Enum_Accidental_value_Flat_2:
		res = "Enum_Accidental_value_Flat_2"
	case Enum_Accidental_value_Flat_3:
		res = "Enum_Accidental_value_Flat_3"
	case Enum_Accidental_value_Flat_4:
		res = "Enum_Accidental_value_Flat_4"
	case Enum_Accidental_value_Sori:
		res = "Enum_Accidental_value_Sori"
	case Enum_Accidental_value_Koron:
		res = "Enum_Accidental_value_Koron"
	case Enum_Accidental_value_Other:
		res = "Enum_Accidental_value_Other"
	}
	return
}

func (enum_accidental_value Enum_Accidental_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Accidental_value_Sharp")
	res = append(res, "Enum_Accidental_value_Natural")
	res = append(res, "Enum_Accidental_value_Flat")
	res = append(res, "Enum_Accidental_value_Double_sharp")
	res = append(res, "Enum_Accidental_value_Sharp_sharp")
	res = append(res, "Enum_Accidental_value_Flat_flat")
	res = append(res, "Enum_Accidental_value_Natural_sharp")
	res = append(res, "Enum_Accidental_value_Natural_flat")
	res = append(res, "Enum_Accidental_value_Quarter_flat")
	res = append(res, "Enum_Accidental_value_Quarter_sharp")
	res = append(res, "Enum_Accidental_value_Three_quarters_flat")
	res = append(res, "Enum_Accidental_value_Three_quarters_sharp")
	res = append(res, "Enum_Accidental_value_Sharp_down")
	res = append(res, "Enum_Accidental_value_Sharp_up")
	res = append(res, "Enum_Accidental_value_Natural_down")
	res = append(res, "Enum_Accidental_value_Natural_up")
	res = append(res, "Enum_Accidental_value_Flat_down")
	res = append(res, "Enum_Accidental_value_Flat_up")
	res = append(res, "Enum_Accidental_value_Double_sharp_down")
	res = append(res, "Enum_Accidental_value_Double_sharp_up")
	res = append(res, "Enum_Accidental_value_Flat_flat_down")
	res = append(res, "Enum_Accidental_value_Flat_flat_up")
	res = append(res, "Enum_Accidental_value_Arrow_down")
	res = append(res, "Enum_Accidental_value_Arrow_up")
	res = append(res, "Enum_Accidental_value_Triple_sharp")
	res = append(res, "Enum_Accidental_value_Triple_flat")
	res = append(res, "Enum_Accidental_value_Slash_quarter_sharp")
	res = append(res, "Enum_Accidental_value_Slash_sharp")
	res = append(res, "Enum_Accidental_value_Slash_flat")
	res = append(res, "Enum_Accidental_value_Double_slash_flat")
	res = append(res, "Enum_Accidental_value_Sharp_1")
	res = append(res, "Enum_Accidental_value_Sharp_2")
	res = append(res, "Enum_Accidental_value_Sharp_3")
	res = append(res, "Enum_Accidental_value_Sharp_5")
	res = append(res, "Enum_Accidental_value_Flat_1")
	res = append(res, "Enum_Accidental_value_Flat_2")
	res = append(res, "Enum_Accidental_value_Flat_3")
	res = append(res, "Enum_Accidental_value_Flat_4")
	res = append(res, "Enum_Accidental_value_Sori")
	res = append(res, "Enum_Accidental_value_Koron")
	res = append(res, "Enum_Accidental_value_Other")

	return
}

func (enum_accidental_value Enum_Accidental_value) CodeValues() (res []string) {

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

// Utility function for Enum_Arrow_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_arrow_direction Enum_Arrow_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_arrow_direction {
	// insertion code per enum code
	case Enum_Arrow_direction_Left:
		res = "left"
	case Enum_Arrow_direction_Up:
		res = "up"
	case Enum_Arrow_direction_Right:
		res = "right"
	case Enum_Arrow_direction_Down:
		res = "down"
	case Enum_Arrow_direction_Northwest:
		res = "northwest"
	case Enum_Arrow_direction_Northeast:
		res = "northeast"
	case Enum_Arrow_direction_Southeast:
		res = "southeast"
	case Enum_Arrow_direction_Southwest:
		res = "southwest"
	case Enum_Arrow_direction_Left_right:
		res = "left right"
	case Enum_Arrow_direction_Up_down:
		res = "up down"
	case Enum_Arrow_direction_Northwest_southeast:
		res = "northwest southeast"
	case Enum_Arrow_direction_Northeast_southwest:
		res = "northeast southwest"
	case Enum_Arrow_direction_Other:
		res = "other"
	}
	return
}

func (enum_arrow_direction *Enum_Arrow_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enum_arrow_direction = Enum_Arrow_direction_Left
		return
	case "up":
		*enum_arrow_direction = Enum_Arrow_direction_Up
		return
	case "right":
		*enum_arrow_direction = Enum_Arrow_direction_Right
		return
	case "down":
		*enum_arrow_direction = Enum_Arrow_direction_Down
		return
	case "northwest":
		*enum_arrow_direction = Enum_Arrow_direction_Northwest
		return
	case "northeast":
		*enum_arrow_direction = Enum_Arrow_direction_Northeast
		return
	case "southeast":
		*enum_arrow_direction = Enum_Arrow_direction_Southeast
		return
	case "southwest":
		*enum_arrow_direction = Enum_Arrow_direction_Southwest
		return
	case "left right":
		*enum_arrow_direction = Enum_Arrow_direction_Left_right
		return
	case "up down":
		*enum_arrow_direction = Enum_Arrow_direction_Up_down
		return
	case "northwest southeast":
		*enum_arrow_direction = Enum_Arrow_direction_Northwest_southeast
		return
	case "northeast southwest":
		*enum_arrow_direction = Enum_Arrow_direction_Northeast_southwest
		return
	case "other":
		*enum_arrow_direction = Enum_Arrow_direction_Other
		return
	default:
		return errUnkownEnum
	}
}

func (enum_arrow_direction *Enum_Arrow_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Arrow_direction_Left":
		*enum_arrow_direction = Enum_Arrow_direction_Left
	case "Enum_Arrow_direction_Up":
		*enum_arrow_direction = Enum_Arrow_direction_Up
	case "Enum_Arrow_direction_Right":
		*enum_arrow_direction = Enum_Arrow_direction_Right
	case "Enum_Arrow_direction_Down":
		*enum_arrow_direction = Enum_Arrow_direction_Down
	case "Enum_Arrow_direction_Northwest":
		*enum_arrow_direction = Enum_Arrow_direction_Northwest
	case "Enum_Arrow_direction_Northeast":
		*enum_arrow_direction = Enum_Arrow_direction_Northeast
	case "Enum_Arrow_direction_Southeast":
		*enum_arrow_direction = Enum_Arrow_direction_Southeast
	case "Enum_Arrow_direction_Southwest":
		*enum_arrow_direction = Enum_Arrow_direction_Southwest
	case "Enum_Arrow_direction_Left_right":
		*enum_arrow_direction = Enum_Arrow_direction_Left_right
	case "Enum_Arrow_direction_Up_down":
		*enum_arrow_direction = Enum_Arrow_direction_Up_down
	case "Enum_Arrow_direction_Northwest_southeast":
		*enum_arrow_direction = Enum_Arrow_direction_Northwest_southeast
	case "Enum_Arrow_direction_Northeast_southwest":
		*enum_arrow_direction = Enum_Arrow_direction_Northeast_southwest
	case "Enum_Arrow_direction_Other":
		*enum_arrow_direction = Enum_Arrow_direction_Other
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_arrow_direction *Enum_Arrow_direction) ToCodeString() (res string) {

	switch *enum_arrow_direction {
	// insertion code per enum code
	case Enum_Arrow_direction_Left:
		res = "Enum_Arrow_direction_Left"
	case Enum_Arrow_direction_Up:
		res = "Enum_Arrow_direction_Up"
	case Enum_Arrow_direction_Right:
		res = "Enum_Arrow_direction_Right"
	case Enum_Arrow_direction_Down:
		res = "Enum_Arrow_direction_Down"
	case Enum_Arrow_direction_Northwest:
		res = "Enum_Arrow_direction_Northwest"
	case Enum_Arrow_direction_Northeast:
		res = "Enum_Arrow_direction_Northeast"
	case Enum_Arrow_direction_Southeast:
		res = "Enum_Arrow_direction_Southeast"
	case Enum_Arrow_direction_Southwest:
		res = "Enum_Arrow_direction_Southwest"
	case Enum_Arrow_direction_Left_right:
		res = "Enum_Arrow_direction_Left_right"
	case Enum_Arrow_direction_Up_down:
		res = "Enum_Arrow_direction_Up_down"
	case Enum_Arrow_direction_Northwest_southeast:
		res = "Enum_Arrow_direction_Northwest_southeast"
	case Enum_Arrow_direction_Northeast_southwest:
		res = "Enum_Arrow_direction_Northeast_southwest"
	case Enum_Arrow_direction_Other:
		res = "Enum_Arrow_direction_Other"
	}
	return
}

func (enum_arrow_direction Enum_Arrow_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Arrow_direction_Left")
	res = append(res, "Enum_Arrow_direction_Up")
	res = append(res, "Enum_Arrow_direction_Right")
	res = append(res, "Enum_Arrow_direction_Down")
	res = append(res, "Enum_Arrow_direction_Northwest")
	res = append(res, "Enum_Arrow_direction_Northeast")
	res = append(res, "Enum_Arrow_direction_Southeast")
	res = append(res, "Enum_Arrow_direction_Southwest")
	res = append(res, "Enum_Arrow_direction_Left_right")
	res = append(res, "Enum_Arrow_direction_Up_down")
	res = append(res, "Enum_Arrow_direction_Northwest_southeast")
	res = append(res, "Enum_Arrow_direction_Northeast_southwest")
	res = append(res, "Enum_Arrow_direction_Other")

	return
}

func (enum_arrow_direction Enum_Arrow_direction) CodeValues() (res []string) {

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

// Utility function for Enum_Arrow_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_arrow_style Enum_Arrow_style) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_arrow_style {
	// insertion code per enum code
	case Enum_Arrow_style_Single:
		res = "single"
	case Enum_Arrow_style_Double:
		res = "double"
	case Enum_Arrow_style_Filled:
		res = "filled"
	case Enum_Arrow_style_Hollow:
		res = "hollow"
	case Enum_Arrow_style_Paired:
		res = "paired"
	case Enum_Arrow_style_Combined:
		res = "combined"
	case Enum_Arrow_style_Other:
		res = "other"
	}
	return
}

func (enum_arrow_style *Enum_Arrow_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*enum_arrow_style = Enum_Arrow_style_Single
		return
	case "double":
		*enum_arrow_style = Enum_Arrow_style_Double
		return
	case "filled":
		*enum_arrow_style = Enum_Arrow_style_Filled
		return
	case "hollow":
		*enum_arrow_style = Enum_Arrow_style_Hollow
		return
	case "paired":
		*enum_arrow_style = Enum_Arrow_style_Paired
		return
	case "combined":
		*enum_arrow_style = Enum_Arrow_style_Combined
		return
	case "other":
		*enum_arrow_style = Enum_Arrow_style_Other
		return
	default:
		return errUnkownEnum
	}
}

func (enum_arrow_style *Enum_Arrow_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Arrow_style_Single":
		*enum_arrow_style = Enum_Arrow_style_Single
	case "Enum_Arrow_style_Double":
		*enum_arrow_style = Enum_Arrow_style_Double
	case "Enum_Arrow_style_Filled":
		*enum_arrow_style = Enum_Arrow_style_Filled
	case "Enum_Arrow_style_Hollow":
		*enum_arrow_style = Enum_Arrow_style_Hollow
	case "Enum_Arrow_style_Paired":
		*enum_arrow_style = Enum_Arrow_style_Paired
	case "Enum_Arrow_style_Combined":
		*enum_arrow_style = Enum_Arrow_style_Combined
	case "Enum_Arrow_style_Other":
		*enum_arrow_style = Enum_Arrow_style_Other
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_arrow_style *Enum_Arrow_style) ToCodeString() (res string) {

	switch *enum_arrow_style {
	// insertion code per enum code
	case Enum_Arrow_style_Single:
		res = "Enum_Arrow_style_Single"
	case Enum_Arrow_style_Double:
		res = "Enum_Arrow_style_Double"
	case Enum_Arrow_style_Filled:
		res = "Enum_Arrow_style_Filled"
	case Enum_Arrow_style_Hollow:
		res = "Enum_Arrow_style_Hollow"
	case Enum_Arrow_style_Paired:
		res = "Enum_Arrow_style_Paired"
	case Enum_Arrow_style_Combined:
		res = "Enum_Arrow_style_Combined"
	case Enum_Arrow_style_Other:
		res = "Enum_Arrow_style_Other"
	}
	return
}

func (enum_arrow_style Enum_Arrow_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Arrow_style_Single")
	res = append(res, "Enum_Arrow_style_Double")
	res = append(res, "Enum_Arrow_style_Filled")
	res = append(res, "Enum_Arrow_style_Hollow")
	res = append(res, "Enum_Arrow_style_Paired")
	res = append(res, "Enum_Arrow_style_Combined")
	res = append(res, "Enum_Arrow_style_Other")

	return
}

func (enum_arrow_style Enum_Arrow_style) CodeValues() (res []string) {

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

// Utility function for Enum_Backward_forward
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_backward_forward Enum_Backward_forward) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_backward_forward {
	// insertion code per enum code
	case Enum_Backward_forward_Backward:
		res = "backward"
	case Enum_Backward_forward_Forward:
		res = "forward"
	}
	return
}

func (enum_backward_forward *Enum_Backward_forward) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "backward":
		*enum_backward_forward = Enum_Backward_forward_Backward
		return
	case "forward":
		*enum_backward_forward = Enum_Backward_forward_Forward
		return
	default:
		return errUnkownEnum
	}
}

func (enum_backward_forward *Enum_Backward_forward) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Backward_forward_Backward":
		*enum_backward_forward = Enum_Backward_forward_Backward
	case "Enum_Backward_forward_Forward":
		*enum_backward_forward = Enum_Backward_forward_Forward
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_backward_forward *Enum_Backward_forward) ToCodeString() (res string) {

	switch *enum_backward_forward {
	// insertion code per enum code
	case Enum_Backward_forward_Backward:
		res = "Enum_Backward_forward_Backward"
	case Enum_Backward_forward_Forward:
		res = "Enum_Backward_forward_Forward"
	}
	return
}

func (enum_backward_forward Enum_Backward_forward) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Backward_forward_Backward")
	res = append(res, "Enum_Backward_forward_Forward")

	return
}

func (enum_backward_forward Enum_Backward_forward) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "backward")
	res = append(res, "forward")

	return
}

// Utility function for Enum_Bar_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_bar_style Enum_Bar_style) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_bar_style {
	// insertion code per enum code
	case Enum_Bar_style_Regular:
		res = "regular"
	case Enum_Bar_style_Dotted:
		res = "dotted"
	case Enum_Bar_style_Dashed:
		res = "dashed"
	case Enum_Bar_style_Heavy:
		res = "heavy"
	case Enum_Bar_style_Light_light:
		res = "light-light"
	case Enum_Bar_style_Light_heavy:
		res = "light-heavy"
	case Enum_Bar_style_Heavy_light:
		res = "heavy-light"
	case Enum_Bar_style_Heavy_heavy:
		res = "heavy-heavy"
	case Enum_Bar_style_Tick:
		res = "tick"
	case Enum_Bar_style_Short:
		res = "short"
	case Enum_Bar_style_None:
		res = "none"
	}
	return
}

func (enum_bar_style *Enum_Bar_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "regular":
		*enum_bar_style = Enum_Bar_style_Regular
		return
	case "dotted":
		*enum_bar_style = Enum_Bar_style_Dotted
		return
	case "dashed":
		*enum_bar_style = Enum_Bar_style_Dashed
		return
	case "heavy":
		*enum_bar_style = Enum_Bar_style_Heavy
		return
	case "light-light":
		*enum_bar_style = Enum_Bar_style_Light_light
		return
	case "light-heavy":
		*enum_bar_style = Enum_Bar_style_Light_heavy
		return
	case "heavy-light":
		*enum_bar_style = Enum_Bar_style_Heavy_light
		return
	case "heavy-heavy":
		*enum_bar_style = Enum_Bar_style_Heavy_heavy
		return
	case "tick":
		*enum_bar_style = Enum_Bar_style_Tick
		return
	case "short":
		*enum_bar_style = Enum_Bar_style_Short
		return
	case "none":
		*enum_bar_style = Enum_Bar_style_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_bar_style *Enum_Bar_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Bar_style_Regular":
		*enum_bar_style = Enum_Bar_style_Regular
	case "Enum_Bar_style_Dotted":
		*enum_bar_style = Enum_Bar_style_Dotted
	case "Enum_Bar_style_Dashed":
		*enum_bar_style = Enum_Bar_style_Dashed
	case "Enum_Bar_style_Heavy":
		*enum_bar_style = Enum_Bar_style_Heavy
	case "Enum_Bar_style_Light_light":
		*enum_bar_style = Enum_Bar_style_Light_light
	case "Enum_Bar_style_Light_heavy":
		*enum_bar_style = Enum_Bar_style_Light_heavy
	case "Enum_Bar_style_Heavy_light":
		*enum_bar_style = Enum_Bar_style_Heavy_light
	case "Enum_Bar_style_Heavy_heavy":
		*enum_bar_style = Enum_Bar_style_Heavy_heavy
	case "Enum_Bar_style_Tick":
		*enum_bar_style = Enum_Bar_style_Tick
	case "Enum_Bar_style_Short":
		*enum_bar_style = Enum_Bar_style_Short
	case "Enum_Bar_style_None":
		*enum_bar_style = Enum_Bar_style_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_bar_style *Enum_Bar_style) ToCodeString() (res string) {

	switch *enum_bar_style {
	// insertion code per enum code
	case Enum_Bar_style_Regular:
		res = "Enum_Bar_style_Regular"
	case Enum_Bar_style_Dotted:
		res = "Enum_Bar_style_Dotted"
	case Enum_Bar_style_Dashed:
		res = "Enum_Bar_style_Dashed"
	case Enum_Bar_style_Heavy:
		res = "Enum_Bar_style_Heavy"
	case Enum_Bar_style_Light_light:
		res = "Enum_Bar_style_Light_light"
	case Enum_Bar_style_Light_heavy:
		res = "Enum_Bar_style_Light_heavy"
	case Enum_Bar_style_Heavy_light:
		res = "Enum_Bar_style_Heavy_light"
	case Enum_Bar_style_Heavy_heavy:
		res = "Enum_Bar_style_Heavy_heavy"
	case Enum_Bar_style_Tick:
		res = "Enum_Bar_style_Tick"
	case Enum_Bar_style_Short:
		res = "Enum_Bar_style_Short"
	case Enum_Bar_style_None:
		res = "Enum_Bar_style_None"
	}
	return
}

func (enum_bar_style Enum_Bar_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Bar_style_Regular")
	res = append(res, "Enum_Bar_style_Dotted")
	res = append(res, "Enum_Bar_style_Dashed")
	res = append(res, "Enum_Bar_style_Heavy")
	res = append(res, "Enum_Bar_style_Light_light")
	res = append(res, "Enum_Bar_style_Light_heavy")
	res = append(res, "Enum_Bar_style_Heavy_light")
	res = append(res, "Enum_Bar_style_Heavy_heavy")
	res = append(res, "Enum_Bar_style_Tick")
	res = append(res, "Enum_Bar_style_Short")
	res = append(res, "Enum_Bar_style_None")

	return
}

func (enum_bar_style Enum_Bar_style) CodeValues() (res []string) {

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

// Utility function for Enum_Beam_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_beam_value Enum_Beam_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_beam_value {
	// insertion code per enum code
	case Enum_Beam_value_Begin:
		res = "begin"
	case Enum_Beam_value_Continue:
		res = "continue"
	case Enum_Beam_value_End:
		res = "end"
	case Enum_Beam_value_Forward_hook:
		res = "forward hook"
	case Enum_Beam_value_Backward_hook:
		res = "backward hook"
	}
	return
}

func (enum_beam_value *Enum_Beam_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "begin":
		*enum_beam_value = Enum_Beam_value_Begin
		return
	case "continue":
		*enum_beam_value = Enum_Beam_value_Continue
		return
	case "end":
		*enum_beam_value = Enum_Beam_value_End
		return
	case "forward hook":
		*enum_beam_value = Enum_Beam_value_Forward_hook
		return
	case "backward hook":
		*enum_beam_value = Enum_Beam_value_Backward_hook
		return
	default:
		return errUnkownEnum
	}
}

func (enum_beam_value *Enum_Beam_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Beam_value_Begin":
		*enum_beam_value = Enum_Beam_value_Begin
	case "Enum_Beam_value_Continue":
		*enum_beam_value = Enum_Beam_value_Continue
	case "Enum_Beam_value_End":
		*enum_beam_value = Enum_Beam_value_End
	case "Enum_Beam_value_Forward_hook":
		*enum_beam_value = Enum_Beam_value_Forward_hook
	case "Enum_Beam_value_Backward_hook":
		*enum_beam_value = Enum_Beam_value_Backward_hook
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_beam_value *Enum_Beam_value) ToCodeString() (res string) {

	switch *enum_beam_value {
	// insertion code per enum code
	case Enum_Beam_value_Begin:
		res = "Enum_Beam_value_Begin"
	case Enum_Beam_value_Continue:
		res = "Enum_Beam_value_Continue"
	case Enum_Beam_value_End:
		res = "Enum_Beam_value_End"
	case Enum_Beam_value_Forward_hook:
		res = "Enum_Beam_value_Forward_hook"
	case Enum_Beam_value_Backward_hook:
		res = "Enum_Beam_value_Backward_hook"
	}
	return
}

func (enum_beam_value Enum_Beam_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Beam_value_Begin")
	res = append(res, "Enum_Beam_value_Continue")
	res = append(res, "Enum_Beam_value_End")
	res = append(res, "Enum_Beam_value_Forward_hook")
	res = append(res, "Enum_Beam_value_Backward_hook")

	return
}

func (enum_beam_value Enum_Beam_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "begin")
	res = append(res, "continue")
	res = append(res, "end")
	res = append(res, "forward hook")
	res = append(res, "backward hook")

	return
}

// Utility function for Enum_Beater_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_beater_value Enum_Beater_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_beater_value {
	// insertion code per enum code
	case Enum_Beater_value_Bow:
		res = "bow"
	case Enum_Beater_value_Chime_hammer:
		res = "chime hammer"
	case Enum_Beater_value_Coin:
		res = "coin"
	case Enum_Beater_value_Drum_stick:
		res = "drum stick"
	case Enum_Beater_value_Finger:
		res = "finger"
	case Enum_Beater_value_Fingernail:
		res = "fingernail"
	case Enum_Beater_value_Fist:
		res = "fist"
	case Enum_Beater_value_Guiro_scraper:
		res = "guiro scraper"
	case Enum_Beater_value_Hammer:
		res = "hammer"
	case Enum_Beater_value_Hand:
		res = "hand"
	case Enum_Beater_value_Jazz_stick:
		res = "jazz stick"
	case Enum_Beater_value_Knitting_needle:
		res = "knitting needle"
	case Enum_Beater_value_Metal_hammer:
		res = "metal hammer"
	case Enum_Beater_value_Slide_brush_on_gong:
		res = "slide brush on gong"
	case Enum_Beater_value_Snare_stick:
		res = "snare stick"
	case Enum_Beater_value_Spoon_mallet:
		res = "spoon mallet"
	case Enum_Beater_value_Superball:
		res = "superball"
	case Enum_Beater_value_Triangle_beater:
		res = "triangle beater"
	case Enum_Beater_value_Triangle_beater_plain:
		res = "triangle beater plain"
	case Enum_Beater_value_Wire_brush:
		res = "wire brush"
	}
	return
}

func (enum_beater_value *Enum_Beater_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bow":
		*enum_beater_value = Enum_Beater_value_Bow
		return
	case "chime hammer":
		*enum_beater_value = Enum_Beater_value_Chime_hammer
		return
	case "coin":
		*enum_beater_value = Enum_Beater_value_Coin
		return
	case "drum stick":
		*enum_beater_value = Enum_Beater_value_Drum_stick
		return
	case "finger":
		*enum_beater_value = Enum_Beater_value_Finger
		return
	case "fingernail":
		*enum_beater_value = Enum_Beater_value_Fingernail
		return
	case "fist":
		*enum_beater_value = Enum_Beater_value_Fist
		return
	case "guiro scraper":
		*enum_beater_value = Enum_Beater_value_Guiro_scraper
		return
	case "hammer":
		*enum_beater_value = Enum_Beater_value_Hammer
		return
	case "hand":
		*enum_beater_value = Enum_Beater_value_Hand
		return
	case "jazz stick":
		*enum_beater_value = Enum_Beater_value_Jazz_stick
		return
	case "knitting needle":
		*enum_beater_value = Enum_Beater_value_Knitting_needle
		return
	case "metal hammer":
		*enum_beater_value = Enum_Beater_value_Metal_hammer
		return
	case "slide brush on gong":
		*enum_beater_value = Enum_Beater_value_Slide_brush_on_gong
		return
	case "snare stick":
		*enum_beater_value = Enum_Beater_value_Snare_stick
		return
	case "spoon mallet":
		*enum_beater_value = Enum_Beater_value_Spoon_mallet
		return
	case "superball":
		*enum_beater_value = Enum_Beater_value_Superball
		return
	case "triangle beater":
		*enum_beater_value = Enum_Beater_value_Triangle_beater
		return
	case "triangle beater plain":
		*enum_beater_value = Enum_Beater_value_Triangle_beater_plain
		return
	case "wire brush":
		*enum_beater_value = Enum_Beater_value_Wire_brush
		return
	default:
		return errUnkownEnum
	}
}

func (enum_beater_value *Enum_Beater_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Beater_value_Bow":
		*enum_beater_value = Enum_Beater_value_Bow
	case "Enum_Beater_value_Chime_hammer":
		*enum_beater_value = Enum_Beater_value_Chime_hammer
	case "Enum_Beater_value_Coin":
		*enum_beater_value = Enum_Beater_value_Coin
	case "Enum_Beater_value_Drum_stick":
		*enum_beater_value = Enum_Beater_value_Drum_stick
	case "Enum_Beater_value_Finger":
		*enum_beater_value = Enum_Beater_value_Finger
	case "Enum_Beater_value_Fingernail":
		*enum_beater_value = Enum_Beater_value_Fingernail
	case "Enum_Beater_value_Fist":
		*enum_beater_value = Enum_Beater_value_Fist
	case "Enum_Beater_value_Guiro_scraper":
		*enum_beater_value = Enum_Beater_value_Guiro_scraper
	case "Enum_Beater_value_Hammer":
		*enum_beater_value = Enum_Beater_value_Hammer
	case "Enum_Beater_value_Hand":
		*enum_beater_value = Enum_Beater_value_Hand
	case "Enum_Beater_value_Jazz_stick":
		*enum_beater_value = Enum_Beater_value_Jazz_stick
	case "Enum_Beater_value_Knitting_needle":
		*enum_beater_value = Enum_Beater_value_Knitting_needle
	case "Enum_Beater_value_Metal_hammer":
		*enum_beater_value = Enum_Beater_value_Metal_hammer
	case "Enum_Beater_value_Slide_brush_on_gong":
		*enum_beater_value = Enum_Beater_value_Slide_brush_on_gong
	case "Enum_Beater_value_Snare_stick":
		*enum_beater_value = Enum_Beater_value_Snare_stick
	case "Enum_Beater_value_Spoon_mallet":
		*enum_beater_value = Enum_Beater_value_Spoon_mallet
	case "Enum_Beater_value_Superball":
		*enum_beater_value = Enum_Beater_value_Superball
	case "Enum_Beater_value_Triangle_beater":
		*enum_beater_value = Enum_Beater_value_Triangle_beater
	case "Enum_Beater_value_Triangle_beater_plain":
		*enum_beater_value = Enum_Beater_value_Triangle_beater_plain
	case "Enum_Beater_value_Wire_brush":
		*enum_beater_value = Enum_Beater_value_Wire_brush
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_beater_value *Enum_Beater_value) ToCodeString() (res string) {

	switch *enum_beater_value {
	// insertion code per enum code
	case Enum_Beater_value_Bow:
		res = "Enum_Beater_value_Bow"
	case Enum_Beater_value_Chime_hammer:
		res = "Enum_Beater_value_Chime_hammer"
	case Enum_Beater_value_Coin:
		res = "Enum_Beater_value_Coin"
	case Enum_Beater_value_Drum_stick:
		res = "Enum_Beater_value_Drum_stick"
	case Enum_Beater_value_Finger:
		res = "Enum_Beater_value_Finger"
	case Enum_Beater_value_Fingernail:
		res = "Enum_Beater_value_Fingernail"
	case Enum_Beater_value_Fist:
		res = "Enum_Beater_value_Fist"
	case Enum_Beater_value_Guiro_scraper:
		res = "Enum_Beater_value_Guiro_scraper"
	case Enum_Beater_value_Hammer:
		res = "Enum_Beater_value_Hammer"
	case Enum_Beater_value_Hand:
		res = "Enum_Beater_value_Hand"
	case Enum_Beater_value_Jazz_stick:
		res = "Enum_Beater_value_Jazz_stick"
	case Enum_Beater_value_Knitting_needle:
		res = "Enum_Beater_value_Knitting_needle"
	case Enum_Beater_value_Metal_hammer:
		res = "Enum_Beater_value_Metal_hammer"
	case Enum_Beater_value_Slide_brush_on_gong:
		res = "Enum_Beater_value_Slide_brush_on_gong"
	case Enum_Beater_value_Snare_stick:
		res = "Enum_Beater_value_Snare_stick"
	case Enum_Beater_value_Spoon_mallet:
		res = "Enum_Beater_value_Spoon_mallet"
	case Enum_Beater_value_Superball:
		res = "Enum_Beater_value_Superball"
	case Enum_Beater_value_Triangle_beater:
		res = "Enum_Beater_value_Triangle_beater"
	case Enum_Beater_value_Triangle_beater_plain:
		res = "Enum_Beater_value_Triangle_beater_plain"
	case Enum_Beater_value_Wire_brush:
		res = "Enum_Beater_value_Wire_brush"
	}
	return
}

func (enum_beater_value Enum_Beater_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Beater_value_Bow")
	res = append(res, "Enum_Beater_value_Chime_hammer")
	res = append(res, "Enum_Beater_value_Coin")
	res = append(res, "Enum_Beater_value_Drum_stick")
	res = append(res, "Enum_Beater_value_Finger")
	res = append(res, "Enum_Beater_value_Fingernail")
	res = append(res, "Enum_Beater_value_Fist")
	res = append(res, "Enum_Beater_value_Guiro_scraper")
	res = append(res, "Enum_Beater_value_Hammer")
	res = append(res, "Enum_Beater_value_Hand")
	res = append(res, "Enum_Beater_value_Jazz_stick")
	res = append(res, "Enum_Beater_value_Knitting_needle")
	res = append(res, "Enum_Beater_value_Metal_hammer")
	res = append(res, "Enum_Beater_value_Slide_brush_on_gong")
	res = append(res, "Enum_Beater_value_Snare_stick")
	res = append(res, "Enum_Beater_value_Spoon_mallet")
	res = append(res, "Enum_Beater_value_Superball")
	res = append(res, "Enum_Beater_value_Triangle_beater")
	res = append(res, "Enum_Beater_value_Triangle_beater_plain")
	res = append(res, "Enum_Beater_value_Wire_brush")

	return
}

func (enum_beater_value Enum_Beater_value) CodeValues() (res []string) {

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

// Utility function for Enum_Bend_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_bend_shape Enum_Bend_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_bend_shape {
	// insertion code per enum code
	case Enum_Bend_shape_Angled:
		res = "angled"
	case Enum_Bend_shape_Curved:
		res = "curved"
	}
	return
}

func (enum_bend_shape *Enum_Bend_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "angled":
		*enum_bend_shape = Enum_Bend_shape_Angled
		return
	case "curved":
		*enum_bend_shape = Enum_Bend_shape_Curved
		return
	default:
		return errUnkownEnum
	}
}

func (enum_bend_shape *Enum_Bend_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Bend_shape_Angled":
		*enum_bend_shape = Enum_Bend_shape_Angled
	case "Enum_Bend_shape_Curved":
		*enum_bend_shape = Enum_Bend_shape_Curved
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_bend_shape *Enum_Bend_shape) ToCodeString() (res string) {

	switch *enum_bend_shape {
	// insertion code per enum code
	case Enum_Bend_shape_Angled:
		res = "Enum_Bend_shape_Angled"
	case Enum_Bend_shape_Curved:
		res = "Enum_Bend_shape_Curved"
	}
	return
}

func (enum_bend_shape Enum_Bend_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Bend_shape_Angled")
	res = append(res, "Enum_Bend_shape_Curved")

	return
}

func (enum_bend_shape Enum_Bend_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "angled")
	res = append(res, "curved")

	return
}

// Utility function for Enum_Breath_mark_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_breath_mark_value Enum_Breath_mark_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_breath_mark_value {
	// insertion code per enum code
	case Enum_Breath_mark_value_:
		res = ""
	case Enum_Breath_mark_value_Comma:
		res = "comma"
	case Enum_Breath_mark_value_Tick:
		res = "tick"
	case Enum_Breath_mark_value_Upbow:
		res = "upbow"
	case Enum_Breath_mark_value_Salzedo:
		res = "salzedo"
	}
	return
}

func (enum_breath_mark_value *Enum_Breath_mark_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "":
		*enum_breath_mark_value = Enum_Breath_mark_value_
		return
	case "comma":
		*enum_breath_mark_value = Enum_Breath_mark_value_Comma
		return
	case "tick":
		*enum_breath_mark_value = Enum_Breath_mark_value_Tick
		return
	case "upbow":
		*enum_breath_mark_value = Enum_Breath_mark_value_Upbow
		return
	case "salzedo":
		*enum_breath_mark_value = Enum_Breath_mark_value_Salzedo
		return
	default:
		return errUnkownEnum
	}
}

func (enum_breath_mark_value *Enum_Breath_mark_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Breath_mark_value_":
		*enum_breath_mark_value = Enum_Breath_mark_value_
	case "Enum_Breath_mark_value_Comma":
		*enum_breath_mark_value = Enum_Breath_mark_value_Comma
	case "Enum_Breath_mark_value_Tick":
		*enum_breath_mark_value = Enum_Breath_mark_value_Tick
	case "Enum_Breath_mark_value_Upbow":
		*enum_breath_mark_value = Enum_Breath_mark_value_Upbow
	case "Enum_Breath_mark_value_Salzedo":
		*enum_breath_mark_value = Enum_Breath_mark_value_Salzedo
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_breath_mark_value *Enum_Breath_mark_value) ToCodeString() (res string) {

	switch *enum_breath_mark_value {
	// insertion code per enum code
	case Enum_Breath_mark_value_:
		res = "Enum_Breath_mark_value_"
	case Enum_Breath_mark_value_Comma:
		res = "Enum_Breath_mark_value_Comma"
	case Enum_Breath_mark_value_Tick:
		res = "Enum_Breath_mark_value_Tick"
	case Enum_Breath_mark_value_Upbow:
		res = "Enum_Breath_mark_value_Upbow"
	case Enum_Breath_mark_value_Salzedo:
		res = "Enum_Breath_mark_value_Salzedo"
	}
	return
}

func (enum_breath_mark_value Enum_Breath_mark_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Breath_mark_value_")
	res = append(res, "Enum_Breath_mark_value_Comma")
	res = append(res, "Enum_Breath_mark_value_Tick")
	res = append(res, "Enum_Breath_mark_value_Upbow")
	res = append(res, "Enum_Breath_mark_value_Salzedo")

	return
}

func (enum_breath_mark_value Enum_Breath_mark_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "")
	res = append(res, "comma")
	res = append(res, "tick")
	res = append(res, "upbow")
	res = append(res, "salzedo")

	return
}

// Utility function for Enum_Caesura_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_caesura_value Enum_Caesura_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_caesura_value {
	// insertion code per enum code
	case Enum_Caesura_value_Normal:
		res = "normal"
	case Enum_Caesura_value_Thick:
		res = "thick"
	case Enum_Caesura_value_Short:
		res = "short"
	case Enum_Caesura_value_Curved:
		res = "curved"
	case Enum_Caesura_value_Single:
		res = "single"
	case Enum_Caesura_value_:
		res = ""
	}
	return
}

func (enum_caesura_value *Enum_Caesura_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enum_caesura_value = Enum_Caesura_value_Normal
		return
	case "thick":
		*enum_caesura_value = Enum_Caesura_value_Thick
		return
	case "short":
		*enum_caesura_value = Enum_Caesura_value_Short
		return
	case "curved":
		*enum_caesura_value = Enum_Caesura_value_Curved
		return
	case "single":
		*enum_caesura_value = Enum_Caesura_value_Single
		return
	case "":
		*enum_caesura_value = Enum_Caesura_value_
		return
	default:
		return errUnkownEnum
	}
}

func (enum_caesura_value *Enum_Caesura_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Caesura_value_Normal":
		*enum_caesura_value = Enum_Caesura_value_Normal
	case "Enum_Caesura_value_Thick":
		*enum_caesura_value = Enum_Caesura_value_Thick
	case "Enum_Caesura_value_Short":
		*enum_caesura_value = Enum_Caesura_value_Short
	case "Enum_Caesura_value_Curved":
		*enum_caesura_value = Enum_Caesura_value_Curved
	case "Enum_Caesura_value_Single":
		*enum_caesura_value = Enum_Caesura_value_Single
	case "Enum_Caesura_value_":
		*enum_caesura_value = Enum_Caesura_value_
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_caesura_value *Enum_Caesura_value) ToCodeString() (res string) {

	switch *enum_caesura_value {
	// insertion code per enum code
	case Enum_Caesura_value_Normal:
		res = "Enum_Caesura_value_Normal"
	case Enum_Caesura_value_Thick:
		res = "Enum_Caesura_value_Thick"
	case Enum_Caesura_value_Short:
		res = "Enum_Caesura_value_Short"
	case Enum_Caesura_value_Curved:
		res = "Enum_Caesura_value_Curved"
	case Enum_Caesura_value_Single:
		res = "Enum_Caesura_value_Single"
	case Enum_Caesura_value_:
		res = "Enum_Caesura_value_"
	}
	return
}

func (enum_caesura_value Enum_Caesura_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Caesura_value_Normal")
	res = append(res, "Enum_Caesura_value_Thick")
	res = append(res, "Enum_Caesura_value_Short")
	res = append(res, "Enum_Caesura_value_Curved")
	res = append(res, "Enum_Caesura_value_Single")
	res = append(res, "Enum_Caesura_value_")

	return
}

func (enum_caesura_value Enum_Caesura_value) CodeValues() (res []string) {

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

// Utility function for Enum_Cancel_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_cancel_location Enum_Cancel_location) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_cancel_location {
	// insertion code per enum code
	case Enum_Cancel_location_Left:
		res = "left"
	case Enum_Cancel_location_Right:
		res = "right"
	case Enum_Cancel_location_Before_barline:
		res = "before-barline"
	}
	return
}

func (enum_cancel_location *Enum_Cancel_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enum_cancel_location = Enum_Cancel_location_Left
		return
	case "right":
		*enum_cancel_location = Enum_Cancel_location_Right
		return
	case "before-barline":
		*enum_cancel_location = Enum_Cancel_location_Before_barline
		return
	default:
		return errUnkownEnum
	}
}

func (enum_cancel_location *Enum_Cancel_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Cancel_location_Left":
		*enum_cancel_location = Enum_Cancel_location_Left
	case "Enum_Cancel_location_Right":
		*enum_cancel_location = Enum_Cancel_location_Right
	case "Enum_Cancel_location_Before_barline":
		*enum_cancel_location = Enum_Cancel_location_Before_barline
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_cancel_location *Enum_Cancel_location) ToCodeString() (res string) {

	switch *enum_cancel_location {
	// insertion code per enum code
	case Enum_Cancel_location_Left:
		res = "Enum_Cancel_location_Left"
	case Enum_Cancel_location_Right:
		res = "Enum_Cancel_location_Right"
	case Enum_Cancel_location_Before_barline:
		res = "Enum_Cancel_location_Before_barline"
	}
	return
}

func (enum_cancel_location Enum_Cancel_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Cancel_location_Left")
	res = append(res, "Enum_Cancel_location_Right")
	res = append(res, "Enum_Cancel_location_Before_barline")

	return
}

func (enum_cancel_location Enum_Cancel_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "before-barline")

	return
}

// Utility function for Enum_Circular_arrow
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_circular_arrow Enum_Circular_arrow) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_circular_arrow {
	// insertion code per enum code
	case Enum_Circular_arrow_Clockwise:
		res = "clockwise"
	case Enum_Circular_arrow_Anticlockwise:
		res = "anticlockwise"
	}
	return
}

func (enum_circular_arrow *Enum_Circular_arrow) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "clockwise":
		*enum_circular_arrow = Enum_Circular_arrow_Clockwise
		return
	case "anticlockwise":
		*enum_circular_arrow = Enum_Circular_arrow_Anticlockwise
		return
	default:
		return errUnkownEnum
	}
}

func (enum_circular_arrow *Enum_Circular_arrow) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Circular_arrow_Clockwise":
		*enum_circular_arrow = Enum_Circular_arrow_Clockwise
	case "Enum_Circular_arrow_Anticlockwise":
		*enum_circular_arrow = Enum_Circular_arrow_Anticlockwise
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_circular_arrow *Enum_Circular_arrow) ToCodeString() (res string) {

	switch *enum_circular_arrow {
	// insertion code per enum code
	case Enum_Circular_arrow_Clockwise:
		res = "Enum_Circular_arrow_Clockwise"
	case Enum_Circular_arrow_Anticlockwise:
		res = "Enum_Circular_arrow_Anticlockwise"
	}
	return
}

func (enum_circular_arrow Enum_Circular_arrow) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Circular_arrow_Clockwise")
	res = append(res, "Enum_Circular_arrow_Anticlockwise")

	return
}

func (enum_circular_arrow Enum_Circular_arrow) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "clockwise")
	res = append(res, "anticlockwise")

	return
}

// Utility function for Enum_Clef_sign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_clef_sign Enum_Clef_sign) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_clef_sign {
	// insertion code per enum code
	case Enum_Clef_sign_G:
		res = "G"
	case Enum_Clef_sign_F:
		res = "F"
	case Enum_Clef_sign_C:
		res = "C"
	case Enum_Clef_sign_Percussion:
		res = "percussion"
	case Enum_Clef_sign_TAB:
		res = "TAB"
	case Enum_Clef_sign_Jianpu:
		res = "jianpu"
	case Enum_Clef_sign_None:
		res = "none"
	}
	return
}

func (enum_clef_sign *Enum_Clef_sign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "G":
		*enum_clef_sign = Enum_Clef_sign_G
		return
	case "F":
		*enum_clef_sign = Enum_Clef_sign_F
		return
	case "C":
		*enum_clef_sign = Enum_Clef_sign_C
		return
	case "percussion":
		*enum_clef_sign = Enum_Clef_sign_Percussion
		return
	case "TAB":
		*enum_clef_sign = Enum_Clef_sign_TAB
		return
	case "jianpu":
		*enum_clef_sign = Enum_Clef_sign_Jianpu
		return
	case "none":
		*enum_clef_sign = Enum_Clef_sign_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_clef_sign *Enum_Clef_sign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Clef_sign_G":
		*enum_clef_sign = Enum_Clef_sign_G
	case "Enum_Clef_sign_F":
		*enum_clef_sign = Enum_Clef_sign_F
	case "Enum_Clef_sign_C":
		*enum_clef_sign = Enum_Clef_sign_C
	case "Enum_Clef_sign_Percussion":
		*enum_clef_sign = Enum_Clef_sign_Percussion
	case "Enum_Clef_sign_TAB":
		*enum_clef_sign = Enum_Clef_sign_TAB
	case "Enum_Clef_sign_Jianpu":
		*enum_clef_sign = Enum_Clef_sign_Jianpu
	case "Enum_Clef_sign_None":
		*enum_clef_sign = Enum_Clef_sign_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_clef_sign *Enum_Clef_sign) ToCodeString() (res string) {

	switch *enum_clef_sign {
	// insertion code per enum code
	case Enum_Clef_sign_G:
		res = "Enum_Clef_sign_G"
	case Enum_Clef_sign_F:
		res = "Enum_Clef_sign_F"
	case Enum_Clef_sign_C:
		res = "Enum_Clef_sign_C"
	case Enum_Clef_sign_Percussion:
		res = "Enum_Clef_sign_Percussion"
	case Enum_Clef_sign_TAB:
		res = "Enum_Clef_sign_TAB"
	case Enum_Clef_sign_Jianpu:
		res = "Enum_Clef_sign_Jianpu"
	case Enum_Clef_sign_None:
		res = "Enum_Clef_sign_None"
	}
	return
}

func (enum_clef_sign Enum_Clef_sign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Clef_sign_G")
	res = append(res, "Enum_Clef_sign_F")
	res = append(res, "Enum_Clef_sign_C")
	res = append(res, "Enum_Clef_sign_Percussion")
	res = append(res, "Enum_Clef_sign_TAB")
	res = append(res, "Enum_Clef_sign_Jianpu")
	res = append(res, "Enum_Clef_sign_None")

	return
}

func (enum_clef_sign Enum_Clef_sign) CodeValues() (res []string) {

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

// Utility function for Enum_Color
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_color Enum_Color) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_color {
	// insertion code per enum code
	}
	return
}

func (enum_color *Enum_Color) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_color *Enum_Color) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_color *Enum_Color) ToCodeString() (res string) {

	switch *enum_color {
	// insertion code per enum code
	}
	return
}

func (enum_color Enum_Color) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_color Enum_Color) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Comma_separated_text
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_comma_separated_text Enum_Comma_separated_text) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_comma_separated_text {
	// insertion code per enum code
	}
	return
}

func (enum_comma_separated_text *Enum_Comma_separated_text) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_comma_separated_text *Enum_Comma_separated_text) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_comma_separated_text *Enum_Comma_separated_text) ToCodeString() (res string) {

	switch *enum_comma_separated_text {
	// insertion code per enum code
	}
	return
}

func (enum_comma_separated_text Enum_Comma_separated_text) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_comma_separated_text Enum_Comma_separated_text) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Degree_symbol_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_degree_symbol_value Enum_Degree_symbol_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_degree_symbol_value {
	// insertion code per enum code
	case Enum_Degree_symbol_value_Major:
		res = "major"
	case Enum_Degree_symbol_value_Minor:
		res = "minor"
	case Enum_Degree_symbol_value_Augmented:
		res = "augmented"
	case Enum_Degree_symbol_value_Diminished:
		res = "diminished"
	case Enum_Degree_symbol_value_Half_diminished:
		res = "half-diminished"
	}
	return
}

func (enum_degree_symbol_value *Enum_Degree_symbol_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Major
		return
	case "minor":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Minor
		return
	case "augmented":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Augmented
		return
	case "diminished":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Diminished
		return
	case "half-diminished":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Half_diminished
		return
	default:
		return errUnkownEnum
	}
}

func (enum_degree_symbol_value *Enum_Degree_symbol_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Degree_symbol_value_Major":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Major
	case "Enum_Degree_symbol_value_Minor":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Minor
	case "Enum_Degree_symbol_value_Augmented":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Augmented
	case "Enum_Degree_symbol_value_Diminished":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Diminished
	case "Enum_Degree_symbol_value_Half_diminished":
		*enum_degree_symbol_value = Enum_Degree_symbol_value_Half_diminished
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_degree_symbol_value *Enum_Degree_symbol_value) ToCodeString() (res string) {

	switch *enum_degree_symbol_value {
	// insertion code per enum code
	case Enum_Degree_symbol_value_Major:
		res = "Enum_Degree_symbol_value_Major"
	case Enum_Degree_symbol_value_Minor:
		res = "Enum_Degree_symbol_value_Minor"
	case Enum_Degree_symbol_value_Augmented:
		res = "Enum_Degree_symbol_value_Augmented"
	case Enum_Degree_symbol_value_Diminished:
		res = "Enum_Degree_symbol_value_Diminished"
	case Enum_Degree_symbol_value_Half_diminished:
		res = "Enum_Degree_symbol_value_Half_diminished"
	}
	return
}

func (enum_degree_symbol_value Enum_Degree_symbol_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Degree_symbol_value_Major")
	res = append(res, "Enum_Degree_symbol_value_Minor")
	res = append(res, "Enum_Degree_symbol_value_Augmented")
	res = append(res, "Enum_Degree_symbol_value_Diminished")
	res = append(res, "Enum_Degree_symbol_value_Half_diminished")

	return
}

func (enum_degree_symbol_value Enum_Degree_symbol_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "augmented")
	res = append(res, "diminished")
	res = append(res, "half-diminished")

	return
}

// Utility function for Enum_Degree_type_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_degree_type_value Enum_Degree_type_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_degree_type_value {
	// insertion code per enum code
	case Enum_Degree_type_value_Add:
		res = "add"
	case Enum_Degree_type_value_Alter:
		res = "alter"
	case Enum_Degree_type_value_Subtract:
		res = "subtract"
	}
	return
}

func (enum_degree_type_value *Enum_Degree_type_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "add":
		*enum_degree_type_value = Enum_Degree_type_value_Add
		return
	case "alter":
		*enum_degree_type_value = Enum_Degree_type_value_Alter
		return
	case "subtract":
		*enum_degree_type_value = Enum_Degree_type_value_Subtract
		return
	default:
		return errUnkownEnum
	}
}

func (enum_degree_type_value *Enum_Degree_type_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Degree_type_value_Add":
		*enum_degree_type_value = Enum_Degree_type_value_Add
	case "Enum_Degree_type_value_Alter":
		*enum_degree_type_value = Enum_Degree_type_value_Alter
	case "Enum_Degree_type_value_Subtract":
		*enum_degree_type_value = Enum_Degree_type_value_Subtract
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_degree_type_value *Enum_Degree_type_value) ToCodeString() (res string) {

	switch *enum_degree_type_value {
	// insertion code per enum code
	case Enum_Degree_type_value_Add:
		res = "Enum_Degree_type_value_Add"
	case Enum_Degree_type_value_Alter:
		res = "Enum_Degree_type_value_Alter"
	case Enum_Degree_type_value_Subtract:
		res = "Enum_Degree_type_value_Subtract"
	}
	return
}

func (enum_degree_type_value Enum_Degree_type_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Degree_type_value_Add")
	res = append(res, "Enum_Degree_type_value_Alter")
	res = append(res, "Enum_Degree_type_value_Subtract")

	return
}

func (enum_degree_type_value Enum_Degree_type_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "add")
	res = append(res, "alter")
	res = append(res, "subtract")

	return
}

// Utility function for Enum_Distance_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_distance_type Enum_Distance_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_distance_type {
	// insertion code per enum code
	}
	return
}

func (enum_distance_type *Enum_Distance_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_distance_type *Enum_Distance_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_distance_type *Enum_Distance_type) ToCodeString() (res string) {

	switch *enum_distance_type {
	// insertion code per enum code
	}
	return
}

func (enum_distance_type Enum_Distance_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_distance_type Enum_Distance_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Effect_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_effect_value Enum_Effect_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_effect_value {
	// insertion code per enum code
	case Enum_Effect_value_Anvil:
		res = "anvil"
	case Enum_Effect_value_Auto_horn:
		res = "auto horn"
	case Enum_Effect_value_Bird_whistle:
		res = "bird whistle"
	case Enum_Effect_value_Cannon:
		res = "cannon"
	case Enum_Effect_value_Duck_call:
		res = "duck call"
	case Enum_Effect_value_Gun_shot:
		res = "gun shot"
	case Enum_Effect_value_Klaxon_horn:
		res = "klaxon horn"
	case Enum_Effect_value_Lions_roar:
		res = "lions roar"
	case Enum_Effect_value_Lotus_flute:
		res = "lotus flute"
	case Enum_Effect_value_Megaphone:
		res = "megaphone"
	case Enum_Effect_value_Police_whistle:
		res = "police whistle"
	case Enum_Effect_value_Siren:
		res = "siren"
	case Enum_Effect_value_Slide_whistle:
		res = "slide whistle"
	case Enum_Effect_value_Thunder_sheet:
		res = "thunder sheet"
	case Enum_Effect_value_Wind_machine:
		res = "wind machine"
	case Enum_Effect_value_Wind_whistle:
		res = "wind whistle"
	}
	return
}

func (enum_effect_value *Enum_Effect_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "anvil":
		*enum_effect_value = Enum_Effect_value_Anvil
		return
	case "auto horn":
		*enum_effect_value = Enum_Effect_value_Auto_horn
		return
	case "bird whistle":
		*enum_effect_value = Enum_Effect_value_Bird_whistle
		return
	case "cannon":
		*enum_effect_value = Enum_Effect_value_Cannon
		return
	case "duck call":
		*enum_effect_value = Enum_Effect_value_Duck_call
		return
	case "gun shot":
		*enum_effect_value = Enum_Effect_value_Gun_shot
		return
	case "klaxon horn":
		*enum_effect_value = Enum_Effect_value_Klaxon_horn
		return
	case "lions roar":
		*enum_effect_value = Enum_Effect_value_Lions_roar
		return
	case "lotus flute":
		*enum_effect_value = Enum_Effect_value_Lotus_flute
		return
	case "megaphone":
		*enum_effect_value = Enum_Effect_value_Megaphone
		return
	case "police whistle":
		*enum_effect_value = Enum_Effect_value_Police_whistle
		return
	case "siren":
		*enum_effect_value = Enum_Effect_value_Siren
		return
	case "slide whistle":
		*enum_effect_value = Enum_Effect_value_Slide_whistle
		return
	case "thunder sheet":
		*enum_effect_value = Enum_Effect_value_Thunder_sheet
		return
	case "wind machine":
		*enum_effect_value = Enum_Effect_value_Wind_machine
		return
	case "wind whistle":
		*enum_effect_value = Enum_Effect_value_Wind_whistle
		return
	default:
		return errUnkownEnum
	}
}

func (enum_effect_value *Enum_Effect_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Effect_value_Anvil":
		*enum_effect_value = Enum_Effect_value_Anvil
	case "Enum_Effect_value_Auto_horn":
		*enum_effect_value = Enum_Effect_value_Auto_horn
	case "Enum_Effect_value_Bird_whistle":
		*enum_effect_value = Enum_Effect_value_Bird_whistle
	case "Enum_Effect_value_Cannon":
		*enum_effect_value = Enum_Effect_value_Cannon
	case "Enum_Effect_value_Duck_call":
		*enum_effect_value = Enum_Effect_value_Duck_call
	case "Enum_Effect_value_Gun_shot":
		*enum_effect_value = Enum_Effect_value_Gun_shot
	case "Enum_Effect_value_Klaxon_horn":
		*enum_effect_value = Enum_Effect_value_Klaxon_horn
	case "Enum_Effect_value_Lions_roar":
		*enum_effect_value = Enum_Effect_value_Lions_roar
	case "Enum_Effect_value_Lotus_flute":
		*enum_effect_value = Enum_Effect_value_Lotus_flute
	case "Enum_Effect_value_Megaphone":
		*enum_effect_value = Enum_Effect_value_Megaphone
	case "Enum_Effect_value_Police_whistle":
		*enum_effect_value = Enum_Effect_value_Police_whistle
	case "Enum_Effect_value_Siren":
		*enum_effect_value = Enum_Effect_value_Siren
	case "Enum_Effect_value_Slide_whistle":
		*enum_effect_value = Enum_Effect_value_Slide_whistle
	case "Enum_Effect_value_Thunder_sheet":
		*enum_effect_value = Enum_Effect_value_Thunder_sheet
	case "Enum_Effect_value_Wind_machine":
		*enum_effect_value = Enum_Effect_value_Wind_machine
	case "Enum_Effect_value_Wind_whistle":
		*enum_effect_value = Enum_Effect_value_Wind_whistle
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_effect_value *Enum_Effect_value) ToCodeString() (res string) {

	switch *enum_effect_value {
	// insertion code per enum code
	case Enum_Effect_value_Anvil:
		res = "Enum_Effect_value_Anvil"
	case Enum_Effect_value_Auto_horn:
		res = "Enum_Effect_value_Auto_horn"
	case Enum_Effect_value_Bird_whistle:
		res = "Enum_Effect_value_Bird_whistle"
	case Enum_Effect_value_Cannon:
		res = "Enum_Effect_value_Cannon"
	case Enum_Effect_value_Duck_call:
		res = "Enum_Effect_value_Duck_call"
	case Enum_Effect_value_Gun_shot:
		res = "Enum_Effect_value_Gun_shot"
	case Enum_Effect_value_Klaxon_horn:
		res = "Enum_Effect_value_Klaxon_horn"
	case Enum_Effect_value_Lions_roar:
		res = "Enum_Effect_value_Lions_roar"
	case Enum_Effect_value_Lotus_flute:
		res = "Enum_Effect_value_Lotus_flute"
	case Enum_Effect_value_Megaphone:
		res = "Enum_Effect_value_Megaphone"
	case Enum_Effect_value_Police_whistle:
		res = "Enum_Effect_value_Police_whistle"
	case Enum_Effect_value_Siren:
		res = "Enum_Effect_value_Siren"
	case Enum_Effect_value_Slide_whistle:
		res = "Enum_Effect_value_Slide_whistle"
	case Enum_Effect_value_Thunder_sheet:
		res = "Enum_Effect_value_Thunder_sheet"
	case Enum_Effect_value_Wind_machine:
		res = "Enum_Effect_value_Wind_machine"
	case Enum_Effect_value_Wind_whistle:
		res = "Enum_Effect_value_Wind_whistle"
	}
	return
}

func (enum_effect_value Enum_Effect_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Effect_value_Anvil")
	res = append(res, "Enum_Effect_value_Auto_horn")
	res = append(res, "Enum_Effect_value_Bird_whistle")
	res = append(res, "Enum_Effect_value_Cannon")
	res = append(res, "Enum_Effect_value_Duck_call")
	res = append(res, "Enum_Effect_value_Gun_shot")
	res = append(res, "Enum_Effect_value_Klaxon_horn")
	res = append(res, "Enum_Effect_value_Lions_roar")
	res = append(res, "Enum_Effect_value_Lotus_flute")
	res = append(res, "Enum_Effect_value_Megaphone")
	res = append(res, "Enum_Effect_value_Police_whistle")
	res = append(res, "Enum_Effect_value_Siren")
	res = append(res, "Enum_Effect_value_Slide_whistle")
	res = append(res, "Enum_Effect_value_Thunder_sheet")
	res = append(res, "Enum_Effect_value_Wind_machine")
	res = append(res, "Enum_Effect_value_Wind_whistle")

	return
}

func (enum_effect_value Enum_Effect_value) CodeValues() (res []string) {

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

// Utility function for Enum_Enclosure_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_enclosure_shape Enum_Enclosure_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_enclosure_shape {
	// insertion code per enum code
	case Enum_Enclosure_shape_Rectangle:
		res = "rectangle"
	case Enum_Enclosure_shape_Square:
		res = "square"
	case Enum_Enclosure_shape_Oval:
		res = "oval"
	case Enum_Enclosure_shape_Circle:
		res = "circle"
	case Enum_Enclosure_shape_Bracket:
		res = "bracket"
	case Enum_Enclosure_shape_Inverted_bracket:
		res = "inverted-bracket"
	case Enum_Enclosure_shape_Triangle:
		res = "triangle"
	case Enum_Enclosure_shape_Diamond:
		res = "diamond"
	case Enum_Enclosure_shape_Pentagon:
		res = "pentagon"
	case Enum_Enclosure_shape_Hexagon:
		res = "hexagon"
	case Enum_Enclosure_shape_Heptagon:
		res = "heptagon"
	case Enum_Enclosure_shape_Octagon:
		res = "octagon"
	case Enum_Enclosure_shape_Nonagon:
		res = "nonagon"
	case Enum_Enclosure_shape_Decagon:
		res = "decagon"
	case Enum_Enclosure_shape_None:
		res = "none"
	}
	return
}

func (enum_enclosure_shape *Enum_Enclosure_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "rectangle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Rectangle
		return
	case "square":
		*enum_enclosure_shape = Enum_Enclosure_shape_Square
		return
	case "oval":
		*enum_enclosure_shape = Enum_Enclosure_shape_Oval
		return
	case "circle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Circle
		return
	case "bracket":
		*enum_enclosure_shape = Enum_Enclosure_shape_Bracket
		return
	case "inverted-bracket":
		*enum_enclosure_shape = Enum_Enclosure_shape_Inverted_bracket
		return
	case "triangle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Triangle
		return
	case "diamond":
		*enum_enclosure_shape = Enum_Enclosure_shape_Diamond
		return
	case "pentagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Pentagon
		return
	case "hexagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Hexagon
		return
	case "heptagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Heptagon
		return
	case "octagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Octagon
		return
	case "nonagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Nonagon
		return
	case "decagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Decagon
		return
	case "none":
		*enum_enclosure_shape = Enum_Enclosure_shape_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_enclosure_shape *Enum_Enclosure_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Enclosure_shape_Rectangle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Rectangle
	case "Enum_Enclosure_shape_Square":
		*enum_enclosure_shape = Enum_Enclosure_shape_Square
	case "Enum_Enclosure_shape_Oval":
		*enum_enclosure_shape = Enum_Enclosure_shape_Oval
	case "Enum_Enclosure_shape_Circle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Circle
	case "Enum_Enclosure_shape_Bracket":
		*enum_enclosure_shape = Enum_Enclosure_shape_Bracket
	case "Enum_Enclosure_shape_Inverted_bracket":
		*enum_enclosure_shape = Enum_Enclosure_shape_Inverted_bracket
	case "Enum_Enclosure_shape_Triangle":
		*enum_enclosure_shape = Enum_Enclosure_shape_Triangle
	case "Enum_Enclosure_shape_Diamond":
		*enum_enclosure_shape = Enum_Enclosure_shape_Diamond
	case "Enum_Enclosure_shape_Pentagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Pentagon
	case "Enum_Enclosure_shape_Hexagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Hexagon
	case "Enum_Enclosure_shape_Heptagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Heptagon
	case "Enum_Enclosure_shape_Octagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Octagon
	case "Enum_Enclosure_shape_Nonagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Nonagon
	case "Enum_Enclosure_shape_Decagon":
		*enum_enclosure_shape = Enum_Enclosure_shape_Decagon
	case "Enum_Enclosure_shape_None":
		*enum_enclosure_shape = Enum_Enclosure_shape_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_enclosure_shape *Enum_Enclosure_shape) ToCodeString() (res string) {

	switch *enum_enclosure_shape {
	// insertion code per enum code
	case Enum_Enclosure_shape_Rectangle:
		res = "Enum_Enclosure_shape_Rectangle"
	case Enum_Enclosure_shape_Square:
		res = "Enum_Enclosure_shape_Square"
	case Enum_Enclosure_shape_Oval:
		res = "Enum_Enclosure_shape_Oval"
	case Enum_Enclosure_shape_Circle:
		res = "Enum_Enclosure_shape_Circle"
	case Enum_Enclosure_shape_Bracket:
		res = "Enum_Enclosure_shape_Bracket"
	case Enum_Enclosure_shape_Inverted_bracket:
		res = "Enum_Enclosure_shape_Inverted_bracket"
	case Enum_Enclosure_shape_Triangle:
		res = "Enum_Enclosure_shape_Triangle"
	case Enum_Enclosure_shape_Diamond:
		res = "Enum_Enclosure_shape_Diamond"
	case Enum_Enclosure_shape_Pentagon:
		res = "Enum_Enclosure_shape_Pentagon"
	case Enum_Enclosure_shape_Hexagon:
		res = "Enum_Enclosure_shape_Hexagon"
	case Enum_Enclosure_shape_Heptagon:
		res = "Enum_Enclosure_shape_Heptagon"
	case Enum_Enclosure_shape_Octagon:
		res = "Enum_Enclosure_shape_Octagon"
	case Enum_Enclosure_shape_Nonagon:
		res = "Enum_Enclosure_shape_Nonagon"
	case Enum_Enclosure_shape_Decagon:
		res = "Enum_Enclosure_shape_Decagon"
	case Enum_Enclosure_shape_None:
		res = "Enum_Enclosure_shape_None"
	}
	return
}

func (enum_enclosure_shape Enum_Enclosure_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Enclosure_shape_Rectangle")
	res = append(res, "Enum_Enclosure_shape_Square")
	res = append(res, "Enum_Enclosure_shape_Oval")
	res = append(res, "Enum_Enclosure_shape_Circle")
	res = append(res, "Enum_Enclosure_shape_Bracket")
	res = append(res, "Enum_Enclosure_shape_Inverted_bracket")
	res = append(res, "Enum_Enclosure_shape_Triangle")
	res = append(res, "Enum_Enclosure_shape_Diamond")
	res = append(res, "Enum_Enclosure_shape_Pentagon")
	res = append(res, "Enum_Enclosure_shape_Hexagon")
	res = append(res, "Enum_Enclosure_shape_Heptagon")
	res = append(res, "Enum_Enclosure_shape_Octagon")
	res = append(res, "Enum_Enclosure_shape_Nonagon")
	res = append(res, "Enum_Enclosure_shape_Decagon")
	res = append(res, "Enum_Enclosure_shape_None")

	return
}

func (enum_enclosure_shape Enum_Enclosure_shape) CodeValues() (res []string) {

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

// Utility function for Enum_Ending_number
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_ending_number Enum_Ending_number) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_ending_number {
	// insertion code per enum code
	}
	return
}

func (enum_ending_number *Enum_Ending_number) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_ending_number *Enum_Ending_number) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_ending_number *Enum_Ending_number) ToCodeString() (res string) {

	switch *enum_ending_number {
	// insertion code per enum code
	}
	return
}

func (enum_ending_number Enum_Ending_number) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_ending_number Enum_Ending_number) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Fan
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_fan Enum_Fan) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_fan {
	// insertion code per enum code
	case Enum_Fan_Accel:
		res = "accel"
	case Enum_Fan_Rit:
		res = "rit"
	case Enum_Fan_None:
		res = "none"
	}
	return
}

func (enum_fan *Enum_Fan) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "accel":
		*enum_fan = Enum_Fan_Accel
		return
	case "rit":
		*enum_fan = Enum_Fan_Rit
		return
	case "none":
		*enum_fan = Enum_Fan_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_fan *Enum_Fan) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Fan_Accel":
		*enum_fan = Enum_Fan_Accel
	case "Enum_Fan_Rit":
		*enum_fan = Enum_Fan_Rit
	case "Enum_Fan_None":
		*enum_fan = Enum_Fan_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_fan *Enum_Fan) ToCodeString() (res string) {

	switch *enum_fan {
	// insertion code per enum code
	case Enum_Fan_Accel:
		res = "Enum_Fan_Accel"
	case Enum_Fan_Rit:
		res = "Enum_Fan_Rit"
	case Enum_Fan_None:
		res = "Enum_Fan_None"
	}
	return
}

func (enum_fan Enum_Fan) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Fan_Accel")
	res = append(res, "Enum_Fan_Rit")
	res = append(res, "Enum_Fan_None")

	return
}

func (enum_fan Enum_Fan) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "accel")
	res = append(res, "rit")
	res = append(res, "none")

	return
}

// Utility function for Enum_Fermata_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_fermata_shape Enum_Fermata_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_fermata_shape {
	// insertion code per enum code
	case Enum_Fermata_shape_Normal:
		res = "normal"
	case Enum_Fermata_shape_Angled:
		res = "angled"
	case Enum_Fermata_shape_Square:
		res = "square"
	case Enum_Fermata_shape_Double_angled:
		res = "double-angled"
	case Enum_Fermata_shape_Double_square:
		res = "double-square"
	case Enum_Fermata_shape_Double_dot:
		res = "double-dot"
	case Enum_Fermata_shape_Half_curve:
		res = "half-curve"
	case Enum_Fermata_shape_Curlew:
		res = "curlew"
	case Enum_Fermata_shape_:
		res = ""
	}
	return
}

func (enum_fermata_shape *Enum_Fermata_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enum_fermata_shape = Enum_Fermata_shape_Normal
		return
	case "angled":
		*enum_fermata_shape = Enum_Fermata_shape_Angled
		return
	case "square":
		*enum_fermata_shape = Enum_Fermata_shape_Square
		return
	case "double-angled":
		*enum_fermata_shape = Enum_Fermata_shape_Double_angled
		return
	case "double-square":
		*enum_fermata_shape = Enum_Fermata_shape_Double_square
		return
	case "double-dot":
		*enum_fermata_shape = Enum_Fermata_shape_Double_dot
		return
	case "half-curve":
		*enum_fermata_shape = Enum_Fermata_shape_Half_curve
		return
	case "curlew":
		*enum_fermata_shape = Enum_Fermata_shape_Curlew
		return
	case "":
		*enum_fermata_shape = Enum_Fermata_shape_
		return
	default:
		return errUnkownEnum
	}
}

func (enum_fermata_shape *Enum_Fermata_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Fermata_shape_Normal":
		*enum_fermata_shape = Enum_Fermata_shape_Normal
	case "Enum_Fermata_shape_Angled":
		*enum_fermata_shape = Enum_Fermata_shape_Angled
	case "Enum_Fermata_shape_Square":
		*enum_fermata_shape = Enum_Fermata_shape_Square
	case "Enum_Fermata_shape_Double_angled":
		*enum_fermata_shape = Enum_Fermata_shape_Double_angled
	case "Enum_Fermata_shape_Double_square":
		*enum_fermata_shape = Enum_Fermata_shape_Double_square
	case "Enum_Fermata_shape_Double_dot":
		*enum_fermata_shape = Enum_Fermata_shape_Double_dot
	case "Enum_Fermata_shape_Half_curve":
		*enum_fermata_shape = Enum_Fermata_shape_Half_curve
	case "Enum_Fermata_shape_Curlew":
		*enum_fermata_shape = Enum_Fermata_shape_Curlew
	case "Enum_Fermata_shape_":
		*enum_fermata_shape = Enum_Fermata_shape_
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_fermata_shape *Enum_Fermata_shape) ToCodeString() (res string) {

	switch *enum_fermata_shape {
	// insertion code per enum code
	case Enum_Fermata_shape_Normal:
		res = "Enum_Fermata_shape_Normal"
	case Enum_Fermata_shape_Angled:
		res = "Enum_Fermata_shape_Angled"
	case Enum_Fermata_shape_Square:
		res = "Enum_Fermata_shape_Square"
	case Enum_Fermata_shape_Double_angled:
		res = "Enum_Fermata_shape_Double_angled"
	case Enum_Fermata_shape_Double_square:
		res = "Enum_Fermata_shape_Double_square"
	case Enum_Fermata_shape_Double_dot:
		res = "Enum_Fermata_shape_Double_dot"
	case Enum_Fermata_shape_Half_curve:
		res = "Enum_Fermata_shape_Half_curve"
	case Enum_Fermata_shape_Curlew:
		res = "Enum_Fermata_shape_Curlew"
	case Enum_Fermata_shape_:
		res = "Enum_Fermata_shape_"
	}
	return
}

func (enum_fermata_shape Enum_Fermata_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Fermata_shape_Normal")
	res = append(res, "Enum_Fermata_shape_Angled")
	res = append(res, "Enum_Fermata_shape_Square")
	res = append(res, "Enum_Fermata_shape_Double_angled")
	res = append(res, "Enum_Fermata_shape_Double_square")
	res = append(res, "Enum_Fermata_shape_Double_dot")
	res = append(res, "Enum_Fermata_shape_Half_curve")
	res = append(res, "Enum_Fermata_shape_Curlew")
	res = append(res, "Enum_Fermata_shape_")

	return
}

func (enum_fermata_shape Enum_Fermata_shape) CodeValues() (res []string) {

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

// Utility function for Enum_Font_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_font_style Enum_Font_style) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_font_style {
	// insertion code per enum code
	case Enum_Font_style_Normal:
		res = "normal"
	case Enum_Font_style_Italic:
		res = "italic"
	}
	return
}

func (enum_font_style *Enum_Font_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enum_font_style = Enum_Font_style_Normal
		return
	case "italic":
		*enum_font_style = Enum_Font_style_Italic
		return
	default:
		return errUnkownEnum
	}
}

func (enum_font_style *Enum_Font_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Font_style_Normal":
		*enum_font_style = Enum_Font_style_Normal
	case "Enum_Font_style_Italic":
		*enum_font_style = Enum_Font_style_Italic
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_font_style *Enum_Font_style) ToCodeString() (res string) {

	switch *enum_font_style {
	// insertion code per enum code
	case Enum_Font_style_Normal:
		res = "Enum_Font_style_Normal"
	case Enum_Font_style_Italic:
		res = "Enum_Font_style_Italic"
	}
	return
}

func (enum_font_style Enum_Font_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Font_style_Normal")
	res = append(res, "Enum_Font_style_Italic")

	return
}

func (enum_font_style Enum_Font_style) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "italic")

	return
}

// Utility function for Enum_Font_weight
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_font_weight Enum_Font_weight) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_font_weight {
	// insertion code per enum code
	case Enum_Font_weight_Normal:
		res = "normal"
	case Enum_Font_weight_Bold:
		res = "bold"
	}
	return
}

func (enum_font_weight *Enum_Font_weight) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*enum_font_weight = Enum_Font_weight_Normal
		return
	case "bold":
		*enum_font_weight = Enum_Font_weight_Bold
		return
	default:
		return errUnkownEnum
	}
}

func (enum_font_weight *Enum_Font_weight) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Font_weight_Normal":
		*enum_font_weight = Enum_Font_weight_Normal
	case "Enum_Font_weight_Bold":
		*enum_font_weight = Enum_Font_weight_Bold
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_font_weight *Enum_Font_weight) ToCodeString() (res string) {

	switch *enum_font_weight {
	// insertion code per enum code
	case Enum_Font_weight_Normal:
		res = "Enum_Font_weight_Normal"
	case Enum_Font_weight_Bold:
		res = "Enum_Font_weight_Bold"
	}
	return
}

func (enum_font_weight Enum_Font_weight) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Font_weight_Normal")
	res = append(res, "Enum_Font_weight_Bold")

	return
}

func (enum_font_weight Enum_Font_weight) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "bold")

	return
}

// Utility function for Enum_Glass_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_glass_value Enum_Glass_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_glass_value {
	// insertion code per enum code
	case Enum_Glass_value_Glass_harmonica:
		res = "glass harmonica"
	case Enum_Glass_value_Glass_harp:
		res = "glass harp"
	case Enum_Glass_value_Wind_chimes:
		res = "wind chimes"
	}
	return
}

func (enum_glass_value *Enum_Glass_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "glass harmonica":
		*enum_glass_value = Enum_Glass_value_Glass_harmonica
		return
	case "glass harp":
		*enum_glass_value = Enum_Glass_value_Glass_harp
		return
	case "wind chimes":
		*enum_glass_value = Enum_Glass_value_Wind_chimes
		return
	default:
		return errUnkownEnum
	}
}

func (enum_glass_value *Enum_Glass_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Glass_value_Glass_harmonica":
		*enum_glass_value = Enum_Glass_value_Glass_harmonica
	case "Enum_Glass_value_Glass_harp":
		*enum_glass_value = Enum_Glass_value_Glass_harp
	case "Enum_Glass_value_Wind_chimes":
		*enum_glass_value = Enum_Glass_value_Wind_chimes
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_glass_value *Enum_Glass_value) ToCodeString() (res string) {

	switch *enum_glass_value {
	// insertion code per enum code
	case Enum_Glass_value_Glass_harmonica:
		res = "Enum_Glass_value_Glass_harmonica"
	case Enum_Glass_value_Glass_harp:
		res = "Enum_Glass_value_Glass_harp"
	case Enum_Glass_value_Wind_chimes:
		res = "Enum_Glass_value_Wind_chimes"
	}
	return
}

func (enum_glass_value Enum_Glass_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Glass_value_Glass_harmonica")
	res = append(res, "Enum_Glass_value_Glass_harp")
	res = append(res, "Enum_Glass_value_Wind_chimes")

	return
}

func (enum_glass_value Enum_Glass_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "glass harmonica")
	res = append(res, "glass harp")
	res = append(res, "wind chimes")

	return
}

// Utility function for Enum_Glyph_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_glyph_type Enum_Glyph_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_glyph_type {
	// insertion code per enum code
	}
	return
}

func (enum_glyph_type *Enum_Glyph_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_glyph_type *Enum_Glyph_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_glyph_type *Enum_Glyph_type) ToCodeString() (res string) {

	switch *enum_glyph_type {
	// insertion code per enum code
	}
	return
}

func (enum_glyph_type Enum_Glyph_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_glyph_type Enum_Glyph_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Group_barline_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_group_barline_value Enum_Group_barline_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_group_barline_value {
	// insertion code per enum code
	case Enum_Group_barline_value_Yes:
		res = "yes"
	case Enum_Group_barline_value_No:
		res = "no"
	case Enum_Group_barline_value_Mensurstrich:
		res = "Mensurstrich"
	}
	return
}

func (enum_group_barline_value *Enum_Group_barline_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enum_group_barline_value = Enum_Group_barline_value_Yes
		return
	case "no":
		*enum_group_barline_value = Enum_Group_barline_value_No
		return
	case "Mensurstrich":
		*enum_group_barline_value = Enum_Group_barline_value_Mensurstrich
		return
	default:
		return errUnkownEnum
	}
}

func (enum_group_barline_value *Enum_Group_barline_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Group_barline_value_Yes":
		*enum_group_barline_value = Enum_Group_barline_value_Yes
	case "Enum_Group_barline_value_No":
		*enum_group_barline_value = Enum_Group_barline_value_No
	case "Enum_Group_barline_value_Mensurstrich":
		*enum_group_barline_value = Enum_Group_barline_value_Mensurstrich
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_group_barline_value *Enum_Group_barline_value) ToCodeString() (res string) {

	switch *enum_group_barline_value {
	// insertion code per enum code
	case Enum_Group_barline_value_Yes:
		res = "Enum_Group_barline_value_Yes"
	case Enum_Group_barline_value_No:
		res = "Enum_Group_barline_value_No"
	case Enum_Group_barline_value_Mensurstrich:
		res = "Enum_Group_barline_value_Mensurstrich"
	}
	return
}

func (enum_group_barline_value Enum_Group_barline_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Group_barline_value_Yes")
	res = append(res, "Enum_Group_barline_value_No")
	res = append(res, "Enum_Group_barline_value_Mensurstrich")

	return
}

func (enum_group_barline_value Enum_Group_barline_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "Mensurstrich")

	return
}

// Utility function for Enum_Group_symbol_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_group_symbol_value Enum_Group_symbol_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_group_symbol_value {
	// insertion code per enum code
	case Enum_Group_symbol_value_None:
		res = "none"
	case Enum_Group_symbol_value_Brace:
		res = "brace"
	case Enum_Group_symbol_value_Line:
		res = "line"
	case Enum_Group_symbol_value_Bracket:
		res = "bracket"
	case Enum_Group_symbol_value_Square:
		res = "square"
	}
	return
}

func (enum_group_symbol_value *Enum_Group_symbol_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_group_symbol_value = Enum_Group_symbol_value_None
		return
	case "brace":
		*enum_group_symbol_value = Enum_Group_symbol_value_Brace
		return
	case "line":
		*enum_group_symbol_value = Enum_Group_symbol_value_Line
		return
	case "bracket":
		*enum_group_symbol_value = Enum_Group_symbol_value_Bracket
		return
	case "square":
		*enum_group_symbol_value = Enum_Group_symbol_value_Square
		return
	default:
		return errUnkownEnum
	}
}

func (enum_group_symbol_value *Enum_Group_symbol_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Group_symbol_value_None":
		*enum_group_symbol_value = Enum_Group_symbol_value_None
	case "Enum_Group_symbol_value_Brace":
		*enum_group_symbol_value = Enum_Group_symbol_value_Brace
	case "Enum_Group_symbol_value_Line":
		*enum_group_symbol_value = Enum_Group_symbol_value_Line
	case "Enum_Group_symbol_value_Bracket":
		*enum_group_symbol_value = Enum_Group_symbol_value_Bracket
	case "Enum_Group_symbol_value_Square":
		*enum_group_symbol_value = Enum_Group_symbol_value_Square
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_group_symbol_value *Enum_Group_symbol_value) ToCodeString() (res string) {

	switch *enum_group_symbol_value {
	// insertion code per enum code
	case Enum_Group_symbol_value_None:
		res = "Enum_Group_symbol_value_None"
	case Enum_Group_symbol_value_Brace:
		res = "Enum_Group_symbol_value_Brace"
	case Enum_Group_symbol_value_Line:
		res = "Enum_Group_symbol_value_Line"
	case Enum_Group_symbol_value_Bracket:
		res = "Enum_Group_symbol_value_Bracket"
	case Enum_Group_symbol_value_Square:
		res = "Enum_Group_symbol_value_Square"
	}
	return
}

func (enum_group_symbol_value Enum_Group_symbol_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Group_symbol_value_None")
	res = append(res, "Enum_Group_symbol_value_Brace")
	res = append(res, "Enum_Group_symbol_value_Line")
	res = append(res, "Enum_Group_symbol_value_Bracket")
	res = append(res, "Enum_Group_symbol_value_Square")

	return
}

func (enum_group_symbol_value Enum_Group_symbol_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "brace")
	res = append(res, "line")
	res = append(res, "bracket")
	res = append(res, "square")

	return
}

// Utility function for Enum_Handbell_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_handbell_value Enum_Handbell_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_handbell_value {
	// insertion code per enum code
	case Enum_Handbell_value_Belltree:
		res = "belltree"
	case Enum_Handbell_value_Damp:
		res = "damp"
	case Enum_Handbell_value_Echo:
		res = "echo"
	case Enum_Handbell_value_Gyro:
		res = "gyro"
	case Enum_Handbell_value_Hand_martellato:
		res = "hand martellato"
	case Enum_Handbell_value_Mallet_lift:
		res = "mallet lift"
	case Enum_Handbell_value_Mallet_table:
		res = "mallet table"
	case Enum_Handbell_value_Martellato:
		res = "martellato"
	case Enum_Handbell_value_Martellato_lift:
		res = "martellato lift"
	case Enum_Handbell_value_Muted_martellato:
		res = "muted martellato"
	case Enum_Handbell_value_Pluck_lift:
		res = "pluck lift"
	case Enum_Handbell_value_Swing:
		res = "swing"
	}
	return
}

func (enum_handbell_value *Enum_Handbell_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "belltree":
		*enum_handbell_value = Enum_Handbell_value_Belltree
		return
	case "damp":
		*enum_handbell_value = Enum_Handbell_value_Damp
		return
	case "echo":
		*enum_handbell_value = Enum_Handbell_value_Echo
		return
	case "gyro":
		*enum_handbell_value = Enum_Handbell_value_Gyro
		return
	case "hand martellato":
		*enum_handbell_value = Enum_Handbell_value_Hand_martellato
		return
	case "mallet lift":
		*enum_handbell_value = Enum_Handbell_value_Mallet_lift
		return
	case "mallet table":
		*enum_handbell_value = Enum_Handbell_value_Mallet_table
		return
	case "martellato":
		*enum_handbell_value = Enum_Handbell_value_Martellato
		return
	case "martellato lift":
		*enum_handbell_value = Enum_Handbell_value_Martellato_lift
		return
	case "muted martellato":
		*enum_handbell_value = Enum_Handbell_value_Muted_martellato
		return
	case "pluck lift":
		*enum_handbell_value = Enum_Handbell_value_Pluck_lift
		return
	case "swing":
		*enum_handbell_value = Enum_Handbell_value_Swing
		return
	default:
		return errUnkownEnum
	}
}

func (enum_handbell_value *Enum_Handbell_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Handbell_value_Belltree":
		*enum_handbell_value = Enum_Handbell_value_Belltree
	case "Enum_Handbell_value_Damp":
		*enum_handbell_value = Enum_Handbell_value_Damp
	case "Enum_Handbell_value_Echo":
		*enum_handbell_value = Enum_Handbell_value_Echo
	case "Enum_Handbell_value_Gyro":
		*enum_handbell_value = Enum_Handbell_value_Gyro
	case "Enum_Handbell_value_Hand_martellato":
		*enum_handbell_value = Enum_Handbell_value_Hand_martellato
	case "Enum_Handbell_value_Mallet_lift":
		*enum_handbell_value = Enum_Handbell_value_Mallet_lift
	case "Enum_Handbell_value_Mallet_table":
		*enum_handbell_value = Enum_Handbell_value_Mallet_table
	case "Enum_Handbell_value_Martellato":
		*enum_handbell_value = Enum_Handbell_value_Martellato
	case "Enum_Handbell_value_Martellato_lift":
		*enum_handbell_value = Enum_Handbell_value_Martellato_lift
	case "Enum_Handbell_value_Muted_martellato":
		*enum_handbell_value = Enum_Handbell_value_Muted_martellato
	case "Enum_Handbell_value_Pluck_lift":
		*enum_handbell_value = Enum_Handbell_value_Pluck_lift
	case "Enum_Handbell_value_Swing":
		*enum_handbell_value = Enum_Handbell_value_Swing
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_handbell_value *Enum_Handbell_value) ToCodeString() (res string) {

	switch *enum_handbell_value {
	// insertion code per enum code
	case Enum_Handbell_value_Belltree:
		res = "Enum_Handbell_value_Belltree"
	case Enum_Handbell_value_Damp:
		res = "Enum_Handbell_value_Damp"
	case Enum_Handbell_value_Echo:
		res = "Enum_Handbell_value_Echo"
	case Enum_Handbell_value_Gyro:
		res = "Enum_Handbell_value_Gyro"
	case Enum_Handbell_value_Hand_martellato:
		res = "Enum_Handbell_value_Hand_martellato"
	case Enum_Handbell_value_Mallet_lift:
		res = "Enum_Handbell_value_Mallet_lift"
	case Enum_Handbell_value_Mallet_table:
		res = "Enum_Handbell_value_Mallet_table"
	case Enum_Handbell_value_Martellato:
		res = "Enum_Handbell_value_Martellato"
	case Enum_Handbell_value_Martellato_lift:
		res = "Enum_Handbell_value_Martellato_lift"
	case Enum_Handbell_value_Muted_martellato:
		res = "Enum_Handbell_value_Muted_martellato"
	case Enum_Handbell_value_Pluck_lift:
		res = "Enum_Handbell_value_Pluck_lift"
	case Enum_Handbell_value_Swing:
		res = "Enum_Handbell_value_Swing"
	}
	return
}

func (enum_handbell_value Enum_Handbell_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Handbell_value_Belltree")
	res = append(res, "Enum_Handbell_value_Damp")
	res = append(res, "Enum_Handbell_value_Echo")
	res = append(res, "Enum_Handbell_value_Gyro")
	res = append(res, "Enum_Handbell_value_Hand_martellato")
	res = append(res, "Enum_Handbell_value_Mallet_lift")
	res = append(res, "Enum_Handbell_value_Mallet_table")
	res = append(res, "Enum_Handbell_value_Martellato")
	res = append(res, "Enum_Handbell_value_Martellato_lift")
	res = append(res, "Enum_Handbell_value_Muted_martellato")
	res = append(res, "Enum_Handbell_value_Pluck_lift")
	res = append(res, "Enum_Handbell_value_Swing")

	return
}

func (enum_handbell_value Enum_Handbell_value) CodeValues() (res []string) {

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

// Utility function for Enum_Harmon_closed_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_harmon_closed_location Enum_Harmon_closed_location) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_harmon_closed_location {
	// insertion code per enum code
	case Enum_Harmon_closed_location_Right:
		res = "right"
	case Enum_Harmon_closed_location_Bottom:
		res = "bottom"
	case Enum_Harmon_closed_location_Left:
		res = "left"
	case Enum_Harmon_closed_location_Top:
		res = "top"
	}
	return
}

func (enum_harmon_closed_location *Enum_Harmon_closed_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Right
		return
	case "bottom":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Bottom
		return
	case "left":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Left
		return
	case "top":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Top
		return
	default:
		return errUnkownEnum
	}
}

func (enum_harmon_closed_location *Enum_Harmon_closed_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Harmon_closed_location_Right":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Right
	case "Enum_Harmon_closed_location_Bottom":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Bottom
	case "Enum_Harmon_closed_location_Left":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Left
	case "Enum_Harmon_closed_location_Top":
		*enum_harmon_closed_location = Enum_Harmon_closed_location_Top
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_harmon_closed_location *Enum_Harmon_closed_location) ToCodeString() (res string) {

	switch *enum_harmon_closed_location {
	// insertion code per enum code
	case Enum_Harmon_closed_location_Right:
		res = "Enum_Harmon_closed_location_Right"
	case Enum_Harmon_closed_location_Bottom:
		res = "Enum_Harmon_closed_location_Bottom"
	case Enum_Harmon_closed_location_Left:
		res = "Enum_Harmon_closed_location_Left"
	case Enum_Harmon_closed_location_Top:
		res = "Enum_Harmon_closed_location_Top"
	}
	return
}

func (enum_harmon_closed_location Enum_Harmon_closed_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Harmon_closed_location_Right")
	res = append(res, "Enum_Harmon_closed_location_Bottom")
	res = append(res, "Enum_Harmon_closed_location_Left")
	res = append(res, "Enum_Harmon_closed_location_Top")

	return
}

func (enum_harmon_closed_location Enum_Harmon_closed_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for Enum_Harmon_closed_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_harmon_closed_value Enum_Harmon_closed_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_harmon_closed_value {
	// insertion code per enum code
	case Enum_Harmon_closed_value_Yes:
		res = "yes"
	case Enum_Harmon_closed_value_No:
		res = "no"
	case Enum_Harmon_closed_value_Half:
		res = "half"
	}
	return
}

func (enum_harmon_closed_value *Enum_Harmon_closed_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_Yes
		return
	case "no":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_No
		return
	case "half":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_Half
		return
	default:
		return errUnkownEnum
	}
}

func (enum_harmon_closed_value *Enum_Harmon_closed_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Harmon_closed_value_Yes":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_Yes
	case "Enum_Harmon_closed_value_No":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_No
	case "Enum_Harmon_closed_value_Half":
		*enum_harmon_closed_value = Enum_Harmon_closed_value_Half
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_harmon_closed_value *Enum_Harmon_closed_value) ToCodeString() (res string) {

	switch *enum_harmon_closed_value {
	// insertion code per enum code
	case Enum_Harmon_closed_value_Yes:
		res = "Enum_Harmon_closed_value_Yes"
	case Enum_Harmon_closed_value_No:
		res = "Enum_Harmon_closed_value_No"
	case Enum_Harmon_closed_value_Half:
		res = "Enum_Harmon_closed_value_Half"
	}
	return
}

func (enum_harmon_closed_value Enum_Harmon_closed_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Harmon_closed_value_Yes")
	res = append(res, "Enum_Harmon_closed_value_No")
	res = append(res, "Enum_Harmon_closed_value_Half")

	return
}

func (enum_harmon_closed_value Enum_Harmon_closed_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for Enum_Harmony_arrangement
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_harmony_arrangement Enum_Harmony_arrangement) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_harmony_arrangement {
	// insertion code per enum code
	case Enum_Harmony_arrangement_Vertical:
		res = "vertical"
	case Enum_Harmony_arrangement_Horizontal:
		res = "horizontal"
	case Enum_Harmony_arrangement_Diagonal:
		res = "diagonal"
	}
	return
}

func (enum_harmony_arrangement *Enum_Harmony_arrangement) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "vertical":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Vertical
		return
	case "horizontal":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Horizontal
		return
	case "diagonal":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Diagonal
		return
	default:
		return errUnkownEnum
	}
}

func (enum_harmony_arrangement *Enum_Harmony_arrangement) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Harmony_arrangement_Vertical":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Vertical
	case "Enum_Harmony_arrangement_Horizontal":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Horizontal
	case "Enum_Harmony_arrangement_Diagonal":
		*enum_harmony_arrangement = Enum_Harmony_arrangement_Diagonal
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_harmony_arrangement *Enum_Harmony_arrangement) ToCodeString() (res string) {

	switch *enum_harmony_arrangement {
	// insertion code per enum code
	case Enum_Harmony_arrangement_Vertical:
		res = "Enum_Harmony_arrangement_Vertical"
	case Enum_Harmony_arrangement_Horizontal:
		res = "Enum_Harmony_arrangement_Horizontal"
	case Enum_Harmony_arrangement_Diagonal:
		res = "Enum_Harmony_arrangement_Diagonal"
	}
	return
}

func (enum_harmony_arrangement Enum_Harmony_arrangement) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Harmony_arrangement_Vertical")
	res = append(res, "Enum_Harmony_arrangement_Horizontal")
	res = append(res, "Enum_Harmony_arrangement_Diagonal")

	return
}

func (enum_harmony_arrangement Enum_Harmony_arrangement) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "vertical")
	res = append(res, "horizontal")
	res = append(res, "diagonal")

	return
}

// Utility function for Enum_Harmony_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_harmony_type Enum_Harmony_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_harmony_type {
	// insertion code per enum code
	case Enum_Harmony_type_Explicit:
		res = "explicit"
	case Enum_Harmony_type_Implied:
		res = "implied"
	case Enum_Harmony_type_Alternate:
		res = "alternate"
	}
	return
}

func (enum_harmony_type *Enum_Harmony_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "explicit":
		*enum_harmony_type = Enum_Harmony_type_Explicit
		return
	case "implied":
		*enum_harmony_type = Enum_Harmony_type_Implied
		return
	case "alternate":
		*enum_harmony_type = Enum_Harmony_type_Alternate
		return
	default:
		return errUnkownEnum
	}
}

func (enum_harmony_type *Enum_Harmony_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Harmony_type_Explicit":
		*enum_harmony_type = Enum_Harmony_type_Explicit
	case "Enum_Harmony_type_Implied":
		*enum_harmony_type = Enum_Harmony_type_Implied
	case "Enum_Harmony_type_Alternate":
		*enum_harmony_type = Enum_Harmony_type_Alternate
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_harmony_type *Enum_Harmony_type) ToCodeString() (res string) {

	switch *enum_harmony_type {
	// insertion code per enum code
	case Enum_Harmony_type_Explicit:
		res = "Enum_Harmony_type_Explicit"
	case Enum_Harmony_type_Implied:
		res = "Enum_Harmony_type_Implied"
	case Enum_Harmony_type_Alternate:
		res = "Enum_Harmony_type_Alternate"
	}
	return
}

func (enum_harmony_type Enum_Harmony_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Harmony_type_Explicit")
	res = append(res, "Enum_Harmony_type_Implied")
	res = append(res, "Enum_Harmony_type_Alternate")

	return
}

func (enum_harmony_type Enum_Harmony_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "explicit")
	res = append(res, "implied")
	res = append(res, "alternate")

	return
}

// Utility function for Enum_Hole_closed_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_hole_closed_location Enum_Hole_closed_location) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_hole_closed_location {
	// insertion code per enum code
	case Enum_Hole_closed_location_Right:
		res = "right"
	case Enum_Hole_closed_location_Bottom:
		res = "bottom"
	case Enum_Hole_closed_location_Left:
		res = "left"
	case Enum_Hole_closed_location_Top:
		res = "top"
	}
	return
}

func (enum_hole_closed_location *Enum_Hole_closed_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enum_hole_closed_location = Enum_Hole_closed_location_Right
		return
	case "bottom":
		*enum_hole_closed_location = Enum_Hole_closed_location_Bottom
		return
	case "left":
		*enum_hole_closed_location = Enum_Hole_closed_location_Left
		return
	case "top":
		*enum_hole_closed_location = Enum_Hole_closed_location_Top
		return
	default:
		return errUnkownEnum
	}
}

func (enum_hole_closed_location *Enum_Hole_closed_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Hole_closed_location_Right":
		*enum_hole_closed_location = Enum_Hole_closed_location_Right
	case "Enum_Hole_closed_location_Bottom":
		*enum_hole_closed_location = Enum_Hole_closed_location_Bottom
	case "Enum_Hole_closed_location_Left":
		*enum_hole_closed_location = Enum_Hole_closed_location_Left
	case "Enum_Hole_closed_location_Top":
		*enum_hole_closed_location = Enum_Hole_closed_location_Top
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_hole_closed_location *Enum_Hole_closed_location) ToCodeString() (res string) {

	switch *enum_hole_closed_location {
	// insertion code per enum code
	case Enum_Hole_closed_location_Right:
		res = "Enum_Hole_closed_location_Right"
	case Enum_Hole_closed_location_Bottom:
		res = "Enum_Hole_closed_location_Bottom"
	case Enum_Hole_closed_location_Left:
		res = "Enum_Hole_closed_location_Left"
	case Enum_Hole_closed_location_Top:
		res = "Enum_Hole_closed_location_Top"
	}
	return
}

func (enum_hole_closed_location Enum_Hole_closed_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Hole_closed_location_Right")
	res = append(res, "Enum_Hole_closed_location_Bottom")
	res = append(res, "Enum_Hole_closed_location_Left")
	res = append(res, "Enum_Hole_closed_location_Top")

	return
}

func (enum_hole_closed_location Enum_Hole_closed_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for Enum_Hole_closed_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_hole_closed_value Enum_Hole_closed_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_hole_closed_value {
	// insertion code per enum code
	case Enum_Hole_closed_value_Yes:
		res = "yes"
	case Enum_Hole_closed_value_No:
		res = "no"
	case Enum_Hole_closed_value_Half:
		res = "half"
	}
	return
}

func (enum_hole_closed_value *Enum_Hole_closed_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enum_hole_closed_value = Enum_Hole_closed_value_Yes
		return
	case "no":
		*enum_hole_closed_value = Enum_Hole_closed_value_No
		return
	case "half":
		*enum_hole_closed_value = Enum_Hole_closed_value_Half
		return
	default:
		return errUnkownEnum
	}
}

func (enum_hole_closed_value *Enum_Hole_closed_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Hole_closed_value_Yes":
		*enum_hole_closed_value = Enum_Hole_closed_value_Yes
	case "Enum_Hole_closed_value_No":
		*enum_hole_closed_value = Enum_Hole_closed_value_No
	case "Enum_Hole_closed_value_Half":
		*enum_hole_closed_value = Enum_Hole_closed_value_Half
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_hole_closed_value *Enum_Hole_closed_value) ToCodeString() (res string) {

	switch *enum_hole_closed_value {
	// insertion code per enum code
	case Enum_Hole_closed_value_Yes:
		res = "Enum_Hole_closed_value_Yes"
	case Enum_Hole_closed_value_No:
		res = "Enum_Hole_closed_value_No"
	case Enum_Hole_closed_value_Half:
		res = "Enum_Hole_closed_value_Half"
	}
	return
}

func (enum_hole_closed_value Enum_Hole_closed_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Hole_closed_value_Yes")
	res = append(res, "Enum_Hole_closed_value_No")
	res = append(res, "Enum_Hole_closed_value_Half")

	return
}

func (enum_hole_closed_value Enum_Hole_closed_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for Enum_Kind_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_kind_value Enum_Kind_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_kind_value {
	// insertion code per enum code
	case Enum_Kind_value_Major:
		res = "major"
	case Enum_Kind_value_Minor:
		res = "minor"
	case Enum_Kind_value_Augmented:
		res = "augmented"
	case Enum_Kind_value_Diminished:
		res = "diminished"
	case Enum_Kind_value_Dominant:
		res = "dominant"
	case Enum_Kind_value_Major_seventh:
		res = "major-seventh"
	case Enum_Kind_value_Minor_seventh:
		res = "minor-seventh"
	case Enum_Kind_value_Diminished_seventh:
		res = "diminished-seventh"
	case Enum_Kind_value_Augmented_seventh:
		res = "augmented-seventh"
	case Enum_Kind_value_Half_diminished:
		res = "half-diminished"
	case Enum_Kind_value_Major_minor:
		res = "major-minor"
	case Enum_Kind_value_Major_sixth:
		res = "major-sixth"
	case Enum_Kind_value_Minor_sixth:
		res = "minor-sixth"
	case Enum_Kind_value_Dominant_ninth:
		res = "dominant-ninth"
	case Enum_Kind_value_Major_ninth:
		res = "major-ninth"
	case Enum_Kind_value_Minor_ninth:
		res = "minor-ninth"
	case Enum_Kind_value_Dominant_11th:
		res = "dominant-11th"
	case Enum_Kind_value_Major_11th:
		res = "major-11th"
	case Enum_Kind_value_Minor_11th:
		res = "minor-11th"
	case Enum_Kind_value_Dominant_13th:
		res = "dominant-13th"
	case Enum_Kind_value_Major_13th:
		res = "major-13th"
	case Enum_Kind_value_Minor_13th:
		res = "minor-13th"
	case Enum_Kind_value_Suspended_second:
		res = "suspended-second"
	case Enum_Kind_value_Suspended_fourth:
		res = "suspended-fourth"
	case Enum_Kind_value_Neapolitan:
		res = "Neapolitan"
	case Enum_Kind_value_Italian:
		res = "Italian"
	case Enum_Kind_value_French:
		res = "French"
	case Enum_Kind_value_German:
		res = "German"
	case Enum_Kind_value_Pedal:
		res = "pedal"
	case Enum_Kind_value_Power:
		res = "power"
	case Enum_Kind_value_Tristan:
		res = "Tristan"
	case Enum_Kind_value_Other:
		res = "other"
	case Enum_Kind_value_None:
		res = "none"
	}
	return
}

func (enum_kind_value *Enum_Kind_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enum_kind_value = Enum_Kind_value_Major
		return
	case "minor":
		*enum_kind_value = Enum_Kind_value_Minor
		return
	case "augmented":
		*enum_kind_value = Enum_Kind_value_Augmented
		return
	case "diminished":
		*enum_kind_value = Enum_Kind_value_Diminished
		return
	case "dominant":
		*enum_kind_value = Enum_Kind_value_Dominant
		return
	case "major-seventh":
		*enum_kind_value = Enum_Kind_value_Major_seventh
		return
	case "minor-seventh":
		*enum_kind_value = Enum_Kind_value_Minor_seventh
		return
	case "diminished-seventh":
		*enum_kind_value = Enum_Kind_value_Diminished_seventh
		return
	case "augmented-seventh":
		*enum_kind_value = Enum_Kind_value_Augmented_seventh
		return
	case "half-diminished":
		*enum_kind_value = Enum_Kind_value_Half_diminished
		return
	case "major-minor":
		*enum_kind_value = Enum_Kind_value_Major_minor
		return
	case "major-sixth":
		*enum_kind_value = Enum_Kind_value_Major_sixth
		return
	case "minor-sixth":
		*enum_kind_value = Enum_Kind_value_Minor_sixth
		return
	case "dominant-ninth":
		*enum_kind_value = Enum_Kind_value_Dominant_ninth
		return
	case "major-ninth":
		*enum_kind_value = Enum_Kind_value_Major_ninth
		return
	case "minor-ninth":
		*enum_kind_value = Enum_Kind_value_Minor_ninth
		return
	case "dominant-11th":
		*enum_kind_value = Enum_Kind_value_Dominant_11th
		return
	case "major-11th":
		*enum_kind_value = Enum_Kind_value_Major_11th
		return
	case "minor-11th":
		*enum_kind_value = Enum_Kind_value_Minor_11th
		return
	case "dominant-13th":
		*enum_kind_value = Enum_Kind_value_Dominant_13th
		return
	case "major-13th":
		*enum_kind_value = Enum_Kind_value_Major_13th
		return
	case "minor-13th":
		*enum_kind_value = Enum_Kind_value_Minor_13th
		return
	case "suspended-second":
		*enum_kind_value = Enum_Kind_value_Suspended_second
		return
	case "suspended-fourth":
		*enum_kind_value = Enum_Kind_value_Suspended_fourth
		return
	case "Neapolitan":
		*enum_kind_value = Enum_Kind_value_Neapolitan
		return
	case "Italian":
		*enum_kind_value = Enum_Kind_value_Italian
		return
	case "French":
		*enum_kind_value = Enum_Kind_value_French
		return
	case "German":
		*enum_kind_value = Enum_Kind_value_German
		return
	case "pedal":
		*enum_kind_value = Enum_Kind_value_Pedal
		return
	case "power":
		*enum_kind_value = Enum_Kind_value_Power
		return
	case "Tristan":
		*enum_kind_value = Enum_Kind_value_Tristan
		return
	case "other":
		*enum_kind_value = Enum_Kind_value_Other
		return
	case "none":
		*enum_kind_value = Enum_Kind_value_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_kind_value *Enum_Kind_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Kind_value_Major":
		*enum_kind_value = Enum_Kind_value_Major
	case "Enum_Kind_value_Minor":
		*enum_kind_value = Enum_Kind_value_Minor
	case "Enum_Kind_value_Augmented":
		*enum_kind_value = Enum_Kind_value_Augmented
	case "Enum_Kind_value_Diminished":
		*enum_kind_value = Enum_Kind_value_Diminished
	case "Enum_Kind_value_Dominant":
		*enum_kind_value = Enum_Kind_value_Dominant
	case "Enum_Kind_value_Major_seventh":
		*enum_kind_value = Enum_Kind_value_Major_seventh
	case "Enum_Kind_value_Minor_seventh":
		*enum_kind_value = Enum_Kind_value_Minor_seventh
	case "Enum_Kind_value_Diminished_seventh":
		*enum_kind_value = Enum_Kind_value_Diminished_seventh
	case "Enum_Kind_value_Augmented_seventh":
		*enum_kind_value = Enum_Kind_value_Augmented_seventh
	case "Enum_Kind_value_Half_diminished":
		*enum_kind_value = Enum_Kind_value_Half_diminished
	case "Enum_Kind_value_Major_minor":
		*enum_kind_value = Enum_Kind_value_Major_minor
	case "Enum_Kind_value_Major_sixth":
		*enum_kind_value = Enum_Kind_value_Major_sixth
	case "Enum_Kind_value_Minor_sixth":
		*enum_kind_value = Enum_Kind_value_Minor_sixth
	case "Enum_Kind_value_Dominant_ninth":
		*enum_kind_value = Enum_Kind_value_Dominant_ninth
	case "Enum_Kind_value_Major_ninth":
		*enum_kind_value = Enum_Kind_value_Major_ninth
	case "Enum_Kind_value_Minor_ninth":
		*enum_kind_value = Enum_Kind_value_Minor_ninth
	case "Enum_Kind_value_Dominant_11th":
		*enum_kind_value = Enum_Kind_value_Dominant_11th
	case "Enum_Kind_value_Major_11th":
		*enum_kind_value = Enum_Kind_value_Major_11th
	case "Enum_Kind_value_Minor_11th":
		*enum_kind_value = Enum_Kind_value_Minor_11th
	case "Enum_Kind_value_Dominant_13th":
		*enum_kind_value = Enum_Kind_value_Dominant_13th
	case "Enum_Kind_value_Major_13th":
		*enum_kind_value = Enum_Kind_value_Major_13th
	case "Enum_Kind_value_Minor_13th":
		*enum_kind_value = Enum_Kind_value_Minor_13th
	case "Enum_Kind_value_Suspended_second":
		*enum_kind_value = Enum_Kind_value_Suspended_second
	case "Enum_Kind_value_Suspended_fourth":
		*enum_kind_value = Enum_Kind_value_Suspended_fourth
	case "Enum_Kind_value_Neapolitan":
		*enum_kind_value = Enum_Kind_value_Neapolitan
	case "Enum_Kind_value_Italian":
		*enum_kind_value = Enum_Kind_value_Italian
	case "Enum_Kind_value_French":
		*enum_kind_value = Enum_Kind_value_French
	case "Enum_Kind_value_German":
		*enum_kind_value = Enum_Kind_value_German
	case "Enum_Kind_value_Pedal":
		*enum_kind_value = Enum_Kind_value_Pedal
	case "Enum_Kind_value_Power":
		*enum_kind_value = Enum_Kind_value_Power
	case "Enum_Kind_value_Tristan":
		*enum_kind_value = Enum_Kind_value_Tristan
	case "Enum_Kind_value_Other":
		*enum_kind_value = Enum_Kind_value_Other
	case "Enum_Kind_value_None":
		*enum_kind_value = Enum_Kind_value_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_kind_value *Enum_Kind_value) ToCodeString() (res string) {

	switch *enum_kind_value {
	// insertion code per enum code
	case Enum_Kind_value_Major:
		res = "Enum_Kind_value_Major"
	case Enum_Kind_value_Minor:
		res = "Enum_Kind_value_Minor"
	case Enum_Kind_value_Augmented:
		res = "Enum_Kind_value_Augmented"
	case Enum_Kind_value_Diminished:
		res = "Enum_Kind_value_Diminished"
	case Enum_Kind_value_Dominant:
		res = "Enum_Kind_value_Dominant"
	case Enum_Kind_value_Major_seventh:
		res = "Enum_Kind_value_Major_seventh"
	case Enum_Kind_value_Minor_seventh:
		res = "Enum_Kind_value_Minor_seventh"
	case Enum_Kind_value_Diminished_seventh:
		res = "Enum_Kind_value_Diminished_seventh"
	case Enum_Kind_value_Augmented_seventh:
		res = "Enum_Kind_value_Augmented_seventh"
	case Enum_Kind_value_Half_diminished:
		res = "Enum_Kind_value_Half_diminished"
	case Enum_Kind_value_Major_minor:
		res = "Enum_Kind_value_Major_minor"
	case Enum_Kind_value_Major_sixth:
		res = "Enum_Kind_value_Major_sixth"
	case Enum_Kind_value_Minor_sixth:
		res = "Enum_Kind_value_Minor_sixth"
	case Enum_Kind_value_Dominant_ninth:
		res = "Enum_Kind_value_Dominant_ninth"
	case Enum_Kind_value_Major_ninth:
		res = "Enum_Kind_value_Major_ninth"
	case Enum_Kind_value_Minor_ninth:
		res = "Enum_Kind_value_Minor_ninth"
	case Enum_Kind_value_Dominant_11th:
		res = "Enum_Kind_value_Dominant_11th"
	case Enum_Kind_value_Major_11th:
		res = "Enum_Kind_value_Major_11th"
	case Enum_Kind_value_Minor_11th:
		res = "Enum_Kind_value_Minor_11th"
	case Enum_Kind_value_Dominant_13th:
		res = "Enum_Kind_value_Dominant_13th"
	case Enum_Kind_value_Major_13th:
		res = "Enum_Kind_value_Major_13th"
	case Enum_Kind_value_Minor_13th:
		res = "Enum_Kind_value_Minor_13th"
	case Enum_Kind_value_Suspended_second:
		res = "Enum_Kind_value_Suspended_second"
	case Enum_Kind_value_Suspended_fourth:
		res = "Enum_Kind_value_Suspended_fourth"
	case Enum_Kind_value_Neapolitan:
		res = "Enum_Kind_value_Neapolitan"
	case Enum_Kind_value_Italian:
		res = "Enum_Kind_value_Italian"
	case Enum_Kind_value_French:
		res = "Enum_Kind_value_French"
	case Enum_Kind_value_German:
		res = "Enum_Kind_value_German"
	case Enum_Kind_value_Pedal:
		res = "Enum_Kind_value_Pedal"
	case Enum_Kind_value_Power:
		res = "Enum_Kind_value_Power"
	case Enum_Kind_value_Tristan:
		res = "Enum_Kind_value_Tristan"
	case Enum_Kind_value_Other:
		res = "Enum_Kind_value_Other"
	case Enum_Kind_value_None:
		res = "Enum_Kind_value_None"
	}
	return
}

func (enum_kind_value Enum_Kind_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Kind_value_Major")
	res = append(res, "Enum_Kind_value_Minor")
	res = append(res, "Enum_Kind_value_Augmented")
	res = append(res, "Enum_Kind_value_Diminished")
	res = append(res, "Enum_Kind_value_Dominant")
	res = append(res, "Enum_Kind_value_Major_seventh")
	res = append(res, "Enum_Kind_value_Minor_seventh")
	res = append(res, "Enum_Kind_value_Diminished_seventh")
	res = append(res, "Enum_Kind_value_Augmented_seventh")
	res = append(res, "Enum_Kind_value_Half_diminished")
	res = append(res, "Enum_Kind_value_Major_minor")
	res = append(res, "Enum_Kind_value_Major_sixth")
	res = append(res, "Enum_Kind_value_Minor_sixth")
	res = append(res, "Enum_Kind_value_Dominant_ninth")
	res = append(res, "Enum_Kind_value_Major_ninth")
	res = append(res, "Enum_Kind_value_Minor_ninth")
	res = append(res, "Enum_Kind_value_Dominant_11th")
	res = append(res, "Enum_Kind_value_Major_11th")
	res = append(res, "Enum_Kind_value_Minor_11th")
	res = append(res, "Enum_Kind_value_Dominant_13th")
	res = append(res, "Enum_Kind_value_Major_13th")
	res = append(res, "Enum_Kind_value_Minor_13th")
	res = append(res, "Enum_Kind_value_Suspended_second")
	res = append(res, "Enum_Kind_value_Suspended_fourth")
	res = append(res, "Enum_Kind_value_Neapolitan")
	res = append(res, "Enum_Kind_value_Italian")
	res = append(res, "Enum_Kind_value_French")
	res = append(res, "Enum_Kind_value_German")
	res = append(res, "Enum_Kind_value_Pedal")
	res = append(res, "Enum_Kind_value_Power")
	res = append(res, "Enum_Kind_value_Tristan")
	res = append(res, "Enum_Kind_value_Other")
	res = append(res, "Enum_Kind_value_None")

	return
}

func (enum_kind_value Enum_Kind_value) CodeValues() (res []string) {

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

// Utility function for Enum_Left_center_right
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_left_center_right Enum_Left_center_right) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_left_center_right {
	// insertion code per enum code
	case Enum_Left_center_right_Left:
		res = "left"
	case Enum_Left_center_right_Center:
		res = "center"
	case Enum_Left_center_right_Right:
		res = "right"
	}
	return
}

func (enum_left_center_right *Enum_Left_center_right) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enum_left_center_right = Enum_Left_center_right_Left
		return
	case "center":
		*enum_left_center_right = Enum_Left_center_right_Center
		return
	case "right":
		*enum_left_center_right = Enum_Left_center_right_Right
		return
	default:
		return errUnkownEnum
	}
}

func (enum_left_center_right *Enum_Left_center_right) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Left_center_right_Left":
		*enum_left_center_right = Enum_Left_center_right_Left
	case "Enum_Left_center_right_Center":
		*enum_left_center_right = Enum_Left_center_right_Center
	case "Enum_Left_center_right_Right":
		*enum_left_center_right = Enum_Left_center_right_Right
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_left_center_right *Enum_Left_center_right) ToCodeString() (res string) {

	switch *enum_left_center_right {
	// insertion code per enum code
	case Enum_Left_center_right_Left:
		res = "Enum_Left_center_right_Left"
	case Enum_Left_center_right_Center:
		res = "Enum_Left_center_right_Center"
	case Enum_Left_center_right_Right:
		res = "Enum_Left_center_right_Right"
	}
	return
}

func (enum_left_center_right Enum_Left_center_right) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Left_center_right_Left")
	res = append(res, "Enum_Left_center_right_Center")
	res = append(res, "Enum_Left_center_right_Right")

	return
}

func (enum_left_center_right Enum_Left_center_right) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "center")
	res = append(res, "right")

	return
}

// Utility function for Enum_Left_right
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_left_right Enum_Left_right) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_left_right {
	// insertion code per enum code
	case Enum_Left_right_Left:
		res = "left"
	case Enum_Left_right_Right:
		res = "right"
	}
	return
}

func (enum_left_right *Enum_Left_right) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enum_left_right = Enum_Left_right_Left
		return
	case "right":
		*enum_left_right = Enum_Left_right_Right
		return
	default:
		return errUnkownEnum
	}
}

func (enum_left_right *Enum_Left_right) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Left_right_Left":
		*enum_left_right = Enum_Left_right_Left
	case "Enum_Left_right_Right":
		*enum_left_right = Enum_Left_right_Right
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_left_right *Enum_Left_right) ToCodeString() (res string) {

	switch *enum_left_right {
	// insertion code per enum code
	case Enum_Left_right_Left:
		res = "Enum_Left_right_Left"
	case Enum_Left_right_Right:
		res = "Enum_Left_right_Right"
	}
	return
}

func (enum_left_right Enum_Left_right) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Left_right_Left")
	res = append(res, "Enum_Left_right_Right")

	return
}

func (enum_left_right Enum_Left_right) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for Enum_Line_end
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_line_end Enum_Line_end) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_line_end {
	// insertion code per enum code
	case Enum_Line_end_Up:
		res = "up"
	case Enum_Line_end_Down:
		res = "down"
	case Enum_Line_end_Both:
		res = "both"
	case Enum_Line_end_Arrow:
		res = "arrow"
	case Enum_Line_end_None:
		res = "none"
	}
	return
}

func (enum_line_end *Enum_Line_end) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enum_line_end = Enum_Line_end_Up
		return
	case "down":
		*enum_line_end = Enum_Line_end_Down
		return
	case "both":
		*enum_line_end = Enum_Line_end_Both
		return
	case "arrow":
		*enum_line_end = Enum_Line_end_Arrow
		return
	case "none":
		*enum_line_end = Enum_Line_end_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_line_end *Enum_Line_end) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Line_end_Up":
		*enum_line_end = Enum_Line_end_Up
	case "Enum_Line_end_Down":
		*enum_line_end = Enum_Line_end_Down
	case "Enum_Line_end_Both":
		*enum_line_end = Enum_Line_end_Both
	case "Enum_Line_end_Arrow":
		*enum_line_end = Enum_Line_end_Arrow
	case "Enum_Line_end_None":
		*enum_line_end = Enum_Line_end_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_line_end *Enum_Line_end) ToCodeString() (res string) {

	switch *enum_line_end {
	// insertion code per enum code
	case Enum_Line_end_Up:
		res = "Enum_Line_end_Up"
	case Enum_Line_end_Down:
		res = "Enum_Line_end_Down"
	case Enum_Line_end_Both:
		res = "Enum_Line_end_Both"
	case Enum_Line_end_Arrow:
		res = "Enum_Line_end_Arrow"
	case Enum_Line_end_None:
		res = "Enum_Line_end_None"
	}
	return
}

func (enum_line_end Enum_Line_end) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Line_end_Up")
	res = append(res, "Enum_Line_end_Down")
	res = append(res, "Enum_Line_end_Both")
	res = append(res, "Enum_Line_end_Arrow")
	res = append(res, "Enum_Line_end_None")

	return
}

func (enum_line_end Enum_Line_end) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "both")
	res = append(res, "arrow")
	res = append(res, "none")

	return
}

// Utility function for Enum_Line_length
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_line_length Enum_Line_length) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_line_length {
	// insertion code per enum code
	case Enum_Line_length_Short:
		res = "short"
	case Enum_Line_length_Medium:
		res = "medium"
	case Enum_Line_length_Long:
		res = "long"
	}
	return
}

func (enum_line_length *Enum_Line_length) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "short":
		*enum_line_length = Enum_Line_length_Short
		return
	case "medium":
		*enum_line_length = Enum_Line_length_Medium
		return
	case "long":
		*enum_line_length = Enum_Line_length_Long
		return
	default:
		return errUnkownEnum
	}
}

func (enum_line_length *Enum_Line_length) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Line_length_Short":
		*enum_line_length = Enum_Line_length_Short
	case "Enum_Line_length_Medium":
		*enum_line_length = Enum_Line_length_Medium
	case "Enum_Line_length_Long":
		*enum_line_length = Enum_Line_length_Long
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_line_length *Enum_Line_length) ToCodeString() (res string) {

	switch *enum_line_length {
	// insertion code per enum code
	case Enum_Line_length_Short:
		res = "Enum_Line_length_Short"
	case Enum_Line_length_Medium:
		res = "Enum_Line_length_Medium"
	case Enum_Line_length_Long:
		res = "Enum_Line_length_Long"
	}
	return
}

func (enum_line_length Enum_Line_length) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Line_length_Short")
	res = append(res, "Enum_Line_length_Medium")
	res = append(res, "Enum_Line_length_Long")

	return
}

func (enum_line_length Enum_Line_length) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "short")
	res = append(res, "medium")
	res = append(res, "long")

	return
}

// Utility function for Enum_Line_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_line_shape Enum_Line_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_line_shape {
	// insertion code per enum code
	case Enum_Line_shape_Straight:
		res = "straight"
	case Enum_Line_shape_Curved:
		res = "curved"
	}
	return
}

func (enum_line_shape *Enum_Line_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "straight":
		*enum_line_shape = Enum_Line_shape_Straight
		return
	case "curved":
		*enum_line_shape = Enum_Line_shape_Curved
		return
	default:
		return errUnkownEnum
	}
}

func (enum_line_shape *Enum_Line_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Line_shape_Straight":
		*enum_line_shape = Enum_Line_shape_Straight
	case "Enum_Line_shape_Curved":
		*enum_line_shape = Enum_Line_shape_Curved
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_line_shape *Enum_Line_shape) ToCodeString() (res string) {

	switch *enum_line_shape {
	// insertion code per enum code
	case Enum_Line_shape_Straight:
		res = "Enum_Line_shape_Straight"
	case Enum_Line_shape_Curved:
		res = "Enum_Line_shape_Curved"
	}
	return
}

func (enum_line_shape Enum_Line_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Line_shape_Straight")
	res = append(res, "Enum_Line_shape_Curved")

	return
}

func (enum_line_shape Enum_Line_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "straight")
	res = append(res, "curved")

	return
}

// Utility function for Enum_Line_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_line_type Enum_Line_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_line_type {
	// insertion code per enum code
	case Enum_Line_type_Solid:
		res = "solid"
	case Enum_Line_type_Dashed:
		res = "dashed"
	case Enum_Line_type_Dotted:
		res = "dotted"
	case Enum_Line_type_Wavy:
		res = "wavy"
	}
	return
}

func (enum_line_type *Enum_Line_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "solid":
		*enum_line_type = Enum_Line_type_Solid
		return
	case "dashed":
		*enum_line_type = Enum_Line_type_Dashed
		return
	case "dotted":
		*enum_line_type = Enum_Line_type_Dotted
		return
	case "wavy":
		*enum_line_type = Enum_Line_type_Wavy
		return
	default:
		return errUnkownEnum
	}
}

func (enum_line_type *Enum_Line_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Line_type_Solid":
		*enum_line_type = Enum_Line_type_Solid
	case "Enum_Line_type_Dashed":
		*enum_line_type = Enum_Line_type_Dashed
	case "Enum_Line_type_Dotted":
		*enum_line_type = Enum_Line_type_Dotted
	case "Enum_Line_type_Wavy":
		*enum_line_type = Enum_Line_type_Wavy
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_line_type *Enum_Line_type) ToCodeString() (res string) {

	switch *enum_line_type {
	// insertion code per enum code
	case Enum_Line_type_Solid:
		res = "Enum_Line_type_Solid"
	case Enum_Line_type_Dashed:
		res = "Enum_Line_type_Dashed"
	case Enum_Line_type_Dotted:
		res = "Enum_Line_type_Dotted"
	case Enum_Line_type_Wavy:
		res = "Enum_Line_type_Wavy"
	}
	return
}

func (enum_line_type Enum_Line_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Line_type_Solid")
	res = append(res, "Enum_Line_type_Dashed")
	res = append(res, "Enum_Line_type_Dotted")
	res = append(res, "Enum_Line_type_Wavy")

	return
}

func (enum_line_type Enum_Line_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "solid")
	res = append(res, "dashed")
	res = append(res, "dotted")
	res = append(res, "wavy")

	return
}

// Utility function for Enum_Line_width_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_line_width_type Enum_Line_width_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_line_width_type {
	// insertion code per enum code
	}
	return
}

func (enum_line_width_type *Enum_Line_width_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_line_width_type *Enum_Line_width_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_line_width_type *Enum_Line_width_type) ToCodeString() (res string) {

	switch *enum_line_width_type {
	// insertion code per enum code
	}
	return
}

func (enum_line_width_type Enum_Line_width_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_line_width_type Enum_Line_width_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Margin_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_margin_type Enum_Margin_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_margin_type {
	// insertion code per enum code
	case Enum_Margin_type_Odd:
		res = "odd"
	case Enum_Margin_type_Even:
		res = "even"
	case Enum_Margin_type_Both:
		res = "both"
	}
	return
}

func (enum_margin_type *Enum_Margin_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "odd":
		*enum_margin_type = Enum_Margin_type_Odd
		return
	case "even":
		*enum_margin_type = Enum_Margin_type_Even
		return
	case "both":
		*enum_margin_type = Enum_Margin_type_Both
		return
	default:
		return errUnkownEnum
	}
}

func (enum_margin_type *Enum_Margin_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Margin_type_Odd":
		*enum_margin_type = Enum_Margin_type_Odd
	case "Enum_Margin_type_Even":
		*enum_margin_type = Enum_Margin_type_Even
	case "Enum_Margin_type_Both":
		*enum_margin_type = Enum_Margin_type_Both
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_margin_type *Enum_Margin_type) ToCodeString() (res string) {

	switch *enum_margin_type {
	// insertion code per enum code
	case Enum_Margin_type_Odd:
		res = "Enum_Margin_type_Odd"
	case Enum_Margin_type_Even:
		res = "Enum_Margin_type_Even"
	case Enum_Margin_type_Both:
		res = "Enum_Margin_type_Both"
	}
	return
}

func (enum_margin_type Enum_Margin_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Margin_type_Odd")
	res = append(res, "Enum_Margin_type_Even")
	res = append(res, "Enum_Margin_type_Both")

	return
}

func (enum_margin_type Enum_Margin_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "odd")
	res = append(res, "even")
	res = append(res, "both")

	return
}

// Utility function for Enum_Measure_numbering_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_measure_numbering_value Enum_Measure_numbering_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_measure_numbering_value {
	// insertion code per enum code
	case Enum_Measure_numbering_value_None:
		res = "none"
	case Enum_Measure_numbering_value_Measure:
		res = "measure"
	case Enum_Measure_numbering_value_System:
		res = "system"
	}
	return
}

func (enum_measure_numbering_value *Enum_Measure_numbering_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_None
		return
	case "measure":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_Measure
		return
	case "system":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_System
		return
	default:
		return errUnkownEnum
	}
}

func (enum_measure_numbering_value *Enum_Measure_numbering_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Measure_numbering_value_None":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_None
	case "Enum_Measure_numbering_value_Measure":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_Measure
	case "Enum_Measure_numbering_value_System":
		*enum_measure_numbering_value = Enum_Measure_numbering_value_System
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_measure_numbering_value *Enum_Measure_numbering_value) ToCodeString() (res string) {

	switch *enum_measure_numbering_value {
	// insertion code per enum code
	case Enum_Measure_numbering_value_None:
		res = "Enum_Measure_numbering_value_None"
	case Enum_Measure_numbering_value_Measure:
		res = "Enum_Measure_numbering_value_Measure"
	case Enum_Measure_numbering_value_System:
		res = "Enum_Measure_numbering_value_System"
	}
	return
}

func (enum_measure_numbering_value Enum_Measure_numbering_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Measure_numbering_value_None")
	res = append(res, "Enum_Measure_numbering_value_Measure")
	res = append(res, "Enum_Measure_numbering_value_System")

	return
}

func (enum_measure_numbering_value Enum_Measure_numbering_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "measure")
	res = append(res, "system")

	return
}

// Utility function for Enum_Measure_text
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_measure_text Enum_Measure_text) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_measure_text {
	// insertion code per enum code
	}
	return
}

func (enum_measure_text *Enum_Measure_text) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_measure_text *Enum_Measure_text) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_measure_text *Enum_Measure_text) ToCodeString() (res string) {

	switch *enum_measure_text {
	// insertion code per enum code
	}
	return
}

func (enum_measure_text Enum_Measure_text) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_measure_text Enum_Measure_text) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Membrane_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_membrane_value Enum_Membrane_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_membrane_value {
	// insertion code per enum code
	case Enum_Membrane_value_Bass_drum:
		res = "bass drum"
	case Enum_Membrane_value_Bass_drum_on_side:
		res = "bass drum on side"
	case Enum_Membrane_value_Bongos:
		res = "bongos"
	case Enum_Membrane_value_Chinese_tomtom:
		res = "Chinese tomtom"
	case Enum_Membrane_value_Conga_drum:
		res = "conga drum"
	case Enum_Membrane_value_Cuica:
		res = "cuica"
	case Enum_Membrane_value_Goblet_drum:
		res = "goblet drum"
	case Enum_Membrane_value_Indo_American_tomtom:
		res = "Indo-American tomtom"
	case Enum_Membrane_value_Japanese_tomtom:
		res = "Japanese tomtom"
	case Enum_Membrane_value_Military_drum:
		res = "military drum"
	case Enum_Membrane_value_Snare_drum:
		res = "snare drum"
	case Enum_Membrane_value_Snare_drum_snares_off:
		res = "snare drum snares off"
	case Enum_Membrane_value_Tabla:
		res = "tabla"
	case Enum_Membrane_value_Tambourine:
		res = "tambourine"
	case Enum_Membrane_value_Tenor_drum:
		res = "tenor drum"
	case Enum_Membrane_value_Timbales:
		res = "timbales"
	case Enum_Membrane_value_Tomtom:
		res = "tomtom"
	}
	return
}

func (enum_membrane_value *Enum_Membrane_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*enum_membrane_value = Enum_Membrane_value_Bass_drum
		return
	case "bass drum on side":
		*enum_membrane_value = Enum_Membrane_value_Bass_drum_on_side
		return
	case "bongos":
		*enum_membrane_value = Enum_Membrane_value_Bongos
		return
	case "Chinese tomtom":
		*enum_membrane_value = Enum_Membrane_value_Chinese_tomtom
		return
	case "conga drum":
		*enum_membrane_value = Enum_Membrane_value_Conga_drum
		return
	case "cuica":
		*enum_membrane_value = Enum_Membrane_value_Cuica
		return
	case "goblet drum":
		*enum_membrane_value = Enum_Membrane_value_Goblet_drum
		return
	case "Indo-American tomtom":
		*enum_membrane_value = Enum_Membrane_value_Indo_American_tomtom
		return
	case "Japanese tomtom":
		*enum_membrane_value = Enum_Membrane_value_Japanese_tomtom
		return
	case "military drum":
		*enum_membrane_value = Enum_Membrane_value_Military_drum
		return
	case "snare drum":
		*enum_membrane_value = Enum_Membrane_value_Snare_drum
		return
	case "snare drum snares off":
		*enum_membrane_value = Enum_Membrane_value_Snare_drum_snares_off
		return
	case "tabla":
		*enum_membrane_value = Enum_Membrane_value_Tabla
		return
	case "tambourine":
		*enum_membrane_value = Enum_Membrane_value_Tambourine
		return
	case "tenor drum":
		*enum_membrane_value = Enum_Membrane_value_Tenor_drum
		return
	case "timbales":
		*enum_membrane_value = Enum_Membrane_value_Timbales
		return
	case "tomtom":
		*enum_membrane_value = Enum_Membrane_value_Tomtom
		return
	default:
		return errUnkownEnum
	}
}

func (enum_membrane_value *Enum_Membrane_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Membrane_value_Bass_drum":
		*enum_membrane_value = Enum_Membrane_value_Bass_drum
	case "Enum_Membrane_value_Bass_drum_on_side":
		*enum_membrane_value = Enum_Membrane_value_Bass_drum_on_side
	case "Enum_Membrane_value_Bongos":
		*enum_membrane_value = Enum_Membrane_value_Bongos
	case "Enum_Membrane_value_Chinese_tomtom":
		*enum_membrane_value = Enum_Membrane_value_Chinese_tomtom
	case "Enum_Membrane_value_Conga_drum":
		*enum_membrane_value = Enum_Membrane_value_Conga_drum
	case "Enum_Membrane_value_Cuica":
		*enum_membrane_value = Enum_Membrane_value_Cuica
	case "Enum_Membrane_value_Goblet_drum":
		*enum_membrane_value = Enum_Membrane_value_Goblet_drum
	case "Enum_Membrane_value_Indo_American_tomtom":
		*enum_membrane_value = Enum_Membrane_value_Indo_American_tomtom
	case "Enum_Membrane_value_Japanese_tomtom":
		*enum_membrane_value = Enum_Membrane_value_Japanese_tomtom
	case "Enum_Membrane_value_Military_drum":
		*enum_membrane_value = Enum_Membrane_value_Military_drum
	case "Enum_Membrane_value_Snare_drum":
		*enum_membrane_value = Enum_Membrane_value_Snare_drum
	case "Enum_Membrane_value_Snare_drum_snares_off":
		*enum_membrane_value = Enum_Membrane_value_Snare_drum_snares_off
	case "Enum_Membrane_value_Tabla":
		*enum_membrane_value = Enum_Membrane_value_Tabla
	case "Enum_Membrane_value_Tambourine":
		*enum_membrane_value = Enum_Membrane_value_Tambourine
	case "Enum_Membrane_value_Tenor_drum":
		*enum_membrane_value = Enum_Membrane_value_Tenor_drum
	case "Enum_Membrane_value_Timbales":
		*enum_membrane_value = Enum_Membrane_value_Timbales
	case "Enum_Membrane_value_Tomtom":
		*enum_membrane_value = Enum_Membrane_value_Tomtom
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_membrane_value *Enum_Membrane_value) ToCodeString() (res string) {

	switch *enum_membrane_value {
	// insertion code per enum code
	case Enum_Membrane_value_Bass_drum:
		res = "Enum_Membrane_value_Bass_drum"
	case Enum_Membrane_value_Bass_drum_on_side:
		res = "Enum_Membrane_value_Bass_drum_on_side"
	case Enum_Membrane_value_Bongos:
		res = "Enum_Membrane_value_Bongos"
	case Enum_Membrane_value_Chinese_tomtom:
		res = "Enum_Membrane_value_Chinese_tomtom"
	case Enum_Membrane_value_Conga_drum:
		res = "Enum_Membrane_value_Conga_drum"
	case Enum_Membrane_value_Cuica:
		res = "Enum_Membrane_value_Cuica"
	case Enum_Membrane_value_Goblet_drum:
		res = "Enum_Membrane_value_Goblet_drum"
	case Enum_Membrane_value_Indo_American_tomtom:
		res = "Enum_Membrane_value_Indo_American_tomtom"
	case Enum_Membrane_value_Japanese_tomtom:
		res = "Enum_Membrane_value_Japanese_tomtom"
	case Enum_Membrane_value_Military_drum:
		res = "Enum_Membrane_value_Military_drum"
	case Enum_Membrane_value_Snare_drum:
		res = "Enum_Membrane_value_Snare_drum"
	case Enum_Membrane_value_Snare_drum_snares_off:
		res = "Enum_Membrane_value_Snare_drum_snares_off"
	case Enum_Membrane_value_Tabla:
		res = "Enum_Membrane_value_Tabla"
	case Enum_Membrane_value_Tambourine:
		res = "Enum_Membrane_value_Tambourine"
	case Enum_Membrane_value_Tenor_drum:
		res = "Enum_Membrane_value_Tenor_drum"
	case Enum_Membrane_value_Timbales:
		res = "Enum_Membrane_value_Timbales"
	case Enum_Membrane_value_Tomtom:
		res = "Enum_Membrane_value_Tomtom"
	}
	return
}

func (enum_membrane_value Enum_Membrane_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Membrane_value_Bass_drum")
	res = append(res, "Enum_Membrane_value_Bass_drum_on_side")
	res = append(res, "Enum_Membrane_value_Bongos")
	res = append(res, "Enum_Membrane_value_Chinese_tomtom")
	res = append(res, "Enum_Membrane_value_Conga_drum")
	res = append(res, "Enum_Membrane_value_Cuica")
	res = append(res, "Enum_Membrane_value_Goblet_drum")
	res = append(res, "Enum_Membrane_value_Indo_American_tomtom")
	res = append(res, "Enum_Membrane_value_Japanese_tomtom")
	res = append(res, "Enum_Membrane_value_Military_drum")
	res = append(res, "Enum_Membrane_value_Snare_drum")
	res = append(res, "Enum_Membrane_value_Snare_drum_snares_off")
	res = append(res, "Enum_Membrane_value_Tabla")
	res = append(res, "Enum_Membrane_value_Tambourine")
	res = append(res, "Enum_Membrane_value_Tenor_drum")
	res = append(res, "Enum_Membrane_value_Timbales")
	res = append(res, "Enum_Membrane_value_Tomtom")

	return
}

func (enum_membrane_value Enum_Membrane_value) CodeValues() (res []string) {

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

// Utility function for Enum_Metal_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_metal_value Enum_Metal_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_metal_value {
	// insertion code per enum code
	case Enum_Metal_value_Agogo:
		res = "agogo"
	case Enum_Metal_value_Almglocken:
		res = "almglocken"
	case Enum_Metal_value_Bell:
		res = "bell"
	case Enum_Metal_value_Bell_plate:
		res = "bell plate"
	case Enum_Metal_value_Bell_tree:
		res = "bell tree"
	case Enum_Metal_value_Brake_drum:
		res = "brake drum"
	case Enum_Metal_value_Cencerro:
		res = "cencerro"
	case Enum_Metal_value_Chain_rattle:
		res = "chain rattle"
	case Enum_Metal_value_Chinese_cymbal:
		res = "Chinese cymbal"
	case Enum_Metal_value_Cowbell:
		res = "cowbell"
	case Enum_Metal_value_Crash_cymbals:
		res = "crash cymbals"
	case Enum_Metal_value_Crotale:
		res = "crotale"
	case Enum_Metal_value_Cymbal_tongs:
		res = "cymbal tongs"
	case Enum_Metal_value_Domed_gong:
		res = "domed gong"
	case Enum_Metal_value_Finger_cymbals:
		res = "finger cymbals"
	case Enum_Metal_value_Flexatone:
		res = "flexatone"
	case Enum_Metal_value_Gong:
		res = "gong"
	case Enum_Metal_value_Hi_hat:
		res = "hi-hat"
	case Enum_Metal_value_High_hat_cymbals:
		res = "high-hat cymbals"
	case Enum_Metal_value_Handbell:
		res = "handbell"
	case Enum_Metal_value_Jaw_harp:
		res = "jaw harp"
	case Enum_Metal_value_Jingle_bells:
		res = "jingle bells"
	case Enum_Metal_value_Musical_saw:
		res = "musical saw"
	case Enum_Metal_value_Shell_bells:
		res = "shell bells"
	case Enum_Metal_value_Sistrum:
		res = "sistrum"
	case Enum_Metal_value_Sizzle_cymbal:
		res = "sizzle cymbal"
	case Enum_Metal_value_Sleigh_bells:
		res = "sleigh bells"
	case Enum_Metal_value_Suspended_cymbal:
		res = "suspended cymbal"
	case Enum_Metal_value_Tam_tam:
		res = "tam tam"
	case Enum_Metal_value_Tam_tam_with_beater:
		res = "tam tam with beater"
	case Enum_Metal_value_Triangle:
		res = "triangle"
	case Enum_Metal_value_Vietnamese_hat:
		res = "Vietnamese hat"
	}
	return
}

func (enum_metal_value *Enum_Metal_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "agogo":
		*enum_metal_value = Enum_Metal_value_Agogo
		return
	case "almglocken":
		*enum_metal_value = Enum_Metal_value_Almglocken
		return
	case "bell":
		*enum_metal_value = Enum_Metal_value_Bell
		return
	case "bell plate":
		*enum_metal_value = Enum_Metal_value_Bell_plate
		return
	case "bell tree":
		*enum_metal_value = Enum_Metal_value_Bell_tree
		return
	case "brake drum":
		*enum_metal_value = Enum_Metal_value_Brake_drum
		return
	case "cencerro":
		*enum_metal_value = Enum_Metal_value_Cencerro
		return
	case "chain rattle":
		*enum_metal_value = Enum_Metal_value_Chain_rattle
		return
	case "Chinese cymbal":
		*enum_metal_value = Enum_Metal_value_Chinese_cymbal
		return
	case "cowbell":
		*enum_metal_value = Enum_Metal_value_Cowbell
		return
	case "crash cymbals":
		*enum_metal_value = Enum_Metal_value_Crash_cymbals
		return
	case "crotale":
		*enum_metal_value = Enum_Metal_value_Crotale
		return
	case "cymbal tongs":
		*enum_metal_value = Enum_Metal_value_Cymbal_tongs
		return
	case "domed gong":
		*enum_metal_value = Enum_Metal_value_Domed_gong
		return
	case "finger cymbals":
		*enum_metal_value = Enum_Metal_value_Finger_cymbals
		return
	case "flexatone":
		*enum_metal_value = Enum_Metal_value_Flexatone
		return
	case "gong":
		*enum_metal_value = Enum_Metal_value_Gong
		return
	case "hi-hat":
		*enum_metal_value = Enum_Metal_value_Hi_hat
		return
	case "high-hat cymbals":
		*enum_metal_value = Enum_Metal_value_High_hat_cymbals
		return
	case "handbell":
		*enum_metal_value = Enum_Metal_value_Handbell
		return
	case "jaw harp":
		*enum_metal_value = Enum_Metal_value_Jaw_harp
		return
	case "jingle bells":
		*enum_metal_value = Enum_Metal_value_Jingle_bells
		return
	case "musical saw":
		*enum_metal_value = Enum_Metal_value_Musical_saw
		return
	case "shell bells":
		*enum_metal_value = Enum_Metal_value_Shell_bells
		return
	case "sistrum":
		*enum_metal_value = Enum_Metal_value_Sistrum
		return
	case "sizzle cymbal":
		*enum_metal_value = Enum_Metal_value_Sizzle_cymbal
		return
	case "sleigh bells":
		*enum_metal_value = Enum_Metal_value_Sleigh_bells
		return
	case "suspended cymbal":
		*enum_metal_value = Enum_Metal_value_Suspended_cymbal
		return
	case "tam tam":
		*enum_metal_value = Enum_Metal_value_Tam_tam
		return
	case "tam tam with beater":
		*enum_metal_value = Enum_Metal_value_Tam_tam_with_beater
		return
	case "triangle":
		*enum_metal_value = Enum_Metal_value_Triangle
		return
	case "Vietnamese hat":
		*enum_metal_value = Enum_Metal_value_Vietnamese_hat
		return
	default:
		return errUnkownEnum
	}
}

func (enum_metal_value *Enum_Metal_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Metal_value_Agogo":
		*enum_metal_value = Enum_Metal_value_Agogo
	case "Enum_Metal_value_Almglocken":
		*enum_metal_value = Enum_Metal_value_Almglocken
	case "Enum_Metal_value_Bell":
		*enum_metal_value = Enum_Metal_value_Bell
	case "Enum_Metal_value_Bell_plate":
		*enum_metal_value = Enum_Metal_value_Bell_plate
	case "Enum_Metal_value_Bell_tree":
		*enum_metal_value = Enum_Metal_value_Bell_tree
	case "Enum_Metal_value_Brake_drum":
		*enum_metal_value = Enum_Metal_value_Brake_drum
	case "Enum_Metal_value_Cencerro":
		*enum_metal_value = Enum_Metal_value_Cencerro
	case "Enum_Metal_value_Chain_rattle":
		*enum_metal_value = Enum_Metal_value_Chain_rattle
	case "Enum_Metal_value_Chinese_cymbal":
		*enum_metal_value = Enum_Metal_value_Chinese_cymbal
	case "Enum_Metal_value_Cowbell":
		*enum_metal_value = Enum_Metal_value_Cowbell
	case "Enum_Metal_value_Crash_cymbals":
		*enum_metal_value = Enum_Metal_value_Crash_cymbals
	case "Enum_Metal_value_Crotale":
		*enum_metal_value = Enum_Metal_value_Crotale
	case "Enum_Metal_value_Cymbal_tongs":
		*enum_metal_value = Enum_Metal_value_Cymbal_tongs
	case "Enum_Metal_value_Domed_gong":
		*enum_metal_value = Enum_Metal_value_Domed_gong
	case "Enum_Metal_value_Finger_cymbals":
		*enum_metal_value = Enum_Metal_value_Finger_cymbals
	case "Enum_Metal_value_Flexatone":
		*enum_metal_value = Enum_Metal_value_Flexatone
	case "Enum_Metal_value_Gong":
		*enum_metal_value = Enum_Metal_value_Gong
	case "Enum_Metal_value_Hi_hat":
		*enum_metal_value = Enum_Metal_value_Hi_hat
	case "Enum_Metal_value_High_hat_cymbals":
		*enum_metal_value = Enum_Metal_value_High_hat_cymbals
	case "Enum_Metal_value_Handbell":
		*enum_metal_value = Enum_Metal_value_Handbell
	case "Enum_Metal_value_Jaw_harp":
		*enum_metal_value = Enum_Metal_value_Jaw_harp
	case "Enum_Metal_value_Jingle_bells":
		*enum_metal_value = Enum_Metal_value_Jingle_bells
	case "Enum_Metal_value_Musical_saw":
		*enum_metal_value = Enum_Metal_value_Musical_saw
	case "Enum_Metal_value_Shell_bells":
		*enum_metal_value = Enum_Metal_value_Shell_bells
	case "Enum_Metal_value_Sistrum":
		*enum_metal_value = Enum_Metal_value_Sistrum
	case "Enum_Metal_value_Sizzle_cymbal":
		*enum_metal_value = Enum_Metal_value_Sizzle_cymbal
	case "Enum_Metal_value_Sleigh_bells":
		*enum_metal_value = Enum_Metal_value_Sleigh_bells
	case "Enum_Metal_value_Suspended_cymbal":
		*enum_metal_value = Enum_Metal_value_Suspended_cymbal
	case "Enum_Metal_value_Tam_tam":
		*enum_metal_value = Enum_Metal_value_Tam_tam
	case "Enum_Metal_value_Tam_tam_with_beater":
		*enum_metal_value = Enum_Metal_value_Tam_tam_with_beater
	case "Enum_Metal_value_Triangle":
		*enum_metal_value = Enum_Metal_value_Triangle
	case "Enum_Metal_value_Vietnamese_hat":
		*enum_metal_value = Enum_Metal_value_Vietnamese_hat
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_metal_value *Enum_Metal_value) ToCodeString() (res string) {

	switch *enum_metal_value {
	// insertion code per enum code
	case Enum_Metal_value_Agogo:
		res = "Enum_Metal_value_Agogo"
	case Enum_Metal_value_Almglocken:
		res = "Enum_Metal_value_Almglocken"
	case Enum_Metal_value_Bell:
		res = "Enum_Metal_value_Bell"
	case Enum_Metal_value_Bell_plate:
		res = "Enum_Metal_value_Bell_plate"
	case Enum_Metal_value_Bell_tree:
		res = "Enum_Metal_value_Bell_tree"
	case Enum_Metal_value_Brake_drum:
		res = "Enum_Metal_value_Brake_drum"
	case Enum_Metal_value_Cencerro:
		res = "Enum_Metal_value_Cencerro"
	case Enum_Metal_value_Chain_rattle:
		res = "Enum_Metal_value_Chain_rattle"
	case Enum_Metal_value_Chinese_cymbal:
		res = "Enum_Metal_value_Chinese_cymbal"
	case Enum_Metal_value_Cowbell:
		res = "Enum_Metal_value_Cowbell"
	case Enum_Metal_value_Crash_cymbals:
		res = "Enum_Metal_value_Crash_cymbals"
	case Enum_Metal_value_Crotale:
		res = "Enum_Metal_value_Crotale"
	case Enum_Metal_value_Cymbal_tongs:
		res = "Enum_Metal_value_Cymbal_tongs"
	case Enum_Metal_value_Domed_gong:
		res = "Enum_Metal_value_Domed_gong"
	case Enum_Metal_value_Finger_cymbals:
		res = "Enum_Metal_value_Finger_cymbals"
	case Enum_Metal_value_Flexatone:
		res = "Enum_Metal_value_Flexatone"
	case Enum_Metal_value_Gong:
		res = "Enum_Metal_value_Gong"
	case Enum_Metal_value_Hi_hat:
		res = "Enum_Metal_value_Hi_hat"
	case Enum_Metal_value_High_hat_cymbals:
		res = "Enum_Metal_value_High_hat_cymbals"
	case Enum_Metal_value_Handbell:
		res = "Enum_Metal_value_Handbell"
	case Enum_Metal_value_Jaw_harp:
		res = "Enum_Metal_value_Jaw_harp"
	case Enum_Metal_value_Jingle_bells:
		res = "Enum_Metal_value_Jingle_bells"
	case Enum_Metal_value_Musical_saw:
		res = "Enum_Metal_value_Musical_saw"
	case Enum_Metal_value_Shell_bells:
		res = "Enum_Metal_value_Shell_bells"
	case Enum_Metal_value_Sistrum:
		res = "Enum_Metal_value_Sistrum"
	case Enum_Metal_value_Sizzle_cymbal:
		res = "Enum_Metal_value_Sizzle_cymbal"
	case Enum_Metal_value_Sleigh_bells:
		res = "Enum_Metal_value_Sleigh_bells"
	case Enum_Metal_value_Suspended_cymbal:
		res = "Enum_Metal_value_Suspended_cymbal"
	case Enum_Metal_value_Tam_tam:
		res = "Enum_Metal_value_Tam_tam"
	case Enum_Metal_value_Tam_tam_with_beater:
		res = "Enum_Metal_value_Tam_tam_with_beater"
	case Enum_Metal_value_Triangle:
		res = "Enum_Metal_value_Triangle"
	case Enum_Metal_value_Vietnamese_hat:
		res = "Enum_Metal_value_Vietnamese_hat"
	}
	return
}

func (enum_metal_value Enum_Metal_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Metal_value_Agogo")
	res = append(res, "Enum_Metal_value_Almglocken")
	res = append(res, "Enum_Metal_value_Bell")
	res = append(res, "Enum_Metal_value_Bell_plate")
	res = append(res, "Enum_Metal_value_Bell_tree")
	res = append(res, "Enum_Metal_value_Brake_drum")
	res = append(res, "Enum_Metal_value_Cencerro")
	res = append(res, "Enum_Metal_value_Chain_rattle")
	res = append(res, "Enum_Metal_value_Chinese_cymbal")
	res = append(res, "Enum_Metal_value_Cowbell")
	res = append(res, "Enum_Metal_value_Crash_cymbals")
	res = append(res, "Enum_Metal_value_Crotale")
	res = append(res, "Enum_Metal_value_Cymbal_tongs")
	res = append(res, "Enum_Metal_value_Domed_gong")
	res = append(res, "Enum_Metal_value_Finger_cymbals")
	res = append(res, "Enum_Metal_value_Flexatone")
	res = append(res, "Enum_Metal_value_Gong")
	res = append(res, "Enum_Metal_value_Hi_hat")
	res = append(res, "Enum_Metal_value_High_hat_cymbals")
	res = append(res, "Enum_Metal_value_Handbell")
	res = append(res, "Enum_Metal_value_Jaw_harp")
	res = append(res, "Enum_Metal_value_Jingle_bells")
	res = append(res, "Enum_Metal_value_Musical_saw")
	res = append(res, "Enum_Metal_value_Shell_bells")
	res = append(res, "Enum_Metal_value_Sistrum")
	res = append(res, "Enum_Metal_value_Sizzle_cymbal")
	res = append(res, "Enum_Metal_value_Sleigh_bells")
	res = append(res, "Enum_Metal_value_Suspended_cymbal")
	res = append(res, "Enum_Metal_value_Tam_tam")
	res = append(res, "Enum_Metal_value_Tam_tam_with_beater")
	res = append(res, "Enum_Metal_value_Triangle")
	res = append(res, "Enum_Metal_value_Vietnamese_hat")

	return
}

func (enum_metal_value Enum_Metal_value) CodeValues() (res []string) {

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

// Utility function for Enum_Mode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_mode Enum_Mode) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_mode {
	// insertion code per enum code
	}
	return
}

func (enum_mode *Enum_Mode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_mode *Enum_Mode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_mode *Enum_Mode) ToCodeString() (res string) {

	switch *enum_mode {
	// insertion code per enum code
	}
	return
}

func (enum_mode Enum_Mode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_mode Enum_Mode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Mute
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_mute Enum_Mute) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_mute {
	// insertion code per enum code
	case Enum_Mute_On:
		res = "on"
	case Enum_Mute_Off:
		res = "off"
	case Enum_Mute_Straight:
		res = "straight"
	case Enum_Mute_Cup:
		res = "cup"
	case Enum_Mute_Harmon_no_stem:
		res = "harmon-no-stem"
	case Enum_Mute_Harmon_stem:
		res = "harmon-stem"
	case Enum_Mute_Bucket:
		res = "bucket"
	case Enum_Mute_Plunger:
		res = "plunger"
	case Enum_Mute_Hat:
		res = "hat"
	case Enum_Mute_Solotone:
		res = "solotone"
	case Enum_Mute_Practice:
		res = "practice"
	case Enum_Mute_Stop_mute:
		res = "stop-mute"
	case Enum_Mute_Stop_hand:
		res = "stop-hand"
	case Enum_Mute_Echo:
		res = "echo"
	case Enum_Mute_Palm:
		res = "palm"
	}
	return
}

func (enum_mute *Enum_Mute) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*enum_mute = Enum_Mute_On
		return
	case "off":
		*enum_mute = Enum_Mute_Off
		return
	case "straight":
		*enum_mute = Enum_Mute_Straight
		return
	case "cup":
		*enum_mute = Enum_Mute_Cup
		return
	case "harmon-no-stem":
		*enum_mute = Enum_Mute_Harmon_no_stem
		return
	case "harmon-stem":
		*enum_mute = Enum_Mute_Harmon_stem
		return
	case "bucket":
		*enum_mute = Enum_Mute_Bucket
		return
	case "plunger":
		*enum_mute = Enum_Mute_Plunger
		return
	case "hat":
		*enum_mute = Enum_Mute_Hat
		return
	case "solotone":
		*enum_mute = Enum_Mute_Solotone
		return
	case "practice":
		*enum_mute = Enum_Mute_Practice
		return
	case "stop-mute":
		*enum_mute = Enum_Mute_Stop_mute
		return
	case "stop-hand":
		*enum_mute = Enum_Mute_Stop_hand
		return
	case "echo":
		*enum_mute = Enum_Mute_Echo
		return
	case "palm":
		*enum_mute = Enum_Mute_Palm
		return
	default:
		return errUnkownEnum
	}
}

func (enum_mute *Enum_Mute) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Mute_On":
		*enum_mute = Enum_Mute_On
	case "Enum_Mute_Off":
		*enum_mute = Enum_Mute_Off
	case "Enum_Mute_Straight":
		*enum_mute = Enum_Mute_Straight
	case "Enum_Mute_Cup":
		*enum_mute = Enum_Mute_Cup
	case "Enum_Mute_Harmon_no_stem":
		*enum_mute = Enum_Mute_Harmon_no_stem
	case "Enum_Mute_Harmon_stem":
		*enum_mute = Enum_Mute_Harmon_stem
	case "Enum_Mute_Bucket":
		*enum_mute = Enum_Mute_Bucket
	case "Enum_Mute_Plunger":
		*enum_mute = Enum_Mute_Plunger
	case "Enum_Mute_Hat":
		*enum_mute = Enum_Mute_Hat
	case "Enum_Mute_Solotone":
		*enum_mute = Enum_Mute_Solotone
	case "Enum_Mute_Practice":
		*enum_mute = Enum_Mute_Practice
	case "Enum_Mute_Stop_mute":
		*enum_mute = Enum_Mute_Stop_mute
	case "Enum_Mute_Stop_hand":
		*enum_mute = Enum_Mute_Stop_hand
	case "Enum_Mute_Echo":
		*enum_mute = Enum_Mute_Echo
	case "Enum_Mute_Palm":
		*enum_mute = Enum_Mute_Palm
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_mute *Enum_Mute) ToCodeString() (res string) {

	switch *enum_mute {
	// insertion code per enum code
	case Enum_Mute_On:
		res = "Enum_Mute_On"
	case Enum_Mute_Off:
		res = "Enum_Mute_Off"
	case Enum_Mute_Straight:
		res = "Enum_Mute_Straight"
	case Enum_Mute_Cup:
		res = "Enum_Mute_Cup"
	case Enum_Mute_Harmon_no_stem:
		res = "Enum_Mute_Harmon_no_stem"
	case Enum_Mute_Harmon_stem:
		res = "Enum_Mute_Harmon_stem"
	case Enum_Mute_Bucket:
		res = "Enum_Mute_Bucket"
	case Enum_Mute_Plunger:
		res = "Enum_Mute_Plunger"
	case Enum_Mute_Hat:
		res = "Enum_Mute_Hat"
	case Enum_Mute_Solotone:
		res = "Enum_Mute_Solotone"
	case Enum_Mute_Practice:
		res = "Enum_Mute_Practice"
	case Enum_Mute_Stop_mute:
		res = "Enum_Mute_Stop_mute"
	case Enum_Mute_Stop_hand:
		res = "Enum_Mute_Stop_hand"
	case Enum_Mute_Echo:
		res = "Enum_Mute_Echo"
	case Enum_Mute_Palm:
		res = "Enum_Mute_Palm"
	}
	return
}

func (enum_mute Enum_Mute) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Mute_On")
	res = append(res, "Enum_Mute_Off")
	res = append(res, "Enum_Mute_Straight")
	res = append(res, "Enum_Mute_Cup")
	res = append(res, "Enum_Mute_Harmon_no_stem")
	res = append(res, "Enum_Mute_Harmon_stem")
	res = append(res, "Enum_Mute_Bucket")
	res = append(res, "Enum_Mute_Plunger")
	res = append(res, "Enum_Mute_Hat")
	res = append(res, "Enum_Mute_Solotone")
	res = append(res, "Enum_Mute_Practice")
	res = append(res, "Enum_Mute_Stop_mute")
	res = append(res, "Enum_Mute_Stop_hand")
	res = append(res, "Enum_Mute_Echo")
	res = append(res, "Enum_Mute_Palm")

	return
}

func (enum_mute Enum_Mute) CodeValues() (res []string) {

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

// Utility function for Enum_Note_size_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_note_size_type Enum_Note_size_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_note_size_type {
	// insertion code per enum code
	case Enum_Note_size_type_Cue:
		res = "cue"
	case Enum_Note_size_type_Grace:
		res = "grace"
	case Enum_Note_size_type_Grace_cue:
		res = "grace-cue"
	case Enum_Note_size_type_Large:
		res = "large"
	}
	return
}

func (enum_note_size_type *Enum_Note_size_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "cue":
		*enum_note_size_type = Enum_Note_size_type_Cue
		return
	case "grace":
		*enum_note_size_type = Enum_Note_size_type_Grace
		return
	case "grace-cue":
		*enum_note_size_type = Enum_Note_size_type_Grace_cue
		return
	case "large":
		*enum_note_size_type = Enum_Note_size_type_Large
		return
	default:
		return errUnkownEnum
	}
}

func (enum_note_size_type *Enum_Note_size_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Note_size_type_Cue":
		*enum_note_size_type = Enum_Note_size_type_Cue
	case "Enum_Note_size_type_Grace":
		*enum_note_size_type = Enum_Note_size_type_Grace
	case "Enum_Note_size_type_Grace_cue":
		*enum_note_size_type = Enum_Note_size_type_Grace_cue
	case "Enum_Note_size_type_Large":
		*enum_note_size_type = Enum_Note_size_type_Large
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_note_size_type *Enum_Note_size_type) ToCodeString() (res string) {

	switch *enum_note_size_type {
	// insertion code per enum code
	case Enum_Note_size_type_Cue:
		res = "Enum_Note_size_type_Cue"
	case Enum_Note_size_type_Grace:
		res = "Enum_Note_size_type_Grace"
	case Enum_Note_size_type_Grace_cue:
		res = "Enum_Note_size_type_Grace_cue"
	case Enum_Note_size_type_Large:
		res = "Enum_Note_size_type_Large"
	}
	return
}

func (enum_note_size_type Enum_Note_size_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Note_size_type_Cue")
	res = append(res, "Enum_Note_size_type_Grace")
	res = append(res, "Enum_Note_size_type_Grace_cue")
	res = append(res, "Enum_Note_size_type_Large")

	return
}

func (enum_note_size_type Enum_Note_size_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "cue")
	res = append(res, "grace")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for Enum_Note_type_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_note_type_value Enum_Note_type_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_note_type_value {
	// insertion code per enum code
	case Enum_Note_type_value_1024th:
		res = "1024th"
	case Enum_Note_type_value_512th:
		res = "512th"
	case Enum_Note_type_value_256th:
		res = "256th"
	case Enum_Note_type_value_128th:
		res = "128th"
	case Enum_Note_type_value_64th:
		res = "64th"
	case Enum_Note_type_value_32nd:
		res = "32nd"
	case Enum_Note_type_value_16th:
		res = "16th"
	case Enum_Note_type_value_Eighth:
		res = "eighth"
	case Enum_Note_type_value_Quarter:
		res = "quarter"
	case Enum_Note_type_value_Half:
		res = "half"
	case Enum_Note_type_value_Whole:
		res = "whole"
	case Enum_Note_type_value_Breve:
		res = "breve"
	case Enum_Note_type_value_Long:
		res = "long"
	case Enum_Note_type_value_Maxima:
		res = "maxima"
	}
	return
}

func (enum_note_type_value *Enum_Note_type_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1024th":
		*enum_note_type_value = Enum_Note_type_value_1024th
		return
	case "512th":
		*enum_note_type_value = Enum_Note_type_value_512th
		return
	case "256th":
		*enum_note_type_value = Enum_Note_type_value_256th
		return
	case "128th":
		*enum_note_type_value = Enum_Note_type_value_128th
		return
	case "64th":
		*enum_note_type_value = Enum_Note_type_value_64th
		return
	case "32nd":
		*enum_note_type_value = Enum_Note_type_value_32nd
		return
	case "16th":
		*enum_note_type_value = Enum_Note_type_value_16th
		return
	case "eighth":
		*enum_note_type_value = Enum_Note_type_value_Eighth
		return
	case "quarter":
		*enum_note_type_value = Enum_Note_type_value_Quarter
		return
	case "half":
		*enum_note_type_value = Enum_Note_type_value_Half
		return
	case "whole":
		*enum_note_type_value = Enum_Note_type_value_Whole
		return
	case "breve":
		*enum_note_type_value = Enum_Note_type_value_Breve
		return
	case "long":
		*enum_note_type_value = Enum_Note_type_value_Long
		return
	case "maxima":
		*enum_note_type_value = Enum_Note_type_value_Maxima
		return
	default:
		return errUnkownEnum
	}
}

func (enum_note_type_value *Enum_Note_type_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Note_type_value_1024th":
		*enum_note_type_value = Enum_Note_type_value_1024th
	case "Enum_Note_type_value_512th":
		*enum_note_type_value = Enum_Note_type_value_512th
	case "Enum_Note_type_value_256th":
		*enum_note_type_value = Enum_Note_type_value_256th
	case "Enum_Note_type_value_128th":
		*enum_note_type_value = Enum_Note_type_value_128th
	case "Enum_Note_type_value_64th":
		*enum_note_type_value = Enum_Note_type_value_64th
	case "Enum_Note_type_value_32nd":
		*enum_note_type_value = Enum_Note_type_value_32nd
	case "Enum_Note_type_value_16th":
		*enum_note_type_value = Enum_Note_type_value_16th
	case "Enum_Note_type_value_Eighth":
		*enum_note_type_value = Enum_Note_type_value_Eighth
	case "Enum_Note_type_value_Quarter":
		*enum_note_type_value = Enum_Note_type_value_Quarter
	case "Enum_Note_type_value_Half":
		*enum_note_type_value = Enum_Note_type_value_Half
	case "Enum_Note_type_value_Whole":
		*enum_note_type_value = Enum_Note_type_value_Whole
	case "Enum_Note_type_value_Breve":
		*enum_note_type_value = Enum_Note_type_value_Breve
	case "Enum_Note_type_value_Long":
		*enum_note_type_value = Enum_Note_type_value_Long
	case "Enum_Note_type_value_Maxima":
		*enum_note_type_value = Enum_Note_type_value_Maxima
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_note_type_value *Enum_Note_type_value) ToCodeString() (res string) {

	switch *enum_note_type_value {
	// insertion code per enum code
	case Enum_Note_type_value_1024th:
		res = "Enum_Note_type_value_1024th"
	case Enum_Note_type_value_512th:
		res = "Enum_Note_type_value_512th"
	case Enum_Note_type_value_256th:
		res = "Enum_Note_type_value_256th"
	case Enum_Note_type_value_128th:
		res = "Enum_Note_type_value_128th"
	case Enum_Note_type_value_64th:
		res = "Enum_Note_type_value_64th"
	case Enum_Note_type_value_32nd:
		res = "Enum_Note_type_value_32nd"
	case Enum_Note_type_value_16th:
		res = "Enum_Note_type_value_16th"
	case Enum_Note_type_value_Eighth:
		res = "Enum_Note_type_value_Eighth"
	case Enum_Note_type_value_Quarter:
		res = "Enum_Note_type_value_Quarter"
	case Enum_Note_type_value_Half:
		res = "Enum_Note_type_value_Half"
	case Enum_Note_type_value_Whole:
		res = "Enum_Note_type_value_Whole"
	case Enum_Note_type_value_Breve:
		res = "Enum_Note_type_value_Breve"
	case Enum_Note_type_value_Long:
		res = "Enum_Note_type_value_Long"
	case Enum_Note_type_value_Maxima:
		res = "Enum_Note_type_value_Maxima"
	}
	return
}

func (enum_note_type_value Enum_Note_type_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Note_type_value_1024th")
	res = append(res, "Enum_Note_type_value_512th")
	res = append(res, "Enum_Note_type_value_256th")
	res = append(res, "Enum_Note_type_value_128th")
	res = append(res, "Enum_Note_type_value_64th")
	res = append(res, "Enum_Note_type_value_32nd")
	res = append(res, "Enum_Note_type_value_16th")
	res = append(res, "Enum_Note_type_value_Eighth")
	res = append(res, "Enum_Note_type_value_Quarter")
	res = append(res, "Enum_Note_type_value_Half")
	res = append(res, "Enum_Note_type_value_Whole")
	res = append(res, "Enum_Note_type_value_Breve")
	res = append(res, "Enum_Note_type_value_Long")
	res = append(res, "Enum_Note_type_value_Maxima")

	return
}

func (enum_note_type_value Enum_Note_type_value) CodeValues() (res []string) {

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

// Utility function for Enum_Notehead_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_notehead_value Enum_Notehead_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_notehead_value {
	// insertion code per enum code
	case Enum_Notehead_value_Slash:
		res = "slash"
	case Enum_Notehead_value_Triangle:
		res = "triangle"
	case Enum_Notehead_value_Diamond:
		res = "diamond"
	case Enum_Notehead_value_Square:
		res = "square"
	case Enum_Notehead_value_Cross:
		res = "cross"
	case Enum_Notehead_value_X:
		res = "x"
	case Enum_Notehead_value_Circle_x:
		res = "circle-x"
	case Enum_Notehead_value_Inverted_triangle:
		res = "inverted triangle"
	case Enum_Notehead_value_Arrow_down:
		res = "arrow down"
	case Enum_Notehead_value_Arrow_up:
		res = "arrow up"
	case Enum_Notehead_value_Circled:
		res = "circled"
	case Enum_Notehead_value_Slashed:
		res = "slashed"
	case Enum_Notehead_value_Back_slashed:
		res = "back slashed"
	case Enum_Notehead_value_Normal:
		res = "normal"
	case Enum_Notehead_value_Cluster:
		res = "cluster"
	case Enum_Notehead_value_Circle_dot:
		res = "circle dot"
	case Enum_Notehead_value_Left_triangle:
		res = "left triangle"
	case Enum_Notehead_value_Rectangle:
		res = "rectangle"
	case Enum_Notehead_value_None:
		res = "none"
	case Enum_Notehead_value_Do:
		res = "do"
	case Enum_Notehead_value_Re:
		res = "re"
	case Enum_Notehead_value_Mi:
		res = "mi"
	case Enum_Notehead_value_Fa:
		res = "fa"
	case Enum_Notehead_value_Fa_up:
		res = "fa up"
	case Enum_Notehead_value_So:
		res = "so"
	case Enum_Notehead_value_La:
		res = "la"
	case Enum_Notehead_value_Ti:
		res = "ti"
	case Enum_Notehead_value_Other:
		res = "other"
	}
	return
}

func (enum_notehead_value *Enum_Notehead_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "slash":
		*enum_notehead_value = Enum_Notehead_value_Slash
		return
	case "triangle":
		*enum_notehead_value = Enum_Notehead_value_Triangle
		return
	case "diamond":
		*enum_notehead_value = Enum_Notehead_value_Diamond
		return
	case "square":
		*enum_notehead_value = Enum_Notehead_value_Square
		return
	case "cross":
		*enum_notehead_value = Enum_Notehead_value_Cross
		return
	case "x":
		*enum_notehead_value = Enum_Notehead_value_X
		return
	case "circle-x":
		*enum_notehead_value = Enum_Notehead_value_Circle_x
		return
	case "inverted triangle":
		*enum_notehead_value = Enum_Notehead_value_Inverted_triangle
		return
	case "arrow down":
		*enum_notehead_value = Enum_Notehead_value_Arrow_down
		return
	case "arrow up":
		*enum_notehead_value = Enum_Notehead_value_Arrow_up
		return
	case "circled":
		*enum_notehead_value = Enum_Notehead_value_Circled
		return
	case "slashed":
		*enum_notehead_value = Enum_Notehead_value_Slashed
		return
	case "back slashed":
		*enum_notehead_value = Enum_Notehead_value_Back_slashed
		return
	case "normal":
		*enum_notehead_value = Enum_Notehead_value_Normal
		return
	case "cluster":
		*enum_notehead_value = Enum_Notehead_value_Cluster
		return
	case "circle dot":
		*enum_notehead_value = Enum_Notehead_value_Circle_dot
		return
	case "left triangle":
		*enum_notehead_value = Enum_Notehead_value_Left_triangle
		return
	case "rectangle":
		*enum_notehead_value = Enum_Notehead_value_Rectangle
		return
	case "none":
		*enum_notehead_value = Enum_Notehead_value_None
		return
	case "do":
		*enum_notehead_value = Enum_Notehead_value_Do
		return
	case "re":
		*enum_notehead_value = Enum_Notehead_value_Re
		return
	case "mi":
		*enum_notehead_value = Enum_Notehead_value_Mi
		return
	case "fa":
		*enum_notehead_value = Enum_Notehead_value_Fa
		return
	case "fa up":
		*enum_notehead_value = Enum_Notehead_value_Fa_up
		return
	case "so":
		*enum_notehead_value = Enum_Notehead_value_So
		return
	case "la":
		*enum_notehead_value = Enum_Notehead_value_La
		return
	case "ti":
		*enum_notehead_value = Enum_Notehead_value_Ti
		return
	case "other":
		*enum_notehead_value = Enum_Notehead_value_Other
		return
	default:
		return errUnkownEnum
	}
}

func (enum_notehead_value *Enum_Notehead_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Notehead_value_Slash":
		*enum_notehead_value = Enum_Notehead_value_Slash
	case "Enum_Notehead_value_Triangle":
		*enum_notehead_value = Enum_Notehead_value_Triangle
	case "Enum_Notehead_value_Diamond":
		*enum_notehead_value = Enum_Notehead_value_Diamond
	case "Enum_Notehead_value_Square":
		*enum_notehead_value = Enum_Notehead_value_Square
	case "Enum_Notehead_value_Cross":
		*enum_notehead_value = Enum_Notehead_value_Cross
	case "Enum_Notehead_value_X":
		*enum_notehead_value = Enum_Notehead_value_X
	case "Enum_Notehead_value_Circle_x":
		*enum_notehead_value = Enum_Notehead_value_Circle_x
	case "Enum_Notehead_value_Inverted_triangle":
		*enum_notehead_value = Enum_Notehead_value_Inverted_triangle
	case "Enum_Notehead_value_Arrow_down":
		*enum_notehead_value = Enum_Notehead_value_Arrow_down
	case "Enum_Notehead_value_Arrow_up":
		*enum_notehead_value = Enum_Notehead_value_Arrow_up
	case "Enum_Notehead_value_Circled":
		*enum_notehead_value = Enum_Notehead_value_Circled
	case "Enum_Notehead_value_Slashed":
		*enum_notehead_value = Enum_Notehead_value_Slashed
	case "Enum_Notehead_value_Back_slashed":
		*enum_notehead_value = Enum_Notehead_value_Back_slashed
	case "Enum_Notehead_value_Normal":
		*enum_notehead_value = Enum_Notehead_value_Normal
	case "Enum_Notehead_value_Cluster":
		*enum_notehead_value = Enum_Notehead_value_Cluster
	case "Enum_Notehead_value_Circle_dot":
		*enum_notehead_value = Enum_Notehead_value_Circle_dot
	case "Enum_Notehead_value_Left_triangle":
		*enum_notehead_value = Enum_Notehead_value_Left_triangle
	case "Enum_Notehead_value_Rectangle":
		*enum_notehead_value = Enum_Notehead_value_Rectangle
	case "Enum_Notehead_value_None":
		*enum_notehead_value = Enum_Notehead_value_None
	case "Enum_Notehead_value_Do":
		*enum_notehead_value = Enum_Notehead_value_Do
	case "Enum_Notehead_value_Re":
		*enum_notehead_value = Enum_Notehead_value_Re
	case "Enum_Notehead_value_Mi":
		*enum_notehead_value = Enum_Notehead_value_Mi
	case "Enum_Notehead_value_Fa":
		*enum_notehead_value = Enum_Notehead_value_Fa
	case "Enum_Notehead_value_Fa_up":
		*enum_notehead_value = Enum_Notehead_value_Fa_up
	case "Enum_Notehead_value_So":
		*enum_notehead_value = Enum_Notehead_value_So
	case "Enum_Notehead_value_La":
		*enum_notehead_value = Enum_Notehead_value_La
	case "Enum_Notehead_value_Ti":
		*enum_notehead_value = Enum_Notehead_value_Ti
	case "Enum_Notehead_value_Other":
		*enum_notehead_value = Enum_Notehead_value_Other
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_notehead_value *Enum_Notehead_value) ToCodeString() (res string) {

	switch *enum_notehead_value {
	// insertion code per enum code
	case Enum_Notehead_value_Slash:
		res = "Enum_Notehead_value_Slash"
	case Enum_Notehead_value_Triangle:
		res = "Enum_Notehead_value_Triangle"
	case Enum_Notehead_value_Diamond:
		res = "Enum_Notehead_value_Diamond"
	case Enum_Notehead_value_Square:
		res = "Enum_Notehead_value_Square"
	case Enum_Notehead_value_Cross:
		res = "Enum_Notehead_value_Cross"
	case Enum_Notehead_value_X:
		res = "Enum_Notehead_value_X"
	case Enum_Notehead_value_Circle_x:
		res = "Enum_Notehead_value_Circle_x"
	case Enum_Notehead_value_Inverted_triangle:
		res = "Enum_Notehead_value_Inverted_triangle"
	case Enum_Notehead_value_Arrow_down:
		res = "Enum_Notehead_value_Arrow_down"
	case Enum_Notehead_value_Arrow_up:
		res = "Enum_Notehead_value_Arrow_up"
	case Enum_Notehead_value_Circled:
		res = "Enum_Notehead_value_Circled"
	case Enum_Notehead_value_Slashed:
		res = "Enum_Notehead_value_Slashed"
	case Enum_Notehead_value_Back_slashed:
		res = "Enum_Notehead_value_Back_slashed"
	case Enum_Notehead_value_Normal:
		res = "Enum_Notehead_value_Normal"
	case Enum_Notehead_value_Cluster:
		res = "Enum_Notehead_value_Cluster"
	case Enum_Notehead_value_Circle_dot:
		res = "Enum_Notehead_value_Circle_dot"
	case Enum_Notehead_value_Left_triangle:
		res = "Enum_Notehead_value_Left_triangle"
	case Enum_Notehead_value_Rectangle:
		res = "Enum_Notehead_value_Rectangle"
	case Enum_Notehead_value_None:
		res = "Enum_Notehead_value_None"
	case Enum_Notehead_value_Do:
		res = "Enum_Notehead_value_Do"
	case Enum_Notehead_value_Re:
		res = "Enum_Notehead_value_Re"
	case Enum_Notehead_value_Mi:
		res = "Enum_Notehead_value_Mi"
	case Enum_Notehead_value_Fa:
		res = "Enum_Notehead_value_Fa"
	case Enum_Notehead_value_Fa_up:
		res = "Enum_Notehead_value_Fa_up"
	case Enum_Notehead_value_So:
		res = "Enum_Notehead_value_So"
	case Enum_Notehead_value_La:
		res = "Enum_Notehead_value_La"
	case Enum_Notehead_value_Ti:
		res = "Enum_Notehead_value_Ti"
	case Enum_Notehead_value_Other:
		res = "Enum_Notehead_value_Other"
	}
	return
}

func (enum_notehead_value Enum_Notehead_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Notehead_value_Slash")
	res = append(res, "Enum_Notehead_value_Triangle")
	res = append(res, "Enum_Notehead_value_Diamond")
	res = append(res, "Enum_Notehead_value_Square")
	res = append(res, "Enum_Notehead_value_Cross")
	res = append(res, "Enum_Notehead_value_X")
	res = append(res, "Enum_Notehead_value_Circle_x")
	res = append(res, "Enum_Notehead_value_Inverted_triangle")
	res = append(res, "Enum_Notehead_value_Arrow_down")
	res = append(res, "Enum_Notehead_value_Arrow_up")
	res = append(res, "Enum_Notehead_value_Circled")
	res = append(res, "Enum_Notehead_value_Slashed")
	res = append(res, "Enum_Notehead_value_Back_slashed")
	res = append(res, "Enum_Notehead_value_Normal")
	res = append(res, "Enum_Notehead_value_Cluster")
	res = append(res, "Enum_Notehead_value_Circle_dot")
	res = append(res, "Enum_Notehead_value_Left_triangle")
	res = append(res, "Enum_Notehead_value_Rectangle")
	res = append(res, "Enum_Notehead_value_None")
	res = append(res, "Enum_Notehead_value_Do")
	res = append(res, "Enum_Notehead_value_Re")
	res = append(res, "Enum_Notehead_value_Mi")
	res = append(res, "Enum_Notehead_value_Fa")
	res = append(res, "Enum_Notehead_value_Fa_up")
	res = append(res, "Enum_Notehead_value_So")
	res = append(res, "Enum_Notehead_value_La")
	res = append(res, "Enum_Notehead_value_Ti")
	res = append(res, "Enum_Notehead_value_Other")

	return
}

func (enum_notehead_value Enum_Notehead_value) CodeValues() (res []string) {

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

// Utility function for Enum_Numeral_mode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_numeral_mode Enum_Numeral_mode) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_numeral_mode {
	// insertion code per enum code
	case Enum_Numeral_mode_Major:
		res = "major"
	case Enum_Numeral_mode_Minor:
		res = "minor"
	case Enum_Numeral_mode_Natural_minor:
		res = "natural minor"
	case Enum_Numeral_mode_Melodic_minor:
		res = "melodic minor"
	case Enum_Numeral_mode_Harmonic_minor:
		res = "harmonic minor"
	}
	return
}

func (enum_numeral_mode *Enum_Numeral_mode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*enum_numeral_mode = Enum_Numeral_mode_Major
		return
	case "minor":
		*enum_numeral_mode = Enum_Numeral_mode_Minor
		return
	case "natural minor":
		*enum_numeral_mode = Enum_Numeral_mode_Natural_minor
		return
	case "melodic minor":
		*enum_numeral_mode = Enum_Numeral_mode_Melodic_minor
		return
	case "harmonic minor":
		*enum_numeral_mode = Enum_Numeral_mode_Harmonic_minor
		return
	default:
		return errUnkownEnum
	}
}

func (enum_numeral_mode *Enum_Numeral_mode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Numeral_mode_Major":
		*enum_numeral_mode = Enum_Numeral_mode_Major
	case "Enum_Numeral_mode_Minor":
		*enum_numeral_mode = Enum_Numeral_mode_Minor
	case "Enum_Numeral_mode_Natural_minor":
		*enum_numeral_mode = Enum_Numeral_mode_Natural_minor
	case "Enum_Numeral_mode_Melodic_minor":
		*enum_numeral_mode = Enum_Numeral_mode_Melodic_minor
	case "Enum_Numeral_mode_Harmonic_minor":
		*enum_numeral_mode = Enum_Numeral_mode_Harmonic_minor
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_numeral_mode *Enum_Numeral_mode) ToCodeString() (res string) {

	switch *enum_numeral_mode {
	// insertion code per enum code
	case Enum_Numeral_mode_Major:
		res = "Enum_Numeral_mode_Major"
	case Enum_Numeral_mode_Minor:
		res = "Enum_Numeral_mode_Minor"
	case Enum_Numeral_mode_Natural_minor:
		res = "Enum_Numeral_mode_Natural_minor"
	case Enum_Numeral_mode_Melodic_minor:
		res = "Enum_Numeral_mode_Melodic_minor"
	case Enum_Numeral_mode_Harmonic_minor:
		res = "Enum_Numeral_mode_Harmonic_minor"
	}
	return
}

func (enum_numeral_mode Enum_Numeral_mode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Numeral_mode_Major")
	res = append(res, "Enum_Numeral_mode_Minor")
	res = append(res, "Enum_Numeral_mode_Natural_minor")
	res = append(res, "Enum_Numeral_mode_Melodic_minor")
	res = append(res, "Enum_Numeral_mode_Harmonic_minor")

	return
}

func (enum_numeral_mode Enum_Numeral_mode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "natural minor")
	res = append(res, "melodic minor")
	res = append(res, "harmonic minor")

	return
}

// Utility function for Enum_On_off
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_on_off Enum_On_off) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_on_off {
	// insertion code per enum code
	case Enum_On_off_On:
		res = "on"
	case Enum_On_off_Off:
		res = "off"
	}
	return
}

func (enum_on_off *Enum_On_off) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*enum_on_off = Enum_On_off_On
		return
	case "off":
		*enum_on_off = Enum_On_off_Off
		return
	default:
		return errUnkownEnum
	}
}

func (enum_on_off *Enum_On_off) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_On_off_On":
		*enum_on_off = Enum_On_off_On
	case "Enum_On_off_Off":
		*enum_on_off = Enum_On_off_Off
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_on_off *Enum_On_off) ToCodeString() (res string) {

	switch *enum_on_off {
	// insertion code per enum code
	case Enum_On_off_On:
		res = "Enum_On_off_On"
	case Enum_On_off_Off:
		res = "Enum_On_off_Off"
	}
	return
}

func (enum_on_off Enum_On_off) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_On_off_On")
	res = append(res, "Enum_On_off_Off")

	return
}

func (enum_on_off Enum_On_off) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "on")
	res = append(res, "off")

	return
}

// Utility function for Enum_Over_under
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_over_under Enum_Over_under) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_over_under {
	// insertion code per enum code
	case Enum_Over_under_Over:
		res = "over"
	case Enum_Over_under_Under:
		res = "under"
	}
	return
}

func (enum_over_under *Enum_Over_under) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "over":
		*enum_over_under = Enum_Over_under_Over
		return
	case "under":
		*enum_over_under = Enum_Over_under_Under
		return
	default:
		return errUnkownEnum
	}
}

func (enum_over_under *Enum_Over_under) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Over_under_Over":
		*enum_over_under = Enum_Over_under_Over
	case "Enum_Over_under_Under":
		*enum_over_under = Enum_Over_under_Under
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_over_under *Enum_Over_under) ToCodeString() (res string) {

	switch *enum_over_under {
	// insertion code per enum code
	case Enum_Over_under_Over:
		res = "Enum_Over_under_Over"
	case Enum_Over_under_Under:
		res = "Enum_Over_under_Under"
	}
	return
}

func (enum_over_under Enum_Over_under) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Over_under_Over")
	res = append(res, "Enum_Over_under_Under")

	return
}

func (enum_over_under Enum_Over_under) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "over")
	res = append(res, "under")

	return
}

// Utility function for Enum_Pedal_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_pedal_type Enum_Pedal_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_pedal_type {
	// insertion code per enum code
	case Enum_Pedal_type_Start:
		res = "start"
	case Enum_Pedal_type_Stop:
		res = "stop"
	case Enum_Pedal_type_Sostenuto:
		res = "sostenuto"
	case Enum_Pedal_type_Change:
		res = "change"
	case Enum_Pedal_type_Continue:
		res = "continue"
	case Enum_Pedal_type_Discontinue:
		res = "discontinue"
	case Enum_Pedal_type_Resume:
		res = "resume"
	}
	return
}

func (enum_pedal_type *Enum_Pedal_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_pedal_type = Enum_Pedal_type_Start
		return
	case "stop":
		*enum_pedal_type = Enum_Pedal_type_Stop
		return
	case "sostenuto":
		*enum_pedal_type = Enum_Pedal_type_Sostenuto
		return
	case "change":
		*enum_pedal_type = Enum_Pedal_type_Change
		return
	case "continue":
		*enum_pedal_type = Enum_Pedal_type_Continue
		return
	case "discontinue":
		*enum_pedal_type = Enum_Pedal_type_Discontinue
		return
	case "resume":
		*enum_pedal_type = Enum_Pedal_type_Resume
		return
	default:
		return errUnkownEnum
	}
}

func (enum_pedal_type *Enum_Pedal_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Pedal_type_Start":
		*enum_pedal_type = Enum_Pedal_type_Start
	case "Enum_Pedal_type_Stop":
		*enum_pedal_type = Enum_Pedal_type_Stop
	case "Enum_Pedal_type_Sostenuto":
		*enum_pedal_type = Enum_Pedal_type_Sostenuto
	case "Enum_Pedal_type_Change":
		*enum_pedal_type = Enum_Pedal_type_Change
	case "Enum_Pedal_type_Continue":
		*enum_pedal_type = Enum_Pedal_type_Continue
	case "Enum_Pedal_type_Discontinue":
		*enum_pedal_type = Enum_Pedal_type_Discontinue
	case "Enum_Pedal_type_Resume":
		*enum_pedal_type = Enum_Pedal_type_Resume
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_pedal_type *Enum_Pedal_type) ToCodeString() (res string) {

	switch *enum_pedal_type {
	// insertion code per enum code
	case Enum_Pedal_type_Start:
		res = "Enum_Pedal_type_Start"
	case Enum_Pedal_type_Stop:
		res = "Enum_Pedal_type_Stop"
	case Enum_Pedal_type_Sostenuto:
		res = "Enum_Pedal_type_Sostenuto"
	case Enum_Pedal_type_Change:
		res = "Enum_Pedal_type_Change"
	case Enum_Pedal_type_Continue:
		res = "Enum_Pedal_type_Continue"
	case Enum_Pedal_type_Discontinue:
		res = "Enum_Pedal_type_Discontinue"
	case Enum_Pedal_type_Resume:
		res = "Enum_Pedal_type_Resume"
	}
	return
}

func (enum_pedal_type Enum_Pedal_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Pedal_type_Start")
	res = append(res, "Enum_Pedal_type_Stop")
	res = append(res, "Enum_Pedal_type_Sostenuto")
	res = append(res, "Enum_Pedal_type_Change")
	res = append(res, "Enum_Pedal_type_Continue")
	res = append(res, "Enum_Pedal_type_Discontinue")
	res = append(res, "Enum_Pedal_type_Resume")

	return
}

func (enum_pedal_type Enum_Pedal_type) CodeValues() (res []string) {

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

// Utility function for Enum_Pitched_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_pitched_value Enum_Pitched_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_pitched_value {
	// insertion code per enum code
	case Enum_Pitched_value_Celesta:
		res = "celesta"
	case Enum_Pitched_value_Chimes:
		res = "chimes"
	case Enum_Pitched_value_Glockenspiel:
		res = "glockenspiel"
	case Enum_Pitched_value_Lithophone:
		res = "lithophone"
	case Enum_Pitched_value_Mallet:
		res = "mallet"
	case Enum_Pitched_value_Marimba:
		res = "marimba"
	case Enum_Pitched_value_Steel_drums:
		res = "steel drums"
	case Enum_Pitched_value_Tubaphone:
		res = "tubaphone"
	case Enum_Pitched_value_Tubular_chimes:
		res = "tubular chimes"
	case Enum_Pitched_value_Vibraphone:
		res = "vibraphone"
	case Enum_Pitched_value_Xylophone:
		res = "xylophone"
	}
	return
}

func (enum_pitched_value *Enum_Pitched_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "celesta":
		*enum_pitched_value = Enum_Pitched_value_Celesta
		return
	case "chimes":
		*enum_pitched_value = Enum_Pitched_value_Chimes
		return
	case "glockenspiel":
		*enum_pitched_value = Enum_Pitched_value_Glockenspiel
		return
	case "lithophone":
		*enum_pitched_value = Enum_Pitched_value_Lithophone
		return
	case "mallet":
		*enum_pitched_value = Enum_Pitched_value_Mallet
		return
	case "marimba":
		*enum_pitched_value = Enum_Pitched_value_Marimba
		return
	case "steel drums":
		*enum_pitched_value = Enum_Pitched_value_Steel_drums
		return
	case "tubaphone":
		*enum_pitched_value = Enum_Pitched_value_Tubaphone
		return
	case "tubular chimes":
		*enum_pitched_value = Enum_Pitched_value_Tubular_chimes
		return
	case "vibraphone":
		*enum_pitched_value = Enum_Pitched_value_Vibraphone
		return
	case "xylophone":
		*enum_pitched_value = Enum_Pitched_value_Xylophone
		return
	default:
		return errUnkownEnum
	}
}

func (enum_pitched_value *Enum_Pitched_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Pitched_value_Celesta":
		*enum_pitched_value = Enum_Pitched_value_Celesta
	case "Enum_Pitched_value_Chimes":
		*enum_pitched_value = Enum_Pitched_value_Chimes
	case "Enum_Pitched_value_Glockenspiel":
		*enum_pitched_value = Enum_Pitched_value_Glockenspiel
	case "Enum_Pitched_value_Lithophone":
		*enum_pitched_value = Enum_Pitched_value_Lithophone
	case "Enum_Pitched_value_Mallet":
		*enum_pitched_value = Enum_Pitched_value_Mallet
	case "Enum_Pitched_value_Marimba":
		*enum_pitched_value = Enum_Pitched_value_Marimba
	case "Enum_Pitched_value_Steel_drums":
		*enum_pitched_value = Enum_Pitched_value_Steel_drums
	case "Enum_Pitched_value_Tubaphone":
		*enum_pitched_value = Enum_Pitched_value_Tubaphone
	case "Enum_Pitched_value_Tubular_chimes":
		*enum_pitched_value = Enum_Pitched_value_Tubular_chimes
	case "Enum_Pitched_value_Vibraphone":
		*enum_pitched_value = Enum_Pitched_value_Vibraphone
	case "Enum_Pitched_value_Xylophone":
		*enum_pitched_value = Enum_Pitched_value_Xylophone
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_pitched_value *Enum_Pitched_value) ToCodeString() (res string) {

	switch *enum_pitched_value {
	// insertion code per enum code
	case Enum_Pitched_value_Celesta:
		res = "Enum_Pitched_value_Celesta"
	case Enum_Pitched_value_Chimes:
		res = "Enum_Pitched_value_Chimes"
	case Enum_Pitched_value_Glockenspiel:
		res = "Enum_Pitched_value_Glockenspiel"
	case Enum_Pitched_value_Lithophone:
		res = "Enum_Pitched_value_Lithophone"
	case Enum_Pitched_value_Mallet:
		res = "Enum_Pitched_value_Mallet"
	case Enum_Pitched_value_Marimba:
		res = "Enum_Pitched_value_Marimba"
	case Enum_Pitched_value_Steel_drums:
		res = "Enum_Pitched_value_Steel_drums"
	case Enum_Pitched_value_Tubaphone:
		res = "Enum_Pitched_value_Tubaphone"
	case Enum_Pitched_value_Tubular_chimes:
		res = "Enum_Pitched_value_Tubular_chimes"
	case Enum_Pitched_value_Vibraphone:
		res = "Enum_Pitched_value_Vibraphone"
	case Enum_Pitched_value_Xylophone:
		res = "Enum_Pitched_value_Xylophone"
	}
	return
}

func (enum_pitched_value Enum_Pitched_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Pitched_value_Celesta")
	res = append(res, "Enum_Pitched_value_Chimes")
	res = append(res, "Enum_Pitched_value_Glockenspiel")
	res = append(res, "Enum_Pitched_value_Lithophone")
	res = append(res, "Enum_Pitched_value_Mallet")
	res = append(res, "Enum_Pitched_value_Marimba")
	res = append(res, "Enum_Pitched_value_Steel_drums")
	res = append(res, "Enum_Pitched_value_Tubaphone")
	res = append(res, "Enum_Pitched_value_Tubular_chimes")
	res = append(res, "Enum_Pitched_value_Vibraphone")
	res = append(res, "Enum_Pitched_value_Xylophone")

	return
}

func (enum_pitched_value Enum_Pitched_value) CodeValues() (res []string) {

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

// Utility function for Enum_Principal_voice_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_principal_voice_symbol Enum_Principal_voice_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_principal_voice_symbol {
	// insertion code per enum code
	case Enum_Principal_voice_symbol_Hauptstimme:
		res = "Hauptstimme"
	case Enum_Principal_voice_symbol_Nebenstimme:
		res = "Nebenstimme"
	case Enum_Principal_voice_symbol_Plain:
		res = "plain"
	case Enum_Principal_voice_symbol_None:
		res = "none"
	}
	return
}

func (enum_principal_voice_symbol *Enum_Principal_voice_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Hauptstimme":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Hauptstimme
		return
	case "Nebenstimme":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Nebenstimme
		return
	case "plain":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Plain
		return
	case "none":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_principal_voice_symbol *Enum_Principal_voice_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Principal_voice_symbol_Hauptstimme":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Hauptstimme
	case "Enum_Principal_voice_symbol_Nebenstimme":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Nebenstimme
	case "Enum_Principal_voice_symbol_Plain":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_Plain
	case "Enum_Principal_voice_symbol_None":
		*enum_principal_voice_symbol = Enum_Principal_voice_symbol_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_principal_voice_symbol *Enum_Principal_voice_symbol) ToCodeString() (res string) {

	switch *enum_principal_voice_symbol {
	// insertion code per enum code
	case Enum_Principal_voice_symbol_Hauptstimme:
		res = "Enum_Principal_voice_symbol_Hauptstimme"
	case Enum_Principal_voice_symbol_Nebenstimme:
		res = "Enum_Principal_voice_symbol_Nebenstimme"
	case Enum_Principal_voice_symbol_Plain:
		res = "Enum_Principal_voice_symbol_Plain"
	case Enum_Principal_voice_symbol_None:
		res = "Enum_Principal_voice_symbol_None"
	}
	return
}

func (enum_principal_voice_symbol Enum_Principal_voice_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Principal_voice_symbol_Hauptstimme")
	res = append(res, "Enum_Principal_voice_symbol_Nebenstimme")
	res = append(res, "Enum_Principal_voice_symbol_Plain")
	res = append(res, "Enum_Principal_voice_symbol_None")

	return
}

func (enum_principal_voice_symbol Enum_Principal_voice_symbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Hauptstimme")
	res = append(res, "Nebenstimme")
	res = append(res, "plain")
	res = append(res, "none")

	return
}

// Utility function for Enum_Right_left_middle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_right_left_middle Enum_Right_left_middle) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_right_left_middle {
	// insertion code per enum code
	case Enum_Right_left_middle_Right:
		res = "right"
	case Enum_Right_left_middle_Left:
		res = "left"
	case Enum_Right_left_middle_Middle:
		res = "middle"
	}
	return
}

func (enum_right_left_middle *Enum_Right_left_middle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*enum_right_left_middle = Enum_Right_left_middle_Right
		return
	case "left":
		*enum_right_left_middle = Enum_Right_left_middle_Left
		return
	case "middle":
		*enum_right_left_middle = Enum_Right_left_middle_Middle
		return
	default:
		return errUnkownEnum
	}
}

func (enum_right_left_middle *Enum_Right_left_middle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Right_left_middle_Right":
		*enum_right_left_middle = Enum_Right_left_middle_Right
	case "Enum_Right_left_middle_Left":
		*enum_right_left_middle = Enum_Right_left_middle_Left
	case "Enum_Right_left_middle_Middle":
		*enum_right_left_middle = Enum_Right_left_middle_Middle
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_right_left_middle *Enum_Right_left_middle) ToCodeString() (res string) {

	switch *enum_right_left_middle {
	// insertion code per enum code
	case Enum_Right_left_middle_Right:
		res = "Enum_Right_left_middle_Right"
	case Enum_Right_left_middle_Left:
		res = "Enum_Right_left_middle_Left"
	case Enum_Right_left_middle_Middle:
		res = "Enum_Right_left_middle_Middle"
	}
	return
}

func (enum_right_left_middle Enum_Right_left_middle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Right_left_middle_Right")
	res = append(res, "Enum_Right_left_middle_Left")
	res = append(res, "Enum_Right_left_middle_Middle")

	return
}

func (enum_right_left_middle Enum_Right_left_middle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "left")
	res = append(res, "middle")

	return
}

// Utility function for Enum_Semi_pitched
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_semi_pitched Enum_Semi_pitched) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_semi_pitched {
	// insertion code per enum code
	case Enum_Semi_pitched_High:
		res = "high"
	case Enum_Semi_pitched_Medium_high:
		res = "medium-high"
	case Enum_Semi_pitched_Medium:
		res = "medium"
	case Enum_Semi_pitched_Medium_low:
		res = "medium-low"
	case Enum_Semi_pitched_Low:
		res = "low"
	case Enum_Semi_pitched_Very_low:
		res = "very-low"
	}
	return
}

func (enum_semi_pitched *Enum_Semi_pitched) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "high":
		*enum_semi_pitched = Enum_Semi_pitched_High
		return
	case "medium-high":
		*enum_semi_pitched = Enum_Semi_pitched_Medium_high
		return
	case "medium":
		*enum_semi_pitched = Enum_Semi_pitched_Medium
		return
	case "medium-low":
		*enum_semi_pitched = Enum_Semi_pitched_Medium_low
		return
	case "low":
		*enum_semi_pitched = Enum_Semi_pitched_Low
		return
	case "very-low":
		*enum_semi_pitched = Enum_Semi_pitched_Very_low
		return
	default:
		return errUnkownEnum
	}
}

func (enum_semi_pitched *Enum_Semi_pitched) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Semi_pitched_High":
		*enum_semi_pitched = Enum_Semi_pitched_High
	case "Enum_Semi_pitched_Medium_high":
		*enum_semi_pitched = Enum_Semi_pitched_Medium_high
	case "Enum_Semi_pitched_Medium":
		*enum_semi_pitched = Enum_Semi_pitched_Medium
	case "Enum_Semi_pitched_Medium_low":
		*enum_semi_pitched = Enum_Semi_pitched_Medium_low
	case "Enum_Semi_pitched_Low":
		*enum_semi_pitched = Enum_Semi_pitched_Low
	case "Enum_Semi_pitched_Very_low":
		*enum_semi_pitched = Enum_Semi_pitched_Very_low
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_semi_pitched *Enum_Semi_pitched) ToCodeString() (res string) {

	switch *enum_semi_pitched {
	// insertion code per enum code
	case Enum_Semi_pitched_High:
		res = "Enum_Semi_pitched_High"
	case Enum_Semi_pitched_Medium_high:
		res = "Enum_Semi_pitched_Medium_high"
	case Enum_Semi_pitched_Medium:
		res = "Enum_Semi_pitched_Medium"
	case Enum_Semi_pitched_Medium_low:
		res = "Enum_Semi_pitched_Medium_low"
	case Enum_Semi_pitched_Low:
		res = "Enum_Semi_pitched_Low"
	case Enum_Semi_pitched_Very_low:
		res = "Enum_Semi_pitched_Very_low"
	}
	return
}

func (enum_semi_pitched Enum_Semi_pitched) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Semi_pitched_High")
	res = append(res, "Enum_Semi_pitched_Medium_high")
	res = append(res, "Enum_Semi_pitched_Medium")
	res = append(res, "Enum_Semi_pitched_Medium_low")
	res = append(res, "Enum_Semi_pitched_Low")
	res = append(res, "Enum_Semi_pitched_Very_low")

	return
}

func (enum_semi_pitched Enum_Semi_pitched) CodeValues() (res []string) {

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

// Utility function for Enum_Show_frets
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_show_frets Enum_Show_frets) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_show_frets {
	// insertion code per enum code
	case Enum_Show_frets_Numbers:
		res = "numbers"
	case Enum_Show_frets_Letters:
		res = "letters"
	}
	return
}

func (enum_show_frets *Enum_Show_frets) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "numbers":
		*enum_show_frets = Enum_Show_frets_Numbers
		return
	case "letters":
		*enum_show_frets = Enum_Show_frets_Letters
		return
	default:
		return errUnkownEnum
	}
}

func (enum_show_frets *Enum_Show_frets) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Show_frets_Numbers":
		*enum_show_frets = Enum_Show_frets_Numbers
	case "Enum_Show_frets_Letters":
		*enum_show_frets = Enum_Show_frets_Letters
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_show_frets *Enum_Show_frets) ToCodeString() (res string) {

	switch *enum_show_frets {
	// insertion code per enum code
	case Enum_Show_frets_Numbers:
		res = "Enum_Show_frets_Numbers"
	case Enum_Show_frets_Letters:
		res = "Enum_Show_frets_Letters"
	}
	return
}

func (enum_show_frets Enum_Show_frets) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Show_frets_Numbers")
	res = append(res, "Enum_Show_frets_Letters")

	return
}

func (enum_show_frets Enum_Show_frets) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "numbers")
	res = append(res, "letters")

	return
}

// Utility function for Enum_Show_tuplet
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_show_tuplet Enum_Show_tuplet) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_show_tuplet {
	// insertion code per enum code
	case Enum_Show_tuplet_Actual:
		res = "actual"
	case Enum_Show_tuplet_Both:
		res = "both"
	case Enum_Show_tuplet_None:
		res = "none"
	}
	return
}

func (enum_show_tuplet *Enum_Show_tuplet) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "actual":
		*enum_show_tuplet = Enum_Show_tuplet_Actual
		return
	case "both":
		*enum_show_tuplet = Enum_Show_tuplet_Both
		return
	case "none":
		*enum_show_tuplet = Enum_Show_tuplet_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_show_tuplet *Enum_Show_tuplet) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Show_tuplet_Actual":
		*enum_show_tuplet = Enum_Show_tuplet_Actual
	case "Enum_Show_tuplet_Both":
		*enum_show_tuplet = Enum_Show_tuplet_Both
	case "Enum_Show_tuplet_None":
		*enum_show_tuplet = Enum_Show_tuplet_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_show_tuplet *Enum_Show_tuplet) ToCodeString() (res string) {

	switch *enum_show_tuplet {
	// insertion code per enum code
	case Enum_Show_tuplet_Actual:
		res = "Enum_Show_tuplet_Actual"
	case Enum_Show_tuplet_Both:
		res = "Enum_Show_tuplet_Both"
	case Enum_Show_tuplet_None:
		res = "Enum_Show_tuplet_None"
	}
	return
}

func (enum_show_tuplet Enum_Show_tuplet) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Show_tuplet_Actual")
	res = append(res, "Enum_Show_tuplet_Both")
	res = append(res, "Enum_Show_tuplet_None")

	return
}

func (enum_show_tuplet Enum_Show_tuplet) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "actual")
	res = append(res, "both")
	res = append(res, "none")

	return
}

// Utility function for Enum_Staff_divide_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_staff_divide_symbol Enum_Staff_divide_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_staff_divide_symbol {
	// insertion code per enum code
	case Enum_Staff_divide_symbol_Down:
		res = "down"
	case Enum_Staff_divide_symbol_Up:
		res = "up"
	case Enum_Staff_divide_symbol_Up_down:
		res = "up-down"
	}
	return
}

func (enum_staff_divide_symbol *Enum_Staff_divide_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Down
		return
	case "up":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Up
		return
	case "up-down":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Up_down
		return
	default:
		return errUnkownEnum
	}
}

func (enum_staff_divide_symbol *Enum_Staff_divide_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Staff_divide_symbol_Down":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Down
	case "Enum_Staff_divide_symbol_Up":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Up
	case "Enum_Staff_divide_symbol_Up_down":
		*enum_staff_divide_symbol = Enum_Staff_divide_symbol_Up_down
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_staff_divide_symbol *Enum_Staff_divide_symbol) ToCodeString() (res string) {

	switch *enum_staff_divide_symbol {
	// insertion code per enum code
	case Enum_Staff_divide_symbol_Down:
		res = "Enum_Staff_divide_symbol_Down"
	case Enum_Staff_divide_symbol_Up:
		res = "Enum_Staff_divide_symbol_Up"
	case Enum_Staff_divide_symbol_Up_down:
		res = "Enum_Staff_divide_symbol_Up_down"
	}
	return
}

func (enum_staff_divide_symbol Enum_Staff_divide_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Staff_divide_symbol_Down")
	res = append(res, "Enum_Staff_divide_symbol_Up")
	res = append(res, "Enum_Staff_divide_symbol_Up_down")

	return
}

func (enum_staff_divide_symbol Enum_Staff_divide_symbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "up-down")

	return
}

// Utility function for Enum_Staff_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_staff_type Enum_Staff_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_staff_type {
	// insertion code per enum code
	case Enum_Staff_type_Ossia:
		res = "ossia"
	case Enum_Staff_type_Editorial:
		res = "editorial"
	case Enum_Staff_type_Cue:
		res = "cue"
	case Enum_Staff_type_Alternate:
		res = "alternate"
	case Enum_Staff_type_Regular:
		res = "regular"
	}
	return
}

func (enum_staff_type *Enum_Staff_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ossia":
		*enum_staff_type = Enum_Staff_type_Ossia
		return
	case "editorial":
		*enum_staff_type = Enum_Staff_type_Editorial
		return
	case "cue":
		*enum_staff_type = Enum_Staff_type_Cue
		return
	case "alternate":
		*enum_staff_type = Enum_Staff_type_Alternate
		return
	case "regular":
		*enum_staff_type = Enum_Staff_type_Regular
		return
	default:
		return errUnkownEnum
	}
}

func (enum_staff_type *Enum_Staff_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Staff_type_Ossia":
		*enum_staff_type = Enum_Staff_type_Ossia
	case "Enum_Staff_type_Editorial":
		*enum_staff_type = Enum_Staff_type_Editorial
	case "Enum_Staff_type_Cue":
		*enum_staff_type = Enum_Staff_type_Cue
	case "Enum_Staff_type_Alternate":
		*enum_staff_type = Enum_Staff_type_Alternate
	case "Enum_Staff_type_Regular":
		*enum_staff_type = Enum_Staff_type_Regular
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_staff_type *Enum_Staff_type) ToCodeString() (res string) {

	switch *enum_staff_type {
	// insertion code per enum code
	case Enum_Staff_type_Ossia:
		res = "Enum_Staff_type_Ossia"
	case Enum_Staff_type_Editorial:
		res = "Enum_Staff_type_Editorial"
	case Enum_Staff_type_Cue:
		res = "Enum_Staff_type_Cue"
	case Enum_Staff_type_Alternate:
		res = "Enum_Staff_type_Alternate"
	case Enum_Staff_type_Regular:
		res = "Enum_Staff_type_Regular"
	}
	return
}

func (enum_staff_type Enum_Staff_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Staff_type_Ossia")
	res = append(res, "Enum_Staff_type_Editorial")
	res = append(res, "Enum_Staff_type_Cue")
	res = append(res, "Enum_Staff_type_Alternate")
	res = append(res, "Enum_Staff_type_Regular")

	return
}

func (enum_staff_type Enum_Staff_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ossia")
	res = append(res, "editorial")
	res = append(res, "cue")
	res = append(res, "alternate")
	res = append(res, "regular")

	return
}

// Utility function for Enum_Start_note
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_start_note Enum_Start_note) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_start_note {
	// insertion code per enum code
	case Enum_Start_note_Upper:
		res = "upper"
	case Enum_Start_note_Main:
		res = "main"
	case Enum_Start_note_Below:
		res = "below"
	}
	return
}

func (enum_start_note *Enum_Start_note) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upper":
		*enum_start_note = Enum_Start_note_Upper
		return
	case "main":
		*enum_start_note = Enum_Start_note_Main
		return
	case "below":
		*enum_start_note = Enum_Start_note_Below
		return
	default:
		return errUnkownEnum
	}
}

func (enum_start_note *Enum_Start_note) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Start_note_Upper":
		*enum_start_note = Enum_Start_note_Upper
	case "Enum_Start_note_Main":
		*enum_start_note = Enum_Start_note_Main
	case "Enum_Start_note_Below":
		*enum_start_note = Enum_Start_note_Below
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_start_note *Enum_Start_note) ToCodeString() (res string) {

	switch *enum_start_note {
	// insertion code per enum code
	case Enum_Start_note_Upper:
		res = "Enum_Start_note_Upper"
	case Enum_Start_note_Main:
		res = "Enum_Start_note_Main"
	case Enum_Start_note_Below:
		res = "Enum_Start_note_Below"
	}
	return
}

func (enum_start_note Enum_Start_note) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Start_note_Upper")
	res = append(res, "Enum_Start_note_Main")
	res = append(res, "Enum_Start_note_Below")

	return
}

func (enum_start_note Enum_Start_note) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upper")
	res = append(res, "main")
	res = append(res, "below")

	return
}

// Utility function for Enum_Start_stop
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_start_stop Enum_Start_stop) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_start_stop {
	// insertion code per enum code
	case Enum_Start_stop_Start:
		res = "start"
	case Enum_Start_stop_Stop:
		res = "stop"
	}
	return
}

func (enum_start_stop *Enum_Start_stop) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_start_stop = Enum_Start_stop_Start
		return
	case "stop":
		*enum_start_stop = Enum_Start_stop_Stop
		return
	default:
		return errUnkownEnum
	}
}

func (enum_start_stop *Enum_Start_stop) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Start_stop_Start":
		*enum_start_stop = Enum_Start_stop_Start
	case "Enum_Start_stop_Stop":
		*enum_start_stop = Enum_Start_stop_Stop
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_start_stop *Enum_Start_stop) ToCodeString() (res string) {

	switch *enum_start_stop {
	// insertion code per enum code
	case Enum_Start_stop_Start:
		res = "Enum_Start_stop_Start"
	case Enum_Start_stop_Stop:
		res = "Enum_Start_stop_Stop"
	}
	return
}

func (enum_start_stop Enum_Start_stop) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Start_stop_Start")
	res = append(res, "Enum_Start_stop_Stop")

	return
}

func (enum_start_stop Enum_Start_stop) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")

	return
}

// Utility function for Enum_Start_stop_continue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_start_stop_continue Enum_Start_stop_continue) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_start_stop_continue {
	// insertion code per enum code
	case Enum_Start_stop_continue_Start:
		res = "start"
	case Enum_Start_stop_continue_Stop:
		res = "stop"
	case Enum_Start_stop_continue_Continue:
		res = "continue"
	}
	return
}

func (enum_start_stop_continue *Enum_Start_stop_continue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_start_stop_continue = Enum_Start_stop_continue_Start
		return
	case "stop":
		*enum_start_stop_continue = Enum_Start_stop_continue_Stop
		return
	case "continue":
		*enum_start_stop_continue = Enum_Start_stop_continue_Continue
		return
	default:
		return errUnkownEnum
	}
}

func (enum_start_stop_continue *Enum_Start_stop_continue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Start_stop_continue_Start":
		*enum_start_stop_continue = Enum_Start_stop_continue_Start
	case "Enum_Start_stop_continue_Stop":
		*enum_start_stop_continue = Enum_Start_stop_continue_Stop
	case "Enum_Start_stop_continue_Continue":
		*enum_start_stop_continue = Enum_Start_stop_continue_Continue
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_start_stop_continue *Enum_Start_stop_continue) ToCodeString() (res string) {

	switch *enum_start_stop_continue {
	// insertion code per enum code
	case Enum_Start_stop_continue_Start:
		res = "Enum_Start_stop_continue_Start"
	case Enum_Start_stop_continue_Stop:
		res = "Enum_Start_stop_continue_Stop"
	case Enum_Start_stop_continue_Continue:
		res = "Enum_Start_stop_continue_Continue"
	}
	return
}

func (enum_start_stop_continue Enum_Start_stop_continue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Start_stop_continue_Start")
	res = append(res, "Enum_Start_stop_continue_Stop")
	res = append(res, "Enum_Start_stop_continue_Continue")

	return
}

func (enum_start_stop_continue Enum_Start_stop_continue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Enum_Start_stop_discontinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_start_stop_discontinue Enum_Start_stop_discontinue) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_start_stop_discontinue {
	// insertion code per enum code
	case Enum_Start_stop_discontinue_Start:
		res = "start"
	case Enum_Start_stop_discontinue_Stop:
		res = "stop"
	case Enum_Start_stop_discontinue_Discontinue:
		res = "discontinue"
	}
	return
}

func (enum_start_stop_discontinue *Enum_Start_stop_discontinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Start
		return
	case "stop":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Stop
		return
	case "discontinue":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Discontinue
		return
	default:
		return errUnkownEnum
	}
}

func (enum_start_stop_discontinue *Enum_Start_stop_discontinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Start_stop_discontinue_Start":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Start
	case "Enum_Start_stop_discontinue_Stop":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Stop
	case "Enum_Start_stop_discontinue_Discontinue":
		*enum_start_stop_discontinue = Enum_Start_stop_discontinue_Discontinue
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_start_stop_discontinue *Enum_Start_stop_discontinue) ToCodeString() (res string) {

	switch *enum_start_stop_discontinue {
	// insertion code per enum code
	case Enum_Start_stop_discontinue_Start:
		res = "Enum_Start_stop_discontinue_Start"
	case Enum_Start_stop_discontinue_Stop:
		res = "Enum_Start_stop_discontinue_Stop"
	case Enum_Start_stop_discontinue_Discontinue:
		res = "Enum_Start_stop_discontinue_Discontinue"
	}
	return
}

func (enum_start_stop_discontinue Enum_Start_stop_discontinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Start_stop_discontinue_Start")
	res = append(res, "Enum_Start_stop_discontinue_Stop")
	res = append(res, "Enum_Start_stop_discontinue_Discontinue")

	return
}

func (enum_start_stop_discontinue Enum_Start_stop_discontinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "discontinue")

	return
}

// Utility function for Enum_Start_stop_single
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_start_stop_single Enum_Start_stop_single) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_start_stop_single {
	// insertion code per enum code
	case Enum_Start_stop_single_Start:
		res = "start"
	case Enum_Start_stop_single_Stop:
		res = "stop"
	case Enum_Start_stop_single_Single:
		res = "single"
	}
	return
}

func (enum_start_stop_single *Enum_Start_stop_single) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_start_stop_single = Enum_Start_stop_single_Start
		return
	case "stop":
		*enum_start_stop_single = Enum_Start_stop_single_Stop
		return
	case "single":
		*enum_start_stop_single = Enum_Start_stop_single_Single
		return
	default:
		return errUnkownEnum
	}
}

func (enum_start_stop_single *Enum_Start_stop_single) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Start_stop_single_Start":
		*enum_start_stop_single = Enum_Start_stop_single_Start
	case "Enum_Start_stop_single_Stop":
		*enum_start_stop_single = Enum_Start_stop_single_Stop
	case "Enum_Start_stop_single_Single":
		*enum_start_stop_single = Enum_Start_stop_single_Single
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_start_stop_single *Enum_Start_stop_single) ToCodeString() (res string) {

	switch *enum_start_stop_single {
	// insertion code per enum code
	case Enum_Start_stop_single_Start:
		res = "Enum_Start_stop_single_Start"
	case Enum_Start_stop_single_Stop:
		res = "Enum_Start_stop_single_Stop"
	case Enum_Start_stop_single_Single:
		res = "Enum_Start_stop_single_Single"
	}
	return
}

func (enum_start_stop_single Enum_Start_stop_single) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Start_stop_single_Start")
	res = append(res, "Enum_Start_stop_single_Stop")
	res = append(res, "Enum_Start_stop_single_Single")

	return
}

func (enum_start_stop_single Enum_Start_stop_single) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")

	return
}

// Utility function for Enum_Stem_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_stem_value Enum_Stem_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_stem_value {
	// insertion code per enum code
	case Enum_Stem_value_Down:
		res = "down"
	case Enum_Stem_value_Up:
		res = "up"
	case Enum_Stem_value_Double:
		res = "double"
	case Enum_Stem_value_None:
		res = "none"
	}
	return
}

func (enum_stem_value *Enum_Stem_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*enum_stem_value = Enum_Stem_value_Down
		return
	case "up":
		*enum_stem_value = Enum_Stem_value_Up
		return
	case "double":
		*enum_stem_value = Enum_Stem_value_Double
		return
	case "none":
		*enum_stem_value = Enum_Stem_value_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_stem_value *Enum_Stem_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Stem_value_Down":
		*enum_stem_value = Enum_Stem_value_Down
	case "Enum_Stem_value_Up":
		*enum_stem_value = Enum_Stem_value_Up
	case "Enum_Stem_value_Double":
		*enum_stem_value = Enum_Stem_value_Double
	case "Enum_Stem_value_None":
		*enum_stem_value = Enum_Stem_value_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_stem_value *Enum_Stem_value) ToCodeString() (res string) {

	switch *enum_stem_value {
	// insertion code per enum code
	case Enum_Stem_value_Down:
		res = "Enum_Stem_value_Down"
	case Enum_Stem_value_Up:
		res = "Enum_Stem_value_Up"
	case Enum_Stem_value_Double:
		res = "Enum_Stem_value_Double"
	case Enum_Stem_value_None:
		res = "Enum_Stem_value_None"
	}
	return
}

func (enum_stem_value Enum_Stem_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Stem_value_Down")
	res = append(res, "Enum_Stem_value_Up")
	res = append(res, "Enum_Stem_value_Double")
	res = append(res, "Enum_Stem_value_None")

	return
}

func (enum_stem_value Enum_Stem_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "double")
	res = append(res, "none")

	return
}

// Utility function for Enum_Step
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_step Enum_Step) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_step {
	// insertion code per enum code
	case Enum_Step_A:
		res = "A"
	case Enum_Step_B:
		res = "B"
	case Enum_Step_C:
		res = "C"
	case Enum_Step_D:
		res = "D"
	case Enum_Step_E:
		res = "E"
	case Enum_Step_F:
		res = "F"
	case Enum_Step_G:
		res = "G"
	}
	return
}

func (enum_step *Enum_Step) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_step = Enum_Step_A
		return
	case "B":
		*enum_step = Enum_Step_B
		return
	case "C":
		*enum_step = Enum_Step_C
		return
	case "D":
		*enum_step = Enum_Step_D
		return
	case "E":
		*enum_step = Enum_Step_E
		return
	case "F":
		*enum_step = Enum_Step_F
		return
	case "G":
		*enum_step = Enum_Step_G
		return
	default:
		return errUnkownEnum
	}
}

func (enum_step *Enum_Step) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Step_A":
		*enum_step = Enum_Step_A
	case "Enum_Step_B":
		*enum_step = Enum_Step_B
	case "Enum_Step_C":
		*enum_step = Enum_Step_C
	case "Enum_Step_D":
		*enum_step = Enum_Step_D
	case "Enum_Step_E":
		*enum_step = Enum_Step_E
	case "Enum_Step_F":
		*enum_step = Enum_Step_F
	case "Enum_Step_G":
		*enum_step = Enum_Step_G
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_step *Enum_Step) ToCodeString() (res string) {

	switch *enum_step {
	// insertion code per enum code
	case Enum_Step_A:
		res = "Enum_Step_A"
	case Enum_Step_B:
		res = "Enum_Step_B"
	case Enum_Step_C:
		res = "Enum_Step_C"
	case Enum_Step_D:
		res = "Enum_Step_D"
	case Enum_Step_E:
		res = "Enum_Step_E"
	case Enum_Step_F:
		res = "Enum_Step_F"
	case Enum_Step_G:
		res = "Enum_Step_G"
	}
	return
}

func (enum_step Enum_Step) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Step_A")
	res = append(res, "Enum_Step_B")
	res = append(res, "Enum_Step_C")
	res = append(res, "Enum_Step_D")
	res = append(res, "Enum_Step_E")
	res = append(res, "Enum_Step_F")
	res = append(res, "Enum_Step_G")

	return
}

func (enum_step Enum_Step) CodeValues() (res []string) {

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

// Utility function for Enum_Stick_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_stick_location Enum_Stick_location) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_stick_location {
	// insertion code per enum code
	case Enum_Stick_location_Center:
		res = "center"
	case Enum_Stick_location_Rim:
		res = "rim"
	case Enum_Stick_location_Cymbal_bell:
		res = "cymbal bell"
	case Enum_Stick_location_Cymbal_edge:
		res = "cymbal edge"
	}
	return
}

func (enum_stick_location *Enum_Stick_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "center":
		*enum_stick_location = Enum_Stick_location_Center
		return
	case "rim":
		*enum_stick_location = Enum_Stick_location_Rim
		return
	case "cymbal bell":
		*enum_stick_location = Enum_Stick_location_Cymbal_bell
		return
	case "cymbal edge":
		*enum_stick_location = Enum_Stick_location_Cymbal_edge
		return
	default:
		return errUnkownEnum
	}
}

func (enum_stick_location *Enum_Stick_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Stick_location_Center":
		*enum_stick_location = Enum_Stick_location_Center
	case "Enum_Stick_location_Rim":
		*enum_stick_location = Enum_Stick_location_Rim
	case "Enum_Stick_location_Cymbal_bell":
		*enum_stick_location = Enum_Stick_location_Cymbal_bell
	case "Enum_Stick_location_Cymbal_edge":
		*enum_stick_location = Enum_Stick_location_Cymbal_edge
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_stick_location *Enum_Stick_location) ToCodeString() (res string) {

	switch *enum_stick_location {
	// insertion code per enum code
	case Enum_Stick_location_Center:
		res = "Enum_Stick_location_Center"
	case Enum_Stick_location_Rim:
		res = "Enum_Stick_location_Rim"
	case Enum_Stick_location_Cymbal_bell:
		res = "Enum_Stick_location_Cymbal_bell"
	case Enum_Stick_location_Cymbal_edge:
		res = "Enum_Stick_location_Cymbal_edge"
	}
	return
}

func (enum_stick_location Enum_Stick_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Stick_location_Center")
	res = append(res, "Enum_Stick_location_Rim")
	res = append(res, "Enum_Stick_location_Cymbal_bell")
	res = append(res, "Enum_Stick_location_Cymbal_edge")

	return
}

func (enum_stick_location Enum_Stick_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "center")
	res = append(res, "rim")
	res = append(res, "cymbal bell")
	res = append(res, "cymbal edge")

	return
}

// Utility function for Enum_Stick_material
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_stick_material Enum_Stick_material) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_stick_material {
	// insertion code per enum code
	case Enum_Stick_material_Soft:
		res = "soft"
	case Enum_Stick_material_Medium:
		res = "medium"
	case Enum_Stick_material_Hard:
		res = "hard"
	case Enum_Stick_material_Shaded:
		res = "shaded"
	case Enum_Stick_material_X:
		res = "x"
	}
	return
}

func (enum_stick_material *Enum_Stick_material) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "soft":
		*enum_stick_material = Enum_Stick_material_Soft
		return
	case "medium":
		*enum_stick_material = Enum_Stick_material_Medium
		return
	case "hard":
		*enum_stick_material = Enum_Stick_material_Hard
		return
	case "shaded":
		*enum_stick_material = Enum_Stick_material_Shaded
		return
	case "x":
		*enum_stick_material = Enum_Stick_material_X
		return
	default:
		return errUnkownEnum
	}
}

func (enum_stick_material *Enum_Stick_material) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Stick_material_Soft":
		*enum_stick_material = Enum_Stick_material_Soft
	case "Enum_Stick_material_Medium":
		*enum_stick_material = Enum_Stick_material_Medium
	case "Enum_Stick_material_Hard":
		*enum_stick_material = Enum_Stick_material_Hard
	case "Enum_Stick_material_Shaded":
		*enum_stick_material = Enum_Stick_material_Shaded
	case "Enum_Stick_material_X":
		*enum_stick_material = Enum_Stick_material_X
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_stick_material *Enum_Stick_material) ToCodeString() (res string) {

	switch *enum_stick_material {
	// insertion code per enum code
	case Enum_Stick_material_Soft:
		res = "Enum_Stick_material_Soft"
	case Enum_Stick_material_Medium:
		res = "Enum_Stick_material_Medium"
	case Enum_Stick_material_Hard:
		res = "Enum_Stick_material_Hard"
	case Enum_Stick_material_Shaded:
		res = "Enum_Stick_material_Shaded"
	case Enum_Stick_material_X:
		res = "Enum_Stick_material_X"
	}
	return
}

func (enum_stick_material Enum_Stick_material) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Stick_material_Soft")
	res = append(res, "Enum_Stick_material_Medium")
	res = append(res, "Enum_Stick_material_Hard")
	res = append(res, "Enum_Stick_material_Shaded")
	res = append(res, "Enum_Stick_material_X")

	return
}

func (enum_stick_material Enum_Stick_material) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "soft")
	res = append(res, "medium")
	res = append(res, "hard")
	res = append(res, "shaded")
	res = append(res, "x")

	return
}

// Utility function for Enum_Stick_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_stick_type Enum_Stick_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_stick_type {
	// insertion code per enum code
	case Enum_Stick_type_Bass_drum:
		res = "bass drum"
	case Enum_Stick_type_Double_bass_drum:
		res = "double bass drum"
	case Enum_Stick_type_Glockenspiel:
		res = "glockenspiel"
	case Enum_Stick_type_Gum:
		res = "gum"
	case Enum_Stick_type_Hammer:
		res = "hammer"
	case Enum_Stick_type_Superball:
		res = "superball"
	case Enum_Stick_type_Timpani:
		res = "timpani"
	case Enum_Stick_type_Wound:
		res = "wound"
	case Enum_Stick_type_Xylophone:
		res = "xylophone"
	case Enum_Stick_type_Yarn:
		res = "yarn"
	}
	return
}

func (enum_stick_type *Enum_Stick_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*enum_stick_type = Enum_Stick_type_Bass_drum
		return
	case "double bass drum":
		*enum_stick_type = Enum_Stick_type_Double_bass_drum
		return
	case "glockenspiel":
		*enum_stick_type = Enum_Stick_type_Glockenspiel
		return
	case "gum":
		*enum_stick_type = Enum_Stick_type_Gum
		return
	case "hammer":
		*enum_stick_type = Enum_Stick_type_Hammer
		return
	case "superball":
		*enum_stick_type = Enum_Stick_type_Superball
		return
	case "timpani":
		*enum_stick_type = Enum_Stick_type_Timpani
		return
	case "wound":
		*enum_stick_type = Enum_Stick_type_Wound
		return
	case "xylophone":
		*enum_stick_type = Enum_Stick_type_Xylophone
		return
	case "yarn":
		*enum_stick_type = Enum_Stick_type_Yarn
		return
	default:
		return errUnkownEnum
	}
}

func (enum_stick_type *Enum_Stick_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Stick_type_Bass_drum":
		*enum_stick_type = Enum_Stick_type_Bass_drum
	case "Enum_Stick_type_Double_bass_drum":
		*enum_stick_type = Enum_Stick_type_Double_bass_drum
	case "Enum_Stick_type_Glockenspiel":
		*enum_stick_type = Enum_Stick_type_Glockenspiel
	case "Enum_Stick_type_Gum":
		*enum_stick_type = Enum_Stick_type_Gum
	case "Enum_Stick_type_Hammer":
		*enum_stick_type = Enum_Stick_type_Hammer
	case "Enum_Stick_type_Superball":
		*enum_stick_type = Enum_Stick_type_Superball
	case "Enum_Stick_type_Timpani":
		*enum_stick_type = Enum_Stick_type_Timpani
	case "Enum_Stick_type_Wound":
		*enum_stick_type = Enum_Stick_type_Wound
	case "Enum_Stick_type_Xylophone":
		*enum_stick_type = Enum_Stick_type_Xylophone
	case "Enum_Stick_type_Yarn":
		*enum_stick_type = Enum_Stick_type_Yarn
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_stick_type *Enum_Stick_type) ToCodeString() (res string) {

	switch *enum_stick_type {
	// insertion code per enum code
	case Enum_Stick_type_Bass_drum:
		res = "Enum_Stick_type_Bass_drum"
	case Enum_Stick_type_Double_bass_drum:
		res = "Enum_Stick_type_Double_bass_drum"
	case Enum_Stick_type_Glockenspiel:
		res = "Enum_Stick_type_Glockenspiel"
	case Enum_Stick_type_Gum:
		res = "Enum_Stick_type_Gum"
	case Enum_Stick_type_Hammer:
		res = "Enum_Stick_type_Hammer"
	case Enum_Stick_type_Superball:
		res = "Enum_Stick_type_Superball"
	case Enum_Stick_type_Timpani:
		res = "Enum_Stick_type_Timpani"
	case Enum_Stick_type_Wound:
		res = "Enum_Stick_type_Wound"
	case Enum_Stick_type_Xylophone:
		res = "Enum_Stick_type_Xylophone"
	case Enum_Stick_type_Yarn:
		res = "Enum_Stick_type_Yarn"
	}
	return
}

func (enum_stick_type Enum_Stick_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Stick_type_Bass_drum")
	res = append(res, "Enum_Stick_type_Double_bass_drum")
	res = append(res, "Enum_Stick_type_Glockenspiel")
	res = append(res, "Enum_Stick_type_Gum")
	res = append(res, "Enum_Stick_type_Hammer")
	res = append(res, "Enum_Stick_type_Superball")
	res = append(res, "Enum_Stick_type_Timpani")
	res = append(res, "Enum_Stick_type_Wound")
	res = append(res, "Enum_Stick_type_Xylophone")
	res = append(res, "Enum_Stick_type_Yarn")

	return
}

func (enum_stick_type Enum_Stick_type) CodeValues() (res []string) {

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

// Utility function for Enum_Syllabic
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_syllabic Enum_Syllabic) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_syllabic {
	// insertion code per enum code
	case Enum_Syllabic_Single:
		res = "single"
	case Enum_Syllabic_Begin:
		res = "begin"
	case Enum_Syllabic_End:
		res = "end"
	case Enum_Syllabic_Middle:
		res = "middle"
	}
	return
}

func (enum_syllabic *Enum_Syllabic) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*enum_syllabic = Enum_Syllabic_Single
		return
	case "begin":
		*enum_syllabic = Enum_Syllabic_Begin
		return
	case "end":
		*enum_syllabic = Enum_Syllabic_End
		return
	case "middle":
		*enum_syllabic = Enum_Syllabic_Middle
		return
	default:
		return errUnkownEnum
	}
}

func (enum_syllabic *Enum_Syllabic) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Syllabic_Single":
		*enum_syllabic = Enum_Syllabic_Single
	case "Enum_Syllabic_Begin":
		*enum_syllabic = Enum_Syllabic_Begin
	case "Enum_Syllabic_End":
		*enum_syllabic = Enum_Syllabic_End
	case "Enum_Syllabic_Middle":
		*enum_syllabic = Enum_Syllabic_Middle
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_syllabic *Enum_Syllabic) ToCodeString() (res string) {

	switch *enum_syllabic {
	// insertion code per enum code
	case Enum_Syllabic_Single:
		res = "Enum_Syllabic_Single"
	case Enum_Syllabic_Begin:
		res = "Enum_Syllabic_Begin"
	case Enum_Syllabic_End:
		res = "Enum_Syllabic_End"
	case Enum_Syllabic_Middle:
		res = "Enum_Syllabic_Middle"
	}
	return
}

func (enum_syllabic Enum_Syllabic) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Syllabic_Single")
	res = append(res, "Enum_Syllabic_Begin")
	res = append(res, "Enum_Syllabic_End")
	res = append(res, "Enum_Syllabic_Middle")

	return
}

func (enum_syllabic Enum_Syllabic) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "single")
	res = append(res, "begin")
	res = append(res, "end")
	res = append(res, "middle")

	return
}

// Utility function for Enum_Symbol_size
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_symbol_size Enum_Symbol_size) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_symbol_size {
	// insertion code per enum code
	case Enum_Symbol_size_Full:
		res = "full"
	case Enum_Symbol_size_Cue:
		res = "cue"
	case Enum_Symbol_size_Grace_cue:
		res = "grace-cue"
	case Enum_Symbol_size_Large:
		res = "large"
	}
	return
}

func (enum_symbol_size *Enum_Symbol_size) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "full":
		*enum_symbol_size = Enum_Symbol_size_Full
		return
	case "cue":
		*enum_symbol_size = Enum_Symbol_size_Cue
		return
	case "grace-cue":
		*enum_symbol_size = Enum_Symbol_size_Grace_cue
		return
	case "large":
		*enum_symbol_size = Enum_Symbol_size_Large
		return
	default:
		return errUnkownEnum
	}
}

func (enum_symbol_size *Enum_Symbol_size) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Symbol_size_Full":
		*enum_symbol_size = Enum_Symbol_size_Full
	case "Enum_Symbol_size_Cue":
		*enum_symbol_size = Enum_Symbol_size_Cue
	case "Enum_Symbol_size_Grace_cue":
		*enum_symbol_size = Enum_Symbol_size_Grace_cue
	case "Enum_Symbol_size_Large":
		*enum_symbol_size = Enum_Symbol_size_Large
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_symbol_size *Enum_Symbol_size) ToCodeString() (res string) {

	switch *enum_symbol_size {
	// insertion code per enum code
	case Enum_Symbol_size_Full:
		res = "Enum_Symbol_size_Full"
	case Enum_Symbol_size_Cue:
		res = "Enum_Symbol_size_Cue"
	case Enum_Symbol_size_Grace_cue:
		res = "Enum_Symbol_size_Grace_cue"
	case Enum_Symbol_size_Large:
		res = "Enum_Symbol_size_Large"
	}
	return
}

func (enum_symbol_size Enum_Symbol_size) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Symbol_size_Full")
	res = append(res, "Enum_Symbol_size_Cue")
	res = append(res, "Enum_Symbol_size_Grace_cue")
	res = append(res, "Enum_Symbol_size_Large")

	return
}

func (enum_symbol_size Enum_Symbol_size) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "full")
	res = append(res, "cue")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for Enum_Sync_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_sync_type Enum_Sync_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_sync_type {
	// insertion code per enum code
	case Enum_Sync_type_None:
		res = "none"
	case Enum_Sync_type_Tempo:
		res = "tempo"
	case Enum_Sync_type_Mostly_tempo:
		res = "mostly-tempo"
	case Enum_Sync_type_Mostly_event:
		res = "mostly-event"
	case Enum_Sync_type_Event:
		res = "event"
	case Enum_Sync_type_Always_event:
		res = "always-event"
	}
	return
}

func (enum_sync_type *Enum_Sync_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_sync_type = Enum_Sync_type_None
		return
	case "tempo":
		*enum_sync_type = Enum_Sync_type_Tempo
		return
	case "mostly-tempo":
		*enum_sync_type = Enum_Sync_type_Mostly_tempo
		return
	case "mostly-event":
		*enum_sync_type = Enum_Sync_type_Mostly_event
		return
	case "event":
		*enum_sync_type = Enum_Sync_type_Event
		return
	case "always-event":
		*enum_sync_type = Enum_Sync_type_Always_event
		return
	default:
		return errUnkownEnum
	}
}

func (enum_sync_type *Enum_Sync_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Sync_type_None":
		*enum_sync_type = Enum_Sync_type_None
	case "Enum_Sync_type_Tempo":
		*enum_sync_type = Enum_Sync_type_Tempo
	case "Enum_Sync_type_Mostly_tempo":
		*enum_sync_type = Enum_Sync_type_Mostly_tempo
	case "Enum_Sync_type_Mostly_event":
		*enum_sync_type = Enum_Sync_type_Mostly_event
	case "Enum_Sync_type_Event":
		*enum_sync_type = Enum_Sync_type_Event
	case "Enum_Sync_type_Always_event":
		*enum_sync_type = Enum_Sync_type_Always_event
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_sync_type *Enum_Sync_type) ToCodeString() (res string) {

	switch *enum_sync_type {
	// insertion code per enum code
	case Enum_Sync_type_None:
		res = "Enum_Sync_type_None"
	case Enum_Sync_type_Tempo:
		res = "Enum_Sync_type_Tempo"
	case Enum_Sync_type_Mostly_tempo:
		res = "Enum_Sync_type_Mostly_tempo"
	case Enum_Sync_type_Mostly_event:
		res = "Enum_Sync_type_Mostly_event"
	case Enum_Sync_type_Event:
		res = "Enum_Sync_type_Event"
	case Enum_Sync_type_Always_event:
		res = "Enum_Sync_type_Always_event"
	}
	return
}

func (enum_sync_type Enum_Sync_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Sync_type_None")
	res = append(res, "Enum_Sync_type_Tempo")
	res = append(res, "Enum_Sync_type_Mostly_tempo")
	res = append(res, "Enum_Sync_type_Mostly_event")
	res = append(res, "Enum_Sync_type_Event")
	res = append(res, "Enum_Sync_type_Always_event")

	return
}

func (enum_sync_type Enum_Sync_type) CodeValues() (res []string) {

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

// Utility function for Enum_System_relation_number
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_system_relation_number Enum_System_relation_number) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_system_relation_number {
	// insertion code per enum code
	case Enum_System_relation_number_Only_top:
		res = "only-top"
	case Enum_System_relation_number_Only_bottom:
		res = "only-bottom"
	case Enum_System_relation_number_Also_top:
		res = "also-top"
	case Enum_System_relation_number_Also_bottom:
		res = "also-bottom"
	case Enum_System_relation_number_None:
		res = "none"
	}
	return
}

func (enum_system_relation_number *Enum_System_relation_number) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "only-top":
		*enum_system_relation_number = Enum_System_relation_number_Only_top
		return
	case "only-bottom":
		*enum_system_relation_number = Enum_System_relation_number_Only_bottom
		return
	case "also-top":
		*enum_system_relation_number = Enum_System_relation_number_Also_top
		return
	case "also-bottom":
		*enum_system_relation_number = Enum_System_relation_number_Also_bottom
		return
	case "none":
		*enum_system_relation_number = Enum_System_relation_number_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_system_relation_number *Enum_System_relation_number) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_System_relation_number_Only_top":
		*enum_system_relation_number = Enum_System_relation_number_Only_top
	case "Enum_System_relation_number_Only_bottom":
		*enum_system_relation_number = Enum_System_relation_number_Only_bottom
	case "Enum_System_relation_number_Also_top":
		*enum_system_relation_number = Enum_System_relation_number_Also_top
	case "Enum_System_relation_number_Also_bottom":
		*enum_system_relation_number = Enum_System_relation_number_Also_bottom
	case "Enum_System_relation_number_None":
		*enum_system_relation_number = Enum_System_relation_number_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_system_relation_number *Enum_System_relation_number) ToCodeString() (res string) {

	switch *enum_system_relation_number {
	// insertion code per enum code
	case Enum_System_relation_number_Only_top:
		res = "Enum_System_relation_number_Only_top"
	case Enum_System_relation_number_Only_bottom:
		res = "Enum_System_relation_number_Only_bottom"
	case Enum_System_relation_number_Also_top:
		res = "Enum_System_relation_number_Also_top"
	case Enum_System_relation_number_Also_bottom:
		res = "Enum_System_relation_number_Also_bottom"
	case Enum_System_relation_number_None:
		res = "Enum_System_relation_number_None"
	}
	return
}

func (enum_system_relation_number Enum_System_relation_number) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_System_relation_number_Only_top")
	res = append(res, "Enum_System_relation_number_Only_bottom")
	res = append(res, "Enum_System_relation_number_Also_top")
	res = append(res, "Enum_System_relation_number_Also_bottom")
	res = append(res, "Enum_System_relation_number_None")

	return
}

func (enum_system_relation_number Enum_System_relation_number) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "only-top")
	res = append(res, "only-bottom")
	res = append(res, "also-top")
	res = append(res, "also-bottom")
	res = append(res, "none")

	return
}

// Utility function for Enum_Tap_hand
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_tap_hand Enum_Tap_hand) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_tap_hand {
	// insertion code per enum code
	case Enum_Tap_hand_Left:
		res = "left"
	case Enum_Tap_hand_Right:
		res = "right"
	}
	return
}

func (enum_tap_hand *Enum_Tap_hand) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*enum_tap_hand = Enum_Tap_hand_Left
		return
	case "right":
		*enum_tap_hand = Enum_Tap_hand_Right
		return
	default:
		return errUnkownEnum
	}
}

func (enum_tap_hand *Enum_Tap_hand) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Tap_hand_Left":
		*enum_tap_hand = Enum_Tap_hand_Left
	case "Enum_Tap_hand_Right":
		*enum_tap_hand = Enum_Tap_hand_Right
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_tap_hand *Enum_Tap_hand) ToCodeString() (res string) {

	switch *enum_tap_hand {
	// insertion code per enum code
	case Enum_Tap_hand_Left:
		res = "Enum_Tap_hand_Left"
	case Enum_Tap_hand_Right:
		res = "Enum_Tap_hand_Right"
	}
	return
}

func (enum_tap_hand Enum_Tap_hand) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Tap_hand_Left")
	res = append(res, "Enum_Tap_hand_Right")

	return
}

func (enum_tap_hand Enum_Tap_hand) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for Enum_Text_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_text_direction Enum_Text_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_text_direction {
	// insertion code per enum code
	case Enum_Text_direction_Ltr:
		res = "ltr"
	case Enum_Text_direction_Rtl:
		res = "rtl"
	case Enum_Text_direction_Lro:
		res = "lro"
	case Enum_Text_direction_Rlo:
		res = "rlo"
	}
	return
}

func (enum_text_direction *Enum_Text_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ltr":
		*enum_text_direction = Enum_Text_direction_Ltr
		return
	case "rtl":
		*enum_text_direction = Enum_Text_direction_Rtl
		return
	case "lro":
		*enum_text_direction = Enum_Text_direction_Lro
		return
	case "rlo":
		*enum_text_direction = Enum_Text_direction_Rlo
		return
	default:
		return errUnkownEnum
	}
}

func (enum_text_direction *Enum_Text_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Text_direction_Ltr":
		*enum_text_direction = Enum_Text_direction_Ltr
	case "Enum_Text_direction_Rtl":
		*enum_text_direction = Enum_Text_direction_Rtl
	case "Enum_Text_direction_Lro":
		*enum_text_direction = Enum_Text_direction_Lro
	case "Enum_Text_direction_Rlo":
		*enum_text_direction = Enum_Text_direction_Rlo
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_text_direction *Enum_Text_direction) ToCodeString() (res string) {

	switch *enum_text_direction {
	// insertion code per enum code
	case Enum_Text_direction_Ltr:
		res = "Enum_Text_direction_Ltr"
	case Enum_Text_direction_Rtl:
		res = "Enum_Text_direction_Rtl"
	case Enum_Text_direction_Lro:
		res = "Enum_Text_direction_Lro"
	case Enum_Text_direction_Rlo:
		res = "Enum_Text_direction_Rlo"
	}
	return
}

func (enum_text_direction Enum_Text_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Text_direction_Ltr")
	res = append(res, "Enum_Text_direction_Rtl")
	res = append(res, "Enum_Text_direction_Lro")
	res = append(res, "Enum_Text_direction_Rlo")

	return
}

func (enum_text_direction Enum_Text_direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ltr")
	res = append(res, "rtl")
	res = append(res, "lro")
	res = append(res, "rlo")

	return
}

// Utility function for Enum_Tied_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_tied_type Enum_Tied_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_tied_type {
	// insertion code per enum code
	case Enum_Tied_type_Start:
		res = "start"
	case Enum_Tied_type_Stop:
		res = "stop"
	case Enum_Tied_type_Continue:
		res = "continue"
	case Enum_Tied_type_Let_ring:
		res = "let-ring"
	}
	return
}

func (enum_tied_type *Enum_Tied_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_tied_type = Enum_Tied_type_Start
		return
	case "stop":
		*enum_tied_type = Enum_Tied_type_Stop
		return
	case "continue":
		*enum_tied_type = Enum_Tied_type_Continue
		return
	case "let-ring":
		*enum_tied_type = Enum_Tied_type_Let_ring
		return
	default:
		return errUnkownEnum
	}
}

func (enum_tied_type *Enum_Tied_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Tied_type_Start":
		*enum_tied_type = Enum_Tied_type_Start
	case "Enum_Tied_type_Stop":
		*enum_tied_type = Enum_Tied_type_Stop
	case "Enum_Tied_type_Continue":
		*enum_tied_type = Enum_Tied_type_Continue
	case "Enum_Tied_type_Let_ring":
		*enum_tied_type = Enum_Tied_type_Let_ring
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_tied_type *Enum_Tied_type) ToCodeString() (res string) {

	switch *enum_tied_type {
	// insertion code per enum code
	case Enum_Tied_type_Start:
		res = "Enum_Tied_type_Start"
	case Enum_Tied_type_Stop:
		res = "Enum_Tied_type_Stop"
	case Enum_Tied_type_Continue:
		res = "Enum_Tied_type_Continue"
	case Enum_Tied_type_Let_ring:
		res = "Enum_Tied_type_Let_ring"
	}
	return
}

func (enum_tied_type Enum_Tied_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Tied_type_Start")
	res = append(res, "Enum_Tied_type_Stop")
	res = append(res, "Enum_Tied_type_Continue")
	res = append(res, "Enum_Tied_type_Let_ring")

	return
}

func (enum_tied_type Enum_Tied_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")
	res = append(res, "let-ring")

	return
}

// Utility function for Enum_Time_only
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_time_only Enum_Time_only) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_time_only {
	// insertion code per enum code
	}
	return
}

func (enum_time_only *Enum_Time_only) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (enum_time_only *Enum_Time_only) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_time_only *Enum_Time_only) ToCodeString() (res string) {

	switch *enum_time_only {
	// insertion code per enum code
	}
	return
}

func (enum_time_only Enum_Time_only) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (enum_time_only Enum_Time_only) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Enum_Time_relation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_time_relation Enum_Time_relation) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_time_relation {
	// insertion code per enum code
	case Enum_Time_relation_Parentheses:
		res = "parentheses"
	case Enum_Time_relation_Bracket:
		res = "bracket"
	case Enum_Time_relation_Equals:
		res = "equals"
	case Enum_Time_relation_Slash:
		res = "slash"
	case Enum_Time_relation_Space:
		res = "space"
	case Enum_Time_relation_Hyphen:
		res = "hyphen"
	}
	return
}

func (enum_time_relation *Enum_Time_relation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "parentheses":
		*enum_time_relation = Enum_Time_relation_Parentheses
		return
	case "bracket":
		*enum_time_relation = Enum_Time_relation_Bracket
		return
	case "equals":
		*enum_time_relation = Enum_Time_relation_Equals
		return
	case "slash":
		*enum_time_relation = Enum_Time_relation_Slash
		return
	case "space":
		*enum_time_relation = Enum_Time_relation_Space
		return
	case "hyphen":
		*enum_time_relation = Enum_Time_relation_Hyphen
		return
	default:
		return errUnkownEnum
	}
}

func (enum_time_relation *Enum_Time_relation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Time_relation_Parentheses":
		*enum_time_relation = Enum_Time_relation_Parentheses
	case "Enum_Time_relation_Bracket":
		*enum_time_relation = Enum_Time_relation_Bracket
	case "Enum_Time_relation_Equals":
		*enum_time_relation = Enum_Time_relation_Equals
	case "Enum_Time_relation_Slash":
		*enum_time_relation = Enum_Time_relation_Slash
	case "Enum_Time_relation_Space":
		*enum_time_relation = Enum_Time_relation_Space
	case "Enum_Time_relation_Hyphen":
		*enum_time_relation = Enum_Time_relation_Hyphen
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_time_relation *Enum_Time_relation) ToCodeString() (res string) {

	switch *enum_time_relation {
	// insertion code per enum code
	case Enum_Time_relation_Parentheses:
		res = "Enum_Time_relation_Parentheses"
	case Enum_Time_relation_Bracket:
		res = "Enum_Time_relation_Bracket"
	case Enum_Time_relation_Equals:
		res = "Enum_Time_relation_Equals"
	case Enum_Time_relation_Slash:
		res = "Enum_Time_relation_Slash"
	case Enum_Time_relation_Space:
		res = "Enum_Time_relation_Space"
	case Enum_Time_relation_Hyphen:
		res = "Enum_Time_relation_Hyphen"
	}
	return
}

func (enum_time_relation Enum_Time_relation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Time_relation_Parentheses")
	res = append(res, "Enum_Time_relation_Bracket")
	res = append(res, "Enum_Time_relation_Equals")
	res = append(res, "Enum_Time_relation_Slash")
	res = append(res, "Enum_Time_relation_Space")
	res = append(res, "Enum_Time_relation_Hyphen")

	return
}

func (enum_time_relation Enum_Time_relation) CodeValues() (res []string) {

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

// Utility function for Enum_Time_separator
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_time_separator Enum_Time_separator) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_time_separator {
	// insertion code per enum code
	case Enum_Time_separator_None:
		res = "none"
	case Enum_Time_separator_Horizontal:
		res = "horizontal"
	case Enum_Time_separator_Diagonal:
		res = "diagonal"
	case Enum_Time_separator_Vertical:
		res = "vertical"
	case Enum_Time_separator_Adjacent:
		res = "adjacent"
	}
	return
}

func (enum_time_separator *Enum_Time_separator) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_time_separator = Enum_Time_separator_None
		return
	case "horizontal":
		*enum_time_separator = Enum_Time_separator_Horizontal
		return
	case "diagonal":
		*enum_time_separator = Enum_Time_separator_Diagonal
		return
	case "vertical":
		*enum_time_separator = Enum_Time_separator_Vertical
		return
	case "adjacent":
		*enum_time_separator = Enum_Time_separator_Adjacent
		return
	default:
		return errUnkownEnum
	}
}

func (enum_time_separator *Enum_Time_separator) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Time_separator_None":
		*enum_time_separator = Enum_Time_separator_None
	case "Enum_Time_separator_Horizontal":
		*enum_time_separator = Enum_Time_separator_Horizontal
	case "Enum_Time_separator_Diagonal":
		*enum_time_separator = Enum_Time_separator_Diagonal
	case "Enum_Time_separator_Vertical":
		*enum_time_separator = Enum_Time_separator_Vertical
	case "Enum_Time_separator_Adjacent":
		*enum_time_separator = Enum_Time_separator_Adjacent
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_time_separator *Enum_Time_separator) ToCodeString() (res string) {

	switch *enum_time_separator {
	// insertion code per enum code
	case Enum_Time_separator_None:
		res = "Enum_Time_separator_None"
	case Enum_Time_separator_Horizontal:
		res = "Enum_Time_separator_Horizontal"
	case Enum_Time_separator_Diagonal:
		res = "Enum_Time_separator_Diagonal"
	case Enum_Time_separator_Vertical:
		res = "Enum_Time_separator_Vertical"
	case Enum_Time_separator_Adjacent:
		res = "Enum_Time_separator_Adjacent"
	}
	return
}

func (enum_time_separator Enum_Time_separator) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Time_separator_None")
	res = append(res, "Enum_Time_separator_Horizontal")
	res = append(res, "Enum_Time_separator_Diagonal")
	res = append(res, "Enum_Time_separator_Vertical")
	res = append(res, "Enum_Time_separator_Adjacent")

	return
}

func (enum_time_separator Enum_Time_separator) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "horizontal")
	res = append(res, "diagonal")
	res = append(res, "vertical")
	res = append(res, "adjacent")

	return
}

// Utility function for Enum_Time_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_time_symbol Enum_Time_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_time_symbol {
	// insertion code per enum code
	case Enum_Time_symbol_Common:
		res = "common"
	case Enum_Time_symbol_Cut:
		res = "cut"
	case Enum_Time_symbol_Single_number:
		res = "single-number"
	case Enum_Time_symbol_Note:
		res = "note"
	case Enum_Time_symbol_Dotted_note:
		res = "dotted-note"
	case Enum_Time_symbol_Normal:
		res = "normal"
	}
	return
}

func (enum_time_symbol *Enum_Time_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "common":
		*enum_time_symbol = Enum_Time_symbol_Common
		return
	case "cut":
		*enum_time_symbol = Enum_Time_symbol_Cut
		return
	case "single-number":
		*enum_time_symbol = Enum_Time_symbol_Single_number
		return
	case "note":
		*enum_time_symbol = Enum_Time_symbol_Note
		return
	case "dotted-note":
		*enum_time_symbol = Enum_Time_symbol_Dotted_note
		return
	case "normal":
		*enum_time_symbol = Enum_Time_symbol_Normal
		return
	default:
		return errUnkownEnum
	}
}

func (enum_time_symbol *Enum_Time_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Time_symbol_Common":
		*enum_time_symbol = Enum_Time_symbol_Common
	case "Enum_Time_symbol_Cut":
		*enum_time_symbol = Enum_Time_symbol_Cut
	case "Enum_Time_symbol_Single_number":
		*enum_time_symbol = Enum_Time_symbol_Single_number
	case "Enum_Time_symbol_Note":
		*enum_time_symbol = Enum_Time_symbol_Note
	case "Enum_Time_symbol_Dotted_note":
		*enum_time_symbol = Enum_Time_symbol_Dotted_note
	case "Enum_Time_symbol_Normal":
		*enum_time_symbol = Enum_Time_symbol_Normal
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_time_symbol *Enum_Time_symbol) ToCodeString() (res string) {

	switch *enum_time_symbol {
	// insertion code per enum code
	case Enum_Time_symbol_Common:
		res = "Enum_Time_symbol_Common"
	case Enum_Time_symbol_Cut:
		res = "Enum_Time_symbol_Cut"
	case Enum_Time_symbol_Single_number:
		res = "Enum_Time_symbol_Single_number"
	case Enum_Time_symbol_Note:
		res = "Enum_Time_symbol_Note"
	case Enum_Time_symbol_Dotted_note:
		res = "Enum_Time_symbol_Dotted_note"
	case Enum_Time_symbol_Normal:
		res = "Enum_Time_symbol_Normal"
	}
	return
}

func (enum_time_symbol Enum_Time_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Time_symbol_Common")
	res = append(res, "Enum_Time_symbol_Cut")
	res = append(res, "Enum_Time_symbol_Single_number")
	res = append(res, "Enum_Time_symbol_Note")
	res = append(res, "Enum_Time_symbol_Dotted_note")
	res = append(res, "Enum_Time_symbol_Normal")

	return
}

func (enum_time_symbol Enum_Time_symbol) CodeValues() (res []string) {

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

// Utility function for Enum_Tip_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_tip_direction Enum_Tip_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_tip_direction {
	// insertion code per enum code
	case Enum_Tip_direction_Up:
		res = "up"
	case Enum_Tip_direction_Down:
		res = "down"
	case Enum_Tip_direction_Left:
		res = "left"
	case Enum_Tip_direction_Right:
		res = "right"
	case Enum_Tip_direction_Northwest:
		res = "northwest"
	case Enum_Tip_direction_Northeast:
		res = "northeast"
	case Enum_Tip_direction_Southeast:
		res = "southeast"
	case Enum_Tip_direction_Southwest:
		res = "southwest"
	}
	return
}

func (enum_tip_direction *Enum_Tip_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enum_tip_direction = Enum_Tip_direction_Up
		return
	case "down":
		*enum_tip_direction = Enum_Tip_direction_Down
		return
	case "left":
		*enum_tip_direction = Enum_Tip_direction_Left
		return
	case "right":
		*enum_tip_direction = Enum_Tip_direction_Right
		return
	case "northwest":
		*enum_tip_direction = Enum_Tip_direction_Northwest
		return
	case "northeast":
		*enum_tip_direction = Enum_Tip_direction_Northeast
		return
	case "southeast":
		*enum_tip_direction = Enum_Tip_direction_Southeast
		return
	case "southwest":
		*enum_tip_direction = Enum_Tip_direction_Southwest
		return
	default:
		return errUnkownEnum
	}
}

func (enum_tip_direction *Enum_Tip_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Tip_direction_Up":
		*enum_tip_direction = Enum_Tip_direction_Up
	case "Enum_Tip_direction_Down":
		*enum_tip_direction = Enum_Tip_direction_Down
	case "Enum_Tip_direction_Left":
		*enum_tip_direction = Enum_Tip_direction_Left
	case "Enum_Tip_direction_Right":
		*enum_tip_direction = Enum_Tip_direction_Right
	case "Enum_Tip_direction_Northwest":
		*enum_tip_direction = Enum_Tip_direction_Northwest
	case "Enum_Tip_direction_Northeast":
		*enum_tip_direction = Enum_Tip_direction_Northeast
	case "Enum_Tip_direction_Southeast":
		*enum_tip_direction = Enum_Tip_direction_Southeast
	case "Enum_Tip_direction_Southwest":
		*enum_tip_direction = Enum_Tip_direction_Southwest
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_tip_direction *Enum_Tip_direction) ToCodeString() (res string) {

	switch *enum_tip_direction {
	// insertion code per enum code
	case Enum_Tip_direction_Up:
		res = "Enum_Tip_direction_Up"
	case Enum_Tip_direction_Down:
		res = "Enum_Tip_direction_Down"
	case Enum_Tip_direction_Left:
		res = "Enum_Tip_direction_Left"
	case Enum_Tip_direction_Right:
		res = "Enum_Tip_direction_Right"
	case Enum_Tip_direction_Northwest:
		res = "Enum_Tip_direction_Northwest"
	case Enum_Tip_direction_Northeast:
		res = "Enum_Tip_direction_Northeast"
	case Enum_Tip_direction_Southeast:
		res = "Enum_Tip_direction_Southeast"
	case Enum_Tip_direction_Southwest:
		res = "Enum_Tip_direction_Southwest"
	}
	return
}

func (enum_tip_direction Enum_Tip_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Tip_direction_Up")
	res = append(res, "Enum_Tip_direction_Down")
	res = append(res, "Enum_Tip_direction_Left")
	res = append(res, "Enum_Tip_direction_Right")
	res = append(res, "Enum_Tip_direction_Northwest")
	res = append(res, "Enum_Tip_direction_Northeast")
	res = append(res, "Enum_Tip_direction_Southeast")
	res = append(res, "Enum_Tip_direction_Southwest")

	return
}

func (enum_tip_direction Enum_Tip_direction) CodeValues() (res []string) {

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

// Utility function for Enum_Top_bottom
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_top_bottom Enum_Top_bottom) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_top_bottom {
	// insertion code per enum code
	case Enum_Top_bottom_Top:
		res = "top"
	case Enum_Top_bottom_Bottom:
		res = "bottom"
	}
	return
}

func (enum_top_bottom *Enum_Top_bottom) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enum_top_bottom = Enum_Top_bottom_Top
		return
	case "bottom":
		*enum_top_bottom = Enum_Top_bottom_Bottom
		return
	default:
		return errUnkownEnum
	}
}

func (enum_top_bottom *Enum_Top_bottom) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Top_bottom_Top":
		*enum_top_bottom = Enum_Top_bottom_Top
	case "Enum_Top_bottom_Bottom":
		*enum_top_bottom = Enum_Top_bottom_Bottom
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_top_bottom *Enum_Top_bottom) ToCodeString() (res string) {

	switch *enum_top_bottom {
	// insertion code per enum code
	case Enum_Top_bottom_Top:
		res = "Enum_Top_bottom_Top"
	case Enum_Top_bottom_Bottom:
		res = "Enum_Top_bottom_Bottom"
	}
	return
}

func (enum_top_bottom Enum_Top_bottom) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Top_bottom_Top")
	res = append(res, "Enum_Top_bottom_Bottom")

	return
}

func (enum_top_bottom Enum_Top_bottom) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "bottom")

	return
}

// Utility function for Enum_Tremolo_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_tremolo_type Enum_Tremolo_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_tremolo_type {
	// insertion code per enum code
	case Enum_Tremolo_type_Start:
		res = "start"
	case Enum_Tremolo_type_Stop:
		res = "stop"
	case Enum_Tremolo_type_Single:
		res = "single"
	case Enum_Tremolo_type_Unmeasured:
		res = "unmeasured"
	}
	return
}

func (enum_tremolo_type *Enum_Tremolo_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*enum_tremolo_type = Enum_Tremolo_type_Start
		return
	case "stop":
		*enum_tremolo_type = Enum_Tremolo_type_Stop
		return
	case "single":
		*enum_tremolo_type = Enum_Tremolo_type_Single
		return
	case "unmeasured":
		*enum_tremolo_type = Enum_Tremolo_type_Unmeasured
		return
	default:
		return errUnkownEnum
	}
}

func (enum_tremolo_type *Enum_Tremolo_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Tremolo_type_Start":
		*enum_tremolo_type = Enum_Tremolo_type_Start
	case "Enum_Tremolo_type_Stop":
		*enum_tremolo_type = Enum_Tremolo_type_Stop
	case "Enum_Tremolo_type_Single":
		*enum_tremolo_type = Enum_Tremolo_type_Single
	case "Enum_Tremolo_type_Unmeasured":
		*enum_tremolo_type = Enum_Tremolo_type_Unmeasured
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_tremolo_type *Enum_Tremolo_type) ToCodeString() (res string) {

	switch *enum_tremolo_type {
	// insertion code per enum code
	case Enum_Tremolo_type_Start:
		res = "Enum_Tremolo_type_Start"
	case Enum_Tremolo_type_Stop:
		res = "Enum_Tremolo_type_Stop"
	case Enum_Tremolo_type_Single:
		res = "Enum_Tremolo_type_Single"
	case Enum_Tremolo_type_Unmeasured:
		res = "Enum_Tremolo_type_Unmeasured"
	}
	return
}

func (enum_tremolo_type Enum_Tremolo_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Tremolo_type_Start")
	res = append(res, "Enum_Tremolo_type_Stop")
	res = append(res, "Enum_Tremolo_type_Single")
	res = append(res, "Enum_Tremolo_type_Unmeasured")

	return
}

func (enum_tremolo_type Enum_Tremolo_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")
	res = append(res, "unmeasured")

	return
}

// Utility function for Enum_Trill_step
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_trill_step Enum_Trill_step) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_trill_step {
	// insertion code per enum code
	case Enum_Trill_step_Whole:
		res = "whole"
	case Enum_Trill_step_Half:
		res = "half"
	case Enum_Trill_step_Unison:
		res = "unison"
	}
	return
}

func (enum_trill_step *Enum_Trill_step) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*enum_trill_step = Enum_Trill_step_Whole
		return
	case "half":
		*enum_trill_step = Enum_Trill_step_Half
		return
	case "unison":
		*enum_trill_step = Enum_Trill_step_Unison
		return
	default:
		return errUnkownEnum
	}
}

func (enum_trill_step *Enum_Trill_step) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Trill_step_Whole":
		*enum_trill_step = Enum_Trill_step_Whole
	case "Enum_Trill_step_Half":
		*enum_trill_step = Enum_Trill_step_Half
	case "Enum_Trill_step_Unison":
		*enum_trill_step = Enum_Trill_step_Unison
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_trill_step *Enum_Trill_step) ToCodeString() (res string) {

	switch *enum_trill_step {
	// insertion code per enum code
	case Enum_Trill_step_Whole:
		res = "Enum_Trill_step_Whole"
	case Enum_Trill_step_Half:
		res = "Enum_Trill_step_Half"
	case Enum_Trill_step_Unison:
		res = "Enum_Trill_step_Unison"
	}
	return
}

func (enum_trill_step Enum_Trill_step) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Trill_step_Whole")
	res = append(res, "Enum_Trill_step_Half")
	res = append(res, "Enum_Trill_step_Unison")

	return
}

func (enum_trill_step Enum_Trill_step) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "unison")

	return
}

// Utility function for Enum_Two_note_turn
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_two_note_turn Enum_Two_note_turn) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_two_note_turn {
	// insertion code per enum code
	case Enum_Two_note_turn_Whole:
		res = "whole"
	case Enum_Two_note_turn_Half:
		res = "half"
	case Enum_Two_note_turn_None:
		res = "none"
	}
	return
}

func (enum_two_note_turn *Enum_Two_note_turn) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*enum_two_note_turn = Enum_Two_note_turn_Whole
		return
	case "half":
		*enum_two_note_turn = Enum_Two_note_turn_Half
		return
	case "none":
		*enum_two_note_turn = Enum_Two_note_turn_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_two_note_turn *Enum_Two_note_turn) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Two_note_turn_Whole":
		*enum_two_note_turn = Enum_Two_note_turn_Whole
	case "Enum_Two_note_turn_Half":
		*enum_two_note_turn = Enum_Two_note_turn_Half
	case "Enum_Two_note_turn_None":
		*enum_two_note_turn = Enum_Two_note_turn_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_two_note_turn *Enum_Two_note_turn) ToCodeString() (res string) {

	switch *enum_two_note_turn {
	// insertion code per enum code
	case Enum_Two_note_turn_Whole:
		res = "Enum_Two_note_turn_Whole"
	case Enum_Two_note_turn_Half:
		res = "Enum_Two_note_turn_Half"
	case Enum_Two_note_turn_None:
		res = "Enum_Two_note_turn_None"
	}
	return
}

func (enum_two_note_turn Enum_Two_note_turn) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Two_note_turn_Whole")
	res = append(res, "Enum_Two_note_turn_Half")
	res = append(res, "Enum_Two_note_turn_None")

	return
}

func (enum_two_note_turn Enum_Two_note_turn) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "none")

	return
}

// Utility function for Enum_Up_down
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_up_down Enum_Up_down) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_up_down {
	// insertion code per enum code
	case Enum_Up_down_Up:
		res = "up"
	case Enum_Up_down_Down:
		res = "down"
	}
	return
}

func (enum_up_down *Enum_Up_down) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enum_up_down = Enum_Up_down_Up
		return
	case "down":
		*enum_up_down = Enum_Up_down_Down
		return
	default:
		return errUnkownEnum
	}
}

func (enum_up_down *Enum_Up_down) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Up_down_Up":
		*enum_up_down = Enum_Up_down_Up
	case "Enum_Up_down_Down":
		*enum_up_down = Enum_Up_down_Down
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_up_down *Enum_Up_down) ToCodeString() (res string) {

	switch *enum_up_down {
	// insertion code per enum code
	case Enum_Up_down_Up:
		res = "Enum_Up_down_Up"
	case Enum_Up_down_Down:
		res = "Enum_Up_down_Down"
	}
	return
}

func (enum_up_down Enum_Up_down) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Up_down_Up")
	res = append(res, "Enum_Up_down_Down")

	return
}

func (enum_up_down Enum_Up_down) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")

	return
}

// Utility function for Enum_Up_down_stop_continue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_up_down_stop_continue Enum_Up_down_stop_continue) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_up_down_stop_continue {
	// insertion code per enum code
	case Enum_Up_down_stop_continue_Up:
		res = "up"
	case Enum_Up_down_stop_continue_Down:
		res = "down"
	case Enum_Up_down_stop_continue_Stop:
		res = "stop"
	case Enum_Up_down_stop_continue_Continue:
		res = "continue"
	}
	return
}

func (enum_up_down_stop_continue *Enum_Up_down_stop_continue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Up
		return
	case "down":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Down
		return
	case "stop":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Stop
		return
	case "continue":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Continue
		return
	default:
		return errUnkownEnum
	}
}

func (enum_up_down_stop_continue *Enum_Up_down_stop_continue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Up_down_stop_continue_Up":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Up
	case "Enum_Up_down_stop_continue_Down":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Down
	case "Enum_Up_down_stop_continue_Stop":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Stop
	case "Enum_Up_down_stop_continue_Continue":
		*enum_up_down_stop_continue = Enum_Up_down_stop_continue_Continue
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_up_down_stop_continue *Enum_Up_down_stop_continue) ToCodeString() (res string) {

	switch *enum_up_down_stop_continue {
	// insertion code per enum code
	case Enum_Up_down_stop_continue_Up:
		res = "Enum_Up_down_stop_continue_Up"
	case Enum_Up_down_stop_continue_Down:
		res = "Enum_Up_down_stop_continue_Down"
	case Enum_Up_down_stop_continue_Stop:
		res = "Enum_Up_down_stop_continue_Stop"
	case Enum_Up_down_stop_continue_Continue:
		res = "Enum_Up_down_stop_continue_Continue"
	}
	return
}

func (enum_up_down_stop_continue Enum_Up_down_stop_continue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Up_down_stop_continue_Up")
	res = append(res, "Enum_Up_down_stop_continue_Down")
	res = append(res, "Enum_Up_down_stop_continue_Stop")
	res = append(res, "Enum_Up_down_stop_continue_Continue")

	return
}

func (enum_up_down_stop_continue Enum_Up_down_stop_continue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Enum_Upright_inverted
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_upright_inverted Enum_Upright_inverted) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_upright_inverted {
	// insertion code per enum code
	case Enum_Upright_inverted_Upright:
		res = "upright"
	case Enum_Upright_inverted_Inverted:
		res = "inverted"
	}
	return
}

func (enum_upright_inverted *Enum_Upright_inverted) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upright":
		*enum_upright_inverted = Enum_Upright_inverted_Upright
		return
	case "inverted":
		*enum_upright_inverted = Enum_Upright_inverted_Inverted
		return
	default:
		return errUnkownEnum
	}
}

func (enum_upright_inverted *Enum_Upright_inverted) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Upright_inverted_Upright":
		*enum_upright_inverted = Enum_Upright_inverted_Upright
	case "Enum_Upright_inverted_Inverted":
		*enum_upright_inverted = Enum_Upright_inverted_Inverted
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_upright_inverted *Enum_Upright_inverted) ToCodeString() (res string) {

	switch *enum_upright_inverted {
	// insertion code per enum code
	case Enum_Upright_inverted_Upright:
		res = "Enum_Upright_inverted_Upright"
	case Enum_Upright_inverted_Inverted:
		res = "Enum_Upright_inverted_Inverted"
	}
	return
}

func (enum_upright_inverted Enum_Upright_inverted) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Upright_inverted_Upright")
	res = append(res, "Enum_Upright_inverted_Inverted")

	return
}

func (enum_upright_inverted Enum_Upright_inverted) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upright")
	res = append(res, "inverted")

	return
}

// Utility function for Enum_Valign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_valign Enum_Valign) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_valign {
	// insertion code per enum code
	case Enum_Valign_Top:
		res = "top"
	case Enum_Valign_Middle:
		res = "middle"
	case Enum_Valign_Bottom:
		res = "bottom"
	case Enum_Valign_Baseline:
		res = "baseline"
	}
	return
}

func (enum_valign *Enum_Valign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enum_valign = Enum_Valign_Top
		return
	case "middle":
		*enum_valign = Enum_Valign_Middle
		return
	case "bottom":
		*enum_valign = Enum_Valign_Bottom
		return
	case "baseline":
		*enum_valign = Enum_Valign_Baseline
		return
	default:
		return errUnkownEnum
	}
}

func (enum_valign *Enum_Valign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Valign_Top":
		*enum_valign = Enum_Valign_Top
	case "Enum_Valign_Middle":
		*enum_valign = Enum_Valign_Middle
	case "Enum_Valign_Bottom":
		*enum_valign = Enum_Valign_Bottom
	case "Enum_Valign_Baseline":
		*enum_valign = Enum_Valign_Baseline
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_valign *Enum_Valign) ToCodeString() (res string) {

	switch *enum_valign {
	// insertion code per enum code
	case Enum_Valign_Top:
		res = "Enum_Valign_Top"
	case Enum_Valign_Middle:
		res = "Enum_Valign_Middle"
	case Enum_Valign_Bottom:
		res = "Enum_Valign_Bottom"
	case Enum_Valign_Baseline:
		res = "Enum_Valign_Baseline"
	}
	return
}

func (enum_valign Enum_Valign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Valign_Top")
	res = append(res, "Enum_Valign_Middle")
	res = append(res, "Enum_Valign_Bottom")
	res = append(res, "Enum_Valign_Baseline")

	return
}

func (enum_valign Enum_Valign) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")
	res = append(res, "baseline")

	return
}

// Utility function for Enum_Valign_image
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_valign_image Enum_Valign_image) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_valign_image {
	// insertion code per enum code
	case Enum_Valign_image_Top:
		res = "top"
	case Enum_Valign_image_Middle:
		res = "middle"
	case Enum_Valign_image_Bottom:
		res = "bottom"
	}
	return
}

func (enum_valign_image *Enum_Valign_image) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*enum_valign_image = Enum_Valign_image_Top
		return
	case "middle":
		*enum_valign_image = Enum_Valign_image_Middle
		return
	case "bottom":
		*enum_valign_image = Enum_Valign_image_Bottom
		return
	default:
		return errUnkownEnum
	}
}

func (enum_valign_image *Enum_Valign_image) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Valign_image_Top":
		*enum_valign_image = Enum_Valign_image_Top
	case "Enum_Valign_image_Middle":
		*enum_valign_image = Enum_Valign_image_Middle
	case "Enum_Valign_image_Bottom":
		*enum_valign_image = Enum_Valign_image_Bottom
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_valign_image *Enum_Valign_image) ToCodeString() (res string) {

	switch *enum_valign_image {
	// insertion code per enum code
	case Enum_Valign_image_Top:
		res = "Enum_Valign_image_Top"
	case Enum_Valign_image_Middle:
		res = "Enum_Valign_image_Middle"
	case Enum_Valign_image_Bottom:
		res = "Enum_Valign_image_Bottom"
	}
	return
}

func (enum_valign_image Enum_Valign_image) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Valign_image_Top")
	res = append(res, "Enum_Valign_image_Middle")
	res = append(res, "Enum_Valign_image_Bottom")

	return
}

func (enum_valign_image Enum_Valign_image) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")

	return
}

// Utility function for Enum_Wedge_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_wedge_type Enum_Wedge_type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_wedge_type {
	// insertion code per enum code
	case Enum_Wedge_type_Crescendo:
		res = "crescendo"
	case Enum_Wedge_type_Diminuendo:
		res = "diminuendo"
	case Enum_Wedge_type_Stop:
		res = "stop"
	case Enum_Wedge_type_Continue:
		res = "continue"
	}
	return
}

func (enum_wedge_type *Enum_Wedge_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "crescendo":
		*enum_wedge_type = Enum_Wedge_type_Crescendo
		return
	case "diminuendo":
		*enum_wedge_type = Enum_Wedge_type_Diminuendo
		return
	case "stop":
		*enum_wedge_type = Enum_Wedge_type_Stop
		return
	case "continue":
		*enum_wedge_type = Enum_Wedge_type_Continue
		return
	default:
		return errUnkownEnum
	}
}

func (enum_wedge_type *Enum_Wedge_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Wedge_type_Crescendo":
		*enum_wedge_type = Enum_Wedge_type_Crescendo
	case "Enum_Wedge_type_Diminuendo":
		*enum_wedge_type = Enum_Wedge_type_Diminuendo
	case "Enum_Wedge_type_Stop":
		*enum_wedge_type = Enum_Wedge_type_Stop
	case "Enum_Wedge_type_Continue":
		*enum_wedge_type = Enum_Wedge_type_Continue
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_wedge_type *Enum_Wedge_type) ToCodeString() (res string) {

	switch *enum_wedge_type {
	// insertion code per enum code
	case Enum_Wedge_type_Crescendo:
		res = "Enum_Wedge_type_Crescendo"
	case Enum_Wedge_type_Diminuendo:
		res = "Enum_Wedge_type_Diminuendo"
	case Enum_Wedge_type_Stop:
		res = "Enum_Wedge_type_Stop"
	case Enum_Wedge_type_Continue:
		res = "Enum_Wedge_type_Continue"
	}
	return
}

func (enum_wedge_type Enum_Wedge_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Wedge_type_Crescendo")
	res = append(res, "Enum_Wedge_type_Diminuendo")
	res = append(res, "Enum_Wedge_type_Stop")
	res = append(res, "Enum_Wedge_type_Continue")

	return
}

func (enum_wedge_type Enum_Wedge_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "crescendo")
	res = append(res, "diminuendo")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Enum_Winged
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_winged Enum_Winged) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_winged {
	// insertion code per enum code
	case Enum_Winged_None:
		res = "none"
	case Enum_Winged_Straight:
		res = "straight"
	case Enum_Winged_Curved:
		res = "curved"
	case Enum_Winged_Double_straight:
		res = "double-straight"
	case Enum_Winged_Double_curved:
		res = "double-curved"
	}
	return
}

func (enum_winged *Enum_Winged) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_winged = Enum_Winged_None
		return
	case "straight":
		*enum_winged = Enum_Winged_Straight
		return
	case "curved":
		*enum_winged = Enum_Winged_Curved
		return
	case "double-straight":
		*enum_winged = Enum_Winged_Double_straight
		return
	case "double-curved":
		*enum_winged = Enum_Winged_Double_curved
		return
	default:
		return errUnkownEnum
	}
}

func (enum_winged *Enum_Winged) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Winged_None":
		*enum_winged = Enum_Winged_None
	case "Enum_Winged_Straight":
		*enum_winged = Enum_Winged_Straight
	case "Enum_Winged_Curved":
		*enum_winged = Enum_Winged_Curved
	case "Enum_Winged_Double_straight":
		*enum_winged = Enum_Winged_Double_straight
	case "Enum_Winged_Double_curved":
		*enum_winged = Enum_Winged_Double_curved
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_winged *Enum_Winged) ToCodeString() (res string) {

	switch *enum_winged {
	// insertion code per enum code
	case Enum_Winged_None:
		res = "Enum_Winged_None"
	case Enum_Winged_Straight:
		res = "Enum_Winged_Straight"
	case Enum_Winged_Curved:
		res = "Enum_Winged_Curved"
	case Enum_Winged_Double_straight:
		res = "Enum_Winged_Double_straight"
	case Enum_Winged_Double_curved:
		res = "Enum_Winged_Double_curved"
	}
	return
}

func (enum_winged Enum_Winged) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Winged_None")
	res = append(res, "Enum_Winged_Straight")
	res = append(res, "Enum_Winged_Curved")
	res = append(res, "Enum_Winged_Double_straight")
	res = append(res, "Enum_Winged_Double_curved")

	return
}

func (enum_winged Enum_Winged) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "straight")
	res = append(res, "curved")
	res = append(res, "double-straight")
	res = append(res, "double-curved")

	return
}

// Utility function for Enum_Wood_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_wood_value Enum_Wood_value) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_wood_value {
	// insertion code per enum code
	case Enum_Wood_value_Bamboo_scraper:
		res = "bamboo scraper"
	case Enum_Wood_value_Board_clapper:
		res = "board clapper"
	case Enum_Wood_value_Cabasa:
		res = "cabasa"
	case Enum_Wood_value_Castanets:
		res = "castanets"
	case Enum_Wood_value_Castanets_with_handle:
		res = "castanets with handle"
	case Enum_Wood_value_Claves:
		res = "claves"
	case Enum_Wood_value_Football_rattle:
		res = "football rattle"
	case Enum_Wood_value_Guiro:
		res = "guiro"
	case Enum_Wood_value_Log_drum:
		res = "log drum"
	case Enum_Wood_value_Maraca:
		res = "maraca"
	case Enum_Wood_value_Maracas:
		res = "maracas"
	case Enum_Wood_value_Quijada:
		res = "quijada"
	case Enum_Wood_value_Rainstick:
		res = "rainstick"
	case Enum_Wood_value_Ratchet:
		res = "ratchet"
	case Enum_Wood_value_Reco_reco:
		res = "reco-reco"
	case Enum_Wood_value_Sandpaper_blocks:
		res = "sandpaper blocks"
	case Enum_Wood_value_Slit_drum:
		res = "slit drum"
	case Enum_Wood_value_Temple_block:
		res = "temple block"
	case Enum_Wood_value_Vibraslap:
		res = "vibraslap"
	case Enum_Wood_value_Whip:
		res = "whip"
	case Enum_Wood_value_Wood_block:
		res = "wood block"
	}
	return
}

func (enum_wood_value *Enum_Wood_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bamboo scraper":
		*enum_wood_value = Enum_Wood_value_Bamboo_scraper
		return
	case "board clapper":
		*enum_wood_value = Enum_Wood_value_Board_clapper
		return
	case "cabasa":
		*enum_wood_value = Enum_Wood_value_Cabasa
		return
	case "castanets":
		*enum_wood_value = Enum_Wood_value_Castanets
		return
	case "castanets with handle":
		*enum_wood_value = Enum_Wood_value_Castanets_with_handle
		return
	case "claves":
		*enum_wood_value = Enum_Wood_value_Claves
		return
	case "football rattle":
		*enum_wood_value = Enum_Wood_value_Football_rattle
		return
	case "guiro":
		*enum_wood_value = Enum_Wood_value_Guiro
		return
	case "log drum":
		*enum_wood_value = Enum_Wood_value_Log_drum
		return
	case "maraca":
		*enum_wood_value = Enum_Wood_value_Maraca
		return
	case "maracas":
		*enum_wood_value = Enum_Wood_value_Maracas
		return
	case "quijada":
		*enum_wood_value = Enum_Wood_value_Quijada
		return
	case "rainstick":
		*enum_wood_value = Enum_Wood_value_Rainstick
		return
	case "ratchet":
		*enum_wood_value = Enum_Wood_value_Ratchet
		return
	case "reco-reco":
		*enum_wood_value = Enum_Wood_value_Reco_reco
		return
	case "sandpaper blocks":
		*enum_wood_value = Enum_Wood_value_Sandpaper_blocks
		return
	case "slit drum":
		*enum_wood_value = Enum_Wood_value_Slit_drum
		return
	case "temple block":
		*enum_wood_value = Enum_Wood_value_Temple_block
		return
	case "vibraslap":
		*enum_wood_value = Enum_Wood_value_Vibraslap
		return
	case "whip":
		*enum_wood_value = Enum_Wood_value_Whip
		return
	case "wood block":
		*enum_wood_value = Enum_Wood_value_Wood_block
		return
	default:
		return errUnkownEnum
	}
}

func (enum_wood_value *Enum_Wood_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Wood_value_Bamboo_scraper":
		*enum_wood_value = Enum_Wood_value_Bamboo_scraper
	case "Enum_Wood_value_Board_clapper":
		*enum_wood_value = Enum_Wood_value_Board_clapper
	case "Enum_Wood_value_Cabasa":
		*enum_wood_value = Enum_Wood_value_Cabasa
	case "Enum_Wood_value_Castanets":
		*enum_wood_value = Enum_Wood_value_Castanets
	case "Enum_Wood_value_Castanets_with_handle":
		*enum_wood_value = Enum_Wood_value_Castanets_with_handle
	case "Enum_Wood_value_Claves":
		*enum_wood_value = Enum_Wood_value_Claves
	case "Enum_Wood_value_Football_rattle":
		*enum_wood_value = Enum_Wood_value_Football_rattle
	case "Enum_Wood_value_Guiro":
		*enum_wood_value = Enum_Wood_value_Guiro
	case "Enum_Wood_value_Log_drum":
		*enum_wood_value = Enum_Wood_value_Log_drum
	case "Enum_Wood_value_Maraca":
		*enum_wood_value = Enum_Wood_value_Maraca
	case "Enum_Wood_value_Maracas":
		*enum_wood_value = Enum_Wood_value_Maracas
	case "Enum_Wood_value_Quijada":
		*enum_wood_value = Enum_Wood_value_Quijada
	case "Enum_Wood_value_Rainstick":
		*enum_wood_value = Enum_Wood_value_Rainstick
	case "Enum_Wood_value_Ratchet":
		*enum_wood_value = Enum_Wood_value_Ratchet
	case "Enum_Wood_value_Reco_reco":
		*enum_wood_value = Enum_Wood_value_Reco_reco
	case "Enum_Wood_value_Sandpaper_blocks":
		*enum_wood_value = Enum_Wood_value_Sandpaper_blocks
	case "Enum_Wood_value_Slit_drum":
		*enum_wood_value = Enum_Wood_value_Slit_drum
	case "Enum_Wood_value_Temple_block":
		*enum_wood_value = Enum_Wood_value_Temple_block
	case "Enum_Wood_value_Vibraslap":
		*enum_wood_value = Enum_Wood_value_Vibraslap
	case "Enum_Wood_value_Whip":
		*enum_wood_value = Enum_Wood_value_Whip
	case "Enum_Wood_value_Wood_block":
		*enum_wood_value = Enum_Wood_value_Wood_block
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_wood_value *Enum_Wood_value) ToCodeString() (res string) {

	switch *enum_wood_value {
	// insertion code per enum code
	case Enum_Wood_value_Bamboo_scraper:
		res = "Enum_Wood_value_Bamboo_scraper"
	case Enum_Wood_value_Board_clapper:
		res = "Enum_Wood_value_Board_clapper"
	case Enum_Wood_value_Cabasa:
		res = "Enum_Wood_value_Cabasa"
	case Enum_Wood_value_Castanets:
		res = "Enum_Wood_value_Castanets"
	case Enum_Wood_value_Castanets_with_handle:
		res = "Enum_Wood_value_Castanets_with_handle"
	case Enum_Wood_value_Claves:
		res = "Enum_Wood_value_Claves"
	case Enum_Wood_value_Football_rattle:
		res = "Enum_Wood_value_Football_rattle"
	case Enum_Wood_value_Guiro:
		res = "Enum_Wood_value_Guiro"
	case Enum_Wood_value_Log_drum:
		res = "Enum_Wood_value_Log_drum"
	case Enum_Wood_value_Maraca:
		res = "Enum_Wood_value_Maraca"
	case Enum_Wood_value_Maracas:
		res = "Enum_Wood_value_Maracas"
	case Enum_Wood_value_Quijada:
		res = "Enum_Wood_value_Quijada"
	case Enum_Wood_value_Rainstick:
		res = "Enum_Wood_value_Rainstick"
	case Enum_Wood_value_Ratchet:
		res = "Enum_Wood_value_Ratchet"
	case Enum_Wood_value_Reco_reco:
		res = "Enum_Wood_value_Reco_reco"
	case Enum_Wood_value_Sandpaper_blocks:
		res = "Enum_Wood_value_Sandpaper_blocks"
	case Enum_Wood_value_Slit_drum:
		res = "Enum_Wood_value_Slit_drum"
	case Enum_Wood_value_Temple_block:
		res = "Enum_Wood_value_Temple_block"
	case Enum_Wood_value_Vibraslap:
		res = "Enum_Wood_value_Vibraslap"
	case Enum_Wood_value_Whip:
		res = "Enum_Wood_value_Whip"
	case Enum_Wood_value_Wood_block:
		res = "Enum_Wood_value_Wood_block"
	}
	return
}

func (enum_wood_value Enum_Wood_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Wood_value_Bamboo_scraper")
	res = append(res, "Enum_Wood_value_Board_clapper")
	res = append(res, "Enum_Wood_value_Cabasa")
	res = append(res, "Enum_Wood_value_Castanets")
	res = append(res, "Enum_Wood_value_Castanets_with_handle")
	res = append(res, "Enum_Wood_value_Claves")
	res = append(res, "Enum_Wood_value_Football_rattle")
	res = append(res, "Enum_Wood_value_Guiro")
	res = append(res, "Enum_Wood_value_Log_drum")
	res = append(res, "Enum_Wood_value_Maraca")
	res = append(res, "Enum_Wood_value_Maracas")
	res = append(res, "Enum_Wood_value_Quijada")
	res = append(res, "Enum_Wood_value_Rainstick")
	res = append(res, "Enum_Wood_value_Ratchet")
	res = append(res, "Enum_Wood_value_Reco_reco")
	res = append(res, "Enum_Wood_value_Sandpaper_blocks")
	res = append(res, "Enum_Wood_value_Slit_drum")
	res = append(res, "Enum_Wood_value_Temple_block")
	res = append(res, "Enum_Wood_value_Vibraslap")
	res = append(res, "Enum_Wood_value_Whip")
	res = append(res, "Enum_Wood_value_Wood_block")

	return
}

func (enum_wood_value Enum_Wood_value) CodeValues() (res []string) {

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

// Utility function for Enum_Yes_no
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_yes_no Enum_Yes_no) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_yes_no {
	// insertion code per enum code
	case Enum_Yes_no_Yes:
		res = "yes"
	case Enum_Yes_no_No:
		res = "no"
	}
	return
}

func (enum_yes_no *Enum_Yes_no) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*enum_yes_no = Enum_Yes_no_Yes
		return
	case "no":
		*enum_yes_no = Enum_Yes_no_No
		return
	default:
		return errUnkownEnum
	}
}

func (enum_yes_no *Enum_Yes_no) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_Yes_no_Yes":
		*enum_yes_no = Enum_Yes_no_Yes
	case "Enum_Yes_no_No":
		*enum_yes_no = Enum_Yes_no_No
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_yes_no *Enum_Yes_no) ToCodeString() (res string) {

	switch *enum_yes_no {
	// insertion code per enum code
	case Enum_Yes_no_Yes:
		res = "Enum_Yes_no_Yes"
	case Enum_Yes_no_No:
		res = "Enum_Yes_no_No"
	}
	return
}

func (enum_yes_no Enum_Yes_no) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_Yes_no_Yes")
	res = append(res, "Enum_Yes_no_No")

	return
}

func (enum_yes_no Enum_Yes_no) CodeValues() (res []string) {

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

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

// Last line of the template
