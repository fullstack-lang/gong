// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAircraftGroundServiceType
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

// GroundLightSystemPropertyType Named source named complex type "GroundLightSystemPropertyType"
type GroundLightSystemPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// LightActivationType Named source named complex type "LightActivationType"
type LightActivationType struct {
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

	// generated from group "aixm:LightActivationPropertyGroup" with order 1326 depth 2
	Group_LightActivationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension `xml:"extension,omitempty"`
}

// A_extension Named source within outer element "extension"
type A_extension struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractLightActivationExtension" of complex type AbstractExtensionType
	AbstractLightActivationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLightActivationExtension,omitempty"`
}

// Group_LightActivationPropertyGroup UnNamed source named group "LightActivationPropertyGroup"
type Group_LightActivationPropertyGroup struct {

	// insertion point for fields

	// generated from element "clicks" of complex type NoNumberType
	Clicks *aixmdatatypes.NoNumberType `xml:"clicks,omitempty"`

	// generated from element "intensityLevel" of complex type CodeLightIntensityType
	IntensityLevel *aixmdatatypes.CodeLightIntensityType `xml:"intensityLevel,omitempty"`

	// generated from element "activation" of complex type CodeSystemActivationType
	Activation *aixmdatatypes.CodeSystemActivationType `xml:"activation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// LightActivationPropertyType Named source named complex type "LightActivationPropertyType"
type LightActivationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "LightActivation" of complex type LightActivationType
	LightActivation *LightActivationType `xml:"LightActivation,omitempty"`
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

// HoldingPatternPropertyType Named source named complex type "HoldingPatternPropertyType"
type HoldingPatternPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// DirectionFinderPropertyType Named source named complex type "DirectionFinderPropertyType"
type DirectionFinderPropertyType struct {
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

// VORPropertyType Named source named complex type "VORPropertyType"
type VORPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
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
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
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
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
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

// OrganisationAuthorityPropertyType Named source named complex type "OrganisationAuthorityPropertyType"
type OrganisationAuthorityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// OrganisationAuthorityType Named source named complex type "OrganisationAuthorityType"
type OrganisationAuthorityType struct {
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

	// generated from element "timeSlice" of complex type OrganisationAuthorityTimeSlicePropertyType
	TimeSlice []*OrganisationAuthorityTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// OrganisationAuthorityTimeSlicePropertyType Named source named complex type "OrganisationAuthorityTimeSlicePropertyType"
type OrganisationAuthorityTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "OrganisationAuthorityTimeSlice" of complex type OrganisationAuthorityTimeSliceType
	OrganisationAuthorityTimeSlice *OrganisationAuthorityTimeSliceType `xml:"OrganisationAuthorityTimeSlice,omitempty"`
}

// OrganisationAuthorityTimeSliceType Named source named complex type "OrganisationAuthorityTimeSliceType"
type OrganisationAuthorityTimeSliceType struct {
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

	// generated from element "designator" of complex type CodeOrganisationDesignatorType
	Designator *aixmdatatypes.CodeOrganisationDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeOrganisationType
	Type_XSD *aixmdatatypes.CodeOrganisationType `xml:"type,omitempty"`

	// generated from element "military" of complex type CodeMilitaryOperationsType
	Military *aixmdatatypes.CodeMilitaryOperationsType `xml:"military,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "relatedOrganisationAuthority" of complex type OrganisationAuthorityAssociationPropertyType
	RelatedOrganisationAuthority []*OrganisationAuthorityAssociationPropertyType `xml:"relatedOrganisationAuthority,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractOrganisationAuthorityExtension" of complex type AbstractExtensionType
	AbstractOrganisationAuthorityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractOrganisationAuthorityExtension,omitempty"`
}

// Group_OrganisationAuthorityPropertyGroup UnNamed source named group "OrganisationAuthorityPropertyGroup"
type Group_OrganisationAuthorityPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "designator" of complex type CodeOrganisationDesignatorType
	Designator *aixmdatatypes.CodeOrganisationDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeOrganisationType
	Type_XSD *aixmdatatypes.CodeOrganisationType `xml:"type,omitempty"`

	// generated from element "military" of complex type CodeMilitaryOperationsType
	Military *aixmdatatypes.CodeMilitaryOperationsType `xml:"military,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "relatedOrganisationAuthority" of complex type OrganisationAuthorityAssociationPropertyType
	RelatedOrganisationAuthority []*OrganisationAuthorityAssociationPropertyType `xml:"relatedOrganisationAuthority,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// OrganisationAuthorityAssociationType Named source named complex type "OrganisationAuthorityAssociationType"
type OrganisationAuthorityAssociationType struct {
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

	// generated from group "aixm:OrganisationAuthorityAssociationPropertyGroup" with order 4067 depth 2
	Group_OrganisationAuthorityAssociationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractOrganisationAuthorityAssociationExtension" of complex type AbstractExtensionType
	AbstractOrganisationAuthorityAssociationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractOrganisationAuthorityAssociationExtension,omitempty"`
}

// Group_OrganisationAuthorityAssociationPropertyGroup UnNamed source named group "OrganisationAuthorityAssociationPropertyGroup"
type Group_OrganisationAuthorityAssociationPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeOrganisationHierarchyType
	Type_XSD *aixmdatatypes.CodeOrganisationHierarchyType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// OrganisationAuthorityAssociationPropertyType Named source named complex type "OrganisationAuthorityAssociationPropertyType"
type OrganisationAuthorityAssociationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "OrganisationAuthorityAssociation" of complex type OrganisationAuthorityAssociationType
	OrganisationAuthorityAssociation *OrganisationAuthorityAssociationType `xml:"OrganisationAuthorityAssociation,omitempty"`
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

// UnitType Named source named complex type "UnitType"
type UnitType struct {
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

	// generated from element "timeSlice" of complex type UnitTimeSlicePropertyType
	TimeSlice []*UnitTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// UnitTimeSlicePropertyType Named source named complex type "UnitTimeSlicePropertyType"
type UnitTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "UnitTimeSlice" of complex type UnitTimeSliceType
	UnitTimeSlice *UnitTimeSliceType `xml:"UnitTimeSlice,omitempty"`
}

// UnitTimeSliceType Named source named complex type "UnitTimeSliceType"
type UnitTimeSliceType struct {
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

	// generated from element "type" of complex type CodeUnitType
	Type_XSD *aixmdatatypes.CodeUnitType `xml:"type,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "designator" of complex type CodeOrganisationDesignatorType
	Designator *aixmdatatypes.CodeOrganisationDesignatorType `xml:"designator,omitempty"`

	// generated from element "military" of complex type CodeMilitaryOperationsType
	Military *aixmdatatypes.CodeMilitaryOperationsType `xml:"military,omitempty"`

	// generated from element "position" of complex type ElevatedPointPropertyType
	Position *ElevatedPointPropertyType `xml:"position,omitempty"`

	// generated from element "airportLocation" of complex type AirportHeliportPropertyType
	AirportLocation *AirportHeliportPropertyType `xml:"airportLocation,omitempty"`

	// generated from element "ownerOrganisation" of complex type OrganisationAuthorityPropertyType
	OwnerOrganisation *OrganisationAuthorityPropertyType `xml:"ownerOrganisation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "relatedUnit" of complex type UnitDependencyPropertyType
	RelatedUnit []*UnitDependencyPropertyType `xml:"relatedUnit,omitempty"`

	// generated from element "availability" of complex type UnitAvailabilityPropertyType
	Availability []*UnitAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUnitExtension" of complex type AbstractExtensionType
	AbstractUnitExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUnitExtension,omitempty"`
}

