// generated code - do not edit
package gml

// insertion point of enum utility functions
// Utility function for Enum_AggregationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_aggregationtype Enum_AggregationType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_aggregationtype {
	// insertion code per enum code
	case Enum_AggregationType_Sequence:
		res = "sequence"
	}
	return
}

func (enum_aggregationtype *Enum_AggregationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sequence":
		*enum_aggregationtype = Enum_AggregationType_Sequence
		return
	default:
		return errUnkownEnum
	}
}

func (enum_aggregationtype *Enum_AggregationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_AggregationType_Sequence":
		*enum_aggregationtype = Enum_AggregationType_Sequence
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_aggregationtype *Enum_AggregationType) ToCodeString() (res string) {

	switch *enum_aggregationtype {
	// insertion code per enum code
	case Enum_AggregationType_Sequence:
		res = "Enum_AggregationType_Sequence"
	}
	return
}

func (enum_aggregationtype Enum_AggregationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_AggregationType_Sequence")

	return
}

func (enum_aggregationtype Enum_AggregationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sequence")

	return
}

// Utility function for Enum_CurveInterpolationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_curveinterpolationtype Enum_CurveInterpolationType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_curveinterpolationtype {
	// insertion code per enum code
	case Enum_CurveInterpolationType_Linear:
		res = "linear"
	case Enum_CurveInterpolationType_Geodesic:
		res = "geodesic"
	case Enum_CurveInterpolationType_CircularArc3Points:
		res = "circularArc3Points"
	case Enum_CurveInterpolationType_CircularArcCenterPointWithRadius:
		res = "circularArcCenterPointWithRadius"
	}
	return
}

func (enum_curveinterpolationtype *Enum_CurveInterpolationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "linear":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Linear
		return
	case "geodesic":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Geodesic
		return
	case "circularArc3Points":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArc3Points
		return
	case "circularArcCenterPointWithRadius":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArcCenterPointWithRadius
		return
	default:
		return errUnkownEnum
	}
}

func (enum_curveinterpolationtype *Enum_CurveInterpolationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CurveInterpolationType_Linear":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Linear
	case "Enum_CurveInterpolationType_Geodesic":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Geodesic
	case "Enum_CurveInterpolationType_CircularArc3Points":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArc3Points
	case "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArcCenterPointWithRadius
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_curveinterpolationtype *Enum_CurveInterpolationType) ToCodeString() (res string) {

	switch *enum_curveinterpolationtype {
	// insertion code per enum code
	case Enum_CurveInterpolationType_Linear:
		res = "Enum_CurveInterpolationType_Linear"
	case Enum_CurveInterpolationType_Geodesic:
		res = "Enum_CurveInterpolationType_Geodesic"
	case Enum_CurveInterpolationType_CircularArc3Points:
		res = "Enum_CurveInterpolationType_CircularArc3Points"
	case Enum_CurveInterpolationType_CircularArcCenterPointWithRadius:
		res = "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius"
	}
	return
}

func (enum_curveinterpolationtype Enum_CurveInterpolationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CurveInterpolationType_Linear")
	res = append(res, "Enum_CurveInterpolationType_Geodesic")
	res = append(res, "Enum_CurveInterpolationType_CircularArc3Points")
	res = append(res, "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius")

	return
}

func (enum_curveinterpolationtype Enum_CurveInterpolationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "linear")
	res = append(res, "geodesic")
	res = append(res, "circularArc3Points")
	res = append(res, "circularArcCenterPointWithRadius")

	return
}

// Utility function for Enum_NilReasonEnumeration_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_nilreasonenumeration_member_1 {
	// insertion code per enum code
	case Enum_NilReasonEnumeration_Member_1_Inapplicable:
		res = "inapplicable"
	case Enum_NilReasonEnumeration_Member_1_Missing:
		res = "missing"
	case Enum_NilReasonEnumeration_Member_1_Template:
		res = "template"
	case Enum_NilReasonEnumeration_Member_1_Unknown:
		res = "unknown"
	case Enum_NilReasonEnumeration_Member_1_Withheld:
		res = "withheld"
	}
	return
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "inapplicable":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Inapplicable
		return
	case "missing":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Missing
		return
	case "template":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Template
		return
	case "unknown":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Unknown
		return
	case "withheld":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Withheld
		return
	default:
		return errUnkownEnum
	}
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_NilReasonEnumeration_Member_1_Inapplicable":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Inapplicable
	case "Enum_NilReasonEnumeration_Member_1_Missing":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Missing
	case "Enum_NilReasonEnumeration_Member_1_Template":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Template
	case "Enum_NilReasonEnumeration_Member_1_Unknown":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Unknown
	case "Enum_NilReasonEnumeration_Member_1_Withheld":
		*enum_nilreasonenumeration_member_1 = Enum_NilReasonEnumeration_Member_1_Withheld
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_nilreasonenumeration_member_1 *Enum_NilReasonEnumeration_Member_1) ToCodeString() (res string) {

	switch *enum_nilreasonenumeration_member_1 {
	// insertion code per enum code
	case Enum_NilReasonEnumeration_Member_1_Inapplicable:
		res = "Enum_NilReasonEnumeration_Member_1_Inapplicable"
	case Enum_NilReasonEnumeration_Member_1_Missing:
		res = "Enum_NilReasonEnumeration_Member_1_Missing"
	case Enum_NilReasonEnumeration_Member_1_Template:
		res = "Enum_NilReasonEnumeration_Member_1_Template"
	case Enum_NilReasonEnumeration_Member_1_Unknown:
		res = "Enum_NilReasonEnumeration_Member_1_Unknown"
	case Enum_NilReasonEnumeration_Member_1_Withheld:
		res = "Enum_NilReasonEnumeration_Member_1_Withheld"
	}
	return
}

