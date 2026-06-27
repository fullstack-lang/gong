package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	treeInstance := &tree.Tree{Name: string(DiagramTreeName)}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(stager.GetRootLibrary(), &treeInstance.RootNodes)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func (stager *Stager) treeLibrary(library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:            library.Name,
		IsExpanded:      library.IsExpandedTmp,
		IsNodeClickable: true,
		IsInEditMode:    library.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if library != stager.GetRootLibrary() {
		addRenameButton(library, libraryNode, stager)
	}
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpandedTmp)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)

	//
	// SubLibraries
	//
	subLibrariesNode := &tree.Node{
		Name:            "Sub Libraries",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsSubLibrariesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, subLibrariesNode)
	subLibrariesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSubLibrariesNodeExpanded)
	subLibrariesNode.OnClick = onNodeClicked(stager, library)

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(subLibrary, &subLibrariesNode.Children)
	}

	// add sub library button
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         subLibrariesNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSubLibrariesNodeExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	//
	// State Machines
	//
	stateMachinesNode := &tree.Node{
		Name:            "State Machines",
		FontStyle:       tree.ITALIC,
		IsExpanded:      library.IsStateMachinesNodeExpanded,
		IsNodeClickable: true,
	}
	libraryNode.Children = append(libraryNode.Children, stateMachinesNode)
	stateMachinesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsStateMachinesNodeExpanded)
	stateMachinesNode.OnClick = onNodeClicked(stager, library)

	// add a statemachine to the library button
	confRootStateMachines := ItemButtonConfiguration[
		StateMachine, *StateMachine,
		Library, *Library,
	]{
		parentNode:                         stateMachinesNode,
		sliceForNewAddedItem:               &library.RootStateMachines,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsStateMachinesNodeExpanded,
	}
	addCreateItemButton(stager, confRootStateMachines)

	for _, stateMachine := range library.RootStateMachines {
		stager.treeStateMachines(stateMachine, stateMachinesNode, &library.StateMachinesWhoseNodeIsExpanded)
	}
}

