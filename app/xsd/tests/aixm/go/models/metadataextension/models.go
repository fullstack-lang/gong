// generated code - do not edit
package metadataextension

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

// From xsd simple type with enumerate restriction "MD_ObligationCode_Type"
type Enum_MD_ObligationCode_Type string

const (
	Enum_MD_ObligationCode_Type_Mandatory Enum_MD_ObligationCode_Type = "mandatory"
	Enum_MD_ObligationCode_Type_Optional Enum_MD_ObligationCode_Type = "optional"
	Enum_MD_ObligationCode_Type_Conditional Enum_MD_ObligationCode_Type = "conditional"
)

// insertion point for gongstructs declarations

// MD_ExtendedElementInformation_Type Named source named complex type "MD_ExtendedElementInformation_Type"
// New metadata element, not found in ISO 19115, which is required to describe geographic data
type MD_ExtendedElementInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "name" of complex type CharacterString_PropertyType
	NameXSD *gml.CharacterString_PropertyType `xml:"name,omitempty"`

	// generated from element "shortName" of complex type CharacterString_PropertyType
	ShortName *gml.CharacterString_PropertyType `xml:"shortName,omitempty"`

	// generated from element "domainCode" of complex type Integer_PropertyType
	DomainCode *gml.Integer_PropertyType `xml:"domainCode,omitempty"`

	// generated from element "definition" of complex type CharacterString_PropertyType
	Definition *gml.CharacterString_PropertyType `xml:"definition,omitempty"`

	// generated from element "obligation" of complex type MD_ObligationCode_PropertyType
	Obligation *MD_ObligationCode_PropertyType `xml:"obligation,omitempty"`

	// generated from element "condition" of complex type CharacterString_PropertyType
	Condition *gml.CharacterString_PropertyType `xml:"condition,omitempty"`

	// generated from element "dataType" of complex type MD_DatatypeCode_PropertyType
	DataType *MD_DatatypeCode_PropertyType `xml:"dataType,omitempty"`

	// generated from element "maximumOccurrence" of complex type CharacterString_PropertyType
	MaximumOccurrence *gml.CharacterString_PropertyType `xml:"maximumOccurrence,omitempty"`

	// generated from element "domainValue" of complex type CharacterString_PropertyType
	DomainValue *gml.CharacterString_PropertyType `xml:"domainValue,omitempty"`

	// generated from element "parentEntity" of complex type CharacterString_PropertyType
	ParentEntity []*gml.CharacterString_PropertyType `xml:"parentEntity,omitempty"`

	// generated from element "rule" of complex type CharacterString_PropertyType
	Rule *gml.CharacterString_PropertyType `xml:"rule,omitempty"`

	// generated from element "rationale" of complex type CharacterString_PropertyType
	Rationale []*gml.CharacterString_PropertyType `xml:"rationale,omitempty"`

	// generated from element "source" of complex type CI_ResponsibleParty_PropertyType
	Source []*gml.CI_ResponsibleParty_PropertyType `xml:"source,omitempty"`
}

// MD_ExtendedElementInformation_PropertyType Named source named complex type "MD_ExtendedElementInformation_PropertyType"
type MD_ExtendedElementInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_ExtendedElementInformation" of complex type MD_ExtendedElementInformation_Type
	MD_ExtendedElementInformation *MD_ExtendedElementInformation_Type `xml:"MD_ExtendedElementInformation,omitempty"`
}

// MD_MetadataExtensionInformation_Type Named source named complex type "MD_MetadataExtensionInformation_Type"
// Information describing metadata extensions.
type MD_MetadataExtensionInformation_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	gcobase.AttributeGroup_ObjectIdentification

	// generated from element "extensionOnLineResource" of complex type CI_OnlineResource_PropertyType
	ExtensionOnLineResource *gml.CI_OnlineResource_PropertyType `xml:"extensionOnLineResource,omitempty"`

	// generated from element "extendedElementInformation" of complex type MD_ExtendedElementInformation_PropertyType
	ExtendedElementInformation []*MD_ExtendedElementInformation_PropertyType `xml:"extendedElementInformation,omitempty"`
}

// MD_MetadataExtensionInformation_PropertyType Named source named complex type "MD_MetadataExtensionInformation_PropertyType"
type MD_MetadataExtensionInformation_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from attribute group "gco:ObjectReference"
	gcobase.AttributeGroup_ObjectReference

	// generated from element "MD_MetadataExtensionInformation" of complex type MD_MetadataExtensionInformation_Type
	MD_MetadataExtensionInformation *MD_MetadataExtensionInformation_Type `xml:"MD_MetadataExtensionInformation,omitempty"`
}

// MD_ObligationCode_PropertyType Named source named complex type "MD_ObligationCode_PropertyType"
type MD_ObligationCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_ObligationCode" of simple type gmd:MD_ObligationCode_Type
	MD_ObligationCode Enum_MD_ObligationCode_Type `xml:"MD_ObligationCode,omitempty"`
}

// MD_DatatypeCode_PropertyType Named source named complex type "MD_DatatypeCode_PropertyType"
type MD_DatatypeCode_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from cross-package attribute ref "gco:nilReason"
	gcobase.Attr_nilReason

	// generated from element "MD_DatatypeCode" of complex type CodeListValue_Type
	MD_DatatypeCode *gcobase.CodeListValue_Type `xml:"MD_DatatypeCode,omitempty"`
}
