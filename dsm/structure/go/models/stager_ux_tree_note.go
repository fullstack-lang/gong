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
			Name: task.GetName(),
		}
		tasksNode.Children = append(tasksNode.Children, nodeTask)
	}
}
