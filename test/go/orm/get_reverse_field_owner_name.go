// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		tmp := GetInstanceDBFromInstance[models.Astruct, AstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.AstructBstruct2Use:
		tmp := GetInstanceDBFromInstance[models.AstructBstruct2Use, AstructBstruct2UseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.AstructBstructUse:
		tmp := GetInstanceDBFromInstance[models.AstructBstructUse, AstructBstructUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.Dstruct:
		tmp := GetInstanceDBFromInstance[models.Dstruct, DstructDB](
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
