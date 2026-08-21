package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
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

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete note",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootNotes = slices.DeleteFunc(library.RootNotes, func(n *Note) bool { return n == note })
			note.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if noteNode.Menu == nil {
		noteNode.Menu = &tree.Menu{Name: "Menu"}
	}
	noteNode.Menu.Buttons = append(noteNode.Menu.Buttons, deleteButton)

	noteNode.OnNameChange = stager.onNameChange(note)
	noteNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, note, &library.NotesWhoseNodeIsExpanded)
	noteNode.OnClick = onNodeClicked(stager, note)

	// Complexities related to the note
	complexitiesNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, complexitiesNode)
	complexitiesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsComplexitysNodeExpanded)

	for _, complexity := range note.Complexities {
		nodeComplexity := &tree.Node{
			Name: complexity.GetName(),
		}
		complexitiesNode.Children = append(complexitiesNode.Children, nodeComplexity)
	}

	// Performances related to the note
	performancesNode := &tree.Node{
		Name:            "Performances",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsPerformancesNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, performancesNode)
	performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsPerformancesNodeExpanded)

	for _, performance := range note.Performances {
		nodePerformance := &tree.Node{
			Name: performance.GetName(),
		}
		performancesNode.Children = append(performancesNode.Children, nodePerformance)
	}

	// Efforts related to the note
	effortsNode := &tree.Node{
		Name:            "Efforts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsEffortsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, effortsNode)
	effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsEffortsNodeExpanded)

	for _, effort := range note.Efforts {
		nodeEffort := &tree.Node{
			Name: effort.GetName(),
		}
		effortsNode.Children = append(effortsNode.Children, nodeEffort)
	}
}
