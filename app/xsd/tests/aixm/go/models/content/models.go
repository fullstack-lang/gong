// generated code - do not edit
package content

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *gcobase.AbstractObject_Type
var _ *gml.DynamicFeatureType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// MD_FeatureCatalogueDescription_Type Named source named complex type "MD_FeatureCatalogueDescription_Type"
// Information identifing the feature catalogue
type MD_FeatureCatalogueDescription_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "complianceCode" of complex type Boolean_PropertyType
	ComplianceCode *gml.Boolean_PropertyType `xml:"complianceCode,omitempty"`

	// generated from element "language" of complex type CharacterString_PropertyType
	Language []*gml.CharacterString_PropertyType `xml:"language,omitempty"`

	// generated from element "includedWithDataset" of complex type Boolean_PropertyType
	IncludedWithDataset *gml.Boolean_PropertyType `xml:"includedWithDataset,omitempty"`

	// generated from element "featureTypes" of complex type GenericName_PropertyType
	FeatureTypes []*gml.GenericName_PropertyType `xml:"featureTypes,omitempty"`

	// generated from element "featureCatalogueCitation" of complex type CI_Citation_PropertyType
	FeatureCatalogueCitation []*gml.CI_Citation_PropertyType `xml:"featureCatalogueCitation,omitempty"`
}

// MD_FeatureCatalogueDescription_PropertyType Named source named complex type "MD_FeatureCatalogueDescription_PropertyType"
type MD_FeatureCatalogueDescription_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_FeatureCatalogueDescription" of complex type MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription *MD_FeatureCatalogueDescription_Type `xml:"MD_FeatureCatalogueDescription,omitempty"`
}

// MD_CoverageDescription_Type Named source named complex type "MD_CoverageDescription_Type"
// Information about the domain of the raster cell
type MD_CoverageDescription_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "attributeDescription" of complex type RecordType_PropertyType
	AttributeDescription *gml.RecordType_PropertyType `xml:"attributeDescription,omitempty"`

	// generated from element "contentType" of complex type MD_CoverageContentTypeCode_PropertyType
	ContentType *MD_CoverageContentTypeCode_PropertyType `xml:"contentType,omitempty"`

	// generated from element "dimension" of complex type MD_RangeDimension_PropertyType
	Dimension []*MD_RangeDimension_PropertyType `xml:"dimension,omitempty"`
}

// MD_CoverageDescription_PropertyType Named source named complex type "MD_CoverageDescription_PropertyType"
type MD_CoverageDescription_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_CoverageDescription" of complex type MD_CoverageDescription_Type
	MD_CoverageDescription *MD_CoverageDescription_Type `xml:"MD_CoverageDescription,omitempty"`

	// generated from element "MD_ImageDescription" of complex type MD_ImageDescription_Type
	MD_ImageDescription *MD_ImageDescription_Type `xml:"MD_ImageDescription,omitempty"`
}

// MD_ImageDescription_Type Named source named complex type "MD_ImageDescription_Type"
// Information about an image's suitability for use
type MD_ImageDescription_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "attributeDescription" of complex type RecordType_PropertyType
	AttributeDescription *gml.RecordType_PropertyType `xml:"attributeDescription,omitempty"`

	// generated from element "contentType" of complex type MD_CoverageContentTypeCode_PropertyType
	ContentType *MD_CoverageContentTypeCode_PropertyType `xml:"contentType,omitempty"`

	// generated from element "dimension" of complex type MD_RangeDimension_PropertyType
	Dimension []*MD_RangeDimension_PropertyType `xml:"dimension,omitempty"`

	// generated from element "illuminationElevationAngle" of complex type Real_PropertyType
	IlluminationElevationAngle *gml.Real_PropertyType `xml:"illuminationElevationAngle,omitempty"`

	// generated from element "illuminationAzimuthAngle" of complex type Real_PropertyType
	IlluminationAzimuthAngle *gml.Real_PropertyType `xml:"illuminationAzimuthAngle,omitempty"`

	// generated from element "imagingCondition" of complex type MD_ImagingConditionCode_PropertyType
	ImagingCondition *MD_ImagingConditionCode_PropertyType `xml:"imagingCondition,omitempty"`

	// generated from element "imageQualityCode" of complex type MD_Identifier_PropertyType
	ImageQualityCode *gml.MD_Identifier_PropertyType `xml:"imageQualityCode,omitempty"`

	// generated from element "cloudCoverPercentage" of complex type Real_PropertyType
	CloudCoverPercentage *gml.Real_PropertyType `xml:"cloudCoverPercentage,omitempty"`

	// generated from element "processingLevelCode" of complex type MD_Identifier_PropertyType
	ProcessingLevelCode *gml.MD_Identifier_PropertyType `xml:"processingLevelCode,omitempty"`

	// generated from element "compressionGenerationQuantity" of complex type Integer_PropertyType
	CompressionGenerationQuantity *gml.Integer_PropertyType `xml:"compressionGenerationQuantity,omitempty"`

	// generated from element "triangulationIndicator" of complex type Boolean_PropertyType
	TriangulationIndicator *gml.Boolean_PropertyType `xml:"triangulationIndicator,omitempty"`

	// generated from element "radiometricCalibrationDataAvailability" of complex type Boolean_PropertyType
	RadiometricCalibrationDataAvailability *gml.Boolean_PropertyType `xml:"radiometricCalibrationDataAvailability,omitempty"`

	// generated from element "cameraCalibrationInformationAvailability" of complex type Boolean_PropertyType
	CameraCalibrationInformationAvailability *gml.Boolean_PropertyType `xml:"cameraCalibrationInformationAvailability,omitempty"`

	// generated from element "filmDistortionInformationAvailability" of complex type Boolean_PropertyType
	FilmDistortionInformationAvailability *gml.Boolean_PropertyType `xml:"filmDistortionInformationAvailability,omitempty"`

	// generated from element "lensDistortionInformationAvailability" of complex type Boolean_PropertyType
	LensDistortionInformationAvailability *gml.Boolean_PropertyType `xml:"lensDistortionInformationAvailability,omitempty"`
}

