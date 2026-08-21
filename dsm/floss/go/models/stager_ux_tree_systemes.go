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
	// Complexities
	//
	complexitiesNode := &tree.Node{
		Name:            "Complexities",
		FontStyle:       tree.ITALIC,
		IsExpanded:      system.IsComplexitysNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, complexitiesNode)
	complexitiesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsComplexitysNodeExpanded)
	complexitiesNode.OnClick = onNodeClicked(stager, system)

	confComplexities := ItemButtonConfiguration[
		Complexity, *Complexity,
		System, *System,
	]{
		parentNode:                         complexitiesNode,
		sliceForNewAddedItem:               &system.Complexities,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsComplexitysNodeExpanded,
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
		IsExpanded:      system.IsPerformancesNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, performancesNode)
	performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsPerformancesNodeExpanded)
	performancesNode.OnClick = onNodeClicked(stager, system)

	confPerformances := ItemButtonConfiguration[
		Performance, *Performance,
		System, *System,
	]{
		parentNode:                         performancesNode,
		sliceForNewAddedItem:               &system.Performances,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsPerformancesNodeExpanded,
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
		IsExpanded:      system.IsEffortsNodeExpanded,
		IsNodeClickable: true,
	}
	systemNode.Children = append(systemNode.Children, effortsNode)
	effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&system.IsEffortsNodeExpanded)
	effortsNode.OnClick = onNodeClicked(stager, system)

	confEfforts := ItemButtonConfiguration[
		Effort, *Effort,
		System, *System,
	]{
		parentNode:                         effortsNode,
		sliceForNewAddedItem:               &system.Efforts,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &system.IsEffortsNodeExpanded,
		parentElement:                      system,
	}
	addCreateItemButton(stager, confEfforts)

	for _, effort := range system.Efforts {
		stager.treeEffortWithinSystem(system, effort, effortsNode)
	}

	//
	// SubSystemes
	//
	confSubSystemes := ItemButtonConfiguration[
		System, *System,
		System, *System,
	]{
		parentNode:                         systemNode,
		sliceForNewAddedItem:               &system.SubSystemes,
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

	for _, system_ := range system.SubSystemes {
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