// Group_UnitPropertyGroup UnNamed source named group "UnitPropertyGroup"
type Group_UnitPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeUnitType
	Type_XSD *aixmdatatypes.CodeUnitType `xml:"type,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "designator" of complex type CodeOrganisationDesignatorType
	Designator *aixmdatatypes.CodeOrganisationDesignatorType `xml:"designator,omitempty"`

	// generated from element "military" of complex type CodeMilitaryOperationsType
	Military *aixmdatatypes.CodeMilitaryOperationsType `xml:"military,omitempty"`

	// generated from element "position" of complex type ElevatedPointPropertyType
	Position *ElevatedPointPropertyType `xml:"position,omitempty"`

	// generated from element "airportLocation" of complex type AirportHeliportPropertyType
	AirportLocation *AirportHeliportPropertyType `xml:"airportLocation,omitempty"`

	// generated from element "ownerOrganisation" of complex type OrganisationAuthorityPropertyType
	OwnerOrganisation *OrganisationAuthorityPropertyType `xml:"ownerOrganisation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "relatedUnit" of complex type UnitDependencyPropertyType
	RelatedUnit []*UnitDependencyPropertyType `xml:"relatedUnit,omitempty"`

	// generated from element "availability" of complex type UnitAvailabilityPropertyType
	Availability []*UnitAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// UnitAvailabilityType Named source named complex type "UnitAvailabilityType"
type UnitAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 4120 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:UnitAvailabilityPropertyGroup" with order 4121 depth 2
	Group_UnitAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractUnitAvailabilityExtension" of complex type AbstractExtensionType
	AbstractUnitAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUnitAvailabilityExtension,omitempty"`
}

// Group_UnitAvailabilityPropertyGroup UnNamed source named group "UnitAvailabilityPropertyGroup"
type Group_UnitAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusOperationsType
	OperationalStatus *aixmdatatypes.CodeStatusOperationsType `xml:"operationalStatus,omitempty"`
}

// UnitAvailabilityPropertyType Named source named complex type "UnitAvailabilityPropertyType"
type UnitAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "UnitAvailability" of complex type UnitAvailabilityType
	UnitAvailability *UnitAvailabilityType `xml:"UnitAvailability,omitempty"`
}

// UnitDependencyType Named source named complex type "UnitDependencyType"
type UnitDependencyType struct {
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

	// generated from group "aixm:UnitDependencyPropertyGroup" with order 4138 depth 2
	Group_UnitDependencyPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUnitDependencyExtension" of complex type AbstractExtensionType
	AbstractUnitDependencyExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUnitDependencyExtension,omitempty"`
}

// Group_UnitDependencyPropertyGroup UnNamed source named group "UnitDependencyPropertyGroup"
type Group_UnitDependencyPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeUnitDependencyType
	Type_XSD *aixmdatatypes.CodeUnitDependencyType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theUnit" of complex type UnitPropertyType
	TheUnit *UnitPropertyType `xml:"theUnit,omitempty"`
}

// UnitDependencyPropertyType Named source named complex type "UnitDependencyPropertyType"
type UnitDependencyPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "UnitDependency" of complex type UnitDependencyType
	UnitDependency *UnitDependencyType `xml:"UnitDependency,omitempty"`
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

// RoutePropertyType Named source named complex type "RoutePropertyType"
type RoutePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
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

	// generated from group "aixm:RoutePortionPropertyGroup" with order 5294 depth 2
	Group_RoutePortionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_9 `xml:"extension,omitempty"`
}

// A_extension_9 Named source within outer element "extension"
type A_extension_9 struct {
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

// AircraftGroundServiceType Named source named complex type "AircraftGroundServiceType"
type AircraftGroundServiceType struct {
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

	// generated from element "timeSlice" of complex type AircraftGroundServiceTimeSlicePropertyType
	TimeSlice []*AircraftGroundServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AircraftGroundServiceTimeSlicePropertyType Named source named complex type "AircraftGroundServiceTimeSlicePropertyType"
type AircraftGroundServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AircraftGroundServiceTimeSlice" of complex type AircraftGroundServiceTimeSliceType
	AircraftGroundServiceTimeSlice *AircraftGroundServiceTimeSliceType `xml:"AircraftGroundServiceTimeSlice,omitempty"`
}

// AircraftGroundServiceTimeSliceType Named source named complex type "AircraftGroundServiceTimeSliceType"
type AircraftGroundServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "type" of complex type CodeAircraftGroundServiceType
	Type_XSD *aixmdatatypes.CodeAircraftGroundServiceType `xml:"type,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAircraftGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAircraftGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAircraftGroundServiceExtension,omitempty"`

	// generated from element "AbstractAirportGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAirportGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportGroundServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_AircraftGroundServicePropertyGroup UnNamed source named group "AircraftGroundServicePropertyGroup"
type Group_AircraftGroundServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportGroundServicePropertyGroup" with order 5702 depth 2
	Group_AirportGroundServicePropertyGroup

	// generated from element "type" of complex type CodeAircraftGroundServiceType
	Type_XSD *aixmdatatypes.CodeAircraftGroundServiceType `xml:"type,omitempty"`
}

// AirportClearanceServiceType Named source named complex type "AirportClearanceServiceType"
type AirportClearanceServiceType struct {
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

	// generated from element "timeSlice" of complex type AirportClearanceServiceTimeSlicePropertyType
	TimeSlice []*AirportClearanceServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportClearanceServiceTimeSlicePropertyType Named source named complex type "AirportClearanceServiceTimeSlicePropertyType"
type AirportClearanceServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportClearanceServiceTimeSlice" of complex type AirportClearanceServiceTimeSliceType
	AirportClearanceServiceTimeSlice *AirportClearanceServiceTimeSliceType `xml:"AirportClearanceServiceTimeSlice,omitempty"`
}

// AirportClearanceServiceTimeSliceType Named source named complex type "AirportClearanceServiceTimeSliceType"
type AirportClearanceServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "snowPlan" of complex type TextInstructionType
	SnowPlan *aixmdatatypes.TextInstructionType `xml:"snowPlan,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirportClearanceServiceExtension" of complex type AbstractExtensionType
	AbstractAirportClearanceServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportClearanceServiceExtension,omitempty"`

	// generated from element "AbstractAirportGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAirportGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportGroundServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_AirportClearanceServicePropertyGroup UnNamed source named group "AirportClearanceServicePropertyGroup"
type Group_AirportClearanceServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportGroundServicePropertyGroup" with order 5729 depth 2
	Group_AirportGroundServicePropertyGroup

	// generated from element "snowPlan" of complex type TextInstructionType
	SnowPlan *aixmdatatypes.TextInstructionType `xml:"snowPlan,omitempty"`
}

