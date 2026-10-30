// generated code - do not edit
package gml

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/measures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *basictypes.CodeType
var _ *gcobase.AbstractObject_Type
var _ *measures.LengthType
var _ *xlink.Simple

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "CurveInterpolationType"
type Enum_CurveInterpolationType string

const (
	Enum_CurveInterpolationType_Linear Enum_CurveInterpolationType = "linear"
	Enum_CurveInterpolationType_Geodesic Enum_CurveInterpolationType = "geodesic"
	Enum_CurveInterpolationType_CircularArc3Points Enum_CurveInterpolationType = "circularArc3Points"
	Enum_CurveInterpolationType_CircularArc2PointWithBulge Enum_CurveInterpolationType = "circularArc2PointWithBulge"
	Enum_CurveInterpolationType_CircularArcCenterPointWithRadius Enum_CurveInterpolationType = "circularArcCenterPointWithRadius"
	Enum_CurveInterpolationType_Elliptical Enum_CurveInterpolationType = "elliptical"
	Enum_CurveInterpolationType_Clothoid Enum_CurveInterpolationType = "clothoid"
	Enum_CurveInterpolationType_Conic Enum_CurveInterpolationType = "conic"
	Enum_CurveInterpolationType_PolynomialSpline Enum_CurveInterpolationType = "polynomialSpline"
	Enum_CurveInterpolationType_CubicSpline Enum_CurveInterpolationType = "cubicSpline"
	Enum_CurveInterpolationType_RationalSpline Enum_CurveInterpolationType = "rationalSpline"
)

// From xsd simple type with enumerate restriction "KnotTypesType"
type Enum_KnotTypesType string

const (
	Enum_KnotTypesType_Uniform Enum_KnotTypesType = "uniform"
	Enum_KnotTypesType_QuasiUniform Enum_KnotTypesType = "quasiUniform"
	Enum_KnotTypesType_PiecewiseBezier Enum_KnotTypesType = "piecewiseBezier"
)

// From xsd simple type with enumerate restriction "SurfaceInterpolationType"
type Enum_SurfaceInterpolationType string

const (
	Enum_SurfaceInterpolationType_None Enum_SurfaceInterpolationType = "none"
	Enum_SurfaceInterpolationType_Planar Enum_SurfaceInterpolationType = "planar"
	Enum_SurfaceInterpolationType_Spherical Enum_SurfaceInterpolationType = "spherical"
	Enum_SurfaceInterpolationType_Elliptical Enum_SurfaceInterpolationType = "elliptical"
	Enum_SurfaceInterpolationType_Conic Enum_SurfaceInterpolationType = "conic"
	Enum_SurfaceInterpolationType_Tin Enum_SurfaceInterpolationType = "tin"
	Enum_SurfaceInterpolationType_ParametricCurve Enum_SurfaceInterpolationType = "parametricCurve"
	Enum_SurfaceInterpolationType_PolynomialSpline Enum_SurfaceInterpolationType = "polynomialSpline"
	Enum_SurfaceInterpolationType_RationalSpline Enum_SurfaceInterpolationType = "rationalSpline"
	Enum_SurfaceInterpolationType_TriangulatedSpline Enum_SurfaceInterpolationType = "triangulatedSpline"
)

// From xsd simple type with enumerate restriction "AggregationType"
type Enum_AggregationType string

const (
	Enum_AggregationType_Set Enum_AggregationType = "set"
	Enum_AggregationType_Bag Enum_AggregationType = "bag"
	Enum_AggregationType_Sequence Enum_AggregationType = "sequence"
	Enum_AggregationType_Array Enum_AggregationType = "array"
	Enum_AggregationType_Record Enum_AggregationType = "record"
	Enum_AggregationType_Table Enum_AggregationType = "table"
)

// From xsd simple type with enumerate restriction "RelatedTimeType_relativePosition"
type Enum_RelatedTimeType_relativePosition string

const (
	Enum_RelatedTimeType_relativePosition_Before Enum_RelatedTimeType_relativePosition = "Before"
	Enum_RelatedTimeType_relativePosition_After Enum_RelatedTimeType_relativePosition = "After"
	Enum_RelatedTimeType_relativePosition_Begins Enum_RelatedTimeType_relativePosition = "Begins"
	Enum_RelatedTimeType_relativePosition_Ends Enum_RelatedTimeType_relativePosition = "Ends"
	Enum_RelatedTimeType_relativePosition_During Enum_RelatedTimeType_relativePosition = "During"
	Enum_RelatedTimeType_relativePosition_Equals Enum_RelatedTimeType_relativePosition = "Equals"
	Enum_RelatedTimeType_relativePosition_Contains Enum_RelatedTimeType_relativePosition = "Contains"
	Enum_RelatedTimeType_relativePosition_Overlaps Enum_RelatedTimeType_relativePosition = "Overlaps"
	Enum_RelatedTimeType_relativePosition_Meets Enum_RelatedTimeType_relativePosition = "Meets"
	Enum_RelatedTimeType_relativePosition_OverlappedBy Enum_RelatedTimeType_relativePosition = "OverlappedBy"
	Enum_RelatedTimeType_relativePosition_MetBy Enum_RelatedTimeType_relativePosition = "MetBy"
	Enum_RelatedTimeType_relativePosition_BegunBy Enum_RelatedTimeType_relativePosition = "BegunBy"
	Enum_RelatedTimeType_relativePosition_EndedBy Enum_RelatedTimeType_relativePosition = "EndedBy"
)

// From xsd simple type with enumerate restriction "TimeIndeterminateValueType"
type Enum_TimeIndeterminateValueType string

const (
	Enum_TimeIndeterminateValueType_After Enum_TimeIndeterminateValueType = "after"
	Enum_TimeIndeterminateValueType_Before Enum_TimeIndeterminateValueType = "before"
	Enum_TimeIndeterminateValueType_Now Enum_TimeIndeterminateValueType = "now"
	Enum_TimeIndeterminateValueType_Unknown Enum_TimeIndeterminateValueType = "unknown"
)

// From xsd simple type with enumerate restriction "TimeUnitType_Member_1"
type Enum_TimeUnitType_Member_1 string

const (
	Enum_TimeUnitType_Member_1_Year Enum_TimeUnitType_Member_1 = "year"
	Enum_TimeUnitType_Member_1_Month Enum_TimeUnitType_Member_1 = "month"
	Enum_TimeUnitType_Member_1_Day Enum_TimeUnitType_Member_1 = "day"
	Enum_TimeUnitType_Member_1_Hour Enum_TimeUnitType_Member_1 = "hour"
	Enum_TimeUnitType_Member_1_Minute Enum_TimeUnitType_Member_1 = "minute"
	Enum_TimeUnitType_Member_1_Second Enum_TimeUnitType_Member_1 = "second"
)

// From xsd simple type with enumerate restriction "CompassPointEnumeration"
type Enum_CompassPointEnumeration string

const (
	Enum_CompassPointEnumeration_N Enum_CompassPointEnumeration = "N"
	Enum_CompassPointEnumeration_NNE Enum_CompassPointEnumeration = "NNE"
	Enum_CompassPointEnumeration_NE Enum_CompassPointEnumeration = "NE"
	Enum_CompassPointEnumeration_ENE Enum_CompassPointEnumeration = "ENE"
	Enum_CompassPointEnumeration_E Enum_CompassPointEnumeration = "E"
	Enum_CompassPointEnumeration_ESE Enum_CompassPointEnumeration = "ESE"
	Enum_CompassPointEnumeration_SE Enum_CompassPointEnumeration = "SE"
	Enum_CompassPointEnumeration_SSE Enum_CompassPointEnumeration = "SSE"
	Enum_CompassPointEnumeration_S Enum_CompassPointEnumeration = "S"
	Enum_CompassPointEnumeration_SSW Enum_CompassPointEnumeration = "SSW"
	Enum_CompassPointEnumeration_SW Enum_CompassPointEnumeration = "SW"
	Enum_CompassPointEnumeration_WSW Enum_CompassPointEnumeration = "WSW"
	Enum_CompassPointEnumeration_W Enum_CompassPointEnumeration = "W"
	Enum_CompassPointEnumeration_WNW Enum_CompassPointEnumeration = "WNW"
	Enum_CompassPointEnumeration_NW Enum_CompassPointEnumeration = "NW"
	Enum_CompassPointEnumeration_NNW Enum_CompassPointEnumeration = "NNW"
)

// From xsd simple type with enumerate restriction "SequenceRuleEnumeration"
type Enum_SequenceRuleEnumeration string

const (
	Enum_SequenceRuleEnumeration_Linear Enum_SequenceRuleEnumeration = "Linear"
	Enum_SequenceRuleEnumeration_Boustrophedonic Enum_SequenceRuleEnumeration = "Boustrophedonic"
	Enum_SequenceRuleEnumeration_Cantor_diagonal Enum_SequenceRuleEnumeration = "Cantor-diagonal"
	Enum_SequenceRuleEnumeration_Spiral Enum_SequenceRuleEnumeration = "Spiral"
	Enum_SequenceRuleEnumeration_Morton Enum_SequenceRuleEnumeration = "Morton"
	Enum_SequenceRuleEnumeration_Hilbert Enum_SequenceRuleEnumeration = "Hilbert"
)

// From xsd simple type with enumerate restriction "MD_TopicCategoryCode_Type"
type Enum_MD_TopicCategoryCode_Type string

const (
	Enum_MD_TopicCategoryCode_Type_Farming Enum_MD_TopicCategoryCode_Type = "farming"
	Enum_MD_TopicCategoryCode_Type_Biota Enum_MD_TopicCategoryCode_Type = "biota"
	Enum_MD_TopicCategoryCode_Type_Boundaries Enum_MD_TopicCategoryCode_Type = "boundaries"
	Enum_MD_TopicCategoryCode_Type_ClimatologyMeteorologyAtmosphere Enum_MD_TopicCategoryCode_Type = "climatologyMeteorologyAtmosphere"
	Enum_MD_TopicCategoryCode_Type_Economy Enum_MD_TopicCategoryCode_Type = "economy"
	Enum_MD_TopicCategoryCode_Type_Elevation Enum_MD_TopicCategoryCode_Type = "elevation"
	Enum_MD_TopicCategoryCode_Type_Environment Enum_MD_TopicCategoryCode_Type = "environment"
	Enum_MD_TopicCategoryCode_Type_GeoscientificInformation Enum_MD_TopicCategoryCode_Type = "geoscientificInformation"
	Enum_MD_TopicCategoryCode_Type_Health Enum_MD_TopicCategoryCode_Type = "health"
	Enum_MD_TopicCategoryCode_Type_ImageryBaseMapsEarthCover Enum_MD_TopicCategoryCode_Type = "imageryBaseMapsEarthCover"
	Enum_MD_TopicCategoryCode_Type_IntelligenceMilitary Enum_MD_TopicCategoryCode_Type = "intelligenceMilitary"
	Enum_MD_TopicCategoryCode_Type_InlandWaters Enum_MD_TopicCategoryCode_Type = "inlandWaters"
	Enum_MD_TopicCategoryCode_Type_Location Enum_MD_TopicCategoryCode_Type = "location"
	Enum_MD_TopicCategoryCode_Type_Oceans Enum_MD_TopicCategoryCode_Type = "oceans"
	Enum_MD_TopicCategoryCode_Type_PlanningCadastre Enum_MD_TopicCategoryCode_Type = "planningCadastre"
	Enum_MD_TopicCategoryCode_Type_Society Enum_MD_TopicCategoryCode_Type = "society"
	Enum_MD_TopicCategoryCode_Type_Structure Enum_MD_TopicCategoryCode_Type = "structure"
	Enum_MD_TopicCategoryCode_Type_Transportation Enum_MD_TopicCategoryCode_Type = "transportation"
	Enum_MD_TopicCategoryCode_Type_UtilitiesCommunication Enum_MD_TopicCategoryCode_Type = "utilitiesCommunication"
)

// From xsd simple type with enumerate restriction "SuccessionType"
type Enum_SuccessionType string

const (
	Enum_SuccessionType_Substitution Enum_SuccessionType = "substitution"
	Enum_SuccessionType_Division Enum_SuccessionType = "division"
	Enum_SuccessionType_Fusion Enum_SuccessionType = "fusion"
	Enum_SuccessionType_Initiation Enum_SuccessionType = "initiation"
)

// From xsd simple type with enumerate restriction "DegreesType_direction"
type Enum_DegreesType_direction string

const (
	Enum_DegreesType_direction_N Enum_DegreesType_direction = "N"
	Enum_DegreesType_direction_E Enum_DegreesType_direction = "E"
	Enum_DegreesType_direction_S Enum_DegreesType_direction = "S"
	Enum_DegreesType_direction_W Enum_DegreesType_direction = "W"
	Enum_DegreesType_direction_N__ Enum_DegreesType_direction = "+"
	Enum_DegreesType_direction_N___1 Enum_DegreesType_direction = "-"
)

// From xsd simple type with enumerate restriction "IncrementOrder"
type Enum_IncrementOrder string

const (
	Enum_IncrementOrder_N__x_y Enum_IncrementOrder = "+x+y"
	Enum_IncrementOrder_N__y_x Enum_IncrementOrder = "+y+x"
	Enum_IncrementOrder_N__x_y_1 Enum_IncrementOrder = "+x-y"
	Enum_IncrementOrder_N__x_y_2 Enum_IncrementOrder = "-x-y"
)

// insertion point for gongstructs declarations

// Group_dynamicProperties UnNamed source named group "dynamicProperties"
type Group_dynamicProperties struct {

	// insertion point for fields

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "history" of complex type HistoryPropertyType
	History *HistoryPropertyType `xml:"history,omitempty"`

	// generated from element "track" of complex type HistoryPropertyType
	Track *HistoryPropertyType `xml:"track,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *StringOrRefType `xml:"dataSource,omitempty"`

	// generated from element "dataSourceReference" of complex type ReferenceType
	DataSourceReference *ReferenceType `xml:"dataSourceReference,omitempty"`
}

// DynamicFeatureType Named source named complex type "DynamicFeatureType"
type DynamicFeatureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from group "gml:dynamicProperties" with order 128 depth 1
	Group_dynamicProperties
}

// DynamicFeatureCollectionType Named source named complex type "DynamicFeatureCollectionType"
type DynamicFeatureCollectionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from group "gml:dynamicProperties" with order 128 depth 1
	Group_dynamicProperties

	// generated from element "dynamicMembers" of complex type DynamicFeatureMemberType
	DynamicMembers *DynamicFeatureMemberType `xml:"dynamicMembers,omitempty"`
}

// DynamicFeatureMemberType Named source named complex type "DynamicFeatureMemberType"
type DynamicFeatureMemberType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature []*DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection []*DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`
}

// AbstractTimeSliceType Named source named complex type "AbstractTimeSliceType"
type AbstractTimeSliceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *StringOrRefType `xml:"dataSource,omitempty"`
}

// HistoryPropertyType Named source named complex type "HistoryPropertyType"
type HistoryPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MovingObjectStatus" of complex type MovingObjectStatusType
	MovingObjectStatus []*MovingObjectStatusType `xml:"MovingObjectStatus,omitempty"`
}

// AbstractFeatureType Named source named complex type "AbstractFeatureType"
// The basic feature model is given by the gml:AbstractFeatureType.
// The content model for gml:AbstractFeatureType adds two specific properties suitable for geographic features to the content model defined in gml:AbstractGMLType. 
// The value of the gml:boundedBy property describes an envelope that encloses the entire feature instance, and is primarily useful for supporting rapid searching for features that occur in a particular location. 
// The value of the gml:location property describes the extent, position or relative location of the feature.
type AbstractFeatureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`
}

// FeaturePropertyType Named source named complex type "FeaturePropertyType"
type FeaturePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature *DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection *DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`

	// generated from element "MultiPointCoverage" of complex type DiscreteCoverageType
	MultiPointCoverage *DiscreteCoverageType `xml:"MultiPointCoverage,omitempty"`

	// generated from element "MultiCurveCoverage" of complex type DiscreteCoverageType
	MultiCurveCoverage *DiscreteCoverageType `xml:"MultiCurveCoverage,omitempty"`

	// generated from element "MultiSurfaceCoverage" of complex type DiscreteCoverageType
	MultiSurfaceCoverage *DiscreteCoverageType `xml:"MultiSurfaceCoverage,omitempty"`

	// generated from element "MultiSolidCoverage" of complex type DiscreteCoverageType
	MultiSolidCoverage *DiscreteCoverageType `xml:"MultiSolidCoverage,omitempty"`

	// generated from element "GridCoverage" of complex type DiscreteCoverageType
	GridCoverage *DiscreteCoverageType `xml:"GridCoverage,omitempty"`

	// generated from element "RectifiedGridCoverage" of complex type DiscreteCoverageType
	RectifiedGridCoverage *DiscreteCoverageType `xml:"RectifiedGridCoverage,omitempty"`

	// generated from element "Observation" of complex type ObservationType
	Observation *ObservationType `xml:"Observation,omitempty"`

	// generated from element "DirectedObservation" of complex type DirectedObservationType
	DirectedObservation *DirectedObservationType `xml:"DirectedObservation,omitempty"`

	// generated from element "DirectedObservationAtDistance" of complex type DirectedObservationAtDistanceType
	DirectedObservationAtDistance *DirectedObservationAtDistanceType `xml:"DirectedObservationAtDistance,omitempty"`

	// generated from element "FeatureCollection" of complex type FeatureCollectionType
	FeatureCollection *FeatureCollectionType `xml:"FeatureCollection,omitempty"`
}

// BoundingShapeType Named source named complex type "BoundingShapeType"
type BoundingShapeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Envelope" of complex type EnvelopeType
	Envelope *EnvelopeType `xml:"Envelope,omitempty"`

	// generated from element "EnvelopeWithTimePeriod" of complex type EnvelopeWithTimePeriodType
	EnvelopeWithTimePeriod *EnvelopeWithTimePeriodType `xml:"EnvelopeWithTimePeriod,omitempty"`

	// generated from element "Null" of simple type gml:NilReasonType
	Null string `xml:"Null,omitempty"`
}

// EnvelopeWithTimePeriodType Named source named complex type "EnvelopeWithTimePeriodType"
type EnvelopeWithTimePeriodType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "lowerCorner" of complex type DirectPositionType
	LowerCorner *DirectPositionType `xml:"lowerCorner,omitempty"`

	// generated from element "upperCorner" of complex type DirectPositionType
	UpperCorner *DirectPositionType `xml:"upperCorner,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from element "beginPosition" of complex type TimePositionType
	BeginPosition *TimePositionType `xml:"beginPosition,omitempty"`

	// generated from element "endPosition" of complex type TimePositionType
	EndPosition *TimePositionType `xml:"endPosition,omitempty"`
}

// AbstractFeatureMemberType Named source named complex type "AbstractFeatureMemberType"
// To create a collection of GML features, a property type shall be derived by extension from gml:AbstractFeatureMemberType.
// By default, this abstract property type does not imply any ownership of the features in the collection. The owns attribute of gml:OwnershipAttributeGroup may be used on a property element instance to assert ownership of a feature in the collection. A collection shall not own a feature already owned by another object.
type AbstractFeatureMemberType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup
}

// AbstractGeometricAggregateType Named source named complex type "AbstractGeometricAggregateType"
type AbstractGeometricAggregateType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup
}

// MultiGeometryType Named source named complex type "MultiGeometryType"
type MultiGeometryType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "geometryMember" of complex type GeometryPropertyType
	GeometryMember []*GeometryPropertyType `xml:"geometryMember,omitempty"`

	// generated from element "geometryMembers" of complex type GeometryArrayPropertyType
	GeometryMembers *GeometryArrayPropertyType `xml:"geometryMembers,omitempty"`
}

// MultiGeometryPropertyType Named source named complex type "MultiGeometryPropertyType"
// A property that has a geometric aggregate as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type MultiGeometryPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`
}

// MultiPointType Named source named complex type "MultiPointType"
type MultiPointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "pointMember" of complex type PointPropertyType
	PointMember []*PointPropertyType `xml:"pointMember,omitempty"`

	// generated from element "pointMembers" of complex type PointArrayPropertyType
	PointMembers *PointArrayPropertyType `xml:"pointMembers,omitempty"`
}

// MultiPointPropertyType Named source named complex type "MultiPointPropertyType"
// A property that has a collection of points as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type MultiPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`
}

// MultiCurveType Named source named complex type "MultiCurveType"
type MultiCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "curveMember" of complex type CurvePropertyType
	CurveMember []*CurvePropertyType `xml:"curveMember,omitempty"`

	// generated from element "curveMembers" of complex type CurveArrayPropertyType
	CurveMembers *CurveArrayPropertyType `xml:"curveMembers,omitempty"`
}

// MultiCurvePropertyType Named source named complex type "MultiCurvePropertyType"
// A property that has a collection of curves as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type MultiCurvePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`
}

// MultiSurfaceType Named source named complex type "MultiSurfaceType"
type MultiSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "surfaceMember" of complex type SurfacePropertyType
	SurfaceMember []*SurfacePropertyType `xml:"surfaceMember,omitempty"`

	// generated from element "surfaceMembers" of complex type SurfaceArrayPropertyType
	SurfaceMembers *SurfaceArrayPropertyType `xml:"surfaceMembers,omitempty"`
}

// MultiSurfacePropertyType Named source named complex type "MultiSurfacePropertyType"
// A property that has a collection of surfaces as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type MultiSurfacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`
}

// MultiSolidType Named source named complex type "MultiSolidType"
type MultiSolidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "solidMember" of complex type SolidPropertyType
	SolidMember []*SolidPropertyType `xml:"solidMember,omitempty"`

	// generated from element "solidMembers" of complex type SolidArrayPropertyType
	SolidMembers *SolidArrayPropertyType `xml:"solidMembers,omitempty"`
}

// MultiSolidPropertyType Named source named complex type "MultiSolidPropertyType"
// A property that has a collection of solids as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type MultiSolidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`
}

// CurveType Named source named complex type "CurveType"
type CurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "segments" of complex type CurveSegmentArrayPropertyType
	Segments *CurveSegmentArrayPropertyType `xml:"segments,omitempty"`
}

// OrientableCurveType Named source named complex type "OrientableCurveType"
type OrientableCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from element "baseCurve" of complex type CurvePropertyType
	BaseCurve *CurvePropertyType `xml:"baseCurve,omitempty"`
}

// AbstractCurveSegmentType Named source named complex type "AbstractCurveSegmentType"
type AbstractCurveSegmentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`
}

// CurveSegmentArrayPropertyType Named source named complex type "CurveSegmentArrayPropertyType"
// gml:CurveSegmentArrayPropertyType is a container for an array of curve segments.
type CurveSegmentArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "LineStringSegment" of complex type LineStringSegmentType
	LineStringSegment []*LineStringSegmentType `xml:"LineStringSegment,omitempty"`

	// generated from element "ArcString" of complex type ArcStringType
	ArcString []*ArcStringType `xml:"ArcString,omitempty"`

	// generated from element "Arc" of complex type ArcType
	Arc []*ArcType `xml:"Arc,omitempty"`

	// generated from element "Circle" of complex type CircleType
	Circle []*CircleType `xml:"Circle,omitempty"`

	// generated from element "ArcStringByBulge" of complex type ArcStringByBulgeType
	ArcStringByBulge []*ArcStringByBulgeType `xml:"ArcStringByBulge,omitempty"`

	// generated from element "ArcByBulge" of complex type ArcByBulgeType
	ArcByBulge []*ArcByBulgeType `xml:"ArcByBulge,omitempty"`

	// generated from element "ArcByCenterPoint" of complex type ArcByCenterPointType
	ArcByCenterPoint []*ArcByCenterPointType `xml:"ArcByCenterPoint,omitempty"`

	// generated from element "CircleByCenterPoint" of complex type CircleByCenterPointType
	CircleByCenterPoint []*CircleByCenterPointType `xml:"CircleByCenterPoint,omitempty"`

	// generated from element "CubicSpline" of complex type CubicSplineType
	CubicSpline []*CubicSplineType `xml:"CubicSpline,omitempty"`

	// generated from element "BSpline" of complex type BSplineType
	BSpline []*BSplineType `xml:"BSpline,omitempty"`

	// generated from element "Bezier" of complex type BezierType
	Bezier []*BezierType `xml:"Bezier,omitempty"`

	// generated from element "OffsetCurve" of complex type OffsetCurveType
	OffsetCurve []*OffsetCurveType `xml:"OffsetCurve,omitempty"`

	// generated from element "Clothoid" of complex type ClothoidType
	Clothoid []*ClothoidType `xml:"Clothoid,omitempty"`

	// generated from element "GeodesicString" of complex type GeodesicStringType
	GeodesicString []*GeodesicStringType `xml:"GeodesicString,omitempty"`

	// generated from element "Geodesic" of complex type GeodesicType
	Geodesic []*GeodesicType `xml:"Geodesic,omitempty"`
}

// LineStringSegmentType Named source named complex type "LineStringSegmentType"
type LineStringSegmentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// ArcStringType Named source named complex type "ArcStringType"
type ArcStringType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// ArcType Named source named complex type "ArcType"
type ArcType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep *PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// CircleType Named source named complex type "CircleType"
type CircleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// ArcStringByBulgeType Named source named complex type "ArcStringByBulgeType"
type ArcStringByBulgeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "bulge" of simple type double
	Bulge float64 `xml:"bulge,omitempty"`

	// generated from element "normal" of complex type VectorType
	Normal []*VectorType `xml:"normal,omitempty"`
}

// ArcByBulgeType Named source named complex type "ArcByBulgeType"
type ArcByBulgeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep *PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "bulge" of simple type double
	Bulge float64 `xml:"bulge,omitempty"`

	// generated from element "normal" of complex type VectorType
	Normal *VectorType `xml:"normal,omitempty"`
}

// ArcByCenterPointType Named source named complex type "ArcByCenterPointType"
type ArcByCenterPointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep *PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "radius" of complex type LengthType
	Radius *measures.LengthType `xml:"radius,omitempty"`

	// generated from element "startAngle" of complex type AngleType
	StartAngle *measures.AngleType `xml:"startAngle,omitempty"`

	// generated from element "endAngle" of complex type AngleType
	EndAngle *measures.AngleType `xml:"endAngle,omitempty"`
}

// CircleByCenterPointType Named source named complex type "CircleByCenterPointType"
type CircleByCenterPointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "numArc"
	NumArc int `xml:"numArc,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep *PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "radius" of complex type LengthType
	Radius *measures.LengthType `xml:"radius,omitempty"`
}

// CubicSplineType Named source named complex type "CubicSplineType"
type CubicSplineType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "degree"
	Degree int `xml:"degree,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "vectorAtStart" of complex type VectorType
	VectorAtStart *VectorType `xml:"vectorAtStart,omitempty"`

	// generated from element "vectorAtEnd" of complex type VectorType
	VectorAtEnd *VectorType `xml:"vectorAtEnd,omitempty"`
}

// BSplineType Named source named complex type "BSplineType"
type BSplineType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "isPolynomial"
	IsPolynomial bool `xml:"isPolynomial,attr,omitempty"`

	// generated from attribute "knotType"
	KnotType Enum_KnotTypesType `xml:"knotType,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "degree" of simple type nonNegativeInteger
	Degree int `xml:"degree,omitempty"`

	// generated from element "knot" of complex type KnotPropertyType
	Knot []*KnotPropertyType `xml:"knot,omitempty"`
}

// KnotType Named source named complex type "KnotType"
type KnotType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "value" of simple type double
	Value float64 `xml:"value,omitempty"`

	// generated from element "multiplicity" of simple type nonNegativeInteger
	Multiplicity int `xml:"multiplicity,omitempty"`

	// generated from element "weight" of simple type double
	Weight float64 `xml:"weight,omitempty"`
}

// KnotPropertyType Named source named complex type "KnotPropertyType"
// gml:KnotPropertyType encapsulates a knot to use it in a geometric type.
type KnotPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "Knot" of complex type KnotType
	Knot *KnotType `xml:"Knot,omitempty"`
}

// BezierType Named source named complex type "BezierType"
type BezierType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from attribute "isPolynomial"
	IsPolynomial bool `xml:"isPolynomial,attr,omitempty"`

	// generated from attribute "knotType"
	KnotType Enum_KnotTypesType `xml:"knotType,attr,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`

	// generated from element "degree" of simple type nonNegativeInteger
	Degree int `xml:"degree,omitempty"`

	// generated from element "knot" of complex type KnotPropertyType
	Knot *KnotPropertyType `xml:"knot,omitempty"`
}

// OffsetCurveType Named source named complex type "OffsetCurveType"
type OffsetCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from element "offsetBase" of complex type CurvePropertyType
	OffsetBase *CurvePropertyType `xml:"offsetBase,omitempty"`

	// generated from element "distance" of complex type LengthType
	Distance *measures.LengthType `xml:"distance,omitempty"`

	// generated from element "refDirection" of complex type VectorType
	RefDirection *VectorType `xml:"refDirection,omitempty"`
}

// AffinePlacementType Named source named complex type "AffinePlacementType"
type AffinePlacementType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "location" of complex type DirectPositionType
	Location *DirectPositionType `xml:"location,omitempty"`

	// generated from element "refDirection" of complex type VectorType
	RefDirection []*VectorType `xml:"refDirection,omitempty"`

	// generated from element "inDimension" of simple type positiveInteger
	InDimension int `xml:"inDimension,omitempty"`

	// generated from element "outDimension" of simple type positiveInteger
	OutDimension int `xml:"outDimension,omitempty"`
}

// ClothoidType Named source named complex type "ClothoidType"
type ClothoidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from anonymous complex type within element "refLocation"
	RefLocation *A_refLocation `xml:"refLocation,omitempty"`

	// generated from element "scaleFactor" of simple type decimal
	ScaleFactor string `xml:"scaleFactor,omitempty"`

	// generated from element "startParameter" of simple type double
	StartParameter float64 `xml:"startParameter,omitempty"`

	// generated from element "endParameter" of simple type double
	EndParameter float64 `xml:"endParameter,omitempty"`
}

// A_refLocation Named source within outer element "refLocation"
type A_refLocation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "AffinePlacement" of complex type AffinePlacementType
	AffinePlacement *AffinePlacementType `xml:"AffinePlacement,omitempty"`
}

// GeodesicStringType Named source named complex type "GeodesicStringType"
type GeodesicStringType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// GeodesicType Named source named complex type "GeodesicType"
type GeodesicType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "numDerivativesAtStart"
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`

	// generated from attribute "numDerivativesAtEnd"
	NumDerivativesAtEnd int `xml:"numDerivativesAtEnd,attr,omitempty"`

	// generated from attribute "numDerivativeInterior"
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`

	// generated from attribute "interpolation"
	Interpolation Enum_CurveInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// SurfaceType Named source named complex type "SurfaceType"
type SurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "patches" of complex type SurfacePatchArrayPropertyType
	Patches *SurfacePatchArrayPropertyType `xml:"patches,omitempty"`

	// generated from element "polygonPatches" of complex type SurfacePatchArrayPropertyType
	PolygonPatches *SurfacePatchArrayPropertyType `xml:"polygonPatches,omitempty"`

	// generated from element "trianglePatches" of complex type SurfacePatchArrayPropertyType
	TrianglePatches *SurfacePatchArrayPropertyType `xml:"trianglePatches,omitempty"`
}

// OrientableSurfaceType Named source named complex type "OrientableSurfaceType"
type OrientableSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from element "baseSurface" of complex type SurfacePropertyType
	BaseSurface *SurfacePropertyType `xml:"baseSurface,omitempty"`
}

// AbstractSurfacePatchType Named source named complex type "AbstractSurfacePatchType"
type AbstractSurfacePatchType struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// SurfacePatchArrayPropertyType Named source named complex type "SurfacePatchArrayPropertyType"
// gml:SurfacePatchArrayPropertyType is a container for a sequence of surface patches.
type SurfacePatchArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "PolygonPatch" of complex type PolygonPatchType
	PolygonPatch []*PolygonPatchType `xml:"PolygonPatch,omitempty"`

	// generated from element "Triangle" of complex type TriangleType
	Triangle []*TriangleType `xml:"Triangle,omitempty"`

	// generated from element "Rectangle" of complex type RectangleType
	Rectangle []*RectangleType `xml:"Rectangle,omitempty"`

	// generated from element "Cone" of complex type ConeType
	Cone []*ConeType `xml:"Cone,omitempty"`

	// generated from element "Cylinder" of complex type CylinderType
	Cylinder []*CylinderType `xml:"Cylinder,omitempty"`

	// generated from element "Sphere" of complex type SphereType
	Sphere []*SphereType `xml:"Sphere,omitempty"`
}

// PolygonPatchType Named source named complex type "PolygonPatchType"
type PolygonPatchType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "interpolation"
	Interpolation Enum_SurfaceInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "exterior" of complex type AbstractRingPropertyType
	Exterior *AbstractRingPropertyType `xml:"exterior,omitempty"`

	// generated from element "interior" of complex type AbstractRingPropertyType
	Interior []*AbstractRingPropertyType `xml:"interior,omitempty"`
}

// TriangleType Named source named complex type "TriangleType"
type TriangleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "interpolation"
	Interpolation Enum_SurfaceInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "exterior" of complex type AbstractRingPropertyType
	Exterior *AbstractRingPropertyType `xml:"exterior,omitempty"`
}

// RectangleType Named source named complex type "RectangleType"
type RectangleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "interpolation"
	Interpolation Enum_SurfaceInterpolationType `xml:"interpolation,attr,omitempty"`

	// generated from element "exterior" of complex type AbstractRingPropertyType
	Exterior *AbstractRingPropertyType `xml:"exterior,omitempty"`
}

// RingType Named source named complex type "RingType"
type RingType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "curveMember" of complex type CurvePropertyType
	CurveMember []*CurvePropertyType `xml:"curveMember,omitempty"`
}

// RingPropertyType Named source named complex type "RingPropertyType"
// A property with the content model of gml:RingPropertyType encapsulates a ring to represent a component of a surface boundary.
type RingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`
}

// Group_PointGrid UnNamed source named group "PointGrid"
type Group_PointGrid struct {

	// insertion point for fields

	// generated from anonymous complex type within element "rows"
	Rows *A_rows `xml:"rows,omitempty"`
}

// A_rows Named source within outer element "rows"
type A_rows struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from anonymous complex type within element "Row"
	Row []*A_Row `xml:"Row,omitempty"`
}

// A_Row Named source within outer element "Row"
type A_Row struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group "gml:geometricPositionListGroup" with order 480 depth 7
	Group_geometricPositionListGroup
}

// AbstractParametricCurveSurfaceType Named source named complex type "AbstractParametricCurveSurfaceType"
type AbstractParametricCurveSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup
}

// AbstractGriddedSurfaceType Named source named complex type "AbstractGriddedSurfaceType"
type AbstractGriddedSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from attribute "rows"
	Rows int `xml:"rows,attr,omitempty"`

	// generated from attribute "columns"
	Columns int `xml:"columns,attr,omitempty"`

	// generated from anonymous complex type within element "rows"
	Rows_1 *A_rows `xml:"rows,omitempty"`
}

// ConeType Named source named complex type "ConeType"
type ConeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from attribute "rows"
	Rows int `xml:"rows,attr,omitempty"`

	// generated from attribute "columns"
	Columns int `xml:"columns,attr,omitempty"`

	// generated from anonymous complex type within element "rows"
	Rows_1 *A_rows `xml:"rows,omitempty"`

	// generated from attribute "horizontalCurveType"
	HorizontalCurveType Enum_CurveInterpolationType `xml:"horizontalCurveType,attr,omitempty"`

	// generated from attribute "verticalCurveType"
	VerticalCurveType Enum_CurveInterpolationType `xml:"verticalCurveType,attr,omitempty"`
}

// CylinderType Named source named complex type "CylinderType"
type CylinderType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from attribute "rows"
	Rows int `xml:"rows,attr,omitempty"`

	// generated from attribute "columns"
	Columns int `xml:"columns,attr,omitempty"`

	// generated from anonymous complex type within element "rows"
	Rows_1 *A_rows `xml:"rows,omitempty"`

	// generated from attribute "horizontalCurveType"
	HorizontalCurveType Enum_CurveInterpolationType `xml:"horizontalCurveType,attr,omitempty"`

	// generated from attribute "verticalCurveType"
	VerticalCurveType Enum_CurveInterpolationType `xml:"verticalCurveType,attr,omitempty"`
}

// SphereType Named source named complex type "SphereType"
type SphereType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from attribute "rows"
	Rows int `xml:"rows,attr,omitempty"`

	// generated from attribute "columns"
	Columns int `xml:"columns,attr,omitempty"`

	// generated from anonymous complex type within element "rows"
	Rows_1 *A_rows `xml:"rows,omitempty"`

	// generated from attribute "horizontalCurveType"
	HorizontalCurveType Enum_CurveInterpolationType `xml:"horizontalCurveType,attr,omitempty"`

	// generated from attribute "verticalCurveType"
	VerticalCurveType Enum_CurveInterpolationType `xml:"verticalCurveType,attr,omitempty"`
}

// TinType Named source named complex type "TinType"
type TinType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "patches" of complex type SurfacePatchArrayPropertyType
	Patches *SurfacePatchArrayPropertyType `xml:"patches,omitempty"`

	// generated from element "polygonPatches" of complex type SurfacePatchArrayPropertyType
	PolygonPatches *SurfacePatchArrayPropertyType `xml:"polygonPatches,omitempty"`

	// generated from element "trianglePatches" of complex type SurfacePatchArrayPropertyType
	TrianglePatches *SurfacePatchArrayPropertyType `xml:"trianglePatches,omitempty"`

	// generated from element "stopLines" of complex type LineStringSegmentArrayPropertyType
	StopLines []*LineStringSegmentArrayPropertyType `xml:"stopLines,omitempty"`

	// generated from element "breakLines" of complex type LineStringSegmentArrayPropertyType
	BreakLines []*LineStringSegmentArrayPropertyType `xml:"breakLines,omitempty"`

	// generated from element "maxLength" of complex type LengthType
	MaxLength *measures.LengthType `xml:"maxLength,omitempty"`

	// generated from anonymous complex type within element "controlPoint"
	ControlPoint *A_controlPoint `xml:"controlPoint,omitempty"`
}

// A_controlPoint Named source within outer element "controlPoint"
type A_controlPoint struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// LineStringSegmentArrayPropertyType Named source named complex type "LineStringSegmentArrayPropertyType"
// gml:LineStringSegmentArrayPropertyType provides a container for line strings.
type LineStringSegmentArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "LineStringSegment" of complex type LineStringSegmentType
	LineStringSegment []*LineStringSegmentType `xml:"LineStringSegment,omitempty"`
}

// AbstractSolidType Named source named complex type "AbstractSolidType"
// gml:AbstractSolidType is an abstraction of a solid to support the different levels of complexity. The solid may always be viewed as a geometric primitive, i.e. is contiguous.
type AbstractSolidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// SolidPropertyType Named source named complex type "SolidPropertyType"
// A property that has a solid as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type SolidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`
}

