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
		stager.treeDiagramFlossEquation(compareAnalysis, diagram, node)
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
		for d_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
			d_.IsChecked = false
		}
		for d_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
			d_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}
}

func (stager *Stager) treeDiagramFlossEquation(
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

	diagramNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			for d_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
				d_.IsChecked = false
			}
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

		rootLib := stager.getRootLibrary()
		if rootLib != nil {
			for _, note := range rootLib.RootNotes {
				stager.treeNoteWithinDiagramFlossEquation(diagram, note, notesNode)
			}

			confNote := ItemButtonConfiguration[
				Note, *Note,
				Library, *Library,
			]{
				parentNode:                         notesNode,
				sliceForNewAddedItem:               &rootLib.RootNotes,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsNotesNodeExpanded,
				parentElement:                      rootLib,
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


