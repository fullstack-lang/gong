// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/markdown/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Content:
		contentInstance := any(concreteInstance).(*models.Content)
		ret2 := backRepo.BackRepoContent.GetContentDBFromContentPtr(contentInstance)
		ret = any(ret2).(*T2)
	case *models.PngImage:
		pngimageInstance := any(concreteInstance).(*models.PngImage)
		ret2 := backRepo.BackRepoPngImage.GetPngImageDBFromPngImagePtr(pngimageInstance)
		ret = any(ret2).(*T2)
	case *models.SvgImage:
		svgimageInstance := any(concreteInstance).(*models.SvgImage)
		ret2 := backRepo.BackRepoSvgImage.GetSvgImageDBFromSvgImagePtr(svgimageInstance)
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
	case *models.Content:
		tmp := GetInstanceDBFromInstance[models.Content, ContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PngImage:
		tmp := GetInstanceDBFromInstance[models.PngImage, PngImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SvgImage:
		tmp := GetInstanceDBFromInstance[models.SvgImage, SvgImageDB](
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
	case *models.Content:
		tmp := GetInstanceDBFromInstance[models.Content, ContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PngImage:
		tmp := GetInstanceDBFromInstance[models.PngImage, PngImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SvgImage:
		tmp := GetInstanceDBFromInstance[models.SvgImage, SvgImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
