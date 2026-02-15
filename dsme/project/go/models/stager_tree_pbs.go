package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePBSRecusriveInDiagram(diagram *Diagram, product *Product, parentNode *tree.Node) {
	productNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		product,
		&diagram.ProductsWhoseNodeIsExpanded,
		&diagram.Product_Shapes,
		&diagram.map_Product_ProductShape,
	)

	addAddItemButton(stager, &diagram.ProductsWhoseNodeIsExpanded, product, nil, productNode, &product.SubProducts, diagram, &diagram.Product_Shapes, &diagram.ProductComposition_Shapes)

	// if product has a parent product, add a button to show/hide the link to the parent
	if parentProduct := product.parentProduct; parentProduct != nil {
		if _, ok := diagram.map_Product_ProductShape[parentProduct]; ok {
			if _, ok := diagram.map_Product_ProductShape[product]; ok {

				showHideCompositionButton := &tree.Button{
					Name: GetGongstructNameFromPointer(product) + " " + string(buttons.BUTTON_add),

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Product_ProductCompositionShape[product]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					// showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
					// 	OnUpdated: stager.OnAddProductCompositionShape(diagram, parentProduct, product),
					// }
					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onAddAssociationShape(stager, parentProduct, product, &diagram.ProductComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onRemoveAssociationShape(stager, compositionShape, &diagram.ProductComposition_Shapes),
					}
				}
				productNode.Buttons = append(productNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, product := range product.SubProducts {
		stager.treePBSRecusriveInDiagram(diagram, product, productNode)
	}
}
