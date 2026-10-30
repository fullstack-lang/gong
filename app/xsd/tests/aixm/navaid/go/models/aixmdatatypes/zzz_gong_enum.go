// generated code - do not edit
package aixmdatatypes

// insertion point of enum utility functions
// Utility function for Enum_CodeAirspacePointPositionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspacepointpositionbasetype_member_1 Enum_CodeAirspacePointPositionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspacepointpositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspacePointPositionBaseType_Member_1_IN:
		res = "IN"
	case Enum_CodeAirspacePointPositionBaseType_Member_1_OUT:
		res = "OUT"
	case Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER:
		res = "BORDER"
	}
	return
}

func (enum_codeairspacepointpositionbasetype_member_1 *Enum_CodeAirspacePointPositionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IN":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_IN
		return
	case "OUT":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_OUT
		return
	case "BORDER":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspacepointpositionbasetype_member_1 *Enum_CodeAirspacePointPositionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspacePointPositionBaseType_Member_1_IN":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_IN
	case "Enum_CodeAirspacePointPositionBaseType_Member_1_OUT":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_OUT
	case "Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER":
		*enum_codeairspacepointpositionbasetype_member_1 = Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspacepointpositionbasetype_member_1 *Enum_CodeAirspacePointPositionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspacepointpositionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspacePointPositionBaseType_Member_1_IN:
		res = "Enum_CodeAirspacePointPositionBaseType_Member_1_IN"
	case Enum_CodeAirspacePointPositionBaseType_Member_1_OUT:
		res = "Enum_CodeAirspacePointPositionBaseType_Member_1_OUT"
	case Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER:
		res = "Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER"
	}
	return
}

func (enum_codeairspacepointpositionbasetype_member_1 Enum_CodeAirspacePointPositionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspacePointPositionBaseType_Member_1_IN")
	res = append(res, "Enum_CodeAirspacePointPositionBaseType_Member_1_OUT")
	res = append(res, "Enum_CodeAirspacePointPositionBaseType_Member_1_BORDER")

	return
}

func (enum_codeairspacepointpositionbasetype_member_1 Enum_CodeAirspacePointPositionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IN")
	res = append(res, "OUT")
	res = append(res, "BORDER")

	return
}

// Utility function for Enum_CodeAirspacePointRoleBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeairspacepointrolebasetype_member_1 Enum_CodeAirspacePointRoleBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeairspacepointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY:
		res = "ENTRY"
	case Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT:
		res = "EXIT"
	case Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT:
		res = "ENTRY_EXIT"
	}
	return
}

func (enum_codeairspacepointrolebasetype_member_1 *Enum_CodeAirspacePointRoleBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENTRY":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY
		return
	case "EXIT":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT
		return
	case "ENTRY_EXIT":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeairspacepointrolebasetype_member_1 *Enum_CodeAirspacePointRoleBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY
	case "Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT
	case "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT":
		*enum_codeairspacepointrolebasetype_member_1 = Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeairspacepointrolebasetype_member_1 *Enum_CodeAirspacePointRoleBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeairspacepointrolebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY:
		res = "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY"
	case Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT:
		res = "Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT"
	case Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT:
		res = "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT"
	}
	return
}

func (enum_codeairspacepointrolebasetype_member_1 Enum_CodeAirspacePointRoleBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY")
	res = append(res, "Enum_CodeAirspacePointRoleBaseType_Member_1_EXIT")
	res = append(res, "Enum_CodeAirspacePointRoleBaseType_Member_1_ENTRY_EXIT")

	return
}

func (enum_codeairspacepointrolebasetype_member_1 Enum_CodeAirspacePointRoleBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENTRY")
	res = append(res, "EXIT")
	res = append(res, "ENTRY_EXIT")

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

// Utility function for Enum_CodeCheckpointCategoryBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecheckpointcategorybasetype_member_1 Enum_CodeCheckpointCategoryBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecheckpointcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCheckpointCategoryBaseType_Member_1_A:
		res = "A"
	case Enum_CodeCheckpointCategoryBaseType_Member_1_G:
		res = "G"
	}
	return
}

func (enum_codecheckpointcategorybasetype_member_1 *Enum_CodeCheckpointCategoryBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codecheckpointcategorybasetype_member_1 = Enum_CodeCheckpointCategoryBaseType_Member_1_A
		return
	case "G":
		*enum_codecheckpointcategorybasetype_member_1 = Enum_CodeCheckpointCategoryBaseType_Member_1_G
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecheckpointcategorybasetype_member_1 *Enum_CodeCheckpointCategoryBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCheckpointCategoryBaseType_Member_1_A":
		*enum_codecheckpointcategorybasetype_member_1 = Enum_CodeCheckpointCategoryBaseType_Member_1_A
	case "Enum_CodeCheckpointCategoryBaseType_Member_1_G":
		*enum_codecheckpointcategorybasetype_member_1 = Enum_CodeCheckpointCategoryBaseType_Member_1_G
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecheckpointcategorybasetype_member_1 *Enum_CodeCheckpointCategoryBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecheckpointcategorybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCheckpointCategoryBaseType_Member_1_A:
		res = "Enum_CodeCheckpointCategoryBaseType_Member_1_A"
	case Enum_CodeCheckpointCategoryBaseType_Member_1_G:
		res = "Enum_CodeCheckpointCategoryBaseType_Member_1_G"
	}
	return
}

func (enum_codecheckpointcategorybasetype_member_1 Enum_CodeCheckpointCategoryBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCheckpointCategoryBaseType_Member_1_A")
	res = append(res, "Enum_CodeCheckpointCategoryBaseType_Member_1_G")

	return
}

func (enum_codecheckpointcategorybasetype_member_1 Enum_CodeCheckpointCategoryBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "G")

	return
}

// Utility function for Enum_CodeCourseQualityILSBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecoursequalityilsbasetype_member_1 Enum_CodeCourseQualityILSBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecoursequalityilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCourseQualityILSBaseType_Member_1_A:
		res = "A"
	case Enum_CodeCourseQualityILSBaseType_Member_1_B:
		res = "B"
	case Enum_CodeCourseQualityILSBaseType_Member_1_C:
		res = "C"
	case Enum_CodeCourseQualityILSBaseType_Member_1_D:
		res = "D"
	case Enum_CodeCourseQualityILSBaseType_Member_1_E:
		res = "E"
	case Enum_CodeCourseQualityILSBaseType_Member_1_T:
		res = "T"
	}
	return
}

func (enum_codecoursequalityilsbasetype_member_1 *Enum_CodeCourseQualityILSBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_A
		return
	case "B":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_B
		return
	case "C":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_C
		return
	case "D":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_D
		return
	case "E":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_E
		return
	case "T":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_T
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecoursequalityilsbasetype_member_1 *Enum_CodeCourseQualityILSBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCourseQualityILSBaseType_Member_1_A":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_A
	case "Enum_CodeCourseQualityILSBaseType_Member_1_B":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_B
	case "Enum_CodeCourseQualityILSBaseType_Member_1_C":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_C
	case "Enum_CodeCourseQualityILSBaseType_Member_1_D":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_D
	case "Enum_CodeCourseQualityILSBaseType_Member_1_E":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_E
	case "Enum_CodeCourseQualityILSBaseType_Member_1_T":
		*enum_codecoursequalityilsbasetype_member_1 = Enum_CodeCourseQualityILSBaseType_Member_1_T
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecoursequalityilsbasetype_member_1 *Enum_CodeCourseQualityILSBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecoursequalityilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCourseQualityILSBaseType_Member_1_A:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_A"
	case Enum_CodeCourseQualityILSBaseType_Member_1_B:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_B"
	case Enum_CodeCourseQualityILSBaseType_Member_1_C:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_C"
	case Enum_CodeCourseQualityILSBaseType_Member_1_D:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_D"
	case Enum_CodeCourseQualityILSBaseType_Member_1_E:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_E"
	case Enum_CodeCourseQualityILSBaseType_Member_1_T:
		res = "Enum_CodeCourseQualityILSBaseType_Member_1_T"
	}
	return
}

func (enum_codecoursequalityilsbasetype_member_1 Enum_CodeCourseQualityILSBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_A")
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_B")
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_C")
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_D")
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_E")
	res = append(res, "Enum_CodeCourseQualityILSBaseType_Member_1_T")

	return
}

func (enum_codecoursequalityilsbasetype_member_1 Enum_CodeCourseQualityILSBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")
	res = append(res, "E")
	res = append(res, "T")

	return
}

// Utility function for Enum_CodeDMEBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedmebasetype_member_1 Enum_CodeDMEBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedmebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDMEBaseType_Member_1_NARROW:
		res = "NARROW"
	case Enum_CodeDMEBaseType_Member_1_PRECISION:
		res = "PRECISION"
	case Enum_CodeDMEBaseType_Member_1_WIDE:
		res = "WIDE"
	}
	return
}

func (enum_codedmebasetype_member_1 *Enum_CodeDMEBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NARROW":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_NARROW
		return
	case "PRECISION":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_PRECISION
		return
	case "WIDE":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_WIDE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedmebasetype_member_1 *Enum_CodeDMEBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDMEBaseType_Member_1_NARROW":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_NARROW
	case "Enum_CodeDMEBaseType_Member_1_PRECISION":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_PRECISION
	case "Enum_CodeDMEBaseType_Member_1_WIDE":
		*enum_codedmebasetype_member_1 = Enum_CodeDMEBaseType_Member_1_WIDE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedmebasetype_member_1 *Enum_CodeDMEBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedmebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDMEBaseType_Member_1_NARROW:
		res = "Enum_CodeDMEBaseType_Member_1_NARROW"
	case Enum_CodeDMEBaseType_Member_1_PRECISION:
		res = "Enum_CodeDMEBaseType_Member_1_PRECISION"
	case Enum_CodeDMEBaseType_Member_1_WIDE:
		res = "Enum_CodeDMEBaseType_Member_1_WIDE"
	}
	return
}

func (enum_codedmebasetype_member_1 Enum_CodeDMEBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDMEBaseType_Member_1_NARROW")
	res = append(res, "Enum_CodeDMEBaseType_Member_1_PRECISION")
	res = append(res, "Enum_CodeDMEBaseType_Member_1_WIDE")

	return
}

func (enum_codedmebasetype_member_1 Enum_CodeDMEBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NARROW")
	res = append(res, "PRECISION")
	res = append(res, "WIDE")

	return
}

// Utility function for Enum_CodeDMEChannelBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedmechannelbasetype_member_1 Enum_CodeDMEChannelBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedmechannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDMEChannelBaseType_Member_1_N_1X:
		res = "1X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_1Y:
		res = "1Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_2X:
		res = "2X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_2Y:
		res = "2Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_3X:
		res = "3X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_3Y:
		res = "3Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_4X:
		res = "4X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_4Y:
		res = "4Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_5X:
		res = "5X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_5Y:
		res = "5Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_6X:
		res = "6X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_6Y:
		res = "6Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_7X:
		res = "7X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_7Y:
		res = "7Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_8X:
		res = "8X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_8Y:
		res = "8Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_9X:
		res = "9X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_9Y:
		res = "9Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_10X:
		res = "10X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_10Y:
		res = "10Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_11X:
		res = "11X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_11Y:
		res = "11Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_12X:
		res = "12X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_12Y:
		res = "12Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_13X:
		res = "13X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_13Y:
		res = "13Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_14X:
		res = "14X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_14Y:
		res = "14Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_15X:
		res = "15X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_15Y:
		res = "15Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_16X:
		res = "16X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_16Y:
		res = "16Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17X:
		res = "17X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17Y:
		res = "17Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17Z:
		res = "17Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18X:
		res = "18X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18W:
		res = "18W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18Y:
		res = "18Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18Z:
		res = "18Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19X:
		res = "19X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19Y:
		res = "19Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19Z:
		res = "19Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20X:
		res = "20X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20W:
		res = "20W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20Y:
		res = "20Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20Z:
		res = "20Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21X:
		res = "21X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21Y:
		res = "21Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21Z:
		res = "21Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22X:
		res = "22X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22W:
		res = "22W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22Y:
		res = "22Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22Z:
		res = "22Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23X:
		res = "23X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23Y:
		res = "23Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23Z:
		res = "23Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24X:
		res = "24X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24W:
		res = "24W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24Y:
		res = "24Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24Z:
		res = "24Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25X:
		res = "25X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25Y:
		res = "25Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25Z:
		res = "25Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26X:
		res = "26X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26W:
		res = "26W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26Y:
		res = "26Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26Z:
		res = "26Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27X:
		res = "27X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27Y:
		res = "27Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27Z:
		res = "27Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28X:
		res = "28X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28W:
		res = "28W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28Y:
		res = "28Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28Z:
		res = "28Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29X:
		res = "29X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29Y:
		res = "29Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29Z:
		res = "29Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30X:
		res = "30X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30W:
		res = "30W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30Y:
		res = "30Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30Z:
		res = "30Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31X:
		res = "31X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31Y:
		res = "31Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31Z:
		res = "31Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32X:
		res = "32X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32W:
		res = "32W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32Y:
		res = "32Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32Z:
		res = "32Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33X:
		res = "33X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33Y:
		res = "33Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33Z:
		res = "33Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34X:
		res = "34X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34W:
		res = "34W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34Y:
		res = "34Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34Z:
		res = "34Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35X:
		res = "35X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35Y:
		res = "35Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35Z:
		res = "35Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36X:
		res = "36X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36W:
		res = "36W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36Y:
		res = "36Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36Z:
		res = "36Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37X:
		res = "37X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37Y:
		res = "37Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37Z:
		res = "37Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38X:
		res = "38X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38W:
		res = "38W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38Y:
		res = "38Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38Z:
		res = "38Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39X:
		res = "39X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39Y:
		res = "39Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39Z:
		res = "39Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40X:
		res = "40X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40W:
		res = "40W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40Y:
		res = "40Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40Z:
		res = "40Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41X:
		res = "41X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41Y:
		res = "41Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41Z:
		res = "41Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42X:
		res = "42X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42W:
		res = "42W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42Y:
		res = "42Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42Z:
		res = "42Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43X:
		res = "43X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43Y:
		res = "43Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43Z:
		res = "43Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44X:
		res = "44X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44W:
		res = "44W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44Y:
		res = "44Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44Z:
		res = "44Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45X:
		res = "45X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45Y:
		res = "45Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45Z:
		res = "45Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46X:
		res = "46X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46W:
		res = "46W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46Y:
		res = "46Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46Z:
		res = "46Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47X:
		res = "47X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47Y:
		res = "47Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47Z:
		res = "47Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48X:
		res = "48X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48W:
		res = "48W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48Y:
		res = "48Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48Z:
		res = "48Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49X:
		res = "49X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49Y:
		res = "49Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49Z:
		res = "49Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50X:
		res = "50X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50W:
		res = "50W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50Y:
		res = "50Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50Z:
		res = "50Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51X:
		res = "51X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51Y:
		res = "51Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51Z:
		res = "51Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52X:
		res = "52X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52W:
		res = "52W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52Y:
		res = "52Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52Z:
		res = "52Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53X:
		res = "53X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53Y:
		res = "53Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53Z:
		res = "53Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54X:
		res = "54X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54W:
		res = "54W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54Y:
		res = "54Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54Z:
		res = "54Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55X:
		res = "55X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55Y:
		res = "55Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55Z:
		res = "55Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56X:
		res = "56X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56W:
		res = "56W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56Y:
		res = "56Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56Z:
		res = "56Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_57X:
		res = "57X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_57Y:
		res = "57Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_58X:
		res = "58X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_58Y:
		res = "58Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_59X:
		res = "59X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_59Y:
		res = "59Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_60X:
		res = "60X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_60Y:
		res = "60Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_61X:
		res = "61X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_61Y:
		res = "61Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_62X:
		res = "62X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_62Y:
		res = "62Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_63X:
		res = "63X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_63Y:
		res = "63Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_64X:
		res = "64X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_64Y:
		res = "64Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_65X:
		res = "65X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_65Y:
		res = "65Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_66X:
		res = "66X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_66Y:
		res = "66Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_67X:
		res = "67X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_67Y:
		res = "67Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_68X:
		res = "68X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_68Y:
		res = "68Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_69X:
		res = "69X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_69Y:
		res = "69Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_70X:
		res = "70X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_70Y:
		res = "70Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_71X:
		res = "71X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_71Y:
		res = "71Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_72X:
		res = "72X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_72Y:
		res = "72Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_73X:
		res = "73X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_73Y:
		res = "73Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_74X:
		res = "74X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_74Y:
		res = "74Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_75X:
		res = "75X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_75Y:
		res = "75Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_76X:
		res = "76X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_76Y:
		res = "76Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_77X:
		res = "77X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_77Y:
		res = "77Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_78X:
		res = "78X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_78Y:
		res = "78Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_79X:
		res = "79X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_79Y:
		res = "79Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80X:
		res = "80X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80Y:
		res = "80Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80Z:
		res = "80Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81X:
		res = "81X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81Y:
		res = "81Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81Z:
		res = "81Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82X:
		res = "82X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82Y:
		res = "82Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82Z:
		res = "82Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83X:
		res = "83X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83Y:
		res = "83Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83Z:
		res = "83Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84X:
		res = "84X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84Y:
		res = "84Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84Z:
		res = "84Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85X:
		res = "85X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85Y:
		res = "85Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85Z:
		res = "85Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86X:
		res = "86X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86Y:
		res = "86Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86Z:
		res = "86Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87X:
		res = "87X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87Y:
		res = "87Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87Z:
		res = "87Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88X:
		res = "88X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88Y:
		res = "88Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88Z:
		res = "88Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89X:
		res = "89X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89Y:
		res = "89Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89Z:
		res = "89Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90X:
		res = "90X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90Y:
		res = "90Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90Z:
		res = "90Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91X:
		res = "91X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91Y:
		res = "91Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91Z:
		res = "91Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92X:
		res = "92X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92Y:
		res = "92Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92Z:
		res = "92Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93X:
		res = "93X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93Y:
		res = "93Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93Z:
		res = "93Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94X:
		res = "94X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94Y:
		res = "94Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94Z:
		res = "94Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95X:
		res = "95X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95Y:
		res = "95Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95Z:
		res = "95Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96X:
		res = "96X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96Y:
		res = "96Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96Z:
		res = "96Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97X:
		res = "97X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97Y:
		res = "97Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97Z:
		res = "97Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98X:
		res = "98X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98Y:
		res = "98Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98Z:
		res = "98Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99X:
		res = "99X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99Y:
		res = "99Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99Z:
		res = "99Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100X:
		res = "100X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100Y:
		res = "100Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100Z:
		res = "100Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101X:
		res = "101X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101Y:
		res = "101Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101Z:
		res = "101Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102X:
		res = "102X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102Y:
		res = "102Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102Z:
		res = "102Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103X:
		res = "103X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103Y:
		res = "103Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103Z:
		res = "103Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104X:
		res = "104X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104Y:
		res = "104Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104Z:
		res = "104Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105X:
		res = "105X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105Y:
		res = "105Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105Z:
		res = "105Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106X:
		res = "106X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106Y:
		res = "106Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106Z:
		res = "106Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107X:
		res = "107X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107Y:
		res = "107Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107Z:
		res = "107Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108X:
		res = "108X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108Y:
		res = "108Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108Z:
		res = "108Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109X:
		res = "109X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109Y:
		res = "109Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109Z:
		res = "109Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110X:
		res = "110X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110Y:
		res = "110Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110Z:
		res = "110Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111X:
		res = "111X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111Y:
		res = "111Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111Z:
		res = "111Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112X:
		res = "112X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112Y:
		res = "112Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112Z:
		res = "112Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113X:
		res = "113X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113Y:
		res = "113Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113Z:
		res = "113Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114X:
		res = "114X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114Y:
		res = "114Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114Z:
		res = "114Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115X:
		res = "115X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115Y:
		res = "115Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115Z:
		res = "115Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116X:
		res = "116X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116Y:
		res = "116Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116Z:
		res = "116Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117X:
		res = "117X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117Y:
		res = "117Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117Z:
		res = "117Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118X:
		res = "118X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118Y:
		res = "118Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118Z:
		res = "118Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119X:
		res = "119X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119Y:
		res = "119Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119Z:
		res = "119Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_120X:
		res = "120X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_120Y:
		res = "120Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_121X:
		res = "121X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_121Y:
		res = "121Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_122X:
		res = "122X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_122Y:
		res = "122Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_123X:
		res = "123X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_123Y:
		res = "123Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_124X:
		res = "124X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_124Y:
		res = "124Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_125X:
		res = "125X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_125Y:
		res = "125Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_126X:
		res = "126X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_126Y:
		res = "126Y"
	}
	return
}

