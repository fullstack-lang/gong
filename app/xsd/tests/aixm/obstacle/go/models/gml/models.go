// generated code - do not edit
package gml

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *xlink.Xlink_SchemaElement

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "NilReasonEnumeration_Member_1"
type Enum_NilReasonEnumeration_Member_1 string

const (
	Enum_NilReasonEnumeration_Member_1_Inapplicable Enum_NilReasonEnumeration_Member_1 = "inapplicable"
	Enum_NilReasonEnumeration_Member_1_Missing Enum_NilReasonEnumeration_Member_1 = "missing"
	Enum_NilReasonEnumeration_Member_1_Template Enum_NilReasonEnumeration_Member_1 = "template"
	Enum_NilReasonEnumeration_Member_1_Unknown Enum_NilReasonEnumeration_Member_1 = "unknown"
	Enum_NilReasonEnumeration_Member_1_Withheld Enum_NilReasonEnumeration_Member_1 = "withheld"
)

// From xsd simple type with enumerate restriction "SignType"
type Enum_SignType string

const (
	Enum_SignType_N__ Enum_SignType = "-"
	Enum_SignType_N___1 Enum_SignType = "+"
)

// From xsd simple type with enumerate restriction "AggregationType"
type Enum_AggregationType string

const (
	Enum_AggregationType_Sequence Enum_AggregationType = "sequence"
)

// From xsd simple type with enumerate restriction "TimeIndeterminateValueType"
type Enum_TimeIndeterminateValueType string

const (
	Enum_TimeIndeterminateValueType_Unknown Enum_TimeIndeterminateValueType = "unknown"
)

// From xsd simple type with enumerate restriction "TimePositionUnion_Member_1"
type Enum_TimePositionUnion_Member_1 string

const (
	Enum_TimePositionUnion_Member_1_Empty Enum_TimePositionUnion_Member_1 = ""
)

// From xsd simple type with enumerate restriction "CurveInterpolationType"
type Enum_CurveInterpolationType string

const (
	Enum_CurveInterpolationType_Linear Enum_CurveInterpolationType = "linear"
	Enum_CurveInterpolationType_Geodesic Enum_CurveInterpolationType = "geodesic"
	Enum_CurveInterpolationType_CircularArc3Points Enum_CurveInterpolationType = "circularArc3Points"
	Enum_CurveInterpolationType_CircularArcCenterPointWithRadius Enum_CurveInterpolationType = "circularArcCenterPointWithRadius"
)

// From xsd simple type with enumerate restriction "SurfaceInterpolationType"
type Enum_SurfaceInterpolationType string

const (
	Enum_SurfaceInterpolationType_Planar Enum_SurfaceInterpolationType = "planar"
)

// insertion point for gongstructs declarations

// CodeType Named source named complex type "CodeType"
// gml:CodeType is a generalized type to be used for a term, keyword or name.
// It adds a XML attribute codeSpace to a term, where the value of the codeSpace attribute (if present) shall indicate a dictionary, thesaurus, classification scheme, authority, or pattern for the term.
type CodeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// CodeWithAuthorityType Named source named complex type "CodeWithAuthorityType"
// gml:CodeWithAuthorityType requires that the codeSpace attribute is provided in an instance.
type CodeWithAuthorityType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// MeasureType Named source named complex type "MeasureType"
// gml:MeasureType supports recording an amount encoded as a value of XML Schema double, together with a units of measure indicated by an attribute uom, short for "units Of measure". The value of the uom attribute identifies a reference system for the amount, usually a ratio or interval scale.
type MeasureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// simpleContent character data
	EnclosedText float64 `xml:",chardata"`
}

// LengthType Named source named complex type "LengthType"
// This is a prototypical definition for a specific measure type defined as a vacuous extension (i.e. aliases) of gml:MeasureType. In this case, the content model supports the description of a length (or distance) quantity, with its units. The unit of measure referenced by uom shall be suitable for a length, such as metres or feet.
type LengthType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// AngleType Named source named complex type "AngleType"
type AngleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// AbstractGMLType Named source named complex type "AbstractGMLType"
type AbstractGMLType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties
}

