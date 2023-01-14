package models

// Tree is a tree of nodes for selecting items to display
type Tree struct {
	Name string

	Type TreeType

	RootNodes []*Node

	// map to access each node according to its id
	// This is usefull since, when navigating a classdiagram
	// one meet a classshape and one need to pick up
	// the corresponding node
	nodeMap map[string]*Node

	// map to navigate from a children node to its parent
	map_Children_Parent map[*Node]*Node
}

func (tree *Tree) UncheckAndDisable() {

	for _, _node := range tree.RootNodes {
		_node.UncheckAndDisableBranch()
	}
}
