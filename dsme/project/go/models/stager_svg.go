package models

import (
	"log"

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

	for _, productShape := range diagram.Product_Shapes {
		stager.svgGenerateProductRect(
			diagram,
			productShape,
			layer)
	}

	for _, compositionShape := range diagram.ProductComposition_Shapes {
		_ = compositionShape
		subProduct := compositionShape.Product
		parentProduct := subProduct.parentProduct

		if subProduct == nil || parentProduct == nil {
			log.Panic("There should be a subProduct and a parentProduct")
		}

		startRect := map_Product_Rect[parentProduct]
		endRect := map_Product_Rect[subProduct]

		stager.svgGenerateLink(
			startRect, endRect,
			&compositionShape.LinkShape, subProduct, layer, false)
	}

	for _, taskShape := range diagram.Task_Shapes {
		stager.svgGenerateTaskRect(
			diagram,
			taskShape,
			layer)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