// SolidArrayPropertyType Named source named complex type "SolidArrayPropertyType"
// gml:SolidArrayPropertyType is a container for an array of solids. The elements are always contained in the array property, referencing geometry elements or arrays of geometry elements is not supported.
type SolidArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Solid" of complex type SolidType
	Solid []*SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid []*CompositeSolidType `xml:"CompositeSolid,omitempty"`
}

// SolidType Named source named complex type "SolidType"
type SolidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "exterior" of complex type ShellPropertyType
	Exterior *ShellPropertyType `xml:"exterior,omitempty"`

	// generated from element "interior" of complex type ShellPropertyType
	Interior []*ShellPropertyType `xml:"interior,omitempty"`
}

// ShellType Named source named complex type "ShellType"
type ShellType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "surfaceMember" of complex type SurfacePropertyType
	SurfaceMember []*SurfacePropertyType `xml:"surfaceMember,omitempty"`
}

// ShellPropertyType Named source named complex type "ShellPropertyType"
// A property with the content model of gml:ShellPropertyType encapsulates a shell to represent a component of a solid boundary.
type ShellPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`
}

// AbstractSurfaceType Named source named complex type "AbstractSurfaceType"
// gml:AbstractSurfaceType is an abstraction of a surface to support the different levels of complexity. A surface is always a continuous region of a plane.
type AbstractSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// SurfacePropertyType Named source named complex type "SurfacePropertyType"
// A property that has a surface as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type SurfacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`
}

// SurfaceArrayPropertyType Named source named complex type "SurfaceArrayPropertyType"
// gml:SurfaceArrayPropertyType is a container for an array of surfaces. The elements are always contained in the array property, referencing geometry elements or arrays of geometry elements via XLinks is not supported.
type SurfaceArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Surface" of complex type SurfaceType
	Surface []*SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface []*SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface []*SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin []*TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface []*OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell []*ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon []*PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface []*CompositeSurfaceType `xml:"CompositeSurface,omitempty"`
}

// PolygonType Named source named complex type "PolygonType"
type PolygonType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "exterior" of complex type AbstractRingPropertyType
	Exterior *AbstractRingPropertyType `xml:"exterior,omitempty"`

	// generated from element "interior" of complex type AbstractRingPropertyType
	Interior []*AbstractRingPropertyType `xml:"interior,omitempty"`
}

// AbstractRingType Named source named complex type "AbstractRingType"
type AbstractRingType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// AbstractRingPropertyType Named source named complex type "AbstractRingPropertyType"
// A property with the content model of gml:AbstractRingPropertyType encapsulates a ring to represent the surface boundary property of a surface.
type AbstractRingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`
}

// LinearRingType Named source named complex type "LinearRingType"
type LinearRingType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// LinearRingPropertyType Named source named complex type "LinearRingPropertyType"
// A property with the content model of gml:LinearRingPropertyType encapsulates a linear ring to represent a component of a surface boundary.
type LinearRingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`
}

// AbstractGeometryType Named source named complex type "AbstractGeometryType"
// All geometry elements are derived directly or indirectly from this abstract supertype. A geometry element may have an identifying attribute (gml:id), one or more names (elements identifier and name) and a description (elements description and descriptionReference) . It may be associated with a spatial reference system (attribute group gml:SRSReferenceGroup).
// The following rules shall be adhered to:
// -	Every geometry type shall derive from this abstract type.
// -	Every geometry element (i.e. an element of a geometry type) shall be directly or indirectly in the substitution group of AbstractGeometry.
type AbstractGeometryType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// AttributeGroup_SRSReferenceGroup UnNamed source named attribute group "SRSReferenceGroup"
type AttributeGroup_SRSReferenceGroup struct {

	// insertion point for fields

	// generated from attribute "srsName"
	SrsName string `xml:"srsName,attr,omitempty"`

	// generated from attribute "srsDimension"
	SrsDimension int `xml:"srsDimension,attr,omitempty"`

	// generated from attribute group "gml:SRSInformationGroup"
	AttributeGroup_SRSInformationGroup
}

// AttributeGroup_SRSInformationGroup UnNamed source named attribute group "SRSInformationGroup"
type AttributeGroup_SRSInformationGroup struct {

	// insertion point for fields

	// generated from attribute "axisLabels"
	AxisLabels string `xml:"axisLabels,attr,omitempty"`

	// generated from attribute "uomLabels"
	UomLabels string `xml:"uomLabels,attr,omitempty"`
}

// GeometryPropertyType Named source named complex type "GeometryPropertyType"
// A geometric property may either be any geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same or another document). Note that either the reference or the contained element shall be given, but not both or none.
// If a feature has a property that takes a geometry element as its value, this is called a geometry property. A generic type for such a geometry property is GeometryPropertyType.
type GeometryPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`
}

// GeometryArrayPropertyType Named source named complex type "GeometryArrayPropertyType"
// If a feature has a property which takes an array of geometry elements as its value, this is called a geometry array property. A generic type for such a geometry property is GeometryArrayPropertyType. 
// The elements are always contained inline in the array property, referencing geometry elements or arrays of geometry elements via XLinks is not supported.
type GeometryArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry []*MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint []*MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve []*MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface []*MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid []*MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid []*SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid []*CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface []*SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface []*SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface []*SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin []*TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface []*OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell []*ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon []*PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface []*CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point []*PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve []*CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve []*OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring []*RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing []*LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString []*LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve []*CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex []*GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid []*GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid []*RectifiedGridType `xml:"RectifiedGrid,omitempty"`
}

// DirectPositionType Named source named complex type "DirectPositionType"
// Direct position instances hold the coordinates for a position within some coordinate reference system (CRS). Since direct positions, as data types, will often be included in larger objects (such as geometry elements) that have references to CRS, the srsName attribute will in general be missing, if this particular direct position is included in a larger element with such a reference to a CRS. In this case, the CRS is implicitly assumed to take on the value of the containing object's CRS.
// if no srsName attribute is given, the CRS shall be specified as part of the larger context this geometry element is part of, typically a geometric object like a point, curve, etc.
type DirectPositionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// DirectPositionListType Named source named complex type "DirectPositionListType"
// posList instances (and other instances with the content model specified by DirectPositionListType) hold the coordinates for a sequence of direct positions within the same coordinate reference system (CRS).
// if no srsName attribute is given, the CRS shall be specified as part of the larger context this geometry element is part of, typically a geometric object like a point, curve, etc. 
// The optional attribute count specifies the number of direct positions in the list. If the attribute count is present then the attribute srsDimension shall be present, too.
// The number of entries in the list is equal to the product of the dimensionality of the coordinate reference system (i.e. it is a derived value of the coordinate reference system definition) and the number of direct positions.
type DirectPositionListType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "count"
	Count int `xml:"count,attr,omitempty"`

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Group_geometricPositionGroup UnNamed source named group "geometricPositionGroup"
type Group_geometricPositionGroup struct {

	// insertion point for fields

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// Group_geometricPositionListGroup UnNamed source named group "geometricPositionListGroup"
type Group_geometricPositionListGroup struct {

	// insertion point for fields

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// VectorType Named source named complex type "VectorType"
// For some applications the components of the position may be adjusted to yield a unit vector.
type VectorType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// EnvelopeType Named source named complex type "EnvelopeType"
type EnvelopeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "lowerCorner" of complex type DirectPositionType
	LowerCorner *DirectPositionType `xml:"lowerCorner,omitempty"`

	// generated from element "upperCorner" of complex type DirectPositionType
	UpperCorner *DirectPositionType `xml:"upperCorner,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// AbstractGeometricPrimitiveType Named source named complex type "AbstractGeometricPrimitiveType"
// gml:AbstractGeometricPrimitiveType is the abstract root type of the geometric primitives. A geometric primitive is a geometric object that is not decomposed further into other primitives in the system. All primitives are oriented in the direction implied by the sequence of their coordinate tuples.
type AbstractGeometricPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// GeometricPrimitivePropertyType Named source named complex type "GeometricPrimitivePropertyType"
// A property that has a geometric primitive as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type GeometricPrimitivePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`
}

// PointType Named source named complex type "PointType"
type PointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// PointPropertyType Named source named complex type "PointPropertyType"
// A property that has a point as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type PointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`
}

// PointArrayPropertyType Named source named complex type "PointArrayPropertyType"
// gml:PointArrayPropertyType is a container for an array of points. The elements are always contained inline in the array property, referencing geometry elements or arrays of geometry elements via XLinks is not supported.
type PointArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Point" of complex type PointType
	Point []*PointType `xml:"Point,omitempty"`
}

// AbstractCurveType Named source named complex type "AbstractCurveType"
// gml:AbstractCurveType is an abstraction of a curve to support the different levels of complexity. The curve may always be viewed as a geometric primitive, i.e. is continuous.
type AbstractCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// CurvePropertyType Named source named complex type "CurvePropertyType"
// A property that has a curve as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type CurvePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`
}

// CurveArrayPropertyType Named source named complex type "CurveArrayPropertyType"
// A container for an array of curves. The elements are always contained in the array property, referencing geometry elements or arrays of geometry elements via XLinks is not supported.
type CurveArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Curve" of complex type CurveType
	Curve []*CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve []*OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring []*RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing []*LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString []*LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve []*CompositeCurveType `xml:"CompositeCurve,omitempty"`
}

// LineStringType Named source named complex type "LineStringType"
type LineStringType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos []*DirectPositionType `xml:"pos,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty []*PointPropertyType `xml:"pointProperty,omitempty"`

	// generated from element "pointRep" of complex type PointPropertyType
	PointRep []*PointPropertyType `xml:"pointRep,omitempty"`

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "coordinates" of complex type CoordinatesType
	Coordinates *basictypes.CoordinatesType `xml:"coordinates,omitempty"`
}

// UnitOfMeasureType Named source named complex type "UnitOfMeasureType"
type UnitOfMeasureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`
}

// UnitDefinitionType Named source named complex type "UnitDefinitionType"
type UnitDefinitionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "quantityType" of complex type StringOrRefType
	QuantityType *StringOrRefType `xml:"quantityType,omitempty"`

	// generated from element "quantityTypeReference" of complex type ReferenceType
	QuantityTypeReference *ReferenceType `xml:"quantityTypeReference,omitempty"`

	// generated from element "catalogSymbol" of complex type CodeType
	CatalogSymbol *basictypes.CodeType `xml:"catalogSymbol,omitempty"`
}

// BaseUnitType Named source named complex type "BaseUnitType"
type BaseUnitType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "quantityType" of complex type StringOrRefType
	QuantityType *StringOrRefType `xml:"quantityType,omitempty"`

	// generated from element "quantityTypeReference" of complex type ReferenceType
	QuantityTypeReference *ReferenceType `xml:"quantityTypeReference,omitempty"`

	// generated from element "catalogSymbol" of complex type CodeType
	CatalogSymbol *basictypes.CodeType `xml:"catalogSymbol,omitempty"`

	// generated from element "unitsSystem" of complex type ReferenceType
	UnitsSystem *ReferenceType `xml:"unitsSystem,omitempty"`
}

// DerivedUnitType Named source named complex type "DerivedUnitType"
type DerivedUnitType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "quantityType" of complex type StringOrRefType
	QuantityType *StringOrRefType `xml:"quantityType,omitempty"`

	// generated from element "quantityTypeReference" of complex type ReferenceType
	QuantityTypeReference *ReferenceType `xml:"quantityTypeReference,omitempty"`

	// generated from element "catalogSymbol" of complex type CodeType
	CatalogSymbol *basictypes.CodeType `xml:"catalogSymbol,omitempty"`

	// generated from element "derivationUnitTerm" of complex type DerivationUnitTermType
	DerivationUnitTerm []*DerivationUnitTermType `xml:"derivationUnitTerm,omitempty"`
}

// DerivationUnitTermType Named source named complex type "DerivationUnitTermType"
type DerivationUnitTermType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// generated from attribute "exponent"
	Exponent int `xml:"exponent,attr,omitempty"`
}

// ConventionalUnitType Named source named complex type "ConventionalUnitType"
type ConventionalUnitType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "quantityType" of complex type StringOrRefType
	QuantityType *StringOrRefType `xml:"quantityType,omitempty"`

	// generated from element "quantityTypeReference" of complex type ReferenceType
	QuantityTypeReference *ReferenceType `xml:"quantityTypeReference,omitempty"`

	// generated from element "catalogSymbol" of complex type CodeType
	CatalogSymbol *basictypes.CodeType `xml:"catalogSymbol,omitempty"`

	// generated from element "conversionToPreferredUnit" of complex type ConversionToPreferredUnitType
	ConversionToPreferredUnit *ConversionToPreferredUnitType `xml:"conversionToPreferredUnit,omitempty"`

	// generated from element "roughConversionToPreferredUnit" of complex type ConversionToPreferredUnitType
	RoughConversionToPreferredUnit *ConversionToPreferredUnitType `xml:"roughConversionToPreferredUnit,omitempty"`

	// generated from element "derivationUnitTerm" of complex type DerivationUnitTermType
	DerivationUnitTerm []*DerivationUnitTermType `xml:"derivationUnitTerm,omitempty"`
}

// ConversionToPreferredUnitType Named source named complex type "ConversionToPreferredUnitType"
// The inherited attribute uom references the preferred unit that this conversion applies to. The conversion of a unit to the preferred unit for this physical quantity type is specified by an arithmetic conversion (scaling and/or offset). The content model extends gml:UnitOfMeasureType, which has a mandatory attribute uom which identifies the preferred unit for the physical quantity type that this conversion applies to. The conversion is specified by a choice of 
// -	gml:factor, which defines the scale factor, or
// -	gml:formula, which defines a formula 
// by which a value using the conventional unit of measure can be converted to obtain the corresponding value using the preferred unit of measure.  
// The formula defines the parameters of a simple formula by which a value using the conventional unit of measure can be converted to the corresponding value using the preferred unit of measure. The formula element contains elements a, b, c and d, whose values use the XML Schema type double. These values are used in the formula y = (a + bx) / (c + dx), where x is a value using this unit, and y is the corresponding value using the base unit. The elements a and d are optional, and if values are not provided, those parameters are considered to be zero. If values are not provided for both a and d, the formula is equivalent to a fraction with numerator and denominator parameters.
type ConversionToPreferredUnitType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// generated from element "factor" of simple type double
	Factor float64 `xml:"factor,omitempty"`

	// generated from element "formula" of complex type FormulaType
	Formula *FormulaType `xml:"formula,omitempty"`
}

// FormulaType Named source named complex type "FormulaType"
type FormulaType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "a" of simple type double
	A float64 `xml:"a,omitempty"`

	// generated from element "b" of simple type double
	B float64 `xml:"b,omitempty"`

	// generated from element "c" of simple type double
	C float64 `xml:"c,omitempty"`

	// generated from element "d" of simple type double
	D float64 `xml:"d,omitempty"`
}

// DefinitionBaseType Named source named complex type "DefinitionBaseType"
type DefinitionBaseType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`
}

// DefinitionType Named source named complex type "DefinitionType"
type DefinitionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`
}

// DictionaryType Named source named complex type "DictionaryType"
type DictionaryType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "dictionaryEntry" of complex type DictionaryEntryType
	DictionaryEntry []*DictionaryEntryType `xml:"dictionaryEntry,omitempty"`

	// generated from element "definitionMember" of complex type DictionaryEntryType
	DefinitionMember []*DictionaryEntryType `xml:"definitionMember,omitempty"`

	// generated from element "indirectEntry" of complex type IndirectEntryType
	IndirectEntry []*IndirectEntryType `xml:"indirectEntry,omitempty"`
}

// DictionaryEntryType Named source named complex type "DictionaryEntryType"
type DictionaryEntryType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Definition" of complex type DefinitionType
	Definition *DefinitionType `xml:"Definition,omitempty"`

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`

	// generated from element "Dictionary" of complex type DictionaryType
	Dictionary *DictionaryType `xml:"Dictionary,omitempty"`

	// generated from element "CoordinateSystemAxis" of complex type CoordinateSystemAxisType
	CoordinateSystemAxis *CoordinateSystemAxisType `xml:"CoordinateSystemAxis,omitempty"`

	// generated from element "EllipsoidalCS" of complex type EllipsoidalCSType
	EllipsoidalCS *EllipsoidalCSType `xml:"EllipsoidalCS,omitempty"`

	// generated from element "CartesianCS" of complex type CartesianCSType
	CartesianCS *CartesianCSType `xml:"CartesianCS,omitempty"`

	// generated from element "VerticalCS" of complex type VerticalCSType
	VerticalCS *VerticalCSType `xml:"VerticalCS,omitempty"`

	// generated from element "TimeCS" of complex type TimeCSType
	TimeCS *TimeCSType `xml:"TimeCS,omitempty"`

	// generated from element "LinearCS" of complex type LinearCSType
	LinearCS *LinearCSType `xml:"LinearCS,omitempty"`

	// generated from element "UserDefinedCS" of complex type UserDefinedCSType
	UserDefinedCS *UserDefinedCSType `xml:"UserDefinedCS,omitempty"`

	// generated from element "SphericalCS" of complex type SphericalCSType
	SphericalCS *SphericalCSType `xml:"SphericalCS,omitempty"`

	// generated from element "PolarCS" of complex type PolarCSType
	PolarCS *PolarCSType `xml:"PolarCS,omitempty"`

	// generated from element "CylindricalCS" of complex type CylindricalCSType
	CylindricalCS *CylindricalCSType `xml:"CylindricalCS,omitempty"`

	// generated from element "AffineCS" of complex type AffineCSType
	AffineCS *AffineCSType `xml:"AffineCS,omitempty"`

	// generated from element "TemporalCS" of complex type TemporalCSType
	TemporalCS *TemporalCSType `xml:"TemporalCS,omitempty"`

	// generated from element "ObliqueCartesianCS" of complex type ObliqueCartesianCSType
	ObliqueCartesianCS *ObliqueCartesianCSType `xml:"ObliqueCartesianCS,omitempty"`

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS *ProjectedCRSType `xml:"ProjectedCRS,omitempty"`

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS *DerivedCRSType `xml:"DerivedCRS,omitempty"`

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS *GeodeticCRSType `xml:"GeodeticCRS,omitempty"`

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS *VerticalCRSType `xml:"VerticalCRS,omitempty"`

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS *EngineeringCRSType `xml:"EngineeringCRS,omitempty"`

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS *ImageCRSType `xml:"ImageCRS,omitempty"`

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS *TemporalCRSType `xml:"TemporalCRS,omitempty"`

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS *GeographicCRSType `xml:"GeographicCRS,omitempty"`

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS *GeocentricCRSType `xml:"GeocentricCRS,omitempty"`

	// generated from element "CompoundCRS" of complex type CompoundCRSType
	CompoundCRS *CompoundCRSType `xml:"CompoundCRS,omitempty"`

	// generated from element "GeodeticDatum" of complex type GeodeticDatumType
	GeodeticDatum *GeodeticDatumType `xml:"GeodeticDatum,omitempty"`

	// generated from element "EngineeringDatum" of complex type EngineeringDatumType
	EngineeringDatum *EngineeringDatumType `xml:"EngineeringDatum,omitempty"`

	// generated from element "ImageDatum" of complex type ImageDatumType
	ImageDatum *ImageDatumType `xml:"ImageDatum,omitempty"`

	// generated from element "VerticalDatum" of complex type VerticalDatumType
	VerticalDatum *VerticalDatumType `xml:"VerticalDatum,omitempty"`

	// generated from element "TemporalDatum" of complex type TemporalDatumType
	TemporalDatum *TemporalDatumType `xml:"TemporalDatum,omitempty"`

	// generated from element "Ellipsoid" of complex type EllipsoidType
	Ellipsoid *EllipsoidType `xml:"Ellipsoid,omitempty"`

	// generated from element "PrimeMeridian" of complex type PrimeMeridianType
	PrimeMeridian *PrimeMeridianType `xml:"PrimeMeridian,omitempty"`

	// generated from element "PassThroughOperation" of complex type PassThroughOperationType
	PassThroughOperation *PassThroughOperationType `xml:"PassThroughOperation,omitempty"`

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`

	// generated from element "ConcatenatedOperation" of complex type ConcatenatedOperationType
	ConcatenatedOperation *ConcatenatedOperationType `xml:"ConcatenatedOperation,omitempty"`

	// generated from element "OperationMethod" of complex type OperationMethodType
	OperationMethod *OperationMethodType `xml:"OperationMethod,omitempty"`

	// generated from element "OperationParameter" of complex type OperationParameterType
	OperationParameter *OperationParameterType `xml:"OperationParameter,omitempty"`

	// generated from element "OperationParameterGroup" of complex type OperationParameterGroupType
	OperationParameterGroup *OperationParameterGroupType `xml:"OperationParameterGroup,omitempty"`

	// generated from element "TimeReferenceSystem" of complex type TimeReferenceSystemType
	TimeReferenceSystem *TimeReferenceSystemType `xml:"TimeReferenceSystem,omitempty"`

	// generated from element "TimeCoordinateSystem" of complex type TimeCoordinateSystemType
	TimeCoordinateSystem *TimeCoordinateSystemType `xml:"TimeCoordinateSystem,omitempty"`

	// generated from element "TimeCalendar" of complex type TimeCalendarType
	TimeCalendar *TimeCalendarType `xml:"TimeCalendar,omitempty"`

	// generated from element "TimeClock" of complex type TimeClockType
	TimeClock *TimeClockType `xml:"TimeClock,omitempty"`

	// generated from element "TimeOrdinalReferenceSystem" of complex type TimeOrdinalReferenceSystemType
	TimeOrdinalReferenceSystem *TimeOrdinalReferenceSystemType `xml:"TimeOrdinalReferenceSystem,omitempty"`

	// generated from element "DefinitionCollection" of complex type DictionaryType
	DefinitionCollection *DictionaryType `xml:"DefinitionCollection,omitempty"`

	// generated from element "DefinitionProxy" of complex type DefinitionProxyType
	DefinitionProxy *DefinitionProxyType `xml:"DefinitionProxy,omitempty"`
}

// AbstractGMLType Named source named complex type "AbstractGMLType"
type AbstractGMLType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties
}

// Group_StandardObjectProperties UnNamed source named group "StandardObjectProperties"
type Group_StandardObjectProperties struct {

	// insertion point for fields

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`
}

// AttributeGroup_AssociationAttributeGroup UnNamed source named attribute group "AssociationAttributeGroup"
type AttributeGroup_AssociationAttributeGroup struct {

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from attribute "gml:remoteSchema"
	Attr_remoteSchema

	// generated from attribute group "xlink:simpleAttrs"
	xlink.AttributeGroup_simpleAttrs
}

// AssociationRoleType Named source named complex type "AssociationRoleType"
type AssociationRoleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup
}

// AttributeGroup_OwnershipAttributeGroup UnNamed source named attribute group "OwnershipAttributeGroup"
type AttributeGroup_OwnershipAttributeGroup struct {

	// insertion point for fields

	// generated from attribute "owns"
	Owns bool `xml:"owns,attr,omitempty"`
}

// ReferenceType Named source named complex type "ReferenceType"
// gml:ReferenceType is intended to be used in application schemas directly, if a property element shall use a "by-reference only" encoding.
type ReferenceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup
}

// InlinePropertyType Named source named complex type "InlinePropertyType"
type InlinePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup
}

// Attr_id UnNamed source top-level attribute "id"
type Attr_id struct {

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`
}

// AbstractMemberType Named source named complex type "AbstractMemberType"
// To create a collection of GML Objects that are not all features, a property type shall be derived by extension from gml:AbstractMemberType.
// This abstract property type is intended to be used only in object types where software shall be able to identify that an instance of such an object type is to be interpreted as a collection of objects.
// By default, this abstract property type does not imply any ownership of the objects in the collection. The owns attribute of gml:OwnershipAttributeGroup may be used on a property element instance to assert ownership of an object in the collection. A collection shall not own an object already owned by another object.
// 
type AbstractMemberType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup
}

// AttributeGroup_AggregationAttributeGroup UnNamed source named attribute group "AggregationAttributeGroup"
type AttributeGroup_AggregationAttributeGroup struct {

	// insertion point for fields

	// generated from attribute "aggregationType"
	AggregationType string `xml:"aggregationType,attr,omitempty"`
}

// AbstractMetadataPropertyType Named source named complex type "AbstractMetadataPropertyType"
// To associate metadata described by any XML Schema with a GML object, a property element shall be defined whose content model is derived by extension from gml:AbstractMetadataPropertyType.
// The value of such a property shall be metadata. The content model of such a property type, i.e. the metadata application schema shall be specified by the GML Application Schema.
// By default, this abstract property type does not imply any ownership of the metadata. The owns attribute of gml:OwnershipAttributeGroup may be used on a metadata property element instance to assert ownership of the metadata.
// If metadata following the conceptual model of ISO 19115 is to be encoded in a GML document, the corresponding Implementation Specification specified in ISO/TS 19139 shall be used to encode the metadata information.
// 
type AbstractMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup
}

// AbstractTimeObjectType Named source named complex type "AbstractTimeObjectType"
type AbstractTimeObjectType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties
}

// AbstractTimePrimitiveType Named source named complex type "AbstractTimePrimitiveType"
type AbstractTimePrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`
}

// TimePrimitivePropertyType Named source named complex type "TimePrimitivePropertyType"
// gml:TimePrimitivePropertyType provides a standard content model for associations between an arbitrary member of the substitution group whose head is gml:AbstractTimePrimitive and another object.
type TimePrimitivePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`
}

// RelatedTimeType Named source named complex type "RelatedTimeType"
// gml:RelatedTimeType provides a content model for indicating the relative position of an arbitrary member of the substitution group whose head is gml:AbstractTimePrimitive. It extends the generic gml:TimePrimitivePropertyType with an XML attribute relativePosition, whose value is selected from the set of 13 temporal relationships identified by Allen (1983)
type RelatedTimeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`

	// generated from attribute "relativePosition"
	RelativePosition Enum_RelatedTimeType_relativePosition `xml:"relativePosition,attr,omitempty"`
}

// AbstractTimeComplexType Named source named complex type "AbstractTimeComplexType"
type AbstractTimeComplexType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties
}

// AbstractTimeGeometricPrimitiveType Named source named complex type "AbstractTimeGeometricPrimitiveType"
type AbstractTimeGeometricPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`
}

// TimeInstantType Named source named complex type "TimeInstantType"
type TimeInstantType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from element "timePosition" of complex type TimePositionType
	TimePosition *TimePositionType `xml:"timePosition,omitempty"`
}

// TimeInstantPropertyType Named source named complex type "TimeInstantPropertyType"
// gml:TimeInstantPropertyType provides for associating a gml:TimeInstant with an object.
type TimeInstantPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`
}

// TimePeriodType Named source named complex type "TimePeriodType"
type TimePeriodType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from element "beginPosition" of complex type TimePositionType
	BeginPosition *TimePositionType `xml:"beginPosition,omitempty"`

	// generated from element "begin" of complex type TimeInstantPropertyType
	Begin *TimeInstantPropertyType `xml:"begin,omitempty"`

	// generated from element "endPosition" of complex type TimePositionType
	EndPosition *TimePositionType `xml:"endPosition,omitempty"`

	// generated from element "end" of complex type TimeInstantPropertyType
	End *TimeInstantPropertyType `xml:"end,omitempty"`

	// generated from group "gml:timeLength" with order 881 depth 2
	Group_timeLength
}

// TimePeriodPropertyType Named source named complex type "TimePeriodPropertyType"
// gml:TimePeriodPropertyType provides for associating a gml:TimePeriod with an object.
type TimePeriodPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`
}

// TimePositionType Named source named complex type "TimePositionType"
// The method for identifying a temporal position is specific to each temporal reference system.  gml:TimePositionType supports the description of temporal position according to the subtypes described in ISO 19108.
// Values based on calendars and clocks use lexical formats that are based on ISO 8601, as described in XML Schema Part 2:2001. A decimal value may be used with coordinate systems such as GPS time or UNIX time. A URI may be used to provide a reference to some era in an ordinal reference system . 
// In common with many of the components modelled as data types in the ISO 19100 series of International Standards, the corresponding GML component has simple content. However, the content model gml:TimePositionType is defined in several steps.
// Three XML attributes appear on gml:TimePositionType:
// A time value shall be associated with a temporal reference system through the frame attribute that provides a URI reference that identifies a description of the reference system. Following ISO 19108, the Gregorian calendar with UTC is the default reference system, but others may also be used. Components for describing temporal reference systems are described in 14.4, but it is not required that the reference system be described in this, as the reference may refer to anything that may be indentified with a URI.  
// For time values using a calendar containing more than one era, the (optional) calendarEraName attribute provides the name of the calendar era.  
// Inexact temporal positions may be expressed using the optional indeterminatePosition attribute.  This takes a value from an enumeration.
type TimePositionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from attribute "calendarEraName"
	CalendarEraName string `xml:"calendarEraName,attr,omitempty"`

	// generated from attribute "indeterminatePosition"
	IndeterminatePosition Enum_TimeIndeterminateValueType `xml:"indeterminatePosition,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Group_timeLength UnNamed source named group "timeLength"
type Group_timeLength struct {

	// insertion point for fields

	// generated from element "duration" of simple type duration
	Duration string `xml:"duration,omitempty"`

	// generated from element "timeInterval" of complex type TimeIntervalLengthType
	TimeInterval *TimeIntervalLengthType `xml:"timeInterval,omitempty"`
}

// TimeIntervalLengthType Named source named complex type "TimeIntervalLengthType"
type TimeIntervalLengthType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "unit"
	Unit string `xml:"unit,attr,omitempty"`

	// generated from attribute "radix"
	Radix int `xml:"radix,attr,omitempty"`

	// generated from attribute "factor"
	Factor int `xml:"factor,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// DirectionPropertyType Named source named complex type "DirectionPropertyType"
type DirectionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DirectionVector" of complex type DirectionVectorType
	DirectionVector *DirectionVectorType `xml:"DirectionVector,omitempty"`

	// generated from element "DirectionDescription" of complex type DirectionDescriptionType
	DirectionDescription *DirectionDescriptionType `xml:"DirectionDescription,omitempty"`

	// generated from element "CompassPoint" of simple type gml:CompassPointEnumeration
	CompassPoint Enum_CompassPointEnumeration `xml:"CompassPoint,omitempty"`

	// generated from element "DirectionKeyword" of complex type CodeType
	DirectionKeyword *basictypes.CodeType `xml:"DirectionKeyword,omitempty"`

	// generated from element "DirectionString" of complex type StringOrRefType
	DirectionString *StringOrRefType `xml:"DirectionString,omitempty"`
}

// DirectionVectorType Named source named complex type "DirectionVectorType"
// Direction vectors are specified by providing components of a vector.
type DirectionVectorType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "vector" of complex type VectorType
	Vector *VectorType `xml:"vector,omitempty"`

	// generated from element "horizontalAngle" of complex type AngleType
	HorizontalAngle *measures.AngleType `xml:"horizontalAngle,omitempty"`

	// generated from element "verticalAngle" of complex type AngleType
	VerticalAngle *measures.AngleType `xml:"verticalAngle,omitempty"`
}

// DirectionDescriptionType Named source named complex type "DirectionDescriptionType"
// direction descriptions are specified by a compass point code, a keyword, a textual description or a reference to a description.
// A gml:compassPoint is specified by a simple enumeration.  	
// In addition, thre elements to contain text-based descriptions of direction are provided.  
// If the direction is specified using a term from a list, gml:keyword should be used, and the list indicated using the value of the codeSpace attribute. 
// if the direction is decribed in prose, gml:direction or gml:reference should be used, allowing the value to be included inline or by reference.
type DirectionDescriptionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "compassPoint" of simple type gml:CompassPointEnumeration
	CompassPoint Enum_CompassPointEnumeration `xml:"compassPoint,omitempty"`

	// generated from element "keyword" of complex type CodeType
	Keyword *basictypes.CodeType `xml:"keyword,omitempty"`

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "reference" of complex type ReferenceType
	Reference_XSD *ReferenceType `xml:"reference,omitempty"`
}

// AbstractTopologyType Named source named complex type "AbstractTopologyType"
// This abstract type supplies the root or base type for all topological elements including primitives and complexes. It inherits AbstractGMLType and hence can be identified using the gml:id attribute.
type AbstractTopologyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties
}

// AbstractTopoPrimitiveType Named source named complex type "AbstractTopoPrimitiveType"
type AbstractTopoPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties
}

// NodeOrEdgePropertyType Named source named complex type "NodeOrEdgePropertyType"
type NodeOrEdgePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Node" of complex type NodeType
	Node *NodeType `xml:"Node,omitempty"`

	// generated from element "Edge" of complex type EdgeType
	Edge *EdgeType `xml:"Edge,omitempty"`
}

// NodePropertyType Named source named complex type "NodePropertyType"
type NodePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Node" of complex type NodeType
	Node *NodeType `xml:"Node,omitempty"`
}

// FaceOrTopoSolidPropertyType Named source named complex type "FaceOrTopoSolidPropertyType"
type FaceOrTopoSolidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Face" of complex type FaceType
	Face *FaceType `xml:"Face,omitempty"`

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid *TopoSolidType `xml:"TopoSolid,omitempty"`
}

// TopoSolidPropertyType Named source named complex type "TopoSolidPropertyType"
type TopoSolidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid *TopoSolidType `xml:"TopoSolid,omitempty"`
}

// NodeType Named source named complex type "NodeType"
type NodeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "container" of complex type FaceOrTopoSolidPropertyType
	Container *FaceOrTopoSolidPropertyType `xml:"container,omitempty"`

	// generated from element "directedEdge" of complex type DirectedEdgePropertyType
	DirectedEdge []*DirectedEdgePropertyType `xml:"directedEdge,omitempty"`

	// generated from element "pointProperty" of complex type PointPropertyType
	PointProperty *PointPropertyType `xml:"pointProperty,omitempty"`
}

// DirectedNodePropertyType Named source named complex type "DirectedNodePropertyType"
type DirectedNodePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Node" of complex type NodeType
	Node *NodeType `xml:"Node,omitempty"`
}

// EdgeType Named source named complex type "EdgeType"
type EdgeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "container" of complex type TopoSolidPropertyType
	Container *TopoSolidPropertyType `xml:"container,omitempty"`

	// generated from element "directedNode" of complex type DirectedNodePropertyType
	DirectedNode *DirectedNodePropertyType `xml:"directedNode,omitempty"`

	// generated from element "directedFace" of complex type DirectedFacePropertyType
	DirectedFace []*DirectedFacePropertyType `xml:"directedFace,omitempty"`

	// generated from element "curveProperty" of complex type CurvePropertyType
	CurveProperty *CurvePropertyType `xml:"curveProperty,omitempty"`
}

// DirectedEdgePropertyType Named source named complex type "DirectedEdgePropertyType"
type DirectedEdgePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Edge" of complex type EdgeType
	Edge *EdgeType `xml:"Edge,omitempty"`
}

// FaceType Named source named complex type "FaceType"
type FaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute "universal"
	Universal bool `xml:"universal,attr,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "isolated" of complex type NodePropertyType
	Isolated []*NodePropertyType `xml:"isolated,omitempty"`

	// generated from element "directedEdge" of complex type DirectedEdgePropertyType
	DirectedEdge []*DirectedEdgePropertyType `xml:"directedEdge,omitempty"`

	// generated from element "directedTopoSolid" of complex type DirectedTopoSolidPropertyType
	DirectedTopoSolid *DirectedTopoSolidPropertyType `xml:"directedTopoSolid,omitempty"`

	// generated from element "surfaceProperty" of complex type SurfacePropertyType
	SurfaceProperty *SurfacePropertyType `xml:"surfaceProperty,omitempty"`
}

// DirectedFacePropertyType Named source named complex type "DirectedFacePropertyType"
type DirectedFacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Face" of complex type FaceType
	Face *FaceType `xml:"Face,omitempty"`
}

// TopoSolidType Named source named complex type "TopoSolidType"
type TopoSolidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute "universal"
	Universal bool `xml:"universal,attr,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "isolated" of complex type NodeOrEdgePropertyType
	Isolated []*NodeOrEdgePropertyType `xml:"isolated,omitempty"`

	// generated from element "directedFace" of complex type DirectedFacePropertyType
	DirectedFace []*DirectedFacePropertyType `xml:"directedFace,omitempty"`

	// generated from element "solidProperty" of complex type SolidPropertyType
	SolidProperty *SolidPropertyType `xml:"solidProperty,omitempty"`
}

// DirectedTopoSolidPropertyType Named source named complex type "DirectedTopoSolidPropertyType"
type DirectedTopoSolidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "orientation"
	Orientation string `xml:"orientation,attr,omitempty"`

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid *TopoSolidType `xml:"TopoSolid,omitempty"`
}

// TopoPointType Named source named complex type "TopoPointType"
type TopoPointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "directedNode" of complex type DirectedNodePropertyType
	DirectedNode *DirectedNodePropertyType `xml:"directedNode,omitempty"`
}

// TopoPointPropertyType Named source named complex type "TopoPointPropertyType"
type TopoPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoPoint" of complex type TopoPointType
	TopoPoint *TopoPointType `xml:"TopoPoint,omitempty"`
}

