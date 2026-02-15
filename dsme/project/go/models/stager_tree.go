package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) tree() {
	stager.treeStage.Reset()

	root := stager.root

	treeInstance := &tree.Tree{Name: "Project Tree"}

	allProjectsNode := &tree.Node{
		Name:       "** Tree of Projects **",
		IsExpanded: true,
	}
	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})
	treeInstance.RootNodes = append(treeInstance.RootNodes, allProjectsNode)

	addAddItemButton(stager, nil, nil, nil, allProjectsNode, &root.Projects, nil, &[]*ProductShape{}, &[]*ProductCompositionShape{})

	for _, project := range root.Projects {
		projectNode := &tree.Node{
			Name:            project.Name,
			IsExpanded:      project.IsExpanded,
			IsNodeClickable: true,
		}
		treeInstance.RootNodes = append(treeInstance.RootNodes, projectNode)
		projectNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateProject(project),
		}

		addAddItemButton(stager, nil, nil, &project.IsExpanded, projectNode, &project.Diagrams, nil, &[]*ProductShape{}, &[]*ProductCompositionShape{})

		for _, diagram := range project.Diagrams {
			diagramNode := &tree.Node{
				Name:              diagram.Name,
				IsExpanded:        diagram.IsExpanded,
				IsNodeClickable:   true,
				HasCheckboxButton: true,
				IsChecked:         diagram.IsChecked,
			}
			projectNode.Children = append(projectNode.Children, diagramNode)
			diagramNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateDiagram(diagram),
			}

			{
				copyButton := &tree.Button{
					Name:            "Diagram Copy",
					Icon:            string(buttons.BUTTON_copy_all),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,
					ToolTipText:     "Copy Diagram",
					Impl: &tree.FunctionalButtonProxy{
						OnUpdated: OnCopyDiagram(stager, diagram),
					},
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
					Impl: &tree.FunctionalButtonProxy{
						OnUpdated: onShowAllInDiagram(stager, diagram),
					},
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
					Impl: &tree.FunctionalButtonProxy{
						OnUpdated: onLayoutPBS(stager, diagram),
					},
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
					Impl: &tree.FunctionalButtonProxy{
						OnUpdated: onLayoutWBS(stager, diagram),
					},
				}
				diagramNode.Buttons = append(diagramNode.Buttons, showAllButton)
			}
			{
				showAllButton := &tree.Button{
					Name:            "Diagram Prefix",
					Icon:            string(buttons.BUTTON_show_chart),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,

					Impl: &tree.FunctionalButtonProxy{
						OnUpdated: func(stage *tree.Stage, button, updatedButton *tree.Button) {
							diagram.ShowPrefix = !diagram.ShowPrefix
							stager.stage.Commit()
						},
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

			pbsNode := &tree.Node{
				Name:            "PBS",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsPBSNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, pbsNode)
			pbsNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded),
			}

			addAddItemButton(stager, nil, nil, &diagram.IsPBSNodeExpanded, pbsNode, &project.RootProducts, diagram, &diagram.Product_Shapes, &diagram.ProductComposition_Shapes)

			for _, product := range project.RootProducts {
				stager.treePBSRecusriveInDiagram(diagram, product, pbsNode)
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
			wbsNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsWBSNodeExpanded),
			}

			addAddItemButton(stager, nil, nil, &diagram.IsWBSNodeExpanded, wbsNode, &project.RootTasks, diagram, &diagram.Task_Shapes, &diagram.TaskComposition_Shapes)

			for _, task := range project.RootTasks {
				stager.treeWBSinDiagram(diagram, task, wbsNode)
			}

			notesNode := &tree.Node{
				Name:            "Notes",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsNotesNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, notesNode)
			notesNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsNotesNodeExpanded),
			}

			addAddItemButton(stager, nil, nil, &diagram.IsNotesNodeExpanded, notesNode, &project.Notes, diagram, &diagram.Note_Shapes, &diagram.NoteProductShapes)

			for _, note := range project.Notes {
				noteNode := addNodeToTree(
					stager,
					diagram,
					notesNode,
					note,
					&diagram.NotesWhoseNodeIsExpanded,
					&diagram.Note_Shapes,
					&diagram.map_Note_NoteShape,
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
								showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
									OnUpdated: onRemoveAssociationShape(stager, noteProductShape, &diagram.NoteProductShapes),
								}
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to product \"" + product.Name + "\""
								showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
									OnUpdated: onAddAssociationShape(stager, note, product, &diagram.NoteProductShapes),
								}
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
								showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
									OnUpdated: onRemoveAssociationShape(stager, noteTaskShape, &diagram.NoteTaskShapes),
								}
							} else {
								showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
								showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
									"\" to task \"" + task.Name + "\""
								showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
									OnUpdated: onAddAssociationShape(stager, note, task, &diagram.NoteTaskShapes),
								}
							}
						}
					}
				}
			}

			resourcesNode := &tree.Node{
				Name:            "Resources",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsResourcesNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, resourcesNode)
			resourcesNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsResourcesNodeExpanded),
			}

			addAddItemButton(stager, nil, nil, &diagram.IsResourcesNodeExpanded, resourcesNode, &project.RootResources, diagram, &diagram.Resource_Shapes, &diagram.ResourceTaskShapes)

			for _, resource := range project.RootResources {
				stager.treeRBSinDiagram(diagram, resource, resourcesNode)
			}

		}

	}

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

// Helper callbacks

func (stager *Stager) OnUpdateProject(project *Project) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			project.IsExpanded = frontNode.IsExpanded
		} else {
			stager.probeForm.FillUpFormFromGongstruct(project, GetPointerToGongstructName[*Project]())
		}
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

func OnCopyDiagram(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		newDiagram := new(Diagram)
		newDiagram.Name = diagram.Name + " copy"
		newDiagram.IsEditable_ = true

		project := stager.stage.Project_Diagrams_reverseMap[diagram]
		Append(&project.Diagrams, newDiagram)

		for _, s := range diagram.Product_Shapes {
			newShape := s.GongCopy().(*ProductShape)
			newDiagram.Product_Shapes = append(newDiagram.Product_Shapes, newShape)
		}

		for _, s := range diagram.ProductComposition_Shapes {
			newShape := s.GongCopy().(*ProductCompositionShape)
			newDiagram.ProductComposition_Shapes = append(newDiagram.ProductComposition_Shapes, newShape)
		}

		for _, s := range diagram.Task_Shapes {
			newShape := s.GongCopy().(*TaskShape)
			newDiagram.Task_Shapes = append(newDiagram.Task_Shapes, newShape)
		}

		for _, s := range diagram.TaskComposition_Shapes {
			newShape := s.GongCopy().(*TaskCompositionShape)
			newDiagram.TaskComposition_Shapes = append(newDiagram.TaskComposition_Shapes, newShape)
		}

		for _, s := range diagram.TaskInputShapes {
			newShape := s.GongCopy().(*TaskInputShape)
			newDiagram.TaskInputShapes = append(newDiagram.TaskInputShapes, newShape)
		}

		for _, s := range diagram.TaskOutputShapes {
			newShape := s.GongCopy().(*TaskOutputShape)
			newDiagram.TaskOutputShapes = append(newDiagram.TaskOutputShapes, newShape)
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
