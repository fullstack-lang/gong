// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongBasicFields":
				if _gongstruct, ok := stage.GongStruct_GongBasicFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		}

	case *models.GongEnum:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongEnumValue:
		switch reverseField.GongstructName {
		// insertion point
		case "GongEnum":
			switch reverseField.Fieldname {
			case "GongEnumValues":
				if _gongenum, ok := stage.GongEnum_GongEnumValues_reverseMap[inst]; ok {
					res = _gongenum.Name
				}
			}
		}

	case *models.GongLink:
		switch reverseField.GongstructName {
		// insertion point
		case "GongNote":
			switch reverseField.Fieldname {
			case "Links":
				if _gongnote, ok := stage.GongNote_Links_reverseMap[inst]; ok {
					res = _gongnote.Name
				}
			}
		}

	case *models.GongNote:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongStruct:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongTimeField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongTimeFields":
				if _gongstruct, ok := stage.GongStruct_GongTimeFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		}

	case *models.Meta:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MetaReference:
		switch reverseField.GongstructName {
		// insertion point
		case "Meta":
			switch reverseField.Fieldname {
			case "MetaReferences":
				if _meta, ok := stage.Meta_MetaReferences_reverseMap[inst]; ok {
					res = _meta.Name
				}
			}
		}

	case *models.ModelPkg:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.PointerToGongStructField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "PointerToGongStructFields":
				if _gongstruct, ok := stage.GongStruct_PointerToGongStructFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
				}
			}
		}

	case *models.SliceOfPointerToGongStructField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "SliceOfPointerToGongStructFields":
				if _gongstruct, ok := stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]; ok {
					res = _gongstruct.Name
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
	case *models.GongBasicField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongBasicFields":
				res = stage.GongStruct_GongBasicFields_reverseMap[inst]
			}
		}

	case *models.GongEnum:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongEnumValue:
		switch reverseField.GongstructName {
		// insertion point
		case "GongEnum":
			switch reverseField.Fieldname {
			case "GongEnumValues":
				res = stage.GongEnum_GongEnumValues_reverseMap[inst]
			}
		}

	case *models.GongLink:
		switch reverseField.GongstructName {
		// insertion point
		case "GongNote":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.GongNote_Links_reverseMap[inst]
			}
		}

	case *models.GongNote:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongStruct:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongTimeField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "GongTimeFields":
				res = stage.GongStruct_GongTimeFields_reverseMap[inst]
			}
		}

	case *models.Meta:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MetaReference:
		switch reverseField.GongstructName {
		// insertion point
		case "Meta":
			switch reverseField.Fieldname {
			case "MetaReferences":
				res = stage.Meta_MetaReferences_reverseMap[inst]
			}
		}

	case *models.ModelPkg:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.PointerToGongStructField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "PointerToGongStructFields":
				res = stage.GongStruct_PointerToGongStructFields_reverseMap[inst]
			}
		}

	case *models.SliceOfPointerToGongStructField:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStruct":
			switch reverseField.Fieldname {
			case "SliceOfPointerToGongStructFields":
				res = stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
