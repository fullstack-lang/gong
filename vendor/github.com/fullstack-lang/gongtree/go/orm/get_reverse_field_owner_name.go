// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtree/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Button:
		tmp := GetInstanceDBFromInstance[models.Button, ButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Buttons":
			if tmp.Node_ButtonsDBID.Int64 != 0 {
				id := uint(tmp.Node_ButtonsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Children":
			if tmp.Node_ChildrenDBID.Int64 != 0 {
				id := uint(tmp.Node_ChildrenDBID.Int64)
				reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
				res = reservePointerTarget.Name
			}
		case "RootNodes":
			if tmp.Tree_RootNodesDBID.Int64 != 0 {
				id := uint(tmp.Tree_RootNodesDBID.Int64)
				reservePointerTarget := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
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
