package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) updateTreeDiagramStage() {

	treeStage := stager.treeDiagramStage

	treeStage.Reset()

	treeInstance := new(tree.Tree)
	treeInstance.Name = string(DiagramTreeName)

	transitionsSet := *GetGongstructInstancesSet[Transition](stager.stage)
	transitionSlice := SortGongstructSetByName(transitionsSet)

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	// architecture node
	architectureNode := &tree.Node{
		Name:            "State Machines",
		FontStyle:       tree.ITALIC,
		IsNodeClickable: true,
		Impl: &diagramArchitectureNodeProxy{
			stager:       stager,
			architecture: stager.architecture,
		},
	}
	treeInstance.RootNodes = append(treeInstance.RootNodes, architectureNode)
	{
		addButton := &tree.Button{
			Name:            "State Machine" + " " + string(buttons.BUTTON_add_box),
			Icon:            string(buttons.BUTTON_add_box),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Add a State Machine",
		}
		architectureNode.Buttons = append(architectureNode.Buttons, addButton)
		addButton.Impl = &ArchitectureAddDiagramButtonProxy{
			stager:      stager,
			architecure: stager.architecture,
		}
	}

	for _, stateMachine := range stager.architecture.StateMachines {
		stateMachineNode := &tree.Node{
			Name:            stateMachine.Name,
			IsExpanded:      stateMachine.IsNodeExpanded,
			IsNodeClickable: true,
			Impl: &diagramStateMachineNodeProxy{
				stager:         stager,
				stateMachine:   stateMachine,
				isNodeExpanded: &stateMachine.IsNodeExpanded,
			},
		}

		treeInstance.RootNodes = append(treeInstance.RootNodes, stateMachineNode)

		{
			addButton := &tree.Button{
				Name:            "Diagram Add" + " " + string(buttons.BUTTON_add_box),
				Icon:            string(buttons.BUTTON_add_box),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Add a Diagram to the state machine",
			}
			stateMachineNode.Buttons = append(stateMachineNode.Buttons, addButton)
			addButton.Impl = &StateMachineAddDiagramButtonProxy{
				stager:       stager,
				stateMachine: stateMachine,
			}
		}

		for _, diagram := range stateMachine.Diagrams {
			diagramNode := new(tree.Node)
			stateMachineNode.Children = append(stateMachineNode.Children, diagramNode)
			diagramNode.Name = diagram.Name
			diagramNode.IsChecked = diagram.IsChecked
			diagramNode.IsExpanded = diagram.IsExpanded

			diagramNode.IsInEditMode = diagram.IsInRenameMode

			diagramNode.IsNodeClickable = true
			diagramProxy := new(Diagram_Tree_DiagramProxy)
			diagramProxy.stager = stager
			diagramProxy.diagram = diagram

			diagramNode.HasCheckboxButton = true
			diagramNode.Impl = diagramProxy

			if !diagram.IsInRenameMode {
				diagramNode.Buttons = append(diagramNode.Buttons,
					&tree.Button{
						Name: diagram.GetName() + " " + string(buttons.BUTTON_edit_note),
						Icon: string(buttons.BUTTON_edit_note),
						Impl: &tree.FunctionalButtonProxy{
							OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
								diagram.IsInRenameMode = true
								stager.stage.Commit()
							},
						},
						HasToolTip:      true,
						ToolTipText:     "Rename the diagram",
						ToolTipPosition: tree.Above,
					})
			} else {
				diagramNode.Buttons = append(diagramNode.Buttons,
					&tree.Button{
						Name: diagram.GetName() + " " + string(buttons.BUTTON_edit_off),
						Icon: string(buttons.BUTTON_edit_off),
						Impl: &tree.FunctionalButtonProxy{
							OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
								diagram.IsInRenameMode = false
								stager.stage.Commit()
							},
						},
						HasToolTip:      true,
						ToolTipText:     "Cancel renaming",
						ToolTipPosition: tree.Above,
					})
			}
			{
				copyButton := &tree.Button{
					Name:            "Diagram Copy" + " " + string(buttons.BUTTON_copy_all),
					Icon:            string(buttons.BUTTON_copy_all),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,
					ToolTipText:     "Copy Diagram",
					Impl: &DiagramCopyButtonProxy{
						stager:  stager,
						diagram: diagram,
					},
				}
				diagramNode.Buttons = append(diagramNode.Buttons, copyButton)
			}
			if diagram.IsChecked {
				addButton := &tree.Button{
					Name:            "Diagram Add" + " " + string(buttons.BUTTON_add_box),
					Icon:            string(buttons.BUTTON_add_box),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,
					ToolTipText:     "Add a State to the State Machine and add it to the diagram",
					Impl: &DiagramAddStateButtonProxy{
						stager:       stager,
						stateMachine: stateMachine,
						diagram:      diagram,
					},
				}
				diagramNode.Buttons = append(diagramNode.Buttons, addButton)
			}

			// for displaying wether the State node is checked
			map_State__StateShape := make(map[*State]*StateShape)
			for _, stateShape := range diagram.State_Shapes {
				map_State__StateShape[stateShape.State] = stateShape
			}

			// for displaying wether the State_Transition node is checked
			map_Transition__TransitionShape := make(map[*Transition]*Transition_Shape)
			for _, transitionShape := range diagram.Transition_Shapes {
				map_Transition__TransitionShape[transitionShape.Transition] = transitionShape
			}

			for _, state := range stateMachine.States {
				diagramStateNode := new(tree.Node)
				diagramStateNode.Name = state.Name
				diagramStateNode.HasCheckboxButton = true
				diagramStateNode.IsNodeClickable = true

				var stateShape *StateShape
				var ok bool
				if stateShape, ok = map_State__StateShape[state]; ok {
					diagramStateNode.IsChecked = true
					diagramStateNode.IsExpanded = stateShape.IsExpanded

					diagramStateNode.OnUpdate = func(_ *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {
						if frontNode.IsExpanded && !stagedNode.IsExpanded {
							stagedNode.IsExpanded = frontNode.IsExpanded
							stateShape.IsExpanded = true
							stager.stage.Commit()
						}
						if !frontNode.IsExpanded && stagedNode.IsExpanded {
							stagedNode.IsExpanded = frontNode.IsExpanded
							stateShape.IsExpanded = false
							stager.stage.Commit()
						}
					}

					diagramStateNode.Buttons = append(diagramStateNode.Buttons,
						&tree.Button{
							Name:            "Show State Shape",
							Icon:            string(buttons.BUTTON_visibility),
							HasToolTip:      true,
							ToolTipPosition: tree.Above,
							ToolTipText:     "Show State Shape",
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								stateShape.SetIsHidden(!stateShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					)
					if !stateShape.GetIsHidden() {
						diagramStateNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
						diagramStateNode.Buttons[0].ToolTipText = "Hide State Shape"
					}

					// range over transitions that have the state as a source or target
					// if the target is present, enable the check button
					for _, transition_ := range transitionSlice {
						if transition_.Start == state && transition_.End != nil {
							transitionNode := new(tree.Node)
							transitionNode.Name = transition_.Name + " --> " + transition_.End.Name
							transitionNode.HasCheckboxButton = true

							proxy := new(DiagramTree_Transition_Proxy)
							proxy.diagram = diagram
							proxy.stager = stager
							proxy.transition = transition_
							transitionNode.Impl = proxy

							// set the transition of the proxy (can be nil, )
							var transitionShape *Transition_Shape
							if transitionShape_, ok := map_Transition__TransitionShape[transition_]; ok {
								transitionShape = transitionShape_
								transitionNode.IsChecked = true
							}

							if transitionShape == nil {
								continue
							}

							proxy.transitionShape = transitionShape

							// disable if end is not present
							if _, ok := map_State__StateShape[transition_.End]; !ok {
								transitionNode.IsCheckboxDisabled = true
							}

							transitionNode.Buttons = append(transitionNode.Buttons,
								&tree.Button{
									Name:            "Show Transition Shape",
									Icon:            string(buttons.BUTTON_visibility),
									HasToolTip:      true,
									ToolTipPosition: tree.Above,
									ToolTipText:     "Show Transition Shape",
									OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
										transitionShape.SetIsHidden(!transitionShape.GetIsHidden())
										stager.stage.Commit()
									},
								},
							)
							if !transitionShape.GetIsHidden() {
								transitionNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
								transitionNode.Buttons[0].ToolTipText = "Hide Transition Shape"
							}

							diagramStateNode.Children = append(diagramStateNode.Children, transitionNode)
						}
					}
				}

				diagramStateNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
					if frontNode.IsChecked && !stagedNode.IsChecked {

						stagedNode.IsChecked = frontNode.IsChecked

						// add the State_Shape
						if stateShape != nil {
							log.Fatalln("adding a shape to an already state shape")
						}

						newStateShapeToDiagram(state, diagram).Stage(stager.stage)

						stager.stage.Commit()
						return
					}
					if !frontNode.IsChecked && stagedNode.IsChecked {

						stagedNode.IsChecked = frontNode.IsChecked

						// one need to remove the State_Shape
						if stateShape == nil {
							log.Fatalln("remove a non existing shape to state")
						}

						stateShape.Unstage(stager.stage)

						idx := slices.Index(diagram.State_Shapes, stateShape)
						diagram.State_Shapes = slices.Delete(diagram.State_Shapes, idx, idx+1)

						stager.stage.Commit()
						return
					}

					stager.probeForm.FillUpFormFromGongstruct(state, "State")
				}

				diagramNode.Children = append(diagramNode.Children, diagramStateNode)
			}
		}
	}

	tree.StageBranch(treeStage, treeInstance)

	treeStage.Commit()
}
