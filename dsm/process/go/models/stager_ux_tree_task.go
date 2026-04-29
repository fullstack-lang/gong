package models

import (
	"log"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treetasks(
	diagramProcess *DiagramProcess,
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
}
