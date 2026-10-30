// generated code - do not edit
package aixmdatatypes

// insertion point of enum utility functions
// Utility function for Enum_CodeATCReportingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeatcreportingbasetype_member_1 Enum_CodeATCReportingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeatcreportingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeATCReportingBaseType_Member_1_COMPULSORY:
		res = "COMPULSORY"
	case Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST:
		res = "ON_REQUEST"
	case Enum_CodeATCReportingBaseType_Member_1_NO_REPORT:
		res = "NO_REPORT"
	}
	return
}

func (enum_codeatcreportingbasetype_member_1 *Enum_CodeATCReportingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "COMPULSORY":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_COMPULSORY
		return
	case "ON_REQUEST":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST
		return
	case "NO_REPORT":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_NO_REPORT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeatcreportingbasetype_member_1 *Enum_CodeATCReportingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeATCReportingBaseType_Member_1_COMPULSORY":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_COMPULSORY
	case "Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST
	case "Enum_CodeATCReportingBaseType_Member_1_NO_REPORT":
		*enum_codeatcreportingbasetype_member_1 = Enum_CodeATCReportingBaseType_Member_1_NO_REPORT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeatcreportingbasetype_member_1 *Enum_CodeATCReportingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeatcreportingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeATCReportingBaseType_Member_1_COMPULSORY:
		res = "Enum_CodeATCReportingBaseType_Member_1_COMPULSORY"
	case Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST:
		res = "Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST"
	case Enum_CodeATCReportingBaseType_Member_1_NO_REPORT:
		res = "Enum_CodeATCReportingBaseType_Member_1_NO_REPORT"
	}
	return
}

func (enum_codeatcreportingbasetype_member_1 Enum_CodeATCReportingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeATCReportingBaseType_Member_1_COMPULSORY")
	res = append(res, "Enum_CodeATCReportingBaseType_Member_1_ON_REQUEST")
	res = append(res, "Enum_CodeATCReportingBaseType_Member_1_NO_REPORT")

	return
}

func (enum_codeatcreportingbasetype_member_1 Enum_CodeATCReportingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "COMPULSORY")
	res = append(res, "ON_REQUEST")
	res = append(res, "NO_REPORT")

	return
}

// Utility function for Enum_CodeAerialRefuellingPointBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaerialrefuellingpointbasetype_member_1 Enum_CodeAerialRefuellingPointBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaerialrefuellingpointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL:
		res = "INITIAL"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL:
		res = "CONTROL"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK:
		res = "CHECK"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT:
		res = "EXIT"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY:
		res = "ENTRY"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR:
		res = "ANCHOR"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN:
		res = "PATTERN"
	}
	return
}

func (enum_codeaerialrefuellingpointbasetype_member_1 *Enum_CodeAerialRefuellingPointBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "INITIAL":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL
		return
	case "CONTROL":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL
		return
	case "CHECK":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK
		return
	case "EXIT":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT
		return
	case "ENTRY":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY
		return
	case "ANCHOR":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR
		return
	case "PATTERN":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaerialrefuellingpointbasetype_member_1 *Enum_CodeAerialRefuellingPointBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR
	case "Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN":
		*enum_codeaerialrefuellingpointbasetype_member_1 = Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaerialrefuellingpointbasetype_member_1 *Enum_CodeAerialRefuellingPointBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaerialrefuellingpointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR"
	case Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN:
		res = "Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN"
	}
	return
}

func (enum_codeaerialrefuellingpointbasetype_member_1 Enum_CodeAerialRefuellingPointBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_INITIAL")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_CONTROL")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_CHECK")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_EXIT")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_ENTRY")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_ANCHOR")
	res = append(res, "Enum_CodeAerialRefuellingPointBaseType_Member_1_PATTERN")

	return
}

func (enum_codeaerialrefuellingpointbasetype_member_1 Enum_CodeAerialRefuellingPointBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "INITIAL")
	res = append(res, "CONTROL")
	res = append(res, "CHECK")
	res = append(res, "EXIT")
	res = append(res, "ENTRY")
	res = append(res, "ANCHOR")
	res = append(res, "PATTERN")

	return
}

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

// Utility function for Enum_CodeAltitudeAdjustmentBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codealtitudeadjustmentbasetype_member_1 Enum_CodeAltitudeAdjustmentBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codealtitudeadjustmentbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA:
		res = "RA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS:
		res = "AS"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT:
		res = "AT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC:
		res = "AC"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI:
		res = "SI"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL:
		res = "XL"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR:
		res = "PR"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA:
		res = "HAA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA:
		res = "MA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT:
		res = "PT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG:
		res = "DG"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS:
		res = "GS"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA:
		res = "CA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT:
		res = "MT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH:
		res = "MAH"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA:
		res = "SA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO:
		res = "AAO"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA:
		res = "VA"
	}
	return
}

func (enum_codealtitudeadjustmentbasetype_member_1 *Enum_CodeAltitudeAdjustmentBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA
		return
	case "AS":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS
		return
	case "AT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT
		return
	case "AC":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC
		return
	case "SI":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI
		return
	case "XL":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL
		return
	case "PR":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR
		return
	case "HAA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA
		return
	case "MA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA
		return
	case "PT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT
		return
	case "DG":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG
		return
	case "GS":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS
		return
	case "CA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA
		return
	case "MT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT
		return
	case "MAH":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH
		return
	case "SA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA
		return
	case "AAO":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO
		return
	case "VA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codealtitudeadjustmentbasetype_member_1 *Enum_CodeAltitudeAdjustmentBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO
	case "Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA":
		*enum_codealtitudeadjustmentbasetype_member_1 = Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codealtitudeadjustmentbasetype_member_1 *Enum_CodeAltitudeAdjustmentBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codealtitudeadjustmentbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO"
	case Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA:
		res = "Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA"
	}
	return
}

func (enum_codealtitudeadjustmentbasetype_member_1 Enum_CodeAltitudeAdjustmentBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_RA")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AS")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AT")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AC")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SI")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_XL")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PR")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_HAA")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MA")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_PT")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_DG")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_GS")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_CA")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MT")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_MAH")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_SA")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_AAO")
	res = append(res, "Enum_CodeAltitudeAdjustmentBaseType_Member_1_VA")

	return
}

func (enum_codealtitudeadjustmentbasetype_member_1 Enum_CodeAltitudeAdjustmentBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RA")
	res = append(res, "AS")
	res = append(res, "AT")
	res = append(res, "AC")
	res = append(res, "SI")
	res = append(res, "XL")
	res = append(res, "PR")
	res = append(res, "HAA")
	res = append(res, "MA")
	res = append(res, "PT")
	res = append(res, "DG")
	res = append(res, "GS")
	res = append(res, "CA")
	res = append(res, "MT")
	res = append(res, "MAH")
	res = append(res, "SA")
	res = append(res, "AAO")
	res = append(res, "VA")

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

// Utility function for Enum_CodeApproachBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachbasetype_member_1 Enum_CodeApproachBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachBaseType_Member_1_ASR:
		res = "ASR"
	case Enum_CodeApproachBaseType_Member_1_ARA:
		res = "ARA"
	case Enum_CodeApproachBaseType_Member_1_ARSR:
		res = "ARSR"
	case Enum_CodeApproachBaseType_Member_1_PAR:
		res = "PAR"
	case Enum_CodeApproachBaseType_Member_1_ILS:
		res = "ILS"
	case Enum_CodeApproachBaseType_Member_1_ILS_DME:
		res = "ILS_DME"
	case Enum_CodeApproachBaseType_Member_1_ILS_PRM:
		res = "ILS_PRM"
	case Enum_CodeApproachBaseType_Member_1_LDA:
		res = "LDA"
	case Enum_CodeApproachBaseType_Member_1_LDA_DME:
		res = "LDA_DME"
	case Enum_CodeApproachBaseType_Member_1_LOC:
		res = "LOC"
	case Enum_CodeApproachBaseType_Member_1_LOC_BC:
		res = "LOC_BC"
	case Enum_CodeApproachBaseType_Member_1_LOC_DME:
		res = "LOC_DME"
	case Enum_CodeApproachBaseType_Member_1_LOC_DME_BC:
		res = "LOC_DME_BC"
	case Enum_CodeApproachBaseType_Member_1_MLS:
		res = "MLS"
	case Enum_CodeApproachBaseType_Member_1_MLS_DME:
		res = "MLS_DME"
	case Enum_CodeApproachBaseType_Member_1_NDB:
		res = "NDB"
	case Enum_CodeApproachBaseType_Member_1_NDB_DME:
		res = "NDB_DME"
	case Enum_CodeApproachBaseType_Member_1_SDF:
		res = "SDF"
	case Enum_CodeApproachBaseType_Member_1_TLS:
		res = "TLS"
	case Enum_CodeApproachBaseType_Member_1_VOR:
		res = "VOR"
	case Enum_CodeApproachBaseType_Member_1_VOR_DME:
		res = "VOR_DME"
	case Enum_CodeApproachBaseType_Member_1_TACAN:
		res = "TACAN"
	case Enum_CodeApproachBaseType_Member_1_VORTAC:
		res = "VORTAC"
	case Enum_CodeApproachBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeApproachBaseType_Member_1_DME_DME:
		res = "DME_DME"
	case Enum_CodeApproachBaseType_Member_1_RNP:
		res = "RNP"
	case Enum_CodeApproachBaseType_Member_1_GPS:
		res = "GPS"
	case Enum_CodeApproachBaseType_Member_1_GLONASS:
		res = "GLONASS"
	case Enum_CodeApproachBaseType_Member_1_GALILEO:
		res = "GALILEO"
	case Enum_CodeApproachBaseType_Member_1_RNAV:
		res = "RNAV"
	case Enum_CodeApproachBaseType_Member_1_IGS:
		res = "IGS"
	}
	return
}

func (enum_codeapproachbasetype_member_1 *Enum_CodeApproachBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ASR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ASR
		return
	case "ARA":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ARA
		return
	case "ARSR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ARSR
		return
	case "PAR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_PAR
		return
	case "ILS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS
		return
	case "ILS_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS_DME
		return
	case "ILS_PRM":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS_PRM
		return
	case "LDA":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LDA
		return
	case "LDA_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LDA_DME
		return
	case "LOC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC
		return
	case "LOC_BC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_BC
		return
	case "LOC_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_DME
		return
	case "LOC_DME_BC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_DME_BC
		return
	case "MLS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_MLS
		return
	case "MLS_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_MLS_DME
		return
	case "NDB":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_NDB
		return
	case "NDB_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_NDB_DME
		return
	case "SDF":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_SDF
		return
	case "TLS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_TLS
		return
	case "VOR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VOR
		return
	case "VOR_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VOR_DME
		return
	case "TACAN":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_TACAN
		return
	case "VORTAC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VORTAC
		return
	case "DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_DME
		return
	case "DME_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_DME_DME
		return
	case "RNP":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_RNP
		return
	case "GPS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GPS
		return
	case "GLONASS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GLONASS
		return
	case "GALILEO":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GALILEO
		return
	case "RNAV":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_RNAV
		return
	case "IGS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_IGS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachbasetype_member_1 *Enum_CodeApproachBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachBaseType_Member_1_ASR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ASR
	case "Enum_CodeApproachBaseType_Member_1_ARA":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ARA
	case "Enum_CodeApproachBaseType_Member_1_ARSR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ARSR
	case "Enum_CodeApproachBaseType_Member_1_PAR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_PAR
	case "Enum_CodeApproachBaseType_Member_1_ILS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS
	case "Enum_CodeApproachBaseType_Member_1_ILS_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS_DME
	case "Enum_CodeApproachBaseType_Member_1_ILS_PRM":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_ILS_PRM
	case "Enum_CodeApproachBaseType_Member_1_LDA":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LDA
	case "Enum_CodeApproachBaseType_Member_1_LDA_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LDA_DME
	case "Enum_CodeApproachBaseType_Member_1_LOC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC
	case "Enum_CodeApproachBaseType_Member_1_LOC_BC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_BC
	case "Enum_CodeApproachBaseType_Member_1_LOC_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_DME
	case "Enum_CodeApproachBaseType_Member_1_LOC_DME_BC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_LOC_DME_BC
	case "Enum_CodeApproachBaseType_Member_1_MLS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_MLS
	case "Enum_CodeApproachBaseType_Member_1_MLS_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_MLS_DME
	case "Enum_CodeApproachBaseType_Member_1_NDB":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_NDB
	case "Enum_CodeApproachBaseType_Member_1_NDB_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_NDB_DME
	case "Enum_CodeApproachBaseType_Member_1_SDF":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_SDF
	case "Enum_CodeApproachBaseType_Member_1_TLS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_TLS
	case "Enum_CodeApproachBaseType_Member_1_VOR":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VOR
	case "Enum_CodeApproachBaseType_Member_1_VOR_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VOR_DME
	case "Enum_CodeApproachBaseType_Member_1_TACAN":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_TACAN
	case "Enum_CodeApproachBaseType_Member_1_VORTAC":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_VORTAC
	case "Enum_CodeApproachBaseType_Member_1_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_DME
	case "Enum_CodeApproachBaseType_Member_1_DME_DME":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_DME_DME
	case "Enum_CodeApproachBaseType_Member_1_RNP":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_RNP
	case "Enum_CodeApproachBaseType_Member_1_GPS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GPS
	case "Enum_CodeApproachBaseType_Member_1_GLONASS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GLONASS
	case "Enum_CodeApproachBaseType_Member_1_GALILEO":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_GALILEO
	case "Enum_CodeApproachBaseType_Member_1_RNAV":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_RNAV
	case "Enum_CodeApproachBaseType_Member_1_IGS":
		*enum_codeapproachbasetype_member_1 = Enum_CodeApproachBaseType_Member_1_IGS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachbasetype_member_1 *Enum_CodeApproachBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachBaseType_Member_1_ASR:
		res = "Enum_CodeApproachBaseType_Member_1_ASR"
	case Enum_CodeApproachBaseType_Member_1_ARA:
		res = "Enum_CodeApproachBaseType_Member_1_ARA"
	case Enum_CodeApproachBaseType_Member_1_ARSR:
		res = "Enum_CodeApproachBaseType_Member_1_ARSR"
	case Enum_CodeApproachBaseType_Member_1_PAR:
		res = "Enum_CodeApproachBaseType_Member_1_PAR"
	case Enum_CodeApproachBaseType_Member_1_ILS:
		res = "Enum_CodeApproachBaseType_Member_1_ILS"
	case Enum_CodeApproachBaseType_Member_1_ILS_DME:
		res = "Enum_CodeApproachBaseType_Member_1_ILS_DME"
	case Enum_CodeApproachBaseType_Member_1_ILS_PRM:
		res = "Enum_CodeApproachBaseType_Member_1_ILS_PRM"
	case Enum_CodeApproachBaseType_Member_1_LDA:
		res = "Enum_CodeApproachBaseType_Member_1_LDA"
	case Enum_CodeApproachBaseType_Member_1_LDA_DME:
		res = "Enum_CodeApproachBaseType_Member_1_LDA_DME"
	case Enum_CodeApproachBaseType_Member_1_LOC:
		res = "Enum_CodeApproachBaseType_Member_1_LOC"
	case Enum_CodeApproachBaseType_Member_1_LOC_BC:
		res = "Enum_CodeApproachBaseType_Member_1_LOC_BC"
	case Enum_CodeApproachBaseType_Member_1_LOC_DME:
		res = "Enum_CodeApproachBaseType_Member_1_LOC_DME"
	case Enum_CodeApproachBaseType_Member_1_LOC_DME_BC:
		res = "Enum_CodeApproachBaseType_Member_1_LOC_DME_BC"
	case Enum_CodeApproachBaseType_Member_1_MLS:
		res = "Enum_CodeApproachBaseType_Member_1_MLS"
	case Enum_CodeApproachBaseType_Member_1_MLS_DME:
		res = "Enum_CodeApproachBaseType_Member_1_MLS_DME"
	case Enum_CodeApproachBaseType_Member_1_NDB:
		res = "Enum_CodeApproachBaseType_Member_1_NDB"
	case Enum_CodeApproachBaseType_Member_1_NDB_DME:
		res = "Enum_CodeApproachBaseType_Member_1_NDB_DME"
	case Enum_CodeApproachBaseType_Member_1_SDF:
		res = "Enum_CodeApproachBaseType_Member_1_SDF"
	case Enum_CodeApproachBaseType_Member_1_TLS:
		res = "Enum_CodeApproachBaseType_Member_1_TLS"
	case Enum_CodeApproachBaseType_Member_1_VOR:
		res = "Enum_CodeApproachBaseType_Member_1_VOR"
	case Enum_CodeApproachBaseType_Member_1_VOR_DME:
		res = "Enum_CodeApproachBaseType_Member_1_VOR_DME"
	case Enum_CodeApproachBaseType_Member_1_TACAN:
		res = "Enum_CodeApproachBaseType_Member_1_TACAN"
	case Enum_CodeApproachBaseType_Member_1_VORTAC:
		res = "Enum_CodeApproachBaseType_Member_1_VORTAC"
	case Enum_CodeApproachBaseType_Member_1_DME:
		res = "Enum_CodeApproachBaseType_Member_1_DME"
	case Enum_CodeApproachBaseType_Member_1_DME_DME:
		res = "Enum_CodeApproachBaseType_Member_1_DME_DME"
	case Enum_CodeApproachBaseType_Member_1_RNP:
		res = "Enum_CodeApproachBaseType_Member_1_RNP"
	case Enum_CodeApproachBaseType_Member_1_GPS:
		res = "Enum_CodeApproachBaseType_Member_1_GPS"
	case Enum_CodeApproachBaseType_Member_1_GLONASS:
		res = "Enum_CodeApproachBaseType_Member_1_GLONASS"
	case Enum_CodeApproachBaseType_Member_1_GALILEO:
		res = "Enum_CodeApproachBaseType_Member_1_GALILEO"
	case Enum_CodeApproachBaseType_Member_1_RNAV:
		res = "Enum_CodeApproachBaseType_Member_1_RNAV"
	case Enum_CodeApproachBaseType_Member_1_IGS:
		res = "Enum_CodeApproachBaseType_Member_1_IGS"
	}
	return
}