// TopoCurveType Named source named complex type "TopoCurveType"
type TopoCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "directedEdge" of complex type DirectedEdgePropertyType
	DirectedEdge []*DirectedEdgePropertyType `xml:"directedEdge,omitempty"`
}

// TopoCurvePropertyType Named source named complex type "TopoCurvePropertyType"
type TopoCurvePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoCurve" of complex type TopoCurveType
	TopoCurve *TopoCurveType `xml:"TopoCurve,omitempty"`
}

// TopoSurfaceType Named source named complex type "TopoSurfaceType"
type TopoSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "directedFace" of complex type DirectedFacePropertyType
	DirectedFace []*DirectedFacePropertyType `xml:"directedFace,omitempty"`
}

// TopoSurfacePropertyType Named source named complex type "TopoSurfacePropertyType"
type TopoSurfacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoSurface" of complex type TopoSurfaceType
	TopoSurface *TopoSurfaceType `xml:"TopoSurface,omitempty"`
}

// TopoVolumeType Named source named complex type "TopoVolumeType"
type TopoVolumeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "directedTopoSolid" of complex type DirectedTopoSolidPropertyType
	DirectedTopoSolid []*DirectedTopoSolidPropertyType `xml:"directedTopoSolid,omitempty"`
}

// TopoVolumePropertyType Named source named complex type "TopoVolumePropertyType"
type TopoVolumePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TopoVolume" of complex type TopoVolumeType
	TopoVolume *TopoVolumeType `xml:"TopoVolume,omitempty"`
}

// TopoComplexType Named source named complex type "TopoComplexType"
type TopoComplexType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute "isMaximal"
	IsMaximal bool `xml:"isMaximal,attr,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "maximalComplex" of complex type TopoComplexPropertyType
	MaximalComplex *TopoComplexPropertyType `xml:"maximalComplex,omitempty"`

	// generated from element "superComplex" of complex type TopoComplexPropertyType
	SuperComplex []*TopoComplexPropertyType `xml:"superComplex,omitempty"`

	// generated from element "subComplex" of complex type TopoComplexPropertyType
	SubComplex []*TopoComplexPropertyType `xml:"subComplex,omitempty"`

	// generated from element "topoPrimitiveMember" of complex type TopoPrimitiveMemberType
	TopoPrimitiveMember []*TopoPrimitiveMemberType `xml:"topoPrimitiveMember,omitempty"`

	// generated from element "topoPrimitiveMembers" of complex type TopoPrimitiveArrayAssociationType
	TopoPrimitiveMembers *TopoPrimitiveArrayAssociationType `xml:"topoPrimitiveMembers,omitempty"`
}

// TopoPrimitiveMemberType Named source named complex type "TopoPrimitiveMemberType"
type TopoPrimitiveMemberType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Node" of complex type NodeType
	Node *NodeType `xml:"Node,omitempty"`

	// generated from element "Edge" of complex type EdgeType
	Edge *EdgeType `xml:"Edge,omitempty"`

	// generated from element "Face" of complex type FaceType
	Face *FaceType `xml:"Face,omitempty"`

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid *TopoSolidType `xml:"TopoSolid,omitempty"`
}

// TopoPrimitiveArrayAssociationType Named source named complex type "TopoPrimitiveArrayAssociationType"
type TopoPrimitiveArrayAssociationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "Node" of complex type NodeType
	Node []*NodeType `xml:"Node,omitempty"`

	// generated from element "Edge" of complex type EdgeType
	Edge []*EdgeType `xml:"Edge,omitempty"`

	// generated from element "Face" of complex type FaceType
	Face []*FaceType `xml:"Face,omitempty"`

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid []*TopoSolidType `xml:"TopoSolid,omitempty"`
}

// TopoComplexPropertyType Named source named complex type "TopoComplexPropertyType"
type TopoComplexPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TopoComplex" of complex type TopoComplexType
	TopoComplex *TopoComplexType `xml:"TopoComplex,omitempty"`
}

// GeometricComplexType Named source named complex type "GeometricComplexType"
type GeometricComplexType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "element" of complex type GeometricPrimitivePropertyType
	Element []*GeometricPrimitivePropertyType `xml:"element,omitempty"`
}

// GeometricComplexPropertyType Named source named complex type "GeometricComplexPropertyType"
// A property that has a geometric complex as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type GeometricComplexPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`
}

// CompositeCurveType Named source named complex type "CompositeCurveType"
type CompositeCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "curveMember" of complex type CurvePropertyType
	CurveMember []*CurvePropertyType `xml:"curveMember,omitempty"`
}

// CompositeSurfaceType Named source named complex type "CompositeSurfaceType"
type CompositeSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "surfaceMember" of complex type SurfacePropertyType
	SurfaceMember []*SurfacePropertyType `xml:"surfaceMember,omitempty"`
}

// CompositeSolidType Named source named complex type "CompositeSolidType"
type CompositeSolidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "solidMember" of complex type SolidPropertyType
	SolidMember []*SolidPropertyType `xml:"solidMember,omitempty"`
}

// AbstractCoverageType Named source named complex type "AbstractCoverageType"
// The base type for coverages is gml:AbstractCoverageType. The basic elements of a coverage can be seen in this content model: the coverage contains gml:domainSet and gml:rangeSet properties. The gml:domainSet property describes the domain of the coverage and the gml:rangeSet property describes the range of the coverage.
type AbstractCoverageType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "domainSet" of complex type DomainSetType
	DomainSet *DomainSetType `xml:"domainSet,omitempty"`

	// generated from element "multiPointDomain" of complex type DomainSetType
	MultiPointDomain *DomainSetType `xml:"multiPointDomain,omitempty"`

	// generated from element "multiCurveDomain" of complex type DomainSetType
	MultiCurveDomain *DomainSetType `xml:"multiCurveDomain,omitempty"`

	// generated from element "multiSurfaceDomain" of complex type DomainSetType
	MultiSurfaceDomain *DomainSetType `xml:"multiSurfaceDomain,omitempty"`

	// generated from element "multiSolidDomain" of complex type DomainSetType
	MultiSolidDomain *DomainSetType `xml:"multiSolidDomain,omitempty"`

	// generated from element "gridDomain" of complex type DomainSetType
	GridDomain *DomainSetType `xml:"gridDomain,omitempty"`

	// generated from element "rectifiedGridDomain" of complex type DomainSetType
	RectifiedGridDomain *DomainSetType `xml:"rectifiedGridDomain,omitempty"`

	// generated from element "rangeSet" of complex type RangeSetType
	RangeSet *RangeSetType `xml:"rangeSet,omitempty"`
}

// DiscreteCoverageType Named source named complex type "DiscreteCoverageType"
type DiscreteCoverageType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "domainSet" of complex type DomainSetType
	DomainSet *DomainSetType `xml:"domainSet,omitempty"`

	// generated from element "multiPointDomain" of complex type DomainSetType
	MultiPointDomain *DomainSetType `xml:"multiPointDomain,omitempty"`

	// generated from element "multiCurveDomain" of complex type DomainSetType
	MultiCurveDomain *DomainSetType `xml:"multiCurveDomain,omitempty"`

	// generated from element "multiSurfaceDomain" of complex type DomainSetType
	MultiSurfaceDomain *DomainSetType `xml:"multiSurfaceDomain,omitempty"`

	// generated from element "multiSolidDomain" of complex type DomainSetType
	MultiSolidDomain *DomainSetType `xml:"multiSolidDomain,omitempty"`

	// generated from element "gridDomain" of complex type DomainSetType
	GridDomain *DomainSetType `xml:"gridDomain,omitempty"`

	// generated from element "rectifiedGridDomain" of complex type DomainSetType
	RectifiedGridDomain *DomainSetType `xml:"rectifiedGridDomain,omitempty"`

	// generated from element "rangeSet" of complex type RangeSetType
	RangeSet *RangeSetType `xml:"rangeSet,omitempty"`

	// generated from element "coverageFunction" of complex type CoverageFunctionType
	CoverageFunction *CoverageFunctionType `xml:"coverageFunction,omitempty"`
}

// AbstractContinuousCoverageType Named source named complex type "AbstractContinuousCoverageType"
type AbstractContinuousCoverageType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "domainSet" of complex type DomainSetType
	DomainSet *DomainSetType `xml:"domainSet,omitempty"`

	// generated from element "multiPointDomain" of complex type DomainSetType
	MultiPointDomain *DomainSetType `xml:"multiPointDomain,omitempty"`

	// generated from element "multiCurveDomain" of complex type DomainSetType
	MultiCurveDomain *DomainSetType `xml:"multiCurveDomain,omitempty"`

	// generated from element "multiSurfaceDomain" of complex type DomainSetType
	MultiSurfaceDomain *DomainSetType `xml:"multiSurfaceDomain,omitempty"`

	// generated from element "multiSolidDomain" of complex type DomainSetType
	MultiSolidDomain *DomainSetType `xml:"multiSolidDomain,omitempty"`

	// generated from element "gridDomain" of complex type DomainSetType
	GridDomain *DomainSetType `xml:"gridDomain,omitempty"`

	// generated from element "rectifiedGridDomain" of complex type DomainSetType
	RectifiedGridDomain *DomainSetType `xml:"rectifiedGridDomain,omitempty"`

	// generated from element "rangeSet" of complex type RangeSetType
	RangeSet *RangeSetType `xml:"rangeSet,omitempty"`

	// generated from element "coverageFunction" of complex type CoverageFunctionType
	CoverageFunction *CoverageFunctionType `xml:"coverageFunction,omitempty"`
}

// DomainSetType Named source named complex type "DomainSetType"
type DomainSetType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`

	// generated from element "TimeTopologyComplex" of complex type TimeTopologyComplexType
	TimeTopologyComplex *TimeTopologyComplexType `xml:"TimeTopologyComplex,omitempty"`
}

// RangeSetType Named source named complex type "RangeSetType"
type RangeSetType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ValueArray" of complex type ValueArrayType
	ValueArray []*ValueArrayType `xml:"ValueArray,omitempty"`

	// generated from element "BooleanList" of simple type gml:booleanOrNilReasonList
	BooleanList string `xml:"BooleanList,omitempty"`

	// generated from element "CategoryList" of complex type CodeOrNilReasonListType
	CategoryList []*basictypes.CodeOrNilReasonListType `xml:"CategoryList,omitempty"`

	// generated from element "CountList" of simple type gml:integerOrNilReasonList
	CountList string `xml:"CountList,omitempty"`

	// generated from element "QuantityList" of complex type MeasureOrNilReasonListType
	QuantityList []*basictypes.MeasureOrNilReasonListType `xml:"QuantityList,omitempty"`

	// generated from element "DataBlock" of complex type DataBlockType
	DataBlock *DataBlockType `xml:"DataBlock,omitempty"`

	// generated from element "File" of complex type FileType
	File *FileType `xml:"File,omitempty"`
}

// DataBlockType Named source named complex type "DataBlockType"
type DataBlockType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "rangeParameters" of complex type AssociationRoleType
	RangeParameters *AssociationRoleType `xml:"rangeParameters,omitempty"`

	// generated from element "tupleList" of complex type CoordinatesType
	TupleList *basictypes.CoordinatesType `xml:"tupleList,omitempty"`

	// generated from element "doubleOrNilReasonTupleList" of simple type gml:doubleOrNilReasonList
	DoubleOrNilReasonTupleList string `xml:"doubleOrNilReasonTupleList,omitempty"`
}

// FileType Named source named complex type "FileType"
type FileType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "rangeParameters" of complex type AssociationRoleType
	RangeParameters *AssociationRoleType `xml:"rangeParameters,omitempty"`

	// generated from element "fileName" of simple type anyURI
	FileName string `xml:"fileName,omitempty"`

	// generated from element "fileReference" of simple type anyURI
	FileReference string `xml:"fileReference,omitempty"`

	// generated from element "fileStructure" of complex type CodeType
	FileStructure *basictypes.CodeType `xml:"fileStructure,omitempty"`

	// generated from element "mimeType" of simple type anyURI
	MimeType string `xml:"mimeType,omitempty"`

	// generated from element "compression" of simple type anyURI
	Compression string `xml:"compression,omitempty"`
}

// CoverageFunctionType Named source named complex type "CoverageFunctionType"
type CoverageFunctionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "MappingRule" of complex type StringOrRefType
	MappingRule *StringOrRefType `xml:"MappingRule,omitempty"`

	// generated from element "CoverageMappingRule" of complex type MappingRuleType
	CoverageMappingRule *MappingRuleType `xml:"CoverageMappingRule,omitempty"`

	// generated from element "GridFunction" of complex type GridFunctionType
	GridFunction *GridFunctionType `xml:"GridFunction,omitempty"`
}

// MappingRuleType Named source named complex type "MappingRuleType"
type MappingRuleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ruleDefinition" of simple type string
	RuleDefinition string `xml:"ruleDefinition,omitempty"`

	// generated from element "ruleReference" of complex type ReferenceType
	RuleReference *ReferenceType `xml:"ruleReference,omitempty"`
}

// GridFunctionType Named source named complex type "GridFunctionType"
type GridFunctionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sequenceRule" of complex type SequenceRuleType
	SequenceRule *SequenceRuleType `xml:"sequenceRule,omitempty"`

	// generated from element "startPoint" of simple type gml:integerList
	StartPoint string `xml:"startPoint,omitempty"`
}

// SequenceRuleType Named source named complex type "SequenceRuleType"
// The gml:SequenceRuleType is derived from the gml:SequenceRuleEnumeration through the addition of an axisOrder attribute.  The gml:SequenceRuleEnumeration is an enumerated type. The rule names are defined in ISO 19123. If no rule name is specified the default is "Linear".
type SequenceRuleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "order"
	Order Enum_IncrementOrder `xml:"order,attr,omitempty"`

	// generated from attribute "axisOrder"
	AxisOrder string `xml:"axisOrder,attr,omitempty"`

	// simpleContent character data
	EnclosedText Enum_SequenceRuleEnumeration `xml:",chardata"`
}

// Boolean Named source element Boolean within schema
type Boolean struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// simpleContent character data
	EnclosedText bool `xml:",chardata"`
}

// Category Named source element Category within schema
// A gml:Category has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
type Category struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Count Named source element Count within schema
type Count struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// simpleContent character data
	EnclosedText int `xml:",chardata"`
}

// Quantity Named source element Quantity within schema
// An XML attribute uom ("unit of measure") is required, whose value is a URI which identifies the definition of a ratio scale or units by which the numeric value shall be multiplied, or an interval or position scale on which the value occurs.
type Quantity struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Group_Value UnNamed source named group "Value"
type Group_Value struct {

	// insertion point for fields

	// generated from anonymous complex type within element "Boolean"
	Boolean *Boolean `xml:"Boolean,omitempty"`

	// generated from anonymous complex type within element "Category"
	Category *Category `xml:"Category,omitempty"`

	// generated from anonymous complex type within element "Count"
	Count *Count `xml:"Count,omitempty"`

	// generated from anonymous complex type within element "Quantity"
	Quantity *Quantity `xml:"Quantity,omitempty"`

	// generated from element "BooleanList" of simple type gml:booleanOrNilReasonList
	BooleanList string `xml:"BooleanList,omitempty"`

	// generated from element "CategoryList" of complex type CodeOrNilReasonListType
	CategoryList *basictypes.CodeOrNilReasonListType `xml:"CategoryList,omitempty"`

	// generated from element "CountList" of simple type gml:integerOrNilReasonList
	CountList string `xml:"CountList,omitempty"`

	// generated from element "QuantityList" of complex type MeasureOrNilReasonListType
	QuantityList *basictypes.MeasureOrNilReasonListType `xml:"QuantityList,omitempty"`

	// generated from element "CompositeValue" of complex type CompositeValueType
	CompositeValue *CompositeValueType `xml:"CompositeValue,omitempty"`

	// generated from element "ValueArray" of complex type ValueArrayType
	ValueArray *ValueArrayType `xml:"ValueArray,omitempty"`

	// generated from element "CategoryExtent" of complex type CategoryExtentType
	CategoryExtent *CategoryExtentType `xml:"CategoryExtent,omitempty"`

	// generated from element "CountExtent" of simple type gml:CountExtentType
	CountExtent string `xml:"CountExtent,omitempty"`

	// generated from element "QuantityExtent" of complex type QuantityExtentType
	QuantityExtent *QuantityExtentType `xml:"QuantityExtent,omitempty"`

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`

	// generated from element "TimeTopologyComplex" of complex type TimeTopologyComplexType
	TimeTopologyComplex *TimeTopologyComplexType `xml:"TimeTopologyComplex,omitempty"`

	// generated from element "Null" of simple type gml:NilReasonType
	Null string `xml:"Null,omitempty"`
}

// ValuePropertyType Named source named complex type "ValuePropertyType"
type ValuePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from group "gml:Value" with order 1218 depth 2
	Group_Value
}

// ValueArrayPropertyType Named source named complex type "ValueArrayPropertyType"
type ValueArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from group "gml:Value" with order 1224 depth 2
	Group_Value
}

// CompositeValueType Named source named complex type "CompositeValueType"
type CompositeValueType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "valueComponent" of complex type ValuePropertyType
	ValueComponent []*ValuePropertyType `xml:"valueComponent,omitempty"`

	// generated from element "valueComponents" of complex type ValueArrayPropertyType
	ValueComponents *ValueArrayPropertyType `xml:"valueComponents,omitempty"`
}

// ValueArrayType Named source named complex type "ValueArrayType"
type ValueArrayType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "valueComponent" of complex type ValuePropertyType
	ValueComponent []*ValuePropertyType `xml:"valueComponent,omitempty"`

	// generated from element "valueComponents" of complex type ValueArrayPropertyType
	ValueComponents *ValueArrayPropertyType `xml:"valueComponents,omitempty"`

	// generated from attribute group "gml:referenceSystem"
	AttributeGroup_referenceSystem
}

// AttributeGroup_referenceSystem UnNamed source named attribute group "referenceSystem"
type AttributeGroup_referenceSystem struct {

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`
}

// CategoryExtentType Named source named complex type "CategoryExtentType"
type CategoryExtentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// QuantityExtentType Named source named complex type "QuantityExtentType"
type QuantityExtentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// BooleanPropertyType Named source named complex type "BooleanPropertyType"
type BooleanPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from anonymous complex type within element "Boolean"
	Boolean *Boolean `xml:"Boolean,omitempty"`
}

// CategoryPropertyType Named source named complex type "CategoryPropertyType"
type CategoryPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from anonymous complex type within element "Category"
	Category *Category `xml:"Category,omitempty"`
}

// QuantityPropertyType Named source named complex type "QuantityPropertyType"
type QuantityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from anonymous complex type within element "Quantity"
	Quantity *Quantity `xml:"Quantity,omitempty"`
}

// CountPropertyType Named source named complex type "CountPropertyType"
type CountPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from anonymous complex type within element "Count"
	Count *Count `xml:"Count,omitempty"`
}

// GridType Named source named complex type "GridType"
type GridType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute "dimension"
	Dimension int `xml:"dimension,attr,omitempty"`

	// generated from element "limits" of complex type GridLimitsType
	Limits *GridLimitsType `xml:"limits,omitempty"`

	// generated from element "axisLabels" of simple type gml:NCNameList
	AxisLabels_1 string `xml:"axisLabels,omitempty"`

	// generated from element "axisName" of simple type string
	AxisName string `xml:"axisName,omitempty"`
}

// GridLimitsType Named source named complex type "GridLimitsType"
type GridLimitsType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "GridEnvelope" of complex type GridEnvelopeType
	GridEnvelope *GridEnvelopeType `xml:"GridEnvelope,omitempty"`
}

// GridEnvelopeType Named source named complex type "GridEnvelopeType"
type GridEnvelopeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "low" of simple type gml:integerList
	Low string `xml:"low,omitempty"`

	// generated from element "high" of simple type gml:integerList
	High string `xml:"high,omitempty"`
}

// RectifiedGridType Named source named complex type "RectifiedGridType"
type RectifiedGridType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute "dimension"
	Dimension int `xml:"dimension,attr,omitempty"`

	// generated from element "limits" of complex type GridLimitsType
	Limits *GridLimitsType `xml:"limits,omitempty"`

	// generated from element "axisLabels" of simple type gml:NCNameList
	AxisLabels_1 string `xml:"axisLabels,omitempty"`

	// generated from element "axisName" of simple type string
	AxisName string `xml:"axisName,omitempty"`

	// generated from element "origin" of complex type PointPropertyType
	Origin *PointPropertyType `xml:"origin,omitempty"`

	// generated from element "offsetVector" of complex type VectorType
	OffsetVector []*VectorType `xml:"offsetVector,omitempty"`
}

// SingleCRSPropertyType Named source named complex type "SingleCRSPropertyType"
// gml:SingleCRSPropertyType is a property type for association roles to a single coordinate reference system, either referencing or containing the definition of that coordinate reference system.
type SingleCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS *ProjectedCRSType `xml:"ProjectedCRS,omitempty"`

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS *DerivedCRSType `xml:"DerivedCRS,omitempty"`

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS *GeodeticCRSType `xml:"GeodeticCRS,omitempty"`

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS *VerticalCRSType `xml:"VerticalCRS,omitempty"`

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS *EngineeringCRSType `xml:"EngineeringCRS,omitempty"`

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS *ImageCRSType `xml:"ImageCRS,omitempty"`

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS *TemporalCRSType `xml:"TemporalCRS,omitempty"`

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS *GeographicCRSType `xml:"GeographicCRS,omitempty"`

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS *GeocentricCRSType `xml:"GeocentricCRS,omitempty"`
}

// AbstractGeneralDerivedCRSType Named source named complex type "AbstractGeneralDerivedCRSType"
type AbstractGeneralDerivedCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "conversion" of complex type GeneralConversionPropertyType
	Conversion *GeneralConversionPropertyType `xml:"conversion,omitempty"`

	// generated from element "definedByConversion" of complex type GeneralConversionPropertyType
	DefinedByConversion *GeneralConversionPropertyType `xml:"definedByConversion,omitempty"`
}

// CompoundCRSType Named source named complex type "CompoundCRSType"
type CompoundCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "componentReferenceSystem" of complex type SingleCRSPropertyType
	ComponentReferenceSystem []*SingleCRSPropertyType `xml:"componentReferenceSystem,omitempty"`

	// generated from element "includesSingleCRS" of complex type SingleCRSPropertyType
	IncludesSingleCRS []*SingleCRSPropertyType `xml:"includesSingleCRS,omitempty"`
}

// CompoundCRSPropertyType Named source named complex type "CompoundCRSPropertyType"
// gml:CompoundCRSPropertyType is a property type for association roles to a compound coordinate reference system, either referencing or containing the definition of that reference system.
type CompoundCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "CompoundCRS" of complex type CompoundCRSType
	CompoundCRS *CompoundCRSType `xml:"CompoundCRS,omitempty"`
}

// GeodeticCRSType Named source named complex type "GeodeticCRSType"
// gml:GeodeticCRS is a coordinate reference system based on a geodetic datum.
type GeodeticCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "ellipsoidalCS" of complex type EllipsoidalCSPropertyType
	EllipsoidalCS *EllipsoidalCSPropertyType `xml:"ellipsoidalCS,omitempty"`

	// generated from element "usesEllipsoidalCS" of complex type EllipsoidalCSPropertyType
	UsesEllipsoidalCS *EllipsoidalCSPropertyType `xml:"usesEllipsoidalCS,omitempty"`

	// generated from element "cartesianCS" of complex type CartesianCSPropertyType
	CartesianCS *CartesianCSPropertyType `xml:"cartesianCS,omitempty"`

	// generated from element "usesCartesianCS" of complex type CartesianCSPropertyType
	UsesCartesianCS *CartesianCSPropertyType `xml:"usesCartesianCS,omitempty"`

	// generated from element "sphericalCS" of complex type SphericalCSPropertyType
	SphericalCS *SphericalCSPropertyType `xml:"sphericalCS,omitempty"`

	// generated from element "usesSphericalCS" of complex type SphericalCSPropertyType
	UsesSphericalCS *SphericalCSPropertyType `xml:"usesSphericalCS,omitempty"`

	// generated from element "geodeticDatum" of complex type GeodeticDatumPropertyType
	GeodeticDatum *GeodeticDatumPropertyType `xml:"geodeticDatum,omitempty"`

	// generated from element "usesGeodeticDatum" of complex type GeodeticDatumPropertyType
	UsesGeodeticDatum *GeodeticDatumPropertyType `xml:"usesGeodeticDatum,omitempty"`
}

// GeodeticCRSPropertyType Named source named complex type "GeodeticCRSPropertyType"
// gml:GeodeticCRSPropertyType is a property type for association roles to a geodetic coordinate reference system, either referencing or containing the definition of that reference system.
type GeodeticCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS *GeodeticCRSType `xml:"GeodeticCRS,omitempty"`
}

// VerticalCRSType Named source named complex type "VerticalCRSType"
type VerticalCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "verticalCS" of complex type VerticalCSPropertyType
	VerticalCS *VerticalCSPropertyType `xml:"verticalCS,omitempty"`

	// generated from element "usesVerticalCS" of complex type VerticalCSPropertyType
	UsesVerticalCS *VerticalCSPropertyType `xml:"usesVerticalCS,omitempty"`

	// generated from element "verticalDatum" of complex type VerticalDatumPropertyType
	VerticalDatum *VerticalDatumPropertyType `xml:"verticalDatum,omitempty"`

	// generated from element "usesVerticalDatum" of complex type VerticalDatumPropertyType
	UsesVerticalDatum *VerticalDatumPropertyType `xml:"usesVerticalDatum,omitempty"`
}

// VerticalCRSPropertyType Named source named complex type "VerticalCRSPropertyType"
// gml:VerticalCRSPropertyType is a property type for association roles to a vertical coordinate reference system, either referencing or containing the definition of that reference system.
type VerticalCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS *VerticalCRSType `xml:"VerticalCRS,omitempty"`
}

// ProjectedCRSType Named source named complex type "ProjectedCRSType"
type ProjectedCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "conversion" of complex type GeneralConversionPropertyType
	Conversion *GeneralConversionPropertyType `xml:"conversion,omitempty"`

	// generated from element "definedByConversion" of complex type GeneralConversionPropertyType
	DefinedByConversion *GeneralConversionPropertyType `xml:"definedByConversion,omitempty"`

	// generated from element "baseGeodeticCRS" of complex type GeodeticCRSPropertyType
	BaseGeodeticCRS *GeodeticCRSPropertyType `xml:"baseGeodeticCRS,omitempty"`

	// generated from element "baseGeographicCRS" of complex type GeographicCRSPropertyType
	BaseGeographicCRS *GeographicCRSPropertyType `xml:"baseGeographicCRS,omitempty"`

	// generated from element "cartesianCS" of complex type CartesianCSPropertyType
	CartesianCS *CartesianCSPropertyType `xml:"cartesianCS,omitempty"`

	// generated from element "usesCartesianCS" of complex type CartesianCSPropertyType
	UsesCartesianCS *CartesianCSPropertyType `xml:"usesCartesianCS,omitempty"`
}

// ProjectedCRSPropertyType Named source named complex type "ProjectedCRSPropertyType"
// gml:ProjectedCRSPropertyType is a property type for association roles to a projected coordinate reference system, either referencing or containing the definition of that reference system.
type ProjectedCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS *ProjectedCRSType `xml:"ProjectedCRS,omitempty"`
}

// DerivedCRSType Named source named complex type "DerivedCRSType"
type DerivedCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "conversion" of complex type GeneralConversionPropertyType
	Conversion *GeneralConversionPropertyType `xml:"conversion,omitempty"`

	// generated from element "definedByConversion" of complex type GeneralConversionPropertyType
	DefinedByConversion *GeneralConversionPropertyType `xml:"definedByConversion,omitempty"`

	// generated from element "baseCRS" of complex type SingleCRSPropertyType
	BaseCRS *SingleCRSPropertyType `xml:"baseCRS,omitempty"`

	// generated from element "derivedCRSType" of complex type CodeWithAuthorityType
	DerivedCRSType *basictypes.CodeWithAuthorityType `xml:"derivedCRSType,omitempty"`

	// generated from element "coordinateSystem" of complex type CoordinateSystemPropertyType
	CoordinateSystem *CoordinateSystemPropertyType `xml:"coordinateSystem,omitempty"`

	// generated from element "usesCS" of complex type CoordinateSystemPropertyType
	UsesCS *CoordinateSystemPropertyType `xml:"usesCS,omitempty"`
}

// DerivedCRSPropertyType Named source named complex type "DerivedCRSPropertyType"
// gml:DerivedCRSPropertyType is a property type for association roles to a non-projected derived coordinate reference system, either referencing or containing the definition of that reference system.
type DerivedCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS *DerivedCRSType `xml:"DerivedCRS,omitempty"`
}

// EngineeringCRSType Named source named complex type "EngineeringCRSType"
type EngineeringCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "affineCS" of complex type AffineCSPropertyType
	AffineCS *AffineCSPropertyType `xml:"affineCS,omitempty"`

	// generated from element "usesAffineCS" of complex type AffineCSPropertyType
	UsesAffineCS *AffineCSPropertyType `xml:"usesAffineCS,omitempty"`

	// generated from element "cartesianCS" of complex type CartesianCSPropertyType
	CartesianCS *CartesianCSPropertyType `xml:"cartesianCS,omitempty"`

	// generated from element "usesCartesianCS" of complex type CartesianCSPropertyType
	UsesCartesianCS *CartesianCSPropertyType `xml:"usesCartesianCS,omitempty"`

	// generated from element "cylindricalCS" of complex type CylindricalCSPropertyType
	CylindricalCS *CylindricalCSPropertyType `xml:"cylindricalCS,omitempty"`

	// generated from element "linearCS" of complex type LinearCSPropertyType
	LinearCS *LinearCSPropertyType `xml:"linearCS,omitempty"`

	// generated from element "polarCS" of complex type PolarCSPropertyType
	PolarCS *PolarCSPropertyType `xml:"polarCS,omitempty"`

	// generated from element "sphericalCS" of complex type SphericalCSPropertyType
	SphericalCS *SphericalCSPropertyType `xml:"sphericalCS,omitempty"`

	// generated from element "usesSphericalCS" of complex type SphericalCSPropertyType
	UsesSphericalCS *SphericalCSPropertyType `xml:"usesSphericalCS,omitempty"`

	// generated from element "userDefinedCS" of complex type UserDefinedCSPropertyType
	UserDefinedCS *UserDefinedCSPropertyType `xml:"userDefinedCS,omitempty"`

	// generated from element "coordinateSystem" of complex type CoordinateSystemPropertyType
	CoordinateSystem *CoordinateSystemPropertyType `xml:"coordinateSystem,omitempty"`

	// generated from element "usesCS" of complex type CoordinateSystemPropertyType
	UsesCS *CoordinateSystemPropertyType `xml:"usesCS,omitempty"`

	// generated from element "engineeringDatum" of complex type EngineeringDatumPropertyType
	EngineeringDatum *EngineeringDatumPropertyType `xml:"engineeringDatum,omitempty"`

	// generated from element "usesEngineeringDatum" of complex type EngineeringDatumPropertyType
	UsesEngineeringDatum *EngineeringDatumPropertyType `xml:"usesEngineeringDatum,omitempty"`
}

// EngineeringCRSPropertyType Named source named complex type "EngineeringCRSPropertyType"
// gml:EngineeringCRSPropertyType is a property type for association roles to an engineering coordinate reference system, either referencing or containing the definition of that reference system.
type EngineeringCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS *EngineeringCRSType `xml:"EngineeringCRS,omitempty"`
}

// ImageCRSType Named source named complex type "ImageCRSType"
type ImageCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "cartesianCS" of complex type CartesianCSPropertyType
	CartesianCS *CartesianCSPropertyType `xml:"cartesianCS,omitempty"`

	// generated from element "usesCartesianCS" of complex type CartesianCSPropertyType
	UsesCartesianCS *CartesianCSPropertyType `xml:"usesCartesianCS,omitempty"`

	// generated from element "affineCS" of complex type AffineCSPropertyType
	AffineCS *AffineCSPropertyType `xml:"affineCS,omitempty"`

	// generated from element "usesAffineCS" of complex type AffineCSPropertyType
	UsesAffineCS *AffineCSPropertyType `xml:"usesAffineCS,omitempty"`

	// generated from element "usesObliqueCartesianCS" of complex type ObliqueCartesianCSPropertyType
	UsesObliqueCartesianCS *ObliqueCartesianCSPropertyType `xml:"usesObliqueCartesianCS,omitempty"`

	// generated from element "imageDatum" of complex type ImageDatumPropertyType
	ImageDatum *ImageDatumPropertyType `xml:"imageDatum,omitempty"`

	// generated from element "usesImageDatum" of complex type ImageDatumPropertyType
	UsesImageDatum *ImageDatumPropertyType `xml:"usesImageDatum,omitempty"`
}

// ImageCRSPropertyType Named source named complex type "ImageCRSPropertyType"
// gml:ImageCRSPropertyType is a property type for association roles to an image coordinate reference system, either referencing or containing the definition of that reference system.
type ImageCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS *ImageCRSType `xml:"ImageCRS,omitempty"`
}

// TemporalCRSType Named source named complex type "TemporalCRSType"
type TemporalCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "timeCS" of complex type TimeCSPropertyType
	TimeCS *TimeCSPropertyType `xml:"timeCS,omitempty"`

	// generated from element "usesTimeCS" of complex type TimeCSPropertyType
	UsesTimeCS *TimeCSPropertyType `xml:"usesTimeCS,omitempty"`

	// generated from element "usesTemporalCS" of complex type TemporalCSPropertyType
	UsesTemporalCS *TemporalCSPropertyType `xml:"usesTemporalCS,omitempty"`

	// generated from element "temporalDatum" of complex type TemporalDatumPropertyType
	TemporalDatum *TemporalDatumPropertyType `xml:"temporalDatum,omitempty"`

	// generated from element "usesTemporalDatum" of complex type TemporalDatumPropertyType
	UsesTemporalDatum *TemporalDatumPropertyType `xml:"usesTemporalDatum,omitempty"`
}

// TemporalCRSPropertyType Named source named complex type "TemporalCRSPropertyType"
// gml:TemporalCRSPropertyType is a property type for association roles to a temporal coordinate reference system, either referencing or containing the definition of that reference system.
type TemporalCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS *TemporalCRSType `xml:"TemporalCRS,omitempty"`
}

// CoordinateSystemAxisType Named source named complex type "CoordinateSystemAxisType"
type CoordinateSystemAxisType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// generated from element "axisAbbrev" of complex type CodeType
	AxisAbbrev *basictypes.CodeType `xml:"axisAbbrev,omitempty"`

	// generated from element "axisDirection" of complex type CodeWithAuthorityType
	AxisDirection *basictypes.CodeWithAuthorityType `xml:"axisDirection,omitempty"`

	// generated from element "minimumValue" of simple type double
	MinimumValue float64 `xml:"minimumValue,omitempty"`

	// generated from element "maximumValue" of simple type double
	MaximumValue float64 `xml:"maximumValue,omitempty"`

	// generated from element "rangeMeaning" of complex type CodeWithAuthorityType
	RangeMeaning *basictypes.CodeWithAuthorityType `xml:"rangeMeaning,omitempty"`
}

// CoordinateSystemAxisPropertyType Named source named complex type "CoordinateSystemAxisPropertyType"
// gml:CoordinateSystemAxisPropertyType is a property type for association roles to a coordinate system axis, either referencing or containing the definition of that axis.
type CoordinateSystemAxisPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "CoordinateSystemAxis" of complex type CoordinateSystemAxisType
	CoordinateSystemAxis *CoordinateSystemAxisType `xml:"CoordinateSystemAxis,omitempty"`
}

// AbstractCoordinateSystemType Named source named complex type "AbstractCoordinateSystemType"
type AbstractCoordinateSystemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// CoordinateSystemPropertyType Named source named complex type "CoordinateSystemPropertyType"
// gml:CoordinateSystemPropertyType is a property type for association roles to a coordinate system, either referencing or containing the definition of that coordinate system.
type CoordinateSystemPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "EllipsoidalCS" of complex type EllipsoidalCSType
	EllipsoidalCS *EllipsoidalCSType `xml:"EllipsoidalCS,omitempty"`

	// generated from element "CartesianCS" of complex type CartesianCSType
	CartesianCS *CartesianCSType `xml:"CartesianCS,omitempty"`

	// generated from element "VerticalCS" of complex type VerticalCSType
	VerticalCS *VerticalCSType `xml:"VerticalCS,omitempty"`

	// generated from element "TimeCS" of complex type TimeCSType
	TimeCS *TimeCSType `xml:"TimeCS,omitempty"`

	// generated from element "LinearCS" of complex type LinearCSType
	LinearCS *LinearCSType `xml:"LinearCS,omitempty"`

	// generated from element "UserDefinedCS" of complex type UserDefinedCSType
	UserDefinedCS *UserDefinedCSType `xml:"UserDefinedCS,omitempty"`

	// generated from element "SphericalCS" of complex type SphericalCSType
	SphericalCS *SphericalCSType `xml:"SphericalCS,omitempty"`

	// generated from element "PolarCS" of complex type PolarCSType
	PolarCS *PolarCSType `xml:"PolarCS,omitempty"`

	// generated from element "CylindricalCS" of complex type CylindricalCSType
	CylindricalCS *CylindricalCSType `xml:"CylindricalCS,omitempty"`

	// generated from element "AffineCS" of complex type AffineCSType
	AffineCS *AffineCSType `xml:"AffineCS,omitempty"`

	// generated from element "TemporalCS" of complex type TemporalCSType
	TemporalCS *TemporalCSType `xml:"TemporalCS,omitempty"`

	// generated from element "ObliqueCartesianCS" of complex type ObliqueCartesianCSType
	ObliqueCartesianCS *ObliqueCartesianCSType `xml:"ObliqueCartesianCS,omitempty"`
}

// EllipsoidalCSType Named source named complex type "EllipsoidalCSType"
type EllipsoidalCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// EllipsoidalCSPropertyType Named source named complex type "EllipsoidalCSPropertyType"
// gml:EllipsoidalCSPropertyType is a property type for association roles to an ellipsoidal coordinate system, either referencing or containing the definition of that coordinate system.
type EllipsoidalCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "EllipsoidalCS" of complex type EllipsoidalCSType
	EllipsoidalCS *EllipsoidalCSType `xml:"EllipsoidalCS,omitempty"`
}

// CartesianCSType Named source named complex type "CartesianCSType"
type CartesianCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// CartesianCSPropertyType Named source named complex type "CartesianCSPropertyType"
// gml:CartesianCSPropertyType is a property type for association roles to a Cartesian coordinate system, either referencing or containing the definition of that coordinate system.
type CartesianCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "CartesianCS" of complex type CartesianCSType
	CartesianCS *CartesianCSType `xml:"CartesianCS,omitempty"`
}

// VerticalCSType Named source named complex type "VerticalCSType"
type VerticalCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// VerticalCSPropertyType Named source named complex type "VerticalCSPropertyType"
// gml:VerticalCSPropertyType is a property type for association roles to a vertical coordinate system, either referencing or containing the definition of that coordinate system.
type VerticalCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "VerticalCS" of complex type VerticalCSType
	VerticalCS *VerticalCSType `xml:"VerticalCS,omitempty"`
}

// TimeCSType Named source named complex type "TimeCSType"
type TimeCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// TimeCSPropertyType Named source named complex type "TimeCSPropertyType"
// gml:TimeCSPropertyType is a property type for association roles to a time coordinate system, either referencing or containing the definition of that coordinate system.
type TimeCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeCS" of complex type TimeCSType
	TimeCS *TimeCSType `xml:"TimeCS,omitempty"`
}

// LinearCSType Named source named complex type "LinearCSType"
type LinearCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// LinearCSPropertyType Named source named complex type "LinearCSPropertyType"
// gml:LinearCSPropertyType is a property type for association roles to a linear coordinate system, either referencing or containing the definition of that coordinate system.
type LinearCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "LinearCS" of complex type LinearCSType
	LinearCS *LinearCSType `xml:"LinearCS,omitempty"`
}

// UserDefinedCSType Named source named complex type "UserDefinedCSType"
type UserDefinedCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// UserDefinedCSPropertyType Named source named complex type "UserDefinedCSPropertyType"
// gml:UserDefinedCSPropertyType is a property type for association roles to a user-defined coordinate system, either referencing or containing the definition of that coordinate system.
type UserDefinedCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "UserDefinedCS" of complex type UserDefinedCSType
	UserDefinedCS *UserDefinedCSType `xml:"UserDefinedCS,omitempty"`
}

// SphericalCSType Named source named complex type "SphericalCSType"
type SphericalCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// SphericalCSPropertyType Named source named complex type "SphericalCSPropertyType"
// gml:SphericalCSPropertyType is property type for association roles to a spherical coordinate system, either referencing or containing the definition of that coordinate system.
type SphericalCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "SphericalCS" of complex type SphericalCSType
	SphericalCS *SphericalCSType `xml:"SphericalCS,omitempty"`
}

// PolarCSType Named source named complex type "PolarCSType"
type PolarCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// PolarCSPropertyType Named source named complex type "PolarCSPropertyType"
// gml:PolarCSPropertyType is a property type for association roles to a polar coordinate system, either referencing or containing the definition of that coordinate system.
type PolarCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "PolarCS" of complex type PolarCSType
	PolarCS *PolarCSType `xml:"PolarCS,omitempty"`
}

// CylindricalCSType Named source named complex type "CylindricalCSType"
type CylindricalCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// CylindricalCSPropertyType Named source named complex type "CylindricalCSPropertyType"
// gml:CylindricalCSPropertyType is a property type for association roles to a cylindrical coordinate system, either referencing or containing the definition of that coordinate system.
type CylindricalCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "CylindricalCS" of complex type CylindricalCSType
	CylindricalCS *CylindricalCSType `xml:"CylindricalCS,omitempty"`
}

// AffineCSType Named source named complex type "AffineCSType"
type AffineCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// AffineCSPropertyType Named source named complex type "AffineCSPropertyType"
// gml:AffineCSPropertyType is a property type for association roles to an affine coordinate system, either referencing or containing the definition of that coordinate system.
type AffineCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "AffineCS" of complex type AffineCSType
	AffineCS *AffineCSType `xml:"AffineCS,omitempty"`
}

// IdentifiedObjectType Named source named complex type "IdentifiedObjectType"
// gml:IdentifiedObjectType provides identification properties of a CRS-related object. In gml:DefinitionType, the gml:identifier element shall be the primary name by which this object is identified, encoding the "name" attribute in the UML model.
// Zero or more of the gml:name elements can be an unordered set of "identifiers", encoding the "identifier" attribute in the UML model. Each of these gml:name elements can reference elsewhere the object's defining information or be an identifier by which this object can be referenced.
// Zero or more other gml:name elements can be an unordered set of "alias" alternative names by which this CRS related object is identified, encoding the "alias" attributes in the UML model. An object may have several aliases, typically used in different contexts. The context for an alias is indicated by the value of its (optional) codeSpace attribute.
// Any needed version information shall be included in the codeSpace attribute of a gml:identifier and gml:name elements. In this use, the gml:remarks element in the gml:DefinitionType shall contain comments on or information about this object, including data source information.
type IdentifiedObjectType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`
}

