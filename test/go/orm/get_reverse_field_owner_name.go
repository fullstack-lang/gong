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
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofa":
				if _astruct, ok := stage.Astruct_Anarrayofa_reverseMap[inst]; ok {
					res = _astruct.Name
				}
			}
		}

	case *models.AstructBstruct2Use:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofb2Use":
				if _astruct, ok := stage.Astruct_Anarrayofb2Use_reverseMap[inst]; ok {
					res = _astruct.Name
				}
			}
		}

	case *models.AstructBstructUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "AnarrayofbUse":
				if _astruct, ok := stage.Astruct_AnarrayofbUse_reverseMap[inst]; ok {
					res = _astruct.Name
				}
			}
		}

	case *models.Bstruct:
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
		case "Dstruct":
			switch reverseField.Fieldname {
			case "Anarrayofb":
				if _dstruct, ok := stage.Dstruct_Anarrayofb_reverseMap[inst]; ok {
					res = _dstruct.Name
				}
			}
		}

	case *models.Dstruct:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Dstruct4s":
				if _astruct, ok := stage.Astruct_Dstruct4s_reverseMap[inst]; ok {
					res = _astruct.Name
				}
			}
		}

	case *models.Fstruct:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Gstruct:
		switch reverseField.GongstructName {
		// insertion point
		case "Dstruct":
			switch reverseField.Fieldname {
			case "Gstructs":
				if _dstruct, ok := stage.Dstruct_Gstructs_reverseMap[inst]; ok {
					res = _dstruct.Name
				}
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
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofa":
				res = stage.Astruct_Anarrayofa_reverseMap[inst]
			}
		}

	case *models.AstructBstruct2Use:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofb2Use":
				res = stage.Astruct_Anarrayofb2Use_reverseMap[inst]
			}
		}

	case *models.AstructBstructUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "AnarrayofbUse":
				res = stage.Astruct_AnarrayofbUse_reverseMap[inst]
			}
		}

	case *models.Bstruct:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Anarrayofb":
				res = stage.Astruct_Anarrayofb_reverseMap[inst]
			case "Anotherarrayofb":
				res = stage.Astruct_Anotherarrayofb_reverseMap[inst]
			}
		case "Dstruct":
			switch reverseField.Fieldname {
			case "Anarrayofb":
				res = stage.Dstruct_Anarrayofb_reverseMap[inst]
			}
		}

	case *models.Dstruct:
		switch reverseField.GongstructName {
		// insertion point
		case "Astruct":
			switch reverseField.Fieldname {
			case "Dstruct4s":
				res = stage.Astruct_Dstruct4s_reverseMap[inst]
			}
		}

	case *models.Fstruct:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Gstruct:
		switch reverseField.GongstructName {
		// insertion point
		case "Dstruct":
			switch reverseField.Fieldname {
			case "Gstructs":
				res = stage.Dstruct_Gstructs_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