func (enum_codeapproachbasetype_member_1 Enum_CodeApproachBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ASR")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ARA")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ARSR")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_PAR")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ILS")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ILS_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_ILS_PRM")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LDA")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LDA_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LOC")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LOC_BC")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LOC_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_LOC_DME_BC")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_MLS")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_MLS_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_NDB")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_NDB_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_SDF")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_TLS")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_VOR")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_VOR_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_TACAN")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_VORTAC")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_DME_DME")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_RNP")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_GPS")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_GLONASS")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_GALILEO")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_RNAV")
	res = append(res, "Enum_CodeApproachBaseType_Member_1_IGS")

	return
}

func (enum_codeapproachbasetype_member_1 Enum_CodeApproachBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ASR")
	res = append(res, "ARA")
	res = append(res, "ARSR")
	res = append(res, "PAR")
	res = append(res, "ILS")
	res = append(res, "ILS_DME")
	res = append(res, "ILS_PRM")
	res = append(res, "LDA")
	res = append(res, "LDA_DME")
	res = append(res, "LOC")
	res = append(res, "LOC_BC")
	res = append(res, "LOC_DME")
	res = append(res, "LOC_DME_BC")
	res = append(res, "MLS")
	res = append(res, "MLS_DME")
	res = append(res, "NDB")
	res = append(res, "NDB_DME")
	res = append(res, "SDF")
	res = append(res, "TLS")
	res = append(res, "VOR")
	res = append(res, "VOR_DME")
	res = append(res, "TACAN")
	res = append(res, "VORTAC")
	res = append(res, "DME")
	res = append(res, "DME_DME")
	res = append(res, "RNP")
	res = append(res, "GPS")
	res = append(res, "GLONASS")
	res = append(res, "GALILEO")
	res = append(res, "RNAV")
	res = append(res, "IGS")

	return
}

// Utility function for Enum_CodeApproachEquipmentAdditionalBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachequipmentadditionalbasetype_member_1 Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachequipmentadditionalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF:
		res = "ADF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR:
		res = "RADAR"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME:
		res = "RADARDME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC:
		res = "VORLOC"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME:
		res = "DUALVORDME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF:
		res = "DUALADF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA:
		res = "ADFMA"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL:
		res = "SPECIAL"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF:
		res = "DUALVHF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3:
		res = "GPSRNP3"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS:
		res = "ADFILS"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME:
		res = "DUALADF_DME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV:
		res = "RADAR_RNAV"
	}
	return
}

func (enum_codeapproachequipmentadditionalbasetype_member_1 *Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ADF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF
		return
	case "DME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME
		return
	case "RADAR":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR
		return
	case "RADARDME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME
		return
	case "VORLOC":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC
		return
	case "DUALVORDME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME
		return
	case "DUALADF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF
		return
	case "ADFMA":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA
		return
	case "SPECIAL":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL
		return
	case "DUALVHF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF
		return
	case "GPSRNP3":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3
		return
	case "ADFILS":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS
		return
	case "DUALADF_DME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME
		return
	case "RADAR_RNAV":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachequipmentadditionalbasetype_member_1 *Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME
	case "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV":
		*enum_codeapproachequipmentadditionalbasetype_member_1 = Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachequipmentadditionalbasetype_member_1 *Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachequipmentadditionalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME"
	case Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV:
		res = "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV"
	}
	return
}

func (enum_codeapproachequipmentadditionalbasetype_member_1 Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADF")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DME")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADARDME")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_VORLOC")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVORDME")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFMA")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_SPECIAL")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALVHF")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_GPSRNP3")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_ADFILS")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_DUALADF_DME")
	res = append(res, "Enum_CodeApproachEquipmentAdditionalBaseType_Member_1_RADAR_RNAV")

	return
}

func (enum_codeapproachequipmentadditionalbasetype_member_1 Enum_CodeApproachEquipmentAdditionalBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ADF")
	res = append(res, "DME")
	res = append(res, "RADAR")
	res = append(res, "RADARDME")
	res = append(res, "VORLOC")
	res = append(res, "DUALVORDME")
	res = append(res, "DUALADF")
	res = append(res, "ADFMA")
	res = append(res, "SPECIAL")
	res = append(res, "DUALVHF")
	res = append(res, "GPSRNP3")
	res = append(res, "ADFILS")
	res = append(res, "DUALADF_DME")
	res = append(res, "RADAR_RNAV")

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

// Utility function for Enum_CodeApproachPrefixBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapproachprefixbasetype_member_1 Enum_CodeApproachPrefixBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapproachprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachPrefixBaseType_Member_1_HI:
		res = "HI"
	case Enum_CodeApproachPrefixBaseType_Member_1_COPTER:
		res = "COPTER"
	case Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING:
		res = "CONVERGING"
	}
	return
}

func (enum_codeapproachprefixbasetype_member_1 *Enum_CodeApproachPrefixBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HI":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_HI
		return
	case "COPTER":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_COPTER
		return
	case "CONVERGING":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapproachprefixbasetype_member_1 *Enum_CodeApproachPrefixBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApproachPrefixBaseType_Member_1_HI":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_HI
	case "Enum_CodeApproachPrefixBaseType_Member_1_COPTER":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_COPTER
	case "Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING":
		*enum_codeapproachprefixbasetype_member_1 = Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapproachprefixbasetype_member_1 *Enum_CodeApproachPrefixBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapproachprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApproachPrefixBaseType_Member_1_HI:
		res = "Enum_CodeApproachPrefixBaseType_Member_1_HI"
	case Enum_CodeApproachPrefixBaseType_Member_1_COPTER:
		res = "Enum_CodeApproachPrefixBaseType_Member_1_COPTER"
	case Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING:
		res = "Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING"
	}
	return
}

func (enum_codeapproachprefixbasetype_member_1 Enum_CodeApproachPrefixBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApproachPrefixBaseType_Member_1_HI")
	res = append(res, "Enum_CodeApproachPrefixBaseType_Member_1_COPTER")
	res = append(res, "Enum_CodeApproachPrefixBaseType_Member_1_CONVERGING")

	return
}

func (enum_codeapproachprefixbasetype_member_1 Enum_CodeApproachPrefixBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HI")
	res = append(res, "COPTER")
	res = append(res, "CONVERGING")

	return
}

// Utility function for Enum_CodeApprovalBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeapprovalbasetype_member_1 Enum_CodeApprovalBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeapprovalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApprovalBaseType_Member_1_APPROVED:
		res = "APPROVED"
	case Enum_CodeApprovalBaseType_Member_1_DISAPPROVED:
		res = "DISAPPROVED"
	case Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED:
		res = "NOT_AUTHORISED"
	case Enum_CodeApprovalBaseType_Member_1_RESTRICTED:
		res = "RESTRICTED"
	}
	return
}

func (enum_codeapprovalbasetype_member_1 *Enum_CodeApprovalBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "APPROVED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_APPROVED
		return
	case "DISAPPROVED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_DISAPPROVED
		return
	case "NOT_AUTHORISED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED
		return
	case "RESTRICTED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_RESTRICTED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeapprovalbasetype_member_1 *Enum_CodeApprovalBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeApprovalBaseType_Member_1_APPROVED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_APPROVED
	case "Enum_CodeApprovalBaseType_Member_1_DISAPPROVED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_DISAPPROVED
	case "Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED
	case "Enum_CodeApprovalBaseType_Member_1_RESTRICTED":
		*enum_codeapprovalbasetype_member_1 = Enum_CodeApprovalBaseType_Member_1_RESTRICTED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeapprovalbasetype_member_1 *Enum_CodeApprovalBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeapprovalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeApprovalBaseType_Member_1_APPROVED:
		res = "Enum_CodeApprovalBaseType_Member_1_APPROVED"
	case Enum_CodeApprovalBaseType_Member_1_DISAPPROVED:
		res = "Enum_CodeApprovalBaseType_Member_1_DISAPPROVED"
	case Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED:
		res = "Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED"
	case Enum_CodeApprovalBaseType_Member_1_RESTRICTED:
		res = "Enum_CodeApprovalBaseType_Member_1_RESTRICTED"
	}
	return
}

func (enum_codeapprovalbasetype_member_1 Enum_CodeApprovalBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeApprovalBaseType_Member_1_APPROVED")
	res = append(res, "Enum_CodeApprovalBaseType_Member_1_DISAPPROVED")
	res = append(res, "Enum_CodeApprovalBaseType_Member_1_NOT_AUTHORISED")
	res = append(res, "Enum_CodeApprovalBaseType_Member_1_RESTRICTED")

	return
}

func (enum_codeapprovalbasetype_member_1 Enum_CodeApprovalBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "APPROVED")
	res = append(res, "DISAPPROVED")
	res = append(res, "NOT_AUTHORISED")
	res = append(res, "RESTRICTED")

	return
}

// Utility function for Enum_CodeArcDirectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codearcdirectionbasetype_member_1 Enum_CodeArcDirectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codearcdirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArcDirectionBaseType_Member_1_CWA:
		res = "CWA"
	case Enum_CodeArcDirectionBaseType_Member_1_CCA:
		res = "CCA"
	}
	return
}

func (enum_codearcdirectionbasetype_member_1 *Enum_CodeArcDirectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CWA":
		*enum_codearcdirectionbasetype_member_1 = Enum_CodeArcDirectionBaseType_Member_1_CWA
		return
	case "CCA":
		*enum_codearcdirectionbasetype_member_1 = Enum_CodeArcDirectionBaseType_Member_1_CCA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codearcdirectionbasetype_member_1 *Enum_CodeArcDirectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeArcDirectionBaseType_Member_1_CWA":
		*enum_codearcdirectionbasetype_member_1 = Enum_CodeArcDirectionBaseType_Member_1_CWA
	case "Enum_CodeArcDirectionBaseType_Member_1_CCA":
		*enum_codearcdirectionbasetype_member_1 = Enum_CodeArcDirectionBaseType_Member_1_CCA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codearcdirectionbasetype_member_1 *Enum_CodeArcDirectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codearcdirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeArcDirectionBaseType_Member_1_CWA:
		res = "Enum_CodeArcDirectionBaseType_Member_1_CWA"
	case Enum_CodeArcDirectionBaseType_Member_1_CCA:
		res = "Enum_CodeArcDirectionBaseType_Member_1_CCA"
	}
	return
}

func (enum_codearcdirectionbasetype_member_1 Enum_CodeArcDirectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeArcDirectionBaseType_Member_1_CWA")
	res = append(res, "Enum_CodeArcDirectionBaseType_Member_1_CCA")

	return
}

func (enum_codearcdirectionbasetype_member_1 Enum_CodeArcDirectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CWA")
	res = append(res, "CCA")

	return
}

// Utility function for Enum_CodeBearingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codebearingbasetype_member_1 Enum_CodeBearingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codebearingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeBearingBaseType_Member_1_TRUE_XSD:
		res = "TRUE"
	case Enum_CodeBearingBaseType_Member_1_MAG:
		res = "MAG"
	case Enum_CodeBearingBaseType_Member_1_RDL:
		res = "RDL"
	case Enum_CodeBearingBaseType_Member_1_TRK:
		res = "TRK"
	case Enum_CodeBearingBaseType_Member_1_HDG:
		res = "HDG"
	}
	return
}

func (enum_codebearingbasetype_member_1 *Enum_CodeBearingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TRUE":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_TRUE_XSD
		return
	case "MAG":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_MAG
		return
	case "RDL":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_RDL
		return
	case "TRK":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_TRK
		return
	case "HDG":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_HDG
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codebearingbasetype_member_1 *Enum_CodeBearingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeBearingBaseType_Member_1_TRUE_XSD":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_TRUE_XSD
	case "Enum_CodeBearingBaseType_Member_1_MAG":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_MAG
	case "Enum_CodeBearingBaseType_Member_1_RDL":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_RDL
	case "Enum_CodeBearingBaseType_Member_1_TRK":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_TRK
	case "Enum_CodeBearingBaseType_Member_1_HDG":
		*enum_codebearingbasetype_member_1 = Enum_CodeBearingBaseType_Member_1_HDG
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codebearingbasetype_member_1 *Enum_CodeBearingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codebearingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeBearingBaseType_Member_1_TRUE_XSD:
		res = "Enum_CodeBearingBaseType_Member_1_TRUE_XSD"
	case Enum_CodeBearingBaseType_Member_1_MAG:
		res = "Enum_CodeBearingBaseType_Member_1_MAG"
	case Enum_CodeBearingBaseType_Member_1_RDL:
		res = "Enum_CodeBearingBaseType_Member_1_RDL"
	case Enum_CodeBearingBaseType_Member_1_TRK:
		res = "Enum_CodeBearingBaseType_Member_1_TRK"
	case Enum_CodeBearingBaseType_Member_1_HDG:
		res = "Enum_CodeBearingBaseType_Member_1_HDG"
	}
	return
}

