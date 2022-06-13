package models

// SliceOfPointerToGongStructField
// swagger:model
type SliceOfPointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int

	CompositeStructName string
}

func (sliceOfPointerToGongStructField *SliceOfPointerToGongStructField) GetIndex() int {
	return sliceOfPointerToGongStructField.Index
}

func (sliceOfPointerToGongStructField *SliceOfPointerToGongStructField) GetCompositeStructName() string {
	return sliceOfPointerToGongStructField.CompositeStructName
}
