// generated code - do not edit
package gml

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
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

// AbstractGMLType Named source named complex type "AbstractGMLType"
type AbstractGMLType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
	Group_StandardObjectProperties
}

// AbstractTimePrimitiveType Named source named complex type "AbstractTimePrimitiveType"
type AbstractTimePrimitiveType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
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

	// generated from group "gml:StandardObjectProperties" with order 156 depth 2
	Group_StandardObjectProperties

	// generated from attribute group "gml:SRSReferenceGroup"
	AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos *DirectPositionType `xml:"pos,omitempty"`
}
