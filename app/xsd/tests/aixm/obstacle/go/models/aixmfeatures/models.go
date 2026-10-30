// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeColourType
var _ *gml.CodeType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// AirportHeliportPropertyType Named source named complex type "AirportHeliportPropertyType"
type AirportHeliportPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// GroundLightSystemPropertyType Named source named complex type "GroundLightSystemPropertyType"
type GroundLightSystemPropertyType struct {
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

	// generated from group "aixm:CurvePropertyGroup" with order 2748 depth 2
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

	// generated from group "aixm:CurvePropertyGroup" with order 2748 depth 2
	Group_CurvePropertyGroup

	// generated from group "aixm:ElevatedCurvePropertyGroup" with order 2759 depth 2
	Group_ElevatedCurvePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension `xml:"extension,omitempty"`
}

// A_extension Named source within outer element "extension"
type A_extension struct {
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

// ElevatedCurvePropertyType Named source named complex type "ElevatedCurvePropertyType"
type ElevatedCurvePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ElevatedCurve" of complex type ElevatedCurveType
	ElevatedCurve *ElevatedCurveType `xml:"ElevatedCurve,omitempty"`
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

	// generated from group "aixm:PointPropertyGroup" with order 2816 depth 2
	Group_PointPropertyGroup

	// generated from group "aixm:ElevatedPointPropertyGroup" with order 2778 depth 2
	Group_ElevatedPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_1 `xml:"extension,omitempty"`
}

// A_extension_1 Named source within outer element "extension"
type A_extension_1 struct {
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

// ElevatedPointPropertyType Named source named complex type "ElevatedPointPropertyType"
type ElevatedPointPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ElevatedPoint" of complex type ElevatedPointType
	ElevatedPoint *ElevatedPointType `xml:"ElevatedPoint,omitempty"`
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2827 depth 2
	Group_SurfacePropertyGroup

	// generated from group "aixm:ElevatedSurfacePropertyGroup" with order 2797 depth 2
	Group_ElevatedSurfacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
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

// ElevatedSurfacePropertyType Named source named complex type "ElevatedSurfacePropertyType"
type ElevatedSurfacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ElevatedSurface" of complex type ElevatedSurfaceType
	ElevatedSurface *ElevatedSurfaceType `xml:"ElevatedSurface,omitempty"`
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

	// generated from group "aixm:PointPropertyGroup" with order 2816 depth 2
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2827 depth 2
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

// MarkerBeaconPropertyType Named source named complex type "MarkerBeaconPropertyType"
type MarkerBeaconPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// NavaidEquipmentPropertyType Named source named complex type "NavaidEquipmentPropertyType"
type NavaidEquipmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// SpecialNavigationStationPropertyType Named source named complex type "SpecialNavigationStationPropertyType"
type SpecialNavigationStationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AeronauticalGroundLightType Named source named complex type "AeronauticalGroundLightType"
type AeronauticalGroundLightType struct {
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

	// generated from element "timeSlice" of complex type AeronauticalGroundLightTimeSlicePropertyType
	TimeSlice []*AeronauticalGroundLightTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AeronauticalGroundLightTimeSlicePropertyType Named source named complex type "AeronauticalGroundLightTimeSlicePropertyType"
type AeronauticalGroundLightTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AeronauticalGroundLightTimeSlice" of complex type AeronauticalGroundLightTimeSliceType
	AeronauticalGroundLightTimeSlice *AeronauticalGroundLightTimeSliceType `xml:"AeronauticalGroundLightTimeSlice,omitempty"`
}

// AeronauticalGroundLightTimeSliceType Named source named complex type "AeronauticalGroundLightTimeSliceType"
type AeronauticalGroundLightTimeSliceType struct {
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

	// generated from element "type" of complex type CodeGroundLightingType
	Type_XSD *aixmdatatypes.CodeGroundLightingType `xml:"type,omitempty"`

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "flashing" of complex type CodeYesNoType
	Flashing *aixmdatatypes.CodeYesNoType `xml:"flashing,omitempty"`

	// generated from element "structureBeacon" of complex type VerticalStructurePropertyType
	StructureBeacon *VerticalStructurePropertyType `xml:"structureBeacon,omitempty"`

	// generated from element "aerodromeBeacon" of complex type AirportHeliportPropertyType
	AerodromeBeacon *AirportHeliportPropertyType `xml:"aerodromeBeacon,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAeronauticalGroundLightExtension" of complex type AbstractExtensionType
	AbstractAeronauticalGroundLightExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAeronauticalGroundLightExtension,omitempty"`
}

// Group_AeronauticalGroundLightPropertyGroup UnNamed source named group "AeronauticalGroundLightPropertyGroup"
type Group_AeronauticalGroundLightPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeGroundLightingType
	Type_XSD *aixmdatatypes.CodeGroundLightingType `xml:"type,omitempty"`

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "flashing" of complex type CodeYesNoType
	Flashing *aixmdatatypes.CodeYesNoType `xml:"flashing,omitempty"`

	// generated from element "structureBeacon" of complex type VerticalStructurePropertyType
	StructureBeacon *VerticalStructurePropertyType `xml:"structureBeacon,omitempty"`

	// generated from element "aerodromeBeacon" of complex type AirportHeliportPropertyType
	AerodromeBeacon *AirportHeliportPropertyType `xml:"aerodromeBeacon,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

	// generated from group "aixm:LinguisticNotePropertyGroup" with order 3860 depth 2
	Group_LinguisticNotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
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

	// generated from group "aixm:NotePropertyGroup" with order 3876 depth 2
	Group_NotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
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

// ObstacleAreaType Named source named complex type "ObstacleAreaType"
type ObstacleAreaType struct {
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

	// generated from element "timeSlice" of complex type ObstacleAreaTimeSlicePropertyType
	TimeSlice []*ObstacleAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ObstacleAreaTimeSlicePropertyType Named source named complex type "ObstacleAreaTimeSlicePropertyType"
type ObstacleAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ObstacleAreaTimeSlice" of complex type ObstacleAreaTimeSliceType
	ObstacleAreaTimeSlice *ObstacleAreaTimeSliceType `xml:"ObstacleAreaTimeSlice,omitempty"`
}

// ObstacleAreaTimeSliceType Named source named complex type "ObstacleAreaTimeSliceType"
type ObstacleAreaTimeSliceType struct {
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

	// generated from group "aixm:ObstacleAreaPropertyGroup" with order 3905 depth 2
	Group_ObstacleAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractObstacleAreaExtension" of complex type AbstractExtensionType
	AbstractObstacleAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractObstacleAreaExtension,omitempty"`
}

