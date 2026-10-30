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
	case Enum_AggregationType_Set:
		res = "set"
	case Enum_AggregationType_Bag:
		res = "bag"
	case Enum_AggregationType_Sequence:
		res = "sequence"
	case Enum_AggregationType_Array:
		res = "array"
	case Enum_AggregationType_Record:
		res = "record"
	case Enum_AggregationType_Table:
		res = "table"
	}
	return
}

func (enum_aggregationtype *Enum_AggregationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "set":
		*enum_aggregationtype = Enum_AggregationType_Set
		return
	case "bag":
		*enum_aggregationtype = Enum_AggregationType_Bag
		return
	case "sequence":
		*enum_aggregationtype = Enum_AggregationType_Sequence
		return
	case "array":
		*enum_aggregationtype = Enum_AggregationType_Array
		return
	case "record":
		*enum_aggregationtype = Enum_AggregationType_Record
		return
	case "table":
		*enum_aggregationtype = Enum_AggregationType_Table
		return
	default:
		return errUnkownEnum
	}
}

func (enum_aggregationtype *Enum_AggregationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_AggregationType_Set":
		*enum_aggregationtype = Enum_AggregationType_Set
	case "Enum_AggregationType_Bag":
		*enum_aggregationtype = Enum_AggregationType_Bag
	case "Enum_AggregationType_Sequence":
		*enum_aggregationtype = Enum_AggregationType_Sequence
	case "Enum_AggregationType_Array":
		*enum_aggregationtype = Enum_AggregationType_Array
	case "Enum_AggregationType_Record":
		*enum_aggregationtype = Enum_AggregationType_Record
	case "Enum_AggregationType_Table":
		*enum_aggregationtype = Enum_AggregationType_Table
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_aggregationtype *Enum_AggregationType) ToCodeString() (res string) {

	switch *enum_aggregationtype {
	// insertion code per enum code
	case Enum_AggregationType_Set:
		res = "Enum_AggregationType_Set"
	case Enum_AggregationType_Bag:
		res = "Enum_AggregationType_Bag"
	case Enum_AggregationType_Sequence:
		res = "Enum_AggregationType_Sequence"
	case Enum_AggregationType_Array:
		res = "Enum_AggregationType_Array"
	case Enum_AggregationType_Record:
		res = "Enum_AggregationType_Record"
	case Enum_AggregationType_Table:
		res = "Enum_AggregationType_Table"
	}
	return
}

func (enum_aggregationtype Enum_AggregationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_AggregationType_Set")
	res = append(res, "Enum_AggregationType_Bag")
	res = append(res, "Enum_AggregationType_Sequence")
	res = append(res, "Enum_AggregationType_Array")
	res = append(res, "Enum_AggregationType_Record")
	res = append(res, "Enum_AggregationType_Table")

	return
}

func (enum_aggregationtype Enum_AggregationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "set")
	res = append(res, "bag")
	res = append(res, "sequence")
	res = append(res, "array")
	res = append(res, "record")
	res = append(res, "table")

	return
}

// Utility function for Enum_CompassPointEnumeration
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_compasspointenumeration Enum_CompassPointEnumeration) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_compasspointenumeration {
	// insertion code per enum code
	case Enum_CompassPointEnumeration_N:
		res = "N"
	case Enum_CompassPointEnumeration_NNE:
		res = "NNE"
	case Enum_CompassPointEnumeration_NE:
		res = "NE"
	case Enum_CompassPointEnumeration_ENE:
		res = "ENE"
	case Enum_CompassPointEnumeration_E:
		res = "E"
	case Enum_CompassPointEnumeration_ESE:
		res = "ESE"
	case Enum_CompassPointEnumeration_SE:
		res = "SE"
	case Enum_CompassPointEnumeration_SSE:
		res = "SSE"
	case Enum_CompassPointEnumeration_S:
		res = "S"
	case Enum_CompassPointEnumeration_SSW:
		res = "SSW"
	case Enum_CompassPointEnumeration_SW:
		res = "SW"
	case Enum_CompassPointEnumeration_WSW:
		res = "WSW"
	case Enum_CompassPointEnumeration_W:
		res = "W"
	case Enum_CompassPointEnumeration_WNW:
		res = "WNW"
	case Enum_CompassPointEnumeration_NW:
		res = "NW"
	case Enum_CompassPointEnumeration_NNW:
		res = "NNW"
	}
	return
}

func (enum_compasspointenumeration *Enum_CompassPointEnumeration) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "N":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_N
		return
	case "NNE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NNE
		return
	case "NE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NE
		return
	case "ENE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_ENE
		return
	case "E":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_E
		return
	case "ESE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_ESE
		return
	case "SE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SE
		return
	case "SSE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SSE
		return
	case "S":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_S
		return
	case "SSW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SSW
		return
	case "SW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SW
		return
	case "WSW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_WSW
		return
	case "W":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_W
		return
	case "WNW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_WNW
		return
	case "NW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NW
		return
	case "NNW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NNW
		return
	default:
		return errUnkownEnum
	}
}

func (enum_compasspointenumeration *Enum_CompassPointEnumeration) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_CompassPointEnumeration_N":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_N
	case "Enum_CompassPointEnumeration_NNE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NNE
	case "Enum_CompassPointEnumeration_NE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NE
	case "Enum_CompassPointEnumeration_ENE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_ENE
	case "Enum_CompassPointEnumeration_E":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_E
	case "Enum_CompassPointEnumeration_ESE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_ESE
	case "Enum_CompassPointEnumeration_SE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SE
	case "Enum_CompassPointEnumeration_SSE":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SSE
	case "Enum_CompassPointEnumeration_S":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_S
	case "Enum_CompassPointEnumeration_SSW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SSW
	case "Enum_CompassPointEnumeration_SW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_SW
	case "Enum_CompassPointEnumeration_WSW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_WSW
	case "Enum_CompassPointEnumeration_W":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_W
	case "Enum_CompassPointEnumeration_WNW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_WNW
	case "Enum_CompassPointEnumeration_NW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NW
	case "Enum_CompassPointEnumeration_NNW":
		*enum_compasspointenumeration = Enum_CompassPointEnumeration_NNW
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_compasspointenumeration *Enum_CompassPointEnumeration) ToCodeString() (res string) {

	switch *enum_compasspointenumeration {
	// insertion code per enum code
	case Enum_CompassPointEnumeration_N:
		res = "Enum_CompassPointEnumeration_N"
	case Enum_CompassPointEnumeration_NNE:
		res = "Enum_CompassPointEnumeration_NNE"
	case Enum_CompassPointEnumeration_NE:
		res = "Enum_CompassPointEnumeration_NE"
	case Enum_CompassPointEnumeration_ENE:
		res = "Enum_CompassPointEnumeration_ENE"
	case Enum_CompassPointEnumeration_E:
		res = "Enum_CompassPointEnumeration_E"
	case Enum_CompassPointEnumeration_ESE:
		res = "Enum_CompassPointEnumeration_ESE"
	case Enum_CompassPointEnumeration_SE:
		res = "Enum_CompassPointEnumeration_SE"
	case Enum_CompassPointEnumeration_SSE:
		res = "Enum_CompassPointEnumeration_SSE"
	case Enum_CompassPointEnumeration_S:
		res = "Enum_CompassPointEnumeration_S"
	case Enum_CompassPointEnumeration_SSW:
		res = "Enum_CompassPointEnumeration_SSW"
	case Enum_CompassPointEnumeration_SW:
		res = "Enum_CompassPointEnumeration_SW"
	case Enum_CompassPointEnumeration_WSW:
		res = "Enum_CompassPointEnumeration_WSW"
	case Enum_CompassPointEnumeration_W:
		res = "Enum_CompassPointEnumeration_W"
	case Enum_CompassPointEnumeration_WNW:
		res = "Enum_CompassPointEnumeration_WNW"
	case Enum_CompassPointEnumeration_NW:
		res = "Enum_CompassPointEnumeration_NW"
	case Enum_CompassPointEnumeration_NNW:
		res = "Enum_CompassPointEnumeration_NNW"
	}
	return
}

