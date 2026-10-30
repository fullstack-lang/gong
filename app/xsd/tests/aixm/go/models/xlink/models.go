// generated code - do not edit
package xlink

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xmlxsd"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *xmlxsd.Xmlxsd_SchemaElement

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "typeType"
type Enum_TypeType string

const (
	Enum_TypeType_Simple Enum_TypeType = "simple"
	Enum_TypeType_Extended Enum_TypeType = "extended"
	Enum_TypeType_Title Enum_TypeType = "title"
	Enum_TypeType_Resource Enum_TypeType = "resource"
	Enum_TypeType_Locator Enum_TypeType = "locator"
	Enum_TypeType_Arc Enum_TypeType = "arc"
)

// From xsd simple type with enumerate restriction "showType"
type Enum_ShowType string

const (
	Enum_ShowType_New_XSD Enum_ShowType = "new"
	Enum_ShowType_Replace Enum_ShowType = "replace"
	Enum_ShowType_Embed Enum_ShowType = "embed"
	Enum_ShowType_Other Enum_ShowType = "other"
	Enum_ShowType_None Enum_ShowType = "none"
)

// From xsd simple type with enumerate restriction "actuateType"
type Enum_ActuateType string

const (
	Enum_ActuateType_OnLoad Enum_ActuateType = "onLoad"
	Enum_ActuateType_OnRequest Enum_ActuateType = "onRequest"
	Enum_ActuateType_Other Enum_ActuateType = "other"
	Enum_ActuateType_None Enum_ActuateType = "none"
)

// insertion point for gongstructs declarations

// Attr_type UnNamed source top-level attribute "type"
type Attr_type struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink type"
	Type_XSD string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
}

// Attr_href UnNamed source top-level attribute "href"
type Attr_href struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink href"
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
}

// Attr_role UnNamed source top-level attribute "role"
type Attr_role struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink role"
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
}

