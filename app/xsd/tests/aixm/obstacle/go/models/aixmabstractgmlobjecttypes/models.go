// generated code - do not edit
package aixmabstractgmlobjecttypes

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *gml.CodeType

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "interpretation_SimpleType"
type Enum_Interpretation_SimpleType string

const (
	Enum_Interpretation_SimpleType_BASELINE Enum_Interpretation_SimpleType = "BASELINE"
	Enum_Interpretation_SimpleType_SNAPSHOT Enum_Interpretation_SimpleType = "SNAPSHOT"
	Enum_Interpretation_SimpleType_TEMPDELTA Enum_Interpretation_SimpleType = "TEMPDELTA"
	Enum_Interpretation_SimpleType_PERMDELTA Enum_Interpretation_SimpleType = "PERMDELTA"
)

// insertion point for gongstructs declarations

// AbstractAIXMFeatureBaseType Named source named complex type "AbstractAIXMFeatureBaseType"
// This derives from gml:DynamicFeatureType, as all AIXM features are expected to have temporal properties modeled using the Timeslice model. While the GML specification does not mandate that dynamic features derive from gml:DynamicFeatureType, many GML-aware applications may only detect a dynamic feature by this type hierarchy. Note that all temporal properties of gml:DynamicFeatureType are removed during this restriction, because they will be added for each AIXM feature specifically, based on the TimeSlice model
type AbstractAIXMFeatureBaseType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`
}

// AbstractAIXMFeatureType Named source named complex type "AbstractAIXMFeatureType"
// Adds the FeatureMetadata, which is common to all AIXM features
type AbstractAIXMFeatureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "featureMetadata" of complex type FeatureMetadataPropertyType
	FeatureMetadata *FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`
}

// AbstractAIXMObjectType Named source named complex type "AbstractAIXMObjectType"
// Base type for AIXM complex types that are NOT features. For example, City, ContactInformation, AirspaceVolume, etc. It derives from AbstractGMLType so that AIXM objects are recognised as GML objects, thus ensuring that GML-aware applications recognise them properly. Retains only the mandatory gml:id attribute.
type AbstractAIXMObjectType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id
}

// AbstractAIXMPropertyType Named source named complex type "AbstractAIXMPropertyType"
// It provides a basis for deriving AIXM feature/object properties. It defines the nilReason attribute and currently, it is only used for properties that are derived from association with an AIM object.
type AbstractAIXMPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`
}

// AbstractAIXMTimeSliceBaseType Named source named complex type "AbstractAIXMTimeSliceBaseType"
// Base type of AIXM Timeslices.  Removes option attributes that are not used in AIXM.
type AbstractAIXMTimeSliceBaseType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`
}

// AbstractAIXMTimeSliceType Named source named complex type "AbstractAIXMTimeSliceType"
// Adds in the AIXM specific properties, such as 'interpretation'.
type AbstractAIXMTimeSliceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "interpretation" of simple type interpretation_SimpleType
	Interpretation Enum_Interpretation_SimpleType `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`
}

// AbstractAIXMMessageBaseType Named source named complex type "AbstractAIXMMessageBaseType"
type AbstractAIXMMessageBaseType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "messageMetadata" of complex type MessageMetadataPropertyType
	MessageMetadata *MessageMetadataPropertyType `xml:"messageMetadata,omitempty"`
}

// AbstractAIXMMessageType Named source named complex type "AbstractAIXMMessageType"
type AbstractAIXMMessageType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "messageMetadata" of complex type MessageMetadataPropertyType
	MessageMetadata *MessageMetadataPropertyType `xml:"messageMetadata,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	gml.AttributeGroup_AggregationAttributeGroup
}

// AbstractExtensionType Named source named complex type "AbstractExtensionType"
type AbstractExtensionType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`
}

// FeatureMetadataPropertyType Named source named complex type "FeatureMetadataPropertyType"
type FeatureMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup
}

// FeatureTimeSliceMetadataPropertyType Named source named complex type "FeatureTimeSliceMetadataPropertyType"
type FeatureTimeSliceMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup
}

// MessageMetadataPropertyType Named source named complex type "MessageMetadataPropertyType"
type MessageMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup
}
