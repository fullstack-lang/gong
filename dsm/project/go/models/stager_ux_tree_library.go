package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.getRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpanded)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	confDiagrams := ItemButtonConfiguration[
		DiagramHierarchy, *DiagramHierarchy, // AT, PAT (Added Element)
		Library, *Library, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagrams)

	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsEditable_ = true
		newDiagram.IsExpanded = true
		for diagram_ := range *GetGongstructInstancesSet[DiagramHierarchy](stager.stage) {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
	}

	for _, diagramHierarchy := range library.Diagrams {
		diagramNode := &tree.Node{
			Name:              diagramHierarchy.Name,
			IsExpanded:        diagramHierarchy.IsExpanded,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagramHierarchy.IsChecked,

			IsInEditMode: diagramHierarchy.isInRenameMode,
		}
		libraryNode.Children = append(libraryNode.Children, diagramNode)

		element := diagramHierarchy
		node := diagramNode

		addRenameButton(element, node, stager)
		diagramNode.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked {
				for diagram_ := range *GetGongstructInstancesSet[DiagramHierarchy](stager.stage) {
					diagram_.IsChecked = false
				}
				diagramHierarchy.IsChecked = true
			} else {
				diagramHierarchy.IsChecked = false
				for diagram_ := range *GetGongstructInstancesSet[DiagramHierarchy](stager.stage) {
					diagram_.IsChecked = false
				}
			}
			stager.stage.Commit()
		}
		diagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsExpanded)
		diagramNode.OnNameChange = stager.onNameChange(diagramHierarchy)
		diagramNode.OnClick = onNodeClicked(stager, diagramHierarchy)
		{
			copyButton := &tree.Button{
				Name:            "Copy Diagram",
				Icon:            string(buttons.BUTTON_copy_all),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Copy Diagram",
				OnClick:         onCopyDiagram(stager, diagramHierarchy),
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
					diagramHierarchy.IsShowPrefix = !diagramHierarchy.IsShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagramHierarchy.IsShowPrefix {
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
					diagramHierarchy.IsTimeDiagram = !diagramHierarchy.IsTimeDiagram
					stager.stage.Commit()
				},
			}
			if !diagramHierarchy.IsTimeDiagram {
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

		pbsNode := &tree.Node{
			Name:            "PBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramHierarchy.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, pbsNode)
		pbsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsPBSNodeExpanded)
		pbsNode.OnClick = onNodeClicked(stager, diagramHierarchy)
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
					parentNodeExpansionBooleanValue:    &diagramHierarchy.IsWBSNodeExpanded,
				},
				receivingDiagram:      diagramHierarchy,
				sliceForNewAddedShape: &diagramHierarchy.Product_Shapes,
			},
			sliceForNewCompositionShapes: &diagramHierarchy.ProductComposition_Shapes,
		}
		addCreateItemShapeAndLinkButton(stager, confPBS)

		for _, product := range library.RootProducts {
			stager.treeProduct(diagramHierarchy, product, pbsNode)
		}

		diagramHierarchy.map_Task_TaskCompositionShape = make(map[*Task]*TaskCompositionShape)
		for _, shape := range diagramHierarchy.TaskComposition_Shapes {
			if shape.Task != nil {
				diagramHierarchy.map_Task_TaskCompositionShape[shape.Task] = shape
			}
		}

		wbsNode := &tree.Node{
			Name:            "WBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramHierarchy.IsWBSNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, wbsNode)
		wbsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsWBSNodeExpanded)
		wbsNode.OnClick = onNodeClicked(stager, diagramHierarchy)

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
					parentNodeExpansionBooleanValue:    &diagramHierarchy.IsWBSNodeExpanded,
				},
				receivingDiagram:      diagramHierarchy,
				sliceForNewAddedShape: &diagramHierarchy.Task_Shapes,
			},
			sliceForNewCompositionShapes: &diagramHierarchy.TaskComposition_Shapes,
		}
		addCreateItemShapeAndLinkButton(stager, confWBS)

		taskGroupsNode := &tree.Node{
			Name:            "TaskGroups",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramHierarchy.IsTaskGroupsNodeExpanded,
			IsNodeClickable: true,
		}
		wbsNode.Children = append(wbsNode.Children, taskGroupsNode)
		taskGroupsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsTaskGroupsNodeExpanded)
		taskGroupsNode.OnClick = onNodeClicked(stager, diagramHierarchy)

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
				parentNodeExpansionBooleanValue:    &diagramHierarchy.IsTaskGroupsNodeExpanded,
			},
			receivingDiagram:      diagramHierarchy,
			sliceForNewAddedShape: &diagramHierarchy.TaskGroupShapes,
		}
		addCreateItemAndShapeButton(stager, confTaskGroups)

		for _, taskGroup := range library.RootTaskGroups {
			taskGroupNodeConf := TreeNodeAndShapeConfigurationWithoutLink[
				*TaskGroup, TaskGroup, // AT, AT_
				*Library, Library, // ParentAT, ParentAT_
				*TaskGroupShape, TaskGroupShape, // CT, CT_
				*DiagramHierarchy, // DiagramType
			]{
				diagram:            diagramHierarchy,
				parentNode:                  taskGroupsNode,
				element:                     taskGroup,
				parentElement:               library,
				elementsWhoseNodeIsExpanded: &diagramHierarchy.TaskGroupsWhoseNodeIsExpanded,
				shapes:                      &diagramHierarchy.TaskGroupShapes,
				shapesMap:                   diagramHierarchy.map_TaskGroup_TaskGroupShape,
			}
			taskGroupNode := addNodeToTreeWithoutLink(stager, taskGroupNodeConf)

			for _, task := range taskGroup.Tasks {
				stager.treeTask(diagramHierarchy, task, taskGroupNode)
			}
		}

		milestonesNode := &tree.Node{
			Name:            "Milestones",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramHierarchy.IsMilestonesNodeExpanded,
			IsNodeClickable: true,
		}
		wbsNode.Children = append(wbsNode.Children, milestonesNode)
		milestonesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsMilestonesNodeExpanded)
		milestonesNode.OnClick = onNodeClicked(stager, diagramHierarchy)

		confMilestones := ItemAndShapeButtonConfiguration[
			Milestone, *Milestone, // AT, PAT (Added Element)
			Milestone, *Milestone, // ParentAT, PParentAT (Parent Element)
			MilestoneShape, *MilestoneShape, // CT, PCT (Concrete Shape)
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Milestone, *Milestone, // AT, PAT (Added Element)
				Milestone, *Milestone, // ParentAT, PParentAT (Parent Element)
			]{
				parentNode:                         milestonesNode,
				sliceForNewAddedItem:               &library.RootMilestones,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagramHierarchy.IsMilestonesNodeExpanded,
			},
			receivingDiagram:      diagramHierarchy,
			sliceForNewAddedShape: &diagramHierarchy.MilestoneShapes,
		}
		addCreateItemAndShapeButton(stager, confMilestones)

		for _, milestone := range library.RootMilestones {
			milestoneNodeConf := TreeNodeAndShapeConfigurationWithoutLink[
				*Milestone, Milestone, // AT, AT_
				*Library, Library, // ParentAT, ParentAT_
				*MilestoneShape, MilestoneShape, // CT, CT_
				*DiagramHierarchy, // DiagramType
			]{
				diagram:            diagramHierarchy,
				parentNode:                  milestonesNode,
				element:                     milestone,
				parentElement:               library,
				elementsWhoseNodeIsExpanded: &diagramHierarchy.MilestonesWhoseNodeIsExpanded,
				shapes:                      &diagramHierarchy.MilestoneShapes,
				shapesMap:                   diagramHierarchy.map_Milestone_MilestoneShape,
			}
			milestoneNode := addNodeToTreeWithoutLink(stager, milestoneNodeConf)
			_ = milestoneNode
		}

		for _, task := range library.RootTasks {
			stager.treeTask(diagramHierarchy, task, wbsNode)
		}

		resourcesNode := &tree.Node{
			Name:            "RBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagramHierarchy.IsResourcesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, resourcesNode)
		resourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsResourcesNodeExpanded)
		resourcesNode.OnClick = onNodeClicked(stager, diagramHierarchy)

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
					parentNodeExpansionBooleanValue:    &diagramHierarchy.IsResourcesNodeExpanded,
				},
				receivingDiagram:      diagramHierarchy,
				sliceForNewAddedShape: &diagramHierarchy.Resource_Shapes,
			},
			sliceForNewCompositionShapes: &diagramHierarchy.ResourceTaskShapes,
		}
		addCreateItemShapeAndLinkButton(stager, confRBS)

		for _, resource := range library.RootResources {
			stager.treeResourceinDiagram(diagramHierarchy, resource, resourcesNode)
		}

		{
			notesNode := &tree.Node{
				Name:            "Notes",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagramHierarchy.IsNotesNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, notesNode)
			notesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagramHierarchy.IsNotesNodeExpanded)
			notesNode.OnClick = onNodeClicked(stager, diagramHierarchy)

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
						parentNodeExpansionBooleanValue:    &diagramHierarchy.IsNotesNodeExpanded,
					},
					receivingDiagram:      diagramHierarchy,
					sliceForNewAddedShape: &diagramHierarchy.Note_Shapes,
				},
				sliceForNewCompositionShapes: &diagramHierarchy.NoteProductShapes,
			}
			addCreateItemShapeAndLinkButton(stager, confNotes)

			for _, note := range library.Notes {
				var dummyMap map[*Note]*NoteProductShape
				var dummySlice *[]*NoteProductShape

				noteNodeConf := TreeNodeShapeAndLinkConfiguration[
					*Note, Note, // AT, AT_
					*NoteShape, NoteShape, // CT, CT_
					*NoteProductShape, NoteProductShape, // ACT, ACT_
					*DiagramHierarchy, // DiagramType
				]{
					TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
						*Note, Note, // AT, AT_
						*NoteShape, NoteShape, // CT, CT_
						*DiagramHierarchy, // DiagramType
					]{
						TreeNodeConfiguration: TreeNodeConfiguration[
							*Note, Note, // AT, AT_
							*DiagramHierarchy, // DiagramType
						]{
							diagram:            diagramHierarchy,
							parentNode:                  notesNode,
							element:                     note,
							parentElement:               (*Note)(nil),
							elementsWhoseNodeIsExpanded: &diagramHierarchy.NotesWhoseNodeIsExpanded,
						},
						shapes:    &diagramHierarchy.Note_Shapes,
						shapesMap: diagramHierarchy.map_Note_NoteShape,
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

					if _, ok := diagramHierarchy.map_Product_ProductShape[product]; ok {
						if _, ok := diagramHierarchy.map_Note_NoteShape[note]; ok {

							noteProductShape, ok := diagramHierarchy.map_Note_NoteProductShape[noteProductKey{Note: note, Product: product}]
							nodeProduct.IsChecked = ok

							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteProductShape, &diagramHierarchy.NoteProductShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, product, &diagramHierarchy.NoteProductShapes)
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
					if _, ok := diagramHierarchy.map_Task_TaskShape[task]; ok {
						if _, ok := diagramHierarchy.map_Note_NoteShape[note]; ok {
							noteTaskShape, ok := diagramHierarchy.map_Note_NoteTaskShape[noteTaskKey{Note: note, Task: task}]
							nodeTask.IsChecked = ok

							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteTaskShape, &diagramHierarchy.NoteTaskShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, task, &diagramHierarchy.NoteTaskShapes)
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
					if _, ok := diagramHierarchy.map_Resource_ResourceShape[resource]; ok {
						if _, ok := diagramHierarchy.map_Note_NoteShape[note]; ok {
							noteResourceShape, ok := diagramHierarchy.map_Note_NoteResourceShape[noteResourceKey{Note: note, Resource: resource}]
							nodeResource.IsChecked = ok
							if ok {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								// what to do when the product node is clicked
								showHideRelationButton.OnClick = onRemoveAssociationShape(stager, noteResourceShape, &diagramHierarchy.NoteResourceShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								showHideRelationButton.OnClick = onAddAssociationShape(stager, note, resource, &diagramHierarchy.NoteResourceShapes)
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
