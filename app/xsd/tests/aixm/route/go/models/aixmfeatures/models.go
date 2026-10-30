// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAerialRefuellingPointType
var _ *gml.CodeType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// AerialRefuellingPropertyType Named source named complex type "AerialRefuellingPropertyType"
type AerialRefuellingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AerialRefuellingType Named source named complex type "AerialRefuellingType"
type AerialRefuellingType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type AerialRefuellingTimeSlicePropertyType
	TimeSlice []*AerialRefuellingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AerialRefuellingTimeSlicePropertyType Named source named complex type "AerialRefuellingTimeSlicePropertyType"
type AerialRefuellingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AerialRefuellingTimeSlice" of complex type AerialRefuellingTimeSliceType
	AerialRefuellingTimeSlice *AerialRefuellingTimeSliceType `xml:"AerialRefuellingTimeSlice,omitempty"`
}

// AerialRefuellingTimeSliceType Named source named complex type "AerialRefuellingTimeSliceType"
type AerialRefuellingTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from element "designatorPrefix" of complex type CodeAerialRefuellingPrefixType
	DesignatorPrefix *aixmdatatypes.CodeAerialRefuellingPrefixType `xml:"designatorPrefix,omitempty"`

	// generated from element "designatorNumber" of complex type NoNumberType
	DesignatorNumber *aixmdatatypes.NoNumberType `xml:"designatorNumber,omitempty"`

	// generated from element "designatorSuffix" of complex type TextDesignatorType
	DesignatorSuffix *aixmdatatypes.TextDesignatorType `xml:"designatorSuffix,omitempty"`

	// generated from element "designatorDirection" of complex type CodeCardinalDirectionType
	DesignatorDirection *aixmdatatypes.CodeCardinalDirectionType `xml:"designatorDirection,omitempty"`

	// generated from element "type" of complex type CodeAerialRefuellingType
	Type_XSD *aixmdatatypes.CodeAerialRefuellingType `xml:"type,omitempty"`

	// generated from element "radarBeaconSetting" of complex type NoNumberType
	RadarBeaconSetting *aixmdatatypes.NoNumberType `xml:"radarBeaconSetting,omitempty"`

	// generated from element "xbandRadarSetting" of complex type NoNumberType
	XbandRadarSetting *aixmdatatypes.NoNumberType `xml:"xbandRadarSetting,omitempty"`

	// generated from element "tankerChannel" of complex type CodeTACANChannelType
	TankerChannel *aixmdatatypes.CodeTACANChannelType `xml:"tankerChannel,omitempty"`

	// generated from element "receiverChannel" of complex type CodeTACANChannelType
	ReceiverChannel *aixmdatatypes.CodeTACANChannelType `xml:"receiverChannel,omitempty"`

	// generated from element "helicopterRoute" of complex type CodeYesNoType
	HelicopterRoute *aixmdatatypes.CodeYesNoType `xml:"helicopterRoute,omitempty"`

	// generated from element "specialRefuelling" of complex type CodeYesNoType
	SpecialRefuelling *aixmdatatypes.CodeYesNoType `xml:"specialRefuelling,omitempty"`

	// generated from element "bidirectionalUse" of complex type CodeYesNoType
	BidirectionalUse *aixmdatatypes.CodeYesNoType `xml:"bidirectionalUse,omitempty"`

	// generated from element "reverseDirectionTurn" of complex type CodeDirectionTurnType
	ReverseDirectionTurn *aixmdatatypes.CodeDirectionTurnType `xml:"reverseDirectionTurn,omitempty"`

	// generated from element "availability" of complex type RouteAvailabilityPropertyType
	Availability []*RouteAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "protectingAirspace" of complex type AirspacePropertyType
	ProtectingAirspace []*AirspacePropertyType `xml:"protectingAirspace,omitempty"`

	// generated from element "track" of complex type AerialRefuellingTrackPropertyType
	Track []*AerialRefuellingTrackPropertyType `xml:"track,omitempty"`

	// generated from element "anchor" of complex type AerialRefuellingAnchorPropertyType
	Anchor []*AerialRefuellingAnchorPropertyType `xml:"anchor,omitempty"`

	// generated from element "oppositeTrack" of complex type AerialRefuellingPropertyType
	OppositeTrack *AerialRefuellingPropertyType `xml:"oppositeTrack,omitempty"`

	// generated from element "managingOrganisation" of complex type AuthorityForAerialRefuellingPropertyType
	ManagingOrganisation []*AuthorityForAerialRefuellingPropertyType `xml:"managingOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension `xml:"extension,omitempty"`
}

// A_extension Named source within outer element "extension"
type A_extension struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAerialRefuellingExtension" of complex type AbstractExtensionType
	AbstractAerialRefuellingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAerialRefuellingExtension,omitempty"`
}

// Group_AerialRefuellingPropertyGroup UnNamed source named group "AerialRefuellingPropertyGroup"
type Group_AerialRefuellingPropertyGroup struct {

	// insertion point for fields

	// generated from element "designatorPrefix" of complex type CodeAerialRefuellingPrefixType
	DesignatorPrefix *aixmdatatypes.CodeAerialRefuellingPrefixType `xml:"designatorPrefix,omitempty"`

	// generated from element "designatorNumber" of complex type NoNumberType
	DesignatorNumber *aixmdatatypes.NoNumberType `xml:"designatorNumber,omitempty"`

	// generated from element "designatorSuffix" of complex type TextDesignatorType
	DesignatorSuffix *aixmdatatypes.TextDesignatorType `xml:"designatorSuffix,omitempty"`

	// generated from element "designatorDirection" of complex type CodeCardinalDirectionType
	DesignatorDirection *aixmdatatypes.CodeCardinalDirectionType `xml:"designatorDirection,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeAerialRefuellingType
	Type_XSD *aixmdatatypes.CodeAerialRefuellingType `xml:"type,omitempty"`

	// generated from element "radarBeaconSetting" of complex type NoNumberType
	RadarBeaconSetting *aixmdatatypes.NoNumberType `xml:"radarBeaconSetting,omitempty"`

	// generated from element "xbandRadarSetting" of complex type NoNumberType
	XbandRadarSetting *aixmdatatypes.NoNumberType `xml:"xbandRadarSetting,omitempty"`

	// generated from element "tankerChannel" of complex type CodeTACANChannelType
	TankerChannel *aixmdatatypes.CodeTACANChannelType `xml:"tankerChannel,omitempty"`

	// generated from element "receiverChannel" of complex type CodeTACANChannelType
	ReceiverChannel *aixmdatatypes.CodeTACANChannelType `xml:"receiverChannel,omitempty"`

	// generated from element "helicopterRoute" of complex type CodeYesNoType
	HelicopterRoute *aixmdatatypes.CodeYesNoType `xml:"helicopterRoute,omitempty"`

	// generated from element "specialRefuelling" of complex type CodeYesNoType
	SpecialRefuelling *aixmdatatypes.CodeYesNoType `xml:"specialRefuelling,omitempty"`

	// generated from element "bidirectionalUse" of complex type CodeYesNoType
	BidirectionalUse *aixmdatatypes.CodeYesNoType `xml:"bidirectionalUse,omitempty"`

	// generated from element "reverseDirectionTurn" of complex type CodeDirectionTurnType
	ReverseDirectionTurn *aixmdatatypes.CodeDirectionTurnType `xml:"reverseDirectionTurn,omitempty"`

	// generated from element "availability" of complex type RouteAvailabilityPropertyType
	Availability []*RouteAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "protectingAirspace" of complex type AirspacePropertyType
	ProtectingAirspace []*AirspacePropertyType `xml:"protectingAirspace,omitempty"`

	// generated from element "track" of complex type AerialRefuellingTrackPropertyType
	Track []*AerialRefuellingTrackPropertyType `xml:"track,omitempty"`

	// generated from element "anchor" of complex type AerialRefuellingAnchorPropertyType
	Anchor []*AerialRefuellingAnchorPropertyType `xml:"anchor,omitempty"`

	// generated from element "oppositeTrack" of complex type AerialRefuellingPropertyType
	OppositeTrack *AerialRefuellingPropertyType `xml:"oppositeTrack,omitempty"`

	// generated from element "managingOrganisation" of complex type AuthorityForAerialRefuellingPropertyType
	ManagingOrganisation []*AuthorityForAerialRefuellingPropertyType `xml:"managingOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AerialRefuellingAnchorType Named source named complex type "AerialRefuellingAnchorType"
type AerialRefuellingAnchorType struct {
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

	// generated from group "aixm:AerialRefuellingAnchorPropertyGroup" with order 416 depth 2
	Group_AerialRefuellingAnchorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_1 `xml:"extension,omitempty"`
}

// A_extension_1 Named source within outer element "extension"
type A_extension_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAerialRefuellingAnchorExtension" of complex type AbstractExtensionType
	AbstractAerialRefuellingAnchorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAerialRefuellingAnchorExtension,omitempty"`
}

// Group_AerialRefuellingAnchorPropertyGroup UnNamed source named group "AerialRefuellingAnchorPropertyGroup"
type Group_AerialRefuellingAnchorPropertyGroup struct {

	// insertion point for fields

	// generated from element "outboundCourse" of complex type ValBearingType
	OutboundCourse *aixmdatatypes.ValBearingType `xml:"outboundCourse,omitempty"`

	// generated from element "outboundCourseType" of complex type CodeCourseType
	OutboundCourseType *aixmdatatypes.CodeCourseType `xml:"outboundCourseType,omitempty"`

	// generated from element "inboundCourse" of complex type ValBearingType
	InboundCourse *aixmdatatypes.ValBearingType `xml:"inboundCourse,omitempty"`

	// generated from element "turnDirection" of complex type CodeDirectionTurnType
	TurnDirection *aixmdatatypes.CodeDirectionTurnType `xml:"turnDirection,omitempty"`

	// generated from element "speedLimit" of complex type ValSpeedType
	SpeedLimit *aixmdatatypes.ValSpeedType `xml:"speedLimit,omitempty"`

	// generated from element "legSeparation" of complex type ValDistanceType
	LegSeparation *aixmdatatypes.ValDistanceType `xml:"legSeparation,omitempty"`

	// generated from element "legLength" of complex type ValDistanceType
	LegLength *aixmdatatypes.ValDistanceType `xml:"legLength,omitempty"`

	// generated from element "refuellingBaseLevel" of complex type ValDistanceVerticalType
	RefuellingBaseLevel *aixmdatatypes.ValDistanceVerticalType `xml:"refuellingBaseLevel,omitempty"`

	// generated from element "refuellingBaseLevelReference" of complex type CodeVerticalReferenceType
	RefuellingBaseLevelReference *aixmdatatypes.CodeVerticalReferenceType `xml:"refuellingBaseLevelReference,omitempty"`

	// generated from element "extent" of complex type SurfacePropertyType
	Extent *SurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "verticalExtent" of complex type AirspaceLayerPropertyType
	VerticalExtent []*AirspaceLayerPropertyType `xml:"verticalExtent,omitempty"`

	// generated from element "point" of complex type AerialRefuellingPointPropertyType
	Point []*AerialRefuellingPointPropertyType `xml:"point,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AerialRefuellingAnchorPropertyType Named source named complex type "AerialRefuellingAnchorPropertyType"
type AerialRefuellingAnchorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AerialRefuellingAnchor" of complex type AerialRefuellingAnchorType
	AerialRefuellingAnchor *AerialRefuellingAnchorType `xml:"AerialRefuellingAnchor,omitempty"`
}

// AerialRefuellingPointType Named source named complex type "AerialRefuellingPointType"
type AerialRefuellingPointType struct {
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

	// generated from group "aixm:SegmentPointPropertyGroup" with order 444 depth 2
	Group_SegmentPointPropertyGroup

	// generated from group "aixm:AerialRefuellingPointPropertyGroup" with order 445 depth 2
	Group_AerialRefuellingPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSegmentPointExtension" of complex type AbstractExtensionType
	AbstractSegmentPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentPointExtension,omitempty"`

	// generated from element "AbstractAerialRefuellingPointExtension" of complex type AbstractExtensionType
	AbstractAerialRefuellingPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAerialRefuellingPointExtension,omitempty"`
}

// Group_AerialRefuellingPointPropertyGroup UnNamed source named group "AerialRefuellingPointPropertyGroup"
type Group_AerialRefuellingPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "sequence" of complex type NoSequenceType
	Sequence *aixmdatatypes.NoSequenceType `xml:"sequence,omitempty"`

	// generated from element "usageType" of complex type CodeAerialRefuellingPointType
	UsageType *aixmdatatypes.CodeAerialRefuellingPointType `xml:"usageType,omitempty"`
}

// AerialRefuellingPointPropertyType Named source named complex type "AerialRefuellingPointPropertyType"
type AerialRefuellingPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AerialRefuellingPoint" of complex type AerialRefuellingPointType
	AerialRefuellingPoint *AerialRefuellingPointType `xml:"AerialRefuellingPoint,omitempty"`
}

// AerialRefuellingTrackType Named source named complex type "AerialRefuellingTrackType"
type AerialRefuellingTrackType struct {
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

	// generated from group "aixm:AerialRefuellingTrackPropertyGroup" with order 463 depth 2
	Group_AerialRefuellingTrackPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAerialRefuellingTrackExtension" of complex type AbstractExtensionType
	AbstractAerialRefuellingTrackExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAerialRefuellingTrackExtension,omitempty"`
}

