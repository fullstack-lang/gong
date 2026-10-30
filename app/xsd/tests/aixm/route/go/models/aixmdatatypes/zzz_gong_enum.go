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

// Utility function for Enum_CodeAerialRefuellingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaerialrefuellingbasetype_member_1 Enum_CodeAerialRefuellingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaerialrefuellingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingBaseType_Member_1_TRACK:
		res = "TRACK"
	case Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR:
		res = "ANCHOR"
	case Enum_CodeAerialRefuellingBaseType_Member_1_BOTH:
		res = "BOTH"
	}
	return
}

func (enum_codeaerialrefuellingbasetype_member_1 *Enum_CodeAerialRefuellingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TRACK":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_TRACK
		return
	case "ANCHOR":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR
		return
	case "BOTH":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_BOTH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaerialrefuellingbasetype_member_1 *Enum_CodeAerialRefuellingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAerialRefuellingBaseType_Member_1_TRACK":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_TRACK
	case "Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR
	case "Enum_CodeAerialRefuellingBaseType_Member_1_BOTH":
		*enum_codeaerialrefuellingbasetype_member_1 = Enum_CodeAerialRefuellingBaseType_Member_1_BOTH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaerialrefuellingbasetype_member_1 *Enum_CodeAerialRefuellingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaerialrefuellingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingBaseType_Member_1_TRACK:
		res = "Enum_CodeAerialRefuellingBaseType_Member_1_TRACK"
	case Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR:
		res = "Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR"
	case Enum_CodeAerialRefuellingBaseType_Member_1_BOTH:
		res = "Enum_CodeAerialRefuellingBaseType_Member_1_BOTH"
	}
	return
}

func (enum_codeaerialrefuellingbasetype_member_1 Enum_CodeAerialRefuellingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAerialRefuellingBaseType_Member_1_TRACK")
	res = append(res, "Enum_CodeAerialRefuellingBaseType_Member_1_ANCHOR")
	res = append(res, "Enum_CodeAerialRefuellingBaseType_Member_1_BOTH")

	return
}

func (enum_codeaerialrefuellingbasetype_member_1 Enum_CodeAerialRefuellingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TRACK")
	res = append(res, "ANCHOR")
	res = append(res, "BOTH")

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

// Utility function for Enum_CodeAerialRefuellingPrefixBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaerialrefuellingprefixbasetype_member_1 Enum_CodeAerialRefuellingPrefixBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaerialrefuellingprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR:
		res = "AR"
	}
	return
}

func (enum_codeaerialrefuellingprefixbasetype_member_1 *Enum_CodeAerialRefuellingPrefixBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AR":
		*enum_codeaerialrefuellingprefixbasetype_member_1 = Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaerialrefuellingprefixbasetype_member_1 *Enum_CodeAerialRefuellingPrefixBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR":
		*enum_codeaerialrefuellingprefixbasetype_member_1 = Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaerialrefuellingprefixbasetype_member_1 *Enum_CodeAerialRefuellingPrefixBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaerialrefuellingprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR:
		res = "Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR"
	}
	return
}

func (enum_codeaerialrefuellingprefixbasetype_member_1 Enum_CodeAerialRefuellingPrefixBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAerialRefuellingPrefixBaseType_Member_1_AR")

	return
}

func (enum_codeaerialrefuellingprefixbasetype_member_1 Enum_CodeAerialRefuellingPrefixBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AR")

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

// Utility function for Enum_CodeCardinalDirectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecardinaldirectionbasetype_member_1 Enum_CodeCardinalDirectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecardinaldirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCardinalDirectionBaseType_Member_1_N:
		res = "N"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NE:
		res = "NE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_E:
		res = "E"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SE:
		res = "SE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_S:
		res = "S"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SW:
		res = "SW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_W:
		res = "W"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NW:
		res = "NW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NNE:
		res = "NNE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_ENE:
		res = "ENE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_ESE:
		res = "ESE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SSE:
		res = "SSE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SSW:
		res = "SSW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_WSW:
		res = "WSW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_WNW:
		res = "WNW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NNW:
		res = "NNW"
	}
	return
}

func (enum_codecardinaldirectionbasetype_member_1 *Enum_CodeCardinalDirectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "N":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_N
		return
	case "NE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NE
		return
	case "E":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_E
		return
	case "SE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SE
		return
	case "S":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_S
		return
	case "SW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SW
		return
	case "W":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_W
		return
	case "NW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NW
		return
	case "NNE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NNE
		return
	case "ENE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_ENE
		return
	case "ESE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_ESE
		return
	case "SSE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SSE
		return
	case "SSW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SSW
		return
	case "WSW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_WSW
		return
	case "WNW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_WNW
		return
	case "NNW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NNW
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecardinaldirectionbasetype_member_1 *Enum_CodeCardinalDirectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCardinalDirectionBaseType_Member_1_N":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_N
	case "Enum_CodeCardinalDirectionBaseType_Member_1_NE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_E":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_E
	case "Enum_CodeCardinalDirectionBaseType_Member_1_SE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_S":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_S
	case "Enum_CodeCardinalDirectionBaseType_Member_1_SW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SW
	case "Enum_CodeCardinalDirectionBaseType_Member_1_W":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_W
	case "Enum_CodeCardinalDirectionBaseType_Member_1_NW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NW
	case "Enum_CodeCardinalDirectionBaseType_Member_1_NNE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NNE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_ENE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_ENE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_ESE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_ESE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_SSE":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SSE
	case "Enum_CodeCardinalDirectionBaseType_Member_1_SSW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_SSW
	case "Enum_CodeCardinalDirectionBaseType_Member_1_WSW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_WSW
	case "Enum_CodeCardinalDirectionBaseType_Member_1_WNW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_WNW
	case "Enum_CodeCardinalDirectionBaseType_Member_1_NNW":
		*enum_codecardinaldirectionbasetype_member_1 = Enum_CodeCardinalDirectionBaseType_Member_1_NNW
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecardinaldirectionbasetype_member_1 *Enum_CodeCardinalDirectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecardinaldirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCardinalDirectionBaseType_Member_1_N:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_N"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_NE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_E:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_E"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_SE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_S:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_S"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_SW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_W:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_W"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_NW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NNE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_NNE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_ENE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_ENE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_ESE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_ESE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SSE:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_SSE"
	case Enum_CodeCardinalDirectionBaseType_Member_1_SSW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_SSW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_WSW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_WSW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_WNW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_WNW"
	case Enum_CodeCardinalDirectionBaseType_Member_1_NNW:
		res = "Enum_CodeCardinalDirectionBaseType_Member_1_NNW"
	}
	return
}

func (enum_codecardinaldirectionbasetype_member_1 Enum_CodeCardinalDirectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_N")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_NE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_E")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_SE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_S")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_SW")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_W")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_NW")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_NNE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_ENE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_ESE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_SSE")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_SSW")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_WSW")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_WNW")
	res = append(res, "Enum_CodeCardinalDirectionBaseType_Member_1_NNW")

	return
}

func (enum_codecardinaldirectionbasetype_member_1 Enum_CodeCardinalDirectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "N")
	res = append(res, "NE")
	res = append(res, "E")
	res = append(res, "SE")
	res = append(res, "S")
	res = append(res, "SW")
	res = append(res, "W")
	res = append(res, "NW")
	res = append(res, "NNE")
	res = append(res, "ENE")
	res = append(res, "ESE")
	res = append(res, "SSE")
	res = append(res, "SSW")
	res = append(res, "WSW")
	res = append(res, "WNW")
	res = append(res, "NNW")

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

// Utility function for Enum_CodeComparisonBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecomparisonbasetype_member_1 Enum_CodeComparisonBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecomparisonbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeComparisonBaseType_Member_1_LESS:
		res = "LESS"
	case Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL:
		res = "LESS_OR_EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_EQUAL:
		res = "EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL:
		res = "GREATER_OR_EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_GREATER:
		res = "GREATER"
	}
	return
}

func (enum_codecomparisonbasetype_member_1 *Enum_CodeComparisonBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LESS":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_LESS
		return
	case "LESS_OR_EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL
		return
	case "EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_EQUAL
		return
	case "GREATER_OR_EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL
		return
	case "GREATER":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_GREATER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecomparisonbasetype_member_1 *Enum_CodeComparisonBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeComparisonBaseType_Member_1_LESS":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_LESS
	case "Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL
	case "Enum_CodeComparisonBaseType_Member_1_EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_EQUAL
	case "Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL
	case "Enum_CodeComparisonBaseType_Member_1_GREATER":
		*enum_codecomparisonbasetype_member_1 = Enum_CodeComparisonBaseType_Member_1_GREATER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecomparisonbasetype_member_1 *Enum_CodeComparisonBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecomparisonbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeComparisonBaseType_Member_1_LESS:
		res = "Enum_CodeComparisonBaseType_Member_1_LESS"
	case Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL:
		res = "Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_EQUAL:
		res = "Enum_CodeComparisonBaseType_Member_1_EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL:
		res = "Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL"
	case Enum_CodeComparisonBaseType_Member_1_GREATER:
		res = "Enum_CodeComparisonBaseType_Member_1_GREATER"
	}
	return
}

func (enum_codecomparisonbasetype_member_1 Enum_CodeComparisonBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeComparisonBaseType_Member_1_LESS")
	res = append(res, "Enum_CodeComparisonBaseType_Member_1_LESS_OR_EQUAL")
	res = append(res, "Enum_CodeComparisonBaseType_Member_1_EQUAL")
	res = append(res, "Enum_CodeComparisonBaseType_Member_1_GREATER_OR_EQUAL")
	res = append(res, "Enum_CodeComparisonBaseType_Member_1_GREATER")

	return
}

func (enum_codecomparisonbasetype_member_1 Enum_CodeComparisonBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LESS")
	res = append(res, "LESS_OR_EQUAL")
	res = append(res, "EQUAL")
	res = append(res, "GREATER_OR_EQUAL")
	res = append(res, "GREATER")

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

// Utility function for Enum_CodeFlightRestrictionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightrestrictionbasetype_member_1 Enum_CodeFlightRestrictionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightrestrictionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightRestrictionBaseType_Member_1_FORBID:
		res = "FORBID"
	case Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY:
		res = "MANDATORY"
	case Enum_CodeFlightRestrictionBaseType_Member_1_CLSD:
		res = "CLSD"
	case Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED:
		res = "ALLOWED"
	}
	return
}

func (enum_codeflightrestrictionbasetype_member_1 *Enum_CodeFlightRestrictionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FORBID":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_FORBID
		return
	case "MANDATORY":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY
		return
	case "CLSD":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_CLSD
		return
	case "ALLOWED":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightrestrictionbasetype_member_1 *Enum_CodeFlightRestrictionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightRestrictionBaseType_Member_1_FORBID":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_FORBID
	case "Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY
	case "Enum_CodeFlightRestrictionBaseType_Member_1_CLSD":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_CLSD
	case "Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED":
		*enum_codeflightrestrictionbasetype_member_1 = Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightrestrictionbasetype_member_1 *Enum_CodeFlightRestrictionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightrestrictionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightRestrictionBaseType_Member_1_FORBID:
		res = "Enum_CodeFlightRestrictionBaseType_Member_1_FORBID"
	case Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY:
		res = "Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY"
	case Enum_CodeFlightRestrictionBaseType_Member_1_CLSD:
		res = "Enum_CodeFlightRestrictionBaseType_Member_1_CLSD"
	case Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED:
		res = "Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED"
	}
	return
}

func (enum_codeflightrestrictionbasetype_member_1 Enum_CodeFlightRestrictionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightRestrictionBaseType_Member_1_FORBID")
	res = append(res, "Enum_CodeFlightRestrictionBaseType_Member_1_MANDATORY")
	res = append(res, "Enum_CodeFlightRestrictionBaseType_Member_1_CLSD")
	res = append(res, "Enum_CodeFlightRestrictionBaseType_Member_1_ALLOWED")

	return
}

