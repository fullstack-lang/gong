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

// Utility function for Enum_CodeAirspaceActivityBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspaceactivitybasetype_member_1 Enum_CodeAirspaceActivityBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspaceactivitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC:
		res = "AD_TFC"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC:
		res = "HELI_TFC"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING:
		res = "TRAINING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS:
		res = "AEROBATICS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW:
		res = "AIRSHOW"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SPORT:
		res = "SPORT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ULM:
		res = "ULM"
	case Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING:
		res = "GLIDING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER:
		res = "PARAGLIDER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING:
		res = "HANGGLIDING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE:
		res = "PARACHUTE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP:
		res = "AIR_DROP"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON:
		res = "BALLOON"
	case Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE:
		res = "RADIOSONDE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT:
		res = "SPACE_FLIGHT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_UAV:
		res = "UAV"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK:
		res = "AERIAL_WORK"
	case Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING:
		res = "CROP_DUSTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING:
		res = "FIRE_FIGHTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS:
		res = "MILOPS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL:
		res = "REFUEL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING:
		res = "JET_CLIMBING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE:
		res = "EXERCISE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TOWING:
		res = "TOWING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER:
		res = "NAVAL_EXER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES:
		res = "MISSILES"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN:
		res = "AIR_GUN"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY:
		res = "ARTILLERY"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING:
		res = "SHOOTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING:
		res = "BLASTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING:
		res = "WATER_BLASTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL:
		res = "ANTI_HAIL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BIRD:
		res = "BIRD"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION:
		res = "BIRD_MIGRATION"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK:
		res = "FIREWORK"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO:
		res = "HI_RADIO"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT:
		res = "HI_LIGHT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_LASER:
		res = "LASER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NATURE:
		res = "NATURE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA:
		res = "FAUNA"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE:
		res = "NO_NOISE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT:
		res = "ACCIDENT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION:
		res = "POPULATION"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP:
		res = "VIP"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES:
		res = "VIP_PRES"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE:
		res = "VIP_VICE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_OIL:
		res = "OIL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_GAS:
		res = "GAS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY:
		res = "REFINERY"
	case Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL:
		res = "CHEMICAL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR:
		res = "NUCLEAR"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL:
		res = "TECHNICAL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ATS:
		res = "ATS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE:
		res = "PROCEDURE"
	}
	return
}

func (enum_codeairspaceactivitybasetype_member_1 *Enum_CodeAirspaceActivityBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AD_TFC":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC
		return
	case "HELI_TFC":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC
		return
	case "TRAINING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING
		return
	case "AEROBATICS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS
		return
	case "AIRSHOW":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW
		return
	case "SPORT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SPORT
		return
	case "ULM":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ULM
		return
	case "GLIDING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING
		return
	case "PARAGLIDER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER
		return
	case "HANGGLIDING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING
		return
	case "PARACHUTE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE
		return
	case "AIR_DROP":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP
		return
	case "BALLOON":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON
		return
	case "RADIOSONDE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE
		return
	case "SPACE_FLIGHT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT
		return
	case "UAV":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_UAV
		return
	case "AERIAL_WORK":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK
		return
	case "CROP_DUSTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING
		return
	case "FIRE_FIGHTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING
		return
	case "MILOPS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS
		return
	case "REFUEL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL
		return
	case "JET_CLIMBING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING
		return
	case "EXERCISE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE
		return
	case "TOWING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TOWING
		return
	case "NAVAL_EXER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER
		return
	case "MISSILES":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES
		return
	case "AIR_GUN":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN
		return
	case "ARTILLERY":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY
		return
	case "SHOOTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING
		return
	case "BLASTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING
		return
	case "WATER_BLASTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING
		return
	case "ANTI_HAIL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL
		return
	case "BIRD":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BIRD
		return
	case "BIRD_MIGRATION":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION
		return
	case "FIREWORK":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK
		return
	case "HI_RADIO":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO
		return
	case "HI_LIGHT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT
		return
	case "LASER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_LASER
		return
	case "NATURE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NATURE
		return
	case "FAUNA":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA
		return
	case "NO_NOISE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE
		return
	case "ACCIDENT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT
		return
	case "POPULATION":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION
		return
	case "VIP":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP
		return
	case "VIP_PRES":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES
		return
	case "VIP_VICE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE
		return
	case "OIL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_OIL
		return
	case "GAS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_GAS
		return
	case "REFINERY":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY
		return
	case "CHEMICAL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL
		return
	case "NUCLEAR":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR
		return
	case "TECHNICAL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL
		return
	case "ATS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ATS
		return
	case "PROCEDURE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspaceactivitybasetype_member_1 *Enum_CodeAirspaceActivityBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC
	case "Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC
	case "Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW
	case "Enum_CodeAirspaceActivityBaseType_Member_1_SPORT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SPORT
	case "Enum_CodeAirspaceActivityBaseType_Member_1_ULM":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ULM
	case "Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER
	case "Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP
	case "Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON
	case "Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT
	case "Enum_CodeAirspaceActivityBaseType_Member_1_UAV":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_UAV
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK
	case "Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS
	case "Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL
	case "Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_TOWING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TOWING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER
	case "Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES
	case "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN
	case "Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY
	case "Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING
	case "Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL
	case "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BIRD
	case "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION
	case "Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK
	case "Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO
	case "Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT
	case "Enum_CodeAirspaceActivityBaseType_Member_1_LASER":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_LASER
	case "Enum_CodeAirspaceActivityBaseType_Member_1_NATURE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NATURE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA
	case "Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT
	case "Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION
	case "Enum_CodeAirspaceActivityBaseType_Member_1_VIP":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP
	case "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES
	case "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE
	case "Enum_CodeAirspaceActivityBaseType_Member_1_OIL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_OIL
	case "Enum_CodeAirspaceActivityBaseType_Member_1_GAS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_GAS
	case "Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY
	case "Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL
	case "Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR
	case "Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL
	case "Enum_CodeAirspaceActivityBaseType_Member_1_ATS":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_ATS
	case "Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE":
		*enum_codeairspaceactivitybasetype_member_1 = Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspaceactivitybasetype_member_1 *Enum_CodeAirspaceActivityBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspaceactivitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SPORT:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_SPORT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ULM:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_ULM"
	case Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON"
	case Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_UAV:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_UAV"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK"
	case Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TOWING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_TOWING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES"
	case Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY"
	case Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BIRD:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD"
	case Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO"
	case Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_LASER:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_LASER"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NATURE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_NATURE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT"
	case Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_VIP"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES"
	case Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE"
	case Enum_CodeAirspaceActivityBaseType_Member_1_OIL:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_OIL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_GAS:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_GAS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY"
	case Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR"
	case Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL"
	case Enum_CodeAirspaceActivityBaseType_Member_1_ATS:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_ATS"
	case Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE:
		res = "Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE"
	}
	return
}