// Group_AerialRefuellingTrackPropertyGroup UnNamed source named group "AerialRefuellingTrackPropertyGroup"
type Group_AerialRefuellingTrackPropertyGroup struct {

	// insertion point for fields

	// generated from element "extent" of complex type CurvePropertyType
	Extent *CurvePropertyType `xml:"extent,omitempty"`

	// generated from element "point" of complex type AerialRefuellingPointPropertyType
	Point []*AerialRefuellingPointPropertyType `xml:"point,omitempty"`

	// generated from element "verticalExtent" of complex type AirspaceLayerPropertyType
	VerticalExtent []*AirspaceLayerPropertyType `xml:"verticalExtent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AerialRefuellingTrackPropertyType Named source named complex type "AerialRefuellingTrackPropertyType"
type AerialRefuellingTrackPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AerialRefuellingTrack" of complex type AerialRefuellingTrackType
	AerialRefuellingTrack *AerialRefuellingTrackType `xml:"AerialRefuellingTrack,omitempty"`
}

// AuthorityForAerialRefuellingType Named source named complex type "AuthorityForAerialRefuellingType"
type AuthorityForAerialRefuellingType struct {
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

	// generated from group "aixm:AuthorityForAerialRefuellingPropertyGroup" with order 482 depth 2
	Group_AuthorityForAerialRefuellingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAuthorityForAerialRefuellingExtension" of complex type AbstractExtensionType
	AbstractAuthorityForAerialRefuellingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAuthorityForAerialRefuellingExtension,omitempty"`
}

// Group_AuthorityForAerialRefuellingPropertyGroup UnNamed source named group "AuthorityForAerialRefuellingPropertyGroup"
type Group_AuthorityForAerialRefuellingPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAuthorityType
	Type_XSD *aixmdatatypes.CodeAuthorityType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// AuthorityForAerialRefuellingPropertyType Named source named complex type "AuthorityForAerialRefuellingPropertyType"
type AuthorityForAerialRefuellingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AuthorityForAerialRefuelling" of complex type AuthorityForAerialRefuellingType
	AuthorityForAerialRefuelling *AuthorityForAerialRefuellingType `xml:"AuthorityForAerialRefuelling,omitempty"`
}

// AirportHeliportPropertyType Named source named complex type "AirportHeliportPropertyType"
type AirportHeliportPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// TouchDownLiftOffPropertyType Named source named complex type "TouchDownLiftOffPropertyType"
type TouchDownLiftOffPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RunwayCentrelinePointPropertyType Named source named complex type "RunwayCentrelinePointPropertyType"
type RunwayCentrelinePointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AirspacePropertyType Named source named complex type "AirspacePropertyType"
type AirspacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// CurveType Named source named complex type "CurveType"
type CurveType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "segments" of complex type CurveSegmentArrayPropertyType
	Segments *gml.CurveSegmentArrayPropertyType `xml:"segments,omitempty"`

	// generated from group "aixm:CurvePropertyGroup" with order 2753 depth 2
	Group_CurvePropertyGroup
}

// Group_CurvePropertyGroup UnNamed source named group "CurvePropertyGroup"
type Group_CurvePropertyGroup struct {

	// insertion point for fields

	// generated from element "horizontalAccuracy" of complex type ValDistanceType
	HorizontalAccuracy *aixmdatatypes.ValDistanceType `xml:"horizontalAccuracy,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// CurvePropertyType Named source named complex type "CurvePropertyType"
type CurvePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Curve" of complex type CurveType
	Curve *CurveType `xml:"Curve,omitempty"`

	// generated from element "ElevatedCurve" of complex type ElevatedCurveType
	ElevatedCurve *ElevatedCurveType `xml:"ElevatedCurve,omitempty"`
}

// ElevatedCurveType Named source named complex type "ElevatedCurveType"
type ElevatedCurveType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "segments" of complex type CurveSegmentArrayPropertyType
	Segments *gml.CurveSegmentArrayPropertyType `xml:"segments,omitempty"`

	// generated from group "aixm:CurvePropertyGroup" with order 2753 depth 2
	Group_CurvePropertyGroup

	// generated from group "aixm:ElevatedCurvePropertyGroup" with order 2764 depth 2
	Group_ElevatedCurvePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractElevatedCurveExtension" of complex type AbstractExtensionType
	AbstractElevatedCurveExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractElevatedCurveExtension,omitempty"`
}

// Group_ElevatedCurvePropertyGroup UnNamed source named group "ElevatedCurvePropertyGroup"
type Group_ElevatedCurvePropertyGroup struct {

	// insertion point for fields

	// generated from element "elevation" of complex type ValDistanceVerticalType
	Elevation *aixmdatatypes.ValDistanceVerticalType `xml:"elevation,omitempty"`

	// generated from element "geoidUndulation" of complex type ValDistanceSignedType
	GeoidUndulation *aixmdatatypes.ValDistanceSignedType `xml:"geoidUndulation,omitempty"`

	// generated from element "verticalDatum" of complex type CodeVerticalDatumType
	VerticalDatum *aixmdatatypes.CodeVerticalDatumType `xml:"verticalDatum,omitempty"`

	// generated from element "verticalAccuracy" of complex type ValDistanceType
	VerticalAccuracy *aixmdatatypes.ValDistanceType `xml:"verticalAccuracy,omitempty"`
}

// ElevatedPointType Named source named complex type "ElevatedPointType"
type ElevatedPointType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos *gml.DirectPositionType `xml:"pos,omitempty"`

	// generated from group "aixm:PointPropertyGroup" with order 2821 depth 2
	Group_PointPropertyGroup

	// generated from group "aixm:ElevatedPointPropertyGroup" with order 2783 depth 2
	Group_ElevatedPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractElevatedPointExtension" of complex type AbstractExtensionType
	AbstractElevatedPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractElevatedPointExtension,omitempty"`
}

// Group_ElevatedPointPropertyGroup UnNamed source named group "ElevatedPointPropertyGroup"
type Group_ElevatedPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "elevation" of complex type ValDistanceVerticalType
	Elevation *aixmdatatypes.ValDistanceVerticalType `xml:"elevation,omitempty"`

	// generated from element "geoidUndulation" of complex type ValDistanceSignedType
	GeoidUndulation *aixmdatatypes.ValDistanceSignedType `xml:"geoidUndulation,omitempty"`

	// generated from element "verticalDatum" of complex type CodeVerticalDatumType
	VerticalDatum *aixmdatatypes.CodeVerticalDatumType `xml:"verticalDatum,omitempty"`

	// generated from element "verticalAccuracy" of complex type ValDistanceType
	VerticalAccuracy *aixmdatatypes.ValDistanceType `xml:"verticalAccuracy,omitempty"`
}

// ElevatedSurfaceType Named source named complex type "ElevatedSurfaceType"
type ElevatedSurfaceType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "patches" of complex type SurfacePatchArrayPropertyType
	Patches *gml.SurfacePatchArrayPropertyType `xml:"patches,omitempty"`

	// generated from group "aixm:SurfacePropertyGroup" with order 2832 depth 2
	Group_SurfacePropertyGroup

	// generated from group "aixm:ElevatedSurfacePropertyGroup" with order 2802 depth 2
	Group_ElevatedSurfacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractElevatedSurfaceExtension" of complex type AbstractExtensionType
	AbstractElevatedSurfaceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractElevatedSurfaceExtension,omitempty"`
}

// Group_ElevatedSurfacePropertyGroup UnNamed source named group "ElevatedSurfacePropertyGroup"
type Group_ElevatedSurfacePropertyGroup struct {

	// insertion point for fields

	// generated from element "elevation" of complex type ValDistanceVerticalType
	Elevation *aixmdatatypes.ValDistanceVerticalType `xml:"elevation,omitempty"`

	// generated from element "geoidUndulation" of complex type ValDistanceSignedType
	GeoidUndulation *aixmdatatypes.ValDistanceSignedType `xml:"geoidUndulation,omitempty"`

	// generated from element "verticalDatum" of complex type CodeVerticalDatumType
	VerticalDatum *aixmdatatypes.CodeVerticalDatumType `xml:"verticalDatum,omitempty"`

	// generated from element "verticalAccuracy" of complex type ValDistanceType
	VerticalAccuracy *aixmdatatypes.ValDistanceType `xml:"verticalAccuracy,omitempty"`
}

// PointType Named source named complex type "PointType"
type PointType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "pos" of complex type DirectPositionType
	Pos *gml.DirectPositionType `xml:"pos,omitempty"`

	// generated from group "aixm:PointPropertyGroup" with order 2821 depth 2
	Group_PointPropertyGroup
}

// Group_PointPropertyGroup UnNamed source named group "PointPropertyGroup"
type Group_PointPropertyGroup struct {

	// insertion point for fields

	// generated from element "horizontalAccuracy" of complex type ValDistanceType
	HorizontalAccuracy *aixmdatatypes.ValDistanceType `xml:"horizontalAccuracy,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// PointPropertyType Named source named complex type "PointPropertyType"
type PointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Point" of complex type PointType
	Point *PointType `xml:"Point,omitempty"`

	// generated from element "ElevatedPoint" of complex type ElevatedPointType
	ElevatedPoint *ElevatedPointType `xml:"ElevatedPoint,omitempty"`
}

// SurfaceType Named source named complex type "SurfaceType"
type SurfaceType struct {
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

	// generated from attribute group "gml:SRSReferenceGroup"
	gml.AttributeGroup_SRSReferenceGroup

	// generated from element "patches" of complex type SurfacePatchArrayPropertyType
	Patches *gml.SurfacePatchArrayPropertyType `xml:"patches,omitempty"`

	// generated from group "aixm:SurfacePropertyGroup" with order 2832 depth 2
	Group_SurfacePropertyGroup
}

// Group_SurfacePropertyGroup UnNamed source named group "SurfacePropertyGroup"
type Group_SurfacePropertyGroup struct {

	// insertion point for fields

	// generated from element "horizontalAccuracy" of complex type ValDistanceType
	HorizontalAccuracy *aixmdatatypes.ValDistanceType `xml:"horizontalAccuracy,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SurfacePropertyType Named source named complex type "SurfacePropertyType"
type SurfacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Surface" of complex type SurfaceType
	Surface *SurfaceType `xml:"Surface,omitempty"`

	// generated from element "ElevatedSurface" of complex type ElevatedSurfaceType
	ElevatedSurface *ElevatedSurfaceType `xml:"ElevatedSurface,omitempty"`
}

// DMEPropertyType Named source named complex type "DMEPropertyType"
type DMEPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// NavaidPropertyType Named source named complex type "NavaidPropertyType"
type NavaidPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AngleIndicationPropertyType Named source named complex type "AngleIndicationPropertyType"
type AngleIndicationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AngleUseType Named source named complex type "AngleUseType"
type AngleUseType struct {
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

	// generated from group "aixm:AngleUsePropertyGroup" with order 3628 depth 2
	Group_AngleUsePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAngleUseExtension" of complex type AbstractExtensionType
	AbstractAngleUseExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAngleUseExtension,omitempty"`
}

// Group_AngleUsePropertyGroup UnNamed source named group "AngleUsePropertyGroup"
type Group_AngleUsePropertyGroup struct {

	// insertion point for fields

	// generated from element "alongCourseGuidance" of complex type CodeYesNoType
	AlongCourseGuidance *aixmdatatypes.CodeYesNoType `xml:"alongCourseGuidance,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theAngleIndication" of complex type AngleIndicationPropertyType
	TheAngleIndication *AngleIndicationPropertyType `xml:"theAngleIndication,omitempty"`
}

// AngleUsePropertyType Named source named complex type "AngleUsePropertyType"
type AngleUsePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AngleUse" of complex type AngleUseType
	AngleUse *AngleUseType `xml:"AngleUse,omitempty"`
}

// DistanceIndicationPropertyType Named source named complex type "DistanceIndicationPropertyType"
type DistanceIndicationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// EnRouteSegmentPointType Named source named complex type "EnRouteSegmentPointType"
type EnRouteSegmentPointType struct {
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

	// generated from group "aixm:SegmentPointPropertyGroup" with order 3681 depth 2
	Group_SegmentPointPropertyGroup

	// generated from group "aixm:EnRouteSegmentPointPropertyGroup" with order 3682 depth 2
	Group_EnRouteSegmentPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_9 `xml:"extension,omitempty"`
}

// A_extension_9 Named source within outer element "extension"
type A_extension_9 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSegmentPointExtension" of complex type AbstractExtensionType
	AbstractSegmentPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentPointExtension,omitempty"`

	// generated from element "AbstractEnRouteSegmentPointExtension" of complex type AbstractExtensionType
	AbstractEnRouteSegmentPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractEnRouteSegmentPointExtension,omitempty"`
}

// Group_EnRouteSegmentPointPropertyGroup UnNamed source named group "EnRouteSegmentPointPropertyGroup"
type Group_EnRouteSegmentPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "roleFreeFlight" of complex type CodeFreeFlightType
	RoleFreeFlight *aixmdatatypes.CodeFreeFlightType `xml:"roleFreeFlight,omitempty"`

	// generated from element "roleRVSM" of complex type CodeRVSMPointRoleType
	RoleRVSM *aixmdatatypes.CodeRVSMPointRoleType `xml:"roleRVSM,omitempty"`

	// generated from element "turnRadius" of complex type ValDistanceType
	TurnRadius *aixmdatatypes.ValDistanceType `xml:"turnRadius,omitempty"`

	// generated from element "roleMilitaryTraining" of complex type CodeMilitaryRoutePointType
	RoleMilitaryTraining *aixmdatatypes.CodeMilitaryRoutePointType `xml:"roleMilitaryTraining,omitempty"`
}

// EnRouteSegmentPointPropertyType Named source named complex type "EnRouteSegmentPointPropertyType"
type EnRouteSegmentPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "EnRouteSegmentPoint" of complex type EnRouteSegmentPointType
	EnRouteSegmentPoint *EnRouteSegmentPointType `xml:"EnRouteSegmentPoint,omitempty"`
}

// PointReferenceType Named source named complex type "PointReferenceType"
type PointReferenceType struct {
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

	// generated from group "aixm:PointReferencePropertyGroup" with order 3702 depth 2
	Group_PointReferencePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPointReferenceExtension" of complex type AbstractExtensionType
	AbstractPointReferenceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPointReferenceExtension,omitempty"`
}

