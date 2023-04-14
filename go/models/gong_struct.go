package models

// GongStruct is a go struct that is selected by the gongc compiler
// swagger:model
type GongStruct struct {
	Name string

	// Fields (stored according to the source file order)
	// swagger:ignore
	Fields []FieldInterface `gorm:"-"`

	// Slice of Fields by their type (not populated by the gongc)
	GongBasicFields                  []*GongBasicField
	GongTimeFields                   []*GongTimeField
	PointerToGongStructFields        []*PointerToGongStructField
	SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField

	// HasOnAfterUpdateSignature is used to generate orchestrator code
	HasOnAfterUpdateSignature bool
}

// HasNameField indicates wether the gong struct has a field with Name "Name"
//
// This is important since
// - only Gong Struct with Name field can be stored in DB
// - only Gong Struct without file
func (gongStruct *GongStruct) HasNameField() (hasNameField bool) {

	// hasNameField default value is false
	for _, field := range gongStruct.Fields {
		switch field := field.(type) {
		case *GongBasicField:
			if field.Name == "Name" {
				hasNameField = true
				continue
			}
		}
	}

	return
}