// AbstractAirportGroundServiceType Named source named complex type "AbstractAirportGroundServiceType"
type AbstractAirportGroundServiceType struct {
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

// Group_AirportGroundServicePropertyGroup UnNamed source named group "AirportGroundServicePropertyGroup"
type Group_AirportGroundServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ServicePropertyGroup" with order 5739 depth 2
	Group_ServicePropertyGroup

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`
}

// AirportSuppliesServiceType Named source named complex type "AirportSuppliesServiceType"
type AirportSuppliesServiceType struct {
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

	// generated from element "timeSlice" of complex type AirportSuppliesServiceTimeSlicePropertyType
	TimeSlice []*AirportSuppliesServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportSuppliesServiceTimeSlicePropertyType Named source named complex type "AirportSuppliesServiceTimeSlicePropertyType"
type AirportSuppliesServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportSuppliesServiceTimeSlice" of complex type AirportSuppliesServiceTimeSliceType
	AirportSuppliesServiceTimeSlice *AirportSuppliesServiceTimeSliceType `xml:"AirportSuppliesServiceTimeSlice,omitempty"`
}

// AirportSuppliesServiceTimeSliceType Named source named complex type "AirportSuppliesServiceTimeSliceType"
type AirportSuppliesServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "fuelSupply" of complex type FuelPropertyType
	FuelSupply []*FuelPropertyType `xml:"fuelSupply,omitempty"`

	// generated from element "oilSupply" of complex type OilPropertyType
	OilSupply []*OilPropertyType `xml:"oilSupply,omitempty"`

	// generated from element "nitrogenSupply" of complex type NitrogenPropertyType
	NitrogenSupply []*NitrogenPropertyType `xml:"nitrogenSupply,omitempty"`

	// generated from element "oxygenSupply" of complex type OxygenPropertyType
	OxygenSupply []*OxygenPropertyType `xml:"oxygenSupply,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirportSuppliesServiceExtension" of complex type AbstractExtensionType
	AbstractAirportSuppliesServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportSuppliesServiceExtension,omitempty"`

	// generated from element "AbstractAirportGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAirportGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportGroundServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_AirportSuppliesServicePropertyGroup UnNamed source named group "AirportSuppliesServicePropertyGroup"
type Group_AirportSuppliesServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportGroundServicePropertyGroup" with order 5766 depth 2
	Group_AirportGroundServicePropertyGroup

	// generated from element "fuelSupply" of complex type FuelPropertyType
	FuelSupply []*FuelPropertyType `xml:"fuelSupply,omitempty"`

	// generated from element "oilSupply" of complex type OilPropertyType
	OilSupply []*OilPropertyType `xml:"oilSupply,omitempty"`

	// generated from element "nitrogenSupply" of complex type NitrogenPropertyType
	NitrogenSupply []*NitrogenPropertyType `xml:"nitrogenSupply,omitempty"`

	// generated from element "oxygenSupply" of complex type OxygenPropertyType
	OxygenSupply []*OxygenPropertyType `xml:"oxygenSupply,omitempty"`
}

// AirTrafficControlServiceType Named source named complex type "AirTrafficControlServiceType"
type AirTrafficControlServiceType struct {
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

	// generated from element "timeSlice" of complex type AirTrafficControlServiceTimeSlicePropertyType
	TimeSlice []*AirTrafficControlServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirTrafficControlServiceTimeSlicePropertyType Named source named complex type "AirTrafficControlServiceTimeSlicePropertyType"
type AirTrafficControlServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirTrafficControlServiceTimeSlice" of complex type AirTrafficControlServiceTimeSliceType
	AirTrafficControlServiceTimeSlice *AirTrafficControlServiceTimeSliceType `xml:"AirTrafficControlServiceTimeSlice,omitempty"`
}

// AirTrafficControlServiceTimeSliceType Named source named complex type "AirTrafficControlServiceTimeSliceType"
type AirTrafficControlServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "radarAssisted" of complex type CodeYesNoType
	RadarAssisted *aixmdatatypes.CodeYesNoType `xml:"radarAssisted,omitempty"`

	// generated from element "dataLinkEnabled" of complex type CodeYesNoType
	DataLinkEnabled *aixmdatatypes.CodeYesNoType `xml:"dataLinkEnabled,omitempty"`

	// generated from element "dataLinkChannel" of complex type CodeCommunicationChannelType
	DataLinkChannel *aixmdatatypes.CodeCommunicationChannelType `xml:"dataLinkChannel,omitempty"`

	// generated from element "type" of complex type CodeServiceATCType
	Type_XSD *aixmdatatypes.CodeServiceATCType `xml:"type,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport []*AirportHeliportPropertyType `xml:"clientAirport,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from element "clientProcedure" of complex type ProcedurePropertyType
	ClientProcedure []*ProcedurePropertyType `xml:"clientProcedure,omitempty"`

	// generated from element "clientHolding" of complex type HoldingPatternPropertyType
	ClientHolding []*HoldingPatternPropertyType `xml:"clientHolding,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`

	// generated from element "aircraftLocator" of complex type DirectionFinderPropertyType
	AircraftLocator *DirectionFinderPropertyType `xml:"aircraftLocator,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirTrafficControlServiceExtension" of complex type AbstractExtensionType
	AbstractAirTrafficControlServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirTrafficControlServiceExtension,omitempty"`

	// generated from element "AbstractTrafficSeparationServiceExtension" of complex type AbstractExtensionType
	AbstractTrafficSeparationServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTrafficSeparationServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_AirTrafficControlServicePropertyGroup UnNamed source named group "AirTrafficControlServicePropertyGroup"
type Group_AirTrafficControlServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:TrafficSeparationServicePropertyGroup" with order 5796 depth 2
	Group_TrafficSeparationServicePropertyGroup

	// generated from element "type" of complex type CodeServiceATCType
	Type_XSD *aixmdatatypes.CodeServiceATCType `xml:"type,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport []*AirportHeliportPropertyType `xml:"clientAirport,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from element "clientProcedure" of complex type ProcedurePropertyType
	ClientProcedure []*ProcedurePropertyType `xml:"clientProcedure,omitempty"`

	// generated from element "clientHolding" of complex type HoldingPatternPropertyType
	ClientHolding []*HoldingPatternPropertyType `xml:"clientHolding,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`

	// generated from element "aircraftLocator" of complex type DirectionFinderPropertyType
	AircraftLocator *DirectionFinderPropertyType `xml:"aircraftLocator,omitempty"`
}

// AirTrafficManagementServiceType Named source named complex type "AirTrafficManagementServiceType"
type AirTrafficManagementServiceType struct {
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

	// generated from element "timeSlice" of complex type AirTrafficManagementServiceTimeSlicePropertyType
	TimeSlice []*AirTrafficManagementServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirTrafficManagementServiceTimeSlicePropertyType Named source named complex type "AirTrafficManagementServiceTimeSlicePropertyType"
type AirTrafficManagementServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirTrafficManagementServiceTimeSlice" of complex type AirTrafficManagementServiceTimeSliceType
	AirTrafficManagementServiceTimeSlice *AirTrafficManagementServiceTimeSliceType `xml:"AirTrafficManagementServiceTimeSlice,omitempty"`
}

// AirTrafficManagementServiceTimeSliceType Named source named complex type "AirTrafficManagementServiceTimeSliceType"
type AirTrafficManagementServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeServiceATFMType
	Type_XSD *aixmdatatypes.CodeServiceATFMType `xml:"type,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirTrafficManagementServiceExtension" of complex type AbstractExtensionType
	AbstractAirTrafficManagementServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirTrafficManagementServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_AirTrafficManagementServicePropertyGroup UnNamed source named group "AirTrafficManagementServicePropertyGroup"
type Group_AirTrafficManagementServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ServicePropertyGroup" with order 5829 depth 2
	Group_ServicePropertyGroup

	// generated from element "type" of complex type CodeServiceATFMType
	Type_XSD *aixmdatatypes.CodeServiceATFMType `xml:"type,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`
}

// CallsignDetailType Named source named complex type "CallsignDetailType"
type CallsignDetailType struct {
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

	// generated from group "aixm:CallsignDetailPropertyGroup" with order 5838 depth 2
	Group_CallsignDetailPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_15 `xml:"extension,omitempty"`
}

// A_extension_15 Named source within outer element "extension"
type A_extension_15 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCallsignDetailExtension" of complex type AbstractExtensionType
	AbstractCallsignDetailExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCallsignDetailExtension,omitempty"`
}