// Group_PointReferencePropertyGroup UnNamed source named group "PointReferencePropertyGroup"
type Group_PointReferencePropertyGroup struct {

	// insertion point for fields

	// generated from element "role" of complex type CodeReferenceRoleType
	Role *aixmdatatypes.CodeReferenceRoleType `xml:"role,omitempty"`

	// generated from element "priorFixTolerance" of complex type ValDistanceSignedType
	PriorFixTolerance *aixmdatatypes.ValDistanceSignedType `xml:"priorFixTolerance,omitempty"`

	// generated from element "postFixTolerance" of complex type ValDistanceSignedType
	PostFixTolerance *aixmdatatypes.ValDistanceSignedType `xml:"postFixTolerance,omitempty"`

	// generated from element "point" of complex type DesignatedPointPropertyType
	Point *DesignatedPointPropertyType `xml:"point,omitempty"`

	// generated from element "facilityAngle" of complex type AngleUsePropertyType
	FacilityAngle []*AngleUsePropertyType `xml:"facilityAngle,omitempty"`

	// generated from element "facilityDistance" of complex type DistanceIndicationPropertyType
	FacilityDistance []*DistanceIndicationPropertyType `xml:"facilityDistance,omitempty"`

	// generated from element "fixToleranceArea" of complex type SurfacePropertyType
	FixToleranceArea *SurfacePropertyType `xml:"fixToleranceArea,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// PointReferencePropertyType Named source named complex type "PointReferencePropertyType"
type PointReferencePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "PointReference" of complex type PointReferenceType
	PointReference *PointReferenceType `xml:"PointReference,omitempty"`
}

// AbstractSegmentPointType Named source named complex type "AbstractSegmentPointType"
type AbstractSegmentPointType struct {
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

// Group_SegmentPointPropertyGroup UnNamed source named group "SegmentPointPropertyGroup"
type Group_SegmentPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "reportingATC" of complex type CodeATCReportingType
	ReportingATC *aixmdatatypes.CodeATCReportingType `xml:"reportingATC,omitempty"`

	// generated from element "flyOver" of complex type CodeYesNoType
	FlyOver *aixmdatatypes.CodeYesNoType `xml:"flyOver,omitempty"`

	// generated from element "waypoint" of complex type CodeYesNoType
	Waypoint *aixmdatatypes.CodeYesNoType `xml:"waypoint,omitempty"`

	// generated from element "radarGuidance" of complex type CodeYesNoType
	RadarGuidance *aixmdatatypes.CodeYesNoType `xml:"radarGuidance,omitempty"`

	// generated from element "facilityMakeup" of complex type PointReferencePropertyType
	FacilityMakeup []*PointReferencePropertyType `xml:"facilityMakeup,omitempty"`

	// generated from element "pointChoice_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	PointChoice_fixDesignatedPoint *DesignatedPointPropertyType `xml:"pointChoice_fixDesignatedPoint,omitempty"`

	// generated from element "pointChoice_navaidSystem" of complex type NavaidPropertyType
	PointChoice_navaidSystem *NavaidPropertyType `xml:"pointChoice_navaidSystem,omitempty"`

	// generated from element "pointChoice_position" of complex type PointPropertyType
	PointChoice_position *PointPropertyType `xml:"pointChoice_position,omitempty"`

	// generated from element "pointChoice_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	PointChoice_runwayPoint *RunwayCentrelinePointPropertyType `xml:"pointChoice_runwayPoint,omitempty"`

	// generated from element "pointChoice_aimingPoint" of complex type TouchDownLiftOffPropertyType
	PointChoice_aimingPoint *TouchDownLiftOffPropertyType `xml:"pointChoice_aimingPoint,omitempty"`

	// generated from element "pointChoice_airportReferencePoint" of complex type AirportHeliportPropertyType
	PointChoice_airportReferencePoint *AirportHeliportPropertyType `xml:"pointChoice_airportReferencePoint,omitempty"`

	// generated from element "extendedServiceVolume" of complex type RadioFrequencyAreaPropertyType
	ExtendedServiceVolume *RadioFrequencyAreaPropertyType `xml:"extendedServiceVolume,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// DesignatedPointPropertyType Named source named complex type "DesignatedPointPropertyType"
type DesignatedPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// LinguisticNoteType Named source named complex type "LinguisticNoteType"
type LinguisticNoteType struct {
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

	// generated from group "aixm:LinguisticNotePropertyGroup" with order 3865 depth 2
	Group_LinguisticNotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractLinguisticNoteExtension" of complex type AbstractExtensionType
	AbstractLinguisticNoteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLinguisticNoteExtension,omitempty"`
}

// Group_LinguisticNotePropertyGroup UnNamed source named group "LinguisticNotePropertyGroup"
type Group_LinguisticNotePropertyGroup struct {

	// insertion point for fields

	// generated from element "note" of complex type TextNoteType
	Note *aixmdatatypes.TextNoteType `xml:"note,omitempty"`
}

// LinguisticNotePropertyType Named source named complex type "LinguisticNotePropertyType"
type LinguisticNotePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "LinguisticNote" of complex type LinguisticNoteType
	LinguisticNote *LinguisticNoteType `xml:"LinguisticNote,omitempty"`
}

// NoteType Named source named complex type "NoteType"
type NoteType struct {
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

	// generated from group "aixm:NotePropertyGroup" with order 3881 depth 2
	Group_NotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNoteExtension" of complex type AbstractExtensionType
	AbstractNoteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNoteExtension,omitempty"`
}

// Group_NotePropertyGroup UnNamed source named group "NotePropertyGroup"
type Group_NotePropertyGroup struct {

	// insertion point for fields

	// generated from element "propertyName" of complex type TextPropertyNameType
	PropertyName *aixmdatatypes.TextPropertyNameType `xml:"propertyName,omitempty"`

	// generated from element "purpose" of complex type CodeNotePurposeType
	Purpose *aixmdatatypes.CodeNotePurposeType `xml:"purpose,omitempty"`

	// generated from element "translatedNote" of complex type LinguisticNotePropertyType
	TranslatedNote []*LinguisticNotePropertyType `xml:"translatedNote,omitempty"`
}

// NotePropertyType Named source named complex type "NotePropertyType"
type NotePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Note" of complex type NoteType
	Note *NoteType `xml:"Note,omitempty"`
}

// VerticalStructurePropertyType Named source named complex type "VerticalStructurePropertyType"
type VerticalStructurePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// OrganisationAuthorityPropertyType Named source named complex type "OrganisationAuthorityPropertyType"
type OrganisationAuthorityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// StandardInstrumentArrivalPropertyType Named source named complex type "StandardInstrumentArrivalPropertyType"
type StandardInstrumentArrivalPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// StandardInstrumentDeparturePropertyType Named source named complex type "StandardInstrumentDeparturePropertyType"
type StandardInstrumentDeparturePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// ChangeOverPointType Named source named complex type "ChangeOverPointType"
type ChangeOverPointType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type ChangeOverPointTimeSlicePropertyType
	TimeSlice []*ChangeOverPointTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ChangeOverPointTimeSlicePropertyType Named source named complex type "ChangeOverPointTimeSlicePropertyType"
type ChangeOverPointTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ChangeOverPointTimeSlice" of complex type ChangeOverPointTimeSliceType
	ChangeOverPointTimeSlice *ChangeOverPointTimeSliceType `xml:"ChangeOverPointTimeSlice,omitempty"`
}

// ChangeOverPointTimeSliceType Named source named complex type "ChangeOverPointTimeSliceType"
type ChangeOverPointTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:ChangeOverPointPropertyGroup" with order 5176 depth 2
	Group_ChangeOverPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractChangeOverPointExtension" of complex type AbstractExtensionType
	AbstractChangeOverPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractChangeOverPointExtension,omitempty"`
}

// Group_ChangeOverPointPropertyGroup UnNamed source named group "ChangeOverPointPropertyGroup"
type Group_ChangeOverPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "location_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	Location_fixDesignatedPoint *DesignatedPointPropertyType `xml:"location_fixDesignatedPoint,omitempty"`

	// generated from element "location_navaidSystem" of complex type NavaidPropertyType
	Location_navaidSystem *NavaidPropertyType `xml:"location_navaidSystem,omitempty"`

	// generated from element "location_position" of complex type PointPropertyType
	Location_position *PointPropertyType `xml:"location_position,omitempty"`

	// generated from element "location_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	Location_runwayPoint *RunwayCentrelinePointPropertyType `xml:"location_runwayPoint,omitempty"`

	// generated from element "location_aimingPoint" of complex type TouchDownLiftOffPropertyType
	Location_aimingPoint *TouchDownLiftOffPropertyType `xml:"location_aimingPoint,omitempty"`

	// generated from element "location_airportReferencePoint" of complex type AirportHeliportPropertyType
	Location_airportReferencePoint *AirportHeliportPropertyType `xml:"location_airportReferencePoint,omitempty"`

	// generated from element "applicableRoutePortion" of complex type RoutePortionPropertyType
	ApplicableRoutePortion *RoutePortionPropertyType `xml:"applicableRoutePortion,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RoutePropertyType Named source named complex type "RoutePropertyType"
type RoutePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RouteType Named source named complex type "RouteType"
type RouteType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type RouteTimeSlicePropertyType
	TimeSlice []*RouteTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RouteTimeSlicePropertyType Named source named complex type "RouteTimeSlicePropertyType"
type RouteTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RouteTimeSlice" of complex type RouteTimeSliceType
	RouteTimeSlice *RouteTimeSliceType `xml:"RouteTimeSlice,omitempty"`
}

// RouteTimeSliceType Named source named complex type "RouteTimeSliceType"
type RouteTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from element "designatorPrefix" of complex type CodeRouteDesignatorPrefixType
	DesignatorPrefix *aixmdatatypes.CodeRouteDesignatorPrefixType `xml:"designatorPrefix,omitempty"`

	// generated from element "designatorSecondLetter" of complex type CodeRouteDesignatorLetterType
	DesignatorSecondLetter *aixmdatatypes.CodeRouteDesignatorLetterType `xml:"designatorSecondLetter,omitempty"`

	// generated from element "designatorNumber" of complex type NoNumberType
	DesignatorNumber *aixmdatatypes.NoNumberType `xml:"designatorNumber,omitempty"`

	// generated from element "multipleIdentifier" of complex type CodeUpperAlphaType
	MultipleIdentifier *aixmdatatypes.CodeUpperAlphaType `xml:"multipleIdentifier,omitempty"`

	// generated from element "locationDesignator" of complex type TextDesignatorType
	LocationDesignator *aixmdatatypes.TextDesignatorType `xml:"locationDesignator,omitempty"`

	// generated from element "type" of complex type CodeRouteType
	Type_XSD *aixmdatatypes.CodeRouteType `xml:"type,omitempty"`

	// generated from element "flightRule" of complex type CodeFlightRuleType
	FlightRule *aixmdatatypes.CodeFlightRuleType `xml:"flightRule,omitempty"`

	// generated from element "internationalUse" of complex type CodeRouteOriginType
	InternationalUse *aixmdatatypes.CodeRouteOriginType `xml:"internationalUse,omitempty"`

	// generated from element "militaryUse" of complex type CodeMilitaryStatusType
	MilitaryUse *aixmdatatypes.CodeMilitaryStatusType `xml:"militaryUse,omitempty"`

	// generated from element "militaryTrainingType" of complex type CodeMilitaryTrainingType
	MilitaryTrainingType *aixmdatatypes.CodeMilitaryTrainingType `xml:"militaryTrainingType,omitempty"`

	// generated from element "userOrganisation" of complex type OrganisationAuthorityPropertyType
	UserOrganisation *OrganisationAuthorityPropertyType `xml:"userOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRouteExtension" of complex type AbstractExtensionType
	AbstractRouteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRouteExtension,omitempty"`
}

// Group_RoutePropertyGroup UnNamed source named group "RoutePropertyGroup"
type Group_RoutePropertyGroup struct {

	// insertion point for fields

	// generated from element "designatorPrefix" of complex type CodeRouteDesignatorPrefixType
	DesignatorPrefix *aixmdatatypes.CodeRouteDesignatorPrefixType `xml:"designatorPrefix,omitempty"`

	// generated from element "designatorSecondLetter" of complex type CodeRouteDesignatorLetterType
	DesignatorSecondLetter *aixmdatatypes.CodeRouteDesignatorLetterType `xml:"designatorSecondLetter,omitempty"`

	// generated from element "designatorNumber" of complex type NoNumberType
	DesignatorNumber *aixmdatatypes.NoNumberType `xml:"designatorNumber,omitempty"`

	// generated from element "multipleIdentifier" of complex type CodeUpperAlphaType
	MultipleIdentifier *aixmdatatypes.CodeUpperAlphaType `xml:"multipleIdentifier,omitempty"`

	// generated from element "locationDesignator" of complex type TextDesignatorType
	LocationDesignator *aixmdatatypes.TextDesignatorType `xml:"locationDesignator,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeRouteType
	Type_XSD *aixmdatatypes.CodeRouteType `xml:"type,omitempty"`

	// generated from element "flightRule" of complex type CodeFlightRuleType
	FlightRule *aixmdatatypes.CodeFlightRuleType `xml:"flightRule,omitempty"`

	// generated from element "internationalUse" of complex type CodeRouteOriginType
	InternationalUse *aixmdatatypes.CodeRouteOriginType `xml:"internationalUse,omitempty"`

	// generated from element "militaryUse" of complex type CodeMilitaryStatusType
	MilitaryUse *aixmdatatypes.CodeMilitaryStatusType `xml:"militaryUse,omitempty"`

	// generated from element "militaryTrainingType" of complex type CodeMilitaryTrainingType
	MilitaryTrainingType *aixmdatatypes.CodeMilitaryTrainingType `xml:"militaryTrainingType,omitempty"`

	// generated from element "userOrganisation" of complex type OrganisationAuthorityPropertyType
	UserOrganisation *OrganisationAuthorityPropertyType `xml:"userOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RouteAvailabilityType Named source named complex type "RouteAvailabilityType"
type RouteAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 5234 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:RouteAvailabilityPropertyGroup" with order 5235 depth 2
	Group_RouteAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_15 `xml:"extension,omitempty"`
}