func (enum_codeflightrestrictionbasetype_member_1 Enum_CodeFlightRestrictionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FORBID")
	res = append(res, "MANDATORY")
	res = append(res, "CLSD")
	res = append(res, "ALLOWED")

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

// Utility function for Enum_CodeFlowConditionOperationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflowconditionoperationbasetype_member_1 Enum_CodeFlowConditionOperationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflowconditionoperationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlowConditionOperationBaseType_Member_1_AND:
		res = "AND"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT:
		res = "ANDNOT"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_OR:
		res = "OR"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ:
		res = "SEQ"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_NONE:
		res = "NONE"
	}
	return
}

func (enum_codeflowconditionoperationbasetype_member_1 *Enum_CodeFlowConditionOperationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AND":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_AND
		return
	case "ANDNOT":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT
		return
	case "OR":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_OR
		return
	case "SEQ":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ
		return
	case "NONE":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_NONE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflowconditionoperationbasetype_member_1 *Enum_CodeFlowConditionOperationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlowConditionOperationBaseType_Member_1_AND":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_AND
	case "Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT
	case "Enum_CodeFlowConditionOperationBaseType_Member_1_OR":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_OR
	case "Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ
	case "Enum_CodeFlowConditionOperationBaseType_Member_1_NONE":
		*enum_codeflowconditionoperationbasetype_member_1 = Enum_CodeFlowConditionOperationBaseType_Member_1_NONE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflowconditionoperationbasetype_member_1 *Enum_CodeFlowConditionOperationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflowconditionoperationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlowConditionOperationBaseType_Member_1_AND:
		res = "Enum_CodeFlowConditionOperationBaseType_Member_1_AND"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT:
		res = "Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_OR:
		res = "Enum_CodeFlowConditionOperationBaseType_Member_1_OR"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ:
		res = "Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ"
	case Enum_CodeFlowConditionOperationBaseType_Member_1_NONE:
		res = "Enum_CodeFlowConditionOperationBaseType_Member_1_NONE"
	}
	return
}

func (enum_codeflowconditionoperationbasetype_member_1 Enum_CodeFlowConditionOperationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlowConditionOperationBaseType_Member_1_AND")
	res = append(res, "Enum_CodeFlowConditionOperationBaseType_Member_1_ANDNOT")
	res = append(res, "Enum_CodeFlowConditionOperationBaseType_Member_1_OR")
	res = append(res, "Enum_CodeFlowConditionOperationBaseType_Member_1_SEQ")
	res = append(res, "Enum_CodeFlowConditionOperationBaseType_Member_1_NONE")

	return
}

func (enum_codeflowconditionoperationbasetype_member_1 Enum_CodeFlowConditionOperationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AND")
	res = append(res, "ANDNOT")
	res = append(res, "OR")
	res = append(res, "SEQ")
	res = append(res, "NONE")

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

// Utility function for Enum_CodeLevelBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelevelbasetype_member_1 Enum_CodeLevelBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelevelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelBaseType_Member_1_UPPER:
		res = "UPPER"
	case Enum_CodeLevelBaseType_Member_1_LOWER:
		res = "LOWER"
	case Enum_CodeLevelBaseType_Member_1_BOTH:
		res = "BOTH"
	}
	return
}

func (enum_codelevelbasetype_member_1 *Enum_CodeLevelBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "UPPER":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_UPPER
		return
	case "LOWER":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_LOWER
		return
	case "BOTH":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_BOTH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelevelbasetype_member_1 *Enum_CodeLevelBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLevelBaseType_Member_1_UPPER":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_UPPER
	case "Enum_CodeLevelBaseType_Member_1_LOWER":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_LOWER
	case "Enum_CodeLevelBaseType_Member_1_BOTH":
		*enum_codelevelbasetype_member_1 = Enum_CodeLevelBaseType_Member_1_BOTH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelevelbasetype_member_1 *Enum_CodeLevelBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelevelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelBaseType_Member_1_UPPER:
		res = "Enum_CodeLevelBaseType_Member_1_UPPER"
	case Enum_CodeLevelBaseType_Member_1_LOWER:
		res = "Enum_CodeLevelBaseType_Member_1_LOWER"
	case Enum_CodeLevelBaseType_Member_1_BOTH:
		res = "Enum_CodeLevelBaseType_Member_1_BOTH"
	}
	return
}

func (enum_codelevelbasetype_member_1 Enum_CodeLevelBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLevelBaseType_Member_1_UPPER")
	res = append(res, "Enum_CodeLevelBaseType_Member_1_LOWER")
	res = append(res, "Enum_CodeLevelBaseType_Member_1_BOTH")

	return
}

func (enum_codelevelbasetype_member_1 Enum_CodeLevelBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "UPPER")
	res = append(res, "LOWER")
	res = append(res, "BOTH")

	return
}

// Utility function for Enum_CodeLocationQualifierBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelocationqualifierbasetype_member_1 Enum_CodeLocationQualifierBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelocationqualifierbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLocationQualifierBaseType_Member_1_DEP:
		res = "DEP"
	case Enum_CodeLocationQualifierBaseType_Member_1_XNG:
		res = "XNG"
	case Enum_CodeLocationQualifierBaseType_Member_1_ARR:
		res = "ARR"
	case Enum_CodeLocationQualifierBaseType_Member_1_ACT:
		res = "ACT"
	case Enum_CodeLocationQualifierBaseType_Member_1_AVBL:
		res = "AVBL"
	}
	return
}

func (enum_codelocationqualifierbasetype_member_1 *Enum_CodeLocationQualifierBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DEP":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_DEP
		return
	case "XNG":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_XNG
		return
	case "ARR":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_ARR
		return
	case "ACT":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_ACT
		return
	case "AVBL":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_AVBL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelocationqualifierbasetype_member_1 *Enum_CodeLocationQualifierBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLocationQualifierBaseType_Member_1_DEP":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_DEP
	case "Enum_CodeLocationQualifierBaseType_Member_1_XNG":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_XNG
	case "Enum_CodeLocationQualifierBaseType_Member_1_ARR":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_ARR
	case "Enum_CodeLocationQualifierBaseType_Member_1_ACT":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_ACT
	case "Enum_CodeLocationQualifierBaseType_Member_1_AVBL":
		*enum_codelocationqualifierbasetype_member_1 = Enum_CodeLocationQualifierBaseType_Member_1_AVBL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelocationqualifierbasetype_member_1 *Enum_CodeLocationQualifierBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelocationqualifierbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLocationQualifierBaseType_Member_1_DEP:
		res = "Enum_CodeLocationQualifierBaseType_Member_1_DEP"
	case Enum_CodeLocationQualifierBaseType_Member_1_XNG:
		res = "Enum_CodeLocationQualifierBaseType_Member_1_XNG"
	case Enum_CodeLocationQualifierBaseType_Member_1_ARR:
		res = "Enum_CodeLocationQualifierBaseType_Member_1_ARR"
	case Enum_CodeLocationQualifierBaseType_Member_1_ACT:
		res = "Enum_CodeLocationQualifierBaseType_Member_1_ACT"
	case Enum_CodeLocationQualifierBaseType_Member_1_AVBL:
		res = "Enum_CodeLocationQualifierBaseType_Member_1_AVBL"
	}
	return
}

func (enum_codelocationqualifierbasetype_member_1 Enum_CodeLocationQualifierBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLocationQualifierBaseType_Member_1_DEP")
	res = append(res, "Enum_CodeLocationQualifierBaseType_Member_1_XNG")
	res = append(res, "Enum_CodeLocationQualifierBaseType_Member_1_ARR")
	res = append(res, "Enum_CodeLocationQualifierBaseType_Member_1_ACT")
	res = append(res, "Enum_CodeLocationQualifierBaseType_Member_1_AVBL")

	return
}

func (enum_codelocationqualifierbasetype_member_1 Enum_CodeLocationQualifierBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DEP")
	res = append(res, "XNG")
	res = append(res, "ARR")
	res = append(res, "ACT")
	res = append(res, "AVBL")

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

// Utility function for Enum_CodeMilitaryTrainingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemilitarytrainingbasetype_member_1 Enum_CodeMilitaryTrainingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemilitarytrainingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryTrainingBaseType_Member_1_IR:
		res = "IR"
	case Enum_CodeMilitaryTrainingBaseType_Member_1_VR:
		res = "VR"
	case Enum_CodeMilitaryTrainingBaseType_Member_1_SR:
		res = "SR"
	}
	return
}

func (enum_codemilitarytrainingbasetype_member_1 *Enum_CodeMilitaryTrainingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_IR
		return
	case "VR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_VR
		return
	case "SR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_SR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemilitarytrainingbasetype_member_1 *Enum_CodeMilitaryTrainingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMilitaryTrainingBaseType_Member_1_IR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_IR
	case "Enum_CodeMilitaryTrainingBaseType_Member_1_VR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_VR
	case "Enum_CodeMilitaryTrainingBaseType_Member_1_SR":
		*enum_codemilitarytrainingbasetype_member_1 = Enum_CodeMilitaryTrainingBaseType_Member_1_SR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemilitarytrainingbasetype_member_1 *Enum_CodeMilitaryTrainingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemilitarytrainingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMilitaryTrainingBaseType_Member_1_IR:
		res = "Enum_CodeMilitaryTrainingBaseType_Member_1_IR"
	case Enum_CodeMilitaryTrainingBaseType_Member_1_VR:
		res = "Enum_CodeMilitaryTrainingBaseType_Member_1_VR"
	case Enum_CodeMilitaryTrainingBaseType_Member_1_SR:
		res = "Enum_CodeMilitaryTrainingBaseType_Member_1_SR"
	}
	return
}

func (enum_codemilitarytrainingbasetype_member_1 Enum_CodeMilitaryTrainingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMilitaryTrainingBaseType_Member_1_IR")
	res = append(res, "Enum_CodeMilitaryTrainingBaseType_Member_1_VR")
	res = append(res, "Enum_CodeMilitaryTrainingBaseType_Member_1_SR")

	return
}

func (enum_codemilitarytrainingbasetype_member_1 Enum_CodeMilitaryTrainingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IR")
	res = append(res, "VR")
	res = append(res, "SR")

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

// Utility function for Enum_CodeRouteAvailabilityBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderouteavailabilitybasetype_member_1 Enum_CodeRouteAvailabilityBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderouteavailabilitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN:
		res = "OPEN"
	case Enum_CodeRouteAvailabilityBaseType_Member_1_COND:
		res = "COND"
	case Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD:
		res = "CLSD"
	}
	return
}

func (enum_coderouteavailabilitybasetype_member_1 *Enum_CodeRouteAvailabilityBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OPEN":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN
		return
	case "COND":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_COND
		return
	case "CLSD":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderouteavailabilitybasetype_member_1 *Enum_CodeRouteAvailabilityBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN
	case "Enum_CodeRouteAvailabilityBaseType_Member_1_COND":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_COND
	case "Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD":
		*enum_coderouteavailabilitybasetype_member_1 = Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderouteavailabilitybasetype_member_1 *Enum_CodeRouteAvailabilityBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderouteavailabilitybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN:
		res = "Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN"
	case Enum_CodeRouteAvailabilityBaseType_Member_1_COND:
		res = "Enum_CodeRouteAvailabilityBaseType_Member_1_COND"
	case Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD:
		res = "Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD"
	}
	return
}

func (enum_coderouteavailabilitybasetype_member_1 Enum_CodeRouteAvailabilityBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteAvailabilityBaseType_Member_1_OPEN")
	res = append(res, "Enum_CodeRouteAvailabilityBaseType_Member_1_COND")
	res = append(res, "Enum_CodeRouteAvailabilityBaseType_Member_1_CLSD")

	return
}

func (enum_coderouteavailabilitybasetype_member_1 Enum_CodeRouteAvailabilityBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OPEN")
	res = append(res, "COND")
	res = append(res, "CLSD")

	return
}

