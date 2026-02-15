package models

import (
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
	addShowHideCompositionButton(
		stager,
		diagram,
		product,
		product.parentProduct,
		productNode,
		diagram.map_Product_ProductShape,
		diagram.map_Product_ProductCompositionShape,
		&diagram.ProductComposition_Shapes,
	)

	for _, product := range product.SubProducts {
		stager.treePBSRecusriveInDiagram(diagram, product, productNode)
	}
}
