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
				if _astruct, ok := stage.Astruct_Anarrayofa_reverseMap[inst]; ok {
					res = _astruct.Name
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
		case "Fstruct":
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
				if _astruct, ok := stage.Astruct_Anarrayofb2Use_reverseMap[inst]; ok {
					res = _astruct.Name
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
		case "Fstruct":
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
				if _astruct, ok := stage.Astruct_AnarrayofbUse_reverseMap[inst]; ok {
					res = _astruct.Name
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
		case "Fstruct":
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
				if _astruct, ok := stage.Astruct_Anarrayofb_reverseMap[inst]; ok {
					res = _astruct.Name
				}
			case "Anotherarrayofb":
				if _astruct, ok := stage.Astruct_Anotherarrayofb_reverseMap[inst]; ok {
					res = _astruct.Name
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
				if _dstruct, ok := stage.Dstruct_Anarrayofb_reverseMap[inst]; ok {
					res = _dstruct.Name
				}
			}
		case "Fstruct":
			switch reverseField.Fieldname {
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
		case "Fstruct":
			switch reverseField.Fieldname {
			}
		}

	case *models.Fstruct:
		tmp := GetInstanceDBFromInstance[models.Fstruct, FstructDB](
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
		case "Fstruct":
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
				res = stage.Astruct_Anarrayofa_reverseMap[inst]
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
		case "Fstruct":
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
				res = stage.Astruct_Anarrayofb2Use_reverseMap[inst]
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
		case "Fstruct":
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
				res = stage.Astruct_AnarrayofbUse_reverseMap[inst]
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
		case "Fstruct":
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
				res = stage.Astruct_Anarrayofb_reverseMap[inst]
			case "Anotherarrayofb":
				res = stage.Astruct_Anotherarrayofb_reverseMap[inst]
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
				res = stage.Dstruct_Anarrayofb_reverseMap[inst]
			}
		case "Fstruct":
			switch reverseField.Fieldname {
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
		case "Fstruct":
			switch reverseField.Fieldname {
			}
		}

	case *models.Fstruct:
		tmp := GetInstanceDBFromInstance[models.Fstruct, FstructDB](
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
		case "Fstruct":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