func (enum_compasspointenumeration Enum_CompassPointEnumeration) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CompassPointEnumeration_N")
	res = append(res, "Enum_CompassPointEnumeration_NNE")
	res = append(res, "Enum_CompassPointEnumeration_NE")
	res = append(res, "Enum_CompassPointEnumeration_ENE")
	res = append(res, "Enum_CompassPointEnumeration_E")
	res = append(res, "Enum_CompassPointEnumeration_ESE")
	res = append(res, "Enum_CompassPointEnumeration_SE")
	res = append(res, "Enum_CompassPointEnumeration_SSE")
	res = append(res, "Enum_CompassPointEnumeration_S")
	res = append(res, "Enum_CompassPointEnumeration_SSW")
	res = append(res, "Enum_CompassPointEnumeration_SW")
	res = append(res, "Enum_CompassPointEnumeration_WSW")
	res = append(res, "Enum_CompassPointEnumeration_W")
	res = append(res, "Enum_CompassPointEnumeration_WNW")
	res = append(res, "Enum_CompassPointEnumeration_NW")
	res = append(res, "Enum_CompassPointEnumeration_NNW")

	return
}

func (enum_compasspointenumeration Enum_CompassPointEnumeration) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "N")
	res = append(res, "NNE")
	res = append(res, "NE")
	res = append(res, "ENE")
	res = append(res, "E")
	res = append(res, "ESE")
	res = append(res, "SE")
	res = append(res, "SSE")
	res = append(res, "S")
	res = append(res, "SSW")
	res = append(res, "SW")
	res = append(res, "WSW")
	res = append(res, "W")
	res = append(res, "WNW")
	res = append(res, "NW")
	res = append(res, "NNW")

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
	case Enum_CurveInterpolationType_CircularArc2PointWithBulge:
		res = "circularArc2PointWithBulge"
	case Enum_CurveInterpolationType_CircularArcCenterPointWithRadius:
		res = "circularArcCenterPointWithRadius"
	case Enum_CurveInterpolationType_Elliptical:
		res = "elliptical"
	case Enum_CurveInterpolationType_Clothoid:
		res = "clothoid"
	case Enum_CurveInterpolationType_Conic:
		res = "conic"
	case Enum_CurveInterpolationType_PolynomialSpline:
		res = "polynomialSpline"
	case Enum_CurveInterpolationType_CubicSpline:
		res = "cubicSpline"
	case Enum_CurveInterpolationType_RationalSpline:
		res = "rationalSpline"
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
	case "circularArc2PointWithBulge":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArc2PointWithBulge
		return
	case "circularArcCenterPointWithRadius":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArcCenterPointWithRadius
		return
	case "elliptical":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Elliptical
		return
	case "clothoid":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Clothoid
		return
	case "conic":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Conic
		return
	case "polynomialSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_PolynomialSpline
		return
	case "cubicSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CubicSpline
		return
	case "rationalSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_RationalSpline
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
	case "Enum_CurveInterpolationType_CircularArc2PointWithBulge":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArc2PointWithBulge
	case "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CircularArcCenterPointWithRadius
	case "Enum_CurveInterpolationType_Elliptical":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Elliptical
	case "Enum_CurveInterpolationType_Clothoid":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Clothoid
	case "Enum_CurveInterpolationType_Conic":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_Conic
	case "Enum_CurveInterpolationType_PolynomialSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_PolynomialSpline
	case "Enum_CurveInterpolationType_CubicSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_CubicSpline
	case "Enum_CurveInterpolationType_RationalSpline":
		*enum_curveinterpolationtype = Enum_CurveInterpolationType_RationalSpline
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
	case Enum_CurveInterpolationType_CircularArc2PointWithBulge:
		res = "Enum_CurveInterpolationType_CircularArc2PointWithBulge"
	case Enum_CurveInterpolationType_CircularArcCenterPointWithRadius:
		res = "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius"
	case Enum_CurveInterpolationType_Elliptical:
		res = "Enum_CurveInterpolationType_Elliptical"
	case Enum_CurveInterpolationType_Clothoid:
		res = "Enum_CurveInterpolationType_Clothoid"
	case Enum_CurveInterpolationType_Conic:
		res = "Enum_CurveInterpolationType_Conic"
	case Enum_CurveInterpolationType_PolynomialSpline:
		res = "Enum_CurveInterpolationType_PolynomialSpline"
	case Enum_CurveInterpolationType_CubicSpline:
		res = "Enum_CurveInterpolationType_CubicSpline"
	case Enum_CurveInterpolationType_RationalSpline:
		res = "Enum_CurveInterpolationType_RationalSpline"
	}
	return
}

func (enum_curveinterpolationtype Enum_CurveInterpolationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_CurveInterpolationType_Linear")
	res = append(res, "Enum_CurveInterpolationType_Geodesic")
	res = append(res, "Enum_CurveInterpolationType_CircularArc3Points")
	res = append(res, "Enum_CurveInterpolationType_CircularArc2PointWithBulge")
	res = append(res, "Enum_CurveInterpolationType_CircularArcCenterPointWithRadius")
	res = append(res, "Enum_CurveInterpolationType_Elliptical")
	res = append(res, "Enum_CurveInterpolationType_Clothoid")
	res = append(res, "Enum_CurveInterpolationType_Conic")
	res = append(res, "Enum_CurveInterpolationType_PolynomialSpline")
	res = append(res, "Enum_CurveInterpolationType_CubicSpline")
	res = append(res, "Enum_CurveInterpolationType_RationalSpline")

	return
}

