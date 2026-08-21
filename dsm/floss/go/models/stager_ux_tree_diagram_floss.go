package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramFloss(
	system *System,
	diagramFloss *DiagramFloss,
	parentNode *tree.Node,
) {
	diagramNode := &tree.Node{
		Name:              diagramFloss.Name,
		IsExpanded:        slices.Contains(system.DiagramFlossWhoseNodeIsExpanded, diagramFloss) == true,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagramFloss.IsChecked,

		IsInEditMode: diagramFloss.isInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, diagramNode)

	element := diagramFloss
	node := diagramNode

	addRenameButton(element, node, stager)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			// uncheck all diagrams
			for diagramFloss_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
				diagramFloss_.IsChecked = false
			}
			for diagramEquation_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
				diagramEquation_.IsChecked = false
			}

			diagramFloss.IsChecked = true
			stager.stage.Commit()
			return
		} else {
			diagramFloss.IsChecked = false
			stager.stage.Commit()
			return
		}
	}

	node.OnClick = onNodeClicked(stager, diagramFloss)
	node.OnNameChange = stager.onNameChange(diagramFloss)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, diagramFloss, &system.DiagramFlossWhoseNodeIsExpanded)

	// prefix button
	{
		showPrefixButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagramFloss.IsShowPrefix = !diagramFloss.IsShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagramFloss.IsShowPrefix {
			showPrefixButton.Icon = string(buttons.BUTTON_label)
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
	}

	//
	// Complexities
	//
	{
		complexitiesNode := &tree.Node{
			Name:            "Complexities",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramFloss.IsComplexitysNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, complexitiesNode)
		complexitiesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramFloss.IsComplexitysNodeExpanded)

		for _, complexity := range system.Complexities {
			stager.treeComplexityWithinDiagramFloss(diagramFloss, system, complexity, complexitiesNode)
		}

		confComplexity := ItemButtonConfiguration[
			Complexity, *Complexity,
			System, *System,
		]{
			parentNode:                         complexitiesNode,
			sliceForNewAddedItem:               &system.Complexities,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramFloss.IsComplexitysNodeExpanded,
			parentElement:                      system,
		}
		addCreateItemButton(stager, confComplexity)
	}

	//
	// Performances
	//
	{
		performancesNode := &tree.Node{
			Name:            "Performances",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramFloss.IsPerformancesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, performancesNode)
		performancesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramFloss.IsPerformancesNodeExpanded)

		for _, performance := range system.Performances {
			stager.treePerformanceWithinDiagramFloss(diagramFloss, system, performance, performancesNode)
		}

		confPerformance := ItemButtonConfiguration[
			Performance, *Performance,
			System, *System,
		]{
			parentNode:                         performancesNode,
			sliceForNewAddedItem:               &system.Performances,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramFloss.IsPerformancesNodeExpanded,
			parentElement:                      system,
		}
		addCreateItemButton(stager, confPerformance)
	}

	//
	// Efforts
	//
	{
		effortsNode := &tree.Node{
			Name:            "Efforts",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramFloss.IsEffortsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, effortsNode)
		effortsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramFloss.IsEffortsNodeExpanded)

		for _, effort := range system.Efforts {
			stager.treeEffortWithinDiagramFloss(diagramFloss, system, effort, effortsNode)
		}

		confEffort := ItemButtonConfiguration[
			Effort, *Effort,
			System, *System,
		]{
			parentNode:                         effortsNode,
			sliceForNewAddedItem:               &system.Efforts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagramFloss.IsEffortsNodeExpanded,
			parentElement:                      system,
		}
		addCreateItemButton(stager, confEffort)
	}

	//
	// Notes
	//
	{
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramFloss.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramFloss.IsNotesNodeExpanded)

		library := system.GetOwningLibrary()
		if library == nil {
			library = stager.getRootLibrary()
		}

		if library != nil {
			for _, note := range library.RootNotes {
				stager.treeNoteWithinDiagramFloss(diagramFloss, note, notesNode)
			}

			confNote := ItemButtonConfiguration[
				Note, *Note,
				Library, *Library,
			]{
				parentNode:                         notesNode,
				sliceForNewAddedItem:               &library.RootNotes,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagramFloss.IsNotesNodeExpanded,
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
							Y:      100 + float64(len(diagramFloss.Note_Shapes))*60,
							Width:  diagramFloss.GetDefaultBoxWidth(),
							Height: diagramFloss.GetDefaultBoxHeigth(),
						},
					}).Stage(stager.stage)
					diagramFloss.Note_Shapes = append(diagramFloss.Note_Shapes, noteShape)
				}
			}
		}
	}
}


