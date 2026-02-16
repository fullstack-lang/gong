package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

func (stager *Stager) svg() {
	stager.svgStage.Reset()
	stage := stager.stage
	_ = stage
	root := stager.root

	svgStage := stager.svgStage

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

	svgObject := (&svg.SVG{Name: `SVG`})

	diagram.map_Product_Rect = make(map[*Product]*svg.Rect)
	diagram.map_Task_Rect = make(map[*Task]*svg.Rect)
	diagram.map_Note_Rect = make(map[*Note]*svg.Rect)
	diagram.map_Resource_Rect = make(map[*Resource]*svg.Rect)

	diagram.map_SvgRect_ProductShape = make(map[*svg.Rect]*ProductShape)
	diagram.map_SvgRect_TaskShape = make(map[*svg.Rect]*TaskShape)
	diagram.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	diagram.map_SvgRect_ResourceShape = make(map[*svg.Rect]*ResourceShape)

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

		if taskShape.Task.IsWithCompletion {
			rectAnchoredPath := new(svg.RectAnchoredPath)
			rect.IsScalingProportionally = false
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

			switch taskShape.Task.Completion {
			case PERCENT_100:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,4.16 v 6.85 h 11.11 v -6.85 z m 0,11.35 v 6.68 h 11.11 v -6.68 z m 0,11.08 v 6.60 h 11.11 v -6.60 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_075:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,15.51 v 6.68 h 11.11 v -6.68 z m 0,11.09 v 6.60 h 11.11 v -6.60 z m 0,11.04 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_050:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.39,26.6 v 6.6 h 11.11 v -6.6 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_025:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z M 6.94,40.34 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_000:
				rectAnchoredPath.Definition = "M 4.99,0 c -2.81,0.28 -4.99,2.65 -4.99,5.53 v 42.95 c 0,3.08 2.51,5.53 5.59,5.53 h 13.88 c 3.08,0 5.53,-2.45 5.53,-5.53 v -42.95 c 0,-3.08 -2.46,-5.53 -5.53,-5.53 h -13.88 c -0.19,0 -0.41,-0.02 -0.60,0 z m 0.60,2.66 h 13.88 c 1.54,0 2.77,1.23 2.77,2.77 v 43.06 c 0,1.54 -1.23,2.77 -2.77,2.77 h -13.88 c -1.54,0 -2.82,-1.23 -2.82,-2.77 v -43.06 c 0,-1.54 1.28,-2.77 2.82,-2.77 z"
			}

			rectAnchoredPath.Name = "Percentage"

			// rectAnchoredPath.ScalePropotionnally = true
			// rectAnchoredPath.AppliedScaling = 0.08

			rectAnchoredPath.Stroke = svg.Lightgray.ToString()
			rectAnchoredPath.StrokeWidth = 2
			rectAnchoredPath.StrokeOpacity = 1

			rectAnchoredPath.Color = svg.Lightgray.ToString()
			rectAnchoredPath.FillOpacity = 0.8

			distanceFromBorder := 10.0
			iconWidth := 25.0

			rectAnchoredPath.X_Offset = distanceFromBorder
			rectAnchoredPath.Y_Offset = distanceFromBorder

			rectAnchoredPath.RectAnchorType = svg.RECT_TOP_LEFT

			// shift the text on the right
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth) {
				title.Content = strutils.WrapString(title.Content, int((rect.Width-(distanceFromBorder+iconWidth))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
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
			log.Println("There should be a note")
			continue
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

	for _, resourceShape := range diagram.Resource_Shapes {
		rect := svgRect(
			stager,
			diagram,
			resourceShape,
			layer)
		diagram.map_Resource_Rect[resourceShape.Resource] = rect
		diagram.map_SvgRect_ResourceShape[rect] = resourceShape
		{
			rectAnchoredPath := new(svg.RectAnchoredPath)
			rect.IsScalingProportionally = false
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

			rectAnchoredPath.Definition = "M 30, 11 a 10, 10 0 1, 1 -20, 0 a 10, 10 0 1, 1 20, 0 M 1, 49 h 38 v -9 c 0 -6.6 -5.4 -12 -12 -12 H 13 c -6.6 0 -12 5.4 -12 12 z"

			rectAnchoredPath.Name = "Person"

			rectAnchoredPath.Stroke = svg.Lightgray.ToString()
			rectAnchoredPath.StrokeWidth = 4
			rectAnchoredPath.StrokeOpacity = 1

			rectAnchoredPath.FillOpacity = 0.0

			distanceFromBorder := 10.0
			iconWidth := 40.0

			rectAnchoredPath.X_Offset = distanceFromBorder
			rectAnchoredPath.Y_Offset = distanceFromBorder

			rectAnchoredPath.RectAnchorType = svg.RECT_TOP_LEFT

			// shift the text on the right
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth) {
				title.Content = strutils.WrapString(title.Content, int((rect.Width-(distanceFromBorder+iconWidth))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0

		}
	}

	for _, noteResourceShape := range diagram.NoteResourceShapes {
		note := noteResourceShape.Note
		resource := noteResourceShape.Resource
		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Resource_Rect[resource]
		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteResourceShape,
			note,
			layer,
			true,
		)
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.HasEndArrow = false
	}

	for _, resourceCompositionShape := range diagram.ResourceComposition_Shapes {
		subResource := resourceCompositionShape.Resource
		parentResource := subResource.parentResource
		if subResource == nil || parentResource == nil {
			log.Panic("There should be a subResource and a parentResource")
		}
		startRect := diagram.map_Resource_Rect[parentResource]
		endRect := diagram.map_Resource_Rect[subResource]
		svgAssociationLink(
			stager,
			startRect, endRect,
			resourceCompositionShape,
			parentResource,
			layer,
			false,
		)
	}

	for _, resourceTaskShape := range diagram.ResourceTaskShapes {
		resource := resourceTaskShape.Resource
		task := resourceTaskShape.Task
		if resource == nil || task == nil {
			log.Panic("There should be a resource and a task")
		}
		startRect := diagram.map_Resource_Rect[resource]
		endRect := diagram.map_Task_Rect[task]
		svgAssociationLink(
			stager,
			startRect, endRect,
			resourceTaskShape,
			resource,
			layer,
			true,
		)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