func (enum_codedmechannelbasetype_member_1 *Enum_CodeDMEChannelBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_1X
		return
	case "1Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_1Y
		return
	case "2X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_2X
		return
	case "2Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_2Y
		return
	case "3X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_3X
		return
	case "3Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_3Y
		return
	case "4X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_4X
		return
	case "4Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_4Y
		return
	case "5X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_5X
		return
	case "5Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_5Y
		return
	case "6X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_6X
		return
	case "6Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_6Y
		return
	case "7X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_7X
		return
	case "7Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_7Y
		return
	case "8X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_8X
		return
	case "8Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_8Y
		return
	case "9X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_9X
		return
	case "9Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_9Y
		return
	case "10X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_10X
		return
	case "10Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_10Y
		return
	case "11X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_11X
		return
	case "11Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_11Y
		return
	case "12X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_12X
		return
	case "12Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_12Y
		return
	case "13X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_13X
		return
	case "13Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_13Y
		return
	case "14X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_14X
		return
	case "14Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_14Y
		return
	case "15X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_15X
		return
	case "15Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_15Y
		return
	case "16X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_16X
		return
	case "16Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_16Y
		return
	case "17X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17X
		return
	case "17Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17Y
		return
	case "17Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17Z
		return
	case "18X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18X
		return
	case "18W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18W
		return
	case "18Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18Y
		return
	case "18Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18Z
		return
	case "19X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19X
		return
	case "19Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19Y
		return
	case "19Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19Z
		return
	case "20X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20X
		return
	case "20W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20W
		return
	case "20Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20Y
		return
	case "20Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20Z
		return
	case "21X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21X
		return
	case "21Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21Y
		return
	case "21Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21Z
		return
	case "22X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22X
		return
	case "22W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22W
		return
	case "22Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22Y
		return
	case "22Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22Z
		return
	case "23X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23X
		return
	case "23Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23Y
		return
	case "23Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23Z
		return
	case "24X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24X
		return
	case "24W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24W
		return
	case "24Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24Y
		return
	case "24Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24Z
		return
	case "25X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25X
		return
	case "25Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25Y
		return
	case "25Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25Z
		return
	case "26X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26X
		return
	case "26W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26W
		return
	case "26Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26Y
		return
	case "26Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26Z
		return
	case "27X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27X
		return
	case "27Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27Y
		return
	case "27Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27Z
		return
	case "28X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28X
		return
	case "28W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28W
		return
	case "28Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28Y
		return
	case "28Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28Z
		return
	case "29X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29X
		return
	case "29Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29Y
		return
	case "29Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29Z
		return
	case "30X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30X
		return
	case "30W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30W
		return
	case "30Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30Y
		return
	case "30Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30Z
		return
	case "31X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31X
		return
	case "31Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31Y
		return
	case "31Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31Z
		return
	case "32X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32X
		return
	case "32W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32W
		return
	case "32Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32Y
		return
	case "32Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32Z
		return
	case "33X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33X
		return
	case "33Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33Y
		return
	case "33Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33Z
		return
	case "34X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34X
		return
	case "34W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34W
		return
	case "34Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34Y
		return
	case "34Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34Z
		return
	case "35X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35X
		return
	case "35Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35Y
		return
	case "35Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35Z
		return
	case "36X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36X
		return
	case "36W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36W
		return
	case "36Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36Y
		return
	case "36Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36Z
		return
	case "37X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37X
		return
	case "37Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37Y
		return
	case "37Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37Z
		return
	case "38X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38X
		return
	case "38W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38W
		return
	case "38Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38Y
		return
	case "38Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38Z
		return
	case "39X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39X
		return
	case "39Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39Y
		return
	case "39Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39Z
		return
	case "40X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40X
		return
	case "40W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40W
		return
	case "40Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40Y
		return
	case "40Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40Z
		return
	case "41X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41X
		return
	case "41Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41Y
		return
	case "41Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41Z
		return
	case "42X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42X
		return
	case "42W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42W
		return
	case "42Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42Y
		return
	case "42Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42Z
		return
	case "43X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43X
		return
	case "43Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43Y
		return
	case "43Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43Z
		return
	case "44X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44X
		return
	case "44W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44W
		return
	case "44Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44Y
		return
	case "44Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44Z
		return
	case "45X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45X
		return
	case "45Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45Y
		return
	case "45Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45Z
		return
	case "46X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46X
		return
	case "46W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46W
		return
	case "46Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46Y
		return
	case "46Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46Z
		return
	case "47X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47X
		return
	case "47Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47Y
		return
	case "47Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47Z
		return
	case "48X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48X
		return
	case "48W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48W
		return
	case "48Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48Y
		return
	case "48Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48Z
		return
	case "49X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49X
		return
	case "49Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49Y
		return
	case "49Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49Z
		return
	case "50X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50X
		return
	case "50W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50W
		return
	case "50Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50Y
		return
	case "50Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50Z
		return
	case "51X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51X
		return
	case "51Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51Y
		return
	case "51Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51Z
		return
	case "52X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52X
		return
	case "52W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52W
		return
	case "52Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52Y
		return
	case "52Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52Z
		return
	case "53X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53X
		return
	case "53Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53Y
		return
	case "53Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53Z
		return
	case "54X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54X
		return
	case "54W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54W
		return
	case "54Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54Y
		return
	case "54Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54Z
		return
	case "55X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55X
		return
	case "55Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55Y
		return
	case "55Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55Z
		return
	case "56X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56X
		return
	case "56W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56W
		return
	case "56Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56Y
		return
	case "56Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56Z
		return
	case "57X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_57X
		return
	case "57Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_57Y
		return
	case "58X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_58X
		return
	case "58Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_58Y
		return
	case "59X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_59X
		return
	case "59Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_59Y
		return
	case "60X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_60X
		return
	case "60Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_60Y
		return
	case "61X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_61X
		return
	case "61Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_61Y
		return
	case "62X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_62X
		return
	case "62Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_62Y
		return
	case "63X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_63X
		return
	case "63Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_63Y
		return
	case "64X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_64X
		return
	case "64Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_64Y
		return
	case "65X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_65X
		return
	case "65Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_65Y
		return
	case "66X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_66X
		return
	case "66Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_66Y
		return
	case "67X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_67X
		return
	case "67Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_67Y
		return
	case "68X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_68X
		return
	case "68Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_68Y
		return
	case "69X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_69X
		return
	case "69Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_69Y
		return
	case "70X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_70X
		return
	case "70Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_70Y
		return
	case "71X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_71X
		return
	case "71Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_71Y
		return
	case "72X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_72X
		return
	case "72Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_72Y
		return
	case "73X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_73X
		return
	case "73Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_73Y
		return
	case "74X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_74X
		return
	case "74Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_74Y
		return
	case "75X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_75X
		return
	case "75Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_75Y
		return
	case "76X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_76X
		return
	case "76Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_76Y
		return
	case "77X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_77X
		return
	case "77Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_77Y
		return
	case "78X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_78X
		return
	case "78Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_78Y
		return
	case "79X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_79X
		return
	case "79Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_79Y
		return
	case "80X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80X
		return
	case "80Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80Y
		return
	case "80Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80Z
		return
	case "81X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81X
		return
	case "81Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81Y
		return
	case "81Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81Z
		return
	case "82X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82X
		return
	case "82Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82Y
		return
	case "82Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82Z
		return
	case "83X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83X
		return
	case "83Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83Y
		return
	case "83Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83Z
		return
	case "84X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84X
		return
	case "84Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84Y
		return
	case "84Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84Z
		return
	case "85X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85X
		return
	case "85Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85Y
		return
	case "85Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85Z
		return
	case "86X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86X
		return
	case "86Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86Y
		return
	case "86Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86Z
		return
	case "87X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87X
		return
	case "87Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87Y
		return
	case "87Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87Z
		return
	case "88X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88X
		return
	case "88Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88Y
		return
	case "88Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88Z
		return
	case "89X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89X
		return
	case "89Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89Y
		return
	case "89Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89Z
		return
	case "90X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90X
		return
	case "90Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90Y
		return
	case "90Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90Z
		return
	case "91X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91X
		return
	case "91Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91Y
		return
	case "91Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91Z
		return
	case "92X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92X
		return
	case "92Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92Y
		return
	case "92Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92Z
		return
	case "93X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93X
		return
	case "93Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93Y
		return
	case "93Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93Z
		return
	case "94X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94X
		return
	case "94Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94Y
		return
	case "94Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94Z
		return
	case "95X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95X
		return
	case "95Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95Y
		return
	case "95Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95Z
		return
	case "96X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96X
		return
	case "96Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96Y
		return
	case "96Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96Z
		return
	case "97X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97X
		return
	case "97Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97Y
		return
	case "97Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97Z
		return
	case "98X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98X
		return
	case "98Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98Y
		return
	case "98Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98Z
		return
	case "99X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99X
		return
	case "99Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99Y
		return
	case "99Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99Z
		return
	case "100X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100X
		return
	case "100Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100Y
		return
	case "100Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100Z
		return
	case "101X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101X
		return
	case "101Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101Y
		return
	case "101Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101Z
		return
	case "102X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102X
		return
	case "102Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102Y
		return
	case "102Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102Z
		return
	case "103X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103X
		return
	case "103Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103Y
		return
	case "103Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103Z
		return
	case "104X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104X
		return
	case "104Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104Y
		return
	case "104Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104Z
		return
	case "105X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105X
		return
	case "105Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105Y
		return
	case "105Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105Z
		return
	case "106X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106X
		return
	case "106Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106Y
		return
	case "106Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106Z
		return
	case "107X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107X
		return
	case "107Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107Y
		return
	case "107Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107Z
		return
	case "108X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108X
		return
	case "108Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108Y
		return
	case "108Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108Z
		return
	case "109X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109X
		return
	case "109Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109Y
		return
	case "109Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109Z
		return
	case "110X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110X
		return
	case "110Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110Y
		return
	case "110Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110Z
		return
	case "111X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111X
		return
	case "111Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111Y
		return
	case "111Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111Z
		return
	case "112X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112X
		return
	case "112Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112Y
		return
	case "112Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112Z
		return
	case "113X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113X
		return
	case "113Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113Y
		return
	case "113Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113Z
		return
	case "114X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114X
		return
	case "114Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114Y
		return
	case "114Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114Z
		return
	case "115X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115X
		return
	case "115Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115Y
		return
	case "115Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115Z
		return
	case "116X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116X
		return
	case "116Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116Y
		return
	case "116Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116Z
		return
	case "117X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117X
		return
	case "117Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117Y
		return
	case "117Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117Z
		return
	case "118X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118X
		return
	case "118Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118Y
		return
	case "118Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118Z
		return
	case "119X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119X
		return
	case "119Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119Y
		return
	case "119Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119Z
		return
	case "120X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_120X
		return
	case "120Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_120Y
		return
	case "121X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_121X
		return
	case "121Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_121Y
		return
	case "122X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_122X
		return
	case "122Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_122Y
		return
	case "123X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_123X
		return
	case "123Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_123Y
		return
	case "124X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_124X
		return
	case "124Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_124Y
		return
	case "125X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_125X
		return
	case "125Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_125Y
		return
	case "126X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_126X
		return
	case "126Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_126Y
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedmechannelbasetype_member_1 *Enum_CodeDMEChannelBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDMEChannelBaseType_Member_1_N_1X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_1X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_1Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_1Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_2X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_2X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_2Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_2Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_3X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_3X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_3Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_3Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_4X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_4X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_4Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_4Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_5X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_5X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_5Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_5Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_6X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_6X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_6Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_6Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_7X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_7X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_7Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_7Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_8X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_8X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_8Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_8Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_9X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_9X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_9Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_9Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_10X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_10X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_10Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_10Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_11X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_11X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_11Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_11Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_12X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_12X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_12Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_12Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_13X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_13X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_13Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_13Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_14X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_14X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_14Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_14Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_15X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_15X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_15Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_15Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_16X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_16X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_16Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_16Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_17X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_17Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_17Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_17Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_18X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_18W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_18Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_18Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_18Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_19X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_19Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_19Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_19Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_20X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_20W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_20Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_20Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_20Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_21X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_21Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_21Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_21Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_22X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_22W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_22Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_22Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_22Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_23X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_23Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_23Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_23Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_24X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_24W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_24Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_24Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_24Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_25X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_25Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_25Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_25Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_26X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_26W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_26Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_26Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_26Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_27X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_27Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_27Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_27Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_28X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_28W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_28Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_28Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_28Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_29X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_29Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_29Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_29Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_30X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_30W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_30Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_30Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_30Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_31X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_31Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_31Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_31Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_32X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_32W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_32Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_32Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_32Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_33X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_33Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_33Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_33Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_34X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_34W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_34Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_34Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_34Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_35X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_35Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_35Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_35Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_36X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_36W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_36Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_36Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_36Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_37X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_37Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_37Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_37Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_38X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_38W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_38Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_38Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_38Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_39X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_39Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_39Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_39Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_40X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_40W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_40Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_40Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_40Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_41X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_41Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_41Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_41Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_42X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_42W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_42Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_42Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_42Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_43X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_43Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_43Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_43Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_44X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_44W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_44Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_44Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_44Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_45X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_45Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_45Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_45Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_46X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_46W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_46Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_46Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_46Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_47X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_47Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_47Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_47Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_48X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_48W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_48Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_48Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_48Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_49X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_49Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_49Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_49Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_50X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_50W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_50Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_50Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_50Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_51X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_51Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_51Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_51Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_52X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_52W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_52Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_52Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_52Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_53X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_53Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_53Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_53Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_54X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_54W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_54Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_54Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_54Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_55X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_55Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_55Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_55Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_56X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_56W":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56W
	case "Enum_CodeDMEChannelBaseType_Member_1_N_56Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_56Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_56Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_57X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_57X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_57Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_57Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_58X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_58X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_58Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_58Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_59X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_59X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_59Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_59Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_60X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_60X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_60Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_60Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_61X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_61X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_61Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_61Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_62X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_62X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_62Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_62Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_63X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_63X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_63Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_63Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_64X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_64X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_64Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_64Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_65X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_65X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_65Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_65Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_66X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_66X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_66Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_66Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_67X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_67X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_67Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_67Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_68X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_68X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_68Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_68Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_69X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_69X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_69Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_69Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_70X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_70X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_70Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_70Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_71X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_71X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_71Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_71Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_72X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_72X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_72Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_72Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_73X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_73X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_73Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_73Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_74X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_74X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_74Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_74Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_75X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_75X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_75Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_75Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_76X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_76X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_76Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_76Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_77X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_77X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_77Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_77Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_78X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_78X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_78Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_78Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_79X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_79X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_79Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_79Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_80X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_80Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_80Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_80Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_81X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_81Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_81Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_81Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_82X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_82Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_82Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_82Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_83X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_83Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_83Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_83Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_84X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_84Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_84Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_84Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_85X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_85Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_85Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_85Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_86X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_86Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_86Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_86Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_87X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_87Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_87Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_87Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_88X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_88Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_88Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_88Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_89X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_89Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_89Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_89Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_90X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_90Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_90Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_90Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_91X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_91Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_91Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_91Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_92X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_92Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_92Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_92Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_93X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_93Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_93Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_93Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_94X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_94Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_94Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_94Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_95X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_95Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_95Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_95Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_96X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_96Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_96Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_96Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_97X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_97Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_97Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_97Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_98X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_98Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_98Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_98Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_99X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_99Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_99Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_99Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_100X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_100Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_100Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_100Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_101X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_101Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_101Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_101Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_102X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_102Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_102Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_102Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_103X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_103Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_103Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_103Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_104X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_104Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_104Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_104Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_105X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_105Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_105Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_105Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_106X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_106Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_106Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_106Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_107X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_107Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_107Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_107Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_108X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_108Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_108Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_108Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_109X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_109Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_109Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_109Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_110X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_110Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_110Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_110Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_111X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_111Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_111Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_111Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_112X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_112Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_112Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_112Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_113X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_113Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_113Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_113Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_114X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_114Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_114Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_114Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_115X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_115Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_115Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_115Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_116X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_116Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_116Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_116Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_117X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_117Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_117Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_117Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_118X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_118Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_118Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_118Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_119X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_119Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_119Z":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_119Z
	case "Enum_CodeDMEChannelBaseType_Member_1_N_120X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_120X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_120Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_120Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_121X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_121X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_121Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_121Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_122X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_122X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_122Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_122Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_123X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_123X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_123Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_123Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_124X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_124X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_124Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_124Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_125X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_125X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_125Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_125Y
	case "Enum_CodeDMEChannelBaseType_Member_1_N_126X":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_126X
	case "Enum_CodeDMEChannelBaseType_Member_1_N_126Y":
		*enum_codedmechannelbasetype_member_1 = Enum_CodeDMEChannelBaseType_Member_1_N_126Y
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedmechannelbasetype_member_1 *Enum_CodeDMEChannelBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedmechannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDMEChannelBaseType_Member_1_N_1X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_1X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_1Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_1Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_2X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_2X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_2Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_2Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_3X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_3X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_3Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_3Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_4X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_4X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_4Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_4Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_5X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_5X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_5Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_5Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_6X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_6X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_6Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_6Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_7X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_7X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_7Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_7Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_8X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_8X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_8Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_8Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_9X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_9X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_9Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_9Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_10X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_10X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_10Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_10Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_11X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_11X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_11Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_11Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_12X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_12X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_12Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_12Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_13X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_13X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_13Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_13Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_14X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_14X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_14Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_14Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_15X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_15X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_15Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_15Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_16X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_16X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_16Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_16Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_17X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_17Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_17Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_17Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_18X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_18W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_18Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_18Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_18Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_19X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_19Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_19Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_19Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_20X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_20W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_20Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_20Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_20Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_21X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_21Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_21Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_21Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_22X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_22W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_22Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_22Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_22Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_23X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_23Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_23Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_23Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_24X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_24W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_24Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_24Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_24Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_25X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_25Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_25Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_25Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_26X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_26W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_26Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_26Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_26Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_27X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_27Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_27Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_27Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_28X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_28W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_28Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_28Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_28Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_29X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_29Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_29Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_29Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_30X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_30W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_30Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_30Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_30Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_31X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_31Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_31Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_31Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_32X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_32W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_32Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_32Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_32Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_33X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_33Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_33Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_33Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_34X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_34W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_34Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_34Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_34Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_35X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_35Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_35Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_35Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_36X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_36W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_36Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_36Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_36Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_37X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_37Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_37Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_37Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_38X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_38W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_38Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_38Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_38Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_39X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_39Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_39Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_39Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_40X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_40W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_40Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_40Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_40Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_41X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_41Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_41Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_41Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_42X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_42W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_42Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_42Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_42Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_43X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_43Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_43Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_43Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_44X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_44W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_44Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_44Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_44Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_45X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_45Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_45Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_45Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_46X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_46W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_46Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_46Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_46Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_47X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_47Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_47Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_47Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_48X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_48W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_48Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_48Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_48Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_49X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_49Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_49Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_49Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_50X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_50W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_50Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_50Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_50Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_51X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_51Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_51Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_51Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_52X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_52W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_52Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_52Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_52Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_53X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_53Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_53Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_53Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_54X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_54W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_54Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_54Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_54Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_55X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_55Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_55Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_55Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_56X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56W:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_56W"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_56Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_56Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_56Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_57X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_57X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_57Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_57Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_58X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_58X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_58Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_58Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_59X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_59X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_59Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_59Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_60X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_60X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_60Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_60Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_61X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_61X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_61Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_61Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_62X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_62X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_62Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_62Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_63X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_63X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_63Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_63Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_64X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_64X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_64Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_64Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_65X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_65X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_65Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_65Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_66X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_66X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_66Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_66Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_67X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_67X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_67Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_67Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_68X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_68X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_68Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_68Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_69X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_69X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_69Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_69Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_70X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_70X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_70Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_70Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_71X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_71X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_71Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_71Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_72X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_72X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_72Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_72Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_73X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_73X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_73Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_73Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_74X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_74X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_74Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_74Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_75X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_75X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_75Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_75Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_76X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_76X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_76Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_76Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_77X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_77X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_77Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_77Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_78X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_78X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_78Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_78Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_79X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_79X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_79Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_79Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_80X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_80Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_80Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_80Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_81X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_81Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_81Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_81Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_82X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_82Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_82Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_82Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_83X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_83Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_83Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_83Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_84X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_84Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_84Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_84Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_85X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_85Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_85Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_85Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_86X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_86Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_86Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_86Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_87X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_87Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_87Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_87Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_88X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_88Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_88Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_88Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_89X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_89Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_89Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_89Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_90X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_90Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_90Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_90Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_91X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_91Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_91Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_91Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_92X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_92Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_92Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_92Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_93X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_93Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_93Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_93Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_94X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_94Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_94Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_94Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_95X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_95Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_95Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_95Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_96X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_96Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_96Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_96Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_97X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_97Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_97Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_97Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_98X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_98Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_98Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_98Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_99X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_99Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_99Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_99Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_100X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_100Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_100Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_100Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_101X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_101Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_101Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_101Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_102X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_102Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_102Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_102Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_103X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_103Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_103Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_103Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_104X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_104Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_104Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_104Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_105X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_105Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_105Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_105Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_106X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_106Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_106Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_106Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_107X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_107Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_107Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_107Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_108X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_108Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_108Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_108Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_109X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_109Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_109Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_109Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_110X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_110Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_110Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_110Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_111X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_111Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_111Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_111Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_112X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_112Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_112Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_112Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_113X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_113Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_113Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_113Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_114X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_114Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_114Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_114Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_115X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_115Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_115Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_115Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_116X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_116Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_116Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_116Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_117X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_117Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_117Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_117Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_118X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_118Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_118Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_118Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_119X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_119Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_119Z:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_119Z"
	case Enum_CodeDMEChannelBaseType_Member_1_N_120X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_120X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_120Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_120Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_121X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_121X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_121Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_121Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_122X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_122X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_122Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_122Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_123X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_123X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_123Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_123Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_124X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_124X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_124Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_124Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_125X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_125X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_125Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_125Y"
	case Enum_CodeDMEChannelBaseType_Member_1_N_126X:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_126X"
	case Enum_CodeDMEChannelBaseType_Member_1_N_126Y:
		res = "Enum_CodeDMEChannelBaseType_Member_1_N_126Y"
	}
	return
}

func (enum_codedmechannelbasetype_member_1 Enum_CodeDMEChannelBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_1X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_1Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_2X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_2Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_3X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_3Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_4X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_4Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_5X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_5Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_6X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_6Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_7X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_7Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_8X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_8Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_9X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_9Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_10X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_10Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_11X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_11Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_12X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_12Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_13X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_13Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_14X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_14Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_15X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_15Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_16X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_16Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_17X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_17Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_17Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_18X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_18W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_18Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_18Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_19X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_19Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_19Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_20X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_20W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_20Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_20Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_21X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_21Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_21Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_22X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_22W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_22Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_22Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_23X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_23Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_23Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_24X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_24W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_24Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_24Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_25X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_25Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_25Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_26X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_26W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_26Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_26Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_27X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_27Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_27Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_28X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_28W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_28Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_28Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_29X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_29Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_29Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_30X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_30W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_30Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_30Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_31X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_31Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_31Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_32X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_32W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_32Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_32Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_33X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_33Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_33Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_34X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_34W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_34Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_34Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_35X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_35Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_35Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_36X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_36W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_36Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_36Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_37X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_37Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_37Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_38X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_38W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_38Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_38Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_39X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_39Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_39Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_40X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_40W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_40Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_40Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_41X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_41Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_41Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_42X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_42W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_42Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_42Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_43X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_43Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_43Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_44X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_44W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_44Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_44Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_45X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_45Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_45Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_46X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_46W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_46Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_46Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_47X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_47Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_47Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_48X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_48W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_48Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_48Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_49X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_49Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_49Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_50X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_50W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_50Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_50Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_51X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_51Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_51Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_52X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_52W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_52Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_52Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_53X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_53Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_53Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_54X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_54W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_54Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_54Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_55X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_55Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_55Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_56X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_56W")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_56Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_56Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_57X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_57Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_58X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_58Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_59X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_59Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_60X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_60Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_61X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_61Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_62X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_62Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_63X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_63Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_64X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_64Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_65X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_65Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_66X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_66Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_67X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_67Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_68X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_68Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_69X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_69Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_70X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_70Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_71X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_71Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_72X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_72Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_73X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_73Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_74X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_74Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_75X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_75Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_76X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_76Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_77X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_77Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_78X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_78Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_79X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_79Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_80X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_80Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_80Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_81X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_81Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_81Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_82X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_82Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_82Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_83X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_83Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_83Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_84X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_84Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_84Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_85X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_85Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_85Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_86X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_86Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_86Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_87X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_87Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_87Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_88X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_88Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_88Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_89X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_89Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_89Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_90X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_90Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_90Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_91X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_91Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_91Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_92X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_92Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_92Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_93X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_93Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_93Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_94X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_94Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_94Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_95X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_95Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_95Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_96X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_96Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_96Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_97X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_97Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_97Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_98X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_98Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_98Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_99X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_99Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_99Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_100X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_100Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_100Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_101X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_101Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_101Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_102X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_102Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_102Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_103X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_103Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_103Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_104X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_104Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_104Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_105X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_105Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_105Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_106X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_106Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_106Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_107X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_107Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_107Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_108X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_108Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_108Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_109X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_109Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_109Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_110X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_110Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_110Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_111X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_111Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_111Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_112X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_112Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_112Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_113X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_113Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_113Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_114X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_114Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_114Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_115X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_115Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_115Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_116X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_116Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_116Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_117X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_117Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_117Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_118X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_118Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_118Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_119X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_119Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_119Z")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_120X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_120Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_121X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_121Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_122X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_122Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_123X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_123Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_124X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_124Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_125X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_125Y")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_126X")
	res = append(res, "Enum_CodeDMEChannelBaseType_Member_1_N_126Y")

	return
}

