package models

import (
	"math/rand/v2"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeNoteWithinDiagramFloss(
	diagramFloss *DiagramFloss,
	note *Note,
	parentNode *tree.Node,
) {
	stage := stager.stage
	noteShape, ok := diagramFloss.map_Note_NoteShape[note]

	noteNode := &tree.Node{
		Name:                    note.GetName(),
		IsExpanded:              slices.Contains(diagramFloss.NotesWhoseNodeIsExpanded, note),
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

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete note",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			rootLib := stager.getRootLibrary()
			if rootLib != nil {
				rootLib.RootNotes = slices.DeleteFunc(rootLib.RootNotes, func(n *Note) bool { return n == note })
			}
			note.Unstage(stager.stage)
			if shape, ok := diagramFloss.map_Note_NoteShape[note]; ok {
				shape.UnstageVoid(stager.stage)
				diagramFloss.Note_Shapes = slices.DeleteFunc(diagramFloss.Note_Shapes, func(s *NoteShape) bool { return s == shape })
			}
			stager.stage.Commit()
		},
	}
	if noteNode.Menu == nil {
		noteNode.Menu = &tree.Menu{Name: "Menu"}
	}
	noteNode.Menu.Buttons = append(noteNode.Menu.Buttons, deleteButton)


	if ok {
		visibilityButton := &tree.Button{
			Name:            diagramFloss.GetName(),
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
	noteNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, note, &diagramFloss.NotesWhoseNodeIsExpanded)
	noteNode.OnClick = onNodeClicked(stager, note)
	noteNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			noteShape := (&NoteShape{
				Name: note.GetName() + " shape",
				Note: note,
				RectShape: RectShape{
					X:      100 + rand.Float64()*100.0,
					Y:      100 + rand.Float64()*100.0,
					Width:  diagramFloss.GetDefaultBoxWidth(),
					Height: diagramFloss.GetDefaultBoxHeigth(),
				},
			}).Stage(stager.stage)
			diagramFloss.Note_Shapes = append(diagramFloss.Note_Shapes, noteShape)
			stage.Commit()
			return
		}
		if !isChecked && ok {
			noteShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}

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
		noteComplexityKey := noteComplexityKey{
			Note:       note,
			Complexity: complexity,
		}
		noteComplexityShape, ok := diagramFloss.map_Note_NoteComplexityShape[noteComplexityKey]
		nodeComplexity := &tree.Node{
			Name:                    complexity.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note complexity shape"
				}
				return "Click to create a link shape for this complexity within this diagram"
			}(),
			IsNodeClickable: true,
		}
		nodeComplexity.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				noteComplexityShape := (&NoteComplexityShape{
					Name:       complexity.GetName() + " shape",
					Complexity: complexity,
					Note:       note,
					LinkShape:  LinkShape{},
				}).Stage(stager.stage)
				diagramFloss.NoteComplexityShapes = append(diagramFloss.NoteComplexityShapes, noteComplexityShape)
				stage.Commit()
				return
			}
			if !isChecked && ok {
				noteComplexityShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramFloss.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					noteComplexityShape.SetIsHidden(!noteComplexityShape.GetIsHidden())
					stage.Commit()
				},
			}
			if noteComplexityShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodeComplexity.Buttons = append(nodeComplexity.Buttons, visibilityButton)
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
		notePerformanceKey := notePerformanceKey{
			Note:        note,
			Performance: performance,
		}
		notePerformanceShape, ok := diagramFloss.map_Note_NotePerformanceShape[notePerformanceKey]
		nodePerformance := &tree.Node{
			Name:                    performance.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note performance shape"
				}
				return "Click to create a link shape for this performance within this diagram"
			}(),
			IsNodeClickable: true,
		}
		nodePerformance.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				notePerformanceShape := (&NotePerformanceShape{
					Name:        performance.GetName() + " shape",
					Performance: performance,
					Note:        note,
					LinkShape:   LinkShape{},
				}).Stage(stager.stage)
				diagramFloss.NotePerformanceShapes = append(diagramFloss.NotePerformanceShapes, notePerformanceShape)
				stage.Commit()
				return
			}
			if !isChecked && ok {
				notePerformanceShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramFloss.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					notePerformanceShape.SetIsHidden(!notePerformanceShape.GetIsHidden())
					stage.Commit()
				},
			}
			if notePerformanceShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodePerformance.Buttons = append(nodePerformance.Buttons, visibilityButton)
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
		noteEffortKey := noteEffortKey{
			Note:   note,
			Effort: effort,
		}
		noteEffortShape, ok := diagramFloss.map_Note_NoteEffortShape[noteEffortKey]
		nodeEffort := &tree.Node{
			Name:                    effort.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note effort shape"
				}
				return "Click to create a link shape for this effort within this diagram"
			}(),
			IsNodeClickable: true,
		}
		nodeEffort.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				noteEffortShape := (&NoteEffortShape{
					Name:      effort.GetName() + " shape",
					Effort:    effort,
					Note:      note,
					LinkShape: LinkShape{},
				}).Stage(stager.stage)
				diagramFloss.NoteEffortShapes = append(diagramFloss.NoteEffortShapes, noteEffortShape)
				stage.Commit()
				return
			}
			if !isChecked && ok {
				noteEffortShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramFloss.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					noteEffortShape.SetIsHidden(!noteEffortShape.GetIsHidden())
					stage.Commit()
				},
			}
			if noteEffortShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodeEffort.Buttons = append(nodeEffort.Buttons, visibilityButton)
		}
		effortsNode.Children = append(effortsNode.Children, nodeEffort)
	}
}
