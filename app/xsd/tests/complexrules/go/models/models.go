// generated code - do not edit
package models

import (
	"encoding/xml"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr
var _ *common.BaseItemType
var _ *refs.RefTargetType

// insertion point for enums declaration

// insertion point for gongstructs declarations

// DerivedItemType Named source named complex type "derivedItemType"
type DerivedItemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "code" of simple type tns:codeSimpleType
	Code string `xml:"code,omitempty"`

	// generated from element "score" of simple type tns:scoreSimpleType
	Score string `xml:"score,omitempty"`

	// generated from attribute "targetRef"
	TargetRef string `xml:"targetRef,attr,omitempty"`

	// generated from element "optionA" of simple type xs:string
	OptionA string `xml:"optionA,omitempty"`

	// generated from element "optionB" of simple type xs:int
	OptionB int `xml:"optionB,omitempty"`
}

// Catalog Named source element catalog within schema
type Catalog struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"catalog"`

	// generated from element "target" of complex type refTargetType
	Target []*refs.RefTargetType `xml:"target,omitempty"`

	// generated from element "item" of complex type derivedItemType
	Item []*DerivedItemType `xml:"item,omitempty"`
}
