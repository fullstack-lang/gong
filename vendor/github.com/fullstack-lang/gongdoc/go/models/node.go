package models

import "log"

// Node is a node in the tree for selecting items to display
// Node implements the visual artifacts of a node (the name, the buttons, the checkbocks, ...)
// The font end read the fields of the node and display the node accordingly
//
// when a user interacts with a node in front, it propagates the action to the back end and
// it is propagated to the [Impl] implementation.
// Conversly, the implementation can set the configuration of the node
type Node struct {
	Name string

	// Impl is the pointer to the implementation of the node in the models of interest
	Impl NodeImplInterface

	IsExpanded bool

	// fields related to the selection of the node by a check box
	HasCheckboxButton  bool
	IsChecked          bool
	IsCheckboxDisabled bool

	// in case the user wants to change the name of the node
	IsInEditMode bool

	Children []*Node

	Buttons []*Button
}

// OnAfterUpdate, notice that node == stagedNode
func (node *Node) OnAfterUpdate(stage *StageStruct, _, frontNode *Node) {

	log.Println("Node, OnAfterUpdate", node.Name)

	if node.Impl != nil {
		node.Impl.OnAfterUpdate(stage, node, frontNode)
	}
}
