// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test2/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Dummy:
		tmp := GetInstanceDBFromInstance[models.Dummy, DummyDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}
