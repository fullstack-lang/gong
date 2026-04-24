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

	addRenameButton(process, processNode, stager)
	processNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(*processsWhoseNodeIsExpanded, process) == -1 {
					*processsWhoseNodeIsExpanded = append(*processsWhoseNodeIsExpanded, process)
				}
			} else {
				if idx := slices.Index(*processsWhoseNodeIsExpanded, process); idx != -1 {
					*processsWhoseNodeIsExpanded = slices.Delete(*processsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(process, GetPointerToGongstructName[*Process]())
		stager.stage.Commit()
	}

	addAddItemButtonSimple(stager, processsWhoseNodeIsExpanded, process, nil, processNode, &process.SubProcesses)

	subProcessesNode := &tree.Node{
		Name:            "SubProcesses",
		FontStyle:       tree.ITALIC,
		IsExpanded:      process.IsSubProcessNodeExpanded,
		IsNodeClickable: true,
	}
	processNode.Children = append(processNode.Children, subProcessesNode)

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

		diagramNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
			if frontNode.IsChecked && !stagedNode.IsChecked {
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
					diagram_.IsChecked = false
				}
				diagram.IsChecked = true
				stagedNode.IsChecked = frontNode.IsChecked
				stager.stage.Commit()
				return
			}
			if !frontNode.IsChecked && stagedNode.IsChecked {
				diagram.IsChecked = false
				stagedNode.IsChecked = frontNode.IsChecked
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
					diagram_.IsChecked = false
				}
				stager.stage.Commit()
				return
			}
			if frontNode.IsExpanded {
				if slices.Index(*processsWhoseNodeIsExpanded, process) == -1 {
					*processsWhoseNodeIsExpanded = append(*processsWhoseNodeIsExpanded, process)
				}
			} else {
				if idx := slices.Index(*processsWhoseNodeIsExpanded, process); idx != -1 {
					*processsWhoseNodeIsExpanded = slices.Delete(*processsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		}

		{
			showPrefixButton := &tree.Button{
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
				showPrefixButton.Icon = string(buttons.BUTTON_label)
				showPrefixButton.ToolTipText = "Show Prefix"
			} else {
				showPrefixButton.Icon = string(buttons.BUTTON_label_off)
				showPrefixButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
		}
	}

}
