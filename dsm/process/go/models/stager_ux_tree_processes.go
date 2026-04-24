package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProcesses(
	process *Process,
	parentNode *tree.Node,
	processsWhoseNodeIsExpanded *[]*Process) {
	processNode := &tree.Node{
		Name:       process.GetName(),
		IsExpanded: slices.Index(*processsWhoseNodeIsExpanded, process) != -1,
	}

	parentNode.Children = append(parentNode.Children, processNode)

	addAddItemButtonSimple(stager, processsWhoseNodeIsExpanded, process, nil, processNode, &process.SubProcesses)

	subProcessesNode := &tree.Node{
		Name:            "SubProcesses",
		FontStyle:       tree.ITALIC,
		IsExpanded:      process.IsSubProcessNodeExpanded,
		IsNodeClickable: true,
	}
	processNode.Children = append(processNode.Children, subProcessesNode)
	subProcessesNode.OnUpdate = stager.OnUpdateExpansion(&process.IsSubProcessNodeExpanded)

	for _, process_ := range process.SubProcesses {
		stager.treeProcesses(process_, subProcessesNode, processsWhoseNodeIsExpanded)
	}

	itemAdderCallback := addAddItemButtonSimple(stager, nil, nil, &process.isExpanded, processNode, &process.DiagramProcesss)

	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.isExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagram := range process.DiagramProcesss {
		diagramNode := &tree.Node{
			Name:              diagram.Name,
			IsExpanded:        diagram.isExpanded,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagram.IsChecked,

			IsInEditMode: diagram.isInRenameMode,
		}
		processNode.Children = append(processNode.Children, diagramNode)

		element := diagram
		node := diagramNode

		addRenameButton(element, node, stager)

		diagramNode.OnUpdate = stager.OnUpdateDiagram(diagram)

		{
			showAllButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,

				OnClick: func() {
					diagram.IsShowPrefix = !diagram.IsShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagram.IsShowPrefix {
				showAllButton.Icon = string(buttons.BUTTON_label)
				showAllButton.ToolTipText = "Show Prefix"
			} else {
				showAllButton.Icon = string(buttons.BUTTON_label_off)
				showAllButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}
	}

}