// Group_CallsignDetailPropertyGroup UnNamed source named group "CallsignDetailPropertyGroup"
type Group_CallsignDetailPropertyGroup struct {

	// insertion point for fields

	// generated from element "callSign" of complex type TextNameType
	CallSign *aixmdatatypes.TextNameType `xml:"callSign,omitempty"`

	// generated from element "language" of complex type CodeLanguageType
	Language *aixmdatatypes.CodeLanguageType `xml:"language,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// CallsignDetailPropertyType Named source named complex type "CallsignDetailPropertyType"
type CallsignDetailPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "CallsignDetail" of complex type CallsignDetailType
	CallsignDetail *CallsignDetailType `xml:"CallsignDetail,omitempty"`
}

// FireFightingServiceType Named source named complex type "FireFightingServiceType"
type FireFightingServiceType struct {
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

	// generated from element "timeSlice" of complex type FireFightingServiceTimeSlicePropertyType
	TimeSlice []*FireFightingServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// FireFightingServiceTimeSlicePropertyType Named source named complex type "FireFightingServiceTimeSlicePropertyType"
type FireFightingServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "FireFightingServiceTimeSlice" of complex type FireFightingServiceTimeSliceType
	FireFightingServiceTimeSlice *FireFightingServiceTimeSliceType `xml:"FireFightingServiceTimeSlice,omitempty"`
}

// FireFightingServiceTimeSliceType Named source named complex type "FireFightingServiceTimeSliceType"
type FireFightingServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "category" of complex type CodeFireFightingType
	Category *aixmdatatypes.CodeFireFightingType `xml:"category,omitempty"`

	// generated from element "standard" of complex type CodeAviationStandardsType
	Standard *aixmdatatypes.CodeAviationStandardsType `xml:"standard,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_16 `xml:"extension,omitempty"`
}

// A_extension_16 Named source within outer element "extension"
type A_extension_16 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFireFightingServiceExtension" of complex type AbstractExtensionType
	AbstractFireFightingServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFireFightingServiceExtension,omitempty"`

	// generated from element "AbstractAirportGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAirportGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportGroundServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_FireFightingServicePropertyGroup UnNamed source named group "FireFightingServicePropertyGroup"
type Group_FireFightingServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportGroundServicePropertyGroup" with order 5877 depth 2
	Group_AirportGroundServicePropertyGroup

	// generated from element "category" of complex type CodeFireFightingType
	Category *aixmdatatypes.CodeFireFightingType `xml:"category,omitempty"`

	// generated from element "standard" of complex type CodeAviationStandardsType
	Standard *aixmdatatypes.CodeAviationStandardsType `xml:"standard,omitempty"`
}

// FuelType Named source named complex type "FuelType"
type FuelType struct {
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

	// generated from group "aixm:FuelPropertyGroup" with order 5884 depth 2
	Group_FuelPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_17 `xml:"extension,omitempty"`
}

// A_extension_17 Named source within outer element "extension"
type A_extension_17 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFuelExtension" of complex type AbstractExtensionType
	AbstractFuelExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFuelExtension,omitempty"`
}

// Group_FuelPropertyGroup UnNamed source named group "FuelPropertyGroup"
type Group_FuelPropertyGroup struct {

	// insertion point for fields

	// generated from element "category" of complex type CodeFuelType
	Category *aixmdatatypes.CodeFuelType `xml:"category,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// FuelPropertyType Named source named complex type "FuelPropertyType"
type FuelPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Fuel" of complex type FuelType
	Fuel *FuelType `xml:"Fuel,omitempty"`
}

// GroundTrafficControlServiceType Named source named complex type "GroundTrafficControlServiceType"
type GroundTrafficControlServiceType struct {
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

	// generated from element "timeSlice" of complex type GroundTrafficControlServiceTimeSlicePropertyType
	TimeSlice []*GroundTrafficControlServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GroundTrafficControlServiceTimeSlicePropertyType Named source named complex type "GroundTrafficControlServiceTimeSlicePropertyType"
type GroundTrafficControlServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GroundTrafficControlServiceTimeSlice" of complex type GroundTrafficControlServiceTimeSliceType
	GroundTrafficControlServiceTimeSlice *GroundTrafficControlServiceTimeSliceType `xml:"GroundTrafficControlServiceTimeSlice,omitempty"`
}

// GroundTrafficControlServiceTimeSliceType Named source named complex type "GroundTrafficControlServiceTimeSliceType"
type GroundTrafficControlServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "radarAssisted" of complex type CodeYesNoType
	RadarAssisted *aixmdatatypes.CodeYesNoType `xml:"radarAssisted,omitempty"`

	// generated from element "dataLinkEnabled" of complex type CodeYesNoType
	DataLinkEnabled *aixmdatatypes.CodeYesNoType `xml:"dataLinkEnabled,omitempty"`

	// generated from element "dataLinkChannel" of complex type CodeCommunicationChannelType
	DataLinkChannel *aixmdatatypes.CodeCommunicationChannelType `xml:"dataLinkChannel,omitempty"`

	// generated from element "type" of complex type CodeServiceGroundControlType
	Type_XSD *aixmdatatypes.CodeServiceGroundControlType `xml:"type,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport *AirportHeliportPropertyType `xml:"clientAirport,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_18 `xml:"extension,omitempty"`
}

// A_extension_18 Named source within outer element "extension"
type A_extension_18 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGroundTrafficControlServiceExtension" of complex type AbstractExtensionType
	AbstractGroundTrafficControlServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundTrafficControlServiceExtension,omitempty"`

	// generated from element "AbstractTrafficSeparationServiceExtension" of complex type AbstractExtensionType
	AbstractTrafficSeparationServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTrafficSeparationServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_GroundTrafficControlServicePropertyGroup UnNamed source named group "GroundTrafficControlServicePropertyGroup"
