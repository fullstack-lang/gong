// generated code - do not edit
package aixmdatatypes

// insertion point of enum utility functions
// Utility function for Enum_CodeAircraftBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftbasetype_member_1 Enum_CodeAircraftBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftBaseType_Member_1_LANDPLANE:
		res = "LANDPLANE"
	case Enum_CodeAircraftBaseType_Member_1_SEAPLANE:
		res = "SEAPLANE"
	case Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN:
		res = "AMPHIBIAN"
	case Enum_CodeAircraftBaseType_Member_1_HELICOPTER:
		res = "HELICOPTER"
	case Enum_CodeAircraftBaseType_Member_1_GYROCOPTER:
		res = "GYROCOPTER"
	case Enum_CodeAircraftBaseType_Member_1_TILT_WING:
		res = "TILT_WING"
	case Enum_CodeAircraftBaseType_Member_1_STOL:
		res = "STOL"
	case Enum_CodeAircraftBaseType_Member_1_GLIDER:
		res = "GLIDER"
	case Enum_CodeAircraftBaseType_Member_1_HANGGLIDER:
		res = "HANGGLIDER"
	case Enum_CodeAircraftBaseType_Member_1_PARAGLIDER:
		res = "PARAGLIDER"
	case Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT:
		res = "ULTRA_LIGHT"
	case Enum_CodeAircraftBaseType_Member_1_BALLOON:
		res = "BALLOON"
	case Enum_CodeAircraftBaseType_Member_1_UAV:
		res = "UAV"
	case Enum_CodeAircraftBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeaircraftbasetype_member_1 *Enum_CodeAircraftBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LANDPLANE":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_LANDPLANE
		return
	case "SEAPLANE":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_SEAPLANE
		return
	case "AMPHIBIAN":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN
		return
	case "HELICOPTER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_HELICOPTER
		return
	case "GYROCOPTER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_GYROCOPTER
		return
	case "TILT_WING":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_TILT_WING
		return
	case "STOL":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_STOL
		return
	case "GLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_GLIDER
		return
	case "HANGGLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_HANGGLIDER
		return
	case "PARAGLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_PARAGLIDER
		return
	case "ULTRA_LIGHT":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT
		return
	case "BALLOON":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_BALLOON
		return
	case "UAV":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_UAV
		return
	case "ALL":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftbasetype_member_1 *Enum_CodeAircraftBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftBaseType_Member_1_LANDPLANE":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_LANDPLANE
	case "Enum_CodeAircraftBaseType_Member_1_SEAPLANE":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_SEAPLANE
	case "Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN
	case "Enum_CodeAircraftBaseType_Member_1_HELICOPTER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_HELICOPTER
	case "Enum_CodeAircraftBaseType_Member_1_GYROCOPTER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_GYROCOPTER
	case "Enum_CodeAircraftBaseType_Member_1_TILT_WING":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_TILT_WING
	case "Enum_CodeAircraftBaseType_Member_1_STOL":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_STOL
	case "Enum_CodeAircraftBaseType_Member_1_GLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_GLIDER
	case "Enum_CodeAircraftBaseType_Member_1_HANGGLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_HANGGLIDER
	case "Enum_CodeAircraftBaseType_Member_1_PARAGLIDER":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_PARAGLIDER
	case "Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT
	case "Enum_CodeAircraftBaseType_Member_1_BALLOON":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_BALLOON
	case "Enum_CodeAircraftBaseType_Member_1_UAV":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_UAV
	case "Enum_CodeAircraftBaseType_Member_1_ALL":
		*enum_codeaircraftbasetype_member_1 = Enum_CodeAircraftBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftbasetype_member_1 *Enum_CodeAircraftBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftBaseType_Member_1_LANDPLANE:
		res = "Enum_CodeAircraftBaseType_Member_1_LANDPLANE"
	case Enum_CodeAircraftBaseType_Member_1_SEAPLANE:
		res = "Enum_CodeAircraftBaseType_Member_1_SEAPLANE"
	case Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN:
		res = "Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN"
	case Enum_CodeAircraftBaseType_Member_1_HELICOPTER:
		res = "Enum_CodeAircraftBaseType_Member_1_HELICOPTER"
	case Enum_CodeAircraftBaseType_Member_1_GYROCOPTER:
		res = "Enum_CodeAircraftBaseType_Member_1_GYROCOPTER"
	case Enum_CodeAircraftBaseType_Member_1_TILT_WING:
		res = "Enum_CodeAircraftBaseType_Member_1_TILT_WING"
	case Enum_CodeAircraftBaseType_Member_1_STOL:
		res = "Enum_CodeAircraftBaseType_Member_1_STOL"
	case Enum_CodeAircraftBaseType_Member_1_GLIDER:
		res = "Enum_CodeAircraftBaseType_Member_1_GLIDER"
	case Enum_CodeAircraftBaseType_Member_1_HANGGLIDER:
		res = "Enum_CodeAircraftBaseType_Member_1_HANGGLIDER"
	case Enum_CodeAircraftBaseType_Member_1_PARAGLIDER:
		res = "Enum_CodeAircraftBaseType_Member_1_PARAGLIDER"
	case Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT:
		res = "Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT"
	case Enum_CodeAircraftBaseType_Member_1_BALLOON:
		res = "Enum_CodeAircraftBaseType_Member_1_BALLOON"
	case Enum_CodeAircraftBaseType_Member_1_UAV:
		res = "Enum_CodeAircraftBaseType_Member_1_UAV"
	case Enum_CodeAircraftBaseType_Member_1_ALL:
		res = "Enum_CodeAircraftBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeaircraftbasetype_member_1 Enum_CodeAircraftBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_LANDPLANE")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_SEAPLANE")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_AMPHIBIAN")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_HELICOPTER")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_GYROCOPTER")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_TILT_WING")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_STOL")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_GLIDER")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_HANGGLIDER")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_PARAGLIDER")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_ULTRA_LIGHT")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_BALLOON")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_UAV")
	res = append(res, "Enum_CodeAircraftBaseType_Member_1_ALL")

	return
}

func (enum_codeaircraftbasetype_member_1 Enum_CodeAircraftBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LANDPLANE")
	res = append(res, "SEAPLANE")
	res = append(res, "AMPHIBIAN")
	res = append(res, "HELICOPTER")
	res = append(res, "GYROCOPTER")
	res = append(res, "TILT_WING")
	res = append(res, "STOL")
	res = append(res, "GLIDER")
	res = append(res, "HANGGLIDER")
	res = append(res, "PARAGLIDER")
	res = append(res, "ULTRA_LIGHT")
	res = append(res, "BALLOON")
	res = append(res, "UAV")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeAircraftCategoryBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftcategorybasetype_member_1 Enum_CodeAircraftCategoryBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftCategoryBaseType_Member_1_A:
		res = "A"
	case Enum_CodeAircraftCategoryBaseType_Member_1_B:
		res = "B"
	case Enum_CodeAircraftCategoryBaseType_Member_1_C:
		res = "C"
	case Enum_CodeAircraftCategoryBaseType_Member_1_D:
		res = "D"
	case Enum_CodeAircraftCategoryBaseType_Member_1_E:
		res = "E"
	case Enum_CodeAircraftCategoryBaseType_Member_1_H:
		res = "H"
	case Enum_CodeAircraftCategoryBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeaircraftcategorybasetype_member_1 *Enum_CodeAircraftCategoryBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_A
		return
	case "B":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_B
		return
	case "C":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_C
		return
	case "D":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_D
		return
	case "E":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_E
		return
	case "H":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_H
		return
	case "ALL":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftcategorybasetype_member_1 *Enum_CodeAircraftCategoryBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftCategoryBaseType_Member_1_A":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_A
	case "Enum_CodeAircraftCategoryBaseType_Member_1_B":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_B
	case "Enum_CodeAircraftCategoryBaseType_Member_1_C":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_C
	case "Enum_CodeAircraftCategoryBaseType_Member_1_D":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_D
	case "Enum_CodeAircraftCategoryBaseType_Member_1_E":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_E
	case "Enum_CodeAircraftCategoryBaseType_Member_1_H":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_H
	case "Enum_CodeAircraftCategoryBaseType_Member_1_ALL":
		*enum_codeaircraftcategorybasetype_member_1 = Enum_CodeAircraftCategoryBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftcategorybasetype_member_1 *Enum_CodeAircraftCategoryBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftCategoryBaseType_Member_1_A:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_A"
	case Enum_CodeAircraftCategoryBaseType_Member_1_B:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_B"
	case Enum_CodeAircraftCategoryBaseType_Member_1_C:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_C"
	case Enum_CodeAircraftCategoryBaseType_Member_1_D:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_D"
	case Enum_CodeAircraftCategoryBaseType_Member_1_E:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_E"
	case Enum_CodeAircraftCategoryBaseType_Member_1_H:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_H"
	case Enum_CodeAircraftCategoryBaseType_Member_1_ALL:
		res = "Enum_CodeAircraftCategoryBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeaircraftcategorybasetype_member_1 Enum_CodeAircraftCategoryBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_A")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_B")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_C")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_D")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_E")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_H")
	res = append(res, "Enum_CodeAircraftCategoryBaseType_Member_1_ALL")

	return
}

func (enum_codeaircraftcategorybasetype_member_1 Enum_CodeAircraftCategoryBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")
	res = append(res, "E")
	res = append(res, "H")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeAircraftEngineBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftenginebasetype_member_1 Enum_CodeAircraftEngineBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftenginebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftEngineBaseType_Member_1_JET:
		res = "JET"
	case Enum_CodeAircraftEngineBaseType_Member_1_PISTON:
		res = "PISTON"
	case Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP:
		res = "TURBOPROP"
	case Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC:
		res = "ELECTRIC"
	case Enum_CodeAircraftEngineBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeaircraftenginebasetype_member_1 *Enum_CodeAircraftEngineBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "JET":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_JET
		return
	case "PISTON":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_PISTON
		return
	case "TURBOPROP":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP
		return
	case "ELECTRIC":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC
		return
	case "ALL":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftenginebasetype_member_1 *Enum_CodeAircraftEngineBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftEngineBaseType_Member_1_JET":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_JET
	case "Enum_CodeAircraftEngineBaseType_Member_1_PISTON":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_PISTON
	case "Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP
	case "Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC
	case "Enum_CodeAircraftEngineBaseType_Member_1_ALL":
		*enum_codeaircraftenginebasetype_member_1 = Enum_CodeAircraftEngineBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftenginebasetype_member_1 *Enum_CodeAircraftEngineBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftenginebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftEngineBaseType_Member_1_JET:
		res = "Enum_CodeAircraftEngineBaseType_Member_1_JET"
	case Enum_CodeAircraftEngineBaseType_Member_1_PISTON:
		res = "Enum_CodeAircraftEngineBaseType_Member_1_PISTON"
	case Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP:
		res = "Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP"
	case Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC:
		res = "Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC"
	case Enum_CodeAircraftEngineBaseType_Member_1_ALL:
		res = "Enum_CodeAircraftEngineBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeaircraftenginebasetype_member_1 Enum_CodeAircraftEngineBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftEngineBaseType_Member_1_JET")
	res = append(res, "Enum_CodeAircraftEngineBaseType_Member_1_PISTON")
	res = append(res, "Enum_CodeAircraftEngineBaseType_Member_1_TURBOPROP")
	res = append(res, "Enum_CodeAircraftEngineBaseType_Member_1_ELECTRIC")
	res = append(res, "Enum_CodeAircraftEngineBaseType_Member_1_ALL")

	return
}

func (enum_codeaircraftenginebasetype_member_1 Enum_CodeAircraftEngineBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "JET")
	res = append(res, "PISTON")
	res = append(res, "TURBOPROP")
	res = append(res, "ELECTRIC")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeAircraftEngineNumberBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftenginenumberbasetype_member_1 Enum_CodeAircraftEngineNumberBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftenginenumberbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1:
		res = "1"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2:
		res = "2"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3:
		res = "3"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4:
		res = "4"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6:
		res = "6"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8:
		res = "8"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C:
		res = "2C"
	}
	return
}

func (enum_codeaircraftenginenumberbasetype_member_1 *Enum_CodeAircraftEngineNumberBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1
		return
	case "2":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2
		return
	case "3":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3
		return
	case "4":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4
		return
	case "6":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6
		return
	case "8":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8
		return
	case "2C":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftenginenumberbasetype_member_1 *Enum_CodeAircraftEngineNumberBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8
	case "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C":
		*enum_codeaircraftenginenumberbasetype_member_1 = Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftenginenumberbasetype_member_1 *Enum_CodeAircraftEngineNumberBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftenginenumberbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8"
	case Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C:
		res = "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C"
	}
	return
}

func (enum_codeaircraftenginenumberbasetype_member_1 Enum_CodeAircraftEngineNumberBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_1")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_3")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_4")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_6")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_8")
	res = append(res, "Enum_CodeAircraftEngineNumberBaseType_Member_1_N_2C")

	return
}

func (enum_codeaircraftenginenumberbasetype_member_1 Enum_CodeAircraftEngineNumberBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1")
	res = append(res, "2")
	res = append(res, "3")
	res = append(res, "4")
	res = append(res, "6")
	res = append(res, "8")
	res = append(res, "2C")

	return
}

// Utility function for Enum_CodeAircraftStandBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftstandbasetype_member_1 Enum_CodeAircraftStandBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftstandbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftStandBaseType_Member_1_NI:
		res = "NI"
	case Enum_CodeAircraftStandBaseType_Member_1_ANG_NI:
		res = "ANG_NI"
	case Enum_CodeAircraftStandBaseType_Member_1_ANG_NO:
		res = "ANG_NO"
	case Enum_CodeAircraftStandBaseType_Member_1_PARL:
		res = "PARL"
	case Enum_CodeAircraftStandBaseType_Member_1_RMT:
		res = "RMT"
	case Enum_CodeAircraftStandBaseType_Member_1_ISOL:
		res = "ISOL"
	}
	return
}

func (enum_codeaircraftstandbasetype_member_1 *Enum_CodeAircraftStandBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NI":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_NI
		return
	case "ANG_NI":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ANG_NI
		return
	case "ANG_NO":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ANG_NO
		return
	case "PARL":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_PARL
		return
	case "RMT":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_RMT
		return
	case "ISOL":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ISOL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftstandbasetype_member_1 *Enum_CodeAircraftStandBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftStandBaseType_Member_1_NI":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_NI
	case "Enum_CodeAircraftStandBaseType_Member_1_ANG_NI":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ANG_NI
	case "Enum_CodeAircraftStandBaseType_Member_1_ANG_NO":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ANG_NO
	case "Enum_CodeAircraftStandBaseType_Member_1_PARL":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_PARL
	case "Enum_CodeAircraftStandBaseType_Member_1_RMT":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_RMT
	case "Enum_CodeAircraftStandBaseType_Member_1_ISOL":
		*enum_codeaircraftstandbasetype_member_1 = Enum_CodeAircraftStandBaseType_Member_1_ISOL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftstandbasetype_member_1 *Enum_CodeAircraftStandBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftstandbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftStandBaseType_Member_1_NI:
		res = "Enum_CodeAircraftStandBaseType_Member_1_NI"
	case Enum_CodeAircraftStandBaseType_Member_1_ANG_NI:
		res = "Enum_CodeAircraftStandBaseType_Member_1_ANG_NI"
	case Enum_CodeAircraftStandBaseType_Member_1_ANG_NO:
		res = "Enum_CodeAircraftStandBaseType_Member_1_ANG_NO"
	case Enum_CodeAircraftStandBaseType_Member_1_PARL:
		res = "Enum_CodeAircraftStandBaseType_Member_1_PARL"
	case Enum_CodeAircraftStandBaseType_Member_1_RMT:
		res = "Enum_CodeAircraftStandBaseType_Member_1_RMT"
	case Enum_CodeAircraftStandBaseType_Member_1_ISOL:
		res = "Enum_CodeAircraftStandBaseType_Member_1_ISOL"
	}
	return
}

func (enum_codeaircraftstandbasetype_member_1 Enum_CodeAircraftStandBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_NI")
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_ANG_NI")
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_ANG_NO")
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_PARL")
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_RMT")
	res = append(res, "Enum_CodeAircraftStandBaseType_Member_1_ISOL")

	return
}

func (enum_codeaircraftstandbasetype_member_1 Enum_CodeAircraftStandBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NI")
	res = append(res, "ANG_NI")
	res = append(res, "ANG_NO")
	res = append(res, "PARL")
	res = append(res, "RMT")
	res = append(res, "ISOL")

	return
}

// Utility function for Enum_CodeAircraftWingspanClassBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftwingspanclassbasetype_member_1 Enum_CodeAircraftWingspanClassBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftwingspanclassbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_I:
		res = "I"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_II:
		res = "II"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_III:
		res = "III"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_IV:
		res = "IV"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_V:
		res = "V"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_VI:
		res = "VI"
	}
	return
}

func (enum_codeaircraftwingspanclassbasetype_member_1 *Enum_CodeAircraftWingspanClassBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "I":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_I
		return
	case "II":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_II
		return
	case "III":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_III
		return
	case "IV":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_IV
		return
	case "V":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_V
		return
	case "VI":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_VI
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftwingspanclassbasetype_member_1 *Enum_CodeAircraftWingspanClassBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_I":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_I
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_II":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_II
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_III":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_III
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_IV":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_IV
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_V":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_V
	case "Enum_CodeAircraftWingspanClassBaseType_Member_1_VI":
		*enum_codeaircraftwingspanclassbasetype_member_1 = Enum_CodeAircraftWingspanClassBaseType_Member_1_VI
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftwingspanclassbasetype_member_1 *Enum_CodeAircraftWingspanClassBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftwingspanclassbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_I:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_I"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_II:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_II"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_III:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_III"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_IV:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_IV"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_V:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_V"
	case Enum_CodeAircraftWingspanClassBaseType_Member_1_VI:
		res = "Enum_CodeAircraftWingspanClassBaseType_Member_1_VI"
	}
	return
}

func (enum_codeaircraftwingspanclassbasetype_member_1 Enum_CodeAircraftWingspanClassBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_I")
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_II")
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_III")
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_IV")
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_V")
	res = append(res, "Enum_CodeAircraftWingspanClassBaseType_Member_1_VI")

	return
}

func (enum_codeaircraftwingspanclassbasetype_member_1 Enum_CodeAircraftWingspanClassBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "I")
	res = append(res, "II")
	res = append(res, "III")
	res = append(res, "IV")
	res = append(res, "V")
	res = append(res, "VI")

	return
}

// Utility function for Enum_CodeAirportHeliportBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairportheliportbasetype_member_1 Enum_CodeAirportHeliportBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairportheliportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportHeliportBaseType_Member_1_AD:
		res = "AD"
	case Enum_CodeAirportHeliportBaseType_Member_1_AH:
		res = "AH"
	case Enum_CodeAirportHeliportBaseType_Member_1_HP:
		res = "HP"
	case Enum_CodeAirportHeliportBaseType_Member_1_LS:
		res = "LS"
	}
	return
}

func (enum_codeairportheliportbasetype_member_1 *Enum_CodeAirportHeliportBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AD":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_AD
		return
	case "AH":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_AH
		return
	case "HP":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_HP
		return
	case "LS":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_LS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairportheliportbasetype_member_1 *Enum_CodeAirportHeliportBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirportHeliportBaseType_Member_1_AD":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_AD
	case "Enum_CodeAirportHeliportBaseType_Member_1_AH":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_AH
	case "Enum_CodeAirportHeliportBaseType_Member_1_HP":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_HP
	case "Enum_CodeAirportHeliportBaseType_Member_1_LS":
		*enum_codeairportheliportbasetype_member_1 = Enum_CodeAirportHeliportBaseType_Member_1_LS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairportheliportbasetype_member_1 *Enum_CodeAirportHeliportBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairportheliportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportHeliportBaseType_Member_1_AD:
		res = "Enum_CodeAirportHeliportBaseType_Member_1_AD"
	case Enum_CodeAirportHeliportBaseType_Member_1_AH:
		res = "Enum_CodeAirportHeliportBaseType_Member_1_AH"
	case Enum_CodeAirportHeliportBaseType_Member_1_HP:
		res = "Enum_CodeAirportHeliportBaseType_Member_1_HP"
	case Enum_CodeAirportHeliportBaseType_Member_1_LS:
		res = "Enum_CodeAirportHeliportBaseType_Member_1_LS"
	}
	return
}

func (enum_codeairportheliportbasetype_member_1 Enum_CodeAirportHeliportBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirportHeliportBaseType_Member_1_AD")
	res = append(res, "Enum_CodeAirportHeliportBaseType_Member_1_AH")
	res = append(res, "Enum_CodeAirportHeliportBaseType_Member_1_HP")
	res = append(res, "Enum_CodeAirportHeliportBaseType_Member_1_LS")

	return
}

func (enum_codeairportheliportbasetype_member_1 Enum_CodeAirportHeliportBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AD")
	res = append(res, "AH")
	res = append(res, "HP")
	res = append(res, "LS")

	return
}

// Utility function for Enum_CodeAirportHeliportCollocationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairportheliportcollocationbasetype_member_1 Enum_CodeAirportHeliportCollocationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairportheliportcollocationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL:
		res = "FULL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY:
		res = "RWY"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL:
		res = "PARTIAL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL:
		res = "UNILATERAL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED:
		res = "SEPARATED"
	}
	return
}

func (enum_codeairportheliportcollocationbasetype_member_1 *Enum_CodeAirportHeliportCollocationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FULL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL
		return
	case "RWY":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY
		return
	case "PARTIAL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL
		return
	case "UNILATERAL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL
		return
	case "SEPARATED":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairportheliportcollocationbasetype_member_1 *Enum_CodeAirportHeliportCollocationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL
	case "Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY
	case "Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL
	case "Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL
	case "Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED":
		*enum_codeairportheliportcollocationbasetype_member_1 = Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairportheliportcollocationbasetype_member_1 *Enum_CodeAirportHeliportCollocationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairportheliportcollocationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL:
		res = "Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY:
		res = "Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL:
		res = "Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL:
		res = "Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL"
	case Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED:
		res = "Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED"
	}
	return
}

func (enum_codeairportheliportcollocationbasetype_member_1 Enum_CodeAirportHeliportCollocationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirportHeliportCollocationBaseType_Member_1_FULL")
	res = append(res, "Enum_CodeAirportHeliportCollocationBaseType_Member_1_RWY")
	res = append(res, "Enum_CodeAirportHeliportCollocationBaseType_Member_1_PARTIAL")
	res = append(res, "Enum_CodeAirportHeliportCollocationBaseType_Member_1_UNILATERAL")
	res = append(res, "Enum_CodeAirportHeliportCollocationBaseType_Member_1_SEPARATED")

	return
}

func (enum_codeairportheliportcollocationbasetype_member_1 Enum_CodeAirportHeliportCollocationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FULL")
	res = append(res, "RWY")
	res = append(res, "PARTIAL")
	res = append(res, "UNILATERAL")
	res = append(res, "SEPARATED")

	return
}

// Utility function for Enum_CodeAirportWarningBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairportwarningbasetype_member_1 Enum_CodeAirportWarningBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairportwarningbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportWarningBaseType_Member_1_WIP:
		res = "WIP"
	case Enum_CodeAirportWarningBaseType_Member_1_EQUIP:
		res = "EQUIP"
	case Enum_CodeAirportWarningBaseType_Member_1_BIRD:
		res = "BIRD"
	case Enum_CodeAirportWarningBaseType_Member_1_ANIMAL:
		res = "ANIMAL"
	case Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL:
		res = "RUBBER_REMOVAL"
	case Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT:
		res = "PARKED_ACFT"
	case Enum_CodeAirportWarningBaseType_Member_1_RESURFACING:
		res = "RESURFACING"
	case Enum_CodeAirportWarningBaseType_Member_1_PAVING:
		res = "PAVING"
	case Enum_CodeAirportWarningBaseType_Member_1_PAINTING:
		res = "PAINTING"
	case Enum_CodeAirportWarningBaseType_Member_1_INSPECTION:
		res = "INSPECTION"
	case Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING:
		res = "GRASS_CUTTING"
	case Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION:
		res = "CALIBRATION"
	}
	return
}

func (enum_codeairportwarningbasetype_member_1 *Enum_CodeAirportWarningBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "WIP":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_WIP
		return
	case "EQUIP":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_EQUIP
		return
	case "BIRD":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_BIRD
		return
	case "ANIMAL":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_ANIMAL
		return
	case "RUBBER_REMOVAL":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL
		return
	case "PARKED_ACFT":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT
		return
	case "RESURFACING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_RESURFACING
		return
	case "PAVING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PAVING
		return
	case "PAINTING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PAINTING
		return
	case "INSPECTION":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_INSPECTION
		return
	case "GRASS_CUTTING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING
		return
	case "CALIBRATION":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairportwarningbasetype_member_1 *Enum_CodeAirportWarningBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirportWarningBaseType_Member_1_WIP":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_WIP
	case "Enum_CodeAirportWarningBaseType_Member_1_EQUIP":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_EQUIP
	case "Enum_CodeAirportWarningBaseType_Member_1_BIRD":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_BIRD
	case "Enum_CodeAirportWarningBaseType_Member_1_ANIMAL":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_ANIMAL
	case "Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL
	case "Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT
	case "Enum_CodeAirportWarningBaseType_Member_1_RESURFACING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_RESURFACING
	case "Enum_CodeAirportWarningBaseType_Member_1_PAVING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PAVING
	case "Enum_CodeAirportWarningBaseType_Member_1_PAINTING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_PAINTING
	case "Enum_CodeAirportWarningBaseType_Member_1_INSPECTION":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_INSPECTION
	case "Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING
	case "Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION":
		*enum_codeairportwarningbasetype_member_1 = Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairportwarningbasetype_member_1 *Enum_CodeAirportWarningBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairportwarningbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirportWarningBaseType_Member_1_WIP:
		res = "Enum_CodeAirportWarningBaseType_Member_1_WIP"
	case Enum_CodeAirportWarningBaseType_Member_1_EQUIP:
		res = "Enum_CodeAirportWarningBaseType_Member_1_EQUIP"
	case Enum_CodeAirportWarningBaseType_Member_1_BIRD:
		res = "Enum_CodeAirportWarningBaseType_Member_1_BIRD"
	case Enum_CodeAirportWarningBaseType_Member_1_ANIMAL:
		res = "Enum_CodeAirportWarningBaseType_Member_1_ANIMAL"
	case Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL:
		res = "Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL"
	case Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT:
		res = "Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT"
	case Enum_CodeAirportWarningBaseType_Member_1_RESURFACING:
		res = "Enum_CodeAirportWarningBaseType_Member_1_RESURFACING"
	case Enum_CodeAirportWarningBaseType_Member_1_PAVING:
		res = "Enum_CodeAirportWarningBaseType_Member_1_PAVING"
	case Enum_CodeAirportWarningBaseType_Member_1_PAINTING:
		res = "Enum_CodeAirportWarningBaseType_Member_1_PAINTING"
	case Enum_CodeAirportWarningBaseType_Member_1_INSPECTION:
		res = "Enum_CodeAirportWarningBaseType_Member_1_INSPECTION"
	case Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING:
		res = "Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING"
	case Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION:
		res = "Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION"
	}
	return
}

func (enum_codeairportwarningbasetype_member_1 Enum_CodeAirportWarningBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_WIP")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_EQUIP")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_BIRD")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_ANIMAL")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_RUBBER_REMOVAL")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_PARKED_ACFT")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_RESURFACING")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_PAVING")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_PAINTING")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_INSPECTION")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_GRASS_CUTTING")
	res = append(res, "Enum_CodeAirportWarningBaseType_Member_1_CALIBRATION")

	return
}

func (enum_codeairportwarningbasetype_member_1 Enum_CodeAirportWarningBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "WIP")
	res = append(res, "EQUIP")
	res = append(res, "BIRD")
	res = append(res, "ANIMAL")
	res = append(res, "RUBBER_REMOVAL")
	res = append(res, "PARKED_ACFT")
	res = append(res, "RESURFACING")
	res = append(res, "PAVING")
	res = append(res, "PAINTING")
	res = append(res, "INSPECTION")
	res = append(res, "GRASS_CUTTING")
	res = append(res, "CALIBRATION")

	return
}

// Utility function for Enum_CodeApproachGuidanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachguidancebasetype_member_1 Enum_CodeApproachGuidanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION:
		res = "NON_PRECISION"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I:
		res = "ILS_PRECISION_CAT_I"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II:
		res = "ILS_PRECISION_CAT_II"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA:
		res = "ILS_PRECISION_CAT_IIIA"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB:
		res = "ILS_PRECISION_CAT_IIIB"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC:
		res = "ILS_PRECISION_CAT_IIIC"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID:
		res = "ILS_PRECISION_CAT_IIID"
	case Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION:
		res = "MLS_PRECISION"
	}
	return
}

