package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeTask(
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
	nodeOutControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, task, &participant.TaskWhoseOutControlFlowsNodeIsExpanded)
	for _, controlFlow := range task.outControlFlows {
		stager.treeControlFlowsWithinTask(diagramProcess, controlFlow, nodeOutControlFlows)
	}

	nodeInControlFlows := &tree.Node{
		Name:            "in control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseInControlFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeInControlFlows)
	nodeInControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, task, &participant.TaskWhoseInControlFlowsNodeIsExpanded)
	for _, controlFlow := range task.inControlFlows {
		stager.treeControlFlowsWithinTask(diagramProcess, controlFlow, nodeInControlFlows)
	}

	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseOutDataFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, task, &participant.TaskWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range task.outDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(participant.TaskWhoseInDataFlowsNodeIsExpanded, task),
	}
	taskNode.Children = append(taskNode.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, task, &participant.TaskWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range task.inDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeInDataFlows)
	}
}