func (enum_codebearingbasetype_member_1 Enum_CodeBearingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeBearingBaseType_Member_1_TRUE_XSD")
	res = append(res, "Enum_CodeBearingBaseType_Member_1_MAG")
	res = append(res, "Enum_CodeBearingBaseType_Member_1_RDL")
	res = append(res, "Enum_CodeBearingBaseType_Member_1_TRK")
	res = append(res, "Enum_CodeBearingBaseType_Member_1_HDG")

	return
}

func (enum_codebearingbasetype_member_1 Enum_CodeBearingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TRUE")
	res = append(res, "MAG")
	res = append(res, "RDL")
	res = append(res, "TRK")
	res = append(res, "HDG")

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

// Utility function for Enum_CodeCourseBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecoursebasetype_member_1 Enum_CodeCourseBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecoursebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCourseBaseType_Member_1_TRUE_TRACK:
		res = "TRUE_TRACK"
	case Enum_CodeCourseBaseType_Member_1_MAG_TRACK:
		res = "MAG_TRACK"
	case Enum_CodeCourseBaseType_Member_1_TRUE_BRG:
		res = "TRUE_BRG"
	case Enum_CodeCourseBaseType_Member_1_MAG_BRG:
		res = "MAG_BRG"
	case Enum_CodeCourseBaseType_Member_1_HDG:
		res = "HDG"
	case Enum_CodeCourseBaseType_Member_1_RDL:
		res = "RDL"
	}
	return
}

func (enum_codecoursebasetype_member_1 *Enum_CodeCourseBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TRUE_TRACK":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_TRUE_TRACK
		return
	case "MAG_TRACK":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_MAG_TRACK
		return
	case "TRUE_BRG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_TRUE_BRG
		return
	case "MAG_BRG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_MAG_BRG
		return
	case "HDG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_HDG
		return
	case "RDL":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_RDL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecoursebasetype_member_1 *Enum_CodeCourseBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCourseBaseType_Member_1_TRUE_TRACK":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_TRUE_TRACK
	case "Enum_CodeCourseBaseType_Member_1_MAG_TRACK":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_MAG_TRACK
	case "Enum_CodeCourseBaseType_Member_1_TRUE_BRG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_TRUE_BRG
	case "Enum_CodeCourseBaseType_Member_1_MAG_BRG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_MAG_BRG
	case "Enum_CodeCourseBaseType_Member_1_HDG":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_HDG
	case "Enum_CodeCourseBaseType_Member_1_RDL":
		*enum_codecoursebasetype_member_1 = Enum_CodeCourseBaseType_Member_1_RDL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecoursebasetype_member_1 *Enum_CodeCourseBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecoursebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCourseBaseType_Member_1_TRUE_TRACK:
		res = "Enum_CodeCourseBaseType_Member_1_TRUE_TRACK"
	case Enum_CodeCourseBaseType_Member_1_MAG_TRACK:
		res = "Enum_CodeCourseBaseType_Member_1_MAG_TRACK"
	case Enum_CodeCourseBaseType_Member_1_TRUE_BRG:
		res = "Enum_CodeCourseBaseType_Member_1_TRUE_BRG"
	case Enum_CodeCourseBaseType_Member_1_MAG_BRG:
		res = "Enum_CodeCourseBaseType_Member_1_MAG_BRG"
	case Enum_CodeCourseBaseType_Member_1_HDG:
		res = "Enum_CodeCourseBaseType_Member_1_HDG"
	case Enum_CodeCourseBaseType_Member_1_RDL:
		res = "Enum_CodeCourseBaseType_Member_1_RDL"
	}
	return
}

func (enum_codecoursebasetype_member_1 Enum_CodeCourseBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCourseBaseType_Member_1_TRUE_TRACK")
	res = append(res, "Enum_CodeCourseBaseType_Member_1_MAG_TRACK")
	res = append(res, "Enum_CodeCourseBaseType_Member_1_TRUE_BRG")
	res = append(res, "Enum_CodeCourseBaseType_Member_1_MAG_BRG")
	res = append(res, "Enum_CodeCourseBaseType_Member_1_HDG")
	res = append(res, "Enum_CodeCourseBaseType_Member_1_RDL")

	return
}

func (enum_codecoursebasetype_member_1 Enum_CodeCourseBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TRUE_TRACK")
	res = append(res, "MAG_TRACK")
	res = append(res, "TRUE_BRG")
	res = append(res, "MAG_BRG")
	res = append(res, "HDG")
	res = append(res, "RDL")

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

// Utility function for Enum_CodeDesignStandardBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedesignstandardbasetype_member_1 Enum_CodeDesignStandardBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedesignstandardbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS:
		res = "PANS_OPS"
	case Enum_CodeDesignStandardBaseType_Member_1_TERPS:
		res = "TERPS"
	case Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS:
		res = "CANADA_TERPS"
	case Enum_CodeDesignStandardBaseType_Member_1_NATO:
		res = "NATO"
	}
	return
}

func (enum_codedesignstandardbasetype_member_1 *Enum_CodeDesignStandardBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PANS_OPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS
		return
	case "TERPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_TERPS
		return
	case "CANADA_TERPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS
		return
	case "NATO":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_NATO
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedesignstandardbasetype_member_1 *Enum_CodeDesignStandardBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS
	case "Enum_CodeDesignStandardBaseType_Member_1_TERPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_TERPS
	case "Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS
	case "Enum_CodeDesignStandardBaseType_Member_1_NATO":
		*enum_codedesignstandardbasetype_member_1 = Enum_CodeDesignStandardBaseType_Member_1_NATO
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedesignstandardbasetype_member_1 *Enum_CodeDesignStandardBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedesignstandardbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS:
		res = "Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS"
	case Enum_CodeDesignStandardBaseType_Member_1_TERPS:
		res = "Enum_CodeDesignStandardBaseType_Member_1_TERPS"
	case Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS:
		res = "Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS"
	case Enum_CodeDesignStandardBaseType_Member_1_NATO:
		res = "Enum_CodeDesignStandardBaseType_Member_1_NATO"
	}
	return
}

func (enum_codedesignstandardbasetype_member_1 Enum_CodeDesignStandardBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDesignStandardBaseType_Member_1_PANS_OPS")
	res = append(res, "Enum_CodeDesignStandardBaseType_Member_1_TERPS")
	res = append(res, "Enum_CodeDesignStandardBaseType_Member_1_CANADA_TERPS")
	res = append(res, "Enum_CodeDesignStandardBaseType_Member_1_NATO")

	return
}

func (enum_codedesignstandardbasetype_member_1 Enum_CodeDesignStandardBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PANS_OPS")
	res = append(res, "TERPS")
	res = append(res, "CANADA_TERPS")
	res = append(res, "NATO")

	return
}

// Utility function for Enum_CodeDirectionReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedirectionreferencebasetype_member_1 Enum_CodeDirectionReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedirectionreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionReferenceBaseType_Member_1_TO:
		res = "TO"
	case Enum_CodeDirectionReferenceBaseType_Member_1_FROM:
		res = "FROM"
	}
	return
}

func (enum_codedirectionreferencebasetype_member_1 *Enum_CodeDirectionReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TO":
		*enum_codedirectionreferencebasetype_member_1 = Enum_CodeDirectionReferenceBaseType_Member_1_TO
		return
	case "FROM":
		*enum_codedirectionreferencebasetype_member_1 = Enum_CodeDirectionReferenceBaseType_Member_1_FROM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedirectionreferencebasetype_member_1 *Enum_CodeDirectionReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDirectionReferenceBaseType_Member_1_TO":
		*enum_codedirectionreferencebasetype_member_1 = Enum_CodeDirectionReferenceBaseType_Member_1_TO
	case "Enum_CodeDirectionReferenceBaseType_Member_1_FROM":
		*enum_codedirectionreferencebasetype_member_1 = Enum_CodeDirectionReferenceBaseType_Member_1_FROM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedirectionreferencebasetype_member_1 *Enum_CodeDirectionReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedirectionreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDirectionReferenceBaseType_Member_1_TO:
		res = "Enum_CodeDirectionReferenceBaseType_Member_1_TO"
	case Enum_CodeDirectionReferenceBaseType_Member_1_FROM:
		res = "Enum_CodeDirectionReferenceBaseType_Member_1_FROM"
	}
	return
}

func (enum_codedirectionreferencebasetype_member_1 Enum_CodeDirectionReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDirectionReferenceBaseType_Member_1_TO")
	res = append(res, "Enum_CodeDirectionReferenceBaseType_Member_1_FROM")

	return
}

func (enum_codedirectionreferencebasetype_member_1 Enum_CodeDirectionReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TO")
	res = append(res, "FROM")

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

// Utility function for Enum_CodeEquipmentUnavailableBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeequipmentunavailablebasetype_member_1 Enum_CodeEquipmentUnavailableBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeequipmentunavailablebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD:
		res = "STANDARD"
	case Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD:
		res = "NONSTANDARD"
	}
	return
}

func (enum_codeequipmentunavailablebasetype_member_1 *Enum_CodeEquipmentUnavailableBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STANDARD":
		*enum_codeequipmentunavailablebasetype_member_1 = Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD
		return
	case "NONSTANDARD":
		*enum_codeequipmentunavailablebasetype_member_1 = Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeequipmentunavailablebasetype_member_1 *Enum_CodeEquipmentUnavailableBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD":
		*enum_codeequipmentunavailablebasetype_member_1 = Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD
	case "Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD":
		*enum_codeequipmentunavailablebasetype_member_1 = Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeequipmentunavailablebasetype_member_1 *Enum_CodeEquipmentUnavailableBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeequipmentunavailablebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD:
		res = "Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD"
	case Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD:
		res = "Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD"
	}
	return
}

func (enum_codeequipmentunavailablebasetype_member_1 Enum_CodeEquipmentUnavailableBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeEquipmentUnavailableBaseType_Member_1_STANDARD")
	res = append(res, "Enum_CodeEquipmentUnavailableBaseType_Member_1_NONSTANDARD")

	return
}

func (enum_codeequipmentunavailablebasetype_member_1 Enum_CodeEquipmentUnavailableBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STANDARD")
	res = append(res, "NONSTANDARD")

	return
}

// Utility function for Enum_CodeFinalGuidanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefinalguidancebasetype_member_1 Enum_CodeFinalGuidanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefinalguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFinalGuidanceBaseType_Member_1_LPV:
		res = "LPV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV:
		res = "LNAV_VNAV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LNAV:
		res = "LNAV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_GLS:
		res = "GLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ASR:
		res = "ASR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ARA:
		res = "ARA"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ARSR:
		res = "ARSR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_PAR:
		res = "PAR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS:
		res = "ILS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME:
		res = "ILS_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM:
		res = "ILS_PRM"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LDA:
		res = "LDA"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME:
		res = "LDA_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC:
		res = "LOC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC:
		res = "LOC_BC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME:
		res = "LOC_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC:
		res = "LOC_DME_BC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_MLS:
		res = "MLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME:
		res = "MLS_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_NDB:
		res = "NDB"
	case Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME:
		res = "NDB_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_SDF:
		res = "SDF"
	case Enum_CodeFinalGuidanceBaseType_Member_1_TLS:
		res = "TLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VOR:
		res = "VOR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME:
		res = "VOR_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_TACAN:
		res = "TACAN"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC:
		res = "VORTAC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LP:
		res = "LP"
	}
	return
}

func (enum_codefinalguidancebasetype_member_1 *Enum_CodeFinalGuidanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LPV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LPV
		return
	case "LNAV_VNAV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV
		return
	case "LNAV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LNAV
		return
	case "GLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_GLS
		return
	case "ASR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ASR
		return
	case "ARA":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ARA
		return
	case "ARSR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ARSR
		return
	case "PAR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_PAR
		return
	case "ILS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS
		return
	case "ILS_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME
		return
	case "ILS_PRM":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM
		return
	case "LDA":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LDA
		return
	case "LDA_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME
		return
	case "LOC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC
		return
	case "LOC_BC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC
		return
	case "LOC_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME
		return
	case "LOC_DME_BC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC
		return
	case "MLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_MLS
		return
	case "MLS_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME
		return
	case "NDB":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_NDB
		return
	case "NDB_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME
		return
	case "SDF":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_SDF
		return
	case "TLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_TLS
		return
	case "VOR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VOR
		return
	case "VOR_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME
		return
	case "TACAN":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_TACAN
		return
	case "VORTAC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC
		return
	case "DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_DME
		return
	case "LP":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LP
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefinalguidancebasetype_member_1 *Enum_CodeFinalGuidanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LPV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LPV
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LNAV
	case "Enum_CodeFinalGuidanceBaseType_Member_1_GLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_GLS
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ASR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ASR
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ARA":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ARA
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ARSR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ARSR
	case "Enum_CodeFinalGuidanceBaseType_Member_1_PAR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_PAR
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ILS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LDA":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LDA
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LOC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC
	case "Enum_CodeFinalGuidanceBaseType_Member_1_MLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_MLS
	case "Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_NDB":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_NDB
	case "Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_SDF":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_SDF
	case "Enum_CodeFinalGuidanceBaseType_Member_1_TLS":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_TLS
	case "Enum_CodeFinalGuidanceBaseType_Member_1_VOR":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VOR
	case "Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_TACAN":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_TACAN
	case "Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC
	case "Enum_CodeFinalGuidanceBaseType_Member_1_DME":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_DME
	case "Enum_CodeFinalGuidanceBaseType_Member_1_LP":
		*enum_codefinalguidancebasetype_member_1 = Enum_CodeFinalGuidanceBaseType_Member_1_LP
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefinalguidancebasetype_member_1 *Enum_CodeFinalGuidanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefinalguidancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFinalGuidanceBaseType_Member_1_LPV:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LPV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LNAV:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV"
	case Enum_CodeFinalGuidanceBaseType_Member_1_GLS:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_GLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ASR:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ASR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ARA:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ARA"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ARSR:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ARSR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_PAR:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_PAR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ILS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LDA:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LDA"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LOC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_MLS:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_MLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_NDB:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_NDB"
	case Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_SDF:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_SDF"
	case Enum_CodeFinalGuidanceBaseType_Member_1_TLS:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_TLS"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VOR:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_VOR"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_TACAN:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_TACAN"
	case Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC"
	case Enum_CodeFinalGuidanceBaseType_Member_1_DME:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_DME"
	case Enum_CodeFinalGuidanceBaseType_Member_1_LP:
		res = "Enum_CodeFinalGuidanceBaseType_Member_1_LP"
	}
	return
}

