// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/tone/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Freqency:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Frequencies":
				if _note, ok := stage.Note_Frequencies_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *models.Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Player:
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
	case *models.Freqency:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Frequencies":
				res = stage.Note_Frequencies_reverseMap[inst]
			}
		}

	case *models.Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Player:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
