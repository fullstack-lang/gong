// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeICAOCountryType
var _ *gml.CodeType

// insertion point for enums declaration

// insertion point for gongstructs declarations

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

	// generated from group "aixm:SegmentPointPropertyGroup" with order 455 depth 2
	Group_SegmentPointPropertyGroup

	// generated from group "aixm:AerialRefuellingPointPropertyGroup" with order 456 depth 2
	Group_AerialRefuellingPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension `xml:"extension,omitempty"`
}

// A_extension Named source within outer element "extension"
type A_extension struct {
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

// AirportHeliportPropertyType Named source named complex type "AirportHeliportPropertyType"
type AirportHeliportPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AltimeterSourcePropertyType Named source named complex type "AltimeterSourcePropertyType"
type AltimeterSourcePropertyType struct {
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

// RunwayDirectionPropertyType Named source named complex type "RunwayDirectionPropertyType"
type RunwayDirectionPropertyType struct {
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

	// generated from group "aixm:CurvePropertyGroup" with order 2764 depth 2
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

	// generated from group "aixm:CurvePropertyGroup" with order 2764 depth 2
	Group_CurvePropertyGroup

	// generated from group "aixm:ElevatedCurvePropertyGroup" with order 2775 depth 2
	Group_ElevatedCurvePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_1 `xml:"extension,omitempty"`
}

// A_extension_1 Named source within outer element "extension"
type A_extension_1 struct {
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

	// generated from group "aixm:PointPropertyGroup" with order 2832 depth 2
	Group_PointPropertyGroup

	// generated from group "aixm:ElevatedPointPropertyGroup" with order 2794 depth 2
	Group_ElevatedPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2843 depth 2
	Group_SurfacePropertyGroup

	// generated from group "aixm:ElevatedSurfacePropertyGroup" with order 2813 depth 2
	Group_ElevatedSurfacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
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

	// generated from group "aixm:PointPropertyGroup" with order 2832 depth 2
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2843 depth 2
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

// HoldingPatternPropertyType Named source named complex type "HoldingPatternPropertyType"
type HoldingPatternPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// HoldingPatternType Named source named complex type "HoldingPatternType"
type HoldingPatternType struct {
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

	// generated from element "timeSlice" of complex type HoldingPatternTimeSlicePropertyType
	TimeSlice []*HoldingPatternTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// HoldingPatternTimeSlicePropertyType Named source named complex type "HoldingPatternTimeSlicePropertyType"
type HoldingPatternTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "HoldingPatternTimeSlice" of complex type HoldingPatternTimeSliceType
	HoldingPatternTimeSlice *HoldingPatternTimeSliceType `xml:"HoldingPatternTimeSlice,omitempty"`
}

// HoldingPatternTimeSliceType Named source named complex type "HoldingPatternTimeSliceType"
type HoldingPatternTimeSliceType struct {
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

	// generated from group "aixm:HoldingPatternPropertyGroup" with order 2865 depth 2
	Group_HoldingPatternPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractHoldingPatternExtension" of complex type AbstractExtensionType
	AbstractHoldingPatternExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractHoldingPatternExtension,omitempty"`
}

// Group_HoldingPatternPropertyGroup UnNamed source named group "HoldingPatternPropertyGroup"
type Group_HoldingPatternPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeHoldingUsageType
	Type_XSD *aixmdatatypes.CodeHoldingUsageType `xml:"type,omitempty"`

	// generated from element "outboundCourse" of complex type ValBearingType
	OutboundCourse *aixmdatatypes.ValBearingType `xml:"outboundCourse,omitempty"`

	// generated from element "outboundCourseType" of complex type CodeCourseType
	OutboundCourseType *aixmdatatypes.CodeCourseType `xml:"outboundCourseType,omitempty"`

	// generated from element "inboundCourse" of complex type ValBearingType
	InboundCourse *aixmdatatypes.ValBearingType `xml:"inboundCourse,omitempty"`

	// generated from element "turnDirection" of complex type CodeDirectionTurnType
	TurnDirection *aixmdatatypes.CodeDirectionTurnType `xml:"turnDirection,omitempty"`

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "speedLimit" of complex type ValSpeedType
	SpeedLimit *aixmdatatypes.ValSpeedType `xml:"speedLimit,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "nonStandardHolding" of complex type CodeYesNoType
	NonStandardHolding *aixmdatatypes.CodeYesNoType `xml:"nonStandardHolding,omitempty"`

	// generated from element "outboundLegSpan_endPoint" of complex type SegmentPointPropertyType
	OutboundLegSpan_endPoint *SegmentPointPropertyType `xml:"outboundLegSpan_endPoint,omitempty"`

	// generated from element "outboundLegSpan_endDistance" of complex type HoldingPatternDistancePropertyType
	OutboundLegSpan_endDistance *HoldingPatternDistancePropertyType `xml:"outboundLegSpan_endDistance,omitempty"`

	// generated from element "outboundLegSpan_endTime" of complex type HoldingPatternDurationPropertyType
	OutboundLegSpan_endTime *HoldingPatternDurationPropertyType `xml:"outboundLegSpan_endTime,omitempty"`

	// generated from element "holdingPoint" of complex type SegmentPointPropertyType
	HoldingPoint *SegmentPointPropertyType `xml:"holdingPoint,omitempty"`

	// generated from element "extent" of complex type CurvePropertyType
	Extent *CurvePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// HoldingPatternDistanceType Named source named complex type "HoldingPatternDistanceType"
type HoldingPatternDistanceType struct {
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

	// generated from group "aixm:HoldingPatternDistancePropertyGroup" with order 2896 depth 2
	Group_HoldingPatternDistancePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractHoldingPatternDistanceExtension" of complex type AbstractExtensionType
	AbstractHoldingPatternDistanceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractHoldingPatternDistanceExtension,omitempty"`
}

// Group_HoldingPatternDistancePropertyGroup UnNamed source named group "HoldingPatternDistancePropertyGroup"
type Group_HoldingPatternDistancePropertyGroup struct {

	// insertion point for fields

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// HoldingPatternDistancePropertyType Named source named complex type "HoldingPatternDistancePropertyType"
type HoldingPatternDistancePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "HoldingPatternDistance" of complex type HoldingPatternDistanceType
	HoldingPatternDistance *HoldingPatternDistanceType `xml:"HoldingPatternDistance,omitempty"`
}

// HoldingPatternDurationType Named source named complex type "HoldingPatternDurationType"
type HoldingPatternDurationType struct {
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

	// generated from group "aixm:HoldingPatternDurationPropertyGroup" with order 2913 depth 2
	Group_HoldingPatternDurationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractHoldingPatternDurationExtension" of complex type AbstractExtensionType
	AbstractHoldingPatternDurationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractHoldingPatternDurationExtension,omitempty"`
}

// Group_HoldingPatternDurationPropertyGroup UnNamed source named group "HoldingPatternDurationPropertyGroup"
type Group_HoldingPatternDurationPropertyGroup struct {

	// insertion point for fields

	// generated from element "duration" of complex type ValDurationType
	Duration *aixmdatatypes.ValDurationType `xml:"duration,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// HoldingPatternDurationPropertyType Named source named complex type "HoldingPatternDurationPropertyType"
type HoldingPatternDurationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "HoldingPatternDuration" of complex type HoldingPatternDurationType
	HoldingPatternDuration *HoldingPatternDurationType `xml:"HoldingPatternDuration,omitempty"`
}

// UnplannedHoldingPropertyType Named source named complex type "UnplannedHoldingPropertyType"
type UnplannedHoldingPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// UnplannedHoldingType Named source named complex type "UnplannedHoldingType"
type UnplannedHoldingType struct {
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

	// generated from element "timeSlice" of complex type UnplannedHoldingTimeSlicePropertyType
	TimeSlice []*UnplannedHoldingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// UnplannedHoldingTimeSlicePropertyType Named source named complex type "UnplannedHoldingTimeSlicePropertyType"
type UnplannedHoldingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "UnplannedHoldingTimeSlice" of complex type UnplannedHoldingTimeSliceType
	UnplannedHoldingTimeSlice *UnplannedHoldingTimeSliceType `xml:"UnplannedHoldingTimeSlice,omitempty"`
}

// UnplannedHoldingTimeSliceType Named source named complex type "UnplannedHoldingTimeSliceType"
type UnplannedHoldingTimeSliceType struct {
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

	// generated from group "aixm:UnplannedHoldingPropertyGroup" with order 2941 depth 2
	Group_UnplannedHoldingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUnplannedHoldingExtension" of complex type AbstractExtensionType
	AbstractUnplannedHoldingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUnplannedHoldingExtension,omitempty"`
}

// Group_UnplannedHoldingPropertyGroup UnNamed source named group "UnplannedHoldingPropertyGroup"
type Group_UnplannedHoldingPropertyGroup struct {

	// insertion point for fields

	// generated from element "unplannedHolding" of complex type CodeApprovalType
	UnplannedHolding *aixmdatatypes.CodeApprovalType `xml:"unplannedHolding,omitempty"`

	// generated from element "authorizedAltitude" of complex type ValDistanceVerticalType
	AuthorizedAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"authorizedAltitude,omitempty"`

	// generated from element "altitudeReference" of complex type CodeVerticalReferenceType
	AltitudeReference *aixmdatatypes.CodeVerticalReferenceType `xml:"altitudeReference,omitempty"`

	// generated from element "controlledAirspace" of complex type CodeYesNoType
	ControlledAirspace *aixmdatatypes.CodeYesNoType `xml:"controlledAirspace,omitempty"`

	// generated from element "holdingPoint" of complex type SegmentPointPropertyType
	HoldingPoint *SegmentPointPropertyType `xml:"holdingPoint,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

// SpecialNavigationSystemPropertyType Named source named complex type "SpecialNavigationSystemPropertyType"
type SpecialNavigationSystemPropertyType struct {
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

	// generated from group "aixm:AngleUsePropertyGroup" with order 3639 depth 2
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

	// generated from group "aixm:SegmentPointPropertyGroup" with order 3692 depth 2
	Group_SegmentPointPropertyGroup

	// generated from group "aixm:EnRouteSegmentPointPropertyGroup" with order 3693 depth 2
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

	// generated from group "aixm:PointReferencePropertyGroup" with order 3713 depth 2
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

// SegmentPointPropertyType Named source named complex type "SegmentPointPropertyType"
type SegmentPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AerialRefuellingPoint" of complex type AerialRefuellingPointType
	AerialRefuellingPoint *AerialRefuellingPointType `xml:"AerialRefuellingPoint,omitempty"`

	// generated from element "EnRouteSegmentPoint" of complex type EnRouteSegmentPointType
	EnRouteSegmentPoint *EnRouteSegmentPointType `xml:"EnRouteSegmentPoint,omitempty"`

	// generated from element "TerminalSegmentPoint" of complex type TerminalSegmentPointType
	TerminalSegmentPoint *TerminalSegmentPointType `xml:"TerminalSegmentPoint,omitempty"`
}

// TerminalSegmentPointType Named source named complex type "TerminalSegmentPointType"
type TerminalSegmentPointType struct {
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

	// generated from group "aixm:SegmentPointPropertyGroup" with order 3759 depth 2
	Group_SegmentPointPropertyGroup

	// generated from group "aixm:TerminalSegmentPointPropertyGroup" with order 3760 depth 2
	Group_TerminalSegmentPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSegmentPointExtension" of complex type AbstractExtensionType
	AbstractSegmentPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentPointExtension,omitempty"`

	// generated from element "AbstractTerminalSegmentPointExtension" of complex type AbstractExtensionType
	AbstractTerminalSegmentPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTerminalSegmentPointExtension,omitempty"`
}

// Group_TerminalSegmentPointPropertyGroup UnNamed source named group "TerminalSegmentPointPropertyGroup"
type Group_TerminalSegmentPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "role" of complex type CodeProcedureFixRoleType
	Role *aixmdatatypes.CodeProcedureFixRoleType `xml:"role,omitempty"`

	// generated from element "leadRadial" of complex type ValBearingType
	LeadRadial *aixmdatatypes.ValBearingType `xml:"leadRadial,omitempty"`

	// generated from element "leadDME" of complex type ValDistanceType
	LeadDME *aixmdatatypes.ValDistanceType `xml:"leadDME,omitempty"`

	// generated from element "indicatorFACF" of complex type CodeYesNoType
	IndicatorFACF *aixmdatatypes.CodeYesNoType `xml:"indicatorFACF,omitempty"`
}

// TerminalSegmentPointPropertyType Named source named complex type "TerminalSegmentPointPropertyType"
type TerminalSegmentPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "TerminalSegmentPoint" of complex type TerminalSegmentPointType
	TerminalSegmentPoint *TerminalSegmentPointType `xml:"TerminalSegmentPoint,omitempty"`
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

	// generated from group "aixm:LinguisticNotePropertyGroup" with order 3876 depth 2
	Group_LinguisticNotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
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

	// generated from group "aixm:NotePropertyGroup" with order 3892 depth 2
	Group_NotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
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

// ApproachAltitudeTableType Named source named complex type "ApproachAltitudeTableType"
type ApproachAltitudeTableType struct {
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

	// generated from group "aixm:ApproachAltitudeTablePropertyGroup" with order 4156 depth 2
	Group_ApproachAltitudeTablePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApproachAltitudeTableExtension" of complex type AbstractExtensionType
	AbstractApproachAltitudeTableExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachAltitudeTableExtension,omitempty"`
}

// Group_ApproachAltitudeTablePropertyGroup UnNamed source named group "ApproachAltitudeTablePropertyGroup"
type Group_ApproachAltitudeTablePropertyGroup struct {

	// insertion point for fields

	// generated from element "measurementPoint" of complex type CodeProcedureDistanceType
	MeasurementPoint *aixmdatatypes.CodeProcedureDistanceType `xml:"measurementPoint,omitempty"`

	// generated from element "altitude" of complex type ValDistanceVerticalType
	Altitude *aixmdatatypes.ValDistanceVerticalType `xml:"altitude,omitempty"`

	// generated from element "altitudeReference" of complex type CodeVerticalReferenceType
	AltitudeReference *aixmdatatypes.CodeVerticalReferenceType `xml:"altitudeReference,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ApproachAltitudeTablePropertyType Named source named complex type "ApproachAltitudeTablePropertyType"
type ApproachAltitudeTablePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApproachAltitudeTable" of complex type ApproachAltitudeTableType
	ApproachAltitudeTable *ApproachAltitudeTableType `xml:"ApproachAltitudeTable,omitempty"`
}

// ApproachConditionType Named source named complex type "ApproachConditionType"
type ApproachConditionType struct {
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

	// generated from group "aixm:ApproachConditionPropertyGroup" with order 4175 depth 2
	Group_ApproachConditionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_15 `xml:"extension,omitempty"`
}

// A_extension_15 Named source within outer element "extension"
type A_extension_15 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApproachConditionExtension" of complex type AbstractExtensionType
	AbstractApproachConditionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachConditionExtension,omitempty"`
}

// Group_ApproachConditionPropertyGroup UnNamed source named group "ApproachConditionPropertyGroup"
type Group_ApproachConditionPropertyGroup struct {

	// insertion point for fields

	// generated from element "finalApproachPath" of complex type CodeMinimaFinalApproachPathType
	FinalApproachPath *aixmdatatypes.CodeMinimaFinalApproachPathType `xml:"finalApproachPath,omitempty"`

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`

	// generated from element "climbGradient" of complex type ValSlopeType
	ClimbGradient *aixmdatatypes.ValSlopeType `xml:"climbGradient,omitempty"`

	// generated from element "minimumSet" of complex type MinimaPropertyType
	MinimumSet *MinimaPropertyType `xml:"minimumSet,omitempty"`

	// generated from element "circlingRestriction" of complex type CirclingRestrictionPropertyType
	CirclingRestriction []*CirclingRestrictionPropertyType `xml:"circlingRestriction,omitempty"`

	// generated from element "aircraftCategory" of complex type AircraftCharacteristicPropertyType
	AircraftCategory []*AircraftCharacteristicPropertyType `xml:"aircraftCategory,omitempty"`

	// generated from element "landingArea" of complex type LandingTakeoffAreaCollectionPropertyType
	LandingArea []*LandingTakeoffAreaCollectionPropertyType `xml:"landingArea,omitempty"`

	// generated from element "altimeter" of complex type AltimeterSourcePropertyType
	Altimeter *AltimeterSourcePropertyType `xml:"altimeter,omitempty"`

	// generated from element "designSurface" of complex type ObstacleAssessmentAreaPropertyType
	DesignSurface []*ObstacleAssessmentAreaPropertyType `xml:"designSurface,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ApproachConditionPropertyType Named source named complex type "ApproachConditionPropertyType"
type ApproachConditionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApproachCondition" of complex type ApproachConditionType
	ApproachCondition *ApproachConditionType `xml:"ApproachCondition,omitempty"`
}

// ApproachDistanceTableType Named source named complex type "ApproachDistanceTableType"
type ApproachDistanceTableType struct {
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

	// generated from group "aixm:ApproachDistanceTablePropertyGroup" with order 4200 depth 2
	Group_ApproachDistanceTablePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_16 `xml:"extension,omitempty"`
}

// A_extension_16 Named source within outer element "extension"
type A_extension_16 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApproachDistanceTableExtension" of complex type AbstractExtensionType
	AbstractApproachDistanceTableExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachDistanceTableExtension,omitempty"`
}

// Group_ApproachDistanceTablePropertyGroup UnNamed source named group "ApproachDistanceTablePropertyGroup"
type Group_ApproachDistanceTablePropertyGroup struct {

	// insertion point for fields

	// generated from element "startingMeasurementPoint" of complex type CodeProcedureDistanceType
	StartingMeasurementPoint *aixmdatatypes.CodeProcedureDistanceType `xml:"startingMeasurementPoint,omitempty"`

	// generated from element "valueHAT" of complex type ValDistanceVerticalType
	ValueHAT *aixmdatatypes.ValDistanceVerticalType `xml:"valueHAT,omitempty"`

	// generated from element "endingMeasurementPoint" of complex type CodeProcedureDistanceType
	EndingMeasurementPoint *aixmdatatypes.CodeProcedureDistanceType `xml:"endingMeasurementPoint,omitempty"`

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ApproachDistanceTablePropertyType Named source named complex type "ApproachDistanceTablePropertyType"
type ApproachDistanceTablePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApproachDistanceTable" of complex type ApproachDistanceTableType
	ApproachDistanceTable *ApproachDistanceTableType `xml:"ApproachDistanceTable,omitempty"`
}

// ApproachTimingTableType Named source named complex type "ApproachTimingTableType"
type ApproachTimingTableType struct {
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

	// generated from group "aixm:ApproachTimingTablePropertyGroup" with order 4220 depth 2
	Group_ApproachTimingTablePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_17 `xml:"extension,omitempty"`
}

// A_extension_17 Named source within outer element "extension"
type A_extension_17 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApproachTimingTableExtension" of complex type AbstractExtensionType
	AbstractApproachTimingTableExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachTimingTableExtension,omitempty"`
}