func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Inapplicable")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Missing")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Template")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Unknown")
	res = append(res, "Enum_NilReasonEnumeration_Member_1_Withheld")

	return
}

func (enum_nilreasonenumeration_member_1 Enum_NilReasonEnumeration_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "inapplicable")
	res = append(res, "missing")
	res = append(res, "template")
	res = append(res, "unknown")
	res = append(res, "withheld")

	return
}

// Utility function for Enum_SignType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_signtype Enum_SignType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_signtype {
	// insertion code per enum code
	case Enum_SignType_N__:
		res = "-"
	case Enum_SignType_N___1:
		res = "+"
	}
	return
}

func (enum_signtype *Enum_SignType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "-":
		*enum_signtype = Enum_SignType_N__
		return
	case "+":
		*enum_signtype = Enum_SignType_N___1
		return
	default:
		return errUnkownEnum
	}
}

func (enum_signtype *Enum_SignType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SignType_N__":
		*enum_signtype = Enum_SignType_N__
	case "Enum_SignType_N___1":
		*enum_signtype = Enum_SignType_N___1
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_signtype *Enum_SignType) ToCodeString() (res string) {

	switch *enum_signtype {
	// insertion code per enum code
	case Enum_SignType_N__:
		res = "Enum_SignType_N__"
	case Enum_SignType_N___1:
		res = "Enum_SignType_N___1"
	}
	return
}

func (enum_signtype Enum_SignType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SignType_N__")
	res = append(res, "Enum_SignType_N___1")

	return
}

func (enum_signtype Enum_SignType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "-")
	res = append(res, "+")

	return
}

// Utility function for Enum_SurfaceInterpolationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_surfaceinterpolationtype {
	// insertion code per enum code
	case Enum_SurfaceInterpolationType_Planar:
		res = "planar"
	}
	return
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "planar":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Planar
		return
	default:
		return errUnkownEnum
	}
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SurfaceInterpolationType_Planar":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Planar
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) ToCodeString() (res string) {

	switch *enum_surfaceinterpolationtype {
	// insertion code per enum code
	case Enum_SurfaceInterpolationType_Planar:
		res = "Enum_SurfaceInterpolationType_Planar"
	}
	return
}

func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SurfaceInterpolationType_Planar")

	return
}

func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "planar")

	return
}

// Utility function for Enum_TimeIndeterminateValueType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_timeindeterminatevaluetype {
	// insertion code per enum code
	case Enum_TimeIndeterminateValueType_Unknown:
		res = "unknown"
	}
	return
}

func (enum_timeindeterminatevaluetype *Enum_TimeIndeterminateValueType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "unknown":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Unknown
		return
	default:
		return errUnkownEnum
	}
}

func (enum_timeindeterminatevaluetype *Enum_TimeIndeterminateValueType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_TimeIndeterminateValueType_Unknown":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Unknown
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_timeindeterminatevaluetype *Enum_TimeIndeterminateValueType) ToCodeString() (res string) {

	switch *enum_timeindeterminatevaluetype {
	// insertion code per enum code
	case Enum_TimeIndeterminateValueType_Unknown:
		res = "Enum_TimeIndeterminateValueType_Unknown"
	}
	return
}

func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_TimeIndeterminateValueType_Unknown")

	return
}

func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "unknown")

	return
}

// Utility function for Enum_TimePositionUnion_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_timepositionunion_member_1 Enum_TimePositionUnion_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_timepositionunion_member_1 {
	// insertion code per enum code
	case Enum_TimePositionUnion_Member_1_Empty:
		res = ""
	}
	return
}

func (enum_timepositionunion_member_1 *Enum_TimePositionUnion_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "":
		*enum_timepositionunion_member_1 = Enum_TimePositionUnion_Member_1_Empty
		return
	default:
		return errUnkownEnum
	}
}

func (enum_timepositionunion_member_1 *Enum_TimePositionUnion_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_TimePositionUnion_Member_1_Empty":
		*enum_timepositionunion_member_1 = Enum_TimePositionUnion_Member_1_Empty
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_timepositionunion_member_1 *Enum_TimePositionUnion_Member_1) ToCodeString() (res string) {

	switch *enum_timepositionunion_member_1 {
	// insertion code per enum code
	case Enum_TimePositionUnion_Member_1_Empty:
		res = "Enum_TimePositionUnion_Member_1_Empty"
	}
	return
}

func (enum_timepositionunion_member_1 Enum_TimePositionUnion_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_TimePositionUnion_Member_1_Empty")

	return
}

func (enum_timepositionunion_member_1 Enum_TimePositionUnion_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "")

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
