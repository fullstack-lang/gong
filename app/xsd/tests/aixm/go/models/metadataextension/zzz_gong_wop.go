// generated code - do not edit
package metadataextension

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type MD_DatatypeCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *MD_DatatypeCode_PropertyType) GongCopyBasicFields(to *MD_DatatypeCode_PropertyType) {
	// insertion point
	*to = *from
}

type MD_ExtendedElementInformation_PropertyType_WOP struct {
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

func (from *MD_ExtendedElementInformation_PropertyType) GongCopyBasicFields(to *MD_ExtendedElementInformation_PropertyType) {
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

type MD_ExtendedElementInformation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_ExtendedElementInformation_Type) GongCopyBasicFields(to *MD_ExtendedElementInformation_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_MetadataExtensionInformation_PropertyType_WOP struct {
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

func (from *MD_MetadataExtensionInformation_PropertyType) GongCopyBasicFields(to *MD_MetadataExtensionInformation_PropertyType) {
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

type MD_MetadataExtensionInformation_Type_WOP struct {
	// insertion point

	Name string

	Id string

	Uuid string
}

func (from *MD_MetadataExtensionInformation_Type) GongCopyBasicFields(to *MD_MetadataExtensionInformation_Type) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Uuid = from.Uuid
}

type MD_ObligationCode_PropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	MD_ObligationCode Enum_MD_ObligationCode_Type
}

func (from *MD_ObligationCode_PropertyType) GongCopyBasicFields(to *MD_ObligationCode_PropertyType) {
	// insertion point
	*to = *from
}

// end of insertion point
