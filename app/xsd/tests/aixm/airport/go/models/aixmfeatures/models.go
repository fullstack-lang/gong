// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeBuoyDesignatorType
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

// AirportHeliportType Named source named complex type "AirportHeliportType"
type AirportHeliportType struct {
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

	// generated from element "timeSlice" of complex type AirportHeliportTimeSlicePropertyType
	TimeSlice []*AirportHeliportTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportHeliportTimeSlicePropertyType Named source named complex type "AirportHeliportTimeSlicePropertyType"
type AirportHeliportTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportHeliportTimeSlice" of complex type AirportHeliportTimeSliceType
	AirportHeliportTimeSlice *AirportHeliportTimeSliceType `xml:"AirportHeliportTimeSlice,omitempty"`
}

// AirportHeliportTimeSliceType Named source named complex type "AirportHeliportTimeSliceType"
type AirportHeliportTimeSliceType struct {
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

	// generated from element "designator" of complex type CodeAirportHeliportDesignatorType
	Designator *aixmdatatypes.CodeAirportHeliportDesignatorType `xml:"designator,omitempty"`

	// generated from element "locationIndicatorICAO" of complex type CodeICAOType
	LocationIndicatorICAO *aixmdatatypes.CodeICAOType `xml:"locationIndicatorICAO,omitempty"`

	// generated from element "designatorIATA" of complex type CodeIATAType
	DesignatorIATA *aixmdatatypes.CodeIATAType `xml:"designatorIATA,omitempty"`

	// generated from element "type" of complex type CodeAirportHeliportType
	Type_XSD *aixmdatatypes.CodeAirportHeliportType `xml:"type,omitempty"`

	// generated from element "certifiedICAO" of complex type CodeYesNoType
	CertifiedICAO *aixmdatatypes.CodeYesNoType `xml:"certifiedICAO,omitempty"`

	// generated from element "privateUse" of complex type CodeYesNoType
	PrivateUse *aixmdatatypes.CodeYesNoType `xml:"privateUse,omitempty"`

	// generated from element "controlType" of complex type CodeMilitaryOperationsType
	ControlType *aixmdatatypes.CodeMilitaryOperationsType `xml:"controlType,omitempty"`

	// generated from element "fieldElevation" of complex type ValDistanceVerticalType
	FieldElevation *aixmdatatypes.ValDistanceVerticalType `xml:"fieldElevation,omitempty"`

	// generated from element "fieldElevationAccuracy" of complex type ValDistanceVerticalType
	FieldElevationAccuracy *aixmdatatypes.ValDistanceVerticalType `xml:"fieldElevationAccuracy,omitempty"`

	// generated from element "verticalDatum" of complex type CodeVerticalDatumType
	VerticalDatum *aixmdatatypes.CodeVerticalDatumType `xml:"verticalDatum,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "magneticVariationChange" of complex type ValMagneticVariationChangeType
	MagneticVariationChange *aixmdatatypes.ValMagneticVariationChangeType `xml:"magneticVariationChange,omitempty"`

	// generated from element "referenceTemperature" of complex type ValTemperatureType
	ReferenceTemperature *aixmdatatypes.ValTemperatureType `xml:"referenceTemperature,omitempty"`

	// generated from element "altimeterCheckLocation" of complex type CodeYesNoType
	AltimeterCheckLocation *aixmdatatypes.CodeYesNoType `xml:"altimeterCheckLocation,omitempty"`

	// generated from element "secondaryPowerSupply" of complex type CodeYesNoType
	SecondaryPowerSupply *aixmdatatypes.CodeYesNoType `xml:"secondaryPowerSupply,omitempty"`

	// generated from element "windDirectionIndicator" of complex type CodeYesNoType
	WindDirectionIndicator *aixmdatatypes.CodeYesNoType `xml:"windDirectionIndicator,omitempty"`

	// generated from element "landingDirectionIndicator" of complex type CodeYesNoType
	LandingDirectionIndicator *aixmdatatypes.CodeYesNoType `xml:"landingDirectionIndicator,omitempty"`

	// generated from element "transitionAltitude" of complex type ValDistanceVerticalType
	TransitionAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"transitionAltitude,omitempty"`

	// generated from element "transitionLevel" of complex type ValFLType
	TransitionLevel *aixmdatatypes.ValFLType `xml:"transitionLevel,omitempty"`

	// generated from element "lowestTemperature" of complex type ValTemperatureType
	LowestTemperature *aixmdatatypes.ValTemperatureType `xml:"lowestTemperature,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "certificationDate" of complex type DateType
	CertificationDate *aixmdatatypes.DateType `xml:"certificationDate,omitempty"`

	// generated from element "certificationExpirationDate" of complex type DateType
	CertificationExpirationDate *aixmdatatypes.DateType `xml:"certificationExpirationDate,omitempty"`

	// generated from element "contaminant" of complex type AirportHeliportContaminationPropertyType
	Contaminant []*AirportHeliportContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "servedCity" of complex type CityPropertyType
	ServedCity []*CityPropertyType `xml:"servedCity,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AirportHeliportResponsibilityOrganisationPropertyType
	ResponsibleOrganisation *AirportHeliportResponsibilityOrganisationPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "ARP" of complex type ElevatedPointPropertyType
	ARP *ElevatedPointPropertyType `xml:"ARP,omitempty"`

	// generated from element "aviationBoundary" of complex type ElevatedSurfacePropertyType
	AviationBoundary *ElevatedSurfacePropertyType `xml:"aviationBoundary,omitempty"`

	// generated from element "altimeterSource" of complex type AltimeterSourcePropertyType
	AltimeterSource []*AltimeterSourcePropertyType `xml:"altimeterSource,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "availability" of complex type AirportHeliportAvailabilityPropertyType
	Availability []*AirportHeliportAvailabilityPropertyType `xml:"availability,omitempty"`

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

	// generated from element "AbstractAirportHeliportExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportExtension,omitempty"`
}

// Group_AirportHeliportPropertyGroup UnNamed source named group "AirportHeliportPropertyGroup"
type Group_AirportHeliportPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type CodeAirportHeliportDesignatorType
	Designator *aixmdatatypes.CodeAirportHeliportDesignatorType `xml:"designator,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "locationIndicatorICAO" of complex type CodeICAOType
	LocationIndicatorICAO *aixmdatatypes.CodeICAOType `xml:"locationIndicatorICAO,omitempty"`

	// generated from element "designatorIATA" of complex type CodeIATAType
	DesignatorIATA *aixmdatatypes.CodeIATAType `xml:"designatorIATA,omitempty"`

	// generated from element "type" of complex type CodeAirportHeliportType
	Type_XSD *aixmdatatypes.CodeAirportHeliportType `xml:"type,omitempty"`

	// generated from element "certifiedICAO" of complex type CodeYesNoType
	CertifiedICAO *aixmdatatypes.CodeYesNoType `xml:"certifiedICAO,omitempty"`

	// generated from element "privateUse" of complex type CodeYesNoType
	PrivateUse *aixmdatatypes.CodeYesNoType `xml:"privateUse,omitempty"`

	// generated from element "controlType" of complex type CodeMilitaryOperationsType
	ControlType *aixmdatatypes.CodeMilitaryOperationsType `xml:"controlType,omitempty"`

	// generated from element "fieldElevation" of complex type ValDistanceVerticalType
	FieldElevation *aixmdatatypes.ValDistanceVerticalType `xml:"fieldElevation,omitempty"`

	// generated from element "fieldElevationAccuracy" of complex type ValDistanceVerticalType
	FieldElevationAccuracy *aixmdatatypes.ValDistanceVerticalType `xml:"fieldElevationAccuracy,omitempty"`

	// generated from element "verticalDatum" of complex type CodeVerticalDatumType
	VerticalDatum *aixmdatatypes.CodeVerticalDatumType `xml:"verticalDatum,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "magneticVariationChange" of complex type ValMagneticVariationChangeType
	MagneticVariationChange *aixmdatatypes.ValMagneticVariationChangeType `xml:"magneticVariationChange,omitempty"`

	// generated from element "referenceTemperature" of complex type ValTemperatureType
	ReferenceTemperature *aixmdatatypes.ValTemperatureType `xml:"referenceTemperature,omitempty"`

	// generated from element "altimeterCheckLocation" of complex type CodeYesNoType
	AltimeterCheckLocation *aixmdatatypes.CodeYesNoType `xml:"altimeterCheckLocation,omitempty"`

	// generated from element "secondaryPowerSupply" of complex type CodeYesNoType
	SecondaryPowerSupply *aixmdatatypes.CodeYesNoType `xml:"secondaryPowerSupply,omitempty"`

	// generated from element "windDirectionIndicator" of complex type CodeYesNoType
	WindDirectionIndicator *aixmdatatypes.CodeYesNoType `xml:"windDirectionIndicator,omitempty"`

	// generated from element "landingDirectionIndicator" of complex type CodeYesNoType
	LandingDirectionIndicator *aixmdatatypes.CodeYesNoType `xml:"landingDirectionIndicator,omitempty"`

	// generated from element "transitionAltitude" of complex type ValDistanceVerticalType
	TransitionAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"transitionAltitude,omitempty"`

	// generated from element "transitionLevel" of complex type ValFLType
	TransitionLevel *aixmdatatypes.ValFLType `xml:"transitionLevel,omitempty"`

	// generated from element "lowestTemperature" of complex type ValTemperatureType
	LowestTemperature *aixmdatatypes.ValTemperatureType `xml:"lowestTemperature,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "certificationDate" of complex type DateType
	CertificationDate *aixmdatatypes.DateType `xml:"certificationDate,omitempty"`

	// generated from element "certificationExpirationDate" of complex type DateType
	CertificationExpirationDate *aixmdatatypes.DateType `xml:"certificationExpirationDate,omitempty"`

	// generated from element "contaminant" of complex type AirportHeliportContaminationPropertyType
	Contaminant []*AirportHeliportContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "servedCity" of complex type CityPropertyType
	ServedCity []*CityPropertyType `xml:"servedCity,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AirportHeliportResponsibilityOrganisationPropertyType
	ResponsibleOrganisation *AirportHeliportResponsibilityOrganisationPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "ARP" of complex type ElevatedPointPropertyType
	ARP *ElevatedPointPropertyType `xml:"ARP,omitempty"`

	// generated from element "aviationBoundary" of complex type ElevatedSurfacePropertyType
	AviationBoundary *ElevatedSurfacePropertyType `xml:"aviationBoundary,omitempty"`

	// generated from element "altimeterSource" of complex type AltimeterSourcePropertyType
	AltimeterSource []*AltimeterSourcePropertyType `xml:"altimeterSource,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "availability" of complex type AirportHeliportAvailabilityPropertyType
	Availability []*AirportHeliportAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirportHeliportAvailabilityType Named source named complex type "AirportHeliportAvailabilityType"
type AirportHeliportAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 599 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:AirportHeliportAvailabilityPropertyGroup" with order 600 depth 2
	Group_AirportHeliportAvailabilityPropertyGroup

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

	// generated from element "AbstractAirportHeliportAvailabilityExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportAvailabilityExtension,omitempty"`
}

// Group_AirportHeliportAvailabilityPropertyGroup UnNamed source named group "AirportHeliportAvailabilityPropertyGroup"
type Group_AirportHeliportAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusAirportType
	OperationalStatus *aixmdatatypes.CodeStatusAirportType `xml:"operationalStatus,omitempty"`

	// generated from element "warning" of complex type CodeAirportWarningType
	Warning *aixmdatatypes.CodeAirportWarningType `xml:"warning,omitempty"`

	// generated from element "usage" of complex type AirportHeliportUsagePropertyType
	Usage []*AirportHeliportUsagePropertyType `xml:"usage,omitempty"`
}

// AirportHeliportAvailabilityPropertyType Named source named complex type "AirportHeliportAvailabilityPropertyType"
type AirportHeliportAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirportHeliportAvailability" of complex type AirportHeliportAvailabilityType
	AirportHeliportAvailability *AirportHeliportAvailabilityType `xml:"AirportHeliportAvailability,omitempty"`
}

// AirportHeliportCollocationType Named source named complex type "AirportHeliportCollocationType"
type AirportHeliportCollocationType struct {
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

	// generated from element "timeSlice" of complex type AirportHeliportCollocationTimeSlicePropertyType
	TimeSlice []*AirportHeliportCollocationTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportHeliportCollocationTimeSlicePropertyType Named source named complex type "AirportHeliportCollocationTimeSlicePropertyType"
type AirportHeliportCollocationTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportHeliportCollocationTimeSlice" of complex type AirportHeliportCollocationTimeSliceType
	AirportHeliportCollocationTimeSlice *AirportHeliportCollocationTimeSliceType `xml:"AirportHeliportCollocationTimeSlice,omitempty"`
}

// AirportHeliportCollocationTimeSliceType Named source named complex type "AirportHeliportCollocationTimeSliceType"
type AirportHeliportCollocationTimeSliceType struct {
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

	// generated from group "aixm:AirportHeliportCollocationPropertyGroup" with order 630 depth 2
	Group_AirportHeliportCollocationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirportHeliportCollocationExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportCollocationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportCollocationExtension,omitempty"`
}

// Group_AirportHeliportCollocationPropertyGroup UnNamed source named group "AirportHeliportCollocationPropertyGroup"
type Group_AirportHeliportCollocationPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAirportHeliportCollocationType
	Type_XSD *aixmdatatypes.CodeAirportHeliportCollocationType `xml:"type,omitempty"`

	// generated from element "hostAirport" of complex type AirportHeliportPropertyType
	HostAirport *AirportHeliportPropertyType `xml:"hostAirport,omitempty"`

	// generated from element "dependentAirport" of complex type AirportHeliportPropertyType
	DependentAirport *AirportHeliportPropertyType `xml:"dependentAirport,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirportHeliportProtectionAreaPropertyType Named source named complex type "AirportHeliportProtectionAreaPropertyType"
type AirportHeliportProtectionAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AbstractAirportHeliportProtectionAreaType Named source named complex type "AbstractAirportHeliportProtectionAreaType"
type AbstractAirportHeliportProtectionAreaType struct {
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

// Group_AirportHeliportProtectionAreaPropertyGroup UnNamed source named group "AirportHeliportProtectionAreaPropertyGroup"
type Group_AirportHeliportProtectionAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "lighting" of complex type CodeYesNoType
	Lighting *aixmdatatypes.CodeYesNoType `xml:"lighting,omitempty"`

	// generated from element "obstacleFree" of complex type CodeYesNoType
	ObstacleFree *aixmdatatypes.CodeYesNoType `xml:"obstacleFree,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AirportHeliportResponsibilityOrganisationType Named source named complex type "AirportHeliportResponsibilityOrganisationType"
type AirportHeliportResponsibilityOrganisationType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 661 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:AirportHeliportResponsibilityOrganisationPropertyGroup" with order 662 depth 2
	Group_AirportHeliportResponsibilityOrganisationPropertyGroup

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

	// generated from element "AbstractAirportHeliportResponsibilityOrganisationExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportResponsibilityOrganisationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportResponsibilityOrganisationExtension,omitempty"`
}

// Group_AirportHeliportResponsibilityOrganisationPropertyGroup UnNamed source named group "AirportHeliportResponsibilityOrganisationPropertyGroup"
type Group_AirportHeliportResponsibilityOrganisationPropertyGroup struct {

	// insertion point for fields

	// generated from element "role" of complex type CodeAuthorityRoleType
	Role *aixmdatatypes.CodeAuthorityRoleType `xml:"role,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// AirportHeliportResponsibilityOrganisationPropertyType Named source named complex type "AirportHeliportResponsibilityOrganisationPropertyType"
type AirportHeliportResponsibilityOrganisationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirportHeliportResponsibilityOrganisation" of complex type AirportHeliportResponsibilityOrganisationType
	AirportHeliportResponsibilityOrganisation *AirportHeliportResponsibilityOrganisationType `xml:"AirportHeliportResponsibilityOrganisation,omitempty"`
}

// AirportHeliportUsageType Named source named complex type "AirportHeliportUsageType"
type AirportHeliportUsageType struct {
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

	// generated from group "aixm:UsageConditionPropertyGroup" with order 680 depth 2
	Group_UsageConditionPropertyGroup

	// generated from group "aixm:AirportHeliportUsagePropertyGroup" with order 681 depth 2
	Group_AirportHeliportUsagePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUsageConditionExtension" of complex type AbstractExtensionType
	AbstractUsageConditionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUsageConditionExtension,omitempty"`

	// generated from element "AbstractAirportHeliportUsageExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportUsageExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportUsageExtension,omitempty"`
}

// Group_AirportHeliportUsagePropertyGroup UnNamed source named group "AirportHeliportUsagePropertyGroup"
type Group_AirportHeliportUsagePropertyGroup struct {

	// insertion point for fields

	// generated from element "operation" of complex type CodeOperationAirportHeliportType
	Operation *aixmdatatypes.CodeOperationAirportHeliportType `xml:"operation,omitempty"`
}

// AirportHeliportUsagePropertyType Named source named complex type "AirportHeliportUsagePropertyType"
type AirportHeliportUsagePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirportHeliportUsage" of complex type AirportHeliportUsageType
	AirportHeliportUsage *AirportHeliportUsageType `xml:"AirportHeliportUsage,omitempty"`
}

// AirportHotSpotType Named source named complex type "AirportHotSpotType"
type AirportHotSpotType struct {
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

	// generated from element "timeSlice" of complex type AirportHotSpotTimeSlicePropertyType
	TimeSlice []*AirportHotSpotTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportHotSpotTimeSlicePropertyType Named source named complex type "AirportHotSpotTimeSlicePropertyType"
type AirportHotSpotTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportHotSpotTimeSlice" of complex type AirportHotSpotTimeSliceType
	AirportHotSpotTimeSlice *AirportHotSpotTimeSliceType `xml:"AirportHotSpotTimeSlice,omitempty"`
}

// AirportHotSpotTimeSliceType Named source named complex type "AirportHotSpotTimeSliceType"
type AirportHotSpotTimeSliceType struct {
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

	// generated from group "aixm:AirportHotSpotPropertyGroup" with order 709 depth 2
	Group_AirportHotSpotPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirportHotSpotExtension" of complex type AbstractExtensionType
	AbstractAirportHotSpotExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHotSpotExtension,omitempty"`
}

