// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeBuoyDesignatorType
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

	// generated from element "AirportHeliport" of complex type AirportHeliportType
	AirportHeliport *aixmfeatures.AirportHeliportType `xml:"AirportHeliport,omitempty"`

	// generated from element "AirportHeliportCollocation" of complex type AirportHeliportCollocationType
	AirportHeliportCollocation *aixmfeatures.AirportHeliportCollocationType `xml:"AirportHeliportCollocation,omitempty"`

	// generated from element "AirportHotSpot" of complex type AirportHotSpotType
	AirportHotSpot *aixmfeatures.AirportHotSpotType `xml:"AirportHotSpot,omitempty"`

	// generated from element "AltimeterSource" of complex type AltimeterSourceType
	AltimeterSource *aixmfeatures.AltimeterSourceType `xml:"AltimeterSource,omitempty"`

	// generated from element "NonMovementArea" of complex type NonMovementAreaType
	NonMovementArea *aixmfeatures.NonMovementAreaType `xml:"NonMovementArea,omitempty"`

	// generated from element "SurveyControlPoint" of complex type SurveyControlPointType
	SurveyControlPoint *aixmfeatures.SurveyControlPointType `xml:"SurveyControlPoint,omitempty"`

	// generated from element "WorkArea" of complex type WorkAreaType
	WorkArea *aixmfeatures.WorkAreaType `xml:"WorkArea,omitempty"`

	// generated from element "Apron" of complex type ApronType
	Apron *aixmfeatures.ApronType `xml:"Apron,omitempty"`

	// generated from element "ApronElement" of complex type ApronElementType
	ApronElement *aixmfeatures.ApronElementType `xml:"ApronElement,omitempty"`

	// generated from element "AircraftStand" of complex type AircraftStandType
	AircraftStand *aixmfeatures.AircraftStandType `xml:"AircraftStand,omitempty"`

	// generated from element "DeicingArea" of complex type DeicingAreaType
	DeicingArea *aixmfeatures.DeicingAreaType `xml:"DeicingArea,omitempty"`

	// generated from element "PassengerLoadingBridge" of complex type PassengerLoadingBridgeType
	PassengerLoadingBridge *aixmfeatures.PassengerLoadingBridgeType `xml:"PassengerLoadingBridge,omitempty"`

	// generated from element "Road" of complex type RoadType
	Road *aixmfeatures.RoadType `xml:"Road,omitempty"`

	// generated from element "TouchDownLiftOff" of complex type TouchDownLiftOffType
	TouchDownLiftOff *aixmfeatures.TouchDownLiftOffType `xml:"TouchDownLiftOff,omitempty"`

	// generated from element "TouchDownLiftOffSafeArea" of complex type TouchDownLiftOffSafeAreaType
	TouchDownLiftOffSafeArea *aixmfeatures.TouchDownLiftOffSafeAreaType `xml:"TouchDownLiftOffSafeArea,omitempty"`

	// generated from element "Runway" of complex type RunwayType
	Runway *aixmfeatures.RunwayType `xml:"Runway,omitempty"`

	// generated from element "RunwayElement" of complex type RunwayElementType
	RunwayElement *aixmfeatures.RunwayElementType `xml:"RunwayElement,omitempty"`

	// generated from element "RunwayDirection" of complex type RunwayDirectionType
	RunwayDirection *aixmfeatures.RunwayDirectionType `xml:"RunwayDirection,omitempty"`

	// generated from element "RunwayCentrelinePoint" of complex type RunwayCentrelinePointType
	RunwayCentrelinePoint *aixmfeatures.RunwayCentrelinePointType `xml:"RunwayCentrelinePoint,omitempty"`

	// generated from element "RunwayBlastPad" of complex type RunwayBlastPadType
	RunwayBlastPad *aixmfeatures.RunwayBlastPadType `xml:"RunwayBlastPad,omitempty"`

	// generated from element "RunwayProtectArea" of complex type RunwayProtectAreaType
	RunwayProtectArea *aixmfeatures.RunwayProtectAreaType `xml:"RunwayProtectArea,omitempty"`

	// generated from element "RunwayVisualRange" of complex type RunwayVisualRangeType
	RunwayVisualRange *aixmfeatures.RunwayVisualRangeType `xml:"RunwayVisualRange,omitempty"`

	// generated from element "VisualGlideSlopeIndicator" of complex type VisualGlideSlopeIndicatorType
	VisualGlideSlopeIndicator *aixmfeatures.VisualGlideSlopeIndicatorType `xml:"VisualGlideSlopeIndicator,omitempty"`

	// generated from element "ArrestingGear" of complex type ArrestingGearType
	ArrestingGear *aixmfeatures.ArrestingGearType `xml:"ArrestingGear,omitempty"`

	// generated from element "Taxiway" of complex type TaxiwayType
	Taxiway *aixmfeatures.TaxiwayType `xml:"Taxiway,omitempty"`

	// generated from element "TaxiwayElement" of complex type TaxiwayElementType
	TaxiwayElement *aixmfeatures.TaxiwayElementType `xml:"TaxiwayElement,omitempty"`

	// generated from element "TaxiHoldingPosition" of complex type TaxiHoldingPositionType
	TaxiHoldingPosition *aixmfeatures.TaxiHoldingPositionType `xml:"TaxiHoldingPosition,omitempty"`

	// generated from element "GuidanceLine" of complex type GuidanceLineType
	GuidanceLine *aixmfeatures.GuidanceLineType `xml:"GuidanceLine,omitempty"`

	// generated from element "SeaplaneLandingArea" of complex type SeaplaneLandingAreaType
	SeaplaneLandingArea *aixmfeatures.SeaplaneLandingAreaType `xml:"SeaplaneLandingArea,omitempty"`

	// generated from element "SeaplaneRampSite" of complex type SeaplaneRampSiteType
	SeaplaneRampSite *aixmfeatures.SeaplaneRampSiteType `xml:"SeaplaneRampSite,omitempty"`

	// generated from element "FloatingDockSite" of complex type FloatingDockSiteType
	FloatingDockSite *aixmfeatures.FloatingDockSiteType `xml:"FloatingDockSite,omitempty"`

	// generated from element "MarkingBuoy" of complex type MarkingBuoyType
	MarkingBuoy *aixmfeatures.MarkingBuoyType `xml:"MarkingBuoy,omitempty"`

	// generated from element "ApproachLightingSystem" of complex type ApproachLightingSystemType
	ApproachLightingSystem *aixmfeatures.ApproachLightingSystemType `xml:"ApproachLightingSystem,omitempty"`

	// generated from element "RunwayDirectionLightSystem" of complex type RunwayDirectionLightSystemType
	RunwayDirectionLightSystem *aixmfeatures.RunwayDirectionLightSystemType `xml:"RunwayDirectionLightSystem,omitempty"`

	// generated from element "RunwayProtectAreaLightSystem" of complex type RunwayProtectAreaLightSystemType
	RunwayProtectAreaLightSystem *aixmfeatures.RunwayProtectAreaLightSystemType `xml:"RunwayProtectAreaLightSystem,omitempty"`

	// generated from element "TaxiwayLightSystem" of complex type TaxiwayLightSystemType
	TaxiwayLightSystem *aixmfeatures.TaxiwayLightSystemType `xml:"TaxiwayLightSystem,omitempty"`

	// generated from element "TaxiHoldingPositionLightSystem" of complex type TaxiHoldingPositionLightSystemType
	TaxiHoldingPositionLightSystem *aixmfeatures.TaxiHoldingPositionLightSystemType `xml:"TaxiHoldingPositionLightSystem,omitempty"`

	// generated from element "GuidanceLineLightSystem" of complex type GuidanceLineLightSystemType
	GuidanceLineLightSystem *aixmfeatures.GuidanceLineLightSystemType `xml:"GuidanceLineLightSystem,omitempty"`

	// generated from element "ApronLightSystem" of complex type ApronLightSystemType
	ApronLightSystem *aixmfeatures.ApronLightSystemType `xml:"ApronLightSystem,omitempty"`

	// generated from element "TouchDownLiftOffLightSystem" of complex type TouchDownLiftOffLightSystemType
	TouchDownLiftOffLightSystem *aixmfeatures.TouchDownLiftOffLightSystemType `xml:"TouchDownLiftOffLightSystem,omitempty"`

	// generated from element "RunwayMarking" of complex type RunwayMarkingType
	RunwayMarking *aixmfeatures.RunwayMarkingType `xml:"RunwayMarking,omitempty"`

	// generated from element "TaxiwayMarking" of complex type TaxiwayMarkingType
	TaxiwayMarking *aixmfeatures.TaxiwayMarkingType `xml:"TaxiwayMarking,omitempty"`

	// generated from element "TaxiHoldingPositionMarking" of complex type TaxiHoldingPositionMarkingType
	TaxiHoldingPositionMarking *aixmfeatures.TaxiHoldingPositionMarkingType `xml:"TaxiHoldingPositionMarking,omitempty"`

	// generated from element "StandMarking" of complex type StandMarkingType
	StandMarking *aixmfeatures.StandMarkingType `xml:"StandMarking,omitempty"`

	// generated from element "ApronMarking" of complex type ApronMarkingType
	ApronMarking *aixmfeatures.ApronMarkingType `xml:"ApronMarking,omitempty"`

	// generated from element "DeicingAreaMarking" of complex type DeicingAreaMarkingType
	DeicingAreaMarking *aixmfeatures.DeicingAreaMarkingType `xml:"DeicingAreaMarking,omitempty"`

	// generated from element "GuidanceLineMarking" of complex type GuidanceLineMarkingType
	GuidanceLineMarking *aixmfeatures.GuidanceLineMarkingType `xml:"GuidanceLineMarking,omitempty"`

	// generated from element "TouchDownLiftOffMarking" of complex type TouchDownLiftOffMarkingType
	TouchDownLiftOffMarking *aixmfeatures.TouchDownLiftOffMarkingType `xml:"TouchDownLiftOffMarking,omitempty"`

	// generated from element "AirportProtectionAreaMarking" of complex type AirportProtectionAreaMarkingType
	AirportProtectionAreaMarking *aixmfeatures.AirportProtectionAreaMarkingType `xml:"AirportProtectionAreaMarking,omitempty"`
}
