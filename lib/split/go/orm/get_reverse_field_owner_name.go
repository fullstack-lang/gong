// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/split/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.AsSplit:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplitArea":
			switch reverseField.Fieldname {
			case "AsSplits":
				if _assplitarea, ok := stage.AsSplitArea_AsSplits_reverseMap[inst]; ok {
					res = _assplitarea.Name
				}
			}
		}

	case *models.AsSplitArea:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplit":
			switch reverseField.Fieldname {
			case "AsSplitAreas":
				if _assplit, ok := stage.AsSplit_AsSplitAreas_reverseMap[inst]; ok {
					res = _assplit.Name
				}
			}
		case "View":
			switch reverseField.Fieldname {
			case "RootAsSplitAreas":
				if _view, ok := stage.View_RootAsSplitAreas_reverseMap[inst]; ok {
					res = _view.Name
				}
			}
		}

	case *models.Form:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.View:
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
	case *models.AsSplit:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplitArea":
			switch reverseField.Fieldname {
			case "AsSplits":
				res = stage.AsSplitArea_AsSplits_reverseMap[inst]
			}
		}

	case *models.AsSplitArea:
		switch reverseField.GongstructName {
		// insertion point
		case "AsSplit":
			switch reverseField.Fieldname {
			case "AsSplitAreas":
				res = stage.AsSplit_AsSplitAreas_reverseMap[inst]
			}
		case "View":
			switch reverseField.Fieldname {
			case "RootAsSplitAreas":
				res = stage.View_RootAsSplitAreas_reverseMap[inst]
			}
		}

	case *models.Form:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.View:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
