// generated code - do not edit
package applicationschema

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type MD_ApplicationSchemaInformation_PropertyType_WOP struct {
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

func (from *MD_ApplicationSchemaInformation_PropertyType) GongCopyBasicFields(to *MD_ApplicationSchemaInformation_PropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Uuidref = from.Uuidref
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

type MD_ApplicationSchemaInformation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_ApplicationSchemaInformation_Type) GongCopyBasicFields(to *MD_ApplicationSchemaInformation_Type) {
	// insertion point
	*to = *from
}

// end of insertion point