func (enum_codeapproachguidancebasetype_member_1 *Enum_CodeApproachGuidanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NON_PRECISION":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION
		return
	case "ILS_PRECISION_CAT_I":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I
		return
	case "ILS_PRECISION_CAT_II":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II
		return
	case "ILS_PRECISION_CAT_IIIA":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA
		return
	case "ILS_PRECISION_CAT_IIIB":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB
		return
	case "ILS_PRECISION_CAT_IIIC":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC
		return
	case "ILS_PRECISION_CAT_IIID":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID
		return
	case "MLS_PRECISION":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachguidancebasetype_member_1 *Enum_CodeApproachGuidanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC
	case "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID
	case "Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION":
		*enum_codeapproachguidancebasetype_member_1 = Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachguidancebasetype_member_1 *Enum_CodeApproachGuidanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC"
	case Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID"
	case Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION:
		res = "Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION"
	}
	return
}

func (enum_codeapproachguidancebasetype_member_1 Enum_CodeApproachGuidanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_NON_PRECISION")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_I")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_II")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIA")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIB")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIIC")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_ILS_PRECISION_CAT_IIID")
	res = append(res, "Enum_CodeApproachGuidanceBaseType_Member_1_MLS_PRECISION")

	return
}

func (enum_codeapproachguidancebasetype_member_1 Enum_CodeApproachGuidanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NON_PRECISION")
	res = append(res, "ILS_PRECISION_CAT_I")
	res = append(res, "ILS_PRECISION_CAT_II")
	res = append(res, "ILS_PRECISION_CAT_IIIA")
	res = append(res, "ILS_PRECISION_CAT_IIIB")
	res = append(res, "ILS_PRECISION_CAT_IIIC")
	res = append(res, "ILS_PRECISION_CAT_IIID")
	res = append(res, "MLS_PRECISION")

	return
}

// Utility function for Enum_CodeApproachLightingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachlightingbasetype_member_1 Enum_CodeApproachLightingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachLightingBaseType_Member_1_ALSAF:
		res = "ALSAF"
	case Enum_CodeApproachLightingBaseType_Member_1_MALS:
		res = "MALS"
	case Enum_CodeApproachLightingBaseType_Member_1_MALSR:
		res = "MALSR"
	case Enum_CodeApproachLightingBaseType_Member_1_SALS:
		res = "SALS"
	case Enum_CodeApproachLightingBaseType_Member_1_SSALS:
		res = "SSALS"
	case Enum_CodeApproachLightingBaseType_Member_1_SSALR:
		res = "SSALR"
	case Enum_CodeApproachLightingBaseType_Member_1_LDIN:
		res = "LDIN"
	case Enum_CodeApproachLightingBaseType_Member_1_ODALS:
		res = "ODALS"
	case Enum_CodeApproachLightingBaseType_Member_1_AFOVRN:
		res = "AFOVRN"
	case Enum_CodeApproachLightingBaseType_Member_1_MILOVRN:
		res = "MILOVRN"
	case Enum_CodeApproachLightingBaseType_Member_1_CALVERT:
		res = "CALVERT"
	}
	return
}

func (enum_codeapproachlightingbasetype_member_1 *Enum_CodeApproachLightingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ALSAF":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_ALSAF
		return
	case "MALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MALS
		return
	case "MALSR":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MALSR
		return
	case "SALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SALS
		return
	case "SSALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SSALS
		return
	case "SSALR":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SSALR
		return
	case "LDIN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_LDIN
		return
	case "ODALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_ODALS
		return
	case "AFOVRN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_AFOVRN
		return
	case "MILOVRN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MILOVRN
		return
	case "CALVERT":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_CALVERT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachlightingbasetype_member_1 *Enum_CodeApproachLightingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachLightingBaseType_Member_1_ALSAF":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_ALSAF
	case "Enum_CodeApproachLightingBaseType_Member_1_MALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MALS
	case "Enum_CodeApproachLightingBaseType_Member_1_MALSR":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MALSR
	case "Enum_CodeApproachLightingBaseType_Member_1_SALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SALS
	case "Enum_CodeApproachLightingBaseType_Member_1_SSALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SSALS
	case "Enum_CodeApproachLightingBaseType_Member_1_SSALR":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_SSALR
	case "Enum_CodeApproachLightingBaseType_Member_1_LDIN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_LDIN
	case "Enum_CodeApproachLightingBaseType_Member_1_ODALS":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_ODALS
	case "Enum_CodeApproachLightingBaseType_Member_1_AFOVRN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_AFOVRN
	case "Enum_CodeApproachLightingBaseType_Member_1_MILOVRN":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_MILOVRN
	case "Enum_CodeApproachLightingBaseType_Member_1_CALVERT":
		*enum_codeapproachlightingbasetype_member_1 = Enum_CodeApproachLightingBaseType_Member_1_CALVERT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachlightingbasetype_member_1 *Enum_CodeApproachLightingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachLightingBaseType_Member_1_ALSAF:
		res = "Enum_CodeApproachLightingBaseType_Member_1_ALSAF"
	case Enum_CodeApproachLightingBaseType_Member_1_MALS:
		res = "Enum_CodeApproachLightingBaseType_Member_1_MALS"
	case Enum_CodeApproachLightingBaseType_Member_1_MALSR:
		res = "Enum_CodeApproachLightingBaseType_Member_1_MALSR"
	case Enum_CodeApproachLightingBaseType_Member_1_SALS:
		res = "Enum_CodeApproachLightingBaseType_Member_1_SALS"
	case Enum_CodeApproachLightingBaseType_Member_1_SSALS:
		res = "Enum_CodeApproachLightingBaseType_Member_1_SSALS"
	case Enum_CodeApproachLightingBaseType_Member_1_SSALR:
		res = "Enum_CodeApproachLightingBaseType_Member_1_SSALR"
	case Enum_CodeApproachLightingBaseType_Member_1_LDIN:
		res = "Enum_CodeApproachLightingBaseType_Member_1_LDIN"
	case Enum_CodeApproachLightingBaseType_Member_1_ODALS:
		res = "Enum_CodeApproachLightingBaseType_Member_1_ODALS"
	case Enum_CodeApproachLightingBaseType_Member_1_AFOVRN:
		res = "Enum_CodeApproachLightingBaseType_Member_1_AFOVRN"
	case Enum_CodeApproachLightingBaseType_Member_1_MILOVRN:
		res = "Enum_CodeApproachLightingBaseType_Member_1_MILOVRN"
	case Enum_CodeApproachLightingBaseType_Member_1_CALVERT:
		res = "Enum_CodeApproachLightingBaseType_Member_1_CALVERT"
	}
	return
}

func (enum_codeapproachlightingbasetype_member_1 Enum_CodeApproachLightingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_ALSAF")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_MALS")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_MALSR")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_SALS")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_SSALS")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_SSALR")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_LDIN")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_ODALS")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_AFOVRN")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_MILOVRN")
	res = append(res, "Enum_CodeApproachLightingBaseType_Member_1_CALVERT")

	return
}

func (enum_codeapproachlightingbasetype_member_1 Enum_CodeApproachLightingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ALSAF")
	res = append(res, "MALS")
	res = append(res, "MALSR")
	res = append(res, "SALS")
	res = append(res, "SSALS")
	res = append(res, "SSALR")
	res = append(res, "LDIN")
	res = append(res, "ODALS")
	res = append(res, "AFOVRN")
	res = append(res, "MILOVRN")
	res = append(res, "CALVERT")

	return
}

// Utility function for Enum_CodeApproachLightingICAOBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachlightingicaobasetype_member_1 Enum_CodeApproachLightingICAOBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachlightingicaobasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE:
		res = "SIMPLE"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1:
		res = "CAT1"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23:
		res = "CAT23"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING:
		res = "CIRCLING"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN:
		res = "LEADIN"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_NONE:
		res = "NONE"
	}
	return
}

func (enum_codeapproachlightingicaobasetype_member_1 *Enum_CodeApproachLightingICAOBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SIMPLE":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE
		return
	case "CAT1":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1
		return
	case "CAT23":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23
		return
	case "CIRCLING":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING
		return
	case "LEADIN":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN
		return
	case "NONE":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_NONE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachlightingicaobasetype_member_1 *Enum_CodeApproachLightingICAOBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN
	case "Enum_CodeApproachLightingICAOBaseType_Member_1_NONE":
		*enum_codeapproachlightingicaobasetype_member_1 = Enum_CodeApproachLightingICAOBaseType_Member_1_NONE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachlightingicaobasetype_member_1 *Enum_CodeApproachLightingICAOBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachlightingicaobasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN"
	case Enum_CodeApproachLightingICAOBaseType_Member_1_NONE:
		res = "Enum_CodeApproachLightingICAOBaseType_Member_1_NONE"
	}
	return
}

func (enum_codeapproachlightingicaobasetype_member_1 Enum_CodeApproachLightingICAOBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_SIMPLE")
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT1")
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_CAT23")
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_CIRCLING")
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_LEADIN")
	res = append(res, "Enum_CodeApproachLightingICAOBaseType_Member_1_NONE")

	return
}

func (enum_codeapproachlightingicaobasetype_member_1 Enum_CodeApproachLightingICAOBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SIMPLE")
	res = append(res, "CAT1")
	res = append(res, "CAT23")
	res = append(res, "CIRCLING")
	res = append(res, "LEADIN")
	res = append(res, "NONE")

	return
}

// Utility function for Enum_CodeApronElementBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapronelementbasetype_member_1 Enum_CodeApronElementBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapronelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApronElementBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeApronElementBaseType_Member_1_PARKING:
		res = "PARKING"
	case Enum_CodeApronElementBaseType_Member_1_RAMP:
		res = "RAMP"
	case Enum_CodeApronElementBaseType_Member_1_CARGO:
		res = "CARGO"
	case Enum_CodeApronElementBaseType_Member_1_FUEL:
		res = "FUEL"
	case Enum_CodeApronElementBaseType_Member_1_HARDSTAND:
		res = "HARDSTAND"
	case Enum_CodeApronElementBaseType_Member_1_MAINT:
		res = "MAINT"
	case Enum_CodeApronElementBaseType_Member_1_MILITARY:
		res = "MILITARY"
	case Enum_CodeApronElementBaseType_Member_1_LOADING:
		res = "LOADING"
	case Enum_CodeApronElementBaseType_Member_1_TAXILANE:
		res = "TAXILANE"
	case Enum_CodeApronElementBaseType_Member_1_TURNAROUND:
		res = "TURNAROUND"
	case Enum_CodeApronElementBaseType_Member_1_TEMPORARY:
		res = "TEMPORARY"
	case Enum_CodeApronElementBaseType_Member_1_STAIRS:
		res = "STAIRS"
	}
	return
}

func (enum_codeapronelementbasetype_member_1 *Enum_CodeApronElementBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_NORMAL
		return
	case "PARKING":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_PARKING
		return
	case "RAMP":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_RAMP
		return
	case "CARGO":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_CARGO
		return
	case "FUEL":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_FUEL
		return
	case "HARDSTAND":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_HARDSTAND
		return
	case "MAINT":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_MAINT
		return
	case "MILITARY":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_MILITARY
		return
	case "LOADING":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_LOADING
		return
	case "TAXILANE":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TAXILANE
		return
	case "TURNAROUND":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TURNAROUND
		return
	case "TEMPORARY":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TEMPORARY
		return
	case "STAIRS":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_STAIRS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapronelementbasetype_member_1 *Enum_CodeApronElementBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApronElementBaseType_Member_1_NORMAL":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_NORMAL
	case "Enum_CodeApronElementBaseType_Member_1_PARKING":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_PARKING
	case "Enum_CodeApronElementBaseType_Member_1_RAMP":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_RAMP
	case "Enum_CodeApronElementBaseType_Member_1_CARGO":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_CARGO
	case "Enum_CodeApronElementBaseType_Member_1_FUEL":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_FUEL
	case "Enum_CodeApronElementBaseType_Member_1_HARDSTAND":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_HARDSTAND
	case "Enum_CodeApronElementBaseType_Member_1_MAINT":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_MAINT
	case "Enum_CodeApronElementBaseType_Member_1_MILITARY":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_MILITARY
	case "Enum_CodeApronElementBaseType_Member_1_LOADING":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_LOADING
	case "Enum_CodeApronElementBaseType_Member_1_TAXILANE":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TAXILANE
	case "Enum_CodeApronElementBaseType_Member_1_TURNAROUND":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TURNAROUND
	case "Enum_CodeApronElementBaseType_Member_1_TEMPORARY":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_TEMPORARY
	case "Enum_CodeApronElementBaseType_Member_1_STAIRS":
		*enum_codeapronelementbasetype_member_1 = Enum_CodeApronElementBaseType_Member_1_STAIRS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapronelementbasetype_member_1 *Enum_CodeApronElementBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapronelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApronElementBaseType_Member_1_NORMAL:
		res = "Enum_CodeApronElementBaseType_Member_1_NORMAL"
	case Enum_CodeApronElementBaseType_Member_1_PARKING:
		res = "Enum_CodeApronElementBaseType_Member_1_PARKING"
	case Enum_CodeApronElementBaseType_Member_1_RAMP:
		res = "Enum_CodeApronElementBaseType_Member_1_RAMP"
	case Enum_CodeApronElementBaseType_Member_1_CARGO:
		res = "Enum_CodeApronElementBaseType_Member_1_CARGO"
	case Enum_CodeApronElementBaseType_Member_1_FUEL:
		res = "Enum_CodeApronElementBaseType_Member_1_FUEL"
	case Enum_CodeApronElementBaseType_Member_1_HARDSTAND:
		res = "Enum_CodeApronElementBaseType_Member_1_HARDSTAND"
	case Enum_CodeApronElementBaseType_Member_1_MAINT:
		res = "Enum_CodeApronElementBaseType_Member_1_MAINT"
	case Enum_CodeApronElementBaseType_Member_1_MILITARY:
		res = "Enum_CodeApronElementBaseType_Member_1_MILITARY"
	case Enum_CodeApronElementBaseType_Member_1_LOADING:
		res = "Enum_CodeApronElementBaseType_Member_1_LOADING"
	case Enum_CodeApronElementBaseType_Member_1_TAXILANE:
		res = "Enum_CodeApronElementBaseType_Member_1_TAXILANE"
	case Enum_CodeApronElementBaseType_Member_1_TURNAROUND:
		res = "Enum_CodeApronElementBaseType_Member_1_TURNAROUND"
	case Enum_CodeApronElementBaseType_Member_1_TEMPORARY:
		res = "Enum_CodeApronElementBaseType_Member_1_TEMPORARY"
	case Enum_CodeApronElementBaseType_Member_1_STAIRS:
		res = "Enum_CodeApronElementBaseType_Member_1_STAIRS"
	}
	return
}

func (enum_codeapronelementbasetype_member_1 Enum_CodeApronElementBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_PARKING")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_RAMP")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_CARGO")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_FUEL")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_HARDSTAND")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_MAINT")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_MILITARY")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_LOADING")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_TAXILANE")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_TURNAROUND")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_TEMPORARY")
	res = append(res, "Enum_CodeApronElementBaseType_Member_1_STAIRS")

	return
}

func (enum_codeapronelementbasetype_member_1 Enum_CodeApronElementBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "PARKING")
	res = append(res, "RAMP")
	res = append(res, "CARGO")
	res = append(res, "FUEL")
	res = append(res, "HARDSTAND")
	res = append(res, "MAINT")
	res = append(res, "MILITARY")
	res = append(res, "LOADING")
	res = append(res, "TAXILANE")
	res = append(res, "TURNAROUND")
	res = append(res, "TEMPORARY")
	res = append(res, "STAIRS")

	return
}

// Utility function for Enum_CodeApronSectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapronsectionbasetype_member_1 Enum_CodeApronSectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapronsectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApronSectionBaseType_Member_1_EDGE:
		res = "EDGE"
	}
	return
}

func (enum_codeapronsectionbasetype_member_1 *Enum_CodeApronSectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EDGE":
		*enum_codeapronsectionbasetype_member_1 = Enum_CodeApronSectionBaseType_Member_1_EDGE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapronsectionbasetype_member_1 *Enum_CodeApronSectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApronSectionBaseType_Member_1_EDGE":
		*enum_codeapronsectionbasetype_member_1 = Enum_CodeApronSectionBaseType_Member_1_EDGE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapronsectionbasetype_member_1 *Enum_CodeApronSectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapronsectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApronSectionBaseType_Member_1_EDGE:
		res = "Enum_CodeApronSectionBaseType_Member_1_EDGE"
	}
	return
}

func (enum_codeapronsectionbasetype_member_1 Enum_CodeApronSectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApronSectionBaseType_Member_1_EDGE")

	return
}

func (enum_codeapronsectionbasetype_member_1 Enum_CodeApronSectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EDGE")

	return
}

// Utility function for Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codearrestinggearenergyabsorbbasetype_member_1 Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codearrestinggearenergyabsorbbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300:
		res = "ROTARY_1300"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800:
		res = "ROTARY_2800"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A:
		res = "ROTARY_34B_1A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B:
		res = "ROTARY_34B_1B"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C:
		res = "ROTARY_34B_1C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F:
		res = "ROTARY_34D_1F"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C:
		res = "ROTARY_44B_2C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D:
		res = "ROTARY_44B_2D"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E:
		res = "ROTARY_44B_2E"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F:
		res = "ROTARY_44B_2F"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H:
		res = "ROTARY_44B_2H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I:
		res = "ROTARY_44B_2I"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L:
		res = "ROTARY_44B_2L"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A:
		res = "ROTARY_44B_3A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H:
		res = "ROTARY_44B_3H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L:
		res = "ROTARY_44B_3L"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C:
		res = "ROTARY_44B_4C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E:
		res = "ROTARY_44B_4E"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H:
		res = "ROTARY_44B_4H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S:
		res = "ROTARY_500S"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4:
		res = "ROTARY_500S_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6:
		res = "ROTARY_500S_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8:
		res = "ROTARY_500S_8"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8:
		res = "ROTARYTRANS_500S_8"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64:
		res = "ROTARY_AAE_64"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A:
		res = "ROTARY_BAK_12A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B:
		res = "ROTARY_BAK_12B"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13:
		res = "ROTARY_BAK_13"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6:
		res = "LINEAR_BAK_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9:
		res = "ROTARY_BAK_9"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3:
		res = "DISK_BEFAB_12_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4:
		res = "DISK_BEFAB_20_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2:
		res = "DISK_BEFAB_21_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4:
		res = "DISK_BEFAB_24_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2:
		res = "DISK_BEFAB_56_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3:
		res = "DISK_BEFAB_6_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2:
		res = "DISK_BEFAB_60_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3:
		res = "DISK_BEFAB_8_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG:
		res = "CHAIN_CHAG"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12:
		res = "ROTARY_DUAL_BAK_12"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15:
		res = "ROTARY_E15"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27:
		res = "ROTARY_E27"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28:
		res = "ROTARY_E28"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5:
		res = "CHAIN_E5"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1:
		res = "CHAIN_E5_1"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2:
		res = "CHAIN_E5_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3:
		res = "CHAIN_E5_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6:
		res = "CHAIN_E6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP:
		res = "ROTARY_CHAIN_JETSTOP"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21:
		res = "MOBILROTARY_M21"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS:
		res = "MOBILROTARY_MAAS"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I:
		res = "MOBILROTARY_MAG_I"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II:
		res = "MOBILROTARY_MAG_II"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III:
		res = "MOBILROTARY_MAG_III"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV:
		res = "MOBILROTARY_MAG_IV"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX:
		res = "MOBILROTARY_MAG_IX"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI:
		res = "MOBILROTARY_MAG_VI"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII:
		res = "MOBILROTARY_MAG_VII"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII:
		res = "MOBILROTARY_MAG_VIII"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X:
		res = "MOBILROTARY_MAG_X"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100:
		res = "TEXTILE_MB_100"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60:
		res = "TEXTILE_MB_60"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER:
		res = "MOBILROTARY_HYDRAULIC_WATER"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21:
		res = "ROTARY_PUAG_MK_21"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A:
		res = "DISK_RAF_MK_12A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6:
		res = "DISK_RAF_MK_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR:
		res = "RAF_PORTABLE_AGEAR"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3:
		res = "DISK_RAF_TYPEA_BEFAB_6_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3:
		res = "DISK_RAF_TYPEB_BEFAB_12_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1:
		res = "ROTARY_RHAG_MK_1"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER:
		res = "ROTARY_HYDRAULIC_WATER"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND:
		res = "BARRIER_DISK_SAFELAND"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1:
		res = "LINEAR_SPRAG_MK_1"
	}
	return
}

func (enum_codearrestinggearenergyabsorbbasetype_member_1 *Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ROTARY_1300":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300
		return
	case "ROTARY_2800":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800
		return
	case "ROTARY_34B_1A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A
		return
	case "ROTARY_34B_1B":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B
		return
	case "ROTARY_34B_1C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C
		return
	case "ROTARY_34D_1F":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F
		return
	case "ROTARY_44B_2C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C
		return
	case "ROTARY_44B_2D":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D
		return
	case "ROTARY_44B_2E":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E
		return
	case "ROTARY_44B_2F":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F
		return
	case "ROTARY_44B_2H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H
		return
	case "ROTARY_44B_2I":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I
		return
	case "ROTARY_44B_2L":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L
		return
	case "ROTARY_44B_3A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A
		return
	case "ROTARY_44B_3H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H
		return
	case "ROTARY_44B_3L":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L
		return
	case "ROTARY_44B_4C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C
		return
	case "ROTARY_44B_4E":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E
		return
	case "ROTARY_44B_4H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H
		return
	case "ROTARY_500S":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S
		return
	case "ROTARY_500S_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4
		return
	case "ROTARY_500S_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6
		return
	case "ROTARY_500S_8":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8
		return
	case "ROTARYTRANS_500S_8":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8
		return
	case "ROTARY_AAE_64":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64
		return
	case "ROTARY_BAK_12A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A
		return
	case "ROTARY_BAK_12B":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B
		return
	case "ROTARY_BAK_13":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13
		return
	case "LINEAR_BAK_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6
		return
	case "ROTARY_BAK_9":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9
		return
	case "DISK_BEFAB_12_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3
		return
	case "DISK_BEFAB_20_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4
		return
	case "DISK_BEFAB_21_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2
		return
	case "DISK_BEFAB_24_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4
		return
	case "DISK_BEFAB_56_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2
		return
	case "DISK_BEFAB_6_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3
		return
	case "DISK_BEFAB_60_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2
		return
	case "DISK_BEFAB_8_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3
		return
	case "CHAIN_CHAG":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG
		return
	case "ROTARY_DUAL_BAK_12":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12
		return
	case "ROTARY_E15":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15
		return
	case "ROTARY_E27":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27
		return
	case "ROTARY_E28":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28
		return
	case "CHAIN_E5":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5
		return
	case "CHAIN_E5_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1
		return
	case "CHAIN_E5_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2
		return
	case "CHAIN_E5_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3
		return
	case "CHAIN_E6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6
		return
	case "ROTARY_CHAIN_JETSTOP":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP
		return
	case "MOBILROTARY_M21":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21
		return
	case "MOBILROTARY_MAAS":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS
		return
	case "MOBILROTARY_MAG_I":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I
		return
	case "MOBILROTARY_MAG_II":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II
		return
	case "MOBILROTARY_MAG_III":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III
		return
	case "MOBILROTARY_MAG_IV":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV
		return
	case "MOBILROTARY_MAG_IX":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX
		return
	case "MOBILROTARY_MAG_VI":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI
		return
	case "MOBILROTARY_MAG_VII":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII
		return
	case "MOBILROTARY_MAG_VIII":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII
		return
	case "MOBILROTARY_MAG_X":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X
		return
	case "TEXTILE_MB_100":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100
		return
	case "TEXTILE_MB_60":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60
		return
	case "MOBILROTARY_HYDRAULIC_WATER":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER
		return
	case "ROTARY_PUAG_MK_21":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21
		return
	case "DISK_RAF_MK_12A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A
		return
	case "DISK_RAF_MK_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6
		return
	case "RAF_PORTABLE_AGEAR":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR
		return
	case "DISK_RAF_TYPEA_BEFAB_6_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3
		return
	case "DISK_RAF_TYPEB_BEFAB_12_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3
		return
	case "ROTARY_RHAG_MK_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1
		return
	case "ROTARY_HYDRAULIC_WATER":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER
		return
	case "BARRIER_DISK_SAFELAND":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND
		return
	case "LINEAR_SPRAG_MK_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codearrestinggearenergyabsorbbasetype_member_1 *Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND
	case "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1":
		*enum_codearrestinggearenergyabsorbbasetype_member_1 = Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codearrestinggearenergyabsorbbasetype_member_1 *Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codearrestinggearenergyabsorbbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND"
	case Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1:
		res = "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1"
	}
	return
}

func (enum_codearrestinggearenergyabsorbbasetype_member_1 Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_1300")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_2800")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1A")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1B")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34B_1C")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_34D_1F")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2C")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2D")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2E")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2F")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2H")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2I")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_2L")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3A")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3H")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_3L")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4C")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4E")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_44B_4H")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_4")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_6")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_500S_8")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARYTRANS_500S_8")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_AAE_64")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12A")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_12B")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_13")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_BAK_6")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_BAK_9")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_12_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_20_4")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_21_2")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_24_4")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_56_2")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_6_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_60_2")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_BEFAB_8_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_CHAG")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_DUAL_BAK_12")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E15")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E27")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_E28")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_1")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_2")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E5_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_CHAIN_E6")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_CHAIN_JETSTOP")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_M21")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAAS")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_I")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_II")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_III")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IV")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_IX")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VI")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VII")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_VIII")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_MAG_X")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_100")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_TEXTILE_MB_60")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_MOBILROTARY_HYDRAULIC_WATER")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_PUAG_MK_21")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_12A")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_MK_6")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_RAF_PORTABLE_AGEAR")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEA_BEFAB_6_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_DISK_RAF_TYPEB_BEFAB_12_3")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_RHAG_MK_1")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_ROTARY_HYDRAULIC_WATER")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_BARRIER_DISK_SAFELAND")
	res = append(res, "Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1_LINEAR_SPRAG_MK_1")

	return
}

func (enum_codearrestinggearenergyabsorbbasetype_member_1 Enum_CodeArrestingGearEnergyAbsorbBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ROTARY_1300")
	res = append(res, "ROTARY_2800")
	res = append(res, "ROTARY_34B_1A")
	res = append(res, "ROTARY_34B_1B")
	res = append(res, "ROTARY_34B_1C")
	res = append(res, "ROTARY_34D_1F")
	res = append(res, "ROTARY_44B_2C")
	res = append(res, "ROTARY_44B_2D")
	res = append(res, "ROTARY_44B_2E")
	res = append(res, "ROTARY_44B_2F")
	res = append(res, "ROTARY_44B_2H")
	res = append(res, "ROTARY_44B_2I")
	res = append(res, "ROTARY_44B_2L")
	res = append(res, "ROTARY_44B_3A")
	res = append(res, "ROTARY_44B_3H")
	res = append(res, "ROTARY_44B_3L")
	res = append(res, "ROTARY_44B_4C")
	res = append(res, "ROTARY_44B_4E")
	res = append(res, "ROTARY_44B_4H")
	res = append(res, "ROTARY_500S")
	res = append(res, "ROTARY_500S_4")
	res = append(res, "ROTARY_500S_6")
	res = append(res, "ROTARY_500S_8")
	res = append(res, "ROTARYTRANS_500S_8")
	res = append(res, "ROTARY_AAE_64")
	res = append(res, "ROTARY_BAK_12A")
	res = append(res, "ROTARY_BAK_12B")
	res = append(res, "ROTARY_BAK_13")
	res = append(res, "LINEAR_BAK_6")
	res = append(res, "ROTARY_BAK_9")
	res = append(res, "DISK_BEFAB_12_3")
	res = append(res, "DISK_BEFAB_20_4")
	res = append(res, "DISK_BEFAB_21_2")
	res = append(res, "DISK_BEFAB_24_4")
	res = append(res, "DISK_BEFAB_56_2")
	res = append(res, "DISK_BEFAB_6_3")
	res = append(res, "DISK_BEFAB_60_2")
	res = append(res, "DISK_BEFAB_8_3")
	res = append(res, "CHAIN_CHAG")
	res = append(res, "ROTARY_DUAL_BAK_12")
	res = append(res, "ROTARY_E15")
	res = append(res, "ROTARY_E27")
	res = append(res, "ROTARY_E28")
	res = append(res, "CHAIN_E5")
	res = append(res, "CHAIN_E5_1")
	res = append(res, "CHAIN_E5_2")
	res = append(res, "CHAIN_E5_3")
	res = append(res, "CHAIN_E6")
	res = append(res, "ROTARY_CHAIN_JETSTOP")
	res = append(res, "MOBILROTARY_M21")
	res = append(res, "MOBILROTARY_MAAS")
	res = append(res, "MOBILROTARY_MAG_I")
	res = append(res, "MOBILROTARY_MAG_II")
	res = append(res, "MOBILROTARY_MAG_III")
	res = append(res, "MOBILROTARY_MAG_IV")
	res = append(res, "MOBILROTARY_MAG_IX")
	res = append(res, "MOBILROTARY_MAG_VI")
	res = append(res, "MOBILROTARY_MAG_VII")
	res = append(res, "MOBILROTARY_MAG_VIII")
	res = append(res, "MOBILROTARY_MAG_X")
	res = append(res, "TEXTILE_MB_100")
	res = append(res, "TEXTILE_MB_60")
	res = append(res, "MOBILROTARY_HYDRAULIC_WATER")
	res = append(res, "ROTARY_PUAG_MK_21")
	res = append(res, "DISK_RAF_MK_12A")
	res = append(res, "DISK_RAF_MK_6")
	res = append(res, "RAF_PORTABLE_AGEAR")
	res = append(res, "DISK_RAF_TYPEA_BEFAB_6_3")
	res = append(res, "DISK_RAF_TYPEB_BEFAB_12_3")
	res = append(res, "ROTARY_RHAG_MK_1")
	res = append(res, "ROTARY_HYDRAULIC_WATER")
	res = append(res, "BARRIER_DISK_SAFELAND")
	res = append(res, "LINEAR_SPRAG_MK_1")

	return
}