// A_extension_15 Named source within outer element "extension"
type A_extension_15 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractRouteAvailabilityExtension" of complex type AbstractExtensionType
	AbstractRouteAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRouteAvailabilityExtension,omitempty"`
}

// Group_RouteAvailabilityPropertyGroup UnNamed source named group "RouteAvailabilityPropertyGroup"
type Group_RouteAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "direction" of complex type CodeDirectionType
	Direction *aixmdatatypes.CodeDirectionType `xml:"direction,omitempty"`

	// generated from element "cardinalDirection" of complex type CodeCardinalDirectionType
	CardinalDirection *aixmdatatypes.CodeCardinalDirectionType `xml:"cardinalDirection,omitempty"`

	// generated from element "status" of complex type CodeRouteAvailabilityType
	Status *aixmdatatypes.CodeRouteAvailabilityType `xml:"status,omitempty"`

	// generated from element "levels" of complex type AirspaceLayerPropertyType
	Levels []*AirspaceLayerPropertyType `xml:"levels,omitempty"`
}

// RouteAvailabilityPropertyType Named source named complex type "RouteAvailabilityPropertyType"
type RouteAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RouteAvailability" of complex type RouteAvailabilityType
	RouteAvailability *RouteAvailabilityType `xml:"RouteAvailability,omitempty"`
}

// RouteDMEType Named source named complex type "RouteDMEType"
type RouteDMEType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type RouteDMETimeSlicePropertyType
	TimeSlice []*RouteDMETimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RouteDMETimeSlicePropertyType Named source named complex type "RouteDMETimeSlicePropertyType"
type RouteDMETimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RouteDMETimeSlice" of complex type RouteDMETimeSliceType
	RouteDMETimeSlice *RouteDMETimeSliceType `xml:"RouteDMETimeSlice,omitempty"`
}

// RouteDMETimeSliceType Named source named complex type "RouteDMETimeSliceType"
type RouteDMETimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:RouteDMEPropertyGroup" with order 5266 depth 2
	Group_RouteDMEPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_16 `xml:"extension,omitempty"`
}

// A_extension_16 Named source within outer element "extension"
type A_extension_16 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRouteDMEExtension" of complex type AbstractExtensionType
	AbstractRouteDMEExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRouteDMEExtension,omitempty"`
}

// Group_RouteDMEPropertyGroup UnNamed source named group "RouteDMEPropertyGroup"
type Group_RouteDMEPropertyGroup struct {

	// insertion point for fields

	// generated from element "criticalDME" of complex type CodeYesNoType
	CriticalDME *aixmdatatypes.CodeYesNoType `xml:"criticalDME,omitempty"`

	// generated from element "satisfactory" of complex type CodeYesNoType
	Satisfactory *aixmdatatypes.CodeYesNoType `xml:"satisfactory,omitempty"`

	// generated from element "referencedDME" of complex type DMEPropertyType
	ReferencedDME *DMEPropertyType `xml:"referencedDME,omitempty"`

	// generated from element "applicableRoutePortion" of complex type RoutePortionPropertyType
	ApplicableRoutePortion *RoutePortionPropertyType `xml:"applicableRoutePortion,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RoutePortionType Named source named complex type "RoutePortionType"
type RoutePortionType struct {
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

	// generated from group "aixm:RoutePortionPropertyGroup" with order 5283 depth 2
	Group_RoutePortionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_17 `xml:"extension,omitempty"`
}

// A_extension_17 Named source within outer element "extension"
type A_extension_17 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRoutePortionExtension" of complex type AbstractExtensionType
	AbstractRoutePortionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRoutePortionExtension,omitempty"`
}

// Group_RoutePortionPropertyGroup UnNamed source named group "RoutePortionPropertyGroup"
type Group_RoutePortionPropertyGroup struct {

	// insertion point for fields

	// generated from element "start_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	Start_fixDesignatedPoint *DesignatedPointPropertyType `xml:"start_fixDesignatedPoint,omitempty"`

	// generated from element "start_navaidSystem" of complex type NavaidPropertyType
	Start_navaidSystem *NavaidPropertyType `xml:"start_navaidSystem,omitempty"`

	// generated from element "start_position" of complex type PointPropertyType
	Start_position *PointPropertyType `xml:"start_position,omitempty"`

	// generated from element "start_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	Start_runwayPoint *RunwayCentrelinePointPropertyType `xml:"start_runwayPoint,omitempty"`

	// generated from element "start_aimingPoint" of complex type TouchDownLiftOffPropertyType
	Start_aimingPoint *TouchDownLiftOffPropertyType `xml:"start_aimingPoint,omitempty"`

	// generated from element "start_airportReferencePoint" of complex type AirportHeliportPropertyType
	Start_airportReferencePoint *AirportHeliportPropertyType `xml:"start_airportReferencePoint,omitempty"`

	// generated from element "intermediatePoint_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	IntermediatePoint_fixDesignatedPoint *DesignatedPointPropertyType `xml:"intermediatePoint_fixDesignatedPoint,omitempty"`

	// generated from element "intermediatePoint_navaidSystem" of complex type NavaidPropertyType
	IntermediatePoint_navaidSystem *NavaidPropertyType `xml:"intermediatePoint_navaidSystem,omitempty"`

	// generated from element "intermediatePoint_position" of complex type PointPropertyType
	IntermediatePoint_position *PointPropertyType `xml:"intermediatePoint_position,omitempty"`

	// generated from element "intermediatePoint_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	IntermediatePoint_runwayPoint *RunwayCentrelinePointPropertyType `xml:"intermediatePoint_runwayPoint,omitempty"`

	// generated from element "intermediatePoint_aimingPoint" of complex type TouchDownLiftOffPropertyType
	IntermediatePoint_aimingPoint *TouchDownLiftOffPropertyType `xml:"intermediatePoint_aimingPoint,omitempty"`

	// generated from element "intermediatePoint_airportReferencePoint" of complex type AirportHeliportPropertyType
	IntermediatePoint_airportReferencePoint *AirportHeliportPropertyType `xml:"intermediatePoint_airportReferencePoint,omitempty"`

	// generated from element "referencedRoute" of complex type RoutePropertyType
	ReferencedRoute *RoutePropertyType `xml:"referencedRoute,omitempty"`

	// generated from element "end_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	End_fixDesignatedPoint *DesignatedPointPropertyType `xml:"end_fixDesignatedPoint,omitempty"`

	// generated from element "end_navaidSystem" of complex type NavaidPropertyType
	End_navaidSystem *NavaidPropertyType `xml:"end_navaidSystem,omitempty"`

	// generated from element "end_position" of complex type PointPropertyType
	End_position *PointPropertyType `xml:"end_position,omitempty"`

	// generated from element "end_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	End_runwayPoint *RunwayCentrelinePointPropertyType `xml:"end_runwayPoint,omitempty"`

	// generated from element "end_aimingPoint" of complex type TouchDownLiftOffPropertyType
	End_aimingPoint *TouchDownLiftOffPropertyType `xml:"end_aimingPoint,omitempty"`

	// generated from element "end_airportReferencePoint" of complex type AirportHeliportPropertyType
	End_airportReferencePoint *AirportHeliportPropertyType `xml:"end_airportReferencePoint,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RoutePortionPropertyType Named source named complex type "RoutePortionPropertyType"
type RoutePortionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RoutePortion" of complex type RoutePortionType
	RoutePortion *RoutePortionType `xml:"RoutePortion,omitempty"`
}

// RouteSegmentType Named source named complex type "RouteSegmentType"
type RouteSegmentType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type RouteSegmentTimeSlicePropertyType
	TimeSlice []*RouteSegmentTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RouteSegmentTimeSlicePropertyType Named source named complex type "RouteSegmentTimeSlicePropertyType"
type RouteSegmentTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RouteSegmentTimeSlice" of complex type RouteSegmentTimeSliceType
	RouteSegmentTimeSlice *RouteSegmentTimeSliceType `xml:"RouteSegmentTimeSlice,omitempty"`
}

// RouteSegmentTimeSliceType Named source named complex type "RouteSegmentTimeSliceType"
type RouteSegmentTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:RouteSegmentPropertyGroup" with order 5332 depth 2
	Group_RouteSegmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_18 `xml:"extension,omitempty"`
}

// A_extension_18 Named source within outer element "extension"
type A_extension_18 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRouteSegmentExtension" of complex type AbstractExtensionType
	AbstractRouteSegmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRouteSegmentExtension,omitempty"`
}

// Group_RouteSegmentPropertyGroup UnNamed source named group "RouteSegmentPropertyGroup"
type Group_RouteSegmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "level" of complex type CodeLevelType
	Level *aixmdatatypes.CodeLevelType `xml:"level,omitempty"`

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "minimumObstacleClearanceAltitude" of complex type ValDistanceVerticalType
	MinimumObstacleClearanceAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumObstacleClearanceAltitude,omitempty"`

	// generated from element "pathType" of complex type CodeRouteSegmentPathType
	PathType *aixmdatatypes.CodeRouteSegmentPathType `xml:"pathType,omitempty"`

	// generated from element "trueTrack" of complex type ValBearingType
	TrueTrack *aixmdatatypes.ValBearingType `xml:"trueTrack,omitempty"`

	// generated from element "magneticTrack" of complex type ValBearingType
	MagneticTrack *aixmdatatypes.ValBearingType `xml:"magneticTrack,omitempty"`

	// generated from element "reverseTrueTrack" of complex type ValBearingType
	ReverseTrueTrack *aixmdatatypes.ValBearingType `xml:"reverseTrueTrack,omitempty"`

	// generated from element "reverseMagneticTrack" of complex type ValBearingType
	ReverseMagneticTrack *aixmdatatypes.ValBearingType `xml:"reverseMagneticTrack,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "widthLeft" of complex type ValDistanceType
	WidthLeft *aixmdatatypes.ValDistanceType `xml:"widthLeft,omitempty"`

	// generated from element "widthRight" of complex type ValDistanceType
	WidthRight *aixmdatatypes.ValDistanceType `xml:"widthRight,omitempty"`

	// generated from element "turnDirection" of complex type CodeDirectionTurnType
	TurnDirection *aixmdatatypes.CodeDirectionTurnType `xml:"turnDirection,omitempty"`

	// generated from element "signalGap" of complex type CodeYesNoType
	SignalGap *aixmdatatypes.CodeYesNoType `xml:"signalGap,omitempty"`

	// generated from element "minimumEnrouteAltitude" of complex type ValDistanceVerticalType
	MinimumEnrouteAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumEnrouteAltitude,omitempty"`

	// generated from element "minimumCrossingAtEnd" of complex type ValDistanceVerticalType
	MinimumCrossingAtEnd *aixmdatatypes.ValDistanceVerticalType `xml:"minimumCrossingAtEnd,omitempty"`

	// generated from element "minimumCrossingAtEndReference" of complex type CodeVerticalReferenceType
	MinimumCrossingAtEndReference *aixmdatatypes.CodeVerticalReferenceType `xml:"minimumCrossingAtEndReference,omitempty"`

	// generated from element "maximumCrossingAtEnd" of complex type ValDistanceVerticalType
	MaximumCrossingAtEnd *aixmdatatypes.ValDistanceVerticalType `xml:"maximumCrossingAtEnd,omitempty"`

	// generated from element "maximumCrossingAtEndReference" of complex type CodeVerticalReferenceType
	MaximumCrossingAtEndReference *aixmdatatypes.CodeVerticalReferenceType `xml:"maximumCrossingAtEndReference,omitempty"`

	// generated from element "navigationType" of complex type CodeRouteNavigationType
	NavigationType *aixmdatatypes.CodeRouteNavigationType `xml:"navigationType,omitempty"`

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`

	// generated from element "designatorSuffix" of complex type CodeRouteDesignatorSuffixType
	DesignatorSuffix *aixmdatatypes.CodeRouteDesignatorSuffixType `xml:"designatorSuffix,omitempty"`

	// generated from element "start" of complex type EnRouteSegmentPointPropertyType
	Start *EnRouteSegmentPointPropertyType `xml:"start,omitempty"`

	// generated from element "routeFormed" of complex type RoutePropertyType
	RouteFormed *RoutePropertyType `xml:"routeFormed,omitempty"`

	// generated from element "evaluationArea" of complex type ObstacleAssessmentAreaPropertyType
	EvaluationArea *ObstacleAssessmentAreaPropertyType `xml:"evaluationArea,omitempty"`

	// generated from element "curveExtent" of complex type CurvePropertyType
	CurveExtent *CurvePropertyType `xml:"curveExtent,omitempty"`

	// generated from element "end" of complex type EnRouteSegmentPointPropertyType
	End *EnRouteSegmentPointPropertyType `xml:"end,omitempty"`

	// generated from element "availability" of complex type RouteAvailabilityPropertyType
	Availability []*RouteAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirspaceBorderCrossingPropertyType Named source named complex type "AirspaceBorderCrossingPropertyType"
type AirspaceBorderCrossingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AirspaceBorderCrossingType Named source named complex type "AirspaceBorderCrossingType"
type AirspaceBorderCrossingType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type AirspaceBorderCrossingTimeSlicePropertyType
	TimeSlice []*AirspaceBorderCrossingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirspaceBorderCrossingTimeSlicePropertyType Named source named complex type "AirspaceBorderCrossingTimeSlicePropertyType"
type AirspaceBorderCrossingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirspaceBorderCrossingTimeSlice" of complex type AirspaceBorderCrossingTimeSliceType
	AirspaceBorderCrossingTimeSlice *AirspaceBorderCrossingTimeSliceType `xml:"AirspaceBorderCrossingTimeSlice,omitempty"`
}

// AirspaceBorderCrossingTimeSliceType Named source named complex type "AirspaceBorderCrossingTimeSliceType"
type AirspaceBorderCrossingTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:AirspaceBorderCrossingPropertyGroup" with order 5386 depth 2
	Group_AirspaceBorderCrossingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_19 `xml:"extension,omitempty"`
}

// A_extension_19 Named source within outer element "extension"
type A_extension_19 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirspaceBorderCrossingExtension" of complex type AbstractExtensionType
	AbstractAirspaceBorderCrossingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceBorderCrossingExtension,omitempty"`
}

