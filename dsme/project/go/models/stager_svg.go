package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	stager.svgStage.Reset()

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

	diagram.map_Product_Rect = make(map[*Product]*svg.Rect)
	diagram.map_Task_Rect = make(map[*Task]*svg.Rect)
	diagram.map_Note_Rect = make(map[*Note]*svg.Rect)
	diagram.map_SvgRect_ProductShape = make(map[*svg.Rect]*ProductShape)
	diagram.map_SvgRect_TaskShape = make(map[*svg.Rect]*TaskShape)
	diagram.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)

	// to implement association between abstract elements by mouse drag
	svgImpl := &svgProxy{
		stager:  stager,
		diagram: diagram,
	}
	svgObject.Impl = svgImpl

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
		diagram.map_Product_Rect[productShape.Product] = rect
		diagram.map_SvgRect_ProductShape[rect] = productShape
	}

	for _, productCompositionShape := range diagram.ProductComposition_Shapes {
		_ = productCompositionShape
		subProduct := productCompositionShape.Product
		parentProduct := subProduct.parentProduct

		if subProduct == nil || parentProduct == nil {
			log.Panic("There should be a subProduct and a parentProduct")
		}

		startRect := diagram.map_Product_Rect[parentProduct]
		endRect := diagram.map_Product_Rect[subProduct]

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
		rect := svgRect(
			stager,
			diagram,
			taskShape,
			layer)
		diagram.map_Task_Rect[taskShape.Task] = rect
		diagram.map_SvgRect_TaskShape[rect] = taskShape
	}

	for _, taskCompositionShape := range diagram.TaskComposition_Shapes {
		_ = taskCompositionShape
		subTask := taskCompositionShape.Task
		parentTask := subTask.parentTask

		if subTask == nil || parentTask == nil {
			log.Panic("There should be a subTask and a parentTask")
		}

		startRect := diagram.map_Task_Rect[parentTask]
		endRect := diagram.map_Task_Rect[subTask]

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

		startRect := diagram.map_Product_Rect[product]
		endRect := diagram.map_Task_Rect[task]

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

		startRect := diagram.map_Task_Rect[task]
		endRect := diagram.map_Product_Rect[product]

		svgAssociationLink(
			stager,
			startRect, endRect,
			taskOutputShape,
			task,
			layer,
			true,
		)
	}

	for _, noteShape := range diagram.Note_Shapes {
		note := noteShape.Note

		if note == nil {
			log.Panic("There should be a note")
		}

		rect := svgRect(
			stager,
			diagram,
			noteShape,
			layer)
		rect.RX = 6
		diagram.map_Note_Rect[note] = rect
		diagram.map_SvgRect_NoteShape[rect] = noteShape

		penLogo := new(svg.RectAnchoredPath)
		penLogo.Stroke = svg.Black.ToString()
		penLogo.StrokeWidth = 2
		penLogo.StrokeOpacity = 1
		penLogo.ScalePropotionnally = true
		penLogo.AppliedScaling = 1

		penLogo.Definition = "M 5 16 L 9 20 L 20 9 L 25 0 L 16 5 Z"
		penLogo.X_Offset = 0
		penLogo.Y_Offset = 5
		penLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, penLogo)
		_ = rect
	}

	for _, noteProductShape := range diagram.NoteProductShapes {
		note := noteProductShape.Note
		product := noteProductShape.Product

		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Product_Rect[product]

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteProductShape,
			note,
			layer,
			true,
		)
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.HasEndArrow = false
	}

	for _, noteTaskShape := range diagram.NoteTaskShapes {
		note := noteTaskShape.Note
		task := noteTaskShape.Task
		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Task_Rect[task]

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteTaskShape,
			note,
			layer,
			true,
		)
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.HasEndArrow = false
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