func (enum_codeairspaceactivitybasetype_member_1 Enum_CodeAirspaceActivityBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AD_TFC")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_HELI_TFC")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_TRAINING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AEROBATICS")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AIRSHOW")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_SPORT")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_ULM")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_GLIDING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_PARAGLIDER")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_HANGGLIDING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_PARACHUTE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_DROP")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_BALLOON")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_RADIOSONDE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_SPACE_FLIGHT")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_UAV")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AERIAL_WORK")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_CROP_DUSTING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_FIRE_FIGHTING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_MILOPS")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_REFUEL")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_JET_CLIMBING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_EXERCISE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_TOWING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_NAVAL_EXER")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_MISSILES")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_AIR_GUN")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_ARTILLERY")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_SHOOTING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_BLASTING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_WATER_BLASTING")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_ANTI_HAIL")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_BIRD_MIGRATION")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_FIREWORK")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_HI_RADIO")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_HI_LIGHT")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_LASER")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_NATURE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_FAUNA")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_NO_NOISE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_ACCIDENT")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_POPULATION")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_VIP")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_PRES")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_VIP_VICE")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_OIL")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_GAS")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_REFINERY")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_CHEMICAL")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_NUCLEAR")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_TECHNICAL")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_ATS")
	res = append(res, "Enum_CodeAirspaceActivityBaseType_Member_1_PROCEDURE")

	return
}

func (enum_codeairspaceactivitybasetype_member_1 Enum_CodeAirspaceActivityBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AD_TFC")
	res = append(res, "HELI_TFC")
	res = append(res, "TRAINING")
	res = append(res, "AEROBATICS")
	res = append(res, "AIRSHOW")
	res = append(res, "SPORT")
	res = append(res, "ULM")
	res = append(res, "GLIDING")
	res = append(res, "PARAGLIDER")
	res = append(res, "HANGGLIDING")
	res = append(res, "PARACHUTE")
	res = append(res, "AIR_DROP")
	res = append(res, "BALLOON")
	res = append(res, "RADIOSONDE")
	res = append(res, "SPACE_FLIGHT")
	res = append(res, "UAV")
	res = append(res, "AERIAL_WORK")
	res = append(res, "CROP_DUSTING")
	res = append(res, "FIRE_FIGHTING")
	res = append(res, "MILOPS")
	res = append(res, "REFUEL")
	res = append(res, "JET_CLIMBING")
	res = append(res, "EXERCISE")
	res = append(res, "TOWING")
	res = append(res, "NAVAL_EXER")
	res = append(res, "MISSILES")
	res = append(res, "AIR_GUN")
	res = append(res, "ARTILLERY")
	res = append(res, "SHOOTING")
	res = append(res, "BLASTING")
	res = append(res, "WATER_BLASTING")
	res = append(res, "ANTI_HAIL")
	res = append(res, "BIRD")
	res = append(res, "BIRD_MIGRATION")
	res = append(res, "FIREWORK")
	res = append(res, "HI_RADIO")
	res = append(res, "HI_LIGHT")
	res = append(res, "LASER")
	res = append(res, "NATURE")
	res = append(res, "FAUNA")
	res = append(res, "NO_NOISE")
	res = append(res, "ACCIDENT")
	res = append(res, "POPULATION")
	res = append(res, "VIP")
	res = append(res, "VIP_PRES")
	res = append(res, "VIP_VICE")
	res = append(res, "OIL")
	res = append(res, "GAS")
	res = append(res, "REFINERY")
	res = append(res, "CHEMICAL")
	res = append(res, "NUCLEAR")
	res = append(res, "TECHNICAL")
	res = append(res, "ATS")
	res = append(res, "PROCEDURE")

	return
}

// Utility function for Enum_CodeAirspaceAggregationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspaceaggregationbasetype_member_1 Enum_CodeAirspaceAggregationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspaceaggregationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceAggregationBaseType_Member_1_BASE:
		res = "BASE"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_UNION:
		res = "UNION"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS:
		res = "INTERS"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR:
		res = "SUBTR"
	}
	return
}

func (enum_codeairspaceaggregationbasetype_member_1 *Enum_CodeAirspaceAggregationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BASE":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_BASE
		return
	case "UNION":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_UNION
		return
	case "INTERS":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS
		return
	case "SUBTR":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspaceaggregationbasetype_member_1 *Enum_CodeAirspaceAggregationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspaceAggregationBaseType_Member_1_BASE":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_BASE
	case "Enum_CodeAirspaceAggregationBaseType_Member_1_UNION":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_UNION
	case "Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS
	case "Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR":
		*enum_codeairspaceaggregationbasetype_member_1 = Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspaceaggregationbasetype_member_1 *Enum_CodeAirspaceAggregationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspaceaggregationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceAggregationBaseType_Member_1_BASE:
		res = "Enum_CodeAirspaceAggregationBaseType_Member_1_BASE"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_UNION:
		res = "Enum_CodeAirspaceAggregationBaseType_Member_1_UNION"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS:
		res = "Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS"
	case Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR:
		res = "Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR"
	}
	return
}