// Utility function for Enum_CodeArrestingGearEngageDeviceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codearrestinggearengagedevicebasetype_member_1 Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codearrestinggearengagedevicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII:
		res = "61QSII"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI:
		res = "62NI"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI:
		res = "63PI"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30:
		res = "NET_A30"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40:
		res = "NET_A40"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT:
		res = "BAK_11_STRUT"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12:
		res = "BAK_12"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK:
		res = "BAK_14_HOOK"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET:
		res = "BAK_15_STANCHION_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK:
		res = "BAK_15_HOOK"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS:
		res = "EMAS"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE:
		res = "HOOK_CABLE"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET:
		res = "HP_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR:
		res = "J_BAR"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER:
		res = "JET_BARRIER"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET:
		res = "MA_1_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE:
		res = "MA_1A_HOOK_CABLE"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET:
		res = "NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H:
		res = "HOOK_H"
	}
	return
}

func (enum_codearrestinggearengagedevicebasetype_member_1 *Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "61QSII":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII
		return
	case "62NI":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI
		return
	case "63PI":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI
		return
	case "NET_A30":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30
		return
	case "NET_A40":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40
		return
	case "BAK_11_STRUT":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT
		return
	case "BAK_12":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12
		return
	case "BAK_14_HOOK":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK
		return
	case "BAK_15_STANCHION_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET
		return
	case "BAK_15_HOOK":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK
		return
	case "EMAS":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS
		return
	case "HOOK_CABLE":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE
		return
	case "HP_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET
		return
	case "J_BAR":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR
		return
	case "JET_BARRIER":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER
		return
	case "MA_1_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET
		return
	case "MA_1A_HOOK_CABLE":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE
		return
	case "NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET
		return
	case "HOOK_H":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codearrestinggearengagedevicebasetype_member_1 *Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET
	case "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H":
		*enum_codearrestinggearengagedevicebasetype_member_1 = Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codearrestinggearengagedevicebasetype_member_1 *Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codearrestinggearengagedevicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET"
	case Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H:
		res = "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H"
	}
	return
}

func (enum_codearrestinggearengagedevicebasetype_member_1 Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_61QSII")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_62NI")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_N_63PI")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A30")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET_A40")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_11_STRUT")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_12")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_14_HOOK")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_STANCHION_NET")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_BAK_15_HOOK")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_EMAS")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_CABLE")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HP_NET")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_J_BAR")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_JET_BARRIER")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1_NET")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_MA_1A_HOOK_CABLE")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_NET")
	res = append(res, "Enum_CodeArrestingGearEngageDeviceBaseType_Member_1_HOOK_H")

	return
}

func (enum_codearrestinggearengagedevicebasetype_member_1 Enum_CodeArrestingGearEngageDeviceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "61QSII")
	res = append(res, "62NI")
	res = append(res, "63PI")
	res = append(res, "NET_A30")
	res = append(res, "NET_A40")
	res = append(res, "BAK_11_STRUT")
	res = append(res, "BAK_12")
	res = append(res, "BAK_14_HOOK")
	res = append(res, "BAK_15_STANCHION_NET")
	res = append(res, "BAK_15_HOOK")
	res = append(res, "EMAS")
	res = append(res, "HOOK_CABLE")
	res = append(res, "HP_NET")
	res = append(res, "J_BAR")
	res = append(res, "JET_BARRIER")
	res = append(res, "MA_1_NET")
	res = append(res, "MA_1A_HOOK_CABLE")
	res = append(res, "NET")
	res = append(res, "HOOK_H")

	return
}

// Utility function for Enum_CodeAuthorityRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeauthorityrolebasetype_member_1 Enum_CodeAuthorityRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeauthorityrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAuthorityRoleBaseType_Member_1_OWN:
		res = "OWN"
	case Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE:
		res = "OPERATE"
	case Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE:
		res = "SUPERVISE"
	}
	return
}

func (enum_codeauthorityrolebasetype_member_1 *Enum_CodeAuthorityRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OWN":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_OWN
		return
	case "OPERATE":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE
		return
	case "SUPERVISE":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeauthorityrolebasetype_member_1 *Enum_CodeAuthorityRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAuthorityRoleBaseType_Member_1_OWN":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_OWN
	case "Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE
	case "Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE":
		*enum_codeauthorityrolebasetype_member_1 = Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeauthorityrolebasetype_member_1 *Enum_CodeAuthorityRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeauthorityrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAuthorityRoleBaseType_Member_1_OWN:
		res = "Enum_CodeAuthorityRoleBaseType_Member_1_OWN"
	case Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE:
		res = "Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE"
	case Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE:
		res = "Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE"
	}
	return
}

func (enum_codeauthorityrolebasetype_member_1 Enum_CodeAuthorityRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAuthorityRoleBaseType_Member_1_OWN")
	res = append(res, "Enum_CodeAuthorityRoleBaseType_Member_1_OPERATE")
	res = append(res, "Enum_CodeAuthorityRoleBaseType_Member_1_SUPERVISE")

	return
}

func (enum_codeauthorityrolebasetype_member_1 Enum_CodeAuthorityRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OWN")
	res = append(res, "OPERATE")
	res = append(res, "SUPERVISE")

	return
}

// Utility function for Enum_CodeBuoyBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codebuoybasetype_member_1 Enum_CodeBuoyBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codebuoybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2:
		res = "BLACK_RED_FL2"
	case Enum_CodeBuoyBaseType_Member_1_GREEN:
		res = "GREEN"
	case Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL:
		res = "GREEN_RED_GFL"
	case Enum_CodeBuoyBaseType_Member_1_Q3_VQ3:
		res = "Q3_VQ3"
	case Enum_CodeBuoyBaseType_Member_1_Q6_VQ6:
		res = "Q6_VQ6"
	case Enum_CodeBuoyBaseType_Member_1_Q9_VQ9:
		res = "Q9_VQ9"
	case Enum_CodeBuoyBaseType_Member_1_Q_VQ:
		res = "Q_VQ"
	case Enum_CodeBuoyBaseType_Member_1_RED:
		res = "RED"
	case Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL:
		res = "RED_GREEN_RFL"
	case Enum_CodeBuoyBaseType_Member_1_RED_WHITE:
		res = "RED_WHITE"
	case Enum_CodeBuoyBaseType_Member_1_WHITE:
		res = "WHITE"
	case Enum_CodeBuoyBaseType_Member_1_YELLOW:
		res = "YELLOW"
	}
	return
}

func (enum_codebuoybasetype_member_1 *Enum_CodeBuoyBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BLACK_RED_FL2":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2
		return
	case "GREEN":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_GREEN
		return
	case "GREEN_RED_GFL":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL
		return
	case "Q3_VQ3":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q3_VQ3
		return
	case "Q6_VQ6":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q6_VQ6
		return
	case "Q9_VQ9":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q9_VQ9
		return
	case "Q_VQ":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q_VQ
		return
	case "RED":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED
		return
	case "RED_GREEN_RFL":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL
		return
	case "RED_WHITE":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED_WHITE
		return
	case "WHITE":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_WHITE
		return
	case "YELLOW":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_YELLOW
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codebuoybasetype_member_1 *Enum_CodeBuoyBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2
	case "Enum_CodeBuoyBaseType_Member_1_GREEN":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_GREEN
	case "Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL
	case "Enum_CodeBuoyBaseType_Member_1_Q3_VQ3":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q3_VQ3
	case "Enum_CodeBuoyBaseType_Member_1_Q6_VQ6":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q6_VQ6
	case "Enum_CodeBuoyBaseType_Member_1_Q9_VQ9":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q9_VQ9
	case "Enum_CodeBuoyBaseType_Member_1_Q_VQ":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_Q_VQ
	case "Enum_CodeBuoyBaseType_Member_1_RED":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED
	case "Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL
	case "Enum_CodeBuoyBaseType_Member_1_RED_WHITE":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_RED_WHITE
	case "Enum_CodeBuoyBaseType_Member_1_WHITE":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_WHITE
	case "Enum_CodeBuoyBaseType_Member_1_YELLOW":
		*enum_codebuoybasetype_member_1 = Enum_CodeBuoyBaseType_Member_1_YELLOW
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codebuoybasetype_member_1 *Enum_CodeBuoyBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codebuoybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2:
		res = "Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2"
	case Enum_CodeBuoyBaseType_Member_1_GREEN:
		res = "Enum_CodeBuoyBaseType_Member_1_GREEN"
	case Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL:
		res = "Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL"
	case Enum_CodeBuoyBaseType_Member_1_Q3_VQ3:
		res = "Enum_CodeBuoyBaseType_Member_1_Q3_VQ3"
	case Enum_CodeBuoyBaseType_Member_1_Q6_VQ6:
		res = "Enum_CodeBuoyBaseType_Member_1_Q6_VQ6"
	case Enum_CodeBuoyBaseType_Member_1_Q9_VQ9:
		res = "Enum_CodeBuoyBaseType_Member_1_Q9_VQ9"
	case Enum_CodeBuoyBaseType_Member_1_Q_VQ:
		res = "Enum_CodeBuoyBaseType_Member_1_Q_VQ"
	case Enum_CodeBuoyBaseType_Member_1_RED:
		res = "Enum_CodeBuoyBaseType_Member_1_RED"
	case Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL:
		res = "Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL"
	case Enum_CodeBuoyBaseType_Member_1_RED_WHITE:
		res = "Enum_CodeBuoyBaseType_Member_1_RED_WHITE"
	case Enum_CodeBuoyBaseType_Member_1_WHITE:
		res = "Enum_CodeBuoyBaseType_Member_1_WHITE"
	case Enum_CodeBuoyBaseType_Member_1_YELLOW:
		res = "Enum_CodeBuoyBaseType_Member_1_YELLOW"
	}
	return
}

func (enum_codebuoybasetype_member_1 Enum_CodeBuoyBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_BLACK_RED_FL2")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_GREEN")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_GREEN_RED_GFL")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_Q3_VQ3")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_Q6_VQ6")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_Q9_VQ9")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_Q_VQ")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_RED")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_RED_GREEN_RFL")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_RED_WHITE")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_WHITE")
	res = append(res, "Enum_CodeBuoyBaseType_Member_1_YELLOW")

	return
}

func (enum_codebuoybasetype_member_1 Enum_CodeBuoyBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BLACK_RED_FL2")
	res = append(res, "GREEN")
	res = append(res, "GREEN_RED_GFL")
	res = append(res, "Q3_VQ3")
	res = append(res, "Q6_VQ6")
	res = append(res, "Q9_VQ9")
	res = append(res, "Q_VQ")
	res = append(res, "RED")
	res = append(res, "RED_GREEN_RFL")
	res = append(res, "RED_WHITE")
	res = append(res, "WHITE")
	res = append(res, "YELLOW")

	return
}

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

// Utility function for Enum_CodeCommunicationModeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecommunicationmodebasetype_member_1 Enum_CodeCommunicationModeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecommunicationmodebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCommunicationModeBaseType_Member_1_HF:
		res = "HF"
	case Enum_CodeCommunicationModeBaseType_Member_1_VHF:
		res = "VHF"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL1:
		res = "VDL1"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL2:
		res = "VDL2"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL4:
		res = "VDL4"
	case Enum_CodeCommunicationModeBaseType_Member_1_AMSS:
		res = "AMSS"
	case Enum_CodeCommunicationModeBaseType_Member_1_ADS_B:
		res = "ADS_B"
	case Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL:
		res = "ADS_B_VDL"
	case Enum_CodeCommunicationModeBaseType_Member_1_HFDL:
		res = "HFDL"
	case Enum_CodeCommunicationModeBaseType_Member_1_VHF_833:
		res = "VHF_833"
	case Enum_CodeCommunicationModeBaseType_Member_1_UHF:
		res = "UHF"
	}
	return
}

func (enum_codecommunicationmodebasetype_member_1 *Enum_CodeCommunicationModeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_HF
		return
	case "VHF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VHF
		return
	case "VDL1":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL1
		return
	case "VDL2":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL2
		return
	case "VDL4":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL4
		return
	case "AMSS":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_AMSS
		return
	case "ADS_B":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_ADS_B
		return
	case "ADS_B_VDL":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL
		return
	case "HFDL":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_HFDL
		return
	case "VHF_833":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VHF_833
		return
	case "UHF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_UHF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecommunicationmodebasetype_member_1 *Enum_CodeCommunicationModeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCommunicationModeBaseType_Member_1_HF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_HF
	case "Enum_CodeCommunicationModeBaseType_Member_1_VHF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VHF
	case "Enum_CodeCommunicationModeBaseType_Member_1_VDL1":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL1
	case "Enum_CodeCommunicationModeBaseType_Member_1_VDL2":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL2
	case "Enum_CodeCommunicationModeBaseType_Member_1_VDL4":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VDL4
	case "Enum_CodeCommunicationModeBaseType_Member_1_AMSS":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_AMSS
	case "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_ADS_B
	case "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL
	case "Enum_CodeCommunicationModeBaseType_Member_1_HFDL":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_HFDL
	case "Enum_CodeCommunicationModeBaseType_Member_1_VHF_833":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_VHF_833
	case "Enum_CodeCommunicationModeBaseType_Member_1_UHF":
		*enum_codecommunicationmodebasetype_member_1 = Enum_CodeCommunicationModeBaseType_Member_1_UHF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecommunicationmodebasetype_member_1 *Enum_CodeCommunicationModeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecommunicationmodebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCommunicationModeBaseType_Member_1_HF:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_HF"
	case Enum_CodeCommunicationModeBaseType_Member_1_VHF:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_VHF"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL1:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_VDL1"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL2:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_VDL2"
	case Enum_CodeCommunicationModeBaseType_Member_1_VDL4:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_VDL4"
	case Enum_CodeCommunicationModeBaseType_Member_1_AMSS:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_AMSS"
	case Enum_CodeCommunicationModeBaseType_Member_1_ADS_B:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B"
	case Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL"
	case Enum_CodeCommunicationModeBaseType_Member_1_HFDL:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_HFDL"
	case Enum_CodeCommunicationModeBaseType_Member_1_VHF_833:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_VHF_833"
	case Enum_CodeCommunicationModeBaseType_Member_1_UHF:
		res = "Enum_CodeCommunicationModeBaseType_Member_1_UHF"
	}
	return
}

func (enum_codecommunicationmodebasetype_member_1 Enum_CodeCommunicationModeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_HF")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_VHF")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_VDL1")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_VDL2")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_VDL4")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_AMSS")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_ADS_B_VDL")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_HFDL")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_VHF_833")
	res = append(res, "Enum_CodeCommunicationModeBaseType_Member_1_UHF")

	return
}

func (enum_codecommunicationmodebasetype_member_1 Enum_CodeCommunicationModeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HF")
	res = append(res, "VHF")
	res = append(res, "VDL1")
	res = append(res, "VDL2")
	res = append(res, "VDL4")
	res = append(res, "AMSS")
	res = append(res, "ADS_B")
	res = append(res, "ADS_B_VDL")
	res = append(res, "HFDL")
	res = append(res, "VHF_833")
	res = append(res, "UHF")

	return
}

// Utility function for Enum_CodeContaminationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecontaminationbasetype_member_1 Enum_CodeContaminationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecontaminationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeContaminationBaseType_Member_1_NONE:
		res = "NONE"
	case Enum_CodeContaminationBaseType_Member_1_DAMP:
		res = "DAMP"
	case Enum_CodeContaminationBaseType_Member_1_WATER:
		res = "WATER"
	case Enum_CodeContaminationBaseType_Member_1_FROST:
		res = "FROST"
	case Enum_CodeContaminationBaseType_Member_1_DRY_SNOW:
		res = "DRY_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_WET_SNOW:
		res = "WET_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_SLUSH:
		res = "SLUSH"
	case Enum_CodeContaminationBaseType_Member_1_ICE:
		res = "ICE"
	case Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW:
		res = "COMPACT_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_RUT:
		res = "RUT"
	case Enum_CodeContaminationBaseType_Member_1_ASH:
		res = "ASH"
	case Enum_CodeContaminationBaseType_Member_1_SAND:
		res = "SAND"
	case Enum_CodeContaminationBaseType_Member_1_OIL:
		res = "OIL"
	case Enum_CodeContaminationBaseType_Member_1_RUBBER:
		res = "RUBBER"
	case Enum_CodeContaminationBaseType_Member_1_GRAS:
		res = "GRAS"
	}
	return
}

func (enum_codecontaminationbasetype_member_1 *Enum_CodeContaminationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NONE":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_NONE
		return
	case "DAMP":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_DAMP
		return
	case "WATER":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_WATER
		return
	case "FROST":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_FROST
		return
	case "DRY_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_DRY_SNOW
		return
	case "WET_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_WET_SNOW
		return
	case "SLUSH":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_SLUSH
		return
	case "ICE":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_ICE
		return
	case "COMPACT_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW
		return
	case "RUT":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_RUT
		return
	case "ASH":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_ASH
		return
	case "SAND":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_SAND
		return
	case "OIL":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_OIL
		return
	case "RUBBER":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_RUBBER
		return
	case "GRAS":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_GRAS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecontaminationbasetype_member_1 *Enum_CodeContaminationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeContaminationBaseType_Member_1_NONE":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_NONE
	case "Enum_CodeContaminationBaseType_Member_1_DAMP":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_DAMP
	case "Enum_CodeContaminationBaseType_Member_1_WATER":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_WATER
	case "Enum_CodeContaminationBaseType_Member_1_FROST":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_FROST
	case "Enum_CodeContaminationBaseType_Member_1_DRY_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_DRY_SNOW
	case "Enum_CodeContaminationBaseType_Member_1_WET_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_WET_SNOW
	case "Enum_CodeContaminationBaseType_Member_1_SLUSH":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_SLUSH
	case "Enum_CodeContaminationBaseType_Member_1_ICE":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_ICE
	case "Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW
	case "Enum_CodeContaminationBaseType_Member_1_RUT":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_RUT
	case "Enum_CodeContaminationBaseType_Member_1_ASH":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_ASH
	case "Enum_CodeContaminationBaseType_Member_1_SAND":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_SAND
	case "Enum_CodeContaminationBaseType_Member_1_OIL":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_OIL
	case "Enum_CodeContaminationBaseType_Member_1_RUBBER":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_RUBBER
	case "Enum_CodeContaminationBaseType_Member_1_GRAS":
		*enum_codecontaminationbasetype_member_1 = Enum_CodeContaminationBaseType_Member_1_GRAS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecontaminationbasetype_member_1 *Enum_CodeContaminationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecontaminationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeContaminationBaseType_Member_1_NONE:
		res = "Enum_CodeContaminationBaseType_Member_1_NONE"
	case Enum_CodeContaminationBaseType_Member_1_DAMP:
		res = "Enum_CodeContaminationBaseType_Member_1_DAMP"
	case Enum_CodeContaminationBaseType_Member_1_WATER:
		res = "Enum_CodeContaminationBaseType_Member_1_WATER"
	case Enum_CodeContaminationBaseType_Member_1_FROST:
		res = "Enum_CodeContaminationBaseType_Member_1_FROST"
	case Enum_CodeContaminationBaseType_Member_1_DRY_SNOW:
		res = "Enum_CodeContaminationBaseType_Member_1_DRY_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_WET_SNOW:
		res = "Enum_CodeContaminationBaseType_Member_1_WET_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_SLUSH:
		res = "Enum_CodeContaminationBaseType_Member_1_SLUSH"
	case Enum_CodeContaminationBaseType_Member_1_ICE:
		res = "Enum_CodeContaminationBaseType_Member_1_ICE"
	case Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW:
		res = "Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW"
	case Enum_CodeContaminationBaseType_Member_1_RUT:
		res = "Enum_CodeContaminationBaseType_Member_1_RUT"
	case Enum_CodeContaminationBaseType_Member_1_ASH:
		res = "Enum_CodeContaminationBaseType_Member_1_ASH"
	case Enum_CodeContaminationBaseType_Member_1_SAND:
		res = "Enum_CodeContaminationBaseType_Member_1_SAND"
	case Enum_CodeContaminationBaseType_Member_1_OIL:
		res = "Enum_CodeContaminationBaseType_Member_1_OIL"
	case Enum_CodeContaminationBaseType_Member_1_RUBBER:
		res = "Enum_CodeContaminationBaseType_Member_1_RUBBER"
	case Enum_CodeContaminationBaseType_Member_1_GRAS:
		res = "Enum_CodeContaminationBaseType_Member_1_GRAS"
	}
	return
}

func (enum_codecontaminationbasetype_member_1 Enum_CodeContaminationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_NONE")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_DAMP")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_WATER")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_FROST")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_DRY_SNOW")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_WET_SNOW")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_SLUSH")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_ICE")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_COMPACT_SNOW")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_RUT")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_ASH")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_SAND")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_OIL")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_RUBBER")
	res = append(res, "Enum_CodeContaminationBaseType_Member_1_GRAS")

	return
}

func (enum_codecontaminationbasetype_member_1 Enum_CodeContaminationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NONE")
	res = append(res, "DAMP")
	res = append(res, "WATER")
	res = append(res, "FROST")
	res = append(res, "DRY_SNOW")
	res = append(res, "WET_SNOW")
	res = append(res, "SLUSH")
	res = append(res, "ICE")
	res = append(res, "COMPACT_SNOW")
	res = append(res, "RUT")
	res = append(res, "ASH")
	res = append(res, "SAND")
	res = append(res, "OIL")
	res = append(res, "RUBBER")
	res = append(res, "GRAS")

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

// Utility function for Enum_CodeDeclaredDistanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedeclareddistancebasetype_member_1 Enum_CodeDeclaredDistanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedeclareddistancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDeclaredDistanceBaseType_Member_1_LDA:
		res = "LDA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TORA:
		res = "TORA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TODA:
		res = "TODA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA:
		res = "ASDA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR:
		res = "DTHR"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH:
		res = "TODAH"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH:
		res = "RTODAH"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH:
		res = "LDAH"
	}
	return
}

func (enum_codedeclareddistancebasetype_member_1 *Enum_CodeDeclaredDistanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LDA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_LDA
		return
	case "TORA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TORA
		return
	case "TODA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TODA
		return
	case "ASDA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA
		return
	case "DTHR":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR
		return
	case "TODAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH
		return
	case "RTODAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH
		return
	case "LDAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedeclareddistancebasetype_member_1 *Enum_CodeDeclaredDistanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_LDA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_LDA
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_TORA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TORA
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_TODA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TODA
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH
	case "Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH":
		*enum_codedeclareddistancebasetype_member_1 = Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedeclareddistancebasetype_member_1 *Enum_CodeDeclaredDistanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedeclareddistancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDeclaredDistanceBaseType_Member_1_LDA:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_LDA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TORA:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_TORA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TODA:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_TODA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH"
	case Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH:
		res = "Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH"
	}
	return
}

func (enum_codedeclareddistancebasetype_member_1 Enum_CodeDeclaredDistanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_LDA")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_TORA")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_TODA")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_ASDA")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_DTHR")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_TODAH")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_RTODAH")
	res = append(res, "Enum_CodeDeclaredDistanceBaseType_Member_1_LDAH")

	return
}

func (enum_codedeclareddistancebasetype_member_1 Enum_CodeDeclaredDistanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LDA")
	res = append(res, "TORA")
	res = append(res, "TODA")
	res = append(res, "ASDA")
	res = append(res, "DTHR")
	res = append(res, "TODAH")
	res = append(res, "RTODAH")
	res = append(res, "LDAH")

	return
}

// Utility function for Enum_CodeDirectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedirectionbasetype_member_1 Enum_CodeDirectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionBaseType_Member_1_FORWARD:
		res = "FORWARD"
	case Enum_CodeDirectionBaseType_Member_1_BACKWARD:
		res = "BACKWARD"
	case Enum_CodeDirectionBaseType_Member_1_BOTH:
		res = "BOTH"
	}
	return
}

func (enum_codedirectionbasetype_member_1 *Enum_CodeDirectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FORWARD":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_FORWARD
		return
	case "BACKWARD":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_BACKWARD
		return
	case "BOTH":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_BOTH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedirectionbasetype_member_1 *Enum_CodeDirectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDirectionBaseType_Member_1_FORWARD":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_FORWARD
	case "Enum_CodeDirectionBaseType_Member_1_BACKWARD":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_BACKWARD
	case "Enum_CodeDirectionBaseType_Member_1_BOTH":
		*enum_codedirectionbasetype_member_1 = Enum_CodeDirectionBaseType_Member_1_BOTH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedirectionbasetype_member_1 *Enum_CodeDirectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionBaseType_Member_1_FORWARD:
		res = "Enum_CodeDirectionBaseType_Member_1_FORWARD"
	case Enum_CodeDirectionBaseType_Member_1_BACKWARD:
		res = "Enum_CodeDirectionBaseType_Member_1_BACKWARD"
	case Enum_CodeDirectionBaseType_Member_1_BOTH:
		res = "Enum_CodeDirectionBaseType_Member_1_BOTH"
	}
	return
}

func (enum_codedirectionbasetype_member_1 Enum_CodeDirectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDirectionBaseType_Member_1_FORWARD")
	res = append(res, "Enum_CodeDirectionBaseType_Member_1_BACKWARD")
	res = append(res, "Enum_CodeDirectionBaseType_Member_1_BOTH")

	return
}

func (enum_codedirectionbasetype_member_1 Enum_CodeDirectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FORWARD")
	res = append(res, "BACKWARD")
	res = append(res, "BOTH")

	return
}

// Utility function for Enum_CodeDirectionTurnBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedirectionturnbasetype_member_1 Enum_CodeDirectionTurnBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedirectionturnbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionTurnBaseType_Member_1_LEFT:
		res = "LEFT"
	case Enum_CodeDirectionTurnBaseType_Member_1_RIGHT:
		res = "RIGHT"
	case Enum_CodeDirectionTurnBaseType_Member_1_EITHER:
		res = "EITHER"
	}
	return
}

func (enum_codedirectionturnbasetype_member_1 *Enum_CodeDirectionTurnBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LEFT":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_LEFT
		return
	case "RIGHT":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_RIGHT
		return
	case "EITHER":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_EITHER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedirectionturnbasetype_member_1 *Enum_CodeDirectionTurnBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDirectionTurnBaseType_Member_1_LEFT":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_LEFT
	case "Enum_CodeDirectionTurnBaseType_Member_1_RIGHT":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_RIGHT
	case "Enum_CodeDirectionTurnBaseType_Member_1_EITHER":
		*enum_codedirectionturnbasetype_member_1 = Enum_CodeDirectionTurnBaseType_Member_1_EITHER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedirectionturnbasetype_member_1 *Enum_CodeDirectionTurnBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedirectionturnbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionTurnBaseType_Member_1_LEFT:
		res = "Enum_CodeDirectionTurnBaseType_Member_1_LEFT"
	case Enum_CodeDirectionTurnBaseType_Member_1_RIGHT:
		res = "Enum_CodeDirectionTurnBaseType_Member_1_RIGHT"
	case Enum_CodeDirectionTurnBaseType_Member_1_EITHER:
		res = "Enum_CodeDirectionTurnBaseType_Member_1_EITHER"
	}
	return
}

func (enum_codedirectionturnbasetype_member_1 Enum_CodeDirectionTurnBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDirectionTurnBaseType_Member_1_LEFT")
	res = append(res, "Enum_CodeDirectionTurnBaseType_Member_1_RIGHT")
	res = append(res, "Enum_CodeDirectionTurnBaseType_Member_1_EITHER")

	return
}

func (enum_codedirectionturnbasetype_member_1 Enum_CodeDirectionTurnBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LEFT")
	res = append(res, "RIGHT")
	res = append(res, "EITHER")

	return
}

// Utility function for Enum_CodeEquipmentAntiCollisionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeequipmentanticollisionbasetype_member_1 Enum_CodeEquipmentAntiCollisionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeequipmentanticollisionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I:
		res = "ACAS_I"
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II:
		res = "ACAS_II"
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS:
		res = "GPWS"
	}
	return
}

