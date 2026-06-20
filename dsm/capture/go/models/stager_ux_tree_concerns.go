package models

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeConcernBSinDiagram(diagram *Diagram, concern *Concern, parentNode *tree.Node) {
	stage := stager.stage

	confNode := TreeNodeShapeAndLinkConfiguration[
		*Concern, Concern,
		*ConcernShape, ConcernShape,
		*ConcernCompositionShape, ConcernCompositionShape,
		*Diagram,
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Concern, Concern,
			*ConcernShape, ConcernShape,
			*Diagram,
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Concern, Concern,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  parentNode,
				element:                     concern,
				parentElement:               concern.parentConcern,
				elementsWhoseNodeIsExpanded: &diagram.ConcernsWhoseNodeIsExpanded,
			},
			shapes:    &diagram.Concern_Shapes,
			shapesMap: diagram.map_Concern_ConcernShape,
		},
		map_Element_CompositionShape: diagram.map_Concern_ConcernCompositionShape,
		compositionShapes:            &diagram.ConcernComposition_Shapes,
	}
	concernNode := addNodeToTree(stager, confNode)

	confSubConcerns := ItemAndShapeButtonConfiguration[
		Concern, *Concern,
		Concern, *Concern,
		ConcernShape, *ConcernShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Concern, *Concern,
			Concern, *Concern,
		]{
			parentNode:                         concernNode,
			sliceForNewAddedItem:               &concern.SubConcerns,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
			parentNodeExpansionSliceEncoding:   &diagram.ConcernsWhoseNodeIsExpanded,
			parentElement:                      concern,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Concern_Shapes,
	}
	addCreateItemAndShapeButton(stager, confSubConcerns)

	for _, task := range concern.SubConcerns {
		stager.treeConcernBSinDiagram(diagram, task, concernNode)
	}

	if len(concern.Inputs) > 0 {
		inputDeliverablesNode := &tree.Node{
			Name:                 fmt.Sprintf("In (%d)", len(concern.Inputs)),
			IsExpanded:           slices.Index(diagram.ConcernsWhoseInputNodeIsExpanded, concern) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		concernNode.Children = append(concernNode.Children, inputDeliverablesNode)

		setCallbacksExpandableNode(stager, inputDeliverablesNode, concern, &diagram.ConcernsWhoseInputNodeIsExpanded)

		for _, deliverable := range concern.Inputs {
			inputDeliverableNode := &tree.Node{
				Name:                    deliverable.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			inputDeliverablesNode.Children = append(inputDeliverablesNode.Children, inputDeliverableNode)

			// if input task is present in diagram as well as the input deliverable
			// display the show/hide input relation button
			if _, ok := diagram.map_Deliverable_DeliverableShape[deliverable]; ok {
				if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {

					inputDeliverableNode.HasCheckboxButton = true

					taskDeliverableKey := concernDeliverableKey{
						Concern: concern,
						Deliverable: deliverable,
					}
					taskInputShape, ok := diagram.map_Concern_TaskInputShape[taskDeliverableKey]
					inputDeliverableNode.IsChecked = ok

					if ok {
						inputDeliverableNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						inputDeliverableNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					inputDeliverableNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, concern, deliverable, &diagram.ConcernInputShapes)
							stager.stage.Commit()
						} else {
							taskInputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					inputDeliverableNode.Buttons = []*tree.Button{
						{
							Name: diagram.GetName(),
							Icon: string(buttons.BUTTON_visibility_off),
							// SVGIcon:         svgIconLinkVisibilityOff,
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								taskInputShape.SetIsHidden(!taskInputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskInputShape.GetIsHidden() {
							inputDeliverableNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							// inputDeliverableNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
							inputDeliverableNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						inputDeliverableNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}

	}
	if len(concern.Outputs) > 0 {
		outputDeliverablesNode := &tree.Node{
			Name:                 fmt.Sprintf("Out (%d)", len(concern.Outputs)),
			IsExpanded:           slices.Index(diagram.ConcernssWhoseOutputNodeIsExpanded, concern) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		concernNode.Children = append(concernNode.Children, outputDeliverablesNode)

		setCallbacksExpandableNode(stager, outputDeliverablesNode, concern, &diagram.ConcernssWhoseOutputNodeIsExpanded)

		for _, deliverable := range concern.Outputs {
			outputDeliverableNode := &tree.Node{
				Name:                    deliverable.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			outputDeliverablesNode.Children = append(outputDeliverablesNode.Children, outputDeliverableNode)

			// if output task is present in diagram as well as the output deliverable
			// display the show/hide output relation button
			if _, ok := diagram.map_Deliverable_DeliverableShape[deliverable]; ok {
				if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {

					outputDeliverableNode.HasCheckboxButton = true

					taskDeliverableKey := concernDeliverableKey{
						Concern: concern,
						Deliverable: deliverable,
					}
					taskOutputShape, ok := diagram.map_Concern_ConcernOutputShape[taskDeliverableKey]
					outputDeliverableNode.IsChecked = ok

					if ok {
						outputDeliverableNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						outputDeliverableNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					outputDeliverableNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, concern, deliverable, &diagram.ConcernOutputShapes)
							stager.stage.Commit()
						} else {
							taskOutputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					outputDeliverableNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								taskOutputShape.SetIsHidden(!taskOutputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskOutputShape.GetIsHidden() {
							outputDeliverableNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							outputDeliverableNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						outputDeliverableNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}

	stakeholders := stager.map_Concern_Stakeholder[concern]

	slices.SortFunc(stakeholders, func(a, b *Stakeholder) int {
		return cmp.Compare(a.Name, b.Name)
	})
	if len(stakeholders) > 0 {
		stakeholdersNode := &tree.Node{
			Name:                 fmt.Sprintf("Stakeholders (%d)", len(stakeholders)),
			IsExpanded:           slices.Index(diagram.ConcernsWhoseStakeholderNodeIsExpanded, concern) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		concernNode.Children = append(concernNode.Children, stakeholdersNode)

		setCallbacksExpandableNode(stager, stakeholdersNode, concern, &diagram.ConcernsWhoseStakeholderNodeIsExpanded)

		for _, stakeholder := range stakeholders {
			n := &tree.Node{
				Name:                    stakeholder.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
				HasSecondCheckboxButton: true,
			}
			stakeholdersNode.Children = append(stakeholdersNode.Children, n)

			n.HasCheckboxButton = true

			stakeholderShape, ok := diagram.map_Stakeholder_StakeholderShape[stakeholder]
			n.IsChecked = ok

			if ok {
				n.CheckboxToolTipText = "Uncheck to remove shape from diagram"
			} else {
				n.CheckboxToolTipText = "Check to add shape to diagram"
			}

			n.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					newShapeToDiagram(stakeholder, diagram, &diagram.Stakeholder_Shapes, stager.stage)
					addAssociationShapeToDiagram(stager, stakeholder, concern, &diagram.StakeholderConcernShapes)
					stager.stage.Commit()
				} else {
					stakeholderShape.UnstageVoid(stager.stage)
					stager.stage.Commit()
				}
			}

			if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {
				if _, ok := diagram.map_Stakeholder_StakeholderShape[stakeholder]; ok {

					n.IsSecondCheckboxDisabled = false

					key := stakeholderConcernKey{Stakeholder: stakeholder, Concern: concern}
					associationShape, ok := diagram.map_StakeholderConcernKey_StakeholderConcernShape[key]
					n.IsSecondCheckboxChecked = ok

					if ok {
						n.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						n.CheckboxToolTipText = "Check to add shape to diagram"
					}

					n.OnIsSecondCheckboxCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, stakeholder, concern, &diagram.StakeholderConcernShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}
				}
			}

			// n.Buttons = []*tree.Button{
			// 	{
			// 		Name: diagram.GetName(),
			// 		Icon: string(buttons.BUTTON_visibility_off),
			// 		// SVGIcon:         svgIconLinkVisibilityOff,
			// 		ToolTipText: func() string {
			// 			return "Remove link from diagram"
			// 		}(),
			// 		HasToolTip:      true,
			// 		ToolTipPosition: tree.Right,
			// 		OnClick: func() {
			// 			stakeholderShape.SetIsHidden(!stakeholderShape.GetIsHidden())
			// 			stage.Commit()
			// 		},
			// 	},
			// }
			// if ok {
			// 	if stakeholderShape.GetIsHidden() {
			// 		n.Buttons[0].Icon = string(buttons.BUTTON_visibility)
			// 		// inputDeliverableNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
			// 		n.Buttons[0].ToolTipText = "Show link on diagram"
			// 	}
			// } else {
			// 	n.Buttons[0].IsDisabled = true
			// }

		}
	}
}

