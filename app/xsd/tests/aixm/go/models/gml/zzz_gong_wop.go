// generated code - do not edit
package gml

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type A_Row_WOP struct {
	// insertion point

	Name string
}

func (from *A_Row) GongCopyBasicFields(to *A_Row) {
	// insertion point
	to.Name = from.Name
}

type A_controlPoint_WOP struct {
	// insertion point

	Name string
}

func (from *A_controlPoint) GongCopyBasicFields(to *A_controlPoint) {
	// insertion point
	to.Name = from.Name
}

type A_refLocation_WOP struct {
	// insertion point

	Name string
}

func (from *A_refLocation) GongCopyBasicFields(to *A_refLocation) {
	// insertion point
	to.Name = from.Name
}

type A_rows_WOP struct {
	// insertion point

	Name string
}

func (from *A_rows) GongCopyBasicFields(to *A_rows) {
	// insertion point
	to.Name = from.Name
}

type AbstractCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *AbstractCRSType) GongCopyBasicFields(to *AbstractCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type AbstractContinuousCoverageType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractContinuousCoverageType) GongCopyBasicFields(to *AbstractContinuousCoverageType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractCoordinateOperationType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string
}

func (from *AbstractCoordinateOperationType) GongCopyBasicFields(to *AbstractCoordinateOperationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
}

type AbstractCoordinateSystemType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *AbstractCoordinateSystemType) GongCopyBasicFields(to *AbstractCoordinateSystemType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type AbstractCoverageType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractCoverageType) GongCopyBasicFields(to *AbstractCoverageType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

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

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractCurveType) GongCopyBasicFields(to *AbstractCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type AbstractDQ_Completeness_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_Completeness_Type) GongCopyBasicFields(to *AbstractDQ_Completeness_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDQ_Element_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_Element_Type) GongCopyBasicFields(to *AbstractDQ_Element_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDQ_LogicalConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_LogicalConsistency_Type) GongCopyBasicFields(to *AbstractDQ_LogicalConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDQ_PositionalAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_PositionalAccuracy_Type) GongCopyBasicFields(to *AbstractDQ_PositionalAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDQ_Result_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_Result_Type) GongCopyBasicFields(to *AbstractDQ_Result_Type) {
	// insertion point
	*to = *from
}

type AbstractDQ_TemporalAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_TemporalAccuracy_Type) GongCopyBasicFields(to *AbstractDQ_TemporalAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDQ_ThematicAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractDQ_ThematicAccuracy_Type) GongCopyBasicFields(to *AbstractDQ_ThematicAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string
}

func (from *AbstractDatumType) GongCopyBasicFields(to *AbstractDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
}

type AbstractEX_GeographicExtent_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractEX_GeographicExtent_Type) GongCopyBasicFields(to *AbstractEX_GeographicExtent_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractFeatureCollectionType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractFeatureCollectionType) GongCopyBasicFields(to *AbstractFeatureCollectionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractFeatureMemberType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *AbstractFeatureMemberType) GongCopyBasicFields(to *AbstractFeatureMemberType) {
	// insertion point
	*to = *from
}

type AbstractFeatureType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractFeatureType) GongCopyBasicFields(to *AbstractFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractGMLType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractGMLType) GongCopyBasicFields(to *AbstractGMLType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractGeneralConversionType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *AbstractGeneralConversionType) GongCopyBasicFields(to *AbstractGeneralConversionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type AbstractGeneralDerivedCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *AbstractGeneralDerivedCRSType) GongCopyBasicFields(to *AbstractGeneralDerivedCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type AbstractGeneralOperationParameterPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *AbstractGeneralOperationParameterPropertyType) GongCopyBasicFields(to *AbstractGeneralOperationParameterPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type AbstractGeneralOperationParameterType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	MinimumOccurs int
}

func (from *AbstractGeneralOperationParameterType) GongCopyBasicFields(to *AbstractGeneralOperationParameterType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.MinimumOccurs = from.MinimumOccurs
}

type AbstractGeneralParameterValuePropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractGeneralParameterValuePropertyType) GongCopyBasicFields(to *AbstractGeneralParameterValuePropertyType) {
	// insertion point
	to.Name = from.Name
}

type AbstractGeneralParameterValueType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractGeneralParameterValueType) GongCopyBasicFields(to *AbstractGeneralParameterValueType) {
	// insertion point
	*to = *from
}

type AbstractGeneralTransformationType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string
}

func (from *AbstractGeneralTransformationType) GongCopyBasicFields(to *AbstractGeneralTransformationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
}

type AbstractGeometricAggregateType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *AbstractGeometricAggregateType) GongCopyBasicFields(to *AbstractGeometricAggregateType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type AbstractGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractGeometricPrimitiveType) GongCopyBasicFields(to *AbstractGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type AbstractGeometryType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractGeometryType) GongCopyBasicFields(to *AbstractGeometryType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type AbstractGriddedSurfaceType_WOP struct {
	// insertion point

	Name string

	AggregationType string

	Rows int

	Columns int
}

func (from *AbstractGriddedSurfaceType) GongCopyBasicFields(to *AbstractGriddedSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.AggregationType = from.AggregationType
	to.Rows = from.Rows
	to.Columns = from.Columns
}

type AbstractMD_Identification_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractMD_Identification_Type) GongCopyBasicFields(to *AbstractMD_Identification_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type AbstractMemberType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *AbstractMemberType) GongCopyBasicFields(to *AbstractMemberType) {
	// insertion point
	*to = *from
}

type AbstractMetaDataType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractMetaDataType) GongCopyBasicFields(to *AbstractMetaDataType) {
	// insertion point
	*to = *from
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

type AbstractParametricCurveSurfaceType_WOP struct {
	// insertion point

	Name string

	AggregationType string
}

func (from *AbstractParametricCurveSurfaceType) GongCopyBasicFields(to *AbstractParametricCurveSurfaceType) {
	// insertion point
	*to = *from
}

type AbstractRS_ReferenceSystem_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractRS_ReferenceSystem_Type) GongCopyBasicFields(to *AbstractRS_ReferenceSystem_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
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

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractRingType) GongCopyBasicFields(to *AbstractRingType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type AbstractSolidType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractSolidType) GongCopyBasicFields(to *AbstractSolidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
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

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *AbstractSurfaceType) GongCopyBasicFields(to *AbstractSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type AbstractTimeComplexType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTimeComplexType) GongCopyBasicFields(to *AbstractTimeComplexType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTimeGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Frame string
}

func (from *AbstractTimeGeometricPrimitiveType) GongCopyBasicFields(to *AbstractTimeGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Frame = from.Frame
}

type AbstractTimeObjectType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTimeObjectType) GongCopyBasicFields(to *AbstractTimeObjectType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTimePrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTimePrimitiveType) GongCopyBasicFields(to *AbstractTimePrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTimeSliceType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTimeSliceType) GongCopyBasicFields(to *AbstractTimeSliceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTimeTopologyPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTimeTopologyPrimitiveType) GongCopyBasicFields(to *AbstractTimeTopologyPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTopoPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTopoPrimitiveType) GongCopyBasicFields(to *AbstractTopoPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AbstractTopologyType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractTopologyType) GongCopyBasicFields(to *AbstractTopologyType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AffineCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *AffineCSPropertyType) GongCopyBasicFields(to *AffineCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type AffineCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *AffineCSType) GongCopyBasicFields(to *AffineCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type AffinePlacementType_WOP struct {
	// insertion point

	Name string

	InDimension int

	OutDimension int
}

func (from *AffinePlacementType) GongCopyBasicFields(to *AffinePlacementType) {
	// insertion point
	to.Name = from.Name
	to.InDimension = from.InDimension
	to.OutDimension = from.OutDimension
}

type AngleChoiceType_WOP struct {
	// insertion point

	Name string
}

func (from *AngleChoiceType) GongCopyBasicFields(to *AngleChoiceType) {
	// insertion point
	to.Name = from.Name
}

type Angle_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Angle_PropertyType) GongCopyBasicFields(to *Angle_PropertyType) {
	// insertion point
	*to = *from
}

type ArcByBulgeType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int

	Bulge float64
}

func (from *ArcByBulgeType) GongCopyBasicFields(to *ArcByBulgeType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
	to.Bulge = from.Bulge
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

type ArcStringByBulgeType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	NumArc int

	Bulge float64
}

func (from *ArcStringByBulgeType) GongCopyBasicFields(to *ArcStringByBulgeType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.NumArc = from.NumArc
	to.Bulge = from.Bulge
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

type ArrayAssociationType_WOP struct {
	// insertion point

	Name string

	Owns bool

	BooleanList string

	CountList string

	CountExtent string
}

func (from *ArrayAssociationType) GongCopyBasicFields(to *ArrayAssociationType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.BooleanList = from.BooleanList
	to.CountList = from.CountList
	to.CountExtent = from.CountExtent
}

type ArrayType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *ArrayType) GongCopyBasicFields(to *ArrayType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type AssociationRoleType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *AssociationRoleType) GongCopyBasicFields(to *AssociationRoleType) {
	// insertion point
	*to = *from
}

type BSplineType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	IsPolynomial bool

	KnotType Enum_KnotTypesType

	Degree int
}

func (from *BSplineType) GongCopyBasicFields(to *BSplineType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.IsPolynomial = from.IsPolynomial
	to.KnotType = from.KnotType
	to.Degree = from.Degree
}

type BagType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *BagType) GongCopyBasicFields(to *BagType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type BaseUnitType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *BaseUnitType) GongCopyBasicFields(to *BaseUnitType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type BezierType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	IsPolynomial bool

	KnotType Enum_KnotTypesType

	Degree int
}

func (from *BezierType) GongCopyBasicFields(to *BezierType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.IsPolynomial = from.IsPolynomial
	to.KnotType = from.KnotType
	to.Degree = from.Degree
}

