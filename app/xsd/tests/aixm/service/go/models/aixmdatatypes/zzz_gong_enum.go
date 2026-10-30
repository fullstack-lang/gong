// generated code - do not edit
package aixmdatatypes

// insertion point of enum utility functions
// Utility function for Enum_CodeAircraftGroundServiceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaircraftgroundservicebasetype_member_1 Enum_CodeAircraftGroundServiceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaircraftgroundservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE:
		res = "DEICE"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND:
		res = "HAND"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR:
		res = "HANGAR"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR:
		res = "REPAIR"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE:
		res = "REMOVE"
	}
	return
}

func (enum_codeaircraftgroundservicebasetype_member_1 *Enum_CodeAircraftGroundServiceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DEICE":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE
		return
	case "HAND":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND
		return
	case "HANGAR":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR
		return
	case "REPAIR":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR
		return
	case "REMOVE":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaircraftgroundservicebasetype_member_1 *Enum_CodeAircraftGroundServiceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE
	case "Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND
	case "Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR
	case "Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR
	case "Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE":
		*enum_codeaircraftgroundservicebasetype_member_1 = Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaircraftgroundservicebasetype_member_1 *Enum_CodeAircraftGroundServiceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaircraftgroundservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE:
		res = "Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND:
		res = "Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR:
		res = "Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR:
		res = "Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR"
	case Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE:
		res = "Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE"
	}
	return
}

func (enum_codeaircraftgroundservicebasetype_member_1 Enum_CodeAircraftGroundServiceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAircraftGroundServiceBaseType_Member_1_DEICE")
	res = append(res, "Enum_CodeAircraftGroundServiceBaseType_Member_1_HAND")
	res = append(res, "Enum_CodeAircraftGroundServiceBaseType_Member_1_HANGAR")
	res = append(res, "Enum_CodeAircraftGroundServiceBaseType_Member_1_REPAIR")
	res = append(res, "Enum_CodeAircraftGroundServiceBaseType_Member_1_REMOVE")

	return
}

func (enum_codeaircraftgroundservicebasetype_member_1 Enum_CodeAircraftGroundServiceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DEICE")
	res = append(res, "HAND")
	res = append(res, "HANGAR")
	res = append(res, "REPAIR")
	res = append(res, "REMOVE")

	return
}

// Utility function for Enum_CodeAviationStandardsBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeaviationstandardsbasetype_member_1 Enum_CodeAviationStandardsBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeaviationstandardsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAviationStandardsBaseType_Member_1_ICAO:
		res = "ICAO"
	case Enum_CodeAviationStandardsBaseType_Member_1_IATA:
		res = "IATA"
	case Enum_CodeAviationStandardsBaseType_Member_1_NATO:
		res = "NATO"
	case Enum_CodeAviationStandardsBaseType_Member_1_FAA:
		res = "FAA"
	}
	return
}

func (enum_codeaviationstandardsbasetype_member_1 *Enum_CodeAviationStandardsBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ICAO":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_ICAO
		return
	case "IATA":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_IATA
		return
	case "NATO":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_NATO
		return
	case "FAA":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_FAA
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeaviationstandardsbasetype_member_1 *Enum_CodeAviationStandardsBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeAviationStandardsBaseType_Member_1_ICAO":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_ICAO
	case "Enum_CodeAviationStandardsBaseType_Member_1_IATA":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_IATA
	case "Enum_CodeAviationStandardsBaseType_Member_1_NATO":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_NATO
	case "Enum_CodeAviationStandardsBaseType_Member_1_FAA":
		*enum_codeaviationstandardsbasetype_member_1 = Enum_CodeAviationStandardsBaseType_Member_1_FAA
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeaviationstandardsbasetype_member_1 *Enum_CodeAviationStandardsBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeaviationstandardsbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeAviationStandardsBaseType_Member_1_ICAO:
		res = "Enum_CodeAviationStandardsBaseType_Member_1_ICAO"
	case Enum_CodeAviationStandardsBaseType_Member_1_IATA:
		res = "Enum_CodeAviationStandardsBaseType_Member_1_IATA"
	case Enum_CodeAviationStandardsBaseType_Member_1_NATO:
		res = "Enum_CodeAviationStandardsBaseType_Member_1_NATO"
	case Enum_CodeAviationStandardsBaseType_Member_1_FAA:
		res = "Enum_CodeAviationStandardsBaseType_Member_1_FAA"
	}
	return
}

func (enum_codeaviationstandardsbasetype_member_1 Enum_CodeAviationStandardsBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeAviationStandardsBaseType_Member_1_ICAO")
	res = append(res, "Enum_CodeAviationStandardsBaseType_Member_1_IATA")
	res = append(res, "Enum_CodeAviationStandardsBaseType_Member_1_NATO")
	res = append(res, "Enum_CodeAviationStandardsBaseType_Member_1_FAA")

	return
}

func (enum_codeaviationstandardsbasetype_member_1 Enum_CodeAviationStandardsBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ICAO")
	res = append(res, "IATA")
	res = append(res, "NATO")
	res = append(res, "FAA")

	return
}

// Utility function for Enum_CodeCommunicationDirectionBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codecommunicationdirectionbasetype_member_1 Enum_CodeCommunicationDirectionBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codecommunicationdirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK:
		res = "UPLINK"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK:
		res = "DOWNLINK"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL:
		res = "BIDIRECTIONAL"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST:
		res = "UPCAST"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST:
		res = "DOWNCAST"
	}
	return
}

func (enum_codecommunicationdirectionbasetype_member_1 *Enum_CodeCommunicationDirectionBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "UPLINK":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK
		return
	case "DOWNLINK":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK
		return
	case "BIDIRECTIONAL":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL
		return
	case "UPCAST":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST
		return
	case "DOWNCAST":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codecommunicationdirectionbasetype_member_1 *Enum_CodeCommunicationDirectionBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK
	case "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK
	case "Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL
	case "Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST
	case "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST":
		*enum_codecommunicationdirectionbasetype_member_1 = Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codecommunicationdirectionbasetype_member_1 *Enum_CodeCommunicationDirectionBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codecommunicationdirectionbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK:
		res = "Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK:
		res = "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL:
		res = "Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST:
		res = "Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST"
	case Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST:
		res = "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST"
	}
	return
}

func (enum_codecommunicationdirectionbasetype_member_1 Enum_CodeCommunicationDirectionBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeCommunicationDirectionBaseType_Member_1_UPLINK")
	res = append(res, "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNLINK")
	res = append(res, "Enum_CodeCommunicationDirectionBaseType_Member_1_BIDIRECTIONAL")
	res = append(res, "Enum_CodeCommunicationDirectionBaseType_Member_1_UPCAST")
	res = append(res, "Enum_CodeCommunicationDirectionBaseType_Member_1_DOWNCAST")

	return
}

func (enum_codecommunicationdirectionbasetype_member_1 Enum_CodeCommunicationDirectionBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "UPLINK")
	res = append(res, "DOWNLINK")
	res = append(res, "BIDIRECTIONAL")
	res = append(res, "UPCAST")
	res = append(res, "DOWNCAST")

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

// Utility function for Enum_CodeFacilityRankingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefacilityrankingbasetype_member_1 Enum_CodeFacilityRankingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefacilityrankingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY:
		res = "PRIMARY"
	case Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY:
		res = "SECONDARY"
	case Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE:
		res = "ALTERNATE"
	case Enum_CodeFacilityRankingBaseType_Member_1_EMERG:
		res = "EMERG"
	case Enum_CodeFacilityRankingBaseType_Member_1_GUARD:
		res = "GUARD"
	}
	return
}

func (enum_codefacilityrankingbasetype_member_1 *Enum_CodeFacilityRankingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PRIMARY":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY
		return
	case "SECONDARY":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY
		return
	case "ALTERNATE":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE
		return
	case "EMERG":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_EMERG
		return
	case "GUARD":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_GUARD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefacilityrankingbasetype_member_1 *Enum_CodeFacilityRankingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY
	case "Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY
	case "Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE
	case "Enum_CodeFacilityRankingBaseType_Member_1_EMERG":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_EMERG
	case "Enum_CodeFacilityRankingBaseType_Member_1_GUARD":
		*enum_codefacilityrankingbasetype_member_1 = Enum_CodeFacilityRankingBaseType_Member_1_GUARD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefacilityrankingbasetype_member_1 *Enum_CodeFacilityRankingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefacilityrankingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY:
		res = "Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY"
	case Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY:
		res = "Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY"
	case Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE:
		res = "Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE"
	case Enum_CodeFacilityRankingBaseType_Member_1_EMERG:
		res = "Enum_CodeFacilityRankingBaseType_Member_1_EMERG"
	case Enum_CodeFacilityRankingBaseType_Member_1_GUARD:
		res = "Enum_CodeFacilityRankingBaseType_Member_1_GUARD"
	}
	return
}

func (enum_codefacilityrankingbasetype_member_1 Enum_CodeFacilityRankingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFacilityRankingBaseType_Member_1_PRIMARY")
	res = append(res, "Enum_CodeFacilityRankingBaseType_Member_1_SECONDARY")
	res = append(res, "Enum_CodeFacilityRankingBaseType_Member_1_ALTERNATE")
	res = append(res, "Enum_CodeFacilityRankingBaseType_Member_1_EMERG")
	res = append(res, "Enum_CodeFacilityRankingBaseType_Member_1_GUARD")

	return
}

func (enum_codefacilityrankingbasetype_member_1 Enum_CodeFacilityRankingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PRIMARY")
	res = append(res, "SECONDARY")
	res = append(res, "ALTERNATE")
	res = append(res, "EMERG")
	res = append(res, "GUARD")

	return
}

// Utility function for Enum_CodeFireFightingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefirefightingbasetype_member_1 Enum_CodeFireFightingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefirefightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFireFightingBaseType_Member_1_H1:
		res = "H1"
	case Enum_CodeFireFightingBaseType_Member_1_H2:
		res = "H2"
	case Enum_CodeFireFightingBaseType_Member_1_H3:
		res = "H3"
	case Enum_CodeFireFightingBaseType_Member_1_A1:
		res = "A1"
	case Enum_CodeFireFightingBaseType_Member_1_A2:
		res = "A2"
	case Enum_CodeFireFightingBaseType_Member_1_A3:
		res = "A3"
	case Enum_CodeFireFightingBaseType_Member_1_A4:
		res = "A4"
	case Enum_CodeFireFightingBaseType_Member_1_A5:
		res = "A5"
	case Enum_CodeFireFightingBaseType_Member_1_A6:
		res = "A6"
	case Enum_CodeFireFightingBaseType_Member_1_A7:
		res = "A7"
	case Enum_CodeFireFightingBaseType_Member_1_A8:
		res = "A8"
	case Enum_CodeFireFightingBaseType_Member_1_A9:
		res = "A9"
	case Enum_CodeFireFightingBaseType_Member_1_A10:
		res = "A10"
	}
	return
}

func (enum_codefirefightingbasetype_member_1 *Enum_CodeFireFightingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "H1":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H1
		return
	case "H2":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H2
		return
	case "H3":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H3
		return
	case "A1":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A1
		return
	case "A2":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A2
		return
	case "A3":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A3
		return
	case "A4":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A4
		return
	case "A5":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A5
		return
	case "A6":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A6
		return
	case "A7":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A7
		return
	case "A8":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A8
		return
	case "A9":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A9
		return
	case "A10":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A10
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefirefightingbasetype_member_1 *Enum_CodeFireFightingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFireFightingBaseType_Member_1_H1":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H1
	case "Enum_CodeFireFightingBaseType_Member_1_H2":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H2
	case "Enum_CodeFireFightingBaseType_Member_1_H3":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_H3
	case "Enum_CodeFireFightingBaseType_Member_1_A1":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A1
	case "Enum_CodeFireFightingBaseType_Member_1_A2":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A2
	case "Enum_CodeFireFightingBaseType_Member_1_A3":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A3
	case "Enum_CodeFireFightingBaseType_Member_1_A4":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A4
	case "Enum_CodeFireFightingBaseType_Member_1_A5":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A5
	case "Enum_CodeFireFightingBaseType_Member_1_A6":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A6
	case "Enum_CodeFireFightingBaseType_Member_1_A7":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A7
	case "Enum_CodeFireFightingBaseType_Member_1_A8":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A8
	case "Enum_CodeFireFightingBaseType_Member_1_A9":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A9
	case "Enum_CodeFireFightingBaseType_Member_1_A10":
		*enum_codefirefightingbasetype_member_1 = Enum_CodeFireFightingBaseType_Member_1_A10
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefirefightingbasetype_member_1 *Enum_CodeFireFightingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefirefightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFireFightingBaseType_Member_1_H1:
		res = "Enum_CodeFireFightingBaseType_Member_1_H1"
	case Enum_CodeFireFightingBaseType_Member_1_H2:
		res = "Enum_CodeFireFightingBaseType_Member_1_H2"
	case Enum_CodeFireFightingBaseType_Member_1_H3:
		res = "Enum_CodeFireFightingBaseType_Member_1_H3"
	case Enum_CodeFireFightingBaseType_Member_1_A1:
		res = "Enum_CodeFireFightingBaseType_Member_1_A1"
	case Enum_CodeFireFightingBaseType_Member_1_A2:
		res = "Enum_CodeFireFightingBaseType_Member_1_A2"
	case Enum_CodeFireFightingBaseType_Member_1_A3:
		res = "Enum_CodeFireFightingBaseType_Member_1_A3"
	case Enum_CodeFireFightingBaseType_Member_1_A4:
		res = "Enum_CodeFireFightingBaseType_Member_1_A4"
	case Enum_CodeFireFightingBaseType_Member_1_A5:
		res = "Enum_CodeFireFightingBaseType_Member_1_A5"
	case Enum_CodeFireFightingBaseType_Member_1_A6:
		res = "Enum_CodeFireFightingBaseType_Member_1_A6"
	case Enum_CodeFireFightingBaseType_Member_1_A7:
		res = "Enum_CodeFireFightingBaseType_Member_1_A7"
	case Enum_CodeFireFightingBaseType_Member_1_A8:
		res = "Enum_CodeFireFightingBaseType_Member_1_A8"
	case Enum_CodeFireFightingBaseType_Member_1_A9:
		res = "Enum_CodeFireFightingBaseType_Member_1_A9"
	case Enum_CodeFireFightingBaseType_Member_1_A10:
		res = "Enum_CodeFireFightingBaseType_Member_1_A10"
	}
	return
}

