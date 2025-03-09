// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/split/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AsSplit:
		assplitInstance := any(concreteInstance).(*models.AsSplit)
		ret2 := backRepo.BackRepoAsSplit.GetAsSplitDBFromAsSplitPtr(assplitInstance)
		ret = any(ret2).(*T2)
	case *models.AsSplitArea:
		assplitareaInstance := any(concreteInstance).(*models.AsSplitArea)
		ret2 := backRepo.BackRepoAsSplitArea.GetAsSplitAreaDBFromAsSplitAreaPtr(assplitareaInstance)
		ret = any(ret2).(*T2)
	case *models.Doc:
		docInstance := any(concreteInstance).(*models.Doc)
		ret2 := backRepo.BackRepoDoc.GetDocDBFromDocPtr(docInstance)
		ret = any(ret2).(*T2)
	case *models.Form:
		formInstance := any(concreteInstance).(*models.Form)
		ret2 := backRepo.BackRepoForm.GetFormDBFromFormPtr(formInstance)
		ret = any(ret2).(*T2)
	case *models.Svg:
		svgInstance := any(concreteInstance).(*models.Svg)
		ret2 := backRepo.BackRepoSvg.GetSvgDBFromSvgPtr(svgInstance)
		ret = any(ret2).(*T2)
	case *models.Table:
		tableInstance := any(concreteInstance).(*models.Table)
		ret2 := backRepo.BackRepoTable.GetTableDBFromTablePtr(tableInstance)
		ret = any(ret2).(*T2)
	case *models.Tree:
		treeInstance := any(concreteInstance).(*models.Tree)
		ret2 := backRepo.BackRepoTree.GetTreeDBFromTreePtr(treeInstance)
		ret = any(ret2).(*T2)
	case *models.View:
		viewInstance := any(concreteInstance).(*models.View)
		ret2 := backRepo.BackRepoView.GetViewDBFromViewPtr(viewInstance)
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
	case *models.AsSplit:
		tmp := GetInstanceDBFromInstance[models.AsSplit, AsSplitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AsSplitArea:
		tmp := GetInstanceDBFromInstance[models.AsSplitArea, AsSplitAreaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Doc:
		tmp := GetInstanceDBFromInstance[models.Doc, DocDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Form:
		tmp := GetInstanceDBFromInstance[models.Form, FormDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Svg:
		tmp := GetInstanceDBFromInstance[models.Svg, SvgDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.View:
		tmp := GetInstanceDBFromInstance[models.View, ViewDB](
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
	case *models.AsSplit:
		tmp := GetInstanceDBFromInstance[models.AsSplit, AsSplitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AsSplitArea:
		tmp := GetInstanceDBFromInstance[models.AsSplitArea, AsSplitAreaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Doc:
		tmp := GetInstanceDBFromInstance[models.Doc, DocDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Form:
		tmp := GetInstanceDBFromInstance[models.Form, FormDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Svg:
		tmp := GetInstanceDBFromInstance[models.Svg, SvgDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.View:
		tmp := GetInstanceDBFromInstance[models.View, ViewDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
