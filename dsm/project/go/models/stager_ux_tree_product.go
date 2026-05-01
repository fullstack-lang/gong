package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProduct(diagram *Diagram, product *Product, parentNode *tree.Node) {
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
