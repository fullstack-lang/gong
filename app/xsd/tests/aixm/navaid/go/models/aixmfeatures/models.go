// generated code - do not edit
package aixmfeatures

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType
var _ *aixmdatatypes.CodeAirspacePointPositionType
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

// TouchDownLiftOffPropertyType Named source named complex type "TouchDownLiftOffPropertyType"
type TouchDownLiftOffPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
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

	// generated from group "aixm:PointPropertyGroup" with order 2837 depth 2
	Group_PointPropertyGroup

	// generated from group "aixm:ElevatedPointPropertyGroup" with order 2799 depth 2
	Group_ElevatedPointPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension `xml:"extension,omitempty"`
}

// A_extension Named source within outer element "extension"
type A_extension struct {
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

	// generated from group "aixm:PointPropertyGroup" with order 2837 depth 2
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

// AuthorityForNavaidEquipmentType Named source named complex type "AuthorityForNavaidEquipmentType"
type AuthorityForNavaidEquipmentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:AuthorityForNavaidEquipmentPropertyGroup" with order 2964 depth 2
	Group_AuthorityForNavaidEquipmentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_1 `xml:"extension,omitempty"`
}

// A_extension_1 Named source within outer element "extension"
type A_extension_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAuthorityForNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractAuthorityForNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAuthorityForNavaidEquipmentExtension,omitempty"`
}

// Group_AuthorityForNavaidEquipmentPropertyGroup UnNamed source named group "AuthorityForNavaidEquipmentPropertyGroup"
type Group_AuthorityForNavaidEquipmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAuthorityRoleType
	Type_XSD *aixmdatatypes.CodeAuthorityRoleType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// AuthorityForNavaidEquipmentPropertyType Named source named complex type "AuthorityForNavaidEquipmentPropertyType"
type AuthorityForNavaidEquipmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AuthorityForNavaidEquipment" of complex type AuthorityForNavaidEquipmentType
	AuthorityForNavaidEquipment *AuthorityForNavaidEquipmentType `xml:"AuthorityForNavaidEquipment,omitempty"`
}

// AuthorityForSpecialNavigationStationType Named source named complex type "AuthorityForSpecialNavigationStationType"
type AuthorityForSpecialNavigationStationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:AuthorityForSpecialNavigationStationPropertyGroup" with order 2982 depth 2
	Group_AuthorityForSpecialNavigationStationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_2 `xml:"extension,omitempty"`
}

// A_extension_2 Named source within outer element "extension"
type A_extension_2 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAuthorityForSpecialNavigationStationExtension" of complex type AbstractExtensionType
	AbstractAuthorityForSpecialNavigationStationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAuthorityForSpecialNavigationStationExtension,omitempty"`
}

// Group_AuthorityForSpecialNavigationStationPropertyGroup UnNamed source named group "AuthorityForSpecialNavigationStationPropertyGroup"
type Group_AuthorityForSpecialNavigationStationPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAuthorityRoleType
	Type_XSD *aixmdatatypes.CodeAuthorityRoleType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// AuthorityForSpecialNavigationStationPropertyType Named source named complex type "AuthorityForSpecialNavigationStationPropertyType"
type AuthorityForSpecialNavigationStationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AuthorityForSpecialNavigationStation" of complex type AuthorityForSpecialNavigationStationType
	AuthorityForSpecialNavigationStation *AuthorityForSpecialNavigationStationType `xml:"AuthorityForSpecialNavigationStation,omitempty"`
}

// AuthorityForSpecialNavigationSystemType Named source named complex type "AuthorityForSpecialNavigationSystemType"
type AuthorityForSpecialNavigationSystemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:AuthorityForSpecialNavigationSystemPropertyGroup" with order 3000 depth 2
	Group_AuthorityForSpecialNavigationSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_3 `xml:"extension,omitempty"`
}

// A_extension_3 Named source within outer element "extension"
type A_extension_3 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAuthorityForSpecialNavigationSystemExtension" of complex type AbstractExtensionType
	AbstractAuthorityForSpecialNavigationSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAuthorityForSpecialNavigationSystemExtension,omitempty"`
}

// Group_AuthorityForSpecialNavigationSystemPropertyGroup UnNamed source named group "AuthorityForSpecialNavigationSystemPropertyGroup"
type Group_AuthorityForSpecialNavigationSystemPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAuthorityRoleType
	Type_XSD *aixmdatatypes.CodeAuthorityRoleType `xml:"type,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theOrganisationAuthority" of complex type OrganisationAuthorityPropertyType
	TheOrganisationAuthority *OrganisationAuthorityPropertyType `xml:"theOrganisationAuthority,omitempty"`
}

// AuthorityForSpecialNavigationSystemPropertyType Named source named complex type "AuthorityForSpecialNavigationSystemPropertyType"
type AuthorityForSpecialNavigationSystemPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "AuthorityForSpecialNavigationSystem" of complex type AuthorityForSpecialNavigationSystemType
	AuthorityForSpecialNavigationSystem *AuthorityForSpecialNavigationSystemType `xml:"AuthorityForSpecialNavigationSystem,omitempty"`
}

// AzimuthType Named source named complex type "AzimuthType"
type AzimuthType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type AzimuthTimeSlicePropertyType
	TimeSlice []*AzimuthTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AzimuthTimeSlicePropertyType Named source named complex type "AzimuthTimeSlicePropertyType"
type AzimuthTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AzimuthTimeSlice" of complex type AzimuthTimeSliceType
	AzimuthTimeSlice *AzimuthTimeSliceType `xml:"AzimuthTimeSlice,omitempty"`
}

// AzimuthTimeSliceType Named source named complex type "AzimuthTimeSliceType"
type AzimuthTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeMLSAzimuthType
	Type_XSD *aixmdatatypes.CodeMLSAzimuthType `xml:"type,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from element "trueBearingAccuracy" of complex type ValAngleType
	TrueBearingAccuracy *aixmdatatypes.ValAngleType `xml:"trueBearingAccuracy,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "angleProportionalLeft" of complex type ValAngleType
	AngleProportionalLeft *aixmdatatypes.ValAngleType `xml:"angleProportionalLeft,omitempty"`

	// generated from element "angleProportionalRight" of complex type ValAngleType
	AngleProportionalRight *aixmdatatypes.ValAngleType `xml:"angleProportionalRight,omitempty"`

	// generated from element "angleCoverLeft" of complex type ValAngleType
	AngleCoverLeft *aixmdatatypes.ValAngleType `xml:"angleCoverLeft,omitempty"`

	// generated from element "angleCoverRight" of complex type ValAngleType
	AngleCoverRight *aixmdatatypes.ValAngleType `xml:"angleCoverRight,omitempty"`

	// generated from element "channel" of complex type CodeMLSChannelType
	Channel *aixmdatatypes.CodeMLSChannelType `xml:"channel,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_4 `xml:"extension,omitempty"`
}

// A_extension_4 Named source within outer element "extension"
type A_extension_4 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAzimuthExtension" of complex type AbstractExtensionType
	AbstractAzimuthExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAzimuthExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_AzimuthPropertyGroup UnNamed source named group "AzimuthPropertyGroup"
type Group_AzimuthPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3038 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "type" of complex type CodeMLSAzimuthType
	Type_XSD *aixmdatatypes.CodeMLSAzimuthType `xml:"type,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from element "trueBearingAccuracy" of complex type ValAngleType
	TrueBearingAccuracy *aixmdatatypes.ValAngleType `xml:"trueBearingAccuracy,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "angleProportionalLeft" of complex type ValAngleType
	AngleProportionalLeft *aixmdatatypes.ValAngleType `xml:"angleProportionalLeft,omitempty"`

	// generated from element "angleProportionalRight" of complex type ValAngleType
	AngleProportionalRight *aixmdatatypes.ValAngleType `xml:"angleProportionalRight,omitempty"`

	// generated from element "angleCoverLeft" of complex type ValAngleType
	AngleCoverLeft *aixmdatatypes.ValAngleType `xml:"angleCoverLeft,omitempty"`

	// generated from element "angleCoverRight" of complex type ValAngleType
	AngleCoverRight *aixmdatatypes.ValAngleType `xml:"angleCoverRight,omitempty"`

	// generated from element "channel" of complex type CodeMLSChannelType
	Channel *aixmdatatypes.CodeMLSChannelType `xml:"channel,omitempty"`
}

// CheckpointINSType Named source named complex type "CheckpointINSType"
type CheckpointINSType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type CheckpointINSTimeSlicePropertyType
	TimeSlice []*CheckpointINSTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// CheckpointINSTimeSlicePropertyType Named source named complex type "CheckpointINSTimeSlicePropertyType"
type CheckpointINSTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "CheckpointINSTimeSlice" of complex type CheckpointINSTimeSliceType
	CheckpointINSTimeSlice *CheckpointINSTimeSliceType `xml:"CheckpointINSTimeSlice,omitempty"`
}

// CheckpointINSTimeSliceType Named source named complex type "CheckpointINSTimeSliceType"
type CheckpointINSTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:CheckpointINSPropertyGroup" with order 3063 depth 2
	Group_CheckpointINSPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_5 `xml:"extension,omitempty"`
}

// A_extension_5 Named source within outer element "extension"
type A_extension_5 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCheckpointINSExtension" of complex type AbstractExtensionType
	AbstractCheckpointINSExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCheckpointINSExtension,omitempty"`

	// generated from element "AbstractNavigationSystemCheckpointExtension" of complex type AbstractExtensionType
	AbstractNavigationSystemCheckpointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavigationSystemCheckpointExtension,omitempty"`
}

// Group_CheckpointINSPropertyGroup UnNamed source named group "CheckpointINSPropertyGroup"
type Group_CheckpointINSPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavigationSystemCheckpointPropertyGroup" with order 3072 depth 2
	Group_NavigationSystemCheckpointPropertyGroup
}

// CheckpointVORType Named source named complex type "CheckpointVORType"
type CheckpointVORType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type CheckpointVORTimeSlicePropertyType
	TimeSlice []*CheckpointVORTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// CheckpointVORTimeSlicePropertyType Named source named complex type "CheckpointVORTimeSlicePropertyType"
type CheckpointVORTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "CheckpointVORTimeSlice" of complex type CheckpointVORTimeSliceType
	CheckpointVORTimeSlice *CheckpointVORTimeSliceType `xml:"CheckpointVORTimeSlice,omitempty"`
}

// CheckpointVORTimeSliceType Named source named complex type "CheckpointVORTimeSliceType"
type CheckpointVORTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:CheckpointVORPropertyGroup" with order 3088 depth 2
	Group_CheckpointVORPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_6 `xml:"extension,omitempty"`
}