func (enum_codefirefightingbasetype_member_1 Enum_CodeFireFightingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_H1")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_H2")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_H3")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A1")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A2")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A3")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A4")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A5")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A6")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A7")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A8")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A9")
	res = append(res, "Enum_CodeFireFightingBaseType_Member_1_A10")

	return
}

func (enum_codefirefightingbasetype_member_1 Enum_CodeFireFightingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "H1")
	res = append(res, "H2")
	res = append(res, "H3")
	res = append(res, "A1")
	res = append(res, "A2")
	res = append(res, "A3")
	res = append(res, "A4")
	res = append(res, "A5")
	res = append(res, "A6")
	res = append(res, "A7")
	res = append(res, "A8")
	res = append(res, "A9")
	res = append(res, "A10")

	return
}

// Utility function for Enum_CodeFlightDestinationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeflightdestinationbasetype_member_1 Enum_CodeFlightDestinationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeflightdestinationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightDestinationBaseType_Member_1_ARR:
		res = "ARR"
	case Enum_CodeFlightDestinationBaseType_Member_1_DEP:
		res = "DEP"
	case Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY:
		res = "OVERFLY"
	case Enum_CodeFlightDestinationBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codeflightdestinationbasetype_member_1 *Enum_CodeFlightDestinationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ARR":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_ARR
		return
	case "DEP":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_DEP
		return
	case "OVERFLY":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY
		return
	case "ALL":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeflightdestinationbasetype_member_1 *Enum_CodeFlightDestinationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFlightDestinationBaseType_Member_1_ARR":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_ARR
	case "Enum_CodeFlightDestinationBaseType_Member_1_DEP":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_DEP
	case "Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY
	case "Enum_CodeFlightDestinationBaseType_Member_1_ALL":
		*enum_codeflightdestinationbasetype_member_1 = Enum_CodeFlightDestinationBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeflightdestinationbasetype_member_1 *Enum_CodeFlightDestinationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeflightdestinationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFlightDestinationBaseType_Member_1_ARR:
		res = "Enum_CodeFlightDestinationBaseType_Member_1_ARR"
	case Enum_CodeFlightDestinationBaseType_Member_1_DEP:
		res = "Enum_CodeFlightDestinationBaseType_Member_1_DEP"
	case Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY:
		res = "Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY"
	case Enum_CodeFlightDestinationBaseType_Member_1_ALL:
		res = "Enum_CodeFlightDestinationBaseType_Member_1_ALL"
	}
	return
}

func (enum_codeflightdestinationbasetype_member_1 Enum_CodeFlightDestinationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFlightDestinationBaseType_Member_1_ARR")
	res = append(res, "Enum_CodeFlightDestinationBaseType_Member_1_DEP")
	res = append(res, "Enum_CodeFlightDestinationBaseType_Member_1_OVERFLY")
	res = append(res, "Enum_CodeFlightDestinationBaseType_Member_1_ALL")

	return
}

func (enum_codeflightdestinationbasetype_member_1 Enum_CodeFlightDestinationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ARR")
	res = append(res, "DEP")
	res = append(res, "OVERFLY")
	res = append(res, "ALL")

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

// Utility function for Enum_CodeFuelBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codefuelbasetype_member_1 Enum_CodeFuelBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codefuelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFuelBaseType_Member_1_AVGAS:
		res = "AVGAS"
	case Enum_CodeFuelBaseType_Member_1_AVGAS_LL:
		res = "AVGAS_LL"
	case Enum_CodeFuelBaseType_Member_1_OCT73:
		res = "OCT73"
	case Enum_CodeFuelBaseType_Member_1_OCT80:
		res = "OCT80"
	case Enum_CodeFuelBaseType_Member_1_OCT82UL:
		res = "OCT82UL"
	case Enum_CodeFuelBaseType_Member_1_OCT80_87:
		res = "OCT80_87"
	case Enum_CodeFuelBaseType_Member_1_OCT91_98:
		res = "OCT91_98"
	case Enum_CodeFuelBaseType_Member_1_OCT100_130:
		res = "OCT100_130"
	case Enum_CodeFuelBaseType_Member_1_OCT108_135:
		res = "OCT108_135"
	case Enum_CodeFuelBaseType_Member_1_OCT115_145:
		res = "OCT115_145"
	case Enum_CodeFuelBaseType_Member_1_MOGAS:
		res = "MOGAS"
	case Enum_CodeFuelBaseType_Member_1_JET:
		res = "JET"
	case Enum_CodeFuelBaseType_Member_1_A:
		res = "A"
	case Enum_CodeFuelBaseType_Member_1_A1:
		res = "A1"
	case Enum_CodeFuelBaseType_Member_1_A1_PLUS:
		res = "A1_PLUS"
	case Enum_CodeFuelBaseType_Member_1_B:
		res = "B"
	case Enum_CodeFuelBaseType_Member_1_JP1:
		res = "JP1"
	case Enum_CodeFuelBaseType_Member_1_JP2:
		res = "JP2"
	case Enum_CodeFuelBaseType_Member_1_JP3:
		res = "JP3"
	case Enum_CodeFuelBaseType_Member_1_JP4:
		res = "JP4"
	case Enum_CodeFuelBaseType_Member_1_JP5:
		res = "JP5"
	case Enum_CodeFuelBaseType_Member_1_JP6:
		res = "JP6"
	case Enum_CodeFuelBaseType_Member_1_JPTS:
		res = "JPTS"
	case Enum_CodeFuelBaseType_Member_1_JP7:
		res = "JP7"
	case Enum_CodeFuelBaseType_Member_1_JP8:
		res = "JP8"
	case Enum_CodeFuelBaseType_Member_1_JP8_HIGHER:
		res = "JP8_HIGHER"
	case Enum_CodeFuelBaseType_Member_1_JP9:
		res = "JP9"
	case Enum_CodeFuelBaseType_Member_1_JP10:
		res = "JP10"
	case Enum_CodeFuelBaseType_Member_1_F18:
		res = "F18"
	case Enum_CodeFuelBaseType_Member_1_F34:
		res = "F34"
	case Enum_CodeFuelBaseType_Member_1_F35:
		res = "F35"
	case Enum_CodeFuelBaseType_Member_1_F40:
		res = "F40"
	case Enum_CodeFuelBaseType_Member_1_F44:
		res = "F44"
	case Enum_CodeFuelBaseType_Member_1_TR0:
		res = "TR0"
	case Enum_CodeFuelBaseType_Member_1_TR4:
		res = "TR4"
	case Enum_CodeFuelBaseType_Member_1_TS1:
		res = "TS1"
	case Enum_CodeFuelBaseType_Member_1_RT:
		res = "RT"
	case Enum_CodeFuelBaseType_Member_1_ALL:
		res = "ALL"
	}
	return
}

func (enum_codefuelbasetype_member_1 *Enum_CodeFuelBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AVGAS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_AVGAS
		return
	case "AVGAS_LL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_AVGAS_LL
		return
	case "OCT73":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT73
		return
	case "OCT80":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT80
		return
	case "OCT82UL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT82UL
		return
	case "OCT80_87":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT80_87
		return
	case "OCT91_98":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT91_98
		return
	case "OCT100_130":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT100_130
		return
	case "OCT108_135":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT108_135
		return
	case "OCT115_145":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT115_145
		return
	case "MOGAS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_MOGAS
		return
	case "JET":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JET
		return
	case "A":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A
		return
	case "A1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A1
		return
	case "A1_PLUS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A1_PLUS
		return
	case "B":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_B
		return
	case "JP1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP1
		return
	case "JP2":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP2
		return
	case "JP3":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP3
		return
	case "JP4":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP4
		return
	case "JP5":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP5
		return
	case "JP6":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP6
		return
	case "JPTS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JPTS
		return
	case "JP7":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP7
		return
	case "JP8":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP8
		return
	case "JP8_HIGHER":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP8_HIGHER
		return
	case "JP9":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP9
		return
	case "JP10":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP10
		return
	case "F18":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F18
		return
	case "F34":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F34
		return
	case "F35":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F35
		return
	case "F40":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F40
		return
	case "F44":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F44
		return
	case "TR0":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TR0
		return
	case "TR4":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TR4
		return
	case "TS1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TS1
		return
	case "RT":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_RT
		return
	case "ALL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_ALL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codefuelbasetype_member_1 *Enum_CodeFuelBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeFuelBaseType_Member_1_AVGAS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_AVGAS
	case "Enum_CodeFuelBaseType_Member_1_AVGAS_LL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_AVGAS_LL
	case "Enum_CodeFuelBaseType_Member_1_OCT73":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT73
	case "Enum_CodeFuelBaseType_Member_1_OCT80":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT80
	case "Enum_CodeFuelBaseType_Member_1_OCT82UL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT82UL
	case "Enum_CodeFuelBaseType_Member_1_OCT80_87":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT80_87
	case "Enum_CodeFuelBaseType_Member_1_OCT91_98":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT91_98
	case "Enum_CodeFuelBaseType_Member_1_OCT100_130":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT100_130
	case "Enum_CodeFuelBaseType_Member_1_OCT108_135":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT108_135
	case "Enum_CodeFuelBaseType_Member_1_OCT115_145":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_OCT115_145
	case "Enum_CodeFuelBaseType_Member_1_MOGAS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_MOGAS
	case "Enum_CodeFuelBaseType_Member_1_JET":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JET
	case "Enum_CodeFuelBaseType_Member_1_A":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A
	case "Enum_CodeFuelBaseType_Member_1_A1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A1
	case "Enum_CodeFuelBaseType_Member_1_A1_PLUS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_A1_PLUS
	case "Enum_CodeFuelBaseType_Member_1_B":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_B
	case "Enum_CodeFuelBaseType_Member_1_JP1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP1
	case "Enum_CodeFuelBaseType_Member_1_JP2":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP2
	case "Enum_CodeFuelBaseType_Member_1_JP3":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP3
	case "Enum_CodeFuelBaseType_Member_1_JP4":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP4
	case "Enum_CodeFuelBaseType_Member_1_JP5":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP5
	case "Enum_CodeFuelBaseType_Member_1_JP6":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP6
	case "Enum_CodeFuelBaseType_Member_1_JPTS":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JPTS
	case "Enum_CodeFuelBaseType_Member_1_JP7":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP7
	case "Enum_CodeFuelBaseType_Member_1_JP8":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP8
	case "Enum_CodeFuelBaseType_Member_1_JP8_HIGHER":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP8_HIGHER
	case "Enum_CodeFuelBaseType_Member_1_JP9":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP9
	case "Enum_CodeFuelBaseType_Member_1_JP10":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_JP10
	case "Enum_CodeFuelBaseType_Member_1_F18":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F18
	case "Enum_CodeFuelBaseType_Member_1_F34":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F34
	case "Enum_CodeFuelBaseType_Member_1_F35":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F35
	case "Enum_CodeFuelBaseType_Member_1_F40":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F40
	case "Enum_CodeFuelBaseType_Member_1_F44":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_F44
	case "Enum_CodeFuelBaseType_Member_1_TR0":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TR0
	case "Enum_CodeFuelBaseType_Member_1_TR4":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TR4
	case "Enum_CodeFuelBaseType_Member_1_TS1":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_TS1
	case "Enum_CodeFuelBaseType_Member_1_RT":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_RT
	case "Enum_CodeFuelBaseType_Member_1_ALL":
		*enum_codefuelbasetype_member_1 = Enum_CodeFuelBaseType_Member_1_ALL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codefuelbasetype_member_1 *Enum_CodeFuelBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codefuelbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeFuelBaseType_Member_1_AVGAS:
		res = "Enum_CodeFuelBaseType_Member_1_AVGAS"
	case Enum_CodeFuelBaseType_Member_1_AVGAS_LL:
		res = "Enum_CodeFuelBaseType_Member_1_AVGAS_LL"
	case Enum_CodeFuelBaseType_Member_1_OCT73:
		res = "Enum_CodeFuelBaseType_Member_1_OCT73"
	case Enum_CodeFuelBaseType_Member_1_OCT80:
		res = "Enum_CodeFuelBaseType_Member_1_OCT80"
	case Enum_CodeFuelBaseType_Member_1_OCT82UL:
		res = "Enum_CodeFuelBaseType_Member_1_OCT82UL"
	case Enum_CodeFuelBaseType_Member_1_OCT80_87:
		res = "Enum_CodeFuelBaseType_Member_1_OCT80_87"
	case Enum_CodeFuelBaseType_Member_1_OCT91_98:
		res = "Enum_CodeFuelBaseType_Member_1_OCT91_98"
	case Enum_CodeFuelBaseType_Member_1_OCT100_130:
		res = "Enum_CodeFuelBaseType_Member_1_OCT100_130"
	case Enum_CodeFuelBaseType_Member_1_OCT108_135:
		res = "Enum_CodeFuelBaseType_Member_1_OCT108_135"
	case Enum_CodeFuelBaseType_Member_1_OCT115_145:
		res = "Enum_CodeFuelBaseType_Member_1_OCT115_145"
	case Enum_CodeFuelBaseType_Member_1_MOGAS:
		res = "Enum_CodeFuelBaseType_Member_1_MOGAS"
	case Enum_CodeFuelBaseType_Member_1_JET:
		res = "Enum_CodeFuelBaseType_Member_1_JET"
	case Enum_CodeFuelBaseType_Member_1_A:
		res = "Enum_CodeFuelBaseType_Member_1_A"
	case Enum_CodeFuelBaseType_Member_1_A1:
		res = "Enum_CodeFuelBaseType_Member_1_A1"
	case Enum_CodeFuelBaseType_Member_1_A1_PLUS:
		res = "Enum_CodeFuelBaseType_Member_1_A1_PLUS"
	case Enum_CodeFuelBaseType_Member_1_B:
		res = "Enum_CodeFuelBaseType_Member_1_B"
	case Enum_CodeFuelBaseType_Member_1_JP1:
		res = "Enum_CodeFuelBaseType_Member_1_JP1"
	case Enum_CodeFuelBaseType_Member_1_JP2:
		res = "Enum_CodeFuelBaseType_Member_1_JP2"
	case Enum_CodeFuelBaseType_Member_1_JP3:
		res = "Enum_CodeFuelBaseType_Member_1_JP3"
	case Enum_CodeFuelBaseType_Member_1_JP4:
		res = "Enum_CodeFuelBaseType_Member_1_JP4"
	case Enum_CodeFuelBaseType_Member_1_JP5:
		res = "Enum_CodeFuelBaseType_Member_1_JP5"
	case Enum_CodeFuelBaseType_Member_1_JP6:
		res = "Enum_CodeFuelBaseType_Member_1_JP6"
	case Enum_CodeFuelBaseType_Member_1_JPTS:
		res = "Enum_CodeFuelBaseType_Member_1_JPTS"
	case Enum_CodeFuelBaseType_Member_1_JP7:
		res = "Enum_CodeFuelBaseType_Member_1_JP7"
	case Enum_CodeFuelBaseType_Member_1_JP8:
		res = "Enum_CodeFuelBaseType_Member_1_JP8"
	case Enum_CodeFuelBaseType_Member_1_JP8_HIGHER:
		res = "Enum_CodeFuelBaseType_Member_1_JP8_HIGHER"
	case Enum_CodeFuelBaseType_Member_1_JP9:
		res = "Enum_CodeFuelBaseType_Member_1_JP9"
	case Enum_CodeFuelBaseType_Member_1_JP10:
		res = "Enum_CodeFuelBaseType_Member_1_JP10"
	case Enum_CodeFuelBaseType_Member_1_F18:
		res = "Enum_CodeFuelBaseType_Member_1_F18"
	case Enum_CodeFuelBaseType_Member_1_F34:
		res = "Enum_CodeFuelBaseType_Member_1_F34"
	case Enum_CodeFuelBaseType_Member_1_F35:
		res = "Enum_CodeFuelBaseType_Member_1_F35"
	case Enum_CodeFuelBaseType_Member_1_F40:
		res = "Enum_CodeFuelBaseType_Member_1_F40"
	case Enum_CodeFuelBaseType_Member_1_F44:
		res = "Enum_CodeFuelBaseType_Member_1_F44"
	case Enum_CodeFuelBaseType_Member_1_TR0:
		res = "Enum_CodeFuelBaseType_Member_1_TR0"
	case Enum_CodeFuelBaseType_Member_1_TR4:
		res = "Enum_CodeFuelBaseType_Member_1_TR4"
	case Enum_CodeFuelBaseType_Member_1_TS1:
		res = "Enum_CodeFuelBaseType_Member_1_TS1"
	case Enum_CodeFuelBaseType_Member_1_RT:
		res = "Enum_CodeFuelBaseType_Member_1_RT"
	case Enum_CodeFuelBaseType_Member_1_ALL:
		res = "Enum_CodeFuelBaseType_Member_1_ALL"
	}
	return
}

