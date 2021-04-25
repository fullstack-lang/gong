package models

// SliceOfPointerToGongStructField
// swagger:model
type SliceOfPointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct
}

func (s *SliceOfPointerToGongStructField) GetName() string {
	return s.Name
}