func (enum_codeequipmentanticollisionbasetype_member_1 *Enum_CodeEquipmentAntiCollisionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ACAS_I":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I
		return
	case "ACAS_II":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II
		return
	case "GPWS":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeequipmentanticollisionbasetype_member_1 *Enum_CodeEquipmentAntiCollisionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I
	case "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II
	case "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS":
		*enum_codeequipmentanticollisionbasetype_member_1 = Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeequipmentanticollisionbasetype_member_1 *Enum_CodeEquipmentAntiCollisionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeequipmentanticollisionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I:
		res = "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I"
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II:
		res = "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II"
	case Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS:
		res = "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS"
	}
	return
}

func (enum_codeequipmentanticollisionbasetype_member_1 Enum_CodeEquipmentAntiCollisionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_I")
	res = append(res, "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_ACAS_II")
	res = append(res, "Enum_CodeEquipmentAntiCollisionBaseType_Member_1_GPWS")

	return
}

func (enum_codeequipmentanticollisionbasetype_member_1 Enum_CodeEquipmentAntiCollisionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ACAS_I")
	res = append(res, "ACAS_II")
	res = append(res, "GPWS")

	return
}

// Utility function for Enum_CodeFlightBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightbasetype_member_1 Enum_CodeFlightBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightBaseType_Member_1_OAT:
		res = "OAT"
	case Enum_CodeFlightBaseType_Member_1_GAT:
		res = "GAT"
	case Enum_CodeFlightBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeflightbasetype_member_1 *Enum_CodeFlightBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OAT":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_OAT
		return
	case "GAT":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_GAT
		return
	case "ALL":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightbasetype_member_1 *Enum_CodeFlightBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightBaseType_Member_1_OAT":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_OAT
	case "Enum_CodeFlightBaseType_Member_1_GAT":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_GAT
	case "Enum_CodeFlightBaseType_Member_1_ALL":
		*enum_codeflightbasetype_member_1 = Enum_CodeFlightBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightbasetype_member_1 *Enum_CodeFlightBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightBaseType_Member_1_OAT:
		res = "Enum_CodeFlightBaseType_Member_1_OAT"
	case Enum_CodeFlightBaseType_Member_1_GAT:
		res = "Enum_CodeFlightBaseType_Member_1_GAT"
	case Enum_CodeFlightBaseType_Member_1_ALL:
		res = "Enum_CodeFlightBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeflightbasetype_member_1 Enum_CodeFlightBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightBaseType_Member_1_OAT")
	res = append(res, "Enum_CodeFlightBaseType_Member_1_GAT")
	res = append(res, "Enum_CodeFlightBaseType_Member_1_ALL")

	return
}

func (enum_codeflightbasetype_member_1 Enum_CodeFlightBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OAT")
	res = append(res, "GAT")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeFlightOriginBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightoriginbasetype_member_1 Enum_CodeFlightOriginBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightoriginbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightOriginBaseType_Member_1_NTL:
		res = "NTL"
	case Enum_CodeFlightOriginBaseType_Member_1_INTL:
		res = "INTL"
	case Enum_CodeFlightOriginBaseType_Member_1_ALL:
		res = "ALL"
	case Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED:
		res = "HOME_BASED"
	}
	return
}

func (enum_codeflightoriginbasetype_member_1 *Enum_CodeFlightOriginBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NTL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_NTL
		return
	case "INTL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_INTL
		return
	case "ALL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_ALL
		return
	case "HOME_BASED":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightoriginbasetype_member_1 *Enum_CodeFlightOriginBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightOriginBaseType_Member_1_NTL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_NTL
	case "Enum_CodeFlightOriginBaseType_Member_1_INTL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_INTL
	case "Enum_CodeFlightOriginBaseType_Member_1_ALL":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_ALL
	case "Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED":
		*enum_codeflightoriginbasetype_member_1 = Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightoriginbasetype_member_1 *Enum_CodeFlightOriginBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightoriginbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightOriginBaseType_Member_1_NTL:
		res = "Enum_CodeFlightOriginBaseType_Member_1_NTL"
	case Enum_CodeFlightOriginBaseType_Member_1_INTL:
		res = "Enum_CodeFlightOriginBaseType_Member_1_INTL"
	case Enum_CodeFlightOriginBaseType_Member_1_ALL:
		res = "Enum_CodeFlightOriginBaseType_Member_1_ALL"
	case Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED:
		res = "Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED"
	}
	return
}

func (enum_codeflightoriginbasetype_member_1 Enum_CodeFlightOriginBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightOriginBaseType_Member_1_NTL")
	res = append(res, "Enum_CodeFlightOriginBaseType_Member_1_INTL")
	res = append(res, "Enum_CodeFlightOriginBaseType_Member_1_ALL")
	res = append(res, "Enum_CodeFlightOriginBaseType_Member_1_HOME_BASED")

	return
}

func (enum_codeflightoriginbasetype_member_1 Enum_CodeFlightOriginBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NTL")
	res = append(res, "INTL")
	res = append(res, "ALL")
	res = append(res, "HOME_BASED")

	return
}

// Utility function for Enum_CodeFlightPurposeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightpurposebasetype_member_1 Enum_CodeFlightPurposeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightpurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED:
		res = "SCHEDULED"
	case Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED:
		res = "NON_SCHEDULED"
	case Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE:
		res = "PRIVATE"
	case Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING:
		res = "AIR_TRAINING"
	case Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK:
		res = "AIR_WORK"
	case Enum_CodeFlightPurposeBaseType_Member_1_ALL:
		res = "ALL"
	case Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT:
		res = "PARTICIPANT"
	}
	return
}

func (enum_codeflightpurposebasetype_member_1 *Enum_CodeFlightPurposeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SCHEDULED":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED
		return
	case "NON_SCHEDULED":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED
		return
	case "PRIVATE":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE
		return
	case "AIR_TRAINING":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING
		return
	case "AIR_WORK":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK
		return
	case "ALL":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_ALL
		return
	case "PARTICIPANT":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightpurposebasetype_member_1 *Enum_CodeFlightPurposeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED
	case "Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED
	case "Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE
	case "Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING
	case "Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK
	case "Enum_CodeFlightPurposeBaseType_Member_1_ALL":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_ALL
	case "Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT":
		*enum_codeflightpurposebasetype_member_1 = Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightpurposebasetype_member_1 *Enum_CodeFlightPurposeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightpurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED"
	case Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED"
	case Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE"
	case Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING"
	case Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK"
	case Enum_CodeFlightPurposeBaseType_Member_1_ALL:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_ALL"
	case Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT:
		res = "Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT"
	}
	return
}

func (enum_codeflightpurposebasetype_member_1 Enum_CodeFlightPurposeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_SCHEDULED")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_NON_SCHEDULED")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_PRIVATE")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_AIR_TRAINING")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_AIR_WORK")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_ALL")
	res = append(res, "Enum_CodeFlightPurposeBaseType_Member_1_PARTICIPANT")

	return
}

func (enum_codeflightpurposebasetype_member_1 Enum_CodeFlightPurposeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SCHEDULED")
	res = append(res, "NON_SCHEDULED")
	res = append(res, "PRIVATE")
	res = append(res, "AIR_TRAINING")
	res = append(res, "AIR_WORK")
	res = append(res, "ALL")
	res = append(res, "PARTICIPANT")

	return
}

// Utility function for Enum_CodeFlightRuleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightrulebasetype_member_1 Enum_CodeFlightRuleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightrulebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightRuleBaseType_Member_1_IFR:
		res = "IFR"
	case Enum_CodeFlightRuleBaseType_Member_1_VFR:
		res = "VFR"
	case Enum_CodeFlightRuleBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeflightrulebasetype_member_1 *Enum_CodeFlightRuleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IFR":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_IFR
		return
	case "VFR":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_VFR
		return
	case "ALL":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightrulebasetype_member_1 *Enum_CodeFlightRuleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightRuleBaseType_Member_1_IFR":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_IFR
	case "Enum_CodeFlightRuleBaseType_Member_1_VFR":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_VFR
	case "Enum_CodeFlightRuleBaseType_Member_1_ALL":
		*enum_codeflightrulebasetype_member_1 = Enum_CodeFlightRuleBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightrulebasetype_member_1 *Enum_CodeFlightRuleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightrulebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightRuleBaseType_Member_1_IFR:
		res = "Enum_CodeFlightRuleBaseType_Member_1_IFR"
	case Enum_CodeFlightRuleBaseType_Member_1_VFR:
		res = "Enum_CodeFlightRuleBaseType_Member_1_VFR"
	case Enum_CodeFlightRuleBaseType_Member_1_ALL:
		res = "Enum_CodeFlightRuleBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeflightrulebasetype_member_1 Enum_CodeFlightRuleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightRuleBaseType_Member_1_IFR")
	res = append(res, "Enum_CodeFlightRuleBaseType_Member_1_VFR")
	res = append(res, "Enum_CodeFlightRuleBaseType_Member_1_ALL")

	return
}

func (enum_codeflightrulebasetype_member_1 Enum_CodeFlightRuleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IFR")
	res = append(res, "VFR")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeFlightStatusBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightstatusbasetype_member_1 Enum_CodeFlightStatusBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightstatusbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightStatusBaseType_Member_1_HEAD:
		res = "HEAD"
	case Enum_CodeFlightStatusBaseType_Member_1_STATE:
		res = "STATE"
	case Enum_CodeFlightStatusBaseType_Member_1_HUM:
		res = "HUM"
	case Enum_CodeFlightStatusBaseType_Member_1_HOSP:
		res = "HOSP"
	case Enum_CodeFlightStatusBaseType_Member_1_SAR:
		res = "SAR"
	case Enum_CodeFlightStatusBaseType_Member_1_ALL:
		res = "ALL"
	case Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY:
		res = "EMERGENCY"
	}
	return
}

func (enum_codeflightstatusbasetype_member_1 *Enum_CodeFlightStatusBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HEAD":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HEAD
		return
	case "STATE":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_STATE
		return
	case "HUM":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HUM
		return
	case "HOSP":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HOSP
		return
	case "SAR":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_SAR
		return
	case "ALL":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_ALL
		return
	case "EMERGENCY":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightstatusbasetype_member_1 *Enum_CodeFlightStatusBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightStatusBaseType_Member_1_HEAD":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HEAD
	case "Enum_CodeFlightStatusBaseType_Member_1_STATE":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_STATE
	case "Enum_CodeFlightStatusBaseType_Member_1_HUM":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HUM
	case "Enum_CodeFlightStatusBaseType_Member_1_HOSP":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_HOSP
	case "Enum_CodeFlightStatusBaseType_Member_1_SAR":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_SAR
	case "Enum_CodeFlightStatusBaseType_Member_1_ALL":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_ALL
	case "Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY":
		*enum_codeflightstatusbasetype_member_1 = Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightstatusbasetype_member_1 *Enum_CodeFlightStatusBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightstatusbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightStatusBaseType_Member_1_HEAD:
		res = "Enum_CodeFlightStatusBaseType_Member_1_HEAD"
	case Enum_CodeFlightStatusBaseType_Member_1_STATE:
		res = "Enum_CodeFlightStatusBaseType_Member_1_STATE"
	case Enum_CodeFlightStatusBaseType_Member_1_HUM:
		res = "Enum_CodeFlightStatusBaseType_Member_1_HUM"
	case Enum_CodeFlightStatusBaseType_Member_1_HOSP:
		res = "Enum_CodeFlightStatusBaseType_Member_1_HOSP"
	case Enum_CodeFlightStatusBaseType_Member_1_SAR:
		res = "Enum_CodeFlightStatusBaseType_Member_1_SAR"
	case Enum_CodeFlightStatusBaseType_Member_1_ALL:
		res = "Enum_CodeFlightStatusBaseType_Member_1_ALL"
	case Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY:
		res = "Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY"
	}
	return
}

func (enum_codeflightstatusbasetype_member_1 Enum_CodeFlightStatusBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_HEAD")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_STATE")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_HUM")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_HOSP")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_SAR")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_ALL")
	res = append(res, "Enum_CodeFlightStatusBaseType_Member_1_EMERGENCY")

	return
}

func (enum_codeflightstatusbasetype_member_1 Enum_CodeFlightStatusBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HEAD")
	res = append(res, "STATE")
	res = append(res, "HUM")
	res = append(res, "HOSP")
	res = append(res, "SAR")
	res = append(res, "ALL")
	res = append(res, "EMERGENCY")

	return
}

// Utility function for Enum_CodeFrictionDeviceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefrictiondevicebasetype_member_1 Enum_CodeFrictionDeviceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefrictiondevicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFrictionDeviceBaseType_Member_1_BRD:
		res = "BRD"
	case Enum_CodeFrictionDeviceBaseType_Member_1_GRT:
		res = "GRT"
	case Enum_CodeFrictionDeviceBaseType_Member_1_MUM:
		res = "MUM"
	case Enum_CodeFrictionDeviceBaseType_Member_1_RFT:
		res = "RFT"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SFH:
		res = "SFH"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SFL:
		res = "SFL"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SKH:
		res = "SKH"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SKL:
		res = "SKL"
	case Enum_CodeFrictionDeviceBaseType_Member_1_TAP:
		res = "TAP"
	}
	return
}

func (enum_codefrictiondevicebasetype_member_1 *Enum_CodeFrictionDeviceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BRD":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_BRD
		return
	case "GRT":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_GRT
		return
	case "MUM":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_MUM
		return
	case "RFT":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_RFT
		return
	case "SFH":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SFH
		return
	case "SFL":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SFL
		return
	case "SKH":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SKH
		return
	case "SKL":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SKL
		return
	case "TAP":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_TAP
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefrictiondevicebasetype_member_1 *Enum_CodeFrictionDeviceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFrictionDeviceBaseType_Member_1_BRD":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_BRD
	case "Enum_CodeFrictionDeviceBaseType_Member_1_GRT":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_GRT
	case "Enum_CodeFrictionDeviceBaseType_Member_1_MUM":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_MUM
	case "Enum_CodeFrictionDeviceBaseType_Member_1_RFT":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_RFT
	case "Enum_CodeFrictionDeviceBaseType_Member_1_SFH":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SFH
	case "Enum_CodeFrictionDeviceBaseType_Member_1_SFL":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SFL
	case "Enum_CodeFrictionDeviceBaseType_Member_1_SKH":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SKH
	case "Enum_CodeFrictionDeviceBaseType_Member_1_SKL":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_SKL
	case "Enum_CodeFrictionDeviceBaseType_Member_1_TAP":
		*enum_codefrictiondevicebasetype_member_1 = Enum_CodeFrictionDeviceBaseType_Member_1_TAP
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefrictiondevicebasetype_member_1 *Enum_CodeFrictionDeviceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefrictiondevicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFrictionDeviceBaseType_Member_1_BRD:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_BRD"
	case Enum_CodeFrictionDeviceBaseType_Member_1_GRT:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_GRT"
	case Enum_CodeFrictionDeviceBaseType_Member_1_MUM:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_MUM"
	case Enum_CodeFrictionDeviceBaseType_Member_1_RFT:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_RFT"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SFH:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_SFH"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SFL:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_SFL"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SKH:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_SKH"
	case Enum_CodeFrictionDeviceBaseType_Member_1_SKL:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_SKL"
	case Enum_CodeFrictionDeviceBaseType_Member_1_TAP:
		res = "Enum_CodeFrictionDeviceBaseType_Member_1_TAP"
	}
	return
}

func (enum_codefrictiondevicebasetype_member_1 Enum_CodeFrictionDeviceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_BRD")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_GRT")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_MUM")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_RFT")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_SFH")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_SFL")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_SKH")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_SKL")
	res = append(res, "Enum_CodeFrictionDeviceBaseType_Member_1_TAP")

	return
}

func (enum_codefrictiondevicebasetype_member_1 Enum_CodeFrictionDeviceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BRD")
	res = append(res, "GRT")
	res = append(res, "MUM")
	res = append(res, "RFT")
	res = append(res, "SFH")
	res = append(res, "SFL")
	res = append(res, "SKH")
	res = append(res, "SKL")
	res = append(res, "TAP")

	return
}

// Utility function for Enum_CodeFrictionEstimateBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefrictionestimatebasetype_member_1 Enum_CodeFrictionEstimateBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefrictionestimatebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFrictionEstimateBaseType_Member_1_GOOD:
		res = "GOOD"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD:
		res = "MEDIUM_GOOD"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM:
		res = "MEDIUM"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR:
		res = "MEDIUM_POOR"
	case Enum_CodeFrictionEstimateBaseType_Member_1_POOR:
		res = "POOR"
	case Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE:
		res = "UNRELIABLE"
	}
	return
}

func (enum_codefrictionestimatebasetype_member_1 *Enum_CodeFrictionEstimateBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GOOD":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_GOOD
		return
	case "MEDIUM_GOOD":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD
		return
	case "MEDIUM":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM
		return
	case "MEDIUM_POOR":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR
		return
	case "POOR":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_POOR
		return
	case "UNRELIABLE":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefrictionestimatebasetype_member_1 *Enum_CodeFrictionEstimateBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFrictionEstimateBaseType_Member_1_GOOD":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_GOOD
	case "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD
	case "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM
	case "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR
	case "Enum_CodeFrictionEstimateBaseType_Member_1_POOR":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_POOR
	case "Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE":
		*enum_codefrictionestimatebasetype_member_1 = Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefrictionestimatebasetype_member_1 *Enum_CodeFrictionEstimateBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefrictionestimatebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFrictionEstimateBaseType_Member_1_GOOD:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_GOOD"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM"
	case Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR"
	case Enum_CodeFrictionEstimateBaseType_Member_1_POOR:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_POOR"
	case Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE:
		res = "Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE"
	}
	return
}

func (enum_codefrictionestimatebasetype_member_1 Enum_CodeFrictionEstimateBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_GOOD")
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_GOOD")
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM")
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_MEDIUM_POOR")
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_POOR")
	res = append(res, "Enum_CodeFrictionEstimateBaseType_Member_1_UNRELIABLE")

	return
}

func (enum_codefrictionestimatebasetype_member_1 Enum_CodeFrictionEstimateBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "GOOD")
	res = append(res, "MEDIUM_GOOD")
	res = append(res, "MEDIUM")
	res = append(res, "MEDIUM_POOR")
	res = append(res, "POOR")
	res = append(res, "UNRELIABLE")

	return
}

// Utility function for Enum_CodeGradeSeparationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codegradeseparationbasetype_member_1 Enum_CodeGradeSeparationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codegradeseparationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS:
		res = "UNDERPASS"
	case Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS:
		res = "OVERPASS"
	}
	return
}

func (enum_codegradeseparationbasetype_member_1 *Enum_CodeGradeSeparationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "UNDERPASS":
		*enum_codegradeseparationbasetype_member_1 = Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS
		return
	case "OVERPASS":
		*enum_codegradeseparationbasetype_member_1 = Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codegradeseparationbasetype_member_1 *Enum_CodeGradeSeparationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS":
		*enum_codegradeseparationbasetype_member_1 = Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS
	case "Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS":
		*enum_codegradeseparationbasetype_member_1 = Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codegradeseparationbasetype_member_1 *Enum_CodeGradeSeparationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codegradeseparationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS:
		res = "Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS"
	case Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS:
		res = "Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS"
	}
	return
}

func (enum_codegradeseparationbasetype_member_1 Enum_CodeGradeSeparationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeGradeSeparationBaseType_Member_1_UNDERPASS")
	res = append(res, "Enum_CodeGradeSeparationBaseType_Member_1_OVERPASS")

	return
}

func (enum_codegradeseparationbasetype_member_1 Enum_CodeGradeSeparationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "UNDERPASS")
	res = append(res, "OVERPASS")

	return
}

// Utility function for Enum_CodeGuidanceLineBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeguidancelinebasetype_member_1 Enum_CodeGuidanceLineBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeguidancelinebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGuidanceLineBaseType_Member_1_RWY:
		res = "RWY"
	case Enum_CodeGuidanceLineBaseType_Member_1_TWY:
		res = "TWY"
	case Enum_CodeGuidanceLineBaseType_Member_1_APRON:
		res = "APRON"
	case Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE:
		res = "GATE_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE:
		res = "LI_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE:
		res = "LO_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE:
		res = "AIR_TLANE"
	}
	return
}

func (enum_codeguidancelinebasetype_member_1 *Enum_CodeGuidanceLineBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RWY":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_RWY
		return
	case "TWY":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_TWY
		return
	case "APRON":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_APRON
		return
	case "GATE_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE
		return
	case "LI_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE
		return
	case "LO_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE
		return
	case "AIR_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeguidancelinebasetype_member_1 *Enum_CodeGuidanceLineBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeGuidanceLineBaseType_Member_1_RWY":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_RWY
	case "Enum_CodeGuidanceLineBaseType_Member_1_TWY":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_TWY
	case "Enum_CodeGuidanceLineBaseType_Member_1_APRON":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_APRON
	case "Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE
	case "Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE
	case "Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE
	case "Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE":
		*enum_codeguidancelinebasetype_member_1 = Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeguidancelinebasetype_member_1 *Enum_CodeGuidanceLineBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeguidancelinebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGuidanceLineBaseType_Member_1_RWY:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_RWY"
	case Enum_CodeGuidanceLineBaseType_Member_1_TWY:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_TWY"
	case Enum_CodeGuidanceLineBaseType_Member_1_APRON:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_APRON"
	case Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE"
	case Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE:
		res = "Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE"
	}
	return
}

func (enum_codeguidancelinebasetype_member_1 Enum_CodeGuidanceLineBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_RWY")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_TWY")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_APRON")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_GATE_TLANE")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_LI_TLANE")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_LO_TLANE")
	res = append(res, "Enum_CodeGuidanceLineBaseType_Member_1_AIR_TLANE")

	return
}

func (enum_codeguidancelinebasetype_member_1 Enum_CodeGuidanceLineBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RWY")
	res = append(res, "TWY")
	res = append(res, "APRON")
	res = append(res, "GATE_TLANE")
	res = append(res, "LI_TLANE")
	res = append(res, "LO_TLANE")
	res = append(res, "AIR_TLANE")

	return
}

// Utility function for Enum_CodeHelicopterPerformanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codehelicopterperformancebasetype_member_1 Enum_CodeHelicopterPerformanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codehelicopterperformancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1:
		res = "1"
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2:
		res = "2"
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3:
		res = "3"
	}
	return
}

func (enum_codehelicopterperformancebasetype_member_1 *Enum_CodeHelicopterPerformanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1
		return
	case "2":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2
		return
	case "3":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codehelicopterperformancebasetype_member_1 *Enum_CodeHelicopterPerformanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1
	case "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2
	case "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3":
		*enum_codehelicopterperformancebasetype_member_1 = Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codehelicopterperformancebasetype_member_1 *Enum_CodeHelicopterPerformanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codehelicopterperformancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1:
		res = "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1"
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2:
		res = "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2"
	case Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3:
		res = "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3"
	}
	return
}

func (enum_codehelicopterperformancebasetype_member_1 Enum_CodeHelicopterPerformanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_1")
	res = append(res, "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_2")
	res = append(res, "Enum_CodeHelicopterPerformanceBaseType_Member_1_N_3")

	return
}

func (enum_codehelicopterperformancebasetype_member_1 Enum_CodeHelicopterPerformanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1")
	res = append(res, "2")
	res = append(res, "3")

	return
}

// Utility function for Enum_CodeHoldingCategoryBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeholdingcategorybasetype_member_1 Enum_CodeHoldingCategoryBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeholdingcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION:
		res = "NON_PRECISION"
	case Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I:
		res = "CAT_I"
	case Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III:
		res = "CAT_II_III"
	}
	return
}

func (enum_codeholdingcategorybasetype_member_1 *Enum_CodeHoldingCategoryBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NON_PRECISION":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION
		return
	case "CAT_I":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I
		return
	case "CAT_II_III":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeholdingcategorybasetype_member_1 *Enum_CodeHoldingCategoryBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION
	case "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I
	case "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III":
		*enum_codeholdingcategorybasetype_member_1 = Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeholdingcategorybasetype_member_1 *Enum_CodeHoldingCategoryBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeholdingcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION:
		res = "Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION"
	case Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I:
		res = "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I"
	case Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III:
		res = "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III"
	}
	return
}

func (enum_codeholdingcategorybasetype_member_1 Enum_CodeHoldingCategoryBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeHoldingCategoryBaseType_Member_1_NON_PRECISION")
	res = append(res, "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_I")
	res = append(res, "Enum_CodeHoldingCategoryBaseType_Member_1_CAT_II_III")

	return
}

func (enum_codeholdingcategorybasetype_member_1 Enum_CodeHoldingCategoryBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NON_PRECISION")
	res = append(res, "CAT_I")
	res = append(res, "CAT_II_III")

	return
}

// Utility function for Enum_CodeLightHoldingPositionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelightholdingpositionbasetype_member_1 Enum_CodeLightHoldingPositionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelightholdingpositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR:
		res = "STOP_BAR"
	case Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN:
		res = "SIGN"
	}
	return
}

func (enum_codelightholdingpositionbasetype_member_1 *Enum_CodeLightHoldingPositionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STOP_BAR":
		*enum_codelightholdingpositionbasetype_member_1 = Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR
		return
	case "SIGN":
		*enum_codelightholdingpositionbasetype_member_1 = Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelightholdingpositionbasetype_member_1 *Enum_CodeLightHoldingPositionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR":
		*enum_codelightholdingpositionbasetype_member_1 = Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR
	case "Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN":
		*enum_codelightholdingpositionbasetype_member_1 = Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelightholdingpositionbasetype_member_1 *Enum_CodeLightHoldingPositionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelightholdingpositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR:
		res = "Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR"
	case Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN:
		res = "Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN"
	}
	return
}

func (enum_codelightholdingpositionbasetype_member_1 Enum_CodeLightHoldingPositionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLightHoldingPositionBaseType_Member_1_STOP_BAR")
	res = append(res, "Enum_CodeLightHoldingPositionBaseType_Member_1_SIGN")

	return
}

func (enum_codelightholdingpositionbasetype_member_1 Enum_CodeLightHoldingPositionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STOP_BAR")
	res = append(res, "SIGN")

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

// Utility function for Enum_CodeLightingJARBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelightingjarbasetype_member_1 Enum_CodeLightingJARBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelightingjarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightingJARBaseType_Member_1_FALS:
		res = "FALS"
	case Enum_CodeLightingJARBaseType_Member_1_IALS:
		res = "IALS"
	case Enum_CodeLightingJARBaseType_Member_1_BALS:
		res = "BALS"
	case Enum_CodeLightingJARBaseType_Member_1_NALS:
		res = "NALS"
	}
	return
}

func (enum_codelightingjarbasetype_member_1 *Enum_CodeLightingJARBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_FALS
		return
	case "IALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_IALS
		return
	case "BALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_BALS
		return
	case "NALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_NALS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelightingjarbasetype_member_1 *Enum_CodeLightingJARBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLightingJARBaseType_Member_1_FALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_FALS
	case "Enum_CodeLightingJARBaseType_Member_1_IALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_IALS
	case "Enum_CodeLightingJARBaseType_Member_1_BALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_BALS
	case "Enum_CodeLightingJARBaseType_Member_1_NALS":
		*enum_codelightingjarbasetype_member_1 = Enum_CodeLightingJARBaseType_Member_1_NALS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelightingjarbasetype_member_1 *Enum_CodeLightingJARBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelightingjarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLightingJARBaseType_Member_1_FALS:
		res = "Enum_CodeLightingJARBaseType_Member_1_FALS"
	case Enum_CodeLightingJARBaseType_Member_1_IALS:
		res = "Enum_CodeLightingJARBaseType_Member_1_IALS"
	case Enum_CodeLightingJARBaseType_Member_1_BALS:
		res = "Enum_CodeLightingJARBaseType_Member_1_BALS"
	case Enum_CodeLightingJARBaseType_Member_1_NALS:
		res = "Enum_CodeLightingJARBaseType_Member_1_NALS"
	}
	return
}

func (enum_codelightingjarbasetype_member_1 Enum_CodeLightingJARBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLightingJARBaseType_Member_1_FALS")
	res = append(res, "Enum_CodeLightingJARBaseType_Member_1_IALS")
	res = append(res, "Enum_CodeLightingJARBaseType_Member_1_BALS")
	res = append(res, "Enum_CodeLightingJARBaseType_Member_1_NALS")

	return
}

func (enum_codelightingjarbasetype_member_1 Enum_CodeLightingJARBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FALS")
	res = append(res, "IALS")
	res = append(res, "BALS")
	res = append(res, "NALS")

	return
}

// Utility function for Enum_CodeLoadingBridgeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeloadingbridgebasetype_member_1 Enum_CodeLoadingBridgeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeloadingbridgebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLoadingBridgeBaseType_Member_1_ARM:
		res = "ARM"
	case Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM:
		res = "MOVABLE_ARM"
	case Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP:
		res = "PORTABLE_RAMP"
	case Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS:
		res = "PORTABLE_STAIRS"
	}
	return
}

