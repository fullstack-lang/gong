package models

// ModelPkg is the go package where the gong source code is located
//
// It contains the list of GongStructs & GongEnum
type ModelPkg struct {
	Name    string // should be "models"
	PkgPath string // for instance "github.com/.../models"

	GongStructs map[string]*GongStruct `gorm:"-"` // sql3Lite does not support maps
	GongEnums   map[string]*GongEnum   `gorm:"-"`
	GongNotes   map[string]*GongNote   `gorm:"-"`
}

// SerializeToStage stages modelPkg and
// recursively stage all structs and all fields of all structs
func (modelPkg *ModelPkg) SerializeToStage() {
	modelPkg.Stage()
	for _, gongStruct := range modelPkg.GongStructs {
		gongStruct.Stage()

		for _, field := range gongStruct.Fields {
			switch field := field.(type) {
			case *GongBasicField:

				field.Stage()
				gongStruct.GongBasicFields = append(gongStruct.GongBasicFields, field)

			case *GongTimeField:
				field.Stage()
				gongStruct.GongTimeFields = append(gongStruct.GongTimeFields, field)

			case *PointerToGongStructField:
				field.Stage()
				gongStruct.PointerToGongStructFields = append(gongStruct.PointerToGongStructFields, field)

			case *SliceOfPointerToGongStructField:
				field.Stage()
				gongStruct.SliceOfPointerToGongStructFields = append(gongStruct.SliceOfPointerToGongStructFields,
					field)
			}
		}

	}
	for _, gongEnum := range modelPkg.GongEnums {
		gongEnum.Stage()

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			gongEnumValue.Stage()
		}
	}
	for _, gongNote := range modelPkg.GongNotes {
		gongNote.Stage()
	}
	Stage.Commit()
}
