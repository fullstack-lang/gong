// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (button *Button) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	return button.GongIsStaged(stage)
}

func (menu *Menu) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Menus[menu]

	return
}

func (stage *Stage) IsStagedMenu(menu *Menu) (ok bool) {

	return menu.GongIsStaged(stage)
}

func (node *Node) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Nodes[node]

	return
}

func (stage *Stage) IsStagedNode(node *Node) (ok bool) {

	return node.GongIsStaged(stage)
}

func (svgicon *SVGIcon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SVGIcons[svgicon]

	return
}

func (stage *Stage) IsStagedSVGIcon(svgicon *SVGIcon) (ok bool) {

	return svgicon.GongIsStaged(stage)
}

func (tree *Tree) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Trees[tree]

	return
}

func (stage *Stage) IsStagedTree(tree *Tree) (ok bool) {

	return tree.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (button *Button) GongStageBranch(stage *Stage) {
	stage.StageBranchButton(button)
}

func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if stage.IsStaged(button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		stage.StageBranch(button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (menu *Menu) GongStageBranch(stage *Stage) {
	stage.StageBranchMenu(menu)
}

func (stage *Stage) StageBranchMenu(menu *Menu) {

	// check if instance is already staged
	if stage.IsStaged(menu) {
		return
	}

	menu.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range menu.Buttons {
		stage.StageBranch(_button)
	}

}

func (node *Node) GongStageBranch(stage *Stage) {
	stage.StageBranchNode(node)
}

func (stage *Stage) StageBranchNode(node *Node) {

	// check if instance is already staged
	if stage.IsStaged(node) {
		return
	}

	node.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if node.PreceedingSVGIcon != nil {
		stage.StageBranch(node.PreceedingSVGIcon)
	}
	if node.Menu != nil {
		stage.StageBranch(node.Menu)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Children {
		stage.StageBranch(_node)
	}
	for _, _button := range node.Buttons {
		stage.StageBranch(_button)
	}

}

func (svgicon *SVGIcon) GongStageBranch(stage *Stage) {
	stage.StageBranchSVGIcon(svgicon)
}

func (stage *Stage) StageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if stage.IsStaged(svgicon) {
		return
	}

	svgicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tree *Tree) GongStageBranch(stage *Stage) {
	stage.StageBranchTree(tree)
}

func (stage *Stage) StageBranchTree(tree *Tree) {

	// check if instance is already staged
	if stage.IsStaged(tree) {
		return
	}

	tree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range tree.RootNodes {
		stage.StageBranch(_node)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Button:
		toT := GongCopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Menu:
		toT := GongCopyBranchMenu(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Node:
		toT := GongCopyBranchNode(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVGIcon:
		toT := GongCopyBranchSVGIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tree:
		toT := GongCopyBranchTree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers
	if buttonFrom.SVGIcon != nil {
		buttonTo.SVGIcon = GongCopyBranchSVGIcon(mapOrigCopy, buttonFrom.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMenu(mapOrigCopy map[any]any, menuFrom *Menu) (menuTo *Menu) {

	// menuFrom has already been copied
	if _menuTo, ok := mapOrigCopy[menuFrom]; ok {
		menuTo = _menuTo.(*Menu)
		return
	}

	menuTo = new(Menu)
	mapOrigCopy[menuFrom] = menuTo
	menuFrom.GongCopyBasicFields(menuTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range menuFrom.Buttons {
		menuTo.Buttons = append(menuTo.Buttons, GongCopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func GongCopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {

	// nodeFrom has already been copied
	if _nodeTo, ok := mapOrigCopy[nodeFrom]; ok {
		nodeTo = _nodeTo.(*Node)
		return
	}

	nodeTo = new(Node)
	mapOrigCopy[nodeFrom] = nodeTo
	nodeFrom.GongCopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers
	if nodeFrom.PreceedingSVGIcon != nil {
		nodeTo.PreceedingSVGIcon = GongCopyBranchSVGIcon(mapOrigCopy, nodeFrom.PreceedingSVGIcon)
	}
	if nodeFrom.Menu != nil {
		nodeTo.Menu = GongCopyBranchMenu(mapOrigCopy, nodeFrom.Menu)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Children {
		nodeTo.Children = append(nodeTo.Children, GongCopyBranchNode(mapOrigCopy, _node))
	}
	for _, _button := range nodeFrom.Buttons {
		nodeTo.Buttons = append(nodeTo.Buttons, GongCopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func GongCopyBranchSVGIcon(mapOrigCopy map[any]any, svgiconFrom *SVGIcon) (svgiconTo *SVGIcon) {

	// svgiconFrom has already been copied
	if _svgiconTo, ok := mapOrigCopy[svgiconFrom]; ok {
		svgiconTo = _svgiconTo.(*SVGIcon)
		return
	}

	svgiconTo = new(SVGIcon)
	mapOrigCopy[svgiconFrom] = svgiconTo
	svgiconFrom.GongCopyBasicFields(svgiconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTree(mapOrigCopy map[any]any, treeFrom *Tree) (treeTo *Tree) {

	// treeFrom has already been copied
	if _treeTo, ok := mapOrigCopy[treeFrom]; ok {
		treeTo = _treeTo.(*Tree)
		return
	}

	treeTo = new(Tree)
	mapOrigCopy[treeFrom] = treeTo
	treeFrom.GongCopyBasicFields(treeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range treeFrom.RootNodes {
		treeTo.RootNodes = append(treeTo.RootNodes, GongCopyBranchNode(mapOrigCopy, _node))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (button *Button) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchButton(button)
}

func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !stage.IsStaged(button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		stage.UnstageBranch(button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (menu *Menu) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMenu(menu)
}

func (stage *Stage) UnstageBranchMenu(menu *Menu) {

	// check if instance is already staged
	if !stage.IsStaged(menu) {
		return
	}

	menu.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range menu.Buttons {
		stage.UnstageBranch(_button)
	}

}

func (node *Node) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNode(node)
}

func (stage *Stage) UnstageBranchNode(node *Node) {

	// check if instance is already staged
	if !stage.IsStaged(node) {
		return
	}

	node.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if node.PreceedingSVGIcon != nil {
		stage.UnstageBranch(node.PreceedingSVGIcon)
	}
	if node.Menu != nil {
		stage.UnstageBranch(node.Menu)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Children {
		stage.UnstageBranch(_node)
	}
	for _, _button := range node.Buttons {
		stage.UnstageBranch(_button)
	}

}

func (svgicon *SVGIcon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSVGIcon(svgicon)
}

func (stage *Stage) UnstageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if !stage.IsStaged(svgicon) {
		return
	}

	svgicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tree *Tree) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTree(tree)
}

func (stage *Stage) UnstageBranchTree(tree *Tree) {

	// check if instance is already staged
	if !stage.IsStaged(tree) {
		return
	}

	tree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range tree.RootNodes {
		stage.UnstageBranch(_node)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Button) GongReconstructPointersFromReferences(stage *Stage, instance *Button) {
	// insertion point for pointers field
	if instance.SVGIcon != nil {
		reference.SVGIcon = stage.SVGIcons_reference[instance.SVGIcon]
	}
	// insertion point for slice of pointers field
}

func (reference *Menu) GongReconstructPointersFromReferences(stage *Stage, instance *Menu) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Buttons = reference.Buttons[:0]
	for _, _b := range instance.Buttons {
		reference.Buttons = append(reference.Buttons, stage.Buttons_reference[_b])
	}
}

func (reference *Node) GongReconstructPointersFromReferences(stage *Stage, instance *Node) {
	// insertion point for pointers field
	if instance.PreceedingSVGIcon != nil {
		reference.PreceedingSVGIcon = stage.SVGIcons_reference[instance.PreceedingSVGIcon]
	}
	if instance.Menu != nil {
		reference.Menu = stage.Menus_reference[instance.Menu]
	}
	// insertion point for slice of pointers field
	reference.Children = reference.Children[:0]
	for _, _b := range instance.Children {
		reference.Children = append(reference.Children, stage.Nodes_reference[_b])
	}
	reference.Buttons = reference.Buttons[:0]
	for _, _b := range instance.Buttons {
		reference.Buttons = append(reference.Buttons, stage.Buttons_reference[_b])
	}
}

func (reference *SVGIcon) GongReconstructPointersFromReferences(stage *Stage, instance *SVGIcon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tree) GongReconstructPointersFromReferences(stage *Stage, instance *Tree) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RootNodes = reference.RootNodes[:0]
	for _, _b := range instance.RootNodes {
		reference.RootNodes = append(reference.RootNodes, stage.Nodes_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *Button) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SVGIcon; _reference != nil {
		reference.SVGIcon = nil
		if _instance, ok := stage.SVGIcons_instance[_reference]; ok {
			reference.SVGIcon = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Menu) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Buttons []*Button
	for _, _reference := range reference.Buttons {
		if _instance, ok := stage.Buttons_instance[_reference]; ok {
			_Buttons = append(_Buttons, _instance)
		}
	}
	reference.Buttons = _Buttons
}

func (reference *Node) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.PreceedingSVGIcon; _reference != nil {
		reference.PreceedingSVGIcon = nil
		if _instance, ok := stage.SVGIcons_instance[_reference]; ok {
			reference.PreceedingSVGIcon = _instance
		}
	}
	if _reference := reference.Menu; _reference != nil {
		reference.Menu = nil
		if _instance, ok := stage.Menus_instance[_reference]; ok {
			reference.Menu = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Children []*Node
	for _, _reference := range reference.Children {
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			_Children = append(_Children, _instance)
		}
	}
	reference.Children = _Children
	var _Buttons []*Button
	for _, _reference := range reference.Buttons {
		if _instance, ok := stage.Buttons_instance[_reference]; ok {
			_Buttons = append(_Buttons, _instance)
		}
	}
	reference.Buttons = _Buttons
}

func (reference *SVGIcon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tree) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RootNodes []*Node
	for _, _reference := range reference.RootNodes {
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			_RootNodes = append(_RootNodes, _instance)
		}
	}
	reference.RootNodes = _RootNodes
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(stage *Stage, buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, button.GongMarshallField(stage, "Name"))
	}
	if button.Icon != buttonOther.Icon {
		diffs = append(diffs, button.GongMarshallField(stage, "Icon"))
	}
	if (button.SVGIcon == nil) != (buttonOther.SVGIcon == nil) {
		diffs = append(diffs, button.GongMarshallField(stage, "SVGIcon"))
	} else if button.SVGIcon != nil && buttonOther.SVGIcon != nil {
		if button.SVGIcon != buttonOther.SVGIcon {
			diffs = append(diffs, button.GongMarshallField(stage, "SVGIcon"))
		}
	}
	if button.IsDisabled != buttonOther.IsDisabled {
		diffs = append(diffs, button.GongMarshallField(stage, "IsDisabled"))
	}
	if button.HasToolTip != buttonOther.HasToolTip {
		diffs = append(diffs, button.GongMarshallField(stage, "HasToolTip"))
	}
	if button.ToolTipText != buttonOther.ToolTipText {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipText"))
	}
	if button.ToolTipPosition != buttonOther.ToolTipPosition {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipPosition"))
	}
	if button.ClientOnX != buttonOther.ClientOnX {
		diffs = append(diffs, button.GongMarshallField(stage, "ClientOnX"))
	}
	if button.ClientOnY != buttonOther.ClientOnY {
		diffs = append(diffs, button.GongMarshallField(stage, "ClientOnY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (menu *Menu) GongDiff(stage *Stage, menuOther *Menu) (diffs []string) {
	// insertion point for field diffs
	if menu.Name != menuOther.Name {
		diffs = append(diffs, menu.GongMarshallField(stage, "Name"))
	}
	ButtonsDifferent := false
	if len(menu.Buttons) != len(menuOther.Buttons) {
		ButtonsDifferent = true
	} else {
		for i := range menu.Buttons {
			if (menu.Buttons[i] == nil) != (menuOther.Buttons[i] == nil) {
				ButtonsDifferent = true
				break
			} else if menu.Buttons[i] != nil && menuOther.Buttons[i] != nil {
				// this is a pointer comparaison
				if menu.Buttons[i] != menuOther.Buttons[i] {
					ButtonsDifferent = true
					break
				}
			}
		}
	}
	if ButtonsDifferent {
		ops := stage.Diff(
			menu,
			"Buttons",
			len(menuOther.Buttons),
			len(menu.Buttons),
			func(i, j int) bool {
				return menuOther.Buttons[i] == menu.Buttons[j]
			},
			func(j int) string {
				return menu.Buttons[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (node *Node) GongDiff(stage *Stage, nodeOther *Node) (diffs []string) {
	// insertion point for field diffs
	if node.Name != nodeOther.Name {
		diffs = append(diffs, node.GongMarshallField(stage, "Name"))
	}
	if node.IsWithPrefix != nodeOther.IsWithPrefix {
		diffs = append(diffs, node.GongMarshallField(stage, "IsWithPrefix"))
	}
	if node.Prefix != nodeOther.Prefix {
		diffs = append(diffs, node.GongMarshallField(stage, "Prefix"))
	}
	if node.FontStyle != nodeOther.FontStyle {
		diffs = append(diffs, node.GongMarshallField(stage, "FontStyle"))
	}
	if node.BackgroundColor != nodeOther.BackgroundColor {
		diffs = append(diffs, node.GongMarshallField(stage, "BackgroundColor"))
	}
	if node.IsExpanded != nodeOther.IsExpanded {
		diffs = append(diffs, node.GongMarshallField(stage, "IsExpanded"))
	}
	if node.HasCheckboxButton != nodeOther.HasCheckboxButton {
		diffs = append(diffs, node.GongMarshallField(stage, "HasCheckboxButton"))
	}
	if node.IsChecked != nodeOther.IsChecked {
		diffs = append(diffs, node.GongMarshallField(stage, "IsChecked"))
	}
	if node.IsCheckboxDisabled != nodeOther.IsCheckboxDisabled {
		diffs = append(diffs, node.GongMarshallField(stage, "IsCheckboxDisabled"))
	}
	if node.CheckboxHasToolTip != nodeOther.CheckboxHasToolTip {
		diffs = append(diffs, node.GongMarshallField(stage, "CheckboxHasToolTip"))
	}
	if node.CheckboxToolTipText != nodeOther.CheckboxToolTipText {
		diffs = append(diffs, node.GongMarshallField(stage, "CheckboxToolTipText"))
	}
	if node.CheckboxToolTipPosition != nodeOther.CheckboxToolTipPosition {
		diffs = append(diffs, node.GongMarshallField(stage, "CheckboxToolTipPosition"))
	}
	if node.HasSecondCheckboxButton != nodeOther.HasSecondCheckboxButton {
		diffs = append(diffs, node.GongMarshallField(stage, "HasSecondCheckboxButton"))
	}
	if node.IsSecondCheckboxChecked != nodeOther.IsSecondCheckboxChecked {
		diffs = append(diffs, node.GongMarshallField(stage, "IsSecondCheckboxChecked"))
	}
	if node.IsSecondCheckboxDisabled != nodeOther.IsSecondCheckboxDisabled {
		diffs = append(diffs, node.GongMarshallField(stage, "IsSecondCheckboxDisabled"))
	}
	if node.SecondCheckboxHasToolTip != nodeOther.SecondCheckboxHasToolTip {
		diffs = append(diffs, node.GongMarshallField(stage, "SecondCheckboxHasToolTip"))
	}
	if node.SecondCheckboxToolTipText != nodeOther.SecondCheckboxToolTipText {
		diffs = append(diffs, node.GongMarshallField(stage, "SecondCheckboxToolTipText"))
	}
	if node.SecondCheckboxToolTipPosition != nodeOther.SecondCheckboxToolTipPosition {
		diffs = append(diffs, node.GongMarshallField(stage, "SecondCheckboxToolTipPosition"))
	}
	if node.TextAfterSecondCheckbox != nodeOther.TextAfterSecondCheckbox {
		diffs = append(diffs, node.GongMarshallField(stage, "TextAfterSecondCheckbox"))
	}
	if node.HasToolTip != nodeOther.HasToolTip {
		diffs = append(diffs, node.GongMarshallField(stage, "HasToolTip"))
	}
	if node.ToolTipText != nodeOther.ToolTipText {
		diffs = append(diffs, node.GongMarshallField(stage, "ToolTipText"))
	}
	if node.ToolTipPosition != nodeOther.ToolTipPosition {
		diffs = append(diffs, node.GongMarshallField(stage, "ToolTipPosition"))
	}
	if node.ClientOnY != nodeOther.ClientOnY {
		diffs = append(diffs, node.GongMarshallField(stage, "ClientOnY"))
	}
	if node.IsInEditMode != nodeOther.IsInEditMode {
		diffs = append(diffs, node.GongMarshallField(stage, "IsInEditMode"))
	}
	if node.IsNodeClickable != nodeOther.IsNodeClickable {
		diffs = append(diffs, node.GongMarshallField(stage, "IsNodeClickable"))
	}
	if node.IsWithPreceedingIcon != nodeOther.IsWithPreceedingIcon {
		diffs = append(diffs, node.GongMarshallField(stage, "IsWithPreceedingIcon"))
	}
	if node.PreceedingIcon != nodeOther.PreceedingIcon {
		diffs = append(diffs, node.GongMarshallField(stage, "PreceedingIcon"))
	}
	if (node.PreceedingSVGIcon == nil) != (nodeOther.PreceedingSVGIcon == nil) {
		diffs = append(diffs, node.GongMarshallField(stage, "PreceedingSVGIcon"))
	} else if node.PreceedingSVGIcon != nil && nodeOther.PreceedingSVGIcon != nil {
		if node.PreceedingSVGIcon != nodeOther.PreceedingSVGIcon {
			diffs = append(diffs, node.GongMarshallField(stage, "PreceedingSVGIcon"))
		}
	}
	ChildrenDifferent := false
	if len(node.Children) != len(nodeOther.Children) {
		ChildrenDifferent = true
	} else {
		for i := range node.Children {
			if (node.Children[i] == nil) != (nodeOther.Children[i] == nil) {
				ChildrenDifferent = true
				break
			} else if node.Children[i] != nil && nodeOther.Children[i] != nil {
				// this is a pointer comparaison
				if node.Children[i] != nodeOther.Children[i] {
					ChildrenDifferent = true
					break
				}
			}
		}
	}
	if ChildrenDifferent {
		ops := stage.Diff(
			node,
			"Children",
			len(nodeOther.Children),
			len(node.Children),
			func(i, j int) bool {
				return nodeOther.Children[i] == node.Children[j]
			},
			func(j int) string {
				return node.Children[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ButtonsDifferent := false
	if len(node.Buttons) != len(nodeOther.Buttons) {
		ButtonsDifferent = true
	} else {
		for i := range node.Buttons {
			if (node.Buttons[i] == nil) != (nodeOther.Buttons[i] == nil) {
				ButtonsDifferent = true
				break
			} else if node.Buttons[i] != nil && nodeOther.Buttons[i] != nil {
				// this is a pointer comparaison
				if node.Buttons[i] != nodeOther.Buttons[i] {
					ButtonsDifferent = true
					break
				}
			}
		}
	}
	if ButtonsDifferent {
		ops := stage.Diff(
			node,
			"Buttons",
			len(nodeOther.Buttons),
			len(node.Buttons),
			func(i, j int) bool {
				return nodeOther.Buttons[i] == node.Buttons[j]
			},
			func(j int) string {
				return node.Buttons[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (node.Menu == nil) != (nodeOther.Menu == nil) {
		diffs = append(diffs, node.GongMarshallField(stage, "Menu"))
	} else if node.Menu != nil && nodeOther.Menu != nil {
		if node.Menu != nodeOther.Menu {
			diffs = append(diffs, node.GongMarshallField(stage, "Menu"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgicon *SVGIcon) GongDiff(stage *Stage, svgiconOther *SVGIcon) (diffs []string) {
	// insertion point for field diffs
	if svgicon.Name != svgiconOther.Name {
		diffs = append(diffs, svgicon.GongMarshallField(stage, "Name"))
	}
	if svgicon.SVG != svgiconOther.SVG {
		diffs = append(diffs, svgicon.GongMarshallField(stage, "SVG"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tree *Tree) GongDiff(stage *Stage, treeOther *Tree) (diffs []string) {
	// insertion point for field diffs
	if tree.Name != treeOther.Name {
		diffs = append(diffs, tree.GongMarshallField(stage, "Name"))
	}
	RootNodesDifferent := false
	if len(tree.RootNodes) != len(treeOther.RootNodes) {
		RootNodesDifferent = true
	} else {
		for i := range tree.RootNodes {
			if (tree.RootNodes[i] == nil) != (treeOther.RootNodes[i] == nil) {
				RootNodesDifferent = true
				break
			} else if tree.RootNodes[i] != nil && treeOther.RootNodes[i] != nil {
				// this is a pointer comparaison
				if tree.RootNodes[i] != treeOther.RootNodes[i] {
					RootNodesDifferent = true
					break
				}
			}
		}
	}
	if RootNodesDifferent {
		ops := stage.Diff(
			tree,
			"RootNodes",
			len(treeOther.RootNodes),
			len(tree.RootNodes),
			func(i, j int) bool {
				return treeOther.RootNodes[i] == tree.RootNodes[j]
			},
			func(j int) string {
				return tree.RootNodes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if tree.HaveSearch != treeOther.HaveSearch {
		diffs = append(diffs, tree.GongMarshallField(stage, "HaveSearch"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
