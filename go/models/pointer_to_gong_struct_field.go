package models

// PointerToGongStructField
// swagger:model
type PointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int

	CompositeStructName string
}

func (pointerToGongStructField *PointerToGongStructField) GetIndex() int {
	return pointerToGongStructField.Index
}

func (pointerToGongStructField *PointerToGongStructField) GetCompositeStructName() string {
	return pointerToGongStructField.CompositeStructName
}
