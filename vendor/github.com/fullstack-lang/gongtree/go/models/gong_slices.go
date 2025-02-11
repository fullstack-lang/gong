// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Node
	// insertion point per field
	clear(stage.Node_Children_reverseMap)
	stage.Node_Children_reverseMap = make(map[*Node]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _node := range node.Children {
			stage.Node_Children_reverseMap[_node] = node
		}
	}
	clear(stage.Node_Buttons_reverseMap)
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
	clear(stage.Tree_RootNodes_reverseMap)
	stage.Tree_RootNodes_reverseMap = make(map[*Node]*Tree)
	for tree := range stage.Trees {
		_ = tree
		for _, _node := range tree.RootNodes {
			stage.Tree_RootNodes_reverseMap[_node] = tree
		}
	}

}