func (enum_codedmechannelbasetype_member_1 Enum_CodeDMEChannelBaseType_Member_1) CodeValues() (res []string) {

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

// Utility function for Enum_CodeDesignatedPointBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedesignatedpointbasetype_member_1 Enum_CodeDesignatedPointBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedesignatedpointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDesignatedPointBaseType_Member_1_ICAO:
		res = "ICAO"
	case Enum_CodeDesignatedPointBaseType_Member_1_COORD:
		res = "COORD"
	case Enum_CodeDesignatedPointBaseType_Member_1_CNF:
		res = "CNF"
	case Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED:
		res = "DESIGNED"
	case Enum_CodeDesignatedPointBaseType_Member_1_MTR:
		res = "MTR"
	case Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL:
		res = "TERMINAL"
	case Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST:
		res = "BRG_DIST"
	}
	return
}

func (enum_codedesignatedpointbasetype_member_1 *Enum_CodeDesignatedPointBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ICAO":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_ICAO
		return
	case "COORD":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_COORD
		return
	case "CNF":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_CNF
		return
	case "DESIGNED":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED
		return
	case "MTR":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_MTR
		return
	case "TERMINAL":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL
		return
	case "BRG_DIST":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedesignatedpointbasetype_member_1 *Enum_CodeDesignatedPointBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDesignatedPointBaseType_Member_1_ICAO":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_ICAO
	case "Enum_CodeDesignatedPointBaseType_Member_1_COORD":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_COORD
	case "Enum_CodeDesignatedPointBaseType_Member_1_CNF":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_CNF
	case "Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED
	case "Enum_CodeDesignatedPointBaseType_Member_1_MTR":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_MTR
	case "Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL
	case "Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST":
		*enum_codedesignatedpointbasetype_member_1 = Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedesignatedpointbasetype_member_1 *Enum_CodeDesignatedPointBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedesignatedpointbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDesignatedPointBaseType_Member_1_ICAO:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_ICAO"
	case Enum_CodeDesignatedPointBaseType_Member_1_COORD:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_COORD"
	case Enum_CodeDesignatedPointBaseType_Member_1_CNF:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_CNF"
	case Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED"
	case Enum_CodeDesignatedPointBaseType_Member_1_MTR:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_MTR"
	case Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL"
	case Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST:
		res = "Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST"
	}
	return
}

func (enum_codedesignatedpointbasetype_member_1 Enum_CodeDesignatedPointBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_ICAO")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_COORD")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_CNF")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_DESIGNED")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_MTR")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_TERMINAL")
	res = append(res, "Enum_CodeDesignatedPointBaseType_Member_1_BRG_DIST")

	return
}

func (enum_codedesignatedpointbasetype_member_1 Enum_CodeDesignatedPointBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ICAO")
	res = append(res, "COORD")
	res = append(res, "CNF")
	res = append(res, "DESIGNED")
	res = append(res, "MTR")
	res = append(res, "TERMINAL")
	res = append(res, "BRG_DIST")

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

// Utility function for Enum_CodeDistanceIndicationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codedistanceindicationbasetype_member_1 Enum_CodeDistanceIndicationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codedistanceindicationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDistanceIndicationBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC:
		res = "GEODETIC"
	}
	return
}

func (enum_codedistanceindicationbasetype_member_1 *Enum_CodeDistanceIndicationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DME":
		*enum_codedistanceindicationbasetype_member_1 = Enum_CodeDistanceIndicationBaseType_Member_1_DME
		return
	case "GEODETIC":
		*enum_codedistanceindicationbasetype_member_1 = Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codedistanceindicationbasetype_member_1 *Enum_CodeDistanceIndicationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeDistanceIndicationBaseType_Member_1_DME":
		*enum_codedistanceindicationbasetype_member_1 = Enum_CodeDistanceIndicationBaseType_Member_1_DME
	case "Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC":
		*enum_codedistanceindicationbasetype_member_1 = Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codedistanceindicationbasetype_member_1 *Enum_CodeDistanceIndicationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codedistanceindicationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeDistanceIndicationBaseType_Member_1_DME:
		res = "Enum_CodeDistanceIndicationBaseType_Member_1_DME"
	case Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC:
		res = "Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC"
	}
	return
}

func (enum_codedistanceindicationbasetype_member_1 Enum_CodeDistanceIndicationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeDistanceIndicationBaseType_Member_1_DME")
	res = append(res, "Enum_CodeDistanceIndicationBaseType_Member_1_GEODETIC")

	return
}

func (enum_codedistanceindicationbasetype_member_1 Enum_CodeDistanceIndicationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DME")
	res = append(res, "GEODETIC")

	return
}

// Utility function for Enum_CodeEmissionBandBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeemissionbandbasetype_member_1 Enum_CodeEmissionBandBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeemissionbandbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEmissionBandBaseType_Member_1_U:
		res = "U"
	case Enum_CodeEmissionBandBaseType_Member_1_H:
		res = "H"
	case Enum_CodeEmissionBandBaseType_Member_1_M:
		res = "M"
	}
	return
}

func (enum_codeemissionbandbasetype_member_1 *Enum_CodeEmissionBandBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "U":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_U
		return
	case "H":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_H
		return
	case "M":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_M
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeemissionbandbasetype_member_1 *Enum_CodeEmissionBandBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeEmissionBandBaseType_Member_1_U":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_U
	case "Enum_CodeEmissionBandBaseType_Member_1_H":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_H
	case "Enum_CodeEmissionBandBaseType_Member_1_M":
		*enum_codeemissionbandbasetype_member_1 = Enum_CodeEmissionBandBaseType_Member_1_M
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeemissionbandbasetype_member_1 *Enum_CodeEmissionBandBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeemissionbandbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeEmissionBandBaseType_Member_1_U:
		res = "Enum_CodeEmissionBandBaseType_Member_1_U"
	case Enum_CodeEmissionBandBaseType_Member_1_H:
		res = "Enum_CodeEmissionBandBaseType_Member_1_H"
	case Enum_CodeEmissionBandBaseType_Member_1_M:
		res = "Enum_CodeEmissionBandBaseType_Member_1_M"
	}
	return
}

func (enum_codeemissionbandbasetype_member_1 Enum_CodeEmissionBandBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeEmissionBandBaseType_Member_1_U")
	res = append(res, "Enum_CodeEmissionBandBaseType_Member_1_H")
	res = append(res, "Enum_CodeEmissionBandBaseType_Member_1_M")

	return
}

func (enum_codeemissionbandbasetype_member_1 Enum_CodeEmissionBandBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "U")
	res = append(res, "H")
	res = append(res, "M")

	return
}

// Utility function for Enum_CodeILSBackCourseBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeilsbackcoursebasetype_member_1 Enum_CodeILSBackCourseBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeilsbackcoursebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeILSBackCourseBaseType_Member_1_YES:
		res = "YES"
	case Enum_CodeILSBackCourseBaseType_Member_1_NO:
		res = "NO"
	case Enum_CodeILSBackCourseBaseType_Member_1_RSTR:
		res = "RSTR"
	}
	return
}

func (enum_codeilsbackcoursebasetype_member_1 *Enum_CodeILSBackCourseBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "YES":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_YES
		return
	case "NO":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_NO
		return
	case "RSTR":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_RSTR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeilsbackcoursebasetype_member_1 *Enum_CodeILSBackCourseBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeILSBackCourseBaseType_Member_1_YES":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_YES
	case "Enum_CodeILSBackCourseBaseType_Member_1_NO":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_NO
	case "Enum_CodeILSBackCourseBaseType_Member_1_RSTR":
		*enum_codeilsbackcoursebasetype_member_1 = Enum_CodeILSBackCourseBaseType_Member_1_RSTR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeilsbackcoursebasetype_member_1 *Enum_CodeILSBackCourseBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeilsbackcoursebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeILSBackCourseBaseType_Member_1_YES:
		res = "Enum_CodeILSBackCourseBaseType_Member_1_YES"
	case Enum_CodeILSBackCourseBaseType_Member_1_NO:
		res = "Enum_CodeILSBackCourseBaseType_Member_1_NO"
	case Enum_CodeILSBackCourseBaseType_Member_1_RSTR:
		res = "Enum_CodeILSBackCourseBaseType_Member_1_RSTR"
	}
	return
}

func (enum_codeilsbackcoursebasetype_member_1 Enum_CodeILSBackCourseBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeILSBackCourseBaseType_Member_1_YES")
	res = append(res, "Enum_CodeILSBackCourseBaseType_Member_1_NO")
	res = append(res, "Enum_CodeILSBackCourseBaseType_Member_1_RSTR")

	return
}

func (enum_codeilsbackcoursebasetype_member_1 Enum_CodeILSBackCourseBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "YES")
	res = append(res, "NO")
	res = append(res, "RSTR")

	return
}

// Utility function for Enum_CodeIntegrityLevelILSBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeintegritylevelilsbasetype_member_1 Enum_CodeIntegrityLevelILSBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeintegritylevelilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1:
		res = "1"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2:
		res = "2"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3:
		res = "3"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4:
		res = "4"
	}
	return
}

func (enum_codeintegritylevelilsbasetype_member_1 *Enum_CodeIntegrityLevelILSBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1
		return
	case "2":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2
		return
	case "3":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3
		return
	case "4":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeintegritylevelilsbasetype_member_1 *Enum_CodeIntegrityLevelILSBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1
	case "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2
	case "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3
	case "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4":
		*enum_codeintegritylevelilsbasetype_member_1 = Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeintegritylevelilsbasetype_member_1 *Enum_CodeIntegrityLevelILSBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeintegritylevelilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1:
		res = "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2:
		res = "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3:
		res = "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3"
	case Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4:
		res = "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4"
	}
	return
}

func (enum_codeintegritylevelilsbasetype_member_1 Enum_CodeIntegrityLevelILSBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_1")
	res = append(res, "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_2")
	res = append(res, "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_3")
	res = append(res, "Enum_CodeIntegrityLevelILSBaseType_Member_1_N_4")

	return
}

func (enum_codeintegritylevelilsbasetype_member_1 Enum_CodeIntegrityLevelILSBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1")
	res = append(res, "2")
	res = append(res, "3")
	res = append(res, "4")

	return
}

// Utility function for Enum_CodeMLSAzimuthBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemlsazimuthbasetype_member_1 Enum_CodeMLSAzimuthBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemlsazimuthbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMLSAzimuthBaseType_Member_1_FWD:
		res = "FWD"
	case Enum_CodeMLSAzimuthBaseType_Member_1_BWD:
		res = "BWD"
	}
	return
}

func (enum_codemlsazimuthbasetype_member_1 *Enum_CodeMLSAzimuthBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FWD":
		*enum_codemlsazimuthbasetype_member_1 = Enum_CodeMLSAzimuthBaseType_Member_1_FWD
		return
	case "BWD":
		*enum_codemlsazimuthbasetype_member_1 = Enum_CodeMLSAzimuthBaseType_Member_1_BWD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemlsazimuthbasetype_member_1 *Enum_CodeMLSAzimuthBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMLSAzimuthBaseType_Member_1_FWD":
		*enum_codemlsazimuthbasetype_member_1 = Enum_CodeMLSAzimuthBaseType_Member_1_FWD
	case "Enum_CodeMLSAzimuthBaseType_Member_1_BWD":
		*enum_codemlsazimuthbasetype_member_1 = Enum_CodeMLSAzimuthBaseType_Member_1_BWD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemlsazimuthbasetype_member_1 *Enum_CodeMLSAzimuthBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemlsazimuthbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMLSAzimuthBaseType_Member_1_FWD:
		res = "Enum_CodeMLSAzimuthBaseType_Member_1_FWD"
	case Enum_CodeMLSAzimuthBaseType_Member_1_BWD:
		res = "Enum_CodeMLSAzimuthBaseType_Member_1_BWD"
	}
	return
}

func (enum_codemlsazimuthbasetype_member_1 Enum_CodeMLSAzimuthBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMLSAzimuthBaseType_Member_1_FWD")
	res = append(res, "Enum_CodeMLSAzimuthBaseType_Member_1_BWD")

	return
}

func (enum_codemlsazimuthbasetype_member_1 Enum_CodeMLSAzimuthBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FWD")
	res = append(res, "BWD")

	return
}