// Group_ObstacleAreaPropertyGroup UnNamed source named group "ObstacleAreaPropertyGroup"
type Group_ObstacleAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeObstacleAreaType
	Type_XSD *aixmdatatypes.CodeObstacleAreaType `xml:"type,omitempty"`

	// generated from element "obstructionIdSurfaceCondition" of complex type CodeObstacleAssessmentSurfaceType
	ObstructionIdSurfaceCondition *aixmdatatypes.CodeObstacleAssessmentSurfaceType `xml:"obstructionIdSurfaceCondition,omitempty"`

	// generated from element "reference_ownerOrganisation" of complex type OrganisationAuthorityPropertyType
	Reference_ownerOrganisation *OrganisationAuthorityPropertyType `xml:"reference_ownerOrganisation,omitempty"`

	// generated from element "reference_ownerRunway" of complex type RunwayDirectionPropertyType
	Reference_ownerRunway *RunwayDirectionPropertyType `xml:"reference_ownerRunway,omitempty"`

	// generated from element "reference_ownerAirport" of complex type AirportHeliportPropertyType
	Reference_ownerAirport *AirportHeliportPropertyType `xml:"reference_ownerAirport,omitempty"`

	// generated from element "surfaceExtent" of complex type SurfacePropertyType
	SurfaceExtent *SurfacePropertyType `xml:"surfaceExtent,omitempty"`

	// generated from element "obstacle" of complex type VerticalStructurePropertyType
	Obstacle []*VerticalStructurePropertyType `xml:"obstacle,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

// VerticalStructureType Named source named complex type "VerticalStructureType"
type VerticalStructureType struct {
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

	// generated from element "timeSlice" of complex type VerticalStructureTimeSlicePropertyType
	TimeSlice []*VerticalStructureTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// VerticalStructureTimeSlicePropertyType Named source named complex type "VerticalStructureTimeSlicePropertyType"
type VerticalStructureTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "VerticalStructureTimeSlice" of complex type VerticalStructureTimeSliceType
	VerticalStructureTimeSlice *VerticalStructureTimeSliceType `xml:"VerticalStructureTimeSlice,omitempty"`
}

// VerticalStructureTimeSliceType Named source named complex type "VerticalStructureTimeSliceType"
type VerticalStructureTimeSliceType struct {
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

	// generated from element "type" of complex type CodeVerticalStructureType
	Type_XSD *aixmdatatypes.CodeVerticalStructureType `xml:"type,omitempty"`

	// generated from element "lighted" of complex type CodeYesNoType
	Lighted *aixmdatatypes.CodeYesNoType `xml:"lighted,omitempty"`

	// generated from element "markingICAOStandard" of complex type CodeYesNoType
	MarkingICAOStandard *aixmdatatypes.CodeYesNoType `xml:"markingICAOStandard,omitempty"`

	// generated from element "group" of complex type CodeYesNoType
	Group *aixmdatatypes.CodeYesNoType `xml:"group,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "radius" of complex type ValDistanceType
	Radius *aixmdatatypes.ValDistanceType `xml:"radius,omitempty"`

	// generated from element "lightingICAOStandard" of complex type CodeYesNoType
	LightingICAOStandard *aixmdatatypes.CodeYesNoType `xml:"lightingICAOStandard,omitempty"`

	// generated from element "synchronisedLighting" of complex type CodeYesNoType
	SynchronisedLighting *aixmdatatypes.CodeYesNoType `xml:"synchronisedLighting,omitempty"`

	// generated from element "marker" of complex type MarkerBeaconPropertyType
	Marker *MarkerBeaconPropertyType `xml:"marker,omitempty"`

	// generated from element "part" of complex type VerticalStructurePartPropertyType
	Part []*VerticalStructurePartPropertyType `xml:"part,omitempty"`

	// generated from element "hostedPassengerService" of complex type PassengerServicePropertyType
	HostedPassengerService []*PassengerServicePropertyType `xml:"hostedPassengerService,omitempty"`

	// generated from element "supportedGroundLight" of complex type GroundLightSystemPropertyType
	SupportedGroundLight []*GroundLightSystemPropertyType `xml:"supportedGroundLight,omitempty"`

	// generated from element "hostedNavaidEquipment" of complex type NavaidEquipmentPropertyType
	HostedNavaidEquipment []*NavaidEquipmentPropertyType `xml:"hostedNavaidEquipment,omitempty"`

	// generated from element "hostedSpecialNavStation" of complex type SpecialNavigationStationPropertyType
	HostedSpecialNavStation []*SpecialNavigationStationPropertyType `xml:"hostedSpecialNavStation,omitempty"`

	// generated from element "hostedUnit" of complex type UnitPropertyType
	HostedUnit []*UnitPropertyType `xml:"hostedUnit,omitempty"`

	// generated from element "hostedOrganisation" of complex type OrganisationAuthorityPropertyType
	HostedOrganisation []*OrganisationAuthorityPropertyType `xml:"hostedOrganisation,omitempty"`

	// generated from element "supportedService" of complex type ServicePropertyType
	SupportedService []*ServicePropertyType `xml:"supportedService,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "lightingAvailability" of complex type VerticalStructureLightingStatusPropertyType
	LightingAvailability []*VerticalStructureLightingStatusPropertyType `xml:"lightingAvailability,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractVerticalStructureExtension" of complex type AbstractExtensionType
	AbstractVerticalStructureExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractVerticalStructureExtension,omitempty"`
}