type Binary_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Binary_PropertyType) GongCopyBasicFields(to *Binary_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type Binary_Type_WOP struct {
	// insertion point

	Name string

	Src string

	EnclosedText string
}

func (from *Binary_Type) GongCopyBasicFields(to *Binary_Type) {
	// insertion point
	*to = *from
}

type Boolean_WOP struct {
	// insertion point

	Name string

	NilReason string

	EnclosedText bool
}

func (from *Boolean) GongCopyBasicFields(to *Boolean) {
	// insertion point
	*to = *from
}

type BooleanPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *BooleanPropertyType) GongCopyBasicFields(to *BooleanPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Boolean_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Boolean bool
}

func (from *Boolean_PropertyType) GongCopyBasicFields(to *Boolean_PropertyType) {
	// insertion point
	*to = *from
}

type BoundedFeatureType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *BoundedFeatureType) GongCopyBasicFields(to *BoundedFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type BoundingShapeType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Null string
}

func (from *BoundingShapeType) GongCopyBasicFields(to *BoundingShapeType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Null = from.Null
}

type CI_Address_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Address_PropertyType) GongCopyBasicFields(to *CI_Address_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Address_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Address_Type) GongCopyBasicFields(to *CI_Address_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_Citation_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Citation_PropertyType) GongCopyBasicFields(to *CI_Citation_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Citation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Citation_Type) GongCopyBasicFields(to *CI_Citation_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_Contact_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Contact_PropertyType) GongCopyBasicFields(to *CI_Contact_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Contact_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Contact_Type) GongCopyBasicFields(to *CI_Contact_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_DateTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *CI_DateTypeCode_PropertyType) GongCopyBasicFields(to *CI_DateTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type CI_Date_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Date_PropertyType) GongCopyBasicFields(to *CI_Date_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Date_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Date_Type) GongCopyBasicFields(to *CI_Date_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_OnLineFunctionCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *CI_OnLineFunctionCode_PropertyType) GongCopyBasicFields(to *CI_OnLineFunctionCode_PropertyType) {
	// insertion point
	*to = *from
}

type CI_OnlineResource_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_OnlineResource_PropertyType) GongCopyBasicFields(to *CI_OnlineResource_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_OnlineResource_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_OnlineResource_Type) GongCopyBasicFields(to *CI_OnlineResource_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_PresentationFormCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *CI_PresentationFormCode_PropertyType) GongCopyBasicFields(to *CI_PresentationFormCode_PropertyType) {
	// insertion point
	*to = *from
}

type CI_ResponsibleParty_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_ResponsibleParty_PropertyType) GongCopyBasicFields(to *CI_ResponsibleParty_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_ResponsibleParty_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_ResponsibleParty_Type) GongCopyBasicFields(to *CI_ResponsibleParty_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_RoleCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *CI_RoleCode_PropertyType) GongCopyBasicFields(to *CI_RoleCode_PropertyType) {
	// insertion point
	*to = *from
}

type CI_Series_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Series_PropertyType) GongCopyBasicFields(to *CI_Series_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Series_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Series_Type) GongCopyBasicFields(to *CI_Series_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CI_Telephone_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CI_Telephone_PropertyType) GongCopyBasicFields(to *CI_Telephone_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CI_Telephone_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *CI_Telephone_Type) GongCopyBasicFields(to *CI_Telephone_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type CRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CRSPropertyType) GongCopyBasicFields(to *CRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CartesianCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CartesianCSPropertyType) GongCopyBasicFields(to *CartesianCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CartesianCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *CartesianCSType) GongCopyBasicFields(to *CartesianCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type Category_WOP struct {
	// insertion point

	Name string

	NilReason string

	EnclosedText string
}

func (from *Category) GongCopyBasicFields(to *Category) {
	// insertion point
	*to = *from
}

type CategoryExtentType_WOP struct {
	// insertion point

	Name string

	EnclosedText string
}

func (from *CategoryExtentType) GongCopyBasicFields(to *CategoryExtentType) {
	// insertion point
	*to = *from
}

type CategoryPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CategoryPropertyType) GongCopyBasicFields(to *CategoryPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CharacterString_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	CharacterString string

	MD_PixelOrientationCode string

	MD_ObligationCode string

	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type
}

func (from *CharacterString_PropertyType) GongCopyBasicFields(to *CharacterString_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.CharacterString = from.CharacterString
	to.MD_PixelOrientationCode = from.MD_PixelOrientationCode
	to.MD_ObligationCode = from.MD_ObligationCode
	to.MD_TopicCategoryCode = from.MD_TopicCategoryCode
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

type ClothoidType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	ScaleFactor string

	StartParameter float64

	EndParameter float64
}

func (from *ClothoidType) GongCopyBasicFields(to *ClothoidType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.ScaleFactor = from.ScaleFactor
	to.StartParameter = from.StartParameter
	to.EndParameter = from.EndParameter
}

type CompositeCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *CompositeCurveType) GongCopyBasicFields(to *CompositeCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type CompositeSolidType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *CompositeSolidType) GongCopyBasicFields(to *CompositeSolidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type CompositeSurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *CompositeSurfaceType) GongCopyBasicFields(to *CompositeSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type CompositeValueType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *CompositeValueType) GongCopyBasicFields(to *CompositeValueType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type CompoundCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CompoundCRSPropertyType) GongCopyBasicFields(to *CompoundCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CompoundCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	AggregationType string
}

func (from *CompoundCRSType) GongCopyBasicFields(to *CompoundCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.AggregationType = from.AggregationType
}

type ConcatenatedOperationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ConcatenatedOperationPropertyType) GongCopyBasicFields(to *ConcatenatedOperationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ConcatenatedOperationType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string

	AggregationType string
}

func (from *ConcatenatedOperationType) GongCopyBasicFields(to *ConcatenatedOperationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
	to.AggregationType = from.AggregationType
}

type ConeType_WOP struct {
	// insertion point

	Name string

	AggregationType string

	Rows int

	Columns int

	HorizontalCurveType Enum_CurveInterpolationType

	VerticalCurveType Enum_CurveInterpolationType
}

func (from *ConeType) GongCopyBasicFields(to *ConeType) {
	// insertion point
	to.Name = from.Name
	to.AggregationType = from.AggregationType
	to.Rows = from.Rows
	to.Columns = from.Columns
	to.HorizontalCurveType = from.HorizontalCurveType
	to.VerticalCurveType = from.VerticalCurveType
}

type ConventionalUnitType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *ConventionalUnitType) GongCopyBasicFields(to *ConventionalUnitType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type ConversionPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ConversionPropertyType) GongCopyBasicFields(to *ConversionPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ConversionToPreferredUnitType_WOP struct {
	// insertion point

	Name string

	Uom string

	Factor float64
}

func (from *ConversionToPreferredUnitType) GongCopyBasicFields(to *ConversionToPreferredUnitType) {
	// insertion point
	to.Name = from.Name
	to.Uom = from.Uom
	to.Factor = from.Factor
}

type ConversionType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string
}

func (from *ConversionType) GongCopyBasicFields(to *ConversionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
}

type CoordinateOperationAccuracy_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CoordinateOperationAccuracy) GongCopyBasicFields(to *CoordinateOperationAccuracy) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CoordinateOperationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CoordinateOperationPropertyType) GongCopyBasicFields(to *CoordinateOperationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CoordinateSystemAxisPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CoordinateSystemAxisPropertyType) GongCopyBasicFields(to *CoordinateSystemAxisPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CoordinateSystemAxisType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Uom string

	MinimumValue float64

	MaximumValue float64
}

func (from *CoordinateSystemAxisType) GongCopyBasicFields(to *CoordinateSystemAxisType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Uom = from.Uom
	to.MinimumValue = from.MinimumValue
	to.MaximumValue = from.MaximumValue
}

type CoordinateSystemPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CoordinateSystemPropertyType) GongCopyBasicFields(to *CoordinateSystemPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Count_WOP struct {
	// insertion point

	Name string

	NilReason string

	EnclosedText int
}

func (from *Count) GongCopyBasicFields(to *Count) {
	// insertion point
	*to = *from
}

type CountPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CountPropertyType) GongCopyBasicFields(to *CountPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Country_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Country_PropertyType) GongCopyBasicFields(to *Country_PropertyType) {
	// insertion point
	*to = *from
}

type CoverageFunctionType_WOP struct {
	// insertion point

	Name string
}

func (from *CoverageFunctionType) GongCopyBasicFields(to *CoverageFunctionType) {
	// insertion point
	to.Name = from.Name
}

type CubicSplineType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int

	Interpolation Enum_CurveInterpolationType

	Degree int
}

func (from *CubicSplineType) GongCopyBasicFields(to *CubicSplineType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
	to.Interpolation = from.Interpolation
	to.Degree = from.Degree
}

type CurveArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *CurveArrayPropertyType) GongCopyBasicFields(to *CurveArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type CurvePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *CurvePropertyType) GongCopyBasicFields(to *CurvePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
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

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *CurveType) GongCopyBasicFields(to *CurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type CylinderType_WOP struct {
	// insertion point

	Name string

	AggregationType string

	Rows int

	Columns int

	HorizontalCurveType Enum_CurveInterpolationType

	VerticalCurveType Enum_CurveInterpolationType
}

func (from *CylinderType) GongCopyBasicFields(to *CylinderType) {
	// insertion point
	to.Name = from.Name
	to.AggregationType = from.AggregationType
	to.Rows = from.Rows
	to.Columns = from.Columns
	to.HorizontalCurveType = from.HorizontalCurveType
	to.VerticalCurveType = from.VerticalCurveType
}

type CylindricalCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *CylindricalCSPropertyType) GongCopyBasicFields(to *CylindricalCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type CylindricalCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *CylindricalCSType) GongCopyBasicFields(to *CylindricalCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type DMSAngleType_WOP struct {
	// insertion point

	Name string

	DecimalMinutes string

	Minutes int

	Seconds string
}

func (from *DMSAngleType) GongCopyBasicFields(to *DMSAngleType) {
	// insertion point
	to.Name = from.Name
	to.DecimalMinutes = from.DecimalMinutes
	to.Minutes = from.Minutes
	to.Seconds = from.Seconds
}

type DQ_AbsoluteExternalPositionalAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongCopyBasicFields(to *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_AbsoluteExternalPositionalAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_AbsoluteExternalPositionalAccuracy_Type) GongCopyBasicFields(to *DQ_AbsoluteExternalPositionalAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_AccuracyOfATimeMeasurement_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_AccuracyOfATimeMeasurement_PropertyType) GongCopyBasicFields(to *DQ_AccuracyOfATimeMeasurement_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_AccuracyOfATimeMeasurement_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_AccuracyOfATimeMeasurement_Type) GongCopyBasicFields(to *DQ_AccuracyOfATimeMeasurement_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_CompletenessCommission_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_CompletenessCommission_PropertyType) GongCopyBasicFields(to *DQ_CompletenessCommission_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_CompletenessCommission_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_CompletenessCommission_Type) GongCopyBasicFields(to *DQ_CompletenessCommission_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_CompletenessOmission_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_CompletenessOmission_PropertyType) GongCopyBasicFields(to *DQ_CompletenessOmission_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_CompletenessOmission_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_CompletenessOmission_Type) GongCopyBasicFields(to *DQ_CompletenessOmission_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_Completeness_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_Completeness_PropertyType) GongCopyBasicFields(to *DQ_Completeness_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_ConceptualConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_ConceptualConsistency_PropertyType) GongCopyBasicFields(to *DQ_ConceptualConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_ConceptualConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_ConceptualConsistency_Type) GongCopyBasicFields(to *DQ_ConceptualConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_ConformanceResult_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_ConformanceResult_PropertyType) GongCopyBasicFields(to *DQ_ConformanceResult_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_ConformanceResult_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_ConformanceResult_Type) GongCopyBasicFields(to *DQ_ConformanceResult_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_DataQuality_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_DataQuality_PropertyType) GongCopyBasicFields(to *DQ_DataQuality_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_DataQuality_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_DataQuality_Type) GongCopyBasicFields(to *DQ_DataQuality_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_DomainConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_DomainConsistency_PropertyType) GongCopyBasicFields(to *DQ_DomainConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_DomainConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_DomainConsistency_Type) GongCopyBasicFields(to *DQ_DomainConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_Element_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_Element_PropertyType) GongCopyBasicFields(to *DQ_Element_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_EvaluationMethodTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *DQ_EvaluationMethodTypeCode_PropertyType) GongCopyBasicFields(to *DQ_EvaluationMethodTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type DQ_FormatConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_FormatConsistency_PropertyType) GongCopyBasicFields(to *DQ_FormatConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_FormatConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_FormatConsistency_Type) GongCopyBasicFields(to *DQ_FormatConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_GriddedDataPositionalAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_GriddedDataPositionalAccuracy_PropertyType) GongCopyBasicFields(to *DQ_GriddedDataPositionalAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_GriddedDataPositionalAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_GriddedDataPositionalAccuracy_Type) GongCopyBasicFields(to *DQ_GriddedDataPositionalAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_LogicalConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_LogicalConsistency_PropertyType) GongCopyBasicFields(to *DQ_LogicalConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_NonQuantitativeAttributeAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongCopyBasicFields(to *DQ_NonQuantitativeAttributeAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_NonQuantitativeAttributeAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_NonQuantitativeAttributeAccuracy_Type) GongCopyBasicFields(to *DQ_NonQuantitativeAttributeAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_PositionalAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_PositionalAccuracy_PropertyType) GongCopyBasicFields(to *DQ_PositionalAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_QuantitativeAttributeAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_QuantitativeAttributeAccuracy_PropertyType) GongCopyBasicFields(to *DQ_QuantitativeAttributeAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_QuantitativeAttributeAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_QuantitativeAttributeAccuracy_Type) GongCopyBasicFields(to *DQ_QuantitativeAttributeAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_QuantitativeResult_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_QuantitativeResult_PropertyType) GongCopyBasicFields(to *DQ_QuantitativeResult_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_QuantitativeResult_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_QuantitativeResult_Type) GongCopyBasicFields(to *DQ_QuantitativeResult_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_RelativeInternalPositionalAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongCopyBasicFields(to *DQ_RelativeInternalPositionalAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_RelativeInternalPositionalAccuracy_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_RelativeInternalPositionalAccuracy_Type) GongCopyBasicFields(to *DQ_RelativeInternalPositionalAccuracy_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_Result_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_Result_PropertyType) GongCopyBasicFields(to *DQ_Result_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_Scope_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_Scope_PropertyType) GongCopyBasicFields(to *DQ_Scope_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_Scope_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_Scope_Type) GongCopyBasicFields(to *DQ_Scope_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_TemporalAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_TemporalAccuracy_PropertyType) GongCopyBasicFields(to *DQ_TemporalAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_TemporalConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_TemporalConsistency_PropertyType) GongCopyBasicFields(to *DQ_TemporalConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_TemporalConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_TemporalConsistency_Type) GongCopyBasicFields(to *DQ_TemporalConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_TemporalValidity_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_TemporalValidity_PropertyType) GongCopyBasicFields(to *DQ_TemporalValidity_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_TemporalValidity_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_TemporalValidity_Type) GongCopyBasicFields(to *DQ_TemporalValidity_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_ThematicAccuracy_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_ThematicAccuracy_PropertyType) GongCopyBasicFields(to *DQ_ThematicAccuracy_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_ThematicClassificationCorrectness_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_ThematicClassificationCorrectness_PropertyType) GongCopyBasicFields(to *DQ_ThematicClassificationCorrectness_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_ThematicClassificationCorrectness_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_ThematicClassificationCorrectness_Type) GongCopyBasicFields(to *DQ_ThematicClassificationCorrectness_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DQ_TopologicalConsistency_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DQ_TopologicalConsistency_PropertyType) GongCopyBasicFields(to *DQ_TopologicalConsistency_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DQ_TopologicalConsistency_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DQ_TopologicalConsistency_Type) GongCopyBasicFields(to *DQ_TopologicalConsistency_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type DS_AssociationTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *DS_AssociationTypeCode_PropertyType) GongCopyBasicFields(to *DS_AssociationTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type DS_Association_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DS_Association_PropertyType) GongCopyBasicFields(to *DS_Association_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DS_Association_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *DS_Association_Type) GongCopyBasicFields(to *DS_Association_Type) {
	// insertion point
	*to = *from
}

type DS_InitiativeTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *DS_InitiativeTypeCode_PropertyType) GongCopyBasicFields(to *DS_InitiativeTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type DataBlockType_WOP struct {
	// insertion point

	Name string

	DoubleOrNilReasonTupleList string
}

func (from *DataBlockType) GongCopyBasicFields(to *DataBlockType) {
	// insertion point
	to.Name = from.Name
	to.DoubleOrNilReasonTupleList = from.DoubleOrNilReasonTupleList
}

type DateTime_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	DateTime string
}

func (from *DateTime_PropertyType) GongCopyBasicFields(to *DateTime_PropertyType) {
	// insertion point
	*to = *from
}

type Date_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Date string

	DateTime string
}

func (from *Date_PropertyType) GongCopyBasicFields(to *Date_PropertyType) {
	// insertion point
	*to = *from
}

type DatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DatumPropertyType) GongCopyBasicFields(to *DatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Decimal_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Decimal string
}

func (from *Decimal_PropertyType) GongCopyBasicFields(to *Decimal_PropertyType) {
	// insertion point
	*to = *from
}