func (stager *Stager) treeStateMachines(
	stateMachine *StateMachine,
	parentNode *tree.Node,
	stateMachinesWhoseNodeIsExpanded *[]*StateMachine,
) {
	stateMachineNode := &tree.Node{
		Name:            stateMachine.GetName(),
		IsExpanded:      slices.Contains(*stateMachinesWhoseNodeIsExpanded, stateMachine),
		IsNodeClickable: true,
		IsInEditMode:    stateMachine.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, stateMachineNode)

	addRenameButton(stateMachine, stateMachineNode, stager)
	stateMachineNode.OnNameChange = stager.onNameChange(stateMachine)
	stateMachineNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, stateMachine, stateMachinesWhoseNodeIsExpanded)
	stateMachineNode.OnClick = onNodeClicked(stager, stateMachine)

	{
		addButton := &tree.Button{
			Name:            "Diagram Add" + " " + string(buttons.BUTTON_add_box),
			Icon:            string(buttons.BUTTON_add_box),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Add a Diagram to the state machine",
			OnClick: func() {
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

	transitionsSet := *GetGongstructInstancesSet[Transition](stager.stage)
	transitionSlice := SortGongstructSetByName(transitionsSet)

	for _, diagram := range stateMachine.Diagrams {
		diagramNode := new(tree.Node)
		stateMachineNode.Children = append(stateMachineNode.Children, diagramNode)
		diagramNode.Name = diagram.Name
		diagramNode.IsChecked = diagram.IsChecked
		diagramNode.IsExpanded = diagram.IsExpanded

		diagramNode.IsInEditMode = diagram.isInRenameMode

		diagramNode.IsNodeClickable = true
		diagramProxy := new(Diagram_Tree_DiagramProxy)
		diagramProxy.stager = stager
		diagramProxy.diagram = diagram

		diagramNode.HasCheckboxButton = true
		diagramNode.Impl = diagramProxy

		if !diagram.isInRenameMode {
			diagramNode.Buttons = append(diagramNode.Buttons,
				&tree.Button{
					Name: diagram.GetName() + " " + string(buttons.BUTTON_edit_note),
					Icon: string(buttons.BUTTON_edit_note),
					OnClick: func() {
						diagram.isInRenameMode = true
						stager.stage.Commit()
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
					OnClick: func() {
						diagram.isInRenameMode = false
						stager.stage.Commit()
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
				OnClick: func() {
					proxy := &DiagramCopyButtonProxy{
						stager:  stager,
						diagram: diagram,
					}
					proxy.ButtonUpdated(nil, nil, nil)
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
				OnClick: func() {
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
			diagramStateNode.IsInEditMode = state.isInRenameMode
			diagramStateNode.IsExpanded = slices.Contains(diagram.StatesWhoseNodeIsExpanded, state)

			if !state.isInRenameMode {
				diagramStateNode.Buttons = append(diagramStateNode.Buttons,
					&tree.Button{
						Name: state.GetName() + " " + string(buttons.BUTTON_edit_note),
						Icon: string(buttons.BUTTON_edit_note),
						OnClick: func() {
							state.isInRenameMode = true
							stager.stage.Commit()
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
						OnClick: func() {
							state.isInRenameMode = false
							stager.stage.Commit()
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
					OnClick: func() {
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

						if transition_.isInRenameMode {
							transitionNode.Name = transition_.Name
						}

						transitionNode.HasCheckboxButton = true
						transitionNode.IsNodeClickable = true
						transitionNode.IsInEditMode = transition_.isInRenameMode

						if !transition_.isInRenameMode {
							transitionNode.Buttons = append(transitionNode.Buttons,
								&tree.Button{
									Name: transition_.GetName() + " " + string(buttons.BUTTON_edit_note),
									Icon: string(buttons.BUTTON_edit_note),
									OnClick: func() {
										transition_.isInRenameMode = true
										stager.stage.Commit()
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
									OnClick: func() {
										transition_.isInRenameMode = false
										stager.stage.Commit()
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
							OnClick: func() {
								if transitionShape != nil {
									transitionShape.SetIsHidden(!transitionShape.GetIsHidden())
									stager.stage.Commit()
								}
							},
						}
						transitionNode.Buttons = append(transitionNode.Buttons, showHideButton)
						if transitionShape != nil && !transitionShape.GetIsHidden() {
							showHideButton.Icon = string(buttons.BUTTON_visibility_off)
							showHideButton.ToolTipText = "Hide Transition Shape"
						}

						diagramStateNode.Children = append(diagramStateNode.Children, transitionNode)
					}
				}
			}

			diagramStateNode.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					// add the State_Shape
					if stateShape != nil {
						log.Fatalln("adding a shape to an already state shape")
					}

					newStateShapeToDiagram(state, diagram).Stage(stager.stage)
					stager.stage.Commit()
				} else {
					// one need to remove the State_Shape
					if stateShape == nil {
						log.Fatalln("remove a non existing shape to state")
					}

					stateShape.Unstage(stager.stage)
					idx := slices.Index(diagram.State_Shapes, stateShape)
					diagram.State_Shapes = slices.Delete(diagram.State_Shapes, idx, idx+1)

					stager.stage.Commit()
				}
			}

			diagramStateNode.OnNameChange = func(newName string) {
				state.Name = newName
				state.isInRenameMode = false
				stager.stage.Commit()
			}

			diagramStateNode.OnIsExpandedChange = func(isExpanded bool) {
				if isExpanded {
					if slices.Index(diagram.StatesWhoseNodeIsExpanded, state) == -1 {
						diagram.StatesWhoseNodeIsExpanded = append(diagram.StatesWhoseNodeIsExpanded, state)
					}
				} else {
					if idx := slices.Index(diagram.StatesWhoseNodeIsExpanded, state); idx != -1 {
						diagram.StatesWhoseNodeIsExpanded = slices.Delete(diagram.StatesWhoseNodeIsExpanded, idx, idx+1)
					}
				}
				stager.stage.Commit()
			}

			diagramStateNode.OnClick = func(frontNode *tree.Node) {
				stager.probeForm.FillUpFormFromGongstruct(state, "State")
			}

			diagramNode.Children = append(diagramNode.Children, diagramStateNode)
		}
	}
}