func (enum_codefuelbasetype_member_1 Enum_CodeFuelBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeFuelBaseType_Member_1_AVGAS")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_AVGAS_LL")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT73")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT80")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT82UL")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT80_87")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT91_98")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT100_130")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT108_135")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_OCT115_145")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_MOGAS")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JET")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_A")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_A1")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_A1_PLUS")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_B")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP1")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP2")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP3")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP4")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP5")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP6")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JPTS")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP7")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP8")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP8_HIGHER")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP9")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_JP10")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_F18")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_F34")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_F35")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_F40")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_F44")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_TR0")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_TR4")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_TS1")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_RT")
	res = append(res, "Enum_CodeFuelBaseType_Member_1_ALL")

	return
}

func (enum_codefuelbasetype_member_1 Enum_CodeFuelBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AVGAS")
	res = append(res, "AVGAS_LL")
	res = append(res, "OCT73")
	res = append(res, "OCT80")
	res = append(res, "OCT82UL")
	res = append(res, "OCT80_87")
	res = append(res, "OCT91_98")
	res = append(res, "OCT100_130")
	res = append(res, "OCT108_135")
	res = append(res, "OCT115_145")
	res = append(res, "MOGAS")
	res = append(res, "JET")
	res = append(res, "A")
	res = append(res, "A1")
	res = append(res, "A1_PLUS")
	res = append(res, "B")
	res = append(res, "JP1")
	res = append(res, "JP2")
	res = append(res, "JP3")
	res = append(res, "JP4")
	res = append(res, "JP5")
	res = append(res, "JP6")
	res = append(res, "JPTS")
	res = append(res, "JP7")
	res = append(res, "JP8")
	res = append(res, "JP8_HIGHER")
	res = append(res, "JP9")
	res = append(res, "JP10")
	res = append(res, "F18")
	res = append(res, "F34")
	res = append(res, "F35")
	res = append(res, "F40")
	res = append(res, "F44")
	res = append(res, "TR0")
	res = append(res, "TR4")
	res = append(res, "TS1")
	res = append(res, "RT")
	res = append(res, "ALL")

	return
}

// Utility function for Enum_CodeIntensityStandByBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeintensitystandbybasetype_member_1 Enum_CodeIntensityStandByBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeintensitystandbybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeIntensityStandByBaseType_Member_1_OFF:
		res = "OFF"
	case Enum_CodeIntensityStandByBaseType_Member_1_LOW:
		res = "LOW"
	}
	return
}

func (enum_codeintensitystandbybasetype_member_1 *Enum_CodeIntensityStandByBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OFF":
		*enum_codeintensitystandbybasetype_member_1 = Enum_CodeIntensityStandByBaseType_Member_1_OFF
		return
	case "LOW":
		*enum_codeintensitystandbybasetype_member_1 = Enum_CodeIntensityStandByBaseType_Member_1_LOW
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeintensitystandbybasetype_member_1 *Enum_CodeIntensityStandByBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeIntensityStandByBaseType_Member_1_OFF":
		*enum_codeintensitystandbybasetype_member_1 = Enum_CodeIntensityStandByBaseType_Member_1_OFF
	case "Enum_CodeIntensityStandByBaseType_Member_1_LOW":
		*enum_codeintensitystandbybasetype_member_1 = Enum_CodeIntensityStandByBaseType_Member_1_LOW
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeintensitystandbybasetype_member_1 *Enum_CodeIntensityStandByBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeintensitystandbybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeIntensityStandByBaseType_Member_1_OFF:
		res = "Enum_CodeIntensityStandByBaseType_Member_1_OFF"
	case Enum_CodeIntensityStandByBaseType_Member_1_LOW:
		res = "Enum_CodeIntensityStandByBaseType_Member_1_LOW"
	}
	return
}

func (enum_codeintensitystandbybasetype_member_1 Enum_CodeIntensityStandByBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeIntensityStandByBaseType_Member_1_OFF")
	res = append(res, "Enum_CodeIntensityStandByBaseType_Member_1_LOW")

	return
}

func (enum_codeintensitystandbybasetype_member_1 Enum_CodeIntensityStandByBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OFF")
	res = append(res, "LOW")

	return
}

// Utility function for Enum_CodeLevelSeriesBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codelevelseriesbasetype_member_1 Enum_CodeLevelSeriesBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codelevelseriesbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelSeriesBaseType_Member_1_ODD:
		res = "ODD"
	case Enum_CodeLevelSeriesBaseType_Member_1_EVEN:
		res = "EVEN"
	case Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF:
		res = "ODD_HALF"
	case Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF:
		res = "EVEN_HALF"
	}
	return
}

func (enum_codelevelseriesbasetype_member_1 *Enum_CodeLevelSeriesBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ODD":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_ODD
		return
	case "EVEN":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_EVEN
		return
	case "ODD_HALF":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF
		return
	case "EVEN_HALF":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codelevelseriesbasetype_member_1 *Enum_CodeLevelSeriesBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLevelSeriesBaseType_Member_1_ODD":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_ODD
	case "Enum_CodeLevelSeriesBaseType_Member_1_EVEN":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_EVEN
	case "Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF
	case "Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF":
		*enum_codelevelseriesbasetype_member_1 = Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codelevelseriesbasetype_member_1 *Enum_CodeLevelSeriesBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codelevelseriesbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelSeriesBaseType_Member_1_ODD:
		res = "Enum_CodeLevelSeriesBaseType_Member_1_ODD"
	case Enum_CodeLevelSeriesBaseType_Member_1_EVEN:
		res = "Enum_CodeLevelSeriesBaseType_Member_1_EVEN"
	case Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF:
		res = "Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF"
	case Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF:
		res = "Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF"
	}
	return
}

func (enum_codelevelseriesbasetype_member_1 Enum_CodeLevelSeriesBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLevelSeriesBaseType_Member_1_ODD")
	res = append(res, "Enum_CodeLevelSeriesBaseType_Member_1_EVEN")
	res = append(res, "Enum_CodeLevelSeriesBaseType_Member_1_ODD_HALF")
	res = append(res, "Enum_CodeLevelSeriesBaseType_Member_1_EVEN_HALF")

	return
}

func (enum_codelevelseriesbasetype_member_1 Enum_CodeLevelSeriesBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ODD")
	res = append(res, "EVEN")
	res = append(res, "ODD_HALF")
	res = append(res, "EVEN_HALF")

	return
}

// Utility function for Enum_CodeLevelTableDesignatorBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeleveltabledesignatorbasetype_member_1 Enum_CodeLevelTableDesignatorBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeleveltabledesignatorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR:
		res = "IFR"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES:
		res = "IFR_METRES"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR:
		res = "VFR"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES:
		res = "VFR_METRES"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM:
		res = "IFR_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM:
		res = "IFR_METRES_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM:
		res = "VFR_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM:
		res = "VFR_METRES_RVSM"
	}
	return
}

func (enum_codeleveltabledesignatorbasetype_member_1 *Enum_CodeLevelTableDesignatorBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "IFR":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR
		return
	case "IFR_METRES":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES
		return
	case "VFR":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR
		return
	case "VFR_METRES":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES
		return
	case "IFR_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM
		return
	case "IFR_METRES_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM
		return
	case "VFR_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM
		return
	case "VFR_METRES_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeleveltabledesignatorbasetype_member_1 *Enum_CodeLevelTableDesignatorBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM
	case "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM":
		*enum_codeleveltabledesignatorbasetype_member_1 = Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeleveltabledesignatorbasetype_member_1 *Enum_CodeLevelTableDesignatorBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeleveltabledesignatorbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM"
	case Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM:
		res = "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM"
	}
	return
}

func (enum_codeleveltabledesignatorbasetype_member_1 Enum_CodeLevelTableDesignatorBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_RVSM")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_IFR_METRES_RVSM")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_RVSM")
	res = append(res, "Enum_CodeLevelTableDesignatorBaseType_Member_1_VFR_METRES_RVSM")

	return
}

func (enum_codeleveltabledesignatorbasetype_member_1 Enum_CodeLevelTableDesignatorBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "IFR")
	res = append(res, "IFR_METRES")
	res = append(res, "VFR")
	res = append(res, "VFR_METRES")
	res = append(res, "IFR_RVSM")
	res = append(res, "IFR_METRES_RVSM")
	res = append(res, "VFR_RVSM")
	res = append(res, "VFR_METRES_RVSM")

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

