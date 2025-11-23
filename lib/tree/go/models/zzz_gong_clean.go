// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	for button := range stage.Buttons {
		_ = button
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, button.SVGIcon) {
			button.SVGIcon = nil
		}
	}

	// Compute reverse map for named struct Node
	for node := range stage.Nodes {
		_ = node
		// insertion point per field
		var _Children []*Node
		for _, _node := range node.Children {
			if IsStaged(stage, _node) {
			 	_Children = append(_Children, _node)
			}
		}
		node.Children = _Children
		var _Buttons []*Button
		for _, _button := range node.Buttons {
			if IsStaged(stage, _button) {
			 	_Buttons = append(_Buttons, _button)
			}
		}
		node.Buttons = _Buttons
		// insertion point per field
		if !IsStaged(stage, node.PreceedingSVGIcon) {
			node.PreceedingSVGIcon = nil
		}
	}

	// Compute reverse map for named struct SVGIcon
	for svgicon := range stage.SVGIcons {
		_ = svgicon
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Tree
	for tree := range stage.Trees {
		_ = tree
		// insertion point per field
		var _RootNodes []*Node
		for _, _node := range tree.RootNodes {
			if IsStaged(stage, _node) {
			 	_RootNodes = append(_RootNodes, _node)
			}
		}
		tree.RootNodes = _RootNodes
		// insertion point per field
	}

}
