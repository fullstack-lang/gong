// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/button/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Button:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Buttons":
				if _group, ok := stage.Group_Buttons_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "Groups":
				if _layout, ok := stage.Layout_Groups_reverseMap[inst]; ok {
					res = _layout.Name
				}
			}
		}

	case *models.Layout:
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
	case *models.Button:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Buttons":
				res = stage.Group_Buttons_reverseMap[inst]
			}
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Layout_Groups_reverseMap[inst]
			}
		}

	case *models.Layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