// Utility function for Enum_CodeNitrogenBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codenitrogenbasetype_member_1 Enum_CodeNitrogenBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codenitrogenbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNitrogenBaseType_Member_1_LPNG:
		res = "LPNG"
	case Enum_CodeNitrogenBaseType_Member_1_HPNG:
		res = "HPNG"
	case Enum_CodeNitrogenBaseType_Member_1_LHNG:
		res = "LHNG"
	case Enum_CodeNitrogenBaseType_Member_1_LNG:
		res = "LNG"
	case Enum_CodeNitrogenBaseType_Member_1_NGRB:
		res = "NGRB"
	case Enum_CodeNitrogenBaseType_Member_1_HNGRB:
		res = "HNGRB"
	case Enum_CodeNitrogenBaseType_Member_1_LNGRB:
		res = "LNGRB"
	case Enum_CodeNitrogenBaseType_Member_1_NG:
		res = "NG"
	}
	return
}

func (enum_codenitrogenbasetype_member_1 *Enum_CodeNitrogenBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LPNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LPNG
		return
	case "HPNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_HPNG
		return
	case "LHNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LHNG
		return
	case "LNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LNG
		return
	case "NGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_NGRB
		return
	case "HNGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_HNGRB
		return
	case "LNGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LNGRB
		return
	case "NG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_NG
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codenitrogenbasetype_member_1 *Enum_CodeNitrogenBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeNitrogenBaseType_Member_1_LPNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LPNG
	case "Enum_CodeNitrogenBaseType_Member_1_HPNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_HPNG
	case "Enum_CodeNitrogenBaseType_Member_1_LHNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LHNG
	case "Enum_CodeNitrogenBaseType_Member_1_LNG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LNG
	case "Enum_CodeNitrogenBaseType_Member_1_NGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_NGRB
	case "Enum_CodeNitrogenBaseType_Member_1_HNGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_HNGRB
	case "Enum_CodeNitrogenBaseType_Member_1_LNGRB":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_LNGRB
	case "Enum_CodeNitrogenBaseType_Member_1_NG":
		*enum_codenitrogenbasetype_member_1 = Enum_CodeNitrogenBaseType_Member_1_NG
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codenitrogenbasetype_member_1 *Enum_CodeNitrogenBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codenitrogenbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeNitrogenBaseType_Member_1_LPNG:
		res = "Enum_CodeNitrogenBaseType_Member_1_LPNG"
	case Enum_CodeNitrogenBaseType_Member_1_HPNG:
		res = "Enum_CodeNitrogenBaseType_Member_1_HPNG"
	case Enum_CodeNitrogenBaseType_Member_1_LHNG:
		res = "Enum_CodeNitrogenBaseType_Member_1_LHNG"
	case Enum_CodeNitrogenBaseType_Member_1_LNG:
		res = "Enum_CodeNitrogenBaseType_Member_1_LNG"
	case Enum_CodeNitrogenBaseType_Member_1_NGRB:
		res = "Enum_CodeNitrogenBaseType_Member_1_NGRB"
	case Enum_CodeNitrogenBaseType_Member_1_HNGRB:
		res = "Enum_CodeNitrogenBaseType_Member_1_HNGRB"
	case Enum_CodeNitrogenBaseType_Member_1_LNGRB:
		res = "Enum_CodeNitrogenBaseType_Member_1_LNGRB"
	case Enum_CodeNitrogenBaseType_Member_1_NG:
		res = "Enum_CodeNitrogenBaseType_Member_1_NG"
	}
	return
}

func (enum_codenitrogenbasetype_member_1 Enum_CodeNitrogenBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_LPNG")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_HPNG")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_LHNG")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_LNG")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_NGRB")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_HNGRB")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_LNGRB")
	res = append(res, "Enum_CodeNitrogenBaseType_Member_1_NG")

	return
}

func (enum_codenitrogenbasetype_member_1 Enum_CodeNitrogenBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LPNG")
	res = append(res, "HPNG")
	res = append(res, "LHNG")
	res = append(res, "LNG")
	res = append(res, "NGRB")
	res = append(res, "HNGRB")
	res = append(res, "LNGRB")
	res = append(res, "NG")

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

// Utility function for Enum_CodeOilBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeoilbasetype_member_1 Enum_CodeOilBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeoilbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOilBaseType_Member_1_PISTON:
		res = "PISTON"
	case Enum_CodeOilBaseType_Member_1_TURBO:
		res = "TURBO"
	case Enum_CodeOilBaseType_Member_1_HYDRAULIC:
		res = "HYDRAULIC"
	}
	return
}

func (enum_codeoilbasetype_member_1 *Enum_CodeOilBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PISTON":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_PISTON
		return
	case "TURBO":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_TURBO
		return
	case "HYDRAULIC":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_HYDRAULIC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeoilbasetype_member_1 *Enum_CodeOilBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOilBaseType_Member_1_PISTON":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_PISTON
	case "Enum_CodeOilBaseType_Member_1_TURBO":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_TURBO
	case "Enum_CodeOilBaseType_Member_1_HYDRAULIC":
		*enum_codeoilbasetype_member_1 = Enum_CodeOilBaseType_Member_1_HYDRAULIC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeoilbasetype_member_1 *Enum_CodeOilBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeoilbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOilBaseType_Member_1_PISTON:
		res = "Enum_CodeOilBaseType_Member_1_PISTON"
	case Enum_CodeOilBaseType_Member_1_TURBO:
		res = "Enum_CodeOilBaseType_Member_1_TURBO"
	case Enum_CodeOilBaseType_Member_1_HYDRAULIC:
		res = "Enum_CodeOilBaseType_Member_1_HYDRAULIC"
	}
	return
}

func (enum_codeoilbasetype_member_1 Enum_CodeOilBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOilBaseType_Member_1_PISTON")
	res = append(res, "Enum_CodeOilBaseType_Member_1_TURBO")
	res = append(res, "Enum_CodeOilBaseType_Member_1_HYDRAULIC")

	return
}

func (enum_codeoilbasetype_member_1 Enum_CodeOilBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PISTON")
	res = append(res, "TURBO")
	res = append(res, "HYDRAULIC")

	return
}

// Utility function for Enum_CodeOrganisationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeorganisationbasetype_member_1 Enum_CodeOrganisationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeorganisationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOrganisationBaseType_Member_1_STATE:
		res = "STATE"
	case Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP:
		res = "STATE_GROUP"
	case Enum_CodeOrganisationBaseType_Member_1_ORG:
		res = "ORG"
	case Enum_CodeOrganisationBaseType_Member_1_INTL_ORG:
		res = "INTL_ORG"
	case Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR:
		res = "ACFT_OPR"
	case Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY:
		res = "HANDLING_AGENCY"
	case Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH:
		res = "NTL_AUTH"
	case Enum_CodeOrganisationBaseType_Member_1_ATS:
		res = "ATS"
	case Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL:
		res = "COMMERCIAL"
	}
	return
}

func (enum_codeorganisationbasetype_member_1 *Enum_CodeOrganisationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STATE":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_STATE
		return
	case "STATE_GROUP":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP
		return
	case "ORG":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ORG
		return
	case "INTL_ORG":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_INTL_ORG
		return
	case "ACFT_OPR":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR
		return
	case "HANDLING_AGENCY":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY
		return
	case "NTL_AUTH":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH
		return
	case "ATS":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ATS
		return
	case "COMMERCIAL":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeorganisationbasetype_member_1 *Enum_CodeOrganisationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOrganisationBaseType_Member_1_STATE":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_STATE
	case "Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP
	case "Enum_CodeOrganisationBaseType_Member_1_ORG":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ORG
	case "Enum_CodeOrganisationBaseType_Member_1_INTL_ORG":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_INTL_ORG
	case "Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR
	case "Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY
	case "Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH
	case "Enum_CodeOrganisationBaseType_Member_1_ATS":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_ATS
	case "Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL":
		*enum_codeorganisationbasetype_member_1 = Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeorganisationbasetype_member_1 *Enum_CodeOrganisationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeorganisationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOrganisationBaseType_Member_1_STATE:
		res = "Enum_CodeOrganisationBaseType_Member_1_STATE"
	case Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP:
		res = "Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP"
	case Enum_CodeOrganisationBaseType_Member_1_ORG:
		res = "Enum_CodeOrganisationBaseType_Member_1_ORG"
	case Enum_CodeOrganisationBaseType_Member_1_INTL_ORG:
		res = "Enum_CodeOrganisationBaseType_Member_1_INTL_ORG"
	case Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR:
		res = "Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR"
	case Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY:
		res = "Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY"
	case Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH:
		res = "Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH"
	case Enum_CodeOrganisationBaseType_Member_1_ATS:
		res = "Enum_CodeOrganisationBaseType_Member_1_ATS"
	case Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL:
		res = "Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL"
	}
	return
}

func (enum_codeorganisationbasetype_member_1 Enum_CodeOrganisationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_STATE")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_STATE_GROUP")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_ORG")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_INTL_ORG")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_ACFT_OPR")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_HANDLING_AGENCY")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_NTL_AUTH")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_ATS")
	res = append(res, "Enum_CodeOrganisationBaseType_Member_1_COMMERCIAL")

	return
}

func (enum_codeorganisationbasetype_member_1 Enum_CodeOrganisationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STATE")
	res = append(res, "STATE_GROUP")
	res = append(res, "ORG")
	res = append(res, "INTL_ORG")
	res = append(res, "ACFT_OPR")
	res = append(res, "HANDLING_AGENCY")
	res = append(res, "NTL_AUTH")
	res = append(res, "ATS")
	res = append(res, "COMMERCIAL")

	return
}

// Utility function for Enum_CodeOrganisationHierarchyBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeorganisationhierarchybasetype_member_1 Enum_CodeOrganisationHierarchyBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeorganisationhierarchybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER:
		res = "MEMBER"
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY:
		res = "OWNED_BY"
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY:
		res = "SUPERVISED_BY"
	}
	return
}

func (enum_codeorganisationhierarchybasetype_member_1 *Enum_CodeOrganisationHierarchyBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MEMBER":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER
		return
	case "OWNED_BY":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY
		return
	case "SUPERVISED_BY":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeorganisationhierarchybasetype_member_1 *Enum_CodeOrganisationHierarchyBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER
	case "Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY
	case "Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY":
		*enum_codeorganisationhierarchybasetype_member_1 = Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeorganisationhierarchybasetype_member_1 *Enum_CodeOrganisationHierarchyBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeorganisationhierarchybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER:
		res = "Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER"
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY:
		res = "Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY"
	case Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY:
		res = "Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY"
	}
	return
}

func (enum_codeorganisationhierarchybasetype_member_1 Enum_CodeOrganisationHierarchyBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOrganisationHierarchyBaseType_Member_1_MEMBER")
	res = append(res, "Enum_CodeOrganisationHierarchyBaseType_Member_1_OWNED_BY")
	res = append(res, "Enum_CodeOrganisationHierarchyBaseType_Member_1_SUPERVISED_BY")

	return
}

func (enum_codeorganisationhierarchybasetype_member_1 Enum_CodeOrganisationHierarchyBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MEMBER")
	res = append(res, "OWNED_BY")
	res = append(res, "SUPERVISED_BY")

	return
}

// Utility function for Enum_CodeOxygenBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeoxygenbasetype_member_1 Enum_CodeOxygenBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeoxygenbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOxygenBaseType_Member_1_LPOX:
		res = "LPOX"
	case Enum_CodeOxygenBaseType_Member_1_HPOX:
		res = "HPOX"
	case Enum_CodeOxygenBaseType_Member_1_LHOX:
		res = "LHOX"
	case Enum_CodeOxygenBaseType_Member_1_LOX:
		res = "LOX"
	case Enum_CodeOxygenBaseType_Member_1_OXRB:
		res = "OXRB"
	case Enum_CodeOxygenBaseType_Member_1_HOXRB:
		res = "HOXRB"
	case Enum_CodeOxygenBaseType_Member_1_LOXRB:
		res = "LOXRB"
	case Enum_CodeOxygenBaseType_Member_1_OX:
		res = "OX"
	}
	return
}

func (enum_codeoxygenbasetype_member_1 *Enum_CodeOxygenBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LPOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LPOX
		return
	case "HPOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_HPOX
		return
	case "LHOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LHOX
		return
	case "LOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LOX
		return
	case "OXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_OXRB
		return
	case "HOXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_HOXRB
		return
	case "LOXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LOXRB
		return
	case "OX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_OX
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeoxygenbasetype_member_1 *Enum_CodeOxygenBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeOxygenBaseType_Member_1_LPOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LPOX
	case "Enum_CodeOxygenBaseType_Member_1_HPOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_HPOX
	case "Enum_CodeOxygenBaseType_Member_1_LHOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LHOX
	case "Enum_CodeOxygenBaseType_Member_1_LOX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LOX
	case "Enum_CodeOxygenBaseType_Member_1_OXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_OXRB
	case "Enum_CodeOxygenBaseType_Member_1_HOXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_HOXRB
	case "Enum_CodeOxygenBaseType_Member_1_LOXRB":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_LOXRB
	case "Enum_CodeOxygenBaseType_Member_1_OX":
		*enum_codeoxygenbasetype_member_1 = Enum_CodeOxygenBaseType_Member_1_OX
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeoxygenbasetype_member_1 *Enum_CodeOxygenBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeoxygenbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeOxygenBaseType_Member_1_LPOX:
		res = "Enum_CodeOxygenBaseType_Member_1_LPOX"
	case Enum_CodeOxygenBaseType_Member_1_HPOX:
		res = "Enum_CodeOxygenBaseType_Member_1_HPOX"
	case Enum_CodeOxygenBaseType_Member_1_LHOX:
		res = "Enum_CodeOxygenBaseType_Member_1_LHOX"
	case Enum_CodeOxygenBaseType_Member_1_LOX:
		res = "Enum_CodeOxygenBaseType_Member_1_LOX"
	case Enum_CodeOxygenBaseType_Member_1_OXRB:
		res = "Enum_CodeOxygenBaseType_Member_1_OXRB"
	case Enum_CodeOxygenBaseType_Member_1_HOXRB:
		res = "Enum_CodeOxygenBaseType_Member_1_HOXRB"
	case Enum_CodeOxygenBaseType_Member_1_LOXRB:
		res = "Enum_CodeOxygenBaseType_Member_1_LOXRB"
	case Enum_CodeOxygenBaseType_Member_1_OX:
		res = "Enum_CodeOxygenBaseType_Member_1_OX"
	}
	return
}