type Group_GroundTrafficControlServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:TrafficSeparationServicePropertyGroup" with order 5922 depth 2
	Group_TrafficSeparationServicePropertyGroup

	// generated from element "type" of complex type CodeServiceGroundControlType
	Type_XSD *aixmdatatypes.CodeServiceGroundControlType `xml:"type,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport *AirportHeliportPropertyType `xml:"clientAirport,omitempty"`
}

// InformationServiceType Named source named complex type "InformationServiceType"
type InformationServiceType struct {
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

	// generated from element "timeSlice" of complex type InformationServiceTimeSlicePropertyType
	TimeSlice []*InformationServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// InformationServiceTimeSlicePropertyType Named source named complex type "InformationServiceTimeSlicePropertyType"
type InformationServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "InformationServiceTimeSlice" of complex type InformationServiceTimeSliceType
	InformationServiceTimeSlice *InformationServiceTimeSliceType `xml:"InformationServiceTimeSlice,omitempty"`
}

// InformationServiceTimeSliceType Named source named complex type "InformationServiceTimeSliceType"
type InformationServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeServiceInformationType
	Type_XSD *aixmdatatypes.CodeServiceInformationType `xml:"type,omitempty"`

	// generated from element "voice" of complex type CodeYesNoType
	Voice *aixmdatatypes.CodeYesNoType `xml:"voice,omitempty"`

	// generated from element "dataLink" of complex type CodeYesNoType
	DataLink *aixmdatatypes.CodeYesNoType `xml:"dataLink,omitempty"`

	// generated from element "recorded" of complex type CodeYesNoType
	Recorded *aixmdatatypes.CodeYesNoType `xml:"recorded,omitempty"`

	// generated from element "navaidBroadcast" of complex type VORPropertyType
	NavaidBroadcast []*VORPropertyType `xml:"navaidBroadcast,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport []*AirportHeliportPropertyType `xml:"clientAirport,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from element "clientProcedure" of complex type ProcedurePropertyType
	ClientProcedure []*ProcedurePropertyType `xml:"clientProcedure,omitempty"`

	// generated from element "clientHolding" of complex type HoldingPatternPropertyType
	ClientHolding []*HoldingPatternPropertyType `xml:"clientHolding,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_19 `xml:"extension,omitempty"`
}

// A_extension_19 Named source within outer element "extension"
type A_extension_19 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractInformationServiceExtension" of complex type AbstractExtensionType
	AbstractInformationServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractInformationServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_InformationServicePropertyGroup UnNamed source named group "InformationServicePropertyGroup"
type Group_InformationServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ServicePropertyGroup" with order 5949 depth 2
	Group_ServicePropertyGroup

	// generated from element "type" of complex type CodeServiceInformationType
	Type_XSD *aixmdatatypes.CodeServiceInformationType `xml:"type,omitempty"`

	// generated from element "voice" of complex type CodeYesNoType
	Voice *aixmdatatypes.CodeYesNoType `xml:"voice,omitempty"`

	// generated from element "dataLink" of complex type CodeYesNoType
	DataLink *aixmdatatypes.CodeYesNoType `xml:"dataLink,omitempty"`

	// generated from element "recorded" of complex type CodeYesNoType
	Recorded *aixmdatatypes.CodeYesNoType `xml:"recorded,omitempty"`

	// generated from element "navaidBroadcast" of complex type VORPropertyType
	NavaidBroadcast []*VORPropertyType `xml:"navaidBroadcast,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientAirport" of complex type AirportHeliportPropertyType
	ClientAirport []*AirportHeliportPropertyType `xml:"clientAirport,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from element "clientProcedure" of complex type ProcedurePropertyType
	ClientProcedure []*ProcedurePropertyType `xml:"clientProcedure,omitempty"`

	// generated from element "clientHolding" of complex type HoldingPatternPropertyType
	ClientHolding []*HoldingPatternPropertyType `xml:"clientHolding,omitempty"`

	// generated from element "clientAerialRefuelling" of complex type AerialRefuellingPropertyType
	ClientAerialRefuelling []*AerialRefuellingPropertyType `xml:"clientAerialRefuelling,omitempty"`
}

// NitrogenType Named source named complex type "NitrogenType"
type NitrogenType struct {
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

	// generated from group "aixm:NitrogenPropertyGroup" with order 5965 depth 2
	Group_NitrogenPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_20 `xml:"extension,omitempty"`
}

// A_extension_20 Named source within outer element "extension"
type A_extension_20 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNitrogenExtension" of complex type AbstractExtensionType
	AbstractNitrogenExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNitrogenExtension,omitempty"`
}

// Group_NitrogenPropertyGroup UnNamed source named group "NitrogenPropertyGroup"
type Group_NitrogenPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeNitrogenType
	Type_XSD *aixmdatatypes.CodeNitrogenType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NitrogenPropertyType Named source named complex type "NitrogenPropertyType"
type NitrogenPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Nitrogen" of complex type NitrogenType
	Nitrogen *NitrogenType `xml:"Nitrogen,omitempty"`
}

// OilType Named source named complex type "OilType"
type OilType struct {
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

	// generated from group "aixm:OilPropertyGroup" with order 5982 depth 2
	Group_OilPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_21 `xml:"extension,omitempty"`
}

// A_extension_21 Named source within outer element "extension"
type A_extension_21 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractOilExtension" of complex type AbstractExtensionType
	AbstractOilExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractOilExtension,omitempty"`
}

// Group_OilPropertyGroup UnNamed source named group "OilPropertyGroup"
type Group_OilPropertyGroup struct {

	// insertion point for fields

	// generated from element "category" of complex type CodeOilType
	Category *aixmdatatypes.CodeOilType `xml:"category,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// OilPropertyType Named source named complex type "OilPropertyType"
type OilPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Oil" of complex type OilType
	Oil *OilType `xml:"Oil,omitempty"`
}

// OxygenType Named source named complex type "OxygenType"
type OxygenType struct {
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

	// generated from group "aixm:OxygenPropertyGroup" with order 5999 depth 2
	Group_OxygenPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_22 `xml:"extension,omitempty"`
}

// A_extension_22 Named source within outer element "extension"
type A_extension_22 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractOxygenExtension" of complex type AbstractExtensionType
	AbstractOxygenExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractOxygenExtension,omitempty"`
}

// Group_OxygenPropertyGroup UnNamed source named group "OxygenPropertyGroup"
type Group_OxygenPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeOxygenType
	Type_XSD *aixmdatatypes.CodeOxygenType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// OxygenPropertyType Named source named complex type "OxygenPropertyType"
type OxygenPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Oxygen" of complex type OxygenType
	Oxygen *OxygenType `xml:"Oxygen,omitempty"`
}

// PassengerServiceType Named source named complex type "PassengerServiceType"
type PassengerServiceType struct {
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

	// generated from element "timeSlice" of complex type PassengerServiceTimeSlicePropertyType
	TimeSlice []*PassengerServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// PassengerServiceTimeSlicePropertyType Named source named complex type "PassengerServiceTimeSlicePropertyType"
type PassengerServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "PassengerServiceTimeSlice" of complex type PassengerServiceTimeSliceType
	PassengerServiceTimeSlice *PassengerServiceTimeSliceType `xml:"PassengerServiceTimeSlice,omitempty"`
}

// PassengerServiceTimeSliceType Named source named complex type "PassengerServiceTimeSliceType"
type PassengerServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport []*AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "type" of complex type CodePassengerServiceType
	Type_XSD *aixmdatatypes.CodePassengerServiceType `xml:"type,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_23 `xml:"extension,omitempty"`
}

// A_extension_23 Named source within outer element "extension"
type A_extension_23 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPassengerServiceExtension" of complex type AbstractExtensionType
	AbstractPassengerServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPassengerServiceExtension,omitempty"`

	// generated from element "AbstractAirportGroundServiceExtension" of complex type AbstractExtensionType
	AbstractAirportGroundServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportGroundServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_PassengerServicePropertyGroup UnNamed source named group "PassengerServicePropertyGroup"
type Group_PassengerServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportGroundServicePropertyGroup" with order 6037 depth 2
	Group_AirportGroundServicePropertyGroup

	// generated from element "type" of complex type CodePassengerServiceType
	Type_XSD *aixmdatatypes.CodePassengerServiceType `xml:"type,omitempty"`
}

// PilotControlledLightingType Named source named complex type "PilotControlledLightingType"
type PilotControlledLightingType struct {
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

	// generated from element "timeSlice" of complex type PilotControlledLightingTimeSlicePropertyType
	TimeSlice []*PilotControlledLightingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// PilotControlledLightingTimeSlicePropertyType Named source named complex type "PilotControlledLightingTimeSlicePropertyType"
type PilotControlledLightingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "PilotControlledLightingTimeSlice" of complex type PilotControlledLightingTimeSliceType
	PilotControlledLightingTimeSlice *PilotControlledLightingTimeSliceType `xml:"PilotControlledLightingTimeSlice,omitempty"`
}

// PilotControlledLightingTimeSliceType Named source named complex type "PilotControlledLightingTimeSliceType"
type PilotControlledLightingTimeSliceType struct {
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

	// generated from group "aixm:PilotControlledLightingPropertyGroup" with order 6054 depth 2
	Group_PilotControlledLightingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_24 `xml:"extension,omitempty"`
}

