// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdoc/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		tmp := GetInstanceDBFromInstance[models.Classdiagram, ClassdiagramDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Classdiagrams":
			if tmp.DiagramPackage_ClassdiagramsDBID.Int64 != 0 {
				id := uint(tmp.DiagramPackage_ClassdiagramsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.DiagramPackage:
		tmp := GetInstanceDBFromInstance[models.DiagramPackage, DiagramPackageDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.Field:
		tmp := GetInstanceDBFromInstance[models.Field, FieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Fields":
			if tmp.GongStructShape_FieldsDBID.Int64 != 0 {
				id := uint(tmp.GongStructShape_FieldsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongEnumShapes":
			if tmp.Classdiagram_GongEnumShapesDBID.Int64 != 0 {
				id := uint(tmp.Classdiagram_GongEnumShapesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongEnumValueEntry:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueEntry, GongEnumValueEntryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongEnumValueEntrys":
			if tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64 != 0 {
				id := uint(tmp.GongEnumShape_GongEnumValueEntrysDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "GongStructShapes":
			if tmp.Classdiagram_GongStructShapesDBID.Int64 != 0 {
				id := uint(tmp.Classdiagram_GongStructShapesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Links":
			if tmp.GongStructShape_LinksDBID.Int64 != 0 {
				id := uint(tmp.GongStructShape_LinksDBID.Int64)
				reservePointerTarget := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.NoteShape:
		tmp := GetInstanceDBFromInstance[models.NoteShape, NoteShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "NoteShapes":
			if tmp.Classdiagram_NoteShapesDBID.Int64 != 0 {
				id := uint(tmp.Classdiagram_NoteShapesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.NoteShapeLink:
		tmp := GetInstanceDBFromInstance[models.NoteShapeLink, NoteShapeLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "NoteShapeLinks":
			if tmp.NoteShape_NoteShapeLinksDBID.Int64 != 0 {
				id := uint(tmp.NoteShape_NoteShapeLinksDBID.Int64)
				reservePointerTarget := backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Position:
		tmp := GetInstanceDBFromInstance[models.Position, PositionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.UmlState:
		tmp := GetInstanceDBFromInstance[models.UmlState, UmlStateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "States":
			if tmp.Umlsc_StatesDBID.Int64 != 0 {
				id := uint(tmp.Umlsc_StatesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Umlsc:
		tmp := GetInstanceDBFromInstance[models.Umlsc, UmlscDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Umlscs":
			if tmp.DiagramPackage_UmlscsDBID.Int64 != 0 {
				id := uint(tmp.DiagramPackage_UmlscsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackagePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Vertice:
		tmp := GetInstanceDBFromInstance[models.Vertice, VerticeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}
