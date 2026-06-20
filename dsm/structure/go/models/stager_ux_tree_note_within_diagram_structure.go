package models

import (
	"math/rand/v2"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeNoteWithinDiagramStructure(
	diagramStructure *DiagramStructure,
	note *Note,
	parentNode *tree.Node,
) {
	stage := stager.stage
	noteShape, ok := diagramStructure.map_Note_NoteShape[note]

	noteNode := &tree.Node{
		Name:                    note.GetName(),
		IsExpanded:              slices.Contains(diagramStructure.NotesWhoseNodeIsExpanded, note),
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
			Name:            diagramStructure.GetName(),
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
	noteNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, note, &diagramStructure.NotesWhoseNodeIsExpanded)
	noteNode.OnClick = onNodeClicked(stager, note)
	noteNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			noteShape := (&NoteShape{
				Name: note.GetName() + " shape",
				Note: note,
				RectShape: RectShape{
					X:      100 + rand.Float64()*100.0,
					Y:      100 + rand.Float64()*100.0,
					Width:  diagramStructure.GetDefaultBoxWidth(),
					Height: diagramStructure.GetDefaultBoxHeigth(),
				},
			}).Stage(stager.stage)
			diagramStructure.Note_Shapes = append(diagramStructure.Note_Shapes, noteShape)
			stage.Commit()
			return
		}
		if !isChecked && ok {
			noteShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}
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
		notePortKey := notePortKey{
			Note: note,
			Port: port,
		}
		_, ok := diagramStructure.map_Note_NotePortShape[notePortKey]
		nodePort := &tree.Node{
			Name:                    port.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note port shape"
				}
				return "Click to create a note port shape for this port within this diagram"
			}(),
			IsNodeClickable: true,
		}
		nodePort.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				notePortShape := (&NotePortShape{
					Name:      port.GetName() + " shape",
					Port:      port,
					Note:      note,
					LinkShape: LinkShape{},
				}).Stage(stager.stage)
				diagramStructure.NotePortShapes = append(diagramStructure.NotePortShapes, notePortShape)
				stage.Commit()
				return
			}
		}
		portsNode.Children = append(portsNode.Children, nodePort)
	}
}