// MD_ImageDescription_PropertyType Named source named complex type "MD_ImageDescription_PropertyType"
type MD_ImageDescription_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_ImageDescription" of complex type MD_ImageDescription_Type
	MD_ImageDescription *MD_ImageDescription_Type `xml:"MD_ImageDescription,omitempty"`
}

// AbstractMD_ContentInformation_Type Named source named complex type "AbstractMD_ContentInformation_Type"
type AbstractMD_ContentInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification
}

// MD_ContentInformation_PropertyType Named source named complex type "MD_ContentInformation_PropertyType"
type MD_ContentInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_FeatureCatalogueDescription" of complex type MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription *MD_FeatureCatalogueDescription_Type `xml:"MD_FeatureCatalogueDescription,omitempty"`

	// generated from element "MD_CoverageDescription" of complex type MD_CoverageDescription_Type
	MD_CoverageDescription *MD_CoverageDescription_Type `xml:"MD_CoverageDescription,omitempty"`

	// generated from element "MD_ImageDescription" of complex type MD_ImageDescription_Type
	MD_ImageDescription *MD_ImageDescription_Type `xml:"MD_ImageDescription,omitempty"`
}

// MD_RangeDimension_Type Named source named complex type "MD_RangeDimension_Type"
// Set of adjacent wavelengths in the electro-magnetic spectrum with a common characteristic, such as the visible band
type MD_RangeDimension_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "sequenceIdentifier" of complex type MemberName_PropertyType
	SequenceIdentifier *gml.MemberName_PropertyType `xml:"sequenceIdentifier,omitempty"`

	// generated from element "descriptor" of complex type CharacterString_PropertyType
	Descriptor *gml.CharacterString_PropertyType `xml:"descriptor,omitempty"`
}

// MD_RangeDimension_PropertyType Named source named complex type "MD_RangeDimension_PropertyType"
type MD_RangeDimension_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_RangeDimension" of complex type MD_RangeDimension_Type
	MD_RangeDimension *MD_RangeDimension_Type `xml:"MD_RangeDimension,omitempty"`

	// generated from element "MD_Band" of complex type MD_Band_Type
	MD_Band *MD_Band_Type `xml:"MD_Band,omitempty"`
}

// MD_Band_Type Named source named complex type "MD_Band_Type"
type MD_Band_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "sequenceIdentifier" of complex type MemberName_PropertyType
	SequenceIdentifier *gml.MemberName_PropertyType `xml:"sequenceIdentifier,omitempty"`

	// generated from element "descriptor" of complex type CharacterString_PropertyType
	Descriptor *gml.CharacterString_PropertyType `xml:"descriptor,omitempty"`

	// generated from element "maxValue" of complex type Real_PropertyType
	MaxValue *gml.Real_PropertyType `xml:"maxValue,omitempty"`

	// generated from element "minValue" of complex type Real_PropertyType
	MinValue *gml.Real_PropertyType `xml:"minValue,omitempty"`

	// generated from element "units" of complex type UomLength_PropertyType
	Units *gml.UomLength_PropertyType `xml:"units,omitempty"`

	// generated from element "peakResponse" of complex type Real_PropertyType
	PeakResponse *gml.Real_PropertyType `xml:"peakResponse,omitempty"`

	// generated from element "bitsPerValue" of complex type Integer_PropertyType
	BitsPerValue *gml.Integer_PropertyType `xml:"bitsPerValue,omitempty"`

	// generated from element "toneGradation" of complex type Integer_PropertyType
	ToneGradation *gml.Integer_PropertyType `xml:"toneGradation,omitempty"`

	// generated from element "scaleFactor" of complex type Real_PropertyType
	ScaleFactor *gml.Real_PropertyType `xml:"scaleFactor,omitempty"`

	// generated from element "offset" of complex type Real_PropertyType
	Offset *gml.Real_PropertyType `xml:"offset,omitempty"`
}

// MD_Band_PropertyType Named source named complex type "MD_Band_PropertyType"
type MD_Band_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Band" of complex type MD_Band_Type
	MD_Band *MD_Band_Type `xml:"MD_Band,omitempty"`
}

// MD_CoverageContentTypeCode_PropertyType Named source named complex type "MD_CoverageContentTypeCode_PropertyType"
type MD_CoverageContentTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_CoverageContentTypeCode" of complex type CodeListValue_Type
	MD_CoverageContentTypeCode *gcobase.CodeListValue_Type `xml:"MD_CoverageContentTypeCode,omitempty"`
}

// MD_ImagingConditionCode_PropertyType Named source named complex type "MD_ImagingConditionCode_PropertyType"
type MD_ImagingConditionCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ImagingConditionCode" of complex type CodeListValue_Type
	MD_ImagingConditionCode *gcobase.CodeListValue_Type `xml:"MD_ImagingConditionCode,omitempty"`
}
