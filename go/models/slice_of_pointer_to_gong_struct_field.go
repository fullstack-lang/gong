package models

// SliceOfPointerToGongStructField
// swagger:model
type SliceOfPointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int
}

func (sliceOfPointerToGongStructField *SliceOfPointerToGongStructField) GetIndex() int {
	return sliceOfPointerToGongStructField.Index
}