// Utility function for Enum_CodeRouteBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutebasetype_member_1 Enum_CodeRouteBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteBaseType_Member_1_ATS:
		res = "ATS"
	case Enum_CodeRouteBaseType_Member_1_NAT:
		res = "NAT"
	}
	return
}

func (enum_coderoutebasetype_member_1 *Enum_CodeRouteBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ATS":
		*enum_coderoutebasetype_member_1 = Enum_CodeRouteBaseType_Member_1_ATS
		return
	case "NAT":
		*enum_coderoutebasetype_member_1 = Enum_CodeRouteBaseType_Member_1_NAT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutebasetype_member_1 *Enum_CodeRouteBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteBaseType_Member_1_ATS":
		*enum_coderoutebasetype_member_1 = Enum_CodeRouteBaseType_Member_1_ATS
	case "Enum_CodeRouteBaseType_Member_1_NAT":
		*enum_coderoutebasetype_member_1 = Enum_CodeRouteBaseType_Member_1_NAT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutebasetype_member_1 *Enum_CodeRouteBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteBaseType_Member_1_ATS:
		res = "Enum_CodeRouteBaseType_Member_1_ATS"
	case Enum_CodeRouteBaseType_Member_1_NAT:
		res = "Enum_CodeRouteBaseType_Member_1_NAT"
	}
	return
}

func (enum_coderoutebasetype_member_1 Enum_CodeRouteBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteBaseType_Member_1_ATS")
	res = append(res, "Enum_CodeRouteBaseType_Member_1_NAT")

	return
}

func (enum_coderoutebasetype_member_1 Enum_CodeRouteBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ATS")
	res = append(res, "NAT")

	return
}

// Utility function for Enum_CodeRouteDesignatorLetterBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutedesignatorletterbasetype_member_1 Enum_CodeRouteDesignatorLetterBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutedesignatorletterbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_A:
		res = "A"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_B:
		res = "B"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_G:
		res = "G"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_H:
		res = "H"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_J:
		res = "J"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_L:
		res = "L"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_M:
		res = "M"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_N:
		res = "N"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_P:
		res = "P"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q:
		res = "Q"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_R:
		res = "R"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_T:
		res = "T"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_V:
		res = "V"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_W:
		res = "W"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y:
		res = "Y"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z:
		res = "Z"
	}
	return
}

func (enum_coderoutedesignatorletterbasetype_member_1 *Enum_CodeRouteDesignatorLetterBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_A
		return
	case "B":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_B
		return
	case "G":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_G
		return
	case "H":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_H
		return
	case "J":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_J
		return
	case "L":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_L
		return
	case "M":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_M
		return
	case "N":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_N
		return
	case "P":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_P
		return
	case "Q":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q
		return
	case "R":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_R
		return
	case "T":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_T
		return
	case "V":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_V
		return
	case "W":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_W
		return
	case "Y":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y
		return
	case "Z":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutedesignatorletterbasetype_member_1 *Enum_CodeRouteDesignatorLetterBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_A":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_A
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_B":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_B
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_G":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_G
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_H":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_H
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_J":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_J
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_L":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_L
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_M":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_M
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_N":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_N
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_P":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_P
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_R":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_R
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_T":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_T
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_V":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_V
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_W":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_W
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y
	case "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z":
		*enum_coderoutedesignatorletterbasetype_member_1 = Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutedesignatorletterbasetype_member_1 *Enum_CodeRouteDesignatorLetterBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutedesignatorletterbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_A:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_A"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_B:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_B"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_G:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_G"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_H:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_H"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_J:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_J"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_L:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_L"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_M:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_M"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_N:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_N"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_P:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_P"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_R:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_R"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_T:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_T"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_V:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_V"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_W:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_W"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y"
	case Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z:
		res = "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z"
	}
	return
}

func (enum_coderoutedesignatorletterbasetype_member_1 Enum_CodeRouteDesignatorLetterBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_A")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_B")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_G")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_H")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_J")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_L")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_M")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_N")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_P")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Q")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_R")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_T")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_V")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_W")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Y")
	res = append(res, "Enum_CodeRouteDesignatorLetterBaseType_Member_1_Z")

	return
}

func (enum_coderoutedesignatorletterbasetype_member_1 Enum_CodeRouteDesignatorLetterBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "G")
	res = append(res, "H")
	res = append(res, "J")
	res = append(res, "L")
	res = append(res, "M")
	res = append(res, "N")
	res = append(res, "P")
	res = append(res, "Q")
	res = append(res, "R")
	res = append(res, "T")
	res = append(res, "V")
	res = append(res, "W")
	res = append(res, "Y")
	res = append(res, "Z")

	return
}

// Utility function for Enum_CodeRouteDesignatorPrefixBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutedesignatorprefixbasetype_member_1 Enum_CodeRouteDesignatorPrefixBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutedesignatorprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K:
		res = "K"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U:
		res = "U"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S:
		res = "S"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T:
		res = "T"
	}
	return
}

func (enum_coderoutedesignatorprefixbasetype_member_1 *Enum_CodeRouteDesignatorPrefixBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "K":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K
		return
	case "U":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U
		return
	case "S":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S
		return
	case "T":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutedesignatorprefixbasetype_member_1 *Enum_CodeRouteDesignatorPrefixBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K
	case "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U
	case "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S
	case "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T":
		*enum_coderoutedesignatorprefixbasetype_member_1 = Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutedesignatorprefixbasetype_member_1 *Enum_CodeRouteDesignatorPrefixBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutedesignatorprefixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K:
		res = "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U:
		res = "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S:
		res = "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S"
	case Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T:
		res = "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T"
	}
	return
}

func (enum_coderoutedesignatorprefixbasetype_member_1 Enum_CodeRouteDesignatorPrefixBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_K")
	res = append(res, "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_U")
	res = append(res, "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_S")
	res = append(res, "Enum_CodeRouteDesignatorPrefixBaseType_Member_1_T")

	return
}

func (enum_coderoutedesignatorprefixbasetype_member_1 Enum_CodeRouteDesignatorPrefixBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "K")
	res = append(res, "U")
	res = append(res, "S")
	res = append(res, "T")

	return
}

// Utility function for Enum_CodeRouteDesignatorSuffixBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutedesignatorsuffixbasetype_member_1 Enum_CodeRouteDesignatorSuffixBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutedesignatorsuffixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F:
		res = "F"
	case Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G:
		res = "G"
	}
	return
}

func (enum_coderoutedesignatorsuffixbasetype_member_1 *Enum_CodeRouteDesignatorSuffixBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "F":
		*enum_coderoutedesignatorsuffixbasetype_member_1 = Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F
		return
	case "G":
		*enum_coderoutedesignatorsuffixbasetype_member_1 = Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutedesignatorsuffixbasetype_member_1 *Enum_CodeRouteDesignatorSuffixBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F":
		*enum_coderoutedesignatorsuffixbasetype_member_1 = Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F
	case "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G":
		*enum_coderoutedesignatorsuffixbasetype_member_1 = Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutedesignatorsuffixbasetype_member_1 *Enum_CodeRouteDesignatorSuffixBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutedesignatorsuffixbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F:
		res = "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F"
	case Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G:
		res = "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G"
	}
	return
}

func (enum_coderoutedesignatorsuffixbasetype_member_1 Enum_CodeRouteDesignatorSuffixBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_F")
	res = append(res, "Enum_CodeRouteDesignatorSuffixBaseType_Member_1_G")

	return
}

func (enum_coderoutedesignatorsuffixbasetype_member_1 Enum_CodeRouteDesignatorSuffixBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "F")
	res = append(res, "G")

	return
}

// Utility function for Enum_CodeRouteNavigationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutenavigationbasetype_member_1 Enum_CodeRouteNavigationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutenavigationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteNavigationBaseType_Member_1_CONV:
		res = "CONV"
	case Enum_CodeRouteNavigationBaseType_Member_1_RNAV:
		res = "RNAV"
	case Enum_CodeRouteNavigationBaseType_Member_1_TACAN:
		res = "TACAN"
	}
	return
}

func (enum_coderoutenavigationbasetype_member_1 *Enum_CodeRouteNavigationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CONV":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_CONV
		return
	case "RNAV":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_RNAV
		return
	case "TACAN":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_TACAN
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutenavigationbasetype_member_1 *Enum_CodeRouteNavigationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteNavigationBaseType_Member_1_CONV":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_CONV
	case "Enum_CodeRouteNavigationBaseType_Member_1_RNAV":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_RNAV
	case "Enum_CodeRouteNavigationBaseType_Member_1_TACAN":
		*enum_coderoutenavigationbasetype_member_1 = Enum_CodeRouteNavigationBaseType_Member_1_TACAN
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutenavigationbasetype_member_1 *Enum_CodeRouteNavigationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutenavigationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteNavigationBaseType_Member_1_CONV:
		res = "Enum_CodeRouteNavigationBaseType_Member_1_CONV"
	case Enum_CodeRouteNavigationBaseType_Member_1_RNAV:
		res = "Enum_CodeRouteNavigationBaseType_Member_1_RNAV"
	case Enum_CodeRouteNavigationBaseType_Member_1_TACAN:
		res = "Enum_CodeRouteNavigationBaseType_Member_1_TACAN"
	}
	return
}

func (enum_coderoutenavigationbasetype_member_1 Enum_CodeRouteNavigationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteNavigationBaseType_Member_1_CONV")
	res = append(res, "Enum_CodeRouteNavigationBaseType_Member_1_RNAV")
	res = append(res, "Enum_CodeRouteNavigationBaseType_Member_1_TACAN")

	return
}

func (enum_coderoutenavigationbasetype_member_1 Enum_CodeRouteNavigationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CONV")
	res = append(res, "RNAV")
	res = append(res, "TACAN")

	return
}

// Utility function for Enum_CodeRouteOriginBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderouteoriginbasetype_member_1 Enum_CodeRouteOriginBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderouteoriginbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteOriginBaseType_Member_1_INTL:
		res = "INTL"
	case Enum_CodeRouteOriginBaseType_Member_1_DOM:
		res = "DOM"
	case Enum_CodeRouteOriginBaseType_Member_1_BOTH:
		res = "BOTH"
	}
	return
}

func (enum_coderouteoriginbasetype_member_1 *Enum_CodeRouteOriginBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "INTL":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_INTL
		return
	case "DOM":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_DOM
		return
	case "BOTH":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_BOTH
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderouteoriginbasetype_member_1 *Enum_CodeRouteOriginBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteOriginBaseType_Member_1_INTL":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_INTL
	case "Enum_CodeRouteOriginBaseType_Member_1_DOM":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_DOM
	case "Enum_CodeRouteOriginBaseType_Member_1_BOTH":
		*enum_coderouteoriginbasetype_member_1 = Enum_CodeRouteOriginBaseType_Member_1_BOTH
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderouteoriginbasetype_member_1 *Enum_CodeRouteOriginBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderouteoriginbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteOriginBaseType_Member_1_INTL:
		res = "Enum_CodeRouteOriginBaseType_Member_1_INTL"
	case Enum_CodeRouteOriginBaseType_Member_1_DOM:
		res = "Enum_CodeRouteOriginBaseType_Member_1_DOM"
	case Enum_CodeRouteOriginBaseType_Member_1_BOTH:
		res = "Enum_CodeRouteOriginBaseType_Member_1_BOTH"
	}
	return
}

func (enum_coderouteoriginbasetype_member_1 Enum_CodeRouteOriginBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteOriginBaseType_Member_1_INTL")
	res = append(res, "Enum_CodeRouteOriginBaseType_Member_1_DOM")
	res = append(res, "Enum_CodeRouteOriginBaseType_Member_1_BOTH")

	return
}

func (enum_coderouteoriginbasetype_member_1 Enum_CodeRouteOriginBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "INTL")
	res = append(res, "DOM")
	res = append(res, "BOTH")

	return
}

