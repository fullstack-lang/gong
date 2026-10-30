// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAircraftCategoryType
var _ *gml.CodeType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// AirspacePropertyType Named source named complex type "AirspacePropertyType"
type AirspacePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AirspaceType Named source named complex type "AirspaceType"
type AirspaceType struct {
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

	// generated from element "timeSlice" of complex type AirspaceTimeSlicePropertyType
	TimeSlice []*AirspaceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirspaceTimeSlicePropertyType Named source named complex type "AirspaceTimeSlicePropertyType"
type AirspaceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirspaceTimeSlice" of complex type AirspaceTimeSliceType
	AirspaceTimeSlice *AirspaceTimeSliceType `xml:"AirspaceTimeSlice,omitempty"`
}

// AirspaceTimeSliceType Named source named complex type "AirspaceTimeSliceType"
type AirspaceTimeSliceType struct {
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

	// generated from element "type" of complex type CodeAirspaceType
	Type_XSD *aixmdatatypes.CodeAirspaceType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeAirspaceDesignatorType
	Designator *aixmdatatypes.CodeAirspaceDesignatorType `xml:"designator,omitempty"`

	// generated from element "localType" of complex type TextNameType
	LocalType *aixmdatatypes.TextNameType `xml:"localType,omitempty"`

	// generated from element "designatorICAO" of complex type CodeYesNoType
	DesignatorICAO *aixmdatatypes.CodeYesNoType `xml:"designatorICAO,omitempty"`

	// generated from element "controlType" of complex type CodeMilitaryOperationsType
	ControlType *aixmdatatypes.CodeMilitaryOperationsType `xml:"controlType,omitempty"`

	// generated from element "upperLowerSeparation" of complex type ValFLType
	UpperLowerSeparation *aixmdatatypes.ValFLType `xml:"upperLowerSeparation,omitempty"`

	// generated from element "class" of complex type AirspaceLayerClassPropertyType
	Class []*AirspaceLayerClassPropertyType `xml:"class,omitempty"`

	// generated from element "protectedRoute" of complex type RoutePropertyType
	ProtectedRoute *RoutePropertyType `xml:"protectedRoute,omitempty"`

	// generated from element "geometryComponent" of complex type AirspaceGeometryComponentPropertyType
	GeometryComponent []*AirspaceGeometryComponentPropertyType `xml:"geometryComponent,omitempty"`

	// generated from element "activation" of complex type AirspaceActivationPropertyType
	Activation []*AirspaceActivationPropertyType `xml:"activation,omitempty"`

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

	// generated from element "AbstractAirspaceExtension" of complex type AbstractExtensionType
	AbstractAirspaceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceExtension,omitempty"`
}

// Group_AirspacePropertyGroup UnNamed source named group "AirspacePropertyGroup"
type Group_AirspacePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAirspaceType
	Type_XSD *aixmdatatypes.CodeAirspaceType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeAirspaceDesignatorType
	Designator *aixmdatatypes.CodeAirspaceDesignatorType `xml:"designator,omitempty"`

	// generated from element "localType" of complex type TextNameType
	LocalType *aixmdatatypes.TextNameType `xml:"localType,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "designatorICAO" of complex type CodeYesNoType
	DesignatorICAO *aixmdatatypes.CodeYesNoType `xml:"designatorICAO,omitempty"`

	// generated from element "controlType" of complex type CodeMilitaryOperationsType
	ControlType *aixmdatatypes.CodeMilitaryOperationsType `xml:"controlType,omitempty"`

	// generated from element "upperLowerSeparation" of complex type ValFLType
	UpperLowerSeparation *aixmdatatypes.ValFLType `xml:"upperLowerSeparation,omitempty"`

	// generated from element "class" of complex type AirspaceLayerClassPropertyType
	Class []*AirspaceLayerClassPropertyType `xml:"class,omitempty"`

	// generated from element "protectedRoute" of complex type RoutePropertyType
	ProtectedRoute *RoutePropertyType `xml:"protectedRoute,omitempty"`

	// generated from element "geometryComponent" of complex type AirspaceGeometryComponentPropertyType
	GeometryComponent []*AirspaceGeometryComponentPropertyType `xml:"geometryComponent,omitempty"`

	// generated from element "activation" of complex type AirspaceActivationPropertyType
	Activation []*AirspaceActivationPropertyType `xml:"activation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirspaceActivationType Named source named complex type "AirspaceActivationType"
