// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Command:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DummyAgent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Status:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UpdateState:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Command:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DummyAgent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Status:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UpdateState:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
