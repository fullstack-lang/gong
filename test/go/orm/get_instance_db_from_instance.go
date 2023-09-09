// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int  | AstructDB | AstructBstruct2UseDB | AstructBstructUseDB | BstructDB | DstructDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
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
	default:
		_ = inst
	}
	return
}
