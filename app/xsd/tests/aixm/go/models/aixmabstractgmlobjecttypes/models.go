// generated code - do not edit
package aixmabstractgmlobjecttypes

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *basictypes.CodeType
var _ *gmd.AbstractDS_Aggregate_Type
var _ *gml.DynamicFeatureType

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

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

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

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*gml.MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *gml.ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *gml.LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *gml.PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "history" of complex type HistoryPropertyType
	History *gml.HistoryPropertyType `xml:"history,omitempty"`

	// generated from element "track" of complex type HistoryPropertyType
	Track *gml.HistoryPropertyType `xml:"track,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *gml.StringOrRefType `xml:"dataSource,omitempty"`

	// generated from element "dataSourceReference" of complex type ReferenceType
	DataSourceReference *gml.ReferenceType `xml:"dataSourceReference,omitempty"`

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

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*gml.MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *gml.ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *gml.StringOrRefType `xml:"dataSource,omitempty"`

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

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*gml.MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *gml.ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *gml.LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *gml.PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "history" of complex type HistoryPropertyType
	History *gml.HistoryPropertyType `xml:"history,omitempty"`

	// generated from element "track" of complex type HistoryPropertyType
	Track *gml.HistoryPropertyType `xml:"track,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *gml.StringOrRefType `xml:"dataSource,omitempty"`

	// generated from element "dataSourceReference" of complex type ReferenceType
	DataSourceReference *gml.ReferenceType `xml:"dataSourceReference,omitempty"`

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

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*gml.MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *gml.ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`

	// generated from element "boundedBy" of complex type BoundingShapeType
	BoundedBy *gml.BoundingShapeType `xml:"boundedBy,omitempty"`

	// generated from element "location" of complex type LocationPropertyType
	Location *gml.LocationPropertyType `xml:"location,omitempty"`

	// generated from element "priorityLocation" of complex type PriorityLocationPropertyType
	PriorityLocation *gml.PriorityLocationPropertyType `xml:"priorityLocation,omitempty"`

	// generated from element "validTime" of complex type TimePrimitivePropertyType
	ValidTime *gml.TimePrimitivePropertyType `xml:"validTime,omitempty"`

	// generated from element "history" of complex type HistoryPropertyType
	History *gml.HistoryPropertyType `xml:"history,omitempty"`

	// generated from element "track" of complex type HistoryPropertyType
	Track *gml.HistoryPropertyType `xml:"track,omitempty"`

	// generated from element "dataSource" of complex type StringOrRefType
	DataSource *gml.StringOrRefType `xml:"dataSource,omitempty"`

	// generated from element "dataSourceReference" of complex type ReferenceType
	DataSourceReference *gml.ReferenceType `xml:"dataSourceReference,omitempty"`

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

	// generated from element "metaDataProperty" of complex type MetaDataPropertyType
	MetaDataProperty []*gml.MetaDataPropertyType `xml:"metaDataProperty,omitempty"`

	// generated from element "description" of complex type StringOrRefType
	Description *gml.StringOrRefType `xml:"description,omitempty"`

	// generated from element "descriptionReference" of complex type ReferenceType
	DescriptionReference *gml.ReferenceType `xml:"descriptionReference,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *basictypes.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*basictypes.CodeType `xml:"name,omitempty"`
}

// FeatureMetadataPropertyType Named source named complex type "FeatureMetadataPropertyType"
type FeatureMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MD_Metadata" of complex type MD_Metadata_Type
	MD_Metadata *gmd.MD_Metadata_Type `xml:"MD_Metadata,omitempty"`
}

// FeatureTimeSliceMetadataPropertyType Named source named complex type "FeatureTimeSliceMetadataPropertyType"
type FeatureTimeSliceMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MD_Metadata" of complex type MD_Metadata_Type
	MD_Metadata *gmd.MD_Metadata_Type `xml:"MD_Metadata,omitempty"`
}

// MessageMetadataPropertyType Named source named complex type "MessageMetadataPropertyType"
type MessageMetadataPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MD_Metadata" of complex type MD_Metadata_Type
	MD_Metadata *gmd.MD_Metadata_Type `xml:"MD_Metadata,omitempty"`
}