func (enum_codeoxygenbasetype_member_1 Enum_CodeOxygenBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_LPOX")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_HPOX")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_LHOX")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_LOX")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_OXRB")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_HOXRB")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_LOXRB")
	res = append(res, "Enum_CodeOxygenBaseType_Member_1_OX")

	return
}

func (enum_codeoxygenbasetype_member_1 Enum_CodeOxygenBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LPOX")
	res = append(res, "HPOX")
	res = append(res, "LHOX")
	res = append(res, "LOX")
	res = append(res, "OXRB")
	res = append(res, "HOXRB")
	res = append(res, "LOXRB")
	res = append(res, "OX")

	return
}

// Utility function for Enum_CodePassengerServiceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepassengerservicebasetype_member_1 Enum_CodePassengerServiceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepassengerservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePassengerServiceBaseType_Member_1_CUST:
		res = "CUST"
	case Enum_CodePassengerServiceBaseType_Member_1_SAN:
		res = "SAN"
	case Enum_CodePassengerServiceBaseType_Member_1_SECUR:
		res = "SECUR"
	case Enum_CodePassengerServiceBaseType_Member_1_VET:
		res = "VET"
	case Enum_CodePassengerServiceBaseType_Member_1_HOTEL:
		res = "HOTEL"
	case Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT:
		res = "TRANSPORT"
	case Enum_CodePassengerServiceBaseType_Member_1_REST:
		res = "REST"
	case Enum_CodePassengerServiceBaseType_Member_1_INFO:
		res = "INFO"
	case Enum_CodePassengerServiceBaseType_Member_1_BANK:
		res = "BANK"
	case Enum_CodePassengerServiceBaseType_Member_1_POST:
		res = "POST"
	case Enum_CodePassengerServiceBaseType_Member_1_MEDIC:
		res = "MEDIC"
	}
	return
}

func (enum_codepassengerservicebasetype_member_1 *Enum_CodePassengerServiceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CUST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_CUST
		return
	case "SAN":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_SAN
		return
	case "SECUR":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_SECUR
		return
	case "VET":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_VET
		return
	case "HOTEL":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_HOTEL
		return
	case "TRANSPORT":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT
		return
	case "REST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_REST
		return
	case "INFO":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_INFO
		return
	case "BANK":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_BANK
		return
	case "POST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_POST
		return
	case "MEDIC":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_MEDIC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepassengerservicebasetype_member_1 *Enum_CodePassengerServiceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePassengerServiceBaseType_Member_1_CUST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_CUST
	case "Enum_CodePassengerServiceBaseType_Member_1_SAN":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_SAN
	case "Enum_CodePassengerServiceBaseType_Member_1_SECUR":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_SECUR
	case "Enum_CodePassengerServiceBaseType_Member_1_VET":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_VET
	case "Enum_CodePassengerServiceBaseType_Member_1_HOTEL":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_HOTEL
	case "Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT
	case "Enum_CodePassengerServiceBaseType_Member_1_REST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_REST
	case "Enum_CodePassengerServiceBaseType_Member_1_INFO":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_INFO
	case "Enum_CodePassengerServiceBaseType_Member_1_BANK":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_BANK
	case "Enum_CodePassengerServiceBaseType_Member_1_POST":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_POST
	case "Enum_CodePassengerServiceBaseType_Member_1_MEDIC":
		*enum_codepassengerservicebasetype_member_1 = Enum_CodePassengerServiceBaseType_Member_1_MEDIC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepassengerservicebasetype_member_1 *Enum_CodePassengerServiceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepassengerservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePassengerServiceBaseType_Member_1_CUST:
		res = "Enum_CodePassengerServiceBaseType_Member_1_CUST"
	case Enum_CodePassengerServiceBaseType_Member_1_SAN:
		res = "Enum_CodePassengerServiceBaseType_Member_1_SAN"
	case Enum_CodePassengerServiceBaseType_Member_1_SECUR:
		res = "Enum_CodePassengerServiceBaseType_Member_1_SECUR"
	case Enum_CodePassengerServiceBaseType_Member_1_VET:
		res = "Enum_CodePassengerServiceBaseType_Member_1_VET"
	case Enum_CodePassengerServiceBaseType_Member_1_HOTEL:
		res = "Enum_CodePassengerServiceBaseType_Member_1_HOTEL"
	case Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT:
		res = "Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT"
	case Enum_CodePassengerServiceBaseType_Member_1_REST:
		res = "Enum_CodePassengerServiceBaseType_Member_1_REST"
	case Enum_CodePassengerServiceBaseType_Member_1_INFO:
		res = "Enum_CodePassengerServiceBaseType_Member_1_INFO"
	case Enum_CodePassengerServiceBaseType_Member_1_BANK:
		res = "Enum_CodePassengerServiceBaseType_Member_1_BANK"
	case Enum_CodePassengerServiceBaseType_Member_1_POST:
		res = "Enum_CodePassengerServiceBaseType_Member_1_POST"
	case Enum_CodePassengerServiceBaseType_Member_1_MEDIC:
		res = "Enum_CodePassengerServiceBaseType_Member_1_MEDIC"
	}
	return
}

func (enum_codepassengerservicebasetype_member_1 Enum_CodePassengerServiceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_CUST")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_SAN")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_SECUR")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_VET")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_HOTEL")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_TRANSPORT")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_REST")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_INFO")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_BANK")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_POST")
	res = append(res, "Enum_CodePassengerServiceBaseType_Member_1_MEDIC")

	return
}

func (enum_codepassengerservicebasetype_member_1 Enum_CodePassengerServiceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CUST")
	res = append(res, "SAN")
	res = append(res, "SECUR")
	res = append(res, "VET")
	res = append(res, "HOTEL")
	res = append(res, "TRANSPORT")
	res = append(res, "REST")
	res = append(res, "INFO")
	res = append(res, "BANK")
	res = append(res, "POST")
	res = append(res, "MEDIC")

	return
}

// Utility function for Enum_CodePilotControlledLightingBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codepilotcontrolledlightingbasetype_member_1 Enum_CodePilotControlledLightingBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codepilotcontrolledlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA:
		res = "STANDARD_FAA"
	case Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD:
		res = "NON_STANDARD"
	}
	return
}

func (enum_codepilotcontrolledlightingbasetype_member_1 *Enum_CodePilotControlledLightingBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "STANDARD_FAA":
		*enum_codepilotcontrolledlightingbasetype_member_1 = Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA
		return
	case "NON_STANDARD":
		*enum_codepilotcontrolledlightingbasetype_member_1 = Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codepilotcontrolledlightingbasetype_member_1 *Enum_CodePilotControlledLightingBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA":
		*enum_codepilotcontrolledlightingbasetype_member_1 = Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA
	case "Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD":
		*enum_codepilotcontrolledlightingbasetype_member_1 = Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codepilotcontrolledlightingbasetype_member_1 *Enum_CodePilotControlledLightingBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codepilotcontrolledlightingbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA:
		res = "Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA"
	case Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD:
		res = "Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD"
	}
	return
}

func (enum_codepilotcontrolledlightingbasetype_member_1 Enum_CodePilotControlledLightingBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodePilotControlledLightingBaseType_Member_1_STANDARD_FAA")
	res = append(res, "Enum_CodePilotControlledLightingBaseType_Member_1_NON_STANDARD")

	return
}

func (enum_codepilotcontrolledlightingbasetype_member_1 Enum_CodePilotControlledLightingBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "STANDARD_FAA")
	res = append(res, "NON_STANDARD")

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

// Utility function for Enum_CodeServiceATCBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeserviceatcbasetype_member_1 Enum_CodeServiceATCBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeserviceatcbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceATCBaseType_Member_1_ACS:
		res = "ACS"
	case Enum_CodeServiceATCBaseType_Member_1_UAC:
		res = "UAC"
	case Enum_CodeServiceATCBaseType_Member_1_OACS:
		res = "OACS"
	case Enum_CodeServiceATCBaseType_Member_1_APP:
		res = "APP"
	case Enum_CodeServiceATCBaseType_Member_1_TWR:
		res = "TWR"
	case Enum_CodeServiceATCBaseType_Member_1_ADVS:
		res = "ADVS"
	case Enum_CodeServiceATCBaseType_Member_1_EFAS:
		res = "EFAS"
	case Enum_CodeServiceATCBaseType_Member_1_CTAF:
		res = "CTAF"
	}
	return
}

func (enum_codeserviceatcbasetype_member_1 *Enum_CodeServiceATCBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ACS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_ACS
		return
	case "UAC":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_UAC
		return
	case "OACS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_OACS
		return
	case "APP":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_APP
		return
	case "TWR":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_TWR
		return
	case "ADVS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_ADVS
		return
	case "EFAS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_EFAS
		return
	case "CTAF":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_CTAF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeserviceatcbasetype_member_1 *Enum_CodeServiceATCBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeServiceATCBaseType_Member_1_ACS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_ACS
	case "Enum_CodeServiceATCBaseType_Member_1_UAC":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_UAC
	case "Enum_CodeServiceATCBaseType_Member_1_OACS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_OACS
	case "Enum_CodeServiceATCBaseType_Member_1_APP":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_APP
	case "Enum_CodeServiceATCBaseType_Member_1_TWR":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_TWR
	case "Enum_CodeServiceATCBaseType_Member_1_ADVS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_ADVS
	case "Enum_CodeServiceATCBaseType_Member_1_EFAS":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_EFAS
	case "Enum_CodeServiceATCBaseType_Member_1_CTAF":
		*enum_codeserviceatcbasetype_member_1 = Enum_CodeServiceATCBaseType_Member_1_CTAF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeserviceatcbasetype_member_1 *Enum_CodeServiceATCBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeserviceatcbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceATCBaseType_Member_1_ACS:
		res = "Enum_CodeServiceATCBaseType_Member_1_ACS"
	case Enum_CodeServiceATCBaseType_Member_1_UAC:
		res = "Enum_CodeServiceATCBaseType_Member_1_UAC"
	case Enum_CodeServiceATCBaseType_Member_1_OACS:
		res = "Enum_CodeServiceATCBaseType_Member_1_OACS"
	case Enum_CodeServiceATCBaseType_Member_1_APP:
		res = "Enum_CodeServiceATCBaseType_Member_1_APP"
	case Enum_CodeServiceATCBaseType_Member_1_TWR:
		res = "Enum_CodeServiceATCBaseType_Member_1_TWR"
	case Enum_CodeServiceATCBaseType_Member_1_ADVS:
		res = "Enum_CodeServiceATCBaseType_Member_1_ADVS"
	case Enum_CodeServiceATCBaseType_Member_1_EFAS:
		res = "Enum_CodeServiceATCBaseType_Member_1_EFAS"
	case Enum_CodeServiceATCBaseType_Member_1_CTAF:
		res = "Enum_CodeServiceATCBaseType_Member_1_CTAF"
	}
	return
}

func (enum_codeserviceatcbasetype_member_1 Enum_CodeServiceATCBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_ACS")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_UAC")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_OACS")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_APP")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_TWR")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_ADVS")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_EFAS")
	res = append(res, "Enum_CodeServiceATCBaseType_Member_1_CTAF")

	return
}

func (enum_codeserviceatcbasetype_member_1 Enum_CodeServiceATCBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ACS")
	res = append(res, "UAC")
	res = append(res, "OACS")
	res = append(res, "APP")
	res = append(res, "TWR")
	res = append(res, "ADVS")
	res = append(res, "EFAS")
	res = append(res, "CTAF")

	return
}

// Utility function for Enum_CodeServiceATFMBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeserviceatfmbasetype_member_1 Enum_CodeServiceATFMBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeserviceatfmbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceATFMBaseType_Member_1_FPL:
		res = "FPL"
	case Enum_CodeServiceATFMBaseType_Member_1_FPLV:
		res = "FPLV"
	case Enum_CodeServiceATFMBaseType_Member_1_ATFM:
		res = "ATFM"
	case Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE:
		res = "CLEARANCE"
	case Enum_CodeServiceATFMBaseType_Member_1_SCHED:
		res = "SCHED"
	}
	return
}