// Group_ApproachTimingTablePropertyGroup UnNamed source named group "ApproachTimingTablePropertyGroup"
type Group_ApproachTimingTablePropertyGroup struct {

	// insertion point for fields

	// generated from element "startingMeasurementPoint" of complex type CodeProcedureDistanceType
	StartingMeasurementPoint *aixmdatatypes.CodeProcedureDistanceType `xml:"startingMeasurementPoint,omitempty"`

	// generated from element "endingMeasurementPoint" of complex type CodeProcedureDistanceType
	EndingMeasurementPoint *aixmdatatypes.CodeProcedureDistanceType `xml:"endingMeasurementPoint,omitempty"`

	// generated from element "time" of complex type ValDurationType
	Time *aixmdatatypes.ValDurationType `xml:"time,omitempty"`

	// generated from element "speed" of complex type ValSpeedType
	Speed *aixmdatatypes.ValSpeedType `xml:"speed,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ApproachTimingTablePropertyType Named source named complex type "ApproachTimingTablePropertyType"
type ApproachTimingTablePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApproachTimingTable" of complex type ApproachTimingTableType
	ApproachTimingTable *ApproachTimingTableType `xml:"ApproachTimingTable,omitempty"`
}

// FinalProfileType Named source named complex type "FinalProfileType"
type FinalProfileType struct {
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

	// generated from group "aixm:FinalProfilePropertyGroup" with order 4240 depth 2
	Group_FinalProfilePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_18 `xml:"extension,omitempty"`
}

// A_extension_18 Named source within outer element "extension"
type A_extension_18 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFinalProfileExtension" of complex type AbstractExtensionType
	AbstractFinalProfileExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFinalProfileExtension,omitempty"`
}

// Group_FinalProfilePropertyGroup UnNamed source named group "FinalProfilePropertyGroup"
type Group_FinalProfilePropertyGroup struct {

	// insertion point for fields

	// generated from element "altitude" of complex type ApproachAltitudeTablePropertyType
	Altitude []*ApproachAltitudeTablePropertyType `xml:"altitude,omitempty"`

	// generated from element "distance" of complex type ApproachDistanceTablePropertyType
	Distance []*ApproachDistanceTablePropertyType `xml:"distance,omitempty"`

	// generated from element "timing" of complex type ApproachTimingTablePropertyType
	Timing []*ApproachTimingTablePropertyType `xml:"timing,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FinalProfilePropertyType Named source named complex type "FinalProfilePropertyType"
type FinalProfilePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FinalProfile" of complex type FinalProfileType
	FinalProfile *FinalProfileType `xml:"FinalProfile,omitempty"`
}

// InstrumentApproachProcedurePropertyType Named source named complex type "InstrumentApproachProcedurePropertyType"
type InstrumentApproachProcedurePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// InstrumentApproachProcedureType Named source named complex type "InstrumentApproachProcedureType"
type InstrumentApproachProcedureType struct {
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

	// generated from element "timeSlice" of complex type InstrumentApproachProcedureTimeSlicePropertyType
	TimeSlice []*InstrumentApproachProcedureTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// InstrumentApproachProcedureTimeSlicePropertyType Named source named complex type "InstrumentApproachProcedureTimeSlicePropertyType"
type InstrumentApproachProcedureTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "InstrumentApproachProcedureTimeSlice" of complex type InstrumentApproachProcedureTimeSliceType
	InstrumentApproachProcedureTimeSlice *InstrumentApproachProcedureTimeSliceType `xml:"InstrumentApproachProcedureTimeSlice,omitempty"`
}

// InstrumentApproachProcedureTimeSliceType Named source named complex type "InstrumentApproachProcedureTimeSliceType"
type InstrumentApproachProcedureTimeSliceType struct {
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

	// generated from element "communicationFailureInstruction" of complex type TextInstructionType
	CommunicationFailureInstruction *aixmdatatypes.TextInstructionType `xml:"communicationFailureInstruction,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "designCriteria" of complex type CodeDesignStandardType
	DesignCriteria *aixmdatatypes.CodeDesignStandardType `xml:"designCriteria,omitempty"`

	// generated from element "codingStandard" of complex type CodeProcedureCodingStandardType
	CodingStandard *aixmdatatypes.CodeProcedureCodingStandardType `xml:"codingStandard,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "RNAV" of complex type CodeYesNoType
	RNAV *aixmdatatypes.CodeYesNoType `xml:"RNAV,omitempty"`

	// generated from element "availability" of complex type ProcedureAvailabilityPropertyType
	Availability []*ProcedureAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "aircraftCharacteristic" of complex type AircraftCharacteristicPropertyType
	AircraftCharacteristic []*AircraftCharacteristicPropertyType `xml:"aircraftCharacteristic,omitempty"`

	// generated from element "flightTransition" of complex type ProcedureTransitionPropertyType
	FlightTransition []*ProcedureTransitionPropertyType `xml:"flightTransition,omitempty"`

	// generated from element "guidanceFacility_specialNavigationSystem" of complex type SpecialNavigationSystemPropertyType
	GuidanceFacility_specialNavigationSystem *SpecialNavigationSystemPropertyType `xml:"guidanceFacility_specialNavigationSystem,omitempty"`

	// generated from element "guidanceFacility_radar" of complex type RadarSystemPropertyType
	GuidanceFacility_radar *RadarSystemPropertyType `xml:"guidanceFacility_radar,omitempty"`

	// generated from element "guidanceFacility_navaid" of complex type NavaidPropertyType
	GuidanceFacility_navaid *NavaidPropertyType `xml:"guidanceFacility_navaid,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "safeAltitude" of complex type SafeAltitudeAreaPropertyType
	SafeAltitude *SafeAltitudeAreaPropertyType `xml:"safeAltitude,omitempty"`

	// generated from element "approachPrefix" of complex type CodeApproachPrefixType
	ApproachPrefix *aixmdatatypes.CodeApproachPrefixType `xml:"approachPrefix,omitempty"`

	// generated from element "approachType" of complex type CodeApproachType
	ApproachType *aixmdatatypes.CodeApproachType `xml:"approachType,omitempty"`

	// generated from element "multipleIdentification" of complex type CodeUpperAlphaType
	MultipleIdentification *aixmdatatypes.CodeUpperAlphaType `xml:"multipleIdentification,omitempty"`

	// generated from element "copterTrack" of complex type ValBearingType
	CopterTrack *aixmdatatypes.ValBearingType `xml:"copterTrack,omitempty"`

	// generated from element "circlingIdentification" of complex type CodeUpperAlphaType
	CirclingIdentification *aixmdatatypes.CodeUpperAlphaType `xml:"circlingIdentification,omitempty"`

	// generated from element "courseReversalInstruction" of complex type TextInstructionType
	CourseReversalInstruction *aixmdatatypes.TextInstructionType `xml:"courseReversalInstruction,omitempty"`

	// generated from element "additionalEquipment" of complex type CodeApproachEquipmentAdditionalType
	AdditionalEquipment *aixmdatatypes.CodeApproachEquipmentAdditionalType `xml:"additionalEquipment,omitempty"`

	// generated from element "channelGNSS" of complex type ValChannelNumberType
	ChannelGNSS *aixmdatatypes.ValChannelNumberType `xml:"channelGNSS,omitempty"`

	// generated from element "WAASReliable" of complex type CodeYesNoType
	WAASReliable *aixmdatatypes.CodeYesNoType `xml:"WAASReliable,omitempty"`

	// generated from element "landing" of complex type LandingTakeoffAreaCollectionPropertyType
	Landing *LandingTakeoffAreaCollectionPropertyType `xml:"landing,omitempty"`

	// generated from element "missedInstruction" of complex type MissedApproachGroupPropertyType
	MissedInstruction []*MissedApproachGroupPropertyType `xml:"missedInstruction,omitempty"`

	// generated from element "finalProfile" of complex type FinalProfilePropertyType
	FinalProfile *FinalProfilePropertyType `xml:"finalProfile,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_19 `xml:"extension,omitempty"`
}

// A_extension_19 Named source within outer element "extension"
type A_extension_19 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractInstrumentApproachProcedureExtension" of complex type AbstractExtensionType
	AbstractInstrumentApproachProcedureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractInstrumentApproachProcedureExtension,omitempty"`

	// generated from element "AbstractProcedureExtension" of complex type AbstractExtensionType
	AbstractProcedureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureExtension,omitempty"`
}

// Group_InstrumentApproachProcedurePropertyGroup UnNamed source named group "InstrumentApproachProcedurePropertyGroup"
type Group_InstrumentApproachProcedurePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ProcedurePropertyGroup" with order 4279 depth 2
	Group_ProcedurePropertyGroup

	// generated from element "approachPrefix" of complex type CodeApproachPrefixType
	ApproachPrefix *aixmdatatypes.CodeApproachPrefixType `xml:"approachPrefix,omitempty"`

	// generated from element "approachType" of complex type CodeApproachType
	ApproachType *aixmdatatypes.CodeApproachType `xml:"approachType,omitempty"`

	// generated from element "multipleIdentification" of complex type CodeUpperAlphaType
	MultipleIdentification *aixmdatatypes.CodeUpperAlphaType `xml:"multipleIdentification,omitempty"`

	// generated from element "copterTrack" of complex type ValBearingType
	CopterTrack *aixmdatatypes.ValBearingType `xml:"copterTrack,omitempty"`

	// generated from element "circlingIdentification" of complex type CodeUpperAlphaType
	CirclingIdentification *aixmdatatypes.CodeUpperAlphaType `xml:"circlingIdentification,omitempty"`

	// generated from element "courseReversalInstruction" of complex type TextInstructionType
	CourseReversalInstruction *aixmdatatypes.TextInstructionType `xml:"courseReversalInstruction,omitempty"`

	// generated from element "additionalEquipment" of complex type CodeApproachEquipmentAdditionalType
	AdditionalEquipment *aixmdatatypes.CodeApproachEquipmentAdditionalType `xml:"additionalEquipment,omitempty"`

	// generated from element "channelGNSS" of complex type ValChannelNumberType
	ChannelGNSS *aixmdatatypes.ValChannelNumberType `xml:"channelGNSS,omitempty"`

	// generated from element "WAASReliable" of complex type CodeYesNoType
	WAASReliable *aixmdatatypes.CodeYesNoType `xml:"WAASReliable,omitempty"`

	// generated from element "landing" of complex type LandingTakeoffAreaCollectionPropertyType
	Landing *LandingTakeoffAreaCollectionPropertyType `xml:"landing,omitempty"`

	// generated from element "missedInstruction" of complex type MissedApproachGroupPropertyType
	MissedInstruction []*MissedApproachGroupPropertyType `xml:"missedInstruction,omitempty"`

	// generated from element "finalProfile" of complex type FinalProfilePropertyType
	FinalProfile *FinalProfilePropertyType `xml:"finalProfile,omitempty"`
}