func (enum_codeloadingbridgebasetype_member_1 *Enum_CodeLoadingBridgeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ARM":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_ARM
		return
	case "MOVABLE_ARM":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM
		return
	case "PORTABLE_RAMP":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP
		return
	case "PORTABLE_STAIRS":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeloadingbridgebasetype_member_1 *Enum_CodeLoadingBridgeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLoadingBridgeBaseType_Member_1_ARM":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_ARM
	case "Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM
	case "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP
	case "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS":
		*enum_codeloadingbridgebasetype_member_1 = Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeloadingbridgebasetype_member_1 *Enum_CodeLoadingBridgeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeloadingbridgebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLoadingBridgeBaseType_Member_1_ARM:
		res = "Enum_CodeLoadingBridgeBaseType_Member_1_ARM"
	case Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM:
		res = "Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM"
	case Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP:
		res = "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP"
	case Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS:
		res = "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS"
	}
	return
}

func (enum_codeloadingbridgebasetype_member_1 Enum_CodeLoadingBridgeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLoadingBridgeBaseType_Member_1_ARM")
	res = append(res, "Enum_CodeLoadingBridgeBaseType_Member_1_MOVABLE_ARM")
	res = append(res, "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_RAMP")
	res = append(res, "Enum_CodeLoadingBridgeBaseType_Member_1_PORTABLE_STAIRS")

	return
}

func (enum_codeloadingbridgebasetype_member_1 Enum_CodeLoadingBridgeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ARM")
	res = append(res, "MOVABLE_ARM")
	res = append(res, "PORTABLE_RAMP")
	res = append(res, "PORTABLE_STAIRS")

	return
}

// Utility function for Enum_CodeLogicalOperatorBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelogicaloperatorbasetype_member_1 Enum_CodeLogicalOperatorBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelogicaloperatorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLogicalOperatorBaseType_Member_1_AND:
		res = "AND"
	case Enum_CodeLogicalOperatorBaseType_Member_1_OR:
		res = "OR"
	case Enum_CodeLogicalOperatorBaseType_Member_1_NOT:
		res = "NOT"
	case Enum_CodeLogicalOperatorBaseType_Member_1_NONE:
		res = "NONE"
	}
	return
}

func (enum_codelogicaloperatorbasetype_member_1 *Enum_CodeLogicalOperatorBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AND":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_AND
		return
	case "OR":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_OR
		return
	case "NOT":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_NOT
		return
	case "NONE":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_NONE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelogicaloperatorbasetype_member_1 *Enum_CodeLogicalOperatorBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLogicalOperatorBaseType_Member_1_AND":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_AND
	case "Enum_CodeLogicalOperatorBaseType_Member_1_OR":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_OR
	case "Enum_CodeLogicalOperatorBaseType_Member_1_NOT":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_NOT
	case "Enum_CodeLogicalOperatorBaseType_Member_1_NONE":
		*enum_codelogicaloperatorbasetype_member_1 = Enum_CodeLogicalOperatorBaseType_Member_1_NONE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelogicaloperatorbasetype_member_1 *Enum_CodeLogicalOperatorBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelogicaloperatorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLogicalOperatorBaseType_Member_1_AND:
		res = "Enum_CodeLogicalOperatorBaseType_Member_1_AND"
	case Enum_CodeLogicalOperatorBaseType_Member_1_OR:
		res = "Enum_CodeLogicalOperatorBaseType_Member_1_OR"
	case Enum_CodeLogicalOperatorBaseType_Member_1_NOT:
		res = "Enum_CodeLogicalOperatorBaseType_Member_1_NOT"
	case Enum_CodeLogicalOperatorBaseType_Member_1_NONE:
		res = "Enum_CodeLogicalOperatorBaseType_Member_1_NONE"
	}
	return
}

func (enum_codelogicaloperatorbasetype_member_1 Enum_CodeLogicalOperatorBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLogicalOperatorBaseType_Member_1_AND")
	res = append(res, "Enum_CodeLogicalOperatorBaseType_Member_1_OR")
	res = append(res, "Enum_CodeLogicalOperatorBaseType_Member_1_NOT")
	res = append(res, "Enum_CodeLogicalOperatorBaseType_Member_1_NONE")

	return
}

func (enum_codelogicaloperatorbasetype_member_1 Enum_CodeLogicalOperatorBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AND")
	res = append(res, "OR")
	res = append(res, "NOT")
	res = append(res, "NONE")

	return
}

// Utility function for Enum_CodeMarkingConditionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemarkingconditionbasetype_member_1 Enum_CodeMarkingConditionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemarkingconditionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkingConditionBaseType_Member_1_GOOD:
		res = "GOOD"
	case Enum_CodeMarkingConditionBaseType_Member_1_FAIR:
		res = "FAIR"
	case Enum_CodeMarkingConditionBaseType_Member_1_POOR:
		res = "POOR"
	case Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT:
		res = "EXCELLENT"
	}
	return
}

func (enum_codemarkingconditionbasetype_member_1 *Enum_CodeMarkingConditionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GOOD":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_GOOD
		return
	case "FAIR":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_FAIR
		return
	case "POOR":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_POOR
		return
	case "EXCELLENT":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemarkingconditionbasetype_member_1 *Enum_CodeMarkingConditionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMarkingConditionBaseType_Member_1_GOOD":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_GOOD
	case "Enum_CodeMarkingConditionBaseType_Member_1_FAIR":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_FAIR
	case "Enum_CodeMarkingConditionBaseType_Member_1_POOR":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_POOR
	case "Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT":
		*enum_codemarkingconditionbasetype_member_1 = Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemarkingconditionbasetype_member_1 *Enum_CodeMarkingConditionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemarkingconditionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkingConditionBaseType_Member_1_GOOD:
		res = "Enum_CodeMarkingConditionBaseType_Member_1_GOOD"
	case Enum_CodeMarkingConditionBaseType_Member_1_FAIR:
		res = "Enum_CodeMarkingConditionBaseType_Member_1_FAIR"
	case Enum_CodeMarkingConditionBaseType_Member_1_POOR:
		res = "Enum_CodeMarkingConditionBaseType_Member_1_POOR"
	case Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT:
		res = "Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT"
	}
	return
}

func (enum_codemarkingconditionbasetype_member_1 Enum_CodeMarkingConditionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMarkingConditionBaseType_Member_1_GOOD")
	res = append(res, "Enum_CodeMarkingConditionBaseType_Member_1_FAIR")
	res = append(res, "Enum_CodeMarkingConditionBaseType_Member_1_POOR")
	res = append(res, "Enum_CodeMarkingConditionBaseType_Member_1_EXCELLENT")

	return
}

func (enum_codemarkingconditionbasetype_member_1 Enum_CodeMarkingConditionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "GOOD")
	res = append(res, "FAIR")
	res = append(res, "POOR")
	res = append(res, "EXCELLENT")

	return
}

// Utility function for Enum_CodeMarkingStyleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemarkingstylebasetype_member_1 Enum_CodeMarkingStyleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemarkingstylebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkingStyleBaseType_Member_1_SOLID:
		res = "SOLID"
	case Enum_CodeMarkingStyleBaseType_Member_1_DASHED:
		res = "DASHED"
	case Enum_CodeMarkingStyleBaseType_Member_1_DOTTED:
		res = "DOTTED"
	}
	return
}

func (enum_codemarkingstylebasetype_member_1 *Enum_CodeMarkingStyleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SOLID":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_SOLID
		return
	case "DASHED":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_DASHED
		return
	case "DOTTED":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_DOTTED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemarkingstylebasetype_member_1 *Enum_CodeMarkingStyleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMarkingStyleBaseType_Member_1_SOLID":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_SOLID
	case "Enum_CodeMarkingStyleBaseType_Member_1_DASHED":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_DASHED
	case "Enum_CodeMarkingStyleBaseType_Member_1_DOTTED":
		*enum_codemarkingstylebasetype_member_1 = Enum_CodeMarkingStyleBaseType_Member_1_DOTTED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemarkingstylebasetype_member_1 *Enum_CodeMarkingStyleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemarkingstylebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkingStyleBaseType_Member_1_SOLID:
		res = "Enum_CodeMarkingStyleBaseType_Member_1_SOLID"
	case Enum_CodeMarkingStyleBaseType_Member_1_DASHED:
		res = "Enum_CodeMarkingStyleBaseType_Member_1_DASHED"
	case Enum_CodeMarkingStyleBaseType_Member_1_DOTTED:
		res = "Enum_CodeMarkingStyleBaseType_Member_1_DOTTED"
	}
	return
}

func (enum_codemarkingstylebasetype_member_1 Enum_CodeMarkingStyleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMarkingStyleBaseType_Member_1_SOLID")
	res = append(res, "Enum_CodeMarkingStyleBaseType_Member_1_DASHED")
	res = append(res, "Enum_CodeMarkingStyleBaseType_Member_1_DOTTED")

	return
}

func (enum_codemarkingstylebasetype_member_1 Enum_CodeMarkingStyleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SOLID")
	res = append(res, "DASHED")
	res = append(res, "DOTTED")

	return
}

// Utility function for Enum_CodeMeteoConditionsBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemeteoconditionsbasetype_member_1 Enum_CodeMeteoConditionsBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemeteoconditionsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMeteoConditionsBaseType_Member_1_IMC:
		res = "IMC"
	case Enum_CodeMeteoConditionsBaseType_Member_1_VMC:
		res = "VMC"
	case Enum_CodeMeteoConditionsBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codemeteoconditionsbasetype_member_1 *Enum_CodeMeteoConditionsBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IMC":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_IMC
		return
	case "VMC":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_VMC
		return
	case "ALL":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemeteoconditionsbasetype_member_1 *Enum_CodeMeteoConditionsBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMeteoConditionsBaseType_Member_1_IMC":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_IMC
	case "Enum_CodeMeteoConditionsBaseType_Member_1_VMC":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_VMC
	case "Enum_CodeMeteoConditionsBaseType_Member_1_ALL":
		*enum_codemeteoconditionsbasetype_member_1 = Enum_CodeMeteoConditionsBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemeteoconditionsbasetype_member_1 *Enum_CodeMeteoConditionsBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemeteoconditionsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMeteoConditionsBaseType_Member_1_IMC:
		res = "Enum_CodeMeteoConditionsBaseType_Member_1_IMC"
	case Enum_CodeMeteoConditionsBaseType_Member_1_VMC:
		res = "Enum_CodeMeteoConditionsBaseType_Member_1_VMC"
	case Enum_CodeMeteoConditionsBaseType_Member_1_ALL:
		res = "Enum_CodeMeteoConditionsBaseType_Member_1_ALL"
	}
	return
}

func (enum_codemeteoconditionsbasetype_member_1 Enum_CodeMeteoConditionsBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMeteoConditionsBaseType_Member_1_IMC")
	res = append(res, "Enum_CodeMeteoConditionsBaseType_Member_1_VMC")
	res = append(res, "Enum_CodeMeteoConditionsBaseType_Member_1_ALL")

	return
}

func (enum_codemeteoconditionsbasetype_member_1 Enum_CodeMeteoConditionsBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IMC")
	res = append(res, "VMC")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeMilitaryOperationsBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemilitaryoperationsbasetype_member_1 Enum_CodeMilitaryOperationsBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemilitaryoperationsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL:
		res = "CIVIL"
	case Enum_CodeMilitaryOperationsBaseType_Member_1_MIL:
		res = "MIL"
	case Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT:
		res = "JOINT"
	}
	return
}

func (enum_codemilitaryoperationsbasetype_member_1 *Enum_CodeMilitaryOperationsBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CIVIL":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL
		return
	case "MIL":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_MIL
		return
	case "JOINT":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemilitaryoperationsbasetype_member_1 *Enum_CodeMilitaryOperationsBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL
	case "Enum_CodeMilitaryOperationsBaseType_Member_1_MIL":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_MIL
	case "Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT":
		*enum_codemilitaryoperationsbasetype_member_1 = Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemilitaryoperationsbasetype_member_1 *Enum_CodeMilitaryOperationsBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemilitaryoperationsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL:
		res = "Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL"
	case Enum_CodeMilitaryOperationsBaseType_Member_1_MIL:
		res = "Enum_CodeMilitaryOperationsBaseType_Member_1_MIL"
	case Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT:
		res = "Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT"
	}
	return
}

func (enum_codemilitaryoperationsbasetype_member_1 Enum_CodeMilitaryOperationsBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMilitaryOperationsBaseType_Member_1_CIVIL")
	res = append(res, "Enum_CodeMilitaryOperationsBaseType_Member_1_MIL")
	res = append(res, "Enum_CodeMilitaryOperationsBaseType_Member_1_JOINT")

	return
}

func (enum_codemilitaryoperationsbasetype_member_1 Enum_CodeMilitaryOperationsBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CIVIL")
	res = append(res, "MIL")
	res = append(res, "JOINT")

	return
}

// Utility function for Enum_CodeMilitaryStatusBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemilitarystatusbasetype_member_1 Enum_CodeMilitaryStatusBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemilitarystatusbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryStatusBaseType_Member_1_MIL:
		res = "MIL"
	case Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL:
		res = "CIVIL"
	case Enum_CodeMilitaryStatusBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codemilitarystatusbasetype_member_1 *Enum_CodeMilitaryStatusBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MIL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_MIL
		return
	case "CIVIL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL
		return
	case "ALL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemilitarystatusbasetype_member_1 *Enum_CodeMilitaryStatusBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMilitaryStatusBaseType_Member_1_MIL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_MIL
	case "Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL
	case "Enum_CodeMilitaryStatusBaseType_Member_1_ALL":
		*enum_codemilitarystatusbasetype_member_1 = Enum_CodeMilitaryStatusBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemilitarystatusbasetype_member_1 *Enum_CodeMilitaryStatusBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemilitarystatusbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryStatusBaseType_Member_1_MIL:
		res = "Enum_CodeMilitaryStatusBaseType_Member_1_MIL"
	case Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL:
		res = "Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL"
	case Enum_CodeMilitaryStatusBaseType_Member_1_ALL:
		res = "Enum_CodeMilitaryStatusBaseType_Member_1_ALL"
	}
	return
}

func (enum_codemilitarystatusbasetype_member_1 Enum_CodeMilitaryStatusBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMilitaryStatusBaseType_Member_1_MIL")
	res = append(res, "Enum_CodeMilitaryStatusBaseType_Member_1_CIVIL")
	res = append(res, "Enum_CodeMilitaryStatusBaseType_Member_1_ALL")

	return
}

func (enum_codemilitarystatusbasetype_member_1 Enum_CodeMilitaryStatusBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MIL")
	res = append(res, "CIVIL")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeNavigationEquipmentBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavigationequipmentbasetype_member_1 Enum_CodeNavigationEquipmentBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavigationequipmentbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationEquipmentBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME:
		res = "VOR_DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME:
		res = "DME_DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN:
		res = "TACAN"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_ILS:
		res = "ILS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_MLS:
		res = "MLS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS:
		res = "GNSS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS:
		res = "WAAS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN:
		res = "LORAN"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_INS:
		res = "INS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_FMS:
		res = "FMS"
	}
	return
}

func (enum_codenavigationequipmentbasetype_member_1 *Enum_CodeNavigationEquipmentBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_DME
		return
	case "VOR_DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME
		return
	case "DME_DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME
		return
	case "TACAN":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN
		return
	case "ILS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_ILS
		return
	case "MLS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_MLS
		return
	case "GNSS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS
		return
	case "WAAS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS
		return
	case "LORAN":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN
		return
	case "INS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_INS
		return
	case "FMS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_FMS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavigationequipmentbasetype_member_1 *Enum_CodeNavigationEquipmentBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_DME
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_ILS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_ILS
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_MLS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_MLS
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_INS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_INS
	case "Enum_CodeNavigationEquipmentBaseType_Member_1_FMS":
		*enum_codenavigationequipmentbasetype_member_1 = Enum_CodeNavigationEquipmentBaseType_Member_1_FMS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavigationequipmentbasetype_member_1 *Enum_CodeNavigationEquipmentBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavigationequipmentbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationEquipmentBaseType_Member_1_DME:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_ILS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_ILS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_MLS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_MLS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_INS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_INS"
	case Enum_CodeNavigationEquipmentBaseType_Member_1_FMS:
		res = "Enum_CodeNavigationEquipmentBaseType_Member_1_FMS"
	}
	return
}

func (enum_codenavigationequipmentbasetype_member_1 Enum_CodeNavigationEquipmentBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_DME")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_VOR_DME")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_DME_DME")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_TACAN")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_ILS")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_MLS")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_GNSS")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_WAAS")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_LORAN")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_INS")
	res = append(res, "Enum_CodeNavigationEquipmentBaseType_Member_1_FMS")

	return
}

func (enum_codenavigationequipmentbasetype_member_1 Enum_CodeNavigationEquipmentBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DME")
	res = append(res, "VOR_DME")
	res = append(res, "DME_DME")
	res = append(res, "TACAN")
	res = append(res, "ILS")
	res = append(res, "MLS")
	res = append(res, "GNSS")
	res = append(res, "WAAS")
	res = append(res, "LORAN")
	res = append(res, "INS")
	res = append(res, "FMS")

	return
}

// Utility function for Enum_CodeNavigationSpecificationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavigationspecificationbasetype_member_1 Enum_CodeNavigationSpecificationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavigationspecificationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10:
		res = "RNAV_10"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5:
		res = "RNAV_5"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2:
		res = "RNAV_2"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1:
		res = "RNAV_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4:
		res = "RNP_4"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2:
		res = "RNP_2"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1:
		res = "BASIC_RNP_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1:
		res = "ADVANCED_RNP_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH:
		res = "RNP_APCH"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH:
		res = "RNP_AR_APCH"
	}
	return
}

func (enum_codenavigationspecificationbasetype_member_1 *Enum_CodeNavigationSpecificationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RNAV_10":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10
		return
	case "RNAV_5":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5
		return
	case "RNAV_2":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2
		return
	case "RNAV_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1
		return
	case "RNP_4":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4
		return
	case "RNP_2":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2
		return
	case "BASIC_RNP_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1
		return
	case "ADVANCED_RNP_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1
		return
	case "RNP_APCH":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH
		return
	case "RNP_AR_APCH":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavigationspecificationbasetype_member_1 *Enum_CodeNavigationSpecificationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH
	case "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH":
		*enum_codenavigationspecificationbasetype_member_1 = Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavigationspecificationbasetype_member_1 *Enum_CodeNavigationSpecificationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavigationspecificationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH"
	case Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH:
		res = "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH"
	}
	return
}

func (enum_codenavigationspecificationbasetype_member_1 Enum_CodeNavigationSpecificationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_10")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_5")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_2")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNAV_1")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_4")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_2")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_BASIC_RNP_1")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_ADVANCED_RNP_1")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_APCH")
	res = append(res, "Enum_CodeNavigationSpecificationBaseType_Member_1_RNP_AR_APCH")

	return
}

func (enum_codenavigationspecificationbasetype_member_1 Enum_CodeNavigationSpecificationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RNAV_10")
	res = append(res, "RNAV_5")
	res = append(res, "RNAV_2")
	res = append(res, "RNAV_1")
	res = append(res, "RNP_4")
	res = append(res, "RNP_2")
	res = append(res, "BASIC_RNP_1")
	res = append(res, "ADVANCED_RNP_1")
	res = append(res, "RNP_APCH")
	res = append(res, "RNP_AR_APCH")

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

// Utility function for Enum_CodeOperationAirportHeliportBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeoperationairportheliportbasetype_member_1 Enum_CodeOperationAirportHeliportBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeoperationairportheliportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING:
		res = "LANDING"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF:
		res = "TAKEOFF"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO:
		res = "TOUCHGO"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH:
		res = "TRAIN_APPROACH"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING:
		res = "ALTN_LANDING"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW:
		res = "AIRSHOW"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeoperationairportheliportbasetype_member_1 *Enum_CodeOperationAirportHeliportBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LANDING":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING
		return
	case "TAKEOFF":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF
		return
	case "TOUCHGO":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO
		return
	case "TRAIN_APPROACH":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH
		return
	case "ALTN_LANDING":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING
		return
	case "AIRSHOW":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW
		return
	case "ALL":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeoperationairportheliportbasetype_member_1 *Enum_CodeOperationAirportHeliportBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW
	case "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL":
		*enum_codeoperationairportheliportbasetype_member_1 = Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeoperationairportheliportbasetype_member_1 *Enum_CodeOperationAirportHeliportBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeoperationairportheliportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW"
	case Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL:
		res = "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeoperationairportheliportbasetype_member_1 Enum_CodeOperationAirportHeliportBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_LANDING")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_TAKEOFF")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_TOUCHGO")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_TRAIN_APPROACH")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALTN_LANDING")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_AIRSHOW")
	res = append(res, "Enum_CodeOperationAirportHeliportBaseType_Member_1_ALL")

	return
}

func (enum_codeoperationairportheliportbasetype_member_1 Enum_CodeOperationAirportHeliportBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LANDING")
	res = append(res, "TAKEOFF")
	res = append(res, "TOUCHGO")
	res = append(res, "TRAIN_APPROACH")
	res = append(res, "ALTN_LANDING")
	res = append(res, "AIRSHOW")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeOperationManoeuvringAreaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeoperationmanoeuvringareabasetype_member_1 Enum_CodeOperationManoeuvringAreaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeoperationmanoeuvringareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING:
		res = "LANDING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF:
		res = "TAKEOFF"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO:
		res = "TOUCHGO"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH:
		res = "TRAIN_APPROACH"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING:
		res = "TAXIING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING:
		res = "CROSSING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW:
		res = "AIRSHOW"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeoperationmanoeuvringareabasetype_member_1 *Enum_CodeOperationManoeuvringAreaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LANDING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING
		return
	case "TAKEOFF":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF
		return
	case "TOUCHGO":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO
		return
	case "TRAIN_APPROACH":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH
		return
	case "TAXIING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING
		return
	case "CROSSING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING
		return
	case "AIRSHOW":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW
		return
	case "ALL":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeoperationmanoeuvringareabasetype_member_1 *Enum_CodeOperationManoeuvringAreaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW
	case "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL":
		*enum_codeoperationmanoeuvringareabasetype_member_1 = Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeoperationmanoeuvringareabasetype_member_1 *Enum_CodeOperationManoeuvringAreaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeoperationmanoeuvringareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW"
	case Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL:
		res = "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeoperationmanoeuvringareabasetype_member_1 Enum_CodeOperationManoeuvringAreaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_LANDING")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAKEOFF")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TOUCHGO")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TRAIN_APPROACH")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_TAXIING")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_CROSSING")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_AIRSHOW")
	res = append(res, "Enum_CodeOperationManoeuvringAreaBaseType_Member_1_ALL")

	return
}

func (enum_codeoperationmanoeuvringareabasetype_member_1 Enum_CodeOperationManoeuvringAreaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LANDING")
	res = append(res, "TAKEOFF")
	res = append(res, "TOUCHGO")
	res = append(res, "TRAIN_APPROACH")
	res = append(res, "TAXIING")
	res = append(res, "CROSSING")
	res = append(res, "AIRSHOW")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodePCNMethodBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepcnmethodbasetype_member_1 Enum_CodePCNMethodBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepcnmethodbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNMethodBaseType_Member_1_TECH:
		res = "TECH"
	case Enum_CodePCNMethodBaseType_Member_1_ACFT:
		res = "ACFT"
	}
	return
}

func (enum_codepcnmethodbasetype_member_1 *Enum_CodePCNMethodBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TECH":
		*enum_codepcnmethodbasetype_member_1 = Enum_CodePCNMethodBaseType_Member_1_TECH
		return
	case "ACFT":
		*enum_codepcnmethodbasetype_member_1 = Enum_CodePCNMethodBaseType_Member_1_ACFT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepcnmethodbasetype_member_1 *Enum_CodePCNMethodBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePCNMethodBaseType_Member_1_TECH":
		*enum_codepcnmethodbasetype_member_1 = Enum_CodePCNMethodBaseType_Member_1_TECH
	case "Enum_CodePCNMethodBaseType_Member_1_ACFT":
		*enum_codepcnmethodbasetype_member_1 = Enum_CodePCNMethodBaseType_Member_1_ACFT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepcnmethodbasetype_member_1 *Enum_CodePCNMethodBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepcnmethodbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNMethodBaseType_Member_1_TECH:
		res = "Enum_CodePCNMethodBaseType_Member_1_TECH"
	case Enum_CodePCNMethodBaseType_Member_1_ACFT:
		res = "Enum_CodePCNMethodBaseType_Member_1_ACFT"
	}
	return
}

func (enum_codepcnmethodbasetype_member_1 Enum_CodePCNMethodBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePCNMethodBaseType_Member_1_TECH")
	res = append(res, "Enum_CodePCNMethodBaseType_Member_1_ACFT")

	return
}

func (enum_codepcnmethodbasetype_member_1 Enum_CodePCNMethodBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TECH")
	res = append(res, "ACFT")

	return
}

// Utility function for Enum_CodePCNPavementBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepcnpavementbasetype_member_1 Enum_CodePCNPavementBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepcnpavementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNPavementBaseType_Member_1_RIGID:
		res = "RIGID"
	case Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE:
		res = "FLEXIBLE"
	}
	return
}

func (enum_codepcnpavementbasetype_member_1 *Enum_CodePCNPavementBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RIGID":
		*enum_codepcnpavementbasetype_member_1 = Enum_CodePCNPavementBaseType_Member_1_RIGID
		return
	case "FLEXIBLE":
		*enum_codepcnpavementbasetype_member_1 = Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepcnpavementbasetype_member_1 *Enum_CodePCNPavementBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePCNPavementBaseType_Member_1_RIGID":
		*enum_codepcnpavementbasetype_member_1 = Enum_CodePCNPavementBaseType_Member_1_RIGID
	case "Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE":
		*enum_codepcnpavementbasetype_member_1 = Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepcnpavementbasetype_member_1 *Enum_CodePCNPavementBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepcnpavementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNPavementBaseType_Member_1_RIGID:
		res = "Enum_CodePCNPavementBaseType_Member_1_RIGID"
	case Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE:
		res = "Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE"
	}
	return
}

func (enum_codepcnpavementbasetype_member_1 Enum_CodePCNPavementBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePCNPavementBaseType_Member_1_RIGID")
	res = append(res, "Enum_CodePCNPavementBaseType_Member_1_FLEXIBLE")

	return
}

func (enum_codepcnpavementbasetype_member_1 Enum_CodePCNPavementBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RIGID")
	res = append(res, "FLEXIBLE")

	return
}

// Utility function for Enum_CodePCNSubgradeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepcnsubgradebasetype_member_1 Enum_CodePCNSubgradeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepcnsubgradebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNSubgradeBaseType_Member_1_A:
		res = "A"
	case Enum_CodePCNSubgradeBaseType_Member_1_B:
		res = "B"
	case Enum_CodePCNSubgradeBaseType_Member_1_C:
		res = "C"
	case Enum_CodePCNSubgradeBaseType_Member_1_D:
		res = "D"
	}
	return
}

func (enum_codepcnsubgradebasetype_member_1 *Enum_CodePCNSubgradeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_A
		return
	case "B":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_B
		return
	case "C":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_C
		return
	case "D":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_D
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepcnsubgradebasetype_member_1 *Enum_CodePCNSubgradeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePCNSubgradeBaseType_Member_1_A":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_A
	case "Enum_CodePCNSubgradeBaseType_Member_1_B":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_B
	case "Enum_CodePCNSubgradeBaseType_Member_1_C":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_C
	case "Enum_CodePCNSubgradeBaseType_Member_1_D":
		*enum_codepcnsubgradebasetype_member_1 = Enum_CodePCNSubgradeBaseType_Member_1_D
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepcnsubgradebasetype_member_1 *Enum_CodePCNSubgradeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepcnsubgradebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNSubgradeBaseType_Member_1_A:
		res = "Enum_CodePCNSubgradeBaseType_Member_1_A"
	case Enum_CodePCNSubgradeBaseType_Member_1_B:
		res = "Enum_CodePCNSubgradeBaseType_Member_1_B"
	case Enum_CodePCNSubgradeBaseType_Member_1_C:
		res = "Enum_CodePCNSubgradeBaseType_Member_1_C"
	case Enum_CodePCNSubgradeBaseType_Member_1_D:
		res = "Enum_CodePCNSubgradeBaseType_Member_1_D"
	}
	return
}

func (enum_codepcnsubgradebasetype_member_1 Enum_CodePCNSubgradeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePCNSubgradeBaseType_Member_1_A")
	res = append(res, "Enum_CodePCNSubgradeBaseType_Member_1_B")
	res = append(res, "Enum_CodePCNSubgradeBaseType_Member_1_C")
	res = append(res, "Enum_CodePCNSubgradeBaseType_Member_1_D")

	return
}

func (enum_codepcnsubgradebasetype_member_1 Enum_CodePCNSubgradeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")

	return
}

// Utility function for Enum_CodePCNTyrePressureBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepcntyrepressurebasetype_member_1 Enum_CodePCNTyrePressureBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepcntyrepressurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNTyrePressureBaseType_Member_1_W:
		res = "W"
	case Enum_CodePCNTyrePressureBaseType_Member_1_X:
		res = "X"
	case Enum_CodePCNTyrePressureBaseType_Member_1_Y:
		res = "Y"
	case Enum_CodePCNTyrePressureBaseType_Member_1_Z:
		res = "Z"
	}
	return
}