func (enum_codeairspaceaggregationbasetype_member_1 Enum_CodeAirspaceAggregationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspaceAggregationBaseType_Member_1_BASE")
	res = append(res, "Enum_CodeAirspaceAggregationBaseType_Member_1_UNION")
	res = append(res, "Enum_CodeAirspaceAggregationBaseType_Member_1_INTERS")
	res = append(res, "Enum_CodeAirspaceAggregationBaseType_Member_1_SUBTR")

	return
}

func (enum_codeairspaceaggregationbasetype_member_1 Enum_CodeAirspaceAggregationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BASE")
	res = append(res, "UNION")
	res = append(res, "INTERS")
	res = append(res, "SUBTR")

	return
}

// Utility function for Enum_CodeAirspaceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspacebasetype_member_1 Enum_CodeAirspaceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceBaseType_Member_1_NAS:
		res = "NAS"
	case Enum_CodeAirspaceBaseType_Member_1_FIR:
		res = "FIR"
	case Enum_CodeAirspaceBaseType_Member_1_FIR_P:
		res = "FIR_P"
	case Enum_CodeAirspaceBaseType_Member_1_UIR:
		res = "UIR"
	case Enum_CodeAirspaceBaseType_Member_1_UIR_P:
		res = "UIR_P"
	case Enum_CodeAirspaceBaseType_Member_1_CTA:
		res = "CTA"
	case Enum_CodeAirspaceBaseType_Member_1_CTA_P:
		res = "CTA_P"
	case Enum_CodeAirspaceBaseType_Member_1_OCA_P:
		res = "OCA_P"
	case Enum_CodeAirspaceBaseType_Member_1_OCA:
		res = "OCA"
	case Enum_CodeAirspaceBaseType_Member_1_UTA:
		res = "UTA"
	case Enum_CodeAirspaceBaseType_Member_1_UTA_P:
		res = "UTA_P"
	case Enum_CodeAirspaceBaseType_Member_1_TMA:
		res = "TMA"
	case Enum_CodeAirspaceBaseType_Member_1_TMA_P:
		res = "TMA_P"
	case Enum_CodeAirspaceBaseType_Member_1_CTR:
		res = "CTR"
	case Enum_CodeAirspaceBaseType_Member_1_CTR_P:
		res = "CTR_P"
	case Enum_CodeAirspaceBaseType_Member_1_OTA:
		res = "OTA"
	case Enum_CodeAirspaceBaseType_Member_1_SECTOR:
		res = "SECTOR"
	case Enum_CodeAirspaceBaseType_Member_1_SECTOR_C:
		res = "SECTOR_C"
	case Enum_CodeAirspaceBaseType_Member_1_TSA:
		res = "TSA"
	case Enum_CodeAirspaceBaseType_Member_1_CBA:
		res = "CBA"
	case Enum_CodeAirspaceBaseType_Member_1_RCA:
		res = "RCA"
	case Enum_CodeAirspaceBaseType_Member_1_RAS:
		res = "RAS"
	case Enum_CodeAirspaceBaseType_Member_1_AWY:
		res = "AWY"
	case Enum_CodeAirspaceBaseType_Member_1_MTR:
		res = "MTR"
	case Enum_CodeAirspaceBaseType_Member_1_P:
		res = "P"
	case Enum_CodeAirspaceBaseType_Member_1_R:
		res = "R"
	case Enum_CodeAirspaceBaseType_Member_1_D:
		res = "D"
	case Enum_CodeAirspaceBaseType_Member_1_ADIZ:
		res = "ADIZ"
	case Enum_CodeAirspaceBaseType_Member_1_NO_FIR:
		res = "NO_FIR"
	case Enum_CodeAirspaceBaseType_Member_1_PART:
		res = "PART"
	case Enum_CodeAirspaceBaseType_Member_1_CLASS:
		res = "CLASS"
	case Enum_CodeAirspaceBaseType_Member_1_POLITICAL:
		res = "POLITICAL"
	case Enum_CodeAirspaceBaseType_Member_1_D_OTHER:
		res = "D_OTHER"
	case Enum_CodeAirspaceBaseType_Member_1_TRA:
		res = "TRA"
	case Enum_CodeAirspaceBaseType_Member_1_A:
		res = "A"
	case Enum_CodeAirspaceBaseType_Member_1_W:
		res = "W"
	case Enum_CodeAirspaceBaseType_Member_1_PROTECT:
		res = "PROTECT"
	case Enum_CodeAirspaceBaseType_Member_1_AMA:
		res = "AMA"
	case Enum_CodeAirspaceBaseType_Member_1_ASR:
		res = "ASR"
	case Enum_CodeAirspaceBaseType_Member_1_ADV:
		res = "ADV"
	case Enum_CodeAirspaceBaseType_Member_1_UADV:
		res = "UADV"
	case Enum_CodeAirspaceBaseType_Member_1_ATZ:
		res = "ATZ"
	case Enum_CodeAirspaceBaseType_Member_1_ATZ_P:
		res = "ATZ_P"
	case Enum_CodeAirspaceBaseType_Member_1_HTZ:
		res = "HTZ"
	case Enum_CodeAirspaceBaseType_Member_1_NAS_P:
		res = "NAS_P"
	}
	return
}

