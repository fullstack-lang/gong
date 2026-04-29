package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treetasks(
	diagramProcess *DiagramProcess,
	participant *Participant,
	task *Task,
	parentNode *tree.Node,
	taskWhoseNodeIsExpanded *[]*Task,
) {
	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagramProcess.map_Task_TaskShape[task]

	taskNode := &tree.Node{
		Name:              task.GetName(),
		IsExpanded:        slices.Index(*taskWhoseNodeIsExpanded, task) != -1,
		IsNodeClickable:   true,
		IsInEditMode:      task.GetIsInRenameMode(),
		HasCheckboxButton: true,
		IsChecked:         ok,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	addRenameButton(task, taskNode, stager)

	if shape, ok := diagramProcess.map_Task_TaskShape[task]; ok {
		taskNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		taskNode.Buttons = append(taskNode.Buttons, visibilityButton)
	}

	taskNode.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			task.SetName(frontNode.Name)
			task.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(task, diagramProcess, &diagramProcess.TaskShapes, stage)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)

			// not necessary since there is a semantic rule (gong clean) that remove the shape from the slice when it is unstaged
			idx := slices.Index(diagramProcess.TaskShapes, shape)
			diagramProcess.TaskShapes = slices.Delete(diagramProcess.TaskShapes, idx, idx+1)
			stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(*taskWhoseNodeIsExpanded, task) == -1 {
					*taskWhoseNodeIsExpanded = append(*taskWhoseNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(*taskWhoseNodeIsExpanded, task); idx != -1 {
					*taskWhoseNodeIsExpanded = slices.Delete(*taskWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(task, GetPointerToGongstructName[*Task]())
		stager.stage.Commit()
	}

	nodeOutControlFlows := &tree.Node{
		Name:            "out control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseOutControlFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeOutControlFlows)
	nodeOutControlFlows.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != nodeOutControlFlows.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(participant.TaskWhoseOutControlFlowsNodeIsExpanded, task) == -1 {
					participant.TaskWhoseOutControlFlowsNodeIsExpanded = append(participant.TaskWhoseOutControlFlowsNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(participant.TaskWhoseOutControlFlowsNodeIsExpanded, task); idx != -1 {
					participant.TaskWhoseOutControlFlowsNodeIsExpanded = slices.Delete(participant.TaskWhoseOutControlFlowsNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
		}
	}
	for _, controlFlow := range task.outcontrolFlows {
		stager.treecontrolflowsWithinTask(diagramProcess, controlFlow, nodeOutControlFlows)
	}

	nodeInControlFlows := &tree.Node{
		Name:            "in control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseInControlFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeInControlFlows)
	nodeInControlFlows.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != nodeInControlFlows.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(participant.TaskWhoseInControlFlowsNodeIsExpanded, task) == -1 {
					participant.TaskWhoseInControlFlowsNodeIsExpanded = append(participant.TaskWhoseInControlFlowsNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(participant.TaskWhoseInControlFlowsNodeIsExpanded, task); idx != -1 {
					participant.TaskWhoseInControlFlowsNodeIsExpanded = slices.Delete(participant.TaskWhoseInControlFlowsNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
		}
	}
	for _, controlFlow := range task.incontrolFlows {
		stager.treecontrolflowsWithinTask(diagramProcess, controlFlow, nodeInControlFlows)
	}
}
