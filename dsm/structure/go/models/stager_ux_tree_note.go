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
			Name: port.GetName(),
		}
		portsNode.Children = append(portsNode.Children, nodePort)
	}
}