func (enum_codeairspacebasetype_member_1 *Enum_CodeAirspaceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NAS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NAS
		return
	case "FIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_FIR
		return
	case "FIR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_FIR_P
		return
	case "UIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UIR
		return
	case "UIR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UIR_P
		return
	case "CTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTA
		return
	case "CTA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTA_P
		return
	case "OCA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OCA_P
		return
	case "OCA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OCA
		return
	case "UTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UTA
		return
	case "UTA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UTA_P
		return
	case "TMA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TMA
		return
	case "TMA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TMA_P
		return
	case "CTR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTR
		return
	case "CTR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTR_P
		return
	case "OTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OTA
		return
	case "SECTOR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_SECTOR
		return
	case "SECTOR_C":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_SECTOR_C
		return
	case "TSA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TSA
		return
	case "CBA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CBA
		return
	case "RCA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_RCA
		return
	case "RAS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_RAS
		return
	case "AWY":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_AWY
		return
	case "MTR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_MTR
		return
	case "P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_P
		return
	case "R":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_R
		return
	case "D":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_D
		return
	case "ADIZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ADIZ
		return
	case "NO_FIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NO_FIR
		return
	case "PART":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_PART
		return
	case "CLASS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CLASS
		return
	case "POLITICAL":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_POLITICAL
		return
	case "D_OTHER":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_D_OTHER
		return
	case "TRA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TRA
		return
	case "A":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_A
		return
	case "W":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_W
		return
	case "PROTECT":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_PROTECT
		return
	case "AMA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_AMA
		return
	case "ASR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ASR
		return
	case "ADV":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ADV
		return
	case "UADV":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UADV
		return
	case "ATZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ATZ
		return
	case "ATZ_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ATZ_P
		return
	case "HTZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_HTZ
		return
	case "NAS_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NAS_P
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspacebasetype_member_1 *Enum_CodeAirspaceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspaceBaseType_Member_1_NAS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NAS
	case "Enum_CodeAirspaceBaseType_Member_1_FIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_FIR
	case "Enum_CodeAirspaceBaseType_Member_1_FIR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_FIR_P
	case "Enum_CodeAirspaceBaseType_Member_1_UIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UIR
	case "Enum_CodeAirspaceBaseType_Member_1_UIR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UIR_P
	case "Enum_CodeAirspaceBaseType_Member_1_CTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTA
	case "Enum_CodeAirspaceBaseType_Member_1_CTA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTA_P
	case "Enum_CodeAirspaceBaseType_Member_1_OCA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OCA_P
	case "Enum_CodeAirspaceBaseType_Member_1_OCA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OCA
	case "Enum_CodeAirspaceBaseType_Member_1_UTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UTA
	case "Enum_CodeAirspaceBaseType_Member_1_UTA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UTA_P
	case "Enum_CodeAirspaceBaseType_Member_1_TMA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TMA
	case "Enum_CodeAirspaceBaseType_Member_1_TMA_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TMA_P
	case "Enum_CodeAirspaceBaseType_Member_1_CTR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTR
	case "Enum_CodeAirspaceBaseType_Member_1_CTR_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CTR_P
	case "Enum_CodeAirspaceBaseType_Member_1_OTA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_OTA
	case "Enum_CodeAirspaceBaseType_Member_1_SECTOR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_SECTOR
	case "Enum_CodeAirspaceBaseType_Member_1_SECTOR_C":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_SECTOR_C
	case "Enum_CodeAirspaceBaseType_Member_1_TSA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TSA
	case "Enum_CodeAirspaceBaseType_Member_1_CBA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CBA
	case "Enum_CodeAirspaceBaseType_Member_1_RCA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_RCA
	case "Enum_CodeAirspaceBaseType_Member_1_RAS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_RAS
	case "Enum_CodeAirspaceBaseType_Member_1_AWY":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_AWY
	case "Enum_CodeAirspaceBaseType_Member_1_MTR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_MTR
	case "Enum_CodeAirspaceBaseType_Member_1_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_P
	case "Enum_CodeAirspaceBaseType_Member_1_R":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_R
	case "Enum_CodeAirspaceBaseType_Member_1_D":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_D
	case "Enum_CodeAirspaceBaseType_Member_1_ADIZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ADIZ
	case "Enum_CodeAirspaceBaseType_Member_1_NO_FIR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NO_FIR
	case "Enum_CodeAirspaceBaseType_Member_1_PART":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_PART
	case "Enum_CodeAirspaceBaseType_Member_1_CLASS":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_CLASS
	case "Enum_CodeAirspaceBaseType_Member_1_POLITICAL":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_POLITICAL
	case "Enum_CodeAirspaceBaseType_Member_1_D_OTHER":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_D_OTHER
	case "Enum_CodeAirspaceBaseType_Member_1_TRA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_TRA
	case "Enum_CodeAirspaceBaseType_Member_1_A":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_A
	case "Enum_CodeAirspaceBaseType_Member_1_W":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_W
	case "Enum_CodeAirspaceBaseType_Member_1_PROTECT":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_PROTECT
	case "Enum_CodeAirspaceBaseType_Member_1_AMA":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_AMA
	case "Enum_CodeAirspaceBaseType_Member_1_ASR":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ASR
	case "Enum_CodeAirspaceBaseType_Member_1_ADV":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ADV
	case "Enum_CodeAirspaceBaseType_Member_1_UADV":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_UADV
	case "Enum_CodeAirspaceBaseType_Member_1_ATZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ATZ
	case "Enum_CodeAirspaceBaseType_Member_1_ATZ_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_ATZ_P
	case "Enum_CodeAirspaceBaseType_Member_1_HTZ":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_HTZ
	case "Enum_CodeAirspaceBaseType_Member_1_NAS_P":
		*enum_codeairspacebasetype_member_1 = Enum_CodeAirspaceBaseType_Member_1_NAS_P
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspacebasetype_member_1 *Enum_CodeAirspaceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceBaseType_Member_1_NAS:
		res = "Enum_CodeAirspaceBaseType_Member_1_NAS"
	case Enum_CodeAirspaceBaseType_Member_1_FIR:
		res = "Enum_CodeAirspaceBaseType_Member_1_FIR"
	case Enum_CodeAirspaceBaseType_Member_1_FIR_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_FIR_P"
	case Enum_CodeAirspaceBaseType_Member_1_UIR:
		res = "Enum_CodeAirspaceBaseType_Member_1_UIR"
	case Enum_CodeAirspaceBaseType_Member_1_UIR_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_UIR_P"
	case Enum_CodeAirspaceBaseType_Member_1_CTA:
		res = "Enum_CodeAirspaceBaseType_Member_1_CTA"
	case Enum_CodeAirspaceBaseType_Member_1_CTA_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_CTA_P"
	case Enum_CodeAirspaceBaseType_Member_1_OCA_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_OCA_P"
	case Enum_CodeAirspaceBaseType_Member_1_OCA:
		res = "Enum_CodeAirspaceBaseType_Member_1_OCA"
	case Enum_CodeAirspaceBaseType_Member_1_UTA:
		res = "Enum_CodeAirspaceBaseType_Member_1_UTA"
	case Enum_CodeAirspaceBaseType_Member_1_UTA_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_UTA_P"
	case Enum_CodeAirspaceBaseType_Member_1_TMA:
		res = "Enum_CodeAirspaceBaseType_Member_1_TMA"
	case Enum_CodeAirspaceBaseType_Member_1_TMA_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_TMA_P"
	case Enum_CodeAirspaceBaseType_Member_1_CTR:
		res = "Enum_CodeAirspaceBaseType_Member_1_CTR"
	case Enum_CodeAirspaceBaseType_Member_1_CTR_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_CTR_P"
	case Enum_CodeAirspaceBaseType_Member_1_OTA:
		res = "Enum_CodeAirspaceBaseType_Member_1_OTA"
	case Enum_CodeAirspaceBaseType_Member_1_SECTOR:
		res = "Enum_CodeAirspaceBaseType_Member_1_SECTOR"
	case Enum_CodeAirspaceBaseType_Member_1_SECTOR_C:
		res = "Enum_CodeAirspaceBaseType_Member_1_SECTOR_C"
	case Enum_CodeAirspaceBaseType_Member_1_TSA:
		res = "Enum_CodeAirspaceBaseType_Member_1_TSA"
	case Enum_CodeAirspaceBaseType_Member_1_CBA:
		res = "Enum_CodeAirspaceBaseType_Member_1_CBA"
	case Enum_CodeAirspaceBaseType_Member_1_RCA:
		res = "Enum_CodeAirspaceBaseType_Member_1_RCA"
	case Enum_CodeAirspaceBaseType_Member_1_RAS:
		res = "Enum_CodeAirspaceBaseType_Member_1_RAS"
	case Enum_CodeAirspaceBaseType_Member_1_AWY:
		res = "Enum_CodeAirspaceBaseType_Member_1_AWY"
	case Enum_CodeAirspaceBaseType_Member_1_MTR:
		res = "Enum_CodeAirspaceBaseType_Member_1_MTR"
	case Enum_CodeAirspaceBaseType_Member_1_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_P"
	case Enum_CodeAirspaceBaseType_Member_1_R:
		res = "Enum_CodeAirspaceBaseType_Member_1_R"
	case Enum_CodeAirspaceBaseType_Member_1_D:
		res = "Enum_CodeAirspaceBaseType_Member_1_D"
	case Enum_CodeAirspaceBaseType_Member_1_ADIZ:
		res = "Enum_CodeAirspaceBaseType_Member_1_ADIZ"
	case Enum_CodeAirspaceBaseType_Member_1_NO_FIR:
		res = "Enum_CodeAirspaceBaseType_Member_1_NO_FIR"
	case Enum_CodeAirspaceBaseType_Member_1_PART:
		res = "Enum_CodeAirspaceBaseType_Member_1_PART"
	case Enum_CodeAirspaceBaseType_Member_1_CLASS:
		res = "Enum_CodeAirspaceBaseType_Member_1_CLASS"
	case Enum_CodeAirspaceBaseType_Member_1_POLITICAL:
		res = "Enum_CodeAirspaceBaseType_Member_1_POLITICAL"
	case Enum_CodeAirspaceBaseType_Member_1_D_OTHER:
		res = "Enum_CodeAirspaceBaseType_Member_1_D_OTHER"
	case Enum_CodeAirspaceBaseType_Member_1_TRA:
		res = "Enum_CodeAirspaceBaseType_Member_1_TRA"
	case Enum_CodeAirspaceBaseType_Member_1_A:
		res = "Enum_CodeAirspaceBaseType_Member_1_A"
	case Enum_CodeAirspaceBaseType_Member_1_W:
		res = "Enum_CodeAirspaceBaseType_Member_1_W"
	case Enum_CodeAirspaceBaseType_Member_1_PROTECT:
		res = "Enum_CodeAirspaceBaseType_Member_1_PROTECT"
	case Enum_CodeAirspaceBaseType_Member_1_AMA:
		res = "Enum_CodeAirspaceBaseType_Member_1_AMA"
	case Enum_CodeAirspaceBaseType_Member_1_ASR:
		res = "Enum_CodeAirspaceBaseType_Member_1_ASR"
	case Enum_CodeAirspaceBaseType_Member_1_ADV:
		res = "Enum_CodeAirspaceBaseType_Member_1_ADV"
	case Enum_CodeAirspaceBaseType_Member_1_UADV:
		res = "Enum_CodeAirspaceBaseType_Member_1_UADV"
	case Enum_CodeAirspaceBaseType_Member_1_ATZ:
		res = "Enum_CodeAirspaceBaseType_Member_1_ATZ"
	case Enum_CodeAirspaceBaseType_Member_1_ATZ_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_ATZ_P"
	case Enum_CodeAirspaceBaseType_Member_1_HTZ:
		res = "Enum_CodeAirspaceBaseType_Member_1_HTZ"
	case Enum_CodeAirspaceBaseType_Member_1_NAS_P:
		res = "Enum_CodeAirspaceBaseType_Member_1_NAS_P"
	}
	return
}

