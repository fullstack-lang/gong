package models

// PointerToGongStructField
// swagger:model
type PointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int
}

func (p *PointerToGongStructField) GetName() string {
	return p.Name
}

func (pointerToGongStructField *PointerToGongStructField) GetIndex() int {
	return pointerToGongStructField.Index
}
