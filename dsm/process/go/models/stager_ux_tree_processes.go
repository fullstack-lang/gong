package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProcesses(
	process *Process,
	parentNode *tree.Node,
	processsWhoseNodeIsExpanded *[]*Process,
) {
	processNode := &tree.Node{
		Name:            process.GetName(),
		IsExpanded:      slices.Index(*processsWhoseNodeIsExpanded, process) != -1,
		IsNodeClickable: true,
		IsInEditMode:    process.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, processNode)

	addRenameButton(process, processNode, stager)
	processNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			process.SetName(frontNode.Name)
			process.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
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

	// Diagrams
	for _, diagramProcess := range process.DiagramProcesss {
		diagramNode := &tree.Node{
			Name:              diagramProcess.Name,
			IsExpanded:        slices.Contains(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess) == true,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagramProcess.IsChecked,

			IsInEditMode: diagramProcess.isInRenameMode,
		}
		processNode.Children = append(processNode.Children, diagramNode)

		element := diagramProcess
		node := diagramNode

		addRenameButton(element, node, stager)

		diagramNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
			if frontNode.Name != stagedNode.Name {
				diagramProcess.Name = frontNode.Name
				diagramProcess.isInRenameMode = false
				stager.stage.Commit()
				return
			}
			if frontNode.IsChecked && !stagedNode.IsChecked {
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
					diagram_.IsChecked = false
				}
				diagramProcess.IsChecked = true
				stagedNode.IsChecked = frontNode.IsChecked
				stager.stage.Commit()
				return
			}
			if !frontNode.IsChecked && stagedNode.IsChecked {
				diagramProcess.IsChecked = false
				stagedNode.IsChecked = frontNode.IsChecked
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
					diagram_.IsChecked = false
				}
				stager.stage.Commit()
				return
			}
			if frontNode.IsExpanded {
				if slices.Index(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess) == -1 {
					process.DiagramProcessWhoseNodeIsExpanded = append(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess)
				}
			} else {
				if idx := slices.Index(process.DiagramProcessWhoseNodeIsExpanded, diagramProcess); idx != -1 {
					process.DiagramProcessWhoseNodeIsExpanded = slices.Delete(process.DiagramProcessWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.probeForm.FillUpFormFromGongstruct(diagramProcess, "Diagram")
		}

		// prefix button
		{
			showPrefixButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,

				OnClick: func() {
					diagramProcess.IsShowPrefix = !diagramProcess.IsShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagramProcess.IsShowPrefix {
				showPrefixButton.Icon = string(buttons.BUTTON_label)
				showPrefixButton.ToolTipText = "Show Prefix"
			} else {
				showPrefixButton.Icon = string(buttons.BUTTON_label_off)
				showPrefixButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
		}

		// Participants
		{
			participantsNode := &tree.Node{
				Name:            "Participants",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagramProcess.IsParticipantsNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, participantsNode)
			participantsNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					if frontNode.IsExpanded {
						diagramProcess.IsParticipantsNodeExpanded = true
					} else {
						diagramProcess.IsParticipantsNodeExpanded = false
					}
					stager.stage.Commit()
					return
				}
			}

			for _, participant := range process.Participants {
				stager.treeParticipants(diagramProcess, participant, participantsNode)
			}
			addAddItemButtonVerySimple(
				stager,
				&diagramProcess.IsParticipantsNodeExpanded,
				participantsNode,
				&process.Participants)

		}
	}

	itemAdderCallback := addAddItemButtonVerySimple(stager, &process.isExpanded, processNode, &process.DiagramProcesss)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.isExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}
	addAddItemButtonSimple(stager, processsWhoseNodeIsExpanded, process, nil, processNode, &process.SubProcesses)

	// SubProcesses
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
}
