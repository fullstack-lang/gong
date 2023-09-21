// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/test/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		tmp := GetInstanceDBFromInstance[models.Astruct, AstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofa":
				if tmp.Astruct_AnarrayofaDBID.Int64 != 0 {
					id := uint(tmp.Astruct_AnarrayofaDBID.Int64)
					reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "AstructBstruct2Use":
			switch reverseField.Fieldname {
			}
		case "AstructBstructUse":
			switch reverseField.Fieldname {
			}
		case "Bstruct":
			switch reverseField.Fieldname {
			}
		case "Dstruct":
			switch reverseField.Fieldname {
			}
		}

	case *models.AstructBstruct2Use:
		tmp := GetInstanceDBFromInstance[models.AstructBstruct2Use, AstructBstruct2UseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofb2Use":
				if tmp.Astruct_Anarrayofb2UseDBID.Int64 != 0 {
					id := uint(tmp.Astruct_Anarrayofb2UseDBID.Int64)
					reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "AstructBstruct2Use":
			switch reverseField.Fieldname {
			}
		case "AstructBstructUse":
			switch reverseField.Fieldname {
			}
		case "Bstruct":
			switch reverseField.Fieldname {
			}
		case "Dstruct":
			switch reverseField.Fieldname {
			}
		}

	case *models.AstructBstructUse:
		tmp := GetInstanceDBFromInstance[models.AstructBstructUse, AstructBstructUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "AnarrayofbUse":
				if tmp.Astruct_AnarrayofbUseDBID.Int64 != 0 {
					id := uint(tmp.Astruct_AnarrayofbUseDBID.Int64)
					reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "AstructBstruct2Use":
			switch reverseField.Fieldname {
			}
		case "AstructBstructUse":
			switch reverseField.Fieldname {
			}
		case "Bstruct":
			switch reverseField.Fieldname {
			}
		case "Dstruct":
			switch reverseField.Fieldname {
			}
		}

	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofb":
				if tmp.Astruct_AnarrayofbDBID.Int64 != 0 {
					id := uint(tmp.Astruct_AnarrayofbDBID.Int64)
					reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
					res = reservePointerTarget.Name
				}
			case "Anotherarrayofb":
				if tmp.Astruct_AnotherarrayofbDBID.Int64 != 0 {
					id := uint(tmp.Astruct_AnotherarrayofbDBID.Int64)
					reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "AstructBstruct2Use":
			switch reverseField.Fieldname {
			}
		case "AstructBstructUse":
			switch reverseField.Fieldname {
			}
		case "Bstruct":
			switch reverseField.Fieldname {
			}
		case "Dstruct":
			switch reverseField.Fieldname {
			case "Anarrayofb":
				if tmp.Dstruct_AnarrayofbDBID.Int64 != 0 {
					id := uint(tmp.Dstruct_AnarrayofbDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDstruct.Map_DstructDBID_DstructPtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.Dstruct:
		tmp := GetInstanceDBFromInstance[models.Dstruct, DstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			}
		case "AstructBstruct2Use":
			switch reverseField.Fieldname {
			}
		case "AstructBstructUse":
			switch reverseField.Fieldname {
			}
		case "Bstruct":
			switch reverseField.Fieldname {
			}
		case "Dstruct":
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
	reverseFieldName string) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Bstruct:
		tmp := GetInstanceDBFromInstance[models.Bstruct, BstructDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Anarrayofb":
			if tmp.Astruct_AnarrayofbDBID.Int64 != 0 {
				id := uint(tmp.Astruct_AnarrayofbDBID.Int64)
				reservePointerTarget := backRepo.BackRepoAstruct.Map_AstructDBID_AstructPtr[id]
				res = reservePointerTarget
			}
		}
	default:
		_ = inst
	}
	return res
}