// A_extension_24 Named source within outer element "extension"
type A_extension_24 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPilotControlledLightingExtension" of complex type AbstractExtensionType
	AbstractPilotControlledLightingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPilotControlledLightingExtension,omitempty"`
}

// Group_PilotControlledLightingPropertyGroup UnNamed source named group "PilotControlledLightingPropertyGroup"
type Group_PilotControlledLightingPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodePilotControlledLightingType
	Type_XSD *aixmdatatypes.CodePilotControlledLightingType `xml:"type,omitempty"`

	// generated from element "duration" of complex type ValDurationType
	Duration *aixmdatatypes.ValDurationType `xml:"duration,omitempty"`

	// generated from element "intensitySteps" of complex type NoNumberType
	IntensitySteps *aixmdatatypes.NoNumberType `xml:"intensitySteps,omitempty"`

	// generated from element "standByIntensity" of complex type CodeIntensityStandByType
	StandByIntensity *aixmdatatypes.CodeIntensityStandByType `xml:"standByIntensity,omitempty"`

	// generated from element "radioFrequency" of complex type ValFrequencyType
	RadioFrequency *aixmdatatypes.ValFrequencyType `xml:"radioFrequency,omitempty"`

	// generated from element "activationInstruction" of complex type TextInstructionType
	ActivationInstruction *aixmdatatypes.TextInstructionType `xml:"activationInstruction,omitempty"`

	// generated from element "controlledLightIntensity" of complex type LightActivationPropertyType
	ControlledLightIntensity []*LightActivationPropertyType `xml:"controlledLightIntensity,omitempty"`

	// generated from element "activatedGroundLighting" of complex type GroundLightSystemPropertyType
	ActivatedGroundLighting []*GroundLightSystemPropertyType `xml:"activatedGroundLighting,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RadioCommunicationChannelPropertyType Named source named complex type "RadioCommunicationChannelPropertyType"
type RadioCommunicationChannelPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RadioCommunicationChannelType Named source named complex type "RadioCommunicationChannelType"
type RadioCommunicationChannelType struct {
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

	// generated from element "timeSlice" of complex type RadioCommunicationChannelTimeSlicePropertyType
	TimeSlice []*RadioCommunicationChannelTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RadioCommunicationChannelTimeSlicePropertyType Named source named complex type "RadioCommunicationChannelTimeSlicePropertyType"
type RadioCommunicationChannelTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RadioCommunicationChannelTimeSlice" of complex type RadioCommunicationChannelTimeSliceType
	RadioCommunicationChannelTimeSlice *RadioCommunicationChannelTimeSliceType `xml:"RadioCommunicationChannelTimeSlice,omitempty"`
}

// RadioCommunicationChannelTimeSliceType Named source named complex type "RadioCommunicationChannelTimeSliceType"
type RadioCommunicationChannelTimeSliceType struct {
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

	// generated from group "aixm:RadioCommunicationChannelPropertyGroup" with order 6086 depth 2
	Group_RadioCommunicationChannelPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_25 `xml:"extension,omitempty"`
}

// A_extension_25 Named source within outer element "extension"
type A_extension_25 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRadioCommunicationChannelExtension" of complex type AbstractExtensionType
	AbstractRadioCommunicationChannelExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadioCommunicationChannelExtension,omitempty"`
}

// Group_RadioCommunicationChannelPropertyGroup UnNamed source named group "RadioCommunicationChannelPropertyGroup"
type Group_RadioCommunicationChannelPropertyGroup struct {

	// insertion point for fields

	// generated from element "mode" of complex type CodeCommunicationModeType
	Mode *aixmdatatypes.CodeCommunicationModeType `xml:"mode,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "frequencyTransmission" of complex type ValFrequencyType
	FrequencyTransmission *aixmdatatypes.ValFrequencyType `xml:"frequencyTransmission,omitempty"`

	// generated from element "frequencyReception" of complex type ValFrequencyType
	FrequencyReception *aixmdatatypes.ValFrequencyType `xml:"frequencyReception,omitempty"`

	// generated from element "channel" of complex type CodeCommunicationChannelType
	Channel *aixmdatatypes.CodeCommunicationChannelType `xml:"channel,omitempty"`

	// generated from element "logon" of complex type TextDesignatorType
	Logon *aixmdatatypes.TextDesignatorType `xml:"logon,omitempty"`

	// generated from element "emissionType" of complex type CodeRadioEmissionType
	EmissionType *aixmdatatypes.CodeRadioEmissionType `xml:"emissionType,omitempty"`

	// generated from element "selectiveCall" of complex type CodeYesNoType
	SelectiveCall *aixmdatatypes.CodeYesNoType `xml:"selectiveCall,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "trafficDirection" of complex type CodeCommunicationDirectionType
	TrafficDirection *aixmdatatypes.CodeCommunicationDirectionType `xml:"trafficDirection,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location []*ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "availability" of complex type RadioCommunicationOperationalStatusPropertyType
	Availability []*RadioCommunicationOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RadioCommunicationOperationalStatusType Named source named complex type "RadioCommunicationOperationalStatusType"
type RadioCommunicationOperationalStatusType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6111 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:RadioCommunicationOperationalStatusPropertyGroup" with order 6112 depth 2
	Group_RadioCommunicationOperationalStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_26 `xml:"extension,omitempty"`
}

// A_extension_26 Named source within outer element "extension"
type A_extension_26 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractRadioCommunicationOperationalStatusExtension" of complex type AbstractExtensionType
	AbstractRadioCommunicationOperationalStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadioCommunicationOperationalStatusExtension,omitempty"`
}

// Group_RadioCommunicationOperationalStatusPropertyGroup UnNamed source named group "RadioCommunicationOperationalStatusPropertyGroup"
type Group_RadioCommunicationOperationalStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusServiceType
	OperationalStatus *aixmdatatypes.CodeStatusServiceType `xml:"operationalStatus,omitempty"`
}

// RadioCommunicationOperationalStatusPropertyType Named source named complex type "RadioCommunicationOperationalStatusPropertyType"
type RadioCommunicationOperationalStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RadioCommunicationOperationalStatus" of complex type RadioCommunicationOperationalStatusType
	RadioCommunicationOperationalStatus *RadioCommunicationOperationalStatusType `xml:"RadioCommunicationOperationalStatus,omitempty"`
}

// SearchRescueServiceType Named source named complex type "SearchRescueServiceType"
type SearchRescueServiceType struct {
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

	// generated from element "timeSlice" of complex type SearchRescueServiceTimeSlicePropertyType
	TimeSlice []*SearchRescueServiceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SearchRescueServiceTimeSlicePropertyType Named source named complex type "SearchRescueServiceTimeSlicePropertyType"
type SearchRescueServiceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SearchRescueServiceTimeSlice" of complex type SearchRescueServiceTimeSliceType
	SearchRescueServiceTimeSlice *SearchRescueServiceTimeSliceType `xml:"SearchRescueServiceTimeSlice,omitempty"`
}

// SearchRescueServiceTimeSliceType Named source named complex type "SearchRescueServiceTimeSliceType"
type SearchRescueServiceTimeSliceType struct {
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

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeServiceSARType
	Type_XSD *aixmdatatypes.CodeServiceSARType `xml:"type,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_27 `xml:"extension,omitempty"`
}

// A_extension_27 Named source within outer element "extension"
type A_extension_27 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSearchRescueServiceExtension" of complex type AbstractExtensionType
	AbstractSearchRescueServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSearchRescueServiceExtension,omitempty"`

	// generated from element "AbstractServiceExtension" of complex type AbstractExtensionType
	AbstractServiceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceExtension,omitempty"`
}

