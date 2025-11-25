package models

import (
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
			addButton.Impl = &SateMachineAddDiagramButtonProxy{
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
						Impl: &DiagramRenameButtonProxy{
							stager:     stager,
							diagram:    diagram,
							buttonType: RENAME,
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
						Impl: &DiagramRenameButtonProxy{
							stager:     stager,
							diagram:    diagram,
							buttonType: RENAME_CANCEL,
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

			for _, state_ := range stateMachine.States {
				diagramStateNode := new(tree.Node)
				diagramStateNode.Name = state_.Name
				diagramStateNode.HasCheckboxButton = true
				diagramStateNode.IsNodeClickable = true

				var stateShape *StateShape
				if stateShape_, ok := map_State__StateShape[state_]; ok {
					diagramStateNode.IsChecked = true
					stateShape = stateShape_
					diagramStateNode.IsExpanded = stateShape_.IsExpanded

					proxy := new(Diagram_Tree_StateShapeProxy)
					diagramStateNode.Impl = proxy
					proxy.stager = stager
					proxy.stateShape = stateShape_

					// range over transitions that have the state as a source or target
					// if the target is present, enable the check button
					for _, transition_ := range transitionSlice {
						if transition_.Start == state_ && transition_.End != nil {
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
							proxy.transitionShape = transitionShape

							// disable if end is not present
							if _, ok := map_State__StateShape[transition_.End]; !ok {
								transitionNode.IsCheckboxDisabled = true
							}

							diagramStateNode.Children = append(diagramStateNode.Children, transitionNode)
						}
					}
				}

				p := new(DiagramTree_State_Proxy)
				p.stager = stager
				p.stateShape = stateShape
				p.state = state_
				p.diagram = diagram
				diagramStateNode.Impl = p

				diagramNode.Children = append(diagramNode.Children, diagramStateNode)
			}
		}
	}

	tree.StageBranch(treeStage, treeInstance)

	treeStage.Commit()
}