func (enum_curveinterpolationtype Enum_CurveInterpolationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "linear")
	res = append(res, "geodesic")
	res = append(res, "circularArc3Points")
	res = append(res, "circularArc2PointWithBulge")
	res = append(res, "circularArcCenterPointWithRadius")
	res = append(res, "elliptical")
	res = append(res, "clothoid")
	res = append(res, "conic")
	res = append(res, "polynomialSpline")
	res = append(res, "cubicSpline")
	res = append(res, "rationalSpline")

	return
}

// Utility function for Enum_DegreesType_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_degreestype_direction Enum_DegreesType_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_degreestype_direction {
	// insertion code per enum code
	case Enum_DegreesType_direction_N:
		res = "N"
	case Enum_DegreesType_direction_E:
		res = "E"
	case Enum_DegreesType_direction_S:
		res = "S"
	case Enum_DegreesType_direction_W:
		res = "W"
	case Enum_DegreesType_direction_N__:
		res = "+"
	case Enum_DegreesType_direction_N___1:
		res = "-"
	}
	return
}

func (enum_degreestype_direction *Enum_DegreesType_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "N":
		*enum_degreestype_direction = Enum_DegreesType_direction_N
		return
	case "E":
		*enum_degreestype_direction = Enum_DegreesType_direction_E
		return
	case "S":
		*enum_degreestype_direction = Enum_DegreesType_direction_S
		return
	case "W":
		*enum_degreestype_direction = Enum_DegreesType_direction_W
		return
	case "+":
		*enum_degreestype_direction = Enum_DegreesType_direction_N__
		return
	case "-":
		*enum_degreestype_direction = Enum_DegreesType_direction_N___1
		return
	default:
		return errUnkownEnum
	}
}

func (enum_degreestype_direction *Enum_DegreesType_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_DegreesType_direction_N":
		*enum_degreestype_direction = Enum_DegreesType_direction_N
	case "Enum_DegreesType_direction_E":
		*enum_degreestype_direction = Enum_DegreesType_direction_E
	case "Enum_DegreesType_direction_S":
		*enum_degreestype_direction = Enum_DegreesType_direction_S
	case "Enum_DegreesType_direction_W":
		*enum_degreestype_direction = Enum_DegreesType_direction_W
	case "Enum_DegreesType_direction_N__":
		*enum_degreestype_direction = Enum_DegreesType_direction_N__
	case "Enum_DegreesType_direction_N___1":
		*enum_degreestype_direction = Enum_DegreesType_direction_N___1
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_degreestype_direction *Enum_DegreesType_direction) ToCodeString() (res string) {

	switch *enum_degreestype_direction {
	// insertion code per enum code
	case Enum_DegreesType_direction_N:
		res = "Enum_DegreesType_direction_N"
	case Enum_DegreesType_direction_E:
		res = "Enum_DegreesType_direction_E"
	case Enum_DegreesType_direction_S:
		res = "Enum_DegreesType_direction_S"
	case Enum_DegreesType_direction_W:
		res = "Enum_DegreesType_direction_W"
	case Enum_DegreesType_direction_N__:
		res = "Enum_DegreesType_direction_N__"
	case Enum_DegreesType_direction_N___1:
		res = "Enum_DegreesType_direction_N___1"
	}
	return
}

func (enum_degreestype_direction Enum_DegreesType_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_DegreesType_direction_N")
	res = append(res, "Enum_DegreesType_direction_E")
	res = append(res, "Enum_DegreesType_direction_S")
	res = append(res, "Enum_DegreesType_direction_W")
	res = append(res, "Enum_DegreesType_direction_N__")
	res = append(res, "Enum_DegreesType_direction_N___1")

	return
}

func (enum_degreestype_direction Enum_DegreesType_direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "N")
	res = append(res, "E")
	res = append(res, "S")
	res = append(res, "W")
	res = append(res, "+")
	res = append(res, "-")

	return
}

// Utility function for Enum_IncrementOrder
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_incrementorder Enum_IncrementOrder) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_incrementorder {
	// insertion code per enum code
	case Enum_IncrementOrder_N__x_y:
		res = "+x+y"
	case Enum_IncrementOrder_N__y_x:
		res = "+y+x"
	case Enum_IncrementOrder_N__x_y_1:
		res = "+x-y"
	case Enum_IncrementOrder_N__x_y_2:
		res = "-x-y"
	}
	return
}

func (enum_incrementorder *Enum_IncrementOrder) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "+x+y":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y
		return
	case "+y+x":
		*enum_incrementorder = Enum_IncrementOrder_N__y_x
		return
	case "+x-y":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y_1
		return
	case "-x-y":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y_2
		return
	default:
		return errUnkownEnum
	}
}

func (enum_incrementorder *Enum_IncrementOrder) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_IncrementOrder_N__x_y":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y
	case "Enum_IncrementOrder_N__y_x":
		*enum_incrementorder = Enum_IncrementOrder_N__y_x
	case "Enum_IncrementOrder_N__x_y_1":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y_1
	case "Enum_IncrementOrder_N__x_y_2":
		*enum_incrementorder = Enum_IncrementOrder_N__x_y_2
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_incrementorder *Enum_IncrementOrder) ToCodeString() (res string) {

	switch *enum_incrementorder {
	// insertion code per enum code
	case Enum_IncrementOrder_N__x_y:
		res = "Enum_IncrementOrder_N__x_y"
	case Enum_IncrementOrder_N__y_x:
		res = "Enum_IncrementOrder_N__y_x"
	case Enum_IncrementOrder_N__x_y_1:
		res = "Enum_IncrementOrder_N__x_y_1"
	case Enum_IncrementOrder_N__x_y_2:
		res = "Enum_IncrementOrder_N__x_y_2"
	}
	return
}

func (enum_incrementorder Enum_IncrementOrder) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_IncrementOrder_N__x_y")
	res = append(res, "Enum_IncrementOrder_N__y_x")
	res = append(res, "Enum_IncrementOrder_N__x_y_1")
	res = append(res, "Enum_IncrementOrder_N__x_y_2")

	return
}

func (enum_incrementorder Enum_IncrementOrder) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "+x+y")
	res = append(res, "+y+x")
	res = append(res, "+x-y")
	res = append(res, "-x-y")

	return
}