// A_extension_6 Named source within outer element "extension"
type A_extension_6 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractCheckpointVORExtension" of complex type AbstractExtensionType
	AbstractCheckpointVORExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractCheckpointVORExtension,omitempty"`

	// generated from element "AbstractNavigationSystemCheckpointExtension" of complex type AbstractExtensionType
	AbstractNavigationSystemCheckpointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavigationSystemCheckpointExtension,omitempty"`
}

// Group_CheckpointVORPropertyGroup UnNamed source named group "CheckpointVORPropertyGroup"
type Group_CheckpointVORPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavigationSystemCheckpointPropertyGroup" with order 3097 depth 2
	Group_NavigationSystemCheckpointPropertyGroup

	// generated from element "checkPointFacility" of complex type VORPropertyType
	CheckPointFacility *VORPropertyType `xml:"checkPointFacility,omitempty"`
}

// DirectionFinderType Named source named complex type "DirectionFinderType"
type DirectionFinderType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type DirectionFinderTimeSlicePropertyType
	TimeSlice []*DirectionFinderTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DirectionFinderTimeSlicePropertyType Named source named complex type "DirectionFinderTimeSlicePropertyType"
type DirectionFinderTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DirectionFinderTimeSlice" of complex type DirectionFinderTimeSliceType
	DirectionFinderTimeSlice *DirectionFinderTimeSliceType `xml:"DirectionFinderTimeSlice,omitempty"`
}

// DirectionFinderTimeSliceType Named source named complex type "DirectionFinderTimeSliceType"
type DirectionFinderTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "doppler" of complex type CodeYesNoType
	Doppler *aixmdatatypes.CodeYesNoType `xml:"doppler,omitempty"`

	// generated from element "informationProvision" of complex type InformationServicePropertyType
	InformationProvision []*InformationServicePropertyType `xml:"informationProvision,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_7 `xml:"extension,omitempty"`
}

// A_extension_7 Named source within outer element "extension"
type A_extension_7 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDirectionFinderExtension" of complex type AbstractExtensionType
	AbstractDirectionFinderExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDirectionFinderExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_DirectionFinderPropertyGroup UnNamed source named group "DirectionFinderPropertyGroup"
type Group_DirectionFinderPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3123 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "doppler" of complex type CodeYesNoType
	Doppler *aixmdatatypes.CodeYesNoType `xml:"doppler,omitempty"`

	// generated from element "informationProvision" of complex type InformationServicePropertyType
	InformationProvision []*InformationServicePropertyType `xml:"informationProvision,omitempty"`
}

// DMEType Named source named complex type "DMEType"
type DMEType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type DMETimeSlicePropertyType
	TimeSlice []*DMETimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DMETimeSlicePropertyType Named source named complex type "DMETimeSlicePropertyType"
type DMETimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DMETimeSlice" of complex type DMETimeSliceType
	DMETimeSlice *DMETimeSliceType `xml:"DMETimeSlice,omitempty"`
}

// DMETimeSliceType Named source named complex type "DMETimeSliceType"
type DMETimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeDMEType
	Type_XSD *aixmdatatypes.CodeDMEType `xml:"type,omitempty"`

	// generated from element "channel" of complex type CodeDMEChannelType
	Channel *aixmdatatypes.CodeDMEChannelType `xml:"channel,omitempty"`

	// generated from element "ghostFrequency" of complex type ValFrequencyType
	GhostFrequency *aixmdatatypes.ValFrequencyType `xml:"ghostFrequency,omitempty"`

	// generated from element "displace" of complex type ValDistanceType
	Displace *aixmdatatypes.ValDistanceType `xml:"displace,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_8 `xml:"extension,omitempty"`
}

// A_extension_8 Named source within outer element "extension"
type A_extension_8 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDMEExtension" of complex type AbstractExtensionType
	AbstractDMEExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDMEExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_DMEPropertyGroup UnNamed source named group "DMEPropertyGroup"
type Group_DMEPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3150 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "type" of complex type CodeDMEType
	Type_XSD *aixmdatatypes.CodeDMEType `xml:"type,omitempty"`

	// generated from element "channel" of complex type CodeDMEChannelType
	Channel *aixmdatatypes.CodeDMEChannelType `xml:"channel,omitempty"`

	// generated from element "ghostFrequency" of complex type ValFrequencyType
	GhostFrequency *aixmdatatypes.ValFrequencyType `xml:"ghostFrequency,omitempty"`

	// generated from element "displace" of complex type ValDistanceType
	Displace *aixmdatatypes.ValDistanceType `xml:"displace,omitempty"`
}

// ElevationType Named source named complex type "ElevationType"
type ElevationType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type ElevationTimeSlicePropertyType
	TimeSlice []*ElevationTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// ElevationTimeSlicePropertyType Named source named complex type "ElevationTimeSlicePropertyType"
type ElevationTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "ElevationTimeSlice" of complex type ElevationTimeSliceType
	ElevationTimeSlice *ElevationTimeSliceType `xml:"ElevationTimeSlice,omitempty"`
}

// ElevationTimeSliceType Named source named complex type "ElevationTimeSliceType"
type ElevationTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "angleNominal" of complex type ValAngleType
	AngleNominal *aixmdatatypes.ValAngleType `xml:"angleNominal,omitempty"`

	// generated from element "angleMinimum" of complex type ValAngleType
	AngleMinimum *aixmdatatypes.ValAngleType `xml:"angleMinimum,omitempty"`

	// generated from element "angleSpan" of complex type ValAngleType
	AngleSpan *aixmdatatypes.ValAngleType `xml:"angleSpan,omitempty"`

	// generated from element "angleAccuracy" of complex type ValAngleType
	AngleAccuracy *aixmdatatypes.ValAngleType `xml:"angleAccuracy,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_9 `xml:"extension,omitempty"`
}

// A_extension_9 Named source within outer element "extension"
type A_extension_9 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractElevationExtension" of complex type AbstractExtensionType
	AbstractElevationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractElevationExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_ElevationPropertyGroup UnNamed source named group "ElevationPropertyGroup"
type Group_ElevationPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3179 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "angleNominal" of complex type ValAngleType
	AngleNominal *aixmdatatypes.ValAngleType `xml:"angleNominal,omitempty"`

	// generated from element "angleMinimum" of complex type ValAngleType
	AngleMinimum *aixmdatatypes.ValAngleType `xml:"angleMinimum,omitempty"`

	// generated from element "angleSpan" of complex type ValAngleType
	AngleSpan *aixmdatatypes.ValAngleType `xml:"angleSpan,omitempty"`

	// generated from element "angleAccuracy" of complex type ValAngleType
	AngleAccuracy *aixmdatatypes.ValAngleType `xml:"angleAccuracy,omitempty"`
}

// GlidepathType Named source named complex type "GlidepathType"
type GlidepathType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type GlidepathTimeSlicePropertyType
	TimeSlice []*GlidepathTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// GlidepathTimeSlicePropertyType Named source named complex type "GlidepathTimeSlicePropertyType"
type GlidepathTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "GlidepathTimeSlice" of complex type GlidepathTimeSliceType
	GlidepathTimeSlice *GlidepathTimeSliceType `xml:"GlidepathTimeSlice,omitempty"`
}

// GlidepathTimeSliceType Named source named complex type "GlidepathTimeSliceType"
type GlidepathTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "slope" of complex type ValAngleType
	Slope *aixmdatatypes.ValAngleType `xml:"slope,omitempty"`

	// generated from element "angleAccuracy" of complex type ValAngleType
	AngleAccuracy *aixmdatatypes.ValAngleType `xml:"angleAccuracy,omitempty"`

	// generated from element "rdh" of complex type ValDistanceVerticalType
	Rdh *aixmdatatypes.ValDistanceVerticalType `xml:"rdh,omitempty"`

	// generated from element "rdhAccuracy" of complex type ValDistanceVerticalType
	RdhAccuracy *aixmdatatypes.ValDistanceVerticalType `xml:"rdhAccuracy,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_10 `xml:"extension,omitempty"`
}

// A_extension_10 Named source within outer element "extension"
type A_extension_10 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractGlidepathExtension" of complex type AbstractExtensionType
	AbstractGlidepathExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractGlidepathExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_GlidepathPropertyGroup UnNamed source named group "GlidepathPropertyGroup"
type Group_GlidepathPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3208 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "slope" of complex type ValAngleType
	Slope *aixmdatatypes.ValAngleType `xml:"slope,omitempty"`

	// generated from element "angleAccuracy" of complex type ValAngleType
	AngleAccuracy *aixmdatatypes.ValAngleType `xml:"angleAccuracy,omitempty"`

	// generated from element "rdh" of complex type ValDistanceVerticalType
	Rdh *aixmdatatypes.ValDistanceVerticalType `xml:"rdh,omitempty"`

	// generated from element "rdhAccuracy" of complex type ValDistanceVerticalType
	RdhAccuracy *aixmdatatypes.ValDistanceVerticalType `xml:"rdhAccuracy,omitempty"`
}

// LocalizerType Named source named complex type "LocalizerType"
type LocalizerType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type LocalizerTimeSlicePropertyType
	TimeSlice []*LocalizerTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// LocalizerTimeSlicePropertyType Named source named complex type "LocalizerTimeSlicePropertyType"
type LocalizerTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "LocalizerTimeSlice" of complex type LocalizerTimeSliceType
	LocalizerTimeSlice *LocalizerTimeSliceType `xml:"LocalizerTimeSlice,omitempty"`
}

// LocalizerTimeSliceType Named source named complex type "LocalizerTimeSliceType"
type LocalizerTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "magneticBearingAccuracy" of complex type ValAngleType
	MagneticBearingAccuracy *aixmdatatypes.ValAngleType `xml:"magneticBearingAccuracy,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from element "trueBearingAccuracy" of complex type ValAngleType
	TrueBearingAccuracy *aixmdatatypes.ValAngleType `xml:"trueBearingAccuracy,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`

	// generated from element "widthCourse" of complex type ValAngleType
	WidthCourse *aixmdatatypes.ValAngleType `xml:"widthCourse,omitempty"`

	// generated from element "widthCourseAccuracy" of complex type ValAngleType
	WidthCourseAccuracy *aixmdatatypes.ValAngleType `xml:"widthCourseAccuracy,omitempty"`

	// generated from element "backCourseUsable" of complex type CodeILSBackCourseType
	BackCourseUsable *aixmdatatypes.CodeILSBackCourseType `xml:"backCourseUsable,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_11 `xml:"extension,omitempty"`
}

// A_extension_11 Named source within outer element "extension"
type A_extension_11 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractLocalizerExtension" of complex type AbstractExtensionType
	AbstractLocalizerExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractLocalizerExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_LocalizerPropertyGroup UnNamed source named group "LocalizerPropertyGroup"
type Group_LocalizerPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3238 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "magneticBearingAccuracy" of complex type ValAngleType
	MagneticBearingAccuracy *aixmdatatypes.ValAngleType `xml:"magneticBearingAccuracy,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from element "trueBearingAccuracy" of complex type ValAngleType
	TrueBearingAccuracy *aixmdatatypes.ValAngleType `xml:"trueBearingAccuracy,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`

	// generated from element "widthCourse" of complex type ValAngleType
	WidthCourse *aixmdatatypes.ValAngleType `xml:"widthCourse,omitempty"`

	// generated from element "widthCourseAccuracy" of complex type ValAngleType
	WidthCourseAccuracy *aixmdatatypes.ValAngleType `xml:"widthCourseAccuracy,omitempty"`

	// generated from element "backCourseUsable" of complex type CodeILSBackCourseType
	BackCourseUsable *aixmdatatypes.CodeILSBackCourseType `xml:"backCourseUsable,omitempty"`
}

