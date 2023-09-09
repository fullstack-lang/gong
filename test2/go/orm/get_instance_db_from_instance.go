// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test2/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int  | DummyDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Dummy:
		dummyInstance := any(concreteInstance).(*models.Dummy)
		ret2 := backRepo.BackRepoDummy.GetDummyDBFromDummyPtr(dummyInstance)
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
	case *models.Dummy:
		tmp := GetInstanceDBFromInstance[models.Dummy, DummyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
