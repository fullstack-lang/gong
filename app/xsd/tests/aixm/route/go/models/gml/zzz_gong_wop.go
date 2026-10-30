// generated code - do not edit
package gml

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AbstractCurveSegmentType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int
}

func (from *AbstractCurveSegmentType) GongCopyBasicFields(to *AbstractCurveSegmentType) {
	// insertion point
	*to = *from
}

type AbstractCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractCurveType) GongCopyBasicFields(to *AbstractCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractFeatureMemberType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractFeatureMemberType) GongCopyBasicFields(to *AbstractFeatureMemberType) {
	// insertion point
	*to = *from
}

type AbstractFeatureType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractFeatureType) GongCopyBasicFields(to *AbstractFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractGMLType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractGMLType) GongCopyBasicFields(to *AbstractGMLType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractGeometricPrimitiveType) GongCopyBasicFields(to *AbstractGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractGeometryType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractGeometryType) GongCopyBasicFields(to *AbstractGeometryType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractMetadataPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *AbstractMetadataPropertyType) GongCopyBasicFields(to *AbstractMetadataPropertyType) {
	// insertion point
	*to = *from
}

type AbstractRingPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractRingPropertyType) GongCopyBasicFields(to *AbstractRingPropertyType) {
	// insertion point
	to.Name = from.Name
}

type AbstractRingType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractRingType) GongCopyBasicFields(to *AbstractRingType) {
	// insertion point
	*to = *from
}

type AbstractSurfacePatchType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractSurfacePatchType) GongCopyBasicFields(to *AbstractSurfacePatchType) {
	// insertion point
	*to = *from
}

type AbstractSurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractSurfaceType) GongCopyBasicFields(to *AbstractSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractTimeGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *AbstractTimeGeometricPrimitiveType) GongCopyBasicFields(to *AbstractTimeGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type AbstractTimeObjectType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimeObjectType) GongCopyBasicFields(to *AbstractTimeObjectType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractTimePrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimePrimitiveType) GongCopyBasicFields(to *AbstractTimePrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractTimeSliceType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimeSliceType) GongCopyBasicFields(to *AbstractTimeSliceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AngleType_WOP struct {
	// insertion point

	Name string

	EnclosedText string
}

func (from *AngleType) GongCopyBasicFields(to *AngleType) {
	// insertion point
	*to = *from
}

type ArcByCenterPointType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int
}

func (from *ArcByCenterPointType) GongCopyBasicFields(to *ArcByCenterPointType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
}

type ArcStringType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int
}

func (from *ArcStringType) GongCopyBasicFields(to *ArcStringType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
}

type ArcType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int
}

func (from *ArcType) GongCopyBasicFields(to *ArcType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
}

type BoundingShapeType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *BoundingShapeType) GongCopyBasicFields(to *BoundingShapeType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type CircleByCenterPointType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int
}

func (from *CircleByCenterPointType) GongCopyBasicFields(to *CircleByCenterPointType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
}

type CircleType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int
}

func (from *CircleType) GongCopyBasicFields(to *CircleType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
}

type CodeType_WOP struct {
	// insertion point

	Name string

	CodeSpace string

	EnclosedText string
}

func (from *CodeType) GongCopyBasicFields(to *CodeType) {
	// insertion point
	*to = *from
}

type CodeWithAuthorityType_WOP struct {
	// insertion point

	Name string

	CodeSpace string

	EnclosedText string
}

func (from *CodeWithAuthorityType) GongCopyBasicFields(to *CodeWithAuthorityType) {
	// insertion point
	*to = *from
}

type CompositeCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string

	AggregationType string
}

func (from *CompositeCurveType) GongCopyBasicFields(to *CompositeCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
	to.AggregationType = from.AggregationType
}

type CurvePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CurvePropertyType) GongCopyBasicFields(to *CurvePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CurveSegmentArrayPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *CurveSegmentArrayPropertyType) GongCopyBasicFields(to *CurveSegmentArrayPropertyType) {
	// insertion point
	to.Name = from.Name
}

type CurveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *CurveType) GongCopyBasicFields(to *CurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type DirectPositionListType_WOP struct {
	// insertion point

	Name string

	SrsName string

	EnclosedText string
}

func (from *DirectPositionListType) GongCopyBasicFields(to *DirectPositionListType) {
	// insertion point
	*to = *from
}

type DirectPositionType_WOP struct {
	// insertion point

	Name string

	SrsName string

	EnclosedText string
}

func (from *DirectPositionType) GongCopyBasicFields(to *DirectPositionType) {
	// insertion point
	*to = *from
}

type DynamicFeatureType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *DynamicFeatureType) GongCopyBasicFields(to *DynamicFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type EnvelopeType_WOP struct {
	// insertion point

	Name string

	SrsName string
}

func (from *EnvelopeType) GongCopyBasicFields(to *EnvelopeType) {
	// insertion point
	to.Name = from.Name
	to.SrsName = from.SrsName
}

type GeodesicStringType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType
}

func (from *GeodesicStringType) GongCopyBasicFields(to *GeodesicStringType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
}

type GeodesicType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType
}

func (from *GeodesicType) GongCopyBasicFields(to *GeodesicType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
}

type LengthType_WOP struct {
	// insertion point

	Name string

	EnclosedText string
}

func (from *LengthType) GongCopyBasicFields(to *LengthType) {
	// insertion point
	*to = *from
}

type LineStringSegmentType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType
}

func (from *LineStringSegmentType) GongCopyBasicFields(to *LineStringSegmentType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
}

type MeasureType_WOP struct {
	// insertion point

	Name string

	Uom string

	EnclosedText float64
}

func (from *MeasureType) GongCopyBasicFields(to *MeasureType) {
	// insertion point
	*to = *from
}

type OrientableCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string

	Orientation Enum_SignType
}

func (from *OrientableCurveType) GongCopyBasicFields(to *OrientableCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
	to.Orientation = from.Orientation
}

type PointPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PointPropertyType) GongCopyBasicFields(to *PointPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PointType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *PointType) GongCopyBasicFields(to *PointType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type PolygonPatchType_WOP struct {
	// insertion point

	Name string

	Interpolation Enum_SurfaceInterpolationType
}

func (from *PolygonPatchType) GongCopyBasicFields(to *PolygonPatchType) {
	// insertion point
	to.Name = from.Name
	to.Interpolation = from.Interpolation
}

type RingType_WOP struct {
	// insertion point

	Name string

	AggregationType string
}

func (from *RingType) GongCopyBasicFields(to *RingType) {
	// insertion point
	to.Name = from.Name
	to.AggregationType = from.AggregationType
}

type SurfacePatchArrayPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *SurfacePatchArrayPropertyType) GongCopyBasicFields(to *SurfacePatchArrayPropertyType) {
	// insertion point
	to.Name = from.Name
}

type SurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *SurfaceType) GongCopyBasicFields(to *SurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type TimeInstantType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *TimeInstantType) GongCopyBasicFields(to *TimeInstantType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type TimePeriodType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *TimePeriodType) GongCopyBasicFields(to *TimePeriodType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type TimePositionType_WOP struct {
	// insertion point

	Name string

	Frame string

	IndeterminatePosition Enum_TimeIndeterminateValueType

	EnclosedText string
}

func (from *TimePositionType) GongCopyBasicFields(to *TimePositionType) {
	// insertion point
	*to = *from
}

type TimePrimitivePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimePrimitivePropertyType) GongCopyBasicFields(to *TimePrimitivePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

// end of insertion point
