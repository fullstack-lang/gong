package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treetasks(
	diagramProcess *DiagramProcess,
	task *Task,
	parentNode *tree.Node,
	taskWhoseNodeIsExpanded *[]*Task,
) {
	taskNode := &tree.Node{
		Name:            task.GetName(),
		IsExpanded:      slices.Index(*taskWhoseNodeIsExpanded, task) != -1,
		IsNodeClickable: true,
		IsInEditMode:    task.GetIsInRenameMode(),
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