func (enum_codeairspacebasetype_member_1 Enum_CodeAirspaceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_NAS")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_FIR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_FIR_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_UIR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_UIR_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CTA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CTA_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_OCA_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_OCA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_UTA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_UTA_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_TMA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_TMA_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CTR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CTR_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_OTA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_SECTOR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_SECTOR_C")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_TSA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CBA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_RCA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_RAS")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_AWY")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_MTR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_R")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_D")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_ADIZ")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_NO_FIR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_PART")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_CLASS")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_POLITICAL")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_D_OTHER")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_TRA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_A")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_W")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_PROTECT")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_AMA")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_ASR")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_ADV")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_UADV")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_ATZ")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_ATZ_P")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_HTZ")
	res = append(res, "Enum_CodeAirspaceBaseType_Member_1_NAS_P")

	return
}

func (enum_codeairspacebasetype_member_1 Enum_CodeAirspaceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NAS")
	res = append(res, "FIR")
	res = append(res, "FIR_P")
	res = append(res, "UIR")
	res = append(res, "UIR_P")
	res = append(res, "CTA")
	res = append(res, "CTA_P")
	res = append(res, "OCA_P")
	res = append(res, "OCA")
	res = append(res, "UTA")
	res = append(res, "UTA_P")
	res = append(res, "TMA")
	res = append(res, "TMA_P")
	res = append(res, "CTR")
	res = append(res, "CTR_P")
	res = append(res, "OTA")
	res = append(res, "SECTOR")
	res = append(res, "SECTOR_C")
	res = append(res, "TSA")
	res = append(res, "CBA")
	res = append(res, "RCA")
	res = append(res, "RAS")
	res = append(res, "AWY")
	res = append(res, "MTR")
	res = append(res, "P")
	res = append(res, "R")
	res = append(res, "D")
	res = append(res, "ADIZ")
	res = append(res, "NO_FIR")
	res = append(res, "PART")
	res = append(res, "CLASS")
	res = append(res, "POLITICAL")
	res = append(res, "D_OTHER")
	res = append(res, "TRA")
	res = append(res, "A")
	res = append(res, "W")
	res = append(res, "PROTECT")
	res = append(res, "AMA")
	res = append(res, "ASR")
	res = append(res, "ADV")
	res = append(res, "UADV")
	res = append(res, "ATZ")
	res = append(res, "ATZ_P")
	res = append(res, "HTZ")
	res = append(res, "NAS_P")

	return
}

