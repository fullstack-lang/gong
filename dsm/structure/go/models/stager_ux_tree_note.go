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

	// ports relarted to the note
	portsNode := &tree.Node{
		Name:            "Ports",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsPortsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, portsNode)
	portsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsPortsNodeExpanded)

	for _, port := range note.Ports {
		nodePort := &tree.Node{
			Name:            port.GetName(),
			IsNodeClickable: true,
			IsInEditMode:    port.GetIsInRenameMode(),
		}
		portsNode.Children = append(portsNode.Children, nodePort)
		addRenameButton(port, nodePort, stager)
		nodePort.OnNameChange = stager.onNameChange(port)
		nodePort.OnClick = onNodeClicked(stager, port)
	}

	// parts related to the note
	partsNode := &tree.Node{
		Name:            "Parts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsPartsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, partsNode)
	partsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsPartsNodeExpanded)

	for _, part := range note.Parts {
		nodePart := &tree.Node{
			Name:            part.GetName(),
			IsNodeClickable: true,
			IsInEditMode:    part.GetIsInRenameMode(),
		}
		partsNode.Children = append(partsNode.Children, nodePart)
		addRenameButton(part, nodePart, stager)
		nodePart.OnNameChange = stager.onNameChange(part)
		nodePart.OnClick = onNodeClicked(stager, part)
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
