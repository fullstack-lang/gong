package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeSystemes(
	system *System,
	parentNode *tree.Node,
	systemsWhoseNodeIsExpanded *[]*System,
) {
	systemNode := &tree.Node{
		Name:            system.GetName(),
		IsExpanded:      slices.Contains(*systemsWhoseNodeIsExpanded, system),
		IsNodeClickable: true,
		IsInEditMode:    system.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, systemNode)

	addRenameButton(system, systemNode, stager)

	// Button to toggle compounding of subsystem CPEs
	{
		compoundButton := &tree.Button{
			Name:            system.GetName() + " Compound CPE",
			Icon:            string(buttons.BUTTON_layers_clear),
			ToolTipText:     "Enable Compounding Subsystem CPEs",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				system.AreCPEsCompoundedFromSubSystems = !system.AreCPEsCompoundedFromSubSystems
				stager.stage.Commit()
			},
		}
		if system.AreCPEsCompoundedFromSubSystems {
			compoundButton.Icon = string(buttons.BUTTON_layers)
			compoundButton.ToolTipText = "Disable Compounding Subsystem CPEs (Currently Active)"
		}
		systemNode.Buttons = append(systemNode.Buttons, compoundButton)
	}

	systemNode.OnNameChange = stager.onNameChange(system)
	systemNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, system, systemsWhoseNodeIsExpanded)
	systemNode.OnClick = onNodeClicked(stager, system)

	//
	// Diagrams within System
	//
	for _, diagram := range system.DiagramFlossEquations {
		stager.treeDiagramFlossEquationWithinSystem(system, diagram, systemNode)
	}

	confDiagrams := ItemButtonConfiguration[
		DiagramFlossEquation, *DiagramFlossEquation,
		System, *System,
	]{
		parentNode:                         systemNode,
		sliceForNewAddedItem:               &system.DiagramFlossEquations,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   systemsWhoseNodeIsExpanded,
		parentElement:                      system,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagrams)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.Scale = 5.0
		for d_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
			d_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	//
	// SubSystemes
	//
	confSubSystemes := ItemButtonConfiguration[
		System, *System,
		System, *System,
	]{
		parentNode:                         systemNode,
		sliceForNewAddedItem:               &system.SubSystems,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   systemsWhoseNodeIsExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confSubSystemes)

	subSystemesNode := &tree.Node{
		Name:            "SubSystems",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsSubSystemNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, subSystemesNode)
	subSystemesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsSubSystemNodeExpanded)
	subSystemesNode.OnClick = onNodeClicked(stager, system)

	for _, system_ := range system.SubSystems {
		stager.treeSystemes(system_, subSystemesNode, systemsWhoseNodeIsExpanded)
	}
}

