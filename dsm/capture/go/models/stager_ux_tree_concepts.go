package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeConceptBSinDiagram(diagram *Diagram, concept *Concept, parentNode *tree.Node) {
	confNode := TreeNodeAndShapeConfigurationWithoutLink[
		*Concept, Concept,
		*Concept, Concept, // Parent is not used
		*ConceptShape, ConceptShape,
		*Diagram,
	]{
		diagram:                     diagram,
		parentNode:                  parentNode,
		element:                     concept,
		parentElement:               nil,
		elementsWhoseNodeIsExpanded: &diagram.ConceptsWhoseNodeIsExpanded,
		shapes:                      &diagram.Concept_Shapes,
		shapesMap:                   diagram.map_Concept_ConceptShape,
	}
	conceptNode := addNodeToTreeWithoutLink(stager, confNode)

	var deliverables []*Deliverable
	for product := range *GetGongstructInstancesSet[Deliverable](stager.stage) {
		if slices.Contains(product.Concepts, concept) {
			deliverables = append(deliverables, product)
		}
	}

	if len(deliverables) > 0 {
		slices.SortFunc(deliverables, func(a, b *Deliverable) int {
			return cmp.Compare(a.Name, b.Name)
		})

		deliverablesNode := &tree.Node{
			Name:                 fmt.Sprintf("Deliverables (%d)", len(deliverables)),
			IsExpanded:           slices.Index(diagram.ConceptsWhoseDeliverablesNodeIsExpanded, concept) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		conceptNode.Children = append(conceptNode.Children, deliverablesNode)

		deliverablesNode.OnUpdate = onUpdateExpandableNode(stager, concept, &diagram.ConceptsWhoseDeliverablesNodeIsExpanded)

		for _, deliverable := range deliverables {
			deliverableNode := &tree.Node{
				Name:                    deliverable.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			deliverablesNode.Children = append(deliverablesNode.Children, deliverableNode)

			if _, ok := diagram.map_Product_ProductShape[deliverable]; ok {
				if _, ok := diagram.map_Concept_ConceptShape[concept]; ok {

					deliverableNode.HasCheckboxButton = true

					key := deliverableConceptKey{
						Deliverable: deliverable,
						Concept:     concept,
					}
					associationShape, ok := diagram.map_DeliverableConceptKey_DeliverableConceptShape[key]
					deliverableNode.IsChecked = ok

					if ok {
						deliverableNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						deliverableNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					deliverableNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
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

					deliverableNode.Buttons = []*tree.Button{
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
							deliverableNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							deliverableNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						deliverableNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
