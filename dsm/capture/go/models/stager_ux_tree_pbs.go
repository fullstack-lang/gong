package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDeliverableRecusriveInDiagram(diagram *Diagram, product *Deliverable, parentNode *tree.Node) {
	confNode := TreeNodeShapeAndLinkConfiguration[
		*Deliverable, Deliverable,
		*ProductShape, ProductShape,
		*ProductCompositionShape, ProductCompositionShape,
		*Diagram,
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Deliverable, Deliverable,
			*ProductShape, ProductShape,
			*Diagram,
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Deliverable, Deliverable,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  parentNode,
				element:                     product,
				parentElement:               product.parentProduct,
				elementsWhoseNodeIsExpanded: &diagram.ProductsWhoseNodeIsExpanded,
			},
			shapes:    &diagram.Product_Shapes,
			shapesMap: diagram.map_Product_ProductShape,
		},
		map_Element_CompositionShape: diagram.map_Product_ProductCompositionShape,
		compositionShapes:            &diagram.ProductComposition_Shapes,
	}
	productNode := addNodeToTree(stager, confNode)

	confSubProducts := ItemAndShapeButtonConfiguration[
		Deliverable, *Deliverable,
		Deliverable, *Deliverable,
		ProductShape, *ProductShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Deliverable, *Deliverable,
			Deliverable, *Deliverable,
		]{
			parentNode:                         productNode,
			sliceForNewAddedItem:               &product.SubProducts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
			parentNodeExpansionSliceEncoding:   &diagram.ProductsWhoseNodeIsExpanded,
			parentElement:                      product,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Product_Shapes,
	}
	addCreateItemAndShapeButton(stager, confSubProducts)

	for _, product := range product.SubProducts {
		stager.treeDeliverableRecusriveInDiagram(diagram, product, productNode)
	}

	if len(product.Concepts) > 0 {
		conceptsNode := &tree.Node{
			Name:                 fmt.Sprintf("Concepts (%d)", len(product.Concepts)),
			IsExpanded:           slices.Index(diagram.ProductsWhoseConceptsNodeIsExpanded, product) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		productNode.Children = append(productNode.Children, conceptsNode)

		conceptsNode.OnUpdate = onUpdateExpandableNode(stager, product, &diagram.ProductsWhoseConceptsNodeIsExpanded)

		for _, concept := range product.Concepts {
			conceptNode := &tree.Node{
				Name:                    concept.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			conceptsNode.Children = append(conceptsNode.Children, conceptNode)

			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Concept_ConceptShape[concept]; ok {

					conceptNode.HasCheckboxButton = true

					key := deliverableConceptKey{
						Deliverable: product,
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
							addAssociationShapeToDiagram(stager, product, concept, &diagram.DeliverableConceptShapes)
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