// Utility function for Enum_CodeMLSChannelBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemlschannelbasetype_member_1 Enum_CodeMLSChannelBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemlschannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMLSChannelBaseType_Member_1_N_500:
		res = "500"
	case Enum_CodeMLSChannelBaseType_Member_1_N_501:
		res = "501"
	case Enum_CodeMLSChannelBaseType_Member_1_N_502:
		res = "502"
	case Enum_CodeMLSChannelBaseType_Member_1_N_503:
		res = "503"
	case Enum_CodeMLSChannelBaseType_Member_1_N_504:
		res = "504"
	case Enum_CodeMLSChannelBaseType_Member_1_N_505:
		res = "505"
	case Enum_CodeMLSChannelBaseType_Member_1_N_506:
		res = "506"
	case Enum_CodeMLSChannelBaseType_Member_1_N_507:
		res = "507"
	case Enum_CodeMLSChannelBaseType_Member_1_N_508:
		res = "508"
	case Enum_CodeMLSChannelBaseType_Member_1_N_509:
		res = "509"
	case Enum_CodeMLSChannelBaseType_Member_1_N_510:
		res = "510"
	case Enum_CodeMLSChannelBaseType_Member_1_N_511:
		res = "511"
	case Enum_CodeMLSChannelBaseType_Member_1_N_512:
		res = "512"
	case Enum_CodeMLSChannelBaseType_Member_1_N_513:
		res = "513"
	case Enum_CodeMLSChannelBaseType_Member_1_N_514:
		res = "514"
	case Enum_CodeMLSChannelBaseType_Member_1_N_515:
		res = "515"
	case Enum_CodeMLSChannelBaseType_Member_1_N_516:
		res = "516"
	case Enum_CodeMLSChannelBaseType_Member_1_N_517:
		res = "517"
	case Enum_CodeMLSChannelBaseType_Member_1_N_518:
		res = "518"
	case Enum_CodeMLSChannelBaseType_Member_1_N_519:
		res = "519"
	case Enum_CodeMLSChannelBaseType_Member_1_N_520:
		res = "520"
	case Enum_CodeMLSChannelBaseType_Member_1_N_521:
		res = "521"
	case Enum_CodeMLSChannelBaseType_Member_1_N_522:
		res = "522"
	case Enum_CodeMLSChannelBaseType_Member_1_N_523:
		res = "523"
	case Enum_CodeMLSChannelBaseType_Member_1_N_524:
		res = "524"
	case Enum_CodeMLSChannelBaseType_Member_1_N_525:
		res = "525"
	case Enum_CodeMLSChannelBaseType_Member_1_N_526:
		res = "526"
	case Enum_CodeMLSChannelBaseType_Member_1_N_527:
		res = "527"
	case Enum_CodeMLSChannelBaseType_Member_1_N_528:
		res = "528"
	case Enum_CodeMLSChannelBaseType_Member_1_N_529:
		res = "529"
	case Enum_CodeMLSChannelBaseType_Member_1_N_530:
		res = "530"
	case Enum_CodeMLSChannelBaseType_Member_1_N_531:
		res = "531"
	case Enum_CodeMLSChannelBaseType_Member_1_N_532:
		res = "532"
	case Enum_CodeMLSChannelBaseType_Member_1_N_533:
		res = "533"
	case Enum_CodeMLSChannelBaseType_Member_1_N_534:
		res = "534"
	case Enum_CodeMLSChannelBaseType_Member_1_N_535:
		res = "535"
	case Enum_CodeMLSChannelBaseType_Member_1_N_536:
		res = "536"
	case Enum_CodeMLSChannelBaseType_Member_1_N_537:
		res = "537"
	case Enum_CodeMLSChannelBaseType_Member_1_N_538:
		res = "538"
	case Enum_CodeMLSChannelBaseType_Member_1_N_539:
		res = "539"
	case Enum_CodeMLSChannelBaseType_Member_1_N_540:
		res = "540"
	case Enum_CodeMLSChannelBaseType_Member_1_N_541:
		res = "541"
	case Enum_CodeMLSChannelBaseType_Member_1_N_542:
		res = "542"
	case Enum_CodeMLSChannelBaseType_Member_1_N_543:
		res = "543"
	case Enum_CodeMLSChannelBaseType_Member_1_N_544:
		res = "544"
	case Enum_CodeMLSChannelBaseType_Member_1_N_545:
		res = "545"
	case Enum_CodeMLSChannelBaseType_Member_1_N_546:
		res = "546"
	case Enum_CodeMLSChannelBaseType_Member_1_N_547:
		res = "547"
	case Enum_CodeMLSChannelBaseType_Member_1_N_548:
		res = "548"
	case Enum_CodeMLSChannelBaseType_Member_1_N_549:
		res = "549"
	case Enum_CodeMLSChannelBaseType_Member_1_N_550:
		res = "550"
	case Enum_CodeMLSChannelBaseType_Member_1_N_551:
		res = "551"
	case Enum_CodeMLSChannelBaseType_Member_1_N_552:
		res = "552"
	case Enum_CodeMLSChannelBaseType_Member_1_N_553:
		res = "553"
	case Enum_CodeMLSChannelBaseType_Member_1_N_554:
		res = "554"
	case Enum_CodeMLSChannelBaseType_Member_1_N_555:
		res = "555"
	case Enum_CodeMLSChannelBaseType_Member_1_N_556:
		res = "556"
	case Enum_CodeMLSChannelBaseType_Member_1_N_557:
		res = "557"
	case Enum_CodeMLSChannelBaseType_Member_1_N_558:
		res = "558"
	case Enum_CodeMLSChannelBaseType_Member_1_N_559:
		res = "559"
	case Enum_CodeMLSChannelBaseType_Member_1_N_560:
		res = "560"
	case Enum_CodeMLSChannelBaseType_Member_1_N_561:
		res = "561"
	case Enum_CodeMLSChannelBaseType_Member_1_N_562:
		res = "562"
	case Enum_CodeMLSChannelBaseType_Member_1_N_563:
		res = "563"
	case Enum_CodeMLSChannelBaseType_Member_1_N_564:
		res = "564"
	case Enum_CodeMLSChannelBaseType_Member_1_N_565:
		res = "565"
	case Enum_CodeMLSChannelBaseType_Member_1_N_566:
		res = "566"
	case Enum_CodeMLSChannelBaseType_Member_1_N_567:
		res = "567"
	case Enum_CodeMLSChannelBaseType_Member_1_N_568:
		res = "568"
	case Enum_CodeMLSChannelBaseType_Member_1_N_569:
		res = "569"
	case Enum_CodeMLSChannelBaseType_Member_1_N_570:
		res = "570"
	case Enum_CodeMLSChannelBaseType_Member_1_N_571:
		res = "571"
	case Enum_CodeMLSChannelBaseType_Member_1_N_572:
		res = "572"
	case Enum_CodeMLSChannelBaseType_Member_1_N_573:
		res = "573"
	case Enum_CodeMLSChannelBaseType_Member_1_N_574:
		res = "574"
	case Enum_CodeMLSChannelBaseType_Member_1_N_575:
		res = "575"
	case Enum_CodeMLSChannelBaseType_Member_1_N_576:
		res = "576"
	case Enum_CodeMLSChannelBaseType_Member_1_N_577:
		res = "577"
	case Enum_CodeMLSChannelBaseType_Member_1_N_578:
		res = "578"
	case Enum_CodeMLSChannelBaseType_Member_1_N_579:
		res = "579"
	case Enum_CodeMLSChannelBaseType_Member_1_N_580:
		res = "580"
	case Enum_CodeMLSChannelBaseType_Member_1_N_581:
		res = "581"
	case Enum_CodeMLSChannelBaseType_Member_1_N_582:
		res = "582"
	case Enum_CodeMLSChannelBaseType_Member_1_N_583:
		res = "583"
	case Enum_CodeMLSChannelBaseType_Member_1_N_584:
		res = "584"
	case Enum_CodeMLSChannelBaseType_Member_1_N_585:
		res = "585"
	case Enum_CodeMLSChannelBaseType_Member_1_N_586:
		res = "586"
	case Enum_CodeMLSChannelBaseType_Member_1_N_587:
		res = "587"
	case Enum_CodeMLSChannelBaseType_Member_1_N_588:
		res = "588"
	case Enum_CodeMLSChannelBaseType_Member_1_N_589:
		res = "589"
	case Enum_CodeMLSChannelBaseType_Member_1_N_590:
		res = "590"
	case Enum_CodeMLSChannelBaseType_Member_1_N_591:
		res = "591"
	case Enum_CodeMLSChannelBaseType_Member_1_N_592:
		res = "592"
	case Enum_CodeMLSChannelBaseType_Member_1_N_593:
		res = "593"
	case Enum_CodeMLSChannelBaseType_Member_1_N_594:
		res = "594"
	case Enum_CodeMLSChannelBaseType_Member_1_N_595:
		res = "595"
	case Enum_CodeMLSChannelBaseType_Member_1_N_596:
		res = "596"
	case Enum_CodeMLSChannelBaseType_Member_1_N_597:
		res = "597"
	case Enum_CodeMLSChannelBaseType_Member_1_N_598:
		res = "598"
	case Enum_CodeMLSChannelBaseType_Member_1_N_599:
		res = "599"
	case Enum_CodeMLSChannelBaseType_Member_1_N_600:
		res = "600"
	case Enum_CodeMLSChannelBaseType_Member_1_N_601:
		res = "601"
	case Enum_CodeMLSChannelBaseType_Member_1_N_602:
		res = "602"
	case Enum_CodeMLSChannelBaseType_Member_1_N_603:
		res = "603"
	case Enum_CodeMLSChannelBaseType_Member_1_N_604:
		res = "604"
	case Enum_CodeMLSChannelBaseType_Member_1_N_605:
		res = "605"
	case Enum_CodeMLSChannelBaseType_Member_1_N_606:
		res = "606"
	case Enum_CodeMLSChannelBaseType_Member_1_N_607:
		res = "607"
	case Enum_CodeMLSChannelBaseType_Member_1_N_608:
		res = "608"
	case Enum_CodeMLSChannelBaseType_Member_1_N_609:
		res = "609"
	case Enum_CodeMLSChannelBaseType_Member_1_N_610:
		res = "610"
	case Enum_CodeMLSChannelBaseType_Member_1_N_611:
		res = "611"
	case Enum_CodeMLSChannelBaseType_Member_1_N_612:
		res = "612"
	case Enum_CodeMLSChannelBaseType_Member_1_N_613:
		res = "613"
	case Enum_CodeMLSChannelBaseType_Member_1_N_614:
		res = "614"
	case Enum_CodeMLSChannelBaseType_Member_1_N_615:
		res = "615"
	case Enum_CodeMLSChannelBaseType_Member_1_N_616:
		res = "616"
	case Enum_CodeMLSChannelBaseType_Member_1_N_617:
		res = "617"
	case Enum_CodeMLSChannelBaseType_Member_1_N_618:
		res = "618"
	case Enum_CodeMLSChannelBaseType_Member_1_N_619:
		res = "619"
	case Enum_CodeMLSChannelBaseType_Member_1_N_620:
		res = "620"
	case Enum_CodeMLSChannelBaseType_Member_1_N_621:
		res = "621"
	case Enum_CodeMLSChannelBaseType_Member_1_N_622:
		res = "622"
	case Enum_CodeMLSChannelBaseType_Member_1_N_623:
		res = "623"
	case Enum_CodeMLSChannelBaseType_Member_1_N_624:
		res = "624"
	case Enum_CodeMLSChannelBaseType_Member_1_N_625:
		res = "625"
	case Enum_CodeMLSChannelBaseType_Member_1_N_626:
		res = "626"
	case Enum_CodeMLSChannelBaseType_Member_1_N_627:
		res = "627"
	case Enum_CodeMLSChannelBaseType_Member_1_N_628:
		res = "628"
	case Enum_CodeMLSChannelBaseType_Member_1_N_629:
		res = "629"
	case Enum_CodeMLSChannelBaseType_Member_1_N_630:
		res = "630"
	case Enum_CodeMLSChannelBaseType_Member_1_N_631:
		res = "631"
	case Enum_CodeMLSChannelBaseType_Member_1_N_632:
		res = "632"
	case Enum_CodeMLSChannelBaseType_Member_1_N_633:
		res = "633"
	case Enum_CodeMLSChannelBaseType_Member_1_N_634:
		res = "634"
	case Enum_CodeMLSChannelBaseType_Member_1_N_635:
		res = "635"
	case Enum_CodeMLSChannelBaseType_Member_1_N_636:
		res = "636"
	case Enum_CodeMLSChannelBaseType_Member_1_N_637:
		res = "637"
	case Enum_CodeMLSChannelBaseType_Member_1_N_638:
		res = "638"
	case Enum_CodeMLSChannelBaseType_Member_1_N_639:
		res = "639"
	case Enum_CodeMLSChannelBaseType_Member_1_N_640:
		res = "640"
	case Enum_CodeMLSChannelBaseType_Member_1_N_641:
		res = "641"
	case Enum_CodeMLSChannelBaseType_Member_1_N_642:
		res = "642"
	case Enum_CodeMLSChannelBaseType_Member_1_N_643:
		res = "643"
	case Enum_CodeMLSChannelBaseType_Member_1_N_644:
		res = "644"
	case Enum_CodeMLSChannelBaseType_Member_1_N_645:
		res = "645"
	case Enum_CodeMLSChannelBaseType_Member_1_N_646:
		res = "646"
	case Enum_CodeMLSChannelBaseType_Member_1_N_647:
		res = "647"
	case Enum_CodeMLSChannelBaseType_Member_1_N_648:
		res = "648"
	case Enum_CodeMLSChannelBaseType_Member_1_N_649:
		res = "649"
	case Enum_CodeMLSChannelBaseType_Member_1_N_650:
		res = "650"
	case Enum_CodeMLSChannelBaseType_Member_1_N_651:
		res = "651"
	case Enum_CodeMLSChannelBaseType_Member_1_N_652:
		res = "652"
	case Enum_CodeMLSChannelBaseType_Member_1_N_653:
		res = "653"
	case Enum_CodeMLSChannelBaseType_Member_1_N_654:
		res = "654"
	case Enum_CodeMLSChannelBaseType_Member_1_N_655:
		res = "655"
	case Enum_CodeMLSChannelBaseType_Member_1_N_656:
		res = "656"
	case Enum_CodeMLSChannelBaseType_Member_1_N_657:
		res = "657"
	case Enum_CodeMLSChannelBaseType_Member_1_N_658:
		res = "658"
	case Enum_CodeMLSChannelBaseType_Member_1_N_659:
		res = "659"
	case Enum_CodeMLSChannelBaseType_Member_1_N_660:
		res = "660"
	case Enum_CodeMLSChannelBaseType_Member_1_N_661:
		res = "661"
	case Enum_CodeMLSChannelBaseType_Member_1_N_662:
		res = "662"
	case Enum_CodeMLSChannelBaseType_Member_1_N_663:
		res = "663"
	case Enum_CodeMLSChannelBaseType_Member_1_N_664:
		res = "664"
	case Enum_CodeMLSChannelBaseType_Member_1_N_665:
		res = "665"
	case Enum_CodeMLSChannelBaseType_Member_1_N_666:
		res = "666"
	case Enum_CodeMLSChannelBaseType_Member_1_N_667:
		res = "667"
	case Enum_CodeMLSChannelBaseType_Member_1_N_668:
		res = "668"
	case Enum_CodeMLSChannelBaseType_Member_1_N_669:
		res = "669"
	case Enum_CodeMLSChannelBaseType_Member_1_N_670:
		res = "670"
	case Enum_CodeMLSChannelBaseType_Member_1_N_671:
		res = "671"
	case Enum_CodeMLSChannelBaseType_Member_1_N_672:
		res = "672"
	case Enum_CodeMLSChannelBaseType_Member_1_N_673:
		res = "673"
	case Enum_CodeMLSChannelBaseType_Member_1_N_674:
		res = "674"
	case Enum_CodeMLSChannelBaseType_Member_1_N_675:
		res = "675"
	case Enum_CodeMLSChannelBaseType_Member_1_N_676:
		res = "676"
	case Enum_CodeMLSChannelBaseType_Member_1_N_677:
		res = "677"
	case Enum_CodeMLSChannelBaseType_Member_1_N_678:
		res = "678"
	case Enum_CodeMLSChannelBaseType_Member_1_N_679:
		res = "679"
	case Enum_CodeMLSChannelBaseType_Member_1_N_680:
		res = "680"
	case Enum_CodeMLSChannelBaseType_Member_1_N_681:
		res = "681"
	case Enum_CodeMLSChannelBaseType_Member_1_N_682:
		res = "682"
	case Enum_CodeMLSChannelBaseType_Member_1_N_683:
		res = "683"
	case Enum_CodeMLSChannelBaseType_Member_1_N_684:
		res = "684"
	case Enum_CodeMLSChannelBaseType_Member_1_N_685:
		res = "685"
	case Enum_CodeMLSChannelBaseType_Member_1_N_686:
		res = "686"
	case Enum_CodeMLSChannelBaseType_Member_1_N_687:
		res = "687"
	case Enum_CodeMLSChannelBaseType_Member_1_N_688:
		res = "688"
	case Enum_CodeMLSChannelBaseType_Member_1_N_689:
		res = "689"
	case Enum_CodeMLSChannelBaseType_Member_1_N_690:
		res = "690"
	case Enum_CodeMLSChannelBaseType_Member_1_N_691:
		res = "691"
	case Enum_CodeMLSChannelBaseType_Member_1_N_692:
		res = "692"
	case Enum_CodeMLSChannelBaseType_Member_1_N_693:
		res = "693"
	case Enum_CodeMLSChannelBaseType_Member_1_N_694:
		res = "694"
	case Enum_CodeMLSChannelBaseType_Member_1_N_695:
		res = "695"
	case Enum_CodeMLSChannelBaseType_Member_1_N_696:
		res = "696"
	case Enum_CodeMLSChannelBaseType_Member_1_N_697:
		res = "697"
	case Enum_CodeMLSChannelBaseType_Member_1_N_698:
		res = "698"
	case Enum_CodeMLSChannelBaseType_Member_1_N_699:
		res = "699"
	}
	return
}

