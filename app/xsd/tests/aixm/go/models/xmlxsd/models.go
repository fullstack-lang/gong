// generated code - do not edit
package xmlxsd

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "lang_Member_1"
type Enum_Lang_Member_1 string

const (
	Enum_Lang_Member_1_Empty Enum_Lang_Member_1 = ""
)

// From xsd simple type with enumerate restriction "space"
type Enum_Space string

const (
	Enum_Space_Default_XSD Enum_Space = "default"
	Enum_Space_Preserve Enum_Space = "preserve"
)

// insertion point for gongstructs declarations

// Attr_lang UnNamed source top-level attribute "lang"
type Attr_lang struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang"
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Attr_space UnNamed source top-level attribute "space"
type Attr_space struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace space"
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`
}

// Attr_base UnNamed source top-level attribute "base"
type Attr_base struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace base"
	Base string `xml:"http://www.w3.org/XML/1998/namespace base,attr,omitempty"`
}

// Attr_id UnNamed source top-level attribute "id"
type Attr_id struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace id"
	Id string `xml:"http://www.w3.org/XML/1998/namespace id,attr,omitempty"`
}

// AttributeGroup_specialAttrs UnNamed source named attribute group "specialAttrs"
type AttributeGroup_specialAttrs struct {

	// insertion point for fields

	// generated from attribute "xml:base"
	Attr_base

	// generated from attribute "xml:lang"
	Attr_lang

	// generated from attribute "xml:space"
	Attr_space

	// generated from attribute "xml:id"
	Attr_id
}

// Xmlxsd_SchemaElement synthetic Gongstruct for schema package xmlxsd
type Xmlxsd_SchemaElement struct {
	Name string `xml:"-"`
}