func (enum_codefinalguidancebasetype_member_1 Enum_CodeFinalGuidanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LPV")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV_VNAV")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LNAV")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_GLS")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ASR")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ARA")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ARSR")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_PAR")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ILS")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_ILS_PRM")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LDA")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LDA_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LOC")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_BC")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LOC_DME_BC")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_MLS")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_MLS_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_NDB")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_NDB_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_SDF")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_TLS")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_VOR")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_VOR_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_TACAN")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_VORTAC")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_DME")
	res = append(res, "Enum_CodeFinalGuidanceBaseType_Member_1_LP")

	return
}

func (enum_codefinalguidancebasetype_member_1 Enum_CodeFinalGuidanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LPV")
	res = append(res, "LNAV_VNAV")
	res = append(res, "LNAV")
	res = append(res, "GLS")
	res = append(res, "ASR")
	res = append(res, "ARA")
	res = append(res, "ARSR")
	res = append(res, "PAR")
	res = append(res, "ILS")
	res = append(res, "ILS_DME")
	res = append(res, "ILS_PRM")
	res = append(res, "LDA")
	res = append(res, "LDA_DME")
	res = append(res, "LOC")
	res = append(res, "LOC_BC")
	res = append(res, "LOC_DME")
	res = append(res, "LOC_DME_BC")
	res = append(res, "MLS")
	res = append(res, "MLS_DME")
	res = append(res, "NDB")
	res = append(res, "NDB_DME")
	res = append(res, "SDF")
	res = append(res, "TLS")
	res = append(res, "VOR")
	res = append(res, "VOR_DME")
	res = append(res, "TACAN")
	res = append(res, "VORTAC")
	res = append(res, "DME")
	res = append(res, "LP")

	return
}

// Utility function for Enum_CodeFreeFlightBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefreeflightbasetype_member_1 Enum_CodeFreeFlightBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefreeflightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFreeFlightBaseType_Member_1_PITCH:
		res = "PITCH"
	case Enum_CodeFreeFlightBaseType_Member_1_CATCH:
		res = "CATCH"
	}
	return
}

func (enum_codefreeflightbasetype_member_1 *Enum_CodeFreeFlightBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PITCH":
		*enum_codefreeflightbasetype_member_1 = Enum_CodeFreeFlightBaseType_Member_1_PITCH
		return
	case "CATCH":
		*enum_codefreeflightbasetype_member_1 = Enum_CodeFreeFlightBaseType_Member_1_CATCH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefreeflightbasetype_member_1 *Enum_CodeFreeFlightBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFreeFlightBaseType_Member_1_PITCH":
		*enum_codefreeflightbasetype_member_1 = Enum_CodeFreeFlightBaseType_Member_1_PITCH
	case "Enum_CodeFreeFlightBaseType_Member_1_CATCH":
		*enum_codefreeflightbasetype_member_1 = Enum_CodeFreeFlightBaseType_Member_1_CATCH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefreeflightbasetype_member_1 *Enum_CodeFreeFlightBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefreeflightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFreeFlightBaseType_Member_1_PITCH:
		res = "Enum_CodeFreeFlightBaseType_Member_1_PITCH"
	case Enum_CodeFreeFlightBaseType_Member_1_CATCH:
		res = "Enum_CodeFreeFlightBaseType_Member_1_CATCH"
	}
	return
}

func (enum_codefreeflightbasetype_member_1 Enum_CodeFreeFlightBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFreeFlightBaseType_Member_1_PITCH")
	res = append(res, "Enum_CodeFreeFlightBaseType_Member_1_CATCH")

	return
}

func (enum_codefreeflightbasetype_member_1 Enum_CodeFreeFlightBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PITCH")
	res = append(res, "CATCH")

	return
}

// Utility function for Enum_CodeHeightReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeheightreferencebasetype_member_1 Enum_CodeHeightReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeheightreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHeightReferenceBaseType_Member_1_HAT:
		res = "HAT"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAA:
		res = "HAA"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAL:
		res = "HAL"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAS:
		res = "HAS"
	}
	return
}

func (enum_codeheightreferencebasetype_member_1 *Enum_CodeHeightReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HAT":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAT
		return
	case "HAA":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAA
		return
	case "HAL":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAL
		return
	case "HAS":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeheightreferencebasetype_member_1 *Enum_CodeHeightReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeHeightReferenceBaseType_Member_1_HAT":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAT
	case "Enum_CodeHeightReferenceBaseType_Member_1_HAA":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAA
	case "Enum_CodeHeightReferenceBaseType_Member_1_HAL":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAL
	case "Enum_CodeHeightReferenceBaseType_Member_1_HAS":
		*enum_codeheightreferencebasetype_member_1 = Enum_CodeHeightReferenceBaseType_Member_1_HAS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeheightreferencebasetype_member_1 *Enum_CodeHeightReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeheightreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHeightReferenceBaseType_Member_1_HAT:
		res = "Enum_CodeHeightReferenceBaseType_Member_1_HAT"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAA:
		res = "Enum_CodeHeightReferenceBaseType_Member_1_HAA"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAL:
		res = "Enum_CodeHeightReferenceBaseType_Member_1_HAL"
	case Enum_CodeHeightReferenceBaseType_Member_1_HAS:
		res = "Enum_CodeHeightReferenceBaseType_Member_1_HAS"
	}
	return
}

func (enum_codeheightreferencebasetype_member_1 Enum_CodeHeightReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeHeightReferenceBaseType_Member_1_HAT")
	res = append(res, "Enum_CodeHeightReferenceBaseType_Member_1_HAA")
	res = append(res, "Enum_CodeHeightReferenceBaseType_Member_1_HAL")
	res = append(res, "Enum_CodeHeightReferenceBaseType_Member_1_HAS")

	return
}

func (enum_codeheightreferencebasetype_member_1 Enum_CodeHeightReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HAT")
	res = append(res, "HAA")
	res = append(res, "HAL")
	res = append(res, "HAS")

	return
}

// Utility function for Enum_CodeHoldingUsageBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeholdingusagebasetype_member_1 Enum_CodeHoldingUsageBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeholdingusagebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingUsageBaseType_Member_1_ENR:
		res = "ENR"
	case Enum_CodeHoldingUsageBaseType_Member_1_TER:
		res = "TER"
	}
	return
}

func (enum_codeholdingusagebasetype_member_1 *Enum_CodeHoldingUsageBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENR":
		*enum_codeholdingusagebasetype_member_1 = Enum_CodeHoldingUsageBaseType_Member_1_ENR
		return
	case "TER":
		*enum_codeholdingusagebasetype_member_1 = Enum_CodeHoldingUsageBaseType_Member_1_TER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeholdingusagebasetype_member_1 *Enum_CodeHoldingUsageBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeHoldingUsageBaseType_Member_1_ENR":
		*enum_codeholdingusagebasetype_member_1 = Enum_CodeHoldingUsageBaseType_Member_1_ENR
	case "Enum_CodeHoldingUsageBaseType_Member_1_TER":
		*enum_codeholdingusagebasetype_member_1 = Enum_CodeHoldingUsageBaseType_Member_1_TER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeholdingusagebasetype_member_1 *Enum_CodeHoldingUsageBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeholdingusagebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingUsageBaseType_Member_1_ENR:
		res = "Enum_CodeHoldingUsageBaseType_Member_1_ENR"
	case Enum_CodeHoldingUsageBaseType_Member_1_TER:
		res = "Enum_CodeHoldingUsageBaseType_Member_1_TER"
	}
	return
}

func (enum_codeholdingusagebasetype_member_1 Enum_CodeHoldingUsageBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeHoldingUsageBaseType_Member_1_ENR")
	res = append(res, "Enum_CodeHoldingUsageBaseType_Member_1_TER")

	return
}

func (enum_codeholdingusagebasetype_member_1 Enum_CodeHoldingUsageBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENR")
	res = append(res, "TER")

	return
}

// Utility function for Enum_CodeHoldingUseBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeholdingusebasetype_member_1 Enum_CodeHoldingUseBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeholdingusebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingUseBaseType_Member_1_PT:
		res = "PT"
	case Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL:
		res = "ARRIVAL"
	case Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH:
		res = "MISSED_APPROACH"
	case Enum_CodeHoldingUseBaseType_Member_1_CLIMB:
		res = "CLIMB"
	case Enum_CodeHoldingUseBaseType_Member_1_ATC:
		res = "ATC"
	}
	return
}

func (enum_codeholdingusebasetype_member_1 *Enum_CodeHoldingUseBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PT":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_PT
		return
	case "ARRIVAL":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL
		return
	case "MISSED_APPROACH":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH
		return
	case "CLIMB":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_CLIMB
		return
	case "ATC":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_ATC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeholdingusebasetype_member_1 *Enum_CodeHoldingUseBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeHoldingUseBaseType_Member_1_PT":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_PT
	case "Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL
	case "Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH
	case "Enum_CodeHoldingUseBaseType_Member_1_CLIMB":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_CLIMB
	case "Enum_CodeHoldingUseBaseType_Member_1_ATC":
		*enum_codeholdingusebasetype_member_1 = Enum_CodeHoldingUseBaseType_Member_1_ATC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeholdingusebasetype_member_1 *Enum_CodeHoldingUseBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeholdingusebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeHoldingUseBaseType_Member_1_PT:
		res = "Enum_CodeHoldingUseBaseType_Member_1_PT"
	case Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL:
		res = "Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL"
	case Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH:
		res = "Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH"
	case Enum_CodeHoldingUseBaseType_Member_1_CLIMB:
		res = "Enum_CodeHoldingUseBaseType_Member_1_CLIMB"
	case Enum_CodeHoldingUseBaseType_Member_1_ATC:
		res = "Enum_CodeHoldingUseBaseType_Member_1_ATC"
	}
	return
}

func (enum_codeholdingusebasetype_member_1 Enum_CodeHoldingUseBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeHoldingUseBaseType_Member_1_PT")
	res = append(res, "Enum_CodeHoldingUseBaseType_Member_1_ARRIVAL")
	res = append(res, "Enum_CodeHoldingUseBaseType_Member_1_MISSED_APPROACH")
	res = append(res, "Enum_CodeHoldingUseBaseType_Member_1_CLIMB")
	res = append(res, "Enum_CodeHoldingUseBaseType_Member_1_ATC")

	return
}

func (enum_codeholdingusebasetype_member_1 Enum_CodeHoldingUseBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PT")
	res = append(res, "ARRIVAL")
	res = append(res, "MISSED_APPROACH")
	res = append(res, "CLIMB")
	res = append(res, "ATC")

	return
}

// Utility function for Enum_CodeMilitaryRoutePointBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemilitaryroutepointbasetype_member_1 Enum_CodeMilitaryRoutePointBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemilitaryroutepointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_S:
		res = "S"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_T:
		res = "T"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_X:
		res = "X"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_AS:
		res = "AS"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_AX:
		res = "AX"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX:
		res = "ASX"
	}
	return
}

func (enum_codemilitaryroutepointbasetype_member_1 *Enum_CodeMilitaryRoutePointBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "S":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_S
		return
	case "T":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_T
		return
	case "X":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_X
		return
	case "AS":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_AS
		return
	case "AX":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_AX
		return
	case "ASX":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemilitaryroutepointbasetype_member_1 *Enum_CodeMilitaryRoutePointBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_S":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_S
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_T":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_T
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_X":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_X
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_AS":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_AS
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_AX":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_AX
	case "Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX":
		*enum_codemilitaryroutepointbasetype_member_1 = Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemilitaryroutepointbasetype_member_1 *Enum_CodeMilitaryRoutePointBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemilitaryroutepointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_S:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_S"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_T:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_T"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_X:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_X"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_AS:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_AS"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_AX:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_AX"
	case Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX:
		res = "Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX"
	}
	return
}

func (enum_codemilitaryroutepointbasetype_member_1 Enum_CodeMilitaryRoutePointBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_S")
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_T")
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_X")
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_AS")
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_AX")
	res = append(res, "Enum_CodeMilitaryRoutePointBaseType_Member_1_ASX")

	return
}

func (enum_codemilitaryroutepointbasetype_member_1 Enum_CodeMilitaryRoutePointBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "S")
	res = append(res, "T")
	res = append(res, "X")
	res = append(res, "AS")
	res = append(res, "AX")
	res = append(res, "ASX")

	return
}

// Utility function for Enum_CodeMinimaFinalApproachPathBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeminimafinalapproachpathbasetype_member_1 Enum_CodeMinimaFinalApproachPathBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeminimafinalapproachpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN:
		res = "STRAIGHT_IN"
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING:
		res = "CIRCLING"
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP:
		res = "SIDESTEP"
	}
	return
}

func (enum_codeminimafinalapproachpathbasetype_member_1 *Enum_CodeMinimaFinalApproachPathBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STRAIGHT_IN":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN
		return
	case "CIRCLING":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING
		return
	case "SIDESTEP":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeminimafinalapproachpathbasetype_member_1 *Enum_CodeMinimaFinalApproachPathBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN
	case "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING
	case "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP":
		*enum_codeminimafinalapproachpathbasetype_member_1 = Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeminimafinalapproachpathbasetype_member_1 *Enum_CodeMinimaFinalApproachPathBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeminimafinalapproachpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN:
		res = "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN"
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING:
		res = "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING"
	case Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP:
		res = "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP"
	}
	return
}

func (enum_codeminimafinalapproachpathbasetype_member_1 Enum_CodeMinimaFinalApproachPathBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_STRAIGHT_IN")
	res = append(res, "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_CIRCLING")
	res = append(res, "Enum_CodeMinimaFinalApproachPathBaseType_Member_1_SIDESTEP")

	return
}

func (enum_codeminimafinalapproachpathbasetype_member_1 Enum_CodeMinimaFinalApproachPathBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STRAIGHT_IN")
	res = append(res, "CIRCLING")
	res = append(res, "SIDESTEP")

	return
}

// Utility function for Enum_CodeMinimumAltitudeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeminimumaltitudebasetype_member_1 Enum_CodeMinimumAltitudeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeminimumaltitudebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimumAltitudeBaseType_Member_1_OCA:
		res = "OCA"
	case Enum_CodeMinimumAltitudeBaseType_Member_1_DA:
		res = "DA"
	case Enum_CodeMinimumAltitudeBaseType_Member_1_MDA:
		res = "MDA"
	}
	return
}

func (enum_codeminimumaltitudebasetype_member_1 *Enum_CodeMinimumAltitudeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OCA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_OCA
		return
	case "DA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_DA
		return
	case "MDA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_MDA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeminimumaltitudebasetype_member_1 *Enum_CodeMinimumAltitudeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMinimumAltitudeBaseType_Member_1_OCA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_OCA
	case "Enum_CodeMinimumAltitudeBaseType_Member_1_DA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_DA
	case "Enum_CodeMinimumAltitudeBaseType_Member_1_MDA":
		*enum_codeminimumaltitudebasetype_member_1 = Enum_CodeMinimumAltitudeBaseType_Member_1_MDA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeminimumaltitudebasetype_member_1 *Enum_CodeMinimumAltitudeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeminimumaltitudebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimumAltitudeBaseType_Member_1_OCA:
		res = "Enum_CodeMinimumAltitudeBaseType_Member_1_OCA"
	case Enum_CodeMinimumAltitudeBaseType_Member_1_DA:
		res = "Enum_CodeMinimumAltitudeBaseType_Member_1_DA"
	case Enum_CodeMinimumAltitudeBaseType_Member_1_MDA:
		res = "Enum_CodeMinimumAltitudeBaseType_Member_1_MDA"
	}
	return
}

