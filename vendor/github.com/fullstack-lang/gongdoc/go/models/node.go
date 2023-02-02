package models

// Node is a node in the tree for selecting items to display
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