type AirspaceActivationType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 2588 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:AirspaceActivationPropertyGroup" with order 2589 depth 2
	Group_AirspaceActivationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_1 `xml:"extension,omitempty"`
}

// A_extension_1 Named source within outer element "extension"
type A_extension_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractAirspaceActivationExtension" of complex type AbstractExtensionType
	AbstractAirspaceActivationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceActivationExtension,omitempty"`
}

// Group_AirspaceActivationPropertyGroup UnNamed source named group "AirspaceActivationPropertyGroup"
type Group_AirspaceActivationPropertyGroup struct {

	// insertion point for fields

	// generated from element "activity" of complex type CodeAirspaceActivityType
	Activity *aixmdatatypes.CodeAirspaceActivityType `xml:"activity,omitempty"`

	// generated from element "status" of complex type CodeStatusAirspaceType
	Status *aixmdatatypes.CodeStatusAirspaceType `xml:"status,omitempty"`

	// generated from element "levels" of complex type AirspaceLayerPropertyType
	Levels []*AirspaceLayerPropertyType `xml:"levels,omitempty"`

	// generated from element "user" of complex type OrganisationAuthorityPropertyType
	User []*OrganisationAuthorityPropertyType `xml:"user,omitempty"`

	// generated from element "aircraft" of complex type AircraftCharacteristicPropertyType
	Aircraft []*AircraftCharacteristicPropertyType `xml:"aircraft,omitempty"`
}

// AirspaceActivationPropertyType Named source named complex type "AirspaceActivationPropertyType"
type AirspaceActivationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceActivation" of complex type AirspaceActivationType
	AirspaceActivation *AirspaceActivationType `xml:"AirspaceActivation,omitempty"`
}

// AirspaceGeometryComponentType Named source named complex type "AirspaceGeometryComponentType"
type AirspaceGeometryComponentType struct {
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

	// generated from group "aixm:AirspaceGeometryComponentPropertyGroup" with order 2610 depth 2
	Group_AirspaceGeometryComponentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirspaceGeometryComponentExtension" of complex type AbstractExtensionType
	AbstractAirspaceGeometryComponentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceGeometryComponentExtension,omitempty"`
}

// Group_AirspaceGeometryComponentPropertyGroup UnNamed source named group "AirspaceGeometryComponentPropertyGroup"
type Group_AirspaceGeometryComponentPropertyGroup struct {

	// insertion point for fields

	// generated from element "operation" of complex type CodeAirspaceAggregationType
	Operation *aixmdatatypes.CodeAirspaceAggregationType `xml:"operation,omitempty"`

	// generated from element "operationSequence" of complex type NoSequenceType
	OperationSequence *aixmdatatypes.NoSequenceType `xml:"operationSequence,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theAirspaceVolume" of complex type AirspaceVolumePropertyType
	TheAirspaceVolume *AirspaceVolumePropertyType `xml:"theAirspaceVolume,omitempty"`
}

// AirspaceGeometryComponentPropertyType Named source named complex type "AirspaceGeometryComponentPropertyType"
type AirspaceGeometryComponentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceGeometryComponent" of complex type AirspaceGeometryComponentType
	AirspaceGeometryComponent *AirspaceGeometryComponentType `xml:"AirspaceGeometryComponent,omitempty"`
}

// AirspaceLayerClassType Named source named complex type "AirspaceLayerClassType"
type AirspaceLayerClassType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 2629 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:AirspaceLayerClassPropertyGroup" with order 2630 depth 2
	Group_AirspaceLayerClassPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractAirspaceLayerClassExtension" of complex type AbstractExtensionType
	AbstractAirspaceLayerClassExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceLayerClassExtension,omitempty"`
}

// Group_AirspaceLayerClassPropertyGroup UnNamed source named group "AirspaceLayerClassPropertyGroup"
type Group_AirspaceLayerClassPropertyGroup struct {

	// insertion point for fields

	// generated from element "classification" of complex type CodeAirspaceClassificationType
	Classification *aixmdatatypes.CodeAirspaceClassificationType `xml:"classification,omitempty"`

	// generated from element "associatedLevels" of complex type AirspaceLayerPropertyType
	AssociatedLevels []*AirspaceLayerPropertyType `xml:"associatedLevels,omitempty"`
}