func (stager *Stager) treeDiagramFlossEquationWithinSystem(
	system *System,
	diagram *DiagramFlossEquation,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagram.GetName(),
		IsExpanded:        slices.Contains(system.DiagramFlossEquationsWhoseNodeIsExpanded, diagram),
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagram.IsChecked,
		IsInEditMode:      diagram.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	addRenameButton(diagram, diagramNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system.DiagramFlossEquations = slices.DeleteFunc(system.DiagramFlossEquations, func(d *DiagramFlossEquation) bool { return d == diagram })
			diagram.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if diagramNode.Menu == nil {
		diagramNode.Menu = &tree.Menu{Name: "Menu"}
	}
	diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, deleteButton)

	// Button for visibility management of quantitative elements
	{
		quantButton := &tree.Button{
			Name:            diagram.GetName() + " Quantitative Visibility",
			Icon:            string(buttons.BUTTON_123),
			ToolTipText:     "Show quantitative values",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.AreQuantitativeElementsVisible = !diagram.AreQuantitativeElementsVisible
				stager.stage.Commit()
			},
		}
		if diagram.AreQuantitativeElementsVisible {
			quantButton.Icon = string(buttons.BUTTON_123)
			quantButton.ToolTipText = "Hide quantitative values"
		} else {
			quantButton.Icon = string(buttons.BUTTON_visibility_off)
			quantButton.ToolTipText = "Show quantitative values"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, quantButton)
	}

	// Button for visibility management of subsystems breakdown
	if system.AreCPEsCompoundedFromSubSystems {
		subsysButton := &tree.Button{
			Name:            diagram.GetName() + " Subsystems Breakdown Visibility",
			Icon:            string(buttons.BUTTON_account_tree),
			ToolTipText:     "Show Subsystems Breakdown",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.AreSubsystemsVisible = !diagram.AreSubsystemsVisible
				stager.stage.Commit()
			},
		}
		if diagram.AreSubsystemsVisible {
			subsysButton.Icon = string(buttons.BUTTON_account_tree)
			subsysButton.ToolTipText = "Hide Subsystems Breakdown (Display System Only)"
		} else {
			subsysButton.Icon = string(buttons.BUTTON_layers_clear)
			subsysButton.ToolTipText = "Display Breakdown at Subsystem Level"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, subsysButton)
	}

	// Button for toggle between font sizes
	{
		fontSizeButton := &tree.Button{
			Name:            diagram.GetName() + " Font Size Toggle",
			Icon:            string(buttons.BUTTON_format_size),
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				switch diagram.FontSize {
				case FONT_SIZE_SMALL:
					diagram.FontSize = FONT_SIZE_NORMAL
				case FONT_SIZE_NORMAL:
					diagram.FontSize = FONT_SIZE_BIG
				case FONT_SIZE_BIG:
					diagram.FontSize = FONT_SIZE_VERY_BIG
				case FONT_SIZE_VERY_BIG:
					diagram.FontSize = FONT_SIZE_SMALL
				default:
					diagram.FontSize = FONT_SIZE_BIG
				}
				stager.stage.Commit()
			},
		}
		currentSizeLabel := "Normal"
		switch diagram.FontSize {
		case FONT_SIZE_SMALL:
			currentSizeLabel = "Small"
		case FONT_SIZE_BIG:
			currentSizeLabel = "Big"
		case FONT_SIZE_VERY_BIG:
			currentSizeLabel = "Very Big"
		default:
			currentSizeLabel = "Normal"
		}
		fontSizeButton.ToolTipText = "Font size: " + currentSizeLabel + " (click to toggle)"
		diagramNode.Buttons = append(diagramNode.Buttons, fontSizeButton)
	}

	diagramNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			for d_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
				d_.IsChecked = false
			}
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	diagramNode.OnClick = onNodeClicked(stager, diagram)
	diagramNode.OnNameChange = stager.onNameChange(diagram)
	diagramNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagram, &system.DiagramFlossEquationsWhoseNodeIsExpanded)

	//
	// Complexities
	//
	complexitiesNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, complexitiesNode)
	complexitiesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsComplexitysNodeExpanded)
	complexitiesNode.OnClick = onNodeClicked(stager, diagram)

	confComplexities := ItemButtonConfiguration[
		Complexity, *Complexity,
		System, *System,
	]{
		parentNode:                         complexitiesNode,
		sliceForNewAddedItem:               &system.Complexities,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &diagram.IsComplexitysNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confComplexities)

	for _, complexity := range system.Complexities {
		stager.treeComplexityWithinSystem(system, complexity, complexitiesNode)
	}

	//
	// Performances
	//
	performancesNode := &tree.Node{
		Name:            "Performances",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsPerformancesNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, performancesNode)
	performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsPerformancesNodeExpanded)
	performancesNode.OnClick = onNodeClicked(stager, diagram)

	confPerformances := ItemButtonConfiguration[
		Performance, *Performance,
		System, *System,
	]{
		parentNode:                         performancesNode,
		sliceForNewAddedItem:               &system.Performances,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &diagram.IsPerformancesNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confPerformances)

	for _, performance := range system.Performances {
		stager.treePerformanceWithinSystem(system, performance, performancesNode)
	}

	//
	// Efforts
	//
	effortsNode := &tree.Node{
		Name:            "Efforts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsEffortsNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, effortsNode)
	effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsEffortsNodeExpanded)
	effortsNode.OnClick = onNodeClicked(stager, diagram)

	confEfforts := ItemButtonConfiguration[
		Effort, *Effort,
		System, *System,
	]{
		parentNode:                         effortsNode,
		sliceForNewAddedItem:               &system.Efforts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &diagram.IsEffortsNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confEfforts)

	for _, effort := range system.Efforts {
		stager.treeEffortWithinSystem(system, effort, effortsNode)
	}

	//
	// Notes
	//
	{
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsNotesNodeExpanded)

		library := system.GetOwningLibrary()
		if library == nil {
			library = stager.getRootLibrary()
		}

		if library != nil {
			for _, note := range library.RootNotes {
				stager.treeNoteWithinDiagramFlossEquation(diagram, note, notesNode)
			}

			confNote := ItemButtonConfiguration[
				Note, *Note,
				Library, *Library,
			]{
				parentNode:                         notesNode,
				sliceForNewAddedItem:               &library.RootNotes,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsNotesNodeExpanded,
				parentElement:                      library,
			}
			adder := addCreateItemButton(stager, confNote)
			adder.OnBeforeCommit = func() {
				newNote := adder.createdItem
				if newNote != nil {
					noteShape := (&NoteShape{
						Name: newNote.GetName() + " shape",
						Note: newNote,
						RectShape: RectShape{
							X:      100,
							Y:      100 + float64(len(diagram.Note_Shapes))*60,
							Width:  diagram.GetDefaultBoxWidth(),
							Height: diagram.GetDefaultBoxHeigth(),
						},
					}).Stage(stager.stage)
					diagram.Note_Shapes = append(diagram.Note_Shapes, noteShape)
				}
			}
		}
	}
}
