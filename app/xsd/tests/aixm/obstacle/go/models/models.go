// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeColourType
var _ *aixmfeatures.AirportHeliportPropertyType
var _ *gml.CodeType
var _ *xlink.Xlink_SchemaElement

// insertion point for enums declaration

// insertion point for gongstructs declarations

// Group_AIXMBasicMessagePropertyGroup UnNamed source named group "AIXMBasicMessagePropertyGroup"
type Group_AIXMBasicMessagePropertyGroup struct {

	// insertion point for fields

	// generated from element "hasMember" of complex type BasicMessageMemberAIXMPropertyType
	HasMember []*BasicMessageMemberAIXMPropertyType `xml:"hasMember,omitempty"`
}

// AIXMBasicMessageType Named source named complex type "AIXMBasicMessageType"
type AIXMBasicMessageType struct {
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
	MessageMetadata *aixmabstractgmlobjecttypes.MessageMetadataPropertyType `xml:"messageMetadata,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	gml.AttributeGroup_AggregationAttributeGroup

	// generated from group "message:AIXMBasicMessagePropertyGroup" with order 6 depth 2
	Group_AIXMBasicMessagePropertyGroup
}

// AIXMBasicMessage Named source root element AIXMBasicMessage of type message:AIXMBasicMessageType
type AIXMBasicMessage struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"AIXMBasicMessage"`

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
	MessageMetadata *aixmabstractgmlobjecttypes.MessageMetadataPropertyType `xml:"messageMetadata,omitempty"`

	// generated from attribute group "gml:AggregationAttributeGroup"
	gml.AttributeGroup_AggregationAttributeGroup

	// generated from group "message:AIXMBasicMessagePropertyGroup" with order 6 depth 2
	Group_AIXMBasicMessagePropertyGroup
}

// BasicMessageMemberAIXMPropertyType Named source named complex type "BasicMessageMemberAIXMPropertyType"
type BasicMessageMemberAIXMPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup

	// generated from element "VerticalStructure" of complex type VerticalStructureType
	VerticalStructure *aixmfeatures.VerticalStructureType `xml:"VerticalStructure,omitempty"`

	// generated from element "ObstacleArea" of complex type ObstacleAreaType
	ObstacleArea *aixmfeatures.ObstacleAreaType `xml:"ObstacleArea,omitempty"`

	// generated from element "AeronauticalGroundLight" of complex type AeronauticalGroundLightType
	AeronauticalGroundLight *aixmfeatures.AeronauticalGroundLightType `xml:"AeronauticalGroundLight,omitempty"`
}
