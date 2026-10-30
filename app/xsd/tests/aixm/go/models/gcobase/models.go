// generated code - do not edit
package gcobase

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *xlink.Simple

// insertion point for enums declaration

// insertion point for gongstructs declarations

// AbstractObject_Type Named source named complex type "AbstractObject_Type"
type AbstractObject_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "gco:ObjectIdentification"
	AttributeGroup_ObjectIdentification
}

// AttributeGroup_ObjectReference UnNamed source named attribute group "ObjectReference"
type AttributeGroup_ObjectReference struct {

	// insertion point for fields

	// generated from attribute "uuidref"
	Uuidref string `xml:"uuidref,attr,omitempty"`

	// generated from attribute group "xlink:simpleAttrs"
	xlink.AttributeGroup_simpleAttrs
}

// Attr_nilReason UnNamed source top-level attribute "nilReason"
type Attr_nilReason struct {

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`
}

// ObjectReference_PropertyType Named source named complex type "ObjectReference_PropertyType"
type ObjectReference_PropertyType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "nilReason"
	NilReason string `xml:"nilReason,attr,omitempty"`

	// generated from attribute group "gco:ObjectReference"
	AttributeGroup_ObjectReference
}

// AttributeGroup_ObjectIdentification UnNamed source named attribute group "ObjectIdentification"
type AttributeGroup_ObjectIdentification struct {

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "uuid"
	Uuid string `xml:"uuid,attr,omitempty"`
}

// CodeListValue_Type Named source named complex type "CodeListValue_Type"
type CodeListValue_Type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeList"
	CodeList string `xml:"codeList,attr,omitempty"`

	// generated from attribute "codeListValue"
	CodeListValue string `xml:"codeListValue,attr,omitempty"`

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// Attr_isoType UnNamed source top-level attribute "isoType"
type Attr_isoType struct {

	// insertion point for fields

	// generated from attribute "isoType"
	IsoType string `xml:"isoType,attr,omitempty"`
}