// Utility function for Enum_KnotTypesType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_knottypestype Enum_KnotTypesType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_knottypestype {
	// insertion code per enum code
	case Enum_KnotTypesType_Uniform:
		res = "uniform"
	case Enum_KnotTypesType_QuasiUniform:
		res = "quasiUniform"
	case Enum_KnotTypesType_PiecewiseBezier:
		res = "piecewiseBezier"
	}
	return
}

func (enum_knottypestype *Enum_KnotTypesType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "uniform":
		*enum_knottypestype = Enum_KnotTypesType_Uniform
		return
	case "quasiUniform":
		*enum_knottypestype = Enum_KnotTypesType_QuasiUniform
		return
	case "piecewiseBezier":
		*enum_knottypestype = Enum_KnotTypesType_PiecewiseBezier
		return
	default:
		return errUnkownEnum
	}
}

func (enum_knottypestype *Enum_KnotTypesType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_KnotTypesType_Uniform":
		*enum_knottypestype = Enum_KnotTypesType_Uniform
	case "Enum_KnotTypesType_QuasiUniform":
		*enum_knottypestype = Enum_KnotTypesType_QuasiUniform
	case "Enum_KnotTypesType_PiecewiseBezier":
		*enum_knottypestype = Enum_KnotTypesType_PiecewiseBezier
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_knottypestype *Enum_KnotTypesType) ToCodeString() (res string) {

	switch *enum_knottypestype {
	// insertion code per enum code
	case Enum_KnotTypesType_Uniform:
		res = "Enum_KnotTypesType_Uniform"
	case Enum_KnotTypesType_QuasiUniform:
		res = "Enum_KnotTypesType_QuasiUniform"
	case Enum_KnotTypesType_PiecewiseBezier:
		res = "Enum_KnotTypesType_PiecewiseBezier"
	}
	return
}

func (enum_knottypestype Enum_KnotTypesType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_KnotTypesType_Uniform")
	res = append(res, "Enum_KnotTypesType_QuasiUniform")
	res = append(res, "Enum_KnotTypesType_PiecewiseBezier")

	return
}

func (enum_knottypestype Enum_KnotTypesType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "uniform")
	res = append(res, "quasiUniform")
	res = append(res, "piecewiseBezier")

	return
}

// Utility function for Enum_MD_TopicCategoryCode_Type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_md_topiccategorycode_type Enum_MD_TopicCategoryCode_Type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_md_topiccategorycode_type {
	// insertion code per enum code
	case Enum_MD_TopicCategoryCode_Type_Farming:
		res = "farming"
	case Enum_MD_TopicCategoryCode_Type_Biota:
		res = "biota"
	case Enum_MD_TopicCategoryCode_Type_Boundaries:
		res = "boundaries"
	case Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere:
		res = "climatologyMeteorologyAtmosphere"
	case Enum_MD_TopicCategoryCode_Type_Economy:
		res = "economy"
	case Enum_MD_TopicCategoryCode_Type_Elevation:
		res = "elevation"
	case Enum_MD_TopicCategoryCode_Type_Environment:
		res = "environment"
	case Enum_MD_TopicCategoryCode_Type_GeoscientificInformation:
		res = "geoscientificInformation"
	case Enum_MD_TopicCategoryCode_Type_Health:
		res = "health"
	case Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover:
		res = "imageryBaseMapsEarthCover"
	case Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary:
		res = "intelligenceMilitary"
	case Enum_MD_TopicCategoryCode_Type_InlandWaters:
		res = "inlandWaters"
	case Enum_MD_TopicCategoryCode_Type_Location:
		res = "location"
	case Enum_MD_TopicCategoryCode_Type_Oceans:
		res = "oceans"
	case Enum_MD_TopicCategoryCode_Type_PlanningCadastre:
		res = "planningCadastre"
	case Enum_MD_TopicCategoryCode_Type_Society:
		res = "society"
	case Enum_MD_TopicCategoryCode_Type_Structure:
		res = "structure"
	case Enum_MD_TopicCategoryCode_Type_Transportation:
		res = "transportation"
	case Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication:
		res = "utilitiesCommunication"
	}
	return
}

func (enum_md_topiccategorycode_type *Enum_MD_TopicCategoryCode_Type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "farming":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Farming
		return
	case "biota":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Biota
		return
	case "boundaries":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Boundaries
		return
	case "climatologyMeteorologyAtmosphere":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere
		return
	case "economy":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Economy
		return
	case "elevation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Elevation
		return
	case "environment":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Environment
		return
	case "geoscientificInformation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_GeoscientificInformation
		return
	case "health":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Health
		return
	case "imageryBaseMapsEarthCover":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover
		return
	case "intelligenceMilitary":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary
		return
	case "inlandWaters":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_InlandWaters
		return
	case "location":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Location
		return
	case "oceans":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Oceans
		return
	case "planningCadastre":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_PlanningCadastre
		return
	case "society":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Society
		return
	case "structure":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Structure
		return
	case "transportation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Transportation
		return
	case "utilitiesCommunication":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication
		return
	default:
		return errUnkownEnum
	}
}

