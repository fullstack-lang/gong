// generated code - do not edit
package spatialrepresentation

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

// From xsd simple type with enumerate restriction "MD_PixelOrientationCode_Type"
type Enum_MD_PixelOrientationCode_Type string

const (
	Enum_MD_PixelOrientationCode_Type_Center Enum_MD_PixelOrientationCode_Type = "center"
	Enum_MD_PixelOrientationCode_Type_LowerLeft Enum_MD_PixelOrientationCode_Type = "lowerLeft"
	Enum_MD_PixelOrientationCode_Type_LowerRight Enum_MD_PixelOrientationCode_Type = "lowerRight"
	Enum_MD_PixelOrientationCode_Type_UpperRight Enum_MD_PixelOrientationCode_Type = "upperRight"
	Enum_MD_PixelOrientationCode_Type_UpperLeft Enum_MD_PixelOrientationCode_Type = "upperLeft"
)

// insertion point for gongstructs declarations

// MD_GridSpatialRepresentation_Type Named source named complex type "MD_GridSpatialRepresentation_Type"
// Types and numbers of raster spatial objects in the dataset
type MD_GridSpatialRepresentation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "numberOfDimensions" of complex type Integer_PropertyType
	NumberOfDimensions *gml.Integer_PropertyType `xml:"numberOfDimensions,omitempty"`

	// generated from element "axisDimensionProperties" of complex type MD_Dimension_PropertyType
	AxisDimensionProperties []*MD_Dimension_PropertyType `xml:"axisDimensionProperties,omitempty"`

	// generated from element "cellGeometry" of complex type MD_CellGeometryCode_PropertyType
	CellGeometry *MD_CellGeometryCode_PropertyType `xml:"cellGeometry,omitempty"`

	// generated from element "transformationParameterAvailability" of complex type Boolean_PropertyType
	TransformationParameterAvailability *gml.Boolean_PropertyType `xml:"transformationParameterAvailability,omitempty"`
}

// MD_GridSpatialRepresentation_PropertyType Named source named complex type "MD_GridSpatialRepresentation_PropertyType"
type MD_GridSpatialRepresentation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_GridSpatialRepresentation" of complex type MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation *MD_GridSpatialRepresentation_Type `xml:"MD_GridSpatialRepresentation,omitempty"`

	// generated from element "MD_Georeferenceable" of complex type MD_Georeferenceable_Type
	MD_Georeferenceable *MD_Georeferenceable_Type `xml:"MD_Georeferenceable,omitempty"`

	// generated from element "MD_Georectified" of complex type MD_Georectified_Type
	MD_Georectified *MD_Georectified_Type `xml:"MD_Georectified,omitempty"`
}

// MD_VectorSpatialRepresentation_Type Named source named complex type "MD_VectorSpatialRepresentation_Type"
// Information about the vector spatial objects in the dataset
type MD_VectorSpatialRepresentation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "topologyLevel" of complex type MD_TopologyLevelCode_PropertyType
	TopologyLevel *MD_TopologyLevelCode_PropertyType `xml:"topologyLevel,omitempty"`

	// generated from element "geometricObjects" of complex type MD_GeometricObjects_PropertyType
	GeometricObjects []*MD_GeometricObjects_PropertyType `xml:"geometricObjects,omitempty"`
}

// MD_VectorSpatialRepresentation_PropertyType Named source named complex type "MD_VectorSpatialRepresentation_PropertyType"
type MD_VectorSpatialRepresentation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_VectorSpatialRepresentation" of complex type MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation *MD_VectorSpatialRepresentation_Type `xml:"MD_VectorSpatialRepresentation,omitempty"`
}

// AbstractMD_SpatialRepresentation_Type Named source named complex type "AbstractMD_SpatialRepresentation_Type"
// Digital mechanism used to represent spatial information
type AbstractMD_SpatialRepresentation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification
}

// MD_SpatialRepresentation_PropertyType Named source named complex type "MD_SpatialRepresentation_PropertyType"
type MD_SpatialRepresentation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_GridSpatialRepresentation" of complex type MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation *MD_GridSpatialRepresentation_Type `xml:"MD_GridSpatialRepresentation,omitempty"`

	// generated from element "MD_Georeferenceable" of complex type MD_Georeferenceable_Type
	MD_Georeferenceable *MD_Georeferenceable_Type `xml:"MD_Georeferenceable,omitempty"`

	// generated from element "MD_Georectified" of complex type MD_Georectified_Type
	MD_Georectified *MD_Georectified_Type `xml:"MD_Georectified,omitempty"`

	// generated from element "MD_VectorSpatialRepresentation" of complex type MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation *MD_VectorSpatialRepresentation_Type `xml:"MD_VectorSpatialRepresentation,omitempty"`
}