// MissedApproachGroupType Named source named complex type "MissedApproachGroupType"
type MissedApproachGroupType struct {
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

	// generated from group "aixm:MissedApproachGroupPropertyGroup" with order 4296 depth 2
	Group_MissedApproachGroupPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_20 `xml:"extension,omitempty"`
}

// A_extension_20 Named source within outer element "extension"
type A_extension_20 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMissedApproachGroupExtension" of complex type AbstractExtensionType
	AbstractMissedApproachGroupExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMissedApproachGroupExtension,omitempty"`
}

// Group_MissedApproachGroupPropertyGroup UnNamed source named group "MissedApproachGroupPropertyGroup"
type Group_MissedApproachGroupPropertyGroup struct {

	// insertion point for fields

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "alternateClimbInstruction" of complex type TextInstructionType
	AlternateClimbInstruction *aixmdatatypes.TextInstructionType `xml:"alternateClimbInstruction,omitempty"`

	// generated from element "alternateClimbAltitude" of complex type ValDistanceVerticalType
	AlternateClimbAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"alternateClimbAltitude,omitempty"`

	// generated from element "altimeter" of complex type AltimeterSourcePropertyType
	Altimeter []*AltimeterSourcePropertyType `xml:"altimeter,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MissedApproachGroupPropertyType Named source named complex type "MissedApproachGroupPropertyType"
type MissedApproachGroupPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "MissedApproachGroup" of complex type MissedApproachGroupType
	MissedApproachGroup *MissedApproachGroupType `xml:"MissedApproachGroup,omitempty"`
}

// TerminalArrivalAreaType Named source named complex type "TerminalArrivalAreaType"
type TerminalArrivalAreaType struct {
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

	// generated from element "timeSlice" of complex type TerminalArrivalAreaTimeSlicePropertyType
	TimeSlice []*TerminalArrivalAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TerminalArrivalAreaTimeSlicePropertyType Named source named complex type "TerminalArrivalAreaTimeSlicePropertyType"
type TerminalArrivalAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TerminalArrivalAreaTimeSlice" of complex type TerminalArrivalAreaTimeSliceType
	TerminalArrivalAreaTimeSlice *TerminalArrivalAreaTimeSliceType `xml:"TerminalArrivalAreaTimeSlice,omitempty"`
}

// TerminalArrivalAreaTimeSliceType Named source named complex type "TerminalArrivalAreaTimeSliceType"
type TerminalArrivalAreaTimeSliceType struct {
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

	// generated from group "aixm:TerminalArrivalAreaPropertyGroup" with order 4327 depth 2
	Group_TerminalArrivalAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_21 `xml:"extension,omitempty"`
}

// A_extension_21 Named source within outer element "extension"
type A_extension_21 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTerminalArrivalAreaExtension" of complex type AbstractExtensionType
	AbstractTerminalArrivalAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTerminalArrivalAreaExtension,omitempty"`
}

// Group_TerminalArrivalAreaPropertyGroup UnNamed source named group "TerminalArrivalAreaPropertyGroup"
type Group_TerminalArrivalAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "arrivalAreaType" of complex type CodeTAAType
	ArrivalAreaType *aixmdatatypes.CodeTAAType `xml:"arrivalAreaType,omitempty"`

	// generated from element "outerBufferWidth" of complex type ValDistanceType
	OuterBufferWidth *aixmdatatypes.ValDistanceType `xml:"outerBufferWidth,omitempty"`

	// generated from element "lateralBufferWidth" of complex type ValDistanceType
	LateralBufferWidth *aixmdatatypes.ValDistanceType `xml:"lateralBufferWidth,omitempty"`

	// generated from element "IF_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	IF_fixDesignatedPoint *DesignatedPointPropertyType `xml:"IF_fixDesignatedPoint,omitempty"`

	// generated from element "IF_navaidSystem" of complex type NavaidPropertyType
	IF_navaidSystem *NavaidPropertyType `xml:"IF_navaidSystem,omitempty"`

	// generated from element "IF_position" of complex type PointPropertyType
	IF_position *PointPropertyType `xml:"IF_position,omitempty"`

	// generated from element "IF_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	IF_runwayPoint *RunwayCentrelinePointPropertyType `xml:"IF_runwayPoint,omitempty"`

	// generated from element "IF_aimingPoint" of complex type TouchDownLiftOffPropertyType
	IF_aimingPoint *TouchDownLiftOffPropertyType `xml:"IF_aimingPoint,omitempty"`

	// generated from element "IF_airportReferencePoint" of complex type AirportHeliportPropertyType
	IF_airportReferencePoint *AirportHeliportPropertyType `xml:"IF_airportReferencePoint,omitempty"`

	// generated from element "IAF_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	IAF_fixDesignatedPoint *DesignatedPointPropertyType `xml:"IAF_fixDesignatedPoint,omitempty"`

	// generated from element "IAF_navaidSystem" of complex type NavaidPropertyType
	IAF_navaidSystem *NavaidPropertyType `xml:"IAF_navaidSystem,omitempty"`

	// generated from element "IAF_position" of complex type PointPropertyType
	IAF_position *PointPropertyType `xml:"IAF_position,omitempty"`

	// generated from element "IAF_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	IAF_runwayPoint *RunwayCentrelinePointPropertyType `xml:"IAF_runwayPoint,omitempty"`

	// generated from element "IAF_aimingPoint" of complex type TouchDownLiftOffPropertyType
	IAF_aimingPoint *TouchDownLiftOffPropertyType `xml:"IAF_aimingPoint,omitempty"`

	// generated from element "IAF_airportReferencePoint" of complex type AirportHeliportPropertyType
	IAF_airportReferencePoint *AirportHeliportPropertyType `xml:"IAF_airportReferencePoint,omitempty"`

	// generated from element "buffer" of complex type SurfacePropertyType
	Buffer *SurfacePropertyType `xml:"buffer,omitempty"`

	// generated from element "sector" of complex type TerminalArrivalAreaSectorPropertyType
	Sector []*TerminalArrivalAreaSectorPropertyType `xml:"sector,omitempty"`

	// generated from element "approachRNAV" of complex type InstrumentApproachProcedurePropertyType
	ApproachRNAV *InstrumentApproachProcedurePropertyType `xml:"approachRNAV,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TerminalArrivalAreaSectorType Named source named complex type "TerminalArrivalAreaSectorType"
type TerminalArrivalAreaSectorType struct {
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

	// generated from group "aixm:TerminalArrivalAreaSectorPropertyGroup" with order 4360 depth 2
	Group_TerminalArrivalAreaSectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_22 `xml:"extension,omitempty"`
}

// A_extension_22 Named source within outer element "extension"
type A_extension_22 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTerminalArrivalAreaSectorExtension" of complex type AbstractExtensionType
	AbstractTerminalArrivalAreaSectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTerminalArrivalAreaSectorExtension,omitempty"`
}

// Group_TerminalArrivalAreaSectorPropertyGroup UnNamed source named group "TerminalArrivalAreaSectorPropertyGroup"
type Group_TerminalArrivalAreaSectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "flyByCode" of complex type CodeYesNoType
	FlyByCode *aixmdatatypes.CodeYesNoType `xml:"flyByCode,omitempty"`

	// generated from element "procedureTurnRequired" of complex type CodeYesNoType
	ProcedureTurnRequired *aixmdatatypes.CodeYesNoType `xml:"procedureTurnRequired,omitempty"`

	// generated from element "altitudeDescription" of complex type CodeAltitudeUseType
	AltitudeDescription *aixmdatatypes.CodeAltitudeUseType `xml:"altitudeDescription,omitempty"`

	// generated from element "sectorDefinition" of complex type CircleSectorPropertyType
	SectorDefinition *CircleSectorPropertyType `xml:"sectorDefinition,omitempty"`

	// generated from element "extent" of complex type SurfacePropertyType
	Extent *SurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "significantObstacle" of complex type ObstructionPropertyType
	SignificantObstacle []*ObstructionPropertyType `xml:"significantObstacle,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TerminalArrivalAreaSectorPropertyType Named source named complex type "TerminalArrivalAreaSectorPropertyType"
type TerminalArrivalAreaSectorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "TerminalArrivalAreaSector" of complex type TerminalArrivalAreaSectorType
	TerminalArrivalAreaSector *TerminalArrivalAreaSectorType `xml:"TerminalArrivalAreaSector,omitempty"`
}

// FASDataBlockType Named source named complex type "FASDataBlockType"
type FASDataBlockType struct {
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

	// generated from group "aixm:FASDataBlockPropertyGroup" with order 4382 depth 2
	Group_FASDataBlockPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_23 `xml:"extension,omitempty"`
}

// A_extension_23 Named source within outer element "extension"
type A_extension_23 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFASDataBlockExtension" of complex type AbstractExtensionType
	AbstractFASDataBlockExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFASDataBlockExtension,omitempty"`
}

// Group_FASDataBlockPropertyGroup UnNamed source named group "FASDataBlockPropertyGroup"
type Group_FASDataBlockPropertyGroup struct {

	// insertion point for fields

	// generated from element "horizontalAlarmLimit" of complex type ValAlarmLimitType
	HorizontalAlarmLimit *aixmdatatypes.ValAlarmLimitType `xml:"horizontalAlarmLimit,omitempty"`

	// generated from element "verticalAlarmLimit" of complex type ValAlarmLimitType
	VerticalAlarmLimit *aixmdatatypes.ValAlarmLimitType `xml:"verticalAlarmLimit,omitempty"`

	// generated from element "thresholdCourseWidth" of complex type ValDistanceType
	ThresholdCourseWidth *aixmdatatypes.ValDistanceType `xml:"thresholdCourseWidth,omitempty"`

	// generated from element "lengthOffset" of complex type ValDistanceType
	LengthOffset *aixmdatatypes.ValDistanceType `xml:"lengthOffset,omitempty"`

	// generated from element "CRCRemainder" of complex type ValHexType
	CRCRemainder *aixmdatatypes.ValHexType `xml:"CRCRemainder,omitempty"`

	// generated from element "operationType" of complex type NoSequenceType
	OperationType *aixmdatatypes.NoSequenceType `xml:"operationType,omitempty"`

	// generated from element "serviceProviderSBAS" of complex type NoSequenceType
	ServiceProviderSBAS *aixmdatatypes.NoSequenceType `xml:"serviceProviderSBAS,omitempty"`

	// generated from element "approachPerformanceDesignator" of complex type NoSequenceType
	ApproachPerformanceDesignator *aixmdatatypes.NoSequenceType `xml:"approachPerformanceDesignator,omitempty"`

	// generated from element "routeIndicator" of complex type CodeRouteIndicatorType
	RouteIndicator *aixmdatatypes.CodeRouteIndicatorType `xml:"routeIndicator,omitempty"`

	// generated from element "referencePathDataSelector" of complex type NoSequenceType
	ReferencePathDataSelector *aixmdatatypes.NoSequenceType `xml:"referencePathDataSelector,omitempty"`

	// generated from element "referencePathIdentifier" of complex type CodeReferencePathIdentifierType
	ReferencePathIdentifier *aixmdatatypes.CodeReferencePathIdentifierType `xml:"referencePathIdentifier,omitempty"`

	// generated from element "codeICAO" of complex type CodeICAOCountryType
	CodeICAO *aixmdatatypes.CodeICAOCountryType `xml:"codeICAO,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FASDataBlockPropertyType Named source named complex type "FASDataBlockPropertyType"
type FASDataBlockPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "FASDataBlock" of complex type FASDataBlockType
	FASDataBlock *FASDataBlockType `xml:"FASDataBlock,omitempty"`
}

// CirclingAreaType Named source named complex type "CirclingAreaType"
type CirclingAreaType struct {
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

	// generated from element "timeSlice" of complex type CirclingAreaTimeSlicePropertyType
	TimeSlice []*CirclingAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// CirclingAreaTimeSlicePropertyType Named source named complex type "CirclingAreaTimeSlicePropertyType"
type CirclingAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "CirclingAreaTimeSlice" of complex type CirclingAreaTimeSliceType
	CirclingAreaTimeSlice *CirclingAreaTimeSliceType `xml:"CirclingAreaTimeSlice,omitempty"`
}

// CirclingAreaTimeSliceType Named source named complex type "CirclingAreaTimeSliceType"
type CirclingAreaTimeSliceType struct {
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

	// generated from group "aixm:CirclingAreaPropertyGroup" with order 4421 depth 2
	Group_CirclingAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_24 `xml:"extension,omitempty"`
}

// A_extension_24 Named source within outer element "extension"
type A_extension_24 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCirclingAreaExtension" of complex type AbstractExtensionType
	AbstractCirclingAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCirclingAreaExtension,omitempty"`
}

// Group_CirclingAreaPropertyGroup UnNamed source named group "CirclingAreaPropertyGroup"
type Group_CirclingAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "extent" of complex type SurfacePropertyType
	Extent *SurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "approach" of complex type InstrumentApproachProcedurePropertyType
	Approach *InstrumentApproachProcedurePropertyType `xml:"approach,omitempty"`

	// generated from element "condition" of complex type ApproachConditionPropertyType
	Condition []*ApproachConditionPropertyType `xml:"condition,omitempty"`

	// generated from element "aircraftCategory" of complex type AircraftCharacteristicPropertyType
	AircraftCategory *AircraftCharacteristicPropertyType `xml:"aircraftCategory,omitempty"`

	// generated from element "designSurface" of complex type ObstacleAssessmentAreaPropertyType
	DesignSurface []*ObstacleAssessmentAreaPropertyType `xml:"designSurface,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// CirclingRestrictionType Named source named complex type "CirclingRestrictionType"
type CirclingRestrictionType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 4439 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:CirclingRestrictionPropertyGroup" with order 4440 depth 2
	Group_CirclingRestrictionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_25 `xml:"extension,omitempty"`
}

// A_extension_25 Named source within outer element "extension"
type A_extension_25 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractCirclingRestrictionExtension" of complex type AbstractExtensionType
	AbstractCirclingRestrictionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCirclingRestrictionExtension,omitempty"`
}

