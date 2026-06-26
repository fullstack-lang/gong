package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagram(library *Library, diagram *Diagram, libraryNode *tree.Node) {
	diagramNode := &tree.Node{
		Name:              diagram.Name,
		IsExpanded:        diagram.IsExpanded,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagram.IsChecked,

		IsInEditMode: diagram.isInRenameMode,
	}
	libraryNode.Children = append(libraryNode.Children, diagramNode)

	element := diagram
	node := diagramNode

	addRenameButton(element, node, stager)
	diagramNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
		} else {
			diagram.IsChecked = false
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
		}
		stager.stage.Commit()
	}
	diagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	diagramNode.OnNameChange = stager.onNameChange(diagram)
	diagramNode.OnClick = onNodeClicked(stager, diagram)
	{
		copyButton := &tree.Button{
			Name:            "Copy Diagram",
			Icon:            string(buttons.BUTTON_copy_all),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Copy Diagram",
			OnClick:         onCopyDiagram(stager, diagram),
		}
		if diagramNode.Menu == nil {
			diagramNode.Menu = &tree.Menu{Name: "Menu"}
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, copyButton)
	}
	{
		showPrefixButton := &tree.Button{
			Name:            "Show Prefix",
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
			showPrefixButton.Name = "Show Prefix"
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.Name = "Hide Prefix"
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		if diagramNode.Menu == nil {
			diagramNode.Menu = &tree.Menu{Name: "Menu"}
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, showPrefixButton)
	}
	{
		timeDiagramButton := &tree.Button{
			Name:            "Time Diagram",
			Icon:            string(buttons.BUTTON_timer),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagram.IsTimeDiagram = !diagram.IsTimeDiagram
				stager.stage.Commit()
			},
		}
		if !diagram.IsTimeDiagram {
			timeDiagramButton.Icon = string(buttons.BUTTON_timer_off)
			timeDiagramButton.Name = "Set as Time Diagram"
			timeDiagramButton.ToolTipText = "Set as Time Diagram"
		} else {
			timeDiagramButton.Icon = string(buttons.BUTTON_timer)
			timeDiagramButton.Name = "Unset Time Diagram"
			timeDiagramButton.ToolTipText = "Unset Time Diagram"
		}
		if diagramNode.Menu == nil {
			diagramNode.Menu = &tree.Menu{Name: "Menu"}
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, timeDiagramButton)
	}
	{
		layoutButton := &tree.Button{
			Name:            "Layout",
			Icon:            string(buttons.BUTTON_format_align_justify),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Layout Diagram",
			OnClick: func() {
				layoutDiagram(diagram, stager)
				stager.stage.Commit()
			},
		}
		if diagramNode.Menu == nil {
			diagramNode.Menu = &tree.Menu{Name: "Menu"}
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, layoutButton)
	}

	pbsNode := &tree.Node{
		Name:            "PBS",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsPBSNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, pbsNode)
	pbsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsPBSNodeExpanded)
	pbsNode.OnClick = onNodeClicked(stager, diagram)
	confPBS := ItemShapeAndLinkButtonConfiguration[
		Product, *Product, // AT, PAT (Added Element)
		Product, *Product, // ParentAT, PParentAT (Parent Element)
		ProductShape, *ProductShape, // CT, PCT (Concrete Shape)
		ProductCompositionShape, *ProductCompositionShape, // ACT, PACT (Association Shape)
	]{
		ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
			Product, *Product, // AT, PAT (Added Element)
			Product, *Product, // ParentAT, PParentAT (Parent Element)
			ProductShape, *ProductShape, // CT, PCT (Concrete Shape)
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Product, *Product, // AT, PAT (Added Element)
				Product, *Product, // ParentAT, PParentAT (Parent Element)
			]{
				parentNode:                         pbsNode,
				sliceForNewAddedItem:               &library.RootProducts,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsPBSNodeExpanded,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Product_Shapes,
		},
		sliceForNewCompositionShapes: &diagram.ProductComposition_Shapes,
	}
	addCreateItemShapeAndLinkButton(stager, confPBS)

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
	wbsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsWBSNodeExpanded)
	wbsNode.OnClick = onNodeClicked(stager, diagram)

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

	taskGroupsNode := &tree.Node{
		Name:            "TaskGroups",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsTaskGroupsNodeExpanded,
		IsNodeClickable: true,
	}
	wbsNode.Children = append(wbsNode.Children, taskGroupsNode)
	taskGroupsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsTaskGroupsNodeExpanded)
	taskGroupsNode.OnClick = onNodeClicked(stager, diagram)

	confTaskGroups := ItemAndShapeButtonConfiguration[
		TaskGroup, *TaskGroup, // AT, PAT (Added Element)
		TaskGroup, *TaskGroup, // ParentAT, PParentAT (Parent Element)
		TaskGroupShape, *TaskGroupShape, // CT, PCT (Concrete Shape)
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			TaskGroup, *TaskGroup, // AT, PAT (Added Element)
			TaskGroup, *TaskGroup, // ParentAT, PParentAT (Parent Element)
		]{
			parentNode:                         taskGroupsNode,
			sliceForNewAddedItem:               &library.RootTaskGroups,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsTaskGroupsNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.TaskGroupShapes,
	}
	addCreateItemAndShapeButton(stager, confTaskGroups)

	for _, taskGroup := range library.RootTaskGroups {
		taskGroupNodeConf := TreeNodeAndShapeConfigurationWithoutLink[
			*TaskGroup, TaskGroup, // AT, AT_
			*Library, Library, // ParentAT, ParentAT_
			*TaskGroupShape, TaskGroupShape, // CT, CT_
			*Diagram, // DiagramType
		]{
			diagram:                     diagram,
			parentNode:                  taskGroupsNode,
			element:                     taskGroup,
			parentElement:               library,
			elementsWhoseNodeIsExpanded: &diagram.TaskGroupsWhoseNodeIsExpanded,
			shapes:                      &diagram.TaskGroupShapes,
			shapesMap:                   diagram.map_TaskGroup_TaskGroupShape,
		}
		taskGroupNode := addNodeToTreeWithoutLink(stager, taskGroupNodeConf)

		for _, task := range taskGroup.Tasks {
			stager.treeTask(diagram, task, taskGroupNode)
		}
	}

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
	resourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsResourcesNodeExpanded)
	resourcesNode.OnClick = onNodeClicked(stager, diagram)

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
		notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsNotesNodeExpanded)
		notesNode.OnClick = onNodeClicked(stager, diagram)

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
				*Note, Note, // AT, AT_
				*NoteShape, NoteShape, // CT, CT_
				*NoteProductShape, NoteProductShape, // ACT, ACT_
				*Diagram, // DiagramType
			]{
				TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
					*Note, Note, // AT, AT_
					*NoteShape, NoteShape, // CT, CT_
					*Diagram, // DiagramType
				]{
					TreeNodeConfiguration: TreeNodeConfiguration[
						*Note, Note, // AT, AT_
						*Diagram, // DiagramType
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
			noteNode := addNodeToTree(stager, noteNodeConf)

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