// Group_AirspaceBorderCrossingPropertyGroup UnNamed source named group "AirspaceBorderCrossingPropertyGroup"
type Group_AirspaceBorderCrossingPropertyGroup struct {

	// insertion point for fields

	// generated from element "exitedAirspace" of complex type AirspacePropertyType
	ExitedAirspace *AirspacePropertyType `xml:"exitedAirspace,omitempty"`

	// generated from element "enteredAirspace" of complex type AirspacePropertyType
	EnteredAirspace *AirspacePropertyType `xml:"enteredAirspace,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AbstractDirectFlightType Named source named complex type "AbstractDirectFlightType"
type AbstractDirectFlightType struct {
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

// Group_DirectFlightPropertyGroup UnNamed source named group "DirectFlightPropertyGroup"
type Group_DirectFlightPropertyGroup struct {

	// insertion point for fields

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// DirectFlightPropertyType Named source named complex type "DirectFlightPropertyType"
type DirectFlightPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "DirectFlightClass" of complex type DirectFlightClassType
	DirectFlightClass *DirectFlightClassType `xml:"DirectFlightClass,omitempty"`

	// generated from element "DirectFlightSegment" of complex type DirectFlightSegmentType
	DirectFlightSegment *DirectFlightSegmentType `xml:"DirectFlightSegment,omitempty"`
}

// DirectFlightClassType Named source named complex type "DirectFlightClassType"
type DirectFlightClassType struct {
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

	// generated from group "aixm:DirectFlightPropertyGroup" with order 5411 depth 2
	Group_DirectFlightPropertyGroup

	// generated from group "aixm:DirectFlightClassPropertyGroup" with order 5412 depth 2
	Group_DirectFlightClassPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_20 `xml:"extension,omitempty"`
}

// A_extension_20 Named source within outer element "extension"
type A_extension_20 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDirectFlightExtension" of complex type AbstractExtensionType
	AbstractDirectFlightExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDirectFlightExtension,omitempty"`

	// generated from element "AbstractDirectFlightClassExtension" of complex type AbstractExtensionType
	AbstractDirectFlightClassExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDirectFlightClassExtension,omitempty"`
}

// Group_DirectFlightClassPropertyGroup UnNamed source named group "DirectFlightClassPropertyGroup"
type Group_DirectFlightClassPropertyGroup struct {

	// insertion point for fields

	// generated from element "exceedLength" of complex type ValDistanceType
	ExceedLength *aixmdatatypes.ValDistanceType `xml:"exceedLength,omitempty"`
}

// DirectFlightSegmentType Named source named complex type "DirectFlightSegmentType"
type DirectFlightSegmentType struct {
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

	// generated from group "aixm:DirectFlightPropertyGroup" with order 5429 depth 2
	Group_DirectFlightPropertyGroup

	// generated from group "aixm:DirectFlightSegmentPropertyGroup" with order 5430 depth 2
	Group_DirectFlightSegmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_21 `xml:"extension,omitempty"`
}

// A_extension_21 Named source within outer element "extension"
type A_extension_21 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDirectFlightExtension" of complex type AbstractExtensionType
	AbstractDirectFlightExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDirectFlightExtension,omitempty"`

	// generated from element "AbstractDirectFlightSegmentExtension" of complex type AbstractExtensionType
	AbstractDirectFlightSegmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDirectFlightSegmentExtension,omitempty"`
}

// Group_DirectFlightSegmentPropertyGroup UnNamed source named group "DirectFlightSegmentPropertyGroup"
type Group_DirectFlightSegmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "end_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	End_fixDesignatedPoint *DesignatedPointPropertyType `xml:"end_fixDesignatedPoint,omitempty"`

	// generated from element "end_navaidSystem" of complex type NavaidPropertyType
	End_navaidSystem *NavaidPropertyType `xml:"end_navaidSystem,omitempty"`

	// generated from element "end_position" of complex type PointPropertyType
	End_position *PointPropertyType `xml:"end_position,omitempty"`

	// generated from element "end_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	End_runwayPoint *RunwayCentrelinePointPropertyType `xml:"end_runwayPoint,omitempty"`

	// generated from element "end_aimingPoint" of complex type TouchDownLiftOffPropertyType
	End_aimingPoint *TouchDownLiftOffPropertyType `xml:"end_aimingPoint,omitempty"`

	// generated from element "end_airportReferencePoint" of complex type AirportHeliportPropertyType
	End_airportReferencePoint *AirportHeliportPropertyType `xml:"end_airportReferencePoint,omitempty"`

	// generated from element "start_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	Start_fixDesignatedPoint *DesignatedPointPropertyType `xml:"start_fixDesignatedPoint,omitempty"`

	// generated from element "start_navaidSystem" of complex type NavaidPropertyType
	Start_navaidSystem *NavaidPropertyType `xml:"start_navaidSystem,omitempty"`

	// generated from element "start_position" of complex type PointPropertyType
	Start_position *PointPropertyType `xml:"start_position,omitempty"`

	// generated from element "start_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	Start_runwayPoint *RunwayCentrelinePointPropertyType `xml:"start_runwayPoint,omitempty"`

	// generated from element "start_aimingPoint" of complex type TouchDownLiftOffPropertyType
	Start_aimingPoint *TouchDownLiftOffPropertyType `xml:"start_aimingPoint,omitempty"`

	// generated from element "start_airportReferencePoint" of complex type AirportHeliportPropertyType
	Start_airportReferencePoint *AirportHeliportPropertyType `xml:"start_airportReferencePoint,omitempty"`
}

// DirectFlightSegmentPropertyType Named source named complex type "DirectFlightSegmentPropertyType"
type DirectFlightSegmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "DirectFlightSegment" of complex type DirectFlightSegmentType
	DirectFlightSegment *DirectFlightSegmentType `xml:"DirectFlightSegment,omitempty"`
}

// FlightConditionCircumstanceType Named source named complex type "FlightConditionCircumstanceType"
type FlightConditionCircumstanceType struct {
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

	// generated from group "aixm:FlightConditionCircumstancePropertyGroup" with order 5460 depth 2
	Group_FlightConditionCircumstancePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_22 `xml:"extension,omitempty"`
}

// A_extension_22 Named source within outer element "extension"
type A_extension_22 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightConditionCircumstanceExtension" of complex type AbstractExtensionType
	AbstractFlightConditionCircumstanceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightConditionCircumstanceExtension,omitempty"`
}

// Group_FlightConditionCircumstancePropertyGroup UnNamed source named group "FlightConditionCircumstancePropertyGroup"
type Group_FlightConditionCircumstancePropertyGroup struct {

	// insertion point for fields

	// generated from element "referenceLocation" of complex type CodeYesNoType
	ReferenceLocation *aixmdatatypes.CodeYesNoType `xml:"referenceLocation,omitempty"`

	// generated from element "relationWithLocation" of complex type CodeLocationQualifierType
	RelationWithLocation *aixmdatatypes.CodeLocationQualifierType `xml:"relationWithLocation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightConditionCircumstancePropertyType Named source named complex type "FlightConditionCircumstancePropertyType"
type FlightConditionCircumstancePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightConditionCircumstance" of complex type FlightConditionCircumstanceType
	FlightConditionCircumstance *FlightConditionCircumstanceType `xml:"FlightConditionCircumstance,omitempty"`
}

// FlightConditionCombinationType Named source named complex type "FlightConditionCombinationType"
type FlightConditionCombinationType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 5478 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:FlightConditionCombinationPropertyGroup" with order 5479 depth 2
	Group_FlightConditionCombinationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_23 `xml:"extension,omitempty"`
}

// A_extension_23 Named source within outer element "extension"
type A_extension_23 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractFlightConditionCombinationExtension" of complex type AbstractExtensionType
	AbstractFlightConditionCombinationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightConditionCombinationExtension,omitempty"`
}

// Group_FlightConditionCombinationPropertyGroup UnNamed source named group "FlightConditionCombinationPropertyGroup"
type Group_FlightConditionCombinationPropertyGroup struct {

	// insertion point for fields

	// generated from element "logicalOperator" of complex type CodeFlowConditionOperationType
	LogicalOperator *aixmdatatypes.CodeFlowConditionOperationType `xml:"logicalOperator,omitempty"`

	// generated from element "element" of complex type FlightConditionElementPropertyType
	Element []*FlightConditionElementPropertyType `xml:"element,omitempty"`
}

// FlightConditionCombinationPropertyType Named source named complex type "FlightConditionCombinationPropertyType"
type FlightConditionCombinationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightConditionCombination" of complex type FlightConditionCombinationType
	FlightConditionCombination *FlightConditionCombinationType `xml:"FlightConditionCombination,omitempty"`
}

// FlightConditionElementType Named source named complex type "FlightConditionElementType"
type FlightConditionElementType struct {
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

	// generated from group "aixm:FlightConditionElementPropertyGroup" with order 5497 depth 2
	Group_FlightConditionElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_24 `xml:"extension,omitempty"`
}

// A_extension_24 Named source within outer element "extension"
type A_extension_24 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightConditionElementExtension" of complex type AbstractExtensionType
	AbstractFlightConditionElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightConditionElementExtension,omitempty"`
}

// Group_FlightConditionElementPropertyGroup UnNamed source named group "FlightConditionElementPropertyGroup"
type Group_FlightConditionElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "index" of complex type NoSequenceType
	Index *aixmdatatypes.NoSequenceType `xml:"index,omitempty"`

	// generated from element "flightCondition_weather" of complex type MeteorologyPropertyType
	FlightCondition_weather *MeteorologyPropertyType `xml:"flightCondition_weather,omitempty"`

	// generated from element "flightCondition_aircraft" of complex type AircraftCharacteristicPropertyType
	FlightCondition_aircraft *AircraftCharacteristicPropertyType `xml:"flightCondition_aircraft,omitempty"`

	// generated from element "flightCondition_flight" of complex type FlightCharacteristicPropertyType
	FlightCondition_flight *FlightCharacteristicPropertyType `xml:"flightCondition_flight,omitempty"`

	// generated from element "flightCondition_directFlightCondition" of complex type DirectFlightPropertyType
	FlightCondition_directFlightCondition *DirectFlightPropertyType `xml:"flightCondition_directFlightCondition,omitempty"`

	// generated from element "flightCondition_borderCrossingCondition" of complex type AirspaceBorderCrossingPropertyType
	FlightCondition_borderCrossingCondition *AirspaceBorderCrossingPropertyType `xml:"flightCondition_borderCrossingCondition,omitempty"`

	// generated from element "flightCondition_operand" of complex type FlightConditionCombinationPropertyType
	FlightCondition_operand *FlightConditionCombinationPropertyType `xml:"flightCondition_operand,omitempty"`

	// generated from element "flightCondition_routePortionCondition" of complex type RoutePortionPropertyType
	FlightCondition_routePortionCondition *RoutePortionPropertyType `xml:"flightCondition_routePortionCondition,omitempty"`

	// generated from element "flightCondition_standardInstrumentDepartureCondition" of complex type StandardInstrumentDeparturePropertyType
	FlightCondition_standardInstrumentDepartureCondition *StandardInstrumentDeparturePropertyType `xml:"flightCondition_standardInstrumentDepartureCondition,omitempty"`

	// generated from element "flightCondition_standardInstrumentArrivalCondition" of complex type StandardInstrumentArrivalPropertyType
	FlightCondition_standardInstrumentArrivalCondition *StandardInstrumentArrivalPropertyType `xml:"flightCondition_standardInstrumentArrivalCondition,omitempty"`

	// generated from element "flightCondition_organisationCondition" of complex type OrganisationAuthorityPropertyType
	FlightCondition_organisationCondition *OrganisationAuthorityPropertyType `xml:"flightCondition_organisationCondition,omitempty"`

	// generated from element "significantPointCondition_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	SignificantPointCondition_fixDesignatedPoint *DesignatedPointPropertyType `xml:"significantPointCondition_fixDesignatedPoint,omitempty"`

	// generated from element "significantPointCondition_navaidSystem" of complex type NavaidPropertyType
	SignificantPointCondition_navaidSystem *NavaidPropertyType `xml:"significantPointCondition_navaidSystem,omitempty"`

	// generated from element "significantPointCondition_position" of complex type PointPropertyType
	SignificantPointCondition_position *PointPropertyType `xml:"significantPointCondition_position,omitempty"`

	// generated from element "significantPointCondition_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	SignificantPointCondition_runwayPoint *RunwayCentrelinePointPropertyType `xml:"significantPointCondition_runwayPoint,omitempty"`

	// generated from element "significantPointCondition_aimingPoint" of complex type TouchDownLiftOffPropertyType
	SignificantPointCondition_aimingPoint *TouchDownLiftOffPropertyType `xml:"significantPointCondition_aimingPoint,omitempty"`

	// generated from element "significantPointCondition_airportReferencePoint" of complex type AirportHeliportPropertyType
	SignificantPointCondition_airportReferencePoint *AirportHeliportPropertyType `xml:"significantPointCondition_airportReferencePoint,omitempty"`

	// generated from element "flightCondition_airspaceCondition" of complex type AirspacePropertyType
	FlightCondition_airspaceCondition *AirspacePropertyType `xml:"flightCondition_airspaceCondition,omitempty"`

	// generated from element "flightCondition_airportHeliportCondition" of complex type AirportHeliportPropertyType
	FlightCondition_airportHeliportCondition *AirportHeliportPropertyType `xml:"flightCondition_airportHeliportCondition,omitempty"`

	// generated from element "flightCondition_aerialRefuellingCondition" of complex type AerialRefuellingPropertyType
	FlightCondition_aerialRefuellingCondition *AerialRefuellingPropertyType `xml:"flightCondition_aerialRefuellingCondition,omitempty"`

	// generated from element "operationalCondition" of complex type FlightConditionCircumstancePropertyType
	OperationalCondition *FlightConditionCircumstancePropertyType `xml:"operationalCondition,omitempty"`

	// generated from element "flightLevel" of complex type FlightRestrictionLevelPropertyType
	FlightLevel []*FlightRestrictionLevelPropertyType `xml:"flightLevel,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightConditionElementPropertyType Named source named complex type "FlightConditionElementPropertyType"
type FlightConditionElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightConditionElement" of complex type FlightConditionElementType
	FlightConditionElement *FlightConditionElementType `xml:"FlightConditionElement,omitempty"`
}

// FlightRestrictionType Named source named complex type "FlightRestrictionType"
type FlightRestrictionType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type FlightRestrictionTimeSlicePropertyType
	TimeSlice []*FlightRestrictionTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// FlightRestrictionTimeSlicePropertyType Named source named complex type "FlightRestrictionTimeSlicePropertyType"
type FlightRestrictionTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "FlightRestrictionTimeSlice" of complex type FlightRestrictionTimeSliceType
	FlightRestrictionTimeSlice *FlightRestrictionTimeSliceType `xml:"FlightRestrictionTimeSlice,omitempty"`
}

// FlightRestrictionTimeSliceType Named source named complex type "FlightRestrictionTimeSliceType"
type FlightRestrictionTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:FlightRestrictionPropertyGroup" with order 5548 depth 2
	Group_FlightRestrictionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_25 `xml:"extension,omitempty"`
}

// A_extension_25 Named source within outer element "extension"
type A_extension_25 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightRestrictionExtension" of complex type AbstractExtensionType
	AbstractFlightRestrictionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightRestrictionExtension,omitempty"`
}