type DefinitionBaseType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DefinitionBaseType) GongCopyBasicFields(to *DefinitionBaseType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type DefinitionProxyType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *DefinitionProxyType) GongCopyBasicFields(to *DefinitionProxyType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type DefinitionType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *DefinitionType) GongCopyBasicFields(to *DefinitionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type DegreesType_WOP struct {
	// insertion point

	Name string

	Direction Enum_DegreesType_direction

	EnclosedText int
}

func (from *DegreesType) GongCopyBasicFields(to *DegreesType) {
	// insertion point
	*to = *from
}

type DerivationUnitTermType_WOP struct {
	// insertion point

	Name string

	Uom string

	Exponent int
}

func (from *DerivationUnitTermType) GongCopyBasicFields(to *DerivationUnitTermType) {
	// insertion point
	*to = *from
}

type DerivedCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DerivedCRSPropertyType) GongCopyBasicFields(to *DerivedCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DerivedCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *DerivedCRSType) GongCopyBasicFields(to *DerivedCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type DerivedUnitType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *DerivedUnitType) GongCopyBasicFields(to *DerivedUnitType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type DictionaryEntryType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DictionaryEntryType) GongCopyBasicFields(to *DictionaryEntryType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DictionaryType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *DictionaryType) GongCopyBasicFields(to *DictionaryType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type DirectPositionListType_WOP struct {
	// insertion point

	Name string

	Count int

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

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

	SrsDimension int

	AxisLabels string

	UomLabels string

	EnclosedText string
}

func (from *DirectPositionType) GongCopyBasicFields(to *DirectPositionType) {
	// insertion point
	*to = *from
}

type DirectedEdgePropertyType_WOP struct {
	// insertion point

	Name string

	Orientation string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *DirectedEdgePropertyType) GongCopyBasicFields(to *DirectedEdgePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Orientation = from.Orientation
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type DirectedFacePropertyType_WOP struct {
	// insertion point

	Name string

	Orientation string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *DirectedFacePropertyType) GongCopyBasicFields(to *DirectedFacePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Orientation = from.Orientation
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type DirectedNodePropertyType_WOP struct {
	// insertion point

	Name string

	Orientation string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *DirectedNodePropertyType) GongCopyBasicFields(to *DirectedNodePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Orientation = from.Orientation
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type DirectedObservationAtDistanceType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DirectedObservationAtDistanceType) GongCopyBasicFields(to *DirectedObservationAtDistanceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type DirectedObservationType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DirectedObservationType) GongCopyBasicFields(to *DirectedObservationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type DirectedTopoSolidPropertyType_WOP struct {
	// insertion point

	Name string

	Orientation string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *DirectedTopoSolidPropertyType) GongCopyBasicFields(to *DirectedTopoSolidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Orientation = from.Orientation
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type DirectionDescriptionType_WOP struct {
	// insertion point

	Name string

	CompassPoint Enum_CompassPointEnumeration

	Description string
}

func (from *DirectionDescriptionType) GongCopyBasicFields(to *DirectionDescriptionType) {
	// insertion point
	to.Name = from.Name
	to.CompassPoint = from.CompassPoint
	to.Description = from.Description
}

type DirectionPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	CompassPoint Enum_CompassPointEnumeration
}

func (from *DirectionPropertyType) GongCopyBasicFields(to *DirectionPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.CompassPoint = from.CompassPoint
}

type DirectionVectorType_WOP struct {
	// insertion point

	Name string
}

func (from *DirectionVectorType) GongCopyBasicFields(to *DirectionVectorType) {
	// insertion point
	to.Name = from.Name
}

type DiscreteCoverageType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DiscreteCoverageType) GongCopyBasicFields(to *DiscreteCoverageType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type Distance_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Distance_PropertyType) GongCopyBasicFields(to *Distance_PropertyType) {
	// insertion point
	*to = *from
}

type DomainOfValidity_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DomainOfValidity) GongCopyBasicFields(to *DomainOfValidity) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DomainSetType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DomainSetType) GongCopyBasicFields(to *DomainSetType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DynamicFeatureCollectionType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DynamicFeatureCollectionType) GongCopyBasicFields(to *DynamicFeatureCollectionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type DynamicFeatureMemberType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *DynamicFeatureMemberType) GongCopyBasicFields(to *DynamicFeatureMemberType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type DynamicFeatureType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *DynamicFeatureType) GongCopyBasicFields(to *DynamicFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type EX_BoundingPolygon_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_BoundingPolygon_PropertyType) GongCopyBasicFields(to *EX_BoundingPolygon_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_BoundingPolygon_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_BoundingPolygon_Type) GongCopyBasicFields(to *EX_BoundingPolygon_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_Extent_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_Extent_PropertyType) GongCopyBasicFields(to *EX_Extent_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_Extent_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_Extent_Type) GongCopyBasicFields(to *EX_Extent_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_GeographicBoundingBox_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_GeographicBoundingBox_PropertyType) GongCopyBasicFields(to *EX_GeographicBoundingBox_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_GeographicBoundingBox_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_GeographicBoundingBox_Type) GongCopyBasicFields(to *EX_GeographicBoundingBox_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_GeographicDescription_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_GeographicDescription_PropertyType) GongCopyBasicFields(to *EX_GeographicDescription_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_GeographicDescription_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_GeographicDescription_Type) GongCopyBasicFields(to *EX_GeographicDescription_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_GeographicExtent_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_GeographicExtent_PropertyType) GongCopyBasicFields(to *EX_GeographicExtent_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_SpatialTemporalExtent_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_SpatialTemporalExtent_PropertyType) GongCopyBasicFields(to *EX_SpatialTemporalExtent_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_SpatialTemporalExtent_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_SpatialTemporalExtent_Type) GongCopyBasicFields(to *EX_SpatialTemporalExtent_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_TemporalExtent_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_TemporalExtent_PropertyType) GongCopyBasicFields(to *EX_TemporalExtent_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_TemporalExtent_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_TemporalExtent_Type) GongCopyBasicFields(to *EX_TemporalExtent_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EX_VerticalExtent_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EX_VerticalExtent_PropertyType) GongCopyBasicFields(to *EX_VerticalExtent_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EX_VerticalExtent_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *EX_VerticalExtent_Type) GongCopyBasicFields(to *EX_VerticalExtent_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type EdgeType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *EdgeType) GongCopyBasicFields(to *EdgeType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type EllipsoidPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EllipsoidPropertyType) GongCopyBasicFields(to *EllipsoidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EllipsoidType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *EllipsoidType) GongCopyBasicFields(to *EllipsoidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type EllipsoidalCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EllipsoidalCSPropertyType) GongCopyBasicFields(to *EllipsoidalCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EllipsoidalCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *EllipsoidalCSType) GongCopyBasicFields(to *EllipsoidalCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type EngineeringCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EngineeringCRSPropertyType) GongCopyBasicFields(to *EngineeringCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EngineeringCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *EngineeringCRSType) GongCopyBasicFields(to *EngineeringCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type EngineeringDatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *EngineeringDatumPropertyType) GongCopyBasicFields(to *EngineeringDatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type EngineeringDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string
}

func (from *EngineeringDatumType) GongCopyBasicFields(to *EngineeringDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
}

type EnvelopeType_WOP struct {
	// insertion point

	Name string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *EnvelopeType) GongCopyBasicFields(to *EnvelopeType) {
	// insertion point
	to.Name = from.Name
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type EnvelopeWithTimePeriodType_WOP struct {
	// insertion point

	Name string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	Frame string
}

func (from *EnvelopeWithTimePeriodType) GongCopyBasicFields(to *EnvelopeWithTimePeriodType) {
	// insertion point
	to.Name = from.Name
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.Frame = from.Frame
}

type FaceOrTopoSolidPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *FaceOrTopoSolidPropertyType) GongCopyBasicFields(to *FaceOrTopoSolidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type FaceType_WOP struct {
	// insertion point

	Name string

	Id string

	Universal bool

	AggregationType string
}

func (from *FaceType) GongCopyBasicFields(to *FaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Universal = from.Universal
	to.AggregationType = from.AggregationType
}

type FeatureArrayPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *FeatureArrayPropertyType) GongCopyBasicFields(to *FeatureArrayPropertyType) {
	// insertion point
	to.Name = from.Name
}

type FeatureCollectionType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *FeatureCollectionType) GongCopyBasicFields(to *FeatureCollectionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type FeaturePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *FeaturePropertyType) GongCopyBasicFields(to *FeaturePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type FileType_WOP struct {
	// insertion point

	Name string

	FileName string

	FileReference string

	MimeType string

	Compression string
}

func (from *FileType) GongCopyBasicFields(to *FileType) {
	// insertion point
	to.Name = from.Name
	to.FileName = from.FileName
	to.FileReference = from.FileReference
	to.MimeType = from.MimeType
	to.Compression = from.Compression
}

type FormulaCitation_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *FormulaCitation) GongCopyBasicFields(to *FormulaCitation) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type FormulaType_WOP struct {
	// insertion point

	Name string

	A float64

	B float64

	C float64

	D float64
}

func (from *FormulaType) GongCopyBasicFields(to *FormulaType) {
	// insertion point
	*to = *from
}

type GM_Object_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GM_Object_PropertyType) GongCopyBasicFields(to *GM_Object_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GM_Point_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GM_Point_PropertyType) GongCopyBasicFields(to *GM_Point_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeneralConversionPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeneralConversionPropertyType) GongCopyBasicFields(to *GeneralConversionPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeneralTransformationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeneralTransformationPropertyType) GongCopyBasicFields(to *GeneralTransformationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GenericMetaDataType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *GenericMetaDataType) GongCopyBasicFields(to *GenericMetaDataType) {
	// insertion point
	*to = *from
}

type GenericName_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *GenericName_PropertyType) GongCopyBasicFields(to *GenericName_PropertyType) {
	// insertion point
	*to = *from
}

type GeocentricCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeocentricCRSPropertyType) GongCopyBasicFields(to *GeocentricCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeocentricCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *GeocentricCRSType) GongCopyBasicFields(to *GeocentricCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
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

type GeodeticCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeodeticCRSPropertyType) GongCopyBasicFields(to *GeodeticCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeodeticCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *GeodeticCRSType) GongCopyBasicFields(to *GeodeticCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type GeodeticDatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeodeticDatumPropertyType) GongCopyBasicFields(to *GeodeticDatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeodeticDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string
}

func (from *GeodeticDatumType) GongCopyBasicFields(to *GeodeticDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
}

type GeographicCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeographicCRSPropertyType) GongCopyBasicFields(to *GeographicCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeographicCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *GeographicCRSType) GongCopyBasicFields(to *GeographicCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type GeometricComplexPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeometricComplexPropertyType) GongCopyBasicFields(to *GeometricComplexPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeometricComplexType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *GeometricComplexType) GongCopyBasicFields(to *GeometricComplexType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type GeometricPrimitivePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *GeometricPrimitivePropertyType) GongCopyBasicFields(to *GeometricPrimitivePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type GeometryArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *GeometryArrayPropertyType) GongCopyBasicFields(to *GeometryArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type GeometryPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *GeometryPropertyType) GongCopyBasicFields(to *GeometryPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type GridEnvelopeType_WOP struct {
	// insertion point

	Name string

	Low string

	High string
}

func (from *GridEnvelopeType) GongCopyBasicFields(to *GridEnvelopeType) {
	// insertion point
	*to = *from
}

type GridFunctionType_WOP struct {
	// insertion point

	Name string

	StartPoint string
}

func (from *GridFunctionType) GongCopyBasicFields(to *GridFunctionType) {
	// insertion point
	to.Name = from.Name
	to.StartPoint = from.StartPoint
}

type GridLimitsType_WOP struct {
	// insertion point

	Name string
}

func (from *GridLimitsType) GongCopyBasicFields(to *GridLimitsType) {
	// insertion point
	to.Name = from.Name
}

type GridType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	Dimension int

	AxisLabels_1 string

	AxisName string
}