// MarkerBeaconType Named source named complex type "MarkerBeaconType"
type MarkerBeaconType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type MarkerBeaconTimeSlicePropertyType
	TimeSlice []*MarkerBeaconTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// MarkerBeaconTimeSlicePropertyType Named source named complex type "MarkerBeaconTimeSlicePropertyType"
type MarkerBeaconTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "MarkerBeaconTimeSlice" of complex type MarkerBeaconTimeSliceType
	MarkerBeaconTimeSlice *MarkerBeaconTimeSliceType `xml:"MarkerBeaconTimeSlice,omitempty"`
}

// MarkerBeaconTimeSliceType Named source named complex type "MarkerBeaconTimeSliceType"
type MarkerBeaconTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "class" of complex type CodeMarkerBeaconSignalType
	Class *aixmdatatypes.CodeMarkerBeaconSignalType `xml:"class,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "axisBearing" of complex type ValBearingType
	AxisBearing *aixmdatatypes.ValBearingType `xml:"axisBearing,omitempty"`

	// generated from element "auralMorseCode" of complex type CodeAuralMorseType
	AuralMorseCode *aixmdatatypes.CodeAuralMorseType `xml:"auralMorseCode,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_12 `xml:"extension,omitempty"`
}

// A_extension_12 Named source within outer element "extension"
type A_extension_12 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractMarkerBeaconExtension" of complex type AbstractExtensionType
	AbstractMarkerBeaconExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractMarkerBeaconExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_MarkerBeaconPropertyGroup UnNamed source named group "MarkerBeaconPropertyGroup"
type Group_MarkerBeaconPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3272 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "class" of complex type CodeMarkerBeaconSignalType
	Class *aixmdatatypes.CodeMarkerBeaconSignalType `xml:"class,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "axisBearing" of complex type ValBearingType
	AxisBearing *aixmdatatypes.ValBearingType `xml:"axisBearing,omitempty"`

	// generated from element "auralMorseCode" of complex type CodeAuralMorseType
	AuralMorseCode *aixmdatatypes.CodeAuralMorseType `xml:"auralMorseCode,omitempty"`
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

// NavaidType Named source named complex type "NavaidType"
type NavaidType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type NavaidTimeSlicePropertyType
	TimeSlice []*NavaidTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// NavaidTimeSlicePropertyType Named source named complex type "NavaidTimeSlicePropertyType"
type NavaidTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "NavaidTimeSlice" of complex type NavaidTimeSliceType
	NavaidTimeSlice *NavaidTimeSliceType `xml:"NavaidTimeSlice,omitempty"`
}

// NavaidTimeSliceType Named source named complex type "NavaidTimeSliceType"
type NavaidTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "type" of complex type CodeNavaidServiceType
	Type_XSD *aixmdatatypes.CodeNavaidServiceType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "purpose" of complex type CodeNavaidPurposeType
	Purpose *aixmdatatypes.CodeNavaidPurposeType `xml:"purpose,omitempty"`

	// generated from element "signalPerformance" of complex type CodeSignalPerformanceILSType
	SignalPerformance *aixmdatatypes.CodeSignalPerformanceILSType `xml:"signalPerformance,omitempty"`

	// generated from element "courseQuality" of complex type CodeCourseQualityILSType
	CourseQuality *aixmdatatypes.CodeCourseQualityILSType `xml:"courseQuality,omitempty"`

	// generated from element "integrityLevel" of complex type CodeIntegrityLevelILSType
	IntegrityLevel *aixmdatatypes.CodeIntegrityLevelILSType `xml:"integrityLevel,omitempty"`

	// generated from element "touchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	TouchDownLiftOff []*TouchDownLiftOffPropertyType `xml:"touchDownLiftOff,omitempty"`

	// generated from element "navaidEquipment" of complex type NavaidComponentPropertyType
	NavaidEquipment []*NavaidComponentPropertyType `xml:"navaidEquipment,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "runwayDirection" of complex type RunwayDirectionPropertyType
	RunwayDirection []*RunwayDirectionPropertyType `xml:"runwayDirection,omitempty"`

	// generated from element "servedAirport" of complex type AirportHeliportPropertyType
	ServedAirport []*AirportHeliportPropertyType `xml:"servedAirport,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_13 `xml:"extension,omitempty"`
}

// A_extension_13 Named source within outer element "extension"
type A_extension_13 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavaidExtension" of complex type AbstractExtensionType
	AbstractNavaidExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidExtension,omitempty"`
}

// Group_NavaidPropertyGroup UnNamed source named group "NavaidPropertyGroup"
type Group_NavaidPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeNavaidServiceType
	Type_XSD *aixmdatatypes.CodeNavaidServiceType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "purpose" of complex type CodeNavaidPurposeType
	Purpose *aixmdatatypes.CodeNavaidPurposeType `xml:"purpose,omitempty"`

	// generated from element "signalPerformance" of complex type CodeSignalPerformanceILSType
	SignalPerformance *aixmdatatypes.CodeSignalPerformanceILSType `xml:"signalPerformance,omitempty"`

	// generated from element "courseQuality" of complex type CodeCourseQualityILSType
	CourseQuality *aixmdatatypes.CodeCourseQualityILSType `xml:"courseQuality,omitempty"`

	// generated from element "integrityLevel" of complex type CodeIntegrityLevelILSType
	IntegrityLevel *aixmdatatypes.CodeIntegrityLevelILSType `xml:"integrityLevel,omitempty"`

	// generated from element "touchDownLiftOff" of complex type TouchDownLiftOffPropertyType
	TouchDownLiftOff []*TouchDownLiftOffPropertyType `xml:"touchDownLiftOff,omitempty"`

	// generated from element "navaidEquipment" of complex type NavaidComponentPropertyType
	NavaidEquipment []*NavaidComponentPropertyType `xml:"navaidEquipment,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "runwayDirection" of complex type RunwayDirectionPropertyType
	RunwayDirection []*RunwayDirectionPropertyType `xml:"runwayDirection,omitempty"`

	// generated from element "servedAirport" of complex type AirportHeliportPropertyType
	ServedAirport []*AirportHeliportPropertyType `xml:"servedAirport,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NavaidComponentType Named source named complex type "NavaidComponentType"
type NavaidComponentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:NavaidComponentPropertyGroup" with order 3319 depth 2
	Group_NavaidComponentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_14 `xml:"extension,omitempty"`
}

// A_extension_14 Named source within outer element "extension"
type A_extension_14 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNavaidComponentExtension" of complex type AbstractExtensionType
	AbstractNavaidComponentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidComponentExtension,omitempty"`
}

// Group_NavaidComponentPropertyGroup UnNamed source named group "NavaidComponentPropertyGroup"
type Group_NavaidComponentPropertyGroup struct {

	// insertion point for fields

	// generated from element "collocationGroup" of complex type NoSequenceType
	CollocationGroup *aixmdatatypes.NoSequenceType `xml:"collocationGroup,omitempty"`

	// generated from element "markerPosition" of complex type CodePositionInILSType
	MarkerPosition *aixmdatatypes.CodePositionInILSType `xml:"markerPosition,omitempty"`

	// generated from element "providesNavigableLocation" of complex type CodeYesNoType
	ProvidesNavigableLocation *aixmdatatypes.CodeYesNoType `xml:"providesNavigableLocation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theNavaidEquipment" of complex type NavaidEquipmentPropertyType
	TheNavaidEquipment *NavaidEquipmentPropertyType `xml:"theNavaidEquipment,omitempty"`
}

// NavaidComponentPropertyType Named source named complex type "NavaidComponentPropertyType"
type NavaidComponentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "NavaidComponent" of complex type NavaidComponentType
	NavaidComponent *NavaidComponentType `xml:"NavaidComponent,omitempty"`
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

// AbstractNavaidEquipmentType Named source named complex type "AbstractNavaidEquipmentType"
type AbstractNavaidEquipmentType struct {
	Name string `xml:"-"`

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

// Group_NavaidEquipmentPropertyGroup UnNamed source named group "NavaidEquipmentPropertyGroup"
type Group_NavaidEquipmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NavaidEquipmentMonitoringType Named source named complex type "NavaidEquipmentMonitoringType"
type NavaidEquipmentMonitoringType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 3360 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:NavaidEquipmentMonitoringPropertyGroup" with order 3361 depth 2
	Group_NavaidEquipmentMonitoringPropertyGroup

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

	// generated from element "AbstractNavaidEquipmentMonitoringExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentMonitoringExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentMonitoringExtension,omitempty"`
}

// Group_NavaidEquipmentMonitoringPropertyGroup UnNamed source named group "NavaidEquipmentMonitoringPropertyGroup"
type Group_NavaidEquipmentMonitoringPropertyGroup struct {

	// insertion point for fields

	// generated from element "monitored" of complex type CodeYesNoType
	Monitored *aixmdatatypes.CodeYesNoType `xml:"monitored,omitempty"`
}

// NavaidEquipmentMonitoringPropertyType Named source named complex type "NavaidEquipmentMonitoringPropertyType"
type NavaidEquipmentMonitoringPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "NavaidEquipmentMonitoring" of complex type NavaidEquipmentMonitoringType
	NavaidEquipmentMonitoring *NavaidEquipmentMonitoringType `xml:"NavaidEquipmentMonitoring,omitempty"`
}

// NavaidOperationalStatusType Named source named complex type "NavaidOperationalStatusType"
type NavaidOperationalStatusType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 3378 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:NavaidOperationalStatusPropertyGroup" with order 3379 depth 2
	Group_NavaidOperationalStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_16 `xml:"extension,omitempty"`
}

// A_extension_16 Named source within outer element "extension"
type A_extension_16 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractNavaidOperationalStatusExtension" of complex type AbstractExtensionType
	AbstractNavaidOperationalStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidOperationalStatusExtension,omitempty"`
}