func (enum_codepcntyrepressurebasetype_member_1 *Enum_CodePCNTyrePressureBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "W":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_W
		return
	case "X":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_X
		return
	case "Y":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_Y
		return
	case "Z":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_Z
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepcntyrepressurebasetype_member_1 *Enum_CodePCNTyrePressureBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePCNTyrePressureBaseType_Member_1_W":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_W
	case "Enum_CodePCNTyrePressureBaseType_Member_1_X":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_X
	case "Enum_CodePCNTyrePressureBaseType_Member_1_Y":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_Y
	case "Enum_CodePCNTyrePressureBaseType_Member_1_Z":
		*enum_codepcntyrepressurebasetype_member_1 = Enum_CodePCNTyrePressureBaseType_Member_1_Z
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepcntyrepressurebasetype_member_1 *Enum_CodePCNTyrePressureBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepcntyrepressurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePCNTyrePressureBaseType_Member_1_W:
		res = "Enum_CodePCNTyrePressureBaseType_Member_1_W"
	case Enum_CodePCNTyrePressureBaseType_Member_1_X:
		res = "Enum_CodePCNTyrePressureBaseType_Member_1_X"
	case Enum_CodePCNTyrePressureBaseType_Member_1_Y:
		res = "Enum_CodePCNTyrePressureBaseType_Member_1_Y"
	case Enum_CodePCNTyrePressureBaseType_Member_1_Z:
		res = "Enum_CodePCNTyrePressureBaseType_Member_1_Z"
	}
	return
}

func (enum_codepcntyrepressurebasetype_member_1 Enum_CodePCNTyrePressureBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePCNTyrePressureBaseType_Member_1_W")
	res = append(res, "Enum_CodePCNTyrePressureBaseType_Member_1_X")
	res = append(res, "Enum_CodePCNTyrePressureBaseType_Member_1_Y")
	res = append(res, "Enum_CodePCNTyrePressureBaseType_Member_1_Z")

	return
}

func (enum_codepcntyrepressurebasetype_member_1 Enum_CodePCNTyrePressureBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "W")
	res = append(res, "X")
	res = append(res, "Y")
	res = append(res, "Z")

	return
}

// Utility function for Enum_CodeProtectAreaSectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprotectareasectionbasetype_member_1 Enum_CodeProtectAreaSectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprotectareasectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE:
		res = "EDGE"
	case Enum_CodeProtectAreaSectionBaseType_Member_1_END:
		res = "END"
	case Enum_CodeProtectAreaSectionBaseType_Member_1_CL:
		res = "CL"
	}
	return
}

func (enum_codeprotectareasectionbasetype_member_1 *Enum_CodeProtectAreaSectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EDGE":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE
		return
	case "END":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_END
		return
	case "CL":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_CL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprotectareasectionbasetype_member_1 *Enum_CodeProtectAreaSectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE
	case "Enum_CodeProtectAreaSectionBaseType_Member_1_END":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_END
	case "Enum_CodeProtectAreaSectionBaseType_Member_1_CL":
		*enum_codeprotectareasectionbasetype_member_1 = Enum_CodeProtectAreaSectionBaseType_Member_1_CL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprotectareasectionbasetype_member_1 *Enum_CodeProtectAreaSectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprotectareasectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE:
		res = "Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE"
	case Enum_CodeProtectAreaSectionBaseType_Member_1_END:
		res = "Enum_CodeProtectAreaSectionBaseType_Member_1_END"
	case Enum_CodeProtectAreaSectionBaseType_Member_1_CL:
		res = "Enum_CodeProtectAreaSectionBaseType_Member_1_CL"
	}
	return
}

func (enum_codeprotectareasectionbasetype_member_1 Enum_CodeProtectAreaSectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProtectAreaSectionBaseType_Member_1_EDGE")
	res = append(res, "Enum_CodeProtectAreaSectionBaseType_Member_1_END")
	res = append(res, "Enum_CodeProtectAreaSectionBaseType_Member_1_CL")

	return
}

func (enum_codeprotectareasectionbasetype_member_1 Enum_CodeProtectAreaSectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EDGE")
	res = append(res, "END")
	res = append(res, "CL")

	return
}

// Utility function for Enum_CodeRVRReadingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codervrreadingbasetype_member_1 Enum_CodeRVRReadingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codervrreadingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVRReadingBaseType_Member_1_TDZ:
		res = "TDZ"
	case Enum_CodeRVRReadingBaseType_Member_1_MID:
		res = "MID"
	case Enum_CodeRVRReadingBaseType_Member_1_TO:
		res = "TO"
	}
	return
}

func (enum_codervrreadingbasetype_member_1 *Enum_CodeRVRReadingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TDZ":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_TDZ
		return
	case "MID":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_MID
		return
	case "TO":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_TO
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codervrreadingbasetype_member_1 *Enum_CodeRVRReadingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRVRReadingBaseType_Member_1_TDZ":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_TDZ
	case "Enum_CodeRVRReadingBaseType_Member_1_MID":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_MID
	case "Enum_CodeRVRReadingBaseType_Member_1_TO":
		*enum_codervrreadingbasetype_member_1 = Enum_CodeRVRReadingBaseType_Member_1_TO
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codervrreadingbasetype_member_1 *Enum_CodeRVRReadingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codervrreadingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVRReadingBaseType_Member_1_TDZ:
		res = "Enum_CodeRVRReadingBaseType_Member_1_TDZ"
	case Enum_CodeRVRReadingBaseType_Member_1_MID:
		res = "Enum_CodeRVRReadingBaseType_Member_1_MID"
	case Enum_CodeRVRReadingBaseType_Member_1_TO:
		res = "Enum_CodeRVRReadingBaseType_Member_1_TO"
	}
	return
}

func (enum_codervrreadingbasetype_member_1 Enum_CodeRVRReadingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRVRReadingBaseType_Member_1_TDZ")
	res = append(res, "Enum_CodeRVRReadingBaseType_Member_1_MID")
	res = append(res, "Enum_CodeRVRReadingBaseType_Member_1_TO")

	return
}

func (enum_codervrreadingbasetype_member_1 Enum_CodeRVRReadingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TDZ")
	res = append(res, "MID")
	res = append(res, "TO")

	return
}

// Utility function for Enum_CodeRVSMBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codervsmbasetype_member_1 Enum_CodeRVSMBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codervsmbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVSMBaseType_Member_1_RVSM:
		res = "RVSM"
	case Enum_CodeRVSMBaseType_Member_1_NON_RVSM:
		res = "NON_RVSM"
	}
	return
}

func (enum_codervsmbasetype_member_1 *Enum_CodeRVSMBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RVSM":
		*enum_codervsmbasetype_member_1 = Enum_CodeRVSMBaseType_Member_1_RVSM
		return
	case "NON_RVSM":
		*enum_codervsmbasetype_member_1 = Enum_CodeRVSMBaseType_Member_1_NON_RVSM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codervsmbasetype_member_1 *Enum_CodeRVSMBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRVSMBaseType_Member_1_RVSM":
		*enum_codervsmbasetype_member_1 = Enum_CodeRVSMBaseType_Member_1_RVSM
	case "Enum_CodeRVSMBaseType_Member_1_NON_RVSM":
		*enum_codervsmbasetype_member_1 = Enum_CodeRVSMBaseType_Member_1_NON_RVSM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codervsmbasetype_member_1 *Enum_CodeRVSMBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codervsmbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVSMBaseType_Member_1_RVSM:
		res = "Enum_CodeRVSMBaseType_Member_1_RVSM"
	case Enum_CodeRVSMBaseType_Member_1_NON_RVSM:
		res = "Enum_CodeRVSMBaseType_Member_1_NON_RVSM"
	}
	return
}

func (enum_codervsmbasetype_member_1 Enum_CodeRVSMBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRVSMBaseType_Member_1_RVSM")
	res = append(res, "Enum_CodeRVSMBaseType_Member_1_NON_RVSM")

	return
}

func (enum_codervsmbasetype_member_1 Enum_CodeRVSMBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RVSM")
	res = append(res, "NON_RVSM")

	return
}

// Utility function for Enum_CodeRoadBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoadbasetype_member_1 Enum_CodeRoadBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoadbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRoadBaseType_Member_1_SERVICE:
		res = "SERVICE"
	case Enum_CodeRoadBaseType_Member_1_PUBLIC:
		res = "PUBLIC"
	}
	return
}

func (enum_coderoadbasetype_member_1 *Enum_CodeRoadBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SERVICE":
		*enum_coderoadbasetype_member_1 = Enum_CodeRoadBaseType_Member_1_SERVICE
		return
	case "PUBLIC":
		*enum_coderoadbasetype_member_1 = Enum_CodeRoadBaseType_Member_1_PUBLIC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoadbasetype_member_1 *Enum_CodeRoadBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRoadBaseType_Member_1_SERVICE":
		*enum_coderoadbasetype_member_1 = Enum_CodeRoadBaseType_Member_1_SERVICE
	case "Enum_CodeRoadBaseType_Member_1_PUBLIC":
		*enum_coderoadbasetype_member_1 = Enum_CodeRoadBaseType_Member_1_PUBLIC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoadbasetype_member_1 *Enum_CodeRoadBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoadbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRoadBaseType_Member_1_SERVICE:
		res = "Enum_CodeRoadBaseType_Member_1_SERVICE"
	case Enum_CodeRoadBaseType_Member_1_PUBLIC:
		res = "Enum_CodeRoadBaseType_Member_1_PUBLIC"
	}
	return
}

func (enum_coderoadbasetype_member_1 Enum_CodeRoadBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRoadBaseType_Member_1_SERVICE")
	res = append(res, "Enum_CodeRoadBaseType_Member_1_PUBLIC")

	return
}

func (enum_coderoadbasetype_member_1 Enum_CodeRoadBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SERVICE")
	res = append(res, "PUBLIC")

	return
}

// Utility function for Enum_CodeRunwayBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwaybasetype_member_1 Enum_CodeRunwayBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayBaseType_Member_1_RWY:
		res = "RWY"
	case Enum_CodeRunwayBaseType_Member_1_FATO:
		res = "FATO"
	}
	return
}

func (enum_coderunwaybasetype_member_1 *Enum_CodeRunwayBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RWY":
		*enum_coderunwaybasetype_member_1 = Enum_CodeRunwayBaseType_Member_1_RWY
		return
	case "FATO":
		*enum_coderunwaybasetype_member_1 = Enum_CodeRunwayBaseType_Member_1_FATO
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwaybasetype_member_1 *Enum_CodeRunwayBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwayBaseType_Member_1_RWY":
		*enum_coderunwaybasetype_member_1 = Enum_CodeRunwayBaseType_Member_1_RWY
	case "Enum_CodeRunwayBaseType_Member_1_FATO":
		*enum_coderunwaybasetype_member_1 = Enum_CodeRunwayBaseType_Member_1_FATO
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwaybasetype_member_1 *Enum_CodeRunwayBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayBaseType_Member_1_RWY:
		res = "Enum_CodeRunwayBaseType_Member_1_RWY"
	case Enum_CodeRunwayBaseType_Member_1_FATO:
		res = "Enum_CodeRunwayBaseType_Member_1_FATO"
	}
	return
}

func (enum_coderunwaybasetype_member_1 Enum_CodeRunwayBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwayBaseType_Member_1_RWY")
	res = append(res, "Enum_CodeRunwayBaseType_Member_1_FATO")

	return
}

func (enum_coderunwaybasetype_member_1 Enum_CodeRunwayBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RWY")
	res = append(res, "FATO")

	return
}

// Utility function for Enum_CodeRunwayElementBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwayelementbasetype_member_1 Enum_CodeRunwayElementBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwayelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayElementBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION:
		res = "INTERSECTION"
	case Enum_CodeRunwayElementBaseType_Member_1_DISPLACED:
		res = "DISPLACED"
	case Enum_CodeRunwayElementBaseType_Member_1_SHOULDER:
		res = "SHOULDER"
	}
	return
}

func (enum_coderunwayelementbasetype_member_1 *Enum_CodeRunwayElementBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_NORMAL
		return
	case "INTERSECTION":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION
		return
	case "DISPLACED":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_DISPLACED
		return
	case "SHOULDER":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_SHOULDER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwayelementbasetype_member_1 *Enum_CodeRunwayElementBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwayElementBaseType_Member_1_NORMAL":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_NORMAL
	case "Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION
	case "Enum_CodeRunwayElementBaseType_Member_1_DISPLACED":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_DISPLACED
	case "Enum_CodeRunwayElementBaseType_Member_1_SHOULDER":
		*enum_coderunwayelementbasetype_member_1 = Enum_CodeRunwayElementBaseType_Member_1_SHOULDER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwayelementbasetype_member_1 *Enum_CodeRunwayElementBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwayelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayElementBaseType_Member_1_NORMAL:
		res = "Enum_CodeRunwayElementBaseType_Member_1_NORMAL"
	case Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION:
		res = "Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION"
	case Enum_CodeRunwayElementBaseType_Member_1_DISPLACED:
		res = "Enum_CodeRunwayElementBaseType_Member_1_DISPLACED"
	case Enum_CodeRunwayElementBaseType_Member_1_SHOULDER:
		res = "Enum_CodeRunwayElementBaseType_Member_1_SHOULDER"
	}
	return
}

func (enum_coderunwayelementbasetype_member_1 Enum_CodeRunwayElementBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwayElementBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeRunwayElementBaseType_Member_1_INTERSECTION")
	res = append(res, "Enum_CodeRunwayElementBaseType_Member_1_DISPLACED")
	res = append(res, "Enum_CodeRunwayElementBaseType_Member_1_SHOULDER")

	return
}

func (enum_coderunwayelementbasetype_member_1 Enum_CodeRunwayElementBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "INTERSECTION")
	res = append(res, "DISPLACED")
	res = append(res, "SHOULDER")

	return
}

// Utility function for Enum_CodeRunwayMarkingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwaymarkingbasetype_member_1 Enum_CodeRunwayMarkingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwaymarkingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION:
		res = "PRECISION"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION:
		res = "NONPRECISION"
	case Enum_CodeRunwayMarkingBaseType_Member_1_BASIC:
		res = "BASIC"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NONE:
		res = "NONE"
	case Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS:
		res = "RUNWAY_NUMBERS"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD:
		res = "NON_STANDARD"
	case Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT:
		res = "HELIPORT"
	}
	return
}

func (enum_coderunwaymarkingbasetype_member_1 *Enum_CodeRunwayMarkingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PRECISION":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION
		return
	case "NONPRECISION":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION
		return
	case "BASIC":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_BASIC
		return
	case "NONE":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NONE
		return
	case "RUNWAY_NUMBERS":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS
		return
	case "NON_STANDARD":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD
		return
	case "HELIPORT":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwaymarkingbasetype_member_1 *Enum_CodeRunwayMarkingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION
	case "Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION
	case "Enum_CodeRunwayMarkingBaseType_Member_1_BASIC":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_BASIC
	case "Enum_CodeRunwayMarkingBaseType_Member_1_NONE":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NONE
	case "Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS
	case "Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD
	case "Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT":
		*enum_coderunwaymarkingbasetype_member_1 = Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwaymarkingbasetype_member_1 *Enum_CodeRunwayMarkingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwaymarkingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION"
	case Enum_CodeRunwayMarkingBaseType_Member_1_BASIC:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_BASIC"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NONE:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_NONE"
	case Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS"
	case Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD"
	case Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT:
		res = "Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT"
	}
	return
}

func (enum_coderunwaymarkingbasetype_member_1 Enum_CodeRunwayMarkingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_PRECISION")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_NONPRECISION")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_BASIC")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_NONE")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_RUNWAY_NUMBERS")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_NON_STANDARD")
	res = append(res, "Enum_CodeRunwayMarkingBaseType_Member_1_HELIPORT")

	return
}

func (enum_coderunwaymarkingbasetype_member_1 Enum_CodeRunwayMarkingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PRECISION")
	res = append(res, "NONPRECISION")
	res = append(res, "BASIC")
	res = append(res, "NONE")
	res = append(res, "RUNWAY_NUMBERS")
	res = append(res, "NON_STANDARD")
	res = append(res, "HELIPORT")

	return
}

// Utility function for Enum_CodeRunwayPointRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwaypointrolebasetype_member_1 Enum_CodeRunwayPointRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwaypointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayPointRoleBaseType_Member_1_START:
		res = "START"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_THR:
		res = "THR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR:
		res = "DISTHR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ:
		res = "TDZ"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_MID:
		res = "MID"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_END:
		res = "END"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN:
		res = "START_RUN"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO:
		res = "LAHSO"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE:
		res = "ABEAM_GLIDESLOPE"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR:
		res = "ABEAM_PAR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION:
		res = "ABEAM_ELEVATION"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR:
		res = "ABEAM_TDR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER:
		res = "ABEAM_RER"
	}
	return
}

func (enum_coderunwaypointrolebasetype_member_1 *Enum_CodeRunwayPointRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "START":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_START
		return
	case "THR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_THR
		return
	case "DISTHR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR
		return
	case "TDZ":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ
		return
	case "MID":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_MID
		return
	case "END":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_END
		return
	case "START_RUN":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN
		return
	case "LAHSO":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO
		return
	case "ABEAM_GLIDESLOPE":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE
		return
	case "ABEAM_PAR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR
		return
	case "ABEAM_ELEVATION":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION
		return
	case "ABEAM_TDR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR
		return
	case "ABEAM_RER":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwaypointrolebasetype_member_1 *Enum_CodeRunwayPointRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_START":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_START
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_THR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_THR
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_MID":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_MID
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_END":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_END
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR
	case "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER":
		*enum_coderunwaypointrolebasetype_member_1 = Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwaypointrolebasetype_member_1 *Enum_CodeRunwayPointRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwaypointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayPointRoleBaseType_Member_1_START:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_START"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_THR:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_THR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_MID:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_MID"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_END:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_END"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR"
	case Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER:
		res = "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER"
	}
	return
}

func (enum_coderunwaypointrolebasetype_member_1 Enum_CodeRunwayPointRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_START")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_THR")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_DISTHR")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_TDZ")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_MID")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_END")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_START_RUN")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_LAHSO")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_GLIDESLOPE")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_PAR")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_ELEVATION")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_TDR")
	res = append(res, "Enum_CodeRunwayPointRoleBaseType_Member_1_ABEAM_RER")

	return
}

func (enum_coderunwaypointrolebasetype_member_1 Enum_CodeRunwayPointRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "START")
	res = append(res, "THR")
	res = append(res, "DISTHR")
	res = append(res, "TDZ")
	res = append(res, "MID")
	res = append(res, "END")
	res = append(res, "START_RUN")
	res = append(res, "LAHSO")
	res = append(res, "ABEAM_GLIDESLOPE")
	res = append(res, "ABEAM_PAR")
	res = append(res, "ABEAM_ELEVATION")
	res = append(res, "ABEAM_TDR")
	res = append(res, "ABEAM_RER")

	return
}

// Utility function for Enum_CodeRunwayProtectionAreaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwayprotectionareabasetype_member_1 Enum_CodeRunwayProtectionAreaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwayprotectionareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY:
		res = "CWY"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA:
		res = "RESA"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ:
		res = "OFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ:
		res = "IOFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ:
		res = "POFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS:
		res = "ILS"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI:
		res = "VGSI"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY:
		res = "STOPWAY"
	}
	return
}

func (enum_coderunwayprotectionareabasetype_member_1 *Enum_CodeRunwayProtectionAreaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CWY":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY
		return
	case "RESA":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA
		return
	case "OFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ
		return
	case "IOFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ
		return
	case "POFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ
		return
	case "ILS":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS
		return
	case "VGSI":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI
		return
	case "STOPWAY":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwayprotectionareabasetype_member_1 *Enum_CodeRunwayProtectionAreaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI
	case "Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY":
		*enum_coderunwayprotectionareabasetype_member_1 = Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwayprotectionareabasetype_member_1 *Enum_CodeRunwayProtectionAreaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwayprotectionareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI"
	case Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY:
		res = "Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY"
	}
	return
}

func (enum_coderunwayprotectionareabasetype_member_1 Enum_CodeRunwayProtectionAreaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_CWY")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_RESA")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_OFZ")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_IOFZ")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_POFZ")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_ILS")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_VGSI")
	res = append(res, "Enum_CodeRunwayProtectionAreaBaseType_Member_1_STOPWAY")

	return
}

func (enum_coderunwayprotectionareabasetype_member_1 Enum_CodeRunwayProtectionAreaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CWY")
	res = append(res, "RESA")
	res = append(res, "OFZ")
	res = append(res, "IOFZ")
	res = append(res, "POFZ")
	res = append(res, "ILS")
	res = append(res, "VGSI")
	res = append(res, "STOPWAY")

	return
}

// Utility function for Enum_CodeRunwaySectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderunwaysectionbasetype_member_1 Enum_CodeRunwaySectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderunwaysectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwaySectionBaseType_Member_1_TDZ:
		res = "TDZ"
	case Enum_CodeRunwaySectionBaseType_Member_1_AIM:
		res = "AIM"
	case Enum_CodeRunwaySectionBaseType_Member_1_CL:
		res = "CL"
	case Enum_CodeRunwaySectionBaseType_Member_1_EDGE:
		res = "EDGE"
	case Enum_CodeRunwaySectionBaseType_Member_1_THR:
		res = "THR"
	case Enum_CodeRunwaySectionBaseType_Member_1_DESIG:
		res = "DESIG"
	case Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR:
		res = "AFT_THR"
	case Enum_CodeRunwaySectionBaseType_Member_1_DTHR:
		res = "DTHR"
	case Enum_CodeRunwaySectionBaseType_Member_1_END:
		res = "END"
	case Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT:
		res = "TWY_INT"
	case Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT:
		res = "RPD_TWY_INT"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD:
		res = "1_THIRD"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD:
		res = "2_THIRD"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD:
		res = "3_THIRD"
	}
	return
}

func (enum_coderunwaysectionbasetype_member_1 *Enum_CodeRunwaySectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TDZ":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_TDZ
		return
	case "AIM":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_AIM
		return
	case "CL":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_CL
		return
	case "EDGE":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_EDGE
		return
	case "THR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_THR
		return
	case "DESIG":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_DESIG
		return
	case "AFT_THR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR
		return
	case "DTHR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_DTHR
		return
	case "END":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_END
		return
	case "TWY_INT":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT
		return
	case "RPD_TWY_INT":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT
		return
	case "1_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD
		return
	case "2_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD
		return
	case "3_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderunwaysectionbasetype_member_1 *Enum_CodeRunwaySectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRunwaySectionBaseType_Member_1_TDZ":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_TDZ
	case "Enum_CodeRunwaySectionBaseType_Member_1_AIM":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_AIM
	case "Enum_CodeRunwaySectionBaseType_Member_1_CL":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_CL
	case "Enum_CodeRunwaySectionBaseType_Member_1_EDGE":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_EDGE
	case "Enum_CodeRunwaySectionBaseType_Member_1_THR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_THR
	case "Enum_CodeRunwaySectionBaseType_Member_1_DESIG":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_DESIG
	case "Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR
	case "Enum_CodeRunwaySectionBaseType_Member_1_DTHR":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_DTHR
	case "Enum_CodeRunwaySectionBaseType_Member_1_END":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_END
	case "Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT
	case "Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT
	case "Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD
	case "Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD
	case "Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD":
		*enum_coderunwaysectionbasetype_member_1 = Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderunwaysectionbasetype_member_1 *Enum_CodeRunwaySectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderunwaysectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRunwaySectionBaseType_Member_1_TDZ:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_TDZ"
	case Enum_CodeRunwaySectionBaseType_Member_1_AIM:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_AIM"
	case Enum_CodeRunwaySectionBaseType_Member_1_CL:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_CL"
	case Enum_CodeRunwaySectionBaseType_Member_1_EDGE:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_EDGE"
	case Enum_CodeRunwaySectionBaseType_Member_1_THR:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_THR"
	case Enum_CodeRunwaySectionBaseType_Member_1_DESIG:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_DESIG"
	case Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR"
	case Enum_CodeRunwaySectionBaseType_Member_1_DTHR:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_DTHR"
	case Enum_CodeRunwaySectionBaseType_Member_1_END:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_END"
	case Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT"
	case Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD"
	case Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD:
		res = "Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD"
	}
	return
}

func (enum_coderunwaysectionbasetype_member_1 Enum_CodeRunwaySectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_TDZ")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_AIM")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_CL")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_EDGE")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_THR")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_DESIG")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_AFT_THR")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_DTHR")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_END")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_TWY_INT")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_RPD_TWY_INT")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_N_1_THIRD")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_N_2_THIRD")
	res = append(res, "Enum_CodeRunwaySectionBaseType_Member_1_N_3_THIRD")

	return
}

func (enum_coderunwaysectionbasetype_member_1 Enum_CodeRunwaySectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TDZ")
	res = append(res, "AIM")
	res = append(res, "CL")
	res = append(res, "EDGE")
	res = append(res, "THR")
	res = append(res, "DESIG")
	res = append(res, "AFT_THR")
	res = append(res, "DTHR")
	res = append(res, "END")
	res = append(res, "TWY_INT")
	res = append(res, "RPD_TWY_INT")
	res = append(res, "1_THIRD")
	res = append(res, "2_THIRD")
	res = append(res, "3_THIRD")

	return
}

// Utility function for Enum_CodeSideBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesidebasetype_member_1 Enum_CodeSideBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesidebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSideBaseType_Member_1_LEFT:
		res = "LEFT"
	case Enum_CodeSideBaseType_Member_1_RIGHT:
		res = "RIGHT"
	case Enum_CodeSideBaseType_Member_1_BOTH:
		res = "BOTH"
	}
	return
}

func (enum_codesidebasetype_member_1 *Enum_CodeSideBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LEFT":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_LEFT
		return
	case "RIGHT":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_RIGHT
		return
	case "BOTH":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_BOTH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesidebasetype_member_1 *Enum_CodeSideBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSideBaseType_Member_1_LEFT":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_LEFT
	case "Enum_CodeSideBaseType_Member_1_RIGHT":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_RIGHT
	case "Enum_CodeSideBaseType_Member_1_BOTH":
		*enum_codesidebasetype_member_1 = Enum_CodeSideBaseType_Member_1_BOTH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesidebasetype_member_1 *Enum_CodeSideBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesidebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSideBaseType_Member_1_LEFT:
		res = "Enum_CodeSideBaseType_Member_1_LEFT"
	case Enum_CodeSideBaseType_Member_1_RIGHT:
		res = "Enum_CodeSideBaseType_Member_1_RIGHT"
	case Enum_CodeSideBaseType_Member_1_BOTH:
		res = "Enum_CodeSideBaseType_Member_1_BOTH"
	}
	return
}

func (enum_codesidebasetype_member_1 Enum_CodeSideBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSideBaseType_Member_1_LEFT")
	res = append(res, "Enum_CodeSideBaseType_Member_1_RIGHT")
	res = append(res, "Enum_CodeSideBaseType_Member_1_BOTH")

	return
}

func (enum_codesidebasetype_member_1 Enum_CodeSideBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LEFT")
	res = append(res, "RIGHT")
	res = append(res, "BOTH")

	return
}

// Utility function for Enum_CodeStatusAirportBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusairportbasetype_member_1 Enum_CodeStatusAirportBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusairportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusAirportBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeStatusAirportBaseType_Member_1_LIMITED:
		res = "LIMITED"
	case Enum_CodeStatusAirportBaseType_Member_1_CLOSED:
		res = "CLOSED"
	}
	return
}

func (enum_codestatusairportbasetype_member_1 *Enum_CodeStatusAirportBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_NORMAL
		return
	case "LIMITED":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_LIMITED
		return
	case "CLOSED":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_CLOSED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusairportbasetype_member_1 *Enum_CodeStatusAirportBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusAirportBaseType_Member_1_NORMAL":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_NORMAL
	case "Enum_CodeStatusAirportBaseType_Member_1_LIMITED":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_LIMITED
	case "Enum_CodeStatusAirportBaseType_Member_1_CLOSED":
		*enum_codestatusairportbasetype_member_1 = Enum_CodeStatusAirportBaseType_Member_1_CLOSED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusairportbasetype_member_1 *Enum_CodeStatusAirportBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusairportbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusAirportBaseType_Member_1_NORMAL:
		res = "Enum_CodeStatusAirportBaseType_Member_1_NORMAL"
	case Enum_CodeStatusAirportBaseType_Member_1_LIMITED:
		res = "Enum_CodeStatusAirportBaseType_Member_1_LIMITED"
	case Enum_CodeStatusAirportBaseType_Member_1_CLOSED:
		res = "Enum_CodeStatusAirportBaseType_Member_1_CLOSED"
	}
	return
}

func (enum_codestatusairportbasetype_member_1 Enum_CodeStatusAirportBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusAirportBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeStatusAirportBaseType_Member_1_LIMITED")
	res = append(res, "Enum_CodeStatusAirportBaseType_Member_1_CLOSED")

	return
}

func (enum_codestatusairportbasetype_member_1 Enum_CodeStatusAirportBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "LIMITED")
	res = append(res, "CLOSED")

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

