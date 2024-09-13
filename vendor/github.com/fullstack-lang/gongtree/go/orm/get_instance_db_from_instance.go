// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtree/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Button:
		buttonInstance := any(concreteInstance).(*models.Button)
		ret2 := backRepo.BackRepoButton.GetButtonDBFromButtonPtr(buttonInstance)
		ret = any(ret2).(*T2)
	case *models.Node:
		nodeInstance := any(concreteInstance).(*models.Node)
		ret2 := backRepo.BackRepoNode.GetNodeDBFromNodePtr(nodeInstance)
		ret = any(ret2).(*T2)
	case *models.SVGIcon:
		svgiconInstance := any(concreteInstance).(*models.SVGIcon)
		ret2 := backRepo.BackRepoSVGIcon.GetSVGIconDBFromSVGIconPtr(svgiconInstance)
		ret = any(ret2).(*T2)
	case *models.Tree:
		treeInstance := any(concreteInstance).(*models.Tree)
		ret2 := backRepo.BackRepoTree.GetTreeDBFromTreePtr(treeInstance)
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
	case *models.Button:
		tmp := GetInstanceDBFromInstance[models.Button, ButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SVGIcon:
		tmp := GetInstanceDBFromInstance[models.SVGIcon, SVGIconDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
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
	case *models.Button:
		tmp := GetInstanceDBFromInstance[models.Button, ButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SVGIcon:
		tmp := GetInstanceDBFromInstance[models.SVGIcon, SVGIconDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