func (enum_codeserviceatfmbasetype_member_1 *Enum_CodeServiceATFMBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FPL":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_FPL
		return
	case "FPLV":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_FPLV
		return
	case "ATFM":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_ATFM
		return
	case "CLEARANCE":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE
		return
	case "SCHED":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_SCHED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeserviceatfmbasetype_member_1 *Enum_CodeServiceATFMBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeServiceATFMBaseType_Member_1_FPL":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_FPL
	case "Enum_CodeServiceATFMBaseType_Member_1_FPLV":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_FPLV
	case "Enum_CodeServiceATFMBaseType_Member_1_ATFM":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_ATFM
	case "Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE
	case "Enum_CodeServiceATFMBaseType_Member_1_SCHED":
		*enum_codeserviceatfmbasetype_member_1 = Enum_CodeServiceATFMBaseType_Member_1_SCHED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeserviceatfmbasetype_member_1 *Enum_CodeServiceATFMBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeserviceatfmbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceATFMBaseType_Member_1_FPL:
		res = "Enum_CodeServiceATFMBaseType_Member_1_FPL"
	case Enum_CodeServiceATFMBaseType_Member_1_FPLV:
		res = "Enum_CodeServiceATFMBaseType_Member_1_FPLV"
	case Enum_CodeServiceATFMBaseType_Member_1_ATFM:
		res = "Enum_CodeServiceATFMBaseType_Member_1_ATFM"
	case Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE:
		res = "Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE"
	case Enum_CodeServiceATFMBaseType_Member_1_SCHED:
		res = "Enum_CodeServiceATFMBaseType_Member_1_SCHED"
	}
	return
}

func (enum_codeserviceatfmbasetype_member_1 Enum_CodeServiceATFMBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeServiceATFMBaseType_Member_1_FPL")
	res = append(res, "Enum_CodeServiceATFMBaseType_Member_1_FPLV")
	res = append(res, "Enum_CodeServiceATFMBaseType_Member_1_ATFM")
	res = append(res, "Enum_CodeServiceATFMBaseType_Member_1_CLEARANCE")
	res = append(res, "Enum_CodeServiceATFMBaseType_Member_1_SCHED")

	return
}

func (enum_codeserviceatfmbasetype_member_1 Enum_CodeServiceATFMBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FPL")
	res = append(res, "FPLV")
	res = append(res, "ATFM")
	res = append(res, "CLEARANCE")
	res = append(res, "SCHED")

	return
}

// Utility function for Enum_CodeServiceGroundControlBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeservicegroundcontrolbasetype_member_1 Enum_CodeServiceGroundControlBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeservicegroundcontrolbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceGroundControlBaseType_Member_1_TWR:
		res = "TWR"
	case Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS:
		res = "SMGCS"
	case Enum_CodeServiceGroundControlBaseType_Member_1_TAXI:
		res = "TAXI"
	}
	return
}

func (enum_codeservicegroundcontrolbasetype_member_1 *Enum_CodeServiceGroundControlBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TWR":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_TWR
		return
	case "SMGCS":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS
		return
	case "TAXI":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_TAXI
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeservicegroundcontrolbasetype_member_1 *Enum_CodeServiceGroundControlBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeServiceGroundControlBaseType_Member_1_TWR":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_TWR
	case "Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS
	case "Enum_CodeServiceGroundControlBaseType_Member_1_TAXI":
		*enum_codeservicegroundcontrolbasetype_member_1 = Enum_CodeServiceGroundControlBaseType_Member_1_TAXI
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeservicegroundcontrolbasetype_member_1 *Enum_CodeServiceGroundControlBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeservicegroundcontrolbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceGroundControlBaseType_Member_1_TWR:
		res = "Enum_CodeServiceGroundControlBaseType_Member_1_TWR"
	case Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS:
		res = "Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS"
	case Enum_CodeServiceGroundControlBaseType_Member_1_TAXI:
		res = "Enum_CodeServiceGroundControlBaseType_Member_1_TAXI"
	}
	return
}

func (enum_codeservicegroundcontrolbasetype_member_1 Enum_CodeServiceGroundControlBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeServiceGroundControlBaseType_Member_1_TWR")
	res = append(res, "Enum_CodeServiceGroundControlBaseType_Member_1_SMGCS")
	res = append(res, "Enum_CodeServiceGroundControlBaseType_Member_1_TAXI")

	return
}

func (enum_codeservicegroundcontrolbasetype_member_1 Enum_CodeServiceGroundControlBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TWR")
	res = append(res, "SMGCS")
	res = append(res, "TAXI")

	return
}

// Utility function for Enum_CodeServiceInformationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeserviceinformationbasetype_member_1 Enum_CodeServiceInformationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeserviceinformationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceInformationBaseType_Member_1_AFIS:
		res = "AFIS"
	case Enum_CodeServiceInformationBaseType_Member_1_AIS:
		res = "AIS"
	case Enum_CodeServiceInformationBaseType_Member_1_ATIS:
		res = "ATIS"
	case Enum_CodeServiceInformationBaseType_Member_1_BRIEFING:
		res = "BRIEFING"
	case Enum_CodeServiceInformationBaseType_Member_1_FIS:
		res = "FIS"
	case Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF:
		res = "OFIS_VHF"
	case Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF:
		res = "OFIS_HF"
	case Enum_CodeServiceInformationBaseType_Member_1_NOTAM:
		res = "NOTAM"
	case Enum_CodeServiceInformationBaseType_Member_1_INFO:
		res = "INFO"
	case Enum_CodeServiceInformationBaseType_Member_1_RAF:
		res = "RAF"
	case Enum_CodeServiceInformationBaseType_Member_1_METAR:
		res = "METAR"
	case Enum_CodeServiceInformationBaseType_Member_1_SIGMET:
		res = "SIGMET"
	case Enum_CodeServiceInformationBaseType_Member_1_TWEB:
		res = "TWEB"
	case Enum_CodeServiceInformationBaseType_Member_1_TAF:
		res = "TAF"
	case Enum_CodeServiceInformationBaseType_Member_1_VOLMET:
		res = "VOLMET"
	case Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER:
		res = "ALTIMETER"
	case Enum_CodeServiceInformationBaseType_Member_1_ASOS:
		res = "ASOS"
	case Enum_CodeServiceInformationBaseType_Member_1_AWOS:
		res = "AWOS"
	}
	return
}

func (enum_codeserviceinformationbasetype_member_1 *Enum_CodeServiceInformationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AFIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AFIS
		return
	case "AIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AIS
		return
	case "ATIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ATIS
		return
	case "BRIEFING":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_BRIEFING
		return
	case "FIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_FIS
		return
	case "OFIS_VHF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF
		return
	case "OFIS_HF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF
		return
	case "NOTAM":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_NOTAM
		return
	case "INFO":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_INFO
		return
	case "RAF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_RAF
		return
	case "METAR":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_METAR
		return
	case "SIGMET":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_SIGMET
		return
	case "TWEB":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_TWEB
		return
	case "TAF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_TAF
		return
	case "VOLMET":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_VOLMET
		return
	case "ALTIMETER":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER
		return
	case "ASOS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ASOS
		return
	case "AWOS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AWOS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeserviceinformationbasetype_member_1 *Enum_CodeServiceInformationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeServiceInformationBaseType_Member_1_AFIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AFIS
	case "Enum_CodeServiceInformationBaseType_Member_1_AIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AIS
	case "Enum_CodeServiceInformationBaseType_Member_1_ATIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ATIS
	case "Enum_CodeServiceInformationBaseType_Member_1_BRIEFING":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_BRIEFING
	case "Enum_CodeServiceInformationBaseType_Member_1_FIS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_FIS
	case "Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF
	case "Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF
	case "Enum_CodeServiceInformationBaseType_Member_1_NOTAM":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_NOTAM
	case "Enum_CodeServiceInformationBaseType_Member_1_INFO":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_INFO
	case "Enum_CodeServiceInformationBaseType_Member_1_RAF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_RAF
	case "Enum_CodeServiceInformationBaseType_Member_1_METAR":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_METAR
	case "Enum_CodeServiceInformationBaseType_Member_1_SIGMET":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_SIGMET
	case "Enum_CodeServiceInformationBaseType_Member_1_TWEB":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_TWEB
	case "Enum_CodeServiceInformationBaseType_Member_1_TAF":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_TAF
	case "Enum_CodeServiceInformationBaseType_Member_1_VOLMET":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_VOLMET
	case "Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER
	case "Enum_CodeServiceInformationBaseType_Member_1_ASOS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_ASOS
	case "Enum_CodeServiceInformationBaseType_Member_1_AWOS":
		*enum_codeserviceinformationbasetype_member_1 = Enum_CodeServiceInformationBaseType_Member_1_AWOS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeserviceinformationbasetype_member_1 *Enum_CodeServiceInformationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeserviceinformationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceInformationBaseType_Member_1_AFIS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_AFIS"
	case Enum_CodeServiceInformationBaseType_Member_1_AIS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_AIS"
	case Enum_CodeServiceInformationBaseType_Member_1_ATIS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_ATIS"
	case Enum_CodeServiceInformationBaseType_Member_1_BRIEFING:
		res = "Enum_CodeServiceInformationBaseType_Member_1_BRIEFING"
	case Enum_CodeServiceInformationBaseType_Member_1_FIS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_FIS"
	case Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF:
		res = "Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF"
	case Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF:
		res = "Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF"
	case Enum_CodeServiceInformationBaseType_Member_1_NOTAM:
		res = "Enum_CodeServiceInformationBaseType_Member_1_NOTAM"
	case Enum_CodeServiceInformationBaseType_Member_1_INFO:
		res = "Enum_CodeServiceInformationBaseType_Member_1_INFO"
	case Enum_CodeServiceInformationBaseType_Member_1_RAF:
		res = "Enum_CodeServiceInformationBaseType_Member_1_RAF"
	case Enum_CodeServiceInformationBaseType_Member_1_METAR:
		res = "Enum_CodeServiceInformationBaseType_Member_1_METAR"
	case Enum_CodeServiceInformationBaseType_Member_1_SIGMET:
		res = "Enum_CodeServiceInformationBaseType_Member_1_SIGMET"
	case Enum_CodeServiceInformationBaseType_Member_1_TWEB:
		res = "Enum_CodeServiceInformationBaseType_Member_1_TWEB"
	case Enum_CodeServiceInformationBaseType_Member_1_TAF:
		res = "Enum_CodeServiceInformationBaseType_Member_1_TAF"
	case Enum_CodeServiceInformationBaseType_Member_1_VOLMET:
		res = "Enum_CodeServiceInformationBaseType_Member_1_VOLMET"
	case Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER:
		res = "Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER"
	case Enum_CodeServiceInformationBaseType_Member_1_ASOS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_ASOS"
	case Enum_CodeServiceInformationBaseType_Member_1_AWOS:
		res = "Enum_CodeServiceInformationBaseType_Member_1_AWOS"
	}
	return
}

func (enum_codeserviceinformationbasetype_member_1 Enum_CodeServiceInformationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_AFIS")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_AIS")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_ATIS")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_BRIEFING")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_FIS")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_OFIS_VHF")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_OFIS_HF")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_NOTAM")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_INFO")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_RAF")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_METAR")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_SIGMET")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_TWEB")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_TAF")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_VOLMET")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_ALTIMETER")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_ASOS")
	res = append(res, "Enum_CodeServiceInformationBaseType_Member_1_AWOS")

	return
}

func (enum_codeserviceinformationbasetype_member_1 Enum_CodeServiceInformationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AFIS")
	res = append(res, "AIS")
	res = append(res, "ATIS")
	res = append(res, "BRIEFING")
	res = append(res, "FIS")
	res = append(res, "OFIS_VHF")
	res = append(res, "OFIS_HF")
	res = append(res, "NOTAM")
	res = append(res, "INFO")
	res = append(res, "RAF")
	res = append(res, "METAR")
	res = append(res, "SIGMET")
	res = append(res, "TWEB")
	res = append(res, "TAF")
	res = append(res, "VOLMET")
	res = append(res, "ALTIMETER")
	res = append(res, "ASOS")
	res = append(res, "AWOS")

	return
}

// Utility function for Enum_CodeServiceSARBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeservicesarbasetype_member_1 Enum_CodeServiceSARBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeservicesarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceSARBaseType_Member_1_ALRS:
		res = "ALRS"
	case Enum_CodeServiceSARBaseType_Member_1_SAR:
		res = "SAR"
	case Enum_CodeServiceSARBaseType_Member_1_RCC:
		res = "RCC"
	}
	return
}

func (enum_codeservicesarbasetype_member_1 *Enum_CodeServiceSARBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ALRS":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_ALRS
		return
	case "SAR":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_SAR
		return
	case "RCC":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_RCC
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeservicesarbasetype_member_1 *Enum_CodeServiceSARBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeServiceSARBaseType_Member_1_ALRS":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_ALRS
	case "Enum_CodeServiceSARBaseType_Member_1_SAR":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_SAR
	case "Enum_CodeServiceSARBaseType_Member_1_RCC":
		*enum_codeservicesarbasetype_member_1 = Enum_CodeServiceSARBaseType_Member_1_RCC
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeservicesarbasetype_member_1 *Enum_CodeServiceSARBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeservicesarbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeServiceSARBaseType_Member_1_ALRS:
		res = "Enum_CodeServiceSARBaseType_Member_1_ALRS"
	case Enum_CodeServiceSARBaseType_Member_1_SAR:
		res = "Enum_CodeServiceSARBaseType_Member_1_SAR"
	case Enum_CodeServiceSARBaseType_Member_1_RCC:
		res = "Enum_CodeServiceSARBaseType_Member_1_RCC"
	}
	return
}

