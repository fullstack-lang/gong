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

	FontStyle FontStyleEnum

	// Impl is the pointer to the implementation of the node in the models of interest
	Impl NodeImplInterface

	// BackgroundColor, if zero value will have the color to default, therwise, the node
	// will have this color
	BackgroundColor string

	IsExpanded bool

	// fields related to the selection of the node by a check box
	HasCheckboxButton  bool
	IsChecked          bool
	IsCheckboxDisabled bool

	// fields related to the selection of the node by a check box
	HasSecondCheckboxButton  bool
	IsSecondCheckboxChecked  bool
	IsSecondCheckboxDisabled bool
	TextAfterSecondCheckbox  string

	// in case the user wants to change the name of the node
	IsInEditMode bool

	// when the user hover on the node, it can click
	IsNodeClickable bool

	// IsWithPreceedingIcon indicates if the node have an icon preceeding the name
	IsWithPreceedingIcon bool

	// PreceedingIcon is the angular material name if [models.IsWithPreceedingIcon] is true
	PreceedingIcon string

	PreceedingSVGIcon *SVGIcon

	Children []*Node

	Buttons []*Button
}

// OnAfterUpdate, notice that node == stagedNode
func (node *Node) OnAfterUpdate(stage *StageStruct, _, frontNode *Node) {

	if node.Impl != nil {
		node.Impl.OnAfterUpdate(stage, node, frontNode)
	}
}
