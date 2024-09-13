// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdoc/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Classdiagram:
		classdiagramInstance := any(concreteInstance).(*models.Classdiagram)
		ret2 := backRepo.BackRepoClassdiagram.GetClassdiagramDBFromClassdiagramPtr(classdiagramInstance)
		ret = any(ret2).(*T2)
	case *models.DiagramPackage:
		diagrampackageInstance := any(concreteInstance).(*models.DiagramPackage)
		ret2 := backRepo.BackRepoDiagramPackage.GetDiagramPackageDBFromDiagramPackagePtr(diagrampackageInstance)
		ret = any(ret2).(*T2)
	case *models.Field:
		fieldInstance := any(concreteInstance).(*models.Field)
		ret2 := backRepo.BackRepoField.GetFieldDBFromFieldPtr(fieldInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnumShape:
		gongenumshapeInstance := any(concreteInstance).(*models.GongEnumShape)
		ret2 := backRepo.BackRepoGongEnumShape.GetGongEnumShapeDBFromGongEnumShapePtr(gongenumshapeInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnumValueEntry:
		gongenumvalueentryInstance := any(concreteInstance).(*models.GongEnumValueEntry)
		ret2 := backRepo.BackRepoGongEnumValueEntry.GetGongEnumValueEntryDBFromGongEnumValueEntryPtr(gongenumvalueentryInstance)
		ret = any(ret2).(*T2)
	case *models.GongStructShape:
		gongstructshapeInstance := any(concreteInstance).(*models.GongStructShape)
		ret2 := backRepo.BackRepoGongStructShape.GetGongStructShapeDBFromGongStructShapePtr(gongstructshapeInstance)
		ret = any(ret2).(*T2)
	case *models.Link:
		linkInstance := any(concreteInstance).(*models.Link)
		ret2 := backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkInstance)
		ret = any(ret2).(*T2)
	case *models.NoteShape:
		noteshapeInstance := any(concreteInstance).(*models.NoteShape)
		ret2 := backRepo.BackRepoNoteShape.GetNoteShapeDBFromNoteShapePtr(noteshapeInstance)
		ret = any(ret2).(*T2)
	case *models.NoteShapeLink:
		noteshapelinkInstance := any(concreteInstance).(*models.NoteShapeLink)
		ret2 := backRepo.BackRepoNoteShapeLink.GetNoteShapeLinkDBFromNoteShapeLinkPtr(noteshapelinkInstance)
		ret = any(ret2).(*T2)
	case *models.Position:
		positionInstance := any(concreteInstance).(*models.Position)
		ret2 := backRepo.BackRepoPosition.GetPositionDBFromPositionPtr(positionInstance)
		ret = any(ret2).(*T2)
	case *models.UmlState:
		umlstateInstance := any(concreteInstance).(*models.UmlState)
		ret2 := backRepo.BackRepoUmlState.GetUmlStateDBFromUmlStatePtr(umlstateInstance)
		ret = any(ret2).(*T2)
	case *models.Umlsc:
		umlscInstance := any(concreteInstance).(*models.Umlsc)
		ret2 := backRepo.BackRepoUmlsc.GetUmlscDBFromUmlscPtr(umlscInstance)
		ret = any(ret2).(*T2)
	case *models.Vertice:
		verticeInstance := any(concreteInstance).(*models.Vertice)
		ret2 := backRepo.BackRepoVertice.GetVerticeDBFromVerticePtr(verticeInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Classdiagram:
		tmp := GetInstanceDBFromInstance[models.Classdiagram, ClassdiagramDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DiagramPackage:
		tmp := GetInstanceDBFromInstance[models.DiagramPackage, DiagramPackageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Field:
		tmp := GetInstanceDBFromInstance[models.Field, FieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValueEntry:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueEntry, GongEnumValueEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteShape:
		tmp := GetInstanceDBFromInstance[models.NoteShape, NoteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteShapeLink:
		tmp := GetInstanceDBFromInstance[models.NoteShapeLink, NoteShapeLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Position:
		tmp := GetInstanceDBFromInstance[models.Position, PositionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UmlState:
		tmp := GetInstanceDBFromInstance[models.UmlState, UmlStateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Umlsc:
		tmp := GetInstanceDBFromInstance[models.Umlsc, UmlscDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Vertice:
		tmp := GetInstanceDBFromInstance[models.Vertice, VerticeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Classdiagram:
		tmp := GetInstanceDBFromInstance[models.Classdiagram, ClassdiagramDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DiagramPackage:
		tmp := GetInstanceDBFromInstance[models.DiagramPackage, DiagramPackageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Field:
		tmp := GetInstanceDBFromInstance[models.Field, FieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValueEntry:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueEntry, GongEnumValueEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteShape:
		tmp := GetInstanceDBFromInstance[models.NoteShape, NoteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteShapeLink:
		tmp := GetInstanceDBFromInstance[models.NoteShapeLink, NoteShapeLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Position:
		tmp := GetInstanceDBFromInstance[models.Position, PositionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UmlState:
		tmp := GetInstanceDBFromInstance[models.UmlState, UmlStateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Umlsc:
		tmp := GetInstanceDBFromInstance[models.Umlsc, UmlscDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Vertice:
		tmp := GetInstanceDBFromInstance[models.Vertice, VerticeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