// Group_FlightRestrictionPropertyGroup UnNamed source named group "FlightRestrictionPropertyGroup"
type Group_FlightRestrictionPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type CodeFlightRestrictionDesignatorType
	Designator *aixmdatatypes.CodeFlightRestrictionDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeFlightRestrictionType
	Type_XSD *aixmdatatypes.CodeFlightRestrictionType `xml:"type,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "flight" of complex type FlightConditionCombinationPropertyType
	Flight *FlightConditionCombinationPropertyType `xml:"flight,omitempty"`

	// generated from element "regulatedRoute" of complex type FlightRestrictionRoutePropertyType
	RegulatedRoute []*FlightRestrictionRoutePropertyType `xml:"regulatedRoute,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightRestrictionLevelType Named source named complex type "FlightRestrictionLevelType"
type FlightRestrictionLevelType struct {
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

	// generated from group "aixm:FlightRestrictionLevelPropertyGroup" with order 5566 depth 2
	Group_FlightRestrictionLevelPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_26 `xml:"extension,omitempty"`
}

// A_extension_26 Named source within outer element "extension"
type A_extension_26 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightRestrictionLevelExtension" of complex type AbstractExtensionType
	AbstractFlightRestrictionLevelExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightRestrictionLevelExtension,omitempty"`
}

// Group_FlightRestrictionLevelPropertyGroup UnNamed source named group "FlightRestrictionLevelPropertyGroup"
type Group_FlightRestrictionLevelPropertyGroup struct {

	// insertion point for fields

	// generated from element "upperLevel" of complex type ValDistanceVerticalType
	UpperLevel *aixmdatatypes.ValDistanceVerticalType `xml:"upperLevel,omitempty"`

	// generated from element "upperLevelReference" of complex type CodeVerticalReferenceType
	UpperLevelReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLevelReference,omitempty"`

	// generated from element "lowerLevel" of complex type ValDistanceVerticalType
	LowerLevel *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLevel,omitempty"`

	// generated from element "lowerLevelReference" of complex type CodeVerticalReferenceType
	LowerLevelReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLevelReference,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightRestrictionLevelPropertyType Named source named complex type "FlightRestrictionLevelPropertyType"
type FlightRestrictionLevelPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightRestrictionLevel" of complex type FlightRestrictionLevelType
	FlightRestrictionLevel *FlightRestrictionLevelType `xml:"FlightRestrictionLevel,omitempty"`
}

// FlightRestrictionRouteType Named source named complex type "FlightRestrictionRouteType"
type FlightRestrictionRouteType struct {
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

	// generated from group "aixm:FlightRestrictionRoutePropertyGroup" with order 5586 depth 2
	Group_FlightRestrictionRoutePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_27 `xml:"extension,omitempty"`
}

// A_extension_27 Named source within outer element "extension"
type A_extension_27 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightRestrictionRouteExtension" of complex type AbstractExtensionType
	AbstractFlightRestrictionRouteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightRestrictionRouteExtension,omitempty"`
}

// Group_FlightRestrictionRoutePropertyGroup UnNamed source named group "FlightRestrictionRoutePropertyGroup"
type Group_FlightRestrictionRoutePropertyGroup struct {

	// insertion point for fields

	// generated from element "priorPermission" of complex type CodeYesNoType
	PriorPermission *aixmdatatypes.CodeYesNoType `xml:"priorPermission,omitempty"`

	// generated from element "routeElement" of complex type FlightRoutingElementPropertyType
	RouteElement []*FlightRoutingElementPropertyType `xml:"routeElement,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightRestrictionRoutePropertyType Named source named complex type "FlightRestrictionRoutePropertyType"
type FlightRestrictionRoutePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightRestrictionRoute" of complex type FlightRestrictionRouteType
	FlightRestrictionRoute *FlightRestrictionRouteType `xml:"FlightRestrictionRoute,omitempty"`
}

// FlightRoutingElementType Named source named complex type "FlightRoutingElementType"
type FlightRoutingElementType struct {
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

	// generated from group "aixm:FlightRoutingElementPropertyGroup" with order 5605 depth 2
	Group_FlightRoutingElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_28 `xml:"extension,omitempty"`
}

// A_extension_28 Named source within outer element "extension"
type A_extension_28 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightRoutingElementExtension" of complex type AbstractExtensionType
	AbstractFlightRoutingElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightRoutingElementExtension,omitempty"`
}

// Group_FlightRoutingElementPropertyGroup UnNamed source named group "FlightRoutingElementPropertyGroup"
type Group_FlightRoutingElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "orderNumber" of complex type NoSequenceType
	OrderNumber *aixmdatatypes.NoSequenceType `xml:"orderNumber,omitempty"`

	// generated from element "speed" of complex type ValSpeedType
	Speed *aixmdatatypes.ValSpeedType `xml:"speed,omitempty"`

	// generated from element "speedReference" of complex type CodeSpeedReferenceType
	SpeedReference *aixmdatatypes.CodeSpeedReferenceType `xml:"speedReference,omitempty"`

	// generated from element "speedCriteria" of complex type CodeComparisonType
	SpeedCriteria *aixmdatatypes.CodeComparisonType `xml:"speedCriteria,omitempty"`

	// generated from element "flightLevel" of complex type FlightRestrictionLevelPropertyType
	FlightLevel []*FlightRestrictionLevelPropertyType `xml:"flightLevel,omitempty"`

	// generated from element "element_directFlightElement" of complex type DirectFlightSegmentPropertyType
	Element_directFlightElement *DirectFlightSegmentPropertyType `xml:"element_directFlightElement,omitempty"`

	// generated from element "element_routePortionElement" of complex type RoutePortionPropertyType
	Element_routePortionElement *RoutePortionPropertyType `xml:"element_routePortionElement,omitempty"`

	// generated from element "element_standardInstrumentDepartureElement" of complex type StandardInstrumentDeparturePropertyType
	Element_standardInstrumentDepartureElement *StandardInstrumentDeparturePropertyType `xml:"element_standardInstrumentDepartureElement,omitempty"`

	// generated from element "element_standardInstrumentArrivalElement" of complex type StandardInstrumentArrivalPropertyType
	Element_standardInstrumentArrivalElement *StandardInstrumentArrivalPropertyType `xml:"element_standardInstrumentArrivalElement,omitempty"`

	// generated from element "pointElement_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	PointElement_fixDesignatedPoint *DesignatedPointPropertyType `xml:"pointElement_fixDesignatedPoint,omitempty"`

	// generated from element "pointElement_navaidSystem" of complex type NavaidPropertyType
	PointElement_navaidSystem *NavaidPropertyType `xml:"pointElement_navaidSystem,omitempty"`

	// generated from element "pointElement_position" of complex type PointPropertyType
	PointElement_position *PointPropertyType `xml:"pointElement_position,omitempty"`

	// generated from element "pointElement_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	PointElement_runwayPoint *RunwayCentrelinePointPropertyType `xml:"pointElement_runwayPoint,omitempty"`

	// generated from element "pointElement_aimingPoint" of complex type TouchDownLiftOffPropertyType
	PointElement_aimingPoint *TouchDownLiftOffPropertyType `xml:"pointElement_aimingPoint,omitempty"`

	// generated from element "pointElement_airportReferencePoint" of complex type AirportHeliportPropertyType
	PointElement_airportReferencePoint *AirportHeliportPropertyType `xml:"pointElement_airportReferencePoint,omitempty"`

	// generated from element "element_airspaceElement" of complex type AirspacePropertyType
	Element_airspaceElement *AirspacePropertyType `xml:"element_airspaceElement,omitempty"`

	// generated from element "element_airportHeliportElement" of complex type AirportHeliportPropertyType
	Element_airportHeliportElement *AirportHeliportPropertyType `xml:"element_airportHeliportElement,omitempty"`

	// generated from element "element_aerialRefuellingElement" of complex type AerialRefuellingPropertyType
	Element_aerialRefuellingElement *AerialRefuellingPropertyType `xml:"element_aerialRefuellingElement,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightRoutingElementPropertyType Named source named complex type "FlightRoutingElementPropertyType"
type FlightRoutingElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightRoutingElement" of complex type FlightRoutingElementType
	FlightRoutingElement *FlightRoutingElementType `xml:"FlightRoutingElement,omitempty"`
}

// RulesProceduresType Named source named complex type "RulesProceduresType"
type RulesProceduresType struct {
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
	FeatureMetadata *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType `xml:"featureMetadata,omitempty"`

	// generated from element "timeSlice" of complex type RulesProceduresTimeSlicePropertyType
	TimeSlice []*RulesProceduresTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RulesProceduresTimeSlicePropertyType Named source named complex type "RulesProceduresTimeSlicePropertyType"
type RulesProceduresTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RulesProceduresTimeSlice" of complex type RulesProceduresTimeSliceType
	RulesProceduresTimeSlice *RulesProceduresTimeSliceType `xml:"RulesProceduresTimeSlice,omitempty"`
}

// RulesProceduresTimeSliceType Named source named complex type "RulesProceduresTimeSliceType"
type RulesProceduresTimeSliceType struct {
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
	Interpretation string `xml:"interpretation,omitempty"`

	// generated from element "sequenceNumber" of simple type sequenceNumber_SimpleType
	SequenceNumber int `xml:"sequenceNumber,omitempty"`

	// generated from element "correctionNumber" of simple type correctionNumber_SimpleType
	CorrectionNumber int `xml:"correctionNumber,omitempty"`

	// generated from element "timeSliceMetadata" of complex type FeatureTimeSliceMetadataPropertyType
	TimeSliceMetadata *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType `xml:"timeSliceMetadata,omitempty"`

	// generated from element "featureLifetime" of complex type TimePrimitivePropertyType
	FeatureLifetime *gml.TimePrimitivePropertyType `xml:"featureLifetime,omitempty"`

	// generated from group "aixm:RulesProceduresPropertyGroup" with order 5652 depth 2
	Group_RulesProceduresPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_29 `xml:"extension,omitempty"`
}

// A_extension_29 Named source within outer element "extension"
type A_extension_29 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRulesProceduresExtension" of complex type AbstractExtensionType
	AbstractRulesProceduresExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRulesProceduresExtension,omitempty"`
}

// Group_RulesProceduresPropertyGroup UnNamed source named group "RulesProceduresPropertyGroup"
type Group_RulesProceduresPropertyGroup struct {

	// insertion point for fields

	// generated from element "category" of complex type CodeRuleProcedureType
	Category *aixmdatatypes.CodeRuleProcedureType `xml:"category,omitempty"`

	// generated from element "title" of complex type CodeRuleProcedureTitleType
	Title *aixmdatatypes.CodeRuleProcedureTitleType `xml:"title,omitempty"`

	// generated from element "content" of complex type XHTMLType
	Content *aixmdatatypes.XHTMLType `xml:"content,omitempty"`

	// generated from element "affectedLocation" of complex type AirportHeliportPropertyType
	AffectedLocation []*AirportHeliportPropertyType `xml:"affectedLocation,omitempty"`

	// generated from element "affectedArea" of complex type AirspacePropertyType
	AffectedArea []*AirspacePropertyType `xml:"affectedArea,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ContactInformationType Named source named complex type "ContactInformationType"
type ContactInformationType struct {
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

	// generated from element "title" of complex type TextNameType
	Title *aixmdatatypes.TextNameType `xml:"title,omitempty"`

	// generated from element "address" of complex type PostalAddressPropertyType
	Address []*PostalAddressPropertyType `xml:"address,omitempty"`

	// generated from element "networkNode" of complex type OnlineContactPropertyType
	NetworkNode []*OnlineContactPropertyType `xml:"networkNode,omitempty"`

	// generated from element "phoneFax" of complex type TelephoneContactPropertyType
	PhoneFax []*TelephoneContactPropertyType `xml:"phoneFax,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_30 `xml:"extension,omitempty"`
}

// A_extension_30 Named source within outer element "extension"
type A_extension_30 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractContactInformationExtension" of complex type AbstractExtensionType
	AbstractContactInformationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractContactInformationExtension,omitempty"`
}