func (enum_codeminimumaltitudebasetype_member_1 Enum_CodeMinimumAltitudeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMinimumAltitudeBaseType_Member_1_OCA")
	res = append(res, "Enum_CodeMinimumAltitudeBaseType_Member_1_DA")
	res = append(res, "Enum_CodeMinimumAltitudeBaseType_Member_1_MDA")

	return
}

func (enum_codeminimumaltitudebasetype_member_1 Enum_CodeMinimumAltitudeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OCA")
	res = append(res, "DA")
	res = append(res, "MDA")

	return
}

// Utility function for Enum_CodeMinimumHeightBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeminimumheightbasetype_member_1 Enum_CodeMinimumHeightBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeminimumheightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimumHeightBaseType_Member_1_DH:
		res = "DH"
	case Enum_CodeMinimumHeightBaseType_Member_1_OCH:
		res = "OCH"
	case Enum_CodeMinimumHeightBaseType_Member_1_MDH:
		res = "MDH"
	}
	return
}

func (enum_codeminimumheightbasetype_member_1 *Enum_CodeMinimumHeightBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_DH
		return
	case "OCH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_OCH
		return
	case "MDH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_MDH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeminimumheightbasetype_member_1 *Enum_CodeMinimumHeightBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMinimumHeightBaseType_Member_1_DH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_DH
	case "Enum_CodeMinimumHeightBaseType_Member_1_OCH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_OCH
	case "Enum_CodeMinimumHeightBaseType_Member_1_MDH":
		*enum_codeminimumheightbasetype_member_1 = Enum_CodeMinimumHeightBaseType_Member_1_MDH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeminimumheightbasetype_member_1 *Enum_CodeMinimumHeightBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeminimumheightbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMinimumHeightBaseType_Member_1_DH:
		res = "Enum_CodeMinimumHeightBaseType_Member_1_DH"
	case Enum_CodeMinimumHeightBaseType_Member_1_OCH:
		res = "Enum_CodeMinimumHeightBaseType_Member_1_OCH"
	case Enum_CodeMinimumHeightBaseType_Member_1_MDH:
		res = "Enum_CodeMinimumHeightBaseType_Member_1_MDH"
	}
	return
}

func (enum_codeminimumheightbasetype_member_1 Enum_CodeMinimumHeightBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMinimumHeightBaseType_Member_1_DH")
	res = append(res, "Enum_CodeMinimumHeightBaseType_Member_1_OCH")
	res = append(res, "Enum_CodeMinimumHeightBaseType_Member_1_MDH")

	return
}

func (enum_codeminimumheightbasetype_member_1 Enum_CodeMinimumHeightBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DH")
	res = append(res, "OCH")
	res = append(res, "MDH")

	return
}

// Utility function for Enum_CodeMissedApproachBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemissedapproachbasetype_member_1 Enum_CodeMissedApproachBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemissedapproachbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMissedApproachBaseType_Member_1_PRIMARY:
		res = "PRIMARY"
	case Enum_CodeMissedApproachBaseType_Member_1_SECONDARY:
		res = "SECONDARY"
	case Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE:
		res = "ALTERNATE"
	case Enum_CodeMissedApproachBaseType_Member_1_TACAN:
		res = "TACAN"
	case Enum_CodeMissedApproachBaseType_Member_1_TACANALT:
		res = "TACANALT"
	case Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT:
		res = "ENGINEOUT"
	}
	return
}

func (enum_codemissedapproachbasetype_member_1 *Enum_CodeMissedApproachBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PRIMARY":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_PRIMARY
		return
	case "SECONDARY":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_SECONDARY
		return
	case "ALTERNATE":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE
		return
	case "TACAN":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_TACAN
		return
	case "TACANALT":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_TACANALT
		return
	case "ENGINEOUT":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemissedapproachbasetype_member_1 *Enum_CodeMissedApproachBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMissedApproachBaseType_Member_1_PRIMARY":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_PRIMARY
	case "Enum_CodeMissedApproachBaseType_Member_1_SECONDARY":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_SECONDARY
	case "Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE
	case "Enum_CodeMissedApproachBaseType_Member_1_TACAN":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_TACAN
	case "Enum_CodeMissedApproachBaseType_Member_1_TACANALT":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_TACANALT
	case "Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT":
		*enum_codemissedapproachbasetype_member_1 = Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemissedapproachbasetype_member_1 *Enum_CodeMissedApproachBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemissedapproachbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMissedApproachBaseType_Member_1_PRIMARY:
		res = "Enum_CodeMissedApproachBaseType_Member_1_PRIMARY"
	case Enum_CodeMissedApproachBaseType_Member_1_SECONDARY:
		res = "Enum_CodeMissedApproachBaseType_Member_1_SECONDARY"
	case Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE:
		res = "Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE"
	case Enum_CodeMissedApproachBaseType_Member_1_TACAN:
		res = "Enum_CodeMissedApproachBaseType_Member_1_TACAN"
	case Enum_CodeMissedApproachBaseType_Member_1_TACANALT:
		res = "Enum_CodeMissedApproachBaseType_Member_1_TACANALT"
	case Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT:
		res = "Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT"
	}
	return
}

func (enum_codemissedapproachbasetype_member_1 Enum_CodeMissedApproachBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_PRIMARY")
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_SECONDARY")
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_ALTERNATE")
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_TACAN")
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_TACANALT")
	res = append(res, "Enum_CodeMissedApproachBaseType_Member_1_ENGINEOUT")

	return
}

func (enum_codemissedapproachbasetype_member_1 Enum_CodeMissedApproachBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PRIMARY")
	res = append(res, "SECONDARY")
	res = append(res, "ALTERNATE")
	res = append(res, "TACAN")
	res = append(res, "TACANALT")
	res = append(res, "ENGINEOUT")

	return
}

// Utility function for Enum_CodeNavigationAreaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavigationareabasetype_member_1 Enum_CodeNavigationAreaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavigationareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationAreaBaseType_Member_1_PNA:
		res = "PNA"
	case Enum_CodeNavigationAreaBaseType_Member_1_OMNI:
		res = "OMNI"
	case Enum_CodeNavigationAreaBaseType_Member_1_DVA:
		res = "DVA"
	}
	return
}

func (enum_codenavigationareabasetype_member_1 *Enum_CodeNavigationAreaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PNA":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_PNA
		return
	case "OMNI":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_OMNI
		return
	case "DVA":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_DVA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavigationareabasetype_member_1 *Enum_CodeNavigationAreaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavigationAreaBaseType_Member_1_PNA":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_PNA
	case "Enum_CodeNavigationAreaBaseType_Member_1_OMNI":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_OMNI
	case "Enum_CodeNavigationAreaBaseType_Member_1_DVA":
		*enum_codenavigationareabasetype_member_1 = Enum_CodeNavigationAreaBaseType_Member_1_DVA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavigationareabasetype_member_1 *Enum_CodeNavigationAreaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavigationareabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationAreaBaseType_Member_1_PNA:
		res = "Enum_CodeNavigationAreaBaseType_Member_1_PNA"
	case Enum_CodeNavigationAreaBaseType_Member_1_OMNI:
		res = "Enum_CodeNavigationAreaBaseType_Member_1_OMNI"
	case Enum_CodeNavigationAreaBaseType_Member_1_DVA:
		res = "Enum_CodeNavigationAreaBaseType_Member_1_DVA"
	}
	return
}

func (enum_codenavigationareabasetype_member_1 Enum_CodeNavigationAreaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavigationAreaBaseType_Member_1_PNA")
	res = append(res, "Enum_CodeNavigationAreaBaseType_Member_1_OMNI")
	res = append(res, "Enum_CodeNavigationAreaBaseType_Member_1_DVA")

	return
}

func (enum_codenavigationareabasetype_member_1 Enum_CodeNavigationAreaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PNA")
	res = append(res, "OMNI")
	res = append(res, "DVA")

	return
}

// Utility function for Enum_CodeNavigationAreaRestrictionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavigationarearestrictionbasetype_member_1 Enum_CodeNavigationAreaRestrictionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavigationarearestrictionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR:
		res = "VECTOR"
	case Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL:
		res = "OMNIDIRECTIONAL"
	}
	return
}

func (enum_codenavigationarearestrictionbasetype_member_1 *Enum_CodeNavigationAreaRestrictionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "VECTOR":
		*enum_codenavigationarearestrictionbasetype_member_1 = Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR
		return
	case "OMNIDIRECTIONAL":
		*enum_codenavigationarearestrictionbasetype_member_1 = Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavigationarearestrictionbasetype_member_1 *Enum_CodeNavigationAreaRestrictionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR":
		*enum_codenavigationarearestrictionbasetype_member_1 = Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR
	case "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL":
		*enum_codenavigationarearestrictionbasetype_member_1 = Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavigationarearestrictionbasetype_member_1 *Enum_CodeNavigationAreaRestrictionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavigationarearestrictionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR:
		res = "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR"
	case Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL:
		res = "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL"
	}
	return
}

func (enum_codenavigationarearestrictionbasetype_member_1 Enum_CodeNavigationAreaRestrictionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_VECTOR")
	res = append(res, "Enum_CodeNavigationAreaRestrictionBaseType_Member_1_OMNIDIRECTIONAL")

	return
}

func (enum_codenavigationarearestrictionbasetype_member_1 Enum_CodeNavigationAreaRestrictionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "VECTOR")
	res = append(res, "OMNIDIRECTIONAL")

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

// Utility function for Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeobstructionidsurfacezonebasetype_member_1 Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeobstructionidsurfacezonebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH:
		res = "APPROACH"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL:
		res = "CONICAL"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL:
		res = "HORIZONTAL"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY:
		res = "PRIMARY"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION:
		res = "TRANSITION"
	}
	return
}

func (enum_codeobstructionidsurfacezonebasetype_member_1 *Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "APPROACH":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH
		return
	case "CONICAL":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL
		return
	case "HORIZONTAL":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL
		return
	case "PRIMARY":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY
		return
	case "TRANSITION":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeobstructionidsurfacezonebasetype_member_1 *Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH
	case "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL
	case "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL
	case "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY
	case "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION":
		*enum_codeobstructionidsurfacezonebasetype_member_1 = Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeobstructionidsurfacezonebasetype_member_1 *Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeobstructionidsurfacezonebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH:
		res = "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL:
		res = "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL:
		res = "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY:
		res = "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY"
	case Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION:
		res = "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION"
	}
	return
}

func (enum_codeobstructionidsurfacezonebasetype_member_1 Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_APPROACH")
	res = append(res, "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_CONICAL")
	res = append(res, "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_HORIZONTAL")
	res = append(res, "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_PRIMARY")
	res = append(res, "Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1_TRANSITION")

	return
}

func (enum_codeobstructionidsurfacezonebasetype_member_1 Enum_CodeObstructionIdSurfaceZoneBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "APPROACH")
	res = append(res, "CONICAL")
	res = append(res, "HORIZONTAL")
	res = append(res, "PRIMARY")
	res = append(res, "TRANSITION")

	return
}

// Utility function for Enum_CodeProcedureAvailabilityBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprocedureavailabilitybasetype_member_1 Enum_CodeProcedureAvailabilityBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprocedureavailabilitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE:
		res = "USABLE"
	case Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE:
		res = "UNUSABLE"
	}
	return
}

func (enum_codeprocedureavailabilitybasetype_member_1 *Enum_CodeProcedureAvailabilityBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "USABLE":
		*enum_codeprocedureavailabilitybasetype_member_1 = Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE
		return
	case "UNUSABLE":
		*enum_codeprocedureavailabilitybasetype_member_1 = Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprocedureavailabilitybasetype_member_1 *Enum_CodeProcedureAvailabilityBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE":
		*enum_codeprocedureavailabilitybasetype_member_1 = Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE
	case "Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE":
		*enum_codeprocedureavailabilitybasetype_member_1 = Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprocedureavailabilitybasetype_member_1 *Enum_CodeProcedureAvailabilityBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprocedureavailabilitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE:
		res = "Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE"
	case Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE:
		res = "Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE"
	}
	return
}

func (enum_codeprocedureavailabilitybasetype_member_1 Enum_CodeProcedureAvailabilityBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProcedureAvailabilityBaseType_Member_1_USABLE")
	res = append(res, "Enum_CodeProcedureAvailabilityBaseType_Member_1_UNUSABLE")

	return
}

func (enum_codeprocedureavailabilitybasetype_member_1 Enum_CodeProcedureAvailabilityBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "USABLE")
	res = append(res, "UNUSABLE")

	return
}

// Utility function for Enum_CodeProcedureCodingStandardBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprocedurecodingstandardbasetype_member_1 Enum_CodeProcedureCodingStandardBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprocedurecodingstandardbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS:
		res = "PANS_OPS"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15:
		res = "ARINC_424_15"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18:
		res = "ARINC_424_18"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19:
		res = "ARINC_424_19"
	}
	return
}

func (enum_codeprocedurecodingstandardbasetype_member_1 *Enum_CodeProcedureCodingStandardBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PANS_OPS":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS
		return
	case "ARINC_424_15":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15
		return
	case "ARINC_424_18":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18
		return
	case "ARINC_424_19":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprocedurecodingstandardbasetype_member_1 *Enum_CodeProcedureCodingStandardBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS
	case "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15
	case "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18
	case "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19":
		*enum_codeprocedurecodingstandardbasetype_member_1 = Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprocedurecodingstandardbasetype_member_1 *Enum_CodeProcedureCodingStandardBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprocedurecodingstandardbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS:
		res = "Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15:
		res = "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18:
		res = "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18"
	case Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19:
		res = "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19"
	}
	return
}

func (enum_codeprocedurecodingstandardbasetype_member_1 Enum_CodeProcedureCodingStandardBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProcedureCodingStandardBaseType_Member_1_PANS_OPS")
	res = append(res, "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_15")
	res = append(res, "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_18")
	res = append(res, "Enum_CodeProcedureCodingStandardBaseType_Member_1_ARINC_424_19")

	return
}

func (enum_codeprocedurecodingstandardbasetype_member_1 Enum_CodeProcedureCodingStandardBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PANS_OPS")
	res = append(res, "ARINC_424_15")
	res = append(res, "ARINC_424_18")
	res = append(res, "ARINC_424_19")

	return
}

// Utility function for Enum_CodeProcedureDistanceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeproceduredistancebasetype_member_1 Enum_CodeProcedureDistanceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeproceduredistancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureDistanceBaseType_Member_1_HAT:
		res = "HAT"
	case Enum_CodeProcedureDistanceBaseType_Member_1_OM:
		res = "OM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_MM:
		res = "MM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_IM:
		res = "IM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_PFAF:
		res = "PFAF"
	case Enum_CodeProcedureDistanceBaseType_Member_1_GSANT:
		res = "GSANT"
	case Enum_CodeProcedureDistanceBaseType_Member_1_FAF:
		res = "FAF"
	case Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD:
		res = "MAP"
	case Enum_CodeProcedureDistanceBaseType_Member_1_THLD:
		res = "THLD"
	case Enum_CodeProcedureDistanceBaseType_Member_1_VDP:
		res = "VDP"
	case Enum_CodeProcedureDistanceBaseType_Member_1_RECH:
		res = "RECH"
	}
	return
}

