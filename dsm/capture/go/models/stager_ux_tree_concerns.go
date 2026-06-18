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
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("In (%d)", len(concern.Inputs)),
			IsExpanded:           slices.Index(diagram.ConcernsWhoseInputNodeIsExpanded, concern) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		concernNode.Children = append(concernNode.Children, inputProductsNode)

		inputProductsNode.OnUpdate = onUpdateExpandableNode(stager, concern, &diagram.ConcernsWhoseInputNodeIsExpanded)

		for _, product := range concern.Inputs {
			inputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			// if input task is present in diagram as well as the input product
			// display the show/hide input relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {

					inputProductNode.HasCheckboxButton = true

					taskProductKey := concernProductKey{
						Concern: concern,
						Product: product,
					}
					taskInputShape, ok := diagram.map_Concern_TaskInputShape[taskProductKey]
					inputProductNode.IsChecked = ok

					if ok {
						inputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						inputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					inputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, concern, product, &diagram.ConcernInputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskInputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					inputProductNode.Buttons = []*tree.Button{
						{
							Name: diagram.GetName(),
							Icon: string(buttons.BUTTON_visibility_off),
							// SVGIcon:         svgIconLinkVisibilityOff,
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								taskInputShape.SetIsHidden(!taskInputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskInputShape.GetIsHidden() {
							inputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							// inputProductNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
							inputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						inputProductNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}

	}
	if len(concern.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("Out (%d)", len(concern.Outputs)),
			IsExpanded:           slices.Index(diagram.ConcernssWhoseOutputNodeIsExpanded, concern) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		concernNode.Children = append(concernNode.Children, outputProductsNode)

		outputProductsNode.OnUpdate = onUpdateExpandableNode(stager, concern, &diagram.ConcernssWhoseOutputNodeIsExpanded)

		for _, product := range concern.Outputs {
			outputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			// if output task is present in diagram as well as the output product
			// display the show/hide output relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {

					outputProductNode.HasCheckboxButton = true

					taskProductKey := concernProductKey{
						Concern: concern,
						Product: product,
					}
					taskOutputShape, ok := diagram.map_Concern_ConcernOutputShape[taskProductKey]
					outputProductNode.IsChecked = ok

					if ok {
						outputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						outputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					outputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, concern, product, &diagram.ConcernOutputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskOutputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					outputProductNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								taskOutputShape.SetIsHidden(!taskOutputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskOutputShape.GetIsHidden() {
							outputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							outputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						outputProductNode.Buttons[0].IsDisabled = true
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

		stakeholdersNode.OnUpdate = onUpdateExpandableNode(stager, concern, &diagram.ConcernsWhoseStakeholderNodeIsExpanded)

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

			n.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsChecked && !stagedNode.IsChecked {
					stagedNode.IsChecked = true
					newShapeToDiagram(stakeholder, diagram, &diagram.Stakeholder_Shapes, stager.stage)
					addAssociationShapeToDiagram(stager, stakeholder, concern, &diagram.StakeholderConcernShapes)
					stager.stage.Commit()
				}
				if !frontNode.IsChecked && stagedNode.IsChecked {
					stagedNode.IsChecked = false
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

					n.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsSecondCheckboxChecked && !stagedNode.IsSecondCheckboxChecked {
							// stagedNode.IsSecondCheckboxChecked = true
							addAssociationShapeToDiagram(stager, stakeholder, concern, &diagram.StakeholderConcernShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsSecondCheckboxChecked && stagedNode.IsSecondCheckboxChecked {
							// stagedNode.IsSecondCheckboxChecked = false
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
			// 		// inputProductNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
			// 		n.Buttons[0].ToolTipText = "Show link on diagram"
			// 	}
			// } else {
			// 	n.Buttons[0].IsDisabled = true
			// }

		}
	}
}