// Group_CirclingRestrictionPropertyGroup UnNamed source named group "CirclingRestrictionPropertyGroup"
type Group_CirclingRestrictionPropertyGroup struct {

	// insertion point for fields

	// generated from element "sectorDescription" of complex type CircleSectorPropertyType
	SectorDescription *CircleSectorPropertyType `xml:"sectorDescription,omitempty"`

	// generated from element "restrictionArea" of complex type SurfacePropertyType
	RestrictionArea *SurfacePropertyType `xml:"restrictionArea,omitempty"`
}

// CirclingRestrictionPropertyType Named source named complex type "CirclingRestrictionPropertyType"
type CirclingRestrictionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "CirclingRestriction" of complex type CirclingRestrictionType
	CirclingRestriction *CirclingRestrictionType `xml:"CirclingRestriction,omitempty"`
}

// EquipmentUnavailableAdjustmentType Named source named complex type "EquipmentUnavailableAdjustmentType"
type EquipmentUnavailableAdjustmentType struct {
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

	// generated from group "aixm:EquipmentUnavailableAdjustmentPropertyGroup" with order 4458 depth 2
	Group_EquipmentUnavailableAdjustmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_26 `xml:"extension,omitempty"`
}

// A_extension_26 Named source within outer element "extension"
type A_extension_26 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractEquipmentUnavailableAdjustmentExtension" of complex type AbstractExtensionType
	AbstractEquipmentUnavailableAdjustmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractEquipmentUnavailableAdjustmentExtension,omitempty"`
}

// Group_EquipmentUnavailableAdjustmentPropertyGroup UnNamed source named group "EquipmentUnavailableAdjustmentPropertyGroup"
type Group_EquipmentUnavailableAdjustmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeEquipmentUnavailableType
	Type_XSD *aixmdatatypes.CodeEquipmentUnavailableType `xml:"type,omitempty"`

	// generated from element "approachLightingInoperative" of complex type CodeYesNoType
	ApproachLightingInoperative *aixmdatatypes.CodeYesNoType `xml:"approachLightingInoperative,omitempty"`

	// generated from element "adjustmentINOPCol" of complex type EquipmentUnavailableAdjustmentColumnPropertyType
	AdjustmentINOPCol []*EquipmentUnavailableAdjustmentColumnPropertyType `xml:"adjustmentINOPCol,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// EquipmentUnavailableAdjustmentPropertyType Named source named complex type "EquipmentUnavailableAdjustmentPropertyType"
type EquipmentUnavailableAdjustmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "EquipmentUnavailableAdjustment" of complex type EquipmentUnavailableAdjustmentType
	EquipmentUnavailableAdjustment *EquipmentUnavailableAdjustmentType `xml:"EquipmentUnavailableAdjustment,omitempty"`
}

// EquipmentUnavailableAdjustmentColumnType Named source named complex type "EquipmentUnavailableAdjustmentColumnType"
type EquipmentUnavailableAdjustmentColumnType struct {
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

	// generated from group "aixm:EquipmentUnavailableAdjustmentColumnPropertyGroup" with order 4477 depth 2
	Group_EquipmentUnavailableAdjustmentColumnPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_27 `xml:"extension,omitempty"`
}

// A_extension_27 Named source within outer element "extension"
type A_extension_27 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractEquipmentUnavailableAdjustmentColumnExtension" of complex type AbstractExtensionType
	AbstractEquipmentUnavailableAdjustmentColumnExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractEquipmentUnavailableAdjustmentColumnExtension,omitempty"`
}

// Group_EquipmentUnavailableAdjustmentColumnPropertyGroup UnNamed source named group "EquipmentUnavailableAdjustmentColumnPropertyGroup"
type Group_EquipmentUnavailableAdjustmentColumnPropertyGroup struct {

	// insertion point for fields

	// generated from element "guidanceEquipment" of complex type CodeApproachType
	GuidanceEquipment *aixmdatatypes.CodeApproachType `xml:"guidanceEquipment,omitempty"`

	// generated from element "landingSystemLights" of complex type CodeYesNoType
	LandingSystemLights *aixmdatatypes.CodeYesNoType `xml:"landingSystemLights,omitempty"`

	// generated from element "equipmentRVR" of complex type CodeYesNoType
	EquipmentRVR *aixmdatatypes.CodeYesNoType `xml:"equipmentRVR,omitempty"`

	// generated from element "visibilityAdjustment" of complex type ValDistanceVerticalType
	VisibilityAdjustment *aixmdatatypes.ValDistanceVerticalType `xml:"visibilityAdjustment,omitempty"`

	// generated from element "approachLightingInoperative" of complex type CodeYesNoType
	ApproachLightingInoperative *aixmdatatypes.CodeYesNoType `xml:"approachLightingInoperative,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// EquipmentUnavailableAdjustmentColumnPropertyType Named source named complex type "EquipmentUnavailableAdjustmentColumnPropertyType"
type EquipmentUnavailableAdjustmentColumnPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "EquipmentUnavailableAdjustmentColumn" of complex type EquipmentUnavailableAdjustmentColumnType
	EquipmentUnavailableAdjustmentColumn *EquipmentUnavailableAdjustmentColumnType `xml:"EquipmentUnavailableAdjustmentColumn,omitempty"`
}

// MinimaType Named source named complex type "MinimaType"
type MinimaType struct {
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

	// generated from group "aixm:MinimaPropertyGroup" with order 4498 depth 2
	Group_MinimaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_28 `xml:"extension,omitempty"`
}

// A_extension_28 Named source within outer element "extension"
type A_extension_28 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMinimaExtension" of complex type AbstractExtensionType
	AbstractMinimaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMinimaExtension,omitempty"`
}

// Group_MinimaPropertyGroup UnNamed source named group "MinimaPropertyGroup"
type Group_MinimaPropertyGroup struct {

	// insertion point for fields

	// generated from element "altitude" of complex type ValDistanceVerticalType
	Altitude *aixmdatatypes.ValDistanceVerticalType `xml:"altitude,omitempty"`

	// generated from element "altitudeCode" of complex type CodeMinimumAltitudeType
	AltitudeCode *aixmdatatypes.CodeMinimumAltitudeType `xml:"altitudeCode,omitempty"`

	// generated from element "altitudeReference" of complex type CodeVerticalReferenceType
	AltitudeReference *aixmdatatypes.CodeVerticalReferenceType `xml:"altitudeReference,omitempty"`

	// generated from element "height" of complex type ValDistanceVerticalType
	Height *aixmdatatypes.ValDistanceVerticalType `xml:"height,omitempty"`

	// generated from element "militaryHeight" of complex type ValDistanceVerticalType
	MilitaryHeight *aixmdatatypes.ValDistanceVerticalType `xml:"militaryHeight,omitempty"`

	// generated from element "radioHeight" of complex type ValDistanceVerticalType
	RadioHeight *aixmdatatypes.ValDistanceVerticalType `xml:"radioHeight,omitempty"`

	// generated from element "heightCode" of complex type CodeMinimumHeightType
	HeightCode *aixmdatatypes.CodeMinimumHeightType `xml:"heightCode,omitempty"`

	// generated from element "heightReference" of complex type CodeHeightReferenceType
	HeightReference *aixmdatatypes.CodeHeightReferenceType `xml:"heightReference,omitempty"`

	// generated from element "visibility" of complex type ValDistanceType
	Visibility *aixmdatatypes.ValDistanceType `xml:"visibility,omitempty"`

	// generated from element "militaryVisibility" of complex type ValDistanceType
	MilitaryVisibility *aixmdatatypes.ValDistanceType `xml:"militaryVisibility,omitempty"`

	// generated from element "mandatoryRVR" of complex type CodeYesNoType
	MandatoryRVR *aixmdatatypes.CodeYesNoType `xml:"mandatoryRVR,omitempty"`

	// generated from element "remoteAltimeterMinima" of complex type CodeYesNoType
	RemoteAltimeterMinima *aixmdatatypes.CodeYesNoType `xml:"remoteAltimeterMinima,omitempty"`

	// generated from element "adjustmentINOP" of complex type EquipmentUnavailableAdjustmentPropertyType
	AdjustmentINOP []*EquipmentUnavailableAdjustmentPropertyType `xml:"adjustmentINOP,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MinimaPropertyType Named source named complex type "MinimaPropertyType"
type MinimaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Minima" of complex type MinimaType
	Minima *MinimaType `xml:"Minima,omitempty"`
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

// StandardInstrumentArrivalType Named source named complex type "StandardInstrumentArrivalType"
type StandardInstrumentArrivalType struct {
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

	// generated from element "timeSlice" of complex type StandardInstrumentArrivalTimeSlicePropertyType
	TimeSlice []*StandardInstrumentArrivalTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandardInstrumentArrivalTimeSlicePropertyType Named source named complex type "StandardInstrumentArrivalTimeSlicePropertyType"
type StandardInstrumentArrivalTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandardInstrumentArrivalTimeSlice" of complex type StandardInstrumentArrivalTimeSliceType
	StandardInstrumentArrivalTimeSlice *StandardInstrumentArrivalTimeSliceType `xml:"StandardInstrumentArrivalTimeSlice,omitempty"`
}

// StandardInstrumentArrivalTimeSliceType Named source named complex type "StandardInstrumentArrivalTimeSliceType"
type StandardInstrumentArrivalTimeSliceType struct {
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

	// generated from element "communicationFailureInstruction" of complex type TextInstructionType
	CommunicationFailureInstruction *aixmdatatypes.TextInstructionType `xml:"communicationFailureInstruction,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "designCriteria" of complex type CodeDesignStandardType
	DesignCriteria *aixmdatatypes.CodeDesignStandardType `xml:"designCriteria,omitempty"`

	// generated from element "codingStandard" of complex type CodeProcedureCodingStandardType
	CodingStandard *aixmdatatypes.CodeProcedureCodingStandardType `xml:"codingStandard,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "RNAV" of complex type CodeYesNoType
	RNAV *aixmdatatypes.CodeYesNoType `xml:"RNAV,omitempty"`

	// generated from element "availability" of complex type ProcedureAvailabilityPropertyType
	Availability []*ProcedureAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "aircraftCharacteristic" of complex type AircraftCharacteristicPropertyType
	AircraftCharacteristic []*AircraftCharacteristicPropertyType `xml:"aircraftCharacteristic,omitempty"`

	// generated from element "flightTransition" of complex type ProcedureTransitionPropertyType
	FlightTransition []*ProcedureTransitionPropertyType `xml:"flightTransition,omitempty"`

	// generated from element "guidanceFacility_specialNavigationSystem" of complex type SpecialNavigationSystemPropertyType
	GuidanceFacility_specialNavigationSystem *SpecialNavigationSystemPropertyType `xml:"guidanceFacility_specialNavigationSystem,omitempty"`

	// generated from element "guidanceFacility_radar" of complex type RadarSystemPropertyType
	GuidanceFacility_radar *RadarSystemPropertyType `xml:"guidanceFacility_radar,omitempty"`

	// generated from element "guidanceFacility_navaid" of complex type NavaidPropertyType
	GuidanceFacility_navaid *NavaidPropertyType `xml:"guidanceFacility_navaid,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "safeAltitude" of complex type SafeAltitudeAreaPropertyType
	SafeAltitude *SafeAltitudeAreaPropertyType `xml:"safeAltitude,omitempty"`

	// generated from element "designator" of complex type TextSIDSTARDesignatorType
	Designator *aixmdatatypes.TextSIDSTARDesignatorType `xml:"designator,omitempty"`

	// generated from element "arrival" of complex type LandingTakeoffAreaCollectionPropertyType
	Arrival *LandingTakeoffAreaCollectionPropertyType `xml:"arrival,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_29 `xml:"extension,omitempty"`
}

// A_extension_29 Named source within outer element "extension"
type A_extension_29 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardInstrumentArrivalExtension" of complex type AbstractExtensionType
	AbstractStandardInstrumentArrivalExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardInstrumentArrivalExtension,omitempty"`

	// generated from element "AbstractProcedureExtension" of complex type AbstractExtensionType
	AbstractProcedureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureExtension,omitempty"`
}

// Group_StandardInstrumentArrivalPropertyGroup UnNamed source named group "StandardInstrumentArrivalPropertyGroup"
type Group_StandardInstrumentArrivalPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ProcedurePropertyGroup" with order 4547 depth 2
	Group_ProcedurePropertyGroup

	// generated from element "designator" of complex type TextSIDSTARDesignatorType
	Designator *aixmdatatypes.TextSIDSTARDesignatorType `xml:"designator,omitempty"`

	// generated from element "arrival" of complex type LandingTakeoffAreaCollectionPropertyType
	Arrival *LandingTakeoffAreaCollectionPropertyType `xml:"arrival,omitempty"`
}

// DepartureArrivalConditionType Named source named complex type "DepartureArrivalConditionType"
type DepartureArrivalConditionType struct {
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

	// generated from group "aixm:DepartureArrivalConditionPropertyGroup" with order 4554 depth 2
	Group_DepartureArrivalConditionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_30 `xml:"extension,omitempty"`
}

// A_extension_30 Named source within outer element "extension"
type A_extension_30 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDepartureArrivalConditionExtension" of complex type AbstractExtensionType
	AbstractDepartureArrivalConditionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDepartureArrivalConditionExtension,omitempty"`
}

// Group_DepartureArrivalConditionPropertyGroup UnNamed source named group "DepartureArrivalConditionPropertyGroup"
type Group_DepartureArrivalConditionPropertyGroup struct {

	// insertion point for fields

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

	// generated from element "engineType" of complex type AircraftCharacteristicPropertyType
	EngineType *AircraftCharacteristicPropertyType `xml:"engineType,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// DepartureArrivalConditionPropertyType Named source named complex type "DepartureArrivalConditionPropertyType"
type DepartureArrivalConditionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "DepartureArrivalCondition" of complex type DepartureArrivalConditionType
	DepartureArrivalCondition *DepartureArrivalConditionType `xml:"DepartureArrivalCondition,omitempty"`
}

// NavigationAreaType Named source named complex type "NavigationAreaType"
type NavigationAreaType struct {
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

	// generated from element "timeSlice" of complex type NavigationAreaTimeSlicePropertyType
	TimeSlice []*NavigationAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// NavigationAreaTimeSlicePropertyType Named source named complex type "NavigationAreaTimeSlicePropertyType"
type NavigationAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "NavigationAreaTimeSlice" of complex type NavigationAreaTimeSliceType
	NavigationAreaTimeSlice *NavigationAreaTimeSliceType `xml:"NavigationAreaTimeSlice,omitempty"`
}

// NavigationAreaTimeSliceType Named source named complex type "NavigationAreaTimeSliceType"
type NavigationAreaTimeSliceType struct {
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

	// generated from group "aixm:NavigationAreaPropertyGroup" with order 4587 depth 2
	Group_NavigationAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_31 `xml:"extension,omitempty"`
}

// A_extension_31 Named source within outer element "extension"
type A_extension_31 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavigationAreaExtension" of complex type AbstractExtensionType
	AbstractNavigationAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavigationAreaExtension,omitempty"`
}

