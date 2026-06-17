package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.rootLibrary
	_ = rootLibrary

	treeInstance := &tree.Tree{Name: "Library Tree"}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(treeInstance, rootLibrary, &treeInstance.RootNodes)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func (stager *Stager) treeLibrary(treeInstance *tree.Tree, library *Library, parentNodes *[]*tree.Node) {
	var libraryNode = &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    library.IsInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.rootLibrary {
		if !library.GetIsInRenameMode() {
			libraryNode.Buttons = append(libraryNode.Buttons,
				&tree.Button{
					Name: library.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
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
					OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
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

	addAddItemButton(stager, nil, nil, &library.IsExpanded, libraryNode, &library.SubLibraries, nil, &[]*ProductShape{}, &[]*ProductCompositionShape{})
	addAddItemButton(stager, nil, nil, &library.IsExpanded, libraryNode, &library.Diagrams, nil, &[]*ProductShape{}, &[]*ProductCompositionShape{})

	for _, diagram := range library.Diagrams {
		diagramNode := &tree.Node{
			Name:              diagram.Name,
			IsExpanded:        diagram.IsExpanded,
			IsNodeClickable:   true,
			HasCheckboxButton: true,
			IsChecked:         diagram.IsChecked,

			IsInEditMode: diagram.IsInRenameMode,
		}
		libraryNode.Children = append(libraryNode.Children, diagramNode)

		element := diagram
		node := diagramNode

		if !element.GetIsInRenameMode() {
			node.Buttons = append(node.Buttons,
				&tree.Button{
					Name: element.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
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
					OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
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
			copyButton := &tree.Button{
				Name:            "Diagram Copy",
				Icon:            string(buttons.BUTTON_copy_all),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Copy Diagram",
				OnUpdate:        onCopyDiagram(stager, diagram),
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
				OnUpdate:        onShowAllInDiagram(stager, diagram),
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
				OnUpdate:        onLayoutPBS(stager, diagram),
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
				OnUpdate:        onLayoutWBS(stager, diagram),
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
		}
		{
			showAllButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,

				OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
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
		wbsNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsConcernsNodeExpanded)

		addAddItemButton(stager, nil, nil, &diagram.IsConcernsNodeExpanded, wbsNode, &library.RootConcerns, diagram, &diagram.Concern_Shapes, &diagram.ConcernComposition_Shapes)

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
		stakeholderNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsStakeholdersNodeExpanded)

		addAddItemButton(stager, nil, nil, &diagram.IsStakeholdersNodeExpanded, stakeholderNode, &library.RootStakeholders, diagram, &diagram.Stakeholder_Shapes, &diagram.StakeholderConcernShapes)

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
		delivarablesNode.OnUpdate = stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded)

		addAddItemButton(stager, nil, nil, &diagram.IsPBSNodeExpanded, delivarablesNode, &library.RootDeliverables, diagram, &diagram.Product_Shapes, &diagram.ProductComposition_Shapes)

		for _, x := range library.RootDeliverables {
			stager.treeDeliverableRecusriveInDiagram(diagram, x, delivarablesNode)
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

			addAddItemButton(stager, nil, nil, &diagram.IsNotesNodeExpanded, notesNode, &library.Notes, diagram, &diagram.Note_Shapes, &diagram.NoteProductShapes)

			for _, note := range library.Notes {
				var dummyMap map[*Note]*NoteProductShape
				var dummySlice *[]*NoteProductShape

				noteNode := addNodeToTree(
					stager,
					diagram,
					notesNode,
					note,
					(*Note)(nil),
					&diagram.NotesWhoseNodeIsExpanded,
					&diagram.Note_Shapes,
					diagram.map_Note_NoteShape,
					dummyMap,
					dummySlice,
				)

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
								showHideRelationButton.OnUpdate = onRemoveAssociationShape(stager, noteProductShape, &diagram.NoteProductShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								showHideRelationButton.OnUpdate = onAddAssociationShape(stager, note, product, &diagram.NoteProductShapes)
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
								showHideRelationButton.OnUpdate = onRemoveAssociationShape(stager, noteTaskShape, &diagram.NoteTaskShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								showHideRelationButton.OnUpdate = onAddAssociationShape(stager, note, task, &diagram.NoteTaskShapes)
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
								showHideRelationButton.OnUpdate = onRemoveAssociationShape(stager, noteResourceShape, &diagram.NoteResourceShapes)
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to resource \"" + resource.Name + "\""
								showHideRelationButton.OnUpdate = onAddAssociationShape(stager, note, resource, &diagram.NoteResourceShapes)
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

// Helper callbacks

func (stager *Stager) OnUpdateLibrary(library *Library) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			library.IsExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			library.Name = frontNode.Name
			library.IsInRenameMode = false
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

func (stager *Stager) OnUpdateDiagram(diagram *Diagram) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsChecked && !stagedNode.IsChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
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
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			diagram.IsExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			diagram.Name = frontNode.Name
			diagram.IsInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
	}
}

func onCopyDiagram(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, updatedButton *tree.Button) {
	return func(_ *tree.Stage, _ *tree.Button) {
		newDiagram := new(Diagram)
		newDiagram.Name = diagram.Name + " copy"
		newDiagram.IsEditable_ = true

		library := stager.stage.Library_Diagrams_reverseMap[diagram]
		Append(&library.Diagrams, newDiagram)

		for _, s := range diagram.Product_Shapes {
			newShape := s.GongCopy().(*ProductShape)
			newDiagram.Product_Shapes = append(newDiagram.Product_Shapes, newShape)
		}

		for _, s := range diagram.ProductComposition_Shapes {
			newShape := s.GongCopy().(*ProductCompositionShape)
			newDiagram.ProductComposition_Shapes = append(newDiagram.ProductComposition_Shapes, newShape)
		}

		for _, s := range diagram.Concern_Shapes {
			newShape := s.GongCopy().(*ConcernShape)
			newDiagram.Concern_Shapes = append(newDiagram.Concern_Shapes, newShape)
		}

		for _, s := range diagram.ConcernComposition_Shapes {
			newShape := s.GongCopy().(*ConcernCompositionShape)
			newDiagram.ConcernComposition_Shapes = append(newDiagram.ConcernComposition_Shapes, newShape)
		}

		for _, s := range diagram.ConcernInputShapes {
			newShape := s.GongCopy().(*ConcernInputShape)
			newDiagram.ConcernInputShapes = append(newDiagram.ConcernInputShapes, newShape)
		}

		for _, s := range diagram.ConcernOutputShapes {
			newShape := s.GongCopy().(*ConcernOutputShape)
			newDiagram.ConcernOutputShapes = append(newDiagram.ConcernOutputShapes, newShape)
		}

		for _, s := range diagram.Note_Shapes {
			newShape := s.GongCopy().(*NoteShape)
			newDiagram.Note_Shapes = append(newDiagram.Note_Shapes, newShape)
		}

		for _, s := range diagram.NoteProductShapes {
			newShape := s.GongCopy().(*NoteProductShape)
			newDiagram.NoteProductShapes = append(newDiagram.NoteProductShapes, newShape)
		}

		for _, s := range diagram.NoteTaskShapes {
			newShape := s.GongCopy().(*NoteTaskShape)
			newDiagram.NoteTaskShapes = append(newDiagram.NoteTaskShapes, newShape)
		}

		StageBranch(stager.stage, newDiagram)
		stager.stage.Commit()
	}
}
