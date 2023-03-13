package models

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

	HasAddChildButton bool

	// to edit the name
	HasEditButton bool
	IsInEditMode  bool

	// to edit the diagram
	HasDrawButton    bool
	HasDrawOffButton bool
	IsInDrawMode     bool
	IsSaved          bool

	HasDeleteButton bool

	Children []*Node
}
