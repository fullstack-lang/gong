// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtree/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Button:
		tmp := GetInstanceDBFromInstance[models.Button, ButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Buttons":
				if tmp != nil && tmp.Node_ButtonsDBID.Int64 != 0 {
					id := uint(tmp.Node_ButtonsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Tree":
			switch reverseField.Fieldname {
			}
		}

	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Children":
				if tmp != nil && tmp.Node_ChildrenDBID.Int64 != 0 {
					id := uint(tmp.Node_ChildrenDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Tree":
			switch reverseField.Fieldname {
			case "RootNodes":
				if tmp != nil && tmp.Tree_RootNodesDBID.Int64 != 0 {
					id := uint(tmp.Tree_RootNodesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Tree":
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
	case *models.Button:
		tmp := GetInstanceDBFromInstance[models.Button, ButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Buttons":
				if tmp != nil && tmp.Node_ButtonsDBID.Int64 != 0 {
					id := uint(tmp.Node_ButtonsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget
				}
			}
		case "Tree":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Children":
				if tmp != nil && tmp.Node_ChildrenDBID.Int64 != 0 {
					id := uint(tmp.Node_ChildrenDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget
				}
			}
		case "Tree":
			switch reverseField.Fieldname {
			case "RootNodes":
				if tmp != nil && tmp.Tree_RootNodesDBID.Int64 != 0 {
					id := uint(tmp.Tree_RootNodesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTree.Map_TreeDBID_TreePtr[id]
					res = reservePointerTarget
				}
			}
		}
	
	case *models.Tree:
		tmp := GetInstanceDBFromInstance[models.Tree, TreeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Button":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Tree":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
