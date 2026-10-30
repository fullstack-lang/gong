// generated code - do not edit
package gmd

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/applicationschema"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/portrayalcatalogue"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *applicationschema.MD_ApplicationSchemaInformation_Type
var _ *content.MD_FeatureCatalogueDescription_Type
var _ *gcobase.AbstractObject_Type
var _ *gml.DynamicFeatureType
var _ *metadataextension.MD_ExtendedElementInformation_Type
var _ *portrayalcatalogue.MD_PortrayalCatalogueReference_Type
var _ *spatialrepresentation.MD_GridSpatialRepresentation_Type

// insertion point for enums declaration

// insertion point for gongstructs declarations

// AbstractDS_Aggregate_Type Named source named complex type "AbstractDS_Aggregate_Type"
// Identifiable collection of datasets
type AbstractDS_Aggregate_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_Aggregate_PropertyType Named source named complex type "DS_Aggregate_PropertyType"
type DS_Aggregate_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_OtherAggregate" of complex type DS_OtherAggregate_Type
	DS_OtherAggregate *DS_OtherAggregate_Type `xml:"DS_OtherAggregate,omitempty"`

	// generated from element "DS_StereoMate" of complex type DS_StereoMate_Type
	DS_StereoMate *DS_StereoMate_Type `xml:"DS_StereoMate,omitempty"`

	// generated from element "DS_Series" of complex type DS_Series_Type
	DS_Series *DS_Series_Type `xml:"DS_Series,omitempty"`

	// generated from element "DS_Platform" of complex type DS_Platform_Type
	DS_Platform *DS_Platform_Type `xml:"DS_Platform,omitempty"`

	// generated from element "DS_Sensor" of complex type DS_Sensor_Type
	DS_Sensor *DS_Sensor_Type `xml:"DS_Sensor,omitempty"`

	// generated from element "DS_ProductionSeries" of complex type DS_ProductionSeries_Type
	DS_ProductionSeries *DS_ProductionSeries_Type `xml:"DS_ProductionSeries,omitempty"`

	// generated from element "DS_Initiative" of complex type DS_Initiative_Type
	DS_Initiative *DS_Initiative_Type `xml:"DS_Initiative,omitempty"`
}

// DS_DataSet_Type Named source named complex type "DS_DataSet_Type"
// Identifiable collection of data
type DS_DataSet_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "has" of complex type MD_Metadata_PropertyType
	Has []*MD_Metadata_PropertyType `xml:"has,omitempty"`

	// generated from element "partOf" of complex type DS_Aggregate_PropertyType
	PartOf []*DS_Aggregate_PropertyType `xml:"partOf,omitempty"`
}

// DS_DataSet_PropertyType Named source named complex type "DS_DataSet_PropertyType"
type DS_DataSet_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_DataSet" of complex type DS_DataSet_Type
	DS_DataSet *DS_DataSet_Type `xml:"DS_DataSet,omitempty"`
}

// DS_OtherAggregate_Type Named source named complex type "DS_OtherAggregate_Type"
type DS_OtherAggregate_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_OtherAggregate_PropertyType Named source named complex type "DS_OtherAggregate_PropertyType"
type DS_OtherAggregate_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_OtherAggregate" of complex type DS_OtherAggregate_Type
	DS_OtherAggregate *DS_OtherAggregate_Type `xml:"DS_OtherAggregate,omitempty"`

	// generated from element "DS_StereoMate" of complex type DS_StereoMate_Type
	DS_StereoMate *DS_StereoMate_Type `xml:"DS_StereoMate,omitempty"`
}

// DS_Series_Type Named source named complex type "DS_Series_Type"
type DS_Series_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_Series_PropertyType Named source named complex type "DS_Series_PropertyType"
type DS_Series_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_Series" of complex type DS_Series_Type
	DS_Series *DS_Series_Type `xml:"DS_Series,omitempty"`

	// generated from element "DS_Platform" of complex type DS_Platform_Type
	DS_Platform *DS_Platform_Type `xml:"DS_Platform,omitempty"`

	// generated from element "DS_Sensor" of complex type DS_Sensor_Type
	DS_Sensor *DS_Sensor_Type `xml:"DS_Sensor,omitempty"`

	// generated from element "DS_ProductionSeries" of complex type DS_ProductionSeries_Type
	DS_ProductionSeries *DS_ProductionSeries_Type `xml:"DS_ProductionSeries,omitempty"`
}