// AbstractCRSType Named source named complex type "AbstractCRSType"
type AbstractCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`
}

// DomainOfValidity Named source element domainOfValidity within schema
// The gml:domainOfValidity property implements an association role to an EX_Extent object as encoded in ISO/TS 19139, either referencing or containing the definition of that extent.
type DomainOfValidity struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "EX_Extent" of complex type EX_Extent_Type
	EX_Extent *EX_Extent_Type `xml:"EX_Extent,omitempty"`
}

// CRSPropertyType Named source named complex type "CRSPropertyType"
// gml:CRSPropertyType is a property type for association roles to a CRS abstract coordinate reference system, either referencing or containing the definition of that CRS.
type CRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS *ProjectedCRSType `xml:"ProjectedCRS,omitempty"`

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS *DerivedCRSType `xml:"DerivedCRS,omitempty"`

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS *GeodeticCRSType `xml:"GeodeticCRS,omitempty"`

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS *VerticalCRSType `xml:"VerticalCRS,omitempty"`

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS *EngineeringCRSType `xml:"EngineeringCRS,omitempty"`

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS *ImageCRSType `xml:"ImageCRS,omitempty"`

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS *TemporalCRSType `xml:"TemporalCRS,omitempty"`

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS *GeographicCRSType `xml:"GeographicCRS,omitempty"`

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS *GeocentricCRSType `xml:"GeocentricCRS,omitempty"`

	// generated from element "CompoundCRS" of complex type CompoundCRSType
	CompoundCRS *CompoundCRSType `xml:"CompoundCRS,omitempty"`
}

// CI_ResponsibleParty_Type Named source named complex type "CI_ResponsibleParty_Type"
// Identification of, and means of communication with, person(s) and organisations associated with the dataset
type CI_ResponsibleParty_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "individualName" of complex type CharacterString_PropertyType
	IndividualName *CharacterString_PropertyType `xml:"individualName,omitempty"`

	// generated from element "organisationName" of complex type CharacterString_PropertyType
	OrganisationName *CharacterString_PropertyType `xml:"organisationName,omitempty"`

	// generated from element "positionName" of complex type CharacterString_PropertyType
	PositionName *CharacterString_PropertyType `xml:"positionName,omitempty"`

	// generated from element "contactInfo" of complex type CI_Contact_PropertyType
	ContactInfo *CI_Contact_PropertyType `xml:"contactInfo,omitempty"`

	// generated from element "role" of complex type CI_RoleCode_PropertyType
	Role *CI_RoleCode_PropertyType `xml:"role,omitempty"`
}

// CI_ResponsibleParty_PropertyType Named source named complex type "CI_ResponsibleParty_PropertyType"
type CI_ResponsibleParty_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_ResponsibleParty" of complex type CI_ResponsibleParty_Type
	CI_ResponsibleParty *CI_ResponsibleParty_Type `xml:"CI_ResponsibleParty,omitempty"`
}

// CI_Citation_Type Named source named complex type "CI_Citation_Type"
// Standardized resource reference
type CI_Citation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "title" of complex type CharacterString_PropertyType
	Title *CharacterString_PropertyType `xml:"title,omitempty"`

	// generated from element "alternateTitle" of complex type CharacterString_PropertyType
	AlternateTitle []*CharacterString_PropertyType `xml:"alternateTitle,omitempty"`

	// generated from element "date" of complex type CI_Date_PropertyType
	Date []*CI_Date_PropertyType `xml:"date,omitempty"`

	// generated from element "edition" of complex type CharacterString_PropertyType
	Edition *CharacterString_PropertyType `xml:"edition,omitempty"`

	// generated from element "editionDate" of complex type Date_PropertyType
	EditionDate *Date_PropertyType `xml:"editionDate,omitempty"`

	// generated from element "identifier" of complex type MD_Identifier_PropertyType
	Identifier []*MD_Identifier_PropertyType `xml:"identifier,omitempty"`

	// generated from element "citedResponsibleParty" of complex type CI_ResponsibleParty_PropertyType
	CitedResponsibleParty []*CI_ResponsibleParty_PropertyType `xml:"citedResponsibleParty,omitempty"`

	// generated from element "presentationForm" of complex type CI_PresentationFormCode_PropertyType
	PresentationForm []*CI_PresentationFormCode_PropertyType `xml:"presentationForm,omitempty"`

	// generated from element "series" of complex type CI_Series_PropertyType
	Series *CI_Series_PropertyType `xml:"series,omitempty"`

	// generated from element "otherCitationDetails" of complex type CharacterString_PropertyType
	OtherCitationDetails *CharacterString_PropertyType `xml:"otherCitationDetails,omitempty"`

	// generated from element "collectiveTitle" of complex type CharacterString_PropertyType
	CollectiveTitle *CharacterString_PropertyType `xml:"collectiveTitle,omitempty"`

	// generated from element "ISBN" of complex type CharacterString_PropertyType
	ISBN *CharacterString_PropertyType `xml:"ISBN,omitempty"`

	// generated from element "ISSN" of complex type CharacterString_PropertyType
	ISSN *CharacterString_PropertyType `xml:"ISSN,omitempty"`
}

// CI_Citation_PropertyType Named source named complex type "CI_Citation_PropertyType"
type CI_Citation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Citation" of complex type CI_Citation_Type
	CI_Citation *CI_Citation_Type `xml:"CI_Citation,omitempty"`
}

// CI_Address_Type Named source named complex type "CI_Address_Type"
// Location of the responsible individual or organisation
type CI_Address_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "deliveryPoint" of complex type CharacterString_PropertyType
	DeliveryPoint []*CharacterString_PropertyType `xml:"deliveryPoint,omitempty"`

	// generated from element "city" of complex type CharacterString_PropertyType
	City *CharacterString_PropertyType `xml:"city,omitempty"`

	// generated from element "administrativeArea" of complex type CharacterString_PropertyType
	AdministrativeArea *CharacterString_PropertyType `xml:"administrativeArea,omitempty"`

	// generated from element "postalCode" of complex type CharacterString_PropertyType
	PostalCode *CharacterString_PropertyType `xml:"postalCode,omitempty"`

	// generated from element "country" of complex type CharacterString_PropertyType
	Country *CharacterString_PropertyType `xml:"country,omitempty"`

	// generated from element "electronicMailAddress" of complex type CharacterString_PropertyType
	ElectronicMailAddress []*CharacterString_PropertyType `xml:"electronicMailAddress,omitempty"`
}

// CI_Address_PropertyType Named source named complex type "CI_Address_PropertyType"
type CI_Address_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Address" of complex type CI_Address_Type
	CI_Address *CI_Address_Type `xml:"CI_Address,omitempty"`
}

// CI_OnlineResource_Type Named source named complex type "CI_OnlineResource_Type"
// Information about online sources from which the dataset, specification, or community profile name and extended metadata elements can be obtained.
type CI_OnlineResource_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "linkage" of complex type URL_PropertyType
	Linkage *URL_PropertyType `xml:"linkage,omitempty"`

	// generated from element "protocol" of complex type CharacterString_PropertyType
	Protocol *CharacterString_PropertyType `xml:"protocol,omitempty"`

	// generated from element "applicationProfile" of complex type CharacterString_PropertyType
	ApplicationProfile *CharacterString_PropertyType `xml:"applicationProfile,omitempty"`

	// generated from element "name" of complex type CharacterString_PropertyType
	NameXSD *CharacterString_PropertyType `xml:"name,omitempty"`

	// generated from element "description" of complex type CharacterString_PropertyType
	Description *CharacterString_PropertyType `xml:"description,omitempty"`

	// generated from element "function" of complex type CI_OnLineFunctionCode_PropertyType
	Function *CI_OnLineFunctionCode_PropertyType `xml:"function,omitempty"`
}

// CI_OnlineResource_PropertyType Named source named complex type "CI_OnlineResource_PropertyType"
type CI_OnlineResource_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_OnlineResource" of complex type CI_OnlineResource_Type
	CI_OnlineResource *CI_OnlineResource_Type `xml:"CI_OnlineResource,omitempty"`
}

// CI_Contact_Type Named source named complex type "CI_Contact_Type"
// Information required enabling contact with the  responsible person and/or organisation
type CI_Contact_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "phone" of complex type CI_Telephone_PropertyType
	Phone *CI_Telephone_PropertyType `xml:"phone,omitempty"`

	// generated from element "address" of complex type CI_Address_PropertyType
	Address *CI_Address_PropertyType `xml:"address,omitempty"`

	// generated from element "onlineResource" of complex type CI_OnlineResource_PropertyType
	OnlineResource *CI_OnlineResource_PropertyType `xml:"onlineResource,omitempty"`

	// generated from element "hoursOfService" of complex type CharacterString_PropertyType
	HoursOfService *CharacterString_PropertyType `xml:"hoursOfService,omitempty"`

	// generated from element "contactInstructions" of complex type CharacterString_PropertyType
	ContactInstructions *CharacterString_PropertyType `xml:"contactInstructions,omitempty"`
}

// CI_Contact_PropertyType Named source named complex type "CI_Contact_PropertyType"
type CI_Contact_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Contact" of complex type CI_Contact_Type
	CI_Contact *CI_Contact_Type `xml:"CI_Contact,omitempty"`
}

// CI_Telephone_Type Named source named complex type "CI_Telephone_Type"
// Telephone numbers for contacting the responsible individual or organisation
type CI_Telephone_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "voice" of complex type CharacterString_PropertyType
	Voice []*CharacterString_PropertyType `xml:"voice,omitempty"`

	// generated from element "facsimile" of complex type CharacterString_PropertyType
	Facsimile []*CharacterString_PropertyType `xml:"facsimile,omitempty"`
}

// CI_Telephone_PropertyType Named source named complex type "CI_Telephone_PropertyType"
type CI_Telephone_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Telephone" of complex type CI_Telephone_Type
	CI_Telephone *CI_Telephone_Type `xml:"CI_Telephone,omitempty"`
}

// CI_Date_Type Named source named complex type "CI_Date_Type"
type CI_Date_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "date" of complex type Date_PropertyType
	Date *Date_PropertyType `xml:"date,omitempty"`

	// generated from element "dateType" of complex type CI_DateTypeCode_PropertyType
	DateType *CI_DateTypeCode_PropertyType `xml:"dateType,omitempty"`
}

// CI_Date_PropertyType Named source named complex type "CI_Date_PropertyType"
type CI_Date_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Date" of complex type CI_Date_Type
	CI_Date *CI_Date_Type `xml:"CI_Date,omitempty"`
}

// CI_Series_Type Named source named complex type "CI_Series_Type"
type CI_Series_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type CharacterString_PropertyType
	NameXSD *CharacterString_PropertyType `xml:"name,omitempty"`

	// generated from element "issueIdentification" of complex type CharacterString_PropertyType
	IssueIdentification *CharacterString_PropertyType `xml:"issueIdentification,omitempty"`

	// generated from element "page" of complex type CharacterString_PropertyType
	Page *CharacterString_PropertyType `xml:"page,omitempty"`
}

// CI_Series_PropertyType Named source named complex type "CI_Series_PropertyType"
type CI_Series_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "CI_Series" of complex type CI_Series_Type
	CI_Series *CI_Series_Type `xml:"CI_Series,omitempty"`
}

// URL_PropertyType Named source named complex type "URL_PropertyType"
type URL_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "URL" of simple type xs:anyURI
	URL string `xml:"URL,omitempty"`
}

// CI_RoleCode_PropertyType Named source named complex type "CI_RoleCode_PropertyType"
type CI_RoleCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CI_RoleCode" of complex type CodeListValue_Type
	CI_RoleCode *gcobase.CodeListValue_Type `xml:"CI_RoleCode,omitempty"`
}

// CI_PresentationFormCode_PropertyType Named source named complex type "CI_PresentationFormCode_PropertyType"
type CI_PresentationFormCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CI_PresentationFormCode" of complex type CodeListValue_Type
	CI_PresentationFormCode *gcobase.CodeListValue_Type `xml:"CI_PresentationFormCode,omitempty"`
}

// CI_OnLineFunctionCode_PropertyType Named source named complex type "CI_OnLineFunctionCode_PropertyType"
type CI_OnLineFunctionCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CI_OnLineFunctionCode" of complex type CodeListValue_Type
	CI_OnLineFunctionCode *gcobase.CodeListValue_Type `xml:"CI_OnLineFunctionCode,omitempty"`
}

// CI_DateTypeCode_PropertyType Named source named complex type "CI_DateTypeCode_PropertyType"
type CI_DateTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CI_DateTypeCode" of complex type CodeListValue_Type
	CI_DateTypeCode *gcobase.CodeListValue_Type `xml:"CI_DateTypeCode,omitempty"`
}

// RS_Identifier_Type Named source named complex type "RS_Identifier_Type"
type RS_Identifier_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "authority" of complex type CI_Citation_PropertyType
	Authority *CI_Citation_PropertyType `xml:"authority,omitempty"`

	// generated from element "code" of complex type CharacterString_PropertyType
	Code *CharacterString_PropertyType `xml:"code,omitempty"`

	// generated from element "codeSpace" of complex type CharacterString_PropertyType
	CodeSpace *CharacterString_PropertyType `xml:"codeSpace,omitempty"`

	// generated from element "version" of complex type CharacterString_PropertyType
	Version *CharacterString_PropertyType `xml:"version,omitempty"`
}

// RS_Identifier_PropertyType Named source named complex type "RS_Identifier_PropertyType"
type RS_Identifier_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "RS_Identifier" of complex type RS_Identifier_Type
	RS_Identifier *RS_Identifier_Type `xml:"RS_Identifier,omitempty"`
}

// MD_ReferenceSystem_Type Named source named complex type "MD_ReferenceSystem_Type"
type MD_ReferenceSystem_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "referenceSystemIdentifier" of complex type RS_Identifier_PropertyType
	ReferenceSystemIdentifier *RS_Identifier_PropertyType `xml:"referenceSystemIdentifier,omitempty"`
}

// MD_ReferenceSystem_PropertyType Named source named complex type "MD_ReferenceSystem_PropertyType"
type MD_ReferenceSystem_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_ReferenceSystem" of complex type MD_ReferenceSystem_Type
	MD_ReferenceSystem *MD_ReferenceSystem_Type `xml:"MD_ReferenceSystem,omitempty"`
}

// MD_Identifier_Type Named source named complex type "MD_Identifier_Type"
type MD_Identifier_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "authority" of complex type CI_Citation_PropertyType
	Authority *CI_Citation_PropertyType `xml:"authority,omitempty"`

	// generated from element "code" of complex type CharacterString_PropertyType
	Code *CharacterString_PropertyType `xml:"code,omitempty"`
}

// MD_Identifier_PropertyType Named source named complex type "MD_Identifier_PropertyType"
type MD_Identifier_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Identifier" of complex type MD_Identifier_Type
	MD_Identifier *MD_Identifier_Type `xml:"MD_Identifier,omitempty"`

	// generated from element "RS_Identifier" of complex type RS_Identifier_Type
	RS_Identifier *RS_Identifier_Type `xml:"RS_Identifier,omitempty"`
}

// AbstractRS_ReferenceSystem_Type Named source named complex type "AbstractRS_ReferenceSystem_Type"
// Description of the spatial and temporal reference systems used in the dataset
type AbstractRS_ReferenceSystem_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type RS_Identifier_PropertyType
	NameXSD *RS_Identifier_PropertyType `xml:"name,omitempty"`

	// generated from element "domainOfValidity" of complex type EX_Extent_PropertyType
	DomainOfValidity []*EX_Extent_PropertyType `xml:"domainOfValidity,omitempty"`
}

// RS_ReferenceSystem_PropertyType Named source named complex type "RS_ReferenceSystem_PropertyType"
type RS_ReferenceSystem_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "AbstractRS_ReferenceSystem" of complex type AbstractRS_ReferenceSystem_Type
	AbstractRS_ReferenceSystem *AbstractRS_ReferenceSystem_Type `xml:"AbstractRS_ReferenceSystem,omitempty"`
}

// EX_TemporalExtent_Type Named source named complex type "EX_TemporalExtent_Type"
// Time period covered by the content of the dataset
type EX_TemporalExtent_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extent" of complex type TM_Primitive_PropertyType
	Extent *TM_Primitive_PropertyType `xml:"extent,omitempty"`
}

// EX_TemporalExtent_PropertyType Named source named complex type "EX_TemporalExtent_PropertyType"
type EX_TemporalExtent_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_TemporalExtent" of complex type EX_TemporalExtent_Type
	EX_TemporalExtent *EX_TemporalExtent_Type `xml:"EX_TemporalExtent,omitempty"`

	// generated from element "EX_SpatialTemporalExtent" of complex type EX_SpatialTemporalExtent_Type
	EX_SpatialTemporalExtent *EX_SpatialTemporalExtent_Type `xml:"EX_SpatialTemporalExtent,omitempty"`
}

// EX_VerticalExtent_Type Named source named complex type "EX_VerticalExtent_Type"
// Vertical domain of dataset
type EX_VerticalExtent_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "minimumValue" of complex type Real_PropertyType
	MinimumValue *Real_PropertyType `xml:"minimumValue,omitempty"`

	// generated from element "maximumValue" of complex type Real_PropertyType
	MaximumValue *Real_PropertyType `xml:"maximumValue,omitempty"`

	// generated from element "verticalCRS" of complex type SC_CRS_PropertyType
	VerticalCRS *SC_CRS_PropertyType `xml:"verticalCRS,omitempty"`
}

// EX_VerticalExtent_PropertyType Named source named complex type "EX_VerticalExtent_PropertyType"
type EX_VerticalExtent_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_VerticalExtent" of complex type EX_VerticalExtent_Type
	EX_VerticalExtent *EX_VerticalExtent_Type `xml:"EX_VerticalExtent,omitempty"`
}

// EX_BoundingPolygon_Type Named source named complex type "EX_BoundingPolygon_Type"
// Boundary enclosing the dataset expressed as the closed set of (x,y) coordinates of the polygon (last point replicates first point)
type EX_BoundingPolygon_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extentTypeCode" of complex type Boolean_PropertyType
	ExtentTypeCode *Boolean_PropertyType `xml:"extentTypeCode,omitempty"`

	// generated from element "polygon" of complex type GM_Object_PropertyType
	Polygon []*GM_Object_PropertyType `xml:"polygon,omitempty"`
}

// EX_BoundingPolygon_PropertyType Named source named complex type "EX_BoundingPolygon_PropertyType"
type EX_BoundingPolygon_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_BoundingPolygon" of complex type EX_BoundingPolygon_Type
	EX_BoundingPolygon *EX_BoundingPolygon_Type `xml:"EX_BoundingPolygon,omitempty"`
}

// EX_Extent_Type Named source named complex type "EX_Extent_Type"
// Information about spatial, vertical, and temporal extent
type EX_Extent_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "description" of complex type CharacterString_PropertyType
	Description *CharacterString_PropertyType `xml:"description,omitempty"`

	// generated from element "geographicElement" of complex type EX_GeographicExtent_PropertyType
	GeographicElement []*EX_GeographicExtent_PropertyType `xml:"geographicElement,omitempty"`

	// generated from element "temporalElement" of complex type EX_TemporalExtent_PropertyType
	TemporalElement []*EX_TemporalExtent_PropertyType `xml:"temporalElement,omitempty"`

	// generated from element "verticalElement" of complex type EX_VerticalExtent_PropertyType
	VerticalElement []*EX_VerticalExtent_PropertyType `xml:"verticalElement,omitempty"`
}

// EX_Extent_PropertyType Named source named complex type "EX_Extent_PropertyType"
type EX_Extent_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_Extent" of complex type EX_Extent_Type
	EX_Extent *EX_Extent_Type `xml:"EX_Extent,omitempty"`
}

// AbstractEX_GeographicExtent_Type Named source named complex type "AbstractEX_GeographicExtent_Type"
// Geographic area of the dataset
type AbstractEX_GeographicExtent_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extentTypeCode" of complex type Boolean_PropertyType
	ExtentTypeCode *Boolean_PropertyType `xml:"extentTypeCode,omitempty"`
}

// EX_GeographicExtent_PropertyType Named source named complex type "EX_GeographicExtent_PropertyType"
type EX_GeographicExtent_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_BoundingPolygon" of complex type EX_BoundingPolygon_Type
	EX_BoundingPolygon *EX_BoundingPolygon_Type `xml:"EX_BoundingPolygon,omitempty"`

	// generated from element "EX_GeographicBoundingBox" of complex type EX_GeographicBoundingBox_Type
	EX_GeographicBoundingBox *EX_GeographicBoundingBox_Type `xml:"EX_GeographicBoundingBox,omitempty"`

	// generated from element "EX_GeographicDescription" of complex type EX_GeographicDescription_Type
	EX_GeographicDescription *EX_GeographicDescription_Type `xml:"EX_GeographicDescription,omitempty"`
}

// EX_GeographicBoundingBox_Type Named source named complex type "EX_GeographicBoundingBox_Type"
// Geographic area of the entire dataset referenced to WGS 84
type EX_GeographicBoundingBox_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extentTypeCode" of complex type Boolean_PropertyType
	ExtentTypeCode *Boolean_PropertyType `xml:"extentTypeCode,omitempty"`

	// generated from element "westBoundLongitude" of complex type Decimal_PropertyType
	WestBoundLongitude *Decimal_PropertyType `xml:"westBoundLongitude,omitempty"`

	// generated from element "eastBoundLongitude" of complex type Decimal_PropertyType
	EastBoundLongitude *Decimal_PropertyType `xml:"eastBoundLongitude,omitempty"`

	// generated from element "southBoundLatitude" of complex type Decimal_PropertyType
	SouthBoundLatitude *Decimal_PropertyType `xml:"southBoundLatitude,omitempty"`

	// generated from element "northBoundLatitude" of complex type Decimal_PropertyType
	NorthBoundLatitude *Decimal_PropertyType `xml:"northBoundLatitude,omitempty"`
}

// EX_GeographicBoundingBox_PropertyType Named source named complex type "EX_GeographicBoundingBox_PropertyType"
type EX_GeographicBoundingBox_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_GeographicBoundingBox" of complex type EX_GeographicBoundingBox_Type
	EX_GeographicBoundingBox *EX_GeographicBoundingBox_Type `xml:"EX_GeographicBoundingBox,omitempty"`
}

// EX_SpatialTemporalExtent_Type Named source named complex type "EX_SpatialTemporalExtent_Type"
// Extent with respect to date and time
type EX_SpatialTemporalExtent_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extent" of complex type TM_Primitive_PropertyType
	Extent *TM_Primitive_PropertyType `xml:"extent,omitempty"`

	// generated from element "spatialExtent" of complex type EX_GeographicExtent_PropertyType
	SpatialExtent []*EX_GeographicExtent_PropertyType `xml:"spatialExtent,omitempty"`
}

// EX_SpatialTemporalExtent_PropertyType Named source named complex type "EX_SpatialTemporalExtent_PropertyType"
type EX_SpatialTemporalExtent_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_SpatialTemporalExtent" of complex type EX_SpatialTemporalExtent_Type
	EX_SpatialTemporalExtent *EX_SpatialTemporalExtent_Type `xml:"EX_SpatialTemporalExtent,omitempty"`
}

// EX_GeographicDescription_Type Named source named complex type "EX_GeographicDescription_Type"
type EX_GeographicDescription_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extentTypeCode" of complex type Boolean_PropertyType
	ExtentTypeCode *Boolean_PropertyType `xml:"extentTypeCode,omitempty"`

	// generated from element "geographicIdentifier" of complex type MD_Identifier_PropertyType
	GeographicIdentifier *MD_Identifier_PropertyType `xml:"geographicIdentifier,omitempty"`
}

// EX_GeographicDescription_PropertyType Named source named complex type "EX_GeographicDescription_PropertyType"
type EX_GeographicDescription_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "EX_GeographicDescription" of complex type EX_GeographicDescription_Type
	EX_GeographicDescription *EX_GeographicDescription_Type `xml:"EX_GeographicDescription,omitempty"`
}

// GM_Point_PropertyType Named source named complex type "GM_Point_PropertyType"
type GM_Point_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`
}

// GM_Object_PropertyType Named source named complex type "GM_Object_PropertyType"
type GM_Object_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`
}

// TypeName_Type Named source named complex type "TypeName_Type"
// A TypeName is a LocalName that references either a recordType or object type in some form of schema. The stored value "aName" is the returned value for the "aName()" operation. This is the types name.  - For parsing from types (or objects) the parsible name normally uses a "." navigation separator, so that it is of the form  [class].[member].[memberOfMember]. ...)
type TypeName_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "aName" of complex type CharacterString_PropertyType
	AName *CharacterString_PropertyType `xml:"aName,omitempty"`
}

// TypeName_PropertyType Named source named complex type "TypeName_PropertyType"
type TypeName_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "TypeName" of complex type TypeName_Type
	TypeName *TypeName_Type `xml:"TypeName,omitempty"`
}

// MemberName_Type Named source named complex type "MemberName_Type"
// A MemberName is a LocalName that references either an attribute slot in a record or  recordType or an attribute, operation, or association role in an object instance or  type description in some form of schema. The stored value "aName" is the returned value for the "aName()" operation.
type MemberName_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "aName" of complex type CharacterString_PropertyType
	AName *CharacterString_PropertyType `xml:"aName,omitempty"`

	// generated from element "attributeType" of complex type TypeName_PropertyType
	AttributeType *TypeName_PropertyType `xml:"attributeType,omitempty"`
}

// MemberName_PropertyType Named source named complex type "MemberName_PropertyType"
type MemberName_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MemberName" of complex type MemberName_Type
	MemberName *MemberName_Type `xml:"MemberName,omitempty"`
}

// Multiplicity_Type Named source named complex type "Multiplicity_Type"
// Use to represent the possible cardinality of a relation. Represented by a set of simple multiplicity ranges.
type Multiplicity_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "range" of complex type MultiplicityRange_PropertyType
	Range_XSD []*MultiplicityRange_PropertyType `xml:"range,omitempty"`
}

// Multiplicity_PropertyType Named source named complex type "Multiplicity_PropertyType"
type Multiplicity_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "Multiplicity" of complex type Multiplicity_Type
	Multiplicity *Multiplicity_Type `xml:"Multiplicity,omitempty"`
}

// MultiplicityRange_Type Named source named complex type "MultiplicityRange_Type"
// A component of a multiplicity, consisting of an non-negative lower bound, and a potentially infinite upper bound.
type MultiplicityRange_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "lower" of complex type Integer_PropertyType
	Lower *Integer_PropertyType `xml:"lower,omitempty"`

	// generated from element "upper" of complex type UnlimitedInteger_PropertyType
	Upper *UnlimitedInteger_PropertyType `xml:"upper,omitempty"`
}

// MultiplicityRange_PropertyType Named source named complex type "MultiplicityRange_PropertyType"
type MultiplicityRange_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MultiplicityRange" of complex type MultiplicityRange_Type
	MultiplicityRange *MultiplicityRange_Type `xml:"MultiplicityRange,omitempty"`
}

// Measure_PropertyType Named source named complex type "Measure_PropertyType"
type Measure_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Measure" of complex type MeasureType
	Measure *basictypes.MeasureType `xml:"Measure,omitempty"`

	// generated from element "Length" of complex type LengthType
	Length *measures.LengthType `xml:"Length,omitempty"`

	// generated from element "Distance" of complex type LengthType
	Distance *measures.LengthType `xml:"Distance,omitempty"`

	// generated from element "Angle" of complex type AngleType
	Angle *measures.AngleType `xml:"Angle,omitempty"`

	// generated from element "Scale" of complex type ScaleType
	Scale *measures.ScaleType `xml:"Scale,omitempty"`
}

// Length_PropertyType Named source named complex type "Length_PropertyType"
type Length_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Length" of complex type LengthType
	Length *measures.LengthType `xml:"Length,omitempty"`

	// generated from element "Distance" of complex type LengthType
	Distance *measures.LengthType `xml:"Distance,omitempty"`
}

// Angle_PropertyType Named source named complex type "Angle_PropertyType"
type Angle_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Angle" of complex type AngleType
	Angle *measures.AngleType `xml:"Angle,omitempty"`
}

// Scale_PropertyType Named source named complex type "Scale_PropertyType"
type Scale_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Scale" of complex type ScaleType
	Scale *measures.ScaleType `xml:"Scale,omitempty"`
}

// Distance_PropertyType Named source named complex type "Distance_PropertyType"
type Distance_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Distance" of complex type LengthType
	Distance *measures.LengthType `xml:"Distance,omitempty"`
}

// CharacterString_PropertyType Named source named complex type "CharacterString_PropertyType"
type CharacterString_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CharacterString" of simple type xs:string
	CharacterString string `xml:"CharacterString,omitempty"`

	// generated from element "MD_PixelOrientationCode" of simple type gmd:MD_PixelOrientationCode_Type
	MD_PixelOrientationCode string `xml:"MD_PixelOrientationCode,omitempty"`

	// generated from element "MD_TopologyLevelCode" of complex type CodeListValue_Type
	MD_TopologyLevelCode *gcobase.CodeListValue_Type `xml:"MD_TopologyLevelCode,omitempty"`

	// generated from element "MD_GeometricObjectTypeCode" of complex type CodeListValue_Type
	MD_GeometricObjectTypeCode *gcobase.CodeListValue_Type `xml:"MD_GeometricObjectTypeCode,omitempty"`

	// generated from element "MD_CellGeometryCode" of complex type CodeListValue_Type
	MD_CellGeometryCode *gcobase.CodeListValue_Type `xml:"MD_CellGeometryCode,omitempty"`

	// generated from element "MD_DimensionNameTypeCode" of complex type CodeListValue_Type
	MD_DimensionNameTypeCode *gcobase.CodeListValue_Type `xml:"MD_DimensionNameTypeCode,omitempty"`

	// generated from element "CI_RoleCode" of complex type CodeListValue_Type
	CI_RoleCode *gcobase.CodeListValue_Type `xml:"CI_RoleCode,omitempty"`

	// generated from element "CI_PresentationFormCode" of complex type CodeListValue_Type
	CI_PresentationFormCode *gcobase.CodeListValue_Type `xml:"CI_PresentationFormCode,omitempty"`

	// generated from element "CI_OnLineFunctionCode" of complex type CodeListValue_Type
	CI_OnLineFunctionCode *gcobase.CodeListValue_Type `xml:"CI_OnLineFunctionCode,omitempty"`

	// generated from element "CI_DateTypeCode" of complex type CodeListValue_Type
	CI_DateTypeCode *gcobase.CodeListValue_Type `xml:"CI_DateTypeCode,omitempty"`

	// generated from element "MD_ObligationCode" of simple type gmd:MD_ObligationCode_Type
	MD_ObligationCode string `xml:"MD_ObligationCode,omitempty"`

	// generated from element "MD_DatatypeCode" of complex type CodeListValue_Type
	MD_DatatypeCode *gcobase.CodeListValue_Type `xml:"MD_DatatypeCode,omitempty"`

	// generated from element "MD_CoverageContentTypeCode" of complex type CodeListValue_Type
	MD_CoverageContentTypeCode *gcobase.CodeListValue_Type `xml:"MD_CoverageContentTypeCode,omitempty"`

	// generated from element "MD_ImagingConditionCode" of complex type CodeListValue_Type
	MD_ImagingConditionCode *gcobase.CodeListValue_Type `xml:"MD_ImagingConditionCode,omitempty"`

	// generated from element "DQ_EvaluationMethodTypeCode" of complex type CodeListValue_Type
	DQ_EvaluationMethodTypeCode *gcobase.CodeListValue_Type `xml:"DQ_EvaluationMethodTypeCode,omitempty"`

	// generated from element "MD_TopicCategoryCode" of simple type gmd:MD_TopicCategoryCode_Type
	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type `xml:"MD_TopicCategoryCode,omitempty"`

	// generated from element "MD_CharacterSetCode" of complex type CodeListValue_Type
	MD_CharacterSetCode *gcobase.CodeListValue_Type `xml:"MD_CharacterSetCode,omitempty"`

	// generated from element "MD_SpatialRepresentationTypeCode" of complex type CodeListValue_Type
	MD_SpatialRepresentationTypeCode *gcobase.CodeListValue_Type `xml:"MD_SpatialRepresentationTypeCode,omitempty"`

	// generated from element "MD_ProgressCode" of complex type CodeListValue_Type
	MD_ProgressCode *gcobase.CodeListValue_Type `xml:"MD_ProgressCode,omitempty"`

	// generated from element "MD_KeywordTypeCode" of complex type CodeListValue_Type
	MD_KeywordTypeCode *gcobase.CodeListValue_Type `xml:"MD_KeywordTypeCode,omitempty"`

	// generated from element "DS_AssociationTypeCode" of complex type CodeListValue_Type
	DS_AssociationTypeCode *gcobase.CodeListValue_Type `xml:"DS_AssociationTypeCode,omitempty"`

	// generated from element "DS_InitiativeTypeCode" of complex type CodeListValue_Type
	DS_InitiativeTypeCode *gcobase.CodeListValue_Type `xml:"DS_InitiativeTypeCode,omitempty"`

	// generated from element "MD_ClassificationCode" of complex type CodeListValue_Type
	MD_ClassificationCode *gcobase.CodeListValue_Type `xml:"MD_ClassificationCode,omitempty"`

	// generated from element "MD_RestrictionCode" of complex type CodeListValue_Type
	MD_RestrictionCode *gcobase.CodeListValue_Type `xml:"MD_RestrictionCode,omitempty"`

	// generated from element "MD_DistributionUnits" of complex type CodeListValue_Type
	MD_DistributionUnits *gcobase.CodeListValue_Type `xml:"MD_DistributionUnits,omitempty"`

	// generated from element "MD_MediumFormatCode" of complex type CodeListValue_Type
	MD_MediumFormatCode *gcobase.CodeListValue_Type `xml:"MD_MediumFormatCode,omitempty"`

	// generated from element "MD_MediumNameCode" of complex type CodeListValue_Type
	MD_MediumNameCode *gcobase.CodeListValue_Type `xml:"MD_MediumNameCode,omitempty"`

	// generated from element "MD_MaintenanceFrequencyCode" of complex type CodeListValue_Type
	MD_MaintenanceFrequencyCode *gcobase.CodeListValue_Type `xml:"MD_MaintenanceFrequencyCode,omitempty"`

	// generated from element "MD_ScopeCode" of complex type CodeListValue_Type
	MD_ScopeCode *gcobase.CodeListValue_Type `xml:"MD_ScopeCode,omitempty"`

	// generated from element "LocalisedCharacterString" of complex type LocalisedCharacterString_Type
	LocalisedCharacterString *LocalisedCharacterString_Type `xml:"LocalisedCharacterString,omitempty"`

	// generated from element "LanguageCode" of complex type CodeListValue_Type
	LanguageCode *gcobase.CodeListValue_Type `xml:"LanguageCode,omitempty"`

	// generated from element "Country" of complex type CodeListValue_Type
	Country *gcobase.CodeListValue_Type `xml:"Country,omitempty"`
}

// Boolean_PropertyType Named source named complex type "Boolean_PropertyType"
type Boolean_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Boolean" of simple type xs:boolean
	Boolean bool `xml:"Boolean,omitempty"`
}

// GenericName_PropertyType Named source named complex type "GenericName_PropertyType"
type GenericName_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "LocalName" of complex type CodeType
	LocalName *basictypes.CodeType `xml:"LocalName,omitempty"`

	// generated from element "ScopedName" of complex type CodeType
	ScopedName *basictypes.CodeType `xml:"ScopedName,omitempty"`
}

// LocalName_PropertyType Named source named complex type "LocalName_PropertyType"
type LocalName_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "LocalName" of complex type CodeType
	LocalName *basictypes.CodeType `xml:"LocalName,omitempty"`
}

// ScopedName_PropertyType Named source named complex type "ScopedName_PropertyType"
type ScopedName_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "ScopedName" of complex type CodeType
	ScopedName *basictypes.CodeType `xml:"ScopedName,omitempty"`
}

// UomAngle_PropertyType Named source named complex type "UomAngle_PropertyType"
type UomAngle_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomLength_PropertyType Named source named complex type "UomLength_PropertyType"
type UomLength_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomScale_PropertyType Named source named complex type "UomScale_PropertyType"
type UomScale_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UnitOfMeasure_PropertyType Named source named complex type "UnitOfMeasure_PropertyType"
type UnitOfMeasure_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomArea_PropertyType Named source named complex type "UomArea_PropertyType"
type UomArea_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomVelocity_PropertyType Named source named complex type "UomVelocity_PropertyType"
type UomVelocity_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomTime_PropertyType Named source named complex type "UomTime_PropertyType"
type UomTime_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// UomVolume_PropertyType Named source named complex type "UomVolume_PropertyType"
type UomVolume_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition *UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit *BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit *DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit *ConventionalUnitType `xml:"ConventionalUnit,omitempty"`
}

// DateTime_PropertyType Named source named complex type "DateTime_PropertyType"
type DateTime_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "DateTime" of simple type xs:dateTime
	DateTime string `xml:"DateTime,omitempty"`
}

// Date_PropertyType Named source named complex type "Date_PropertyType"
type Date_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Date" of simple type gco:Date_Type
	Date string `xml:"Date,omitempty"`

	// generated from element "DateTime" of simple type xs:dateTime
	DateTime string `xml:"DateTime,omitempty"`
}

// Number_PropertyType Named source named complex type "Number_PropertyType"
type Number_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Real" of simple type xs:double
	Real_XSD float64 `xml:"Real,omitempty"`

	// generated from element "Decimal" of simple type xs:decimal
	Decimal string `xml:"Decimal,omitempty"`

	// generated from element "Integer" of simple type xs:integer
	Integer int `xml:"Integer,omitempty"`
}

// Decimal_PropertyType Named source named complex type "Decimal_PropertyType"
type Decimal_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Decimal" of simple type xs:decimal
	Decimal string `xml:"Decimal,omitempty"`
}

// Real_PropertyType Named source named complex type "Real_PropertyType"
type Real_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Real" of simple type xs:double
	Real_XSD float64 `xml:"Real,omitempty"`
}

// Integer_PropertyType Named source named complex type "Integer_PropertyType"
type Integer_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Integer" of simple type xs:integer
	Integer int `xml:"Integer,omitempty"`
}

// UnlimitedInteger_Type Named source named complex type "UnlimitedInteger_Type"
type UnlimitedInteger_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "isInfinite"
	IsInfinite bool `xml:"isInfinite,attr,omitempty"`

	// simpleContent character data
	EnclosedText int `xml:",chardata"`
}

// UnlimitedInteger_PropertyType Named source named complex type "UnlimitedInteger_PropertyType"
type UnlimitedInteger_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "UnlimitedInteger" of complex type UnlimitedInteger_Type
	UnlimitedInteger *UnlimitedInteger_Type `xml:"UnlimitedInteger,omitempty"`
}

// Record_PropertyType Named source named complex type "Record_PropertyType"
type Record_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from untyped element "Record"
	Record string `xml:"Record,omitempty"`
}

// RecordType_Type Named source named complex type "RecordType_Type"
type RecordType_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:simpleAttrs"
	xlink.AttributeGroup_simpleAttrs

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// RecordType_PropertyType Named source named complex type "RecordType_PropertyType"
type RecordType_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "RecordType" of complex type RecordType_Type
	RecordType *RecordType_Type `xml:"RecordType,omitempty"`
}

// Binary_Type Named source named complex type "Binary_Type"
type Binary_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "src"
	Src string `xml:"src,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Binary_PropertyType Named source named complex type "Binary_PropertyType"