func (enum_md_topiccategorycode_type *Enum_MD_TopicCategoryCode_Type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_MD_TopicCategoryCode_Type_Farming":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Farming
	case "Enum_MD_TopicCategoryCode_Type_Biota":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Biota
	case "Enum_MD_TopicCategoryCode_Type_Boundaries":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Boundaries
	case "Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere
	case "Enum_MD_TopicCategoryCode_Type_Economy":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Economy
	case "Enum_MD_TopicCategoryCode_Type_Elevation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Elevation
	case "Enum_MD_TopicCategoryCode_Type_Environment":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Environment
	case "Enum_MD_TopicCategoryCode_Type_GeoscientificInformation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_GeoscientificInformation
	case "Enum_MD_TopicCategoryCode_Type_Health":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Health
	case "Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover
	case "Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary
	case "Enum_MD_TopicCategoryCode_Type_InlandWaters":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_InlandWaters
	case "Enum_MD_TopicCategoryCode_Type_Location":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Location
	case "Enum_MD_TopicCategoryCode_Type_Oceans":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Oceans
	case "Enum_MD_TopicCategoryCode_Type_PlanningCadastre":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_PlanningCadastre
	case "Enum_MD_TopicCategoryCode_Type_Society":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Society
	case "Enum_MD_TopicCategoryCode_Type_Structure":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Structure
	case "Enum_MD_TopicCategoryCode_Type_Transportation":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_Transportation
	case "Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication":
		*enum_md_topiccategorycode_type = Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_md_topiccategorycode_type *Enum_MD_TopicCategoryCode_Type) ToCodeString() (res string) {

	switch *enum_md_topiccategorycode_type {
	// insertion code per enum code
	case Enum_MD_TopicCategoryCode_Type_Farming:
		res = "Enum_MD_TopicCategoryCode_Type_Farming"
	case Enum_MD_TopicCategoryCode_Type_Biota:
		res = "Enum_MD_TopicCategoryCode_Type_Biota"
	case Enum_MD_TopicCategoryCode_Type_Boundaries:
		res = "Enum_MD_TopicCategoryCode_Type_Boundaries"
	case Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere:
		res = "Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere"
	case Enum_MD_TopicCategoryCode_Type_Economy:
		res = "Enum_MD_TopicCategoryCode_Type_Economy"
	case Enum_MD_TopicCategoryCode_Type_Elevation:
		res = "Enum_MD_TopicCategoryCode_Type_Elevation"
	case Enum_MD_TopicCategoryCode_Type_Environment:
		res = "Enum_MD_TopicCategoryCode_Type_Environment"
	case Enum_MD_TopicCategoryCode_Type_GeoscientificInformation:
		res = "Enum_MD_TopicCategoryCode_Type_GeoscientificInformation"
	case Enum_MD_TopicCategoryCode_Type_Health:
		res = "Enum_MD_TopicCategoryCode_Type_Health"
	case Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover:
		res = "Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover"
	case Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary:
		res = "Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary"
	case Enum_MD_TopicCategoryCode_Type_InlandWaters:
		res = "Enum_MD_TopicCategoryCode_Type_InlandWaters"
	case Enum_MD_TopicCategoryCode_Type_Location:
		res = "Enum_MD_TopicCategoryCode_Type_Location"
	case Enum_MD_TopicCategoryCode_Type_Oceans:
		res = "Enum_MD_TopicCategoryCode_Type_Oceans"
	case Enum_MD_TopicCategoryCode_Type_PlanningCadastre:
		res = "Enum_MD_TopicCategoryCode_Type_PlanningCadastre"
	case Enum_MD_TopicCategoryCode_Type_Society:
		res = "Enum_MD_TopicCategoryCode_Type_Society"
	case Enum_MD_TopicCategoryCode_Type_Structure:
		res = "Enum_MD_TopicCategoryCode_Type_Structure"
	case Enum_MD_TopicCategoryCode_Type_Transportation:
		res = "Enum_MD_TopicCategoryCode_Type_Transportation"
	case Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication:
		res = "Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication"
	}
	return
}

func (enum_md_topiccategorycode_type Enum_MD_TopicCategoryCode_Type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Farming")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Biota")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Boundaries")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Economy")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Elevation")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Environment")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_GeoscientificInformation")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Health")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_InlandWaters")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Location")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Oceans")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_PlanningCadastre")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Society")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Structure")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_Transportation")
	res = append(res, "Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication")

	return
}

func (enum_md_topiccategorycode_type Enum_MD_TopicCategoryCode_Type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "farming")
	res = append(res, "biota")
	res = append(res, "boundaries")
	res = append(res, "climatologyMeteorologyAtmosphere")
	res = append(res, "economy")
	res = append(res, "elevation")
	res = append(res, "environment")
	res = append(res, "geoscientificInformation")
	res = append(res, "health")
	res = append(res, "imageryBaseMapsEarthCover")
	res = append(res, "intelligenceMilitary")
	res = append(res, "inlandWaters")
	res = append(res, "location")
	res = append(res, "oceans")
	res = append(res, "planningCadastre")
	res = append(res, "society")
	res = append(res, "structure")
	res = append(res, "transportation")
	res = append(res, "utilitiesCommunication")

	return
}

// Utility function for Enum_RelatedTimeType_relativePosition
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_relatedtimetype_relativeposition Enum_RelatedTimeType_relativePosition) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_relatedtimetype_relativeposition {
	// insertion code per enum code
	case Enum_RelatedTimeType_relativePosition_Before:
		res = "Before"
	case Enum_RelatedTimeType_relativePosition_After:
		res = "After"
	case Enum_RelatedTimeType_relativePosition_Begins:
		res = "Begins"
	case Enum_RelatedTimeType_relativePosition_Ends:
		res = "Ends"
	case Enum_RelatedTimeType_relativePosition_During:
		res = "During"
	case Enum_RelatedTimeType_relativePosition_Equals:
		res = "Equals"
	case Enum_RelatedTimeType_relativePosition_Contains:
		res = "Contains"
	case Enum_RelatedTimeType_relativePosition_Overlaps:
		res = "Overlaps"
	case Enum_RelatedTimeType_relativePosition_Meets:
		res = "Meets"
	case Enum_RelatedTimeType_relativePosition_OverlappedBy:
		res = "OverlappedBy"
	case Enum_RelatedTimeType_relativePosition_MetBy:
		res = "MetBy"
	case Enum_RelatedTimeType_relativePosition_BegunBy:
		res = "BegunBy"
	case Enum_RelatedTimeType_relativePosition_EndedBy:
		res = "EndedBy"
	}
	return
}

func (enum_relatedtimetype_relativeposition *Enum_RelatedTimeType_relativePosition) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Before":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Before
		return
	case "After":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_After
		return
	case "Begins":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Begins
		return
	case "Ends":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Ends
		return
	case "During":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_During
		return
	case "Equals":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Equals
		return
	case "Contains":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Contains
		return
	case "Overlaps":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Overlaps
		return
	case "Meets":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Meets
		return
	case "OverlappedBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_OverlappedBy
		return
	case "MetBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_MetBy
		return
	case "BegunBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_BegunBy
		return
	case "EndedBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_EndedBy
		return
	default:
		return errUnkownEnum
	}
}

