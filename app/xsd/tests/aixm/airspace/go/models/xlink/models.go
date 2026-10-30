// generated code - do not edit
package xlink

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

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

// Xlink_SchemaElement synthetic Gongstruct for schema package xlink
type Xlink_SchemaElement struct {
	Name string `xml:"-"`
}
