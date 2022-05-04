package models

// PointerToGongStructField
// swagger:model
type PointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int
}

func (pointerToGongStructField *PointerToGongStructField) GetIndex() int {
	return pointerToGongStructField.Index
}
