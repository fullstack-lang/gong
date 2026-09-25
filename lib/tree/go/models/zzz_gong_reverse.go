// generated code - do not edit
package models

// insertion point
func (inst *Button) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Menu":
		switch reverseField.Fieldname {
		case "Buttons":
			if _menu, ok := stage.Menu_Buttons_reverseMap[inst]; ok {
				res = _menu.Name
			}
		}
	case "Node":
		switch reverseField.Fieldname {
		case "Buttons":
			if _node, ok := stage.Node_Buttons_reverseMap[inst]; ok {
				res = _node.Name
			}
		}
	}
	return
}

func (inst *Menu) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Node) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
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
	return
}

func (inst *SVGIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tree) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