// DS_Initiative_Type Named source named complex type "DS_Initiative_Type"
type DS_Initiative_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_Initiative_PropertyType Named source named complex type "DS_Initiative_PropertyType"
type DS_Initiative_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_Initiative" of complex type DS_Initiative_Type
	DS_Initiative *DS_Initiative_Type `xml:"DS_Initiative,omitempty"`
}

// DS_Platform_Type Named source named complex type "DS_Platform_Type"
type DS_Platform_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_Platform_PropertyType Named source named complex type "DS_Platform_PropertyType"
type DS_Platform_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_Platform" of complex type DS_Platform_Type
	DS_Platform *DS_Platform_Type `xml:"DS_Platform,omitempty"`
}

// DS_Sensor_Type Named source named complex type "DS_Sensor_Type"
type DS_Sensor_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_Sensor_PropertyType Named source named complex type "DS_Sensor_PropertyType"
type DS_Sensor_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_Sensor" of complex type DS_Sensor_Type
	DS_Sensor *DS_Sensor_Type `xml:"DS_Sensor,omitempty"`
}

// DS_ProductionSeries_Type Named source named complex type "DS_ProductionSeries_Type"
type DS_ProductionSeries_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_ProductionSeries_PropertyType Named source named complex type "DS_ProductionSeries_PropertyType"
type DS_ProductionSeries_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_ProductionSeries" of complex type DS_ProductionSeries_Type
	DS_ProductionSeries *DS_ProductionSeries_Type `xml:"DS_ProductionSeries,omitempty"`
}

// DS_StereoMate_Type Named source named complex type "DS_StereoMate_Type"
type DS_StereoMate_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "composedOf" of complex type DS_DataSet_PropertyType
	ComposedOf []*DS_DataSet_PropertyType `xml:"composedOf,omitempty"`

	// generated from element "seriesMetadata" of complex type MD_Metadata_PropertyType
	SeriesMetadata []*MD_Metadata_PropertyType `xml:"seriesMetadata,omitempty"`

	// generated from element "subset" of complex type DS_Aggregate_PropertyType
	Subset []*DS_Aggregate_PropertyType `xml:"subset,omitempty"`

	// generated from element "superset" of complex type DS_Aggregate_PropertyType
	Superset []*DS_Aggregate_PropertyType `xml:"superset,omitempty"`
}

// DS_StereoMate_PropertyType Named source named complex type "DS_StereoMate_PropertyType"
type DS_StereoMate_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "DS_StereoMate" of complex type DS_StereoMate_Type
	DS_StereoMate *DS_StereoMate_Type `xml:"DS_StereoMate,omitempty"`
}