type Binary_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Binary" of complex type Binary_Type
	Binary *Binary_Type `xml:"Binary,omitempty"`
}

// TM_Primitive_PropertyType Named source named complex type "TM_Primitive_PropertyType"
type TM_Primitive_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`
}

// TM_PeriodDuration_PropertyType Named source named complex type "TM_PeriodDuration_PropertyType"
type TM_PeriodDuration_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "TM_PeriodDuration" of simple type xs:duration
	TM_PeriodDuration string `xml:"TM_PeriodDuration,omitempty"`
}

// SC_CRS_PropertyType Named source named complex type "SC_CRS_PropertyType"
type SC_CRS_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS *ProjectedCRSType `xml:"ProjectedCRS,omitempty"`

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS *DerivedCRSType `xml:"DerivedCRS,omitempty"`

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS *GeodeticCRSType `xml:"GeodeticCRS,omitempty"`

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS *VerticalCRSType `xml:"VerticalCRS,omitempty"`

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS *EngineeringCRSType `xml:"EngineeringCRS,omitempty"`

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS *ImageCRSType `xml:"ImageCRS,omitempty"`

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS *TemporalCRSType `xml:"TemporalCRS,omitempty"`

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS *GeographicCRSType `xml:"GeographicCRS,omitempty"`

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS *GeocentricCRSType `xml:"GeocentricCRS,omitempty"`

	// generated from element "CompoundCRS" of complex type CompoundCRSType
	CompoundCRS *CompoundCRSType `xml:"CompoundCRS,omitempty"`
}

// LI_ProcessStep_Type Named source named complex type "LI_ProcessStep_Type"
type LI_ProcessStep_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "description" of complex type CharacterString_PropertyType
	Description *CharacterString_PropertyType `xml:"description,omitempty"`

	// generated from element "rationale" of complex type CharacterString_PropertyType
	Rationale *CharacterString_PropertyType `xml:"rationale,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime *DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "processor" of complex type CI_ResponsibleParty_PropertyType
	Processor []*CI_ResponsibleParty_PropertyType `xml:"processor,omitempty"`

	// generated from element "source" of complex type LI_Source_PropertyType
	Source []*LI_Source_PropertyType `xml:"source,omitempty"`
}

// LI_ProcessStep_PropertyType Named source named complex type "LI_ProcessStep_PropertyType"
type LI_ProcessStep_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "LI_ProcessStep" of complex type LI_ProcessStep_Type
	LI_ProcessStep *LI_ProcessStep_Type `xml:"LI_ProcessStep,omitempty"`
}

// LI_Source_Type Named source named complex type "LI_Source_Type"
type LI_Source_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "description" of complex type CharacterString_PropertyType
	Description *CharacterString_PropertyType `xml:"description,omitempty"`

	// generated from element "scaleDenominator" of complex type MD_RepresentativeFraction_PropertyType
	ScaleDenominator *MD_RepresentativeFraction_PropertyType `xml:"scaleDenominator,omitempty"`

	// generated from element "sourceReferenceSystem" of complex type MD_ReferenceSystem_PropertyType
	SourceReferenceSystem *MD_ReferenceSystem_PropertyType `xml:"sourceReferenceSystem,omitempty"`

	// generated from element "sourceCitation" of complex type CI_Citation_PropertyType
	SourceCitation *CI_Citation_PropertyType `xml:"sourceCitation,omitempty"`

	// generated from element "sourceExtent" of complex type EX_Extent_PropertyType
	SourceExtent []*EX_Extent_PropertyType `xml:"sourceExtent,omitempty"`

	// generated from element "sourceStep" of complex type LI_ProcessStep_PropertyType
	SourceStep []*LI_ProcessStep_PropertyType `xml:"sourceStep,omitempty"`
}

// LI_Source_PropertyType Named source named complex type "LI_Source_PropertyType"
type LI_Source_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "LI_Source" of complex type LI_Source_Type
	LI_Source *LI_Source_Type `xml:"LI_Source,omitempty"`
}

// LI_Lineage_Type Named source named complex type "LI_Lineage_Type"
type LI_Lineage_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "statement" of complex type CharacterString_PropertyType
	Statement *CharacterString_PropertyType `xml:"statement,omitempty"`

	// generated from element "processStep" of complex type LI_ProcessStep_PropertyType
	ProcessStep []*LI_ProcessStep_PropertyType `xml:"processStep,omitempty"`

	// generated from element "source" of complex type LI_Source_PropertyType
	Source []*LI_Source_PropertyType `xml:"source,omitempty"`
}

// LI_Lineage_PropertyType Named source named complex type "LI_Lineage_PropertyType"
type LI_Lineage_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "LI_Lineage" of complex type LI_Lineage_Type
	LI_Lineage *LI_Lineage_Type `xml:"LI_Lineage,omitempty"`
}

// DQ_ConformanceResult_Type Named source named complex type "DQ_ConformanceResult_Type"
// quantitative_result from Quality Procedures -  - renamed to remove implied use limitiation.
type DQ_ConformanceResult_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "specification" of complex type CI_Citation_PropertyType
	Specification *CI_Citation_PropertyType `xml:"specification,omitempty"`

	// generated from element "explanation" of complex type CharacterString_PropertyType
	Explanation *CharacterString_PropertyType `xml:"explanation,omitempty"`

	// generated from element "pass" of complex type Boolean_PropertyType
	Pass *Boolean_PropertyType `xml:"pass,omitempty"`
}

// DQ_ConformanceResult_PropertyType Named source named complex type "DQ_ConformanceResult_PropertyType"
type DQ_ConformanceResult_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_ConformanceResult" of complex type DQ_ConformanceResult_Type
	DQ_ConformanceResult *DQ_ConformanceResult_Type `xml:"DQ_ConformanceResult,omitempty"`
}

// DQ_QuantitativeResult_Type Named source named complex type "DQ_QuantitativeResult_Type"
// Quantitative_conformance_measure from Quality Procedures.  -  - Renamed to remove implied use limitation -  - OCL - -- result is type specified by valueDomain - result.tupleType = valueDomain
type DQ_QuantitativeResult_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "valueType" of complex type RecordType_PropertyType
	ValueType *RecordType_PropertyType `xml:"valueType,omitempty"`

	// generated from element "valueUnit" of complex type UnitOfMeasure_PropertyType
	ValueUnit *UnitOfMeasure_PropertyType `xml:"valueUnit,omitempty"`

	// generated from element "errorStatistic" of complex type CharacterString_PropertyType
	ErrorStatistic *CharacterString_PropertyType `xml:"errorStatistic,omitempty"`

	// generated from element "value" of complex type Record_PropertyType
	Value []*Record_PropertyType `xml:"value,omitempty"`
}

// DQ_QuantitativeResult_PropertyType Named source named complex type "DQ_QuantitativeResult_PropertyType"
type DQ_QuantitativeResult_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_QuantitativeResult" of complex type DQ_QuantitativeResult_Type
	DQ_QuantitativeResult *DQ_QuantitativeResult_Type `xml:"DQ_QuantitativeResult,omitempty"`
}

// AbstractDQ_Result_Type Named source named complex type "AbstractDQ_Result_Type"
type AbstractDQ_Result_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification
}

// DQ_Result_PropertyType Named source named complex type "DQ_Result_PropertyType"
type DQ_Result_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_ConformanceResult" of complex type DQ_ConformanceResult_Type
	DQ_ConformanceResult *DQ_ConformanceResult_Type `xml:"DQ_ConformanceResult,omitempty"`

	// generated from element "DQ_QuantitativeResult" of complex type DQ_QuantitativeResult_Type
	DQ_QuantitativeResult *DQ_QuantitativeResult_Type `xml:"DQ_QuantitativeResult,omitempty"`
}

// DQ_TemporalValidity_Type Named source named complex type "DQ_TemporalValidity_Type"
type DQ_TemporalValidity_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_TemporalValidity_PropertyType Named source named complex type "DQ_TemporalValidity_PropertyType"
type DQ_TemporalValidity_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TemporalValidity" of complex type DQ_TemporalValidity_Type
	DQ_TemporalValidity *DQ_TemporalValidity_Type `xml:"DQ_TemporalValidity,omitempty"`
}

// DQ_TemporalConsistency_Type Named source named complex type "DQ_TemporalConsistency_Type"
type DQ_TemporalConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_TemporalConsistency_PropertyType Named source named complex type "DQ_TemporalConsistency_PropertyType"
type DQ_TemporalConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TemporalConsistency" of complex type DQ_TemporalConsistency_Type
	DQ_TemporalConsistency *DQ_TemporalConsistency_Type `xml:"DQ_TemporalConsistency,omitempty"`
}

// DQ_AccuracyOfATimeMeasurement_Type Named source named complex type "DQ_AccuracyOfATimeMeasurement_Type"
type DQ_AccuracyOfATimeMeasurement_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_AccuracyOfATimeMeasurement_PropertyType Named source named complex type "DQ_AccuracyOfATimeMeasurement_PropertyType"
type DQ_AccuracyOfATimeMeasurement_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_AccuracyOfATimeMeasurement" of complex type DQ_AccuracyOfATimeMeasurement_Type
	DQ_AccuracyOfATimeMeasurement *DQ_AccuracyOfATimeMeasurement_Type `xml:"DQ_AccuracyOfATimeMeasurement,omitempty"`
}

// DQ_QuantitativeAttributeAccuracy_Type Named source named complex type "DQ_QuantitativeAttributeAccuracy_Type"
type DQ_QuantitativeAttributeAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_QuantitativeAttributeAccuracy_PropertyType Named source named complex type "DQ_QuantitativeAttributeAccuracy_PropertyType"
type DQ_QuantitativeAttributeAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_QuantitativeAttributeAccuracy" of complex type DQ_QuantitativeAttributeAccuracy_Type
	DQ_QuantitativeAttributeAccuracy *DQ_QuantitativeAttributeAccuracy_Type `xml:"DQ_QuantitativeAttributeAccuracy,omitempty"`
}

// DQ_NonQuantitativeAttributeAccuracy_Type Named source named complex type "DQ_NonQuantitativeAttributeAccuracy_Type"
type DQ_NonQuantitativeAttributeAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_NonQuantitativeAttributeAccuracy_PropertyType Named source named complex type "DQ_NonQuantitativeAttributeAccuracy_PropertyType"
type DQ_NonQuantitativeAttributeAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_NonQuantitativeAttributeAccuracy" of complex type DQ_NonQuantitativeAttributeAccuracy_Type
	DQ_NonQuantitativeAttributeAccuracy *DQ_NonQuantitativeAttributeAccuracy_Type `xml:"DQ_NonQuantitativeAttributeAccuracy,omitempty"`
}

// DQ_ThematicClassificationCorrectness_Type Named source named complex type "DQ_ThematicClassificationCorrectness_Type"
type DQ_ThematicClassificationCorrectness_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_ThematicClassificationCorrectness_PropertyType Named source named complex type "DQ_ThematicClassificationCorrectness_PropertyType"
type DQ_ThematicClassificationCorrectness_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_ThematicClassificationCorrectness" of complex type DQ_ThematicClassificationCorrectness_Type
	DQ_ThematicClassificationCorrectness *DQ_ThematicClassificationCorrectness_Type `xml:"DQ_ThematicClassificationCorrectness,omitempty"`
}

// DQ_RelativeInternalPositionalAccuracy_Type Named source named complex type "DQ_RelativeInternalPositionalAccuracy_Type"
type DQ_RelativeInternalPositionalAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_RelativeInternalPositionalAccuracy_PropertyType Named source named complex type "DQ_RelativeInternalPositionalAccuracy_PropertyType"
type DQ_RelativeInternalPositionalAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_RelativeInternalPositionalAccuracy" of complex type DQ_RelativeInternalPositionalAccuracy_Type
	DQ_RelativeInternalPositionalAccuracy *DQ_RelativeInternalPositionalAccuracy_Type `xml:"DQ_RelativeInternalPositionalAccuracy,omitempty"`
}

// DQ_GriddedDataPositionalAccuracy_Type Named source named complex type "DQ_GriddedDataPositionalAccuracy_Type"
type DQ_GriddedDataPositionalAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_GriddedDataPositionalAccuracy_PropertyType Named source named complex type "DQ_GriddedDataPositionalAccuracy_PropertyType"
type DQ_GriddedDataPositionalAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_GriddedDataPositionalAccuracy" of complex type DQ_GriddedDataPositionalAccuracy_Type
	DQ_GriddedDataPositionalAccuracy *DQ_GriddedDataPositionalAccuracy_Type `xml:"DQ_GriddedDataPositionalAccuracy,omitempty"`
}

// DQ_AbsoluteExternalPositionalAccuracy_Type Named source named complex type "DQ_AbsoluteExternalPositionalAccuracy_Type"
type DQ_AbsoluteExternalPositionalAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_AbsoluteExternalPositionalAccuracy_PropertyType Named source named complex type "DQ_AbsoluteExternalPositionalAccuracy_PropertyType"
type DQ_AbsoluteExternalPositionalAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_AbsoluteExternalPositionalAccuracy" of complex type DQ_AbsoluteExternalPositionalAccuracy_Type
	DQ_AbsoluteExternalPositionalAccuracy *DQ_AbsoluteExternalPositionalAccuracy_Type `xml:"DQ_AbsoluteExternalPositionalAccuracy,omitempty"`
}

// DQ_TopologicalConsistency_Type Named source named complex type "DQ_TopologicalConsistency_Type"
type DQ_TopologicalConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_TopologicalConsistency_PropertyType Named source named complex type "DQ_TopologicalConsistency_PropertyType"
type DQ_TopologicalConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TopologicalConsistency" of complex type DQ_TopologicalConsistency_Type
	DQ_TopologicalConsistency *DQ_TopologicalConsistency_Type `xml:"DQ_TopologicalConsistency,omitempty"`
}

// DQ_FormatConsistency_Type Named source named complex type "DQ_FormatConsistency_Type"
type DQ_FormatConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_FormatConsistency_PropertyType Named source named complex type "DQ_FormatConsistency_PropertyType"
type DQ_FormatConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_FormatConsistency" of complex type DQ_FormatConsistency_Type
	DQ_FormatConsistency *DQ_FormatConsistency_Type `xml:"DQ_FormatConsistency,omitempty"`
}

// DQ_DomainConsistency_Type Named source named complex type "DQ_DomainConsistency_Type"
type DQ_DomainConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_DomainConsistency_PropertyType Named source named complex type "DQ_DomainConsistency_PropertyType"
type DQ_DomainConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_DomainConsistency" of complex type DQ_DomainConsistency_Type
	DQ_DomainConsistency *DQ_DomainConsistency_Type `xml:"DQ_DomainConsistency,omitempty"`
}

// DQ_ConceptualConsistency_Type Named source named complex type "DQ_ConceptualConsistency_Type"
type DQ_ConceptualConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_ConceptualConsistency_PropertyType Named source named complex type "DQ_ConceptualConsistency_PropertyType"
type DQ_ConceptualConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_ConceptualConsistency" of complex type DQ_ConceptualConsistency_Type
	DQ_ConceptualConsistency *DQ_ConceptualConsistency_Type `xml:"DQ_ConceptualConsistency,omitempty"`
}

// DQ_CompletenessOmission_Type Named source named complex type "DQ_CompletenessOmission_Type"
type DQ_CompletenessOmission_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_CompletenessOmission_PropertyType Named source named complex type "DQ_CompletenessOmission_PropertyType"
type DQ_CompletenessOmission_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_CompletenessOmission" of complex type DQ_CompletenessOmission_Type
	DQ_CompletenessOmission *DQ_CompletenessOmission_Type `xml:"DQ_CompletenessOmission,omitempty"`
}

// DQ_CompletenessCommission_Type Named source named complex type "DQ_CompletenessCommission_Type"
type DQ_CompletenessCommission_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_CompletenessCommission_PropertyType Named source named complex type "DQ_CompletenessCommission_PropertyType"
type DQ_CompletenessCommission_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_CompletenessCommission" of complex type DQ_CompletenessCommission_Type
	DQ_CompletenessCommission *DQ_CompletenessCommission_Type `xml:"DQ_CompletenessCommission,omitempty"`
}

// AbstractDQ_TemporalAccuracy_Type Named source named complex type "AbstractDQ_TemporalAccuracy_Type"
type AbstractDQ_TemporalAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_TemporalAccuracy_PropertyType Named source named complex type "DQ_TemporalAccuracy_PropertyType"
type DQ_TemporalAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TemporalValidity" of complex type DQ_TemporalValidity_Type
	DQ_TemporalValidity *DQ_TemporalValidity_Type `xml:"DQ_TemporalValidity,omitempty"`

	// generated from element "DQ_TemporalConsistency" of complex type DQ_TemporalConsistency_Type
	DQ_TemporalConsistency *DQ_TemporalConsistency_Type `xml:"DQ_TemporalConsistency,omitempty"`

	// generated from element "DQ_AccuracyOfATimeMeasurement" of complex type DQ_AccuracyOfATimeMeasurement_Type
	DQ_AccuracyOfATimeMeasurement *DQ_AccuracyOfATimeMeasurement_Type `xml:"DQ_AccuracyOfATimeMeasurement,omitempty"`
}

// AbstractDQ_ThematicAccuracy_Type Named source named complex type "AbstractDQ_ThematicAccuracy_Type"
type AbstractDQ_ThematicAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_ThematicAccuracy_PropertyType Named source named complex type "DQ_ThematicAccuracy_PropertyType"
type DQ_ThematicAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_QuantitativeAttributeAccuracy" of complex type DQ_QuantitativeAttributeAccuracy_Type
	DQ_QuantitativeAttributeAccuracy *DQ_QuantitativeAttributeAccuracy_Type `xml:"DQ_QuantitativeAttributeAccuracy,omitempty"`

	// generated from element "DQ_NonQuantitativeAttributeAccuracy" of complex type DQ_NonQuantitativeAttributeAccuracy_Type
	DQ_NonQuantitativeAttributeAccuracy *DQ_NonQuantitativeAttributeAccuracy_Type `xml:"DQ_NonQuantitativeAttributeAccuracy,omitempty"`

	// generated from element "DQ_ThematicClassificationCorrectness" of complex type DQ_ThematicClassificationCorrectness_Type
	DQ_ThematicClassificationCorrectness *DQ_ThematicClassificationCorrectness_Type `xml:"DQ_ThematicClassificationCorrectness,omitempty"`
}

// AbstractDQ_PositionalAccuracy_Type Named source named complex type "AbstractDQ_PositionalAccuracy_Type"
type AbstractDQ_PositionalAccuracy_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_PositionalAccuracy_PropertyType Named source named complex type "DQ_PositionalAccuracy_PropertyType"
type DQ_PositionalAccuracy_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_RelativeInternalPositionalAccuracy" of complex type DQ_RelativeInternalPositionalAccuracy_Type
	DQ_RelativeInternalPositionalAccuracy *DQ_RelativeInternalPositionalAccuracy_Type `xml:"DQ_RelativeInternalPositionalAccuracy,omitempty"`

	// generated from element "DQ_GriddedDataPositionalAccuracy" of complex type DQ_GriddedDataPositionalAccuracy_Type
	DQ_GriddedDataPositionalAccuracy *DQ_GriddedDataPositionalAccuracy_Type `xml:"DQ_GriddedDataPositionalAccuracy,omitempty"`

	// generated from element "DQ_AbsoluteExternalPositionalAccuracy" of complex type DQ_AbsoluteExternalPositionalAccuracy_Type
	DQ_AbsoluteExternalPositionalAccuracy *DQ_AbsoluteExternalPositionalAccuracy_Type `xml:"DQ_AbsoluteExternalPositionalAccuracy,omitempty"`
}

// AbstractDQ_LogicalConsistency_Type Named source named complex type "AbstractDQ_LogicalConsistency_Type"
type AbstractDQ_LogicalConsistency_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_LogicalConsistency_PropertyType Named source named complex type "DQ_LogicalConsistency_PropertyType"
type DQ_LogicalConsistency_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TopologicalConsistency" of complex type DQ_TopologicalConsistency_Type
	DQ_TopologicalConsistency *DQ_TopologicalConsistency_Type `xml:"DQ_TopologicalConsistency,omitempty"`

	// generated from element "DQ_FormatConsistency" of complex type DQ_FormatConsistency_Type
	DQ_FormatConsistency *DQ_FormatConsistency_Type `xml:"DQ_FormatConsistency,omitempty"`

	// generated from element "DQ_DomainConsistency" of complex type DQ_DomainConsistency_Type
	DQ_DomainConsistency *DQ_DomainConsistency_Type `xml:"DQ_DomainConsistency,omitempty"`

	// generated from element "DQ_ConceptualConsistency" of complex type DQ_ConceptualConsistency_Type
	DQ_ConceptualConsistency *DQ_ConceptualConsistency_Type `xml:"DQ_ConceptualConsistency,omitempty"`
}

// AbstractDQ_Completeness_Type Named source named complex type "AbstractDQ_Completeness_Type"
type AbstractDQ_Completeness_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_Completeness_PropertyType Named source named complex type "DQ_Completeness_PropertyType"
type DQ_Completeness_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_CompletenessOmission" of complex type DQ_CompletenessOmission_Type
	DQ_CompletenessOmission *DQ_CompletenessOmission_Type `xml:"DQ_CompletenessOmission,omitempty"`

	// generated from element "DQ_CompletenessCommission" of complex type DQ_CompletenessCommission_Type
	DQ_CompletenessCommission *DQ_CompletenessCommission_Type `xml:"DQ_CompletenessCommission,omitempty"`
}

// AbstractDQ_Element_Type Named source named complex type "AbstractDQ_Element_Type"
type AbstractDQ_Element_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "nameOfMeasure" of complex type CharacterString_PropertyType
	NameOfMeasure []*CharacterString_PropertyType `xml:"nameOfMeasure,omitempty"`

	// generated from element "measureIdentification" of complex type MD_Identifier_PropertyType
	MeasureIdentification *MD_Identifier_PropertyType `xml:"measureIdentification,omitempty"`

	// generated from element "measureDescription" of complex type CharacterString_PropertyType
	MeasureDescription *CharacterString_PropertyType `xml:"measureDescription,omitempty"`

	// generated from element "evaluationMethodType" of complex type DQ_EvaluationMethodTypeCode_PropertyType
	EvaluationMethodType *DQ_EvaluationMethodTypeCode_PropertyType `xml:"evaluationMethodType,omitempty"`

	// generated from element "evaluationMethodDescription" of complex type CharacterString_PropertyType
	EvaluationMethodDescription *CharacterString_PropertyType `xml:"evaluationMethodDescription,omitempty"`

	// generated from element "evaluationProcedure" of complex type CI_Citation_PropertyType
	EvaluationProcedure *CI_Citation_PropertyType `xml:"evaluationProcedure,omitempty"`

	// generated from element "dateTime" of complex type DateTime_PropertyType
	DateTime []*DateTime_PropertyType `xml:"dateTime,omitempty"`

	// generated from element "result" of complex type DQ_Result_PropertyType
	Result *DQ_Result_PropertyType `xml:"result,omitempty"`
}

// DQ_Element_PropertyType Named source named complex type "DQ_Element_PropertyType"
type DQ_Element_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_TemporalValidity" of complex type DQ_TemporalValidity_Type
	DQ_TemporalValidity *DQ_TemporalValidity_Type `xml:"DQ_TemporalValidity,omitempty"`

	// generated from element "DQ_TemporalConsistency" of complex type DQ_TemporalConsistency_Type
	DQ_TemporalConsistency *DQ_TemporalConsistency_Type `xml:"DQ_TemporalConsistency,omitempty"`

	// generated from element "DQ_AccuracyOfATimeMeasurement" of complex type DQ_AccuracyOfATimeMeasurement_Type
	DQ_AccuracyOfATimeMeasurement *DQ_AccuracyOfATimeMeasurement_Type `xml:"DQ_AccuracyOfATimeMeasurement,omitempty"`

	// generated from element "DQ_QuantitativeAttributeAccuracy" of complex type DQ_QuantitativeAttributeAccuracy_Type
	DQ_QuantitativeAttributeAccuracy *DQ_QuantitativeAttributeAccuracy_Type `xml:"DQ_QuantitativeAttributeAccuracy,omitempty"`

	// generated from element "DQ_NonQuantitativeAttributeAccuracy" of complex type DQ_NonQuantitativeAttributeAccuracy_Type
	DQ_NonQuantitativeAttributeAccuracy *DQ_NonQuantitativeAttributeAccuracy_Type `xml:"DQ_NonQuantitativeAttributeAccuracy,omitempty"`

	// generated from element "DQ_ThematicClassificationCorrectness" of complex type DQ_ThematicClassificationCorrectness_Type
	DQ_ThematicClassificationCorrectness *DQ_ThematicClassificationCorrectness_Type `xml:"DQ_ThematicClassificationCorrectness,omitempty"`

	// generated from element "DQ_RelativeInternalPositionalAccuracy" of complex type DQ_RelativeInternalPositionalAccuracy_Type
	DQ_RelativeInternalPositionalAccuracy *DQ_RelativeInternalPositionalAccuracy_Type `xml:"DQ_RelativeInternalPositionalAccuracy,omitempty"`

	// generated from element "DQ_GriddedDataPositionalAccuracy" of complex type DQ_GriddedDataPositionalAccuracy_Type
	DQ_GriddedDataPositionalAccuracy *DQ_GriddedDataPositionalAccuracy_Type `xml:"DQ_GriddedDataPositionalAccuracy,omitempty"`

	// generated from element "DQ_AbsoluteExternalPositionalAccuracy" of complex type DQ_AbsoluteExternalPositionalAccuracy_Type
	DQ_AbsoluteExternalPositionalAccuracy *DQ_AbsoluteExternalPositionalAccuracy_Type `xml:"DQ_AbsoluteExternalPositionalAccuracy,omitempty"`

	// generated from element "DQ_TopologicalConsistency" of complex type DQ_TopologicalConsistency_Type
	DQ_TopologicalConsistency *DQ_TopologicalConsistency_Type `xml:"DQ_TopologicalConsistency,omitempty"`

	// generated from element "DQ_FormatConsistency" of complex type DQ_FormatConsistency_Type
	DQ_FormatConsistency *DQ_FormatConsistency_Type `xml:"DQ_FormatConsistency,omitempty"`

	// generated from element "DQ_DomainConsistency" of complex type DQ_DomainConsistency_Type
	DQ_DomainConsistency *DQ_DomainConsistency_Type `xml:"DQ_DomainConsistency,omitempty"`

	// generated from element "DQ_ConceptualConsistency" of complex type DQ_ConceptualConsistency_Type
	DQ_ConceptualConsistency *DQ_ConceptualConsistency_Type `xml:"DQ_ConceptualConsistency,omitempty"`

	// generated from element "DQ_CompletenessOmission" of complex type DQ_CompletenessOmission_Type
	DQ_CompletenessOmission *DQ_CompletenessOmission_Type `xml:"DQ_CompletenessOmission,omitempty"`

	// generated from element "DQ_CompletenessCommission" of complex type DQ_CompletenessCommission_Type
	DQ_CompletenessCommission *DQ_CompletenessCommission_Type `xml:"DQ_CompletenessCommission,omitempty"`
}

// DQ_DataQuality_Type Named source named complex type "DQ_DataQuality_Type"
type DQ_DataQuality_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "scope" of complex type DQ_Scope_PropertyType
	Scope *DQ_Scope_PropertyType `xml:"scope,omitempty"`

	// generated from element "report" of complex type DQ_Element_PropertyType
	Report []*DQ_Element_PropertyType `xml:"report,omitempty"`

	// generated from element "lineage" of complex type LI_Lineage_PropertyType
	Lineage *LI_Lineage_PropertyType `xml:"lineage,omitempty"`
}

// DQ_DataQuality_PropertyType Named source named complex type "DQ_DataQuality_PropertyType"
type DQ_DataQuality_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_DataQuality" of complex type DQ_DataQuality_Type
	DQ_DataQuality *DQ_DataQuality_Type `xml:"DQ_DataQuality,omitempty"`
}

// DQ_Scope_Type Named source named complex type "DQ_Scope_Type"
type DQ_Scope_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "level" of complex type MD_ScopeCode_PropertyType
	Level *MD_ScopeCode_PropertyType `xml:"level,omitempty"`

	// generated from element "extent" of complex type EX_Extent_PropertyType
	Extent *EX_Extent_PropertyType `xml:"extent,omitempty"`

	// generated from element "levelDescription" of complex type MD_ScopeDescription_PropertyType
	LevelDescription []*MD_ScopeDescription_PropertyType `xml:"levelDescription,omitempty"`
}

// DQ_Scope_PropertyType Named source named complex type "DQ_Scope_PropertyType"
type DQ_Scope_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DQ_Scope" of complex type DQ_Scope_Type
	DQ_Scope *DQ_Scope_Type `xml:"DQ_Scope,omitempty"`
}

// DQ_EvaluationMethodTypeCode_PropertyType Named source named complex type "DQ_EvaluationMethodTypeCode_PropertyType"
type DQ_EvaluationMethodTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "DQ_EvaluationMethodTypeCode" of complex type CodeListValue_Type
	DQ_EvaluationMethodTypeCode *gcobase.CodeListValue_Type `xml:"DQ_EvaluationMethodTypeCode,omitempty"`
}

// AbstractMD_Identification_Type Named source named complex type "AbstractMD_Identification_Type"
// Basic information about data
type AbstractMD_Identification_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "citation" of complex type CI_Citation_PropertyType
	Citation *CI_Citation_PropertyType `xml:"citation,omitempty"`

	// generated from element "abstract" of complex type CharacterString_PropertyType
	Abstract *CharacterString_PropertyType `xml:"abstract,omitempty"`

	// generated from element "purpose" of complex type CharacterString_PropertyType
	Purpose *CharacterString_PropertyType `xml:"purpose,omitempty"`

	// generated from element "credit" of complex type CharacterString_PropertyType
	Credit []*CharacterString_PropertyType `xml:"credit,omitempty"`

	// generated from element "status" of complex type MD_ProgressCode_PropertyType
	Status []*MD_ProgressCode_PropertyType `xml:"status,omitempty"`

	// generated from element "pointOfContact" of complex type CI_ResponsibleParty_PropertyType
	PointOfContact []*CI_ResponsibleParty_PropertyType `xml:"pointOfContact,omitempty"`

	// generated from element "resourceMaintenance" of complex type MD_MaintenanceInformation_PropertyType
	ResourceMaintenance []*MD_MaintenanceInformation_PropertyType `xml:"resourceMaintenance,omitempty"`

	// generated from element "graphicOverview" of complex type MD_BrowseGraphic_PropertyType
	GraphicOverview []*MD_BrowseGraphic_PropertyType `xml:"graphicOverview,omitempty"`

	// generated from element "resourceFormat" of complex type MD_Format_PropertyType
	ResourceFormat []*MD_Format_PropertyType `xml:"resourceFormat,omitempty"`

	// generated from element "descriptiveKeywords" of complex type MD_Keywords_PropertyType
	DescriptiveKeywords []*MD_Keywords_PropertyType `xml:"descriptiveKeywords,omitempty"`

	// generated from element "resourceSpecificUsage" of complex type MD_Usage_PropertyType
	ResourceSpecificUsage []*MD_Usage_PropertyType `xml:"resourceSpecificUsage,omitempty"`

	// generated from element "resourceConstraints" of complex type MD_Constraints_PropertyType
	ResourceConstraints []*MD_Constraints_PropertyType `xml:"resourceConstraints,omitempty"`

	// generated from element "aggregationInfo" of complex type MD_AggregateInformation_PropertyType
	AggregationInfo []*MD_AggregateInformation_PropertyType `xml:"aggregationInfo,omitempty"`
}

