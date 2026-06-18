package models

import (
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
}
