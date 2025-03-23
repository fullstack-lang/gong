// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/load/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.FileToDownload:
		filetodownloadInstance := any(concreteInstance).(*models.FileToDownload)
		ret2 := backRepo.BackRepoFileToDownload.GetFileToDownloadDBFromFileToDownloadPtr(filetodownloadInstance)
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
	case *models.FileToDownload:
		tmp := GetInstanceDBFromInstance[models.FileToDownload, FileToDownloadDB](
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
	case *models.FileToDownload:
		tmp := GetInstanceDBFromInstance[models.FileToDownload, FileToDownloadDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
