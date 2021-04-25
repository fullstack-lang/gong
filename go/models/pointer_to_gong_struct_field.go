package models

// PointerToGongStructField
// swagger:model
type PointerToGongStructField struct {
	Name       string
	GongStruct *GongStruct
}

func (p *PointerToGongStructField) GetName() string {
	return p.Name
}
