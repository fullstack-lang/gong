// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test/test/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Astruct:
		astructInstance := any(concreteInstance).(*models.Astruct)
		ret2 := backRepo.BackRepoAstruct.GetAstructDBFromAstructPtr(astructInstance)
		ret = any(ret2).(*T2)
	case *models.AstructBstruct2Use:
		astructbstruct2useInstance := any(concreteInstance).(*models.AstructBstruct2Use)
		ret2 := backRepo.BackRepoAstructBstruct2Use.GetAstructBstruct2UseDBFromAstructBstruct2UsePtr(astructbstruct2useInstance)
		ret = any(ret2).(*T2)
	case *models.AstructBstructUse:
		astructbstructuseInstance := any(concreteInstance).(*models.AstructBstructUse)
		ret2 := backRepo.BackRepoAstructBstructUse.GetAstructBstructUseDBFromAstructBstructUsePtr(astructbstructuseInstance)
		ret = any(ret2).(*T2)
	case *models.Bstruct:
		bstructInstance := any(concreteInstance).(*models.Bstruct)
		ret2 := backRepo.BackRepoBstruct.GetBstructDBFromBstructPtr(bstructInstance)
		ret = any(ret2).(*T2)
	case *models.Dstruct:
		dstructInstance := any(concreteInstance).(*models.Dstruct)
		ret2 := backRepo.BackRepoDstruct.GetDstructDBFromDstructPtr(dstructInstance)
		ret = any(ret2).(*T2)
	case *models.F0123456789012345678901234567890:
		f0123456789012345678901234567890Instance := any(concreteInstance).(*models.F0123456789012345678901234567890)
		ret2 := backRepo.BackRepoF0123456789012345678901234567890.GetF0123456789012345678901234567890DBFromF0123456789012345678901234567890Ptr(f0123456789012345678901234567890Instance)
		ret = any(ret2).(*T2)
	case *models.Gstruct:
		gstructInstance := any(concreteInstance).(*models.Gstruct)
		ret2 := backRepo.BackRepoGstruct.GetGstructDBFromGstructPtr(gstructInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Astruct:
		tmp := GetInstanceDBFromInstance[models.Astruct, AstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AstructBstruct2Use:
		tmp := GetInstanceDBFromInstance[models.AstructBstruct2Use, AstructBstruct2UseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AstructBstructUse:
		tmp := GetInstanceDBFromInstance[models.AstructBstructUse, AstructBstructUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dstruct:
		tmp := GetInstanceDBFromInstance[models.Dstruct, DstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.F0123456789012345678901234567890:
		tmp := GetInstanceDBFromInstance[models.F0123456789012345678901234567890, F0123456789012345678901234567890DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Gstruct:
		tmp := GetInstanceDBFromInstance[models.Gstruct, GstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Astruct:
		tmp := GetInstanceDBFromInstance[models.Astruct, AstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AstructBstruct2Use:
		tmp := GetInstanceDBFromInstance[models.AstructBstruct2Use, AstructBstruct2UseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AstructBstructUse:
		tmp := GetInstanceDBFromInstance[models.AstructBstructUse, AstructBstructUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dstruct:
		tmp := GetInstanceDBFromInstance[models.Dstruct, DstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.F0123456789012345678901234567890:
		tmp := GetInstanceDBFromInstance[models.F0123456789012345678901234567890, F0123456789012345678901234567890DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Gstruct:
		tmp := GetInstanceDBFromInstance[models.Gstruct, GstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
