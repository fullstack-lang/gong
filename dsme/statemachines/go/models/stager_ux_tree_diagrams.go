package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagrams() {

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
		OnUpdate: func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(stager.architecture, "Architecture")

			stager.stage.Commit()
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
			OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
				s := stager.stage
				newDiagram := (&StateMachine{
					Name: "New StateMachine",
				}).Stage(s)

				stager.architecture.StateMachines = append(stager.architecture.StateMachines, newDiagram)
				stager.stage.Commit()
			},
		}
		architectureNode.Buttons = append(architectureNode.Buttons, addButton)
	}

	for _, stateMachine := range stager.architecture.StateMachines {
		stateMachineNode := &tree.Node{
			Name:            stateMachine.Name,
			IsExpanded:      stateMachine.IsNodeExpanded,
			IsNodeClickable: true,
			OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					stateMachine.IsNodeExpanded = !stateMachine.IsNodeExpanded
					stager.stage.Commit()
					return
				}
				stager.probeForm.FillUpFormFromGongstruct(stateMachine, stateMachine.GongGetGongstructName())

				stager.stage.Commit()
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
				OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
					s := stager.stage
					newDiagram := (&Diagram{
						Name:        "New Diagram",
						IsEditable_: true,
					}).Stage(s)

					stateMachine.Diagrams = append(stateMachine.Diagrams, newDiagram)
					stager.stage.Commit()
				},
			}
			stateMachineNode.Buttons = append(stateMachineNode.Buttons, addButton)
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
					OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
						s := stager.stage
						newState := new(State).Stage(s)

						newState.Name = "New State"
						stateMachine.States = append(stateMachine.States, newState)

						newStateShapeToDiagram(newState, diagram).Stage(stager.stage)

						stager.stage.Commit()
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
				diagramStateNode.IsInEditMode = state.IsInRenameMode
				diagramStateNode.IsExpanded = slices.Index(diagram.StatesWhoseNodeIsExpanded, state) != -1

				if !state.IsInRenameMode {
					diagramStateNode.Buttons = append(diagramStateNode.Buttons,
						&tree.Button{
							Name: state.GetName() + " " + string(buttons.BUTTON_edit_note),
							Icon: string(buttons.BUTTON_edit_note),
							Impl: &tree.FunctionalButtonProxy{
								OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
									state.IsInRenameMode = true
									stager.stage.Commit()
								},
							},
							HasToolTip:      true,
							ToolTipText:     "Rename the state",
							ToolTipPosition: tree.Above,
						})
				} else {
					diagramStateNode.Buttons = append(diagramStateNode.Buttons,
						&tree.Button{
							Name: state.GetName() + " " + string(buttons.BUTTON_edit_off),
							Icon: string(buttons.BUTTON_edit_off),
							Impl: &tree.FunctionalButtonProxy{
								OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
									state.IsInRenameMode = false
									stager.stage.Commit()
								},
							},
							HasToolTip:      true,
							ToolTipText:     "Cancel renaming",
							ToolTipPosition: tree.Above,
						})
				}

				var stateShape *StateShape
				var ok bool
				if stateShape, ok = map_State__StateShape[state]; ok {
					diagramStateNode.IsChecked = true

					howHideButton := &tree.Button{
						Name:            "Show State Shape",
						Icon:            string(buttons.BUTTON_visibility),
						HasToolTip:      true,
						ToolTipPosition: tree.Above,
						ToolTipText:     "Show State Shape",
						OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
							stateShape.SetIsHidden(!stateShape.GetIsHidden())
							stager.stage.Commit()
						},
					}
					diagramStateNode.Buttons = append(diagramStateNode.Buttons, howHideButton)

					if !stateShape.GetIsHidden() {
						howHideButton.Icon = string(buttons.BUTTON_visibility_off)
						howHideButton.ToolTipText = "Hide State Shape"
					}

					// range over transitions that have the state as a source or target
					// if the target is present, enable the check button
					for _, transition_ := range transitionSlice {
						if transition_.Start == state && transition_.End != nil {
							transitionNode := new(tree.Node)
							transitionNode.Name = transition_.Name + " --> " + transition_.End.Name

							if transition_.IsInRenameMode {
								transitionNode.Name = transition_.Name
							}

							transitionNode.HasCheckboxButton = true
							transitionNode.IsNodeClickable = true
							transitionNode.IsInEditMode = transition_.IsInRenameMode

							if !transition_.IsInRenameMode {
								transitionNode.Buttons = append(transitionNode.Buttons,
									&tree.Button{
										Name: transition_.GetName() + " " + string(buttons.BUTTON_edit_note),
										Icon: string(buttons.BUTTON_edit_note),
										Impl: &tree.FunctionalButtonProxy{
											OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
												transition_.IsInRenameMode = true
												stager.stage.Commit()
											},
										},
										HasToolTip:      true,
										ToolTipText:     "Rename the transition",
										ToolTipPosition: tree.Above,
									})
							} else {
								transitionNode.Buttons = append(transitionNode.Buttons,
									&tree.Button{
										Name: transition_.GetName() + " " + string(buttons.BUTTON_edit_off),
										Icon: string(buttons.BUTTON_edit_off),
										Impl: &tree.FunctionalButtonProxy{
											OnUpdated: func(_ *tree.Stage, _, _ *tree.Button) {
												transition_.IsInRenameMode = false
												stager.stage.Commit()
											},
										},
										HasToolTip:      true,
										ToolTipText:     "Cancel renaming",
										ToolTipPosition: tree.Above,
									})
							}

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

							showHideButton := &tree.Button{
								Name:            "Show Transition Shape",
								Icon:            string(buttons.BUTTON_visibility),
								HasToolTip:      true,
								ToolTipPosition: tree.Above,
								ToolTipText:     "Show Transition Shape",
								OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
									transitionShape.SetIsHidden(!transitionShape.GetIsHidden())
									stager.stage.Commit()
								},
							}
							transitionNode.Buttons = append(transitionNode.Buttons, showHideButton)
							if !transitionShape.GetIsHidden() {
								showHideButton.Icon = string(buttons.BUTTON_visibility_off)
								showHideButton.ToolTipText = "Hide Transition Shape"
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

					if frontNode.Name != stagedNode.Name {
						state.Name = frontNode.Name
						state.IsInRenameMode = false
						stager.stage.Commit()
						return
					}

					if frontNode.IsExpanded != stagedNode.IsExpanded {
						stagedNode.IsExpanded = frontNode.IsExpanded
						if frontNode.IsExpanded {
							if slices.Index(diagram.StatesWhoseNodeIsExpanded, state) == -1 {
								diagram.StatesWhoseNodeIsExpanded = append(diagram.StatesWhoseNodeIsExpanded, state)
							}
						} else {
							if idx := slices.Index(diagram.StatesWhoseNodeIsExpanded, state); idx != -1 {
								diagram.StatesWhoseNodeIsExpanded = slices.Delete(diagram.StatesWhoseNodeIsExpanded, idx, idx+1)
							}
						}
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
