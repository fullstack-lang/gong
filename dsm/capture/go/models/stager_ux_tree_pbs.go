package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDeliverableRecusriveInDiagram(diagram *Diagram, deliverable *Deliverable, parentNode *tree.Node) {
	confNode := TreeNodeShapeAndLinkConfiguration[
		*Deliverable, Deliverable,
		*DeliverableShape, DeliverableShape,
		*DeliverableCompositionShape, DeliverableCompositionShape,
		*Diagram,
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Deliverable, Deliverable,
			*DeliverableShape, DeliverableShape,
			*Diagram,
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Deliverable, Deliverable,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  parentNode,
				element:                     deliverable,
				parentElement:               deliverable.parentDeliverable,
				elementsWhoseNodeIsExpanded: &diagram.DeliverablesWhoseNodeIsExpanded,
			},
			shapes:    &diagram.Deliverable_Shapes,
			shapesMap: diagram.map_Deliverable_DeliverableShape,
		},
		map_Element_CompositionShape: diagram.map_Deliverable_DeliverableCompositionShape,
		compositionShapes:            &diagram.DeliverableComposition_Shapes,
	}
	deliverableNode := addNodeToTree(stager, confNode)

	confSubDeliverables := ItemAndShapeButtonConfiguration[
		Deliverable, *Deliverable,
		Deliverable, *Deliverable,
		DeliverableShape, *DeliverableShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Deliverable, *Deliverable,
			Deliverable, *Deliverable,
		]{
			parentNode:                         deliverableNode,
			sliceForNewAddedItem:               &deliverable.SubDeliverables,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
			parentNodeExpansionSliceEncoding:   &diagram.DeliverablesWhoseNodeIsExpanded,
			parentElement:                      deliverable,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Deliverable_Shapes,
	}
	addCreateItemAndShapeButton(stager, confSubDeliverables)

	for _, deliverable := range deliverable.SubDeliverables {
		stager.treeDeliverableRecusriveInDiagram(diagram, deliverable, deliverableNode)
	}

	if len(deliverable.Concepts) > 0 {
		conceptsNode := &tree.Node{
			Name:                 fmt.Sprintf("Concepts (%d)", len(deliverable.Concepts)),
			IsExpanded:           slices.Index(diagram.DeliverablesWhoseConceptsNodeIsExpanded, deliverable) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		deliverableNode.Children = append(deliverableNode.Children, conceptsNode)

		conceptsNode.OnUpdate = onUpdateExpandableNode(stager, deliverable, &diagram.DeliverablesWhoseConceptsNodeIsExpanded)

		for _, concept := range deliverable.Concepts {
			conceptNode := &tree.Node{
				Name:                    concept.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			conceptsNode.Children = append(conceptsNode.Children, conceptNode)

			if _, ok := diagram.map_Deliverable_DeliverableShape[deliverable]; ok {
				if _, ok := diagram.map_Concept_ConceptShape[concept]; ok {

					conceptNode.HasCheckboxButton = true

					key := deliverableConceptKey{
						Deliverable: deliverable,
						Concept:     concept,
					}
					associationShape, ok := diagram.map_DeliverableConceptKey_DeliverableConceptShape[key]
					conceptNode.IsChecked = ok

					if ok {
						conceptNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						conceptNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					conceptNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, deliverable, concept, &diagram.DeliverableConceptShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					conceptNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								associationShape.SetIsHidden(!associationShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					}
					if ok {
						if associationShape.GetIsHidden() {
							conceptNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							conceptNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						conceptNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
