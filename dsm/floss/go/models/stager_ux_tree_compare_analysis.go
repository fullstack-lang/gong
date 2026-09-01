package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeCompareAnalysisWithinLibrary(
	library *Library,
	compareAnalysis *CompareAnalysis,
	parentNode *tree.Node,
) {
	node := &tree.Node{
		Name:            compareAnalysis.GetName(),
		IsExpanded:      slices.Contains(library.CompareAnalysisWhoseNodeIsExpanded, compareAnalysis),
		IsNodeClickable: true,
		IsInEditMode:    compareAnalysis.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(compareAnalysis, node, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete compare analysis",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootCompareAnalysis = slices.DeleteFunc(library.RootCompareAnalysis, func(ca *CompareAnalysis) bool { return ca == compareAnalysis })
			compareAnalysis.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}
	node.Menu.Buttons = append(node.Menu.Buttons, deleteButton)

	node.OnNameChange = stager.onNameChange(compareAnalysis)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, compareAnalysis, &library.CompareAnalysisWhoseNodeIsExpanded)
	node.OnClick = onNodeClicked(stager, compareAnalysis)

	//
	// Diagrams within CompareAnalysis
	//
	for _, diagram := range compareAnalysis.DiagramFlossEquations {
		stager.treeDiagramFlossEquation(library, compareAnalysis, diagram, node)
	}

	confDiagrams := ItemButtonConfiguration[
		DiagramFlossEquation, *DiagramFlossEquation,
		CompareAnalysis, *CompareAnalysis,
	]{
		parentNode:                         node,
		sliceForNewAddedItem:               &compareAnalysis.DiagramFlossEquations,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   &library.CompareAnalysisWhoseNodeIsExpanded,
		parentElement:                      compareAnalysis,
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
}

func (stager *Stager) treeDiagramFlossEquation(
	library *Library,
	compareAnalysis *CompareAnalysis,
	diagram *DiagramFlossEquation,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagram.GetName(),
		IsExpanded:        slices.Contains(compareAnalysis.DiagramFlossEquationsWhoseNodeIsExpanded, diagram),
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagram.IsChecked,
		IsInEditMode:      diagram.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	addRenameButton(diagram, diagramNode, stager)

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
	{
		hasCompounded := (compareAnalysis.ToSystem != nil && compareAnalysis.ToSystem.AreCPEsCompoundedFromSubSystems) ||
			(compareAnalysis.FromSystem != nil && compareAnalysis.FromSystem.AreCPEsCompoundedFromSubSystems)
		if hasCompounded {
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
	}

	// Button for toggle between 6 columns (pairs) and 3 columns (delta computation)
	{
		deltaColButton := &tree.Button{
			Name:            diagram.GetName() + " Delta Columns Mode Toggle",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.IsInDelta3ColumnsMode = !diagram.IsInDelta3ColumnsMode
				stager.stage.Commit()
			},
		}
		if diagram.IsInDelta3ColumnsMode {
			deltaColButton.Icon = string(buttons.BUTTON_view_week)
			deltaColButton.ToolTipText = "Switch to 6 columns mode (V1 & V2 pairs)"
		} else {
			deltaColButton.Icon = string(buttons.BUTTON_view_column)
			deltaColButton.ToolTipText = "Switch to 3 columns Delta mode (V2 - V1 computation)"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, deltaColButton)
	}

	// Button for toggle hiding common elements between source & target
	{
		commonElementsButton := &tree.Button{
			Name:            diagram.GetName() + " Common Elements Visibility Toggle",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.AreCommonElementsHidden = !diagram.AreCommonElementsHidden
				stager.stage.Commit()
			},
		}
		if diagram.AreCommonElementsHidden {
			commonElementsButton.Icon = string(buttons.BUTTON_difference)
			commonElementsButton.ToolTipText = "Show common elements between source & target (Currently Hidden)"
		} else {
			commonElementsButton.Icon = string(buttons.BUTTON_filter_alt_off)
			commonElementsButton.ToolTipText = "Hide common elements between source & target"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, commonElementsButton)
	}

	// Button for toggle diff arrows (Added complexity/performance/effort notes & links)
	{
		cpeArrowsButton := &tree.Button{
			Name:            diagram.GetName() + " Diff Arrows Visibility Toggle",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.AreCPEArrowsVisible = !diagram.AreCPEArrowsVisible
				stager.stage.Commit()
			},
		}
		if diagram.AreCPEArrowsVisible {
			cpeArrowsButton.Icon = string(buttons.BUTTON_trending_up)
			cpeArrowsButton.ToolTipText = "Hide Diff Arrows (Complexity/Performance/Effort)"
		} else {
			cpeArrowsButton.Icon = string(buttons.BUTTON_trending_flat)
			cpeArrowsButton.ToolTipText = "Show Diff Arrows (Complexity/Performance/Effort)"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, cpeArrowsButton)
	}

	// Button for visibility management of column titles (Complexity, Performance, Effort)
	{
		colTitlesButton := &tree.Button{
			Name:            diagram.GetName() + " Column Titles Visibility",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				diagram.AreColumnTitlesVisible = !diagram.AreColumnTitlesVisible
				stager.stage.Commit()
			},
		}
		if diagram.AreColumnTitlesVisible {
			colTitlesButton.Icon = string(buttons.BUTTON_label)
			colTitlesButton.ToolTipText = "Hide Column Titles (Complexity, Performance, Effort)"
		} else {
			colTitlesButton.Icon = string(buttons.BUTTON_label_off)
			colTitlesButton.ToolTipText = "Show Column Titles (Complexity, Performance, Effort)"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, colTitlesButton)
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
	diagramNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagram, &compareAnalysis.DiagramFlossEquationsWhoseNodeIsExpanded)

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

		if library == nil {
			library = compareAnalysis.GetOwningLibrary()
		}
		if library == nil {
			library = stager.getRootLibrary()
		}

		if library != nil {
			notes := stager.collectAllNotesForDiagram(library, diagram)
			for _, note := range notes {
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