// Group_VerticalStructurePropertyGroup UnNamed source named group "VerticalStructurePropertyGroup"
type Group_VerticalStructurePropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeVerticalStructureType
	Type_XSD *aixmdatatypes.CodeVerticalStructureType `xml:"type,omitempty"`

	// generated from element "lighted" of complex type CodeYesNoType
	Lighted *aixmdatatypes.CodeYesNoType `xml:"lighted,omitempty"`

	// generated from element "markingICAOStandard" of complex type CodeYesNoType
	MarkingICAOStandard *aixmdatatypes.CodeYesNoType `xml:"markingICAOStandard,omitempty"`

	// generated from element "group" of complex type CodeYesNoType
	Group *aixmdatatypes.CodeYesNoType `xml:"group,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "radius" of complex type ValDistanceType
	Radius *aixmdatatypes.ValDistanceType `xml:"radius,omitempty"`

	// generated from element "lightingICAOStandard" of complex type CodeYesNoType
	LightingICAOStandard *aixmdatatypes.CodeYesNoType `xml:"lightingICAOStandard,omitempty"`

	// generated from element "synchronisedLighting" of complex type CodeYesNoType
	SynchronisedLighting *aixmdatatypes.CodeYesNoType `xml:"synchronisedLighting,omitempty"`

	// generated from element "marker" of complex type MarkerBeaconPropertyType
	Marker *MarkerBeaconPropertyType `xml:"marker,omitempty"`

	// generated from element "part" of complex type VerticalStructurePartPropertyType
	Part []*VerticalStructurePartPropertyType `xml:"part,omitempty"`

	// generated from element "hostedPassengerService" of complex type PassengerServicePropertyType
	HostedPassengerService []*PassengerServicePropertyType `xml:"hostedPassengerService,omitempty"`

	// generated from element "supportedGroundLight" of complex type GroundLightSystemPropertyType
	SupportedGroundLight []*GroundLightSystemPropertyType `xml:"supportedGroundLight,omitempty"`

	// generated from element "hostedNavaidEquipment" of complex type NavaidEquipmentPropertyType
	HostedNavaidEquipment []*NavaidEquipmentPropertyType `xml:"hostedNavaidEquipment,omitempty"`

	// generated from element "hostedSpecialNavStation" of complex type SpecialNavigationStationPropertyType
	HostedSpecialNavStation []*SpecialNavigationStationPropertyType `xml:"hostedSpecialNavStation,omitempty"`

	// generated from element "hostedUnit" of complex type UnitPropertyType
	HostedUnit []*UnitPropertyType `xml:"hostedUnit,omitempty"`

	// generated from element "hostedOrganisation" of complex type OrganisationAuthorityPropertyType
	HostedOrganisation []*OrganisationAuthorityPropertyType `xml:"hostedOrganisation,omitempty"`

	// generated from element "supportedService" of complex type ServicePropertyType
	SupportedService []*ServicePropertyType `xml:"supportedService,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "lightingAvailability" of complex type VerticalStructureLightingStatusPropertyType
	LightingAvailability []*VerticalStructureLightingStatusPropertyType `xml:"lightingAvailability,omitempty"`
}