// MD_Metadata_Type Named source named complex type "MD_Metadata_Type"
// Information about the metadata
type MD_Metadata_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "fileIdentifier" of complex type CharacterString_PropertyType
	FileIdentifier *gml.CharacterString_PropertyType `xml:"fileIdentifier,omitempty"`

	// generated from element "language" of complex type CharacterString_PropertyType
	Language *gml.CharacterString_PropertyType `xml:"language,omitempty"`

	// generated from element "characterSet" of complex type MD_CharacterSetCode_PropertyType
	CharacterSet *gml.MD_CharacterSetCode_PropertyType `xml:"characterSet,omitempty"`

	// generated from element "parentIdentifier" of complex type CharacterString_PropertyType
	ParentIdentifier *gml.CharacterString_PropertyType `xml:"parentIdentifier,omitempty"`

	// generated from element "hierarchyLevel" of complex type MD_ScopeCode_PropertyType
	HierarchyLevel []*gml.MD_ScopeCode_PropertyType `xml:"hierarchyLevel,omitempty"`

	// generated from element "hierarchyLevelName" of complex type CharacterString_PropertyType
	HierarchyLevelName []*gml.CharacterString_PropertyType `xml:"hierarchyLevelName,omitempty"`

	// generated from element "contact" of complex type CI_ResponsibleParty_PropertyType
	Contact []*gml.CI_ResponsibleParty_PropertyType `xml:"contact,omitempty"`

	// generated from element "dateStamp" of complex type Date_PropertyType
	DateStamp *gml.Date_PropertyType `xml:"dateStamp,omitempty"`

	// generated from element "metadataStandardName" of complex type CharacterString_PropertyType
	MetadataStandardName *gml.CharacterString_PropertyType `xml:"metadataStandardName,omitempty"`

	// generated from element "metadataStandardVersion" of complex type CharacterString_PropertyType
	MetadataStandardVersion *gml.CharacterString_PropertyType `xml:"metadataStandardVersion,omitempty"`

	// generated from element "dataSetURI" of complex type CharacterString_PropertyType
	DataSetURI *gml.CharacterString_PropertyType `xml:"dataSetURI,omitempty"`

	// generated from element "locale" of complex type PT_Locale_PropertyType
	Locale []*gml.PT_Locale_PropertyType `xml:"locale,omitempty"`

	// generated from element "spatialRepresentationInfo" of complex type MD_SpatialRepresentation_PropertyType
	SpatialRepresentationInfo []*spatialrepresentation.MD_SpatialRepresentation_PropertyType `xml:"spatialRepresentationInfo,omitempty"`

	// generated from element "referenceSystemInfo" of complex type MD_ReferenceSystem_PropertyType
	ReferenceSystemInfo []*gml.MD_ReferenceSystem_PropertyType `xml:"referenceSystemInfo,omitempty"`

	// generated from element "metadataExtensionInfo" of complex type MD_MetadataExtensionInformation_PropertyType
	MetadataExtensionInfo []*metadataextension.MD_MetadataExtensionInformation_PropertyType `xml:"metadataExtensionInfo,omitempty"`

	// generated from element "identificationInfo" of complex type MD_Identification_PropertyType
	IdentificationInfo []*gml.MD_Identification_PropertyType `xml:"identificationInfo,omitempty"`

	// generated from element "contentInfo" of complex type MD_ContentInformation_PropertyType
	ContentInfo []*content.MD_ContentInformation_PropertyType `xml:"contentInfo,omitempty"`

	// generated from element "distributionInfo" of complex type MD_Distribution_PropertyType
	DistributionInfo *gml.MD_Distribution_PropertyType `xml:"distributionInfo,omitempty"`

	// generated from element "dataQualityInfo" of complex type DQ_DataQuality_PropertyType
	DataQualityInfo []*gml.DQ_DataQuality_PropertyType `xml:"dataQualityInfo,omitempty"`

	// generated from element "portrayalCatalogueInfo" of complex type MD_PortrayalCatalogueReference_PropertyType
	PortrayalCatalogueInfo []*portrayalcatalogue.MD_PortrayalCatalogueReference_PropertyType `xml:"portrayalCatalogueInfo,omitempty"`

	// generated from element "metadataConstraints" of complex type MD_Constraints_PropertyType
	MetadataConstraints []*gml.MD_Constraints_PropertyType `xml:"metadataConstraints,omitempty"`

	// generated from element "applicationSchemaInfo" of complex type MD_ApplicationSchemaInformation_PropertyType
	ApplicationSchemaInfo []*applicationschema.MD_ApplicationSchemaInformation_PropertyType `xml:"applicationSchemaInfo,omitempty"`

	// generated from element "metadataMaintenance" of complex type MD_MaintenanceInformation_PropertyType
	MetadataMaintenance *gml.MD_MaintenanceInformation_PropertyType `xml:"metadataMaintenance,omitempty"`

	// generated from element "series" of complex type DS_Aggregate_PropertyType
	Series []*DS_Aggregate_PropertyType `xml:"series,omitempty"`

	// generated from element "describes" of complex type DS_DataSet_PropertyType
	Describes []*DS_DataSet_PropertyType `xml:"describes,omitempty"`

	// generated from element "propertyType" of complex type ObjectReference_PropertyType
	PropertyType []*gcobase.ObjectReference_PropertyType `xml:"propertyType,omitempty"`

	// generated from element "featureType" of complex type ObjectReference_PropertyType
	FeatureType []*gcobase.ObjectReference_PropertyType `xml:"featureType,omitempty"`

	// generated from element "featureAttribute" of complex type ObjectReference_PropertyType
	FeatureAttribute []*gcobase.ObjectReference_PropertyType `xml:"featureAttribute,omitempty"`
}

// MD_Metadata_PropertyType Named source named complex type "MD_Metadata_PropertyType"
type MD_Metadata_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Metadata" of complex type MD_Metadata_Type
	MD_Metadata *MD_Metadata_Type `xml:"MD_Metadata,omitempty"`
}