func (enum_relatedtimetype_relativeposition *Enum_RelatedTimeType_relativePosition) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_RelatedTimeType_relativePosition_Before":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Before
	case "Enum_RelatedTimeType_relativePosition_After":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_After
	case "Enum_RelatedTimeType_relativePosition_Begins":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Begins
	case "Enum_RelatedTimeType_relativePosition_Ends":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Ends
	case "Enum_RelatedTimeType_relativePosition_During":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_During
	case "Enum_RelatedTimeType_relativePosition_Equals":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Equals
	case "Enum_RelatedTimeType_relativePosition_Contains":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Contains
	case "Enum_RelatedTimeType_relativePosition_Overlaps":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Overlaps
	case "Enum_RelatedTimeType_relativePosition_Meets":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_Meets
	case "Enum_RelatedTimeType_relativePosition_OverlappedBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_OverlappedBy
	case "Enum_RelatedTimeType_relativePosition_MetBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_MetBy
	case "Enum_RelatedTimeType_relativePosition_BegunBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_BegunBy
	case "Enum_RelatedTimeType_relativePosition_EndedBy":
		*enum_relatedtimetype_relativeposition = Enum_RelatedTimeType_relativePosition_EndedBy
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_relatedtimetype_relativeposition *Enum_RelatedTimeType_relativePosition) ToCodeString() (res string) {

	switch *enum_relatedtimetype_relativeposition {
	// insertion code per enum code
	case Enum_RelatedTimeType_relativePosition_Before:
		res = "Enum_RelatedTimeType_relativePosition_Before"
	case Enum_RelatedTimeType_relativePosition_After:
		res = "Enum_RelatedTimeType_relativePosition_After"
	case Enum_RelatedTimeType_relativePosition_Begins:
		res = "Enum_RelatedTimeType_relativePosition_Begins"
	case Enum_RelatedTimeType_relativePosition_Ends:
		res = "Enum_RelatedTimeType_relativePosition_Ends"
	case Enum_RelatedTimeType_relativePosition_During:
		res = "Enum_RelatedTimeType_relativePosition_During"
	case Enum_RelatedTimeType_relativePosition_Equals:
		res = "Enum_RelatedTimeType_relativePosition_Equals"
	case Enum_RelatedTimeType_relativePosition_Contains:
		res = "Enum_RelatedTimeType_relativePosition_Contains"
	case Enum_RelatedTimeType_relativePosition_Overlaps:
		res = "Enum_RelatedTimeType_relativePosition_Overlaps"
	case Enum_RelatedTimeType_relativePosition_Meets:
		res = "Enum_RelatedTimeType_relativePosition_Meets"
	case Enum_RelatedTimeType_relativePosition_OverlappedBy:
		res = "Enum_RelatedTimeType_relativePosition_OverlappedBy"
	case Enum_RelatedTimeType_relativePosition_MetBy:
		res = "Enum_RelatedTimeType_relativePosition_MetBy"
	case Enum_RelatedTimeType_relativePosition_BegunBy:
		res = "Enum_RelatedTimeType_relativePosition_BegunBy"
	case Enum_RelatedTimeType_relativePosition_EndedBy:
		res = "Enum_RelatedTimeType_relativePosition_EndedBy"
	}
	return
}

func (enum_relatedtimetype_relativeposition Enum_RelatedTimeType_relativePosition) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_RelatedTimeType_relativePosition_Before")
	res = append(res, "Enum_RelatedTimeType_relativePosition_After")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Begins")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Ends")
	res = append(res, "Enum_RelatedTimeType_relativePosition_During")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Equals")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Contains")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Overlaps")
	res = append(res, "Enum_RelatedTimeType_relativePosition_Meets")
	res = append(res, "Enum_RelatedTimeType_relativePosition_OverlappedBy")
	res = append(res, "Enum_RelatedTimeType_relativePosition_MetBy")
	res = append(res, "Enum_RelatedTimeType_relativePosition_BegunBy")
	res = append(res, "Enum_RelatedTimeType_relativePosition_EndedBy")

	return
}

func (enum_relatedtimetype_relativeposition Enum_RelatedTimeType_relativePosition) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Before")
	res = append(res, "After")
	res = append(res, "Begins")
	res = append(res, "Ends")
	res = append(res, "During")
	res = append(res, "Equals")
	res = append(res, "Contains")
	res = append(res, "Overlaps")
	res = append(res, "Meets")
	res = append(res, "OverlappedBy")
	res = append(res, "MetBy")
	res = append(res, "BegunBy")
	res = append(res, "EndedBy")

	return
}

// Utility function for Enum_SequenceRuleEnumeration
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_sequenceruleenumeration Enum_SequenceRuleEnumeration) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_sequenceruleenumeration {
	// insertion code per enum code
	case Enum_SequenceRuleEnumeration_Linear:
		res = "Linear"
	case Enum_SequenceRuleEnumeration_Boustrophedonic:
		res = "Boustrophedonic"
	case Enum_SequenceRuleEnumeration_Cantor_diagonal:
		res = "Cantor-diagonal"
	case Enum_SequenceRuleEnumeration_Spiral:
		res = "Spiral"
	case Enum_SequenceRuleEnumeration_Morton:
		res = "Morton"
	case Enum_SequenceRuleEnumeration_Hilbert:
		res = "Hilbert"
	}
	return
}

func (enum_sequenceruleenumeration *Enum_SequenceRuleEnumeration) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Linear":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Linear
		return
	case "Boustrophedonic":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Boustrophedonic
		return
	case "Cantor-diagonal":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Cantor_diagonal
		return
	case "Spiral":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Spiral
		return
	case "Morton":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Morton
		return
	case "Hilbert":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Hilbert
		return
	default:
		return errUnkownEnum
	}
}

func (enum_sequenceruleenumeration *Enum_SequenceRuleEnumeration) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SequenceRuleEnumeration_Linear":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Linear
	case "Enum_SequenceRuleEnumeration_Boustrophedonic":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Boustrophedonic
	case "Enum_SequenceRuleEnumeration_Cantor_diagonal":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Cantor_diagonal
	case "Enum_SequenceRuleEnumeration_Spiral":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Spiral
	case "Enum_SequenceRuleEnumeration_Morton":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Morton
	case "Enum_SequenceRuleEnumeration_Hilbert":
		*enum_sequenceruleenumeration = Enum_SequenceRuleEnumeration_Hilbert
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_sequenceruleenumeration *Enum_SequenceRuleEnumeration) ToCodeString() (res string) {

	switch *enum_sequenceruleenumeration {
	// insertion code per enum code
	case Enum_SequenceRuleEnumeration_Linear:
		res = "Enum_SequenceRuleEnumeration_Linear"
	case Enum_SequenceRuleEnumeration_Boustrophedonic:
		res = "Enum_SequenceRuleEnumeration_Boustrophedonic"
	case Enum_SequenceRuleEnumeration_Cantor_diagonal:
		res = "Enum_SequenceRuleEnumeration_Cantor_diagonal"
	case Enum_SequenceRuleEnumeration_Spiral:
		res = "Enum_SequenceRuleEnumeration_Spiral"
	case Enum_SequenceRuleEnumeration_Morton:
		res = "Enum_SequenceRuleEnumeration_Morton"
	case Enum_SequenceRuleEnumeration_Hilbert:
		res = "Enum_SequenceRuleEnumeration_Hilbert"
	}
	return
}