// VerticalStructureLightingStatusType Named source named complex type "VerticalStructureLightingStatusType"
type VerticalStructureLightingStatusType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 3970 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:VerticalStructureLightingStatusPropertyGroup" with order 3971 depth 2
	Group_VerticalStructureLightingStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractVerticalStructureLightingStatusExtension" of complex type AbstractExtensionType
	AbstractVerticalStructureLightingStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractVerticalStructureLightingStatusExtension,omitempty"`
}

// Group_VerticalStructureLightingStatusPropertyGroup UnNamed source named group "VerticalStructureLightingStatusPropertyGroup"
type Group_VerticalStructureLightingStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`
}

// VerticalStructureLightingStatusPropertyType Named source named complex type "VerticalStructureLightingStatusPropertyType"
type VerticalStructureLightingStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "VerticalStructureLightingStatus" of complex type VerticalStructureLightingStatusType
	VerticalStructureLightingStatus *VerticalStructureLightingStatusType `xml:"VerticalStructureLightingStatus,omitempty"`
}

// VerticalStructurePartType Named source named complex type "VerticalStructurePartType"
type VerticalStructurePartType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 3988 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:VerticalStructurePartPropertyGroup" with order 3989 depth 2
	Group_VerticalStructurePartPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_9 `xml:"extension,omitempty"`
}

// A_extension_9 Named source within outer element "extension"
type A_extension_9 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractVerticalStructurePartExtension" of complex type AbstractExtensionType
	AbstractVerticalStructurePartExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractVerticalStructurePartExtension,omitempty"`
}

