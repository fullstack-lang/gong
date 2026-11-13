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
				Name:   note.GetName() + " shape",
				Note:   note,
				X:      100 + rand.Float64()*100.0,
				Y:      100 + rand.Float64()*100.0,
				Width:  diagramProcess.GetDefaultBoxWidth(),
				Height: diagramProcess.GetDefaultBoxHeigth(),
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
		noteTaskShape, ok := diagramProcess.map_Note_NoteTaskShape[noteTaskKey]
		nodeTask := &tree.Node{
			Name:                    task.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			IsInEditMode:            task.GetIsInRenameMode(),
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
		addRenameButton(task, nodeTask, stager)
		nodeTask.OnNameChange = stager.onNameChange(task)
		nodeTask.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(task, "Task")
			if ok && noteTaskShape != nil {
				fillUpShapeAppearanceForm(stager, noteTaskShape)
			} else {
				clearShapeAppearanceForm(stager)
			}
			stager.stage.CommitWithSuspendedCallbacks()
		}
		nodeTask.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				addAssociationShapeToDiagram(stager, note, task, &diagramProcess.NoteTaskShapes)
				stage.Commit()
				return
			}
			if !isChecked && ok && noteTaskShape != nil {
				idx := slices.Index(diagramProcess.NoteTaskShapes, noteTaskShape)
				if idx != -1 {
					diagramProcess.NoteTaskShapes = slices.Delete(diagramProcess.NoteTaskShapes, idx, idx+1)
				}
				noteTaskShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramProcess.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					noteTaskShape.SetIsHidden(!noteTaskShape.GetIsHidden())
					stage.Commit()
				},
			}
			if noteTaskShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodeTask.Buttons = append(nodeTask.Buttons, visibilityButton)
		}
		tasksNode.Children = append(tasksNode.Children, nodeTask)
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
		noteControlFlowKey := noteControlFlowKey{
			Note:        note,
			ControlFlow: controlFlow,
		}
		noteControlFlowShape, ok := diagramProcess.map_Note_NoteControlFlowShape[noteControlFlowKey]
		nodeControlFlow := &tree.Node{
			Name:                    controlFlow.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			IsInEditMode:            controlFlow.GetIsInRenameMode(),
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note control flow shape"
				}
				return "Click to create a note control flow shape for this control flow within this diagram"
			}(),
			IsNodeClickable: true,
		}
		addRenameButton(controlFlow, nodeControlFlow, stager)
		nodeControlFlow.OnNameChange = stager.onNameChange(controlFlow)
		nodeControlFlow.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(controlFlow, "ControlFlow")
			if ok && noteControlFlowShape != nil {
				fillUpShapeAppearanceForm(stager, noteControlFlowShape)
			} else {
				clearShapeAppearanceForm(stager)
			}
			stager.stage.CommitWithSuspendedCallbacks()
		}
		nodeControlFlow.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				addAssociationShapeToDiagram(stager, note, controlFlow, &diagramProcess.NoteControlFlowShapes)
				stage.Commit()
				return
			}
			if !isChecked && ok && noteControlFlowShape != nil {
				idx := slices.Index(diagramProcess.NoteControlFlowShapes, noteControlFlowShape)
				if idx != -1 {
					diagramProcess.NoteControlFlowShapes = slices.Delete(diagramProcess.NoteControlFlowShapes, idx, idx+1)
				}
				noteControlFlowShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramProcess.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					noteControlFlowShape.SetIsHidden(!noteControlFlowShape.GetIsHidden())
					stage.Commit()
				},
			}
			if noteControlFlowShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodeControlFlow.Buttons = append(nodeControlFlow.Buttons, visibilityButton)
		}
		controlFlowsNode.Children = append(controlFlowsNode.Children, nodeControlFlow)
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
		noteDataFlowKey := noteDataFlowKey{
			Note:     note,
			DataFlow: dataFlow,
		}
		noteDataFlowShape, ok := diagramProcess.map_Note_NoteDataFlowShape[noteDataFlowKey]
		nodeDataFlow := &tree.Node{
			Name:                    dataFlow.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			IsInEditMode:            dataFlow.GetIsInRenameMode(),
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note data flow shape"
				}
				return "Click to create a note data flow shape for this data flow within this diagram"
			}(),
			IsNodeClickable: true,
		}
		addRenameButton(dataFlow, nodeDataFlow, stager)
		nodeDataFlow.OnNameChange = stager.onNameChange(dataFlow)
		nodeDataFlow.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(dataFlow, "DataFlow")
			if ok && noteDataFlowShape != nil {
				fillUpShapeAppearanceForm(stager, noteDataFlowShape)
			} else {
				clearShapeAppearanceForm(stager)
			}
			stager.stage.CommitWithSuspendedCallbacks()
		}
		nodeDataFlow.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				addAssociationShapeToDiagram(stager, note, dataFlow, &diagramProcess.NoteDataFlowShapes)
				stage.Commit()
				return
			}
			if !isChecked && ok && noteDataFlowShape != nil {
				idx := slices.Index(diagramProcess.NoteDataFlowShapes, noteDataFlowShape)
				if idx != -1 {
					diagramProcess.NoteDataFlowShapes = slices.Delete(diagramProcess.NoteDataFlowShapes, idx, idx+1)
				}
				noteDataFlowShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramProcess.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					noteDataFlowShape.SetIsHidden(!noteDataFlowShape.GetIsHidden())
					stage.Commit()
				},
			}
			if noteDataFlowShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodeDataFlow.Buttons = append(nodeDataFlow.Buttons, visibilityButton)
		}
		dataFlowsNode.Children = append(dataFlowsNode.Children, nodeDataFlow)
	}
}
