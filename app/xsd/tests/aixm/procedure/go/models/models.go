// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeICAOCountryType
var _ *aixmfeatures.AerialRefuellingPointType
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

	// generated from element "StandardInstrumentDeparture" of complex type StandardInstrumentDepartureType
	StandardInstrumentDeparture *aixmfeatures.StandardInstrumentDepartureType `xml:"StandardInstrumentDeparture,omitempty"`

	// generated from element "StandardInstrumentArrival" of complex type StandardInstrumentArrivalType
	StandardInstrumentArrival *aixmfeatures.StandardInstrumentArrivalType `xml:"StandardInstrumentArrival,omitempty"`

	// generated from element "InstrumentApproachProcedure" of complex type InstrumentApproachProcedureType
	InstrumentApproachProcedure *aixmfeatures.InstrumentApproachProcedureType `xml:"InstrumentApproachProcedure,omitempty"`

	// generated from element "NavigationArea" of complex type NavigationAreaType
	NavigationArea *aixmfeatures.NavigationAreaType `xml:"NavigationArea,omitempty"`

	// generated from element "TerminalArrivalArea" of complex type TerminalArrivalAreaType
	TerminalArrivalArea *aixmfeatures.TerminalArrivalAreaType `xml:"TerminalArrivalArea,omitempty"`

	// generated from element "CirclingArea" of complex type CirclingAreaType
	CirclingArea *aixmfeatures.CirclingAreaType `xml:"CirclingArea,omitempty"`

	// generated from element "DepartureLeg" of complex type DepartureLegType
	DepartureLeg *aixmfeatures.DepartureLegType `xml:"DepartureLeg,omitempty"`

	// generated from element "ArrivalLeg" of complex type ArrivalLegType
	ArrivalLeg *aixmfeatures.ArrivalLegType `xml:"ArrivalLeg,omitempty"`

	// generated from element "ArrivalFeederLeg" of complex type ArrivalFeederLegType
	ArrivalFeederLeg *aixmfeatures.ArrivalFeederLegType `xml:"ArrivalFeederLeg,omitempty"`

	// generated from element "InitialLeg" of complex type InitialLegType
	InitialLeg *aixmfeatures.InitialLegType `xml:"InitialLeg,omitempty"`

	// generated from element "IntermediateLeg" of complex type IntermediateLegType
	IntermediateLeg *aixmfeatures.IntermediateLegType `xml:"IntermediateLeg,omitempty"`

	// generated from element "FinalLeg" of complex type FinalLegType
	FinalLeg *aixmfeatures.FinalLegType `xml:"FinalLeg,omitempty"`

	// generated from element "MissedApproachLeg" of complex type MissedApproachLegType
	MissedApproachLeg *aixmfeatures.MissedApproachLegType `xml:"MissedApproachLeg,omitempty"`

	// generated from element "ProcedureDME" of complex type ProcedureDMEType
	ProcedureDME *aixmfeatures.ProcedureDMEType `xml:"ProcedureDME,omitempty"`

	// generated from element "NavigationAreaRestriction" of complex type NavigationAreaRestrictionType
	NavigationAreaRestriction *aixmfeatures.NavigationAreaRestrictionType `xml:"NavigationAreaRestriction,omitempty"`

	// generated from element "SafeAltitudeArea" of complex type SafeAltitudeAreaType
	SafeAltitudeArea *aixmfeatures.SafeAltitudeAreaType `xml:"SafeAltitudeArea,omitempty"`

	// generated from element "HoldingPattern" of complex type HoldingPatternType
	HoldingPattern *aixmfeatures.HoldingPatternType `xml:"HoldingPattern,omitempty"`

	// generated from element "UnplannedHolding" of complex type UnplannedHoldingType
	UnplannedHolding *aixmfeatures.UnplannedHoldingType `xml:"UnplannedHolding,omitempty"`

	// generated from element "HoldingAssessment" of complex type HoldingAssessmentType
	HoldingAssessment *aixmfeatures.HoldingAssessmentType `xml:"HoldingAssessment,omitempty"`
}
