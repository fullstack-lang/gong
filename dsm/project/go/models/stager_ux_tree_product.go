package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProduct(diagram *Diagram, product *Product, parentNode *tree.Node) {
	productNodeConf := TreeNodeShapeAndLinkConfiguration[
		*Product, Product, // AT, AT_
		*ProductShape, ProductShape, // CT, CT_
		*ProductCompositionShape, ProductCompositionShape, // ACT, ACT_
		*Diagram, // DiagramType
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Product, Product, // AT, AT_
			*ProductShape, ProductShape, // CT, CT_
			*Diagram, // DiagramType
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Product, Product, // AT, AT_
				*Diagram, // DiagramType
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
	productNode := addNodeToTree(stager, productNodeConf)
	productNode.IsWithPreceedingIcon = true
	productNode.PreceedingIcon = string(buttons.BUTTON_category)

	if product.IsImport && product.ReferencedProduct != nil {
		productNode.Name = "🔗 " + product.ReferencedProduct.Name
		productNode.CheckboxToolTipText = "Add imported product to diagram"
	}

	productShape, ok := diagram.map_Product_ProductShape[product]

	addLayoutButtons(stager, diagram, productNode, product, productShape, ok)

	if ok && product.ReferencedProduct != nil {
		button := &tree.Button{
			Name:            "Show type on diagram",
			Icon:            string(buttons.BUTTON_label),
			ToolTipText:     "Show type on diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				productShape.IsShowType = !productShape.IsShowType
				stager.stage.Commit()
			},
		}
		if productShape.IsShowType {
			button.Name = "Hide type on diagram"
			button.ToolTipText = "Hide type on diagram"
			button.Icon = string(buttons.BUTTON_label_off)
		}
		if productNode.Menu == nil {
			productNode.Menu = &tree.Menu{Name: "Menu"}
		}
		productNode.Menu.Buttons = append(productNode.Menu.Buttons, button)
	}

	conf := ItemShapeAndLinkButtonConfiguration[
		Product, *Product, // AT, PAT (Added Element)
		Product, *Product, // ParentAT, PParentAT (Parent Element)
		ProductShape, *ProductShape, // CT, PCT (Concrete Shape)
		ProductCompositionShape, *ProductCompositionShape, // ACT, PACT (Association Shape),
	]{
		ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
			Product, *Product, // AT, PAT (Added Element)
			Product, *Product, // ParentAT, PParentAT (Parent Element)
			ProductShape, *ProductShape, // CT, PCT (Concrete Shape)
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Product, *Product, // AT, PAT (Added Element)
				Product, *Product, // ParentAT, PParentAT (Parent Element)
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
		},
		sliceForNewCompositionShapes: &diagram.ProductComposition_Shapes,
	}

	addCreateItemShapeAndLinkButton(stager, conf)

	if product.ReferencedProduct != nil {
		refProduct := product.ReferencedProduct
		refNode := &tree.Node{
			Name:                    "Type: " + refProduct.GetName(),
			IsNodeClickable:         true,
			IsWithPreceedingIcon:    true,
			PreceedingIcon:          string(buttons.BUTTON_category),
			IsInEditMode:            refProduct.GetIsInRenameMode(),
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Right,
		}
		productNode.Children = append(productNode.Children, refNode)
		addRenameButton(refProduct, refNode, stager)
		refNode.OnNameChange = stager.onNameChange(refProduct)
		refNode.OnClick = onNodeClicked(stager, refProduct)

		// if both product and referenced product have shapes in the diagram
		if _, okStart := diagram.map_Product_ProductShape[product]; okStart {
			if _, okEnd := diagram.map_Product_ProductShape[refProduct]; okEnd {
				refNode.HasCheckboxButton = true

				key := productReferenceKey{
					Product:           product,
					ReferencedProduct: refProduct,
				}
				productRefShape, okLink := diagram.map_Product_ProductReferenceShape[key]
				refNode.IsChecked = okLink

				if okLink {
					refNode.CheckboxToolTipText = "Uncheck to remove reference link from diagram"
				} else {
					refNode.CheckboxToolTipText = "Check to add reference link to diagram"
				}

				refNode.OnIsCheckedChanged = func(isChecked bool) {
					if isChecked {
						addAssociationShapeToDiagram(stager, product, refProduct, &diagram.ProductReference_Shapes)
						stager.stage.Commit()
					} else {
						if productRefShape != nil {
							productRefShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}
				}

				refNode.Buttons = []*tree.Button{
					{
						Name:            diagram.GetName(),
						Icon:            string(buttons.BUTTON_visibility_off),
						ToolTipText:     "Hide link from diagram",
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
						OnClick: func() {
							if productRefShape != nil {
								productRefShape.SetIsHidden(!productRefShape.GetIsHidden())
								stager.stage.Commit()
							}
						},
					},
				}
				if okLink {
					if productRefShape.GetIsHidden() {
						refNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
						refNode.Buttons[0].ToolTipText = "Show link on diagram"
					}
				} else {
					refNode.Buttons[0].IsDisabled = true
				}
			}
		}
	}

	for _, product := range product.SubProducts {
		stager.treeProduct(diagram, product, productNode)
	}
}