func (enum_codeservicesarbasetype_member_1 Enum_CodeServiceSARBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeServiceSARBaseType_Member_1_ALRS")
	res = append(res, "Enum_CodeServiceSARBaseType_Member_1_SAR")
	res = append(res, "Enum_CodeServiceSARBaseType_Member_1_RCC")

	return
}

func (enum_codeservicesarbasetype_member_1 Enum_CodeServiceSARBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ALRS")
	res = append(res, "SAR")
	res = append(res, "RCC")

	return
}

// Utility function for Enum_CodeSpecialDateBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codespecialdatebasetype_member_1 Enum_CodeSpecialDateBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codespecialdatebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialDateBaseType_Member_1_HOL:
		res = "HOL"
	case Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI:
		res = "BUSY_FRI"
	}
	return
}

func (enum_codespecialdatebasetype_member_1 *Enum_CodeSpecialDateBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "HOL":
		*enum_codespecialdatebasetype_member_1 = Enum_CodeSpecialDateBaseType_Member_1_HOL
		return
	case "BUSY_FRI":
		*enum_codespecialdatebasetype_member_1 = Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codespecialdatebasetype_member_1 *Enum_CodeSpecialDateBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSpecialDateBaseType_Member_1_HOL":
		*enum_codespecialdatebasetype_member_1 = Enum_CodeSpecialDateBaseType_Member_1_HOL
	case "Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI":
		*enum_codespecialdatebasetype_member_1 = Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codespecialdatebasetype_member_1 *Enum_CodeSpecialDateBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codespecialdatebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSpecialDateBaseType_Member_1_HOL:
		res = "Enum_CodeSpecialDateBaseType_Member_1_HOL"
	case Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI:
		res = "Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI"
	}
	return
}

func (enum_codespecialdatebasetype_member_1 Enum_CodeSpecialDateBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSpecialDateBaseType_Member_1_HOL")
	res = append(res, "Enum_CodeSpecialDateBaseType_Member_1_BUSY_FRI")

	return
}

func (enum_codespecialdatebasetype_member_1 Enum_CodeSpecialDateBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "HOL")
	res = append(res, "BUSY_FRI")

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

// Utility function for Enum_CodeStatusServiceBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codestatusservicebasetype_member_1 Enum_CodeStatusServiceBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codestatusservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusServiceBaseType_Member_1_NORMAL:
		res = "NORMAL"
	case Enum_CodeStatusServiceBaseType_Member_1_LIMITED:
		res = "LIMITED"
	case Enum_CodeStatusServiceBaseType_Member_1_ONTEST:
		res = "ONTEST"
	case Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE:
		res = "UNSERVICEABLE"
	}
	return
}

func (enum_codestatusservicebasetype_member_1 *Enum_CodeStatusServiceBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NORMAL":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_NORMAL
		return
	case "LIMITED":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_LIMITED
		return
	case "ONTEST":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_ONTEST
		return
	case "UNSERVICEABLE":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codestatusservicebasetype_member_1 *Enum_CodeStatusServiceBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeStatusServiceBaseType_Member_1_NORMAL":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_NORMAL
	case "Enum_CodeStatusServiceBaseType_Member_1_LIMITED":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_LIMITED
	case "Enum_CodeStatusServiceBaseType_Member_1_ONTEST":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_ONTEST
	case "Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE":
		*enum_codestatusservicebasetype_member_1 = Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codestatusservicebasetype_member_1 *Enum_CodeStatusServiceBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codestatusservicebasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeStatusServiceBaseType_Member_1_NORMAL:
		res = "Enum_CodeStatusServiceBaseType_Member_1_NORMAL"
	case Enum_CodeStatusServiceBaseType_Member_1_LIMITED:
		res = "Enum_CodeStatusServiceBaseType_Member_1_LIMITED"
	case Enum_CodeStatusServiceBaseType_Member_1_ONTEST:
		res = "Enum_CodeStatusServiceBaseType_Member_1_ONTEST"
	case Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE:
		res = "Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE"
	}
	return
}

func (enum_codestatusservicebasetype_member_1 Enum_CodeStatusServiceBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeStatusServiceBaseType_Member_1_NORMAL")
	res = append(res, "Enum_CodeStatusServiceBaseType_Member_1_LIMITED")
	res = append(res, "Enum_CodeStatusServiceBaseType_Member_1_ONTEST")
	res = append(res, "Enum_CodeStatusServiceBaseType_Member_1_UNSERVICEABLE")

	return
}

func (enum_codestatusservicebasetype_member_1 Enum_CodeStatusServiceBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NORMAL")
	res = append(res, "LIMITED")
	res = append(res, "ONTEST")
	res = append(res, "UNSERVICEABLE")

	return
}

// Utility function for Enum_CodeSystemActivationBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codesystemactivationbasetype_member_1 Enum_CodeSystemActivationBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codesystemactivationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSystemActivationBaseType_Member_1_ON:
		res = "ON"
	case Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF:
		res = "ON_OR_OFF"
	case Enum_CodeSystemActivationBaseType_Member_1_OFF:
		res = "OFF"
	}
	return
}

func (enum_codesystemactivationbasetype_member_1 *Enum_CodeSystemActivationBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ON":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_ON
		return
	case "ON_OR_OFF":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF
		return
	case "OFF":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_OFF
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codesystemactivationbasetype_member_1 *Enum_CodeSystemActivationBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeSystemActivationBaseType_Member_1_ON":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_ON
	case "Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF
	case "Enum_CodeSystemActivationBaseType_Member_1_OFF":
		*enum_codesystemactivationbasetype_member_1 = Enum_CodeSystemActivationBaseType_Member_1_OFF
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codesystemactivationbasetype_member_1 *Enum_CodeSystemActivationBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codesystemactivationbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeSystemActivationBaseType_Member_1_ON:
		res = "Enum_CodeSystemActivationBaseType_Member_1_ON"
	case Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF:
		res = "Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF"
	case Enum_CodeSystemActivationBaseType_Member_1_OFF:
		res = "Enum_CodeSystemActivationBaseType_Member_1_OFF"
	}
	return
}

func (enum_codesystemactivationbasetype_member_1 Enum_CodeSystemActivationBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeSystemActivationBaseType_Member_1_ON")
	res = append(res, "Enum_CodeSystemActivationBaseType_Member_1_ON_OR_OFF")
	res = append(res, "Enum_CodeSystemActivationBaseType_Member_1_OFF")

	return
}

func (enum_codesystemactivationbasetype_member_1 Enum_CodeSystemActivationBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ON")
	res = append(res, "ON_OR_OFF")
	res = append(res, "OFF")

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

// Utility function for Enum_CodeUnitBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeunitbasetype_member_1 Enum_CodeUnitBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeunitbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUnitBaseType_Member_1_ACC:
		res = "ACC"
	case Enum_CodeUnitBaseType_Member_1_ADSU:
		res = "ADSU"
	case Enum_CodeUnitBaseType_Member_1_ADVC:
		res = "ADVC"
	case Enum_CodeUnitBaseType_Member_1_ALPS:
		res = "ALPS"
	case Enum_CodeUnitBaseType_Member_1_AOF:
		res = "AOF"
	case Enum_CodeUnitBaseType_Member_1_APP:
		res = "APP"
	case Enum_CodeUnitBaseType_Member_1_APP_ARR:
		res = "APP_ARR"
	case Enum_CodeUnitBaseType_Member_1_APP_DEP:
		res = "APP_DEP"
	case Enum_CodeUnitBaseType_Member_1_ARO:
		res = "ARO"
	case Enum_CodeUnitBaseType_Member_1_ATCC:
		res = "ATCC"
	case Enum_CodeUnitBaseType_Member_1_ATFMU:
		res = "ATFMU"
	case Enum_CodeUnitBaseType_Member_1_ATMU:
		res = "ATMU"
	case Enum_CodeUnitBaseType_Member_1_ATSU:
		res = "ATSU"
	case Enum_CodeUnitBaseType_Member_1_BOF:
		res = "BOF"
	case Enum_CodeUnitBaseType_Member_1_BS:
		res = "BS"
	case Enum_CodeUnitBaseType_Member_1_COM:
		res = "COM"
	case Enum_CodeUnitBaseType_Member_1_FCST:
		res = "FCST"
	case Enum_CodeUnitBaseType_Member_1_FIC:
		res = "FIC"
	case Enum_CodeUnitBaseType_Member_1_GCA:
		res = "GCA"
	case Enum_CodeUnitBaseType_Member_1_MET:
		res = "MET"
	case Enum_CodeUnitBaseType_Member_1_MWO:
		res = "MWO"
	case Enum_CodeUnitBaseType_Member_1_NOF:
		res = "NOF"
	case Enum_CodeUnitBaseType_Member_1_OAC:
		res = "OAC"
	case Enum_CodeUnitBaseType_Member_1_PAR:
		res = "PAR"
	case Enum_CodeUnitBaseType_Member_1_RAD:
		res = "RAD"
	case Enum_CodeUnitBaseType_Member_1_RAFC:
		res = "RAFC"
	case Enum_CodeUnitBaseType_Member_1_RCC:
		res = "RCC"
	case Enum_CodeUnitBaseType_Member_1_RSC:
		res = "RSC"
	case Enum_CodeUnitBaseType_Member_1_SAR:
		res = "SAR"
	case Enum_CodeUnitBaseType_Member_1_SMC:
		res = "SMC"
	case Enum_CodeUnitBaseType_Member_1_SMR:
		res = "SMR"
	case Enum_CodeUnitBaseType_Member_1_SRA:
		res = "SRA"
	case Enum_CodeUnitBaseType_Member_1_SSR:
		res = "SSR"
	case Enum_CodeUnitBaseType_Member_1_TAR:
		res = "TAR"
	case Enum_CodeUnitBaseType_Member_1_TWR:
		res = "TWR"
	case Enum_CodeUnitBaseType_Member_1_UAC:
		res = "UAC"
	case Enum_CodeUnitBaseType_Member_1_UDF:
		res = "UDF"
	case Enum_CodeUnitBaseType_Member_1_UIC:
		res = "UIC"
	case Enum_CodeUnitBaseType_Member_1_VDF:
		res = "VDF"
	case Enum_CodeUnitBaseType_Member_1_WAFC:
		res = "WAFC"
	case Enum_CodeUnitBaseType_Member_1_ARTCC:
		res = "ARTCC"
	case Enum_CodeUnitBaseType_Member_1_FSS:
		res = "FSS"
	case Enum_CodeUnitBaseType_Member_1_TRACON:
		res = "TRACON"
	case Enum_CodeUnitBaseType_Member_1_MIL:
		res = "MIL"
	case Enum_CodeUnitBaseType_Member_1_MILOPS:
		res = "MILOPS"
	}
	return
}