// Utility function for Enum_CodeRouteSegmentPathBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderoutesegmentpathbasetype_member_1 Enum_CodeRouteSegmentPathBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderoutesegmentpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteSegmentPathBaseType_Member_1_GRC:
		res = "GRC"
	case Enum_CodeRouteSegmentPathBaseType_Member_1_RHL:
		res = "RHL"
	case Enum_CodeRouteSegmentPathBaseType_Member_1_GDS:
		res = "GDS"
	}
	return
}

func (enum_coderoutesegmentpathbasetype_member_1 *Enum_CodeRouteSegmentPathBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "GRC":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_GRC
		return
	case "RHL":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_RHL
		return
	case "GDS":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_GDS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderoutesegmentpathbasetype_member_1 *Enum_CodeRouteSegmentPathBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRouteSegmentPathBaseType_Member_1_GRC":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_GRC
	case "Enum_CodeRouteSegmentPathBaseType_Member_1_RHL":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_RHL
	case "Enum_CodeRouteSegmentPathBaseType_Member_1_GDS":
		*enum_coderoutesegmentpathbasetype_member_1 = Enum_CodeRouteSegmentPathBaseType_Member_1_GDS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderoutesegmentpathbasetype_member_1 *Enum_CodeRouteSegmentPathBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderoutesegmentpathbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRouteSegmentPathBaseType_Member_1_GRC:
		res = "Enum_CodeRouteSegmentPathBaseType_Member_1_GRC"
	case Enum_CodeRouteSegmentPathBaseType_Member_1_RHL:
		res = "Enum_CodeRouteSegmentPathBaseType_Member_1_RHL"
	case Enum_CodeRouteSegmentPathBaseType_Member_1_GDS:
		res = "Enum_CodeRouteSegmentPathBaseType_Member_1_GDS"
	}
	return
}

func (enum_coderoutesegmentpathbasetype_member_1 Enum_CodeRouteSegmentPathBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRouteSegmentPathBaseType_Member_1_GRC")
	res = append(res, "Enum_CodeRouteSegmentPathBaseType_Member_1_RHL")
	res = append(res, "Enum_CodeRouteSegmentPathBaseType_Member_1_GDS")

	return
}

func (enum_coderoutesegmentpathbasetype_member_1 Enum_CodeRouteSegmentPathBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "GRC")
	res = append(res, "RHL")
	res = append(res, "GDS")

	return
}

// Utility function for Enum_CodeRuleProcedureBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderuleprocedurebasetype_member_1 Enum_CodeRuleProcedureBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderuleprocedurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRuleProcedureBaseType_Member_1_RULE:
		res = "RULE"
	case Enum_CodeRuleProcedureBaseType_Member_1_LAW:
		res = "LAW"
	case Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE:
		res = "PROCEDURE"
	case Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE:
		res = "PRACTICE"
	case Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF:
		res = "ICAO_DIFF"
	}
	return
}

func (enum_coderuleprocedurebasetype_member_1 *Enum_CodeRuleProcedureBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RULE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_RULE
		return
	case "LAW":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_LAW
		return
	case "PROCEDURE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE
		return
	case "PRACTICE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE
		return
	case "ICAO_DIFF":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderuleprocedurebasetype_member_1 *Enum_CodeRuleProcedureBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRuleProcedureBaseType_Member_1_RULE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_RULE
	case "Enum_CodeRuleProcedureBaseType_Member_1_LAW":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_LAW
	case "Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE
	case "Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE
	case "Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF":
		*enum_coderuleprocedurebasetype_member_1 = Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderuleprocedurebasetype_member_1 *Enum_CodeRuleProcedureBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderuleprocedurebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRuleProcedureBaseType_Member_1_RULE:
		res = "Enum_CodeRuleProcedureBaseType_Member_1_RULE"
	case Enum_CodeRuleProcedureBaseType_Member_1_LAW:
		res = "Enum_CodeRuleProcedureBaseType_Member_1_LAW"
	case Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE:
		res = "Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE"
	case Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE:
		res = "Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE"
	case Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF:
		res = "Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF"
	}
	return
}

func (enum_coderuleprocedurebasetype_member_1 Enum_CodeRuleProcedureBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRuleProcedureBaseType_Member_1_RULE")
	res = append(res, "Enum_CodeRuleProcedureBaseType_Member_1_LAW")
	res = append(res, "Enum_CodeRuleProcedureBaseType_Member_1_PROCEDURE")
	res = append(res, "Enum_CodeRuleProcedureBaseType_Member_1_PRACTICE")
	res = append(res, "Enum_CodeRuleProcedureBaseType_Member_1_ICAO_DIFF")

	return
}

func (enum_coderuleprocedurebasetype_member_1 Enum_CodeRuleProcedureBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RULE")
	res = append(res, "LAW")
	res = append(res, "PROCEDURE")
	res = append(res, "PRACTICE")
	res = append(res, "ICAO_DIFF")

	return
}

// Utility function for Enum_CodeRuleProcedureTitleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderuleproceduretitlebasetype_member_1 Enum_CodeRuleProcedureTitleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderuleproceduretitlebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT:
		res = "ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED:
		res = "ENTRY_TRANSIT_DEPARTURE_SCHEDULED"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED:
		res = "ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE:
		res = "ENTRY_TRANSIT_DEPARTURE_PRIVATE"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT:
		res = "PUBLIC_HEALTH_MEASURES_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS:
		res = "CUSTOMS_REQUIREMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS:
		res = "IMMIGRATION_REQUIREMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS:
		res = "PUBLIC_HEALTH_MEASURES_PASSENGERS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO:
		res = "ENTRY_TRANSIT_DEPARTURE_CARGO"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS:
		res = "AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS:
		res = "NATIONAL_REGULATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS:
		res = "INTERNATIONAL_AGREEMENTS_CONVENTIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES:
		res = "DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS:
		res = "MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS:
		res = "ABBREVIATIONS_AIS_PUBLICATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES:
		res = "AERODROME_HELIPORT_CHARGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES:
		res = "AIR_NAVIGATION_SERVICES_CHARGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL:
		res = "FLIGHT_RULES_GENERAL"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES:
		res = "VISUAL_FLIGHT_RULES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES:
		res = "INSTRUMENT_FLIGHT_RULES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION:
		res = "ATS_AIRSPACE_CLASSIFICATION"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES:
		res = "HOLDING_APPROACH_DEPARTURE_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES:
		res = "ATS_SURVEILLANCE_SERVICES_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES:
		res = "ALTIMETER_SETTING_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES:
		res = "REGIONAL_SUPPLEMENTARY_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT:
		res = "AIR_TRAFFIC_FLOW_MANAGEMENT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING:
		res = "FLIGHT_PLANNING"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES:
		res = "ADDRESSING_FLIGHT_PLAN_MESSAGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT:
		res = "INTERCEPTION_CIVIL_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE:
		res = "UNLAWFUL_INTERFERENCE"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS:
		res = "AIR_TRAFFIC_INCIDENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY:
		res = "AERODROME_HELIPORT_AVAILABILITY"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS:
		res = "LOCAL_TRAFFIC_REGULATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES:
		res = "NOISE_ABATEMENT_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES:
		res = "AERODROME_FLIGHT_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION:
		res = "AERODROME_BIRD_CONCENTRATION"
	}
	return
}

func (enum_coderuleproceduretitlebasetype_member_1 *Enum_CodeRuleProcedureTitleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT
		return
	case "ENTRY_TRANSIT_DEPARTURE_SCHEDULED":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED
		return
	case "ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED
		return
	case "ENTRY_TRANSIT_DEPARTURE_PRIVATE":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE
		return
	case "PUBLIC_HEALTH_MEASURES_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT
		return
	case "CUSTOMS_REQUIREMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS
		return
	case "IMMIGRATION_REQUIREMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS
		return
	case "PUBLIC_HEALTH_MEASURES_PASSENGERS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS
		return
	case "ENTRY_TRANSIT_DEPARTURE_CARGO":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO
		return
	case "AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS
		return
	case "NATIONAL_REGULATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS
		return
	case "INTERNATIONAL_AGREEMENTS_CONVENTIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS
		return
	case "DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES
		return
	case "MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS
		return
	case "ABBREVIATIONS_AIS_PUBLICATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS
		return
	case "AERODROME_HELIPORT_CHARGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES
		return
	case "AIR_NAVIGATION_SERVICES_CHARGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES
		return
	case "FLIGHT_RULES_GENERAL":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL
		return
	case "VISUAL_FLIGHT_RULES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES
		return
	case "INSTRUMENT_FLIGHT_RULES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES
		return
	case "ATS_AIRSPACE_CLASSIFICATION":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION
		return
	case "HOLDING_APPROACH_DEPARTURE_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES
		return
	case "ATS_SURVEILLANCE_SERVICES_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES
		return
	case "ALTIMETER_SETTING_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES
		return
	case "REGIONAL_SUPPLEMENTARY_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES
		return
	case "AIR_TRAFFIC_FLOW_MANAGEMENT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT
		return
	case "FLIGHT_PLANNING":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING
		return
	case "ADDRESSING_FLIGHT_PLAN_MESSAGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES
		return
	case "INTERCEPTION_CIVIL_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT
		return
	case "UNLAWFUL_INTERFERENCE":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE
		return
	case "AIR_TRAFFIC_INCIDENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS
		return
	case "AERODROME_HELIPORT_AVAILABILITY":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY
		return
	case "LOCAL_TRAFFIC_REGULATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS
		return
	case "NOISE_ABATEMENT_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES
		return
	case "AERODROME_FLIGHT_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES
		return
	case "AERODROME_BIRD_CONCENTRATION":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderuleproceduretitlebasetype_member_1 *Enum_CodeRuleProcedureTitleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES
	case "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION":
		*enum_coderuleproceduretitlebasetype_member_1 = Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderuleproceduretitlebasetype_member_1 *Enum_CodeRuleProcedureTitleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderuleproceduretitlebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES"
	case Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION:
		res = "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION"
	}
	return
}

func (enum_coderuleproceduretitlebasetype_member_1 Enum_CodeRuleProcedureTitleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_SCHEDULED")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_PRIVATE")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_AIRCRAFT")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_CUSTOMS_REQUIREMENTS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_IMMIGRATION_REQUIREMENTS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_PUBLIC_HEALTH_MEASURES_PASSENGERS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ENTRY_TRANSIT_DEPARTURE_CARGO")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_NATIONAL_REGULATIONS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERNATIONAL_AGREEMENTS_CONVENTIONS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ABBREVIATIONS_AIS_PUBLICATIONS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_CHARGES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_NAVIGATION_SERVICES_CHARGES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_RULES_GENERAL")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_VISUAL_FLIGHT_RULES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_INSTRUMENT_FLIGHT_RULES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_AIRSPACE_CLASSIFICATION")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_HOLDING_APPROACH_DEPARTURE_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ATS_SURVEILLANCE_SERVICES_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ALTIMETER_SETTING_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_REGIONAL_SUPPLEMENTARY_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_FLOW_MANAGEMENT")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_FLIGHT_PLANNING")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_ADDRESSING_FLIGHT_PLAN_MESSAGES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_INTERCEPTION_CIVIL_AIRCRAFT")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_UNLAWFUL_INTERFERENCE")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AIR_TRAFFIC_INCIDENTS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_HELIPORT_AVAILABILITY")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_LOCAL_TRAFFIC_REGULATIONS")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_NOISE_ABATEMENT_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_FLIGHT_PROCEDURES")
	res = append(res, "Enum_CodeRuleProcedureTitleBaseType_Member_1_AERODROME_BIRD_CONCENTRATION")

	return
}