// Utility function for Enum_CodeSurfaceCompositionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesurfacecompositionbasetype_member_1 Enum_CodeSurfaceCompositionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesurfacecompositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH:
		res = "ASPH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS:
		res = "ASPH_GRASS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC:
		res = "CONC"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH:
		res = "CONC_ASPH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS:
		res = "CONC_GRS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS:
		res = "GRASS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_SAND:
		res = "SAND"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_WATER:
		res = "WATER"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM:
		res = "BITUM"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK:
		res = "BRICK"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM:
		res = "MACADAM"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_STONE:
		res = "STONE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL:
		res = "CORAL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY:
		res = "CLAY"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE:
		res = "LATERITE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL:
		res = "GRAVEL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH:
		res = "EARTH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ICE:
		res = "ICE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW:
		res = "SNOW"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE:
		res = "MEMBRANE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_METAL:
		res = "METAL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MATS:
		res = "MATS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL:
		res = "PIERCED_STEEL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD:
		res = "WOOD"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX:
		res = "NON_BITUM_MIX"
	}
	return
}

func (enum_codesurfacecompositionbasetype_member_1 *Enum_CodeSurfaceCompositionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ASPH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH
		return
	case "ASPH_GRASS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS
		return
	case "CONC":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC
		return
	case "CONC_ASPH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH
		return
	case "CONC_GRS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS
		return
	case "GRASS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS
		return
	case "SAND":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_SAND
		return
	case "WATER":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_WATER
		return
	case "BITUM":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM
		return
	case "BRICK":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK
		return
	case "MACADAM":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM
		return
	case "STONE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_STONE
		return
	case "CORAL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL
		return
	case "CLAY":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY
		return
	case "LATERITE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE
		return
	case "GRAVEL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL
		return
	case "EARTH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH
		return
	case "ICE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ICE
		return
	case "SNOW":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW
		return
	case "MEMBRANE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE
		return
	case "METAL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_METAL
		return
	case "MATS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MATS
		return
	case "PIERCED_STEEL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL
		return
	case "WOOD":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD
		return
	case "NON_BITUM_MIX":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesurfacecompositionbasetype_member_1 *Enum_CodeSurfaceCompositionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_SAND":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_SAND
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_WATER":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_WATER
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_STONE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_STONE
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_ICE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_ICE
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_METAL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_METAL
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_MATS":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_MATS
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD
	case "Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX":
		*enum_codesurfacecompositionbasetype_member_1 = Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesurfacecompositionbasetype_member_1 *Enum_CodeSurfaceCompositionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesurfacecompositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_SAND:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_SAND"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_WATER:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_WATER"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_STONE:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_STONE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_ICE:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_ICE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_METAL:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_METAL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_MATS:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_MATS"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD"
	case Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX:
		res = "Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX"
	}
	return
}

func (enum_codesurfacecompositionbasetype_member_1 Enum_CodeSurfaceCompositionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_ASPH_GRASS")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_ASPH")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_CONC_GRS")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_GRASS")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_SAND")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_WATER")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_BITUM")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_BRICK")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_MACADAM")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_STONE")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_CORAL")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_CLAY")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_LATERITE")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_GRAVEL")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_EARTH")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_ICE")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_SNOW")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_MEMBRANE")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_METAL")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_MATS")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_PIERCED_STEEL")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_WOOD")
	res = append(res, "Enum_CodeSurfaceCompositionBaseType_Member_1_NON_BITUM_MIX")

	return
}

func (enum_codesurfacecompositionbasetype_member_1 Enum_CodeSurfaceCompositionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ASPH")
	res = append(res, "ASPH_GRASS")
	res = append(res, "CONC")
	res = append(res, "CONC_ASPH")
	res = append(res, "CONC_GRS")
	res = append(res, "GRASS")
	res = append(res, "SAND")
	res = append(res, "WATER")
	res = append(res, "BITUM")
	res = append(res, "BRICK")
	res = append(res, "MACADAM")
	res = append(res, "STONE")
	res = append(res, "CORAL")
	res = append(res, "CLAY")
	res = append(res, "LATERITE")
	res = append(res, "GRAVEL")
	res = append(res, "EARTH")
	res = append(res, "ICE")
	res = append(res, "SNOW")
	res = append(res, "MEMBRANE")
	res = append(res, "METAL")
	res = append(res, "MATS")
	res = append(res, "PIERCED_STEEL")
	res = append(res, "WOOD")
	res = append(res, "NON_BITUM_MIX")

	return
}

// Utility function for Enum_CodeSurfaceConditionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesurfaceconditionbasetype_member_1 Enum_CodeSurfaceConditionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesurfaceconditionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfaceConditionBaseType_Member_1_GOOD:
		res = "GOOD"
	case Enum_CodeSurfaceConditionBaseType_Member_1_FAIR:
		res = "FAIR"
	case Enum_CodeSurfaceConditionBaseType_Member_1_POOR:
		res = "POOR"
	case Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE:
		res = "UNSAFE"
	case Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED:
		res = "DEFORMED"
	}
	return
}

func (enum_codesurfaceconditionbasetype_member_1 *Enum_CodeSurfaceConditionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GOOD":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_GOOD
		return
	case "FAIR":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_FAIR
		return
	case "POOR":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_POOR
		return
	case "UNSAFE":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE
		return
	case "DEFORMED":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesurfaceconditionbasetype_member_1 *Enum_CodeSurfaceConditionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSurfaceConditionBaseType_Member_1_GOOD":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_GOOD
	case "Enum_CodeSurfaceConditionBaseType_Member_1_FAIR":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_FAIR
	case "Enum_CodeSurfaceConditionBaseType_Member_1_POOR":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_POOR
	case "Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE
	case "Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED":
		*enum_codesurfaceconditionbasetype_member_1 = Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesurfaceconditionbasetype_member_1 *Enum_CodeSurfaceConditionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesurfaceconditionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfaceConditionBaseType_Member_1_GOOD:
		res = "Enum_CodeSurfaceConditionBaseType_Member_1_GOOD"
	case Enum_CodeSurfaceConditionBaseType_Member_1_FAIR:
		res = "Enum_CodeSurfaceConditionBaseType_Member_1_FAIR"
	case Enum_CodeSurfaceConditionBaseType_Member_1_POOR:
		res = "Enum_CodeSurfaceConditionBaseType_Member_1_POOR"
	case Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE:
		res = "Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE"
	case Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED:
		res = "Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED"
	}
	return
}

func (enum_codesurfaceconditionbasetype_member_1 Enum_CodeSurfaceConditionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSurfaceConditionBaseType_Member_1_GOOD")
	res = append(res, "Enum_CodeSurfaceConditionBaseType_Member_1_FAIR")
	res = append(res, "Enum_CodeSurfaceConditionBaseType_Member_1_POOR")
	res = append(res, "Enum_CodeSurfaceConditionBaseType_Member_1_UNSAFE")
	res = append(res, "Enum_CodeSurfaceConditionBaseType_Member_1_DEFORMED")

	return
}

func (enum_codesurfaceconditionbasetype_member_1 Enum_CodeSurfaceConditionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "GOOD")
	res = append(res, "FAIR")
	res = append(res, "POOR")
	res = append(res, "UNSAFE")
	res = append(res, "DEFORMED")

	return
}

// Utility function for Enum_CodeSurfacePreparationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesurfacepreparationbasetype_member_1 Enum_CodeSurfacePreparationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesurfacepreparationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL:
		res = "NATURAL"
	case Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED:
		res = "ROLLED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED:
		res = "COMPACTED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_GRADED:
		res = "GRADED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED:
		res = "GROOVED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_OILED:
		res = "OILED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_PAVED:
		res = "PAVED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_PFC:
		res = "PFC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_AFSC:
		res = "AFSC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_RFSC:
		res = "RFSC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED:
		res = "NON_GROOVED"
	}
	return
}

func (enum_codesurfacepreparationbasetype_member_1 *Enum_CodeSurfacePreparationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NATURAL":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL
		return
	case "ROLLED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED
		return
	case "COMPACTED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED
		return
	case "GRADED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_GRADED
		return
	case "GROOVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED
		return
	case "OILED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_OILED
		return
	case "PAVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_PAVED
		return
	case "PFC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_PFC
		return
	case "AFSC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_AFSC
		return
	case "RFSC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_RFSC
		return
	case "NON_GROOVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesurfacepreparationbasetype_member_1 *Enum_CodeSurfacePreparationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL
	case "Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_GRADED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_GRADED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_OILED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_OILED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_PAVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_PAVED
	case "Enum_CodeSurfacePreparationBaseType_Member_1_PFC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_PFC
	case "Enum_CodeSurfacePreparationBaseType_Member_1_AFSC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_AFSC
	case "Enum_CodeSurfacePreparationBaseType_Member_1_RFSC":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_RFSC
	case "Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED":
		*enum_codesurfacepreparationbasetype_member_1 = Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesurfacepreparationbasetype_member_1 *Enum_CodeSurfacePreparationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesurfacepreparationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL"
	case Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_GRADED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_GRADED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_OILED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_OILED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_PAVED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_PAVED"
	case Enum_CodeSurfacePreparationBaseType_Member_1_PFC:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_PFC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_AFSC:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_AFSC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_RFSC:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_RFSC"
	case Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED:
		res = "Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED"
	}
	return
}

func (enum_codesurfacepreparationbasetype_member_1 Enum_CodeSurfacePreparationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_NATURAL")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_ROLLED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_COMPACTED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_GRADED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_GROOVED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_OILED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_PAVED")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_PFC")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_AFSC")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_RFSC")
	res = append(res, "Enum_CodeSurfacePreparationBaseType_Member_1_NON_GROOVED")

	return
}

func (enum_codesurfacepreparationbasetype_member_1 Enum_CodeSurfacePreparationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NATURAL")
	res = append(res, "ROLLED")
	res = append(res, "COMPACTED")
	res = append(res, "GRADED")
	res = append(res, "GROOVED")
	res = append(res, "OILED")
	res = append(res, "PAVED")
	res = append(res, "PFC")
	res = append(res, "AFSC")
	res = append(res, "RFSC")
	res = append(res, "NON_GROOVED")

	return
}

// Utility function for Enum_CodeTLOFSectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetlofsectionbasetype_member_1 Enum_CodeTLOFSectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetlofsectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTLOFSectionBaseType_Member_1_AIM:
		res = "AIM"
	case Enum_CodeTLOFSectionBaseType_Member_1_EDGE:
		res = "EDGE"
	}
	return
}

func (enum_codetlofsectionbasetype_member_1 *Enum_CodeTLOFSectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AIM":
		*enum_codetlofsectionbasetype_member_1 = Enum_CodeTLOFSectionBaseType_Member_1_AIM
		return
	case "EDGE":
		*enum_codetlofsectionbasetype_member_1 = Enum_CodeTLOFSectionBaseType_Member_1_EDGE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetlofsectionbasetype_member_1 *Enum_CodeTLOFSectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTLOFSectionBaseType_Member_1_AIM":
		*enum_codetlofsectionbasetype_member_1 = Enum_CodeTLOFSectionBaseType_Member_1_AIM
	case "Enum_CodeTLOFSectionBaseType_Member_1_EDGE":
		*enum_codetlofsectionbasetype_member_1 = Enum_CodeTLOFSectionBaseType_Member_1_EDGE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetlofsectionbasetype_member_1 *Enum_CodeTLOFSectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetlofsectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTLOFSectionBaseType_Member_1_AIM:
		res = "Enum_CodeTLOFSectionBaseType_Member_1_AIM"
	case Enum_CodeTLOFSectionBaseType_Member_1_EDGE:
		res = "Enum_CodeTLOFSectionBaseType_Member_1_EDGE"
	}
	return
}

func (enum_codetlofsectionbasetype_member_1 Enum_CodeTLOFSectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTLOFSectionBaseType_Member_1_AIM")
	res = append(res, "Enum_CodeTLOFSectionBaseType_Member_1_EDGE")

	return
}

func (enum_codetlofsectionbasetype_member_1 Enum_CodeTLOFSectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AIM")
	res = append(res, "EDGE")

	return
}

// Utility function for Enum_CodeTaxiwayBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetaxiwaybasetype_member_1 Enum_CodeTaxiwayBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetaxiwaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwayBaseType_Member_1_AIR:
		res = "AIR"
	case Enum_CodeTaxiwayBaseType_Member_1_GND:
		res = "GND"
	case Enum_CodeTaxiwayBaseType_Member_1_EXIT:
		res = "EXIT"
	case Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT:
		res = "FASTEXIT"
	case Enum_CodeTaxiwayBaseType_Member_1_STUB:
		res = "STUB"
	case Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND:
		res = "TURN_AROUND"
	case Enum_CodeTaxiwayBaseType_Member_1_PARALLEL:
		res = "PARALLEL"
	case Enum_CodeTaxiwayBaseType_Member_1_BYPASS:
		res = "BYPASS"
	}
	return
}

func (enum_codetaxiwaybasetype_member_1 *Enum_CodeTaxiwayBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AIR":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_AIR
		return
	case "GND":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_GND
		return
	case "EXIT":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_EXIT
		return
	case "FASTEXIT":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT
		return
	case "STUB":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_STUB
		return
	case "TURN_AROUND":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND
		return
	case "PARALLEL":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_PARALLEL
		return
	case "BYPASS":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_BYPASS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetaxiwaybasetype_member_1 *Enum_CodeTaxiwayBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTaxiwayBaseType_Member_1_AIR":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_AIR
	case "Enum_CodeTaxiwayBaseType_Member_1_GND":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_GND
	case "Enum_CodeTaxiwayBaseType_Member_1_EXIT":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_EXIT
	case "Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT
	case "Enum_CodeTaxiwayBaseType_Member_1_STUB":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_STUB
	case "Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND
	case "Enum_CodeTaxiwayBaseType_Member_1_PARALLEL":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_PARALLEL
	case "Enum_CodeTaxiwayBaseType_Member_1_BYPASS":
		*enum_codetaxiwaybasetype_member_1 = Enum_CodeTaxiwayBaseType_Member_1_BYPASS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetaxiwaybasetype_member_1 *Enum_CodeTaxiwayBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetaxiwaybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwayBaseType_Member_1_AIR:
		res = "Enum_CodeTaxiwayBaseType_Member_1_AIR"
	case Enum_CodeTaxiwayBaseType_Member_1_GND:
		res = "Enum_CodeTaxiwayBaseType_Member_1_GND"
	case Enum_CodeTaxiwayBaseType_Member_1_EXIT:
		res = "Enum_CodeTaxiwayBaseType_Member_1_EXIT"
	case Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT:
		res = "Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT"
	case Enum_CodeTaxiwayBaseType_Member_1_STUB:
		res = "Enum_CodeTaxiwayBaseType_Member_1_STUB"
	case Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND:
		res = "Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND"
	case Enum_CodeTaxiwayBaseType_Member_1_PARALLEL:
		res = "Enum_CodeTaxiwayBaseType_Member_1_PARALLEL"
	case Enum_CodeTaxiwayBaseType_Member_1_BYPASS:
		res = "Enum_CodeTaxiwayBaseType_Member_1_BYPASS"
	}
	return
}

func (enum_codetaxiwaybasetype_member_1 Enum_CodeTaxiwayBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_AIR")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_GND")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_EXIT")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_FASTEXIT")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_STUB")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_TURN_AROUND")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_PARALLEL")
	res = append(res, "Enum_CodeTaxiwayBaseType_Member_1_BYPASS")

	return
}

func (enum_codetaxiwaybasetype_member_1 Enum_CodeTaxiwayBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AIR")
	res = append(res, "GND")
	res = append(res, "EXIT")
	res = append(res, "FASTEXIT")
	res = append(res, "STUB")
	res = append(res, "TURN_AROUND")
	res = append(res, "PARALLEL")
	res = append(res, "BYPASS")

	return
}

// Utility function for Enum_CodeTaxiwayElementBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetaxiwayelementbasetype_member_1 Enum_CodeTaxiwayElementBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetaxiwayelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION:
		res = "INTERSECTION"
	case Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER:
		res = "SHOULDER"
	case Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY:
		res = "HOLDING_BAY"
	}
	return
}

func (enum_codetaxiwayelementbasetype_member_1 *Enum_CodeTaxiwayElementBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL
		return
	case "INTERSECTION":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION
		return
	case "SHOULDER":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER
		return
	case "HOLDING_BAY":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetaxiwayelementbasetype_member_1 *Enum_CodeTaxiwayElementBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL
	case "Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION
	case "Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER
	case "Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY":
		*enum_codetaxiwayelementbasetype_member_1 = Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetaxiwayelementbasetype_member_1 *Enum_CodeTaxiwayElementBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetaxiwayelementbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL:
		res = "Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL"
	case Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION:
		res = "Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION"
	case Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER:
		res = "Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER"
	case Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY:
		res = "Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY"
	}
	return
}

func (enum_codetaxiwayelementbasetype_member_1 Enum_CodeTaxiwayElementBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTaxiwayElementBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeTaxiwayElementBaseType_Member_1_INTERSECTION")
	res = append(res, "Enum_CodeTaxiwayElementBaseType_Member_1_SHOULDER")
	res = append(res, "Enum_CodeTaxiwayElementBaseType_Member_1_HOLDING_BAY")

	return
}

func (enum_codetaxiwayelementbasetype_member_1 Enum_CodeTaxiwayElementBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "INTERSECTION")
	res = append(res, "SHOULDER")
	res = append(res, "HOLDING_BAY")

	return
}

// Utility function for Enum_CodeTaxiwaySectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetaxiwaysectionbasetype_member_1 Enum_CodeTaxiwaySectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetaxiwaysectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwaySectionBaseType_Member_1_CL:
		res = "CL"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE:
		res = "EDGE"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_END:
		res = "END"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT:
		res = "RWY_INT"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT:
		res = "TWY_INT"
	}
	return
}

func (enum_codetaxiwaysectionbasetype_member_1 *Enum_CodeTaxiwaySectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CL":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_CL
		return
	case "EDGE":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE
		return
	case "END":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_END
		return
	case "RWY_INT":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT
		return
	case "TWY_INT":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetaxiwaysectionbasetype_member_1 *Enum_CodeTaxiwaySectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTaxiwaySectionBaseType_Member_1_CL":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_CL
	case "Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE
	case "Enum_CodeTaxiwaySectionBaseType_Member_1_END":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_END
	case "Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT
	case "Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT":
		*enum_codetaxiwaysectionbasetype_member_1 = Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetaxiwaysectionbasetype_member_1 *Enum_CodeTaxiwaySectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetaxiwaysectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTaxiwaySectionBaseType_Member_1_CL:
		res = "Enum_CodeTaxiwaySectionBaseType_Member_1_CL"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE:
		res = "Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_END:
		res = "Enum_CodeTaxiwaySectionBaseType_Member_1_END"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT:
		res = "Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT"
	case Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT:
		res = "Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT"
	}
	return
}

func (enum_codetaxiwaysectionbasetype_member_1 Enum_CodeTaxiwaySectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTaxiwaySectionBaseType_Member_1_CL")
	res = append(res, "Enum_CodeTaxiwaySectionBaseType_Member_1_EDGE")
	res = append(res, "Enum_CodeTaxiwaySectionBaseType_Member_1_END")
	res = append(res, "Enum_CodeTaxiwaySectionBaseType_Member_1_RWY_INT")
	res = append(res, "Enum_CodeTaxiwaySectionBaseType_Member_1_TWY_INT")

	return
}

func (enum_codetaxiwaysectionbasetype_member_1 Enum_CodeTaxiwaySectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CL")
	res = append(res, "EDGE")
	res = append(res, "END")
	res = append(res, "RWY_INT")
	res = append(res, "TWY_INT")

	return
}

// Utility function for Enum_CodeTelecomNetworkBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetelecomnetworkbasetype_member_1 Enum_CodeTelecomNetworkBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetelecomnetworkbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTelecomNetworkBaseType_Member_1_AFTN:
		res = "AFTN"
	case Enum_CodeTelecomNetworkBaseType_Member_1_AMHS:
		res = "AMHS"
	case Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET:
		res = "INTERNET"
	case Enum_CodeTelecomNetworkBaseType_Member_1_SITA:
		res = "SITA"
	case Enum_CodeTelecomNetworkBaseType_Member_1_ACARS:
		res = "ACARS"
	case Enum_CodeTelecomNetworkBaseType_Member_1_ADNS:
		res = "ADNS"
	}
	return
}

func (enum_codetelecomnetworkbasetype_member_1 *Enum_CodeTelecomNetworkBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AFTN":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_AFTN
		return
	case "AMHS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_AMHS
		return
	case "INTERNET":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET
		return
	case "SITA":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_SITA
		return
	case "ACARS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_ACARS
		return
	case "ADNS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_ADNS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetelecomnetworkbasetype_member_1 *Enum_CodeTelecomNetworkBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTelecomNetworkBaseType_Member_1_AFTN":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_AFTN
	case "Enum_CodeTelecomNetworkBaseType_Member_1_AMHS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_AMHS
	case "Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET
	case "Enum_CodeTelecomNetworkBaseType_Member_1_SITA":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_SITA
	case "Enum_CodeTelecomNetworkBaseType_Member_1_ACARS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_ACARS
	case "Enum_CodeTelecomNetworkBaseType_Member_1_ADNS":
		*enum_codetelecomnetworkbasetype_member_1 = Enum_CodeTelecomNetworkBaseType_Member_1_ADNS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetelecomnetworkbasetype_member_1 *Enum_CodeTelecomNetworkBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetelecomnetworkbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTelecomNetworkBaseType_Member_1_AFTN:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_AFTN"
	case Enum_CodeTelecomNetworkBaseType_Member_1_AMHS:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_AMHS"
	case Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET"
	case Enum_CodeTelecomNetworkBaseType_Member_1_SITA:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_SITA"
	case Enum_CodeTelecomNetworkBaseType_Member_1_ACARS:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_ACARS"
	case Enum_CodeTelecomNetworkBaseType_Member_1_ADNS:
		res = "Enum_CodeTelecomNetworkBaseType_Member_1_ADNS"
	}
	return
}

func (enum_codetelecomnetworkbasetype_member_1 Enum_CodeTelecomNetworkBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_AFTN")
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_AMHS")
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_INTERNET")
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_SITA")
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_ACARS")
	res = append(res, "Enum_CodeTelecomNetworkBaseType_Member_1_ADNS")

	return
}

func (enum_codetelecomnetworkbasetype_member_1 Enum_CodeTelecomNetworkBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AFTN")
	res = append(res, "AMHS")
	res = append(res, "INTERNET")
	res = append(res, "SITA")
	res = append(res, "ACARS")
	res = append(res, "ADNS")

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

// Utility function for Enum_CodeTransponderBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetransponderbasetype_member_1 Enum_CodeTransponderBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetransponderbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTransponderBaseType_Member_1_MODE_1:
		res = "MODE_1"
	case Enum_CodeTransponderBaseType_Member_1_MODE_2:
		res = "MODE_2"
	case Enum_CodeTransponderBaseType_Member_1_MODE_3A:
		res = "MODE_3A"
	case Enum_CodeTransponderBaseType_Member_1_MODE_4:
		res = "MODE_4"
	case Enum_CodeTransponderBaseType_Member_1_MODE_5:
		res = "MODE_5"
	case Enum_CodeTransponderBaseType_Member_1_MODE_C:
		res = "MODE_C"
	case Enum_CodeTransponderBaseType_Member_1_MODE_S:
		res = "MODE_S"
	}
	return
}

func (enum_codetransponderbasetype_member_1 *Enum_CodeTransponderBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MODE_1":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_1
		return
	case "MODE_2":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_2
		return
	case "MODE_3A":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_3A
		return
	case "MODE_4":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_4
		return
	case "MODE_5":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_5
		return
	case "MODE_C":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_C
		return
	case "MODE_S":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_S
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetransponderbasetype_member_1 *Enum_CodeTransponderBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTransponderBaseType_Member_1_MODE_1":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_1
	case "Enum_CodeTransponderBaseType_Member_1_MODE_2":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_2
	case "Enum_CodeTransponderBaseType_Member_1_MODE_3A":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_3A
	case "Enum_CodeTransponderBaseType_Member_1_MODE_4":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_4
	case "Enum_CodeTransponderBaseType_Member_1_MODE_5":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_5
	case "Enum_CodeTransponderBaseType_Member_1_MODE_C":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_C
	case "Enum_CodeTransponderBaseType_Member_1_MODE_S":
		*enum_codetransponderbasetype_member_1 = Enum_CodeTransponderBaseType_Member_1_MODE_S
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetransponderbasetype_member_1 *Enum_CodeTransponderBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetransponderbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTransponderBaseType_Member_1_MODE_1:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_1"
	case Enum_CodeTransponderBaseType_Member_1_MODE_2:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_2"
	case Enum_CodeTransponderBaseType_Member_1_MODE_3A:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_3A"
	case Enum_CodeTransponderBaseType_Member_1_MODE_4:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_4"
	case Enum_CodeTransponderBaseType_Member_1_MODE_5:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_5"
	case Enum_CodeTransponderBaseType_Member_1_MODE_C:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_C"
	case Enum_CodeTransponderBaseType_Member_1_MODE_S:
		res = "Enum_CodeTransponderBaseType_Member_1_MODE_S"
	}
	return
}

func (enum_codetransponderbasetype_member_1 Enum_CodeTransponderBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_1")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_2")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_3A")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_4")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_5")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_C")
	res = append(res, "Enum_CodeTransponderBaseType_Member_1_MODE_S")

	return
}

func (enum_codetransponderbasetype_member_1 Enum_CodeTransponderBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MODE_1")
	res = append(res, "MODE_2")
	res = append(res, "MODE_3A")
	res = append(res, "MODE_4")
	res = append(res, "MODE_5")
	res = append(res, "MODE_C")
	res = append(res, "MODE_S")

	return
}

// Utility function for Enum_CodeUsageLimitationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeusagelimitationbasetype_member_1 Enum_CodeUsageLimitationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeusagelimitationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUsageLimitationBaseType_Member_1_PERMIT:
		res = "PERMIT"
	case Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL:
		res = "CONDITIONAL"
	case Enum_CodeUsageLimitationBaseType_Member_1_FORBID:
		res = "FORBID"
	case Enum_CodeUsageLimitationBaseType_Member_1_RESERV:
		res = "RESERV"
	}
	return
}

func (enum_codeusagelimitationbasetype_member_1 *Enum_CodeUsageLimitationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PERMIT":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_PERMIT
		return
	case "CONDITIONAL":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL
		return
	case "FORBID":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_FORBID
		return
	case "RESERV":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_RESERV
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeusagelimitationbasetype_member_1 *Enum_CodeUsageLimitationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeUsageLimitationBaseType_Member_1_PERMIT":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_PERMIT
	case "Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL
	case "Enum_CodeUsageLimitationBaseType_Member_1_FORBID":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_FORBID
	case "Enum_CodeUsageLimitationBaseType_Member_1_RESERV":
		*enum_codeusagelimitationbasetype_member_1 = Enum_CodeUsageLimitationBaseType_Member_1_RESERV
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeusagelimitationbasetype_member_1 *Enum_CodeUsageLimitationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeusagelimitationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUsageLimitationBaseType_Member_1_PERMIT:
		res = "Enum_CodeUsageLimitationBaseType_Member_1_PERMIT"
	case Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL:
		res = "Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL"
	case Enum_CodeUsageLimitationBaseType_Member_1_FORBID:
		res = "Enum_CodeUsageLimitationBaseType_Member_1_FORBID"
	case Enum_CodeUsageLimitationBaseType_Member_1_RESERV:
		res = "Enum_CodeUsageLimitationBaseType_Member_1_RESERV"
	}
	return
}

func (enum_codeusagelimitationbasetype_member_1 Enum_CodeUsageLimitationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeUsageLimitationBaseType_Member_1_PERMIT")
	res = append(res, "Enum_CodeUsageLimitationBaseType_Member_1_CONDITIONAL")
	res = append(res, "Enum_CodeUsageLimitationBaseType_Member_1_FORBID")
	res = append(res, "Enum_CodeUsageLimitationBaseType_Member_1_RESERV")

	return
}

func (enum_codeusagelimitationbasetype_member_1 Enum_CodeUsageLimitationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PERMIT")
	res = append(res, "CONDITIONAL")
	res = append(res, "FORBID")
	res = append(res, "RESERV")

	return
}

