// generated code - do not edit
package gml

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractCurveSegmentTypes[abstractcurvesegmenttype]
	return ok
}

func (abstractcurvetype *AbstractCurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractCurveTypes[abstractcurvetype]
	return ok
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractFeatureMemberTypes[abstractfeaturemembertype]
	return ok
}

func (abstractfeaturetype *AbstractFeatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractFeatureTypes[abstractfeaturetype]
	return ok
}

func (abstractgmltype *AbstractGMLType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGMLTypes[abstractgmltype]
	return ok
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGeometricPrimitiveTypes[abstractgeometricprimitivetype]
	return ok
}

func (abstractgeometrytype *AbstractGeometryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGeometryTypes[abstractgeometrytype]
	return ok
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractMetadataPropertyTypes[abstractmetadatapropertytype]
	return ok
}

func (abstractringpropertytype *AbstractRingPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractRingPropertyTypes[abstractringpropertytype]
	return ok
}

func (abstractringtype *AbstractRingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractRingTypes[abstractringtype]
	return ok
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractSurfacePatchTypes[abstractsurfacepatchtype]
	return ok
}

func (abstractsurfacetype *AbstractSurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractSurfaceTypes[abstractsurfacetype]
	return ok
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeGeometricPrimitiveTypes[abstracttimegeometricprimitivetype]
	return ok
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeObjectTypes[abstracttimeobjecttype]
	return ok
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimePrimitiveTypes[abstracttimeprimitivetype]
	return ok
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeSliceTypes[abstracttimeslicetype]
	return ok
}

func (angletype *AngleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AngleTypes[angletype]
	return ok
}

func (arcbycenterpointtype *ArcByCenterPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcByCenterPointTypes[arcbycenterpointtype]
	return ok
}

func (arcstringtype *ArcStringType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcStringTypes[arcstringtype]
	return ok
}

func (arctype *ArcType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcTypes[arctype]
	return ok
}

func (boundingshapetype *BoundingShapeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BoundingShapeTypes[boundingshapetype]
	return ok
}

func (circlebycenterpointtype *CircleByCenterPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CircleByCenterPointTypes[circlebycenterpointtype]
	return ok
}

func (circletype *CircleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CircleTypes[circletype]
	return ok
}

func (codetype *CodeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTypes[codetype]
	return ok
}

func (codewithauthoritytype *CodeWithAuthorityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWithAuthorityTypes[codewithauthoritytype]
	return ok
}

func (compositecurvetype *CompositeCurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CompositeCurveTypes[compositecurvetype]
	return ok
}

func (curvepropertytype *CurvePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurvePropertyTypes[curvepropertytype]
	return ok
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurveSegmentArrayPropertyTypes[curvesegmentarraypropertytype]
	return ok
}

func (curvetype *CurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurveTypes[curvetype]
	return ok
}

func (directpositionlisttype *DirectPositionListType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DirectPositionListTypes[directpositionlisttype]
	return ok
}

func (directpositiontype *DirectPositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DirectPositionTypes[directpositiontype]
	return ok
}

func (dynamicfeaturetype *DynamicFeatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DynamicFeatureTypes[dynamicfeaturetype]
	return ok
}

func (envelopetype *EnvelopeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EnvelopeTypes[envelopetype]
	return ok
}

func (geodesicstringtype *GeodesicStringType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeodesicStringTypes[geodesicstringtype]
	return ok
}

func (geodesictype *GeodesicType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeodesicTypes[geodesictype]
	return ok
}

func (lengthtype *LengthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LengthTypes[lengthtype]
	return ok
}

func (linestringsegmenttype *LineStringSegmentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LineStringSegmentTypes[linestringsegmenttype]
	return ok
}

func (measuretype *MeasureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureTypes[measuretype]
	return ok
}

func (orientablecurvetype *OrientableCurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OrientableCurveTypes[orientablecurvetype]
	return ok
}

func (pointpropertytype *PointPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointPropertyTypes[pointpropertytype]
	return ok
}

func (pointtype *PointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointTypes[pointtype]
	return ok
}

func (polygonpatchtype *PolygonPatchType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PolygonPatchTypes[polygonpatchtype]
	return ok
}

func (ringtype *RingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RingTypes[ringtype]
	return ok
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SurfacePatchArrayPropertyTypes[surfacepatcharraypropertytype]
	return ok
}

func (surfacetype *SurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SurfaceTypes[surfacetype]
	return ok
}

func (timeinstanttype *TimeInstantType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeInstantTypes[timeinstanttype]
	return ok
}

func (timeperiodtype *TimePeriodType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePeriodTypes[timeperiodtype]
	return ok
}

