// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/cursor/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Cursor:
		cursorInstance := any(concreteInstance).(*models.Cursor)
		ret2 := backRepo.BackRepoCursor.GetCursorDBFromCursorPtr(cursorInstance)
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
	case *models.Cursor:
		tmp := GetInstanceDBFromInstance[models.Cursor, CursorDB](
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
	case *models.Cursor:
		tmp := GetInstanceDBFromInstance[models.Cursor, CursorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