// Group_NavaidOperationalStatusPropertyGroup UnNamed source named group "NavaidOperationalStatusPropertyGroup"
type Group_NavaidOperationalStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusNavaidType
	OperationalStatus *aixmdatatypes.CodeStatusNavaidType `xml:"operationalStatus,omitempty"`

	// generated from element "signalType" of complex type CodeRadioSignalType
	SignalType *aixmdatatypes.CodeRadioSignalType `xml:"signalType,omitempty"`
}

// NavaidOperationalStatusPropertyType Named source named complex type "NavaidOperationalStatusPropertyType"
type NavaidOperationalStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "NavaidOperationalStatus" of complex type NavaidOperationalStatusType
	NavaidOperationalStatus *NavaidOperationalStatusType `xml:"NavaidOperationalStatus,omitempty"`
}

// AbstractNavigationSystemCheckpointType Named source named complex type "AbstractNavigationSystemCheckpointType"
type AbstractNavigationSystemCheckpointType struct {
	Name string `xml:"-"`

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

// Group_NavigationSystemCheckpointPropertyGroup UnNamed source named group "NavigationSystemCheckpointPropertyGroup"
type Group_NavigationSystemCheckpointPropertyGroup struct {

	// insertion point for fields

	// generated from element "category" of complex type CodeCheckpointCategoryType
	Category *aixmdatatypes.CodeCheckpointCategoryType `xml:"category,omitempty"`

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

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "angle" of complex type ValBearingType
	Angle *aixmdatatypes.ValBearingType `xml:"angle,omitempty"`

	// generated from element "position" of complex type ElevatedPointPropertyType
	Position *ElevatedPointPropertyType `xml:"position,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport *AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// NDBType Named source named complex type "NDBType"
type NDBType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type NDBTimeSlicePropertyType
	TimeSlice []*NDBTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// NDBTimeSlicePropertyType Named source named complex type "NDBTimeSlicePropertyType"
type NDBTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "NDBTimeSlice" of complex type NDBTimeSliceType
	NDBTimeSlice *NDBTimeSliceType `xml:"NDBTimeSlice,omitempty"`
}

// NDBTimeSliceType Named source named complex type "NDBTimeSliceType"
type NDBTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "class" of complex type CodeNDBUsageType
	Class *aixmdatatypes.CodeNDBUsageType `xml:"class,omitempty"`

	// generated from element "emissionBand" of complex type CodeEmissionBandType
	EmissionBand *aixmdatatypes.CodeEmissionBandType `xml:"emissionBand,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_17 `xml:"extension,omitempty"`
}

// A_extension_17 Named source within outer element "extension"
type A_extension_17 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractNDBExtension" of complex type AbstractExtensionType
	AbstractNDBExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNDBExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_NDBPropertyGroup UnNamed source named group "NDBPropertyGroup"
type Group_NDBPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3436 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "class" of complex type CodeNDBUsageType
	Class *aixmdatatypes.CodeNDBUsageType `xml:"class,omitempty"`

	// generated from element "emissionBand" of complex type CodeEmissionBandType
	EmissionBand *aixmdatatypes.CodeEmissionBandType `xml:"emissionBand,omitempty"`
}

// SDFType Named source named complex type "SDFType"
type SDFType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type SDFTimeSlicePropertyType
	TimeSlice []*SDFTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SDFTimeSlicePropertyType Named source named complex type "SDFTimeSlicePropertyType"
type SDFTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SDFTimeSlice" of complex type SDFTimeSliceType
	SDFTimeSlice *SDFTimeSliceType `xml:"SDFTimeSlice,omitempty"`
}

// SDFTimeSliceType Named source named complex type "SDFTimeSliceType"
type SDFTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_18 `xml:"extension,omitempty"`
}

// A_extension_18 Named source within outer element "extension"
type A_extension_18 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSDFExtension" of complex type AbstractExtensionType
	AbstractSDFExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSDFExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_SDFPropertyGroup UnNamed source named group "SDFPropertyGroup"
type Group_SDFPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3464 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "magneticBearing" of complex type ValBearingType
	MagneticBearing *aixmdatatypes.ValBearingType `xml:"magneticBearing,omitempty"`

	// generated from element "trueBearing" of complex type ValBearingType
	TrueBearing *aixmdatatypes.ValBearingType `xml:"trueBearing,omitempty"`
}

// SpecialNavigationStationType Named source named complex type "SpecialNavigationStationType"
type SpecialNavigationStationType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type SpecialNavigationStationTimeSlicePropertyType
	TimeSlice []*SpecialNavigationStationTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SpecialNavigationStationTimeSlicePropertyType Named source named complex type "SpecialNavigationStationTimeSlicePropertyType"
type SpecialNavigationStationTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SpecialNavigationStationTimeSlice" of complex type SpecialNavigationStationTimeSliceType
	SpecialNavigationStationTimeSlice *SpecialNavigationStationTimeSliceType `xml:"SpecialNavigationStationTimeSlice,omitempty"`
}

// SpecialNavigationStationTimeSliceType Named source named complex type "SpecialNavigationStationTimeSliceType"
type SpecialNavigationStationTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "type" of complex type CodeSpecialNavigationStationType
	Type_XSD *aixmdatatypes.CodeSpecialNavigationStationType `xml:"type,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "emission" of complex type CodeRadioEmissionType
	Emission *aixmdatatypes.CodeRadioEmissionType `xml:"emission,omitempty"`

	// generated from element "systemChain" of complex type SpecialNavigationSystemPropertyType
	SystemChain *SpecialNavigationSystemPropertyType `xml:"systemChain,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AuthorityForSpecialNavigationStationPropertyType
	ResponsibleOrganisation *AuthorityForSpecialNavigationStationPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "position" of complex type ElevatedPointPropertyType
	Position *ElevatedPointPropertyType `xml:"position,omitempty"`

	// generated from element "availability" of complex type SpecialNavigationStationStatusPropertyType
	Availability []*SpecialNavigationStationStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_19 `xml:"extension,omitempty"`
}

// A_extension_19 Named source within outer element "extension"
type A_extension_19 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSpecialNavigationStationExtension" of complex type AbstractExtensionType
	AbstractSpecialNavigationStationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSpecialNavigationStationExtension,omitempty"`
}

// Group_SpecialNavigationStationPropertyGroup UnNamed source named group "SpecialNavigationStationPropertyGroup"
type Group_SpecialNavigationStationPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "type" of complex type CodeSpecialNavigationStationType
	Type_XSD *aixmdatatypes.CodeSpecialNavigationStationType `xml:"type,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "emission" of complex type CodeRadioEmissionType
	Emission *aixmdatatypes.CodeRadioEmissionType `xml:"emission,omitempty"`

	// generated from element "systemChain" of complex type SpecialNavigationSystemPropertyType
	SystemChain *SpecialNavigationSystemPropertyType `xml:"systemChain,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AuthorityForSpecialNavigationStationPropertyType
	ResponsibleOrganisation *AuthorityForSpecialNavigationStationPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "position" of complex type ElevatedPointPropertyType
	Position *ElevatedPointPropertyType `xml:"position,omitempty"`

	// generated from element "availability" of complex type SpecialNavigationStationStatusPropertyType
	Availability []*SpecialNavigationStationStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SpecialNavigationStationStatusType Named source named complex type "SpecialNavigationStationStatusType"
type SpecialNavigationStationStatusType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 3504 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:SpecialNavigationStationStatusPropertyGroup" with order 3505 depth 2
	Group_SpecialNavigationStationStatusPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_20 `xml:"extension,omitempty"`
}

// A_extension_20 Named source within outer element "extension"
type A_extension_20 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPropertiesWithScheduleExtension" of complex type AbstractExtensionType
	AbstractPropertiesWithScheduleExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPropertiesWithScheduleExtension,omitempty"`

	// generated from element "AbstractSpecialNavigationStationStatusExtension" of complex type AbstractExtensionType
	AbstractSpecialNavigationStationStatusExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSpecialNavigationStationStatusExtension,omitempty"`
}

// Group_SpecialNavigationStationStatusPropertyGroup UnNamed source named group "SpecialNavigationStationStatusPropertyGroup"
type Group_SpecialNavigationStationStatusPropertyGroup struct {

	// insertion point for fields

	// generated from element "operationalStatus" of complex type CodeStatusNavaidType
	OperationalStatus *aixmdatatypes.CodeStatusNavaidType `xml:"operationalStatus,omitempty"`
}

// SpecialNavigationStationStatusPropertyType Named source named complex type "SpecialNavigationStationStatusPropertyType"
type SpecialNavigationStationStatusPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SpecialNavigationStationStatus" of complex type SpecialNavigationStationStatusType
	SpecialNavigationStationStatus *SpecialNavigationStationStatusType `xml:"SpecialNavigationStationStatus,omitempty"`
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

// SpecialNavigationSystemType Named source named complex type "SpecialNavigationSystemType"
type SpecialNavigationSystemType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type SpecialNavigationSystemTimeSlicePropertyType
	TimeSlice []*SpecialNavigationSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SpecialNavigationSystemTimeSlicePropertyType Named source named complex type "SpecialNavigationSystemTimeSlicePropertyType"
type SpecialNavigationSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SpecialNavigationSystemTimeSlice" of complex type SpecialNavigationSystemTimeSliceType
	SpecialNavigationSystemTimeSlice *SpecialNavigationSystemTimeSliceType `xml:"SpecialNavigationSystemTimeSlice,omitempty"`
}

// SpecialNavigationSystemTimeSliceType Named source named complex type "SpecialNavigationSystemTimeSliceType"
type SpecialNavigationSystemTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "type" of complex type CodeSpecialNavigationSystemType
	Type_XSD *aixmdatatypes.CodeSpecialNavigationSystemType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeSpecialNavigationChainDesignatorType
	Designator *aixmdatatypes.CodeSpecialNavigationChainDesignatorType `xml:"designator,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AuthorityForSpecialNavigationSystemPropertyType
	ResponsibleOrganisation *AuthorityForSpecialNavigationSystemPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_21 `xml:"extension,omitempty"`
}

// A_extension_21 Named source within outer element "extension"
type A_extension_21 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSpecialNavigationSystemExtension" of complex type AbstractExtensionType
	AbstractSpecialNavigationSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSpecialNavigationSystemExtension,omitempty"`
}

// Group_SpecialNavigationSystemPropertyGroup UnNamed source named group "SpecialNavigationSystemPropertyGroup"
type Group_SpecialNavigationSystemPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeSpecialNavigationSystemType
	Type_XSD *aixmdatatypes.CodeSpecialNavigationSystemType `xml:"type,omitempty"`

	// generated from element "designator" of complex type CodeSpecialNavigationChainDesignatorType
	Designator *aixmdatatypes.CodeSpecialNavigationChainDesignatorType `xml:"designator,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "responsibleOrganisation" of complex type AuthorityForSpecialNavigationSystemPropertyType
	ResponsibleOrganisation *AuthorityForSpecialNavigationSystemPropertyType `xml:"responsibleOrganisation,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// TACANType Named source named complex type "TACANType"