func (enum_sequenceruleenumeration Enum_SequenceRuleEnumeration) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SequenceRuleEnumeration_Linear")
	res = append(res, "Enum_SequenceRuleEnumeration_Boustrophedonic")
	res = append(res, "Enum_SequenceRuleEnumeration_Cantor_diagonal")
	res = append(res, "Enum_SequenceRuleEnumeration_Spiral")
	res = append(res, "Enum_SequenceRuleEnumeration_Morton")
	res = append(res, "Enum_SequenceRuleEnumeration_Hilbert")

	return
}

func (enum_sequenceruleenumeration Enum_SequenceRuleEnumeration) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Linear")
	res = append(res, "Boustrophedonic")
	res = append(res, "Cantor-diagonal")
	res = append(res, "Spiral")
	res = append(res, "Morton")
	res = append(res, "Hilbert")

	return
}

// Utility function for Enum_SuccessionType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_successiontype Enum_SuccessionType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_successiontype {
	// insertion code per enum code
	case Enum_SuccessionType_Substitution:
		res = "substitution"
	case Enum_SuccessionType_Division:
		res = "division"
	case Enum_SuccessionType_Fusion:
		res = "fusion"
	case Enum_SuccessionType_Initiation:
		res = "initiation"
	}
	return
}

func (enum_successiontype *Enum_SuccessionType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "substitution":
		*enum_successiontype = Enum_SuccessionType_Substitution
		return
	case "division":
		*enum_successiontype = Enum_SuccessionType_Division
		return
	case "fusion":
		*enum_successiontype = Enum_SuccessionType_Fusion
		return
	case "initiation":
		*enum_successiontype = Enum_SuccessionType_Initiation
		return
	default:
		return errUnkownEnum
	}
}

func (enum_successiontype *Enum_SuccessionType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SuccessionType_Substitution":
		*enum_successiontype = Enum_SuccessionType_Substitution
	case "Enum_SuccessionType_Division":
		*enum_successiontype = Enum_SuccessionType_Division
	case "Enum_SuccessionType_Fusion":
		*enum_successiontype = Enum_SuccessionType_Fusion
	case "Enum_SuccessionType_Initiation":
		*enum_successiontype = Enum_SuccessionType_Initiation
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_successiontype *Enum_SuccessionType) ToCodeString() (res string) {

	switch *enum_successiontype {
	// insertion code per enum code
	case Enum_SuccessionType_Substitution:
		res = "Enum_SuccessionType_Substitution"
	case Enum_SuccessionType_Division:
		res = "Enum_SuccessionType_Division"
	case Enum_SuccessionType_Fusion:
		res = "Enum_SuccessionType_Fusion"
	case Enum_SuccessionType_Initiation:
		res = "Enum_SuccessionType_Initiation"
	}
	return
}

func (enum_successiontype Enum_SuccessionType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SuccessionType_Substitution")
	res = append(res, "Enum_SuccessionType_Division")
	res = append(res, "Enum_SuccessionType_Fusion")
	res = append(res, "Enum_SuccessionType_Initiation")

	return
}

func (enum_successiontype Enum_SuccessionType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "substitution")
	res = append(res, "division")
	res = append(res, "fusion")
	res = append(res, "initiation")

	return
}

// Utility function for Enum_SurfaceInterpolationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_surfaceinterpolationtype {
	// insertion code per enum code
	case Enum_SurfaceInterpolationType_None:
		res = "none"
	case Enum_SurfaceInterpolationType_Planar:
		res = "planar"
	case Enum_SurfaceInterpolationType_Spherical:
		res = "spherical"
	case Enum_SurfaceInterpolationType_Elliptical:
		res = "elliptical"
	case Enum_SurfaceInterpolationType_Conic:
		res = "conic"
	case Enum_SurfaceInterpolationType_Tin:
		res = "tin"
	case Enum_SurfaceInterpolationType_ParametricCurve:
		res = "parametricCurve"
	case Enum_SurfaceInterpolationType_PolynomialSpline:
		res = "polynomialSpline"
	case Enum_SurfaceInterpolationType_RationalSpline:
		res = "rationalSpline"
	case Enum_SurfaceInterpolationType_TriangulatedSpline:
		res = "triangulatedSpline"
	}
	return
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_None
		return
	case "planar":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Planar
		return
	case "spherical":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Spherical
		return
	case "elliptical":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Elliptical
		return
	case "conic":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Conic
		return
	case "tin":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Tin
		return
	case "parametricCurve":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_ParametricCurve
		return
	case "polynomialSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_PolynomialSpline
		return
	case "rationalSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_RationalSpline
		return
	case "triangulatedSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_TriangulatedSpline
		return
	default:
		return errUnkownEnum
	}
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_SurfaceInterpolationType_None":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_None
	case "Enum_SurfaceInterpolationType_Planar":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Planar
	case "Enum_SurfaceInterpolationType_Spherical":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Spherical
	case "Enum_SurfaceInterpolationType_Elliptical":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Elliptical
	case "Enum_SurfaceInterpolationType_Conic":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Conic
	case "Enum_SurfaceInterpolationType_Tin":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_Tin
	case "Enum_SurfaceInterpolationType_ParametricCurve":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_ParametricCurve
	case "Enum_SurfaceInterpolationType_PolynomialSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_PolynomialSpline
	case "Enum_SurfaceInterpolationType_RationalSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_RationalSpline
	case "Enum_SurfaceInterpolationType_TriangulatedSpline":
		*enum_surfaceinterpolationtype = Enum_SurfaceInterpolationType_TriangulatedSpline
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_surfaceinterpolationtype *Enum_SurfaceInterpolationType) ToCodeString() (res string) {

	switch *enum_surfaceinterpolationtype {
	// insertion code per enum code
	case Enum_SurfaceInterpolationType_None:
		res = "Enum_SurfaceInterpolationType_None"
	case Enum_SurfaceInterpolationType_Planar:
		res = "Enum_SurfaceInterpolationType_Planar"
	case Enum_SurfaceInterpolationType_Spherical:
		res = "Enum_SurfaceInterpolationType_Spherical"
	case Enum_SurfaceInterpolationType_Elliptical:
		res = "Enum_SurfaceInterpolationType_Elliptical"
	case Enum_SurfaceInterpolationType_Conic:
		res = "Enum_SurfaceInterpolationType_Conic"
	case Enum_SurfaceInterpolationType_Tin:
		res = "Enum_SurfaceInterpolationType_Tin"
	case Enum_SurfaceInterpolationType_ParametricCurve:
		res = "Enum_SurfaceInterpolationType_ParametricCurve"
	case Enum_SurfaceInterpolationType_PolynomialSpline:
		res = "Enum_SurfaceInterpolationType_PolynomialSpline"
	case Enum_SurfaceInterpolationType_RationalSpline:
		res = "Enum_SurfaceInterpolationType_RationalSpline"
	case Enum_SurfaceInterpolationType_TriangulatedSpline:
		res = "Enum_SurfaceInterpolationType_TriangulatedSpline"
	}
	return
}

