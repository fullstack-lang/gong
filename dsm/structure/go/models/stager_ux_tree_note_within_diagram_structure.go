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
				Name:   note.GetName() + " shape",
				Note:   note,
				X:      100 + rand.Float64()*100.0,
				Y:      100 + rand.Float64()*100.0,
				Width:  diagramStructure.GetDefaultBoxWidth(),
				Height: diagramStructure.GetDefaultBoxHeigth(),
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
		notePortShape, ok := diagramStructure.map_Note_NotePortShape[notePortKey]
		nodePort := &tree.Node{
			Name:                    port.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			IsInEditMode:            port.GetIsInRenameMode(),
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
		addRenameButton(port, nodePort, stager)
		nodePort.OnNameChange = stager.onNameChange(port)
		nodePort.OnClick = onNodeClicked(stager, port)
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
			if !isChecked && ok {
				notePortShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramStructure.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					notePortShape.SetIsHidden(!notePortShape.GetIsHidden())
					stage.Commit()
				},
			}
			if notePortShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodePort.Buttons = append(nodePort.Buttons, visibilityButton)
		}
		portsNode.Children = append(portsNode.Children, nodePort)
	}

	// parts relarted to the note
	partsNode := &tree.Node{
		Name:            "Parts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      note.IsPartsNodeExpanded,
		IsNodeClickable: true,
	}
	noteNode.Children = append(noteNode.Children, partsNode)
	partsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&note.IsPartsNodeExpanded)

	for _, part := range note.Parts {
		notePartKey := notePartKey{
			Note: note,
			Part: part,
		}
		notePartShape, ok := diagramStructure.map_Note_NotePartShape[notePartKey]
		nodePart := &tree.Node{
			Name:                    part.GetName(),
			HasCheckboxButton:       true,
			IsChecked:               ok,
			IsInEditMode:            part.GetIsInRenameMode(),
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if ok {
					return "Click to remove the note part shape"
				}
				return "Click to create a note part shape for this part within this diagram"
			}(),
			IsNodeClickable: true,
		}
		addRenameButton(part, nodePart, stager)
		nodePart.OnNameChange = stager.onNameChange(part)
		nodePart.OnClick = onNodeClicked(stager, part)
		nodePart.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !ok {
				notePartShape := (&NotePartShape{
					Name:      part.GetName() + " shape",
					Part:      part,
					Note:      note,
					LinkShape: LinkShape{},
				}).Stage(stager.stage)
				diagramStructure.NotePartShapes = append(diagramStructure.NotePartShapes, notePartShape)
				stage.Commit()
				return
			}
			if !isChecked && ok {
				notePartShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramStructure.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					notePartShape.SetIsHidden(!notePartShape.GetIsHidden())
					stage.Commit()
				},
			}
			if notePartShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			nodePart.Buttons = append(nodePart.Buttons, visibilityButton)
		}
		partsNode.Children = append(partsNode.Children, nodePart)
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
		noteDataFlowShape, ok := diagramStructure.map_Note_NoteDataFlowShape[noteDataFlowKey]
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
				noteDataFlowShape := (&NoteDataFlowShape{
					Name:      note.GetName() + " to " + dataFlow.GetName(),
					DataFlow:  dataFlow,
					Note:      note,
					LinkShape: LinkShape{},
				}).Stage(stager.stage)
				diagramStructure.NoteDataFlowShapes = append(diagramStructure.NoteDataFlowShapes, noteDataFlowShape)
				stage.Commit()
				return
			}
			if !isChecked && ok {
				noteDataFlowShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		if ok {
			visibilityButton := &tree.Button{
				Name:            diagramStructure.GetName(),
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