func (enum_coderuleproceduretitlebasetype_member_1 Enum_CodeRuleProcedureTitleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENTRY_TRANSIT_DEPARTURE_OF_AIRCRAFT")
	res = append(res, "ENTRY_TRANSIT_DEPARTURE_SCHEDULED")
	res = append(res, "ENTRY_TRANSIT_DEPARTURE_NON_SCHEDULED")
	res = append(res, "ENTRY_TRANSIT_DEPARTURE_PRIVATE")
	res = append(res, "PUBLIC_HEALTH_MEASURES_AIRCRAFT")
	res = append(res, "CUSTOMS_REQUIREMENTS")
	res = append(res, "IMMIGRATION_REQUIREMENTS")
	res = append(res, "PUBLIC_HEALTH_MEASURES_PASSENGERS")
	res = append(res, "ENTRY_TRANSIT_DEPARTURE_CARGO")
	res = append(res, "AIRCRAFT_INSTRUMENTS_EQUIPMENT_FLIGHT_DOCUMENTS")
	res = append(res, "NATIONAL_REGULATIONS")
	res = append(res, "INTERNATIONAL_AGREEMENTS_CONVENTIONS")
	res = append(res, "DIFFERENCES_ICAO_STANDARDS_RECOMMENDED_PRACTICES_PROCEDURES")
	res = append(res, "MEASURING_SYSTEM_AIRCRAFT_MARKINGS_HOLIDAYS")
	res = append(res, "ABBREVIATIONS_AIS_PUBLICATIONS")
	res = append(res, "AERODROME_HELIPORT_CHARGES")
	res = append(res, "AIR_NAVIGATION_SERVICES_CHARGES")
	res = append(res, "FLIGHT_RULES_GENERAL")
	res = append(res, "VISUAL_FLIGHT_RULES")
	res = append(res, "INSTRUMENT_FLIGHT_RULES")
	res = append(res, "ATS_AIRSPACE_CLASSIFICATION")
	res = append(res, "HOLDING_APPROACH_DEPARTURE_PROCEDURES")
	res = append(res, "ATS_SURVEILLANCE_SERVICES_PROCEDURES")
	res = append(res, "ALTIMETER_SETTING_PROCEDURES")
	res = append(res, "REGIONAL_SUPPLEMENTARY_PROCEDURES")
	res = append(res, "AIR_TRAFFIC_FLOW_MANAGEMENT")
	res = append(res, "FLIGHT_PLANNING")
	res = append(res, "ADDRESSING_FLIGHT_PLAN_MESSAGES")
	res = append(res, "INTERCEPTION_CIVIL_AIRCRAFT")
	res = append(res, "UNLAWFUL_INTERFERENCE")
	res = append(res, "AIR_TRAFFIC_INCIDENTS")
	res = append(res, "AERODROME_HELIPORT_AVAILABILITY")
	res = append(res, "LOCAL_TRAFFIC_REGULATIONS")
	res = append(res, "NOISE_ABATEMENT_PROCEDURES")
	res = append(res, "AERODROME_FLIGHT_PROCEDURES")
	res = append(res, "AERODROME_BIRD_CONCENTRATION")

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

// Utility function for Enum_CodeTACANChannelBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codetacanchannelbasetype_member_1 Enum_CodeTACANChannelBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codetacanchannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTACANChannelBaseType_Member_1_N_1X:
		res = "1X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_1Y:
		res = "1Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_2X:
		res = "2X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_2Y:
		res = "2Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_3X:
		res = "3X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_3Y:
		res = "3Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_4X:
		res = "4X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_4Y:
		res = "4Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_5X:
		res = "5X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_5Y:
		res = "5Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_6X:
		res = "6X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_6Y:
		res = "6Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_7X:
		res = "7X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_7Y:
		res = "7Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_8X:
		res = "8X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_8Y:
		res = "8Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_9X:
		res = "9X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_9Y:
		res = "9Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_10X:
		res = "10X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_10Y:
		res = "10Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_11X:
		res = "11X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_11Y:
		res = "11Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_12X:
		res = "12X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_12Y:
		res = "12Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_13X:
		res = "13X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_13Y:
		res = "13Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_14X:
		res = "14X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_14Y:
		res = "14Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_15X:
		res = "15X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_15Y:
		res = "15Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_16X:
		res = "16X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_16Y:
		res = "16Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17X:
		res = "17X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17Y:
		res = "17Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17Z:
		res = "17Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18X:
		res = "18X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18W:
		res = "18W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18Y:
		res = "18Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18Z:
		res = "18Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19X:
		res = "19X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19Y:
		res = "19Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19Z:
		res = "19Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20X:
		res = "20X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20W:
		res = "20W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20Y:
		res = "20Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20Z:
		res = "20Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21X:
		res = "21X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21Y:
		res = "21Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21Z:
		res = "21Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22X:
		res = "22X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22W:
		res = "22W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22Y:
		res = "22Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22Z:
		res = "22Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23X:
		res = "23X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23Y:
		res = "23Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23Z:
		res = "23Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24X:
		res = "24X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24W:
		res = "24W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24Y:
		res = "24Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24Z:
		res = "24Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25X:
		res = "25X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25Y:
		res = "25Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25Z:
		res = "25Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26X:
		res = "26X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26W:
		res = "26W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26Y:
		res = "26Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26Z:
		res = "26Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27X:
		res = "27X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27Y:
		res = "27Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27Z:
		res = "27Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28X:
		res = "28X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28W:
		res = "28W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28Y:
		res = "28Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28Z:
		res = "28Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29X:
		res = "29X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29Y:
		res = "29Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29Z:
		res = "29Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30X:
		res = "30X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30W:
		res = "30W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30Y:
		res = "30Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30Z:
		res = "30Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31X:
		res = "31X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31Y:
		res = "31Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31Z:
		res = "31Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32X:
		res = "32X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32W:
		res = "32W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32Y:
		res = "32Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32Z:
		res = "32Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33X:
		res = "33X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33Y:
		res = "33Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33Z:
		res = "33Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34X:
		res = "34X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34W:
		res = "34W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34Y:
		res = "34Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34Z:
		res = "34Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35X:
		res = "35X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35Y:
		res = "35Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35Z:
		res = "35Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36X:
		res = "36X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36W:
		res = "36W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36Y:
		res = "36Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36Z:
		res = "36Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37X:
		res = "37X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37Y:
		res = "37Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37Z:
		res = "37Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38X:
		res = "38X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38W:
		res = "38W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38Y:
		res = "38Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38Z:
		res = "38Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39X:
		res = "39X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39Y:
		res = "39Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39Z:
		res = "39Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40X:
		res = "40X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40W:
		res = "40W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40Y:
		res = "40Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40Z:
		res = "40Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41X:
		res = "41X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41Y:
		res = "41Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41Z:
		res = "41Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42X:
		res = "42X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42W:
		res = "42W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42Y:
		res = "42Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42Z:
		res = "42Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43X:
		res = "43X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43Y:
		res = "43Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43Z:
		res = "43Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44X:
		res = "44X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44W:
		res = "44W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44Y:
		res = "44Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44Z:
		res = "44Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45X:
		res = "45X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45Y:
		res = "45Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45Z:
		res = "45Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46X:
		res = "46X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46W:
		res = "46W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46Y:
		res = "46Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46Z:
		res = "46Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47X:
		res = "47X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47Y:
		res = "47Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47Z:
		res = "47Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48X:
		res = "48X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48W:
		res = "48W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48Y:
		res = "48Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48Z:
		res = "48Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49X:
		res = "49X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49Y:
		res = "49Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49Z:
		res = "49Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50X:
		res = "50X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50W:
		res = "50W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50Y:
		res = "50Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50Z:
		res = "50Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51X:
		res = "51X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51Y:
		res = "51Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51Z:
		res = "51Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52X:
		res = "52X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52W:
		res = "52W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52Y:
		res = "52Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52Z:
		res = "52Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53X:
		res = "53X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53Y:
		res = "53Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53Z:
		res = "53Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54X:
		res = "54X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54W:
		res = "54W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54Y:
		res = "54Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54Z:
		res = "54Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55X:
		res = "55X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55Y:
		res = "55Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55Z:
		res = "55Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56X:
		res = "56X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56W:
		res = "56W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56Y:
		res = "56Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56Z:
		res = "56Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_57X:
		res = "57X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_57Y:
		res = "57Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_58X:
		res = "58X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_58Y:
		res = "58Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_59X:
		res = "59X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_59Y:
		res = "59Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_60X:
		res = "60X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_60Y:
		res = "60Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_61X:
		res = "61X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_61Y:
		res = "61Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_62X:
		res = "62X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_62Y:
		res = "62Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_63X:
		res = "63X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_63Y:
		res = "63Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_64X:
		res = "64X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_64Y:
		res = "64Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_65X:
		res = "65X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_65Y:
		res = "65Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_66X:
		res = "66X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_66Y:
		res = "66Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_67X:
		res = "67X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_67Y:
		res = "67Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_68X:
		res = "68X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_68Y:
		res = "68Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_69X:
		res = "69X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_69Y:
		res = "69Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_70X:
		res = "70X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_70Y:
		res = "70Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_71X:
		res = "71X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_71Y:
		res = "71Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_72X:
		res = "72X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_72Y:
		res = "72Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_73X:
		res = "73X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_73Y:
		res = "73Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_74X:
		res = "74X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_74Y:
		res = "74Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_75X:
		res = "75X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_75Y:
		res = "75Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_76X:
		res = "76X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_76Y:
		res = "76Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_77X:
		res = "77X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_77Y:
		res = "77Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_78X:
		res = "78X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_78Y:
		res = "78Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_79X:
		res = "79X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_79Y:
		res = "79Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80X:
		res = "80X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80Y:
		res = "80Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80Z:
		res = "80Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81X:
		res = "81X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81Y:
		res = "81Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81Z:
		res = "81Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82X:
		res = "82X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82Y:
		res = "82Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82Z:
		res = "82Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83X:
		res = "83X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83Y:
		res = "83Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83Z:
		res = "83Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84X:
		res = "84X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84Y:
		res = "84Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84Z:
		res = "84Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85X:
		res = "85X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85Y:
		res = "85Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85Z:
		res = "85Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86X:
		res = "86X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86Y:
		res = "86Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86Z:
		res = "86Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87X:
		res = "87X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87Y:
		res = "87Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87Z:
		res = "87Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88X:
		res = "88X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88Y:
		res = "88Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88Z:
		res = "88Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89X:
		res = "89X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89Y:
		res = "89Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89Z:
		res = "89Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90X:
		res = "90X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90Y:
		res = "90Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90Z:
		res = "90Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91X:
		res = "91X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91Y:
		res = "91Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91Z:
		res = "91Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92X:
		res = "92X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92Y:
		res = "92Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92Z:
		res = "92Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93X:
		res = "93X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93Y:
		res = "93Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93Z:
		res = "93Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94X:
		res = "94X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94Y:
		res = "94Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94Z:
		res = "94Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95X:
		res = "95X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95Y:
		res = "95Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95Z:
		res = "95Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96X:
		res = "96X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96Y:
		res = "96Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96Z:
		res = "96Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97X:
		res = "97X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97Y:
		res = "97Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97Z:
		res = "97Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98X:
		res = "98X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98Y:
		res = "98Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98Z:
		res = "98Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99X:
		res = "99X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99Y:
		res = "99Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99Z:
		res = "99Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100X:
		res = "100X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100Y:
		res = "100Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100Z:
		res = "100Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101X:
		res = "101X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101Y:
		res = "101Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101Z:
		res = "101Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102X:
		res = "102X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102Y:
		res = "102Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102Z:
		res = "102Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103X:
		res = "103X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103Y:
		res = "103Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103Z:
		res = "103Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104X:
		res = "104X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104Y:
		res = "104Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104Z:
		res = "104Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105X:
		res = "105X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105Y:
		res = "105Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105Z:
		res = "105Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106X:
		res = "106X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106Y:
		res = "106Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106Z:
		res = "106Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107X:
		res = "107X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107Y:
		res = "107Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107Z:
		res = "107Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108X:
		res = "108X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108Y:
		res = "108Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108Z:
		res = "108Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109X:
		res = "109X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109Y:
		res = "109Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109Z:
		res = "109Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110X:
		res = "110X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110Y:
		res = "110Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110Z:
		res = "110Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111X:
		res = "111X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111Y:
		res = "111Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111Z:
		res = "111Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112X:
		res = "112X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112Y:
		res = "112Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112Z:
		res = "112Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113X:
		res = "113X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113Y:
		res = "113Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113Z:
		res = "113Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114X:
		res = "114X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114Y:
		res = "114Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114Z:
		res = "114Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115X:
		res = "115X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115Y:
		res = "115Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115Z:
		res = "115Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116X:
		res = "116X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116Y:
		res = "116Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116Z:
		res = "116Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117X:
		res = "117X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117Y:
		res = "117Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117Z:
		res = "117Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118X:
		res = "118X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118Y:
		res = "118Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118Z:
		res = "118Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119X:
		res = "119X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119Y:
		res = "119Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119Z:
		res = "119Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_120X:
		res = "120X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_120Y:
		res = "120Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_121X:
		res = "121X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_121Y:
		res = "121Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_122X:
		res = "122X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_122Y:
		res = "122Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_123X:
		res = "123X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_123Y:
		res = "123Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_124X:
		res = "124X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_124Y:
		res = "124Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_125X:
		res = "125X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_125Y:
		res = "125Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_126X:
		res = "126X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_126Y:
		res = "126Y"
	}
	return
}