// Utility function for Enum_CodeVASISBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codevasisbasetype_member_1 Enum_CodeVASISBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codevasisbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVASISBaseType_Member_1_PAPI:
		res = "PAPI"
	case Enum_CodeVASISBaseType_Member_1_APAPI:
		res = "APAPI"
	case Enum_CodeVASISBaseType_Member_1_HAPI:
		res = "HAPI"
	case Enum_CodeVASISBaseType_Member_1_VASIS:
		res = "VASIS"
	case Enum_CodeVASISBaseType_Member_1_AVASIS:
		res = "AVASIS"
	case Enum_CodeVASISBaseType_Member_1_TVASIS:
		res = "TVASIS"
	case Enum_CodeVASISBaseType_Member_1_ATVASIS:
		res = "ATVASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_VASIS:
		res = "3B_VASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS:
		res = "3B_AVASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS:
		res = "3B_ATVASIS"
	case Enum_CodeVASISBaseType_Member_1_PVASI:
		res = "PVASI"
	case Enum_CodeVASISBaseType_Member_1_TRCV:
		res = "TRCV"
	case Enum_CodeVASISBaseType_Member_1_PNI:
		res = "PNI"
	case Enum_CodeVASISBaseType_Member_1_ILU:
		res = "ILU"
	case Enum_CodeVASISBaseType_Member_1_OLS:
		res = "OLS"
	case Enum_CodeVASISBaseType_Member_1_LCVASI:
		res = "LCVASI"
	}
	return
}

func (enum_codevasisbasetype_member_1 *Enum_CodeVASISBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PAPI
		return
	case "APAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_APAPI
		return
	case "HAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_HAPI
		return
	case "VASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_VASIS
		return
	case "AVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_AVASIS
		return
	case "TVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_TVASIS
		return
	case "ATVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_ATVASIS
		return
	case "3B_VASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_VASIS
		return
	case "3B_AVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS
		return
	case "3B_ATVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS
		return
	case "PVASI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PVASI
		return
	case "TRCV":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_TRCV
		return
	case "PNI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PNI
		return
	case "ILU":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_ILU
		return
	case "OLS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_OLS
		return
	case "LCVASI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_LCVASI
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codevasisbasetype_member_1 *Enum_CodeVASISBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVASISBaseType_Member_1_PAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PAPI
	case "Enum_CodeVASISBaseType_Member_1_APAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_APAPI
	case "Enum_CodeVASISBaseType_Member_1_HAPI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_HAPI
	case "Enum_CodeVASISBaseType_Member_1_VASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_VASIS
	case "Enum_CodeVASISBaseType_Member_1_AVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_AVASIS
	case "Enum_CodeVASISBaseType_Member_1_TVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_TVASIS
	case "Enum_CodeVASISBaseType_Member_1_ATVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_ATVASIS
	case "Enum_CodeVASISBaseType_Member_1_N_3B_VASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_VASIS
	case "Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS
	case "Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS
	case "Enum_CodeVASISBaseType_Member_1_PVASI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PVASI
	case "Enum_CodeVASISBaseType_Member_1_TRCV":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_TRCV
	case "Enum_CodeVASISBaseType_Member_1_PNI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_PNI
	case "Enum_CodeVASISBaseType_Member_1_ILU":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_ILU
	case "Enum_CodeVASISBaseType_Member_1_OLS":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_OLS
	case "Enum_CodeVASISBaseType_Member_1_LCVASI":
		*enum_codevasisbasetype_member_1 = Enum_CodeVASISBaseType_Member_1_LCVASI
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codevasisbasetype_member_1 *Enum_CodeVASISBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codevasisbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVASISBaseType_Member_1_PAPI:
		res = "Enum_CodeVASISBaseType_Member_1_PAPI"
	case Enum_CodeVASISBaseType_Member_1_APAPI:
		res = "Enum_CodeVASISBaseType_Member_1_APAPI"
	case Enum_CodeVASISBaseType_Member_1_HAPI:
		res = "Enum_CodeVASISBaseType_Member_1_HAPI"
	case Enum_CodeVASISBaseType_Member_1_VASIS:
		res = "Enum_CodeVASISBaseType_Member_1_VASIS"
	case Enum_CodeVASISBaseType_Member_1_AVASIS:
		res = "Enum_CodeVASISBaseType_Member_1_AVASIS"
	case Enum_CodeVASISBaseType_Member_1_TVASIS:
		res = "Enum_CodeVASISBaseType_Member_1_TVASIS"
	case Enum_CodeVASISBaseType_Member_1_ATVASIS:
		res = "Enum_CodeVASISBaseType_Member_1_ATVASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_VASIS:
		res = "Enum_CodeVASISBaseType_Member_1_N_3B_VASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS:
		res = "Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS"
	case Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS:
		res = "Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS"
	case Enum_CodeVASISBaseType_Member_1_PVASI:
		res = "Enum_CodeVASISBaseType_Member_1_PVASI"
	case Enum_CodeVASISBaseType_Member_1_TRCV:
		res = "Enum_CodeVASISBaseType_Member_1_TRCV"
	case Enum_CodeVASISBaseType_Member_1_PNI:
		res = "Enum_CodeVASISBaseType_Member_1_PNI"
	case Enum_CodeVASISBaseType_Member_1_ILU:
		res = "Enum_CodeVASISBaseType_Member_1_ILU"
	case Enum_CodeVASISBaseType_Member_1_OLS:
		res = "Enum_CodeVASISBaseType_Member_1_OLS"
	case Enum_CodeVASISBaseType_Member_1_LCVASI:
		res = "Enum_CodeVASISBaseType_Member_1_LCVASI"
	}
	return
}

func (enum_codevasisbasetype_member_1 Enum_CodeVASISBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVASISBaseType_Member_1_PAPI")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_APAPI")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_HAPI")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_VASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_AVASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_TVASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_ATVASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_N_3B_VASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_N_3B_AVASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_N_3B_ATVASIS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_PVASI")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_TRCV")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_PNI")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_ILU")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_OLS")
	res = append(res, "Enum_CodeVASISBaseType_Member_1_LCVASI")

	return
}

func (enum_codevasisbasetype_member_1 Enum_CodeVASISBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PAPI")
	res = append(res, "APAPI")
	res = append(res, "HAPI")
	res = append(res, "VASIS")
	res = append(res, "AVASIS")
	res = append(res, "TVASIS")
	res = append(res, "ATVASIS")
	res = append(res, "3B_VASIS")
	res = append(res, "3B_AVASIS")
	res = append(res, "3B_ATVASIS")
	res = append(res, "PVASI")
	res = append(res, "TRCV")
	res = append(res, "PNI")
	res = append(res, "ILU")
	res = append(res, "OLS")
	res = append(res, "LCVASI")

	return
}

// Utility function for Enum_CodeValueInterpretationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codevalueinterpretationbasetype_member_1 Enum_CodeValueInterpretationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codevalueinterpretationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeValueInterpretationBaseType_Member_1_ABOVE:
		res = "ABOVE"
	case Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE:
		res = "AT_OR_ABOVE"
	case Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW:
		res = "AT_OR_BELOW"
	case Enum_CodeValueInterpretationBaseType_Member_1_BELOW:
		res = "BELOW"
	}
	return
}

func (enum_codevalueinterpretationbasetype_member_1 *Enum_CodeValueInterpretationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ABOVE":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_ABOVE
		return
	case "AT_OR_ABOVE":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE
		return
	case "AT_OR_BELOW":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW
		return
	case "BELOW":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_BELOW
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codevalueinterpretationbasetype_member_1 *Enum_CodeValueInterpretationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeValueInterpretationBaseType_Member_1_ABOVE":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_ABOVE
	case "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE
	case "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW
	case "Enum_CodeValueInterpretationBaseType_Member_1_BELOW":
		*enum_codevalueinterpretationbasetype_member_1 = Enum_CodeValueInterpretationBaseType_Member_1_BELOW
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codevalueinterpretationbasetype_member_1 *Enum_CodeValueInterpretationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codevalueinterpretationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeValueInterpretationBaseType_Member_1_ABOVE:
		res = "Enum_CodeValueInterpretationBaseType_Member_1_ABOVE"
	case Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE:
		res = "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE"
	case Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW:
		res = "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW"
	case Enum_CodeValueInterpretationBaseType_Member_1_BELOW:
		res = "Enum_CodeValueInterpretationBaseType_Member_1_BELOW"
	}
	return
}

func (enum_codevalueinterpretationbasetype_member_1 Enum_CodeValueInterpretationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeValueInterpretationBaseType_Member_1_ABOVE")
	res = append(res, "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_ABOVE")
	res = append(res, "Enum_CodeValueInterpretationBaseType_Member_1_AT_OR_BELOW")
	res = append(res, "Enum_CodeValueInterpretationBaseType_Member_1_BELOW")

	return
}

func (enum_codevalueinterpretationbasetype_member_1 Enum_CodeValueInterpretationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ABOVE")
	res = append(res, "AT_OR_ABOVE")
	res = append(res, "AT_OR_BELOW")
	res = append(res, "BELOW")

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

// Utility function for Enum_CodeVisualDockingGuidanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codevisualdockingguidancebasetype_member_1 Enum_CodeVisualDockingGuidanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codevisualdockingguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS:
		res = "AGNIS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA:
		res = "PAPA"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE:
		res = "SAFE_GATE"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC:
		res = "SAFE_DOC"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS:
		res = "APIS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS:
		res = "A_VDGS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP:
		res = "AGNIS_STOP"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA:
		res = "AGNIS_PAPA"
	}
	return
}

func (enum_codevisualdockingguidancebasetype_member_1 *Enum_CodeVisualDockingGuidanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AGNIS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS
		return
	case "PAPA":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA
		return
	case "SAFE_GATE":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE
		return
	case "SAFE_DOC":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC
		return
	case "APIS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS
		return
	case "A_VDGS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS
		return
	case "AGNIS_STOP":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP
		return
	case "AGNIS_PAPA":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codevisualdockingguidancebasetype_member_1 *Enum_CodeVisualDockingGuidanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP
	case "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA":
		*enum_codevisualdockingguidancebasetype_member_1 = Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codevisualdockingguidancebasetype_member_1 *Enum_CodeVisualDockingGuidanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codevisualdockingguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP"
	case Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA:
		res = "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA"
	}
	return
}

func (enum_codevisualdockingguidancebasetype_member_1 Enum_CodeVisualDockingGuidanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_PAPA")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_GATE")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_SAFE_DOC")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_APIS")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_A_VDGS")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_STOP")
	res = append(res, "Enum_CodeVisualDockingGuidanceBaseType_Member_1_AGNIS_PAPA")

	return
}

func (enum_codevisualdockingguidancebasetype_member_1 Enum_CodeVisualDockingGuidanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AGNIS")
	res = append(res, "PAPA")
	res = append(res, "SAFE_GATE")
	res = append(res, "SAFE_DOC")
	res = append(res, "APIS")
	res = append(res, "A_VDGS")
	res = append(res, "AGNIS_STOP")
	res = append(res, "AGNIS_PAPA")

	return
}

// Utility function for Enum_CodeWakeTurbulenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codewaketurbulencebasetype_member_1 Enum_CodeWakeTurbulenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codewaketurbulencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeWakeTurbulenceBaseType_Member_1_LOW:
		res = "LOW"
	case Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM:
		res = "MEDIUM"
	case Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH:
		res = "HIGH"
	}
	return
}

func (enum_codewaketurbulencebasetype_member_1 *Enum_CodeWakeTurbulenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LOW":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_LOW
		return
	case "MEDIUM":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM
		return
	case "HIGH":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codewaketurbulencebasetype_member_1 *Enum_CodeWakeTurbulenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeWakeTurbulenceBaseType_Member_1_LOW":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_LOW
	case "Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM
	case "Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH":
		*enum_codewaketurbulencebasetype_member_1 = Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codewaketurbulencebasetype_member_1 *Enum_CodeWakeTurbulenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codewaketurbulencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeWakeTurbulenceBaseType_Member_1_LOW:
		res = "Enum_CodeWakeTurbulenceBaseType_Member_1_LOW"
	case Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM:
		res = "Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM"
	case Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH:
		res = "Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH"
	}
	return
}

func (enum_codewaketurbulencebasetype_member_1 Enum_CodeWakeTurbulenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeWakeTurbulenceBaseType_Member_1_LOW")
	res = append(res, "Enum_CodeWakeTurbulenceBaseType_Member_1_MEDIUM")
	res = append(res, "Enum_CodeWakeTurbulenceBaseType_Member_1_HIGH")

	return
}

func (enum_codewaketurbulencebasetype_member_1 Enum_CodeWakeTurbulenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LOW")
	res = append(res, "MEDIUM")
	res = append(res, "HIGH")

	return
}

// Utility function for Enum_CodeWorkAreaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeworkareabasetype_member_1 Enum_CodeWorkAreaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeworkareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION:
		res = "CONSTRUCTION"
	case Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK:
		res = "SURFACEWORK"
	case Enum_CodeWorkAreaBaseType_Member_1_PARKED:
		res = "PARKED"
	}
	return
}

func (enum_codeworkareabasetype_member_1 *Enum_CodeWorkAreaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CONSTRUCTION":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION
		return
	case "SURFACEWORK":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK
		return
	case "PARKED":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_PARKED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeworkareabasetype_member_1 *Enum_CodeWorkAreaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION
	case "Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK
	case "Enum_CodeWorkAreaBaseType_Member_1_PARKED":
		*enum_codeworkareabasetype_member_1 = Enum_CodeWorkAreaBaseType_Member_1_PARKED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeworkareabasetype_member_1 *Enum_CodeWorkAreaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeworkareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION:
		res = "Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION"
	case Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK:
		res = "Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK"
	case Enum_CodeWorkAreaBaseType_Member_1_PARKED:
		res = "Enum_CodeWorkAreaBaseType_Member_1_PARKED"
	}
	return
}

func (enum_codeworkareabasetype_member_1 Enum_CodeWorkAreaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeWorkAreaBaseType_Member_1_CONSTRUCTION")
	res = append(res, "Enum_CodeWorkAreaBaseType_Member_1_SURFACEWORK")
	res = append(res, "Enum_CodeWorkAreaBaseType_Member_1_PARKED")

	return
}

func (enum_codeworkareabasetype_member_1 Enum_CodeWorkAreaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CONSTRUCTION")
	res = append(res, "SURFACEWORK")
	res = append(res, "PARKED")

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

// Utility function for Enum_UomDepthType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomdepthtype_member_1 Enum_UomDepthType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomdepthtype_member_1 {
	// insertion code per enum code
	case Enum_UomDepthType_Member_1_MM:
		res = "MM"
	case Enum_UomDepthType_Member_1_CM:
		res = "CM"
	case Enum_UomDepthType_Member_1_IN:
		res = "IN"
	case Enum_UomDepthType_Member_1_FT:
		res = "FT"
	}
	return
}

func (enum_uomdepthtype_member_1 *Enum_UomDepthType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MM":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_MM
		return
	case "CM":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_CM
		return
	case "IN":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_IN
		return
	case "FT":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_FT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomdepthtype_member_1 *Enum_UomDepthType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomDepthType_Member_1_MM":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_MM
	case "Enum_UomDepthType_Member_1_CM":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_CM
	case "Enum_UomDepthType_Member_1_IN":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_IN
	case "Enum_UomDepthType_Member_1_FT":
		*enum_uomdepthtype_member_1 = Enum_UomDepthType_Member_1_FT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomdepthtype_member_1 *Enum_UomDepthType_Member_1) ToCodeString() (res string) {

	switch *enum_uomdepthtype_member_1 {
	// insertion code per enum code
	case Enum_UomDepthType_Member_1_MM:
		res = "Enum_UomDepthType_Member_1_MM"
	case Enum_UomDepthType_Member_1_CM:
		res = "Enum_UomDepthType_Member_1_CM"
	case Enum_UomDepthType_Member_1_IN:
		res = "Enum_UomDepthType_Member_1_IN"
	case Enum_UomDepthType_Member_1_FT:
		res = "Enum_UomDepthType_Member_1_FT"
	}
	return
}

func (enum_uomdepthtype_member_1 Enum_UomDepthType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomDepthType_Member_1_MM")
	res = append(res, "Enum_UomDepthType_Member_1_CM")
	res = append(res, "Enum_UomDepthType_Member_1_IN")
	res = append(res, "Enum_UomDepthType_Member_1_FT")

	return
}

func (enum_uomdepthtype_member_1 Enum_UomDepthType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MM")
	res = append(res, "CM")
	res = append(res, "IN")
	res = append(res, "FT")

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

// Utility function for Enum_UomFLType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomfltype_member_1 Enum_UomFLType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomfltype_member_1 {
	// insertion code per enum code
	case Enum_UomFLType_Member_1_FL:
		res = "FL"
	case Enum_UomFLType_Member_1_SM:
		res = "SM"
	}
	return
}

func (enum_uomfltype_member_1 *Enum_UomFLType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FL":
		*enum_uomfltype_member_1 = Enum_UomFLType_Member_1_FL
		return
	case "SM":
		*enum_uomfltype_member_1 = Enum_UomFLType_Member_1_SM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomfltype_member_1 *Enum_UomFLType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomFLType_Member_1_FL":
		*enum_uomfltype_member_1 = Enum_UomFLType_Member_1_FL
	case "Enum_UomFLType_Member_1_SM":
		*enum_uomfltype_member_1 = Enum_UomFLType_Member_1_SM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomfltype_member_1 *Enum_UomFLType_Member_1) ToCodeString() (res string) {

	switch *enum_uomfltype_member_1 {
	// insertion code per enum code
	case Enum_UomFLType_Member_1_FL:
		res = "Enum_UomFLType_Member_1_FL"
	case Enum_UomFLType_Member_1_SM:
		res = "Enum_UomFLType_Member_1_SM"
	}
	return
}

func (enum_uomfltype_member_1 Enum_UomFLType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomFLType_Member_1_FL")
	res = append(res, "Enum_UomFLType_Member_1_SM")

	return
}

func (enum_uomfltype_member_1 Enum_UomFLType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FL")
	res = append(res, "SM")

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

// Utility function for Enum_UomPressureType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uompressuretype_member_1 Enum_UomPressureType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uompressuretype_member_1 {
	// insertion code per enum code
	case Enum_UomPressureType_Member_1_PA:
		res = "PA"
	case Enum_UomPressureType_Member_1_MPA:
		res = "MPA"
	case Enum_UomPressureType_Member_1_PSI:
		res = "PSI"
	case Enum_UomPressureType_Member_1_BAR:
		res = "BAR"
	case Enum_UomPressureType_Member_1_TORR:
		res = "TORR"
	case Enum_UomPressureType_Member_1_ATM:
		res = "ATM"
	case Enum_UomPressureType_Member_1_HPA:
		res = "HPA"
	}
	return
}

func (enum_uompressuretype_member_1 *Enum_UomPressureType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_PA
		return
	case "MPA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_MPA
		return
	case "PSI":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_PSI
		return
	case "BAR":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_BAR
		return
	case "TORR":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_TORR
		return
	case "ATM":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_ATM
		return
	case "HPA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_HPA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uompressuretype_member_1 *Enum_UomPressureType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomPressureType_Member_1_PA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_PA
	case "Enum_UomPressureType_Member_1_MPA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_MPA
	case "Enum_UomPressureType_Member_1_PSI":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_PSI
	case "Enum_UomPressureType_Member_1_BAR":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_BAR
	case "Enum_UomPressureType_Member_1_TORR":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_TORR
	case "Enum_UomPressureType_Member_1_ATM":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_ATM
	case "Enum_UomPressureType_Member_1_HPA":
		*enum_uompressuretype_member_1 = Enum_UomPressureType_Member_1_HPA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uompressuretype_member_1 *Enum_UomPressureType_Member_1) ToCodeString() (res string) {

	switch *enum_uompressuretype_member_1 {
	// insertion code per enum code
	case Enum_UomPressureType_Member_1_PA:
		res = "Enum_UomPressureType_Member_1_PA"
	case Enum_UomPressureType_Member_1_MPA:
		res = "Enum_UomPressureType_Member_1_MPA"
	case Enum_UomPressureType_Member_1_PSI:
		res = "Enum_UomPressureType_Member_1_PSI"
	case Enum_UomPressureType_Member_1_BAR:
		res = "Enum_UomPressureType_Member_1_BAR"
	case Enum_UomPressureType_Member_1_TORR:
		res = "Enum_UomPressureType_Member_1_TORR"
	case Enum_UomPressureType_Member_1_ATM:
		res = "Enum_UomPressureType_Member_1_ATM"
	case Enum_UomPressureType_Member_1_HPA:
		res = "Enum_UomPressureType_Member_1_HPA"
	}
	return
}

func (enum_uompressuretype_member_1 Enum_UomPressureType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomPressureType_Member_1_PA")
	res = append(res, "Enum_UomPressureType_Member_1_MPA")
	res = append(res, "Enum_UomPressureType_Member_1_PSI")
	res = append(res, "Enum_UomPressureType_Member_1_BAR")
	res = append(res, "Enum_UomPressureType_Member_1_TORR")
	res = append(res, "Enum_UomPressureType_Member_1_ATM")
	res = append(res, "Enum_UomPressureType_Member_1_HPA")

	return
}

func (enum_uompressuretype_member_1 Enum_UomPressureType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PA")
	res = append(res, "MPA")
	res = append(res, "PSI")
	res = append(res, "BAR")
	res = append(res, "TORR")
	res = append(res, "ATM")
	res = append(res, "HPA")

	return
}

// Utility function for Enum_UomSpeedType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomspeedtype_member_1 Enum_UomSpeedType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomspeedtype_member_1 {
	// insertion code per enum code
	case Enum_UomSpeedType_Member_1_KM_H:
		res = "KM_H"
	case Enum_UomSpeedType_Member_1_KT:
		res = "KT"
	case Enum_UomSpeedType_Member_1_MACH:
		res = "MACH"
	case Enum_UomSpeedType_Member_1_M_MIN:
		res = "M_MIN"
	case Enum_UomSpeedType_Member_1_FT_MIN:
		res = "FT_MIN"
	case Enum_UomSpeedType_Member_1_M_SEC:
		res = "M_SEC"
	case Enum_UomSpeedType_Member_1_FT_SEC:
		res = "FT_SEC"
	case Enum_UomSpeedType_Member_1_MPH:
		res = "MPH"
	}
	return
}

func (enum_uomspeedtype_member_1 *Enum_UomSpeedType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "KM_H":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_KM_H
		return
	case "KT":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_KT
		return
	case "MACH":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_MACH
		return
	case "M_MIN":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_M_MIN
		return
	case "FT_MIN":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_FT_MIN
		return
	case "M_SEC":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_M_SEC
		return
	case "FT_SEC":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_FT_SEC
		return
	case "MPH":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_MPH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomspeedtype_member_1 *Enum_UomSpeedType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomSpeedType_Member_1_KM_H":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_KM_H
	case "Enum_UomSpeedType_Member_1_KT":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_KT
	case "Enum_UomSpeedType_Member_1_MACH":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_MACH
	case "Enum_UomSpeedType_Member_1_M_MIN":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_M_MIN
	case "Enum_UomSpeedType_Member_1_FT_MIN":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_FT_MIN
	case "Enum_UomSpeedType_Member_1_M_SEC":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_M_SEC
	case "Enum_UomSpeedType_Member_1_FT_SEC":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_FT_SEC
	case "Enum_UomSpeedType_Member_1_MPH":
		*enum_uomspeedtype_member_1 = Enum_UomSpeedType_Member_1_MPH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomspeedtype_member_1 *Enum_UomSpeedType_Member_1) ToCodeString() (res string) {

	switch *enum_uomspeedtype_member_1 {
	// insertion code per enum code
	case Enum_UomSpeedType_Member_1_KM_H:
		res = "Enum_UomSpeedType_Member_1_KM_H"
	case Enum_UomSpeedType_Member_1_KT:
		res = "Enum_UomSpeedType_Member_1_KT"
	case Enum_UomSpeedType_Member_1_MACH:
		res = "Enum_UomSpeedType_Member_1_MACH"
	case Enum_UomSpeedType_Member_1_M_MIN:
		res = "Enum_UomSpeedType_Member_1_M_MIN"
	case Enum_UomSpeedType_Member_1_FT_MIN:
		res = "Enum_UomSpeedType_Member_1_FT_MIN"
	case Enum_UomSpeedType_Member_1_M_SEC:
		res = "Enum_UomSpeedType_Member_1_M_SEC"
	case Enum_UomSpeedType_Member_1_FT_SEC:
		res = "Enum_UomSpeedType_Member_1_FT_SEC"
	case Enum_UomSpeedType_Member_1_MPH:
		res = "Enum_UomSpeedType_Member_1_MPH"
	}
	return
}

func (enum_uomspeedtype_member_1 Enum_UomSpeedType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomSpeedType_Member_1_KM_H")
	res = append(res, "Enum_UomSpeedType_Member_1_KT")
	res = append(res, "Enum_UomSpeedType_Member_1_MACH")
	res = append(res, "Enum_UomSpeedType_Member_1_M_MIN")
	res = append(res, "Enum_UomSpeedType_Member_1_FT_MIN")
	res = append(res, "Enum_UomSpeedType_Member_1_M_SEC")
	res = append(res, "Enum_UomSpeedType_Member_1_FT_SEC")
	res = append(res, "Enum_UomSpeedType_Member_1_MPH")

	return
}

func (enum_uomspeedtype_member_1 Enum_UomSpeedType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "KM_H")
	res = append(res, "KT")
	res = append(res, "MACH")
	res = append(res, "M_MIN")
	res = append(res, "FT_MIN")
	res = append(res, "M_SEC")
	res = append(res, "FT_SEC")
	res = append(res, "MPH")

	return
}

// Utility function for Enum_UomTemperatureType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomtemperaturetype_member_1 Enum_UomTemperatureType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomtemperaturetype_member_1 {
	// insertion code per enum code
	case Enum_UomTemperatureType_Member_1_C:
		res = "C"
	case Enum_UomTemperatureType_Member_1_F:
		res = "F"
	case Enum_UomTemperatureType_Member_1_K:
		res = "K"
	}
	return
}

func (enum_uomtemperaturetype_member_1 *Enum_UomTemperatureType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "C":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_C
		return
	case "F":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_F
		return
	case "K":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_K
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomtemperaturetype_member_1 *Enum_UomTemperatureType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomTemperatureType_Member_1_C":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_C
	case "Enum_UomTemperatureType_Member_1_F":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_F
	case "Enum_UomTemperatureType_Member_1_K":
		*enum_uomtemperaturetype_member_1 = Enum_UomTemperatureType_Member_1_K
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomtemperaturetype_member_1 *Enum_UomTemperatureType_Member_1) ToCodeString() (res string) {

	switch *enum_uomtemperaturetype_member_1 {
	// insertion code per enum code
	case Enum_UomTemperatureType_Member_1_C:
		res = "Enum_UomTemperatureType_Member_1_C"
	case Enum_UomTemperatureType_Member_1_F:
		res = "Enum_UomTemperatureType_Member_1_F"
	case Enum_UomTemperatureType_Member_1_K:
		res = "Enum_UomTemperatureType_Member_1_K"
	}
	return
}

func (enum_uomtemperaturetype_member_1 Enum_UomTemperatureType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomTemperatureType_Member_1_C")
	res = append(res, "Enum_UomTemperatureType_Member_1_F")
	res = append(res, "Enum_UomTemperatureType_Member_1_K")

	return
}

func (enum_uomtemperaturetype_member_1 Enum_UomTemperatureType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "C")
	res = append(res, "F")
	res = append(res, "K")

	return
}

// Utility function for Enum_UomWeightType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomweighttype_member_1 Enum_UomWeightType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomweighttype_member_1 {
	// insertion code per enum code
	case Enum_UomWeightType_Member_1_KG:
		res = "KG"
	case Enum_UomWeightType_Member_1_T:
		res = "T"
	case Enum_UomWeightType_Member_1_LB:
		res = "LB"
	case Enum_UomWeightType_Member_1_TON:
		res = "TON"
	}
	return
}

func (enum_uomweighttype_member_1 *Enum_UomWeightType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "KG":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_KG
		return
	case "T":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_T
		return
	case "LB":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_LB
		return
	case "TON":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_TON
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomweighttype_member_1 *Enum_UomWeightType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomWeightType_Member_1_KG":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_KG
	case "Enum_UomWeightType_Member_1_T":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_T
	case "Enum_UomWeightType_Member_1_LB":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_LB
	case "Enum_UomWeightType_Member_1_TON":
		*enum_uomweighttype_member_1 = Enum_UomWeightType_Member_1_TON
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomweighttype_member_1 *Enum_UomWeightType_Member_1) ToCodeString() (res string) {

	switch *enum_uomweighttype_member_1 {
	// insertion code per enum code
	case Enum_UomWeightType_Member_1_KG:
		res = "Enum_UomWeightType_Member_1_KG"
	case Enum_UomWeightType_Member_1_T:
		res = "Enum_UomWeightType_Member_1_T"
	case Enum_UomWeightType_Member_1_LB:
		res = "Enum_UomWeightType_Member_1_LB"
	case Enum_UomWeightType_Member_1_TON:
		res = "Enum_UomWeightType_Member_1_TON"
	}
	return
}

func (enum_uomweighttype_member_1 Enum_UomWeightType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomWeightType_Member_1_KG")
	res = append(res, "Enum_UomWeightType_Member_1_T")
	res = append(res, "Enum_UomWeightType_Member_1_LB")
	res = append(res, "Enum_UomWeightType_Member_1_TON")

	return
}

func (enum_uomweighttype_member_1 Enum_UomWeightType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "KG")
	res = append(res, "T")
	res = append(res, "LB")
	res = append(res, "TON")

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