// MD_Identification_PropertyType Named source named complex type "MD_Identification_PropertyType"
type MD_Identification_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_DataIdentification" of complex type MD_DataIdentification_Type
	MD_DataIdentification *MD_DataIdentification_Type `xml:"MD_DataIdentification,omitempty"`

	// generated from element "MD_ServiceIdentification" of complex type MD_ServiceIdentification_Type
	MD_ServiceIdentification *MD_ServiceIdentification_Type `xml:"MD_ServiceIdentification,omitempty"`
}

// MD_BrowseGraphic_Type Named source named complex type "MD_BrowseGraphic_Type"
// Graphic that provides an illustration of the dataset (should include a legend for the graphic)
type MD_BrowseGraphic_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "fileName" of complex type CharacterString_PropertyType
	FileName *CharacterString_PropertyType `xml:"fileName,omitempty"`

	// generated from element "fileDescription" of complex type CharacterString_PropertyType
	FileDescription *CharacterString_PropertyType `xml:"fileDescription,omitempty"`

	// generated from element "fileType" of complex type CharacterString_PropertyType
	FileType *CharacterString_PropertyType `xml:"fileType,omitempty"`
}

// MD_BrowseGraphic_PropertyType Named source named complex type "MD_BrowseGraphic_PropertyType"
type MD_BrowseGraphic_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_BrowseGraphic" of complex type MD_BrowseGraphic_Type
	MD_BrowseGraphic *MD_BrowseGraphic_Type `xml:"MD_BrowseGraphic,omitempty"`
}

// MD_DataIdentification_Type Named source named complex type "MD_DataIdentification_Type"
type MD_DataIdentification_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "citation" of complex type CI_Citation_PropertyType
	Citation *CI_Citation_PropertyType `xml:"citation,omitempty"`

	// generated from element "abstract" of complex type CharacterString_PropertyType
	Abstract *CharacterString_PropertyType `xml:"abstract,omitempty"`

	// generated from element "purpose" of complex type CharacterString_PropertyType
	Purpose *CharacterString_PropertyType `xml:"purpose,omitempty"`

	// generated from element "credit" of complex type CharacterString_PropertyType
	Credit []*CharacterString_PropertyType `xml:"credit,omitempty"`

	// generated from element "status" of complex type MD_ProgressCode_PropertyType
	Status []*MD_ProgressCode_PropertyType `xml:"status,omitempty"`

	// generated from element "pointOfContact" of complex type CI_ResponsibleParty_PropertyType
	PointOfContact []*CI_ResponsibleParty_PropertyType `xml:"pointOfContact,omitempty"`

	// generated from element "resourceMaintenance" of complex type MD_MaintenanceInformation_PropertyType
	ResourceMaintenance []*MD_MaintenanceInformation_PropertyType `xml:"resourceMaintenance,omitempty"`

	// generated from element "graphicOverview" of complex type MD_BrowseGraphic_PropertyType
	GraphicOverview []*MD_BrowseGraphic_PropertyType `xml:"graphicOverview,omitempty"`

	// generated from element "resourceFormat" of complex type MD_Format_PropertyType
	ResourceFormat []*MD_Format_PropertyType `xml:"resourceFormat,omitempty"`

	// generated from element "descriptiveKeywords" of complex type MD_Keywords_PropertyType
	DescriptiveKeywords []*MD_Keywords_PropertyType `xml:"descriptiveKeywords,omitempty"`

	// generated from element "resourceSpecificUsage" of complex type MD_Usage_PropertyType
	ResourceSpecificUsage []*MD_Usage_PropertyType `xml:"resourceSpecificUsage,omitempty"`

	// generated from element "resourceConstraints" of complex type MD_Constraints_PropertyType
	ResourceConstraints []*MD_Constraints_PropertyType `xml:"resourceConstraints,omitempty"`

	// generated from element "aggregationInfo" of complex type MD_AggregateInformation_PropertyType
	AggregationInfo []*MD_AggregateInformation_PropertyType `xml:"aggregationInfo,omitempty"`

	// generated from element "spatialRepresentationType" of complex type MD_SpatialRepresentationTypeCode_PropertyType
	SpatialRepresentationType []*MD_SpatialRepresentationTypeCode_PropertyType `xml:"spatialRepresentationType,omitempty"`

	// generated from element "spatialResolution" of complex type MD_Resolution_PropertyType
	SpatialResolution []*MD_Resolution_PropertyType `xml:"spatialResolution,omitempty"`

	// generated from element "language" of complex type CharacterString_PropertyType
	Language []*CharacterString_PropertyType `xml:"language,omitempty"`

	// generated from element "characterSet" of complex type MD_CharacterSetCode_PropertyType
	CharacterSet []*MD_CharacterSetCode_PropertyType `xml:"characterSet,omitempty"`

	// generated from element "topicCategory" of complex type MD_TopicCategoryCode_PropertyType
	TopicCategory []*MD_TopicCategoryCode_PropertyType `xml:"topicCategory,omitempty"`

	// generated from element "environmentDescription" of complex type CharacterString_PropertyType
	EnvironmentDescription *CharacterString_PropertyType `xml:"environmentDescription,omitempty"`

	// generated from element "extent" of complex type EX_Extent_PropertyType
	Extent []*EX_Extent_PropertyType `xml:"extent,omitempty"`

	// generated from element "supplementalInformation" of complex type CharacterString_PropertyType
	SupplementalInformation *CharacterString_PropertyType `xml:"supplementalInformation,omitempty"`
}

// MD_DataIdentification_PropertyType Named source named complex type "MD_DataIdentification_PropertyType"
type MD_DataIdentification_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_DataIdentification" of complex type MD_DataIdentification_Type
	MD_DataIdentification *MD_DataIdentification_Type `xml:"MD_DataIdentification,omitempty"`
}

// MD_ServiceIdentification_Type Named source named complex type "MD_ServiceIdentification_Type"
// See 19119 for further info
type MD_ServiceIdentification_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "citation" of complex type CI_Citation_PropertyType
	Citation *CI_Citation_PropertyType `xml:"citation,omitempty"`

	// generated from element "abstract" of complex type CharacterString_PropertyType
	Abstract *CharacterString_PropertyType `xml:"abstract,omitempty"`

	// generated from element "purpose" of complex type CharacterString_PropertyType
	Purpose *CharacterString_PropertyType `xml:"purpose,omitempty"`

	// generated from element "credit" of complex type CharacterString_PropertyType
	Credit []*CharacterString_PropertyType `xml:"credit,omitempty"`

	// generated from element "status" of complex type MD_ProgressCode_PropertyType
	Status []*MD_ProgressCode_PropertyType `xml:"status,omitempty"`

	// generated from element "pointOfContact" of complex type CI_ResponsibleParty_PropertyType
	PointOfContact []*CI_ResponsibleParty_PropertyType `xml:"pointOfContact,omitempty"`

	// generated from element "resourceMaintenance" of complex type MD_MaintenanceInformation_PropertyType
	ResourceMaintenance []*MD_MaintenanceInformation_PropertyType `xml:"resourceMaintenance,omitempty"`

	// generated from element "graphicOverview" of complex type MD_BrowseGraphic_PropertyType
	GraphicOverview []*MD_BrowseGraphic_PropertyType `xml:"graphicOverview,omitempty"`

	// generated from element "resourceFormat" of complex type MD_Format_PropertyType
	ResourceFormat []*MD_Format_PropertyType `xml:"resourceFormat,omitempty"`

	// generated from element "descriptiveKeywords" of complex type MD_Keywords_PropertyType
	DescriptiveKeywords []*MD_Keywords_PropertyType `xml:"descriptiveKeywords,omitempty"`

	// generated from element "resourceSpecificUsage" of complex type MD_Usage_PropertyType
	ResourceSpecificUsage []*MD_Usage_PropertyType `xml:"resourceSpecificUsage,omitempty"`

	// generated from element "resourceConstraints" of complex type MD_Constraints_PropertyType
	ResourceConstraints []*MD_Constraints_PropertyType `xml:"resourceConstraints,omitempty"`

	// generated from element "aggregationInfo" of complex type MD_AggregateInformation_PropertyType
	AggregationInfo []*MD_AggregateInformation_PropertyType `xml:"aggregationInfo,omitempty"`
}

// MD_ServiceIdentification_PropertyType Named source named complex type "MD_ServiceIdentification_PropertyType"
type MD_ServiceIdentification_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_ServiceIdentification" of complex type MD_ServiceIdentification_Type
	MD_ServiceIdentification *MD_ServiceIdentification_Type `xml:"MD_ServiceIdentification,omitempty"`
}

// MD_RepresentativeFraction_Type Named source named complex type "MD_RepresentativeFraction_Type"
type MD_RepresentativeFraction_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "denominator" of complex type Integer_PropertyType
	Denominator *Integer_PropertyType `xml:"denominator,omitempty"`
}

// MD_RepresentativeFraction_PropertyType Named source named complex type "MD_RepresentativeFraction_PropertyType"
type MD_RepresentativeFraction_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_RepresentativeFraction" of complex type MD_RepresentativeFraction_Type
	MD_RepresentativeFraction *MD_RepresentativeFraction_Type `xml:"MD_RepresentativeFraction,omitempty"`
}

// MD_Usage_Type Named source named complex type "MD_Usage_Type"
// Brief description of ways in which the dataset is currently used.
type MD_Usage_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "specificUsage" of complex type CharacterString_PropertyType
	SpecificUsage *CharacterString_PropertyType `xml:"specificUsage,omitempty"`

	// generated from element "usageDateTime" of complex type DateTime_PropertyType
	UsageDateTime *DateTime_PropertyType `xml:"usageDateTime,omitempty"`

	// generated from element "userDeterminedLimitations" of complex type CharacterString_PropertyType
	UserDeterminedLimitations *CharacterString_PropertyType `xml:"userDeterminedLimitations,omitempty"`

	// generated from element "userContactInfo" of complex type CI_ResponsibleParty_PropertyType
	UserContactInfo []*CI_ResponsibleParty_PropertyType `xml:"userContactInfo,omitempty"`
}

// MD_Usage_PropertyType Named source named complex type "MD_Usage_PropertyType"
type MD_Usage_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Usage" of complex type MD_Usage_Type
	MD_Usage *MD_Usage_Type `xml:"MD_Usage,omitempty"`
}

// MD_Keywords_Type Named source named complex type "MD_Keywords_Type"
// Keywords, their type and reference source
type MD_Keywords_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "keyword" of complex type CharacterString_PropertyType
	Keyword []*CharacterString_PropertyType `xml:"keyword,omitempty"`

	// generated from element "type" of complex type MD_KeywordTypeCode_PropertyType
	Type_XSD *MD_KeywordTypeCode_PropertyType `xml:"type,omitempty"`

	// generated from element "thesaurusName" of complex type CI_Citation_PropertyType
	ThesaurusName *CI_Citation_PropertyType `xml:"thesaurusName,omitempty"`
}

// MD_Keywords_PropertyType Named source named complex type "MD_Keywords_PropertyType"
type MD_Keywords_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Keywords" of complex type MD_Keywords_Type
	MD_Keywords *MD_Keywords_Type `xml:"MD_Keywords,omitempty"`
}

// DS_Association_Type Named source named complex type "DS_Association_Type"
type DS_Association_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification
}

// DS_Association_PropertyType Named source named complex type "DS_Association_PropertyType"
type DS_Association_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_Association" of complex type DS_Association_Type
	DS_Association *DS_Association_Type `xml:"DS_Association,omitempty"`
}

// MD_AggregateInformation_Type Named source named complex type "MD_AggregateInformation_Type"
// Encapsulates the dataset aggregation information
type MD_AggregateInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "aggregateDataSetName" of complex type CI_Citation_PropertyType
	AggregateDataSetName *CI_Citation_PropertyType `xml:"aggregateDataSetName,omitempty"`

	// generated from element "aggregateDataSetIdentifier" of complex type MD_Identifier_PropertyType
	AggregateDataSetIdentifier *MD_Identifier_PropertyType `xml:"aggregateDataSetIdentifier,omitempty"`

	// generated from element "associationType" of complex type DS_AssociationTypeCode_PropertyType
	AssociationType *DS_AssociationTypeCode_PropertyType `xml:"associationType,omitempty"`

	// generated from element "initiativeType" of complex type DS_InitiativeTypeCode_PropertyType
	InitiativeType *DS_InitiativeTypeCode_PropertyType `xml:"initiativeType,omitempty"`
}

// MD_AggregateInformation_PropertyType Named source named complex type "MD_AggregateInformation_PropertyType"
type MD_AggregateInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_AggregateInformation" of complex type MD_AggregateInformation_Type
	MD_AggregateInformation *MD_AggregateInformation_Type `xml:"MD_AggregateInformation,omitempty"`
}

// MD_Resolution_Type Named source named complex type "MD_Resolution_Type"
type MD_Resolution_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "equivalentScale" of complex type MD_RepresentativeFraction_PropertyType
	EquivalentScale *MD_RepresentativeFraction_PropertyType `xml:"equivalentScale,omitempty"`

	// generated from element "distance" of complex type Distance_PropertyType
	Distance *Distance_PropertyType `xml:"distance,omitempty"`
}

// MD_Resolution_PropertyType Named source named complex type "MD_Resolution_PropertyType"
type MD_Resolution_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_Resolution" of complex type MD_Resolution_Type
	MD_Resolution *MD_Resolution_Type `xml:"MD_Resolution,omitempty"`
}

// MD_TopicCategoryCode_PropertyType Named source named complex type "MD_TopicCategoryCode_PropertyType"
type MD_TopicCategoryCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_TopicCategoryCode" of simple type gmd:MD_TopicCategoryCode_Type
	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type `xml:"MD_TopicCategoryCode,omitempty"`
}

// MD_CharacterSetCode_PropertyType Named source named complex type "MD_CharacterSetCode_PropertyType"
type MD_CharacterSetCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_CharacterSetCode" of complex type CodeListValue_Type
	MD_CharacterSetCode *gcobase.CodeListValue_Type `xml:"MD_CharacterSetCode,omitempty"`
}

// MD_SpatialRepresentationTypeCode_PropertyType Named source named complex type "MD_SpatialRepresentationTypeCode_PropertyType"
type MD_SpatialRepresentationTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_SpatialRepresentationTypeCode" of complex type CodeListValue_Type
	MD_SpatialRepresentationTypeCode *gcobase.CodeListValue_Type `xml:"MD_SpatialRepresentationTypeCode,omitempty"`
}

// MD_ProgressCode_PropertyType Named source named complex type "MD_ProgressCode_PropertyType"
type MD_ProgressCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ProgressCode" of complex type CodeListValue_Type
	MD_ProgressCode *gcobase.CodeListValue_Type `xml:"MD_ProgressCode,omitempty"`
}

// MD_KeywordTypeCode_PropertyType Named source named complex type "MD_KeywordTypeCode_PropertyType"
type MD_KeywordTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_KeywordTypeCode" of complex type CodeListValue_Type
	MD_KeywordTypeCode *gcobase.CodeListValue_Type `xml:"MD_KeywordTypeCode,omitempty"`
}

// DS_AssociationTypeCode_PropertyType Named source named complex type "DS_AssociationTypeCode_PropertyType"
type DS_AssociationTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "DS_AssociationTypeCode" of complex type CodeListValue_Type
	DS_AssociationTypeCode *gcobase.CodeListValue_Type `xml:"DS_AssociationTypeCode,omitempty"`
}

// DS_InitiativeTypeCode_PropertyType Named source named complex type "DS_InitiativeTypeCode_PropertyType"
type DS_InitiativeTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "DS_InitiativeTypeCode" of complex type CodeListValue_Type
	DS_InitiativeTypeCode *gcobase.CodeListValue_Type `xml:"DS_InitiativeTypeCode,omitempty"`
}

// MD_Constraints_Type Named source named complex type "MD_Constraints_Type"
// Restrictions on the access and use of a dataset or metadata
type MD_Constraints_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "useLimitation" of complex type CharacterString_PropertyType
	UseLimitation []*CharacterString_PropertyType `xml:"useLimitation,omitempty"`
}

// MD_Constraints_PropertyType Named source named complex type "MD_Constraints_PropertyType"
type MD_Constraints_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Constraints" of complex type MD_Constraints_Type
	MD_Constraints *MD_Constraints_Type `xml:"MD_Constraints,omitempty"`

	// generated from element "MD_LegalConstraints" of complex type MD_LegalConstraints_Type
	MD_LegalConstraints *MD_LegalConstraints_Type `xml:"MD_LegalConstraints,omitempty"`

	// generated from element "MD_SecurityConstraints" of complex type MD_SecurityConstraints_Type
	MD_SecurityConstraints *MD_SecurityConstraints_Type `xml:"MD_SecurityConstraints,omitempty"`
}

// MD_LegalConstraints_Type Named source named complex type "MD_LegalConstraints_Type"
// Restrictions and legal prerequisites for accessing and using the dataset.
type MD_LegalConstraints_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "useLimitation" of complex type CharacterString_PropertyType
	UseLimitation []*CharacterString_PropertyType `xml:"useLimitation,omitempty"`

	// generated from element "accessConstraints" of complex type MD_RestrictionCode_PropertyType
	AccessConstraints []*MD_RestrictionCode_PropertyType `xml:"accessConstraints,omitempty"`

	// generated from element "useConstraints" of complex type MD_RestrictionCode_PropertyType
	UseConstraints []*MD_RestrictionCode_PropertyType `xml:"useConstraints,omitempty"`

	// generated from element "otherConstraints" of complex type CharacterString_PropertyType
	OtherConstraints []*CharacterString_PropertyType `xml:"otherConstraints,omitempty"`
}

// MD_LegalConstraints_PropertyType Named source named complex type "MD_LegalConstraints_PropertyType"
type MD_LegalConstraints_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_LegalConstraints" of complex type MD_LegalConstraints_Type
	MD_LegalConstraints *MD_LegalConstraints_Type `xml:"MD_LegalConstraints,omitempty"`
}

// MD_SecurityConstraints_Type Named source named complex type "MD_SecurityConstraints_Type"
// Handling restrictions imposed on the dataset because of national security, privacy, or other concerns
type MD_SecurityConstraints_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "useLimitation" of complex type CharacterString_PropertyType
	UseLimitation []*CharacterString_PropertyType `xml:"useLimitation,omitempty"`

	// generated from element "classification" of complex type MD_ClassificationCode_PropertyType
	Classification *MD_ClassificationCode_PropertyType `xml:"classification,omitempty"`

	// generated from element "userNote" of complex type CharacterString_PropertyType
	UserNote *CharacterString_PropertyType `xml:"userNote,omitempty"`

	// generated from element "classificationSystem" of complex type CharacterString_PropertyType
	ClassificationSystem *CharacterString_PropertyType `xml:"classificationSystem,omitempty"`

	// generated from element "handlingDescription" of complex type CharacterString_PropertyType
	HandlingDescription *CharacterString_PropertyType `xml:"handlingDescription,omitempty"`
}

// MD_SecurityConstraints_PropertyType Named source named complex type "MD_SecurityConstraints_PropertyType"
type MD_SecurityConstraints_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_SecurityConstraints" of complex type MD_SecurityConstraints_Type
	MD_SecurityConstraints *MD_SecurityConstraints_Type `xml:"MD_SecurityConstraints,omitempty"`
}

// MD_ClassificationCode_PropertyType Named source named complex type "MD_ClassificationCode_PropertyType"
type MD_ClassificationCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ClassificationCode" of complex type CodeListValue_Type
	MD_ClassificationCode *gcobase.CodeListValue_Type `xml:"MD_ClassificationCode,omitempty"`
}

// MD_RestrictionCode_PropertyType Named source named complex type "MD_RestrictionCode_PropertyType"
type MD_RestrictionCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_RestrictionCode" of complex type CodeListValue_Type
	MD_RestrictionCode *gcobase.CodeListValue_Type `xml:"MD_RestrictionCode,omitempty"`
}

// MD_Medium_Type Named source named complex type "MD_Medium_Type"
// Information about the media on which the data can be distributed
type MD_Medium_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type MD_MediumNameCode_PropertyType
	NameXSD *MD_MediumNameCode_PropertyType `xml:"name,omitempty"`

	// generated from element "density" of complex type Real_PropertyType
	Density []*Real_PropertyType `xml:"density,omitempty"`

	// generated from element "densityUnits" of complex type CharacterString_PropertyType
	DensityUnits *CharacterString_PropertyType `xml:"densityUnits,omitempty"`

	// generated from element "volumes" of complex type Integer_PropertyType
	Volumes *Integer_PropertyType `xml:"volumes,omitempty"`

	// generated from element "mediumFormat" of complex type MD_MediumFormatCode_PropertyType
	MediumFormat []*MD_MediumFormatCode_PropertyType `xml:"mediumFormat,omitempty"`

	// generated from element "mediumNote" of complex type CharacterString_PropertyType
	MediumNote *CharacterString_PropertyType `xml:"mediumNote,omitempty"`
}

// MD_Medium_PropertyType Named source named complex type "MD_Medium_PropertyType"
type MD_Medium_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Medium" of complex type MD_Medium_Type
	MD_Medium *MD_Medium_Type `xml:"MD_Medium,omitempty"`
}

// MD_DigitalTransferOptions_Type Named source named complex type "MD_DigitalTransferOptions_Type"
// Technical means and media by which a dataset is obtained from the distributor
type MD_DigitalTransferOptions_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "unitsOfDistribution" of complex type CharacterString_PropertyType
	UnitsOfDistribution *CharacterString_PropertyType `xml:"unitsOfDistribution,omitempty"`

	// generated from element "transferSize" of complex type Real_PropertyType
	TransferSize *Real_PropertyType `xml:"transferSize,omitempty"`

	// generated from element "onLine" of complex type CI_OnlineResource_PropertyType
	OnLine []*CI_OnlineResource_PropertyType `xml:"onLine,omitempty"`

	// generated from element "offLine" of complex type MD_Medium_PropertyType
	OffLine *MD_Medium_PropertyType `xml:"offLine,omitempty"`
}

// MD_DigitalTransferOptions_PropertyType Named source named complex type "MD_DigitalTransferOptions_PropertyType"
type MD_DigitalTransferOptions_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_DigitalTransferOptions" of complex type MD_DigitalTransferOptions_Type
	MD_DigitalTransferOptions *MD_DigitalTransferOptions_Type `xml:"MD_DigitalTransferOptions,omitempty"`
}

// MD_StandardOrderProcess_Type Named source named complex type "MD_StandardOrderProcess_Type"
// Common ways in which the dataset may be obtained or received, and related instructions and fee information
type MD_StandardOrderProcess_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "fees" of complex type CharacterString_PropertyType
	Fees *CharacterString_PropertyType `xml:"fees,omitempty"`

	// generated from element "plannedAvailableDateTime" of complex type DateTime_PropertyType
	PlannedAvailableDateTime *DateTime_PropertyType `xml:"plannedAvailableDateTime,omitempty"`

	// generated from element "orderingInstructions" of complex type CharacterString_PropertyType
	OrderingInstructions *CharacterString_PropertyType `xml:"orderingInstructions,omitempty"`

	// generated from element "turnaround" of complex type CharacterString_PropertyType
	Turnaround *CharacterString_PropertyType `xml:"turnaround,omitempty"`
}

// MD_StandardOrderProcess_PropertyType Named source named complex type "MD_StandardOrderProcess_PropertyType"
type MD_StandardOrderProcess_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_StandardOrderProcess" of complex type MD_StandardOrderProcess_Type
	MD_StandardOrderProcess *MD_StandardOrderProcess_Type `xml:"MD_StandardOrderProcess,omitempty"`
}

// MD_Distributor_Type Named source named complex type "MD_Distributor_Type"
// Information about the distributor
type MD_Distributor_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "distributorContact" of complex type CI_ResponsibleParty_PropertyType
	DistributorContact *CI_ResponsibleParty_PropertyType `xml:"distributorContact,omitempty"`

	// generated from element "distributionOrderProcess" of complex type MD_StandardOrderProcess_PropertyType
	DistributionOrderProcess []*MD_StandardOrderProcess_PropertyType `xml:"distributionOrderProcess,omitempty"`

	// generated from element "distributorFormat" of complex type MD_Format_PropertyType
	DistributorFormat []*MD_Format_PropertyType `xml:"distributorFormat,omitempty"`

	// generated from element "distributorTransferOptions" of complex type MD_DigitalTransferOptions_PropertyType
	DistributorTransferOptions []*MD_DigitalTransferOptions_PropertyType `xml:"distributorTransferOptions,omitempty"`
}

// MD_Distributor_PropertyType Named source named complex type "MD_Distributor_PropertyType"
type MD_Distributor_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Distributor" of complex type MD_Distributor_Type
	MD_Distributor *MD_Distributor_Type `xml:"MD_Distributor,omitempty"`
}

// MD_Distribution_Type Named source named complex type "MD_Distribution_Type"
// Information about the distributor of and options for obtaining the dataset
type MD_Distribution_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "distributionFormat" of complex type MD_Format_PropertyType
	DistributionFormat []*MD_Format_PropertyType `xml:"distributionFormat,omitempty"`

	// generated from element "distributor" of complex type MD_Distributor_PropertyType
	Distributor []*MD_Distributor_PropertyType `xml:"distributor,omitempty"`

	// generated from element "transferOptions" of complex type MD_DigitalTransferOptions_PropertyType
	TransferOptions []*MD_DigitalTransferOptions_PropertyType `xml:"transferOptions,omitempty"`
}

// MD_Distribution_PropertyType Named source named complex type "MD_Distribution_PropertyType"
type MD_Distribution_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Distribution" of complex type MD_Distribution_Type
	MD_Distribution *MD_Distribution_Type `xml:"MD_Distribution,omitempty"`
}

// MD_Format_Type Named source named complex type "MD_Format_Type"
// Description of the form of the data to be distributed
type MD_Format_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type CharacterString_PropertyType
	NameXSD *CharacterString_PropertyType `xml:"name,omitempty"`

	// generated from element "version" of complex type CharacterString_PropertyType
	Version *CharacterString_PropertyType `xml:"version,omitempty"`

	// generated from element "amendmentNumber" of complex type CharacterString_PropertyType
	AmendmentNumber *CharacterString_PropertyType `xml:"amendmentNumber,omitempty"`

	// generated from element "specification" of complex type CharacterString_PropertyType
	Specification *CharacterString_PropertyType `xml:"specification,omitempty"`

	// generated from element "fileDecompressionTechnique" of complex type CharacterString_PropertyType
	FileDecompressionTechnique *CharacterString_PropertyType `xml:"fileDecompressionTechnique,omitempty"`

	// generated from element "formatDistributor" of complex type MD_Distributor_PropertyType
	FormatDistributor []*MD_Distributor_PropertyType `xml:"formatDistributor,omitempty"`
}

// MD_Format_PropertyType Named source named complex type "MD_Format_PropertyType"
type MD_Format_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Format" of complex type MD_Format_Type
	MD_Format *MD_Format_Type `xml:"MD_Format,omitempty"`
}

// MD_DistributionUnits_PropertyType Named source named complex type "MD_DistributionUnits_PropertyType"
type MD_DistributionUnits_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_DistributionUnits" of complex type CodeListValue_Type
	MD_DistributionUnits *gcobase.CodeListValue_Type `xml:"MD_DistributionUnits,omitempty"`
}

// MD_MediumFormatCode_PropertyType Named source named complex type "MD_MediumFormatCode_PropertyType"
type MD_MediumFormatCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_MediumFormatCode" of complex type CodeListValue_Type
	MD_MediumFormatCode *gcobase.CodeListValue_Type `xml:"MD_MediumFormatCode,omitempty"`
}

// MD_MediumNameCode_PropertyType Named source named complex type "MD_MediumNameCode_PropertyType"
type MD_MediumNameCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_MediumNameCode" of complex type CodeListValue_Type
	MD_MediumNameCode *gcobase.CodeListValue_Type `xml:"MD_MediumNameCode,omitempty"`
}

// MD_MaintenanceInformation_Type Named source named complex type "MD_MaintenanceInformation_Type"
// Information about the scope and frequency of updating
type MD_MaintenanceInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "maintenanceAndUpdateFrequency" of complex type MD_MaintenanceFrequencyCode_PropertyType
	MaintenanceAndUpdateFrequency *MD_MaintenanceFrequencyCode_PropertyType `xml:"maintenanceAndUpdateFrequency,omitempty"`

	// generated from element "dateOfNextUpdate" of complex type Date_PropertyType
	DateOfNextUpdate *Date_PropertyType `xml:"dateOfNextUpdate,omitempty"`

	// generated from element "userDefinedMaintenanceFrequency" of complex type TM_PeriodDuration_PropertyType
	UserDefinedMaintenanceFrequency *TM_PeriodDuration_PropertyType `xml:"userDefinedMaintenanceFrequency,omitempty"`

	// generated from element "updateScope" of complex type MD_ScopeCode_PropertyType
	UpdateScope []*MD_ScopeCode_PropertyType `xml:"updateScope,omitempty"`

	// generated from element "updateScopeDescription" of complex type MD_ScopeDescription_PropertyType
	UpdateScopeDescription []*MD_ScopeDescription_PropertyType `xml:"updateScopeDescription,omitempty"`

	// generated from element "maintenanceNote" of complex type CharacterString_PropertyType
	MaintenanceNote []*CharacterString_PropertyType `xml:"maintenanceNote,omitempty"`

	// generated from element "contact" of complex type CI_ResponsibleParty_PropertyType
	Contact []*CI_ResponsibleParty_PropertyType `xml:"contact,omitempty"`
}

// MD_MaintenanceInformation_PropertyType Named source named complex type "MD_MaintenanceInformation_PropertyType"
type MD_MaintenanceInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_MaintenanceInformation" of complex type MD_MaintenanceInformation_Type
	MD_MaintenanceInformation *MD_MaintenanceInformation_Type `xml:"MD_MaintenanceInformation,omitempty"`
}

// MD_ScopeDescription_Type Named source named complex type "MD_ScopeDescription_Type"
// Description of the class of information covered by the information
type MD_ScopeDescription_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "attributes" of complex type ObjectReference_PropertyType
	Attributes []*gcobase.ObjectReference_PropertyType `xml:"attributes,omitempty"`

	// generated from element "features" of complex type ObjectReference_PropertyType
	Features []*gcobase.ObjectReference_PropertyType `xml:"features,omitempty"`

	// generated from element "featureInstances" of complex type ObjectReference_PropertyType
	FeatureInstances []*gcobase.ObjectReference_PropertyType `xml:"featureInstances,omitempty"`

	// generated from element "attributeInstances" of complex type ObjectReference_PropertyType
	AttributeInstances []*gcobase.ObjectReference_PropertyType `xml:"attributeInstances,omitempty"`

	// generated from element "dataset" of complex type CharacterString_PropertyType
	Dataset *CharacterString_PropertyType `xml:"dataset,omitempty"`

	// generated from element "other" of complex type CharacterString_PropertyType
	Other *CharacterString_PropertyType `xml:"other,omitempty"`
}

// MD_ScopeDescription_PropertyType Named source named complex type "MD_ScopeDescription_PropertyType"
type MD_ScopeDescription_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ScopeDescription" of complex type MD_ScopeDescription_Type
	MD_ScopeDescription *MD_ScopeDescription_Type `xml:"MD_ScopeDescription,omitempty"`
}

// MD_MaintenanceFrequencyCode_PropertyType Named source named complex type "MD_MaintenanceFrequencyCode_PropertyType"
type MD_MaintenanceFrequencyCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_MaintenanceFrequencyCode" of complex type CodeListValue_Type
	MD_MaintenanceFrequencyCode *gcobase.CodeListValue_Type `xml:"MD_MaintenanceFrequencyCode,omitempty"`
}

// MD_ScopeCode_PropertyType Named source named complex type "MD_ScopeCode_PropertyType"
type MD_ScopeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ScopeCode" of complex type CodeListValue_Type
	MD_ScopeCode *gcobase.CodeListValue_Type `xml:"MD_ScopeCode,omitempty"`
}

// PT_FreeText_Type Named source named complex type "PT_FreeText_Type"
type PT_FreeText_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "textGroup" of complex type LocalisedCharacterString_PropertyType
	TextGroup []*LocalisedCharacterString_PropertyType `xml:"textGroup,omitempty"`
}

// PT_FreeText_PropertyType Named source named complex type "PT_FreeText_PropertyType"
type PT_FreeText_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "CharacterString" of simple type xs:string
	CharacterString string `xml:"CharacterString,omitempty"`

	// generated from element "MD_PixelOrientationCode" of simple type gmd:MD_PixelOrientationCode_Type
	MD_PixelOrientationCode string `xml:"MD_PixelOrientationCode,omitempty"`

	// generated from element "MD_TopologyLevelCode" of complex type CodeListValue_Type
	MD_TopologyLevelCode *gcobase.CodeListValue_Type `xml:"MD_TopologyLevelCode,omitempty"`

	// generated from element "MD_GeometricObjectTypeCode" of complex type CodeListValue_Type
	MD_GeometricObjectTypeCode *gcobase.CodeListValue_Type `xml:"MD_GeometricObjectTypeCode,omitempty"`

	// generated from element "MD_CellGeometryCode" of complex type CodeListValue_Type
	MD_CellGeometryCode *gcobase.CodeListValue_Type `xml:"MD_CellGeometryCode,omitempty"`

	// generated from element "MD_DimensionNameTypeCode" of complex type CodeListValue_Type
	MD_DimensionNameTypeCode *gcobase.CodeListValue_Type `xml:"MD_DimensionNameTypeCode,omitempty"`

	// generated from element "CI_RoleCode" of complex type CodeListValue_Type
	CI_RoleCode *gcobase.CodeListValue_Type `xml:"CI_RoleCode,omitempty"`

	// generated from element "CI_PresentationFormCode" of complex type CodeListValue_Type
	CI_PresentationFormCode *gcobase.CodeListValue_Type `xml:"CI_PresentationFormCode,omitempty"`

	// generated from element "CI_OnLineFunctionCode" of complex type CodeListValue_Type
	CI_OnLineFunctionCode *gcobase.CodeListValue_Type `xml:"CI_OnLineFunctionCode,omitempty"`

	// generated from element "CI_DateTypeCode" of complex type CodeListValue_Type
	CI_DateTypeCode *gcobase.CodeListValue_Type `xml:"CI_DateTypeCode,omitempty"`

	// generated from element "MD_ObligationCode" of simple type gmd:MD_ObligationCode_Type
	MD_ObligationCode string `xml:"MD_ObligationCode,omitempty"`

	// generated from element "MD_DatatypeCode" of complex type CodeListValue_Type
	MD_DatatypeCode *gcobase.CodeListValue_Type `xml:"MD_DatatypeCode,omitempty"`

	// generated from element "MD_CoverageContentTypeCode" of complex type CodeListValue_Type
	MD_CoverageContentTypeCode *gcobase.CodeListValue_Type `xml:"MD_CoverageContentTypeCode,omitempty"`

	// generated from element "MD_ImagingConditionCode" of complex type CodeListValue_Type
	MD_ImagingConditionCode *gcobase.CodeListValue_Type `xml:"MD_ImagingConditionCode,omitempty"`

	// generated from element "DQ_EvaluationMethodTypeCode" of complex type CodeListValue_Type
	DQ_EvaluationMethodTypeCode *gcobase.CodeListValue_Type `xml:"DQ_EvaluationMethodTypeCode,omitempty"`

	// generated from element "MD_TopicCategoryCode" of simple type gmd:MD_TopicCategoryCode_Type
	MD_TopicCategoryCode Enum_MD_TopicCategoryCode_Type `xml:"MD_TopicCategoryCode,omitempty"`

	// generated from element "MD_CharacterSetCode" of complex type CodeListValue_Type
	MD_CharacterSetCode *gcobase.CodeListValue_Type `xml:"MD_CharacterSetCode,omitempty"`

	// generated from element "MD_SpatialRepresentationTypeCode" of complex type CodeListValue_Type
	MD_SpatialRepresentationTypeCode *gcobase.CodeListValue_Type `xml:"MD_SpatialRepresentationTypeCode,omitempty"`

	// generated from element "MD_ProgressCode" of complex type CodeListValue_Type
	MD_ProgressCode *gcobase.CodeListValue_Type `xml:"MD_ProgressCode,omitempty"`

	// generated from element "MD_KeywordTypeCode" of complex type CodeListValue_Type
	MD_KeywordTypeCode *gcobase.CodeListValue_Type `xml:"MD_KeywordTypeCode,omitempty"`

	// generated from element "DS_AssociationTypeCode" of complex type CodeListValue_Type
	DS_AssociationTypeCode *gcobase.CodeListValue_Type `xml:"DS_AssociationTypeCode,omitempty"`

	// generated from element "DS_InitiativeTypeCode" of complex type CodeListValue_Type
	DS_InitiativeTypeCode *gcobase.CodeListValue_Type `xml:"DS_InitiativeTypeCode,omitempty"`

	// generated from element "MD_ClassificationCode" of complex type CodeListValue_Type
	MD_ClassificationCode *gcobase.CodeListValue_Type `xml:"MD_ClassificationCode,omitempty"`

	// generated from element "MD_RestrictionCode" of complex type CodeListValue_Type
	MD_RestrictionCode *gcobase.CodeListValue_Type `xml:"MD_RestrictionCode,omitempty"`

	// generated from element "MD_DistributionUnits" of complex type CodeListValue_Type
	MD_DistributionUnits *gcobase.CodeListValue_Type `xml:"MD_DistributionUnits,omitempty"`

	// generated from element "MD_MediumFormatCode" of complex type CodeListValue_Type
	MD_MediumFormatCode *gcobase.CodeListValue_Type `xml:"MD_MediumFormatCode,omitempty"`

	// generated from element "MD_MediumNameCode" of complex type CodeListValue_Type
	MD_MediumNameCode *gcobase.CodeListValue_Type `xml:"MD_MediumNameCode,omitempty"`

	// generated from element "MD_MaintenanceFrequencyCode" of complex type CodeListValue_Type
	MD_MaintenanceFrequencyCode *gcobase.CodeListValue_Type `xml:"MD_MaintenanceFrequencyCode,omitempty"`

	// generated from element "MD_ScopeCode" of complex type CodeListValue_Type
	MD_ScopeCode *gcobase.CodeListValue_Type `xml:"MD_ScopeCode,omitempty"`

	// generated from element "LocalisedCharacterString" of complex type LocalisedCharacterString_Type
	LocalisedCharacterString *LocalisedCharacterString_Type `xml:"LocalisedCharacterString,omitempty"`

	// generated from element "LanguageCode" of complex type CodeListValue_Type
	LanguageCode *gcobase.CodeListValue_Type `xml:"LanguageCode,omitempty"`

	// generated from element "Country" of complex type CodeListValue_Type
	Country *gcobase.CodeListValue_Type `xml:"Country,omitempty"`

	// generated from element "PT_FreeText" of complex type PT_FreeText_Type
	PT_FreeText *PT_FreeText_Type `xml:"PT_FreeText,omitempty"`
}

