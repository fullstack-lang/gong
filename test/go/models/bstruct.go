package models

// Bstruct demonstrate basic metabaron
//
//	swagger:model Bstruct
type Bstruct struct {
	Name string

	Floatfield, Floatfield2 float64

	Intfield int

	// swagger:ignore this field is ignored by gongc
	Struct interface{} `gorm:"-"` // instruction for gorm ORM

	// swagger:ignore
	// gong:ignore
	ToBeIgnored int
}

// presence of this function will let gong generates orchestrator code
func (bstruct *Bstruct) OnAfterUpdate(stage *StageStruct, stagedInstance, frontInstance *Bstruct) {

	//
}