// Group_ContactInformationPropertyGroup UnNamed source named group "ContactInformationPropertyGroup"
type Group_ContactInformationPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "title" of complex type TextNameType
	Title *aixmdatatypes.TextNameType `xml:"title,omitempty"`

	// generated from element "address" of complex type PostalAddressPropertyType
	Address []*PostalAddressPropertyType `xml:"address,omitempty"`

	// generated from element "networkNode" of complex type OnlineContactPropertyType
	NetworkNode []*OnlineContactPropertyType `xml:"networkNode,omitempty"`

	// generated from element "phoneFax" of complex type TelephoneContactPropertyType
	PhoneFax []*TelephoneContactPropertyType `xml:"phoneFax,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ContactInformationPropertyType Named source named complex type "ContactInformationPropertyType"
type ContactInformationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ContactInformation" of complex type ContactInformationType
	ContactInformation *ContactInformationType `xml:"ContactInformation,omitempty"`
}

// OnlineContactType Named source named complex type "OnlineContactType"
type OnlineContactType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6216 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:OnlineContactPropertyGroup" with order 6217 depth 2
	Group_OnlineContactPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_31 `xml:"extension,omitempty"`
}

// A_extension_31 Named source within outer element "extension"
type A_extension_31 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractOnlineContactExtension" of complex type AbstractExtensionType
	AbstractOnlineContactExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractOnlineContactExtension,omitempty"`
}

// Group_OnlineContactPropertyGroup UnNamed source named group "OnlineContactPropertyGroup"
type Group_OnlineContactPropertyGroup struct {

	// insertion point for fields

	// generated from element "network" of complex type CodeTelecomNetworkType
	Network *aixmdatatypes.CodeTelecomNetworkType `xml:"network,omitempty"`

	// generated from element "linkage" of complex type TextAddressType
	Linkage *aixmdatatypes.TextAddressType `xml:"linkage,omitempty"`

	// generated from element "protocol" of complex type TextNameType
	Protocol *aixmdatatypes.TextNameType `xml:"protocol,omitempty"`

	// generated from element "eMail" of complex type TextAddressType
	EMail *aixmdatatypes.TextAddressType `xml:"eMail,omitempty"`
}

// OnlineContactPropertyType Named source named complex type "OnlineContactPropertyType"
type OnlineContactPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "OnlineContact" of complex type OnlineContactType
	OnlineContact *OnlineContactType `xml:"OnlineContact,omitempty"`
}

// PostalAddressType Named source named complex type "PostalAddressType"
type PostalAddressType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6237 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:PostalAddressPropertyGroup" with order 6238 depth 2
	Group_PostalAddressPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_32 `xml:"extension,omitempty"`
}

// A_extension_32 Named source within outer element "extension"
type A_extension_32 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractPostalAddressExtension" of complex type AbstractExtensionType
	AbstractPostalAddressExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPostalAddressExtension,omitempty"`
}

// Group_PostalAddressPropertyGroup UnNamed source named group "PostalAddressPropertyGroup"
type Group_PostalAddressPropertyGroup struct {

	// insertion point for fields

	// generated from element "deliveryPoint" of complex type TextAddressType
	DeliveryPoint *aixmdatatypes.TextAddressType `xml:"deliveryPoint,omitempty"`

	// generated from element "city" of complex type TextNameType
	City *aixmdatatypes.TextNameType `xml:"city,omitempty"`

	// generated from element "administrativeArea" of complex type TextNameType
	AdministrativeArea *aixmdatatypes.TextNameType `xml:"administrativeArea,omitempty"`

	// generated from element "postalCode" of complex type TextNameType
	PostalCode *aixmdatatypes.TextNameType `xml:"postalCode,omitempty"`

	// generated from element "country" of complex type TextNameType
	Country *aixmdatatypes.TextNameType `xml:"country,omitempty"`
}

// PostalAddressPropertyType Named source named complex type "PostalAddressPropertyType"
type PostalAddressPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "PostalAddress" of complex type PostalAddressType
	PostalAddress *PostalAddressType `xml:"PostalAddress,omitempty"`
}

// TelephoneContactType Named source named complex type "TelephoneContactType"
type TelephoneContactType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6259 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:TelephoneContactPropertyGroup" with order 6260 depth 2
	Group_TelephoneContactPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_33 `xml:"extension,omitempty"`
}

// A_extension_33 Named source within outer element "extension"
type A_extension_33 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractTelephoneContactExtension" of complex type AbstractExtensionType
	AbstractTelephoneContactExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTelephoneContactExtension,omitempty"`
}

// Group_TelephoneContactPropertyGroup UnNamed source named group "TelephoneContactPropertyGroup"
type Group_TelephoneContactPropertyGroup struct {

	// insertion point for fields

	// generated from element "voice" of complex type TextPhoneType
	Voice *aixmdatatypes.TextPhoneType `xml:"voice,omitempty"`

	// generated from element "facsimile" of complex type TextPhoneType
	Facsimile *aixmdatatypes.TextPhoneType `xml:"facsimile,omitempty"`
}

// TelephoneContactPropertyType Named source named complex type "TelephoneContactPropertyType"
type TelephoneContactPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "TelephoneContact" of complex type TelephoneContactType
	TelephoneContact *TelephoneContactType `xml:"TelephoneContact,omitempty"`
}

// AircraftCharacteristicType Named source named complex type "AircraftCharacteristicType"
type AircraftCharacteristicType struct {
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

	// generated from group "aixm:AircraftCharacteristicPropertyGroup" with order 6278 depth 2
	Group_AircraftCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_34 `xml:"extension,omitempty"`
}

// A_extension_34 Named source within outer element "extension"
type A_extension_34 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAircraftCharacteristicExtension" of complex type AbstractExtensionType
	AbstractAircraftCharacteristicExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAircraftCharacteristicExtension,omitempty"`
}

// Group_AircraftCharacteristicPropertyGroup UnNamed source named group "AircraftCharacteristicPropertyGroup"
type Group_AircraftCharacteristicPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAircraftType
	Type_XSD *aixmdatatypes.CodeAircraftType `xml:"type,omitempty"`

	// generated from element "engine" of complex type CodeAircraftEngineType
	Engine *aixmdatatypes.CodeAircraftEngineType `xml:"engine,omitempty"`

	// generated from element "numberEngine" of complex type CodeAircraftEngineNumberType
	NumberEngine *aixmdatatypes.CodeAircraftEngineNumberType `xml:"numberEngine,omitempty"`

	// generated from element "typeAircraftICAO" of complex type CodeAircraftICAOType
	TypeAircraftICAO *aixmdatatypes.CodeAircraftICAOType `xml:"typeAircraftICAO,omitempty"`

	// generated from element "aircraftLandingCategory" of complex type CodeAircraftCategoryType
	AircraftLandingCategory *aixmdatatypes.CodeAircraftCategoryType `xml:"aircraftLandingCategory,omitempty"`

	// generated from element "wingSpan" of complex type ValDistanceType
	WingSpan *aixmdatatypes.ValDistanceType `xml:"wingSpan,omitempty"`

	// generated from element "wingSpanInterpretation" of complex type CodeValueInterpretationType
	WingSpanInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"wingSpanInterpretation,omitempty"`

	// generated from element "classWingSpan" of complex type CodeAircraftWingspanClassType
	ClassWingSpan *aixmdatatypes.CodeAircraftWingspanClassType `xml:"classWingSpan,omitempty"`

	// generated from element "weight" of complex type ValWeightType
	Weight *aixmdatatypes.ValWeightType `xml:"weight,omitempty"`

	// generated from element "weightInterpretation" of complex type CodeValueInterpretationType
	WeightInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"weightInterpretation,omitempty"`

	// generated from element "passengers" of complex type NoNumberType
	Passengers *aixmdatatypes.NoNumberType `xml:"passengers,omitempty"`

	// generated from element "passengersInterpretation" of complex type CodeValueInterpretationType
	PassengersInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"passengersInterpretation,omitempty"`

	// generated from element "speed" of complex type ValSpeedType
	Speed *aixmdatatypes.ValSpeedType `xml:"speed,omitempty"`

	// generated from element "speedInterpretation" of complex type CodeValueInterpretationType
	SpeedInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"speedInterpretation,omitempty"`

	// generated from element "wakeTurbulence" of complex type CodeWakeTurbulenceType
	WakeTurbulence *aixmdatatypes.CodeWakeTurbulenceType `xml:"wakeTurbulence,omitempty"`

	// generated from element "navigationEquipment" of complex type CodeNavigationEquipmentType
	NavigationEquipment *aixmdatatypes.CodeNavigationEquipmentType `xml:"navigationEquipment,omitempty"`

	// generated from element "navigationSpecification" of complex type CodeNavigationSpecificationType
	NavigationSpecification *aixmdatatypes.CodeNavigationSpecificationType `xml:"navigationSpecification,omitempty"`

	// generated from element "verticalSeparationCapability" of complex type CodeRVSMType
	VerticalSeparationCapability *aixmdatatypes.CodeRVSMType `xml:"verticalSeparationCapability,omitempty"`

	// generated from element "antiCollisionAndSeparationEquipment" of complex type CodeEquipmentAntiCollisionType
	AntiCollisionAndSeparationEquipment *aixmdatatypes.CodeEquipmentAntiCollisionType `xml:"antiCollisionAndSeparationEquipment,omitempty"`

	// generated from element "communicationEquipment" of complex type CodeCommunicationModeType
	CommunicationEquipment *aixmdatatypes.CodeCommunicationModeType `xml:"communicationEquipment,omitempty"`

	// generated from element "surveillanceEquipment" of complex type CodeTransponderType
	SurveillanceEquipment *aixmdatatypes.CodeTransponderType `xml:"surveillanceEquipment,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AircraftCharacteristicPropertyType Named source named complex type "AircraftCharacteristicPropertyType"
type AircraftCharacteristicPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AircraftCharacteristic" of complex type AircraftCharacteristicType
	AircraftCharacteristic *AircraftCharacteristicType `xml:"AircraftCharacteristic,omitempty"`
}

// FlightCharacteristicType Named source named complex type "FlightCharacteristicType"
type FlightCharacteristicType struct {
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

	// generated from group "aixm:FlightCharacteristicPropertyGroup" with order 6315 depth 2
	Group_FlightCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_35 `xml:"extension,omitempty"`
}

// A_extension_35 Named source within outer element "extension"
type A_extension_35 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFlightCharacteristicExtension" of complex type AbstractExtensionType
	AbstractFlightCharacteristicExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFlightCharacteristicExtension,omitempty"`
}

// Group_FlightCharacteristicPropertyGroup UnNamed source named group "FlightCharacteristicPropertyGroup"
type Group_FlightCharacteristicPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeFlightType
	Type_XSD *aixmdatatypes.CodeFlightType `xml:"type,omitempty"`

	// generated from element "rule" of complex type CodeFlightRuleType
	Rule *aixmdatatypes.CodeFlightRuleType `xml:"rule,omitempty"`

	// generated from element "status" of complex type CodeFlightStatusType
	Status *aixmdatatypes.CodeFlightStatusType `xml:"status,omitempty"`

	// generated from element "military" of complex type CodeMilitaryStatusType
	Military *aixmdatatypes.CodeMilitaryStatusType `xml:"military,omitempty"`

	// generated from element "origin" of complex type CodeFlightOriginType
	Origin *aixmdatatypes.CodeFlightOriginType `xml:"origin,omitempty"`

	// generated from element "purpose" of complex type CodeFlightPurposeType
	Purpose *aixmdatatypes.CodeFlightPurposeType `xml:"purpose,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FlightCharacteristicPropertyType Named source named complex type "FlightCharacteristicPropertyType"
type FlightCharacteristicPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FlightCharacteristic" of complex type FlightCharacteristicType
	FlightCharacteristic *FlightCharacteristicType `xml:"FlightCharacteristic,omitempty"`
}

// AirspaceLayerType Named source named complex type "AirspaceLayerType"
type AirspaceLayerType struct {
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

	// generated from group "aixm:AirspaceLayerPropertyGroup" with order 6337 depth 2
	Group_AirspaceLayerPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_36 `xml:"extension,omitempty"`
}

// A_extension_36 Named source within outer element "extension"
type A_extension_36 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirspaceLayerExtension" of complex type AbstractExtensionType
	AbstractAirspaceLayerExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceLayerExtension,omitempty"`
}

// Group_AirspaceLayerPropertyGroup UnNamed source named group "AirspaceLayerPropertyGroup"
type Group_AirspaceLayerPropertyGroup struct {

	// insertion point for fields

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "altitudeInterpretation" of complex type CodeAltitudeUseType
	AltitudeInterpretation *aixmdatatypes.CodeAltitudeUseType `xml:"altitudeInterpretation,omitempty"`

	// generated from element "discreteLevelSeries" of complex type StandardLevelColumnPropertyType
	DiscreteLevelSeries *StandardLevelColumnPropertyType `xml:"discreteLevelSeries,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirspaceLayerPropertyType Named source named complex type "AirspaceLayerPropertyType"
type AirspaceLayerPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceLayer" of complex type AirspaceLayerType
	AirspaceLayer *AirspaceLayerType `xml:"AirspaceLayer,omitempty"`
}

// MeteorologyType Named source named complex type "MeteorologyType"
type MeteorologyType struct {
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

	// generated from group "aixm:MeteorologyPropertyGroup" with order 6426 depth 2
	Group_MeteorologyPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_37 `xml:"extension,omitempty"`
}

// A_extension_37 Named source within outer element "extension"
type A_extension_37 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMeteorologyExtension" of complex type AbstractExtensionType
	AbstractMeteorologyExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMeteorologyExtension,omitempty"`
}

// Group_MeteorologyPropertyGroup UnNamed source named group "MeteorologyPropertyGroup"
type Group_MeteorologyPropertyGroup struct {

	// insertion point for fields

	// generated from element "flightConditions" of complex type CodeMeteoConditionsType
	FlightConditions *aixmdatatypes.CodeMeteoConditionsType `xml:"flightConditions,omitempty"`

	// generated from element "visibility" of complex type ValDistanceType
	Visibility *aixmdatatypes.ValDistanceType `xml:"visibility,omitempty"`

	// generated from element "visibilityInterpretation" of complex type CodeValueInterpretationType
	VisibilityInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"visibilityInterpretation,omitempty"`

	// generated from element "runwayVisualRange" of complex type ValDistanceType
	RunwayVisualRange *aixmdatatypes.ValDistanceType `xml:"runwayVisualRange,omitempty"`

	// generated from element "runwayVisualRangeInterpretation" of complex type CodeValueInterpretationType
	RunwayVisualRangeInterpretation *aixmdatatypes.CodeValueInterpretationType `xml:"runwayVisualRangeInterpretation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MeteorologyPropertyType Named source named complex type "MeteorologyPropertyType"
type MeteorologyPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Meteorology" of complex type MeteorologyType
	Meteorology *MeteorologyType `xml:"Meteorology,omitempty"`
}