// Group_NavigationAreaPropertyGroup UnNamed source named group "NavigationAreaPropertyGroup"
type Group_NavigationAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "navigationAreaType" of complex type CodeNavigationAreaType
	NavigationAreaType *aixmdatatypes.CodeNavigationAreaType `xml:"navigationAreaType,omitempty"`

	// generated from element "minimumCeiling" of complex type ValDistanceVerticalType
	MinimumCeiling *aixmdatatypes.ValDistanceVerticalType `xml:"minimumCeiling,omitempty"`

	// generated from element "minimumVisibility" of complex type ValDistanceType
	MinimumVisibility *aixmdatatypes.ValDistanceType `xml:"minimumVisibility,omitempty"`

	// generated from element "departure" of complex type StandardInstrumentDeparturePropertyType
	Departure *StandardInstrumentDeparturePropertyType `xml:"departure,omitempty"`

	// generated from element "sector" of complex type NavigationAreaSectorPropertyType
	Sector []*NavigationAreaSectorPropertyType `xml:"sector,omitempty"`

	// generated from element "centrePoint_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	CentrePoint_fixDesignatedPoint *DesignatedPointPropertyType `xml:"centrePoint_fixDesignatedPoint,omitempty"`

	// generated from element "centrePoint_navaidSystem" of complex type NavaidPropertyType
	CentrePoint_navaidSystem *NavaidPropertyType `xml:"centrePoint_navaidSystem,omitempty"`

	// generated from element "centrePoint_position" of complex type PointPropertyType
	CentrePoint_position *PointPropertyType `xml:"centrePoint_position,omitempty"`

	// generated from element "centrePoint_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	CentrePoint_runwayPoint *RunwayCentrelinePointPropertyType `xml:"centrePoint_runwayPoint,omitempty"`

	// generated from element "centrePoint_aimingPoint" of complex type TouchDownLiftOffPropertyType
	CentrePoint_aimingPoint *TouchDownLiftOffPropertyType `xml:"centrePoint_aimingPoint,omitempty"`

	// generated from element "centrePoint_airportReferencePoint" of complex type AirportHeliportPropertyType
	CentrePoint_airportReferencePoint *AirportHeliportPropertyType `xml:"centrePoint_airportReferencePoint,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NavigationAreaSectorType Named source named complex type "NavigationAreaSectorType"
type NavigationAreaSectorType struct {
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

	// generated from group "aixm:NavigationAreaSectorPropertyGroup" with order 4612 depth 2
	Group_NavigationAreaSectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_32 `xml:"extension,omitempty"`
}

// A_extension_32 Named source within outer element "extension"
type A_extension_32 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavigationAreaSectorExtension" of complex type AbstractExtensionType
	AbstractNavigationAreaSectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavigationAreaSectorExtension,omitempty"`
}

// Group_NavigationAreaSectorPropertyGroup UnNamed source named group "NavigationAreaSectorPropertyGroup"
type Group_NavigationAreaSectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "sectorDefinition" of complex type CircleSectorPropertyType
	SectorDefinition *CircleSectorPropertyType `xml:"sectorDefinition,omitempty"`

	// generated from element "significantObstacle" of complex type ObstructionPropertyType
	SignificantObstacle []*ObstructionPropertyType `xml:"significantObstacle,omitempty"`

	// generated from element "extent" of complex type SurfacePropertyType
	Extent *SurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "sectorCriteria" of complex type SectorDesignPropertyType
	SectorCriteria []*SectorDesignPropertyType `xml:"sectorCriteria,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NavigationAreaSectorPropertyType Named source named complex type "NavigationAreaSectorPropertyType"
type NavigationAreaSectorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "NavigationAreaSector" of complex type NavigationAreaSectorType
	NavigationAreaSector *NavigationAreaSectorType `xml:"NavigationAreaSector,omitempty"`
}

// SectorDesignType Named source named complex type "SectorDesignType"
type SectorDesignType struct {
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

	// generated from group "aixm:SectorDesignPropertyGroup" with order 4632 depth 2
	Group_SectorDesignPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_33 `xml:"extension,omitempty"`
}

// A_extension_33 Named source within outer element "extension"
type A_extension_33 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSectorDesignExtension" of complex type AbstractExtensionType
	AbstractSectorDesignExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSectorDesignExtension,omitempty"`
}

// Group_SectorDesignPropertyGroup UnNamed source named group "SectorDesignPropertyGroup"
type Group_SectorDesignPropertyGroup struct {

	// insertion point for fields

	// generated from element "turnDirection" of complex type CodeDirectionTurnType
	TurnDirection *aixmdatatypes.CodeDirectionTurnType `xml:"turnDirection,omitempty"`

	// generated from element "designGradient" of complex type ValSlopeType
	DesignGradient *aixmdatatypes.ValSlopeType `xml:"designGradient,omitempty"`

	// generated from element "terminationAltitude" of complex type ValDistanceVerticalType
	TerminationAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"terminationAltitude,omitempty"`

	// generated from element "turnPermitted" of complex type CodeYesNoType
	TurnPermitted *aixmdatatypes.CodeYesNoType `xml:"turnPermitted,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SectorDesignPropertyType Named source named complex type "SectorDesignPropertyType"
type SectorDesignPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SectorDesign" of complex type SectorDesignType
	SectorDesign *SectorDesignType `xml:"SectorDesign,omitempty"`
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

// StandardInstrumentDepartureType Named source named complex type "StandardInstrumentDepartureType"
type StandardInstrumentDepartureType struct {
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

	// generated from element "timeSlice" of complex type StandardInstrumentDepartureTimeSlicePropertyType
	TimeSlice []*StandardInstrumentDepartureTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandardInstrumentDepartureTimeSlicePropertyType Named source named complex type "StandardInstrumentDepartureTimeSlicePropertyType"
type StandardInstrumentDepartureTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandardInstrumentDepartureTimeSlice" of complex type StandardInstrumentDepartureTimeSliceType
	StandardInstrumentDepartureTimeSlice *StandardInstrumentDepartureTimeSliceType `xml:"StandardInstrumentDepartureTimeSlice,omitempty"`
}

// StandardInstrumentDepartureTimeSliceType Named source named complex type "StandardInstrumentDepartureTimeSliceType"
type StandardInstrumentDepartureTimeSliceType struct {
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

	// generated from element "communicationFailureInstruction" of complex type TextInstructionType
	CommunicationFailureInstruction *aixmdatatypes.TextInstructionType `xml:"communicationFailureInstruction,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "designCriteria" of complex type CodeDesignStandardType
	DesignCriteria *aixmdatatypes.CodeDesignStandardType `xml:"designCriteria,omitempty"`

	// generated from element "codingStandard" of complex type CodeProcedureCodingStandardType
	CodingStandard *aixmdatatypes.CodeProcedureCodingStandardType `xml:"codingStandard,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "RNAV" of complex type CodeYesNoType
	RNAV *aixmdatatypes.CodeYesNoType `xml:"RNAV,omitempty"`

	// generated from element "availability" of complex type ProcedureAvailabilityPropertyType
	Availability []*ProcedureAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "aircraftCharacteristic" of complex type AircraftCharacteristicPropertyType
	AircraftCharacteristic []*AircraftCharacteristicPropertyType `xml:"aircraftCharacteristic,omitempty"`

	// generated from element "flightTransition" of complex type ProcedureTransitionPropertyType
	FlightTransition []*ProcedureTransitionPropertyType `xml:"flightTransition,omitempty"`

	// generated from element "guidanceFacility_specialNavigationSystem" of complex type SpecialNavigationSystemPropertyType
	GuidanceFacility_specialNavigationSystem *SpecialNavigationSystemPropertyType `xml:"guidanceFacility_specialNavigationSystem,omitempty"`

	// generated from element "guidanceFacility_radar" of complex type RadarSystemPropertyType
	GuidanceFacility_radar *RadarSystemPropertyType `xml:"guidanceFacility_radar,omitempty"`

	// generated from element "guidanceFacility_navaid" of complex type NavaidPropertyType
	GuidanceFacility_navaid *NavaidPropertyType `xml:"guidanceFacility_navaid,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "safeAltitude" of complex type SafeAltitudeAreaPropertyType
	SafeAltitude *SafeAltitudeAreaPropertyType `xml:"safeAltitude,omitempty"`

	// generated from element "designator" of complex type TextSIDSTARDesignatorType
	Designator *aixmdatatypes.TextSIDSTARDesignatorType `xml:"designator,omitempty"`

	// generated from element "contingencyRoute" of complex type CodeYesNoType
	ContingencyRoute *aixmdatatypes.CodeYesNoType `xml:"contingencyRoute,omitempty"`

	// generated from element "takeoff" of complex type LandingTakeoffAreaCollectionPropertyType
	Takeoff *LandingTakeoffAreaCollectionPropertyType `xml:"takeoff,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_34 `xml:"extension,omitempty"`
}

// A_extension_34 Named source within outer element "extension"
type A_extension_34 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardInstrumentDepartureExtension" of complex type AbstractExtensionType
	AbstractStandardInstrumentDepartureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardInstrumentDepartureExtension,omitempty"`

	// generated from element "AbstractProcedureExtension" of complex type AbstractExtensionType
	AbstractProcedureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureExtension,omitempty"`
}

// Group_StandardInstrumentDeparturePropertyGroup UnNamed source named group "StandardInstrumentDeparturePropertyGroup"
type Group_StandardInstrumentDeparturePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ProcedurePropertyGroup" with order 4672 depth 2
	Group_ProcedurePropertyGroup

	// generated from element "designator" of complex type TextSIDSTARDesignatorType
	Designator *aixmdatatypes.TextSIDSTARDesignatorType `xml:"designator,omitempty"`

	// generated from element "contingencyRoute" of complex type CodeYesNoType
	ContingencyRoute *aixmdatatypes.CodeYesNoType `xml:"contingencyRoute,omitempty"`

	// generated from element "takeoff" of complex type LandingTakeoffAreaCollectionPropertyType
	Takeoff *LandingTakeoffAreaCollectionPropertyType `xml:"takeoff,omitempty"`
}

// SafeAltitudeAreaPropertyType Named source named complex type "SafeAltitudeAreaPropertyType"
type SafeAltitudeAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// SafeAltitudeAreaType Named source named complex type "SafeAltitudeAreaType"
type SafeAltitudeAreaType struct {
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

	// generated from element "timeSlice" of complex type SafeAltitudeAreaTimeSlicePropertyType
	TimeSlice []*SafeAltitudeAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SafeAltitudeAreaTimeSlicePropertyType Named source named complex type "SafeAltitudeAreaTimeSlicePropertyType"
type SafeAltitudeAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SafeAltitudeAreaTimeSlice" of complex type SafeAltitudeAreaTimeSliceType
	SafeAltitudeAreaTimeSlice *SafeAltitudeAreaTimeSliceType `xml:"SafeAltitudeAreaTimeSlice,omitempty"`
}

// SafeAltitudeAreaTimeSliceType Named source named complex type "SafeAltitudeAreaTimeSliceType"
type SafeAltitudeAreaTimeSliceType struct {
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

	// generated from group "aixm:SafeAltitudeAreaPropertyGroup" with order 4691 depth 2
	Group_SafeAltitudeAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_35 `xml:"extension,omitempty"`
}

// A_extension_35 Named source within outer element "extension"
type A_extension_35 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSafeAltitudeAreaExtension" of complex type AbstractExtensionType
	AbstractSafeAltitudeAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSafeAltitudeAreaExtension,omitempty"`
}

// Group_SafeAltitudeAreaPropertyGroup UnNamed source named group "SafeAltitudeAreaPropertyGroup"
type Group_SafeAltitudeAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "safeAreaType" of complex type CodeSafeAltitudeType
	SafeAreaType *aixmdatatypes.CodeSafeAltitudeType `xml:"safeAreaType,omitempty"`

	// generated from element "centrePoint_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	CentrePoint_fixDesignatedPoint *DesignatedPointPropertyType `xml:"centrePoint_fixDesignatedPoint,omitempty"`

	// generated from element "centrePoint_navaidSystem" of complex type NavaidPropertyType
	CentrePoint_navaidSystem *NavaidPropertyType `xml:"centrePoint_navaidSystem,omitempty"`

	// generated from element "centrePoint_position" of complex type PointPropertyType
	CentrePoint_position *PointPropertyType `xml:"centrePoint_position,omitempty"`

	// generated from element "centrePoint_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	CentrePoint_runwayPoint *RunwayCentrelinePointPropertyType `xml:"centrePoint_runwayPoint,omitempty"`

	// generated from element "centrePoint_aimingPoint" of complex type TouchDownLiftOffPropertyType
	CentrePoint_aimingPoint *TouchDownLiftOffPropertyType `xml:"centrePoint_aimingPoint,omitempty"`

	// generated from element "centrePoint_airportReferencePoint" of complex type AirportHeliportPropertyType
	CentrePoint_airportReferencePoint *AirportHeliportPropertyType `xml:"centrePoint_airportReferencePoint,omitempty"`

	// generated from element "sector" of complex type SafeAltitudeAreaSectorPropertyType
	Sector []*SafeAltitudeAreaSectorPropertyType `xml:"sector,omitempty"`

	// generated from element "location" of complex type AirportHeliportPropertyType
	Location []*AirportHeliportPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SafeAltitudeAreaSectorType Named source named complex type "SafeAltitudeAreaSectorType"
type SafeAltitudeAreaSectorType struct {
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

	// generated from group "aixm:SafeAltitudeAreaSectorPropertyGroup" with order 4714 depth 2
	Group_SafeAltitudeAreaSectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_36 `xml:"extension,omitempty"`
}

