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
		shapes:                      &diagramProcess.Task_Shapes,
		shapesMap:                   diagramProcess.map_Task_TaskShape,
	}
	taskNode := addNodeToTreeWithoutLink(stager, taskNodeConf)

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
	for _, controlFlow := range task.outControlFlows {
		stager.treeControlFlowsWithinTask(diagramProcess, controlFlow, nodeOutControlFlows)
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
	for _, controlFlow := range task.inControlFlows {
		stager.treeControlFlowsWithinTask(diagramProcess, controlFlow, nodeInControlFlows)
	}

	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseOutDataFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != nodeOutDataFlows.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(participant.TaskWhoseOutDataFlowsNodeIsExpanded, task) == -1 {
					participant.TaskWhoseOutDataFlowsNodeIsExpanded = append(participant.TaskWhoseOutDataFlowsNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(participant.TaskWhoseOutDataFlowsNodeIsExpanded, task); idx != -1 {
					participant.TaskWhoseOutDataFlowsNodeIsExpanded = slices.Delete(participant.TaskWhoseOutDataFlowsNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
		}
	}
	for _, dataFlow := range task.outDataFlows {
		stager.treeDataFlowsWithinTask(diagramProcess, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseInDataFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeInDataFlows)
	nodeInDataFlows.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != nodeInDataFlows.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(participant.TaskWhoseInDataFlowsNodeIsExpanded, task) == -1 {
					participant.TaskWhoseInDataFlowsNodeIsExpanded = append(participant.TaskWhoseInDataFlowsNodeIsExpanded, task)
				}
			} else {
				if idx := slices.Index(participant.TaskWhoseInDataFlowsNodeIsExpanded, task); idx != -1 {
					participant.TaskWhoseInDataFlowsNodeIsExpanded = slices.Delete(participant.TaskWhoseInDataFlowsNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
		}
	}
	for _, dataFlow := range task.inDataFlows {
		stager.treeDataFlowsWithinTask(diagramProcess, dataFlow, nodeInDataFlows)
	}
}
