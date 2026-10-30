// generated code - do not edit
package common

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// insertion point for gongstructs declarations

// BaseItemType Named source named complex type "baseItemType"
type BaseItemType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "code" of simple type tns:codeSimpleType
	Code string `xml:"code,omitempty"`

	// generated from element "score" of simple type tns:scoreSimpleType
	Score string `xml:"score,omitempty"`
}
