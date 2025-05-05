// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AttributeShape:
		attributeshapeInstance := any(concreteInstance).(*models.AttributeShape)
		ret2 := backRepo.BackRepoAttributeShape.GetAttributeShapeDBFromAttributeShapePtr(attributeshapeInstance)
		ret = any(ret2).(*T2)
	case *models.Classdiagram:
		classdiagramInstance := any(concreteInstance).(*models.Classdiagram)
		ret2 := backRepo.BackRepoClassdiagram.GetClassdiagramDBFromClassdiagramPtr(classdiagramInstance)
		ret = any(ret2).(*T2)
	case *models.DiagramPackage:
		diagrampackageInstance := any(concreteInstance).(*models.DiagramPackage)
		ret2 := backRepo.BackRepoDiagramPackage.GetDiagramPackageDBFromDiagramPackagePtr(diagrampackageInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnumShape:
		gongenumshapeInstance := any(concreteInstance).(*models.GongEnumShape)
		ret2 := backRepo.BackRepoGongEnumShape.GetGongEnumShapeDBFromGongEnumShapePtr(gongenumshapeInstance)
		ret = any(ret2).(*T2)
	case *models.GongEnumValueShape:
		gongenumvalueshapeInstance := any(concreteInstance).(*models.GongEnumValueShape)
		ret2 := backRepo.BackRepoGongEnumValueShape.GetGongEnumValueShapeDBFromGongEnumValueShapePtr(gongenumvalueshapeInstance)
		ret = any(ret2).(*T2)
	case *models.GongNoteLinkShape:
		gongnotelinkshapeInstance := any(concreteInstance).(*models.GongNoteLinkShape)
		ret2 := backRepo.BackRepoGongNoteLinkShape.GetGongNoteLinkShapeDBFromGongNoteLinkShapePtr(gongnotelinkshapeInstance)
		ret = any(ret2).(*T2)
	case *models.GongNoteShape:
		gongnoteshapeInstance := any(concreteInstance).(*models.GongNoteShape)
		ret2 := backRepo.BackRepoGongNoteShape.GetGongNoteShapeDBFromGongNoteShapePtr(gongnoteshapeInstance)
		ret = any(ret2).(*T2)
	case *models.GongStructShape:
		gongstructshapeInstance := any(concreteInstance).(*models.GongStructShape)
		ret2 := backRepo.BackRepoGongStructShape.GetGongStructShapeDBFromGongStructShapePtr(gongstructshapeInstance)
		ret = any(ret2).(*T2)
	case *models.LinkShape:
		linkshapeInstance := any(concreteInstance).(*models.LinkShape)
		ret2 := backRepo.BackRepoLinkShape.GetLinkShapeDBFromLinkShapePtr(linkshapeInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AttributeShape:
		tmp := GetInstanceDBFromInstance[models.AttributeShape, AttributeShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
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
	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValueShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueShape, GongEnumValueShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNoteLinkShape:
		tmp := GetInstanceDBFromInstance[models.GongNoteLinkShape, GongNoteLinkShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNoteShape:
		tmp := GetInstanceDBFromInstance[models.GongNoteShape, GongNoteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LinkShape:
		tmp := GetInstanceDBFromInstance[models.LinkShape, LinkShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AttributeShape:
		tmp := GetInstanceDBFromInstance[models.AttributeShape, AttributeShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
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
	case *models.GongEnumShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumShape, GongEnumShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongEnumValueShape:
		tmp := GetInstanceDBFromInstance[models.GongEnumValueShape, GongEnumValueShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNoteLinkShape:
		tmp := GetInstanceDBFromInstance[models.GongNoteLinkShape, GongNoteLinkShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongNoteShape:
		tmp := GetInstanceDBFromInstance[models.GongNoteShape, GongNoteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GongStructShape:
		tmp := GetInstanceDBFromInstance[models.GongStructShape, GongStructShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LinkShape:
		tmp := GetInstanceDBFromInstance[models.LinkShape, LinkShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
