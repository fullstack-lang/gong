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
				if tmp != nil && tmp.DiagramPackage_ClassdiagramsDBID.Int64 != 0 {
					id := uint(tmp.DiagramPackage_ClassdiagramsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.GongStructShape_FieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStructShape_FieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Classdiagram_GongEnumShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_GongEnumShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64 != 0 {
					id := uint(tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Classdiagram_GongStructShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_GongStructShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.GongStructShape_LinksDBID.Int64 != 0 {
					id := uint(tmp.GongStructShape_LinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Classdiagram_NoteShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_NoteShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.NoteShape_NoteShapeLinksDBID.Int64 != 0 {
					id := uint(tmp.NoteShape_NoteShapeLinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Umlsc_StatesDBID.Int64 != 0 {
					id := uint(tmp.Umlsc_StatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.DiagramPackage_UmlscsDBID.Int64 != 0 {
					id := uint(tmp.DiagramPackage_UmlscsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.DiagramPackage_ClassdiagramsDBID.Int64 != 0 {
					id := uint(tmp.DiagramPackage_ClassdiagramsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.GongStructShape_FieldsDBID.Int64 != 0 {
					id := uint(tmp.GongStructShape_FieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.Classdiagram_GongEnumShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_GongEnumShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64 != 0 {
					id := uint(tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.Classdiagram_GongStructShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_GongStructShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.GongStructShape_LinksDBID.Int64 != 0 {
					id := uint(tmp.GongStructShape_LinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.Classdiagram_NoteShapesDBID.Int64 != 0 {
					id := uint(tmp.Classdiagram_NoteShapesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.NoteShape_NoteShapeLinksDBID.Int64 != 0 {
					id := uint(tmp.NoteShape_NoteShapeLinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.Umlsc_StatesDBID.Int64 != 0 {
					id := uint(tmp.Umlsc_StatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.DiagramPackage_UmlscsDBID.Int64 != 0 {
					id := uint(tmp.DiagramPackage_UmlscsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
					res = reservePointerTarget
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
	return res
}