// AirspaceLayerClassPropertyType Named source named complex type "AirspaceLayerClassPropertyType"
type AirspaceLayerClassPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceLayerClass" of complex type AirspaceLayerClassType
	AirspaceLayerClass *AirspaceLayerClassType `xml:"AirspaceLayerClass,omitempty"`
}

// AirspaceVolumeType Named source named complex type "AirspaceVolumeType"
type AirspaceVolumeType struct {
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

	// generated from group "aixm:AirspaceVolumePropertyGroup" with order 2648 depth 2
	Group_AirspaceVolumePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirspaceVolumeExtension" of complex type AbstractExtensionType
	AbstractAirspaceVolumeExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceVolumeExtension,omitempty"`
}

// Group_AirspaceVolumePropertyGroup UnNamed source named group "AirspaceVolumePropertyGroup"
type Group_AirspaceVolumePropertyGroup struct {

	// insertion point for fields

	// generated from element "upperLimit" of complex type ValDistanceVerticalType
	UpperLimit *aixmdatatypes.ValDistanceVerticalType `xml:"upperLimit,omitempty"`

	// generated from element "upperLimitReference" of complex type CodeVerticalReferenceType
	UpperLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"upperLimitReference,omitempty"`

	// generated from element "maximumLimit" of complex type ValDistanceVerticalType
	MaximumLimit *aixmdatatypes.ValDistanceVerticalType `xml:"maximumLimit,omitempty"`

	// generated from element "maximumLimitReference" of complex type CodeVerticalReferenceType
	MaximumLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"maximumLimitReference,omitempty"`

	// generated from element "lowerLimit" of complex type ValDistanceVerticalType
	LowerLimit *aixmdatatypes.ValDistanceVerticalType `xml:"lowerLimit,omitempty"`

	// generated from element "lowerLimitReference" of complex type CodeVerticalReferenceType
	LowerLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"lowerLimitReference,omitempty"`

	// generated from element "minimumLimit" of complex type ValDistanceVerticalType
	MinimumLimit *aixmdatatypes.ValDistanceVerticalType `xml:"minimumLimit,omitempty"`

	// generated from element "minimumLimitReference" of complex type CodeVerticalReferenceType
	MinimumLimitReference *aixmdatatypes.CodeVerticalReferenceType `xml:"minimumLimitReference,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "horizontalProjection" of complex type SurfacePropertyType
	HorizontalProjection *SurfacePropertyType `xml:"horizontalProjection,omitempty"`

	// generated from element "centreline" of complex type CurvePropertyType
	Centreline *CurvePropertyType `xml:"centreline,omitempty"`

	// generated from element "contributorAirspace" of complex type AirspaceVolumeDependencyPropertyType
	ContributorAirspace *AirspaceVolumeDependencyPropertyType `xml:"contributorAirspace,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirspaceVolumePropertyType Named source named complex type "AirspaceVolumePropertyType"
type AirspaceVolumePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceVolume" of complex type AirspaceVolumeType
	AirspaceVolume *AirspaceVolumeType `xml:"AirspaceVolume,omitempty"`
}

// AirspaceVolumeDependencyType Named source named complex type "AirspaceVolumeDependencyType"
type AirspaceVolumeDependencyType struct {
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

	// generated from group "aixm:AirspaceVolumeDependencyPropertyGroup" with order 2676 depth 2
	Group_AirspaceVolumeDependencyPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirspaceVolumeDependencyExtension" of complex type AbstractExtensionType
	AbstractAirspaceVolumeDependencyExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirspaceVolumeDependencyExtension,omitempty"`
}