// Utility function for Enum_CodeAirspaceClassificationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspaceclassificationbasetype_member_1 Enum_CodeAirspaceClassificationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspaceclassificationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceClassificationBaseType_Member_1_A:
		res = "A"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_B:
		res = "B"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_C:
		res = "C"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_D:
		res = "D"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_E:
		res = "E"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_F:
		res = "F"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_G:
		res = "G"
	}
	return
}

func (enum_codeairspaceclassificationbasetype_member_1 *Enum_CodeAirspaceClassificationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_A
		return
	case "B":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_B
		return
	case "C":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_C
		return
	case "D":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_D
		return
	case "E":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_E
		return
	case "F":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_F
		return
	case "G":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_G
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspaceclassificationbasetype_member_1 *Enum_CodeAirspaceClassificationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_A":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_A
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_B":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_B
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_C":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_C
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_D":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_D
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_E":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_E
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_F":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_F
	case "Enum_CodeAirspaceClassificationBaseType_Member_1_G":
		*enum_codeairspaceclassificationbasetype_member_1 = Enum_CodeAirspaceClassificationBaseType_Member_1_G
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspaceclassificationbasetype_member_1 *Enum_CodeAirspaceClassificationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspaceclassificationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceClassificationBaseType_Member_1_A:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_A"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_B:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_B"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_C:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_C"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_D:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_D"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_E:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_E"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_F:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_F"
	case Enum_CodeAirspaceClassificationBaseType_Member_1_G:
		res = "Enum_CodeAirspaceClassificationBaseType_Member_1_G"
	}
	return
}

func (enum_codeairspaceclassificationbasetype_member_1 Enum_CodeAirspaceClassificationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_A")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_B")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_C")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_D")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_E")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_F")
	res = append(res, "Enum_CodeAirspaceClassificationBaseType_Member_1_G")

	return
}

func (enum_codeairspaceclassificationbasetype_member_1 Enum_CodeAirspaceClassificationBaseType_Member_1) CodeValues() (res []string) {

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

// Utility function for Enum_CodeAirspaceDependencyBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspacedependencybasetype_member_1 Enum_CodeAirspaceDependencyBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspacedependencybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY:
		res = "FULL_GEOMETRY"
	case Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION:
		res = "HORZ_PROJECTION"
	}
	return
}

func (enum_codeairspacedependencybasetype_member_1 *Enum_CodeAirspaceDependencyBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FULL_GEOMETRY":
		*enum_codeairspacedependencybasetype_member_1 = Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY
		return
	case "HORZ_PROJECTION":
		*enum_codeairspacedependencybasetype_member_1 = Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspacedependencybasetype_member_1 *Enum_CodeAirspaceDependencyBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY":
		*enum_codeairspacedependencybasetype_member_1 = Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY
	case "Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION":
		*enum_codeairspacedependencybasetype_member_1 = Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspacedependencybasetype_member_1 *Enum_CodeAirspaceDependencyBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspacedependencybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY:
		res = "Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY"
	case Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION:
		res = "Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION"
	}
	return
}

func (enum_codeairspacedependencybasetype_member_1 Enum_CodeAirspaceDependencyBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspaceDependencyBaseType_Member_1_FULL_GEOMETRY")
	res = append(res, "Enum_CodeAirspaceDependencyBaseType_Member_1_HORZ_PROJECTION")

	return
}

func (enum_codeairspacedependencybasetype_member_1 Enum_CodeAirspaceDependencyBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FULL_GEOMETRY")
	res = append(res, "HORZ_PROJECTION")

	return
}