func (enum_codemlschannelbasetype_member_1 *Enum_CodeMLSChannelBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "500":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_500
		return
	case "501":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_501
		return
	case "502":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_502
		return
	case "503":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_503
		return
	case "504":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_504
		return
	case "505":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_505
		return
	case "506":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_506
		return
	case "507":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_507
		return
	case "508":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_508
		return
	case "509":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_509
		return
	case "510":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_510
		return
	case "511":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_511
		return
	case "512":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_512
		return
	case "513":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_513
		return
	case "514":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_514
		return
	case "515":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_515
		return
	case "516":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_516
		return
	case "517":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_517
		return
	case "518":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_518
		return
	case "519":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_519
		return
	case "520":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_520
		return
	case "521":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_521
		return
	case "522":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_522
		return
	case "523":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_523
		return
	case "524":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_524
		return
	case "525":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_525
		return
	case "526":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_526
		return
	case "527":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_527
		return
	case "528":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_528
		return
	case "529":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_529
		return
	case "530":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_530
		return
	case "531":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_531
		return
	case "532":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_532
		return
	case "533":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_533
		return
	case "534":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_534
		return
	case "535":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_535
		return
	case "536":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_536
		return
	case "537":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_537
		return
	case "538":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_538
		return
	case "539":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_539
		return
	case "540":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_540
		return
	case "541":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_541
		return
	case "542":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_542
		return
	case "543":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_543
		return
	case "544":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_544
		return
	case "545":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_545
		return
	case "546":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_546
		return
	case "547":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_547
		return
	case "548":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_548
		return
	case "549":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_549
		return
	case "550":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_550
		return
	case "551":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_551
		return
	case "552":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_552
		return
	case "553":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_553
		return
	case "554":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_554
		return
	case "555":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_555
		return
	case "556":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_556
		return
	case "557":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_557
		return
	case "558":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_558
		return
	case "559":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_559
		return
	case "560":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_560
		return
	case "561":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_561
		return
	case "562":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_562
		return
	case "563":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_563
		return
	case "564":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_564
		return
	case "565":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_565
		return
	case "566":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_566
		return
	case "567":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_567
		return
	case "568":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_568
		return
	case "569":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_569
		return
	case "570":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_570
		return
	case "571":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_571
		return
	case "572":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_572
		return
	case "573":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_573
		return
	case "574":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_574
		return
	case "575":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_575
		return
	case "576":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_576
		return
	case "577":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_577
		return
	case "578":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_578
		return
	case "579":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_579
		return
	case "580":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_580
		return
	case "581":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_581
		return
	case "582":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_582
		return
	case "583":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_583
		return
	case "584":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_584
		return
	case "585":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_585
		return
	case "586":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_586
		return
	case "587":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_587
		return
	case "588":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_588
		return
	case "589":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_589
		return
	case "590":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_590
		return
	case "591":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_591
		return
	case "592":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_592
		return
	case "593":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_593
		return
	case "594":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_594
		return
	case "595":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_595
		return
	case "596":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_596
		return
	case "597":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_597
		return
	case "598":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_598
		return
	case "599":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_599
		return
	case "600":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_600
		return
	case "601":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_601
		return
	case "602":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_602
		return
	case "603":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_603
		return
	case "604":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_604
		return
	case "605":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_605
		return
	case "606":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_606
		return
	case "607":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_607
		return
	case "608":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_608
		return
	case "609":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_609
		return
	case "610":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_610
		return
	case "611":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_611
		return
	case "612":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_612
		return
	case "613":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_613
		return
	case "614":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_614
		return
	case "615":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_615
		return
	case "616":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_616
		return
	case "617":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_617
		return
	case "618":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_618
		return
	case "619":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_619
		return
	case "620":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_620
		return
	case "621":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_621
		return
	case "622":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_622
		return
	case "623":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_623
		return
	case "624":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_624
		return
	case "625":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_625
		return
	case "626":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_626
		return
	case "627":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_627
		return
	case "628":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_628
		return
	case "629":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_629
		return
	case "630":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_630
		return
	case "631":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_631
		return
	case "632":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_632
		return
	case "633":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_633
		return
	case "634":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_634
		return
	case "635":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_635
		return
	case "636":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_636
		return
	case "637":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_637
		return
	case "638":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_638
		return
	case "639":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_639
		return
	case "640":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_640
		return
	case "641":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_641
		return
	case "642":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_642
		return
	case "643":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_643
		return
	case "644":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_644
		return
	case "645":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_645
		return
	case "646":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_646
		return
	case "647":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_647
		return
	case "648":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_648
		return
	case "649":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_649
		return
	case "650":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_650
		return
	case "651":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_651
		return
	case "652":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_652
		return
	case "653":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_653
		return
	case "654":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_654
		return
	case "655":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_655
		return
	case "656":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_656
		return
	case "657":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_657
		return
	case "658":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_658
		return
	case "659":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_659
		return
	case "660":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_660
		return
	case "661":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_661
		return
	case "662":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_662
		return
	case "663":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_663
		return
	case "664":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_664
		return
	case "665":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_665
		return
	case "666":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_666
		return
	case "667":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_667
		return
	case "668":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_668
		return
	case "669":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_669
		return
	case "670":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_670
		return
	case "671":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_671
		return
	case "672":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_672
		return
	case "673":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_673
		return
	case "674":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_674
		return
	case "675":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_675
		return
	case "676":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_676
		return
	case "677":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_677
		return
	case "678":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_678
		return
	case "679":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_679
		return
	case "680":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_680
		return
	case "681":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_681
		return
	case "682":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_682
		return
	case "683":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_683
		return
	case "684":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_684
		return
	case "685":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_685
		return
	case "686":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_686
		return
	case "687":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_687
		return
	case "688":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_688
		return
	case "689":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_689
		return
	case "690":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_690
		return
	case "691":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_691
		return
	case "692":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_692
		return
	case "693":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_693
		return
	case "694":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_694
		return
	case "695":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_695
		return
	case "696":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_696
		return
	case "697":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_697
		return
	case "698":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_698
		return
	case "699":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_699
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemlschannelbasetype_member_1 *Enum_CodeMLSChannelBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMLSChannelBaseType_Member_1_N_500":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_500
	case "Enum_CodeMLSChannelBaseType_Member_1_N_501":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_501
	case "Enum_CodeMLSChannelBaseType_Member_1_N_502":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_502
	case "Enum_CodeMLSChannelBaseType_Member_1_N_503":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_503
	case "Enum_CodeMLSChannelBaseType_Member_1_N_504":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_504
	case "Enum_CodeMLSChannelBaseType_Member_1_N_505":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_505
	case "Enum_CodeMLSChannelBaseType_Member_1_N_506":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_506
	case "Enum_CodeMLSChannelBaseType_Member_1_N_507":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_507
	case "Enum_CodeMLSChannelBaseType_Member_1_N_508":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_508
	case "Enum_CodeMLSChannelBaseType_Member_1_N_509":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_509
	case "Enum_CodeMLSChannelBaseType_Member_1_N_510":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_510
	case "Enum_CodeMLSChannelBaseType_Member_1_N_511":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_511
	case "Enum_CodeMLSChannelBaseType_Member_1_N_512":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_512
	case "Enum_CodeMLSChannelBaseType_Member_1_N_513":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_513
	case "Enum_CodeMLSChannelBaseType_Member_1_N_514":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_514
	case "Enum_CodeMLSChannelBaseType_Member_1_N_515":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_515
	case "Enum_CodeMLSChannelBaseType_Member_1_N_516":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_516
	case "Enum_CodeMLSChannelBaseType_Member_1_N_517":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_517
	case "Enum_CodeMLSChannelBaseType_Member_1_N_518":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_518
	case "Enum_CodeMLSChannelBaseType_Member_1_N_519":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_519
	case "Enum_CodeMLSChannelBaseType_Member_1_N_520":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_520
	case "Enum_CodeMLSChannelBaseType_Member_1_N_521":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_521
	case "Enum_CodeMLSChannelBaseType_Member_1_N_522":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_522
	case "Enum_CodeMLSChannelBaseType_Member_1_N_523":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_523
	case "Enum_CodeMLSChannelBaseType_Member_1_N_524":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_524
	case "Enum_CodeMLSChannelBaseType_Member_1_N_525":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_525
	case "Enum_CodeMLSChannelBaseType_Member_1_N_526":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_526
	case "Enum_CodeMLSChannelBaseType_Member_1_N_527":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_527
	case "Enum_CodeMLSChannelBaseType_Member_1_N_528":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_528
	case "Enum_CodeMLSChannelBaseType_Member_1_N_529":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_529
	case "Enum_CodeMLSChannelBaseType_Member_1_N_530":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_530
	case "Enum_CodeMLSChannelBaseType_Member_1_N_531":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_531
	case "Enum_CodeMLSChannelBaseType_Member_1_N_532":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_532
	case "Enum_CodeMLSChannelBaseType_Member_1_N_533":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_533
	case "Enum_CodeMLSChannelBaseType_Member_1_N_534":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_534
	case "Enum_CodeMLSChannelBaseType_Member_1_N_535":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_535
	case "Enum_CodeMLSChannelBaseType_Member_1_N_536":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_536
	case "Enum_CodeMLSChannelBaseType_Member_1_N_537":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_537
	case "Enum_CodeMLSChannelBaseType_Member_1_N_538":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_538
	case "Enum_CodeMLSChannelBaseType_Member_1_N_539":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_539
	case "Enum_CodeMLSChannelBaseType_Member_1_N_540":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_540
	case "Enum_CodeMLSChannelBaseType_Member_1_N_541":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_541
	case "Enum_CodeMLSChannelBaseType_Member_1_N_542":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_542
	case "Enum_CodeMLSChannelBaseType_Member_1_N_543":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_543
	case "Enum_CodeMLSChannelBaseType_Member_1_N_544":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_544
	case "Enum_CodeMLSChannelBaseType_Member_1_N_545":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_545
	case "Enum_CodeMLSChannelBaseType_Member_1_N_546":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_546
	case "Enum_CodeMLSChannelBaseType_Member_1_N_547":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_547
	case "Enum_CodeMLSChannelBaseType_Member_1_N_548":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_548
	case "Enum_CodeMLSChannelBaseType_Member_1_N_549":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_549
	case "Enum_CodeMLSChannelBaseType_Member_1_N_550":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_550
	case "Enum_CodeMLSChannelBaseType_Member_1_N_551":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_551
	case "Enum_CodeMLSChannelBaseType_Member_1_N_552":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_552
	case "Enum_CodeMLSChannelBaseType_Member_1_N_553":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_553
	case "Enum_CodeMLSChannelBaseType_Member_1_N_554":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_554
	case "Enum_CodeMLSChannelBaseType_Member_1_N_555":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_555
	case "Enum_CodeMLSChannelBaseType_Member_1_N_556":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_556
	case "Enum_CodeMLSChannelBaseType_Member_1_N_557":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_557
	case "Enum_CodeMLSChannelBaseType_Member_1_N_558":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_558
	case "Enum_CodeMLSChannelBaseType_Member_1_N_559":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_559
	case "Enum_CodeMLSChannelBaseType_Member_1_N_560":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_560
	case "Enum_CodeMLSChannelBaseType_Member_1_N_561":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_561
	case "Enum_CodeMLSChannelBaseType_Member_1_N_562":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_562
	case "Enum_CodeMLSChannelBaseType_Member_1_N_563":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_563
	case "Enum_CodeMLSChannelBaseType_Member_1_N_564":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_564
	case "Enum_CodeMLSChannelBaseType_Member_1_N_565":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_565
	case "Enum_CodeMLSChannelBaseType_Member_1_N_566":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_566
	case "Enum_CodeMLSChannelBaseType_Member_1_N_567":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_567
	case "Enum_CodeMLSChannelBaseType_Member_1_N_568":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_568
	case "Enum_CodeMLSChannelBaseType_Member_1_N_569":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_569
	case "Enum_CodeMLSChannelBaseType_Member_1_N_570":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_570
	case "Enum_CodeMLSChannelBaseType_Member_1_N_571":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_571
	case "Enum_CodeMLSChannelBaseType_Member_1_N_572":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_572
	case "Enum_CodeMLSChannelBaseType_Member_1_N_573":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_573
	case "Enum_CodeMLSChannelBaseType_Member_1_N_574":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_574
	case "Enum_CodeMLSChannelBaseType_Member_1_N_575":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_575
	case "Enum_CodeMLSChannelBaseType_Member_1_N_576":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_576
	case "Enum_CodeMLSChannelBaseType_Member_1_N_577":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_577
	case "Enum_CodeMLSChannelBaseType_Member_1_N_578":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_578
	case "Enum_CodeMLSChannelBaseType_Member_1_N_579":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_579
	case "Enum_CodeMLSChannelBaseType_Member_1_N_580":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_580
	case "Enum_CodeMLSChannelBaseType_Member_1_N_581":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_581
	case "Enum_CodeMLSChannelBaseType_Member_1_N_582":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_582
	case "Enum_CodeMLSChannelBaseType_Member_1_N_583":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_583
	case "Enum_CodeMLSChannelBaseType_Member_1_N_584":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_584
	case "Enum_CodeMLSChannelBaseType_Member_1_N_585":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_585
	case "Enum_CodeMLSChannelBaseType_Member_1_N_586":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_586
	case "Enum_CodeMLSChannelBaseType_Member_1_N_587":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_587
	case "Enum_CodeMLSChannelBaseType_Member_1_N_588":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_588
	case "Enum_CodeMLSChannelBaseType_Member_1_N_589":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_589
	case "Enum_CodeMLSChannelBaseType_Member_1_N_590":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_590
	case "Enum_CodeMLSChannelBaseType_Member_1_N_591":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_591
	case "Enum_CodeMLSChannelBaseType_Member_1_N_592":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_592
	case "Enum_CodeMLSChannelBaseType_Member_1_N_593":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_593
	case "Enum_CodeMLSChannelBaseType_Member_1_N_594":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_594
	case "Enum_CodeMLSChannelBaseType_Member_1_N_595":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_595
	case "Enum_CodeMLSChannelBaseType_Member_1_N_596":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_596
	case "Enum_CodeMLSChannelBaseType_Member_1_N_597":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_597
	case "Enum_CodeMLSChannelBaseType_Member_1_N_598":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_598
	case "Enum_CodeMLSChannelBaseType_Member_1_N_599":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_599
	case "Enum_CodeMLSChannelBaseType_Member_1_N_600":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_600
	case "Enum_CodeMLSChannelBaseType_Member_1_N_601":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_601
	case "Enum_CodeMLSChannelBaseType_Member_1_N_602":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_602
	case "Enum_CodeMLSChannelBaseType_Member_1_N_603":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_603
	case "Enum_CodeMLSChannelBaseType_Member_1_N_604":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_604
	case "Enum_CodeMLSChannelBaseType_Member_1_N_605":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_605
	case "Enum_CodeMLSChannelBaseType_Member_1_N_606":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_606
	case "Enum_CodeMLSChannelBaseType_Member_1_N_607":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_607
	case "Enum_CodeMLSChannelBaseType_Member_1_N_608":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_608
	case "Enum_CodeMLSChannelBaseType_Member_1_N_609":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_609
	case "Enum_CodeMLSChannelBaseType_Member_1_N_610":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_610
	case "Enum_CodeMLSChannelBaseType_Member_1_N_611":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_611
	case "Enum_CodeMLSChannelBaseType_Member_1_N_612":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_612
	case "Enum_CodeMLSChannelBaseType_Member_1_N_613":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_613
	case "Enum_CodeMLSChannelBaseType_Member_1_N_614":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_614
	case "Enum_CodeMLSChannelBaseType_Member_1_N_615":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_615
	case "Enum_CodeMLSChannelBaseType_Member_1_N_616":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_616
	case "Enum_CodeMLSChannelBaseType_Member_1_N_617":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_617
	case "Enum_CodeMLSChannelBaseType_Member_1_N_618":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_618
	case "Enum_CodeMLSChannelBaseType_Member_1_N_619":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_619
	case "Enum_CodeMLSChannelBaseType_Member_1_N_620":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_620
	case "Enum_CodeMLSChannelBaseType_Member_1_N_621":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_621
	case "Enum_CodeMLSChannelBaseType_Member_1_N_622":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_622
	case "Enum_CodeMLSChannelBaseType_Member_1_N_623":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_623
	case "Enum_CodeMLSChannelBaseType_Member_1_N_624":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_624
	case "Enum_CodeMLSChannelBaseType_Member_1_N_625":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_625
	case "Enum_CodeMLSChannelBaseType_Member_1_N_626":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_626
	case "Enum_CodeMLSChannelBaseType_Member_1_N_627":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_627
	case "Enum_CodeMLSChannelBaseType_Member_1_N_628":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_628
	case "Enum_CodeMLSChannelBaseType_Member_1_N_629":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_629
	case "Enum_CodeMLSChannelBaseType_Member_1_N_630":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_630
	case "Enum_CodeMLSChannelBaseType_Member_1_N_631":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_631
	case "Enum_CodeMLSChannelBaseType_Member_1_N_632":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_632
	case "Enum_CodeMLSChannelBaseType_Member_1_N_633":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_633
	case "Enum_CodeMLSChannelBaseType_Member_1_N_634":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_634
	case "Enum_CodeMLSChannelBaseType_Member_1_N_635":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_635
	case "Enum_CodeMLSChannelBaseType_Member_1_N_636":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_636
	case "Enum_CodeMLSChannelBaseType_Member_1_N_637":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_637
	case "Enum_CodeMLSChannelBaseType_Member_1_N_638":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_638
	case "Enum_CodeMLSChannelBaseType_Member_1_N_639":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_639
	case "Enum_CodeMLSChannelBaseType_Member_1_N_640":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_640
	case "Enum_CodeMLSChannelBaseType_Member_1_N_641":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_641
	case "Enum_CodeMLSChannelBaseType_Member_1_N_642":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_642
	case "Enum_CodeMLSChannelBaseType_Member_1_N_643":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_643
	case "Enum_CodeMLSChannelBaseType_Member_1_N_644":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_644
	case "Enum_CodeMLSChannelBaseType_Member_1_N_645":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_645
	case "Enum_CodeMLSChannelBaseType_Member_1_N_646":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_646
	case "Enum_CodeMLSChannelBaseType_Member_1_N_647":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_647
	case "Enum_CodeMLSChannelBaseType_Member_1_N_648":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_648
	case "Enum_CodeMLSChannelBaseType_Member_1_N_649":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_649
	case "Enum_CodeMLSChannelBaseType_Member_1_N_650":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_650
	case "Enum_CodeMLSChannelBaseType_Member_1_N_651":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_651
	case "Enum_CodeMLSChannelBaseType_Member_1_N_652":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_652
	case "Enum_CodeMLSChannelBaseType_Member_1_N_653":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_653
	case "Enum_CodeMLSChannelBaseType_Member_1_N_654":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_654
	case "Enum_CodeMLSChannelBaseType_Member_1_N_655":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_655
	case "Enum_CodeMLSChannelBaseType_Member_1_N_656":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_656
	case "Enum_CodeMLSChannelBaseType_Member_1_N_657":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_657
	case "Enum_CodeMLSChannelBaseType_Member_1_N_658":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_658
	case "Enum_CodeMLSChannelBaseType_Member_1_N_659":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_659
	case "Enum_CodeMLSChannelBaseType_Member_1_N_660":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_660
	case "Enum_CodeMLSChannelBaseType_Member_1_N_661":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_661
	case "Enum_CodeMLSChannelBaseType_Member_1_N_662":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_662
	case "Enum_CodeMLSChannelBaseType_Member_1_N_663":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_663
	case "Enum_CodeMLSChannelBaseType_Member_1_N_664":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_664
	case "Enum_CodeMLSChannelBaseType_Member_1_N_665":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_665
	case "Enum_CodeMLSChannelBaseType_Member_1_N_666":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_666
	case "Enum_CodeMLSChannelBaseType_Member_1_N_667":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_667
	case "Enum_CodeMLSChannelBaseType_Member_1_N_668":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_668
	case "Enum_CodeMLSChannelBaseType_Member_1_N_669":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_669
	case "Enum_CodeMLSChannelBaseType_Member_1_N_670":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_670
	case "Enum_CodeMLSChannelBaseType_Member_1_N_671":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_671
	case "Enum_CodeMLSChannelBaseType_Member_1_N_672":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_672
	case "Enum_CodeMLSChannelBaseType_Member_1_N_673":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_673
	case "Enum_CodeMLSChannelBaseType_Member_1_N_674":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_674
	case "Enum_CodeMLSChannelBaseType_Member_1_N_675":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_675
	case "Enum_CodeMLSChannelBaseType_Member_1_N_676":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_676
	case "Enum_CodeMLSChannelBaseType_Member_1_N_677":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_677
	case "Enum_CodeMLSChannelBaseType_Member_1_N_678":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_678
	case "Enum_CodeMLSChannelBaseType_Member_1_N_679":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_679
	case "Enum_CodeMLSChannelBaseType_Member_1_N_680":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_680
	case "Enum_CodeMLSChannelBaseType_Member_1_N_681":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_681
	case "Enum_CodeMLSChannelBaseType_Member_1_N_682":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_682
	case "Enum_CodeMLSChannelBaseType_Member_1_N_683":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_683
	case "Enum_CodeMLSChannelBaseType_Member_1_N_684":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_684
	case "Enum_CodeMLSChannelBaseType_Member_1_N_685":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_685
	case "Enum_CodeMLSChannelBaseType_Member_1_N_686":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_686
	case "Enum_CodeMLSChannelBaseType_Member_1_N_687":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_687
	case "Enum_CodeMLSChannelBaseType_Member_1_N_688":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_688
	case "Enum_CodeMLSChannelBaseType_Member_1_N_689":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_689
	case "Enum_CodeMLSChannelBaseType_Member_1_N_690":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_690
	case "Enum_CodeMLSChannelBaseType_Member_1_N_691":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_691
	case "Enum_CodeMLSChannelBaseType_Member_1_N_692":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_692
	case "Enum_CodeMLSChannelBaseType_Member_1_N_693":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_693
	case "Enum_CodeMLSChannelBaseType_Member_1_N_694":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_694
	case "Enum_CodeMLSChannelBaseType_Member_1_N_695":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_695
	case "Enum_CodeMLSChannelBaseType_Member_1_N_696":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_696
	case "Enum_CodeMLSChannelBaseType_Member_1_N_697":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_697
	case "Enum_CodeMLSChannelBaseType_Member_1_N_698":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_698
	case "Enum_CodeMLSChannelBaseType_Member_1_N_699":
		*enum_codemlschannelbasetype_member_1 = Enum_CodeMLSChannelBaseType_Member_1_N_699
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemlschannelbasetype_member_1 *Enum_CodeMLSChannelBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemlschannelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMLSChannelBaseType_Member_1_N_500:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_500"
	case Enum_CodeMLSChannelBaseType_Member_1_N_501:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_501"
	case Enum_CodeMLSChannelBaseType_Member_1_N_502:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_502"
	case Enum_CodeMLSChannelBaseType_Member_1_N_503:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_503"
	case Enum_CodeMLSChannelBaseType_Member_1_N_504:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_504"
	case Enum_CodeMLSChannelBaseType_Member_1_N_505:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_505"
	case Enum_CodeMLSChannelBaseType_Member_1_N_506:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_506"
	case Enum_CodeMLSChannelBaseType_Member_1_N_507:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_507"
	case Enum_CodeMLSChannelBaseType_Member_1_N_508:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_508"
	case Enum_CodeMLSChannelBaseType_Member_1_N_509:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_509"
	case Enum_CodeMLSChannelBaseType_Member_1_N_510:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_510"
	case Enum_CodeMLSChannelBaseType_Member_1_N_511:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_511"
	case Enum_CodeMLSChannelBaseType_Member_1_N_512:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_512"
	case Enum_CodeMLSChannelBaseType_Member_1_N_513:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_513"
	case Enum_CodeMLSChannelBaseType_Member_1_N_514:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_514"
	case Enum_CodeMLSChannelBaseType_Member_1_N_515:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_515"
	case Enum_CodeMLSChannelBaseType_Member_1_N_516:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_516"
	case Enum_CodeMLSChannelBaseType_Member_1_N_517:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_517"
	case Enum_CodeMLSChannelBaseType_Member_1_N_518:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_518"
	case Enum_CodeMLSChannelBaseType_Member_1_N_519:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_519"
	case Enum_CodeMLSChannelBaseType_Member_1_N_520:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_520"
	case Enum_CodeMLSChannelBaseType_Member_1_N_521:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_521"
	case Enum_CodeMLSChannelBaseType_Member_1_N_522:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_522"
	case Enum_CodeMLSChannelBaseType_Member_1_N_523:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_523"
	case Enum_CodeMLSChannelBaseType_Member_1_N_524:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_524"
	case Enum_CodeMLSChannelBaseType_Member_1_N_525:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_525"
	case Enum_CodeMLSChannelBaseType_Member_1_N_526:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_526"
	case Enum_CodeMLSChannelBaseType_Member_1_N_527:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_527"
	case Enum_CodeMLSChannelBaseType_Member_1_N_528:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_528"
	case Enum_CodeMLSChannelBaseType_Member_1_N_529:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_529"
	case Enum_CodeMLSChannelBaseType_Member_1_N_530:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_530"
	case Enum_CodeMLSChannelBaseType_Member_1_N_531:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_531"
	case Enum_CodeMLSChannelBaseType_Member_1_N_532:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_532"
	case Enum_CodeMLSChannelBaseType_Member_1_N_533:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_533"
	case Enum_CodeMLSChannelBaseType_Member_1_N_534:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_534"
	case Enum_CodeMLSChannelBaseType_Member_1_N_535:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_535"
	case Enum_CodeMLSChannelBaseType_Member_1_N_536:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_536"
	case Enum_CodeMLSChannelBaseType_Member_1_N_537:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_537"
	case Enum_CodeMLSChannelBaseType_Member_1_N_538:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_538"
	case Enum_CodeMLSChannelBaseType_Member_1_N_539:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_539"
	case Enum_CodeMLSChannelBaseType_Member_1_N_540:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_540"
	case Enum_CodeMLSChannelBaseType_Member_1_N_541:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_541"
	case Enum_CodeMLSChannelBaseType_Member_1_N_542:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_542"
	case Enum_CodeMLSChannelBaseType_Member_1_N_543:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_543"
	case Enum_CodeMLSChannelBaseType_Member_1_N_544:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_544"
	case Enum_CodeMLSChannelBaseType_Member_1_N_545:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_545"
	case Enum_CodeMLSChannelBaseType_Member_1_N_546:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_546"
	case Enum_CodeMLSChannelBaseType_Member_1_N_547:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_547"
	case Enum_CodeMLSChannelBaseType_Member_1_N_548:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_548"
	case Enum_CodeMLSChannelBaseType_Member_1_N_549:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_549"
	case Enum_CodeMLSChannelBaseType_Member_1_N_550:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_550"
	case Enum_CodeMLSChannelBaseType_Member_1_N_551:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_551"
	case Enum_CodeMLSChannelBaseType_Member_1_N_552:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_552"
	case Enum_CodeMLSChannelBaseType_Member_1_N_553:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_553"
	case Enum_CodeMLSChannelBaseType_Member_1_N_554:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_554"
	case Enum_CodeMLSChannelBaseType_Member_1_N_555:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_555"
	case Enum_CodeMLSChannelBaseType_Member_1_N_556:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_556"
	case Enum_CodeMLSChannelBaseType_Member_1_N_557:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_557"
	case Enum_CodeMLSChannelBaseType_Member_1_N_558:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_558"
	case Enum_CodeMLSChannelBaseType_Member_1_N_559:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_559"
	case Enum_CodeMLSChannelBaseType_Member_1_N_560:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_560"
	case Enum_CodeMLSChannelBaseType_Member_1_N_561:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_561"
	case Enum_CodeMLSChannelBaseType_Member_1_N_562:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_562"
	case Enum_CodeMLSChannelBaseType_Member_1_N_563:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_563"
	case Enum_CodeMLSChannelBaseType_Member_1_N_564:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_564"
	case Enum_CodeMLSChannelBaseType_Member_1_N_565:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_565"
	case Enum_CodeMLSChannelBaseType_Member_1_N_566:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_566"
	case Enum_CodeMLSChannelBaseType_Member_1_N_567:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_567"
	case Enum_CodeMLSChannelBaseType_Member_1_N_568:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_568"
	case Enum_CodeMLSChannelBaseType_Member_1_N_569:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_569"
	case Enum_CodeMLSChannelBaseType_Member_1_N_570:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_570"
	case Enum_CodeMLSChannelBaseType_Member_1_N_571:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_571"
	case Enum_CodeMLSChannelBaseType_Member_1_N_572:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_572"
	case Enum_CodeMLSChannelBaseType_Member_1_N_573:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_573"
	case Enum_CodeMLSChannelBaseType_Member_1_N_574:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_574"
	case Enum_CodeMLSChannelBaseType_Member_1_N_575:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_575"
	case Enum_CodeMLSChannelBaseType_Member_1_N_576:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_576"
	case Enum_CodeMLSChannelBaseType_Member_1_N_577:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_577"
	case Enum_CodeMLSChannelBaseType_Member_1_N_578:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_578"
	case Enum_CodeMLSChannelBaseType_Member_1_N_579:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_579"
	case Enum_CodeMLSChannelBaseType_Member_1_N_580:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_580"
	case Enum_CodeMLSChannelBaseType_Member_1_N_581:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_581"
	case Enum_CodeMLSChannelBaseType_Member_1_N_582:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_582"
	case Enum_CodeMLSChannelBaseType_Member_1_N_583:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_583"
	case Enum_CodeMLSChannelBaseType_Member_1_N_584:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_584"
	case Enum_CodeMLSChannelBaseType_Member_1_N_585:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_585"
	case Enum_CodeMLSChannelBaseType_Member_1_N_586:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_586"
	case Enum_CodeMLSChannelBaseType_Member_1_N_587:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_587"
	case Enum_CodeMLSChannelBaseType_Member_1_N_588:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_588"
	case Enum_CodeMLSChannelBaseType_Member_1_N_589:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_589"
	case Enum_CodeMLSChannelBaseType_Member_1_N_590:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_590"
	case Enum_CodeMLSChannelBaseType_Member_1_N_591:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_591"
	case Enum_CodeMLSChannelBaseType_Member_1_N_592:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_592"
	case Enum_CodeMLSChannelBaseType_Member_1_N_593:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_593"
	case Enum_CodeMLSChannelBaseType_Member_1_N_594:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_594"
	case Enum_CodeMLSChannelBaseType_Member_1_N_595:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_595"
	case Enum_CodeMLSChannelBaseType_Member_1_N_596:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_596"
	case Enum_CodeMLSChannelBaseType_Member_1_N_597:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_597"
	case Enum_CodeMLSChannelBaseType_Member_1_N_598:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_598"
	case Enum_CodeMLSChannelBaseType_Member_1_N_599:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_599"
	case Enum_CodeMLSChannelBaseType_Member_1_N_600:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_600"
	case Enum_CodeMLSChannelBaseType_Member_1_N_601:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_601"
	case Enum_CodeMLSChannelBaseType_Member_1_N_602:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_602"
	case Enum_CodeMLSChannelBaseType_Member_1_N_603:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_603"
	case Enum_CodeMLSChannelBaseType_Member_1_N_604:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_604"
	case Enum_CodeMLSChannelBaseType_Member_1_N_605:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_605"
	case Enum_CodeMLSChannelBaseType_Member_1_N_606:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_606"
	case Enum_CodeMLSChannelBaseType_Member_1_N_607:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_607"
	case Enum_CodeMLSChannelBaseType_Member_1_N_608:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_608"
	case Enum_CodeMLSChannelBaseType_Member_1_N_609:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_609"
	case Enum_CodeMLSChannelBaseType_Member_1_N_610:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_610"
	case Enum_CodeMLSChannelBaseType_Member_1_N_611:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_611"
	case Enum_CodeMLSChannelBaseType_Member_1_N_612:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_612"
	case Enum_CodeMLSChannelBaseType_Member_1_N_613:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_613"
	case Enum_CodeMLSChannelBaseType_Member_1_N_614:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_614"
	case Enum_CodeMLSChannelBaseType_Member_1_N_615:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_615"
	case Enum_CodeMLSChannelBaseType_Member_1_N_616:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_616"
	case Enum_CodeMLSChannelBaseType_Member_1_N_617:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_617"
	case Enum_CodeMLSChannelBaseType_Member_1_N_618:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_618"
	case Enum_CodeMLSChannelBaseType_Member_1_N_619:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_619"
	case Enum_CodeMLSChannelBaseType_Member_1_N_620:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_620"
	case Enum_CodeMLSChannelBaseType_Member_1_N_621:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_621"
	case Enum_CodeMLSChannelBaseType_Member_1_N_622:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_622"
	case Enum_CodeMLSChannelBaseType_Member_1_N_623:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_623"
	case Enum_CodeMLSChannelBaseType_Member_1_N_624:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_624"
	case Enum_CodeMLSChannelBaseType_Member_1_N_625:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_625"
	case Enum_CodeMLSChannelBaseType_Member_1_N_626:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_626"
	case Enum_CodeMLSChannelBaseType_Member_1_N_627:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_627"
	case Enum_CodeMLSChannelBaseType_Member_1_N_628:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_628"
	case Enum_CodeMLSChannelBaseType_Member_1_N_629:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_629"
	case Enum_CodeMLSChannelBaseType_Member_1_N_630:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_630"
	case Enum_CodeMLSChannelBaseType_Member_1_N_631:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_631"
	case Enum_CodeMLSChannelBaseType_Member_1_N_632:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_632"
	case Enum_CodeMLSChannelBaseType_Member_1_N_633:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_633"
	case Enum_CodeMLSChannelBaseType_Member_1_N_634:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_634"
	case Enum_CodeMLSChannelBaseType_Member_1_N_635:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_635"
	case Enum_CodeMLSChannelBaseType_Member_1_N_636:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_636"
	case Enum_CodeMLSChannelBaseType_Member_1_N_637:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_637"
	case Enum_CodeMLSChannelBaseType_Member_1_N_638:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_638"
	case Enum_CodeMLSChannelBaseType_Member_1_N_639:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_639"
	case Enum_CodeMLSChannelBaseType_Member_1_N_640:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_640"
	case Enum_CodeMLSChannelBaseType_Member_1_N_641:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_641"
	case Enum_CodeMLSChannelBaseType_Member_1_N_642:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_642"
	case Enum_CodeMLSChannelBaseType_Member_1_N_643:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_643"
	case Enum_CodeMLSChannelBaseType_Member_1_N_644:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_644"
	case Enum_CodeMLSChannelBaseType_Member_1_N_645:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_645"
	case Enum_CodeMLSChannelBaseType_Member_1_N_646:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_646"
	case Enum_CodeMLSChannelBaseType_Member_1_N_647:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_647"
	case Enum_CodeMLSChannelBaseType_Member_1_N_648:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_648"
	case Enum_CodeMLSChannelBaseType_Member_1_N_649:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_649"
	case Enum_CodeMLSChannelBaseType_Member_1_N_650:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_650"
	case Enum_CodeMLSChannelBaseType_Member_1_N_651:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_651"
	case Enum_CodeMLSChannelBaseType_Member_1_N_652:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_652"
	case Enum_CodeMLSChannelBaseType_Member_1_N_653:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_653"
	case Enum_CodeMLSChannelBaseType_Member_1_N_654:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_654"
	case Enum_CodeMLSChannelBaseType_Member_1_N_655:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_655"
	case Enum_CodeMLSChannelBaseType_Member_1_N_656:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_656"
	case Enum_CodeMLSChannelBaseType_Member_1_N_657:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_657"
	case Enum_CodeMLSChannelBaseType_Member_1_N_658:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_658"
	case Enum_CodeMLSChannelBaseType_Member_1_N_659:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_659"
	case Enum_CodeMLSChannelBaseType_Member_1_N_660:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_660"
	case Enum_CodeMLSChannelBaseType_Member_1_N_661:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_661"
	case Enum_CodeMLSChannelBaseType_Member_1_N_662:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_662"
	case Enum_CodeMLSChannelBaseType_Member_1_N_663:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_663"
	case Enum_CodeMLSChannelBaseType_Member_1_N_664:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_664"
	case Enum_CodeMLSChannelBaseType_Member_1_N_665:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_665"
	case Enum_CodeMLSChannelBaseType_Member_1_N_666:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_666"
	case Enum_CodeMLSChannelBaseType_Member_1_N_667:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_667"
	case Enum_CodeMLSChannelBaseType_Member_1_N_668:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_668"
	case Enum_CodeMLSChannelBaseType_Member_1_N_669:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_669"
	case Enum_CodeMLSChannelBaseType_Member_1_N_670:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_670"
	case Enum_CodeMLSChannelBaseType_Member_1_N_671:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_671"
	case Enum_CodeMLSChannelBaseType_Member_1_N_672:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_672"
	case Enum_CodeMLSChannelBaseType_Member_1_N_673:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_673"
	case Enum_CodeMLSChannelBaseType_Member_1_N_674:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_674"
	case Enum_CodeMLSChannelBaseType_Member_1_N_675:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_675"
	case Enum_CodeMLSChannelBaseType_Member_1_N_676:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_676"
	case Enum_CodeMLSChannelBaseType_Member_1_N_677:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_677"
	case Enum_CodeMLSChannelBaseType_Member_1_N_678:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_678"
	case Enum_CodeMLSChannelBaseType_Member_1_N_679:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_679"
	case Enum_CodeMLSChannelBaseType_Member_1_N_680:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_680"
	case Enum_CodeMLSChannelBaseType_Member_1_N_681:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_681"
	case Enum_CodeMLSChannelBaseType_Member_1_N_682:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_682"
	case Enum_CodeMLSChannelBaseType_Member_1_N_683:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_683"
	case Enum_CodeMLSChannelBaseType_Member_1_N_684:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_684"
	case Enum_CodeMLSChannelBaseType_Member_1_N_685:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_685"
	case Enum_CodeMLSChannelBaseType_Member_1_N_686:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_686"
	case Enum_CodeMLSChannelBaseType_Member_1_N_687:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_687"
	case Enum_CodeMLSChannelBaseType_Member_1_N_688:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_688"
	case Enum_CodeMLSChannelBaseType_Member_1_N_689:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_689"
	case Enum_CodeMLSChannelBaseType_Member_1_N_690:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_690"
	case Enum_CodeMLSChannelBaseType_Member_1_N_691:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_691"
	case Enum_CodeMLSChannelBaseType_Member_1_N_692:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_692"
	case Enum_CodeMLSChannelBaseType_Member_1_N_693:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_693"
	case Enum_CodeMLSChannelBaseType_Member_1_N_694:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_694"
	case Enum_CodeMLSChannelBaseType_Member_1_N_695:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_695"
	case Enum_CodeMLSChannelBaseType_Member_1_N_696:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_696"
	case Enum_CodeMLSChannelBaseType_Member_1_N_697:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_697"
	case Enum_CodeMLSChannelBaseType_Member_1_N_698:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_698"
	case Enum_CodeMLSChannelBaseType_Member_1_N_699:
		res = "Enum_CodeMLSChannelBaseType_Member_1_N_699"
	}
	return
}

