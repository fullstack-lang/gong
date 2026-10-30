// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAirspacePointPositionType
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

	// generated from element "Navaid" of complex type NavaidType
	Navaid *aixmfeatures.NavaidType `xml:"Navaid,omitempty"`

	// generated from element "VOR" of complex type VORType
	VOR *aixmfeatures.VORType `xml:"VOR,omitempty"`

	// generated from element "DME" of complex type DMEType
	DME *aixmfeatures.DMEType `xml:"DME,omitempty"`

	// generated from element "TACAN" of complex type TACANType
	TACAN *aixmfeatures.TACANType `xml:"TACAN,omitempty"`

	// generated from element "NDB" of complex type NDBType
	NDB *aixmfeatures.NDBType `xml:"NDB,omitempty"`

	// generated from element "Localizer" of complex type LocalizerType
	Localizer *aixmfeatures.LocalizerType `xml:"Localizer,omitempty"`

	// generated from element "Glidepath" of complex type GlidepathType
	Glidepath *aixmfeatures.GlidepathType `xml:"Glidepath,omitempty"`

	// generated from element "MarkerBeacon" of complex type MarkerBeaconType
	MarkerBeacon *aixmfeatures.MarkerBeaconType `xml:"MarkerBeacon,omitempty"`

	// generated from element "SDF" of complex type SDFType
	SDF *aixmfeatures.SDFType `xml:"SDF,omitempty"`

	// generated from element "DirectionFinder" of complex type DirectionFinderType
	DirectionFinder *aixmfeatures.DirectionFinderType `xml:"DirectionFinder,omitempty"`

	// generated from element "Azimuth" of complex type AzimuthType
	Azimuth *aixmfeatures.AzimuthType `xml:"Azimuth,omitempty"`

	// generated from element "Elevation" of complex type ElevationType
	Elevation *aixmfeatures.ElevationType `xml:"Elevation,omitempty"`

	// generated from element "CheckpointVOR" of complex type CheckpointVORType
	CheckpointVOR *aixmfeatures.CheckpointVORType `xml:"CheckpointVOR,omitempty"`

	// generated from element "CheckpointINS" of complex type CheckpointINSType
	CheckpointINS *aixmfeatures.CheckpointINSType `xml:"CheckpointINS,omitempty"`

	// generated from element "SpecialNavigationSystem" of complex type SpecialNavigationSystemType
	SpecialNavigationSystem *aixmfeatures.SpecialNavigationSystemType `xml:"SpecialNavigationSystem,omitempty"`

	// generated from element "SpecialNavigationStation" of complex type SpecialNavigationStationType
	SpecialNavigationStation *aixmfeatures.SpecialNavigationStationType `xml:"SpecialNavigationStation,omitempty"`

	// generated from element "DesignatedPoint" of complex type DesignatedPointType
	DesignatedPoint *aixmfeatures.DesignatedPointType `xml:"DesignatedPoint,omitempty"`

	// generated from element "SignificantPointInAirspace" of complex type SignificantPointInAirspaceType
	SignificantPointInAirspace *aixmfeatures.SignificantPointInAirspaceType `xml:"SignificantPointInAirspace,omitempty"`

	// generated from element "AngleIndication" of complex type AngleIndicationType
	AngleIndication *aixmfeatures.AngleIndicationType `xml:"AngleIndication,omitempty"`

	// generated from element "DistanceIndication" of complex type DistanceIndicationType
	DistanceIndication *aixmfeatures.DistanceIndicationType `xml:"DistanceIndication,omitempty"`

	// generated from element "RadarSystem" of complex type RadarSystemType
	RadarSystem *aixmfeatures.RadarSystemType `xml:"RadarSystem,omitempty"`

	// generated from element "PrimarySurveillanceRadar" of complex type PrimarySurveillanceRadarType
	PrimarySurveillanceRadar *aixmfeatures.PrimarySurveillanceRadarType `xml:"PrimarySurveillanceRadar,omitempty"`

	// generated from element "SecondarySurveillanceRadar" of complex type SecondarySurveillanceRadarType
	SecondarySurveillanceRadar *aixmfeatures.SecondarySurveillanceRadarType `xml:"SecondarySurveillanceRadar,omitempty"`

	// generated from element "PrecisionApproachRadar" of complex type PrecisionApproachRadarType
	PrecisionApproachRadar *aixmfeatures.PrecisionApproachRadarType `xml:"PrecisionApproachRadar,omitempty"`
}