// A_extension_36 Named source within outer element "extension"
type A_extension_36 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSafeAltitudeAreaSectorExtension" of complex type AbstractExtensionType
	AbstractSafeAltitudeAreaSectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSafeAltitudeAreaSectorExtension,omitempty"`
}

// Group_SafeAltitudeAreaSectorPropertyGroup UnNamed source named group "SafeAltitudeAreaSectorPropertyGroup"
type Group_SafeAltitudeAreaSectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "bufferWidth" of complex type ValDistanceType
	BufferWidth *aixmdatatypes.ValDistanceType `xml:"bufferWidth,omitempty"`

	// generated from element "extent" of complex type SurfacePropertyType
	Extent *SurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "significantObstacle" of complex type ObstructionPropertyType
	SignificantObstacle []*ObstructionPropertyType `xml:"significantObstacle,omitempty"`

	// generated from element "sectorDefinition" of complex type CircleSectorPropertyType
	SectorDefinition *CircleSectorPropertyType `xml:"sectorDefinition,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SafeAltitudeAreaSectorPropertyType Named source named complex type "SafeAltitudeAreaSectorPropertyType"
type SafeAltitudeAreaSectorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SafeAltitudeAreaSector" of complex type SafeAltitudeAreaSectorType
	SafeAltitudeAreaSector *SafeAltitudeAreaSectorType `xml:"SafeAltitudeAreaSector,omitempty"`
}

// AbstractApproachLegType Named source named complex type "AbstractApproachLegType"
type AbstractApproachLegType struct {
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
}

// Group_ApproachLegPropertyGroup UnNamed source named group "ApproachLegPropertyGroup"
type Group_ApproachLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:SegmentLegPropertyGroup" with order 4738 depth 2
	Group_SegmentLegPropertyGroup

	// generated from element "approach" of complex type InstrumentApproachProcedurePropertyType
	Approach *InstrumentApproachProcedurePropertyType `xml:"approach,omitempty"`
}

// ArrivalFeederLegType Named source named complex type "ArrivalFeederLegType"
type ArrivalFeederLegType struct {
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

	// generated from element "timeSlice" of complex type ArrivalFeederLegTimeSlicePropertyType
	TimeSlice []*ArrivalFeederLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ArrivalFeederLegTimeSlicePropertyType Named source named complex type "ArrivalFeederLegTimeSlicePropertyType"
type ArrivalFeederLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ArrivalFeederLegTimeSlice" of complex type ArrivalFeederLegTimeSliceType
	ArrivalFeederLegTimeSlice *ArrivalFeederLegTimeSliceType `xml:"ArrivalFeederLegTimeSlice,omitempty"`
}

// ArrivalFeederLegTimeSliceType Named source named complex type "ArrivalFeederLegTimeSliceType"
type ArrivalFeederLegTimeSliceType struct {
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

	// generated from group "aixm:ArrivalFeederLegPropertyGroup" with order 4755 depth 2
	Group_ArrivalFeederLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_37 `xml:"extension,omitempty"`
}

// A_extension_37 Named source within outer element "extension"
type A_extension_37 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractArrivalFeederLegExtension" of complex type AbstractExtensionType
	AbstractArrivalFeederLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractArrivalFeederLegExtension,omitempty"`

	// generated from element "AbstractApproachLegExtension" of complex type AbstractExtensionType
	AbstractApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_ArrivalFeederLegPropertyGroup UnNamed source named group "ArrivalFeederLegPropertyGroup"
type Group_ArrivalFeederLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ApproachLegPropertyGroup" with order 4765 depth 2
	Group_ApproachLegPropertyGroup

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`
}

// ArrivalLegType Named source named complex type "ArrivalLegType"
type ArrivalLegType struct {
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

	// generated from element "timeSlice" of complex type ArrivalLegTimeSlicePropertyType
	TimeSlice []*ArrivalLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ArrivalLegTimeSlicePropertyType Named source named complex type "ArrivalLegTimeSlicePropertyType"
type ArrivalLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ArrivalLegTimeSlice" of complex type ArrivalLegTimeSliceType
	ArrivalLegTimeSlice *ArrivalLegTimeSliceType `xml:"ArrivalLegTimeSlice,omitempty"`
}

// ArrivalLegTimeSliceType Named source named complex type "ArrivalLegTimeSliceType"
type ArrivalLegTimeSliceType struct {
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

	// generated from group "aixm:ArrivalLegPropertyGroup" with order 4782 depth 2
	Group_ArrivalLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_38 `xml:"extension,omitempty"`
}

// A_extension_38 Named source within outer element "extension"
type A_extension_38 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractArrivalLegExtension" of complex type AbstractExtensionType
	AbstractArrivalLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractArrivalLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_ArrivalLegPropertyGroup UnNamed source named group "ArrivalLegPropertyGroup"
type Group_ArrivalLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:SegmentLegPropertyGroup" with order 4791 depth 2
	Group_SegmentLegPropertyGroup

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`

	// generated from element "arrival" of complex type StandardInstrumentArrivalPropertyType
	Arrival *StandardInstrumentArrivalPropertyType `xml:"arrival,omitempty"`
}

// DepartureLegType Named source named complex type "DepartureLegType"
type DepartureLegType struct {
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

	// generated from element "timeSlice" of complex type DepartureLegTimeSlicePropertyType
	TimeSlice []*DepartureLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DepartureLegTimeSlicePropertyType Named source named complex type "DepartureLegTimeSlicePropertyType"
type DepartureLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DepartureLegTimeSlice" of complex type DepartureLegTimeSliceType
	DepartureLegTimeSlice *DepartureLegTimeSliceType `xml:"DepartureLegTimeSlice,omitempty"`
}

// DepartureLegTimeSliceType Named source named complex type "DepartureLegTimeSliceType"
type DepartureLegTimeSliceType struct {
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

	// generated from group "aixm:DepartureLegPropertyGroup" with order 4809 depth 2
	Group_DepartureLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_39 `xml:"extension,omitempty"`
}

// A_extension_39 Named source within outer element "extension"
type A_extension_39 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDepartureLegExtension" of complex type AbstractExtensionType
	AbstractDepartureLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDepartureLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_DepartureLegPropertyGroup UnNamed source named group "DepartureLegPropertyGroup"
type Group_DepartureLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:SegmentLegPropertyGroup" with order 4818 depth 2
	Group_SegmentLegPropertyGroup

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`

	// generated from element "minimumObstacleClearanceAltitude" of complex type ValDistanceVerticalType
	MinimumObstacleClearanceAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumObstacleClearanceAltitude,omitempty"`

	// generated from element "departure" of complex type StandardInstrumentDeparturePropertyType
	Departure *StandardInstrumentDeparturePropertyType `xml:"departure,omitempty"`

	// generated from element "condition" of complex type DepartureArrivalConditionPropertyType
	Condition []*DepartureArrivalConditionPropertyType `xml:"condition,omitempty"`
}

// FinalLegType Named source named complex type "FinalLegType"
type FinalLegType struct {
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

	// generated from element "timeSlice" of complex type FinalLegTimeSlicePropertyType
	TimeSlice []*FinalLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// FinalLegTimeSlicePropertyType Named source named complex type "FinalLegTimeSlicePropertyType"
type FinalLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "FinalLegTimeSlice" of complex type FinalLegTimeSliceType
	FinalLegTimeSlice *FinalLegTimeSliceType `xml:"FinalLegTimeSlice,omitempty"`
}

// FinalLegTimeSliceType Named source named complex type "FinalLegTimeSliceType"
type FinalLegTimeSliceType struct {
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

	// generated from group "aixm:FinalLegPropertyGroup" with order 4838 depth 2
	Group_FinalLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_40 `xml:"extension,omitempty"`
}

// A_extension_40 Named source within outer element "extension"
type A_extension_40 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFinalLegExtension" of complex type AbstractExtensionType
	AbstractFinalLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFinalLegExtension,omitempty"`

	// generated from element "AbstractApproachLegExtension" of complex type AbstractExtensionType
	AbstractApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_FinalLegPropertyGroup UnNamed source named group "FinalLegPropertyGroup"
type Group_FinalLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ApproachLegPropertyGroup" with order 4848 depth 2
	Group_ApproachLegPropertyGroup

	// generated from element "guidanceSystem" of complex type CodeFinalGuidanceType
	GuidanceSystem *aixmdatatypes.CodeFinalGuidanceType `xml:"guidanceSystem,omitempty"`

	// generated from element "landingSystemCategory" of complex type CodeApproachGuidanceType
	LandingSystemCategory *aixmdatatypes.CodeApproachGuidanceType `xml:"landingSystemCategory,omitempty"`

	// generated from element "minimumBaroVnavTemperature" of complex type ValTemperatureType
	MinimumBaroVnavTemperature *aixmdatatypes.ValTemperatureType `xml:"minimumBaroVnavTemperature,omitempty"`

	// generated from element "rnpDMEAuthorized" of complex type CodeYesNoType
	RnpDMEAuthorized *aixmdatatypes.CodeYesNoType `xml:"rnpDMEAuthorized,omitempty"`

	// generated from element "courseOffsetAngle" of complex type ValBearingType
	CourseOffsetAngle *aixmdatatypes.ValBearingType `xml:"courseOffsetAngle,omitempty"`

	// generated from element "courseOffsetSide" of complex type CodeSideType
	CourseOffsetSide *aixmdatatypes.CodeSideType `xml:"courseOffsetSide,omitempty"`

	// generated from element "courseCentrelineDistance" of complex type ValDistanceType
	CourseCentrelineDistance *aixmdatatypes.ValDistanceType `xml:"courseCentrelineDistance,omitempty"`

	// generated from element "courseOffsetDistance" of complex type ValDistanceType
	CourseOffsetDistance *aixmdatatypes.ValDistanceType `xml:"courseOffsetDistance,omitempty"`

	// generated from element "courseCentrelineIntersect" of complex type CodeRelativePositionType
	CourseCentrelineIntersect *aixmdatatypes.CodeRelativePositionType `xml:"courseCentrelineIntersect,omitempty"`

	// generated from element "condition" of complex type ApproachConditionPropertyType
	Condition []*ApproachConditionPropertyType `xml:"condition,omitempty"`

	// generated from element "finalPathAlignmentPoint_fixDesignatedPoint" of complex type DesignatedPointPropertyType
	FinalPathAlignmentPoint_fixDesignatedPoint *DesignatedPointPropertyType `xml:"finalPathAlignmentPoint_fixDesignatedPoint,omitempty"`

	// generated from element "finalPathAlignmentPoint_navaidSystem" of complex type NavaidPropertyType
	FinalPathAlignmentPoint_navaidSystem *NavaidPropertyType `xml:"finalPathAlignmentPoint_navaidSystem,omitempty"`

	// generated from element "finalPathAlignmentPoint_position" of complex type PointPropertyType
	FinalPathAlignmentPoint_position *PointPropertyType `xml:"finalPathAlignmentPoint_position,omitempty"`

	// generated from element "finalPathAlignmentPoint_runwayPoint" of complex type RunwayCentrelinePointPropertyType
	FinalPathAlignmentPoint_runwayPoint *RunwayCentrelinePointPropertyType `xml:"finalPathAlignmentPoint_runwayPoint,omitempty"`

	// generated from element "finalPathAlignmentPoint_aimingPoint" of complex type TouchDownLiftOffPropertyType
	FinalPathAlignmentPoint_aimingPoint *TouchDownLiftOffPropertyType `xml:"finalPathAlignmentPoint_aimingPoint,omitempty"`

	// generated from element "finalPathAlignmentPoint_airportReferencePoint" of complex type AirportHeliportPropertyType
	FinalPathAlignmentPoint_airportReferencePoint *AirportHeliportPropertyType `xml:"finalPathAlignmentPoint_airportReferencePoint,omitempty"`

	// generated from element "visualDescentPoint" of complex type TerminalSegmentPointPropertyType
	VisualDescentPoint *TerminalSegmentPointPropertyType `xml:"visualDescentPoint,omitempty"`

	// generated from element "FASData" of complex type FASDataBlockPropertyType
	FASData *FASDataBlockPropertyType `xml:"FASData,omitempty"`
}

// HoldingUseType Named source named complex type "HoldingUseType"
type HoldingUseType struct {
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

	// generated from group "aixm:HoldingUsePropertyGroup" with order 4872 depth 2
	Group_HoldingUsePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_41 `xml:"extension,omitempty"`
}

// A_extension_41 Named source within outer element "extension"
type A_extension_41 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractHoldingUseExtension" of complex type AbstractExtensionType
	AbstractHoldingUseExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractHoldingUseExtension,omitempty"`
}