func (enum_codemlschannelbasetype_member_1 Enum_CodeMLSChannelBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_500")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_501")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_502")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_503")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_504")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_505")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_506")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_507")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_508")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_509")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_510")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_511")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_512")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_513")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_514")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_515")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_516")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_517")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_518")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_519")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_520")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_521")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_522")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_523")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_524")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_525")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_526")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_527")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_528")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_529")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_530")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_531")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_532")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_533")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_534")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_535")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_536")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_537")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_538")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_539")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_540")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_541")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_542")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_543")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_544")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_545")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_546")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_547")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_548")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_549")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_550")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_551")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_552")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_553")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_554")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_555")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_556")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_557")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_558")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_559")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_560")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_561")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_562")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_563")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_564")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_565")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_566")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_567")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_568")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_569")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_570")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_571")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_572")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_573")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_574")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_575")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_576")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_577")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_578")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_579")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_580")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_581")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_582")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_583")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_584")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_585")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_586")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_587")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_588")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_589")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_590")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_591")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_592")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_593")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_594")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_595")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_596")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_597")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_598")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_599")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_600")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_601")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_602")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_603")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_604")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_605")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_606")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_607")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_608")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_609")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_610")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_611")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_612")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_613")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_614")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_615")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_616")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_617")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_618")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_619")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_620")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_621")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_622")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_623")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_624")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_625")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_626")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_627")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_628")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_629")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_630")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_631")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_632")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_633")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_634")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_635")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_636")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_637")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_638")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_639")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_640")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_641")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_642")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_643")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_644")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_645")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_646")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_647")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_648")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_649")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_650")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_651")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_652")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_653")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_654")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_655")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_656")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_657")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_658")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_659")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_660")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_661")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_662")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_663")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_664")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_665")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_666")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_667")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_668")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_669")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_670")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_671")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_672")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_673")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_674")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_675")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_676")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_677")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_678")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_679")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_680")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_681")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_682")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_683")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_684")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_685")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_686")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_687")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_688")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_689")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_690")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_691")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_692")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_693")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_694")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_695")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_696")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_697")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_698")
	res = append(res, "Enum_CodeMLSChannelBaseType_Member_1_N_699")

	return
}

func (enum_codemlschannelbasetype_member_1 Enum_CodeMLSChannelBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "500")
	res = append(res, "501")
	res = append(res, "502")
	res = append(res, "503")
	res = append(res, "504")
	res = append(res, "505")
	res = append(res, "506")
	res = append(res, "507")
	res = append(res, "508")
	res = append(res, "509")
	res = append(res, "510")
	res = append(res, "511")
	res = append(res, "512")
	res = append(res, "513")
	res = append(res, "514")
	res = append(res, "515")
	res = append(res, "516")
	res = append(res, "517")
	res = append(res, "518")
	res = append(res, "519")
	res = append(res, "520")
	res = append(res, "521")
	res = append(res, "522")
	res = append(res, "523")
	res = append(res, "524")
	res = append(res, "525")
	res = append(res, "526")
	res = append(res, "527")
	res = append(res, "528")
	res = append(res, "529")
	res = append(res, "530")
	res = append(res, "531")
	res = append(res, "532")
	res = append(res, "533")
	res = append(res, "534")
	res = append(res, "535")
	res = append(res, "536")
	res = append(res, "537")
	res = append(res, "538")
	res = append(res, "539")
	res = append(res, "540")
	res = append(res, "541")
	res = append(res, "542")
	res = append(res, "543")
	res = append(res, "544")
	res = append(res, "545")
	res = append(res, "546")
	res = append(res, "547")
	res = append(res, "548")
	res = append(res, "549")
	res = append(res, "550")
	res = append(res, "551")
	res = append(res, "552")
	res = append(res, "553")
	res = append(res, "554")
	res = append(res, "555")
	res = append(res, "556")
	res = append(res, "557")
	res = append(res, "558")
	res = append(res, "559")
	res = append(res, "560")
	res = append(res, "561")
	res = append(res, "562")
	res = append(res, "563")
	res = append(res, "564")
	res = append(res, "565")
	res = append(res, "566")
	res = append(res, "567")
	res = append(res, "568")
	res = append(res, "569")
	res = append(res, "570")
	res = append(res, "571")
	res = append(res, "572")
	res = append(res, "573")
	res = append(res, "574")
	res = append(res, "575")
	res = append(res, "576")
	res = append(res, "577")
	res = append(res, "578")
	res = append(res, "579")
	res = append(res, "580")
	res = append(res, "581")
	res = append(res, "582")
	res = append(res, "583")
	res = append(res, "584")
	res = append(res, "585")
	res = append(res, "586")
	res = append(res, "587")
	res = append(res, "588")
	res = append(res, "589")
	res = append(res, "590")
	res = append(res, "591")
	res = append(res, "592")
	res = append(res, "593")
	res = append(res, "594")
	res = append(res, "595")
	res = append(res, "596")
	res = append(res, "597")
	res = append(res, "598")
	res = append(res, "599")
	res = append(res, "600")
	res = append(res, "601")
	res = append(res, "602")
	res = append(res, "603")
	res = append(res, "604")
	res = append(res, "605")
	res = append(res, "606")
	res = append(res, "607")
	res = append(res, "608")
	res = append(res, "609")
	res = append(res, "610")
	res = append(res, "611")
	res = append(res, "612")
	res = append(res, "613")
	res = append(res, "614")
	res = append(res, "615")
	res = append(res, "616")
	res = append(res, "617")
	res = append(res, "618")
	res = append(res, "619")
	res = append(res, "620")
	res = append(res, "621")
	res = append(res, "622")
	res = append(res, "623")
	res = append(res, "624")
	res = append(res, "625")
	res = append(res, "626")
	res = append(res, "627")
	res = append(res, "628")
	res = append(res, "629")
	res = append(res, "630")
	res = append(res, "631")
	res = append(res, "632")
	res = append(res, "633")
	res = append(res, "634")
	res = append(res, "635")
	res = append(res, "636")
	res = append(res, "637")
	res = append(res, "638")
	res = append(res, "639")
	res = append(res, "640")
	res = append(res, "641")
	res = append(res, "642")
	res = append(res, "643")
	res = append(res, "644")
	res = append(res, "645")
	res = append(res, "646")
	res = append(res, "647")
	res = append(res, "648")
	res = append(res, "649")
	res = append(res, "650")
	res = append(res, "651")
	res = append(res, "652")
	res = append(res, "653")
	res = append(res, "654")
	res = append(res, "655")
	res = append(res, "656")
	res = append(res, "657")
	res = append(res, "658")
	res = append(res, "659")
	res = append(res, "660")
	res = append(res, "661")
	res = append(res, "662")
	res = append(res, "663")
	res = append(res, "664")
	res = append(res, "665")
	res = append(res, "666")
	res = append(res, "667")
	res = append(res, "668")
	res = append(res, "669")
	res = append(res, "670")
	res = append(res, "671")
	res = append(res, "672")
	res = append(res, "673")
	res = append(res, "674")
	res = append(res, "675")
	res = append(res, "676")
	res = append(res, "677")
	res = append(res, "678")
	res = append(res, "679")
	res = append(res, "680")
	res = append(res, "681")
	res = append(res, "682")
	res = append(res, "683")
	res = append(res, "684")
	res = append(res, "685")
	res = append(res, "686")
	res = append(res, "687")
	res = append(res, "688")
	res = append(res, "689")
	res = append(res, "690")
	res = append(res, "691")
	res = append(res, "692")
	res = append(res, "693")
	res = append(res, "694")
	res = append(res, "695")
	res = append(res, "696")
	res = append(res, "697")
	res = append(res, "698")
	res = append(res, "699")

	return
}

// Utility function for Enum_CodeMarkerBeaconSignalBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codemarkerbeaconsignalbasetype_member_1 Enum_CodeMarkerBeaconSignalBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codemarkerbeaconsignalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN:
		res = "FAN"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN:
		res = "LOW_PWR_FAN"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z:
		res = "Z"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES:
		res = "BONES"
	}
	return
}

func (enum_codemarkerbeaconsignalbasetype_member_1 *Enum_CodeMarkerBeaconSignalBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FAN":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN
		return
	case "LOW_PWR_FAN":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN
		return
	case "Z":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z
		return
	case "BONES":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codemarkerbeaconsignalbasetype_member_1 *Enum_CodeMarkerBeaconSignalBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN
	case "Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN
	case "Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z
	case "Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES":
		*enum_codemarkerbeaconsignalbasetype_member_1 = Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codemarkerbeaconsignalbasetype_member_1 *Enum_CodeMarkerBeaconSignalBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codemarkerbeaconsignalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN:
		res = "Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN:
		res = "Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z:
		res = "Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z"
	case Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES:
		res = "Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES"
	}
	return
}

func (enum_codemarkerbeaconsignalbasetype_member_1 Enum_CodeMarkerBeaconSignalBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeMarkerBeaconSignalBaseType_Member_1_FAN")
	res = append(res, "Enum_CodeMarkerBeaconSignalBaseType_Member_1_LOW_PWR_FAN")
	res = append(res, "Enum_CodeMarkerBeaconSignalBaseType_Member_1_Z")
	res = append(res, "Enum_CodeMarkerBeaconSignalBaseType_Member_1_BONES")

	return
}

func (enum_codemarkerbeaconsignalbasetype_member_1 Enum_CodeMarkerBeaconSignalBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FAN")
	res = append(res, "LOW_PWR_FAN")
	res = append(res, "Z")
	res = append(res, "BONES")

	return
}

// Utility function for Enum_CodeNDBUsageBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codendbusagebasetype_member_1 Enum_CodeNDBUsageBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codendbusagebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNDBUsageBaseType_Member_1_ENR:
		res = "ENR"
	case Enum_CodeNDBUsageBaseType_Member_1_L:
		res = "L"
	case Enum_CodeNDBUsageBaseType_Member_1_MAR:
		res = "MAR"
	}
	return
}

func (enum_codendbusagebasetype_member_1 *Enum_CodeNDBUsageBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENR":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_ENR
		return
	case "L":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_L
		return
	case "MAR":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_MAR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codendbusagebasetype_member_1 *Enum_CodeNDBUsageBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNDBUsageBaseType_Member_1_ENR":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_ENR
	case "Enum_CodeNDBUsageBaseType_Member_1_L":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_L
	case "Enum_CodeNDBUsageBaseType_Member_1_MAR":
		*enum_codendbusagebasetype_member_1 = Enum_CodeNDBUsageBaseType_Member_1_MAR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codendbusagebasetype_member_1 *Enum_CodeNDBUsageBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codendbusagebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNDBUsageBaseType_Member_1_ENR:
		res = "Enum_CodeNDBUsageBaseType_Member_1_ENR"
	case Enum_CodeNDBUsageBaseType_Member_1_L:
		res = "Enum_CodeNDBUsageBaseType_Member_1_L"
	case Enum_CodeNDBUsageBaseType_Member_1_MAR:
		res = "Enum_CodeNDBUsageBaseType_Member_1_MAR"
	}
	return
}

func (enum_codendbusagebasetype_member_1 Enum_CodeNDBUsageBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNDBUsageBaseType_Member_1_ENR")
	res = append(res, "Enum_CodeNDBUsageBaseType_Member_1_L")
	res = append(res, "Enum_CodeNDBUsageBaseType_Member_1_MAR")

	return
}

func (enum_codendbusagebasetype_member_1 Enum_CodeNDBUsageBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENR")
	res = append(res, "L")
	res = append(res, "MAR")

	return
}

// Utility function for Enum_CodeNavaidPurposeBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavaidpurposebasetype_member_1 Enum_CodeNavaidPurposeBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavaidpurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL:
		res = "TERMINAL"
	case Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE:
		res = "ENROUTE"
	case Enum_CodeNavaidPurposeBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codenavaidpurposebasetype_member_1 *Enum_CodeNavaidPurposeBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TERMINAL":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL
		return
	case "ENROUTE":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE
		return
	case "ALL":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavaidpurposebasetype_member_1 *Enum_CodeNavaidPurposeBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL
	case "Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE
	case "Enum_CodeNavaidPurposeBaseType_Member_1_ALL":
		*enum_codenavaidpurposebasetype_member_1 = Enum_CodeNavaidPurposeBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavaidpurposebasetype_member_1 *Enum_CodeNavaidPurposeBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavaidpurposebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL:
		res = "Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL"
	case Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE:
		res = "Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE"
	case Enum_CodeNavaidPurposeBaseType_Member_1_ALL:
		res = "Enum_CodeNavaidPurposeBaseType_Member_1_ALL"
	}
	return
}

func (enum_codenavaidpurposebasetype_member_1 Enum_CodeNavaidPurposeBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavaidPurposeBaseType_Member_1_TERMINAL")
	res = append(res, "Enum_CodeNavaidPurposeBaseType_Member_1_ENROUTE")
	res = append(res, "Enum_CodeNavaidPurposeBaseType_Member_1_ALL")

	return
}

