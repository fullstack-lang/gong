package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeNoteWithinDiagramProcess(
	diagramProcess *DiagramProcess,
	note *Note,
	parentNode *tree.Node,
) {
	stage := stager.stage
	noteShape, ok := diagramProcess.map_Note_NoteShape[note]

	noteNode := &tree.Node{
		Name:                    note.GetName(),
		IsExpanded:              slices.Contains(diagramProcess.NotesWhoseNodeIsExpanded, note),
		IsNodeClickable:         true,
		IsInEditMode:            note.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the note shape"
			}
			return "Click to create a note shape for this note within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, noteNode)
	addRenameButton(note, noteNode, stager)

	if ok {
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				noteShape.SetIsHidden(!noteShape.GetIsHidden())
				stage.Commit()
			},
		}
		if noteShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		noteNode.Buttons = append(noteNode.Buttons, visibilityButton)
	}

	noteNode.OnNameChange = stager.onNameChange(note)
	noteNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, note, &diagramProcess.NotesWhoseNodeIsExpanded)
	noteNode.OnClick = onNodeClicked(stager, note)
	noteNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			noteShape := (&NoteShape{
				Name: note.GetName() + " shape",
				Note: note,
			}).Stage(stager.stage)
			diagramProcess.Note_Shapes = append(diagramProcess.Note_Shapes, noteShape)
			stage.Commit()
			return
		}
		if !isChecked && ok {
			noteShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}
}