// PT_Locale_Type Named source named complex type "PT_Locale_Type"
type PT_Locale_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "languageCode" of complex type LanguageCode_PropertyType
	LanguageCode *LanguageCode_PropertyType `xml:"languageCode,omitempty"`

	// generated from element "country" of complex type Country_PropertyType
	Country *Country_PropertyType `xml:"country,omitempty"`

	// generated from element "characterEncoding" of complex type MD_CharacterSetCode_PropertyType
	CharacterEncoding *MD_CharacterSetCode_PropertyType `xml:"characterEncoding,omitempty"`
}

// PT_Locale_PropertyType Named source named complex type "PT_Locale_PropertyType"
type PT_Locale_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "PT_Locale" of complex type PT_Locale_Type
	PT_Locale *PT_Locale_Type `xml:"PT_Locale,omitempty"`
}

// LocalisedCharacterString_Type Named source named complex type "LocalisedCharacterString_Type"
type LocalisedCharacterString_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "locale"
	Locale string `xml:"locale,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// LocalisedCharacterString_PropertyType Named source named complex type "LocalisedCharacterString_PropertyType"
type LocalisedCharacterString_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "LocalisedCharacterString" of complex type LocalisedCharacterString_Type
	LocalisedCharacterString *LocalisedCharacterString_Type `xml:"LocalisedCharacterString,omitempty"`
}

// PT_LocaleContainer_Type Named source named complex type "PT_LocaleContainer_Type"
type PT_LocaleContainer_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "description" of complex type CharacterString_PropertyType
	Description *CharacterString_PropertyType `xml:"description,omitempty"`

	// generated from element "locale" of complex type PT_Locale_PropertyType
	Locale *PT_Locale_PropertyType `xml:"locale,omitempty"`

	// generated from element "date" of complex type CI_Date_PropertyType
	Date []*CI_Date_PropertyType `xml:"date,omitempty"`

	// generated from element "responsibleParty" of complex type CI_ResponsibleParty_PropertyType
	ResponsibleParty []*CI_ResponsibleParty_PropertyType `xml:"responsibleParty,omitempty"`

	// generated from element "localisedString" of complex type LocalisedCharacterString_PropertyType
	LocalisedString []*LocalisedCharacterString_PropertyType `xml:"localisedString,omitempty"`
}

// PT_LocaleContainer_PropertyType Named source named complex type "PT_LocaleContainer_PropertyType"
type PT_LocaleContainer_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "PT_LocaleContainer" of complex type PT_LocaleContainer_Type
	PT_LocaleContainer *PT_LocaleContainer_Type `xml:"PT_LocaleContainer,omitempty"`
}

// LanguageCode_PropertyType Named source named complex type "LanguageCode_PropertyType"
type LanguageCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "LanguageCode" of complex type CodeListValue_Type
	LanguageCode *gcobase.CodeListValue_Type `xml:"LanguageCode,omitempty"`
}

// Country_PropertyType Named source named complex type "Country_PropertyType"
type Country_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "Country" of complex type CodeListValue_Type
	Country *gcobase.CodeListValue_Type `xml:"Country,omitempty"`
}

// AbstractDatumType Named source named complex type "AbstractDatumType"
type AbstractDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`
}

// DatumPropertyType Named source named complex type "DatumPropertyType"
// gml:DatumPropertyType is a property type for association roles to a datum, either referencing or containing the definition of that datum.
type DatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeodeticDatum" of complex type GeodeticDatumType
	GeodeticDatum *GeodeticDatumType `xml:"GeodeticDatum,omitempty"`

	// generated from element "EngineeringDatum" of complex type EngineeringDatumType
	EngineeringDatum *EngineeringDatumType `xml:"EngineeringDatum,omitempty"`

	// generated from element "ImageDatum" of complex type ImageDatumType
	ImageDatum *ImageDatumType `xml:"ImageDatum,omitempty"`

	// generated from element "VerticalDatum" of complex type VerticalDatumType
	VerticalDatum *VerticalDatumType `xml:"VerticalDatum,omitempty"`

	// generated from element "TemporalDatum" of complex type TemporalDatumType
	TemporalDatum *TemporalDatumType `xml:"TemporalDatum,omitempty"`
}

// GeodeticDatumType Named source named complex type "GeodeticDatumType"
type GeodeticDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`

	// generated from element "primeMeridian" of complex type PrimeMeridianPropertyType
	PrimeMeridian *PrimeMeridianPropertyType `xml:"primeMeridian,omitempty"`

	// generated from element "usesPrimeMeridian" of complex type PrimeMeridianPropertyType
	UsesPrimeMeridian *PrimeMeridianPropertyType `xml:"usesPrimeMeridian,omitempty"`

	// generated from element "ellipsoid" of complex type EllipsoidPropertyType
	Ellipsoid *EllipsoidPropertyType `xml:"ellipsoid,omitempty"`

	// generated from element "usesEllipsoid" of complex type EllipsoidPropertyType
	UsesEllipsoid *EllipsoidPropertyType `xml:"usesEllipsoid,omitempty"`
}

// GeodeticDatumPropertyType Named source named complex type "GeodeticDatumPropertyType"
// gml:GeodeticDatumPropertyType is a property type for association roles to a geodetic datum, either referencing or containing the definition of that datum.
type GeodeticDatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeodeticDatum" of complex type GeodeticDatumType
	GeodeticDatum *GeodeticDatumType `xml:"GeodeticDatum,omitempty"`
}

// EllipsoidType Named source named complex type "EllipsoidType"
type EllipsoidType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "semiMajorAxis" of complex type MeasureType
	SemiMajorAxis *basictypes.MeasureType `xml:"semiMajorAxis,omitempty"`

	// generated from anonymous complex type within element "secondDefiningParameter"
	SecondDefiningParameter *SecondDefiningParameter `xml:"secondDefiningParameter,omitempty"`
}

// SecondDefiningParameter Named source element secondDefiningParameter within schema
// gml:secondDefiningParameter is a property containing the definition of the second parameter that defines the shape of an ellipsoid. An ellipsoid requires two defining parameters: semi-major axis and inverse flattening or semi-major axis and semi-minor axis. When the reference body is a sphere rather than an ellipsoid, only a single defining parameter is required, namely the radius of the sphere; in that case, the semi-major axis "degenerates" into the radius of the sphere.
// The inverseFlattening element contains the inverse flattening value of the ellipsoid. This value is a scale factor (or ratio). It uses gml:LengthType with the restriction that the unit of measure referenced by the uom attribute must be suitable for a scale factor, such as percent, permil, or parts-per-million.
// The semiMinorAxis element contains the length of the semi-minor axis of the ellipsoid. When the isSphere element is included, the ellipsoid is degenerate and is actually a sphere. The sphere is completely defined by the semi-major axis, which is the radius of the sphere.
type SecondDefiningParameter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from anonymous complex type within element "SecondDefiningParameter"
	SecondDefiningParameter *SecondDefiningParameter_1 `xml:"SecondDefiningParameter,omitempty"`
}

// SecondDefiningParameter_1 Named source element SecondDefiningParameter within schema
type SecondDefiningParameter_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "inverseFlattening" of complex type MeasureType
	InverseFlattening *basictypes.MeasureType `xml:"inverseFlattening,omitempty"`

	// generated from element "semiMinorAxis" of complex type LengthType
	SemiMinorAxis *measures.LengthType `xml:"semiMinorAxis,omitempty"`

	// generated from element "isSphere" of simple type boolean
	IsSphere bool `xml:"isSphere,omitempty"`
}

// EllipsoidPropertyType Named source named complex type "EllipsoidPropertyType"
// gml:EllipsoidPropertyType is a property type for association roles to an ellipsoid, either referencing or containing the definition of that ellipsoid.
type EllipsoidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Ellipsoid" of complex type EllipsoidType
	Ellipsoid *EllipsoidType `xml:"Ellipsoid,omitempty"`
}

// PrimeMeridianType Named source named complex type "PrimeMeridianType"
type PrimeMeridianType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "greenwichLongitude" of complex type AngleType
	GreenwichLongitude *measures.AngleType `xml:"greenwichLongitude,omitempty"`
}

// PrimeMeridianPropertyType Named source named complex type "PrimeMeridianPropertyType"
// gml:PrimeMeridianPropertyType is a property type for association roles to a prime meridian, either referencing or containing the definition of that meridian.
type PrimeMeridianPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "PrimeMeridian" of complex type PrimeMeridianType
	PrimeMeridian *PrimeMeridianType `xml:"PrimeMeridian,omitempty"`
}

// EngineeringDatumType Named source named complex type "EngineeringDatumType"
type EngineeringDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`
}

// EngineeringDatumPropertyType Named source named complex type "EngineeringDatumPropertyType"
// gml:EngineeringDatumPropertyType is a property type for association roles to an engineering datum, either referencing or containing the definition of that datum.
type EngineeringDatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "EngineeringDatum" of complex type EngineeringDatumType
	EngineeringDatum *EngineeringDatumType `xml:"EngineeringDatum,omitempty"`
}

// ImageDatumType Named source named complex type "ImageDatumType"
type ImageDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`

	// generated from element "pixelInCell" of complex type CodeWithAuthorityType
	PixelInCell *basictypes.CodeWithAuthorityType `xml:"pixelInCell,omitempty"`
}

// ImageDatumPropertyType Named source named complex type "ImageDatumPropertyType"
// gml:ImageDatumPropertyType is a property type for association roles to an image datum, either referencing or containing the definition of that datum.
type ImageDatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ImageDatum" of complex type ImageDatumType
	ImageDatum *ImageDatumType `xml:"ImageDatum,omitempty"`
}

// VerticalDatumType Named source named complex type "VerticalDatumType"
type VerticalDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`
}

// VerticalDatumPropertyType Named source named complex type "VerticalDatumPropertyType"
// gml:VerticalDatumPropertyType is property type for association roles to a vertical datum, either referencing or containing the definition of that datum.
type VerticalDatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "VerticalDatum" of complex type VerticalDatumType
	VerticalDatum *VerticalDatumType `xml:"VerticalDatum,omitempty"`
}

// TemporalDatumType Named source named complex type "TemporalDatumType"
type TemporalDatumType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "anchorDefinition" of complex type CodeType
	AnchorDefinition *basictypes.CodeType `xml:"anchorDefinition,omitempty"`

	// generated from element "anchorPoint" of complex type CodeType
	AnchorPoint *basictypes.CodeType `xml:"anchorPoint,omitempty"`

	// generated from element "realizationEpoch" of simple type date
	RealizationEpoch string `xml:"realizationEpoch,omitempty"`

	// generated from element "origin" of simple type dateTime
	Origin string `xml:"origin,omitempty"`
}

// TemporalDatumBaseType Named source named complex type "TemporalDatumBaseType"
// The TemporalDatumBaseType partially defines the origin of a temporal coordinate reference system. This type restricts the AbstractDatumType to remove the "anchorDefinition" and "realizationEpoch" elements.
type TemporalDatumBaseType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`
}

// TemporalDatumPropertyType Named source named complex type "TemporalDatumPropertyType"
// gml:TemporalDatumPropertyType is a property type for association roles to a temporal datum, either referencing or containing the definition of that datum.
type TemporalDatumPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TemporalDatum" of complex type TemporalDatumType
	TemporalDatum *TemporalDatumType `xml:"TemporalDatum,omitempty"`
}

// AbstractCoordinateOperationType Named source named complex type "AbstractCoordinateOperationType"
type AbstractCoordinateOperationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`
}

// CoordinateOperationAccuracy Named source element coordinateOperationAccuracy within schema
// gml:coordinateOperationAccuracy is an association role to a DQ_PositionalAccuracy object as encoded in ISO/TS 19139, either referencing or containing the definition of that positional accuracy. That object contains an estimate of the impact of this coordinate operation on point accuracy. That is, it gives position error estimates for the target coordinates of this coordinate operation, assuming no errors in the source coordinates.
type CoordinateOperationAccuracy struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DQ_RelativeInternalPositionalAccuracy" of complex type DQ_RelativeInternalPositionalAccuracy_Type
	DQ_RelativeInternalPositionalAccuracy *DQ_RelativeInternalPositionalAccuracy_Type `xml:"DQ_RelativeInternalPositionalAccuracy,omitempty"`

	// generated from element "DQ_GriddedDataPositionalAccuracy" of complex type DQ_GriddedDataPositionalAccuracy_Type
	DQ_GriddedDataPositionalAccuracy *DQ_GriddedDataPositionalAccuracy_Type `xml:"DQ_GriddedDataPositionalAccuracy,omitempty"`

	// generated from element "DQ_AbsoluteExternalPositionalAccuracy" of complex type DQ_AbsoluteExternalPositionalAccuracy_Type
	DQ_AbsoluteExternalPositionalAccuracy *DQ_AbsoluteExternalPositionalAccuracy_Type `xml:"DQ_AbsoluteExternalPositionalAccuracy,omitempty"`
}

// CoordinateOperationPropertyType Named source named complex type "CoordinateOperationPropertyType"
// gml:CoordinateOperationPropertyType is a property type for association roles to a coordinate operation, either referencing or containing the definition of that coordinate operation.
type CoordinateOperationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "PassThroughOperation" of complex type PassThroughOperationType
	PassThroughOperation *PassThroughOperationType `xml:"PassThroughOperation,omitempty"`

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`

	// generated from element "ConcatenatedOperation" of complex type ConcatenatedOperationType
	ConcatenatedOperation *ConcatenatedOperationType `xml:"ConcatenatedOperation,omitempty"`
}

// SingleOperationPropertyType Named source named complex type "SingleOperationPropertyType"
// gml:SingleOperationPropertyType is a property type for association roles to a single operation, either referencing or containing the definition of that single operation.
type SingleOperationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "PassThroughOperation" of complex type PassThroughOperationType
	PassThroughOperation *PassThroughOperationType `xml:"PassThroughOperation,omitempty"`

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`
}

// AbstractGeneralConversionType Named source named complex type "AbstractGeneralConversionType"
type AbstractGeneralConversionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`
}

// GeneralConversionPropertyType Named source named complex type "GeneralConversionPropertyType"
// gml:GeneralConversionPropertyType is a property type for association roles to a general conversion, either referencing or containing the definition of that conversion.
type GeneralConversionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`
}

// AbstractGeneralTransformationType Named source named complex type "AbstractGeneralTransformationType"
type AbstractGeneralTransformationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`
}

// GeneralTransformationPropertyType Named source named complex type "GeneralTransformationPropertyType"
// gml:GeneralTransformationPropertyType is a property type for association roles to a general transformation, either referencing or containing the definition of that transformation.
type GeneralTransformationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`
}

// ConcatenatedOperationType Named source named complex type "ConcatenatedOperationType"
// gml:ConcatenatedOperation is an ordered sequence of two or more coordinate operations. This sequence of operations is constrained by the requirement that the source coordinate reference system of step (n+1) must be the same as the target coordinate reference system of step (n). The source coordinate reference system of the first step and the target coordinate reference system of the last step are the source and target coordinate reference system associated with the concatenated operation. Instead of a forward operation, an inverse operation may be used for one or more of the operation steps mentioned above, if the inverse operation is uniquely defined by the forward operation.
// The gml:coordOperation property elements are an ordered sequence of associations to the two or more operations used by this concatenated operation. The AggregationAttributeGroup should be used to specify that the coordOperation associations are ordered.
type ConcatenatedOperationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "coordOperation" of complex type CoordinateOperationPropertyType
	CoordOperation []*CoordinateOperationPropertyType `xml:"coordOperation,omitempty"`

	// generated from element "usesSingleOperation" of complex type CoordinateOperationPropertyType
	UsesSingleOperation []*CoordinateOperationPropertyType `xml:"usesSingleOperation,omitempty"`

	// generated from element "usesOperation" of complex type CoordinateOperationPropertyType
	UsesOperation []*CoordinateOperationPropertyType `xml:"usesOperation,omitempty"`
}

// ConcatenatedOperationPropertyType Named source named complex type "ConcatenatedOperationPropertyType"
// gml:ConcatenatedOperationPropertyType is a property type for association roles to a concatenated operation, either referencing or containing the definition of that concatenated operation.
type ConcatenatedOperationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ConcatenatedOperation" of complex type ConcatenatedOperationType
	ConcatenatedOperation *ConcatenatedOperationType `xml:"ConcatenatedOperation,omitempty"`
}

// PassThroughOperationType Named source named complex type "PassThroughOperationType"
type PassThroughOperationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "modifiedCoordinate" of simple type positiveInteger
	ModifiedCoordinate int `xml:"modifiedCoordinate,omitempty"`

	// generated from element "coordOperation" of complex type CoordinateOperationPropertyType
	CoordOperation *CoordinateOperationPropertyType `xml:"coordOperation,omitempty"`

	// generated from element "usesSingleOperation" of complex type CoordinateOperationPropertyType
	UsesSingleOperation *CoordinateOperationPropertyType `xml:"usesSingleOperation,omitempty"`

	// generated from element "usesOperation" of complex type CoordinateOperationPropertyType
	UsesOperation *CoordinateOperationPropertyType `xml:"usesOperation,omitempty"`
}

// PassThroughOperationPropertyType Named source named complex type "PassThroughOperationPropertyType"
// gml:PassThroughOperationPropertyType is a property type for association roles to a pass through operation, either referencing or containing the definition of that pass through operation.
type PassThroughOperationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "PassThroughOperation" of complex type PassThroughOperationType
	PassThroughOperation *PassThroughOperationType `xml:"PassThroughOperation,omitempty"`
}

// ConversionType Named source named complex type "ConversionType"
type ConversionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`

	// generated from element "method" of complex type OperationMethodPropertyType
	Method *OperationMethodPropertyType `xml:"method,omitempty"`

	// generated from element "usesMethod" of complex type OperationMethodPropertyType
	UsesMethod *OperationMethodPropertyType `xml:"usesMethod,omitempty"`

	// generated from element "parameterValue" of complex type AbstractGeneralParameterValuePropertyType
	ParameterValue []*AbstractGeneralParameterValuePropertyType `xml:"parameterValue,omitempty"`

	// generated from element "usesValue" of complex type AbstractGeneralParameterValuePropertyType
	UsesValue []*AbstractGeneralParameterValuePropertyType `xml:"usesValue,omitempty"`

	// generated from element "includesValue" of complex type AbstractGeneralParameterValuePropertyType
	IncludesValue []*AbstractGeneralParameterValuePropertyType `xml:"includesValue,omitempty"`
}

// ConversionPropertyType Named source named complex type "ConversionPropertyType"
// gml:ConversionPropertyType is a property type for association roles to a concrete general-purpose conversion, either referencing or containing the definition of that conversion.
type ConversionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`
}

// TransformationType Named source named complex type "TransformationType"
type TransformationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity *DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "operationVersion" of simple type string
	OperationVersion string `xml:"operationVersion,omitempty"`

	// generated from anonymous complex type within element "coordinateOperationAccuracy"
	CoordinateOperationAccuracy []*CoordinateOperationAccuracy `xml:"coordinateOperationAccuracy,omitempty"`

	// generated from element "sourceCRS" of complex type CRSPropertyType
	SourceCRS *CRSPropertyType `xml:"sourceCRS,omitempty"`

	// generated from element "targetCRS" of complex type CRSPropertyType
	TargetCRS *CRSPropertyType `xml:"targetCRS,omitempty"`

	// generated from element "method" of complex type OperationMethodPropertyType
	Method *OperationMethodPropertyType `xml:"method,omitempty"`

	// generated from element "usesMethod" of complex type OperationMethodPropertyType
	UsesMethod *OperationMethodPropertyType `xml:"usesMethod,omitempty"`

	// generated from element "parameterValue" of complex type AbstractGeneralParameterValuePropertyType
	ParameterValue []*AbstractGeneralParameterValuePropertyType `xml:"parameterValue,omitempty"`

	// generated from element "usesValue" of complex type AbstractGeneralParameterValuePropertyType
	UsesValue []*AbstractGeneralParameterValuePropertyType `xml:"usesValue,omitempty"`

	// generated from element "includesValue" of complex type AbstractGeneralParameterValuePropertyType
	IncludesValue []*AbstractGeneralParameterValuePropertyType `xml:"includesValue,omitempty"`
}

// TransformationPropertyType Named source named complex type "TransformationPropertyType"
// gml:TransformationPropertyType is a property type for association roles to a transformation, either referencing or containing the definition of that transformation.
type TransformationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`
}

// AbstractGeneralParameterValueType Named source named complex type "AbstractGeneralParameterValueType"
type AbstractGeneralParameterValueType struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// AbstractGeneralParameterValuePropertyType Named source named complex type "AbstractGeneralParameterValuePropertyType"
// gml:AbstractGeneralParameterValuePropertyType is a  property type for inline association roles to a parameter value or group of parameter values, always containing the values.
type AbstractGeneralParameterValuePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ParameterValue" of complex type ParameterValueType
	ParameterValue *ParameterValueType `xml:"ParameterValue,omitempty"`

	// generated from element "ParameterValueGroup" of complex type ParameterValueGroupType
	ParameterValueGroup *ParameterValueGroupType `xml:"ParameterValueGroup,omitempty"`
}

// ParameterValueType Named source named complex type "ParameterValueType"
type ParameterValueType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "value" of complex type MeasureType
	Value *basictypes.MeasureType `xml:"value,omitempty"`

	// generated from element "dmsAngleValue" of complex type DMSAngleType
	DmsAngleValue *DMSAngleType `xml:"dmsAngleValue,omitempty"`

	// generated from element "stringValue" of simple type string
	StringValue string `xml:"stringValue,omitempty"`

	// generated from element "integerValue" of simple type positiveInteger
	IntegerValue int `xml:"integerValue,omitempty"`

	// generated from element "booleanValue" of simple type boolean
	BooleanValue bool `xml:"booleanValue,omitempty"`

	// generated from element "valueList" of complex type MeasureListType
	ValueList *basictypes.MeasureListType `xml:"valueList,omitempty"`

	// generated from element "integerValueList" of simple type gml:integerList
	IntegerValueList string `xml:"integerValueList,omitempty"`

	// generated from element "valueFile" of simple type anyURI
	ValueFile string `xml:"valueFile,omitempty"`

	// generated from element "operationParameter" of complex type OperationParameterPropertyType
	OperationParameter *OperationParameterPropertyType `xml:"operationParameter,omitempty"`

	// generated from element "valueOfParameter" of complex type OperationParameterPropertyType
	ValueOfParameter *OperationParameterPropertyType `xml:"valueOfParameter,omitempty"`
}

// ParameterValueGroupType Named source named complex type "ParameterValueGroupType"
type ParameterValueGroupType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "parameterValue" of complex type AbstractGeneralParameterValuePropertyType
	ParameterValue []*AbstractGeneralParameterValuePropertyType `xml:"parameterValue,omitempty"`

	// generated from element "usesValue" of complex type AbstractGeneralParameterValuePropertyType
	UsesValue []*AbstractGeneralParameterValuePropertyType `xml:"usesValue,omitempty"`

	// generated from element "includesValue" of complex type AbstractGeneralParameterValuePropertyType
	IncludesValue []*AbstractGeneralParameterValuePropertyType `xml:"includesValue,omitempty"`

	// generated from element "group" of complex type OperationParameterGroupPropertyType
	Group *OperationParameterGroupPropertyType `xml:"group,omitempty"`

	// generated from element "valuesOfGroup" of complex type OperationParameterGroupPropertyType
	ValuesOfGroup *OperationParameterGroupPropertyType `xml:"valuesOfGroup,omitempty"`
}

// OperationMethodType Named source named complex type "OperationMethodType"
type OperationMethodType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "formulaCitation"
	FormulaCitation *FormulaCitation `xml:"formulaCitation,omitempty"`

	// generated from element "formula" of complex type CodeType
	Formula *basictypes.CodeType `xml:"formula,omitempty"`

	// generated from element "methodFormula" of complex type CodeType
	MethodFormula *basictypes.CodeType `xml:"methodFormula,omitempty"`

	// generated from element "sourceDimensions" of simple type positiveInteger
	SourceDimensions int `xml:"sourceDimensions,omitempty"`

	// generated from element "targetDimensions" of simple type positiveInteger
	TargetDimensions int `xml:"targetDimensions,omitempty"`

	// generated from element "parameter" of complex type AbstractGeneralOperationParameterPropertyType
	Parameter []*AbstractGeneralOperationParameterPropertyType `xml:"parameter,omitempty"`

	// generated from element "includesParameter" of complex type AbstractGeneralOperationParameterPropertyType
	IncludesParameter []*AbstractGeneralOperationParameterPropertyType `xml:"includesParameter,omitempty"`

	// generated from element "generalOperationParameter" of complex type AbstractGeneralOperationParameterPropertyType
	GeneralOperationParameter []*AbstractGeneralOperationParameterPropertyType `xml:"generalOperationParameter,omitempty"`

	// generated from element "usesParameter" of complex type AbstractGeneralOperationParameterPropertyType
	UsesParameter []*AbstractGeneralOperationParameterPropertyType `xml:"usesParameter,omitempty"`
}

// FormulaCitation Named source element formulaCitation within schema
// gml:formulaCitation provides a reference to a publication giving the formula(s) or procedure used by an coordinate operation method.
type FormulaCitation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "CI_Citation" of complex type CI_Citation_Type
	CI_Citation *CI_Citation_Type `xml:"CI_Citation,omitempty"`
}

// OperationMethodPropertyType Named source named complex type "OperationMethodPropertyType"
// gml:OperationMethodPropertyType is a property type for association roles to a concrete general-purpose operation method, either referencing or containing the definition of that method.
type OperationMethodPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "OperationMethod" of complex type OperationMethodType
	OperationMethod *OperationMethodType `xml:"OperationMethod,omitempty"`
}

// AbstractGeneralOperationParameterType Named source named complex type "AbstractGeneralOperationParameterType"
type AbstractGeneralOperationParameterType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "minimumOccurs" of simple type nonNegativeInteger
	MinimumOccurs int `xml:"minimumOccurs,omitempty"`
}

// AbstractGeneralOperationParameterPropertyType Named source named complex type "AbstractGeneralOperationParameterPropertyType"
// gml:AbstractGeneralOperationParameterPropertyType is a property type for association roles to an operation parameter or group, either referencing or containing the definition of that parameter or group.
type AbstractGeneralOperationParameterPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "OperationParameter" of complex type OperationParameterType
	OperationParameter *OperationParameterType `xml:"OperationParameter,omitempty"`

	// generated from element "OperationParameterGroup" of complex type OperationParameterGroupType
	OperationParameterGroup *OperationParameterGroupType `xml:"OperationParameterGroup,omitempty"`
}

// OperationParameterType Named source named complex type "OperationParameterType"
type OperationParameterType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "minimumOccurs" of simple type nonNegativeInteger
	MinimumOccurs int `xml:"minimumOccurs,omitempty"`
}

// OperationParameterPropertyType Named source named complex type "OperationParameterPropertyType"
// gml:OperationParameterPropertyType is a property type for association roles to an operation parameter, either referencing or containing the definition of that parameter.
type OperationParameterPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "OperationParameter" of complex type OperationParameterType
	OperationParameter *OperationParameterType `xml:"OperationParameter,omitempty"`
}

// OperationParameterGroupType Named source named complex type "OperationParameterGroupType"
type OperationParameterGroupType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "minimumOccurs" of simple type nonNegativeInteger
	MinimumOccurs int `xml:"minimumOccurs,omitempty"`

	// generated from element "maximumOccurs" of simple type positiveInteger
	MaximumOccurs int `xml:"maximumOccurs,omitempty"`

	// generated from element "parameter" of complex type AbstractGeneralOperationParameterPropertyType
	Parameter []*AbstractGeneralOperationParameterPropertyType `xml:"parameter,omitempty"`

	// generated from element "includesParameter" of complex type AbstractGeneralOperationParameterPropertyType
	IncludesParameter []*AbstractGeneralOperationParameterPropertyType `xml:"includesParameter,omitempty"`

	// generated from element "generalOperationParameter" of complex type AbstractGeneralOperationParameterPropertyType
	GeneralOperationParameter []*AbstractGeneralOperationParameterPropertyType `xml:"generalOperationParameter,omitempty"`

	// generated from element "usesParameter" of complex type AbstractGeneralOperationParameterPropertyType
	UsesParameter []*AbstractGeneralOperationParameterPropertyType `xml:"usesParameter,omitempty"`
}

// OperationParameterGroupPropertyType Named source named complex type "OperationParameterGroupPropertyType"
// gml:OperationParameterPropertyType is a property type for association roles to an operation parameter group, either referencing or containing the definition of that parameter group.
type OperationParameterGroupPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "OperationParameterGroup" of complex type OperationParameterGroupType
	OperationParameterGroup *OperationParameterGroupType `xml:"OperationParameterGroup,omitempty"`
}

// ObservationType Named source named complex type "ObservationType"
type ObservationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "using" of complex type ProcedurePropertyType
	Using *ProcedurePropertyType `xml:"using,omitempty"`

	// generated from element "target" of complex type TargetPropertyType
	Target *TargetPropertyType `xml:"target,omitempty"`

	// generated from element "subject" of complex type TargetPropertyType
	Subject *TargetPropertyType `xml:"subject,omitempty"`

	// generated from element "resultOf" of complex type ResultType
	ResultOf *ResultType `xml:"resultOf,omitempty"`
}

// ProcedurePropertyType Named source named complex type "ProcedurePropertyType"
type ProcedurePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature *DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection *DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`

	// generated from element "MultiPointCoverage" of complex type DiscreteCoverageType
	MultiPointCoverage *DiscreteCoverageType `xml:"MultiPointCoverage,omitempty"`

	// generated from element "MultiCurveCoverage" of complex type DiscreteCoverageType
	MultiCurveCoverage *DiscreteCoverageType `xml:"MultiCurveCoverage,omitempty"`

	// generated from element "MultiSurfaceCoverage" of complex type DiscreteCoverageType
	MultiSurfaceCoverage *DiscreteCoverageType `xml:"MultiSurfaceCoverage,omitempty"`

	// generated from element "MultiSolidCoverage" of complex type DiscreteCoverageType
	MultiSolidCoverage *DiscreteCoverageType `xml:"MultiSolidCoverage,omitempty"`

	// generated from element "GridCoverage" of complex type DiscreteCoverageType
	GridCoverage *DiscreteCoverageType `xml:"GridCoverage,omitempty"`

	// generated from element "RectifiedGridCoverage" of complex type DiscreteCoverageType
	RectifiedGridCoverage *DiscreteCoverageType `xml:"RectifiedGridCoverage,omitempty"`

	// generated from element "Observation" of complex type ObservationType
	Observation *ObservationType `xml:"Observation,omitempty"`

	// generated from element "DirectedObservation" of complex type DirectedObservationType
	DirectedObservation *DirectedObservationType `xml:"DirectedObservation,omitempty"`

	// generated from element "DirectedObservationAtDistance" of complex type DirectedObservationAtDistanceType
	DirectedObservationAtDistance *DirectedObservationAtDistanceType `xml:"DirectedObservationAtDistance,omitempty"`

	// generated from element "FeatureCollection" of complex type FeatureCollectionType
	FeatureCollection *FeatureCollectionType `xml:"FeatureCollection,omitempty"`
}

// TargetPropertyType Named source named complex type "TargetPropertyType"
type TargetPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature *DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection *DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`

	// generated from element "MultiPointCoverage" of complex type DiscreteCoverageType
	MultiPointCoverage *DiscreteCoverageType `xml:"MultiPointCoverage,omitempty"`

	// generated from element "MultiCurveCoverage" of complex type DiscreteCoverageType
	MultiCurveCoverage *DiscreteCoverageType `xml:"MultiCurveCoverage,omitempty"`

	// generated from element "MultiSurfaceCoverage" of complex type DiscreteCoverageType
	MultiSurfaceCoverage *DiscreteCoverageType `xml:"MultiSurfaceCoverage,omitempty"`

	// generated from element "MultiSolidCoverage" of complex type DiscreteCoverageType
	MultiSolidCoverage *DiscreteCoverageType `xml:"MultiSolidCoverage,omitempty"`

	// generated from element "GridCoverage" of complex type DiscreteCoverageType
	GridCoverage *DiscreteCoverageType `xml:"GridCoverage,omitempty"`

	// generated from element "RectifiedGridCoverage" of complex type DiscreteCoverageType
	RectifiedGridCoverage *DiscreteCoverageType `xml:"RectifiedGridCoverage,omitempty"`

	// generated from element "Observation" of complex type ObservationType
	Observation *ObservationType `xml:"Observation,omitempty"`

	// generated from element "DirectedObservation" of complex type DirectedObservationType
	DirectedObservation *DirectedObservationType `xml:"DirectedObservation,omitempty"`

	// generated from element "DirectedObservationAtDistance" of complex type DirectedObservationAtDistanceType
	DirectedObservationAtDistance *DirectedObservationAtDistanceType `xml:"DirectedObservationAtDistance,omitempty"`

	// generated from element "FeatureCollection" of complex type FeatureCollectionType
	FeatureCollection *FeatureCollectionType `xml:"FeatureCollection,omitempty"`

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`
}

// ResultType Named source named complex type "ResultType"
type ResultType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup
}

// DirectedObservationType Named source named complex type "DirectedObservationType"
type DirectedObservationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "using" of complex type ProcedurePropertyType
	Using *ProcedurePropertyType `xml:"using,omitempty"`

	// generated from element "target" of complex type TargetPropertyType
	Target *TargetPropertyType `xml:"target,omitempty"`

	// generated from element "subject" of complex type TargetPropertyType
	Subject *TargetPropertyType `xml:"subject,omitempty"`

	// generated from element "resultOf" of complex type ResultType
	ResultOf *ResultType `xml:"resultOf,omitempty"`

	// generated from element "direction" of complex type DirectionPropertyType
	Direction *DirectionPropertyType `xml:"direction,omitempty"`
}

// DirectedObservationAtDistanceType Named source named complex type "DirectedObservationAtDistanceType"
type DirectedObservationAtDistanceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "using" of complex type ProcedurePropertyType
	Using *ProcedurePropertyType `xml:"using,omitempty"`

	// generated from element "target" of complex type TargetPropertyType
	Target *TargetPropertyType `xml:"target,omitempty"`

	// generated from element "subject" of complex type TargetPropertyType
	Subject *TargetPropertyType `xml:"subject,omitempty"`

	// generated from element "resultOf" of complex type ResultType
	ResultOf *ResultType `xml:"resultOf,omitempty"`

	// generated from element "direction" of complex type DirectionPropertyType
	Direction *DirectionPropertyType `xml:"direction,omitempty"`

	// generated from element "distance" of complex type MeasureType
	Distance *basictypes.MeasureType `xml:"distance,omitempty"`
}

// TimeReferenceSystemType Named source named complex type "TimeReferenceSystemType"
type TimeReferenceSystemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "domainOfValidity" of simple type string
	DomainOfValidity string `xml:"domainOfValidity,omitempty"`
}

// TimeCoordinateSystemType Named source named complex type "TimeCoordinateSystemType"
type TimeCoordinateSystemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "domainOfValidity" of simple type string
	DomainOfValidity string `xml:"domainOfValidity,omitempty"`

	// generated from element "originPosition" of complex type TimePositionType
	OriginPosition *TimePositionType `xml:"originPosition,omitempty"`

	// generated from element "origin" of complex type TimeInstantPropertyType
	Origin *TimeInstantPropertyType `xml:"origin,omitempty"`

	// generated from element "interval" of complex type TimeIntervalLengthType
	Interval *TimeIntervalLengthType `xml:"interval,omitempty"`
}

// TimeCalendarType Named source named complex type "TimeCalendarType"
type TimeCalendarType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "domainOfValidity" of simple type string
	DomainOfValidity string `xml:"domainOfValidity,omitempty"`

	// generated from element "referenceFrame" of complex type TimeCalendarEraPropertyType
	ReferenceFrame []*TimeCalendarEraPropertyType `xml:"referenceFrame,omitempty"`
}

// TimeCalendarEraType Named source named complex type "TimeCalendarEraType"
type TimeCalendarEraType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "referenceEvent" of complex type StringOrRefType
	ReferenceEvent *StringOrRefType `xml:"referenceEvent,omitempty"`

	// generated from element "referenceDate" of simple type gml:CalDate
	ReferenceDate string `xml:"referenceDate,omitempty"`

	// generated from element "julianReference" of simple type decimal
	JulianReference string `xml:"julianReference,omitempty"`

	// generated from element "epochOfUse" of complex type TimePeriodPropertyType
	EpochOfUse *TimePeriodPropertyType `xml:"epochOfUse,omitempty"`
}

// TimeCalendarPropertyType Named source named complex type "TimeCalendarPropertyType"
// gml:TimeCalendarPropertyType provides for associating a gml:TimeCalendar with an object.
type TimeCalendarPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeCalendar" of complex type TimeCalendarType
	TimeCalendar *TimeCalendarType `xml:"TimeCalendar,omitempty"`
}

// TimeCalendarEraPropertyType Named source named complex type "TimeCalendarEraPropertyType"
// gml:TimeCalendarEraPropertyType provides for associating a gml:TimeCalendarEra with an object.
type TimeCalendarEraPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeCalendarEra" of complex type TimeCalendarEraType
	TimeCalendarEra *TimeCalendarEraType `xml:"TimeCalendarEra,omitempty"`
}