// Group_AirportHotSpotPropertyGroup UnNamed source named group "AirportHotSpotPropertyGroup"
type Group_AirportHotSpotPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "instruction" of complex type TextInstructionType
	Instruction *aixmdatatypes.TextInstructionType `xml:"instruction,omitempty"`

	// generated from element "area" of complex type ElevatedSurfacePropertyType
	Area *ElevatedSurfacePropertyType `xml:"area,omitempty"`

	// generated from element "affectedAirport" of complex type AirportHeliportPropertyType
	AffectedAirport *AirportHeliportPropertyType `xml:"affectedAirport,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

// AltimeterSourceType Named source named complex type "AltimeterSourceType"
type AltimeterSourceType struct {
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

	// generated from element "timeSlice" of complex type AltimeterSourceTimeSlicePropertyType
	TimeSlice []*AltimeterSourceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AltimeterSourceTimeSlicePropertyType Named source named complex type "AltimeterSourceTimeSlicePropertyType"
type AltimeterSourceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AltimeterSourceTimeSlice" of complex type AltimeterSourceTimeSliceType
	AltimeterSourceTimeSlice *AltimeterSourceTimeSliceType `xml:"AltimeterSourceTimeSlice,omitempty"`
}

// AltimeterSourceTimeSliceType Named source named complex type "AltimeterSourceTimeSliceType"
type AltimeterSourceTimeSliceType struct {
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

	// generated from group "aixm:AltimeterSourcePropertyGroup" with order 737 depth 2
	Group_AltimeterSourcePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAltimeterSourceExtension" of complex type AbstractExtensionType
	AbstractAltimeterSourceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAltimeterSourceExtension,omitempty"`
}

// Group_AltimeterSourcePropertyGroup UnNamed source named group "AltimeterSourcePropertyGroup"
type Group_AltimeterSourcePropertyGroup struct {

	// insertion point for fields

	// generated from element "isRemote" of complex type CodeYesNoType
	IsRemote *aixmdatatypes.CodeYesNoType `xml:"isRemote,omitempty"`

	// generated from element "isPrimary" of complex type CodeYesNoType
	IsPrimary *aixmdatatypes.CodeYesNoType `xml:"isPrimary,omitempty"`

	// generated from element "availability" of complex type AltimeterSourceStatusPropertyType
	Availability []*AltimeterSourceStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AltimeterSourceStatusType Named source named complex type "AltimeterSourceStatusType"
type AltimeterSourceStatusType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 753 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:AltimeterSourceStatusPropertyGroup" with order 754 depth 2
	Group_AltimeterSourceStatusPropertyGroup

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

	// generated from element "AbstractAltimeterSourceStatusExtension" of complex type AbstractExtensionType
	AbstractAltimeterSourceStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAltimeterSourceStatusExtension,omitempty"`
}

// Group_AltimeterSourceStatusPropertyGroup UnNamed source named group "AltimeterSourceStatusPropertyGroup"
type Group_AltimeterSourceStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusOperationsType
	OperationalStatus *aixmdatatypes.CodeStatusOperationsType `xml:"operationalStatus,omitempty"`
}

// AltimeterSourceStatusPropertyType Named source named complex type "AltimeterSourceStatusPropertyType"
type AltimeterSourceStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AltimeterSourceStatus" of complex type AltimeterSourceStatusType
	AltimeterSourceStatus *AltimeterSourceStatusType `xml:"AltimeterSourceStatus,omitempty"`
}

// CityType Named source named complex type "CityType"
type CityType struct {
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

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCityExtension" of complex type AbstractExtensionType
	AbstractCityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCityExtension,omitempty"`
}

// Group_CityPropertyGroup UnNamed source named group "CityPropertyGroup"
type Group_CityPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// CityPropertyType Named source named complex type "CityPropertyType"
type CityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "City" of complex type CityType
	City *CityType `xml:"City,omitempty"`
}

// ConditionCombinationType Named source named complex type "ConditionCombinationType"
type ConditionCombinationType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 788 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:ConditionCombinationPropertyGroup" with order 789 depth 2
	Group_ConditionCombinationPropertyGroup

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

	// generated from element "AbstractConditionCombinationExtension" of complex type AbstractExtensionType
	AbstractConditionCombinationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractConditionCombinationExtension,omitempty"`
}

// Group_ConditionCombinationPropertyGroup UnNamed source named group "ConditionCombinationPropertyGroup"
type Group_ConditionCombinationPropertyGroup struct {

	// insertion point for fields

	// generated from element "logicalOperator" of complex type CodeLogicalOperatorType
	LogicalOperator *aixmdatatypes.CodeLogicalOperatorType `xml:"logicalOperator,omitempty"`

	// generated from element "weather" of complex type MeteorologyPropertyType
	Weather []*MeteorologyPropertyType `xml:"weather,omitempty"`

	// generated from element "aircraft" of complex type AircraftCharacteristicPropertyType
	Aircraft []*AircraftCharacteristicPropertyType `xml:"aircraft,omitempty"`

	// generated from element "flight" of complex type FlightCharacteristicPropertyType
	Flight []*FlightCharacteristicPropertyType `xml:"flight,omitempty"`

	// generated from element "subCondition" of complex type ConditionCombinationPropertyType
	SubCondition []*ConditionCombinationPropertyType `xml:"subCondition,omitempty"`
}

// ConditionCombinationPropertyType Named source named complex type "ConditionCombinationPropertyType"
type ConditionCombinationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ConditionCombination" of complex type ConditionCombinationType
	ConditionCombination *ConditionCombinationType `xml:"ConditionCombination,omitempty"`
}

// NonMovementAreaType Named source named complex type "NonMovementAreaType"
type NonMovementAreaType struct {
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

	// generated from element "timeSlice" of complex type NonMovementAreaTimeSlicePropertyType
	TimeSlice []*NonMovementAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// NonMovementAreaTimeSlicePropertyType Named source named complex type "NonMovementAreaTimeSlicePropertyType"
type NonMovementAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "NonMovementAreaTimeSlice" of complex type NonMovementAreaTimeSliceType
	NonMovementAreaTimeSlice *NonMovementAreaTimeSliceType `xml:"NonMovementAreaTimeSlice,omitempty"`
}

// NonMovementAreaTimeSliceType Named source named complex type "NonMovementAreaTimeSliceType"
type NonMovementAreaTimeSliceType struct {
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

	// generated from group "aixm:NonMovementAreaPropertyGroup" with order 821 depth 2
	Group_NonMovementAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNonMovementAreaExtension" of complex type AbstractExtensionType
	AbstractNonMovementAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNonMovementAreaExtension,omitempty"`
}

// Group_NonMovementAreaPropertyGroup UnNamed source named group "NonMovementAreaPropertyGroup"
type Group_NonMovementAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SurfaceCharacteristicsType Named source named complex type "SurfaceCharacteristicsType"
type SurfaceCharacteristicsType struct {
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

	// generated from group "aixm:SurfaceCharacteristicsPropertyGroup" with order 836 depth 2
	Group_SurfaceCharacteristicsPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceCharacteristicsExtension" of complex type AbstractExtensionType
	AbstractSurfaceCharacteristicsExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceCharacteristicsExtension,omitempty"`
}

// Group_SurfaceCharacteristicsPropertyGroup UnNamed source named group "SurfaceCharacteristicsPropertyGroup"
type Group_SurfaceCharacteristicsPropertyGroup struct {

	// insertion point for fields

	// generated from element "composition" of complex type CodeSurfaceCompositionType
	Composition *aixmdatatypes.CodeSurfaceCompositionType `xml:"composition,omitempty"`

	// generated from element "preparation" of complex type CodeSurfacePreparationType
	Preparation *aixmdatatypes.CodeSurfacePreparationType `xml:"preparation,omitempty"`

	// generated from element "surfaceCondition" of complex type CodeSurfaceConditionType
	SurfaceCondition *aixmdatatypes.CodeSurfaceConditionType `xml:"surfaceCondition,omitempty"`

	// generated from element "classPCN" of complex type ValPCNType
	ClassPCN *aixmdatatypes.ValPCNType `xml:"classPCN,omitempty"`

	// generated from element "pavementTypePCN" of complex type CodePCNPavementType
	PavementTypePCN *aixmdatatypes.CodePCNPavementType `xml:"pavementTypePCN,omitempty"`

	// generated from element "pavementSubgradePCN" of complex type CodePCNSubgradeType
	PavementSubgradePCN *aixmdatatypes.CodePCNSubgradeType `xml:"pavementSubgradePCN,omitempty"`

	// generated from element "maxTyrePressurePCN" of complex type CodePCNTyrePressureType
	MaxTyrePressurePCN *aixmdatatypes.CodePCNTyrePressureType `xml:"maxTyrePressurePCN,omitempty"`

	// generated from element "evaluationMethodPCN" of complex type CodePCNMethodType
	EvaluationMethodPCN *aixmdatatypes.CodePCNMethodType `xml:"evaluationMethodPCN,omitempty"`

	// generated from element "classLCN" of complex type ValLCNType
	ClassLCN *aixmdatatypes.ValLCNType `xml:"classLCN,omitempty"`

	// generated from element "weightSIWL" of complex type ValWeightType
	WeightSIWL *aixmdatatypes.ValWeightType `xml:"weightSIWL,omitempty"`

	// generated from element "tyrePressureSIWL" of complex type ValPressureType
	TyrePressureSIWL *aixmdatatypes.ValPressureType `xml:"tyrePressureSIWL,omitempty"`

	// generated from element "weightAUW" of complex type ValWeightType
	WeightAUW *aixmdatatypes.ValWeightType `xml:"weightAUW,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SurfaceCharacteristicsPropertyType Named source named complex type "SurfaceCharacteristicsPropertyType"
type SurfaceCharacteristicsPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SurfaceCharacteristics" of complex type SurfaceCharacteristicsType
	SurfaceCharacteristics *SurfaceCharacteristicsType `xml:"SurfaceCharacteristics,omitempty"`
}

// SurveyControlPointType Named source named complex type "SurveyControlPointType"
type SurveyControlPointType struct {
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

	// generated from element "timeSlice" of complex type SurveyControlPointTimeSlicePropertyType
	TimeSlice []*SurveyControlPointTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SurveyControlPointTimeSlicePropertyType Named source named complex type "SurveyControlPointTimeSlicePropertyType"
type SurveyControlPointTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SurveyControlPointTimeSlice" of complex type SurveyControlPointTimeSliceType
	SurveyControlPointTimeSlice *SurveyControlPointTimeSliceType `xml:"SurveyControlPointTimeSlice,omitempty"`
}

// SurveyControlPointTimeSliceType Named source named complex type "SurveyControlPointTimeSliceType"
type SurveyControlPointTimeSliceType struct {
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

	// generated from group "aixm:SurveyControlPointPropertyGroup" with order 875 depth 2
	Group_SurveyControlPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurveyControlPointExtension" of complex type AbstractExtensionType
	AbstractSurveyControlPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurveyControlPointExtension,omitempty"`
}

// Group_SurveyControlPointPropertyGroup UnNamed source named group "SurveyControlPointPropertyGroup"
type Group_SurveyControlPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextNameType
	Designator *aixmdatatypes.TextNameType `xml:"designator,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AbstractUsageConditionType Named source named complex type "AbstractUsageConditionType"
type AbstractUsageConditionType struct {
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

// Group_UsageConditionPropertyGroup UnNamed source named group "UsageConditionPropertyGroup"
type Group_UsageConditionPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeUsageLimitationType
	Type_XSD *aixmdatatypes.CodeUsageLimitationType `xml:"type,omitempty"`

	// generated from element "priorPermission" of complex type ValDurationType
	PriorPermission *aixmdatatypes.ValDurationType `xml:"priorPermission,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "selection" of complex type ConditionCombinationPropertyType
	Selection *ConditionCombinationPropertyType `xml:"selection,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// WorkAreaType Named source named complex type "WorkAreaType"
type WorkAreaType struct {
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

	// generated from element "timeSlice" of complex type WorkAreaTimeSlicePropertyType
	TimeSlice []*WorkAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// WorkAreaTimeSlicePropertyType Named source named complex type "WorkAreaTimeSlicePropertyType"
type WorkAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "WorkAreaTimeSlice" of complex type WorkAreaTimeSliceType
	WorkAreaTimeSlice *WorkAreaTimeSliceType `xml:"WorkAreaTimeSlice,omitempty"`
}

// WorkAreaTimeSliceType Named source named complex type "WorkAreaTimeSliceType"
type WorkAreaTimeSliceType struct {
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

	// generated from group "aixm:WorkAreaPropertyGroup" with order 916 depth 2
	Group_WorkAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractWorkAreaExtension" of complex type AbstractExtensionType
	AbstractWorkAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractWorkAreaExtension,omitempty"`
}

// Group_WorkAreaPropertyGroup UnNamed source named group "WorkAreaPropertyGroup"
type Group_WorkAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeWorkAreaType
	Type_XSD *aixmdatatypes.CodeWorkAreaType `xml:"type,omitempty"`

	// generated from element "plannedOperational" of complex type DateType
	PlannedOperational *aixmdatatypes.DateType `xml:"plannedOperational,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "activation" of complex type WorkareaActivityPropertyType
	Activation []*WorkareaActivityPropertyType `xml:"activation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// WorkareaActivityType Named source named complex type "WorkareaActivityType"
type WorkareaActivityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 934 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:WorkareaActivityPropertyGroup" with order 935 depth 2
	Group_WorkareaActivityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractWorkareaActivityExtension" of complex type AbstractExtensionType
	AbstractWorkareaActivityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractWorkareaActivityExtension,omitempty"`
}

// Group_WorkareaActivityPropertyGroup UnNamed source named group "WorkareaActivityPropertyGroup"
type Group_WorkareaActivityPropertyGroup struct {

	// insertion point for fields

	// generated from element "isActive" of complex type CodeYesNoType
	IsActive *aixmdatatypes.CodeYesNoType `xml:"isActive,omitempty"`
}

// WorkareaActivityPropertyType Named source named complex type "WorkareaActivityPropertyType"
type WorkareaActivityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "WorkareaActivity" of complex type WorkareaActivityType
	WorkareaActivity *WorkareaActivityType `xml:"WorkareaActivity,omitempty"`
}

// AircraftStandPropertyType Named source named complex type "AircraftStandPropertyType"
type AircraftStandPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AircraftStandType Named source named complex type "AircraftStandType"
type AircraftStandType struct {
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

	// generated from element "timeSlice" of complex type AircraftStandTimeSlicePropertyType
	TimeSlice []*AircraftStandTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AircraftStandTimeSlicePropertyType Named source named complex type "AircraftStandTimeSlicePropertyType"
type AircraftStandTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AircraftStandTimeSlice" of complex type AircraftStandTimeSliceType
	AircraftStandTimeSlice *AircraftStandTimeSliceType `xml:"AircraftStandTimeSlice,omitempty"`
}

// AircraftStandTimeSliceType Named source named complex type "AircraftStandTimeSliceType"
type AircraftStandTimeSliceType struct {
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

	// generated from group "aixm:AircraftStandPropertyGroup" with order 963 depth 2
	Group_AircraftStandPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_15 `xml:"extension,omitempty"`
}

// A_extension_15 Named source within outer element "extension"
type A_extension_15 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAircraftStandExtension" of complex type AbstractExtensionType
	AbstractAircraftStandExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAircraftStandExtension,omitempty"`
}

// Group_AircraftStandPropertyGroup UnNamed source named group "AircraftStandPropertyGroup"
type Group_AircraftStandPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeAircraftStandType
	Type_XSD *aixmdatatypes.CodeAircraftStandType `xml:"type,omitempty"`

	// generated from element "visualDockingSystem" of complex type CodeVisualDockingGuidanceType
	VisualDockingSystem *aixmdatatypes.CodeVisualDockingGuidanceType `xml:"visualDockingSystem,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "apronLocation" of complex type ApronElementPropertyType
	ApronLocation *ApronElementPropertyType `xml:"apronLocation,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "contaminant" of complex type AircraftStandContaminationPropertyType
	Contaminant []*AircraftStandContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ApronAreaAvailabilityPropertyType
	Availability []*ApronAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// ApronPropertyType Named source named complex type "ApronPropertyType"
type ApronPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// ApronType Named source named complex type "ApronType"
type ApronType struct {
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

	// generated from element "timeSlice" of complex type ApronTimeSlicePropertyType
	TimeSlice []*ApronTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ApronTimeSlicePropertyType Named source named complex type "ApronTimeSlicePropertyType"
type ApronTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ApronTimeSlice" of complex type ApronTimeSliceType
	ApronTimeSlice *ApronTimeSliceType `xml:"ApronTimeSlice,omitempty"`
}

// ApronTimeSliceType Named source named complex type "ApronTimeSliceType"
type ApronTimeSliceType struct {
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

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "contaminant" of complex type ApronContaminationPropertyType
	Contaminant []*ApronContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ApronAreaAvailabilityPropertyType
	Availability []*ApronAreaAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_16 `xml:"extension,omitempty"`
}

// A_extension_16 Named source within outer element "extension"
type A_extension_16 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApronExtension" of complex type AbstractExtensionType
	AbstractApronExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronExtension,omitempty"`
}

// Group_ApronPropertyGroup UnNamed source named group "ApronPropertyGroup"
type Group_ApronPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "contaminant" of complex type ApronContaminationPropertyType
	Contaminant []*ApronContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ApronAreaAvailabilityPropertyType
	Availability []*ApronAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// ApronAreaAvailabilityType Named source named complex type "ApronAreaAvailabilityType"
type ApronAreaAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 1015 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:ApronAreaAvailabilityPropertyGroup" with order 1016 depth 2
	Group_ApronAreaAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_17 `xml:"extension,omitempty"`
}

// A_extension_17 Named source within outer element "extension"
type A_extension_17 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractApronAreaAvailabilityExtension" of complex type AbstractExtensionType
	AbstractApronAreaAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronAreaAvailabilityExtension,omitempty"`
}

// Group_ApronAreaAvailabilityPropertyGroup UnNamed source named group "ApronAreaAvailabilityPropertyGroup"
type Group_ApronAreaAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusAirportType
	OperationalStatus *aixmdatatypes.CodeStatusAirportType `xml:"operationalStatus,omitempty"`

	// generated from element "warning" of complex type CodeAirportWarningType
	Warning *aixmdatatypes.CodeAirportWarningType `xml:"warning,omitempty"`

	// generated from element "usage" of complex type ApronAreaUsagePropertyType
	Usage []*ApronAreaUsagePropertyType `xml:"usage,omitempty"`
}

// ApronAreaAvailabilityPropertyType Named source named complex type "ApronAreaAvailabilityPropertyType"
type ApronAreaAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApronAreaAvailability" of complex type ApronAreaAvailabilityType
	ApronAreaAvailability *ApronAreaAvailabilityType `xml:"ApronAreaAvailability,omitempty"`
}

// ApronAreaUsageType Named source named complex type "ApronAreaUsageType"
type ApronAreaUsageType struct {
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

	// generated from group "aixm:UsageConditionPropertyGroup" with order 1035 depth 2
	Group_UsageConditionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_18 `xml:"extension,omitempty"`
}

