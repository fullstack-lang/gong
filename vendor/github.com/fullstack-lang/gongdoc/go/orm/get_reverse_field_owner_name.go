// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdoc/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		switch reverseField.GongstructName {
		// insertion point
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Classdiagrams":
				if _diagrampackage, ok := stage.DiagramPackage_Classdiagrams_reverseMap[inst]; ok {
					res = _diagrampackage.Name
				}
			}
		}

	case *models.DiagramPackage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Field:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Fields":
				if _gongstructshape, ok := stage.GongStructShape_Fields_reverseMap[inst]; ok {
					res = _gongstructshape.Name
				}
			}
		}

	case *models.GongEnumShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongEnumShapes":
				if _classdiagram, ok := stage.Classdiagram_GongEnumShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		}

	case *models.GongEnumValueEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "GongEnumShape":
			switch reverseField.Fieldname {
			case "GongEnumValueEntrys":
				if _gongenumshape, ok := stage.GongEnumShape_GongEnumValueEntrys_reverseMap[inst]; ok {
					res = _gongenumshape.Name
				}
			}
		}

	case *models.GongStructShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongStructShapes":
				if _classdiagram, ok := stage.Classdiagram_GongStructShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Links":
				if _gongstructshape, ok := stage.GongStructShape_Links_reverseMap[inst]; ok {
					res = _gongstructshape.Name
				}
			}
		}

	case *models.NoteShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "NoteShapes":
				if _classdiagram, ok := stage.Classdiagram_NoteShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		}

	case *models.NoteShapeLink:
		switch reverseField.GongstructName {
		// insertion point
		case "NoteShape":
			switch reverseField.Fieldname {
			case "NoteShapeLinks":
				if _noteshape, ok := stage.NoteShape_NoteShapeLinks_reverseMap[inst]; ok {
					res = _noteshape.Name
				}
			}
		}

	case *models.Position:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UmlState:
		switch reverseField.GongstructName {
		// insertion point
		case "Umlsc":
			switch reverseField.Fieldname {
			case "States":
				if _umlsc, ok := stage.Umlsc_States_reverseMap[inst]; ok {
					res = _umlsc.Name
				}
			}
		}

	case *models.Umlsc:
		switch reverseField.GongstructName {
		// insertion point
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Umlscs":
				if _diagrampackage, ok := stage.DiagramPackage_Umlscs_reverseMap[inst]; ok {
					res = _diagrampackage.Name
				}
			}
		}

	case *models.Vertice:
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
	case *models.Classdiagram:
		switch reverseField.GongstructName {
		// insertion point
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Classdiagrams":
				res = stage.DiagramPackage_Classdiagrams_reverseMap[inst]
			}
		}

	case *models.DiagramPackage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Field:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Fields":
				res = stage.GongStructShape_Fields_reverseMap[inst]
			}
		}

	case *models.GongEnumShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongEnumShapes":
				res = stage.Classdiagram_GongEnumShapes_reverseMap[inst]
			}
		}

	case *models.GongEnumValueEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "GongEnumShape":
			switch reverseField.Fieldname {
			case "GongEnumValueEntrys":
				res = stage.GongEnumShape_GongEnumValueEntrys_reverseMap[inst]
			}
		}

	case *models.GongStructShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongStructShapes":
				res = stage.Classdiagram_GongStructShapes_reverseMap[inst]
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.GongStructShape_Links_reverseMap[inst]
			}
		}

	case *models.NoteShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "NoteShapes":
				res = stage.Classdiagram_NoteShapes_reverseMap[inst]
			}
		}

	case *models.NoteShapeLink:
		switch reverseField.GongstructName {
		// insertion point
		case "NoteShape":
			switch reverseField.Fieldname {
			case "NoteShapeLinks":
				res = stage.NoteShape_NoteShapeLinks_reverseMap[inst]
			}
		}

	case *models.Position:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UmlState:
		switch reverseField.GongstructName {
		// insertion point
		case "Umlsc":
			switch reverseField.Fieldname {
			case "States":
				res = stage.Umlsc_States_reverseMap[inst]
			}
		}

	case *models.Umlsc:
		switch reverseField.GongstructName {
		// insertion point
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Umlscs":
				res = stage.DiagramPackage_Umlscs_reverseMap[inst]
			}
		}

	case *models.Vertice:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