// Utility function for Enum_CodeAltitudeUseBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codealtitudeusebasetype_member_1 Enum_CodeAltitudeUseBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codealtitudeusebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER:
		res = "ABOVE_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER:
		res = "BELOW_UPPER"
	case Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER:
		res = "AT_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN:
		res = "BETWEEN"
	case Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED:
		res = "RECOMMENDED"
	case Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER:
		res = "EXPECT_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED:
		res = "AS_ASSIGNED"
	}
	return
}

func (enum_codealtitudeusebasetype_member_1 *Enum_CodeAltitudeUseBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ABOVE_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER
		return
	case "BELOW_UPPER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER
		return
	case "AT_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER
		return
	case "BETWEEN":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN
		return
	case "RECOMMENDED":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED
		return
	case "EXPECT_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER
		return
	case "AS_ASSIGNED":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codealtitudeusebasetype_member_1 *Enum_CodeAltitudeUseBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER
	case "Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER
	case "Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER
	case "Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN
	case "Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED
	case "Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER
	case "Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED":
		*enum_codealtitudeusebasetype_member_1 = Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codealtitudeusebasetype_member_1 *Enum_CodeAltitudeUseBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codealtitudeusebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER"
	case Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN"
	case Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED"
	case Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER"
	case Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED:
		res = "Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED"
	}
	return
}

func (enum_codealtitudeusebasetype_member_1 Enum_CodeAltitudeUseBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_ABOVE_LOWER")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_BELOW_UPPER")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_AT_LOWER")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_BETWEEN")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_RECOMMENDED")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_EXPECT_LOWER")
	res = append(res, "Enum_CodeAltitudeUseBaseType_Member_1_AS_ASSIGNED")

	return
}

func (enum_codealtitudeusebasetype_member_1 Enum_CodeAltitudeUseBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ABOVE_LOWER")
	res = append(res, "BELOW_UPPER")
	res = append(res, "AT_LOWER")
	res = append(res, "BETWEEN")
	res = append(res, "RECOMMENDED")
	res = append(res, "EXPECT_LOWER")
	res = append(res, "AS_ASSIGNED")

	return
}

// Utility function for Enum_CodeAuthorityBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeauthoritybasetype_member_1 Enum_CodeAuthorityBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeauthoritybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAuthorityBaseType_Member_1_OWN:
		res = "OWN"
	case Enum_CodeAuthorityBaseType_Member_1_DLGT:
		res = "DLGT"
	case Enum_CodeAuthorityBaseType_Member_1_AIS:
		res = "AIS"
	}
	return
}

func (enum_codeauthoritybasetype_member_1 *Enum_CodeAuthorityBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OWN":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_OWN
		return
	case "DLGT":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_DLGT
		return
	case "AIS":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_AIS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeauthoritybasetype_member_1 *Enum_CodeAuthorityBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAuthorityBaseType_Member_1_OWN":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_OWN
	case "Enum_CodeAuthorityBaseType_Member_1_DLGT":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_DLGT
	case "Enum_CodeAuthorityBaseType_Member_1_AIS":
		*enum_codeauthoritybasetype_member_1 = Enum_CodeAuthorityBaseType_Member_1_AIS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeauthoritybasetype_member_1 *Enum_CodeAuthorityBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeauthoritybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAuthorityBaseType_Member_1_OWN:
		res = "Enum_CodeAuthorityBaseType_Member_1_OWN"
	case Enum_CodeAuthorityBaseType_Member_1_DLGT:
		res = "Enum_CodeAuthorityBaseType_Member_1_DLGT"
	case Enum_CodeAuthorityBaseType_Member_1_AIS:
		res = "Enum_CodeAuthorityBaseType_Member_1_AIS"
	}
	return
}

func (enum_codeauthoritybasetype_member_1 Enum_CodeAuthorityBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAuthorityBaseType_Member_1_OWN")
	res = append(res, "Enum_CodeAuthorityBaseType_Member_1_DLGT")
	res = append(res, "Enum_CodeAuthorityBaseType_Member_1_AIS")

	return
}

func (enum_codeauthoritybasetype_member_1 Enum_CodeAuthorityBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OWN")
	res = append(res, "DLGT")
	res = append(res, "AIS")

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

// Utility function for Enum_CodeGeoBorderBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codegeoborderbasetype_member_1 Enum_CodeGeoBorderBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codegeoborderbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGeoBorderBaseType_Member_1_STATE:
		res = "STATE"
	case Enum_CodeGeoBorderBaseType_Member_1_WATER:
		res = "WATER"
	case Enum_CodeGeoBorderBaseType_Member_1_COAST:
		res = "COAST"
	case Enum_CodeGeoBorderBaseType_Member_1_RIVER:
		res = "RIVER"
	case Enum_CodeGeoBorderBaseType_Member_1_BANK:
		res = "BANK"
	}
	return
}

func (enum_codegeoborderbasetype_member_1 *Enum_CodeGeoBorderBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STATE":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_STATE
		return
	case "WATER":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_WATER
		return
	case "COAST":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_COAST
		return
	case "RIVER":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_RIVER
		return
	case "BANK":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_BANK
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codegeoborderbasetype_member_1 *Enum_CodeGeoBorderBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeGeoBorderBaseType_Member_1_STATE":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_STATE
	case "Enum_CodeGeoBorderBaseType_Member_1_WATER":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_WATER
	case "Enum_CodeGeoBorderBaseType_Member_1_COAST":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_COAST
	case "Enum_CodeGeoBorderBaseType_Member_1_RIVER":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_RIVER
	case "Enum_CodeGeoBorderBaseType_Member_1_BANK":
		*enum_codegeoborderbasetype_member_1 = Enum_CodeGeoBorderBaseType_Member_1_BANK
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codegeoborderbasetype_member_1 *Enum_CodeGeoBorderBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codegeoborderbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeGeoBorderBaseType_Member_1_STATE:
		res = "Enum_CodeGeoBorderBaseType_Member_1_STATE"
	case Enum_CodeGeoBorderBaseType_Member_1_WATER:
		res = "Enum_CodeGeoBorderBaseType_Member_1_WATER"
	case Enum_CodeGeoBorderBaseType_Member_1_COAST:
		res = "Enum_CodeGeoBorderBaseType_Member_1_COAST"
	case Enum_CodeGeoBorderBaseType_Member_1_RIVER:
		res = "Enum_CodeGeoBorderBaseType_Member_1_RIVER"
	case Enum_CodeGeoBorderBaseType_Member_1_BANK:
		res = "Enum_CodeGeoBorderBaseType_Member_1_BANK"
	}
	return
}