func (timepositiontype *TimePositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePositionTypes[timepositiontype]
	return ok
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePrimitivePropertyTypes[timeprimitivepropertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractcurvesegmenttype) {
		return
	}

	abstractcurvesegmenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractcurvetype *AbstractCurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractcurvetype) {
		return
	}

	abstractcurvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractcurvetype.Identifier != nil {
		stage.StageBranch(abstractcurvetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractcurvetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractfeaturemembertype) {
		return
	}

	abstractfeaturemembertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractfeaturetype *AbstractFeatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractfeaturetype) {
		return
	}

	abstractfeaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetype.Identifier != nil {
		stage.StageBranch(abstractfeaturetype.Identifier)
	}
	if abstractfeaturetype.BoundedBy != nil {
		stage.StageBranch(abstractfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgmltype *AbstractGMLType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgmltype) {
		return
	}

	abstractgmltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltype.Identifier != nil {
		stage.StageBranch(abstractgmltype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgeometricprimitivetype) {
		return
	}

	abstractgeometricprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetype.Identifier != nil {
		stage.StageBranch(abstractgeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgeometrytype *AbstractGeometryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgeometrytype) {
		return
	}

	abstractgeometrytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytype.Identifier != nil {
		stage.StageBranch(abstractgeometrytype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractmetadatapropertytype) {
		return
	}

	abstractmetadatapropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractringpropertytype *AbstractRingPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractringpropertytype) {
		return
	}

	abstractringpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractringpropertytype.Ring != nil {
		stage.StageBranch(abstractringpropertytype.Ring)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractringtype *AbstractRingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractringtype) {
		return
	}

	abstractringtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractsurfacepatchtype) {
		return
	}

	abstractsurfacepatchtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractsurfacetype *AbstractSurfaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractsurfacetype) {
		return
	}

	abstractsurfacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractsurfacetype.Identifier != nil {
		stage.StageBranch(abstractsurfacetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractsurfacetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimegeometricprimitivetype) {
		return
	}

	abstracttimegeometricprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetype.Identifier != nil {
		stage.StageBranch(abstracttimegeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeobjecttype) {
		return
	}

	abstracttimeobjecttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttype.Identifier != nil {
		stage.StageBranch(abstracttimeobjecttype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeprimitivetype) {
		return
	}

	abstracttimeprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetype.Identifier != nil {
		stage.StageBranch(abstracttimeprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeslicetype *AbstractTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeslicetype) {
		return
	}

	abstracttimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetype.Identifier != nil {
		stage.StageBranch(abstracttimeslicetype.Identifier)
	}
	if abstracttimeslicetype.ValidTime != nil {
		stage.StageBranch(abstracttimeslicetype.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (angletype *AngleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(angletype) {
		return
	}

	angletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcbycenterpointtype *ArcByCenterPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arcbycenterpointtype) {
		return
	}

	arcbycenterpointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arcbycenterpointtype.Pos != nil {
		stage.StageBranch(arcbycenterpointtype.Pos)
	}
	if arcbycenterpointtype.PointProperty != nil {
		stage.StageBranch(arcbycenterpointtype.PointProperty)
	}
	if arcbycenterpointtype.PosList != nil {
		stage.StageBranch(arcbycenterpointtype.PosList)
	}
	if arcbycenterpointtype.Radius != nil {
		stage.StageBranch(arcbycenterpointtype.Radius)
	}
	if arcbycenterpointtype.StartAngle != nil {
		stage.StageBranch(arcbycenterpointtype.StartAngle)
	}
	if arcbycenterpointtype.EndAngle != nil {
		stage.StageBranch(arcbycenterpointtype.EndAngle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcstringtype *ArcStringType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arcstringtype) {
		return
	}

	arcstringtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arcstringtype.PosList != nil {
		stage.StageBranch(arcstringtype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range arcstringtype.Pos {
		stage.StageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range arcstringtype.PointProperty {
		stage.StageBranch(_pointpropertytype)
	}

}

func (arctype *ArcType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arctype) {
		return
	}

	arctype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arctype.Pos != nil {
		stage.StageBranch(arctype.Pos)
	}
	if arctype.PointProperty != nil {
		stage.StageBranch(arctype.PointProperty)
	}
	if arctype.PosList != nil {
		stage.StageBranch(arctype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (boundingshapetype *BoundingShapeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(boundingshapetype) {
		return
	}

	boundingshapetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetype.Envelope != nil {
		stage.StageBranch(boundingshapetype.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlebycenterpointtype *CircleByCenterPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(circlebycenterpointtype) {
		return
	}

	circlebycenterpointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circlebycenterpointtype.Pos != nil {
		stage.StageBranch(circlebycenterpointtype.Pos)
	}
	if circlebycenterpointtype.PointProperty != nil {
		stage.StageBranch(circlebycenterpointtype.PointProperty)
	}
	if circlebycenterpointtype.PosList != nil {
		stage.StageBranch(circlebycenterpointtype.PosList)
	}
	if circlebycenterpointtype.Radius != nil {
		stage.StageBranch(circlebycenterpointtype.Radius)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circletype *CircleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(circletype) {
		return
	}

	circletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circletype.PosList != nil {
		stage.StageBranch(circletype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range circletype.Pos {
		stage.StageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range circletype.PointProperty {
		stage.StageBranch(_pointpropertytype)
	}

}

func (codetype *CodeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetype) {
		return
	}

	codetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (compositecurvetype *CompositeCurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(compositecurvetype) {
		return
	}

	compositecurvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compositecurvetype.Identifier != nil {
		stage.StageBranch(compositecurvetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range compositecurvetype.NameXSD {
		stage.StageBranch(_codetype)
	}
	for _, _curvepropertytype := range compositecurvetype.CurveMember {
		stage.StageBranch(_curvepropertytype)
	}

}

func (curvepropertytype *CurvePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curvepropertytype) {
		return
	}

	curvepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytype.Curve != nil {
		stage.StageBranch(curvepropertytype.Curve)
	}
	if curvepropertytype.OrientableCurve != nil {
		stage.StageBranch(curvepropertytype.OrientableCurve)
	}
	if curvepropertytype.CompositeCurve != nil {
		stage.StageBranch(curvepropertytype.CompositeCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curvesegmentarraypropertytype) {
		return
	}

	curvesegmentarraypropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linestringsegmenttype := range curvesegmentarraypropertytype.LineStringSegment {
		stage.StageBranch(_linestringsegmenttype)
	}
	for _, _arcstringtype := range curvesegmentarraypropertytype.ArcString {
		stage.StageBranch(_arcstringtype)
	}
	for _, _arctype := range curvesegmentarraypropertytype.Arc {
		stage.StageBranch(_arctype)
	}
	for _, _circletype := range curvesegmentarraypropertytype.Circle {
		stage.StageBranch(_circletype)
	}
	for _, _arcbycenterpointtype := range curvesegmentarraypropertytype.ArcByCenterPoint {
		stage.StageBranch(_arcbycenterpointtype)
	}
	for _, _circlebycenterpointtype := range curvesegmentarraypropertytype.CircleByCenterPoint {
		stage.StageBranch(_circlebycenterpointtype)
	}
	for _, _geodesicstringtype := range curvesegmentarraypropertytype.GeodesicString {
		stage.StageBranch(_geodesicstringtype)
	}
	for _, _geodesictype := range curvesegmentarraypropertytype.Geodesic {
		stage.StageBranch(_geodesictype)
	}

}

func (curvetype *CurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curvetype) {
		return
	}

	curvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvetype.Identifier != nil {
		stage.StageBranch(curvetype.Identifier)
	}
	if curvetype.Segments != nil {
		stage.StageBranch(curvetype.Segments)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range curvetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (directpositionlisttype *DirectPositionListType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(directpositionlisttype) {
		return
	}

	directpositionlisttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directpositiontype *DirectPositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(directpositiontype) {
		return
	}

	directpositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamicfeaturetype *DynamicFeatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dynamicfeaturetype) {
		return
	}

	dynamicfeaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetype.Identifier != nil {
		stage.StageBranch(dynamicfeaturetype.Identifier)
	}
	if dynamicfeaturetype.BoundedBy != nil {
		stage.StageBranch(dynamicfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (envelopetype *EnvelopeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(envelopetype) {
		return
	}

	envelopetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if envelopetype.LowerCorner != nil {
		stage.StageBranch(envelopetype.LowerCorner)
	}
	if envelopetype.UpperCorner != nil {
		stage.StageBranch(envelopetype.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geodesicstringtype *GeodesicStringType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geodesicstringtype) {
		return
	}

	geodesicstringtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geodesicstringtype.PosList != nil {
		stage.StageBranch(geodesicstringtype.PosList)
	}
	if geodesicstringtype.Pos != nil {
		stage.StageBranch(geodesicstringtype.Pos)
	}
	if geodesicstringtype.PointProperty != nil {
		stage.StageBranch(geodesicstringtype.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geodesictype *GeodesicType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geodesictype) {
		return
	}

	geodesictype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geodesictype.PosList != nil {
		stage.StageBranch(geodesictype.PosList)
	}
	if geodesictype.Pos != nil {
		stage.StageBranch(geodesictype.Pos)
	}
	if geodesictype.PointProperty != nil {
		stage.StageBranch(geodesictype.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lengthtype *LengthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lengthtype) {
		return
	}

	lengthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linestringsegmenttype *LineStringSegmentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linestringsegmenttype) {
		return
	}

	linestringsegmenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linestringsegmenttype.PosList != nil {
		stage.StageBranch(linestringsegmenttype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range linestringsegmenttype.Pos {
		stage.StageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range linestringsegmenttype.PointProperty {
		stage.StageBranch(_pointpropertytype)
	}

}

func (measuretype *MeasureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measuretype) {
		return
	}

	measuretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (orientablecurvetype *OrientableCurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(orientablecurvetype) {
		return
	}

	orientablecurvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if orientablecurvetype.Identifier != nil {
		stage.StageBranch(orientablecurvetype.Identifier)
	}
	if orientablecurvetype.BaseCurve != nil {
		stage.StageBranch(orientablecurvetype.BaseCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range orientablecurvetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (pointpropertytype *PointPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pointpropertytype) {
		return
	}

	pointpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointpropertytype.Point != nil {
		stage.StageBranch(pointpropertytype.Point)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pointtype) {
		return
	}

	pointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointtype.Identifier != nil {
		stage.StageBranch(pointtype.Identifier)
	}
	if pointtype.Pos != nil {
		stage.StageBranch(pointtype.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (polygonpatchtype *PolygonPatchType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(polygonpatchtype) {
		return
	}

	polygonpatchtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if polygonpatchtype.Exterior != nil {
		stage.StageBranch(polygonpatchtype.Exterior)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ringtype *RingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ringtype) {
		return
	}

	ringtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _curvepropertytype := range ringtype.CurveMember {
		stage.StageBranch(_curvepropertytype)
	}

}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(surfacepatcharraypropertytype) {
		return
	}

	surfacepatcharraypropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygonpatchtype := range surfacepatcharraypropertytype.PolygonPatch {
		stage.StageBranch(_polygonpatchtype)
	}

}

func (surfacetype *SurfaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(surfacetype) {
		return
	}

	surfacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if surfacetype.Identifier != nil {
		stage.StageBranch(surfacetype.Identifier)
	}
	if surfacetype.Patches != nil {
		stage.StageBranch(surfacetype.Patches)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range surfacetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timeinstanttype *TimeInstantType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeinstanttype) {
		return
	}

	timeinstanttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttype.Identifier != nil {
		stage.StageBranch(timeinstanttype.Identifier)
	}
	if timeinstanttype.TimePosition != nil {
		stage.StageBranch(timeinstanttype.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timeperiodtype *TimePeriodType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeperiodtype) {
		return
	}

	timeperiodtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtype.Identifier != nil {
		stage.StageBranch(timeperiodtype.Identifier)
	}
	if timeperiodtype.BeginPosition != nil {
		stage.StageBranch(timeperiodtype.BeginPosition)
	}
	if timeperiodtype.EndPosition != nil {
		stage.StageBranch(timeperiodtype.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timepositiontype *TimePositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timepositiontype) {
		return
	}

	timepositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeprimitivepropertytype) {
		return
	}

	timeprimitivepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytype.TimeInstant != nil {
		stage.StageBranch(timeprimitivepropertytype.TimeInstant)
	}
	if timeprimitivepropertytype.TimePeriod != nil {
		stage.StageBranch(timeprimitivepropertytype.TimePeriod)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AbstractCurveSegmentType:
		toT := GongCopyBranchAbstractCurveSegmentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractCurveType:
		toT := GongCopyBranchAbstractCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractFeatureMemberType:
		toT := GongCopyBranchAbstractFeatureMemberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractFeatureType:
		toT := GongCopyBranchAbstractFeatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGMLType:
		toT := GongCopyBranchAbstractGMLType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGeometricPrimitiveType:
		toT := GongCopyBranchAbstractGeometricPrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGeometryType:
		toT := GongCopyBranchAbstractGeometryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractMetadataPropertyType:
		toT := GongCopyBranchAbstractMetadataPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractRingPropertyType:
		toT := GongCopyBranchAbstractRingPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractRingType:
		toT := GongCopyBranchAbstractRingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractSurfacePatchType:
		toT := GongCopyBranchAbstractSurfacePatchType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractSurfaceType:
		toT := GongCopyBranchAbstractSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeGeometricPrimitiveType:
		toT := GongCopyBranchAbstractTimeGeometricPrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeObjectType:
		toT := GongCopyBranchAbstractTimeObjectType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimePrimitiveType:
		toT := GongCopyBranchAbstractTimePrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeSliceType:
		toT := GongCopyBranchAbstractTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AngleType:
		toT := GongCopyBranchAngleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcByCenterPointType:
		toT := GongCopyBranchArcByCenterPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcStringType:
		toT := GongCopyBranchArcStringType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcType:
		toT := GongCopyBranchArcType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BoundingShapeType:
		toT := GongCopyBranchBoundingShapeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleByCenterPointType:
		toT := GongCopyBranchCircleByCenterPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleType:
		toT := GongCopyBranchCircleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeType:
		toT := GongCopyBranchCodeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeWithAuthorityType:
		toT := GongCopyBranchCodeWithAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CompositeCurveType:
		toT := GongCopyBranchCompositeCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurvePropertyType:
		toT := GongCopyBranchCurvePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurveSegmentArrayPropertyType:
		toT := GongCopyBranchCurveSegmentArrayPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurveType:
		toT := GongCopyBranchCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectPositionListType:
		toT := GongCopyBranchDirectPositionListType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectPositionType:
		toT := GongCopyBranchDirectPositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DynamicFeatureType:
		toT := GongCopyBranchDynamicFeatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EnvelopeType:
		toT := GongCopyBranchEnvelopeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeodesicStringType:
		toT := GongCopyBranchGeodesicStringType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeodesicType:
		toT := GongCopyBranchGeodesicType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LengthType:
		toT := GongCopyBranchLengthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LineStringSegmentType:
		toT := GongCopyBranchLineStringSegmentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureType:
		toT := GongCopyBranchMeasureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OrientableCurveType:
		toT := GongCopyBranchOrientableCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointPropertyType:
		toT := GongCopyBranchPointPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointType:
		toT := GongCopyBranchPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PolygonPatchType:
		toT := GongCopyBranchPolygonPatchType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RingType:
		toT := GongCopyBranchRingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SurfacePatchArrayPropertyType:
		toT := GongCopyBranchSurfacePatchArrayPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SurfaceType:
		toT := GongCopyBranchSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeInstantType:
		toT := GongCopyBranchTimeInstantType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePeriodType:
		toT := GongCopyBranchTimePeriodType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePositionType:
		toT := GongCopyBranchTimePositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePrimitivePropertyType:
		toT := GongCopyBranchTimePrimitivePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractCurveSegmentType(mapOrigCopy map[any]any, abstractcurvesegmenttypeFrom *AbstractCurveSegmentType) (abstractcurvesegmenttypeTo *AbstractCurveSegmentType) {
	var alreadyCopied bool
	abstractcurvesegmenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractcurvesegmenttypeFrom)
	if alreadyCopied {
		return
	}
	abstractcurvesegmenttypeFrom.GongCopyBasicFields(abstractcurvesegmenttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractCurveType(mapOrigCopy map[any]any, abstractcurvetypeFrom *AbstractCurveType) (abstractcurvetypeTo *AbstractCurveType) {
	var alreadyCopied bool
	abstractcurvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractcurvetypeFrom)
	if alreadyCopied {
		return
	}
	abstractcurvetypeFrom.GongCopyBasicFields(abstractcurvetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractcurvetypeFrom.Identifier != nil {
		abstractcurvetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractcurvetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractcurvetypeFrom.NameXSD {
		abstractcurvetypeTo.NameXSD = append(abstractcurvetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractFeatureMemberType(mapOrigCopy map[any]any, abstractfeaturemembertypeFrom *AbstractFeatureMemberType) (abstractfeaturemembertypeTo *AbstractFeatureMemberType) {
	var alreadyCopied bool
	abstractfeaturemembertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractfeaturemembertypeFrom)
	if alreadyCopied {
		return
	}
	abstractfeaturemembertypeFrom.GongCopyBasicFields(abstractfeaturemembertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractFeatureType(mapOrigCopy map[any]any, abstractfeaturetypeFrom *AbstractFeatureType) (abstractfeaturetypeTo *AbstractFeatureType) {
	var alreadyCopied bool
	abstractfeaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractfeaturetypeFrom)
	if alreadyCopied {
		return
	}
	abstractfeaturetypeFrom.GongCopyBasicFields(abstractfeaturetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetypeFrom.Identifier != nil {
		abstractfeaturetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractfeaturetypeFrom.Identifier)
	}
	if abstractfeaturetypeFrom.BoundedBy != nil {
		abstractfeaturetypeTo.BoundedBy = GongCopyBranchBoundingShapeType(mapOrigCopy, abstractfeaturetypeFrom.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetypeFrom.NameXSD {
		abstractfeaturetypeTo.NameXSD = append(abstractfeaturetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGMLType(mapOrigCopy map[any]any, abstractgmltypeFrom *AbstractGMLType) (abstractgmltypeTo *AbstractGMLType) {
	var alreadyCopied bool
	abstractgmltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgmltypeFrom)
	if alreadyCopied {
		return
	}
	abstractgmltypeFrom.GongCopyBasicFields(abstractgmltypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltypeFrom.Identifier != nil {
		abstractgmltypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgmltypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltypeFrom.NameXSD {
		abstractgmltypeTo.NameXSD = append(abstractgmltypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGeometricPrimitiveType(mapOrigCopy map[any]any, abstractgeometricprimitivetypeFrom *AbstractGeometricPrimitiveType) (abstractgeometricprimitivetypeTo *AbstractGeometricPrimitiveType) {
	var alreadyCopied bool
	abstractgeometricprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgeometricprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstractgeometricprimitivetypeFrom.GongCopyBasicFields(abstractgeometricprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetypeFrom.Identifier != nil {
		abstractgeometricprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgeometricprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetypeFrom.NameXSD {
		abstractgeometricprimitivetypeTo.NameXSD = append(abstractgeometricprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGeometryType(mapOrigCopy map[any]any, abstractgeometrytypeFrom *AbstractGeometryType) (abstractgeometrytypeTo *AbstractGeometryType) {
	var alreadyCopied bool
	abstractgeometrytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgeometrytypeFrom)
	if alreadyCopied {
		return
	}
	abstractgeometrytypeFrom.GongCopyBasicFields(abstractgeometrytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytypeFrom.Identifier != nil {
		abstractgeometrytypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgeometrytypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytypeFrom.NameXSD {
		abstractgeometrytypeTo.NameXSD = append(abstractgeometrytypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractMetadataPropertyType(mapOrigCopy map[any]any, abstractmetadatapropertytypeFrom *AbstractMetadataPropertyType) (abstractmetadatapropertytypeTo *AbstractMetadataPropertyType) {
	var alreadyCopied bool
	abstractmetadatapropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractmetadatapropertytypeFrom)
	if alreadyCopied {
		return
	}
	abstractmetadatapropertytypeFrom.GongCopyBasicFields(abstractmetadatapropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractRingPropertyType(mapOrigCopy map[any]any, abstractringpropertytypeFrom *AbstractRingPropertyType) (abstractringpropertytypeTo *AbstractRingPropertyType) {
	var alreadyCopied bool
	abstractringpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractringpropertytypeFrom)
	if alreadyCopied {
		return
	}
	abstractringpropertytypeFrom.GongCopyBasicFields(abstractringpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractringpropertytypeFrom.Ring != nil {
		abstractringpropertytypeTo.Ring = GongCopyBranchRingType(mapOrigCopy, abstractringpropertytypeFrom.Ring)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractRingType(mapOrigCopy map[any]any, abstractringtypeFrom *AbstractRingType) (abstractringtypeTo *AbstractRingType) {
	var alreadyCopied bool
	abstractringtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractringtypeFrom)
	if alreadyCopied {
		return
	}
	abstractringtypeFrom.GongCopyBasicFields(abstractringtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractSurfacePatchType(mapOrigCopy map[any]any, abstractsurfacepatchtypeFrom *AbstractSurfacePatchType) (abstractsurfacepatchtypeTo *AbstractSurfacePatchType) {
	var alreadyCopied bool
	abstractsurfacepatchtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractsurfacepatchtypeFrom)
	if alreadyCopied {
		return
	}
	abstractsurfacepatchtypeFrom.GongCopyBasicFields(abstractsurfacepatchtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractSurfaceType(mapOrigCopy map[any]any, abstractsurfacetypeFrom *AbstractSurfaceType) (abstractsurfacetypeTo *AbstractSurfaceType) {
	var alreadyCopied bool
	abstractsurfacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractsurfacetypeFrom)
	if alreadyCopied {
		return
	}
	abstractsurfacetypeFrom.GongCopyBasicFields(abstractsurfacetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractsurfacetypeFrom.Identifier != nil {
		abstractsurfacetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractsurfacetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractsurfacetypeFrom.NameXSD {
		abstractsurfacetypeTo.NameXSD = append(abstractsurfacetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimeGeometricPrimitiveType(mapOrigCopy map[any]any, abstracttimegeometricprimitivetypeFrom *AbstractTimeGeometricPrimitiveType) (abstracttimegeometricprimitivetypeTo *AbstractTimeGeometricPrimitiveType) {
	var alreadyCopied bool
	abstracttimegeometricprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimegeometricprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimegeometricprimitivetypeFrom.GongCopyBasicFields(abstracttimegeometricprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetypeFrom.Identifier != nil {
		abstracttimegeometricprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimegeometricprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetypeFrom.NameXSD {
		abstracttimegeometricprimitivetypeTo.NameXSD = append(abstracttimegeometricprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimeObjectType(mapOrigCopy map[any]any, abstracttimeobjecttypeFrom *AbstractTimeObjectType) (abstracttimeobjecttypeTo *AbstractTimeObjectType) {
	var alreadyCopied bool
	abstracttimeobjecttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeobjecttypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeobjecttypeFrom.GongCopyBasicFields(abstracttimeobjecttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttypeFrom.Identifier != nil {
		abstracttimeobjecttypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeobjecttypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttypeFrom.NameXSD {
		abstracttimeobjecttypeTo.NameXSD = append(abstracttimeobjecttypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimePrimitiveType(mapOrigCopy map[any]any, abstracttimeprimitivetypeFrom *AbstractTimePrimitiveType) (abstracttimeprimitivetypeTo *AbstractTimePrimitiveType) {
	var alreadyCopied bool
	abstracttimeprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeprimitivetypeFrom.GongCopyBasicFields(abstracttimeprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetypeFrom.Identifier != nil {
		abstracttimeprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetypeFrom.NameXSD {
		abstracttimeprimitivetypeTo.NameXSD = append(abstracttimeprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimeSliceType(mapOrigCopy map[any]any, abstracttimeslicetypeFrom *AbstractTimeSliceType) (abstracttimeslicetypeTo *AbstractTimeSliceType) {
	var alreadyCopied bool
	abstracttimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeslicetypeFrom.GongCopyBasicFields(abstracttimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetypeFrom.Identifier != nil {
		abstracttimeslicetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeslicetypeFrom.Identifier)
	}
	if abstracttimeslicetypeFrom.ValidTime != nil {
		abstracttimeslicetypeTo.ValidTime = GongCopyBranchTimePrimitivePropertyType(mapOrigCopy, abstracttimeslicetypeFrom.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetypeFrom.NameXSD {
		abstracttimeslicetypeTo.NameXSD = append(abstracttimeslicetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAngleType(mapOrigCopy map[any]any, angletypeFrom *AngleType) (angletypeTo *AngleType) {
	var alreadyCopied bool
	angletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, angletypeFrom)
	if alreadyCopied {
		return
	}
	angletypeFrom.GongCopyBasicFields(angletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcByCenterPointType(mapOrigCopy map[any]any, arcbycenterpointtypeFrom *ArcByCenterPointType) (arcbycenterpointtypeTo *ArcByCenterPointType) {
	var alreadyCopied bool
	arcbycenterpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arcbycenterpointtypeFrom)
	if alreadyCopied {
		return
	}
	arcbycenterpointtypeFrom.GongCopyBasicFields(arcbycenterpointtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if arcbycenterpointtypeFrom.Pos != nil {
		arcbycenterpointtypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, arcbycenterpointtypeFrom.Pos)
	}
	if arcbycenterpointtypeFrom.PointProperty != nil {
		arcbycenterpointtypeTo.PointProperty = GongCopyBranchPointPropertyType(mapOrigCopy, arcbycenterpointtypeFrom.PointProperty)
	}
	if arcbycenterpointtypeFrom.PosList != nil {
		arcbycenterpointtypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, arcbycenterpointtypeFrom.PosList)
	}
	if arcbycenterpointtypeFrom.Radius != nil {
		arcbycenterpointtypeTo.Radius = GongCopyBranchLengthType(mapOrigCopy, arcbycenterpointtypeFrom.Radius)
	}
	if arcbycenterpointtypeFrom.StartAngle != nil {
		arcbycenterpointtypeTo.StartAngle = GongCopyBranchAngleType(mapOrigCopy, arcbycenterpointtypeFrom.StartAngle)
	}
	if arcbycenterpointtypeFrom.EndAngle != nil {
		arcbycenterpointtypeTo.EndAngle = GongCopyBranchAngleType(mapOrigCopy, arcbycenterpointtypeFrom.EndAngle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcStringType(mapOrigCopy map[any]any, arcstringtypeFrom *ArcStringType) (arcstringtypeTo *ArcStringType) {
	var alreadyCopied bool
	arcstringtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arcstringtypeFrom)
	if alreadyCopied {
		return
	}
	arcstringtypeFrom.GongCopyBasicFields(arcstringtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if arcstringtypeFrom.PosList != nil {
		arcstringtypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, arcstringtypeFrom.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range arcstringtypeFrom.Pos {
		arcstringtypeTo.Pos = append(arcstringtypeTo.Pos, GongCopyBranchDirectPositionType(mapOrigCopy, _directpositiontype))
	}
	for _, _pointpropertytype := range arcstringtypeFrom.PointProperty {
		arcstringtypeTo.PointProperty = append(arcstringtypeTo.PointProperty, GongCopyBranchPointPropertyType(mapOrigCopy, _pointpropertytype))
	}

	return
}

func GongCopyBranchArcType(mapOrigCopy map[any]any, arctypeFrom *ArcType) (arctypeTo *ArcType) {
	var alreadyCopied bool
	arctypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arctypeFrom)
	if alreadyCopied {
		return
	}
	arctypeFrom.GongCopyBasicFields(arctypeTo)

	//insertion point for the staging of instances referenced by pointers
	if arctypeFrom.Pos != nil {
		arctypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, arctypeFrom.Pos)
	}
	if arctypeFrom.PointProperty != nil {
		arctypeTo.PointProperty = GongCopyBranchPointPropertyType(mapOrigCopy, arctypeFrom.PointProperty)
	}
	if arctypeFrom.PosList != nil {
		arctypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, arctypeFrom.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBoundingShapeType(mapOrigCopy map[any]any, boundingshapetypeFrom *BoundingShapeType) (boundingshapetypeTo *BoundingShapeType) {
	var alreadyCopied bool
	boundingshapetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, boundingshapetypeFrom)
	if alreadyCopied {
		return
	}
	boundingshapetypeFrom.GongCopyBasicFields(boundingshapetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetypeFrom.Envelope != nil {
		boundingshapetypeTo.Envelope = GongCopyBranchEnvelopeType(mapOrigCopy, boundingshapetypeFrom.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircleByCenterPointType(mapOrigCopy map[any]any, circlebycenterpointtypeFrom *CircleByCenterPointType) (circlebycenterpointtypeTo *CircleByCenterPointType) {
	var alreadyCopied bool
	circlebycenterpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, circlebycenterpointtypeFrom)
	if alreadyCopied {
		return
	}
	circlebycenterpointtypeFrom.GongCopyBasicFields(circlebycenterpointtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if circlebycenterpointtypeFrom.Pos != nil {
		circlebycenterpointtypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, circlebycenterpointtypeFrom.Pos)
	}
	if circlebycenterpointtypeFrom.PointProperty != nil {
		circlebycenterpointtypeTo.PointProperty = GongCopyBranchPointPropertyType(mapOrigCopy, circlebycenterpointtypeFrom.PointProperty)
	}
	if circlebycenterpointtypeFrom.PosList != nil {
		circlebycenterpointtypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, circlebycenterpointtypeFrom.PosList)
	}
	if circlebycenterpointtypeFrom.Radius != nil {
		circlebycenterpointtypeTo.Radius = GongCopyBranchLengthType(mapOrigCopy, circlebycenterpointtypeFrom.Radius)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircleType(mapOrigCopy map[any]any, circletypeFrom *CircleType) (circletypeTo *CircleType) {
	var alreadyCopied bool
	circletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, circletypeFrom)
	if alreadyCopied {
		return
	}
	circletypeFrom.GongCopyBasicFields(circletypeTo)

	//insertion point for the staging of instances referenced by pointers
	if circletypeFrom.PosList != nil {
		circletypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, circletypeFrom.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range circletypeFrom.Pos {
		circletypeTo.Pos = append(circletypeTo.Pos, GongCopyBranchDirectPositionType(mapOrigCopy, _directpositiontype))
	}
	for _, _pointpropertytype := range circletypeFrom.PointProperty {
		circletypeTo.PointProperty = append(circletypeTo.PointProperty, GongCopyBranchPointPropertyType(mapOrigCopy, _pointpropertytype))
	}

	return
}

func GongCopyBranchCodeType(mapOrigCopy map[any]any, codetypeFrom *CodeType) (codetypeTo *CodeType) {
	var alreadyCopied bool
	codetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetypeFrom)
	if alreadyCopied {
		return
	}
	codetypeFrom.GongCopyBasicFields(codetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeWithAuthorityType(mapOrigCopy map[any]any, codewithauthoritytypeFrom *CodeWithAuthorityType) (codewithauthoritytypeTo *CodeWithAuthorityType) {
	var alreadyCopied bool
	codewithauthoritytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codewithauthoritytypeFrom)
	if alreadyCopied {
		return
	}
	codewithauthoritytypeFrom.GongCopyBasicFields(codewithauthoritytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCompositeCurveType(mapOrigCopy map[any]any, compositecurvetypeFrom *CompositeCurveType) (compositecurvetypeTo *CompositeCurveType) {
	var alreadyCopied bool
	compositecurvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, compositecurvetypeFrom)
	if alreadyCopied {
		return
	}
	compositecurvetypeFrom.GongCopyBasicFields(compositecurvetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if compositecurvetypeFrom.Identifier != nil {
		compositecurvetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, compositecurvetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range compositecurvetypeFrom.NameXSD {
		compositecurvetypeTo.NameXSD = append(compositecurvetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}
	for _, _curvepropertytype := range compositecurvetypeFrom.CurveMember {
		compositecurvetypeTo.CurveMember = append(compositecurvetypeTo.CurveMember, GongCopyBranchCurvePropertyType(mapOrigCopy, _curvepropertytype))
	}

	return
}

func GongCopyBranchCurvePropertyType(mapOrigCopy map[any]any, curvepropertytypeFrom *CurvePropertyType) (curvepropertytypeTo *CurvePropertyType) {
	var alreadyCopied bool
	curvepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curvepropertytypeFrom)
	if alreadyCopied {
		return
	}
	curvepropertytypeFrom.GongCopyBasicFields(curvepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytypeFrom.Curve != nil {
		curvepropertytypeTo.Curve = GongCopyBranchCurveType(mapOrigCopy, curvepropertytypeFrom.Curve)
	}
	if curvepropertytypeFrom.OrientableCurve != nil {
		curvepropertytypeTo.OrientableCurve = GongCopyBranchOrientableCurveType(mapOrigCopy, curvepropertytypeFrom.OrientableCurve)
	}
	if curvepropertytypeFrom.CompositeCurve != nil {
		curvepropertytypeTo.CompositeCurve = GongCopyBranchCompositeCurveType(mapOrigCopy, curvepropertytypeFrom.CompositeCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCurveSegmentArrayPropertyType(mapOrigCopy map[any]any, curvesegmentarraypropertytypeFrom *CurveSegmentArrayPropertyType) (curvesegmentarraypropertytypeTo *CurveSegmentArrayPropertyType) {
	var alreadyCopied bool
	curvesegmentarraypropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curvesegmentarraypropertytypeFrom)
	if alreadyCopied {
		return
	}
	curvesegmentarraypropertytypeFrom.GongCopyBasicFields(curvesegmentarraypropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linestringsegmenttype := range curvesegmentarraypropertytypeFrom.LineStringSegment {
		curvesegmentarraypropertytypeTo.LineStringSegment = append(curvesegmentarraypropertytypeTo.LineStringSegment, GongCopyBranchLineStringSegmentType(mapOrigCopy, _linestringsegmenttype))
	}
	for _, _arcstringtype := range curvesegmentarraypropertytypeFrom.ArcString {
		curvesegmentarraypropertytypeTo.ArcString = append(curvesegmentarraypropertytypeTo.ArcString, GongCopyBranchArcStringType(mapOrigCopy, _arcstringtype))
	}
	for _, _arctype := range curvesegmentarraypropertytypeFrom.Arc {
		curvesegmentarraypropertytypeTo.Arc = append(curvesegmentarraypropertytypeTo.Arc, GongCopyBranchArcType(mapOrigCopy, _arctype))
	}
	for _, _circletype := range curvesegmentarraypropertytypeFrom.Circle {
		curvesegmentarraypropertytypeTo.Circle = append(curvesegmentarraypropertytypeTo.Circle, GongCopyBranchCircleType(mapOrigCopy, _circletype))
	}
	for _, _arcbycenterpointtype := range curvesegmentarraypropertytypeFrom.ArcByCenterPoint {
		curvesegmentarraypropertytypeTo.ArcByCenterPoint = append(curvesegmentarraypropertytypeTo.ArcByCenterPoint, GongCopyBranchArcByCenterPointType(mapOrigCopy, _arcbycenterpointtype))
	}
	for _, _circlebycenterpointtype := range curvesegmentarraypropertytypeFrom.CircleByCenterPoint {
		curvesegmentarraypropertytypeTo.CircleByCenterPoint = append(curvesegmentarraypropertytypeTo.CircleByCenterPoint, GongCopyBranchCircleByCenterPointType(mapOrigCopy, _circlebycenterpointtype))
	}
	for _, _geodesicstringtype := range curvesegmentarraypropertytypeFrom.GeodesicString {
		curvesegmentarraypropertytypeTo.GeodesicString = append(curvesegmentarraypropertytypeTo.GeodesicString, GongCopyBranchGeodesicStringType(mapOrigCopy, _geodesicstringtype))
	}
	for _, _geodesictype := range curvesegmentarraypropertytypeFrom.Geodesic {
		curvesegmentarraypropertytypeTo.Geodesic = append(curvesegmentarraypropertytypeTo.Geodesic, GongCopyBranchGeodesicType(mapOrigCopy, _geodesictype))
	}

	return
}

func GongCopyBranchCurveType(mapOrigCopy map[any]any, curvetypeFrom *CurveType) (curvetypeTo *CurveType) {
	var alreadyCopied bool
	curvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curvetypeFrom)
	if alreadyCopied {
		return
	}
	curvetypeFrom.GongCopyBasicFields(curvetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if curvetypeFrom.Identifier != nil {
		curvetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, curvetypeFrom.Identifier)
	}
	if curvetypeFrom.Segments != nil {
		curvetypeTo.Segments = GongCopyBranchCurveSegmentArrayPropertyType(mapOrigCopy, curvetypeFrom.Segments)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range curvetypeFrom.NameXSD {
		curvetypeTo.NameXSD = append(curvetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchDirectPositionListType(mapOrigCopy map[any]any, directpositionlisttypeFrom *DirectPositionListType) (directpositionlisttypeTo *DirectPositionListType) {
	var alreadyCopied bool
	directpositionlisttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directpositionlisttypeFrom)
	if alreadyCopied {
		return
	}
	directpositionlisttypeFrom.GongCopyBasicFields(directpositionlisttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDirectPositionType(mapOrigCopy map[any]any, directpositiontypeFrom *DirectPositionType) (directpositiontypeTo *DirectPositionType) {
	var alreadyCopied bool
	directpositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directpositiontypeFrom)
	if alreadyCopied {
		return
	}
	directpositiontypeFrom.GongCopyBasicFields(directpositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDynamicFeatureType(mapOrigCopy map[any]any, dynamicfeaturetypeFrom *DynamicFeatureType) (dynamicfeaturetypeTo *DynamicFeatureType) {
	var alreadyCopied bool
	dynamicfeaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dynamicfeaturetypeFrom)
	if alreadyCopied {
		return
	}
	dynamicfeaturetypeFrom.GongCopyBasicFields(dynamicfeaturetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetypeFrom.Identifier != nil {
		dynamicfeaturetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, dynamicfeaturetypeFrom.Identifier)
	}
	if dynamicfeaturetypeFrom.BoundedBy != nil {
		dynamicfeaturetypeTo.BoundedBy = GongCopyBranchBoundingShapeType(mapOrigCopy, dynamicfeaturetypeFrom.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetypeFrom.NameXSD {
		dynamicfeaturetypeTo.NameXSD = append(dynamicfeaturetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchEnvelopeType(mapOrigCopy map[any]any, envelopetypeFrom *EnvelopeType) (envelopetypeTo *EnvelopeType) {
	var alreadyCopied bool
	envelopetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, envelopetypeFrom)
	if alreadyCopied {
		return
	}
	envelopetypeFrom.GongCopyBasicFields(envelopetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if envelopetypeFrom.LowerCorner != nil {
		envelopetypeTo.LowerCorner = GongCopyBranchDirectPositionType(mapOrigCopy, envelopetypeFrom.LowerCorner)
	}
	if envelopetypeFrom.UpperCorner != nil {
		envelopetypeTo.UpperCorner = GongCopyBranchDirectPositionType(mapOrigCopy, envelopetypeFrom.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGeodesicStringType(mapOrigCopy map[any]any, geodesicstringtypeFrom *GeodesicStringType) (geodesicstringtypeTo *GeodesicStringType) {
	var alreadyCopied bool
	geodesicstringtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geodesicstringtypeFrom)
	if alreadyCopied {
		return
	}
	geodesicstringtypeFrom.GongCopyBasicFields(geodesicstringtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if geodesicstringtypeFrom.PosList != nil {
		geodesicstringtypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, geodesicstringtypeFrom.PosList)
	}
	if geodesicstringtypeFrom.Pos != nil {
		geodesicstringtypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, geodesicstringtypeFrom.Pos)
	}
	if geodesicstringtypeFrom.PointProperty != nil {
		geodesicstringtypeTo.PointProperty = GongCopyBranchPointPropertyType(mapOrigCopy, geodesicstringtypeFrom.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGeodesicType(mapOrigCopy map[any]any, geodesictypeFrom *GeodesicType) (geodesictypeTo *GeodesicType) {
	var alreadyCopied bool
	geodesictypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geodesictypeFrom)
	if alreadyCopied {
		return
	}
	geodesictypeFrom.GongCopyBasicFields(geodesictypeTo)

	//insertion point for the staging of instances referenced by pointers
	if geodesictypeFrom.PosList != nil {
		geodesictypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, geodesictypeFrom.PosList)
	}
	if geodesictypeFrom.Pos != nil {
		geodesictypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, geodesictypeFrom.Pos)
	}
	if geodesictypeFrom.PointProperty != nil {
		geodesictypeTo.PointProperty = GongCopyBranchPointPropertyType(mapOrigCopy, geodesictypeFrom.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLengthType(mapOrigCopy map[any]any, lengthtypeFrom *LengthType) (lengthtypeTo *LengthType) {
	var alreadyCopied bool
	lengthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lengthtypeFrom)
	if alreadyCopied {
		return
	}
	lengthtypeFrom.GongCopyBasicFields(lengthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLineStringSegmentType(mapOrigCopy map[any]any, linestringsegmenttypeFrom *LineStringSegmentType) (linestringsegmenttypeTo *LineStringSegmentType) {
	var alreadyCopied bool
	linestringsegmenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linestringsegmenttypeFrom)
	if alreadyCopied {
		return
	}
	linestringsegmenttypeFrom.GongCopyBasicFields(linestringsegmenttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if linestringsegmenttypeFrom.PosList != nil {
		linestringsegmenttypeTo.PosList = GongCopyBranchDirectPositionListType(mapOrigCopy, linestringsegmenttypeFrom.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range linestringsegmenttypeFrom.Pos {
		linestringsegmenttypeTo.Pos = append(linestringsegmenttypeTo.Pos, GongCopyBranchDirectPositionType(mapOrigCopy, _directpositiontype))
	}
	for _, _pointpropertytype := range linestringsegmenttypeFrom.PointProperty {
		linestringsegmenttypeTo.PointProperty = append(linestringsegmenttypeTo.PointProperty, GongCopyBranchPointPropertyType(mapOrigCopy, _pointpropertytype))
	}

	return
}

func GongCopyBranchMeasureType(mapOrigCopy map[any]any, measuretypeFrom *MeasureType) (measuretypeTo *MeasureType) {
	var alreadyCopied bool
	measuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measuretypeFrom)
	if alreadyCopied {
		return
	}
	measuretypeFrom.GongCopyBasicFields(measuretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOrientableCurveType(mapOrigCopy map[any]any, orientablecurvetypeFrom *OrientableCurveType) (orientablecurvetypeTo *OrientableCurveType) {
	var alreadyCopied bool
	orientablecurvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, orientablecurvetypeFrom)
	if alreadyCopied {
		return
	}
	orientablecurvetypeFrom.GongCopyBasicFields(orientablecurvetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if orientablecurvetypeFrom.Identifier != nil {
		orientablecurvetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, orientablecurvetypeFrom.Identifier)
	}
	if orientablecurvetypeFrom.BaseCurve != nil {
		orientablecurvetypeTo.BaseCurve = GongCopyBranchCurvePropertyType(mapOrigCopy, orientablecurvetypeFrom.BaseCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range orientablecurvetypeFrom.NameXSD {
		orientablecurvetypeTo.NameXSD = append(orientablecurvetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchPointPropertyType(mapOrigCopy map[any]any, pointpropertytypeFrom *PointPropertyType) (pointpropertytypeTo *PointPropertyType) {
	var alreadyCopied bool
	pointpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointpropertytypeFrom)
	if alreadyCopied {
		return
	}
	pointpropertytypeFrom.GongCopyBasicFields(pointpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if pointpropertytypeFrom.Point != nil {
		pointpropertytypeTo.Point = GongCopyBranchPointType(mapOrigCopy, pointpropertytypeFrom.Point)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointType(mapOrigCopy map[any]any, pointtypeFrom *PointType) (pointtypeTo *PointType) {
	var alreadyCopied bool
	pointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointtypeFrom)
	if alreadyCopied {
		return
	}
	pointtypeFrom.GongCopyBasicFields(pointtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if pointtypeFrom.Identifier != nil {
		pointtypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, pointtypeFrom.Identifier)
	}
	if pointtypeFrom.Pos != nil {
		pointtypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, pointtypeFrom.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtypeFrom.NameXSD {
		pointtypeTo.NameXSD = append(pointtypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchPolygonPatchType(mapOrigCopy map[any]any, polygonpatchtypeFrom *PolygonPatchType) (polygonpatchtypeTo *PolygonPatchType) {
	var alreadyCopied bool
	polygonpatchtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, polygonpatchtypeFrom)
	if alreadyCopied {
		return
	}
	polygonpatchtypeFrom.GongCopyBasicFields(polygonpatchtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if polygonpatchtypeFrom.Exterior != nil {
		polygonpatchtypeTo.Exterior = GongCopyBranchAbstractRingPropertyType(mapOrigCopy, polygonpatchtypeFrom.Exterior)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRingType(mapOrigCopy map[any]any, ringtypeFrom *RingType) (ringtypeTo *RingType) {
	var alreadyCopied bool
	ringtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ringtypeFrom)
	if alreadyCopied {
		return
	}
	ringtypeFrom.GongCopyBasicFields(ringtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _curvepropertytype := range ringtypeFrom.CurveMember {
		ringtypeTo.CurveMember = append(ringtypeTo.CurveMember, GongCopyBranchCurvePropertyType(mapOrigCopy, _curvepropertytype))
	}

	return
}

func GongCopyBranchSurfacePatchArrayPropertyType(mapOrigCopy map[any]any, surfacepatcharraypropertytypeFrom *SurfacePatchArrayPropertyType) (surfacepatcharraypropertytypeTo *SurfacePatchArrayPropertyType) {
	var alreadyCopied bool
	surfacepatcharraypropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, surfacepatcharraypropertytypeFrom)
	if alreadyCopied {
		return
	}
	surfacepatcharraypropertytypeFrom.GongCopyBasicFields(surfacepatcharraypropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygonpatchtype := range surfacepatcharraypropertytypeFrom.PolygonPatch {
		surfacepatcharraypropertytypeTo.PolygonPatch = append(surfacepatcharraypropertytypeTo.PolygonPatch, GongCopyBranchPolygonPatchType(mapOrigCopy, _polygonpatchtype))
	}

	return
}

func GongCopyBranchSurfaceType(mapOrigCopy map[any]any, surfacetypeFrom *SurfaceType) (surfacetypeTo *SurfaceType) {
	var alreadyCopied bool
	surfacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, surfacetypeFrom)
	if alreadyCopied {
		return
	}
	surfacetypeFrom.GongCopyBasicFields(surfacetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if surfacetypeFrom.Identifier != nil {
		surfacetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, surfacetypeFrom.Identifier)
	}
	if surfacetypeFrom.Patches != nil {
		surfacetypeTo.Patches = GongCopyBranchSurfacePatchArrayPropertyType(mapOrigCopy, surfacetypeFrom.Patches)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range surfacetypeFrom.NameXSD {
		surfacetypeTo.NameXSD = append(surfacetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimeInstantType(mapOrigCopy map[any]any, timeinstanttypeFrom *TimeInstantType) (timeinstanttypeTo *TimeInstantType) {
	var alreadyCopied bool
	timeinstanttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeinstanttypeFrom)
	if alreadyCopied {
		return
	}
	timeinstanttypeFrom.GongCopyBasicFields(timeinstanttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttypeFrom.Identifier != nil {
		timeinstanttypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, timeinstanttypeFrom.Identifier)
	}
	if timeinstanttypeFrom.TimePosition != nil {
		timeinstanttypeTo.TimePosition = GongCopyBranchTimePositionType(mapOrigCopy, timeinstanttypeFrom.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttypeFrom.NameXSD {
		timeinstanttypeTo.NameXSD = append(timeinstanttypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimePeriodType(mapOrigCopy map[any]any, timeperiodtypeFrom *TimePeriodType) (timeperiodtypeTo *TimePeriodType) {
	var alreadyCopied bool
	timeperiodtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeperiodtypeFrom)
	if alreadyCopied {
		return
	}
	timeperiodtypeFrom.GongCopyBasicFields(timeperiodtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtypeFrom.Identifier != nil {
		timeperiodtypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, timeperiodtypeFrom.Identifier)
	}
	if timeperiodtypeFrom.BeginPosition != nil {
		timeperiodtypeTo.BeginPosition = GongCopyBranchTimePositionType(mapOrigCopy, timeperiodtypeFrom.BeginPosition)
	}
	if timeperiodtypeFrom.EndPosition != nil {
		timeperiodtypeTo.EndPosition = GongCopyBranchTimePositionType(mapOrigCopy, timeperiodtypeFrom.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtypeFrom.NameXSD {
		timeperiodtypeTo.NameXSD = append(timeperiodtypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimePositionType(mapOrigCopy map[any]any, timepositiontypeFrom *TimePositionType) (timepositiontypeTo *TimePositionType) {
	var alreadyCopied bool
	timepositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timepositiontypeFrom)
	if alreadyCopied {
		return
	}
	timepositiontypeFrom.GongCopyBasicFields(timepositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimePrimitivePropertyType(mapOrigCopy map[any]any, timeprimitivepropertytypeFrom *TimePrimitivePropertyType) (timeprimitivepropertytypeTo *TimePrimitivePropertyType) {
	var alreadyCopied bool
	timeprimitivepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeprimitivepropertytypeFrom)
	if alreadyCopied {
		return
	}
	timeprimitivepropertytypeFrom.GongCopyBasicFields(timeprimitivepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytypeFrom.TimeInstant != nil {
		timeprimitivepropertytypeTo.TimeInstant = GongCopyBranchTimeInstantType(mapOrigCopy, timeprimitivepropertytypeFrom.TimeInstant)
	}
	if timeprimitivepropertytypeFrom.TimePeriod != nil {
		timeprimitivepropertytypeTo.TimePeriod = GongCopyBranchTimePeriodType(mapOrigCopy, timeprimitivepropertytypeFrom.TimePeriod)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractcurvesegmenttype) {
		return
	}

	abstractcurvesegmenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractcurvetype *AbstractCurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractcurvetype) {
		return
	}

	abstractcurvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractcurvetype.Identifier != nil {
		stage.UnstageBranch(abstractcurvetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractcurvetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractfeaturemembertype) {
		return
	}

	abstractfeaturemembertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractfeaturetype *AbstractFeatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractfeaturetype) {
		return
	}

	abstractfeaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetype.Identifier != nil {
		stage.UnstageBranch(abstractfeaturetype.Identifier)
	}
	if abstractfeaturetype.BoundedBy != nil {
		stage.UnstageBranch(abstractfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgmltype *AbstractGMLType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgmltype) {
		return
	}

	abstractgmltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltype.Identifier != nil {
		stage.UnstageBranch(abstractgmltype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgeometricprimitivetype) {
		return
	}

	abstractgeometricprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstractgeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgeometrytype *AbstractGeometryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgeometrytype) {
		return
	}

	abstractgeometrytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytype.Identifier != nil {
		stage.UnstageBranch(abstractgeometrytype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractmetadatapropertytype) {
		return
	}

	abstractmetadatapropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractringpropertytype *AbstractRingPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractringpropertytype) {
		return
	}

	abstractringpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractringpropertytype.Ring != nil {
		stage.UnstageBranch(abstractringpropertytype.Ring)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractringtype *AbstractRingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractringtype) {
		return
	}

	abstractringtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractsurfacepatchtype) {
		return
	}

	abstractsurfacepatchtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractsurfacetype *AbstractSurfaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractsurfacetype) {
		return
	}

	abstractsurfacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractsurfacetype.Identifier != nil {
		stage.UnstageBranch(abstractsurfacetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractsurfacetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimegeometricprimitivetype) {
		return
	}

	abstracttimegeometricprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstracttimegeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeobjecttype) {
		return
	}

	abstracttimeobjecttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttype.Identifier != nil {
		stage.UnstageBranch(abstracttimeobjecttype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeprimitivetype) {
		return
	}

	abstracttimeprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstracttimeprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeslicetype *AbstractTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeslicetype) {
		return
	}

	abstracttimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetype.Identifier != nil {
		stage.UnstageBranch(abstracttimeslicetype.Identifier)
	}
	if abstracttimeslicetype.ValidTime != nil {
		stage.UnstageBranch(abstracttimeslicetype.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (angletype *AngleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(angletype) {
		return
	}

	angletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcbycenterpointtype *ArcByCenterPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arcbycenterpointtype) {
		return
	}

	arcbycenterpointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arcbycenterpointtype.Pos != nil {
		stage.UnstageBranch(arcbycenterpointtype.Pos)
	}
	if arcbycenterpointtype.PointProperty != nil {
		stage.UnstageBranch(arcbycenterpointtype.PointProperty)
	}
	if arcbycenterpointtype.PosList != nil {
		stage.UnstageBranch(arcbycenterpointtype.PosList)
	}
	if arcbycenterpointtype.Radius != nil {
		stage.UnstageBranch(arcbycenterpointtype.Radius)
	}
	if arcbycenterpointtype.StartAngle != nil {
		stage.UnstageBranch(arcbycenterpointtype.StartAngle)
	}
	if arcbycenterpointtype.EndAngle != nil {
		stage.UnstageBranch(arcbycenterpointtype.EndAngle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcstringtype *ArcStringType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arcstringtype) {
		return
	}

	arcstringtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arcstringtype.PosList != nil {
		stage.UnstageBranch(arcstringtype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range arcstringtype.Pos {
		stage.UnstageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range arcstringtype.PointProperty {
		stage.UnstageBranch(_pointpropertytype)
	}

}

func (arctype *ArcType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arctype) {
		return
	}

	arctype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arctype.Pos != nil {
		stage.UnstageBranch(arctype.Pos)
	}
	if arctype.PointProperty != nil {
		stage.UnstageBranch(arctype.PointProperty)
	}
	if arctype.PosList != nil {
		stage.UnstageBranch(arctype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (boundingshapetype *BoundingShapeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(boundingshapetype) {
		return
	}

	boundingshapetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetype.Envelope != nil {
		stage.UnstageBranch(boundingshapetype.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlebycenterpointtype *CircleByCenterPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(circlebycenterpointtype) {
		return
	}

	circlebycenterpointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circlebycenterpointtype.Pos != nil {
		stage.UnstageBranch(circlebycenterpointtype.Pos)
	}
	if circlebycenterpointtype.PointProperty != nil {
		stage.UnstageBranch(circlebycenterpointtype.PointProperty)
	}
	if circlebycenterpointtype.PosList != nil {
		stage.UnstageBranch(circlebycenterpointtype.PosList)
	}
	if circlebycenterpointtype.Radius != nil {
		stage.UnstageBranch(circlebycenterpointtype.Radius)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circletype *CircleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(circletype) {
		return
	}

	circletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if circletype.PosList != nil {
		stage.UnstageBranch(circletype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range circletype.Pos {
		stage.UnstageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range circletype.PointProperty {
		stage.UnstageBranch(_pointpropertytype)
	}

}

func (codetype *CodeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetype) {
		return
	}

	codetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (compositecurvetype *CompositeCurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(compositecurvetype) {
		return
	}

	compositecurvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compositecurvetype.Identifier != nil {
		stage.UnstageBranch(compositecurvetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range compositecurvetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}
	for _, _curvepropertytype := range compositecurvetype.CurveMember {
		stage.UnstageBranch(_curvepropertytype)
	}

}

func (curvepropertytype *CurvePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curvepropertytype) {
		return
	}

	curvepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytype.Curve != nil {
		stage.UnstageBranch(curvepropertytype.Curve)
	}
	if curvepropertytype.OrientableCurve != nil {
		stage.UnstageBranch(curvepropertytype.OrientableCurve)
	}
	if curvepropertytype.CompositeCurve != nil {
		stage.UnstageBranch(curvepropertytype.CompositeCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curvesegmentarraypropertytype) {
		return
	}

	curvesegmentarraypropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linestringsegmenttype := range curvesegmentarraypropertytype.LineStringSegment {
		stage.UnstageBranch(_linestringsegmenttype)
	}
	for _, _arcstringtype := range curvesegmentarraypropertytype.ArcString {
		stage.UnstageBranch(_arcstringtype)
	}
	for _, _arctype := range curvesegmentarraypropertytype.Arc {
		stage.UnstageBranch(_arctype)
	}
	for _, _circletype := range curvesegmentarraypropertytype.Circle {
		stage.UnstageBranch(_circletype)
	}
	for _, _arcbycenterpointtype := range curvesegmentarraypropertytype.ArcByCenterPoint {
		stage.UnstageBranch(_arcbycenterpointtype)
	}
	for _, _circlebycenterpointtype := range curvesegmentarraypropertytype.CircleByCenterPoint {
		stage.UnstageBranch(_circlebycenterpointtype)
	}
	for _, _geodesicstringtype := range curvesegmentarraypropertytype.GeodesicString {
		stage.UnstageBranch(_geodesicstringtype)
	}
	for _, _geodesictype := range curvesegmentarraypropertytype.Geodesic {
		stage.UnstageBranch(_geodesictype)
	}

}

func (curvetype *CurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curvetype) {
		return
	}

	curvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvetype.Identifier != nil {
		stage.UnstageBranch(curvetype.Identifier)
	}
	if curvetype.Segments != nil {
		stage.UnstageBranch(curvetype.Segments)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range curvetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (directpositionlisttype *DirectPositionListType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(directpositionlisttype) {
		return
	}

	directpositionlisttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directpositiontype *DirectPositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(directpositiontype) {
		return
	}

	directpositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamicfeaturetype *DynamicFeatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dynamicfeaturetype) {
		return
	}

	dynamicfeaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetype.Identifier != nil {
		stage.UnstageBranch(dynamicfeaturetype.Identifier)
	}
	if dynamicfeaturetype.BoundedBy != nil {
		stage.UnstageBranch(dynamicfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (envelopetype *EnvelopeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(envelopetype) {
		return
	}

	envelopetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if envelopetype.LowerCorner != nil {
		stage.UnstageBranch(envelopetype.LowerCorner)
	}
	if envelopetype.UpperCorner != nil {
		stage.UnstageBranch(envelopetype.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geodesicstringtype *GeodesicStringType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geodesicstringtype) {
		return
	}

	geodesicstringtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geodesicstringtype.PosList != nil {
		stage.UnstageBranch(geodesicstringtype.PosList)
	}
	if geodesicstringtype.Pos != nil {
		stage.UnstageBranch(geodesicstringtype.Pos)
	}
	if geodesicstringtype.PointProperty != nil {
		stage.UnstageBranch(geodesicstringtype.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geodesictype *GeodesicType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geodesictype) {
		return
	}

	geodesictype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geodesictype.PosList != nil {
		stage.UnstageBranch(geodesictype.PosList)
	}
	if geodesictype.Pos != nil {
		stage.UnstageBranch(geodesictype.Pos)
	}
	if geodesictype.PointProperty != nil {
		stage.UnstageBranch(geodesictype.PointProperty)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lengthtype *LengthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lengthtype) {
		return
	}

	lengthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linestringsegmenttype *LineStringSegmentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linestringsegmenttype) {
		return
	}

	linestringsegmenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linestringsegmenttype.PosList != nil {
		stage.UnstageBranch(linestringsegmenttype.PosList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directpositiontype := range linestringsegmenttype.Pos {
		stage.UnstageBranch(_directpositiontype)
	}
	for _, _pointpropertytype := range linestringsegmenttype.PointProperty {
		stage.UnstageBranch(_pointpropertytype)
	}

}

func (measuretype *MeasureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measuretype) {
		return
	}

	measuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (orientablecurvetype *OrientableCurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(orientablecurvetype) {
		return
	}

	orientablecurvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if orientablecurvetype.Identifier != nil {
		stage.UnstageBranch(orientablecurvetype.Identifier)
	}
	if orientablecurvetype.BaseCurve != nil {
		stage.UnstageBranch(orientablecurvetype.BaseCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range orientablecurvetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (pointpropertytype *PointPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pointpropertytype) {
		return
	}

	pointpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointpropertytype.Point != nil {
		stage.UnstageBranch(pointpropertytype.Point)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pointtype) {
		return
	}

	pointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointtype.Identifier != nil {
		stage.UnstageBranch(pointtype.Identifier)
	}
	if pointtype.Pos != nil {
		stage.UnstageBranch(pointtype.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (polygonpatchtype *PolygonPatchType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(polygonpatchtype) {
		return
	}

	polygonpatchtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if polygonpatchtype.Exterior != nil {
		stage.UnstageBranch(polygonpatchtype.Exterior)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ringtype *RingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ringtype) {
		return
	}

	ringtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _curvepropertytype := range ringtype.CurveMember {
		stage.UnstageBranch(_curvepropertytype)
	}

}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(surfacepatcharraypropertytype) {
		return
	}

	surfacepatcharraypropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _polygonpatchtype := range surfacepatcharraypropertytype.PolygonPatch {
		stage.UnstageBranch(_polygonpatchtype)
	}

}

func (surfacetype *SurfaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(surfacetype) {
		return
	}

	surfacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if surfacetype.Identifier != nil {
		stage.UnstageBranch(surfacetype.Identifier)
	}
	if surfacetype.Patches != nil {
		stage.UnstageBranch(surfacetype.Patches)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range surfacetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timeinstanttype *TimeInstantType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeinstanttype) {
		return
	}

	timeinstanttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttype.Identifier != nil {
		stage.UnstageBranch(timeinstanttype.Identifier)
	}
	if timeinstanttype.TimePosition != nil {
		stage.UnstageBranch(timeinstanttype.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timeperiodtype *TimePeriodType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeperiodtype) {
		return
	}

	timeperiodtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtype.Identifier != nil {
		stage.UnstageBranch(timeperiodtype.Identifier)
	}
	if timeperiodtype.BeginPosition != nil {
		stage.UnstageBranch(timeperiodtype.BeginPosition)
	}
	if timeperiodtype.EndPosition != nil {
		stage.UnstageBranch(timeperiodtype.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timepositiontype *TimePositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timepositiontype) {
		return
	}

	timepositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeprimitivepropertytype) {
		return
	}

	timeprimitivepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytype.TimeInstant != nil {
		stage.UnstageBranch(timeprimitivepropertytype.TimeInstant)
	}
	if timeprimitivepropertytype.TimePeriod != nil {
		stage.UnstageBranch(timeprimitivepropertytype.TimePeriod)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AbstractCurveSegmentType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractCurveSegmentType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractCurveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractCurveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractFeatureMemberType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractFeatureMemberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractFeatureType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractFeatureType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BoundedBy, stage.BoundingShapeTypes_reference, instance.BoundedBy)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGMLType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGMLType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGeometricPrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGeometricPrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGeometryType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGeometryType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractMetadataPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractMetadataPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractRingPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractRingPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Ring, stage.RingTypes_reference, instance.Ring)
	// insertion point for slice of pointers field
}

func (reference *AbstractRingType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractRingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractSurfacePatchType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractSurfacePatchType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractSurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractSurfaceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimeGeometricPrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeGeometricPrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimeObjectType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeObjectType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimePrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimePrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.ValidTime, stage.TimePrimitivePropertyTypes_reference, instance.ValidTime)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AngleType) GongReconstructPointersFromReferences(stage *Stage, instance *AngleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ArcByCenterPointType) GongReconstructPointersFromReferences(stage *Stage, instance *ArcByCenterPointType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructPointer(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	__gong__reconstructPointer(&reference.Radius, stage.LengthTypes_reference, instance.Radius)
	__gong__reconstructPointer(&reference.StartAngle, stage.AngleTypes_reference, instance.StartAngle)
	__gong__reconstructPointer(&reference.EndAngle, stage.AngleTypes_reference, instance.EndAngle)
	// insertion point for slice of pointers field
}

func (reference *ArcStringType) GongReconstructPointersFromReferences(stage *Stage, instance *ArcStringType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
}

func (reference *ArcType) GongReconstructPointersFromReferences(stage *Stage, instance *ArcType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructPointer(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	// insertion point for slice of pointers field
}

func (reference *BoundingShapeType) GongReconstructPointersFromReferences(stage *Stage, instance *BoundingShapeType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Envelope, stage.EnvelopeTypes_reference, instance.Envelope)
	// insertion point for slice of pointers field
}

func (reference *CircleByCenterPointType) GongReconstructPointersFromReferences(stage *Stage, instance *CircleByCenterPointType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructPointer(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	__gong__reconstructPointer(&reference.Radius, stage.LengthTypes_reference, instance.Radius)
	// insertion point for slice of pointers field
}

func (reference *CircleType) GongReconstructPointersFromReferences(stage *Stage, instance *CircleType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
}

func (reference *CodeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWithAuthorityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CompositeCurveType) GongReconstructPointersFromReferences(stage *Stage, instance *CompositeCurveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
	__gong__reconstructSliceOfPointersFromReferences(&reference.CurveMember, stage.CurvePropertyTypes_reference, instance.CurveMember)
}

func (reference *CurvePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *CurvePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Curve, stage.CurveTypes_reference, instance.Curve)
	__gong__reconstructPointer(&reference.OrientableCurve, stage.OrientableCurveTypes_reference, instance.OrientableCurve)
	__gong__reconstructPointer(&reference.CompositeCurve, stage.CompositeCurveTypes_reference, instance.CompositeCurve)
	// insertion point for slice of pointers field
}

func (reference *CurveSegmentArrayPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *CurveSegmentArrayPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.LineStringSegment, stage.LineStringSegmentTypes_reference, instance.LineStringSegment)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ArcString, stage.ArcStringTypes_reference, instance.ArcString)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arc, stage.ArcTypes_reference, instance.Arc)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Circle, stage.CircleTypes_reference, instance.Circle)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ArcByCenterPoint, stage.ArcByCenterPointTypes_reference, instance.ArcByCenterPoint)
	__gong__reconstructSliceOfPointersFromReferences(&reference.CircleByCenterPoint, stage.CircleByCenterPointTypes_reference, instance.CircleByCenterPoint)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeodesicString, stage.GeodesicStringTypes_reference, instance.GeodesicString)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Geodesic, stage.GeodesicTypes_reference, instance.Geodesic)
}

func (reference *CurveType) GongReconstructPointersFromReferences(stage *Stage, instance *CurveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.Segments, stage.CurveSegmentArrayPropertyTypes_reference, instance.Segments)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *DirectPositionListType) GongReconstructPointersFromReferences(stage *Stage, instance *DirectPositionListType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DirectPositionType) GongReconstructPointersFromReferences(stage *Stage, instance *DirectPositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DynamicFeatureType) GongReconstructPointersFromReferences(stage *Stage, instance *DynamicFeatureType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BoundedBy, stage.BoundingShapeTypes_reference, instance.BoundedBy)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *EnvelopeType) GongReconstructPointersFromReferences(stage *Stage, instance *EnvelopeType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LowerCorner, stage.DirectPositionTypes_reference, instance.LowerCorner)
	__gong__reconstructPointer(&reference.UpperCorner, stage.DirectPositionTypes_reference, instance.UpperCorner)
	// insertion point for slice of pointers field
}

func (reference *GeodesicStringType) GongReconstructPointersFromReferences(stage *Stage, instance *GeodesicStringType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructPointer(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
	// insertion point for slice of pointers field
}

func (reference *GeodesicType) GongReconstructPointersFromReferences(stage *Stage, instance *GeodesicType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructPointer(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
	// insertion point for slice of pointers field
}

func (reference *LengthType) GongReconstructPointersFromReferences(stage *Stage, instance *LengthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LineStringSegmentType) GongReconstructPointersFromReferences(stage *Stage, instance *LineStringSegmentType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PosList, stage.DirectPositionListTypes_reference, instance.PosList)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PointProperty, stage.PointPropertyTypes_reference, instance.PointProperty)
}

func (reference *MeasureType) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OrientableCurveType) GongReconstructPointersFromReferences(stage *Stage, instance *OrientableCurveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BaseCurve, stage.CurvePropertyTypes_reference, instance.BaseCurve)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *PointPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *PointPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Point, stage.PointTypes_reference, instance.Point)
	// insertion point for slice of pointers field
}

func (reference *PointType) GongReconstructPointersFromReferences(stage *Stage, instance *PointType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *PolygonPatchType) GongReconstructPointersFromReferences(stage *Stage, instance *PolygonPatchType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Exterior, stage.AbstractRingPropertyTypes_reference, instance.Exterior)
	// insertion point for slice of pointers field
}

func (reference *RingType) GongReconstructPointersFromReferences(stage *Stage, instance *RingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.CurveMember, stage.CurvePropertyTypes_reference, instance.CurveMember)
}

func (reference *SurfacePatchArrayPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *SurfacePatchArrayPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.PolygonPatch, stage.PolygonPatchTypes_reference, instance.PolygonPatch)
}

func (reference *SurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *SurfaceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.Patches, stage.SurfacePatchArrayPropertyTypes_reference, instance.Patches)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimeInstantType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeInstantType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.TimePosition, stage.TimePositionTypes_reference, instance.TimePosition)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimePeriodType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePeriodType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BeginPosition, stage.TimePositionTypes_reference, instance.BeginPosition)
	__gong__reconstructPointer(&reference.EndPosition, stage.TimePositionTypes_reference, instance.EndPosition)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimePositionType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TimePrimitivePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePrimitivePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TimeInstant, stage.TimeInstantTypes_reference, instance.TimeInstant)
	__gong__reconstructPointer(&reference.TimePeriod, stage.TimePeriodTypes_reference, instance.TimePeriod)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractCurveSegmentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractCurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractFeatureMemberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractFeatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BoundedBy, stage.BoundingShapeTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGMLType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGeometricPrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGeometryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractMetadataPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractRingPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Ring, stage.RingTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AbstractRingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractSurfacePatchType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractSurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimeGeometricPrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimeObjectType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimePrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ValidTime, stage.TimePrimitivePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AngleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ArcByCenterPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointProperty, stage.PointPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Radius, stage.LengthTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.StartAngle, stage.AngleTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.EndAngle, stage.AngleTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ArcStringType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PointProperty, stage.PointPropertyTypes_instance)
}

func (reference *ArcType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointProperty, stage.PointPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *BoundingShapeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Envelope, stage.EnvelopeTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *CircleByCenterPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointProperty, stage.PointPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Radius, stage.LengthTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *CircleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PointProperty, stage.PointPropertyTypes_instance)
}

func (reference *CodeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CompositeCurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.CurveMember, stage.CurvePropertyTypes_instance)
}

func (reference *CurvePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Curve, stage.CurveTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.OrientableCurve, stage.OrientableCurveTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.CompositeCurve, stage.CompositeCurveTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *CurveSegmentArrayPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.LineStringSegment, stage.LineStringSegmentTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ArcString, stage.ArcStringTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arc, stage.ArcTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Circle, stage.CircleTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ArcByCenterPoint, stage.ArcByCenterPointTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.CircleByCenterPoint, stage.CircleByCenterPointTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeodesicString, stage.GeodesicStringTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Geodesic, stage.GeodesicTypes_instance)
}

func (reference *CurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Segments, stage.CurveSegmentArrayPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *DirectPositionListType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DirectPositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DynamicFeatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BoundedBy, stage.BoundingShapeTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *EnvelopeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LowerCorner, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.UpperCorner, stage.DirectPositionTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *GeodesicStringType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointProperty, stage.PointPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *GeodesicType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointProperty, stage.PointPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *LengthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LineStringSegmentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PosList, stage.DirectPositionListTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pos, stage.DirectPositionTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PointProperty, stage.PointPropertyTypes_instance)
}

func (reference *MeasureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OrientableCurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BaseCurve, stage.CurvePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *PointPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Point, stage.PointTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *PointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *PolygonPatchType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Exterior, stage.AbstractRingPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *RingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.CurveMember, stage.CurvePropertyTypes_instance)
}

func (reference *SurfacePatchArrayPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.PolygonPatch, stage.PolygonPatchTypes_instance)
}

func (reference *SurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Patches, stage.SurfacePatchArrayPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimeInstantType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.TimePosition, stage.TimePositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimePeriodType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BeginPosition, stage.TimePositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.EndPosition, stage.TimePositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimePositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TimePrimitivePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TimeInstant, stage.TimeInstantTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.TimePeriod, stage.TimePeriodTypes_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongDiff(stage *Stage, abstractcurvesegmenttypeOther *AbstractCurveSegmentType) (diffs []string) {
	// insertion point for field diffs
	if abstractcurvesegmenttype.Name != abstractcurvesegmenttypeOther.Name {
		diffs = append(diffs, abstractcurvesegmenttype.GongMarshallField(stage, "Name"))
	}
	if abstractcurvesegmenttype.NumDerivativesAtStart != abstractcurvesegmenttypeOther.NumDerivativesAtStart {
		diffs = append(diffs, abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if abstractcurvesegmenttype.NumDerivativesAtEnd != abstractcurvesegmenttypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if abstractcurvesegmenttype.NumDerivativeInterior != abstractcurvesegmenttypeOther.NumDerivativeInterior {
		diffs = append(diffs, abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativeInterior"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractcurvetype *AbstractCurveType) GongDiff(stage *Stage, abstractcurvetypeOther *AbstractCurveType) (diffs []string) {
	// insertion point for field diffs
	if abstractcurvetype.Name != abstractcurvetypeOther.Name {
		diffs = append(diffs, abstractcurvetype.GongMarshallField(stage, "Name"))
	}
	if abstractcurvetype.Id != abstractcurvetypeOther.Id {
		diffs = append(diffs, abstractcurvetype.GongMarshallField(stage, "Id"))
	}
	if abstractcurvetype.Description != abstractcurvetypeOther.Description {
		diffs = append(diffs, abstractcurvetype.GongMarshallField(stage, "Description"))
	}
	if abstractcurvetype.Identifier != abstractcurvetypeOther.Identifier {
		diffs = append(diffs, abstractcurvetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractcurvetype, "NameXSD", abstractcurvetypeOther.NameXSD, abstractcurvetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractcurvetype.SrsName != abstractcurvetypeOther.SrsName {
		diffs = append(diffs, abstractcurvetype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongDiff(stage *Stage, abstractfeaturemembertypeOther *AbstractFeatureMemberType) (diffs []string) {
	// insertion point for field diffs
	if abstractfeaturemembertype.Name != abstractfeaturemembertypeOther.Name {
		diffs = append(diffs, abstractfeaturemembertype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractfeaturetype *AbstractFeatureType) GongDiff(stage *Stage, abstractfeaturetypeOther *AbstractFeatureType) (diffs []string) {
	// insertion point for field diffs
	if abstractfeaturetype.Name != abstractfeaturetypeOther.Name {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Name"))
	}
	if abstractfeaturetype.Id != abstractfeaturetypeOther.Id {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Id"))
	}
	if abstractfeaturetype.Description != abstractfeaturetypeOther.Description {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Description"))
	}
	if abstractfeaturetype.Identifier != abstractfeaturetypeOther.Identifier {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractfeaturetype, "NameXSD", abstractfeaturetypeOther.NameXSD, abstractfeaturetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractfeaturetype.BoundedBy != abstractfeaturetypeOther.BoundedBy {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgmltype *AbstractGMLType) GongDiff(stage *Stage, abstractgmltypeOther *AbstractGMLType) (diffs []string) {
	// insertion point for field diffs
	if abstractgmltype.Name != abstractgmltypeOther.Name {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Name"))
	}
	if abstractgmltype.Id != abstractgmltypeOther.Id {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Id"))
	}
	if abstractgmltype.Description != abstractgmltypeOther.Description {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Description"))
	}
	if abstractgmltype.Identifier != abstractgmltypeOther.Identifier {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgmltype, "NameXSD", abstractgmltypeOther.NameXSD, abstractgmltype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongDiff(stage *Stage, abstractgeometricprimitivetypeOther *AbstractGeometricPrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstractgeometricprimitivetype.Name != abstractgeometricprimitivetypeOther.Name {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstractgeometricprimitivetype.Id != abstractgeometricprimitivetypeOther.Id {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstractgeometricprimitivetype.Description != abstractgeometricprimitivetypeOther.Description {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstractgeometricprimitivetype.Identifier != abstractgeometricprimitivetypeOther.Identifier {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgeometricprimitivetype, "NameXSD", abstractgeometricprimitivetypeOther.NameXSD, abstractgeometricprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractgeometricprimitivetype.SrsName != abstractgeometricprimitivetypeOther.SrsName {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgeometrytype *AbstractGeometryType) GongDiff(stage *Stage, abstractgeometrytypeOther *AbstractGeometryType) (diffs []string) {
	// insertion point for field diffs
	if abstractgeometrytype.Name != abstractgeometrytypeOther.Name {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Name"))
	}
	if abstractgeometrytype.Id != abstractgeometrytypeOther.Id {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Id"))
	}
	if abstractgeometrytype.Description != abstractgeometrytypeOther.Description {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Description"))
	}
	if abstractgeometrytype.Identifier != abstractgeometrytypeOther.Identifier {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgeometrytype, "NameXSD", abstractgeometrytypeOther.NameXSD, abstractgeometrytype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractgeometrytype.SrsName != abstractgeometrytypeOther.SrsName {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongDiff(stage *Stage, abstractmetadatapropertytypeOther *AbstractMetadataPropertyType) (diffs []string) {
	// insertion point for field diffs
	if abstractmetadatapropertytype.Name != abstractmetadatapropertytypeOther.Name {
		diffs = append(diffs, abstractmetadatapropertytype.GongMarshallField(stage, "Name"))
	}
	if abstractmetadatapropertytype.Owns != abstractmetadatapropertytypeOther.Owns {
		diffs = append(diffs, abstractmetadatapropertytype.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractringpropertytype *AbstractRingPropertyType) GongDiff(stage *Stage, abstractringpropertytypeOther *AbstractRingPropertyType) (diffs []string) {
	// insertion point for field diffs
	if abstractringpropertytype.Name != abstractringpropertytypeOther.Name {
		diffs = append(diffs, abstractringpropertytype.GongMarshallField(stage, "Name"))
	}
	if abstractringpropertytype.Ring != abstractringpropertytypeOther.Ring {
		diffs = append(diffs, abstractringpropertytype.GongMarshallField(stage, "Ring"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractringtype *AbstractRingType) GongDiff(stage *Stage, abstractringtypeOther *AbstractRingType) (diffs []string) {
	// insertion point for field diffs
	if abstractringtype.Name != abstractringtypeOther.Name {
		diffs = append(diffs, abstractringtype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongDiff(stage *Stage, abstractsurfacepatchtypeOther *AbstractSurfacePatchType) (diffs []string) {
	// insertion point for field diffs
	if abstractsurfacepatchtype.Name != abstractsurfacepatchtypeOther.Name {
		diffs = append(diffs, abstractsurfacepatchtype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractsurfacetype *AbstractSurfaceType) GongDiff(stage *Stage, abstractsurfacetypeOther *AbstractSurfaceType) (diffs []string) {
	// insertion point for field diffs
	if abstractsurfacetype.Name != abstractsurfacetypeOther.Name {
		diffs = append(diffs, abstractsurfacetype.GongMarshallField(stage, "Name"))
	}
	if abstractsurfacetype.Id != abstractsurfacetypeOther.Id {
		diffs = append(diffs, abstractsurfacetype.GongMarshallField(stage, "Id"))
	}
	if abstractsurfacetype.Description != abstractsurfacetypeOther.Description {
		diffs = append(diffs, abstractsurfacetype.GongMarshallField(stage, "Description"))
	}
	if abstractsurfacetype.Identifier != abstractsurfacetypeOther.Identifier {
		diffs = append(diffs, abstractsurfacetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractsurfacetype, "NameXSD", abstractsurfacetypeOther.NameXSD, abstractsurfacetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractsurfacetype.SrsName != abstractsurfacetypeOther.SrsName {
		diffs = append(diffs, abstractsurfacetype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongDiff(stage *Stage, abstracttimegeometricprimitivetypeOther *AbstractTimeGeometricPrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimegeometricprimitivetype.Name != abstracttimegeometricprimitivetypeOther.Name {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimegeometricprimitivetype.Id != abstracttimegeometricprimitivetypeOther.Id {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimegeometricprimitivetype.Description != abstracttimegeometricprimitivetypeOther.Description {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimegeometricprimitivetype.Identifier != abstracttimegeometricprimitivetypeOther.Identifier {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimegeometricprimitivetype, "NameXSD", abstracttimegeometricprimitivetypeOther.NameXSD, abstracttimegeometricprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstracttimegeometricprimitivetype.Frame != abstracttimegeometricprimitivetypeOther.Frame {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Frame"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeobjecttype *AbstractTimeObjectType) GongDiff(stage *Stage, abstracttimeobjecttypeOther *AbstractTimeObjectType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeobjecttype.Name != abstracttimeobjecttypeOther.Name {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeobjecttype.Id != abstracttimeobjecttypeOther.Id {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeobjecttype.Description != abstracttimeobjecttypeOther.Description {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeobjecttype.Identifier != abstracttimeobjecttypeOther.Identifier {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeobjecttype, "NameXSD", abstracttimeobjecttypeOther.NameXSD, abstracttimeobjecttype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongDiff(stage *Stage, abstracttimeprimitivetypeOther *AbstractTimePrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeprimitivetype.Name != abstracttimeprimitivetypeOther.Name {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeprimitivetype.Id != abstracttimeprimitivetypeOther.Id {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeprimitivetype.Description != abstracttimeprimitivetypeOther.Description {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeprimitivetype.Identifier != abstracttimeprimitivetypeOther.Identifier {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeprimitivetype, "NameXSD", abstracttimeprimitivetypeOther.NameXSD, abstracttimeprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeslicetype *AbstractTimeSliceType) GongDiff(stage *Stage, abstracttimeslicetypeOther *AbstractTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeslicetype.Name != abstracttimeslicetypeOther.Name {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeslicetype.Id != abstracttimeslicetypeOther.Id {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeslicetype.Description != abstracttimeslicetypeOther.Description {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeslicetype.Identifier != abstracttimeslicetypeOther.Identifier {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeslicetype, "NameXSD", abstracttimeslicetypeOther.NameXSD, abstracttimeslicetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstracttimeslicetype.ValidTime != abstracttimeslicetypeOther.ValidTime {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "ValidTime"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (angletype *AngleType) GongDiff(stage *Stage, angletypeOther *AngleType) (diffs []string) {
	// insertion point for field diffs
	if angletype.Name != angletypeOther.Name {
		diffs = append(diffs, angletype.GongMarshallField(stage, "Name"))
	}
	if angletype.EnclosedText != angletypeOther.EnclosedText {
		diffs = append(diffs, angletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arcbycenterpointtype *ArcByCenterPointType) GongDiff(stage *Stage, arcbycenterpointtypeOther *ArcByCenterPointType) (diffs []string) {
	// insertion point for field diffs
	if arcbycenterpointtype.Name != arcbycenterpointtypeOther.Name {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "Name"))
	}
	if arcbycenterpointtype.NumDerivativesAtStart != arcbycenterpointtypeOther.NumDerivativesAtStart {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if arcbycenterpointtype.NumDerivativesAtEnd != arcbycenterpointtypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if arcbycenterpointtype.NumDerivativeInterior != arcbycenterpointtypeOther.NumDerivativeInterior {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if arcbycenterpointtype.Interpolation != arcbycenterpointtypeOther.Interpolation {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "Interpolation"))
	}
	if arcbycenterpointtype.NumArc != arcbycenterpointtypeOther.NumArc {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "NumArc"))
	}
	if arcbycenterpointtype.Pos != arcbycenterpointtypeOther.Pos {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "Pos"))
	}
	if arcbycenterpointtype.PointProperty != arcbycenterpointtypeOther.PointProperty {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "PointProperty"))
	}
	if arcbycenterpointtype.PosList != arcbycenterpointtypeOther.PosList {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "PosList"))
	}
	if arcbycenterpointtype.Radius != arcbycenterpointtypeOther.Radius {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "Radius"))
	}
	if arcbycenterpointtype.StartAngle != arcbycenterpointtypeOther.StartAngle {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "StartAngle"))
	}
	if arcbycenterpointtype.EndAngle != arcbycenterpointtypeOther.EndAngle {
		diffs = append(diffs, arcbycenterpointtype.GongMarshallField(stage, "EndAngle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arcstringtype *ArcStringType) GongDiff(stage *Stage, arcstringtypeOther *ArcStringType) (diffs []string) {
	// insertion point for field diffs
	if arcstringtype.Name != arcstringtypeOther.Name {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "Name"))
	}
	if arcstringtype.NumDerivativesAtStart != arcstringtypeOther.NumDerivativesAtStart {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if arcstringtype.NumDerivativesAtEnd != arcstringtypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if arcstringtype.NumDerivativeInterior != arcstringtypeOther.NumDerivativeInterior {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if arcstringtype.Interpolation != arcstringtypeOther.Interpolation {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "Interpolation"))
	}
	if arcstringtype.NumArc != arcstringtypeOther.NumArc {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "NumArc"))
	}
	if ops := __gong__diffSliceOfPointers(stage, arcstringtype, "Pos", arcstringtypeOther.Pos, arcstringtype.Pos); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, arcstringtype, "PointProperty", arcstringtypeOther.PointProperty, arcstringtype.PointProperty); ops != "" {
		diffs = append(diffs, ops)
	}
	if arcstringtype.PosList != arcstringtypeOther.PosList {
		diffs = append(diffs, arcstringtype.GongMarshallField(stage, "PosList"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arctype *ArcType) GongDiff(stage *Stage, arctypeOther *ArcType) (diffs []string) {
	// insertion point for field diffs
	if arctype.Name != arctypeOther.Name {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Name"))
	}
	if arctype.NumDerivativesAtStart != arctypeOther.NumDerivativesAtStart {
		diffs = append(diffs, arctype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if arctype.NumDerivativesAtEnd != arctypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, arctype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if arctype.NumDerivativeInterior != arctypeOther.NumDerivativeInterior {
		diffs = append(diffs, arctype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if arctype.Interpolation != arctypeOther.Interpolation {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Interpolation"))
	}
	if arctype.NumArc != arctypeOther.NumArc {
		diffs = append(diffs, arctype.GongMarshallField(stage, "NumArc"))
	}
	if arctype.Pos != arctypeOther.Pos {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Pos"))
	}
	if arctype.PointProperty != arctypeOther.PointProperty {
		diffs = append(diffs, arctype.GongMarshallField(stage, "PointProperty"))
	}
	if arctype.PosList != arctypeOther.PosList {
		diffs = append(diffs, arctype.GongMarshallField(stage, "PosList"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (boundingshapetype *BoundingShapeType) GongDiff(stage *Stage, boundingshapetypeOther *BoundingShapeType) (diffs []string) {
	// insertion point for field diffs
	if boundingshapetype.Name != boundingshapetypeOther.Name {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "Name"))
	}
	if boundingshapetype.NilReason != boundingshapetypeOther.NilReason {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "NilReason"))
	}
	if boundingshapetype.Envelope != boundingshapetypeOther.Envelope {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "Envelope"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circlebycenterpointtype *CircleByCenterPointType) GongDiff(stage *Stage, circlebycenterpointtypeOther *CircleByCenterPointType) (diffs []string) {
	// insertion point for field diffs
	if circlebycenterpointtype.Name != circlebycenterpointtypeOther.Name {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "Name"))
	}
	if circlebycenterpointtype.NumDerivativesAtStart != circlebycenterpointtypeOther.NumDerivativesAtStart {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if circlebycenterpointtype.NumDerivativesAtEnd != circlebycenterpointtypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if circlebycenterpointtype.NumDerivativeInterior != circlebycenterpointtypeOther.NumDerivativeInterior {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if circlebycenterpointtype.Interpolation != circlebycenterpointtypeOther.Interpolation {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "Interpolation"))
	}
	if circlebycenterpointtype.NumArc != circlebycenterpointtypeOther.NumArc {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "NumArc"))
	}
	if circlebycenterpointtype.Pos != circlebycenterpointtypeOther.Pos {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "Pos"))
	}
	if circlebycenterpointtype.PointProperty != circlebycenterpointtypeOther.PointProperty {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "PointProperty"))
	}
	if circlebycenterpointtype.PosList != circlebycenterpointtypeOther.PosList {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "PosList"))
	}
	if circlebycenterpointtype.Radius != circlebycenterpointtypeOther.Radius {
		diffs = append(diffs, circlebycenterpointtype.GongMarshallField(stage, "Radius"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circletype *CircleType) GongDiff(stage *Stage, circletypeOther *CircleType) (diffs []string) {
	// insertion point for field diffs
	if circletype.Name != circletypeOther.Name {
		diffs = append(diffs, circletype.GongMarshallField(stage, "Name"))
	}
	if circletype.NumDerivativesAtStart != circletypeOther.NumDerivativesAtStart {
		diffs = append(diffs, circletype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if circletype.NumDerivativesAtEnd != circletypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, circletype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if circletype.NumDerivativeInterior != circletypeOther.NumDerivativeInterior {
		diffs = append(diffs, circletype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if circletype.Interpolation != circletypeOther.Interpolation {
		diffs = append(diffs, circletype.GongMarshallField(stage, "Interpolation"))
	}
	if circletype.NumArc != circletypeOther.NumArc {
		diffs = append(diffs, circletype.GongMarshallField(stage, "NumArc"))
	}
	if ops := __gong__diffSliceOfPointers(stage, circletype, "Pos", circletypeOther.Pos, circletype.Pos); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, circletype, "PointProperty", circletypeOther.PointProperty, circletype.PointProperty); ops != "" {
		diffs = append(diffs, ops)
	}
	if circletype.PosList != circletypeOther.PosList {
		diffs = append(diffs, circletype.GongMarshallField(stage, "PosList"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetype *CodeType) GongDiff(stage *Stage, codetypeOther *CodeType) (diffs []string) {
	// insertion point for field diffs
	if codetype.Name != codetypeOther.Name {
		diffs = append(diffs, codetype.GongMarshallField(stage, "Name"))
	}
	if codetype.CodeSpace != codetypeOther.CodeSpace {
		diffs = append(diffs, codetype.GongMarshallField(stage, "CodeSpace"))
	}
	if codetype.EnclosedText != codetypeOther.EnclosedText {
		diffs = append(diffs, codetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codewithauthoritytype *CodeWithAuthorityType) GongDiff(stage *Stage, codewithauthoritytypeOther *CodeWithAuthorityType) (diffs []string) {
	// insertion point for field diffs
	if codewithauthoritytype.Name != codewithauthoritytypeOther.Name {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "Name"))
	}
	if codewithauthoritytype.CodeSpace != codewithauthoritytypeOther.CodeSpace {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "CodeSpace"))
	}
	if codewithauthoritytype.EnclosedText != codewithauthoritytypeOther.EnclosedText {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (compositecurvetype *CompositeCurveType) GongDiff(stage *Stage, compositecurvetypeOther *CompositeCurveType) (diffs []string) {
	// insertion point for field diffs
	if compositecurvetype.Name != compositecurvetypeOther.Name {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "Name"))
	}
	if compositecurvetype.Id != compositecurvetypeOther.Id {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "Id"))
	}
	if compositecurvetype.Description != compositecurvetypeOther.Description {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "Description"))
	}
	if compositecurvetype.Identifier != compositecurvetypeOther.Identifier {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, compositecurvetype, "NameXSD", compositecurvetypeOther.NameXSD, compositecurvetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if compositecurvetype.SrsName != compositecurvetypeOther.SrsName {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "SrsName"))
	}
	if compositecurvetype.AggregationType != compositecurvetypeOther.AggregationType {
		diffs = append(diffs, compositecurvetype.GongMarshallField(stage, "AggregationType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, compositecurvetype, "CurveMember", compositecurvetypeOther.CurveMember, compositecurvetype.CurveMember); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curvepropertytype *CurvePropertyType) GongDiff(stage *Stage, curvepropertytypeOther *CurvePropertyType) (diffs []string) {
	// insertion point for field diffs
	if curvepropertytype.Name != curvepropertytypeOther.Name {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Name"))
	}
	if curvepropertytype.NilReason != curvepropertytypeOther.NilReason {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if curvepropertytype.Type_XSD != curvepropertytypeOther.Type_XSD {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if curvepropertytype.Href != curvepropertytypeOther.Href {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Href"))
	}
	if curvepropertytype.Role != curvepropertytypeOther.Role {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Role"))
	}
	if curvepropertytype.Arcrole != curvepropertytypeOther.Arcrole {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if curvepropertytype.Title != curvepropertytypeOther.Title {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Title"))
	}
	if curvepropertytype.Show != curvepropertytypeOther.Show {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Show"))
	}
	if curvepropertytype.Actuate != curvepropertytypeOther.Actuate {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Actuate"))
	}
	if curvepropertytype.Curve != curvepropertytypeOther.Curve {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Curve"))
	}
	if curvepropertytype.OrientableCurve != curvepropertytypeOther.OrientableCurve {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "OrientableCurve"))
	}
	if curvepropertytype.CompositeCurve != curvepropertytypeOther.CompositeCurve {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "CompositeCurve"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongDiff(stage *Stage, curvesegmentarraypropertytypeOther *CurveSegmentArrayPropertyType) (diffs []string) {
	// insertion point for field diffs
	if curvesegmentarraypropertytype.Name != curvesegmentarraypropertytypeOther.Name {
		diffs = append(diffs, curvesegmentarraypropertytype.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "LineStringSegment", curvesegmentarraypropertytypeOther.LineStringSegment, curvesegmentarraypropertytype.LineStringSegment); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "ArcString", curvesegmentarraypropertytypeOther.ArcString, curvesegmentarraypropertytype.ArcString); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "Arc", curvesegmentarraypropertytypeOther.Arc, curvesegmentarraypropertytype.Arc); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "Circle", curvesegmentarraypropertytypeOther.Circle, curvesegmentarraypropertytype.Circle); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "ArcByCenterPoint", curvesegmentarraypropertytypeOther.ArcByCenterPoint, curvesegmentarraypropertytype.ArcByCenterPoint); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "CircleByCenterPoint", curvesegmentarraypropertytypeOther.CircleByCenterPoint, curvesegmentarraypropertytype.CircleByCenterPoint); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "GeodesicString", curvesegmentarraypropertytypeOther.GeodesicString, curvesegmentarraypropertytype.GeodesicString); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, curvesegmentarraypropertytype, "Geodesic", curvesegmentarraypropertytypeOther.Geodesic, curvesegmentarraypropertytype.Geodesic); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curvetype *CurveType) GongDiff(stage *Stage, curvetypeOther *CurveType) (diffs []string) {
	// insertion point for field diffs
	if curvetype.Name != curvetypeOther.Name {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Name"))
	}
	if curvetype.Id != curvetypeOther.Id {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Id"))
	}
	if curvetype.Description != curvetypeOther.Description {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Description"))
	}
	if curvetype.Identifier != curvetypeOther.Identifier {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, curvetype, "NameXSD", curvetypeOther.NameXSD, curvetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if curvetype.SrsName != curvetypeOther.SrsName {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "SrsName"))
	}
	if curvetype.Segments != curvetypeOther.Segments {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Segments"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (directpositionlisttype *DirectPositionListType) GongDiff(stage *Stage, directpositionlisttypeOther *DirectPositionListType) (diffs []string) {
	// insertion point for field diffs
	if directpositionlisttype.Name != directpositionlisttypeOther.Name {
		diffs = append(diffs, directpositionlisttype.GongMarshallField(stage, "Name"))
	}
	if directpositionlisttype.SrsName != directpositionlisttypeOther.SrsName {
		diffs = append(diffs, directpositionlisttype.GongMarshallField(stage, "SrsName"))
	}
	if directpositionlisttype.EnclosedText != directpositionlisttypeOther.EnclosedText {
		diffs = append(diffs, directpositionlisttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (directpositiontype *DirectPositionType) GongDiff(stage *Stage, directpositiontypeOther *DirectPositionType) (diffs []string) {
	// insertion point for field diffs
	if directpositiontype.Name != directpositiontypeOther.Name {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "Name"))
	}
	if directpositiontype.SrsName != directpositiontypeOther.SrsName {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "SrsName"))
	}
	if directpositiontype.EnclosedText != directpositiontypeOther.EnclosedText {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dynamicfeaturetype *DynamicFeatureType) GongDiff(stage *Stage, dynamicfeaturetypeOther *DynamicFeatureType) (diffs []string) {
	// insertion point for field diffs
	if dynamicfeaturetype.Name != dynamicfeaturetypeOther.Name {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Name"))
	}
	if dynamicfeaturetype.Id != dynamicfeaturetypeOther.Id {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Id"))
	}
	if dynamicfeaturetype.Description != dynamicfeaturetypeOther.Description {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Description"))
	}
	if dynamicfeaturetype.Identifier != dynamicfeaturetypeOther.Identifier {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamicfeaturetype, "NameXSD", dynamicfeaturetypeOther.NameXSD, dynamicfeaturetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if dynamicfeaturetype.BoundedBy != dynamicfeaturetypeOther.BoundedBy {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (envelopetype *EnvelopeType) GongDiff(stage *Stage, envelopetypeOther *EnvelopeType) (diffs []string) {
	// insertion point for field diffs
	if envelopetype.Name != envelopetypeOther.Name {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "Name"))
	}
	if envelopetype.SrsName != envelopetypeOther.SrsName {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "SrsName"))
	}
	if envelopetype.LowerCorner != envelopetypeOther.LowerCorner {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "LowerCorner"))
	}
	if envelopetype.UpperCorner != envelopetypeOther.UpperCorner {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "UpperCorner"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geodesicstringtype *GeodesicStringType) GongDiff(stage *Stage, geodesicstringtypeOther *GeodesicStringType) (diffs []string) {
	// insertion point for field diffs
	if geodesicstringtype.Name != geodesicstringtypeOther.Name {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "Name"))
	}
	if geodesicstringtype.NumDerivativesAtStart != geodesicstringtypeOther.NumDerivativesAtStart {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if geodesicstringtype.NumDerivativesAtEnd != geodesicstringtypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if geodesicstringtype.NumDerivativeInterior != geodesicstringtypeOther.NumDerivativeInterior {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if geodesicstringtype.Interpolation != geodesicstringtypeOther.Interpolation {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "Interpolation"))
	}
	if geodesicstringtype.PosList != geodesicstringtypeOther.PosList {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "PosList"))
	}
	if geodesicstringtype.Pos != geodesicstringtypeOther.Pos {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "Pos"))
	}
	if geodesicstringtype.PointProperty != geodesicstringtypeOther.PointProperty {
		diffs = append(diffs, geodesicstringtype.GongMarshallField(stage, "PointProperty"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geodesictype *GeodesicType) GongDiff(stage *Stage, geodesictypeOther *GeodesicType) (diffs []string) {
	// insertion point for field diffs
	if geodesictype.Name != geodesictypeOther.Name {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "Name"))
	}
	if geodesictype.NumDerivativesAtStart != geodesictypeOther.NumDerivativesAtStart {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if geodesictype.NumDerivativesAtEnd != geodesictypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if geodesictype.NumDerivativeInterior != geodesictypeOther.NumDerivativeInterior {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if geodesictype.Interpolation != geodesictypeOther.Interpolation {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "Interpolation"))
	}
	if geodesictype.PosList != geodesictypeOther.PosList {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "PosList"))
	}
	if geodesictype.Pos != geodesictypeOther.Pos {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "Pos"))
	}
	if geodesictype.PointProperty != geodesictypeOther.PointProperty {
		diffs = append(diffs, geodesictype.GongMarshallField(stage, "PointProperty"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lengthtype *LengthType) GongDiff(stage *Stage, lengthtypeOther *LengthType) (diffs []string) {
	// insertion point for field diffs
	if lengthtype.Name != lengthtypeOther.Name {
		diffs = append(diffs, lengthtype.GongMarshallField(stage, "Name"))
	}
	if lengthtype.EnclosedText != lengthtypeOther.EnclosedText {
		diffs = append(diffs, lengthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linestringsegmenttype *LineStringSegmentType) GongDiff(stage *Stage, linestringsegmenttypeOther *LineStringSegmentType) (diffs []string) {
	// insertion point for field diffs
	if linestringsegmenttype.Name != linestringsegmenttypeOther.Name {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "Name"))
	}
	if linestringsegmenttype.NumDerivativesAtStart != linestringsegmenttypeOther.NumDerivativesAtStart {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "NumDerivativesAtStart"))
	}
	if linestringsegmenttype.NumDerivativesAtEnd != linestringsegmenttypeOther.NumDerivativesAtEnd {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "NumDerivativesAtEnd"))
	}
	if linestringsegmenttype.NumDerivativeInterior != linestringsegmenttypeOther.NumDerivativeInterior {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	if linestringsegmenttype.Interpolation != linestringsegmenttypeOther.Interpolation {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "Interpolation"))
	}
	if ops := __gong__diffSliceOfPointers(stage, linestringsegmenttype, "Pos", linestringsegmenttypeOther.Pos, linestringsegmenttype.Pos); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, linestringsegmenttype, "PointProperty", linestringsegmenttypeOther.PointProperty, linestringsegmenttype.PointProperty); ops != "" {
		diffs = append(diffs, ops)
	}
	if linestringsegmenttype.PosList != linestringsegmenttypeOther.PosList {
		diffs = append(diffs, linestringsegmenttype.GongMarshallField(stage, "PosList"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measuretype *MeasureType) GongDiff(stage *Stage, measuretypeOther *MeasureType) (diffs []string) {
	// insertion point for field diffs
	if measuretype.Name != measuretypeOther.Name {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "Name"))
	}
	if measuretype.Uom != measuretypeOther.Uom {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "Uom"))
	}
	if measuretype.EnclosedText != measuretypeOther.EnclosedText {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (orientablecurvetype *OrientableCurveType) GongDiff(stage *Stage, orientablecurvetypeOther *OrientableCurveType) (diffs []string) {
	// insertion point for field diffs
	if orientablecurvetype.Name != orientablecurvetypeOther.Name {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "Name"))
	}
	if orientablecurvetype.Id != orientablecurvetypeOther.Id {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "Id"))
	}
	if orientablecurvetype.Description != orientablecurvetypeOther.Description {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "Description"))
	}
	if orientablecurvetype.Identifier != orientablecurvetypeOther.Identifier {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, orientablecurvetype, "NameXSD", orientablecurvetypeOther.NameXSD, orientablecurvetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if orientablecurvetype.SrsName != orientablecurvetypeOther.SrsName {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "SrsName"))
	}
	if orientablecurvetype.Orientation != orientablecurvetypeOther.Orientation {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "Orientation"))
	}
	if orientablecurvetype.BaseCurve != orientablecurvetypeOther.BaseCurve {
		diffs = append(diffs, orientablecurvetype.GongMarshallField(stage, "BaseCurve"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointpropertytype *PointPropertyType) GongDiff(stage *Stage, pointpropertytypeOther *PointPropertyType) (diffs []string) {
	// insertion point for field diffs
	if pointpropertytype.Name != pointpropertytypeOther.Name {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Name"))
	}
	if pointpropertytype.NilReason != pointpropertytypeOther.NilReason {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if pointpropertytype.Type_XSD != pointpropertytypeOther.Type_XSD {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if pointpropertytype.Href != pointpropertytypeOther.Href {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Href"))
	}
	if pointpropertytype.Role != pointpropertytypeOther.Role {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Role"))
	}
	if pointpropertytype.Arcrole != pointpropertytypeOther.Arcrole {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if pointpropertytype.Title != pointpropertytypeOther.Title {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Title"))
	}
	if pointpropertytype.Show != pointpropertytypeOther.Show {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Show"))
	}
	if pointpropertytype.Actuate != pointpropertytypeOther.Actuate {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Actuate"))
	}
	if pointpropertytype.Point != pointpropertytypeOther.Point {
		diffs = append(diffs, pointpropertytype.GongMarshallField(stage, "Point"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointtype *PointType) GongDiff(stage *Stage, pointtypeOther *PointType) (diffs []string) {
	// insertion point for field diffs
	if pointtype.Name != pointtypeOther.Name {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Name"))
	}
	if pointtype.Id != pointtypeOther.Id {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Id"))
	}
	if pointtype.Description != pointtypeOther.Description {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Description"))
	}
	if pointtype.Identifier != pointtypeOther.Identifier {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, pointtype, "NameXSD", pointtypeOther.NameXSD, pointtype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if pointtype.SrsName != pointtypeOther.SrsName {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "SrsName"))
	}
	if pointtype.Pos != pointtypeOther.Pos {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Pos"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (polygonpatchtype *PolygonPatchType) GongDiff(stage *Stage, polygonpatchtypeOther *PolygonPatchType) (diffs []string) {
	// insertion point for field diffs
	if polygonpatchtype.Name != polygonpatchtypeOther.Name {
		diffs = append(diffs, polygonpatchtype.GongMarshallField(stage, "Name"))
	}
	if polygonpatchtype.Interpolation != polygonpatchtypeOther.Interpolation {
		diffs = append(diffs, polygonpatchtype.GongMarshallField(stage, "Interpolation"))
	}
	if polygonpatchtype.Exterior != polygonpatchtypeOther.Exterior {
		diffs = append(diffs, polygonpatchtype.GongMarshallField(stage, "Exterior"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ringtype *RingType) GongDiff(stage *Stage, ringtypeOther *RingType) (diffs []string) {
	// insertion point for field diffs
	if ringtype.Name != ringtypeOther.Name {
		diffs = append(diffs, ringtype.GongMarshallField(stage, "Name"))
	}
	if ringtype.AggregationType != ringtypeOther.AggregationType {
		diffs = append(diffs, ringtype.GongMarshallField(stage, "AggregationType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ringtype, "CurveMember", ringtypeOther.CurveMember, ringtype.CurveMember); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongDiff(stage *Stage, surfacepatcharraypropertytypeOther *SurfacePatchArrayPropertyType) (diffs []string) {
	// insertion point for field diffs
	if surfacepatcharraypropertytype.Name != surfacepatcharraypropertytypeOther.Name {
		diffs = append(diffs, surfacepatcharraypropertytype.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, surfacepatcharraypropertytype, "PolygonPatch", surfacepatcharraypropertytypeOther.PolygonPatch, surfacepatcharraypropertytype.PolygonPatch); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (surfacetype *SurfaceType) GongDiff(stage *Stage, surfacetypeOther *SurfaceType) (diffs []string) {
	// insertion point for field diffs
	if surfacetype.Name != surfacetypeOther.Name {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Name"))
	}
	if surfacetype.Id != surfacetypeOther.Id {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Id"))
	}
	if surfacetype.Description != surfacetypeOther.Description {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Description"))
	}
	if surfacetype.Identifier != surfacetypeOther.Identifier {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, surfacetype, "NameXSD", surfacetypeOther.NameXSD, surfacetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if surfacetype.SrsName != surfacetypeOther.SrsName {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "SrsName"))
	}
	if surfacetype.Patches != surfacetypeOther.Patches {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Patches"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeinstanttype *TimeInstantType) GongDiff(stage *Stage, timeinstanttypeOther *TimeInstantType) (diffs []string) {
	// insertion point for field diffs
	if timeinstanttype.Name != timeinstanttypeOther.Name {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Name"))
	}
	if timeinstanttype.Id != timeinstanttypeOther.Id {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Id"))
	}
	if timeinstanttype.Description != timeinstanttypeOther.Description {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Description"))
	}
	if timeinstanttype.Identifier != timeinstanttypeOther.Identifier {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, timeinstanttype, "NameXSD", timeinstanttypeOther.NameXSD, timeinstanttype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if timeinstanttype.Frame != timeinstanttypeOther.Frame {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Frame"))
	}
	if timeinstanttype.TimePosition != timeinstanttypeOther.TimePosition {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "TimePosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeperiodtype *TimePeriodType) GongDiff(stage *Stage, timeperiodtypeOther *TimePeriodType) (diffs []string) {
	// insertion point for field diffs
	if timeperiodtype.Name != timeperiodtypeOther.Name {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Name"))
	}
	if timeperiodtype.Id != timeperiodtypeOther.Id {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Id"))
	}
	if timeperiodtype.Description != timeperiodtypeOther.Description {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Description"))
	}
	if timeperiodtype.Identifier != timeperiodtypeOther.Identifier {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, timeperiodtype, "NameXSD", timeperiodtypeOther.NameXSD, timeperiodtype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if timeperiodtype.Frame != timeperiodtypeOther.Frame {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Frame"))
	}
	if timeperiodtype.BeginPosition != timeperiodtypeOther.BeginPosition {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "BeginPosition"))
	}
	if timeperiodtype.EndPosition != timeperiodtypeOther.EndPosition {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "EndPosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timepositiontype *TimePositionType) GongDiff(stage *Stage, timepositiontypeOther *TimePositionType) (diffs []string) {
	// insertion point for field diffs
	if timepositiontype.Name != timepositiontypeOther.Name {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "Name"))
	}
	if timepositiontype.Frame != timepositiontypeOther.Frame {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "Frame"))
	}
	if timepositiontype.IndeterminatePosition != timepositiontypeOther.IndeterminatePosition {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "IndeterminatePosition"))
	}
	if timepositiontype.EnclosedText != timepositiontypeOther.EnclosedText {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongDiff(stage *Stage, timeprimitivepropertytypeOther *TimePrimitivePropertyType) (diffs []string) {
	// insertion point for field diffs
	if timeprimitivepropertytype.Name != timeprimitivepropertytypeOther.Name {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Name"))
	}
	if timeprimitivepropertytype.NilReason != timeprimitivepropertytypeOther.NilReason {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if timeprimitivepropertytype.Type_XSD != timeprimitivepropertytypeOther.Type_XSD {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if timeprimitivepropertytype.Href != timeprimitivepropertytypeOther.Href {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Href"))
	}
	if timeprimitivepropertytype.Role != timeprimitivepropertytypeOther.Role {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Role"))
	}
	if timeprimitivepropertytype.Arcrole != timeprimitivepropertytypeOther.Arcrole {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if timeprimitivepropertytype.Title != timeprimitivepropertytypeOther.Title {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Title"))
	}
	if timeprimitivepropertytype.Show != timeprimitivepropertytypeOther.Show {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Show"))
	}
	if timeprimitivepropertytype.Actuate != timeprimitivepropertytypeOther.Actuate {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Actuate"))
	}
	if timeprimitivepropertytype.TimeInstant != timeprimitivepropertytypeOther.TimeInstant {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "TimeInstant"))
	}
	if timeprimitivepropertytype.TimePeriod != timeprimitivepropertytypeOther.TimePeriod {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "TimePeriod"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