// Group_HoldingUsePropertyGroup UnNamed source named group "HoldingUsePropertyGroup"
type Group_HoldingUsePropertyGroup struct {

	// insertion point for fields

	// generated from element "holdingUse" of complex type CodeHoldingUseType
	HoldingUse *aixmdatatypes.CodeHoldingUseType `xml:"holdingUse,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "instructedAltitude" of complex type ValDistanceVerticalType
	InstructedAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"instructedAltitude,omitempty"`

	// generated from element "instructionAltitudeReference" of complex type CodeVerticalReferenceType
	InstructionAltitudeReference *aixmdatatypes.CodeVerticalReferenceType `xml:"instructionAltitudeReference,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theHoldingPattern" of complex type HoldingPatternPropertyType
	TheHoldingPattern *HoldingPatternPropertyType `xml:"theHoldingPattern,omitempty"`
}

// HoldingUsePropertyType Named source named complex type "HoldingUsePropertyType"
type HoldingUsePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "HoldingUse" of complex type HoldingUseType
	HoldingUse *HoldingUseType `xml:"HoldingUse,omitempty"`
}

// InitialLegType Named source named complex type "InitialLegType"
type InitialLegType struct {
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

	// generated from element "timeSlice" of complex type InitialLegTimeSlicePropertyType
	TimeSlice []*InitialLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// InitialLegTimeSlicePropertyType Named source named complex type "InitialLegTimeSlicePropertyType"
type InitialLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "InitialLegTimeSlice" of complex type InitialLegTimeSliceType
	InitialLegTimeSlice *InitialLegTimeSliceType `xml:"InitialLegTimeSlice,omitempty"`
}

// InitialLegTimeSliceType Named source named complex type "InitialLegTimeSliceType"
type InitialLegTimeSliceType struct {
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

	// generated from group "aixm:InitialLegPropertyGroup" with order 4904 depth 2
	Group_InitialLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_42 `xml:"extension,omitempty"`
}

// A_extension_42 Named source within outer element "extension"
type A_extension_42 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractInitialLegExtension" of complex type AbstractExtensionType
	AbstractInitialLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractInitialLegExtension,omitempty"`

	// generated from element "AbstractApproachLegExtension" of complex type AbstractExtensionType
	AbstractApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_InitialLegPropertyGroup UnNamed source named group "InitialLegPropertyGroup"
type Group_InitialLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ApproachLegPropertyGroup" with order 4914 depth 2
	Group_ApproachLegPropertyGroup

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`
}

// IntermediateLegType Named source named complex type "IntermediateLegType"
type IntermediateLegType struct {
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

	// generated from element "timeSlice" of complex type IntermediateLegTimeSlicePropertyType
	TimeSlice []*IntermediateLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// IntermediateLegTimeSlicePropertyType Named source named complex type "IntermediateLegTimeSlicePropertyType"
type IntermediateLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "IntermediateLegTimeSlice" of complex type IntermediateLegTimeSliceType
	IntermediateLegTimeSlice *IntermediateLegTimeSliceType `xml:"IntermediateLegTimeSlice,omitempty"`
}

// IntermediateLegTimeSliceType Named source named complex type "IntermediateLegTimeSliceType"
type IntermediateLegTimeSliceType struct {
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

	// generated from group "aixm:IntermediateLegPropertyGroup" with order 4931 depth 2
	Group_IntermediateLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_43 `xml:"extension,omitempty"`
}

// A_extension_43 Named source within outer element "extension"
type A_extension_43 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractIntermediateLegExtension" of complex type AbstractExtensionType
	AbstractIntermediateLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractIntermediateLegExtension,omitempty"`

	// generated from element "AbstractApproachLegExtension" of complex type AbstractExtensionType
	AbstractApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_IntermediateLegPropertyGroup UnNamed source named group "IntermediateLegPropertyGroup"
type Group_IntermediateLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ApproachLegPropertyGroup" with order 4941 depth 2
	Group_ApproachLegPropertyGroup

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`
}

// LandingTakeoffAreaCollectionType Named source named complex type "LandingTakeoffAreaCollectionType"
type LandingTakeoffAreaCollectionType struct {
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

	// generated from group "aixm:LandingTakeoffAreaCollectionPropertyGroup" with order 4947 depth 2
	Group_LandingTakeoffAreaCollectionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_44 `xml:"extension,omitempty"`
}

// A_extension_44 Named source within outer element "extension"
type A_extension_44 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractLandingTakeoffAreaCollectionExtension" of complex type AbstractExtensionType
	AbstractLandingTakeoffAreaCollectionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLandingTakeoffAreaCollectionExtension,omitempty"`
}

// Group_LandingTakeoffAreaCollectionPropertyGroup UnNamed source named group "LandingTakeoffAreaCollectionPropertyGroup"
type Group_LandingTakeoffAreaCollectionPropertyGroup struct {

	// insertion point for fields

	// generated from element "runway" of complex type RunwayDirectionPropertyType
	Runway []*RunwayDirectionPropertyType `xml:"runway,omitempty"`

	// generated from element "TLOF" of complex type TouchDownLiftOffPropertyType
	TLOF []*TouchDownLiftOffPropertyType `xml:"TLOF,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// LandingTakeoffAreaCollectionPropertyType Named source named complex type "LandingTakeoffAreaCollectionPropertyType"
type LandingTakeoffAreaCollectionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "LandingTakeoffAreaCollection" of complex type LandingTakeoffAreaCollectionType
	LandingTakeoffAreaCollection *LandingTakeoffAreaCollectionType `xml:"LandingTakeoffAreaCollection,omitempty"`
}

// MissedApproachLegType Named source named complex type "MissedApproachLegType"
type MissedApproachLegType struct {
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

	// generated from element "timeSlice" of complex type MissedApproachLegTimeSlicePropertyType
	TimeSlice []*MissedApproachLegTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// MissedApproachLegTimeSlicePropertyType Named source named complex type "MissedApproachLegTimeSlicePropertyType"
type MissedApproachLegTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MissedApproachLegTimeSlice" of complex type MissedApproachLegTimeSliceType
	MissedApproachLegTimeSlice *MissedApproachLegTimeSliceType `xml:"MissedApproachLegTimeSlice,omitempty"`
}

// MissedApproachLegTimeSliceType Named source named complex type "MissedApproachLegTimeSliceType"
type MissedApproachLegTimeSliceType struct {
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

	// generated from group "aixm:MissedApproachLegPropertyGroup" with order 4976 depth 2
	Group_MissedApproachLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_45 `xml:"extension,omitempty"`
}

// A_extension_45 Named source within outer element "extension"
type A_extension_45 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMissedApproachLegExtension" of complex type AbstractExtensionType
	AbstractMissedApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMissedApproachLegExtension,omitempty"`

	// generated from element "AbstractApproachLegExtension" of complex type AbstractExtensionType
	AbstractApproachLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLegExtension,omitempty"`

	// generated from element "AbstractSegmentLegExtension" of complex type AbstractExtensionType
	AbstractSegmentLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSegmentLegExtension,omitempty"`
}

// Group_MissedApproachLegPropertyGroup UnNamed source named group "MissedApproachLegPropertyGroup"
type Group_MissedApproachLegPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ApproachLegPropertyGroup" with order 4986 depth 2
	Group_ApproachLegPropertyGroup

	// generated from element "type" of complex type CodeMissedApproachType
	Type_XSD *aixmdatatypes.CodeMissedApproachType `xml:"type,omitempty"`

	// generated from element "thresholdAfterMAPT" of complex type CodeYesNoType
	ThresholdAfterMAPT *aixmdatatypes.CodeYesNoType `xml:"thresholdAfterMAPT,omitempty"`

	// generated from element "heightMAPT" of complex type ValDistanceVerticalType
	HeightMAPT *aixmdatatypes.ValDistanceVerticalType `xml:"heightMAPT,omitempty"`

	// generated from element "requiredNavigationPerformance" of complex type CodeRNPType
	RequiredNavigationPerformance *aixmdatatypes.CodeRNPType `xml:"requiredNavigationPerformance,omitempty"`

	// generated from element "condition" of complex type ApproachConditionPropertyType
	Condition []*ApproachConditionPropertyType `xml:"condition,omitempty"`
}

// NavigationAreaRestrictionType Named source named complex type "NavigationAreaRestrictionType"
type NavigationAreaRestrictionType struct {
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

	// generated from element "timeSlice" of complex type NavigationAreaRestrictionTimeSlicePropertyType
	TimeSlice []*NavigationAreaRestrictionTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// NavigationAreaRestrictionTimeSlicePropertyType Named source named complex type "NavigationAreaRestrictionTimeSlicePropertyType"
type NavigationAreaRestrictionTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "NavigationAreaRestrictionTimeSlice" of complex type NavigationAreaRestrictionTimeSliceType
	NavigationAreaRestrictionTimeSlice *NavigationAreaRestrictionTimeSliceType `xml:"NavigationAreaRestrictionTimeSlice,omitempty"`
}

// NavigationAreaRestrictionTimeSliceType Named source named complex type "NavigationAreaRestrictionTimeSliceType"
type NavigationAreaRestrictionTimeSliceType struct {
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

	// generated from group "aixm:NavigationAreaRestrictionPropertyGroup" with order 5007 depth 2
	Group_NavigationAreaRestrictionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_46 `xml:"extension,omitempty"`
}

// A_extension_46 Named source within outer element "extension"
type A_extension_46 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavigationAreaRestrictionExtension" of complex type AbstractExtensionType
	AbstractNavigationAreaRestrictionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavigationAreaRestrictionExtension,omitempty"`
}

// Group_NavigationAreaRestrictionPropertyGroup UnNamed source named group "NavigationAreaRestrictionPropertyGroup"
type Group_NavigationAreaRestrictionPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeNavigationAreaRestrictionType
	Type_XSD *aixmdatatypes.CodeNavigationAreaRestrictionType `xml:"type,omitempty"`

	// generated from element "procedure" of complex type ProcedurePropertyType
	Procedure []*ProcedurePropertyType `xml:"procedure,omitempty"`

	// generated from element "designSurface" of complex type ObstacleAssessmentAreaPropertyType
	DesignSurface *ObstacleAssessmentAreaPropertyType `xml:"designSurface,omitempty"`

	// generated from element "sectorDefinition" of complex type CircleSectorPropertyType
	SectorDefinition *CircleSectorPropertyType `xml:"sectorDefinition,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ProcedurePropertyType Named source named complex type "ProcedurePropertyType"
type ProcedurePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AbstractProcedureType Named source named complex type "AbstractProcedureType"
type AbstractProcedureType struct {
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
}

// Group_ProcedurePropertyGroup UnNamed source named group "ProcedurePropertyGroup"
type Group_ProcedurePropertyGroup struct {

	// insertion point for fields

	// generated from element "communicationFailureInstruction" of complex type TextInstructionType
	CommunicationFailureInstruction *aixmdatatypes.TextInstructionType `xml:"communicationFailureInstruction,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "designCriteria" of complex type CodeDesignStandardType
	DesignCriteria *aixmdatatypes.CodeDesignStandardType `xml:"designCriteria,omitempty"`

	// generated from element "codingStandard" of complex type CodeProcedureCodingStandardType
	CodingStandard *aixmdatatypes.CodeProcedureCodingStandardType `xml:"codingStandard,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "RNAV" of complex type CodeYesNoType
	RNAV *aixmdatatypes.CodeYesNoType `xml:"RNAV,omitempty"`

	// generated from element "availability" of complex type ProcedureAvailabilityPropertyType
	Availability []*ProcedureAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "aircraftCharacteristic" of complex type AircraftCharacteristicPropertyType
	AircraftCharacteristic []*AircraftCharacteristicPropertyType `xml:"aircraftCharacteristic,omitempty"`

	// generated from element "flightTransition" of complex type ProcedureTransitionPropertyType
	FlightTransition []*ProcedureTransitionPropertyType `xml:"flightTransition,omitempty"`

	// generated from element "guidanceFacility_specialNavigationSystem" of complex type SpecialNavigationSystemPropertyType
	GuidanceFacility_specialNavigationSystem *SpecialNavigationSystemPropertyType `xml:"guidanceFacility_specialNavigationSystem,omitempty"`

	// generated from element "guidanceFacility_radar" of complex type RadarSystemPropertyType
	GuidanceFacility_radar *RadarSystemPropertyType `xml:"guidanceFacility_radar,omitempty"`

	// generated from element "guidanceFacility_navaid" of complex type NavaidPropertyType
	GuidanceFacility_navaid *NavaidPropertyType `xml:"guidanceFacility_navaid,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "safeAltitude" of complex type SafeAltitudeAreaPropertyType
	SafeAltitude *SafeAltitudeAreaPropertyType `xml:"safeAltitude,omitempty"`
}

// ProcedureDMEType Named source named complex type "ProcedureDMEType"
type ProcedureDMEType struct {
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

	// generated from element "timeSlice" of complex type ProcedureDMETimeSlicePropertyType
	TimeSlice []*ProcedureDMETimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ProcedureDMETimeSlicePropertyType Named source named complex type "ProcedureDMETimeSlicePropertyType"
type ProcedureDMETimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ProcedureDMETimeSlice" of complex type ProcedureDMETimeSliceType
	ProcedureDMETimeSlice *ProcedureDMETimeSliceType `xml:"ProcedureDMETimeSlice,omitempty"`
}

// ProcedureDMETimeSliceType Named source named complex type "ProcedureDMETimeSliceType"
type ProcedureDMETimeSliceType struct {
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

	// generated from group "aixm:ProcedureDMEPropertyGroup" with order 5060 depth 2
	Group_ProcedureDMEPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_47 `xml:"extension,omitempty"`
}

// A_extension_47 Named source within outer element "extension"
type A_extension_47 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractProcedureDMEExtension" of complex type AbstractExtensionType
	AbstractProcedureDMEExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureDMEExtension,omitempty"`
}

// Group_ProcedureDMEPropertyGroup UnNamed source named group "ProcedureDMEPropertyGroup"
type Group_ProcedureDMEPropertyGroup struct {

	// insertion point for fields

	// generated from element "criticalDME" of complex type CodeYesNoType
	CriticalDME *aixmdatatypes.CodeYesNoType `xml:"criticalDME,omitempty"`

	// generated from element "satisfactory" of complex type CodeYesNoType
	Satisfactory *aixmdatatypes.CodeYesNoType `xml:"satisfactory,omitempty"`

	// generated from element "DME" of complex type DMEPropertyType
	DME *DMEPropertyType `xml:"DME,omitempty"`

	// generated from element "segmentLeg" of complex type SegmentLegPropertyType
	SegmentLeg *SegmentLegPropertyType `xml:"segmentLeg,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ProcedureTransitionType Named source named complex type "ProcedureTransitionType"
type ProcedureTransitionType struct {
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

	// generated from group "aixm:ProcedureTransitionPropertyGroup" with order 5077 depth 2
	Group_ProcedureTransitionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_48 `xml:"extension,omitempty"`
}

// A_extension_48 Named source within outer element "extension"
type A_extension_48 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractProcedureTransitionExtension" of complex type AbstractExtensionType
	AbstractProcedureTransitionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureTransitionExtension,omitempty"`
}

// Group_ProcedureTransitionPropertyGroup UnNamed source named group "ProcedureTransitionPropertyGroup"
type Group_ProcedureTransitionPropertyGroup struct {

	// insertion point for fields

	// generated from element "transitionId" of complex type CodeDesignatedPointDesignatorType
	TransitionId *aixmdatatypes.CodeDesignatedPointDesignatorType `xml:"transitionId,omitempty"`

	// generated from element "type" of complex type CodeProcedurePhaseType
	Type_XSD *aixmdatatypes.CodeProcedurePhaseType `xml:"type,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "vectorHeading" of complex type ValBearingType
	VectorHeading *aixmdatatypes.ValBearingType `xml:"vectorHeading,omitempty"`

	// generated from element "departureRunwayTransition" of complex type LandingTakeoffAreaCollectionPropertyType
	DepartureRunwayTransition *LandingTakeoffAreaCollectionPropertyType `xml:"departureRunwayTransition,omitempty"`

	// generated from element "trajectory" of complex type CurvePropertyType
	Trajectory *CurvePropertyType `xml:"trajectory,omitempty"`

	// generated from element "transitionLeg" of complex type ProcedureTransitionLegPropertyType
	TransitionLeg []*ProcedureTransitionLegPropertyType `xml:"transitionLeg,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ProcedureTransitionPropertyType Named source named complex type "ProcedureTransitionPropertyType"
type ProcedureTransitionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ProcedureTransition" of complex type ProcedureTransitionType
	ProcedureTransition *ProcedureTransitionType `xml:"ProcedureTransition,omitempty"`
}