// Group_StandardObjectProperties UnNamed source named group "StandardObjectProperties"
type Group_StandardObjectProperties struct {

	// insertion point for fields

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*CodeType `xml:"name,omitempty"`
}

// AttributeGroup_AssociationAttributeGroup UnNamed source named attribute group "AssociationAttributeGroup"
type AttributeGroup_AssociationAttributeGroup struct {

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from attribute group "xlink:simpleAttrs"
	xlink.AttributeGroup_simpleAttrs
}

// AttributeGroup_OwnershipAttributeGroup UnNamed source named attribute group "OwnershipAttributeGroup"
type AttributeGroup_OwnershipAttributeGroup struct {

	// insertion point for fields

	// generated from attribute "owns"
	Owns bool `xml:"owns,attr,omitempty"`
}

// Attr_id UnNamed source top-level attribute "id"
type Attr_id struct {

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`
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

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`
}

// BoundingShapeType Named source named complex type "BoundingShapeType"
type BoundingShapeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Envelope" of complex type EnvelopeType
	Envelope *EnvelopeType `xml:"Envelope,omitempty"`
}

// AbstractFeatureMemberType Named source named complex type "AbstractFeatureMemberType"
// To create a collection of GML features, a property type shall be derived by extension from gml:AbstractFeatureMemberType.
// By default, this abstract property type does not imply any ownership of the features in the collection. The owns attribute of gml:OwnershipAttributeGroup may be used on a property element instance to assert ownership of a feature in the collection. A collection shall not own a feature already owned by another object.
type AbstractFeatureMemberType struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Group_dynamicProperties UnNamed source named group "dynamicProperties"
type Group_dynamicProperties struct {

	// insertion point for fields
}

// DynamicFeatureType Named source named complex type "DynamicFeatureType"
type DynamicFeatureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *BoundingShapeType `xml:"boundedBy,omitempty"`
}

// AbstractTimeSliceType Named source named complex type "AbstractTimeSliceType"
type AbstractTimeSliceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *TimePrimitivePropertyType `xml:"validTime,omitempty"`
}

// AbstractTimeObjectType Named source named complex type "AbstractTimeObjectType"
type AbstractTimeObjectType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties
}

// AbstractTimePrimitiveType Named source named complex type "AbstractTimePrimitiveType"
type AbstractTimePrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties
}

// TimePrimitivePropertyType Named source named complex type "TimePrimitivePropertyType"
// gml:TimePrimitivePropertyType provides a standard content model for associations between an arbitrary member of the substitution group whose head is gml:AbstractTimePrimitive and another object.
type TimePrimitivePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "TimeInstant" of complex type TimeInstantType
	TimeInstant *TimeInstantType `xml:"TimeInstant,omitempty"`

	// generated from element "TimePeriod" of complex type TimePeriodType
	TimePeriod *TimePeriodType `xml:"TimePeriod,omitempty"`
}

// AbstractTimeGeometricPrimitiveType Named source named complex type "AbstractTimeGeometricPrimitiveType"
type AbstractTimeGeometricPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`
}

// TimeInstantType Named source named complex type "TimeInstantType"
type TimeInstantType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from element "timePosition" of complex type TimePositionType
	TimePosition *TimePositionType `xml:"timePosition,omitempty"`
}

// TimePeriodType Named source named complex type "TimePeriodType"
type TimePeriodType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute "frame"
	Frame string `xml:"frame,attr,omitempty"`

	// generated from element "beginPosition" of complex type TimePositionType
	BeginPosition *TimePositionType `xml:"beginPosition,omitempty"`

	// generated from element "endPosition" of complex type TimePositionType
	EndPosition *TimePositionType `xml:"endPosition,omitempty"`
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

	// generated from attribute "indeterminatePosition"
	IndeterminatePosition Enum_TimeIndeterminateValueType `xml:"indeterminatePosition,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
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

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// AttributeGroup_SRSReferenceGroup UnNamed source named attribute group "SRSReferenceGroup"
type AttributeGroup_SRSReferenceGroup struct {

	// insertion point for fields

	// generated from attribute "srsName"
	SrsName string `xml:"srsName,attr,omitempty"`
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
}

