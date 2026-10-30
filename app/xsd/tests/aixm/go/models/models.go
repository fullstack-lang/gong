// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/applicationschema"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/measures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/portrayalcatalogue"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xmlxsd"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeBuoyDesignatorType
var _ *aixmfeatures.AerialRefuellingPropertyType
var _ *applicationschema.MD_ApplicationSchemaInformation_Type
var _ *basictypes.CodeType
var _ *content.MD_FeatureCatalogueDescription_Type
var _ *gcobase.AbstractObject_Type
var _ *gmd.AbstractDS_Aggregate_Type
var _ *gml.DynamicFeatureType
var _ *measures.LengthType
var _ *metadataextension.MD_ExtendedElementInformation_Type
var _ *portrayalcatalogue.MD_PortrayalCatalogueReference_Type
var _ *spatialrepresentation.MD_GridSpatialRepresentation_Type
var _ *xlink.Simple
var _ *xmlxsd.Xmlxsd_SchemaElement

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

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup

	// generated from element "AerialRefuelling" of complex type AerialRefuellingType
	AerialRefuelling *aixmfeatures.AerialRefuellingType `xml:"AerialRefuelling,omitempty"`

	// generated from element "AirportHeliport" of complex type AirportHeliportType
	AirportHeliport *aixmfeatures.AirportHeliportType `xml:"AirportHeliport,omitempty"`

	// generated from element "AirportHeliportCollocation" of complex type AirportHeliportCollocationType
	AirportHeliportCollocation *aixmfeatures.AirportHeliportCollocationType `xml:"AirportHeliportCollocation,omitempty"`

	// generated from element "TouchDownLiftOffSafeArea" of complex type TouchDownLiftOffSafeAreaType
	TouchDownLiftOffSafeArea *aixmfeatures.TouchDownLiftOffSafeAreaType `xml:"TouchDownLiftOffSafeArea,omitempty"`

	// generated from element "RunwayProtectArea" of complex type RunwayProtectAreaType
	RunwayProtectArea *aixmfeatures.RunwayProtectAreaType `xml:"RunwayProtectArea,omitempty"`

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

	// generated from element "AircraftStand" of complex type AircraftStandType
	AircraftStand *aixmfeatures.AircraftStandType `xml:"AircraftStand,omitempty"`

	// generated from element "Apron" of complex type ApronType
	Apron *aixmfeatures.ApronType `xml:"Apron,omitempty"`

	// generated from element "ApronElement" of complex type ApronElementType
	ApronElement *aixmfeatures.ApronElementType `xml:"ApronElement,omitempty"`

	// generated from element "DeicingArea" of complex type DeicingAreaType
	DeicingArea *aixmfeatures.DeicingAreaType `xml:"DeicingArea,omitempty"`

	// generated from element "PassengerLoadingBridge" of complex type PassengerLoadingBridgeType
	PassengerLoadingBridge *aixmfeatures.PassengerLoadingBridgeType `xml:"PassengerLoadingBridge,omitempty"`

	// generated from element "Road" of complex type RoadType
	Road *aixmfeatures.RoadType `xml:"Road,omitempty"`

	// generated from element "TouchDownLiftOff" of complex type TouchDownLiftOffType
	TouchDownLiftOff *aixmfeatures.TouchDownLiftOffType `xml:"TouchDownLiftOff,omitempty"`

	// generated from element "ApproachLightingSystem" of complex type ApproachLightingSystemType
	ApproachLightingSystem *aixmfeatures.ApproachLightingSystemType `xml:"ApproachLightingSystem,omitempty"`

	// generated from element "ApronLightSystem" of complex type ApronLightSystemType
	ApronLightSystem *aixmfeatures.ApronLightSystemType `xml:"ApronLightSystem,omitempty"`

	// generated from element "GuidanceLineLightSystem" of complex type GuidanceLineLightSystemType
	GuidanceLineLightSystem *aixmfeatures.GuidanceLineLightSystemType `xml:"GuidanceLineLightSystem,omitempty"`

	// generated from element "RunwayDirectionLightSystem" of complex type RunwayDirectionLightSystemType
	RunwayDirectionLightSystem *aixmfeatures.RunwayDirectionLightSystemType `xml:"RunwayDirectionLightSystem,omitempty"`

	// generated from element "RunwayProtectAreaLightSystem" of complex type RunwayProtectAreaLightSystemType
	RunwayProtectAreaLightSystem *aixmfeatures.RunwayProtectAreaLightSystemType `xml:"RunwayProtectAreaLightSystem,omitempty"`

	// generated from element "TaxiHoldingPositionLightSystem" of complex type TaxiHoldingPositionLightSystemType
	TaxiHoldingPositionLightSystem *aixmfeatures.TaxiHoldingPositionLightSystemType `xml:"TaxiHoldingPositionLightSystem,omitempty"`

	// generated from element "TaxiwayLightSystem" of complex type TaxiwayLightSystemType
	TaxiwayLightSystem *aixmfeatures.TaxiwayLightSystemType `xml:"TaxiwayLightSystem,omitempty"`

	// generated from element "TouchDownLiftOffLightSystem" of complex type TouchDownLiftOffLightSystemType
	TouchDownLiftOffLightSystem *aixmfeatures.TouchDownLiftOffLightSystemType `xml:"TouchDownLiftOffLightSystem,omitempty"`

	// generated from element "VisualGlideSlopeIndicator" of complex type VisualGlideSlopeIndicatorType
	VisualGlideSlopeIndicator *aixmfeatures.VisualGlideSlopeIndicatorType `xml:"VisualGlideSlopeIndicator,omitempty"`

	// generated from element "AirportProtectionAreaMarking" of complex type AirportProtectionAreaMarkingType
	AirportProtectionAreaMarking *aixmfeatures.AirportProtectionAreaMarkingType `xml:"AirportProtectionAreaMarking,omitempty"`

	// generated from element "ApronMarking" of complex type ApronMarkingType
	ApronMarking *aixmfeatures.ApronMarkingType `xml:"ApronMarking,omitempty"`

	// generated from element "DeicingAreaMarking" of complex type DeicingAreaMarkingType
	DeicingAreaMarking *aixmfeatures.DeicingAreaMarkingType `xml:"DeicingAreaMarking,omitempty"`

	// generated from element "GuidanceLineMarking" of complex type GuidanceLineMarkingType
	GuidanceLineMarking *aixmfeatures.GuidanceLineMarkingType `xml:"GuidanceLineMarking,omitempty"`

	// generated from element "RunwayMarking" of complex type RunwayMarkingType
	RunwayMarking *aixmfeatures.RunwayMarkingType `xml:"RunwayMarking,omitempty"`

	// generated from element "StandMarking" of complex type StandMarkingType
	StandMarking *aixmfeatures.StandMarkingType `xml:"StandMarking,omitempty"`

	// generated from element "TaxiHoldingPositionMarking" of complex type TaxiHoldingPositionMarkingType
	TaxiHoldingPositionMarking *aixmfeatures.TaxiHoldingPositionMarkingType `xml:"TaxiHoldingPositionMarking,omitempty"`

	// generated from element "TaxiwayMarking" of complex type TaxiwayMarkingType
	TaxiwayMarking *aixmfeatures.TaxiwayMarkingType `xml:"TaxiwayMarking,omitempty"`

	// generated from element "TouchDownLiftOffMarking" of complex type TouchDownLiftOffMarkingType
	TouchDownLiftOffMarking *aixmfeatures.TouchDownLiftOffMarkingType `xml:"TouchDownLiftOffMarking,omitempty"`

	// generated from element "ArrestingGear" of complex type ArrestingGearType
	ArrestingGear *aixmfeatures.ArrestingGearType `xml:"ArrestingGear,omitempty"`

	// generated from element "Runway" of complex type RunwayType
	Runway *aixmfeatures.RunwayType `xml:"Runway,omitempty"`

	// generated from element "RunwayBlastPad" of complex type RunwayBlastPadType
	RunwayBlastPad *aixmfeatures.RunwayBlastPadType `xml:"RunwayBlastPad,omitempty"`

	// generated from element "RunwayCentrelinePoint" of complex type RunwayCentrelinePointType
	RunwayCentrelinePoint *aixmfeatures.RunwayCentrelinePointType `xml:"RunwayCentrelinePoint,omitempty"`

	// generated from element "RunwayDirection" of complex type RunwayDirectionType
	RunwayDirection *aixmfeatures.RunwayDirectionType `xml:"RunwayDirection,omitempty"`

	// generated from element "RunwayElement" of complex type RunwayElementType
	RunwayElement *aixmfeatures.RunwayElementType `xml:"RunwayElement,omitempty"`

	// generated from element "RunwayVisualRange" of complex type RunwayVisualRangeType
	RunwayVisualRange *aixmfeatures.RunwayVisualRangeType `xml:"RunwayVisualRange,omitempty"`

	// generated from element "FloatingDockSite" of complex type FloatingDockSiteType
	FloatingDockSite *aixmfeatures.FloatingDockSiteType `xml:"FloatingDockSite,omitempty"`

	// generated from element "MarkingBuoy" of complex type MarkingBuoyType
	MarkingBuoy *aixmfeatures.MarkingBuoyType `xml:"MarkingBuoy,omitempty"`

	// generated from element "SeaplaneLandingArea" of complex type SeaplaneLandingAreaType
	SeaplaneLandingArea *aixmfeatures.SeaplaneLandingAreaType `xml:"SeaplaneLandingArea,omitempty"`

	// generated from element "SeaplaneRampSite" of complex type SeaplaneRampSiteType
	SeaplaneRampSite *aixmfeatures.SeaplaneRampSiteType `xml:"SeaplaneRampSite,omitempty"`

	// generated from element "GuidanceLine" of complex type GuidanceLineType
	GuidanceLine *aixmfeatures.GuidanceLineType `xml:"GuidanceLine,omitempty"`

	// generated from element "TaxiHoldingPosition" of complex type TaxiHoldingPositionType
	TaxiHoldingPosition *aixmfeatures.TaxiHoldingPositionType `xml:"TaxiHoldingPosition,omitempty"`

	// generated from element "Taxiway" of complex type TaxiwayType
	Taxiway *aixmfeatures.TaxiwayType `xml:"Taxiway,omitempty"`

	// generated from element "TaxiwayElement" of complex type TaxiwayElementType
	TaxiwayElement *aixmfeatures.TaxiwayElementType `xml:"TaxiwayElement,omitempty"`

	// generated from element "Airspace" of complex type AirspaceType
	Airspace *aixmfeatures.AirspaceType `xml:"Airspace,omitempty"`

	// generated from element "AuthorityForAirspace" of complex type AuthorityForAirspaceType
	AuthorityForAirspace *aixmfeatures.AuthorityForAirspaceType `xml:"AuthorityForAirspace,omitempty"`

	// generated from element "GeoBorder" of complex type GeoBorderType
	GeoBorder *aixmfeatures.GeoBorderType `xml:"GeoBorder,omitempty"`

	// generated from element "HoldingPattern" of complex type HoldingPatternType
	HoldingPattern *aixmfeatures.HoldingPatternType `xml:"HoldingPattern,omitempty"`

	// generated from element "UnplannedHolding" of complex type UnplannedHoldingType
	UnplannedHolding *aixmfeatures.UnplannedHoldingType `xml:"UnplannedHolding,omitempty"`

	// generated from element "Navaid" of complex type NavaidType
	Navaid *aixmfeatures.NavaidType `xml:"Navaid,omitempty"`

	// generated from element "Azimuth" of complex type AzimuthType
	Azimuth *aixmfeatures.AzimuthType `xml:"Azimuth,omitempty"`

	// generated from element "DirectionFinder" of complex type DirectionFinderType
	DirectionFinder *aixmfeatures.DirectionFinderType `xml:"DirectionFinder,omitempty"`

	// generated from element "DME" of complex type DMEType
	DME *aixmfeatures.DMEType `xml:"DME,omitempty"`

	// generated from element "Elevation" of complex type ElevationType
	Elevation *aixmfeatures.ElevationType `xml:"Elevation,omitempty"`

	// generated from element "Glidepath" of complex type GlidepathType
	Glidepath *aixmfeatures.GlidepathType `xml:"Glidepath,omitempty"`

	// generated from element "Localizer" of complex type LocalizerType
	Localizer *aixmfeatures.LocalizerType `xml:"Localizer,omitempty"`

	// generated from element "MarkerBeacon" of complex type MarkerBeaconType
	MarkerBeacon *aixmfeatures.MarkerBeaconType `xml:"MarkerBeacon,omitempty"`

	// generated from element "NDB" of complex type NDBType
	NDB *aixmfeatures.NDBType `xml:"NDB,omitempty"`

	// generated from element "SDF" of complex type SDFType
	SDF *aixmfeatures.SDFType `xml:"SDF,omitempty"`

	// generated from element "TACAN" of complex type TACANType
	TACAN *aixmfeatures.TACANType `xml:"TACAN,omitempty"`

	// generated from element "VOR" of complex type VORType
	VOR *aixmfeatures.VORType `xml:"VOR,omitempty"`

	// generated from element "CheckpointINS" of complex type CheckpointINSType
	CheckpointINS *aixmfeatures.CheckpointINSType `xml:"CheckpointINS,omitempty"`

	// generated from element "CheckpointVOR" of complex type CheckpointVORType
	CheckpointVOR *aixmfeatures.CheckpointVORType `xml:"CheckpointVOR,omitempty"`

	// generated from element "SpecialNavigationStation" of complex type SpecialNavigationStationType
	SpecialNavigationStation *aixmfeatures.SpecialNavigationStationType `xml:"SpecialNavigationStation,omitempty"`

	// generated from element "SpecialNavigationSystem" of complex type SpecialNavigationSystemType
	SpecialNavigationSystem *aixmfeatures.SpecialNavigationSystemType `xml:"SpecialNavigationSystem,omitempty"`

	// generated from element "AngleIndication" of complex type AngleIndicationType
	AngleIndication *aixmfeatures.AngleIndicationType `xml:"AngleIndication,omitempty"`

	// generated from element "DistanceIndication" of complex type DistanceIndicationType
	DistanceIndication *aixmfeatures.DistanceIndicationType `xml:"DistanceIndication,omitempty"`

	// generated from element "DesignatedPoint" of complex type DesignatedPointType
	DesignatedPoint *aixmfeatures.DesignatedPointType `xml:"DesignatedPoint,omitempty"`

	// generated from element "SignificantPointInAirspace" of complex type SignificantPointInAirspaceType
	SignificantPointInAirspace *aixmfeatures.SignificantPointInAirspaceType `xml:"SignificantPointInAirspace,omitempty"`

	// generated from element "AeronauticalGroundLight" of complex type AeronauticalGroundLightType
	AeronauticalGroundLight *aixmfeatures.AeronauticalGroundLightType `xml:"AeronauticalGroundLight,omitempty"`

	// generated from element "ObstacleArea" of complex type ObstacleAreaType
	ObstacleArea *aixmfeatures.ObstacleAreaType `xml:"ObstacleArea,omitempty"`

	// generated from element "VerticalStructure" of complex type VerticalStructureType
	VerticalStructure *aixmfeatures.VerticalStructureType `xml:"VerticalStructure,omitempty"`

	// generated from element "OrganisationAuthority" of complex type OrganisationAuthorityType
	OrganisationAuthority *aixmfeatures.OrganisationAuthorityType `xml:"OrganisationAuthority,omitempty"`

	// generated from element "Unit" of complex type UnitType
	Unit *aixmfeatures.UnitType `xml:"Unit,omitempty"`

	// generated from element "TerminalArrivalArea" of complex type TerminalArrivalAreaType
	TerminalArrivalArea *aixmfeatures.TerminalArrivalAreaType `xml:"TerminalArrivalArea,omitempty"`

	// generated from element "CirclingArea" of complex type CirclingAreaType
	CirclingArea *aixmfeatures.CirclingAreaType `xml:"CirclingArea,omitempty"`

	// generated from element "NavigationArea" of complex type NavigationAreaType
	NavigationArea *aixmfeatures.NavigationAreaType `xml:"NavigationArea,omitempty"`

	// generated from element "SafeAltitudeArea" of complex type SafeAltitudeAreaType
	SafeAltitudeArea *aixmfeatures.SafeAltitudeAreaType `xml:"SafeAltitudeArea,omitempty"`

	// generated from element "NavigationAreaRestriction" of complex type NavigationAreaRestrictionType
	NavigationAreaRestriction *aixmfeatures.NavigationAreaRestrictionType `xml:"NavigationAreaRestriction,omitempty"`

	// generated from element "InstrumentApproachProcedure" of complex type InstrumentApproachProcedureType
	InstrumentApproachProcedure *aixmfeatures.InstrumentApproachProcedureType `xml:"InstrumentApproachProcedure,omitempty"`

	// generated from element "StandardInstrumentArrival" of complex type StandardInstrumentArrivalType
	StandardInstrumentArrival *aixmfeatures.StandardInstrumentArrivalType `xml:"StandardInstrumentArrival,omitempty"`

	// generated from element "StandardInstrumentDeparture" of complex type StandardInstrumentDepartureType
	StandardInstrumentDeparture *aixmfeatures.StandardInstrumentDepartureType `xml:"StandardInstrumentDeparture,omitempty"`

	// generated from element "ProcedureDME" of complex type ProcedureDMEType
	ProcedureDME *aixmfeatures.ProcedureDMEType `xml:"ProcedureDME,omitempty"`

	// generated from element "ArrivalFeederLeg" of complex type ArrivalFeederLegType
	ArrivalFeederLeg *aixmfeatures.ArrivalFeederLegType `xml:"ArrivalFeederLeg,omitempty"`

	// generated from element "FinalLeg" of complex type FinalLegType
	FinalLeg *aixmfeatures.FinalLegType `xml:"FinalLeg,omitempty"`

	// generated from element "InitialLeg" of complex type InitialLegType
	InitialLeg *aixmfeatures.InitialLegType `xml:"InitialLeg,omitempty"`

	// generated from element "IntermediateLeg" of complex type IntermediateLegType
	IntermediateLeg *aixmfeatures.IntermediateLegType `xml:"IntermediateLeg,omitempty"`

	// generated from element "MissedApproachLeg" of complex type MissedApproachLegType
	MissedApproachLeg *aixmfeatures.MissedApproachLegType `xml:"MissedApproachLeg,omitempty"`

	// generated from element "ArrivalLeg" of complex type ArrivalLegType
	ArrivalLeg *aixmfeatures.ArrivalLegType `xml:"ArrivalLeg,omitempty"`

	// generated from element "DepartureLeg" of complex type DepartureLegType
	DepartureLeg *aixmfeatures.DepartureLegType `xml:"DepartureLeg,omitempty"`

	// generated from element "ChangeOverPoint" of complex type ChangeOverPointType
	ChangeOverPoint *aixmfeatures.ChangeOverPointType `xml:"ChangeOverPoint,omitempty"`

	// generated from element "Route" of complex type RouteType
	Route *aixmfeatures.RouteType `xml:"Route,omitempty"`

	// generated from element "RouteDME" of complex type RouteDMEType
	RouteDME *aixmfeatures.RouteDMEType `xml:"RouteDME,omitempty"`

	// generated from element "RouteSegment" of complex type RouteSegmentType
	RouteSegment *aixmfeatures.RouteSegmentType `xml:"RouteSegment,omitempty"`

	// generated from element "AirspaceBorderCrossing" of complex type AirspaceBorderCrossingType
	AirspaceBorderCrossing *aixmfeatures.AirspaceBorderCrossingType `xml:"AirspaceBorderCrossing,omitempty"`

	// generated from element "FlightRestriction" of complex type FlightRestrictionType
	FlightRestriction *aixmfeatures.FlightRestrictionType `xml:"FlightRestriction,omitempty"`

	// generated from element "RulesProcedures" of complex type RulesProceduresType
	RulesProcedures *aixmfeatures.RulesProceduresType `xml:"RulesProcedures,omitempty"`

	// generated from element "PilotControlledLighting" of complex type PilotControlledLightingType
	PilotControlledLighting *aixmfeatures.PilotControlledLightingType `xml:"PilotControlledLighting,omitempty"`

	// generated from element "RadioCommunicationChannel" of complex type RadioCommunicationChannelType
	RadioCommunicationChannel *aixmfeatures.RadioCommunicationChannelType `xml:"RadioCommunicationChannel,omitempty"`

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

	// generated from element "AirTrafficManagementService" of complex type AirTrafficManagementServiceType
	AirTrafficManagementService *aixmfeatures.AirTrafficManagementServiceType `xml:"AirTrafficManagementService,omitempty"`

	// generated from element "InformationService" of complex type InformationServiceType
	InformationService *aixmfeatures.InformationServiceType `xml:"InformationService,omitempty"`

	// generated from element "SearchRescueService" of complex type SearchRescueServiceType
	SearchRescueService *aixmfeatures.SearchRescueServiceType `xml:"SearchRescueService,omitempty"`

	// generated from element "AirTrafficControlService" of complex type AirTrafficControlServiceType
	AirTrafficControlService *aixmfeatures.AirTrafficControlServiceType `xml:"AirTrafficControlService,omitempty"`

	// generated from element "GroundTrafficControlService" of complex type GroundTrafficControlServiceType
	GroundTrafficControlService *aixmfeatures.GroundTrafficControlServiceType `xml:"GroundTrafficControlService,omitempty"`

	// generated from element "RadioFrequencyArea" of complex type RadioFrequencyAreaType
	RadioFrequencyArea *aixmfeatures.RadioFrequencyAreaType `xml:"RadioFrequencyArea,omitempty"`

	// generated from element "SpecialDate" of complex type SpecialDateType
	SpecialDate *aixmfeatures.SpecialDateType `xml:"SpecialDate,omitempty"`

	// generated from element "StandardLevelColumn" of complex type StandardLevelColumnType
	StandardLevelColumn *aixmfeatures.StandardLevelColumnType `xml:"StandardLevelColumn,omitempty"`

	// generated from element "StandardLevelSector" of complex type StandardLevelSectorType
	StandardLevelSector *aixmfeatures.StandardLevelSectorType `xml:"StandardLevelSector,omitempty"`

	// generated from element "StandardLevelTable" of complex type StandardLevelTableType
	StandardLevelTable *aixmfeatures.StandardLevelTableType `xml:"StandardLevelTable,omitempty"`

	// generated from element "HoldingAssessment" of complex type HoldingAssessmentType
	HoldingAssessment *aixmfeatures.HoldingAssessmentType `xml:"HoldingAssessment,omitempty"`

	// generated from element "PrecisionApproachRadar" of complex type PrecisionApproachRadarType
	PrecisionApproachRadar *aixmfeatures.PrecisionApproachRadarType `xml:"PrecisionApproachRadar,omitempty"`

	// generated from element "PrimarySurveillanceRadar" of complex type PrimarySurveillanceRadarType
	PrimarySurveillanceRadar *aixmfeatures.PrimarySurveillanceRadarType `xml:"PrimarySurveillanceRadar,omitempty"`

	// generated from element "SecondarySurveillanceRadar" of complex type SecondarySurveillanceRadarType
	SecondarySurveillanceRadar *aixmfeatures.SecondarySurveillanceRadarType `xml:"SecondarySurveillanceRadar,omitempty"`

	// generated from element "RadarSystem" of complex type RadarSystemType
	RadarSystem *aixmfeatures.RadarSystemType `xml:"RadarSystem,omitempty"`
}