func (enum_codeunitbasetype_member_1 *Enum_CodeUnitBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ACC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ACC
		return
	case "ADSU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ADSU
		return
	case "ADVC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ADVC
		return
	case "ALPS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ALPS
		return
	case "AOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_AOF
		return
	case "APP":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP
		return
	case "APP_ARR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP_ARR
		return
	case "APP_DEP":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP_DEP
		return
	case "ARO":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ARO
		return
	case "ATCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATCC
		return
	case "ATFMU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATFMU
		return
	case "ATMU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATMU
		return
	case "ATSU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATSU
		return
	case "BOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_BOF
		return
	case "BS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_BS
		return
	case "COM":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_COM
		return
	case "FCST":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FCST
		return
	case "FIC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FIC
		return
	case "GCA":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_GCA
		return
	case "MET":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MET
		return
	case "MWO":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MWO
		return
	case "NOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_NOF
		return
	case "OAC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_OAC
		return
	case "PAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_PAR
		return
	case "RAD":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RAD
		return
	case "RAFC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RAFC
		return
	case "RCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RCC
		return
	case "RSC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RSC
		return
	case "SAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SAR
		return
	case "SMC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SMC
		return
	case "SMR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SMR
		return
	case "SRA":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SRA
		return
	case "SSR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SSR
		return
	case "TAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TAR
		return
	case "TWR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TWR
		return
	case "UAC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UAC
		return
	case "UDF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UDF
		return
	case "UIC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UIC
		return
	case "VDF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_VDF
		return
	case "WAFC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_WAFC
		return
	case "ARTCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ARTCC
		return
	case "FSS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FSS
		return
	case "TRACON":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TRACON
		return
	case "MIL":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MIL
		return
	case "MILOPS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MILOPS
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeunitbasetype_member_1 *Enum_CodeUnitBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeUnitBaseType_Member_1_ACC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ACC
	case "Enum_CodeUnitBaseType_Member_1_ADSU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ADSU
	case "Enum_CodeUnitBaseType_Member_1_ADVC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ADVC
	case "Enum_CodeUnitBaseType_Member_1_ALPS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ALPS
	case "Enum_CodeUnitBaseType_Member_1_AOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_AOF
	case "Enum_CodeUnitBaseType_Member_1_APP":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP
	case "Enum_CodeUnitBaseType_Member_1_APP_ARR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP_ARR
	case "Enum_CodeUnitBaseType_Member_1_APP_DEP":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_APP_DEP
	case "Enum_CodeUnitBaseType_Member_1_ARO":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ARO
	case "Enum_CodeUnitBaseType_Member_1_ATCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATCC
	case "Enum_CodeUnitBaseType_Member_1_ATFMU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATFMU
	case "Enum_CodeUnitBaseType_Member_1_ATMU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATMU
	case "Enum_CodeUnitBaseType_Member_1_ATSU":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ATSU
	case "Enum_CodeUnitBaseType_Member_1_BOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_BOF
	case "Enum_CodeUnitBaseType_Member_1_BS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_BS
	case "Enum_CodeUnitBaseType_Member_1_COM":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_COM
	case "Enum_CodeUnitBaseType_Member_1_FCST":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FCST
	case "Enum_CodeUnitBaseType_Member_1_FIC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FIC
	case "Enum_CodeUnitBaseType_Member_1_GCA":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_GCA
	case "Enum_CodeUnitBaseType_Member_1_MET":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MET
	case "Enum_CodeUnitBaseType_Member_1_MWO":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MWO
	case "Enum_CodeUnitBaseType_Member_1_NOF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_NOF
	case "Enum_CodeUnitBaseType_Member_1_OAC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_OAC
	case "Enum_CodeUnitBaseType_Member_1_PAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_PAR
	case "Enum_CodeUnitBaseType_Member_1_RAD":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RAD
	case "Enum_CodeUnitBaseType_Member_1_RAFC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RAFC
	case "Enum_CodeUnitBaseType_Member_1_RCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RCC
	case "Enum_CodeUnitBaseType_Member_1_RSC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_RSC
	case "Enum_CodeUnitBaseType_Member_1_SAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SAR
	case "Enum_CodeUnitBaseType_Member_1_SMC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SMC
	case "Enum_CodeUnitBaseType_Member_1_SMR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SMR
	case "Enum_CodeUnitBaseType_Member_1_SRA":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SRA
	case "Enum_CodeUnitBaseType_Member_1_SSR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_SSR
	case "Enum_CodeUnitBaseType_Member_1_TAR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TAR
	case "Enum_CodeUnitBaseType_Member_1_TWR":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TWR
	case "Enum_CodeUnitBaseType_Member_1_UAC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UAC
	case "Enum_CodeUnitBaseType_Member_1_UDF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UDF
	case "Enum_CodeUnitBaseType_Member_1_UIC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_UIC
	case "Enum_CodeUnitBaseType_Member_1_VDF":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_VDF
	case "Enum_CodeUnitBaseType_Member_1_WAFC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_WAFC
	case "Enum_CodeUnitBaseType_Member_1_ARTCC":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_ARTCC
	case "Enum_CodeUnitBaseType_Member_1_FSS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_FSS
	case "Enum_CodeUnitBaseType_Member_1_TRACON":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_TRACON
	case "Enum_CodeUnitBaseType_Member_1_MIL":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MIL
	case "Enum_CodeUnitBaseType_Member_1_MILOPS":
		*enum_codeunitbasetype_member_1 = Enum_CodeUnitBaseType_Member_1_MILOPS
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeunitbasetype_member_1 *Enum_CodeUnitBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeunitbasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUnitBaseType_Member_1_ACC:
		res = "Enum_CodeUnitBaseType_Member_1_ACC"
	case Enum_CodeUnitBaseType_Member_1_ADSU:
		res = "Enum_CodeUnitBaseType_Member_1_ADSU"
	case Enum_CodeUnitBaseType_Member_1_ADVC:
		res = "Enum_CodeUnitBaseType_Member_1_ADVC"
	case Enum_CodeUnitBaseType_Member_1_ALPS:
		res = "Enum_CodeUnitBaseType_Member_1_ALPS"
	case Enum_CodeUnitBaseType_Member_1_AOF:
		res = "Enum_CodeUnitBaseType_Member_1_AOF"
	case Enum_CodeUnitBaseType_Member_1_APP:
		res = "Enum_CodeUnitBaseType_Member_1_APP"
	case Enum_CodeUnitBaseType_Member_1_APP_ARR:
		res = "Enum_CodeUnitBaseType_Member_1_APP_ARR"
	case Enum_CodeUnitBaseType_Member_1_APP_DEP:
		res = "Enum_CodeUnitBaseType_Member_1_APP_DEP"
	case Enum_CodeUnitBaseType_Member_1_ARO:
		res = "Enum_CodeUnitBaseType_Member_1_ARO"
	case Enum_CodeUnitBaseType_Member_1_ATCC:
		res = "Enum_CodeUnitBaseType_Member_1_ATCC"
	case Enum_CodeUnitBaseType_Member_1_ATFMU:
		res = "Enum_CodeUnitBaseType_Member_1_ATFMU"
	case Enum_CodeUnitBaseType_Member_1_ATMU:
		res = "Enum_CodeUnitBaseType_Member_1_ATMU"
	case Enum_CodeUnitBaseType_Member_1_ATSU:
		res = "Enum_CodeUnitBaseType_Member_1_ATSU"
	case Enum_CodeUnitBaseType_Member_1_BOF:
		res = "Enum_CodeUnitBaseType_Member_1_BOF"
	case Enum_CodeUnitBaseType_Member_1_BS:
		res = "Enum_CodeUnitBaseType_Member_1_BS"
	case Enum_CodeUnitBaseType_Member_1_COM:
		res = "Enum_CodeUnitBaseType_Member_1_COM"
	case Enum_CodeUnitBaseType_Member_1_FCST:
		res = "Enum_CodeUnitBaseType_Member_1_FCST"
	case Enum_CodeUnitBaseType_Member_1_FIC:
		res = "Enum_CodeUnitBaseType_Member_1_FIC"
	case Enum_CodeUnitBaseType_Member_1_GCA:
		res = "Enum_CodeUnitBaseType_Member_1_GCA"
	case Enum_CodeUnitBaseType_Member_1_MET:
		res = "Enum_CodeUnitBaseType_Member_1_MET"
	case Enum_CodeUnitBaseType_Member_1_MWO:
		res = "Enum_CodeUnitBaseType_Member_1_MWO"
	case Enum_CodeUnitBaseType_Member_1_NOF:
		res = "Enum_CodeUnitBaseType_Member_1_NOF"
	case Enum_CodeUnitBaseType_Member_1_OAC:
		res = "Enum_CodeUnitBaseType_Member_1_OAC"
	case Enum_CodeUnitBaseType_Member_1_PAR:
		res = "Enum_CodeUnitBaseType_Member_1_PAR"
	case Enum_CodeUnitBaseType_Member_1_RAD:
		res = "Enum_CodeUnitBaseType_Member_1_RAD"
	case Enum_CodeUnitBaseType_Member_1_RAFC:
		res = "Enum_CodeUnitBaseType_Member_1_RAFC"
	case Enum_CodeUnitBaseType_Member_1_RCC:
		res = "Enum_CodeUnitBaseType_Member_1_RCC"
	case Enum_CodeUnitBaseType_Member_1_RSC:
		res = "Enum_CodeUnitBaseType_Member_1_RSC"
	case Enum_CodeUnitBaseType_Member_1_SAR:
		res = "Enum_CodeUnitBaseType_Member_1_SAR"
	case Enum_CodeUnitBaseType_Member_1_SMC:
		res = "Enum_CodeUnitBaseType_Member_1_SMC"
	case Enum_CodeUnitBaseType_Member_1_SMR:
		res = "Enum_CodeUnitBaseType_Member_1_SMR"
	case Enum_CodeUnitBaseType_Member_1_SRA:
		res = "Enum_CodeUnitBaseType_Member_1_SRA"
	case Enum_CodeUnitBaseType_Member_1_SSR:
		res = "Enum_CodeUnitBaseType_Member_1_SSR"
	case Enum_CodeUnitBaseType_Member_1_TAR:
		res = "Enum_CodeUnitBaseType_Member_1_TAR"
	case Enum_CodeUnitBaseType_Member_1_TWR:
		res = "Enum_CodeUnitBaseType_Member_1_TWR"
	case Enum_CodeUnitBaseType_Member_1_UAC:
		res = "Enum_CodeUnitBaseType_Member_1_UAC"
	case Enum_CodeUnitBaseType_Member_1_UDF:
		res = "Enum_CodeUnitBaseType_Member_1_UDF"
	case Enum_CodeUnitBaseType_Member_1_UIC:
		res = "Enum_CodeUnitBaseType_Member_1_UIC"
	case Enum_CodeUnitBaseType_Member_1_VDF:
		res = "Enum_CodeUnitBaseType_Member_1_VDF"
	case Enum_CodeUnitBaseType_Member_1_WAFC:
		res = "Enum_CodeUnitBaseType_Member_1_WAFC"
	case Enum_CodeUnitBaseType_Member_1_ARTCC:
		res = "Enum_CodeUnitBaseType_Member_1_ARTCC"
	case Enum_CodeUnitBaseType_Member_1_FSS:
		res = "Enum_CodeUnitBaseType_Member_1_FSS"
	case Enum_CodeUnitBaseType_Member_1_TRACON:
		res = "Enum_CodeUnitBaseType_Member_1_TRACON"
	case Enum_CodeUnitBaseType_Member_1_MIL:
		res = "Enum_CodeUnitBaseType_Member_1_MIL"
	case Enum_CodeUnitBaseType_Member_1_MILOPS:
		res = "Enum_CodeUnitBaseType_Member_1_MILOPS"
	}
	return
}

func (enum_codeunitbasetype_member_1 Enum_CodeUnitBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ACC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ADSU")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ADVC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ALPS")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_AOF")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_APP")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_APP_ARR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_APP_DEP")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ARO")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ATCC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ATFMU")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ATMU")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ATSU")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_BOF")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_BS")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_COM")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_FCST")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_FIC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_GCA")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_MET")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_MWO")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_NOF")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_OAC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_PAR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_RAD")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_RAFC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_RCC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_RSC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_SAR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_SMC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_SMR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_SRA")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_SSR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_TAR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_TWR")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_UAC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_UDF")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_UIC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_VDF")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_WAFC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_ARTCC")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_FSS")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_TRACON")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_MIL")
	res = append(res, "Enum_CodeUnitBaseType_Member_1_MILOPS")

	return
}

func (enum_codeunitbasetype_member_1 Enum_CodeUnitBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ACC")
	res = append(res, "ADSU")
	res = append(res, "ADVC")
	res = append(res, "ALPS")
	res = append(res, "AOF")
	res = append(res, "APP")
	res = append(res, "APP_ARR")
	res = append(res, "APP_DEP")
	res = append(res, "ARO")
	res = append(res, "ATCC")
	res = append(res, "ATFMU")
	res = append(res, "ATMU")
	res = append(res, "ATSU")
	res = append(res, "BOF")
	res = append(res, "BS")
	res = append(res, "COM")
	res = append(res, "FCST")
	res = append(res, "FIC")
	res = append(res, "GCA")
	res = append(res, "MET")
	res = append(res, "MWO")
	res = append(res, "NOF")
	res = append(res, "OAC")
	res = append(res, "PAR")
	res = append(res, "RAD")
	res = append(res, "RAFC")
	res = append(res, "RCC")
	res = append(res, "RSC")
	res = append(res, "SAR")
	res = append(res, "SMC")
	res = append(res, "SMR")
	res = append(res, "SRA")
	res = append(res, "SSR")
	res = append(res, "TAR")
	res = append(res, "TWR")
	res = append(res, "UAC")
	res = append(res, "UDF")
	res = append(res, "UIC")
	res = append(res, "VDF")
	res = append(res, "WAFC")
	res = append(res, "ARTCC")
	res = append(res, "FSS")
	res = append(res, "TRACON")
	res = append(res, "MIL")
	res = append(res, "MILOPS")

	return
}

// Utility function for Enum_CodeUnitDependencyBaseType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_codeunitdependencybasetype_member_1 Enum_CodeUnitDependencyBaseType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_codeunitdependencybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUnitDependencyBaseType_Member_1_OWNER:
		res = "OWNER"
	case Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER:
		res = "PROVIDER"
	case Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE:
		res = "ALTERNATE"
	}
	return
}

func (enum_codeunitdependencybasetype_member_1 *Enum_CodeUnitDependencyBaseType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "OWNER":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_OWNER
		return
	case "PROVIDER":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER
		return
	case "ALTERNATE":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE
		return
	default:
		return errUnkownEnum
	}
}

func (enum_codeunitdependencybasetype_member_1 *Enum_CodeUnitDependencyBaseType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CodeUnitDependencyBaseType_Member_1_OWNER":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_OWNER
	case "Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER
	case "Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE":
		*enum_codeunitdependencybasetype_member_1 = Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_codeunitdependencybasetype_member_1 *Enum_CodeUnitDependencyBaseType_Member_1) ToCodeString() (res string) {

	switch *enum_codeunitdependencybasetype_member_1 {
	// insertion code per enum code
	case Enum_CodeUnitDependencyBaseType_Member_1_OWNER:
		res = "Enum_CodeUnitDependencyBaseType_Member_1_OWNER"
	case Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER:
		res = "Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER"
	case Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE:
		res = "Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE"
	}
	return
}

func (enum_codeunitdependencybasetype_member_1 Enum_CodeUnitDependencyBaseType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CodeUnitDependencyBaseType_Member_1_OWNER")
	res = append(res, "Enum_CodeUnitDependencyBaseType_Member_1_PROVIDER")
	res = append(res, "Enum_CodeUnitDependencyBaseType_Member_1_ALTERNATE")

	return
}

func (enum_codeunitdependencybasetype_member_1 Enum_CodeUnitDependencyBaseType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "OWNER")
	res = append(res, "PROVIDER")
	res = append(res, "ALTERNATE")

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