// A_extension_18 Named source within outer element "extension"
type A_extension_18 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUsageConditionExtension" of complex type AbstractExtensionType
	AbstractUsageConditionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUsageConditionExtension,omitempty"`

	// generated from element "AbstractApronAreaUsageExtension" of complex type AbstractExtensionType
	AbstractApronAreaUsageExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronAreaUsageExtension,omitempty"`
}

// Group_ApronAreaUsagePropertyGroup UnNamed source named group "ApronAreaUsagePropertyGroup"
type Group_ApronAreaUsagePropertyGroup struct {

	// insertion point for fields
}

// ApronAreaUsagePropertyType Named source named complex type "ApronAreaUsagePropertyType"
type ApronAreaUsagePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApronAreaUsage" of complex type ApronAreaUsageType
	ApronAreaUsage *ApronAreaUsageType `xml:"ApronAreaUsage,omitempty"`
}

// ApronElementPropertyType Named source named complex type "ApronElementPropertyType"
type ApronElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// ApronElementType Named source named complex type "ApronElementType"
type ApronElementType struct {
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

	// generated from element "timeSlice" of complex type ApronElementTimeSlicePropertyType
	TimeSlice []*ApronElementTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ApronElementTimeSlicePropertyType Named source named complex type "ApronElementTimeSlicePropertyType"
type ApronElementTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ApronElementTimeSlice" of complex type ApronElementTimeSliceType
	ApronElementTimeSlice *ApronElementTimeSliceType `xml:"ApronElementTimeSlice,omitempty"`
}

// ApronElementTimeSliceType Named source named complex type "ApronElementTimeSliceType"
type ApronElementTimeSliceType struct {
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

	// generated from group "aixm:ApronElementPropertyGroup" with order 1063 depth 2
	Group_ApronElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_19 `xml:"extension,omitempty"`
}

// A_extension_19 Named source within outer element "extension"
type A_extension_19 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApronElementExtension" of complex type AbstractExtensionType
	AbstractApronElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronElementExtension,omitempty"`
}

// Group_ApronElementPropertyGroup UnNamed source named group "ApronElementPropertyGroup"
type Group_ApronElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeApronElementType
	Type_XSD *aixmdatatypes.CodeApronElementType `xml:"type,omitempty"`

	// generated from element "jetwayAvailability" of complex type CodeYesNoType
	JetwayAvailability *aixmdatatypes.CodeYesNoType `xml:"jetwayAvailability,omitempty"`

	// generated from element "towingAvailability" of complex type CodeYesNoType
	TowingAvailability *aixmdatatypes.CodeYesNoType `xml:"towingAvailability,omitempty"`

	// generated from element "dockingAvailability" of complex type CodeYesNoType
	DockingAvailability *aixmdatatypes.CodeYesNoType `xml:"dockingAvailability,omitempty"`

	// generated from element "groundPowerAvailability" of complex type CodeYesNoType
	GroundPowerAvailability *aixmdatatypes.CodeYesNoType `xml:"groundPowerAvailability,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "associatedApron" of complex type ApronPropertyType
	AssociatedApron *ApronPropertyType `xml:"associatedApron,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "supplyService" of complex type AirportSuppliesServicePropertyType
	SupplyService []*AirportSuppliesServicePropertyType `xml:"supplyService,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ApronAreaAvailabilityPropertyType
	Availability []*ApronAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// DeicingAreaPropertyType Named source named complex type "DeicingAreaPropertyType"
type DeicingAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// DeicingAreaType Named source named complex type "DeicingAreaType"
type DeicingAreaType struct {
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

	// generated from element "timeSlice" of complex type DeicingAreaTimeSlicePropertyType
	TimeSlice []*DeicingAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DeicingAreaTimeSlicePropertyType Named source named complex type "DeicingAreaTimeSlicePropertyType"
type DeicingAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DeicingAreaTimeSlice" of complex type DeicingAreaTimeSliceType
	DeicingAreaTimeSlice *DeicingAreaTimeSliceType `xml:"DeicingAreaTimeSlice,omitempty"`
}

// DeicingAreaTimeSliceType Named source named complex type "DeicingAreaTimeSliceType"
type DeicingAreaTimeSliceType struct {
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

	// generated from group "aixm:DeicingAreaPropertyGroup" with order 1099 depth 2
	Group_DeicingAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_20 `xml:"extension,omitempty"`
}

// A_extension_20 Named source within outer element "extension"
type A_extension_20 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDeicingAreaExtension" of complex type AbstractExtensionType
	AbstractDeicingAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDeicingAreaExtension,omitempty"`
}

// Group_DeicingAreaPropertyGroup UnNamed source named group "DeicingAreaPropertyGroup"
type Group_DeicingAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "associatedApron" of complex type ApronPropertyType
	AssociatedApron *ApronPropertyType `xml:"associatedApron,omitempty"`

	// generated from element "taxiwayLocation" of complex type TaxiwayPropertyType
	TaxiwayLocation *TaxiwayPropertyType `xml:"taxiwayLocation,omitempty"`

	// generated from element "standLocation" of complex type AircraftStandPropertyType
	StandLocation *AircraftStandPropertyType `xml:"standLocation,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ApronAreaAvailabilityPropertyType
	Availability []*ApronAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// PassengerLoadingBridgeType Named source named complex type "PassengerLoadingBridgeType"
type PassengerLoadingBridgeType struct {
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

	// generated from element "timeSlice" of complex type PassengerLoadingBridgeTimeSlicePropertyType
	TimeSlice []*PassengerLoadingBridgeTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// PassengerLoadingBridgeTimeSlicePropertyType Named source named complex type "PassengerLoadingBridgeTimeSlicePropertyType"
type PassengerLoadingBridgeTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "PassengerLoadingBridgeTimeSlice" of complex type PassengerLoadingBridgeTimeSliceType
	PassengerLoadingBridgeTimeSlice *PassengerLoadingBridgeTimeSliceType `xml:"PassengerLoadingBridgeTimeSlice,omitempty"`
}

// PassengerLoadingBridgeTimeSliceType Named source named complex type "PassengerLoadingBridgeTimeSliceType"
type PassengerLoadingBridgeTimeSliceType struct {
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

	// generated from group "aixm:PassengerLoadingBridgePropertyGroup" with order 1129 depth 2
	Group_PassengerLoadingBridgePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_21 `xml:"extension,omitempty"`
}

// A_extension_21 Named source within outer element "extension"
type A_extension_21 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPassengerLoadingBridgeExtension" of complex type AbstractExtensionType
	AbstractPassengerLoadingBridgeExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPassengerLoadingBridgeExtension,omitempty"`
}

// Group_PassengerLoadingBridgePropertyGroup UnNamed source named group "PassengerLoadingBridgePropertyGroup"
type Group_PassengerLoadingBridgePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeLoadingBridgeType
	Type_XSD *aixmdatatypes.CodeLoadingBridgeType `xml:"type,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "associatedStand" of complex type AircraftStandPropertyType
	AssociatedStand []*AircraftStandPropertyType `xml:"associatedStand,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RoadType Named source named complex type "RoadType"
type RoadType struct {
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

	// generated from element "timeSlice" of complex type RoadTimeSlicePropertyType
	TimeSlice []*RoadTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RoadTimeSlicePropertyType Named source named complex type "RoadTimeSlicePropertyType"
type RoadTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RoadTimeSlice" of complex type RoadTimeSliceType
	RoadTimeSlice *RoadTimeSliceType `xml:"RoadTimeSlice,omitempty"`
}

// RoadTimeSliceType Named source named complex type "RoadTimeSliceType"
type RoadTimeSliceType struct {
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

	// generated from group "aixm:RoadPropertyGroup" with order 1156 depth 2
	Group_RoadPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_22 `xml:"extension,omitempty"`
}

// A_extension_22 Named source within outer element "extension"
type A_extension_22 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRoadExtension" of complex type AbstractExtensionType
	AbstractRoadExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRoadExtension,omitempty"`
}

// Group_RoadPropertyGroup UnNamed source named group "RoadPropertyGroup"
type Group_RoadPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextNameType
	Designator *aixmdatatypes.TextNameType `xml:"designator,omitempty"`

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`

	// generated from element "type" of complex type CodeRoadType
	Type_XSD *aixmdatatypes.CodeRoadType `xml:"type,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "associatedAirport" of complex type AirportHeliportPropertyType
	AssociatedAirport *AirportHeliportPropertyType `xml:"associatedAirport,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "accessibleStand" of complex type AircraftStandPropertyType
	AccessibleStand []*AircraftStandPropertyType `xml:"accessibleStand,omitempty"`

	// generated from element "surfaceExtent" of complex type ElevatedSurfacePropertyType
	SurfaceExtent *ElevatedSurfacePropertyType `xml:"surfaceExtent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

// TouchDownLiftOffType Named source named complex type "TouchDownLiftOffType"
type TouchDownLiftOffType struct {
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

	// generated from element "timeSlice" of complex type TouchDownLiftOffTimeSlicePropertyType
	TimeSlice []*TouchDownLiftOffTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TouchDownLiftOffTimeSlicePropertyType Named source named complex type "TouchDownLiftOffTimeSlicePropertyType"
type TouchDownLiftOffTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TouchDownLiftOffTimeSlice" of complex type TouchDownLiftOffTimeSliceType
	TouchDownLiftOffTimeSlice *TouchDownLiftOffTimeSliceType `xml:"TouchDownLiftOffTimeSlice,omitempty"`
}

// TouchDownLiftOffTimeSliceType Named source named complex type "TouchDownLiftOffTimeSliceType"
type TouchDownLiftOffTimeSliceType struct {
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

	// generated from group "aixm:TouchDownLiftOffPropertyGroup" with order 1188 depth 2
	Group_TouchDownLiftOffPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_23 `xml:"extension,omitempty"`
}

// A_extension_23 Named source within outer element "extension"
type A_extension_23 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTouchDownLiftOffExtension" of complex type AbstractExtensionType
	AbstractTouchDownLiftOffExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTouchDownLiftOffExtension,omitempty"`
}

// Group_TouchDownLiftOffPropertyGroup UnNamed source named group "TouchDownLiftOffPropertyGroup"
type Group_TouchDownLiftOffPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "slope" of complex type ValSlopeType
	Slope *aixmdatatypes.ValSlopeType `xml:"slope,omitempty"`

	// generated from element "helicopterClass" of complex type CodeHelicopterPerformanceType
	HelicopterClass *aixmdatatypes.CodeHelicopterPerformanceType `xml:"helicopterClass,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "aimingPoint" of complex type ElevatedPointPropertyType
	AimingPoint *ElevatedPointPropertyType `xml:"aimingPoint,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "approachTakeOffArea" of complex type RunwayPropertyType
	ApproachTakeOffArea *RunwayPropertyType `xml:"approachTakeOffArea,omitempty"`

	// generated from element "contaminant" of complex type TouchDownLiftOffContaminationPropertyType
	Contaminant []*TouchDownLiftOffContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// TouchDownLiftOffSafeAreaType Named source named complex type "TouchDownLiftOffSafeAreaType"
type TouchDownLiftOffSafeAreaType struct {
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

	// generated from element "timeSlice" of complex type TouchDownLiftOffSafeAreaTimeSlicePropertyType
	TimeSlice []*TouchDownLiftOffSafeAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TouchDownLiftOffSafeAreaTimeSlicePropertyType Named source named complex type "TouchDownLiftOffSafeAreaTimeSlicePropertyType"
type TouchDownLiftOffSafeAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TouchDownLiftOffSafeAreaTimeSlice" of complex type TouchDownLiftOffSafeAreaTimeSliceType
	TouchDownLiftOffSafeAreaTimeSlice *TouchDownLiftOffSafeAreaTimeSliceType `xml:"TouchDownLiftOffSafeAreaTimeSlice,omitempty"`
}

// TouchDownLiftOffSafeAreaTimeSliceType Named source named complex type "TouchDownLiftOffSafeAreaTimeSliceType"
type TouchDownLiftOffSafeAreaTimeSliceType struct {
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

	// generated from group "aixm:TouchDownLiftOffSafeAreaPropertyGroup" with order 1225 depth 2
	Group_TouchDownLiftOffSafeAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_24 `xml:"extension,omitempty"`
}

// A_extension_24 Named source within outer element "extension"
type A_extension_24 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTouchDownLiftOffSafeAreaExtension" of complex type AbstractExtensionType
	AbstractTouchDownLiftOffSafeAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTouchDownLiftOffSafeAreaExtension,omitempty"`

	// generated from element "AbstractAirportHeliportProtectionAreaExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportProtectionAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportProtectionAreaExtension,omitempty"`
}

// Group_TouchDownLiftOffSafeAreaPropertyGroup UnNamed source named group "TouchDownLiftOffSafeAreaPropertyGroup"
type Group_TouchDownLiftOffSafeAreaPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportHeliportProtectionAreaPropertyGroup" with order 1234 depth 2
	Group_AirportHeliportProtectionAreaPropertyGroup

	// generated from element "protectedTouchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	ProtectedTouchDownLiftOff *TouchDownLiftOffPropertyType `xml:"protectedTouchDownLiftOff,omitempty"`
}

// ApproachLightingSystemType Named source named complex type "ApproachLightingSystemType"
type ApproachLightingSystemType struct {
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

	// generated from element "timeSlice" of complex type ApproachLightingSystemTimeSlicePropertyType
	TimeSlice []*ApproachLightingSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ApproachLightingSystemTimeSlicePropertyType Named source named complex type "ApproachLightingSystemTimeSlicePropertyType"
type ApproachLightingSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ApproachLightingSystemTimeSlice" of complex type ApproachLightingSystemTimeSliceType
	ApproachLightingSystemTimeSlice *ApproachLightingSystemTimeSliceType `xml:"ApproachLightingSystemTimeSlice,omitempty"`
}

// ApproachLightingSystemTimeSliceType Named source named complex type "ApproachLightingSystemTimeSliceType"
type ApproachLightingSystemTimeSliceType struct {
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

	// generated from group "aixm:ApproachLightingSystemPropertyGroup" with order 1251 depth 2
	Group_ApproachLightingSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_25 `xml:"extension,omitempty"`
}

// A_extension_25 Named source within outer element "extension"
type A_extension_25 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApproachLightingSystemExtension" of complex type AbstractExtensionType
	AbstractApproachLightingSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApproachLightingSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_ApproachLightingSystemPropertyGroup UnNamed source named group "ApproachLightingSystemPropertyGroup"
type Group_ApproachLightingSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1260 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "classICAO" of complex type CodeApproachLightingICAOType
	ClassICAO *aixmdatatypes.CodeApproachLightingICAOType `xml:"classICAO,omitempty"`

	// generated from element "type" of complex type CodeApproachLightingType
	Type_XSD *aixmdatatypes.CodeApproachLightingType `xml:"type,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "sequencedFlashing" of complex type CodeYesNoType
	SequencedFlashing *aixmdatatypes.CodeYesNoType `xml:"sequencedFlashing,omitempty"`

	// generated from element "alignmentIndicator" of complex type CodeYesNoType
	AlignmentIndicator *aixmdatatypes.CodeYesNoType `xml:"alignmentIndicator,omitempty"`

	// generated from element "servedRunwayDirection" of complex type RunwayDirectionPropertyType
	ServedRunwayDirection *RunwayDirectionPropertyType `xml:"servedRunwayDirection,omitempty"`
}

// ApronLightSystemType Named source named complex type "ApronLightSystemType"
type ApronLightSystemType struct {
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

	// generated from element "timeSlice" of complex type ApronLightSystemTimeSlicePropertyType
	TimeSlice []*ApronLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ApronLightSystemTimeSlicePropertyType Named source named complex type "ApronLightSystemTimeSlicePropertyType"
type ApronLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ApronLightSystemTimeSlice" of complex type ApronLightSystemTimeSliceType
	ApronLightSystemTimeSlice *ApronLightSystemTimeSliceType `xml:"ApronLightSystemTimeSlice,omitempty"`
}

// ApronLightSystemTimeSliceType Named source named complex type "ApronLightSystemTimeSliceType"
type ApronLightSystemTimeSliceType struct {
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

	// generated from group "aixm:ApronLightSystemPropertyGroup" with order 1282 depth 2
	Group_ApronLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_26 `xml:"extension,omitempty"`
}

// A_extension_26 Named source within outer element "extension"
type A_extension_26 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApronLightSystemExtension" of complex type AbstractExtensionType
	AbstractApronLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_ApronLightSystemPropertyGroup UnNamed source named group "ApronLightSystemPropertyGroup"
type Group_ApronLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1291 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "position" of complex type CodeApronSectionType
	Position *aixmdatatypes.CodeApronSectionType `xml:"position,omitempty"`

	// generated from element "lightedApron" of complex type ApronPropertyType
	LightedApron *ApronPropertyType `xml:"lightedApron,omitempty"`
}

// GroundLightingAvailabilityType Named source named complex type "GroundLightingAvailabilityType"
type GroundLightingAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 1298 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:GroundLightingAvailabilityPropertyGroup" with order 1299 depth 2
	Group_GroundLightingAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_27 `xml:"extension,omitempty"`
}

// A_extension_27 Named source within outer element "extension"
type A_extension_27 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractGroundLightingAvailabilityExtension" of complex type AbstractExtensionType
	AbstractGroundLightingAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightingAvailabilityExtension,omitempty"`
}

// Group_GroundLightingAvailabilityPropertyGroup UnNamed source named group "GroundLightingAvailabilityPropertyGroup"
type Group_GroundLightingAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusOperationsType
	OperationalStatus *aixmdatatypes.CodeStatusOperationsType `xml:"operationalStatus,omitempty"`
}

// GroundLightingAvailabilityPropertyType Named source named complex type "GroundLightingAvailabilityPropertyType"
type GroundLightingAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "GroundLightingAvailability" of complex type GroundLightingAvailabilityType
	GroundLightingAvailability *GroundLightingAvailabilityType `xml:"GroundLightingAvailability,omitempty"`
}