// RadioFrequencyAreaPropertyType Named source named complex type "RadioFrequencyAreaPropertyType"
type RadioFrequencyAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AbstractPropertiesWithScheduleType Named source named complex type "AbstractPropertiesWithScheduleType"
type AbstractPropertiesWithScheduleType struct {
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

// Group_PropertiesWithSchedulePropertyGroup UnNamed source named group "PropertiesWithSchedulePropertyGroup"
type Group_PropertiesWithSchedulePropertyGroup struct {

	// insertion point for fields

	// generated from element "timeInterval" of complex type TimesheetPropertyType
	TimeInterval []*TimesheetPropertyType `xml:"timeInterval,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "specialDateAuthority" of complex type OrganisationAuthorityPropertyType
	SpecialDateAuthority []*OrganisationAuthorityPropertyType `xml:"specialDateAuthority,omitempty"`
}

// TimesheetType Named source named complex type "TimesheetType"
type TimesheetType struct {
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

	// generated from group "aixm:TimesheetPropertyGroup" with order 6523 depth 2
	Group_TimesheetPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_38 `xml:"extension,omitempty"`
}

// A_extension_38 Named source within outer element "extension"
type A_extension_38 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTimesheetExtension" of complex type AbstractExtensionType
	AbstractTimesheetExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTimesheetExtension,omitempty"`
}

// Group_TimesheetPropertyGroup UnNamed source named group "TimesheetPropertyGroup"
type Group_TimesheetPropertyGroup struct {

	// insertion point for fields

	// generated from element "timeReference" of complex type CodeTimeReferenceType
	TimeReference *aixmdatatypes.CodeTimeReferenceType `xml:"timeReference,omitempty"`

	// generated from element "startDate" of complex type DateMonthDayType
	StartDate *aixmdatatypes.DateMonthDayType `xml:"startDate,omitempty"`

	// generated from element "endDate" of complex type DateMonthDayType
	EndDate *aixmdatatypes.DateMonthDayType `xml:"endDate,omitempty"`

	// generated from element "day" of complex type CodeDayType
	Day *aixmdatatypes.CodeDayType `xml:"day,omitempty"`

	// generated from element "dayTil" of complex type CodeDayType
	DayTil *aixmdatatypes.CodeDayType `xml:"dayTil,omitempty"`

	// generated from element "startTime" of complex type TimeType
	StartTime *aixmdatatypes.TimeType `xml:"startTime,omitempty"`

	// generated from element "startEvent" of complex type CodeTimeEventType
	StartEvent *aixmdatatypes.CodeTimeEventType `xml:"startEvent,omitempty"`

	// generated from element "startTimeRelativeEvent" of complex type ValDurationType
	StartTimeRelativeEvent *aixmdatatypes.ValDurationType `xml:"startTimeRelativeEvent,omitempty"`

	// generated from element "startEventInterpretation" of complex type CodeTimeEventCombinationType
	StartEventInterpretation *aixmdatatypes.CodeTimeEventCombinationType `xml:"startEventInterpretation,omitempty"`

	// generated from element "endTime" of complex type TimeType
	EndTime *aixmdatatypes.TimeType `xml:"endTime,omitempty"`

	// generated from element "endEvent" of complex type CodeTimeEventType
	EndEvent *aixmdatatypes.CodeTimeEventType `xml:"endEvent,omitempty"`

	// generated from element "endTimeRelativeEvent" of complex type ValDurationType
	EndTimeRelativeEvent *aixmdatatypes.ValDurationType `xml:"endTimeRelativeEvent,omitempty"`

	// generated from element "endEventInterpretation" of complex type CodeTimeEventCombinationType
	EndEventInterpretation *aixmdatatypes.CodeTimeEventCombinationType `xml:"endEventInterpretation,omitempty"`

	// generated from element "daylightSavingAdjust" of complex type CodeYesNoType
	DaylightSavingAdjust *aixmdatatypes.CodeYesNoType `xml:"daylightSavingAdjust,omitempty"`

	// generated from element "excluded" of complex type CodeYesNoType
	Excluded *aixmdatatypes.CodeYesNoType `xml:"excluded,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TimesheetPropertyType Named source named complex type "TimesheetPropertyType"
type TimesheetPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Timesheet" of complex type TimesheetType
	Timesheet *TimesheetType `xml:"Timesheet,omitempty"`
}

// StandardLevelColumnPropertyType Named source named complex type "StandardLevelColumnPropertyType"
type StandardLevelColumnPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AltitudeAdjustmentType Named source named complex type "AltitudeAdjustmentType"
type AltitudeAdjustmentType struct {
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

	// generated from group "aixm:AltitudeAdjustmentPropertyGroup" with order 6656 depth 2
	Group_AltitudeAdjustmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_39 `xml:"extension,omitempty"`
}

// A_extension_39 Named source within outer element "extension"
type A_extension_39 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAltitudeAdjustmentExtension" of complex type AbstractExtensionType
	AbstractAltitudeAdjustmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAltitudeAdjustmentExtension,omitempty"`
}

// Group_AltitudeAdjustmentPropertyGroup UnNamed source named group "AltitudeAdjustmentPropertyGroup"
type Group_AltitudeAdjustmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "altitudeAdjustmentType" of complex type CodeAltitudeAdjustmentType
	AltitudeAdjustmentType *aixmdatatypes.CodeAltitudeAdjustmentType `xml:"altitudeAdjustmentType,omitempty"`

	// generated from element "primaryAlternateMinimum" of complex type CodeYesNoType
	PrimaryAlternateMinimum *aixmdatatypes.CodeYesNoType `xml:"primaryAlternateMinimum,omitempty"`

	// generated from element "altitudeAdjustment" of complex type ValDistanceVerticalType
	AltitudeAdjustment *aixmdatatypes.ValDistanceVerticalType `xml:"altitudeAdjustment,omitempty"`

	// generated from element "localRemoteCode" of complex type CodeYesNoType
	LocalRemoteCode *aixmdatatypes.CodeYesNoType `xml:"localRemoteCode,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AltitudeAdjustmentPropertyType Named source named complex type "AltitudeAdjustmentPropertyType"
type AltitudeAdjustmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AltitudeAdjustment" of complex type AltitudeAdjustmentType
	AltitudeAdjustment *AltitudeAdjustmentType `xml:"AltitudeAdjustment,omitempty"`
}

// ObstacleAssessmentAreaType Named source named complex type "ObstacleAssessmentAreaType"
type ObstacleAssessmentAreaType struct {
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

	// generated from group "aixm:ObstacleAssessmentAreaPropertyGroup" with order 6713 depth 2
	Group_ObstacleAssessmentAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_40 `xml:"extension,omitempty"`
}

// A_extension_40 Named source within outer element "extension"
type A_extension_40 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractObstacleAssessmentAreaExtension" of complex type AbstractExtensionType
	AbstractObstacleAssessmentAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractObstacleAssessmentAreaExtension,omitempty"`
}

// Group_ObstacleAssessmentAreaPropertyGroup UnNamed source named group "ObstacleAssessmentAreaPropertyGroup"
type Group_ObstacleAssessmentAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeObstacleAssessmentSurfaceType
	Type_XSD *aixmdatatypes.CodeObstacleAssessmentSurfaceType `xml:"type,omitempty"`

	// generated from element "sectionNumber" of complex type NoNumberType
	SectionNumber *aixmdatatypes.NoNumberType `xml:"sectionNumber,omitempty"`

	// generated from element "slope" of complex type ValSlopeType
	Slope *aixmdatatypes.ValSlopeType `xml:"slope,omitempty"`

	// generated from element "assessedAltitude" of complex type ValDistanceVerticalType
	AssessedAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"assessedAltitude,omitempty"`

	// generated from element "slopeLowerAltitude" of complex type ValDistanceVerticalType
	SlopeLowerAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"slopeLowerAltitude,omitempty"`

	// generated from element "gradientLowHigh" of complex type ValSlopeType
	GradientLowHigh *aixmdatatypes.ValSlopeType `xml:"gradientLowHigh,omitempty"`

	// generated from element "surfaceZone" of complex type CodeObstructionIdSurfaceZoneType
	SurfaceZone *aixmdatatypes.CodeObstructionIdSurfaceZoneType `xml:"surfaceZone,omitempty"`

	// generated from element "safetyRegulation" of complex type TextNameType
	SafetyRegulation *aixmdatatypes.TextNameType `xml:"safetyRegulation,omitempty"`

	// generated from element "aircraftCategory" of complex type AircraftCharacteristicPropertyType
	AircraftCategory []*AircraftCharacteristicPropertyType `xml:"aircraftCategory,omitempty"`

	// generated from element "significantObstacle" of complex type ObstructionPropertyType
	SignificantObstacle []*ObstructionPropertyType `xml:"significantObstacle,omitempty"`

	// generated from element "surface" of complex type SurfacePropertyType
	Surface *SurfacePropertyType `xml:"surface,omitempty"`

	// generated from element "startingCurve" of complex type CurvePropertyType
	StartingCurve *CurvePropertyType `xml:"startingCurve,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ObstacleAssessmentAreaPropertyType Named source named complex type "ObstacleAssessmentAreaPropertyType"
type ObstacleAssessmentAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ObstacleAssessmentArea" of complex type ObstacleAssessmentAreaType
	ObstacleAssessmentArea *ObstacleAssessmentAreaType `xml:"ObstacleAssessmentArea,omitempty"`
}

// ObstaclePlacementType Named source named complex type "ObstaclePlacementType"
type ObstaclePlacementType struct {
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

	// generated from group "aixm:ObstaclePlacementPropertyGroup" with order 6741 depth 2
	Group_ObstaclePlacementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_41 `xml:"extension,omitempty"`
}

// A_extension_41 Named source within outer element "extension"
type A_extension_41 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractObstaclePlacementExtension" of complex type AbstractExtensionType
	AbstractObstaclePlacementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractObstaclePlacementExtension,omitempty"`
}

// Group_ObstaclePlacementPropertyGroup UnNamed source named group "ObstaclePlacementPropertyGroup"
type Group_ObstaclePlacementPropertyGroup struct {

	// insertion point for fields

	// generated from element "obstacleBearing" of complex type ValBearingType
	ObstacleBearing *aixmdatatypes.ValBearingType `xml:"obstacleBearing,omitempty"`

	// generated from element "obstacleDistance" of complex type ValDistanceType
	ObstacleDistance *aixmdatatypes.ValDistanceType `xml:"obstacleDistance,omitempty"`

	// generated from element "pointType" of complex type TextNameType
	PointType *aixmdatatypes.TextNameType `xml:"pointType,omitempty"`

	// generated from element "obstaclePlacement" of complex type CodeSideType
	ObstaclePlacement *aixmdatatypes.CodeSideType `xml:"obstaclePlacement,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ObstaclePlacementPropertyType Named source named complex type "ObstaclePlacementPropertyType"
type ObstaclePlacementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ObstaclePlacement" of complex type ObstaclePlacementType
	ObstaclePlacement *ObstaclePlacementType `xml:"ObstaclePlacement,omitempty"`
}

// ObstructionType Named source named complex type "ObstructionType"
type ObstructionType struct {
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

	// generated from group "aixm:ObstructionPropertyGroup" with order 6761 depth 2
	Group_ObstructionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_42 `xml:"extension,omitempty"`
}

// A_extension_42 Named source within outer element "extension"
type A_extension_42 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractObstructionExtension" of complex type AbstractExtensionType
	AbstractObstructionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractObstructionExtension,omitempty"`
}

// Group_ObstructionPropertyGroup UnNamed source named group "ObstructionPropertyGroup"
type Group_ObstructionPropertyGroup struct {

	// insertion point for fields

	// generated from element "requiredClearance" of complex type ValDistanceType
	RequiredClearance *aixmdatatypes.ValDistanceType `xml:"requiredClearance,omitempty"`

	// generated from element "minimumAltitude" of complex type ValDistanceVerticalType
	MinimumAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumAltitude,omitempty"`

	// generated from element "surfacePenetration" of complex type CodeYesNoType
	SurfacePenetration *aixmdatatypes.CodeYesNoType `xml:"surfacePenetration,omitempty"`

	// generated from element "slopePenetration" of complex type ValAngleType
	SlopePenetration *aixmdatatypes.ValAngleType `xml:"slopePenetration,omitempty"`

	// generated from element "controlling" of complex type CodeYesNoType
	Controlling *aixmdatatypes.CodeYesNoType `xml:"controlling,omitempty"`

	// generated from element "closeIn" of complex type CodeYesNoType
	CloseIn *aixmdatatypes.CodeYesNoType `xml:"closeIn,omitempty"`

	// generated from element "theVerticalStructure" of complex type VerticalStructurePropertyType
	TheVerticalStructure *VerticalStructurePropertyType `xml:"theVerticalStructure,omitempty"`

	// generated from element "adjustment" of complex type AltitudeAdjustmentPropertyType
	Adjustment []*AltitudeAdjustmentPropertyType `xml:"adjustment,omitempty"`

	// generated from element "obstaclePlacement" of complex type ObstaclePlacementPropertyType
	ObstaclePlacement []*ObstaclePlacementPropertyType `xml:"obstaclePlacement,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ObstructionPropertyType Named source named complex type "ObstructionPropertyType"
type ObstructionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Obstruction" of complex type ObstructionType
	Obstruction *ObstructionType `xml:"Obstruction,omitempty"`
}