func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_SurfaceInterpolationType_None")
	res = append(res, "Enum_SurfaceInterpolationType_Planar")
	res = append(res, "Enum_SurfaceInterpolationType_Spherical")
	res = append(res, "Enum_SurfaceInterpolationType_Elliptical")
	res = append(res, "Enum_SurfaceInterpolationType_Conic")
	res = append(res, "Enum_SurfaceInterpolationType_Tin")
	res = append(res, "Enum_SurfaceInterpolationType_ParametricCurve")
	res = append(res, "Enum_SurfaceInterpolationType_PolynomialSpline")
	res = append(res, "Enum_SurfaceInterpolationType_RationalSpline")
	res = append(res, "Enum_SurfaceInterpolationType_TriangulatedSpline")

	return
}

func (enum_surfaceinterpolationtype Enum_SurfaceInterpolationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "planar")
	res = append(res, "spherical")
	res = append(res, "elliptical")
	res = append(res, "conic")
	res = append(res, "tin")
	res = append(res, "parametricCurve")
	res = append(res, "polynomialSpline")
	res = append(res, "rationalSpline")
	res = append(res, "triangulatedSpline")

	return
}

// Utility function for Enum_TimeIndeterminateValueType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_timeindeterminatevaluetype {
	// insertion code per enum code
	case Enum_TimeIndeterminateValueType_After:
		res = "after"
	case Enum_TimeIndeterminateValueType_Before:
		res = "before"
	case Enum_TimeIndeterminateValueType_Now:
		res = "now"
	case Enum_TimeIndeterminateValueType_Unknown:
		res = "unknown"
	}
	return
}

func (enum_timeindeterminatevaluetype *Enum_TimeIndeterminateValueType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "after":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_After
		return
	case "before":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Before
		return
	case "now":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Now
		return
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
	case "Enum_TimeIndeterminateValueType_After":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_After
	case "Enum_TimeIndeterminateValueType_Before":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Before
	case "Enum_TimeIndeterminateValueType_Now":
		*enum_timeindeterminatevaluetype = Enum_TimeIndeterminateValueType_Now
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
	case Enum_TimeIndeterminateValueType_After:
		res = "Enum_TimeIndeterminateValueType_After"
	case Enum_TimeIndeterminateValueType_Before:
		res = "Enum_TimeIndeterminateValueType_Before"
	case Enum_TimeIndeterminateValueType_Now:
		res = "Enum_TimeIndeterminateValueType_Now"
	case Enum_TimeIndeterminateValueType_Unknown:
		res = "Enum_TimeIndeterminateValueType_Unknown"
	}
	return
}

func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_TimeIndeterminateValueType_After")
	res = append(res, "Enum_TimeIndeterminateValueType_Before")
	res = append(res, "Enum_TimeIndeterminateValueType_Now")
	res = append(res, "Enum_TimeIndeterminateValueType_Unknown")

	return
}

func (enum_timeindeterminatevaluetype Enum_TimeIndeterminateValueType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "after")
	res = append(res, "before")
	res = append(res, "now")
	res = append(res, "unknown")

	return
}

// Utility function for Enum_TimeUnitType_Member_1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_timeunittype_member_1 Enum_TimeUnitType_Member_1) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_timeunittype_member_1 {
	// insertion code per enum code
	case Enum_TimeUnitType_Member_1_Year:
		res = "year"
	case Enum_TimeUnitType_Member_1_Month:
		res = "month"
	case Enum_TimeUnitType_Member_1_Day:
		res = "day"
	case Enum_TimeUnitType_Member_1_Hour:
		res = "hour"
	case Enum_TimeUnitType_Member_1_Minute:
		res = "minute"
	case Enum_TimeUnitType_Member_1_Second:
		res = "second"
	}
	return
}

func (enum_timeunittype_member_1 *Enum_TimeUnitType_Member_1) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "year":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Year
		return
	case "month":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Month
		return
	case "day":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Day
		return
	case "hour":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Hour
		return
	case "minute":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Minute
		return
	case "second":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Second
		return
	default:
		return errUnkownEnum
	}
}

func (enum_timeunittype_member_1 *Enum_TimeUnitType_Member_1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_TimeUnitType_Member_1_Year":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Year
	case "Enum_TimeUnitType_Member_1_Month":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Month
	case "Enum_TimeUnitType_Member_1_Day":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Day
	case "Enum_TimeUnitType_Member_1_Hour":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Hour
	case "Enum_TimeUnitType_Member_1_Minute":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Minute
	case "Enum_TimeUnitType_Member_1_Second":
		*enum_timeunittype_member_1 = Enum_TimeUnitType_Member_1_Second
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_timeunittype_member_1 *Enum_TimeUnitType_Member_1) ToCodeString() (res string) {

	switch *enum_timeunittype_member_1 {
	// insertion code per enum code
	case Enum_TimeUnitType_Member_1_Year:
		res = "Enum_TimeUnitType_Member_1_Year"
	case Enum_TimeUnitType_Member_1_Month:
		res = "Enum_TimeUnitType_Member_1_Month"
	case Enum_TimeUnitType_Member_1_Day:
		res = "Enum_TimeUnitType_Member_1_Day"
	case Enum_TimeUnitType_Member_1_Hour:
		res = "Enum_TimeUnitType_Member_1_Hour"
	case Enum_TimeUnitType_Member_1_Minute:
		res = "Enum_TimeUnitType_Member_1_Minute"
	case Enum_TimeUnitType_Member_1_Second:
		res = "Enum_TimeUnitType_Member_1_Second"
	}
	return
}

func (enum_timeunittype_member_1 Enum_TimeUnitType_Member_1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_TimeUnitType_Member_1_Year")
	res = append(res, "Enum_TimeUnitType_Member_1_Month")
	res = append(res, "Enum_TimeUnitType_Member_1_Day")
	res = append(res, "Enum_TimeUnitType_Member_1_Hour")
	res = append(res, "Enum_TimeUnitType_Member_1_Minute")
	res = append(res, "Enum_TimeUnitType_Member_1_Second")

	return
}

func (enum_timeunittype_member_1 Enum_TimeUnitType_Member_1) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "year")
	res = append(res, "month")
	res = append(res, "day")
	res = append(res, "hour")
	res = append(res, "minute")
	res = append(res, "second")

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