// Group_AirspaceVolumeDependencyPropertyGroup UnNamed source named group "AirspaceVolumeDependencyPropertyGroup"
type Group_AirspaceVolumeDependencyPropertyGroup struct {

	// insertion point for fields

	// generated from element "dependency" of complex type CodeAirspaceDependencyType
	Dependency *aixmdatatypes.CodeAirspaceDependencyType `xml:"dependency,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theAirspace" of complex type AirspacePropertyType
	TheAirspace *AirspacePropertyType `xml:"theAirspace,omitempty"`
}

// AirspaceVolumeDependencyPropertyType Named source named complex type "AirspaceVolumeDependencyPropertyType"
type AirspaceVolumeDependencyPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirspaceVolumeDependency" of complex type AirspaceVolumeDependencyType
	AirspaceVolumeDependency *AirspaceVolumeDependencyType `xml:"AirspaceVolumeDependency,omitempty"`
}

// AuthorityForAirspaceType Named source named complex type "AuthorityForAirspaceType"
type AuthorityForAirspaceType struct {
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

	// generated from element "timeSlice" of complex type AuthorityForAirspaceTimeSlicePropertyType
	TimeSlice []*AuthorityForAirspaceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AuthorityForAirspaceTimeSlicePropertyType Named source named complex type "AuthorityForAirspaceTimeSlicePropertyType"
type AuthorityForAirspaceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AuthorityForAirspaceTimeSlice" of complex type AuthorityForAirspaceTimeSliceType
	AuthorityForAirspaceTimeSlice *AuthorityForAirspaceTimeSliceType `xml:"AuthorityForAirspaceTimeSlice,omitempty"`
}

// AuthorityForAirspaceTimeSliceType Named source named complex type "AuthorityForAirspaceTimeSliceType"
type AuthorityForAirspaceTimeSliceType struct {
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

	// generated from group "aixm:AuthorityForAirspacePropertyGroup" with order 2705 depth 2
	Group_AuthorityForAirspacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAuthorityForAirspaceExtension" of complex type AbstractExtensionType
	AbstractAuthorityForAirspaceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAuthorityForAirspaceExtension,omitempty"`
}

// Group_AuthorityForAirspacePropertyGroup UnNamed source named group "AuthorityForAirspacePropertyGroup"
type Group_AuthorityForAirspacePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAuthorityType
	Type_XSD *aixmdatatypes.CodeAuthorityType `xml:"type,omitempty"`

	// generated from element "responsibleOrganisation" of complex type OrganisationAuthorityPropertyType
	ResponsibleOrganisation *OrganisationAuthorityPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "assignedAirspace" of complex type AirspacePropertyType
	AssignedAirspace *AirspacePropertyType `xml:"assignedAirspace,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// GeoBorderType Named source named complex type "GeoBorderType"
type GeoBorderType struct {
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

	// generated from element "timeSlice" of complex type GeoBorderTimeSlicePropertyType
	TimeSlice []*GeoBorderTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GeoBorderTimeSlicePropertyType Named source named complex type "GeoBorderTimeSlicePropertyType"
type GeoBorderTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GeoBorderTimeSlice" of complex type GeoBorderTimeSliceType
	GeoBorderTimeSlice *GeoBorderTimeSliceType `xml:"GeoBorderTimeSlice,omitempty"`
}

// GeoBorderTimeSliceType Named source named complex type "GeoBorderTimeSliceType"
type GeoBorderTimeSliceType struct {
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

	// generated from element "type" of complex type CodeGeoBorderType
	Type_XSD *aixmdatatypes.CodeGeoBorderType `xml:"type,omitempty"`

	// generated from element "border" of complex type CurvePropertyType
	Border *CurvePropertyType `xml:"border,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGeoBorderExtension" of complex type AbstractExtensionType
	AbstractGeoBorderExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGeoBorderExtension,omitempty"`
}

// Group_GeoBorderPropertyGroup UnNamed source named group "GeoBorderPropertyGroup"
type Group_GeoBorderPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeGeoBorderType
	Type_XSD *aixmdatatypes.CodeGeoBorderType `xml:"type,omitempty"`

	// generated from element "border" of complex type CurvePropertyType
	Border *CurvePropertyType `xml:"border,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

	// generated from group "aixm:CurvePropertyGroup" with order 2748 depth 2
	Group_CurvePropertyGroup

	// generated from group "aixm:ElevatedCurvePropertyGroup" with order 2759 depth 2
	Group_ElevatedCurvePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
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
	Extension []*A_extension_9 `xml:"extension,omitempty"`
}

// A_extension_9 Named source within outer element "extension"
type A_extension_9 struct {
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
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
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
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
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

// RoutePropertyType Named source named complex type "RoutePropertyType"
type RoutePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
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

	// generated from group "aixm:AircraftCharacteristicPropertyGroup" with order 6273 depth 2
	Group_AircraftCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
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

	// generated from group "aixm:AirspaceLayerPropertyGroup" with order 6332 depth 2
	Group_AirspaceLayerPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
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
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
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
