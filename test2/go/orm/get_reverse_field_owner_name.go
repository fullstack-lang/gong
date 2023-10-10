// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test2/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.A:
		tmp := GetInstanceDBFromInstance[models.A, ADB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}

	case *models.B:
		tmp := GetInstanceDBFromInstance[models.B, BDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				if tmp != nil && tmp.A_BsDBID.Int64 != 0 {
					id := uint(tmp.A_BsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoA.Map_ADBID_APtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "B":
			switch reverseField.Fieldname {
			}
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
	case *models.A:
		tmp := GetInstanceDBFromInstance[models.A, ADB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.B:
		tmp := GetInstanceDBFromInstance[models.B, BDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				if tmp != nil && tmp.A_BsDBID.Int64 != 0 {
					id := uint(tmp.A_BsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoA.Map_ADBID_APtr[id]
					res = reservePointerTarget
				}
			}
		case "B":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