// MD_Georeferenceable_Type Named source named complex type "MD_Georeferenceable_Type"
type MD_Georeferenceable_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "numberOfDimensions" of complex type Integer_PropertyType
	NumberOfDimensions *gml.Integer_PropertyType `xml:"numberOfDimensions,omitempty"`

	// generated from element "axisDimensionProperties" of complex type MD_Dimension_PropertyType
	AxisDimensionProperties []*MD_Dimension_PropertyType `xml:"axisDimensionProperties,omitempty"`

	// generated from element "cellGeometry" of complex type MD_CellGeometryCode_PropertyType
	CellGeometry *MD_CellGeometryCode_PropertyType `xml:"cellGeometry,omitempty"`

	// generated from element "transformationParameterAvailability" of complex type Boolean_PropertyType
	TransformationParameterAvailability *gml.Boolean_PropertyType `xml:"transformationParameterAvailability,omitempty"`

	// generated from element "controlPointAvailability" of complex type Boolean_PropertyType
	ControlPointAvailability *gml.Boolean_PropertyType `xml:"controlPointAvailability,omitempty"`

	// generated from element "orientationParameterAvailability" of complex type Boolean_PropertyType
	OrientationParameterAvailability *gml.Boolean_PropertyType `xml:"orientationParameterAvailability,omitempty"`

	// generated from element "orientationParameterDescription" of complex type CharacterString_PropertyType
	OrientationParameterDescription *gml.CharacterString_PropertyType `xml:"orientationParameterDescription,omitempty"`

	// generated from element "georeferencedParameters" of complex type Record_PropertyType
	GeoreferencedParameters *gml.Record_PropertyType `xml:"georeferencedParameters,omitempty"`

	// generated from element "parameterCitation" of complex type CI_Citation_PropertyType
	ParameterCitation []*gml.CI_Citation_PropertyType `xml:"parameterCitation,omitempty"`
}

// MD_Georeferenceable_PropertyType Named source named complex type "MD_Georeferenceable_PropertyType"
type MD_Georeferenceable_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Georeferenceable" of complex type MD_Georeferenceable_Type
	MD_Georeferenceable *MD_Georeferenceable_Type `xml:"MD_Georeferenceable,omitempty"`
}

// MD_Dimension_Type Named source named complex type "MD_Dimension_Type"
type MD_Dimension_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "dimensionName" of complex type MD_DimensionNameTypeCode_PropertyType
	DimensionName *MD_DimensionNameTypeCode_PropertyType `xml:"dimensionName,omitempty"`

	// generated from element "dimensionSize" of complex type Integer_PropertyType
	DimensionSize *gml.Integer_PropertyType `xml:"dimensionSize,omitempty"`

	// generated from element "resolution" of complex type Measure_PropertyType
	Resolution *gml.Measure_PropertyType `xml:"resolution,omitempty"`
}

// MD_Dimension_PropertyType Named source named complex type "MD_Dimension_PropertyType"
type MD_Dimension_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Dimension" of complex type MD_Dimension_Type
	MD_Dimension *MD_Dimension_Type `xml:"MD_Dimension,omitempty"`
}

