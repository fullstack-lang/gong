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

	if product.IsImport && product.ReferencedProduct != nil {
		productNode.Name = "🔗 " + product.ReferencedProduct.Name
		productNode.CheckboxToolTipText = "Add imported product to diagram"
	}

	toggleAbstractLayoutButton := &tree.Button{
		Name: "Toggle Abstract Layout Direction to " + func() string {
			if product.LayoutDirection == Vertical {
				return "Horizontal"
			} else {
				return "Vertical"
			}
		}(),
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			if product.LayoutDirection == Vertical {
				product.LayoutDirection = Horizontal
			} else {
				product.LayoutDirection = Vertical
			}
			stager.stage.Commit()
		},
	}

	if product.LayoutDirection == Vertical {
		toggleAbstractLayoutButton.Icon = string(buttons.BUTTON_swap_horiz)
		toggleAbstractLayoutButton.ToolTipText = "Set layout to Horizontal"
	} else {
		toggleAbstractLayoutButton.Icon = string(buttons.BUTTON_swap_vert)
		toggleAbstractLayoutButton.ToolTipText = "Set layout to Vertical"
	}

	if productNode.Menu == nil {
		productNode.Menu = &tree.Menu{Name: "Menu"}
	}
	productNode.Menu.Buttons = append(productNode.Menu.Buttons, toggleAbstractLayoutButton)

	productShape, isPresent := diagram.map_Product_ProductShape[product]

	if isPresent {
		toggleLayoutButton := &tree.Button{
			Name: "Toggle Concrete Layout Direction to " + func() string {
				if productShape.LayoutDirection == Vertical {
					return "Horizontal"
				} else {
					return "Vertical"
				}
			}(),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				if productShape.LayoutDirection == Vertical {
					productShape.LayoutDirection = Horizontal
				} else {
					productShape.LayoutDirection = Vertical
				}
				stager.stage.Commit()
			},
		}

		if productShape.LayoutDirection == Vertical {
			toggleLayoutButton.Icon = string(buttons.BUTTON_swap_horiz)
			toggleLayoutButton.ToolTipText = "Set concrete layout to Horizontal"
		} else {
			toggleLayoutButton.Icon = string(buttons.BUTTON_swap_vert)
			toggleLayoutButton.ToolTipText = "Set concrete layout to Vertical"
		}

		toggleOverrideButton := &tree.Button{
			Name: "Toggle Override Layout Direction",
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				productShape.OverideLayoutDirection = !productShape.OverideLayoutDirection
				stager.stage.Commit()
			},
		}

		if productShape.OverideLayoutDirection {
			toggleOverrideButton.Icon = string(buttons.BUTTON_check_box)
			toggleOverrideButton.ToolTipText = "Disable layout override"
		} else {
			toggleOverrideButton.Icon = string(buttons.BUTTON_check_box_outline_blank)
			toggleOverrideButton.ToolTipText = "Enable layout override"
		}

		productNode.Menu.Buttons = append(productNode.Menu.Buttons, toggleLayoutButton)
		productNode.Menu.Buttons = append(productNode.Menu.Buttons, toggleOverrideButton)
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

	for _, product := range product.SubProducts {
		stager.treeProduct(diagram, product, productNode)
	}
}