type TACANType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type TACANTimeSlicePropertyType
	TimeSlice []*TACANTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// TACANTimeSlicePropertyType Named source named complex type "TACANTimeSlicePropertyType"
type TACANTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "TACANTimeSlice" of complex type TACANTimeSliceType
	TACANTimeSlice *TACANTimeSliceType `xml:"TACANTimeSlice,omitempty"`
}

// TACANTimeSliceType Named source named complex type "TACANTimeSliceType"
type TACANTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "channel" of complex type CodeTACANChannelType
	Channel *aixmdatatypes.CodeTACANChannelType `xml:"channel,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_22 `xml:"extension,omitempty"`
}

// A_extension_22 Named source within outer element "extension"
type A_extension_22 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractTACANExtension" of complex type AbstractExtensionType
	AbstractTACANExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractTACANExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_TACANPropertyGroup UnNamed source named group "TACANPropertyGroup"
type Group_TACANPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3570 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "channel" of complex type CodeTACANChannelType
	Channel *aixmdatatypes.CodeTACANChannelType `xml:"channel,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`
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

// VORType Named source named complex type "VORType"
type VORType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type VORTimeSlicePropertyType
	TimeSlice []*VORTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// VORTimeSlicePropertyType Named source named complex type "VORTimeSlicePropertyType"
type VORTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "VORTimeSlice" of complex type VORTimeSliceType
	VORTimeSlice *VORTimeSliceType `xml:"VORTimeSlice,omitempty"`
}

// VORTimeSliceType Named source named complex type "VORTimeSliceType"
type VORTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeNavaidDesignatorType
	Designator *aixmdatatypes.CodeNavaidDesignatorType `xml:"designator,omitempty"`

	// generated from element "emissionClass" of complex type CodeRadioEmissionType
	EmissionClass *aixmdatatypes.CodeRadioEmissionType `xml:"emissionClass,omitempty"`

	// generated from element "mobile" of complex type CodeYesNoType
	Mobile *aixmdatatypes.CodeYesNoType `xml:"mobile,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "flightChecked" of complex type CodeYesNoType
	FlightChecked *aixmdatatypes.CodeYesNoType `xml:"flightChecked,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "authority" of complex type AuthorityForNavaidEquipmentPropertyType
	Authority []*AuthorityForNavaidEquipmentPropertyType `xml:"authority,omitempty"`

	// generated from element "monitoring" of complex type NavaidEquipmentMonitoringPropertyType
	Monitoring []*NavaidEquipmentMonitoringPropertyType `xml:"monitoring,omitempty"`

	// generated from element "availability" of complex type NavaidOperationalStatusPropertyType
	Availability []*NavaidOperationalStatusPropertyType `xml:"availability,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "type" of complex type CodeVORType
	Type_XSD *aixmdatatypes.CodeVORType `xml:"type,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "zeroBearingDirection" of complex type CodeNorthReferenceType
	ZeroBearingDirection *aixmdatatypes.CodeNorthReferenceType `xml:"zeroBearingDirection,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_23 `xml:"extension,omitempty"`
}

// A_extension_23 Named source within outer element "extension"
type A_extension_23 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractVORExtension" of complex type AbstractExtensionType
	AbstractVORExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractVORExtension,omitempty"`

	// generated from element "AbstractNavaidEquipmentExtension" of complex type AbstractExtensionType
	AbstractNavaidEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractNavaidEquipmentExtension,omitempty"`
}

// Group_VORPropertyGroup UnNamed source named group "VORPropertyGroup"
type Group_VORPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:NavaidEquipmentPropertyGroup" with order 3597 depth 2
	Group_NavaidEquipmentPropertyGroup

	// generated from element "type" of complex type CodeVORType
	Type_XSD *aixmdatatypes.CodeVORType `xml:"type,omitempty"`

	// generated from element "frequency" of complex type ValFrequencyType
	Frequency *aixmdatatypes.ValFrequencyType `xml:"frequency,omitempty"`

	// generated from element "zeroBearingDirection" of complex type CodeNorthReferenceType
	ZeroBearingDirection *aixmdatatypes.CodeNorthReferenceType `xml:"zeroBearingDirection,omitempty"`

	// generated from element "declination" of complex type ValMagneticVariationType
	Declination *aixmdatatypes.ValMagneticVariationType `xml:"declination,omitempty"`
}

// AngleIndicationType Named source named complex type "AngleIndicationType"
type AngleIndicationType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type AngleIndicationTimeSlicePropertyType
	TimeSlice []*AngleIndicationTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// AngleIndicationTimeSlicePropertyType Named source named complex type "AngleIndicationTimeSlicePropertyType"
type AngleIndicationTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AngleIndicationTimeSlice" of complex type AngleIndicationTimeSliceType
	AngleIndicationTimeSlice *AngleIndicationTimeSliceType `xml:"AngleIndicationTimeSlice,omitempty"`
}

// AngleIndicationTimeSliceType Named source named complex type "AngleIndicationTimeSliceType"
type AngleIndicationTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:AngleIndicationPropertyGroup" with order 3617 depth 2
	Group_AngleIndicationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_24 `xml:"extension,omitempty"`
}

// A_extension_24 Named source within outer element "extension"
type A_extension_24 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractAngleIndicationExtension" of complex type AbstractExtensionType
	AbstractAngleIndicationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractAngleIndicationExtension,omitempty"`
}

// Group_AngleIndicationPropertyGroup UnNamed source named group "AngleIndicationPropertyGroup"
type Group_AngleIndicationPropertyGroup struct {

	// insertion point for fields

	// generated from element "angle" of complex type ValBearingType
	Angle *aixmdatatypes.ValBearingType `xml:"angle,omitempty"`

	// generated from element "angleType" of complex type CodeBearingType
	AngleType *aixmdatatypes.CodeBearingType `xml:"angleType,omitempty"`

	// generated from element "indicationDirection" of complex type CodeDirectionReferenceType
	IndicationDirection *aixmdatatypes.CodeDirectionReferenceType `xml:"indicationDirection,omitempty"`

	// generated from element "trueAngle" of complex type ValBearingType
	TrueAngle *aixmdatatypes.ValBearingType `xml:"trueAngle,omitempty"`

	// generated from element "cardinalDirection" of complex type CodeCardinalDirectionType
	CardinalDirection *aixmdatatypes.CodeCardinalDirectionType `xml:"cardinalDirection,omitempty"`

	// generated from element "minimumReceptionAltitude" of complex type ValDistanceVerticalType
	MinimumReceptionAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumReceptionAltitude,omitempty"`

	// generated from element "fix" of complex type DesignatedPointPropertyType
	Fix *DesignatedPointPropertyType `xml:"fix,omitempty"`

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

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// DistanceIndicationType Named source named complex type "DistanceIndicationType"
type DistanceIndicationType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type DistanceIndicationTimeSlicePropertyType
	TimeSlice []*DistanceIndicationTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DistanceIndicationTimeSlicePropertyType Named source named complex type "DistanceIndicationTimeSlicePropertyType"
type DistanceIndicationTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DistanceIndicationTimeSlice" of complex type DistanceIndicationTimeSliceType
	DistanceIndicationTimeSlice *DistanceIndicationTimeSliceType `xml:"DistanceIndicationTimeSlice,omitempty"`
}

// DistanceIndicationTimeSliceType Named source named complex type "DistanceIndicationTimeSliceType"
type DistanceIndicationTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:DistanceIndicationPropertyGroup" with order 3673 depth 2
	Group_DistanceIndicationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_25 `xml:"extension,omitempty"`
}

// A_extension_25 Named source within outer element "extension"
type A_extension_25 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDistanceIndicationExtension" of complex type AbstractExtensionType
	AbstractDistanceIndicationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDistanceIndicationExtension,omitempty"`
}

// Group_DistanceIndicationPropertyGroup UnNamed source named group "DistanceIndicationPropertyGroup"
type Group_DistanceIndicationPropertyGroup struct {

	// insertion point for fields

	// generated from element "distance" of complex type ValDistanceType
	Distance *aixmdatatypes.ValDistanceType `xml:"distance,omitempty"`

	// generated from element "minimumReceptionAltitude" of complex type ValDistanceVerticalType
	MinimumReceptionAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"minimumReceptionAltitude,omitempty"`

	// generated from element "type" of complex type CodeDistanceIndicationType
	Type_XSD *aixmdatatypes.CodeDistanceIndicationType `xml:"type,omitempty"`

	// generated from element "fix" of complex type DesignatedPointPropertyType
	Fix *DesignatedPointPropertyType `xml:"fix,omitempty"`

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

// DesignatedPointType Named source named complex type "DesignatedPointType"
type DesignatedPointType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type DesignatedPointTimeSlicePropertyType
	TimeSlice []*DesignatedPointTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// DesignatedPointTimeSlicePropertyType Named source named complex type "DesignatedPointTimeSlicePropertyType"
type DesignatedPointTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "DesignatedPointTimeSlice" of complex type DesignatedPointTimeSliceType
	DesignatedPointTimeSlice *DesignatedPointTimeSliceType `xml:"DesignatedPointTimeSlice,omitempty"`
}

// DesignatedPointTimeSliceType Named source named complex type "DesignatedPointTimeSliceType"
type DesignatedPointTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "designator" of complex type CodeDesignatedPointDesignatorType
	Designator *aixmdatatypes.CodeDesignatedPointDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeDesignatedPointType
	Type_XSD *aixmdatatypes.CodeDesignatedPointType `xml:"type,omitempty"`

	// generated from element "location" of complex type PointPropertyType
	Location *PointPropertyType `xml:"location,omitempty"`

	// generated from element "aimingPoint" of complex type TouchDownLiftOffPropertyType
	AimingPoint *TouchDownLiftOffPropertyType `xml:"aimingPoint,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport *AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "runwayPoint" of complex type RunwayCentrelinePointPropertyType
	RunwayPoint *RunwayCentrelinePointPropertyType `xml:"runwayPoint,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_26 `xml:"extension,omitempty"`
}

// A_extension_26 Named source within outer element "extension"
type A_extension_26 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractDesignatedPointExtension" of complex type AbstractExtensionType
	AbstractDesignatedPointExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractDesignatedPointExtension,omitempty"`
}