func (enum_codetacanchannelbasetype_member_1 *Enum_CodeTACANChannelBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_1X
		return
	case "1Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_1Y
		return
	case "2X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_2X
		return
	case "2Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_2Y
		return
	case "3X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_3X
		return
	case "3Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_3Y
		return
	case "4X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_4X
		return
	case "4Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_4Y
		return
	case "5X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_5X
		return
	case "5Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_5Y
		return
	case "6X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_6X
		return
	case "6Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_6Y
		return
	case "7X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_7X
		return
	case "7Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_7Y
		return
	case "8X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_8X
		return
	case "8Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_8Y
		return
	case "9X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_9X
		return
	case "9Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_9Y
		return
	case "10X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_10X
		return
	case "10Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_10Y
		return
	case "11X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_11X
		return
	case "11Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_11Y
		return
	case "12X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_12X
		return
	case "12Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_12Y
		return
	case "13X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_13X
		return
	case "13Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_13Y
		return
	case "14X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_14X
		return
	case "14Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_14Y
		return
	case "15X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_15X
		return
	case "15Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_15Y
		return
	case "16X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_16X
		return
	case "16Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_16Y
		return
	case "17X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17X
		return
	case "17Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17Y
		return
	case "17Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17Z
		return
	case "18X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18X
		return
	case "18W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18W
		return
	case "18Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18Y
		return
	case "18Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18Z
		return
	case "19X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19X
		return
	case "19Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19Y
		return
	case "19Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19Z
		return
	case "20X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20X
		return
	case "20W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20W
		return
	case "20Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20Y
		return
	case "20Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20Z
		return
	case "21X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21X
		return
	case "21Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21Y
		return
	case "21Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21Z
		return
	case "22X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22X
		return
	case "22W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22W
		return
	case "22Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22Y
		return
	case "22Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22Z
		return
	case "23X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23X
		return
	case "23Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23Y
		return
	case "23Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23Z
		return
	case "24X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24X
		return
	case "24W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24W
		return
	case "24Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24Y
		return
	case "24Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24Z
		return
	case "25X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25X
		return
	case "25Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25Y
		return
	case "25Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25Z
		return
	case "26X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26X
		return
	case "26W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26W
		return
	case "26Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26Y
		return
	case "26Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26Z
		return
	case "27X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27X
		return
	case "27Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27Y
		return
	case "27Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27Z
		return
	case "28X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28X
		return
	case "28W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28W
		return
	case "28Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28Y
		return
	case "28Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28Z
		return
	case "29X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29X
		return
	case "29Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29Y
		return
	case "29Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29Z
		return
	case "30X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30X
		return
	case "30W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30W
		return
	case "30Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30Y
		return
	case "30Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30Z
		return
	case "31X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31X
		return
	case "31Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31Y
		return
	case "31Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31Z
		return
	case "32X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32X
		return
	case "32W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32W
		return
	case "32Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32Y
		return
	case "32Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32Z
		return
	case "33X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33X
		return
	case "33Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33Y
		return
	case "33Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33Z
		return
	case "34X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34X
		return
	case "34W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34W
		return
	case "34Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34Y
		return
	case "34Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34Z
		return
	case "35X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35X
		return
	case "35Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35Y
		return
	case "35Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35Z
		return
	case "36X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36X
		return
	case "36W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36W
		return
	case "36Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36Y
		return
	case "36Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36Z
		return
	case "37X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37X
		return
	case "37Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37Y
		return
	case "37Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37Z
		return
	case "38X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38X
		return
	case "38W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38W
		return
	case "38Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38Y
		return
	case "38Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38Z
		return
	case "39X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39X
		return
	case "39Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39Y
		return
	case "39Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39Z
		return
	case "40X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40X
		return
	case "40W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40W
		return
	case "40Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40Y
		return
	case "40Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40Z
		return
	case "41X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41X
		return
	case "41Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41Y
		return
	case "41Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41Z
		return
	case "42X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42X
		return
	case "42W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42W
		return
	case "42Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42Y
		return
	case "42Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42Z
		return
	case "43X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43X
		return
	case "43Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43Y
		return
	case "43Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43Z
		return
	case "44X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44X
		return
	case "44W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44W
		return
	case "44Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44Y
		return
	case "44Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44Z
		return
	case "45X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45X
		return
	case "45Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45Y
		return
	case "45Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45Z
		return
	case "46X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46X
		return
	case "46W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46W
		return
	case "46Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46Y
		return
	case "46Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46Z
		return
	case "47X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47X
		return
	case "47Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47Y
		return
	case "47Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47Z
		return
	case "48X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48X
		return
	case "48W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48W
		return
	case "48Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48Y
		return
	case "48Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48Z
		return
	case "49X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49X
		return
	case "49Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49Y
		return
	case "49Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49Z
		return
	case "50X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50X
		return
	case "50W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50W
		return
	case "50Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50Y
		return
	case "50Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50Z
		return
	case "51X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51X
		return
	case "51Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51Y
		return
	case "51Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51Z
		return
	case "52X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52X
		return
	case "52W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52W
		return
	case "52Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52Y
		return
	case "52Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52Z
		return
	case "53X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53X
		return
	case "53Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53Y
		return
	case "53Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53Z
		return
	case "54X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54X
		return
	case "54W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54W
		return
	case "54Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54Y
		return
	case "54Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54Z
		return
	case "55X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55X
		return
	case "55Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55Y
		return
	case "55Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55Z
		return
	case "56X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56X
		return
	case "56W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56W
		return
	case "56Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56Y
		return
	case "56Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56Z
		return
	case "57X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_57X
		return
	case "57Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_57Y
		return
	case "58X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_58X
		return
	case "58Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_58Y
		return
	case "59X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_59X
		return
	case "59Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_59Y
		return
	case "60X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_60X
		return
	case "60Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_60Y
		return
	case "61X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_61X
		return
	case "61Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_61Y
		return
	case "62X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_62X
		return
	case "62Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_62Y
		return
	case "63X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_63X
		return
	case "63Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_63Y
		return
	case "64X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_64X
		return
	case "64Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_64Y
		return
	case "65X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_65X
		return
	case "65Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_65Y
		return
	case "66X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_66X
		return
	case "66Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_66Y
		return
	case "67X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_67X
		return
	case "67Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_67Y
		return
	case "68X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_68X
		return
	case "68Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_68Y
		return
	case "69X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_69X
		return
	case "69Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_69Y
		return
	case "70X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_70X
		return
	case "70Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_70Y
		return
	case "71X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_71X
		return
	case "71Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_71Y
		return
	case "72X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_72X
		return
	case "72Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_72Y
		return
	case "73X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_73X
		return
	case "73Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_73Y
		return
	case "74X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_74X
		return
	case "74Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_74Y
		return
	case "75X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_75X
		return
	case "75Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_75Y
		return
	case "76X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_76X
		return
	case "76Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_76Y
		return
	case "77X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_77X
		return
	case "77Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_77Y
		return
	case "78X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_78X
		return
	case "78Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_78Y
		return
	case "79X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_79X
		return
	case "79Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_79Y
		return
	case "80X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80X
		return
	case "80Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80Y
		return
	case "80Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80Z
		return
	case "81X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81X
		return
	case "81Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81Y
		return
	case "81Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81Z
		return
	case "82X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82X
		return
	case "82Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82Y
		return
	case "82Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82Z
		return
	case "83X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83X
		return
	case "83Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83Y
		return
	case "83Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83Z
		return
	case "84X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84X
		return
	case "84Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84Y
		return
	case "84Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84Z
		return
	case "85X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85X
		return
	case "85Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85Y
		return
	case "85Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85Z
		return
	case "86X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86X
		return
	case "86Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86Y
		return
	case "86Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86Z
		return
	case "87X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87X
		return
	case "87Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87Y
		return
	case "87Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87Z
		return
	case "88X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88X
		return
	case "88Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88Y
		return
	case "88Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88Z
		return
	case "89X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89X
		return
	case "89Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89Y
		return
	case "89Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89Z
		return
	case "90X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90X
		return
	case "90Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90Y
		return
	case "90Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90Z
		return
	case "91X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91X
		return
	case "91Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91Y
		return
	case "91Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91Z
		return
	case "92X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92X
		return
	case "92Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92Y
		return
	case "92Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92Z
		return
	case "93X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93X
		return
	case "93Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93Y
		return
	case "93Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93Z
		return
	case "94X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94X
		return
	case "94Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94Y
		return
	case "94Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94Z
		return
	case "95X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95X
		return
	case "95Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95Y
		return
	case "95Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95Z
		return
	case "96X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96X
		return
	case "96Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96Y
		return
	case "96Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96Z
		return
	case "97X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97X
		return
	case "97Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97Y
		return
	case "97Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97Z
		return
	case "98X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98X
		return
	case "98Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98Y
		return
	case "98Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98Z
		return
	case "99X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99X
		return
	case "99Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99Y
		return
	case "99Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99Z
		return
	case "100X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100X
		return
	case "100Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100Y
		return
	case "100Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100Z
		return
	case "101X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101X
		return
	case "101Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101Y
		return
	case "101Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101Z
		return
	case "102X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102X
		return
	case "102Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102Y
		return
	case "102Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102Z
		return
	case "103X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103X
		return
	case "103Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103Y
		return
	case "103Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103Z
		return
	case "104X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104X
		return
	case "104Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104Y
		return
	case "104Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104Z
		return
	case "105X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105X
		return
	case "105Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105Y
		return
	case "105Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105Z
		return
	case "106X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106X
		return
	case "106Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106Y
		return
	case "106Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106Z
		return
	case "107X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107X
		return
	case "107Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107Y
		return
	case "107Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107Z
		return
	case "108X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108X
		return
	case "108Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108Y
		return
	case "108Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108Z
		return
	case "109X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109X
		return
	case "109Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109Y
		return
	case "109Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109Z
		return
	case "110X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110X
		return
	case "110Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110Y
		return
	case "110Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110Z
		return
	case "111X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111X
		return
	case "111Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111Y
		return
	case "111Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111Z
		return
	case "112X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112X
		return
	case "112Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112Y
		return
	case "112Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112Z
		return
	case "113X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113X
		return
	case "113Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113Y
		return
	case "113Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113Z
		return
	case "114X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114X
		return
	case "114Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114Y
		return
	case "114Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114Z
		return
	case "115X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115X
		return
	case "115Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115Y
		return
	case "115Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115Z
		return
	case "116X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116X
		return
	case "116Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116Y
		return
	case "116Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116Z
		return
	case "117X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117X
		return
	case "117Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117Y
		return
	case "117Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117Z
		return
	case "118X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118X
		return
	case "118Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118Y
		return
	case "118Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118Z
		return
	case "119X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119X
		return
	case "119Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119Y
		return
	case "119Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119Z
		return
	case "120X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_120X
		return
	case "120Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_120Y
		return
	case "121X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_121X
		return
	case "121Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_121Y
		return
	case "122X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_122X
		return
	case "122Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_122Y
		return
	case "123X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_123X
		return
	case "123Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_123Y
		return
	case "124X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_124X
		return
	case "124Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_124Y
		return
	case "125X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_125X
		return
	case "125Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_125Y
		return
	case "126X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_126X
		return
	case "126Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_126Y
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codetacanchannelbasetype_member_1 *Enum_CodeTACANChannelBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeTACANChannelBaseType_Member_1_N_1X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_1X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_1Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_1Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_2X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_2X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_2Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_2Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_3X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_3X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_3Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_3Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_4X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_4X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_4Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_4Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_5X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_5X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_5Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_5Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_6X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_6X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_6Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_6Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_7X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_7X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_7Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_7Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_8X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_8X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_8Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_8Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_9X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_9X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_9Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_9Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_10X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_10X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_10Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_10Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_11X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_11X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_11Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_11Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_12X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_12X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_12Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_12Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_13X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_13X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_13Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_13Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_14X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_14X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_14Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_14Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_15X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_15X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_15Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_15Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_16X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_16X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_16Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_16Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_17X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_17Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_17Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_17Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_18X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_18W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_18Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_18Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_18Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_19X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_19Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_19Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_19Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_20X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_20W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_20Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_20Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_20Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_21X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_21Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_21Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_21Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_22X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_22W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_22Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_22Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_22Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_23X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_23Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_23Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_23Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_24X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_24W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_24Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_24Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_24Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_25X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_25Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_25Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_25Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_26X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_26W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_26Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_26Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_26Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_27X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_27Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_27Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_27Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_28X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_28W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_28Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_28Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_28Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_29X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_29Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_29Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_29Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_30X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_30W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_30Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_30Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_30Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_31X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_31Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_31Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_31Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_32X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_32W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_32Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_32Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_32Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_33X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_33Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_33Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_33Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_34X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_34W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_34Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_34Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_34Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_35X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_35Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_35Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_35Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_36X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_36W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_36Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_36Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_36Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_37X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_37Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_37Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_37Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_38X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_38W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_38Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_38Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_38Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_39X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_39Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_39Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_39Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_40X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_40W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_40Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_40Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_40Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_41X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_41Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_41Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_41Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_42X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_42W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_42Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_42Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_42Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_43X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_43Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_43Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_43Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_44X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_44W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_44Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_44Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_44Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_45X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_45Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_45Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_45Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_46X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_46W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_46Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_46Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_46Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_47X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_47Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_47Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_47Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_48X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_48W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_48Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_48Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_48Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_49X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_49Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_49Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_49Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_50X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_50W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_50Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_50Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_50Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_51X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_51Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_51Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_51Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_52X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_52W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_52Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_52Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_52Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_53X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_53Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_53Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_53Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_54X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_54W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_54Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_54Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_54Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_55X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_55Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_55Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_55Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_56X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_56W":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56W
	case "Enum_CodeTACANChannelBaseType_Member_1_N_56Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_56Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_56Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_57X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_57X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_57Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_57Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_58X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_58X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_58Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_58Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_59X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_59X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_59Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_59Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_60X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_60X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_60Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_60Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_61X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_61X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_61Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_61Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_62X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_62X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_62Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_62Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_63X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_63X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_63Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_63Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_64X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_64X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_64Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_64Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_65X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_65X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_65Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_65Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_66X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_66X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_66Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_66Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_67X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_67X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_67Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_67Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_68X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_68X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_68Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_68Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_69X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_69X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_69Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_69Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_70X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_70X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_70Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_70Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_71X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_71X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_71Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_71Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_72X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_72X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_72Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_72Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_73X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_73X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_73Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_73Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_74X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_74X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_74Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_74Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_75X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_75X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_75Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_75Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_76X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_76X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_76Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_76Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_77X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_77X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_77Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_77Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_78X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_78X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_78Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_78Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_79X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_79X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_79Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_79Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_80X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_80Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_80Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_80Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_81X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_81Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_81Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_81Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_82X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_82Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_82Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_82Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_83X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_83Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_83Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_83Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_84X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_84Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_84Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_84Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_85X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_85Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_85Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_85Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_86X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_86Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_86Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_86Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_87X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_87Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_87Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_87Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_88X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_88Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_88Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_88Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_89X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_89Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_89Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_89Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_90X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_90Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_90Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_90Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_91X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_91Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_91Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_91Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_92X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_92Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_92Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_92Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_93X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_93Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_93Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_93Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_94X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_94Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_94Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_94Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_95X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_95Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_95Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_95Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_96X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_96Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_96Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_96Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_97X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_97Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_97Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_97Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_98X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_98Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_98Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_98Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_99X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_99Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_99Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_99Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_100X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_100Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_100Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_100Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_101X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_101Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_101Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_101Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_102X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_102Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_102Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_102Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_103X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_103Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_103Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_103Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_104X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_104Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_104Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_104Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_105X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_105Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_105Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_105Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_106X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_106Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_106Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_106Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_107X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_107Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_107Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_107Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_108X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_108Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_108Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_108Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_109X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_109Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_109Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_109Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_110X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_110Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_110Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_110Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_111X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_111Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_111Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_111Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_112X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_112Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_112Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_112Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_113X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_113Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_113Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_113Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_114X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_114Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_114Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_114Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_115X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_115Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_115Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_115Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_116X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_116Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_116Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_116Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_117X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_117Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_117Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_117Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_118X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_118Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_118Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_118Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_119X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_119Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_119Z":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_119Z
	case "Enum_CodeTACANChannelBaseType_Member_1_N_120X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_120X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_120Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_120Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_121X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_121X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_121Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_121Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_122X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_122X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_122Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_122Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_123X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_123X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_123Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_123Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_124X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_124X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_124Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_124Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_125X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_125X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_125Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_125Y
	case "Enum_CodeTACANChannelBaseType_Member_1_N_126X":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_126X
	case "Enum_CodeTACANChannelBaseType_Member_1_N_126Y":
		*enum_codetacanchannelbasetype_member_1 = Enum_CodeTACANChannelBaseType_Member_1_N_126Y
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codetacanchannelbasetype_member_1 *Enum_CodeTACANChannelBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codetacanchannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeTACANChannelBaseType_Member_1_N_1X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_1X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_1Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_1Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_2X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_2X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_2Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_2Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_3X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_3X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_3Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_3Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_4X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_4X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_4Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_4Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_5X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_5X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_5Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_5Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_6X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_6X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_6Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_6Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_7X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_7X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_7Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_7Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_8X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_8X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_8Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_8Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_9X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_9X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_9Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_9Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_10X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_10X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_10Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_10Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_11X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_11X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_11Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_11Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_12X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_12X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_12Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_12Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_13X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_13X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_13Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_13Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_14X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_14X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_14Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_14Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_15X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_15X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_15Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_15Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_16X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_16X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_16Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_16Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_17X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_17Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_17Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_17Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_18X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_18W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_18Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_18Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_18Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_19X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_19Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_19Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_19Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_20X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_20W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_20Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_20Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_20Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_21X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_21Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_21Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_21Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_22X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_22W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_22Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_22Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_22Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_23X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_23Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_23Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_23Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_24X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_24W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_24Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_24Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_24Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_25X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_25Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_25Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_25Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_26X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_26W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_26Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_26Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_26Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_27X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_27Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_27Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_27Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_28X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_28W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_28Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_28Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_28Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_29X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_29Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_29Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_29Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_30X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_30W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_30Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_30Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_30Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_31X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_31Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_31Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_31Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_32X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_32W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_32Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_32Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_32Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_33X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_33Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_33Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_33Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_34X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_34W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_34Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_34Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_34Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_35X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_35Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_35Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_35Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_36X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_36W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_36Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_36Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_36Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_37X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_37Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_37Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_37Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_38X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_38W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_38Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_38Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_38Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_39X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_39Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_39Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_39Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_40X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_40W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_40Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_40Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_40Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_41X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_41Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_41Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_41Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_42X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_42W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_42Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_42Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_42Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_43X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_43Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_43Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_43Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_44X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_44W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_44Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_44Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_44Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_45X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_45Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_45Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_45Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_46X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_46W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_46Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_46Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_46Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_47X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_47Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_47Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_47Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_48X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_48W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_48Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_48Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_48Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_49X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_49Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_49Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_49Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_50X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_50W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_50Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_50Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_50Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_51X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_51Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_51Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_51Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_52X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_52W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_52Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_52Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_52Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_53X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_53Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_53Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_53Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_54X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_54W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_54Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_54Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_54Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_55X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_55Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_55Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_55Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_56X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56W:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_56W"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_56Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_56Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_56Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_57X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_57X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_57Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_57Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_58X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_58X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_58Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_58Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_59X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_59X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_59Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_59Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_60X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_60X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_60Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_60Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_61X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_61X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_61Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_61Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_62X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_62X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_62Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_62Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_63X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_63X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_63Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_63Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_64X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_64X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_64Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_64Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_65X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_65X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_65Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_65Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_66X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_66X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_66Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_66Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_67X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_67X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_67Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_67Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_68X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_68X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_68Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_68Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_69X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_69X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_69Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_69Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_70X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_70X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_70Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_70Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_71X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_71X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_71Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_71Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_72X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_72X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_72Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_72Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_73X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_73X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_73Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_73Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_74X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_74X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_74Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_74Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_75X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_75X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_75Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_75Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_76X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_76X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_76Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_76Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_77X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_77X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_77Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_77Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_78X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_78X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_78Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_78Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_79X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_79X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_79Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_79Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_80X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_80Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_80Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_80Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_81X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_81Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_81Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_81Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_82X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_82Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_82Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_82Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_83X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_83Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_83Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_83Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_84X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_84Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_84Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_84Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_85X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_85Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_85Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_85Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_86X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_86Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_86Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_86Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_87X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_87Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_87Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_87Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_88X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_88Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_88Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_88Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_89X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_89Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_89Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_89Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_90X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_90Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_90Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_90Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_91X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_91Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_91Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_91Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_92X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_92Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_92Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_92Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_93X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_93Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_93Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_93Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_94X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_94Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_94Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_94Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_95X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_95Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_95Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_95Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_96X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_96Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_96Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_96Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_97X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_97Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_97Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_97Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_98X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_98Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_98Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_98Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_99X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_99Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_99Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_99Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_100X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_100Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_100Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_100Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_101X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_101Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_101Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_101Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_102X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_102Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_102Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_102Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_103X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_103Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_103Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_103Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_104X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_104Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_104Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_104Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_105X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_105Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_105Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_105Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_106X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_106Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_106Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_106Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_107X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_107Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_107Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_107Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_108X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_108Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_108Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_108Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_109X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_109Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_109Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_109Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_110X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_110Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_110Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_110Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_111X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_111Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_111Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_111Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_112X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_112Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_112Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_112Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_113X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_113Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_113Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_113Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_114X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_114Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_114Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_114Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_115X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_115Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_115Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_115Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_116X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_116Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_116Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_116Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_117X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_117Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_117Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_117Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_118X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_118Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_118Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_118Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_119X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_119Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_119Z:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_119Z"
	case Enum_CodeTACANChannelBaseType_Member_1_N_120X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_120X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_120Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_120Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_121X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_121X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_121Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_121Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_122X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_122X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_122Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_122Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_123X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_123X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_123Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_123Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_124X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_124X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_124Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_124Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_125X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_125X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_125Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_125Y"
	case Enum_CodeTACANChannelBaseType_Member_1_N_126X:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_126X"
	case Enum_CodeTACANChannelBaseType_Member_1_N_126Y:
		res = "Enum_CodeTACANChannelBaseType_Member_1_N_126Y"
	}
	return
}