func (enum_codenavaidpurposebasetype_member_1 Enum_CodeNavaidPurposeBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TERMINAL")
	res = append(res, "ENROUTE")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeNavaidServiceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenavaidservicebasetype_member_1 Enum_CodeNavaidServiceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenavaidservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavaidServiceBaseType_Member_1_VOR:
		res = "VOR"
	case Enum_CodeNavaidServiceBaseType_Member_1_DME:
		res = "DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB:
		res = "NDB"
	case Enum_CodeNavaidServiceBaseType_Member_1_TACAN:
		res = "TACAN"
	case Enum_CodeNavaidServiceBaseType_Member_1_MKR:
		res = "MKR"
	case Enum_CodeNavaidServiceBaseType_Member_1_ILS:
		res = "ILS"
	case Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME:
		res = "ILS_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_MLS:
		res = "MLS"
	case Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME:
		res = "MLS_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_VORTAC:
		res = "VORTAC"
	case Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME:
		res = "VOR_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME:
		res = "NDB_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_TLS:
		res = "TLS"
	case Enum_CodeNavaidServiceBaseType_Member_1_LOC:
		res = "LOC"
	case Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME:
		res = "LOC_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR:
		res = "NDB_MKR"
	case Enum_CodeNavaidServiceBaseType_Member_1_DF:
		res = "DF"
	case Enum_CodeNavaidServiceBaseType_Member_1_SDF:
		res = "SDF"
	}
	return
}

func (enum_codenavaidservicebasetype_member_1 *Enum_CodeNavaidServiceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "VOR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VOR
		return
	case "DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_DME
		return
	case "NDB":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB
		return
	case "TACAN":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_TACAN
		return
	case "MKR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MKR
		return
	case "ILS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_ILS
		return
	case "ILS_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME
		return
	case "MLS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MLS
		return
	case "MLS_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME
		return
	case "VORTAC":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VORTAC
		return
	case "VOR_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME
		return
	case "NDB_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME
		return
	case "TLS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_TLS
		return
	case "LOC":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_LOC
		return
	case "LOC_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME
		return
	case "NDB_MKR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR
		return
	case "DF":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_DF
		return
	case "SDF":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_SDF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenavaidservicebasetype_member_1 *Enum_CodeNavaidServiceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNavaidServiceBaseType_Member_1_VOR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VOR
	case "Enum_CodeNavaidServiceBaseType_Member_1_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_NDB":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB
	case "Enum_CodeNavaidServiceBaseType_Member_1_TACAN":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_TACAN
	case "Enum_CodeNavaidServiceBaseType_Member_1_MKR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MKR
	case "Enum_CodeNavaidServiceBaseType_Member_1_ILS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_ILS
	case "Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_MLS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MLS
	case "Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_VORTAC":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VORTAC
	case "Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_TLS":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_TLS
	case "Enum_CodeNavaidServiceBaseType_Member_1_LOC":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_LOC
	case "Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME
	case "Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR
	case "Enum_CodeNavaidServiceBaseType_Member_1_DF":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_DF
	case "Enum_CodeNavaidServiceBaseType_Member_1_SDF":
		*enum_codenavaidservicebasetype_member_1 = Enum_CodeNavaidServiceBaseType_Member_1_SDF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenavaidservicebasetype_member_1 *Enum_CodeNavaidServiceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenavaidservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNavaidServiceBaseType_Member_1_VOR:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_VOR"
	case Enum_CodeNavaidServiceBaseType_Member_1_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_NDB"
	case Enum_CodeNavaidServiceBaseType_Member_1_TACAN:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_TACAN"
	case Enum_CodeNavaidServiceBaseType_Member_1_MKR:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_MKR"
	case Enum_CodeNavaidServiceBaseType_Member_1_ILS:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_ILS"
	case Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_MLS:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_MLS"
	case Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_VORTAC:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_VORTAC"
	case Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_TLS:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_TLS"
	case Enum_CodeNavaidServiceBaseType_Member_1_LOC:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_LOC"
	case Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME"
	case Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR"
	case Enum_CodeNavaidServiceBaseType_Member_1_DF:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_DF"
	case Enum_CodeNavaidServiceBaseType_Member_1_SDF:
		res = "Enum_CodeNavaidServiceBaseType_Member_1_SDF"
	}
	return
}

func (enum_codenavaidservicebasetype_member_1 Enum_CodeNavaidServiceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_VOR")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_NDB")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_TACAN")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_MKR")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_ILS")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_ILS_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_MLS")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_MLS_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_VORTAC")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_VOR_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_NDB_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_TLS")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_LOC")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_LOC_DME")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_NDB_MKR")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_DF")
	res = append(res, "Enum_CodeNavaidServiceBaseType_Member_1_SDF")

	return
}

func (enum_codenavaidservicebasetype_member_1 Enum_CodeNavaidServiceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "VOR")
	res = append(res, "DME")
	res = append(res, "NDB")
	res = append(res, "TACAN")
	res = append(res, "MKR")
	res = append(res, "ILS")
	res = append(res, "ILS_DME")
	res = append(res, "MLS")
	res = append(res, "MLS_DME")
	res = append(res, "VORTAC")
	res = append(res, "VOR_DME")
	res = append(res, "NDB_DME")
	res = append(res, "TLS")
	res = append(res, "LOC")
	res = append(res, "LOC_DME")
	res = append(res, "NDB_MKR")
	res = append(res, "DF")
	res = append(res, "SDF")

	return
}

// Utility function for Enum_CodeNorthReferenceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenorthreferencebasetype_member_1 Enum_CodeNorthReferenceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenorthreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD:
		res = "TRUE"
	case Enum_CodeNorthReferenceBaseType_Member_1_MAG:
		res = "MAG"
	case Enum_CodeNorthReferenceBaseType_Member_1_GRID:
		res = "GRID"
	}
	return
}

func (enum_codenorthreferencebasetype_member_1 *Enum_CodeNorthReferenceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TRUE":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD
		return
	case "MAG":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_MAG
		return
	case "GRID":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_GRID
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenorthreferencebasetype_member_1 *Enum_CodeNorthReferenceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD
	case "Enum_CodeNorthReferenceBaseType_Member_1_MAG":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_MAG
	case "Enum_CodeNorthReferenceBaseType_Member_1_GRID":
		*enum_codenorthreferencebasetype_member_1 = Enum_CodeNorthReferenceBaseType_Member_1_GRID
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenorthreferencebasetype_member_1 *Enum_CodeNorthReferenceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenorthreferencebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD:
		res = "Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD"
	case Enum_CodeNorthReferenceBaseType_Member_1_MAG:
		res = "Enum_CodeNorthReferenceBaseType_Member_1_MAG"
	case Enum_CodeNorthReferenceBaseType_Member_1_GRID:
		res = "Enum_CodeNorthReferenceBaseType_Member_1_GRID"
	}
	return
}

func (enum_codenorthreferencebasetype_member_1 Enum_CodeNorthReferenceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNorthReferenceBaseType_Member_1_TRUE_XSD")
	res = append(res, "Enum_CodeNorthReferenceBaseType_Member_1_MAG")
	res = append(res, "Enum_CodeNorthReferenceBaseType_Member_1_GRID")

	return
}

func (enum_codenorthreferencebasetype_member_1 Enum_CodeNorthReferenceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TRUE")
	res = append(res, "MAG")
	res = append(res, "GRID")

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

// Utility function for Enum_CodePARBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeparbasetype_member_1 Enum_CodePARBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeparbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePARBaseType_Member_1_FPN16:
		res = "FPN16"
	case Enum_CodePARBaseType_Member_1_FPN62:
		res = "FPN62"
	case Enum_CodePARBaseType_Member_1_GPN22:
		res = "GPN22"
	case Enum_CodePARBaseType_Member_1_GPN25:
		res = "GPN25"
	case Enum_CodePARBaseType_Member_1_MPN14K:
		res = "MPN14K"
	case Enum_CodePARBaseType_Member_1_TPN19:
		res = "TPN19"
	}
	return
}

func (enum_codeparbasetype_member_1 *Enum_CodePARBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FPN16":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_FPN16
		return
	case "FPN62":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_FPN62
		return
	case "GPN22":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_GPN22
		return
	case "GPN25":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_GPN25
		return
	case "MPN14K":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_MPN14K
		return
	case "TPN19":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_TPN19
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeparbasetype_member_1 *Enum_CodePARBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePARBaseType_Member_1_FPN16":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_FPN16
	case "Enum_CodePARBaseType_Member_1_FPN62":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_FPN62
	case "Enum_CodePARBaseType_Member_1_GPN22":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_GPN22
	case "Enum_CodePARBaseType_Member_1_GPN25":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_GPN25
	case "Enum_CodePARBaseType_Member_1_MPN14K":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_MPN14K
	case "Enum_CodePARBaseType_Member_1_TPN19":
		*enum_codeparbasetype_member_1 = Enum_CodePARBaseType_Member_1_TPN19
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeparbasetype_member_1 *Enum_CodePARBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeparbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePARBaseType_Member_1_FPN16:
		res = "Enum_CodePARBaseType_Member_1_FPN16"
	case Enum_CodePARBaseType_Member_1_FPN62:
		res = "Enum_CodePARBaseType_Member_1_FPN62"
	case Enum_CodePARBaseType_Member_1_GPN22:
		res = "Enum_CodePARBaseType_Member_1_GPN22"
	case Enum_CodePARBaseType_Member_1_GPN25:
		res = "Enum_CodePARBaseType_Member_1_GPN25"
	case Enum_CodePARBaseType_Member_1_MPN14K:
		res = "Enum_CodePARBaseType_Member_1_MPN14K"
	case Enum_CodePARBaseType_Member_1_TPN19:
		res = "Enum_CodePARBaseType_Member_1_TPN19"
	}
	return
}

func (enum_codeparbasetype_member_1 Enum_CodePARBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePARBaseType_Member_1_FPN16")
	res = append(res, "Enum_CodePARBaseType_Member_1_FPN62")
	res = append(res, "Enum_CodePARBaseType_Member_1_GPN22")
	res = append(res, "Enum_CodePARBaseType_Member_1_GPN25")
	res = append(res, "Enum_CodePARBaseType_Member_1_MPN14K")
	res = append(res, "Enum_CodePARBaseType_Member_1_TPN19")

	return
}

func (enum_codeparbasetype_member_1 Enum_CodePARBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FPN16")
	res = append(res, "FPN62")
	res = append(res, "GPN22")
	res = append(res, "GPN25")
	res = append(res, "MPN14K")
	res = append(res, "TPN19")

	return
}

// Utility function for Enum_CodePositionInILSBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepositioninilsbasetype_member_1 Enum_CodePositionInILSBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepositioninilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePositionInILSBaseType_Member_1_OUTER:
		res = "OUTER"
	case Enum_CodePositionInILSBaseType_Member_1_MIDDLE:
		res = "MIDDLE"
	case Enum_CodePositionInILSBaseType_Member_1_INNER:
		res = "INNER"
	case Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE:
		res = "BACKCOURSE"
	}
	return
}

func (enum_codepositioninilsbasetype_member_1 *Enum_CodePositionInILSBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OUTER":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_OUTER
		return
	case "MIDDLE":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_MIDDLE
		return
	case "INNER":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_INNER
		return
	case "BACKCOURSE":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepositioninilsbasetype_member_1 *Enum_CodePositionInILSBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePositionInILSBaseType_Member_1_OUTER":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_OUTER
	case "Enum_CodePositionInILSBaseType_Member_1_MIDDLE":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_MIDDLE
	case "Enum_CodePositionInILSBaseType_Member_1_INNER":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_INNER
	case "Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE":
		*enum_codepositioninilsbasetype_member_1 = Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepositioninilsbasetype_member_1 *Enum_CodePositionInILSBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepositioninilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePositionInILSBaseType_Member_1_OUTER:
		res = "Enum_CodePositionInILSBaseType_Member_1_OUTER"
	case Enum_CodePositionInILSBaseType_Member_1_MIDDLE:
		res = "Enum_CodePositionInILSBaseType_Member_1_MIDDLE"
	case Enum_CodePositionInILSBaseType_Member_1_INNER:
		res = "Enum_CodePositionInILSBaseType_Member_1_INNER"
	case Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE:
		res = "Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE"
	}
	return
}

func (enum_codepositioninilsbasetype_member_1 Enum_CodePositionInILSBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePositionInILSBaseType_Member_1_OUTER")
	res = append(res, "Enum_CodePositionInILSBaseType_Member_1_MIDDLE")
	res = append(res, "Enum_CodePositionInILSBaseType_Member_1_INNER")
	res = append(res, "Enum_CodePositionInILSBaseType_Member_1_BACKCOURSE")

	return
}

func (enum_codepositioninilsbasetype_member_1 Enum_CodePositionInILSBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OUTER")
	res = append(res, "MIDDLE")
	res = append(res, "INNER")
	res = append(res, "BACKCOURSE")

	return
}

// Utility function for Enum_CodePrimaryRadarBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeprimaryradarbasetype_member_1 Enum_CodePrimaryRadarBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeprimaryradarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePrimaryRadarBaseType_Member_1_ASR:
		res = "ASR"
	case Enum_CodePrimaryRadarBaseType_Member_1_ARSR:
		res = "ARSR"
	}
	return
}

func (enum_codeprimaryradarbasetype_member_1 *Enum_CodePrimaryRadarBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ASR":
		*enum_codeprimaryradarbasetype_member_1 = Enum_CodePrimaryRadarBaseType_Member_1_ASR
		return
	case "ARSR":
		*enum_codeprimaryradarbasetype_member_1 = Enum_CodePrimaryRadarBaseType_Member_1_ARSR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeprimaryradarbasetype_member_1 *Enum_CodePrimaryRadarBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePrimaryRadarBaseType_Member_1_ASR":
		*enum_codeprimaryradarbasetype_member_1 = Enum_CodePrimaryRadarBaseType_Member_1_ASR
	case "Enum_CodePrimaryRadarBaseType_Member_1_ARSR":
		*enum_codeprimaryradarbasetype_member_1 = Enum_CodePrimaryRadarBaseType_Member_1_ARSR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeprimaryradarbasetype_member_1 *Enum_CodePrimaryRadarBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeprimaryradarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePrimaryRadarBaseType_Member_1_ASR:
		res = "Enum_CodePrimaryRadarBaseType_Member_1_ASR"
	case Enum_CodePrimaryRadarBaseType_Member_1_ARSR:
		res = "Enum_CodePrimaryRadarBaseType_Member_1_ARSR"
	}
	return
}

func (enum_codeprimaryradarbasetype_member_1 Enum_CodePrimaryRadarBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePrimaryRadarBaseType_Member_1_ASR")
	res = append(res, "Enum_CodePrimaryRadarBaseType_Member_1_ARSR")

	return
}

func (enum_codeprimaryradarbasetype_member_1 Enum_CodePrimaryRadarBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ASR")
	res = append(res, "ARSR")

	return
}

// Utility function for Enum_CodeRadarServiceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderadarservicebasetype_member_1 Enum_CodeRadarServiceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderadarservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadarServiceBaseType_Member_1_PAR:
		res = "PAR"
	case Enum_CodeRadarServiceBaseType_Member_1_ARSR:
		res = "ARSR"
	case Enum_CodeRadarServiceBaseType_Member_1_ASR:
		res = "ASR"
	case Enum_CodeRadarServiceBaseType_Member_1_SSR:
		res = "SSR"
	}
	return
}

func (enum_coderadarservicebasetype_member_1 *Enum_CodeRadarServiceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PAR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_PAR
		return
	case "ARSR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_ARSR
		return
	case "ASR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_ASR
		return
	case "SSR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_SSR
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderadarservicebasetype_member_1 *Enum_CodeRadarServiceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRadarServiceBaseType_Member_1_PAR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_PAR
	case "Enum_CodeRadarServiceBaseType_Member_1_ARSR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_ARSR
	case "Enum_CodeRadarServiceBaseType_Member_1_ASR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_ASR
	case "Enum_CodeRadarServiceBaseType_Member_1_SSR":
		*enum_coderadarservicebasetype_member_1 = Enum_CodeRadarServiceBaseType_Member_1_SSR
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderadarservicebasetype_member_1 *Enum_CodeRadarServiceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderadarservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadarServiceBaseType_Member_1_PAR:
		res = "Enum_CodeRadarServiceBaseType_Member_1_PAR"
	case Enum_CodeRadarServiceBaseType_Member_1_ARSR:
		res = "Enum_CodeRadarServiceBaseType_Member_1_ARSR"
	case Enum_CodeRadarServiceBaseType_Member_1_ASR:
		res = "Enum_CodeRadarServiceBaseType_Member_1_ASR"
	case Enum_CodeRadarServiceBaseType_Member_1_SSR:
		res = "Enum_CodeRadarServiceBaseType_Member_1_SSR"
	}
	return
}

func (enum_coderadarservicebasetype_member_1 Enum_CodeRadarServiceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRadarServiceBaseType_Member_1_PAR")
	res = append(res, "Enum_CodeRadarServiceBaseType_Member_1_ARSR")
	res = append(res, "Enum_CodeRadarServiceBaseType_Member_1_ASR")
	res = append(res, "Enum_CodeRadarServiceBaseType_Member_1_SSR")

	return
}

func (enum_coderadarservicebasetype_member_1 Enum_CodeRadarServiceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PAR")
	res = append(res, "ARSR")
	res = append(res, "ASR")
	res = append(res, "SSR")

	return
}

// Utility function for Enum_CodeRadioEmissionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderadioemissionbasetype_member_1 Enum_CodeRadioEmissionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderadioemissionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadioEmissionBaseType_Member_1_A2:
		res = "A2"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3A:
		res = "A3A"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3B:
		res = "A3B"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3E:
		res = "A3E"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3H:
		res = "A3H"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3J:
		res = "A3J"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3L:
		res = "A3L"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3U:
		res = "A3U"
	case Enum_CodeRadioEmissionBaseType_Member_1_J3E:
		res = "J3E"
	case Enum_CodeRadioEmissionBaseType_Member_1_NONA1A:
		res = "NONA1A"
	case Enum_CodeRadioEmissionBaseType_Member_1_NONA2A:
		res = "NONA2A"
	case Enum_CodeRadioEmissionBaseType_Member_1_PON:
		res = "PON"
	case Enum_CodeRadioEmissionBaseType_Member_1_A8W:
		res = "A8W"
	case Enum_CodeRadioEmissionBaseType_Member_1_A9W:
		res = "A9W"
	case Enum_CodeRadioEmissionBaseType_Member_1_NOX:
		res = "NOX"
	case Enum_CodeRadioEmissionBaseType_Member_1_G1D:
		res = "G1D"
	}
	return
}

func (enum_coderadioemissionbasetype_member_1 *Enum_CodeRadioEmissionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A2":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A2
		return
	case "A3A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3A
		return
	case "A3B":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3B
		return
	case "A3E":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3E
		return
	case "A3H":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3H
		return
	case "A3J":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3J
		return
	case "A3L":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3L
		return
	case "A3U":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3U
		return
	case "J3E":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_J3E
		return
	case "NONA1A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NONA1A
		return
	case "NONA2A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NONA2A
		return
	case "PON":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_PON
		return
	case "A8W":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A8W
		return
	case "A9W":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A9W
		return
	case "NOX":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NOX
		return
	case "G1D":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_G1D
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderadioemissionbasetype_member_1 *Enum_CodeRadioEmissionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRadioEmissionBaseType_Member_1_A2":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A2
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3A
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3B":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3B
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3E":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3E
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3H":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3H
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3J":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3J
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3L":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3L
	case "Enum_CodeRadioEmissionBaseType_Member_1_A3U":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A3U
	case "Enum_CodeRadioEmissionBaseType_Member_1_J3E":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_J3E
	case "Enum_CodeRadioEmissionBaseType_Member_1_NONA1A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NONA1A
	case "Enum_CodeRadioEmissionBaseType_Member_1_NONA2A":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NONA2A
	case "Enum_CodeRadioEmissionBaseType_Member_1_PON":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_PON
	case "Enum_CodeRadioEmissionBaseType_Member_1_A8W":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A8W
	case "Enum_CodeRadioEmissionBaseType_Member_1_A9W":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_A9W
	case "Enum_CodeRadioEmissionBaseType_Member_1_NOX":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_NOX
	case "Enum_CodeRadioEmissionBaseType_Member_1_G1D":
		*enum_coderadioemissionbasetype_member_1 = Enum_CodeRadioEmissionBaseType_Member_1_G1D
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderadioemissionbasetype_member_1 *Enum_CodeRadioEmissionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderadioemissionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadioEmissionBaseType_Member_1_A2:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A2"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3A:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3A"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3B:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3B"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3E:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3E"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3H:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3H"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3J:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3J"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3L:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3L"
	case Enum_CodeRadioEmissionBaseType_Member_1_A3U:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A3U"
	case Enum_CodeRadioEmissionBaseType_Member_1_J3E:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_J3E"
	case Enum_CodeRadioEmissionBaseType_Member_1_NONA1A:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_NONA1A"
	case Enum_CodeRadioEmissionBaseType_Member_1_NONA2A:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_NONA2A"
	case Enum_CodeRadioEmissionBaseType_Member_1_PON:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_PON"
	case Enum_CodeRadioEmissionBaseType_Member_1_A8W:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A8W"
	case Enum_CodeRadioEmissionBaseType_Member_1_A9W:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_A9W"
	case Enum_CodeRadioEmissionBaseType_Member_1_NOX:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_NOX"
	case Enum_CodeRadioEmissionBaseType_Member_1_G1D:
		res = "Enum_CodeRadioEmissionBaseType_Member_1_G1D"
	}
	return
}