// Group_DesignatedPointPropertyGroup UnNamed source named group "DesignatedPointPropertyGroup"
type Group_DesignatedPointPropertyGroup struct {

	// insertion point for fields

	// generated from element "designator" of complex type CodeDesignatedPointDesignatorType
	Designator *aixmdatatypes.CodeDesignatedPointDesignatorType `xml:"designator,omitempty"`

	// generated from element "type" of complex type CodeDesignatedPointType
	Type_XSD *aixmdatatypes.CodeDesignatedPointType `xml:"type,omitempty"`

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "location" of complex type PointPropertyType
	Location *PointPropertyType `xml:"location,omitempty"`

	// generated from element "aimingPoint" of complex type TouchDownLiftOffPropertyType
	AimingPoint *TouchDownLiftOffPropertyType `xml:"aimingPoint,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport *AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "runwayPoint" of complex type RunwayCentrelinePointPropertyType
	RunwayPoint *RunwayCentrelinePointPropertyType `xml:"runwayPoint,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// SignificantPointInAirspaceType Named source named complex type "SignificantPointInAirspaceType"
type SignificantPointInAirspaceType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type SignificantPointInAirspaceTimeSlicePropertyType
	TimeSlice []*SignificantPointInAirspaceTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SignificantPointInAirspaceTimeSlicePropertyType Named source named complex type "SignificantPointInAirspaceTimeSlicePropertyType"
type SignificantPointInAirspaceTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SignificantPointInAirspaceTimeSlice" of complex type SignificantPointInAirspaceTimeSliceType
	SignificantPointInAirspaceTimeSlice *SignificantPointInAirspaceTimeSliceType `xml:"SignificantPointInAirspaceTimeSlice,omitempty"`
}

// SignificantPointInAirspaceTimeSliceType Named source named complex type "SignificantPointInAirspaceTimeSliceType"
type SignificantPointInAirspaceTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:SignificantPointInAirspacePropertyGroup" with order 3827 depth 2
	Group_SignificantPointInAirspacePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_27 `xml:"extension,omitempty"`
}

// A_extension_27 Named source within outer element "extension"
type A_extension_27 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSignificantPointInAirspaceExtension" of complex type AbstractExtensionType
	AbstractSignificantPointInAirspaceExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSignificantPointInAirspaceExtension,omitempty"`
}

// Group_SignificantPointInAirspacePropertyGroup UnNamed source named group "SignificantPointInAirspacePropertyGroup"
type Group_SignificantPointInAirspacePropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeAirspacePointRoleType
	Type_XSD *aixmdatatypes.CodeAirspacePointRoleType `xml:"type,omitempty"`

	// generated from element "relativeLocation" of complex type CodeAirspacePointPositionType
	RelativeLocation *aixmdatatypes.CodeAirspacePointPositionType `xml:"relativeLocation,omitempty"`

	// generated from element "containingAirspace" of complex type AirspacePropertyType
	ContainingAirspace *AirspacePropertyType `xml:"containingAirspace,omitempty"`

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

	// generated from group "aixm:LinguisticNotePropertyGroup" with order 3881 depth 2
	Group_LinguisticNotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_28 `xml:"extension,omitempty"`
}

// A_extension_28 Named source within outer element "extension"
type A_extension_28 struct {
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

	// generated from group "aixm:NotePropertyGroup" with order 3897 depth 2
	Group_NotePropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_29 `xml:"extension,omitempty"`
}

// A_extension_29 Named source within outer element "extension"
type A_extension_29 struct {
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

// UnitPropertyType Named source named complex type "UnitPropertyType"
type UnitPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// InformationServicePropertyType Named source named complex type "InformationServicePropertyType"
type InformationServicePropertyType struct {
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6232 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:OnlineContactPropertyGroup" with order 6233 depth 2
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6253 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:PostalAddressPropertyGroup" with order 6254 depth 2
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

	// generated from group "aixm:PropertiesWithSchedulePropertyGroup" with order 6275 depth 2
	Group_PropertiesWithSchedulePropertyGroup

	// generated from group "aixm:TelephoneContactPropertyGroup" with order 6276 depth 2
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

	// generated from group "aixm:TimesheetPropertyGroup" with order 6539 depth 2
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

// PrecisionApproachRadarType Named source named complex type "PrecisionApproachRadarType"
type PrecisionApproachRadarType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type PrecisionApproachRadarTimeSlicePropertyType
	TimeSlice []*PrecisionApproachRadarTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// PrecisionApproachRadarTimeSlicePropertyType Named source named complex type "PrecisionApproachRadarTimeSlicePropertyType"
type PrecisionApproachRadarTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "PrecisionApproachRadarTimeSlice" of complex type PrecisionApproachRadarTimeSliceType
	PrecisionApproachRadarTimeSlice *PrecisionApproachRadarTimeSliceType `xml:"PrecisionApproachRadarTimeSlice,omitempty"`
}

// PrecisionApproachRadarTimeSliceType Named source named complex type "PrecisionApproachRadarTimeSliceType"
type PrecisionApproachRadarTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "serialNumber" of complex type TextDesignatorType
	SerialNumber *aixmdatatypes.TextDesignatorType `xml:"serialNumber,omitempty"`

	// generated from element "range" of complex type ValDistanceType
	Range_XSD *aixmdatatypes.ValDistanceType `xml:"range,omitempty"`

	// generated from element "rangeAccuracy" of complex type ValDistanceType
	RangeAccuracy *aixmdatatypes.ValDistanceType `xml:"rangeAccuracy,omitempty"`

	// generated from element "dualChannel" of complex type CodeYesNoType
	DualChannel *aixmdatatypes.CodeYesNoType `xml:"dualChannel,omitempty"`

	// generated from element "movingTargetIndicator" of complex type CodeYesNoType
	MovingTargetIndicator *aixmdatatypes.CodeYesNoType `xml:"movingTargetIndicator,omitempty"`

	// generated from element "standbyPower" of complex type CodeStandbyPowerType
	StandbyPower *aixmdatatypes.CodeStandbyPowerType `xml:"standbyPower,omitempty"`

	// generated from element "digital" of complex type CodeYesNoType
	Digital *aixmdatatypes.CodeYesNoType `xml:"digital,omitempty"`

	// generated from element "militaryUseOnly" of complex type CodeYesNoType
	MilitaryUseOnly *aixmdatatypes.CodeYesNoType `xml:"militaryUseOnly,omitempty"`

	// generated from element "specialUseOnly" of complex type CodeYesNoType
	SpecialUseOnly *aixmdatatypes.CodeYesNoType `xml:"specialUseOnly,omitempty"`

	// generated from element "specialAircraftOnly" of complex type CodeYesNoType
	SpecialAircraftOnly *aixmdatatypes.CodeYesNoType `xml:"specialAircraftOnly,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "precisionApproachRadarType" of complex type CodePARType
	PrecisionApproachRadarType *aixmdatatypes.CodePARType `xml:"precisionApproachRadarType,omitempty"`

	// generated from element "slope" of complex type ValAngleType
	Slope *aixmdatatypes.ValAngleType `xml:"slope,omitempty"`

	// generated from element "slopeAccuracy" of complex type ValAngleType
	SlopeAccuracy *aixmdatatypes.ValAngleType `xml:"slopeAccuracy,omitempty"`

	// generated from element "reflector" of complex type ReflectorPropertyType
	Reflector []*ReflectorPropertyType `xml:"reflector,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_35 `xml:"extension,omitempty"`
}

// A_extension_35 Named source within outer element "extension"
type A_extension_35 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPrecisionApproachRadarExtension" of complex type AbstractExtensionType
	AbstractPrecisionApproachRadarExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPrecisionApproachRadarExtension,omitempty"`

	// generated from element "AbstractRadarEquipmentExtension" of complex type AbstractExtensionType
	AbstractRadarEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadarEquipmentExtension,omitempty"`
}

// Group_PrecisionApproachRadarPropertyGroup UnNamed source named group "PrecisionApproachRadarPropertyGroup"
type Group_PrecisionApproachRadarPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:RadarEquipmentPropertyGroup" with order 6822 depth 2
	Group_RadarEquipmentPropertyGroup

	// generated from element "precisionApproachRadarType" of complex type CodePARType
	PrecisionApproachRadarType *aixmdatatypes.CodePARType `xml:"precisionApproachRadarType,omitempty"`

	// generated from element "slope" of complex type ValAngleType
	Slope *aixmdatatypes.ValAngleType `xml:"slope,omitempty"`

	// generated from element "slopeAccuracy" of complex type ValAngleType
	SlopeAccuracy *aixmdatatypes.ValAngleType `xml:"slopeAccuracy,omitempty"`

	// generated from element "reflector" of complex type ReflectorPropertyType
	Reflector []*ReflectorPropertyType `xml:"reflector,omitempty"`
}

// PrimarySurveillanceRadarType Named source named complex type "PrimarySurveillanceRadarType"
type PrimarySurveillanceRadarType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type PrimarySurveillanceRadarTimeSlicePropertyType
	TimeSlice []*PrimarySurveillanceRadarTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// PrimarySurveillanceRadarTimeSlicePropertyType Named source named complex type "PrimarySurveillanceRadarTimeSlicePropertyType"
type PrimarySurveillanceRadarTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "PrimarySurveillanceRadarTimeSlice" of complex type PrimarySurveillanceRadarTimeSliceType
	PrimarySurveillanceRadarTimeSlice *PrimarySurveillanceRadarTimeSliceType `xml:"PrimarySurveillanceRadarTimeSlice,omitempty"`
}

