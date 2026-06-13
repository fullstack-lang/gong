package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProduct(diagramHierarchy *DiagramHierarchy, product *Product, parentNode *tree.Node) {
	productNodeConf := TreeNodeShapeAndLinkConfiguration[
		*Product, Product, // AT, AT_
		*ProductShape, ProductShape, // CT, CT_
		*ProductCompositionShape, ProductCompositionShape, // ACT, ACT_
		*DiagramHierarchy, // DiagramType
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Product, Product, // AT, AT_
			*ProductShape, ProductShape, // CT, CT_
			*DiagramHierarchy, // DiagramType
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Product, Product, // AT, AT_
				*DiagramHierarchy, // DiagramType
			]{
				diagram:            diagramHierarchy,
				parentNode:                  parentNode,
				element:                     product,
				parentElement:               product.parentProduct,
				elementsWhoseNodeIsExpanded: &diagramHierarchy.ProductsWhoseNodeIsExpanded,
			},
			shapes:    &diagramHierarchy.Product_Shapes,
			shapesMap: diagramHierarchy.map_Product_ProductShape,
		},
		map_Element_CompositionShape: diagramHierarchy.map_Product_ProductCompositionShape,
		compositionShapes:            &diagramHierarchy.ProductComposition_Shapes,
	}
	productNode := addNodeToTree(stager, productNodeConf)

	if product.IsImport && product.ReferencedProduct != nil {
		productNode.Name = "🔗 " + product.ReferencedProduct.Name
		productNode.CheckboxToolTipText = "Add imported product to diagram"
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
				parentNodeExpansionSliceEncoding:   &diagramHierarchy.ProductsWhoseNodeIsExpanded,
				parentElement:                      product,
			},
			receivingDiagram:      diagramHierarchy,
			sliceForNewAddedShape: &diagramHierarchy.Product_Shapes,
		},
		sliceForNewCompositionShapes: &diagramHierarchy.ProductComposition_Shapes,
	}

	addCreateItemShapeAndLinkButton(stager, conf)

	for _, product := range product.SubProducts {
		stager.treeProduct(diagramHierarchy, product, productNode)
	}
}