func (from *GridType) GongCopyBasicFields(to *GridType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.Dimension = from.Dimension
	to.AxisLabels_1 = from.AxisLabels_1
	to.AxisName = from.AxisName
}

type HistoryPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *HistoryPropertyType) GongCopyBasicFields(to *HistoryPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type IdentifiedObjectType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *IdentifiedObjectType) GongCopyBasicFields(to *IdentifiedObjectType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type ImageCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ImageCRSPropertyType) GongCopyBasicFields(to *ImageCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ImageCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *ImageCRSType) GongCopyBasicFields(to *ImageCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type ImageDatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ImageDatumPropertyType) GongCopyBasicFields(to *ImageDatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ImageDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string
}

func (from *ImageDatumType) GongCopyBasicFields(to *ImageDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
}

type IndirectEntryType_WOP struct {
	// insertion point

	Name string
}

func (from *IndirectEntryType) GongCopyBasicFields(to *IndirectEntryType) {
	// insertion point
	to.Name = from.Name
}

type InlinePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *InlinePropertyType) GongCopyBasicFields(to *InlinePropertyType) {
	// insertion point
	*to = *from
}

type Integer_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Integer int
}

func (from *Integer_PropertyType) GongCopyBasicFields(to *Integer_PropertyType) {
	// insertion point
	*to = *from
}

type KnotPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *KnotPropertyType) GongCopyBasicFields(to *KnotPropertyType) {
	// insertion point
	to.Name = from.Name
}

type KnotType_WOP struct {
	// insertion point

	Name string

	Value float64

	Multiplicity int

	Weight float64
}

func (from *KnotType) GongCopyBasicFields(to *KnotType) {
	// insertion point
	*to = *from
}

type LI_Lineage_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *LI_Lineage_PropertyType) GongCopyBasicFields(to *LI_Lineage_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type LI_Lineage_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *LI_Lineage_Type) GongCopyBasicFields(to *LI_Lineage_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type LI_ProcessStep_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *LI_ProcessStep_PropertyType) GongCopyBasicFields(to *LI_ProcessStep_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type LI_ProcessStep_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *LI_ProcessStep_Type) GongCopyBasicFields(to *LI_ProcessStep_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type LI_Source_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *LI_Source_PropertyType) GongCopyBasicFields(to *LI_Source_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type LI_Source_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *LI_Source_Type) GongCopyBasicFields(to *LI_Source_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type LanguageCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *LanguageCode_PropertyType) GongCopyBasicFields(to *LanguageCode_PropertyType) {
	// insertion point
	*to = *from
}

type Length_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Length_PropertyType) GongCopyBasicFields(to *Length_PropertyType) {
	// insertion point
	*to = *from
}

type LineStringSegmentArrayPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *LineStringSegmentArrayPropertyType) GongCopyBasicFields(to *LineStringSegmentArrayPropertyType) {
	// insertion point
	to.Name = from.Name
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

type LineStringType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *LineStringType) GongCopyBasicFields(to *LineStringType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type LinearCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *LinearCSPropertyType) GongCopyBasicFields(to *LinearCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type LinearCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *LinearCSType) GongCopyBasicFields(to *LinearCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type LinearRingPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *LinearRingPropertyType) GongCopyBasicFields(to *LinearRingPropertyType) {
	// insertion point
	to.Name = from.Name
}

type LinearRingType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *LinearRingType) GongCopyBasicFields(to *LinearRingType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type LocalName_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *LocalName_PropertyType) GongCopyBasicFields(to *LocalName_PropertyType) {
	// insertion point
	*to = *from
}

type LocalisedCharacterString_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *LocalisedCharacterString_PropertyType) GongCopyBasicFields(to *LocalisedCharacterString_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type LocalisedCharacterString_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Locale string

	EnclosedText string
}

func (from *LocalisedCharacterString_Type) GongCopyBasicFields(to *LocalisedCharacterString_Type) {
	// insertion point
	*to = *from
}

type LocationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Null string
}

func (from *LocationPropertyType) GongCopyBasicFields(to *LocationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Null = from.Null
}

type MD_AggregateInformation_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_AggregateInformation_PropertyType) GongCopyBasicFields(to *MD_AggregateInformation_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_AggregateInformation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_AggregateInformation_Type) GongCopyBasicFields(to *MD_AggregateInformation_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_BrowseGraphic_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_BrowseGraphic_PropertyType) GongCopyBasicFields(to *MD_BrowseGraphic_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_BrowseGraphic_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_BrowseGraphic_Type) GongCopyBasicFields(to *MD_BrowseGraphic_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_CharacterSetCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_CharacterSetCode_PropertyType) GongCopyBasicFields(to *MD_CharacterSetCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_ClassificationCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_ClassificationCode_PropertyType) GongCopyBasicFields(to *MD_ClassificationCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_Constraints_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Constraints_PropertyType) GongCopyBasicFields(to *MD_Constraints_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Constraints_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Constraints_Type) GongCopyBasicFields(to *MD_Constraints_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_DataIdentification_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_DataIdentification_PropertyType) GongCopyBasicFields(to *MD_DataIdentification_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_DataIdentification_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_DataIdentification_Type) GongCopyBasicFields(to *MD_DataIdentification_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_DigitalTransferOptions_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_DigitalTransferOptions_PropertyType) GongCopyBasicFields(to *MD_DigitalTransferOptions_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_DigitalTransferOptions_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_DigitalTransferOptions_Type) GongCopyBasicFields(to *MD_DigitalTransferOptions_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_DistributionUnits_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_DistributionUnits_PropertyType) GongCopyBasicFields(to *MD_DistributionUnits_PropertyType) {
	// insertion point
	*to = *from
}

type MD_Distribution_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Distribution_PropertyType) GongCopyBasicFields(to *MD_Distribution_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Distribution_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Distribution_Type) GongCopyBasicFields(to *MD_Distribution_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_Distributor_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Distributor_PropertyType) GongCopyBasicFields(to *MD_Distributor_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Distributor_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Distributor_Type) GongCopyBasicFields(to *MD_Distributor_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_Format_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Format_PropertyType) GongCopyBasicFields(to *MD_Format_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Format_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Format_Type) GongCopyBasicFields(to *MD_Format_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_Identification_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Identification_PropertyType) GongCopyBasicFields(to *MD_Identification_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Identifier_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Identifier_PropertyType) GongCopyBasicFields(to *MD_Identifier_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Identifier_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Identifier_Type) GongCopyBasicFields(to *MD_Identifier_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_KeywordTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_KeywordTypeCode_PropertyType) GongCopyBasicFields(to *MD_KeywordTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_Keywords_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Keywords_PropertyType) GongCopyBasicFields(to *MD_Keywords_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Keywords_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Keywords_Type) GongCopyBasicFields(to *MD_Keywords_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_LegalConstraints_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_LegalConstraints_PropertyType) GongCopyBasicFields(to *MD_LegalConstraints_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_LegalConstraints_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_LegalConstraints_Type) GongCopyBasicFields(to *MD_LegalConstraints_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_MaintenanceFrequencyCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_MaintenanceFrequencyCode_PropertyType) GongCopyBasicFields(to *MD_MaintenanceFrequencyCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_MaintenanceInformation_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_MaintenanceInformation_PropertyType) GongCopyBasicFields(to *MD_MaintenanceInformation_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_MaintenanceInformation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_MaintenanceInformation_Type) GongCopyBasicFields(to *MD_MaintenanceInformation_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_MediumFormatCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_MediumFormatCode_PropertyType) GongCopyBasicFields(to *MD_MediumFormatCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_MediumNameCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_MediumNameCode_PropertyType) GongCopyBasicFields(to *MD_MediumNameCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_Medium_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Medium_PropertyType) GongCopyBasicFields(to *MD_Medium_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Medium_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Medium_Type) GongCopyBasicFields(to *MD_Medium_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_ProgressCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_ProgressCode_PropertyType) GongCopyBasicFields(to *MD_ProgressCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_ReferenceSystem_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_ReferenceSystem_PropertyType) GongCopyBasicFields(to *MD_ReferenceSystem_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_ReferenceSystem_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_ReferenceSystem_Type) GongCopyBasicFields(to *MD_ReferenceSystem_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_RepresentativeFraction_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_RepresentativeFraction_PropertyType) GongCopyBasicFields(to *MD_RepresentativeFraction_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_RepresentativeFraction_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_RepresentativeFraction_Type) GongCopyBasicFields(to *MD_RepresentativeFraction_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_Resolution_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_Resolution_PropertyType) GongCopyBasicFields(to *MD_Resolution_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type MD_Resolution_Type_WOP struct {
	// insertion point

	Name string
}

func (from *MD_Resolution_Type) GongCopyBasicFields(to *MD_Resolution_Type) {
	// insertion point
	to.Name = from.Name
}

type MD_RestrictionCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_RestrictionCode_PropertyType) GongCopyBasicFields(to *MD_RestrictionCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_ScopeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_ScopeCode_PropertyType) GongCopyBasicFields(to *MD_ScopeCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_ScopeDescription_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_ScopeDescription_PropertyType) GongCopyBasicFields(to *MD_ScopeDescription_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type MD_ScopeDescription_Type_WOP struct {
	// insertion point

	Name string
}

func (from *MD_ScopeDescription_Type) GongCopyBasicFields(to *MD_ScopeDescription_Type) {
	// insertion point
	to.Name = from.Name
}

type MD_SecurityConstraints_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_SecurityConstraints_PropertyType) GongCopyBasicFields(to *MD_SecurityConstraints_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_SecurityConstraints_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_SecurityConstraints_Type) GongCopyBasicFields(to *MD_SecurityConstraints_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_ServiceIdentification_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_ServiceIdentification_PropertyType) GongCopyBasicFields(to *MD_ServiceIdentification_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_ServiceIdentification_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_ServiceIdentification_Type) GongCopyBasicFields(to *MD_ServiceIdentification_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_SpatialRepresentationTypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_SpatialRepresentationTypeCode_PropertyType) GongCopyBasicFields(to *MD_SpatialRepresentationTypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_StandardOrderProcess_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_StandardOrderProcess_PropertyType) GongCopyBasicFields(to *MD_StandardOrderProcess_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_StandardOrderProcess_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_StandardOrderProcess_Type) GongCopyBasicFields(to *MD_StandardOrderProcess_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_TopicCategoryCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type
}

