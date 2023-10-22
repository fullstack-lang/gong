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
				if _node, ok := stage.Node_Buttons_reverseMap[inst]; ok {
					res = _node.Name
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
				if _node, ok := stage.Node_Children_reverseMap[inst]; ok {
					res = _node.Name
				}
			}
		case "Tree":
			switch reverseField.Fieldname {
			case "RootNodes":
				if _tree, ok := stage.Tree_RootNodes_reverseMap[inst]; ok {
					res = _tree.Name
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
				res = stage.Node_Buttons_reverseMap[inst]
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
				res = stage.Node_Children_reverseMap[inst]
			}
		case "Tree":
			switch reverseField.Fieldname {
			case "RootNodes":
				res = stage.Tree_RootNodes_reverseMap[inst]
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
