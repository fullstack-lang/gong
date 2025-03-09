// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test/test3/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.A:
		aInstance := any(concreteInstance).(*models.A)
		ret2 := backRepo.BackRepoA.GetADBFromAPtr(aInstance)
		ret = any(ret2).(*T2)
	case *models.B:
		bInstance := any(concreteInstance).(*models.B)
		ret2 := backRepo.BackRepoB.GetBDBFromBPtr(bInstance)
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
	case *models.A:
		tmp := GetInstanceDBFromInstance[models.A, ADB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.B:
		tmp := GetInstanceDBFromInstance[models.B, BDB](
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
	case *models.A:
		tmp := GetInstanceDBFromInstance[models.A, ADB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.B:
		tmp := GetInstanceDBFromInstance[models.B, BDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
