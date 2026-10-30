// generated code - do not edit
package applicationschema

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

// MD_ApplicationSchemaInformation_Type Named source named complex type "MD_ApplicationSchemaInformation_Type"
// Information about the application schema used to build the dataset
type MD_ApplicationSchemaInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type CI_Citation_PropertyType
	NameXSD *gml.CI_Citation_PropertyType `xml:"name,omitempty"`

	// generated from element "schemaLanguage" of complex type CharacterString_PropertyType
	SchemaLanguage *gml.CharacterString_PropertyType `xml:"schemaLanguage,omitempty"`

	// generated from element "constraintLanguage" of complex type CharacterString_PropertyType
	ConstraintLanguage *gml.CharacterString_PropertyType `xml:"constraintLanguage,omitempty"`

	// generated from element "schemaAscii" of complex type CharacterString_PropertyType
	SchemaAscii *gml.CharacterString_PropertyType `xml:"schemaAscii,omitempty"`

	// generated from element "graphicsFile" of complex type Binary_PropertyType
	GraphicsFile *gml.Binary_PropertyType `xml:"graphicsFile,omitempty"`

	// generated from element "softwareDevelopmentFile" of complex type Binary_PropertyType
	SoftwareDevelopmentFile *gml.Binary_PropertyType `xml:"softwareDevelopmentFile,omitempty"`

	// generated from element "softwareDevelopmentFileFormat" of complex type CharacterString_PropertyType
	SoftwareDevelopmentFileFormat *gml.CharacterString_PropertyType `xml:"softwareDevelopmentFileFormat,omitempty"`
}

// MD_ApplicationSchemaInformation_PropertyType Named source named complex type "MD_ApplicationSchemaInformation_PropertyType"
type MD_ApplicationSchemaInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_ApplicationSchemaInformation" of complex type MD_ApplicationSchemaInformation_Type
	MD_ApplicationSchemaInformation *MD_ApplicationSchemaInformation_Type `xml:"MD_ApplicationSchemaInformation,omitempty"`
}
