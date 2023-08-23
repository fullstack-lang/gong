package orm

import (
	"github.com/fullstack-lang/gong/test/go/models"
)

type GongstrucDB interface {
	// insertion point for generic types
	AstructDB | AstructBstruct2UseDB | AstructBstructUseDB | BstructDB | DstructDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstrucDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	case *models.Astruct:
		astructInstance := any(concreteInstance).(*models.Astruct)
		ret2 := backRepo.BackRepoAstruct.GetAstructDBFromAstructPtr(astructInstance)
		ret = any(ret2).(*T2)
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	case *models.Astruct:
		tmp := GetInstanceDBFromInstance[models.Astruct, AstructDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	}
	return
}
