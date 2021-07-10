package models

// SliceOfPointerToGongStructField
// swagger:model
type SliceOfPointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct

	Index int
}

func (s *SliceOfPointerToGongStructField) GetName() string {
	return s.Name
}

func (sliceOfPointerToGongStructField *SliceOfPointerToGongStructField) GetIndex() int {
	return sliceOfPointerToGongStructField.Index
}