func (from *MD_TopicCategoryCode_PropertyType) GongCopyBasicFields(to *MD_TopicCategoryCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_Usage_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MD_Usage_PropertyType) GongCopyBasicFields(to *MD_Usage_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_Usage_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_Usage_Type) GongCopyBasicFields(to *MD_Usage_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MappingRuleType_WOP struct {
	// insertion point

	Name string

	RuleDefinition string
}

func (from *MappingRuleType) GongCopyBasicFields(to *MappingRuleType) {
	// insertion point
	to.Name = from.Name
	to.RuleDefinition = from.RuleDefinition
}

type Measure_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Measure_PropertyType) GongCopyBasicFields(to *Measure_PropertyType) {
	// insertion point
	*to = *from
}

type MemberName_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MemberName_PropertyType) GongCopyBasicFields(to *MemberName_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MemberName_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MemberName_Type) GongCopyBasicFields(to *MemberName_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MetaDataPropertyType_WOP struct {
	// insertion point

	Name string

	About string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MetaDataPropertyType) GongCopyBasicFields(to *MetaDataPropertyType) {
	// insertion point
	to.Name = from.Name
	to.About = from.About
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MovingObjectStatusType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *MovingObjectStatusType) GongCopyBasicFields(to *MovingObjectStatusType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type MultiCurvePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *MultiCurvePropertyType) GongCopyBasicFields(to *MultiCurvePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type MultiCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *MultiCurveType) GongCopyBasicFields(to *MultiCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type MultiGeometryPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *MultiGeometryPropertyType) GongCopyBasicFields(to *MultiGeometryPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type MultiGeometryType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *MultiGeometryType) GongCopyBasicFields(to *MultiGeometryType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type MultiPointPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *MultiPointPropertyType) GongCopyBasicFields(to *MultiPointPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type MultiPointType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *MultiPointType) GongCopyBasicFields(to *MultiPointType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type MultiSolidPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *MultiSolidPropertyType) GongCopyBasicFields(to *MultiSolidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type MultiSolidType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *MultiSolidType) GongCopyBasicFields(to *MultiSolidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type MultiSurfacePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *MultiSurfacePropertyType) GongCopyBasicFields(to *MultiSurfacePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type MultiSurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *MultiSurfaceType) GongCopyBasicFields(to *MultiSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type MultiplicityRange_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *MultiplicityRange_PropertyType) GongCopyBasicFields(to *MultiplicityRange_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MultiplicityRange_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MultiplicityRange_Type) GongCopyBasicFields(to *MultiplicityRange_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type Multiplicity_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *Multiplicity_PropertyType) GongCopyBasicFields(to *Multiplicity_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Multiplicity_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *Multiplicity_Type) GongCopyBasicFields(to *Multiplicity_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type NodeOrEdgePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *NodeOrEdgePropertyType) GongCopyBasicFields(to *NodeOrEdgePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type NodePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *NodePropertyType) GongCopyBasicFields(to *NodePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type NodeType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *NodeType) GongCopyBasicFields(to *NodeType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type Number_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Real_XSD float64

	Decimal string

	Integer int
}

func (from *Number_PropertyType) GongCopyBasicFields(to *Number_PropertyType) {
	// insertion point
	*to = *from
}

type ObliqueCartesianCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ObliqueCartesianCSPropertyType) GongCopyBasicFields(to *ObliqueCartesianCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ObliqueCartesianCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *ObliqueCartesianCSType) GongCopyBasicFields(to *ObliqueCartesianCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type ObservationType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *ObservationType) GongCopyBasicFields(to *ObservationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type OffsetCurveType_WOP struct {
	// insertion point

	Name string

	NumDerivativesAtStart int

	NumDerivativesAtEnd int

	NumDerivativeInterior int
}

func (from *OffsetCurveType) GongCopyBasicFields(to *OffsetCurveType) {
	// insertion point
	to.Name = from.Name
	to.NumDerivativesAtStart = from.NumDerivativesAtStart
	to.NumDerivativesAtEnd = from.NumDerivativesAtEnd
	to.NumDerivativeInterior = from.NumDerivativeInterior
}

type OperationMethodPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *OperationMethodPropertyType) GongCopyBasicFields(to *OperationMethodPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type OperationMethodType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	SourceDimensions int

	TargetDimensions int
}

func (from *OperationMethodType) GongCopyBasicFields(to *OperationMethodType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.SourceDimensions = from.SourceDimensions
	to.TargetDimensions = from.TargetDimensions
}

type OperationParameterGroupPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *OperationParameterGroupPropertyType) GongCopyBasicFields(to *OperationParameterGroupPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type OperationParameterGroupType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	MinimumOccurs int

	MaximumOccurs int
}

func (from *OperationParameterGroupType) GongCopyBasicFields(to *OperationParameterGroupType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.MinimumOccurs = from.MinimumOccurs
	to.MaximumOccurs = from.MaximumOccurs
}

type OperationParameterPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *OperationParameterPropertyType) GongCopyBasicFields(to *OperationParameterPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type OperationParameterType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	MinimumOccurs int
}

func (from *OperationParameterType) GongCopyBasicFields(to *OperationParameterType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.MinimumOccurs = from.MinimumOccurs
}

type OperationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *OperationPropertyType) GongCopyBasicFields(to *OperationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type OrientableCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	Orientation string
}

func (from *OrientableCurveType) GongCopyBasicFields(to *OrientableCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.Orientation = from.Orientation
}

type OrientableSurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	Orientation string
}

func (from *OrientableSurfaceType) GongCopyBasicFields(to *OrientableSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.Orientation = from.Orientation
}

type PT_FreeText_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	CharacterString string

	MD_PixelOrientationCode string

	MD_ObligationCode string

	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type
}

func (from *PT_FreeText_PropertyType) GongCopyBasicFields(to *PT_FreeText_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.CharacterString = from.CharacterString
	to.MD_PixelOrientationCode = from.MD_PixelOrientationCode
	to.MD_ObligationCode = from.MD_ObligationCode
	to.MD_TopicCategoryCode = from.MD_TopicCategoryCode
}

type PT_FreeText_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *PT_FreeText_Type) GongCopyBasicFields(to *PT_FreeText_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type PT_LocaleContainer_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PT_LocaleContainer_PropertyType) GongCopyBasicFields(to *PT_LocaleContainer_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PT_LocaleContainer_Type_WOP struct {
	// insertion point

	Name string
}

func (from *PT_LocaleContainer_Type) GongCopyBasicFields(to *PT_LocaleContainer_Type) {
	// insertion point
	to.Name = from.Name
}

type PT_Locale_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PT_Locale_PropertyType) GongCopyBasicFields(to *PT_Locale_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PT_Locale_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *PT_Locale_Type) GongCopyBasicFields(to *PT_Locale_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type ParameterValueGroupType_WOP struct {
	// insertion point

	Name string
}

func (from *ParameterValueGroupType) GongCopyBasicFields(to *ParameterValueGroupType) {
	// insertion point
	to.Name = from.Name
}

type ParameterValueType_WOP struct {
	// insertion point

	Name string

	StringValue string

	IntegerValue int

	BooleanValue bool

	IntegerValueList string

	ValueFile string
}

func (from *ParameterValueType) GongCopyBasicFields(to *ParameterValueType) {
	// insertion point
	to.Name = from.Name
	to.StringValue = from.StringValue
	to.IntegerValue = from.IntegerValue
	to.BooleanValue = from.BooleanValue
	to.IntegerValueList = from.IntegerValueList
	to.ValueFile = from.ValueFile
}

type PassThroughOperationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PassThroughOperationPropertyType) GongCopyBasicFields(to *PassThroughOperationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PassThroughOperationType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string

	AggregationType string

	ModifiedCoordinate int
}

func (from *PassThroughOperationType) GongCopyBasicFields(to *PassThroughOperationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
	to.AggregationType = from.AggregationType
	to.ModifiedCoordinate = from.ModifiedCoordinate
}

type PointArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *PointArrayPropertyType) GongCopyBasicFields(to *PointArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type PointPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *PointPropertyType) GongCopyBasicFields(to *PointPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type PointType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *PointType) GongCopyBasicFields(to *PointType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type PolarCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PolarCSPropertyType) GongCopyBasicFields(to *PolarCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PolarCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *PolarCSType) GongCopyBasicFields(to *PolarCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
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

type PolygonType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *PolygonType) GongCopyBasicFields(to *PolygonType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type PrimeMeridianPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *PrimeMeridianPropertyType) GongCopyBasicFields(to *PrimeMeridianPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type PrimeMeridianType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *PrimeMeridianType) GongCopyBasicFields(to *PrimeMeridianType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type PriorityLocationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Null string

	Priority string
}

func (from *PriorityLocationPropertyType) GongCopyBasicFields(to *PriorityLocationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Null = from.Null
	to.Priority = from.Priority
}

type ProcedurePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ProcedurePropertyType) GongCopyBasicFields(to *ProcedurePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ProjectedCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ProjectedCRSPropertyType) GongCopyBasicFields(to *ProjectedCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type ProjectedCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *ProjectedCRSType) GongCopyBasicFields(to *ProjectedCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type Quantity_WOP struct {
	// insertion point

	Name string

	NilReason string

	EnclosedText string
}

func (from *Quantity) GongCopyBasicFields(to *Quantity) {
	// insertion point
	*to = *from
}

type QuantityExtentType_WOP struct {
	// insertion point

	Name string

	EnclosedText string
}

func (from *QuantityExtentType) GongCopyBasicFields(to *QuantityExtentType) {
	// insertion point
	*to = *from
}

type QuantityPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *QuantityPropertyType) GongCopyBasicFields(to *QuantityPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type RS_Identifier_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *RS_Identifier_PropertyType) GongCopyBasicFields(to *RS_Identifier_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type RS_Identifier_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *RS_Identifier_Type) GongCopyBasicFields(to *RS_Identifier_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type RS_ReferenceSystem_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *RS_ReferenceSystem_PropertyType) GongCopyBasicFields(to *RS_ReferenceSystem_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type RangeSetType_WOP struct {
	// insertion point

	Name string

	BooleanList string

	CountList string
}

func (from *RangeSetType) GongCopyBasicFields(to *RangeSetType) {
	// insertion point
	to.Name = from.Name
	to.BooleanList = from.BooleanList
	to.CountList = from.CountList
}

type Real_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Real_XSD float64
}

func (from *Real_PropertyType) GongCopyBasicFields(to *Real_PropertyType) {
	// insertion point
	*to = *from
}

type RecordType_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *RecordType_PropertyType) GongCopyBasicFields(to *RecordType_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type RecordType_Type_WOP struct {
	// insertion point

	Name string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	EnclosedText string
}

func (from *RecordType_Type) GongCopyBasicFields(to *RecordType_Type) {
	// insertion point
	*to = *from
}

type Record_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Record string
}

func (from *Record_PropertyType) GongCopyBasicFields(to *Record_PropertyType) {
	// insertion point
	*to = *from
}

type RectangleType_WOP struct {
	// insertion point

	Name string

	Interpolation Enum_SurfaceInterpolationType
}

func (from *RectangleType) GongCopyBasicFields(to *RectangleType) {
	// insertion point
	to.Name = from.Name
	to.Interpolation = from.Interpolation
}

type RectifiedGridType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	Dimension int

	AxisLabels_1 string

	AxisName string
}

func (from *RectifiedGridType) GongCopyBasicFields(to *RectifiedGridType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.Dimension = from.Dimension
	to.AxisLabels_1 = from.AxisLabels_1
	to.AxisName = from.AxisName
}

type ReferenceType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ReferenceType) GongCopyBasicFields(to *ReferenceType) {
	// insertion point
	*to = *from
}

type RelatedTimeType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool

	RelativePosition Enum_RelatedTimeType_relativePosition
}

func (from *RelatedTimeType) GongCopyBasicFields(to *RelatedTimeType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
	to.RelativePosition = from.RelativePosition
}

type ResultType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ResultType) GongCopyBasicFields(to *ResultType) {
	// insertion point
	*to = *from
}

type RingPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *RingPropertyType) GongCopyBasicFields(to *RingPropertyType) {
	// insertion point
	to.Name = from.Name
}

type RingType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *RingType) GongCopyBasicFields(to *RingType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type SC_CRS_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *SC_CRS_PropertyType) GongCopyBasicFields(to *SC_CRS_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type Scale_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *Scale_PropertyType) GongCopyBasicFields(to *Scale_PropertyType) {
	// insertion point
	*to = *from
}

type ScopedName_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *ScopedName_PropertyType) GongCopyBasicFields(to *ScopedName_PropertyType) {
	// insertion point
	*to = *from
}

type SecondDefiningParameter_WOP struct {
	// insertion point

	Name string
}

func (from *SecondDefiningParameter) GongCopyBasicFields(to *SecondDefiningParameter) {
	// insertion point
	to.Name = from.Name
}

type SecondDefiningParameter_1_WOP struct {
	// insertion point

	Name string

	IsSphere bool
}

func (from *SecondDefiningParameter_1) GongCopyBasicFields(to *SecondDefiningParameter_1) {
	// insertion point
	*to = *from
}

type SequenceRuleType_WOP struct {
	// insertion point

	Name string

	Order Enum_IncrementOrder

	AxisOrder string

	EnclosedText Enum_SequenceRuleEnumeration
}

func (from *SequenceRuleType) GongCopyBasicFields(to *SequenceRuleType) {
	// insertion point
	*to = *from
}

type ShellPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *ShellPropertyType) GongCopyBasicFields(to *ShellPropertyType) {
	// insertion point
	to.Name = from.Name
}

type ShellType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string

	AggregationType string
}

func (from *ShellType) GongCopyBasicFields(to *ShellType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
	to.AggregationType = from.AggregationType
}

type SingleCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *SingleCRSPropertyType) GongCopyBasicFields(to *SingleCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type SingleOperationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *SingleOperationPropertyType) GongCopyBasicFields(to *SingleOperationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type SolidArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *SolidArrayPropertyType) GongCopyBasicFields(to *SolidArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type SolidPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *SolidPropertyType) GongCopyBasicFields(to *SolidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type SolidType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *SolidType) GongCopyBasicFields(to *SolidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type SphereType_WOP struct {
	// insertion point

	Name string

	AggregationType string

	Rows int

	Columns int

	HorizontalCurveType Enum_CurveInterpolationType

	VerticalCurveType Enum_CurveInterpolationType
}

func (from *SphereType) GongCopyBasicFields(to *SphereType) {
	// insertion point
	to.Name = from.Name
	to.AggregationType = from.AggregationType
	to.Rows = from.Rows
	to.Columns = from.Columns
	to.HorizontalCurveType = from.HorizontalCurveType
	to.VerticalCurveType = from.VerticalCurveType
}

type SphericalCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *SphericalCSPropertyType) GongCopyBasicFields(to *SphericalCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type SphericalCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *SphericalCSType) GongCopyBasicFields(to *SphericalCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type StringOrRefType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	EnclosedText string
}

func (from *StringOrRefType) GongCopyBasicFields(to *StringOrRefType) {
	// insertion point
	*to = *from
}

type SurfaceArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *SurfaceArrayPropertyType) GongCopyBasicFields(to *SurfaceArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type SurfacePatchArrayPropertyType_WOP struct {
	// insertion point

	Name string
}

func (from *SurfacePatchArrayPropertyType) GongCopyBasicFields(to *SurfacePatchArrayPropertyType) {
	// insertion point
	to.Name = from.Name
}

type SurfacePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *SurfacePropertyType) GongCopyBasicFields(to *SurfacePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type SurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *SurfaceType) GongCopyBasicFields(to *SurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type TM_PeriodDuration_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	TM_PeriodDuration string
}

func (from *TM_PeriodDuration_PropertyType) GongCopyBasicFields(to *TM_PeriodDuration_PropertyType) {
	// insertion point
	*to = *from
}

type TM_Primitive_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TM_Primitive_PropertyType) GongCopyBasicFields(to *TM_Primitive_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TargetPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TargetPropertyType) GongCopyBasicFields(to *TargetPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TemporalCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TemporalCRSPropertyType) GongCopyBasicFields(to *TemporalCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TemporalCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *TemporalCRSType) GongCopyBasicFields(to *TemporalCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type TemporalCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TemporalCSPropertyType) GongCopyBasicFields(to *TemporalCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TemporalCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *TemporalCSType) GongCopyBasicFields(to *TemporalCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type TemporalDatumBaseType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *TemporalDatumBaseType) GongCopyBasicFields(to *TemporalDatumBaseType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type TemporalDatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TemporalDatumPropertyType) GongCopyBasicFields(to *TemporalDatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TemporalDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string

	Origin string
}

func (from *TemporalDatumType) GongCopyBasicFields(to *TemporalDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
	to.Origin = from.Origin
}

type TimeCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimeCSPropertyType) GongCopyBasicFields(to *TimeCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TimeCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *TimeCSType) GongCopyBasicFields(to *TimeCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type TimeCalendarEraPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimeCalendarEraPropertyType) GongCopyBasicFields(to *TimeCalendarEraPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TimeCalendarEraType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	ReferenceDate string

	JulianReference string
}

func (from *TimeCalendarEraType) GongCopyBasicFields(to *TimeCalendarEraType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.ReferenceDate = from.ReferenceDate
	to.JulianReference = from.JulianReference
}

type TimeCalendarPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimeCalendarPropertyType) GongCopyBasicFields(to *TimeCalendarPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TimeCalendarType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	DomainOfValidity string
}

func (from *TimeCalendarType) GongCopyBasicFields(to *TimeCalendarType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.DomainOfValidity = from.DomainOfValidity
}

type TimeClockPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimeClockPropertyType) GongCopyBasicFields(to *TimeClockPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TimeClockType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	DomainOfValidity string

	ReferenceTime string

	UtcReference string
}

func (from *TimeClockType) GongCopyBasicFields(to *TimeClockType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.DomainOfValidity = from.DomainOfValidity
	to.ReferenceTime = from.ReferenceTime
	to.UtcReference = from.UtcReference
}

type TimeCoordinateSystemType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	DomainOfValidity string
}

func (from *TimeCoordinateSystemType) GongCopyBasicFields(to *TimeCoordinateSystemType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.DomainOfValidity = from.DomainOfValidity
}

type TimeEdgePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimeEdgePropertyType) GongCopyBasicFields(to *TimeEdgePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimeEdgeType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *TimeEdgeType) GongCopyBasicFields(to *TimeEdgeType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type TimeInstantPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimeInstantPropertyType) GongCopyBasicFields(to *TimeInstantPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimeInstantType_WOP struct {
	// insertion point

	Name string

	Id string

	Frame string
}

func (from *TimeInstantType) GongCopyBasicFields(to *TimeInstantType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Frame = from.Frame
}

type TimeIntervalLengthType_WOP struct {
	// insertion point

	Name string

	Unit string

	Radix int

	Factor int

	EnclosedText string
}

func (from *TimeIntervalLengthType) GongCopyBasicFields(to *TimeIntervalLengthType) {
	// insertion point
	*to = *from
}

type TimeNodePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimeNodePropertyType) GongCopyBasicFields(to *TimeNodePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimeNodeType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *TimeNodeType) GongCopyBasicFields(to *TimeNodeType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type TimeOrdinalEraPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimeOrdinalEraPropertyType) GongCopyBasicFields(to *TimeOrdinalEraPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TimeOrdinalEraType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *TimeOrdinalEraType) GongCopyBasicFields(to *TimeOrdinalEraType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type TimeOrdinalReferenceSystemType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	DomainOfValidity string
}

func (from *TimeOrdinalReferenceSystemType) GongCopyBasicFields(to *TimeOrdinalReferenceSystemType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.DomainOfValidity = from.DomainOfValidity
}

type TimePeriodPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimePeriodPropertyType) GongCopyBasicFields(to *TimePeriodPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimePeriodType_WOP struct {
	// insertion point

	Name string

	Id string

	Frame string

	Duration string
}

func (from *TimePeriodType) GongCopyBasicFields(to *TimePeriodType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Frame = from.Frame
	to.Duration = from.Duration
}

type TimePositionType_WOP struct {
	// insertion point

	Name string

	Frame string

	CalendarEraName string

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

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimePrimitivePropertyType) GongCopyBasicFields(to *TimePrimitivePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimeReferenceSystemType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	DomainOfValidity string
}

func (from *TimeReferenceSystemType) GongCopyBasicFields(to *TimeReferenceSystemType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.DomainOfValidity = from.DomainOfValidity
}

type TimeTopologyComplexPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimeTopologyComplexPropertyType) GongCopyBasicFields(to *TimeTopologyComplexPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TimeTopologyComplexType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *TimeTopologyComplexType) GongCopyBasicFields(to *TimeTopologyComplexType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type TimeTopologyPrimitivePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TimeTopologyPrimitivePropertyType) GongCopyBasicFields(to *TimeTopologyPrimitivePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TinType_WOP struct {
	// insertion point

	Name string

	Id string

	SrsName string

	SrsDimension int

	AxisLabels string

	UomLabels string
}

func (from *TinType) GongCopyBasicFields(to *TinType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SrsName = from.SrsName
	to.SrsDimension = from.SrsDimension
	to.AxisLabels = from.AxisLabels
	to.UomLabels = from.UomLabels
}

type TopoComplexPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TopoComplexPropertyType) GongCopyBasicFields(to *TopoComplexPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TopoComplexType_WOP struct {
	// insertion point

	Name string

	Id string

	IsMaximal bool

	AggregationType string
}

func (from *TopoComplexType) GongCopyBasicFields(to *TopoComplexType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.IsMaximal = from.IsMaximal
	to.AggregationType = from.AggregationType
}

type TopoCurvePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *TopoCurvePropertyType) GongCopyBasicFields(to *TopoCurvePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type TopoCurveType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *TopoCurveType) GongCopyBasicFields(to *TopoCurveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type TopoPointPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *TopoPointPropertyType) GongCopyBasicFields(to *TopoPointPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type TopoPointType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *TopoPointType) GongCopyBasicFields(to *TopoPointType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
}

type TopoPrimitiveArrayAssociationType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *TopoPrimitiveArrayAssociationType) GongCopyBasicFields(to *TopoPrimitiveArrayAssociationType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type TopoPrimitiveMemberType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TopoPrimitiveMemberType) GongCopyBasicFields(to *TopoPrimitiveMemberType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TopoSolidPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool
}

func (from *TopoSolidPropertyType) GongCopyBasicFields(to *TopoSolidPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
}

type TopoSolidType_WOP struct {
	// insertion point

	Name string

	Id string

	Universal bool

	AggregationType string
}

func (from *TopoSolidType) GongCopyBasicFields(to *TopoSolidType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Universal = from.Universal
	to.AggregationType = from.AggregationType
}

type TopoSurfacePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *TopoSurfacePropertyType) GongCopyBasicFields(to *TopoSurfacePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type TopoSurfaceType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *TopoSurfaceType) GongCopyBasicFields(to *TopoSurfaceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type TopoVolumePropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *TopoVolumePropertyType) GongCopyBasicFields(to *TopoVolumePropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
}

type TopoVolumeType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string
}

func (from *TopoVolumeType) GongCopyBasicFields(to *TopoVolumeType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
}

type TransformationPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TransformationPropertyType) GongCopyBasicFields(to *TransformationPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TransformationType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	OperationVersion string
}

func (from *TransformationType) GongCopyBasicFields(to *TransformationType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.OperationVersion = from.OperationVersion
}

type TriangleType_WOP struct {
	// insertion point

	Name string

	Interpolation Enum_SurfaceInterpolationType
}

func (from *TriangleType) GongCopyBasicFields(to *TriangleType) {
	// insertion point
	to.Name = from.Name
	to.Interpolation = from.Interpolation
}

type TypeName_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TypeName_PropertyType) GongCopyBasicFields(to *TypeName_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type TypeName_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *TypeName_Type) GongCopyBasicFields(to *TypeName_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type URL_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	URL string
}

func (from *URL_PropertyType) GongCopyBasicFields(to *URL_PropertyType) {
	// insertion point
	*to = *from
}

type UnitDefinitionType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string
}

func (from *UnitDefinitionType) GongCopyBasicFields(to *UnitDefinitionType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
}

type UnitOfMeasureType_WOP struct {
	// insertion point

	Name string

	Uom string
}

func (from *UnitOfMeasureType) GongCopyBasicFields(to *UnitOfMeasureType) {
	// insertion point
	*to = *from
}

type UnitOfMeasure_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UnitOfMeasure_PropertyType) GongCopyBasicFields(to *UnitOfMeasure_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UnlimitedInteger_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *UnlimitedInteger_PropertyType) GongCopyBasicFields(to *UnlimitedInteger_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type UnlimitedInteger_Type_WOP struct {
	// insertion point

	Name string

	IsInfinite bool

	EnclosedText int
}

func (from *UnlimitedInteger_Type) GongCopyBasicFields(to *UnlimitedInteger_Type) {
	// insertion point
	*to = *from
}

type UomAngle_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomAngle_PropertyType) GongCopyBasicFields(to *UomAngle_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomArea_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomArea_PropertyType) GongCopyBasicFields(to *UomArea_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomLength_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomLength_PropertyType) GongCopyBasicFields(to *UomLength_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomScale_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomScale_PropertyType) GongCopyBasicFields(to *UomScale_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomTime_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomTime_PropertyType) GongCopyBasicFields(to *UomTime_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomVelocity_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomVelocity_PropertyType) GongCopyBasicFields(to *UomVelocity_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UomVolume_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UomVolume_PropertyType) GongCopyBasicFields(to *UomVolume_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UserDefinedCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *UserDefinedCSPropertyType) GongCopyBasicFields(to *UserDefinedCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type UserDefinedCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *UserDefinedCSType) GongCopyBasicFields(to *UserDefinedCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type ValueArrayPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	BooleanList string

	CountList string

	CountExtent string

	Null string
}

func (from *ValueArrayPropertyType) GongCopyBasicFields(to *ValueArrayPropertyType) {
	// insertion point
	to.Name = from.Name
	to.Owns = from.Owns
	to.BooleanList = from.BooleanList
	to.CountList = from.CountList
	to.CountExtent = from.CountExtent
	to.Null = from.Null
}

type ValueArrayType_WOP struct {
	// insertion point

	Name string

	Id string

	AggregationType string

	CodeSpace string

	Uom string
}

func (from *ValueArrayType) GongCopyBasicFields(to *ValueArrayType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.AggregationType = from.AggregationType
	to.CodeSpace = from.CodeSpace
	to.Uom = from.Uom
}

type ValuePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string

	Owns bool

	BooleanList string

	CountList string

	CountExtent string

	Null string
}

func (from *ValuePropertyType) GongCopyBasicFields(to *ValuePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.Owns = from.Owns
	to.BooleanList = from.BooleanList
	to.CountList = from.CountList
	to.CountExtent = from.CountExtent
	to.Null = from.Null
}

type VectorType_WOP struct {
	// insertion point

	Name string

	EnclosedText string
}

func (from *VectorType) GongCopyBasicFields(to *VectorType) {
	// insertion point
	*to = *from
}

type VerticalCRSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *VerticalCRSPropertyType) GongCopyBasicFields(to *VerticalCRSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type VerticalCRSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string
}

func (from *VerticalCRSType) GongCopyBasicFields(to *VerticalCRSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
}

type VerticalCSPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *VerticalCSPropertyType) GongCopyBasicFields(to *VerticalCSPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type VerticalCSType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	AggregationType string
}

func (from *VerticalCSType) GongCopyBasicFields(to *VerticalCSType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.AggregationType = from.AggregationType
}

type VerticalDatumPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *VerticalDatumPropertyType) GongCopyBasicFields(to *VerticalDatumPropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.RemoteSchema = from.RemoteSchema
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type VerticalDatumType_WOP struct {
	// insertion point

	Name string

	Id string

	Remarks string

	Scope string

	RealizationEpoch string
}

func (from *VerticalDatumType) GongCopyBasicFields(to *VerticalDatumType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Remarks = from.Remarks
	to.Scope = from.Scope
	to.RealizationEpoch = from.RealizationEpoch
}

// end of insertion point