// Group_SearchRescueServicePropertyGroup UnNamed source named group "SearchRescueServicePropertyGroup"
type Group_SearchRescueServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ServicePropertyGroup" with order 6149 depth 2
	Group_ServicePropertyGroup

	// generated from element "type" of complex type CodeServiceSARType
	Type_XSD *aixmdatatypes.CodeServiceSARType `xml:"type,omitempty"`

	// generated from element "clientAirspace" of complex type AirspacePropertyType
	ClientAirspace []*AirspacePropertyType `xml:"clientAirspace,omitempty"`

	// generated from element "clientRoute" of complex type RoutePortionPropertyType
	ClientRoute []*RoutePortionPropertyType `xml:"clientRoute,omitempty"`
}

// AbstractServiceType Named source named complex type "AbstractServiceType"
type AbstractServiceType struct {
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

// Group_ServicePropertyGroup UnNamed source named group "ServicePropertyGroup"
type Group_ServicePropertyGroup struct {

	// insertion point for fields

	// generated from element "flightOperations" of complex type CodeFlightDestinationType
	FlightOperations *aixmdatatypes.CodeFlightDestinationType `xml:"flightOperations,omitempty"`

	// generated from element "rank" of complex type CodeFacilityRankingType
	Rank *aixmdatatypes.CodeFacilityRankingType `xml:"rank,omitempty"`

	// generated from element "compliantICAO" of complex type CodeYesNoType
	CompliantICAO *aixmdatatypes.CodeYesNoType `xml:"compliantICAO,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "serviceProvider" of complex type UnitPropertyType
	ServiceProvider *UnitPropertyType `xml:"serviceProvider,omitempty"`

	// generated from element "call-sign" of complex type CallsignDetailPropertyType
	Call_sign []*CallsignDetailPropertyType `xml:"call-sign,omitempty"`

	// generated from element "radioCommunication" of complex type RadioCommunicationChannelPropertyType
	RadioCommunication []*RadioCommunicationChannelPropertyType `xml:"radioCommunication,omitempty"`

	// generated from element "groundCommunication" of complex type ContactInformationPropertyType
	GroundCommunication []*ContactInformationPropertyType `xml:"groundCommunication,omitempty"`

	// generated from element "availability" of complex type ServiceOperationalStatusPropertyType
	Availability []*ServiceOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ServiceOperationalStatusType Named source named complex type "ServiceOperationalStatusType"
type ServiceOperationalStatusType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6176 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:ServiceOperationalStatusPropertyGroup" with order 6177 depth 2
	Group_ServiceOperationalStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_28 `xml:"extension,omitempty"`
}

// A_extension_28 Named source within outer element "extension"
type A_extension_28 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractServiceOperationalStatusExtension" of complex type AbstractExtensionType
	AbstractServiceOperationalStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractServiceOperationalStatusExtension,omitempty"`
}

// Group_ServiceOperationalStatusPropertyGroup UnNamed source named group "ServiceOperationalStatusPropertyGroup"
type Group_ServiceOperationalStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusServiceType
	OperationalStatus *aixmdatatypes.CodeStatusServiceType `xml:"operationalStatus,omitempty"`
}

// ServiceOperationalStatusPropertyType Named source named complex type "ServiceOperationalStatusPropertyType"
type ServiceOperationalStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ServiceOperationalStatus" of complex type ServiceOperationalStatusType
	ServiceOperationalStatus *ServiceOperationalStatusType `xml:"ServiceOperationalStatus,omitempty"`
}

// AbstractTrafficSeparationServiceType Named source named complex type "AbstractTrafficSeparationServiceType"
type AbstractTrafficSeparationServiceType struct {
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

// Group_TrafficSeparationServicePropertyGroup UnNamed source named group "TrafficSeparationServicePropertyGroup"
type Group_TrafficSeparationServicePropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:ServicePropertyGroup" with order 6198 depth 2
	Group_ServicePropertyGroup

	// generated from element "radarAssisted" of complex type CodeYesNoType
	RadarAssisted *aixmdatatypes.CodeYesNoType `xml:"radarAssisted,omitempty"`

	// generated from element "dataLinkEnabled" of complex type CodeYesNoType
	DataLinkEnabled *aixmdatatypes.CodeYesNoType `xml:"dataLinkEnabled,omitempty"`

	// generated from element "dataLinkChannel" of complex type CodeCommunicationChannelType
	DataLinkChannel *aixmdatatypes.CodeCommunicationChannelType `xml:"dataLinkChannel,omitempty"`
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
	Extension []*A_extension_29 `xml:"extension,omitempty"`
}

// A_extension_29 Named source within outer element "extension"
type A_extension_29 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6227 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:OnlineContactPropertyGroup" with order 6228 depth 2
	Group_OnlineContactPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_30 `xml:"extension,omitempty"`
}

// A_extension_30 Named source within outer element "extension"
type A_extension_30 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6248 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:PostalAddressPropertyGroup" with order 6249 depth 2
	Group_PostalAddressPropertyGroup

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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6270 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:TelephoneContactPropertyGroup" with order 6271 depth 2
	Group_TelephoneContactPropertyGroup

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

// SpecialDateType Named source named complex type "SpecialDateType"
type SpecialDateType struct {
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

	// generated from element "timeSlice" of complex type SpecialDateTimeSlicePropertyType
	TimeSlice []*SpecialDateTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SpecialDateTimeSlicePropertyType Named source named complex type "SpecialDateTimeSlicePropertyType"
type SpecialDateTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SpecialDateTimeSlice" of complex type SpecialDateTimeSliceType
	SpecialDateTimeSlice *SpecialDateTimeSliceType `xml:"SpecialDateTimeSlice,omitempty"`
}

// SpecialDateTimeSliceType Named source named complex type "SpecialDateTimeSliceType"
type SpecialDateTimeSliceType struct {
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

	// generated from element "type" of complex type CodeSpecialDateType
	Type_XSD *aixmdatatypes.CodeSpecialDateType `xml:"type,omitempty"`

	// generated from element "dateDay" of complex type DateMonthDayType
	DateDay *aixmdatatypes.DateMonthDayType `xml:"dateDay,omitempty"`

	// generated from element "dateYear" of complex type DateYearType
	DateYear *aixmdatatypes.DateYearType `xml:"dateYear,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "authority" of complex type OrganisationAuthorityPropertyType
	Authority *OrganisationAuthorityPropertyType `xml:"authority,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_33 `xml:"extension,omitempty"`
}

// A_extension_33 Named source within outer element "extension"
type A_extension_33 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSpecialDateExtension" of complex type AbstractExtensionType
	AbstractSpecialDateExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSpecialDateExtension,omitempty"`
}

// Group_SpecialDatePropertyGroup UnNamed source named group "SpecialDatePropertyGroup"
type Group_SpecialDatePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeSpecialDateType
	Type_XSD *aixmdatatypes.CodeSpecialDateType `xml:"type,omitempty"`

	// generated from element "dateDay" of complex type DateMonthDayType
	DateDay *aixmdatatypes.DateMonthDayType `xml:"dateDay,omitempty"`

	// generated from element "dateYear" of complex type DateYearType
	DateYear *aixmdatatypes.DateYearType `xml:"dateYear,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "authority" of complex type OrganisationAuthorityPropertyType
	Authority *OrganisationAuthorityPropertyType `xml:"authority,omitempty"`
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
	Extension []*A_extension_34 `xml:"extension,omitempty"`
}

// A_extension_34 Named source within outer element "extension"
type A_extension_34 struct {
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

// StandardLevelType Named source named complex type "StandardLevelType"
type StandardLevelType struct {
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

	// generated from group "aixm:StandardLevelPropertyGroup" with order 6565 depth 2
	Group_StandardLevelPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_35 `xml:"extension,omitempty"`
}

// A_extension_35 Named source within outer element "extension"
type A_extension_35 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardLevelExtension" of complex type AbstractExtensionType
	AbstractStandardLevelExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardLevelExtension,omitempty"`
}

