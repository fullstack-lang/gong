// generated code - do not edit
package gml

// insertion point
func (inst *AbstractCurveSegmentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractFeatureMemberType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractFeatureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGMLType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeometricPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeometryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractMetadataPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractRingPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractRingType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractSurfacePatchType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeGeometricPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeObjectType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimePrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeSliceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AngleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ArcByCenterPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcByCenterPoint":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcByCenterPoint_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcStringType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcString":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcString_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Arc":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Arc_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *BoundingShapeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CircleByCenterPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "CircleByCenterPoint":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_CircleByCenterPoint_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CircleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Circle":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Circle_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CodeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractCurveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractcurvetype, ok := stage.AbstractCurveType_NameXSD_reverseMap[inst]; ok {
				res = _abstractcurvetype.Name
			}
		}
	case "AbstractFeatureType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractfeaturetype, ok := stage.AbstractFeatureType_NameXSD_reverseMap[inst]; ok {
				res = _abstractfeaturetype.Name
			}
		}
	case "AbstractGMLType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractgmltype, ok := stage.AbstractGMLType_NameXSD_reverseMap[inst]; ok {
				res = _abstractgmltype.Name
			}
		}
	case "AbstractGeometricPrimitiveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractgeometricprimitivetype, ok := stage.AbstractGeometricPrimitiveType_NameXSD_reverseMap[inst]; ok {
				res = _abstractgeometricprimitivetype.Name
			}
		}
	case "AbstractGeometryType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractgeometrytype, ok := stage.AbstractGeometryType_NameXSD_reverseMap[inst]; ok {
				res = _abstractgeometrytype.Name
			}
		}
	case "AbstractSurfaceType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstractsurfacetype, ok := stage.AbstractSurfaceType_NameXSD_reverseMap[inst]; ok {
				res = _abstractsurfacetype.Name
			}
		}
	case "AbstractTimeGeometricPrimitiveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstracttimegeometricprimitivetype, ok := stage.AbstractTimeGeometricPrimitiveType_NameXSD_reverseMap[inst]; ok {
				res = _abstracttimegeometricprimitivetype.Name
			}
		}
	case "AbstractTimeObjectType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstracttimeobjecttype, ok := stage.AbstractTimeObjectType_NameXSD_reverseMap[inst]; ok {
				res = _abstracttimeobjecttype.Name
			}
		}
	case "AbstractTimePrimitiveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstracttimeprimitivetype, ok := stage.AbstractTimePrimitiveType_NameXSD_reverseMap[inst]; ok {
				res = _abstracttimeprimitivetype.Name
			}
		}
	case "AbstractTimeSliceType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _abstracttimeslicetype, ok := stage.AbstractTimeSliceType_NameXSD_reverseMap[inst]; ok {
				res = _abstracttimeslicetype.Name
			}
		}
	case "CompositeCurveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _compositecurvetype, ok := stage.CompositeCurveType_NameXSD_reverseMap[inst]; ok {
				res = _compositecurvetype.Name
			}
		}
	case "CurveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _curvetype, ok := stage.CurveType_NameXSD_reverseMap[inst]; ok {
				res = _curvetype.Name
			}
		}
	case "DynamicFeatureType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _dynamicfeaturetype, ok := stage.DynamicFeatureType_NameXSD_reverseMap[inst]; ok {
				res = _dynamicfeaturetype.Name
			}
		}
	case "OrientableCurveType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _orientablecurvetype, ok := stage.OrientableCurveType_NameXSD_reverseMap[inst]; ok {
				res = _orientablecurvetype.Name
			}
		}
	case "PointType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _pointtype, ok := stage.PointType_NameXSD_reverseMap[inst]; ok {
				res = _pointtype.Name
			}
		}
	case "SurfaceType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _surfacetype, ok := stage.SurfaceType_NameXSD_reverseMap[inst]; ok {
				res = _surfacetype.Name
			}
		}
	case "TimeInstantType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _timeinstanttype, ok := stage.TimeInstantType_NameXSD_reverseMap[inst]; ok {
				res = _timeinstanttype.Name
			}
		}
	case "TimePeriodType":
		switch reverseField.Fieldname {
		case "NameXSD":
			if _timeperiodtype, ok := stage.TimePeriodType_NameXSD_reverseMap[inst]; ok {
				res = _timeperiodtype.Name
			}
		}
	}
	return
}

func (inst *CodeWithAuthorityType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CompositeCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CurvePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompositeCurveType":
		switch reverseField.Fieldname {
		case "CurveMember":
			if _compositecurvetype, ok := stage.CompositeCurveType_CurveMember_reverseMap[inst]; ok {
				res = _compositecurvetype.Name
			}
		}
	case "RingType":
		switch reverseField.Fieldname {
		case "CurveMember":
			if _ringtype, ok := stage.RingType_CurveMember_reverseMap[inst]; ok {
				res = _ringtype.Name
			}
		}
	}
	return
}

func (inst *CurveSegmentArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectPositionListType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectPositionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArcStringType":
		switch reverseField.Fieldname {
		case "Pos":
			if _arcstringtype, ok := stage.ArcStringType_Pos_reverseMap[inst]; ok {
				res = _arcstringtype.Name
			}
		}
	case "CircleType":
		switch reverseField.Fieldname {
		case "Pos":
			if _circletype, ok := stage.CircleType_Pos_reverseMap[inst]; ok {
				res = _circletype.Name
			}
		}
	case "LineStringSegmentType":
		switch reverseField.Fieldname {
		case "Pos":
			if _linestringsegmenttype, ok := stage.LineStringSegmentType_Pos_reverseMap[inst]; ok {
				res = _linestringsegmenttype.Name
			}
		}
	}
	return
}

func (inst *DynamicFeatureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EnvelopeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeodesicStringType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "GeodesicString":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_GeodesicString_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *GeodesicType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Geodesic":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Geodesic_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *LengthType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LineStringSegmentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "LineStringSegment":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_LineStringSegment_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MeasureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OrientableCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PointPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArcStringType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _arcstringtype, ok := stage.ArcStringType_PointProperty_reverseMap[inst]; ok {
				res = _arcstringtype.Name
			}
		}
	case "CircleType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _circletype, ok := stage.CircleType_PointProperty_reverseMap[inst]; ok {
				res = _circletype.Name
			}
		}
	case "LineStringSegmentType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _linestringsegmenttype, ok := stage.LineStringSegmentType_PointProperty_reverseMap[inst]; ok {
				res = _linestringsegmenttype.Name
			}
		}
	}
	return
}

func (inst *PointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PolygonPatchType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "PolygonPatch":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_PolygonPatch_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *RingType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SurfacePatchArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeInstantType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimePeriodType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimePositionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimePrimitivePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