// TimeClockType Named source named complex type "TimeClockType"
type TimeClockType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "domainOfValidity" of simple type string
	DomainOfValidity string `xml:"domainOfValidity,omitempty"`

	// generated from element "referenceEvent" of complex type StringOrRefType
	ReferenceEvent *StringOrRefType `xml:"referenceEvent,omitempty"`

	// generated from element "referenceTime" of simple type time
	ReferenceTime string `xml:"referenceTime,omitempty"`

	// generated from element "utcReference" of simple type time
	UtcReference string `xml:"utcReference,omitempty"`

	// generated from element "dateBasis" of complex type TimeCalendarPropertyType
	DateBasis []*TimeCalendarPropertyType `xml:"dateBasis,omitempty"`
}

// TimeClockPropertyType Named source named complex type "TimeClockPropertyType"
// gml:TimeClockPropertyType provides for associating a gml:TimeClock with an object.
type TimeClockPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeClock" of complex type TimeClockType
	TimeClock *TimeClockType `xml:"TimeClock,omitempty"`
}

// TimeOrdinalReferenceSystemType Named source named complex type "TimeOrdinalReferenceSystemType"
type TimeOrdinalReferenceSystemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "domainOfValidity" of simple type string
	DomainOfValidity string `xml:"domainOfValidity,omitempty"`

	// generated from element "component" of complex type TimeOrdinalEraPropertyType
	Component []*TimeOrdinalEraPropertyType `xml:"component,omitempty"`
}

// TimeOrdinalEraType Named source named complex type "TimeOrdinalEraType"
type TimeOrdinalEraType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from element "start" of complex type TimeNodePropertyType
	Start *TimeNodePropertyType `xml:"start,omitempty"`

	// generated from element "end" of complex type TimeNodePropertyType
	End *TimeNodePropertyType `xml:"end,omitempty"`

	// generated from element "extent" of complex type TimePeriodPropertyType
	Extent *TimePeriodPropertyType `xml:"extent,omitempty"`

	// generated from element "member" of complex type TimeOrdinalEraPropertyType
	Member []*TimeOrdinalEraPropertyType `xml:"member,omitempty"`

	// generated from element "group" of complex type ReferenceType
	Group *ReferenceType `xml:"group,omitempty"`
}

// TimeOrdinalEraPropertyType Named source named complex type "TimeOrdinalEraPropertyType"
// gml:TimeOrdinalEraPropertyType provides for associating a gml:TimeOrdinalEra with an object.
type TimeOrdinalEraPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeOrdinalEra" of complex type TimeOrdinalEraType
	TimeOrdinalEra *TimeOrdinalEraType `xml:"TimeOrdinalEra,omitempty"`
}

// AbstractTimeTopologyPrimitiveType Named source named complex type "AbstractTimeTopologyPrimitiveType"
type AbstractTimeTopologyPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from element "complex" of complex type ReferenceType
	Complex_XSD *ReferenceType `xml:"complex,omitempty"`
}

// TimeTopologyPrimitivePropertyType Named source named complex type "TimeTopologyPrimitivePropertyType"
// gml:TimeTopologyPrimitivePropertyType provides for associating a gml:AbstractTimeTopologyPrimitive with an object.
type TimeTopologyPrimitivePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`
}

// TimeTopologyComplexType Named source named complex type "TimeTopologyComplexType"
type TimeTopologyComplexType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "primitive" of complex type TimeTopologyPrimitivePropertyType
	Primitive []*TimeTopologyPrimitivePropertyType `xml:"primitive,omitempty"`
}

// TimeTopologyComplexPropertyType Named source named complex type "TimeTopologyComplexPropertyType"
// gml:TimeTopologyComplexPropertyType provides for associating a gml:TimeTopologyComplex with an object.
type TimeTopologyComplexPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeTopologyComplex" of complex type TimeTopologyComplexType
	TimeTopologyComplex *TimeTopologyComplexType `xml:"TimeTopologyComplex,omitempty"`
}

// TimeNodeType Named source named complex type "TimeNodeType"
type TimeNodeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from element "complex" of complex type ReferenceType
	Complex_XSD *ReferenceType `xml:"complex,omitempty"`

	// generated from element "previousEdge" of complex type TimeEdgePropertyType
	PreviousEdge []*TimeEdgePropertyType `xml:"previousEdge,omitempty"`

	// generated from element "nextEdge" of complex type TimeEdgePropertyType
	NextEdge []*TimeEdgePropertyType `xml:"nextEdge,omitempty"`

	// generated from element "position" of complex type TimeInstantPropertyType
	Position *TimeInstantPropertyType `xml:"position,omitempty"`
}

// TimeNodePropertyType Named source named complex type "TimeNodePropertyType"
// gml:TimeNodePropertyType provides for associating a gml:TimeNode with an object
type TimeNodePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode *TimeNodeType `xml:"TimeNode,omitempty"`
}

// TimeEdgeType Named source named complex type "TimeEdgeType"
type TimeEdgeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "relatedTime" of complex type RelatedTimeType
	RelatedTime []*RelatedTimeType `xml:"relatedTime,omitempty"`

	// generated from element "complex" of complex type ReferenceType
	Complex_XSD *ReferenceType `xml:"complex,omitempty"`

	// generated from element "start" of complex type TimeNodePropertyType
	Start *TimeNodePropertyType `xml:"start,omitempty"`

	// generated from element "end" of complex type TimeNodePropertyType
	End *TimeNodePropertyType `xml:"end,omitempty"`

	// generated from element "extent" of complex type TimePeriodPropertyType
	Extent *TimePeriodPropertyType `xml:"extent,omitempty"`
}

// TimeEdgePropertyType Named source named complex type "TimeEdgePropertyType"
// gml:TimeEdgePropertyType provides for associating a gml:TimeEdge with an object.
type TimeEdgePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge *TimeEdgeType `xml:"TimeEdge,omitempty"`
}

// OperationPropertyType Named source named complex type "OperationPropertyType"
type OperationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Conversion" of complex type ConversionType
	Conversion *ConversionType `xml:"Conversion,omitempty"`

	// generated from element "Transformation" of complex type TransformationType
	Transformation *TransformationType `xml:"Transformation,omitempty"`
}

// TemporalCSType Named source named complex type "TemporalCSType"
type TemporalCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// TemporalCSPropertyType Named source named complex type "TemporalCSPropertyType"
type TemporalCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TemporalCS" of complex type TemporalCSType
	TemporalCS *TemporalCSType `xml:"TemporalCS,omitempty"`
}

// ObliqueCartesianCSType Named source named complex type "ObliqueCartesianCSType"
type ObliqueCartesianCSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "axis" of complex type CoordinateSystemAxisPropertyType
	Axis []*CoordinateSystemAxisPropertyType `xml:"axis,omitempty"`

	// generated from element "usesAxis" of complex type CoordinateSystemAxisPropertyType
	UsesAxis []*CoordinateSystemAxisPropertyType `xml:"usesAxis,omitempty"`
}

// ObliqueCartesianCSPropertyType Named source named complex type "ObliqueCartesianCSPropertyType"
type ObliqueCartesianCSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "ObliqueCartesianCS" of complex type ObliqueCartesianCSType
	ObliqueCartesianCS *ObliqueCartesianCSType `xml:"ObliqueCartesianCS,omitempty"`
}

// GeographicCRSType Named source named complex type "GeographicCRSType"
type GeographicCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "usesEllipsoidalCS" of complex type EllipsoidalCSPropertyType
	UsesEllipsoidalCS *EllipsoidalCSPropertyType `xml:"usesEllipsoidalCS,omitempty"`

	// generated from element "usesGeodeticDatum" of complex type GeodeticDatumPropertyType
	UsesGeodeticDatum *GeodeticDatumPropertyType `xml:"usesGeodeticDatum,omitempty"`
}

// GeographicCRSPropertyType Named source named complex type "GeographicCRSPropertyType"
type GeographicCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS *GeographicCRSType `xml:"GeographicCRS,omitempty"`
}

// GeocentricCRSType Named source named complex type "GeocentricCRSType"
type GeocentricCRSType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from anonymous complex type within element "domainOfValidity"
	DomainOfValidity []*DomainOfValidity `xml:"domainOfValidity,omitempty"`

	// generated from element "scope" of simple type string
	Scope string `xml:"scope,omitempty"`

	// generated from element "usesCartesianCS" of complex type CartesianCSPropertyType
	UsesCartesianCS *CartesianCSPropertyType `xml:"usesCartesianCS,omitempty"`

	// generated from element "usesSphericalCS" of complex type SphericalCSPropertyType
	UsesSphericalCS *SphericalCSPropertyType `xml:"usesSphericalCS,omitempty"`

	// generated from element "usesGeodeticDatum" of complex type GeodeticDatumPropertyType
	UsesGeodeticDatum *GeodeticDatumPropertyType `xml:"usesGeodeticDatum,omitempty"`
}

// GeocentricCRSPropertyType Named source named complex type "GeocentricCRSPropertyType"
type GeocentricCRSPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS *GeocentricCRSType `xml:"GeocentricCRS,omitempty"`
}

// Attr_uom UnNamed source top-level attribute "uom"
type Attr_uom struct {

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`
}

// DMSAngleType Named source named complex type "DMSAngleType"
type DMSAngleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "degrees" of complex type DegreesType
	Degrees *DegreesType `xml:"degrees,omitempty"`

	// generated from element "decimalMinutes" of simple type gml:DecimalMinutesType
	DecimalMinutes string `xml:"decimalMinutes,omitempty"`

	// generated from element "minutes" of simple type gml:ArcMinutesType
	Minutes int `xml:"minutes,omitempty"`

	// generated from element "seconds" of simple type gml:ArcSecondsType
	Seconds string `xml:"seconds,omitempty"`
}

// DegreesType Named source named complex type "DegreesType"
type DegreesType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "direction"
	Direction Enum_DegreesType_direction `xml:"direction,attr,omitempty"`

	// simpleContent character data
	EnclosedText int `xml:",chardata"`
}

// AngleChoiceType Named source named complex type "AngleChoiceType"
type AngleChoiceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "angle" of complex type AngleType
	Angle *measures.AngleType `xml:"angle,omitempty"`

	// generated from element "dmsAngle" of complex type DMSAngleType
	DmsAngle *DMSAngleType `xml:"dmsAngle,omitempty"`
}

// Attr_remoteSchema UnNamed source top-level attribute "remoteSchema"
type Attr_remoteSchema struct {

	// insertion point for fields

	// generated from attribute "remoteSchema"
	RemoteSchema string `xml:"remoteSchema,attr,omitempty"`
}

// ArrayAssociationType Named source named complex type "ArrayAssociationType"
type ArrayAssociationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	AttributeGroup_OwnershipAttributeGroup

	// generated from element "LineStringSegment" of complex type LineStringSegmentType
	LineStringSegment []*LineStringSegmentType `xml:"LineStringSegment,omitempty"`

	// generated from element "ArcString" of complex type ArcStringType
	ArcString []*ArcStringType `xml:"ArcString,omitempty"`

	// generated from element "Arc" of complex type ArcType
	Arc []*ArcType `xml:"Arc,omitempty"`

	// generated from element "Circle" of complex type CircleType
	Circle []*CircleType `xml:"Circle,omitempty"`

	// generated from element "ArcStringByBulge" of complex type ArcStringByBulgeType
	ArcStringByBulge []*ArcStringByBulgeType `xml:"ArcStringByBulge,omitempty"`

	// generated from element "ArcByBulge" of complex type ArcByBulgeType
	ArcByBulge []*ArcByBulgeType `xml:"ArcByBulge,omitempty"`

	// generated from element "ArcByCenterPoint" of complex type ArcByCenterPointType
	ArcByCenterPoint []*ArcByCenterPointType `xml:"ArcByCenterPoint,omitempty"`

	// generated from element "CircleByCenterPoint" of complex type CircleByCenterPointType
	CircleByCenterPoint []*CircleByCenterPointType `xml:"CircleByCenterPoint,omitempty"`

	// generated from element "CubicSpline" of complex type CubicSplineType
	CubicSpline []*CubicSplineType `xml:"CubicSpline,omitempty"`

	// generated from element "BSpline" of complex type BSplineType
	BSpline []*BSplineType `xml:"BSpline,omitempty"`

	// generated from element "Bezier" of complex type BezierType
	Bezier []*BezierType `xml:"Bezier,omitempty"`

	// generated from element "OffsetCurve" of complex type OffsetCurveType
	OffsetCurve []*OffsetCurveType `xml:"OffsetCurve,omitempty"`

	// generated from element "Clothoid" of complex type ClothoidType
	Clothoid []*ClothoidType `xml:"Clothoid,omitempty"`

	// generated from element "GeodesicString" of complex type GeodesicStringType
	GeodesicString []*GeodesicStringType `xml:"GeodesicString,omitempty"`

	// generated from element "Geodesic" of complex type GeodesicType
	Geodesic []*GeodesicType `xml:"Geodesic,omitempty"`

	// generated from element "AffinePlacement" of complex type AffinePlacementType
	AffinePlacement []*AffinePlacementType `xml:"AffinePlacement,omitempty"`

	// generated from element "Envelope" of complex type EnvelopeType
	Envelope []*EnvelopeType `xml:"Envelope,omitempty"`

	// generated from element "EnvelopeWithTimePeriod" of complex type EnvelopeWithTimePeriodType
	EnvelopeWithTimePeriod []*EnvelopeWithTimePeriodType `xml:"EnvelopeWithTimePeriod,omitempty"`

	// generated from element "MovingObjectStatus" of complex type MovingObjectStatusType
	MovingObjectStatus []*MovingObjectStatusType `xml:"MovingObjectStatus,omitempty"`

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature []*DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection []*DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`

	// generated from element "MultiPointCoverage" of complex type DiscreteCoverageType
	MultiPointCoverage []*DiscreteCoverageType `xml:"MultiPointCoverage,omitempty"`

	// generated from element "MultiCurveCoverage" of complex type DiscreteCoverageType
	MultiCurveCoverage []*DiscreteCoverageType `xml:"MultiCurveCoverage,omitempty"`

	// generated from element "MultiSurfaceCoverage" of complex type DiscreteCoverageType
	MultiSurfaceCoverage []*DiscreteCoverageType `xml:"MultiSurfaceCoverage,omitempty"`

	// generated from element "MultiSolidCoverage" of complex type DiscreteCoverageType
	MultiSolidCoverage []*DiscreteCoverageType `xml:"MultiSolidCoverage,omitempty"`

	// generated from element "GridCoverage" of complex type DiscreteCoverageType
	GridCoverage []*DiscreteCoverageType `xml:"GridCoverage,omitempty"`

	// generated from element "RectifiedGridCoverage" of complex type DiscreteCoverageType
	RectifiedGridCoverage []*DiscreteCoverageType `xml:"RectifiedGridCoverage,omitempty"`

	// generated from element "Observation" of complex type ObservationType
	Observation []*ObservationType `xml:"Observation,omitempty"`

	// generated from element "DirectedObservation" of complex type DirectedObservationType
	DirectedObservation []*DirectedObservationType `xml:"DirectedObservation,omitempty"`

	// generated from element "DirectedObservationAtDistance" of complex type DirectedObservationAtDistanceType
	DirectedObservationAtDistance []*DirectedObservationAtDistanceType `xml:"DirectedObservationAtDistance,omitempty"`

	// generated from element "FeatureCollection" of complex type FeatureCollectionType
	FeatureCollection []*FeatureCollectionType `xml:"FeatureCollection,omitempty"`

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry []*MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint []*MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve []*MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface []*MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid []*MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid []*SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid []*CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface []*SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface []*SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface []*SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin []*TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface []*OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell []*ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon []*PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface []*CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point []*PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve []*CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve []*OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring []*RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing []*LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString []*LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve []*CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex []*GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid []*GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid []*RectifiedGridType `xml:"RectifiedGrid,omitempty"`

	// generated from element "Definition" of complex type DefinitionType
	Definition []*DefinitionType `xml:"Definition,omitempty"`

	// generated from element "UnitDefinition" of complex type UnitDefinitionType
	UnitDefinition []*UnitDefinitionType `xml:"UnitDefinition,omitempty"`

	// generated from element "BaseUnit" of complex type BaseUnitType
	BaseUnit []*BaseUnitType `xml:"BaseUnit,omitempty"`

	// generated from element "DerivedUnit" of complex type DerivedUnitType
	DerivedUnit []*DerivedUnitType `xml:"DerivedUnit,omitempty"`

	// generated from element "ConventionalUnit" of complex type ConventionalUnitType
	ConventionalUnit []*ConventionalUnitType `xml:"ConventionalUnit,omitempty"`

	// generated from element "Dictionary" of complex type DictionaryType
	Dictionary []*DictionaryType `xml:"Dictionary,omitempty"`

	// generated from element "CoordinateSystemAxis" of complex type CoordinateSystemAxisType
	CoordinateSystemAxis []*CoordinateSystemAxisType `xml:"CoordinateSystemAxis,omitempty"`

	// generated from element "EllipsoidalCS" of complex type EllipsoidalCSType
	EllipsoidalCS []*EllipsoidalCSType `xml:"EllipsoidalCS,omitempty"`

	// generated from element "CartesianCS" of complex type CartesianCSType
	CartesianCS []*CartesianCSType `xml:"CartesianCS,omitempty"`

	// generated from element "VerticalCS" of complex type VerticalCSType
	VerticalCS []*VerticalCSType `xml:"VerticalCS,omitempty"`

	// generated from element "TimeCS" of complex type TimeCSType
	TimeCS []*TimeCSType `xml:"TimeCS,omitempty"`

	// generated from element "LinearCS" of complex type LinearCSType
	LinearCS []*LinearCSType `xml:"LinearCS,omitempty"`

	// generated from element "UserDefinedCS" of complex type UserDefinedCSType
	UserDefinedCS []*UserDefinedCSType `xml:"UserDefinedCS,omitempty"`

	// generated from element "SphericalCS" of complex type SphericalCSType
	SphericalCS []*SphericalCSType `xml:"SphericalCS,omitempty"`

	// generated from element "PolarCS" of complex type PolarCSType
	PolarCS []*PolarCSType `xml:"PolarCS,omitempty"`

	// generated from element "CylindricalCS" of complex type CylindricalCSType
	CylindricalCS []*CylindricalCSType `xml:"CylindricalCS,omitempty"`

	// generated from element "AffineCS" of complex type AffineCSType
	AffineCS []*AffineCSType `xml:"AffineCS,omitempty"`

	// generated from element "TemporalCS" of complex type TemporalCSType
	TemporalCS []*TemporalCSType `xml:"TemporalCS,omitempty"`

	// generated from element "ObliqueCartesianCS" of complex type ObliqueCartesianCSType
	ObliqueCartesianCS []*ObliqueCartesianCSType `xml:"ObliqueCartesianCS,omitempty"`

	// generated from element "ProjectedCRS" of complex type ProjectedCRSType
	ProjectedCRS []*ProjectedCRSType `xml:"ProjectedCRS,omitempty"`

	// generated from element "DerivedCRS" of complex type DerivedCRSType
	DerivedCRS []*DerivedCRSType `xml:"DerivedCRS,omitempty"`

	// generated from element "GeodeticCRS" of complex type GeodeticCRSType
	GeodeticCRS []*GeodeticCRSType `xml:"GeodeticCRS,omitempty"`

	// generated from element "VerticalCRS" of complex type VerticalCRSType
	VerticalCRS []*VerticalCRSType `xml:"VerticalCRS,omitempty"`

	// generated from element "EngineeringCRS" of complex type EngineeringCRSType
	EngineeringCRS []*EngineeringCRSType `xml:"EngineeringCRS,omitempty"`

	// generated from element "ImageCRS" of complex type ImageCRSType
	ImageCRS []*ImageCRSType `xml:"ImageCRS,omitempty"`

	// generated from element "TemporalCRS" of complex type TemporalCRSType
	TemporalCRS []*TemporalCRSType `xml:"TemporalCRS,omitempty"`

	// generated from element "GeographicCRS" of complex type GeographicCRSType
	GeographicCRS []*GeographicCRSType `xml:"GeographicCRS,omitempty"`

	// generated from element "GeocentricCRS" of complex type GeocentricCRSType
	GeocentricCRS []*GeocentricCRSType `xml:"GeocentricCRS,omitempty"`

	// generated from element "CompoundCRS" of complex type CompoundCRSType
	CompoundCRS []*CompoundCRSType `xml:"CompoundCRS,omitempty"`

	// generated from element "GeodeticDatum" of complex type GeodeticDatumType
	GeodeticDatum []*GeodeticDatumType `xml:"GeodeticDatum,omitempty"`

	// generated from element "EngineeringDatum" of complex type EngineeringDatumType
	EngineeringDatum []*EngineeringDatumType `xml:"EngineeringDatum,omitempty"`

	// generated from element "ImageDatum" of complex type ImageDatumType
	ImageDatum []*ImageDatumType `xml:"ImageDatum,omitempty"`

	// generated from element "VerticalDatum" of complex type VerticalDatumType
	VerticalDatum []*VerticalDatumType `xml:"VerticalDatum,omitempty"`

	// generated from element "TemporalDatum" of complex type TemporalDatumType
	TemporalDatum []*TemporalDatumType `xml:"TemporalDatum,omitempty"`

	// generated from element "Ellipsoid" of complex type EllipsoidType
	Ellipsoid []*EllipsoidType `xml:"Ellipsoid,omitempty"`

	// generated from element "PrimeMeridian" of complex type PrimeMeridianType
	PrimeMeridian []*PrimeMeridianType `xml:"PrimeMeridian,omitempty"`

	// generated from element "PassThroughOperation" of complex type PassThroughOperationType
	PassThroughOperation []*PassThroughOperationType `xml:"PassThroughOperation,omitempty"`

	// generated from element "Conversion" of complex type ConversionType
	Conversion []*ConversionType `xml:"Conversion,omitempty"`

	// generated from element "Transformation" of complex type TransformationType
	Transformation []*TransformationType `xml:"Transformation,omitempty"`

	// generated from element "ConcatenatedOperation" of complex type ConcatenatedOperationType
	ConcatenatedOperation []*ConcatenatedOperationType `xml:"ConcatenatedOperation,omitempty"`

	// generated from element "OperationMethod" of complex type OperationMethodType
	OperationMethod []*OperationMethodType `xml:"OperationMethod,omitempty"`

	// generated from element "OperationParameter" of complex type OperationParameterType
	OperationParameter []*OperationParameterType `xml:"OperationParameter,omitempty"`

	// generated from element "OperationParameterGroup" of complex type OperationParameterGroupType
	OperationParameterGroup []*OperationParameterGroupType `xml:"OperationParameterGroup,omitempty"`

	// generated from element "TimeReferenceSystem" of complex type TimeReferenceSystemType
	TimeReferenceSystem []*TimeReferenceSystemType `xml:"TimeReferenceSystem,omitempty"`

	// generated from element "TimeCoordinateSystem" of complex type TimeCoordinateSystemType
	TimeCoordinateSystem []*TimeCoordinateSystemType `xml:"TimeCoordinateSystem,omitempty"`

	// generated from element "TimeCalendar" of complex type TimeCalendarType
	TimeCalendar []*TimeCalendarType `xml:"TimeCalendar,omitempty"`

	// generated from element "TimeClock" of complex type TimeClockType
	TimeClock []*TimeClockType `xml:"TimeClock,omitempty"`

	// generated from element "TimeOrdinalReferenceSystem" of complex type TimeOrdinalReferenceSystemType
	TimeOrdinalReferenceSystem []*TimeOrdinalReferenceSystemType `xml:"TimeOrdinalReferenceSystem,omitempty"`

	// generated from element "DefinitionCollection" of complex type DictionaryType
	DefinitionCollection []*DictionaryType `xml:"DefinitionCollection,omitempty"`

	// generated from element "DefinitionProxy" of complex type DefinitionProxyType
	DefinitionProxy []*DefinitionProxyType `xml:"DefinitionProxy,omitempty"`

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant []*TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod []*TimePeriodType `xml:"TimePeriod,omitempty"`

	// generated from element "TimeNode" of complex type TimeNodeType
	TimeNode []*TimeNodeType `xml:"TimeNode,omitempty"`

	// generated from element "TimeEdge" of complex type TimeEdgeType
	TimeEdge []*TimeEdgeType `xml:"TimeEdge,omitempty"`

	// generated from element "TimeTopologyComplex" of complex type TimeTopologyComplexType
	TimeTopologyComplex []*TimeTopologyComplexType `xml:"TimeTopologyComplex,omitempty"`

	// generated from element "Node" of complex type NodeType
	Node []*NodeType `xml:"Node,omitempty"`

	// generated from element "Edge" of complex type EdgeType
	Edge []*EdgeType `xml:"Edge,omitempty"`

	// generated from element "Face" of complex type FaceType
	Face []*FaceType `xml:"Face,omitempty"`

	// generated from element "TopoSolid" of complex type TopoSolidType
	TopoSolid []*TopoSolidType `xml:"TopoSolid,omitempty"`

	// generated from element "TopoComplex" of complex type TopoComplexType
	TopoComplex []*TopoComplexType `xml:"TopoComplex,omitempty"`

	// generated from element "Bag" of complex type BagType
	Bag []*BagType `xml:"Bag,omitempty"`

	// generated from element "Array" of complex type ArrayType
	Array []*ArrayType `xml:"Array,omitempty"`

	// generated from element "DataBlock" of complex type DataBlockType
	DataBlock []*DataBlockType `xml:"DataBlock,omitempty"`

	// generated from element "File" of complex type FileType
	File []*FileType `xml:"File,omitempty"`

	// generated from element "coverageFunction" of complex type CoverageFunctionType
	CoverageFunction []*CoverageFunctionType `xml:"coverageFunction,omitempty"`

	// generated from element "CoverageMappingRule" of complex type MappingRuleType
	CoverageMappingRule []*MappingRuleType `xml:"CoverageMappingRule,omitempty"`

	// generated from element "GridFunction" of complex type GridFunctionType
	GridFunction []*GridFunctionType `xml:"GridFunction,omitempty"`

	// generated from anonymous complex type within element "Boolean"
	Boolean []*Boolean `xml:"Boolean,omitempty"`

	// generated from anonymous complex type within element "Category"
	Category []*Category `xml:"Category,omitempty"`

	// generated from anonymous complex type within element "Count"
	Count []*Count `xml:"Count,omitempty"`

	// generated from anonymous complex type within element "Quantity"
	Quantity []*Quantity `xml:"Quantity,omitempty"`

	// generated from element "BooleanList" of simple type gml:booleanOrNilReasonList
	BooleanList string `xml:"BooleanList,omitempty"`

	// generated from element "CategoryList" of complex type CodeOrNilReasonListType
	CategoryList []*basictypes.CodeOrNilReasonListType `xml:"CategoryList,omitempty"`

	// generated from element "CountList" of simple type gml:integerOrNilReasonList
	CountList string `xml:"CountList,omitempty"`

	// generated from element "QuantityList" of complex type MeasureOrNilReasonListType
	QuantityList []*basictypes.MeasureOrNilReasonListType `xml:"QuantityList,omitempty"`

	// generated from element "CompositeValue" of complex type CompositeValueType
	CompositeValue []*CompositeValueType `xml:"CompositeValue,omitempty"`

	// generated from element "ValueArray" of complex type ValueArrayType
	ValueArray []*ValueArrayType `xml:"ValueArray,omitempty"`

	// generated from element "CategoryExtent" of complex type CategoryExtentType
	CategoryExtent []*CategoryExtentType `xml:"CategoryExtent,omitempty"`

	// generated from element "CountExtent" of simple type gml:CountExtentType
	CountExtent string `xml:"CountExtent,omitempty"`

	// generated from element "QuantityExtent" of complex type QuantityExtentType
	QuantityExtent []*QuantityExtentType `xml:"QuantityExtent,omitempty"`

	// generated from element "ParameterValue" of complex type ParameterValueType
	ParameterValue []*ParameterValueType `xml:"ParameterValue,omitempty"`

	// generated from element "ParameterValueGroup" of complex type ParameterValueGroupType
	ParameterValueGroup []*ParameterValueGroupType `xml:"ParameterValueGroup,omitempty"`

	// generated from element "GenericMetaData" of complex type GenericMetaDataType
	GenericMetaData []*GenericMetaDataType `xml:"GenericMetaData,omitempty"`
}

// StringOrRefType Named source named complex type "StringOrRefType"
type StringOrRefType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// BagType Named source named complex type "BagType"
type BagType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "member" of complex type AssociationRoleType
	Member []*AssociationRoleType `xml:"member,omitempty"`

	// generated from element "members" of complex type ArrayAssociationType
	Members *ArrayAssociationType `xml:"members,omitempty"`
}

// ArrayType Named source named complex type "ArrayType"
type ArrayType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "members" of complex type ArrayAssociationType
	Members *ArrayAssociationType `xml:"members,omitempty"`
}

// MetaDataPropertyType Named source named complex type "MetaDataPropertyType"
type MetaDataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "about"
	About string `xml:"about,attr,omitempty"`

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "GenericMetaData" of complex type GenericMetaDataType
	GenericMetaData *GenericMetaDataType `xml:"GenericMetaData,omitempty"`
}

// AbstractMetaDataType Named source named complex type "AbstractMetaDataType"
type AbstractMetaDataType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`
}

// GenericMetaDataType Named source named complex type "GenericMetaDataType"
type GenericMetaDataType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`
}

// LocationPropertyType Named source named complex type "LocationPropertyType"
type LocationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`

	// generated from element "LocationKeyWord" of complex type CodeType
	LocationKeyWord *basictypes.CodeType `xml:"LocationKeyWord,omitempty"`

	// generated from element "LocationString" of complex type StringOrRefType
	LocationString *StringOrRefType `xml:"LocationString,omitempty"`

	// generated from element "Null" of simple type gml:NilReasonType
	Null string `xml:"Null,omitempty"`
}

// PriorityLocationPropertyType Named source named complex type "PriorityLocationPropertyType"
type PriorityLocationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "MultiGeometry" of complex type MultiGeometryType
	MultiGeometry *MultiGeometryType `xml:"MultiGeometry,omitempty"`

	// generated from element "MultiPoint" of complex type MultiPointType
	MultiPoint *MultiPointType `xml:"MultiPoint,omitempty"`

	// generated from element "MultiCurve" of complex type MultiCurveType
	MultiCurve *MultiCurveType `xml:"MultiCurve,omitempty"`

	// generated from element "MultiSurface" of complex type MultiSurfaceType
	MultiSurface *MultiSurfaceType `xml:"MultiSurface,omitempty"`

	// generated from element "MultiSolid" of complex type MultiSolidType
	MultiSolid *MultiSolidType `xml:"MultiSolid,omitempty"`

	// generated from element "Solid" of complex type SolidType
	Solid *SolidType `xml:"Solid,omitempty"`

	// generated from element "CompositeSolid" of complex type CompositeSolidType
	CompositeSolid *CompositeSolidType `xml:"CompositeSolid,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "PolyhedralSurface" of complex type SurfaceType
	PolyhedralSurface *SurfaceType `xml:"PolyhedralSurface,omitempty"`

	// generated from element "TriangulatedSurface" of complex type SurfaceType
	TriangulatedSurface *SurfaceType `xml:"TriangulatedSurface,omitempty"`

	// generated from element "Tin" of complex type TinType
	Tin *TinType `xml:"Tin,omitempty"`

	// generated from element "OrientableSurface" of complex type OrientableSurfaceType
	OrientableSurface *OrientableSurfaceType `xml:"OrientableSurface,omitempty"`

	// generated from element "Shell" of complex type ShellType
	Shell *ShellType `xml:"Shell,omitempty"`

	// generated from element "Polygon" of complex type PolygonType
	Polygon *PolygonType `xml:"Polygon,omitempty"`

	// generated from element "CompositeSurface" of complex type CompositeSurfaceType
	CompositeSurface *CompositeSurfaceType `xml:"CompositeSurface,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`

	// generated from element "LinearRing" of complex type LinearRingType
	LinearRing *LinearRingType `xml:"LinearRing,omitempty"`

	// generated from element "LineString" of complex type LineStringType
	LineString *LineStringType `xml:"LineString,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`

	// generated from element "GeometricComplex" of complex type GeometricComplexType
	GeometricComplex *GeometricComplexType `xml:"GeometricComplex,omitempty"`

	// generated from element "Grid" of complex type GridType
	Grid *GridType `xml:"Grid,omitempty"`

	// generated from element "RectifiedGrid" of complex type RectifiedGridType
	RectifiedGrid *RectifiedGridType `xml:"RectifiedGrid,omitempty"`

	// generated from element "LocationKeyWord" of complex type CodeType
	LocationKeyWord *basictypes.CodeType `xml:"LocationKeyWord,omitempty"`

	// generated from element "LocationString" of complex type StringOrRefType
	LocationString *StringOrRefType `xml:"LocationString,omitempty"`

	// generated from element "Null" of simple type gml:NilReasonType
	Null string `xml:"Null,omitempty"`

	// generated from attribute "priority"
	Priority string `xml:"priority,attr,omitempty"`
}

// FeatureArrayPropertyType Named source named complex type "FeatureArrayPropertyType"
type FeatureArrayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DynamicFeature" of complex type DynamicFeatureType
	DynamicFeature []*DynamicFeatureType `xml:"DynamicFeature,omitempty"`

	// generated from element "DynamicFeatureCollection" of complex type DynamicFeatureCollectionType
	DynamicFeatureCollection []*DynamicFeatureCollectionType `xml:"DynamicFeatureCollection,omitempty"`

	// generated from element "MultiPointCoverage" of complex type DiscreteCoverageType
	MultiPointCoverage []*DiscreteCoverageType `xml:"MultiPointCoverage,omitempty"`

	// generated from element "MultiCurveCoverage" of complex type DiscreteCoverageType
	MultiCurveCoverage []*DiscreteCoverageType `xml:"MultiCurveCoverage,omitempty"`

	// generated from element "MultiSurfaceCoverage" of complex type DiscreteCoverageType
	MultiSurfaceCoverage []*DiscreteCoverageType `xml:"MultiSurfaceCoverage,omitempty"`

	// generated from element "MultiSolidCoverage" of complex type DiscreteCoverageType
	MultiSolidCoverage []*DiscreteCoverageType `xml:"MultiSolidCoverage,omitempty"`

	// generated from element "GridCoverage" of complex type DiscreteCoverageType
	GridCoverage []*DiscreteCoverageType `xml:"GridCoverage,omitempty"`

	// generated from element "RectifiedGridCoverage" of complex type DiscreteCoverageType
	RectifiedGridCoverage []*DiscreteCoverageType `xml:"RectifiedGridCoverage,omitempty"`

	// generated from element "Observation" of complex type ObservationType
	Observation []*ObservationType `xml:"Observation,omitempty"`

	// generated from element "DirectedObservation" of complex type DirectedObservationType
	DirectedObservation []*DirectedObservationType `xml:"DirectedObservation,omitempty"`

	// generated from element "DirectedObservationAtDistance" of complex type DirectedObservationAtDistanceType
	DirectedObservationAtDistance []*DirectedObservationAtDistanceType `xml:"DirectedObservationAtDistance,omitempty"`

	// generated from element "FeatureCollection" of complex type FeatureCollectionType
	FeatureCollection []*FeatureCollectionType `xml:"FeatureCollection,omitempty"`
}

// BoundedFeatureType Named source named complex type "BoundedFeatureType"
type BoundedFeatureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 3588 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`
}

// AbstractFeatureCollectionType Named source named complex type "AbstractFeatureCollectionType"
type AbstractFeatureCollectionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "featureMember" of complex type FeaturePropertyType
	FeatureMember []*FeaturePropertyType `xml:"featureMember,omitempty"`

	// generated from element "featureMembers" of complex type FeatureArrayPropertyType
	FeatureMembers *FeatureArrayPropertyType `xml:"featureMembers,omitempty"`
}

// FeatureCollectionType Named source named complex type "FeatureCollectionType"
type FeatureCollectionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "featureMember" of complex type FeaturePropertyType
	FeatureMember []*FeaturePropertyType `xml:"featureMember,omitempty"`

	// generated from element "featureMembers" of complex type FeatureArrayPropertyType
	FeatureMembers *FeatureArrayPropertyType `xml:"featureMembers,omitempty"`
}

// IndirectEntryType Named source named complex type "IndirectEntryType"
type IndirectEntryType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DefinitionProxy" of complex type DefinitionProxyType
	DefinitionProxy *DefinitionProxyType `xml:"DefinitionProxy,omitempty"`
}

// DefinitionProxyType Named source named complex type "DefinitionProxyType"
type DefinitionProxyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "remarks" of simple type string
	Remarks string `xml:"remarks,omitempty"`

	// generated from element "definitionRef" of complex type ReferenceType
	DefinitionRef *ReferenceType `xml:"definitionRef,omitempty"`
}

// MovingObjectStatusType Named source named complex type "MovingObjectStatusType"
type MovingObjectStatusType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 753 depth 2
	Group_StandardObjectProperties

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *StringOrRefType `xml:"dataSource,omitempty"`

	// generated from element "position" of complex type GeometryPropertyType
	Position *GeometryPropertyType `xml:"position,omitempty"`

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`

	// generated from element "locationName" of complex type CodeType
	LocationName *basictypes.CodeType `xml:"locationName,omitempty"`

	// generated from element "locationReference" of complex type ReferenceType
	LocationReference *ReferenceType `xml:"locationReference,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "speed" of complex type MeasureType
	Speed *basictypes.MeasureType `xml:"speed,omitempty"`

	// generated from element "bearing" of complex type DirectionPropertyType
	Bearing *DirectionPropertyType `xml:"bearing,omitempty"`

	// generated from element "acceleration" of complex type MeasureType
	Acceleration *basictypes.MeasureType `xml:"acceleration,omitempty"`

	// generated from element "elevation" of complex type MeasureType
	Elevation *basictypes.MeasureType `xml:"elevation,omitempty"`

	// generated from element "status" of complex type StringOrRefType
	Status *StringOrRefType `xml:"status,omitempty"`

	// generated from element "statusReference" of complex type ReferenceType
	StatusReference *ReferenceType `xml:"statusReference,omitempty"`
}