// AbstractGroundLightSystemType Named source named complex type "AbstractGroundLightSystemType"
type AbstractGroundLightSystemType struct {
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

// Group_GroundLightSystemPropertyGroup UnNamed source named group "GroundLightSystemPropertyGroup"
type Group_GroundLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from element "emergencyLighting" of complex type CodeYesNoType
	EmergencyLighting *aixmdatatypes.CodeYesNoType `xml:"emergencyLighting,omitempty"`

	// generated from element "intensityLevel" of complex type CodeLightIntensityType
	IntensityLevel *aixmdatatypes.CodeLightIntensityType `xml:"intensityLevel,omitempty"`

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "element" of complex type LightElementPropertyType
	Element []*LightElementPropertyType `xml:"element,omitempty"`

	// generated from element "availability" of complex type GroundLightingAvailabilityPropertyType
	Availability []*GroundLightingAvailabilityPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// GuidanceLineLightSystemType Named source named complex type "GuidanceLineLightSystemType"
type GuidanceLineLightSystemType struct {
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

	// generated from element "timeSlice" of complex type GuidanceLineLightSystemTimeSlicePropertyType
	TimeSlice []*GuidanceLineLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GuidanceLineLightSystemTimeSlicePropertyType Named source named complex type "GuidanceLineLightSystemTimeSlicePropertyType"
type GuidanceLineLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GuidanceLineLightSystemTimeSlice" of complex type GuidanceLineLightSystemTimeSliceType
	GuidanceLineLightSystemTimeSlice *GuidanceLineLightSystemTimeSliceType `xml:"GuidanceLineLightSystemTimeSlice,omitempty"`
}

// GuidanceLineLightSystemTimeSliceType Named source named complex type "GuidanceLineLightSystemTimeSliceType"
type GuidanceLineLightSystemTimeSliceType struct {
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

	// generated from group "aixm:GuidanceLineLightSystemPropertyGroup" with order 1341 depth 2
	Group_GuidanceLineLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_28 `xml:"extension,omitempty"`
}

// A_extension_28 Named source within outer element "extension"
type A_extension_28 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGuidanceLineLightSystemExtension" of complex type AbstractExtensionType
	AbstractGuidanceLineLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGuidanceLineLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_GuidanceLineLightSystemPropertyGroup UnNamed source named group "GuidanceLineLightSystemPropertyGroup"
type Group_GuidanceLineLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1350 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "lightedGuidanceLine" of complex type GuidanceLinePropertyType
	LightedGuidanceLine *GuidanceLinePropertyType `xml:"lightedGuidanceLine,omitempty"`
}

// RunwayDirectionLightSystemType Named source named complex type "RunwayDirectionLightSystemType"
type RunwayDirectionLightSystemType struct {
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

	// generated from element "timeSlice" of complex type RunwayDirectionLightSystemTimeSlicePropertyType
	TimeSlice []*RunwayDirectionLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayDirectionLightSystemTimeSlicePropertyType Named source named complex type "RunwayDirectionLightSystemTimeSlicePropertyType"
type RunwayDirectionLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayDirectionLightSystemTimeSlice" of complex type RunwayDirectionLightSystemTimeSliceType
	RunwayDirectionLightSystemTimeSlice *RunwayDirectionLightSystemTimeSliceType `xml:"RunwayDirectionLightSystemTimeSlice,omitempty"`
}

// RunwayDirectionLightSystemTimeSliceType Named source named complex type "RunwayDirectionLightSystemTimeSliceType"
type RunwayDirectionLightSystemTimeSliceType struct {
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

	// generated from group "aixm:RunwayDirectionLightSystemPropertyGroup" with order 1386 depth 2
	Group_RunwayDirectionLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_29 `xml:"extension,omitempty"`
}

// A_extension_29 Named source within outer element "extension"
type A_extension_29 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayDirectionLightSystemExtension" of complex type AbstractExtensionType
	AbstractRunwayDirectionLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayDirectionLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_RunwayDirectionLightSystemPropertyGroup UnNamed source named group "RunwayDirectionLightSystemPropertyGroup"
type Group_RunwayDirectionLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1395 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "position" of complex type CodeRunwaySectionType
	Position *aixmdatatypes.CodeRunwaySectionType `xml:"position,omitempty"`

	// generated from element "associatedRunwayDirection" of complex type RunwayDirectionPropertyType
	AssociatedRunwayDirection *RunwayDirectionPropertyType `xml:"associatedRunwayDirection,omitempty"`
}

// RunwayProtectAreaLightSystemType Named source named complex type "RunwayProtectAreaLightSystemType"
type RunwayProtectAreaLightSystemType struct {
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

	// generated from element "timeSlice" of complex type RunwayProtectAreaLightSystemTimeSlicePropertyType
	TimeSlice []*RunwayProtectAreaLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayProtectAreaLightSystemTimeSlicePropertyType Named source named complex type "RunwayProtectAreaLightSystemTimeSlicePropertyType"
type RunwayProtectAreaLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayProtectAreaLightSystemTimeSlice" of complex type RunwayProtectAreaLightSystemTimeSliceType
	RunwayProtectAreaLightSystemTimeSlice *RunwayProtectAreaLightSystemTimeSliceType `xml:"RunwayProtectAreaLightSystemTimeSlice,omitempty"`
}

// RunwayProtectAreaLightSystemTimeSliceType Named source named complex type "RunwayProtectAreaLightSystemTimeSliceType"
type RunwayProtectAreaLightSystemTimeSliceType struct {
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

	// generated from group "aixm:RunwayProtectAreaLightSystemPropertyGroup" with order 1413 depth 2
	Group_RunwayProtectAreaLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_30 `xml:"extension,omitempty"`
}

// A_extension_30 Named source within outer element "extension"
type A_extension_30 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayProtectAreaLightSystemExtension" of complex type AbstractExtensionType
	AbstractRunwayProtectAreaLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayProtectAreaLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_RunwayProtectAreaLightSystemPropertyGroup UnNamed source named group "RunwayProtectAreaLightSystemPropertyGroup"
type Group_RunwayProtectAreaLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1422 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "position" of complex type CodeProtectAreaSectionType
	Position *aixmdatatypes.CodeProtectAreaSectionType `xml:"position,omitempty"`

	// generated from element "lightedArea" of complex type RunwayProtectAreaPropertyType
	LightedArea *RunwayProtectAreaPropertyType `xml:"lightedArea,omitempty"`
}

// TaxiHoldingPositionLightSystemType Named source named complex type "TaxiHoldingPositionLightSystemType"
type TaxiHoldingPositionLightSystemType struct {
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

	// generated from element "timeSlice" of complex type TaxiHoldingPositionLightSystemTimeSlicePropertyType
	TimeSlice []*TaxiHoldingPositionLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiHoldingPositionLightSystemTimeSlicePropertyType Named source named complex type "TaxiHoldingPositionLightSystemTimeSlicePropertyType"
type TaxiHoldingPositionLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiHoldingPositionLightSystemTimeSlice" of complex type TaxiHoldingPositionLightSystemTimeSliceType
	TaxiHoldingPositionLightSystemTimeSlice *TaxiHoldingPositionLightSystemTimeSliceType `xml:"TaxiHoldingPositionLightSystemTimeSlice,omitempty"`
}

// TaxiHoldingPositionLightSystemTimeSliceType Named source named complex type "TaxiHoldingPositionLightSystemTimeSliceType"
type TaxiHoldingPositionLightSystemTimeSliceType struct {
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

	// generated from group "aixm:TaxiHoldingPositionLightSystemPropertyGroup" with order 1440 depth 2
	Group_TaxiHoldingPositionLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_31 `xml:"extension,omitempty"`
}

// A_extension_31 Named source within outer element "extension"
type A_extension_31 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiHoldingPositionLightSystemExtension" of complex type AbstractExtensionType
	AbstractTaxiHoldingPositionLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiHoldingPositionLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_TaxiHoldingPositionLightSystemPropertyGroup UnNamed source named group "TaxiHoldingPositionLightSystemPropertyGroup"
type Group_TaxiHoldingPositionLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1449 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "type" of complex type CodeLightHoldingPositionType
	Type_XSD *aixmdatatypes.CodeLightHoldingPositionType `xml:"type,omitempty"`

	// generated from element "taxiHolding" of complex type TaxiHoldingPositionPropertyType
	TaxiHolding *TaxiHoldingPositionPropertyType `xml:"taxiHolding,omitempty"`
}

// TaxiwayLightSystemType Named source named complex type "TaxiwayLightSystemType"
type TaxiwayLightSystemType struct {
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

	// generated from element "timeSlice" of complex type TaxiwayLightSystemTimeSlicePropertyType
	TimeSlice []*TaxiwayLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiwayLightSystemTimeSlicePropertyType Named source named complex type "TaxiwayLightSystemTimeSlicePropertyType"
type TaxiwayLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiwayLightSystemTimeSlice" of complex type TaxiwayLightSystemTimeSliceType
	TaxiwayLightSystemTimeSlice *TaxiwayLightSystemTimeSliceType `xml:"TaxiwayLightSystemTimeSlice,omitempty"`
}

// TaxiwayLightSystemTimeSliceType Named source named complex type "TaxiwayLightSystemTimeSliceType"
type TaxiwayLightSystemTimeSliceType struct {
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

	// generated from group "aixm:TaxiwayLightSystemPropertyGroup" with order 1467 depth 2
	Group_TaxiwayLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_32 `xml:"extension,omitempty"`
}

// A_extension_32 Named source within outer element "extension"
type A_extension_32 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiwayLightSystemExtension" of complex type AbstractExtensionType
	AbstractTaxiwayLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiwayLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_TaxiwayLightSystemPropertyGroup UnNamed source named group "TaxiwayLightSystemPropertyGroup"
type Group_TaxiwayLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1476 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "position" of complex type CodeTaxiwaySectionType
	Position *aixmdatatypes.CodeTaxiwaySectionType `xml:"position,omitempty"`

	// generated from element "lightedTaxiway" of complex type TaxiwayPropertyType
	LightedTaxiway *TaxiwayPropertyType `xml:"lightedTaxiway,omitempty"`
}

// TouchDownLiftOffLightSystemType Named source named complex type "TouchDownLiftOffLightSystemType"
type TouchDownLiftOffLightSystemType struct {
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

	// generated from element "timeSlice" of complex type TouchDownLiftOffLightSystemTimeSlicePropertyType
	TimeSlice []*TouchDownLiftOffLightSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TouchDownLiftOffLightSystemTimeSlicePropertyType Named source named complex type "TouchDownLiftOffLightSystemTimeSlicePropertyType"
type TouchDownLiftOffLightSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TouchDownLiftOffLightSystemTimeSlice" of complex type TouchDownLiftOffLightSystemTimeSliceType
	TouchDownLiftOffLightSystemTimeSlice *TouchDownLiftOffLightSystemTimeSliceType `xml:"TouchDownLiftOffLightSystemTimeSlice,omitempty"`
}

// TouchDownLiftOffLightSystemTimeSliceType Named source named complex type "TouchDownLiftOffLightSystemTimeSliceType"
type TouchDownLiftOffLightSystemTimeSliceType struct {
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

	// generated from group "aixm:TouchDownLiftOffLightSystemPropertyGroup" with order 1494 depth 2
	Group_TouchDownLiftOffLightSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_33 `xml:"extension,omitempty"`
}

// A_extension_33 Named source within outer element "extension"
type A_extension_33 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTouchDownLiftOffLightSystemExtension" of complex type AbstractExtensionType
	AbstractTouchDownLiftOffLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTouchDownLiftOffLightSystemExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_TouchDownLiftOffLightSystemPropertyGroup UnNamed source named group "TouchDownLiftOffLightSystemPropertyGroup"
type Group_TouchDownLiftOffLightSystemPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 1503 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "position" of complex type CodeTLOFSectionType
	Position *aixmdatatypes.CodeTLOFSectionType `xml:"position,omitempty"`

	// generated from element "lightedTouchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	LightedTouchDownLiftOff *TouchDownLiftOffPropertyType `xml:"lightedTouchDownLiftOff,omitempty"`
}

// AirportProtectionAreaMarkingType Named source named complex type "AirportProtectionAreaMarkingType"
type AirportProtectionAreaMarkingType struct {
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

	// generated from element "timeSlice" of complex type AirportProtectionAreaMarkingTimeSlicePropertyType
	TimeSlice []*AirportProtectionAreaMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AirportProtectionAreaMarkingTimeSlicePropertyType Named source named complex type "AirportProtectionAreaMarkingTimeSlicePropertyType"
type AirportProtectionAreaMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AirportProtectionAreaMarkingTimeSlice" of complex type AirportProtectionAreaMarkingTimeSliceType
	AirportProtectionAreaMarkingTimeSlice *AirportProtectionAreaMarkingTimeSliceType `xml:"AirportProtectionAreaMarkingTimeSlice,omitempty"`
}

// AirportProtectionAreaMarkingTimeSliceType Named source named complex type "AirportProtectionAreaMarkingTimeSliceType"
type AirportProtectionAreaMarkingTimeSliceType struct {
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

	// generated from group "aixm:AirportProtectionAreaMarkingPropertyGroup" with order 1521 depth 2
	Group_AirportProtectionAreaMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_34 `xml:"extension,omitempty"`
}

// A_extension_34 Named source within outer element "extension"
type A_extension_34 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAirportProtectionAreaMarkingExtension" of complex type AbstractExtensionType
	AbstractAirportProtectionAreaMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportProtectionAreaMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_AirportProtectionAreaMarkingPropertyGroup UnNamed source named group "AirportProtectionAreaMarkingPropertyGroup"
type Group_AirportProtectionAreaMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1530 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markingLocation" of complex type CodeProtectAreaSectionType
	MarkingLocation *aixmdatatypes.CodeProtectAreaSectionType `xml:"markingLocation,omitempty"`

	// generated from element "markedProtectionArea" of complex type AirportHeliportProtectionAreaPropertyType
	MarkedProtectionArea *AirportHeliportProtectionAreaPropertyType `xml:"markedProtectionArea,omitempty"`
}

// ApronMarkingType Named source named complex type "ApronMarkingType"
type ApronMarkingType struct {
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

	// generated from element "timeSlice" of complex type ApronMarkingTimeSlicePropertyType
	TimeSlice []*ApronMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ApronMarkingTimeSlicePropertyType Named source named complex type "ApronMarkingTimeSlicePropertyType"
type ApronMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ApronMarkingTimeSlice" of complex type ApronMarkingTimeSliceType
	ApronMarkingTimeSlice *ApronMarkingTimeSliceType `xml:"ApronMarkingTimeSlice,omitempty"`
}

// ApronMarkingTimeSliceType Named source named complex type "ApronMarkingTimeSliceType"
type ApronMarkingTimeSliceType struct {
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

	// generated from group "aixm:ApronMarkingPropertyGroup" with order 1548 depth 2
	Group_ApronMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_35 `xml:"extension,omitempty"`
}

// A_extension_35 Named source within outer element "extension"
type A_extension_35 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractApronMarkingExtension" of complex type AbstractExtensionType
	AbstractApronMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_ApronMarkingPropertyGroup UnNamed source named group "ApronMarkingPropertyGroup"
type Group_ApronMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1557 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markingLocation" of complex type CodeApronSectionType
	MarkingLocation *aixmdatatypes.CodeApronSectionType `xml:"markingLocation,omitempty"`

	// generated from element "markedApron" of complex type ApronPropertyType
	MarkedApron *ApronPropertyType `xml:"markedApron,omitempty"`
}

// DeicingAreaMarkingType Named source named complex type "DeicingAreaMarkingType"
type DeicingAreaMarkingType struct {
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

	// generated from element "timeSlice" of complex type DeicingAreaMarkingTimeSlicePropertyType
	TimeSlice []*DeicingAreaMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DeicingAreaMarkingTimeSlicePropertyType Named source named complex type "DeicingAreaMarkingTimeSlicePropertyType"
type DeicingAreaMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DeicingAreaMarkingTimeSlice" of complex type DeicingAreaMarkingTimeSliceType
	DeicingAreaMarkingTimeSlice *DeicingAreaMarkingTimeSliceType `xml:"DeicingAreaMarkingTimeSlice,omitempty"`
}

// DeicingAreaMarkingTimeSliceType Named source named complex type "DeicingAreaMarkingTimeSliceType"
type DeicingAreaMarkingTimeSliceType struct {
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

	// generated from group "aixm:DeicingAreaMarkingPropertyGroup" with order 1575 depth 2
	Group_DeicingAreaMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_36 `xml:"extension,omitempty"`
}

// A_extension_36 Named source within outer element "extension"
type A_extension_36 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDeicingAreaMarkingExtension" of complex type AbstractExtensionType
	AbstractDeicingAreaMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDeicingAreaMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_DeicingAreaMarkingPropertyGroup UnNamed source named group "DeicingAreaMarkingPropertyGroup"
type Group_DeicingAreaMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1584 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markedDeicingArea" of complex type DeicingAreaPropertyType
	MarkedDeicingArea *DeicingAreaPropertyType `xml:"markedDeicingArea,omitempty"`
}

// GuidanceLineMarkingType Named source named complex type "GuidanceLineMarkingType"
type GuidanceLineMarkingType struct {
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

	// generated from element "timeSlice" of complex type GuidanceLineMarkingTimeSlicePropertyType
	TimeSlice []*GuidanceLineMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GuidanceLineMarkingTimeSlicePropertyType Named source named complex type "GuidanceLineMarkingTimeSlicePropertyType"
type GuidanceLineMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GuidanceLineMarkingTimeSlice" of complex type GuidanceLineMarkingTimeSliceType
	GuidanceLineMarkingTimeSlice *GuidanceLineMarkingTimeSliceType `xml:"GuidanceLineMarkingTimeSlice,omitempty"`
}

// GuidanceLineMarkingTimeSliceType Named source named complex type "GuidanceLineMarkingTimeSliceType"
type GuidanceLineMarkingTimeSliceType struct {
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

	// generated from group "aixm:GuidanceLineMarkingPropertyGroup" with order 1601 depth 2
	Group_GuidanceLineMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_37 `xml:"extension,omitempty"`
}

// A_extension_37 Named source within outer element "extension"
type A_extension_37 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGuidanceLineMarkingExtension" of complex type AbstractExtensionType
	AbstractGuidanceLineMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGuidanceLineMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_GuidanceLineMarkingPropertyGroup UnNamed source named group "GuidanceLineMarkingPropertyGroup"
type Group_GuidanceLineMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1610 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markedGuidanceLine" of complex type GuidanceLinePropertyType
	MarkedGuidanceLine *GuidanceLinePropertyType `xml:"markedGuidanceLine,omitempty"`
}

