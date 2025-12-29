// generated code - do not edit
package models

import "time"
var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Node
	// insertion point per field
	stage.Node_Children_reverseMap = make(map[*Node]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _node := range node.Children {
			stage.Node_Children_reverseMap[_node] = node
		}
	}
	stage.Node_Buttons_reverseMap = make(map[*Button]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _button := range node.Buttons {
			stage.Node_Buttons_reverseMap[_button] = node
		}
	}

	// Compute reverse map for named struct SVGIcon
	// insertion point per field

	// Compute reverse map for named struct Tree
	// insertion point per field
	stage.Tree_RootNodes_reverseMap = make(map[*Node]*Tree)
	for tree := range stage.Trees {
		_ = tree
		for _, _node := range tree.RootNodes {
			stage.Tree_RootNodes_reverseMap[_node] = tree
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Nodes {
		res = append(res, instance)
	}

	for instance := range stage.SVGIcons {
		res = append(res, instance)
	}

	for instance := range stage.Trees {
		res = append(res, instance)
	}

	return
}


// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := *button
	return &newInstance
}

func (node *Node) GongCopy() GongstructIF {
	newInstance := *node
	return &newInstance
}

func (svgicon *SVGIcon) GongCopy() GongstructIF {
	newInstance := *svgicon
	return &newInstance
}

func (tree *Tree) GongCopy() GongstructIF {
	newInstance := *tree
	return &newInstance
}


func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenDeletedInstances int
	
	// insertion point per named struct
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if _, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of Button "+button.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of Button "+button.Name,
				)
			}
		}
	}

	lenNewInstances += len(buttons_newInstances)
	lenDeletedInstances += len(buttons_deletedInstances)
	var nodes_newInstances []*Node
	var nodes_deletedInstances []*Node

	// parse all staged instances and check if they have a reference
	for node := range stage.Nodes {
		if _, ok := stage.Nodes_reference[node]; !ok {
			nodes_newInstances = append(nodes_newInstances, node)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of Node "+node.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for node := range stage.Nodes_reference {
		if _, ok := stage.Nodes[node]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, node)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of Node "+node.Name,
				)
			}
		}
	}

	lenNewInstances += len(nodes_newInstances)
	lenDeletedInstances += len(nodes_deletedInstances)
	var svgicons_newInstances []*SVGIcon
	var svgicons_deletedInstances []*SVGIcon

	// parse all staged instances and check if they have a reference
	for svgicon := range stage.SVGIcons {
		if _, ok := stage.SVGIcons_reference[svgicon]; !ok {
			svgicons_newInstances = append(svgicons_newInstances, svgicon)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of SVGIcon "+svgicon.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svgicon := range stage.SVGIcons_reference {
		if _, ok := stage.SVGIcons[svgicon]; !ok {
			svgicons_deletedInstances = append(svgicons_deletedInstances, svgicon)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of SVGIcon "+svgicon.Name,
				)
			}
		}
	}

	lenNewInstances += len(svgicons_newInstances)
	lenDeletedInstances += len(svgicons_deletedInstances)
	var trees_newInstances []*Tree
	var trees_deletedInstances []*Tree

	// parse all staged instances and check if they have a reference
	for tree := range stage.Trees {
		if _, ok := stage.Trees_reference[tree]; !ok {
			trees_newInstances = append(trees_newInstances, tree)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of Tree "+tree.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tree := range stage.Trees_reference {
		if _, ok := stage.Trees[tree]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, tree)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of Tree "+tree.Name,
				)
			}
		}
	}

	lenNewInstances += len(trees_newInstances)
	lenDeletedInstances += len(trees_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Buttons_reference = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	for instance := range stage.Nodes {
		stage.Nodes_reference[instance] = instance
	}

	stage.SVGIcons_reference = make(map[*SVGIcon]*SVGIcon)
	for instance := range stage.SVGIcons {
		stage.SVGIcons_reference[instance] = instance
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	for instance := range stage.Trees {
		stage.Trees_reference[instance] = instance
	}

}