func (enum_codeproceduredistancebasetype_member_1 *Enum_CodeProcedureDistanceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HAT":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_HAT
		return
	case "OM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_OM
		return
	case "MM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_MM
		return
	case "IM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_IM
		return
	case "PFAF":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_PFAF
		return
	case "GSANT":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_GSANT
		return
	case "FAF":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_FAF
		return
	case "MAP":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD
		return
	case "THLD":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_THLD
		return
	case "VDP":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_VDP
		return
	case "RECH":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_RECH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeproceduredistancebasetype_member_1 *Enum_CodeProcedureDistanceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProcedureDistanceBaseType_Member_1_HAT":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_HAT
	case "Enum_CodeProcedureDistanceBaseType_Member_1_OM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_OM
	case "Enum_CodeProcedureDistanceBaseType_Member_1_MM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_MM
	case "Enum_CodeProcedureDistanceBaseType_Member_1_IM":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_IM
	case "Enum_CodeProcedureDistanceBaseType_Member_1_PFAF":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_PFAF
	case "Enum_CodeProcedureDistanceBaseType_Member_1_GSANT":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_GSANT
	case "Enum_CodeProcedureDistanceBaseType_Member_1_FAF":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_FAF
	case "Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD
	case "Enum_CodeProcedureDistanceBaseType_Member_1_THLD":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_THLD
	case "Enum_CodeProcedureDistanceBaseType_Member_1_VDP":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_VDP
	case "Enum_CodeProcedureDistanceBaseType_Member_1_RECH":
		*enum_codeproceduredistancebasetype_member_1 = Enum_CodeProcedureDistanceBaseType_Member_1_RECH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeproceduredistancebasetype_member_1 *Enum_CodeProcedureDistanceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeproceduredistancebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureDistanceBaseType_Member_1_HAT:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_HAT"
	case Enum_CodeProcedureDistanceBaseType_Member_1_OM:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_OM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_MM:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_MM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_IM:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_IM"
	case Enum_CodeProcedureDistanceBaseType_Member_1_PFAF:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_PFAF"
	case Enum_CodeProcedureDistanceBaseType_Member_1_GSANT:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_GSANT"
	case Enum_CodeProcedureDistanceBaseType_Member_1_FAF:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_FAF"
	case Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD"
	case Enum_CodeProcedureDistanceBaseType_Member_1_THLD:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_THLD"
	case Enum_CodeProcedureDistanceBaseType_Member_1_VDP:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_VDP"
	case Enum_CodeProcedureDistanceBaseType_Member_1_RECH:
		res = "Enum_CodeProcedureDistanceBaseType_Member_1_RECH"
	}
	return
}

func (enum_codeproceduredistancebasetype_member_1 Enum_CodeProcedureDistanceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_HAT")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_OM")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_MM")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_IM")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_PFAF")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_GSANT")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_FAF")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_MAP_XSD")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_THLD")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_VDP")
	res = append(res, "Enum_CodeProcedureDistanceBaseType_Member_1_RECH")

	return
}

func (enum_codeproceduredistancebasetype_member_1 Enum_CodeProcedureDistanceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HAT")
	res = append(res, "OM")
	res = append(res, "MM")
	res = append(res, "IM")
	res = append(res, "PFAF")
	res = append(res, "GSANT")
	res = append(res, "FAF")
	res = append(res, "MAP")
	res = append(res, "THLD")
	res = append(res, "VDP")
	res = append(res, "RECH")

	return
}

// Utility function for Enum_CodeProcedureFixRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprocedurefixrolebasetype_member_1 Enum_CodeProcedureFixRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprocedurefixrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FAF:
		res = "FAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP:
		res = "FPAP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FROP:
		res = "FROP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FTP:
		res = "FTP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IAF:
		res = "IAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD:
		res = "IF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF:
		res = "IF_IAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_LTP:
		res = "LTP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF:
		res = "MAHF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT:
		res = "MAPT"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_SDF:
		res = "SDF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_TP:
		res = "TP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_VDP:
		res = "VDP"
	}
	return
}

func (enum_codeprocedurefixrolebasetype_member_1 *Enum_CodeProcedureFixRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FAF
		return
	case "FPAP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP
		return
	case "FROP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FROP
		return
	case "FTP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FTP
		return
	case "IAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IAF
		return
	case "IF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD
		return
	case "IF_IAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF
		return
	case "LTP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_LTP
		return
	case "MAHF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF
		return
	case "MAPT":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT
		return
	case "SDF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_SDF
		return
	case "TP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_TP
		return
	case "VDP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_VDP
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprocedurefixrolebasetype_member_1 *Enum_CodeProcedureFixRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_FAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FAF
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_FROP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FROP
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_FTP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_FTP
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_IAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IAF
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_LTP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_LTP
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_SDF":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_SDF
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_TP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_TP
	case "Enum_CodeProcedureFixRoleBaseType_Member_1_VDP":
		*enum_codeprocedurefixrolebasetype_member_1 = Enum_CodeProcedureFixRoleBaseType_Member_1_VDP
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprocedurefixrolebasetype_member_1 *Enum_CodeProcedureFixRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprocedurefixrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FAF:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_FAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FROP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_FROP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_FTP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_FTP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IAF:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_IAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_LTP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_LTP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_SDF:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_SDF"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_TP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_TP"
	case Enum_CodeProcedureFixRoleBaseType_Member_1_VDP:
		res = "Enum_CodeProcedureFixRoleBaseType_Member_1_VDP"
	}
	return
}

func (enum_codeprocedurefixrolebasetype_member_1 Enum_CodeProcedureFixRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_FAF")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_FPAP")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_FROP")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_FTP")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_IAF")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_XSD")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_IF_IAF")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_LTP")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_MAHF")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_MAPT")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_SDF")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_TP")
	res = append(res, "Enum_CodeProcedureFixRoleBaseType_Member_1_VDP")

	return
}

func (enum_codeprocedurefixrolebasetype_member_1 Enum_CodeProcedureFixRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FAF")
	res = append(res, "FPAP")
	res = append(res, "FROP")
	res = append(res, "FTP")
	res = append(res, "IAF")
	res = append(res, "IF")
	res = append(res, "IF_IAF")
	res = append(res, "LTP")
	res = append(res, "MAHF")
	res = append(res, "MAPT")
	res = append(res, "SDF")
	res = append(res, "TP")
	res = append(res, "VDP")

	return
}

// Utility function for Enum_CodeProcedurePhaseBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprocedurephasebasetype_member_1 Enum_CodeProcedurePhaseBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprocedurephasebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedurePhaseBaseType_Member_1_RWY:
		res = "RWY"
	case Enum_CodeProcedurePhaseBaseType_Member_1_COMMON:
		res = "COMMON"
	case Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE:
		res = "EN_ROUTE"
	case Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH:
		res = "APPROACH"
	case Enum_CodeProcedurePhaseBaseType_Member_1_FINAL:
		res = "FINAL"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED:
		res = "MISSED"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P:
		res = "MISSED_P"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S:
		res = "MISSED_S"
	case Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT:
		res = "ENGINE_OUT"
	}
	return
}

func (enum_codeprocedurephasebasetype_member_1 *Enum_CodeProcedurePhaseBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RWY":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_RWY
		return
	case "COMMON":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_COMMON
		return
	case "EN_ROUTE":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE
		return
	case "APPROACH":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH
		return
	case "FINAL":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_FINAL
		return
	case "MISSED":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED
		return
	case "MISSED_P":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P
		return
	case "MISSED_S":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S
		return
	case "ENGINE_OUT":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprocedurephasebasetype_member_1 *Enum_CodeProcedurePhaseBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeProcedurePhaseBaseType_Member_1_RWY":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_RWY
	case "Enum_CodeProcedurePhaseBaseType_Member_1_COMMON":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_COMMON
	case "Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE
	case "Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH
	case "Enum_CodeProcedurePhaseBaseType_Member_1_FINAL":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_FINAL
	case "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED
	case "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P
	case "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S
	case "Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT":
		*enum_codeprocedurephasebasetype_member_1 = Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprocedurephasebasetype_member_1 *Enum_CodeProcedurePhaseBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprocedurephasebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeProcedurePhaseBaseType_Member_1_RWY:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_RWY"
	case Enum_CodeProcedurePhaseBaseType_Member_1_COMMON:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_COMMON"
	case Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE"
	case Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH"
	case Enum_CodeProcedurePhaseBaseType_Member_1_FINAL:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_FINAL"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P"
	case Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S"
	case Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT:
		res = "Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT"
	}
	return
}

func (enum_codeprocedurephasebasetype_member_1 Enum_CodeProcedurePhaseBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_RWY")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_COMMON")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_EN_ROUTE")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_APPROACH")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_FINAL")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_P")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_MISSED_S")
	res = append(res, "Enum_CodeProcedurePhaseBaseType_Member_1_ENGINE_OUT")

	return
}

func (enum_codeprocedurephasebasetype_member_1 Enum_CodeProcedurePhaseBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RWY")
	res = append(res, "COMMON")
	res = append(res, "EN_ROUTE")
	res = append(res, "APPROACH")
	res = append(res, "FINAL")
	res = append(res, "MISSED")
	res = append(res, "MISSED_P")
	res = append(res, "MISSED_S")
	res = append(res, "ENGINE_OUT")

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

// Utility function for Enum_CodeRVSMPointRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codervsmpointrolebasetype_member_1 Enum_CodeRVSMPointRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codervsmpointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVSMPointRoleBaseType_Member_1_IN:
		res = "IN"
	case Enum_CodeRVSMPointRoleBaseType_Member_1_OUT:
		res = "OUT"
	case Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT:
		res = "IN_OUT"
	}
	return
}

func (enum_codervsmpointrolebasetype_member_1 *Enum_CodeRVSMPointRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IN":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_IN
		return
	case "OUT":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_OUT
		return
	case "IN_OUT":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codervsmpointrolebasetype_member_1 *Enum_CodeRVSMPointRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRVSMPointRoleBaseType_Member_1_IN":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_IN
	case "Enum_CodeRVSMPointRoleBaseType_Member_1_OUT":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_OUT
	case "Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT":
		*enum_codervsmpointrolebasetype_member_1 = Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codervsmpointrolebasetype_member_1 *Enum_CodeRVSMPointRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codervsmpointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRVSMPointRoleBaseType_Member_1_IN:
		res = "Enum_CodeRVSMPointRoleBaseType_Member_1_IN"
	case Enum_CodeRVSMPointRoleBaseType_Member_1_OUT:
		res = "Enum_CodeRVSMPointRoleBaseType_Member_1_OUT"
	case Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT:
		res = "Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT"
	}
	return
}

func (enum_codervsmpointrolebasetype_member_1 Enum_CodeRVSMPointRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRVSMPointRoleBaseType_Member_1_IN")
	res = append(res, "Enum_CodeRVSMPointRoleBaseType_Member_1_OUT")
	res = append(res, "Enum_CodeRVSMPointRoleBaseType_Member_1_IN_OUT")

	return
}

func (enum_codervsmpointrolebasetype_member_1 Enum_CodeRVSMPointRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IN")
	res = append(res, "OUT")
	res = append(res, "IN_OUT")

	return
}

// Utility function for Enum_CodeReferenceRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codereferencerolebasetype_member_1 Enum_CodeReferenceRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codereferencerolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION:
		res = "INTERSECTION"
	case Enum_CodeReferenceRoleBaseType_Member_1_RECNAV:
		res = "RECNAV"
	case Enum_CodeReferenceRoleBaseType_Member_1_ATD:
		res = "ATD"
	case Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME:
		res = "RAD_DME"
	}
	return
}

func (enum_codereferencerolebasetype_member_1 *Enum_CodeReferenceRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "INTERSECTION":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION
		return
	case "RECNAV":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_RECNAV
		return
	case "ATD":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_ATD
		return
	case "RAD_DME":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codereferencerolebasetype_member_1 *Enum_CodeReferenceRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION
	case "Enum_CodeReferenceRoleBaseType_Member_1_RECNAV":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_RECNAV
	case "Enum_CodeReferenceRoleBaseType_Member_1_ATD":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_ATD
	case "Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME":
		*enum_codereferencerolebasetype_member_1 = Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codereferencerolebasetype_member_1 *Enum_CodeReferenceRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codereferencerolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION:
		res = "Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION"
	case Enum_CodeReferenceRoleBaseType_Member_1_RECNAV:
		res = "Enum_CodeReferenceRoleBaseType_Member_1_RECNAV"
	case Enum_CodeReferenceRoleBaseType_Member_1_ATD:
		res = "Enum_CodeReferenceRoleBaseType_Member_1_ATD"
	case Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME:
		res = "Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME"
	}
	return
}

func (enum_codereferencerolebasetype_member_1 Enum_CodeReferenceRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeReferenceRoleBaseType_Member_1_INTERSECTION")
	res = append(res, "Enum_CodeReferenceRoleBaseType_Member_1_RECNAV")
	res = append(res, "Enum_CodeReferenceRoleBaseType_Member_1_ATD")
	res = append(res, "Enum_CodeReferenceRoleBaseType_Member_1_RAD_DME")

	return
}

func (enum_codereferencerolebasetype_member_1 Enum_CodeReferenceRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "INTERSECTION")
	res = append(res, "RECNAV")
	res = append(res, "ATD")
	res = append(res, "RAD_DME")

	return
}

// Utility function for Enum_CodeRelativePositionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderelativepositionbasetype_member_1 Enum_CodeRelativePositionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderelativepositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRelativePositionBaseType_Member_1_BEFORE:
		res = "BEFORE"
	case Enum_CodeRelativePositionBaseType_Member_1_AT:
		res = "AT"
	case Enum_CodeRelativePositionBaseType_Member_1_AFTER:
		res = "AFTER"
	}
	return
}

func (enum_coderelativepositionbasetype_member_1 *Enum_CodeRelativePositionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BEFORE":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_BEFORE
		return
	case "AT":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_AT
		return
	case "AFTER":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_AFTER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderelativepositionbasetype_member_1 *Enum_CodeRelativePositionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRelativePositionBaseType_Member_1_BEFORE":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_BEFORE
	case "Enum_CodeRelativePositionBaseType_Member_1_AT":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_AT
	case "Enum_CodeRelativePositionBaseType_Member_1_AFTER":
		*enum_coderelativepositionbasetype_member_1 = Enum_CodeRelativePositionBaseType_Member_1_AFTER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderelativepositionbasetype_member_1 *Enum_CodeRelativePositionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderelativepositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRelativePositionBaseType_Member_1_BEFORE:
		res = "Enum_CodeRelativePositionBaseType_Member_1_BEFORE"
	case Enum_CodeRelativePositionBaseType_Member_1_AT:
		res = "Enum_CodeRelativePositionBaseType_Member_1_AT"
	case Enum_CodeRelativePositionBaseType_Member_1_AFTER:
		res = "Enum_CodeRelativePositionBaseType_Member_1_AFTER"
	}
	return
}

func (enum_coderelativepositionbasetype_member_1 Enum_CodeRelativePositionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRelativePositionBaseType_Member_1_BEFORE")
	res = append(res, "Enum_CodeRelativePositionBaseType_Member_1_AT")
	res = append(res, "Enum_CodeRelativePositionBaseType_Member_1_AFTER")

	return
}

func (enum_coderelativepositionbasetype_member_1 Enum_CodeRelativePositionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BEFORE")
	res = append(res, "AT")
	res = append(res, "AFTER")

	return
}

