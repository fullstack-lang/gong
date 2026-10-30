// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAircraftGroundServiceType
var _ *aixmfeatures.AerialRefuellingPropertyType
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

	// generated from element "OrganisationAuthority" of complex type OrganisationAuthorityType
	OrganisationAuthority *aixmfeatures.OrganisationAuthorityType `xml:"OrganisationAuthority,omitempty"`

	// generated from element "Unit" of complex type UnitType
	Unit *aixmfeatures.UnitType `xml:"Unit,omitempty"`

	// generated from element "AirTrafficControlService" of complex type AirTrafficControlServiceType
	AirTrafficControlService *aixmfeatures.AirTrafficControlServiceType `xml:"AirTrafficControlService,omitempty"`

	// generated from element "AirTrafficManagementService" of complex type AirTrafficManagementServiceType
	AirTrafficManagementService *aixmfeatures.AirTrafficManagementServiceType `xml:"AirTrafficManagementService,omitempty"`

	// generated from element "GroundTrafficControlService" of complex type GroundTrafficControlServiceType
	GroundTrafficControlService *aixmfeatures.GroundTrafficControlServiceType `xml:"GroundTrafficControlService,omitempty"`

	// generated from element "InformationService" of complex type InformationServiceType
	InformationService *aixmfeatures.InformationServiceType `xml:"InformationService,omitempty"`

	// generated from element "SearchRescueService" of complex type SearchRescueServiceType
	SearchRescueService *aixmfeatures.SearchRescueServiceType `xml:"SearchRescueService,omitempty"`

	// generated from element "AircraftGroundService" of complex type AircraftGroundServiceType
	AircraftGroundService *aixmfeatures.AircraftGroundServiceType `xml:"AircraftGroundService,omitempty"`

	// generated from element "AirportClearanceService" of complex type AirportClearanceServiceType
	AirportClearanceService *aixmfeatures.AirportClearanceServiceType `xml:"AirportClearanceService,omitempty"`

	// generated from element "AirportSuppliesService" of complex type AirportSuppliesServiceType
	AirportSuppliesService *aixmfeatures.AirportSuppliesServiceType `xml:"AirportSuppliesService,omitempty"`

	// generated from element "FireFightingService" of complex type FireFightingServiceType
	FireFightingService *aixmfeatures.FireFightingServiceType `xml:"FireFightingService,omitempty"`

	// generated from element "PassengerService" of complex type PassengerServiceType
	PassengerService *aixmfeatures.PassengerServiceType `xml:"PassengerService,omitempty"`

	// generated from element "PilotControlledLighting" of complex type PilotControlledLightingType
	PilotControlledLighting *aixmfeatures.PilotControlledLightingType `xml:"PilotControlledLighting,omitempty"`

	// generated from element "RadioCommunicationChannel" of complex type RadioCommunicationChannelType
	RadioCommunicationChannel *aixmfeatures.RadioCommunicationChannelType `xml:"RadioCommunicationChannel,omitempty"`

	// generated from element "SpecialDate" of complex type SpecialDateType
	SpecialDate *aixmfeatures.SpecialDateType `xml:"SpecialDate,omitempty"`

	// generated from element "StandardLevelTable" of complex type StandardLevelTableType
	StandardLevelTable *aixmfeatures.StandardLevelTableType `xml:"StandardLevelTable,omitempty"`

	// generated from element "StandardLevelColumn" of complex type StandardLevelColumnType
	StandardLevelColumn *aixmfeatures.StandardLevelColumnType `xml:"StandardLevelColumn,omitempty"`

	// generated from element "StandardLevelSector" of complex type StandardLevelSectorType
	StandardLevelSector *aixmfeatures.StandardLevelSectorType `xml:"StandardLevelSector,omitempty"`

	// generated from element "StandardLevel" of complex type StandardLevelType
	StandardLevel *aixmfeatures.StandardLevelType `xml:"StandardLevel,omitempty"`
}
