package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeNote(
	library *Library,
	note *Note,
	parentNode *tree.Node,
) {
	noteNode := &tree.Node{
		Name:            note.GetName(),
		IsExpanded:      slices.Contains(library.NotesWhoseNodeIsExpanded, note),
		IsNodeClickable: true,
		IsInEditMode:    note.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, noteNode)

	addRenameButton(note, noteNode, stager)

	noteNode.OnNameChange = stager.onNameChange(note)
	noteNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, note, &library.NotesWhoseNodeIsExpanded)
	noteNode.OnClick = onNodeClicked(stager, note)

	// tasks relarted to the note
	tasksNode := &tree.Node{
		Name:            "Tasks",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsTasksNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, tasksNode)
	tasksNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsTasksNodeExpanded)

	for _, task := range note.Tasks {
		nodeTask := &tree.Node{
			Name:            task.GetName(),
			IsNodeClickable: true,
			IsInEditMode:    task.GetIsInRenameMode(),
		}
		tasksNode.Children = append(tasksNode.Children, nodeTask)
		addRenameButton(task, nodeTask, stager)
		nodeTask.OnNameChange = stager.onNameChange(task)
		nodeTask.OnClick = onNodeClicked(stager, task)
	}

	// control flows related to the note
	controlFlowsNode := &tree.Node{
		Name:            "Control Flows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsControlFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, controlFlowsNode)
	controlFlowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsControlFlowsNodeExpanded)

	for _, controlFlow := range note.ControlFlows {
		nodeControlFlow := &tree.Node{
			Name:            controlFlow.GetName(),
			IsNodeClickable: true,
			IsInEditMode:    controlFlow.GetIsInRenameMode(),
		}
		controlFlowsNode.Children = append(controlFlowsNode.Children, nodeControlFlow)
		addRenameButton(controlFlow, nodeControlFlow, stager)
		nodeControlFlow.OnNameChange = stager.onNameChange(controlFlow)
		nodeControlFlow.OnClick = onNodeClicked(stager, controlFlow)
	}

	// data flows related to the note
	dataFlowsNode := &tree.Node{
		Name:            "Data Flows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsDataFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, dataFlowsNode)
	dataFlowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsDataFlowsNodeExpanded)

	for _, dataFlow := range note.DataFlows {
		nodeDataFlow := &tree.Node{
			Name:            dataFlow.GetName(),
			IsNodeClickable: true,
			IsInEditMode:    dataFlow.GetIsInRenameMode(),
		}
		dataFlowsNode.Children = append(dataFlowsNode.Children, nodeDataFlow)
		addRenameButton(dataFlow, nodeDataFlow, stager)
		nodeDataFlow.OnNameChange = stager.onNameChange(dataFlow)
		nodeDataFlow.OnClick = onNodeClicked(stager, dataFlow)
	}
}