// Group_VerticalStructurePartPropertyGroup UnNamed source named group "VerticalStructurePartPropertyGroup"
type Group_VerticalStructurePartPropertyGroup struct {

	// insertion point for fields

	// generated from element "verticalExtent" of complex type ValDistanceType
	VerticalExtent *aixmdatatypes.ValDistanceType `xml:"verticalExtent,omitempty"`

	// generated from element "verticalExtentAccuracy" of complex type ValDistanceType
	VerticalExtentAccuracy *aixmdatatypes.ValDistanceType `xml:"verticalExtentAccuracy,omitempty"`

	// generated from element "type" of complex type CodeVerticalStructureType
	Type_XSD *aixmdatatypes.CodeVerticalStructureType `xml:"type,omitempty"`

	// generated from element "constructionStatus" of complex type CodeStatusConstructionType
	ConstructionStatus *aixmdatatypes.CodeStatusConstructionType `xml:"constructionStatus,omitempty"`

	// generated from element "markingPattern" of complex type CodeVerticalStructureMarkingType
	MarkingPattern *aixmdatatypes.CodeVerticalStructureMarkingType `xml:"markingPattern,omitempty"`

	// generated from element "markingFirstColour" of complex type CodeColourType
	MarkingFirstColour *aixmdatatypes.CodeColourType `xml:"markingFirstColour,omitempty"`

	// generated from element "markingSecondColour" of complex type CodeColourType
	MarkingSecondColour *aixmdatatypes.CodeColourType `xml:"markingSecondColour,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "frangible" of complex type CodeYesNoType
	Frangible *aixmdatatypes.CodeYesNoType `xml:"frangible,omitempty"`

	// generated from element "visibleMaterial" of complex type CodeVerticalStructureMaterialType
	VisibleMaterial *aixmdatatypes.CodeVerticalStructureMaterialType `xml:"visibleMaterial,omitempty"`

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "horizontalProjection_surfaceExtent" of complex type ElevatedSurfacePropertyType
	HorizontalProjection_surfaceExtent *ElevatedSurfacePropertyType `xml:"horizontalProjection_surfaceExtent,omitempty"`

	// generated from element "horizontalProjection_linearExtent" of complex type ElevatedCurvePropertyType
	HorizontalProjection_linearExtent *ElevatedCurvePropertyType `xml:"horizontalProjection_linearExtent,omitempty"`

	// generated from element "horizontalProjection_location" of complex type ElevatedPointPropertyType
	HorizontalProjection_location *ElevatedPointPropertyType `xml:"horizontalProjection_location,omitempty"`

	// generated from element "lighting" of complex type LightElementPropertyType
	Lighting []*LightElementPropertyType `xml:"lighting,omitempty"`
}

// VerticalStructurePartPropertyType Named source named complex type "VerticalStructurePartPropertyType"
type VerticalStructurePartPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "VerticalStructurePart" of complex type VerticalStructurePartType
	VerticalStructurePart *VerticalStructurePartType `xml:"VerticalStructurePart,omitempty"`
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

// UnitPropertyType Named source named complex type "UnitPropertyType"
type UnitPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// PassengerServicePropertyType Named source named complex type "PassengerServicePropertyType"
type PassengerServicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// ServicePropertyType Named source named complex type "ServicePropertyType"
type ServicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// LightElementType Named source named complex type "LightElementType"
type LightElementType struct {
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

	// generated from group "aixm:LightElementPropertyGroup" with order 6381 depth 2
	Group_LightElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractLightElementExtension" of complex type AbstractExtensionType
	AbstractLightElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLightElementExtension,omitempty"`
}

// Group_LightElementPropertyGroup UnNamed source named group "LightElementPropertyGroup"
type Group_LightElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "intensityLevel" of complex type CodeLightIntensityType
	IntensityLevel *aixmdatatypes.CodeLightIntensityType `xml:"intensityLevel,omitempty"`

	// generated from element "intensity" of complex type ValLightIntensityType
	Intensity *aixmdatatypes.ValLightIntensityType `xml:"intensity,omitempty"`

	// generated from element "type" of complex type CodeLightSourceType
	Type_XSD *aixmdatatypes.CodeLightSourceType `xml:"type,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type LightElementStatusPropertyType
	Availability []*LightElementStatusPropertyType `xml:"availability,omitempty"`
}

// LightElementPropertyType Named source named complex type "LightElementPropertyType"
type LightElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "LightElement" of complex type LightElementType
	LightElement *LightElementType `xml:"LightElement,omitempty"`
}

// LightElementStatusType Named source named complex type "LightElementStatusType"
type LightElementStatusType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6403 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:LightElementStatusPropertyGroup" with order 6404 depth 2
	Group_LightElementStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractLightElementStatusExtension" of complex type AbstractExtensionType
	AbstractLightElementStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLightElementStatusExtension,omitempty"`
}

// Group_LightElementStatusPropertyGroup UnNamed source named group "LightElementStatusPropertyGroup"
type Group_LightElementStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`
}

// LightElementStatusPropertyType Named source named complex type "LightElementStatusPropertyType"
type LightElementStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "LightElementStatus" of complex type LightElementStatusType
	LightElementStatus *LightElementStatusType `xml:"LightElementStatus,omitempty"`
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

	// generated from group "aixm:TimesheetPropertyGroup" with order 6518 depth 2
	Group_TimesheetPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
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
