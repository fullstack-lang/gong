// generated code - do not edit
package gcobase

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AbstractObject_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *AbstractObject_Type) GongCopyBasicFields(to *AbstractObject_Type) {
	// insertion point
	*to = *from
}

type CodeListValue_Type_WOP struct {
	// insertion point

	Name string

	CodeList string

	CodeListValue string

	CodeSpace string

	EnclosedText string
}

func (from *CodeListValue_Type) GongCopyBasicFields(to *CodeListValue_Type) {
	// insertion point
	*to = *from
}

type ObjectReference_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Uuidref string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *ObjectReference_PropertyType) GongCopyBasicFields(to *ObjectReference_PropertyType) {
	// insertion point
	*to = *from
}

// end of insertion point
