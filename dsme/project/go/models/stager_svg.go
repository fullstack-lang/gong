package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	stager.svgStage.Reset()

	map_Product_Rect := make(map[*Product]*svg.Rect)

	svgStage := stager.svgStage
	svgObject := (&svg.SVG{Name: `SVG`})

	var diagram *Diagram
	{
		for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
			if diagram_.IsChecked {
				diagram = diagram_
			}
		}
	}

	if diagram == nil {
		stager.svgStage.Commit()
		return
	}

	// To be implemented when drawing compositions between Products
	// svgImpl := &svgProxy{
	// 	stager:                   stager,
	// 	diagram:                  diagram,
	// 	map_SvgRect_ProductShape: make(map[*svg.Rect]*ProductShape),
	// }
	// svgObject.Impl = svgImpl

	svgObject.Name = diagram.Name
	svgObject.IsEditable = diagram.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	ProductsSet := make(map[*Product]any)

	for _, ProductShape_ := range diagram.Product_Shapes {
		ProductsSet[ProductShape_.Product] = true
	}

	for _, ProductShape := range diagram.Product_Shapes {

		if ProductShape.Product == nil {
			continue
		}
		Product := ProductShape.Product

		rect := stager.svgGenerateRect(
			diagram,
			ProductShape,
			layer)
		map_Product_Rect[Product] = rect

		// To be implemented when drawing compositions between Products
		// svgImpl.map_SvgRect_ProductShape[rect] = ProductShape
	}

	for _, transtionShape := range diagram.Composition_Shapes {
		_ = transtionShape
		// subProduct := transtionShape.Product

		// if subProduct == nil || subProduct.Start == nil || subProduct.End == nil {
		// 	continue
		// }

		// startRect := map_Product_Rect[subProduct.Start]
		// endRect := map_Product_Rect[subProduct]

		// stager.svgGenerateLink(
		// 	startRect, endRect,
		// 	&transtionShape.LinkShape, subProduct, layer, false)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