func (enum_codegeoborderbasetype_member_1 Enum_CodeGeoBorderBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeGeoBorderBaseType_Member_1_STATE")
	res = append(res, "Enum_CodeGeoBorderBaseType_Member_1_WATER")
	res = append(res, "Enum_CodeGeoBorderBaseType_Member_1_COAST")
	res = append(res, "Enum_CodeGeoBorderBaseType_Member_1_RIVER")
	res = append(res, "Enum_CodeGeoBorderBaseType_Member_1_BANK")

	return
}

func (enum_codegeoborderbasetype_member_1 Enum_CodeGeoBorderBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STATE")
	res = append(res, "WATER")
	res = append(res, "COAST")
	res = append(res, "RIVER")
	res = append(res, "BANK")

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

// Utility function for Enum_CodeStatusAirspaceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusairspacebasetype_member_1 Enum_CodeStatusAirspaceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusairspacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION:
		res = "AVBL_FOR_ACTIVATION"
	case Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE:
		res = "ACTIVE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE:
		res = "IN_USE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE:
		res = "INACTIVE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT:
		res = "INTERMITTENT"
	}
	return
}

func (enum_codestatusairspacebasetype_member_1 *Enum_CodeStatusAirspaceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AVBL_FOR_ACTIVATION":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION
		return
	case "ACTIVE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE
		return
	case "IN_USE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE
		return
	case "INACTIVE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE
		return
	case "INTERMITTENT":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusairspacebasetype_member_1 *Enum_CodeStatusAirspaceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION
	case "Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE
	case "Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE
	case "Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE
	case "Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT":
		*enum_codestatusairspacebasetype_member_1 = Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusairspacebasetype_member_1 *Enum_CodeStatusAirspaceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusairspacebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION:
		res = "Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION"
	case Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE:
		res = "Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE:
		res = "Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE:
		res = "Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE"
	case Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT:
		res = "Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT"
	}
	return
}

func (enum_codestatusairspacebasetype_member_1 Enum_CodeStatusAirspaceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusAirspaceBaseType_Member_1_AVBL_FOR_ACTIVATION")
	res = append(res, "Enum_CodeStatusAirspaceBaseType_Member_1_ACTIVE")
	res = append(res, "Enum_CodeStatusAirspaceBaseType_Member_1_IN_USE")
	res = append(res, "Enum_CodeStatusAirspaceBaseType_Member_1_INACTIVE")
	res = append(res, "Enum_CodeStatusAirspaceBaseType_Member_1_INTERMITTENT")

	return
}

func (enum_codestatusairspacebasetype_member_1 Enum_CodeStatusAirspaceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AVBL_FOR_ACTIVATION")
	res = append(res, "ACTIVE")
	res = append(res, "IN_USE")
	res = append(res, "INACTIVE")
	res = append(res, "INTERMITTENT")

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

// Utility function for Enum_CodeVerticalReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeverticalreferencebasetype_member_1 Enum_CodeVerticalReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeverticalreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalReferenceBaseType_Member_1_SFC:
		res = "SFC"
	case Enum_CodeVerticalReferenceBaseType_Member_1_MSL:
		res = "MSL"
	case Enum_CodeVerticalReferenceBaseType_Member_1_W84:
		res = "W84"
	case Enum_CodeVerticalReferenceBaseType_Member_1_STD:
		res = "STD"
	}
	return
}

func (enum_codeverticalreferencebasetype_member_1 *Enum_CodeVerticalReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SFC":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_SFC
		return
	case "MSL":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_MSL
		return
	case "W84":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_W84
		return
	case "STD":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_STD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeverticalreferencebasetype_member_1 *Enum_CodeVerticalReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVerticalReferenceBaseType_Member_1_SFC":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_SFC
	case "Enum_CodeVerticalReferenceBaseType_Member_1_MSL":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_MSL
	case "Enum_CodeVerticalReferenceBaseType_Member_1_W84":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_W84
	case "Enum_CodeVerticalReferenceBaseType_Member_1_STD":
		*enum_codeverticalreferencebasetype_member_1 = Enum_CodeVerticalReferenceBaseType_Member_1_STD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeverticalreferencebasetype_member_1 *Enum_CodeVerticalReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeverticalreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVerticalReferenceBaseType_Member_1_SFC:
		res = "Enum_CodeVerticalReferenceBaseType_Member_1_SFC"
	case Enum_CodeVerticalReferenceBaseType_Member_1_MSL:
		res = "Enum_CodeVerticalReferenceBaseType_Member_1_MSL"
	case Enum_CodeVerticalReferenceBaseType_Member_1_W84:
		res = "Enum_CodeVerticalReferenceBaseType_Member_1_W84"
	case Enum_CodeVerticalReferenceBaseType_Member_1_STD:
		res = "Enum_CodeVerticalReferenceBaseType_Member_1_STD"
	}
	return
}

func (enum_codeverticalreferencebasetype_member_1 Enum_CodeVerticalReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVerticalReferenceBaseType_Member_1_SFC")
	res = append(res, "Enum_CodeVerticalReferenceBaseType_Member_1_MSL")
	res = append(res, "Enum_CodeVerticalReferenceBaseType_Member_1_W84")
	res = append(res, "Enum_CodeVerticalReferenceBaseType_Member_1_STD")

	return
}

func (enum_codeverticalreferencebasetype_member_1 Enum_CodeVerticalReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SFC")
	res = append(res, "MSL")
	res = append(res, "W84")
	res = append(res, "STD")

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