func (enum_coderadioemissionbasetype_member_1 Enum_CodeRadioEmissionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A2")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3A")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3B")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3E")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3H")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3J")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3L")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A3U")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_J3E")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_NONA1A")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_NONA2A")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_PON")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A8W")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_A9W")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_NOX")
	res = append(res, "Enum_CodeRadioEmissionBaseType_Member_1_G1D")

	return
}

func (enum_coderadioemissionbasetype_member_1 Enum_CodeRadioEmissionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A2")
	res = append(res, "A3A")
	res = append(res, "A3B")
	res = append(res, "A3E")
	res = append(res, "A3H")
	res = append(res, "A3J")
	res = append(res, "A3L")
	res = append(res, "A3U")
	res = append(res, "J3E")
	res = append(res, "NONA1A")
	res = append(res, "NONA2A")
	res = append(res, "PON")
	res = append(res, "A8W")
	res = append(res, "A9W")
	res = append(res, "NOX")
	res = append(res, "G1D")

	return
}

// Utility function for Enum_CodeRadioSignalBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_coderadiosignalbasetype_member_1 Enum_CodeRadioSignalBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_coderadiosignalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH:
		res = "AZIMUTH"
	case Enum_CodeRadioSignalBaseType_Member_1_DISTANCE:
		res = "DISTANCE"
	case Enum_CodeRadioSignalBaseType_Member_1_BEAM:
		res = "BEAM"
	case Enum_CodeRadioSignalBaseType_Member_1_VOICE:
		res = "VOICE"
	case Enum_CodeRadioSignalBaseType_Member_1_DATALINK:
		res = "DATALINK"
	}
	return
}

func (enum_coderadiosignalbasetype_member_1 *Enum_CodeRadioSignalBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AZIMUTH":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH
		return
	case "DISTANCE":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_DISTANCE
		return
	case "BEAM":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_BEAM
		return
	case "VOICE":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_VOICE
		return
	case "DATALINK":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_DATALINK
		return
	default:
		return errUnkownEnum
	}
}

func (enum_coderadiosignalbasetype_member_1 *Enum_CodeRadioSignalBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH
	case "Enum_CodeRadioSignalBaseType_Member_1_DISTANCE":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_DISTANCE
	case "Enum_CodeRadioSignalBaseType_Member_1_BEAM":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_BEAM
	case "Enum_CodeRadioSignalBaseType_Member_1_VOICE":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_VOICE
	case "Enum_CodeRadioSignalBaseType_Member_1_DATALINK":
		*enum_coderadiosignalbasetype_member_1 = Enum_CodeRadioSignalBaseType_Member_1_DATALINK
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_coderadiosignalbasetype_member_1 *Enum_CodeRadioSignalBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_coderadiosignalbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH:
		res = "Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH"
	case Enum_CodeRadioSignalBaseType_Member_1_DISTANCE:
		res = "Enum_CodeRadioSignalBaseType_Member_1_DISTANCE"
	case Enum_CodeRadioSignalBaseType_Member_1_BEAM:
		res = "Enum_CodeRadioSignalBaseType_Member_1_BEAM"
	case Enum_CodeRadioSignalBaseType_Member_1_VOICE:
		res = "Enum_CodeRadioSignalBaseType_Member_1_VOICE"
	case Enum_CodeRadioSignalBaseType_Member_1_DATALINK:
		res = "Enum_CodeRadioSignalBaseType_Member_1_DATALINK"
	}
	return
}

func (enum_coderadiosignalbasetype_member_1 Enum_CodeRadioSignalBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeRadioSignalBaseType_Member_1_AZIMUTH")
	res = append(res, "Enum_CodeRadioSignalBaseType_Member_1_DISTANCE")
	res = append(res, "Enum_CodeRadioSignalBaseType_Member_1_BEAM")
	res = append(res, "Enum_CodeRadioSignalBaseType_Member_1_VOICE")
	res = append(res, "Enum_CodeRadioSignalBaseType_Member_1_DATALINK")

	return
}

func (enum_coderadiosignalbasetype_member_1 Enum_CodeRadioSignalBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AZIMUTH")
	res = append(res, "DISTANCE")
	res = append(res, "BEAM")
	res = append(res, "VOICE")
	res = append(res, "DATALINK")

	return
}

// Utility function for Enum_CodeReflectorBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codereflectorbasetype_member_1 Enum_CodeReflectorBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codereflectorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN:
		res = "TOUCHDOWN"
	case Enum_CodeReflectorBaseType_Member_1_RUNWAY_END:
		res = "RUNWAY_END"
	case Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD:
		res = "REFERENCE"
	}
	return
}

func (enum_codereflectorbasetype_member_1 *Enum_CodeReflectorBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TOUCHDOWN":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN
		return
	case "RUNWAY_END":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_RUNWAY_END
		return
	case "REFERENCE":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codereflectorbasetype_member_1 *Enum_CodeReflectorBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN
	case "Enum_CodeReflectorBaseType_Member_1_RUNWAY_END":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_RUNWAY_END
	case "Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD":
		*enum_codereflectorbasetype_member_1 = Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codereflectorbasetype_member_1 *Enum_CodeReflectorBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codereflectorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN:
		res = "Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN"
	case Enum_CodeReflectorBaseType_Member_1_RUNWAY_END:
		res = "Enum_CodeReflectorBaseType_Member_1_RUNWAY_END"
	case Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD:
		res = "Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD"
	}
	return
}

func (enum_codereflectorbasetype_member_1 Enum_CodeReflectorBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeReflectorBaseType_Member_1_TOUCHDOWN")
	res = append(res, "Enum_CodeReflectorBaseType_Member_1_RUNWAY_END")
	res = append(res, "Enum_CodeReflectorBaseType_Member_1_REFERENCE_XSD")

	return
}

func (enum_codereflectorbasetype_member_1 Enum_CodeReflectorBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TOUCHDOWN")
	res = append(res, "RUNWAY_END")
	res = append(res, "REFERENCE")

	return
}

// Utility function for Enum_CodeSignalPerformanceILSBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesignalperformanceilsbasetype_member_1 Enum_CodeSignalPerformanceILSBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesignalperformanceilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_I:
		res = "I"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_II:
		res = "II"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_III:
		res = "III"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA:
		res = "IIIA"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB:
		res = "IIIB"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC:
		res = "IIIC"
	}
	return
}

func (enum_codesignalperformanceilsbasetype_member_1 *Enum_CodeSignalPerformanceILSBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "I":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_I
		return
	case "II":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_II
		return
	case "III":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_III
		return
	case "IIIA":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA
		return
	case "IIIB":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB
		return
	case "IIIC":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesignalperformanceilsbasetype_member_1 *Enum_CodeSignalPerformanceILSBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_I":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_I
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_II":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_II
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_III":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_III
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB
	case "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC":
		*enum_codesignalperformanceilsbasetype_member_1 = Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesignalperformanceilsbasetype_member_1 *Enum_CodeSignalPerformanceILSBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesignalperformanceilsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_I:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_I"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_II:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_II"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_III:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_III"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB"
	case Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC:
		res = "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC"
	}
	return
}

func (enum_codesignalperformanceilsbasetype_member_1 Enum_CodeSignalPerformanceILSBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_I")
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_II")
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_III")
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIA")
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIB")
	res = append(res, "Enum_CodeSignalPerformanceILSBaseType_Member_1_IIIC")

	return
}

func (enum_codesignalperformanceilsbasetype_member_1 Enum_CodeSignalPerformanceILSBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "I")
	res = append(res, "II")
	res = append(res, "III")
	res = append(res, "IIIA")
	res = append(res, "IIIB")
	res = append(res, "IIIC")

	return
}

// Utility function for Enum_CodeSpecialNavigationStationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codespecialnavigationstationbasetype_member_1 Enum_CodeSpecialNavigationStationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codespecialnavigationstationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER:
		res = "MASTER"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE:
		res = "SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE:
		res = "RED_SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE:
		res = "GREEN_SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE:
		res = "PURPLE_SLAVE"
	}
	return
}

func (enum_codespecialnavigationstationbasetype_member_1 *Enum_CodeSpecialNavigationStationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MASTER":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER
		return
	case "SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE
		return
	case "RED_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE
		return
	case "GREEN_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE
		return
	case "PURPLE_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codespecialnavigationstationbasetype_member_1 *Enum_CodeSpecialNavigationStationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER
	case "Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE
	case "Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE
	case "Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE
	case "Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE":
		*enum_codespecialnavigationstationbasetype_member_1 = Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codespecialnavigationstationbasetype_member_1 *Enum_CodeSpecialNavigationStationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codespecialnavigationstationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER:
		res = "Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE:
		res = "Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE:
		res = "Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE:
		res = "Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE"
	case Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE:
		res = "Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE"
	}
	return
}

func (enum_codespecialnavigationstationbasetype_member_1 Enum_CodeSpecialNavigationStationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSpecialNavigationStationBaseType_Member_1_MASTER")
	res = append(res, "Enum_CodeSpecialNavigationStationBaseType_Member_1_SLAVE")
	res = append(res, "Enum_CodeSpecialNavigationStationBaseType_Member_1_RED_SLAVE")
	res = append(res, "Enum_CodeSpecialNavigationStationBaseType_Member_1_GREEN_SLAVE")
	res = append(res, "Enum_CodeSpecialNavigationStationBaseType_Member_1_PURPLE_SLAVE")

	return
}

func (enum_codespecialnavigationstationbasetype_member_1 Enum_CodeSpecialNavigationStationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MASTER")
	res = append(res, "SLAVE")
	res = append(res, "RED_SLAVE")
	res = append(res, "GREEN_SLAVE")
	res = append(res, "PURPLE_SLAVE")

	return
}

// Utility function for Enum_CodeSpecialNavigationSystemBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codespecialnavigationsystembasetype_member_1 Enum_CodeSpecialNavigationSystemBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codespecialnavigationsystembasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA:
		res = "LORANA"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC:
		res = "LORANC"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND:
		res = "LORAND"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA:
		res = "DECCA"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS:
		res = "GNSS"
	}
	return
}

func (enum_codespecialnavigationsystembasetype_member_1 *Enum_CodeSpecialNavigationSystemBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LORANA":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA
		return
	case "LORANC":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC
		return
	case "LORAND":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND
		return
	case "DECCA":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA
		return
	case "GNSS":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codespecialnavigationsystembasetype_member_1 *Enum_CodeSpecialNavigationSystemBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA
	case "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC
	case "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND
	case "Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA
	case "Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS":
		*enum_codespecialnavigationsystembasetype_member_1 = Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codespecialnavigationsystembasetype_member_1 *Enum_CodeSpecialNavigationSystemBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codespecialnavigationsystembasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA:
		res = "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC:
		res = "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND:
		res = "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA:
		res = "Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA"
	case Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS:
		res = "Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS"
	}
	return
}

func (enum_codespecialnavigationsystembasetype_member_1 Enum_CodeSpecialNavigationSystemBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANA")
	res = append(res, "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORANC")
	res = append(res, "Enum_CodeSpecialNavigationSystemBaseType_Member_1_LORAND")
	res = append(res, "Enum_CodeSpecialNavigationSystemBaseType_Member_1_DECCA")
	res = append(res, "Enum_CodeSpecialNavigationSystemBaseType_Member_1_GNSS")

	return
}

func (enum_codespecialnavigationsystembasetype_member_1 Enum_CodeSpecialNavigationSystemBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LORANA")
	res = append(res, "LORANC")
	res = append(res, "LORAND")
	res = append(res, "DECCA")
	res = append(res, "GNSS")

	return
}

// Utility function for Enum_CodeStandbyPowerBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestandbypowerbasetype_member_1 Enum_CodeStandbyPowerBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestandbypowerbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStandbyPowerBaseType_Member_1_BATTERY:
		res = "BATTERY"
	case Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL:
		res = "COMMERCIAL"
	case Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR:
		res = "GENERATOR"
	case Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN:
		res = "UNKNOWN"
	case Enum_CodeStandbyPowerBaseType_Member_1_NONE:
		res = "NONE"
	}
	return
}

func (enum_codestandbypowerbasetype_member_1 *Enum_CodeStandbyPowerBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BATTERY":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_BATTERY
		return
	case "COMMERCIAL":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL
		return
	case "GENERATOR":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR
		return
	case "UNKNOWN":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN
		return
	case "NONE":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_NONE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestandbypowerbasetype_member_1 *Enum_CodeStandbyPowerBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStandbyPowerBaseType_Member_1_BATTERY":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_BATTERY
	case "Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL
	case "Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR
	case "Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN
	case "Enum_CodeStandbyPowerBaseType_Member_1_NONE":
		*enum_codestandbypowerbasetype_member_1 = Enum_CodeStandbyPowerBaseType_Member_1_NONE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestandbypowerbasetype_member_1 *Enum_CodeStandbyPowerBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestandbypowerbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStandbyPowerBaseType_Member_1_BATTERY:
		res = "Enum_CodeStandbyPowerBaseType_Member_1_BATTERY"
	case Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL:
		res = "Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL"
	case Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR:
		res = "Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR"
	case Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN:
		res = "Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN"
	case Enum_CodeStandbyPowerBaseType_Member_1_NONE:
		res = "Enum_CodeStandbyPowerBaseType_Member_1_NONE"
	}
	return
}

func (enum_codestandbypowerbasetype_member_1 Enum_CodeStandbyPowerBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStandbyPowerBaseType_Member_1_BATTERY")
	res = append(res, "Enum_CodeStandbyPowerBaseType_Member_1_COMMERCIAL")
	res = append(res, "Enum_CodeStandbyPowerBaseType_Member_1_GENERATOR")
	res = append(res, "Enum_CodeStandbyPowerBaseType_Member_1_UNKNOWN")
	res = append(res, "Enum_CodeStandbyPowerBaseType_Member_1_NONE")

	return
}

func (enum_codestandbypowerbasetype_member_1 Enum_CodeStandbyPowerBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BATTERY")
	res = append(res, "COMMERCIAL")
	res = append(res, "GENERATOR")
	res = append(res, "UNKNOWN")
	res = append(res, "NONE")

	return
}

// Utility function for Enum_CodeStatusNavaidBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusnavaidbasetype_member_1 Enum_CodeStatusNavaidBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusnavaidbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL:
		res = "OPERATIONAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE:
		res = "UNSERVICEABLE"
	case Enum_CodeStatusNavaidBaseType_Member_1_ONTEST:
		res = "ONTEST"
	case Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT:
		res = "INTERRUPT"
	case Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL:
		res = "PARTIAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL:
		res = "CONDITIONAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION:
		res = "FALSE_INDICATION"
	case Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE:
		res = "FALSE_POSSIBLE"
	case Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED:
		res = "DISPLACED"
	case Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION:
		res = "IN_CONSTRUCTION"
	}
	return
}

func (enum_codestatusnavaidbasetype_member_1 *Enum_CodeStatusNavaidBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OPERATIONAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL
		return
	case "UNSERVICEABLE":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE
		return
	case "ONTEST":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_ONTEST
		return
	case "INTERRUPT":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT
		return
	case "PARTIAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL
		return
	case "CONDITIONAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL
		return
	case "FALSE_INDICATION":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION
		return
	case "FALSE_POSSIBLE":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE
		return
	case "DISPLACED":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED
		return
	case "IN_CONSTRUCTION":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusnavaidbasetype_member_1 *Enum_CodeStatusNavaidBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL
	case "Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE
	case "Enum_CodeStatusNavaidBaseType_Member_1_ONTEST":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_ONTEST
	case "Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT
	case "Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL
	case "Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL
	case "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION
	case "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE
	case "Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED
	case "Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION":
		*enum_codestatusnavaidbasetype_member_1 = Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusnavaidbasetype_member_1 *Enum_CodeStatusNavaidBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusnavaidbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE"
	case Enum_CodeStatusNavaidBaseType_Member_1_ONTEST:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_ONTEST"
	case Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT"
	case Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL"
	case Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION"
	case Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE"
	case Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED"
	case Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION:
		res = "Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION"
	}
	return
}

func (enum_codestatusnavaidbasetype_member_1 Enum_CodeStatusNavaidBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_OPERATIONAL")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_UNSERVICEABLE")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_ONTEST")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_INTERRUPT")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_PARTIAL")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_CONDITIONAL")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_INDICATION")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_FALSE_POSSIBLE")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_DISPLACED")
	res = append(res, "Enum_CodeStatusNavaidBaseType_Member_1_IN_CONSTRUCTION")

	return
}

func (enum_codestatusnavaidbasetype_member_1 Enum_CodeStatusNavaidBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OPERATIONAL")
	res = append(res, "UNSERVICEABLE")
	res = append(res, "ONTEST")
	res = append(res, "INTERRUPT")
	res = append(res, "PARTIAL")
	res = append(res, "CONDITIONAL")
	res = append(res, "FALSE_INDICATION")
	res = append(res, "FALSE_POSSIBLE")
	res = append(res, "DISPLACED")
	res = append(res, "IN_CONSTRUCTION")

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

// Utility function for Enum_CodeVORBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codevorbasetype_member_1 Enum_CodeVORBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codevorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVORBaseType_Member_1_VOR:
		res = "VOR"
	case Enum_CodeVORBaseType_Member_1_DVOR:
		res = "DVOR"
	case Enum_CodeVORBaseType_Member_1_VOT:
		res = "VOT"
	}
	return
}

func (enum_codevorbasetype_member_1 *Enum_CodeVORBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "VOR":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_VOR
		return
	case "DVOR":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_DVOR
		return
	case "VOT":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_VOT
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codevorbasetype_member_1 *Enum_CodeVORBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeVORBaseType_Member_1_VOR":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_VOR
	case "Enum_CodeVORBaseType_Member_1_DVOR":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_DVOR
	case "Enum_CodeVORBaseType_Member_1_VOT":
		*enum_codevorbasetype_member_1 = Enum_CodeVORBaseType_Member_1_VOT
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codevorbasetype_member_1 *Enum_CodeVORBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codevorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeVORBaseType_Member_1_VOR:
		res = "Enum_CodeVORBaseType_Member_1_VOR"
	case Enum_CodeVORBaseType_Member_1_DVOR:
		res = "Enum_CodeVORBaseType_Member_1_DVOR"
	case Enum_CodeVORBaseType_Member_1_VOT:
		res = "Enum_CodeVORBaseType_Member_1_VOT"
	}
	return
}

func (enum_codevorbasetype_member_1 Enum_CodeVORBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeVORBaseType_Member_1_VOR")
	res = append(res, "Enum_CodeVORBaseType_Member_1_DVOR")
	res = append(res, "Enum_CodeVORBaseType_Member_1_VOT")

	return
}

func (enum_codevorbasetype_member_1 Enum_CodeVORBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "VOR")
	res = append(res, "DVOR")
	res = append(res, "VOT")

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

// Utility function for Enum_UomFrequencyType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_uomfrequencytype_member_1 Enum_UomFrequencyType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_uomfrequencytype_member_1 {
	// insertion code per enum code
	case Enum_UomFrequencyType_Member_1_HZ:
		res = "HZ"
	case Enum_UomFrequencyType_Member_1_KHZ:
		res = "KHZ"
	case Enum_UomFrequencyType_Member_1_MHZ:
		res = "MHZ"
	case Enum_UomFrequencyType_Member_1_GHZ:
		res = "GHZ"
	}
	return
}

func (enum_uomfrequencytype_member_1 *Enum_UomFrequencyType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_HZ
		return
	case "KHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_KHZ
		return
	case "MHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_MHZ
		return
	case "GHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_GHZ
		return
	default:
		return errUnkownEnum
	}
}

func (enum_uomfrequencytype_member_1 *Enum_UomFrequencyType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_UomFrequencyType_Member_1_HZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_HZ
	case "Enum_UomFrequencyType_Member_1_KHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_KHZ
	case "Enum_UomFrequencyType_Member_1_MHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_MHZ
	case "Enum_UomFrequencyType_Member_1_GHZ":
		*enum_uomfrequencytype_member_1 = Enum_UomFrequencyType_Member_1_GHZ
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_uomfrequencytype_member_1 *Enum_UomFrequencyType_Member_1) ToCodeString() (res string) {

	switch *enum_uomfrequencytype_member_1 {
	// insertion code per enum code
	case Enum_UomFrequencyType_Member_1_HZ:
		res = "Enum_UomFrequencyType_Member_1_HZ"
	case Enum_UomFrequencyType_Member_1_KHZ:
		res = "Enum_UomFrequencyType_Member_1_KHZ"
	case Enum_UomFrequencyType_Member_1_MHZ:
		res = "Enum_UomFrequencyType_Member_1_MHZ"
	case Enum_UomFrequencyType_Member_1_GHZ:
		res = "Enum_UomFrequencyType_Member_1_GHZ"
	}
	return
}

func (enum_uomfrequencytype_member_1 Enum_UomFrequencyType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_UomFrequencyType_Member_1_HZ")
	res = append(res, "Enum_UomFrequencyType_Member_1_KHZ")
	res = append(res, "Enum_UomFrequencyType_Member_1_MHZ")
	res = append(res, "Enum_UomFrequencyType_Member_1_GHZ")

	return
}

func (enum_uomfrequencytype_member_1 Enum_UomFrequencyType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HZ")
	res = append(res, "KHZ")
	res = append(res, "MHZ")
	res = append(res, "GHZ")

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
