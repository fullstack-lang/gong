package models

import (
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
	productNode := addNodeToTreeWithConf(stager, productNodeConf)

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
