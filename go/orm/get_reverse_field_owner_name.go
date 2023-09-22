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
				if tmp.GongStruct_GongBasicFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_GongBasicFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongEnum_GongEnumValuesDBID.Int64 != 0 {
					id := uint(tmp.GongEnum_GongEnumValuesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongNote_LinksDBID.Int64 != 0 {
					id := uint(tmp.GongNote_LinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongStruct_GongTimeFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_GongTimeFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.Meta_MetaReferencesDBID.Int64 != 0 {
					id := uint(tmp.Meta_MetaReferencesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongStruct_PointerToGongStructFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_PointerToGongStructFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget.Name
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
				if tmp.GongStruct_GongBasicFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_GongBasicFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget
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
				if tmp.GongEnum_GongEnumValuesDBID.Int64 != 0 {
					id := uint(tmp.GongEnum_GongEnumValuesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[id]
					res = reservePointerTarget
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
				if tmp.GongNote_LinksDBID.Int64 != 0 {
					id := uint(tmp.GongNote_LinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[id]
					res = reservePointerTarget
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
				if tmp.GongStruct_GongTimeFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_GongTimeFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget
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
				if tmp.Meta_MetaReferencesDBID.Int64 != 0 {
					id := uint(tmp.Meta_MetaReferencesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[id]
					res = reservePointerTarget
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
				if tmp.GongStruct_PointerToGongStructFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_PointerToGongStructFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget
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
				if tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
					res = reservePointerTarget
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
	return res
}
