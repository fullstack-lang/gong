// generated code - do not edit
package x

import (
	"github.com/fullstack-lang/gong/test2/go/models/x"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | ADB
}

func GetInstanceDBFromInstance[T x.Gongstruct, T2 GongstructDB](
	stage *x.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *x.A:
		aInstance := any(concreteInstance).(*x.A)
		ret2 := backRepo.BackRepoA.GetADBFromAPtr(aInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T x.Gongstruct](
	stage *x.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *x.A:
		tmp := GetInstanceDBFromInstance[x.A, ADB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T x.PointerToGongstruct](
	stage *x.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *x.A:
		tmp := GetInstanceDBFromInstance[x.A, ADB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