// AbstractMarkingType Named source named complex type "AbstractMarkingType"
type AbstractMarkingType struct {
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

// Group_MarkingPropertyGroup UnNamed source named group "MarkingPropertyGroup"
type Group_MarkingPropertyGroup struct {

	// insertion point for fields

	// generated from element "markingICAOStandard" of complex type CodeYesNoType
	MarkingICAOStandard *aixmdatatypes.CodeYesNoType `xml:"markingICAOStandard,omitempty"`

	// generated from element "condition" of complex type CodeMarkingConditionType
	Condition *aixmdatatypes.CodeMarkingConditionType `xml:"condition,omitempty"`

	// generated from element "element" of complex type MarkingElementPropertyType
	Element []*MarkingElementPropertyType `xml:"element,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MarkingElementType Named source named complex type "MarkingElementType"
type MarkingElementType struct {
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

	// generated from group "aixm:MarkingElementPropertyGroup" with order 1628 depth 2
	Group_MarkingElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_38 `xml:"extension,omitempty"`
}

// A_extension_38 Named source within outer element "extension"
type A_extension_38 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMarkingElementExtension" of complex type AbstractExtensionType
	AbstractMarkingElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingElementExtension,omitempty"`
}

// Group_MarkingElementPropertyGroup UnNamed source named group "MarkingElementPropertyGroup"
type Group_MarkingElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "style" of complex type CodeMarkingStyleType
	Style *aixmdatatypes.CodeMarkingStyleType `xml:"style,omitempty"`

	// generated from element "extent_curveExtent" of complex type ElevatedCurvePropertyType
	Extent_curveExtent *ElevatedCurvePropertyType `xml:"extent_curveExtent,omitempty"`

	// generated from element "extent_surfaceExtent" of complex type ElevatedSurfacePropertyType
	Extent_surfaceExtent *ElevatedSurfacePropertyType `xml:"extent_surfaceExtent,omitempty"`

	// generated from element "extent_location" of complex type ElevatedPointPropertyType
	Extent_location *ElevatedPointPropertyType `xml:"extent_location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MarkingElementPropertyType Named source named complex type "MarkingElementPropertyType"
type MarkingElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "MarkingElement" of complex type MarkingElementType
	MarkingElement *MarkingElementType `xml:"MarkingElement,omitempty"`
}

// RunwayMarkingType Named source named complex type "RunwayMarkingType"
type RunwayMarkingType struct {
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

	// generated from element "timeSlice" of complex type RunwayMarkingTimeSlicePropertyType
	TimeSlice []*RunwayMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayMarkingTimeSlicePropertyType Named source named complex type "RunwayMarkingTimeSlicePropertyType"
type RunwayMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayMarkingTimeSlice" of complex type RunwayMarkingTimeSliceType
	RunwayMarkingTimeSlice *RunwayMarkingTimeSliceType `xml:"RunwayMarkingTimeSlice,omitempty"`
}

// RunwayMarkingTimeSliceType Named source named complex type "RunwayMarkingTimeSliceType"
type RunwayMarkingTimeSliceType struct {
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

	// generated from group "aixm:RunwayMarkingPropertyGroup" with order 1661 depth 2
	Group_RunwayMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_39 `xml:"extension,omitempty"`
}

// A_extension_39 Named source within outer element "extension"
type A_extension_39 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayMarkingExtension" of complex type AbstractExtensionType
	AbstractRunwayMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_RunwayMarkingPropertyGroup UnNamed source named group "RunwayMarkingPropertyGroup"
type Group_RunwayMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1670 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markingLocation" of complex type CodeRunwaySectionType
	MarkingLocation *aixmdatatypes.CodeRunwaySectionType `xml:"markingLocation,omitempty"`

	// generated from element "markedRunway" of complex type RunwayPropertyType
	MarkedRunway *RunwayPropertyType `xml:"markedRunway,omitempty"`
}

// StandMarkingType Named source named complex type "StandMarkingType"
type StandMarkingType struct {
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

	// generated from element "timeSlice" of complex type StandMarkingTimeSlicePropertyType
	TimeSlice []*StandMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// StandMarkingTimeSlicePropertyType Named source named complex type "StandMarkingTimeSlicePropertyType"
type StandMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "StandMarkingTimeSlice" of complex type StandMarkingTimeSliceType
	StandMarkingTimeSlice *StandMarkingTimeSliceType `xml:"StandMarkingTimeSlice,omitempty"`
}

// StandMarkingTimeSliceType Named source named complex type "StandMarkingTimeSliceType"
type StandMarkingTimeSliceType struct {
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

	// generated from group "aixm:StandMarkingPropertyGroup" with order 1688 depth 2
	Group_StandMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_40 `xml:"extension,omitempty"`
}

// A_extension_40 Named source within outer element "extension"
type A_extension_40 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractStandMarkingExtension" of complex type AbstractExtensionType
	AbstractStandMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractStandMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_StandMarkingPropertyGroup UnNamed source named group "StandMarkingPropertyGroup"
type Group_StandMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1697 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markedStand" of complex type AircraftStandPropertyType
	MarkedStand *AircraftStandPropertyType `xml:"markedStand,omitempty"`
}

// TaxiHoldingPositionMarkingType Named source named complex type "TaxiHoldingPositionMarkingType"
type TaxiHoldingPositionMarkingType struct {
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

	// generated from element "timeSlice" of complex type TaxiHoldingPositionMarkingTimeSlicePropertyType
	TimeSlice []*TaxiHoldingPositionMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiHoldingPositionMarkingTimeSlicePropertyType Named source named complex type "TaxiHoldingPositionMarkingTimeSlicePropertyType"
type TaxiHoldingPositionMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiHoldingPositionMarkingTimeSlice" of complex type TaxiHoldingPositionMarkingTimeSliceType
	TaxiHoldingPositionMarkingTimeSlice *TaxiHoldingPositionMarkingTimeSliceType `xml:"TaxiHoldingPositionMarkingTimeSlice,omitempty"`
}

// TaxiHoldingPositionMarkingTimeSliceType Named source named complex type "TaxiHoldingPositionMarkingTimeSliceType"
type TaxiHoldingPositionMarkingTimeSliceType struct {
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

	// generated from group "aixm:TaxiHoldingPositionMarkingPropertyGroup" with order 1714 depth 2
	Group_TaxiHoldingPositionMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_41 `xml:"extension,omitempty"`
}

// A_extension_41 Named source within outer element "extension"
type A_extension_41 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiHoldingPositionMarkingExtension" of complex type AbstractExtensionType
	AbstractTaxiHoldingPositionMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiHoldingPositionMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_TaxiHoldingPositionMarkingPropertyGroup UnNamed source named group "TaxiHoldingPositionMarkingPropertyGroup"
type Group_TaxiHoldingPositionMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1723 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markedTaxiHold" of complex type TaxiHoldingPositionPropertyType
	MarkedTaxiHold *TaxiHoldingPositionPropertyType `xml:"markedTaxiHold,omitempty"`
}

// TaxiwayMarkingType Named source named complex type "TaxiwayMarkingType"
type TaxiwayMarkingType struct {
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

	// generated from element "timeSlice" of complex type TaxiwayMarkingTimeSlicePropertyType
	TimeSlice []*TaxiwayMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiwayMarkingTimeSlicePropertyType Named source named complex type "TaxiwayMarkingTimeSlicePropertyType"
type TaxiwayMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiwayMarkingTimeSlice" of complex type TaxiwayMarkingTimeSliceType
	TaxiwayMarkingTimeSlice *TaxiwayMarkingTimeSliceType `xml:"TaxiwayMarkingTimeSlice,omitempty"`
}

// TaxiwayMarkingTimeSliceType Named source named complex type "TaxiwayMarkingTimeSliceType"
type TaxiwayMarkingTimeSliceType struct {
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

	// generated from group "aixm:TaxiwayMarkingPropertyGroup" with order 1740 depth 2
	Group_TaxiwayMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_42 `xml:"extension,omitempty"`
}

// A_extension_42 Named source within outer element "extension"
type A_extension_42 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiwayMarkingExtension" of complex type AbstractExtensionType
	AbstractTaxiwayMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiwayMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_TaxiwayMarkingPropertyGroup UnNamed source named group "TaxiwayMarkingPropertyGroup"
type Group_TaxiwayMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1749 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markingLocation" of complex type CodeTaxiwaySectionType
	MarkingLocation *aixmdatatypes.CodeTaxiwaySectionType `xml:"markingLocation,omitempty"`

	// generated from element "markedTaxiway" of complex type TaxiwayPropertyType
	MarkedTaxiway *TaxiwayPropertyType `xml:"markedTaxiway,omitempty"`

	// generated from element "markedElement" of complex type TaxiwayElementPropertyType
	MarkedElement *TaxiwayElementPropertyType `xml:"markedElement,omitempty"`
}

// TouchDownLiftOffMarkingType Named source named complex type "TouchDownLiftOffMarkingType"
type TouchDownLiftOffMarkingType struct {
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

	// generated from element "timeSlice" of complex type TouchDownLiftOffMarkingTimeSlicePropertyType
	TimeSlice []*TouchDownLiftOffMarkingTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TouchDownLiftOffMarkingTimeSlicePropertyType Named source named complex type "TouchDownLiftOffMarkingTimeSlicePropertyType"
type TouchDownLiftOffMarkingTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TouchDownLiftOffMarkingTimeSlice" of complex type TouchDownLiftOffMarkingTimeSliceType
	TouchDownLiftOffMarkingTimeSlice *TouchDownLiftOffMarkingTimeSliceType `xml:"TouchDownLiftOffMarkingTimeSlice,omitempty"`
}

// TouchDownLiftOffMarkingTimeSliceType Named source named complex type "TouchDownLiftOffMarkingTimeSliceType"
type TouchDownLiftOffMarkingTimeSliceType struct {
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

	// generated from group "aixm:TouchDownLiftOffMarkingPropertyGroup" with order 1768 depth 2
	Group_TouchDownLiftOffMarkingPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_43 `xml:"extension,omitempty"`
}

// A_extension_43 Named source within outer element "extension"
type A_extension_43 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTouchDownLiftOffMarkingExtension" of complex type AbstractExtensionType
	AbstractTouchDownLiftOffMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTouchDownLiftOffMarkingExtension,omitempty"`

	// generated from element "AbstractMarkingExtension" of complex type AbstractExtensionType
	AbstractMarkingExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingExtension,omitempty"`
}

// Group_TouchDownLiftOffMarkingPropertyGroup UnNamed source named group "TouchDownLiftOffMarkingPropertyGroup"
type Group_TouchDownLiftOffMarkingPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:MarkingPropertyGroup" with order 1777 depth 2
	Group_MarkingPropertyGroup

	// generated from element "markingLocation" of complex type CodeTLOFSectionType
	MarkingLocation *aixmdatatypes.CodeTLOFSectionType `xml:"markingLocation,omitempty"`

	// generated from element "markedTouchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	MarkedTouchDownLiftOff *TouchDownLiftOffPropertyType `xml:"markedTouchDownLiftOff,omitempty"`
}

// ArrestingGearType Named source named complex type "ArrestingGearType"
type ArrestingGearType struct {
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

	// generated from element "timeSlice" of complex type ArrestingGearTimeSlicePropertyType
	TimeSlice []*ArrestingGearTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ArrestingGearTimeSlicePropertyType Named source named complex type "ArrestingGearTimeSlicePropertyType"
type ArrestingGearTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ArrestingGearTimeSlice" of complex type ArrestingGearTimeSliceType
	ArrestingGearTimeSlice *ArrestingGearTimeSliceType `xml:"ArrestingGearTimeSlice,omitempty"`
}

// ArrestingGearTimeSliceType Named source named complex type "ArrestingGearTimeSliceType"
type ArrestingGearTimeSliceType struct {
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

	// generated from group "aixm:ArrestingGearPropertyGroup" with order 1795 depth 2
	Group_ArrestingGearPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_44 `xml:"extension,omitempty"`
}

// A_extension_44 Named source within outer element "extension"
type A_extension_44 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractArrestingGearExtension" of complex type AbstractExtensionType
	AbstractArrestingGearExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractArrestingGearExtension,omitempty"`
}

// Group_ArrestingGearPropertyGroup UnNamed source named group "ArrestingGearPropertyGroup"
type Group_ArrestingGearPropertyGroup struct {

	// insertion point for fields

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "engageDevice" of complex type CodeArrestingGearEngageDeviceType
	EngageDevice *aixmdatatypes.CodeArrestingGearEngageDeviceType `xml:"engageDevice,omitempty"`

	// generated from element "absorbType" of complex type CodeArrestingGearEnergyAbsorbType
	AbsorbType *aixmdatatypes.CodeArrestingGearEnergyAbsorbType `xml:"absorbType,omitempty"`

	// generated from element "bidirectional" of complex type CodeYesNoType
	Bidirectional *aixmdatatypes.CodeYesNoType `xml:"bidirectional,omitempty"`

	// generated from element "location" of complex type ValDistanceType
	Location *aixmdatatypes.ValDistanceType `xml:"location,omitempty"`

	// generated from element "runwayDirection" of complex type RunwayDirectionPropertyType
	RunwayDirection []*RunwayDirectionPropertyType `xml:"runwayDirection,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "extent_surfaceExtent" of complex type ElevatedSurfacePropertyType
	Extent_surfaceExtent *ElevatedSurfacePropertyType `xml:"extent_surfaceExtent,omitempty"`

	// generated from element "extent_curveExtent" of complex type ElevatedCurvePropertyType
	Extent_curveExtent *ElevatedCurvePropertyType `xml:"extent_curveExtent,omitempty"`

	// generated from element "extent_pointExtent" of complex type ElevatedPointPropertyType
	Extent_pointExtent *ElevatedPointPropertyType `xml:"extent_pointExtent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ManoeuvringAreaAvailabilityType Named source named complex type "ManoeuvringAreaAvailabilityType"
type ManoeuvringAreaAvailabilityType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 1821 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:ManoeuvringAreaAvailabilityPropertyGroup" with order 1822 depth 2
	Group_ManoeuvringAreaAvailabilityPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_45 `xml:"extension,omitempty"`
}

// A_extension_45 Named source within outer element "extension"
type A_extension_45 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractManoeuvringAreaAvailabilityExtension" of complex type AbstractExtensionType
	AbstractManoeuvringAreaAvailabilityExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractManoeuvringAreaAvailabilityExtension,omitempty"`
}

// Group_ManoeuvringAreaAvailabilityPropertyGroup UnNamed source named group "ManoeuvringAreaAvailabilityPropertyGroup"
type Group_ManoeuvringAreaAvailabilityPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusAirportType
	OperationalStatus *aixmdatatypes.CodeStatusAirportType `xml:"operationalStatus,omitempty"`

	// generated from element "warning" of complex type CodeAirportWarningType
	Warning *aixmdatatypes.CodeAirportWarningType `xml:"warning,omitempty"`

	// generated from element "usage" of complex type ManoeuvringAreaUsagePropertyType
	Usage []*ManoeuvringAreaUsagePropertyType `xml:"usage,omitempty"`
}

// ManoeuvringAreaAvailabilityPropertyType Named source named complex type "ManoeuvringAreaAvailabilityPropertyType"
type ManoeuvringAreaAvailabilityPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ManoeuvringAreaAvailability" of complex type ManoeuvringAreaAvailabilityType
	ManoeuvringAreaAvailability *ManoeuvringAreaAvailabilityType `xml:"ManoeuvringAreaAvailability,omitempty"`
}

// ManoeuvringAreaUsageType Named source named complex type "ManoeuvringAreaUsageType"
type ManoeuvringAreaUsageType struct {
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

	// generated from group "aixm:UsageConditionPropertyGroup" with order 1841 depth 2
	Group_UsageConditionPropertyGroup

	// generated from group "aixm:ManoeuvringAreaUsagePropertyGroup" with order 1842 depth 2
	Group_ManoeuvringAreaUsagePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_46 `xml:"extension,omitempty"`
}

// A_extension_46 Named source within outer element "extension"
type A_extension_46 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractUsageConditionExtension" of complex type AbstractExtensionType
	AbstractUsageConditionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractUsageConditionExtension,omitempty"`

	// generated from element "AbstractManoeuvringAreaUsageExtension" of complex type AbstractExtensionType
	AbstractManoeuvringAreaUsageExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractManoeuvringAreaUsageExtension,omitempty"`
}

// Group_ManoeuvringAreaUsagePropertyGroup UnNamed source named group "ManoeuvringAreaUsagePropertyGroup"
type Group_ManoeuvringAreaUsagePropertyGroup struct {

	// insertion point for fields

	// generated from element "operation" of complex type CodeOperationManoeuvringAreaType
	Operation *aixmdatatypes.CodeOperationManoeuvringAreaType `xml:"operation,omitempty"`
}

// ManoeuvringAreaUsagePropertyType Named source named complex type "ManoeuvringAreaUsagePropertyType"
type ManoeuvringAreaUsagePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ManoeuvringAreaUsage" of complex type ManoeuvringAreaUsageType
	ManoeuvringAreaUsage *ManoeuvringAreaUsageType `xml:"ManoeuvringAreaUsage,omitempty"`
}

// NavaidEquipmentDistanceType Named source named complex type "NavaidEquipmentDistanceType"
type NavaidEquipmentDistanceType struct {
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

	// generated from group "aixm:NavaidEquipmentDistancePropertyGroup" with order 1859 depth 2
	Group_NavaidEquipmentDistancePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_47 `xml:"extension,omitempty"`
}

// A_extension_47 Named source within outer element "extension"
type A_extension_47 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavaidEquipmentDistanceExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentDistanceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentDistanceExtension,omitempty"`
}

// Group_NavaidEquipmentDistancePropertyGroup UnNamed source named group "NavaidEquipmentDistancePropertyGroup"
type Group_NavaidEquipmentDistancePropertyGroup struct {

	// insertion point for fields

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "distanceAccuracy" of complex type ValDistanceType
	DistanceAccuracy *aixmdatatypes.ValDistanceType `xml:"distanceAccuracy,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theNavaidEquipment" of complex type NavaidEquipmentPropertyType
	TheNavaidEquipment *NavaidEquipmentPropertyType `xml:"theNavaidEquipment,omitempty"`
}

// NavaidEquipmentDistancePropertyType Named source named complex type "NavaidEquipmentDistancePropertyType"
type NavaidEquipmentDistancePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "NavaidEquipmentDistance" of complex type NavaidEquipmentDistanceType
	NavaidEquipmentDistance *NavaidEquipmentDistanceType `xml:"NavaidEquipmentDistance,omitempty"`
}

// RunwayPropertyType Named source named complex type "RunwayPropertyType"
type RunwayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RunwayType Named source named complex type "RunwayType"
type RunwayType struct {
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

	// generated from element "timeSlice" of complex type RunwayTimeSlicePropertyType
	TimeSlice []*RunwayTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayTimeSlicePropertyType Named source named complex type "RunwayTimeSlicePropertyType"
type RunwayTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayTimeSlice" of complex type RunwayTimeSliceType
	RunwayTimeSlice *RunwayTimeSliceType `xml:"RunwayTimeSlice,omitempty"`
}

// RunwayTimeSliceType Named source named complex type "RunwayTimeSliceType"
type RunwayTimeSliceType struct {
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

	// generated from group "aixm:RunwayPropertyGroup" with order 1889 depth 2
	Group_RunwayPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_48 `xml:"extension,omitempty"`
}

// A_extension_48 Named source within outer element "extension"
type A_extension_48 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayExtension" of complex type AbstractExtensionType
	AbstractRunwayExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayExtension,omitempty"`
}