// MD_Georectified_Type Named source named complex type "MD_Georectified_Type"
type MD_Georectified_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "numberOfDimensions" of complex type Integer_PropertyType
	NumberOfDimensions *gml.Integer_PropertyType `xml:"numberOfDimensions,omitempty"`

	// generated from element "axisDimensionProperties" of complex type MD_Dimension_PropertyType
	AxisDimensionProperties []*MD_Dimension_PropertyType `xml:"axisDimensionProperties,omitempty"`

	// generated from element "cellGeometry" of complex type MD_CellGeometryCode_PropertyType
	CellGeometry *MD_CellGeometryCode_PropertyType `xml:"cellGeometry,omitempty"`

	// generated from element "transformationParameterAvailability" of complex type Boolean_PropertyType
	TransformationParameterAvailability *gml.Boolean_PropertyType `xml:"transformationParameterAvailability,omitempty"`

	// generated from element "checkPointAvailability" of complex type Boolean_PropertyType
	CheckPointAvailability *gml.Boolean_PropertyType `xml:"checkPointAvailability,omitempty"`

	// generated from element "checkPointDescription" of complex type CharacterString_PropertyType
	CheckPointDescription *gml.CharacterString_PropertyType `xml:"checkPointDescription,omitempty"`

	// generated from element "cornerPoints" of complex type GM_Point_PropertyType
	CornerPoints []*gml.GM_Point_PropertyType `xml:"cornerPoints,omitempty"`

	// generated from element "centerPoint" of complex type GM_Point_PropertyType
	CenterPoint *gml.GM_Point_PropertyType `xml:"centerPoint,omitempty"`

	// generated from element "pointInPixel" of complex type MD_PixelOrientationCode_PropertyType
	PointInPixel *MD_PixelOrientationCode_PropertyType `xml:"pointInPixel,omitempty"`

	// generated from element "transformationDimensionDescription" of complex type CharacterString_PropertyType
	TransformationDimensionDescription *gml.CharacterString_PropertyType `xml:"transformationDimensionDescription,omitempty"`

	// generated from element "transformationDimensionMapping" of complex type CharacterString_PropertyType
	TransformationDimensionMapping *gml.CharacterString_PropertyType `xml:"transformationDimensionMapping,omitempty"`
}

// MD_Georectified_PropertyType Named source named complex type "MD_Georectified_PropertyType"
type MD_Georectified_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_Georectified" of complex type MD_Georectified_Type
	MD_Georectified *MD_Georectified_Type `xml:"MD_Georectified,omitempty"`
}

// MD_GeometricObjects_Type Named source named complex type "MD_GeometricObjects_Type"
type MD_GeometricObjects_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "geometricObjectType" of complex type MD_GeometricObjectTypeCode_PropertyType
	GeometricObjectType *MD_GeometricObjectTypeCode_PropertyType `xml:"geometricObjectType,omitempty"`

	// generated from element "geometricObjectCount" of complex type Integer_PropertyType
	GeometricObjectCount *gml.Integer_PropertyType `xml:"geometricObjectCount,omitempty"`
}

// MD_GeometricObjects_PropertyType Named source named complex type "MD_GeometricObjects_PropertyType"
type MD_GeometricObjects_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_GeometricObjects" of complex type MD_GeometricObjects_Type
	MD_GeometricObjects *MD_GeometricObjects_Type `xml:"MD_GeometricObjects,omitempty"`
}

// MD_PixelOrientationCode_PropertyType Named source named complex type "MD_PixelOrientationCode_PropertyType"
type MD_PixelOrientationCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_PixelOrientationCode" of simple type gmd:MD_PixelOrientationCode_Type
	MD_PixelOrientationCode Enum_MD_PixelOrientationCode_Type `xml:"MD_PixelOrientationCode,omitempty"`
}

// MD_TopologyLevelCode_PropertyType Named source named complex type "MD_TopologyLevelCode_PropertyType"
type MD_TopologyLevelCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_TopologyLevelCode" of complex type CodeListValue_Type
	MD_TopologyLevelCode *gcobase.CodeListValue_Type `xml:"MD_TopologyLevelCode,omitempty"`
}

// MD_GeometricObjectTypeCode_PropertyType Named source named complex type "MD_GeometricObjectTypeCode_PropertyType"
type MD_GeometricObjectTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_GeometricObjectTypeCode" of complex type CodeListValue_Type
	MD_GeometricObjectTypeCode *gcobase.CodeListValue_Type `xml:"MD_GeometricObjectTypeCode,omitempty"`
}

// MD_CellGeometryCode_PropertyType Named source named complex type "MD_CellGeometryCode_PropertyType"
type MD_CellGeometryCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_CellGeometryCode" of complex type CodeListValue_Type
	MD_CellGeometryCode *gcobase.CodeListValue_Type `xml:"MD_CellGeometryCode,omitempty"`
}

// MD_DimensionNameTypeCode_PropertyType Named source named complex type "MD_DimensionNameTypeCode_PropertyType"
type MD_DimensionNameTypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_DimensionNameTypeCode" of complex type CodeListValue_Type
	MD_DimensionNameTypeCode *gcobase.CodeListValue_Type `xml:"MD_DimensionNameTypeCode,omitempty"`
}
