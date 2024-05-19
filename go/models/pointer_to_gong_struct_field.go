package models

// PointerToGongStructField
// swagger:model
type PointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int

	CompositeStructName string

	// IsType denotes wether this association is to a type
	//
	// a pointer to a gongstruct with "gong:type" magic node means that the object points to a type
	// This means that when a branch is copied, the duplicate meachanism
	// will stop here and only duplicate the pointer, not the type
	IsType bool
}

func (pointerToGongStructField *PointerToGongStructField) GetIndex() int {
	return pointerToGongStructField.Index
}

func (pointerToGongStructField *PointerToGongStructField) GetCompositeStructName() string {
	return pointerToGongStructField.CompositeStructName
}
