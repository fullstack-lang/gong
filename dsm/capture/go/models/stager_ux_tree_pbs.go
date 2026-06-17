package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDeliverableRecusriveInDiagram(diagram *Diagram, product *Deliverable, parentNode *tree.Node) {
	productNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		product,
		product.parentProduct,
		&diagram.ProductsWhoseNodeIsExpanded,
		&diagram.Product_Shapes,
		diagram.map_Product_ProductShape,
		diagram.map_Product_ProductCompositionShape,
		&diagram.ProductComposition_Shapes,
	)

	addAddItemButton(stager, &diagram.ProductsWhoseNodeIsExpanded, product, nil, productNode, &product.SubProducts, diagram, &diagram.Product_Shapes, &diagram.ProductComposition_Shapes)

	for _, product := range product.SubProducts {
		stager.treeDeliverableRecusriveInDiagram(diagram, product, productNode)
	}
}