// Group_RunwayPropertyGroup UnNamed source named group "RunwayPropertyGroup"
type Group_RunwayPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeRunwayType
	Type_XSD *aixmdatatypes.CodeRunwayType `xml:"type,omitempty"`

	// generated from element "nominalLength" of complex type ValDistanceType
	NominalLength *aixmdatatypes.ValDistanceType `xml:"nominalLength,omitempty"`

	// generated from element "lengthAccuracy" of complex type ValDistanceType
	LengthAccuracy *aixmdatatypes.ValDistanceType `xml:"lengthAccuracy,omitempty"`

	// generated from element "nominalWidth" of complex type ValDistanceType
	NominalWidth *aixmdatatypes.ValDistanceType `xml:"nominalWidth,omitempty"`

	// generated from element "widthAccuracy" of complex type ValDistanceType
	WidthAccuracy *aixmdatatypes.ValDistanceType `xml:"widthAccuracy,omitempty"`

	// generated from element "widthShoulder" of complex type ValDistanceType
	WidthShoulder *aixmdatatypes.ValDistanceType `xml:"widthShoulder,omitempty"`

	// generated from element "lengthStrip" of complex type ValDistanceType
	LengthStrip *aixmdatatypes.ValDistanceType `xml:"lengthStrip,omitempty"`

	// generated from element "widthStrip" of complex type ValDistanceType
	WidthStrip *aixmdatatypes.ValDistanceType `xml:"widthStrip,omitempty"`

	// generated from element "lengthOffset" of complex type ValDistanceSignedType
	LengthOffset *aixmdatatypes.ValDistanceSignedType `xml:"lengthOffset,omitempty"`

	// generated from element "widthOffset" of complex type ValDistanceSignedType
	WidthOffset *aixmdatatypes.ValDistanceSignedType `xml:"widthOffset,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "overallContaminant" of complex type RunwayContaminationPropertyType
	OverallContaminant []*RunwayContaminationPropertyType `xml:"overallContaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "areaContaminant" of complex type RunwaySectionContaminationPropertyType
	AreaContaminant []*RunwaySectionContaminationPropertyType `xml:"areaContaminant,omitempty"`
}

// RunwayBlastPadType Named source named complex type "RunwayBlastPadType"
type RunwayBlastPadType struct {
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

	// generated from element "timeSlice" of complex type RunwayBlastPadTimeSlicePropertyType
	TimeSlice []*RunwayBlastPadTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayBlastPadTimeSlicePropertyType Named source named complex type "RunwayBlastPadTimeSlicePropertyType"
type RunwayBlastPadTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayBlastPadTimeSlice" of complex type RunwayBlastPadTimeSliceType
	RunwayBlastPadTimeSlice *RunwayBlastPadTimeSliceType `xml:"RunwayBlastPadTimeSlice,omitempty"`
}

// RunwayBlastPadTimeSliceType Named source named complex type "RunwayBlastPadTimeSliceType"
type RunwayBlastPadTimeSliceType struct {
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

	// generated from group "aixm:RunwayBlastPadPropertyGroup" with order 1929 depth 2
	Group_RunwayBlastPadPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_49 `xml:"extension,omitempty"`
}

// A_extension_49 Named source within outer element "extension"
type A_extension_49 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayBlastPadExtension" of complex type AbstractExtensionType
	AbstractRunwayBlastPadExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayBlastPadExtension,omitempty"`
}

// Group_RunwayBlastPadPropertyGroup UnNamed source named group "RunwayBlastPadPropertyGroup"
type Group_RunwayBlastPadPropertyGroup struct {

	// insertion point for fields

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`

	// generated from element "usedRunwayDirection" of complex type RunwayDirectionPropertyType
	UsedRunwayDirection *RunwayDirectionPropertyType `xml:"usedRunwayDirection,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
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

// RunwayCentrelinePointType Named source named complex type "RunwayCentrelinePointType"
type RunwayCentrelinePointType struct {
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

	// generated from element "timeSlice" of complex type RunwayCentrelinePointTimeSlicePropertyType
	TimeSlice []*RunwayCentrelinePointTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayCentrelinePointTimeSlicePropertyType Named source named complex type "RunwayCentrelinePointTimeSlicePropertyType"
type RunwayCentrelinePointTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayCentrelinePointTimeSlice" of complex type RunwayCentrelinePointTimeSliceType
	RunwayCentrelinePointTimeSlice *RunwayCentrelinePointTimeSliceType `xml:"RunwayCentrelinePointTimeSlice,omitempty"`
}

// RunwayCentrelinePointTimeSliceType Named source named complex type "RunwayCentrelinePointTimeSliceType"
type RunwayCentrelinePointTimeSliceType struct {
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

	// generated from group "aixm:RunwayCentrelinePointPropertyGroup" with order 1958 depth 2
	Group_RunwayCentrelinePointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_50 `xml:"extension,omitempty"`
}

// A_extension_50 Named source within outer element "extension"
type A_extension_50 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayCentrelinePointExtension" of complex type AbstractExtensionType
	AbstractRunwayCentrelinePointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayCentrelinePointExtension,omitempty"`
}

// Group_RunwayCentrelinePointPropertyGroup UnNamed source named group "RunwayCentrelinePointPropertyGroup"
type Group_RunwayCentrelinePointPropertyGroup struct {

	// insertion point for fields

	// generated from element "role" of complex type CodeRunwayPointRoleType
	Role *aixmdatatypes.CodeRunwayPointRoleType `xml:"role,omitempty"`

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "onRunway" of complex type RunwayDirectionPropertyType
	OnRunway *RunwayDirectionPropertyType `xml:"onRunway,omitempty"`

	// generated from element "associatedDeclaredDistance" of complex type RunwayDeclaredDistancePropertyType
	AssociatedDeclaredDistance []*RunwayDeclaredDistancePropertyType `xml:"associatedDeclaredDistance,omitempty"`

	// generated from element "navaidEquipment" of complex type NavaidEquipmentDistancePropertyType
	NavaidEquipment []*NavaidEquipmentDistancePropertyType `xml:"navaidEquipment,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RunwayDeclaredDistanceType Named source named complex type "RunwayDeclaredDistanceType"
type RunwayDeclaredDistanceType struct {
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

	// generated from group "aixm:RunwayDeclaredDistancePropertyGroup" with order 1977 depth 2
	Group_RunwayDeclaredDistancePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_51 `xml:"extension,omitempty"`
}

// A_extension_51 Named source within outer element "extension"
type A_extension_51 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayDeclaredDistanceExtension" of complex type AbstractExtensionType
	AbstractRunwayDeclaredDistanceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayDeclaredDistanceExtension,omitempty"`
}

// Group_RunwayDeclaredDistancePropertyGroup UnNamed source named group "RunwayDeclaredDistancePropertyGroup"
type Group_RunwayDeclaredDistancePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeDeclaredDistanceType
	Type_XSD *aixmdatatypes.CodeDeclaredDistanceType `xml:"type,omitempty"`

	// generated from element "declaredValue" of complex type RunwayDeclaredDistanceValuePropertyType
	DeclaredValue []*RunwayDeclaredDistanceValuePropertyType `xml:"declaredValue,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RunwayDeclaredDistancePropertyType Named source named complex type "RunwayDeclaredDistancePropertyType"
type RunwayDeclaredDistancePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RunwayDeclaredDistance" of complex type RunwayDeclaredDistanceType
	RunwayDeclaredDistance *RunwayDeclaredDistanceType `xml:"RunwayDeclaredDistance,omitempty"`
}

// RunwayDeclaredDistanceValueType Named source named complex type "RunwayDeclaredDistanceValueType"
type RunwayDeclaredDistanceValueType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 1995 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:RunwayDeclaredDistanceValuePropertyGroup" with order 1996 depth 2
	Group_RunwayDeclaredDistanceValuePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_52 `xml:"extension,omitempty"`
}

// A_extension_52 Named source within outer element "extension"
type A_extension_52 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractRunwayDeclaredDistanceValueExtension" of complex type AbstractExtensionType
	AbstractRunwayDeclaredDistanceValueExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayDeclaredDistanceValueExtension,omitempty"`
}

// Group_RunwayDeclaredDistanceValuePropertyGroup UnNamed source named group "RunwayDeclaredDistanceValuePropertyGroup"
type Group_RunwayDeclaredDistanceValuePropertyGroup struct {

	// insertion point for fields

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "distanceAccuracy" of complex type ValDistanceType
	DistanceAccuracy *aixmdatatypes.ValDistanceType `xml:"distanceAccuracy,omitempty"`
}

// RunwayDeclaredDistanceValuePropertyType Named source named complex type "RunwayDeclaredDistanceValuePropertyType"
type RunwayDeclaredDistanceValuePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RunwayDeclaredDistanceValue" of complex type RunwayDeclaredDistanceValueType
	RunwayDeclaredDistanceValue *RunwayDeclaredDistanceValueType `xml:"RunwayDeclaredDistanceValue,omitempty"`
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

// RunwayDirectionType Named source named complex type "RunwayDirectionType"
type RunwayDirectionType struct {
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

	// generated from element "timeSlice" of complex type RunwayDirectionTimeSlicePropertyType
	TimeSlice []*RunwayDirectionTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayDirectionTimeSlicePropertyType Named source named complex type "RunwayDirectionTimeSlicePropertyType"
type RunwayDirectionTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayDirectionTimeSlice" of complex type RunwayDirectionTimeSliceType
	RunwayDirectionTimeSlice *RunwayDirectionTimeSliceType `xml:"RunwayDirectionTimeSlice,omitempty"`
}

// RunwayDirectionTimeSliceType Named source named complex type "RunwayDirectionTimeSliceType"
type RunwayDirectionTimeSliceType struct {
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

	// generated from group "aixm:RunwayDirectionPropertyGroup" with order 2025 depth 2
	Group_RunwayDirectionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_53 `xml:"extension,omitempty"`
}

// A_extension_53 Named source within outer element "extension"
type A_extension_53 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayDirectionExtension" of complex type AbstractExtensionType
	AbstractRunwayDirectionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayDirectionExtension,omitempty"`
}

// Group_RunwayDirectionPropertyGroup UnNamed source named group "RunwayDirectionPropertyGroup"
type Group_RunwayDirectionPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from element "trueBearingAccuracy" of complex type ValAngleType
	TrueBearingAccuracy *aixmdatatypes.ValAngleType `xml:"trueBearingAccuracy,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "patternVFR" of complex type CodeDirectionTurnType
	PatternVFR *aixmdatatypes.CodeDirectionTurnType `xml:"patternVFR,omitempty"`

	// generated from element "slopeTDZ" of complex type ValSlopeType
	SlopeTDZ *aixmdatatypes.ValSlopeType `xml:"slopeTDZ,omitempty"`

	// generated from element "elevationTDZ" of complex type ValDistanceVerticalType
	ElevationTDZ *aixmdatatypes.ValDistanceVerticalType `xml:"elevationTDZ,omitempty"`

	// generated from element "elevationTDZAccuracy" of complex type ValDistanceType
	ElevationTDZAccuracy *aixmdatatypes.ValDistanceType `xml:"elevationTDZAccuracy,omitempty"`

	// generated from element "approachMarkingType" of complex type CodeRunwayMarkingType
	ApproachMarkingType *aixmdatatypes.CodeRunwayMarkingType `xml:"approachMarkingType,omitempty"`

	// generated from element "approachMarkingCondition" of complex type CodeMarkingConditionType
	ApproachMarkingCondition *aixmdatatypes.CodeMarkingConditionType `xml:"approachMarkingCondition,omitempty"`

	// generated from element "classLightingJAR" of complex type CodeLightingJARType
	ClassLightingJAR *aixmdatatypes.CodeLightingJARType `xml:"classLightingJAR,omitempty"`

	// generated from element "precisionApproachGuidance" of complex type CodeApproachGuidanceType
	PrecisionApproachGuidance *aixmdatatypes.CodeApproachGuidanceType `xml:"precisionApproachGuidance,omitempty"`

	// generated from element "usedRunway" of complex type RunwayPropertyType
	UsedRunway *RunwayPropertyType `xml:"usedRunway,omitempty"`

	// generated from element "startingElement" of complex type RunwayElementPropertyType
	StartingElement *RunwayElementPropertyType `xml:"startingElement,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// RunwayElementPropertyType Named source named complex type "RunwayElementPropertyType"
type RunwayElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RunwayElementType Named source named complex type "RunwayElementType"
type RunwayElementType struct {
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

	// generated from element "timeSlice" of complex type RunwayElementTimeSlicePropertyType
	TimeSlice []*RunwayElementTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayElementTimeSlicePropertyType Named source named complex type "RunwayElementTimeSlicePropertyType"
type RunwayElementTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayElementTimeSlice" of complex type RunwayElementTimeSliceType
	RunwayElementTimeSlice *RunwayElementTimeSliceType `xml:"RunwayElementTimeSlice,omitempty"`
}

// RunwayElementTimeSliceType Named source named complex type "RunwayElementTimeSliceType"
type RunwayElementTimeSliceType struct {
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

	// generated from group "aixm:RunwayElementPropertyGroup" with order 2064 depth 2
	Group_RunwayElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_54 `xml:"extension,omitempty"`
}

// A_extension_54 Named source within outer element "extension"
type A_extension_54 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayElementExtension" of complex type AbstractExtensionType
	AbstractRunwayElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayElementExtension,omitempty"`
}

// Group_RunwayElementPropertyGroup UnNamed source named group "RunwayElementPropertyGroup"
type Group_RunwayElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeRunwayElementType
	Type_XSD *aixmdatatypes.CodeRunwayElementType `xml:"type,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "gradeSeparation" of complex type CodeGradeSeparationType
	GradeSeparation *aixmdatatypes.CodeGradeSeparationType `xml:"gradeSeparation,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedRunway" of complex type RunwayPropertyType
	AssociatedRunway []*RunwayPropertyType `xml:"associatedRunway,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// RunwayProtectAreaPropertyType Named source named complex type "RunwayProtectAreaPropertyType"
type RunwayProtectAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// RunwayProtectAreaType Named source named complex type "RunwayProtectAreaType"
type RunwayProtectAreaType struct {
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

	// generated from element "timeSlice" of complex type RunwayProtectAreaTimeSlicePropertyType
	TimeSlice []*RunwayProtectAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayProtectAreaTimeSlicePropertyType Named source named complex type "RunwayProtectAreaTimeSlicePropertyType"
type RunwayProtectAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayProtectAreaTimeSlice" of complex type RunwayProtectAreaTimeSliceType
	RunwayProtectAreaTimeSlice *RunwayProtectAreaTimeSliceType `xml:"RunwayProtectAreaTimeSlice,omitempty"`
}

// RunwayProtectAreaTimeSliceType Named source named complex type "RunwayProtectAreaTimeSliceType"
type RunwayProtectAreaTimeSliceType struct {
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

	// generated from group "aixm:RunwayProtectAreaPropertyGroup" with order 2096 depth 2
	Group_RunwayProtectAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_55 `xml:"extension,omitempty"`
}

// A_extension_55 Named source within outer element "extension"
type A_extension_55 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayProtectAreaExtension" of complex type AbstractExtensionType
	AbstractRunwayProtectAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayProtectAreaExtension,omitempty"`

	// generated from element "AbstractAirportHeliportProtectionAreaExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportProtectionAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportProtectionAreaExtension,omitempty"`
}

// Group_RunwayProtectAreaPropertyGroup UnNamed source named group "RunwayProtectAreaPropertyGroup"
type Group_RunwayProtectAreaPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:AirportHeliportProtectionAreaPropertyGroup" with order 2105 depth 2
	Group_AirportHeliportProtectionAreaPropertyGroup

	// generated from element "type" of complex type CodeRunwayProtectionAreaType
	Type_XSD *aixmdatatypes.CodeRunwayProtectionAreaType `xml:"type,omitempty"`

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`

	// generated from element "protectedRunwayDirection" of complex type RunwayDirectionPropertyType
	ProtectedRunwayDirection *RunwayDirectionPropertyType `xml:"protectedRunwayDirection,omitempty"`
}

// RunwayVisualRangeType Named source named complex type "RunwayVisualRangeType"
type RunwayVisualRangeType struct {
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

	// generated from element "timeSlice" of complex type RunwayVisualRangeTimeSlicePropertyType
	TimeSlice []*RunwayVisualRangeTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RunwayVisualRangeTimeSlicePropertyType Named source named complex type "RunwayVisualRangeTimeSlicePropertyType"
type RunwayVisualRangeTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RunwayVisualRangeTimeSlice" of complex type RunwayVisualRangeTimeSliceType
	RunwayVisualRangeTimeSlice *RunwayVisualRangeTimeSliceType `xml:"RunwayVisualRangeTimeSlice,omitempty"`
}

// RunwayVisualRangeTimeSliceType Named source named complex type "RunwayVisualRangeTimeSliceType"
type RunwayVisualRangeTimeSliceType struct {
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

	// generated from group "aixm:RunwayVisualRangePropertyGroup" with order 2124 depth 2
	Group_RunwayVisualRangePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_56 `xml:"extension,omitempty"`
}

// A_extension_56 Named source within outer element "extension"
type A_extension_56 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRunwayVisualRangeExtension" of complex type AbstractExtensionType
	AbstractRunwayVisualRangeExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayVisualRangeExtension,omitempty"`
}

// Group_RunwayVisualRangePropertyGroup UnNamed source named group "RunwayVisualRangePropertyGroup"
type Group_RunwayVisualRangePropertyGroup struct {

	// insertion point for fields

	// generated from element "readingPosition" of complex type CodeRVRReadingType
	ReadingPosition *aixmdatatypes.CodeRVRReadingType `xml:"readingPosition,omitempty"`

	// generated from element "associatedRunwayDirection" of complex type RunwayDirectionPropertyType
	AssociatedRunwayDirection []*RunwayDirectionPropertyType `xml:"associatedRunwayDirection,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// VisualGlideSlopeIndicatorType Named source named complex type "VisualGlideSlopeIndicatorType"
type VisualGlideSlopeIndicatorType struct {
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

	// generated from element "timeSlice" of complex type VisualGlideSlopeIndicatorTimeSlicePropertyType
	TimeSlice []*VisualGlideSlopeIndicatorTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// VisualGlideSlopeIndicatorTimeSlicePropertyType Named source named complex type "VisualGlideSlopeIndicatorTimeSlicePropertyType"
type VisualGlideSlopeIndicatorTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "VisualGlideSlopeIndicatorTimeSlice" of complex type VisualGlideSlopeIndicatorTimeSliceType
	VisualGlideSlopeIndicatorTimeSlice *VisualGlideSlopeIndicatorTimeSliceType `xml:"VisualGlideSlopeIndicatorTimeSlice,omitempty"`
}

// VisualGlideSlopeIndicatorTimeSliceType Named source named complex type "VisualGlideSlopeIndicatorTimeSliceType"
type VisualGlideSlopeIndicatorTimeSliceType struct {
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

	// generated from group "aixm:VisualGlideSlopeIndicatorPropertyGroup" with order 2151 depth 2
	Group_VisualGlideSlopeIndicatorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_57 `xml:"extension,omitempty"`
}

// A_extension_57 Named source within outer element "extension"
type A_extension_57 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractVisualGlideSlopeIndicatorExtension" of complex type AbstractExtensionType
	AbstractVisualGlideSlopeIndicatorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractVisualGlideSlopeIndicatorExtension,omitempty"`

	// generated from element "AbstractGroundLightSystemExtension" of complex type AbstractExtensionType
	AbstractGroundLightSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGroundLightSystemExtension,omitempty"`
}

