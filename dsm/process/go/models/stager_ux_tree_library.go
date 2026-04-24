package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.isExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.rootLibrary {
		if !library.GetIsInRenameMode() {
			libraryNode.Buttons = append(libraryNode.Buttons,
				&tree.Button{
					Name: library.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnClick: func() {
						library.SetIsInRenameMode(true)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Rename the " + GetGongstructNameFromPointer(library),
					ToolTipPosition: tree.Above,
				})
		} else {
			libraryNode.Buttons = append(libraryNode.Buttons,
				&tree.Button{
					Name: library.GetName() + " " + string(buttons.BUTTON_edit_off),
					Icon: string(buttons.BUTTON_edit_off),
					OnClick: func() {
						library.SetIsInRenameMode(false)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Cancel renaming",
					ToolTipPosition: tree.Above,
				})
		}
	}

	libraryNode.OnUpdate = stager.OnUpdateLibrary(library)
	addAddItemButtonSimple(stager, nil, nil, &library.isExpanded, libraryNode, &library.SubLibraries)

	itemAdderCallback := addAddItemButtonSimple(stager, nil, nil, &library.isExpanded, libraryNode, &library.DiagramProcesss)

	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.isExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagram := range library.DiagramProcesss {
		diagramNode := &tree.Node{
			Name:              diagram.Name,
			IsExpanded:        diagram.isExpanded,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagram.IsChecked,

			IsInEditMode: diagram.isInRenameMode,
		}
		libraryNode.Children = append(libraryNode.Children, diagramNode)

		element := diagram
		node := diagramNode

		if !element.GetIsInRenameMode() {
			node.Buttons = append(node.Buttons,
				&tree.Button{
					Name: element.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnClick: func() {
						element.SetIsInRenameMode(true)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Rename the " + GetGongstructNameFromPointer(element),
					ToolTipPosition: tree.Above,
				})
		} else {
			node.Buttons = append(node.Buttons,
				&tree.Button{
					Name: element.GetName() + " " + string(buttons.BUTTON_edit_off),
					Icon: string(buttons.BUTTON_edit_off),
					OnClick: func() {
						element.SetIsInRenameMode(false)
						stager.stage.Commit()
					},
					HasToolTip:      true,
					ToolTipText:     "Cancel renaming",
					ToolTipPosition: tree.Above,
				})
		}

		diagramNode.OnUpdate = stager.OnUpdateDiagram(diagram)

		{
			showAllButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,

				OnClick: func() {
					diagram.IsShowPrefix = !diagram.IsShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagram.IsShowPrefix {
				showAllButton.Icon = string(buttons.BUTTON_label)
				showAllButton.ToolTipText = "Show Prefix"
			} else {
				showAllButton.Icon = string(buttons.BUTTON_label_off)
				showAllButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}

		processesNode := &tree.Node{
			Name:            "PBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsProcesssNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, processesNode)
		processesNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsProcesssNodeExpanded)

		addAddItemButton(stager, nil, nil, &diagram.IsProcesssNodeExpanded, processesNode, &library.RootProcesses, diagram, &diagram.Process_Shapes, &diagram.ProcessComposition_Shapes)

		for _, product := range library.RootProcesses {
			stager.treeProcessesWithinDiagram(diagram, product, processesNode)
		}
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}
}

// Helper callbacks

func (stager *Stager) OnUpdateLibrary(library *Library) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			library.isExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			library.Name = frontNode.Name
			library.isInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(library, GetPointerToGongstructName[*Library]())
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			*isExpanded = !*isExpanded
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateDiagram(diagram *DiagramProcess) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsChecked && !stagedNode.IsChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
			stagedNode.IsChecked = frontNode.IsChecked
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			diagram.IsChecked = false
			stagedNode.IsChecked = frontNode.IsChecked
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
				diagram_.IsChecked = false
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			diagram.isExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			diagram.Name = frontNode.Name
			diagram.isInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
	}
}
