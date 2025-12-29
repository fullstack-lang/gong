// generated code - do not edit
package models

// insertion point
func (inst *Button) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
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

func (inst *Node) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *SVGIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Tree) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *Button) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Node":
		switch reverseField.Fieldname {
		case "Buttons":
			res = stage.Node_Buttons_reverseMap[inst]
		}
	}
	return res
}

func (inst *Node) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
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
	return res
}

func (inst *SVGIcon) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Tree) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
