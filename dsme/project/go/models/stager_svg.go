package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	stager.svgStage.Reset()

	map_Product_Rect := make(map[*Product]*svg.Rect)
	map_Task_Rect := make(map[*Task]*svg.Rect)

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
		rect := svgRect(
			stager,
			diagram,
			productShape,
			layer)
		rect.RX = 3
		map_Product_Rect[productShape.Product] = rect
	}

	for _, productCompositionShape := range diagram.ProductComposition_Shapes {
		_ = productCompositionShape
		subProduct := productCompositionShape.Product
		parentProduct := subProduct.parentProduct

		if subProduct == nil || parentProduct == nil {
			log.Panic("There should be a subProduct and a parentProduct")
		}

		startRect := map_Product_Rect[parentProduct]
		endRect := map_Product_Rect[subProduct]

		svgAssociationLink(
			stager,
			startRect, endRect,
			productCompositionShape,
			// when one clicks on the link, this is the form of the parent product
			parentProduct,
			layer,
			false)
	}

	for _, taskShape := range diagram.Task_Shapes {
		map_Task_Rect[taskShape.Task] = svgRect(
			stager,
			diagram,
			taskShape,
			layer)
	}

	for _, taskCompositionShape := range diagram.TaskComposition_Shapes {
		_ = taskCompositionShape
		subTask := taskCompositionShape.Task
		parentTask := subTask.parentTask

		if subTask == nil || parentTask == nil {
			log.Panic("There should be a subTask and a parentTask")
		}

		startRect := map_Task_Rect[parentTask]
		endRect := map_Task_Rect[subTask]

		svgAssociationLink(
			stager,
			startRect, endRect,
			taskCompositionShape,
			parentTask,
			layer,
			false)
	}

	for _, taskInputShape := range diagram.TaskInputShapes {
		task := taskInputShape.Task
		product := taskInputShape.Product

		if task == nil || product == nil {
			log.Panic("There should be a task and a product")
		}

		startRect := map_Product_Rect[product]
		endRect := map_Task_Rect[task]

		svgAssociationLink(
			stager,
			startRect, endRect,
			taskInputShape,
			task,
			layer,
			true,
		)
	}

	for _, taskOutputShape := range diagram.TaskOutputShapes {
		task := taskOutputShape.Task
		product := taskOutputShape.Product

		if task == nil || product == nil {
			log.Panic("There should be a task and a product")
		}

		startRect := map_Task_Rect[task]
		endRect := map_Product_Rect[product]

		svgAssociationLink(
			stager,
			startRect, endRect,
			taskOutputShape,
			task,
			layer,
			true,
		)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
