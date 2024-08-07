// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.GongBasicField:
		gongbasicfieldInstance := any(concreteInstance).(*models.GongBasicField)
		ret2 := backRepo.BackRepoGongBasicField.GetGongBasicFieldDBFromGongBasicFieldPtr(gongbasicfieldInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnum:
		gongenumInstance := any(concreteInstance).(*models.GongEnum)
		ret2 := backRepo.BackRepoGongEnum.GetGongEnumDBFromGongEnumPtr(gongenumInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnumValue:
		gongenumvalueInstance := any(concreteInstance).(*models.GongEnumValue)
		ret2 := backRepo.BackRepoGongEnumValue.GetGongEnumValueDBFromGongEnumValuePtr(gongenumvalueInstance)
		ret = any(ret2).(*T2)
	case *models.GongLink:
		gonglinkInstance := any(concreteInstance).(*models.GongLink)
		ret2 := backRepo.BackRepoGongLink.GetGongLinkDBFromGongLinkPtr(gonglinkInstance)
		ret = any(ret2).(*T2)
	case *models.GongNote:
		gongnoteInstance := any(concreteInstance).(*models.GongNote)
		ret2 := backRepo.BackRepoGongNote.GetGongNoteDBFromGongNotePtr(gongnoteInstance)
		ret = any(ret2).(*T2)
	case *models.GongStruct:
		gongstructInstance := any(concreteInstance).(*models.GongStruct)
		ret2 := backRepo.BackRepoGongStruct.GetGongStructDBFromGongStructPtr(gongstructInstance)
		ret = any(ret2).(*T2)
	case *models.GongTimeField:
		gongtimefieldInstance := any(concreteInstance).(*models.GongTimeField)
		ret2 := backRepo.BackRepoGongTimeField.GetGongTimeFieldDBFromGongTimeFieldPtr(gongtimefieldInstance)
		ret = any(ret2).(*T2)
	case *models.Meta:
		metaInstance := any(concreteInstance).(*models.Meta)
		ret2 := backRepo.BackRepoMeta.GetMetaDBFromMetaPtr(metaInstance)
		ret = any(ret2).(*T2)
	case *models.MetaReference:
		metareferenceInstance := any(concreteInstance).(*models.MetaReference)
		ret2 := backRepo.BackRepoMetaReference.GetMetaReferenceDBFromMetaReferencePtr(metareferenceInstance)
		ret = any(ret2).(*T2)
	case *models.ModelPkg:
		modelpkgInstance := any(concreteInstance).(*models.ModelPkg)
		ret2 := backRepo.BackRepoModelPkg.GetModelPkgDBFromModelPkgPtr(modelpkgInstance)
		ret = any(ret2).(*T2)
	case *models.PointerToGongStructField:
		pointertogongstructfieldInstance := any(concreteInstance).(*models.PointerToGongStructField)
		ret2 := backRepo.BackRepoPointerToGongStructField.GetPointerToGongStructFieldDBFromPointerToGongStructFieldPtr(pointertogongstructfieldInstance)
		ret = any(ret2).(*T2)
	case *models.SliceOfPointerToGongStructField:
		sliceofpointertogongstructfieldInstance := any(concreteInstance).(*models.SliceOfPointerToGongStructField)
		ret2 := backRepo.BackRepoSliceOfPointerToGongStructField.GetSliceOfPointerToGongStructFieldDBFromSliceOfPointerToGongStructFieldPtr(sliceofpointertogongstructfieldInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.GongBasicField:
		tmp := GetInstanceDBFromInstance[models.GongBasicField, GongBasicFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnum:
		tmp := GetInstanceDBFromInstance[models.GongEnum, GongEnumDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValue:
		tmp := GetInstanceDBFromInstance[models.GongEnumValue, GongEnumValueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongLink:
		tmp := GetInstanceDBFromInstance[models.GongLink, GongLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNote:
		tmp := GetInstanceDBFromInstance[models.GongNote, GongNoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStruct:
		tmp := GetInstanceDBFromInstance[models.GongStruct, GongStructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongTimeField:
		tmp := GetInstanceDBFromInstance[models.GongTimeField, GongTimeFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Meta:
		tmp := GetInstanceDBFromInstance[models.Meta, MetaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MetaReference:
		tmp := GetInstanceDBFromInstance[models.MetaReference, MetaReferenceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ModelPkg:
		tmp := GetInstanceDBFromInstance[models.ModelPkg, ModelPkgDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.PointerToGongStructField, PointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SliceOfPointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.SliceOfPointerToGongStructField, SliceOfPointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.GongBasicField:
		tmp := GetInstanceDBFromInstance[models.GongBasicField, GongBasicFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnum:
		tmp := GetInstanceDBFromInstance[models.GongEnum, GongEnumDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValue:
		tmp := GetInstanceDBFromInstance[models.GongEnumValue, GongEnumValueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongLink:
		tmp := GetInstanceDBFromInstance[models.GongLink, GongLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNote:
		tmp := GetInstanceDBFromInstance[models.GongNote, GongNoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStruct:
		tmp := GetInstanceDBFromInstance[models.GongStruct, GongStructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongTimeField:
		tmp := GetInstanceDBFromInstance[models.GongTimeField, GongTimeFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Meta:
		tmp := GetInstanceDBFromInstance[models.Meta, MetaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MetaReference:
		tmp := GetInstanceDBFromInstance[models.MetaReference, MetaReferenceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ModelPkg:
		tmp := GetInstanceDBFromInstance[models.ModelPkg, ModelPkgDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.PointerToGongStructField, PointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SliceOfPointerToGongStructField:
		tmp := GetInstanceDBFromInstance[models.SliceOfPointerToGongStructField, SliceOfPointerToGongStructFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
