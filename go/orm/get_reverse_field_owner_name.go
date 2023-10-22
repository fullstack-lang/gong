// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		tmp := GetInstanceDBFromInstance[models.GongBasicField, GongBasicFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongBasicFields":
				if _gongstruct, ok := stage.GongStruct_GongBasicFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnum:
		tmp := GetInstanceDBFromInstance[models.GongEnum, GongEnumDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumValue:
		tmp := GetInstanceDBFromInstance[models.GongEnumValue, GongEnumValueDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			case "GongEnumValues":
				if _gongenum, ok := stage.GongEnum_GongEnumValues_reverseMap[inst]; ok {
					res = _gongenum.Name
				}
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongLink:
		tmp := GetInstanceDBFromInstance[models.GongLink, GongLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			case "Links":
				if _gongnote, ok := stage.GongNote_Links_reverseMap[inst]; ok {
					res = _gongnote.Name
				}
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongNote:
		tmp := GetInstanceDBFromInstance[models.GongNote, GongNoteDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongStruct:
		tmp := GetInstanceDBFromInstance[models.GongStruct, GongStructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongTimeField:
		tmp := GetInstanceDBFromInstance[models.GongTimeField, GongTimeFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongTimeFields":
				if _gongstruct, ok := stage.GongStruct_GongTimeFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.Meta:
		tmp := GetInstanceDBFromInstance[models.Meta, MetaDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.MetaReference:
		tmp := GetInstanceDBFromInstance[models.MetaReference, MetaReferenceDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			case "MetaReferences":
				if _meta, ok := stage.Meta_MetaReferences_reverseMap[inst]; ok {
					res = _meta.Name
				}
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.ModelPkg:
		tmp := GetInstanceDBFromInstance[models.ModelPkg, ModelPkgDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.PointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.PointerToGongStructField, PointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "PointerToGongStructFields":
				if _gongstruct, ok := stage.GongStruct_PointerToGongStructFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.SliceOfPointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.SliceOfPointerToGongStructField, SliceOfPointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "SliceOfPointerToGongStructFields":
				if _gongstruct, ok := stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		tmp := GetInstanceDBFromInstance[models.GongBasicField, GongBasicFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongBasicFields":
				res = stage.GongStruct_GongBasicFields_reverseMap[inst]
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnum:
		tmp := GetInstanceDBFromInstance[models.GongEnum, GongEnumDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumValue:
		tmp := GetInstanceDBFromInstance[models.GongEnumValue, GongEnumValueDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			case "GongEnumValues":
				res = stage.GongEnum_GongEnumValues_reverseMap[inst]
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongLink:
		tmp := GetInstanceDBFromInstance[models.GongLink, GongLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.GongNote_Links_reverseMap[inst]
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongNote:
		tmp := GetInstanceDBFromInstance[models.GongNote, GongNoteDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongStruct:
		tmp := GetInstanceDBFromInstance[models.GongStruct, GongStructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongTimeField:
		tmp := GetInstanceDBFromInstance[models.GongTimeField, GongTimeFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongTimeFields":
				res = stage.GongStruct_GongTimeFields_reverseMap[inst]
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.Meta:
		tmp := GetInstanceDBFromInstance[models.Meta, MetaDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.MetaReference:
		tmp := GetInstanceDBFromInstance[models.MetaReference, MetaReferenceDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			case "MetaReferences":
				res = stage.Meta_MetaReferences_reverseMap[inst]
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.ModelPkg:
		tmp := GetInstanceDBFromInstance[models.ModelPkg, ModelPkgDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.PointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.PointerToGongStructField, PointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "PointerToGongStructFields":
				res = stage.GongStruct_PointerToGongStructFields_reverseMap[inst]
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	case *models.SliceOfPointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.SliceOfPointerToGongStructField, SliceOfPointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "GongBasicField":
			switch reverseField.Fieldname {
			}
		case "GongEnum":
			switch reverseField.Fieldname {
			}
		case "GongEnumValue":
			switch reverseField.Fieldname {
			}
		case "GongLink":
			switch reverseField.Fieldname {
			}
		case "GongNote":
			switch reverseField.Fieldname {
			}
		case "GongStruct":
			switch reverseField.Fieldname {
			case "SliceOfPointerToGongStructFields":
				res = stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]
			}
		case "GongTimeField":
			switch reverseField.Fieldname {
			}
		case "Meta":
			switch reverseField.Fieldname {
			}
		case "MetaReference":
			switch reverseField.Fieldname {
			}
		case "ModelPkg":
			switch reverseField.Fieldname {
			}
		case "PointerToGongStructField":
			switch reverseField.Fieldname {
			}
		case "SliceOfPointerToGongStructField":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
