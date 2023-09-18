// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		tmp := GetInstanceDBFromInstance[models.GongBasicField, GongBasicFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongBasicFields":
			if tmp.GongStruct_GongBasicFieldsDBID.Int64 != 0 {
				id := uint(tmp.GongStruct_GongBasicFieldsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongEnum:
		tmp := GetInstanceDBFromInstance[models.GongEnum, GongEnumDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.GongEnumValue:
		tmp := GetInstanceDBFromInstance[models.GongEnumValue, GongEnumValueDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongEnumValues":
			if tmp.GongEnum_GongEnumValuesDBID.Int64 != 0 {
				id := uint(tmp.GongEnum_GongEnumValuesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongLink:
		tmp := GetInstanceDBFromInstance[models.GongLink, GongLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Links":
			if tmp.GongNote_LinksDBID.Int64 != 0 {
				id := uint(tmp.GongNote_LinksDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongNote:
		tmp := GetInstanceDBFromInstance[models.GongNote, GongNoteDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.GongStruct:
		tmp := GetInstanceDBFromInstance[models.GongStruct, GongStructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.GongTimeField:
		tmp := GetInstanceDBFromInstance[models.GongTimeField, GongTimeFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongTimeFields":
			if tmp.GongStruct_GongTimeFieldsDBID.Int64 != 0 {
				id := uint(tmp.GongStruct_GongTimeFieldsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Meta:
		tmp := GetInstanceDBFromInstance[models.Meta, MetaDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.MetaReference:
		tmp := GetInstanceDBFromInstance[models.MetaReference, MetaReferenceDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "MetaReferences":
			if tmp.Meta_MetaReferencesDBID.Int64 != 0 {
				id := uint(tmp.Meta_MetaReferencesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.ModelPkg:
		tmp := GetInstanceDBFromInstance[models.ModelPkg, ModelPkgDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.PointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.PointerToGongStructField, PointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "PointerToGongStructFields":
			if tmp.GongStruct_PointerToGongStructFieldsDBID.Int64 != 0 {
				id := uint(tmp.GongStruct_PointerToGongStructFieldsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.SliceOfPointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.SliceOfPointerToGongStructField, SliceOfPointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "SliceOfPointerToGongStructFields":
			if tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 != 0 {
				id := uint(tmp.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]
				res = reservePointerTarget.Name
			}
		}

	default:
		_ = inst
	}
	return
}
