package models

// ModelPkg is the go package where the gong source code is located
type ModelPkg struct {
	Name    string // should be "models"
	PkgPath string // for instance "github.com/.../models"

	GongStructs map[string]*GongStruct `gorm:"-"` // sql3Lite does not support maps
	GongEnums   map[string]*GongEnum   `gorm:"-"`
}

// SerializeToStage stages modelPkg and
// recursively stage all structs and all fields of all structs
func (modelPkg *ModelPkg) SerializeToStage() {
	modelPkg.Stage()
	for _, gongStruct := range modelPkg.GongStructs {
		gongStruct.Stage()

		for _, field := range gongStruct.Fields {
			switch field.(type) {
			case *GongBasicField:
				gongBasicField := field.(*GongBasicField)
				_ = gongBasicField
				gongBasicField.Stage()
				gongStruct.GongBasicFields = append(gongStruct.GongBasicFields, gongBasicField)

			case *GongTimeField:
				gongTimeField := field.(*GongTimeField)
				_ = gongTimeField
				gongTimeField.Stage()
				gongStruct.GongTimeFields = append(gongStruct.GongTimeFields, gongTimeField)

			case *PointerToGongStructField:
				pointerToStructField := field.(*PointerToGongStructField)
				_ = pointerToStructField
				pointerToStructField.Stage()
				gongStruct.PointerToGongStructFields = append(gongStruct.PointerToGongStructFields, pointerToStructField)

			case *SliceOfPointerToGongStructField:
				sliceOfPointerToGongStruct := field.(*SliceOfPointerToGongStructField)
				_ = sliceOfPointerToGongStruct
				sliceOfPointerToGongStruct.Stage()
				gongStruct.SliceOfPointerToGongStructFields = append(gongStruct.SliceOfPointerToGongStructFields,
					sliceOfPointerToGongStruct)
			}
		}

	}
	for _, gongEnum := range modelPkg.GongEnums {
		gongEnum.Stage()

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			gongEnumValue.Stage()
		}
	}
	Stage.Commit()
}