// Group_VisualGlideSlopeIndicatorPropertyGroup UnNamed source named group "VisualGlideSlopeIndicatorPropertyGroup"
type Group_VisualGlideSlopeIndicatorPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:GroundLightSystemPropertyGroup" with order 2160 depth 2
	Group_GroundLightSystemPropertyGroup

	// generated from element "type" of complex type CodeVASISType
	Type_XSD *aixmdatatypes.CodeVASISType `xml:"type,omitempty"`

	// generated from element "position" of complex type CodeSideType
	Position *aixmdatatypes.CodeSideType `xml:"position,omitempty"`

	// generated from element "numberBox" of complex type NoNumberType
	NumberBox *aixmdatatypes.NoNumberType `xml:"numberBox,omitempty"`

	// generated from element "portable" of complex type CodeYesNoType
	Portable *aixmdatatypes.CodeYesNoType `xml:"portable,omitempty"`

	// generated from element "slopeAngle" of complex type ValAngleType
	SlopeAngle *aixmdatatypes.ValAngleType `xml:"slopeAngle,omitempty"`

	// generated from element "minimumEyeHeightOverThreshold" of complex type ValDistanceVerticalType
	MinimumEyeHeightOverThreshold *aixmdatatypes.ValDistanceVerticalType `xml:"minimumEyeHeightOverThreshold,omitempty"`

	// generated from element "runwayDirection" of complex type RunwayDirectionPropertyType
	RunwayDirection *RunwayDirectionPropertyType `xml:"runwayDirection,omitempty"`
}

// FloatingDockSitePropertyType Named source named complex type "FloatingDockSitePropertyType"
type FloatingDockSitePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// FloatingDockSiteType Named source named complex type "FloatingDockSiteType"
type FloatingDockSiteType struct {
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

	// generated from element "timeSlice" of complex type FloatingDockSiteTimeSlicePropertyType
	TimeSlice []*FloatingDockSiteTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// FloatingDockSiteTimeSlicePropertyType Named source named complex type "FloatingDockSiteTimeSlicePropertyType"
type FloatingDockSiteTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "FloatingDockSiteTimeSlice" of complex type FloatingDockSiteTimeSliceType
	FloatingDockSiteTimeSlice *FloatingDockSiteTimeSliceType `xml:"FloatingDockSiteTimeSlice,omitempty"`
}

// FloatingDockSiteTimeSliceType Named source named complex type "FloatingDockSiteTimeSliceType"
type FloatingDockSiteTimeSliceType struct {
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

	// generated from group "aixm:FloatingDockSitePropertyGroup" with order 2183 depth 2
	Group_FloatingDockSitePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_58 `xml:"extension,omitempty"`
}

// A_extension_58 Named source within outer element "extension"
type A_extension_58 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractFloatingDockSiteExtension" of complex type AbstractExtensionType
	AbstractFloatingDockSiteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractFloatingDockSiteExtension,omitempty"`
}

// Group_FloatingDockSitePropertyGroup UnNamed source named group "FloatingDockSitePropertyGroup"
type Group_FloatingDockSitePropertyGroup struct {

	// insertion point for fields

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// MarkingBuoyType Named source named complex type "MarkingBuoyType"
type MarkingBuoyType struct {
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

	// generated from element "timeSlice" of complex type MarkingBuoyTimeSlicePropertyType
	TimeSlice []*MarkingBuoyTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// MarkingBuoyTimeSlicePropertyType Named source named complex type "MarkingBuoyTimeSlicePropertyType"
type MarkingBuoyTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MarkingBuoyTimeSlice" of complex type MarkingBuoyTimeSliceType
	MarkingBuoyTimeSlice *MarkingBuoyTimeSliceType `xml:"MarkingBuoyTimeSlice,omitempty"`
}

// MarkingBuoyTimeSliceType Named source named complex type "MarkingBuoyTimeSliceType"
type MarkingBuoyTimeSliceType struct {
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

	// generated from group "aixm:MarkingBuoyPropertyGroup" with order 2208 depth 2
	Group_MarkingBuoyPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_59 `xml:"extension,omitempty"`
}

// A_extension_59 Named source within outer element "extension"
type A_extension_59 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMarkingBuoyExtension" of complex type AbstractExtensionType
	AbstractMarkingBuoyExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkingBuoyExtension,omitempty"`
}

// Group_MarkingBuoyPropertyGroup UnNamed source named group "MarkingBuoyPropertyGroup"
type Group_MarkingBuoyPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type CodeBuoyDesignatorType
	Designator *aixmdatatypes.CodeBuoyDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeBuoyType
	Type_XSD *aixmdatatypes.CodeBuoyType `xml:"type,omitempty"`

	// generated from element "colour" of complex type CodeColourType
	Colour *aixmdatatypes.CodeColourType `xml:"colour,omitempty"`

	// generated from element "theSeaplaneLandingArea" of complex type SeaplaneLandingAreaPropertyType
	TheSeaplaneLandingArea *SeaplaneLandingAreaPropertyType `xml:"theSeaplaneLandingArea,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SeaplaneLandingAreaPropertyType Named source named complex type "SeaplaneLandingAreaPropertyType"
type SeaplaneLandingAreaPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// SeaplaneLandingAreaType Named source named complex type "SeaplaneLandingAreaType"
type SeaplaneLandingAreaType struct {
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

	// generated from element "timeSlice" of complex type SeaplaneLandingAreaTimeSlicePropertyType
	TimeSlice []*SeaplaneLandingAreaTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SeaplaneLandingAreaTimeSlicePropertyType Named source named complex type "SeaplaneLandingAreaTimeSlicePropertyType"
type SeaplaneLandingAreaTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SeaplaneLandingAreaTimeSlice" of complex type SeaplaneLandingAreaTimeSliceType
	SeaplaneLandingAreaTimeSlice *SeaplaneLandingAreaTimeSliceType `xml:"SeaplaneLandingAreaTimeSlice,omitempty"`
}

// SeaplaneLandingAreaTimeSliceType Named source named complex type "SeaplaneLandingAreaTimeSliceType"
type SeaplaneLandingAreaTimeSliceType struct {
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

	// generated from group "aixm:SeaplaneLandingAreaPropertyGroup" with order 2237 depth 2
	Group_SeaplaneLandingAreaPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_60 `xml:"extension,omitempty"`
}

// A_extension_60 Named source within outer element "extension"
type A_extension_60 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSeaplaneLandingAreaExtension" of complex type AbstractExtensionType
	AbstractSeaplaneLandingAreaExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSeaplaneLandingAreaExtension,omitempty"`
}

// Group_SeaplaneLandingAreaPropertyGroup UnNamed source named group "SeaplaneLandingAreaPropertyGroup"
type Group_SeaplaneLandingAreaPropertyGroup struct {

	// insertion point for fields

	// generated from element "rampSite" of complex type SeaplaneRampSitePropertyType
	RampSite []*SeaplaneRampSitePropertyType `xml:"rampSite,omitempty"`

	// generated from element "dockSite" of complex type FloatingDockSitePropertyType
	DockSite []*FloatingDockSitePropertyType `xml:"dockSite,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// SeaplaneRampSitePropertyType Named source named complex type "SeaplaneRampSitePropertyType"
type SeaplaneRampSitePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// SeaplaneRampSiteType Named source named complex type "SeaplaneRampSiteType"
type SeaplaneRampSiteType struct {
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

	// generated from element "timeSlice" of complex type SeaplaneRampSiteTimeSlicePropertyType
	TimeSlice []*SeaplaneRampSiteTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SeaplaneRampSiteTimeSlicePropertyType Named source named complex type "SeaplaneRampSiteTimeSlicePropertyType"
type SeaplaneRampSiteTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SeaplaneRampSiteTimeSlice" of complex type SeaplaneRampSiteTimeSliceType
	SeaplaneRampSiteTimeSlice *SeaplaneRampSiteTimeSliceType `xml:"SeaplaneRampSiteTimeSlice,omitempty"`
}

// SeaplaneRampSiteTimeSliceType Named source named complex type "SeaplaneRampSiteTimeSliceType"
type SeaplaneRampSiteTimeSliceType struct {
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

	// generated from group "aixm:SeaplaneRampSitePropertyGroup" with order 2265 depth 2
	Group_SeaplaneRampSitePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_61 `xml:"extension,omitempty"`
}

// A_extension_61 Named source within outer element "extension"
type A_extension_61 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSeaplaneRampSiteExtension" of complex type AbstractExtensionType
	AbstractSeaplaneRampSiteExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSeaplaneRampSiteExtension,omitempty"`
}

// Group_SeaplaneRampSitePropertyGroup UnNamed source named group "SeaplaneRampSitePropertyGroup"
type Group_SeaplaneRampSitePropertyGroup struct {

	// insertion point for fields

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "centreline" of complex type ElevatedCurvePropertyType
	Centreline *ElevatedCurvePropertyType `xml:"centreline,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// AircraftStandContaminationType Named source named complex type "AircraftStandContaminationType"
type AircraftStandContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2280 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_62 `xml:"extension,omitempty"`
}

// A_extension_62 Named source within outer element "extension"
type A_extension_62 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractAircraftStandContaminationExtension" of complex type AbstractExtensionType
	AbstractAircraftStandContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAircraftStandContaminationExtension,omitempty"`
}

// Group_AircraftStandContaminationPropertyGroup UnNamed source named group "AircraftStandContaminationPropertyGroup"
type Group_AircraftStandContaminationPropertyGroup struct {

	// insertion point for fields
}

// AircraftStandContaminationPropertyType Named source named complex type "AircraftStandContaminationPropertyType"
type AircraftStandContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AircraftStandContamination" of complex type AircraftStandContaminationType
	AircraftStandContamination *AircraftStandContaminationType `xml:"AircraftStandContamination,omitempty"`
}

// AirportHeliportContaminationType Named source named complex type "AirportHeliportContaminationType"
type AirportHeliportContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2297 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_63 `xml:"extension,omitempty"`
}

// A_extension_63 Named source within outer element "extension"
type A_extension_63 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractAirportHeliportContaminationExtension" of complex type AbstractExtensionType
	AbstractAirportHeliportContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAirportHeliportContaminationExtension,omitempty"`
}

// Group_AirportHeliportContaminationPropertyGroup UnNamed source named group "AirportHeliportContaminationPropertyGroup"
type Group_AirportHeliportContaminationPropertyGroup struct {

	// insertion point for fields
}

// AirportHeliportContaminationPropertyType Named source named complex type "AirportHeliportContaminationPropertyType"
type AirportHeliportContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AirportHeliportContamination" of complex type AirportHeliportContaminationType
	AirportHeliportContamination *AirportHeliportContaminationType `xml:"AirportHeliportContamination,omitempty"`
}

// ApronContaminationType Named source named complex type "ApronContaminationType"
type ApronContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2314 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_64 `xml:"extension,omitempty"`
}

// A_extension_64 Named source within outer element "extension"
type A_extension_64 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractApronContaminationExtension" of complex type AbstractExtensionType
	AbstractApronContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractApronContaminationExtension,omitempty"`
}

// Group_ApronContaminationPropertyGroup UnNamed source named group "ApronContaminationPropertyGroup"
type Group_ApronContaminationPropertyGroup struct {

	// insertion point for fields
}

// ApronContaminationPropertyType Named source named complex type "ApronContaminationPropertyType"
type ApronContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "ApronContamination" of complex type ApronContaminationType
	ApronContamination *ApronContaminationType `xml:"ApronContamination,omitempty"`
}

// RidgeType Named source named complex type "RidgeType"
type RidgeType struct {
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

	// generated from group "aixm:RidgePropertyGroup" with order 2331 depth 2
	Group_RidgePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_65 `xml:"extension,omitempty"`
}

// A_extension_65 Named source within outer element "extension"
type A_extension_65 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRidgeExtension" of complex type AbstractExtensionType
	AbstractRidgeExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRidgeExtension,omitempty"`
}

// Group_RidgePropertyGroup UnNamed source named group "RidgePropertyGroup"
type Group_RidgePropertyGroup struct {

	// insertion point for fields

	// generated from element "side" of complex type CodeSideType
	Side *aixmdatatypes.CodeSideType `xml:"side,omitempty"`

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "depth" of complex type ValDepthType
	Depth *aixmdatatypes.ValDepthType `xml:"depth,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RidgePropertyType Named source named complex type "RidgePropertyType"
type RidgePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Ridge" of complex type RidgeType
	Ridge *RidgeType `xml:"Ridge,omitempty"`
}

// RunwayContaminationType Named source named complex type "RunwayContaminationType"
type RunwayContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2350 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from group "aixm:RunwayContaminationPropertyGroup" with order 2351 depth 2
	Group_RunwayContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_66 `xml:"extension,omitempty"`
}

// A_extension_66 Named source within outer element "extension"
type A_extension_66 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractRunwayContaminationExtension" of complex type AbstractExtensionType
	AbstractRunwayContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwayContaminationExtension,omitempty"`
}

// Group_RunwayContaminationPropertyGroup UnNamed source named group "RunwayContaminationPropertyGroup"
type Group_RunwayContaminationPropertyGroup struct {

	// insertion point for fields

	// generated from element "clearedLength" of complex type ValDistanceType
	ClearedLength *aixmdatatypes.ValDistanceType `xml:"clearedLength,omitempty"`

	// generated from element "clearedWidth" of complex type ValDistanceType
	ClearedWidth *aixmdatatypes.ValDistanceType `xml:"clearedWidth,omitempty"`

	// generated from element "clearedSide" of complex type CodeSideType
	ClearedSide *aixmdatatypes.CodeSideType `xml:"clearedSide,omitempty"`

	// generated from element "furtherClearanceLength" of complex type ValDistanceType
	FurtherClearanceLength *aixmdatatypes.ValDistanceType `xml:"furtherClearanceLength,omitempty"`

	// generated from element "furtherClearanceWidth" of complex type ValDistanceType
	FurtherClearanceWidth *aixmdatatypes.ValDistanceType `xml:"furtherClearanceWidth,omitempty"`

	// generated from element "obscuredLightsSide" of complex type CodeSideType
	ObscuredLightsSide *aixmdatatypes.CodeSideType `xml:"obscuredLightsSide,omitempty"`

	// generated from element "clearedLengthBegin" of complex type ValDistanceType
	ClearedLengthBegin *aixmdatatypes.ValDistanceType `xml:"clearedLengthBegin,omitempty"`

	// generated from element "taxiwayAvailable" of complex type CodeYesNoType
	TaxiwayAvailable *aixmdatatypes.CodeYesNoType `xml:"taxiwayAvailable,omitempty"`

	// generated from element "apronAvailable" of complex type CodeYesNoType
	ApronAvailable *aixmdatatypes.CodeYesNoType `xml:"apronAvailable,omitempty"`
}

// RunwayContaminationPropertyType Named source named complex type "RunwayContaminationPropertyType"
type RunwayContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RunwayContamination" of complex type RunwayContaminationType
	RunwayContamination *RunwayContaminationType `xml:"RunwayContamination,omitempty"`
}

// RunwaySectionContaminationType Named source named complex type "RunwaySectionContaminationType"
type RunwaySectionContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2376 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from group "aixm:RunwaySectionContaminationPropertyGroup" with order 2377 depth 2
	Group_RunwaySectionContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_67 `xml:"extension,omitempty"`
}

// A_extension_67 Named source within outer element "extension"
type A_extension_67 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractRunwaySectionContaminationExtension" of complex type AbstractExtensionType
	AbstractRunwaySectionContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRunwaySectionContaminationExtension,omitempty"`
}

// Group_RunwaySectionContaminationPropertyGroup UnNamed source named group "RunwaySectionContaminationPropertyGroup"
type Group_RunwaySectionContaminationPropertyGroup struct {

	// insertion point for fields

	// generated from element "section" of complex type CodeRunwaySectionType
	Section *aixmdatatypes.CodeRunwaySectionType `xml:"section,omitempty"`
}

// RunwaySectionContaminationPropertyType Named source named complex type "RunwaySectionContaminationPropertyType"
type RunwaySectionContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RunwaySectionContamination" of complex type RunwaySectionContaminationType
	RunwaySectionContamination *RunwaySectionContaminationType `xml:"RunwaySectionContamination,omitempty"`
}

