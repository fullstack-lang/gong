// generated code - do not edit
package refs

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "statusEnum"
type Enum_StatusEnum string

const (
	Enum_StatusEnum_ACTIVE Enum_StatusEnum = "ACTIVE"
	Enum_StatusEnum_INACTIVE Enum_StatusEnum = "INACTIVE"
	Enum_StatusEnum_DEPRECATED Enum_StatusEnum = "DEPRECATED"
)

// insertion point for gongstructs declarations

// AttributeGroup_auditAttributes UnNamed source named attribute group "auditAttributes"
type AttributeGroup_auditAttributes struct {

	// insertion point for fields

	// generated from attribute "createdBy"
	CreatedBy string `xml:"createdBy,attr,omitempty"`
}

// RefTargetType Named source named complex type "refTargetType"
type RefTargetType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id"
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "status"
	Status Enum_StatusEnum `xml:"status,attr,omitempty"`

	// generated from attribute group "refs:auditAttributes"
	AttributeGroup_auditAttributes

	// generated from element "label" of simple type xs:string
	Label string `xml:"label,omitempty"`
}