// Utility function for Enum_CodeSafeAltitudeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesafealtitudebasetype_member_1 Enum_CodeSafeAltitudeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesafealtitudebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSafeAltitudeBaseType_Member_1_MSA:
		res = "MSA"
	case Enum_CodeSafeAltitudeBaseType_Member_1_ESA:
		res = "ESA"
	}
	return
}

func (enum_codesafealtitudebasetype_member_1 *Enum_CodeSafeAltitudeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MSA":
		*enum_codesafealtitudebasetype_member_1 = Enum_CodeSafeAltitudeBaseType_Member_1_MSA
		return
	case "ESA":
		*enum_codesafealtitudebasetype_member_1 = Enum_CodeSafeAltitudeBaseType_Member_1_ESA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesafealtitudebasetype_member_1 *Enum_CodeSafeAltitudeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSafeAltitudeBaseType_Member_1_MSA":
		*enum_codesafealtitudebasetype_member_1 = Enum_CodeSafeAltitudeBaseType_Member_1_MSA
	case "Enum_CodeSafeAltitudeBaseType_Member_1_ESA":
		*enum_codesafealtitudebasetype_member_1 = Enum_CodeSafeAltitudeBaseType_Member_1_ESA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesafealtitudebasetype_member_1 *Enum_CodeSafeAltitudeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesafealtitudebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSafeAltitudeBaseType_Member_1_MSA:
		res = "Enum_CodeSafeAltitudeBaseType_Member_1_MSA"
	case Enum_CodeSafeAltitudeBaseType_Member_1_ESA:
		res = "Enum_CodeSafeAltitudeBaseType_Member_1_ESA"
	}
	return
}

func (enum_codesafealtitudebasetype_member_1 Enum_CodeSafeAltitudeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSafeAltitudeBaseType_Member_1_MSA")
	res = append(res, "Enum_CodeSafeAltitudeBaseType_Member_1_ESA")

	return
}

func (enum_codesafealtitudebasetype_member_1 Enum_CodeSafeAltitudeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MSA")
	res = append(res, "ESA")

	return
}

// Utility function for Enum_CodeSegmentPathBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesegmentpathbasetype_member_1 Enum_CodeSegmentPathBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesegmentpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSegmentPathBaseType_Member_1_AF:
		res = "AF"
	case Enum_CodeSegmentPathBaseType_Member_1_HF:
		res = "HF"
	case Enum_CodeSegmentPathBaseType_Member_1_HA:
		res = "HA"
	case Enum_CodeSegmentPathBaseType_Member_1_HM:
		res = "HM"
	case Enum_CodeSegmentPathBaseType_Member_1_IF_XSD:
		res = "IF"
	case Enum_CodeSegmentPathBaseType_Member_1_PI:
		res = "PI"
	case Enum_CodeSegmentPathBaseType_Member_1_PT:
		res = "PT"
	case Enum_CodeSegmentPathBaseType_Member_1_TF:
		res = "TF"
	case Enum_CodeSegmentPathBaseType_Member_1_CA:
		res = "CA"
	case Enum_CodeSegmentPathBaseType_Member_1_CD:
		res = "CD"
	case Enum_CodeSegmentPathBaseType_Member_1_CI:
		res = "CI"
	case Enum_CodeSegmentPathBaseType_Member_1_CR:
		res = "CR"
	case Enum_CodeSegmentPathBaseType_Member_1_CF:
		res = "CF"
	case Enum_CodeSegmentPathBaseType_Member_1_DF:
		res = "DF"
	case Enum_CodeSegmentPathBaseType_Member_1_FA:
		res = "FA"
	case Enum_CodeSegmentPathBaseType_Member_1_FC:
		res = "FC"
	case Enum_CodeSegmentPathBaseType_Member_1_FT:
		res = "FT"
	case Enum_CodeSegmentPathBaseType_Member_1_FM:
		res = "FM"
	case Enum_CodeSegmentPathBaseType_Member_1_VM:
		res = "VM"
	case Enum_CodeSegmentPathBaseType_Member_1_FD:
		res = "FD"
	case Enum_CodeSegmentPathBaseType_Member_1_VR:
		res = "VR"
	case Enum_CodeSegmentPathBaseType_Member_1_VD:
		res = "VD"
	case Enum_CodeSegmentPathBaseType_Member_1_VI:
		res = "VI"
	case Enum_CodeSegmentPathBaseType_Member_1_VA:
		res = "VA"
	case Enum_CodeSegmentPathBaseType_Member_1_RF:
		res = "RF"
	}
	return
}

func (enum_codesegmentpathbasetype_member_1 *Enum_CodeSegmentPathBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_AF
		return
	case "HF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HF
		return
	case "HA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HA
		return
	case "HM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HM
		return
	case "IF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_IF_XSD
		return
	case "PI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_PI
		return
	case "PT":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_PT
		return
	case "TF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_TF
		return
	case "CA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CA
		return
	case "CD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CD
		return
	case "CI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CI
		return
	case "CR":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CR
		return
	case "CF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CF
		return
	case "DF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_DF
		return
	case "FA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FA
		return
	case "FC":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FC
		return
	case "FT":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FT
		return
	case "FM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FM
		return
	case "VM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VM
		return
	case "FD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FD
		return
	case "VR":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VR
		return
	case "VD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VD
		return
	case "VI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VI
		return
	case "VA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VA
		return
	case "RF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_RF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesegmentpathbasetype_member_1 *Enum_CodeSegmentPathBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSegmentPathBaseType_Member_1_AF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_AF
	case "Enum_CodeSegmentPathBaseType_Member_1_HF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HF
	case "Enum_CodeSegmentPathBaseType_Member_1_HA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HA
	case "Enum_CodeSegmentPathBaseType_Member_1_HM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_HM
	case "Enum_CodeSegmentPathBaseType_Member_1_IF_XSD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_IF_XSD
	case "Enum_CodeSegmentPathBaseType_Member_1_PI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_PI
	case "Enum_CodeSegmentPathBaseType_Member_1_PT":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_PT
	case "Enum_CodeSegmentPathBaseType_Member_1_TF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_TF
	case "Enum_CodeSegmentPathBaseType_Member_1_CA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CA
	case "Enum_CodeSegmentPathBaseType_Member_1_CD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CD
	case "Enum_CodeSegmentPathBaseType_Member_1_CI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CI
	case "Enum_CodeSegmentPathBaseType_Member_1_CR":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CR
	case "Enum_CodeSegmentPathBaseType_Member_1_CF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_CF
	case "Enum_CodeSegmentPathBaseType_Member_1_DF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_DF
	case "Enum_CodeSegmentPathBaseType_Member_1_FA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FA
	case "Enum_CodeSegmentPathBaseType_Member_1_FC":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FC
	case "Enum_CodeSegmentPathBaseType_Member_1_FT":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FT
	case "Enum_CodeSegmentPathBaseType_Member_1_FM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FM
	case "Enum_CodeSegmentPathBaseType_Member_1_VM":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VM
	case "Enum_CodeSegmentPathBaseType_Member_1_FD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_FD
	case "Enum_CodeSegmentPathBaseType_Member_1_VR":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VR
	case "Enum_CodeSegmentPathBaseType_Member_1_VD":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VD
	case "Enum_CodeSegmentPathBaseType_Member_1_VI":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VI
	case "Enum_CodeSegmentPathBaseType_Member_1_VA":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_VA
	case "Enum_CodeSegmentPathBaseType_Member_1_RF":
		*enum_codesegmentpathbasetype_member_1 = Enum_CodeSegmentPathBaseType_Member_1_RF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesegmentpathbasetype_member_1 *Enum_CodeSegmentPathBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesegmentpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSegmentPathBaseType_Member_1_AF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_AF"
	case Enum_CodeSegmentPathBaseType_Member_1_HF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_HF"
	case Enum_CodeSegmentPathBaseType_Member_1_HA:
		res = "Enum_CodeSegmentPathBaseType_Member_1_HA"
	case Enum_CodeSegmentPathBaseType_Member_1_HM:
		res = "Enum_CodeSegmentPathBaseType_Member_1_HM"
	case Enum_CodeSegmentPathBaseType_Member_1_IF_XSD:
		res = "Enum_CodeSegmentPathBaseType_Member_1_IF_XSD"
	case Enum_CodeSegmentPathBaseType_Member_1_PI:
		res = "Enum_CodeSegmentPathBaseType_Member_1_PI"
	case Enum_CodeSegmentPathBaseType_Member_1_PT:
		res = "Enum_CodeSegmentPathBaseType_Member_1_PT"
	case Enum_CodeSegmentPathBaseType_Member_1_TF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_TF"
	case Enum_CodeSegmentPathBaseType_Member_1_CA:
		res = "Enum_CodeSegmentPathBaseType_Member_1_CA"
	case Enum_CodeSegmentPathBaseType_Member_1_CD:
		res = "Enum_CodeSegmentPathBaseType_Member_1_CD"
	case Enum_CodeSegmentPathBaseType_Member_1_CI:
		res = "Enum_CodeSegmentPathBaseType_Member_1_CI"
	case Enum_CodeSegmentPathBaseType_Member_1_CR:
		res = "Enum_CodeSegmentPathBaseType_Member_1_CR"
	case Enum_CodeSegmentPathBaseType_Member_1_CF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_CF"
	case Enum_CodeSegmentPathBaseType_Member_1_DF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_DF"
	case Enum_CodeSegmentPathBaseType_Member_1_FA:
		res = "Enum_CodeSegmentPathBaseType_Member_1_FA"
	case Enum_CodeSegmentPathBaseType_Member_1_FC:
		res = "Enum_CodeSegmentPathBaseType_Member_1_FC"
	case Enum_CodeSegmentPathBaseType_Member_1_FT:
		res = "Enum_CodeSegmentPathBaseType_Member_1_FT"
	case Enum_CodeSegmentPathBaseType_Member_1_FM:
		res = "Enum_CodeSegmentPathBaseType_Member_1_FM"
	case Enum_CodeSegmentPathBaseType_Member_1_VM:
		res = "Enum_CodeSegmentPathBaseType_Member_1_VM"
	case Enum_CodeSegmentPathBaseType_Member_1_FD:
		res = "Enum_CodeSegmentPathBaseType_Member_1_FD"
	case Enum_CodeSegmentPathBaseType_Member_1_VR:
		res = "Enum_CodeSegmentPathBaseType_Member_1_VR"
	case Enum_CodeSegmentPathBaseType_Member_1_VD:
		res = "Enum_CodeSegmentPathBaseType_Member_1_VD"
	case Enum_CodeSegmentPathBaseType_Member_1_VI:
		res = "Enum_CodeSegmentPathBaseType_Member_1_VI"
	case Enum_CodeSegmentPathBaseType_Member_1_VA:
		res = "Enum_CodeSegmentPathBaseType_Member_1_VA"
	case Enum_CodeSegmentPathBaseType_Member_1_RF:
		res = "Enum_CodeSegmentPathBaseType_Member_1_RF"
	}
	return
}

func (enum_codesegmentpathbasetype_member_1 Enum_CodeSegmentPathBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_AF")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_HF")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_HA")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_HM")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_IF_XSD")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_PI")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_PT")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_TF")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_CA")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_CD")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_CI")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_CR")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_CF")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_DF")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_FA")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_FC")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_FT")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_FM")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_VM")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_FD")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_VR")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_VD")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_VI")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_VA")
	res = append(res, "Enum_CodeSegmentPathBaseType_Member_1_RF")

	return
}

func (enum_codesegmentpathbasetype_member_1 Enum_CodeSegmentPathBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AF")
	res = append(res, "HF")
	res = append(res, "HA")
	res = append(res, "HM")
	res = append(res, "IF")
	res = append(res, "PI")
	res = append(res, "PT")
	res = append(res, "TF")
	res = append(res, "CA")
	res = append(res, "CD")
	res = append(res, "CI")
	res = append(res, "CR")
	res = append(res, "CF")
	res = append(res, "DF")
	res = append(res, "FA")
	res = append(res, "FC")
	res = append(res, "FT")
	res = append(res, "FM")
	res = append(res, "VM")
	res = append(res, "FD")
	res = append(res, "VR")
	res = append(res, "VD")
	res = append(res, "VI")
	res = append(res, "VA")
	res = append(res, "RF")

	return
}

// Utility function for Enum_CodeSegmentTerminationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesegmentterminationbasetype_member_1 Enum_CodeSegmentTerminationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesegmentterminationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE:
		res = "ALTITUDE"
	case Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE:
		res = "DISTANCE"
	case Enum_CodeSegmentTerminationBaseType_Member_1_DURATION:
		res = "DURATION"
	case Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT:
		res = "INTERCEPT"
	}
	return
}

func (enum_codesegmentterminationbasetype_member_1 *Enum_CodeSegmentTerminationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ALTITUDE":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE
		return
	case "DISTANCE":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE
		return
	case "DURATION":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_DURATION
		return
	case "INTERCEPT":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesegmentterminationbasetype_member_1 *Enum_CodeSegmentTerminationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE
	case "Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE
	case "Enum_CodeSegmentTerminationBaseType_Member_1_DURATION":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_DURATION
	case "Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT":
		*enum_codesegmentterminationbasetype_member_1 = Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesegmentterminationbasetype_member_1 *Enum_CodeSegmentTerminationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesegmentterminationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE:
		res = "Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE"
	case Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE:
		res = "Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE"
	case Enum_CodeSegmentTerminationBaseType_Member_1_DURATION:
		res = "Enum_CodeSegmentTerminationBaseType_Member_1_DURATION"
	case Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT:
		res = "Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT"
	}
	return
}

func (enum_codesegmentterminationbasetype_member_1 Enum_CodeSegmentTerminationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSegmentTerminationBaseType_Member_1_ALTITUDE")
	res = append(res, "Enum_CodeSegmentTerminationBaseType_Member_1_DISTANCE")
	res = append(res, "Enum_CodeSegmentTerminationBaseType_Member_1_DURATION")
	res = append(res, "Enum_CodeSegmentTerminationBaseType_Member_1_INTERCEPT")

	return
}

func (enum_codesegmentterminationbasetype_member_1 Enum_CodeSegmentTerminationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ALTITUDE")
	res = append(res, "DISTANCE")
	res = append(res, "DURATION")
	res = append(res, "INTERCEPT")

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

// Utility function for Enum_CodeSpeedReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codespeedreferencebasetype_member_1 Enum_CodeSpeedReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codespeedreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpeedReferenceBaseType_Member_1_IAS:
		res = "IAS"
	case Enum_CodeSpeedReferenceBaseType_Member_1_TAS:
		res = "TAS"
	case Enum_CodeSpeedReferenceBaseType_Member_1_GS:
		res = "GS"
	}
	return
}

func (enum_codespeedreferencebasetype_member_1 *Enum_CodeSpeedReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IAS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_IAS
		return
	case "TAS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_TAS
		return
	case "GS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_GS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codespeedreferencebasetype_member_1 *Enum_CodeSpeedReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSpeedReferenceBaseType_Member_1_IAS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_IAS
	case "Enum_CodeSpeedReferenceBaseType_Member_1_TAS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_TAS
	case "Enum_CodeSpeedReferenceBaseType_Member_1_GS":
		*enum_codespeedreferencebasetype_member_1 = Enum_CodeSpeedReferenceBaseType_Member_1_GS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codespeedreferencebasetype_member_1 *Enum_CodeSpeedReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codespeedreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpeedReferenceBaseType_Member_1_IAS:
		res = "Enum_CodeSpeedReferenceBaseType_Member_1_IAS"
	case Enum_CodeSpeedReferenceBaseType_Member_1_TAS:
		res = "Enum_CodeSpeedReferenceBaseType_Member_1_TAS"
	case Enum_CodeSpeedReferenceBaseType_Member_1_GS:
		res = "Enum_CodeSpeedReferenceBaseType_Member_1_GS"
	}
	return
}