// PrimarySurveillanceRadarTimeSliceType Named source named complex type "PrimarySurveillanceRadarTimeSliceType"
type PrimarySurveillanceRadarTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "serialNumber" of complex type TextDesignatorType
	SerialNumber *aixmdatatypes.TextDesignatorType `xml:"serialNumber,omitempty"`

	// generated from element "range" of complex type ValDistanceType
	Range_XSD *aixmdatatypes.ValDistanceType `xml:"range,omitempty"`

	// generated from element "rangeAccuracy" of complex type ValDistanceType
	RangeAccuracy *aixmdatatypes.ValDistanceType `xml:"rangeAccuracy,omitempty"`

	// generated from element "dualChannel" of complex type CodeYesNoType
	DualChannel *aixmdatatypes.CodeYesNoType `xml:"dualChannel,omitempty"`

	// generated from element "movingTargetIndicator" of complex type CodeYesNoType
	MovingTargetIndicator *aixmdatatypes.CodeYesNoType `xml:"movingTargetIndicator,omitempty"`

	// generated from element "standbyPower" of complex type CodeStandbyPowerType
	StandbyPower *aixmdatatypes.CodeStandbyPowerType `xml:"standbyPower,omitempty"`

	// generated from element "digital" of complex type CodeYesNoType
	Digital *aixmdatatypes.CodeYesNoType `xml:"digital,omitempty"`

	// generated from element "militaryUseOnly" of complex type CodeYesNoType
	MilitaryUseOnly *aixmdatatypes.CodeYesNoType `xml:"militaryUseOnly,omitempty"`

	// generated from element "specialUseOnly" of complex type CodeYesNoType
	SpecialUseOnly *aixmdatatypes.CodeYesNoType `xml:"specialUseOnly,omitempty"`

	// generated from element "specialAircraftOnly" of complex type CodeYesNoType
	SpecialAircraftOnly *aixmdatatypes.CodeYesNoType `xml:"specialAircraftOnly,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "verticalCoverageAltitude" of complex type ValDistanceVerticalType
	VerticalCoverageAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"verticalCoverageAltitude,omitempty"`

	// generated from element "verticalCoverageDistance" of complex type ValDistanceType
	VerticalCoverageDistance *aixmdatatypes.ValDistanceType `xml:"verticalCoverageDistance,omitempty"`

	// generated from element "verticalCoverageAzimuth" of complex type ValBearingType
	VerticalCoverageAzimuth *aixmdatatypes.ValBearingType `xml:"verticalCoverageAzimuth,omitempty"`

	// generated from element "antennaTiltFixed" of complex type CodeYesNoType
	AntennaTiltFixed *aixmdatatypes.CodeYesNoType `xml:"antennaTiltFixed,omitempty"`

	// generated from element "tiltAngle" of complex type ValAngleType
	TiltAngle *aixmdatatypes.ValAngleType `xml:"tiltAngle,omitempty"`

	// generated from element "automatedRadarTerminalSystem" of complex type TextDesignatorType
	AutomatedRadarTerminalSystem *aixmdatatypes.TextDesignatorType `xml:"automatedRadarTerminalSystem,omitempty"`

	// generated from element "groundStation" of complex type SurveillanceGroundStationPropertyType
	GroundStation []*SurveillanceGroundStationPropertyType `xml:"groundStation,omitempty"`

	// generated from element "type" of complex type CodePrimaryRadarType
	Type_XSD *aixmdatatypes.CodePrimaryRadarType `xml:"type,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_36 `xml:"extension,omitempty"`
}

// A_extension_36 Named source within outer element "extension"
type A_extension_36 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractPrimarySurveillanceRadarExtension" of complex type AbstractExtensionType
	AbstractPrimarySurveillanceRadarExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractPrimarySurveillanceRadarExtension,omitempty"`

	// generated from element "AbstractSurveillanceRadarExtension" of complex type AbstractExtensionType
	AbstractSurveillanceRadarExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurveillanceRadarExtension,omitempty"`

	// generated from element "AbstractRadarEquipmentExtension" of complex type AbstractExtensionType
	AbstractRadarEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadarEquipmentExtension,omitempty"`
}

// Group_PrimarySurveillanceRadarPropertyGroup UnNamed source named group "PrimarySurveillanceRadarPropertyGroup"
type Group_PrimarySurveillanceRadarPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:SurveillanceRadarPropertyGroup" with order 6852 depth 2
	Group_SurveillanceRadarPropertyGroup

	// generated from element "type" of complex type CodePrimaryRadarType
	Type_XSD *aixmdatatypes.CodePrimaryRadarType `xml:"type,omitempty"`
}

// RadarComponentType Named source named complex type "RadarComponentType"
type RadarComponentType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:RadarComponentPropertyGroup" with order 6858 depth 2
	Group_RadarComponentPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_37 `xml:"extension,omitempty"`
}

// A_extension_37 Named source within outer element "extension"
type A_extension_37 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRadarComponentExtension" of complex type AbstractExtensionType
	AbstractRadarComponentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadarComponentExtension,omitempty"`
}

// Group_RadarComponentPropertyGroup UnNamed source named group "RadarComponentPropertyGroup"
type Group_RadarComponentPropertyGroup struct {

	// insertion point for fields

	// generated from element "collocationGroup" of complex type NoSequenceType
	CollocationGroup *aixmdatatypes.NoSequenceType `xml:"collocationGroup,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theRadarEquipment" of complex type RadarEquipmentPropertyType
	TheRadarEquipment *RadarEquipmentPropertyType `xml:"theRadarEquipment,omitempty"`
}

// RadarComponentPropertyType Named source named complex type "RadarComponentPropertyType"
type RadarComponentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "RadarComponent" of complex type RadarComponentType
	RadarComponent *RadarComponentType `xml:"RadarComponent,omitempty"`
}

// RadarEquipmentPropertyType Named source named complex type "RadarEquipmentPropertyType"
type RadarEquipmentPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from attribute group "gml:AssociationAttributeGroup"
	gml.AttributeGroup_AssociationAttributeGroup
}

// AbstractRadarEquipmentType Named source named complex type "AbstractRadarEquipmentType"
type AbstractRadarEquipmentType struct {
	Name string `xml:"-"`

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

// Group_RadarEquipmentPropertyGroup UnNamed source named group "RadarEquipmentPropertyGroup"
type Group_RadarEquipmentPropertyGroup struct {

	// insertion point for fields

	// generated from element "name" of complex type TextNameType
	NameXSD *aixmdatatypes.TextNameType `xml:"name,omitempty"`

	// generated from element "serialNumber" of complex type TextDesignatorType
	SerialNumber *aixmdatatypes.TextDesignatorType `xml:"serialNumber,omitempty"`

	// generated from element "range" of complex type ValDistanceType
	Range_XSD *aixmdatatypes.ValDistanceType `xml:"range,omitempty"`

	// generated from element "rangeAccuracy" of complex type ValDistanceType
	RangeAccuracy *aixmdatatypes.ValDistanceType `xml:"rangeAccuracy,omitempty"`

	// generated from element "dualChannel" of complex type CodeYesNoType
	DualChannel *aixmdatatypes.CodeYesNoType `xml:"dualChannel,omitempty"`

	// generated from element "movingTargetIndicator" of complex type CodeYesNoType
	MovingTargetIndicator *aixmdatatypes.CodeYesNoType `xml:"movingTargetIndicator,omitempty"`

	// generated from element "standbyPower" of complex type CodeStandbyPowerType
	StandbyPower *aixmdatatypes.CodeStandbyPowerType `xml:"standbyPower,omitempty"`

	// generated from element "digital" of complex type CodeYesNoType
	Digital *aixmdatatypes.CodeYesNoType `xml:"digital,omitempty"`

	// generated from element "militaryUseOnly" of complex type CodeYesNoType
	MilitaryUseOnly *aixmdatatypes.CodeYesNoType `xml:"militaryUseOnly,omitempty"`

	// generated from element "specialUseOnly" of complex type CodeYesNoType
	SpecialUseOnly *aixmdatatypes.CodeYesNoType `xml:"specialUseOnly,omitempty"`

	// generated from element "specialAircraftOnly" of complex type CodeYesNoType
	SpecialAircraftOnly *aixmdatatypes.CodeYesNoType `xml:"specialAircraftOnly,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// RadarSystemType Named source named complex type "RadarSystemType"
type RadarSystemType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type RadarSystemTimeSlicePropertyType
	TimeSlice []*RadarSystemTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// RadarSystemTimeSlicePropertyType Named source named complex type "RadarSystemTimeSlicePropertyType"
type RadarSystemTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "RadarSystemTimeSlice" of complex type RadarSystemTimeSliceType
	RadarSystemTimeSlice *RadarSystemTimeSliceType `xml:"RadarSystemTimeSlice,omitempty"`
}

// RadarSystemTimeSliceType Named source named complex type "RadarSystemTimeSliceType"
type RadarSystemTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from group "aixm:RadarSystemPropertyGroup" with order 6912 depth 2
	Group_RadarSystemPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_38 `xml:"extension,omitempty"`
}

// A_extension_38 Named source within outer element "extension"
type A_extension_38 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractRadarSystemExtension" of complex type AbstractExtensionType
	AbstractRadarSystemExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadarSystemExtension,omitempty"`
}

// Group_RadarSystemPropertyGroup UnNamed source named group "RadarSystemPropertyGroup"
type Group_RadarSystemPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeRadarServiceType
	Type_XSD *aixmdatatypes.CodeRadarServiceType `xml:"type,omitempty"`

	// generated from element "model" of complex type TextNameType
	Model *aixmdatatypes.TextNameType `xml:"model,omitempty"`

	// generated from element "generalTerrainMonitor" of complex type CodeYesNoType
	GeneralTerrainMonitor *aixmdatatypes.CodeYesNoType `xml:"generalTerrainMonitor,omitempty"`

	// generated from element "broadcastIdentifier" of complex type TextDesignatorType
	BroadcastIdentifier *aixmdatatypes.TextDesignatorType `xml:"broadcastIdentifier,omitempty"`

	// generated from element "radarEquipment" of complex type RadarComponentPropertyType
	RadarEquipment []*RadarComponentPropertyType `xml:"radarEquipment,omitempty"`

	// generated from element "office" of complex type OrganisationAuthorityPropertyType
	Office []*OrganisationAuthorityPropertyType `xml:"office,omitempty"`

	// generated from element "airportHeliport" of complex type AirportHeliportPropertyType
	AirportHeliport *AirportHeliportPropertyType `xml:"airportHeliport,omitempty"`

	// generated from element "PARRunway" of complex type RunwayPropertyType
	PARRunway []*RunwayPropertyType `xml:"PARRunway,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ReflectorType Named source named complex type "ReflectorType"
type ReflectorType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:ReflectorPropertyGroup" with order 6934 depth 2
	Group_ReflectorPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_39 `xml:"extension,omitempty"`
}

// A_extension_39 Named source within outer element "extension"
type A_extension_39 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractReflectorExtension" of complex type AbstractExtensionType
	AbstractReflectorExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractReflectorExtension,omitempty"`
}

// Group_ReflectorPropertyGroup UnNamed source named group "ReflectorPropertyGroup"
type Group_ReflectorPropertyGroup struct {

	// insertion point for fields

	// generated from element "type" of complex type CodeReflectorType
	Type_XSD *aixmdatatypes.CodeReflectorType `xml:"type,omitempty"`

	// generated from element "touchdownReflector" of complex type ElevatedPointPropertyType
	TouchdownReflector *ElevatedPointPropertyType `xml:"touchdownReflector,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`
}

// ReflectorPropertyType Named source named complex type "ReflectorPropertyType"
type ReflectorPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "Reflector" of complex type ReflectorType
	Reflector *ReflectorType `xml:"Reflector,omitempty"`
}

// SecondarySurveillanceRadarType Named source named complex type "SecondarySurveillanceRadarType"
type SecondarySurveillanceRadarType struct {
	Name string `xml:"-"`

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

	// generated from element "timeSlice" of complex type SecondarySurveillanceRadarTimeSlicePropertyType
	TimeSlice []*SecondarySurveillanceRadarTimeSlicePropertyType `xml:"timeSlice,omitempty"`
}

