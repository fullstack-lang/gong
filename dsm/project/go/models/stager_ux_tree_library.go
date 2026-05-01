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
		addRenameButton(library, libraryNode, stager)
	}

	libraryNode.OnUpdate = stager.OnUpdateLibrary(library)

	confSubLibraries := ItemButtonConfiguration[
		Library, *Library, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.isExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.isExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagrams)

	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.isExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagram := range library.Diagrams {
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

		addRenameButton(element, node, stager)

		diagramNode.OnUpdate = stager.onUpdateDiagram(diagram)

		{
			copyButton := &tree.Button{
				Name:            "Diagram Copy",
				Icon:            string(buttons.BUTTON_copy_all),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Copy Diagram",
				OnClick:         onCopyDiagram(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, copyButton)
		}
		{
			showPrefixButton := &tree.Button{
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
				showPrefixButton.Icon = string(buttons.BUTTON_label)
				showPrefixButton.ToolTipText = "Show Prefix"
			} else {
				showPrefixButton.Icon = string(buttons.BUTTON_label_off)
				showPrefixButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
		}

		pbsNode := &tree.Node{
			Name:            "PBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, pbsNode)
		pbsNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded)

		conf := ItemButtonConfiguration[
			Library, *Library, // AT, PAT (Added Element)
			Library, *Library, // ParentAT, PParentAT (Parent Element)
		]{
			parentNode:                         libraryNode,
			sliceForNewAddedItem:               &library.SubLibraries,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &library.isExpanded,
			parentElement:                      library,
		}
		addCreateItemButton(stager, conf)

		for _, product := range library.RootProducts {
			stager.treeProduct(diagram, product, pbsNode)
		}

		diagram.map_Task_TaskCompositionShape = make(map[*Task]*TaskCompositionShape)
		for _, shape := range diagram.TaskComposition_Shapes {
			if shape.Task != nil {
				diagram.map_Task_TaskCompositionShape[shape.Task] = shape
			}
		}

		wbsNode := &tree.Node{
			Name:            "WBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsWBSNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, wbsNode)
		wbsNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsWBSNodeExpanded)

		confWBS := ItemShapeAndLinkButtonConfiguration[
			Task, *Task, // AT, PAT (Added Element)
			Task, *Task, // ParentAT, PParentAT (Parent Element)
			TaskShape, *TaskShape, // CT, PCT (Concrete Shape)
			TaskCompositionShape, *TaskCompositionShape, // ACT, PACT (Association Shape)
		]{
			ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
				Task, *Task, // AT, PAT (Added Element)
				Task, *Task, // ParentAT, PParentAT (Parent Element)
				TaskShape, *TaskShape, // CT, PCT (Concrete Shape)
			]{
				ItemButtonConfiguration: ItemButtonConfiguration[
					Task, *Task, // AT, PAT (Added Element)
					Task, *Task, // ParentAT, PParentAT (Parent Element)
				]{
					parentNode:                         wbsNode,
					sliceForNewAddedItem:               &library.RootTasks,
					isParentNodeExpandedByAddOperation: true,
					parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
					parentNodeExpansionBooleanValue:    &diagram.IsWBSNodeExpanded,
				},
				receivingDiagram:      diagram,
				sliceForNewAddedShape: &diagram.Task_Shapes,
			},
			sliceForNewCompositionShapes: &diagram.TaskComposition_Shapes,
		}
		addCreateItemShapeAndLinkButton(stager, confWBS)

		for _, task := range library.RootTasks {
			stager.treeTask(diagram, task, wbsNode)
		}

		resourcesNode := &tree.Node{
			Name:            "RBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsResourcesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, resourcesNode)
		resourcesNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsResourcesNodeExpanded)

		confRBS := ItemShapeAndLinkButtonConfiguration[
			Resource, *Resource, // AT, PAT (Added Element)
			Resource, *Resource, // ParentAT, PParentAT (Parent Element)
			ResourceShape, *ResourceShape, // CT, PCT (Concrete Shape)
			ResourceTaskShape, *ResourceTaskShape, // ACT, PACT (Association Shape)
		]{
			ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
				Resource, *Resource, // AT, PAT (Added Element)
				Resource, *Resource, // ParentAT, PParentAT (Parent Element)
				ResourceShape, *ResourceShape, // CT, PCT (Concrete Shape)
			]{
				ItemButtonConfiguration: ItemButtonConfiguration[
					Resource, *Resource, // AT, PAT (Added Element)
					Resource, *Resource, // ParentAT, PParentAT (Parent Element)
				]{
					parentNode:                         resourcesNode,
					sliceForNewAddedItem:               &library.RootResources,
					isParentNodeExpandedByAddOperation: true,
					parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
					parentNodeExpansionBooleanValue:    &diagram.IsResourcesNodeExpanded,
				},
				receivingDiagram:      diagram,
				sliceForNewAddedShape: &diagram.Resource_Shapes,
			},
			sliceForNewCompositionShapes: &diagram.ResourceTaskShapes,
		}
		addCreateItemShapeAndLinkButton(stager, confRBS)

		for _, resource := range library.RootResources {
			stager.treeResourceinDiagram(diagram, resource, resourcesNode)
		}

		{
			notesNode := &tree.Node{
				Name:            "Notes",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsNotesNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, notesNode)
			notesNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsNotesNodeExpanded)

			confNotes := ItemShapeAndLinkButtonConfiguration[
				Note, *Note, // AT, PAT (Added Element)
				Note, *Note, // ParentAT, PParentAT (Parent Element)
				NoteShape, *NoteShape, // CT, PCT (Concrete Shape)
				NoteProductShape, *NoteProductShape, // ACT, PACT (Association Shape)
			]{
				ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
					Note, *Note, // AT, PAT (Added Element)
					Note, *Note, // ParentAT, PParentAT (Parent Element)
					NoteShape, *NoteShape, // CT, PCT (Concrete Shape)
				]{
					ItemButtonConfiguration: ItemButtonConfiguration[
						Note, *Note, // AT, PAT (Added Element)
						Note, *Note, // ParentAT, PParentAT (Parent Element)
					]{
						parentNode:                         notesNode,
						sliceForNewAddedItem:               &library.Notes,
						isParentNodeExpandedByAddOperation: true,
						parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
						parentNodeExpansionBooleanValue:    &diagram.IsNotesNodeExpanded,
					},
					receivingDiagram:      diagram,
					sliceForNewAddedShape: &diagram.Note_Shapes,
				},
				sliceForNewCompositionShapes: &diagram.NoteProductShapes,
			}
			addCreateItemShapeAndLinkButton(stager, confNotes)

			for _, note := range library.Notes {
				var dummyMap map[*Note]*NoteProductShape
				var dummySlice *[]*NoteProductShape

				noteNodeConf := TreeNodeShapeAndLinkConfiguration[
					*Note, Note,
					*NoteShape, NoteShape,
					*NoteProductShape, NoteProductShape,
					*Diagram,
				]{
					TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
						*Note, Note,
						*NoteShape, NoteShape,
						*Diagram,
					]{
						TreeNodeConfiguration: TreeNodeConfiguration[
							*Note, Note,
							*Diagram,
						]{
							diagram:                     diagram,
							parentNode:                  notesNode,
							element:                     note,
							parentElement:               (*Note)(nil),
							elementsWhoseNodeIsExpanded: &diagram.NotesWhoseNodeIsExpanded,
						},
						shapes:    &diagram.Note_Shapes,
						shapesMap: diagram.map_Note_NoteShape,
					},
					map_Element_CompositionShape: dummyMap,
					compositionShapes:            dummySlice,
				}
				noteNode := addNodeToTreeWithConf(stager, noteNodeConf)

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
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteProductShape, &diagram.NoteProductShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, product, &diagram.NoteProductShapes)
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
					if _, ok := diagram.map_Task_TaskShape[task]; ok {
						if _, ok := diagram.map_Note_NoteShape[note]; ok {
							noteTaskShape, ok := diagram.map_Note_NoteTaskShape[noteTaskKey{Note: note, Task: task}]
							nodeTask.IsChecked = ok

							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteTaskShape, &diagram.NoteTaskShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, task, &diagram.NoteTaskShapes)
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
					if _, ok := diagram.map_Resource_ResourceShape[resource]; ok {
						if _, ok := diagram.map_Note_NoteShape[note]; ok {
							noteResourceShape, ok := diagram.map_Note_NoteResourceShape[noteResourceKey{Note: note, Resource: resource}]
							nodeResource.IsChecked = ok
							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteResourceShape, &diagram.NoteResourceShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, resource, &diagram.NoteResourceShapes)
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
