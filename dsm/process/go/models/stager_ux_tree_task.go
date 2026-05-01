package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treetasks(
	diagramProcess *DiagramProcess,
	participant *Participant,
	task *Task,
	parentNode *tree.Node,
	taskWhoseNodeIsExpanded *[]*Task,
) {

	taskNodeConf := TreeNodeAndShapeConfigurationWithoutLink[
		*Task, Task,
		*Participant, Participant,
		*TaskShape, TaskShape,
		*DiagramProcess,
	]{
		diagram:                     diagramProcess,
		parentNode:                  parentNode,
		element:                     task,
		parentElement:               participant,
		elementsWhoseNodeIsExpanded: taskWhoseNodeIsExpanded,
		shapes:                      &diagramProcess.TaskShapes,
		shapesMap:                   diagramProcess.map_Task_TaskShape,
	}
	taskNode := addNodeToTreeWithoutLinkWithConf(stager, taskNodeConf)

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