// ProcedureTransitionLegType Named source named complex type "ProcedureTransitionLegType"
type ProcedureTransitionLegType struct {
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

	// generated from group "aixm:ProcedureTransitionLegPropertyGroup" with order 5100 depth 2
	Group_ProcedureTransitionLegPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_49 `xml:"extension,omitempty"`
}

// A_extension_49 Named source within outer element "extension"
type A_extension_49 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractProcedureTransitionLegExtension" of complex type AbstractExtensionType
	AbstractProcedureTransitionLegExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureTransitionLegExtension,omitempty"`
}

// Group_ProcedureTransitionLegPropertyGroup UnNamed source named group "ProcedureTransitionLegPropertyGroup"
type Group_ProcedureTransitionLegPropertyGroup struct {

	// insertion point for fields

	// generated from element "seqNumberARINC" of complex type NoSequenceType
	SeqNumberARINC *aixmdatatypes.NoSequenceType `xml:"seqNumberARINC,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theSegmentLeg" of complex type SegmentLegPropertyType
	TheSegmentLeg *SegmentLegPropertyType `xml:"theSegmentLeg,omitempty"`
}

// ProcedureTransitionLegPropertyType Named source named complex type "ProcedureTransitionLegPropertyType"
type ProcedureTransitionLegPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ProcedureTransitionLeg" of complex type ProcedureTransitionLegType
	ProcedureTransitionLeg *ProcedureTransitionLegType `xml:"ProcedureTransitionLeg,omitempty"`
}

// SegmentLegPropertyType Named source named complex type "SegmentLegPropertyType"
type SegmentLegPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AbstractSegmentLegType Named source named complex type "AbstractSegmentLegType"
type AbstractSegmentLegType struct {
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
}

// Group_SegmentLegPropertyGroup UnNamed source named group "SegmentLegPropertyGroup"
type Group_SegmentLegPropertyGroup struct {

	// insertion point for fields

	// generated from element "endConditionDesignator" of complex type CodeSegmentTerminationType
	EndConditionDesignator *aixmdatatypes.CodeSegmentTerminationType `xml:"endConditionDesignator,omitempty"`

	// generated from element "legPath" of complex type CodeTrajectoryType
	LegPath *aixmdatatypes.CodeTrajectoryType `xml:"legPath,omitempty"`

	// generated from element "legTypeARINC" of complex type CodeSegmentPathType
	LegTypeARINC *aixmdatatypes.CodeSegmentPathType `xml:"legTypeARINC,omitempty"`

	// generated from element "course" of complex type ValBearingType
	Course *aixmdatatypes.ValBearingType `xml:"course,omitempty"`

	// generated from element "courseType" of complex type CodeCourseType
	CourseType *aixmdatatypes.CodeCourseType `xml:"courseType,omitempty"`

	// generated from element "courseDirection" of complex type CodeDirectionReferenceType
	CourseDirection *aixmdatatypes.CodeDirectionReferenceType `xml:"courseDirection,omitempty"`

	// generated from element "turnDirection" of complex type CodeDirectionTurnType
	TurnDirection *aixmdatatypes.CodeDirectionTurnType `xml:"turnDirection,omitempty"`

	// generated from element "speedLimit" of complex type ValSpeedType
	SpeedLimit *aixmdatatypes.ValSpeedType `xml:"speedLimit,omitempty"`

	// generated from element "speedReference" of complex type CodeSpeedReferenceType
	SpeedReference *aixmdatatypes.CodeSpeedReferenceType `xml:"speedReference,omitempty"`

	// generated from element "speedInterpretation" of complex type CodeAltitudeUseType
	SpeedInterpretation *aixmdatatypes.CodeAltitudeUseType `xml:"speedInterpretation,omitempty"`

	// generated from element "bankAngle" of complex type ValAngleType
	BankAngle *aixmdatatypes.ValAngleType `xml:"bankAngle,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "duration" of complex type ValDurationType
	Duration *aixmdatatypes.ValDurationType `xml:"duration,omitempty"`

	// generated from element "procedureTurnRequired" of complex type CodeYesNoType
	ProcedureTurnRequired *aixmdatatypes.CodeYesNoType `xml:"procedureTurnRequired,omitempty"`

	// generated from element "upperLimitAltitude" of complex type ValDistanceVerticalType
	UpperLimitAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimitAltitude,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimitAltitude" of complex type ValDistanceVerticalType
	LowerLimitAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimitAltitude,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "altitudeInterpretation" of complex type CodeAltitudeUseType
	AltitudeInterpretation *aixmdatatypes.CodeAltitudeUseType `xml:"altitudeInterpretation,omitempty"`

	// generated from element "altitudeOverrideATC" of complex type ValDistanceVerticalType
	AltitudeOverrideATC *aixmdatatypes.ValDistanceVerticalType `xml:"altitudeOverrideATC,omitempty"`

	// generated from element "altitudeOverrideReference" of complex type CodeVerticalReferenceType
	AltitudeOverrideReference *aixmdatatypes.CodeVerticalReferenceType `xml:"altitudeOverrideReference,omitempty"`

	// generated from element "verticalAngle" of complex type ValAngleType
	VerticalAngle *aixmdatatypes.ValAngleType `xml:"verticalAngle,omitempty"`

	// generated from element "startPoint" of complex type TerminalSegmentPointPropertyType
	StartPoint *TerminalSegmentPointPropertyType `xml:"startPoint,omitempty"`

	// generated from element "endPoint" of complex type TerminalSegmentPointPropertyType
	EndPoint *TerminalSegmentPointPropertyType `xml:"endPoint,omitempty"`

	// generated from element "trajectory" of complex type CurvePropertyType
	Trajectory *CurvePropertyType `xml:"trajectory,omitempty"`

	// generated from element "arcCentre" of complex type TerminalSegmentPointPropertyType
	ArcCentre *TerminalSegmentPointPropertyType `xml:"arcCentre,omitempty"`

	// generated from element "angle" of complex type AngleIndicationPropertyType
	Angle *AngleIndicationPropertyType `xml:"angle,omitempty"`

	// generated from element "distance" of complex type DistanceIndicationPropertyType
	Distance *DistanceIndicationPropertyType `xml:"distance,omitempty"`

	// generated from element "aircraftCategory" of complex type AircraftCharacteristicPropertyType
	AircraftCategory []*AircraftCharacteristicPropertyType `xml:"aircraftCategory,omitempty"`

	// generated from element "holding" of complex type HoldingUsePropertyType
	Holding *HoldingUsePropertyType `xml:"holding,omitempty"`

	// generated from element "designSurface" of complex type ObstacleAssessmentAreaPropertyType
	DesignSurface []*ObstacleAssessmentAreaPropertyType `xml:"designSurface,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ProcedureAvailabilityType Named source named complex type "ProcedureAvailabilityType"
type ProcedureAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 5158 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:ProcedureAvailabilityPropertyGroup" with order 5159 depth 2
	Group_ProcedureAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_50 `xml:"extension,omitempty"`
}

// A_extension_50 Named source within outer element "extension"
type A_extension_50 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractProcedureAvailabilityExtension" of complex type AbstractExtensionType
	AbstractProcedureAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractProcedureAvailabilityExtension,omitempty"`
}

// Group_ProcedureAvailabilityPropertyGroup UnNamed source named group "ProcedureAvailabilityPropertyGroup"
type Group_ProcedureAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "status" of complex type CodeProcedureAvailabilityType
	Status *aixmdatatypes.CodeProcedureAvailabilityType `xml:"status,omitempty"`
}

// ProcedureAvailabilityPropertyType Named source named complex type "ProcedureAvailabilityPropertyType"
type ProcedureAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ProcedureAvailability" of complex type ProcedureAvailabilityType
	ProcedureAvailability *ProcedureAvailabilityType `xml:"ProcedureAvailability,omitempty"`
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

	// generated from group "aixm:AircraftCharacteristicPropertyGroup" with order 6289 depth 2
	Group_AircraftCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_51 `xml:"extension,omitempty"`
}

// A_extension_51 Named source within outer element "extension"
type A_extension_51 struct {
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

// CircleSectorType Named source named complex type "CircleSectorType"
type CircleSectorType struct {
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

	// generated from group "aixm:CircleSectorPropertyGroup" with order 6370 depth 2
	Group_CircleSectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_52 `xml:"extension,omitempty"`
}

// A_extension_52 Named source within outer element "extension"
type A_extension_52 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCircleSectorExtension" of complex type AbstractExtensionType
	AbstractCircleSectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCircleSectorExtension,omitempty"`
}

// Group_CircleSectorPropertyGroup UnNamed source named group "CircleSectorPropertyGroup"
type Group_CircleSectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "arcDirection" of complex type CodeArcDirectionType
	ArcDirection *aixmdatatypes.CodeArcDirectionType `xml:"arcDirection,omitempty"`

	// generated from element "fromAngle" of complex type ValBearingType
	FromAngle *aixmdatatypes.ValBearingType `xml:"fromAngle,omitempty"`

	// generated from element "toAngle" of complex type ValBearingType
	ToAngle *aixmdatatypes.ValBearingType `xml:"toAngle,omitempty"`

	// generated from element "angleType" of complex type CodeBearingType
	AngleType *aixmdatatypes.CodeBearingType `xml:"angleType,omitempty"`

	// generated from element "angleDirectionReference" of complex type CodeDirectionReferenceType
	AngleDirectionReference *aixmdatatypes.CodeDirectionReferenceType `xml:"angleDirectionReference,omitempty"`

	// generated from element "innerDistance" of complex type ValDistanceType
	InnerDistance *aixmdatatypes.ValDistanceType `xml:"innerDistance,omitempty"`

	// generated from element "outerDistance" of complex type ValDistanceType
	OuterDistance *aixmdatatypes.ValDistanceType `xml:"outerDistance,omitempty"`

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// CircleSectorPropertyType Named source named complex type "CircleSectorPropertyType"
type CircleSectorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "CircleSector" of complex type CircleSectorType
	CircleSector *CircleSectorType `xml:"CircleSector,omitempty"`
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

	// generated from group "aixm:TimesheetPropertyGroup" with order 6534 depth 2
	Group_TimesheetPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_53 `xml:"extension,omitempty"`
}

// A_extension_53 Named source within outer element "extension"
type A_extension_53 struct {
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

	// generated from group "aixm:AltitudeAdjustmentPropertyGroup" with order 6667 depth 2
	Group_AltitudeAdjustmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_54 `xml:"extension,omitempty"`
}

// A_extension_54 Named source within outer element "extension"
type A_extension_54 struct {
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

// HoldingAssessmentType Named source named complex type "HoldingAssessmentType"
type HoldingAssessmentType struct {
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

	// generated from element "timeSlice" of complex type HoldingAssessmentTimeSlicePropertyType
	TimeSlice []*HoldingAssessmentTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// HoldingAssessmentTimeSlicePropertyType Named source named complex type "HoldingAssessmentTimeSlicePropertyType"
type HoldingAssessmentTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "HoldingAssessmentTimeSlice" of complex type HoldingAssessmentTimeSliceType
	HoldingAssessmentTimeSlice *HoldingAssessmentTimeSliceType `xml:"HoldingAssessmentTimeSlice,omitempty"`
}

// HoldingAssessmentTimeSliceType Named source named complex type "HoldingAssessmentTimeSliceType"
type HoldingAssessmentTimeSliceType struct {
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

	// generated from group "aixm:HoldingAssessmentPropertyGroup" with order 6698 depth 2
	Group_HoldingAssessmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_55 `xml:"extension,omitempty"`
}

// A_extension_55 Named source within outer element "extension"
type A_extension_55 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractHoldingAssessmentExtension" of complex type AbstractExtensionType
	AbstractHoldingAssessmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractHoldingAssessmentExtension,omitempty"`
}

// Group_HoldingAssessmentPropertyGroup UnNamed source named group "HoldingAssessmentPropertyGroup"
type Group_HoldingAssessmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "speedLimit" of complex type ValSpeedType
	SpeedLimit *aixmdatatypes.ValSpeedType `xml:"speedLimit,omitempty"`

	// generated from element "patternTemplate" of complex type TextNameType
	PatternTemplate *aixmdatatypes.TextNameType `xml:"patternTemplate,omitempty"`

	// generated from element "turbulentAir" of complex type CodeYesNoType
	TurbulentAir *aixmdatatypes.CodeYesNoType `xml:"turbulentAir,omitempty"`

	// generated from element "legLengthToward" of complex type ValDistanceType
	LegLengthToward *aixmdatatypes.ValDistanceType `xml:"legLengthToward,omitempty"`

	// generated from element "legLengthAway" of complex type ValDistanceType
	LegLengthAway *aixmdatatypes.ValDistanceType `xml:"legLengthAway,omitempty"`

	// generated from element "holdingPoint" of complex type SegmentPointPropertyType
	HoldingPoint *SegmentPointPropertyType `xml:"holdingPoint,omitempty"`

	// generated from element "unplannedHolding" of complex type UnplannedHoldingPropertyType
	UnplannedHolding *UnplannedHoldingPropertyType `xml:"unplannedHolding,omitempty"`

	// generated from element "assessedHoldingPattern" of complex type HoldingPatternPropertyType
	AssessedHoldingPattern *HoldingPatternPropertyType `xml:"assessedHoldingPattern,omitempty"`

	// generated from element "obstacleAssessment" of complex type ObstacleAssessmentAreaPropertyType
	ObstacleAssessment []*ObstacleAssessmentAreaPropertyType `xml:"obstacleAssessment,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

	// generated from group "aixm:ObstacleAssessmentAreaPropertyGroup" with order 6724 depth 2
	Group_ObstacleAssessmentAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_56 `xml:"extension,omitempty"`
}

// A_extension_56 Named source within outer element "extension"
type A_extension_56 struct {
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

	// generated from group "aixm:ObstaclePlacementPropertyGroup" with order 6752 depth 2
	Group_ObstaclePlacementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_57 `xml:"extension,omitempty"`
}

// A_extension_57 Named source within outer element "extension"
type A_extension_57 struct {
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

	// generated from group "aixm:ObstructionPropertyGroup" with order 6772 depth 2
	Group_ObstructionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_58 `xml:"extension,omitempty"`
}

// A_extension_58 Named source within outer element "extension"
type A_extension_58 struct {
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

// RadarSystemPropertyType Named source named complex type "RadarSystemPropertyType"
type RadarSystemPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}