func (enum_codetacanchannelbasetype_member_1 Enum_CodeTACANChannelBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_1X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_1Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_2X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_2Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_3X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_3Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_4X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_4Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_5X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_5Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_6X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_6Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_7X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_7Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_8X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_8Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_9X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_9Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_10X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_10Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_11X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_11Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_12X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_12Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_13X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_13Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_14X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_14Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_15X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_15Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_16X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_16Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_17X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_17Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_17Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_18X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_18W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_18Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_18Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_19X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_19Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_19Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_20X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_20W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_20Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_20Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_21X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_21Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_21Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_22X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_22W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_22Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_22Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_23X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_23Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_23Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_24X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_24W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_24Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_24Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_25X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_25Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_25Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_26X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_26W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_26Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_26Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_27X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_27Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_27Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_28X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_28W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_28Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_28Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_29X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_29Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_29Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_30X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_30W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_30Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_30Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_31X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_31Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_31Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_32X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_32W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_32Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_32Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_33X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_33Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_33Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_34X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_34W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_34Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_34Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_35X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_35Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_35Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_36X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_36W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_36Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_36Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_37X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_37Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_37Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_38X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_38W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_38Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_38Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_39X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_39Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_39Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_40X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_40W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_40Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_40Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_41X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_41Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_41Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_42X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_42W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_42Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_42Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_43X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_43Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_43Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_44X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_44W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_44Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_44Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_45X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_45Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_45Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_46X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_46W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_46Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_46Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_47X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_47Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_47Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_48X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_48W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_48Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_48Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_49X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_49Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_49Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_50X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_50W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_50Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_50Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_51X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_51Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_51Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_52X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_52W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_52Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_52Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_53X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_53Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_53Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_54X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_54W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_54Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_54Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_55X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_55Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_55Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_56X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_56W")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_56Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_56Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_57X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_57Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_58X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_58Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_59X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_59Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_60X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_60Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_61X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_61Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_62X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_62Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_63X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_63Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_64X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_64Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_65X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_65Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_66X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_66Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_67X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_67Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_68X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_68Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_69X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_69Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_70X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_70Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_71X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_71Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_72X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_72Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_73X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_73Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_74X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_74Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_75X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_75Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_76X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_76Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_77X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_77Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_78X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_78Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_79X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_79Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_80X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_80Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_80Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_81X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_81Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_81Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_82X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_82Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_82Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_83X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_83Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_83Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_84X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_84Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_84Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_85X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_85Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_85Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_86X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_86Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_86Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_87X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_87Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_87Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_88X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_88Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_88Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_89X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_89Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_89Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_90X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_90Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_90Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_91X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_91Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_91Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_92X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_92Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_92Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_93X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_93Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_93Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_94X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_94Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_94Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_95X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_95Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_95Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_96X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_96Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_96Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_97X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_97Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_97Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_98X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_98Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_98Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_99X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_99Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_99Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_100X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_100Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_100Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_101X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_101Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_101Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_102X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_102Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_102Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_103X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_103Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_103Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_104X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_104Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_104Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_105X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_105Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_105Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_106X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_106Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_106Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_107X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_107Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_107Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_108X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_108Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_108Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_109X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_109Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_109Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_110X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_110Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_110Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_111X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_111Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_111Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_112X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_112Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_112Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_113X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_113Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_113Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_114X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_114Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_114Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_115X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_115Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_115Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_116X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_116Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_116Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_117X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_117Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_117Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_118X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_118Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_118Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_119X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_119Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_119Z")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_120X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_120Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_121X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_121Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_122X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_122Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_123X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_123Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_124X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_124Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_125X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_125Y")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_126X")
	res = append(res, "Enum_CodeTACANChannelBaseType_Member_1_N_126Y")

	return
}

