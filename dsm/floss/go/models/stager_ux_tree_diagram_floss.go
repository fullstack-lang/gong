package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramFloss(
	system *System,
	diagramFloss *DiagramFloss,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagramFloss.Name,
		IsExpanded:        slices.Contains(system.DiagramFlossWhoseNodeIsExpanded, diagramFloss) == true,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagramFloss.IsChecked,

		IsInEditMode: diagramFloss.isInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	element := diagramFloss
	node := diagramNode

	addRenameButton(element, node, stager)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			// uncheck all diagrams
			for diagramFloss_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
				diagramFloss_.IsChecked = false
			}

			diagramFloss.IsChecked = true
			stager.stage.Commit()
			return
		} else {
			diagramFloss.IsChecked = false
			stager.stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, diagramFloss)
	node.OnNameChange = stager.onNameChange(diagramFloss)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagramFloss, &system.DiagramFlossWhoseNodeIsExpanded)

	// prefix button
	{
		showPrefixButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagramFloss.IsShowPrefix = !diagramFloss.IsShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagramFloss.IsShowPrefix {
			showPrefixButton.Icon = string(buttons.BUTTON_label)
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
	}
}