// Group_StandardLevelPropertyGroup UnNamed source named group "StandardLevelPropertyGroup"
type Group_StandardLevelPropertyGroup struct {

	// insertion point for fields

	// generated from element "verticalDistance" of complex type ValDistanceVerticalType
	VerticalDistance *aixmdatatypes.ValDistanceVerticalType `xml:"verticalDistance,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// StandardLevelPropertyType Named source named complex type "StandardLevelPropertyType"
type StandardLevelPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "StandardLevel" of complex type StandardLevelType
	StandardLevel *StandardLevelType `xml:"StandardLevel,omitempty"`
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

// StandardLevelColumnType Named source named complex type "StandardLevelColumnType"
type StandardLevelColumnType struct {
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

	// generated from element "timeSlice" of complex type StandardLevelColumnTimeSlicePropertyType
	TimeSlice []*StandardLevelColumnTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandardLevelColumnTimeSlicePropertyType Named source named complex type "StandardLevelColumnTimeSlicePropertyType"
type StandardLevelColumnTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandardLevelColumnTimeSlice" of complex type StandardLevelColumnTimeSliceType
	StandardLevelColumnTimeSlice *StandardLevelColumnTimeSliceType `xml:"StandardLevelColumnTimeSlice,omitempty"`
}

// StandardLevelColumnTimeSliceType Named source named complex type "StandardLevelColumnTimeSliceType"
type StandardLevelColumnTimeSliceType struct {
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

	// generated from group "aixm:StandardLevelColumnPropertyGroup" with order 6593 depth 2
	Group_StandardLevelColumnPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_36 `xml:"extension,omitempty"`
}

// A_extension_36 Named source within outer element "extension"
type A_extension_36 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardLevelColumnExtension" of complex type AbstractExtensionType
	AbstractStandardLevelColumnExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardLevelColumnExtension,omitempty"`
}

// Group_StandardLevelColumnPropertyGroup UnNamed source named group "StandardLevelColumnPropertyGroup"
type Group_StandardLevelColumnPropertyGroup struct {

	// insertion point for fields

	// generated from element "series" of complex type CodeLevelSeriesType
	Series *aixmdatatypes.CodeLevelSeriesType `xml:"series,omitempty"`

	// generated from element "unitOfMeasurement" of simple type aixm:CodeDistanceVerticalUomType
	UnitOfMeasurement string `xml:"unitOfMeasurement,omitempty"`

	// generated from element "separation" of complex type CodeRVSMType
	Separation *aixmdatatypes.CodeRVSMType `xml:"separation,omitempty"`

	// generated from element "level" of complex type StandardLevelPropertyType
	Level []*StandardLevelPropertyType `xml:"level,omitempty"`

	// generated from element "levelTable" of complex type StandardLevelTablePropertyType
	LevelTable *StandardLevelTablePropertyType `xml:"levelTable,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// StandardLevelSectorType Named source named complex type "StandardLevelSectorType"
type StandardLevelSectorType struct {
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

	// generated from element "timeSlice" of complex type StandardLevelSectorTimeSlicePropertyType
	TimeSlice []*StandardLevelSectorTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandardLevelSectorTimeSlicePropertyType Named source named complex type "StandardLevelSectorTimeSlicePropertyType"
type StandardLevelSectorTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandardLevelSectorTimeSlice" of complex type StandardLevelSectorTimeSliceType
	StandardLevelSectorTimeSlice *StandardLevelSectorTimeSliceType `xml:"StandardLevelSectorTimeSlice,omitempty"`
}

// StandardLevelSectorTimeSliceType Named source named complex type "StandardLevelSectorTimeSliceType"
type StandardLevelSectorTimeSliceType struct {
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

	// generated from group "aixm:StandardLevelSectorPropertyGroup" with order 6622 depth 2
	Group_StandardLevelSectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_37 `xml:"extension,omitempty"`
}

// A_extension_37 Named source within outer element "extension"
type A_extension_37 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardLevelSectorExtension" of complex type AbstractExtensionType
	AbstractStandardLevelSectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardLevelSectorExtension,omitempty"`
}

// Group_StandardLevelSectorPropertyGroup UnNamed source named group "StandardLevelSectorPropertyGroup"
type Group_StandardLevelSectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "flightRule" of complex type CodeFlightRuleType
	FlightRule *aixmdatatypes.CodeFlightRuleType `xml:"flightRule,omitempty"`

	// generated from element "fromTrack" of complex type ValBearingType
	FromTrack *aixmdatatypes.ValBearingType `xml:"fromTrack,omitempty"`

	// generated from element "toTrack" of complex type ValBearingType
	ToTrack *aixmdatatypes.ValBearingType `xml:"toTrack,omitempty"`

	// generated from element "angleType" of complex type CodeNorthReferenceType
	AngleType *aixmdatatypes.CodeNorthReferenceType `xml:"angleType,omitempty"`

	// generated from element "applicableAirspace" of complex type AirspacePropertyType
	ApplicableAirspace []*AirspacePropertyType `xml:"applicableAirspace,omitempty"`

	// generated from element "applicableLevelColumn" of complex type StandardLevelColumnPropertyType
	ApplicableLevelColumn *StandardLevelColumnPropertyType `xml:"applicableLevelColumn,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// StandardLevelTablePropertyType Named source named complex type "StandardLevelTablePropertyType"
type StandardLevelTablePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// StandardLevelTableType Named source named complex type "StandardLevelTableType"
type StandardLevelTableType struct {
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

	// generated from element "timeSlice" of complex type StandardLevelTableTimeSlicePropertyType
	TimeSlice []*StandardLevelTableTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandardLevelTableTimeSlicePropertyType Named source named complex type "StandardLevelTableTimeSlicePropertyType"
type StandardLevelTableTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandardLevelTableTimeSlice" of complex type StandardLevelTableTimeSliceType
	StandardLevelTableTimeSlice *StandardLevelTableTimeSliceType `xml:"StandardLevelTableTimeSlice,omitempty"`
}

// StandardLevelTableTimeSliceType Named source named complex type "StandardLevelTableTimeSliceType"
type StandardLevelTableTimeSliceType struct {
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

	// generated from element "standardICAO" of complex type CodeYesNoType
	StandardICAO *aixmdatatypes.CodeYesNoType `xml:"standardICAO,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_38 `xml:"extension,omitempty"`
}

// A_extension_38 Named source within outer element "extension"
type A_extension_38 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandardLevelTableExtension" of complex type AbstractExtensionType
	AbstractStandardLevelTableExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandardLevelTableExtension,omitempty"`
}

// Group_StandardLevelTablePropertyGroup UnNamed source named group "StandardLevelTablePropertyGroup"
type Group_StandardLevelTablePropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type CodeLevelTableDesignatorType
	NameXSD *aixmdatatypes.CodeLevelTableDesignatorType `xml:"name,omitempty"`

	// generated from element "standardICAO" of complex type CodeYesNoType
	StandardICAO *aixmdatatypes.CodeYesNoType `xml:"standardICAO,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}