func (enum_codetacanchannelbasetype_member_1 Enum_CodeTACANChannelBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1X")
	res = append(res, "1Y")
	res = append(res, "2X")
	res = append(res, "2Y")
	res = append(res, "3X")
	res = append(res, "3Y")
	res = append(res, "4X")
	res = append(res, "4Y")
	res = append(res, "5X")
	res = append(res, "5Y")
	res = append(res, "6X")
	res = append(res, "6Y")
	res = append(res, "7X")
	res = append(res, "7Y")
	res = append(res, "8X")
	res = append(res, "8Y")
	res = append(res, "9X")
	res = append(res, "9Y")
	res = append(res, "10X")
	res = append(res, "10Y")
	res = append(res, "11X")
	res = append(res, "11Y")
	res = append(res, "12X")
	res = append(res, "12Y")
	res = append(res, "13X")
	res = append(res, "13Y")
	res = append(res, "14X")
	res = append(res, "14Y")
	res = append(res, "15X")
	res = append(res, "15Y")
	res = append(res, "16X")
	res = append(res, "16Y")
	res = append(res, "17X")
	res = append(res, "17Y")
	res = append(res, "17Z")
	res = append(res, "18X")
	res = append(res, "18W")
	res = append(res, "18Y")
	res = append(res, "18Z")
	res = append(res, "19X")
	res = append(res, "19Y")
	res = append(res, "19Z")
	res = append(res, "20X")
	res = append(res, "20W")
	res = append(res, "20Y")
	res = append(res, "20Z")
	res = append(res, "21X")
	res = append(res, "21Y")
	res = append(res, "21Z")
	res = append(res, "22X")
	res = append(res, "22W")
	res = append(res, "22Y")
	res = append(res, "22Z")
	res = append(res, "23X")
	res = append(res, "23Y")
	res = append(res, "23Z")
	res = append(res, "24X")
	res = append(res, "24W")
	res = append(res, "24Y")
	res = append(res, "24Z")
	res = append(res, "25X")
	res = append(res, "25Y")
	res = append(res, "25Z")
	res = append(res, "26X")
	res = append(res, "26W")
	res = append(res, "26Y")
	res = append(res, "26Z")
	res = append(res, "27X")
	res = append(res, "27Y")
	res = append(res, "27Z")
	res = append(res, "28X")
	res = append(res, "28W")
	res = append(res, "28Y")
	res = append(res, "28Z")
	res = append(res, "29X")
	res = append(res, "29Y")
	res = append(res, "29Z")
	res = append(res, "30X")
	res = append(res, "30W")
	res = append(res, "30Y")
	res = append(res, "30Z")
	res = append(res, "31X")
	res = append(res, "31Y")
	res = append(res, "31Z")
	res = append(res, "32X")
	res = append(res, "32W")
	res = append(res, "32Y")
	res = append(res, "32Z")
	res = append(res, "33X")
	res = append(res, "33Y")
	res = append(res, "33Z")
	res = append(res, "34X")
	res = append(res, "34W")
	res = append(res, "34Y")
	res = append(res, "34Z")
	res = append(res, "35X")
	res = append(res, "35Y")
	res = append(res, "35Z")
	res = append(res, "36X")
	res = append(res, "36W")
	res = append(res, "36Y")
	res = append(res, "36Z")
	res = append(res, "37X")
	res = append(res, "37Y")
	res = append(res, "37Z")
	res = append(res, "38X")
	res = append(res, "38W")
	res = append(res, "38Y")
	res = append(res, "38Z")
	res = append(res, "39X")
	res = append(res, "39Y")
	res = append(res, "39Z")
	res = append(res, "40X")
	res = append(res, "40W")
	res = append(res, "40Y")
	res = append(res, "40Z")
	res = append(res, "41X")
	res = append(res, "41Y")
	res = append(res, "41Z")
	res = append(res, "42X")
	res = append(res, "42W")
	res = append(res, "42Y")
	res = append(res, "42Z")
	res = append(res, "43X")
	res = append(res, "43Y")
	res = append(res, "43Z")
	res = append(res, "44X")
	res = append(res, "44W")
	res = append(res, "44Y")
	res = append(res, "44Z")
	res = append(res, "45X")
	res = append(res, "45Y")
	res = append(res, "45Z")
	res = append(res, "46X")
	res = append(res, "46W")
	res = append(res, "46Y")
	res = append(res, "46Z")
	res = append(res, "47X")
	res = append(res, "47Y")
	res = append(res, "47Z")
	res = append(res, "48X")
	res = append(res, "48W")
	res = append(res, "48Y")
	res = append(res, "48Z")
	res = append(res, "49X")
	res = append(res, "49Y")
	res = append(res, "49Z")
	res = append(res, "50X")
	res = append(res, "50W")
	res = append(res, "50Y")
	res = append(res, "50Z")
	res = append(res, "51X")
	res = append(res, "51Y")
	res = append(res, "51Z")
	res = append(res, "52X")
	res = append(res, "52W")
	res = append(res, "52Y")
	res = append(res, "52Z")
	res = append(res, "53X")
	res = append(res, "53Y")
	res = append(res, "53Z")
	res = append(res, "54X")
	res = append(res, "54W")
	res = append(res, "54Y")
	res = append(res, "54Z")
	res = append(res, "55X")
	res = append(res, "55Y")
	res = append(res, "55Z")
	res = append(res, "56X")
	res = append(res, "56W")
	res = append(res, "56Y")
	res = append(res, "56Z")
	res = append(res, "57X")
	res = append(res, "57Y")
	res = append(res, "58X")
	res = append(res, "58Y")
	res = append(res, "59X")
	res = append(res, "59Y")
	res = append(res, "60X")
	res = append(res, "60Y")
	res = append(res, "61X")
	res = append(res, "61Y")
	res = append(res, "62X")
	res = append(res, "62Y")
	res = append(res, "63X")
	res = append(res, "63Y")
	res = append(res, "64X")
	res = append(res, "64Y")
	res = append(res, "65X")
	res = append(res, "65Y")
	res = append(res, "66X")
	res = append(res, "66Y")
	res = append(res, "67X")
	res = append(res, "67Y")
	res = append(res, "68X")
	res = append(res, "68Y")
	res = append(res, "69X")
	res = append(res, "69Y")
	res = append(res, "70X")
	res = append(res, "70Y")
	res = append(res, "71X")
	res = append(res, "71Y")
	res = append(res, "72X")
	res = append(res, "72Y")
	res = append(res, "73X")
	res = append(res, "73Y")
	res = append(res, "74X")
	res = append(res, "74Y")
	res = append(res, "75X")
	res = append(res, "75Y")
	res = append(res, "76X")
	res = append(res, "76Y")
	res = append(res, "77X")
	res = append(res, "77Y")
	res = append(res, "78X")
	res = append(res, "78Y")
	res = append(res, "79X")
	res = append(res, "79Y")
	res = append(res, "80X")
	res = append(res, "80Y")
	res = append(res, "80Z")
	res = append(res, "81X")
	res = append(res, "81Y")
	res = append(res, "81Z")
	res = append(res, "82X")
	res = append(res, "82Y")
	res = append(res, "82Z")
	res = append(res, "83X")
	res = append(res, "83Y")
	res = append(res, "83Z")
	res = append(res, "84X")
	res = append(res, "84Y")
	res = append(res, "84Z")
	res = append(res, "85X")
	res = append(res, "85Y")
	res = append(res, "85Z")
	res = append(res, "86X")
	res = append(res, "86Y")
	res = append(res, "86Z")
	res = append(res, "87X")
	res = append(res, "87Y")
	res = append(res, "87Z")
	res = append(res, "88X")
	res = append(res, "88Y")
	res = append(res, "88Z")
	res = append(res, "89X")
	res = append(res, "89Y")
	res = append(res, "89Z")
	res = append(res, "90X")
	res = append(res, "90Y")
	res = append(res, "90Z")
	res = append(res, "91X")
	res = append(res, "91Y")
	res = append(res, "91Z")
	res = append(res, "92X")
	res = append(res, "92Y")
	res = append(res, "92Z")
	res = append(res, "93X")
	res = append(res, "93Y")
	res = append(res, "93Z")
	res = append(res, "94X")
	res = append(res, "94Y")
	res = append(res, "94Z")
	res = append(res, "95X")
	res = append(res, "95Y")
	res = append(res, "95Z")
	res = append(res, "96X")
	res = append(res, "96Y")
	res = append(res, "96Z")
	res = append(res, "97X")
	res = append(res, "97Y")
	res = append(res, "97Z")
	res = append(res, "98X")
	res = append(res, "98Y")
	res = append(res, "98Z")
	res = append(res, "99X")
	res = append(res, "99Y")
	res = append(res, "99Z")
	res = append(res, "100X")
	res = append(res, "100Y")
	res = append(res, "100Z")
	res = append(res, "101X")
	res = append(res, "101Y")
	res = append(res, "101Z")
	res = append(res, "102X")
	res = append(res, "102Y")
	res = append(res, "102Z")
	res = append(res, "103X")
	res = append(res, "103Y")
	res = append(res, "103Z")
	res = append(res, "104X")
	res = append(res, "104Y")
	res = append(res, "104Z")
	res = append(res, "105X")
	res = append(res, "105Y")
	res = append(res, "105Z")
	res = append(res, "106X")
	res = append(res, "106Y")
	res = append(res, "106Z")
	res = append(res, "107X")
	res = append(res, "107Y")
	res = append(res, "107Z")
	res = append(res, "108X")
	res = append(res, "108Y")
	res = append(res, "108Z")
	res = append(res, "109X")
	res = append(res, "109Y")
	res = append(res, "109Z")
	res = append(res, "110X")
	res = append(res, "110Y")
	res = append(res, "110Z")
	res = append(res, "111X")
	res = append(res, "111Y")
	res = append(res, "111Z")
	res = append(res, "112X")
	res = append(res, "112Y")
	res = append(res, "112Z")
	res = append(res, "113X")
	res = append(res, "113Y")
	res = append(res, "113Z")
	res = append(res, "114X")
	res = append(res, "114Y")
	res = append(res, "114Z")
	res = append(res, "115X")
	res = append(res, "115Y")
	res = append(res, "115Z")
	res = append(res, "116X")
	res = append(res, "116Y")
	res = append(res, "116Z")
	res = append(res, "117X")
	res = append(res, "117Y")
	res = append(res, "117Z")
	res = append(res, "118X")
	res = append(res, "118Y")
	res = append(res, "118Z")
	res = append(res, "119X")
	res = append(res, "119Y")
	res = append(res, "119Z")
	res = append(res, "120X")
	res = append(res, "120Y")
	res = append(res, "121X")
	res = append(res, "121Y")
	res = append(res, "122X")
	res = append(res, "122Y")
	res = append(res, "123X")
	res = append(res, "123Y")
	res = append(res, "124X")
	res = append(res, "124Y")
	res = append(res, "125X")
	res = append(res, "125Y")
	res = append(res, "126X")
	res = append(res, "126Y")

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