// SecondarySurveillanceRadarTimeSlicePropertyType Named source named complex type "SecondarySurveillanceRadarTimeSlicePropertyType"
type SecondarySurveillanceRadarTimeSlicePropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "SecondarySurveillanceRadarTimeSlice" of complex type SecondarySurveillanceRadarTimeSliceType
	SecondarySurveillanceRadarTimeSlice *SecondarySurveillanceRadarTimeSliceType `xml:"SecondarySurveillanceRadarTimeSlice,omitempty"`
}

// SecondarySurveillanceRadarTimeSliceType Named source named complex type "SecondarySurveillanceRadarTimeSliceType"
type SecondarySurveillanceRadarTimeSliceType struct {
	Name string `xml:"-"`

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

	// generated from element "serialNumber" of complex type TextDesignatorType
	SerialNumber *aixmdatatypes.TextDesignatorType `xml:"serialNumber,omitempty"`

	// generated from element "range" of complex type ValDistanceType
	Range_XSD *aixmdatatypes.ValDistanceType `xml:"range,omitempty"`

	// generated from element "rangeAccuracy" of complex type ValDistanceType
	RangeAccuracy *aixmdatatypes.ValDistanceType `xml:"rangeAccuracy,omitempty"`

	// generated from element "dualChannel" of complex type CodeYesNoType
	DualChannel *aixmdatatypes.CodeYesNoType `xml:"dualChannel,omitempty"`

	// generated from element "movingTargetIndicator" of complex type CodeYesNoType
	MovingTargetIndicator *aixmdatatypes.CodeYesNoType `xml:"movingTargetIndicator,omitempty"`

	// generated from element "standbyPower" of complex type CodeStandbyPowerType
	StandbyPower *aixmdatatypes.CodeStandbyPowerType `xml:"standbyPower,omitempty"`

	// generated from element "digital" of complex type CodeYesNoType
	Digital *aixmdatatypes.CodeYesNoType `xml:"digital,omitempty"`

	// generated from element "militaryUseOnly" of complex type CodeYesNoType
	MilitaryUseOnly *aixmdatatypes.CodeYesNoType `xml:"militaryUseOnly,omitempty"`

	// generated from element "specialUseOnly" of complex type CodeYesNoType
	SpecialUseOnly *aixmdatatypes.CodeYesNoType `xml:"specialUseOnly,omitempty"`

	// generated from element "specialAircraftOnly" of complex type CodeYesNoType
	SpecialAircraftOnly *aixmdatatypes.CodeYesNoType `xml:"specialAircraftOnly,omitempty"`

	// generated from element "magneticVariation" of complex type ValMagneticVariationType
	MagneticVariation *aixmdatatypes.ValMagneticVariationType `xml:"magneticVariation,omitempty"`

	// generated from element "magneticVariationAccuracy" of complex type ValAngleType
	MagneticVariationAccuracy *aixmdatatypes.ValAngleType `xml:"magneticVariationAccuracy,omitempty"`

	// generated from element "dateMagneticVariation" of complex type DateYearType
	DateMagneticVariation *aixmdatatypes.DateYearType `xml:"dateMagneticVariation,omitempty"`

	// generated from element "contact" of complex type ContactInformationPropertyType
	Contact []*ContactInformationPropertyType `xml:"contact,omitempty"`

	// generated from element "location" of complex type ElevatedPointPropertyType
	Location *ElevatedPointPropertyType `xml:"location,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "verticalCoverageAltitude" of complex type ValDistanceVerticalType
	VerticalCoverageAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"verticalCoverageAltitude,omitempty"`

	// generated from element "verticalCoverageDistance" of complex type ValDistanceType
	VerticalCoverageDistance *aixmdatatypes.ValDistanceType `xml:"verticalCoverageDistance,omitempty"`

	// generated from element "verticalCoverageAzimuth" of complex type ValBearingType
	VerticalCoverageAzimuth *aixmdatatypes.ValBearingType `xml:"verticalCoverageAzimuth,omitempty"`

	// generated from element "antennaTiltFixed" of complex type CodeYesNoType
	AntennaTiltFixed *aixmdatatypes.CodeYesNoType `xml:"antennaTiltFixed,omitempty"`

	// generated from element "tiltAngle" of complex type ValAngleType
	TiltAngle *aixmdatatypes.ValAngleType `xml:"tiltAngle,omitempty"`

	// generated from element "automatedRadarTerminalSystem" of complex type TextDesignatorType
	AutomatedRadarTerminalSystem *aixmdatatypes.TextDesignatorType `xml:"automatedRadarTerminalSystem,omitempty"`

	// generated from element "groundStation" of complex type SurveillanceGroundStationPropertyType
	GroundStation []*SurveillanceGroundStationPropertyType `xml:"groundStation,omitempty"`

	// generated from element "transponder" of complex type CodeTransponderType
	Transponder *aixmdatatypes.CodeTransponderType `xml:"transponder,omitempty"`

	// generated from element "autonomous" of complex type CodeYesNoType
	Autonomous *aixmdatatypes.CodeYesNoType `xml:"autonomous,omitempty"`

	// generated from element "monopulse" of complex type CodeYesNoType
	Monopulse *aixmdatatypes.CodeYesNoType `xml:"monopulse,omitempty"`

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_40 `xml:"extension,omitempty"`
}

// A_extension_40 Named source within outer element "extension"
type A_extension_40 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSecondarySurveillanceRadarExtension" of complex type AbstractExtensionType
	AbstractSecondarySurveillanceRadarExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSecondarySurveillanceRadarExtension,omitempty"`

	// generated from element "AbstractSurveillanceRadarExtension" of complex type AbstractExtensionType
	AbstractSurveillanceRadarExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurveillanceRadarExtension,omitempty"`

	// generated from element "AbstractRadarEquipmentExtension" of complex type AbstractExtensionType
	AbstractRadarEquipmentExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractRadarEquipmentExtension,omitempty"`
}

// Group_SecondarySurveillanceRadarPropertyGroup UnNamed source named group "SecondarySurveillanceRadarPropertyGroup"
type Group_SecondarySurveillanceRadarPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:SurveillanceRadarPropertyGroup" with order 6973 depth 2
	Group_SurveillanceRadarPropertyGroup

	// generated from element "transponder" of complex type CodeTransponderType
	Transponder *aixmdatatypes.CodeTransponderType `xml:"transponder,omitempty"`

	// generated from element "autonomous" of complex type CodeYesNoType
	Autonomous *aixmdatatypes.CodeYesNoType `xml:"autonomous,omitempty"`

	// generated from element "monopulse" of complex type CodeYesNoType
	Monopulse *aixmdatatypes.CodeYesNoType `xml:"monopulse,omitempty"`
}

// SurveillanceGroundStationType Named source named complex type "SurveillanceGroundStationType"
type SurveillanceGroundStationType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gml:id"
	gml.Attr_id

	// generated from element "description" of simple type string
	Description string `xml:"description,omitempty"`

	// generated from element "identifier" of complex type CodeWithAuthorityType
	Identifier *gml.CodeWithAuthorityType `xml:"identifier,omitempty"`

	// generated from element "name" of complex type CodeType
	NameXSD []*gml.CodeType `xml:"name,omitempty"`

	// generated from group "aixm:SurveillanceGroundStationPropertyGroup" with order 6981 depth 2
	Group_SurveillanceGroundStationPropertyGroup

	// generated from anonymous complex type within element "extension"
	Extension []*A_extension_41 `xml:"extension,omitempty"`
}

// A_extension_41 Named source within outer element "extension"
type A_extension_41 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gml:OwnershipAttributeGroup"
	gml.AttributeGroup_OwnershipAttributeGroup

	// generated from element "AbstractSurveillanceGroundStationExtension" of complex type AbstractExtensionType
	AbstractSurveillanceGroundStationExtension []*aixmabstractgmlobjecttypes.AbstractExtensionType `xml:"AbstractSurveillanceGroundStationExtension,omitempty"`
}

// Group_SurveillanceGroundStationPropertyGroup UnNamed source named group "SurveillanceGroundStationPropertyGroup"
type Group_SurveillanceGroundStationPropertyGroup struct {

	// insertion point for fields

	// generated from element "videoMap" of complex type CodeYesNoType
	VideoMap *aixmdatatypes.CodeYesNoType `xml:"videoMap,omitempty"`

	// generated from element "annotation" of complex type NotePropertyType
	Annotation []*NotePropertyType `xml:"annotation,omitempty"`

	// generated from element "theUnit" of complex type UnitPropertyType
	TheUnit *UnitPropertyType `xml:"theUnit,omitempty"`
}

// SurveillanceGroundStationPropertyType Named source named complex type "SurveillanceGroundStationPropertyType"
type SurveillanceGroundStationPropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from element "SurveillanceGroundStation" of complex type SurveillanceGroundStationType
	SurveillanceGroundStation *SurveillanceGroundStationType `xml:"SurveillanceGroundStation,omitempty"`
}

// AbstractSurveillanceRadarType Named source named complex type "AbstractSurveillanceRadarType"
type AbstractSurveillanceRadarType struct {
	Name string `xml:"-"`

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

// Group_SurveillanceRadarPropertyGroup UnNamed source named group "SurveillanceRadarPropertyGroup"
type Group_SurveillanceRadarPropertyGroup struct {

	// insertion point for fields

	// generated from group "aixm:RadarEquipmentPropertyGroup" with order 7003 depth 2
	Group_RadarEquipmentPropertyGroup

	// generated from element "verticalCoverageAltitude" of complex type ValDistanceVerticalType
	VerticalCoverageAltitude *aixmdatatypes.ValDistanceVerticalType `xml:"verticalCoverageAltitude,omitempty"`

	// generated from element "verticalCoverageDistance" of complex type ValDistanceType
	VerticalCoverageDistance *aixmdatatypes.ValDistanceType `xml:"verticalCoverageDistance,omitempty"`

	// generated from element "verticalCoverageAzimuth" of complex type ValBearingType
	VerticalCoverageAzimuth *aixmdatatypes.ValBearingType `xml:"verticalCoverageAzimuth,omitempty"`

	// generated from element "antennaTiltFixed" of complex type CodeYesNoType
	AntennaTiltFixed *aixmdatatypes.CodeYesNoType `xml:"antennaTiltFixed,omitempty"`

	// generated from element "tiltAngle" of complex type ValAngleType
	TiltAngle *aixmdatatypes.ValAngleType `xml:"tiltAngle,omitempty"`

	// generated from element "automatedRadarTerminalSystem" of complex type TextDesignatorType
	AutomatedRadarTerminalSystem *aixmdatatypes.TextDesignatorType `xml:"automatedRadarTerminalSystem,omitempty"`

	// generated from element "groundStation" of complex type SurveillanceGroundStationPropertyType
	GroundStation []*SurveillanceGroundStationPropertyType `xml:"groundStation,omitempty"`
}