// AbstractSurfaceContaminationType Named source named complex type "AbstractSurfaceContaminationType"
type AbstractSurfaceContaminationType struct {
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

// Group_SurfaceContaminationPropertyGroup UnNamed source named group "SurfaceContaminationPropertyGroup"
type Group_SurfaceContaminationPropertyGroup struct {

	// insertion point for fields

	// generated from element "observationTime" of complex type DateTimeType
	ObservationTime *aixmdatatypes.DateTimeType `xml:"observationTime,omitempty"`

	// generated from element "depth" of complex type ValDepthType
	Depth *aixmdatatypes.ValDepthType `xml:"depth,omitempty"`

	// generated from element "frictionCoefficient" of complex type ValFrictionType
	FrictionCoefficient *aixmdatatypes.ValFrictionType `xml:"frictionCoefficient,omitempty"`

	// generated from element "frictionEstimation" of complex type CodeFrictionEstimateType
	FrictionEstimation *aixmdatatypes.CodeFrictionEstimateType `xml:"frictionEstimation,omitempty"`

	// generated from element "frictionDevice" of complex type CodeFrictionDeviceType
	FrictionDevice *aixmdatatypes.CodeFrictionDeviceType `xml:"frictionDevice,omitempty"`

	// generated from element "obscuredLights" of complex type CodeYesNoType
	ObscuredLights *aixmdatatypes.CodeYesNoType `xml:"obscuredLights,omitempty"`

	// generated from element "furtherClearanceTime" of complex type TimeType
	FurtherClearanceTime *aixmdatatypes.TimeType `xml:"furtherClearanceTime,omitempty"`

	// generated from element "furtherTotalClearance" of complex type CodeYesNoType
	FurtherTotalClearance *aixmdatatypes.CodeYesNoType `xml:"furtherTotalClearance,omitempty"`

	// generated from element "nextObservationTime" of complex type DateTimeType
	NextObservationTime *aixmdatatypes.DateTimeType `xml:"nextObservationTime,omitempty"`

	// generated from element "proportion" of complex type ValPercentType
	Proportion *aixmdatatypes.ValPercentType `xml:"proportion,omitempty"`

	// generated from element "criticalRidge" of complex type RidgePropertyType
	CriticalRidge []*RidgePropertyType `xml:"criticalRidge,omitempty"`

	// generated from element "layer" of complex type SurfaceContaminationLayerPropertyType
	Layer []*SurfaceContaminationLayerPropertyType `xml:"layer,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SurfaceContaminationLayerType Named source named complex type "SurfaceContaminationLayerType"
type SurfaceContaminationLayerType struct {
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

	// generated from group "aixm:SurfaceContaminationLayerPropertyGroup" with order 2416 depth 2
	Group_SurfaceContaminationLayerPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_68 `xml:"extension,omitempty"`
}

// A_extension_68 Named source within outer element "extension"
type A_extension_68 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationLayerExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationLayerExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationLayerExtension,omitempty"`
}

// Group_SurfaceContaminationLayerPropertyGroup UnNamed source named group "SurfaceContaminationLayerPropertyGroup"
type Group_SurfaceContaminationLayerPropertyGroup struct {

	// insertion point for fields

	// generated from element "layerOrder" of complex type NoSequenceType
	LayerOrder *aixmdatatypes.NoSequenceType `xml:"layerOrder,omitempty"`

	// generated from element "type" of complex type CodeContaminationType
	Type_XSD *aixmdatatypes.CodeContaminationType `xml:"type,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent []*ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SurfaceContaminationLayerPropertyType Named source named complex type "SurfaceContaminationLayerPropertyType"
type SurfaceContaminationLayerPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SurfaceContaminationLayer" of complex type SurfaceContaminationLayerType
	SurfaceContaminationLayer *SurfaceContaminationLayerType `xml:"SurfaceContaminationLayer,omitempty"`
}

// TaxiwayContaminationType Named source named complex type "TaxiwayContaminationType"
type TaxiwayContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2435 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from group "aixm:TaxiwayContaminationPropertyGroup" with order 2436 depth 2
	Group_TaxiwayContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_69 `xml:"extension,omitempty"`
}

// A_extension_69 Named source within outer element "extension"
type A_extension_69 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractTaxiwayContaminationExtension" of complex type AbstractExtensionType
	AbstractTaxiwayContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiwayContaminationExtension,omitempty"`
}

// Group_TaxiwayContaminationPropertyGroup UnNamed source named group "TaxiwayContaminationPropertyGroup"
type Group_TaxiwayContaminationPropertyGroup struct {

	// insertion point for fields

	// generated from element "clearedWidth" of complex type ValDistanceType
	ClearedWidth *aixmdatatypes.ValDistanceType `xml:"clearedWidth,omitempty"`
}

// TaxiwayContaminationPropertyType Named source named complex type "TaxiwayContaminationPropertyType"
type TaxiwayContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "TaxiwayContamination" of complex type TaxiwayContaminationType
	TaxiwayContamination *TaxiwayContaminationType `xml:"TaxiwayContamination,omitempty"`
}

// TouchDownLiftOffContaminationType Named source named complex type "TouchDownLiftOffContaminationType"
type TouchDownLiftOffContaminationType struct {
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

	// generated from group "aixm:SurfaceContaminationPropertyGroup" with order 2453 depth 2
	Group_SurfaceContaminationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_70 `xml:"extension,omitempty"`
}

// A_extension_70 Named source within outer element "extension"
type A_extension_70 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurfaceContaminationExtension" of complex type AbstractExtensionType
	AbstractSurfaceContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurfaceContaminationExtension,omitempty"`

	// generated from element "AbstractTouchDownLiftOffContaminationExtension" of complex type AbstractExtensionType
	AbstractTouchDownLiftOffContaminationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTouchDownLiftOffContaminationExtension,omitempty"`
}

// Group_TouchDownLiftOffContaminationPropertyGroup UnNamed source named group "TouchDownLiftOffContaminationPropertyGroup"
type Group_TouchDownLiftOffContaminationPropertyGroup struct {

	// insertion point for fields
}

// TouchDownLiftOffContaminationPropertyType Named source named complex type "TouchDownLiftOffContaminationPropertyType"
type TouchDownLiftOffContaminationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "TouchDownLiftOffContamination" of complex type TouchDownLiftOffContaminationType
	TouchDownLiftOffContamination *TouchDownLiftOffContaminationType `xml:"TouchDownLiftOffContamination,omitempty"`
}

// GuidanceLinePropertyType Named source named complex type "GuidanceLinePropertyType"
type GuidanceLinePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// GuidanceLineType Named source named complex type "GuidanceLineType"
type GuidanceLineType struct {
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

	// generated from element "timeSlice" of complex type GuidanceLineTimeSlicePropertyType
	TimeSlice []*GuidanceLineTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GuidanceLineTimeSlicePropertyType Named source named complex type "GuidanceLineTimeSlicePropertyType"
type GuidanceLineTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GuidanceLineTimeSlice" of complex type GuidanceLineTimeSliceType
	GuidanceLineTimeSlice *GuidanceLineTimeSliceType `xml:"GuidanceLineTimeSlice,omitempty"`
}

// GuidanceLineTimeSliceType Named source named complex type "GuidanceLineTimeSliceType"
type GuidanceLineTimeSliceType struct {
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

	// generated from group "aixm:GuidanceLinePropertyGroup" with order 2481 depth 2
	Group_GuidanceLinePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_71 `xml:"extension,omitempty"`
}

// A_extension_71 Named source within outer element "extension"
type A_extension_71 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGuidanceLineExtension" of complex type AbstractExtensionType
	AbstractGuidanceLineExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGuidanceLineExtension,omitempty"`
}

// Group_GuidanceLinePropertyGroup UnNamed source named group "GuidanceLinePropertyGroup"
type Group_GuidanceLinePropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextNameType
	Designator *aixmdatatypes.TextNameType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeGuidanceLineType
	Type_XSD *aixmdatatypes.CodeGuidanceLineType `xml:"type,omitempty"`

	// generated from element "maxSpeed" of complex type ValSpeedType
	MaxSpeed *aixmdatatypes.ValSpeedType `xml:"maxSpeed,omitempty"`

	// generated from element "usageDirection" of complex type CodeDirectionType
	UsageDirection *aixmdatatypes.CodeDirectionType `xml:"usageDirection,omitempty"`

	// generated from element "connectedTouchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	ConnectedTouchDownLiftOff []*TouchDownLiftOffPropertyType `xml:"connectedTouchDownLiftOff,omitempty"`

	// generated from element "connectedRunwayCentrelinePoint" of complex type RunwayCentrelinePointPropertyType
	ConnectedRunwayCentrelinePoint []*RunwayCentrelinePointPropertyType `xml:"connectedRunwayCentrelinePoint,omitempty"`

	// generated from element "connectedApron" of complex type ApronPropertyType
	ConnectedApron []*ApronPropertyType `xml:"connectedApron,omitempty"`

	// generated from element "connectedStand" of complex type AircraftStandPropertyType
	ConnectedStand []*AircraftStandPropertyType `xml:"connectedStand,omitempty"`

	// generated from element "extent" of complex type ElevatedCurvePropertyType
	Extent *ElevatedCurvePropertyType `xml:"extent,omitempty"`

	// generated from element "connectedTaxiway" of complex type TaxiwayPropertyType
	ConnectedTaxiway []*TaxiwayPropertyType `xml:"connectedTaxiway,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TaxiHoldingPositionPropertyType Named source named complex type "TaxiHoldingPositionPropertyType"
type TaxiHoldingPositionPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// TaxiHoldingPositionType Named source named complex type "TaxiHoldingPositionType"
type TaxiHoldingPositionType struct {
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

	// generated from element "timeSlice" of complex type TaxiHoldingPositionTimeSlicePropertyType
	TimeSlice []*TaxiHoldingPositionTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiHoldingPositionTimeSlicePropertyType Named source named complex type "TaxiHoldingPositionTimeSlicePropertyType"
type TaxiHoldingPositionTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiHoldingPositionTimeSlice" of complex type TaxiHoldingPositionTimeSliceType
	TaxiHoldingPositionTimeSlice *TaxiHoldingPositionTimeSliceType `xml:"TaxiHoldingPositionTimeSlice,omitempty"`
}

// TaxiHoldingPositionTimeSliceType Named source named complex type "TaxiHoldingPositionTimeSliceType"
type TaxiHoldingPositionTimeSliceType struct {
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

	// generated from group "aixm:TaxiHoldingPositionPropertyGroup" with order 2515 depth 2
	Group_TaxiHoldingPositionPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_72 `xml:"extension,omitempty"`
}

// A_extension_72 Named source within outer element "extension"
type A_extension_72 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiHoldingPositionExtension" of complex type AbstractExtensionType
	AbstractTaxiHoldingPositionExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiHoldingPositionExtension,omitempty"`
}

// Group_TaxiHoldingPositionPropertyGroup UnNamed source named group "TaxiHoldingPositionPropertyGroup"
type Group_TaxiHoldingPositionPropertyGroup struct {

	// insertion point for fields

	// generated from element "landingCategory" of complex type CodeHoldingCategoryType
	LandingCategory *aixmdatatypes.CodeHoldingCategoryType `xml:"landingCategory,omitempty"`

	// generated from element "status" of complex type CodeStatusOperationsType
	Status *aixmdatatypes.CodeStatusOperationsType `xml:"status,omitempty"`

	// generated from element "associatedGuidanceLine" of complex type GuidanceLinePropertyType
	AssociatedGuidanceLine *GuidanceLinePropertyType `xml:"associatedGuidanceLine,omitempty"`

	// generated from element "protectedRunway" of complex type RunwayPropertyType
	ProtectedRunway []*RunwayPropertyType `xml:"protectedRunway,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TaxiwayPropertyType Named source named complex type "TaxiwayPropertyType"
type TaxiwayPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// TaxiwayType Named source named complex type "TaxiwayType"
type TaxiwayType struct {
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

	// generated from element "timeSlice" of complex type TaxiwayTimeSlicePropertyType
	TimeSlice []*TaxiwayTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiwayTimeSlicePropertyType Named source named complex type "TaxiwayTimeSlicePropertyType"
type TaxiwayTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiwayTimeSlice" of complex type TaxiwayTimeSliceType
	TaxiwayTimeSlice *TaxiwayTimeSliceType `xml:"TaxiwayTimeSlice,omitempty"`
}

// TaxiwayTimeSliceType Named source named complex type "TaxiwayTimeSliceType"
type TaxiwayTimeSliceType struct {
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

	// generated from group "aixm:TaxiwayPropertyGroup" with order 2544 depth 2
	Group_TaxiwayPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_73 `xml:"extension,omitempty"`
}

// A_extension_73 Named source within outer element "extension"
type A_extension_73 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiwayExtension" of complex type AbstractExtensionType
	AbstractTaxiwayExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiwayExtension,omitempty"`
}

// Group_TaxiwayPropertyGroup UnNamed source named group "TaxiwayPropertyGroup"
type Group_TaxiwayPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type TextDesignatorType
	Designator *aixmdatatypes.TextDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeTaxiwayType
	Type_XSD *aixmdatatypes.CodeTaxiwayType `xml:"type,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "widthShoulder" of complex type ValDistanceType
	WidthShoulder *aixmdatatypes.ValDistanceType `xml:"widthShoulder,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "abandoned" of complex type CodeYesNoType
	Abandoned *aixmdatatypes.CodeYesNoType `xml:"abandoned,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedAirportHeliport" of complex type AirportHeliportPropertyType
	AssociatedAirportHeliport *AirportHeliportPropertyType `xml:"associatedAirportHeliport,omitempty"`

	// generated from element "contaminant" of complex type TaxiwayContaminationPropertyType
	Contaminant []*TaxiwayContaminationPropertyType `xml:"contaminant,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
}

// TaxiwayElementPropertyType Named source named complex type "TaxiwayElementPropertyType"
type TaxiwayElementPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// TaxiwayElementType Named source named complex type "TaxiwayElementType"
type TaxiwayElementType struct {
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

	// generated from element "timeSlice" of complex type TaxiwayElementTimeSlicePropertyType
	TimeSlice []*TaxiwayElementTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TaxiwayElementTimeSlicePropertyType Named source named complex type "TaxiwayElementTimeSlicePropertyType"
type TaxiwayElementTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TaxiwayElementTimeSlice" of complex type TaxiwayElementTimeSliceType
	TaxiwayElementTimeSlice *TaxiwayElementTimeSliceType `xml:"TaxiwayElementTimeSlice,omitempty"`
}

// TaxiwayElementTimeSliceType Named source named complex type "TaxiwayElementTimeSliceType"
type TaxiwayElementTimeSliceType struct {
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

	// generated from group "aixm:TaxiwayElementPropertyGroup" with order 2578 depth 2
	Group_TaxiwayElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_74 `xml:"extension,omitempty"`
}

// A_extension_74 Named source within outer element "extension"
type A_extension_74 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTaxiwayElementExtension" of complex type AbstractExtensionType
	AbstractTaxiwayElementExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTaxiwayElementExtension,omitempty"`
}

// Group_TaxiwayElementPropertyGroup UnNamed source named group "TaxiwayElementPropertyGroup"
type Group_TaxiwayElementPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeTaxiwayElementType
	Type_XSD *aixmdatatypes.CodeTaxiwayElementType `xml:"type,omitempty"`

	// generated from element "length" of complex type ValDistanceType
	Length *aixmdatatypes.ValDistanceType `xml:"length,omitempty"`

	// generated from element "width" of complex type ValDistanceType
	Width *aixmdatatypes.ValDistanceType `xml:"width,omitempty"`

	// generated from element "gradeSeparation" of complex type CodeGradeSeparationType
	GradeSeparation *aixmdatatypes.CodeGradeSeparationType `xml:"gradeSeparation,omitempty"`

	// generated from element "surfaceProperties" of complex type SurfaceCharacteristicsPropertyType
	SurfaceProperties *SurfaceCharacteristicsPropertyType `xml:"surfaceProperties,omitempty"`

	// generated from element "associatedTaxiway" of complex type TaxiwayPropertyType
	AssociatedTaxiway *TaxiwayPropertyType `xml:"associatedTaxiway,omitempty"`

	// generated from element "extent" of complex type ElevatedSurfacePropertyType
	Extent *ElevatedSurfacePropertyType `xml:"extent,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "availability" of complex type ManoeuvringAreaAvailabilityPropertyType
	Availability []*ManoeuvringAreaAvailabilityPropertyType `xml:"availability,omitempty"`
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

	// generated from group "aixm:CurvePropertyGroup" with order 2794 depth 2
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

	// generated from group "aixm:CurvePropertyGroup" with order 2794 depth 2
	Group_CurvePropertyGroup

	// generated from group "aixm:ElevatedCurvePropertyGroup" with order 2805 depth 2
	Group_ElevatedCurvePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_75 `xml:"extension,omitempty"`
}

// A_extension_75 Named source within outer element "extension"
type A_extension_75 struct {
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

	// generated from group "aixm:PointPropertyGroup" with order 2862 depth 2
	Group_PointPropertyGroup

	// generated from group "aixm:ElevatedPointPropertyGroup" with order 2824 depth 2
	Group_ElevatedPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_76 `xml:"extension,omitempty"`
}

// A_extension_76 Named source within outer element "extension"
type A_extension_76 struct {
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2873 depth 2
	Group_SurfacePropertyGroup

	// generated from group "aixm:ElevatedSurfacePropertyGroup" with order 2843 depth 2
	Group_ElevatedSurfacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_77 `xml:"extension,omitempty"`
}

// A_extension_77 Named source within outer element "extension"
type A_extension_77 struct {
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

	// generated from group "aixm:PointPropertyGroup" with order 2862 depth 2
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

	// generated from group "aixm:SurfacePropertyGroup" with order 2873 depth 2
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

// NavaidEquipmentPropertyType Named source named complex type "NavaidEquipmentPropertyType"
type NavaidEquipmentPropertyType struct {
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

	// generated from group "aixm:LinguisticNotePropertyGroup" with order 3906 depth 2
	Group_LinguisticNotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_78 `xml:"extension,omitempty"`
}

// A_extension_78 Named source within outer element "extension"
type A_extension_78 struct {
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

	// generated from group "aixm:NotePropertyGroup" with order 3922 depth 2
	Group_NotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_79 `xml:"extension,omitempty"`
}

// A_extension_79 Named source within outer element "extension"
type A_extension_79 struct {
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

// AirportSuppliesServicePropertyType Named source named complex type "AirportSuppliesServicePropertyType"
type AirportSuppliesServicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
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
	Extension []*A_extension_80 `xml:"extension,omitempty"`
}

// A_extension_80 Named source within outer element "extension"
type A_extension_80 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6257 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:OnlineContactPropertyGroup" with order 6258 depth 2
	Group_OnlineContactPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_81 `xml:"extension,omitempty"`
}

// A_extension_81 Named source within outer element "extension"
type A_extension_81 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6278 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:PostalAddressPropertyGroup" with order 6279 depth 2
	Group_PostalAddressPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_82 `xml:"extension,omitempty"`
}

// A_extension_82 Named source within outer element "extension"
type A_extension_82 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6300 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:TelephoneContactPropertyGroup" with order 6301 depth 2
	Group_TelephoneContactPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_83 `xml:"extension,omitempty"`
}

// A_extension_83 Named source within outer element "extension"
type A_extension_83 struct {
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

	// generated from group "aixm:AircraftCharacteristicPropertyGroup" with order 6319 depth 2
	Group_AircraftCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_84 `xml:"extension,omitempty"`
}

// A_extension_84 Named source within outer element "extension"
type A_extension_84 struct {
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

	// generated from group "aixm:FlightCharacteristicPropertyGroup" with order 6356 depth 2
	Group_FlightCharacteristicPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_85 `xml:"extension,omitempty"`
}

// A_extension_85 Named source within outer element "extension"
type A_extension_85 struct {
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

	// generated from group "aixm:LightElementPropertyGroup" with order 6427 depth 2
	Group_LightElementPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_86 `xml:"extension,omitempty"`
}

// A_extension_86 Named source within outer element "extension"
type A_extension_86 struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6449 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:LightElementStatusPropertyGroup" with order 6450 depth 2
	Group_LightElementStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_87 `xml:"extension,omitempty"`
}

// A_extension_87 Named source within outer element "extension"
type A_extension_87 struct {
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

	// generated from group "aixm:MeteorologyPropertyGroup" with order 6467 depth 2
	Group_MeteorologyPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_88 `xml:"extension,omitempty"`
}

// A_extension_88 Named source within outer element "extension"
type A_extension_88 struct {
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

	// generated from group "aixm:TimesheetPropertyGroup" with order 6564 depth 2
	Group_TimesheetPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_89 `xml:"extension,omitempty"`
}

// A_extension_89 Named source within outer element "extension"
type A_extension_89 struct {
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
