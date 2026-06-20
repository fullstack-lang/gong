package models

import (
	"math/rand/v2"
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
				RectShape: RectShape{
					X:      100 + rand.Float64()*100.0,
					Y:      100 + rand.Float64()*100.0,
					Width:  diagramProcess.GetDefaultBoxWidth(),
					Height: diagramProcess.GetDefaultBoxHeigth(),
				},
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
		noteTaskKey := noteTaskKey{
			Note: note,
			Task: task,
		}
		_, ok := diagramProcess.map_Note_NoteTaskShape[noteTaskKey]
		nodeTask := &tree.Node{
			Name:                    task.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note task shape"
				}
				return "Click to create a note task shape for this task within this diagram"
			}(),
			IsNodeClickable: true,
		}
		nodeTask.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				noteTaskShape := (&NoteTaskShape{
					Name:      task.GetName() + " shape",
					Task:      task,
					Note:      note,
					LinkShape: LinkShape{},
				}).Stage(stager.stage)
				diagramProcess.NoteTaskShapes = append(diagramProcess.NoteTaskShapes, noteTaskShape)
				stage.Commit()
				return
			}
		}
		tasksNode.Children = append(tasksNode.Children, nodeTask)
	}
}