// Attr_arcrole UnNamed source top-level attribute "arcrole"
type Attr_arcrole struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink arcrole"
	Arcrole string `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
}

// Attr_title UnNamed source top-level attribute "title"
type Attr_title struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink title"
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
}

// Attr_show UnNamed source top-level attribute "show"
type Attr_show struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink show"
	Show string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
}

// Attr_actuate UnNamed source top-level attribute "actuate"
type Attr_actuate struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink actuate"
	Actuate string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

// Attr_label UnNamed source top-level attribute "label"
type Attr_label struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink label"
	Label string `xml:"http://www.w3.org/1999/xlink label,attr,omitempty"`
}

// Attr_from UnNamed source top-level attribute "from"
type Attr_from struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink from"
	From string `xml:"http://www.w3.org/1999/xlink from,attr,omitempty"`
}

// Attr_to UnNamed source top-level attribute "to"
type Attr_to struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink to"
	To string `xml:"http://www.w3.org/1999/xlink to,attr,omitempty"`
}

// AttributeGroup_simpleAttrs UnNamed source named attribute group "simpleAttrs"
type AttributeGroup_simpleAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from attribute "xlink:href"
	Attr_href

	// generated from attribute "xlink:role"
	Attr_role

	// generated from attribute "xlink:arcrole"
	Attr_arcrole

	// generated from attribute "xlink:title"
	Attr_title

	// generated from attribute "xlink:show"
	Attr_show

	// generated from attribute "xlink:actuate"
	Attr_actuate
}

// Group_simpleModel UnNamed source named group "simpleModel"
type Group_simpleModel struct {

	// insertion point for fields
}

// Simple Named source named complex type "simple"
// 
//     Intended for use as the type of user-declared elements to make them
//     simple links.
//    
type Simple struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:simpleAttrs"
	AttributeGroup_simpleAttrs
}

// AttributeGroup_extendedAttrs UnNamed source named attribute group "extendedAttrs"
type AttributeGroup_extendedAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from attribute "xlink:role"
	Attr_role

	// generated from attribute "xlink:title"
	Attr_title
}

// Group_extendedModel UnNamed source named group "extendedModel"
type Group_extendedModel struct {

	// insertion point for fields

	// generated from element "title" of complex type titleEltType
	Title *TitleEltType `xml:"title,omitempty"`

	// generated from element "resource" of complex type resourceType
	Resource *ResourceType `xml:"resource,omitempty"`

	// generated from element "locator" of complex type locatorType
	Locator *LocatorType `xml:"locator,omitempty"`

	// generated from element "arc" of complex type arcType
	Arc *ArcType `xml:"arc,omitempty"`
}

// Extended Named source named complex type "extended"
// 
//     Intended for use as the type of user-declared elements to make them
//     extended links.
//     Note that the elements referenced in the content model are all abstract.
//     The intention is that by simply declaring elements with these as their
//     substitutionGroup, all the right things will happen.
//    
type Extended struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:extendedAttrs"
	AttributeGroup_extendedAttrs

	// generated from element "title" of complex type titleEltType
	Title_1 *TitleEltType `xml:"title,omitempty"`

	// generated from element "resource" of complex type resourceType
	Resource *ResourceType `xml:"resource,omitempty"`

	// generated from element "locator" of complex type locatorType
	Locator *LocatorType `xml:"locator,omitempty"`

	// generated from element "arc" of complex type arcType
	Arc *ArcType `xml:"arc,omitempty"`
}

// AttributeGroup_titleAttrs UnNamed source named attribute group "titleAttrs"
type AttributeGroup_titleAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from cross-package attribute "xml:lang"
	xmlxsd.Attr_lang
}

// Group_titleModel UnNamed source named group "titleModel"
type Group_titleModel struct {

	// insertion point for fields
}

// TitleEltType Named source named complex type "titleEltType"
type TitleEltType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:titleAttrs"
	AttributeGroup_titleAttrs
}

// AttributeGroup_resourceAttrs UnNamed source named attribute group "resourceAttrs"
type AttributeGroup_resourceAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from attribute "xlink:role"
	Attr_role

	// generated from attribute "xlink:title"
	Attr_title

	// generated from attribute "xlink:label"
	Attr_label
}

// Group_resourceModel UnNamed source named group "resourceModel"
type Group_resourceModel struct {

	// insertion point for fields
}

// ResourceType Named source named complex type "resourceType"
type ResourceType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:resourceAttrs"
	AttributeGroup_resourceAttrs
}

// AttributeGroup_locatorAttrs UnNamed source named attribute group "locatorAttrs"
type AttributeGroup_locatorAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from attribute "xlink:href"
	Attr_href

	// generated from attribute "xlink:role"
	Attr_role

	// generated from attribute "xlink:title"
	Attr_title

	// generated from attribute "xlink:label"
	Attr_label
}

// Group_locatorModel UnNamed source named group "locatorModel"
type Group_locatorModel struct {

	// insertion point for fields

	// generated from element "title" of complex type titleEltType
	Title []*TitleEltType `xml:"title,omitempty"`
}

// LocatorType Named source named complex type "locatorType"
type LocatorType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:locatorAttrs"
	AttributeGroup_locatorAttrs

	// generated from element "title" of complex type titleEltType
	Title_1 []*TitleEltType `xml:"title,omitempty"`
}

// AttributeGroup_arcAttrs UnNamed source named attribute group "arcAttrs"
type AttributeGroup_arcAttrs struct {

	// insertion point for fields

	// generated from attribute "xlink:type"
	Attr_type

	// generated from attribute "xlink:arcrole"
	Attr_arcrole

	// generated from attribute "xlink:title"
	Attr_title

	// generated from attribute "xlink:show"
	Attr_show

	// generated from attribute "xlink:actuate"
	Attr_actuate

	// generated from attribute "xlink:from"
	Attr_from

	// generated from attribute "xlink:to"
	Attr_to
}

// Group_arcModel UnNamed source named group "arcModel"
type Group_arcModel struct {

	// insertion point for fields

	// generated from element "title" of complex type titleEltType
	Title []*TitleEltType `xml:"title,omitempty"`
}

// ArcType Named source named complex type "arcType"
type ArcType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "xlink:arcAttrs"
	AttributeGroup_arcAttrs

	// generated from element "title" of complex type titleEltType
	Title_1 []*TitleEltType `xml:"title,omitempty"`
}
