package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	var libraryNode = &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.GetIsInRenameMode(),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.GetRootLibrary() {
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

	libraryNode.OnNameChange = func(newName string) {
		library.Name = newName
		library.SetIsInRenameMode(false)
		stager.stage.Commit()
	}
	libraryNode.OnIsExpandedChange = func(isExpanded bool) {
		library.IsExpanded = isExpanded
		stager.stage.Commit()
	}
	libraryNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(library, GetPointerToGongstructName[*Library]())
		stager.stage.Commit()
	}

	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
	}
	addCreateItemButton(stager, confDiagrams)

	for _, diagram := range library.Diagrams {
		diagramNode := &tree.Node{
			Name:              diagram.Name,
			IsExpanded:        diagram.IsExpanded,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagram.IsChecked,

			IsInEditMode: diagram.GetIsInRenameMode(),
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

		diagramNode.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked {
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
					diagram_.IsChecked = false
				}
				diagram.IsChecked = true
			} else {
				diagram.IsChecked = false
				// reset all ddiagram selection
				for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
					diagram_.IsChecked = false
				}
			}
			stager.stage.Commit()
		}
		diagramNode.OnIsExpandedChange = func(isExpanded bool) {
			diagram.IsExpanded = isExpanded
			stager.stage.Commit()
		}
		diagramNode.OnNameChange = func(newName string) {
			diagram.Name = newName
			diagram.SetIsInRenameMode(false)
			stager.stage.Commit()
		}
		diagramNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
			stager.stage.Commit()
		}

		{
			copyButton := &tree.Button{
				Name:            "Diagram Copy",
				Icon:            string(buttons.BUTTON_copy_all),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Copy Diagram",
				OnClick:        onCopyDiagram(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, copyButton)
		}
		{
			showAllButton := &tree.Button{
				Name:            "Diagram Show All",
				Icon:            string(buttons.BUTTON_all_out),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Show All Elements in the Diagram",
				OnClick:        onShowAllInDiagram(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}
		{
			showAllButton := &tree.Button{
				Name:            "Diagram Show As PBS Tree",
				Icon:            string(buttons.BUTTON_account_tree),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Show Show As PBS Tree",
				OnClick:        onLayoutPBS(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}
		{
			showAllButton := &tree.Button{
				Name:            "Diagram Show As WBS Tree",
				Icon:            string(buttons.BUTTON_account_circle),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Show Show As WBS Tree",
				OnClick:        onLayoutWBS(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}
		{
			showAllButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,

				OnClick: func() {
					diagram.ShowPrefix = !diagram.ShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagram.ShowPrefix {
				showAllButton.Icon = string(buttons.BUTTON_label)
				showAllButton.ToolTipText = "Show Prefix"
			} else {
				showAllButton.Icon = string(buttons.BUTTON_label_off)
				showAllButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}

		wbsNode := &tree.Node{
			Name:            "Concerns",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsConcernsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, wbsNode)
		wbsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsConcernsNodeExpanded)

		confRootConcerns := ItemAndShapeButtonConfiguration[
			Concern, *Concern,
			Library, *Library,
			ConcernShape, *ConcernShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Concern, *Concern,
				Library, *Library,
			]{
				parentNode:                         wbsNode,
				sliceForNewAddedItem:               &library.RootConcerns,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsConcernsNodeExpanded,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Concern_Shapes,
		}
		addCreateItemAndShapeButton(stager, confRootConcerns)

		diagram.map_Concern_ConcernCompositionShape = make(map[*Concern]*ConcernCompositionShape)
		for _, shape := range diagram.ConcernComposition_Shapes {
			if shape.Concern != nil {
				diagram.map_Concern_ConcernCompositionShape[shape.Concern] = shape
			}
		}

		stakeholderNode := &tree.Node{
			Name:            "Stakeholders",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsStakeholdersNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, stakeholderNode)
		stakeholderNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsStakeholdersNodeExpanded)

		confRootStakeholders := ItemAndShapeButtonConfiguration[
			Stakeholder, *Stakeholder,
			Library, *Library,
			StakeholderShape, *StakeholderShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Stakeholder, *Stakeholder,
				Library, *Library,
			]{
				parentNode:                         stakeholderNode,
				sliceForNewAddedItem:               &library.RootStakeholders,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsStakeholdersNodeExpanded,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Stakeholder_Shapes,
		}
		addCreateItemAndShapeButton(stager, confRootStakeholders)

		for _, x := range library.RootStakeholders {
			stager.treeStakeholderBSinDiagram(diagram, x, stakeholderNode)
		}

		for _, x := range library.RootConcerns {
			stager.treeConcernBSinDiagram(diagram, x, wbsNode)
		}

		delivarablesNode := &tree.Node{
			Name:            "Deliverables",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, delivarablesNode)
		delivarablesNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded)

		confRootDeliverables := ItemAndShapeButtonConfiguration[
			Deliverable, *Deliverable,
			Library, *Library,
			ProductShape, *ProductShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Deliverable, *Deliverable,
				Library, *Library,
			]{
				parentNode:                         delivarablesNode,
				sliceForNewAddedItem:               &library.RootDeliverables,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsPBSNodeExpanded,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Product_Shapes,
		}
		addCreateItemAndShapeButton(stager, confRootDeliverables)

		for _, x := range library.RootDeliverables {
			stager.treeDeliverableRecusriveInDiagram(diagram, x, delivarablesNode)
		}

		requirementsNode := &tree.Node{
			Name:            "Requirements",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsRequirementsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, requirementsNode)
		requirementsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsRequirementsNodeExpanded)

		confRootRequirements := ItemButtonConfiguration[
			Requirement, *Requirement,
			Library, *Library,
		]{
			parentNode:                         requirementsNode,
			sliceForNewAddedItem:               &library.RootRequirements,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsRequirementsNodeExpanded,
		}
		addCreateItemButton(stager, confRootRequirements)

		for _, req := range library.RootRequirements {
			stager.treeRequirementBSinDiagram(diagram, req, requirementsNode)
		}

		conceptsNode := &tree.Node{
			Name:            "Concepts",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsConceptsNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, conceptsNode)
		conceptsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsConceptsNodeExpanded)

		confRootConcepts := ItemButtonConfiguration[
			Concept, *Concept,
			Library, *Library,
		]{
			parentNode:                         conceptsNode,
			sliceForNewAddedItem:               &library.RootConcepts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsConceptsNodeExpanded,
		}
		addCreateItemButton(stager, confRootConcepts)

		for _, concept := range library.RootConcepts {
			stager.treeConceptBSinDiagram(diagram, concept, conceptsNode)
		}

		{
			notesNode := &tree.Node{
				Name:            "Notes",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsNotesNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, notesNode)
			notesNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsNotesNodeExpanded)

			confNotes := ItemAndShapeButtonConfiguration[
				Note, *Note,
				Library, *Library,
				NoteShape, *NoteShape,
			]{
				ItemButtonConfiguration: ItemButtonConfiguration[
					Note, *Note,
					Library, *Library,
				]{
					parentNode:                         notesNode,
					sliceForNewAddedItem:               &library.Notes,
					isParentNodeExpandedByAddOperation: true,
					parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
					parentNodeExpansionBooleanValue:    &diagram.IsNotesNodeExpanded,
				},
				receivingDiagram:      diagram,
				sliceForNewAddedShape: &diagram.Note_Shapes,
			}
			addCreateItemAndShapeButton(stager, confNotes)

			for _, note := range library.Notes {
				confNode := TreeNodeAndShapeConfigurationWithoutLink[
					*Note, Note,
					*Note, Note, // Parent is not used
					*NoteShape, NoteShape,
					*Diagram,
				]{
					diagram:                     diagram,
					parentNode:                  notesNode,
					element:                     note,
					parentElement:               nil,
					elementsWhoseNodeIsExpanded: &diagram.NotesWhoseNodeIsExpanded,
					shapes:                      &diagram.Note_Shapes,
					shapesMap:                   diagram.map_Note_NoteShape,
				}
				noteNode := addNodeToTreeWithoutLink(stager, confNode)

				// allow display of associations note to products
				for _, product := range note.Products {
					nodeProduct := &tree.Node{
						Name:            product.Name,
						IsNodeClickable: true,
					}
					noteNode.Children = append(noteNode.Children, nodeProduct)

					showHideRelationButton := &tree.Button{
						Name: GetGongstructNameFromPointer(product) + "- showHideRelationButton" + note.Name + " - " + product.Name,

						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					nodeProduct.Buttons = append(nodeProduct.Buttons, showHideRelationButton)

					if _, ok := diagram.map_Product_ProductShape[product]; ok {
						if _, ok := diagram.map_Note_NoteShape[note]; ok {

							noteProductShape, ok := diagram.map_Note_NoteProductShape[noteProductKey{Note: note, Product: product}]
							nodeProduct.IsChecked = ok

							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteProductShape, &diagram.NoteProductShapes)() }
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, product, &diagram.NoteProductShapes)() }
							}
						}
					}
				}

				for _, task := range note.Tasks {
					nodeTask := &tree.Node{
						Name:            task.Name,
						IsNodeClickable: true,
					}
					noteNode.Children = append(noteNode.Children, nodeTask)
					showHideRelationButton := &tree.Button{
						Name:            GetGongstructNameFromPointer(task),
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					nodeTask.Buttons = append(nodeTask.Buttons, showHideRelationButton)
					if _, ok := diagram.map_Concern_ConcernShape[task]; ok {
						if _, ok := diagram.map_Note_NoteShape[note]; ok {
							noteTaskShape, ok := diagram.map_Note_NoteTaskShape[noteTaskKey{Note: note, Task: task}]
							nodeTask.IsChecked = ok

							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteTaskShape, &diagram.NoteTaskShapes)() }
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, task, &diagram.NoteTaskShapes)() }
							}
						}
					}
				}

				for _, resource := range note.Resources {
					nodeResource := &tree.Node{
						Name:            resource.Name,
						IsNodeClickable: true,
					}
					noteNode.Children = append(noteNode.Children, nodeResource)
					showHideRelationButton := &tree.Button{
						Name:            GetGongstructNameFromPointer(resource),
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					nodeResource.Buttons = append(nodeResource.Buttons, showHideRelationButton)
					if _, ok := diagram.map_Stakeholder_StakeholderShape[resource]; ok {
						if _, ok := diagram.map_Note_NoteShape[note]; ok {
							noteResourceShape, ok := diagram.map_Note_NoteResourceShape[noteResourceKey{Note: note, Stakeholder: resource}]
							nodeResource.IsChecked = ok
							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteResourceShape, &diagram.NoteResourceShapes)() }
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, resource, &diagram.NoteResourceShapes)() }
							}
						}
					}
				}
			}
		}
	}

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(treeInstance, subLibrary, &libraryNode.Children)
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(isExpanded bool) {
	return func(newIsExpanded bool) {
		if *isExpanded != newIsExpanded {
			*isExpanded = newIsExpanded
		}
		stager.stage.Commit()
	}
}
