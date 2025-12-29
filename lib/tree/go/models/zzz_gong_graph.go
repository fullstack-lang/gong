// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *Node:
		ok = stage.IsStagedNode(target)

	case *SVGIcon:
		ok = stage.IsStagedSVGIcon(target)

	case *Tree:
		ok = stage.IsStagedTree(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *Node:
		ok = stage.IsStagedNode(target)

	case *SVGIcon:
		ok = stage.IsStagedSVGIcon(target)

	case *Tree:
		ok = stage.IsStagedTree(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedNode(node *Node) (ok bool) {

	_, ok = stage.Nodes[node]

	return
}

func (stage *Stage) IsStagedSVGIcon(svgicon *SVGIcon) (ok bool) {

	_, ok = stage.SVGIcons[svgicon]

	return
}

func (stage *Stage) IsStagedTree(tree *Tree) (ok bool) {

	_, ok = stage.Trees[tree]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Button:
		stage.StageBranchButton(target)

	case *Node:
		stage.StageBranchNode(target)

	case *SVGIcon:
		stage.StageBranchSVGIcon(target)

	case *Tree:
		stage.StageBranchTree(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if IsStaged(stage, button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		StageBranch(stage, button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNode(node *Node) {

	// check if instance is already staged
	if IsStaged(stage, node) {
		return
	}

	node.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if node.PreceedingSVGIcon != nil {
		StageBranch(stage, node.PreceedingSVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Children {
		StageBranch(stage, _node)
	}
	for _, _button := range node.Buttons {
		StageBranch(stage, _button)
	}

}

func (stage *Stage) StageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if IsStaged(stage, svgicon) {
		return
	}

	svgicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTree(tree *Tree) {

	// check if instance is already staged
	if IsStaged(stage, tree) {
		return
	}

	tree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range tree.RootNodes {
		StageBranch(stage, _node)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Button:
		toT := CopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Node:
		toT := CopyBranchNode(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVGIcon:
		toT := CopyBranchSVGIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tree:
		toT := CopyBranchTree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.CopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers
	if buttonFrom.SVGIcon != nil {
		buttonTo.SVGIcon = CopyBranchSVGIcon(mapOrigCopy, buttonFrom.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {

	// nodeFrom has already been copied
	if _nodeTo, ok := mapOrigCopy[nodeFrom]; ok {
		nodeTo = _nodeTo.(*Node)
		return
	}

	nodeTo = new(Node)
	mapOrigCopy[nodeFrom] = nodeTo
	nodeFrom.CopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers
	if nodeFrom.PreceedingSVGIcon != nil {
		nodeTo.PreceedingSVGIcon = CopyBranchSVGIcon(mapOrigCopy, nodeFrom.PreceedingSVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Children {
		nodeTo.Children = append(nodeTo.Children, CopyBranchNode(mapOrigCopy, _node))
	}
	for _, _button := range nodeFrom.Buttons {
		nodeTo.Buttons = append(nodeTo.Buttons, CopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func CopyBranchSVGIcon(mapOrigCopy map[any]any, svgiconFrom *SVGIcon) (svgiconTo *SVGIcon) {

	// svgiconFrom has already been copied
	if _svgiconTo, ok := mapOrigCopy[svgiconFrom]; ok {
		svgiconTo = _svgiconTo.(*SVGIcon)
		return
	}

	svgiconTo = new(SVGIcon)
	mapOrigCopy[svgiconFrom] = svgiconTo
	svgiconFrom.CopyBasicFields(svgiconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTree(mapOrigCopy map[any]any, treeFrom *Tree) (treeTo *Tree) {

	// treeFrom has already been copied
	if _treeTo, ok := mapOrigCopy[treeFrom]; ok {
		treeTo = _treeTo.(*Tree)
		return
	}

	treeTo = new(Tree)
	mapOrigCopy[treeFrom] = treeTo
	treeFrom.CopyBasicFields(treeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range treeFrom.RootNodes {
		treeTo.RootNodes = append(treeTo.RootNodes, CopyBranchNode(mapOrigCopy, _node))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Button:
		stage.UnstageBranchButton(target)

	case *Node:
		stage.UnstageBranchNode(target)

	case *SVGIcon:
		stage.UnstageBranchSVGIcon(target)

	case *Tree:
		stage.UnstageBranchTree(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !IsStaged(stage, button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		UnstageBranch(stage, button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNode(node *Node) {

	// check if instance is already staged
	if !IsStaged(stage, node) {
		return
	}

	node.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if node.PreceedingSVGIcon != nil {
		UnstageBranch(stage, node.PreceedingSVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Children {
		UnstageBranch(stage, _node)
	}
	for _, _button := range node.Buttons {
		UnstageBranch(stage, _button)
	}

}

func (stage *Stage) UnstageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if !IsStaged(stage, svgicon) {
		return
	}

	svgicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTree(tree *Tree) {

	// check if instance is already staged
	if !IsStaged(stage, tree) {
		return
	}

	tree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range tree.RootNodes {
		UnstageBranch(stage, _node)
	}

}


// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, "Name")
	}
	if button.Icon != buttonOther.Icon {
		diffs = append(diffs, "Icon")
	}
	if (button.SVGIcon == nil) != (buttonOther.SVGIcon == nil) {
		diffs = append(diffs, "SVGIcon")
	} else if button.SVGIcon != nil && buttonOther.SVGIcon != nil {
		if button.SVGIcon != buttonOther.SVGIcon {
			diffs = append(diffs, "SVGIcon")
		}
	}
	if button.IsDisabled != buttonOther.IsDisabled {
		diffs = append(diffs, "IsDisabled")
	}
	if button.HasToolTip != buttonOther.HasToolTip {
		diffs = append(diffs, "HasToolTip")
	}
	if button.ToolTipText != buttonOther.ToolTipText {
		diffs = append(diffs, "ToolTipText")
	}
	if button.ToolTipPosition != buttonOther.ToolTipPosition {
		diffs = append(diffs, "ToolTipPosition")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (node *Node) GongDiff(nodeOther *Node) (diffs []string) {
	// insertion point for field diffs
	if node.Name != nodeOther.Name {
		diffs = append(diffs, "Name")
	}
	if node.FontStyle != nodeOther.FontStyle {
		diffs = append(diffs, "FontStyle")
	}
	if node.BackgroundColor != nodeOther.BackgroundColor {
		diffs = append(diffs, "BackgroundColor")
	}
	if node.IsExpanded != nodeOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if node.HasCheckboxButton != nodeOther.HasCheckboxButton {
		diffs = append(diffs, "HasCheckboxButton")
	}
	if node.IsChecked != nodeOther.IsChecked {
		diffs = append(diffs, "IsChecked")
	}
	if node.IsCheckboxDisabled != nodeOther.IsCheckboxDisabled {
		diffs = append(diffs, "IsCheckboxDisabled")
	}
	if node.CheckboxHasToolTip != nodeOther.CheckboxHasToolTip {
		diffs = append(diffs, "CheckboxHasToolTip")
	}
	if node.CheckboxToolTipText != nodeOther.CheckboxToolTipText {
		diffs = append(diffs, "CheckboxToolTipText")
	}
	if node.CheckboxToolTipPosition != nodeOther.CheckboxToolTipPosition {
		diffs = append(diffs, "CheckboxToolTipPosition")
	}
	if node.HasSecondCheckboxButton != nodeOther.HasSecondCheckboxButton {
		diffs = append(diffs, "HasSecondCheckboxButton")
	}
	if node.IsSecondCheckboxChecked != nodeOther.IsSecondCheckboxChecked {
		diffs = append(diffs, "IsSecondCheckboxChecked")
	}
	if node.IsSecondCheckboxDisabled != nodeOther.IsSecondCheckboxDisabled {
		diffs = append(diffs, "IsSecondCheckboxDisabled")
	}
	if node.SecondCheckboxHasToolTip != nodeOther.SecondCheckboxHasToolTip {
		diffs = append(diffs, "SecondCheckboxHasToolTip")
	}
	if node.SecondCheckboxToolTipText != nodeOther.SecondCheckboxToolTipText {
		diffs = append(diffs, "SecondCheckboxToolTipText")
	}
	if node.SecondCheckboxToolTipPosition != nodeOther.SecondCheckboxToolTipPosition {
		diffs = append(diffs, "SecondCheckboxToolTipPosition")
	}
	if node.TextAfterSecondCheckbox != nodeOther.TextAfterSecondCheckbox {
		diffs = append(diffs, "TextAfterSecondCheckbox")
	}
	if node.HasToolTip != nodeOther.HasToolTip {
		diffs = append(diffs, "HasToolTip")
	}
	if node.ToolTipText != nodeOther.ToolTipText {
		diffs = append(diffs, "ToolTipText")
	}
	if node.ToolTipPosition != nodeOther.ToolTipPosition {
		diffs = append(diffs, "ToolTipPosition")
	}
	if node.IsInEditMode != nodeOther.IsInEditMode {
		diffs = append(diffs, "IsInEditMode")
	}
	if node.IsNodeClickable != nodeOther.IsNodeClickable {
		diffs = append(diffs, "IsNodeClickable")
	}
	if node.IsWithPreceedingIcon != nodeOther.IsWithPreceedingIcon {
		diffs = append(diffs, "IsWithPreceedingIcon")
	}
	if node.PreceedingIcon != nodeOther.PreceedingIcon {
		diffs = append(diffs, "PreceedingIcon")
	}
	if (node.PreceedingSVGIcon == nil) != (nodeOther.PreceedingSVGIcon == nil) {
		diffs = append(diffs, "PreceedingSVGIcon")
	} else if node.PreceedingSVGIcon != nil && nodeOther.PreceedingSVGIcon != nil {
		if node.PreceedingSVGIcon != nodeOther.PreceedingSVGIcon {
			diffs = append(diffs, "PreceedingSVGIcon")
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
                if len(node.Children[i].GongDiff(nodeOther.Children[i])) > 0 {
                    ChildrenDifferent = true
                    break
                }
            }
        }
    }
    if ChildrenDifferent {
        diffs = append(diffs, "Children")
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
                if len(node.Buttons[i].GongDiff(nodeOther.Buttons[i])) > 0 {
                    ButtonsDifferent = true
                    break
                }
            }
        }
    }
    if ButtonsDifferent {
        diffs = append(diffs, "Buttons")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgicon *SVGIcon) GongDiff(svgiconOther *SVGIcon) (diffs []string) {
	// insertion point for field diffs
	if svgicon.Name != svgiconOther.Name {
		diffs = append(diffs, "Name")
	}
	if svgicon.SVG != svgiconOther.SVG {
		diffs = append(diffs, "SVG")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tree *Tree) GongDiff(treeOther *Tree) (diffs []string) {
	// insertion point for field diffs
	if tree.Name != treeOther.Name {
		diffs = append(diffs, "Name")
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
                if len(tree.RootNodes[i].GongDiff(treeOther.RootNodes[i])) > 0 {
                    RootNodesDifferent = true
                    break
                }
            }
        }
    }
    if RootNodesDifferent {
        diffs = append(diffs, "RootNodes")
    }

	return
}
