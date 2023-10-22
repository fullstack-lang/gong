// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdoc/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		tmp := GetInstanceDBFromInstance[models.Classdiagram, ClassdiagramDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Classdiagrams":
				if _diagrampackage, ok := stage.DiagramPackage_Classdiagrams_reverseMap[inst]; ok {
					res = _diagrampackage.Name
				}
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.DiagramPackage:
		tmp := GetInstanceDBFromInstance[models.DiagramPackage, DiagramPackageDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Field:
		tmp := GetInstanceDBFromInstance[models.Field, FieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Fields":
				if _gongstructshape, ok := stage.GongStructShape_Fields_reverseMap[inst]; ok {
					res = _gongstructshape.Name
				}
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongEnumShapes":
				if _classdiagram, ok := stage.Classdiagram_GongEnumShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumValueEntry:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueEntry, GongEnumValueEntryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			case "GongEnumValueEntrys":
				if _gongenumshape, ok := stage.GongEnumShape_GongEnumValueEntrys_reverseMap[inst]; ok {
					res = _gongenumshape.Name
				}
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongStructShapes":
				if _classdiagram, ok := stage.Classdiagram_GongStructShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Links":
				if _gongstructshape, ok := stage.GongStructShape_Links_reverseMap[inst]; ok {
					res = _gongstructshape.Name
				}
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.NoteShape:
		tmp := GetInstanceDBFromInstance[models.NoteShape, NoteShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "NoteShapes":
				if _classdiagram, ok := stage.Classdiagram_NoteShapes_reverseMap[inst]; ok {
					res = _classdiagram.Name
				}
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.NoteShapeLink:
		tmp := GetInstanceDBFromInstance[models.NoteShapeLink, NoteShapeLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			case "NoteShapeLinks":
				if _noteshape, ok := stage.NoteShape_NoteShapeLinks_reverseMap[inst]; ok {
					res = _noteshape.Name
				}
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Position:
		tmp := GetInstanceDBFromInstance[models.Position, PositionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.UmlState:
		tmp := GetInstanceDBFromInstance[models.UmlState, UmlStateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			case "States":
				if _umlsc, ok := stage.Umlsc_States_reverseMap[inst]; ok {
					res = _umlsc.Name
				}
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Umlsc:
		tmp := GetInstanceDBFromInstance[models.Umlsc, UmlscDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Umlscs":
				if _diagrampackage, ok := stage.DiagramPackage_Umlscs_reverseMap[inst]; ok {
					res = _diagrampackage.Name
				}
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Vertice:
		tmp := GetInstanceDBFromInstance[models.Vertice, VerticeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
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
	case *models.Classdiagram:
		tmp := GetInstanceDBFromInstance[models.Classdiagram, ClassdiagramDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Classdiagrams":
				res = stage.DiagramPackage_Classdiagrams_reverseMap[inst]
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.DiagramPackage:
		tmp := GetInstanceDBFromInstance[models.DiagramPackage, DiagramPackageDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Field:
		tmp := GetInstanceDBFromInstance[models.Field, FieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Fields":
				res = stage.GongStructShape_Fields_reverseMap[inst]
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongEnumShapes":
				res = stage.Classdiagram_GongEnumShapes_reverseMap[inst]
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongEnumValueEntry:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueEntry, GongEnumValueEntryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			case "GongEnumValueEntrys":
				res = stage.GongEnumShape_GongEnumValueEntrys_reverseMap[inst]
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongStructShapes":
				res = stage.Classdiagram_GongStructShapes_reverseMap[inst]
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.GongStructShape_Links_reverseMap[inst]
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.NoteShape:
		tmp := GetInstanceDBFromInstance[models.NoteShape, NoteShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "NoteShapes":
				res = stage.Classdiagram_NoteShapes_reverseMap[inst]
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.NoteShapeLink:
		tmp := GetInstanceDBFromInstance[models.NoteShapeLink, NoteShapeLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			case "NoteShapeLinks":
				res = stage.NoteShape_NoteShapeLinks_reverseMap[inst]
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Position:
		tmp := GetInstanceDBFromInstance[models.Position, PositionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.UmlState:
		tmp := GetInstanceDBFromInstance[models.UmlState, UmlStateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			case "States":
				res = stage.Umlsc_States_reverseMap[inst]
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Umlsc:
		tmp := GetInstanceDBFromInstance[models.Umlsc, UmlscDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Umlscs":
				res = stage.DiagramPackage_Umlscs_reverseMap[inst]
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	case *models.Vertice:
		tmp := GetInstanceDBFromInstance[models.Vertice, VerticeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			}
		case "DiagramPackage":
			switch reverseField.Fieldname {
			}
		case "Field":
			switch reverseField.Fieldname {
			}
		case "GongEnumShape":
			switch reverseField.Fieldname {
			}
		case "GongEnumValueEntry":
			switch reverseField.Fieldname {
			}
		case "GongStructShape":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "NoteShape":
			switch reverseField.Fieldname {
			}
		case "NoteShapeLink":
			switch reverseField.Fieldname {
			}
		case "Position":
			switch reverseField.Fieldname {
			}
		case "UmlState":
			switch reverseField.Fieldname {
			}
		case "Umlsc":
			switch reverseField.Fieldname {
			}
		case "Vertice":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