// AbstractGeometricPrimitiveType Named source named complex type "AbstractGeometricPrimitiveType"
// gml:AbstractGeometricPrimitiveType is the abstract root type of the geometric primitives. A geometric primitive is a geometric object that is not decomposed further into other primitives in the system. All primitives are oriented in the direction implied by the sequence of their coordinate tuples.
type AbstractGeometricPrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// PointType Named source named complex type "PointType"
type PointType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`
}

// PointPropertyType Named source named complex type "PointPropertyType"
// A property that has a point as its value domain may either be an appropriate geometry element encapsulated in an element of this type or an XLink reference to a remote geometry element (where remote includes geometry elements located elsewhere in the same document). Either the reference or the contained element shall be given, but neither both nor none.
type PointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	AttributeGroup_AssociationAttributeGroup

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`
}

// AbstractCurveType Named source named complex type "AbstractCurveType"
// gml:AbstractCurveType is an abstraction of a curve to support the different levels of complexity. The curve may always be viewed as a geometric primitive, i.e. is continuous.
type AbstractCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
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

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "OrientableCurve" of complex type OrientableCurveType
	OrientableCurve *OrientableCurveType `xml:"OrientableCurve,omitempty"`

	// generated from element "CompositeCurve" of complex type CompositeCurveType
	CompositeCurve *CompositeCurveType `xml:"CompositeCurve,omitempty"`
}

// AbstractSurfaceType Named source named complex type "AbstractSurfaceType"
// gml:AbstractSurfaceType is an abstraction of a surface to support the different levels of complexity. A surface is always a continuous region of a plane.
type AbstractSurfaceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup
}

// AbstractRingType Named source named complex type "AbstractRingType"
type AbstractRingType struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// AbstractRingPropertyType Named source named complex type "AbstractRingPropertyType"
// A property with the content model of gml:AbstractRingPropertyType encapsulates a ring to represent the surface boundary property of a surface.
type AbstractRingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "Ring" of complex type RingType
	Ring *RingType `xml:"Ring,omitempty"`
}

// CurveType Named source named complex type "CurveType"
type CurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute "orientation"
	Orientation Enum_SignType `xml:"orientation,attr,omitempty"`

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

	// generated from element "ArcByCenterPoint" of complex type ArcByCenterPointType
	ArcByCenterPoint []*ArcByCenterPointType `xml:"ArcByCenterPoint,omitempty"`

	// generated from element "CircleByCenterPoint" of complex type CircleByCenterPointType
	CircleByCenterPoint []*CircleByCenterPointType `xml:"CircleByCenterPoint,omitempty"`

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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`
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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`
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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`
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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`
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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "radius" of complex type LengthType
	Radius *LengthType `xml:"radius,omitempty"`

	// generated from element "startAngle" of complex type AngleType
	StartAngle *AngleType `xml:"startAngle,omitempty"`

	// generated from element "endAngle" of complex type AngleType
	EndAngle *AngleType `xml:"endAngle,omitempty"`
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

	// generated from element "posList" of complex type DirectPositionListType
	PosList *DirectPositionListType `xml:"posList,omitempty"`

	// generated from element "radius" of complex type LengthType
	Radius *LengthType `xml:"radius,omitempty"`
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

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "patches" of complex type SurfacePatchArrayPropertyType
	Patches *SurfacePatchArrayPropertyType `xml:"patches,omitempty"`
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
}

// PolygonPatchType Named source named complex type "PolygonPatchType"
type PolygonPatchType struct {
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

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "curveMember" of complex type CurvePropertyType
	CurveMember []*CurvePropertyType `xml:"curveMember,omitempty"`
}

// CompositeCurveType Named source named complex type "CompositeCurveType"
type CompositeCurveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 140 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from attribute group "gml:AggregationAttributeGroup"
	AttributeGroup_AggregationAttributeGroup

	// generated from element "curveMember" of complex type CurvePropertyType
	CurveMember []*CurvePropertyType `xml:"curveMember,omitempty"`
}