func (enum_codespeedreferencebasetype_member_1 Enum_CodeSpeedReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSpeedReferenceBaseType_Member_1_IAS")
	res = append(res, "Enum_CodeSpeedReferenceBaseType_Member_1_TAS")
	res = append(res, "Enum_CodeSpeedReferenceBaseType_Member_1_GS")

	return
}

func (enum_codespeedreferencebasetype_member_1 Enum_CodeSpeedReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IAS")
	res = append(res, "TAS")
	res = append(res, "GS")

	return
}

// Utility function for Enum_CodeTAABaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetaabasetype_member_1 Enum_CodeTAABaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetaabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTAABaseType_Member_1_LEFT_BASE:
		res = "LEFT_BASE"
	case Enum_CodeTAABaseType_Member_1_RIGHT_BASE:
		res = "RIGHT_BASE"
	case Enum_CodeTAABaseType_Member_1_STRAIGHT_IN:
		res = "STRAIGHT_IN"
	}
	return
}

func (enum_codetaabasetype_member_1 *Enum_CodeTAABaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LEFT_BASE":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_LEFT_BASE
		return
	case "RIGHT_BASE":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_RIGHT_BASE
		return
	case "STRAIGHT_IN":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_STRAIGHT_IN
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetaabasetype_member_1 *Enum_CodeTAABaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTAABaseType_Member_1_LEFT_BASE":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_LEFT_BASE
	case "Enum_CodeTAABaseType_Member_1_RIGHT_BASE":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_RIGHT_BASE
	case "Enum_CodeTAABaseType_Member_1_STRAIGHT_IN":
		*enum_codetaabasetype_member_1 = Enum_CodeTAABaseType_Member_1_STRAIGHT_IN
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetaabasetype_member_1 *Enum_CodeTAABaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetaabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTAABaseType_Member_1_LEFT_BASE:
		res = "Enum_CodeTAABaseType_Member_1_LEFT_BASE"
	case Enum_CodeTAABaseType_Member_1_RIGHT_BASE:
		res = "Enum_CodeTAABaseType_Member_1_RIGHT_BASE"
	case Enum_CodeTAABaseType_Member_1_STRAIGHT_IN:
		res = "Enum_CodeTAABaseType_Member_1_STRAIGHT_IN"
	}
	return
}

func (enum_codetaabasetype_member_1 Enum_CodeTAABaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTAABaseType_Member_1_LEFT_BASE")
	res = append(res, "Enum_CodeTAABaseType_Member_1_RIGHT_BASE")
	res = append(res, "Enum_CodeTAABaseType_Member_1_STRAIGHT_IN")

	return
}

func (enum_codetaabasetype_member_1 Enum_CodeTAABaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LEFT_BASE")
	res = append(res, "RIGHT_BASE")
	res = append(res, "STRAIGHT_IN")

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

// Utility function for Enum_CodeTrajectoryBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetrajectorybasetype_member_1 Enum_CodeTrajectoryBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetrajectorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT:
		res = "STRAIGHT"
	case Enum_CodeTrajectoryBaseType_Member_1_ARC:
		res = "ARC"
	case Enum_CodeTrajectoryBaseType_Member_1_PT:
		res = "PT"
	case Enum_CodeTrajectoryBaseType_Member_1_BASETURN:
		res = "BASETURN"
	case Enum_CodeTrajectoryBaseType_Member_1_HOLDING:
		res = "HOLDING"
	}
	return
}

func (enum_codetrajectorybasetype_member_1 *Enum_CodeTrajectoryBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STRAIGHT":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT
		return
	case "ARC":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_ARC
		return
	case "PT":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_PT
		return
	case "BASETURN":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_BASETURN
		return
	case "HOLDING":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_HOLDING
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetrajectorybasetype_member_1 *Enum_CodeTrajectoryBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT
	case "Enum_CodeTrajectoryBaseType_Member_1_ARC":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_ARC
	case "Enum_CodeTrajectoryBaseType_Member_1_PT":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_PT
	case "Enum_CodeTrajectoryBaseType_Member_1_BASETURN":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_BASETURN
	case "Enum_CodeTrajectoryBaseType_Member_1_HOLDING":
		*enum_codetrajectorybasetype_member_1 = Enum_CodeTrajectoryBaseType_Member_1_HOLDING
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetrajectorybasetype_member_1 *Enum_CodeTrajectoryBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetrajectorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT:
		res = "Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT"
	case Enum_CodeTrajectoryBaseType_Member_1_ARC:
		res = "Enum_CodeTrajectoryBaseType_Member_1_ARC"
	case Enum_CodeTrajectoryBaseType_Member_1_PT:
		res = "Enum_CodeTrajectoryBaseType_Member_1_PT"
	case Enum_CodeTrajectoryBaseType_Member_1_BASETURN:
		res = "Enum_CodeTrajectoryBaseType_Member_1_BASETURN"
	case Enum_CodeTrajectoryBaseType_Member_1_HOLDING:
		res = "Enum_CodeTrajectoryBaseType_Member_1_HOLDING"
	}
	return
}

func (enum_codetrajectorybasetype_member_1 Enum_CodeTrajectoryBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTrajectoryBaseType_Member_1_STRAIGHT")
	res = append(res, "Enum_CodeTrajectoryBaseType_Member_1_ARC")
	res = append(res, "Enum_CodeTrajectoryBaseType_Member_1_PT")
	res = append(res, "Enum_CodeTrajectoryBaseType_Member_1_BASETURN")
	res = append(res, "Enum_CodeTrajectoryBaseType_Member_1_HOLDING")

	return
}

func (enum_codetrajectorybasetype_member_1 Enum_CodeTrajectoryBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STRAIGHT")
	res = append(res, "ARC")
	res = append(res, "PT")
	res = append(res, "BASETURN")
	res = append(res, "HOLDING")

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

// Utility function for Enum_CodeUpperAlphaBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeupperalphabasetype_member_1 Enum_CodeUpperAlphaBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeupperalphabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUpperAlphaBaseType_Member_1_A:
		res = "A"
	case Enum_CodeUpperAlphaBaseType_Member_1_B:
		res = "B"
	case Enum_CodeUpperAlphaBaseType_Member_1_C:
		res = "C"
	case Enum_CodeUpperAlphaBaseType_Member_1_D:
		res = "D"
	case Enum_CodeUpperAlphaBaseType_Member_1_E:
		res = "E"
	case Enum_CodeUpperAlphaBaseType_Member_1_F:
		res = "F"
	case Enum_CodeUpperAlphaBaseType_Member_1_G:
		res = "G"
	case Enum_CodeUpperAlphaBaseType_Member_1_H:
		res = "H"
	case Enum_CodeUpperAlphaBaseType_Member_1_I:
		res = "I"
	case Enum_CodeUpperAlphaBaseType_Member_1_J:
		res = "J"
	case Enum_CodeUpperAlphaBaseType_Member_1_K:
		res = "K"
	case Enum_CodeUpperAlphaBaseType_Member_1_L:
		res = "L"
	case Enum_CodeUpperAlphaBaseType_Member_1_M:
		res = "M"
	case Enum_CodeUpperAlphaBaseType_Member_1_N:
		res = "N"
	case Enum_CodeUpperAlphaBaseType_Member_1_O:
		res = "O"
	case Enum_CodeUpperAlphaBaseType_Member_1_P:
		res = "P"
	case Enum_CodeUpperAlphaBaseType_Member_1_Q:
		res = "Q"
	case Enum_CodeUpperAlphaBaseType_Member_1_R:
		res = "R"
	case Enum_CodeUpperAlphaBaseType_Member_1_S:
		res = "S"
	case Enum_CodeUpperAlphaBaseType_Member_1_T:
		res = "T"
	case Enum_CodeUpperAlphaBaseType_Member_1_U:
		res = "U"
	case Enum_CodeUpperAlphaBaseType_Member_1_V:
		res = "V"
	case Enum_CodeUpperAlphaBaseType_Member_1_W:
		res = "W"
	case Enum_CodeUpperAlphaBaseType_Member_1_X:
		res = "X"
	case Enum_CodeUpperAlphaBaseType_Member_1_Y:
		res = "Y"
	case Enum_CodeUpperAlphaBaseType_Member_1_Z:
		res = "Z"
	}
	return
}

func (enum_codeupperalphabasetype_member_1 *Enum_CodeUpperAlphaBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_A
		return
	case "B":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_B
		return
	case "C":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_C
		return
	case "D":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_D
		return
	case "E":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_E
		return
	case "F":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_F
		return
	case "G":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_G
		return
	case "H":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_H
		return
	case "I":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_I
		return
	case "J":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_J
		return
	case "K":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_K
		return
	case "L":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_L
		return
	case "M":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_M
		return
	case "N":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_N
		return
	case "O":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_O
		return
	case "P":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_P
		return
	case "Q":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Q
		return
	case "R":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_R
		return
	case "S":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_S
		return
	case "T":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_T
		return
	case "U":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_U
		return
	case "V":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_V
		return
	case "W":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_W
		return
	case "X":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_X
		return
	case "Y":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Y
		return
	case "Z":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Z
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeupperalphabasetype_member_1 *Enum_CodeUpperAlphaBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeUpperAlphaBaseType_Member_1_A":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_A
	case "Enum_CodeUpperAlphaBaseType_Member_1_B":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_B
	case "Enum_CodeUpperAlphaBaseType_Member_1_C":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_C
	case "Enum_CodeUpperAlphaBaseType_Member_1_D":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_D
	case "Enum_CodeUpperAlphaBaseType_Member_1_E":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_E
	case "Enum_CodeUpperAlphaBaseType_Member_1_F":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_F
	case "Enum_CodeUpperAlphaBaseType_Member_1_G":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_G
	case "Enum_CodeUpperAlphaBaseType_Member_1_H":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_H
	case "Enum_CodeUpperAlphaBaseType_Member_1_I":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_I
	case "Enum_CodeUpperAlphaBaseType_Member_1_J":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_J
	case "Enum_CodeUpperAlphaBaseType_Member_1_K":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_K
	case "Enum_CodeUpperAlphaBaseType_Member_1_L":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_L
	case "Enum_CodeUpperAlphaBaseType_Member_1_M":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_M
	case "Enum_CodeUpperAlphaBaseType_Member_1_N":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_N
	case "Enum_CodeUpperAlphaBaseType_Member_1_O":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_O
	case "Enum_CodeUpperAlphaBaseType_Member_1_P":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_P
	case "Enum_CodeUpperAlphaBaseType_Member_1_Q":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Q
	case "Enum_CodeUpperAlphaBaseType_Member_1_R":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_R
	case "Enum_CodeUpperAlphaBaseType_Member_1_S":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_S
	case "Enum_CodeUpperAlphaBaseType_Member_1_T":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_T
	case "Enum_CodeUpperAlphaBaseType_Member_1_U":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_U
	case "Enum_CodeUpperAlphaBaseType_Member_1_V":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_V
	case "Enum_CodeUpperAlphaBaseType_Member_1_W":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_W
	case "Enum_CodeUpperAlphaBaseType_Member_1_X":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_X
	case "Enum_CodeUpperAlphaBaseType_Member_1_Y":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Y
	case "Enum_CodeUpperAlphaBaseType_Member_1_Z":
		*enum_codeupperalphabasetype_member_1 = Enum_CodeUpperAlphaBaseType_Member_1_Z
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeupperalphabasetype_member_1 *Enum_CodeUpperAlphaBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeupperalphabasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUpperAlphaBaseType_Member_1_A:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_A"
	case Enum_CodeUpperAlphaBaseType_Member_1_B:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_B"
	case Enum_CodeUpperAlphaBaseType_Member_1_C:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_C"
	case Enum_CodeUpperAlphaBaseType_Member_1_D:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_D"
	case Enum_CodeUpperAlphaBaseType_Member_1_E:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_E"
	case Enum_CodeUpperAlphaBaseType_Member_1_F:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_F"
	case Enum_CodeUpperAlphaBaseType_Member_1_G:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_G"
	case Enum_CodeUpperAlphaBaseType_Member_1_H:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_H"
	case Enum_CodeUpperAlphaBaseType_Member_1_I:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_I"
	case Enum_CodeUpperAlphaBaseType_Member_1_J:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_J"
	case Enum_CodeUpperAlphaBaseType_Member_1_K:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_K"
	case Enum_CodeUpperAlphaBaseType_Member_1_L:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_L"
	case Enum_CodeUpperAlphaBaseType_Member_1_M:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_M"
	case Enum_CodeUpperAlphaBaseType_Member_1_N:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_N"
	case Enum_CodeUpperAlphaBaseType_Member_1_O:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_O"
	case Enum_CodeUpperAlphaBaseType_Member_1_P:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_P"
	case Enum_CodeUpperAlphaBaseType_Member_1_Q:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_Q"
	case Enum_CodeUpperAlphaBaseType_Member_1_R:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_R"
	case Enum_CodeUpperAlphaBaseType_Member_1_S:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_S"
	case Enum_CodeUpperAlphaBaseType_Member_1_T:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_T"
	case Enum_CodeUpperAlphaBaseType_Member_1_U:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_U"
	case Enum_CodeUpperAlphaBaseType_Member_1_V:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_V"
	case Enum_CodeUpperAlphaBaseType_Member_1_W:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_W"
	case Enum_CodeUpperAlphaBaseType_Member_1_X:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_X"
	case Enum_CodeUpperAlphaBaseType_Member_1_Y:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_Y"
	case Enum_CodeUpperAlphaBaseType_Member_1_Z:
		res = "Enum_CodeUpperAlphaBaseType_Member_1_Z"
	}
	return
}

func (enum_codeupperalphabasetype_member_1 Enum_CodeUpperAlphaBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_A")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_B")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_C")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_D")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_E")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_F")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_G")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_H")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_I")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_J")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_K")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_L")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_M")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_N")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_O")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_P")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_Q")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_R")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_S")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_T")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_U")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_V")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_W")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_X")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_Y")
	res = append(res, "Enum_CodeUpperAlphaBaseType_Member_1_Z")

	return
}

func (enum_codeupperalphabasetype_member_1 Enum_CodeUpperAlphaBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")
	res = append(res, "E")
	res = append(res, "F")
	res = append(res, "G")
	res = append(res, "H")
	res = append(res, "I")
	res = append(res, "J")
	res = append(res, "K")
	res = append(res, "L")
	res = append(res, "M")
	res = append(res, "N")
	res = append(res, "O")
	res = append(res, "P")
	res = append(res, "Q")
	res = append(res, "R")
	res = append(res, "S")
	res = append(res, "T")
	res = append(res, "U")
	res = append(res, "V")
	res = append(res, "W")
	res = append(res, "X")
	res = append(res, "Y")
	res = append(res, "Z")

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
