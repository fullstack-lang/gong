// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Chapter:
		chapterInstance := any(concreteInstance).(*models.Chapter)
		ret2 := backRepo.BackRepoChapter.GetChapterDBFromChapterPtr(chapterInstance)
		ret = any(ret2).(*T2)
	case *models.Content:
		contentInstance := any(concreteInstance).(*models.Content)
		ret2 := backRepo.BackRepoContent.GetContentDBFromContentPtr(contentInstance)
		ret = any(ret2).(*T2)
	case *models.Page:
		pageInstance := any(concreteInstance).(*models.Page)
		ret2 := backRepo.BackRepoPage.GetPageDBFromPagePtr(pageInstance)
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
	case *models.Chapter:
		tmp := GetInstanceDBFromInstance[models.Chapter, ChapterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Content:
		tmp := GetInstanceDBFromInstance[models.Content, ContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page:
		tmp := GetInstanceDBFromInstance[models.Page, PageDB](
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
	case *models.Chapter:
		tmp := GetInstanceDBFromInstance[models.Chapter, ChapterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Content:
		tmp := GetInstanceDBFromInstance[models.Content, ContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page:
		tmp := GetInstanceDBFromInstance[models.Page, PageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
